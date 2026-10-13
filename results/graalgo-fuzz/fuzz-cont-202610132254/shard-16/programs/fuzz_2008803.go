package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
func genFn0(p0 int32, p1 int8) int {
    p1 += (-(p1))
    for fi7 := 0; fi7 < 7; fi7++ {
        p1 = (p1 & p1)
        _ = fi7
    }
    p0 = p0
    p1 -= p1
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    return 485
}

func main() {
    v1 := []uint32{}
    v2 := make(map[uint32]bool)
    var v3 uint16 = 566
    var v4 int8 = -128
    v5 := make(map[string]uint32)
    v6 := []int{42, 9223372036854775807}
    v3 -= 139
    fmt.Println(v4)
    v4 = (v4 >> ((255) & 63))
    if true {
        var v8 int = 479
        v8 -= -9223372036854775808
        v3 += 2
        _ = v8
    } else {
        v9 := GenS1{F0: 749, F1: 687}
        if true {
            v3 -= 271
        }
        var v10 int8 = v4
        _ = v10
        _ = v9
    }
    fmt.Println(v4)
    fmt.Println(v3)
    v11 := GenS1{F0: 396, F1: 611}
    v3 += v3
    v12 := GenS1{F0: -591, F1: 18446744073709551615}
    var v13 uint64 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
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
}
