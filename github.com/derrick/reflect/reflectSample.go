package reflect

import (
	"reflect"
	"fmt"
)

type passengerCar struct{
	Make string
	Series string
	Model string
	VehicleType string
}

func ReflectSample(){
	fmt.Println("***************reflectSample*****************")
	p := passengerCar{"audi","a4","2.0","car"}
	v := reflect.ValueOf(p)
	t := reflect.TypeOf(p)
	fmt.Println("ValueOf:",v)
	fmt.Println("TypeOf:",t)
}
