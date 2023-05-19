package helpers

import (
	"context"
	"database/sql"
)

func InsertIntoDB(ctx context.Context, db *sql.DB, data string) error {
	stmt, err := db.PrepareContext(ctx, "INSERT INTO quotation values(?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(data)
	if err != nil {
		return err
	}

	return nil
}
