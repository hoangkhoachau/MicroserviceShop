package api

import (
	ginserver "github.com/hoangkhoachau/MicroserviceShop/basket/api/gin-server"
	"github.com/hoangkhoachau/MicroserviceShop/basket/config"
	"github.com/hoangkhoachau/MicroserviceShop/basket/internal/factories"
)

type ApiFactory struct {
	cfg            config.SystemConfig
	serviceFactory factories.ServiceFactory
}

func NewApiFactory(cfg config.SystemConfig, serviceFactory factories.ServiceFactory) *ApiFactory {
	return &ApiFactory{cfg: cfg, serviceFactory: serviceFactory}
}

func (a *ApiFactory) GetApi() IApi {
	switch a.cfg.Server {
	case "gin":
		return ginserver.NewGinServer(a.serviceFactory, a.cfg)
	}
	return nil
}
