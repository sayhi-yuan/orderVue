<template>
    <LoginLayout v-if="!isLogin" @login="onLogin"/>
    <MainLayout v-else/>
</template>

<script setup>
    import { ref } from 'vue';

    import MainLayout from './layouts/MainLayout.vue'
    import LoginLayout from './layouts/LoginLayout.vue'

    // 登录
    const isLogin = ref(false)
    const onLogin = (status) => {
        isLogin.value = status
    };
</script>
    
<style>
</style>
