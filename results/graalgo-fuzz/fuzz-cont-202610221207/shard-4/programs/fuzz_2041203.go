package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 int64 }
func genFn0() uint32 {
    v4 := GenS1{F0: 2, F1: 0, F2: 360}
    var v5 uint = 42
    v5 += (v5 - (v5 / ((14) | 1)))
    switch 127 {
    case 9223372036854775807:
        v6 := []int{-793, 99, 32768}
        _ = v6
    case 667:
        var v7 int8 = -128
        _ = v7
    case 32767:
        fmt.Println(v5)
    }
    _ = v4
    _ = v5
    return 42
}

func main() {
    var v1 int = 9223372036854775807
    var v2 float64 = (0.1 + float64(1.5))
    var v3 uint = 2
    var v8 uint8 = 127
    var v9 float64 = ((v2 - 1.5) / float64(((1.5 / 1.0) * 1.5)))
    v10 := []int{v1, 4294967295}
    if len(v10) > 0 {
        v10[0] = v1
    }
    var v11 float64 = 0.1
    var v12 int32 = -94
    if len(v10) > 0 {
        v10[0] = 247
    }
    fmt.Println(v9)
    var v13 uint32 = 2
    var v14 uint32 = v13
    fmt.Println(v11, v9)
    var v15 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
