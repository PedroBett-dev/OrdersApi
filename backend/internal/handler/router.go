package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/seu-usuario/minha-api/internal/config"
)

func InitRoutes(cfg config.Config) *gin.Engine {

	router := gin.Default()

	return router

}
