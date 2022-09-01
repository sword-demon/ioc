package main

import (
	"fmt"
	. "ioc/Injector"
	"ioc/services"
)

func main() {
	//uid := 123
	//userService := services.NewUserService(services.NewOrderService())
	//userService.GetUserInfo(uid)
	//userService.GetOrderInfo(uid)

	BeanFactory.Set(services.NewOrderService())

	userService := services.NewUserService()
	BeanFactory.Apply(userService)
	fmt.Println(userService.Order)
}
