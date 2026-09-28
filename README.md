# FleetSignal: AI-Powered Autonomous Fleet Operations

**Winner Strategy for Motorq Hackathon** 🚀

FleetSignal is an end-to-end, privacy-preserved fleet management architecture designed to handle high-throughput telemetry data, predict vehicle maintenance via Machine Learning, and autonomously execute mitigation strategies using Agentic AI (LangGraph).

## The Architecture
Our system leverages the exact tech stack Motorq champions (Go, Kafka, React) augmented with state-of-the-art AI tooling.

1. **Simulator (`/simulator`)**: Generates 1,000s of vehicle telemetry points and incidents per second.
2. **High-Throughput Ingestion (`/backend/go-fleet-api`)**: A pure Go Kafka consumer pipeline processing real-time telemetry into ClickHouse (for time-series aggregations) and PostgreSQL (for ACID transactions).
3. **Analytics & Privacy Engine (`/backend/python-hf-api`)**: Python FastAPI backend employing `k-anonymity` and Laplace Noise for differential privacy, ensuring PII (like VINs) are protected before ML models calculate Alpha Factors.
4. **Agentic Copilot (`/agents/fleet-ops-agent`)**: A LangGraph state machine powered by Gemini. It queries pgvector embeddings to determine if a vehicle fault can be mitigated via an Over-The-Air (OTA) update or requires a physical work order.
5. **Glassmorphism Frontend (`/frontend`)**: A React/Vite dashboard leveraging pristine Vanilla CSS with dynamic micro-animations to wow the judges.

## Getting Started

### Prerequisites
- Docker & Docker Compose
- A Google Gemini API Key (`GOOGLE_API_KEY`)

### Quickstart (Hackathon Demo)
1. **Configure Environment:** Create a `.env` file in the root directory and add your API key:
   ```bash
   GOOGLE_API_KEY=your_gemini_api_key
   ```
2. **Launch the Infrastructure:**
   ```bash
   cd infrastructure
   docker-compose up --build
   ```
3. **Start the Simulator** (in a new terminal):
   ```bash
   cd simulator
   pip install -r requirements.txt
   python main.py
   ```
4. **Access the Dashboard:** Open [http://localhost:3000](http://localhost:3000)

## Why This Wins
*   **Scale**: ClickHouse + Kafka handles firehose data natively.
*   **Intelligence**: LangGraph enables autonomous decision-making, moving beyond simple static dashboards.
*   **Aesthetics**: Glassmorphism and bespoke React UI tokens.
*   **Resource Efficiency**: The entire microservice cluster runs in ~1.5GB of RAM via heavily constrained Docker deployments.
