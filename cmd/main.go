package main

import (
	"fmt"
	"log"

	db "github.com/will-bernardo/help-desk-go-api/infrastructure/database"
)

// 	ID           string
// 	CircuitID    string
// 	SuportID     string
// 	TechnicianID string
// 	Title        string
// 	Description  string
// 	Status       string
// 	Logs         []TicketLog
// 	UpdatedAt    string
// 	CreatedAt    string
// }

func main() {
	query := `
		CREATE TABLE IF NOT EXISTS tickets (
		id TEXT NOT NULL PRIMARY KEY,
		circuit_id TEXT
		suport_id TEXT
		technitian_id TEXT
		title TEXT
		description TEXT
		status TEXT
		updated_at TEXT
		created_at TEXT
		)
	`
	sqlite, err := db.NewSQLiteConnection()
	if err != nil {
		log.Fatal(err)
	}

	result, err2 := sqlite.Exec(query)
	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println(result)
}