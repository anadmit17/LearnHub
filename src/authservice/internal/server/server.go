package server

import (
	"net/http"

	"authservice/internal/config"
	"authservice/internal/service"
)

type HTTPServer struct {
	config      *config.Config
	authService *service.AuthService
	server      *http.Server
	mux         *http.ServeMux
}
