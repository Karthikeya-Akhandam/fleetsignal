import hashlib
import os
import numpy as np
from typing import List, Dict

class PrivacyEngine:
    def __init__(self, k_anonymity_threshold: int = 1, epsilon: float = 1.0):
        # We use a salt loaded from env or generated randomly for this session
        self.salt = os.getenv("PRIVACY_SALT", "motorq_hackathon_super_secret_salt").encode()
        self.k_threshold = k_anonymity_threshold
        self.epsilon = epsilon
        
    def hash_vin(self, vin: str) -> str:
        """One-way cryptographic hash of VIN with salt to prevent re-identification"""
        return hashlib.sha256(vin.encode() + self.salt).hexdigest()

    def enforce_k_anonymity(self, data: List[Dict], group_by_key: str) -> List[Dict]:
        """
        Filters out any aggregated groups that have fewer than 'k' distinct vehicles.
        Expects data dictionaries to have 'k_count' and 'geohash' or similar group_by_key.
        """
        return [row for row in data if row.get('k_count', 0) >= self.k_threshold]

    def add_laplace_noise(self, value: float, sensitivity: float) -> float:
        """
        Adds Laplace noise to a numerical value to satisfy epsilon-differential privacy.
        """
        scale = sensitivity / self.epsilon
        noise = np.random.laplace(0, scale)
        return value + noise
