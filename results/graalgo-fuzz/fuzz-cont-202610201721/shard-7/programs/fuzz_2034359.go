package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 uint; F1 uint64; F2 string }

func main() {
    var v1 float64 = 3.14159
    var v2 int8 = 0
    var v3 int64 = 443
    var v4 float32 = float32(3.14159)
    var v5 uint16 = 0
    v6 := []int{}
    v7 := GenS1{F0: 18446744073709551615}
    for fi8 := 0; fi8 < 1; fi8++ {
        v5 += v5
        _ = fi8
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        for fi10 := 0; fi10 < 4; fi10++ {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            switch (fi9 + 32767) {
            case 0:
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            case 9223372036854775807:
                v4 += (float32(float64((float32(0.0) / float32(3.14159)))) / 1.0)
            case 886:
                var v11 int32 = 0
                _ = v11
            default:
                v5 = 1
            }
            var v12 uint8 = 0
            _ = fi10
            _ = v12
        }
        v4 = float32(v1)
        v1 += 0.1
        _ = fi9
    }
    var v13 uint = 2
    var v14 uint8 = 228
    var v15 int8 = v2
    v6 = append(v6, 9223372036854775807)
    switch "" {
    case "foo bar":
        var v16 uint16 = (^(v5))
        _ = v16
    case "hello":
        v17 := GenS2{F0: (v14 - (v14 & 42)), F1: int(v5)}
        _ = v17
    case "\n\t":
        v4 += -2.25
    default:
        v18 := GenS3{F0: v2, F1: 4294967295}
        _ = v18
    }
    switch -698 {
    case 1:
        _ = len(v6)
    case -1:
        v6 = append(v6, 208)
    default:
        if ((v3 ^ (v3 ^ 9223372036854775807)) < v3) {
            v4 += v4
        } else {
            v14 = 0
            var v19 float64 = 1e10
            _ = v19
        }
    }
    if false {
        if (!false) {
            for fi20 := 0; fi20 < 12; fi20++ {
                v3 = (-(v3))
                v2 = (v2 << (((^(v13))) & 63))
                for fi21 := 0; fi21 < 19; fi21++ {
                    if false {
                        fmt.Println(v13, v4)
                        _ = len(v6)
                        fmt.Println(fi21, v15, v3)
                    } else {
                        v4 -= ((((0.1 / v4) / v4) * v4) - 1e10)
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    }
                    v22 := GenS3{F0: 100, F1: v13}
                    fmt.Println(v4, fi21)
                    _ = fi21
                    _ = v22
                }
                _ = fi20
            }
        }
        var v23 uint = (v13 >> (((v13 | v13)) & 63))
        var v24 int32 = 42
        _ = v23
        _ = v24
    }
    v25 := GenS2{F0: v14, F1: 128}
    for fi26 := 0; fi26 < 2; fi26++ {
        switch (^(fi26)) {
        case 42:
            switch fi26 {
            case 114:
                if (!(449 <= 42)) {
                    fmt.Println(v2, v5, v1)
                } else {
                    for fi27 := 0; fi27 < 16; fi27++ {
                        v6 = append(v6, fi27)
                        v13 = v13
                        v28 := GenS1{F0: 859}
                        v6 = append(v6, 996)
                        v3 = (v3 % (((v3 >> ((v13) & 63))) | 1))
                        _ = fi27
                        _ = v28
                    }
                    v29 := GenS1{F0: 1}
                    _ = v29
                }
            }
        case -839:
            v4 -= (v4 / ((((0.1 - v4) * (-2.25 + v4)) * float32(float64((float32(((v1 * (((v1 / 1.0) / 1e10) + -2.25)) / (v1 - (float64(v4) * v1)))) - v4)))) + (v4 / ((v4 / -2.25) + v4))))
        default:
            fmt.Println(v13, v2, v3)
        }
        _ = fi26
    }
    v30 := GenS4{F0: (v13 % ((v13) | 1)), F1: 0, F2: "a"}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v25
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
