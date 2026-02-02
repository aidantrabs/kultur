package handler

import (
    "net/http"

    "github.com/labstack/echo/v4"
)

func (h *Handler) Health(c echo.Context) error {
    ctx := c.Request().Context()

    if err := h.pool.Ping(ctx); err != nil {
        return c.JSON(http.StatusServiceUnavailable, map[string]string{
            "status":   "unhealthy",
            "database": "disconnected",
        })
    }

    return c.JSON(http.StatusOK, map[string]string{
        "status":   "healthy",
        "database": "connected",
    })
}
