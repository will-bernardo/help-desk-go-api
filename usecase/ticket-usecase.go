package usecase

import (
	"log"

	"github.com/will-bernardo/help-desk-go-api/entity"
	"github.com/will-bernardo/help-desk-go-api/usecase/dto"
	"github.com/will-bernardo/help-desk-go-api/usecase/repository"
)

//Criar um ticket
//Orquestração

type CreatetTicket struct {
	Repository repository.TicketRepository
}

func NewCreateTicket(repository repository.TicketRepository) *CreatetTicket {
	return &CreatetTicket{
		Repository: repository,
	}
}

func (ct *CreatetTicket) Execute(input dto.CreateTicketDtoInput) (dto.CreateTicketDtoOutput, error) {
	ticket := entity.NewTicket()
	ticket.ID = input.ID
	ticket.CircuitID = input.Circuit
	ticket.Title = input.Title
	ticket.Description = input.Description

	//Aplica regras
	invalidTicket := ticket.IsValid()

	if invalidTicket != nil {
		log.Fatal("ticket inválido")
	}

	//Insere no banco
	ct.Repository.Insert(*ticket)

	output := dto.CreateTicketDtoOutput{
		Protocol: 69696969,
	}

	return output, nil
}
