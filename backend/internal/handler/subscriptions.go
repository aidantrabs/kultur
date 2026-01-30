package handler

import (
    "crypto/rand"
    "encoding/hex"
    "net/http"

    "github.com/aidantrabs/trinbago-hackathon/backend/internal/db"
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

    confirmToken, err := generateToken()
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "failed to generate token")
    }

    unsubToken, err := generateToken()
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "failed to generate token")
    }

    sub, err := h.queries.CreateSubscription(ctx, db.CreateSubscriptionParams{
        Email:             req.Email,
        DigestWeekly:      pgtype.Bool{Bool: req.DigestWeekly, Valid: true},
        FestivalReminders: []byte("[]"),
        ConfirmationToken: pgtype.Text{String: confirmToken, Valid: true},
        UnsubscribeToken:  unsubToken,
    })
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "failed to create subscription")
    }

    // TODO: send confirmation email via Resend

    return c.JSON(http.StatusCreated, map[string]interface{}{
        "message": "subscription created, please check your email to confirm",
        "id":      sub.ID,
    })
}

func (h *Handler) ConfirmSubscription(c echo.Context) error {
    ctx := c.Request().Context()
    token := c.Param("token")

    sub, err := h.queries.GetSubscriptionByConfirmationToken(ctx, pgtype.Text{String: token, Valid: true})
    if err != nil {
        return echo.NewHTTPError(http.StatusNotFound, "invalid confirmation token")
    }

    if err := h.queries.ConfirmSubscription(ctx, sub.ID); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "failed to confirm subscription")
    }

    return c.JSON(http.StatusOK, map[string]string{
        "message": "subscription confirmed",
    })
}

func (h *Handler) Unsubscribe(c echo.Context) error {
    ctx := c.Request().Context()
    token := c.Param("token")

    sub, err := h.queries.GetSubscriptionByUnsubscribeToken(ctx, token)
    if err != nil {
        return echo.NewHTTPError(http.StatusNotFound, "invalid unsubscribe token")
    }

    if err := h.queries.DeleteSubscription(ctx, sub.ID); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "failed to unsubscribe")
    }

    return c.JSON(http.StatusOK, map[string]string{
        "message": "unsubscribed successfully",
    })
}

func generateToken() (string, error) {
    bytes := make([]byte, 32)
    if _, err := rand.Read(bytes); err != nil {
        return "", err
    }
    return hex.EncodeToString(bytes), nil
}
