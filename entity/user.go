package entity

import "time"

type User struct {
	ID           uint64    `json:"id" db:"id"`
	Name         string    `json:"name" db:"name"`
	Email        *string   `json:"email" db:"email"`
	Phone        *string   `json:"phone" db:"phone"`
	PasswordHash string    `json:"-" db:"password_hash"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}
