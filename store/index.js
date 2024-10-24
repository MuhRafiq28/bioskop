export const state = () => ({
  user: null,
});

export const mutations = {
  SET_USER(state, user) {
    if (user && user.id) {
      state.user = user;
    }
  },
  CLEAR_USER(state) {
    state.user = null;
  }
};

export const actions = {
  nuxtClientInit({ commit }) {
    const user = JSON.parse(localStorage.getItem('user'));
    console.log('LocalStorage User:', user); // Tambahkan log ini untuk cek isi localStorage
    if (user && user.id) {
      commit('SET_USER', user);
      console.log('User loaded into Vuex:', user); // Cek apakah user berhasil dimasukkan ke Vuex
    }
  },

  async login({ commit }, credentials) {
    try {
      const response = await this.$axios.post('/api/login', credentials);
      const user = response.data.user;
      if (user && user.id) {
        localStorage.setItem('user', JSON.stringify(user));
        commit('SET_USER', user);
      }
    } catch (error) {
      console.error('Login failed:', error);
    }
  },

  logout({ commit }) {
    commit('CLEAR_USER');
    localStorage.removeItem('user');
  }
};
