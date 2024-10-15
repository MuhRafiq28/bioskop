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
      <button type="submit" class="btn btn-success">Simpan</button>
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
    async updateMovie() {
      try {
        await this.$axios.put(`/movies/${this.$route.query.id}`, this.movie);
        this.$router.push('/admin/movies');
      } catch (err) {
        this.error = 'Gagal memperbarui film.';
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
