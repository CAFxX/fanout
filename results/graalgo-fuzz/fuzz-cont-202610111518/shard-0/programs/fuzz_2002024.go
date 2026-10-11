package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 int8; F1 int; F2 uint16 }
type GenS3 struct { F0 uint8; F1 int8; F2 int32 }
type GenS4 struct { F0 string; F1 int16; F2 uint }
type GenS5 struct { F0 int8; F1 uint; F2 uint32 }
func genFn0(p0 uint64) uint16 {
    v4 := GenS1{F0: 18446744073709551615, F1: 65535}
    v5 := GenS2{F0: 1, F1: -2147483648, F2: 0}
    _ = p0
    _ = v4
    _ = v5
    return 2
}
func genFn1(p0 int, p1 uint64) int32 {
    switch p0 {
    case -352:
        v6 := GenS3{F0: 9, F1: 13, F2: 1}
        _ = v6
    default:
        fmt.Println(p0, p0)
    }
    v7 := GenS4{F0: "", F1: -32768, F2: 597}
    _ = p0
    _ = p1
    _ = v7
    return -2147483648
}

func main() {
    var v1 int16 = 766
    var v2 uint = 1
    var v3 uint16 = 32767
    switch "hello" {
    case "hello":
        var v8 uint32 = 4294967295
        _ = v8
    default:
        v3 -= 239
    }
    for fi9 := 0; fi9 < 1; fi9++ {
        fmt.Println(v1)
        v3 += v3
        _ = fi9
    }
    v10 := []int{-2147483648}
    var v11 float32 = 0.0
    v11 -= v11
    var v12 uint = (v2 >> (((^(v2))) & 63))
    v2 -= 0
    for fi13 := 0; fi13 < 3; fi13++ {
        for fi14 := 0; fi14 < 11; fi14++ {
            switch 42 {
            case 4294967295:
                if len(v10) > 0 {
                    v10[0] = (fi13 / ((127) | 1))
                }
            case 42:
                var v15 float32 = 0.1
                _ = v15
            case -9223372036854775808:
                if (!(v1 != v1)) {
                    var v16 uint32 = 42
                    var v17 int32 = int32(fi13)
                    _ = v16
                    _ = v17
                } else {
                    fmt.Println(v11, v3, v12)
                    fmt.Println(fi13)
                    if true {
                        v11 = ((v11 / v11) * 1.5)
                        var v18 int = fi13
                        _ = v18
                    } else {
                        var v19 bool = (2 < 350)
                        _ = len(v10)
                        _ = v19
                    }
                    v12 = (v2 - v12)
                }
            default:
                var v20 float64 = 3.14159
                _ = v20
            }
            v10 = append(v10, fi13)
            v1 = v1
            _ = fi14
        }
        if true {
            fmt.Println(v1, v1, v1)
        }
        _ = fi13
    }
    v21 := GenS5{F0: -18, F1: v2, F2: 963}
    v1 = (v1 / (((v1 & (v1 * (v1 & (v1 - v1))))) | 1))
    for fi22 := 0; fi22 < 4; fi22++ {
        switch -108 {
        case 738:
            v12 = (v12 << (((v2 >> ((v2) & 63))) & 63))
        case -197:
            for fi23 := 0; fi23 < 8; fi23++ {
                v12 += (v12 - 4294967295)
                _ = fi23
            }
        case 0:
            var v24 float32 = (((v11 * v11) + (v11 - (float32(0.1) * (float32((((0.1 - 3.14159) / 3.14159) * 1.5)) - 1.5)))) / v11)
            _ = v24
        default:
            v12 = 267
        }
        if len(v10) > 0 {
            v10[0] = fi22
        }
        _ = fi22
    }
    if len(v10) > 0 {
        v10[0] = 2147483647
    }
    v2 -= v12
    var v25 uint8 = uint8(v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v21
    _ = v25
    _ = v3
}
