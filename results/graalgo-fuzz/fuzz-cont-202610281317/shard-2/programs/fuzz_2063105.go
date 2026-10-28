package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 uint16 }
type GenS2 struct { F0 uint16; F1 int16; F2 bool }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 int8; F1 int32; F2 int16 }
type GenS5 struct { F0 uint64; F1 bool; F2 uint16 }
type GenS6 struct { F0 int64; F1 uint; F2 string }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int64; F1 uint }
type GenS9 struct { F0 string; F1 uint; F2 uint8 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 bool; F1 uint32; F2 string }
type GenS12 struct { F0 int8; F1 bool }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 int16; F1 int; F2 int8 }
type GenS15 struct { F0 uint16; F1 int64; F2 int32 }
type GenS16 struct { F0 uint; F1 int16; F2 uint64 }
type GenS17 struct { F0 string; F1 uint; F2 uint16 }

func main() {
    var v1 uint8 = 122
    var v2 uint16 = 255
    var v3 uint64 = 0
    v4 := []uint32{}
    var v5 int32 = 1
    var v6 uint = 4294967295
    if (!(!(2 < 42))) {
        if true {
            var v7 uint16 = 537
            for fi8 := 0; fi8 < 15; fi8++ {
                v9 := GenS1{F0: false, F1: (((("hello" + "") + "\n\t") + "\n\t") + ""), F2: (v7 << ((v6) & 63))}
                _ = fi8
                _ = v9
            }
            _ = v7
        } else {
            v4 = append(v4, 65535)
            if (42 >= 562) {
                var v10 uint64 = v3
                v11 := GenS1{F0: ((v3 * (^(v3))) == 42), F1: ("hello" + "a"), F2: v2}
                _ = v10
                _ = v11
            } else {
                v12 := GenS2{F0: v2, F1: -675, F2: (0.0 < 0.1)}
                var v13 uint8 = (^(v1))
                _ = v12
                _ = v13
            }
            if false {
                if false {
                    var v14 bool = false
                    _ = v14
                }
            }
        }
        for fi15 := 0; fi15 < 19; fi15++ {
            switch (fi15 / (((fi15 ^ (-(fi15)))) | 1)) {
            case 2147483647:
                v16 := GenS3{F0: -32768, F1: v6}
                _ = v16
            case 32768:
                fmt.Println(v2, fi15, v3)
            default:
                var v17 uint = (v6 * v6)
                _ = v17
            }
            v18 := GenS4{F0: 100, F1: (v5 >> ((567) & 63)), F2: 42}
            v2 += v2
            _ = fi15
            _ = v18
        }
        for fi19 := 0; fi19 < 17; fi19++ {
            fmt.Println(v3, fi19, v1)
            _ = len(v4)
            fmt.Println(v5, v6)
            _ = fi19
        }
    }
    v5 = (v5 ^ (v5 >> (((^(v6))) & 63)))
    for fi20 := 0; fi20 < 13; fi20++ {
        for fi21 := 0; fi21 < 20; fi21++ {
            var v22 int16 = int16(v6)
            var v23 float32 = 0.0
            v1 += uint8(v2)
            _ = fi21
            _ = v22
            _ = v23
        }
        fmt.Println(v3, v2, v3)
        _ = len(v4)
        _ = fi20
    }
    for fi24 := 0; fi24 < 11; fi24++ {
        fmt.Println(fi24, v3, v1)
        _ = fi24
    }
    for fi25 := 0; fi25 < 13; fi25++ {
        v2 += (^(v2))
        switch (("foo bar" + (("\n\t" + ("a" + "")) + (("hello" + ("\n\t" + "\n\t")) + ("foo bar" + "")))) + "a") {
        case "":
            v3 = uint64(v1)
        case "a":
            v3 = 0
        default:
            _ = len(v4)
        }
        v5 += (^(v5))
        _ = fi25
    }
    v3 -= v3
    fmt.Println(v5)
    v3 = (v3 << ((v6) & 63))
    v26 := GenS5{F0: v3, F1: true, F2: (v2 | 0)}
    v2 = (^(v2))
    for fi27 := 0; fi27 < 1; fi27++ {
        v2 += v2
        v6 -= v6
        _ = fi27
    }
    var v28 int32 = v5
    var v29 int = 42
    if len(v4) > 0 {
        v4[0] = 2
    }
    v4 = append(v4, 628)
    for fi30 := 0; fi30 < 17; fi30++ {
        switch "hello" {
        case "hello":
            switch fi30 {
            case 127:
                v4 = append(v4, 42)
            case -1:
                for fi31 := 0; fi31 < 14; fi31++ {
                    v28 = v28
                    if (1 > -706) {
                        if true {
                            var v32 float32 = (3.14159 - (3.14159 / (float32(0.1) / 1.0)))
                            switch -478 {
                            case 4294967295:
                                if true {
                                    if (!false) {
                                        if (false && true) {
                                            var v33 float64 = 0.0
                                            switch (-(v29)) {
                                            case 218:
                                                switch "" {
                                                case "hello":
                                                    v5 = (v28 / ((42) | 1))
                                                case "a":
                                                    v6 += uint(v3)
                                                default:
                                                    if true {
                                                        if (2 > (v1 ^ (v1 % ((v1) | 1)))) {
                                                            v34 := GenS2{F0: v2, F1: 32767, F2: true}
                                                            _ = v34
                                                        } else {
                                                            switch "\n\t" {
                                                            case "foo bar":
                                                                v32 -= (0.0 / (((0.0 + float32(v33)) / float32((v33 - float64((((v32 / v32) - v32) / 1.5))))) + ((v32 / v32) / (((v32 - (float32(v33) - v32)) * v32) + float32(v33)))))
                                                            case "a":
                                                                if true {
                                                                    fmt.Println(v6, fi31, fi31)
                                                                }
                                                            case "hello":
                                                                if ("foo bar" != "\n\t") {
                                                                    fmt.Println(fi30, v2)
                                                                    _ = len(v4)
                                                                    switch "foo bar" {
                                                                    case "a":
                                                                        if true {
                                                                            v35 := GenS3{F0: -955, F1: (v6 % ((18446744073709551615) | 1))}
                                                                            _ = v35
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                            var v36 uint8 = v1
                                                            fmt.Println(v32)
                                                            v36 += (^(v36))
                                                            switch fi31 {
                                                            case 128:
                                                                v37 := GenS6{F0: 42, F1: v6, F2: "foo bar"}
                                                                v6 = (v6 & (v6 | uint(v3)))
                                                                _ = v37
                                                            case 0:
                                                                if false {
                                                                    switch "a" {
                                                                    case "hello":
                                                                        if false {
                                                                            fmt.Println(v32, fi31, fi30)
                                                                            v36 = uint8(v6)
                                                                            fmt.Println(v33, v1)
                                                                        } else {
                                                                            v3 += v3
                                                                            if true {
                                                                                if false {
                                                                                    var v38 int8 = -128
                                                                                    switch "\n\t" {
                                                                                    case "a":
                                                                                        v6 -= 143
                                                                                    }
                                                                                    _ = v38
                                                                                }
                                                                                _ = len(v4)
                                                                                var v39 uint16 = uint16(v36)
                                                                                _ = v39
                                                                            } else {
                                                                                v40 := GenS3{F0: 0, F1: (v6 / ((uint(v2)) | 1))}
                                                                                if ((v3 * 1) >= v3) {
                                                                                    var v41 uint64 = v3
                                                                                    fmt.Println(v6, fi31, v3)
                                                                                    var v42 uint = 255
                                                                                    var v43 float64 = -2.25
                                                                                    _ = v41
                                                                                    _ = v42
                                                                                    _ = v43
                                                                                }
                                                                                _ = v40
                                                                            }
                                                                            v5 = v28
                                                                        }
                                                                    case "":
                                                                        switch fi31 {
                                                                        case 128:
                                                                            v44 := GenS6{F0: -9223372036854775808, F1: 622, F2: ("hello" + (("" + (("" + "hello") + "hello")) + (("a" + "foo bar") + "hello")))}
                                                                            _ = v44
                                                                        }
                                                                    case "a":
                                                                        v33 += (v33 - v33)
                                                                    default:
                                                                        if (!((v3 | v3) != (^(v3)))) {
                                                                            switch "foo bar" {
                                                                            case "\n\t":
                                                                                v45 := GenS7{F0: 0}
                                                                                _ = v45
                                                                            case "hello":
                                                                                var v46 float64 = (((float64((0.0 - 0.1)) * (v33 / 1e10)) / (1e10 - float64((-2.25 * (v32 / 0.1))))) - -2.25)
                                                                                v47 := GenS4{F0: -1, F1: (v28 / (((v28 / ((1) | 1))) | 1)), F2: 869}
                                                                                switch -9223372036854775808 {
                                                                                case 1:
                                                                                    v29 = fi30
                                                                                case 127:
                                                                                    v48 := GenS8{F0: -657, F1: (^(v6))}
                                                                                    _ = v48
                                                                                case 199:
                                                                                    v49 := GenS1{F0: true, F1: "foo bar", F2: v2}
                                                                                    switch (v29 + (fi30 + (fi30 / ((fi31) | 1)))) {
                                                                                    case 9223372036854775807:
                                                                                        var v50 int16 = -670
                                                                                        _ = v50
                                                                                    case 527:
                                                                                        var v51 uint64 = (v3 + (v3 % ((2) | 1)))
                                                                                        _ = v51
                                                                                    }
                                                                                    _ = len(v4)
                                                                                    v5 -= 705
                                                                                    if len(v4) >= 2 {
                                                                                        v4 = v4[1:len(v4)]
                                                                                    }
                                                                                    v52 := GenS2{F0: (v2 >> ((602) & 63)), F1: 32767, F2: true}
                                                                                    switch fi31 {
                                                                                    case 771:
                                                                                        switch 42 {
                                                                                        case 9223372036854775807:
                                                                                            if len(v4) > 0 {
                                                                                                v4[0] = 989
                                                                                            }
                                                                                        case -1:
                                                                                            v33 -= 0.0
                                                                                        default:
                                                                                            switch (fi31 + int(v5)) {
                                                                                            case -2147483648:
                                                                                                v53 := GenS2{F0: (^(v2)), F1: -51, F2: false}
                                                                                                _ = v53
                                                                                            }
                                                                                        }
                                                                                        var v54 uint64 = 42
                                                                                        v55 := GenS2{F0: v2, F1: -1, F2: true}
                                                                                        _ = v54
                                                                                        _ = v55
                                                                                    case -439:
                                                                                        v2 = 153
                                                                                    default:
                                                                                        var v56 string = "foo bar"
                                                                                        _ = v56
                                                                                    }
                                                                                    if len(v4) > 0 {
                                                                                        v4[0] = 42
                                                                                    }
                                                                                    _ = v49
                                                                                    _ = v52
                                                                                default:
                                                                                    if false {
                                                                                        if (false || ((v2 + (v2 / ((v2) | 1))) >= 1)) {
                                                                                            if false {
                                                                                                if ((!false) && ((fi30 & (-(fi31))) != 2147483647)) {
                                                                                                    var v57 int8 = 1
                                                                                                    var v58 int32 = int32(v3)
                                                                                                    _ = v57
                                                                                                    _ = v58
                                                                                                }
                                                                                            }
                                                                                        } else {
                                                                                            if false {
                                                                                                _ = len(v4)
                                                                                            } else {
                                                                                                var v59 int64 = 1
                                                                                                var v60 int32 = (v5 >> ((42) & 63))
                                                                                                _ = v59
                                                                                                _ = v60
                                                                                            }
                                                                                            v3 = 18446744073709551615
                                                                                            _ = len(v4)
                                                                                        }
                                                                                        var v61 bool = (4294967295 == (v6 | (^(v6))))
                                                                                        v29 = v29
                                                                                        _ = v61
                                                                                    } else {
                                                                                        _ = len(v4)
                                                                                    }
                                                                                }
                                                                                _ = v46
                                                                                _ = v47
                                                                            }
                                                                            v1 = uint8(v3)
                                                                            v1 = 127
                                                                        }
                                                                    }
                                                                }
                                                            case 32767:
                                                                fmt.Println(fi31, v28, v2)
                                                            default:
                                                                if true {
                                                                    if len(v4) > 0 {
                                                                        v4[0] = 42
                                                                    }
                                                                }
                                                            }
                                                            _ = v36
                                                        }
                                                        switch "a" {
                                                        case "a":
                                                            v33 = ((float64(1e10) + v33) + v33)
                                                        case "":
                                                            if true {
                                                                v3 = (^(v3))
                                                                switch "hello" {
                                                                case "\n\t":
                                                                    var v62 bool = false
                                                                    _ = v62
                                                                case "a":
                                                                    fmt.Println(v1, v1)
                                                                default:
                                                                    v3 -= v3
                                                                }
                                                                v4 = append(v4, 0)
                                                            } else {
                                                                var v63 float64 = 3.14159
                                                                _ = v63
                                                            }
                                                        case "\n\t":
                                                            v29 = (fi30 & fi30)
                                                        }
                                                    } else {
                                                        var v64 int32 = (v28 << (((^(v6))) & 63))
                                                        if true {
                                                            fmt.Println(v3)
                                                        } else {
                                                            v65 := GenS5{F0: (v3 - (^(v3))), F1: false, F2: v2}
                                                            fmt.Println(v28, v2, v3)
                                                            v33 -= 1.5
                                                            _ = v65
                                                        }
                                                        v64 = v64
                                                        var v66 int16 = 42
                                                        switch int(v28) {
                                                        case 32768:
                                                            var v67 uint64 = v3
                                                            v2 = v2
                                                            if false {
                                                                _ = len(v4)
                                                                v68 := GenS9{F0: "\n\t", F1: (v6 >> (((v6 % (((^(v6))) | 1))) & 63)), F2: v1}
                                                                if false {
                                                                    var v69 int64 = 1
                                                                    v67 += uint64(v6)
                                                                    v1 = v1
                                                                    _ = v69
                                                                }
                                                                var v70 bool = (0.0 > v32)
                                                                _ = v68
                                                                _ = v70
                                                            } else {
                                                                if (372 <= 0) {
                                                                    if (0.0 == 3.14159) {
                                                                        fmt.Println(v64, v33, v66)
                                                                    } else {
                                                                        v71 := GenS9{F0: "foo bar", F1: 65535, F2: 64}
                                                                        var v72 int16 = 1000
                                                                        _ = v71
                                                                        _ = v72
                                                                    }
                                                                    var v73 int8 = int8(v6)
                                                                    _ = v73
                                                                } else {
                                                                    switch (fi30 - (^(fi31))) {
                                                                    case 871:
                                                                        if len(v4) > 0 {
                                                                            v4[0] = 707
                                                                        }
                                                                    case 969:
                                                                        v74 := GenS10{F0: 9223372036854775807}
                                                                        _ = v74
                                                                    case 4294967295:
                                                                        v75 := GenS5{F0: (^(v67)), F1: true, F2: 0}
                                                                        _ = v75
                                                                    }
                                                                    if (!false) {
                                                                        v3 = v67
                                                                        var v76 float32 = float32(0.1)
                                                                        switch "" {
                                                                        case "":
                                                                            fmt.Println(fi30, v28, v66)
                                                                        case "hello":
                                                                            v77 := GenS8{F0: -837, F1: v6}
                                                                            _ = v77
                                                                        case "\n\t":
                                                                            v3 = v67
                                                                        default:
                                                                            var v78 uint = v6
                                                                            _ = v78
                                                                        }
                                                                        _ = v76
                                                                    } else {
                                                                        fmt.Println(v28, v32, v6)
                                                                        fmt.Println(v32, fi31, v33)
                                                                        v28 += v5
                                                                    }
                                                                    switch "\n\t" {
                                                                    case "hello":
                                                                        v32 += (v32 - v32)
                                                                    default:
                                                                        switch "foo bar" {
                                                                        case "\n\t":
                                                                            switch (("\n\t" + (("\n\t" + "foo bar") + "")) + "") {
                                                                            case "foo bar":
                                                                                v3 += 18446744073709551615
                                                                            }
                                                                        case "a":
                                                                            v79 := GenS11{F0: true, F1: 2, F2: (("a" + ("hello" + "\n\t")) + "a")}
                                                                            _ = v79
                                                                        case "hello":
                                                                            v2 -= 110
                                                                        }
                                                                    }
                                                                }
                                                                var v80 uint32 = 1
                                                                _ = v80
                                                            }
                                                            _ = v67
                                                        case 1:
                                                            if true {
                                                                v32 += 3.14159
                                                                var v81 int = fi31
                                                                switch "" {
                                                                case "":
                                                                    v66 -= (v66 << (((^(v6))) & 63))
                                                                case "hello":
                                                                    v28 += 2147483647
                                                                case "\n\t":
                                                                    switch (-(v29)) {
                                                                    case -1:
                                                                        if len(v4) >= 2 {
                                                                            v4 = v4[1:len(v4)]
                                                                        }
                                                                    case 1:
                                                                        v81 = fi30
                                                                    case 127:
                                                                        var v82 uint = (v6 << (((^(v6))) & 63))
                                                                        if (v2 > v2) {
                                                                            if (v5 < 2147483647) {
                                                                                var v83 int32 = int32(v2)
                                                                                v82 = 1
                                                                                v84 := GenS1{F0: false, F1: "hello", F2: (v2 + (v2 - (v2 / ((v2) | 1))))}
                                                                                _ = v83
                                                                                _ = v84
                                                                            } else {
                                                                                fmt.Println(v32, v64)
                                                                            }
                                                                            switch fi30 {
                                                                            case -16:
                                                                                v66 = v66
                                                                            }
                                                                            if len(v4) > 0 {
                                                                                v4[0] = 65535
                                                                            }
                                                                            if false {
                                                                                var v85 float64 = v33
                                                                                if ((fi31 <= v81) && (4294967295 != 348)) {
                                                                                    v5 = (v5 | v28)
                                                                                } else {
                                                                                    v86 := GenS12{F0: -41, F1: false}
                                                                                    if true {
                                                                                        _ = len(v4)
                                                                                    } else {
                                                                                        if false {
                                                                                            var v87 int32 = v28
                                                                                            _ = v87
                                                                                        }
                                                                                        v64 = (v5 * -421)
                                                                                        v33 += 1.5
                                                                                    }
                                                                                    _ = v86
                                                                                }
                                                                                v88 := GenS7{F0: int8(v1)}
                                                                                _ = v85
                                                                                _ = v88
                                                                            }
                                                                        }
                                                                        _ = v82
                                                                    }
                                                                }
                                                                fmt.Println(v29)
                                                                _ = v81
                                                            }
                                                        }
                                                        _ = v64
                                                        _ = v66
                                                    }
                                                }
                                            case 0:
                                                switch "a" {
                                                case "hello":
                                                    switch (v29 >> ((v6) & 63)) {
                                                    case 210:
                                                        var v89 int32 = 127
                                                        _ = v89
                                                    case -2147483648:
                                                        v5 -= int32(v2)
                                                        var v90 uint8 = 2
                                                        _ = len(v4)
                                                        var v91 int64 = 931
                                                        var v92 int32 = (v28 % ((v28) | 1))
                                                        _ = v90
                                                        _ = v91
                                                        _ = v92
                                                    }
                                                case "foo bar":
                                                    switch (-(v29)) {
                                                    case -695:
                                                        _ = len(v4)
                                                    }
                                                default:
                                                    v33 = 1.5
                                                }
                                            case 4294967295:
                                                if false {
                                                    if true {
                                                        v4 = append(v4, 2)
                                                    }
                                                } else {
                                                    fmt.Println(v1)
                                                    var v93 string = ("hello" + "a")
                                                    _ = len(v4)
                                                    _ = v93
                                                }
                                            }
                                            _ = v33
                                        } else {
                                            switch "\n\t" {
                                            case "foo bar":
                                                if true {
                                                    var v94 int16 = -1
                                                    var v95 int64 = 0
                                                    _ = v94
                                                    _ = v95
                                                } else {
                                                    v5 = v28
                                                    if (false && (!(!false))) {
                                                        v4 = append(v4, 65535)
                                                    }
                                                    v3 += (v3 >> (((v6 >> ((0) & 63))) & 63))
                                                }
                                            default:
                                                if true {
                                                    _ = len(v4)
                                                    var v96 int = (fi31 << ((255) & 63))
                                                    var v97 int64 = 963
                                                    v98 := GenS13{F0: 100}
                                                    _ = v96
                                                    _ = v97
                                                    _ = v98
                                                } else {
                                                    var v99 bool = ((!false) || (v32 != v32))
                                                    switch (v29 / ((fi31) | 1)) {
                                                    case 127:
                                                        v2 = uint16(v1)
                                                    default:
                                                        v4 = append(v4, 0)
                                                    }
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    var v100 int = (v29 | 9223372036854775807)
                                                    var v101 float32 = v32
                                                    v3 = 42
                                                    _ = v100
                                                    _ = v101
                                                    _ = v99
                                                }
                                                if (!(0 > 2147483647)) {
                                                    v102 := GenS14{F0: 951, F1: fi30, F2: 42}
                                                    v2 -= (^(v2))
                                                    var v103 float64 = -2.25
                                                    _ = v102
                                                    _ = v103
                                                }
                                                fmt.Println(fi31)
                                            }
                                            v5 -= v28
                                        }
                                        var v104 int8 = -128
                                        var v105 string = "a"
                                        _ = v104
                                        _ = v105
                                    }
                                    v2 += v2
                                    switch fi30 {
                                    case 128:
                                        var v106 string = "foo bar"
                                        _ = v106
                                    case 2147483647:
                                        _ = len(v4)
                                    case -1:
                                        if true {
                                            var v107 float32 = (v32 - ((1e10 + float32((0.1 - 1e10))) * v32))
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            _ = v107
                                        }
                                    }
                                } else {
                                    switch (^(fi30)) {
                                    case 0:
                                        v32 += (v32 / v32)
                                    default:
                                        v6 += (^(v6))
                                    }
                                }
                            case 2147483647:
                                if ("foo bar" <= ("\n\t" + ("foo bar" + "\n\t"))) {
                                    v108 := GenS11{F0: true, F1: uint32(v2), F2: "hello"}
                                    _ = v108
                                }
                            }
                            _ = v32
                        } else {
                            v109 := GenS15{F0: 618, F1: 1, F2: -1}
                            if (!true) {
                                fmt.Println(v6, v5, v28)
                                var v110 uint8 = (^(v1))
                                _ = v110
                            } else {
                                var v111 uint = v6
                                _ = v111
                            }
                            fmt.Println(v3, v2, v2)
                            _ = v109
                        }
                    } else {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        fmt.Println(v5, v28, v1)
                        var v112 uint8 = v1
                        v113 := GenS3{F0: 1, F1: v6}
                        _ = v112
                        _ = v113
                    }
                    if (!true) {
                        if (!false) {
                            v114 := GenS5{F0: 2, F1: (192 > (v28 | v5)), F2: (^(v2))}
                            var v115 string = "a"
                            fmt.Println(v115)
                            _ = v114
                            _ = v115
                        }
                        var v116 uint32 = 565
                        v28 = -1
                        _ = v116
                    }
                    _ = fi31
                }
            }
        default:
            v4 = append(v4, 0)
        }
        var v117 uint64 = uint64(v1)
        v3 += 0
        _ = fi30
        _ = v117
    }
    for fi118 := 0; fi118 < 11; fi118++ {
        var v119 uint64 = 0
        _ = len(v4)
        _ = fi118
        _ = v119
    }
    for fi120 := 0; fi120 < 2; fi120++ {
        switch "" {
        case "\n\t":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case "foo bar":
            v3 = 77
        case "":
            for fi121 := 0; fi121 < 18; fi121++ {
                _ = len(v4)
                v28 = v5
                _ = fi121
            }
        default:
            switch (fi120 & (fi120 + v29)) {
            case 9223372036854775807:
                if false {
                    v2 += v2
                } else {
                    switch "hello" {
                    case "":
                        v6 += (^(v6))
                    case "a":
                        var v122 int = fi120
                        _ = v122
                    case "hello":
                        var v123 bool = (uint32(v1) < 4294967295)
                        _ = v123
                    default:
                        for fi124 := 0; fi124 < 4; fi124++ {
                            if len(v4) > 0 {
                                v4[0] = 0
                            }
                            if true {
                                if false {
                                    v1 = v1
                                    if (886 >= 421) {
                                        var v125 int16 = int16(v2)
                                        if len(v4) > 0 {
                                            v4[0] = 1
                                        }
                                        v126 := GenS2{F0: (v2 * 2), F1: (v125 / (((v125 | (v125 % ((v125) | 1)))) | 1)), F2: false}
                                        _ = v125
                                        _ = v126
                                    } else {
                                        if len(v4) > 0 {
                                            v4[0] = 65535
                                        }
                                    }
                                    fmt.Println(v29)
                                } else {
                                    v127 := GenS16{F0: 42, F1: -1, F2: v3}
                                    if (((("foo bar" + "foo bar") + "\n\t") + "\n\t") < ("foo bar" + "\n\t")) {
                                        var v128 string = ""
                                        var v129 uint64 = 18446744073709551615
                                        v29 += (fi120 % (((fi124 * int(v1))) | 1))
                                        var v130 uint8 = (^(v1))
                                        _ = v128
                                        _ = v129
                                        _ = v130
                                    }
                                    var v131 int = fi120
                                    _ = v127
                                    _ = v131
                                }
                                v132 := GenS1{F0: true, F1: ("hello" + ("a" + "")), F2: (^(v2))}
                                _ = v132
                            }
                            _ = fi124
                        }
                    }
                    var v133 int8 = 1
                    switch "hello" {
                    case "foo bar":
                        switch ("\n\t" + "") {
                        case "":
                            fmt.Println(v29, v1)
                        case "\n\t":
                            var v134 bool = (true || true)
                            _ = v134
                        }
                    }
                    _ = v133
                }
            case 42:
                switch (((("a" + "foo bar") + "") + "\n\t") + (("\n\t" + (("" + ("" + "a")) + "\n\t")) + ("\n\t" + "hello"))) {
                case "\n\t":
                    switch (-(fi120)) {
                    case 9223372036854775807:
                        v2 -= (v2 & v2)
                    case 2147483647:
                        var v135 string = "hello"
                        _ = v135
                    case -865:
                        var v136 int = fi120
                        _ = v136
                    }
                case "":
                    if false {
                        var v137 float64 = (1e10 + (1.5 + 0.1))
                        _ = v137
                    } else {
                        for fi138 := 0; fi138 < 14; fi138++ {
                            _ = len(v4)
                            _ = fi138
                        }
                        v139 := GenS17{F0: ("" + ("" + ((("\n\t" + "\n\t") + "foo bar") + "hello"))), F1: 255, F2: (v2 & 244)}
                        var v140 int = int(v6)
                        _ = v139
                        _ = v140
                    }
                default:
                    var v141 bool = (-787 >= (v29 / (((-(fi120))) | 1)))
                    _ = v141
                }
            case 918:
                v1 = v1
            }
        }
        if (262 != uint64(v1)) {
            switch "foo bar" {
            case "\n\t":
                v6 = v6
            case "a":
                v2 = (^(v2))
            }
            v5 -= v5
        }
        v1 += (v1 * 42)
        _ = fi120
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
