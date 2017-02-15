package main

import "fmt"

//定义一个类型
type vehicle struct{
	make string
	series string
	model string
}
//可以定义匿名字段，只提供类型，不写字段名
type vehicleDefine struct{
	vehicle //定义了一个匿名的vehicle
	plateNo string
	model string //vehicleDefine.model优先访问，可以通过vehicleDefine.vehicle.model访问vehicle中的model
}

var V vehicle

func structSample(){
	fmt.Println("***************StructSample*******************")
	V.make = "Audi"
	V.series = "Q7"
	V.model = "3.0T"
	fmt.Println("字段赋值方式初始化",V)

	BMW740 := vehicle{"BMW","7","740"}
	fmt.Println("按照字段顺序初始化 BMW740",BMW740)

	BenzClass := vehicle{series:"C",make:"Benz",model:"210"}
	fmt.Println("按照字段名字 BenzClass",BenzClass)

	//这里BenzBclass是指针
	BenzBclass := new(vehicle)
	fmt.Println("new 指针 BenzBclass",BenzBclass)

	mainLoss := vehicleDefine{vehicle{"Cadillac","CTS","2.0T"},"沪A00001","this is mainLoss's model"}
	fmt.Println("这是一个匿名字段的struct",mainLoss)
	fmt.Println("可以直接访问匿名字段mainLoss.make：",mainLoss.make)
	fmt.Println("也可以通过匿名字段的名称访问mainLoss.vehicle.make:",mainLoss.vehicle.make)
}

func main(){
	structSample()
}