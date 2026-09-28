import pytest
from privacy import anonymize_vin, apply_differential_privacy

def test_anonymize_vin():
    vin = "WBA0000000001"
    anon1 = anonymize_vin(vin)
    anon2 = anonymize_vin(vin)
    
    # Must be deterministic for the same VIN
    assert anon1 == anon2
    
    # Must be 64 chars long (SHA-256 hex digest)
    assert len(anon1) == 64
    assert anon1 != vin

def test_differential_privacy():
    value = 100.0
    epsilon = 0.5
    
    dp_val1 = apply_differential_privacy(value, epsilon)
    dp_val2 = apply_differential_privacy(value, epsilon)
    
    # Noise should be applied, highly unlikely to equal exact input
    assert dp_val1 != value
    
    # It should not be completely random, but clustered around the true value (usually within a range)
    # Epsilon 0.5 means scale of 2.0 (1/0.5), so mostly within +/- 10
    assert 50.0 < dp_val1 < 150.0
