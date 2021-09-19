module go-micro-frame

go 1.16

require (
	github.com/armon/go-metrics v0.3.2 // indirect
	github.com/gin-gonic/gin v1.7.4
	github.com/go-redis/redis/v8 v8.11.3
	github.com/go-redsync/redsync/v4 v4.3.0
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/grpc-opentracing v0.0.0-20180507213350-8e809c8a8645
	github.com/hashicorp/consul/api v1.10.1
	github.com/hashicorp/go-immutable-radix v1.1.0 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/mattn/go-isatty v0.0.13 // indirect
	github.com/olivere/elastic/v7 v7.0.29
	github.com/opentracing/opentracing-go v1.2.0
	github.com/satori/go.uuid v1.2.0
	github.com/spf13/viper v1.8.1
	github.com/streadway/amqp v1.0.0
	github.com/stretchr/testify v1.7.0
	github.com/uber/jaeger-client-go v2.29.1+incompatible
	go.uber.org/zap v1.19.1
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5 // indirect
	golang.org/x/net v0.0.0-20210805182204-aaa1db679c0d
	google.golang.org/grpc v1.40.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gorm.io/gorm v1.21.13
	microframe.com/consul v0.0.0-00010101000000-000000000000
	microframe.com/jaeger v0.0.0-00010101000000-000000000000
	microframe.com/logger v0.0.0-00010101000000-000000000000
	microframe.com/mysql v0.0.0-00010101000000-000000000000
	microframe.com/nacos v0.0.0-00010101000000-000000000000
)

replace (
	microframe.com/consul => ../microframe.com/consul
	microframe.com/jaeger => ../microframe.com/jaeger
	microframe.com/logger => ../microframe.com/logger
	microframe.com/mysql => ../microframe.com/database/mysql
	microframe.com/nacos => ../microframe.com/nacos
)
