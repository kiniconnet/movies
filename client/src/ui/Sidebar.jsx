import React from 'react'
import { NavLink } from 'react-router-dom'

const Sidebar = ({jwtToken}) => {
  return (
    <div>
        {jwtToken == "" 
        ?<p></p>
        :(<div className='bg-platinum blur-2 rounded-lg text-oxfordBlue p-2 mb-3'>
          <ul className='flex gap-4'>
            <li>
                <NavLink to="/">Home</NavLink>
            </li>
            <li>
                <NavLink to="/all-movies">All Movies</NavLink>
            </li>
        </ul>
        </div>)}
    </div>
  )
}

export default Sidebar