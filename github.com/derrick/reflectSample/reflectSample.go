package reflectSample

import (
	"reflect"
	"fmt"
)

type PassengerCar struct{
	Make string
	Series string
	Model string
	VehicleType string
}

func (p *PassengerCar) GetModelName() string{
	return "this is PassengerCar's GetModelName:" + p.Make + "-" + p.Series + "-" + p.Model
}

func ReflectSample(){
	fmt.Println("***************ValueOf*****************")
	p := PassengerCar{"audi","a4","2.0","car"}
	v := reflect.ValueOf(&p)
	/*
	Addr()返回指针的地址,如果CanAddr()返回false，会产生panic错误
	Addr()通常用来获取struct或者slice的指针，用来给需要point receiver的方法传参
	传引用的话CanXXX方法才会返回true，GO默认都是传值的副本
	*/
	fmt.Println("ValueOf:",v)
	fmt.Println("CanAddr()是否可以调用Addr:",v.Elem().FieldByName("Make").CanAddr())
	fmt.Println("Addr:",v.Elem().FieldByName("Make").Addr())
	fmt.Println("***************Set的使用*****************")
	fmt.Println("CanSet()是否可以调用set:",v.Elem().FieldByName("Make").CanSet())
	//Set的参数必须为Value对象
	v.Elem().FieldByName("Make").Set(reflect.ValueOf("BMW"))
	fmt.Println("Set后的结果:",v)

	fmt.Println("***************Call的使用*****************")
	//如果此处v的类型不是func，那么会报panics错误
	result := v.MethodByName("GetModelName").Call([]reflect.Value{})
	fmt.Println(result)
	fmt.Println("***************Cap的使用*****************")
	arr := [...]string{"One","Two","Three"}
	//Cap只能供数组使用，否则会报panics错误
	len := reflect.ValueOf(arr).Cap()
	fmt.Println(len)
	fmt.Println("***************Elem的使用*****************")
	//Elem返回interface v包含的值或者指针 v的指向，如果不是interface或者指针会报panics错误
	elem := v.Elem()
	fmt.Println("Elem:",elem)
	fmt.Println("***************Filed的使用*****************")
	//Field返回struct v的第i个字段，如果v不是struct或者i大于v的字段数量会报panics错误
	f := v.Elem().Field(0)
	fmt.Println("Filed0:",f)
	fmt.Println("***************TypeOf*****************")
	t := reflect.TypeOf(p)
	fmt.Println("TypeOf:",t)
	
}
