<template>
  <div class="container">
    <h1>Daftar Film</h1>
    <button @click="$router.push('/admin/add-movie')" class="btn btn-success mb-3">Tambah Film</button>

    <div v-if="movies.length === 0">
      <p>Belum ada film.</p>
    </div>

    <table class="table table-striped" v-if="movies.length > 0">
      <thead>
        <tr>
          <th>Gambar</th>
          <th>Judul</th>
          <th>Genre</th>
          <th>Deskripsi</th>
          <th>Tanggal Rilis</th>
          <th>Aksi</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="movie in movies" :key="movie.id">
          <td>
            <img v-if="movie.image_url" :src="`http://localhost:8080${movie.image_url}`" :alt="movie.title"
              style="width: 100px; height: auto;" />
            <span v-else>Tidak ada gambar</span>
          </td>
          <td>{{ movie.title }}</td>
          <td>{{ movie.genre }}</td>
          <td>{{ movie.description }}</td>
          <td>{{ formatDate(movie.release_date) }}</td>
          <td>
            <button @click="editMovie(movie.id)" class="btn btn-warning">Edit</button>
            <button @click="deleteMovie(movie.id)" class="btn btn-danger">Hapus</button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  middleware: ['admin'],
  data() {
    return {
      movies: [],
      error: null,
    };
  },
  async mounted() {
    try {
      const response = await this.$axios.get('/movies');
      this.movies = response.data;
    } catch (err) {
      this.error = 'Gagal mengambil data film.';
    }
  },
  methods: {
    formatDate(date) {
      return new Date(date).toLocaleDateString();
    },
    editMovie(id) {
      this.$router.push(`/admin/edit-movie?id=${id}`);
    },
    async deleteMovie(id) {
      if (confirm('Apakah Anda yakin ingin menghapus film ini?')) {
        const token = localStorage.getItem('token'); // Ambil token dari localStorage
        console.log(`Menghapus film dengan ID: ${id}`); // Log ID
        try {
          const response = await fetch(`http://localhost:8080/movies/${id}`, {
            method: 'DELETE',
            headers: {
              'Authorization': `Bearer ${token}`,
              'Content-Type': 'application/json'
            }
          });

          if (!response.ok) {
            const errorText = await response.text(); // Ambil pesan error dari server
            throw new Error(`Network response was not ok: ${errorText}`);
          }

          // Tindakan setelah penghapusan berhasil
          this.movies = this.movies.filter(movie => movie.id !== id);
        } catch (error) {
          console.error('Error:', error);
          this.error = 'Gagal menghapus film.';
        }
      }
    }
  }
};
</script>

<style scoped>
.table {
  width: 100%;
}
</style>
