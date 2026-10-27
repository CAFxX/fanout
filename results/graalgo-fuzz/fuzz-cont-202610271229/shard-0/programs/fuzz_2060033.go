package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 uint64; F1 uint; F2 int32 }
func genFn0() uint32 {
    switch -9223372036854775808 {
    case 0:
        v5 := []int{2147483647, 9223372036854775807, 127}
        _ = v5
    case -613:
        for fi6 := 0; fi6 < 10; fi6++ {
            fmt.Println(fi6)
            _ = fi6
        }
    }
    v7 := []int{-915, 127, 1, -200}
    _ = v7
    return 4294967295
}

func main() {
    var v1 int8 = 17
    var v2 uint16 = 255
    v3 := make(map[int]uint32)
    v4 := []uint32{346, 42, 2147483648, 0}
    if (0 != 9223372036854775807) {
        var v8 bool = false
        v9 := GenS1{F0: v1, F1: -1}
        _ = v8
        _ = v9
    } else {
        var v10 float64 = 1e10
        var v11 uint = uint(v2)
        var v12 int16 = 42
        _ = v10
        _ = v11
        _ = v12
    }
    v4 = append(v4, 2147483648)
    var v13 float32 = (((0.0 - 3.14159) / float32(1e10)) + 1e10)
    v14 := GenS1{F0: -22, F1: 0}
    fmt.Println(v13, v1, v2)
    var v15 int64 = 42
    v1 = (v1 / ((v1) | 1))
    fmt.Println(v2, v15, v13)
    fmt.Println(v15, v15)
    v16 := GenS1{F0: v1, F1: -927}
    fmt.Println(v1, v15)
    v17 := GenS2{F0: 0, F1: 42, F2: 2147483647}
    for fi18 := 0; fi18 < 13; fi18++ {
        var v19 int = -9223372036854775808
        _ = fi18
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
