package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"go.uber.org/zap"
)

// 直接读取 nacos 的配置信息
func GetClientContent(host string, port uint64, namespace string, user string, password string, dataid string, group string) (content string, err error) {
	//从nacos中读取配置信息
	configClient, err := GetConfigClient(host, port, namespace, user, password)
	if err != nil {
		return "", err
	}

	content, err = configClient.GetConfig(vo.ConfigParam{
		DataId: dataid,
		Group:  group})

	if err != nil {
		return "", err
	}
	zap.S().Infof("从nacos读取到的全部配置如下：", content)
	return content, nil
}

// 获取nacos的 IConfigClient
func GetConfigClient(host string, port uint64, namespace string, user string, password string) (config_client.IConfigClient, error) {
	//从nacos中读取配置信息
	sc := []constant.ServerConfig{
		{
			IpAddr: host,
			Port:   port,
		},
	}
	cc := constant.ClientConfig{
		NamespaceId:         namespace, // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "debug",
		Username:            user,
		Password:            password,
	}
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		return nil, err
	}

	return configClient, nil
}
