import React, { useState, useEffect } from "react";
import {
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Paper,
  Typography,
} from "@mui/material";
import { ReactElement } from "react";
import axios from "axios";

type PhoneNumbers = {
  country: string;
  country_code: number;
  phone_number: string;
  state: boolean;
};

export const BasicTable = (): ReactElement => {
  const [data, setData] = useState<PhoneNumbers[]>([]);
  const [filterState, setFilterState] = useState(false);

  useEffect(() => {
    axios
      .get("http://localhost:666/phones")
      .then(function (response) {
        // handle success
        setData(response.data);
      })
      .catch(function (error) {
        console.log(error);
      });
  }, []);

  return (
    <div style={{ marginTop: "100px" }}>
      <Typography
        variant="h1"
        style={{ fontWeight: "bold", marginBottom: "20px" }}
      >
        Phone Numbers
      </Typography>
      <TableContainer component={Paper}>
        <Table sx={{ minWidth: 650 }} aria-label="simple table">
          <TableHead style={{ backgroundColor: "#707070" }}>
            <TableRow>
              <TableCell style={{ fontWeight: "bold" }}>Country</TableCell>
              <TableCell style={{ fontWeight: "bold" }} align="right">
                State
              </TableCell>
              <TableCell style={{ fontWeight: "bold" }} align="right">
                Country Code
              </TableCell>
              <TableCell style={{ fontWeight: "bold" }} align="right">
                Phone Number
              </TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {data.map((row, id) => (
              <TableRow
                key={id}
                sx={{ "&:last-child td, &:last-child th": { border: 0 } }}
              >
                <TableCell align="right">{row.country}</TableCell>
                <TableCell align="right">{row.state.toString()}</TableCell>
                <TableCell align="right">{row.country_code}</TableCell>
                <TableCell align="right">{row.phone_number}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </div>
  );
};
