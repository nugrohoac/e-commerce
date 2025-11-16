package auth

import (
	"context"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/nugrohoac/e-commerce/application/model"
	"github.com/nugrohoac/e-commerce/constant"
	"github.com/nugrohoac/e-commerce/entity"
	"github.com/nugrohoac/e-commerce/infrastructure/repository/user"
)

type AuthService interface {
	Login(ctx context.Context, identifier, password string) (*model.AuthResponse, error)
}

type authService struct {
	userRepo user.UserRepository
}

func (a authService) Login(ctx context.Context, identifier, password string) (*model.AuthResponse, error) {
	var (
		userEntity *entity.User
		err        error
	)

	if strings.Contains(identifier, "@") {
		userEntity, err = a.userRepo.GetByEmail(ctx, identifier)
		if err != nil {
			return nil, err
		}
	} else {
		userEntity, err = a.userRepo.GetByPhone(ctx, identifier)
		if err != nil {
			return nil, err
		}
	}

	if bcrypt.CompareHashAndPassword([]byte(userEntity.PasswordHash), []byte(password)) != nil {
		return nil, constant.ErrInvalidLogin
	}

	return &model.AuthResponse{
		Token: "", // TODO: generate token
		User: model.User{
			ID:    userEntity.ID,
			Name:  userEntity.Name,
			Email: userEntity.Email,
			Phone: userEntity.Phone,
		},
	}, nil
}

func NewAuthService(userRepo user.UserRepository) AuthService {
	return authService{userRepo: userRepo}
}
