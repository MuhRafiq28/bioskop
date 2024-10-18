<template>
  <div>
    <h1>Daftar Pengguna</h1>
    <table>
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
            <nuxt-link :to="`/admin/edit-user/${user.id}`">Edit</nuxt-link> |
            <button @click="deleteUser(user.id)">Hapus</button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  data() {
    return {
      users: []
    };
  },
  async fetch() {
    try {
      const response = await this.$axios.get('/api/users');
      this.users = response.data;
    } catch (error) {
      console.error('Error fetching users:', error);
    }
  },
  methods: {
    async deleteUser(id) {
      if (confirm('Apakah anda yakin ingin menghapus user ini?')) {
        try {
          await this.$axios.delete(`/api/users/${id}`);
          this.users = this.users.filter(user => user.id !== id);
        } catch (error) {
          console.error('Error deleting user:', error);
        }
      }
    }
  }
};
</script>

<style scoped>
table {
  width: 100%;
  border-collapse: collapse;
}
th, td {
  border: 1px solid black;
  padding: 8px;
  text-align: left;
}
</style>
