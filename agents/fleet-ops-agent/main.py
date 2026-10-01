from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
import uvicorn
import os
from graph import app as agent_app

from fastapi.middleware.cors import CORSMiddleware

app = FastAPI(title="Fleet Ops Agent API")

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

class IncidentRequest(BaseModel):
    incident_id: str
    vin: str
    fault_code: str
    severity: str
    description: str

@app.post("/api/v1/agent/evaluate-incident")
async def evaluate_incident(req: IncidentRequest):
    try:
        initial_state = {
            "messages": [],
            "incident_id": req.incident_id,
            "vin": req.vin,
            "fault_code": req.fault_code,
            "severity": req.severity,
            "description": req.description,
            "requires_ota": False,
            "requires_work_order": False,
            "ota_version": "",
            "work_order_cost": 0.0,
            "work_order_notes": "",
            "resolved": False
        }
        
        final_state = agent_app.invoke(initial_state)
        
        return {
            "status": "success",
            "decision": {
                "requires_ota": final_state["requires_ota"],
                "ota_version": final_state["ota_version"],
                "requires_work_order": final_state["requires_work_order"],
                "work_order_cost": final_state["work_order_cost"],
                "resolved": final_state["resolved"]
            }
        }
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))

if __name__ == "__main__":
    port = int(os.getenv("AGENT_PORT", "8001"))
    uvicorn.run(app, host="0.0.0.0", port=port)
