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

var PgTsConfigMap = newPgTsConfigMapTable()

type PgTsConfigMapTable struct {
	postgres.Table

	//Columns
	Mapcfg       postgres.ColumnString
	Maptokentype postgres.ColumnInteger
	Mapseqno     postgres.ColumnInteger
	Mapdict      postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

// creates new PgTsConfigMapTable with assigned alias
func (a *PgTsConfigMapTable) AS(alias string) *PgTsConfigMapTable {
	aliasTable := newPgTsConfigMapTable()

	aliasTable.Table.AS(alias)

	return aliasTable
}

func newPgTsConfigMapTable() *PgTsConfigMapTable {
	var (
		MapcfgColumn       = postgres.StringColumn("mapcfg")
		MaptokentypeColumn = postgres.IntegerColumn("maptokentype")
		MapseqnoColumn     = postgres.IntegerColumn("mapseqno")
		MapdictColumn      = postgres.StringColumn("mapdict")
	)

	return &PgTsConfigMapTable{
		Table: postgres.NewTable("pg_catalog", "pg_ts_config_map", MapcfgColumn, MaptokentypeColumn, MapseqnoColumn, MapdictColumn),

		//Columns
		Mapcfg:       MapcfgColumn,
		Maptokentype: MaptokentypeColumn,
		Mapseqno:     MapseqnoColumn,
		Mapdict:      MapdictColumn,

		AllColumns:     postgres.ColumnList{MapcfgColumn, MaptokentypeColumn, MapseqnoColumn, MapdictColumn},
		MutableColumns: postgres.ColumnList{MapcfgColumn, MaptokentypeColumn, MapseqnoColumn, MapdictColumn},
	}
}