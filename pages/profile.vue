<template>
  <div class="container mt-5">
    <h2 class="text-center">Profil Pengguna</h2>
    <div v-if="user" class="profile-details">
      <p><strong>Nama:</strong> {{ user.name }}</p>
      <p><strong>Email:</strong> {{ user.email }}</p>
      <p><strong>Role:</strong> {{ user.role }}</p>
      <router-link :to="{ name: 'edit-profile', query: { id: user.id } }" class="btn btn-primary mt-3">
        Edit Profil
      </router-link>
    </div>
    <div v-if="error" class="alert alert-danger mt-3">{{ error }}</div>
    <button @click="kembali" class="btn btn-secondary mt-3">Kembali</button>
  </div>
</template>

<script>
export default {
  data() {
    return {
      user: null,
      error: ''
    };
  },
  async mounted() {
  const id = this.$route.query.id;
  console.log("ID dari query params:", id); // Tambahkan log untuk memeriksa id
  if (!id) {
    this.error = 'ID pengguna tidak ditemukan.';
    return;
  }

  try {
    const token = localStorage.getItem('token');
    const response = await this.$axios.get('/users/' + id, {
      headers: {
        Authorization: `Bearer ${token}`
      }
    });
    this.user = response.data;
  } catch (error) {
    this.error = 'Gagal mengambil data profil pengguna.';
  }
},
methods: {
    kembali() {
      // Cek role dari user
      if (this.user.role === 'admin') {
        this.$router.push('/homeadmin'); // Arahkan ke halaman admin
      } else {
        this.$router.push('/homeuser'); // Arahkan ke halaman user biasa
      }
    }
  },
}
</script>

<style scoped>
.container {
  max-width: 600px;
  margin: 0 auto;
}
</style>
