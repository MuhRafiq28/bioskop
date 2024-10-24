<template>
  <div class="users">
    <AppNavbar />
  <div class="container">
    <h2 class="text-center">Manajemen Pengguna</h2>
    <router-link :to="{ name: 'add-user' }" class="btn btn-primary mb-3">Tambah Pengguna</router-link>

    <table class="table table-bordered">
      <thead>
        <tr>
          <th>ID</th>
          <th>Nama</th>
          <th>Email</th>
          <th>Role</th>
          <th>Aksi</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="user in users" :key="user.id">
          <td>{{ user.id }}</td>
          <td>{{ user.name }}</td>
          <td>{{ user.email }}</td>
          <td>{{ user.role }}</td>
          <td>
            <router-link :to="{ name: 'edit-user', query: { id: user.id } }" class="btn btn-warning btn-sm">Edit</router-link>
            <button @click="deleteUser(user.id)" class="btn btn-danger btn-sm">Hapus</button>
          </td>
        </tr>
      </tbody>
    </table>
    <div v-if="message" class="alert alert-info mt-3">{{ message }}</div>
  </div>
</div>
</template>

<script>
import AppNavbar from '../../components/AppNavbar.vue';

export default {
  name:'users',
  components: {
    AppNavbar
  },
  data() {
    return {
      users: [],
      message: ''
    };
  },
  async mounted() {
    try {
      const token = localStorage.getItem('token');
      const response = await this.$axios.get('/users', {
        headers: {
          Authorization: `Bearer ${token}`
        }
      });
      this.users = response.data;
    } catch (error) {
      this.message = 'Gagal mengambil data pengguna.';
    }
  },
  methods: {
    async deleteUser(id) {
      if (confirm('Apakah kamu yakin ingin menghapus pengguna ini?')) {
        try {
          const token = localStorage.getItem('token');
          await this.$axios.delete(`/users/${id}`, {
            headers: {
              Authorization: `Bearer ${token}`
            }
          });
          this.users = this.users.filter(user => user.id !== id);
          this.message = 'Pengguna berhasil dihapus.';
        } catch (error) {
          this.message = 'Gagal menghapus pengguna.';
        }
      }
    }
  }
};
</script>

<style scoped>
.container {
  max-width: 800px;
  margin: 0 auto;
  margin-top: 100px;
}
</style>
