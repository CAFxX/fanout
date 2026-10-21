package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 uint8; F1 uint; F2 int16 }
type GenS3 struct { F0 uint8; F1 string; F2 int8 }
func genFn0(p0 int8, p1 uint64, p2 uint16) int32 {
    fmt.Println(p2, p0)
    var v4 uint8 = 128
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return int32(p2)
}
func genFn1(p0 int64) int16 {
    v5 := GenS1{F0: 516, F1: 42}
    if (2 == 1) {
        p0 -= (p0 % ((0) | 1))
    } else {
        var v6 int16 = 32767
        _ = v6
    }
    switch -723 {
    case 42:
        p0 = (p0 ^ (p0 - p0))
    case 1:
        if true {
            p0 -= 9223372036854775807
            p0 -= p0
            var v7 int64 = -1
            _ = v7
        } else {
            var v8 int64 = 9223372036854775807
            _ = v8
        }
    default:
        fmt.Println(p0)
    }
    v9 := []int{127}
    _ = p0
    _ = v5
    _ = v9
    return -32768
}

func main() {
    var v1 int = -172
    var v2 int = v1
    v3 := []int{918, (v2 - (v1 & v2))}
    v10 := GenS1{F0: 2, F1: 558}
    var v11 float64 = float64(0.1)
    if (!true) {
        if true {
            v12 := GenS1{F0: 2, F1: 116}
            v13 := GenS2{F0: 255, F1: 42, F2: 42}
            var v14 float32 = (0.1 / (((0.0 - (1e10 / (1.5 * 1e10))) + 0.1) + -2.25))
            _ = v12
            _ = v13
            _ = v14
        }
        if len(v3) > 0 {
            v3[0] = (-(v1))
        }
    } else {
        v11 = 1e10
        v15 := GenS1{F0: 124, F1: int32(v1)}
        _ = v15
    }
    v16 := GenS2{F0: 156, F1: 18446744073709551615, F2: -905}
    fmt.Println(v11)
    v17 := GenS2{F0: 2, F1: 487, F2: genFn1(-9223372036854775808)}
    fmt.Println(v11, v1, v11)
    fmt.Println(v2, v11)
    if true {
        v18 := GenS3{F0: 255, F1: "foo bar", F2: 0}
        _ = v18
    }
    v3 = append(v3, v2)
    var v19 float32 = 1.5
    if (((0 != 2) || true) || true) {
        if (46 != 9223372036854775807) {
            if ((((v19 + ((1e10 / 3.14159) * (v19 * v19))) + ((v19 - (0.0 + v19)) * float32(v11))) - 1.5) <= float32(0.1)) {
                var v20 int8 = 42
                var v21 float64 = 3.14159
                _ = len(v3)
                _ = v20
                _ = v21
            }
            v1 = (v1 & (v1 & (v1 % ((v2) | 1))))
            var v22 int64 = -9223372036854775808
            _ = v22
        }
    }
    if true {
        v2 = v2
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v3
}
