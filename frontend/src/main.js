import {createApp} from 'vue'
import Antd from 'ant-design-vue';
import 'ant-design-vue/dist/antd.css';
import App from './App.vue'

import Route from './route/route'

const app = createApp(App);

app.use(Route).use(Antd).mount('#app');
