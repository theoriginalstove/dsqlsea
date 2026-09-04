package dbmanager

import (
	"fmt"
	"net/url"
)

// MySQLReformatURI rewrites the URI of a database created by CreateDatabase
// into the DSN form the go-sql-driver/mysql driver expects.
func MySQLReformatURI(original string) (string, error) {
	u, err := url.Parse(original)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s@tcp(%s)%s?multiStatements=true&parseTime=true&tls=true", u.User, u.Host, u.Path), nil
}
