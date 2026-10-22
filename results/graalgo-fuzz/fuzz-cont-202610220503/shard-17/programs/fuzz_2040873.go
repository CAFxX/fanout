package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 int }
func genFn0(p0 uint32) uint32 {
    v7 := []int{-647}
    var v8 uint8 = 0
    _ = p0
    _ = v7
    _ = v8
    return 791
}
func genFn1(p0 uint, p1 int) int64 {
    p1 -= (p1 % ((0) | 1))
    v9 := []int{}
    p0 -= (p0 * (p0 + (p0 | p0)))
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    _ = p0
    _ = p1
    _ = v9
    return -6
}

func main() {
    var v1 uint = 255
    var v2 uint = v1
    var v3 uint16 = 42
    var v4 float32 = (1e10 - ((0.1 / 1.0) + (1.5 / (0.1 / -2.25))))
    v5 := []int{777, 32767, 851}
    v6 := make(map[uint32]uint32)
    var v10 int = -972
    v1 -= (v1 >> ((v1) & 63))
    if (18446744073709551615 != 42) {
        for fi11 := 0; fi11 < 13; fi11++ {
            var v12 int8 = 127
            _ = fi11
            _ = v12
        }
        fmt.Println(v1)
    } else {
        v13 := GenS1{F0: 100, F1: 0}
        var v14 int32 = 0
        fmt.Println(v2)
        _ = v13
        _ = v14
    }
    var v15 bool = (2147483647 <= -444)
    var v16 bool = true
    v17 := GenS1{F0: 0, F1: -166}
    switch v10 {
    case 2147483647:
        if v16 {
            var v18 int8 = 1
            _ = v18
        }
    case 32767:
        switch v10 {
        case 2147483647:
            v19 := GenS1{F0: int8(v10), F1: 805}
            _ = v19
        }
    default:
        if v15 {
            for fi20 := 0; fi20 < 11; fi20++ {
                var v21 uint32 = 423
                _ = fi20
                _ = v21
            }
        } else {
            var v22 int8 = 33
            v23 := GenS1{F0: 0, F1: -64}
            v10 += (v10 - (v10 ^ (v10 << (((v2 * (^(v2)))) & 63))))
            _ = v22
            _ = v23
        }
    }
    v4 += float32(float64(1.5))
    v3 += v3
    v24 := GenS2{F0: (^(v10))}
    v1 += 1
    fmt.Println(v16, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
