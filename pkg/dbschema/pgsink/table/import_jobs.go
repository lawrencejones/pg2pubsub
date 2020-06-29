//
// Code generated by go-jet DO NOT EDIT.
// Generated at Monday, 25-May-20 13:35:21 BST
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/postgres"
)

var ImportJobs = newImportJobsTable()

type ImportJobsTable struct {
	postgres.Table

	//Columns
	ID             postgres.ColumnInteger
	SubscriptionID postgres.ColumnString
	TableName      postgres.ColumnString
	Cursor         postgres.ColumnString
	CompletedAt    postgres.ColumnTimestampz
	ExpiredAt      postgres.ColumnTimestampz
	UpdatedAt      postgres.ColumnTimestampz
	CreatedAt      postgres.ColumnTimestampz
	Error          postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

// creates new ImportJobsTable with assigned alias
func (a *ImportJobsTable) AS(alias string) *ImportJobsTable {
	aliasTable := newImportJobsTable()

	aliasTable.Table.AS(alias)

	return aliasTable
}

func newImportJobsTable() *ImportJobsTable {
	var (
		IDColumn             = postgres.IntegerColumn("id")
		SubscriptionIDColumn = postgres.StringColumn("subscription_id")
		TableNameColumn      = postgres.StringColumn("table_name")
		CursorColumn         = postgres.StringColumn("cursor")
		CompletedAtColumn    = postgres.TimestampzColumn("completed_at")
		ExpiredAtColumn      = postgres.TimestampzColumn("expired_at")
		UpdatedAtColumn      = postgres.TimestampzColumn("updated_at")
		CreatedAtColumn      = postgres.TimestampzColumn("created_at")
		ErrorColumn          = postgres.StringColumn("error")
	)

	return &ImportJobsTable{
		Table: postgres.NewTable("pgsink", "import_jobs", IDColumn, SubscriptionIDColumn, TableNameColumn, CursorColumn, CompletedAtColumn, ExpiredAtColumn, UpdatedAtColumn, CreatedAtColumn, ErrorColumn),

		//Columns
		ID:             IDColumn,
		SubscriptionID: SubscriptionIDColumn,
		TableName:      TableNameColumn,
		Cursor:         CursorColumn,
		CompletedAt:    CompletedAtColumn,
		ExpiredAt:      ExpiredAtColumn,
		UpdatedAt:      UpdatedAtColumn,
		CreatedAt:      CreatedAtColumn,
		Error:          ErrorColumn,

		AllColumns:     postgres.ColumnList{IDColumn, SubscriptionIDColumn, TableNameColumn, CursorColumn, CompletedAtColumn, ExpiredAtColumn, UpdatedAtColumn, CreatedAtColumn, ErrorColumn},
		MutableColumns: postgres.ColumnList{SubscriptionIDColumn, TableNameColumn, CursorColumn, CompletedAtColumn, ExpiredAtColumn, UpdatedAtColumn, CreatedAtColumn, ErrorColumn},
	}
}