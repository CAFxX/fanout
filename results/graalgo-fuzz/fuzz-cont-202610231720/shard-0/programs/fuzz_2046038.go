package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 uint }

func main() {
    var v1 uint32 = 2147483648
    var v2 int = 2147483647
    var v3 int16 = -1
    if false {
        v4 := GenS1{F0: 0, F1: v1, F2: 0}
        var v5 int8 = 1
        var v6 uint64 = 0
        _ = v4
        _ = v5
        _ = v6
    }
    v1 = v1
    var v7 float32 = 1.5
    var v8 int = v2
    var v9 uint = 904
    v3 -= 1
    v7 += 1.5
    var v10 bool = (3.14159 == 1.5)
    var v11 uint32 = (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
