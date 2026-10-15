package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 string }

func main() {
    var v1 int = -427
    var v2 float32 = 1e10
    var v3 int8 = 100
    var v4 float64 = 3.14159
    var v5 int32 = -1
    v6 := []int{}
    var v7 int8 = 127
    var v8 float64 = ((((v4 * v4) - float64((1e10 / v2))) + 3.14159) / (v4 / (v4 - -2.25)))
    var v9 uint = 42
    fmt.Println(v5, v1, v2)
    v3 = v7
    v4 -= 3.14159
    fmt.Println(v1, v4)
    v10 := GenS1{F0: v5, F1: 1, F2: "\n\t"}
    var v11 float32 = (float32(v8) - v2)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
