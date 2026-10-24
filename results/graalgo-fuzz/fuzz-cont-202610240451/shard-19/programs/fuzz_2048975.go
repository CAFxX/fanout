package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
func genFn0() uint32 {
    v6 := []int{}
    v6 = append(v6, -888)
    _ = v6
    return 0
}
func genFn1(p0 int, p1 bool, p2 uint64) uint8 {
    for fi7 := 0; fi7 < 3; fi7++ {
        v8 := GenS1{F0: (0 == -1), F1: 233}
        var v9 uint8 = 127
        if p1 {
            v9 = v9
            var v10 int16 = -610
            var v11 int = (p0 % (((fi7 | (p0 - (p0 + p0)))) | 1))
            _ = v10
            _ = v11
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    var v12 uint64 = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 197
}

func main() {
    var v1 int = 636
    v2 := make(map[int]bool)
    var v3 bool = (false && (1 <= -851))
    v4 := make(map[uint32]int)
    v5 := make(map[uint32]int)
    fmt.Println(v3, v1)
    var v13 int16 = 0
    for fi14 := 0; fi14 < 8; fi14++ {
        for fi15 := 0; fi15 < 12; fi15++ {
            var v16 int32 = -2147483648
            v4[genFn0()] = fi14
            v17 := []int{}
            _ = fi15
            _ = v16
            _ = v17
        }
        var v18 int16 = (v13 / ((-468) | 1))
        v18 = v18
        _ = fi14
        _ = v18
    }
    v19 := GenS1{F0: (v3 && false), F1: 1}
    for fi20 := 0; fi20 < 11; fi20++ {
        v1 = 1
        _ = fi20
    }
    for fi21 := 0; fi21 < 3; fi21++ {
        v13 = (-(v13))
        _ = fi21
    }
    v22 := []int{(v1 - v1), v1, v1, (v1 % (((-(v1))) | 1))}
    var v23 float32 = (float32(1.5) + 1.5)
    v23 = (v23 / 1e10)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
