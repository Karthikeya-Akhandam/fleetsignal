.PHONY: help up down logs sim db build clean run

help: ## Show this help message
	@echo "FleetSignal Makefile Commands:"
	@echo "  make up      - Start all services and databases via Docker Compose"
	@echo "  make down    - Stop and remove all containers"
	@echo "  make logs    - View logs for all running services"
	@echo "  make sim     - Run the telemetry simulator (requires local python)"
	@echo "  make db      - Start ONLY the databases (Postgres, ClickHouse, Redis, Kafka)"
	@echo "  make build   - Rebuild all Docker images"
	@echo "  make clean   - Remove all containers, networks, and volumes (WARNING: Data loss)"
	@echo "  make run     - Clean, build, start all services, and run the simulator"

run: clean up
	@echo "Waiting for Kafka and APIs to initialize..."
	@sleep 15
	@echo "Starting simulator..."
	cd simulator && python main.py

up:
	cd infrastructure && docker-compose up -d

down:
	cd infrastructure && docker-compose down

logs:
	cd infrastructure && docker-compose logs -f

sim:
	cd simulator && python main.py

db:
	cd infrastructure && docker-compose up -d postgres clickhouse redis zookeeper kafka

build:
	cd infrastructure && docker-compose build

clean:
	cd infrastructure && docker-compose down -v
