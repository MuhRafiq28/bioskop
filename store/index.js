export const state = () => ({
  user: null,
});

export const mutations = {
  setUser(state, user) {
    state.user = user;
  },
  logout(state) {
    state.user = null;
  },
};

export const actions = {
  login({ commit }, user) {
    commit('setUser', user);
  },
  logout({ commit }) {
    commit('logout');
  },
};
