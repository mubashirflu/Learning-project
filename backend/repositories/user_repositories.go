package repositories

import (
	"backend/models"
	"context"
	"database/sql"
)

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}
func (repository *UserRepository) CreateUser(
	ctx context.Context,
	input CreateUserInput,
) (models.User, error) {
	var user models.User

	err := repository.db.QueryRowContext(
		ctx,
		`
        INSERT INTO users (name, email, password)
        VALUES ($1, $2, $3)
        RETURNING id, name, email, password, created_at
        `,
		input.Name,
		input.Email,
		input.Password,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)

	return user, err
}

func (repository *UserRepository) GetUserByEmail(
	ctx context.Context,
	email string,
) (models.User, error) {
	var user models.User

	err := repository.db.QueryRowContext(
		ctx,
		`
        SELECT id, name, email, password, created_at
        FROM users
        WHERE email = $1
        `,
		email,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)

	return user, err
}
