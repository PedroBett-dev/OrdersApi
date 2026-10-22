package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/seu-usuario/minha-api/internal/apperror"
	"github.com/seu-usuario/minha-api/internal/entities"
)

type clientService interface {
	GetAll() ([]entities.Client, error)
	Create(client entities.Client) (entities.Client, error)
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
		RespondError(c, err)
		return
	}

	c.JSON(http.StatusOK, clients)
}

/*--------------------------------------------------------*/
func (h *ClientHandler) Create(c *gin.Context) {
	var client entities.Client
	if err := c.ShouldBindJSON(&client); err != nil {
		RespondError(c, apperror.New(http.StatusBadRequest, err.Error()))
		return
	}

	created, err := h.service.Create(client)
	if err != nil {
		RespondError(c, err)
		return
	}

	c.JSON(http.StatusCreated, created)
}
