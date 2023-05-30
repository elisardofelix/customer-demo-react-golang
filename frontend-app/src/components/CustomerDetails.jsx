
import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';



const CustomerDetails = ({ isLoading, setIsLoading }) => {
    const [customer, setCustomer] = useState([]);
    let { id } = useParams();

  useEffect(() => {
    setIsLoading(true)
    fetch(`http://localhost:9999/api/customers/${id}`)
      .then(res => res.json())
      .then(result => setCustomer(result.data))
      .finally(() => setIsLoading(false));
  }, [id]);

  return (
      !isLoading && <div>
          <center><h3 style={{ color:"black" }}>Customer Details</h3></center>
            <TableContainer component={Paper}>
              <Table aria-label="customer details table">
                <TableHead>
                  <TableRow>
                    <TableCell>ID</TableCell>
                    <TableCell align="center">First Name</TableCell>
                    <TableCell align="center">Last Name</TableCell>
                    <TableCell align="center">Email</TableCell>
                    <TableCell align="center">Gender</TableCell>
                    <TableCell align="center">Company</TableCell>
                    <TableCell align="center">City</TableCell>
                    <TableCell align="center">Title</TableCell>
                  </TableRow>
                </TableHead>
                <TableBody>
                  <TableRow>
                    <TableCell component="th" scope="row">
                      {customer.id}
                    </TableCell>
                    <TableCell align="center">{customer.first_name}</TableCell>
                    <TableCell align="center">{customer.last_name}</TableCell>
                    <TableCell align="center">{customer.email}</TableCell>
                    <TableCell align="center">{customer.gender}</TableCell>
                    <TableCell align="center">{customer.company}</TableCell>
                    <TableCell align="center">{customer.city}</TableCell>
                    <TableCell align="center">{customer.title}</TableCell>
                  </TableRow>
                </TableBody>
              </Table>
            </TableContainer>
      </div>
  );
};

export default CustomerDetails;
