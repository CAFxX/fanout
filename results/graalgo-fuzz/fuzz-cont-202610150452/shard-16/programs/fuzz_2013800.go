package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 uint }
func genFn0(p0 uint32, p1 int64, p2 int32) uint16 {
    v6 := GenS1{F0: p1, F1: 0}
    var v7 uint16 = 32767
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return v7
}
func genFn1(p0 int8) int16 {
    v8 := []int{int(p0), -2147483648}
    v9 := GenS1{F0: 715, F1: p0}
    _ = p0
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 uint8 = 42
    var v2 uint64 = 42
    v3 := make(map[uint32]bool)
    var v4 int32 = -2147483648
    var v5 int16 = 643
    v1 -= (v1 - uint8(v2))
    v10 := []int{1, 2147483647, 32768}
    v4 += 42
    fmt.Println(v5, v5, v1)
    v2 = (v2 + 1)
    for fi11 := 0; fi11 < 14; fi11++ {
        var v12 int8 = 0
        if (0.1 != float64((1e10 / 1e10))) {
            v4 -= v4
            var v13 float32 = float32(1.5)
            var v14 float32 = (v13 - v13)
            _ = v13
            _ = v14
        }
        _ = fi11
        _ = v12
    }
    var v15 uint64 = (v2 >> ((2) & 63))
    v16 := GenS2{F0: 0}
    v17 := GenS2{F0: 65535}
    var v18 uint32 = 2
    fmt.Println(v1, v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
