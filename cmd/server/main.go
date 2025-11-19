package main

import (
	"agen_edc/config"
	"agen_edc/internal/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// load .env if present
	_ = godotenv.Load()
	cfg := config.Load()
	r := routes.SetupRouter(cfg)
	log.Printf("Starting server on :%s", cfg.AppPort)
	if err := r.Run(":" + cfg.AppPort); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
