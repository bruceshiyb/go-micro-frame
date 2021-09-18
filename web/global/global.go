package global

import (
	ut "github.com/go-playground/universal-translator"

	"go-micro-frame-web/config"
	"go-micro-frame-web/proto"
)

var (
	Trans ut.Translator
	ServerConfig *config.ServerConfig = &config.ServerConfig{}

	UserSrvClient proto.UserClient

	NacosConfig *config.NacosConfig = &config.NacosConfig{}
)
