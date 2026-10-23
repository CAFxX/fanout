package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
func genFn0(p0 uint16, p1 int, p2 int) int8 {
    for fi6 := 0; fi6 < 6; fi6++ {
        p1 -= (p2 % ((p2) | 1))
        switch (p1 ^ fi6) {
        case 9223372036854775807:
            var v7 uint = 18446744073709551615
            _ = v7
        }
        p2 -= (p2 * (^(p1)))
        _ = fi6
    }
    if true {
        v8 := []int{32768, p2, p2, (p2 & p1)}
        for fi9 := 0; fi9 < 15; fi9++ {
            _ = len(v8)
            var v10 float64 = 0.0
            _ = fi9
            _ = v10
        }
        _ = v8
    } else {
        switch 32768 {
        case 0:
            p2 += p2
        }
        var v11 uint16 = (p0 >> ((834) & 63))
        for fi12 := 0; fi12 < 7; fi12++ {
            var v13 float64 = 0.1
            _ = fi12
            _ = v13
        }
        _ = v11
    }
    var v14 uint = 65535
    var v15 uint8 = 225
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    return -12
}
func genFn1() uint16 {
    v16 := GenS1{F0: 32767, F1: 321}
    var v17 float64 = ((3.14159 + 1.5) / -2.25)
    _ = v16
    _ = v17
    return 1
}

func main() {
    v1 := []uint32{}
    v2 := make(map[uint32]int)
    var v3 bool = true
    v4 := make(map[uint32]int)
    v5 := make(map[int]bool)
    v3 = true
    var v18 int8 = 78
    v3 = false
    var v19 int = -870
    for fi20 := 0; fi20 < 14; fi20++ {
        var v21 int32 = 42
        _ = fi20
        _ = v21
    }
    var v22 int32 = -1
    for fi23 := 0; fi23 < 13; fi23++ {
        var v24 int8 = 100
        var v25 bool = v3
        _ = fi23
        _ = v24
        _ = v25
    }
    var v26 int32 = (v22 << ((0) & 63))
    v22 += (v22 - 444)
    v18 = 89
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
