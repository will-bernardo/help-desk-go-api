package usecase

import (
	"log"

	"github.com/will-bernardo/help-desk-go-api/entity"
	"github.com/will-bernardo/help-desk-go-api/usecase/dto"
	"github.com/will-bernardo/help-desk-go-api/usecase/repository"
)

type CreatetTicket struct {
	Repository repository.TicketRepository
}

func NewCreateTicket(repo repository.TicketRepository) *CreatetTicket {
	return &CreatetTicket{
		Repository: repo,
	}
}

//1. Receber dados
//2. Executar regras da entidade para a funcionalidade especifica
//3. Persistir dados
//4. Retornar dados

func (ct *CreatetTicket) Execute(input dto.CreateTicketDtoInput) (dto.CreateTicketDtoOutput, error) {
	ticket := entity.NewTicket()
	ticket.UserID = input.UserID
	ticket.CircuitID = input.Circuit
	ticket.Subject = input.Title
	ticket.Description = input.Description

	//Aplica regras
	invalidTicket := ticket.IsValid()
	if invalidTicket != nil {
		log.Fatal("ticket inválido")
	}

	//Insere no banco
	err := ct.Repository.Insert(ticket.ID, ticket.CircuitID, ticket.Subject, ticket.Description)
	if err != nil {
		log.Fatal(err)
	}

	//Retorna resultado
	output := dto.CreateTicketDtoOutput{
		Protocol: 69696969,
	}

	return output, nil
}
