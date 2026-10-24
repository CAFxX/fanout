package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 int32; F1 uint64 }
func genFn0(p0 uint8, p1 int64, p2 int16) uint64 {
    var v7 int32 = 517
    v8 := []int{-9223372036854775808, 1, 128, 333}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 int32 = 2147483647
    var v2 int32 = (-(v1))
    v3 := []int{}
    v4 := make(map[int]int)
    var v5 float32 = (3.14159 - float32(0.0))
    var v6 float32 = v5
    fmt.Println(v6, v2)
    var v9 float64 = 0.1
    v10 := GenS1{F0: 2, F1: 42}
    v11 := GenS1{F0: 680, F1: 4294967295}
    switch 4294967295 {
    case 128:
        v2 = (v1 / ((42) | 1))
    default:
        v12 := GenS1{F0: 1, F1: 255}
        _ = v12
    }
    v3 = append(v3, 4294967295)
    v13 := GenS2{F0: v2, F1: genFn0(128, 1, -32768)}
    v5 = (v5 * v5)
    v3 = append(v3, -2147483648)
    var v14 float32 = float32((v9 + v9))
    if (false || false) {
        var v15 bool = false
        _ = v15
    } else {
        var v16 uint8 = 255
        _ = v16
    }
    switch 32767 {
    case -623:
        v17 := v4[660]
        _ = v17
    }
    v5 += v5
    v2 += 85
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
