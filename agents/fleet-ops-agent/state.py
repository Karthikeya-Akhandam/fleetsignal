from typing import TypedDict, Annotated, Sequence
import operator
from langchain_core.messages import BaseMessage

class AgentState(TypedDict):
    # Chat history
    messages: Annotated[Sequence[BaseMessage], operator.add]
    
    # Context state
    incident_id: str
    vin: str
    fault_code: str
    severity: str
    description: str
    
    # Agent decisions
    requires_ota: bool
    requires_work_order: bool
    ota_version: str
    work_order_cost: float
    work_order_notes: str
    
    # Completion flag
    resolved: bool
