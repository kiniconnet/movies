import React, { useState } from 'react';
import toast from 'react-hot-toast';
import { Link, useNavigate, useOutletContext } from 'react-router-dom';
import { useMutation } from '@tanstack/react-query';
import { BASE_URL } from '../main';

const Login = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const { setJwtToken } = useOutletContext();
  const navigate = useNavigate();

   
  // Define the mutation function for login
  const { mutate: login, isLoading } = useMutation({
    mutationFn: async (payload) => {
      const response = await fetch(`${BASE_URL}/authenticate`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(payload),
      });

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      const data = await response.json();
      console.log("Server response data:", data);
      return data;
    },

    onSuccess: (data) => {
      console.log("Full server response:", data);

      // Trying to access the token using different possible keys
      const accessToken = data.data.access_token;

      if (!accessToken || accessToken.trim() === "") {
        toast.error("Login failed: No valid access token received from the server.");
        return;
      }
    
      console.log("Access token:", accessToken);

      setJwtToken(data.data.access_token); // Store the JWT token
      toast.success('Logged in successfully');
      navigate('/'); // Redirect to the home page
    },
    onError: (error) => {
      console.error('Error during login:', error);
      toast.error(error.message || 'An error occurred. Please try again.');
    },
  });

  // Handle form submission
  const handleSubmit = (event) => {
    event.preventDefault();

    const payload = {
      email,
      password,
    };

    console.log('Sending login request with payload:', payload);

    // Trigger the mutation
    login(payload);
  };

  return (
    <div className="min-h-96 flex items-center justify-center bg-gray-50">
      <div className="w-full max-w-md p-8 space-y-6 bg-oxfordBlue rounded-lg shadow-md">
        <h2 className="text-2xl font-bold text-center text-gray-800">Login</h2>
        <form onSubmit={handleSubmit} className="space-y-4">
          {/* Email Input */}
          <div>
            <label htmlFor="email" className="block text-sm font-medium text-gray-700">
              Email:
            </label>
            <input
              type="email"
              id="email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              required
              aria-label="Email address"
              className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-oxfordBlue focus:border-oxfordBlue sm:text-sm text-black"
            />
          </div>

          {/* Password Input */}
          <div>
            <label htmlFor="password" className="block text-sm font-medium text-gray-700">
              Password:
            </label>
            <input
              type="password"
              id="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
              minLength={8}
              aria-label="Password"
              className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-oxfordBlue focus:border-oxfordBlue sm:text-sm text-black"
            />
          </div>

          {/* Submit Button */}
          <button
            type="submit"
            aria-label="Login"
            disabled={isLoading} // Disable the button while the request is in progress
            className={`w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white ${
              isLoading ? 'bg-gray-400' : 'bg-indigo-600 hover:bg-indigo-700'
            } focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500`}
          >
            {isLoading ? 'Logging in...' : 'Login'}
          </button>
        </form>

        <p className='text-orangeWeb'>Don't have an account? <Link to="/signup" className='text-white underline'>Sign Up</Link></p>
      </div>
    </div>
  );
};

export default Login;