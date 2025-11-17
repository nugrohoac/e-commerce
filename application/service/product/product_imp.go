package product

import (
	"context"

	"github.com/nugrohoac/e-commerce/application/model"
	"github.com/nugrohoac/e-commerce/infrastructure/repository/product"
)

type service struct {
	productRepository product.Repository
}

func (s service) ListProducts(ctx context.Context) ([]model.ProductResponse, error) {
	stocks, err := s.productRepository.GetProducts(ctx)
	if err != nil {
		return nil, err
	}

	res := make([]model.ProductResponse, len(stocks))
	for i, stock := range stocks {
		res[i] = model.ProductResponse{
			ID:         stock.ID,
			Name:       stock.Name,
			SKU:        stock.SKU,
			Price:      stock.Price,
			TotalStock: stock.TotalStock,
		}
	}

	return res, nil
}

func NewService(productRepository product.Repository) Service {
	return service{productRepository: productRepository}
}
