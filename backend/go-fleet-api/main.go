package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/Karthikeya-Akhandam/fleetsignal/backend/go-fleet-api/config"
	"github.com/Karthikeya-Akhandam/fleetsignal/backend/go-fleet-api/repository"
	"github.com/Karthikeya-Akhandam/fleetsignal/backend/go-fleet-api/routers"
)

func main() {
	// Load .env if it exists
	if err := godotenv.Load("../../.env"); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	cfg := config.LoadConfig()

	if err := repository.InitPostgres(cfg); err != nil {
		log.Fatalf("Failed to connect to Postgres: %v", err)
	}
	defer repository.ClosePostgres()

	if err := repository.InitClickHouse(cfg); err != nil {
		log.Fatalf("Failed to connect to ClickHouse: %v", err)
	}
	defer repository.CloseClickHouse()

	router := gin.Default()

	routers.RegisterRoutes(router)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "healthy",
			"service": "fleet-api",
		})
	})

	log.Printf("Starting Fleet API on port %s", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
