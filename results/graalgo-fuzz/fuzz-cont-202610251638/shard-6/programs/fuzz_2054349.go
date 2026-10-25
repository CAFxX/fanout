package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 uint; F2 int }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint64 = 660
    var v3 float64 = (0.0 - (-2.25 / 1.5))
    v4 := []int{}
    if (v2 != v2) {
        v3 = -2.25
        fmt.Println(v3, v2)
        v5 := GenS1{F0: (v2 + (v2 % ((211) | 1)))}
        _ = v5
    }
    var v6 int32 = 468
    if (false && false) {
        v7 := GenS1{F0: (v2 & (v2 >> ((42) & 63)))}
        _ = v7
    }
    v4 = append(v4, -788)
    v8 := GenS2{F0: 32767, F1: 18446744073709551615, F2: 628}
    var v9 uint64 = v2
    var v10 float64 = v3
    fmt.Println(v10)
    v2 += (v9 ^ v9)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v8
    _ = v9
}
