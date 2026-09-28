package repository

import (
	"context"
	"fmt"
	"time"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/Karthikeya-Akhandam/fleetsignal/backend/go-fleet-api/config"
)

var CHConn driver.Conn

func InitClickHouse(cfg config.Config) error {
	addr := fmt.Sprintf("%s:%s", cfg.ClickHouseHost, cfg.ClickHousePort)
	
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{addr},
		Auth: clickhouse.Auth{
			Database: cfg.ClickHouseDB,
			Username: cfg.ClickHouseUser,
			Password: cfg.ClickHousePassword,
		},
		DialTimeout:     time.Second * 5,
		MaxOpenConns:    5,
		MaxIdleConns:    2,
	})

	if err != nil {
		return err
	}

	if err := conn.Ping(context.Background()); err != nil {
		return err
	}

	CHConn = conn
	return nil
}

func CloseClickHouse() {
	if CHConn != nil {
		CHConn.Close()
	}
}
