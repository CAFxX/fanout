package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 uint32 }
type GenS2 struct { F0 int8; F1 uint64; F2 uint16 }

func main() {
    v1 := []int{32767, 2147483647, 32767, -2147483648}
    var v2 uint64 = 12
    var v3 float32 = 3.14159
    var v4 int64 = -9223372036854775808
    v5 := GenS1{F0: 255, F1: 185, F2: 4294967295}
    var v6 float32 = v3
    for fi7 := 0; fi7 < 5; fi7++ {
        v2 += (v2 ^ v2)
        v4 = v4
        _ = fi7
    }
    if len(v1) > 0 {
        v1[0] = 32767
    }
    _ = len(v1)
    var v8 int16 = 1000
    v6 = float32(float64(v6))
    if (-379 != 32767) {
        var v9 float64 = 0.1
        v8 -= v8
        _ = v9
    } else {
        var v10 uint8 = 230
        _ = v10
    }
    var v11 uint = 460
    switch "" {
    case "foo bar":
        v12 := GenS2{F0: 56, F1: (v2 * v2), F2: uint16(v11)}
        _ = v12
    default:
        switch 1 {
        case -76:
            fmt.Println(v2, v3)
        case -285:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case -143:
            var v13 float32 = v6
            _ = v13
        }
    }
    var v14 int64 = v4
    var v15 uint = 4294967295
    for fi16 := 0; fi16 < 20; fi16++ {
        switch (fi16 << (((^(v11))) & 63)) {
        case -9223372036854775808:
            v14 += 556
        case 42:
            _ = len(v1)
        default:
            _ = len(v1)
        }
        v17 := GenS2{F0: -128, F1: (v2 - 621), F2: 255}
        var v18 uint16 = 255
        _ = fi16
        _ = v17
        _ = v18
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
