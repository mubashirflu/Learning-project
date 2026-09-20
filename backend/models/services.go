package models

import "time"

type Services struct {
	ID               uint      `json:"id"`
	USER_ID          uint      `json:"user_id"`
	NAME             string    `json:"name"`
	DESCRIPTION      string    `json:"description"`
	PRICE            float64   `json:"price"`
	DURATION_MINUTES int       `json:"duration_minutes"`
	CREATE_AT        time.Time `json:"created_at"`
}
