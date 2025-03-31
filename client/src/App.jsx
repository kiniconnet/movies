import React from 'react'
import Navbar from './components/Navbar'
import { createBrowserRouter, Route, RouterProvider, Routes } from 'react-router-dom'
import Home from './pages/Home'
import AllMovies from './pages/AllMovies'
import Login from './pages/Login'
import AppLayout from './ui/AppLayout'
import SignUp from './pages/Signup'

const router = createBrowserRouter([
  {
    element: <AppLayout />,
    children: [
      {
        path: '/',
        exact: true,
        element: <Home />
      },
      {
        path: '/all-movies',
        element: <AllMovies />
      },
      {
        path: '/login',
        element: <Login />
      },
      {
        path: '/signup',
        element: <SignUp />
      },
    ],
  },
])

const App = () => {
  return  (
  <RouterProvider router={router}/>
  )
}

export default App