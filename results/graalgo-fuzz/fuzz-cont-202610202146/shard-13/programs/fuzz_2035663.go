package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int }

func main() {
    v1 := []uint32{42, 65535}
    var v2 int16 = -32768
    v3 := make(map[uint32]bool)
    var v4 float32 = (0.1 + 1e10)
    var v5 int16 = -1
    var v6 int32 = 2147483647
    var v7 int32 = 133
    fmt.Println(v5)
    var v8 uint64 = 730
    var v9 uint8 = 89
    var v10 float64 = (1e10 / 1.0)
    var v11 uint16 = 32767
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v6 = (v7 + (v6 * v6))
    var v12 uint64 = v8
    var v13 uint16 = v11
    for fi14 := 0; fi14 < 12; fi14++ {
        v15 := GenS1{F0: 32768, F1: v11, F2: fi14}
        fmt.Println(v5, v10, v6)
        v16 := GenS1{F0: (fi14 ^ -9223372036854775808), F1: (v13 | v13), F2: (fi14 << ((65535) & 63))}
        _ = fi14
        _ = v15
        _ = v16
    }
    v8 += v8
    var v17 uint = 4294967295
    if (65535 != 1) {
        var v18 int16 = v5
        var v19 int8 = -128
        switch 0 {
        case -841:
            if true {
                v20 := GenS1{F0: 127, F1: 53, F2: 734}
                delete(v3, 596)
                _ = v20
            } else {
                v19 += v19
            }
        case 860:
            v21 := GenS1{F0: -695, F1: (v11 << (((v17 - v17)) & 63)), F2: -2147483648}
            _ = v21
        default:
            var v22 float32 = v4
            _ = v22
        }
        _ = v18
        _ = v19
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
