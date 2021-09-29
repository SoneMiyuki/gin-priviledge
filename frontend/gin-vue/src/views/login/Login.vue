<template>
    <div class="register">
        <b-row class="mt-5">
            <b-col md = "8" offset-md="2" lg="6" offset-lg="3">
        <b-card title="登录页面">
    <b-form>
            <b-form-group label="手机">
        <b-form-input
          v-model="$v.user.telephone.$model"
          type="number"
          required
          placeholder="请输入手机"
        ></b-form-input>
      <b-form-invalid-feedback :state="validateState('telephone')">
        手机号必须为11位
      </b-form-invalid-feedback>
      <b-form-valid-feedback :state="validateState('telephone')">
        手机号填写成功
      </b-form-valid-feedback>
      </b-form-group>
            <b-form-group label="密码">
        <b-form-input
          v-model="$v.user.password.$model"
          type="password"
          placeholder="请输入密码"
        ></b-form-input>
      </b-form-group>
      <b-form-group>
        <b-button variant="outline-primary" block @click="login">登录</b-button>
      </b-form-group>
    </b-form>
    </b-card>
    </b-col>
        </b-row>
    </div>
</template>
<script>
import { required, minLength, maxLength } from 'vuelidate/lib/validators';
import customValidator from '@/helper/validator';
import storageService from '@/service/storageService';

export default {
  data() {
    return {
      user: {
        telephone: '',
        password: '',
      },
      validation: null,
    };
  },
  validations: {
    user: {
      telephone: {
        required,
        telephone: customValidator.telephoneValidator,
      },
      password: {
        required,
        minLength: minLength(6),
        maxLength: maxLength(15),
      },
    },
  },
  methods: {
    validateState(name) {
      const { $dirty, $error } = this.$v.user[name];
      return $dirty ? !$error : null;
    },
    login() {
      this.$v.user.$touch();
      if (this.$v.user.$anyError) {
        return;
      }
      console.log('');
      const api = 'http://localhost:8090/api/auth/login';
      this.axios.post(api, { ...this.user }).then((res) => {
        storageService.set(storageService.USER_TOKEN, res.data.data.token);
        this.$router.replace({ name: 'Home' });
      }).catch((err) => {
        console.log('err:', err.response.data.msg);
      });
    },
  },
};
</script>
<style lang="scss" scoped>

</style>
