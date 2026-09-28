import { create } from 'zustand';

export const useStore = create((set) => ({
  // State
  incidents: [],
  selectedIncident: null,
  isAnalyzing: false,
  agentDecision: null,

  // Actions
  setIncidents: (incidents) => set({ incidents }),
  
  selectIncident: (incident) => set({ 
    selectedIncident: incident,
    agentDecision: null // Reset decision when selecting new incident
  }),

  setAnalyzing: (status) => set({ isAnalyzing: status }),
  
  setAgentDecision: (decision) => set({ agentDecision: decision }),
}));
