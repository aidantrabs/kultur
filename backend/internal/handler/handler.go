package handler

import (
    "github.com/aidantrabs/kultur/backend/internal/db"
    "github.com/aidantrabs/kultur/backend/internal/email"
    "github.com/aidantrabs/kultur/backend/internal/service"
    "github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
    pool          *pgxpool.Pool
    festivals     *service.FestivalService
    memories      *service.MemoryService
    subscriptions *service.SubscriptionService
}

type Config struct {
    ResendAPIKey string
    FromEmail    string
    BaseURL      string
}

func New(pool *pgxpool.Pool, cfg Config) *Handler {
    queries := db.New(pool)
    festivalSvc := service.NewFestivalService(queries)

    emailSvc := email.NewService(email.Config{
        APIKey:    cfg.ResendAPIKey,
        FromEmail: cfg.FromEmail,
        BaseURL:   cfg.BaseURL,
    })

    return &Handler{
        pool:          pool,
        festivals:     festivalSvc,
        memories:      service.NewMemoryService(queries, festivalSvc),
        subscriptions: service.NewSubscriptionService(queries, emailSvc),
    }
}
