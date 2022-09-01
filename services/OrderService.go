package services

import "fmt"

type IOrder interface {
	Name() string
}

type OrderService struct {
	Version string
	DB      *DBService `inject:"-"`
}

func NewOrderService() *OrderService {
	return &OrderService{Version: "1.0"}
}

func (orderService *OrderService) GetOrderInfo(uid int) {
	fmt.Println("获取用户id=", uid, "的订单信息")
}

func (orderService *OrderService) Name() string {
	return "orderService"
}
