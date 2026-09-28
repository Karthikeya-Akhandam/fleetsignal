package ingestion

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/Karthikeya-Akhandam/fleetsignal/backend/go-fleet-api/models"
	"github.com/Karthikeya-Akhandam/fleetsignal/backend/go-fleet-api/repository"
)

func StartTelemetryConsumer(ctx context.Context, reader *kafka.Reader) {
	batchSize := 1000
	var batch []models.TelemetryPayload

	for {
		select {
		case <-ctx.Done():
			flushTelemetryBatch(batch)
			return
		default:
			// FetchMessage does not auto-commit offsets
			msg, err := reader.FetchMessage(ctx)
			if err != nil {
				if err == context.Canceled {
					return
				}
				log.Printf("Telemetry consumer error: %v\n", err)
				continue
			}

			var payload models.TelemetryPayload
			if err := json.Unmarshal(msg.Value, &payload); err == nil {
				batch = append(batch, payload)
			}

			if len(batch) >= batchSize {
				if err := flushTelemetryBatch(batch); err == nil {
					reader.CommitMessages(ctx, msg) // Commit the last message in the batch
					batch = batch[:0] // clear batch
				} else {
					log.Printf("Failed to flush telemetry batch: %v", err)
				}
			}
		}
	}
}

func flushTelemetryBatch(batch []models.TelemetryPayload) error {
	if len(batch) == 0 {
		return nil
	}

	batchConn, err := repository.CHConn.PrepareBatch(context.Background(), "INSERT INTO vehicle_telemetry")
	if err != nil {
		return err
	}

	for _, t := range batch {
		ts, _ := time.Parse(time.RFC3339, t.Timestamp)
		err := batchConn.Append(
			ts, t.VIN, t.Latitude, t.Longitude, float32(t.SpeedKMH), 
			int32(t.RPM), float32(t.EngineLoad), float32(t.BatteryVoltage), 
			float32(t.FuelLevel), float32(t.BatterySOC), t.ErrorCode, 
			t.HVACStatus, int32(t.CargoWeightKG),
		)
		if err != nil {
			return err
		}
	}

	return batchConn.Send()
}
