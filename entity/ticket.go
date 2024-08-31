package entity

type TicketLog struct {
	Description string
	CreatedAt   string
}

type Ticket struct {
	ID           string
	CircuitID    string
	SuportID     string
	TechnicianID string
	Title        string
	Description  string
	Status       string
	Logs         []TicketLog
	UpdatedAt    string
	CreatedAt    string
}

func NewTicket() *Ticket {
	return &Ticket{}
}

//Validar se o circuito é válido
//Validar se circuito pertence ao cliente
