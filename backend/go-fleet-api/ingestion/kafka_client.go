package ingestion

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/Karthikeya-Akhandam/fleetsignal/backend/go-fleet-api/config"
)

type KafkaConsumers struct {
	TelemetryConsumer *kafka.Consumer
	IncidentConsumer  *kafka.Consumer
}

func InitConsumers(cfg config.Config) (*KafkaConsumers, error) {
	// Telemetry consumer setup
	telemetryConsumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  cfg.KafkaBrokers,
		"group.id":           "fleet-telemetry-group",
		"auto.offset.reset":  "earliest",
		"enable.auto.commit": false, // Manual commits for exactly-once processing
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create telemetry consumer: %w", err)
	}

	err = telemetryConsumer.SubscribeTopics([]string{"vehicle.telemetry"}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to subscribe to telemetry: %w", err)
	}

	// Incident consumer setup
	incidentConsumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  cfg.KafkaBrokers,
		"group.id":           "fleet-incident-group",
		"auto.offset.reset":  "earliest",
		"enable.auto.commit": false,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create incident consumer: %w", err)
	}

	err = incidentConsumer.SubscribeTopics([]string{"incident.alerts"}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to subscribe to incidents: %w", err)
	}

	return &KafkaConsumers{
		TelemetryConsumer: telemetryConsumer,
		IncidentConsumer:  incidentConsumer,
	}, nil
}

func (kc *KafkaConsumers) Close() {
	if kc.TelemetryConsumer != nil {
		kc.TelemetryConsumer.Close()
	}
	if kc.IncidentConsumer != nil {
		kc.IncidentConsumer.Close()
	}
}
