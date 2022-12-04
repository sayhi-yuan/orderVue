<template>
    <div class="login-container">
        <a-form ref="form" :model="form" class="login-form">
        <h2 class="title">进销存管理系统登录</h2>
        <a-form-item>
            <a-input class="inputBox" placeholder="请输入账号">
                <template #prefix><UserOutlined style="color: rgba(0, 0, 0, 0.25)" /></template>
            </a-input>
        </a-form-item>
        <a-form-item>
            <a-input-password class="inputBox" placeholder="请输入密码" >
                <template #prefix><LockOutlined style="color: rgba(0, 0, 0, 0.25)" /></template>
            </a-input-password>
        </a-form-item>

        <a-form-item>
            <a-button class="submit" type="primary" @click="onLogin">登录</a-button>
        </a-form-item>
        </a-form>
    </div>
</template>

<script setup>
    import { defineEmits } from 'vue';
    import { UserOutlined, LockOutlined } from '@ant-design/icons-vue';

    // 定义事件
    const emits = defineEmits([
        'login', // 登录事件
    ]);

    // 登录
    const onLogin = () => {
        emits('login', true)
        // 存储local storage
        window.localStorage.setItem('token', '测试')
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