package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool; F1 bool; F2 uint32 }
type GenS4 struct { F0 int; F1 int16; F2 uint16 }
type GenS5 struct { F0 int16; F1 bool }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint32; F1 uint; F2 uint16 }
type GenS9 struct { F0 bool; F1 uint32 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 uint32; F1 uint32 }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 int16; F1 int; F2 int64 }
type GenS14 struct { F0 int; F1 int64 }
func genFn0(p0 int, p1 bool) int8 {
    p0 += -1
    v6 := GenS1{F0: 145}
    var v7 uint = 1
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 127
}
func genFn1() int8 {
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 int64 = -979
        var v10 uint = 1
        v11 := GenS2{F0: (^(fi8))}
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    if (0.1 > -2.25) {
        v12 := GenS2{F0: 211}
        v13 := []int{15, 32767}
        _ = v12
        _ = v13
    } else {
        if (4294967295 <= 65535) {
            v14 := []int{94, 127, 128}
            var v15 int16 = 221
            var v16 float32 = 1e10
            _ = v14
            _ = v15
            _ = v16
        }
        v17 := []int{-77, 741}
        var v18 float64 = 0.0
        _ = v17
        _ = v18
    }
    v19 := GenS1{F0: 752}
    _ = v19
    return 42
}

func main() {
    var v1 bool = true
    var v2 float64 = 0.0
    v3 := make(map[int]uint32)
    var v4 uint = 521
    var v5 int8 = 127
    if (((42 >= 42) || (560 > 2)) || v1) {
        for fi20 := 0; fi20 < 5; fi20++ {
            v3[fi20] = 2147483648
            fmt.Println(v5)
            _ = fi20
        }
    }
    if true {
        fmt.Println(v1, v5, v4)
        for fi21 := 0; fi21 < 3; fi21++ {
            for fi22 := 0; fi22 < 1; fi22++ {
                var v23 uint32 = 2147483648
                _ = fi22
                _ = v23
            }
            if (1000 <= -293) {
                for fi24 := 0; fi24 < 2; fi24++ {
                    var v25 int = (-(fi24))
                    var v26 int8 = genFn0((fi24 ^ fi24), (!(0.1 > 0.1)))
                    v4 -= v4
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                if ((true || false) && ((-(v5)) >= (v5 * (^(v5))))) {
                    v27 := GenS3{F0: ("hello" <= ("foo bar" + (("hello" + "a") + ("" + "hello")))), F1: ((((-9223372036854775808 < 1) || ((424 == 4294967295) || (true && false))) || ((!(!(uint64(v4) != 18446744073709551615))) || v1)) && v1), F2: 0}
                    for fi28 := 0; fi28 < 11; fi28++ {
                        var v29 int16 = 0
                        if v1 {
                            switch (fi21 * 9223372036854775807) {
                            case 42:
                                fmt.Println(v1, v29, v4)
                            }
                        }
                        _ = fi28
                        _ = v29
                    }
                    fmt.Println(v4, v5, v1)
                    _ = v27
                }
                var v30 int8 = 1
                _ = v30
            } else {
                v5 -= (v5 | (v5 - (v5 << (((v4 ^ v4)) & 63))))
            }
            _ = fi21
        }
        fmt.Println(v5)
    } else {
        if true {
            var v31 bool = (1e10 < 0.1)
            v31 = v1
            switch 0 {
            case 128:
                if v31 {
                    for fi32 := 0; fi32 < 14; fi32++ {
                        v33 := GenS4{F0: fi32, F1: -1, F2: 32767}
                        _ = fi32
                        _ = v33
                    }
                    v5 += 1
                } else {
                    var v34 string = "hello"
                    _ = v34
                }
            case 0:
                v4 = 18446744073709551615
            case 4294967295:
                if (0 >= 4294967295) {
                    v35 := []int{-529, -30, -252, 9223372036854775807}
                    for fi36 := 0; fi36 < 19; fi36++ {
                        switch ("" + "a") {
                        case "foo bar":
                            var v37 float32 = 1e10
                            _ = v37
                        case "":
                            v4 = v4
                        case "\n\t":
                            delete(v3, fi36)
                        default:
                            var v38 uint16 = 0
                            _ = v38
                        }
                        delete(v3, fi36)
                        v39 := GenS3{F0: (32768 <= fi36), F1: (false || v1), F2: 2147483648}
                        _ = fi36
                        _ = v39
                    }
                    _ = v35
                }
            }
            _ = v31
        } else {
            if true {
                var v40 float64 = v2
                for fi41 := 0; fi41 < 1; fi41++ {
                    switch fi41 {
                    case 2147483647:
                        delete(v3, fi41)
                    case 32768:
                        fmt.Println(v4, v4)
                    }
                    v40 -= (v2 / v40)
                    _ = fi41
                }
                for fi42 := 0; fi42 < 12; fi42++ {
                    var v43 uint8 = 128
                    _ = fi42
                    _ = v43
                }
                _ = v40
            }
        }
        if (0 == 127) {
            var v44 uint32 = 4294967295
            v45 := GenS3{F0: v1, F1: v1, F2: (v44 ^ v44)}
            _ = v44
            _ = v45
        }
        var v46 float64 = (v2 + v2)
        _ = v46
    }
    v1 = false
    v5 += 117
    switch 509 {
    case 4294967295:
        fmt.Println(v5)
    default:
        var v47 float64 = v2
        _ = v47
    }
    if v1 {
        for fi48 := 0; fi48 < 1; fi48++ {
            for fi49 := 0; fi49 < 8; fi49++ {
                v5 = genFn0(int(v4), (255 > 0))
                if v1 {
                    v50 := GenS1{F0: 42}
                    if v1 {
                        var v51 uint64 = 4294967295
                        fmt.Println(fi48, v1, v5)
                        fmt.Println(v2, v1, v51)
                        _ = v51
                    } else {
                        var v52 int16 = 32767
                        _ = v52
                    }
                    v3[(fi49 | (^(fi48)))] = 1
                    _ = v50
                } else {
                    v5 = v5
                }
                v53 := []int{(-(fi49)), (fi49 >> ((v4) & 63)), (fi49 + 0), 4294967295}
                _ = fi49
                _ = v53
            }
            for fi54 := 0; fi54 < 19; fi54++ {
                v55 := GenS5{F0: 42, F1: false}
                var v56 uint64 = 182
                switch 261 {
                case 2147483647:
                    delete(v3, 396)
                case -528:
                    var v57 uint32 = 2147483648
                    _ = v57
                }
                v58 := GenS6{F0: 765}
                _ = fi54
                _ = v55
                _ = v56
                _ = v58
            }
            v59 := v3[-373]
            _ = fi48
            _ = v59
        }
    } else {
        var v60 uint8 = 0
        fmt.Println(v1, v1, v1)
        var v61 int32 = int32(v5)
        _ = v60
        _ = v61
    }
    switch "a" {
    case "":
        if (!false) {
            v62 := []int{128, -9223372036854775808, 32767, 9223372036854775807}
            _ = v62
        } else {
            if ((v2 * v2) <= v2) {
                for fi63 := 0; fi63 < 16; fi63++ {
                    v4 -= 0
                    v4 += (^(v4))
                    for fi64 := 0; fi64 < 20; fi64++ {
                        v65 := []int{}
                        _ = fi64
                        _ = v65
                    }
                    _ = fi63
                }
                v1 = (v5 >= v5)
            } else {
                fmt.Println(v1, v1, v4)
                for fi66 := 0; fi66 < 18; fi66++ {
                    if (v1 && v1) {
                        for fi67 := 0; fi67 < 6; fi67++ {
                            if true {
                                if (!(4294967295 == 42)) {
                                    if (v1 || true) {
                                        v68 := []int{}
                                        switch fi66 {
                                        case -926:
                                            var v69 int8 = v5
                                            v69 = v5
                                            _ = v69
                                        case 9223372036854775807:
                                            var v70 uint8 = 130
                                            _ = v70
                                        case 1:
                                            var v71 uint8 = 255
                                            _ = v71
                                        default:
                                            v4 = v4
                                        }
                                        v72 := GenS6{F0: 0}
                                        _ = v68
                                        _ = v72
                                    } else {
                                        v73 := []int{fi67}
                                        v74 := GenS7{F0: 18446744073709551615}
                                        var v75 uint = (v4 >> ((18446744073709551615) & 63))
                                        _ = v73
                                        _ = v74
                                        _ = v75
                                    }
                                } else {
                                    switch "hello" {
                                    case "\n\t":
                                        if (!(false || v1)) {
                                            var v76 int32 = 1
                                            if false {
                                                if (115 != 9) {
                                                    fmt.Println(v4, fi67)
                                                    var v77 float32 = -2.25
                                                    _ = v77
                                                } else {
                                                    if (v1 && (false || ((-2.25 - float32(0.0)) > 0.1))) {
                                                        switch "\n\t" {
                                                        case "foo bar":
                                                            var v78 uint32 = uint32(v4)
                                                            _ = v78
                                                        case "hello":
                                                            var v79 uint8 = 1
                                                            _ = v79
                                                        case "\n\t":
                                                            var v80 uint32 = 0
                                                            _ = v80
                                                        default:
                                                            if v1 {
                                                                v3[-2147483648] = 4294967295
                                                                if ((v1 || ((v1 || v1) || ((v76 | int32(v5)) >= (v76 % (((-(v76))) | 1))))) && ((-(fi66)) > -1)) {
                                                                    var v81 int32 = v76
                                                                    var v82 int8 = genFn1()
                                                                    v5 += (v82 * (v82 - v5))
                                                                    _ = v81
                                                                    _ = v82
                                                                }
                                                                v5 -= v5
                                                            } else {
                                                                var v83 string = "hello"
                                                                v84 := []int{(^(fi66)), fi66}
                                                                v83 = "hello"
                                                                if len(v84) >= 2 {
                                                                    v84 = v84[1:len(v84)]
                                                                }
                                                                var v85 int64 = -1
                                                                v4 -= v4
                                                                _ = v83
                                                                _ = v84
                                                                _ = v85
                                                            }
                                                        }
                                                    }
                                                }
                                            }
                                            _ = v76
                                        }
                                    case "":
                                        var v86 uint16 = 340
                                        _ = v86
                                    case "foo bar":
                                        switch "a" {
                                        case "a":
                                            switch "a" {
                                            case "\n\t":
                                                v87 := GenS5{F0: 1, F1: v1}
                                                _ = v87
                                            case "":
                                                switch "foo bar" {
                                                case "hello":
                                                    var v88 uint8 = 205
                                                    _ = v88
                                                }
                                                v3[fi66] = 4294967295
                                                v89 := GenS7{F0: 0}
                                                var v90 uint64 = 2
                                                _ = v89
                                                _ = v90
                                            default:
                                                switch fi67 {
                                                case 127:
                                                    if (!v1) {
                                                        v91 := v3[fi67]
                                                        v92 := GenS4{F0: (-(fi67)), F1: 32767, F2: 65535}
                                                        _ = v91
                                                        _ = v92
                                                    } else {
                                                        _ = len(v3)
                                                    }
                                                }
                                            }
                                        }
                                    default:
                                        switch fi67 {
                                        case -530:
                                            switch "foo bar" {
                                            case "foo bar":
                                                v4 = (v4 % ((837) | 1))
                                            default:
                                                var v93 uint = (v4 / ((1) | 1))
                                                _ = v93
                                            }
                                        case 0:
                                            if (65535 >= 255) {
                                                fmt.Println(fi66, v1)
                                                var v94 string = ""
                                                v5 = v5
                                                if true {
                                                    v94 = ""
                                                    var v95 int32 = -1
                                                    if ((((fi66 * 42) == fi67) && true) || (42 > -1)) {
                                                        var v96 uint8 = 2
                                                        v94 = "\n\t"
                                                        _ = v96
                                                    }
                                                    _ = v95
                                                } else {
                                                    switch v94 {
                                                    case "\n\t":
                                                        switch "\n\t" {
                                                        case "a":
                                                            var v97 string = (("" + (v94 + ("hello" + ""))) + "foo bar")
                                                            _ = v97
                                                        case "hello":
                                                            fmt.Println(v5)
                                                        default:
                                                            fmt.Println(v2, fi67, v94)
                                                        }
                                                    default:
                                                        var v98 float64 = ((v2 - 1e10) - (((v2 - v2) / float64(-2.25)) + (((v2 * v2) - (1.5 * (float64((1.5 / (1e10 + 0.1))) + v2))) + (v2 / -2.25))))
                                                        _ = v98
                                                    }
                                                    if true {
                                                        v2 = 0.1
                                                        v2 = v2
                                                        v99 := GenS2{F0: 127}
                                                        _ = v99
                                                    }
                                                }
                                                _ = v94
                                            } else {
                                                fmt.Println(v5, v1)
                                                var v100 uint8 = 2
                                                var v101 float32 = 1.5
                                                _ = v100
                                                _ = v101
                                            }
                                            var v102 int = fi66
                                            switch "" {
                                            case "a":
                                                var v103 uint16 = 32767
                                                _ = v103
                                            default:
                                                var v104 uint16 = 255
                                                _ = v104
                                            }
                                            var v105 int64 = -1
                                            var v106 int = fi67
                                            _ = v102
                                            _ = v105
                                            _ = v106
                                        case 4294967295:
                                            if v1 {
                                                v1 = true
                                            } else {
                                                var v107 float32 = 3.14159
                                                var v108 int32 = -922
                                                _ = v107
                                                _ = v108
                                            }
                                        }
                                    }
                                    v5 = -1
                                    if v1 {
                                        var v109 string = ""
                                        v110 := GenS1{F0: 18446744073709551615}
                                        fmt.Println(fi66, v2, v1)
                                        v109 = "hello"
                                        switch (fi67 % (((^(fi67))) | 1)) {
                                        case 128:
                                            v111 := GenS6{F0: 4294967295}
                                            _ = v111
                                        case -2147483648:
                                            var v112 int = fi67
                                            _ = v112
                                        case -1:
                                            v113 := []int{(fi66 << (((v4 >> ((4294967295) & 63))) & 63))}
                                            _ = v113
                                        default:
                                            if (v109 > "hello") {
                                                var v114 float32 = (0.0 - 3.14159)
                                                _ = v114
                                            }
                                        }
                                        _ = v109
                                        _ = v110
                                    }
                                    if (1 != 0) {
                                        v115 := []int{fi66}
                                        var v116 float64 = float64((-2.25 * ((0.0 / 0.1) / 1.0)))
                                        _ = v115
                                        _ = v116
                                    }
                                }
                                switch (-(fi67)) {
                                case 2147483647:
                                    v5 += v5
                                default:
                                    if v1 {
                                        v5 -= v5
                                    }
                                }
                                v4 = (v4 >> ((520) & 63))
                            } else {
                                var v117 int8 = -128
                                _ = v117
                            }
                            _ = fi67
                        }
                    } else {
                        v118 := GenS8{F0: 470, F1: 65535, F2: 936}
                        _ = v118
                    }
                    for fi119 := 0; fi119 < 5; fi119++ {
                        fmt.Println(v2)
                        var v120 uint8 = 127
                        switch (fi66 ^ fi66) {
                        case -715:
                            var v121 float64 = (3.14159 + v2)
                            _ = v121
                        case 32768:
                            var v122 int = 488
                            _ = v122
                        }
                        _ = fi119
                        _ = v120
                    }
                    _ = fi66
                }
                v5 += v5
            }
            v5 -= (v5 - (^(v5)))
            v4 = (^(v4))
        }
    case "\n\t":
        fmt.Println(v4, v1, v1)
    }
    if v1 {
        v123 := []int{32768}
        if ((!(true || (858 > 9223372036854775807))) && v1) {
            v2 -= ((v2 + v2) - -2.25)
            v2 += 0.0
            if v1 {
                if ((127 <= -2147483648) || false) {
                    if ((false || v1) && (true || true)) {
                        v1 = true
                    } else {
                        v124 := GenS1{F0: 0}
                        _ = v124
                    }
                    var v125 int = -2147483648
                    _ = v125
                }
                var v126 int16 = 209
                v1 = v1
                _ = v126
            }
        } else {
            if ((v2 / float64(3.14159)) != 0.1) {
                delete(v3, -645)
                if (!(!(v5 > genFn1()))) {
                    v3[42] = 2147483648
                } else {
                    for fi127 := 0; fi127 < 8; fi127++ {
                        fmt.Println(v2)
                        var v128 bool = v1
                        _ = fi127
                        _ = v128
                    }
                    for fi129 := 0; fi129 < 12; fi129++ {
                        v3[(fi129 << ((v4) & 63))] = 42
                        for fi130 := 0; fi130 < 1; fi130++ {
                            if len(v123) > 0 {
                                v123[0] = (fi130 - -2147483648)
                            }
                            _ = fi130
                        }
                        _ = fi129
                    }
                }
            }
            v2 -= v2
            if false {
                switch -559 {
                case 0:
                    fmt.Println(v2, v2)
                case 1:
                    var v131 int = 128
                    _ = v131
                case 4294967295:
                    if true {
                        for fi132 := 0; fi132 < 5; fi132++ {
                            var v133 string = (("hello" + ("hello" + (("hello" + "a") + "foo bar"))) + "a")
                            v5 = (v5 / (((-(v5))) | 1))
                            v134 := GenS9{F0: v1, F1: 4294967295}
                            _ = fi132
                            _ = v133
                            _ = v134
                        }
                        v135 := GenS10{F0: 42}
                        _ = v135
                    }
                }
                v136 := GenS11{F0: 0, F1: 2147483648}
                fmt.Println(v5, v5)
                _ = v136
            } else {
                for fi137 := 0; fi137 < 10; fi137++ {
                    v138 := GenS6{F0: 584}
                    _ = fi137
                    _ = v138
                }
                for fi139 := 0; fi139 < 20; fi139++ {
                    v4 = (^(v4))
                    if ((v4 | v4) == (v4 | (v4 | 18446744073709551615))) {
                        switch "a" {
                        case "foo bar":
                            var v140 string = ("" + "hello")
                            _ = v140
                        case "a":
                            for fi141 := 0; fi141 < 4; fi141++ {
                                var v142 int = 4294967295
                                v143 := GenS12{F0: 65535}
                                v4 -= 12
                                if ((172 != 2) || (v1 || v1)) {
                                    if (1 > 0) {
                                        v144 := GenS13{F0: -1, F1: (fi141 + (v142 * v142)), F2: 0}
                                        var v145 uint16 = 255
                                        var v146 int8 = 1
                                        _ = v144
                                        _ = v145
                                        _ = v146
                                    } else {
                                        v147 := GenS8{F0: 0, F1: (^(v4)), F2: 65535}
                                        v2 += 3.14159
                                        var v148 uint32 = 2147483648
                                        _ = v147
                                        _ = v148
                                    }
                                    _ = len(v3)
                                    v5 += 127
                                    var v149 int32 = 2147483647
                                    _ = v149
                                } else {
                                    if v1 {
                                        var v150 uint = (v4 & (^(v4)))
                                        var v151 float32 = 3.14159
                                        _ = v150
                                        _ = v151
                                    }
                                    var v152 int64 = 0
                                    _ = v152
                                }
                                if len(v123) >= 2 {
                                    v123 = v123[1:len(v123)]
                                }
                                var v153 string = "\n\t"
                                v4 = v4
                                _ = fi141
                                _ = v142
                                _ = v143
                                _ = v153
                            }
                        case "\n\t":
                            if ((fi139 << (((^(v4))) & 63)) != (fi139 + 0)) {
                                var v154 int32 = -1
                                _ = v154
                            } else {
                                var v155 uint32 = 65535
                                v156 := v3[470]
                                _ = v155
                                _ = v156
                            }
                        default:
                            delete(v3, fi139)
                        }
                    }
                    fmt.Println(v1, v5)
                    _ = fi139
                }
                v1 = v1
            }
        }
        _ = v123
    } else {
        if (165 != 0) {
            v157 := v3[85]
            for fi158 := 0; fi158 < 1; fi158++ {
                for fi159 := 0; fi159 < 13; fi159++ {
                    var v160 uint = (v4 - (^(v4)))
                    var v161 uint64 = 1
                    _ = fi159
                    _ = v160
                    _ = v161
                }
                v162 := []int{}
                _ = fi158
                _ = v162
            }
            _ = v157
        }
        var v163 int16 = -32768
        _ = v163
    }
    var v164 float64 = (0.1 * v2)
    var v165 string = "\n\t"
    for fi166 := 0; fi166 < 10; fi166++ {
        for fi167 := 0; fi167 < 16; fi167++ {
            var v168 uint = (v4 / ((v4) | 1))
            if (-9223372036854775808 == -1) {
                v169 := GenS6{F0: 973}
                v170 := v3[127]
                if (!(18446744073709551615 < 94)) {
                    switch ((v165 + v165) + v165) {
                    case "hello":
                        var v171 string = (v165 + v165)
                        _ = v171
                    case "\n\t":
                        v164 = (v2 * (float64(1.5) / (-2.25 + float64(((3.14159 / (float32(((v164 * v164) / v164)) - ((((float32(((((v2 / v164) * -2.25) / (3.14159 * v2)) * ((v164 / 1e10) * v164))) + (((0.0 - 1e10) - (1.5 * -2.25)) - 3.14159)) + (((1.5 * ((-2.25 / ((1.5 / 1.0) + float32((1.5 * v2)))) * 1e10)) + (1.5 + 1e10)) + (0.1 / 1.0))) / (((3.14159 * 1e10) * ((0.0 / (-2.25 / (-2.25 + 3.14159))) / (0.1 * float32((1e10 * (v164 * float64(1.5))))))) / -2.25)) - 0.1))) - float32(float64(-2.25)))))))
                    default:
                        var v172 int64 = 1
                        v173 := GenS4{F0: fi166, F1: 1000, F2: 32767}
                        _ = v172
                        _ = v173
                    }
                }
                _ = v169
                _ = v170
            } else {
                v168 += (v4 % (((v168 | (v4 * (v4 - v168)))) | 1))
            }
            if false {
                if (true && (577 != 2147483647)) {
                    v5 = genFn1()
                    if (v1 || (1.5 < ((v164 * v164) - v2))) {
                        v4 -= v4
                        var v174 int32 = 1
                        if v1 {
                            fmt.Println(v164, fi166, v1)
                        }
                        _ = v174
                    }
                    v164 += (float64(((0.1 + (3.14159 + (((((((((float32((v2 / v2)) - 0.1) + (-2.25 / 1.5)) + 0.0) - ((1.5 / (3.14159 + float32(0.1))) + float32(1.5))) * 1.5) * 0.1) - ((1e10 + (1e10 - 1.5)) * 0.1)) + -2.25) - 0.0))) - ((((((0.1 / 1e10) * -2.25) - float32(v164)) - (-2.25 * (1.5 * 3.14159))) + 0.1) - 1.5))) / 1.0)
                }
            } else {
                if false {
                    v175 := v3[fi167]
                    _ = v175
                } else {
                    v176 := GenS5{F0: 521, F1: v1}
                    _ = v176
                }
                v5 -= v5
                v177 := GenS14{F0: fi167, F1: -1}
                _ = v177
            }
            var v178 bool = (11 <= genFn0(fi167, v1))
            _ = fi167
            _ = v168
            _ = v178
        }
        fmt.Println(v165, v1)
        _ = fi166
    }
    fmt.Println("v1:", v1)
    fmt.Println("v164:", v164)
    fmt.Println("v165:", v165)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v164
    _ = v165
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
