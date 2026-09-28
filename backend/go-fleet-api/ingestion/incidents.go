package ingestion

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
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

func StartIncidentConsumer(ctx context.Context, consumer *kafka.Consumer) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			msg, err := consumer.ReadMessage(100 * time.Millisecond)
			if err == nil {
				var inc IncidentMessage
				if err := json.Unmarshal(msg.Value, &inc); err == nil {
					// We immediately insert incidents, no batching needed as they are rare
					err := insertIncident(ctx, inc)
					if err == nil {
						consumer.CommitMessage(msg)
					} else {
						log.Printf("Failed to insert incident: %v", err)
					}
				}
			} else if !err.(kafka.Error).IsTimeout() {
				log.Printf("Incident consumer error: %v\n", err)
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
