package repositories

import (
	"context"
	"database/sql"

	"backend/models"
)

type ServiceInput struct {
	UserID          uint    `json:"user_id"`
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Price           float64 `json:"price"`
	DurationMinutes int     `json:"duration_minutes"`
}

type UpdateServiceInput struct {
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Price           float64 `json:"price"`
	DurationMinutes int     `json:"duration_minutes"`
}

type ServiceRepository struct {
	db *sql.DB
}

func (repository *ServiceRepository) GetServicesByUser(ctx context.Context, userID uint) (any, any) {
	panic("unimplemented")
}

func NewServiceRepository(db *sql.DB) *ServiceRepository {
	return &ServiceRepository{db: db}
}

func (repository *ServiceRepository) CreateService(
	ctx context.Context,
	input ServiceInput,
) (models.Services, error) {
	var service models.Services

	err := repository.db.QueryRowContext(
		ctx,
		`
		INSERT INTO services (user_id, name, description, price, duration_minutes)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, user_id, name, description, price, duration_minutes, created_at
		`,
		input.UserID,
		input.Name,
		input.Description,
		input.Price,
		input.DurationMinutes,
	).Scan(
		&service.ID,
		&service.USER_ID,
		&service.NAME,
		&service.DESCRIPTION,
		&service.PRICE,
		&service.DURATION_MINUTES,
		&service.CREATE_AT,
	)

	return service, err
}
func (repository *ServiceRepository) GetServiceByID(ctx context.Context, serviceID uint) (models.Services, error) {
	var service models.Services
	err := repository.db.QueryRowContext(ctx, `
		SELECT id, user_id, name, description, price, duration_minutes, created_at
		FROM services
		WHERE id = $1`, serviceID).Scan(
		&service.ID,
		&service.USER_ID,
		&service.NAME,
		&service.DESCRIPTION,
		&service.PRICE,
		&service.DURATION_MINUTES,
		&service.CREATE_AT,
	)
	return service, err
}

func (repository *ServiceRepository) GetServicesByUserID(
	ctx context.Context,
	userID uint,
) ([]models.Services, error) {
	rows, err := repository.db.QueryContext(
		ctx,
		`
		SELECT id, user_id, name, description, price, duration_minutes, created_at
		FROM services
		WHERE user_id = $1
		ORDER BY id
		`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	services := make([]models.Services, 0)
	for rows.Next() {
		var service models.Services

		err := rows.Scan(
			&service.ID,
			&service.USER_ID,
			&service.NAME,
			&service.DESCRIPTION,
			&service.PRICE,
			&service.DURATION_MINUTES,
			&service.CREATE_AT,
		)
		if err != nil {
			return nil, err
		}

		services = append(services, service)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return services, nil
}

func (repository *ServiceRepository) UpdateService(
	ctx context.Context,
	serviceID uint,
	userID uint,
	input UpdateServiceInput,
) (models.Services, error) {
	var service models.Services

	err := repository.db.QueryRowContext(
		ctx,
		`
		UPDATE services
		SET name = $1,
			description = $2,
			price = $3,
			duration_minutes = $4
		WHERE id = $5 AND user_id = $6
		RETURNING id, user_id, name, description, price, duration_minutes, created_at
		`,
		input.Name,
		input.Description,
		input.Price,
		input.DurationMinutes,
		serviceID,
		userID,
	).Scan(
		&service.ID,
		&service.USER_ID,
		&service.NAME,
		&service.DESCRIPTION,
		&service.PRICE,
		&service.DURATION_MINUTES,
		&service.CREATE_AT,
	)

	return service, err
}

func (repository *ServiceRepository) DeleteService(
	ctx context.Context,
	serviceID uint,
	userID uint,
) error {
	result, err := repository.db.ExecContext(
		ctx,
		`
        DELETE FROM services
        WHERE id = $1 AND user_id = $2
        `,
		serviceID,
		userID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
