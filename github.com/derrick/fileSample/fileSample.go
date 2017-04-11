package main

import (
	//"encoding/xml"
	//"io/ioutil"
	"os"
	"fmt"
)

func main() {
	dirSample()
	fileSample()
}

func dirSample()  {
	//创建一个目录
	os.Mkdir("1stDir",0777)
	//创建多级目录
	os.MkdirAll("1stDir/2/3",0777)
	//重命名
	os.Rename("1stDir","1stDir-bak")
	//删除目录
	os.Remove("1stDir")
	//删除多级目录
	os.Remove("1stDir")
}

func fileSample(){
	//创建文件
	file, error := os.Create("1stFile")
	if error != nil {
		fmt.Println(file, error)
		return
	}
	defer file.Close()
	b := []byte("hello file")
	//写入文件
	file.Write(b)
	//直接写入string
	file.WriteString("this is a string")
}