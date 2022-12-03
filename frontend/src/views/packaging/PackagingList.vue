<template>
    <a-row>
        <a-button class="editable-add-btn" style="margin-bottom: 8px;" @click="handleAdd">新增</a-button>
        <Add @switch="switchIsShowAdd" :is-show-add="isShowAdd" />
    </a-row>

    <a-form-item label="包裹信息查询">
        <a-input-search
        v-model:value="searchKey"
        placeholder="请输入包裹信息或封包信息查询"
        enter-button
        @search="onSearch"
        :style="{width: '400px', float: 'left'}"
        />
    </a-form-item>
    <a-table :columns="columns" :data-source="data" :pagination="false" bordered :style="{lineHeight: '0.3'}">
        <template #bodyCell="{ column, text, record }">
            <template v-if="['name', 'borrow', 'repayment'].includes(column.dataIndex)">
                <div>
                <a-input
                    v-if="editableData[record.key]"
                    v-model:value="editableData[record.key][column.dataIndex]"
                    style="margin: -5px 0"
                />
                <template v-else>
                    {{ text }}
                </template>
                </div>
            </template>
            <template v-else-if="column.dataIndex === 'operation'">                      
                <div class="editable-row-operations">
                <span v-if="editableData[record.key]">
                    <a-typography-link @click="save(record.key)">保存</a-typography-link>
                    <a-popconfirm title="确定要取消吗?" @confirm="cancel(record.key)">
                        <a>取消</a>
                    </a-popconfirm>
                </span>
                <span v-else>
                    <a @click="edit(record.key)">编辑</a>
                    <a @click="onDelete(record.key)">删除</a>
                </span>
                </div>
            </template> 
        </template>
    </a-table>

    <br>
    <a-pagination v-model:current="current" :total="500" :style="{float: 'right'}"/>
</template>

<script lang="ts">
import { cloneDeep } from 'lodash-es';
import type { TableColumnsType } from 'ant-design-vue';
import { computed, defineComponent, ref, reactive } from 'vue';
import type { UnwrapRef } from 'vue';

import Add from './Add.vue';

export default defineComponent({
    components: {
            Add,
        },
    setup() {
        // 表单头部，和对应数据字段的key
        const columns = ref<TableColumnsType>([
            {
                title: '车辆信息',
                dataIndex: 'name',
                width: ""
            },
            {
                title: '封包信息',
                dataIndex: 'borrow',
                width: ""
            },
            {
                title: '包裹信息',
                dataIndex: 'repayment',
                width: ""
            },
            {
                title: '封包时间',
                dataIndex: 'repayment',
                width: "22%"
            },
            {
                title: '操作',
                dataIndex: 'operation',
                width: "13%"
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

        // 搜索
        const searchKey = "";
        const onSearch = () => {
            
        }

        // 新增
        const isShowAdd = ref(false)
        const handleAdd = () => {
            isShowAdd.value = true;
        };

        const switchIsShowAdd = () => {
            isShowAdd.value = !isShowAdd.value
        }


        interface DataItem {
            key: string;
            name: string;
            age: number;
        }

                // 数据
        const data = ref([
            {   key: "1",
                name: 'John Brown',
                borrow: 10,
                repayment: "2022-12-03 19:18:09",
            },
            {
                key: "2",
                name: 'John Brown',
                borrow: 10,
                repayment: "2022-12-03 19:18:09",
            },
        ]);

        const editableData: UnwrapRef<Record<string, DataItem>> = reactive({});

        const edit = (key: string) => {
                editableData[key] = cloneDeep(data.value.filter(item => key === item.key)[0]);
            };

        const save = (key: string) => {
            Object.assign(data.value.filter(item => key === item.key)[0], editableData[key]);
            delete editableData[key];
        };

        const cancel = (key: string) => {
            delete editableData[key];
        };

        const onDelete = (key: string) => {
            data.value = data.value.filter(item => item.key !== key);
        };

        const current = ref<number>(1);    
        return {
            data,
            columns,
            fixedColumns,
            fixedData,

            // 搜索
            searchKey,
            onSearch,
            // 新增
            isShowAdd,
            handleAdd,
            switchIsShowAdd,

            edit,
            save,
            cancel,
            onDelete,
            editableData,

            current,
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


    .editable-row-operations a {
        margin-right: 8px;
    }
</style>
  