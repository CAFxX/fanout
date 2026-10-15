package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 string }
type GenS2 struct { F0 uint; F1 int; F2 bool }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 uint8; F1 uint16; F2 uint16 }

func main() {
    var v1 int = 9223372036854775807
    var v2 float64 = 0.0
    v3 := []int{1, 127, (v1 % ((v1) | 1)), v1}
    var v4 uint = 255
    var v5 float64 = v2
    var v6 uint8 = 105
    v1 = 127
    var v7 int16 = 42
    for fi8 := 0; fi8 < 17; fi8++ {
        _ = len(v3)
        var v9 uint8 = (v6 - 1)
        for fi10 := 0; fi10 < 13; fi10++ {
            v11 := GenS1{F0: 65535, F1: (^(v4)), F2: ("" + ("\n\t" + "hello"))}
            fmt.Println(v1)
            _ = fi10
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    _ = len(v3)
    switch v1 {
    case 854:
        var v12 float64 = v2
        _ = v12
    case 9223372036854775807:
        for fi13 := 0; fi13 < 14; fi13++ {
            v5 = ((v2 * v2) / (v5 + (((v2 / (v2 * v2)) / (v5 / v2)) - v5)))
            v14 := GenS1{F0: 42, F1: v4, F2: "foo bar"}
            _ = fi13
            _ = v14
        }
    default:
        v15 := GenS2{F0: v4, F1: (^(v1)), F2: (0 != 42)}
        _ = v15
    }
    v7 -= (v7 - (v7 * int16(v1)))
    var v16 uint = v4
    v1 = (v1 | v1)
    fmt.Println(v2, v4, v1)
    v17 := GenS1{F0: 18446744073709551615, F1: v4, F2: "\n\t"}
    var v18 uint32 = 2147483647
    var v19 int32 = -193
    for fi20 := 0; fi20 < 9; fi20++ {
        var v21 int64 = -1
        if false {
            switch "foo bar" {
            case "\n\t":
                if (378 < 279) {
                    for fi22 := 0; fi22 < 1; fi22++ {
                        if (float64(0.1) <= ((0.0 * (v2 + -2.25)) - (0.0 * v5))) {
                            v23 := GenS2{F0: v16, F1: (v1 << (((^(v4))) & 63)), F2: ((-(v19)) == (v19 / ((v19) | 1)))}
                            _ = v23
                        } else {
                            v24 := GenS3{F0: 0, F1: -88}
                            v7 -= v7
                            _ = v24
                        }
                        _ = fi22
                    }
                    v21 = (-(v21))
                    for fi25 := 0; fi25 < 10; fi25++ {
                        v26 := GenS4{F0: v6, F1: 586, F2: 42}
                        v19 -= -38
                        _ = fi25
                        _ = v26
                    }
                }
            case "":
                for fi27 := 0; fi27 < 17; fi27++ {
                    v21 += (v21 & v21)
                    if (!true) {
                        fmt.Println(v16, v1)
                        var v28 int = v1
                        _ = v28
                    }
                    fmt.Println(v7, v19)
                    var v29 float32 = 0.1
                    _ = fi27
                    _ = v29
                }
            case "a":
                fmt.Println(v5, v21)
            default:
                for fi30 := 0; fi30 < 3; fi30++ {
                    v7 = v7
                    var v31 uint = v4
                    v32 := GenS4{F0: v6, F1: 435, F2: 86}
                    _ = fi30
                    _ = v31
                    _ = v32
                }
            }
        }
        fmt.Println(v7)
        _ = fi20
        _ = v21
    }
    v16 -= v16
    v2 += v2
    v18 -= 42
    var v33 float32 = (1e10 / (float32(((v2 - v2) / v2)) - (0.0 - 1.5)))
    v3 = append(v3, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
