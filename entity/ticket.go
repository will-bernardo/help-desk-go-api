package entity

import "errors"

type Ticket struct {
	ID           string
	UserID       string
	CircuitID    string
	SuportID     string
	TechnicianID string
	Subject      string
	Description  string
	Status       string
	UpdatedAt    string
	CreatedAt    string
}

func NewTicket(circuit string) *Ticket {
	return &Ticket{
		CircuitID: circuit,
	}
}

func (t *Ticket) ValidateCircuit() (string, error) {
	if len(t.CircuitID) == 0 {
		return "", errors.New("Please provide a valid circuit identifier")
	}
	return t.CircuitID, nil
}
