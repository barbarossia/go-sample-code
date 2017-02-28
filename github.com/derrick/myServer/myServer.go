package main

import (
	"fmt"
	"net/http"
	"log"
)

func helloWeb(response http.ResponseWriter, request *http.Request){
	fmt.Println(request.Method)
	fmt.Println(request.Header)
	barray := []byte("hello world")
	response.Write(barray)
}

func main() {
	//hello world
	http.HandleFunc("/", helloWeb) //设置访问的路由
	http.HandleFunc("/sayHello", helloWeb) //mapping request path
	err := http.ListenAndServe(":9001", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}