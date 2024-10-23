<template>
  <div class="add-movie">
    <h1>Add Movie</h1>
    <form @submit.prevent="addMovie">
      <input type="text" v-model="title" placeholder="Title" required />
      <input type="text" v-model="genre" placeholder="Genre" required />
      <textarea v-model="description" placeholder="Description" required></textarea>
      <input type="date" v-model="releaseDate" required />
      <input type="number" v-model="price" placeholder="Price" required />
      <input type="text" v-model="trailerUrl" placeholder="Trailer URL" required />
      <input type="number" v-model="ticketQuantity" placeholder="Ticket Quantity" required />
      <input type="file" @change="onFileChange" required />
      <button type="submit">Add Movie</button>
    </form>
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
      price: null,
      trailerUrl: '',
      ticketQuantity: null,
      imageFile: null,
    };
  },
  methods: {
    onFileChange(event) {
      this.imageFile = event.target.files[0]; // Menyimpan file gambar yang dipilih
    },
    async addMovie() {
      // Validasi semua field
      if (!this.title || !this.genre || !this.description ||
          !this.releaseDate || !this.price || !this.trailerUrl ||
          !this.ticketQuantity || !this.imageFile) {
        alert('Semua field diperlukan!');
        return;
      }

      const formData = new FormData();
      formData.append('title', this.title);
      formData.append('genre', this.genre);
      formData.append('description', this.description);
      formData.append('release_date', this.releaseDate); // Pastikan ini menggunakan 'release_date'
      formData.append('harga', this.price); // Ganti 'price' menjadi 'harga'
      formData.append('trailer_url', this.trailerUrl); // Ganti 'trailerUrl' menjadi 'trailer_url'
      formData.append('jumlah_tiket', this.ticketQuantity); // Ganti 'ticketQuantity' menjadi 'jumlah_tiket'
      formData.append('image', this.imageFile);

      try {
        const response = await this.$axios.post('http://localhost:8080/movies', formData, {
          headers: {
            'Content-Type': 'multipart/form-data',
            'Authorization': `Bearer ${localStorage.getItem('token')}`
          }
        });
        console.log('Movie added successfully:', response.data);

        // Tampilkan alert sebelum mengarahkan
        alert('Film berhasil ditambahkan!');

        // Arahkan ke halaman movies setelah berhasil menambahkan film
        this.$router.push('/admin/movies');
      } catch (error) {
        console.error('Error adding movie:', error.response?.data || error.message);
        alert(`Error: ${error.response?.data?.message || 'Gagal menambahkan film.'}`);
      }
    }
  },
};
</script>

<style scoped>
.add-movie {
  max-width: 600px;
  margin: 0 auto;
  padding: 20px;
  border: 1px solid #ccc;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  background-color: #f9f9f9;
}

.add-movie h1 {
  text-align: center;
  margin-bottom: 20px;
}

.add-movie input,
.add-movie textarea,
.add-movie button {
  width: 100%;
  padding: 10px;
  margin: 10px 0;
  border: 1px solid #ddd;
  border-radius: 4px;
}

.add-movie button {
  background-color: #28a745;
  color: white;
  border: none;
  cursor: pointer;
}

.add-movie button:hover {
  background-color: #218838;
}
</style>
