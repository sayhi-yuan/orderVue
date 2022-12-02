<template>
    <a-button class="editable-add-btn" style="margin-bottom: 8px; float: left;" @click="handleAdd">新增</a-button>
    <MerchantsAdd @switch="switchIsShowAdd" :is-show-add="isShowAdd" />

    <a-table :columns="columns" :data-source="dataSource" bordered :style="{lineHeight: '0.3'}">
        <template #bodyCell="{ column, text, record }">
            <template v-if="['name', 'age', 'address'].includes(column.dataIndex)">
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
</template>


<script lang="ts">
    import { cloneDeep } from 'lodash-es';
    import { defineComponent, isShallow, reactive, ref } from 'vue';
    import type { UnwrapRef } from 'vue';

    import MerchantsAdd from './MerchantsAdd.vue';

    const columns = [
        {
            title: '商家名称',
            dataIndex: 'name',
            width: '25%',
        },
        {
            title: '商家地址',
            dataIndex: 'age',
            width: '45%',
        },
        {
            title: '操作',
            dataIndex: 'operation',
            width: '11%',
        },
    ];

    interface DataItem {
        key: string;
        name: string;
        age: number;
    }
    const data: DataItem[] = [];
    for (let i = 0; i < 100; i++) {
        data.push({
            key: i.toString(),
            name: `Edrward ${i}`,
            age: 32,
        });
    }

    export default defineComponent({
        components: {
            MerchantsAdd,
        },

        setup() {
            const dataSource = ref(data);
            const editableData: UnwrapRef<Record<string, DataItem>> = reactive({});
            const isShowAdd = ref(false)

            const edit = (key: string) => {
                    editableData[key] = cloneDeep(dataSource.value.filter(item => key === item.key)[0]);
                };

            const save = (key: string) => {
                Object.assign(dataSource.value.filter(item => key === item.key)[0], editableData[key]);
                delete editableData[key];
            };

            const cancel = (key: string) => {
                delete editableData[key];
            };

            const onDelete = (key: string) => {
                dataSource.value = dataSource.value.filter(item => item.key !== key);
            };

            const handleAdd = () => {
                isShowAdd.value = true;
            };

            const switchIsShowAdd = () => {
                isShowAdd.value = !isShowAdd.value
            }

            return {
                dataSource,
                columns,
                editingKey: '',
                editableData,

                // 是否展示新增窗口
                isShowAdd,

                edit,
                save,
                cancel,
                onDelete,
                handleAdd,
                switchIsShowAdd,
            };
        },
    });
</script>

<style scoped>
    .editable-row-operations a {
        margin-right: 8px;
    }
</style>
  