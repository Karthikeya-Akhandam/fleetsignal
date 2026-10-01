# Architecture (High-Level Design)

## Overview
FleetSignal employs a modular, event-driven microservice architecture designed to handle extremely high-throughput append-only data while persisting stateful operations consistently.

- **Ingestion**: A Go-based simulator streams JSON telemetry via MQTT into a Kafka broker (`vehicle.telemetry`).
- **Processing**: A Go-based Fleet API and Kafka Consumer groups and aggregates the incoming streams.
- **Data Store**: 
  - **ClickHouse**: Handles high-throughput append-only time-series data for analytics.
  - **PostgreSQL**: Stores persistent relational entities (Vehicles, Users, Resolved Incidents).
  - **Redis**: Caches API responses and active WebSocket sessions.

## Technology Stack Justification

| Layer | Choice | Justification |
|-------|--------|---------------|
| Ingestion / Messaging | **Apache Kafka** | Chosen over RabbitMQ for its high-throughput partitioning, replayability, and durability for massive telemetry spikes. |
| Time-Series Store | **ClickHouse** | Columnar database heavily optimised for high-speed aggregations (e.g., `GROUP BY geohash`). Relational DBs would bottleneck on 100K+ inserts/sec. |
| Backend Services | **Go (Golang)** | Excellent concurrency (Goroutines) for processing massive Kafka streams with low memory footprint. |
| Frontend | **React (Vite) + Recharts** | Fast, component-driven UI with high-performance SVG charting. |

## Data Flow Pipeline
1. Simulated vehicles emit packets at `0.2Hz` to the `vehicle.telemetry` topic.
2. The Go-based Fleet Agent consumes these records, batches them, and issues high-speed inserts into ClickHouse (`vehicle_telemetry` table).
3. The Agent simultaneously listens to `incident.alerts` and records severity-based anomalies directly into PostgreSQL.
4. Python HF-API queries ClickHouse using aggregation logic.
5. The API filters data through the Privacy Engine (Laplace Noise + $k$-anonymity) and serves it securely to the React Dashboard.
