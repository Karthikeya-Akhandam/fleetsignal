from langgraph.graph import StateGraph, END
from state import AgentState
from analyzer import analyze_incident
from ota import execute_ota_mitigation
from work_order import create_work_order

def build_graph():
    workflow = StateGraph(AgentState)
    
    # Add nodes
    workflow.add_node("analyzer", analyze_incident)
    workflow.add_node("ota_mitigation", execute_ota_mitigation)
    workflow.add_node("work_order_creation", create_work_order)
    
    # Set entry point
    workflow.set_entry_point("analyzer")
    
    # Routing logic
    def route_mitigation(state: AgentState):
        if state.get("requires_ota"):
            return "ota_mitigation"
        if state.get("requires_work_order"):
            return "work_order_creation"
        return END

    workflow.add_conditional_edges(
        "analyzer",
        route_mitigation,
        {
            "ota_mitigation": "ota_mitigation",
            "work_order_creation": "work_order_creation",
            END: END
        }
    )
    
    # Even if OTA happens, it might still need a physical work order (e.g. software fix + sensor replace)
    def route_post_ota(state: AgentState):
        if state.get("requires_work_order"):
            return "work_order_creation"
        return END
        
    workflow.add_conditional_edges(
        "ota_mitigation",
        route_post_ota,
        {
            "work_order_creation": "work_order_creation",
            END: END
        }
    )
    
    # Work order is terminal
    workflow.add_edge("work_order_creation", END)
    
    return workflow.compile()

app = build_graph()
