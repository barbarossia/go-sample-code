package main

import (
    "fmt"
    "github.com/derrick/stringutil"
    "github.com/derrick/vardef"
    "github.com/derrick/condition"
    "github.com/derrick/method"
    "github.com/derrick/structSample"
    "github.com/derrick/ooSample"
    "github.com/derrick/interfaceSample"
    "github.com/derrick/reflectSample"
)

func main() {
    fmt.Println("hello, world\n")
    fmt.Println(stringutil.Reverse("!oG ,olleH"))
    vardef.InnerType()
    vardef.VarDefSample()
    vardef.ArrayDef()
    vardef.MapDef()
    condition.GotoSample()

    i := 1
    i1 := method.PointerFunc(&i)
    fmt.Println("i + 1 = ",i1)
    fmt.Println("i = ",i)

    method.DeferSample()
    method.PubFilter()

    structSample.StructSample()

    a4 := ooSample.Vehicle{"Audi","A4","2.0T"}
    fmt.Println(a4.GetModelName())

    ooSample.MethodExtendsSample()

    fmt.Println("***************interfaceSample*******************")


    car := interfaceSample.Car{interfaceSample.PassengerCar{"Audi","A4","2.0T","轿车"}}
    suv := interfaceSample.SUV{interfaceSample.PassengerCar{"Audi","Q5","2.0T","SUV"}}
    sportscar := interfaceSample.SportsCar{interfaceSample.PassengerCar{"Audi","R8","4.0T","跑车"}}

    car.Ignition()
    suv.Ignition()
    sportscar.Ignition()

    /*var v interfaceSample.Vehicle
    v = car
    v.Ignition()
    v = suv
    v.Ignition()
    v = sportscar
    v.Ignition()*/

    varray := make([]interfaceSample.Vehicle, 3)
    varray[0],varray[1],varray[2] = car,suv,sportscar
    for _, value := range varray{
        value.Ignition()
    }

    interfaceSample.InstanceOf()

    reflectSample.ReflectSample()
}
