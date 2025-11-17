package warehouse

import (
	"context"

	"github.com/nugrohoac/e-commerce/application/model"
)

type Service interface {
	TransferStock(ctx context.Context, req model.WarehouseTransferRequest) (*model.WarehouseTransferResponse, error)
}
