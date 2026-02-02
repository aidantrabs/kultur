package handler

import (
    "net/http"

    "github.com/aidantrabs/kultur/backend/internal/email"
    "github.com/labstack/echo/v4"
)

type TestEmailRequest struct {
    Email string `json:"email"`
}

func (h *Handler) TestWelcomeEmail(c echo.Context) error {
    var req TestEmailRequest
    if err := c.Bind(&req); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
    }

    if req.Email == "" {
        return echo.NewHTTPError(http.StatusBadRequest, "email is required")
    }

    if err := h.email.SendWelcome(req.Email, "test-unsubscribe-token"); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }

    return c.JSON(http.StatusOK, map[string]string{
        "message": "welcome email sent",
        "to":      req.Email,
    })
}

func (h *Handler) TestFestivalReminder(c echo.Context) error {
    var req TestEmailRequest
    if err := c.Bind(&req); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
    }

    if req.Email == "" {
        return echo.NewHTTPError(http.StatusBadRequest, "email is required")
    }

    if err := h.email.SendFestivalReminder(req.Email, "Trinidad Carnival", "carnival", "test-unsubscribe-token", 7); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }

    return c.JSON(http.StatusOK, map[string]string{
        "message": "festival reminder email sent",
        "to":      req.Email,
    })
}

func (h *Handler) TestWeeklyDigest(c echo.Context) error {
    var req TestEmailRequest
    if err := c.Bind(&req); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
    }

    if req.Email == "" {
        return echo.NewHTTPError(http.StatusBadRequest, "email is required")
    }

    testFestivals := []email.FestivalDigestItem{
        {
            Name:     "Trinidad Carnival",
            Slug:     "carnival",
            Date:     "February 16-17, 2026",
            Heritage: "Mixed Heritage",
            Region:   "Nationwide",
        },
        {
            Name:     "Hosay",
            Slug:     "hosay",
            Date:     "February 20, 2026",
            Heritage: "Indian Heritage",
            Region:   "St. James",
        },
        {
            Name:     "Phagwa",
            Slug:     "phagwa",
            Date:     "March 14, 2026",
            Heritage: "Indian Heritage",
            Region:   "Central Trinidad",
        },
    }

    if err := h.email.SendWeeklyDigest(req.Email, testFestivals, "test-unsubscribe-token"); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }

    return c.JSON(http.StatusOK, map[string]string{
        "message": "weekly digest email sent",
        "to":      req.Email,
    })
}
