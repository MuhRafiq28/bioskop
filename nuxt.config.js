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

  auth: {
    strategies: {
      local: {
        endpoints: {
          login: { url: 'api/auth/login', method: 'post' },
          logout: { url: 'api/auth/logout', method: 'post' },
          user: { url: 'api/user', method: 'get' }
        }
      }
    }
  },
  router: {
    extendRoutes(routes, resolve) {
      routes.push(
        {
          name: 'users',
          path: '/admin/users',
          component: resolve(__dirname, 'pages/admin/users.vue')
        },
        {
          name: 'add-user',
          path: '/admin/add-user',
          component: resolve(__dirname, 'pages/admin/add-user.vue')
        },
        {
          name: 'edit-user',
          path: '/admin/edit-user',
          component: resolve(__dirname, 'pages/admin/edit-user.vue')
        }
      );
    }
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

  plugins: [
    '~/plugins/bootstrap-vue.js',
    { src: '~/plugins/auth.js', mode: 'client' },
  ],

  // Build Configuration
  build: {}
}
