from state import AgentState
from analyzer import get_pg_connection

def create_work_order(state: AgentState) -> AgentState:
    """
    Creates a physical work order for a vehicle if physical repair is needed.
    """
    if not state.get("requires_work_order"):
        return state
        
    conn = get_pg_connection()
    try:
        cur = conn.cursor()
        
        # Insert Work Order
        query = """
            INSERT INTO work_orders (incident_id, vin, status, description, cost_estimate)
            VALUES (%s, %s, %s, %s, %s)
        """
        cur.execute(query, (
            state["incident_id"], 
            state["vin"], 
            "PENDING_DISPATCH", 
            state["work_order_notes"], 
            state["work_order_cost"]
        ))
        
        # We do NOT mark the incident as resolved because the work order is pending physical action.
        state["resolved"] = False
        
        conn.commit()
        cur.close()
    except Exception as e:
        print(f"Failed to create work order: {e}")
    finally:
        conn.close()
        
    return state
