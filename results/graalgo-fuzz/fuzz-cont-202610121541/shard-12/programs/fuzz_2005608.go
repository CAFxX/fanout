package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
func genFn0(p0 int8) uint32 {
    var v4 float32 = float32(float64(-2.25))
    v5 := []int{-886, 128, -224}
    _ = p0
    _ = v4
    _ = v5
    return 178
}
func genFn1(p0 uint32, p1 int8) int64 {
    if (true && false) {
        v6 := GenS1{F0: -235, F1: 9223372036854775807}
        _ = v6
    }
    p0 -= (p0 + p0)
    _ = p0
    _ = p1
    return -9223372036854775808
}

func main() {
    var v1 int16 = 1000
    v2 := make(map[string]bool)
    var v3 int16 = v1
    if (43 >= 18446744073709551615) {
        v3 = (v1 ^ -358)
        v1 -= (v3 / ((42) | 1))
    }
    v3 += (v1 - (v1 >> ((656) & 63)))
    v1 -= (^(v3))
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 int64 = 42
        fmt.Println(v8, fi7, fi7)
        _ = fi7
        _ = v8
    }
    v9 := v2["foo bar"]
    v10 := []int{127}
    if (962 >= 164) {
        var v11 int16 = (v1 << ((708) & 63))
        _ = v11
    }
    var v12 uint = 160
    var v13 int8 = 100
    v13 = 57
    for fi14 := 0; fi14 < 10; fi14++ {
        var v15 int8 = (v13 << (((v12 | (^(v12)))) & 63))
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 9; fi16++ {
        var v17 int8 = (v13 ^ (v13 | v13))
        fmt.Println(v1)
        _ = fi16
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v9
}
