//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type PgStatProgressBasebackup struct {
	Pid                 *int32
	Phase               *string
	BackupTotal         *int64
	BackupStreamed      *int64
	TablespacesTotal    *int64
	TablespacesStreamed *int64
}