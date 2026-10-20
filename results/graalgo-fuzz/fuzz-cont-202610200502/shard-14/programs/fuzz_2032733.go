package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 uint32; F2 bool }
func genFn0(p0 bool, p1 int8, p2 int) int8 {
    for fi7 := 0; fi7 < 3; fi7++ {
        p0 = p0
        _ = fi7
    }
    v8 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 119
}

func main() {
    var v1 float64 = 1.5
    v2 := make(map[int]string)
    var v3 float32 = 0.1
    var v4 int64 = 0
    v5 := []uint32{2, 265}
    v6 := []uint32{}
    v9 := GenS1{F0: 0}
    fmt.Println(v1)
    v4 = -1
    if (v1 > v1) {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        fmt.Println(v1, v1)
        v3 = float32(float64(-2.25))
    } else {
        v10 := GenS1{F0: -658}
        _ = v10
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v11 := GenS2{F0: v4, F1: 0, F2: (!(255 <= 42))}
    var v12 uint64 = 137
    v3 -= v3
    v3 += -2.25
    v13 := GenS1{F0: 88}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
