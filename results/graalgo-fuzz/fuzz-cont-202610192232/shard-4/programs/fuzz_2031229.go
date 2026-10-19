package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
func genFn0(p0 uint, p1 string) int32 {
    var v5 int32 = 0
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 int16 = -32768
        var v8 uint32 = 334
        var v9 int32 = 619
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    var v10 uint8 = 0
    _ = p0
    _ = p1
    _ = v10
    _ = v5
    return (-(v5))
}
func genFn1(p0 uint32) bool {
    v11 := []int{-152, 0, 32768, 32768}
    for fi12 := 0; fi12 < 18; fi12++ {
        fmt.Println(p0, fi12)
        _ = fi12
    }
    _ = p0
    _ = v11
    return (!true)
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := make(map[int]int)
    var v3 int32 = 1
    var v4 int = -478
    v1 = v1
    v13 := []int{(v4 - 42), (v4 | v4)}
    fmt.Println(v1, v1)
    var v14 uint = 255
    v2[v4] = v4
    v15 := GenS1{F0: (v4 << (((v14 + (v14 % (((v14 + (v14 ^ v14))) | 1)))) & 63)), F1: -128}
    for fi16 := 0; fi16 < 9; fi16++ {
        var v17 float64 = 1.5
        var v18 float32 = 1e10
        _ = fi16
        _ = v17
        _ = v18
    }
    var v19 uint32 = 184
    fmt.Println(v4, v14, v3)
    v20 := GenS1{F0: v4, F1: 79}
    v1 += (v1 >> ((v14) & 63))
    fmt.Println(v19, v1)
    if (true && ((!true) || (!((0.0 * 1.5) < 0.1)))) {
        v19 += 58
        if (255 >= 42) {
            v3 -= v3
        } else {
            v4 -= v4
            fmt.Println(v4, v4, v14)
        }
    } else {
        var v21 int16 = 42
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
