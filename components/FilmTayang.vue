<template>
  <div class="container d-flex justify-content-between align-items-center">
    <div v-for="(film, index) in film" :key="index" class="md:m-0 bg-abu m-5 rounded-lg p-3 col-11">
      <h1>Film Yag Akan Tayang</h1>
      <div class="m-1 row justify-content-md-center bg-hitam rounded-lg">
        <div class="col col-lg-2">
          1 of 3
        </div>
        <div class="col-lg-6 text-white">
          {{ film.judul }}<br>
          {{ film.deskripsi }}<br>
          {{ film.harga }}
        </div>
        <div class="col col-lg-1">
          <button class="btn btn-success float-bottom">
            Pesan
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'FilmTayang',
  data(){
    return{
      film: [
        {
          img : '#',
          judul: 'Godzila X Kong',
          deskripsi: 'cerita tentang monster yang hidup di intin bumi  yang kemudian hari naik ke  permukaan bumi yaitu tempat manusia hidup',
          waktu : 'waktu tayang 13.20',
          harga: 'Rp. 999.999'
        }
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
