<template>
    <div class="container">
        <div class="loginBox">
            <a-form 
                class="loginForm"
                ref="loginFormRef"
                :label-col="labelCol"
                :model="formData" 
                :rules="rules"
            >
                <h1><b>登录界面</b></h1>
                <!-- 用户名输入 -->
                 <a-form-item label='用户名' name="username">
                    <a-input v-model:value="formData.username" placeholder="请输入用户名...">
                    </a-input>
                </a-form-item>

                <!-- 密码输入 -->
                <a-form-item label='密码' name="password">
                    <a-input-password v-model:value="formData.password" placeholder="请输入密码...">
                    </a-input-password>
                </a-form-item>
                
                <!-- 登陆按钮 -->
                <a-button 
                    type="primary"
                    class="login-form-button"
                    @click="login">登陆</a-button>
            </a-form>
        </div>
    </div>

    <!-- 测试 -->
    <div v-if="false">
        <!-- 使用一个Ant组件中的按钮 -->
        <a-button type="primary">Primary Button</a-button>
    </div>
</template>

<script>
import { message } from 'ant-design-vue'


    export default {
        data() {
            return  {
                // 添加表单元素
                formData: {
                    username: '',
                    password: '',
                },
                // 通过label-col来控制labe宽度
                labelCol: {style:{width:'4.5em'}},
                // 设置表单输入的规则
                rules: {
                    username: [
                        { 
                            required:true, 
                            message: '请输入用户名'},
                        { 
                            min: 4,
                            max: 12,
                            message: '用户名必须在4到12个字符之间',
                        },
                    ],
                    password: [
                        { 
                            required:true, 
                            message: '请输入与密码'
                        },
                        {
                            min: 6,
                            max: 20,
                            message: '密码必须在6到20个字符之前',
                        }
                    ]
                }
            }
        },
        methods:{
            // 输入函数
            login: function(){
                const result = this.$refs.loginFormRef.validate()

                result.then(()=>{
                    console.log("ok")
                }).catch((errEsg)=>{
                    this.$message.error("非法输入，请重新输入！")
                })
            }
        }
    }
</script>

<style>
.container {
    height: 100%;
    background-color: #1e1e1e;
}

.loginBox{
    width: 400px;
    height: 300px;

    background-color: white;
    
    position: absolute;
    top: 50%;
    left: 70%;
    transform: translate(-50%, -50%);
    
    margin: 5px;
    padding: 20px;
    border-radius: 10px;
}
.login-form-button{
    width: 100%;
}
</style>