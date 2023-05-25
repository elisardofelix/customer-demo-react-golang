import React from 'react';

const HoverLoader = ({isLoading}) => {
  return (
    <div>
      {isLoading && (
        <div
          className="hover-loader"
          style={{
            display: 'flex',
            justifyContent: 'center',
            alignItems: 'center',
            height: '100vh',
          }}
        >
          <div
            className="spinner"
            style={{
              border: '4px solid #f3f3f3',
              borderTop: '4px solid #3498db',
              borderRadius: '50%',
              width: '40px',
              height: '40px',
              animation: 'spin 2s linear infinite',
            }}
          ></div>
        </div>
      )}
    </div>
  );
};

export default HoverLoader;