## Ant表单问题
在ant表中，我们可以通过设置`rule`的值来控制输入验证。但在确认提交时，我们需要得到验证器的验证结果。

在项目源代码中，是这样处理验证结果的（这里省略了部分中间代码）。通过向表单函数`validate`里面传入一个回传函数，就能对结果进行处理。
```javascript
login() {
    this.$refs.loginFormRef.validate(async (valid) => {
    if (!valid) 
        return this.$message.error('输入非法数据，请重新输入')

    this.$router.push('/index')
    })
},
```
但是在复现过程中，我们发现浏览器会提示错误
> Warning: [antdv: Form] 
> validateFields/validateField/validate not support callback, please use promise instead

通过错误提示，我们的值现在这个函数不支持传入回调函数了。在一篇博客中，我得知了该函数的两种用法。
> 表单验证功能之 validate 方法：对症表单进行校验的方法
> 
> 1. 若传入一个回调函数，该回调函数会在校验结束后被调用，并传入两个参数：是否校验成功和未通过校验的字段。
> 1. 若不传入回调函数，则会返回一个 promise。
> 
> 引用链接：https://blog.csdn.net/WuLex/article/details/127823251

既然前面一个方法会报错，就只能时候后面一个方法

根据javascript的相关文档，我们发现Promise有三个状态
> * 待定（pending）：初始状态，既没有被兑现，也没有被拒绝。
> * 已兑现（fulfilled）：意味着操作成功完成。
> * 已拒绝（rejected）：意味着操作失败。
> 
> 引用链接: https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Promise#%E9%9D%99%E6%80%81%E6%96%B9%E6%B3%95

而对这些状态也对应着3个实例函数

``` javascript
Promise.prototype.catch()
Promise.prototype.then()
Promise.prototype.finally()
``` 

因此上述代码就可以修改为，在浏览器中就不会报错了。
``` javascript
const result = this.$refs.loginFormRef.validate()
        
result.then(()=>{
    console.log("ok")
}).catch(()=>{
    console.log("error")
})
```


