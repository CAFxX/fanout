package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 int64; F1 uint64; F2 uint32 }
type GenS3 struct { F0 int64; F1 int64 }
func genFn0(p0 uint16, p1 bool, p2 uint64) string {
    p2 += (p2 | 4294967295)
    for fi6 := 0; fi6 < 5; fi6++ {
        v7 := GenS1{F0: 313, F1: 32}
        _ = fi6
        _ = v7
    }
    v8 := GenS2{F0: 1, F1: p2, F2: uint32(p2)}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return "hello"
}
func genFn1(p0 uint32) uint {
    p0 -= p0
    v9 := []int{}
    if len(v9) > 0 {
        v9[0] = 32768
    }
    v10 := GenS2{F0: 0, F1: 0, F2: (p0 - p0)}
    p0 -= (^(p0))
    _ = p0
    _ = v10
    _ = v9
    return 65535
}

func main() {
    v1 := make(map[string]int)
    var v2 int8 = 7
    v3 := []uint32{388, 135, 2}
    var v4 uint16 = 0
    var v5 int16 = 316
    for fi11 := 0; fi11 < 16; fi11++ {
        var v12 uint16 = (v4 + (v4 << ((18446744073709551615) & 63)))
        _ = fi11
        _ = v12
    }
    var v13 float64 = (1.5 * ((0.0 - 3.14159) / -2.25))
    var v14 bool = (0 > 2)
    switch 4294967295 {
    case 466:
        if len(v3) > 0 {
            v3[0] = 650
        }
    case -639:
        var v15 int32 = -1
        _ = v15
    case 9223372036854775807:
        for fi16 := 0; fi16 < 13; fi16++ {
            var v17 int16 = (v5 % (((-(v5))) | 1))
            var v18 int = 32768
            v2 = (v2 ^ v2)
            _ = fi16
            _ = v17
            _ = v18
        }
    default:
        v19 := GenS1{F0: 4294967295, F1: 1}
        _ = v19
    }
    _ = len(v3)
    for fi20 := 0; fi20 < 16; fi20++ {
        var v21 int64 = 1
        var v22 uint32 = 2147483647
        _ = fi20
        _ = v21
        _ = v22
    }
    var v23 int64 = -658
    v24 := GenS2{F0: -1, F1: 4294967295, F2: 4294967295}
    var v25 uint64 = 1
    var v26 int = int(v25)
    v2 = v2
    fmt.Println(v4)
    v5 -= (v5 << ((4294967295) & 63))
    switch "\n\t" {
    case "foo bar":
        v5 = (v5 * v5)
    default:
        fmt.Println(v4, v14)
    }
    v3 = append(v3, 2147483648)
    fmt.Println(v2, v23, v2)
    v5 = (v5 & (v5 ^ -1))
    if v14 {
        if (false || (v14 || ((v14 && false) || (true || false)))) {
            v27 := GenS3{F0: v23, F1: (v23 - v23)}
            var v28 uint16 = 1
            _ = v27
            _ = v28
        }
        if ((v23 % ((v23) | 1)) == v23) {
            fmt.Println(v26, v26)
            fmt.Println(v23)
            for fi29 := 0; fi29 < 11; fi29++ {
                if (91 < 65) {
                    var v30 uint32 = 42
                    _ = v30
                }
                if v14 {
                    fmt.Println(fi29, v14, v2)
                } else {
                    delete(v1, "\n\t")
                    if v14 {
                        v23 = (v23 - -9223372036854775808)
                    }
                }
                _ = fi29
            }
        }
    } else {
        var v31 float64 = v13
        if (!v14) {
            v4 = v4
            var v32 uint = 1
            _ = v32
        }
        _ = v31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
