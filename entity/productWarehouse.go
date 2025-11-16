package entity

import "time"

type ProductWarehouse struct {
	ID          uint64    `json:"id" db:"id"`
	ProductID   uint64    `json:"product_id" db:"product_id"`
	WarehouseID uint64    `json:"warehouse_id" db:"warehouse_id"`
	Quantity    int       `json:"quantity" db:"quantity"`
	ReservedQty int       `json:"reserved_quantity" db:"reserved_quantity"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
