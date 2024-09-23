package entity

type Ticket struct {
	ID           string
	UserID       string
	CircuitID    string
	SuportID     string
	TechnicianID string
	Subject      string
	Description  string
	Status       string
	Logs         []TicketLog
	UpdatedAt    string
	CreatedAt    string
}

func NewTicket() *Ticket {
	return &Ticket{}
}

//Validar se circuito pertence ao cliente
//Validar se o circuito é válido

func (t *Ticket) IsValid() error {
	return nil
}
