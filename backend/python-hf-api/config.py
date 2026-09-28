import os
from dotenv import load_dotenv

load_dotenv()

class Config:
    POSTGRES_USER = os.getenv("POSTGRES_USER", "motorq")
    POSTGRES_PASSWORD = os.getenv("POSTGRES_PASSWORD", "motorq_password")
    POSTGRES_DB = os.getenv("POSTGRES_DB", "fleetsignal")
    POSTGRES_HOST = os.getenv("POSTGRES_HOST", "localhost")
    POSTGRES_PORT = os.getenv("POSTGRES_PORT", "5432")
    
    CLICKHOUSE_USER = os.getenv("CLICKHOUSE_USER", "motorq")
    CLICKHOUSE_PASSWORD = os.getenv("CLICKHOUSE_PASSWORD", "motorq_password")
    CLICKHOUSE_DB = os.getenv("CLICKHOUSE_DB", "fleetsignal")
    CLICKHOUSE_HOST = os.getenv("CLICKHOUSE_HOST", "localhost")
    CLICKHOUSE_PORT = int(os.getenv("CLICKHOUSE_PORT", "8123"))

    PORT = int(os.getenv("HF_API_PORT", "8000"))
    
config = Config()
