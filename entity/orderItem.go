package entity

import "time"

type OrderItem struct {
	ID          uint64    `json:"id" db:"id"`
	OrderID     uint64    `json:"order_id" db:"order_id"`
	ProductID   uint64    `json:"product_id" db:"product_id"`
	WarehouseID uint64    `json:"warehouse_id" db:"warehouse_id"`
	Qty         int       `json:"qty" db:"qty"`
	Price       float64   `json:"price" db:"price"` // copy from product.price
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
