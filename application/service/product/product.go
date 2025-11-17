package product

import (
	"context"

	"github.com/nugrohoac/e-commerce/application/model"
)

type Service interface {
	ListProducts(ctx context.Context) ([]model.ProductResponse, error)
}
