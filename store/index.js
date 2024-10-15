// store/index.js
export const state = () => ({
  user: null // State untuk menyimpan data pengguna
});

export const mutations = {
  SET_USER(state, user) {
    state.user = user; // Menyimpan data pengguna di state
  },
  CLEAR_USER(state) {
    state.user = null; // Menghapus data pengguna
  }
};

export const actions = {
  async login({ commit }, credentials) {
    try {
      const response = await this.$axios.post('/api/login', credentials);
      const user = response.data.user;
      commit('SET_USER', user); // Pastikan 'user' memiliki properti 'role'
    } catch (error) {
      console.error('Login failed:', error);
    }
  },
  logout({ commit }) {
    commit('CLEAR_USER'); // Menghapus data pengguna saat logout
  }
};
