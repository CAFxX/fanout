package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 uint8 }

func main() {
    v1 := []int{636}
    v2 := make(map[int]string)
    var v3 int = 0
    var v4 int = v3
    v5 := make(map[string]bool)
    v6 := make(map[string]uint32)
    v4 = (v4 + (v3 + v4))
    v4 -= (-(v3))
    v7 := GenS1{F0: 345, F1: 1, F2: 2}
    fmt.Println(v4, v3)
    v8 := GenS1{F0: 456, F1: 1, F2: 1}
    v4 = (v4 << ((776) & 63))
    fmt.Println(v4, v4)
    for fi9 := 0; fi9 < 12; fi9++ {
        var v10 int64 = -91
        var v11 float32 = 1.5
        _ = fi9
        _ = v10
        _ = v11
    }
    if (32767 < 255) {
        fmt.Println(v3, v3)
        var v12 int16 = 32767
        _ = v12
    }
    v4 -= -634
    fmt.Println(v4, v4)
    if len(v1) > 0 {
        v1[0] = (v4 - (v3 % (((v4 % ((v3) | 1))) | 1)))
    }
    var v13 int8 = -128
    var v14 int32 = 1
    v13 -= v13
    fmt.Println(v3, v13)
    v13 = (v13 >> ((5) & 63))
    var v15 uint32 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
