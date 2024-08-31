package entity

type Client struct {
	User
	Circuit string
	Plan    string
	CPF     string
	CNPJ    string
	//In a real case this entity must include address info
}

func NewClient() *Client{
	return &Client{}
}