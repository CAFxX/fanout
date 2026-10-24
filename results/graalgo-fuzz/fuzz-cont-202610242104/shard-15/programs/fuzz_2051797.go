package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint32; F1 int64; F2 int16 }
type GenS5 struct { F0 int64; F1 int16; F2 int32 }
type GenS6 struct { F0 int64; F1 int32; F2 int8 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 int8; F1 string; F2 int16 }
func genFn0(p0 uint64, p1 uint32, p2 uint32) int32 {
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 uint32 = p1
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: -2147483648}
    for fi10 := 0; fi10 < 20; fi10++ {
        v11 := GenS1{F0: fi10}
        switch "foo bar" {
        case "\n\t":
            v12 := GenS2{F0: 42, F1: 32767}
            _ = v12
        }
        _ = fi10
        _ = v11
    }
    v13 := []int{}
    switch "hello" {
    case "hello":
        fmt.Println(p1)
    case "foo bar":
        var v14 uint = 65535
        _ = v14
    case "":
        if false {
            fmt.Println(p0, p0)
            p0 = p0
            p0 -= (p0 * p0)
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v9
    return 556
}
func genFn1(p0 bool) int64 {
    v15 := []int{-1}
    p0 = (("" > "foo bar") || p0)
    var v16 int = 9223372036854775807
    fmt.Println(v16, v16)
    fmt.Println(p0, v16)
    _ = p0
    _ = v15
    _ = v16
    return -9223372036854775808
}

func main() {
    v1 := []int{32767, -2147483648, 128, -1}
    var v2 int8 = 127
    var v3 int8 = (v2 + (v2 - v2))
    var v4 int = -99
    var v5 int8 = -81
    v6 := []int{(v4 / ((v4) | 1)), 736, (v4 & (v4 ^ v4)), (v4 & -236)}
    _ = len(v6)
    for fi17 := 0; fi17 < 8; fi17++ {
        var v18 bool = (2 >= 603)
        for fi19 := 0; fi19 < 13; fi19++ {
            fmt.Println(fi17, v5, fi17)
            _ = fi19
        }
        _ = fi17
        _ = v18
    }
    switch "a" {
    case "foo bar":
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    default:
        var v20 float32 = (0.0 - float32(0.1))
        _ = v20
    }
    fmt.Println(v4, v5)
    var v21 float32 = (0.1 / -2.25)
    fmt.Println(v2)
    if false {
        switch v4 {
        case -1:
            var v22 uint64 = 420
            _ = v22
        case 57:
            var v23 uint = 4294967295
            _ = v23
        case 9223372036854775807:
            var v24 float64 = 1e10
            _ = v24
        default:
            v25 := GenS1{F0: -549}
            _ = v25
        }
    }
    if true {
        if len(v1) > 0 {
            v1[0] = int(v3)
        }
    } else {
        switch (v4 | (v4 ^ 2147483647)) {
        case -1:
            v3 -= (-(v3))
        case 42:
            for fi26 := 0; fi26 < 1; fi26++ {
                if ("foo bar" >= "a") {
                    if len(v6) > 0 {
                        v6[0] = (^(v4))
                    }
                    var v27 string = "a"
                    v28 := GenS2{F0: 42, F1: 0}
                    _ = v27
                    _ = v28
                }
                if true {
                    v4 += fi26
                }
                for fi29 := 0; fi29 < 10; fi29++ {
                    fmt.Println(v3, v4, v5)
                    v30 := GenS2{F0: 600, F1: -1}
                    _ = fi29
                    _ = v30
                }
                _ = fi26
            }
        case -9223372036854775808:
            for fi31 := 0; fi31 < 9; fi31++ {
                if false {
                    var v32 float32 = v21
                    _ = v32
                }
                v21 += (((v21 * (((-2.25 - (((v21 + -2.25) - v21) - v21)) * v21) - v21)) * 1.5) * v21)
                switch v4 {
                case 2147483647:
                    v33 := GenS3{F0: 831}
                    _ = v33
                case 1:
                    fmt.Println(fi31)
                case 4294967295:
                    v1 = append(v1, fi31)
                }
                _ = fi31
            }
        default:
            _ = len(v6)
        }
    }
    v34 := GenS2{F0: 2, F1: -1}
    for fi35 := 0; fi35 < 6; fi35++ {
        var v36 uint = 65535
        if true {
            _ = len(v1)
        }
        if len(v6) > 0 {
            v6[0] = v4
        }
        _ = fi35
        _ = v36
    }
    if false {
        v2 = (v2 << ((1) & 63))
        v2 += v2
    }
    if false {
        if (false || false) {
            if len(v1) > 0 {
                v1[0] = int(v5)
            }
            v6 = append(v6, (v4 << ((363) & 63)))
        }
        var v37 int8 = 42
        _ = len(v1)
        _ = v37
    }
    fmt.Println(v4, v21)
    _ = len(v1)
    for fi38 := 0; fi38 < 12; fi38++ {
        v6 = append(v6, 529)
        _ = fi38
    }
    for fi39 := 0; fi39 < 2; fi39++ {
        if (float32((0.0 - float64(v21))) > (v21 + ((v21 / v21) * v21))) {
            for fi40 := 0; fi40 < 10; fi40++ {
                v41 := GenS3{F0: v4}
                _ = fi40
                _ = v41
            }
            for fi42 := 0; fi42 < 19; fi42++ {
                fmt.Println(v4)
                fmt.Println(v3, v5, fi39)
                _ = len(v1)
                v3 -= (v3 / (((v5 / (((v2 ^ (v5 >> ((0) & 63)))) | 1))) | 1))
                _ = fi42
            }
            for fi43 := 0; fi43 < 3; fi43++ {
                if false {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    var v44 int64 = 1
                    if len(v1) > 0 {
                        v1[0] = (fi39 ^ fi43)
                    }
                    _ = v44
                } else {
                    if false {
                        if false {
                            var v45 int64 = -9223372036854775808
                            v4 = (fi39 >> ((920) & 63))
                            _ = v45
                        } else {
                            v3 -= (v5 * (v2 & 71))
                            fmt.Println(v4)
                        }
                        switch "" {
                        case "hello":
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        case "a":
                            fmt.Println(v4, fi39, v3)
                        default:
                            switch v4 {
                            case 127:
                                v5 = v3
                                var v46 int = fi43
                                var v47 uint8 = 128
                                _ = v46
                                _ = v47
                            case -805:
                                var v48 int8 = (v5 & (v5 / (((v5 + v5)) | 1)))
                                _ = v48
                            default:
                                v4 += 491
                            }
                        }
                        v49 := GenS4{F0: 380, F1: -1, F2: 1000}
                        _ = v49
                    }
                    if (673 < 131) {
                        switch (fi39 ^ (fi39 | (^(fi43)))) {
                        case -368:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        case -67:
                            if false {
                                switch (fi39 + 2147483647) {
                                case 9223372036854775807:
                                    if ((true || (4294967295 == 65535)) || (fi43 < v4)) {
                                        v50 := GenS4{F0: 1, F1: 699, F2: 1}
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        _ = v50
                                    }
                                    if false {
                                        var v51 string = "hello"
                                        _ = v51
                                    }
                                    v52 := GenS5{F0: 0, F1: 1, F2: -1}
                                    _ = v52
                                case -366:
                                    if false {
                                        var v53 string = ""
                                        _ = v53
                                    } else {
                                        var v54 bool = true
                                        if (0.1 < -2.25) {
                                            var v55 uint16 = 32767
                                            v56 := GenS3{F0: fi43}
                                            if len(v6) > 0 {
                                                v6[0] = fi39
                                            }
                                            _ = v55
                                            _ = v56
                                        } else {
                                            v4 += (^(fi39))
                                            v5 -= v3
                                        }
                                        v54 = ((fi43 * (v4 - (fi39 + v4))) < (^(v4)))
                                        _ = v54
                                    }
                                case 32768:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    switch fi43 {
                                    case -880:
                                        fmt.Println(v2, v2)
                                    case 1:
                                        switch ("foo bar" + "hello") {
                                        case "\n\t":
                                            var v57 uint16 = 631
                                            _ = v57
                                        case "":
                                            v5 = v5
                                        }
                                    }
                                    var v58 int8 = 125
                                    _ = v58
                                default:
                                    var v59 uint64 = 1
                                    _ = v59
                                }
                            }
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        var v60 uint64 = 4294967295
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        _ = v60
                    } else {
                        switch ("foo bar" + "\n\t") {
                        case "\n\t":
                            v21 = (v21 / 1e10)
                        case "":
                            if false {
                                v3 = (-(v5))
                                v21 += 3.14159
                                switch "" {
                                case "\n\t":
                                    v4 += (v4 / ((32768) | 1))
                                default:
                                    var v61 int16 = 1000
                                    _ = v61
                                }
                                v5 -= (v3 | (v5 | (v5 | -128)))
                            } else {
                                switch -740 {
                                case 4294967295:
                                    v3 -= v2
                                case 9223372036854775807:
                                    v6 = append(v6, fi43)
                                case 2147483647:
                                    if (-451 <= -2147483648) {
                                        if (!(v4 == 32768)) {
                                            if len(v1) > 0 {
                                                v1[0] = fi43
                                            }
                                            switch "a" {
                                            case "\n\t":
                                                v3 -= (v3 << ((65535) & 63))
                                            case "":
                                                var v62 int8 = 115
                                                if true {
                                                    v63 := GenS5{F0: 42, F1: 1000, F2: genFn0(18446744073709551615, 1, 0)}
                                                    var v64 bool = false
                                                    _ = v63
                                                    _ = v64
                                                }
                                                var v65 int = fi43
                                                _ = v62
                                                _ = v65
                                            case "foo bar":
                                                v2 = -1
                                            }
                                        } else {
                                            if false {
                                                if false {
                                                    var v66 float32 = float32(1e10)
                                                    _ = v66
                                                } else {
                                                    _ = len(v6)
                                                }
                                                var v67 int16 = -1
                                                _ = v67
                                            } else {
                                                v3 = v3
                                                if false {
                                                    v68 := GenS3{F0: fi43}
                                                    v2 = (-(v3))
                                                    v3 -= (v5 * v2)
                                                    var v69 uint8 = 2
                                                    _ = v68
                                                    _ = v69
                                                } else {
                                                    var v70 uint16 = 255
                                                    _ = v70
                                                }
                                                var v71 int16 = -834
                                                _ = v71
                                            }
                                            v72 := GenS6{F0: -843, F1: 1, F2: -50}
                                            v21 -= v21
                                            _ = v72
                                        }
                                        if (false || false) {
                                            var v73 int16 = -32768
                                            if false {
                                                if (true && false) {
                                                    v74 := GenS2{F0: 172, F1: v73}
                                                    _ = v74
                                                } else {
                                                    v5 += v5
                                                    if (genFn0(42, 365, 934) > 42) {
                                                        var v75 uint64 = 42
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                        case "hello":
                                                            if (!(!false)) {
                                                                if len(v1) > 0 {
                                                                    v1[0] = v4
                                                                }
                                                            }
                                                        default:
                                                            v76 := GenS4{F0: 65535, F1: 42, F2: (v73 % (((v73 / ((42) | 1))) | 1))}
                                                            _ = v76
                                                        }
                                                        switch (v4 & (-(v4))) {
                                                        case 4294967295:
                                                            var v77 int32 = 104
                                                            _ = v77
                                                        }
                                                        _ = len(v6)
                                                        _ = v75
                                                    } else {
                                                        v2 = (v2 & v5)
                                                        v5 = (^(v5))
                                                    }
                                                    _ = len(v6)
                                                }
                                                v2 += v5
                                                if true {
                                                    v21 += v21
                                                    v6 = append(v6, (fi43 << ((704) & 63)))
                                                    if (((v73 ^ (v73 + (^(v73)))) != 42) && false) {
                                                        v78 := GenS5{F0: 9223372036854775807, F1: -32768, F2: genFn0(42, 0, 0)}
                                                        _ = v78
                                                    } else {
                                                        if (42 < 0) {
                                                            v79 := GenS5{F0: 390, F1: v73, F2: 975}
                                                            var v80 int32 = -979
                                                            _ = v79
                                                            _ = v80
                                                        }
                                                        fmt.Println(fi39, v5)
                                                    }
                                                } else {
                                                    if false {
                                                        fmt.Println(v21, fi43)
                                                    }
                                                    if (0.1 != -2.25) {
                                                        v5 += (v5 / ((v3) | 1))
                                                    } else {
                                                        var v81 uint32 = 108
                                                        var v82 uint = 518
                                                        if true {
                                                            v5 = v3
                                                            v83 := GenS6{F0: 311, F1: 0, F2: 42}
                                                            if true {
                                                                var v84 uint = 0
                                                                var v85 uint64 = 125
                                                                _ = v84
                                                                _ = v85
                                                            }
                                                            _ = v83
                                                        } else {
                                                            var v86 int8 = v3
                                                            fmt.Println(v2)
                                                            if (true || false) {
                                                                v82 -= (^(v82))
                                                                if true {
                                                                    if len(v6) > 0 {
                                                                        v6[0] = v4
                                                                    }
                                                                    var v87 string = ("hello" + "a")
                                                                    v81 -= v81
                                                                    v2 = (v2 & 1)
                                                                    var v88 int8 = v2
                                                                    v6 = append(v6, (fi39 | (fi39 ^ -9223372036854775808)))
                                                                    _ = v87
                                                                    _ = v88
                                                                } else {
                                                                    fmt.Println(v2)
                                                                    _ = len(v6)
                                                                    var v89 int64 = genFn1(true)
                                                                    _ = v89
                                                                }
                                                            } else {
                                                                var v90 string = "\n\t"
                                                                if len(v6) >= 2 {
                                                                    v6 = v6[1:len(v6)]
                                                                }
                                                                fmt.Println(fi39, v82, v86)
                                                                _ = v90
                                                            }
                                                            _ = v86
                                                        }
                                                        _ = v81
                                                        _ = v82
                                                    }
                                                }
                                            } else {
                                                if (!(true && (2 == 2))) {
                                                    switch "a" {
                                                    case "":
                                                        v91 := GenS7{F0: 701}
                                                        _ = v91
                                                    default:
                                                        if false {
                                                            var v92 float32 = v21
                                                            _ = v92
                                                        }
                                                    }
                                                } else {
                                                    switch fi39 {
                                                    case -2147483648:
                                                        var v93 uint8 = 128
                                                        _ = v93
                                                    default:
                                                        if len(v1) > 0 {
                                                            v1[0] = -2147483648
                                                        }
                                                    }
                                                }
                                                v3 = (v5 & int8(fi43))
                                            }
                                            if false {
                                                v1 = append(v1, fi39)
                                                if (-2.25 != 0.0) {
                                                    v4 -= (fi43 ^ v4)
                                                    if (1 == 1) {
                                                        fmt.Println(v5)
                                                        v2 = v5
                                                        var v94 uint8 = 42
                                                        _ = v94
                                                    }
                                                    var v95 float32 = float32(0.0)
                                                    _ = v95
                                                }
                                                switch "" {
                                                case "a":
                                                    v96 := GenS8{F0: 2}
                                                    _ = v96
                                                case "":
                                                    v97 := GenS1{F0: 9223372036854775807}
                                                    _ = v97
                                                case "foo bar":
                                                    if true {
                                                        v98 := GenS7{F0: 2147483647}
                                                        _ = v98
                                                    }
                                                default:
                                                    if (!((1 >= 1) || ((!false) && (!(-2147483648 < genFn0(1, 2, 0)))))) {
                                                        if (!true) {
                                                            if (0 != 42) {
                                                                v4 -= 128
                                                                fmt.Println(v73, v73, v4)
                                                                if (65535 > 4294967295) {
                                                                    _ = len(v1)
                                                                    var v99 int16 = v73
                                                                    v100 := GenS3{F0: v4}
                                                                    v6 = append(v6, fi39)
                                                                    fmt.Println(fi43, fi43)
                                                                    _ = v100
                                                                    _ = v99
                                                                } else {
                                                                    v73 = v73
                                                                    v2 -= v5
                                                                }
                                                            }
                                                        } else {
                                                            var v101 uint64 = 0
                                                            switch "foo bar" {
                                                            case "a":
                                                                v102 := GenS5{F0: int64(fi39), F1: 742, F2: 42}
                                                                _ = v102
                                                            case "foo bar":
                                                                if true {
                                                                    fmt.Println(fi43, fi39, v73)
                                                                } else {
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                }
                                                            }
                                                            if len(v6) > 0 {
                                                                v6[0] = fi39
                                                            }
                                                            var v103 int8 = (^(v5))
                                                            _ = v101
                                                            _ = v103
                                                        }
                                                        var v104 int8 = (v3 + v5)
                                                        _ = v104
                                                    }
                                                }
                                            } else {
                                                switch v4 {
                                                case 9223372036854775807:
                                                    if ((false || false) || (v4 != 128)) {
                                                        var v105 float32 = v21
                                                        _ = v105
                                                    }
                                                default:
                                                    fmt.Println(v21)
                                                }
                                            }
                                            _ = v73
                                        } else {
                                            switch "a" {
                                            case "a":
                                                var v106 string = ("foo bar" + "hello")
                                                _ = v106
                                            default:
                                                v107 := GenS6{F0: -1, F1: -2147483648, F2: (v5 & int8(fi43))}
                                                _ = v107
                                            }
                                        }
                                    }
                                default:
                                    v108 := GenS9{F0: 42}
                                    _ = v108
                                }
                                if false {
                                    var v109 int16 = 32767
                                    _ = v109
                                }
                                fmt.Println(v3, v3, fi39)
                            }
                        }
                        var v110 uint16 = 233
                        var v111 int16 = 950
                        var v112 uint32 = 0
                        if (true && true) {
                            v112 -= 2147483647
                        } else {
                            if len(v6) > 0 {
                                v6[0] = (^(v4))
                            }
                            if (18446744073709551615 > 4294967295) {
                                if false {
                                    var v113 uint16 = v110
                                    v111 -= 122
                                    _ = v113
                                }
                                v112 -= (v112 + (v112 >> ((710) & 63)))
                            }
                        }
                        _ = v110
                        _ = v111
                        _ = v112
                    }
                    if false {
                        var v114 bool = false
                        if v114 {
                            var v115 int8 = v5
                            fmt.Println(v4, v4, fi39)
                            fmt.Println(v3, fi43, v21)
                            _ = v115
                        } else {
                            v4 += fi43
                            if v114 {
                                var v116 string = "foo bar"
                                _ = v116
                            } else {
                                if false {
                                    switch "foo bar" {
                                    case "hello":
                                        fmt.Println(fi39, fi43)
                                    case "":
                                        v5 += v2
                                    default:
                                        _ = len(v6)
                                    }
                                    var v117 int = v4
                                    fmt.Println(v5, v114)
                                    _ = v117
                                } else {
                                    v2 += (v2 + 127)
                                    var v118 float64 = float64((float32(((1e10 / (float64(v21) + 1e10)) / 1e10)) - (float32(3.14159) + (0.0 * v21))))
                                    if v114 {
                                        switch "a" {
                                        case "hello":
                                            fmt.Println(fi43)
                                        }
                                    } else {
                                        v119 := GenS10{F0: 0}
                                        v114 = v114
                                        v114 = v114
                                        _ = v119
                                    }
                                    _ = v118
                                }
                            }
                            if (false && v114) {
                                fmt.Println(v4, v5)
                                v6 = append(v6, 127)
                            }
                            v114 = true
                            if (v114 && (!(42 >= 255))) {
                                var v120 int16 = 32767
                                _ = v120
                            }
                            v21 = v21
                        }
                        _ = v114
                    } else {
                        var v121 float64 = 1e10
                        _ = v121
                    }
                }
                _ = fi43
            }
        }
        v122 := GenS11{F0: (v5 | (v3 | -125)), F1: (("" + ("a" + "foo bar")) + "a"), F2: -32768}
        fmt.Println(v21, v5)
        _ = fi39
        _ = v122
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
