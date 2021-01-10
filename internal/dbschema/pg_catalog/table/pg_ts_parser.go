//
// Code generated by go-jet DO NOT EDIT.
// Generated at Tuesday, 12-May-20 07:59:32 BST
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/postgres"
)

var PgTsParser = newPgTsParserTable()

type PgTsParserTable struct {
	postgres.Table

	//Columns
	Oid          postgres.ColumnString
	Prsname      postgres.ColumnString
	Prsnamespace postgres.ColumnString
	Prsstart     postgres.ColumnString
	Prstoken     postgres.ColumnString
	Prsend       postgres.ColumnString
	Prsheadline  postgres.ColumnString
	Prslextype   postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

// creates new PgTsParserTable with assigned alias
func (a *PgTsParserTable) AS(alias string) *PgTsParserTable {
	aliasTable := newPgTsParserTable()

	aliasTable.Table.AS(alias)

	return aliasTable
}

func newPgTsParserTable() *PgTsParserTable {
	var (
		OidColumn          = postgres.StringColumn("oid")
		PrsnameColumn      = postgres.StringColumn("prsname")
		PrsnamespaceColumn = postgres.StringColumn("prsnamespace")
		PrsstartColumn     = postgres.StringColumn("prsstart")
		PrstokenColumn     = postgres.StringColumn("prstoken")
		PrsendColumn       = postgres.StringColumn("prsend")
		PrsheadlineColumn  = postgres.StringColumn("prsheadline")
		PrslextypeColumn   = postgres.StringColumn("prslextype")
	)

	return &PgTsParserTable{
		Table: postgres.NewTable("pg_catalog", "pg_ts_parser", OidColumn, PrsnameColumn, PrsnamespaceColumn, PrsstartColumn, PrstokenColumn, PrsendColumn, PrsheadlineColumn, PrslextypeColumn),

		//Columns
		Oid:          OidColumn,
		Prsname:      PrsnameColumn,
		Prsnamespace: PrsnamespaceColumn,
		Prsstart:     PrsstartColumn,
		Prstoken:     PrstokenColumn,
		Prsend:       PrsendColumn,
		Prsheadline:  PrsheadlineColumn,
		Prslextype:   PrslextypeColumn,

		AllColumns:     postgres.ColumnList{OidColumn, PrsnameColumn, PrsnamespaceColumn, PrsstartColumn, PrstokenColumn, PrsendColumn, PrsheadlineColumn, PrslextypeColumn},
		MutableColumns: postgres.ColumnList{OidColumn, PrsnameColumn, PrsnamespaceColumn, PrsstartColumn, PrstokenColumn, PrsendColumn, PrsheadlineColumn, PrslextypeColumn},
	}
}