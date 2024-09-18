package repository

import "github.com/will-bernardo/help-desk-go-api/entity"

//Preferir não receber a entidade como parâmetro, receber propriedades
//especificas dela ao invés disso
//ex: title String, description String...
type TicketRepository interface {
	Insert(ticket entity.Ticket) error
}