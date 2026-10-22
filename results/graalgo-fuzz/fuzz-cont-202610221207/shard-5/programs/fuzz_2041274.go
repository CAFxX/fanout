package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }

func main() {
    var v1 bool = (703 <= 0)
    var v2 int = 691
    var v3 string = ""
    var v4 int32 = -698
    v5 := GenS1{F0: 42, F1: 354}
    v2 = v2
    var v6 uint = 42
    var v7 int64 = 1
    v8 := []int{(v2 ^ 1), 557, v2, v2}
    v9 := GenS1{F0: 100, F1: v4}
    if v1 {
        var v10 float32 = ((0.0 - 3.14159) + (float32(0.0) - (-2.25 + -2.25)))
        _ = v10
    }
    var v11 int64 = v7
    fmt.Println(v2, v1, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
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
