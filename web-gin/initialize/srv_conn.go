package initialize

import (
	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"go-micro-frame-web/global"
	"go-micro-frame-web/proto"
	"go.uber.org/zap"
	"microframe.com/consul"
)

func InitSrvConn() {
	consulInfo := global.ServerConfig.ConsulInfo
	registerClient := consul.NewRegistryClient(global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)
	userConn, err := registerClient.Discovery(consulInfo.Host, consulInfo.Port, global.ServerConfig.UserSrvInfo.Name)
	if err != nil {
		zap.S().Fatal("[InitSrvConn] 连接 【用户服务失败】")
	}

	userSrvClient := proto.NewUserClient(userConn)
	global.UserSrvClient = userSrvClient
}
