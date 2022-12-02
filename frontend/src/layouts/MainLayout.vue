<template>
    <a-layout style="min-height: 100vh">
        <a-layout-sider v-model:collapsed="collapsed" collapsible>
            <a-layout-header style="background: #fff; padding: 0; height: 40px; line-height: 40px;">LOGO</a-layout-header>
        <a-menu v-model:selectedKeys="selectedKeys" theme="dark" mode="inline">
            <a-sub-menu key="sub1">
                <template #title>
                    <span>
                        <bar-chart-outlined />
                        <span>Dashboard</span>
                    </span>
                </template>
                <a-menu-item key="11">
                    <router-link :to="{ path: '/' }">分析页</router-link>
                </a-menu-item>
            </a-sub-menu>
            <a-sub-menu key="sub2">
                <template #title>
                    <span>
                        <fire-outlined />
                        <span>营销管理</span>
                    </span>
                </template>
                <a-menu-item key="21">
                    <router-link :to="{ path: '/marketing-manage/list' }">营销列表</router-link>
                </a-menu-item>
            </a-sub-menu>
            <a-sub-menu key="sub3">
                <template #title>
                    <span>
                        <bulb-outlined />
                        <span>订单管理</span>
                    </span>
                </template>
                <a-menu-item key="31">
                    <router-link :to="{ path: '/order/list' }">订单列表</router-link>
                </a-menu-item>
            </a-sub-menu>
            <a-sub-menu key="sub4">
                <template #title>
                    <span>
                        <appstore-outlined />
                        <span>商品管理</span>
                    </span>
                </template>
                <a-menu-item key="41">
                    <router-link :to="{ path: '/goods/list' }">商品列表</router-link>
                </a-menu-item>
            </a-sub-menu>
            <a-sub-menu key="sub5">
                <template #title>
                    <span>
                        <team-outlined />
                        <span>商家管理</span>
                    </span>
                </template>
                <a-menu-item key="51">
                    <router-link :to="{ path: '/merchants/list' }">商家列表</router-link>
                </a-menu-item>
            </a-sub-menu>
            <a-sub-menu key="sub6">
                <template #title>
                    <span>
                        <credit-card-outlined />
                        <span>打包发货</span>
                    </span>
                </template>
                <a-menu-item key="61">
                    <router-link :to="{ path: '/packaging/list' }">打包列表</router-link>
                </a-menu-item>
            </a-sub-menu>
            <!-- <a-menu-item key="9">
                <file-outlined />
                <span>File</span>
            </a-menu-item> -->
        </a-menu>
        </a-layout-sider>
        <a-layout>
            <!-- 内容的头顶 -->
            <a-layout-header style="background: #fff; padding: 0; height: 40px;">
                <a-avatar style="background-color: #1890ff; float: right; margin-right: 18px; margin-top: 5px;">
                    <template #icon><UserOutlined /></template>
                </a-avatar>
            </a-layout-header> 

            <a-layout-content style="margin: 0 8px">
                <a-layout-content>
                    <a-breadcrumb :style="{margin: '16px 0', textAlign: 'left'}">
                        <a-breadcrumb-item v-for="tab in tabList">{{ tab }}</a-breadcrumb-item>
                    </a-breadcrumb>
                    <div :style="{ padding: '18px', background: '#fff', minHeight: '560px' }">
                        <router-view></router-view>
                    </div>
                </a-layout-content>
            </a-layout-content>
            <a-layout-footer style="text-align: center">
                Ant Design ©2022 Created by Ant UED
            </a-layout-footer>
        </a-layout>
    </a-layout>
</template>

<script lang="ts">
    import {
        TeamOutlined,
        FireOutlined,
        BarChartOutlined,
        CreditCardOutlined,
        AppstoreOutlined,
        BulbOutlined,
        UserOutlined,
    } from '@ant-design/icons-vue';
    import { defineComponent, ref } from 'vue';
    export default defineComponent({
        components: {
            TeamOutlined,
            FireOutlined,
            BarChartOutlined,
            CreditCardOutlined,
            AppstoreOutlined,
            BulbOutlined,
            UserOutlined
        },
        data() {
            return {
                collapsed: ref<boolean>(false),
                selectedKeys: ref<string[]>(['1']),
            };
        },

        computed: {
            tabList(){
                return [this.$route.meta.module, this.$route.meta.name]
            },
        },
    });
</script>

<style>
    #components-layout-demo-side .logo {
        height: 32px;
        margin: 16px;
        background: rgba(255, 255, 255, 0.3);
    }

    .site-layout .site-layout-background {
        background: #fff;
    }
    [data-theme='dark'] .site-layout .site-layout-background {
        background: #141414;
    }
</style>