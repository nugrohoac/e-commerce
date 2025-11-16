package auth

import (
	"net/http"

	"github.com/nugrohoac/e-commerce/application/model"
	"github.com/nugrohoac/e-commerce/application/service/auth"

	"github.com/labstack/echo/v4"
)

type authHandler struct {
	authService auth.AuthService
}

func (a authHandler) Login(c echo.Context) error {
	var req model.LoginRequest

	// Parse & validate JSON request
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid_request"})
	}

	if req.Identifier == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "identifier_and_password_required"})
	}

	// Call service layer
	response, err := a.authService.Login(c.Request().Context(), req.Identifier, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid_credentials"})
	}

	return c.JSON(http.StatusOK, response)
}

func RegisterAuthHandler(e *echo.Echo, authService auth.AuthService) {
	handler := authHandler{authService: authService}
	route := e.Group("/auth")

	route.POST("/login", handler.Login)
}
