package warehouse

import (
	"context"

	"github.com/nugrohoac/e-commerce/entity"
)

type Repository interface {
	BeginTx(ctx context.Context) (*Tx, error)
	CommitTx(tx *Tx) error
	RollbackTx(tx *Tx) error

	// transfer

	LockStockForUpdate(ctx context.Context, tx *Tx, productID, warehouseID uint64) error
	ReduceStock(ctx context.Context, tx *Tx, productID, warehouseID uint64, qty int) error
	IncreaseStock(ctx context.Context, tx *Tx, productID, warehouseID uint64, qty int) error
	InsertTransferLog(ctx context.Context, tx *Tx, productID, fromWh, toWh uint64, qty int) error

	UpdateWarehouseStatus(ctx context.Context, ID uint64, status string) error
	GetByID(ctx context.Context, ID uint64) (*entity.Warehouse, error)
}
