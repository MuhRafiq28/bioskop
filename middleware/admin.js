export default function ({ store, redirect }) {
  // Mengecek apakah sedang di client side
  if (process.client) {
    // Ambil data user dari localStorage
    const user = localStorage.getItem('user');

    if (user) {
      // Simpan user ke Vuex store
      store.commit('SET_USER', JSON.parse(user));
    } else {
      // Jika tidak ada user, redirect ke halaman login
      return redirect('/login');
    }

    // Pastikan role user adalah admin
    const currentUser = store.state.user;
    if (!currentUser || currentUser.role !== 'admin') {
      return redirect('/');
    }
  }
}
