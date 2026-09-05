# Dynamic queries

This is where sqlc and dsqlsea diverge, as at the time of when this was forked, `sqlc` did not (or currently does not) support dynamic queries. Dynamic queries are often used for filtering, such as with a list endpoint in a REST API. Sometimes we need `select` statements where the `where` and `order by` may change from user to user.


```sql schema.sql
--schema.sql
CREATE TABLE records (
    id          BIGSERIAL PRIMARY KEY,
    tenant_id   BIGINT NOT NULL,
    name        TEXT NOT NULL,
    age         INT NOT NULL,
    status      TEXT NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);
```

To generate a dynamic query, annotate a query with `:dynamicmany` or `:dynamicone`, and then specify which columns from the select query will be optional parameters with `@dynamic` or `@dynamic-sort`. These columns will be required to use the macro `sqlc.arg(<column_name>)`.

```sql query.sql
-- query.sql

-- name: ListRecords :dynamicmany
-- @dynamic name
-- @dynamic age
-- @dynamic-sort name, age, created_at
SELECT id, name, age, created_at
FROM records
WHERE tenant_id = sqlc.arg(tenant_id)
    AND name = sqlc.arg(name)
    AND age > sqlc.arg(age);
```

We can then use the generated code as so:

```go
package main

import (
    "context"
	"errors"
    "log/slog"
    "os"

	"github.com/jackc/pgx/v5"

    "db"
)

func main() {
    ctx := context.Background()
    uri := "postgres://postgres:supersecure@localhost:5432/postgres"

    pdb, err := pgx.Connect(ctx, uri)
    if err != nil {
        slog.Error("failed to connect to postgres")
        os.Exit(1)
    }
    defer pdb.Close(ctx)

    q := db.New(pdb)

    tenantID := 1
    overTwentyFive, err := q.ListRecords(ctx, tenantID, ListRecordsOpts{}.Age(25).OrderBy(ListRecordsOrderByAge, true))
    if err != nil {
        slog.Error("failed listing users over the age of 25", slog.Int("tenant_id", tenantID))
    }
    slog.Info(
        "users over the age of 25 sorted oldest to youngest", 
        slog.Any("users", overTwentyFive), 
        slog.Int("tenant_id", tenantID),
    )

    everyAlice, err := q.ListRecords(ctx, tenantID, ListRecordsOpts{}.Name("Alice"))
    if err != nil {
        slog.Error("failed listing all Alices", slog.Int("tenant_id", tenantID))
    }
    slog.Info("all alices for this tenant", slog.Any("users", everyAlice))
}
```

The generated `go` code will output as:
```go
package db

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

type DBTX interface {
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	QueryRow(context.Context, string, ...interface{}) pgx.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

type Queries struct {
	db DBTX
}

func (q *Queries) WithTx(tx pgx.Tx) *Queries {
	return &Queries{
		db: tx,
	}
}

type Record struct {
	ID        int64
	TenantID  int64
	Name      string
	Age       int32
	Status    string
	CreatedAt pgtype.Timestamptz
}

const listRecords = `-- name: ListRecords :many
SELECT id, name, age, created_at FROM records
WHERE tenant_id = $1
`

type ListRecordsRow struct {
	ID        int64
	Name      string
	Age       int32
	CreatedAt pgtype.Timestamptz
}

type ListRecordsOpts struct {
	name    *string
	age     *int32
	orderBy []string
}

func (o ListRecordsOpts) Name(v string) ListRecordsOpts {
	o.name = &v
	return o
}

func (o ListRecordsOpts) Age(v int32) ListRecordsOpts {
	o.age = &v
	return o
}

type ListRecordsOrderByColumn string

const (
	ListRecordsOrderByName      ListRecordsOrderByColumn = "name"
	ListRecordsOrderByAge       ListRecordsOrderByColumn = "age"
	ListRecordsOrderByCreatedAt ListRecordsOrderByColumn = "created_at"
)

func (o ListRecordsOpts) OrderBy(col ListRecordsOrderByColumn, desc bool) ListRecordsOpts {
	dir := " ASC"
	if desc {
		dir = " DESC"
	}
	o.orderBy = append(o.orderBy, string(col)+dir)
	return o
}

func (q *Queries) ListRecords(ctx context.Context, tenantID int64, opts ListRecordsOpts) ([]ListRecordsRow, error) {
	query := listRecords
	queryParams := []interface{}{tenantID}
	conds := make([]string, 0, 2)
	n := 1
	_ = n
	if opts.name != nil {
		n++
		conds = append(conds, fmt.Sprintf("records.name = $%d", n))
		queryParams = append(queryParams, *opts.name)
	}
	if opts.age != nil {
		n++
		conds = append(conds, fmt.Sprintf("records.age > $%d", n))
		queryParams = append(queryParams, *opts.age)
	}
	if len(conds) > 0 {
		query += " AND " + strings.Join(conds, " AND ")
	}
	if len(opts.orderBy) > 0 {
		query += " ORDER BY " + strings.Join(opts.orderBy, ", ")
	}
	rows, err := q.db.Query(ctx, query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListRecordsRow
	for rows.Next() {
		var i ListRecordsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Age,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
```
