package db

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func NewDBConnection() (*sql.DB, error) {
	connStr := os.Getenv("DATABASE_URI")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
