package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }

func main() {
    var v1 float64 = 1e10
    var v2 int8 = 46
    var v3 float64 = 0.0
    v3 = v3
    v4 := []int{-9223372036854775808, -1, 9223372036854775807, 781}
    for fi5 := 0; fi5 < 6; fi5++ {
        v1 = float64(1.5)
        var v6 float32 = (1e10 + -2.25)
        for fi7 := 0; fi7 < 19; fi7++ {
            var v8 uint = 973
            v3 += v1
            _ = fi7
            _ = v8
        }
        _ = fi5
        _ = v6
    }
    v3 = v1
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v9 float64 = (((((float64(1.5) + v3) / (v1 / (v1 - v3))) - ((v3 + float64((float32(1e10) + 0.0))) - v3)) / v1) * v3)
    v1 = ((1.5 / 0.1) / -2.25)
    if true {
        var v10 uint8 = 171
        _ = v10
    } else {
        var v11 int16 = 631
        _ = v11
    }
    if (false && ((v2 ^ v2) < v2)) {
        if len(v4) > 0 {
            v4[0] = 1
        }
        v2 -= v2
    }
    v9 += (v1 / (float64((float32(float64((1.5 / 3.14159))) - 1e10)) - v9))
    var v12 float32 = 1.5
    v9 = v3
    for fi13 := 0; fi13 < 19; fi13++ {
        v12 -= v12
        switch (^(fi13)) {
        case 0:
            var v14 uint16 = 1
            _ = v14
        default:
            v4 = append(v4, 23)
        }
        v15 := GenS1{F0: 128, F1: 1}
        _ = fi13
        _ = v15
    }
    var v16 int = -9223372036854775808
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v17 uint16 = 0
    var v18 uint = 4294967295
    fmt.Println(v9, v12, v9)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
