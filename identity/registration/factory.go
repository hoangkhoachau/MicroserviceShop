package registration

import (
	"github.com/hoangkhoachau/MicroserviceShop/identity/config"
	"github.com/hoangkhoachau/MicroserviceShop/identity/registration/consul"
)

type registrationFactory struct {
	appConfig config.AppConfig
}

func NewRegistrationFactory(appConfig config.AppConfig) *registrationFactory {
	return &registrationFactory{appConfig: appConfig}
}

func (r *registrationFactory) GetRegistrationService() IServiceRegistration {
	switch r.appConfig.System.ServiceDiscovery {
	case "consul":
		return consul.NewConsul(r.appConfig)
	}
	return nil
}
