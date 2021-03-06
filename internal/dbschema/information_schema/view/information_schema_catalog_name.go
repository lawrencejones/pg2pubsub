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

var InformationSchemaCatalogName = newInformationSchemaCatalogNameTable()

type informationSchemaCatalogNameTable struct {
	postgres.Table

	//Columns
	CatalogName postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type InformationSchemaCatalogNameTable struct {
	informationSchemaCatalogNameTable

	EXCLUDED informationSchemaCatalogNameTable
}

// AS creates new InformationSchemaCatalogNameTable with assigned alias
func (a *InformationSchemaCatalogNameTable) AS(alias string) *InformationSchemaCatalogNameTable {
	aliasTable := newInformationSchemaCatalogNameTable()
	aliasTable.Table.AS(alias)
	return aliasTable
}

func newInformationSchemaCatalogNameTable() *InformationSchemaCatalogNameTable {
	return &InformationSchemaCatalogNameTable{
		informationSchemaCatalogNameTable: newInformationSchemaCatalogNameTableImpl("information_schema", "information_schema_catalog_name"),
		EXCLUDED:                          newInformationSchemaCatalogNameTableImpl("", "excluded"),
	}
}

func newInformationSchemaCatalogNameTableImpl(schemaName, tableName string) informationSchemaCatalogNameTable {
	var (
		CatalogNameColumn = postgres.StringColumn("catalog_name")
		allColumns        = postgres.ColumnList{CatalogNameColumn}
		mutableColumns    = postgres.ColumnList{CatalogNameColumn}
	)

	return informationSchemaCatalogNameTable{
		Table: postgres.NewTable(schemaName, tableName, allColumns...),

		//Columns
		CatalogName: CatalogNameColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
