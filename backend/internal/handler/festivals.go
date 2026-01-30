package handler

import (
    "net/http"

    "github.com/labstack/echo/v4"
)

func (h *Handler) ListFestivals(c echo.Context) error {
    ctx := c.Request().Context()

    region := c.QueryParam("region")
    heritage := c.QueryParam("heritage")

    var festivals interface{}
    var err error

    switch {
    case region != "":
        festivals, err = h.queries.ListFestivalsByRegion(ctx, region)
    case heritage != "":
        festivals, err = h.queries.ListFestivalsByHeritage(ctx, heritage)
    default:
        festivals, err = h.queries.ListFestivals(ctx)
    }

    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "failed to fetch festivals")
    }

    return c.JSON(http.StatusOK, festivals)
}

func (h *Handler) ListUpcomingFestivals(c echo.Context) error {
    ctx := c.Request().Context()

    festivals, err := h.queries.ListUpcomingFestivals(ctx)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "failed to fetch upcoming festivals")
    }

    return c.JSON(http.StatusOK, festivals)
}

func (h *Handler) GetFestival(c echo.Context) error {
    ctx := c.Request().Context()
    slug := c.Param("slug")

    festival, err := h.queries.GetFestivalBySlug(ctx, slug)
    if err != nil {
        return echo.NewHTTPError(http.StatusNotFound, "festival not found")
    }

    return c.JSON(http.StatusOK, festival)
}
