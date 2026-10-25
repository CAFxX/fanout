package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 int64; F1 uint64; F2 int32 }
type GenS3 struct { F0 bool; F1 int; F2 uint8 }
type GenS4 struct { F0 int16; F1 uint64 }
type GenS5 struct { F0 string; F1 int16; F2 int8 }
type GenS6 struct { F0 int; F1 int8 }
type GenS7 struct { F0 uint64; F1 int; F2 int8 }
type GenS8 struct { F0 uint; F1 int16; F2 int64 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 string }
type GenS11 struct { F0 int64; F1 uint8 }
type GenS12 struct { F0 uint16; F1 uint; F2 int32 }
type GenS13 struct { F0 string }
type GenS14 struct { F0 bool; F1 uint }
type GenS15 struct { F0 int8; F1 int64; F2 string }
type GenS16 struct { F0 int16 }
type GenS17 struct { F0 int32; F1 uint8; F2 bool }
type GenS18 struct { F0 uint64; F1 uint32 }
func genFn0(p0 int32, p1 uint16) int {
    v6 := GenS1{F0: -2147483648, F1: 128}
    for fi7 := 0; fi7 < 9; fi7++ {
        v8 := []int{42}
        if true {
            for fi9 := 0; fi9 < 17; fi9++ {
                if (1e10 <= -2.25) {
                    p1 -= 32767
                } else {
                    fmt.Println(p0, fi9)
                    var v10 float32 = float32(((1.5 + (((((0.0 / 0.1) - 0.1) - -2.25) - 3.14159) / 1.0)) - ((1e10 * 0.0) * (0.1 / (1e10 * 1.5)))))
                    _ = v10
                }
                _ = fi9
            }
            v8 = append(v8, (fi7 % ((fi7) | 1)))
        } else {
            p0 = (p0 * p0)
        }
        _ = fi7
        _ = v8
    }
    fmt.Println(p1, p1)
    for fi11 := 0; fi11 < 16; fi11++ {
        fmt.Println(p1, p0)
        _ = fi11
    }
    v12 := []int{-535, 2147483647}
    _ = p0
    _ = p1
    _ = v12
    _ = v6
    return int(p1)
}
func genFn1(p0 uint16, p1 string, p2 uint64) uint32 {
    for fi13 := 0; fi13 < 12; fi13++ {
        for fi14 := 0; fi14 < 6; fi14++ {
            if (fi14 < -550) {
                if (1e10 < (0.1 * 1.5)) {
                    v15 := []int{32767, int(p0), fi13, fi14}
                    var v16 string = ""
                    var v17 uint64 = p2
                    _ = v15
                    _ = v16
                    _ = v17
                }
                p1 = "\n\t"
            }
            p1 = ((p1 + (p1 + p1)) + p1)
            _ = fi14
        }
        _ = fi13
    }
    for fi18 := 0; fi18 < 15; fi18++ {
        v19 := []int{(fi18 | genFn0(42, p0)), genFn0(42, (p0 / (((p0 + (p0 / (((p0 - (p0 ^ (p0 >> ((2) & 63))))) | 1)))) | 1))), (fi18 + (fi18 % ((fi18) | 1)))}
        for fi20 := 0; fi20 < 14; fi20++ {
            switch (fi20 + (fi20 / (((fi18 | (^(fi18)))) | 1))) {
            case 192:
                if true {
                    switch fi20 {
                    case 9223372036854775807:
                        fmt.Println(p0, p2, fi20)
                    case 32767:
                        if (73 == 75) {
                            p2 += uint64(p0)
                            fmt.Println(fi18, fi20, fi18)
                            p0 = 2
                        }
                    case 589:
                        p1 = "hello"
                    default:
                        var v21 uint = 1
                        _ = v21
                    }
                    p0 += p0
                }
            case 32767:
                v22 := GenS2{F0: -1, F1: 42, F2: -446}
                _ = v22
            default:
                if ((false && (!(!(((3.14159 * 1.5) - 0.1) != 0.1)))) || (false && (!((3.14159 / 3.14159) > (-2.25 * ((-2.25 * 0.0) - (-2.25 + 1.5))))))) {
                    var v23 float32 = ((0.1 + float32((-2.25 * ((1e10 * (-2.25 * 1e10)) + ((1.5 - 3.14159) + 3.14159))))) - ((0.0 * (float32(float64((0.1 * 1e10))) / (((1.5 - 0.1) * 1e10) / 1.0))) + 0.1))
                    fmt.Println(fi18)
                    if false {
                        var v24 float64 = (1e10 - 1e10)
                        if false {
                            fmt.Println(fi20, v23)
                            p1 = "hello"
                        } else {
                            var v25 float64 = v24
                            var v26 string = "foo bar"
                            if len(v19) >= 2 {
                                v19 = v19[1:len(v19)]
                            }
                            _ = v25
                            _ = v26
                        }
                        var v27 uint16 = p0
                        _ = v24
                        _ = v27
                    } else {
                        var v28 int8 = 42
                        _ = v28
                    }
                    _ = v23
                }
            }
            v29 := GenS3{F0: true, F1: (fi18 - (fi20 >> ((uint(p0)) & 63))), F2: 255}
            _ = fi20
            _ = v29
        }
        _ = fi18
        _ = v19
    }
    v30 := []int{-836}
    _ = len(v30)
    _ = p0
    _ = p1
    _ = p2
    _ = v30
    return 152
}

func main() {
    var v1 bool = false
    var v2 uint64 = 2
    var v3 float32 = -2.25
    var v4 int16 = -814
    var v5 int8 = 42
    for fi31 := 0; fi31 < 18; fi31++ {
        v2 += v2
        v3 = v3
        var v32 uint = 110
        _ = fi31
        _ = v32
    }
    v33 := GenS3{F0: true, F1: 4294967295, F2: 1}
    v1 = (v3 > float32(((((((1e10 - float64(v3)) + -2.25) / 3.14159) / 1.0) / -2.25) / 3.14159)))
    fmt.Println(v4, v2)
    var v34 uint = 65535
    v35 := []int{}
    switch "foo bar" {
    case "hello":
        if (!false) {
            v4 += (v4 % ((v4) | 1))
            if true {
                fmt.Println(v4, v34, v4)
                var v36 uint8 = 213
                _ = v36
            }
        } else {
            if ((false || v1) && (true || (-252 < 42))) {
                var v37 int64 = int64(v34)
                v38 := GenS4{F0: int16(v5), F1: v2}
                _ = v37
                _ = v38
            }
            fmt.Println(v2, v2)
            v39 := GenS3{F0: v1, F1: -486, F2: 35}
            _ = v39
        }
    }
    var v40 float64 = 3.14159
    var v41 uint = 18446744073709551615
    v42 := GenS5{F0: ("" + "hello"), F1: v4, F2: (v5 / (((^(v5))) | 1))}
    if ((v5 + v5) <= v5) {
        for fi43 := 0; fi43 < 16; fi43++ {
            fmt.Println(v34)
            if false {
                var v44 string = "a"
                var v45 int8 = 127
                var v46 uint16 = 1
                _ = v44
                _ = v45
                _ = v46
            }
            _ = fi43
        }
        fmt.Println(v5, v41)
        v47 := GenS2{F0: 9223372036854775807, F1: v2, F2: 1}
        _ = v47
    }
    v5 = (v5 + v5)
    if (2147483647 > 2) {
        _ = len(v35)
        switch "hello" {
        case "foo bar":
            if (127 > -2147483648) {
                var v48 int8 = (v5 & v5)
                v3 -= (v3 / v3)
                v49 := GenS6{F0: 9223372036854775807, F1: (v48 % ((v48) | 1))}
                _ = v48
                _ = v49
            }
        case "a":
            var v50 uint64 = 474
            _ = v50
        default:
            v35 = append(v35, 817)
        }
    }
    if false {
        var v51 int32 = -132
        v41 = (^(v34))
        _ = v51
    } else {
        if v1 {
            v41 = (v41 * v34)
            for fi52 := 0; fi52 < 18; fi52++ {
                v53 := GenS6{F0: fi52, F1: v5}
                var v54 bool = (128 == 98)
                switch (((("a" + "") + "hello") + ("hello" + "hello")) + "\n\t") {
                case "\n\t":
                    v54 = ((-9223372036854775808 > -9223372036854775808) || (v5 > (-(v5))))
                case "a":
                    _ = len(v35)
                default:
                    if v54 {
                        for fi55 := 0; fi55 < 17; fi55++ {
                            v2 = (v2 + v2)
                            _ = len(v35)
                            v56 := GenS7{F0: (v2 & 223), F1: genFn0(-2147483648, 340), F2: 17}
                            _ = fi55
                            _ = v56
                        }
                        v5 = (v5 + (v5 >> ((v34) & 63)))
                        v57 := GenS8{F0: 4294967295, F1: 1, F2: -1}
                        _ = v57
                    } else {
                        var v58 uint64 = (^(v2))
                        switch ("foo bar" + "foo bar") {
                        case "\n\t":
                            if len(v35) >= 2 {
                                v35 = v35[1:len(v35)]
                            }
                        }
                        if (genFn0(-1, 255) <= (^(fi52))) {
                            switch "\n\t" {
                            case "a":
                                for fi59 := 0; fi59 < 10; fi59++ {
                                    if ((false && false) || true) {
                                        fmt.Println(v4, fi59, v4)
                                        v60 := GenS4{F0: -594, F1: (v2 / ((v2) | 1))}
                                        _ = v60
                                    }
                                    _ = fi59
                                }
                            case "\n\t":
                                v3 = float32(1.5)
                            case "":
                                v40 += v40
                            }
                            switch "hello" {
                            case "foo bar":
                                fmt.Println(v40, v1)
                            case "":
                                for fi61 := 0; fi61 < 10; fi61++ {
                                    var v62 int8 = int8(fi52)
                                    v5 += int8(v2)
                                    v3 = ((v3 / ((float32(float64((1.5 * 1.5))) - 1.5) / v3)) - v3)
                                    _ = fi61
                                    _ = v62
                                }
                            default:
                                v2 -= (v2 - (v58 | (^(v2))))
                            }
                        } else {
                            v40 = ((1.5 / 3.14159) - v40)
                            fmt.Println(fi52)
                        }
                        _ = v58
                    }
                }
                _ = fi52
                _ = v53
                _ = v54
            }
            for fi63 := 0; fi63 < 11; fi63++ {
                v64 := GenS7{F0: (^(v2)), F1: (fi63 << ((v34) & 63)), F2: int8(v34)}
                _ = fi63
                _ = v64
            }
        } else {
            for fi65 := 0; fi65 < 4; fi65++ {
                fmt.Println(v5, v4)
                fmt.Println(v34, v2, v3)
                _ = fi65
            }
            if v1 {
                for fi66 := 0; fi66 < 14; fi66++ {
                    v2 += 0
                    _ = fi66
                }
                fmt.Println(v40, v4)
                if len(v35) >= 2 {
                    v35 = v35[1:len(v35)]
                }
            } else {
                for fi67 := 0; fi67 < 14; fi67++ {
                    v2 += v2
                    for fi68 := 0; fi68 < 7; fi68++ {
                        switch fi68 {
                        case -1:
                            var v69 int16 = 32767
                            _ = v69
                        case 4294967295:
                            var v70 int64 = 42
                            _ = v70
                        }
                        if len(v35) > 0 {
                            v35[0] = fi68
                        }
                        switch "" {
                        case "foo bar":
                            fmt.Println(v3, fi67)
                        case "":
                            var v71 uint64 = v2
                            _ = v71
                        case "a":
                            if v1 {
                                var v72 int = (fi68 - 128)
                                fmt.Println(v72, v4, v4)
                                _ = v72
                            } else {
                                v4 += v4
                                v41 = v41
                                if false {
                                    if (v2 == 300) {
                                        var v73 int8 = (-(v5))
                                        _ = v73
                                    } else {
                                        var v74 int32 = -2147483648
                                        _ = v74
                                    }
                                    if v1 {
                                        v1 = (321 < 65535)
                                    } else {
                                        v4 += (v4 >> (((^(v41))) & 63))
                                        switch (fi67 / (((fi68 << (((^(v41))) & 63))) | 1)) {
                                        case 2147483647:
                                            _ = len(v35)
                                            switch (fi67 >> ((v34) & 63)) {
                                            case 191:
                                                if false {
                                                    v40 = 1e10
                                                    var v75 int = (fi67 >> ((18) & 63))
                                                    _ = v75
                                                } else {
                                                    switch "a" {
                                                    case "foo bar":
                                                        switch -9223372036854775808 {
                                                        case 42:
                                                            if len(v35) >= 2 {
                                                                v35 = v35[1:len(v35)]
                                                            }
                                                        }
                                                        if (v2 < v2) {
                                                            var v76 uint16 = 670
                                                            v77 := GenS9{F0: ""}
                                                            _ = v76
                                                            _ = v77
                                                        } else {
                                                            fmt.Println(v1, v3, v3)
                                                            v34 = (v34 << ((v34) & 63))
                                                            var v78 int16 = -166
                                                            _ = v78
                                                        }
                                                        var v79 int = fi67
                                                        _ = v79
                                                    default:
                                                        v80 := GenS10{F0: "a"}
                                                        _ = v80
                                                    }
                                                    fmt.Println(fi67, fi67, v40)
                                                    fmt.Println(v1, v3)
                                                }
                                            default:
                                                fmt.Println(v2)
                                            }
                                            var v81 uint16 = 255
                                            if true {
                                                if (9223372036854775807 != 9223372036854775807) {
                                                    v1 = false
                                                } else {
                                                    if v1 {
                                                        var v82 int16 = (-(v4))
                                                        switch "\n\t" {
                                                        case "":
                                                            v82 += (v82 / (((v4 ^ v4)) | 1))
                                                        case "hello":
                                                            _ = len(v35)
                                                        case "foo bar":
                                                            v83 := GenS5{F0: "a", F1: (v82 + v4), F2: v5}
                                                            _ = v83
                                                        }
                                                        _ = len(v35)
                                                        _ = v82
                                                    } else {
                                                        if v1 {
                                                            _ = len(v35)
                                                            var v84 uint64 = (v2 << ((55) & 63))
                                                            _ = v84
                                                        } else {
                                                            switch "" {
                                                            case "":
                                                                v85 := GenS11{F0: int64(v81), F1: 42}
                                                                _ = v85
                                                            }
                                                        }
                                                    }
                                                    v4 = int16(v5)
                                                    var v86 int32 = 42
                                                    _ = v86
                                                }
                                                v87 := GenS12{F0: (v81 ^ 867), F1: 1, F2: int32(v34)}
                                                v35 = append(v35, (fi68 % ((269) | 1)))
                                                _ = v87
                                            } else {
                                                v88 := GenS13{F0: ("hello" + "foo bar")}
                                                v81 = v81
                                                switch (fi68 / ((fi67) | 1)) {
                                                case 0:
                                                    v2 += v2
                                                case 128:
                                                    var v89 int8 = v5
                                                    _ = v89
                                                case -2147483648:
                                                    fmt.Println(v5, v1, fi68)
                                                    v34 = v41
                                                default:
                                                    v3 += v3
                                                }
                                                v90 := GenS14{F0: (!(-32768 > (^(v4)))), F1: 440}
                                                _ = v88
                                                _ = v90
                                            }
                                            fmt.Println(v5, v41, v34)
                                            _ = v81
                                        }
                                        var v91 float64 = v40
                                        _ = v91
                                    }
                                    var v92 int8 = (v5 ^ v5)
                                    _ = v92
                                }
                            }
                        default:
                            var v93 uint64 = v2
                            _ = v93
                        }
                        _ = fi68
                    }
                    switch (fi67 << ((v34) & 63)) {
                    case -339:
                        for fi94 := 0; fi94 < 10; fi94++ {
                            if (v40 < 1.5) {
                                switch ("hello" + (("hello" + (("a" + "foo bar") + "a")) + "a")) {
                                case "a":
                                    var v95 float64 = (v40 - (((-2.25 / 0.1) - v40) / v40))
                                    _ = v95
                                case "foo bar":
                                    switch fi67 {
                                    case 32768:
                                        if len(v35) > 0 {
                                            v35[0] = -2147483648
                                        }
                                    case -861:
                                        v4 += v4
                                    default:
                                        var v96 int64 = 9223372036854775807
                                        _ = v96
                                    }
                                case "\n\t":
                                    var v97 uint32 = 4294967295
                                    _ = v97
                                }
                                _ = len(v35)
                            } else {
                                fmt.Println(v3, v1, v3)
                                if (genFn0(-1, 0) <= (fi94 / (((fi67 ^ 32768)) | 1))) {
                                    var v98 int8 = -1
                                    v99 := GenS9{F0: (("a" + "a") + "")}
                                    _ = v98
                                    _ = v99
                                }
                                var v100 uint = v34
                                _ = v100
                            }
                            v101 := GenS14{F0: v1, F1: (v41 << ((v41) & 63))}
                            _ = fi94
                            _ = v101
                        }
                    }
                    _ = fi67
                }
                if v1 {
                    if ((!v1) || false) {
                        for fi102 := 0; fi102 < 17; fi102++ {
                            v5 = (v5 - v5)
                            _ = fi102
                        }
                        switch "" {
                        case "":
                            var v103 int32 = 0
                            _ = v103
                        case "\n\t":
                            for fi104 := 0; fi104 < 20; fi104++ {
                                v5 -= v5
                                _ = fi104
                            }
                        case "a":
                            v35 = append(v35, -573)
                        }
                        var v105 int64 = -9223372036854775808
                        _ = v105
                    } else {
                        if len(v35) > 0 {
                            v35[0] = genFn0(-1, 255)
                        }
                        _ = len(v35)
                        if false {
                            v106 := GenS15{F0: v5, F1: 0, F2: ("foo bar" + "a")}
                            _ = v106
                        } else {
                            var v107 bool = (true && v1)
                            _ = v107
                        }
                    }
                    var v108 bool = v1
                    _ = v108
                }
            }
            switch "hello" {
            case "a":
                if (((float64(((0.0 + float32((3.14159 / v40))) + (v3 + 0.0))) * v40) >= 3.14159) && (v1 && v1)) {
                    fmt.Println(v4, v3, v3)
                    if v1 {
                        if len(v35) >= 2 {
                            v35 = v35[1:len(v35)]
                        }
                        v35 = append(v35, genFn0(491, 671))
                    } else {
                        if (v1 || v1) {
                            var v109 float64 = v40
                            _ = v109
                        }
                        if len(v35) >= 2 {
                            v35 = v35[1:len(v35)]
                        }
                    }
                } else {
                    if ((128 <= 128) && (genFn0(42, 0) == -9223372036854775808)) {
                        if (v40 > v40) {
                            if (0 < 2) {
                                fmt.Println(v2, v3, v40)
                                if len(v35) >= 2 {
                                    v35 = v35[1:len(v35)]
                                }
                                fmt.Println(v4)
                            } else {
                                for fi110 := 0; fi110 < 8; fi110++ {
                                    var v111 uint8 = 217
                                    if v1 {
                                        switch fi110 {
                                        case 128:
                                            switch ("" + "\n\t") {
                                            case "hello":
                                                if (v111 > 2) {
                                                    if len(v35) >= 2 {
                                                        v35 = v35[1:len(v35)]
                                                    }
                                                    var v112 int8 = -60
                                                    fmt.Println(v40, v5)
                                                    _ = v112
                                                }
                                            case "":
                                                var v113 int = fi110
                                                _ = v113
                                            default:
                                                fmt.Println(v5, v1)
                                            }
                                        default:
                                            v2 -= (v2 & 295)
                                        }
                                        var v114 uint64 = v2
                                        _ = v114
                                    }
                                    v34 = (v41 % ((42) | 1))
                                    _ = fi110
                                    _ = v111
                                }
                                v115 := GenS16{F0: 42}
                                switch "hello" {
                                case "foo bar":
                                    var v116 int = -944
                                    _ = v116
                                case "\n\t":
                                    var v117 int16 = (v4 << ((255) & 63))
                                    _ = v117
                                default:
                                    if len(v35) >= 2 {
                                        v35 = v35[1:len(v35)]
                                    }
                                }
                                _ = v115
                            }
                        } else {
                            var v118 string = "\n\t"
                            if (441 != 127) {
                                for fi119 := 0; fi119 < 13; fi119++ {
                                    if v1 {
                                        for fi120 := 0; fi120 < 4; fi120++ {
                                            var v121 float32 = float32(v40)
                                            var v122 bool = (!true)
                                            v123 := GenS17{F0: 42, F1: 231, F2: false}
                                            _ = fi120
                                            _ = v121
                                            _ = v122
                                            _ = v123
                                        }
                                    } else {
                                        v124 := GenS3{F0: (!v1), F1: -1, F2: 42}
                                        _ = v124
                                    }
                                    _ = fi119
                                }
                            }
                            v125 := GenS15{F0: 26, F1: 233, F2: (v118 + v118)}
                            _ = v118
                            _ = v125
                        }
                    } else {
                        fmt.Println(v3)
                    }
                    v41 = 4294967295
                }
            }
        }
    }
    v40 = v40
    for fi126 := 0; fi126 < 19; fi126++ {
        for fi127 := 0; fi127 < 14; fi127++ {
            fmt.Println(v1, v41)
            v128 := GenS18{F0: 705, F1: 353}
            _ = fi127
            _ = v128
        }
        var v129 uint8 = 255
        _ = fi126
        _ = v129
    }
    fmt.Println(v34)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
}
