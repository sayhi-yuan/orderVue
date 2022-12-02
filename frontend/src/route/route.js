import { createRouter, createWebHistory } from 'vue-router';

// Dashboard
import DashboardAnalysisPage from '../views/dashboard/AnalysisPage.vue';
// 营销管理
import MarketingList from '../views/marketing/MarketingList.vue'
// 订单管理
import OrderList from '../views/order/OrderList.vue'
// 商品管理
import GoodsList from '../views/goods/GoodsList.vue';
// 商家管理
import MerchantsList from '../views/merchants/MerchantsList.vue'
// 打包发货
import PackagingList from '../views/packaging/PackagingList.vue'

// 路由
const routes = [
    // Dashboard
    {
        path: '/',
        meta: {
            module: "Dashboard"
        },
        children: [
            {
                path: "",
                component: DashboardAnalysisPage,
                meta: {
                    name: "分析页"
                },
            },
        ],

    },

    // 营销管理
    {
        path: '/marketing-manage',
        meta: {
            module: "营销管理"
        },
        children: [
            {
                path: "list",
                component: MarketingList,
                meta: {
                    name: "营销列表"
                },
            },
        ],
    },

    // 订单管理
    {
        path: '/order',
        meta: {
            module: "订单管理"
        },
        children: [
            {
                path: "list",
                component: OrderList,
                meta: {
                    name: "订单列表"
                },
            },
        ],
    },

    // 商品管理
    {
        path: '/goods',
        meta: {
            module: '商品管理',
        },
        children: [
            {
                path: 'list',
                component: GoodsList,
                meta: {
                    name: '商品列表',
                },
            },
        ],
    },

    // 商家管理
    {
        path: '/merchants',
        meta: {
            module: '商家管理',
        },
        children: [
            {
                path: 'list',
                component: MerchantsList,
                meta: {
                    name: '商家列表',
                },
            },
        ],
    },

    // 打包发货
    {
        path: '/packaging',
        meta: {
            module: '打包发货',
        },
        children: [
            {
                path: 'list',
                component: PackagingList,
                meta: {
                    name: '发货列表',
                },
            },
        ],
    },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router;