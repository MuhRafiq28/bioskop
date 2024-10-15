<template>
  <div>
<<<<<<< HEAD
    <Hero />
    <daftar-film />
=======
    <AppNavbar />
    <Hero class="mt-8" />
    <DaftarFilm />
    <FilmTayang />
>>>>>>> b8e0901da59fa3f8fb3588fe1ed099f640fa4f0b
  </div>
</template>

<script>
<<<<<<< HEAD
import DaftarFilm from '../components/Daftar-film.vue';
=======
import AppNavbar from '../components/AppNavbar.vue';
import FilmTayang from '../components/FilmTayang.vue';
>>>>>>> b8e0901da59fa3f8fb3588fe1ed099f640fa4f0b
import Hero from '../components/hero.vue';

export default {
  name: 'IndexPage',
  components: {
    Hero,
<<<<<<< HEAD
    DaftarFilm
=======
    AppNavbar,
    FilmTayang
>>>>>>> b8e0901da59fa3f8fb3588fe1ed099f640fa4f0b
  }
}
</script>

<style scoped>
.mt-8{
  margin-top: 100px;
}

</style>
