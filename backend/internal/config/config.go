package config

import (
    "os"

    "github.com/joho/godotenv"
)

type Config struct {
    Port           string
    DatabaseURL    string
    ResendAPIKey   string
    AllowedOrigins string
    AdminAPIKey    string
    BaseURL        string
    FromEmail      string
}

func Load() (*Config, error) {
    // load .env (ignored in prod)
    godotenv.Load()

    return &Config{
        Port:           getEnv("PORT", "8080"),
        DatabaseURL:    getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/hackathon?sslmode=disable"),
        ResendAPIKey:   getEnv("RESEND_API_KEY", ""),
        AllowedOrigins: getEnv("ALLOWED_ORIGINS", "http://localhost:5173"),
        AdminAPIKey:    getEnv("ADMIN_API_KEY", ""),
        BaseURL:        getEnv("BASE_URL", "http://localhost:8080"),
        FromEmail:      getEnv("FROM_EMAIL", "onboarding@resend.dev"),
    }, nil
}

func getEnv(key, fallback string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }

    return fallback
}
