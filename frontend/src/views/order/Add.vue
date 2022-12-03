<template>
    <a-drawer
    title="新增订单"
    placement="right"
    :closable="false"
    :visible="isShowAdd"
    :style="{ position: 'absolute' }"
    >
        <a-form :model="formState" :label-col="labelCol" :wrapper-col="wrapperCol">
            <a-form-item label="供方货号">
                <a-select
                    ref="select"
                    v-model:value="formState.name"
                    >
                    <a-select-option value="1">商家1号</a-select-option>
                    <a-select-option value="2">商家2号</a-select-option>
                    <a-select-option value="3">商家3号</a-select-option>
                    <a-select-option value="4">商家4号</a-select-option>
                </a-select>
            </a-form-item>
            <a-form-item label="前台订单">
                <a-input v-model:value="formState.name" />
            </a-form-item>
            <a-form-item label="商家名称">
                <a-select
                    ref="select"
                    v-model:value="formState.name"
                    >
                    <a-select-option value="1">商家1号</a-select-option>
                    <a-select-option value="2">商家2号</a-select-option>
                    <a-select-option value="3">商家3号</a-select-option>
                    <a-select-option value="4">商家4号</a-select-option>
                </a-select>
            </a-form-item>
            <a-form-item label="商家货号">
                <a-select
                    ref="select"
                    v-model:value="formState.name"
                    >
                    <a-select-option value="1">商家1号</a-select-option>
                    <a-select-option value="2">商家2号</a-select-option>
                    <a-select-option value="3">商家3号</a-select-option>
                    <a-select-option value="4">商家4号</a-select-option>
                </a-select>
            </a-form-item>
            <a-form-item label="商家价格">
                <a-input v-model:value="formState.name" />
            </a-form-item>
            <a-form-item label="商家地址">
                <a-input v-model:value="formState.address" type="textarea" />
            </a-form-item>

            <!-- 添加选项 -->
            <a-form-item label="商家颜色">
                <a-select
                    ref="select"
                    v-model:value="formState.name"
                    >
                    <a-select-option value="1">商家1号</a-select-option>
                    <a-select-option value="2">商家2号</a-select-option>
                    <a-select-option value="3">商家3号</a-select-option>
                    <a-select-option value="4">商家4号</a-select-option>
                </a-select>
            </a-form-item>
            <a-form-item label="尺码">
                <a-input v-model:value="formState.name" />
            </a-form-item>
            <a-form-item label="数量">
                <a-input v-model:value="formState.name" />
            </a-form-item>

            <a-form-item label="">
                <a-button type="primary" style="margin-right: 12px" @click="onSave">添加选项</a-button>
                <a-button type="primary" style="margin-right: 12px" @click="onSave">保存</a-button>
                <a-button @click="onClose">取消</a-button>
            </a-form-item>
        </a-form>
    </a-drawer>
</template>

<script lang="ts">
    import { defineComponent, reactive } from 'vue';
    import type { UnwrapRef } from 'vue';
    import { message } from 'ant-design-vue';
  
    interface FormState {
        name: string;
        address: string;

        color: number;
    }

    export default defineComponent({
        props: {
            isShowAdd: {
                type: Boolean,
            },
        },

        methods: {
            onClose(){
                this.$emit('switch');
                // 同时清空数据
            },

            onSave(){
                this.$emit('switch');
                message.success('添加成功');
                message.error('添加失败');
                // 同时清空数据
            }
        },

        setup() {
            const afterVisibleChange = (bool: boolean) => {
            };

            const showDrawer = () => {
            // props.isShowAddWindow = true;
            };


            const formState: UnwrapRef<FormState> = reactive({
                    name: '',
                    address: '',
                    color: 0,
                });

            return {
                afterVisibleChange,
                showDrawer,

                // 新增商家
                formState,
                labelCol: { style: { width: '80px' } },
                wrapperCol: { span: 18 },
            };
        },
    });
</script>

 
  
  