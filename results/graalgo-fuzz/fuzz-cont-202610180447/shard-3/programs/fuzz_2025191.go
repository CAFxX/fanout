package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 int16 }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 int; F1 int64; F2 uint64 }
type GenS5 struct { F0 int32; F1 uint32 }
type GenS6 struct { F0 string; F1 uint; F2 uint }
func genFn0(p0 string, p1 string, p2 int) uint16 {
    if (1 <= 0) {
        var v6 bool = (p2 <= p2)
        var v7 int16 = 0
        for fi8 := 0; fi8 < 13; fi8++ {
            fmt.Println(p0, p0)
            if (2 == 826) {
                v9 := []int{(p2 - -589), (p2 & fi8), fi8, (p2 / ((-9223372036854775808) | 1))}
                if (false || (v6 || (fi8 != fi8))) {
                    var v10 float64 = (0.1 * 0.0)
                    var v11 int8 = 100
                    _ = v10
                    _ = v11
                }
                _ = len(v9)
                _ = v9
            } else {
                v6 = (1 < 18446744073709551615)
            }
            v12 := GenS1{F0: 918, F1: (v7 | v7), F2: (-(v7))}
            _ = fi8
            _ = v12
        }
        _ = v6
        _ = v7
    } else {
        var v13 int8 = 100
        p1 = p1
        _ = v13
    }
    p1 = p0
    switch p2 {
    case -9223372036854775808:
        var v14 int64 = 1
        _ = v14
    case -827:
        switch p0 {
        case "\n\t":
            var v15 int64 = -9223372036854775808
            _ = v15
        case "a":
            switch (p2 | p2) {
            case 42:
                var v16 float64 = 0.0
                _ = v16
            case -49:
                var v17 uint = 42
                _ = v17
            default:
                var v18 string = "\n\t"
                _ = v18
            }
        case "foo bar":
            p0 = (p0 + p0)
        }
    default:
        v19 := GenS1{F0: 469, F1: 42, F2: -32768}
        _ = v19
    }
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    return 32767
}

func main() {
    v1 := []uint32{100, 2147483647, 2147483647, 278}
    v2 := make(map[uint32]string)
    var v3 int32 = 2147483647
    var v4 string = "a"
    var v5 uint = 4294967295
    if true {
        v20 := GenS1{F0: 42, F1: 1, F2: -1}
        v4 = v4
        _ = v20
    }
    v21 := GenS2{F0: 100, F1: true}
    v5 = 69
    var v22 float32 = ((0.0 * ((-2.25 / -2.25) * -2.25)) - 3.14159)
    v23 := GenS3{F0: v4, F1: -70}
    v24 := GenS1{F0: 887, F1: -277, F2: 844}
    if true {
        v22 -= (v22 - 0.0)
        v25 := GenS4{F0: -125, F1: 594, F2: 0}
        _ = v25
    }
    v1 = append(v1, 4294967295)
    if true {
        for fi26 := 0; fi26 < 13; fi26++ {
            v5 += (v5 - (v5 % (((^(v5))) | 1)))
            switch (-(fi26)) {
            case 32767:
                v4 = (v4 + v4)
            case 127:
                v27 := GenS2{F0: 0, F1: true}
                _ = v27
            default:
                v3 += 42
            }
            _ = fi26
        }
        var v28 int32 = v3
        for fi29 := 0; fi29 < 3; fi29++ {
            if (false || (-128 < 1)) {
                v28 = (v3 % ((-2147483648) | 1))
            } else {
                switch fi29 {
                case 2147483647:
                    fmt.Println(v4)
                case 4294967295:
                    var v30 int16 = 1
                    _ = v30
                case 32768:
                    var v31 int64 = -697
                    _ = v31
                default:
                    v22 += (((v22 - float32(((-2.25 - 1.5) * 1.5))) - float32((float64((v22 * (float32(float64(float32(1.5))) / 1.0))) - -2.25))) * float32((((0.1 / float64(((((v22 * v22) + ((1.5 + v22) * v22)) - (0.1 - ((((1e10 / 0.1) / (0.1 + (v22 * float32(1e10)))) / ((v22 / (v22 + float32(1e10))) - ((float32((((((-2.25 - (1.5 / 1.5)) - ((float64(1e10) - 0.1) + 1e10)) - ((1.5 * 1e10) + 3.14159)) * float64(3.14159)) * float64((v22 + v22)))) * v22) * (float32(1e10) * v22)))) / v22))) * v22))) * (0.1 + 0.1)) / 1.0)))
                }
                v32 := GenS5{F0: -927, F1: 2}
                _ = v32
            }
            v33 := GenS6{F0: "a", F1: (v5 / ((v5) | 1)), F2: v5}
            _ = fi29
            _ = v33
        }
        _ = v28
    }
    for fi34 := 0; fi34 < 8; fi34++ {
        var v35 float32 = 0.0
        _ = fi34
        _ = v35
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
