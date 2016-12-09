package main

import (
    "fmt"
    "github.com/derrick/stringutil"
    "github.com/derrick/vardef"
)

func main() {
    fmt.Printf("hello, world\n")
    fmt.Printf(stringutil.Reverse("!oG ,olleH"))
    vardef.InnerType()
    vardef.VarDefSample()
}
