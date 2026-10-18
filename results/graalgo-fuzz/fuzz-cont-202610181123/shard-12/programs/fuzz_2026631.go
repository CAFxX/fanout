package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 uint8; F1 uint8; F2 uint16 }
func genFn0(p0 int64, p1 uint16, p2 uint32) uint32 {
    fmt.Println(p0, p2, p0)
    var v4 bool = (42 > 77)
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return (p2 / (((^(p2))) | 1))
}
func genFn1() int16 {
    v5 := GenS1{F0: 453}
    var v6 int16 = 0
    _ = v5
    _ = v6
    return v6
}

func main() {
    var v1 float32 = (3.14159 * 0.1)
    var v2 int = 270
    v3 := make(map[string]uint32)
    switch (v2 / ((v2) | 1)) {
    case -9223372036854775808:
        v7 := GenS1{F0: 2}
        _ = v7
    case 499:
        var v8 int8 = -128
        _ = v8
    default:
        v9 := GenS1{F0: 42}
        _ = v9
    }
    v10 := GenS2{F0: 42, F1: -1}
    v1 = 3.14159
    var v11 bool = true
    v1 -= 1e10
    v2 = (v2 * (v2 ^ (^(v2))))
    for fi12 := 0; fi12 < 16; fi12++ {
        v2 += -9223372036854775808
        v13 := GenS1{F0: 255}
        v14 := GenS3{F0: 149, F1: 64, F2: 469}
        _ = fi12
        _ = v13
        _ = v14
    }
    if v11 {
        for fi15 := 0; fi15 < 9; fi15++ {
            var v16 int32 = 2147483647
            var v17 int16 = -663
            _ = fi15
            _ = v16
            _ = v17
        }
        var v18 int32 = -487
        _ = v18
    }
    var v19 float32 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v3
}
