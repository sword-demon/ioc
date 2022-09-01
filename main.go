package main

import (
	"fmt"
	"ioc/Config"
	. "ioc/Injector"
	"ioc/services"
)

func main() {
	//uid := 123
	//userService := services.NewUserService(services.NewOrderService())
	//userService.GetUserInfo(uid)
	//userService.GetOrderInfo(uid)

	serviceConfig := Config.NewServiceConfig()
	//BeanFactory.ExprMap = map[string]interface{}{
	//	"ServiceConfig": serviceConfig,
	//}
	BeanFactory.Set(serviceConfig)

	userService := services.NewUserService()
	BeanFactory.Apply(userService)
	fmt.Println(userService.Order)
}
