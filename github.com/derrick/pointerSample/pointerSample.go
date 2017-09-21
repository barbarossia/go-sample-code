package main

import (
	"fmt"
)

//此处param传递的只是原始对象的值的副本，++并不会改变原来对象的值
func addWithVal(param int) {
	param++
}

//必须传递原始对象的值才能改变原始值
//指针指向了内存中值存储的内存地址，*int代表指针指向了一个int
func addWithPointer(param *int) {
	//此处指针指向了++后的value在内存中的地址
	*param++
	//此处如果使用param++会编译错误，因为param不是一个int
	//param++
}

func square(x *float64) {
	*x = *x * *x
}

func main() {
	var x = 1
	addWithVal(x)
	fmt.Println("addWithVal:", x)
	//使用&符号来获取变量的内存地址，&x和addWithPointer中的*int指向了同一块内存地址
	addWithPointer(&x)
	fmt.Println("addWithPointer:", x)

	//new关键字会分配内存并返回指针
	newX := new(int)
	addWithPointer(newX)
	fmt.Println("addWithPointer:", *newX)

	y := 1.5
	square(&y)
	fmt.Println("addWithPointer:", y)
}
