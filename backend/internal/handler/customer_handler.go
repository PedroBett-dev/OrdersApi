package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/seu-usuario/minha-api/internal/apperror"
	"github.com/seu-usuario/minha-api/internal/entities"
)

type customerService interface {
	GetAll() ([]entities.Customer, error)
	GetById(id string) (entities.Customer, error)
	GetAddress(id string) (entities.Address, error)
	Create(customer entities.Customer) (entities.Customer, error)
}

type CustomerHandler struct {
	service customerService
}

func NewCustomerHandler(s customerService) *CustomerHandler {
	return &CustomerHandler{service: s}
}

func (h *CustomerHandler) GetAll(c *gin.Context) {
	customers, err := h.service.GetAll()
	if err != nil {
		RespondError(c, err)
		return
	}
	c.JSON(http.StatusOK, customers)
}

func (h *CustomerHandler) GetById(c *gin.Context) {
	customer, err := h.service.GetById(strings.ToUpper(c.Param("id")))
	if err != nil {
		RespondError(c, err)
		return
	}
	c.JSON(http.StatusOK, customer)
}

func (h *CustomerHandler) GetAddress(c *gin.Context) {
	address, err := h.service.GetAddress(strings.ToUpper(c.Param("id")))
	if err != nil {
		RespondError(c, err)
		return
	}
	c.JSON(http.StatusOK, address)
}

func (h *CustomerHandler) Create(c *gin.Context) {
	var customer entities.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		RespondError(c, apperror.New(http.StatusBadRequest, err.Error()))
		return
	}
	created, err := h.service.Create(customer)
	if err != nil {
		RespondError(c, err)
		return
	}
	c.JSON(http.StatusCreated, created)
}
