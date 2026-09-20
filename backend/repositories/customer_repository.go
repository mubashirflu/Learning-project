package repositories

import (
	"backend/models"
	"context"
	"database/sql"
	"fmt"
)

type CustomerRepositoryInputs struct {
	UserID uint    `json:"user_id"`
	Name   string  `json:"name"`
	Email  string  `json:"email"`
	Phone  float64 `json:"phone"`
}
type CustomerRepository struct {
	db *sql.DB
}

func NewCustomerReposiotyr(db *sql.DB) *CustomerRepository {
	return &CustomerRepository{db: db}
}
func (repositories *CustomerRepository) CreateCustomer(ctx context.Context,
	input CustomerRepositoryInputs,
) (models.Customers, error) {
	var customers models.Customers
	err := repositories.db.QueryRowContext(
		ctx,
		`
        INSERT INTO customers (uesr_id,name, email, phone,crated_at)
        VALUES ($1, $2, $3,$4,)
        RETURNING user_id, name, email, phone, created_at
        `,
		input.UserID,
		input.Name,
		input.Email,
		input.Phone,
	).Scan(
		&customers.USER_ID,
		&customers.NAME,
		&customers.EMAIL,
		&customers.PHONE,
		&customers.CREATE_AT,
	)
	return customers, err
}
func (repository *CustomerRepository) GetServicesByUserID(
	ctx context.Context,
	userID uint,
) ([]models.Customers, error) {
	rows, err := repository.db.QueryContext(
		ctx,
		`
		SELECT id, user_id, name, email, phone, created_at
		FROM customers
		WHERE user_id = $1
		ORDER BY id
		`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	customers := make([]models.Customers, 0)

	for rows.Next() {
		var customer models.Customers

		err := rows.Scan(
			&customer.ID,
			&customer.USER_ID,
			&customer.NAME,
			&customer.EMAIL,
			&customer.PHONE,
			&customer.CREATE_AT,
		)
		if err != nil {
			return nil, err
		}

		customers = append(customers, customer)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return customers, nil
}
func (repository *CustomerRepository) UpdateCustomer(
	ctx context.Context,
	customer models.Customers,
) error {
	_, err := repository.db.ExecContext(
		ctx,
		`
		UPDATE customers
		SET
			name = $1,
			email = $2,
			phone = $3
		WHERE id = $4
		`,
		customer.NAME,
		customer.EMAIL,
		customer.PHONE,
		customer.ID,
	)

	return err
}
func (repository *CustomerRepository) DeleteCustomer(
	ctx context.Context,
	customerID uint,
) error {
	result, err := repository.db.ExecContext(
		ctx,
		`
		DELETE FROM customers
		WHERE id = $1
		`,
		customerID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("customer with id %d not found", customerID)
	}

	return nil
}
