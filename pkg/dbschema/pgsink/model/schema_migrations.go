//
// Code generated by go-jet DO NOT EDIT.
// Generated at Saturday, 02-Jan-21 20:38:10 GMT
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"time"
)

type SchemaMigrations struct {
	ID        int32 `sql:"primary_key"`
	VersionID int64
	IsApplied bool
	Tstamp    *time.Time
}
