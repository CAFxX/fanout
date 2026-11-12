package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 int8; F1 uint32 }
type GenS4 struct { F0 bool; F1 uint8 }
type GenS5 struct { F0 int64; F1 uint8 }
type GenS6 struct { F0 uint16; F1 uint }
type GenS7 struct { F0 uint16; F1 bool; F2 uint64 }
type GenS8 struct { F0 uint; F1 bool; F2 uint64 }
type GenS9 struct { F0 int16; F1 uint32 }
type GenS10 struct { F0 uint8; F1 int32; F2 int }
type GenS11 struct { F0 uint8; F1 int64; F2 uint8 }
func genFn0(p0 string, p1 int, p2 int8) uint32 {
    var v4 int16 = -678
    if true {
        var v5 float32 = 0.1
        for fi6 := 0; fi6 < 18; fi6++ {
            p1 -= fi6
            for fi7 := 0; fi7 < 7; fi7++ {
                var v8 uint8 = 195
                var v9 int32 = 1
                v10 := GenS1{F0: 2, F1: -822}
                v11 := GenS2{F0: (v9 % ((-1) | 1)), F1: 42}
                var v12 float64 = 0.1
                switch "\n\t" {
                case "foo bar":
                    v13 := GenS3{F0: p2, F1: 2}
                    _ = v13
                default:
                    var v14 uint16 = 0
                    var v15 string = p0
                    _ = v14
                    _ = v15
                }
                _ = fi7
                _ = v10
                _ = v11
                _ = v12
                _ = v8
                _ = v9
            }
            var v16 float32 = v5
            _ = fi6
            _ = v16
        }
        _ = v5
    }
    for fi17 := 0; fi17 < 2; fi17++ {
        if true {
            var v18 int = int(v4)
            v19 := []int{p1, (v18 & (v18 + (^(p1)))), fi17, (v18 >> ((18446744073709551615) & 63))}
            for fi20 := 0; fi20 < 8; fi20++ {
                if len(v19) >= 2 {
                    v19 = v19[1:len(v19)]
                }
                if (264 == 67) {
                    v19 = append(v19, v18)
                    v4 = v4
                    p1 += 0
                } else {
                    if (false || ((float64((-2.25 / ((3.14159 - (0.0 - (1.5 / 1.5))) + (((0.1 - (0.0 * (float32(0.0) / (0.0 - 0.1)))) - 0.1) * ((0.1 / float32(1.5)) - (float32(-2.25) * 3.14159)))))) * 0.0) != 1.5)) {
                        if true {
                            p2 = -1
                        }
                        var v21 int16 = v4
                        if ((v18 | -1) == (fi20 >> ((314) & 63))) {
                            if len(v19) >= 2 {
                                v19 = v19[1:len(v19)]
                            }
                        }
                        if (1.5 >= ((-2.25 * 3.14159) / 0.1)) {
                            v22 := GenS2{F0: 0, F1: 42}
                            _ = v22
                        }
                        _ = v21
                    } else {
                        p0 = p0
                        p1 = (v18 >> ((1) & 63))
                        var v23 string = p0
                        _ = v23
                    }
                }
                var v24 uint32 = 1
                _ = fi20
                _ = v24
            }
            _ = v18
            _ = v19
        } else {
            var v25 bool = true
            v25 = ((!v25) || false)
            _ = v25
        }
        _ = fi17
    }
    if ((65535 == 255) && (!false)) {
        for fi26 := 0; fi26 < 19; fi26++ {
            v27 := GenS2{F0: 42, F1: 4294967295}
            var v28 int32 = 1
            var v29 int64 = -1
            _ = fi26
            _ = v27
            _ = v28
            _ = v29
        }
    }
    if false {
        fmt.Println(p2, p2)
        for fi30 := 0; fi30 < 9; fi30++ {
            switch p0 {
            case "hello":
                var v31 uint64 = 42
                _ = v31
            case "a":
                p1 = (fi30 + 9223372036854775807)
            case "":
                var v32 float64 = 0.1
                _ = v32
            default:
                v4 += (v4 * v4)
            }
            _ = fi30
        }
        for fi33 := 0; fi33 < 13; fi33++ {
            var v34 uint8 = 1
            _ = fi33
            _ = v34
        }
    } else {
        switch 32768 {
        case -188:
            for fi35 := 0; fi35 < 17; fi35++ {
                for fi36 := 0; fi36 < 1; fi36++ {
                    var v37 uint32 = 0
                    if (2 != 138) {
                        fmt.Println(fi36, fi35, v37)
                        p1 += fi36
                        p2 = 100
                        v38 := []int{p1, 0, 128, fi36}
                        v37 = v37
                        var v39 int = (-(fi36))
                        _ = v38
                        _ = v39
                    } else {
                        if (true || (-9223372036854775808 <= 42)) {
                            if true {
                                switch (((p0 + p0) + (p0 + p0)) + p0) {
                                case "a":
                                    if true {
                                        fmt.Println(p1, p1, p1)
                                        v37 = (v37 + v37)
                                    } else {
                                        v40 := GenS1{F0: 279, F1: -1}
                                        _ = v40
                                    }
                                case "foo bar":
                                    p2 -= p2
                                case "":
                                    switch 32768 {
                                    case -9223372036854775808:
                                        v41 := []int{int(v37)}
                                        _ = v41
                                    }
                                default:
                                    v42 := []int{-32}
                                    _ = v42
                                }
                            }
                            v43 := GenS3{F0: (p2 & int8(p1)), F1: (v37 ^ (v37 % (((^(v37))) | 1)))}
                            v44 := []int{fi36, (-(fi36)), (fi35 ^ -2147483648), (^(p1))}
                            _ = v43
                            _ = v44
                        } else {
                            v45 := []int{910, 9223372036854775807, (p1 / ((p1) | 1))}
                            v46 := GenS4{F0: false, F1: 1}
                            _ = v45
                            _ = v46
                        }
                        var v47 uint64 = 42
                        var v48 uint16 = uint16(v47)
                        _ = v47
                        _ = v48
                    }
                    _ = fi36
                    _ = v37
                }
                switch (fi35 % ((int(v4)) | 1)) {
                case 122:
                    var v49 uint16 = 0
                    _ = v49
                case 2147483647:
                    for fi50 := 0; fi50 < 6; fi50++ {
                        if false {
                            if (-2.25 <= float32((0.0 * ((0.0 / 1.0) - (1e10 + (1.5 / 3.14159)))))) {
                                if (!(p2 != (^(p2)))) {
                                    v51 := []int{fi50, fi35}
                                    var v52 int64 = 42
                                    _ = v51
                                    _ = v52
                                }
                                fmt.Println(fi35, fi50, p0)
                            }
                            if (false || false) {
                                switch int(p2) {
                                case -907:
                                    p2 += p2
                                case 127:
                                    if true {
                                        var v53 uint8 = 255
                                        switch 127 {
                                        case -203:
                                            switch p1 {
                                            case 122:
                                                p1 = fi35
                                            case -832:
                                                fmt.Println(fi35)
                                            case 42:
                                                var v54 int32 = -1
                                                _ = v54
                                            }
                                        case 42:
                                            v55 := []int{985, fi50, fi35}
                                            _ = v55
                                        case -329:
                                            var v56 bool = true
                                            _ = v56
                                        }
                                        if (!true) {
                                            v4 -= -472
                                        }
                                        _ = v53
                                    } else {
                                        var v57 int = (fi50 >> ((665) & 63))
                                        p0 = "foo bar"
                                        var v58 uint32 = 2147483647
                                        _ = v57
                                        _ = v58
                                    }
                                case 1:
                                    p2 = (p2 & p2)
                                default:
                                    var v59 string = p0
                                    _ = v59
                                }
                            } else {
                                var v60 int32 = 42
                                var v61 uint32 = 2
                                var v62 int16 = 1
                                switch "" {
                                case "\n\t":
                                    switch 4294967295 {
                                    case 42:
                                        v4 -= 42
                                    case -905:
                                        v63 := []int{}
                                        _ = v63
                                    default:
                                        if (p0 > p0) {
                                            v64 := GenS4{F0: (!true), F1: 255}
                                            v65 := []int{fi50, p1, fi35}
                                            v65 = append(v65, fi50)
                                            var v66 int32 = v60
                                            if (true && false) {
                                                if true {
                                                    var v67 int8 = p2
                                                    fmt.Println(v66, v61)
                                                    switch p0 {
                                                    case "":
                                                        var v68 uint = 4294967295
                                                        _ = v68
                                                    case "foo bar":
                                                        v69 := GenS4{F0: false, F1: 135}
                                                        _ = v69
                                                    default:
                                                        var v70 uint = 255
                                                        _ = v70
                                                    }
                                                    switch fi50 {
                                                    case -774:
                                                        _ = len(v65)
                                                    case 42:
                                                        fmt.Println(p0, p2, v62)
                                                    case -429:
                                                        var v71 int = (p1 >> ((255) & 63))
                                                        _ = v71
                                                    default:
                                                        _ = len(v65)
                                                    }
                                                    if false {
                                                        var v72 float32 = ((float32((0.1 * 0.1)) - float32((1.5 * ((0.0 * 0.1) + (1.5 - 3.14159))))) * (0.0 * (1e10 - float32(3.14159))))
                                                        _ = v72
                                                    } else {
                                                        switch p0 {
                                                        case "\n\t":
                                                            v60 = (-(v66))
                                                        case "":
                                                            var v73 uint16 = 255
                                                            _ = v73
                                                        case "foo bar":
                                                            _ = len(v65)
                                                        }
                                                        if len(v65) >= 2 {
                                                            v65 = v65[1:len(v65)]
                                                        }
                                                    }
                                                    var v74 uint16 = 255
                                                    var v75 string = "a"
                                                    var v76 uint16 = (v74 + (v74 & (v74 & v74)))
                                                    _ = v67
                                                    _ = v74
                                                    _ = v75
                                                    _ = v76
                                                } else {
                                                    var v77 uint32 = (v61 & 65535)
                                                    fmt.Println(v66, p2, fi35)
                                                    if ("" != p0) {
                                                        if false {
                                                            v61 += v77
                                                        }
                                                        p2 = 127
                                                    } else {
                                                        fmt.Println(fi35)
                                                        v78 := GenS3{F0: (p2 & (-(p2))), F1: v77}
                                                        _ = v78
                                                    }
                                                    _ = v77
                                                }
                                                if len(v65) > 0 {
                                                    v65[0] = (p1 / ((p1) | 1))
                                                }
                                                v79 := GenS5{F0: -790, F1: 1}
                                                _ = v79
                                            } else {
                                                v80 := GenS1{F0: 18446744073709551615, F1: 90}
                                                v66 = -2147483648
                                                v66 = v60
                                                fmt.Println(fi50)
                                                _ = v80
                                            }
                                            _ = v64
                                            _ = v65
                                            _ = v66
                                        }
                                    }
                                case "a":
                                    v81 := GenS6{F0: 2, F1: 65535}
                                    _ = v81
                                case "":
                                    v60 = v60
                                    if true {
                                        p0 = "hello"
                                        var v82 uint16 = 255
                                        var v83 int64 = 1
                                        v82 = 255
                                        v84 := []int{int(v60)}
                                        _ = v82
                                        _ = v83
                                        _ = v84
                                    } else {
                                        switch (p0 + p0) {
                                        case "":
                                            fmt.Println(v60, p1, p0)
                                        }
                                        fmt.Println(v4, v62)
                                        switch "hello" {
                                        case "a":
                                            v85 := []int{fi35, (p1 - (fi50 | fi35))}
                                            _ = v85
                                        default:
                                            var v86 int8 = p2
                                            _ = v86
                                        }
                                    }
                                    switch 128 {
                                    case 32768:
                                        var v87 string = "a"
                                        _ = v87
                                    default:
                                        switch (fi35 << ((206) & 63)) {
                                        case 2147483647:
                                            p2 = p2
                                        case -18:
                                            var v88 int32 = v60
                                            _ = v88
                                        case 4294967295:
                                            var v89 uint64 = 18446744073709551615
                                            _ = v89
                                        }
                                    }
                                }
                                _ = v60
                                _ = v61
                                _ = v62
                            }
                            v90 := GenS6{F0: 65535, F1: 255}
                            v91 := []int{-51, 32767, fi35, 169}
                            _ = v90
                            _ = v91
                        }
                        v92 := []int{(fi50 - (fi35 - 4294967295)), p1, (fi50 | int(p2)), (fi50 >> ((0) & 63))}
                        _ = fi50
                        _ = v92
                    }
                default:
                    var v93 uint = 18446744073709551615
                    _ = v93
                }
                _ = fi35
            }
        case -11:
            if true {
                var v94 string = ("" + "a")
                _ = v94
            } else {
                switch p0 {
                case "hello":
                    if false {
                        fmt.Println(p1)
                    } else {
                        var v95 int64 = -953
                        var v96 uint16 = 42
                        _ = v95
                        _ = v96
                    }
                case "foo bar":
                    for fi97 := 0; fi97 < 2; fi97++ {
                        if false {
                            fmt.Println(p1)
                            for fi98 := 0; fi98 < 10; fi98++ {
                                var v99 bool = (false || false)
                                var v100 float64 = (-2.25 / 3.14159)
                                _ = fi98
                                _ = v100
                                _ = v99
                            }
                            var v101 uint = 0
                            _ = v101
                        } else {
                            switch (p0 + p0) {
                            case "a":
                                if (false && true) {
                                    p0 = p0
                                    var v102 float32 = 0.0
                                    _ = v102
                                }
                            case "hello":
                                v103 := []int{int(v4), (fi97 & (-(p1))), p1, -1}
                                _ = v103
                            }
                        }
                        for fi104 := 0; fi104 < 10; fi104++ {
                            var v105 bool = false
                            _ = fi104
                            _ = v105
                        }
                        _ = fi97
                    }
                case "":
                    v106 := []int{(p1 >> ((601) & 63))}
                    _ = v106
                default:
                    p2 -= p2
                }
                if true {
                    var v107 uint = 1
                    v4 += (v4 % (((v4 % (((v4 % ((v4) | 1))) | 1))) | 1))
                    _ = v107
                } else {
                    var v108 float32 = (((1e10 - 1.5) - -2.25) + (0.1 * 0.1))
                    if (1 >= 2) {
                        v109 := GenS4{F0: true, F1: 1}
                        _ = v109
                    }
                    _ = v108
                }
            }
        default:
            fmt.Println(p2, p0, p0)
        }
        var v110 int64 = 1
        v111 := GenS2{F0: 2147483647, F1: 2}
        _ = v110
        _ = v111
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 4294967295
}

func main() {
    v1 := make(map[string]bool)
    var v2 int16 = -1
    var v3 int16 = v2
    var v112 int64 = 105
    for fi113 := 0; fi113 < 7; fi113++ {
        var v114 float64 = 3.14159
        var v115 bool = (2147483648 == 2)
        _ = fi113
        _ = v114
        _ = v115
    }
    v1[""] = (0.1 <= float64(float32((-2.25 - (1e10 * 1.5)))))
    var v116 int8 = 127
    var v117 uint = 65535
    var v118 uint32 = 4294967295
    v119 := GenS7{F0: 32767, F1: (v116 >= v116), F2: 4294967295}
    var v120 float32 = (1e10 + 1e10)
    v120 = (float32(0.1) * 0.1)
    v121 := GenS6{F0: 1, F1: v117}
    if ((v3 - -98) != (v3 | v2)) {
        if false {
            var v122 string = ""
            if false {
                v123 := GenS2{F0: -1, F1: 447}
                _ = v123
            } else {
                switch 2147483647 {
                case -103:
                    for fi124 := 0; fi124 < 18; fi124++ {
                        var v125 uint32 = v118
                        v126 := GenS6{F0: 530, F1: v117}
                        for fi127 := 0; fi127 < 9; fi127++ {
                            if (v3 >= (v2 * (v3 | v2))) {
                                v112 = (v112 - (-(v112)))
                            }
                            v128 := GenS8{F0: (v117 ^ 255), F1: true, F2: 2}
                            _ = fi127
                            _ = v128
                        }
                        _ = fi124
                        _ = v125
                        _ = v126
                    }
                case 42:
                    var v129 uint32 = (v118 | 42)
                    _ = v129
                default:
                    fmt.Println(v117, v2)
                }
                var v130 int16 = 0
                if (uint8(v117) < 255) {
                    var v131 string = v122
                    if (((!(0 != 42)) && (!(uint64(v117) != 2))) && true) {
                        var v132 float64 = 0.0
                        _ = v132
                    }
                    _ = v131
                } else {
                    for fi133 := 0; fi133 < 9; fi133++ {
                        for fi134 := 0; fi134 < 14; fi134++ {
                            var v135 float32 = 1e10
                            v136 := GenS3{F0: -1, F1: 2147483648}
                            v137 := GenS8{F0: v117, F1: false, F2: 2}
                            _ = fi134
                            _ = v135
                            _ = v136
                            _ = v137
                        }
                        var v138 float64 = 3.14159
                        switch 127 {
                        case 0:
                            fmt.Println(v130)
                        case 42:
                            switch "foo bar" {
                            case "foo bar":
                                v3 = -104
                            }
                        case 128:
                            var v139 uint = v117
                            _ = v139
                        }
                        _ = fi133
                        _ = v138
                    }
                }
                _ = v130
            }
            _ = v122
        } else {
            v140 := GenS7{F0: 0, F1: (true && true), F2: 4294967295}
            _ = v140
        }
        if false {
            v141 := GenS5{F0: (v112 << (((^(v117))) & 63)), F1: 190}
            var v142 uint8 = 2
            for fi143 := 0; fi143 < 4; fi143++ {
                for fi144 := 0; fi144 < 3; fi144++ {
                    if false {
                        var v145 int16 = (^(v2))
                        _ = v145
                    } else {
                        v118 = 2147483648
                        v117 = v117
                        if (v112 < int64(v117)) {
                            fmt.Println(v116, v116, v120)
                            if true {
                                v117 += v117
                                var v146 float64 = 1e10
                                v116 += int8(v117)
                                _ = v146
                            } else {
                                switch ("foo bar" + ("a" + "a")) {
                                case "":
                                    v3 += v3
                                case "hello":
                                    var v147 int8 = v116
                                    _ = v147
                                case "\n\t":
                                    v148 := []int{fi143, (fi143 << (((v117 | v117)) & 63))}
                                    _ = v148
                                default:
                                    v149 := GenS7{F0: 989, F1: (0 < 409), F2: 42}
                                    _ = v149
                                }
                                v150 := []int{}
                                v1[""] = (!(v2 < v2))
                                _ = v150
                            }
                        }
                        var v151 int64 = v112
                        var v152 uint8 = v142
                        _ = v151
                        _ = v152
                    }
                    var v153 int16 = (-(v2))
                    var v154 int = fi144
                    _ = fi144
                    _ = v153
                    _ = v154
                }
                v117 = 1
                _ = fi143
            }
            _ = v141
            _ = v142
        } else {
            v155 := GenS9{F0: -32768, F1: (^(v118))}
            _ = v155
        }
    } else {
        for fi156 := 0; fi156 < 10; fi156++ {
            for fi157 := 0; fi157 < 19; fi157++ {
                var v158 string = ("foo bar" + (("a" + "hello") + ("hello" + "foo bar")))
                var v159 int16 = (v3 / (((v3 | v3)) | 1))
                switch fi157 {
                case 32767:
                    v158 = v158
                }
                var v160 int16 = 1000
                _ = fi157
                _ = v158
                _ = v159
                _ = v160
            }
            var v161 uint32 = 2
            var v162 uint32 = (^(v118))
            _ = fi156
            _ = v161
            _ = v162
        }
        switch 2147483647 {
        case -388:
            v116 -= v116
        }
    }
    v117 = (v117 / ((v117) | 1))
    var v163 int16 = 42
    switch 324 {
    case -1:
        for fi164 := 0; fi164 < 19; fi164++ {
            var v165 int64 = 1
            var v166 string = "foo bar"
            var v167 int64 = v165
            _ = fi164
            _ = v165
            _ = v166
            _ = v167
        }
    case -236:
        for fi168 := 0; fi168 < 20; fi168++ {
            v117 = uint(v118)
            switch "foo bar" {
            case "hello":
                var v169 int64 = -79
                _ = v169
            case "a":
                var v170 int16 = 1
                _ = v170
            case "foo bar":
                v112 = (v112 | v112)
            default:
                if ((v117 & v117) == 953) {
                    v171 := GenS10{F0: uint8(v117), F1: 318, F2: (-(fi168))}
                    var v172 int16 = v3
                    v173 := GenS4{F0: ((v120 / v120) > (v120 / v120)), F1: 127}
                    _ = v171
                    _ = v172
                    _ = v173
                }
            }
            fmt.Println(v117)
            _ = fi168
        }
    case 4294967295:
        var v174 int64 = v112
        _ = v174
    }
    v175 := GenS11{F0: 255, F1: -1, F2: 68}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v112:", v112)
    fmt.Println("v116:", v116)
    fmt.Println("v117:", v117)
    fmt.Println("v118:", v118)
    fmt.Println("v120:", v120)
    fmt.Println("v163:", v163)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v112
    _ = v116
    _ = v117
    _ = v118
    _ = v119
    _ = v120
    _ = v121
    _ = v163
    _ = v175
    _ = v2
    _ = v3
}
