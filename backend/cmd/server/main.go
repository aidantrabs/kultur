package main

import (
		"log"

		"github.com/aidantrabs/trinbago-hackathon/backend/internal/config"
)

func main() {
		cfg, err := config.Load()
		if err != nil {
				log.Fatal("Failed to load config:", err)
		}

		log.Printf("Server configured on port %s", cfg.Port)
		log.Printf("Database URL: %s", cfg.DatabaseURL)
}
