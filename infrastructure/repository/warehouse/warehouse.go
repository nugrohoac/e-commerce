package warehouse

import "context"

type Repository interface {
	BeginTx(ctx context.Context) (*Tx, error)
	CommitTx(tx *Tx) error
	RollbackTx(tx *Tx) error

	LockStockForUpdate(ctx context.Context, tx *Tx, productID, warehouseID uint64) error
	ReduceStock(ctx context.Context, tx *Tx, productID, warehouseID uint64, qty int) error
	IncreaseStock(ctx context.Context, tx *Tx, productID, warehouseID uint64, qty int) error
	InsertTransferLog(ctx context.Context, tx *Tx, productID, fromWh, toWh uint64, qty int) error
}
