package model

type CheckoutItem struct {
	ProductID uint64 `json:"product_id" validate:"required"`
	Qty       int    `json:"qty" validate:"required,gt=0"`
}

type CheckoutRequest struct {
	UserID uint64         `json:"user_id" validate:"required"`
	Items  []CheckoutItem `json:"items" validate:"required,dive"`
}

type CheckoutResponse struct {
	OrderID uint64 `json:"order_id"`
	Status  string `json:"status"`
}
