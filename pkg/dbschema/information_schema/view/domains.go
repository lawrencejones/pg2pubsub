//
// Code generated by go-jet DO NOT EDIT.
// Generated at Tuesday, 12-May-20 09:15:06 BST
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package view

import (
	"github.com/go-jet/jet/postgres"
)

var Domains = newDomainsTable()

type DomainsTable struct {
	postgres.Table

	//Columns
	DomainCatalog          postgres.ColumnString
	DomainSchema           postgres.ColumnString
	DomainName             postgres.ColumnString
	DataType               postgres.ColumnString
	CharacterMaximumLength postgres.ColumnInteger
	CharacterOctetLength   postgres.ColumnInteger
	CharacterSetCatalog    postgres.ColumnString
	CharacterSetSchema     postgres.ColumnString
	CharacterSetName       postgres.ColumnString
	CollationCatalog       postgres.ColumnString
	CollationSchema        postgres.ColumnString
	CollationName          postgres.ColumnString
	NumericPrecision       postgres.ColumnInteger
	NumericPrecisionRadix  postgres.ColumnInteger
	NumericScale           postgres.ColumnInteger
	DatetimePrecision      postgres.ColumnInteger
	IntervalType           postgres.ColumnString
	IntervalPrecision      postgres.ColumnInteger
	DomainDefault          postgres.ColumnString
	UdtCatalog             postgres.ColumnString
	UdtSchema              postgres.ColumnString
	UdtName                postgres.ColumnString
	ScopeCatalog           postgres.ColumnString
	ScopeSchema            postgres.ColumnString
	ScopeName              postgres.ColumnString
	MaximumCardinality     postgres.ColumnInteger
	DtdIdentifier          postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

// creates new DomainsTable with assigned alias
func (a *DomainsTable) AS(alias string) *DomainsTable {
	aliasTable := newDomainsTable()

	aliasTable.Table.AS(alias)

	return aliasTable
}

func newDomainsTable() *DomainsTable {
	var (
		DomainCatalogColumn          = postgres.StringColumn("domain_catalog")
		DomainSchemaColumn           = postgres.StringColumn("domain_schema")
		DomainNameColumn             = postgres.StringColumn("domain_name")
		DataTypeColumn               = postgres.StringColumn("data_type")
		CharacterMaximumLengthColumn = postgres.IntegerColumn("character_maximum_length")
		CharacterOctetLengthColumn   = postgres.IntegerColumn("character_octet_length")
		CharacterSetCatalogColumn    = postgres.StringColumn("character_set_catalog")
		CharacterSetSchemaColumn     = postgres.StringColumn("character_set_schema")
		CharacterSetNameColumn       = postgres.StringColumn("character_set_name")
		CollationCatalogColumn       = postgres.StringColumn("collation_catalog")
		CollationSchemaColumn        = postgres.StringColumn("collation_schema")
		CollationNameColumn          = postgres.StringColumn("collation_name")
		NumericPrecisionColumn       = postgres.IntegerColumn("numeric_precision")
		NumericPrecisionRadixColumn  = postgres.IntegerColumn("numeric_precision_radix")
		NumericScaleColumn           = postgres.IntegerColumn("numeric_scale")
		DatetimePrecisionColumn      = postgres.IntegerColumn("datetime_precision")
		IntervalTypeColumn           = postgres.StringColumn("interval_type")
		IntervalPrecisionColumn      = postgres.IntegerColumn("interval_precision")
		DomainDefaultColumn          = postgres.StringColumn("domain_default")
		UdtCatalogColumn             = postgres.StringColumn("udt_catalog")
		UdtSchemaColumn              = postgres.StringColumn("udt_schema")
		UdtNameColumn                = postgres.StringColumn("udt_name")
		ScopeCatalogColumn           = postgres.StringColumn("scope_catalog")
		ScopeSchemaColumn            = postgres.StringColumn("scope_schema")
		ScopeNameColumn              = postgres.StringColumn("scope_name")
		MaximumCardinalityColumn     = postgres.IntegerColumn("maximum_cardinality")
		DtdIdentifierColumn          = postgres.StringColumn("dtd_identifier")
	)

	return &DomainsTable{
		Table: postgres.NewTable("information_schema", "domains", DomainCatalogColumn, DomainSchemaColumn, DomainNameColumn, DataTypeColumn, CharacterMaximumLengthColumn, CharacterOctetLengthColumn, CharacterSetCatalogColumn, CharacterSetSchemaColumn, CharacterSetNameColumn, CollationCatalogColumn, CollationSchemaColumn, CollationNameColumn, NumericPrecisionColumn, NumericPrecisionRadixColumn, NumericScaleColumn, DatetimePrecisionColumn, IntervalTypeColumn, IntervalPrecisionColumn, DomainDefaultColumn, UdtCatalogColumn, UdtSchemaColumn, UdtNameColumn, ScopeCatalogColumn, ScopeSchemaColumn, ScopeNameColumn, MaximumCardinalityColumn, DtdIdentifierColumn),

		//Columns
		DomainCatalog:          DomainCatalogColumn,
		DomainSchema:           DomainSchemaColumn,
		DomainName:             DomainNameColumn,
		DataType:               DataTypeColumn,
		CharacterMaximumLength: CharacterMaximumLengthColumn,
		CharacterOctetLength:   CharacterOctetLengthColumn,
		CharacterSetCatalog:    CharacterSetCatalogColumn,
		CharacterSetSchema:     CharacterSetSchemaColumn,
		CharacterSetName:       CharacterSetNameColumn,
		CollationCatalog:       CollationCatalogColumn,
		CollationSchema:        CollationSchemaColumn,
		CollationName:          CollationNameColumn,
		NumericPrecision:       NumericPrecisionColumn,
		NumericPrecisionRadix:  NumericPrecisionRadixColumn,
		NumericScale:           NumericScaleColumn,
		DatetimePrecision:      DatetimePrecisionColumn,
		IntervalType:           IntervalTypeColumn,
		IntervalPrecision:      IntervalPrecisionColumn,
		DomainDefault:          DomainDefaultColumn,
		UdtCatalog:             UdtCatalogColumn,
		UdtSchema:              UdtSchemaColumn,
		UdtName:                UdtNameColumn,
		ScopeCatalog:           ScopeCatalogColumn,
		ScopeSchema:            ScopeSchemaColumn,
		ScopeName:              ScopeNameColumn,
		MaximumCardinality:     MaximumCardinalityColumn,
		DtdIdentifier:          DtdIdentifierColumn,

		AllColumns:     postgres.ColumnList{DomainCatalogColumn, DomainSchemaColumn, DomainNameColumn, DataTypeColumn, CharacterMaximumLengthColumn, CharacterOctetLengthColumn, CharacterSetCatalogColumn, CharacterSetSchemaColumn, CharacterSetNameColumn, CollationCatalogColumn, CollationSchemaColumn, CollationNameColumn, NumericPrecisionColumn, NumericPrecisionRadixColumn, NumericScaleColumn, DatetimePrecisionColumn, IntervalTypeColumn, IntervalPrecisionColumn, DomainDefaultColumn, UdtCatalogColumn, UdtSchemaColumn, UdtNameColumn, ScopeCatalogColumn, ScopeSchemaColumn, ScopeNameColumn, MaximumCardinalityColumn, DtdIdentifierColumn},
		MutableColumns: postgres.ColumnList{DomainCatalogColumn, DomainSchemaColumn, DomainNameColumn, DataTypeColumn, CharacterMaximumLengthColumn, CharacterOctetLengthColumn, CharacterSetCatalogColumn, CharacterSetSchemaColumn, CharacterSetNameColumn, CollationCatalogColumn, CollationSchemaColumn, CollationNameColumn, NumericPrecisionColumn, NumericPrecisionRadixColumn, NumericScaleColumn, DatetimePrecisionColumn, IntervalTypeColumn, IntervalPrecisionColumn, DomainDefaultColumn, UdtCatalogColumn, UdtSchemaColumn, UdtNameColumn, ScopeCatalogColumn, ScopeSchemaColumn, ScopeNameColumn, MaximumCardinalityColumn, DtdIdentifierColumn},
	}
}