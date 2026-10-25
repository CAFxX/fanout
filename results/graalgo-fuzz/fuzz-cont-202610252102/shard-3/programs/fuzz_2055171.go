package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 int16 }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 uint64; F1 int; F2 int32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int64; F1 uint8 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int; F1 int64; F2 uint }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint8; F1 bool }
type GenS10 struct { F0 uint64; F1 uint64; F2 int }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 int; F1 int; F2 uint }
type GenS13 struct { F0 uint32; F1 string; F2 uint32 }
func genFn0() int8 {
    for fi6 := 0; fi6 < 2; fi6++ {
        v7 := []int{fi6, (fi6 * fi6), (fi6 % ((127) | 1)), (fi6 << ((65535) & 63))}
        for fi8 := 0; fi8 < 10; fi8++ {
            v9 := GenS1{F0: "", F1: -9223372036854775808, F2: 42}
            _ = fi8
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    if false {
        for fi10 := 0; fi10 < 17; fi10++ {
            var v11 int = fi10
            v12 := []int{(v11 & fi10)}
            _ = fi10
            _ = v11
            _ = v12
        }
        for fi13 := 0; fi13 < 9; fi13++ {
            if (0.1 < 0.1) {
                var v14 int64 = 8
                for fi15 := 0; fi15 < 19; fi15++ {
                    v14 = v14
                    switch -267 {
                    case 463:
                        if (!(((0.0 * (0.1 / -2.25)) <= 1e10) || false)) {
                            if (false && false) {
                                switch "\n\t" {
                                case "\n\t":
                                    fmt.Println(fi13)
                                }
                            } else {
                                var v16 string = "\n\t"
                                v16 = ((((("foo bar" + v16) + "") + v16) + ("a" + "a")) + v16)
                                v17 := []int{(fi15 % (((fi15 % ((-1) | 1))) | 1)), fi15}
                                _ = v16
                                _ = v17
                            }
                            fmt.Println(fi13, v14, v14)
                            var v18 bool = (!true)
                            _ = v18
                        } else {
                            v14 -= (v14 ^ v14)
                        }
                    case -104:
                        v14 += (v14 - (v14 & (v14 & 1)))
                        v14 = (v14 - 643)
                        var v19 uint32 = 0
                        var v20 int8 = 0
                        _ = v19
                        _ = v20
                    }
                    _ = fi15
                }
                _ = v14
            } else {
                var v21 int8 = -82
                v21 = 1
                var v22 string = "a"
                _ = v21
                _ = v22
            }
            switch (fi13 * (fi13 << ((94) & 63))) {
            case -2147483648:
                var v23 float32 = 0.0
                _ = v23
            default:
                switch "foo bar" {
                case "foo bar":
                    var v24 uint64 = 142
                    _ = v24
                case "":
                    v25 := []int{(fi13 << ((18446744073709551615) & 63))}
                    _ = v25
                default:
                    v26 := []int{fi13, -758}
                    _ = v26
                }
            }
            var v27 int = (fi13 ^ (fi13 >> ((42) & 63)))
            _ = fi13
            _ = v27
        }
        v28 := []int{9223372036854775807, 42, 429}
        _ = v28
    }
    v29 := GenS2{F0: -49, F1: 733}
    switch -2147483648 {
    case 2147483647:
        v30 := GenS2{F0: -128, F1: 736}
        _ = v30
    case 372:
        for fi31 := 0; fi31 < 12; fi31++ {
            if true {
                v32 := GenS1{F0: "hello", F1: 9223372036854775807, F2: 0}
                switch "" {
                case "hello":
                    v33 := GenS2{F0: 100, F1: 65535}
                    _ = v33
                default:
                    switch "\n\t" {
                    case "hello":
                        for fi34 := 0; fi34 < 9; fi34++ {
                            var v35 int32 = 0
                            v35 -= (-(v35))
                            v36 := GenS2{F0: 1, F1: 255}
                            _ = fi34
                            _ = v35
                            _ = v36
                        }
                    case "\n\t":
                        var v37 int32 = 768
                        _ = v37
                    case "":
                        var v38 float64 = -2.25
                        _ = v38
                    default:
                        var v39 uint8 = 1
                        _ = v39
                    }
                }
                _ = v32
            } else {
                var v40 int8 = 55
                _ = v40
            }
            for fi41 := 0; fi41 < 1; fi41++ {
                fmt.Println(fi31, fi41)
                var v42 int16 = 717
                var v43 bool = false
                v42 = -130
                _ = fi41
                _ = v42
                _ = v43
            }
            _ = fi31
        }
    }
    _ = v29
    return 1
}
func genFn1() bool {
    v44 := GenS1{F0: ("foo bar" + ("a" + "")), F1: 0, F2: 0}
    if (!((0 >= -32768) && (((3.14159 - (-2.25 - (1e10 * -2.25))) >= 1.5) || (0.0 > 1e10)))) {
        v45 := []int{9223372036854775807, 127, 32768}
        var v46 int = 32767
        _ = v45
        _ = v46
    } else {
        for fi47 := 0; fi47 < 13; fi47++ {
            var v48 int32 = 1
            v48 += (v48 | -2147483648)
            if false {
                var v49 int64 = -192
                _ = v49
            } else {
                for fi50 := 0; fi50 < 12; fi50++ {
                    v51 := []int{fi47, -9223372036854775808, fi50}
                    var v52 float32 = 0.1
                    v48 += (v48 * (v48 * (v48 >> ((884) & 63))))
                    _ = fi50
                    _ = v51
                    _ = v52
                }
                fmt.Println(fi47, v48)
                v48 = -991
            }
            _ = fi47
            _ = v48
        }
        for fi53 := 0; fi53 < 20; fi53++ {
            if true {
                switch "" {
                case "hello":
                    if ("\n\t" == (("\n\t" + "a") + "hello")) {
                        v54 := []int{}
                        var v55 string = "foo bar"
                        for fi56 := 0; fi56 < 1; fi56++ {
                            v57 := GenS3{F0: 2, F1: -55, F2: 1}
                            _ = fi56
                            _ = v57
                        }
                        _ = v54
                        _ = v55
                    } else {
                        switch (fi53 | (fi53 % ((32768) | 1))) {
                        case 2147483647:
                            for fi58 := 0; fi58 < 4; fi58++ {
                                var v59 int32 = 42
                                _ = fi58
                                _ = v59
                            }
                        }
                        v60 := []int{(fi53 >> ((255) & 63)), fi53, (fi53 - 914)}
                        var v61 uint = 0
                        _ = v60
                        _ = v61
                    }
                }
                var v62 int = (-(fi53))
                for fi63 := 0; fi63 < 4; fi63++ {
                    v62 -= fi53
                    var v64 uint32 = 0
                    v64 = v64
                    _ = fi63
                    _ = v64
                }
                _ = v62
            } else {
                var v65 float64 = -2.25
                for fi66 := 0; fi66 < 12; fi66++ {
                    if (621 <= -9223372036854775808) {
                        v67 := []int{42, (fi53 ^ fi66)}
                        v65 = float64(0.0)
                        var v68 int8 = -34
                        _ = v67
                        _ = v68
                    } else {
                        switch ("hello" + "hello") {
                        case "a":
                            var v69 float64 = ((v65 * ((v65 * v65) - v65)) / v65)
                            _ = v69
                        case "foo bar":
                            v65 = v65
                        default:
                            fmt.Println(fi53)
                        }
                        var v70 int16 = 1000
                        switch (fi53 * 42) {
                        case 9223372036854775807:
                            v71 := GenS2{F0: 86, F1: 65535}
                            _ = v71
                        }
                        if false {
                            var v72 int32 = -1
                            switch (fi66 | (^(fi53))) {
                            case 32768:
                                if (true && (!false)) {
                                    var v73 float32 = 1e10
                                    _ = v73
                                }
                            case 563:
                                v74 := GenS3{F0: 4294967295, F1: (fi66 ^ -138), F2: 1}
                                _ = v74
                            default:
                                fmt.Println(v65, v72)
                            }
                            var v75 float64 = (1.5 / (v65 * v65))
                            _ = v72
                            _ = v75
                        }
                        v65 -= (v65 - v65)
                        _ = v70
                    }
                    if true {
                        if true {
                            v76 := []int{fi53, fi53, fi66}
                            var v77 int64 = -9223372036854775808
                            var v78 float64 = v65
                            var v79 uint32 = 560
                            v80 := GenS4{F0: 42}
                            _ = v76
                            _ = v77
                            _ = v78
                            _ = v79
                            _ = v80
                        } else {
                            if false {
                                var v81 uint16 = 254
                                _ = v81
                            } else {
                                if (-66 > -1) {
                                    var v82 uint8 = 183
                                    var v83 string = ""
                                    v83 = v83
                                    _ = v82
                                    _ = v83
                                }
                                var v84 int16 = 32767
                                _ = v84
                            }
                        }
                        v85 := []int{(fi53 / (((fi66 | (fi53 % (((fi53 >> ((18446744073709551615) & 63))) | 1)))) | 1)), fi66}
                        v65 = 1e10
                        _ = v85
                    } else {
                        v86 := []int{fi53}
                        _ = v86
                    }
                    _ = fi66
                }
                var v87 int32 = 0
                _ = v65
                _ = v87
            }
            _ = fi53
        }
        var v88 int64 = 0
        _ = v88
    }
    var v89 int = 42
    _ = v44
    _ = v89
    return false
}

func main() {
    var v1 int64 = -1
    v2 := []int{1, 9223372036854775807}
    var v3 float64 = 0.0
    v4 := make(map[string]uint32)
    var v5 uint16 = 42
    switch "a" {
    case "":
        if ("a" < (("foo bar" + "a") + "\n\t")) {
            switch -940 {
            case -9223372036854775808:
                for fi90 := 0; fi90 < 7; fi90++ {
                    v3 += 3.14159
                    v91 := GenS5{F0: (v1 * (v1 ^ 42)), F1: 1}
                    _ = fi90
                    _ = v91
                }
            case 1:
                var v92 uint = 478
                _ = v92
            default:
                v93 := GenS1{F0: "\n\t", F1: v1, F2: -30}
                _ = v93
            }
            _ = len(v4)
            var v94 uint64 = 236
            _ = v94
        } else {
            if len(v2) > 0 {
                v2[0] = -747
            }
            v5 = v5
        }
    case "\n\t":
        fmt.Println(v3, v3, v5)
    default:
        var v95 int = 127
        _ = v95
    }
    if true {
        v96 := v4["hello"]
        _ = v96
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        fmt.Println(v3)
        var v97 uint16 = v5
        _ = v97
    }
    if ((!(true && true)) || true) {
        if (127 >= 128) {
            var v98 uint64 = 497
            _ = v98
        } else {
            if (!false) {
                for fi99 := 0; fi99 < 1; fi99++ {
                    if true {
                        v5 -= v5
                    }
                    _ = fi99
                }
                switch "hello" {
                case "":
                    var v100 string = "hello"
                    _ = v100
                case "\n\t":
                    v2 = append(v2, 0)
                }
            } else {
                var v101 float64 = (1e10 / (v3 * v3))
                v102 := v4[("a" + "a")]
                _ = v101
                _ = v102
            }
            for fi103 := 0; fi103 < 4; fi103++ {
                var v104 uint16 = v5
                var v105 uint32 = 42
                _ = fi103
                _ = v104
                _ = v105
            }
            switch 127 {
            case 842:
                _ = len(v4)
            case 32767:
                if false {
                    for fi106 := 0; fi106 < 2; fi106++ {
                        switch 4294967295 {
                        case 32768:
                            switch fi106 {
                            case -622:
                                switch ((((("hello" + "a") + "hello") + "a") + "") + "hello") {
                                case "hello":
                                    if (false && true) {
                                        var v107 int = 42
                                        _ = v107
                                    }
                                case "a":
                                    fmt.Println(v5, fi106)
                                case "\n\t":
                                    var v108 float32 = -2.25
                                    _ = v108
                                default:
                                    for fi109 := 0; fi109 < 17; fi109++ {
                                        v110 := GenS2{F0: 42, F1: (v5 - (^(v5)))}
                                        var v111 uint64 = 4294967295
                                        if (!(((42 >= 42) && true) && true)) {
                                            v5 += 431
                                        } else {
                                            fmt.Println(v1)
                                        }
                                        _ = fi109
                                        _ = v110
                                        _ = v111
                                    }
                                }
                            default:
                                v112 := GenS6{F0: ""}
                                _ = v112
                            }
                        case -9223372036854775808:
                            for fi113 := 0; fi113 < 15; fi113++ {
                                if ((false && true) || (42 <= 897)) {
                                    var v114 uint = 65535
                                    var v115 float64 = (((float64(3.14159) - 1.5) - (v3 / v3)) * (1e10 - 3.14159))
                                    v5 += v5
                                    _ = v114
                                    _ = v115
                                }
                                switch ("a" + ("hello" + ("a" + "a"))) {
                                case "\n\t":
                                    v116 := GenS5{F0: (-(v1)), F1: 102}
                                    _ = v116
                                case "a":
                                    var v117 int8 = 1
                                    _ = v117
                                case "hello":
                                    if len(v2) > 0 {
                                        v2[0] = (fi106 & int(v1))
                                    }
                                }
                                v118 := GenS2{F0: -57, F1: 42}
                                var v119 float32 = (3.14159 - -2.25)
                                var v120 uint8 = 1
                                _ = fi113
                                _ = v118
                                _ = v119
                                _ = v120
                            }
                        case 32767:
                            var v121 bool = false
                            _ = v121
                        }
                        v3 = float64(float32(-2.25))
                        _ = fi106
                    }
                } else {
                    fmt.Println(v5, v1)
                    v122 := GenS4{F0: 32767}
                    _ = v122
                }
            }
        }
    } else {
        for fi123 := 0; fi123 < 6; fi123++ {
            for fi124 := 0; fi124 < 4; fi124++ {
                switch fi124 {
                case 4294967295:
                    v1 += -9223372036854775808
                case 32767:
                    v3 = (v3 * (1.5 * (float64((1.5 / (((0.1 - 1e10) - 1.5) * 3.14159))) * (((((float64((float32(v3) * (0.1 + (((3.14159 - (1e10 + ((-2.25 + 1e10) - (((0.1 / 3.14159) / 1.0) / 1.0)))) / 1.5) + float32(1e10))))) + float64(3.14159)) * 3.14159) / -2.25) * v3) * v3))))
                default:
                    fmt.Println(v5, v1, fi124)
                }
                v1 += 9223372036854775807
                _ = fi124
            }
            _ = fi123
        }
    }
    v3 = 0.0
    if false {
        v125 := GenS7{F0: 4294967295, F1: v1, F2: 42}
        switch -1 {
        case 2147483647:
            for fi126 := 0; fi126 < 15; fi126++ {
                v127 := GenS6{F0: "hello"}
                if false {
                    var v128 int8 = -128
                    _ = v128
                } else {
                    v4["hello"] = 923
                    if true {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        var v129 uint = 18446744073709551615
                        if true {
                            v4[("hello" + ("a" + "foo bar"))] = 678
                            for fi130 := 0; fi130 < 7; fi130++ {
                                v3 -= v3
                                _ = fi130
                            }
                            if (false && (false && (v129 >= (v129 - v129)))) {
                                var v131 float32 = 1.5
                                switch fi126 {
                                case 0:
                                    for fi132 := 0; fi132 < 6; fi132++ {
                                        fmt.Println(v5)
                                        if (42 < 1) {
                                            v133 := GenS5{F0: v1, F1: 0}
                                            _ = v133
                                        } else {
                                            v131 = 0.1
                                            fmt.Println(v5)
                                        }
                                        _ = fi132
                                    }
                                }
                                for fi134 := 0; fi134 < 16; fi134++ {
                                    var v135 float64 = v3
                                    var v136 int = (fi134 / ((fi126) | 1))
                                    _ = fi134
                                    _ = v135
                                    _ = v136
                                }
                                _ = v131
                            } else {
                                var v137 int16 = 1
                                switch fi126 {
                                case -815:
                                    for fi138 := 0; fi138 < 5; fi138++ {
                                        v1 -= -812
                                        _ = fi138
                                    }
                                case 42:
                                    v139 := GenS8{F0: -916}
                                    _ = v139
                                case -748:
                                    v140 := GenS6{F0: "\n\t"}
                                    _ = v140
                                default:
                                    v137 = 1000
                                }
                                v5 = v5
                                _ = v137
                            }
                        } else {
                            var v141 uint = v129
                            v2 = append(v2, fi126)
                            _ = v141
                        }
                        _ = v129
                    } else {
                        var v142 int = fi126
                        _ = v142
                    }
                    switch ("foo bar" + ("a" + ("foo bar" + ""))) {
                    case "a":
                        v1 -= (v1 | v1)
                    case "hello":
                        v1 -= v1
                    case "\n\t":
                        if false {
                            for fi143 := 0; fi143 < 7; fi143++ {
                                var v144 uint8 = 255
                                _ = fi143
                                _ = v144
                            }
                            v145 := GenS9{F0: 42, F1: (((((0.1 + 1.5) / (0.0 - (float32(v3) * 0.1))) * 0.0) == -2.25) || true)}
                            switch 32767 {
                            case 128:
                                v146 := GenS10{F0: 42, F1: 28, F2: 128}
                                _ = v146
                            default:
                                v5 = 65535
                            }
                            _ = v145
                        } else {
                            for fi147 := 0; fi147 < 12; fi147++ {
                                if len(v2) > 0 {
                                    v2[0] = fi147
                                }
                                if (0 != 1) {
                                    _ = len(v4)
                                    v5 = (v5 << ((65535) & 63))
                                } else {
                                    switch fi147 {
                                    case -1:
                                        var v148 uint32 = 374
                                        _ = v148
                                    default:
                                        v149 := v4["\n\t"]
                                        _ = v149
                                    }
                                    switch int(v1) {
                                    case 4294967295:
                                        switch fi147 {
                                        case 42:
                                            v1 = (v1 + 0)
                                        case 128:
                                            var v150 int64 = v1
                                            _ = v150
                                        case -1:
                                            v5 += v5
                                        default:
                                            fmt.Println(fi147, fi147)
                                        }
                                    default:
                                        v1 += 150
                                    }
                                    var v151 uint = 352
                                    _ = v151
                                }
                                var v152 uint16 = v5
                                _ = fi147
                                _ = v152
                            }
                        }
                    default:
                        if false {
                            v3 = v3
                        }
                    }
                }
                _ = fi126
                _ = v127
            }
        default:
            switch 0 {
            case 84:
                var v153 float64 = (((float64(3.14159) / v3) + 3.14159) + ((1e10 + 0.1) + (v3 * v3)))
                _ = v153
            case 42:
                if false {
                    v3 = (1e10 / 3.14159)
                    v3 = ((float64(3.14159) + ((v3 / (float64(0.0) / v3)) / v3)) / 0.1)
                    var v154 int8 = genFn0()
                    _ = v154
                }
            default:
                var v155 uint64 = 4294967295
                _ = v155
            }
        }
        _ = v125
    }
    var v156 float64 = v3
    v157 := GenS6{F0: "\n\t"}
    v1 = v1
    for fi158 := 0; fi158 < 19; fi158++ {
        _ = len(v2)
        if false {
            var v159 int16 = -384
            if (1 >= 0) {
                v160 := GenS11{F0: 193}
                var v161 string = ""
                v162 := GenS4{F0: v159}
                _ = v160
                _ = v161
                _ = v162
            }
            var v163 int = 32767
            _ = v159
            _ = v163
        } else {
            var v164 int8 = 0
            v165 := v4[(("hello" + ("a" + "")) + "foo bar")]
            var v166 int8 = v164
            _ = v164
            _ = v165
            _ = v166
        }
        _ = fi158
    }
    if true {
        if len(v2) > 0 {
            v2[0] = -1
        }
        v167 := GenS4{F0: 0}
        v1 = (^(v1))
        _ = v167
    } else {
        v156 = ((float64(((((float32(((1e10 - v3) / v3)) + (0.0 - -2.25)) / (((1.5 * 1.5) - 0.0) * (0.0 / float32((v3 * v3))))) - float32(1e10)) * (float32(v156) - 1.5))) * v156) * v156)
        var v168 string = ("a" + ("hello" + "foo bar"))
        if (!(false && true)) {
            v3 = ((float64((0.1 - 3.14159)) + ((v156 / (float64(0.0) - v3)) / v156)) * 1.5)
            if true {
                if true {
                    v5 += (v5 + v5)
                    var v169 uint64 = 2
                    if (v169 == (v169 / ((v169) | 1))) {
                        var v170 float32 = (float32(v156) * (((0.1 / (float32(v3) + (1e10 * 0.1))) / 1e10) - 1e10))
                        for fi171 := 0; fi171 < 2; fi171++ {
                            if (v169 < v169) {
                                switch "" {
                                case "":
                                    var v172 int32 = 762
                                    _ = v172
                                case "\n\t":
                                    v173 := GenS2{F0: -128, F1: v5}
                                    _ = v173
                                }
                            } else {
                                for fi174 := 0; fi174 < 17; fi174++ {
                                    if ((v168 > v168) || false) {
                                        var v175 float64 = 0.0
                                        var v176 int = (fi171 / (((fi171 % ((32767) | 1))) | 1))
                                        _ = v175
                                        _ = v176
                                    }
                                    v5 -= (v5 - 65535)
                                    v169 = (v169 >> ((23) & 63))
                                    _ = fi174
                                }
                            }
                            _ = fi171
                        }
                        if true {
                            if true {
                                v177 := GenS11{F0: (v169 >> ((0) & 63))}
                                var v178 uint16 = v5
                                _ = v177
                                _ = v178
                            }
                        }
                        _ = v170
                    } else {
                        switch 501 {
                        case 4294967295:
                            var v179 uint = 65535
                            _ = v179
                        default:
                            if (1 > 578) {
                                v1 -= -9223372036854775808
                                var v180 int = 32767
                                v3 -= (v3 / ((v3 * 0.0) * (1.5 * v3)))
                                _ = v180
                            } else {
                                if true {
                                    for fi181 := 0; fi181 < 5; fi181++ {
                                        v156 = 1.5
                                        var v182 int16 = -32768
                                        _ = fi181
                                        _ = v182
                                    }
                                    for fi183 := 0; fi183 < 3; fi183++ {
                                        var v184 uint32 = 761
                                        fmt.Println(v169, v5, v3)
                                        _ = fi183
                                        _ = v184
                                    }
                                }
                                fmt.Println(v1, v156)
                            }
                        }
                        for fi185 := 0; fi185 < 11; fi185++ {
                            for fi186 := 0; fi186 < 3; fi186++ {
                                v5 = uint16(v169)
                                v187 := GenS4{F0: -1}
                                if (65535 >= 739) {
                                    var v188 uint64 = 42
                                    _ = v188
                                } else {
                                    v169 += v169
                                    if (((fi186 / ((0) | 1)) >= (fi186 | (fi186 & (-(fi186))))) || (!true)) {
                                        v189 := v4["hello"]
                                        var v190 float32 = 0.1
                                        _ = v189
                                        _ = v190
                                    } else {
                                        fmt.Println(v156, v169, v156)
                                        v169 = (v169 ^ v169)
                                        switch "\n\t" {
                                        case "hello":
                                            var v191 uint16 = v5
                                            _ = v191
                                        case "\n\t":
                                            var v192 uint = 360
                                            fmt.Println(v1, v1)
                                            _ = v192
                                        default:
                                            var v193 uint32 = 42
                                            _ = v193
                                        }
                                    }
                                    v168 = v168
                                    v169 += (v169 >> ((459) & 63))
                                }
                                _ = len(v2)
                                if ((!false) && ((-2147483648 == 42) || (!false))) {
                                    var v194 uint = 2
                                    _ = v194
                                } else {
                                    v156 -= 3.14159
                                    v195 := GenS8{F0: -243}
                                    var v196 int32 = 0
                                    switch "" {
                                    case "foo bar":
                                        v169 -= 155
                                    case "\n\t":
                                        switch ((v168 + "a") + (v168 + (v168 + v168))) {
                                        case "":
                                            _ = len(v4)
                                        case "\n\t":
                                            v3 = v156
                                        }
                                        v4[""] = 0
                                    default:
                                        v3 += ((0.1 / 0.1) + v156)
                                    }
                                    _ = v195
                                    _ = v196
                                }
                                _ = fi186
                                _ = v187
                            }
                            _ = fi185
                        }
                        for fi197 := 0; fi197 < 17; fi197++ {
                            var v198 uint64 = v169
                            v156 -= (0.0 - 1e10)
                            _ = fi197
                            _ = v198
                        }
                    }
                    _ = v169
                } else {
                    for fi199 := 0; fi199 < 14; fi199++ {
                        v200 := GenS4{F0: 0}
                        _ = fi199
                        _ = v200
                    }
                    if (!true) {
                        var v201 uint8 = 211
                        switch (v168 + v168) {
                        case "foo bar":
                            if (!(v168 <= "")) {
                                for fi202 := 0; fi202 < 14; fi202++ {
                                    v3 = v156
                                    _ = fi202
                                }
                                var v203 uint64 = 318
                                _ = v203
                            }
                        case "hello":
                            for fi204 := 0; fi204 < 1; fi204++ {
                                v1 = (v1 ^ 1)
                                _ = fi204
                            }
                        default:
                            var v205 uint = 0
                            _ = v205
                        }
                        if false {
                            if true {
                                delete(v4, v168)
                            } else {
                                switch v168 {
                                case "":
                                    if false {
                                        for fi206 := 0; fi206 < 20; fi206++ {
                                            if true {
                                                if (true && (2 <= 779)) {
                                                    v3 = (0.1 + (float64((1e10 - -2.25)) + 0.1))
                                                } else {
                                                    fmt.Println(v5)
                                                }
                                                v5 = (v5 << ((65535) & 63))
                                            }
                                            _ = fi206
                                        }
                                        for fi207 := 0; fi207 < 1; fi207++ {
                                            v156 += ((float64(-2.25) - 3.14159) + v156)
                                            v201 = (v201 | v201)
                                            fmt.Println(v168)
                                            _ = fi207
                                        }
                                        var v208 int8 = 0
                                        _ = v208
                                    }
                                default:
                                    var v209 int16 = -1
                                    _ = v209
                                }
                                if false {
                                    if len(v2) > 0 {
                                        v2[0] = 0
                                    }
                                    var v210 float64 = float64(1.5)
                                    if false {
                                        switch -232 {
                                        case -783:
                                            if len(v2) > 0 {
                                                v2[0] = 105
                                            }
                                        case 32767:
                                            _ = len(v4)
                                        }
                                        if true {
                                            var v211 bool = ((false && (127 == -1)) && false)
                                            v212 := GenS4{F0: 0}
                                            var v213 int = 32767
                                            _ = v211
                                            _ = v212
                                            _ = v213
                                        }
                                        if (true || (1 == 828)) {
                                            v214 := GenS8{F0: 1000}
                                            _ = v214
                                        }
                                    }
                                    _ = v210
                                } else {
                                    for fi215 := 0; fi215 < 18; fi215++ {
                                        v1 = v1
                                        var v216 int8 = -128
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        _ = fi215
                                        _ = v216
                                    }
                                    if true {
                                        v168 = ""
                                    }
                                }
                                if false {
                                    if (4294967295 != 2) {
                                        v217 := v4[((v168 + ("a" + "hello")) + v168)]
                                        if false {
                                            v201 -= (v201 / (((v201 + 127)) | 1))
                                        }
                                        v201 += (v201 - (v201 | v201))
                                        _ = v217
                                    }
                                    fmt.Println(v1, v168)
                                    var v218 uint8 = 127
                                    _ = v218
                                } else {
                                    v5 += uint16(v201)
                                    fmt.Println(v168, v5)
                                    v201 += v201
                                }
                            }
                            for fi219 := 0; fi219 < 5; fi219++ {
                                var v220 int64 = 42
                                _ = fi219
                                _ = v220
                            }
                            var v221 uint32 = 42
                            _ = v221
                        }
                        _ = v201
                    } else {
                        v222 := GenS3{F0: 4294967295, F1: 127, F2: 759}
                        v223 := GenS12{F0: 42, F1: 32768, F2: 4294967295}
                        _ = v222
                        _ = v223
                    }
                    v224 := GenS13{F0: 0, F1: v168, F2: 938}
                    _ = v224
                }
                if true {
                    if (0 <= 2147483648) {
                        v156 += 0.0
                    }
                } else {
                    if false {
                        var v225 uint = 1
                        switch "" {
                        case "\n\t":
                            fmt.Println(v5)
                        case "foo bar":
                            fmt.Println(v225)
                        default:
                            if (288 == (v5 | v5)) {
                                for fi226 := 0; fi226 < 5; fi226++ {
                                    var v227 uint64 = 0
                                    for fi228 := 0; fi228 < 14; fi228++ {
                                        var v229 uint32 = 760
                                        v156 = v156
                                        v229 -= v229
                                        _ = fi228
                                        _ = v229
                                    }
                                    v3 -= (v156 - v156)
                                    _ = fi226
                                    _ = v227
                                }
                                if (true && true) {
                                    switch "a" {
                                    case "foo bar":
                                        fmt.Println(v1, v5, v1)
                                    case "":
                                        var v230 uint64 = 1
                                        _ = v230
                                    case "\n\t":
                                        var v231 int = 0
                                        _ = v231
                                    default:
                                        v156 = ((-2.25 * (v3 + v3)) * ((-2.25 - (1e10 * (v156 + 0.0))) / (v3 / 1.0)))
                                    }
                                }
                                v232 := GenS9{F0: 1, F1: false}
                                _ = v232
                            }
                        }
                        _ = v225
                    } else {
                        if true {
                            v156 = float64(float32(3.14159))
                        }
                        for fi233 := 0; fi233 < 20; fi233++ {
                            var v234 int32 = 776
                            v4[v168] = 1
                            _ = fi233
                            _ = v234
                        }
                        v2 = append(v2, 9223372036854775807)
                    }
                    var v235 int16 = -11
                    _ = v235
                }
            } else {
                var v236 int8 = 1
                delete(v4, "\n\t")
                v236 = v236
                _ = v236
            }
        }
        _ = v168
    }
    var v237 float32 = (1e10 - 0.1)
    v156 -= (((float64(float32((float64((float32(1.5) + (3.14159 * v237))) - v156))) * float64(((((v237 + v237) * 0.1) / (v237 * v237)) + v237))) * ((float64((3.14159 / 0.1)) / v3) * (1.5 + float64((float32(v156) / float32(float64(((v237 - 0.0) - 0.1)))))))) - v156)
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v156:", v156)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v237:", v237)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v156
    _ = v157
    _ = v2
    _ = v237
    _ = v3
    _ = v4
    _ = v5
}
