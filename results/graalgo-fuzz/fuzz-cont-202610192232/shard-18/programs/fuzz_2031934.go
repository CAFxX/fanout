package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 uint16; F1 uint; F2 int }

func main() {
    var v1 int16 = 39
    var v2 uint32 = 1
    var v3 float32 = (1.5 - (3.14159 / 1.0))
    v4 := make(map[int]bool)
    var v5 float64 = (0.0 - (((-2.25 * (3.14159 - 0.1)) - 1e10) - ((0.1 * ((1e10 / 3.14159) + 0.1)) * 3.14159)))
    for fi6 := 0; fi6 < 12; fi6++ {
        var v7 uint8 = 127
        var v8 int32 = 1
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: -128, F1: 118}
    for fi10 := 0; fi10 < 1; fi10++ {
        fmt.Println(v3)
        _ = fi10
    }
    var v11 uint32 = 2
    var v12 uint64 = 0
    v13 := v4[2147483647]
    for fi14 := 0; fi14 < 7; fi14++ {
        v15 := GenS1{F0: int8(v1), F1: 100}
        _ = fi14
        _ = v15
    }
    if (0 <= (v1 >> ((18446744073709551615) & 63))) {
        if v13 {
            for fi16 := 0; fi16 < 4; fi16++ {
                v4[(fi16 & 2147483647)] = v13
                var v17 int8 = 0
                v18 := GenS2{F0: 1, F1: 255, F2: (fi16 >> ((42) & 63))}
                _ = fi16
                _ = v17
                _ = v18
            }
            v5 = 0.1
        }
        v19 := GenS1{F0: 121, F1: 100}
        var v20 float64 = (v5 - ((v5 - float64(v3)) * float64(v3)))
        _ = v19
        _ = v20
    } else {
        if ((9223372036854775807 < 32767) && true) {
            var v21 int8 = 0
            v1 += (v1 & v1)
            var v22 int64 = -421
            _ = v21
            _ = v22
        }
        v3 = 0.0
        switch -367 {
        case -1:
            v2 = (v2 >> ((1) & 63))
        case 1:
            _ = len(v4)
        }
    }
    v3 += (0.0 + 0.1)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
