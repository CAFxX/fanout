package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 int16 }
type GenS2 struct { F0 int8; F1 uint64; F2 int }
type GenS3 struct { F0 uint64; F1 uint16; F2 uint }
type GenS4 struct { F0 int16; F1 int64; F2 int8 }
type GenS5 struct { F0 uint8; F1 int; F2 int16 }

func main() {
    v1 := []string{("" + "foo bar")}
    var v2 int32 = 0
    v3 := make(map[uint32]bool)
    if false {
        if false {
            var v4 int64 = 1
            _ = len(v1)
            v2 -= 42
            _ = v4
        }
    }
    var v5 float32 = float32(((0.0 / 1.5) / ((1.5 / 1.0) / ((1e10 - (1.5 / ((0.1 - ((-2.25 - 0.0) + (-2.25 - (float64((1e10 + 1e10)) / 1e10)))) - (0.1 + 0.0)))) + (1.5 * 0.0)))))
    v2 += v2
    fmt.Println(v2)
    v2 -= v2
    fmt.Println(v2, v2)
    v6 := GenS1{F0: int64(v2), F1: 1, F2: -1}
    var v7 int64 = 42
    switch ("" + "") {
    case "\n\t":
        if true {
            var v8 int8 = -36
            v2 = v2
            for fi9 := 0; fi9 < 1; fi9++ {
                v1 = append(v1, "hello")
                if (!true) {
                    for fi10 := 0; fi10 < 15; fi10++ {
                        v5 = v5
                        _ = fi10
                    }
                    if (864 < 530) {
                        var v11 uint8 = 255
                        v12 := GenS2{F0: (v8 + v8), F1: 630, F2: 1}
                        _ = v11
                        _ = v12
                    }
                    v13 := GenS3{F0: 18446744073709551615, F1: 0, F2: 79}
                    _ = v13
                }
                v14 := GenS1{F0: (v7 / (((v7 & 1)) | 1)), F1: (v7 | v7), F2: 279}
                _ = fi9
                _ = v14
            }
            _ = v8
        } else {
            v5 += v5
            for fi15 := 0; fi15 < 13; fi15++ {
                v16 := GenS4{F0: 1000, F1: (v7 % ((v7) | 1)), F2: 1}
                _ = fi15
                _ = v16
            }
            var v17 uint32 = 572
            _ = v17
        }
    default:
        fmt.Println(v5)
    }
    switch 32768 {
    case 127:
        v18 := GenS4{F0: 379, F1: (v7 - (v7 * (^(v7)))), F2: 127}
        _ = v18
    case 75:
        v19 := GenS3{F0: 4294967295, F1: 987, F2: 4294967295}
        _ = v19
    case 32768:
        _ = len(v1)
    default:
        v2 = v2
    }
    if true {
        for fi20 := 0; fi20 < 11; fi20++ {
            for fi21 := 0; fi21 < 18; fi21++ {
                v2 = 115
                _ = fi21
            }
            fmt.Println(v2, fi20)
            _ = fi20
        }
        var v22 uint32 = 42
        switch 127 {
        case 0:
            for fi23 := 0; fi23 < 14; fi23++ {
                var v24 uint64 = 2
                if (true || (2147483647 < fi23)) {
                    for fi25 := 0; fi25 < 16; fi25++ {
                        delete(v3, (v22 + 4294967295))
                        switch 4294967295 {
                        case 689:
                            var v26 int = fi23
                            _ = v26
                        case -9223372036854775808:
                            _ = len(v1)
                        }
                        _ = fi25
                    }
                    if ((v2 & -2147483648) > (v2 / ((int32(v7)) | 1))) {
                        var v27 uint32 = (v22 ^ (v22 | (^(v22))))
                        _ = v27
                    }
                    v28 := GenS4{F0: -32768, F1: v7, F2: 42}
                    _ = v28
                }
                _ = fi23
                _ = v24
            }
        default:
            switch -9223372036854775808 {
            case 9223372036854775807:
                var v29 float64 = (3.14159 / 1e10)
                _ = v29
            case -9223372036854775808:
                v7 = (v7 + 42)
            default:
                if (!true) {
                    v7 += (v7 - v7)
                    delete(v3, v22)
                    var v30 uint = 4294967295
                    _ = v30
                } else {
                    switch 1 {
                    case 74:
                        v7 = (^(v7))
                    }
                    v31 := GenS5{F0: 57, F1: 32768, F2: 168}
                    v2 += (v2 ^ 1)
                    _ = v31
                }
            }
        }
        _ = v22
    }
    fmt.Println(v5)
    for fi32 := 0; fi32 < 10; fi32++ {
        switch fi32 {
        case 2147483647:
            var v33 uint8 = 255
            _ = v33
        case -825:
            v1 = append(v1, "a")
        }
        _ = len(v1)
        _ = fi32
    }
    v5 += (((1.5 / (float32(0.0) - v5)) * ((v5 + 1e10) - 0.0)) / v5)
    fmt.Println(v7)
    v34 := GenS1{F0: v7, F1: (^(v7)), F2: 944}
    if (!(true || (!false))) {
        var v35 uint32 = 1
        _ = v35
    }
    fmt.Println(v5, v5, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v5
    _ = v6
    _ = v7
}
