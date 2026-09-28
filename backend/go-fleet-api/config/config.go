package config

import (
	"os"
)

type Config struct {
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
	ClickHouseHost   string
	ClickHousePort   string
	ClickHouseUser   string
	ClickHousePassword string
	ClickHouseDB     string
	KafkaBrokers     string
	Port             string
}

func LoadConfig() Config {
	return Config{
		PostgresHost:     getEnv("POSTGRES_HOST", "localhost"),
		PostgresPort:     getEnv("POSTGRES_PORT", "5432"),
		PostgresUser:     getEnv("POSTGRES_USER", "motorq"),
		PostgresPassword: getEnv("POSTGRES_PASSWORD", "motorq_password"),
		PostgresDB:       getEnv("POSTGRES_DB", "fleetsignal"),
		ClickHouseHost:   getEnv("CLICKHOUSE_HOST", "localhost"),
		ClickHousePort:   getEnv("CLICKHOUSE_PORT", "9000"), // Native port for go client
		ClickHouseUser:   getEnv("CLICKHOUSE_USER", "motorq"),
		ClickHousePassword: getEnv("CLICKHOUSE_PASSWORD", "motorq_password"),
		ClickHouseDB:     getEnv("CLICKHOUSE_DB", "fleetsignal"),
		KafkaBrokers:     getEnv("KAFKA_BOOTSTRAP_SERVERS", "localhost:9092"),
		Port:             getEnv("FLEET_API_PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
