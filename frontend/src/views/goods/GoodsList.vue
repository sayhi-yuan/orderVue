<template>
    <a-row>
        <a-button class="editable-add-btn" style="margin-bottom: 16px; float: left;" @click="addGoods">新增</a-button>
        <GoodsAdd @switch="switchIsShowAdd" :is-show-add="isShowAdd" />
    </a-row> 
    <!-- 查询 -->
    <a-row>
        <a-col :span="7">
            <a-form-item label="供方货号" name="username">
                <a-input>
                    <UserOutlined class="site-form-item-icon" />
                </a-input>
            </a-form-item>
        </a-col>
        <a-col :span="1"></a-col>
        <a-col :span="7">
            <a-form-item label="商家货号" name="password">
                <a-input>
                    <LockOutlined class="site-form-item-icon" />
                </a-input>
            </a-form-item>
        </a-col>
        <a-col :span="5"></a-col>
        <a-col :span="1">
            <a-form-item>
                <a-button type="primary">查询</a-button>
            </a-form-item>
        </a-col>
    </a-row>

    <a-table :columns="columns" :data-source="data" :pagination="false" bordered :style="{lineHeight: '0.3'}">
        <template #summary>
        </template>
    </a-table>
    <br/>
    <a-pagination
      v-model:current="current"
      v-model:page-size="pageSize"
      :total="10000"
      :show-total="total => `共 ${total} 条`"
      :style="{float: 'right'}"
    />
</template>

<script lang="ts">
import type { TableColumnsType } from 'ant-design-vue';
import { computed, defineComponent, ref } from 'vue';

import GoodsAdd from './GoodsAdd.vue'

export default defineComponent({
    components: {
        GoodsAdd,
    },
    setup() {
        // 表单头部，和对应数据字段的key
        const columns = ref<TableColumnsType>([
            {
                title: '供方货号',
                dataIndex: 'name',
            },
            {
                title: '商家名称',
                dataIndex: 'borrow',
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
                title: '商家价格',
                dataIndex: 'repayment',
            },
            {
                title: '前台售价',
                dataIndex: 'repayment',
            },
            {
                title: '商品条码',
                dataIndex: 'repayment',
            },
            {
                title: '商家地址',
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
        const addGoods = () => {
            isShowAdd.value = true;
        }

        const current = ref<number>(1);
        const pageSize = ref<number>(20);
        return {
            data,
            columns,
            totals,
            fixedColumns,
            fixedData,

            current,
            pageSize,

            isShowAdd,
            switchIsShowAdd,

            // 添加商品
            addGoods,
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
  