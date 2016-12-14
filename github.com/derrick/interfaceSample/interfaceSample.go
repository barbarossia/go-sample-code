package interfaceSample

import "fmt"

//空interface，不包含任何的method
//NullInterface的对象可以存储任意类型的值
type NullInterface interface {

}

type Vehicle interface {
	Ignition()
	Drive()
	Brake()
}

type PassengerCar struct{
	Make string
	Series string
	Model string
	VehicleType string
}

type Car struct {
	PassengerCar
}

type SUV struct {
	PassengerCar
}

type SportsCar struct {
	PassengerCar
}
//PassengerCar实现了Ignition()方法
func (p PassengerCar) Ignition(){
	fmt.Printf("The %s is launched:\n",p.VehicleType)
}
//PassengerCar实现了Ignition()方法
func (p PassengerCar) Drive(){
	fmt.Printf("The %s is driving:\n",p.VehicleType)
}
//PassengerCar实现了Ignition()方法
func (p PassengerCar) Brake(){
	fmt.Printf("The %s is braking:\n",p.VehicleType)
}
//Car重载了PassengerCar的Ignition()方法
func (c Car) Ignition(){
	fmt.Println("The car is launched:",c)
}

//使用空interface作为参数，可以接受任意类型的值作为参数
//使用空interface作为返回值，可以返回任意类型的值
func nullInterfaceSample(inf NullInterface) NullInterface{
	var temp NullInterface
	temp =  1
	temp = "this is string"
	temp = false
	return temp;
}
//类型判断，interface可以存储任何值，可以通过Comma-ok断言和switch来判读类型
func InstanceOf(){
	arr := make([]NullInterface,3)
	arr[0],arr[1],arr[2] = 1,"s",false
	fmt.Println("***************InstanceOf*****************")
	fmt.Println("Comma-ok:")
	//通过Comma-ok判断
	for index,ele := range arr{
		if value, ok :=  ele.(int); ok{
			fmt.Printf("arr[%d] this is a int, its value is %d \n",index,value)
		}else if value, ok :=  ele.(string); ok{
			fmt.Printf("arr[%d] this is a string, its value is %s \n",index,value)
		}else if value, ok :=  ele.(bool); ok{
			fmt.Printf("arr[%d] this is a bool, its value is %t \n",index,value)
		}else{
			fmt.Println("annother type")
		}
	}
	fmt.Println("switch:")
	//通过switch判断
	for index,ele := range arr {
		switch value := ele.(type){
			case int:
				fmt.Printf("arr[%d] this is a int, its value is %d \n",index,value)
			case string:
				fmt.Printf("arr[%d] this is a int, its value is %s \n",index,value)
			case bool:
				fmt.Printf("arr[%d] this is a int, its value is %t \n",index,value)
			default:
				fmt.Println("annother type")

		}
	}
}