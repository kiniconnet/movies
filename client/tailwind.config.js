/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    colors: {
        black: "oklch(0.208 0.042 265.755)",
        oxfordBlue: "hsla(221, 51%, 16%, 1)",
        orangeWeb: "hsla(37, 98%, 53%, 1)",
        platinum: "oklch(0.554 0.046 257.417)",
        green:"rgb(21 128 61 / var(--tw-bg-opacity, 1))",
        white: "hsla(0, 0%, 100%, 1)",
    },
    fontFamily: {
      sans: ['Graphik', 'sans-serif'],
      serif: ['Merriweather', 'serif'],
      poppins: ['Poppins'],
    },
    extend: {},
  },
  plugins: [],
}