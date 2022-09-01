package Injector

import (
	"github.com/shenyisyn/goft-expr/src/expr"
	"log"
	"reflect"
)

var BeanFactory *BeanFactoryImpl

func init() {
	BeanFactory = NewBeanFactoryImpl()
}

type BeanFactoryImpl struct {
	beanMapper BeanMapper
	ExprMap    map[string]interface{}
}

func (beanFactoryImpl *BeanFactoryImpl) Set(vList ...interface{}) {
	if vList == nil || len(vList) == 0 {
		return
	}
	for _, v := range vList {
		beanFactoryImpl.beanMapper.add(v)
	}
}

func (beanFactoryImpl *BeanFactoryImpl) Get(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	getV := beanFactoryImpl.beanMapper.get(v)
	if getV.IsValid() {
		return getV.Interface()
	}
	return nil
}

// Apply 处理依赖注入
func (beanFactoryImpl *BeanFactoryImpl) Apply(bean interface{}) {
	if bean == nil {
		return
	}
	// 反射
	v := reflect.ValueOf(bean)
	// 判断是否是指针类型
	if v.Kind() == reflect.Ptr {
		// 代表是指针类型
		v = v.Elem() // 等于所指向的对象
	}
	// 判断是否是一个结构体 如果是string就没必要进行处理了
	if v.Kind() != reflect.Struct {
		return
	}
	for i := 0; i < v.NumField(); i++ {
		// 取得 struct 的 tag field
		field := v.Type().Field(i)
		// 获取tag的属性值
		// 还得判断属性是否首字母大写
		if v.Field(i).CanSet() && field.Tag.Get("inject") != "" {
			// 判断是否已经有了，不重复初始化实例
			if getV := beanFactoryImpl.Get(field.Type); getV != nil {
				v.Field(i).Set(reflect.ValueOf(getV))
				continue
			}
			// 兼容写 - 的方式
			if field.Tag.Get("inject") != "-" {
				// 表达式方式支持
				log.Println("使用了表达式的方式")
				ret := expr.BeanExpr(field.Tag.Get("inject"), beanFactoryImpl.ExprMap)
				if ret != nil && !ret.IsEmpty() {
					retValue := ret[0]
					if retValue != nil {
						beanFactoryImpl.Set(retValue)
						v.Field(i).Set(reflect.ValueOf(retValue))
					}
				}
			}
		}
	}
}

func NewBeanFactoryImpl() *BeanFactoryImpl {
	return &BeanFactoryImpl{beanMapper: make(BeanMapper), ExprMap: make(map[string]interface{})}
}
