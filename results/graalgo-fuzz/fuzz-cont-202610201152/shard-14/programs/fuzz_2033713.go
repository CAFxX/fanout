package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 int64 }

func main() {
    var v1 int64 = 0
    var v2 int16 = 834
    var v3 uint32 = 42
    var v4 uint16 = 65535
    for fi5 := 0; fi5 < 1; fi5++ {
        v6 := GenS1{F0: v1, F1: 100, F2: (v1 % ((v1) | 1))}
        switch (^(fi5)) {
        case 2147483647:
            var v7 int32 = 0
            _ = v7
        case 0:
            v3 -= v3
        default:
            var v8 int32 = 0
            _ = v8
        }
        _ = fi5
        _ = v6
    }
    v9 := []int{-2147483648, 1}
    if (v2 < (v2 | v2)) {
        v2 += (v2 % ((v2) | 1))
        for fi10 := 0; fi10 < 10; fi10++ {
            v9 = append(v9, 0)
            var v11 int32 = int32(v2)
            _ = fi10
            _ = v11
        }
    }
    v12 := GenS1{F0: v1, F1: -1, F2: (v1 << ((2) & 63))}
    v1 = (-(v1))
    if len(v9) > 0 {
        v9[0] = 32767
    }
    var v13 float32 = ((((float32(1e10) / ((1.5 / 0.1) + (3.14159 / 3.14159))) * 1e10) * 0.1) + 1.5)
    for fi14 := 0; fi14 < 17; fi14++ {
        var v15 int64 = -9223372036854775808
        _ = fi14
        _ = v15
    }
    if len(v9) > 0 {
        v9[0] = 236
    }
    if len(v9) > 0 {
        v9[0] = 509
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
