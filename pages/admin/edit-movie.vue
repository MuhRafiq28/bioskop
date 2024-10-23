<template>
  <div class="edit-movie">
    <h1>Edit Film</h1>
    <form @submit.prevent="updateMovie">
      <input v-model="movie.title" placeholder="Judul" class="form-control" required />
      <input v-model="movie.genre" placeholder="Genre" class="form-control" required />
      <textarea v-model="movie.description" placeholder="Deskripsi" class="form-control" required></textarea>
      <input type="date" v-model="movie.release_date" class="form-control" required />
      <input type="number" v-model="movie.harga" placeholder="Harga" class="form-control" required />
      <input type="text" v-model="movie.trailer_url" placeholder="Trailer URL" class="form-control" required />
      <input type="number" v-model="movie.jumlah_tiket" placeholder="Jumlah Tiket" class="form-control" required />
      <input type="file" @change="onFileChange" class="form-control" />

      <div v-if="movie.image_url"> <!-- Menampilkan gambar jika ada -->
        <h5>Gambar Saat Ini:</h5>
        <img :src="movie.image_url" alt="Current Movie Image" class="img-thumbnail" style="max-width: 200px; max-height: 200px;" />
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
        harga: null,
        trailer_url: '',
        jumlah_tiket: null,
        image_url: '', // Menyimpan URL gambar
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

      // Jika gambar disimpan dengan URL, masukkan ke movie.image_url
      this.movie.image_url = response.data.image; // Pastikan nama ini sesuai dengan field di backend
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
      formData.append('harga', this.movie.harga); // Menambahkan harga
      formData.append('trailer_url', this.movie.trailer_url); // Menambahkan trailer_url
      formData.append('jumlah_tiket', this.movie.jumlah_tiket); // Menambahkan jumlah_tiket

      if (this.imageFile) {
        formData.append('image', this.imageFile);
      }

      // Log data yang dikirim
      console.log('Data yang akan dikirim:', {
        title: this.movie.title,
        genre: this.movie.genre,
        description: this.movie.description,
        release_date: this.movie.release_date,
        harga: this.movie.harga,
        trailer_url: this.movie.trailer_url,
        jumlah_tiket: this.movie.jumlah_tiket,
        image: this.imageFile ? this.imageFile.name : 'Tidak ada gambar'
      });

      try {
        await this.$axios.put(`/movies/${this.$route.query.id}`, formData, {
          headers: {
            'Content-Type': 'multipart/form-data',
            'Authorization': `Bearer ${localStorage.getItem('token')}`
          },
        });

        // Tampilkan alert setelah berhasil mengupdate film
        alert('Film berhasil diperbarui!');

        // Arahkan ke halaman movies setelah berhasil memperbarui film
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
.edit-movie {
  max-width: 600px;
  margin: auto;
  padding: 20px;
  border: 1px solid #ccc;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  background-color: #f9f9f9;
}

.edit-movie h1 {
  text-align: center;
  margin-bottom: 20px;
}

.edit-movie input,
.edit-movie textarea,
.edit-movie button {
  width: 100%;
  padding: 10px;
  margin: 10px 0;
  border: 1px solid #ddd;
  border-radius: 4px;
}

.edit-movie button {
  background-color: #28a745;
  color: white;
  border: none;
  cursor: pointer;
}

.edit-movie button:hover {
  background-color: #218838;
}

.img-thumbnail {
  margin-top: 10px;
}
</style>
