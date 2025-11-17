package order

import (
	"context"

	"github.com/nugrohoac/e-commerce/application/model"
	"github.com/nugrohoac/e-commerce/constant"
	"github.com/nugrohoac/e-commerce/entity"
	"github.com/nugrohoac/e-commerce/infrastructure/repository/order"
	"github.com/nugrohoac/e-commerce/infrastructure/repository/product"
)

type service struct {
	orderRepo   order.Repository
	productRepo product.Repository
}

func (s service) Checkout(ctx context.Context, req model.CheckoutRequest) (*model.CheckoutResponse, error) {
	tx, err := s.orderRepo.BeginTx(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = s.orderRepo.RollbackTx(tx)
	}()

	// Create order
	o := &entity.Order{
		UserID: req.UserID,
		Status: "pending",
	}
	orderID, err := s.orderRepo.CreateOrder(ctx, tx, o)
	if err != nil {
		return nil, err
	}

	// Process items
	for _, item := range req.Items {
		// Find warehouse
		warehouseID, err := s.orderRepo.FindWarehouseWithStockForUpdate(ctx, tx, item.ProductID, item.Qty)
		if err != nil {
			return nil, err
		}

		// Get product price
		prod, err := s.productRepo.GetByID(ctx, item.ProductID)
		if err != nil {
			return nil, err
		}

		// Create order item
		if err = s.orderRepo.CreateOrderItem(ctx, tx, orderID, item.ProductID, warehouseID, item.Qty, prod.Price); err != nil {
			return nil, err
		}

		// Reserve stock
		if err = s.orderRepo.ReserveStock(ctx, tx, item.ProductID, warehouseID, item.Qty, orderID); err != nil {
			return nil, err
		}

		o.TotalPrice += float64(item.Qty) * prod.Price
	}

	// update total price of order
	if err = s.orderRepo.UpdateTotalPrice(ctx, tx, orderID, o.TotalPrice); err != nil {
		return nil, err
	}

	// All good → commit
	if err = s.orderRepo.CommitTx(tx); err != nil {
		return nil, err
	}

	return &model.CheckoutResponse{
		OrderID: orderID,
		Status:  "pending",
	}, nil
}

func (s service) Pay(ctx context.Context, orderID uint64) (string, error) {
	tx, err := s.orderRepo.BeginTx(ctx)
	if err != nil {
		return "", err
	}

	defer func() {
		_ = s.orderRepo.RollbackTx(tx)
	}()

	// 1. Get all active reservations
	reservations, err := s.orderRepo.GetReservationsByOrder(ctx, tx, orderID)
	if err != nil {
		return "", err
	}

	if len(reservations) == 0 {
		return "", constant.ErrNoActiveReservation
	}

	// 2. Deduct real stock for each reservation
	for _, r := range reservations {
		if err = s.orderRepo.DeductStock(ctx, tx, r.ProductID, r.WarehouseID, r.Qty); err != nil {
			return "", err
		}
	}

	// 3. Mark reservations converted
	if err = s.orderRepo.MarkReservationConverted(ctx, tx, orderID); err != nil {
		return "", err
	}

	// 4. Update order status
	if err = s.orderRepo.MarkOrderPaid(ctx, tx, orderID); err != nil {
		return "", err
	}

	// 5. Commit
	if err = s.orderRepo.CommitTx(tx); err != nil {
		return "", err
	}

	return "paid", nil
}

func NewService(orderRepo order.Repository, productRepo product.Repository) Service {
	return service{orderRepo: orderRepo, productRepo: productRepo}
}
