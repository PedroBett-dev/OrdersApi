package main

import (
	"log"

	"github.com/seu-usuario/minha-api/internal/config"
	"github.com/seu-usuario/minha-api/internal/handler"
	"github.com/seu-usuario/minha-api/internal/repository"
	rauter "github.com/seu-usuario/minha-api/internal/router"
)

func main() {

	cfg := config.InitConfig()

	log.Printf("api configurada na porta %s", cfg.GetApiPort())

	db, err := repository.InitRepo(cfg)
	if err != nil {
		log.Fatalf("erro ao conectar no postgres: %v", err)
	}

	handlers := handler.InitHandlers(db)

	router := rauter.InitRoutes(cfg, handlers)

	if err := router.Run(":" + cfg.GetApiPort()); err != nil {
		log.Fatalf("erro ao iniciar a api: %v", err)
	}

}
