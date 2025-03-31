import React, { useState } from 'react';
import toast from 'react-hot-toast';
import { useNavigate } from 'react-router-dom';
import { useMutation } from '@tanstack/react-query';
import { BASE_URL } from '../main';

const SignUp = () => {
  // State variables for form inputs
  const [firstName, setFirstName] = useState('');
  const [lastName, setLastName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const navigate = useNavigate();

  // Define the mutation function for signing up
  const { mutate: signUp, isLoading } = useMutation({
    mutationFn: async (payload) => {
      const response = await fetch(`${BASE_URL}/signup`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(payload),
      });

      if (!response.ok) {
        const errorText = await response.text();
        throw new Error(errorText || `HTTP error! status: ${response.status}`);
      }

      const data = await response.json();
      console.log("Server response data:", data);
      return data;
    },

    onSuccess: (data) => {
      console.log("User registered successfully:", data);

      // Assuming the server returns a success message or user data
      toast.success('Account created successfully');
      navigate('/login'); // Redirect to the login page after successful registration
    },
    onError: (error) => {
      console.error('Error during signup:', error);
      toast.error(error.message || 'An error occurred. Please try again.');
    },
  });

  // Handle form submission
  const handleSubmit = (event) => {
    event.preventDefault();

    // Validate inputs
    if (!firstName || !lastName || !email || !password) {
      toast.error('All fields are required');
      return;
    }

    if (password.length < 8) {
      toast.error('Password must be at least 8 characters long');
      return;
    }

    // Prepare payload for the API request
    const payload = {
      first_name: firstName,
      last_name: lastName,
      email,
      password,
    };

    console.log('Sending signup request with payload:', payload);

    // Trigger the mutation
    signUp(payload);
  };

  return (
    <div className="min-h-96 flex items-center justify-center bg-gray-50">
      <div className="w-full max-w-md p-8 space-y-6 bg-oxfordBlue rounded-lg shadow-md">
        <h2 className="text-2xl font-bold text-center text-gray-800">Sign Up</h2>
        <form onSubmit={handleSubmit} className="space-y-4">
          {/* First Name Input */}
          <div>
            <label htmlFor="firstName" className="block text-sm font-medium text-gray-700">
              First Name:
            </label>
            <input
              type="text"
              id="firstName"
              value={firstName}
              onChange={(e) => setFirstName(e.target.value)}
              required
              aria-label="First name"
              className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-oxfordBlue focus:border-oxfordBlue sm:text-sm text-black"
            />
          </div>

          {/* Last Name Input */}
          <div>
            <label htmlFor="lastName" className="block text-sm font-medium text-gray-700">
              Last Name:
            </label>
            <input
              type="text"
              id="lastName"
              value={lastName}
              onChange={(e) => setLastName(e.target.value)}
              required
              aria-label="Last name"
              className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-oxfordBlue focus:border-oxfordBlue sm:text-sm text-black"
            />
          </div>

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
            aria-label="Sign up"
            disabled={isLoading} // Disable the button while the request is in progress
            className={`w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white ${
              isLoading ? 'bg-gray-400' : 'bg-indigo-600 hover:bg-indigo-700'
            } focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500`}
          >
            {isLoading ? 'Signing up...' : 'Sign Up'}
          </button>
        </form>
      </div>
    </div>
  );
};

export default SignUp;