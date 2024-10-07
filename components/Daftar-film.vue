<template>
  <div class="container p-4 mt-3 m-lg-5">
    <div class="judul">
      <h1>Film Rekomendasi</h1>
    </div>
    <div class="d-flex justify-content-between">
      <div class="container-car d-flex pt-4 pb-4 align-content-center">
        <div class="gambar-poster d-flex align-items-center">
          <img class="poster" src="images/godzila.jpg" alt="" />
        </div>
        <div class="descripsi-poster ml-2 p-4 ">
          <h1><strong> Godzila X Kong </strong></h1>
          <p class="m-0"><strong> Si Botak </strong></p>
          <p class="m-0">Konoha</p>
          <p class="mt-4">
            cerita tentang monster yang hidup di intin bumi yang kemudian hari
            naik ke permukaan bumi yaitu tempat manusia hidup
          </p>
        </div>
      </div>

      <div class="container-car d-flex pt-4 pb-4 align-content-center ml-3">
        <div class="gambar-poster">
          <img class="poster" src="images/godzila.jpg" alt="" />
        </div>
        <div class="descripsi-poster ml-2 p-4">
          <h1><strong> Godzila X Kong </strong></h1>
          <p class="m-0"><strong> Si Botak </strong></p>
          <p class="m-0">Konoha</p>
          <p class="mt-4">
            cerita tentang monster yang hidup di intin bumi yang kemudian hari
            naik ke permukaan bumi yaitu tempat manusia hidup
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  nama: "DaftarFilm",
};
</script>

<style>
.poster {
  width: 140px;
  height: auto;
  border-radius: 10px;
}
</style>
