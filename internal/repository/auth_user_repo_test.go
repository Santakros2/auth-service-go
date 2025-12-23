package repository_test

import (
	"context"
	"testing"

	"auth-service/internal/domain"
	"auth-service/internal/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestCreateRepo(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := repository.NewUserRepository(db)

	user := &domain.AuthUser{
		ID:       "uuid-123",
		Email:    "test@example.com",
		Password: "hash",
		Role:     "USER",
		IsActive: true,
		IsLocked: false,
	}

	mock.ExpectExec(`INSERT INTO auth_users`).
		WithArgs(
			user.ID,
			user.Email,
			user.Password,
			user.Role,
			user.IsActive,
			user.IsLocked,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Create(context.Background(), user)

	require.NoError(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}
