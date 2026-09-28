import React from 'react';
import './components.css';

export function Button({ children, variant = 'primary', onClick, className = '', ...props }) {
  const baseClass = variant === 'primary' ? 'btn-primary' : 'btn-secondary';
  
  return (
    <button 
      className={`btn ${baseClass} ${className}`} 
      onClick={onClick}
      {...props}
    >
      {children}
    </button>
  );
}
