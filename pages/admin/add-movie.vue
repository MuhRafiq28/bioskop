<template>
  <div class="container">
    <h1>Tambah Film</h1>
    <form @submit.prevent="addMovie">
      <div class="form-group">
        <label for="title">Judul</label>
        <input type="text" v-model="title" id="title" class="form-control" required />
      </div>
      <div class="form-group">
        <label for="genre">Genre</label>
        <input type="text" v-model="genre" id="genre" class="form-control" required />
      </div>
      <div class="form-group">
        <label for="description">Deskripsi</label>
        <textarea v-model="description" id="description" class="form-control" required></textarea>
      </div>
      <div class="form-group">
        <label for="releaseDate">Tanggal Rilis</label>
        <input type="date" v-model="releaseDate" id="releaseDate" class="form-control" required />
      </div>
      <div class="form-group">
        <label for="image">Gambar</label>
        <input type="file" @change="onFileChange" id="image" class="form-control" required />
      </div>
      <button type="submit" class="btn btn-success">Tambah Film</button>
    </form>
    <div v-if="error" class="alert alert-danger mt-3">{{ error }}</div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      title: '',
      genre: '',
      description: '',
      releaseDate: '',
      imageFile: null,
      error: null,
    };
  },
  methods: {
    onFileChange(event) {
      this.imageFile = event.target.files[0];
    },
    async addMovie() {
      const formData = new FormData();
      formData.append('title', this.title);
      formData.append('genre', this.genre);
      formData.append('description', this.description);
      formData.append('release_date', this.releaseDate);
      formData.append('image', this.imageFile); // pastikan ini berisi file gambar yang valid

      try {
        const response = await this.$axios.post('http://localhost:8080/movies', formData, {
          headers: {
            Authorization: `Bearer ${this.token}`,
          },
        });
        alert('Film berhasil ditambahkan: ' + response.data.message);
        this.$router.push('/homeadmin');
      } catch (error) {
        this.error = 'Gagal menambahkan film: ' + error.response.data.message;
      }
    }
  }
};
</script>

<style scoped>
.container {
  margin-top: 20px;
}
</style>
