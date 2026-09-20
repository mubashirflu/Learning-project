package services

import (
	"backend/models"
	"backend/repositories"
	"context"
	"fmt"
	"strings"
)

type CreateServiceInput struct {
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

type ServiceService struct {
	serviceRepo *repositories.ServiceRepository
	userRepo    *repositories.UserRepository
}

func (service *ServiceService) Create(context context.Context, userID uint, input CreateServiceInput) (any, any) {
	panic("unimplemented")
}

func NewServiceService(
	serviceRepo *repositories.ServiceRepository,
	userRepo *repositories.UserRepository,
) *ServiceService {
	return &ServiceService{
		serviceRepo: serviceRepo,
		userRepo:    userRepo,
	}
}

// Create Service
func (service *ServiceService) CreateService(
	ctx context.Context,
	userID uint,
	input CreateServiceInput,
) (models.Services, error) {

	if userID == 0 {
		return models.Services{}, fmt.Errorf("user id is required")
	}

	name := strings.TrimSpace(input.Name)
	description := strings.TrimSpace(input.Description)

	if name == "" {
		return models.Services{}, fmt.Errorf("service name is required")
	}

	if description == "" {
		return models.Services{}, fmt.Errorf("service description is required")
	}

	if input.Price < 0 {
		return models.Services{}, fmt.Errorf("price cannot be negative")
	}

	if input.DurationMinutes <= 0 {
		return models.Services{}, fmt.Errorf(
			"duration must be greater than zero",
		)
	}

	// Check user exists
	_, err := service.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return models.Services{}, fmt.Errorf(
			"verify user: %w",
			err,
		)
	}

	// Create service
	serviceInput := repositories.ServiceInput{
		UserID:          userID,
		Name:            name,
		Description:     description,
		Price:           input.Price,
		DurationMinutes: input.DurationMinutes,
	}

	newService, err := service.serviceRepo.CreateService(
		ctx,
		serviceInput,
	)

	if err != nil {
		return models.Services{}, fmt.Errorf(
			"create service: %w",
			err,
		)
	}

	return newService, nil
}

// Get all services of a user
func (service *ServiceService) GetServices(
	ctx context.Context,
	userID uint,
) ([]models.Services, error) {

	if userID == 0 {
		return nil, fmt.Errorf("user id is required")
	}

	result, err := service.serviceRepo.GetServicesByUser(
		ctx,
		userID,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"get services: %w",
			err,
		)
	}

	// Repository currently returns any
	services, ok := result.([]models.Services)
	if !ok {
		return nil, fmt.Errorf(
			"invalid services data returned from repository",
		)
	}

	return services, nil
}

// Get single service
func (service *ServiceService) GetService(
	ctx context.Context,
	userID uint,
	serviceID uint,
) (models.Services, error) {

	if userID == 0 {
		return models.Services{}, fmt.Errorf("user id is required")
	}

	if serviceID == 0 {
		return models.Services{}, fmt.Errorf("service id is required")
	}

	existingService, err := service.serviceRepo.GetServiceByID(
		ctx,
		serviceID,
	)
	if err != nil {
		return models.Services{}, fmt.Errorf(
			"get service: %w",
			err,
		)
	}

	// Ownership check
	if existingService.USER_ID != userID {
		return models.Services{}, fmt.Errorf(
			"you do not have permission to access this service",
		)
	}

	return existingService, nil
}

// Update Service
func (service *ServiceService) UpdateService(
	ctx context.Context,
	userID uint,
	serviceID uint,
	input UpdateServiceInput,
) (models.Services, error) {

	if userID == 0 {
		return models.Services{}, fmt.Errorf("user id is required")
	}

	if serviceID == 0 {
		return models.Services{}, fmt.Errorf("service id is required")
	}

	name := strings.TrimSpace(input.Name)
	description := strings.TrimSpace(input.Description)

	if name == "" {
		return models.Services{}, fmt.Errorf("service name is required")
	}

	if description == "" {
		return models.Services{}, fmt.Errorf("service description is required")
	}

	if input.Price < 0 {
		return models.Services{}, fmt.Errorf(
			"price cannot be negative",
		)
	}

	if input.DurationMinutes <= 0 {
		return models.Services{}, fmt.Errorf(
			"duration must be greater than zero",
		)
	}

	// Get existing service
	existingService, err := service.serviceRepo.GetServiceByID(
		ctx,
		serviceID,
	)
	if err != nil {
		return models.Services{}, fmt.Errorf(
			"get service: %w",
			err,
		)
	}

	// Ownership check
	if existingService.USER_ID != userID {
		return models.Services{}, fmt.Errorf(
			"you do not have permission to update this service",
		)
	}

	updateInput := repositories.UpdateServiceInput{
		Name:            name,
		Description:     description,
		Price:           input.Price,
		DurationMinutes: input.DurationMinutes,
	}

	updatedService, err := service.serviceRepo.UpdateService(
		ctx,
		userID,
		serviceID,
		updateInput,
	)
	if err != nil {
		return models.Services{}, fmt.Errorf(
			"update service: %w",
			err,
		)
	}

	return updatedService, nil
}

// Delete Service
func (service *ServiceService) DeleteService(
	ctx context.Context,
	userID uint,
	serviceID uint,
) error {

	if userID == 0 {
		return fmt.Errorf("user id is required")
	}

	if serviceID == 0 {
		return fmt.Errorf("service id is required")
	}

	// Get service
	existingService, err := service.serviceRepo.GetServiceByID(
		ctx,
		serviceID,
	)
	if err != nil {
		return fmt.Errorf(
			"get service: %w",
			err,
		)
	}

	// Ownership check
	if existingService.USER_ID != userID {
		return fmt.Errorf(
			"you do not have permission to delete this service",
		)
	}

	if err := service.serviceRepo.DeleteService(
		ctx,
		userID,
		serviceID,
	); err != nil {
		return fmt.Errorf(
			"delete service: %w",
			err,
		)
	}

	return nil
}
