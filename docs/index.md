<div class="hero-section">
  <h1 class="hero-title">FleetSignal Intelligence</h1>
  <p class="hero-subtitle">Privacy-Preserving Predictive Maintenance & Fleet Stress Intelligence. Ingest millions of telemetry events in real-time without exposing sensitive driver data.</p>
  <a href="https://drive.google.com/drive/folders/11s27XqaA0thk6hIcXIZipTtEbOMBbODU?usp=sharing" class="cta-button" target="_blank">View Hackathon Submission (Google Drive)</a>
</div>

## The Problem

Fleet managers face a paradox: they must ingest and analyse millions of telemetry events in real time to prevent costly breakdowns, but doing so often violates strict data privacy laws (e.g., GDPR, DPDP Act 2023).

<div class="card-grid">
  <div class="feature-card">
    <div class="feature-icon">⚠️</div>
    <div class="feature-title">The Cost of Ignorance</div>
    <div class="feature-desc">Operating in high-variance temperature regions without predictive maintenance leads to accelerated hardware degradation.</div>
    <div class="metric-highlight">34% Faster Battery Drain</div>
  </div>
  <div class="feature-card">
    <div class="feature-icon">🔒</div>
    <div class="feature-title">The Privacy Barrier</div>
    <div class="feature-desc">Existing telematics solutions expose raw GPS tracks and vehicle identification numbers, violating data protection regulations.</div>
    <div class="metric-highlight">Zero Compliance</div>
  </div>
</div>

---

## Our Solution

**FleetSignal** solves this by providing a highly scalable, event-driven telematics platform that ingests raw telemetry, applies real-time differential privacy and $k$-anonymity, and serves aggregate predictive insights.

<div class="card-grid">
  <div class="feature-card">
    <div class="feature-icon">📊</div>
    <div class="feature-title">Predictive Feature Importance</div>
    <div class="feature-desc">Displays geospatial features sorted by their predictive importance for fleet stress using our Python ML backend.</div>
  </div>
  <div class="feature-card">
    <div class="feature-icon">🛡️</div>
    <div class="feature-title">Dynamic k-Anonymity</div>
    <div class="feature-desc">If fewer than k=5 vehicles occupy a geohash, the data is entirely suppressed to prevent reverse-engineering.</div>
  </div>
  <div class="feature-card">
    <div class="feature-icon">⚡</div>
    <div class="feature-title">Real-time Anomaly Alerts</div>
    <div class="feature-desc">Go Fleet-API consumes Kafka incident streams instantly, saving severity-based anomalies directly to Postgres.</div>
  </div>
</div>

---

## System Architecture

FleetSignal employs a modular, event-driven microservice architecture designed to handle extremely high-throughput append-only data while persisting stateful operations consistently.

<div class="diagram-container">
```mermaid
graph TD
    subgraph "Edge / Vehicles"
        Sim[Python Telemetry Simulator]
    end

    subgraph "Message Broker"
        Kafka[Apache Kafka]
    end

    subgraph "Core Processing"
        GoAPI[Go Fleet API Consumer]
    end

    subgraph "Databases"
        CH[(ClickHouse\nTime-Series)]
        PG[(PostgreSQL\nRelational)]
    end

    subgraph "Intelligence & Privacy"
        PyAPI[Python HF-API\nPrivacy Engine]
    end

    subgraph "Presentation"
        UI[React Dashboard]
    end

    Sim -- MQTT / TCP --> Kafka
    Kafka -- Batch Consume --> GoAPI
    GoAPI -- Insert Telemetry --> CH
    GoAPI -- Insert Incidents --> PG
    UI -- Fetch Insights --> PyAPI
    PyAPI -- Aggregate Queries --> CH
    UI -- Fetch Alerts --> GoAPI
```
</div>

### Technology Stack Justification

| Layer | Choice | Why This, and What We Rejected |
|-------|--------|---------------------------------|
| **Ingestion** | Apache Kafka | High-throughput, partitioned ordering. Rejected RabbitMQ (can't handle replay easily). |
| **Time-Series** | ClickHouse | Columnar aggregation speed. Rejected PostgreSQL for telemetry due to insert bottlenecks. |
| **Backend API** | Go (Golang) | Excellent concurrency for Kafka consuming. Rejected Node.js for CPU bound tasks. |
| **Privacy / ML** | Python (FastAPI)| Native Pandas/Numpy support for differential privacy math. |
| **Frontend** | React / Vite | State management and Recharts capabilities. |

---

## Low-Level Design & Privacy Flow

The privacy engine bridges the gap between high-frequency automotive telemetry and strict privacy laws, proving that actionable ML insights can be generated without compromising driver confidentiality.

<div class="diagram-container">
```mermaid
sequenceDiagram
    participant UI as React Dashboard
    participant Py as Python HF-API
    participant CH as ClickHouse
    
    UI->>Py: GET /api/v1/alpha/freight-tonnage
    Py->>CH: SELECT sum(value) GROUP BY geohash
    CH-->>Py: Raw Aggregates
    Py->>Py: Check k-anonymity (count >= 5)
    Py->>Py: Apply Laplace Noise to value
    Py-->>UI: Anonymised JSON
    UI->>UI: Render Recharts BarChart
```
</div>

### Differential Privacy (Laplace Noise)
Differential privacy provides a mathematical guarantee that the output of a query does not significantly change whether any single individual's data is included or not. We add statistical **Laplace noise** to all raw numerical aggregates (like Grid Stress or Freight Tonnage values).

### Performance Benchmarks

| NFR | Target | Achieved | How Measured |
|-----|--------|----------|--------------|
| **Ingest Throughput** | 1,000+ events/sec | > 5,000 / sec | Kafka Consumer Lag limits |
| **API Latency** | p95 < 200 ms | ~ 45 ms | Chrome DevTools Network Tab |
| **Availability** | Resilient broker | Yes | Docker auto-restart |
