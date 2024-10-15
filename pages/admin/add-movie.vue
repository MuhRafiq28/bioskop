<template>
  <div class="container">
    <h1>Tambah Film</h1>
    <form @submit.prevent="addMovie" enctype="multipart/form-data">
      <div class="form-group">
        <label>Judul</label>
        <input v-model="movie.title" class="form-control" required />
      </div>
      <div class="form-group">
        <label>Genre</label>
        <input v-model="movie.genre" class="form-control" required />
      </div>
      <div class="form-group">
        <label>Deskripsi</label>
        <textarea v-model="movie.description" class="form-control" required></textarea>
      </div>
      <div class="form-group">
        <label>Tanggal Rilis</label>
        <input type="date" v-model="movie.release_date" class="form-control" required />
      </div>
      <div class="form-group">
        <label>Gambar</label>
        <input type="file" @change="handleFileUpload" class="form-control" required />
      </div>
      <button type="submit" class="btn btn-success">Simpan</button>

      <!-- Tampilkan error jika ada -->
      <p v-if="error" class="text-danger">{{ error }}</p>
    </form>
  </div>
</template>

<script>
export default {
  middleware: ['admin'],
  data() {
    return {
      movie: {
        title: '',
        genre: '',
        description: '',
        release_date: '',
      },
      image: null,
      error: null,
    };
  },
  methods: {
    handleFileUpload(event) {
      const file = event.target.files[0];

      // Validasi ukuran file dan tipe file
      if (file && file.size > 2 * 1024 * 1024) { // Maks 2MB
        this.error = 'Ukuran file terlalu besar. Maksimal 2MB.';
        this.image = null;
      } else if (file && !file.type.startsWith('image/')) {
        this.error = 'File harus berupa gambar.';
        this.image = null;
      } else {
        this.image = file;
        this.error = null; // Hapus pesan error jika validasi berhasil
      }
    },
    async addMovie() {
      // Cek jika tidak ada gambar yang dipilih
      if (!this.image) {
        this.error = 'Silakan pilih gambar terlebih dahulu.';
        return;
      }

      const formData = new FormData();
      formData.append('title', this.movie.title);
      formData.append('genre', this.movie.genre);
      formData.append('description', this.movie.description);
      formData.append('release_date', this.movie.release_date);
      formData.append('image', this.image);

      try {
        await this.$axios.post('/movies', formData);
        alert('Film berhasil ditambahkan.');
        this.$router.push('/admin/movies');
      } catch (err) {
        if (err.response && err.response.data) {
          this.error = err.response.data.message || 'Gagal menambahkan film.';
        } else {
          this.error = 'Gagal menambahkan film.';
        }
      }
    }
  }
};
</script>

<style scoped>
.container {
  max-width: 600px;
  margin: auto;
}

.text-danger {
  color: red;
}
</style>
