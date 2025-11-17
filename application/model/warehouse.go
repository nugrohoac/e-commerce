package model

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
