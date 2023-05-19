package helpers

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func BootstrapDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		return nil, err
	}

	return db, nil
}
