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
}
