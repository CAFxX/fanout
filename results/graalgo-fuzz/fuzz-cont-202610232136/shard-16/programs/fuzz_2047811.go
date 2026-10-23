package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 int8 }
func genFn0() uint32 {
    var v6 float32 = -2.25
    var v7 float64 = 1e10
    v7 = float64(float32(float64((v6 / 1.5))))
    var v8 float64 = (((((v7 + v7) / (v7 * v7)) - 0.0) - 1.5) - -2.25)
    _ = v6
    _ = v7
    _ = v8
    return 2147483647
}
func genFn1(p0 int) int64 {
    var v9 int16 = 1000
    v10 := []int{(p0 >> ((2) & 63)), p0, p0, (p0 * (p0 - (p0 | p0)))}
    v9 -= 0
    _ = len(v10)
    _ = p0
    _ = v10
    _ = v9
    return 42
}

func main() {
    var v1 uint32 = 818
    var v2 int8 = 100
    var v3 int8 = v2
    v4 := []int{0, 42, 545}
    var v5 int = 4294967295
    v2 -= 100
    var v11 int32 = 493
    _ = len(v4)
    v12 := GenS1{F0: 127}
    var v13 int = v5
    for fi14 := 0; fi14 < 7; fi14++ {
        for fi15 := 0; fi15 < 5; fi15++ {
            v5 += (fi14 | (fi15 / (((v5 * (-(v5)))) | 1)))
            _ = fi15
        }
        v2 -= (v2 >> ((860) & 63))
        _ = fi14
    }
    v5 = v5
    v5 = (v5 | v13)
    if true {
        for fi16 := 0; fi16 < 2; fi16++ {
            v17 := GenS2{F0: 42, F1: (v3 / (((v3 | v3)) | 1))}
            _ = fi16
            _ = v17
        }
        v2 += v3
    } else {
        v18 := GenS2{F0: 1, F1: (v2 % (((v3 << ((131) & 63))) | 1))}
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
