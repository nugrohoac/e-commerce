package entity

import "time"

type StockReservation struct {
	ID          uint64    `json:"id" db:"id"`
	OrderID     *uint64   `json:"order_id" db:"order_id"` // nullable
	ProductID   uint64    `json:"product_id" db:"product_id"`
	WarehouseID uint64    `json:"warehouse_id" db:"warehouse_id"`
	Qty         int       `json:"qty" db:"qty"`
	ExpiresAt   time.Time `json:"expires_at" db:"expires_at"`
	Status      string    `json:"status" db:"status"` // active / released / converted_to_order
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}
