import React, { useState } from 'react'
import Header from './Header'
import Sidebar from './Sidebar'
import { Outlet } from 'react-router-dom'
import { Toaster } from 'react-hot-toast'


const AppLayout = () => {
  const [jwtToken, setJwtToken] = useState("");
  


  return (
    <div className='w-11/12 mx-auto'>
        <Header jwtToken={jwtToken} setJwtToken={setJwtToken}/>
        <Sidebar jwtToken={jwtToken} setJwtToken={setJwtToken}/>
        <main>
              <Toaster position='top-center'/>
             <Outlet  context={{
              jwtToken, 
              setJwtToken,
             }}/>
        </main>
    </div>
  )
}

export default AppLayout