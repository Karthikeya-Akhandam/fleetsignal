import React from 'react';
import { Outlet, Link } from 'react-router-dom';
import './components.css';

export function Layout() {
  return (
    <div style={{ display: 'flex', flexDirection: 'column', minHeight: '100vh' }}>
      <header className="glass-panel" style={{ margin: '16px', padding: '16px 24px', display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
        <h1 style={{ fontSize: '24px', fontWeight: 800, background: 'var(--gradient-brand)', WebkitBackgroundClip: 'text', WebkitTextFillColor: 'transparent' }}>
          FleetSignal
        </h1>
        <nav style={{ display: 'flex', gap: '20px' }}>
          <Link to="/" style={{ color: 'var(--text-primary)', textDecoration: 'none', fontWeight: 500 }}>Dashboard</Link>
          <Link to="/incidents" style={{ color: 'var(--text-primary)', textDecoration: 'none', fontWeight: 500 }}>Incidents</Link>
          <Link to="/analytics" style={{ color: 'var(--text-primary)', textDecoration: 'none', fontWeight: 500 }}>Analytics</Link>
        </nav>
      </header>
      
      <main style={{ flex: 1, padding: '0 24px 24px', maxWidth: '1400px', margin: '0 auto', width: '100%' }}>
        <Outlet />
      </main>
    </div>
  );
}
