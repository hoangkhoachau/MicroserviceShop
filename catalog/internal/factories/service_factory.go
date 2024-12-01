package factories

import (
	service "github.com/hoangkhoachau/MicroserviceShop/catalog/internal/app/service"
	serviceImp "github.com/hoangkhoachau/MicroserviceShop/catalog/internal/infrastructure/service"
)

type ServiceFactory struct {
	repositoryFactory RepositoryFactory
}

func NewServiceFactory(repositoryFactory RepositoryFactory) *ServiceFactory {
	return &ServiceFactory{repositoryFactory: repositoryFactory}
}

func (s *ServiceFactory) GetProductService() service.IProductService {
	return serviceImp.NewProductService(s.repositoryFactory.GetProductRepository())
}
