package usecase

import (
	"log"
	"time"

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
//2. Executar regras da entidade
//3. Persistir dados
//4. Retornar dados

func (ct *CreatetTicket) Execute(input dto.CreateTicketDtoInput) (dto.CreateTicketDtoOutput, error) {
	ticket := entity.NewTicket()
	ticket.ID = input.ID
	ticket.CircuitID = input.Circuit
	ticket.Title = input.Title
	ticket.Description = input.Description
	ticket.UpdatedAt = time.Now().Format(time.RFC3339)
	ticket.CreatedAt = time.Now().Format(time.RFC3339)

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
