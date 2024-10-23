<template>
  <div class="container d-flex align-items-center justify-content-center vh-100">
    <div class="m-2 m-md-5 bg-abu rounded-lg p-5 col-11 col-md-6 align-items-center d-block text-center">
      <h1><strong>LOGIN</strong></h1>
      <p>Silahkan Login <strong>Masukkan Data Anda</strong></p>
      <form @submit.prevent="login">
        <div class="mb-3">
          <input v-model="email" placeholder="Masukkan Email" type="email" id="email" required />
        </div>
        <div class="mb-3">
          <input v-model="password" placeholder="Masukkan Password" type="password" id="password" required />
        </div>
        <router-link to="/register"><button type="button" class="btn btn-danger mr-2">Daftar</button></router-link>
        <button type="submit" :disabled="loading" class="btn btn-success">Masuk</button>
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
      email: '',
      password: '',
      loading: false,
      error: null,
    };
  },
  methods: {
    async login() {
  this.loading = true;
  this.error = null;
  try {
    const response = await this.$axios.post('/login', {
      email: this.email,
      password: this.password,
    });

    const user = response.data;
    if (user && user.id) {
      localStorage.setItem('user', JSON.stringify(user));
      localStorage.setItem('token', user.token); // Simpan token
      this.$store.commit('SET_USER', user);

      // Set header default untuk Axios
      this.$axios.defaults.headers.common['Authorization'] = `Bearer ${localStorage.getItem('token')}`;

      // Navigasi berdasarkan role
      if (user.role === 'admin') {
        this.$router.push('/homeadmin');
      } else if (user.role === 'user') {
        this.$router.push('/homeuser');
      } else {
        this.$router.push('/');
      }
    } else {
      this.error = 'Data pengguna tidak valid';
    }
  } catch (error) {
    this.error = error.response?.data?.message || 'Login gagal, coba lagi.';
  } finally {
    this.loading = false;
  }
}
  }
};
</script>

<style scoped>
.bg-abu {
  background: #D9D9D9;
}
</style>
