package repository

import "github.com/will-bernardo/help-desk-go-api/entity"

type TicketRepository interface {
	Select() entity.Ticket
	Insert(id string, circuit string, title string, description string) error
}