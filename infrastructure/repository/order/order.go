package order

import (
	"context"

	"github.com/nugrohoac/e-commerce/entity"
)

type Repository interface {
	BeginTx(ctx context.Context) (*Tx, error)
	CommitTx(tx *Tx) error
	RollbackTx(tx *Tx) error

	FindWarehouseWithStockForUpdate(ctx context.Context, tx *Tx, productID uint64, qty int) (uint64, error)
	CreateOrder(ctx context.Context, tx *Tx, o *entity.Order) (uint64, error)
	CreateOrderItem(ctx context.Context, tx *Tx, orderID uint64, productID uint64, warehouseID uint64, qty int, price float64) error
	ReserveStock(ctx context.Context, tx *Tx, productID uint64, warehouseID uint64, qty int, orderID uint64) error
	UpdateTotalPrice(ctx context.Context, tx *Tx, orderID uint64, total float64, ) error
}
