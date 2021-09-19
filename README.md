# go-micro-frame

go-micro-frame框架，是一套开源组件组合而成的微服务框架。

没有框架的强约束，没有学习上的成本。只需要搭建积木的方式组合自己的框架来快速开展业务。框架只保留了微服务的核心功能，使用者可以完全自主的进行改造。

文档参考：https://github.com/jettjia/go-micro-fram-doc

项目介绍
```
grpc            【提供grpc接口】
web-gin         【go实现web接口，gin调用grpc服务，提供http接口】
web-hyperf      【php实现web接口，hyperf调用grpc服务，提供http接口】 
web-springboot  【java实现web接口，springboot调用grpc服务，提供http接口】
```

运行项目
```
1.运行 grpc 服务，里面有代码说明和使用说明
2.运行 web-gin 服务，这里实现了调用 grpc服务，里面有代码说明和使用说明
```

组件模块介绍：

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
v1.1.X 完成
已经实现了模块介绍中的功能
```

```
v1.2.X 开发中
会把microframe.com里的模块全部实现，进行统一维护，方便更新和迭代替换。
nacos：已封装
logger: 已封装
jaeger: 已封装
mysql：已封装，连接池方式
consul-register: 已封装
```

```
v1.3.X 开发中
会增加k8s的自动化发布脚本，Prometheus、Grafana监控等
```

```
v1.4.X 规划
web层多实现
web-springboot: java版本的web, 利用springboot框架调用grpc服务
web-hyperf: php版本的web，利用 hyperf框架调用grpc服务
```

```
v2.0 规划
会改造到istio或者dapr的三代微服务方式中，会单独另起一个项目进行维护
```

