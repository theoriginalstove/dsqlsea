package clickhouse

import (
	"github.com/theoriginalstove/dsqlsea/internal/sql/catalog"
)

func defaultSchema(name string) *catalog.Schema {
	return &catalog.Schema{Name: name}
}
