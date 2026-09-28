from pydantic import BaseModel, Field
from typing import List, Optional
from datetime import datetime

class AnonymizedFactor(BaseModel):
    geohash: str = Field(..., description="Geohash precision 5 string")
    timestamp: datetime
    k_count: int = Field(..., description="Number of vehicles in this aggregation to satisfy k-anonymity")
    factor_value: float = Field(..., description="Calculated alpha factor metric with laplace noise applied")

class AlphaFactorResponse(BaseModel):
    factor_name: str
    description: str
    data: List[AnonymizedFactor]
    metadata: dict
