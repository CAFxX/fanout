package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 int64; F1 uint16; F2 int8 }

func main() {
    var v1 int8 = -1
    v2 := []int{32767, -836, -9223372036854775808, -2147483648}
    var v3 int = -576
    for fi4 := 0; fi4 < 2; fi4++ {
        var v5 uint16 = 781
        _ = len(v2)
        _ = fi4
        _ = v5
    }
    v1 += v1
    if (true || false) {
        v3 = 50
    } else {
        v3 += 4294967295
    }
    v3 = v3
    if len(v2) > 0 {
        v2[0] = v3
    }
    var v6 uint8 = 1
    v7 := GenS1{F0: 42}
    v1 = v1
    fmt.Println(v3, v1, v3)
    v3 -= 2147483647
    v8 := GenS2{F0: 100, F1: (v6 * v6)}
    if false {
        v9 := GenS2{F0: 128, F1: 127}
        _ = v9
    } else {
        if len(v2) > 0 {
            v2[0] = v3
        }
        var v10 int8 = -128
        _ = v10
    }
    if (1.5 >= (3.14159 * -2.25)) {
        v1 -= (v1 + v1)
        var v11 uint16 = 133
        _ = v11
    }
    v3 = (v3 >> ((0) & 63))
    v12 := GenS3{F0: 0, F1: 32767, F2: v1}
    switch (v3 + -941) {
    case 127:
        v3 += v3
    case -889:
        v1 = v1
    case -517:
        if (!true) {
            v2 = append(v2, 864)
            var v13 float64 = 0.1
            fmt.Println(v3)
            _ = v13
        } else {
            var v14 float64 = 1e10
            v3 -= (^(v3))
            _ = v14
        }
    default:
        fmt.Println(v6, v6)
    }
    if len(v2) > 0 {
        v2[0] = (v3 ^ (v3 * v3))
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
