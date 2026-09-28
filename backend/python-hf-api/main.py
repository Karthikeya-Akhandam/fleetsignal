from fastapi import FastAPI
import uvicorn
from config import config
from db import db_clients
from routers import router as alpha_router

app = FastAPI(
    title="FleetSignal Hedge Fund API",
    description="Privacy-preserved Alternative Data API for Hedge Funds",
    version="1.0.0"
)

@app.on_event("startup")
async def startup_event():
    # Initialize DB connections
    db_clients.connect()
    print("Database connections established.")

@app.on_event("shutdown")
async def shutdown_event():
    db_clients.close()
    print("Database connections closed.")

app.include_router(alpha_router, prefix="/api/v1")

@app.get("/health")
async def health_check():
    return {"status": "healthy", "service": "hf-api"}

if __name__ == "__main__":
    uvicorn.run("main:app", host="0.0.0.0", port=config.PORT, reload=True)
