package warehouse

import (
	"context"

	"github.com/nugrohoac/e-commerce/application/model"
	"github.com/nugrohoac/e-commerce/constant"
	"github.com/nugrohoac/e-commerce/constant/enum"
	"github.com/nugrohoac/e-commerce/infrastructure/repository/warehouse"
)

type service struct {
	warehouseRepo warehouse.Repository
}

const messageStockTransferred = "stock transferred"

func (s service) TransferStock(
	ctx context.Context,
	req model.WarehouseTransferRequest) (*model.WarehouseTransferResponse, error) {
	tx, err := s.warehouseRepo.BeginTx(ctx)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = s.warehouseRepo.RollbackTx(tx)
	}()

	if err = s.warehouseRepo.LockStockForUpdate(ctx, tx, req.ProductID, req.FromWarehouseID); err != nil {
		return nil, err
	}
	if err = s.warehouseRepo.LockStockForUpdate(ctx, tx, req.ProductID, req.ToWarehouseID); err != nil {
		return nil, err
	}

	if err = s.warehouseRepo.ReduceStock(ctx, tx, req.ProductID, req.FromWarehouseID, req.Qty); err != nil {
		return nil, constant.ErrNoWarehouseHasStock
	}

	if err = s.warehouseRepo.IncreaseStock(ctx, tx, req.ProductID, req.ToWarehouseID, req.Qty); err != nil {
		return nil, err
	}

	if err = s.warehouseRepo.InsertTransferLog(ctx, tx, req.ProductID, req.FromWarehouseID, req.ToWarehouseID, req.Qty); err != nil {
		return nil, err
	}

	if err = s.warehouseRepo.CommitTx(tx); err != nil {
		return nil, err
	}

	return &model.WarehouseTransferResponse{
		Status:  constant.StatusSuccess,
		Message: messageStockTransferred,
	}, nil
}

func (s service) Activate(ctx context.Context, ID uint64) (*model.Warehouse, error) {
	err := s.warehouseRepo.UpdateWarehouseStatus(ctx, ID, enum.WarehouseStatusActive)
	if err != nil {
		return nil, err
	}

	wh, err := s.warehouseRepo.GetByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	return &model.Warehouse{
		ShopID: wh.ShopID,
		Name:   wh.Name,
		Status: wh.Status,
	}, nil
}

func (s service) Deactivate(ctx context.Context, ID uint64) (*model.Warehouse, error) {
	err := s.warehouseRepo.UpdateWarehouseStatus(ctx, ID, enum.WarehouseStatusInActive)
	if err != nil {
		return nil, err
	}

	wh, err := s.warehouseRepo.GetByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	return &model.Warehouse{
		ShopID: wh.ShopID,
		Name:   wh.Name,
		Status: wh.Status,
	}, nil
}

func NewService(warehouseRepo warehouse.Repository) Service {
	return service{warehouseRepo: warehouseRepo}
}
