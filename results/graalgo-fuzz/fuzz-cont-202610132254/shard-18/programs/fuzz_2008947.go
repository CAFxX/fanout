package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 int8 }
type GenS2 struct { F0 int8; F1 int16; F2 uint32 }
type GenS3 struct { F0 string; F1 uint32 }
func genFn0(p0 int8) bool {
    p0 -= 38
    p0 += (p0 & p0)
    fmt.Println(p0)
    var v7 float32 = ((-2.25 - (0.0 + (1.5 * 1e10))) - -2.25)
    _ = p0
    _ = v7
    return true
}
func genFn1(p0 int8, p1 bool, p2 string) uint32 {
    for fi8 := 0; fi8 < 15; fi8++ {
        for fi9 := 0; fi9 < 4; fi9++ {
            p2 = p2
            _ = fi9
        }
        v10 := GenS1{F0: -9223372036854775808, F1: "hello", F2: (p0 & -128)}
        v11 := GenS1{F0: 1, F1: p2, F2: (p0 ^ p0)}
        _ = fi8
        _ = v10
        _ = v11
    }
    fmt.Println(p2, p0, p2)
    p0 = (p0 | -1)
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 uint8 = 112
    var v2 int32 = 917
    v3 := make(map[uint32]bool)
    var v4 int = 32768
    v5 := []int{(^(v4)), (v4 % ((1) | 1)), v4}
    var v6 uint32 = 145
    v5 = append(v5, 1)
    fmt.Println(v4, v1, v4)
    for fi12 := 0; fi12 < 18; fi12++ {
        var v13 uint = 42
        _ = fi12
        _ = v13
    }
    var v14 bool = (9223372036854775807 >= -9223372036854775808)
    switch "\n\t" {
    case "hello":
        v4 -= 32768
    case "a":
        for fi15 := 0; fi15 < 11; fi15++ {
            if (2 < 18446744073709551615) {
                delete(v3, (^(v6)))
            } else {
                v3[1] = (1 != 1)
                fmt.Println(v6, v1, v14)
                v16 := GenS2{F0: 105, F1: 1, F2: (v6 & v6)}
                _ = v16
            }
            v17 := v3[2]
            _ = fi15
            _ = v17
        }
    case "":
        var v18 int64 = 0
        _ = v18
    default:
        var v19 float32 = ((1e10 * 1.5) + 0.1)
        _ = v19
    }
    for fi20 := 0; fi20 < 9; fi20++ {
        var v21 uint64 = 756
        switch "foo bar" {
        case "foo bar":
            v22 := GenS3{F0: (("hello" + "hello") + "a"), F1: 42}
            _ = v22
        case "":
            var v23 float32 = (0.0 / 1.5)
            _ = v23
        case "hello":
            switch 0 {
            case -516:
                var v24 uint64 = v21
                _ = v24
            default:
                if len(v5) > 0 {
                    v5[0] = (v4 >> ((255) & 63))
                }
            }
        default:
            var v25 uint8 = v1
            _ = v25
        }
        v6 = v6
        _ = fi20
        _ = v21
    }
    var v26 bool = ((v4 / (((v4 ^ v4)) | 1)) >= (v4 & v4))
    for fi27 := 0; fi27 < 11; fi27++ {
        v28 := GenS3{F0: "\n\t", F1: (v6 << ((42) & 63))}
        _ = fi27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
