package postgresql

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDB(connString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	return db, nil
}
