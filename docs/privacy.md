# Privacy Engine

The Privacy Engine bridges the gap between high-frequency automotive telemetry and strict privacy laws, proving that actionable ML insights can be generated without compromising driver confidentiality.

## Implementation Details

Our Python-based `hf-api` utilises the **FastAPI** framework and leverages standard data-science libraries (Numpy/Pandas) to manipulate streams mathematically before JSON serialization.

### Differential Privacy (Laplace Noise)
Differential privacy provides a mathematical guarantee that the output of a query does not significantly change whether any single individual's data is included or not.
- We add statistical **Laplace noise** to all raw numerical aggregates (like Grid Stress or Freight Tonnage values).
- The noise prevents direct reverse-engineering of the true values while preserving the macroscopic statistical trends needed by fleet dispatchers.

### $k$-Anonymity
A dataset is said to have $k$-anonymity if the information for each person contained in the release cannot be distinguished from at least $k-1$ individuals whose information also appears in the release.
- For our geospatial analysis, if a specific region (geohash) has fewer than $k=5$ unique vehicles transmitting telemetry, the query output for that region is **completely suppressed**.
- This handles edge cases where Laplace Noise alone might be insufficient (e.g., if only 1 vehicle exists in a rural geohash).

## Threat Model
| Threat | Mitigation Control |
|--------|--------------------|
| Database Extraction | $k$-anonymisation prevents querying granular identities. |
| Time-Series Sniffing | Differential privacy masks time-based activity logs of a specific user. |
| Application Layer Attacks | Endpoints strictly return aggregated arrays instead of individual raw event rows. |
