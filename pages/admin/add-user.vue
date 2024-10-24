<template>
  <div class="container mt-5">
    <h2 class="text-center">Tambah Pengguna Baru</h2>
    <form @submit.prevent="addUser">
      <div class="form-group">
        <label for="name">Nama:</label>
        <input type="text" id="name" class="form-control" v-model="user.name" required />
      </div>

      <div class="form-group">
        <label for="email">Email:</label>
        <input type="email" id="email" class="form-control" v-model="user.email" required />
      </div>

      <div class="form-group">
        <label for="password">Password:</label>
        <input type="password" id="password" class="form-control" v-model="user.password" required />
      </div>

      <div class="form-group">
        <label for="role">Role:</label>
        <select id="role" class="form-control" v-model="user.role">
          <option value="user">User</option>
          <option value="admin">Admin</option>
        </select>
      </div>

      <button type="submit" class="btn btn-primary mt-3">Simpan</button>
    </form>

    <div v-if="message" class="alert alert-info mt-3">{{ message }}</div>
    <router-link :to="{ name: 'users' }" class="btn btn-secondary mt-3">Kembali ke Daftar Pengguna</router-link>
  </div>
</template>

<script>
export default {
  data() {
    return {
      user: {
        name: '',
        email: '',
        password: '',
        role: 'user'
      },
      message: ''
    };
  },
  methods: {
    async addUser() {
      try {
        const token = localStorage.getItem('token');
        await this.$axios.post('/users', this.user, {
          headers: {
            Authorization: `Bearer ${token}`
          }
        });
        this.message = 'Pengguna berhasil ditambahkan.';
      } catch (error) {
        this.message = 'Gagal menambahkan pengguna.';
      }
    }
  }
};
</script>

<style scoped>
.container {
  max-width: 600px;
  margin: 0 auto;
}
</style>
