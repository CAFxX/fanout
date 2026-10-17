package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 int16 }

func main() {
    v1 := []int{-260, -9223372036854775808, 9223372036854775807, 2147483647}
    var v2 float32 = -2.25
    var v3 int8 = 0
    var v4 int8 = 42
    var v5 float32 = (v2 * v2)
    if (((175 < -750) && true) || (234 <= 2)) {
        fmt.Println(v2)
        var v6 float32 = 1.5
        v4 -= (v4 - v3)
        _ = v6
    }
    var v7 int64 = 1
    if false {
        v7 += v7
    }
    fmt.Println(v7, v7, v2)
    v5 = v2
    var v8 int = 85
    v9 := GenS1{F0: 4294967295, F1: 511, F2: 42}
    var v10 float64 = 1.5
    if len(v1) > 0 {
        v1[0] = (v8 + v8)
    }
    var v11 uint32 = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
