package product

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/labstack/gommon/log"

	"github.com/nugrohoac/e-commerce/entity"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return repository{db: db}
}

func (r repository) GetByID(ctx context.Context, ID uint64) (*entity.Product, error) {
	query, args, err := sq.Select("id", "name", "sku", "price", "created_at", "updated_at").
		From("product").
		Where(sq.Eq{"id": ID}).
		Limit(1).
		ToSql()
	if err != nil {
		return nil, err
	}

	p := &entity.Product{}
	err = r.db.QueryRowContext(ctx, query, args...).Scan(
		&p.ID,
		&p.Name,
		&p.SKU,
		&p.Price,
		&p.CreatedAt,
		&p.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return p, nil
}

func (r repository) GetProducts(ctx context.Context) ([]entity.ProductStock, error) {
	query := sq.Select(
		"p.id",
		"p.name",
		"p.sku",
		"p.price",
		"COALESCE(SUM(pw.quantity - pw.reserved_quantity), 0) AS total_stock",
	).From("product p").
		LeftJoin("product_warehouse pw ON pw.product_id = p.id").
		LeftJoin("warehouse w ON w.id = pw.warehouse_id AND w.status = 'active'").
		GroupBy("p.id", "p.name", "p.sku", "p.price")

	sqlStr, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, sqlStr, args...)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = rows.Close(); err != nil {
			log.Error(err)
		}
	}()

	var items []entity.ProductStock
	for rows.Next() {
		var ps entity.ProductStock
		if err = rows.Scan(
			&ps.ID,
			&ps.Name,
			&ps.SKU,
			&ps.Price,
			&ps.TotalStock,
		); err != nil {
			return nil, err
		}
		items = append(items, ps)
	}

	return items, nil
}
