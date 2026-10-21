package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 uint64; F1 string; F2 int }
type GenS3 struct { F0 int32; F1 bool }
func genFn0(p0 uint8) int32 {
    if false {
        v4 := []int{-27}
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        switch 4294967295 {
        case 0:
            fmt.Println(p0)
        case 2147483647:
            var v5 uint32 = 903
            _ = v5
        default:
            if false {
                v6 := GenS1{F0: 127, F1: 65535}
                var v7 string = "\n\t"
                _ = v6
                _ = v7
            }
        }
        _ = v4
    } else {
        for fi8 := 0; fi8 < 19; fi8++ {
            switch fi8 {
            case -654:
                fmt.Println(p0)
            case -592:
                p0 = p0
            case -1:
                p0 -= p0
            default:
                switch (fi8 * fi8) {
                case 517:
                    for fi9 := 0; fi9 < 17; fi9++ {
                        v10 := []int{(fi8 ^ fi8), 4294967295}
                        if false {
                            _ = len(v10)
                            p0 = (p0 % ((p0) | 1))
                            switch fi9 {
                            case 32768:
                                switch fi8 {
                                case 127:
                                    var v11 int8 = -128
                                    _ = v11
                                }
                            default:
                                var v12 uint64 = 514
                                _ = v12
                            }
                            var v13 int32 = 1
                            fmt.Println(v13, p0)
                            _ = v13
                        } else {
                            var v14 int64 = 42
                            _ = v14
                        }
                        if (4294967295 <= 2) {
                            v15 := GenS1{F0: -128, F1: 152}
                            _ = v15
                        }
                        _ = fi9
                        _ = v10
                    }
                case -2147483648:
                    var v16 string = "foo bar"
                    _ = v16
                default:
                    v17 := GenS1{F0: 33, F1: 0}
                    _ = v17
                }
            }
            v18 := GenS1{F0: 0, F1: 32767}
            _ = fi8
            _ = v18
        }
        p0 = p0
    }
    fmt.Println(p0)
    _ = p0
    return -633
}

func main() {
    v1 := []string{}
    v2 := []int{-965, -9223372036854775808}
    var v3 int32 = 2147483647
    if true {
        fmt.Println(v3)
    }
    v3 -= (-(v3))
    v3 = genFn0(255)
    var v19 int64 = -519
    if true {
        var v20 int32 = 454
        var v21 uint8 = 127
        _ = v20
        _ = v21
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v3 = genFn0(127)
    switch 127 {
    case 1:
        v22 := GenS1{F0: 127, F1: 837}
        _ = v22
    case 0:
        var v23 uint32 = 42
        _ = v23
    case 32767:
        v24 := GenS2{F0: 18446744073709551615, F1: ("foo bar" + ((("a" + "foo bar") + "foo bar") + "a")), F2: int(v19)}
        _ = v24
    default:
        var v25 uint8 = 0
        _ = v25
    }
    fmt.Println(v3, v19)
    switch 583 {
    case 9223372036854775807:
        v3 = v3
    }
    var v26 int32 = genFn0(82)
    var v27 uint8 = 128
    v28 := GenS3{F0: v3, F1: ((-2.25 - ((((3.14159 / ((float32(0.0) - float32(((1e10 * (0.0 * 1e10)) + 0.1))) + float32(1.5))) + -2.25) + (0.0 / 1e10)) * (0.0 * 1e10))) != (0.0 - 3.14159))}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if len(v1) > 0 {
        v1[0] = "a"
    }
    var v29 uint64 = 42
    v29 += 866
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
}
