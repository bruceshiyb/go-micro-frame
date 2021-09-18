# go-micro-frame-doc

go-micro-frame框架，是一套开源组件组合而成的微服务框架。

没有框架的强约束，没有学习上的成本。只需要搭建积木的方式组合自己的框架来快速开展业务。框架只保留了微服务的核心功能，使用者可以完全自主的进行改造。

文档参考：https://github.com/jettjia/go-micro-fram-doc

v1.0.0 已经完成，已经实现了下面介绍的模块

v1.1.0 开发中，将一些公用的模块进行再封装和统一管理，方便快速的进行模块替换。

```
gorm		【orm】
gin		【web服务】
grpc、proto	【rpc微服务】
zap 		【日志】
viper		【配置读取】
consul 		【服务注册和发现】
nacos		【配置中心】
grpc自带	【负载均衡】
es		【搜索】
分布式锁	        【redis实现】
幂等性		【grpc重试，分布式锁处理，token处理等】
jaeger		【链路追踪】
sentinel	【限流、熔断、降级】
kong		【网关】
amqp            【amqp，消息队列，比如：rabbitmq】
cron            【分布式定时任务】
分布式事务	【rocketmq，事务消息方式；方式2：seata-golang】
分布式mysql	【gaea分库分表】
```

