export default ({ store }) => {
  if (process.client) {
    // Ambil data user dari localStorage
    const user = localStorage.getItem('user');
    console.log('User from localStorage:', user);

    if (user) {
      // Simpan user ke Vuex store
      store.commit('SET_USER', JSON.parse(user));
      console.log('User loaded to Vuex store:', JSON.parse(user));
    }
  }
}
