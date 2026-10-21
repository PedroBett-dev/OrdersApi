package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ApiConfig apiConfig
	DbConfig  dbConfig
}

type apiConfig struct {
	port string
}

type dbConfig struct {
	host     string
	port     string
	user     string
	password string
	name     string
	sslMode  string
}

/*--------------------------------------------------------*/
func InitConfig() Config {
	_ = godotenv.Load()

	return Config{
		ApiConfig: initApiConfig(),
		DbConfig:  initDbConfig(),
	}
}

/*--------------------------------------------------------*/
func initApiConfig() apiConfig {
	return apiConfig{
		port: getEnvOrDefault("PORT", "8080"),
	}
}

/*--------------------------------------------------------*/
func initDbConfig() dbConfig {
	return dbConfig{
		host:     getEnvOrDefault("DB_HOST", "localhost"),
		port:     getEnvOrDefault("DB_PORT", "5432"),
		user:     os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASSWORD"),
		name:     os.Getenv("DB_NAME"),
		sslMode:  getEnvOrDefault("DB_SSLMODE", "disable"),
	}
}

/*--------------------------------------------------------*/
func getEnvOrDefault(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
