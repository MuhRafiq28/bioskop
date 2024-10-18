<template>
  <div class="edit-user">
    <h1>Edit Pengguna</h1>
    <form @submit.prevent="updateUser">
      <div>
        <label for="name">Nama:</label>
        <input v-model="form.name" type="text" id="name" required />
      </div>
      <div>
        <label for="email">Email:</label>
        <input v-model="form.email" type="email" id="email" required />
      </div>
      <div>
        <label for="role">Role:</label>
        <select v-model="form.role" id="role" required>
          <option value="admin">Admin</option>
          <option value="user">User</option>
        </select>
      </div>
      <button type="submit">Perbarui Pengguna</button>
    </form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      form: {
        name: '',
        email: '',
        role: ''
      }
    };
  },
  async fetch() {
    const userId = this.$route.params.id;
    try {
      const response = await this.$axios.get(`/api/users/${userId}`);
      this.form = response.data;
    } catch (error) {
      console.error('Error fetching user:', error);
    }
  },
  methods: {
    async updateUser() {
      const userId = this.$route.params.id;
      try {
        await this.$axios.put(`/api/users/${userId}`, this.form);
        this.$router.push('/admin/users');
      } catch (error) {
        console.error('Error updating user:', error);
      }
    }
  }
};
</script>

<style scoped>
.edit-user {
  max-width: 600px;
  margin: 0 auto;
}
</style>
