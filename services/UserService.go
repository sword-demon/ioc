package services

import "fmt"

type UserService struct {
	Order *OrderService `inject:"-"`
}

func NewUserService() *UserService {
	return &UserService{}
}

func (userService *UserService) GetUserInfo(uid int) {
	fmt.Println("用户id=", uid)
}

//func (userService *UserService) GetOrderInfo(uid int) {
//	userService.Order.GetOrderInfo(uid)
//}
