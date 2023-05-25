import React, { useState, useEffect } from "react";
import { DataGrid } from '@mui/x-data-grid';


const columns = [
  { field: 'id', headerName: 'ID', width: 70 ,
      renderCell: (params) => {
            return <a href={`/customers/${params.id}`}>{params.id}</a>;
      },
  },
  { field: 'first_name', headerName: 'First name', width: 130 },
  { field: 'company', headerName: 'Company', width: 130 },
  { field: 'gender', headerName: 'Gender', width: 130 },
  { field: 'city', headerName: 'City', width: 130 }
];



function DataTable({ isLoading, setIsLoading }) {
  let [data, setData] = useState([])

  useEffect(() => {
    setIsLoading(true)
    fetch('http://localhost:9999/api/customers')
      .then(res => res.json())
      .then(result => setData(result.data))
      .catch(err => console.error(err))
      .finally(()=> setIsLoading(false));
  }, []);

  return (
      !isLoading && <div>
        <center><h3>Customers List</h3></center>
            <div style={{ height: 400, width: '100%' }}>
                <DataGrid
                  rows={data}
                  columns={columns}
                  initialState={{
                    pagination: {
                      paginationModel: { page: 0, pageSize: 5 },
                    },
                  }}
                  pageSizeOptions={[5, 10]}
                />
            </div>
      </div>
    );

}

export default DataTable;
