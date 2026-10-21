package main

import (
	"log"

	"github.com/seu-usuario/minha-api/internal/config"
	"github.com/seu-usuario/minha-api/internal/handler"
	"github.com/seu-usuario/minha-api/internal/repository"
)

func main() {

	cfg := config.InitConfig()

	log.Printf("api configurada na porta %s", cfg.GetApiPort())

	db, err := repository.InitRepo(cfg)
	if err != nil {
		log.Fatalf("erro ao conectar no postgres: %v", err)
	}
	_ = db

	router := handler.InitRoutes(cfg)

	router.Run(cfg.GetApiPort())

}
