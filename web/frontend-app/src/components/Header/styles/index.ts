import { AppBar, styled, Typography } from '@mui/material'

export const LayoutGrid = styled(AppBar)({
  'display': 'flex',
  'padding': '12px 60px',
  'backgroundColor': '#484848',
  'alignItems': 'center',
  'boxShadow': 'none',
  '@media only screen and (max-width: 600px)': {
    alignItems: 'center',
  },
})

export const CustomTypography = styled(Typography)({
  color: 'black',
  fontWeight: 'bold',
  fontSize: '20px',
  textTransform: 'uppercase'
})