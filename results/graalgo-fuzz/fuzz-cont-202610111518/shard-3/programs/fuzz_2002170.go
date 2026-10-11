package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 int64; F1 int64 }

func main() {
    var v1 int8 = 78
    var v2 uint64 = 146
    var v3 float64 = 0.1
    v4 := []uint32{2147483648, 42, uint32(v2), 42}
    var v5 int = -9223372036854775808
    v1 = 0
    var v6 uint8 = 1
    if (((!(-892 != -1)) && true) || true) {
        var v7 float64 = 3.14159
        for fi8 := 0; fi8 < 16; fi8++ {
            for fi9 := 0; fi9 < 12; fi9++ {
                v10 := GenS1{F0: 65535, F1: 330}
                _ = fi9
                _ = v10
            }
            var v11 bool = true
            _ = fi8
            _ = v11
        }
        _ = v7
    } else {
        var v12 uint64 = v2
        _ = v12
    }
    if (742 < 2147483648) {
        var v13 uint16 = 285
        var v14 bool = true
        _ = v13
        _ = v14
    } else {
        var v15 bool = ("a" < "")
        _ = v15
    }
    v4 = append(v4, 2147483648)
    v3 = 0.1
    v2 -= 539
    var v16 int = (-(v5))
    if (int8(v5) <= 0) {
        _ = len(v4)
        var v17 bool = false
        v18 := GenS1{F0: 435, F1: 0}
        _ = v17
        _ = v18
    }
    v19 := GenS1{F0: 65535, F1: 604}
    if len(v4) > 0 {
        v4[0] = 4294967295
    }
    v4 = append(v4, uint32(v6))
    if len(v4) > 0 {
        v4[0] = 42
    }
    for fi20 := 0; fi20 < 15; fi20++ {
        fmt.Println(fi20, fi20)
        var v21 uint8 = 42
        v22 := GenS2{F0: -697, F1: 0}
        _ = fi20
        _ = v21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
