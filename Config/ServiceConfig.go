package Config

import "ioc/services"

type ServiceConfig struct{}

func NewServiceConfig() *ServiceConfig {
	return &ServiceConfig{}
}

func (sc *ServiceConfig) OrderService() *services.OrderService {
	return services.NewOrderService()
}

func (sc ServiceConfig) DBService() *services.DBService {
	return services.NewDBService()
}