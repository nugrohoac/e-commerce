package warehouse

import (
	"github.com/labstack/echo/v4"
	"github.com/nugrohoac/e-commerce/application/model"
	"github.com/nugrohoac/e-commerce/application/service/warehouse"
)

type handler struct {
	warehouseSvc warehouse.Service
}

func (h *handler) TransferStock(c echo.Context) error {
	var req model.WarehouseTransferRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, echo.Map{"error": "invalid request"})
	}

	if req.Qty <= 0 {
		return c.JSON(400, echo.Map{"error": "qty must be > 0"})
	}

	resp, err := h.warehouseSvc.TransferStock(c.Request().Context(), req)
	if err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	return c.JSON(200, resp)
}

func RegisterHandler(e *echo.Echo, warehouseSvc warehouse.Service) {
	h := handler{warehouseSvc: warehouseSvc}
	group := e.Group("/warehouse")
	group.POST("/transfer", h.TransferStock)
}
