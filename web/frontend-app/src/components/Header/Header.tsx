import React, { ReactElement } from 'react'
import { CustomTypography, LayoutGrid } from './styles'

export const Header = (): ReactElement => {
  return (
    <LayoutGrid>      
      <CustomTypography >Jumia Challenger</CustomTypography>
    </LayoutGrid>
  )
}
