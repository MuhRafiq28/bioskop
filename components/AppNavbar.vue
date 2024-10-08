<template>
  <div class="fixed-top">
    <b-navbar toggleable="lg" type="dark" variant="dark" class="shadow-sm ">
      <div class="container">
        <b-navbar-brand href="#" class="font-weight-bold">NgeBioskop</b-navbar-brand>

        <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

        <b-collapse id="nav-collapse" is-nav>
          <b-navbar-nav class="ml-auto">
            <b-nav-item>
              <router-link class="nav-link text-white" to="/">Home</router-link>
            </b-nav-item>
            <b-nav-item>
              <router-link class="nav-link text-white" to="/#">Pesanan</router-link>
            </b-nav-item>
          </b-navbar-nav>

          <b-navbar-nav>
            <b-nav-item v-if="!user">
              <router-link to="/login"><b-button variant="outline-light" class="px-3">Login</b-button></router-link>
            </b-nav-item>
            <b-nav-item v-if="user">
              <b-button class="px-3 text-dark" style="background: #D9D9D9;">{{ user.name }}</b-button>
            </b-nav-item>
          </b-navbar-nav>
        </b-collapse>
      </div>
    </b-navbar>
  </div>
</template>

<script>
export default {
  name: 'AppNavbar',
  computed: {
    user() {
      return this.$store.state.user; // Ambil data pengguna dari Vuex store
    },
  },
}
</script>
