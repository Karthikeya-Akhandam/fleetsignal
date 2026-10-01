import React, { useState, useEffect } from 'react';
import { Card } from '../components/Card';
import { hfApi, fleetApi } from '../api/client';
import { BarChart, Bar, XAxis, YAxis, CartesianGrid, Tooltip, ResponsiveContainer } from 'recharts';
import { TrendingUp, Activity, Battery, Thermometer } from 'lucide-react';

export function Analytics() {
  const [freightData, setFreightData] = useState([]);
  const [gridStress, setGridStress] = useState([]);
  const [anomaly, setAnomaly] = useState(null);
  
  useEffect(() => {
    // Interceptor returns the parsed JSON directly!
    hfApi.get('/alpha/freight-tonnage')
      .then(json => {
        if (json && json.data) {
          const formatted = json.data
            .sort((a, b) => b.factor_value - a.factor_value)
            .slice(0, 5)
            .map(d => ({
              feature: d.geohash,
              importance: d.factor_value
            }));
          setFreightData(formatted);
        }
      })
      .catch(err => console.error('freight error:', err));
      
    hfApi.get('/alpha/grid-stress')
      .then(json => {
        if (json && json.data) {
          const sorted = json.data.sort((a, b) => b.factor_value - a.factor_value);
          setGridStress(sorted);
        }
      })
      .catch(err => console.error('grid stress error:', err));

    fleetApi.get('/incidents')
      .then(data => {
         if (data && data.length > 0) {
            setAnomaly(data[0]); // the most recent one
         }
      })
      .catch(err => console.error('incident error:', err));
  }, []);

  const displayData = freightData;

  return (
    <div className="animate-fade-in" style={{ display: 'flex', flexDirection: 'column', gap: '24px' }}>
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
        <h2>Alpha Analytics (Privacy-Preserved)</h2>
        <span style={{ color: 'var(--accent-purple)', fontWeight: 600, display: 'flex', alignItems: 'center', gap: '8px' }}>
          <TrendingUp size={20} /> k-Anonymity Enabled
        </span>
      </div>

      <div style={{ display: 'grid', gridTemplateColumns: 'repeat(auto-fit, minmax(300px, 1fr))', gap: '24px' }}>
        <Card style={{ display: 'flex', flexDirection: 'column' }}>
          <h3 style={{ marginBottom: '24px' }}>Predictive Feature Importance</h3>
          <div style={{ width: '100%', height: '300px' }}>
            <ResponsiveContainer width="100%" height="100%">
              <BarChart data={displayData} layout="vertical" margin={{ top: 0, right: 0, left: 20, bottom: 0 }}>
                <CartesianGrid strokeDasharray="3 3" stroke="var(--border-subtle)" horizontal={false} />
                <XAxis type="number" stroke="var(--text-muted)" />
                <YAxis dataKey="feature" type="category" stroke="var(--text-primary)" width={100} />
                <Tooltip 
                  cursor={{ fill: 'rgba(255,255,255,0.05)' }}
                  contentStyle={{ background: 'var(--bg-elevated)', border: '1px solid var(--border-subtle)', borderRadius: '8px' }} 
                />
                <Bar dataKey="importance" fill="var(--accent-purple)" radius={[0, 4, 4, 0]} />
              </BarChart>
            </ResponsiveContainer>
          </div>
        </Card>

        <div style={{ display: 'flex', flexDirection: 'column', gap: '24px' }}>
          <Card>
            <h4 style={{ color: 'var(--text-muted)', marginBottom: '16px', display: 'flex', alignItems: 'center', gap: '8px' }}>
              <Battery size={20} color="var(--brand-primary)" /> Top Insight
            </h4>
            <p style={{ fontSize: '18px', lineHeight: 1.6, margin: 0 }}>
              {gridStress.length > 0 ? (
                <>High stress detected in <strong style={{ color: 'var(--brand-primary)' }}>{gridStress[0].geohash}</strong> with stress index {gridStress[0].factor_value.toFixed(2)}.</>
              ) : (
                <>Vehicles operating in regions with <strong style={{ color: 'var(--brand-primary)' }}>high ambient temperature variance</strong> show a 34% faster battery degradation rate.</>
              )}
            </p>
          </Card>
          
          <Card>
            <h4 style={{ color: 'var(--text-muted)', marginBottom: '16px', display: 'flex', alignItems: 'center', gap: '8px' }}>
              <Thermometer size={20} color="var(--accent-danger)" /> Anomaly Detection
            </h4>
            <p style={{ fontSize: '18px', lineHeight: 1.6, margin: 0 }}>
              {anomaly ? (
                 <>Incident <strong style={{ color: 'var(--accent-danger)' }}>{anomaly.fault_code}</strong> detected in Vehicle {anomaly.vin}: {anomaly.description}.</>
              ) : (
                 <>Motor temperature spikes detected in <strong style={{ color: 'var(--accent-danger)' }}>Model Y Batch 4</strong> during rapid acceleration events.</>
              )}
            </p>
          </Card>
        </div>
      </div>
    </div>
  );
}
