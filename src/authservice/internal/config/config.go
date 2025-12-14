package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort    string
	DBConn        string
	JWTSecret     string
	JWTAccessTTL  string
	JWTRefreshTTL string
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("WARNING:", err)
	}

	return &Config{
		ServerPort:    getEnv("PORT", "8080"),
		DBConn:        getEnv("DB_CONN", "user=postgres password=secret host=localhost port=5432 dbname=auth"),
		JWTSecret:     getEnv("JWT_SECRET", "my-jwt-secret"),
		JWTAccessTTL:  getEnv("JWT_ACCESS_TTL", "15m"),
		JWTRefreshTTL: getEnv("JWT_REFRESH_TTL", "72h"),
	}
}

func (c *Config) String() string {
	return fmt.Sprintf("Port: %s, DBConn: %s", c.ServerPort, c.DBConn)
}
