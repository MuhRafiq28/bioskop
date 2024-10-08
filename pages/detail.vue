<template>
  <div class="container mb-5">
    <AppNavbar />
    <div class="container-detailTiket mt-5 p-5 rounded-lg pl-5 pr-5">
      <div class="detailTiket">
        <div class="img">
          <img class="rounded-lg" src="/images/dilan-1991.jpg" alt="Tiket"> <!-- Mengambil gambar dari folder static -->
        </div>
        <div class="desTiket">
          <h1>Dilan 1991</h1>
          <h4>romantisme</h4>
          <p>menemukan cinta di sekolah</p>
        </div>
      </div>
      <div class="hargaTiket d-flex justify-content-between">
        <div class="detail">
        <h3>Bandung</h3>
        <p>29/09/07</p>
        <p>23:00</p>
      </div>
      <div class="harga">
        <h3 class="mb-2">Rp.100</h3>
        <button class="rounded-sm">Pesan</button>
      </div>
      </div>
      <div class="hargaTiket d-flex justify-content-between">
        <div class="detail">
        <h3>Bandung</h3>
        <p>29/09/07</p>
        <p>23:00</p>
      </div>
      <div class="harga">
        <h3 class="mb-2">Rp.100</h3>
        <button class="rounded-sm">Pesan</button>
      </div>
      </div>
      <div class="hargaTiket d-flex justify-content-between">
        <div class="detail">
        <h3>Bandung</h3>
        <p>29/09/07</p>
        <p>23:00</p>
      </div>
      <div class="harga">
        <h3 class="mb-2">Rp.100</h3>
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
  }
}
</script>

<style>
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
  max-width: 200px;
  height: auto;
}

.desTiket {
  margin-left: 20px;
}

.hargaTiket {
  margin-top: 20px;
}
</style>
