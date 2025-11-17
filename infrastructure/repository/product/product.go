package product

import (
	"context"

	"github.com/nugrohoac/e-commerce/entity"
)

type Repository interface {
	GetProducts(ctx context.Context) ([]entity.ProductStock, error)
}
