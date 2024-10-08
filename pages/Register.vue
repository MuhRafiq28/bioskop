<template>
  <div class="container d-flex align-items-center justify-content-center vh-100">
    <div class="m-2 m-md-5 bg-abu rounded-lg p-5 col-11 col-md-6 align-items-center d-block text-center">
      <h1><strong>REGISTER</strong></h1>
      <p>Silahkan Register <strong>Isi Form Dengan Sesuai</strong></p>
      <form @submit.prevent="register">
        <div class="mb-3">
          <input v-model="name" placeholder="Masukan Nama" type="text" id="name" required />
        </div>
        <div class="mb-3">
          <input v-model="email" placeholder="Masukan Email" type="email" id="email" required />
        </div>
        <div class="mb-3">
          <input v-model="password" placeholder="Masukan Password" type="password" id="password" required />
        </div>
        <router-link to="/login"><button class="btn btn-danger mr-2">Kembali</button></router-link>
        <button type="submit" :disabled="loading" class="btn btn-success">Daftar</button>
        <div v-if="loading">Loading...</div>
        <div v-if="error" style="color: red;">{{ error }}</div>
      </form>
    </div>
  </div>
</template>



<script>
export default {
  data() {
    return {
      name: '',
      email: '',
      password: '',
      role: 'user',
      loading: false,
      error: null,
    };
  },
  methods: {
    async register() {
      this.loading = true; // Set loading to true
      this.error = null;   // Reset error message
      try {
        const response = await this.$axios.post('http://localhost:8080/register', {
          name: this.name,
          email: this.email,
          password: this.password,
          role: this.role,
        });
        console.log(response.data);
        alert('Registrasi berhasil!');
        this.$router.push('/login');
      } catch (error) {
        console.error(error);
        this.error = error.response?.data?.message || 'Registrasi gagal, coba lagi.'; // Set error message
      } finally {
        this.loading = false; // Set loading to false
      }
    },
  },
};
</script>
<style scoped>
.p-13{
  padding: 100px;
}
.bg-abu {
  background: #D9D9D9;
}
</style>
