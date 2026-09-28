from state import AgentState
from analyzer import get_pg_connection

def execute_ota_mitigation(state: AgentState) -> AgentState:
    """
    Simulates sending an OTA update to a vehicle.
    Records the OTA action in the database.
    """
    if not state.get("requires_ota"):
        return state
        
    conn = get_pg_connection()
    try:
        cur = conn.cursor()
        
        # Insert OTA record
        query = """
            INSERT INTO ota_updates (vin, version, status)
            VALUES (%s, %s, %s)
        """
        cur.execute(query, (state["vin"], state["ota_version"], "COMPLETED"))
        
        # Mark incident resolved if no physical repair needed
        if not state.get("requires_work_order"):
            resolve_query = """
                UPDATE incidents 
                SET resolved = true, resolution_notes = %s
                WHERE incident_id = %s
            """
            cur.execute(resolve_query, (f"Resolved via OTA update to {state['ota_version']}", state["incident_id"]))
            state["resolved"] = True
            
        conn.commit()
        cur.close()
    except Exception as e:
        print(f"Failed to execute OTA: {e}")
    finally:
        conn.close()
        
    return state
