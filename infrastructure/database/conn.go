package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func NewDBConnection() (*sql.DB, error) {
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASSWORD")
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_name := os.Getenv("DB_NAME")
	db_ssl_mode := os.Getenv("DB_SSL_MODE")

	format := "user=%s password=%s host=%s port=%s dbname=%s sslmode=%s"
	connStr := fmt.Sprintf(format, db_user, db_pass, db_host, db_port, db_name, db_ssl_mode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
