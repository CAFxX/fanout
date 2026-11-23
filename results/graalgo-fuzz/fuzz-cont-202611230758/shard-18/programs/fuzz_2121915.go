package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 int32; F1 int; F2 string }
type GenS3 struct { F0 uint32; F1 int8 }
type GenS4 struct { F0 int; F1 int32 }
type GenS5 struct { F0 int8; F1 int32; F2 int16 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint64; F1 int8 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int; F1 int }
type GenS10 struct { F0 uint32; F1 uint8; F2 int }
type GenS11 struct { F0 int64; F1 uint }
func genFn0(p0 int) bool {
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 uint8 = 1
        _ = fi6
        _ = v7
    }
    fmt.Println(p0)
    fmt.Println(p0)
    if (!true) {
        fmt.Println(p0)
        fmt.Println(p0)
    } else {
        p0 += -9223372036854775808
        var v8 int16 = 32767
        _ = v8
    }
    p0 = p0
    _ = p0
    return ((p0 | (p0 << ((0) & 63))) > 9223372036854775807)
}
func genFn1(p0 string, p1 string) bool {
    v9 := GenS1{F0: -667, F1: -2147483648}
    fmt.Println(p0, p0)
    p1 = p0
    for fi10 := 0; fi10 < 14; fi10++ {
        var v11 int8 = -128
        var v12 uint64 = 291
        _ = fi10
        _ = v11
        _ = v12
    }
    if true {
        fmt.Println(p1, p1)
    } else {
        var v13 int8 = -1
        for fi14 := 0; fi14 < 14; fi14++ {
            var v15 uint8 = 212
            v16 := GenS1{F0: -1, F1: -1}
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v9
    return false
}

func main() {
    v1 := make(map[int]int)
    var v2 bool = true
    var v3 uint16 = 42
    var v4 float64 = 3.14159
    var v5 uint32 = 2147483647
    v4 += (v4 / float64(1e10))
    if true {
        v17 := []int{}
        var v18 int16 = -32768
        _ = v17
        _ = v18
    }
    v19 := GenS1{F0: -2147483648, F1: 0}
    for fi20 := 0; fi20 < 15; fi20++ {
        v21 := GenS2{F0: 2147483647, F1: 868, F2: "\n\t"}
        _ = fi20
        _ = v21
    }
    v3 += uint16(v5)
    fmt.Println(v3, v5)
    var v22 uint32 = 42
    for fi23 := 0; fi23 < 9; fi23++ {
        var v24 int32 = 42
        for fi25 := 0; fi25 < 3; fi25++ {
            if false {
                switch "a" {
                case "foo bar":
                    if (!v2) {
                        var v26 int = fi25
                        v27 := GenS2{F0: v24, F1: (^(fi23)), F2: "\n\t"}
                        _ = v26
                        _ = v27
                    } else {
                        v22 += 65535
                        v22 -= v5
                    }
                case "":
                    v3 -= v3
                case "\n\t":
                    fmt.Println(v24, fi23, v5)
                default:
                    var v28 int32 = -1
                    _ = v28
                }
                v29 := GenS1{F0: (v24 * -1), F1: v24}
                _ = v29
            } else {
                switch "hello" {
                case "\n\t":
                    fmt.Println(v3, v5)
                }
                v4 = float64((float32(0.1) / 1.0))
            }
            _ = fi25
        }
        _ = fi23
        _ = v24
    }
    v30 := GenS1{F0: 0, F1: 1}
    var v31 bool = (v2 || v2)
    switch ("" + "hello") {
    case "foo bar":
        if true {
            var v32 float64 = float64(float32(float64((3.14159 - 0.1))))
            fmt.Println(v31, v2)
            v22 -= (v22 & v22)
            _ = v32
        }
    case "a":
        _ = len(v1)
    default:
        v33 := []int{-2147483648, -656, 9223372036854775807}
        _ = v33
    }
    for fi34 := 0; fi34 < 17; fi34++ {
        for fi35 := 0; fi35 < 12; fi35++ {
            var v36 int16 = 32767
            switch fi35 {
            case -489:
                if v31 {
                    switch (("" + "a") + ("a" + ("" + "\n\t"))) {
                    case "hello":
                        v31 = v31
                    case "a":
                        fmt.Println(fi34, v36, v5)
                        var v37 int32 = 42
                        _ = v37
                    default:
                        var v38 int64 = -1
                        _ = v38
                    }
                    var v39 uint32 = (v22 >> ((890) & 63))
                    v40 := GenS2{F0: 0, F1: fi34, F2: "foo bar"}
                    _ = v39
                    _ = v40
                } else {
                    var v41 float64 = v4
                    _ = v41
                }
            case 2147483647:
                var v42 int16 = 731
                _ = v42
            case 32767:
                if (v31 && (82 >= 42)) {
                    v2 = v31
                    if false {
                        v31 = v2
                        fmt.Println(v2, v3)
                        switch "a" {
                        case "\n\t":
                            v43 := GenS1{F0: -2147483648, F1: 0}
                            _ = v43
                        case "hello":
                            var v44 float64 = float64((0.0 + 0.1))
                            _ = v44
                        }
                    } else {
                        v45 := []int{(^(fi35))}
                        var v46 float64 = (float64(float32(1.5)) / float64(0.1))
                        if v2 {
                            v47 := GenS2{F0: 10, F1: 1, F2: "\n\t"}
                            switch fi35 {
                            case 32767:
                                v48 := v1[fi35]
                                _ = v48
                            case -2147483648:
                                fmt.Println(v2, v2, v2)
                            case 32768:
                                fmt.Println(v46, fi34, fi34)
                            default:
                                v49 := GenS1{F0: -1, F1: 2147483647}
                                _ = v49
                            }
                            _ = v47
                        }
                        _ = v45
                        _ = v46
                    }
                }
            }
            _ = fi35
            _ = v36
        }
        _ = fi34
    }
    if (340 > -2147483648) {
        var v50 uint32 = v22
        if v31 {
            var v51 int8 = 120
            _ = v51
        }
        v52 := []int{-9223372036854775808}
        _ = v50
        _ = v52
    } else {
        switch 32768 {
        case 2147483647:
            if v31 {
                v3 -= (v3 * v3)
                for fi53 := 0; fi53 < 12; fi53++ {
                    for fi54 := 0; fi54 < 20; fi54++ {
                        if v2 {
                            if (!(v31 && v2)) {
                                if v31 {
                                    v55 := []int{int(v22), -794}
                                    if true {
                                        v1[fi53] = (fi53 / (((fi54 * fi54)) | 1))
                                        if false {
                                            if (9223372036854775807 <= 1) {
                                                v31 = (!(v31 && (!true)))
                                                _ = len(v55)
                                                if v31 {
                                                    var v56 string = "hello"
                                                    fmt.Println(v31, fi53, v31)
                                                    _ = v56
                                                } else {
                                                    if false {
                                                        if (306 != 194) {
                                                            v57 := GenS3{F0: v22, F1: -128}
                                                            var v58 uint64 = 647
                                                            _ = v57
                                                            _ = v58
                                                        }
                                                        switch (fi54 * fi54) {
                                                        case 1:
                                                            fmt.Println(v22, v4, v2)
                                                        case 647:
                                                            fmt.Println(fi54, fi54)
                                                        }
                                                    }
                                                }
                                            }
                                            v59 := GenS1{F0: -792, F1: 2147483647}
                                            var v60 int16 = 630
                                            fmt.Println(v60, v5)
                                            _ = v59
                                            _ = v60
                                        } else {
                                            v61 := GenS4{F0: fi53, F1: -1}
                                            _ = v61
                                        }
                                    } else {
                                        v31 = false
                                        if v31 {
                                            v62 := GenS5{F0: -79, F1: 1, F2: 32767}
                                            _ = v62
                                        } else {
                                            var v63 uint = 950
                                            _ = v63
                                        }
                                    }
                                    _ = v55
                                }
                                var v64 uint64 = 0
                                v65 := GenS5{F0: -1, F1: 0, F2: 1}
                                var v66 uint64 = (v64 / (((v64 & v64)) | 1))
                                switch "foo bar" {
                                case "":
                                    var v67 float32 = ((0.1 * ((0.1 + (3.14159 + (((float32((v4 * ((float64(float32((-2.25 / 1.0))) - ((-2.25 - v4) - v4)) / v4))) - 1.5) / 0.1) - float32((v4 + float64(1e10)))))) + ((float32((0.1 / v4)) + (float32(1.5) + 1.5)) + float32(-2.25)))) * (0.0 - 1.5))
                                    _ = v67
                                }
                                _ = v64
                                _ = v65
                                _ = v66
                            } else {
                                v22 = v5
                            }
                            v31 = false
                        }
                        v2 = v31
                        var v68 bool = v2
                        _ = fi54
                        _ = v68
                    }
                    _ = fi53
                }
            }
        case -860:
            fmt.Println(v2, v5)
        default:
            v22 += (^(v22))
        }
        switch -2147483648 {
        case -9223372036854775808:
            for fi69 := 0; fi69 < 10; fi69++ {
                if false {
                    v3 = 669
                    for fi70 := 0; fi70 < 19; fi70++ {
                        v5 += v5
                        if false {
                            v5 = (v5 * (v5 & v5))
                        } else {
                            switch (fi69 ^ 823) {
                            case 127:
                                v71 := GenS2{F0: 0, F1: (fi69 >> ((42) & 63)), F2: ("hello" + ("" + "foo bar"))}
                                _ = v71
                            case 32767:
                                var v72 int8 = 0
                                _ = v72
                            case 32768:
                                v73 := GenS5{F0: 0, F1: 42, F2: -1}
                                _ = v73
                            }
                            if v31 {
                                var v74 uint32 = (v5 % ((65535) | 1))
                                _ = v74
                            }
                            if v2 {
                                if v31 {
                                    v22 = (v22 >> ((411) & 63))
                                    v1[fi70] = (fi70 + (fi69 & (-(fi69))))
                                } else {
                                    v22 += (v22 * (v5 << ((42) & 63)))
                                    v4 += v4
                                    v5 += uint32(v3)
                                    switch "a" {
                                    case "a":
                                        var v75 uint64 = 18446744073709551615
                                        var v76 int16 = -1
                                        _ = v75
                                        _ = v76
                                    case "hello":
                                        if true {
                                            var v77 int16 = 0
                                            v78 := []int{(-(fi70)), (-(fi69))}
                                            if (1 < 2) {
                                                var v79 bool = v2
                                                switch "hello" {
                                                case "":
                                                    var v80 int16 = (-(v77))
                                                    _ = v80
                                                case "hello":
                                                    var v81 string = ("foo bar" + "a")
                                                    switch v81 {
                                                    case "foo bar":
                                                        delete(v1, fi70)
                                                    case "a":
                                                        var v82 uint16 = 99
                                                        _ = v82
                                                    }
                                                    _ = v81
                                                case "foo bar":
                                                    v5 = (v5 * (v5 | 800))
                                                default:
                                                    fmt.Println(v22, v31, v22)
                                                }
                                                _ = v79
                                            } else {
                                                var v83 int8 = 1
                                                var v84 float64 = ((v4 * v4) + float64(0.1))
                                                _ = v83
                                                _ = v84
                                            }
                                            _ = v77
                                            _ = v78
                                        }
                                    }
                                }
                                _ = len(v1)
                                var v85 uint32 = (v22 * v5)
                                _ = v85
                            }
                        }
                        _ = fi70
                    }
                } else {
                    var v86 int = (fi69 << ((384) & 63))
                    _ = v86
                }
                v5 += v22
                switch "a" {
                case "foo bar":
                    switch ("a" + ("" + "\n\t")) {
                    case "foo bar":
                        var v87 int64 = -178
                        _ = v87
                    case "\n\t":
                        v88 := GenS5{F0: 74, F1: 1, F2: -823}
                        _ = v88
                    case "hello":
                        var v89 int32 = 0
                        _ = v89
                    }
                }
                _ = fi69
            }
        case 2147483647:
            _ = len(v1)
        default:
            if false {
                switch -1 {
                case 1:
                    v90 := []int{128}
                    _ = v90
                case 127:
                    if v2 {
                        if v31 {
                            v91 := v1[-2147483648]
                            var v92 uint = 65535
                            _ = v91
                            _ = v92
                        }
                        v3 = (v3 & (v3 * v3))
                        v93 := []int{59}
                        _ = v93
                    }
                }
            }
        }
        for fi94 := 0; fi94 < 2; fi94++ {
            if true {
                switch fi94 {
                case -9223372036854775808:
                    var v95 int = fi94
                    _ = v95
                case 14:
                    for fi96 := 0; fi96 < 4; fi96++ {
                        v31 = (!true)
                        switch (-(fi96)) {
                        case 127:
                            if v2 {
                                switch "a" {
                                case "":
                                    v97 := GenS3{F0: (v5 + (v22 - (v5 << ((65535) & 63)))), F1: 4}
                                    _ = v97
                                default:
                                    v3 = (^(v3))
                                    v31 = (!((398 < 4294967295) && true))
                                }
                                if false {
                                    v3 += (v3 << ((65535) & 63))
                                }
                            } else {
                                var v98 uint64 = 637
                                var v99 uint64 = 1
                                if false {
                                    v2 = false
                                    v99 -= (v99 | (v99 << ((1) & 63)))
                                }
                                _ = v98
                                _ = v99
                            }
                        case -9223372036854775808:
                            v3 = v3
                        case -2147483648:
                            var v100 int16 = 179
                            _ = v100
                        default:
                            v101 := []int{42, fi96, int(v3)}
                            _ = v101
                        }
                        _ = fi96
                    }
                default:
                    if (-38 != 1) {
                        for fi102 := 0; fi102 < 11; fi102++ {
                            v103 := GenS6{F0: 4294967295}
                            var v104 float64 = 0.0
                            fmt.Println(v31, v5)
                            _ = fi102
                            _ = v103
                            _ = v104
                        }
                        if true {
                            for fi105 := 0; fi105 < 15; fi105++ {
                                v106 := GenS6{F0: 197}
                                _ = fi105
                                _ = v106
                            }
                            v107 := v1[(fi94 >> ((275) & 63))]
                            _ = v107
                        }
                        if v31 {
                            v5 = 0
                            var v108 int32 = 0
                            _ = v108
                        } else {
                            fmt.Println(v3, v31, v31)
                            v2 = (519 > 0)
                        }
                    } else {
                        v109 := []int{fi94, (fi94 * int(v3))}
                        if (!v2) {
                            v110 := GenS7{F0: 42, F1: -9}
                            v109 = append(v109, (fi94 * fi94))
                            _ = v110
                        }
                        _ = v109
                    }
                }
                v4 = (v4 - v4)
            } else {
                if v2 {
                    if (v2 || (!v31)) {
                        for fi111 := 0; fi111 < 17; fi111++ {
                            v112 := []int{2147483647}
                            _ = fi111
                            _ = v112
                        }
                        switch "foo bar" {
                        case "":
                            v5 += 2147483648
                        case "hello":
                            var v113 string = "\n\t"
                            _ = v113
                        case "foo bar":
                            var v114 uint32 = (^(v22))
                            _ = v114
                        default:
                            switch "a" {
                            case "foo bar":
                                switch "hello" {
                                case "\n\t":
                                    for fi115 := 0; fi115 < 8; fi115++ {
                                        v4 = (v4 + (v4 - 0.1))
                                        v116 := []int{(fi94 % (((fi94 + 1)) | 1)), fi94, (fi94 / (((fi94 / ((fi115) | 1))) | 1)), (fi115 ^ (fi94 << ((0) & 63)))}
                                        _ = fi115
                                        _ = v116
                                    }
                                case "foo bar":
                                    for fi117 := 0; fi117 < 18; fi117++ {
                                        v4 -= ((v4 * (0.1 / (v4 * v4))) + (1e10 - float64(1.5)))
                                        if v31 {
                                            switch (fi94 * (fi117 + fi117)) {
                                            case 862:
                                                var v118 uint16 = 255
                                                _ = v118
                                            default:
                                                if v31 {
                                                    var v119 bool = (32767 <= v3)
                                                    v120 := GenS5{F0: 127, F1: 0, F2: 42}
                                                    _ = v119
                                                    _ = v120
                                                } else {
                                                    v121 := GenS5{F0: 42, F1: 66, F2: 32767}
                                                    if (float64((3.14159 / 1.5)) == ((v4 * v4) / 1.5)) {
                                                        v31 = v2
                                                        v31 = v31
                                                        v4 -= v4
                                                    }
                                                    var v122 uint8 = 247
                                                    _ = v121
                                                    _ = v122
                                                }
                                            }
                                        } else {
                                            v3 += 0
                                            var v123 uint16 = v3
                                            _ = v123
                                        }
                                        var v124 bool = ((v4 * 0.0) != v4)
                                        _ = fi117
                                        _ = v124
                                    }
                                }
                            case "\n\t":
                                if (((((3.14159 * (float32(((1.5 * v4) / 1.0)) * (((-2.25 + 3.14159) + ((((float32((3.14159 / 0.1)) / (((((((float32((1.5 + (0.1 + 3.14159))) + 0.1) + 0.1) - (float32((1.5 * 0.0)) / (0.0 / (float32(v4) - 3.14159)))) - 3.14159) - 1e10) * (0.1 / 1.5)) - ((1.5 * 1e10) * 1.5))) / (float32(float64(3.14159)) * (1.5 - 0.1))) - -2.25) * 0.0)) + -2.25))) * float32(v4)) + (((1.5 * 1.5) / 1e10) - 3.14159)) + (1e10 - 0.1)) >= -2.25) {
                                    if v2 {
                                        for fi125 := 0; fi125 < 17; fi125++ {
                                            var v126 int = fi94
                                            var v127 uint32 = (v5 & (v5 | uint32(v3)))
                                            if false {
                                                var v128 uint32 = 421
                                                v129 := GenS8{F0: v31}
                                                _ = v128
                                                _ = v129
                                            }
                                            _ = fi125
                                            _ = v126
                                            _ = v127
                                        }
                                    }
                                    if ((!v2) || v2) {
                                        v2 = (791 < 42)
                                        fmt.Println(v22, v4)
                                        v130 := GenS2{F0: 1, F1: 2147483647, F2: "a"}
                                        _ = v130
                                    } else {
                                        v22 += (v5 % ((65535) | 1))
                                        switch 516 {
                                        case 9223372036854775807:
                                            switch "foo bar" {
                                            case "a":
                                                v22 -= 1
                                            default:
                                                v2 = v2
                                            }
                                        }
                                    }
                                } else {
                                    v131 := GenS7{F0: 1, F1: 42}
                                    _ = v131
                                }
                            case "":
                                v22 = 2147483647
                            default:
                                v1[fi94] = (^(fi94))
                            }
                        }
                    } else {
                        for fi132 := 0; fi132 < 17; fi132++ {
                            v133 := GenS9{F0: (-(fi94)), F1: fi94}
                            if v2 {
                                delete(v1, 4294967295)
                                var v134 int8 = 1
                                switch (fi132 / ((fi94) | 1)) {
                                case 32767:
                                    var v135 float32 = -2.25
                                    _ = v135
                                }
                                var v136 int32 = -2147483648
                                v1[fi132] = (fi94 % ((-1) | 1))
                                _ = v134
                                _ = v136
                            } else {
                                if (fi132 > fi94) {
                                    var v137 uint32 = (^(v5))
                                    v2 = false
                                    _ = v137
                                }
                                var v138 uint64 = 4294967295
                                _ = v138
                            }
                            fmt.Println(v5)
                            var v139 float64 = v4
                            if v31 {
                                var v140 uint32 = 65535
                                if false {
                                    v141 := GenS4{F0: (^(fi94)), F1: 42}
                                    _ = v141
                                }
                                v142 := GenS9{F0: int(v5), F1: (fi94 % (((fi94 + fi94)) | 1))}
                                _ = v140
                                _ = v142
                            } else {
                                var v143 uint64 = 0
                                var v144 int16 = -896
                                _ = v143
                                _ = v144
                            }
                            _ = fi132
                            _ = v133
                            _ = v139
                        }
                        v145 := v1[fi94]
                        _ = v145
                    }
                    fmt.Println(v4, v4, v31)
                    delete(v1, (fi94 | (fi94 ^ fi94)))
                }
                for fi146 := 0; fi146 < 2; fi146++ {
                    var v147 float64 = v4
                    _ = fi146
                    _ = v147
                }
                switch ("hello" + "\n\t") {
                case "":
                    if (127 >= -558) {
                        if true {
                            v148 := []int{}
                            _ = v148
                        } else {
                            var v149 int32 = 42
                            _ = v149
                        }
                    } else {
                        switch (fi94 & fi94) {
                        case 9223372036854775807:
                            v5 = 4294967295
                        case -9223372036854775808:
                            v150 := GenS10{F0: 1, F1: 193, F2: fi94}
                            _ = v150
                        }
                        var v151 int = -999
                        v152 := []int{}
                        _ = v151
                        _ = v152
                    }
                case "\n\t":
                    v153 := GenS11{F0: 0, F1: 0}
                    _ = v153
                case "foo bar":
                    v5 = v5
                }
            }
            v5 += 2
            switch -1 {
            case 42:
                for fi154 := 0; fi154 < 11; fi154++ {
                    if true {
                        switch ("" + "\n\t") {
                        case "\n\t":
                            v31 = (-1 >= 127)
                        }
                        fmt.Println(fi94)
                        var v155 uint = 4294967295
                        _ = v155
                    }
                    _ = len(v1)
                    var v156 uint16 = v3
                    _ = fi154
                    _ = v156
                }
            }
            _ = fi94
        }
    }
    var v157 int32 = -1
    if true {
        v158 := []int{}
        v159 := GenS11{F0: 1, F1: 181}
        _ = v158
        _ = v159
    } else {
        var v160 uint32 = v5
        _ = v160
    }
    for fi161 := 0; fi161 < 10; fi161++ {
        v2 = (-3 <= 622)
        _ = fi161
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v157:", v157)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v157
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
