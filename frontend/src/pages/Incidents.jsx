import React, { useEffect } from 'react';
import { Card } from '../components/Card';
import { Button } from '../components/Button';
import { useStore } from '../store/useStore';
import { fleetApi } from '../api/client';
import { AlertCircle, Wrench, ChevronRight } from 'lucide-react';
import { IncidentModal } from '../components/IncidentModal';

export function Incidents() {
  const { incidents, setIncidents, selectIncident } = useStore();

  useEffect(() => {
    const fetchIncidents = () => {
      fleetApi.get('/incidents')
        .then(data => {
          if (data) {
            setIncidents(data);
          }
        })
        .catch(err => console.error("Failed to fetch incidents:", err));
    };

    fetchIncidents();
    const interval = setInterval(fetchIncidents, 5000);
    return () => clearInterval(interval);
  }, [setIncidents]);

  return (
    <div className="animate-fade-in" style={{ display: 'flex', flexDirection: 'column', gap: '24px' }}>
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
        <h2>Incident Management</h2>
        <Button variant="secondary">Filter Incidents</Button>
      </div>

      <div style={{ display: 'flex', flexDirection: 'column', gap: '16px' }}>
        {incidents.map((incident) => (
          <Card key={incident.incident_id} className="incident-row" style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', cursor: 'pointer' }}>
            <div style={{ display: 'flex', alignItems: 'center', gap: '20px' }}>
              <div style={{ 
                padding: '12px', 
                background: incident.severity === 'CRITICAL' ? 'rgba(239, 68, 68, 0.1)' : incident.severity === 'HIGH' ? 'rgba(245, 158, 11, 0.1)' : 'rgba(59, 130, 246, 0.1)', 
                color: incident.severity === 'CRITICAL' ? 'var(--accent-danger)' : incident.severity === 'HIGH' ? 'var(--accent-warning)' : 'var(--brand-primary)',
                borderRadius: '50%' 
              }}>
                <AlertCircle size={24} />
              </div>
              
              <div>
                <h4 style={{ margin: '0 0 4px 0', display: 'flex', alignItems: 'center', gap: '8px' }}>
                  {incident.vin}
                  <span style={{ fontSize: '12px', padding: '2px 8px', background: 'var(--bg-secondary)', borderRadius: '12px', border: '1px solid var(--border-subtle)' }}>
                    {incident.fault_code}
                  </span>
                </h4>
                <p style={{ margin: 0, color: 'var(--text-muted)', fontSize: '14px' }}>{incident.description}</p>
              </div>
            </div>

            <div style={{ display: 'flex', alignItems: 'center', gap: '16px' }}>
              <span style={{ color: 'var(--text-muted)', fontSize: '13px' }}>
                {new Date(incident.timestamp).toLocaleTimeString([], {hour: '2-digit', minute:'2-digit'})}
              </span>
              <Button onClick={() => selectIncident(incident)} style={{ display: 'flex', alignItems: 'center', gap: '8px' }}>
                <Wrench size={16} /> Resolve
              </Button>
            </div>
          </Card>
        ))}
        {incidents.length === 0 && (
          <div style={{ textAlign: 'center', padding: '40px', color: 'var(--text-muted)' }}>
            No active incidents found.
          </div>
        )}
      </div>
      
      <IncidentModal />
    </div>
  );
}
