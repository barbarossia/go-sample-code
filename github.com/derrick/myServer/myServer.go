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

func sayHi2u(response http.ResponseWriter, request *http.Request){
	request.ParseForm()
	fmt.Println(request.Form) //整个form是一个map，key：[value,value]
	fmt.Println(request.Form["username"]) //返回key是username的数组
	request.Form.Get("username")
	//barray := []byte("Hello:"+request.Form["username"][0])
	barray := []byte("Hello:"+request.Form.Get("username"))//也可以直接通过Get方法获取
	response.Write(barray)
}

func loadInfo(response http.ResponseWriter, request *http.Request){
	response.Header().Add("Access-Control-Allow-Origin", "*")
	request.ParseForm()
	request.Form.Get("id")
	//barray := []byte("Hello:"+request.Form["username"][0])
	barray := []byte("name: " + request.Form.Get("id") +" </br> age:28 </br> gender:male")//也可以直接通过Get方法获取
	response.Write(barray)
}

func main() {
	//hello world
	//http.HandleFunc("/", helloWeb) //设置访问的路由
	http.HandleFunc("/sayHello", helloWeb) //mapping request path
	http.HandleFunc("/sayHi2u", sayHi2u) //mapping request path
	http.HandleFunc("/loadInfo", loadInfo) //mapping request path
	err := http.ListenAndServe(":9001", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

