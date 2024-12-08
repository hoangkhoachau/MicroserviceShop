package api

import (
	ginserver "github.com/hoangkhoachau/MicroserviceShop/api-gateway/api/gin-server"
	"github.com/hoangkhoachau/MicroserviceShop/api-gateway/config"
)

type apiFactory struct {
	routesConfig config.RoutesConfig
	api          IApi
}

func NewApiFactory(routesConfig config.RoutesConfig) *apiFactory {
	return &apiFactory{routesConfig: routesConfig}
}

func (a *apiFactory) GetApi() IApi {
	switch a.routesConfig.Server {
	case "gin":
		return ginserver.NewGinServer(a.routesConfig)
	}
	return nil
}
