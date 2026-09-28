import clickhouse_connect
import psycopg2
from config import config

class DatabaseClients:
    def __init__(self):
        self.ch_client = None
        self.pg_conn = None

    def connect(self):
        # Connect to ClickHouse for Telemetry querying
        self.ch_client = clickhouse_connect.get_client(
            host=config.CLICKHOUSE_HOST,
            port=config.CLICKHOUSE_PORT,
            username=config.CLICKHOUSE_USER,
            password=config.CLICKHOUSE_PASSWORD,
            database=config.CLICKHOUSE_DB
        )
        
        # Connect to Postgres for Incident/Vehicle querying
        self.pg_conn = psycopg2.connect(
            host=config.POSTGRES_HOST,
            port=config.POSTGRES_PORT,
            user=config.POSTGRES_USER,
            password=config.POSTGRES_PASSWORD,
            database=config.POSTGRES_DB
        )

    def close(self):
        if self.ch_client:
            self.ch_client.close()
        if self.pg_conn:
            self.pg_conn.close()

db_clients = DatabaseClients()
