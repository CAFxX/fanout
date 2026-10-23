package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }

func main() {
    var v1 float32 = float32((1.5 / 0.1))
    v2 := make(map[int]uint32)
    v3 := []uint32{}
    var v4 uint64 = 1
    var v5 int16 = 1
    var v6 int8 = 100
    var v7 int8 = (^(v6))
    v2[42] = 2147483648
    v1 -= v1
    switch 9 {
    case 2147483647:
        if ("hello" <= "a") {
            v4 = v4
        } else {
            v2[1] = 377
        }
    default:
        v8 := GenS1{F0: 65535, F1: 484}
        _ = v8
    }
    if false {
        fmt.Println(v4, v4, v5)
    }
    var v9 int = 98
    v10 := GenS1{F0: 42, F1: -1}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
