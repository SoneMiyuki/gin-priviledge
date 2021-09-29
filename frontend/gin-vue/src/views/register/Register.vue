<template>
    <div class="registery">
        <b-row class="mt-5">
            <b-col md = "8" offset-md="2" lg="6" offset-lg="3">
        <b-card title="注册页面">
    <b-form>
      <b-form-group label="姓名">
        <b-form-input
          v-model="$v.user.name.$model"
          type="text"
          placeholder="请输入用户名"
        ></b-form-input>
      </b-form-group>
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
        <b-button variant="outline-primary" block @click="register">注册</b-button>
      </b-form-group>
    </b-form>
    </b-card>
    </b-col>
        </b-row>
    </div>
</template>
<script>
import { required, minLength, maxLength } from 'vuelidate/lib/validators';

import { mapActions } from 'vuex';
import customValidator from '@/helper/validator';

export default {
  data() {
    return {
      user: {
        name: '',
        telephone: '',
        password: '',
      },
      validation: null,
    };
  },
  validations: {
    user: {
      name: {

      },
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
    ...mapActions('userModule', { userRegister: 'register' }),
    validateState(name) {
      const { $dirty, $error } = this.$v.user[name];
      return $dirty ? !$error : null;
    },
    register() {
      this.$v.user.$touch();
      if (this.$v.user.$anyError) {
        return;
      }
      this.userRegister(this.user).then(() => {
        this.$router.replace({ name: 'Home' });
      }).catch((err) => {
        console.log('err:', err.response.data.msg);
        this.$bvToast.toast(err.response.data.msg, {
          title: '数据验证错误',
          variant: 'danger',
          solid: true,
        });
      });
    },
  },
};
</script>
<style lang="scss" scoped>

</style>
