package repository

import (
	"database/sql"

	"github.com/will-bernardo/help-desk-go-api/entity"
)

type TicketRepository struct {
	db *sql.DB
}

func NewTicketRepository(db *sql.DB) *TicketRepository {
	return &TicketRepository{
		db: db,
	}
}

func (repo *TicketRepository) Select() ([]entity.Ticket, error) {
	var tickets []entity.Ticket

	//consulta
	rows, err := repo.db.Query("select * from tickets")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//iteração no resultado
	for rows.Next() {
		var ticket entity.Ticket
		if err = rows.Scan(
			&ticket.ID,
			&ticket.UserID,
			&ticket.CircuitID,
			&ticket.SuportID,
			&ticket.TechnicianID,
			&ticket.Subject,
			&ticket.Description,
			&ticket.Status,
			&ticket.CreatedAt,
			&ticket.UpdatedAt,
		); err != nil {
			return nil, err
		}
		tickets = append(tickets, ticket)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tickets, nil
}

func (repo *TicketRepository) Insert(
	id string,
	circuit string,
	title string,
	description string,
) error {
	_, err := repo.db.Exec("insert into users (id, name, age) values ('abc', 'ronaldo', 33)")
	if err != nil {
		return err
	}

	return nil
}
