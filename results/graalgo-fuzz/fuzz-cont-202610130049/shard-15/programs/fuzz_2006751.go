package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string; F1 uint8; F2 int }
type GenS4 struct { F0 uint; F1 uint }
type GenS5 struct { F0 int16; F1 uint }
type GenS6 struct { F0 int16; F1 string }
type GenS7 struct { F0 int8; F1 int64; F2 int8 }
type GenS8 struct { F0 int16; F1 int8; F2 int }
type GenS9 struct { F0 string }
type GenS10 struct { F0 uint8; F1 uint64 }
type GenS11 struct { F0 uint64; F1 int64 }
type GenS12 struct { F0 uint16; F1 uint8 }
type GenS13 struct { F0 bool; F1 uint32 }
type GenS14 struct { F0 int32 }
func genFn0() uint32 {
    v7 := GenS1{F0: 376, F1: 0, F2: 128}
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 int64 = 1
        _ = fi8
        _ = v9
    }
    var v10 int16 = 32767
    v11 := []int{-9223372036854775808, -16, 498}
    if (false || (-128 <= 1)) {
        for fi12 := 0; fi12 < 12; fi12++ {
            v11 = append(v11, fi12)
            var v13 uint64 = 876
            if false {
                if (1 != 42) {
                    for fi14 := 0; fi14 < 11; fi14++ {
                        fmt.Println(v10, fi12, v13)
                        switch "\n\t" {
                        case "":
                            v10 = (v10 * (v10 & 175))
                        case "a":
                            fmt.Println(fi14)
                        case "\n\t":
                            var v15 uint8 = 2
                            _ = v15
                        }
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                        _ = fi14
                    }
                    switch fi12 {
                    case -151:
                        var v16 uint16 = 42
                        _ = v16
                    case 977:
                        switch "foo bar" {
                        case "a":
                            if false {
                                v11 = append(v11, (fi12 / ((fi12) | 1)))
                                v13 = (v13 + v13)
                            }
                        case "foo bar":
                            fmt.Println(v13, v10, fi12)
                        case "hello":
                            var v17 uint64 = (v13 | (v13 / ((v13) | 1)))
                            _ = v17
                        }
                    }
                    if false {
                        switch "hello" {
                        case "foo bar":
                            v18 := GenS2{F0: (false || false)}
                            _ = v18
                        case "hello":
                            v19 := GenS3{F0: "\n\t", F1: 0, F2: fi12}
                            _ = v19
                        case "a":
                            v13 = v13
                        default:
                            v20 := GenS4{F0: 484, F1: 938}
                            _ = v20
                        }
                        fmt.Println(fi12, v13, v10)
                        fmt.Println(v13, fi12, v10)
                    } else {
                        if true {
                            v13 += v13
                        } else {
                            var v21 uint16 = 979
                            v21 += 0
                            v10 = v10
                            _ = v21
                        }
                    }
                } else {
                    switch (fi12 ^ (fi12 >> ((65535) & 63))) {
                    case -112:
                        fmt.Println(fi12)
                    }
                    for fi22 := 0; fi22 < 19; fi22++ {
                        fmt.Println(fi22, v13)
                        var v23 int32 = -1
                        if ((0 != 50) || false) {
                            fmt.Println(fi12, v13)
                            v10 += 1
                            v13 += v13
                        } else {
                            fmt.Println(v10)
                            v23 = v23
                            v10 = (v10 + v10)
                        }
                        _ = fi22
                        _ = v23
                    }
                    for fi24 := 0; fi24 < 18; fi24++ {
                        var v25 string = "hello"
                        var v26 int64 = 1
                        if true {
                            switch (fi12 >> ((85) & 63)) {
                            case 42:
                                var v27 string = v25
                                _ = v27
                            case 127:
                                fmt.Println(v26, fi12)
                            case 9223372036854775807:
                                if (false || false) {
                                    var v28 int = (fi24 | -145)
                                    _ = v28
                                } else {
                                    v13 = (v13 & (^(v13)))
                                    fmt.Println(fi24, v13)
                                }
                            default:
                                _ = len(v11)
                                v25 = ""
                                switch v25 {
                                case "hello":
                                    _ = len(v11)
                                case "\n\t":
                                    var v29 uint32 = 0
                                    _ = v29
                                case "a":
                                    fmt.Println(fi24, fi24, v26)
                                default:
                                    _ = len(v11)
                                }
                                if (v25 <= "\n\t") {
                                    v11 = append(v11, 204)
                                } else {
                                    fmt.Println(v10)
                                    switch v25 {
                                    case "a":
                                        _ = len(v11)
                                    default:
                                        v30 := GenS4{F0: 4294967295, F1: 538}
                                        _ = v30
                                    }
                                    if (42 > 127) {
                                        var v31 uint = 4294967295
                                        v13 -= (v13 << (((v31 - v31)) & 63))
                                        var v32 float32 = 1.5
                                        _ = v31
                                        _ = v32
                                    } else {
                                        if len(v11) > 0 {
                                            v11[0] = fi24
                                        }
                                    }
                                    var v33 int16 = v10
                                    _ = v33
                                }
                            }
                            if true {
                                var v34 uint64 = v13
                                var v35 string = v25
                                switch v35 {
                                case "\n\t":
                                    if false {
                                        v13 = (^(v13))
                                        if (true || ((((v35 + ((v25 + v35) + v25)) + v35) != ((v25 + "\n\t") + v35)) || false)) {
                                            var v36 int8 = -1
                                            var v37 uint8 = 211
                                            switch v35 {
                                            case "a":
                                                v37 = 1
                                            case "foo bar":
                                                v38 := GenS2{F0: false}
                                                _ = v38
                                            }
                                            _ = v36
                                            _ = v37
                                        } else {
                                            var v39 int = -1
                                            if len(v11) > 0 {
                                                v11[0] = fi12
                                            }
                                            if (0 < -1) {
                                                v40 := GenS5{F0: v10, F1: 255}
                                                _ = v40
                                            } else {
                                                v39 = fi24
                                                if (2 >= 86) {
                                                    v41 := GenS1{F0: 388, F1: 0, F2: 0}
                                                    var v42 uint8 = 105
                                                    _ = v41
                                                    _ = v42
                                                }
                                                v11 = append(v11, v39)
                                            }
                                            v13 -= v34
                                            _ = v39
                                        }
                                    } else {
                                        if len(v11) > 0 {
                                            v11[0] = fi12
                                        }
                                        fmt.Println(v26, fi24)
                                        var v43 string = ("a" + ((v25 + (v25 + "hello")) + (v25 + ((v35 + "a") + ""))))
                                        _ = v43
                                    }
                                case "a":
                                    fmt.Println(v25, v10)
                                case "hello":
                                    v44 := GenS3{F0: v35, F1: 127, F2: (^(fi12))}
                                    _ = v44
                                }
                                _ = v34
                                _ = v35
                            }
                            if false {
                                v26 += v26
                                v45 := GenS6{F0: int16(fi24), F1: ((v25 + v25) + v25)}
                                var v46 uint = 1
                                _ = v45
                                _ = v46
                            } else {
                                if ((!false) && (!((42 >= 4294967295) && ((-(fi12)) < fi24)))) {
                                    v11 = append(v11, fi24)
                                } else {
                                    v26 += v26
                                    switch (-(fi12)) {
                                    case -9223372036854775808:
                                        var v47 int32 = 919
                                        _ = v47
                                    case 128:
                                        v48 := GenS7{F0: -111, F1: (v26 * (-(v26))), F2: -1}
                                        _ = v48
                                    case 1:
                                        var v49 int64 = -9223372036854775808
                                        _ = v49
                                    }
                                    var v50 float64 = 0.1
                                    _ = v50
                                }
                            }
                        } else {
                            fmt.Println(v26, fi12)
                            var v51 uint32 = 850
                            var v52 uint64 = (v13 & 42)
                            fmt.Println(v51)
                            if (1000 <= 42) {
                                if ((-(fi24)) >= 2147483647) {
                                    var v53 uint8 = 128
                                    _ = v53
                                } else {
                                    v11 = append(v11, (fi12 % (((-(fi24))) | 1)))
                                    v11 = append(v11, (fi24 - fi12))
                                    v11 = append(v11, (^(fi12)))
                                }
                                v26 -= v26
                            } else {
                                v11 = append(v11, int(v52))
                                v13 += v13
                                v13 = (v52 + v52)
                            }
                            v11 = append(v11, fi24)
                            _ = v51
                            _ = v52
                        }
                        v54 := GenS2{F0: false}
                        _ = fi24
                        _ = v25
                        _ = v26
                        _ = v54
                    }
                }
                var v55 float32 = 0.0
                _ = v55
            }
            _ = fi12
            _ = v13
        }
        v10 = (-(v10))
    } else {
        fmt.Println(v10)
    }
    _ = v10
    _ = v11
    _ = v7
    return 711
}
func genFn1() int {
    for fi56 := 0; fi56 < 11; fi56++ {
        var v57 uint8 = 128
        var v58 uint16 = 32767
        _ = fi56
        _ = v57
        _ = v58
    }
    var v59 uint64 = 84
    if false {
        if ("hello" >= ((("hello" + "\n\t") + (("" + "") + "\n\t")) + "foo bar")) {
            if false {
                v59 -= (v59 | (v59 + v59))
            }
            fmt.Println(v59)
            if false {
                v60 := []int{-1, 2147483647, 9223372036854775807, 1}
                var v61 float32 = 1.5
                v62 := GenS7{F0: -81, F1: -1, F2: -1}
                _ = v60
                _ = v61
                _ = v62
            } else {
                fmt.Println(v59)
                var v63 uint16 = 255
                var v64 uint8 = 127
                _ = v63
                _ = v64
            }
        } else {
            fmt.Println(v59)
            v65 := []int{-422, 2147483647}
            _ = v65
        }
        var v66 int64 = 9223372036854775807
        v67 := []int{305, 594}
        _ = v66
        _ = v67
    }
    v68 := GenS4{F0: 18446744073709551615, F1: 65535}
    _ = v59
    _ = v68
    return 0
}

func main() {
    v1 := []int{-1, -1, 42}
    var v2 uint64 = 4294967295
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]bool)
    var v5 float64 = float64((1.5 / -2.25))
    var v6 uint32 = 2147483647
    var v69 int32 = 2147483647
    v70 := v4[(v6 << ((0) & 63))]
    for fi71 := 0; fi71 < 19; fi71++ {
        fmt.Println(fi71, v6)
        var v72 float64 = float64(0.0)
        var v73 bool = v70
        _ = fi71
        _ = v72
        _ = v73
    }
    var v74 uint = 1
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if v70 {
        _ = len(v1)
        if true {
            if true {
                _ = len(v3)
                var v75 float64 = v5
                _ = v75
            } else {
                v70 = true
                if len(v1) > 0 {
                    v1[0] = -310
                }
                v76 := GenS7{F0: 44, F1: 9223372036854775807, F2: 70}
                _ = v76
            }
            if (v70 || v70) {
                v77 := GenS3{F0: "a", F1: 42, F2: 127}
                fmt.Println(v2, v6, v70)
                for fi78 := 0; fi78 < 16; fi78++ {
                    var v79 uint16 = 255
                    var v80 bool = false
                    v81 := GenS7{F0: -128, F1: 0, F2: -99}
                    _ = fi78
                    _ = v79
                    _ = v80
                    _ = v81
                }
                _ = v77
            }
        }
    } else {
        fmt.Println(v5, v70, v5)
        if v70 {
            var v82 int16 = -32768
            _ = v82
        } else {
            for fi83 := 0; fi83 < 16; fi83++ {
                if v70 {
                    v70 = false
                    _ = len(v3)
                    for fi84 := 0; fi84 < 2; fi84++ {
                        if v70 {
                            fmt.Println(v5)
                            var v85 int8 = -128
                            _ = v85
                        } else {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            var v86 uint8 = 2
                            if v70 {
                                if (v5 == v5) {
                                    var v87 uint64 = (v2 & (v2 & v2))
                                    var v88 int32 = (v69 << ((65535) & 63))
                                    _ = v87
                                    _ = v88
                                } else {
                                    v89 := GenS3{F0: ("a" + "hello"), F1: (v86 % (((v86 | (v86 ^ (^(v86))))) | 1)), F2: (^(fi84))}
                                    var v90 string = "\n\t"
                                    _ = v89
                                    _ = v90
                                }
                                v2 = v2
                                v91 := GenS3{F0: "foo bar", F1: (v86 - 1), F2: -17}
                                _ = v91
                            } else {
                                if (true && false) {
                                    var v92 int32 = 0
                                    var v93 uint = (v74 % (((v74 << ((2) & 63))) | 1))
                                    v94 := GenS7{F0: 42, F1: -9223372036854775808, F2: 126}
                                    v5 += v5
                                    v6 = (v6 + uint32(v2))
                                    if ((v74 >> ((v74) & 63)) > v74) {
                                        if (!v70) {
                                            if v70 {
                                                v86 = (v86 + 127)
                                            }
                                            switch "foo bar" {
                                            case "hello":
                                                if true {
                                                    v95 := GenS8{F0: -1, F1: 1, F2: fi83}
                                                    v6 -= (v6 & (v6 ^ uint32(v86)))
                                                    var v96 int64 = -1
                                                    if v70 {
                                                        v86 = 237
                                                    }
                                                    _ = v95
                                                    _ = v96
                                                }
                                            case "":
                                                var v97 int16 = int16(fi84)
                                                _ = v97
                                            default:
                                                v98 := GenS9{F0: "foo bar"}
                                                v99 := GenS10{F0: (v86 % ((0) | 1)), F1: 448}
                                                _ = v98
                                                _ = v99
                                            }
                                        }
                                        _ = len(v1)
                                        v2 -= v2
                                    }
                                    var v100 uint16 = 254
                                    if v70 {
                                        v86 = v86
                                        var v101 uint32 = v6
                                        var v102 string = "foo bar"
                                        _ = v101
                                        _ = v102
                                    } else {
                                        v100 -= uint16(v6)
                                    }
                                    fmt.Println(fi83)
                                    _ = v100
                                    _ = v92
                                    _ = v93
                                    _ = v94
                                }
                                if true {
                                    var v103 int16 = 42
                                    var v104 uint16 = 42
                                    _ = v103
                                    _ = v104
                                } else {
                                    var v105 uint = (^(v74))
                                    _ = v105
                                }
                            }
                            _ = v86
                        }
                        _ = fi84
                    }
                }
                _ = fi83
            }
            if false {
                for fi106 := 0; fi106 < 10; fi106++ {
                    v5 += v5
                    v107 := v4[(v6 >> (((v74 ^ (v74 - (v74 | (v74 | 1))))) & 63))]
                    _ = fi106
                    _ = v107
                }
                for fi108 := 0; fi108 < 16; fi108++ {
                    if true {
                        if (32767 < 369) {
                            for fi109 := 0; fi109 < 5; fi109++ {
                                var v110 int32 = (v69 << ((v74) & 63))
                                _ = fi109
                                _ = v110
                            }
                        } else {
                            v69 = v69
                        }
                    } else {
                        switch (fi108 + -449) {
                        case 326:
                            var v111 bool = false
                            _ = v111
                        default:
                            if v70 {
                                v6 = v6
                                _ = len(v4)
                            }
                        }
                    }
                    v112 := GenS11{F0: (v2 % (((v2 & v2)) | 1)), F1: -1}
                    _ = fi108
                    _ = v112
                }
                if v70 {
                    for fi113 := 0; fi113 < 3; fi113++ {
                        var v114 int32 = -2147483648
                        _ = fi113
                        _ = v114
                    }
                } else {
                    var v115 int = 127
                    v116 := v4[genFn0()]
                    _ = v115
                    _ = v116
                }
            }
        }
    }
    var v117 float32 = (1e10 * (1.5 * 1.5))
    if (v74 <= 42) {
        switch (("\n\t" + "hello") + ("hello" + "")) {
        case "\n\t":
            switch "foo bar" {
            case "hello":
                switch "\n\t" {
                case "":
                    switch "foo bar" {
                    case "a":
                        v118 := GenS5{F0: int16(v74), F1: v74}
                        _ = v118
                    case "hello":
                        var v119 int64 = 388
                        _ = v119
                    case "foo bar":
                        for fi120 := 0; fi120 < 15; fi120++ {
                            v1 = append(v1, (fi120 | 42))
                            _ = fi120
                        }
                    default:
                        v121 := GenS12{F0: 2, F1: 151}
                        _ = v121
                    }
                case "a":
                    for fi122 := 0; fi122 < 7; fi122++ {
                        v117 -= (v117 / -2.25)
                        if len(v1) > 0 {
                            v1[0] = (fi122 << (((^(v74))) & 63))
                        }
                        _ = fi122
                    }
                case "foo bar":
                    switch "a" {
                    case "":
                        v123 := GenS8{F0: -673, F1: 0, F2: -596}
                        _ = v123
                    case "hello":
                        v124 := GenS5{F0: 1000, F1: v74}
                        _ = v124
                    case "foo bar":
                        fmt.Println(v69)
                    }
                default:
                    for fi125 := 0; fi125 < 18; fi125++ {
                        var v126 bool = v70
                        switch "" {
                        case "":
                            var v127 float32 = v117
                            _ = v127
                        case "\n\t":
                            switch "foo bar" {
                            case "a":
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            case "\n\t":
                                if false {
                                    v69 -= -2147483648
                                    var v128 int8 = -1
                                    var v129 bool = (!false)
                                    _ = v128
                                    _ = v129
                                }
                            }
                        }
                        if v126 {
                            for fi130 := 0; fi130 < 5; fi130++ {
                                if ((v6 << ((uint(v2)) & 63)) >= 354) {
                                    v131 := GenS12{F0: 42, F1: 0}
                                    _ = v131
                                }
                                v69 = -2147483648
                                _ = fi130
                            }
                        }
                        _ = fi125
                        _ = v126
                    }
                }
            default:
                v132 := GenS12{F0: 2, F1: 128}
                _ = v132
            }
        case "":
            v133 := v4[v6]
            _ = v133
        }
        if (540 > (v2 ^ (v2 << (((v74 * v74)) & 63)))) {
            var v134 uint8 = 128
            var v135 uint8 = (v134 / ((v134) | 1))
            switch 913 {
            case 0:
                for fi136 := 0; fi136 < 13; fi136++ {
                    for fi137 := 0; fi137 < 15; fi137++ {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = fi137
                    }
                    v1 = append(v1, 32768)
                    v70 = (v70 && (int8(v135) <= 47))
                    _ = fi136
                }
            default:
                fmt.Println(v2, v6, v5)
            }
            _ = v134
            _ = v135
        } else {
            v70 = false
            if true {
                switch -9223372036854775808 {
                case 1:
                    v6 += 163
                default:
                    for fi138 := 0; fi138 < 17; fi138++ {
                        var v139 int8 = -51
                        switch "\n\t" {
                        case "foo bar":
                            switch "\n\t" {
                            case "hello":
                                if (663 <= v74) {
                                    v5 += v5
                                } else {
                                    fmt.Println(v6, v69)
                                }
                            }
                        case "\n\t":
                            v2 += v2
                        case "hello":
                            for fi140 := 0; fi140 < 9; fi140++ {
                                var v141 int32 = -365
                                var v142 int = (fi140 / ((4294967295) | 1))
                                _ = fi140
                                _ = v141
                                _ = v142
                            }
                        }
                        _ = fi138
                        _ = v139
                    }
                }
                var v143 int16 = 32767
                _ = v143
            }
        }
        v144 := GenS1{F0: 255, F1: 255, F2: 128}
        _ = v144
    } else {
        if true {
            for fi145 := 0; fi145 < 19; fi145++ {
                for fi146 := 0; fi146 < 5; fi146++ {
                    v5 = v5
                    v147 := GenS2{F0: (v74 == (v74 + v74))}
                    _ = fi146
                    _ = v147
                }
                switch ("\n\t" + "") {
                case "hello":
                    v148 := GenS6{F0: 32767, F1: "a"}
                    _ = v148
                default:
                    for fi149 := 0; fi149 < 12; fi149++ {
                        var v150 uint64 = v2
                        _ = fi149
                        _ = v150
                    }
                }
                var v151 int16 = 0
                _ = fi145
                _ = v151
            }
            switch 4294967295 {
            case 42:
                var v152 float32 = v117
                _ = v152
            default:
                v4[genFn0()] = (int32(v6) != (v69 + (v69 * v69)))
            }
        } else {
            switch "hello" {
            case "":
                for fi153 := 0; fi153 < 1; fi153++ {
                    var v154 uint32 = (v6 / (((^(v6))) | 1))
                    _ = fi153
                    _ = v154
                }
            case "\n\t":
                fmt.Println(v74)
            case "hello":
                v155 := GenS7{F0: 0, F1: 42, F2: 42}
                _ = v155
            }
            delete(v4, (v6 + v6))
            for fi156 := 0; fi156 < 13; fi156++ {
                fmt.Println(v70)
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi156
            }
        }
        v74 = 1
        for fi157 := 0; fi157 < 16; fi157++ {
            for fi158 := 0; fi158 < 19; fi158++ {
                v6 = (v6 / (((v6 % ((v6) | 1))) | 1))
                if len(v1) > 0 {
                    v1[0] = (-(fi157))
                }
                _ = fi158
            }
            var v159 int32 = (v69 / ((int32(fi157)) | 1))
            v160 := GenS11{F0: (v2 - (^(v2))), F1: -1}
            _ = fi157
            _ = v159
            _ = v160
        }
    }
    var v161 uint32 = (v6 << ((4294967295) & 63))
    var v162 string = (((((("" + "") + ("\n\t" + "")) + ("" + "foo bar")) + ("hello" + "foo bar")) + "\n\t") + "")
    v162 = v162
    for fi163 := 0; fi163 < 14; fi163++ {
        if v70 {
            var v164 uint16 = 2
            var v165 int16 = -521
            switch "hello" {
            case "foo bar":
                v166 := GenS11{F0: v2, F1: 202}
                _ = v166
            case "\n\t":
                if true {
                    v167 := GenS13{F0: ((v164 ^ (^(v164))) > v164), F1: genFn0()}
                    v168 := GenS13{F0: v70, F1: v6}
                    _ = v167
                    _ = v168
                }
            case "hello":
                v6 = genFn0()
            }
            _ = v164
            _ = v165
        }
        if v70 {
            var v169 uint = (v74 - 946)
            var v170 uint8 = 149
            _ = v169
            _ = v170
        } else {
            for fi171 := 0; fi171 < 6; fi171++ {
                fmt.Println(v69)
                _ = fi171
            }
            var v172 uint32 = uint32(v2)
            switch fi163 {
            case 127:
                var v173 uint64 = (v2 ^ v2)
                _ = v173
            default:
                var v174 uint64 = (v2 * 18446744073709551615)
                _ = v174
            }
            _ = v172
        }
        for fi175 := 0; fi175 < 13; fi175++ {
            var v176 int8 = 14
            _ = len(v1)
            switch v162 {
            case "a":
                if true {
                    switch "hello" {
                    case "foo bar":
                        var v177 bool = (!(true && v70))
                        _ = v177
                    default:
                        v178 := GenS12{F0: 255, F1: 42}
                        _ = v178
                    }
                    v179 := GenS14{F0: v69}
                    _ = v179
                }
            case "":
                v180 := GenS5{F0: -1, F1: 2}
                _ = v180
            }
            _ = fi175
            _ = v176
        }
        _ = fi163
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v117:", v117)
    fmt.Println("v161:", v161)
    fmt.Println("v162:", v162)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v117
    _ = v161
    _ = v162
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v69
    _ = v70
    _ = v74
}
