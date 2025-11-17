package order

import (
	"github.com/labstack/echo/v4"
	"github.com/nugrohoac/e-commerce/application/model"
	"github.com/nugrohoac/e-commerce/application/service/order"
)

type handler struct {
	orderService order.Service
}

func (h *handler) Checkout(c echo.Context) error {
	var req model.CheckoutRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(400, echo.Map{"error": "invalid_request"})
	}

	response, err := h.orderService.Checkout(c.Request().Context(), req)
	if err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	return c.JSON(200, response)
}

func RegisterHandler(e *echo.Echo, orderService order.Service) {
	h := &handler{orderService: orderService}
	group := e.Group("/order")
	group.POST("/checkout", h.Checkout)
}
