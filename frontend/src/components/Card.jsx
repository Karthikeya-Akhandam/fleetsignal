import React from 'react';
import './components.css';

export function Card({ children, className = '' }) {
  return (
    <div className={`glass-card ${className}`}>
      {children}
    </div>
  );
}
