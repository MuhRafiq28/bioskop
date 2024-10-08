export default {
  // Global page headers
  head: {
    title: 'bioskop',
    htmlAttrs: {
      lang: 'en'
    },
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: '' },
      { name: 'format-detection', content: 'telephone=no' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' },
      {
        rel: 'stylesheet',
        href: 'https://fonts.googleapis.com/css2?family=Life+Savers:wght@400;700&display=swap'
      }
    ]
  },

  // Global CSS
  css: [],

  // Plugins to run before rendering page
  plugins: [],

  // Auto import components
  components: true,

  // Modules for dev and build
  buildModules: [],

  // Modules
  modules: [
    // https://go.nuxtjs.dev/bootstrap
    'bootstrap-vue/nuxt',
    '@nuxtjs/axios', // Tambahkan ini
  ],

  // Axios module configuration
  axios: {
    baseURL: 'http://localhost:8080',
  },

  // Build Configuration
  build: {}
}
