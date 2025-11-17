package model

import "time"

type WarehouseTransferRequest struct {
	ProductID       uint64 `json:"product_id"`
	FromWarehouseID uint64 `json:"from_warehouse_id"`
	ToWarehouseID   uint64 `json:"to_warehouse_id"`
	Qty             int    `json:"qty"`
}

type WarehouseTransferResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Warehouse struct {
	ID        uint64    `json:"id" db:"id"`
	ShopID    uint64    `json:"shop_id" db:"shop_id"`
	Name      string    `json:"name" db:"name"`
	Status    string    `json:"status" db:"status"` // active / inactive
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
