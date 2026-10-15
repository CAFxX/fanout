package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 uint32; F2 string }

func main() {
    var v1 int64 = 42
    var v2 float32 = 0.1
    var v3 float64 = ((1e10 / 1.0) * 3.14159)
    var v4 uint8 = 255
    var v5 uint16 = 902
    v6 := []int{}
    v7 := GenS1{F0: 2147483647}
    v3 -= ((0.1 - v3) - (v3 / v3))
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    for fi8 := 0; fi8 < 18; fi8++ {
        v1 += (-(v1))
        v6 = append(v6, (^(fi8)))
        _ = fi8
    }
    if len(v6) > 0 {
        v6[0] = 32767
    }
    switch 9223372036854775807 {
    case 32767:
        var v9 bool = ((!(-131 >= -9223372036854775808)) || false)
        _ = v9
    case 1:
        v6 = append(v6, -2147483648)
    case -915:
        var v10 int64 = (-(v1))
        _ = v10
    default:
        for fi11 := 0; fi11 < 2; fi11++ {
            if false {
                v4 -= v4
                v12 := GenS1{F0: 670}
                v13 := GenS2{F0: 42, F1: 1, F2: "\n\t"}
                _ = v12
                _ = v13
            }
            fmt.Println(v5)
            _ = fi11
        }
    }
    v2 = v2
    fmt.Println(v5, v4, v3)
    v2 += 0.1
    v14 := GenS1{F0: 0}
    v5 = (v5 + (^(v5)))
    var v15 int16 = 42
    fmt.Println(v15)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
