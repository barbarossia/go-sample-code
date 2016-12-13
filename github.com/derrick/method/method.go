package method

import (
	"fmt"
)

//多返回值，返回值不指定名称
func mutiReturn(a,b int) (int,int){
	return a+b,a-b
}
//多返回值，指定名称,官方推荐写法，可读性好一些
func mutiReturnWithName(a,b int)(sum int,diff int){
	sum = a + b
	diff = a - b
	return sum, diff
}
//变参函数 这里的arg其实是一个string的slice
func indefiniteQuantityArg(arg...string){

}
//传指针到函数,这样的话i就不会是一个值的copy了，i本身在func中就会被改变
func PointerFunc(i *int) int{
	fmt.Println("***************PointerFuc*******************")
	*i = *i +1
	return *i
}
//defer Go本身就支持defer，在return或者函数退出之前安装先进后出的顺序执行
func DeferSample(){
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}
//定义testInt type
type testInt func(int) bool
func isOdd(i int) bool{
	if i%2==0{
		return false
	}else{
		return true
	}
}
func isEven(i int) bool{
	if i%2==0{
		return true
	}else{
		return false
	}
}
//定义一个func，将之前定义的type testInt作为参数传入
func intFilter(sliceA []int,f testInt) []int{
	fmt.Println("***************Intfilter*******************")
	var result []int
	for _,value:= range sliceA{
		if f(value){
			result = append(result,value)
		}
	}
	return result
}
//调用intFilter
func PubFilter(){
	slice := []int {1, 2, 3, 4, 5,6, 7}
	fmt.Println("slice = ", slice)
	odd := intFilter(slice, isOdd)    // 函数当做值来传递了
	fmt.Println("Odd elements of slice are: ", odd)
	even := intFilter(slice, isEven)  // 函数当做值来传递了
	fmt.Println("Even elements of slice are: ", even)
}