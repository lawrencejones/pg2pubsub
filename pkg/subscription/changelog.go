package subscription

import (
	"github.com/lawrencejones/pgsink/pkg/changelog"
	"github.com/lawrencejones/pgsink/pkg/decode"
	"github.com/lawrencejones/pgsink/pkg/logical"

	kitlog "github.com/go-kit/kit/log"
)

// BuildChangelog produces a stream of changelog entries from raw logical messages
// produced by a subscription.
func BuildChangelog(logger kitlog.Logger, decoder decode.Decoder, stream *Stream) changelog.Changelog {
	output := make(changelog.Changelog)

	// TODO: If this is ever modified to marshal entries in parallel, this will complicate
	// any acknowledgement pipeline. Double check assumptions about acknowledgement order
	// before removing ordering.
	go func() {
		registry, raw := logical.BuildRegistry(logger, decoder, stream.Messages())
		for msg := range Sequence(raw) {
			timestamp, lsn := msg.Begin.Timestamp, msg.Begin.LSN
			switch entry := msg.Entry.(type) {
			case *logical.Relation:
				schema := changelog.SchemaFromRelation(timestamp, &lsn, entry)
				output <- changelog.Entry{Schema: &schema}
			case *logical.Insert, *logical.Update, *logical.Delete:
				modification := &changelog.Modification{
					Timestamp: timestamp,
					LSN:       &lsn,
				}

				var (
					relation *logical.Relation
					err      error
				)
				relation, modification.Before, modification.After, err = registry.Marshal(msg.Entry)
				// This shouldn't panic, but it does for now. The primary reason this might fail
				// is the decoder being unable to recognise a type, which means we can't continue
				// subscribing.
				//
				// We need better ways of recovering from this, such as removing the problem table
				// from the subscription.
				if err != nil {
					panic(err.Error())
				}

				modification.Namespace, modification.Name = relation.Namespace, relation.Name

				output <- changelog.Entry{Modification: modification}
			default:
				// ignore this message type
			}
		}

		close(output)
	}()

	return output
}

// SequencedMessage wraps logical messages with the begin message associated with the
// transaction that the message was contained within.
type SequencedMessage struct {
	Begin    logical.Begin
	Sequence uint64
	Entry    interface{} // the original message
}

// Sequence receives a channel containing logical replication messages and produces
// a channel which annotates each message with commit information. Sequenced structs can
// be tracked back to a specific LSN, and logically ordered by sequence number, ensuring
// we can detect the authoriative row value even if the same row is updated many times
// within the same transaction.
//
// This will almost always be used like so:
//
//     Sequence(stream.Messages())
//
// Where stream is an active Stream.
func Sequence(messages <-chan interface{}) <-chan SequencedMessage {
	output := make(chan SequencedMessage)

	go func() {
		var currentTransaction *logical.Begin
		var sequence uint64

		for msg := range messages {
			switch msg := msg.(type) {
			case *logical.Begin:
				currentTransaction = msg
				sequence = 0
			case *logical.Commit:
				currentTransaction = nil
			default:
				sequence++
				output <- SequencedMessage{
					Begin:    *currentTransaction,
					Sequence: sequence,
					Entry:    msg,
				}
			}
		}

		close(output)
	}()

	return output
}
