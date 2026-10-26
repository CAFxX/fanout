package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 int; F1 int64; F2 int16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint64; F1 uint }
type GenS6 struct { F0 int8; F1 uint32; F2 int8 }
type GenS7 struct { F0 uint; F1 uint; F2 int8 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 string; F1 int16 }
type GenS12 struct { F0 int32 }
func genFn0(p0 int32) int16 {
    for fi7 := 0; fi7 < 13; fi7++ {
        if ((2 <= 1) || true) {
            if ((792 > 813) && true) {
                p0 -= 2147483647
                if true {
                    for fi8 := 0; fi8 < 1; fi8++ {
                        v9 := GenS1{F0: 42}
                        _ = fi8
                        _ = v9
                    }
                    v10 := []int{int(p0), fi7, (fi7 >> ((18446744073709551615) & 63))}
                    _ = v10
                }
                var v11 float32 = float32(1e10)
                _ = v11
            }
            var v12 uint32 = 775
            if (-2147483648 > (fi7 * fi7)) {
                p0 = (-(p0))
            }
            _ = v12
        }
        v13 := []int{}
        _ = fi7
        _ = v13
    }
    p0 = (p0 - (p0 / (((^(p0))) | 1)))
    p0 = p0
    _ = p0
    return 1
}

func main() {
    var v1 int32 = 0
    var v2 int64 = -213
    var v3 uint = 255
    var v4 uint16 = 42
    v5 := make(map[int]int)
    var v6 string = "\n\t"
    v3 = v3
    var v14 uint8 = 119
    var v15 float64 = 1.5
    v16 := GenS1{F0: (v4 - 65535)}
    if (-3 == 100) {
        for fi17 := 0; fi17 < 8; fi17++ {
            v18 := []int{}
            _ = fi17
            _ = v18
        }
        if (-680 != 0) {
            var v19 uint16 = uint16(v3)
            v20 := []int{2147483647, -9223372036854775808}
            _ = v19
            _ = v20
        }
    } else {
        if (v2 < v2) {
            switch v6 {
            case "":
                var v21 bool = true
                _ = v21
            }
        }
        if ((uint8(v3) != v14) && (!(978 != 42))) {
            v3 = (^(v3))
        } else {
            v3 = 710
            v22 := GenS1{F0: v4}
            _ = v22
        }
    }
    _ = len(v5)
    if true {
        if ((v14 << ((v3) & 63)) > (v14 + (v14 << ((v3) & 63)))) {
            fmt.Println(v1, v6)
        }
        for fi23 := 0; fi23 < 3; fi23++ {
            var v24 bool = (v3 <= v3)
            if v24 {
                fmt.Println(v6, v2, v24)
                var v25 float64 = v15
                _ = v25
            }
            var v26 uint8 = (v14 * 0)
            _ = fi23
            _ = v24
            _ = v26
        }
        var v27 uint64 = uint64(v3)
        _ = v27
    }
    if (v3 != v3) {
        var v28 float64 = (v15 - v15)
        _ = v28
    }
    var v29 bool = (!true)
    v15 = ((((v15 / float64(float32(v15))) + 0.0) - 0.1) + 0.1)
    for fi30 := 0; fi30 < 3; fi30++ {
        fmt.Println(v4)
        v15 = 1.5
        _ = fi30
    }
    var v31 uint64 = 18446744073709551615
    for fi32 := 0; fi32 < 2; fi32++ {
        v4 -= v4
        _ = fi32
    }
    if v29 {
        for fi33 := 0; fi33 < 20; fi33++ {
            v34 := v5[(fi33 >> ((18446744073709551615) & 63))]
            for fi35 := 0; fi35 < 13; fi35++ {
                v2 += v2
                var v36 uint16 = v4
                if ((v31 * v31) >= 2) {
                    if ((v2 >> ((18446744073709551615) & 63)) < v2) {
                        v37 := GenS1{F0: 42}
                        delete(v5, (fi33 - (fi35 & v34)))
                        _ = v37
                    } else {
                        v34 += fi33
                        fmt.Println(v31)
                        v36 = 32767
                    }
                }
                _ = fi35
                _ = v36
            }
            if false {
                v3 -= v3
                fmt.Println(v6, v15)
                for fi38 := 0; fi38 < 18; fi38++ {
                    fmt.Println(v6, v2, fi38)
                    fmt.Println(v1, v4)
                    _ = fi38
                }
            } else {
                v2 -= v2
                switch int(v31) {
                case -530:
                    v5[89] = v34
                default:
                    switch v6 {
                    case "hello":
                        for fi39 := 0; fi39 < 1; fi39++ {
                            fmt.Println(fi39, v4, v1)
                            var v40 string = "foo bar"
                            v15 += (((v15 - 1e10) + v15) * v15)
                            var v41 uint = 42
                            _ = fi39
                            _ = v40
                            _ = v41
                        }
                    case "a":
                        v34 = (v34 + -9223372036854775808)
                    case "":
                        v42 := GenS1{F0: v4}
                        _ = v42
                    default:
                        var v43 uint32 = 4294967295
                        _ = v43
                    }
                }
            }
            _ = fi33
            _ = v34
        }
        v1 = int32(v14)
    } else {
        v44 := GenS2{F0: 759, F1: (^(v31))}
        _ = v44
    }
    for fi45 := 0; fi45 < 19; fi45++ {
        fmt.Println(fi45, v6)
        _ = fi45
    }
    v6 = (v6 + ((v6 + v6) + v6))
    if (830 == v2) {
        if false {
            v46 := GenS1{F0: 42}
            var v47 int = 876
            _ = v46
            _ = v47
        } else {
            var v48 int16 = genFn0(v1)
            _ = v48
        }
        var v49 float32 = (0.1 * float32(1e10))
        _ = v49
    } else {
        if v29 {
            v31 = v31
            switch ("\n\t" + "foo bar") {
            case "hello":
                v50 := v5[-335]
                _ = v50
            case "\n\t":
                v51 := GenS3{F0: -9223372036854775808, F1: 1, F2: 32767}
                _ = v51
            }
            fmt.Println(v4, v6, v31)
        } else {
            if ((^(v14)) >= v14) {
                var v52 int32 = v1
                for fi53 := 0; fi53 < 19; fi53++ {
                    if v29 {
                        for fi54 := 0; fi54 < 7; fi54++ {
                            switch v6 {
                            case "hello":
                                v52 -= (v52 ^ v1)
                            default:
                                v55 := GenS2{F0: 1000, F1: v31}
                                fmt.Println(fi54, v31)
                                v52 -= int32(v2)
                                _ = v55
                            }
                            _ = fi54
                        }
                        fmt.Println(v31, v15)
                    } else {
                        for fi56 := 0; fi56 < 19; fi56++ {
                            v57 := GenS2{F0: 191, F1: v31}
                            _ = fi56
                            _ = v57
                        }
                    }
                    if false {
                        if (v15 != v15) {
                            if v29 {
                                for fi58 := 0; fi58 < 16; fi58++ {
                                    fmt.Println(v29, v6)
                                    _ = fi58
                                }
                            }
                            if v29 {
                                var v59 uint32 = 2
                                v1 += v52
                                _ = v59
                            }
                        } else {
                            switch fi53 {
                            case 9223372036854775807:
                                var v60 uint32 = 42
                                _ = v60
                            case 0:
                                v61 := []int{}
                                _ = v61
                            default:
                                v52 += v1
                            }
                            for fi62 := 0; fi62 < 2; fi62++ {
                                fmt.Println(v29, fi62)
                                v2 -= v2
                                _ = fi62
                            }
                            switch fi53 {
                            case -2147483648:
                                v63 := GenS2{F0: -32768, F1: v31}
                                _ = v63
                            case 32767:
                                switch "hello" {
                                case "":
                                    fmt.Println(v31, fi53, v15)
                                default:
                                    fmt.Println(v31, v15, v14)
                                }
                            case 97:
                                fmt.Println(v31, v52, v2)
                            }
                        }
                        if true {
                            if ((v4 + uint16(v3)) <= v4) {
                                v31 -= 1
                                switch "a" {
                                case "hello":
                                    v52 += (v52 & v52)
                                default:
                                    v29 = v29
                                }
                                var v64 float32 = 1.5
                                _ = v64
                            }
                            v6 = "\n\t"
                            v31 -= (^(v31))
                        } else {
                            fmt.Println(v14, v31)
                            switch fi53 {
                            case 32767:
                                switch fi53 {
                                case 1:
                                    for fi65 := 0; fi65 < 20; fi65++ {
                                        switch (fi65 | fi53) {
                                        case 903:
                                            if false {
                                                v66 := GenS1{F0: 1}
                                                var v67 bool = (v14 == 128)
                                                v68 := GenS3{F0: (fi53 | 127), F1: -1, F2: 1000}
                                                v1 = int32(v2)
                                                _ = v66
                                                _ = v67
                                                _ = v68
                                            }
                                            v1 = v52
                                            var v69 int16 = 220
                                            _ = v69
                                        case 32768:
                                            var v70 uint16 = (v4 >> ((v3) & 63))
                                            _ = v70
                                        case 2147483647:
                                            switch v6 {
                                            case "\n\t":
                                                if ((!(456 <= 2)) || true) {
                                                    v52 = v1
                                                }
                                            case "foo bar":
                                                fmt.Println(v15, v52, v29)
                                            }
                                        }
                                        if (!true) {
                                            var v71 uint16 = (v4 ^ v4)
                                            if (238 <= (^(v14))) {
                                                v72 := GenS4{F0: (v71 & (v4 - 42))}
                                                switch (fi65 | (fi65 >> (((v3 << ((v3) & 63))) & 63))) {
                                                case -2147483648:
                                                    var v73 string = v6
                                                    _ = v73
                                                case -9223372036854775808:
                                                    if v29 {
                                                        var v74 int = 1
                                                        var v75 int32 = (^(v1))
                                                        if (v29 || v29) {
                                                            delete(v5, fi65)
                                                            if false {
                                                                v76 := []int{}
                                                                _ = v76
                                                            } else {
                                                                var v77 int = (fi65 | -1)
                                                                _ = v77
                                                            }
                                                            switch (v6 + v6) {
                                                            case "":
                                                                v2 -= (v2 ^ (v2 & v2))
                                                            case "a":
                                                                v4 += 0
                                                            case "\n\t":
                                                                fmt.Println(v15)
                                                            default:
                                                                v78 := []int{(-(fi65))}
                                                                _ = v78
                                                            }
                                                            var v79 string = v6
                                                            _ = v79
                                                        }
                                                        _ = v74
                                                        _ = v75
                                                    }
                                                    v80 := GenS3{F0: 1, F1: (^(v2)), F2: -27}
                                                    _ = v80
                                                }
                                                _ = v72
                                            }
                                            var v81 int8 = -38
                                            _ = v71
                                            _ = v81
                                        }
                                        _ = fi65
                                    }
                                case 483:
                                    var v82 int = 201
                                    _ = v82
                                default:
                                    if ((v14 == v14) && v29) {
                                        fmt.Println(v15)
                                    }
                                }
                            case 581:
                                v83 := GenS1{F0: 32767}
                                _ = v83
                            case -837:
                                v84 := []int{(^(fi53)), (fi53 % (((fi53 / (((fi53 - (fi53 >> ((834) & 63)))) | 1))) | 1))}
                                _ = v84
                            default:
                                var v85 int16 = -512
                                _ = v85
                            }
                            v3 = 2
                        }
                        switch "hello" {
                        case "hello":
                            fmt.Println(v14)
                        }
                    } else {
                        if (4 <= 1) {
                            var v86 uint = v3
                            switch fi53 {
                            case 128:
                                switch v6 {
                                case "\n\t":
                                    switch v6 {
                                    case "foo bar":
                                        var v87 uint8 = (^(v14))
                                        _ = v87
                                    case "a":
                                        for fi88 := 0; fi88 < 12; fi88++ {
                                            _ = len(v5)
                                            _ = fi88
                                        }
                                    default:
                                        for fi89 := 0; fi89 < 17; fi89++ {
                                            v90 := GenS5{F0: v31, F1: (v86 | (v3 / ((1) | 1)))}
                                            v91 := GenS6{F0: -1, F1: 4294967295, F2: int8(v31)}
                                            if (v29 || (v6 >= "")) {
                                                v92 := []int{}
                                                v29 = true
                                                _ = v92
                                            } else {
                                                if (!true) {
                                                    var v93 uint16 = v4
                                                    if (!((v93 % (((^(v93))) | 1)) <= (v4 - (v93 | 922)))) {
                                                        fmt.Println(v14, v86)
                                                        switch (fi89 ^ int(v31)) {
                                                        case -73:
                                                            if v29 {
                                                                switch "" {
                                                                case "foo bar":
                                                                    var v94 int8 = 58
                                                                    if false {
                                                                        v94 = (v94 - v94)
                                                                    } else {
                                                                        var v95 uint16 = v93
                                                                        _ = v95
                                                                    }
                                                                    if false {
                                                                        if (fi89 != (fi89 << ((uint(v14)) & 63))) {
                                                                            var v96 uint8 = v14
                                                                            v97 := GenS7{F0: v3, F1: (^(v3)), F2: (v94 & (v94 * 127))}
                                                                            var v98 uint32 = 973
                                                                            _ = v96
                                                                            _ = v97
                                                                            _ = v98
                                                                        }
                                                                        v2 = (v2 * -9223372036854775808)
                                                                    } else {
                                                                        fmt.Println(fi89, v29, v4)
                                                                        v99 := []int{}
                                                                        var v100 uint = v86
                                                                        _ = v100
                                                                        _ = v99
                                                                    }
                                                                    _ = v94
                                                                case "hello":
                                                                    switch v6 {
                                                                    case "\n\t":
                                                                        var v101 float64 = (v15 - (float64(float32(v15)) + float64(0.0)))
                                                                        _ = v101
                                                                    }
                                                                case "":
                                                                    var v102 uint = v86
                                                                    _ = v102
                                                                }
                                                            }
                                                        default:
                                                            fmt.Println(v29, v29, fi89)
                                                        }
                                                        var v103 uint8 = v14
                                                        var v104 int64 = v2
                                                        v103 += (v14 & 1)
                                                        _ = v103
                                                        _ = v104
                                                    }
                                                    fmt.Println(v93, v29, v29)
                                                    _ = v93
                                                }
                                                if false {
                                                    v105 := GenS1{F0: 65535}
                                                    _ = v105
                                                } else {
                                                    v29 = ((0 != 2) && (v31 <= v31))
                                                }
                                                if false {
                                                    switch 9223372036854775807 {
                                                    case 421:
                                                        fmt.Println(v29, v86)
                                                    default:
                                                        if (!(((-2.25 > (3.14159 - 1e10)) || (v14 != (v14 % ((1) | 1)))) || (127 <= 100))) {
                                                            if v29 {
                                                                v106 := GenS1{F0: v4}
                                                                _ = v106
                                                            } else {
                                                                var v107 uint8 = v14
                                                                v52 = v52
                                                                _ = v107
                                                            }
                                                            if ("" <= (v6 + v6)) {
                                                                v3 = (v86 >> ((v86) & 63))
                                                            }
                                                            var v108 uint16 = (v4 * 65535)
                                                            _ = v108
                                                        }
                                                    }
                                                    v3 = (v86 - (^(v86)))
                                                    switch (v6 + "\n\t") {
                                                    case "a":
                                                        var v109 uint16 = v4
                                                        _ = v109
                                                    case "foo bar":
                                                        if true {
                                                            if (float32(0.1) > -2.25) {
                                                                if ((-(v2)) >= (v2 << (((v3 * (v3 / (((v3 << ((2) & 63))) | 1)))) & 63))) {
                                                                    switch 32768 {
                                                                    case 32768:
                                                                        v110 := GenS4{F0: (^(v4))}
                                                                        _ = v110
                                                                    case 163:
                                                                        var v111 int16 = 349
                                                                        _ = v111
                                                                    case 0:
                                                                        if ((^(v4)) > 775) {
                                                                            v112 := []int{}
                                                                            v4 += 255
                                                                            v15 -= v15
                                                                            _ = v112
                                                                        }
                                                                        v113 := []int{(fi89 | (fi53 & fi89))}
                                                                        _ = v113
                                                                    default:
                                                                        v114 := GenS2{F0: 1000, F1: (v31 - (^(v31)))}
                                                                        _ = v114
                                                                    }
                                                                    fmt.Println(v14)
                                                                } else {
                                                                    fmt.Println(fi53)
                                                                    switch v6 {
                                                                    case "a":
                                                                        switch 9223372036854775807 {
                                                                        case -426:
                                                                            v5[42] = (fi89 ^ (^(fi89)))
                                                                        }
                                                                    default:
                                                                        if (255 == 65535) {
                                                                            v29 = false
                                                                            v15 += float64(-2.25)
                                                                            v115 := GenS8{F0: (v2 & v2)}
                                                                            _ = v115
                                                                        } else {
                                                                            v116 := []int{(fi53 >> (((v3 * 18446744073709551615)) & 63)), (fi89 & 496), int(v86)}
                                                                            _ = v116
                                                                        }
                                                                    }
                                                                    if v29 {
                                                                        var v117 string = v6
                                                                        _ = v117
                                                                    }
                                                                    switch -74 {
                                                                    case 128:
                                                                        _ = len(v5)
                                                                    }
                                                                    v118 := GenS7{F0: v86, F1: v86, F2: -107}
                                                                    _ = v118
                                                                }
                                                                switch v6 {
                                                                case "hello":
                                                                    var v119 int32 = 0
                                                                    _ = v119
                                                                case "a":
                                                                    _ = len(v5)
                                                                case "foo bar":
                                                                    if (v29 || (!(!false))) {
                                                                        _ = len(v5)
                                                                    }
                                                                default:
                                                                    _ = len(v5)
                                                                }
                                                            } else {
                                                                v120 := GenS9{F0: int16(v2)}
                                                                _ = v120
                                                            }
                                                        } else {
                                                            v121 := GenS10{F0: v3}
                                                            if (!(!v29)) {
                                                                if (!(104 != -99)) {
                                                                    var v122 int32 = (v52 / (((v1 - -356)) | 1))
                                                                    v123 := []int{1, 4294967295, fi53}
                                                                    v31 -= v31
                                                                    v124 := GenS11{F0: v6, F1: 42}
                                                                    if v29 {
                                                                        fmt.Println(v14)
                                                                    }
                                                                    var v125 uint8 = 86
                                                                    var v126 int16 = -124
                                                                    _ = v122
                                                                    _ = v123
                                                                    _ = v124
                                                                    _ = v125
                                                                    _ = v126
                                                                }
                                                            } else {
                                                                delete(v5, fi89)
                                                                v127 := []int{int(v31), fi89, (-(fi89))}
                                                                fmt.Println(v52, v29)
                                                                v31 = v31
                                                                var v128 uint = 812
                                                                _ = v127
                                                                _ = v128
                                                            }
                                                            switch 2147483647 {
                                                            case 0:
                                                                if false {
                                                                    if true {
                                                                        var v129 uint16 = (v4 | 255)
                                                                        var v130 int32 = -2147483648
                                                                        switch v6 {
                                                                        case "foo bar":
                                                                            fmt.Println(fi53, v3, v86)
                                                                        case "\n\t":
                                                                            var v131 uint = v86
                                                                            _ = v131
                                                                        case "a":
                                                                            v132 := GenS8{F0: v2}
                                                                            _ = v132
                                                                        default:
                                                                            var v133 uint64 = v31
                                                                            _ = v133
                                                                        }
                                                                        _ = v129
                                                                        _ = v130
                                                                    } else {
                                                                        if (-949 <= (fi53 % ((1) | 1))) {
                                                                            v134 := GenS12{F0: (v52 ^ (v1 ^ 0))}
                                                                            v135 := []int{817, fi89, fi89, (fi53 - fi89)}
                                                                            var v136 int16 = 1
                                                                            _ = v134
                                                                            _ = v135
                                                                            _ = v136
                                                                        } else {
                                                                            v137 := []int{}
                                                                            _ = v137
                                                                        }
                                                                        v5[(fi53 / ((fi89) | 1))] = 923
                                                                        var v138 uint16 = (v4 >> ((1) & 63))
                                                                        _ = v138
                                                                    }
                                                                    var v139 uint8 = 2
                                                                    _ = v139
                                                                }
                                                            case -1:
                                                                var v140 uint32 = 1
                                                                _ = v140
                                                            case 42:
                                                                v29 = v29
                                                            }
                                                            _ = v121
                                                        }
                                                    }
                                                }
                                                var v141 uint32 = 42
                                                var v142 uint32 = (^(v141))
                                                _ = v141
                                                _ = v142
                                            }
                                            _ = fi89
                                            _ = v90
                                            _ = v91
                                        }
                                    }
                                }
                            case 42:
                                var v143 uint32 = 65535
                                _ = v143
                            }
                            switch fi53 {
                            case 128:
                                for fi144 := 0; fi144 < 12; fi144++ {
                                    v145 := v5[(fi53 >> ((602) & 63))]
                                    fmt.Println(fi144, v31)
                                    _ = fi144
                                    _ = v145
                                }
                            }
                            _ = v86
                        } else {
                            fmt.Println(v6, v4, v29)
                        }
                        fmt.Println(fi53, v6, v29)
                    }
                    var v146 int16 = 922
                    _ = fi53
                    _ = v146
                }
                _ = v52
            }
        }
    }
    v15 += v15
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v29
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
