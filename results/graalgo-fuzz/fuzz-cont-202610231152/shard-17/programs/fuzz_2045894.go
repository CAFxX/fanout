package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 uint16 }
func genFn0(p0 uint8) bool {
    p0 += p0
    var v7 float32 = 0.0
    p0 = (^(p0))
    p0 += (p0 | p0)
    v8 := GenS1{F0: false, F1: -81}
    _ = p0
    _ = v7
    _ = v8
    return (42 < 0)
}
func genFn1(p0 uint64) int {
    p0 += (p0 % ((p0) | 1))
    var v9 int32 = -875
    p0 = p0
    v10 := []int{}
    _ = p0
    _ = v10
    _ = v9
    return 0
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint64 = 1
    var v3 int16 = -32768
    v4 := []uint32{}
    var v5 uint64 = v2
    v6 := make(map[int]uint32)
    for fi11 := 0; fi11 < 10; fi11++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi11
    }
    fmt.Println(v2, v5)
    v3 -= (v3 & (v3 / (((v3 << ((18446744073709551615) & 63))) | 1)))
    if (!(-2.25 >= 0.0)) {
        v12 := GenS2{F0: 0}
        fmt.Println(v2, v2)
        _ = v12
    }
    var v13 uint16 = 0
    if false {
        var v14 uint = 2
        if (false && false) {
            v5 = v5
        } else {
            v15 := GenS2{F0: (v13 * v13)}
            _ = v15
        }
        _ = v14
    } else {
        var v16 int32 = 448
        _ = v16
    }
    fmt.Println(v5, v13, v3)
    v17 := GenS1{F0: (-396 == 9223372036854775807), F1: -748}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
