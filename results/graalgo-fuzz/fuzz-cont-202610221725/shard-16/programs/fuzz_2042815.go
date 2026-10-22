package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 string }
type GenS2 struct { F0 uint8; F1 uint8; F2 bool }
type GenS3 struct { F0 int16 }

func main() {
    var v1 int16 = 903
    var v2 int = -9223372036854775808
    var v3 int = v2
    var v4 int32 = 0
    switch (v2 >> ((2) & 63)) {
    case -1:
        var v5 float64 = -2.25
        _ = v5
    }
    v6 := GenS1{F0: -32768, F1: 100, F2: ""}
    if ((-2.25 + ((0.1 * -2.25) * (float32(1e10) + 1e10))) > 0.1) {
        v7 := GenS2{F0: 210, F1: 128, F2: false}
        _ = v7
    }
    var v8 float64 = (3.14159 + (0.1 / 1.0))
    v1 += (-(v1))
    switch (v2 - (v2 & v2)) {
    case -516:
        v3 = (v2 ^ (v3 ^ 9223372036854775807))
    case -89:
        if false {
            for fi9 := 0; fi9 < 7; fi9++ {
                v3 = v2
                fmt.Println(fi9, fi9)
                _ = fi9
            }
        }
    default:
        v4 += (v4 >> ((0) & 63))
    }
    var v10 int64 = 14
    for fi11 := 0; fi11 < 20; fi11++ {
        v1 = -140
        v8 = float64(-2.25)
        _ = fi11
    }
    v12 := []int{128, int(v4), v3}
    v1 = (v1 - 189)
    var v13 uint32 = 2
    var v14 int32 = (v4 ^ v4)
    for fi15 := 0; fi15 < 12; fi15++ {
        v2 = 827
        _ = fi15
    }
    v16 := GenS3{F0: (v1 % ((v1) | 1))}
    v4 = (v4 | v4)
    v17 := GenS3{F0: (v1 - 172)}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v8
}
