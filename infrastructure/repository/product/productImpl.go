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

func NewRepository(db *sql.DB) Repository {
	return repository{db: db}
}
