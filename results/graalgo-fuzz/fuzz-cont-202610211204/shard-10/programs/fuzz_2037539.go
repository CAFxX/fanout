package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int }
func genFn0(p0 int) uint8 {
    p0 -= p0
    v6 := GenS1{F0: 1, F1: -1}
    _ = p0
    _ = v6
    return 0
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint = 443
    var v3 int32 = -773
    var v4 uint16 = 42
    var v5 uint16 = (v4 | (v4 * 365))
    for fi7 := 0; fi7 < 18; fi7++ {
        v4 = 32767
        switch (-(fi7)) {
        case 0:
            if true {
                for fi8 := 0; fi8 < 19; fi8++ {
                    v2 = v2
                    v2 += (v2 ^ (^(v2)))
                    _ = fi8
                }
                v9 := GenS2{F0: 1, F1: (fi7 << (((v2 + (v2 >> ((18446744073709551615) & 63)))) & 63))}
                _ = v9
            } else {
                if ((!((v4 * (v4 + 32767)) >= v5)) && (2 > 60)) {
                    var v10 int8 = -19
                    var v11 uint32 = 1
                    var v12 bool = (1 > (fi7 & 42))
                    _ = v10
                    _ = v11
                    _ = v12
                }
                switch "hello" {
                case "a":
                    var v13 float64 = 3.14159
                    _ = v13
                case "hello":
                    v14 := []int{-9223372036854775808, 538, 127}
                    _ = v14
                default:
                    for fi15 := 0; fi15 < 13; fi15++ {
                        var v16 float64 = float64(0.0)
                        _ = fi15
                        _ = v16
                    }
                }
            }
        case 32767:
            switch 1 {
            case -9223372036854775808:
                v17 := GenS3{F0: 164}
                _ = v17
            case 1:
                v2 = (v2 | v2)
            default:
                v4 = v4
            }
        case 127:
            var v18 uint64 = 1
            _ = v18
        }
        _ = fi7
    }
    var v19 uint8 = genFn0(-1)
    for fi20 := 0; fi20 < 18; fi20++ {
        var v21 int8 = 42
        _ = fi20
        _ = v21
    }
    if (true && (uint32(v19) == 65535)) {
        v22 := []int{1, 938, 32768, -713}
        _ = v22
    }
    v23 := []int{-1, 1, 32768, -131}
    var v24 int16 = -378
    v25 := GenS4{F0: (v24 - (v24 % ((v24) | 1)))}
    for fi26 := 0; fi26 < 17; fi26++ {
        v27 := GenS1{F0: v5, F1: -337}
        delete(v1, fi26)
        if len(v23) > 0 {
            v23[0] = (fi26 & fi26)
        }
        _ = fi26
        _ = v27
    }
    v28 := GenS5{F0: 42}
    v19 = v19
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
