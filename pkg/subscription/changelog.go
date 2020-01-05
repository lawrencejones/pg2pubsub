package subscription

import (
	"fmt"

	kitlog "github.com/go-kit/kit/log"
	"github.com/lawrencejones/pg2sink/pkg/changelog"
	"github.com/lawrencejones/pg2sink/pkg/logical"
)

// BuildChangelog produces a stream of changelog entries from raw logical messages
// produced by a subscription.
func BuildChangelog(logger kitlog.Logger, raw <-chan interface{}) changelog.Changelog {
	output := make(changelog.Changelog)

	// TODO: If this is ever modified to marshal entries in parallel, this will complicate
	// any acknowledgement pipeline. Double check assumptions about acknowledgement order
	// before removing ordering.
	go func() {
		registry, raw := logical.BuildRegistry(logger, raw)
		for msg := range Sequence(raw) {
			timestamp, lsn := msg.Begin.Timestamp, msg.Begin.LSN
			if relation, ok := msg.Entry.(*logical.Relation); ok {
				schema := changelog.SchemaFromRelation(timestamp, &lsn, relation)
				output <- changelog.Entry{Schema: &schema}
			} else {
				modification := &changelog.Modification{
					Timestamp: timestamp,
					LSN:       &lsn,
				}

				var relation *logical.Relation
				relation, modification.Before, modification.After = registry.Marshal(msg.Entry)
				modification.Namespace = fmt.Sprintf("%s.%s", relation.Namespace, relation.Name)

				output <- changelog.Entry{Modification: modification}
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
//     Sequence(sub.Receive())
//
// Where sub is a Subscription.
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
