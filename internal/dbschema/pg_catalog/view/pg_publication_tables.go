//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package view

import (
	"github.com/go-jet/jet/v2/postgres"
)

var PgPublicationTables = newPgPublicationTablesTable()

type pgPublicationTablesTable struct {
	postgres.Table

	//Columns
	Pubname    postgres.ColumnString
	Schemaname postgres.ColumnString
	Tablename  postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type PgPublicationTablesTable struct {
	pgPublicationTablesTable

	EXCLUDED pgPublicationTablesTable
}

// AS creates new PgPublicationTablesTable with assigned alias
func (a *PgPublicationTablesTable) AS(alias string) *PgPublicationTablesTable {
	aliasTable := newPgPublicationTablesTable()
	aliasTable.Table.AS(alias)
	return aliasTable
}

func newPgPublicationTablesTable() *PgPublicationTablesTable {
	return &PgPublicationTablesTable{
		pgPublicationTablesTable: newPgPublicationTablesTableImpl("pg_catalog", "pg_publication_tables"),
		EXCLUDED:                 newPgPublicationTablesTableImpl("", "excluded"),
	}
}

func newPgPublicationTablesTableImpl(schemaName, tableName string) pgPublicationTablesTable {
	var (
		PubnameColumn    = postgres.StringColumn("pubname")
		SchemanameColumn = postgres.StringColumn("schemaname")
		TablenameColumn  = postgres.StringColumn("tablename")
		allColumns       = postgres.ColumnList{PubnameColumn, SchemanameColumn, TablenameColumn}
		mutableColumns   = postgres.ColumnList{PubnameColumn, SchemanameColumn, TablenameColumn}
	)

	return pgPublicationTablesTable{
		Table: postgres.NewTable(schemaName, tableName, allColumns...),

		//Columns
		Pubname:    PubnameColumn,
		Schemaname: SchemanameColumn,
		Tablename:  TablenameColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
