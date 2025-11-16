package entity

import "time"

type Order struct {
	ID         uint64    `json:"id" db:"id"`
	UserID     uint64    `json:"user_id" db:"user_id"`
	Status     string    `json:"status" db:"status"` // pending / paid / expired / canceled
	TotalPrice float64   `json:"total_price" db:"total_price"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}
