package entity_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/will-bernardo/help-desk-go-api/entity"
)

func TestValidateCircuit(t *testing.T) {
	ticket := entity.NewTicket("123456")
	result, _ := ticket.ValidateCircuit()
	expected := "123456"
	assert.Equal(t, expected, result)
}

func TestCreateTicket(t *testing.T) {
	result := entity.NewTicket("123456")
	expected := "123456"
	assert.Equal(t, expected, result.CircuitID)
}
