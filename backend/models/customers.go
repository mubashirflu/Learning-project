package models

import "time"

type Customers struct {
	ID        uint      `json:"id"`
	USER_ID   uint      `json:"user_id"`
	NAME      string    `json:"name"`
	EMAIL     string    `json:"email"`
	PHONE     string    `json:"phone"`
	CREATE_AT time.Time `json:"created_at"`
}
