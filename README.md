# go-micro-frame

go-micro-frame框架，是一套开源组件组合而成的微服务框架。

没有框架的强约束，没有学习上的成本。只需要搭建积木的方式组合自己的框架来快速开展业务。框架只保留了微服务的核心功能，使用者可以完全自主的进行改造。

文档参考：https://github.com/jettjia/go-micro-fram-doc

模块介绍：

```
gorm		【orm】
gin		【web服务】
grpc、proto	【rpc微服务】
zap 		【日志】
viper		【配置读取】
consul 		【服务注册和发现】
nacos		【配置中心】
grpc-lb 	【负载均衡】
es		【搜索】
分布式锁	        【redis实现】
幂等性		【grpc重试，分布式锁处理，token处理等】
jaeger		【链路追踪】
sentinel	【限流、熔断、降级】
kong		【网关】
amqp            【amqp，消息队列，比如：rabbitmq】
cron            【分布式定时任务;go:go-cron,java:xxl-job】
分布式事务	【方式1：rocketmq，事务消息方式；方式2：seata-golang,暂缺】
分布式mysql	【go: gaea分库分表; java: shardingsphere-proxy】
```

版本说明：

```
v1.1.1 完成
已经实现了模块介绍中的功能
```

```
v1.2.0 开发中
会把microframe.com里的模块全部实现，进行统一维护，方便更新和迭代替换。
nacos：已封装
logger: 已封装
jaeger: 已封装
mysql：已封装，连接池方式
```

```
v1.3.0 开发中
会增加k8s的自动化发布脚本，Prometheus、Grafana监控等
```

```
v1.4.0 规划
会改造到istio或者dapr的三代微服务方式中，会单独另起一个项目进行维护
```

