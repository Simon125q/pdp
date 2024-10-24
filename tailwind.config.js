/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./**/*.html", "./**/*.templ", "./**/*.go"],
  safelist: [],
  theme: {
    colors: {
      "tul-dark-red": "#8b0002",
      "tul-medium-red": "#a43537",
      "tul-light-red": "#c47979",
      "tul-dark-gray": "#1d1d1d",
    },
  },
};
