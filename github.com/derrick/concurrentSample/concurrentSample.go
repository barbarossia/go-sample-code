package main

import (
	"fmt"
	"runtime"
)

func sayHello(name string){
	fmt.Println("***************concurrentSample.SayHello*****************")
	//runtime.Gosched()表示将CPUC时间让出，在某个时间点继续执行，而恢复执行的时间点由Go决定
	for i := 0; i < 3; i++ {
		runtime.Gosched()
		fmt.Println(name, ":", i)
	}
}

func main()  {
	//使用go关键字启动一个goroutine
	go sayHello("goroutine")
	sayHello("direct")
}


func chanSample(){
	fmt.Println("***************concurrentSample.ChanSample*****************")
	//定义一个string类型的chan，chan必须指定类型，并且只能传递指定类型的值
	//chan是blocking的，发送和接收都必须等到发送者和接收者ready才可以
	cs := make(chan string)
	go func() { cs <- "ping" }()
	s := <- cs
	fmt.Println(s)
}


func bufferedChanSample(){
	fmt.Println("***************concurrentSample.BufferedChanSample*****************")
	cs := make(chan string, 2)
	cs <- "ping 1"
	cs <- "ping 2"
	fmt.Println(<-cs)
	fmt.Println(<-cs)
}