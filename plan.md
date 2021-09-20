# 开发版本 v1.2.X

已经完成模块
```
logger
nacos
jaeger
database/mysql
consul-register
consul-discovery
snowflake
rabbitMq
elasticSearch
sentinel
```

进行中模块
```
redis: 不随着项目启动
redis-三种模式： standalone cluster sentinel
redisSyncLock-分布式锁


rocketMq: 不随着项目启动
kafka: 不随着项目启动
csv
```

初始化改造
```
初始化时候，增加统一的方法 initApp.go


默认只初始化：
logger: 日志
config: nacos配置读取
db: mysql初始化
jaeger: 链路追踪
缓存：redis

可选初始化:
es
rabbitmq


```




