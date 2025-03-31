import React from 'react'
import toast from 'react-hot-toast';
import { Link, useNavigate } from 'react-router-dom'

const Header = ({jwtToken, setJwtToken}) => {
  
  const navigate = useNavigate();


  const LogOut = () => {
    setJwtToken("");
    toast.success("Logged Out Successfully")
    navigate("/login")
  }

  return (
        <div className='flex items-center w-10/12 mx-auto my-3 justify-between bg-oxfordBlue py-3 px-5 rounded-3xl shadow-lg'>
            {/* logo section */}
            <div>
                <Link to="/" className='text-orangeWeb'>TMH <span className='text-white'>|</span> <span className='underline underline-offset-8 text-platinum'>Movies</span></Link> 
            </div>

            {/*Authentication menu*/}
            <div>
              { jwtToken === ""
               ?<Link to="/login" className='bg-platinum py-1 px-3 rounded-lg'>Login</Link>
                : <Link to="#!" onClick={LogOut} className='bg-orangeWeb py-1 px-3 rounded-lg'>LogOut</Link> 
              }
            </div>
        </div>
  )
}

export default Header