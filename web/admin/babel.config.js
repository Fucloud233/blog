module.exports = {
  presets: [
    '@vue/cli-plugin-babel/preset'
  ],

  // 使用babel加载Ant组件
  "plugins": [
    ["import", { "libraryName": "ant-design-vue", "libraryDirectory": "es", "style": "css" }] // `style: true` 会加载 less 文件
  ]
  
}
