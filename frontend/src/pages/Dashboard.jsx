import React from 'react';
import { Card } from '../components/Card';
import { Activity, AlertTriangle, CheckCircle, Zap } from 'lucide-react';
import { AreaChart, Area, XAxis, YAxis, CartesianGrid, Tooltip, ResponsiveContainer } from 'recharts';

const mockData = [
  { time: '10:00', incidents: 2, updates: 5 },
  { time: '11:00', incidents: 5, updates: 8 },
  { time: '12:00', incidents: 3, updates: 12 },
  { time: '13:00', incidents: 8, updates: 15 },
  { time: '14:00', incidents: 4, updates: 9 },
  { time: '15:00', incidents: 1, updates: 4 },
];

export function Dashboard() {
  return (
    <div className="animate-fade-in" style={{ display: 'flex', flexDirection: 'column', gap: '24px' }}>
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
        <h2>Dashboard Overview</h2>
        <span style={{ color: 'var(--text-muted)' }}>Real-time Fleet Status</span>
      </div>

      <div style={{ display: 'grid', gridTemplateColumns: 'repeat(auto-fit, minmax(240px, 1fr))', gap: '24px' }}>
        <Card>
          <div style={{ display: 'flex', alignItems: 'center', gap: '16px' }}>
            <div style={{ padding: '12px', background: 'rgba(59, 130, 246, 0.1)', borderRadius: '12px', color: 'var(--brand-primary)' }}>
              <Activity size={24} />
            </div>
            <div>
              <p style={{ color: 'var(--text-muted)', fontSize: '14px' }}>Active Vehicles</p>
              <h3 style={{ fontSize: '24px', margin: 0 }}>1,248</h3>
            </div>
          </div>
        </Card>
        
        <Card>
          <div style={{ display: 'flex', alignItems: 'center', gap: '16px' }}>
            <div style={{ padding: '12px', background: 'rgba(239, 68, 68, 0.1)', borderRadius: '12px', color: 'var(--accent-danger)' }}>
              <AlertTriangle size={24} />
            </div>
            <div>
              <p style={{ color: 'var(--text-muted)', fontSize: '14px' }}>Critical Incidents</p>
              <h3 style={{ fontSize: '24px', margin: 0 }}>23</h3>
            </div>
          </div>
        </Card>

        <Card>
          <div style={{ display: 'flex', alignItems: 'center', gap: '16px' }}>
            <div style={{ padding: '12px', background: 'rgba(16, 185, 129, 0.1)', borderRadius: '12px', color: 'var(--accent-success)' }}>
              <CheckCircle size={24} />
            </div>
            <div>
              <p style={{ color: 'var(--text-muted)', fontSize: '14px' }}>Resolved via OTA</p>
              <h3 style={{ fontSize: '24px', margin: 0 }}>156</h3>
            </div>
          </div>
        </Card>
      </div>

      <Card className="chart-container" style={{ height: '400px', display: 'flex', flexDirection: 'column' }}>
        <h3 style={{ marginBottom: '24px' }}>Incident & Update Activity</h3>
        <div style={{ flex: 1, width: '100%', minHeight: 0 }}>
          <ResponsiveContainer width="100%" height="100%">
            <AreaChart data={mockData}>
              <defs>
                <linearGradient id="colorIncidents" x1="0" y1="0" x2="0" y2="1">
                  <stop offset="5%" stopColor="var(--accent-danger)" stopOpacity={0.3}/>
                  <stop offset="95%" stopColor="var(--accent-danger)" stopOpacity={0}/>
                </linearGradient>
                <linearGradient id="colorUpdates" x1="0" y1="0" x2="0" y2="1">
                  <stop offset="5%" stopColor="var(--brand-primary)" stopOpacity={0.3}/>
                  <stop offset="95%" stopColor="var(--brand-primary)" stopOpacity={0}/>
                </linearGradient>
              </defs>
              <CartesianGrid strokeDasharray="3 3" stroke="var(--border-subtle)" vertical={false} />
              <XAxis dataKey="time" stroke="var(--text-muted)" />
              <YAxis stroke="var(--text-muted)" />
              <Tooltip 
                contentStyle={{ background: 'var(--bg-elevated)', border: '1px solid var(--border-subtle)', borderRadius: '8px' }}
                itemStyle={{ color: 'var(--text-primary)' }}
              />
              <Area type="monotone" dataKey="incidents" stroke="var(--accent-danger)" fillOpacity={1} fill="url(#colorIncidents)" />
              <Area type="monotone" dataKey="updates" stroke="var(--brand-primary)" fillOpacity={1} fill="url(#colorUpdates)" />
            </AreaChart>
          </ResponsiveContainer>
        </div>
      </Card>
    </div>
  );
}
