import storageService from '@/service/storageService';
import userService from '@/service/userService';

const userModule = {
  namespaced: true,
  state: {
    token: storageService.get(storageService.USER_TOKEN),
    userInfo: JSON.parse(storageService.get(storageService.USER_INFO)),
  },

  mutations: {
    SET_TOKEN(state, token) {
      // update cache
      storageService.set(storageService.USER_TOKEN, token);
      // update state
      state.token = token;
    },
    SET_USERINFO(state, userInfo) {
      // update cache
      storageService.set(storageService.USER_INFO, JSON.stringify(userInfo));
      // update state
      state.userInfo = userInfo;
    },
  },

  actions: {
    register(context, { name, telephone, password }) {
      return new Promise((resolve, reject) => {
        userService.register({ name, telephone, password }).then((res) => {
          context.commit('SET_TOKEN', res.data.data.token);
          return userService.info();
        }).then((res) => {
          // console.log(res.data.data.user);
          context.commit('SET_USERINFO', res.data.data.user);
          resolve(res);
        }).catch((err) => {
          reject(err);
        });
      });
    },
  },
};

export default userModule;
