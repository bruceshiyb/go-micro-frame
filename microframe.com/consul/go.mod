module consul

go 1.16

require (
	github.com/hashicorp/consul/api v1.10.1
	github.com/opentracing/opentracing-go v1.2.0
	github.com/satori/go.uuid v1.2.0
	go.uber.org/zap v1.19.1
	google.golang.org/grpc v1.40.0
	microframe.com/otgrpc v0.0.0-00010101000000-000000000000
)

replace microframe.com/otgrpc => ../otgrpc
