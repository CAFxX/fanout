package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64; F1 int; F2 int64 }
type GenS4 struct { F0 int64; F1 int32 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int64; F1 uint8; F2 int8 }
func genFn0(p0 int64, p1 int32, p2 int64) uint32 {
    p1 = 0
    for fi4 := 0; fi4 < 3; fi4++ {
        for fi5 := 0; fi5 < 7; fi5++ {
            switch ("" + "\n\t") {
            case "\n\t":
                p0 -= (p2 * p0)
            case "foo bar":
                v6 := []int{(fi5 / ((1) | 1)), -283}
                _ = v6
            case "":
                var v7 uint16 = 442
                _ = v7
            default:
                p1 += (^(p1))
            }
            var v8 int16 = 1000
            if false {
                if false {
                    switch fi4 {
                    case 32767:
                        v9 := GenS1{F0: 1, F1: 255, F2: 42}
                        _ = v9
                    default:
                        if (fi4 != (fi4 / (((fi4 << ((2) & 63))) | 1))) {
                            if (!(42 <= 255)) {
                                v10 := []int{fi4, (fi5 + 9223372036854775807)}
                                _ = v10
                            }
                            v11 := GenS2{F0: 428}
                            _ = v11
                        } else {
                            p2 = p0
                            var v12 int64 = (p0 << ((0) & 63))
                            _ = v12
                        }
                    }
                    v13 := []int{}
                    if (515 <= 4294967295) {
                        if ((0 != 114) || true) {
                            switch "foo bar" {
                            case "hello":
                                v14 := GenS1{F0: 0, F1: 0, F2: 32767}
                                _ = v14
                            case "":
                                fmt.Println(fi5, p1, v8)
                            }
                        }
                    } else {
                        v15 := GenS3{F0: 42, F1: 32767, F2: (p0 / ((9223372036854775807) | 1))}
                        _ = v15
                    }
                    switch (("a" + ("a" + ("\n\t" + ""))) + "a") {
                    case "a":
                        v16 := GenS1{F0: 0, F1: 32767, F2: 65535}
                        _ = v16
                    default:
                        v17 := GenS4{F0: p0, F1: p1}
                        _ = v17
                    }
                    _ = v13
                }
                v18 := []int{}
                _ = v18
            } else {
                var v19 uint8 = 121
                _ = v19
            }
            _ = fi5
            _ = v8
        }
        fmt.Println(fi4)
        fmt.Println(p0, p1, p1)
        _ = fi4
    }
    v20 := []int{0, -2147483648, -9223372036854775808}
    var v21 uint8 = 155
    v22 := GenS2{F0: 370}
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    _ = v21
    _ = v22
    return 0
}

func main() {
    var v1 uint8 = 254
    v2 := make(map[string]uint32)
    v3 := []int{-238, 4294967295}
    var v23 int32 = 42
    var v24 uint32 = 2147483647
    fmt.Println(v24, v24, v23)
    v24 = v24
    v23 += 0
    var v25 string = "foo bar"
    switch 42 {
    case 1:
        fmt.Println(v25, v25)
    }
    v23 += -904
    var v26 uint8 = 251
    v27 := v2[v25]
    if (-1 == 9223372036854775807) {
        switch -588 {
        case 1:
            fmt.Println(v23)
        case -9223372036854775808:
            v28 := GenS5{F0: 2}
            _ = v28
        case 42:
            v26 += (v26 & 1)
        }
        if true {
            v25 = v25
            for fi29 := 0; fi29 < 6; fi29++ {
                for fi30 := 0; fi30 < 9; fi30++ {
                    var v31 float32 = 0.1
                    _ = fi30
                    _ = v31
                }
                var v32 int8 = -55
                _ = fi29
                _ = v32
            }
            for fi33 := 0; fi33 < 18; fi33++ {
                v34 := GenS2{F0: 250}
                _ = fi33
                _ = v34
            }
        } else {
            for fi35 := 0; fi35 < 17; fi35++ {
                v36 := GenS6{F0: 4294967295}
                var v37 uint = 751
                _ = fi35
                _ = v36
                _ = v37
            }
            for fi38 := 0; fi38 < 17; fi38++ {
                if ((42 > 402) || ((false || ((uint16(v27) > 176) && false)) || ((^(v1)) >= 127))) {
                    v26 -= uint8(v24)
                    v39 := GenS5{F0: 0}
                    _ = v39
                } else {
                    v27 -= v24
                    fmt.Println(v25)
                    if (3.14159 < (1e10 - ((3.14159 + (float64((-2.25 * 1.5)) / -2.25)) + -2.25))) {
                        var v40 int32 = 1
                        var v41 uint = 4294967295
                        _ = v40
                        _ = v41
                    } else {
                        for fi42 := 0; fi42 < 16; fi42++ {
                            if true {
                                v27 -= (v24 | v27)
                            } else {
                                var v43 int16 = -702
                                v44 := GenS7{F0: -9223372036854775808, F1: (v26 & 2), F2: 123}
                                var v45 int64 = 9223372036854775807
                                _ = v43
                                _ = v44
                                _ = v45
                            }
                            v24 = (v27 % ((v24) | 1))
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = fi42
                        }
                    }
                }
                var v46 int64 = -591
                switch v25 {
                case "":
                    v47 := GenS4{F0: 762, F1: (-(v23))}
                    _ = v47
                case "foo bar":
                    switch 4294967295 {
                    case 32768:
                        var v48 string = v25
                        _ = v48
                    case 36:
                        v1 += 187
                    case -706:
                        switch (-(fi38)) {
                        case 4294967295:
                            var v49 uint32 = (v24 + 549)
                            _ = v49
                        case -2147483648:
                            switch 32768 {
                            case -2147483648:
                                fmt.Println(v23, v23, v26)
                            }
                        case -381:
                            _ = len(v3)
                        }
                    default:
                        var v50 bool = (false && true)
                        _ = v50
                    }
                case "\n\t":
                    var v51 float64 = (0.1 * float64(0.0))
                    _ = v51
                default:
                    v27 = v27
                }
                _ = fi38
                _ = v46
            }
            switch 2147483647 {
            case 0:
                fmt.Println(v24)
            default:
                var v52 uint8 = (v1 >> ((1) & 63))
                _ = v52
            }
        }
    }
    var v53 uint = 1
    for fi54 := 0; fi54 < 15; fi54++ {
        var v55 uint64 = 0
        if ((fi54 & (-(fi54))) <= fi54) {
            switch "hello" {
            case "":
                _ = len(v2)
            case "foo bar":
                v2[((v25 + v25) + (v25 + "a"))] = (v27 % ((genFn0(1, -2147483648, -9223372036854775808)) | 1))
            default:
                v56 := v2[v25]
                _ = v56
            }
        }
        switch (fi54 * fi54) {
        case 2147483647:
            var v57 float32 = float32(((0.0 * 1.5) + -2.25))
            _ = v57
        case 537:
            v27 += genFn0(-1, v23, -219)
        case 127:
            var v58 uint8 = v1
            _ = v58
        default:
            for fi59 := 0; fi59 < 6; fi59++ {
                v3 = append(v3, (fi59 % ((int(v53)) | 1)))
                var v60 uint32 = genFn0(9223372036854775807, -443, 42)
                _ = fi59
                _ = v60
            }
        }
        _ = fi54
        _ = v55
    }
    for fi61 := 0; fi61 < 15; fi61++ {
        v24 = genFn0(42, v23, 1)
        var v62 uint8 = v1
        _ = fi61
        _ = v62
    }
    for fi63 := 0; fi63 < 2; fi63++ {
        switch -1 {
        case 128:
            fmt.Println(v27)
        default:
            v1 = (^(v26))
        }
        _ = fi63
    }
    fmt.Println(v1)
    for fi64 := 0; fi64 < 20; fi64++ {
        switch fi64 {
        case 4294967295:
            var v65 float64 = 1.5
            _ = v65
        case 770:
            if false {
                var v66 int64 = -1
                if false {
                    delete(v2, (v25 + (v25 + "")))
                    switch (v25 + "a") {
                    case "hello":
                        switch 975 {
                        case 832:
                            v53 -= 65535
                        case -569:
                            v26 += (^(v1))
                        default:
                            v26 = v1
                        }
                    case "a":
                        delete(v2, v25)
                    }
                }
                for fi67 := 0; fi67 < 18; fi67++ {
                    switch (fi64 - fi64) {
                    case 127:
                        switch "" {
                        case "hello":
                            v26 = v26
                        case "":
                            if (!true) {
                                switch (-(fi64)) {
                                case -664:
                                    v68 := GenS6{F0: 1}
                                    v27 += (v24 & v24)
                                    v23 = v23
                                    if false {
                                        if (true || true) {
                                            v24 += (v24 << (((v53 << (((v53 << (((v53 - (v53 / (((v53 - 1)) | 1)))) & 63))) & 63))) & 63))
                                            var v69 uint = (v53 + (v53 / ((uint(v24)) | 1)))
                                            if (((v66 | 738) == 116) || (42 == (fi67 | fi64))) {
                                                if true {
                                                    if true {
                                                        var v70 uint8 = (v1 ^ 2)
                                                        v3 = append(v3, (fi64 << ((v69) & 63)))
                                                        _ = v70
                                                    } else {
                                                        _ = len(v3)
                                                        fmt.Println(v24, v53)
                                                        v26 += 0
                                                        var v71 string = (v25 + "foo bar")
                                                        v26 = (^(v1))
                                                        _ = v71
                                                    }
                                                    _ = len(v3)
                                                    v25 = "\n\t"
                                                }
                                                v53 += v53
                                                var v72 uint32 = v24
                                                v25 = v25
                                                switch v25 {
                                                case "\n\t":
                                                    v23 = int32(v66)
                                                case "":
                                                    var v73 float64 = 0.0
                                                    switch fi67 {
                                                    case -2147483648:
                                                        if false {
                                                            var v74 float64 = 3.14159
                                                            if (1e10 != 3.14159) {
                                                                v24 += uint32(v26)
                                                                v25 = "\n\t"
                                                                var v75 int = fi67
                                                                _ = v75
                                                            }
                                                            _ = v74
                                                        }
                                                    case -938:
                                                        fmt.Println(fi64, v66, v66)
                                                    case 4294967295:
                                                        fmt.Println(v69)
                                                    }
                                                    _ = v73
                                                case "hello":
                                                    fmt.Println(fi67)
                                                }
                                                _ = v72
                                            } else {
                                                v66 += 1
                                                var v76 uint = (v53 & (v53 - v69))
                                                _ = v76
                                            }
                                            _ = v69
                                        }
                                        _ = len(v3)
                                    }
                                    _ = v68
                                case 9223372036854775807:
                                    v66 = v66
                                case 2147483647:
                                    var v77 float64 = ((0.1 * 1e10) / 1.5)
                                    _ = v77
                                }
                                switch fi67 {
                                case -395:
                                    v78 := GenS3{F0: (v66 - (v66 - v66)), F1: fi67, F2: v66}
                                    _ = v78
                                case -2147483648:
                                    v79 := GenS6{F0: 2}
                                    _ = v79
                                }
                                if (838 <= 425) {
                                    if true {
                                        var v80 uint64 = 2
                                        v81 := GenS4{F0: (v66 - 880), F1: v23}
                                        v82 := GenS7{F0: v66, F1: (v1 - 2), F2: 0}
                                        _ = v80
                                        _ = v81
                                        _ = v82
                                    }
                                    v66 += -666
                                    var v83 uint = (v53 | (^(v53)))
                                    _ = v83
                                }
                            }
                        case "\n\t":
                            var v84 uint16 = 255
                            _ = v84
                        default:
                            var v85 uint16 = 42
                            _ = v85
                        }
                    }
                    fmt.Println(v66, v1)
                    _ = fi67
                }
                _ = v66
            } else {
                var v86 uint8 = v1
                _ = v86
            }
        default:
            var v87 int16 = 32767
            _ = v87
        }
        if true {
            var v88 uint8 = 0
            _ = v88
        }
        fmt.Println(v27, v27, v27)
        _ = fi64
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v53
}
