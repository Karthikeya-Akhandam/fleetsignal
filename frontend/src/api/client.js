import axios from 'axios';

// The Go Fleet API for high-throughput reads/writes
export const fleetApi = axios.create({
  baseURL: import.meta.env.VITE_FLEET_API_URL || 'http://localhost:8080/api/v1',
  headers: {
    'Content-Type': 'application/json',
  },
});

// The Python HF API for complex analytics and privacy-preserved ML
export const hfApi = axios.create({
  baseURL: import.meta.env.VITE_HF_API_URL || 'http://localhost:8000/api/v1',
  headers: {
    'Content-Type': 'application/json',
  },
});

// The Python Fleet Ops Agent API for LLM decisions
export const agentApi = axios.create({
  baseURL: import.meta.env.VITE_AGENT_API_URL || 'http://localhost:8001/api/v1/agent',
  headers: {
    'Content-Type': 'application/json',
  },
});

// Generic error handler
const handleApiError = (error) => {
  console.error('API Error:', error.response?.data || error.message);
  throw error;
};

// Response interceptors
fleetApi.interceptors.response.use((res) => res.data, handleApiError);
hfApi.interceptors.response.use((res) => res.data, handleApiError);
agentApi.interceptors.response.use((res) => res.data, handleApiError);
