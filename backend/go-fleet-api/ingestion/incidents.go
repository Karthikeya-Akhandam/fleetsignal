package ingestion

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/Karthikeya-Akhandam/fleetsignal/backend/go-fleet-api/repository"
	"github.com/Karthikeya-Akhandam/fleetsignal/backend/go-fleet-api/service"
)

type IncidentMessage struct {
	Timestamp   string `json:"timestamp"`
	VIN         string `json:"vin"`
	FaultCode   string `json:"fault_code"`
	Severity    string `json:"severity"`
	Description string `json:"description"`
}

func StartIncidentConsumer(ctx context.Context, reader *kafka.Reader) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			msg, err := reader.FetchMessage(ctx)
			if err != nil {
				if err == context.Canceled {
					return
				}
				log.Printf("Incident consumer error: %v\n", err)
				continue
			}

			var inc IncidentMessage
			if err := json.Unmarshal(msg.Value, &inc); err == nil {
				err := insertIncident(ctx, inc)
				if err == nil {
					reader.CommitMessages(ctx, msg)
				} else {
					log.Printf("Failed to insert incident: %v", err)
				}
			}
		}
	}
}

func insertIncident(ctx context.Context, inc IncidentMessage) error {
	ts, _ := time.Parse(time.RFC3339, inc.Timestamp)
	
	embeddingText := fmt.Sprintf("Incident %s: %s", inc.FaultCode, inc.Description)
	embedding, _ := service.GenerateEmbedding(ctx, embeddingText)
	
	query := `
		INSERT INTO incidents (vin, timestamp, fault_code, severity, embedding)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := repository.PGPool.Exec(ctx, query, inc.VIN, ts, inc.FaultCode, inc.Severity, embedding)
	return err
}
