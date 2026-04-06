// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"grpc-common/ucenter/ucclient"
	"ucenter-api/internal/config"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	UCRegisterRpc ucclient.RegisterService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UCRegisterRpc: ucclient.NewRegisterService(zrpc.MustNewClient(c.UCenterRpc)),
	}
}
