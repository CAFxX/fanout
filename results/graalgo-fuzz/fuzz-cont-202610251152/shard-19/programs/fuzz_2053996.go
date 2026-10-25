package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 int8; F1 uint32; F2 uint16 }

func main() {
    var v1 int = 4294967295
    var v2 float32 = 1.5
    var v3 float64 = 1.5
    v4 := []int{}
    if len(v4) > 0 {
        v4[0] = (v1 ^ (v1 >> ((255) & 63)))
    }
    if (!true) {
        switch v1 {
        case 127:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case 947:
            for fi5 := 0; fi5 < 13; fi5++ {
                fmt.Println(v3, v1, v2)
                _ = fi5
            }
        case -2147483648:
            v6 := GenS1{F0: 1, F1: 4294967295}
            _ = v6
        default:
            v1 = (v1 + (v1 - (v1 | v1)))
        }
        v7 := GenS2{F0: 42, F1: 94, F2: 32767}
        _ = v7
    }
    fmt.Println(v1, v3, v2)
    var v8 int64 = -378
    if true {
        for fi9 := 0; fi9 < 16; fi9++ {
            v1 = (v1 / (((v1 >> ((18446744073709551615) & 63))) | 1))
            var v10 uint = 65535
            v4 = append(v4, (fi9 / ((v1) | 1)))
            _ = fi9
            _ = v10
        }
        if false {
            var v11 float64 = (v3 + v3)
            _ = v11
        }
    }
    v1 += (^(v1))
    v12 := GenS2{F0: -99, F1: 288, F2: 65535}
    for fi13 := 0; fi13 < 11; fi13++ {
        var v14 uint64 = 686
        v2 = (v2 - v2)
        _ = fi13
        _ = v14
    }
    fmt.Println(v8, v1, v2)
    var v15 float32 = float32(((0.1 - (((v3 / (v3 * v3)) + v3) / v3)) + (float64((((v2 - v2) * (float32((float64(1e10) - v3)) + (-2.25 - v2))) + v2)) * ((0.1 / float64(v2)) + v3))))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
