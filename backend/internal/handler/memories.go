package handler

import (
    "net/http"

    "github.com/aidantrabs/trinbago-hackathon/backend/internal/db"
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
    slug := c.Param("slug")

    festival, err := h.queries.GetFestivalBySlug(ctx, slug)
    if err != nil {
        return echo.NewHTTPError(http.StatusNotFound, "festival not found")
    }

    memories, err := h.queries.ListMemoriesByFestival(ctx, pgtype.UUID{Bytes: festival.ID.Bytes, Valid: true})
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

    festivalUUID, err := uuid.Parse(req.FestivalID)
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "invalid festival_id")
    }

    memory, err := h.queries.CreateMemory(ctx, db.CreateMemoryParams{
        FestivalID:   pgtype.UUID{Bytes: festivalUUID, Valid: true},
        AuthorName:   pgtype.Text{String: req.AuthorName, Valid: req.AuthorName != ""},
        AuthorEmail:  pgtype.Text{String: req.AuthorEmail, Valid: req.AuthorEmail != ""},
        Content:      req.Content,
        YearOfMemory: pgtype.Text{String: req.YearOfMemory, Valid: req.YearOfMemory != ""},
    })
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "failed to create memory")
    }

    return c.JSON(http.StatusCreated, memory)
}
