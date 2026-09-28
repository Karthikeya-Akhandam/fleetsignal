import random
from typing import Optional
from datetime import datetime, timezone
from models import IncidentPayload
from vehicle import Vehicle

class EventGenerator:
    def __init__(self):
        # Dictionary of possible faults and their severities
        self.fault_catalog = {
            "P0420": ("Catalyst System Efficiency Below Threshold", "LOW"),
            "P0300": ("Random/Multiple Cylinder Misfire Detected", "HIGH"),
            "C1446": ("Brake Switch Circuit Failure", "CRITICAL"),
            "U0100": ("Lost Communication with ECM/PCM", "CRITICAL"),
            "B1001": ("Airbag Deployment Command", "CRITICAL"),
            "P0A80": ("Replace Hybrid Battery Pack", "HIGH"),
        }

    def evaluate_vehicle(self, vehicle: Vehicle) -> Optional[IncidentPayload]:
        """
        Evaluates a vehicle's state to determine if an anomaly/incident should be generated.
        Returns an IncidentPayload if a fault occurs, else None.
        """
        # Very low probability per tick to simulate rare events
        if random.random() < 0.0001:
            fault_code = random.choice(list(self.fault_catalog.keys()))
            desc, severity = self.fault_catalog[fault_code]
            
            # Set the error code on the vehicle so it shows up in telemetry
            vehicle.error_code = fault_code
            
            # For specific critical faults, force physical symptoms
            if severity == "CRITICAL":
                vehicle.speed_kmh = max(0.0, vehicle.speed_kmh - 20.0) # sudden decel
            
            return IncidentPayload(
                timestamp=datetime.now(timezone.utc).isoformat(),
                vin=vehicle.vin,
                fault_code=fault_code,
                severity=severity,
                description=desc
            )
            
        return None
