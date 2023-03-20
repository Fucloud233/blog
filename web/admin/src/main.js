import { createApp, VueElement } from 'vue'
import App from './App.vue'
import router from './router'
import axios from 'axios'
import { Button, Form, Input } from 'ant-design-vue'

// 创建一个Vue示例
const app = createApp(App)

// 将axios挂载在Vue上
axios.defaults.baseURL = 'http://localhost:3000/api/v1'
VueElement.prototype.$http = axios

// 加载组件
app.use(router)
app.use(Button)
app.use(Form)
app.use(Input)

// 将Vue搭载在app上
app.mount('#app')
