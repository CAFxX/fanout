package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 uint }

func main() {
    v1 := []uint32{}
    var v2 bool = (!true)
    var v3 int32 = 2147483647
    var v4 int = -2147483648
    var v5 int16 = 0
    var v6 bool = (!(1 <= 42))
    v7 := GenS1{F0: 42}
    var v8 uint = 0
    v9 := GenS1{F0: 263}
    if v6 {
        var v10 int8 = -83
        var v11 uint8 = 42
        _ = v10
        _ = v11
    } else {
        var v12 int32 = v3
        var v13 float32 = (((3.14159 * (((0.1 - 3.14159) + 1.5) - 3.14159)) + 1e10) - float32(1.5))
        v13 += (v13 - v13)
        _ = v12
        _ = v13
    }
    v3 -= v3
    v14 := GenS2{F0: -929, F1: (v8 % ((v8) | 1))}
    v1 = append(v1, 2)
    var v15 uint16 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
