package helpers

import (
	"database/sql"
)

func CreateDBTables(db *sql.DB) error {
	stmt, err := db.Prepare(
		`CREATE TABLE IF NOT EXISTS quotation (bid TEXT NOT NULL);`,
	)
	if err != nil {
		return err
	}
	defer stmt.Close()

	stmt.Exec()

	return nil
}
