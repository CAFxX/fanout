package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 int64 }
type GenS2 struct { F0 uint32; F1 int }
func genFn0(p0 uint32, p1 int8, p2 uint64) uint {
    if ((3.14159 != (float32(0.0) * (((((0.1 / 3.14159) * (1e10 - (1.5 / float32((-2.25 * ((((1.5 / 1.5) / float64(0.1)) - 3.14159) / (1.5 + 1.5))))))) + (3.14159 * -2.25)) * 3.14159) * 1e10))) || false) {
        fmt.Println(p1, p0, p1)
    } else {
        p1 -= (p1 & p1)
        for fi5 := 0; fi5 < 17; fi5++ {
            p1 += (p1 / (((p1 / ((p1) | 1))) | 1))
            v6 := []int{(-(fi5)), -481}
            _ = fi5
            _ = v6
        }
    }
    var v7 float64 = 1.5
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 2
}
func genFn1(p0 uint16) string {
    fmt.Println(p0)
    v8 := []int{-600, 298, 1}
    p0 = 32767
    _ = p0
    _ = v8
    return ("a" + "hello")
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 string = "a"
    v3 := make(map[int]uint32)
    var v4 int16 = -278
    v9 := GenS1{F0: 0, F1: 734, F2: 1}
    var v10 uint = 2
    if true {
        if (!false) {
            v11 := GenS2{F0: 2, F1: -1}
            _ = v11
        }
    }
    if true {
        switch 128 {
        case 32768:
            v4 += -897
        case 1:
            v2 = v2
        default:
            var v12 int8 = 42
            _ = v12
        }
        var v13 uint64 = 999
        if false {
            delete(v1, 42)
            v4 = 42
            var v14 int8 = 42
            _ = v14
        }
        _ = v13
    }
    fmt.Println(v2, v10)
    v10 += genFn0(0, -1, 439)
    var v15 int32 = -956
    v16 := GenS2{F0: 397, F1: 448}
    v17 := []int{22, 2147483647, 2147483647, -2147483648}
    v15 = (v15 << (((v10 >> (((^(v10))) & 63))) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
