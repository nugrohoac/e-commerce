package warehouse

import (
	"net/http"
	"strconv"

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

func (h *handler) ActivateWarehouse(c echo.Context) error {
	ID, err := h.getWarehouseID(c)
	if err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	wh, err := h.warehouseSvc.Activate(c.Request().Context(), ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "internal service error",
		})
	}

	return c.JSON(http.StatusOK, wh)
}

func (h *handler) DeactivateWarehouse(c echo.Context) error {
	ID, err := h.getWarehouseID(c)
	if err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	wh, err := h.warehouseSvc.Deactivate(c.Request().Context(), ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "internal service error",
		})
	}

	return c.JSON(http.StatusOK, wh)
}

func (h *handler) getWarehouseID(c echo.Context) (uint64, error) {
	ID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return uint64(0), err
	}

	return ID, nil
}

func RegisterHandler(e *echo.Echo, warehouseSvc warehouse.Service) {
	h := handler{warehouseSvc: warehouseSvc}
	group := e.Group("/warehouse")
	group.POST("/transfer", h.TransferStock)
	group.PATCH("/:id/activate", h.ActivateWarehouse)
	group.PATCH("/:id/deactivate", h.DeactivateWarehouse)
}
