package services

import (
	"backend/models"
	"backend/repositories"
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type RegisterUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

// Register
func (service *UserService) Register(
	ctx context.Context,
	input RegisterUserInput,
) (models.User, error) {

	// Basic validation
	if input.Name == "" {
		return models.User{}, fmt.Errorf("name is required")
	}

	if input.Email == "" {
		return models.User{}, fmt.Errorf("email is required")
	}

	if input.Password == "" {
		return models.User{}, fmt.Errorf("password is required")
	}

	if len(input.Password) < 6 {
		return models.User{}, fmt.Errorf("password must be at least 6 characters")
	}

	// Check if email already exists
	existingUser, err := service.userRepo.GetUserByEmail(ctx, input.Email)

	if err == nil && existingUser.ID != 0 {
		return models.User{}, fmt.Errorf("email already registered")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(input.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return models.User{}, fmt.Errorf("hash password: %w", err)
	}

	// Create user
	user, err := service.userRepo.CreateUser(
		ctx,
		repositories.CreateUserInput{
			Name:     input.Name,
			Email:    input.Email,
			Password: string(hashedPassword),
		},
	)
	if err != nil {
		return models.User{}, fmt.Errorf("create user: %w", err)
	}

	// Never return password hash to caller
	user.Password = ""

	return user, nil
}

// Login
func (service *UserService) Login(
	ctx context.Context,
	input LoginUserInput,
) (models.User, error) {

	if input.Email == "" {
		return models.User{}, fmt.Errorf("email is required")
	}

	if input.Password == "" {
		return models.User{}, fmt.Errorf("password is required")
	}

	// Get user by email
	user, err := service.userRepo.GetUserByEmail(ctx, input.Email)
	if err != nil {
		return models.User{}, fmt.Errorf("invalid email or password")
	}

	// Compare password with stored hash
	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(input.Password),
	)
	if err != nil {
		return models.User{}, fmt.Errorf("invalid email or password")
	}

	// Don't expose password hash
	user.Password = ""

	return user, nil
}

// Get user by email
func (service *UserService) GetUserByEmail(
	ctx context.Context,
	email string,
) (models.User, error) {

	if email == "" {
		return models.User{}, fmt.Errorf("email is required")
	}

	user, err := service.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return models.User{}, fmt.Errorf("get user: %w", err)
	}

	user.Password = ""

	return user, nil
}
