package main

import (
    "context"
    "log"
    "net/http"
    "strings"
    "time"

    "github.com/aidantrabs/kultur/backend/internal/config"
    "github.com/aidantrabs/kultur/backend/internal/db"
    "github.com/aidantrabs/kultur/backend/internal/handler"
    "github.com/aidantrabs/kultur/backend/internal/middleware"
    "github.com/labstack/echo/v4"
    echomw "github.com/labstack/echo/v4/middleware"
)

func main() {
    ctx := context.Background()

    cfg, err := config.Load()
    if err != nil {
        log.Fatal("failed to load config:", err)
    }

    pool, err := db.Connect(ctx, cfg.DatabaseURL)
    if err != nil {
        log.Fatal("failed to connect to database:", err)
    }
    defer pool.Close()

    h := handler.New(pool, handler.Config{
        ResendAPIKey: cfg.ResendAPIKey,
        FromEmail:    cfg.FromEmail,
        BaseURL:      cfg.BaseURL,
    })

    e := echo.New()
    e.HideBanner = true

    // global middleware
    e.Use(echomw.Logger())
    e.Use(echomw.Recover())
    e.Use(echomw.CORSWithConfig(echomw.CORSConfig{
        AllowOrigins:     strings.Split(cfg.AllowedOrigins, ","),
        AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodOptions},
        AllowHeaders:     []string{echo.HeaderContentType, "X-API-Key"},
        AllowCredentials: true,
    }))

    // rate limiters
    memoryRateLimiter := middleware.NewRateLimiter(5, time.Hour)       // 5/hour for memories
    subscribeRateLimiter := middleware.NewRateLimiter(10, time.Hour)  // 10/hour for subscribe

    // health check
    e.GET("/health", h.Health)

    // public api routes
    api := e.Group("/api")

    // festivals (public)
    api.GET("/festivals", h.ListFestivals)
    api.GET("/festivals/upcoming", h.ListUpcomingFestivals)
    api.GET("/festivals/calendar", h.ListFestivalsByYear)
    api.GET("/festivals/:slug", h.GetFestival)
    api.GET("/festivals/:slug/dates", h.GetFestivalDates)
    api.GET("/festivals/:slug/memories", h.ListMemoriesByFestival)

    // memories (public, rate limited)
    api.POST("/memories", h.CreateMemory, memoryRateLimiter.Middleware())

    // subscriptions (public)
    api.POST("/subscribe", h.Subscribe, subscribeRateLimiter.Middleware())
    api.GET("/subscribe/confirm/:token", h.ConfirmSubscription)
    api.GET("/unsubscribe/:token", h.Unsubscribe)

    // admin routes (protected)
    admin := api.Group("/admin", middleware.APIKeyAuth(cfg.AdminAPIKey))

    // admin: memories
    admin.GET("/memories", h.ListAllMemories)
    admin.PATCH("/memories/:id", h.UpdateMemoryStatus)
    admin.DELETE("/memories/:id", h.DeleteMemory)

    // admin: subscriptions
    admin.GET("/subscriptions", h.ListAllSubscriptions)
    admin.DELETE("/subscriptions/:id", h.DeleteSubscription)

    // admin: festivals
    admin.POST("/festivals", h.CreateFestival)
    admin.PUT("/festivals/:id", h.UpdateFestival)
    admin.DELETE("/festivals/:id", h.DeleteFestival)

    // admin: festival dates
    admin.POST("/festival-dates", h.CreateFestivalDate)
    admin.PUT("/festival-dates/:id", h.UpdateFestivalDate)
    admin.DELETE("/festival-dates/:id", h.DeleteFestivalDate)

    log.Printf("server starting on port %s", cfg.Port)
    e.Logger.Fatal(e.Start(":" + cfg.Port))
}
