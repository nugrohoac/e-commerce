package order

import (
	"strconv"

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

func (h *handler) PayOrder(c echo.Context) error {
	orderIDParam := c.Param("id")

	orderIDUint, err := strconv.ParseUint(orderIDParam, 10, 64)
	if err != nil {
		return c.JSON(400, echo.Map{"error": "invalid order id"})
	}

	status, err := h.orderService.Pay(c.Request().Context(), orderIDUint)
	if err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	return c.JSON(200, echo.Map{
		"order_id": orderIDUint,
		"status":   status,
	})

}

func RegisterHandler(e *echo.Echo, orderService order.Service) {
	h := &handler{orderService: orderService}
	group := e.Group("/order")
	group.POST("/checkout", h.Checkout)
	group.POST("/:id/pay", h.PayOrder)
}
