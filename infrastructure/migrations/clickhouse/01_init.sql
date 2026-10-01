CREATE TABLE IF NOT EXISTS fleetsignal.vehicle_telemetry (
    timestamp DateTime64(3),
    vin String,
    latitude Float64,
    longitude Float64,
    speed_kmh Float32,
    rpm Int32,
    engine_load Float32,
    battery_voltage Float32,
    fuel_level Float32,
    battery_soc Float32,
    error_code String,
    hvac_status String,
    cargo_weight_kg Int32
) ENGINE = MergeTree()
ORDER BY (vin, timestamp);
