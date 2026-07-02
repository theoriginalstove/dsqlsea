package compiler

import (
	"github.com/theoriginalstove/dsqlsea/internal/sql/catalog"
)

type Result struct {
	Catalog *catalog.Catalog
	Queries []*Query
}
