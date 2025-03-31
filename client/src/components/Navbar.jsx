import React from 'react'
import { Link } from 'react-router-dom'

const Navbar = () => {
  return (
    <div className='container'>
        <div>
            {/* for logo */}
            <div>
                <Link to="/">Trusteegain</Link>
            </div>

            {/* for navbar list */}
            <div>
                <ul>
                    <li><Link to="/">Home</Link></li>
                    <li><Link to="/all-movies">All Movies</Link></li>
                    <li><Link to="/login">Login</Link></li>
                </ul>
            </div>
        </div>        
    </div>
  )
}

export default Navbar