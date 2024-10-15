<template>
  <div>
    <AppNavbar />
    <HeroLanding class="mt-8" />
  </div>
</template>

<script>
import AppNavbar from '../components/AppNavbar.vue';
import HeroLanding from '../components/landing/HeroLanding.vue';

export default {
  name: 'IndexPage',
  components: {
    HeroLanding,
    AppNavbar,
  }
}
</script>

<style scoped>
.mt-8{
  margin-top: 100px;
}

</style>
