package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64; F1 uint8; F2 uint16 }
func genFn0() int8 {
    var v5 int8 = 1
    v5 += (v5 | v5)
    v5 -= (^(v5))
    v5 = 127
    _ = v5
    return 100
}
func genFn1(p0 uint64, p1 uint16, p2 string) string {
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 bool = (42 == 65535)
        v8 := GenS1{F0: 0, F1: 236}
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := []int{127, 1, -147, 42}
    p1 -= (p1 >> ((18446744073709551615) & 63))
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return p2
}

func main() {
    var v1 uint64 = 508
    v2 := make(map[uint32]bool)
    var v3 uint = 255
    var v4 int32 = 81
    var v10 float64 = ((-2.25 / 3.14159) / float64(1.5))
    fmt.Println(v3)
    var v11 uint64 = 2
    for fi12 := 0; fi12 < 14; fi12++ {
        v4 = (v4 & v4)
        for fi13 := 0; fi13 < 8; fi13++ {
            v11 -= uint64(v3)
            fmt.Println(v11)
            v14 := v2[944]
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    var v15 int8 = -1
    v11 = (v1 + v11)
    switch "foo bar" {
    case "foo bar":
        fmt.Println(v1, v3)
    default:
        for fi16 := 0; fi16 < 14; fi16++ {
            v10 += 0.0
            v1 = v11
            v11 -= v1
            _ = fi16
        }
    }
    for fi17 := 0; fi17 < 16; fi17++ {
        v1 = 1
        var v18 int16 = -99
        _ = fi17
        _ = v18
    }
    var v19 int8 = v15
    v3 = v3
    v20 := []int{-1, 127, -636, 9223372036854775807}
    v11 += v11
    v4 = v4
    v21 := GenS1{F0: 0, F1: 189}
    v22 := GenS1{F0: 227, F1: 127}
    switch 32768 {
    case 1:
        if (289 < 1) {
            _ = len(v20)
            v19 -= (v19 * (v15 * v19))
            if len(v20) > 0 {
                v20[0] = -9223372036854775808
            }
        }
    case 46:
        if false {
            _ = len(v20)
            if (!false) {
                switch -9223372036854775808 {
                case 32768:
                    var v23 uint32 = 406
                    _ = v23
                case 1:
                    _ = len(v2)
                case -627:
                    v24 := GenS2{F0: uint(v11)}
                    _ = v24
                }
                v10 = v10
            } else {
                v25 := GenS3{F0: -1, F1: 42, F2: 592}
                var v26 bool = false
                _ = v25
                _ = v26
            }
            for fi27 := 0; fi27 < 8; fi27++ {
                var v28 int64 = -713
                delete(v2, 2)
                v29 := GenS3{F0: (v28 ^ (v28 + (v28 + (v28 / ((-9223372036854775808) | 1))))), F1: 128, F2: 65535}
                _ = fi27
                _ = v28
                _ = v29
            }
        }
    case -9223372036854775808:
        var v30 uint8 = 2
        _ = v30
    default:
        v1 = (v1 * (^(v11)))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
