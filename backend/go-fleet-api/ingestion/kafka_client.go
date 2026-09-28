package ingestion

import (
	"github.com/segmentio/kafka-go"
	"github.com/Karthikeya-Akhandam/fleetsignal/backend/go-fleet-api/config"
)

type KafkaConsumers struct {
	TelemetryReader *kafka.Reader
	IncidentReader  *kafka.Reader
}

func InitConsumers(cfg config.Config) (*KafkaConsumers, error) {
	telemetryReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{cfg.KafkaBrokers},
		GroupID:   "fleet-telemetry-group",
		Topic:     "vehicle.telemetry",
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})

	incidentReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{cfg.KafkaBrokers},
		GroupID:   "fleet-incident-group",
		Topic:     "incident.alerts",
		MinBytes:  10e3,
		MaxBytes:  10e6,
	})

	return &KafkaConsumers{
		TelemetryReader: telemetryReader,
		IncidentReader:  incidentReader,
	}, nil
}

func (kc *KafkaConsumers) Close() {
	if kc.TelemetryReader != nil {
		kc.TelemetryReader.Close()
	}
	if kc.IncidentReader != nil {
		kc.IncidentReader.Close()
	}
}
