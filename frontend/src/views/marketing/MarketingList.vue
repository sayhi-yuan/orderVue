<template>
    <a-row>
        <a-col :span="8">
            <a-table :columns="columns" :data-source="data" :pagination="false" bordered :style="{lineHeight: '0.3'}">
                <template #summary></template>
            </a-table>
        </a-col>
        <a-col :span="16">
            <a-table :columns="columns" :data-source="data" :pagination="false" bordered :style="{lineHeight: '0.3'}">
                <template #summary></template>
            </a-table>
        </a-col>
    </a-row>
</template>

<script lang="ts">
import type { TableColumnsType } from 'ant-design-vue';
import { computed, defineComponent, ref } from 'vue';

export default defineComponent({
    setup() {
        const columns = ref<TableColumnsType>([
            {
                title: 'Name1',
                dataIndex: 'name',
            },
            {
                title: 'Borrow',
                dataIndex: 'borrow',
            },
            {
                title: 'Repayment',
                dataIndex: 'repayment',
            },
        ]);

        const data = ref([
            {
                name: 'Brown',
                borrow: 10,
                repayment: 33,
            },
            {
                name: 'Brown',
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

        return {
            data,
            columns,
            totals,
            fixedColumns,
            fixedData,
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
  