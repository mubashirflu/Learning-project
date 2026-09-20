package models

import "time"

type Appointments struct {
	ID               uint      `json:"id"`
	CUSTOMER_ID      uint      `json:"customer_id"`
	SERVICE_ID       uint      `json:"service_id"`
	APPOINTMENT_DATE time.Time `json:"appointment_date"`
	START_TIME       time.Time `json:"start_time"`
	END_TIME         time.Time `json:"end_time"`
	STATUS           string    `json:"status"`
	CREATED_AT       time.Time `json:"created_at"`
}
