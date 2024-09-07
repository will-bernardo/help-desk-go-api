package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	db "github.com/will-bernardo/help-desk-go-api/infrastructure/database"
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

	err1 := db.Ping()
	if err1 != nil {
		log.Fatal(err1)
	}

	fmt.Println(os.Getenv("Conected to database"))
}
