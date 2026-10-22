package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/seu-usuario/minha-api/internal/entities"
)

type ProductService interface {
	GetAll() ([]entities.Product, error)
}

type ProductHandler struct {
	service ProductService
}

func NewProductHandler(s ProductService) *ProductHandler { return &ProductHandler{service: s} }

func (h *ProductHandler) GetAll(c *gin.Context) {
	products, err := h.service.GetAll()
	if err != nil {
		RespondError(c, err)
		return
	}
	c.JSON(http.StatusOK, products)
}
