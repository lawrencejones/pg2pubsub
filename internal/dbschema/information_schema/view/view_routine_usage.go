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

var ViewRoutineUsage = newViewRoutineUsageTable()

type viewRoutineUsageTable struct {
	postgres.Table

	//Columns
	TableCatalog    postgres.ColumnString
	TableSchema     postgres.ColumnString
	TableName       postgres.ColumnString
	SpecificCatalog postgres.ColumnString
	SpecificSchema  postgres.ColumnString
	SpecificName    postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type ViewRoutineUsageTable struct {
	viewRoutineUsageTable

	EXCLUDED viewRoutineUsageTable
}

// AS creates new ViewRoutineUsageTable with assigned alias
func (a *ViewRoutineUsageTable) AS(alias string) *ViewRoutineUsageTable {
	aliasTable := newViewRoutineUsageTable()
	aliasTable.Table.AS(alias)
	return aliasTable
}

func newViewRoutineUsageTable() *ViewRoutineUsageTable {
	return &ViewRoutineUsageTable{
		viewRoutineUsageTable: newViewRoutineUsageTableImpl("information_schema", "view_routine_usage"),
		EXCLUDED:              newViewRoutineUsageTableImpl("", "excluded"),
	}
}

func newViewRoutineUsageTableImpl(schemaName, tableName string) viewRoutineUsageTable {
	var (
		TableCatalogColumn    = postgres.StringColumn("table_catalog")
		TableSchemaColumn     = postgres.StringColumn("table_schema")
		TableNameColumn       = postgres.StringColumn("table_name")
		SpecificCatalogColumn = postgres.StringColumn("specific_catalog")
		SpecificSchemaColumn  = postgres.StringColumn("specific_schema")
		SpecificNameColumn    = postgres.StringColumn("specific_name")
		allColumns            = postgres.ColumnList{TableCatalogColumn, TableSchemaColumn, TableNameColumn, SpecificCatalogColumn, SpecificSchemaColumn, SpecificNameColumn}
		mutableColumns        = postgres.ColumnList{TableCatalogColumn, TableSchemaColumn, TableNameColumn, SpecificCatalogColumn, SpecificSchemaColumn, SpecificNameColumn}
	)

	return viewRoutineUsageTable{
		Table: postgres.NewTable(schemaName, tableName, allColumns...),

		//Columns
		TableCatalog:    TableCatalogColumn,
		TableSchema:     TableSchemaColumn,
		TableName:       TableNameColumn,
		SpecificCatalog: SpecificCatalogColumn,
		SpecificSchema:  SpecificSchemaColumn,
		SpecificName:    SpecificNameColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
