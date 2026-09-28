import random
from datetime import datetime, timezone
import math
from typing import Optional
from models import TelemetryPayload
import config

class Vehicle:
    def __init__(self, vin: str):
        self.vin = vin
        self.latitude = random.uniform(34.0, 42.0) # Approx US coordinates
        self.longitude = random.uniform(-118.0, -73.0)
        self.speed_kmh = random.uniform(0.0, 105.0)
        self.rpm = int(random.uniform(800, 2500))
        self.engine_load = random.uniform(10.0, 95.0)
        self.battery_voltage = random.uniform(12.2, 14.4)
        self.fuel_level = random.uniform(10.0, 100.0)
        self.battery_soc = random.uniform(20.0, 100.0)
        self.cargo_weight_kg = int(random.uniform(0, 20000))
        self.hvac_status = "ON" if random.random() > 0.5 else "OFF"
        self.error_code: Optional[str] = None
        
        # Internal physics state
        self.heading_rad = random.uniform(0, 2 * math.pi)

    def tick(self) -> TelemetryPayload:
        # Simulate physics step
        delta_t = 1.0 / config.TICK_RATE_HZ
        
        # Random walk for speed and heading
        self.speed_kmh += random.uniform(-5.0, 5.0)
        self.speed_kmh = max(0.0, min(self.speed_kmh, 120.0))
        
        if self.speed_kmh > 0:
            self.heading_rad += random.uniform(-0.1, 0.1)
            # Distance in km
            distance_km = (self.speed_kmh * delta_t) / 3600.0
            # Rough approximation of degrees roughly 111km per degree
            self.latitude += (math.cos(self.heading_rad) * distance_km) / 111.0
            self.longitude += (math.sin(self.heading_rad) * distance_km) / (111.0 * math.cos(math.radians(self.latitude)))
            
            # RPM and Load correlates with speed and cargo
            self.rpm = int(800 + (self.speed_kmh / 120.0) * 1700)
            self.engine_load = 20.0 + (self.cargo_weight_kg / 20000.0) * 50.0 + random.uniform(-5, 5)
            self.fuel_level -= 0.001 * self.engine_load * delta_t
        else:
            self.rpm = 800
            self.engine_load = random.uniform(5.0, 15.0)
            
        self.fuel_level = max(0.0, self.fuel_level)

        return TelemetryPayload(
            timestamp=datetime.now(timezone.utc).isoformat(),
            vin=self.vin,
            latitude=round(self.latitude, 6),
            longitude=round(self.longitude, 6),
            speed_kmh=round(self.speed_kmh, 2),
            rpm=self.rpm,
            engine_load=round(self.engine_load, 2),
            battery_voltage=round(self.battery_voltage, 2),
            fuel_level=round(self.fuel_level, 2),
            battery_soc=round(self.battery_soc, 2),
            error_code=self.error_code,
            hvac_status=self.hvac_status,
            cargo_weight_kg=self.cargo_weight_kg
        )
