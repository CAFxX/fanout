package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 int8; F1 int32; F2 uint }
type GenS3 struct { F0 uint8; F1 int; F2 uint32 }
type GenS4 struct { F0 uint8; F1 bool }
func genFn0() uint {
    v4 := GenS1{F0: 42, F1: 128}
    v5 := []int{42, 32768, 128, 2147483647}
    if len(v5) > 0 {
        v5[0] = 9223372036854775807
    }
    if false {
        var v6 float64 = float64((1e10 - 0.1))
        _ = len(v5)
        var v7 string = ""
        _ = v6
        _ = v7
    }
    var v8 int8 = 1
    _ = v4
    _ = v5
    _ = v8
    return 237
}
func genFn1(p0 int8) uint8 {
    p0 += p0
    for fi9 := 0; fi9 < 2; fi9++ {
        p0 -= p0
        _ = fi9
    }
    var v10 int64 = 42
    var v11 int16 = 32767
    fmt.Println(p0, v11, p0)
    _ = p0
    _ = v10
    _ = v11
    return 255
}

func main() {
    var v1 uint32 = 2147483648
    var v2 uint = 255
    v3 := make(map[uint32]bool)
    switch 32767 {
    case 32768:
        switch "hello" {
        case "hello":
            v3[(v1 >> (((^(v2))) & 63))] = (!(-1 <= 814))
        default:
            fmt.Println(v2)
        }
    case 0:
        v12 := []int{-328, -9223372036854775808, 190}
        _ = v12
    case 127:
        v13 := GenS2{F0: -62, F1: 1, F2: (v2 >> ((v2) & 63))}
        _ = v13
    default:
        switch "hello" {
        case "":
            for fi14 := 0; fi14 < 6; fi14++ {
                for fi15 := 0; fi15 < 2; fi15++ {
                    if false {
                        var v16 uint32 = v1
                        v2 = 4294967295
                        _ = v16
                    }
                    v17 := []int{fi15, (-(fi14)), int(v2), (fi15 ^ (fi15 ^ fi14))}
                    var v18 uint16 = 0
                    _ = fi15
                    _ = v17
                    _ = v18
                }
                if (false || true) {
                    v19 := GenS3{F0: 2, F1: (fi14 >> ((v2) & 63)), F2: (v1 - (v1 | 159))}
                    v1 += (v1 >> ((genFn0()) & 63))
                    v20 := v3[704]
                    _ = v19
                    _ = v20
                } else {
                    for fi21 := 0; fi21 < 20; fi21++ {
                        switch "" {
                        case "foo bar":
                            v1 += v1
                        }
                        v22 := GenS1{F0: 221, F1: genFn1(42)}
                        if ((1 < -633) && (v2 > v2)) {
                            v2 += (v2 | 18446744073709551615)
                        }
                        _ = fi21
                        _ = v22
                    }
                }
                _ = fi14
            }
        default:
            fmt.Println(v1, v2)
        }
    }
    v2 = 4294967295
    var v23 uint16 = 42
    if true {
        v2 += (v2 ^ v2)
        v2 = 360
    } else {
        var v24 uint16 = 255
        fmt.Println(v1)
        var v25 int8 = 37
        _ = v24
        _ = v25
    }
    v23 = (v23 + 42)
    v26 := GenS4{F0: 0, F1: true}
    var v27 uint32 = (v1 + (v1 * (v1 * 2147483647)))
    fmt.Println(v2)
    switch -1 {
    case -1:
        fmt.Println(v27, v23, v1)
    case -984:
        v27 -= (v1 + v1)
    default:
        if (v2 > genFn0()) {
            v28 := GenS3{F0: genFn1(1), F1: 558, F2: (^(v27))}
            var v29 int = 42
            _ = v28
            _ = v29
        }
    }
    delete(v3, v1)
    for fi30 := 0; fi30 < 8; fi30++ {
        v1 += v1
        switch ("hello" + "a") {
        case "a":
            v31 := GenS3{F0: 0, F1: fi30, F2: 620}
            _ = v31
        default:
            var v32 uint8 = 127
            _ = v32
        }
        fmt.Println(v27, v1)
        _ = fi30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v23
    _ = v26
    _ = v27
    _ = v3
}
