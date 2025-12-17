package server

import (
	"encoding/json"
	"log"
	"net/http"

	"authservice/internal/models"
)

func (s *HTTPServer) handleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()
	
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error decoding request body:", err)
		return
	}

	s.authService.RegisterUser(user)
	w.WriteHeader(http.StatusCreated)
}

func (s *HTTPServer) handleLogin(w http.ResponseWriter, r *http.Request) {
	// POST /api/auth/login Authenticates a user and issues tokens (e.g., JWT)
	w.WriteHeader(http.StatusNotImplemented)
}

func (s *HTTPServer) handleRefresh(w http.ResponseWriter, r *http.Request) {
	// GET /api/auth/refresh Issues new access tokens using a refresh token
	w.WriteHeader(http.StatusNotImplemented)
}

func (s *HTTPServer) handleLogout(w http.ResponseWriter, r *http.Request) {
	// DELETE /api/auth/logout Invalidates a user's session or token
	w.WriteHeader(http.StatusNotImplemented)
}

func (s *HTTPServer) handleUser(w http.ResponseWriter, r *http.Request) {
	// GET /api/user Retrieves logged-in user details (protected route)
	w.WriteHeader(http.StatusNotImplemented)
}

func (s *HTTPServer) handleSwagger(w http.ResponseWriter, r *http.Request) {
	// GET /api/auth/swagger
	w.WriteHeader(http.StatusNotImplemented)
}

func (s *HTTPServer) handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
