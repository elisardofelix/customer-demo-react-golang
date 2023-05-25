import React from 'react';
import { render, act } from '@testing-library/react';
import { describe, it,expect, vi } from 'vitest'
import { MemoryRouter, Route, Routes } from 'react-router-dom';
import CustomerDetails from './CustomerDetails';


global.fetch = vi.fn();


function createFetchResponse(data) {
  return { json: () => new Promise((resolve) => resolve(data)) }
}

describe('CustomerDetails component', () => {
    it('renders customer details correctly', async () => {

      const customerDetails = 
       {data : {
        id: 1,
        first_name: "Laura",
        last_name: "Richards",
        email: "lrichards0@reverbnation.com",
        gender: "Female",
        company: "Meezzy",
        city: "Warner, NH",
        title: "Biostatistician III",
        created_at: "2023-01-30T21:53:49.182774Z",
        updated_at: "2023-01-30T21:53:49.182774Z"
        }};
    
    
      fetch.mockResolvedValue(createFetchResponse(customerDetails))


      let component;
        

      await act(async () => {
        component = render(
          <MemoryRouter initialEntries={['/customers/1']}>
                 <Routes>
              <Route path='/customers/:id' element={<CustomerDetails isLoading={false} setIsLoading={(state)=>{}} />} />
              </Routes>
        </MemoryRouter>
      );
    });
      
  
      // Verificar que se haya obtenido correctamente el componente renderizado
      expect(component.getByText("Laura")).toBeDefined();
      expect(component.getByText("Richards")).toBeDefined();
      expect(component.getByText("lrichards0@reverbnation.com")).toBeDefined();
      expect(component.getByText("Female")).toBeDefined();

    });
  });
  