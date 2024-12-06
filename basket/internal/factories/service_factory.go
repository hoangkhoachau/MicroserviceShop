package factories

import (
	service "github.com/hoangkhoachau/MicroserviceShop/basket/internal/app/service"
	serviceImp "github.com/hoangkhoachau/MicroserviceShop/basket/internal/infrastructure/service"
)

type ServiceFactory struct {
	eventBusFactory   EventBusFactory
	repositoryFactory RepositoryFactory
}

func NewServiceFactory(repositoryFactory RepositoryFactory, eventBusFactory EventBusFactory) *ServiceFactory {
	return &ServiceFactory{repositoryFactory: repositoryFactory, eventBusFactory: eventBusFactory}
}

func (s *ServiceFactory) GetBasketService() service.IBasketService {
	return serviceImp.NewBasketService(s.repositoryFactory.GetBasketRepository(), s.eventBusFactory.GetEventBus())
}
