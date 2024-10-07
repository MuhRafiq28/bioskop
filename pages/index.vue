<template>
  <div>
    <Hero />
    <daftar-film />
  </div>
</template>

<script>
import DaftarFilm from '../components/Daftar-film.vue';
import Hero from '../components/hero.vue';

export default {
  name: 'IndexPage',
  components: {
    Hero,
    DaftarFilm
  }
}
</script>
