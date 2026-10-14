package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 int }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32; F1 int8 }
type GenS5 struct { F0 bool; F1 string; F2 int16 }
type GenS6 struct { F0 uint8; F1 uint32 }
type GenS7 struct { F0 int16; F1 int }
type GenS8 struct { F0 int32; F1 int8; F2 int }
type GenS9 struct { F0 uint32; F1 uint32; F2 uint64 }
type GenS10 struct { F0 uint64; F1 uint32 }
type GenS11 struct { F0 int8; F1 int32; F2 int16 }
type GenS12 struct { F0 int64; F1 uint; F2 int32 }

func main() {
    var v1 bool = true
    var v2 int16 = 1000
    v3 := []uint32{537}
    var v4 uint32 = 1
    for fi5 := 0; fi5 < 13; fi5++ {
        switch (^(fi5)) {
        case 127:
            for fi6 := 0; fi6 < 11; fi6++ {
                var v7 float64 = float64(-2.25)
                _ = fi6
                _ = v7
            }
        case 750:
            v4 += v4
        }
        v1 = false
        v4 = (v4 ^ v4)
        _ = fi5
    }
    fmt.Println(v4)
    if len(v3) > 0 {
        v3[0] = v4
    }
    if (229 > 42) {
        switch 32767 {
        case 375:
            var v8 float32 = (((3.14159 - 0.0) / float32((0.1 - (-2.25 * 0.1)))) / 1.0)
            _ = v8
        case 2147483647:
            v9 := GenS1{F0: 2, F1: (false || true), F2: 42}
            _ = v9
        }
    }
    _ = len(v3)
    for fi10 := 0; fi10 < 13; fi10++ {
        var v11 int64 = -9223372036854775808
        var v12 uint8 = 128
        for fi13 := 0; fi13 < 20; fi13++ {
            switch ("" + ("" + (("" + ("\n\t" + "a")) + ((("foo bar" + ((("\n\t" + "foo bar") + ("a" + "hello")) + "foo bar")) + "hello") + "a")))) {
            case "a":
                if v1 {
                    v14 := GenS2{F0: (v2 | (v2 - v2)), F1: 18446744073709551615}
                    v15 := GenS1{F0: uint(v12), F1: (v1 && ((((0.1 * float64((-2.25 / 1.0))) - float64(((1e10 + 0.1) - -2.25))) / -2.25) <= 1e10)), F2: fi13}
                    _ = v14
                    _ = v15
                }
                fmt.Println(v2, fi10)
                _ = len(v3)
                switch (fi13 - (fi13 % ((fi10) | 1))) {
                case -52:
                    if (false || v1) {
                        fmt.Println(v2)
                        var v16 bool = (0.0 != -2.25)
                        _ = v16
                    }
                case -231:
                    var v17 uint64 = 0
                    _ = v17
                case -9223372036854775808:
                    v18 := GenS2{F0: 1, F1: 1}
                    v11 = (v11 - (v11 << ((885) & 63)))
                    _ = v18
                default:
                    v12 += 42
                }
            case "hello":
                v11 -= v11
            case "":
                var v19 int = -663
                _ = v19
            }
            _ = fi13
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    if (!(!true)) {
        var v20 uint8 = 255
        _ = v20
    } else {
        for fi21 := 0; fi21 < 3; fi21++ {
            switch "hello" {
            case "a":
                var v22 int32 = 231
                _ = v22
            case "\n\t":
                v23 := GenS3{F0: -84}
                _ = v23
            default:
                var v24 int = (fi21 * (fi21 - 4294967295))
                _ = v24
            }
            _ = fi21
        }
        var v25 uint32 = (v4 + (^(v4)))
        _ = v25
    }
    var v26 uint16 = 255
    var v27 int64 = -9223372036854775808
    v3 = append(v3, v4)
    fmt.Println(v4, v27)
    var v28 string = "hello"
    for fi29 := 0; fi29 < 19; fi29++ {
        switch v28 {
        case "hello":
            v2 = -32768
        case "\n\t":
            fmt.Println(v2, v26, v1)
        }
        if v1 {
            for fi30 := 0; fi30 < 14; fi30++ {
                switch "a" {
                case "":
                    v1 = v1
                case "foo bar":
                    switch v28 {
                    case "":
                        v4 = 2147483648
                    case "foo bar":
                        var v31 uint16 = uint16(v4)
                        v31 = v26
                        _ = v31
                    }
                case "hello":
                    var v32 int32 = 610
                    _ = v32
                }
                _ = fi30
            }
            v4 = 0
            var v33 bool = v1
            _ = v33
        } else {
            if false {
                switch (^(fi29)) {
                case 32767:
                    v4 = (v4 << ((824) & 63))
                case -1:
                    for fi34 := 0; fi34 < 17; fi34++ {
                        v35 := GenS1{F0: 65535, F1: ((v1 || (!false)) && (1 == 142)), F2: fi29}
                        v36 := GenS4{F0: 2147483647, F1: 127}
                        _ = fi34
                        _ = v35
                        _ = v36
                    }
                }
                v2 = (-(v2))
                switch "a" {
                case "hello":
                    if (137 >= 0) {
                        switch ((v28 + v28) + v28) {
                        case "":
                            v4 -= v4
                        case "\n\t":
                            for fi37 := 0; fi37 < 9; fi37++ {
                                if true {
                                    if (false || ((^(fi37)) != (fi37 & (fi37 | 9223372036854775807)))) {
                                        var v38 uint16 = (v26 << ((uint(v26)) & 63))
                                        var v39 int32 = 29
                                        _ = v38
                                        _ = v39
                                    } else {
                                        if (v1 && v1) {
                                            fmt.Println(fi29)
                                            _ = len(v3)
                                            var v40 float64 = (((3.14159 / (0.0 - (1.5 / 0.1))) / 0.1) * 0.1)
                                            _ = v40
                                        } else {
                                            var v41 uint32 = v4
                                            if false {
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                var v42 uint16 = (v26 << ((18446744073709551615) & 63))
                                                _ = v42
                                            }
                                            var v43 int8 = 1
                                            switch int(v43) {
                                            case -483:
                                                fmt.Println(v2, v2, v27)
                                            case 32768:
                                                v3 = append(v3, 2)
                                            case 1:
                                                if false {
                                                    var v44 float32 = 0.1
                                                    _ = v44
                                                } else {
                                                    var v45 int16 = v2
                                                    switch ((((v28 + v28) + "hello") + v28) + v28) {
                                                    case "":
                                                        fmt.Println(v28, v27, v41)
                                                    case "hello":
                                                        v46 := GenS2{F0: -1, F1: 2}
                                                        _ = v46
                                                    default:
                                                        v27 = -9223372036854775808
                                                    }
                                                    _ = v45
                                                }
                                                if true {
                                                    v1 = ((v43 ^ v43) == v43)
                                                    _ = len(v3)
                                                }
                                                fmt.Println(v4, v26)
                                                v3 = append(v3, v4)
                                                v47 := GenS2{F0: v2, F1: 0}
                                                _ = v47
                                            default:
                                                if len(v3) > 0 {
                                                    v3[0] = (v4 | (v4 ^ (^(v4))))
                                                }
                                            }
                                            switch (("\n\t" + (v28 + v28)) + v28) {
                                            case "hello":
                                                v41 = (^(v41))
                                            }
                                            if ((v2 ^ -1) != (v2 << ((1) & 63))) {
                                                if v1 {
                                                    switch "foo bar" {
                                                    case "a":
                                                        fmt.Println(v28, v26, v43)
                                                    case "hello":
                                                        if (v1 || (4294967295 < uint(v26))) {
                                                            v48 := GenS1{F0: 42, F1: (255 != 33), F2: (^(fi37))}
                                                            v49 := GenS5{F0: v1, F1: v28, F2: (v2 * int16(fi37))}
                                                            v27 = (-(v27))
                                                            _ = v48
                                                            _ = v49
                                                        }
                                                    case "":
                                                        v3 = append(v3, 2147483647)
                                                    }
                                                    var v50 int32 = 949
                                                    var v51 uint8 = 255
                                                    _ = v50
                                                    _ = v51
                                                }
                                                if (false && v1) {
                                                    v1 = v1
                                                    switch "" {
                                                    case "a":
                                                        v43 -= (v43 + int8(v27))
                                                    }
                                                } else {
                                                    switch "" {
                                                    case "hello":
                                                        if v1 {
                                                            _ = len(v3)
                                                            v52 := GenS4{F0: 42, F1: (-(v43))}
                                                            _ = v52
                                                        }
                                                    case "foo bar":
                                                        v2 += v2
                                                    case "a":
                                                        switch (^(fi37)) {
                                                        case 32767:
                                                            if ((!(v2 >= -1)) && v1) {
                                                                v3 = append(v3, v41)
                                                            }
                                                            v41 -= (v41 | v4)
                                                        case 1:
                                                            if v1 {
                                                                var v53 uint = 9
                                                                v26 = (v26 / (((v26 - uint16(v4))) | 1))
                                                                switch "hello" {
                                                                case "":
                                                                    v27 -= (v27 & (v27 + v27))
                                                                case "\n\t":
                                                                    var v54 float32 = (1e10 / 1e10)
                                                                    _ = v54
                                                                default:
                                                                    switch v28 {
                                                                    case "\n\t":
                                                                        if (1e10 >= float64(-2.25)) {
                                                                            v1 = v1
                                                                        } else {
                                                                            if len(v3) > 0 {
                                                                                v3[0] = (v41 / ((661) | 1))
                                                                            }
                                                                            var v55 float32 = 3.14159
                                                                            switch (fi37 % ((fi29) | 1)) {
                                                                            case 4294967295:
                                                                                if (0 > 42) {
                                                                                    fmt.Println(v55, v43, v53)
                                                                                    fmt.Println(v4, v1, v26)
                                                                                }
                                                                                v1 = ((v26 << (((v53 % ((v53) | 1))) & 63)) <= (v26 / ((65535) | 1)))
                                                                            case 549:
                                                                                fmt.Println(v4, v4, v28)
                                                                            case -1:
                                                                                if len(v3) > 0 {
                                                                                    v3[0] = (v41 | 2147483647)
                                                                                }
                                                                            default:
                                                                                fmt.Println(v28)
                                                                            }
                                                                            _ = v55
                                                                        }
                                                                    default:
                                                                        var v56 int = fi37
                                                                        _ = v56
                                                                    }
                                                                }
                                                                _ = v53
                                                            }
                                                        case -459:
                                                            fmt.Println(v28, v27)
                                                        }
                                                    default:
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    }
                                                    v28 = "\n\t"
                                                }
                                            } else {
                                                var v57 float32 = (-2.25 / 1.0)
                                                if ((!v1) || v1) {
                                                    if (!false) {
                                                        v58 := GenS6{F0: 42, F1: (v41 << ((4294967295) & 63))}
                                                        v59 := GenS4{F0: 2147483647, F1: (v43 % (((v43 + int8(fi37))) | 1))}
                                                        _ = v58
                                                        _ = v59
                                                    }
                                                    v1 = (true && v1)
                                                } else {
                                                    v60 := GenS1{F0: 255, F1: true, F2: 42}
                                                    if true {
                                                        fmt.Println(v57, fi29, v4)
                                                        fmt.Println(fi37, v57, v41)
                                                    } else {
                                                        if v1 {
                                                            var v61 uint16 = v26
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            _ = v61
                                                        }
                                                        fmt.Println(v57)
                                                    }
                                                    var v62 int16 = v2
                                                    v63 := GenS7{F0: (-(v62)), F1: fi37}
                                                    _ = v60
                                                    _ = v62
                                                    _ = v63
                                                }
                                                switch (^(fi29)) {
                                                case 32768:
                                                    v1 = (42 != 647)
                                                }
                                                _ = v57
                                            }
                                            _ = v41
                                            _ = v43
                                        }
                                        if (!false) {
                                            if (v1 || v1) {
                                                v64 := GenS8{F0: -1, F1: 42, F2: 32768}
                                                _ = v64
                                            }
                                            if len(v3) > 0 {
                                                v3[0] = (v4 >> ((2) & 63))
                                            }
                                            v1 = (v28 >= ("" + (v28 + "foo bar")))
                                        } else {
                                            v4 -= (v4 * (^(v4)))
                                            v28 = v28
                                            if v1 {
                                                fmt.Println(v28, v4)
                                            } else {
                                                v65 := GenS5{F0: true, F1: v28, F2: v2}
                                                v28 = v28
                                                v3 = append(v3, v4)
                                                fmt.Println(fi29, v28, v27)
                                                var v66 uint64 = 4294967295
                                                _ = v65
                                                _ = v66
                                            }
                                        }
                                    }
                                } else {
                                    var v67 int8 = -64
                                    if false {
                                        var v68 float64 = 3.14159
                                        var v69 int8 = v67
                                        _ = v68
                                        _ = v69
                                    }
                                    v4 += (v4 ^ v4)
                                    _ = v67
                                }
                                fmt.Println(v26)
                                v27 = -234
                                _ = fi37
                            }
                        case "foo bar":
                            var v70 int16 = 0
                            _ = v70
                        default:
                            var v71 string = "\n\t"
                            _ = v71
                        }
                        v1 = ("foo bar" != (("" + v28) + v28))
                        v28 = v28
                    } else {
                        if (v1 && v1) {
                            switch ((v28 + v28) + ((v28 + v28) + ("" + "hello"))) {
                            case "foo bar":
                                for fi72 := 0; fi72 < 13; fi72++ {
                                    if len(v3) > 0 {
                                        v3[0] = (v4 + (^(v4)))
                                    }
                                    _ = fi72
                                }
                            default:
                                v73 := GenS4{F0: -2147483648, F1: 19}
                                _ = v73
                            }
                        } else {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            var v74 float32 = float32(float64(1e10))
                            if false {
                                switch fi29 {
                                case 370:
                                    for fi75 := 0; fi75 < 19; fi75++ {
                                        switch "\n\t" {
                                        case "a":
                                            var v76 uint32 = v4
                                            _ = v76
                                        case "\n\t":
                                            v77 := GenS4{F0: -1, F1: 22}
                                            _ = v77
                                        default:
                                            switch "" {
                                            case "hello":
                                                v27 = 788
                                            case "foo bar":
                                                v27 = 42
                                                switch v28 {
                                                case "a":
                                                    v78 := GenS4{F0: -2147483648, F1: -128}
                                                    _ = v78
                                                case "":
                                                    var v79 float32 = (v74 * (float32(((3.14159 * 0.1) + (0.1 / 1e10))) + v74))
                                                    _ = v79
                                                }
                                            default:
                                                v74 = 1e10
                                            }
                                        }
                                        fmt.Println(v74)
                                        var v80 int8 = -1
                                        _ = fi75
                                        _ = v80
                                    }
                                case 955:
                                    var v81 int32 = 0
                                    _ = v81
                                default:
                                    fmt.Println(v74)
                                }
                            } else {
                                v27 = v27
                                var v82 uint8 = 128
                                if true {
                                    v82 += (v82 << ((4294967295) & 63))
                                    v83 := GenS2{F0: (v2 * v2), F1: 253}
                                    fmt.Println(v82, v1)
                                    _ = v83
                                } else {
                                    fmt.Println(fi29, v4, v74)
                                    if v1 {
                                        var v84 uint32 = 65535
                                        _ = v84
                                    } else {
                                        v3 = append(v3, 439)
                                        for fi85 := 0; fi85 < 10; fi85++ {
                                            v86 := GenS9{F0: v4, F1: 717, F2: 18446744073709551615}
                                            var v87 uint = 121
                                            _ = fi85
                                            _ = v86
                                            _ = v87
                                        }
                                        if true {
                                            var v88 int32 = int32(v26)
                                            v89 := GenS6{F0: (v82 & 1), F1: (v4 | 2)}
                                            var v90 float64 = (1.5 - 0.1)
                                            _ = v88
                                            _ = v89
                                            _ = v90
                                        }
                                    }
                                    var v91 float32 = float32((float64(v74) + -2.25))
                                    _ = v91
                                }
                                _ = v82
                            }
                            _ = v74
                        }
                        v3 = append(v3, v4)
                        switch (fi29 * fi29) {
                        case 32767:
                            if (!((255 != 1) || true)) {
                                var v92 int = fi29
                                if (-369 != fi29) {
                                    fmt.Println(v2, v4)
                                    var v93 int8 = -1
                                    _ = v93
                                } else {
                                    if v1 {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        switch ("a" + "a") {
                                        case "":
                                            switch (((("a" + v28) + v28) + "\n\t") + (v28 + (v28 + v28))) {
                                            case "\n\t":
                                                switch (fi29 << ((931) & 63)) {
                                                case 32767:
                                                    var v94 uint8 = 14
                                                    _ = v94
                                                case 1:
                                                    v95 := GenS10{F0: 0, F1: v4}
                                                    _ = v95
                                                default:
                                                    var v96 int64 = (v27 & (v27 << ((2) & 63)))
                                                    _ = v96
                                                }
                                            }
                                        }
                                    } else {
                                        if len(v3) > 0 {
                                            v3[0] = (v4 ^ v4)
                                        }
                                    }
                                }
                                _ = v92
                            } else {
                                fmt.Println(v1, v28, fi29)
                                fmt.Println(v28, v4)
                                fmt.Println(v26, fi29, v2)
                            }
                        case -920:
                            var v97 float64 = 0.0
                            _ = v97
                        case -9223372036854775808:
                            fmt.Println(v2)
                        default:
                            if false {
                                var v98 int = fi29
                                if true {
                                    var v99 float64 = float64((float32((0.1 * 3.14159)) + 0.1))
                                    if v1 {
                                        var v100 int64 = v27
                                        var v101 string = ((v28 + v28) + v28)
                                        v27 -= v100
                                        _ = v100
                                        _ = v101
                                    }
                                    _ = v99
                                } else {
                                    v27 -= 1
                                }
                                _ = v98
                            } else {
                                for fi102 := 0; fi102 < 4; fi102++ {
                                    v103 := GenS11{F0: int8(v4), F1: 2147483647, F2: v2}
                                    var v104 int8 = -128
                                    var v105 uint32 = v4
                                    _ = fi102
                                    _ = v103
                                    _ = v104
                                    _ = v105
                                }
                            }
                        }
                    }
                default:
                    for fi106 := 0; fi106 < 8; fi106++ {
                        var v107 int32 = 42
                        _ = fi106
                        _ = v107
                    }
                }
            }
            v108 := GenS12{F0: v27, F1: 65535, F2: -836}
            _ = v108
        }
        var v109 int = fi29
        _ = fi29
        _ = v109
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
