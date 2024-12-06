package factories

import (
	service "github.com/hoangkhoachau/MicroserviceShop/order/internal/app/service"
	serviceImp "github.com/hoangkhoachau/MicroserviceShop/order/internal/infrastructure/service"
)

type ServiceFactory struct {
	eventBusFactory   EventBusFactory
	repositoryFactory RepositoryFactory
}

func NewServiceFactory(repositoryFactory RepositoryFactory, eventBusFactory EventBusFactory) *ServiceFactory {
	return &ServiceFactory{repositoryFactory: repositoryFactory, eventBusFactory: eventBusFactory}
}

func (s *ServiceFactory) GetOrderService() service.IOrderService {
	return serviceImp.NewOrderService(s.repositoryFactory.GetOrderRepository(), s.eventBusFactory.GetEventBus())
}
