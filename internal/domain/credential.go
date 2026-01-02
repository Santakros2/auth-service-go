package domain

import "time"

type AuthUser struct {
	ID       string
	Email    string
	Password string
	Role     string
	IsActive bool
	IsLocked bool
}

type RefreshToken struct {
	ID        string
	UserID    string
	TokenHash string
	ExpireAt  time.Time
	Revoked   bool
}

type LoginResponse struct {
	Valid bool
}


