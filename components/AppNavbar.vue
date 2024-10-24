navbar

<template>
  <div class="fixed-top mt-3 col-12 mb-col-11 ml-0 mb-ml-5 ">
    <b-navbar toggleable="lg" class="content p-2 p-md-0  rounded-5">
      <div class="container">
        <b-navbar-brand href="#" class="judul btn btn-white font-weight-bold">NgeBioskop</b-navbar-brand>

        <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>
        <b-collapse v-if="!user" id="nav-collapse" is-nav>
          <b-navbar-nav class="ml-auto">
            <b-nav-item>
              <router-link to="/login">
                <b-button class="px-3 bg-dark">Login <b-icon-door-open-fill></b-icon-door-open-fill></b-button>
              </router-link>
            </b-nav-item>
          </b-navbar-nav>
        </b-collapse>

        <b-collapse v-if="user" id="nav-collapse" is-nav>
          <b-navbar-nav class="ml-auto">
            <b-nav-item>
              <a @click="home" class="nav-link btn btn-white text-dark" to="/">Home</a>
            </b-nav-item>
            <b-nav-item>
              <router-link class="nav-link btn btn-white text-dark" to="/hasilCheckout">Pesanan
                <b-icon-cart></b-icon-cart></router-link>
            </b-nav-item>
            <b-nav-item>
              <router-link class="nav-link btn btn-white text-dark" to="/film">Detail Film
                <b-icon-film></b-icon-film></router-link>
            </b-nav-item>
          </b-navbar-nav>
          <b-navbar-nav>
            <b-nav-item>
              <nuxt-link :to="{ name: 'profile', query: { id: user.id } }">
                <b-button class="px-3 text-white bg-dark">
                  {{ user.name }} <b-icon-person></b-icon-person>
                </b-button>
              </nuxt-link>
              <b-button variant="outline-danger" class="ml-2" @click="logout">Logout
                <b-icon-door-closed></b-icon-door-closed></b-button>
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
  methods: {
    logout() {
      if (confirm("Apakah Anda Yakin Akan Logout")) {
        this.$store.dispatch('logout'); // Panggil action logout di Vuex
        this.$router.push('/'); // Arahkan pengguna kembali ke halaman utama setelah logout
      }
    },
    home() {
      // Cek role dari user
      if (this.user.role === 'admin') {
        this.$router.push('/homeadmin'); // Arahkan ke halaman admin
      } else {
        this.$router.push('/homeuser'); // Arahkan ke halaman user biasa
      }
    }
  },
}
</script>

<style scoped>
.content {
  background-color: rgba(255, 255, 255, 0.5);
}

.rounded-5 {
  border-radius: 20px !important;
}

.judul {
  font-family: 'Life Savers', cursive;
}

.btn-white {
  background-color: rgba(255, 255, 255, 0.8);
}

.btn-white:hover {
  background-color: rgba(240, 240, 240, 0.8);
}
</style>
