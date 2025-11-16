package entity

import "time"

type WarehouseTransfer struct {
	ID              uint64    `json:"id" db:"id"`
	FromWarehouseID uint64    `json:"from_warehouse_id" db:"from_warehouse_id"`
	ToWarehouseID   uint64    `json:"to_warehouse_id" db:"to_warehouse_id"`
	ProductID       uint64    `json:"product_id" db:"product_id"`
	Qty             int       `json:"qty" db:"qty"`
	Status          string    `json:"status" db:"status"` // pending / completed
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
}
