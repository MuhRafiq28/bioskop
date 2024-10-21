<template>
  <div class="container d-flex">
    <!-- Desktop -->
    <div class="col-lg-6 ml-0 mb-ml-0 ml-md-5 d-none d-md-block">
      <h1 class="judul_d font-weight-bold mb-3">Website Pembelian Tiket Bioskop Terpercaya</h1>
      <div class="col-lg-9">
        <p class="description_d">Ngebioskop website pembelian tiket yang sudah dipakai oleh banyak orang, harga tiketnya
          terjangkau, dan sudah sangat dipercaya</p>
      </div>
      <div v-if="!user" class="col-lg-6 d-flex justify-content-between align-items-center mt-3">
        <router-link to="/login"><button class="btn btn-primary">Pesan Sekarang</button></router-link>
        <router-link to="/#"><button class="btn btn-success">About Me</button></router-link>
      </div>
      <div v-if="user" class="col-lg-6 d-flex justify-content-between align-items-center mt-3">
        <button @click="masuk" class="btn btn-primary">Masuk {{ user.name }}</button>
        <router-link to="/#"><button class="btn btn-success">About Me</button></router-link>
      </div>
    </div>

    <!-- Auto Slide Gambar di Desktop dengan Transisi -->
    <div class="col-lg-6 d-none d-md-flex justify-content-center align-items-center">
      <transition name="slide-fade" mode="out-in">
        <div class="card bayangan" :key="film[currentIndex].judul">
          <img :src="film[currentIndex].img" :alt="film[currentIndex].judul" style="width: 280px; height: 350px;">
          <div class="card-body">
            <h2 class="card-title font-weight-bold text-center text-white">{{ film[currentIndex].judul }}</h2>
          </div>
        </div>
      </transition>
      <transition name="slide-fade" mode="out-in">
        <div class="card depan" :key="film[currentIndex].judul + 'depan'">
          <img :src="film[currentIndex].img" :alt="film[currentIndex].judul" style="width: 280px; height: 350px;">
          <div class="card-body">
            <h2 class="card-title font-weight-bold text-center text-white">{{ film[currentIndex].judul }}</h2>
          </div>
        </div>
      </transition>
    </div>

    <!-- Mobile -->
    <div class="flex flex-wrap">
      <div class="col-lg-6 d-md-none d-flex">
        <transition name="slide-fade" mode="out-in">
          <div class="col-lg-6 d-flex  mb-5" :key="film[currentIndex].judul + 'mobile'">
            <div class="card bayangan_m">
              <img :src="film[currentIndex].img" :alt="film[currentIndex].judul" style="width: 280px; height: 290px;">
              <div class="card-body">
                <h2 class="card-title font-weight-bold text-center text-white">{{ film[currentIndex].judul }}</h2>
              </div>
            </div>
          </div>
        </transition>
        <transition name="slide-fade" mode="out-in">
          <div class="card depan_m" :key="film[currentIndex].judul + 'depan'">
            <img :src="film[currentIndex].img" :alt="film[currentIndex].judul" style="width: 280px; height: 290px;">
            <div class="card-body">
              <h2 class="card-title font-weight-bold text-center text-white">{{ film[currentIndex].judul }}</h2>
            </div>
          </div>
        </transition>
      </div>

      <div class="col-12 d-md-none">
        <h1 class="judul_m font-weight-bold mb-3">Website Pembelian Tiket Bioskop Terpercaya</h1>
        <div class="col-lg-9">
          <p class="description_m">Ngebioskop website pembelian tiket yang sudah dipakai oleh banyak orang, harga
            tiketnya terjangkau, dan sudah sangat dipercaya</p>
        </div>
        <div v-if="!user" class="col-lg-6 d-flex justify-content-between align-items-center mt-3">
          <router-link to="/login"><button class="btn btn-primary">Pesan Sekarang</button></router-link>
          <router-link to="/#"><button class="btn btn-success">About Me</button></router-link>
        </div>
        <div v-if="user" class="col-lg-6 d-flex justify-content-between align-items-center mt-3">
          <button @click="masuk" class="btn btn-primary">Masuk {{ user.name }}</button>
          <router-link to="/#"><button class="btn btn-success">About Me</button></router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "HeroLanding",
  computed: {
    user() {
      return this.$store.state.user; // Ambil data pengguna dari Vuex store
    },
  },
  data() {
    return {
      film: [
        {
          img: 'images/godzila.jpg',
          judul: 'Godzila X Kong',
        },
        {
          img: 'images/dilan-1991.jpg',
          judul: 'Dilan 1991',
        },
        {
          img: 'images/dilan-1991.jpg',
          judul: 'Avatar 2',
        },
      ],
      currentIndex: 0,
    };
  },
  mounted() {
    this.startAutoSlide();
  },
  methods: {
    startAutoSlide() {
      setInterval(() => {
        this.currentIndex = (this.currentIndex + 1) % this.film.length;
      }, 3000);
    },
    masuk() {
      // Cek role dari user
      if (this.user.role === 'admin') {
        this.$router.push('/homeadmin'); // Arahkan ke halaman admin
      } else {
        this.$router.push('/homeuser'); // Arahkan ke halaman user biasa
      }
    }
  },
};
</script>

<style scoped>
.judul_d {
  font-size: 60px;
}

.description_d {
  font-size: 20px;
}

.judul_m {
  font-size: 30px;
}

.description_m {
  font-size: 10px;
}

.depan {
  position: absolute;
  top: 10px;
  left: 120px;
  background-color: rgb(185, 48, 20);
  box-shadow: inset 0 46px 69px rgba(0, 0, 0, 0.1);
}

.bayangan {
  background-color: rgb(156, 40, 17);
  border-radius: 5px;
}

.depan_m {
  position: absolute;
  top: 20px;
  left: 50px;
  background-color: rgb(185, 48, 20);
  box-shadow: inset 0 46px 69px rgba(0, 0, 0, 0.1);
}

.bayangan_m {
  background-color: rgb(156, 40, 17);
  border-radius: 5px;
}

/* Transisi Gambar */
.slide-fade-enter-active,
.slide-fade-leave-active {
  transition: all 0.5s ease;
}

.slide-fade-enter,
.slide-fade-leave-to
  {
  opacity: 0;
  transform: translateX(10px);
}
</style>
