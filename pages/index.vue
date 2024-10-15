<template>
  <div>
<<<<<<< HEAD
    <Hero />
    <daftar-film />
=======
    <AppNavbar />
<<<<<<< HEAD
    <HeroLanding class="mt-8" />
=======
    <Hero class="mt-8" />
    <DaftarFilm />
    <FilmTayang />
>>>>>>> b8e0901da59fa3f8fb3588fe1ed099f640fa4f0b
>>>>>>> 8953ce946ff4774439496c89f836ee1eb63d3db3
  </div>
</template>

<script>
<<<<<<< HEAD
import DaftarFilm from '../components/Daftar-film.vue';
=======
import AppNavbar from '../components/AppNavbar.vue';
<<<<<<< HEAD
import HeroLanding from '../components/landing/HeroLanding.vue';
=======
import FilmTayang from '../components/FilmTayang.vue';
>>>>>>> b8e0901da59fa3f8fb3588fe1ed099f640fa4f0b
import Hero from '../components/hero.vue';
>>>>>>> 8953ce946ff4774439496c89f836ee1eb63d3db3

export default {
  name: 'IndexPage',
  components: {
<<<<<<< HEAD
    HeroLanding,
    AppNavbar,
=======
    Hero,
<<<<<<< HEAD
    DaftarFilm
=======
    AppNavbar,
    FilmTayang
>>>>>>> b8e0901da59fa3f8fb3588fe1ed099f640fa4f0b
>>>>>>> 8953ce946ff4774439496c89f836ee1eb63d3db3
  }
}
</script>

<style scoped>
.mt-8{
  margin-top: 100px;
}

</style>
