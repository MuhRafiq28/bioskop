<template>
  <div class="container d-flex align-items-center">
    <div class="m-2 m-md-5 bg-abu rounded-lg p-3 col-11">
      <h1 class="mb-3" style="font-family: 'Life Savers', cursive;">Film Yag Akan Tayang</h1>
      <div v-for="(film, index) in film" :key="index" class="m-1 row justify-content-md-center rounded-lg bg-hitam p-2 mb-5 mb-md-2 ">
        <div class="col col-lg-2 d-flex align-items-center justify-content-center">
          <img class="rounded-lg" :src="film.img" :alt="film.judul" style="width: 150px;">
        </div>
        <div class="col-lg-8 text-white ">
          <h1 class="ml-3 m-0 mt-3 mb-mt-0 fw-bold" style="font-family: 'Life Savers', cursive;">{{ film.judul }}</h1><br>
          <p class="ml-3 m-0">{{ film.deskripsi }}</p><br>
          <p class="ml-3 m-0">{{ film.waktu }}</p><br>
          <h5 class="ml-3 m-0 mb-2">{{ film.harga }}</h5>
        </div>
        <div class="col col-lg-2 d-flex align-items-end justify-content-between">
          <router-link to="/checkout" class="btn btn-success float-bottom mt-3 mb-mt-0 mb-3 ml-4 ">
            Pesan
          </router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'FilmTayang',
  data() {
    return {
      film: [
        {
          img: 'images/godzila.jpg',
          judul: 'Godzila X Kong',
          deskripsi: 'cerita tentang monster yang hidup di intin bumi  yang kemudian hari naik ke  permukaan bumi yaitu tempat manusia hidup',
          waktu: 'waktu tayang 13.20',
          harga: 'Rp. 999.999'
        },
        {
          img: 'images/dilan-1991.jpg',
          judul: 'Dilan 1991',
          deskripsi: 'cerita tentang anak sekolah yang bernama dilan seorang panglima perang di sekolahnnya, yang tiba tiba tertarik pada seorang murid baru cewe cantik ',
          waktu : 'waktu tayang 22.20',
          harga: 'Rp. 777.777'
        },
      ]
    }
  }
}
</script>

<style scoped>
.bg-abu {
  background: #D9D9D9;
}

.bg-hitam {
  background: #3F3D56;
}
</style>
