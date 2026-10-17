package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 uint; F2 int32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32; F1 uint64; F2 uint16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint64; F1 int32 }
type GenS7 struct { F0 uint32; F1 int64; F2 uint16 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 string; F1 bool; F2 uint64 }
type GenS10 struct { F0 int64; F1 int8; F2 int8 }
type GenS11 struct { F0 uint; F1 uint32; F2 int32 }
type GenS12 struct { F0 uint8; F1 uint32; F2 int }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 uint32; F1 int16; F2 int16 }
func genFn0(p0 uint64) uint16 {
    fmt.Println(p0)
    var v4 int32 = -1
    if (false && true) {
        var v5 uint64 = 4294967295
        v4 = 42
        if (!(302 == -1)) {
            fmt.Println(p0)
            if (false && false) {
                v6 := GenS1{F0: 930}
                v4 = (v4 << ((67) & 63))
                switch (("hello" + "") + "a") {
                case "\n\t":
                    var v7 uint = 4294967295
                    _ = v7
                case "":
                    for fi8 := 0; fi8 < 15; fi8++ {
                        var v9 uint16 = 1
                        if false {
                            fmt.Println(v9, v5, v5)
                            if false {
                                var v10 float32 = -2.25
                                _ = v10
                            } else {
                                var v11 uint32 = 2
                                _ = v11
                            }
                        }
                        v12 := GenS1{F0: (fi8 % ((fi8) | 1))}
                        _ = fi8
                        _ = v12
                        _ = v9
                    }
                default:
                    v13 := []int{-1, -475, 32768, 127}
                    _ = v13
                }
                _ = v6
            } else {
                fmt.Println(v4, p0, v4)
                switch ((("\n\t" + "") + "hello") + "") {
                case "a":
                    for fi14 := 0; fi14 < 11; fi14++ {
                        var v15 uint8 = 210
                        var v16 int8 = 42
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                case "foo bar":
                    var v17 uint32 = 1
                    _ = v17
                case "":
                    v4 = v4
                default:
                    var v18 uint64 = p0
                    _ = v18
                }
                var v19 int16 = -395
                _ = v19
            }
            if false {
                v4 += int32(v5)
                for fi20 := 0; fi20 < 16; fi20++ {
                    switch ("" + (("a" + "foo bar") + "hello")) {
                    case "hello":
                        p0 += v5
                    case "foo bar":
                        if true {
                            fmt.Println(p0, v5, fi20)
                            var v21 uint8 = 72
                            var v22 int8 = -127
                            _ = v21
                            _ = v22
                        } else {
                            fmt.Println(p0)
                        }
                    }
                    _ = fi20
                }
            } else {
                v23 := []int{}
                if len(v23) >= 2 {
                    v23 = v23[1:len(v23)]
                }
                _ = v23
            }
        }
        _ = v5
    } else {
        v4 += v4
        var v24 uint8 = 166
        _ = v24
    }
    _ = p0
    _ = v4
    return 2
}
func genFn1() string {
    if ((32767 < 42) || (-2147483648 > -2147483648)) {
        var v25 uint8 = 1
        v25 = (v25 | v25)
        v25 += v25
        _ = v25
    } else {
        if (-128 == 0) {
            var v26 int = -1
            if true {
                var v27 int16 = 42
                if true {
                    switch ("foo bar" + "") {
                    case "hello":
                        v26 = v26
                    case "a":
                        if false {
                            var v28 int64 = -348
                            _ = v28
                        }
                    default:
                        v29 := GenS2{F0: -1, F1: 103, F2: -894}
                        _ = v29
                    }
                    var v30 int32 = 0
                    v31 := []int{(v26 >> ((302) & 63))}
                    _ = v30
                    _ = v31
                } else {
                    var v32 float32 = (-2.25 - -2.25)
                    v33 := []int{v26}
                    for fi34 := 0; fi34 < 15; fi34++ {
                        if len(v33) > 0 {
                            v33[0] = 0
                        }
                        var v35 float32 = v32
                        if (199 == 114) {
                            if len(v33) >= 2 {
                                v33 = v33[1:len(v33)]
                            }
                            v36 := GenS2{F0: 0, F1: 65535, F2: -487}
                            _ = v36
                        }
                        _ = fi34
                        _ = v35
                    }
                    _ = v32
                    _ = v33
                }
                _ = v27
            } else {
                var v37 float64 = (0.1 - ((0.0 / -2.25) - (0.0 + 1e10)))
                _ = v37
            }
            _ = v26
        } else {
            var v38 int64 = 9223372036854775807
            v38 += (v38 & (v38 & -137))
            _ = v38
        }
        var v39 float32 = 0.0
        _ = v39
    }
    if false {
        switch ("hello" + (("hello" + "\n\t") + ("hello" + "foo bar"))) {
        case "":
            switch 638 {
            case 127:
                if (("foo bar" + "a") != "") {
                    switch "foo bar" {
                    case "\n\t":
                        for fi40 := 0; fi40 < 3; fi40++ {
                            for fi41 := 0; fi41 < 6; fi41++ {
                                var v42 int32 = -2147483648
                                switch (("" + ("foo bar" + "hello")) + "\n\t") {
                                case "":
                                    var v43 int64 = -9223372036854775808
                                    _ = v43
                                case "\n\t":
                                    v44 := []int{-1, -352, fi40, 4294967295}
                                    _ = v44
                                case "hello":
                                    v42 = (-(v42))
                                default:
                                    var v45 bool = false
                                    v46 := []int{}
                                    fmt.Println(v42, fi41, fi40)
                                    switch ("" + (("hello" + ("a" + "a")) + "")) {
                                    case "foo bar":
                                        if (("\n\t" <= ("" + "foo bar")) && v45) {
                                            v42 = 10
                                            v42 = -1
                                        }
                                    case "hello":
                                        var v47 int32 = 2147483647
                                        _ = v47
                                    case "\n\t":
                                        var v48 int = 128
                                        _ = v48
                                    default:
                                        var v49 string = ("hello" + "")
                                        _ = v49
                                    }
                                    fmt.Println(v45)
                                    _ = v45
                                    _ = v46
                                }
                                var v50 float32 = (0.0 + -2.25)
                                _ = fi41
                                _ = v42
                                _ = v50
                            }
                            _ = fi40
                        }
                    case "hello":
                        for fi51 := 0; fi51 < 9; fi51++ {
                            var v52 bool = (1 >= 1)
                            _ = fi51
                            _ = v52
                        }
                    case "foo bar":
                        var v53 bool = (-914 < 518)
                        _ = v53
                    }
                    for fi54 := 0; fi54 < 7; fi54++ {
                        fmt.Println(fi54)
                        var v55 uint64 = 43
                        _ = fi54
                        _ = v55
                    }
                }
            case 204:
                var v56 uint64 = 42
                _ = v56
            case 224:
                var v57 int8 = 1
                _ = v57
            default:
                switch (("hello" + "\n\t") + "hello") {
                case "foo bar":
                    switch 479 {
                    case 55:
                        var v58 int32 = 195
                        _ = v58
                    }
                case "":
                    for fi59 := 0; fi59 < 6; fi59++ {
                        v60 := GenS2{F0: 411, F1: 1, F2: 762}
                        if (((0.0 - 0.0) + 3.14159) >= -2.25) {
                            if ((!true) || (false || (2 != 1))) {
                                var v61 uint64 = 490
                                var v62 float32 = (0.0 + -2.25)
                                _ = v61
                                _ = v62
                            } else {
                                for fi63 := 0; fi63 < 11; fi63++ {
                                    if ((-2.25 * 1e10) > 3.14159) {
                                        var v64 float64 = (-2.25 * 1e10)
                                        v64 = 0.0
                                        _ = v64
                                    }
                                    _ = fi63
                                }
                                if true {
                                    fmt.Println(fi59)
                                    var v65 uint8 = 128
                                    _ = v65
                                }
                                switch ("a" + "") {
                                case "foo bar":
                                    for fi66 := 0; fi66 < 17; fi66++ {
                                        var v67 uint16 = 42
                                        v68 := GenS3{F0: 65}
                                        v67 += v67
                                        v69 := GenS4{F0: 797, F1: 4294967295, F2: (v67 & v67)}
                                        if false {
                                            fmt.Println(fi59, fi59)
                                            v70 := GenS3{F0: 0}
                                            v71 := []int{(fi59 >> ((956) & 63)), (fi66 % (((-(fi66))) | 1))}
                                            _ = v70
                                            _ = v71
                                        }
                                        v72 := []int{(fi59 << ((0) & 63))}
                                        _ = fi66
                                        _ = v67
                                        _ = v68
                                        _ = v69
                                        _ = v72
                                    }
                                default:
                                    v73 := []int{fi59, 4294967295, 32768}
                                    _ = v73
                                }
                            }
                            fmt.Println(fi59)
                            var v74 uint8 = 188
                            _ = v74
                        } else {
                            var v75 int = -531
                            fmt.Println(fi59, fi59)
                            _ = v75
                        }
                        if ((459 >= -1) || false) {
                            if (!(!(!(42 <= -536)))) {
                                var v76 bool = true
                                _ = v76
                            }
                            fmt.Println(fi59)
                            v77 := GenS2{F0: -9223372036854775808, F1: 1, F2: -2147483648}
                            _ = v77
                        } else {
                            var v78 uint = 255
                            _ = v78
                        }
                        _ = fi59
                        _ = v60
                    }
                case "a":
                    if false {
                        var v79 int32 = 12
                        _ = v79
                    } else {
                        var v80 bool = (533 == 413)
                        _ = v80
                    }
                }
            }
        case "a":
            var v81 float64 = -2.25
            _ = v81
        case "foo bar":
            var v82 bool = (true || (1 != 255))
            _ = v82
        }
    } else {
        switch "foo bar" {
        case "\n\t":
            for fi83 := 0; fi83 < 12; fi83++ {
                var v84 uint16 = genFn0(2)
                if (-2147483648 <= 42) {
                    var v85 bool = false
                    _ = v85
                }
                if (2147483647 < 2147483647) {
                    fmt.Println(v84)
                }
                _ = fi83
                _ = v84
            }
        }
        var v86 int = -137
        v86 = v86
        _ = v86
    }
    var v87 float32 = (-2.25 - -2.25)
    _ = v87
    return "foo bar"
}

func main() {
    var v1 int16 = 32767
    var v2 uint64 = 2
    var v3 uint64 = v2
    v88 := []int{2147483647, -790}
    if len(v88) > 0 {
        v88[0] = 2147483647
    }
    v89 := GenS5{F0: -32768}
    v90 := GenS4{F0: 2147483647, F1: (v3 & 372), F2: 0}
    var v91 uint = 804
    var v92 int64 = -1
    var v93 uint16 = 0
    v94 := GenS1{F0: -561}
    switch 9223372036854775807 {
    case -302:
        for fi95 := 0; fi95 < 13; fi95++ {
            switch "a" {
            case "a":
                for fi96 := 0; fi96 < 14; fi96++ {
                    var v97 int16 = (-(v1))
                    switch (fi96 / ((int(v1)) | 1)) {
                    case 1:
                        v98 := GenS6{F0: (v2 >> ((v91) & 63)), F1: 0}
                        switch "" {
                        case "":
                            var v99 int32 = 1
                            _ = v99
                        case "\n\t":
                            if len(v88) > 0 {
                                v88[0] = fi96
                            }
                        }
                        v1 = 42
                        _ = v98
                    default:
                        var v100 int16 = (v1 / (((v1 / ((v97) | 1))) | 1))
                        _ = v100
                    }
                    _ = fi96
                    _ = v97
                }
            case "foo bar":
                fmt.Println(v93, v93, v1)
            case "\n\t":
                switch "\n\t" {
                case "":
                    v91 = (v91 ^ (^(v91)))
                case "foo bar":
                    v2 += (v2 >> ((v91) & 63))
                case "\n\t":
                    var v101 uint8 = 127
                    _ = v101
                default:
                    var v102 uint64 = v2
                    _ = v102
                }
            }
            v1 += (v1 / ((v1) | 1))
            _ = fi95
        }
    case -9223372036854775808:
        var v103 uint32 = 80
        _ = v103
    default:
        if false {
            if (true && (true || true)) {
                var v104 int16 = -759
                switch 42 {
                case 128:
                    var v105 int = 32768
                    _ = v105
                }
                _ = v104
            }
            var v106 float32 = float32((-2.25 * (((-2.25 - 3.14159) * (0.0 - (0.1 - 3.14159))) * 1.5)))
            var v107 string = ""
            _ = v106
            _ = v107
        } else {
            if false {
                if true {
                    v108 := GenS6{F0: v2, F1: -2147483648}
                    var v109 uint32 = 2147483647
                    var v110 uint32 = v109
                    _ = v108
                    _ = v109
                    _ = v110
                }
                v93 = (v93 + v93)
            } else {
                var v111 int8 = 99
                var v112 uint64 = 42
                for fi113 := 0; fi113 < 20; fi113++ {
                    fmt.Println(v111)
                    for fi114 := 0; fi114 < 13; fi114++ {
                        switch 746 {
                        case -2147483648:
                            v92 -= v92
                        default:
                            v115 := GenS7{F0: 65535, F1: -9223372036854775808, F2: v93}
                            _ = v115
                        }
                        switch (fi114 % ((fi114) | 1)) {
                        case 128:
                            v91 -= v91
                        default:
                            v116 := GenS8{F0: ("a" + "foo bar")}
                            switch 127 {
                            case 2147483647:
                                fmt.Println(v112, fi113)
                            case 127:
                                switch ((("hello" + "hello") + "a") + "hello") {
                                case "hello":
                                    if false {
                                        v117 := GenS3{F0: 128}
                                        var v118 uint32 = 2
                                        v92 = int64(v118)
                                        _ = v117
                                        _ = v118
                                    } else {
                                        v119 := GenS9{F0: ("a" + ""), F1: (v93 <= genFn0(4294967295)), F2: 0}
                                        var v120 int32 = int32(v91)
                                        if (false || (0.1 < (1.5 / 1.0))) {
                                            if false {
                                                var v121 int8 = (v111 << (((v91 + (^(v91)))) & 63))
                                                _ = v121
                                            } else {
                                                if (1 == 0) {
                                                    v120 -= (v120 | (v120 % ((v120) | 1)))
                                                    if true {
                                                        if (fi114 >= (fi113 & (fi113 | (^(fi113))))) {
                                                            switch "hello" {
                                                            case "hello":
                                                                v91 = 923
                                                            case "foo bar":
                                                                fmt.Println(v120, v3, v3)
                                                            case "a":
                                                                if false {
                                                                    switch "\n\t" {
                                                                    case "\n\t":
                                                                        fmt.Println(fi113, v111)
                                                                    case "foo bar":
                                                                        var v122 uint64 = (v2 & 2)
                                                                        _ = v122
                                                                    case "hello":
                                                                        if (!(1 != (^(v2)))) {
                                                                            var v123 bool = true
                                                                            var v124 uint = (v91 + v91)
                                                                            var v125 float64 = float64(float32(1.5))
                                                                            _ = v123
                                                                            _ = v124
                                                                            _ = v125
                                                                        }
                                                                    default:
                                                                        var v126 bool = (!false)
                                                                        var v127 int16 = -32768
                                                                        v128 := GenS10{F0: v92, F1: (v111 | v111), F2: (v111 + v111)}
                                                                        _ = v126
                                                                        _ = v127
                                                                        _ = v128
                                                                    }
                                                                }
                                                            }
                                                            var v129 bool = true
                                                            _ = v129
                                                        } else {
                                                            v88 = append(v88, fi114)
                                                            var v130 int64 = v92
                                                            var v131 uint8 = 230
                                                            _ = v130
                                                            _ = v131
                                                        }
                                                    } else {
                                                        v112 = 18446744073709551615
                                                        fmt.Println(v112, v3)
                                                        v120 = v120
                                                    }
                                                } else {
                                                    if (!((v120 != (v120 * v120)) && (9223372036854775807 == fi113))) {
                                                        var v132 string = ("" + ("" + ("foo bar" + "a")))
                                                        _ = v132
                                                    }
                                                }
                                                if false {
                                                    fmt.Println(v112)
                                                } else {
                                                    var v133 int16 = (^(v1))
                                                    v133 = 1
                                                    if (0.1 == float32((3.14159 - (1e10 + -2.25)))) {
                                                        var v134 int64 = v92
                                                        v93 -= (v93 << ((v91) & 63))
                                                        v135 := GenS11{F0: (v91 / ((v91) | 1)), F1: 2147483648, F2: 0}
                                                        _ = v134
                                                        _ = v135
                                                    }
                                                    _ = v133
                                                }
                                                fmt.Println(v111, v93, v91)
                                            }
                                            v91 -= 397
                                            switch (^(fi113)) {
                                            case -1:
                                                if len(v88) >= 2 {
                                                    v88 = v88[1:len(v88)]
                                                }
                                            case -370:
                                                v92 += (v92 * (v92 * 9223372036854775807))
                                            default:
                                                v136 := GenS12{F0: 255, F1: 979, F2: fi114}
                                                _ = v136
                                            }
                                        }
                                        v2 -= (v3 / (((v3 ^ 18446744073709551615)) | 1))
                                        var v137 int16 = int16(v2)
                                        _ = v119
                                        _ = v120
                                        _ = v137
                                    }
                                case "\n\t":
                                    if false {
                                        switch fi113 {
                                        case -9223372036854775808:
                                            v93 = 222
                                        }
                                        if len(v88) > 0 {
                                            v88[0] = (fi114 ^ fi114)
                                        }
                                    }
                                }
                            case 128:
                                v2 += v2
                            }
                            var v138 float64 = (1e10 * 0.1)
                            v139 := GenS4{F0: 924, F1: v2, F2: (v93 * v93)}
                            v138 = (1e10 / v138)
                            _ = v116
                            _ = v138
                            _ = v139
                        }
                        if false {
                            fmt.Println(v93, v3)
                            var v140 uint8 = 127
                            var v141 string = ""
                            _ = v140
                            _ = v141
                        }
                        _ = fi114
                    }
                    switch (("a" + "foo bar") + "foo bar") {
                    case "a":
                        for fi142 := 0; fi142 < 5; fi142++ {
                            var v143 int8 = v111
                            v144 := GenS13{F0: -396}
                            _ = fi142
                            _ = v143
                            _ = v144
                        }
                    case "\n\t":
                        var v145 uint64 = v3
                        _ = v145
                    case "foo bar":
                        var v146 uint64 = v2
                        _ = v146
                    default:
                        if (!true) {
                            if true {
                                if false {
                                    v112 -= (v112 ^ 758)
                                }
                                switch "\n\t" {
                                case "\n\t":
                                    var v147 bool = (2 == (v111 % ((v111) | 1)))
                                    _ = v147
                                default:
                                    v148 := GenS14{F0: 2147483648, F1: (v1 ^ v1), F2: v1}
                                    _ = v148
                                }
                                v149 := GenS10{F0: (^(v92)), F1: 42, F2: v111}
                                _ = v149
                            }
                            fmt.Println(v1, v3, v93)
                        } else {
                            v111 -= v111
                            switch (fi113 - (fi113 / ((fi113) | 1))) {
                            case -1:
                                if ((v112 ^ (v2 << ((v91) & 63))) < (v3 + 1)) {
                                    switch "\n\t" {
                                    case "":
                                        var v150 int8 = (v111 >> ((v91) & 63))
                                        _ = v150
                                    case "\n\t":
                                        var v151 uint = uint(v2)
                                        _ = v151
                                    case "foo bar":
                                        v3 -= 1
                                    }
                                    v2 = 0
                                    var v152 float32 = ((0.1 * 3.14159) - (-2.25 * (float32(1e10) - -2.25)))
                                    _ = v152
                                } else {
                                    var v153 int32 = int32(v111)
                                    _ = v153
                                }
                            case 1:
                                v91 += (v91 % (((v91 ^ (v91 << ((uint(v93)) & 63)))) | 1))
                            default:
                                var v154 uint32 = 165
                                _ = v154
                            }
                        }
                    }
                    _ = fi113
                }
                _ = v111
                _ = v112
            }
            v155 := GenS4{F0: -570, F1: (v2 / ((v2) | 1)), F2: v93}
            _ = v155
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v88:", len(v88), v88)
    fmt.Println("v91:", v91)
    fmt.Println("v92:", v92)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v2
    _ = v3
    _ = v88
    _ = v89
    _ = v90
    _ = v91
    _ = v92
    _ = v93
    _ = v94
}
