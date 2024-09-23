package dto

type CreateTicketDtoInput struct {
	UserID      string
	Circuit     string
	Title       string
	Description string
}

type CreateTicketDtoOutput struct {
	Protocol int32
}
