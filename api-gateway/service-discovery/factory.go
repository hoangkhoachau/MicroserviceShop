package servicediscovery

import (
	"github.com/hoangkhoachau/MicroserviceShop/api-gateway/config"
	"github.com/hoangkhoachau/MicroserviceShop/api-gateway/service-discovery/consul"
)

type ServiceDiscoveryFactory struct {
	cfg config.RoutesConfig
}

func NewServiceDiscoveryFactory(cfg config.RoutesConfig) *ServiceDiscoveryFactory {
	return &ServiceDiscoveryFactory{cfg: cfg}
}

func (s *ServiceDiscoveryFactory) GetServiceDiscovery() IServiceDiscovery {
	switch s.cfg.ServiceDiscovery {
	case "consul":
		return consul.NewConsul(s.cfg.Consul)
	}
	return nil
}
