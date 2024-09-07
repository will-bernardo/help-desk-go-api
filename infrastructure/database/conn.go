package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var (
	db_user     = os.Getenv("DB_USER")
	db_pass     = os.Getenv("DB_PASSWORD")
	db_host     = os.Getenv("DB_HOST")
	db_port     = os.Getenv("DB_PORT")
	db_name     = os.Getenv("DB_NAME")
	db_ssl_mode = os.Getenv("DB_SSL_MODE")
)

var format = "user=%s password=%s host=%s port=%d dbname=%s sslmode=%s"
var connStr = fmt.Sprintf(format, db_user, db_pass, db_host, db_port, db_name, db_ssl_mode)

func NewDBConnection() (*sql.DB, error) {
	fmt.Println(connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
