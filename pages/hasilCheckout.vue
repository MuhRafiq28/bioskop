<template>
  <div class="p-5">
    <AppNavbar />
    <div class="container-2 mt-7">
      <div class="status p-4 d-flex align-items-center">
        <h5 class="judul">
          <b-icon-basket3></b-icon-basket3><span class="ml-3">Belanja</span>
        </h5>
        <p class="ml-3 waktu-teransaksi">23 Sep 24</p>
        <div
          class="kotak ml-3 d-flex align-items-center justify-content-center"
        >
          Selesai
        </div>
      </div>
      <hr />
      <div class="barang d-flex p-4">
        <div class="foto-barang">
          <h5>Ngebioskop</h5>
          <img class="rounded-sm" src="images/dilan-1991.jpg" alt="" />
        </div>

        <div
          class="detail-barang d-flex justify-content-between align-items-center"
        >
          <div class=" ml-2 nama-barang">
            <h5>Dilan-1991</h5>
            <p class="jum-barang">1X Barang</p>
          </div>
          <div class="harga-barang text-right">
            <p>Total Belanja</p>
            <h5>Rp. 123.000</h5>
            <div
              class="kotak-b ml-3 d-flex align-items-center justify-content-center"
            >
              Beli lagi
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import AppNavbar from "../components/AppNavbar.vue";

export default {
  name: "hasilCheckout",
  comments: { AppNavbar },
};
</script>

<style>
.judul {
  margin-top: 2px;
  margin-bottom: 0;
}
.waktu-teransaksi {
  margin-bottom: 0;
}
.jum-barang {
  color: #594F4F;
}
hr {
  margin: 0;
}
.mt-7 {
  margin-top: 100px;
}
.kotak-b {
  width: 85px;
  height: 29px;
  background: #D9D9D9;
  color: black;
  border-radius: 15px;
  font: bold;
  float: right;
}
.kotak {
  width: 85px;
  height: 29px;
  background: #32EC42;
  color: black;
  border-radius: 15px;
  font: bold;
  float: right;
}
.container-2 {
  box-shadow: 5px 5px 15px rgba(0, 0, 0, 0.3);
  border-radius: 20px;
}
.foto-barang img {
  width: 110px;
  height: 109px;
}
.detail-barang {
  width: 100%; /* Pastikan elemen container memiliki lebar penuh */
}

.harga-barang {
  text-align: right; /* Agar konten harga tersusun rapi di sisi kanan */
}

.nama-barang,
.harga-barang {
  flex: 1; /* Pastikan elemen memiliki fleksibilitas dalam pembagian ruang */
}
</style>
