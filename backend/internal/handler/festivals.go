package handler

import (
    "errors"
    "net/http"
    "strconv"
    "time"

    "github.com/aidantrabs/kultur/backend/internal/db"
    "github.com/aidantrabs/kultur/backend/internal/service"
    "github.com/google/uuid"
    "github.com/jackc/pgx/v5/pgtype"
    "github.com/labstack/echo/v4"
)

func (h *Handler) ListFestivals(c echo.Context) error {
    ctx := c.Request().Context()

    festivals, err := h.festivals.List(ctx, service.ListFestivalsParams{
        Region:   c.QueryParam("region"),
        Heritage: c.QueryParam("heritage"),
    })
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "failed to fetch festivals")
    }

    return c.JSON(http.StatusOK, festivals)
}

func (h *Handler) ListUpcomingFestivals(c echo.Context) error {
    ctx := c.Request().Context()

    festivals, err := h.festivals.ListUpcoming(ctx)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "failed to fetch upcoming festivals")
    }

    return c.JSON(http.StatusOK, festivals)
}

func (h *Handler) ListFestivalsByYear(c echo.Context) error {
    ctx := c.Request().Context()

    yearStr := c.QueryParam("year")
    if yearStr == "" {
        yearStr = strconv.Itoa(time.Now().Year())
    }

    year, err := strconv.Atoi(yearStr)
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "invalid year")
    }

    festivals, err := h.festivals.ListByYear(ctx, int32(year))
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "failed to fetch festivals")
    }

    return c.JSON(http.StatusOK, festivals)
}

func (h *Handler) GetFestival(c echo.Context) error {
    ctx := c.Request().Context()

    festival, err := h.festivals.GetBySlug(ctx, c.Param("slug"))
    if errors.Is(err, service.ErrFestivalNotFound) {
        return echo.NewHTTPError(http.StatusNotFound, "festival not found")
    }
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "failed to fetch festival")
    }

    return c.JSON(http.StatusOK, festival)
}

func (h *Handler) GetFestivalDates(c echo.Context) error {
    ctx := c.Request().Context()

    festival, err := h.festivals.GetBySlug(ctx, c.Param("slug"))
    if errors.Is(err, service.ErrFestivalNotFound) {
        return echo.NewHTTPError(http.StatusNotFound, "festival not found")
    }
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "failed to fetch festival")
    }

    dates, err := h.festivals.GetDates(ctx, festival.ID)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "failed to fetch festival dates")
    }

    return c.JSON(http.StatusOK, dates)
}

type CreateFestivalRequest struct {
    Slug             string `json:"slug"`
    Name             string `json:"name"`
    DateType         string `json:"date_type"`
    Region           string `json:"region"`
    HeritageType     string `json:"heritage_type"`
    FestivalType     string `json:"festival_type"`
    Summary          string `json:"summary"`
    Story            string `json:"story"`
    WhatToExpect     string `json:"what_to_expect"`
    HowToParticipate string `json:"how_to_participate"`
    PracticalInfo    string `json:"practical_info"`
    CoverImageUrl    string `json:"cover_image_url"`
    IsPublished      bool   `json:"is_published"`
}

func (h *Handler) CreateFestival(c echo.Context) error {
    ctx := c.Request().Context()

    var req CreateFestivalRequest
    if err := c.Bind(&req); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
    }

    if req.Slug == "" || req.Name == "" {
        return echo.NewHTTPError(http.StatusBadRequest, "slug and name are required")
    }

    festival, err := h.festivals.Create(ctx, db.CreateFestivalParams{
        Slug:             req.Slug,
        Name:             req.Name,
        DateType:         req.DateType,
        Region:           req.Region,
        HeritageType:     req.HeritageType,
        FestivalType:     req.FestivalType,
        Summary:          req.Summary,
        Story:            pgtype.Text{String: req.Story, Valid: req.Story != ""},
        WhatToExpect:     pgtype.Text{String: req.WhatToExpect, Valid: req.WhatToExpect != ""},
        HowToParticipate: pgtype.Text{String: req.HowToParticipate, Valid: req.HowToParticipate != ""},
        PracticalInfo:    pgtype.Text{String: req.PracticalInfo, Valid: req.PracticalInfo != ""},
        CoverImageUrl:    pgtype.Text{String: req.CoverImageUrl, Valid: req.CoverImageUrl != ""},
        GalleryImages:    []byte("[]"),
        VideoEmbeds:      []byte("[]"),
        IsPublished:      pgtype.Bool{Bool: req.IsPublished, Valid: true},
    })
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "failed to create festival")
    }

    return c.JSON(http.StatusCreated, festival)
}

func (h *Handler) UpdateFestival(c echo.Context) error {
    ctx := c.Request().Context()

    id, err := uuid.Parse(c.Param("id"))
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "invalid festival id")
    }

    var req CreateFestivalRequest
    if err := c.Bind(&req); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
    }

    festival, err := h.festivals.Update(ctx, db.UpdateFestivalParams{
        ID:               pgtype.UUID{Bytes: id, Valid: true},
        Slug:             req.Slug,
        Name:             req.Name,
        DateType:         req.DateType,
        Region:           req.Region,
        HeritageType:     req.HeritageType,
        FestivalType:     req.FestivalType,
        Summary:          req.Summary,
        Story:            pgtype.Text{String: req.Story, Valid: req.Story != ""},
        WhatToExpect:     pgtype.Text{String: req.WhatToExpect, Valid: req.WhatToExpect != ""},
        HowToParticipate: pgtype.Text{String: req.HowToParticipate, Valid: req.HowToParticipate != ""},
        PracticalInfo:    pgtype.Text{String: req.PracticalInfo, Valid: req.PracticalInfo != ""},
        CoverImageUrl:    pgtype.Text{String: req.CoverImageUrl, Valid: req.CoverImageUrl != ""},
        GalleryImages:    []byte("[]"),
        VideoEmbeds:      []byte("[]"),
        IsPublished:      pgtype.Bool{Bool: req.IsPublished, Valid: true},
    })
    if errors.Is(err, service.ErrFestivalNotFound) {
        return echo.NewHTTPError(http.StatusNotFound, "festival not found")
    }
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "failed to update festival")
    }

    return c.JSON(http.StatusOK, festival)
}

func (h *Handler) DeleteFestival(c echo.Context) error {
    ctx := c.Request().Context()

    id, err := uuid.Parse(c.Param("id"))
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "invalid festival id")
    }

    if err := h.festivals.Delete(ctx, pgtype.UUID{Bytes: id, Valid: true}); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "failed to delete festival")
    }

    return c.NoContent(http.StatusNoContent)
}

type CreateFestivalDateRequest struct {
    FestivalID  string `json:"festival_id"`
    Year        int    `json:"year"`
    StartDate   string `json:"start_date"`
    EndDate     string `json:"end_date"`
    IsTentative bool   `json:"is_tentative"`
}

func (h *Handler) CreateFestivalDate(c echo.Context) error {
    ctx := c.Request().Context()

    var req CreateFestivalDateRequest
    if err := c.Bind(&req); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
    }

    festivalID, err := uuid.Parse(req.FestivalID)
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "invalid festival_id")
    }

    startDate, err := time.Parse("2006-01-02", req.StartDate)
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "invalid start_date format (use YYYY-MM-DD)")
    }

    var endDate pgtype.Date
    if req.EndDate != "" {
        ed, err := time.Parse("2006-01-02", req.EndDate)
        if err != nil {
            return echo.NewHTTPError(http.StatusBadRequest, "invalid end_date format (use YYYY-MM-DD)")
        }

        endDate = pgtype.Date{Time: ed, Valid: true}
    }

    date, err := h.festivals.CreateDate(ctx, db.CreateFestivalDateParams{
        FestivalID:  pgtype.UUID{Bytes: festivalID, Valid: true},
        Year:        int32(req.Year),
        StartDate:   pgtype.Date{Time: startDate, Valid: true},
        EndDate:     endDate,
        IsTentative: pgtype.Bool{Bool: req.IsTentative, Valid: true},
    })
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "failed to create festival date")
    }

    return c.JSON(http.StatusCreated, date)
}

type UpdateFestivalDateRequest struct {
    StartDate   string `json:"start_date"`
    EndDate     string `json:"end_date"`
    IsTentative bool   `json:"is_tentative"`
}

func (h *Handler) UpdateFestivalDate(c echo.Context) error {
    ctx := c.Request().Context()

    id, err := uuid.Parse(c.Param("id"))
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "invalid date id")
    }

    var req UpdateFestivalDateRequest
    if err := c.Bind(&req); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
    }

    startDate, err := time.Parse("2006-01-02", req.StartDate)
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "invalid start_date format (use YYYY-MM-DD)")
    }

    var endDate pgtype.Date
    if req.EndDate != "" {
        ed, err := time.Parse("2006-01-02", req.EndDate)
        if err != nil {
            return echo.NewHTTPError(http.StatusBadRequest, "invalid end_date format (use YYYY-MM-DD)")
        }

        endDate = pgtype.Date{Time: ed, Valid: true}
    }

    date, err := h.festivals.UpdateDate(ctx, db.UpdateFestivalDateParams{
        ID:          pgtype.UUID{Bytes: id, Valid: true},
        StartDate:   pgtype.Date{Time: startDate, Valid: true},
        EndDate:     endDate,
        IsTentative: pgtype.Bool{Bool: req.IsTentative, Valid: true},
    })
    if errors.Is(err, service.ErrFestivalDateNotFound) {
        return echo.NewHTTPError(http.StatusNotFound, "festival date not found")
    }
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "failed to update festival date")
    }

    return c.JSON(http.StatusOK, date)
}

func (h *Handler) DeleteFestivalDate(c echo.Context) error {
    ctx := c.Request().Context()

    id, err := uuid.Parse(c.Param("id"))
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "invalid date id")
    }

    if err := h.festivals.DeleteDate(ctx, pgtype.UUID{Bytes: id, Valid: true}); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "failed to delete festival date")
    }

    return c.NoContent(http.StatusNoContent)
}
