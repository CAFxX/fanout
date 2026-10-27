package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 uint; F1 int64; F2 bool }
type GenS3 struct { F0 bool }
func genFn0(p0 bool, p1 bool, p2 int32) uint64 {
    var v5 int64 = 289
    fmt.Println(p1, p0, p1)
    v6 := GenS1{F0: true, F1: 655}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 965
}
func genFn1(p0 int64) int32 {
    var v7 uint64 = 747
    switch 128 {
    case -1:
        for fi8 := 0; fi8 < 13; fi8++ {
            var v9 float32 = ((1e10 / 3.14159) * ((-2.25 - -2.25) - float32(1.5)))
            for fi10 := 0; fi10 < 12; fi10++ {
                v11 := GenS2{F0: 2, F1: -9223372036854775808, F2: false}
                v12 := []int{}
                p0 += 253
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = fi8
            _ = v9
        }
    default:
        if (true || (true || false)) {
            v13 := GenS2{F0: 42, F1: (^(p0)), F2: ((p0 & p0) < p0)}
            _ = v13
        } else {
            var v14 int64 = 0
            _ = v14
        }
    }
    var v15 int8 = 1
    _ = p0
    _ = v15
    _ = v7
    return 2147483647
}

func main() {
    var v1 uint8 = 210
    var v2 float64 = 0.1
    var v3 int64 = 22
    var v4 uint = 4294967295
    if true {
        v16 := GenS1{F0: false, F1: v4}
        var v17 int = 945
        var v18 int64 = v3
        _ = v16
        _ = v17
        _ = v18
    } else {
        v4 += (v4 + (v4 >> (((v4 / ((v4) | 1))) & 63)))
    }
    v19 := []int{127, -9223372036854775808, 2147483647}
    fmt.Println(v1)
    v4 = (v4 | v4)
    v20 := GenS3{F0: true}
    var v21 float64 = (v2 * (1e10 / v2))
    if (v21 > 0.0) {
        v19 = append(v19, 127)
        var v22 bool = (4294967295 != 1)
        _ = v22
    }
    v23 := GenS1{F0: false, F1: 2}
    fmt.Println(v4, v21, v4)
    var v24 int8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
