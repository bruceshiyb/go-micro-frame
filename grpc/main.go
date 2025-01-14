package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/opentracing/opentracing-go"
	"github.com/satori/go.uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"

	"go-micro-frame/domain/repository"
	service2 "go-micro-frame/domain/service"
	"go-micro-frame/global"
	"go-micro-frame/handler"
	"go-micro-frame/initialize"
	"go-micro-frame/proto"
	"microframe.com/consul"
	mylogger "microframe.com/logger"
	"microframe.com/otgrpc"
	"microframe.com/publicUtil"
)

func main() {
	// 判断是否生成 随机的 微服务端口号
	IP := flag.String("ip", "0.0.0.0", "ip地址")
	Port := flag.Int("port", 0, "端口")
	flag.Parse()

	zap.S().Info("port: ", *Port)

	// 初始化
	initialize.InitSrv()

	*Port = int(global.ServerConfig.Port)
	if global.ServerConfig.Env != "dev" {
		*Port, _ = publicUtil.GetFreePort()
	}

	zap.S().Info("main函数：", global.ServerConfig)

	// 启动grpc，并注册到 consul，并且使用 jaeger
	server := grpc.NewServer(grpc.UnaryInterceptor(otgrpc.OpenTracingServerInterceptor(opentracing.GlobalTracer())))

	// 创建实例
	userService := service2.NewUserService(repository.NewUserRepository())

	proto.RegisterUserServer(server, &handler.UserServer{UserService: userService})
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	//注册服务健康检查
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	//启动服务
	go func() {
		err = server.Serve(lis)
		if err != nil {
			panic("failed to start grpc:" + err.Error())
		}
	}()

	//服务注册
	registerClient := consul.NewRegistryClient(global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)
	serviceId := fmt.Sprintf("%s", uuid.NewV4())
	err = registerClient.Register(global.ServerConfig.Host, *Port, global.ServerConfig.Name, global.ServerConfig.Tags, serviceId)
	if err != nil {
		zap.S().Panic("服务注册失败:", err.Error())
	}
	zap.S().Debugf("启动grpc服务端口： %d", *Port)
	mylogger.Info("", mylogger.Int("启动grpc服务端口,port:", *Port))
	/////////////////////////////////

	//接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err = registerClient.DeRegister(serviceId); err != nil {
		zap.S().Info("注销失败:", err.Error())
	} else {
		zap.S().Info("注销成功")
	}

}
