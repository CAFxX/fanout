package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 bool }
func genFn0() uint8 {
    var v5 float64 = 0.0
    v6 := []int{2147483647}
    var v7 uint16 = 691
    _ = v5
    _ = v6
    _ = v7
    return 98
}
func genFn1() uint8 {
    var v8 float32 = 0.0
    if false {
        v8 = v8
    } else {
        v8 = v8
    }
    v8 = v8
    fmt.Println(v8)
    v9 := []int{127, 42}
    _ = v8
    _ = v9
    return 42
}

func main() {
    v1 := []string{"hello", "foo bar", "a", "hello"}
    var v2 int8 = 127
    var v3 uint64 = 843
    v4 := []uint32{2147483648}
    switch "\n\t" {
    case "\n\t":
        _ = len(v1)
    default:
        v2 = 77
    }
    var v10 float32 = float32(-2.25)
    _ = len(v4)
    fmt.Println(v10, v2, v2)
    if len(v4) > 0 {
        v4[0] = 4294967295
    }
    var v11 float64 = 1e10
    if false {
        fmt.Println(v11)
        var v12 uint16 = 114
        _ = v12
    }
    v13 := GenS1{F0: uint32(v3), F1: 112, F2: false}
    v4 = append(v4, 2)
    var v14 int32 = 42
    v14 -= 0
    var v15 string = "hello"
    if false {
        var v16 uint64 = v3
        fmt.Println(v10, v16)
        var v17 bool = (!false)
        _ = v16
        _ = v17
    }
    _ = len(v4)
    if false {
        v14 -= (v14 / ((-46) | 1))
        v18 := GenS1{F0: 2147483647, F1: 255, F2: ((255 != 0) || false)}
        _ = len(v4)
        _ = v18
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
