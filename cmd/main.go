package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	db "github.com/will-bernardo/help-desk-go-api/infra/database"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := db.NewDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Conected to database")

	rows, err := db.Query("select * from users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id   string
			name string
			age  int32
		)

		if err = rows.Scan(&id, &name, &age); err != nil {
			log.Fatal(err)
		}

		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("id: %s, name: %s, age: %d\n", id, name, age)
	}
}
