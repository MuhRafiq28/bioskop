<template>
  <div class="container mt-5">
    <h2 class="text-center">Edit Profil Pengguna</h2>
    <form @submit.prevent="updateProfile"> <!-- Pastikan memanggil updateProfile di sini -->
      <div class="form-group">
        <label for="name">Nama:</label>
        <input type="text" id="name" class="form-control" v-model="user.name" required />
      </div>

      <div class="form-group">
        <label for="email">Email:</label>
        <input type="email" id="email" class="form-control" v-model="user.email" required />
      </div>

      <div class="form-group">
        <label for="password">Password Baru:</label>
        <input type="password" id="password" class="form-control" v-model="user.password"
          placeholder="Kosongkan jika tidak ingin mengubah" autocomplete="new-password" />
      </div>


      <button type="submit" class="btn btn-primary mt-3">Perbarui Profil</button>
    </form>

    <div v-if="message" class="alert alert-info mt-3">{{ message }}</div>
    <router-link :to="{ name: 'profile', query: { id: user.id } }" class="btn btn-secondary mt-3">Kembali ke
      Profil</router-link>
  </div>
</template>

<script>
export default {
  data() {
    return {
      user: {
        id: '',
        name: '',
        email: '',
        password: ''
      },
      message: ''
    };
  },
  async mounted() {
    try {
      const token = localStorage.getItem('token');
      const response = await this.$axios.get('/users/' + this.$route.query.id, {
        headers: {
          Authorization: `Bearer ${token}`
        }
      });
      this.user = response.data;
    } catch (error) {
      this.message = 'Gagal mengambil data pengguna.';
    }
  },
  methods: {
    async updateProfile() {
      try {
        const token = localStorage.getItem('token');
        const response = await this.$axios.put(`/users/${this.user.id}`, this.user, {
          headers: {
            Authorization: `Bearer ${token}`
          }
        });

        console.log('Response after update:', response.data); // Tambahkan log ini

        // Memperbarui data pengguna di Vuex
        this.$store.commit('SET_USER', response.data);

        // Jika server mengembalikan token baru, simpan token baru
        if (response.data.token) {
          localStorage.setItem('token', response.data.token);
        }

        localStorage.setItem('user', JSON.stringify(response.data)); // Update localStorage setelah profil diupdate
        this.message = 'Profil berhasil diperbarui.';
      } catch (error) {
        this.message = 'Gagal memperbarui profil.';
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
