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
            <b-nav-item>
              <b-button variant="outline-light" class="px-3">Login</b-button>
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
}
</script>
<style scoped>
.nav-link {
  transition: color 0.3s ease;
}

.nav-link:hover {
  color: #5DC4D2 !important;
}

.b-navbar-brand {
  font-family: 'Poppins', sans-serif;
  font-size: 1.5rem;
  color: #fff !important;
}

.b-navbar {
  background-color: #007bff;
}

.shadow-sm {
  box-shadow: 0 0.5rem 1rem rgba(0, 0, 0, 0.15);
}
</style>
