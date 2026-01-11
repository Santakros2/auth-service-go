package handlers

import (
	"auth-service/internal/domain"
	"auth-service/internal/service"
	"encoding/json"
	"log"
	"net/http"
)

type AuthHandler struct {
	Service *service.Service
}

func NewHandler(s *service.Service) AuthHandler {
	return AuthHandler{Service: s}
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	AccessToken  string `json:"access"`
	RefreshToken string `json:"refresh"`
}

type token struct {
	RefreshToken string `json:"refresh"`
}

func (h *AuthHandler) LoginHandle(w http.ResponseWriter, r *http.Request) {

	// Storing Input
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validating Input
	if req.Email == "" || req.Password == "" {
		http.Error(w, "email and password required", http.StatusBadRequest)
		return
	}
	log.Println(req.Password)

	// Calling Service layer
	access, refresh, err := h.Service.Login(
		r.Context(), req.Email, req.Password,
	)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Response sending
	json.NewEncoder(w).Encode(AuthResponse{
		AccessToken:  access,
		RefreshToken: refresh,
	})
}

func (h *AuthHandler) RefreshHandler(w http.ResponseWriter, r *http.Request) {

	var refresh token

	// Input taking
	if err := json.NewDecoder(r.Body).Decode(&refresh); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if refresh.RefreshToken == "" {
		http.Error(w, "no refresh", http.StatusBadRequest)
		return
	}

	access, newRefresh, err := h.Service.Refresh(r.Context(), refresh.RefreshToken)

	if err != nil {
		http.Error(w, "Invalid credential", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(AuthResponse{
		AccessToken:  access,
		RefreshToken: newRefresh,
	})
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	var refresh token

	if err := json.NewDecoder(r.Body).Decode(refresh); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if refresh.RefreshToken == "" {
		http.Error(w, "no refresh", http.StatusBadRequest)
		return
	}

	err := h.Service.Logout(r.Context(), refresh.RefreshToken)

	if err != nil {
		http.Error(w, "Invalid Credential", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(http.StatusNoContent)
}

func (h *AuthHandler) LogoutAll(w http.ResponseWriter, r *http.Request) {

}

func (h *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	var user domain.SignupRequest

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if user.Email == "" || user.Password == "" {
		http.Error(w, "email and password required", http.StatusBadRequest)
		return
	}

	err := h.Service.Signup(r.Context(), user.Email, user.Password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
