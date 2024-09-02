package repository

import "github.com/will-bernardo/help-desk-go-api/entity"

type TicketRepository interface {
	Insert(ticket entity.Ticket) error
}
