package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64; F1 int16; F2 int }

func main() {
    v1 := make(map[int]bool)
    var v2 int32 = 0
    var v3 uint8 = 2
    var v4 uint = 0
    var v5 float64 = 0.1
    v6 := []uint32{2147483647}
    var v7 int32 = v2
    v4 = (v4 | v4)
    v8 := GenS1{F0: false, F1: v7}
    v2 = 885
    for fi9 := 0; fi9 < 12; fi9++ {
        var v10 int32 = int32(v4)
        v10 -= (-(v7))
        v11 := GenS1{F0: (true && true), F1: (v2 << ((130) & 63))}
        _ = fi9
        _ = v10
        _ = v11
    }
    fmt.Println(v4, v2)
    v7 = (v7 | (v2 - -1))
    if (true || true) {
        var v12 float64 = (1e10 * ((v5 * ((1e10 / (0.1 + v5)) - 3.14159)) + 3.14159))
        _ = v12
    } else {
        v13 := GenS2{F0: 0}
        _ = v13
    }
    v3 = v3
    v14 := GenS3{F0: 558, F1: 1000, F2: -328}
    v6 = append(v6, 65535)
    v4 -= uint(v3)
    v3 = v3
    for fi15 := 0; fi15 < 18; fi15++ {
        v4 = (v4 | (v4 ^ v4))
        _ = fi15
    }
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
