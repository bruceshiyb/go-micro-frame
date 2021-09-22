package main

import (
	"fmt"

	"os"
	"os/signal"
	"syscall"

	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"

	"go-micro-frame-web/global"
	"go-micro-frame-web/initialize"
	"microframe.com/consul"
	mylogger "microframe.com/logger"
	"microframe.com/publicUtil"
)

func main() {
	// 初始化web
	initialize.InitWeb()

	// 初始化routers
	Router := initialize.Routers()

	/////////////////////////////////////////////
	// 随机生成 port, 如果是本地开发环境端口号固定，线上环境启动获取端口号
	if global.ServerConfig.Env != "dev" {
		global.ServerConfig.Port, _ = publicUtil.GetFreePort()
	}

	//注册服务健康检查
	grpcServer := grpc.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, health.NewServer())

	// 注册服务到 consul中
	registerClient := consul.NewRegistryClient(global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)
	serviceId := fmt.Sprintf("%s", uuid.NewV4())
	err := registerClient.RegisterHttp(global.ServerConfig.Host, global.ServerConfig.Port, global.ServerConfig.Name, global.ServerConfig.Tags, serviceId)
	if err != nil {
		zap.S().Panic("服务注册失败:", err.Error())
	}
	/////////////////////////////////////////////

	// 启动 web服务
	zap.S().Debugf("启动web服务的端口： %d", global.ServerConfig.Port)
	mylogger.Info("", mylogger.Int("启动web服务端口,port:", global.ServerConfig.Port))
	go func() {
		if err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
			zap.S().Panic("启动失败:", err.Error())
		}
	}()

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
