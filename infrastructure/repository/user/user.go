package user

import (
	"context"

	"github.com/nugrohoac/e-commerce/entity"
)

//go:generate mockery --name Repository --output ../../../mocks/infrastructure/repository/user
type Repository interface {
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	GetByPhone(ctx context.Context, phone string) (*entity.User, error)
}
