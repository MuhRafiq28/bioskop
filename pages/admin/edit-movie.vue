<template>
  <div class="container">
    <h1>Edit Film</h1>
    <form @submit.prevent="updateMovie">
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
        <input type="file" @change="onFileChange" class="form-control" />
      </div>
      <button type="submit" class="btn btn-success">Simpan</button>
      <div v-if="error" class="text-danger">{{ error }}</div>
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
      imageFile: null,
      error: null,
    };
  },
  async mounted() {
    const id = this.$route.query.id;
    try {
      const response = await this.$axios.get(`/movies/${id}`);
      this.movie = response.data;
    } catch (err) {
      this.error = 'Gagal mengambil data film.';
    }
  },
  methods: {
    onFileChange(event) {
      this.imageFile = event.target.files[0];
    },
    async updateMovie() {
      // Cek apakah format tanggal valid
      const datePattern = /^\d{4}-\d{2}-\d{2}$/; // YYYY-MM-DD
      if (!datePattern.test(this.movie.release_date)) {
        this.error = 'Format tanggal tidak valid, gunakan YYYY-MM-DD';
        return; // Hentikan proses jika format tidak valid
      }

      const formData = new FormData();
      formData.append('title', this.movie.title);
      formData.append('genre', this.movie.genre);
      formData.append('description', this.movie.description);
      formData.append('release_date', this.movie.release_date);

      if (this.imageFile) {
        formData.append('image', this.imageFile);
      }

      // Log data yang dikirim
      console.log('Data yang akan dikirim:', {
        title: this.movie.title,
        genre: this.movie.genre,
        description: this.movie.description,
        release_date: this.movie.release_date,
        image: this.imageFile ? this.imageFile.name : 'Tidak ada gambar'
      });

      try {
        await this.$axios.put(`/movies/${this.$route.query.id}`, formData, {
          headers: {
            'Content-Type': 'multipart/form-data',
          },
        });
        this.$router.push('/admin/movies');
      } catch (err) {
        this.error = 'Gagal memperbarui film.';
        console.error(err.response.data); // Log error dari server
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
</style>
