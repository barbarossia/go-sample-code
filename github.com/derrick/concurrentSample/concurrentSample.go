package main

import (
	"fmt"
	"runtime"
	"time"
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
	/*go sayHello("goroutine")
	sayHello("direct")
	bufferedChanSample()
	//异步调用
	done := make(chan bool,1)
	go synChan(done)
	//没有这行代码，synChan中的内容将不会打印，程序会在worker启动前就退出了
	<-done
	chanDirection()*/

	/*c := make(chan string,4)
	c <- "1"
	c <- "2"
	c <- "3"
	c <- "4"
	//生产者通过内置的close关闭channel
	在消费方可以通过语法v, ok := <-ch测试channel是否被关闭。如果ok返回false，那么说明channel已经没有任何数据并且已经被关闭
	close(c)
	//use range to iterate
	for i := range c{
		fmt.Println(i)
	}*/
	/*selectSample()*/
	timeOutSample()
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
	//如果向chanel中发送的数据超过buffer的容量，会报fatal error: all goroutines are asleep - deadlock!
	//cs <- "ping 3"
	fmt.Println(<-cs)
	fmt.Println(<-cs)
}

func synChan(done chan bool){
	fmt.Println("***************concurrentSample.synChan*****************")
	fmt.Println("working......")
	time.Sleep(2000)
	fmt.Println("done")
	done <- true
}

func chanDirection(){
	fmt.Println("***************concurrentSample.chanDirection*****************")
	//normal chan
	normalChan := make(chan string)
	go func() { normalChan <- "normal" }()
	s := <- normalChan
	fmt.Println(s)
	//send only
	sendChan := make(chan<- string,1)
	sendChan <- "send only"
	go sendChanFunc(sendChan)
	//receive only
	recevieChan := make(<- chan string,1)
	//this line will catch compile error
	//recevieChan <- "ping"
	go recevieChanFunc(recevieChan)
}

func sendChanFunc(c chan <- string){
	//this line will catch compile error
	//s := <-c
	//fmt.Println(s)
}

func recevieChanFunc(c <- chan string){
	s := <-c
	fmt.Println(s)
}
//select用来监听channel上的数据流动，可以理解为channel的switch
//select默认是阻塞的，只有当监听的channel中有发送或接收可以进行时才会运行，当多个channel都准备好的时候，select是随机的选择一个执行的。
//default就是当监听的channel都没有准备好的时候，默认执行的
func selectSample(){
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <- c1:
				fmt.Println(msg1)
			case msg2 := <- c2:
				fmt.Println(msg2)
			default :
				fmt.Println("this is default")

			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
//timeout sample
func timeOutSample(){
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}
}