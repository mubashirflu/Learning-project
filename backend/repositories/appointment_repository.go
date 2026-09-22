package repositories

import (
	"backend/models"
	"context"
	"database/sql"
	"time"
)

type AppointmentInput struct {
	CustomerID      uint      `json:"customer_id"`
	ServiceID       uint      `json:"service_id"`
	AppointmentDate time.Time `json:"appointment_date"`
	StartTime       time.Time `json:"start_time"`
	EndTime         time.Time `json:"end_time"`
	Status          string    `json:"status"`
}
type Appointments struct {
	db *sql.DB
}

func NewAppointments(db *sql.DB) *Appointments {
	return &Appointments{db: db}
}
func (repositories *Appointments) CreateAppointment(ctx context.Context, input AppointmentInput) (models.Appointments, error) {
	var appointments models.Appointments
	err := repositories.db.QueryRowContext(
		ctx,
		`INSERT INTO appointments(customer_id,service_id,appointment_date,start_time,end_time,status)
			VALUES ($1,$2,$3,$4,$5,$6)
			RETURNING id,customer_id,service_id,appointment_date,start_time,end_time,status,created_at
			`,
		input.CustomerID,
		input.ServiceID,
		input.AppointmentDate,
		input.StartTime,
		input.EndTime,
		input.Status).Scan(
		&appointments.ID,
		&appointments.CUSTOMER_ID,
		&appointments.SERVICE_ID,
		&appointments.APPOINTMENT_DATE,
		&appointments.START_TIME,
		&appointments.END_TIME,
		&appointments.STATUS,
		&appointments.CREATED_AT,
	)
	return appointments, err
}

func (repositories *Appointments) GetAppointmentsByCustomer(
	ctx context.Context,
	customerID uint,
) ([]models.Appointments, error) {
	return repositories.getAppointments(ctx, `
			SELECT id, customer_id, service_id, appointment_date, start_time, end_time, status, created_at
			FROM appointments
			WHERE customer_id = $1
			ORDER BY appointment_date, start_time, id
		`, customerID)
}

func (repositories *Appointments) GetAppointmentsByDate(
	ctx context.Context,
	serviceID uint,
	appointmentDate time.Time,
) ([]models.Appointments, error) {
	return repositories.getAppointments(ctx, `
			SELECT id, customer_id, service_id, appointment_date, start_time, end_time, status, created_at
			FROM appointments
			WHERE service_id = $1 AND appointment_date = $2
			ORDER BY start_time, id
		`, serviceID, appointmentDate)
}

func (repositories *Appointments) CheckAvailability(
	ctx context.Context,
	serviceID uint,
	appointmentDate time.Time,
	startTime time.Time,
	endTime time.Time,
) (bool, error) {
	var hasOverlap bool
	err := repositories.db.QueryRowContext(
		ctx,
		`SELECT EXISTS (
				SELECT 1
				FROM appointments
				WHERE service_id = $1
				AND appointment_date = $2
				AND status <> 'cancelled'
				AND start_time < $4
				AND end_time > $3
			)`,
		serviceID,
		appointmentDate,
		startTime,
		endTime,
	).Scan(&hasOverlap)

	return !hasOverlap, err
}

func (repositories *Appointments) CancelAppointment(
	ctx context.Context,
	appointmentID uint,
) error {
	result, err := repositories.db.ExecContext(
		ctx,
		`UPDATE appointments
			SET status = 'cancelled'
			WHERE id = $1`,
		appointmentID,
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

func (repositories *Appointments) getAppointments(
	ctx context.Context,
	query string,
	args ...any,
) ([]models.Appointments, error) {
	rows, err := repositories.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	appointments := make([]models.Appointments, 0)
	for rows.Next() {
		var appointment models.Appointments
		err := rows.Scan(
			&appointment.ID,
			&appointment.CUSTOMER_ID,
			&appointment.SERVICE_ID,
			&appointment.APPOINTMENT_DATE,
			&appointment.START_TIME,
			&appointment.END_TIME,
			&appointment.STATUS,
			&appointment.CREATED_AT,
		)
		if err != nil {
			return nil, err
		}
		appointments = append(appointments, appointment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return appointments, nil
}
func (repositories *Appointments) GetAppointmentsByUserAndDate(
	ctx context.Context,
	userID uint,
	appointmentDate time.Time,
) ([]models.Appointments, error) {

	rows, err := repositories.db.QueryContext(
		ctx,
		`
			SELECT
				a.id,
				a.customer_id,
				a.service_id,
				a.appointment_date,
				a.start_time,
				a.end_time,
				a.status,
				a.created_at
			FROM appointments a
			JOIN services s
				ON s.id = a.service_id
			WHERE s.user_id = $1
			AND a.appointment_date = $2
			ORDER BY a.start_time, a.id
			`,
		userID,
		appointmentDate,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	appointments := make([]models.Appointments, 0)

	for rows.Next() {
		var appointment models.Appointments

		err := rows.Scan(
			&appointment.ID,
			&appointment.CUSTOMER_ID,
			&appointment.SERVICE_ID,
			&appointment.APPOINTMENT_DATE,
			&appointment.START_TIME,
			&appointment.END_TIME,
			&appointment.STATUS,
			&appointment.CREATED_AT,
		)
		if err != nil {
			return nil, err
		}

		appointments = append(appointments, appointment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return appointments, nil
}
