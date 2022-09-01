package main

import (
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
	userService.GetUserInfo(2)
	// user->order->db
}
