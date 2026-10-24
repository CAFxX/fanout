package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int; F1 bool }
func genFn0(p0 int8, p1 int16, p2 string) int32 {
    switch 42 {
    case 76:
        p0 = p0
    case 9223372036854775807:
        fmt.Println(p2, p1, p0)
    default:
        v7 := GenS1{F0: 0, F1: (0 < 18446744073709551615)}
        _ = v7
    }
    var v8 float32 = (float32(float64(((float32(float64(1.5)) - 1.5) / ((float32(0.0) / 3.14159) - float32(1.5))))) - float32(0.0))
    v9 := []int{821, 32768}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return -46
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 uint = 42
    var v3 int = 0
    var v4 uint32 = 4294967295
    var v5 int8 = -1
    v6 := make(map[string]int)
    v10 := []int{}
    v2 -= v2
    if false {
        fmt.Println(v1)
        switch 32768 {
        case -818:
            v3 = (v3 | v3)
        }
        fmt.Println(v1, v3, v4)
    }
    v11 := GenS1{F0: 42, F1: false}
    var v12 uint16 = 260
    for fi13 := 0; fi13 < 9; fi13++ {
        var v14 int = fi13
        _ = fi13
        _ = v14
    }
    var v15 uint8 = 128
    for fi16 := 0; fi16 < 12; fi16++ {
        var v17 int64 = v1
        _ = fi16
        _ = v17
    }
    v2 = 42
    fmt.Println(v5, v15)
    v5 += 42
    v1 = v1
    v18 := GenS2{F0: 61}
    var v19 float32 = 0.1
    v20 := GenS3{F0: (v3 % ((v3) | 1)), F1: true}
    var v21 int64 = -9223372036854775808
    v4 = (^(v4))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
