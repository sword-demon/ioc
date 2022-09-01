package main

import (
	"fmt"
	"ioc/Config"
	. "ioc/Injector"
	"ioc/services"
)

func main() {

	serviceConfig := Config.NewServiceConfig()
	//BeanFactory.ExprMap = map[string]interface{}{
	//	"ServiceConfig": serviceConfig,
	//}
	BeanFactory.Config(serviceConfig)

	userService := services.NewUserService()
	BeanFactory.Apply(userService)
	fmt.Println(userService.Order.DB)
	// user->order->db
}
