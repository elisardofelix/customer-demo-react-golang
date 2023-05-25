import React from 'react';
import { describe, it, expect, vi } from 'vitest'
import { render, act } from '@testing-library/react';
import { MemoryRouter } from 'react-router-dom';
import DataTable from './DataTable';

global.fetch = vi.fn();


function createFetchResponse(data) {
  return { json: () => new Promise((resolve) => resolve(data)) }
}

describe('suite', () => {

  const data = 
    { 
      data :[{
      id: 1,
      first_name: 'Eber',
      company: 'LalaCom',
      gender: 'Male',
      city: 'Santo Domingo'
    },
    {
      id: 2,
      first_name: 'Sara',
      company: 'EliCorp',
      gender: 'Female',
      city: 'Doral'
    }
  ]};
    
  fetch.mockResolvedValue(createFetchResponse(data))

  let component

  it('renders DataTable component with data first_name', async () => {
    
    await act(async () => {
      component = render(
      <MemoryRouter>
        <DataTable isLoading={false} setIsLoading={(state)=>{}} />
      </MemoryRouter>
    );
    });

    

    data.data.forEach(datum => {
      expect(component.getByText(datum.first_name)).toBeDefined();
      expect(component.getByText(datum.company)).toBeDefined();
    });
  });

}) 

