package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
func genFn0(p0 int16, p1 int64, p2 uint) int64 {
    for fi6 := 0; fi6 < 17; fi6++ {
        v7 := GenS1{F0: 32, F1: 0}
        p0 -= -212
        v8 := []int{(fi6 ^ (^(fi6))), 543, (fi6 - (fi6 % ((fi6) | 1)))}
        _ = fi6
        _ = v7
        _ = v8
    }
    p1 += 0
    v9 := GenS1{F0: 50, F1: 0}
    p0 += 450
    v10 := []int{127, 2147483647}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return p1
}

func main() {
    var v1 float64 = (1e10 * 1e10)
    v2 := make(map[uint32]uint32)
    v3 := []uint32{42, 961}
    v4 := []int{4294967295, 127, 32768}
    var v5 int64 = 1
    fmt.Println(v1, v1)
    for fi11 := 0; fi11 < 10; fi11++ {
        fmt.Println(v1)
        v5 += genFn0(971, v5, 65535)
        _ = fi11
    }
    var v12 int16 = 811
    for fi13 := 0; fi13 < 10; fi13++ {
        var v14 float32 = float32(float64(1.5))
        _ = fi13
        _ = v14
    }
    _ = len(v2)
    var v15 int8 = -128
    v16 := v2[65535]
    for fi17 := 0; fi17 < 1; fi17++ {
        var v18 uint16 = 42
        _ = fi17
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
