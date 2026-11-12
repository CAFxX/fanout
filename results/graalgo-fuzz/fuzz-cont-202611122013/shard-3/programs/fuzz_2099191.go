package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16; F1 int8 }
func genFn0() uint {
    var v4 bool = true
    var v5 int64 = 680
    v6 := []int{32768, 32767, 333, 1}
    switch 339 {
    case 127:
        if len(v6) > 0 {
            v6[0] = -1
        }
    case 32768:
        v7 := GenS1{F0: 42, F1: 2}
        _ = v7
    default:
        for fi8 := 0; fi8 < 12; fi8++ {
            var v9 string = "hello"
            var v10 uint64 = 0
            var v11 int32 = 2147483647
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
    }
    _ = v4
    _ = v5
    _ = v6
    return 42
}
func genFn1() int64 {
    v12 := GenS2{F0: 768}
    for fi13 := 0; fi13 < 16; fi13++ {
        var v14 float64 = 0.0
        _ = fi13
        _ = v14
    }
    var v15 uint16 = 727
    v15 -= 655
    _ = v12
    _ = v15
    return 42
}

func main() {
    var v1 float32 = 1.5
    v2 := make(map[int]bool)
    v3 := make(map[int]int)
    for fi16 := 0; fi16 < 14; fi16++ {
        v1 = (v1 - ((-2.25 + ((3.14159 * (v1 / v1)) * v1)) / (v1 + v1)))
        v1 += 0.1
        var v17 int32 = -2147483648
        _ = fi16
        _ = v17
    }
    var v18 uint16 = 0
    var v19 float64 = (float64((v1 + v1)) * 0.0)
    v1 -= v1
    if (false && (42 == -51)) {
        var v20 uint = 42
        _ = v20
    } else {
        var v21 int32 = 2147483647
        fmt.Println(v1, v19, v1)
        _ = v21
    }
    v22 := GenS2{F0: 9223372036854775807}
    if (false || true) {
        v23 := []int{518, 148}
        for fi24 := 0; fi24 < 14; fi24++ {
            v1 = v1
            _ = fi24
        }
        v25 := GenS1{F0: -804, F1: 1}
        _ = v23
        _ = v25
    }
    v26 := []int{1, -1, -1, -8}
    if len(v26) > 0 {
        v26[0] = 42
    }
    v27 := v2[32768]
    v18 -= 2
    v28 := GenS3{F0: v18, F1: 1}
    v29 := GenS1{F0: -561, F1: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
}
