package auth

import (
	"context"

	"github.com/nugrohoac/e-commerce/application/model"
)

type Service interface {
	Login(ctx context.Context, identifier, password string) (*model.AuthResponse, error)
}
