package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 int32 }

func main() {
    var v1 uint16 = 119
    v2 := make(map[string]bool)
    v3 := make(map[int]uint32)
    v4 := make(map[string]int)
    var v5 float32 = 1.5
    v6 := GenS1{F0: 0, F1: 0, F2: 444}
    v1 = v1
    var v7 float32 = 0.1
    switch ("" + "") {
    case "a":
        v1 += 255
    default:
        v8 := GenS1{F0: 31, F1: 0, F2: 1}
        _ = v8
    }
    v7 -= (v7 * (v5 - v7))
    if (-1 > 0) {
        v9 := []int{473, 0, 901, 9223372036854775807}
        v5 = (float32((0.1 - float64(1e10))) - (v5 * (1.5 / v5)))
        _ = v9
    } else {
        v7 += v7
        var v10 uint64 = 1
        _ = v10
    }
    v11 := []int{2147483647, 922}
    fmt.Println(v7, v5)
    if ("" > "") {
        fmt.Println(v1)
        var v12 float32 = float32(3.14159)
        v13 := GenS1{F0: 42, F1: 4294967295, F2: 42}
        _ = v12
        _ = v13
    } else {
        v14 := v3[-1]
        var v15 uint16 = 65535
        _ = v14
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
