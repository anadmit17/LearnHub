package main

import (
	"authservice/internal/config"
	"authservice/internal/repository"
	"authservice/internal/server"
	"authservice/internal/service"
)

func main() {
	config := config.NewConfig()
	repo := repository.NewInMemoryUserRepository()
	authService := service.NewDefaultAuthService(repo)
	server := server.NewHTTPServer(config, authService)
	server.RunServer()
}