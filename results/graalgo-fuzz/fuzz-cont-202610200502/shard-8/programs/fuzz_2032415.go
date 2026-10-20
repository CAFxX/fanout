package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 uint16 }

func main() {
    var v1 uint16 = 2
    var v2 float64 = ((3.14159 * (1e10 / (3.14159 + 1e10))) - -2.25)
    var v3 uint16 = 349
    v2 += v2
    for fi4 := 0; fi4 < 6; fi4++ {
        v3 = (v1 - (v1 & (v1 >> ((51) & 63))))
        fmt.Println(v3, fi4, v3)
        var v5 int64 = -9223372036854775808
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := GenS1{F0: (v1 / (((v1 / ((v1) | 1))) | 1)), F1: 42, F2: (v1 + (v1 / ((v1) | 1)))}
        v2 = (v2 + v2)
        var v8 int8 = -59
        _ = fi6
        _ = v7
        _ = v8
    }
    fmt.Println(v3)
    var v9 bool = true
    var v10 int16 = 1000
    for fi11 := 0; fi11 < 4; fi11++ {
        v9 = (-2.25 <= (-2.25 + (3.14159 * 0.0)))
        var v12 float32 = 0.1
        _ = fi11
        _ = v12
    }
    var v13 uint16 = 2
    var v14 uint64 = 1
    for fi15 := 0; fi15 < 8; fi15++ {
        v1 = v3
        switch (fi15 & (fi15 << ((255) & 63))) {
        case 4294967295:
            fmt.Println(v2)
        default:
            fmt.Println(v13)
        }
        _ = fi15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v9
}
