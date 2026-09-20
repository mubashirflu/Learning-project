package services

import (
	"backend/models"
	"backend/repositories"
	"context"
	"fmt"
	"strings"
)

type CreateCustomerInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type UpdateCustomerInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type CustomerService struct {
	customerRepo *repositories.CustomerRepository
	userRepo     *repositories.UserRepository
}

func NewCustomerService(
	customerRepo *repositories.CustomerRepository,
	userRepo *repositories.UserRepository,
) *CustomerService {
	return &CustomerService{
		customerRepo: customerRepo,
		userRepo:     userRepo,
	}
}

// Create Customer
func (service *CustomerService) CreateCustomer(
	ctx context.Context,
	userID uint,
	input CreateCustomerInput,
) (models.Customers, error) {

	if userID == 0 {
		return models.Customers{}, fmt.Errorf("user id is required")
	}

	input.Name = strings.TrimSpace(input.Name)
	input.Email = strings.TrimSpace(input.Email)
	input.Phone = strings.TrimSpace(input.Phone)

	if input.Name == "" {
		return models.Customers{}, fmt.Errorf("customer name is required")
	}

	if input.Email == "" {
		return models.Customers{}, fmt.Errorf("customer email is required")
	}

	if input.Phone == "" {
		return models.Customers{}, fmt.Errorf("customer phone is required")
	}

	customer, err := service.customerRepo.CreateCustomer(
		ctx,
		repositories.CustomerRepositoryInputs{
			UserID: userID,
			Name:   input.Name,
			Email:  input.Email,
			Phone:  input.Phone,
		},
	)
	if err != nil {
		return models.Customers{}, fmt.Errorf("create customer: %w", err)
	}

	return customer, nil
}

// Get Customers of logged-in user
func (service *CustomerService) GetCustomers(
	ctx context.Context,
	userID uint,
) ([]models.Customers, error) {

	if userID == 0 {
		return nil, fmt.Errorf("user id is required")
	}

	customers, err := service.customerRepo.GetCustomersByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("get customers: %w", err)
	}

	return customers, nil
}

// Update Customer
func (service *CustomerService) UpdateCustomer(
	ctx context.Context,
	userID uint,
	customerID uint,
	input UpdateCustomerInput,
) (models.Customers, error) {

	if userID == 0 {
		return models.Customers{}, fmt.Errorf("user id is required")
	}

	if customerID == 0 {
		return models.Customers{}, fmt.Errorf("customer id is required")
	}

	input.Name = strings.TrimSpace(input.Name)
	input.Email = strings.TrimSpace(input.Email)
	input.Phone = strings.TrimSpace(input.Phone)

	if input.Name == "" {
		return models.Customers{}, fmt.Errorf("customer name is required")
	}

	if input.Email == "" {
		return models.Customers{}, fmt.Errorf("customer email is required")
	}

	if input.Phone == "" {
		return models.Customers{}, fmt.Errorf("customer phone is required")
	}

	customer, err := service.customerRepo.GetCustomerByID(ctx, customerID)
	if err != nil {
		return models.Customers{}, fmt.Errorf("get customer: %w", err)
	}

	// Ownership check
	if customer.USER_ID != userID {
		return models.Customers{}, fmt.Errorf("you do not have permission to update this customer")
	}

	customer.NAME = input.Name
	customer.EMAIL = input.Email
	customer.PHONE = input.Phone
	if err := service.customerRepo.UpdateCustomer(ctx, customer); err != nil {
		return models.Customers{}, fmt.Errorf("update customer: %w", err)
	}

	return customer, nil
}

// Delete Customer
func (service *CustomerService) DeleteCustomer(
	ctx context.Context,
	userID uint,
	customerID uint,
) error {

	if userID == 0 {
		return fmt.Errorf("user id is required")
	}

	if customerID == 0 {
		return fmt.Errorf("customer id is required")
	}

	customer, err := service.customerRepo.GetCustomerByID(ctx, customerID)
	if err != nil {
		return fmt.Errorf("get customer: %w", err)
	}

	// Ownership check
	if customer.USER_ID != userID {
		return fmt.Errorf("you do not have permission to delete this customer")
	}

	if err := service.customerRepo.DeleteCustomer(ctx, customerID); err != nil {
		return fmt.Errorf("delete customer: %w", err)
	}

	return nil
}
