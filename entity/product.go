package entity

import "time"

type Product struct {
	ID        uint64    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	SKU       string    `json:"sku" db:"sku"`
	Price     float64   `json:"price" db:"price"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type ProductStock struct {
	ID         uint64  `json:"id"`
	Name       string  `json:"name"`
	SKU        string  `json:"sku"`
	Price      float64 `json:"price"`
	TotalStock int     `json:"total_stock"`
}
