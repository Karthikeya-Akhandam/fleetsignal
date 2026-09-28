CREATE EXTENSION IF NOT EXISTS vector;

CREATE TABLE vehicles (
    vin VARCHAR(17) PRIMARY KEY,
    model VARCHAR(50) NOT NULL,
    year INT NOT NULL,
    capacity_kg INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE incidents (
    incident_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    vin VARCHAR(17) REFERENCES vehicles(vin),
    timestamp TIMESTAMP NOT NULL,
    fault_code VARCHAR(20) NOT NULL,
    severity VARCHAR(10) NOT NULL,
    resolved BOOLEAN DEFAULT FALSE,
    resolution_notes TEXT,
    embedding vector(768),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE work_orders (
    order_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    incident_id UUID REFERENCES incidents(incident_id),
    vin VARCHAR(17) REFERENCES vehicles(vin),
    status VARCHAR(20) DEFAULT 'PENDING',
    description TEXT NOT NULL,
    cost_estimate DECIMAL(10, 2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE ota_updates (
    update_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    vin VARCHAR(17) REFERENCES vehicles(vin),
    version VARCHAR(20) NOT NULL,
    status VARCHAR(20) DEFAULT 'PENDING',
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
