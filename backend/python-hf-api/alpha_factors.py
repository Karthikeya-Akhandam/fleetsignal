from db import db_clients
from privacy import PrivacyEngine
from schemas import AnonymizedFactor
from datetime import datetime

privacy_engine = PrivacyEngine()

class AlphaFactorEngine:
    def get_freight_tonnage_index(self, start_time: str, end_time: str):
        """
        Alpha Factor 1: Freight Tonnage Index
        Aggregates cargo_weight_kg by geohash (precision 4) to detect supply chain density.
        """
        # ClickHouse query: aggregate weight by geohash, enforce k-anonymity
        query = f"""
        SELECT 
            geohashEncode(longitude, latitude, 4) as geohash,
            toStartOfHour(timestamp) as time_bucket,
            count(DISTINCT vin) as k_count,
            sum(cargo_weight_kg) as total_tonnage
        FROM vehicle_telemetry
        WHERE timestamp >= '{start_time}' AND timestamp <= '{end_time}'
        GROUP BY geohash, time_bucket
        HAVING k_count >= {privacy_engine.k_threshold}
        """
        
        result = db_clients.ch_client.query(query)
        
        processed_data = []
        for row in result.result_rows:
            geohash, time_bucket, k_count, total_tonnage = row
            # Apply Differential Privacy noise to the aggregated tonnage
            noisy_tonnage = privacy_engine.add_laplace_noise(total_tonnage, sensitivity=20000.0)
            
            processed_data.append(AnonymizedFactor(
                geohash=geohash,
                timestamp=time_bucket,
                k_count=k_count,
                factor_value=max(0, noisy_tonnage)
            ))
            
        return processed_data

    def get_grid_stress_demand(self, start_time: str, end_time: str):
        """
        Alpha Factor 3: Grid-Stress MW Demand
        Aggregates EV battery load by geohash to predict local grid stress.
        """
        query = f"""
        SELECT 
            geohashEncode(longitude, latitude, 5) as geohash,
            toStartOfHour(timestamp) as time_bucket,
            count(DISTINCT vin) as k_count,
            avg(battery_soc) as avg_soc,
            sum(engine_load) as total_load
        FROM vehicle_telemetry
        WHERE timestamp >= '{start_time}' AND timestamp <= '{end_time}'
        GROUP BY geohash, time_bucket
        HAVING k_count >= {privacy_engine.k_threshold}
        """
        
        result = db_clients.ch_client.query(query)
        
        processed_data = []
        for row in result.result_rows:
            geohash, time_bucket, k_count, avg_soc, total_load = row
            
            # Combine logic: if load is high and SOC is low, charging demand is high
            stress_index = total_load / (avg_soc + 1.0)
            noisy_index = privacy_engine.add_laplace_noise(stress_index, sensitivity=10.0)
            
            processed_data.append(AnonymizedFactor(
                geohash=geohash,
                timestamp=time_bucket,
                k_count=k_count,
                factor_value=max(0, noisy_index)
            ))
            
        return processed_data

alpha_engine = AlphaFactorEngine()
