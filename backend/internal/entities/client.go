package entities

type Client struct {
	Id          string  `json:"id"`
	Name        string  `json:"name" binding:"required,max=50"`
	Email       string  `json:"email" binding:"required,max=255"`
	Cpf         string  `json:"cpf" binding:"required"`
	PhoneNumber string  `json:"phoneNumber" binding:"required"`
	Address     Address `json:"address" gorm:"embedded" binding:"required"`
}

type Address struct {
	Cep          string `json:"cep" binding:"required"`
	Street       string `json:"street" binding:"required,max=100"`
	Number       string `json:"number" binding:"required,max=10"`
	Complement   string `json:"complement" binding:"max=50"`
	Neighborhood string `json:"neighborhood" binding:"required,max=50"`
	City         string `json:"city" binding:"required,max=50"`
	State        string `json:"state" binding:"required,len=2"`
}
