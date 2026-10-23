package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 uint; F1 int8; F2 int32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int; F1 int64 }
type GenS6 struct { F0 uint64; F1 int8; F2 int64 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int; F1 uint16 }
type GenS9 struct { F0 uint64; F1 uint16 }
type GenS10 struct { F0 int64; F1 string }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 uint8; F1 string; F2 int16 }
type GenS14 struct { F0 uint; F1 uint64; F2 string }
func genFn0() uint8 {
    if (2 >= 42) {
        var v4 uint8 = 1
        _ = v4
    } else {
        if false {
            v5 := GenS1{F0: 42}
            v6 := GenS1{F0: 451}
            _ = v5
            _ = v6
        }
    }
    switch ((("a" + ("a" + "")) + "hello") + "foo bar") {
    case "hello":
        var v7 string = "a"
        _ = v7
    default:
        switch -283 {
        case -2147483648:
            if true {
                if false {
                    var v8 int32 = -23
                    _ = v8
                } else {
                    var v9 int64 = -11
                    var v10 int64 = v9
                    _ = v10
                    _ = v9
                }
            }
        default:
            switch 2147483647 {
            case -914:
                switch 1 {
                case -1:
                    switch ("a" + "\n\t") {
                    case "foo bar":
                        v11 := GenS2{F0: 42, F1: 103}
                        _ = v11
                    case "\n\t":
                        if ((55 == -128) || false) {
                            var v12 float32 = (0.1 - (-2.25 - 0.0))
                            switch "\n\t" {
                            case "a":
                                for fi13 := 0; fi13 < 19; fi13++ {
                                    for fi14 := 0; fi14 < 12; fi14++ {
                                        v12 += v12
                                        _ = fi14
                                    }
                                    for fi15 := 0; fi15 < 18; fi15++ {
                                        var v16 float64 = float64(0.0)
                                        _ = fi15
                                        _ = v16
                                    }
                                    if (255 >= 18446744073709551615) {
                                        var v17 int = fi13
                                        _ = v17
                                    } else {
                                        v18 := GenS3{F0: 805, F1: -1, F2: 219}
                                        v19 := []int{(fi13 + (fi13 << ((65535) & 63)))}
                                        _ = v18
                                        _ = v19
                                    }
                                    _ = fi13
                                }
                            case "hello":
                                fmt.Println(v12)
                            case "foo bar":
                                var v20 int16 = 1000
                                _ = v20
                            }
                            _ = v12
                        } else {
                            if false {
                                v21 := []int{42, 946, 127}
                                for fi22 := 0; fi22 < 3; fi22++ {
                                    var v23 int8 = 24
                                    _ = fi22
                                    _ = v23
                                }
                                var v24 float32 = -2.25
                                _ = v21
                                _ = v24
                            } else {
                                var v25 int32 = -2147483648
                                switch "hello" {
                                case "foo bar":
                                    var v26 int = 32768
                                    _ = v26
                                default:
                                    if (2 >= 42) {
                                        v27 := []int{32768, 4294967295, -171}
                                        _ = v27
                                    } else {
                                        v25 = v25
                                        switch -1 {
                                        case 32768:
                                            v28 := []int{714}
                                            _ = v28
                                        default:
                                            v29 := []int{}
                                            _ = v29
                                        }
                                        var v30 uint = 4294967295
                                        _ = v30
                                    }
                                }
                                _ = v25
                            }
                            var v31 uint = 428
                            _ = v31
                        }
                    case "hello":
                        var v32 int32 = -1
                        _ = v32
                    }
                default:
                    switch 32767 {
                    case -683:
                        var v33 int8 = 100
                        _ = v33
                    case 0:
                        var v34 float64 = float64(0.1)
                        _ = v34
                    case -109:
                        v35 := []int{-404, 1, 729, -2147483648}
                        _ = v35
                    default:
                        var v36 float32 = 0.1
                        _ = v36
                    }
                }
            case 128:
                switch -499 {
                case 32767:
                    var v37 int = 1
                    _ = v37
                default:
                    v38 := GenS4{F0: 450}
                    _ = v38
                }
            case 690:
                if false {
                    switch "a" {
                    case "a":
                        if (-2147483648 < -616) {
                            if true {
                                for fi39 := 0; fi39 < 1; fi39++ {
                                    v40 := []int{(fi39 - (^(fi39))), (fi39 | fi39)}
                                    var v41 uint = 0
                                    _ = fi39
                                    _ = v40
                                    _ = v41
                                }
                                var v42 int64 = -1
                                var v43 int32 = 0
                                _ = v42
                                _ = v43
                            }
                            v44 := []int{-107, -2147483648}
                            if true {
                                if (255 < 42) {
                                    var v45 uint16 = 2
                                    _ = v45
                                } else {
                                    var v46 int16 = -32768
                                    v46 += (v46 + (v46 * v46))
                                    var v47 string = ("a" + ("a" + "foo bar"))
                                    _ = v46
                                    _ = v47
                                }
                                var v48 uint32 = 2
                                if len(v44) >= 2 {
                                    v44 = v44[1:len(v44)]
                                }
                                _ = v48
                            }
                            _ = v44
                        } else {
                            switch "" {
                            case "hello":
                                if false {
                                    for fi49 := 0; fi49 < 6; fi49++ {
                                        var v50 int64 = 1
                                        _ = fi49
                                        _ = v50
                                    }
                                    var v51 int64 = -414
                                    v51 -= (v51 % ((v51) | 1))
                                    _ = v51
                                } else {
                                    var v52 float64 = 0.0
                                    _ = v52
                                }
                            case "a":
                                switch 128 {
                                case 0:
                                    var v53 uint32 = 2
                                    _ = v53
                                default:
                                    v54 := []int{}
                                    _ = v54
                                }
                            case "":
                                var v55 int = 899
                                _ = v55
                            }
                        }
                    case "foo bar":
                        var v56 int32 = 2147483647
                        _ = v56
                    }
                } else {
                    switch 2147483647 {
                    case -9223372036854775808:
                        var v57 uint16 = 32767
                        _ = v57
                    default:
                        if false {
                            v58 := GenS4{F0: 742}
                            if true {
                                var v59 int8 = 0
                                v60 := GenS5{F0: -1, F1: -1}
                                for fi61 := 0; fi61 < 14; fi61++ {
                                    var v62 bool = true
                                    _ = fi61
                                    _ = v62
                                }
                                _ = v59
                                _ = v60
                            }
                            var v63 int = -290
                            _ = v58
                            _ = v63
                        } else {
                            var v64 string = ("a" + ((("a" + "a") + "") + "foo bar"))
                            v65 := []int{1, 9223372036854775807}
                            _ = v64
                            _ = v65
                        }
                    }
                    v66 := []int{32768}
                    _ = v66
                }
            default:
                var v67 uint = 720
                _ = v67
            }
        }
    }
    return 0
}
func genFn1(p0 uint8, p1 bool, p2 bool) int {
    var v68 int16 = -878
    var v69 int = 128
    if (0.1 <= float64(3.14159)) {
        switch (v69 / (((^(v69))) | 1)) {
        case 127:
            p2 = false
        case 2147483647:
            p0 += p0
        case -2147483648:
            v70 := GenS3{F0: 42, F1: -128, F2: -108}
            _ = v70
        default:
            if (42 < 42) {
                v71 := []int{v69}
                _ = v71
            }
        }
    } else {
        if p2 {
            v72 := GenS4{F0: 2}
            v73 := []int{0, (v69 >> ((18446744073709551615) & 63))}
            var v74 float64 = 0.0
            _ = v72
            _ = v73
            _ = v74
        } else {
            p0 += 2
        }
        p0 = p0
        p0 += 217
    }
    var v75 int64 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v68
    _ = v69
    _ = v75
    return 0
}

func main() {
    var v1 bool = (true && false)
    var v2 int64 = 1
    var v3 uint = 65535
    switch "\n\t" {
    case "foo bar":
        var v76 string = "foo bar"
        _ = v76
    }
    fmt.Println(v2)
    switch ("hello" + "hello") {
    case "\n\t":
        var v77 string = ("hello" + "\n\t")
        _ = v77
    case "hello":
        var v78 uint = (^(v3))
        _ = v78
    case "foo bar":
        v3 = v3
    }
    if (748 >= 255) {
        if ("\n\t" < "") {
            if ((v3 >> ((342) & 63)) > (^(v3))) {
                v2 = 1
                if true {
                    var v79 uint16 = 65535
                    v80 := GenS2{F0: 806, F1: 52}
                    v81 := []int{201, 1, 361}
                    _ = v79
                    _ = v80
                    _ = v81
                } else {
                    for fi82 := 0; fi82 < 9; fi82++ {
                        var v83 int = 0
                        if v1 {
                            var v84 bool = false
                            if v1 {
                                var v85 bool = (!(v2 != 1))
                                _ = v85
                            } else {
                                v2 += v2
                                if (v1 && (42 != 980)) {
                                    fmt.Println(fi82)
                                    var v86 int8 = -73
                                    if v84 {
                                        var v87 int8 = v86
                                        _ = v87
                                    }
                                    _ = v86
                                } else {
                                    if (!v84) {
                                        v2 += v2
                                        for fi88 := 0; fi88 < 10; fi88++ {
                                            v2 -= v2
                                            var v89 uint64 = 2
                                            switch genFn1(127, (2147483648 < 2147483647), (v1 || v1)) {
                                            case 2147483647:
                                                v1 = v1
                                            case 32767:
                                                fmt.Println(v3, fi82, v2)
                                            case -371:
                                                v1 = v1
                                            default:
                                                var v90 float32 = 1.5
                                                _ = v90
                                            }
                                            v91 := []int{(fi88 | (v83 | (v83 >> ((372) & 63))))}
                                            v92 := GenS6{F0: v89, F1: 127, F2: v2}
                                            _ = fi88
                                            _ = v89
                                            _ = v91
                                            _ = v92
                                        }
                                        v1 = (1 < 218)
                                    } else {
                                        v93 := GenS4{F0: 835}
                                        if v1 {
                                            if (!v84) {
                                                v3 = v3
                                                v94 := GenS6{F0: 4294967295, F1: -22, F2: (v2 * (v2 + v2))}
                                                if (!v84) {
                                                    v95 := GenS4{F0: 1}
                                                    _ = v95
                                                }
                                                _ = v94
                                            }
                                        }
                                        _ = v93
                                    }
                                    var v96 int16 = 42
                                    var v97 uint64 = 943
                                    _ = v96
                                    _ = v97
                                }
                                var v98 bool = ((!(3.14159 < 0.1)) && v84)
                                _ = v98
                            }
                            _ = v84
                        }
                        _ = fi82
                        _ = v83
                    }
                }
                var v99 int8 = 42
                _ = v99
            } else {
                fmt.Println(v2)
            }
        }
        fmt.Println(v1, v2, v2)
        v2 -= 924
    } else {
        var v100 int16 = 32767
        var v101 string = "foo bar"
        var v102 bool = (1 <= 2)
        _ = v100
        _ = v101
        _ = v102
    }
    if (905 < v2) {
        fmt.Println(v2)
    } else {
        fmt.Println(v2, v2)
        for fi103 := 0; fi103 < 20; fi103++ {
            if false {
                if (70 <= 58) {
                    switch "foo bar" {
                    case "\n\t":
                        v104 := []int{761}
                        _ = v104
                    case "":
                        var v105 int16 = -32768
                        _ = v105
                    case "hello":
                        v1 = v1
                    }
                    v106 := []int{}
                    _ = v106
                }
            }
            switch "foo bar" {
            case "a":
                for fi107 := 0; fi107 < 6; fi107++ {
                    v108 := GenS2{F0: fi103, F1: 42}
                    fmt.Println(fi107, v1)
                    switch "" {
                    case "":
                        v3 = (^(v3))
                    case "a":
                        var v109 uint = (v3 + (v3 << ((0) & 63)))
                        _ = v109
                    }
                    if false {
                        v110 := GenS7{F0: 692}
                        var v111 uint16 = 42
                        if v1 {
                            var v112 uint = v3
                            switch fi107 {
                            case 0:
                                fmt.Println(v111, v111)
                            case 32767:
                                v1 = true
                            default:
                                var v113 int32 = 0
                                _ = v113
                            }
                            _ = v112
                        } else {
                            v114 := []int{fi107, (fi103 & (fi107 ^ (-(fi107)))), genFn1(127, (((!false) || v1) && v1), false)}
                            var v115 uint = uint(v111)
                            v116 := GenS1{F0: 1}
                            _ = v114
                            _ = v115
                            _ = v116
                        }
                        _ = v110
                        _ = v111
                    } else {
                        if v1 {
                            var v117 uint = v3
                            switch "" {
                            case "hello":
                                if (-1 <= v2) {
                                    if v1 {
                                        v118 := []int{fi103, -197}
                                        _ = v118
                                    }
                                    v3 -= v117
                                }
                            default:
                                v3 = (v117 - 805)
                            }
                            _ = v117
                        } else {
                            if v1 {
                                if v1 {
                                    var v119 int64 = v2
                                    var v120 uint16 = 2
                                    var v121 int = fi107
                                    _ = v119
                                    _ = v120
                                    _ = v121
                                } else {
                                    v2 = v2
                                    var v122 uint64 = 2
                                    _ = v122
                                }
                                var v123 float32 = -2.25
                                _ = v123
                            }
                            v124 := GenS5{F0: fi103, F1: (^(v2))}
                            _ = v124
                        }
                        v125 := []int{fi103, 2147483647, 27, (fi107 | genFn1(157, v1, (int16(fi103) == 32767)))}
                        _ = v125
                    }
                    fmt.Println(fi107, fi103, v2)
                    if ((^(v3)) < v3) {
                        switch "a" {
                        case "foo bar":
                            v126 := GenS8{F0: (fi107 ^ -2147483648), F1: 0}
                            _ = v126
                        case "a":
                            fmt.Println(v3)
                        case "hello":
                            var v127 float64 = (3.14159 / 1.0)
                            _ = v127
                        }
                        v128 := []int{(fi107 + fi107), genFn1(127, v1, ((815 > 579) && (true || ("hello" >= "\n\t")))), genFn1(127, v1, v1)}
                        _ = v128
                    }
                    _ = fi107
                    _ = v108
                }
            case "":
                v3 = v3
            }
            if (v1 && (-2147483648 > -95)) {
                v1 = true
                switch fi103 {
                case 0:
                    v129 := GenS9{F0: 394, F1: 65535}
                    _ = v129
                case 4294967295:
                    v130 := GenS2{F0: 32768, F1: 76}
                    _ = v130
                }
                if (97 != 127) {
                    for fi131 := 0; fi131 < 18; fi131++ {
                        v132 := GenS2{F0: (fi103 / (((fi131 | (^(fi103)))) | 1)), F1: 2}
                        switch (fi131 / ((genFn1(118, v1, v1)) | 1)) {
                        case 128:
                            fmt.Println(v2)
                        case -1:
                            v133 := GenS9{F0: 2, F1: 2}
                            v1 = true
                            _ = v133
                        case 32767:
                            v134 := []int{}
                            _ = v134
                        }
                        var v135 uint32 = 65535
                        v136 := GenS3{F0: (v3 | 148), F1: 0, F2: -933}
                        _ = fi131
                        _ = v132
                        _ = v135
                        _ = v136
                    }
                    var v137 int64 = (v2 & (v2 ^ -1))
                    _ = v137
                } else {
                    v138 := []int{genFn1(1, v1, v1), fi103, fi103}
                    _ = v138
                }
            }
            _ = fi103
        }
        if (!v1) {
            for fi139 := 0; fi139 < 19; fi139++ {
                fmt.Println(v3, v3)
                switch "\n\t" {
                case "":
                    var v140 int32 = 423
                    _ = v140
                default:
                    v3 -= v3
                }
                if ("\n\t" >= "\n\t") {
                    var v141 int32 = 571
                    _ = v141
                } else {
                    if v1 {
                        var v142 uint64 = 18446744073709551615
                        _ = v142
                    } else {
                        v143 := GenS3{F0: v3, F1: 88, F2: -2147483648}
                        if ((-(fi139)) > genFn1(genFn0(), true, v1)) {
                            v144 := []int{fi139, fi139, (fi139 | (^(fi139))), fi139}
                            _ = v144
                        } else {
                            v145 := []int{128}
                            v3 -= 4294967295
                            _ = v145
                        }
                        if v1 {
                            v146 := []int{42, (fi139 - (fi139 * (-(fi139))))}
                            for fi147 := 0; fi147 < 4; fi147++ {
                                v1 = v1
                                var v148 int16 = -32768
                                _ = fi147
                                _ = v148
                            }
                            var v149 float64 = 1e10
                            _ = v146
                            _ = v149
                        } else {
                            if ((("foo bar" + "\n\t") + "foo bar") <= "foo bar") {
                                v150 := []int{fi139, fi139, -732, (fi139 % ((127) | 1))}
                                v151 := GenS2{F0: -969, F1: 2}
                                switch "foo bar" {
                                case "a":
                                    v152 := GenS9{F0: 1, F1: 65535}
                                    _ = v152
                                case "\n\t":
                                    if v1 {
                                        switch "" {
                                        case "hello":
                                            var v153 int32 = 498
                                            _ = v153
                                        default:
                                            if ((v1 || (true && (42 == 242))) || false) {
                                                for fi154 := 0; fi154 < 2; fi154++ {
                                                    switch genFn1(255, true, (v1 || false)) {
                                                    case -9223372036854775808:
                                                        var v155 float64 = 0.1
                                                        _ = v155
                                                    default:
                                                        v156 := GenS10{F0: (v2 / ((v2) | 1)), F1: (("" + "") + "hello")}
                                                        _ = v156
                                                    }
                                                    _ = fi154
                                                }
                                                v2 = v2
                                                fmt.Println(fi139, fi139)
                                            }
                                        }
                                    } else {
                                        fmt.Println(v3)
                                        fmt.Println(v1, fi139)
                                        var v157 int8 = 7
                                        _ = v157
                                    }
                                }
                                _ = v150
                                _ = v151
                            } else {
                                for fi158 := 0; fi158 < 13; fi158++ {
                                    if (1 <= 2) {
                                        var v159 int32 = -1
                                        v160 := GenS11{F0: -128}
                                        fmt.Println(v2, v1)
                                        _ = v159
                                        _ = v160
                                    } else {
                                        var v161 float64 = ((float64(-2.25) / (((float64(float32(float64(-2.25))) + 0.1) + 0.0) - 1.5)) / 1.5)
                                        switch genFn1(25, ((v2 << ((2) & 63)) <= 9223372036854775807), false) {
                                        case -78:
                                            v161 = float64(((0.1 / 3.14159) + float32(v161)))
                                        case 9223372036854775807:
                                            fmt.Println(fi139)
                                        }
                                        _ = v161
                                    }
                                    v1 = false
                                    v162 := GenS1{F0: 1}
                                    _ = fi158
                                    _ = v162
                                }
                            }
                            fmt.Println(v2, v2, v1)
                        }
                        _ = v143
                    }
                    v1 = false
                    for fi163 := 0; fi163 < 20; fi163++ {
                        if (!((^(v2)) != 9223372036854775807)) {
                            v164 := []int{(fi139 << ((v3) & 63)), genFn1(0, v1, ((0.1 + 0.0) != 1e10)), -545, 1}
                            v2 += v2
                            fmt.Println(v3, v3, v3)
                            _ = v164
                        } else {
                            var v165 uint8 = 128
                            var v166 int16 = 340
                            _ = v165
                            _ = v166
                        }
                        v1 = v1
                        fmt.Println(v2)
                        _ = fi163
                    }
                }
                _ = fi139
            }
            for fi167 := 0; fi167 < 12; fi167++ {
                v1 = true
                var v168 uint = v3
                _ = fi167
                _ = v168
            }
            var v169 int16 = 1
            _ = v169
        } else {
            if v1 {
                for fi170 := 0; fi170 < 18; fi170++ {
                    var v171 int16 = 386
                    v172 := []int{(fi170 * genFn1(0, (3.14159 > (((((float64(0.0) - float64(((0.0 + 3.14159) * (-2.25 * 0.0)))) - float64((3.14159 / -2.25))) / float64(float32(float64(1e10)))) - -2.25) / 3.14159)), v1)), (^(fi170)), (^(fi170)), 42}
                    switch genFn1(1, (!v1), v1) {
                    case 1:
                        v173 := GenS12{F0: 88}
                        _ = v173
                    case 32767:
                        v174 := GenS5{F0: (fi170 % (((fi170 + genFn1(127, ((^(v2)) <= v2), (!((v171 > v171) || (0 != fi170)))))) | 1)), F1: v2}
                        _ = v174
                    default:
                        if len(v172) >= 2 {
                            v172 = v172[1:len(v172)]
                        }
                    }
                    _ = fi170
                    _ = v171
                    _ = v172
                }
                v175 := GenS3{F0: v3, F1: 39, F2: 1}
                _ = v175
            } else {
                for fi176 := 0; fi176 < 4; fi176++ {
                    fmt.Println(fi176)
                    v177 := []int{}
                    _ = fi176
                    _ = v177
                }
                switch genFn1(2, ((float32(0.1) - float32(0.0)) != -2.25), (v1 || (4294967295 != 980))) {
                case -883:
                    v178 := []int{genFn1(173, v1, v1), -725, 127, -1}
                    _ = v178
                case 32767:
                    v179 := []int{127}
                    _ = v179
                default:
                    v180 := GenS8{F0: 2147483647, F1: 32767}
                    _ = v180
                }
            }
            var v181 int64 = (v2 >> (((v3 / ((339) | 1))) & 63))
            var v182 int = genFn1(0, (v1 && v1), true)
            _ = v181
            _ = v182
        }
    }
    var v183 float32 = -2.25
    var v184 int8 = -1
    v3 += v3
    fmt.Println(v184, v183, v1)
    for fi185 := 0; fi185 < 7; fi185++ {
        v186 := []int{fi185}
        _ = fi185
        _ = v186
    }
    v187 := GenS13{F0: 1, F1: "hello", F2: -1}
    if true {
        if v1 {
            v2 -= (v2 % ((-1) | 1))
            var v188 int16 = 1000
            v3 -= 65535
            _ = v188
        }
        v189 := GenS14{F0: 1, F1: 1, F2: "\n\t"}
        _ = v189
    }
    var v190 float32 = (v183 - 0.0)
    fmt.Println("v1:", v1)
    fmt.Println("v183:", v183)
    fmt.Println("v184:", v184)
    fmt.Println("v190:", v190)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v183
    _ = v184
    _ = v187
    _ = v190
    _ = v2
    _ = v3
}
