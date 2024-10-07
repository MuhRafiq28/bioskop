<template>
  <div class="container p-4">
    <div class="judul text-center">
      <h1>Film Rekomendasi</h1>
    </div>
    <div class="slideshow-container">
      <div class="row">
        <div 
          class="col-6"
          v-for="(film, index) in displayedFilms"
          :key="index"
        >
          <div class="film-card  ">
            <img class="poster" :src="film.poster" alt="" />
            <div class="film-details ml-3">
              <h2 class="film-title"><strong>{{ film.judul }}</strong></h2>
              <p class="film-maker"><strong>{{ film.pembuat }}</strong></p>
              <p class="film-description">{{ film.deskripsi }}</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "DaftarFilm",
  data() {
    return {
      films: [
        {
          judul: "Godzila X Kong",
          pembuat: "Si Botak",
          deskripsi: "Cerita tentang monster yang hidup di dalam bumi dan naik ke permukaan.",
          poster: "images/godzila.jpg",
        },
        {
          judul: "Film Lain",
          pembuat: "Si Botak",
          deskripsi: "Deskripsi untuk film lain.",
          poster: "images/godzila.jpg",
        },
        {
          judul: "Film Ketiga",
          pembuat: "Si Botak",
          deskripsi: "Deskripsi untuk film ketiga.",
          poster: "images/godzila.jpg",
        },
        {
          judul: "Film Keempat",
          pembuat: "Si Botak",
          deskripsi: "Deskripsi untuk film keempat.",
          poster: "images/godzila.jpg",
        },
      ],
      currentSlide: 0,
    };
  },
  computed: {
    displayedFilms() {
      return [
        this.films[this.currentSlide % this.films.length],
        this.films[(this.currentSlide + 1) % this.films.length],
      ];
    },
  },
  mounted() {
    this.startSlideshow();
  },
  methods: {
    startSlideshow() {
      setInterval(() => {
        this.currentSlide = (this.currentSlide + 1) % this.films.length;
      }, 3000);
    },
  },
};
</script>

<style>
.poster {
  width: 140px;
  height: auto;
  border-radius: 10px;
}
.slideshow-container {
  position: relative;
  max-width: 800px;
  margin: auto;
}
.row {
  display: flex;
  flex-wrap: wrap;
}
.col-6 {
  width: 50%;
  padding: 10px;
}
.film-card {
  display: flex;
  align-items: flex-start;
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 8px;
  background-color: #f9f9f9;
  margin-left: 50px;
}
.film-details {
  margin-left: 15px;
  flex: 1;
  margin-top: 10px;
}
.film-title {
  font-size: 1.25rem;
  margin-bottom: 5px;
}
.film-maker, .film-description {
  margin: 0;
}


@media (max-width: 1100px) {
  .col-6 {
    width: 100%;
  }
}


@media (max-width: 768px) {
  .film-card {
    display: block;
    align-items: center;
  }
  .poster {
    width: 80%;
  }
}
</style>
