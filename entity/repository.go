package entity

type TicketRepository interface {
	Insert(ticket Ticket) error
}
