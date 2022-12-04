<template>
    <div class="login-container">
        <a-form class="login-form">
            <h2 class="title">进销存管理系统登录</h2>
            <a-form-item>
                <a-input class="inputBox" placeholder="请输入数据库" v-model="loginData.database">
                    <template #prefix><HddOutlined style="color: rgba(0, 0, 0, 0.25)" /></template>
                </a-input>
            </a-form-item>
            <a-form-item>
                <a-input class="inputBox" placeholder="请输入账号" v-model="loginData.name">
                    <template #prefix><UserOutlined style="color: rgba(0, 0, 0, 0.25)" /></template>
                </a-input>
            </a-form-item>
            <a-form-item>
                <a-input-password class="inputBox" placeholder="请输入密码" v-model="loginData.password">
                    <template #prefix><LockOutlined style="color: rgba(0, 0, 0, 0.25)" /></template>
                </a-input-password>
            </a-form-item>
            <a-form-item>
                <a-button class="submit" type="primary" @click="onLogin">登录</a-button>
            </a-form-item>
        </a-form>
    </div>
</template>

<script setup lang="ts">
    import { defineEmits, UnwrapRef, reactive } from 'vue';
    import { UserOutlined, LockOutlined, HddOutlined } from '@ant-design/icons-vue';

    // go
    import { Login } from '../../wailsjs/go/service/Service.js';

    // 定义事件
    const emits = defineEmits([
        'login', // 登录事件
    ]);

    // 登录结构体
    interface LoginData {
        name:  string;
        password: string;
        database: string;
    }

    const loginData: UnwrapRef<LoginData> = reactive({
        name: '',
        password: '',
        database: '',
    })
    // 登录
    const onLogin = () => {
        emits('login', false)
        Login(loginData.name, loginData.password, loginData.database).then((result) => {
            console.log(result)
        });
    }
</script>

<style>
    .login-form {
        width: 565px;
        height: 372px;
        margin: 0 auto;
        padding: 160px 110px;
    }

    /* 背景 */
    .login-container {
        position: absolute;
        width: 100%;
        height: 100%;
        background: url("../assets/images/login.png");
    }

    /* Log */
    .login-title {
        color: #fff;
        text-align: center;
        margin: 150px 0;
        font-size: 48px;
    }
    /* 登陆按钮 */
    .submit{
        width: 100%;
        height: 45px;
        font-size: 16px;
    }
    /* 用户登陆标题 */
    .title{
        margin-bottom: 50px;
        color: #fff;
        font-weight: 700;
        font-size: 24px;
        font-family: Microsoft Yahei;
    }
    /* 输入框 */
    .inputBox{
        height: 45px;
    }
    /* 输入框内左边距50px */
    .ant-input-affix-wrapper .ant-input:not(:first-child) {
        padding-left: 50px;
    }
</style>