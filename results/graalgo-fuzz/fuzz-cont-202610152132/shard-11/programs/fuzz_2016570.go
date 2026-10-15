package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 uint32; F1 uint32; F2 uint }
func genFn0(p0 uint) int8 {
    p0 -= p0
    p0 -= p0
    _ = p0
    return 0
}
func genFn1(p0 uint8, p1 uint32) int {
    p0 -= p0
    var v4 int32 = -1
    p1 = 0
    v5 := GenS1{F0: (p1 << ((900) & 63)), F1: ""}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 32767
}

func main() {
    var v1 int32 = 0
    v2 := make(map[int]string)
    v3 := make(map[uint32]string)
    if (42 <= 121) {
        fmt.Println(v1)
    } else {
        for fi6 := 0; fi6 < 20; fi6++ {
            var v7 float64 = (0.1 * (1.5 * 1.5))
            _ = fi6
            _ = v7
        }
        var v8 uint32 = 65535
        var v9 int8 = -128
        _ = v8
        _ = v9
    }
    v1 += 42
    v1 = v1
    var v10 float32 = float32(3.14159)
    var v11 int64 = 42
    for fi12 := 0; fi12 < 16; fi12++ {
        var v13 uint16 = 42
        fmt.Println(fi12, v10, v13)
        v3[42] = "foo bar"
        _ = fi12
        _ = v13
    }
    v1 -= -902
    switch genFn1(24, 267) {
    case 32767:
        delete(v2, 2147483647)
    case -1:
        var v14 string = "\n\t"
        _ = v14
    case -9223372036854775808:
        v15 := []int{4294967295, -2147483648, 145, -1}
        _ = v15
    default:
        v11 = -442
    }
    for fi16 := 0; fi16 < 13; fi16++ {
        v17 := GenS2{F0: -9223372036854775808, F1: v1}
        _ = fi16
        _ = v17
    }
    fmt.Println(v1, v10, v1)
    switch 0 {
    case 42:
        v18 := GenS2{F0: (v11 + v11), F1: int32(v11)}
        _ = v18
    case 0:
        v19 := GenS3{F0: 4294967295, F1: 2, F2: 292}
        _ = v19
    case 9223372036854775807:
        var v20 string = "foo bar"
        _ = v20
    default:
        if false {
            v11 -= (v11 / ((v11) | 1))
            if (42 < 0) {
                var v21 int64 = 1
                v22 := []int{genFn1(0, 65535), 32768, 1}
                if false {
                    v1 += (^(v1))
                    v11 = v11
                    for fi23 := 0; fi23 < 12; fi23++ {
                        v24 := GenS2{F0: (-(v11)), F1: -2147483648}
                        _ = fi23
                        _ = v24
                    }
                }
                _ = v21
                _ = v22
            }
        } else {
            var v25 uint16 = 2
            var v26 float32 = v10
            _ = v25
            _ = v26
        }
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
}
