package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 uint32; F1 bool; F2 int64 }
type GenS4 struct { F0 uint16; F1 uint }

func main() {
    var v1 uint64 = 4294967295
    v2 := []string{}
    v3 := make(map[int]bool)
    if true {
        v1 += (v1 | (v1 / (((v1 ^ v1)) | 1)))
    }
    v4 := GenS1{F0: 387, F1: 0}
    var v5 int16 = 765
    if true {
        v6 := GenS2{F0: 255, F1: ((v5 & (v5 >> ((579) & 63))) >= -32768)}
        _ = v6
    }
    v1 -= 18446744073709551615
    if (-1 < -531) {
        _ = len(v2)
    }
    if (1 >= 42) {
        v5 = (v5 & v5)
        fmt.Println(v1)
        if (!true) {
            v7 := GenS3{F0: 900, F1: (0.1 != (0.0 / ((1e10 * (float32(0.1) * -2.25)) / 3.14159))), F2: 42}
            _ = v7
        } else {
            fmt.Println(v5)
            fmt.Println(v1)
        }
    } else {
        var v8 uint32 = 494
        var v9 int64 = -9223372036854775808
        v8 = (v8 * 940)
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        var v11 bool = (42 > -2)
        for fi12 := 0; fi12 < 20; fi12++ {
            v1 -= 4294967295
            var v13 uint8 = 255
            _ = fi12
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    v14 := v3[544]
    v5 = v5
    v15 := GenS1{F0: 4294967295, F1: 42}
    v16 := GenS1{F0: 255, F1: 42}
    v17 := GenS4{F0: 1, F1: 42}
    fmt.Println(v1, v5, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
