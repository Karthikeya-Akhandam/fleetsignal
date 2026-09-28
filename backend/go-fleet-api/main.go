package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/Karthikeya-Akhandam/fleetsignal/backend/go-fleet-api/config"
	"github.com/Karthikeya-Akhandam/fleetsignal/backend/go-fleet-api/repository"
	"github.com/Karthikeya-Akhandam/fleetsignal/backend/go-fleet-api/routers"
	"github.com/Karthikeya-Akhandam/fleetsignal/backend/go-fleet-api/ingestion"
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

	// Start Kafka Consumers
	consumers, err := ingestion.InitConsumers(cfg)
	if err != nil {
		log.Fatalf("Failed to init consumers: %v", err)
	}
	defer consumers.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go ingestion.StartTelemetryConsumer(ctx, consumers.TelemetryReader)
	go ingestion.StartIncidentConsumer(ctx, consumers.IncidentReader)

	// Graceful shutdown
	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}

	go func() {
		log.Printf("Starting Fleet API on port %s", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down Fleet API...")

	cancel() // Stops Kafka consumers

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Fleet API exiting")
}
