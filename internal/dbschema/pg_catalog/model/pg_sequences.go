//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type PgSequences struct {
	Schemaname    *string
	Sequencename  *string
	Sequenceowner *string
	DataType      *string
	StartValue    *int64
	MinValue      *int64
	MaxValue      *int64
	IncrementBy   *int64
	Cycle         *bool
	CacheSize     *int64
	LastValue     *int64
}
