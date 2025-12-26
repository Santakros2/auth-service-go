package service

import (
	"auth-service/internal/domain"
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

func (s *Service) Login(ctx context.Context, email string, password string) (*domain.LoginResponse, error) {
	if email == "" {
		return nil, fmt.Errorf("please enter valid email")
	}

	if password == "" {
		return nil, fmt.Errorf("please enter valid password")
	}

	user, err := s.Repo.FindByMail(ctx, email)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	if !security.PasswordCheck(password, user.Password) {
		return nil, fmt.Errorf("invalid email or password")
	}

	if !user.IsActive || user.IsLocked {
		return nil, fmt.Errorf("account disabled")
	}

	return &domain.LoginResponse{Valid: true}, nil
}
