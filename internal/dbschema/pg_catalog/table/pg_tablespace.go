//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var PgTablespace = newPgTablespaceTable()

type pgTablespaceTable struct {
	postgres.Table

	//Columns
	Oid        postgres.ColumnString
	Spcname    postgres.ColumnString
	Spcowner   postgres.ColumnString
	Spcacl     postgres.ColumnString
	Spcoptions postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type PgTablespaceTable struct {
	pgTablespaceTable

	EXCLUDED pgTablespaceTable
}

// AS creates new PgTablespaceTable with assigned alias
func (a *PgTablespaceTable) AS(alias string) *PgTablespaceTable {
	aliasTable := newPgTablespaceTable()
	aliasTable.Table.AS(alias)
	return aliasTable
}

func newPgTablespaceTable() *PgTablespaceTable {
	return &PgTablespaceTable{
		pgTablespaceTable: newPgTablespaceTableImpl("pg_catalog", "pg_tablespace"),
		EXCLUDED:          newPgTablespaceTableImpl("", "excluded"),
	}
}

func newPgTablespaceTableImpl(schemaName, tableName string) pgTablespaceTable {
	var (
		OidColumn        = postgres.StringColumn("oid")
		SpcnameColumn    = postgres.StringColumn("spcname")
		SpcownerColumn   = postgres.StringColumn("spcowner")
		SpcaclColumn     = postgres.StringColumn("spcacl")
		SpcoptionsColumn = postgres.StringColumn("spcoptions")
		allColumns       = postgres.ColumnList{OidColumn, SpcnameColumn, SpcownerColumn, SpcaclColumn, SpcoptionsColumn}
		mutableColumns   = postgres.ColumnList{OidColumn, SpcnameColumn, SpcownerColumn, SpcaclColumn, SpcoptionsColumn}
	)

	return pgTablespaceTable{
		Table: postgres.NewTable(schemaName, tableName, allColumns...),

		//Columns
		Oid:        OidColumn,
		Spcname:    SpcnameColumn,
		Spcowner:   SpcownerColumn,
		Spcacl:     SpcaclColumn,
		Spcoptions: SpcoptionsColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
