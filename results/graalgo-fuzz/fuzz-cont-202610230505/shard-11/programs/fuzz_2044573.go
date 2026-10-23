package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 uint8; F1 bool }
func genFn0(p0 uint) int16 {
    var v7 int8 = -1
    v8 := []int{}
    p0 += p0
    for fi9 := 0; fi9 < 12; fi9++ {
        var v10 int16 = -932
        fmt.Println(v7)
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 int64, p1 uint32) int {
    p1 += 342
    v11 := GenS1{F0: -1, F1: 127}
    _ = p0
    _ = p1
    _ = v11
    return -2147483648
}

func main() {
    var v1 uint64 = 1
    v2 := []uint32{2147483648, 876, 287, 2147483648}
    v3 := make(map[uint32]int)
    var v4 int64 = 0
    var v5 float64 = 0.0
    var v6 int64 = v4
    v12 := GenS1{F0: -32768, F1: 127}
    v5 -= (0.0 + (float64(0.0) - v5))
    v13 := GenS2{F0: 63, F1: true}
    if true {
        var v14 int8 = 90
        v4 -= v6
        var v15 int = 9223372036854775807
        _ = v14
        _ = v15
    } else {
        var v16 int64 = v4
        _ = v16
    }
    _ = len(v3)
    v1 -= (^(v1))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v6, v5)
    v6 = v6
    fmt.Println(v4, v4)
    for fi17 := 0; fi17 < 4; fi17++ {
        fmt.Println(v5)
        _ = fi17
    }
    var v18 uint64 = v1
    fmt.Println(v5, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
