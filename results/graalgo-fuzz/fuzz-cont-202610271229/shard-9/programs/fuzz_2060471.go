package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 uint16 }
func genFn0(p0 uint64, p1 int16, p2 int) uint32 {
    fmt.Println(p1, p0, p2)
    for fi6 := 0; fi6 < 2; fi6++ {
        if ((1 > 1) && (float32(((((float64(float32(0.1)) - ((1e10 / 1e10) + -2.25)) + (3.14159 * (float64(0.0) / 3.14159))) * (1.5 + 0.1)) * 1.5)) >= 1e10)) {
            var v7 uint = 18446744073709551615
            _ = v7
        }
        p0 = 2
        _ = fi6
    }
    switch (p2 + p2) {
    case 504:
        v8 := []int{p2}
        _ = v8
    case 4294967295:
        v9 := []int{p2, -9223372036854775808}
        _ = v9
    case 127:
        var v10 int = (^(p2))
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    return 318
}
func genFn1(p0 int8, p1 int) uint64 {
    p1 = -75
    var v11 int16 = 42
    p1 -= (p1 & -785)
    if true {
        var v12 int64 = 1
        v13 := []int{(p1 >> ((255) & 63)), (p1 | (p1 | p1))}
        v12 -= 0
        _ = v12
        _ = v13
    } else {
        var v14 uint64 = 4294967295
        v15 := GenS1{F0: 255}
        var v16 float64 = 0.0
        _ = v14
        _ = v15
        _ = v16
    }
    var v17 string = "hello"
    _ = p0
    _ = p1
    _ = v11
    _ = v17
    return 833
}

func main() {
    var v1 float64 = 0.0
    v2 := make(map[string]int)
    var v3 bool = false
    v4 := []uint32{65535, 2147483648, 2}
    var v5 uint8 = 49
    v3 = true
    for fi18 := 0; fi18 < 12; fi18++ {
        v4 = append(v4, 2)
        _ = fi18
    }
    v3 = v3
    v19 := GenS1{F0: 681}
    if ((v5 > (v5 % ((v5) | 1))) || (42 != 0)) {
        var v20 uint = 255
        _ = v20
    }
    var v21 int16 = 0
    var v22 uint = 409
    var v23 int32 = 1
    v3 = false
    v5 = v5
    if v3 {
        v24 := GenS2{F0: v21, F1: 255}
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
