package handler

import (
	"github.com/seu-usuario/minha-api/internal/repository"
	"github.com/seu-usuario/minha-api/internal/service"
	"gorm.io/gorm"
)

type Handlers struct {
	Client *ClientHandler
}

/*--------------------------------------------------------*/
func InitHandlers(db *gorm.DB) *Handlers {

	return &Handlers{
		Client: initClientHandler(db),
	}
}

/*--------------------------------------------------------*/
func initClientHandler(db *gorm.DB) *ClientHandler {
	clientRepo := repository.NewClientRepository(db)
	clientService := service.NewClientService(clientRepo)
	return NewClientHandler(clientService)
}
