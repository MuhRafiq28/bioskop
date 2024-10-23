<template>
  <div>
    <AppNavbar />
    <div class="container mt-8">
      <!-- Search bar -->
      <div class="row mb-4">
        <div class="col-md-6">
          <input
            type="text"
            v-model="searchQuery"
            placeholder="Cari film..."
            class="form-control"
          />
        </div>
        <div class="col-md-2">
          <button @click="searchMovies" class="btn btn-primary w-100">Search</button>
        </div>
      </div>

      <div class="row">
        <!-- Tampilkan pesan error jika gagal mendapatkan data -->
        <div v-if="error" class="alert alert-danger">{{ error }}</div>

        <!-- Jika tidak ada film -->
        <div v-if="filteredMovies.length === 0 && !error" class="alert alert-info">Tidak ada film untuk ditampilkan.</div>

        <!-- Looping untuk setiap movie -->
        <div v-for="movie in filteredMovies" :key="movie.id" class="col-md-3 col-sm-6 mb-4">
          <div class="card h-100">
            <img
              v-if="movie.image_url"
              :src="`http://localhost:8080${movie.image_url}`"
              :alt="movie.title"
              class="card-img-top"
              style="height: 400px; "
              @click="pesan(movie.id)"
            />
            <span v-else class="text-center">Tidak ada gambar</span>
            <div class="card-body d-flex flex-column align-items-center">
              <h5 class="card-title m-0 p-0">{{ movie.title }}</h5>
              <p class="card-text m-0 p-0">{{ movie.genre }}</p>
              <button @click="pesan(movie.id)" class="btn btn-success mt-2">Pesan</button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import AppNavbar from '../components/AppNavbar.vue';

export default {
  name: 'film',
  components: {
    AppNavbar
  },
  data() {
    return {
      movies: [],
      searchQuery: '',
      error: null,
    };
  },
  computed: {
    filteredMovies() {
      return this.movies.filter(movie =>
        movie.title.toLowerCase().includes(this.searchQuery.toLowerCase())
      );
    }
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
    pesan(id) {
      this.$router.push(`/detail?id=${id}`);
    },
    searchMovies() {
      // Fungsi ini akan berjalan setiap kali tombol search ditekan
      // Namun, pencarian sudah dilakukan di computed `filteredMovies`
    }
  }
}
</script>

<style scoped>
.container {
  margin: auto;
  padding: 20px;
}

.card {
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

.card-img-top {
  max-height: 320px;
  object-fit: cover;
  transition: transform 0.3s ease-in-out;
}

.card-img-top:hover {
  transform: scale(1.1);
}

.row {
  display: flex;
  flex-wrap: wrap;
  justify-content: space-around;
}

.alert {
  margin-top: 20px;
}

.mt-8 {
  margin-top: 100px;
}

.mb-4 {
  margin-bottom: 1.5rem;
}
</style>
