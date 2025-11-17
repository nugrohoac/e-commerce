package auth

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"

	authModel "github.com/nugrohoac/e-commerce/application/model"
	"github.com/nugrohoac/e-commerce/entity"
	userRepoMock "github.com/nugrohoac/e-commerce/mocks/infrastructure/repository/user"
	"github.com/nugrohoac/e-commerce/testdata"
)

func TestLogin(t *testing.T) {
	email := "jhon@gmail.com"
	password := "123456"
	cost := bcrypt.DefaultCost
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		require.NoError(t, err)
	}

	user := entity.User{
		ID:           0,
		Name:         "jhon",
		Email:        &email,
		PasswordHash: string(hashedPassword),
	}

	tests := map[string]struct {
		paramIdentifier  string
		paramPassword    string
		getByEmail       testdata.FuncCaller
		getByID          testdata.FuncCaller
		expectedResponse *authModel.AuthResponse
		expectedError    error
	}{
		"success with email and password": {
			paramIdentifier: email,
			paramPassword:   password,
			getByEmail: testdata.FuncCaller{
				IsCalled: true,
				Input:    []interface{}{context.Background(), email},
				Output:   []interface{}{&user, nil},
			},
			getByID: testdata.FuncCaller{},
			expectedResponse: &authModel.AuthResponse{
				Token: "",
				User: authModel.User{
					ID:    user.ID,
					Name:  user.Name,
					Email: user.Email,
					Phone: user.Phone,
				},
			},
			expectedError: nil,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			userRepo := userRepoMock.NewRepository(t)

			if test.getByEmail.IsCalled {
				userRepo.On("GetByEmail", test.getByEmail.Input...).
					Return(test.getByEmail.Output...).
					Once()
			}

			if test.getByID.IsCalled {
				userRepo.On("GetByID", test.getByEmail.Input...).
					Return(test.getByEmail.Output...).
					Once()
			}

			authSvc := NewService(userRepo)
			authResponse, err := authSvc.Login(context.Background(), test.paramIdentifier, test.paramPassword)
			assert.Equal(t, test.expectedError, err)
			assert.Equal(t, test.expectedResponse, authResponse)
		})
	}
}
