<template>
    <div>
        <!-- 使用一个Ant组件中的按钮 -->
        <a-button type="primary">Primary Button</a-button>
    </div>
</template>

<script>
    export default {
        data() {
            return  {
                // 添加表单元素
                formData: {
                    username: '',
                    password: '',
                }
            }
        }
    }
</script>