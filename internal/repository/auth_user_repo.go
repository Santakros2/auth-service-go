package repository

import (
	"auth-service/internal/domain"
	"context"
	"database/sql"
	"log"
)

type AuthUserRepository interface {
	FindByMail(ctx context.Context, email string) (*domain.AuthUser, error)
	Create(ctx context.Context, user *domain.AuthUser) error
	SaveRefresh(ctx context.Context, refresh *domain.RefreshToken) error
}

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(ctx context.Context, user *domain.AuthUser) error {
	query := `INSERT INTO auth_users(id, email, password_hash, role, is_active, is_locked) VALUES(?, ?, ?, ?, ?, ?)`

	_, err := r.DB.ExecContext(ctx, query, user.ID, user.Email, user.Password, user.Role, user.IsActive, user.IsLocked)

	if err != nil {
		log.Println("error in database ", err)
		return err
	}

	return nil

}

func (r *UserRepository) FindByMail(ctx context.Context, email string) (*domain.AuthUser, error) {
	query := `SELECT id, email, password_hash, role, is_active, is_locked FROM auth_users where email = ?`
	row := r.DB.QueryRowContext(ctx, query, email)

	var user domain.AuthUser

	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Role, &user.IsActive, &user.IsLocked)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil

}

func (r *UserRepository) SaveRefresh(ctx context.Context, refresh *domain.RefreshToken) error {
	query := `INSERT INTO auth_users(user_id, token_hash, expire_at, revoked) VALUES(?, ?, ?, ?)`

	_, err := r.DB.ExecContext(ctx, query, refresh.ID, refresh.TokenHash, refresh.ExpireAt, refresh.Revoked)

	if err != nil {
		log.Println("err in db", err)
		return err
	}

	return nil

}
