import React, { useState, useEffect } from 'react';
import { Card } from '../components/Card';
import { hfApi } from '../api/client';
import { BarChart, Bar, XAxis, YAxis, CartesianGrid, Tooltip, ResponsiveContainer } from 'recharts';
import { TrendingUp, Activity, Battery, Thermometer } from 'lucide-react';

const mockAlphaData = [
  { feature: 'Battery Deg', importance: 0.85 },
  { feature: 'Motor Temp', importance: 0.62 },
  { feature: 'Brake Wear', importance: 0.55 },
  { feature: 'HVAC Load', importance: 0.41 },
  { feature: 'Tire Press', importance: 0.28 }
];

export function Analytics() {
  const [factors, setFactors] = useState([]);

  useEffect(() => {
    // Attempt to fetch from Python HF API, fallback to mock data
    hfApi.get('/alpha-factors')
      .then(res => setFactors(res.factors || mockAlphaData))
      .catch(() => setFactors(mockAlphaData));
  }, []);

  return (
    <div className="animate-fade-in" style={{ display: 'flex', flexDirection: 'column', gap: '24px' }}>
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
        <h2>Alpha Analytics (Privacy-Preserved)</h2>
        <span style={{ color: 'var(--accent-purple)', fontWeight: 600, display: 'flex', alignItems: 'center', gap: '8px' }}>
          <TrendingUp size={20} /> k-Anonymity Enabled
        </span>
      </div>

      <div style={{ display: 'grid', gridTemplateColumns: 'repeat(auto-fit, minmax(300px, 1fr))', gap: '24px' }}>
        <Card style={{ height: '400px', display: 'flex', flexDirection: 'column' }}>
          <h3 style={{ marginBottom: '24px' }}>Predictive Feature Importance</h3>
          <div style={{ flex: 1, width: '100%', minHeight: 0 }}>
            <ResponsiveContainer width="100%" height="100%">
              <BarChart data={factors} layout="vertical" margin={{ top: 0, right: 0, left: 20, bottom: 0 }}>
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
              Vehicles operating in regions with <strong style={{ color: 'var(--brand-primary)' }}>high ambient temperature variance</strong> show a 34% faster battery degradation rate.
            </p>
          </Card>
          
          <Card>
            <h4 style={{ color: 'var(--text-muted)', marginBottom: '16px', display: 'flex', alignItems: 'center', gap: '8px' }}>
              <Thermometer size={20} color="var(--accent-danger)" /> Anomaly Detection
            </h4>
            <p style={{ fontSize: '18px', lineHeight: 1.6, margin: 0 }}>
              Motor temperature spikes detected in <strong style={{ color: 'var(--accent-danger)' }}>Model Y Batch 4</strong> during rapid acceleration events.
            </p>
          </Card>
        </div>
      </div>
    </div>
  );
}
