package models

import (
	"encoding/json"
	"testing"
)

func TestTelemetryPayloadUnmarshal(t *testing.T) {
	jsonPayload := `{
		"timestamp": "2026-09-28T10:00:00Z",
		"vin": "WBA123456789",
		"latitude": 37.7749,
		"longitude": -122.4194,
		"speed_kmh": 65.5,
		"rpm": 2500,
		"engine_load": 45.2,
		"battery_voltage": 12.4,
		"fuel_level": 75.0,
		"battery_soc": 98.5,
		"error_code": "NONE",
		"hvac_status": "ON",
		"cargo_weight_kg": 500
	}`

	var payload TelemetryPayload
	err := json.Unmarshal([]byte(jsonPayload), &payload)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	if payload.VIN != "WBA123456789" {
		t.Errorf("Expected VIN 'WBA123456789', got '%s'", payload.VIN)
	}
	if payload.SpeedKMH != 65.5 {
		t.Errorf("Expected speed 65.5, got %f", payload.SpeedKMH)
	}
	if payload.RPM != 2500 {
		t.Errorf("Expected RPM 2500, got %d", payload.RPM)
	}
}
