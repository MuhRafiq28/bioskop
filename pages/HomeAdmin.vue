<template>
  <div>
    <AppNavbar />
    <Hero class="mt-8" />
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

<style scoped>
.mt-8{
  margin-top: 80px;
}
</style>
