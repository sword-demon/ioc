package services

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type UserService struct {
	Order *OrderService `inject:"-"`
	DB    *gorm.DB      `inject:"-"`
}

func NewUserService() *UserService {
	return &UserService{}
}

type UserModel struct {
	Id       int    `gorm:"column:id"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
}

func (userService *UserService) GetUserInfo(uid int) {
	user := &UserModel{}
	userService.DB.Raw("select * from user where id=?", uid).First(user)
	//fmt.Println("用户id=", uid)
	fmt.Println(user)
}

//func (userService *UserService) GetOrderInfo(uid int) {
//	userService.Order.GetOrderInfo(uid)
//}
