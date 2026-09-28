from dataclasses import dataclass
from typing import Optional
from datetime import datetime

@dataclass
class TelemetryPayload:
    timestamp: str
    vin: str
    latitude: float
    longitude: float
    speed_kmh: float
    rpm: int
    engine_load: float
    battery_voltage: float
    fuel_level: float
    battery_soc: float
    error_code: Optional[str]
    hvac_status: str
    cargo_weight_kg: int

@dataclass
class IncidentPayload:
    timestamp: str
    vin: str
    fault_code: str
    severity: str
    description: str
