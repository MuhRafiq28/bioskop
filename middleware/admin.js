// middleware/admin.js
export default function({ store, redirect }) {
  const user = store.state.user; // Ambil state pengguna dari Vuex

  // Periksa apakah user ada dan memiliki role 'admin'
  if (!user || user.role !== 'admin') {
    return redirect('/'); // Redirect jika tidak ada user atau bukan admin
  }
}
