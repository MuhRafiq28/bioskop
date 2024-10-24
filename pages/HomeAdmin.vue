<template>
  <div>
    <AppNavbar />
    <Hero class="mt-8" />
    <div class="container">
      <nuxt-link to="/admin/movies" class="btn btn-primary">Kelola Movies <b-icon-film></b-icon-film></nuxt-link>
      <nuxt-link to="/admin/users" class="btn btn-primary">Kelola Pengguna <b-icon-person-badge></b-icon-person-badge></nuxt-link>
    </div>
    <DaftarFilm />
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

<style scoped>
.mt-8 {
  margin-top: 80px;
}
</style>
