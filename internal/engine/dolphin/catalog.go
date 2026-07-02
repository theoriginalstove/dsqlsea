package dolphin

import (
	"github.com/theoriginalstove/dsqlsea/internal/sql/catalog"
)

func NewCatalog() *catalog.Catalog {
	def := "public" // TODO: What is the default database for MySQL?
	return &catalog.Catalog{
		DefaultSchema: def,
		Schemas: []*catalog.Schema{
			defaultSchema(def),
		},
		Extensions: map[string]struct{}{},
	}
}
