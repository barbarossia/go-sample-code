package main

import (
    "fmt"
    "github.com/derrick/stringutil"
    "github.com/derrick/vardef"
    "github.com/derrick/condition"
    "github.com/derrick/method"
    "github.com/derrick/structSample"
    "github.com/derrick/ooSample"
)

func main() {
    fmt.Printf("hello, world\n")
    fmt.Printf(stringutil.Reverse("!oG ,olleH"))
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

}
