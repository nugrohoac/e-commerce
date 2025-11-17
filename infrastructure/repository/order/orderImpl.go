package order

import (
	"context"
	"database/sql"
	"errors"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/nugrohoac/e-commerce/constant"
	"github.com/nugrohoac/e-commerce/entity"
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

func (r repository) FindWarehouseWithStockForUpdate(ctx context.Context, tx *Tx, productID uint64, qty int) (uint64, error) {
	query, args, err := sq.Select("pw.warehouse_id").
		From("product_warehouse pw").
		Join("warehouse w ON w.id = pw.warehouse_id AND w.status = 'active'").
		Where(sq.Eq{"pw.product_id": productID}).
		Where("(pw.quantity - pw.reserved_quantity) >= ?", qty).
		OrderBy("(pw.quantity - pw.reserved_quantity) DESC").
		Limit(1).
		Suffix("FOR UPDATE"). // for locking rows
		ToSql()
	if err != nil {
		return 0, err
	}

	var warehouseID uint64
	err = tx.Tx.QueryRowContext(ctx, query, args...).Scan(&warehouseID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, constant.ErrNoWarehouseHasStock
		}
		return 0, err
	}

	return warehouseID, nil
}

func (r repository) CreateOrder(ctx context.Context, tx *Tx, o *entity.Order) (uint64, error) {
	timeNow := time.Now()

	query, args, err := sq.Insert("`order`").
		Columns("user_id", "status", "total_price", "created_at", "updated_at").
		Values(
			o.UserID,
			o.Status,
			o.TotalPrice,
			timeNow,
			timeNow,
		).
		ToSql()
	if err != nil {
		return 0, err
	}

	res, err := tx.Tx.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(id), nil
}

func (r repository) CreateOrderItem(ctx context.Context, tx *Tx, orderID uint64, productID uint64, warehouseID uint64, qty int, price float64) error {
	timeNow := time.Now()

	query, args, err := sq.Insert("order_item").
		Columns("order_id", "product_id", "warehouse_id", "qty", "price", "created_at", "updated_at").
		Values(
			orderID,
			productID,
			warehouseID,
			qty,
			price,
			timeNow,
			timeNow,
		).
		ToSql()
	if err != nil {
		return err
	}

	if _, err = tx.Tx.ExecContext(ctx, query, args...); err != nil {
		return err
	}

	return nil
}

func (r repository) ReserveStock(ctx context.Context, tx *Tx, productID uint64, warehouseID uint64, qty int, orderID uint64) error {
	increaseExpr := sq.Expr("reserved_quantity + ?", qty)

	query, args, err := sq.Update("product_warehouse").
		Set("reserved_quantity", increaseExpr).
		Where(sq.Eq{"product_id": productID}).
		Where(sq.Eq{"warehouse_id": warehouseID}).
		ToSql()
	if err != nil {
		return err
	}

	if _, err = tx.Tx.ExecContext(ctx, query, args...); err != nil {
		return err
	}

	// Insert reservation record
	timeNow := time.Now()
	query, args, err = sq.Insert("stock_reservation").
		Columns("order_id", "product_id", "warehouse_id", "qty", "status", "expires_at", "created_at").
		Values(
			orderID,
			productID,
			warehouseID,
			qty,
			"active",
			timeNow.Add(10*time.Minute),
			timeNow,
		).
		ToSql()
	if err != nil {
		return err
	}

	if _, err = tx.Tx.ExecContext(ctx, query, args...); err != nil {
		return err
	}

	return nil
}

func (r repository) UpdateTotalPrice(ctx context.Context, tx *Tx, ID uint64, total float64) error {
	query, args, err := sq.Update("`order`").
		SetMap(sq.Eq{"total_price": total, "updated_at": time.Now()}).
		Where(sq.Eq{"id": ID}).
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
	return &repository{
		db: db,
	}
}
