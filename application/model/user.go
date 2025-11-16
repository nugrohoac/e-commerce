package model

import "time"

type User struct {
	ID           uint64    `json:"-"`
	Name         string    `json:"name"`
	Email        *string   `json:"email"`
	Phone        *string   `json:"phone"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}
