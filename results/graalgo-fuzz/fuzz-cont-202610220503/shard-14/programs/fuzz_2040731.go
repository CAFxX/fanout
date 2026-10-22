package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0() int64 {
    var v4 float32 = (0.1 - ((3.14159 / -2.25) * float32(0.1)))
    var v5 uint32 = 1
    _ = v4
    _ = v5
    return 201
}

func main() {
    var v1 uint64 = 2
    var v2 int16 = 1000
    var v3 uint32 = 2
    var v6 int = -141
    var v7 string = ("a" + ("foo bar" + ""))
    var v8 int64 = -9223372036854775808
    v1 = 2
    var v9 string = "a"
    v10 := GenS1{F0: 0}
    v1 -= v1
    v3 = 2147483647
    fmt.Println(v8)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
