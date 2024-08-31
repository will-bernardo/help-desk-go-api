package usecase

import "github.com/will-bernardo/help-desk-go-api/entity"

//Criar um ticket
//Orquestração

type CreatetTicketUseCase struct {
	Repository entity.TicketRepository
}

func NewCreateTicket(repository entity.TicketRepository) *CreatetTicketUseCase {
	return &CreatetTicketUseCase{
		Repository: repository,
	}
}
