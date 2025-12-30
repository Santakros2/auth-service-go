package handlers

import (
	"auth-service/internal/service"
	"net/http"
)

type AuthHandler struct {
	Service *service.Service
}

func NewHandler(s *service.Service) AuthHandler {
	return AuthHandler{Service: s}
}

func (h *AuthHandler) LoginHandle(w http.ResponseWriter, r *http.Request) {
	// var user domain.AuthUser
}
