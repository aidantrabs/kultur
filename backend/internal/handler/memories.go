package handler

import (
    "errors"
    "net/http"

    "github.com/aidantrabs/kultur/backend/internal/service"
    "github.com/google/uuid"
    "github.com/jackc/pgx/v5/pgtype"
    "github.com/labstack/echo/v4"
)

type CreateMemoryRequest struct {
    FestivalID   string `json:"festival_id"`
    AuthorName   string `json:"author_name"`
    AuthorEmail  string `json:"author_email"`
    Content      string `json:"content"`
    YearOfMemory string `json:"year_of_memory"`
}

func (h *Handler) ListMemoriesByFestival(c echo.Context) error {
    ctx := c.Request().Context()

    memories, err := h.memories.ListByFestivalSlug(ctx, c.Param("slug"))
    if errors.Is(err, service.ErrFestivalNotFound) {
        return echo.NewHTTPError(http.StatusNotFound, "festival not found")
    }
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "failed to fetch memories")
    }

    return c.JSON(http.StatusOK, memories)
}

func (h *Handler) CreateMemory(c echo.Context) error {
    ctx := c.Request().Context()

    var req CreateMemoryRequest
    if err := c.Bind(&req); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
    }

    if req.Content == "" {
        return echo.NewHTTPError(http.StatusBadRequest, "content is required")
    }

    if req.FestivalID == "" {
        return echo.NewHTTPError(http.StatusBadRequest, "festival_id is required")
    }

    festivalUUID, err := uuid.Parse(req.FestivalID)
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "invalid festival_id")
    }

    memory, err := h.memories.Create(ctx, service.CreateMemoryParams{
        FestivalID:   pgtype.UUID{Bytes: festivalUUID, Valid: true},
        AuthorName:   req.AuthorName,
        AuthorEmail:  req.AuthorEmail,
        Content:      req.Content,
        YearOfMemory: req.YearOfMemory,
    })
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "failed to create memory")
    }

    return c.JSON(http.StatusCreated, memory)
}

func (h *Handler) ListAllMemories(c echo.Context) error {
    ctx := c.Request().Context()

    memories, err := h.memories.ListAll(ctx)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "failed to fetch memories")
    }

    return c.JSON(http.StatusOK, memories)
}

type UpdateMemoryStatusRequest struct {
    Status string `json:"status"`
}

func (h *Handler) UpdateMemoryStatus(c echo.Context) error {
    ctx := c.Request().Context()

    id, err := uuid.Parse(c.Param("id"))
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "invalid memory id")
    }

    var req UpdateMemoryStatusRequest
    if err := c.Bind(&req); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
    }

    if req.Status == "" {
        return echo.NewHTTPError(http.StatusBadRequest, "status is required")
    }

    if req.Status != "approved" && req.Status != "rejected" && req.Status != "pending" {
        return echo.NewHTTPError(http.StatusBadRequest, "status must be approved, rejected, or pending")
    }

    err = h.memories.UpdateStatus(ctx, pgtype.UUID{Bytes: id, Valid: true}, req.Status)
    if errors.Is(err, service.ErrMemoryNotFound) {
        return echo.NewHTTPError(http.StatusNotFound, "memory not found")
    }

		if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "failed to update memory status")
    }

    return c.JSON(http.StatusOK, map[string]string{"status": "updated"})
}

func (h *Handler) DeleteMemory(c echo.Context) error {
    ctx := c.Request().Context()

    id, err := uuid.Parse(c.Param("id"))
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "invalid memory id")
    }

    if err := h.memories.Delete(ctx, pgtype.UUID{Bytes: id, Valid: true}); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "failed to delete memory")
    }

    return c.NoContent(http.StatusNoContent)
}
