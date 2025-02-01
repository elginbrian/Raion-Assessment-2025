package repository

import (
	"context"
	"fmt"
	contract "raion-assessment/domain/contract"
	entity "raion-assessment/domain/entity"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type authRepository struct {
	db *pgxpool.Pool
}

func NewAuthRepository(db *pgxpool.Pool) contract.IAuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	query := "SELECT id, name, email, bio, image_url, password_hash FROM users WHERE email = $1"

	row := r.db.QueryRow(ctx, query, email)

	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Bio, &user.ImageURL, &user.PasswordHash); err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error scanning row: %w", err)
	}
    
	return &user, nil
}