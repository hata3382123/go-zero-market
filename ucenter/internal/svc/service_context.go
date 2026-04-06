package svc

import (
	"database/sql"
	"mscoin-common/msdb"
	"ucenter/internal/config"
	"ucenter/internal/database"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/syncx"
)

type ServiceContext struct {
	Config config.Config
	Cache  cache.Cache
	DB     *msdb.MsDB
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisCache := cache.New(c.StoreRedis, syncx.NewSingleFlight(), cache.NewStat("mscoin"), sql.ErrNoRows)
	return &ServiceContext{
		Config: c,
		Cache:  redisCache,
		DB:     database.ConnMysql(c.Mysql.DataSource),
	}
}
