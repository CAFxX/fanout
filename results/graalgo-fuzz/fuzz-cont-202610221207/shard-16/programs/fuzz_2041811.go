package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 uint32 }
func genFn0(p0 uint16) int64 {
    for fi7 := 0; fi7 < 11; fi7++ {
        p0 = 32767
        var v8 int8 = -85
        _ = fi7
        _ = v8
    }
    p0 -= (p0 / ((p0) | 1))
    v9 := GenS1{F0: p0}
    _ = p0
    _ = v9
    return -1
}
func genFn1(p0 uint8, p1 uint32, p2 uint8) uint {
    p2 -= p2
    v10 := []int{0, 538}
    p2 += (p2 & 0)
    p2 = p0
    var v11 uint32 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 42
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]int)
    var v4 uint32 = 65535
    var v5 int = -2147483648
    var v6 int = (v5 >> ((238) & 63))
    var v12 uint16 = 2
    for fi13 := 0; fi13 < 20; fi13++ {
        _ = len(v1)
        _ = fi13
    }
    v4 = (^(v4))
    v6 -= v6
    var v14 uint32 = 87
    var v15 int16 = 42
    v14 = (^(v14))
    if (false || true) {
        var v16 bool = false
        v17 := GenS1{F0: 788}
        if false {
            var v18 int64 = -834
            v19 := GenS1{F0: (v12 << ((1) & 63))}
            v20 := GenS1{F0: (^(v12))}
            _ = v18
            _ = v19
            _ = v20
        } else {
            v21 := GenS2{F0: uint32(v12), F1: v14}
            var v22 float64 = ((0.0 / (1e10 * 3.14159)) / 3.14159)
            var v23 float64 = (1e10 - (v22 * v22))
            _ = v21
            _ = v22
            _ = v23
        }
        _ = v16
        _ = v17
    } else {
        fmt.Println(v12, v5, v15)
    }
    v24 := GenS1{F0: v12}
    v25 := []int{128, v6, v6, -633}
    var v26 int = (v6 >> ((1) & 63))
    var v27 uint64 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
