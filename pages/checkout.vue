<template>
  <div>
    <AppNavbar />
    <!-- peropil pemesanan -->
    <div class="container p-5">
      <div class="detailfilem d-flex justify-content-between mt-7 p-4">
        <div class="img">
          <img class="jpg w-5 rounded-c" src="images/dilan-1991.jpg" alt="" />
        </div>
        <div class="text ml-50px align-content-center">
          <h1>Godzila X Kong</h1>
          <p class="mt-3">
            cerita tentang monster yang hidup di intin bumi yang kemudian hari
            naik ke permukaan bumi yaitu tempat manusia hidup
          </p>
          <p class="mt-7">
            <span class="ml-2">Minutes</span>
          </p>
        </div>
      </div>
      <hr />

       <!-- detail pemesanan -->
      <template>
  <div>
    <div
      v-for="(detail, index) in details"
      :key="index"
      class="detailharga d-flex justify-content-between align-items-center"
    >
      <div class="info">
        <h2>{{ detail.city }}</h2>
        <p>{{ detail.date }}</p>
        <p>{{ detail.time }}</p>
      </div>

      <div class="harga d-flex ">
        <h3>Rp. <strong>{{ detail.price }}</strong></h3>
        <button class=" ml-2 btn btn-primary">Pesan</button>
        <br><hr>
      </div>
    </div>
  </div>
</template>


    </div>
  </div>
</template>

<script>
import AppNavbar from "../components/AppNavbar.vue";
export default {
  name: "checkout",
  components: { AppNavbar },
  data() {
    return {
      details: [
        {
          city: 'Bandung',
          date: '29-07-2024',
          time: '20 : 23',
          price: '123.000',
        },
        {
          city: 'Jakarta',
          date: '30-07-2024',
          time: '15 : 00',
          price: '150.000',
        },
        {
          city: 'Surabaya',
          date: '01-08-2024',
          time: '18 : 45',
          price: '175.000',
        },
      ],
    };
  },
};
</script>

<style scoped>
.harga h3 {
  margin-bottom: 0;
  margin-top: 2px;
}
.mt-7 {
  margin-top: 70px;
}

.w-5 {
  width: 250px;
}

.ml-50px {
  margin-left: 50px;
}

.ct {
  color: white;
}

.mt-50px {
  margin-top: 50px;
}

.rounded-c {
  border-radius: 15px;
}

button {
  font-family: inherit;
  border: none;
  outline: 1px dotted rgb(37, 37, 37);
  outline-offset: -4px;
  cursor: pointer;
  background: hsl(0deg 0% 75%);
  box-shadow:
    inset -1px -1px #292929,
    inset 1px 1px #fff,
    inset -2px -2px rgb(158, 158, 158),
    inset 2px 2px #ffffff;
  font-size: 14px;
  text-transform: uppercase;
  letter-spacing: 2px;
  padding: 5px 30px;
}

button:active {
  box-shadow:
    inset -1px -1px #fff,
    inset 1px 1px #292929,
    inset -2px -2px #ffffff,
    inset 2px 2px rgb(158, 158, 158);
}

</style>
