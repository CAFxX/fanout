package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 uint32 }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 int32; F1 uint32; F2 int64 }
func genFn0(p0 string) uint {
    var v5 int64 = -69
    v6 := []int{-1, 4294967295, 4294967295}
    v5 -= -9223372036854775808
    _ = p0
    _ = v5
    _ = v6
    return 0
}
func genFn1(p0 uint64, p1 int16) uint32 {
    var v7 uint64 = p0
    var v8 int16 = p1
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 534
}

func main() {
    var v1 uint32 = 793
    var v2 uint16 = 267
    v3 := []int{988, 4294967295, 2147483647, 32768}
    var v4 int16 = 0
    var v9 float64 = -2.25
    v2 = v2
    for fi10 := 0; fi10 < 16; fi10++ {
        if false {
            for fi11 := 0; fi11 < 7; fi11++ {
                v12 := GenS1{F0: 1, F1: 146, F2: (v1 ^ (v1 / ((439) | 1)))}
                v2 += (v2 * (v2 << ((255) & 63)))
                _ = fi11
                _ = v12
            }
            v13 := GenS1{F0: 84, F1: 127, F2: (v1 - v1)}
            _ = v13
        } else {
            for fi14 := 0; fi14 < 11; fi14++ {
                fmt.Println(v9, fi10)
                if len(v3) > 0 {
                    v3[0] = fi14
                }
                var v15 int16 = (v4 ^ v4)
                _ = fi14
                _ = v15
            }
            v3 = append(v3, fi10)
            _ = len(v3)
        }
        for fi16 := 0; fi16 < 6; fi16++ {
            v17 := GenS2{F0: 127, F1: (2 > 63)}
            v1 = (^(v1))
            fmt.Println(v1, v9)
            var v18 int64 = 9
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = fi10
    }
    fmt.Println(v2, v2, v1)
    if true {
        if (2 == 0) {
            v4 -= (v4 - v4)
        }
        var v19 float32 = ((-2.25 + 0.1) * ((((float32((float64(-2.25) / v9)) + 0.1) - 0.0) * 0.0) + 0.1))
        _ = v19
    } else {
        v9 = (v9 / 1e10)
        fmt.Println(v2)
    }
    v1 = uint32(v2)
    v20 := GenS3{F0: -652, F1: v1, F2: 444}
    var v21 uint64 = 278
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v9
}
