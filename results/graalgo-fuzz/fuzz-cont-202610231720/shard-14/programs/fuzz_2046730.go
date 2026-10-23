package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
func genFn0(p0 uint64, p1 int8, p2 uint) uint {
    switch 1 {
    case -920:
        p0 = 731
    case 1:
        var v4 int32 = -629
        _ = v4
    default:
        var v5 int = 127
        _ = v5
    }
    fmt.Println(p2)
    p0 += (p0 & p0)
    p2 = (p2 * 65535)
    var v6 int = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p2
}
func genFn1(p0 int32, p1 int16, p2 int) uint {
    v7 := GenS1{F0: 255, F1: 255}
    fmt.Println(p0)
    if true {
        p1 += int16(p2)
    }
    p1 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 0
}

func main() {
    v1 := make(map[string]int)
    v2 := []uint32{2147483647, 992}
    var v3 uint8 = 146
    v3 = (^(v3))
    fmt.Println(v3)
    for fi8 := 0; fi8 < 16; fi8++ {
        for fi9 := 0; fi9 < 8; fi9++ {
            var v10 float32 = 0.1
            if (((-636 != -284) && true) && false) {
                if false {
                    var v11 int64 = 42
                    v12 := GenS1{F0: 200, F1: 42}
                    if (!false) {
                        v10 = v10
                        var v13 uint = 2
                        var v14 int16 = -1
                        _ = v13
                        _ = v14
                    }
                    var v15 uint8 = 128
                    _ = v11
                    _ = v12
                    _ = v15
                } else {
                    v10 -= (((0.1 - (0.1 - (v10 / (((v10 - (v10 / (v10 / v10))) + float32(-2.25)) - v10)))) / 3.14159) / 1.5)
                }
            }
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    var v16 uint16 = 2
    switch ("foo bar" + "\n\t") {
    case "foo bar":
        _ = len(v2)
    case "\n\t":
        fmt.Println(v16, v3)
    case "a":
        v17 := GenS1{F0: (v16 + (v16 & v16)), F1: (v16 | v16)}
        _ = v17
    default:
        switch 9223372036854775807 {
        case 691:
            var v18 uint = 438
            _ = v18
        case 32767:
            var v19 uint = 1
            _ = v19
        case 9223372036854775807:
            v3 += (v3 * (v3 - 127))
        default:
            _ = len(v2)
        }
    }
    fmt.Println(v3, v3)
    fmt.Println(v16, v3)
    for fi20 := 0; fi20 < 2; fi20++ {
        v21 := GenS1{F0: (v16 | (v16 + v16)), F1: 284}
        for fi22 := 0; fi22 < 11; fi22++ {
            var v23 int8 = 0
            if len(v2) > 0 {
                v2[0] = 0
            }
            fmt.Println(v16, fi20, fi20)
            v24 := GenS1{F0: v16, F1: (^(v16))}
            var v25 float32 = -2.25
            _ = fi22
            _ = v23
            _ = v24
            _ = v25
        }
        _ = fi20
        _ = v21
    }
    _ = len(v2)
    v16 = (^(v16))
    switch ("foo bar" + "hello") {
    case "a":
        var v26 int8 = -1
        _ = v26
    case "":
        v3 += v3
    }
    var v27 int8 = -128
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v2
    _ = v27
    _ = v3
}
