package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 string }

func main() {
    var v1 string = "hello"
    var v2 string = (("hello" + "foo bar") + v1)
    var v3 int16 = 1000
    v1 = v1
    var v4 uint32 = 1
    v5 := GenS1{F0: 4294967295, F1: 42, F2: v2}
    var v6 uint64 = 4294967295
    var v7 int8 = 30
    v6 += (^(v6))
    var v8 float64 = float64(0.0)
    fmt.Println(v3)
    var v9 uint16 = 2
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
