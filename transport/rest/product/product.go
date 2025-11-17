package product

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nugrohoac/e-commerce/application/service/product"
)

type handler struct {
	productService product.Service
}

func (h *handler) ListProducts(c echo.Context) error {
	ctx := c.Request().Context()

	products, err := h.productService.ListProducts(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": products,
	})
}

func RegisterHandler(e *echo.Echo, authService product.Service) {
	h := handler{productService: authService}
	group := e.Group("/product")

	group.GET("", h.ListProducts)
}
