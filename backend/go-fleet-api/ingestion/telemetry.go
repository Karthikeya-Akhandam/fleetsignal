package ingestion

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/Karthikeya-Akhandam/fleetsignal/backend/go-fleet-api/models"
	"github.com/Karthikeya-Akhandam/fleetsignal/backend/go-fleet-api/repository"
)

func StartTelemetryConsumer(ctx context.Context, consumer *kafka.Consumer) {
	batchSize := 1000
	var batch []models.TelemetryPayload

	for {
		select {
		case <-ctx.Done():
			flushTelemetryBatch(batch)
			return
		default:
			msg, err := consumer.ReadMessage(100 * time.Millisecond)
			if err == nil {
				var payload models.TelemetryPayload
				if err := json.Unmarshal(msg.Value, &payload); err == nil {
					batch = append(batch, payload)
				}

				if len(batch) >= batchSize {
					if err := flushTelemetryBatch(batch); err == nil {
						consumer.CommitMessage(msg)
						batch = batch[:0] // clear batch
					} else {
						log.Printf("Failed to flush telemetry batch: %v", err)
					}
				}
			} else if !err.(kafka.Error).IsTimeout() {
				log.Printf("Telemetry consumer error: %v\n", err)
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
