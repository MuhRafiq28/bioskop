export default function ({ store, redirect }) {
  if (process.client) {
    const user = localStorage.getItem('user');

    if (user) {
      const parsedUser = JSON.parse(user);
      if (parsedUser && parsedUser.id) {
        store.commit('SET_USER', parsedUser);
      }
    } else {
      return redirect('/login');
    }
  }
}
