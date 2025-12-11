package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func NewConfig() *Config {
	envFiles := []string{".env", "server/.env"}
	if err := godotenv.Load(envFiles...); err != nil {
		log.Printf("loading env files %v: %v", envFiles, err)
	}

	return &Config{
		ServerHost: getEnv("HOST", "localhost"),
		ServerPort: getEnv("PORT", "8080"),
		DBConn:     getEnv("DBConn", "postgres://postgres:password@localhost:5432/learn_hub_auth?sslmode=disable"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
