from fastapi import APIRouter, HTTPException, Query
from datetime import datetime, timedelta
from schemas import AlphaFactorResponse
from alpha_factors import alpha_engine

router = APIRouter()

@router.get("/alpha/freight-tonnage", response_model=AlphaFactorResponse)
async def get_freight_tonnage(
    start_time: str = Query(default_factory=lambda: (datetime.utcnow() - timedelta(days=1)).strftime("%Y-%m-%d %H:%M:%S")),
    end_time: str = Query(default_factory=lambda: datetime.utcnow().strftime("%Y-%m-%d %H:%M:%S"))
):
    try:
        data = alpha_engine.get_freight_tonnage_index(start_time, end_time)
        return AlphaFactorResponse(
            factor_name="Freight Tonnage Index",
            description="Aggregated cargo weight by region. Correlates with supply chain volume and industrial output.",
            data=data,
            metadata={"start_time": start_time, "end_time": end_time, "privacy": "k=5, epsilon=1.0"}
        )
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))

@router.get("/alpha/grid-stress", response_model=AlphaFactorResponse)
async def get_grid_stress(
    start_time: str = Query(default_factory=lambda: (datetime.utcnow() - timedelta(days=1)).strftime("%Y-%m-%d %H:%M:%S")),
    end_time: str = Query(default_factory=lambda: datetime.utcnow().strftime("%Y-%m-%d %H:%M:%S"))
):
    try:
        data = alpha_engine.get_grid_stress_demand(start_time, end_time)
        return AlphaFactorResponse(
            factor_name="Grid-Stress MW Demand",
            description="Aggregated EV battery charging load by region. Correlates with localized grid stress.",
            data=data,
            metadata={"start_time": start_time, "end_time": end_time, "privacy": "k=5, epsilon=1.0"}
        )
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))
