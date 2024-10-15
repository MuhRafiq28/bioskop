<template>
  <div>
    <AppNavbar />
    <h1>Admin</h1>
    <Hero class="mt-5" />
    <DaftarFilm />
    <nuxt-link to="/admin/movies">Kelola Pengguna</nuxt-link>
    <FilmTayang />
  </div>
</template>

<script>
import AppNavbar from '../components/AppNavbar.vue';
import FilmTayang from '../components/FilmTayang.vue';
import Hero from '../components/hero.vue';

export default {
  middleware: ['admin'],
  name: 'IndexPage',
  components: {
    Hero,
    AppNavbar,
    FilmTayang
  }
}
</script>
