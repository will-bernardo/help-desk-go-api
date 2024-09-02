package db

import (
	"database/sql"

	_ "github.com/ncruces/go-sqlite3"
)

func NewSQLiteConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./sqlite.db")
	if err != nil {
		return nil, err
	}

	return db, nil
}
