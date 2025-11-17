package model

type ProductResponse struct {
	ID         uint64  `json:"id"`
	Name       string  `json:"name"`
	SKU        string  `json:"sku"`
	Price      float64 `json:"price"`
	TotalStock int     `json:"total_stock"`
}
