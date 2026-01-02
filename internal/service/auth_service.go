package service

import (
	"auth-service/internal/repository"
	"auth-service/internal/security"
	"context"
	"fmt"
)

type Service struct {
	Repo repository.AuthUserRepository
}

func NewService(repo repository.AuthUserRepository) *Service {
	return &Service{Repo: repo}
}

func (s *Service) Login(ctx context.Context, email string, password string) (string, string, error) {
	if email == "" {
		return "", "", fmt.Errorf("please enter valid email")
	}

	if password == "" {
		return "", "", fmt.Errorf("please enter valid password")
	}

	user, err := s.Repo.FindByMail(ctx, email)

	if err != nil {
		return "", "", err
	}

	if user == nil {
		return "", "", fmt.Errorf("invalid email or password")
	}

	if !security.PasswordCheck(password, user.Password) {
		return "", "", fmt.Errorf("invalid email or password")
	}

	if !user.IsActive || user.IsLocked {
		return "", "", fmt.Errorf("account disabled")
	}

	return "", "", nil
}

func (s *Service) Refresh(ctx context.Context, refresh string) (string, string, error) {
	return "", "", nil
}

func (s *Service) Logout(ctx context.Context, refresh string) error {
	return nil
}
