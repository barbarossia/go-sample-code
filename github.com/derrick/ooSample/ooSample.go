package ooSample

import "fmt"

type Vehicle struct{
	Make string
	Series string
	Model string
}
//通过这种方式来为struct添加method，这里的ReceiverType为Vehicle，及通过在方法名前面添加ReceiverType来为对应的type添加method
// func (r ReceiverType) funcName(parameters) (results)
func (v Vehicle) GetModelName() string{
	return v.Make + "-" + v.Series + "-" + v.Model
}
//receiver也支持指针，这种情况下会改变原始的值，不传指针的话就只是一个值的copy
//如果一个method的receiver是*T,你可以在一个T类型的实例变量V上面调用这个method，而不需要&V去调用这个method类似的
//如果一个method的receiver是T，你可以在一个*T类型的变量P上面调用这个method，而不需要 *P去调用这个method
func (v *Vehicle) setModelName(vr Vehicle){
	v.Model = vr.Make
}

type Car struct {
	Vehicle
	PlateNo string
}
//Car本身没有GetModelName和setModelName，在定义了Vehicle匿名字段后，也可以直接调用
func MethodExtendsSample(){
	fmt.Println("***************methodExtendsSample*******************")
	car := Car{Vehicle{"Ford","蒙迪欧","2.0T"},"沪A00001"}
	fmt.Println("Car可以直接调用Vehicle的GetModelName方法:",car.GetModelName())
}