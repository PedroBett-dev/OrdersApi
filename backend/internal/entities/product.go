package entities

type Product struct {
	Id     string `json:"id"`
	Name   string `json:"name" binding:"required,max=50"`
	Price  bool   `json:"price" binding:"required,min=0"`
	Status string `json:"status" binding:"required,oneof=active inactive"`
}
