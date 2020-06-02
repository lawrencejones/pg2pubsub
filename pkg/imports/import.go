package imports

import (
	"context"
	"fmt"
	"strings"

	"github.com/lawrencejones/pgsink/pkg/dbschema/pgsink/model"
	"github.com/lawrencejones/pgsink/pkg/logical"

	kitlog "github.com/go-kit/kit/log"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"go.opencensus.io/trace"
)

// querier allows each helper to accept either transaction or connection objects
type querier interface {
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

// Import is built for each job in the database, having resolved contextual information
// that can help run the job from the database whenever the job was enqueued.
type Import struct {
	TableName         string
	PrimaryKey        string
	PrimaryKeyScanner logical.ValueScanner
	Relation          *logical.Relation
	Scanners          []interface{}
	Cursor            interface{}
}

// Build queries the database for information required to perform an import, given an
// import job to process.
func Build(ctx context.Context, logger kitlog.Logger, tx querier, job model.ImportJobs) (*Import, error) {
	// We should query for the primary key as the first thing we do, as this may fail if the
	// table is misconfigured. It's better to fail here, before we've pushed anything into
	// the changelog, than after pushing the schema when we discover the table is
	// incompatible.
	logger.Log("event", "lookup_primary_key", "msg", "querying Postgres for relations primary key column")
	primaryKey, err := getPrimaryKeyColumn(ctx, tx, job.TableName)
	if err != nil {
		return nil, fmt.Errorf("failed to lookup primary key: %w", err)
	}

	logger.Log("event", "build_relation", "msg", "querying Postgres for relation type information")
	relation, err := buildRelation(ctx, tx, job.TableName, primaryKey)
	if err != nil {
		return nil, fmt.Errorf("failed to build relation for table: %w", err)
	}

	// Build scanners for decoding column types. We'll need the primary key scanner for
	// interpreting the cursor.
	primaryKeyScanner, scanners := buildScanners(relation, primaryKey)

	// We need to translate the import_jobs.cursor value, which is text, into a type that
	// will be supported for querying into the table. We can use the primaryKeyScanner for
	// this, which ensures we reliably encode/decode Postgres types.
	var cursor interface{}
	if job.Cursor != nil {
		if err := primaryKeyScanner.Scan(*job.Cursor); err != nil {
			return nil, fmt.Errorf("incompatible cursor in import_jobs table: %w", err)
		}

		cursor = primaryKeyScanner.Get()
	}

	cfg := &Import{
		TableName:         job.TableName,
		PrimaryKey:        primaryKey,
		PrimaryKeyScanner: primaryKeyScanner,
		Relation:          relation,
		Scanners:          scanners,
		Cursor:            cursor,
	}

	return cfg, nil
}

// buildScanners produces pgx type scanners, returning a scanner for the relation primary
// key and a slice of scanners for the other columns.
func buildScanners(relation *logical.Relation, primaryKey string) (primaryKeyScanner logical.ValueScanner, scanners []interface{}) {
	// Go can't handle splatting non-empty-interface types into a parameter list of
	// empty-interfaces, so we have to construct an interface{} slice of scanners.
	scanners = make([]interface{}, len(relation.Columns))
	for idx, column := range relation.Columns {
		scanner := logical.TypeForOID(column.Type)
		scanners[idx] = scanner

		// We'll need this scanner to convert the cursor value between what the table accepts
		// and what we'll store in import_jobs
		if column.Name == primaryKey {
			primaryKeyScanner = scanner
		}
	}

	return
}

// buildRelation generates the logical.Relation structure by querying Postgres catalog
// tables. Importantly, this populates the relation.Columns slice, providing type
// information that can later be used to marshal Golang types.
func buildRelation(ctx context.Context, tx querier, tableName, primaryKeyColumn string) (*logical.Relation, error) {
	ctx, span := trace.StartSpan(ctx, "pkg/imports.buildRelation")
	defer span.End()

	// Eg. oid = 16411, namespace = public, relname = example
	query := `
	select pg_class.oid as oid
	     , nspname as namespace
	     , relname as name
		from pg_class join pg_namespace on pg_class.relnamespace=pg_namespace.oid
	 where pg_class.oid = $1::regclass::oid;
	`

	relation := &logical.Relation{Columns: []logical.Column{}}
	err := tx.QueryRow(ctx, query, tableName).Scan(&relation.ID, &relation.Namespace, &relation.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to identify table namespace and name: %w", err)
	}

	// Eg. name = id, type = 20
	columnQuery := `
	select attname as name
			 , atttypid as type
	  from pg_attribute
	 where attrelid = $1 and attnum > 0 and not attisdropped
	 order by attnum;
	`

	rows, err := tx.Query(ctx, columnQuery, relation.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to query pg_attribute for relation columns: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		column := logical.Column{}
		if err := rows.Scan(&column.Name, &column.Type); err != nil {
			return nil, err
		}

		// We don't strictly require this, but it's asking for trouble to generate
		// logical.Relation structs that have incorrect key metadata.
		if column.Name == primaryKeyColumn {
			column.Key = true
		}

		relation.Columns = append(relation.Columns, column)
	}

	return relation, nil
}

// buildQuery creates a query string for the given relation, with an optional cursor.
// Prepended to the columns is now(), which enables us to timestamp our imported rows to
// the database time.
func buildQuery(relation *logical.Relation, primaryKey string, limit int, cursor interface{}) string {
	columnNames := make([]string, len(relation.Columns))
	for idx, column := range relation.Columns {
		columnNames[idx] = column.Name
	}

	query := fmt.Sprintf(`select now(), %s from %s`, strings.Join(columnNames, ", "), relation.String())
	if cursor != nil {
		query += fmt.Sprintf(` where %s > $1`, primaryKey)
	}
	query += fmt.Sprintf(` order by %s limit %d`, primaryKey, limit)

	return query
}

type multiplePrimaryKeysError []string

func (m multiplePrimaryKeysError) Error() string {
	return fmt.Sprintf("unsupported multiple primary keys: %s", strings.Join(m, ", "))
}

var NoPrimaryKeyError = fmt.Errorf("no primary key found")

// getPrimaryKeyColumn identifies the primary key column of the given table. It only
// supports tables with primary keys, and of those, only single column primary keys.
func getPrimaryKeyColumn(ctx context.Context, tx querier, tableName string) (string, error) {
	ctx, span := trace.StartSpan(ctx, "pkg/imports.getPrimaryKeyColumn")
	defer span.End()

	query := `
	select array_agg(pg_attribute.attname)
	from pg_index join pg_attribute
	on pg_attribute.attrelid = pg_index.indrelid and pg_attribute.attnum = ANY(pg_index.indkey)
	where pg_index.indrelid = $1::regclass
	and pg_index.indisprimary;
	`

	primaryKeysTextArray := pgtype.TextArray{}
	err := tx.QueryRow(ctx, query, tableName).Scan(&primaryKeysTextArray)
	if err != nil {
		return "", err
	}

	var primaryKeys []string
	if err := primaryKeysTextArray.AssignTo(&primaryKeys); err != nil {
		return "", err
	}

	if len(primaryKeys) == 0 {
		return "", NoPrimaryKeyError
	} else if len(primaryKeys) > 1 {
		return "", multiplePrimaryKeysError(primaryKeys)
	}

	return primaryKeys[0], nil
}
