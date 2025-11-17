package warehouse

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
)

type Tx struct {
	Tx *sql.Tx
}
type repository struct {
	db *sql.DB
}

func (r repository) BeginTx(ctx context.Context) (*Tx, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &Tx{Tx: tx}, nil
}

func (r repository) CommitTx(tx *Tx) error {
	return tx.Tx.Commit()
}

func (r repository) RollbackTx(tx *Tx) error {
	return tx.Tx.Rollback()
}

func (r repository) LockStockForUpdate(ctx context.Context, tx *Tx, productID, warehouseID uint64) error {
	query, args, err := sq.
		Select("id").
		From("product_warehouse").
		Where(sq.Eq{"product_id": productID, "warehouse_id": warehouseID}).
		Suffix("FOR UPDATE").
		ToSql()
	if err != nil {
		return err
	}

	if _, err = tx.Tx.ExecContext(ctx, query, args...); err != nil {
		return err
	}

	return nil
}

func (r repository) ReduceStock(ctx context.Context, tx *Tx, productID, warehouseID uint64, qty int) error {
	query, args, err := sq.
		Update("product_warehouse").
		Set("quantity", sq.Expr("quantity - ?", qty)).
		Where(sq.Eq{"product_id": productID, "warehouse_id": warehouseID}).
		Where("quantity >= ?", qty).
		ToSql()
	if err != nil {
		return err
	}

	if _, err = tx.Tx.ExecContext(ctx, query, args...); err != nil {
		return err
	}

	return nil
}

func (r repository) IncreaseStock(ctx context.Context, tx *Tx, productID, warehouseID uint64, qty int) error {
	query, args, err := sq.
		Update("product_warehouse").
		Set("quantity", sq.Expr("quantity + ?", qty)).
		Where(sq.Eq{"product_id": productID, "warehouse_id": warehouseID}).
		ToSql()
	if err != nil {
		return err
	}

	if _, err = tx.Tx.ExecContext(ctx, query, args...); err != nil {
		return err
	}

	return nil
}

func (r repository) InsertTransferLog(ctx context.Context, tx *Tx, productID, fromWh, toWh uint64, qty int) error {
	query, args, err := sq.
		Insert("warehouse_transfer").
		Columns("product_id", "from_warehouse_id", "to_warehouse_id", "qty").
		Values(productID, fromWh, toWh, qty).
		ToSql()
	if err != nil {
		return err
	}

	if _, err = tx.Tx.ExecContext(ctx, query, args...); err != nil {
		return err
	}

	return nil
}

func NewRepository(db *sql.DB) Repository {
	return repository{db: db}
}
