package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int32 }
type GenS2 struct { F0 uint64 }
func genFn0() int8 {
    var v7 int32 = -1
    v7 = (v7 << ((722) & 63))
    v8 := GenS1{F0: 18446744073709551615, F1: 32767, F2: v7}
    _ = v7
    _ = v8
    return 41
}
func genFn1(p0 int8, p1 uint) uint8 {
    for fi9 := 0; fi9 < 4; fi9++ {
        fmt.Println(fi9, p1, fi9)
        p1 += p1
        _ = fi9
    }
    v10 := []int{-2147483648, -93, 32767, 2147483647}
    for fi11 := 0; fi11 < 8; fi11++ {
        _ = len(v10)
        _ = fi11
    }
    _ = len(v10)
    _ = p0
    _ = p1
    _ = v10
    return 210
}

func main() {
    var v1 int32 = -1
    v2 := []uint32{4294967295, 315, 729}
    v3 := []string{}
    var v4 int32 = (v1 << ((405) & 63))
    var v5 float64 = (0.1 - 0.1)
    v6 := make(map[uint32]bool)
    var v12 uint64 = 42
    v1 += 1
    var v13 uint32 = 237
    for fi14 := 0; fi14 < 17; fi14++ {
        v15 := GenS2{F0: (^(v12))}
        _ = fi14
        _ = v15
    }
    v16 := GenS2{F0: (v12 * (v12 * v12))}
    v4 = v1
    if false {
        var v17 float64 = (-2.25 / v5)
        v4 = v4
        _ = v17
    }
    fmt.Println(v4, v4, v13)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
