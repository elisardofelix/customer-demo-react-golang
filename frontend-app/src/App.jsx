import React, { useState } from 'react';
import DataTable from './components/DataTable';
import CustomerDetails from './components/CustomerDetails';
import { BrowserRouter as Router, Routes, Route, } from 'react-router-dom';
import './App.css';
import HoverLoader from './components/HoverLoader';

function App() {
  const [isLoading, setIsLoading] = useState(true);


  return (
    <div className="App">
      <HoverLoader isLoading={isLoading} />
      <header className="App-header" style={{display : (isLoading ? "none": "block")}}>
        <nav>
          <h1>Customers Demo</h1>
        </nav>
      </header>
      <main className="App-body">
        <Router>
            <Routes>
              <Route path='/' element={<DataTable isLoading={isLoading} setIsLoading={setIsLoading} />} />
              <Route path='/customers/:id' element={<CustomerDetails isLoading={isLoading} setIsLoading={setIsLoading}/>} />
            </Routes>
        </Router>
      </main>
    </div>
     
  );
}

export default App;
