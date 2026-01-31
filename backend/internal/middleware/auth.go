package middleware

import (
    "net/http"

    "github.com/labstack/echo/v4"
)

func APIKeyAuth(apiKey string) echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            if apiKey == "" {
                return echo.NewHTTPError(http.StatusInternalServerError, "admin API key not configured")
            }

            key := c.Request().Header.Get("X-API-Key")
            if key == "" {
                return echo.NewHTTPError(http.StatusUnauthorized, "missing API key")
            }

            if key != apiKey {
                return echo.NewHTTPError(http.StatusUnauthorized, "invalid API key")
            }

            return next(c)
        }
    }
}
