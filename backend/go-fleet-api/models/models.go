package models

import "time"

type TelemetryPayload struct {
	Timestamp      string  `json:"timestamp"`
	VIN            string  `json:"vin"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	SpeedKMH       float64 `json:"speed_kmh"`
	RPM            int     `json:"rpm"`
	EngineLoad     float64 `json:"engine_load"`
	BatteryVoltage float64 `json:"battery_voltage"`
	FuelLevel      float64 `json:"fuel_level"`
	BatterySOC     float64 `json:"battery_soc"`
	ErrorCode      *string `json:"error_code,omitempty"`
	HVACStatus     string  `json:"hvac_status"`
	CargoWeightKG  int     `json:"cargo_weight_kg"`
}

type Incident struct {
	IncidentID      string    `json:"incident_id"`
	VIN             string    `json:"vin"`
	Timestamp       time.Time `json:"timestamp"`
	FaultCode       string    `json:"fault_code"`
	Severity        string    `json:"severity"`
	Resolved        bool      `json:"resolved"`
	ResolutionNotes string    `json:"resolution_notes"`
	CreatedAt       time.Time `json:"created_at"`
}

type WorkOrder struct {
	OrderID      string    `json:"order_id"`
	IncidentID   string    `json:"incident_id"`
	VIN          string    `json:"vin"`
	Status       string    `json:"status"`
	Description  string    `json:"description"`
	CostEstimate float64   `json:"cost_estimate"`
	CreatedAt    time.Time `json:"created_at"`
}
