<template>
  <div>
    <AppNavbar />
    <h1>User</h1>
    <Hero class="mt-5" />
    <DaftarFilm />
    <FilmTayang />
  </div>
</template>

<script>
import AppNavbar from '../components/AppNavbar.vue';
import FilmTayang from '../components/FilmTayang.vue';
import Hero from '../components/hero.vue';

export default {
  name: 'IndexPage',
  components: {
    Hero,
    AppNavbar,
    FilmTayang
  }
}
</script>
