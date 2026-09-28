import React from 'react';
import { Button } from './Button';
import { useStore } from '../store/useStore';
import { X, Cpu, Wrench } from 'lucide-react';
import { agentApi } from '../api/client';

export function IncidentModal() {
  const { selectedIncident, selectIncident, setAnalyzing, isAnalyzing, setAgentDecision, agentDecision } = useStore();

  if (!selectedIncident) return null;

  const handleAskAgent = async () => {
    setAnalyzing(true);
    setAgentDecision(null);
    try {
      // API call to the Python Fleet Ops Agent
      const res = await agentApi.post('/evaluate-incident', selectedIncident);
      setAgentDecision(res.decision);
    } catch (err) {
      console.error(err);
      // Fallback mock decision if backend isn't running
      setTimeout(() => {
        setAgentDecision({
          requires_ota: selectedIncident.severity !== 'CRITICAL',
          ota_version: 'v2.4.1-patch',
          requires_work_order: selectedIncident.severity === 'CRITICAL',
          work_order_cost: 450.00,
          resolved: selectedIncident.severity !== 'CRITICAL'
        });
      }, 1500);
    } finally {
      setTimeout(() => setAnalyzing(false), 1500); // Simulate thought process delay
    }
  };

  return (
    <div className="modal-overlay" onClick={() => selectIncident(null)}>
      <div className="modal-content" onClick={e => e.stopPropagation()}>
        <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: '24px' }}>
          <h3 style={{ margin: 0, fontSize: '20px' }}>Incident Copilot</h3>
          <X style={{ cursor: 'pointer', color: 'var(--text-muted)' }} onClick={() => selectIncident(null)} />
        </div>

        <div style={{ background: 'var(--bg-primary)', padding: '16px', borderRadius: '12px', marginBottom: '24px' }}>
          <p style={{ margin: '0 0 8px 0', fontSize: '14px', color: 'var(--text-muted)' }}>Vehicle: <strong>{selectedIncident.vin}</strong></p>
          <p style={{ margin: '0 0 8px 0', fontSize: '14px', color: 'var(--text-muted)' }}>Fault Code: <strong style={{ color: 'var(--text-primary)' }}>{selectedIncident.fault_code}</strong></p>
          <p style={{ margin: 0, fontSize: '14px', color: 'var(--text-muted)' }}>Description: <strong style={{ color: 'var(--text-primary)' }}>{selectedIncident.description}</strong></p>
        </div>

        {!agentDecision && !isAnalyzing && (
          <Button variant="primary" style={{ width: '100%', gap: '8px' }} onClick={handleAskAgent}>
            <Cpu size={18} /> Ask AI Copilot for Mitigation
          </Button>
        )}

        {isAnalyzing && (
          <div style={{ textAlign: 'center', padding: '24px', color: 'var(--brand-primary)' }}>
            <Cpu size={32} className="animate-pulse" style={{ animation: 'fadeIn 1s infinite alternate' }} />
            <p style={{ marginTop: '16px' }}>LangGraph Agent is analyzing...</p>
          </div>
        )}

        {agentDecision && (
          <div className="animate-fade-in" style={{ borderTop: '1px solid var(--border-subtle)', paddingTop: '24px' }}>
            <h4 style={{ marginBottom: '16px', display: 'flex', alignItems: 'center', gap: '8px', color: 'var(--accent-success)' }}>
              <CheckCircle size={18} /> AI Recommended Action
            </h4>
            
            {agentDecision.requires_ota && (
              <div style={{ padding: '12px', background: 'rgba(16, 185, 129, 0.1)', borderRadius: '8px', marginBottom: '12px' }}>
                <p style={{ margin: 0, fontSize: '14px', color: 'var(--accent-success)' }}>
                  <strong>OTA Update Ready:</strong> {agentDecision.ota_version}
                </p>
              </div>
            )}

            {agentDecision.requires_work_order && (
              <div style={{ padding: '12px', background: 'rgba(245, 158, 11, 0.1)', borderRadius: '8px', marginBottom: '12px' }}>
                <p style={{ margin: 0, fontSize: '14px', color: 'var(--accent-warning)' }}>
                  <strong>Work Order Required:</strong> Estimated ${agentDecision.work_order_cost}
                </p>
              </div>
            )}

            <div style={{ display: 'flex', gap: '12px', marginTop: '24px' }}>
              <Button variant="secondary" style={{ flex: 1 }} onClick={() => selectIncident(null)}>Dismiss</Button>
              <Button variant="primary" style={{ flex: 2 }}>Execute Actions</Button>
            </div>
          </div>
        )}
      </div>
    </div>
  );
}

// Dummy icon for use in this file until lucide is imported properly
function CheckCircle(props) {
  return <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" {...props}><path d="M22 11.08V12a10 10 0 1 1-5.93-9.14"></path><polyline points="22 4 12 14.01 9 11.01"></polyline></svg>;
}
