package services

import (
	"backend/models"
	"backend/repositories"
	"context"
	"fmt"
	"time"
)

type CreateAppointmentInput struct {
	CustomerID      uint      `json:"customer_id"`
	ServiceID       uint      `json:"service_id"`
	AppointmentDate time.Time `json:"appointment_date"`
	StartTime       time.Time `json:"start_time"`
}

type AppointmentService struct {
	appointments *repositories.Appointments
	customers    *repositories.CustomerRepository
	services     *repositories.ServiceRepository
}

func NewAppointmentService(
	appointments *repositories.Appointments,
	customers *repositories.CustomerRepository,
	services *repositories.ServiceRepository,
) *AppointmentService {
	return &AppointmentService{
		appointments: appointments,
		customers:    customers,
		services:     services,
	}
}

func (service *AppointmentService) CreateAppointment(
	ctx context.Context,
	input CreateAppointmentInput,
) (models.Appointments, error) {
	if input.CustomerID == 0 {
		return models.Appointments{}, fmt.Errorf("customer id is required")
	}
	if input.ServiceID == 0 {
		return models.Appointments{}, fmt.Errorf("service id is required")
	}
	if input.AppointmentDate.IsZero() {
		return models.Appointments{}, fmt.Errorf("appointment date is required")
	}
	if input.StartTime.IsZero() {
		return models.Appointments{}, fmt.Errorf("start time is required")
	}

	if _, err := service.customers.GetCustomerByID(ctx, input.CustomerID); err != nil {
		return models.Appointments{}, fmt.Errorf("verify customer: %w", err)
	}

	serviceModel, err := service.services.GetServiceByID(ctx, input.ServiceID)
	if err != nil {
		return models.Appointments{}, fmt.Errorf("verify service: %w", err)
	}
	if serviceModel.DURATION_MINUTES <= 0 {
		return models.Appointments{}, fmt.Errorf("service duration must be greater than zero")
	}

	endTime := input.StartTime.Add(time.Duration(serviceModel.DURATION_MINUTES) * time.Minute)
	if endTime.Day() != input.StartTime.Day() {
		return models.Appointments{}, fmt.Errorf("appointment cannot cross midnight")
	}

	available, err := service.appointments.CheckAvailability(
		ctx,
		input.ServiceID,
		input.AppointmentDate,
		input.StartTime,
		endTime,
	)
	if err != nil {
		return models.Appointments{}, fmt.Errorf("check appointment availability: %w", err)
	}
	if !available {
		return models.Appointments{}, fmt.Errorf("appointment time is not available")
	}

	return service.appointments.CreateAppointment(ctx, repositories.AppointmentInput{
		CustomerID:      input.CustomerID,
		ServiceID:       input.ServiceID,
		AppointmentDate: input.AppointmentDate,
		StartTime:       input.StartTime,
		EndTime:         endTime,
		Status:          "booked",
	})
}

// Get appointments of a customer
func (service *AppointmentService) GetAppointmentsByCustomer(
	ctx context.Context,
	customerID uint,
) ([]models.Appointments, error) {

	if customerID == 0 {
		return nil, fmt.Errorf("customer id is required")
	}

	// Verify customer exists
	if _, err := service.customers.GetCustomerByID(
		ctx,
		customerID,
	); err != nil {
		return nil, fmt.Errorf("verify customer: %w", err)
	}

	appointments, err := service.appointments.GetAppointmentsByCustomer(
		ctx,
		customerID,
	)

	if err != nil {
		return nil, fmt.Errorf(
			"get customer appointments: %w",
			err,
		)
	}

	return appointments, nil
}
func (service *AppointmentService) CancelAppointment(
	ctx context.Context,
	appointmentID uint,
) error {

	if appointmentID == 0 {
		return fmt.Errorf("appointment id is required")
	}

	err := service.appointments.CancelAppointment(
		ctx,
		appointmentID,
	)

	if err != nil {
		return fmt.Errorf(
			"cancel appointment: %w",
			err,
		)
	}

	return nil
}
func (service *AppointmentService) GetAppointmentsByService(
	ctx context.Context,
	serviceID uint,
	appointmentDate time.Time,
) ([]models.Appointments, error) {

	if serviceID == 0 {
		return nil, fmt.Errorf("service id is required")
	}

	if appointmentDate.IsZero() {
		return nil, fmt.Errorf("appointment date is required")
	}

	// Verify service exists
	if _, err := service.services.GetServiceByID(
		ctx,
		serviceID,
	); err != nil {
		return nil, fmt.Errorf(
			"verify service: %w",
			err,
		)
	}

	appointments, err := service.appointments.GetAppointmentsByDate(
		ctx,
		serviceID,
		appointmentDate,
	)

	if err != nil {
		return nil, fmt.Errorf(
			"get service appointments: %w",
			err,
		)
	}

	return appointments, nil
}
