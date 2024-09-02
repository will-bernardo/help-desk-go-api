package dto

type CreateTicketDtoInput struct {
	ID          string
	Circuit     string
	Title       string
	Description string
}

type CreateTicketDtoOutput struct {
	Protocol int32
}
