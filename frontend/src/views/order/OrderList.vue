<template>
    <a-button class="editable-add-btn" style="margin-bottom: 16px; float: left" @click="addOrder">新增订单</a-button>
    <Add @switch="switchIsShowAdd" :is-show-add="isShowAdd" />

    <a-button type="primary" :style="{float: 'right'}">
        <template #icon>
        <DownloadOutlined />
        </template>
        导出今日订单
    </a-button>
    <a-button type="primary" :style="{marginRight: '25px', float: 'right'}">
        <template #icon>
        <DownloadOutlined />
        </template>
        导出订单
    </a-button>

    <a-table :columns="columns" :data-source="data" :pagination="false" bordered :style="{lineHeight: '0.3'}">
        <template #summary>
        </template>
    </a-table>
</template>

<script lang="ts">
import type { TableColumnsType } from 'ant-design-vue';
import { computed, defineComponent, ref } from 'vue';
import { DownloadOutlined } from '@ant-design/icons-vue';

import Add from './Add.vue'

export default defineComponent({
    components: {
        Add,
        DownloadOutlined,
    },
    setup() {
        // 表单头部，和对应数据字段的key
        const columns = ref<TableColumnsType>([
            {
                title: '供方货号',
                dataIndex: 'name',
            },
            {
                title: '前台订单',
                dataIndex: 'borrow',
            },
            {
                title: '商家名称',
                dataIndex: 'repayment',
            },
            {
                title: '商家货号',
                dataIndex: 'repayment',
            },
            {
                title: '商家颜色',
                dataIndex: 'repayment',
            },
            {
                title: '商家数量',
                dataIndex: 'repayment',
            },
            {
                title: '商家地址',
                dataIndex: 'repayment',
            },
            {
                title: '订单条码',
                dataIndex: 'repayment',
            },
        ]);

        // 数据
        const data = ref([
            {
                name: 'John Brown',
                borrow: 10,
                repayment: 33,
            },
            {
                name: 'John Brown',
                borrow: 10,
                repayment: 33,
            },
        ]);

        const fixedColumns = ref<TableColumnsType>([
            {
                title: 'Name',
                dataIndex: 'name',
                fixed: true,
                width: 100,
            },
            {
                title: 'Description',
                dataIndex: 'description',
            },
        ]);

        const fixedData = ref<{ key: number; name: string; description: string }[]>([]);
        for (let i = 0; i < 20; i += 1) {
            fixedData.value.push({
                key: i,
                name: ['Light', 'Bamboo', 'Little'][i % 3],
                description: 'Everything that has a beginning, has an end.',
            });
        }

        const totals = computed(() => {
            let totalBorrow = 0;
            let totalRepayment = 0;

            data.value.forEach(({ borrow, repayment }) => {
                totalBorrow += borrow;
                totalRepayment += repayment;
            });
            
            return { totalBorrow, totalRepayment };
        });

        const isShowAdd = ref(false)

        const switchIsShowAdd = () => {
            isShowAdd.value = !isShowAdd.value
        }

        // 添加商品
        const addOrder = () => {
            isShowAdd.value = true;
        }

        return {
            data,
            columns,
            totals,
            fixedColumns,
            fixedData,

            // 新增
            isShowAdd,
            switchIsShowAdd,
            addOrder,
        };
    },
});
</script>

<style>
    #components-table-demo-summary tfoot th,
    #components-table-demo-summary tfoot td {
        background: #fafafa;
    }

    [data-theme='dark'] #components-table-demo-summary tfoot th,
    [data-theme='dark'] #components-table-demo-summary tfoot td {
        background: #1d1d1d;
    }
</style>
  