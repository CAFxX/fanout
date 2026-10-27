package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 int = -9223372036854775808
    var v2 uint8 = 1
    var v3 int8 = 59
    v4 := []int{(v1 - 42), v1}
    v5 := []int{v1, (-(v1)), v1}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v6 int16 = -32768
    if false {
        var v7 uint8 = (v2 | 0)
        _ = v7
    }
    v6 = (-(v6))
    v8 := GenS1{F0: -2147483648}
    fmt.Println(v1, v1, v3)
    var v9 bool = false
    v10 := GenS1{F0: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
