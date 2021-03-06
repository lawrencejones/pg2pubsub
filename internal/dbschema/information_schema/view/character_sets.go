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

var CharacterSets = newCharacterSetsTable()

type characterSetsTable struct {
	postgres.Table

	//Columns
	CharacterSetCatalog   postgres.ColumnString
	CharacterSetSchema    postgres.ColumnString
	CharacterSetName      postgres.ColumnString
	CharacterRepertoire   postgres.ColumnString
	FormOfUse             postgres.ColumnString
	DefaultCollateCatalog postgres.ColumnString
	DefaultCollateSchema  postgres.ColumnString
	DefaultCollateName    postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type CharacterSetsTable struct {
	characterSetsTable

	EXCLUDED characterSetsTable
}

// AS creates new CharacterSetsTable with assigned alias
func (a *CharacterSetsTable) AS(alias string) *CharacterSetsTable {
	aliasTable := newCharacterSetsTable()
	aliasTable.Table.AS(alias)
	return aliasTable
}

func newCharacterSetsTable() *CharacterSetsTable {
	return &CharacterSetsTable{
		characterSetsTable: newCharacterSetsTableImpl("information_schema", "character_sets"),
		EXCLUDED:           newCharacterSetsTableImpl("", "excluded"),
	}
}

func newCharacterSetsTableImpl(schemaName, tableName string) characterSetsTable {
	var (
		CharacterSetCatalogColumn   = postgres.StringColumn("character_set_catalog")
		CharacterSetSchemaColumn    = postgres.StringColumn("character_set_schema")
		CharacterSetNameColumn      = postgres.StringColumn("character_set_name")
		CharacterRepertoireColumn   = postgres.StringColumn("character_repertoire")
		FormOfUseColumn             = postgres.StringColumn("form_of_use")
		DefaultCollateCatalogColumn = postgres.StringColumn("default_collate_catalog")
		DefaultCollateSchemaColumn  = postgres.StringColumn("default_collate_schema")
		DefaultCollateNameColumn    = postgres.StringColumn("default_collate_name")
		allColumns                  = postgres.ColumnList{CharacterSetCatalogColumn, CharacterSetSchemaColumn, CharacterSetNameColumn, CharacterRepertoireColumn, FormOfUseColumn, DefaultCollateCatalogColumn, DefaultCollateSchemaColumn, DefaultCollateNameColumn}
		mutableColumns              = postgres.ColumnList{CharacterSetCatalogColumn, CharacterSetSchemaColumn, CharacterSetNameColumn, CharacterRepertoireColumn, FormOfUseColumn, DefaultCollateCatalogColumn, DefaultCollateSchemaColumn, DefaultCollateNameColumn}
	)

	return characterSetsTable{
		Table: postgres.NewTable(schemaName, tableName, allColumns...),

		//Columns
		CharacterSetCatalog:   CharacterSetCatalogColumn,
		CharacterSetSchema:    CharacterSetSchemaColumn,
		CharacterSetName:      CharacterSetNameColumn,
		CharacterRepertoire:   CharacterRepertoireColumn,
		FormOfUse:             FormOfUseColumn,
		DefaultCollateCatalog: DefaultCollateCatalogColumn,
		DefaultCollateSchema:  DefaultCollateSchemaColumn,
		DefaultCollateName:    DefaultCollateNameColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
