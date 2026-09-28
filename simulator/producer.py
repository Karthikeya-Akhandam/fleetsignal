import json
from confluent_kafka import Producer
import config
from models import TelemetryPayload, IncidentPayload
import dataclasses

class KafkaPublisher:
    def __init__(self):
        conf = {
            'bootstrap.servers': config.KAFKA_BOOTSTRAP_SERVERS,
            'client.id': 'simulator-producer',
            'linger.ms': 5, # Batching delay for higher throughput
            'batch.size': 16384
        }
        self.producer = Producer(conf)

    def delivery_report(self, err, msg):
        if err is not None:
            print(f"Message delivery failed: {err}")
            
    def publish_telemetry(self, payload: TelemetryPayload):
        data = json.dumps(dataclasses.asdict(payload)).encode('utf-8')
        # Partitioning by VIN ensures all messages for a single vehicle go to the same partition
        self.producer.produce(
            topic=config.TELEMETRY_TOPIC,
            key=payload.vin.encode('utf-8'),
            value=data,
            on_delivery=self.delivery_report
        )
        self.producer.poll(0) # Non-blocking poll for delivery events

    def publish_incident(self, payload: IncidentPayload):
        data = json.dumps(dataclasses.asdict(payload)).encode('utf-8')
        self.producer.produce(
            topic=config.INCIDENT_TOPIC,
            key=payload.vin.encode('utf-8'),
            value=data,
            on_delivery=self.delivery_report
        )
        self.producer.poll(0)

    def flush(self):
        self.producer.flush()
