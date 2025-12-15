package server

import (
	"log"
	"net/http"

	"authservice/internal/config"
	"authservice/internal/service"
)

type HTTPServer struct {
	config      *config.Config
	authService service.AuthService
	server      *http.Server
	mux         *http.ServeMux
}

func NewHTTPServer(config *config.Config, authService service.AuthService) *HTTPServer {
	httpServer := HTTPServer{
		config:      config,
		authService: authService,
		mux: http.NewServeMux(),
	}
	
	httpServer.registerHandlers()
	server := http.Server{
		Addr: ":" + config.ServerPort,
		Handler: httpServer.mux,
	}

	httpServer.server = &server
	return &httpServer
}

func (s *HTTPServer) registerHandlers() {
	s.mux.HandleFunc("/healthcheck", s.handleHealthCheck)
	s.mux.HandleFunc("/api/auth/register", s.handleRegister)
	s.mux.HandleFunc("/api/auth/login", s.handleLogin)
	s.mux.HandleFunc("/api/auth/refresh", s.handleRefresh)
	s.mux.HandleFunc("/api/auth/logout", s.handleLogout)
	s.mux.HandleFunc("/api/user", s.handleUser)
	s.mux.HandleFunc("/api/auth/swagger", s.handleSwagger)
}

func (s *HTTPServer) RunServer() {
	log.Fatal(s.server.ListenAndServe())
}
