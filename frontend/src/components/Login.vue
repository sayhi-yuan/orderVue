<script lang="ts">
import { defineComponent, reactive } from 'vue';
import { message } from 'ant-design-vue';

import { Login } from '../../wailsjs/go/service/Service.js';

interface FormState {
  username: string;
  password: string;
  database: string;
  remember: boolean;

  resultText: any;
}

export default defineComponent({
  setup() {
    const formState = reactive<FormState>({
      username: '',
      password: '',
      database: '',
      remember: true,
      resultText: '',
    });
    const onFinish = (values: any) => {
      Login(formState.username,formState.password,formState.database).then(result => {
        if (result.code != 200){
          message.error(result.message);
          return
        }
        formState.resultText = result
      })
    };

    const onFinishFailed = (errorInfo: any) => {
      console.log('Failed:', errorInfo);
    };
    return {
      formState,
      onFinish,
      onFinishFailed,
    };
  },
});
</script>

<template>

  <div id="result" class="result">{{ formState.resultText }}</div>

  <a-form class="login"
      :model="formState"
      name="basic"
      :label-col="{ span: 8 }"
      :wrapper-col="{ span: 16 }"
      autocomplete="off"
      @finish="onFinish"
      @finishFailed="onFinishFailed">
    <a-form-item
        label="Username"
        name="username"
        :rules="[{ required: true, message: 'Please input your username!' }]">
      <a-input v-model:value="formState.username" />
    </a-form-item>

    <a-form-item
        label="Password"
        name="password"
        :rules="[{ required: true, message: 'Please input your password!' }]">
      <a-input-password v-model:value="formState.password" />
    </a-form-item>

    <a-form-item
        label="Database"
        name="database"
        :rules="[{ required: true, message: 'Please input your database!' }]">
      <a-input-password v-model:value="formState.database" />
    </a-form-item>

    <a-form-item name="remember" :wrapper-col="{ offset: 8, span: 16 }">
      <a-checkbox v-model:checked="formState.remember">Remember me</a-checkbox>
    </a-form-item>

    <a-form-item :wrapper-col="{ offset: 8, span: 16 }">
      <a-button type="primary" html-type="submit">Submit</a-button>
    </a-form-item>
  </a-form>
</template>

<style scoped>
 .login{

 }
</style>