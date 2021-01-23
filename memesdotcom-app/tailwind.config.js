module.exports = {
  purge: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  darkMode: false, // or 'media' or 'class'
  theme: {
    backgroundColor: theme => ({
      ...theme('colors'),
      dark: '#1e272e'
    }),
    extend: {
      fontSize: {
        xxs: '.625rem'
      },
      spacing: {
        72: '18rem',
        84: '21rem',
        96: '24rem',
        144: '36rem'
      },
      color: {

      }
    }
  },
  variants: {
    extend: {}
  },
  plugins: []
}
