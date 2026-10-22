package handler

import (
	"github.com/seu-usuario/minha-api/internal/repository"
	"github.com/seu-usuario/minha-api/internal/service"
	"gorm.io/gorm"
)

type Handlers struct {
	Customer *CustomerHandler
	Product  *ProductHandler
}

func InitHandlers(db *gorm.DB) *Handlers {

	return &Handlers{
		Customer: initCustomerHandler(db),
		Product:  initProductHandler(db),
	}
}

func initCustomerHandler(db *gorm.DB) *CustomerHandler {
	customerRepo := repository.NewCustomerRepository(db)
	customerService := service.NewCustomerService(customerRepo)
	return NewCustomerHandler(customerService)
}

func initProductHandler(db *gorm.DB) *ProductHandler {
	productRepo := repository.NewProductRepository(db)
	ProductService := service.NewProductService(productRepo)
	return NewProductHandler(ProductService)
}
