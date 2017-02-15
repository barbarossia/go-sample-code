package main

import "fmt"

func ifSample(){
	b := true;
	//条件不需要括号
	if b {

	}else{

	}
	i := 10;
	//Go可以在条件中定义变量，这个变量的作用域只在条件逻辑内，如下的x只在条件内有效
	if x := square(i); x>99{
		fmt.Println("the square of i greater than 99")
	}else{
		fmt.Println("the square of i less than 99")
	}

	//多条件
	if i == 0{

	}else if i >100{

	}else{

	}

}
func square(i int) int{
	return i*i;
}

func gotoSample(){
	fmt.Println("***************GotoSample*******************")
	i := 0;
	SomePlace:
	fmt.Println("i:",i)
	i++
	if i <10 {
		goto SomePlace
	}
}
//for可以迭代，也可以作为while使用
func forSample(){
	fmt.Println("***************ForSample*******************")
	//for expression1 变量声明，expression2 条件判断，expression3 每轮循环结束调用
	for i:=1;i<10;i++{
		fmt.Println("this is print in a for:",i)
	}
	add := 0
	//忽略;就变成了while
	for add<10 {
		add += add
	}
	//break
	for i:=10;i<10;i++{
		if i==5{
			break
		}
		fmt.Println("break sample:",i)
	}
	//continue
	for i:=10;i<10;i++{
		if i==5{
			continue
		}
		fmt.Println("continue sample:",i)
	}

}
//switch
func switchSample(){
	fmt.Println("***************SwitchSample*******************")
	//switch和其他语言没有什么区别，只要switch的类型一直就可以了
	//默认在每个case后做break操作，所以不需要手工写break
	//可以加入fallthrough来强制执行后面的代码
	i:=10
	switch i {
	case 1:
		fmt.Println("switch sample:",i)
	case 2,3,4:
		fmt.Println("switch sample,equal to 2,3 or 4",i)
	case 5:
		fmt.Println("switch sample fallthrough 5",i)
		fallthrough
	case 6:
		fmt.Println("switch sample 6",i)
	default:
		fmt.Println("default")
	}
}

func main() {
	ifSample()
	gotoSample()
	forSample()
	switchSample()
}