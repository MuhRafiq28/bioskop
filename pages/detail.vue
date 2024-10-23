<template>
  <div class="container mb-5">
    <AppNavbar />
    <div class="container-detailTiket p-5 rounded-lg pl-5 pr-5 mt-8">
      <div class="detailTiket">
        <div class="img col-lg-6" @click="showTrailer">
          <img
            v-if="movie.image_url"
            :src="`http://localhost:8080${movie.image_url}`"
            :alt="movie.title"
            class="card-img-top"
            style="height: 400px; cursor: pointer;"
          />
          <span v-else class="text-center">Tidak ada gambar</span>
        </div>
        <div class="desTiket col-lg-6 d-flex flex-column align-items-center justify-content-center text-center">
          <h1 class="display-3"><strong>{{ movie.title }}</strong></h1>
          <h4>{{ movie.genre }}</h4>
          <p>{{ movie.description }}</p>
        </div>
      </div>
      <div v-if="showVideo" class="video-container">
        <iframe
          width="100%"
          height="400"
          :src="trailerUrl"
          frameborder="0"
          allowfullscreen
        ></iframe>
        <button class="btn btn-danger mt-2" @click="showVideo = false">Tutup</button>
      </div>
      <div class="hargaTiket d-flex justify-content-between">
        <div class="detail">
          <p>{{ movie.release_date }}</p>
          <p>{{ movie.ticketQuantity }}</p>
        </div>
        <div class="harga">
          <h3 class="mb-2">{{ movie.harga }}</h3>
          <button class="rounded-sm">Pesan</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import AppNavbar from '~/components/AppNavbar.vue'; // Pastikan jalur file benar

export default {
  name: "detail",
  components: {
    AppNavbar
  },
  data() {
    return {
      movie: {
        title: '',
        genre: '',
        description: '',
        release_date: '',
        ticketQuantity: '',
        harga: '', // Pastikan ini ada di data film
        imageFile: '',
        error: '',
        trailer_url: ''
      },
      showVideo: false,
      trailerUrl: '',
    };
  },
  async mounted() {
    const id = this.$route.query.id;
    try {
      const response = await this.$axios.get(`/movies/${id}`);
      this.movie = response.data;

      // Ubah URL trailer menjadi format embed
      if (this.movie.trailer_url) {
        const videoId = this.movie.trailer_url.split('v=')[1]; // Mengambil ID dari URL
        if (videoId) {
          this.trailerUrl = `https://www.youtube.com/embed/${videoId}`;
        }
      }
    } catch (err) {
      this.error = 'Gagal mengambil data film.';
    }
  },
  methods: {
    showTrailer() {
      if (this.trailerUrl) {
        this.showVideo = true; // Tampilkan video
      } else {
        alert('Trailer tidak tersedia.');
      }
    },
  }
}
</script>

<style scoped>
.container-detailTiket {
  border: 1px solid #ccc;
  padding: 20px;
  max-width: 800px;
  margin: 0 auto;
}

.detailTiket {
  display: flex;
  justify-content: space-between;
}

.img img {
  height: auto;
}

.desTiket {
  margin-left: 20px;
}

.hargaTiket {
  margin-top: 20px;
}

.video-container {
  margin-top: 20px;
  text-align: center;
}

.mt-8 {
  margin-top: 100px;
}
</style>
