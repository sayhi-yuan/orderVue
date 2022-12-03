<template>
    <a-drawer
    title="新增包裹信息"
    placement="right"
    :closable="false"
    :visible="isShowAdd"
    :style="{ position: 'absolute' }"
    >
        <!-- 表单 -->
        <a-form :model="formState" :label-col="labelCol" :wrapper-col="wrapperCol">
            <a-form-item label="车辆信息">
                <a-input v-model:value="formState.name" />
            </a-form-item>
            <a-form-item label="封包信息">
                <a-input v-model:value="formState.address" type="textarea" />
            </a-form-item>
            <a-form-item label="包裹信息">
                <a-input v-model:value="formState.address" type="textarea" />
            </a-form-item>
            <a-form-item :wrapper-col="{ span: 14, offset: 4 }">
                <a-button type="primary" @click="onSave">保存</a-button>
                <a-button style="margin-left: 10px" @click="onClose">取消</a-button>
            </a-form-item>
        </a-form>
    </a-drawer>
</template>

<script lang="ts">
    import { defineComponent, reactive, toRaw } from 'vue';
    import type { UnwrapRef } from 'vue';
    import { message } from 'ant-design-vue';
  
    interface FormState {
        name: string;
        address: string;
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
                });

            return {
                afterVisibleChange,
                showDrawer,

                // 新增商家
                formState,
                labelCol: { style: { width: '75px' } },
                wrapperCol: { span: 18 },
            };
        },
    });
</script>

 
  
  