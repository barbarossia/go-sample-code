package vardef

import(
	"fmt"
)
//常量定义
const i = 10000
const s = "hello"
const flag  = false

//全局变量
var global = "it's global"

//应该这么定义
const(
	c1 = 1
	c2 = "sss"
)
var(
	g1 = "g1"
	g2 = "g2"
)

func VarDefSample(){
	fmt.Println("***************VarDefSample*******************")
	//定义类型为string的a变量
	var a string
	//定义类型为string的b和c俩个变量
	var b,c string
	//定义类型为string的a1初始值为a1
	var a1 string = "a1"
	//定义类型为string的d和e，初始值为tom,jerry
	var d,e string = "tom","jerry"
	//定义d1和e1,类型根据后面的值判断
	var d1,e1 = "tom",1
	//根据初始化的值自动推导出相应的类型,只能用在func内部，所以var一般用来声明全局变量
	d2,e2 := "jerry",2
	//_是特殊变量名，任何的赋值都会丢弃
	_,_1 := 3,4
	//以上定义的所有变量都没有使用，编译就会报错
	fmt.Println(a,b,c,a1,d,e,d1,e1,d2,e2,_1)
}

func InnerType(){
	fmt.Println("***************InnerType*******************")
	//Boolean
	b := true
	if b {
		b = false
	}
	//数值类型int int8 int16 int32 int64 rune, uint8 uint16 uint32 uint64 byte
	//rune=int32,byte=uint8
	//float32 float64 complex128 complex64
	//这些类型的变量之间不允许互相赋值或操作，不然会在编译时引起编译器报错

	//字符串 字符串值是不可变的
	str, str2 := "tom","jerry"
	//使用+来连接字符串
	str3 := str + str2
	fmt.Println(str)
	fmt.Println(str3)
	//多行字符串
	mutilineStr := `hello
	world`
	fmt.Println(mutilineStr)
}

//枚举定义
// iota，这个关键字用来声明enum的时候采用，它默认开始值是0，const中每增加一行加1
func EnumDef(){
	fmt.Println("***************EnumDef*******************")
	const(
		x = iota
		y = iota
		z = iota
		w // 常量声明省略值时，默认和之前一个值的字面相同。这里隐式地说w = iota，因此w == 3。其实上面y和z可同样不用"= iota"
	)
	const v = iota // 每遇到一个const关键字，iota就会重置，此时v == 0

	const (
		e, f, g = iota, iota, iota //e=0,f=0,g=0 iota在同一行值相同
	)

}
//数组定义
//数组不能改变长度
func ArrayDef(){
	fmt.Println("***************ArrayDef*******************")
	arr := [3]int{1,2,3} //定义长度为3的int数组，并初始化值

	arr1 := [3]int{1, 2, 3} // 声明了一个长度为3的int数组
	arr2 := [10]int{1, 2, 3} // 声明了一个长度为10的int数组，其中前三个元素初始化为1、2、3，其它默认为0
	arr3 := [...]int{4, 5, 6} // 可以省略长度而采用`...`的方式，Go会自动根据元素个数来计算长度
	fmt.Println(arr[0])
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	//多维数组
	doubleArray := [2][2]int{[2]int{1,2},[2]int{1,2}}
	//简化为
	doubleArray = [2][2]int{{1,2},{1,2}}
	fmt.Println(doubleArray)
}
//slice
//因为数组默认不能改变长度，声明数组必须要指定长度，比如[3]int和[4]int是俩种类型，在将数组作为参数传入的时候，传入的其实是副本
//slice是引用，在定义的时候不需要指定长度
//slice的默认开始位置是0，ar[:n]等价于ar[0:n]
//slice的第二个序列默认是数组的长度，ar[n:]等价于ar[n:len(ar)]
//如果从一个数组里面直接获取slice，可以这样ar[:]，因为默认第一个序列是0，第二个是数组的长度，即等价于ar[0:len(ar)]
func SliceDef(){
	fmt.Println("***************SliceDef*******************")
	//slice定义
	var islice []int
	slice1 := []int{1,2,3,4}
	fmt.Println(islice)
	fmt.Println(slice1)
	// 声明一个含有10个元素元素类型为byte的数组
	var ar = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	// 声明两个含有byte的slice
	var a, b []byte
	// a指向数组的第3个元素开始，并到第五个元素结束，
	a = ar[2:5]
	//现在a含有的元素: ar[2]、ar[3]和ar[4]
	// b是数组ar的另一个slice
	b = ar[3:5]
	// b的元素是：ar[3]和ar[4]
	fmt.Println(a)
	fmt.Println(b)

	// 声明一个数组
	var array = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	// 声明两个slice
	var aSlice, bSlice []byte

	// 演示一些简便操作
	aSlice = array[:3] // 等价于aSlice = array[0:3] aSlice包含元素: a,b,c
	aSlice = array[5:] // 等价于aSlice = array[5:10] aSlice包含元素: f,g,h,i,j
	aSlice = array[:]  // 等价于aSlice = array[0:10] 这样aSlice包含了全部的元素

	// 从slice中获取slice
	aSlice = array[3:7]  // aSlice包含元素: d,e,f,g，len=4，cap=7
	bSlice = aSlice[1:3] // bSlice 包含aSlice[1], aSlice[2] 也就是含有: e,f
	bSlice = aSlice[:3]  // bSlice 包含 aSlice[0], aSlice[1], aSlice[2] 也就是含有: d,e,f
	bSlice = aSlice[0:5] // 对slice的slice可以在cap范围内扩展，此时bSlice包含：d,e,f,g,h
	bSlice = aSlice[:]   // bSlice包含所有aSlice的元素: d,e,f,g
	fmt.Println(bSlice)
	//slice是引用类型，所以当引用改变其中元素的值时，其它的所有引用都会改变该值，例如上面的aSlice和bSlice，如果修改了aSlice中元素的值，那么bSlice相对应的值也会改变。
}
//map
func MapDef(){

}