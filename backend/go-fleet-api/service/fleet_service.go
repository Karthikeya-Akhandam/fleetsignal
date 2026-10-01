package service

import (
	"context"
	"time"
	"github.com/Karthikeya-Akhandam/fleetsignal/backend/go-fleet-api/models"
	"github.com/Karthikeya-Akhandam/fleetsignal/backend/go-fleet-api/repository"
)

type FleetService struct{}

func NewFleetService() *FleetService {
	return &FleetService{}
}

func (s *FleetService) GetActiveIncidents(ctx context.Context) ([]models.Incident, error) {
	query := `
		SELECT incident_id, vin, timestamp, fault_code, severity, resolved, COALESCE(resolution_notes, ''), created_at 
		FROM incidents 
		WHERE resolved = false 
		ORDER BY timestamp DESC
	`
	rows, err := repository.PGPool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var incidents []models.Incident
	for rows.Next() {
		var i models.Incident
		if err := rows.Scan(&i.IncidentID, &i.VIN, &i.Timestamp, &i.FaultCode, &i.Severity, &i.Resolved, &i.ResolutionNotes, &i.CreatedAt); err != nil {
			return nil, err
		}
		incidents = append(incidents, i)
	}
	return incidents, nil
}

func (s *FleetService) ResolveIncident(ctx context.Context, incidentID string, notes string) error {
	query := `
		UPDATE incidents 
		SET resolved = true, resolution_notes = $1 
		WHERE incident_id = $2
	`
	_, err := repository.PGPool.Exec(ctx, query, notes, incidentID)
	return err
}

type FleetStats struct {
	ActiveVehicles   int `json:"active_vehicles"`
	CriticalIncidents int `json:"critical_incidents"`
	ResolvedOTA      int `json:"resolved_ota"`
}

func (s *FleetService) GetFleetStats(ctx context.Context) (*FleetStats, error) {
	var stats FleetStats
	
	queryPG := `SELECT count(*) FROM incidents WHERE resolved = false`
	_ = repository.PGPool.QueryRow(ctx, queryPG).Scan(&stats.CriticalIncidents)
	
	queryPGResolved := `SELECT count(*) FROM incidents WHERE resolved = true`
	_ = repository.PGPool.QueryRow(ctx, queryPGResolved).Scan(&stats.ResolvedOTA)
	
	queryCH := `SELECT uniq(vin) FROM fleetsignal.vehicle_telemetry WHERE timestamp > now() - interval 1 hour`
	var active uint64
	if err := repository.CHConn.QueryRow(ctx, queryCH).Scan(&active); err == nil {
		stats.ActiveVehicles = int(active)
	} else {
		stats.ActiveVehicles = 0
	}
	
	return &stats, nil
}

type TelemetryHistory struct {
	Time      string `json:"time"`
	Incidents int    `json:"incidents"`
	Updates   int    `json:"updates"`
}

func (s *FleetService) GetTelemetryHistory(ctx context.Context) ([]TelemetryHistory, error) {
	queryCH := `
		SELECT 
			toStartOfHour(timestamp) as hr,
			count(*) as updates
		FROM fleetsignal.vehicle_telemetry
		WHERE timestamp >= now() - interval 6 hour
		GROUP BY hr
		ORDER BY hr ASC
	`
	rows, err := repository.CHConn.Query(ctx, queryCH)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var history []TelemetryHistory
	for rows.Next() {
		var hr time.Time
		var updates uint64
		if err := rows.Scan(&hr, &updates); err == nil {
			history = append(history, TelemetryHistory{
				Time:      hr.Format("15:00"),
				Updates:   int(updates),
				Incidents: int(updates / 10), // Rough approximation for demonstration
			})
		}
	}
	
	if len(history) == 0 {
		return []TelemetryHistory{}, nil
	}
	
	return history, nil
}
