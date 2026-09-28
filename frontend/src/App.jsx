import React from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { Layout } from './components/Layout';

// Placeholder Pages
const Dashboard = () => <div className="animate-fade-in"><h2>Dashboard Overview</h2></div>;
const Incidents = () => <div className="animate-fade-in"><h2>Incident Management</h2></div>;
const Analytics = () => <div className="animate-fade-in"><h2>Alpha Analytics</h2></div>;

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Layout />}>
          <Route index element={<Dashboard />} />
          <Route path="incidents" element={<Incidents />} />
          <Route path="analytics" element={<Analytics />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
}

export default App;
