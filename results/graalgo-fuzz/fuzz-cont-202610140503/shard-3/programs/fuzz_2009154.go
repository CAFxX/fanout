package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 string }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int32) int {
    p0 += (p0 / ((1) | 1))
    var v5 bool = (265 != 4294967295)
    v6 := []int{127, 328}
    v5 = (!v5)
    var v7 uint16 = 0
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return -359
}
func genFn1(p0 int16, p1 int64, p2 uint) int64 {
    fmt.Println(p0, p2, p1)
    p1 = p1
    var v8 float64 = 0.1
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return -9223372036854775808
}

func main() {
    var v1 uint = 65535
    v2 := make(map[string]int)
    v3 := make(map[int]bool)
    var v4 uint64 = 533
    for fi9 := 0; fi9 < 8; fi9++ {
        v10 := GenS1{F0: 0, F1: 4294967295, F2: "foo bar"}
        _ = fi9
        _ = v10
    }
    if true {
        v11 := GenS2{F0: 127}
        _ = v11
    } else {
        var v12 int64 = 1
        v1 -= v1
        v1 -= v1
        _ = v12
    }
    v1 = (v1 + v1)
    fmt.Println(v1)
    v13 := GenS2{F0: 42}
    for fi14 := 0; fi14 < 13; fi14++ {
        var v15 string = "foo bar"
        var v16 uint32 = 42
        _ = fi14
        _ = v15
        _ = v16
    }
    var v17 float64 = ((0.0 + 0.1) / 1e10)
    v17 -= (v17 * (0.0 - 0.1))
    v18 := []int{genFn0(-1)}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
