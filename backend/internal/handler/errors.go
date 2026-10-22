package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/seu-usuario/minha-api/internal/apperror"
)

/*--------------------------------------------------------*/
func RespondError(c *gin.Context, err error) {
	var appErr *apperror.AppError
	if errors.As(err, &appErr) {
		c.JSON(appErr.Status, gin.H{"error": appErr.Message})
		return
	}

	log.Printf("erro interno: %v", err)
	c.JSON(http.StatusInternalServerError, gin.H{"error": "erro interno"})
}
