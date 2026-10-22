package rauter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/seu-usuario/minha-api/internal/config"
	"github.com/seu-usuario/minha-api/internal/handler"
)

/*--------------------------------------------------------*/
func InitRoutes(cfg config.Config, handlers *handler.Handlers) *gin.Engine {

	r := gin.Default()

	heathRoutes(r)
	clientsRoutes(r, handlers.Client)

	return r
}

/*--------------------------------------------------------*/
func heathRoutes(r *gin.Engine) {

	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "I'm fine!")
	})
}

/*--------------------------------------------------------*/
func clientsRoutes(r *gin.Engine, h *handler.ClientHandler) {
	r.GET("/clients", h.GetAll)
	r.POST("/clients", h.Create)
}
