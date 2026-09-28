import time
import random
import string
import argparse
from typing import List
from vehicle import Vehicle
from events import EventGenerator
from producer import KafkaPublisher
import config

def generate_vin() -> str:
    """Generates a random 17-character VIN"""
    chars = string.ascii_uppercase + string.digits
    # Avoid I, O, Q in VINs
    chars = ''.join(c for c in chars if c not in 'IOQ')
    return ''.join(random.choice(chars) for _ in range(17))

def main():
    parser = argparse.ArgumentParser(description='Fleet Data Simulator')
    parser.add_argument('--vehicles', type=int, default=config.NUM_VEHICLES, help='Number of vehicles to simulate')
    parser.add_argument('--tick-rate', type=float, default=config.TICK_RATE_HZ, help='Simulation ticks per second')
    args = parser.parse_args()

    print(f"Initializing simulator with {args.vehicles} vehicles at {args.tick_rate}Hz...")
    
    # Initialize components
    vehicles: List[Vehicle] = [Vehicle(generate_vin()) for _ in range(args.vehicles)]
    event_gen = EventGenerator()
    publisher = KafkaPublisher()
    
    print("Simulator running. Press Ctrl+C to stop.")
    
    tick_duration = 1.0 / args.tick_rate
    
    try:
        while True:
            start_time = time.time()
            
            for vehicle in vehicles:
                # 1. Update vehicle state
                telemetry = vehicle.tick()
                
                # 2. Check for anomalies
                incident = event_gen.evaluate_vehicle(vehicle)
                
                # 3. Publish data
                publisher.publish_telemetry(telemetry)
                if incident:
                    print(f"\n[INCIDENT] VIN: {incident.vin} | Fault: {incident.fault_code} | Severity: {incident.severity}")
                    publisher.publish_incident(incident)
            
            # Maintain tick rate
            elapsed = time.time() - start_time
            sleep_time = tick_duration - elapsed
            if sleep_time > 0:
                time.sleep(sleep_time)
                
    except KeyboardInterrupt:
        print("\nStopping simulator...")
    finally:
        publisher.flush()
        print("Simulator stopped.")

if __name__ == "__main__":
    main()
