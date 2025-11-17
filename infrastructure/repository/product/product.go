package product

import (
	"context"

	"github.com/nugrohoac/e-commerce/entity"
)

type Repository interface {
	GetProducts(ctx context.Context) ([]entity.ProductStock, error)
	GetByID(ctx context.Context, ID uint64) (*entity.Product, error)
}
