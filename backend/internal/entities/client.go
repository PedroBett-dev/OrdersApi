package entities

type Client struct {
	Id    string `json:"id"`
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Cpf   string `json:"cpf" binding:"required"`
}
