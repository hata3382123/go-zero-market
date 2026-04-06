package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql      MysqlConfig     `json:",optional"`
	StoreRedis cache.CacheConf `json:"StoreRedis,optional"` // conf 按 json 标签解析，勿只写 yaml 标签否则字段会被跳过
	Captcha    CaptchaConf
}
type MysqlConfig struct {
	DataSource string `json:",optional"`
}
type CaptchaConf struct {
	Vid string
	Key string
}
