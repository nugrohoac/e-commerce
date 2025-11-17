package order

import (
	"context"

	"github.com/nugrohoac/e-commerce/application/model"
)

type Service interface {
	Checkout(ctx context.Context, req model.CheckoutRequest) (*model.CheckoutResponse, error)
	Pay(ctx context.Context, orderID uint64) (string, error)
}
