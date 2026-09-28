package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/Karthikeya-Akhandam/fleetsignal/backend/go-fleet-api/config"
)

var PGPool *pgxpool.Pool

func InitPostgres(cfg config.Config) error {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresDB)
	
	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return err
	}
	// Limit connections to save RAM
	poolConfig.MaxConns = 10

	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return err
	}

	if err := pool.Ping(context.Background()); err != nil {
		return err
	}

	PGPool = pool
	return nil
}

func ClosePostgres() {
	if PGPool != nil {
		PGPool.Close()
	}
}
