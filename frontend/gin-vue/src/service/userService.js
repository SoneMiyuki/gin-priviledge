import request from '@/utils/request';

const register = ({ name, telephone, password }) => {
  // console.log(`url: ${process.env.VUE_APP_BASE_URL}`);
  return request.post('auth/register', { name, telephone, password });
};

const login = ({ telephone, password }) => {
  // console.log(`url: ${process.env.VUE_APP_BASE_URL}`);
  return request.post('auth/login', { telephone, password });
};

const info = () => {
  return request.get('auth/info');
};

export default {
  register,
  info,
  login,
};
