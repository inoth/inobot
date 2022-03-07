# inobot
```
inobot run --topic="test_topic" --replica=1
```

#### 迭代规划
* 日志
    * 方案A：nsq hook 写入 es
    * 方案B：直接写到文件中，使用filebeat采集
* 脚本工具库
    * http 请求
    * 通用加密 / 解密
    * go to lua 模块，具体需要封装内容暂定
* 安全
    * 脚本加密
    * 执行鉴权
* 其他
    * 脚本功能类型分类
    * 根据 tag 分流权重
    * 配置中心分发脚本 / 热更拉取配置中心
