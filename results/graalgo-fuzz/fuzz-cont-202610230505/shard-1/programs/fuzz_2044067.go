package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 int }
type GenS2 struct { F0 uint }
func genFn0(p0 uint32, p1 uint64, p2 int8) bool {
    if true {
        fmt.Println(p0, p1, p2)
    } else {
        switch 446 {
        case 42:
            fmt.Println(p2, p1, p0)
        default:
            fmt.Println(p0)
        }
    }
    p1 = (p1 >> ((897) & 63))
    switch 0 {
    case -2147483648:
        fmt.Println(p2, p1)
    default:
        v6 := GenS1{F0: p1, F1: 0, F2: 127}
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    return true
}
func genFn1(p0 bool) int32 {
    p0 = p0
    if (1.5 >= 1e10) {
        p0 = p0
        v7 := GenS1{F0: 0, F1: 42, F2: 127}
        if p0 {
            v8 := GenS2{F0: 540}
            var v9 int32 = 1
            _ = v8
            _ = v9
        }
        _ = v7
    }
    v10 := GenS2{F0: 1}
    p0 = p0
    v11 := []int{42, 726}
    _ = p0
    _ = v10
    _ = v11
    return -871
}

func main() {
    v1 := []int{42, 127, 42}
    var v2 int8 = 69
    var v3 uint = 198
    v4 := make(map[uint32]uint32)
    v5 := make(map[int]int)
    v3 -= 18446744073709551615
    for fi12 := 0; fi12 < 19; fi12++ {
        _ = len(v1)
        var v13 uint = (v3 * v3)
        v2 += v2
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 4; fi14++ {
        v3 = (v3 & v3)
        var v15 int32 = 241
        _ = fi14
        _ = v15
    }
    var v16 uint16 = 1
    v16 = (v16 % ((42) | 1))
    var v17 int = 2147483647
    var v18 int16 = 0
    for fi19 := 0; fi19 < 1; fi19++ {
        _ = len(v1)
        _ = fi19
    }
    v2 += v2
    var v20 int = v17
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
