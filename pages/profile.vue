<template>
  <div class="profile">
    <h1>Edit Profil Saya</h1>
    <form @submit.prevent="updateProfile">
      <div>
        <label for="name">Nama:</label>
        <input v-model="form.name" type="text" id="name" required />
      </div>
      <div>
        <label for="email">Email:</label>
        <input v-model="form.email" type="email" id="email" required />
      </div>
      <div>
        <label for="password">Password (Kosongkan jika tidak ingin mengubah):</label>
        <input v-model="form.password" type="password" id="password" />
      </div>
      <button type="submit">Perbarui Profil</button>
    </form>
  </div>
</template>

<script>
export default {
  middleware: 'auth',
  data() {
    return {
      form: {
        name: '',
        email: '',
        password: ''
      }
    };
  },
  async fetch() {
    const userId = this.$store.state.user?.id;
    if (!userId) {
      return this.$router.push('/login');
    }

    try {
      const response = await this.$axios.get(`/api/users/${userId}`);
      this.form = response.data;
    } catch (error) {
      console.error('Error fetching profile:', error);
    }
  },

  methods: {
    async updateProfile() {
      const userId = this.$auth.user.id; // Ambil ID dari auth user
      try {
        const response = await this.$axios.put(`/api/users/${userId}`, this.form);
        if (response.status === 200) {
          // Tindakan setelah berhasil
          this.$router.push('/'); // Arahkan kembali ke halaman utama setelah update berhasil
        }
      } catch (error) {
        console.error('Error updating profile:', error);
        this.error = error.response?.data?.message || 'Gagal memperbarui profil.';
      }
    }
  }
}
</script>

<style scoped>
.profile {
  max-width: 600px;
  margin: 0 auto;
}
</style>
