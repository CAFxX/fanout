package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int; F1 uint32 }
type GenS6 struct { F0 int32; F1 int8 }
type GenS7 struct { F0 uint8; F1 uint32; F2 uint8 }
type GenS8 struct { F0 int16; F1 int64 }
type GenS9 struct { F0 int64; F1 int16 }
type GenS10 struct { F0 uint32; F1 bool }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint8 = 205
    v3 := []int{128}
    var v4 uint64 = 18446744073709551615
    v2 += v2
    if false {
        v5 := GenS1{F0: 4294967295, F1: (v4 / (((v4 & v4)) | 1))}
        if (!false) {
            v2 -= (v2 + v2)
            v6 := GenS2{F0: 42, F1: int32(v2)}
            _ = v6
        } else {
            var v7 bool = true
            if v7 {
                if (32767 == 0) {
                    switch 127 {
                    case 42:
                        v4 = (^(v4))
                    }
                    if (2147483647 > 711) {
                        for fi8 := 0; fi8 < 13; fi8++ {
                            if v7 {
                                for fi9 := 0; fi9 < 14; fi9++ {
                                    if true {
                                        var v10 uint8 = v2
                                        var v11 uint8 = 1
                                        var v12 float64 = 0.1
                                        _ = v10
                                        _ = v11
                                        _ = v12
                                    }
                                    var v13 uint64 = v4
                                    switch ("a" + "foo bar") {
                                    case "foo bar":
                                        fmt.Println(v7, v7)
                                    case "hello":
                                        var v14 uint32 = 4
                                        _ = v14
                                    }
                                    switch fi8 {
                                    case -9223372036854775808:
                                        v15 := GenS3{F0: 127}
                                        _ = v15
                                    }
                                    _ = fi9
                                    _ = v13
                                }
                                var v16 uint16 = 0
                                if v7 {
                                    v17 := GenS2{F0: 79, F1: 0}
                                    _ = v17
                                }
                                _ = v16
                            }
                            _ = fi8
                        }
                        fmt.Println(v2)
                        v4 -= 81
                    }
                } else {
                    if ((v2 * (v2 ^ 255)) < v2) {
                        if true {
                            var v18 int32 = 0
                            switch "foo bar" {
                            case "foo bar":
                                if (v2 < 193) {
                                    var v19 int16 = -819
                                    v20 := GenS2{F0: v18, F1: v18}
                                    for fi21 := 0; fi21 < 9; fi21++ {
                                        v3 = append(v3, 238)
                                        var v22 uint16 = 32767
                                        _ = fi21
                                        _ = v22
                                    }
                                    _ = v19
                                    _ = v20
                                } else {
                                    v4 += (v4 + 2)
                                    v1[2147483648] = 4294967295
                                    fmt.Println(v18)
                                }
                            }
                            for fi23 := 0; fi23 < 17; fi23++ {
                                v24 := v1[142]
                                _ = fi23
                                _ = v24
                            }
                            _ = v18
                        }
                    }
                }
            } else {
                fmt.Println(v7, v7)
                var v25 uint64 = (v4 ^ (v4 | v4))
                _ = v25
            }
            v7 = v7
            _ = v7
        }
        _ = v5
    } else {
        fmt.Println(v2, v2)
        fmt.Println(v2, v2)
    }
    for fi26 := 0; fi26 < 4; fi26++ {
        var v27 bool = (1 >= 2147483647)
        v2 -= (v2 * (^(v2)))
        _ = fi26
        _ = v27
    }
    if false {
        switch ("a" + "\n\t") {
        case "":
            switch "" {
            case "hello":
                for fi28 := 0; fi28 < 7; fi28++ {
                    if false {
                        if false {
                            v29 := GenS4{F0: (v4 < v4)}
                            var v30 int = (fi28 >> ((18446744073709551615) & 63))
                            for fi31 := 0; fi31 < 16; fi31++ {
                                var v32 uint8 = v2
                                _ = fi31
                                _ = v32
                            }
                            _ = v29
                            _ = v30
                        }
                        switch (fi28 / (((fi28 ^ 32768)) | 1)) {
                        case -433:
                            v4 = (v4 ^ v4)
                        case 0:
                            var v33 int = (^(fi28))
                            _ = v33
                        default:
                            v34 := v1[37]
                            _ = v34
                        }
                        v35 := GenS5{F0: (fi28 * 684), F1: 474}
                        _ = v35
                    } else {
                        for fi36 := 0; fi36 < 2; fi36++ {
                            fmt.Println(v4)
                            _ = fi36
                        }
                    }
                    var v37 float32 = 1e10
                    var v38 string = ("hello" + "foo bar")
                    _ = fi28
                    _ = v37
                    _ = v38
                }
            case "a":
                var v39 int8 = -128
                _ = v39
            default:
                v4 -= 2
            }
        }
        var v40 int64 = 9223372036854775807
        for fi41 := 0; fi41 < 20; fi41++ {
            var v42 int8 = 100
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if (false || (-499 > (fi41 / ((4294967295) | 1)))) {
                var v43 uint8 = (v2 % ((v2) | 1))
                _ = v43
            }
            _ = fi41
            _ = v42
        }
        _ = v40
    }
    switch ("hello" + "foo bar") {
    case "\n\t":
        switch "a" {
        case "hello":
            fmt.Println(v4, v4)
        }
    case "hello":
        switch "\n\t" {
        case "\n\t":
            v4 = v4
        case "a":
            var v44 int32 = -711
            _ = v44
        case "":
            for fi45 := 0; fi45 < 5; fi45++ {
                var v46 bool = false
                switch "\n\t" {
                case "foo bar":
                    var v47 uint = 1
                    _ = v47
                case "\n\t":
                    if (float64(1.5) == ((-2.25 - float64(float32(((((1e10 - -2.25) + ((-2.25 + (3.14159 * 0.0)) + ((float64(1.5) + (float64(0.1) * 0.0)) / (0.1 - float64(1.5))))) * 1e10) / 1.5)))) - 0.0)) {
                        var v48 string = "\n\t"
                        for fi49 := 0; fi49 < 6; fi49++ {
                            delete(v1, 345)
                            var v50 uint64 = (v4 ^ (v4 % ((v4) | 1)))
                            v50 += (v50 << ((2) & 63))
                            _ = fi49
                            _ = v50
                        }
                        v51 := GenS6{F0: -2147483648, F1: -1}
                        _ = v48
                        _ = v51
                    }
                }
                _ = fi45
                _ = v46
            }
        default:
            v52 := GenS3{F0: 115}
            _ = v52
        }
    case "a":
        var v53 float32 = 0.1
        _ = v53
    }
    if ((v4 + 311) != uint64(v2)) {
        v2 = (v2 & v2)
    } else {
        fmt.Println(v4, v2)
        var v54 uint16 = 255
        _ = v54
    }
    if (true || true) {
        var v55 int8 = 110
        var v56 int8 = (-(v55))
        _ = v55
        _ = v56
    }
    switch 1 {
    case 32767:
        var v57 float64 = float64(1e10)
        _ = v57
    case 519:
        fmt.Println(v2, v2)
    }
    for fi58 := 0; fi58 < 15; fi58++ {
        v59 := GenS7{F0: v2, F1: 2, F2: v2}
        v60 := GenS5{F0: fi58, F1: 2}
        for fi61 := 0; fi61 < 5; fi61++ {
            switch (fi58 - fi61) {
            case 666:
                v62 := GenS8{F0: 749, F1: 9223372036854775807}
                _ = v62
            case 128:
                switch fi61 {
                case 0:
                    v63 := GenS5{F0: -936, F1: 2}
                    _ = v63
                case -9223372036854775808:
                    v64 := v1[506]
                    _ = v64
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            if (uint16(v2) >= 255) {
                v4 -= (v4 | (v4 * v4))
                if false {
                    var v65 bool = true
                    if v65 {
                        v66 := GenS9{F0: -789, F1: -477}
                        v65 = true
                        _ = v66
                    } else {
                        switch "a" {
                        case "hello":
                            v65 = (2147483647 <= 769)
                        case "":
                            var v67 string = ("hello" + "\n\t")
                            _ = v67
                        }
                        v68 := GenS1{F0: (v4 & v4), F1: v4}
                        var v69 uint8 = (v2 * v2)
                        _ = v68
                        _ = v69
                    }
                    _ = v65
                } else {
                    switch 127 {
                    case 1:
                        switch -1 {
                        case 503:
                            switch (("\n\t" + ("\n\t" + "\n\t")) + (((("" + "\n\t") + "foo bar") + "foo bar") + "foo bar")) {
                            case "a":
                                if (!(false || true)) {
                                    switch ("a" + ("" + "a")) {
                                    case "foo bar":
                                        v2 = (v2 & 128)
                                    case "a":
                                        var v70 bool = (((!true) || false) || ((^(v4)) != 4294967295))
                                        _ = v70
                                    }
                                } else {
                                    fmt.Println(fi61)
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    v2 -= v2
                                }
                            case "\n\t":
                                if false {
                                    v2 += (v2 & (v2 / ((235) | 1)))
                                    switch -9223372036854775808 {
                                    case 820:
                                        fmt.Println(fi61)
                                    default:
                                        if (true || false) {
                                            var v71 int64 = 1
                                            if false {
                                                var v72 uint64 = v4
                                                var v73 uint16 = 346
                                                if (("" + "foo bar") < (("" + "") + "\n\t")) {
                                                    var v74 int = 4294967295
                                                    v4 -= (v72 * (v4 % ((2) | 1)))
                                                    _ = v74
                                                } else {
                                                    v72 = v4
                                                }
                                                _ = v72
                                                _ = v73
                                            }
                                            _ = v71
                                        }
                                    }
                                }
                            default:
                                v4 = (^(v4))
                            }
                        }
                    case 0:
                        v4 += (v4 >> ((255) & 63))
                        v4 -= 1
                    }
                    v75 := GenS7{F0: v2, F1: 733, F2: (v2 * (^(v2)))}
                    if true {
                        var v76 float32 = 0.0
                        var v77 int64 = 9223372036854775807
                        v77 += (v77 % (((v77 << ((uint(v2)) & 63))) | 1))
                        _ = v76
                        _ = v77
                    } else {
                        _ = len(v1)
                        v4 = v4
                        fmt.Println(fi61)
                    }
                    v78 := GenS5{F0: (fi61 - (fi58 >> ((900) & 63))), F1: 171}
                    fmt.Println(v2, fi58)
                    var v79 int64 = 0
                    _ = v75
                    _ = v78
                    _ = v79
                }
            }
            _ = fi61
        }
        _ = fi58
        _ = v59
        _ = v60
    }
    var v80 uint32 = 0
    for fi81 := 0; fi81 < 5; fi81++ {
        v82 := GenS10{F0: (v80 * v80), F1: false}
        _ = fi81
        _ = v82
    }
    _ = len(v1)
    var v83 bool = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v80:", v80)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v80
    _ = v83
}
