package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 uint64 }
type GenS2 struct { F0 bool; F1 uint32; F2 bool }
func genFn0(p0 uint, p1 string) uint64 {
    v4 := []int{}
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = v4
    return 1
}
func genFn1(p0 uint16) int16 {
    var v5 int = 826
    fmt.Println(v5, v5)
    var v6 int8 = 42
    v6 = 127
    v6 -= v6
    _ = p0
    _ = v5
    _ = v6
    return -419
}

func main() {
    var v1 float64 = -2.25
    var v2 float32 = (float32(-2.25) / (1e10 - 0.1))
    var v3 int8 = 1
    var v7 int32 = 2147483647
    v3 += (v3 / ((100) | 1))
    v3 += (-(v3))
    var v8 int16 = 1
    if true {
        var v9 int = -355
        v10 := []int{v9, (v9 - v9), (v9 ^ (-(v9))), v9}
        switch v9 {
        case 858:
            for fi11 := 0; fi11 < 11; fi11++ {
                var v12 float32 = 0.0
                v13 := GenS1{F0: 18446744073709551615, F1: (v8 * v8), F2: 4294967295}
                _ = fi11
                _ = v12
                _ = v13
            }
        case -1:
            var v14 int64 = 9223372036854775807
            _ = v14
        default:
            if (float32(-2.25) >= -2.25) {
                _ = len(v10)
                switch -464 {
                case 127:
                    var v15 int = v9
                    _ = v15
                default:
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                }
                if false {
                    fmt.Println(v2)
                    if false {
                        v2 = float32(v1)
                    }
                    v16 := GenS2{F0: ((2 >= 18446744073709551615) || true), F1: 823, F2: (false && (true && (892 < 766)))}
                    _ = v16
                }
            } else {
                var v17 uint = 237
                fmt.Println(v3, v8)
                var v18 float32 = -2.25
                _ = v17
                _ = v18
            }
        }
        _ = v10
        _ = v9
    }
    if (false && (false && false)) {
        v3 = -128
    } else {
        var v19 int16 = (^(v8))
        _ = v19
    }
    v3 -= (v3 << ((2) & 63))
    v20 := []int{}
    var v21 uint16 = 84
    _ = len(v20)
    var v22 int64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v7
    _ = v8
}
