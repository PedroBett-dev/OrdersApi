package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/seu-usuario/minha-api/internal/entities"
)

type clientService interface {
	GetAll() ([]entities.Client, error)
}

type ClientHandler struct {
	service clientService
}

/*--------------------------------------------------------*/
func NewClientHandler(service clientService) *ClientHandler {
	return &ClientHandler{service: service}
}

/*--------------------------------------------------------*/
func (h *ClientHandler) GetAll(c *gin.Context) {
	clients, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao buscar clientes"})
		return
	}

	c.JSON(http.StatusOK, clients)
}
