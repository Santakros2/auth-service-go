package service

import (
	"auth-service/internal/domain"
	"auth-service/internal/repository"
	"auth-service/internal/security"
	"auth-service/pkg/encrypt"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	Repo repository.AuthUserRepository
}

func NewService(repo repository.AuthUserRepository) *Service {
	return &Service{Repo: repo}
}

func (s *Service) Login(ctx context.Context, email string, password string) (string, string, error) {

	// Validation for the input
	if email == "" {
		return "", "", fmt.Errorf("please enter valid email")
	}

	if password == "" {
		return "", "", fmt.Errorf("please enter valid password")
	}

	// Getting the User info
	user, err := s.Repo.FindByMail(ctx, email)
	log.Println(user.Email)
	if err != nil {
		return "", "", err
	}

	if user == nil {
		return "", "", fmt.Errorf("invalid email or password")
	}

	// Validation for password
	if !security.PasswordCheck(password, user.Password) {
		return "", "", fmt.Errorf("invalid email or password")
	}

	if !user.IsActive || user.IsLocked {
		return "", "", fmt.Errorf("account disabled")
	}

	// Token Generation
	tokenPair, err := security.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		return "", "", err
	}

	// Hash the Refresh Token
	hashRefresh := encrypt.HashToken(tokenPair.RefreshToken)

	// Create struct to store in db
	ref := domain.RefreshToken{
		ID:        uuid.New().String(),
		UserID:    user.ID,
		TokenHash: hashRefresh,
		ExpireAt:  time.Now().Add(30 * 24 * time.Hour),
		Revoked:   false,
	}
	// Saving Refresh Token
	if err := s.Repo.SaveRefresh(ctx, &ref); err != nil {
		log.Print("could not save refresh")
		return "", "", err
	}

	return tokenPair.AccessToken, tokenPair.RefreshToken, nil
}

func (s *Service) Refresh(ctx context.Context, refresh string) (string, string, error) {
	return "", "", nil
}

func (s *Service) Logout(ctx context.Context, refresh string) error {
	return nil
}
