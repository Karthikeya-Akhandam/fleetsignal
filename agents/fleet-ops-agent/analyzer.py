import os
import psycopg2
from langchain_google_genai import ChatGoogleGenerativeAI
from langchain_core.messages import SystemMessage, HumanMessage
from state import AgentState

llm = ChatGoogleGenerativeAI(model="gemini-1.5-flash", temperature=0)

def get_pg_connection():
    return psycopg2.connect(
        host=os.getenv("POSTGRES_HOST", "localhost"),
        port=os.getenv("POSTGRES_PORT", "5432"),
        user=os.getenv("POSTGRES_USER", "motorq"),
        password=os.getenv("POSTGRES_PASSWORD", "motorq_password"),
        database=os.getenv("POSTGRES_DB", "fleetsignal")
    )

def analyze_incident(state: AgentState) -> AgentState:
    """
    Analyzes an incident to determine if it requires an OTA update or a physical work order.
    """
    # In a real scenario, we would use the fault embedding to query similar past incidents from pgvector
    # For this hackathon scope, we will use the LLM directly on the fault code and description
    
    prompt = f"""
    You are an expert Fleet Operations AI. Analyze this vehicle incident:
    VIN: {state['vin']}
    Fault Code: {state['fault_code']}
    Severity: {state['severity']}
    Description: {state['description']}
    
    Determine:
    1. Can this be fixed with a software Over-The-Air (OTA) update? (Yes/No)
    2. Does it require physical repair at a shop (Work Order)? (Yes/No)
    
    If OTA is required, suggest a version string like 'v2.4.1-patch'.
    If Work Order is required, estimate a cost in USD and provide notes.
    
    Respond in JSON format:
    {{"requires_ota": bool, "requires_work_order": bool, "ota_version": "string", "work_order_cost": float, "work_order_notes": "string"}}
    """
    
    response = llm.invoke([
        SystemMessage(content="You are a JSON-only response bot."),
        HumanMessage(content=prompt)
    ])
    
    import json
    try:
        # Strip markdown formatting if any
        raw_text = response.content.replace("```json", "").replace("```", "").strip()
        decision = json.loads(raw_text)
        
        return {
            **state,
            "requires_ota": decision.get("requires_ota", False),
            "requires_work_order": decision.get("requires_work_order", False),
            "ota_version": decision.get("ota_version", ""),
            "work_order_cost": decision.get("work_order_cost", 0.0),
            "work_order_notes": decision.get("work_order_notes", "")
        }
    except Exception as e:
        # Fallback
        return {**state, "requires_work_order": True, "work_order_notes": "Failed to parse AI decision."}
