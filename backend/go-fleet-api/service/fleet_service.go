package service

import (
	"context"
	"github.com/Karthikeya-Akhandam/fleetsignal/backend/go-fleet-api/models"
	"github.com/Karthikeya-Akhandam/fleetsignal/backend/go-fleet-api/repository"
)

type FleetService struct{}

func NewFleetService() *FleetService {
	return &FleetService{}
}

func (s *FleetService) GetActiveIncidents(ctx context.Context) ([]models.Incident, error) {
	query := `
		SELECT incident_id, vin, timestamp, fault_code, severity, resolved, resolution_notes, created_at 
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
