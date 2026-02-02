package handler

import (
    "errors"
    "net/http"

    "github.com/aidantrabs/kultur/backend/internal/service"
    "github.com/google/uuid"
    "github.com/jackc/pgx/v5/pgtype"
    "github.com/labstack/echo/v4"
)

type SubscribeRequest struct {
    Email        string `json:"email"`
    DigestWeekly bool   `json:"digest_weekly"`
}

func (h *Handler) Subscribe(c echo.Context) error {
    ctx := c.Request().Context()

    var req SubscribeRequest
    if err := c.Bind(&req); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
    }

    if req.Email == "" {
        return echo.NewHTTPError(http.StatusBadRequest, "email is required")
    }

    sub, err := h.subscriptions.Create(ctx, service.CreateSubscriptionParams{
        Email:        req.Email,
        DigestWeekly: req.DigestWeekly,
    })
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "failed to create subscription")
    }

    return c.JSON(http.StatusCreated, map[string]interface{}{
        "message": "subscription created, please check your email to confirm",
        "id":      sub.ID,
    })
}

func (h *Handler) ConfirmSubscription(c echo.Context) error {
    ctx := c.Request().Context()

    err := h.subscriptions.Confirm(ctx, c.Param("token"))
    if errors.Is(err, service.ErrInvalidToken) {
        return echo.NewHTTPError(http.StatusNotFound, "invalid confirmation token")
    }
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "failed to confirm subscription")
    }

    return c.JSON(http.StatusOK, map[string]string{
        "message": "subscription confirmed",
    })
}

func (h *Handler) Unsubscribe(c echo.Context) error {
    ctx := c.Request().Context()

    err := h.subscriptions.Unsubscribe(ctx, c.Param("token"))
    if errors.Is(err, service.ErrInvalidToken) {
        return echo.NewHTTPError(http.StatusNotFound, "invalid unsubscribe token")
    }
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "failed to unsubscribe")
    }

    return c.JSON(http.StatusOK, map[string]string{
        "message": "unsubscribed successfully",
    })
}

func (h *Handler) ListAllSubscriptions(c echo.Context) error {
    ctx := c.Request().Context()

    subs, err := h.subscriptions.ListAll(ctx)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "failed to fetch subscriptions")
    }

    return c.JSON(http.StatusOK, subs)
}

func (h *Handler) DeleteSubscription(c echo.Context) error {
    ctx := c.Request().Context()

    id, err := uuid.Parse(c.Param("id"))
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "invalid subscription id")
    }

    if err := h.subscriptions.Delete(ctx, pgtype.UUID{Bytes: id, Valid: true}); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "failed to delete subscription")
    }

    return c.NoContent(http.StatusNoContent)
}
