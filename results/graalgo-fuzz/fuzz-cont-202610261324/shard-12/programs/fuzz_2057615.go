package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 uint16; F1 int8 }

func main() {
    var v1 int8 = 100
    v2 := make(map[int]int)
    var v3 bool = true
    var v4 int = 2147483647
    switch 32768 {
    case -535:
        var v5 int64 = 42
        _ = v5
    default:
        if v3 {
            v6 := []int{v4}
            v3 = v3
            _ = v6
        } else {
            fmt.Println(v1, v3)
        }
    }
    v7 := GenS1{F0: (v1 - v1), F1: -1}
    v1 = (v1 & v1)
    for fi8 := 0; fi8 < 20; fi8++ {
        if (0 >= 0) {
            v9 := []int{v4, (v4 * fi8), fi8, 582}
            _ = v9
        }
        _ = fi8
    }
    v1 += (v1 & v1)
    v4 -= (v4 % ((1) | 1))
    var v10 int32 = -739
    fmt.Println(v1)
    v11 := GenS2{F0: 42, F1: -26}
    for fi12 := 0; fi12 < 18; fi12++ {
        for fi13 := 0; fi13 < 15; fi13++ {
            var v14 float64 = 0.1
            v1 -= 79
            _ = fi13
            _ = v14
        }
        fmt.Println(fi12, v1, v4)
        fmt.Println(v10)
        _ = fi12
    }
    v15 := []int{(^(v4)), v4}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
