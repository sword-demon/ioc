package Injector

import "reflect"

type BeanMapper map[reflect.Type]reflect.Value

func (beanMapper BeanMapper) add(bean interface{}) {
	t := reflect.TypeOf(bean)
	// 判断的t的类型如果不是指针就panic
	if t.Kind() != reflect.Ptr {
		panic("bean must be a pointer")
	}
	beanMapper[t] = reflect.ValueOf(bean)
}

func (beanMapper BeanMapper) get(bean interface{}) reflect.Value {
	//t := reflect.TypeOf(bean)
	var t reflect.Type
	// 添加断言判断
	if bt, ok := bean.(reflect.Type); ok {
		t = bt
	} else {
		t = reflect.TypeOf(bean)
	}
	if v, ok := beanMapper[t]; ok {
		return v
	}
	return reflect.Value{}
}
