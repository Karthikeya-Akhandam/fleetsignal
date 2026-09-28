import os
from dotenv import load_dotenv

load_dotenv()

# Kafka Configuration
KAFKA_BOOTSTRAP_SERVERS = os.getenv("KAFKA_BOOTSTRAP_SERVERS", "localhost:9092")
TELEMETRY_TOPIC = "vehicle.telemetry"
INCIDENT_TOPIC = "incident.alerts"

# Simulation Settings
NUM_VEHICLES = int(os.getenv("SIM_NUM_VEHICLES", "10000"))
TICK_RATE_HZ = float(os.getenv("SIM_TICK_RATE_HZ", "1.0"))

# Physics Constants
GRAVITY = 9.81
AIR_DENSITY = 1.225
DRAG_COEFFICIENT = 0.65 # Heavy duty truck
FRONTAL_AREA = 10.0
