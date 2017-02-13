package concurrentSample

import (
	"fmt"
	"runtime"
)

func SayHello(name string){
	fmt.Println("***************concurrentSample.SayHello*****************")
	for i := 0; i < 3; i++ {
		runtime.Gosched()
		fmt.Println(name, ":", i)
	}
}