package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 int8; F1 uint32; F2 uint8 }
type GenS3 struct { F0 uint8; F1 string; F2 int64 }
type GenS4 struct { F0 uint8; F1 int; F2 uint8 }
type GenS5 struct { F0 uint16; F1 int64 }
type GenS6 struct { F0 int32; F1 uint64; F2 string }
type GenS7 struct { F0 uint32; F1 uint64 }
type GenS8 struct { F0 uint8; F1 int8; F2 uint }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 uint8; F1 uint16 }
type GenS12 struct { F0 int32; F1 bool }
func genFn0(p0 uint32, p1 int32, p2 int16) uint {
    for fi4 := 0; fi4 < 19; fi4++ {
        v5 := []int{(fi4 + fi4)}
        _ = fi4
        _ = v5
    }
    var v6 float32 = 1.5
    var v7 int64 = 42
    fmt.Println(v6, v7, v7)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 65535
}
func genFn1(p0 bool) bool {
    v8 := []int{32767, -2147483648, 859, 827}
    var v9 int16 = 218
    _ = p0
    _ = v8
    _ = v9
    return true
}

func main() {
    var v1 uint64 = 1
    v2 := make(map[uint32]int)
    var v3 string = ("foo bar" + "\n\t")
    for fi10 := 0; fi10 < 15; fi10++ {
        if (!false) {
            var v11 bool = (42 > 253)
            fmt.Println(v1, fi10, fi10)
            _ = v11
        } else {
            v12 := GenS1{F0: false, F1: 255}
            v13 := GenS1{F0: true, F1: 106}
            var v14 uint32 = 2147483647
            _ = v12
            _ = v13
            _ = v14
        }
        if (2 < 2) {
            if true {
                v15 := GenS2{F0: -128, F1: 373, F2: 42}
                _ = v15
            }
            if false {
                switch "\n\t" {
                case "":
                    for fi16 := 0; fi16 < 15; fi16++ {
                        v17 := v2[412]
                        _ = fi16
                        _ = v17
                    }
                }
                var v18 string = "\n\t"
                _ = v18
            } else {
                var v19 int8 = -89
                fmt.Println(fi10)
                _ = v19
            }
            var v20 int16 = -1
            _ = v20
        }
        v21 := GenS1{F0: true, F1: 119}
        _ = fi10
        _ = v21
    }
    for fi22 := 0; fi22 < 19; fi22++ {
        if true {
            if (!true) {
                for fi23 := 0; fi23 < 15; fi23++ {
                    var v24 uint8 = 0
                    var v25 uint = genFn0(1, 2147483647, -32768)
                    var v26 int8 = 1
                    _ = fi23
                    _ = v24
                    _ = v25
                    _ = v26
                }
                var v27 bool = (v3 < "hello")
                _ = v27
            }
            if (42 < 4294967295) {
                if false {
                    if false {
                        if (1 == 31) {
                            var v28 uint32 = 65535
                            if (false && (1 == 0)) {
                                var v29 float32 = (-2.25 - float32(float64(0.1)))
                                if ((v3 + "foo bar") <= "hello") {
                                    v30 := []int{(fi22 ^ (fi22 / ((fi22) | 1))), (fi22 >> ((genFn0(v28, 2147483647, 246)) & 63))}
                                    _ = v30
                                }
                                _ = v29
                            }
                            _ = v28
                        } else {
                            v1 = v1
                            for fi31 := 0; fi31 < 11; fi31++ {
                                fmt.Println(fi31, fi31, v1)
                                v32 := GenS2{F0: -123, F1: 2147483648, F2: 255}
                                if true {
                                    if false {
                                        v33 := []int{}
                                        fmt.Println(fi31, fi22)
                                        _ = v33
                                    } else {
                                        var v34 float64 = ((0.0 * -2.25) - (float64(0.0) - 0.1))
                                        _ = v34
                                    }
                                    var v35 uint16 = 255
                                    switch fi22 {
                                    case -1:
                                        if true {
                                            fmt.Println(v3, fi22)
                                            v36 := []int{fi22, (fi31 & (fi22 % ((-1) | 1))), fi22, (^(fi22))}
                                            if false {
                                                if (true && false) {
                                                    _ = len(v36)
                                                    var v37 uint32 = 2147483647
                                                    _ = v37
                                                } else {
                                                    switch (v3 + "\n\t") {
                                                    case "":
                                                        var v38 int = 4294967295
                                                        _ = v38
                                                    default:
                                                        delete(v2, uint32(v1))
                                                    }
                                                    v35 -= v35
                                                    var v39 int16 = 390
                                                    fmt.Println(v39)
                                                    _ = v39
                                                }
                                            }
                                            _ = v36
                                        }
                                    }
                                    _ = v35
                                } else {
                                    v3 = (v3 + ((("" + ("" + "foo bar")) + ("" + "hello")) + v3))
                                    var v40 int16 = -32768
                                    v1 += (v1 & 2)
                                    v41 := GenS2{F0: -1, F1: uint32(v1), F2: 42}
                                    v42 := []int{}
                                    _ = v40
                                    _ = v41
                                    _ = v42
                                }
                                fmt.Println(v1)
                                switch (v3 + (((("" + "foo bar") + "hello") + v3) + "a")) {
                                case "foo bar":
                                    v43 := []int{(-(fi22)), fi31, fi31}
                                    _ = v43
                                case "a":
                                    if (!false) {
                                        v44 := []int{}
                                        switch (v3 + ("" + v3)) {
                                        case "a":
                                            var v45 uint8 = 2
                                            _ = v45
                                        case "hello":
                                            v3 = ((v3 + v3) + v3)
                                        default:
                                            var v46 uint64 = 4294967295
                                            _ = v46
                                        }
                                        _ = v44
                                    } else {
                                        var v47 uint = 18446744073709551615
                                        switch (fi22 << ((v47) & 63)) {
                                        case 9223372036854775807:
                                            var v48 int8 = 0
                                            _ = v48
                                        case -2147483648:
                                            var v49 int32 = -1
                                            _ = v49
                                        case 32767:
                                            var v50 int32 = 0
                                            _ = v50
                                        }
                                        v51 := v2[0]
                                        _ = v47
                                        _ = v51
                                    }
                                case "\n\t":
                                    var v52 uint32 = 4294967295
                                    _ = v52
                                }
                                _ = fi31
                                _ = v32
                            }
                            switch -24 {
                            case 4294967295:
                                v53 := GenS3{F0: 127, F1: v3, F2: -9223372036854775808}
                                _ = v53
                            case -689:
                                v1 = v1
                            default:
                                v54 := GenS2{F0: 42, F1: 2, F2: 2}
                                _ = v54
                            }
                        }
                        if true {
                            v1 += (v1 ^ v1)
                            if (-32768 >= -32768) {
                                v55 := GenS3{F0: 127, F1: v3, F2: 38}
                                var v56 uint64 = (v1 - 434)
                                for fi57 := 0; fi57 < 18; fi57++ {
                                    var v58 uint64 = (v1 % ((v56) | 1))
                                    _ = fi57
                                    _ = v58
                                }
                                _ = v55
                                _ = v56
                            }
                        } else {
                            v1 = v1
                            v59 := []int{-1, fi22}
                            v1 += 1
                            _ = v59
                        }
                    } else {
                        var v60 uint16 = 436
                        _ = v60
                    }
                } else {
                    v1 -= 2
                }
                var v61 int32 = 255
                var v62 uint64 = (v1 | (v1 * v1))
                _ = v61
                _ = v62
            } else {
                v3 = v3
            }
            var v63 int32 = -941
            _ = v63
        } else {
            var v64 uint8 = 0
            v65 := v2[1]
            switch "" {
            case "a":
                v66 := []int{}
                _ = v66
            case "foo bar":
                for fi67 := 0; fi67 < 13; fi67++ {
                    var v68 uint8 = 127
                    _ = len(v2)
                    if (65 > 100) {
                        switch (v65 % (((v65 % ((v65) | 1))) | 1)) {
                        case -855:
                            v1 = 2
                        }
                        fmt.Println(fi22, v3)
                    } else {
                        v69 := GenS4{F0: (v64 << ((1) & 63)), F1: (fi22 & 127), F2: v68}
                        if true {
                            switch (fi22 & fi22) {
                            case 666:
                                v70 := v2[825]
                                _ = v70
                            case -1:
                                var v71 float32 = -2.25
                                _ = v71
                            default:
                                switch v65 {
                                case 127:
                                    fmt.Println(fi22, v65)
                                }
                            }
                            if true {
                                switch (v3 + (v3 + "a")) {
                                case "":
                                    v72 := GenS5{F0: 0, F1: 9223372036854775807}
                                    _ = v72
                                case "\n\t":
                                    v1 -= (^(v1))
                                case "hello":
                                    var v73 int = (fi67 + fi67)
                                    _ = v73
                                default:
                                    fmt.Println(v68, v1, fi67)
                                }
                            }
                            v74 := GenS2{F0: 0, F1: 2, F2: 94}
                            var v75 uint = 4294967295
                            _ = v74
                            _ = v75
                        } else {
                            if false {
                                if (true && ((v65 * v65) >= (fi67 ^ 0))) {
                                    if true {
                                        v76 := GenS6{F0: 0, F1: v1, F2: "hello"}
                                        _ = v76
                                    } else {
                                        switch "hello" {
                                        case "hello":
                                            switch v65 {
                                            case 2147483647:
                                                v1 = (^(v1))
                                            case 1:
                                                v77 := GenS7{F0: uint32(v1), F1: (^(v1))}
                                                _ = v77
                                            }
                                        }
                                        var v78 uint16 = 255
                                        v1 = (v1 + 4294967295)
                                        _ = v78
                                    }
                                    switch (-(fi67)) {
                                    case 4294967295:
                                        v64 += (v64 * v64)
                                    case 32767:
                                        if true {
                                            var v79 uint16 = 761
                                            _ = v79
                                        }
                                    default:
                                        var v80 int8 = 1
                                        _ = v80
                                    }
                                } else {
                                    v81 := v2[0]
                                    v82 := GenS1{F0: false, F1: (v68 % ((v64) | 1))}
                                    _ = v81
                                    _ = v82
                                }
                                fmt.Println(v68, v3)
                                v1 += 0
                                switch (("a" + "foo bar") + ("foo bar" + v3)) {
                                case "a":
                                    _ = len(v2)
                                }
                            }
                        }
                        _ = v69
                    }
                    _ = fi67
                    _ = v68
                }
            default:
                if false {
                    for fi83 := 0; fi83 < 18; fi83++ {
                        fmt.Println(v3, v3, v65)
                        v3 = v3
                        v84 := []int{(fi22 ^ fi22), (v65 - (fi83 ^ (^(fi83)))), (^(fi22))}
                        var v85 int64 = 1
                        _ = fi83
                        _ = v84
                        _ = v85
                    }
                    switch v65 {
                    case 4294967295:
                        if ((^(fi22)) > (fi22 & v65)) {
                            if false {
                                for fi86 := 0; fi86 < 7; fi86++ {
                                    var v87 float64 = 3.14159
                                    v1 = (v1 - 18446744073709551615)
                                    if ((!(675 <= 832)) && true) {
                                        v64 = 2
                                    }
                                    _ = fi86
                                    _ = v87
                                }
                                v3 = ("\n\t" + "\n\t")
                            } else {
                                if (!(!(!true))) {
                                    for fi88 := 0; fi88 < 9; fi88++ {
                                        if ((float32(1e10) - 1e10) != 1e10) {
                                            var v89 int8 = 127
                                            v64 = v64
                                            if true {
                                                var v90 int32 = 1
                                                _ = v90
                                            } else {
                                                var v91 uint = 11
                                                v3 = ""
                                                _ = v91
                                            }
                                            _ = v89
                                        } else {
                                            v65 += v65
                                            var v92 uint16 = 0
                                            _ = v92
                                        }
                                        if true {
                                            v3 = ""
                                        }
                                        _ = fi88
                                    }
                                    v93 := GenS3{F0: 255, F1: (v3 + "hello"), F2: -677}
                                    _ = v93
                                }
                                v94 := GenS4{F0: 79, F1: fi22, F2: 128}
                                var v95 uint64 = 1
                                _ = v94
                                _ = v95
                            }
                            var v96 uint = genFn0(1, -1, 1000)
                            for fi97 := 0; fi97 < 15; fi97++ {
                                fmt.Println(v64, v1)
                                fmt.Println(fi22, v96)
                                v1 += 2
                                _ = fi97
                            }
                            _ = v96
                        } else {
                            if true {
                                v98 := []int{(-(fi22)), (v65 * fi22)}
                                if (true && true) {
                                    var v99 string = ("" + v3)
                                    _ = v99
                                }
                                var v100 string = "foo bar"
                                _ = v100
                                _ = v98
                            } else {
                                v3 = v3
                            }
                            if false {
                                var v101 uint64 = (v1 | v1)
                                _ = v101
                            } else {
                                var v102 int32 = -723
                                if false {
                                    delete(v2, 4294967295)
                                    v65 = fi22
                                } else {
                                    switch fi22 {
                                    case 1:
                                        for fi103 := 0; fi103 < 10; fi103++ {
                                            fmt.Println(v3, v102)
                                            if (false || true) {
                                                v102 = v102
                                            } else {
                                                v1 = (v1 / ((0) | 1))
                                                switch fi22 {
                                                case 127:
                                                    delete(v2, 164)
                                                case -9223372036854775808:
                                                    if (-433 < -347) {
                                                        v104 := v2[65535]
                                                        var v105 float32 = (0.1 + 3.14159)
                                                        _ = v104
                                                        _ = v105
                                                    } else {
                                                        v2[0] = v65
                                                        v106 := v2[4294967295]
                                                        var v107 uint8 = (v64 & 204)
                                                        if (4294967295 <= (v65 + -953)) {
                                                            var v108 string = ("foo bar" + (((v3 + "hello") + v3) + v3))
                                                            var v109 int64 = -1
                                                            _ = v108
                                                            _ = v109
                                                        }
                                                        if true {
                                                            var v110 uint8 = v107
                                                            var v111 float64 = 0.0
                                                            _ = v110
                                                            _ = v111
                                                        } else {
                                                            v112 := GenS8{F0: v64, F1: 1, F2: 18446744073709551615}
                                                            v113 := []int{-29, (fi103 / ((v106) | 1))}
                                                            _ = v112
                                                            _ = v113
                                                        }
                                                        _ = v106
                                                        _ = v107
                                                    }
                                                    v114 := []int{fi22}
                                                    var v115 uint16 = 42
                                                    _ = v114
                                                    _ = v115
                                                case 32768:
                                                    if false {
                                                        if (false && true) {
                                                            if (-1 >= -4) {
                                                                v116 := []int{}
                                                                _ = v116
                                                            } else {
                                                                var v117 uint64 = 42
                                                                _ = v117
                                                            }
                                                        }
                                                        delete(v2, 42)
                                                    } else {
                                                        switch (fi103 - (v65 + fi103)) {
                                                        case 2147483647:
                                                            v118 := GenS9{F0: 347}
                                                            var v119 int8 = -25
                                                            _ = v118
                                                            _ = v119
                                                        default:
                                                            v120 := []int{(v65 ^ v65), (fi103 << ((18446744073709551615) & 63)), (-(fi22)), (fi22 / ((422) | 1))}
                                                            _ = v120
                                                        }
                                                    }
                                                }
                                            }
                                            v121 := v2[2147483648]
                                            v122 := GenS8{F0: v64, F1: 127, F2: 255}
                                            v3 = "foo bar"
                                            var v123 int8 = -128
                                            _ = fi103
                                            _ = v121
                                            _ = v122
                                            _ = v123
                                        }
                                    case 42:
                                        for fi124 := 0; fi124 < 9; fi124++ {
                                            v125 := GenS1{F0: (9223372036854775807 > (v65 & v65)), F1: v64}
                                            _ = fi124
                                            _ = v125
                                        }
                                    case 558:
                                        if (4294967295 <= 0) {
                                            v65 -= v65
                                            if false {
                                                for fi126 := 0; fi126 < 20; fi126++ {
                                                    v1 += v1
                                                    var v127 int32 = v102
                                                    v3 = (v3 + v3)
                                                    v2[uint32(v64)] = v65
                                                    _ = fi126
                                                    _ = v127
                                                }
                                            }
                                        } else {
                                            v1 = (^(v1))
                                            var v128 int16 = 1000
                                            v129 := GenS10{F0: 95}
                                            _ = v128
                                            _ = v129
                                        }
                                    default:
                                        if ((!true) && ((0.0 == 3.14159) || true)) {
                                            var v130 uint64 = v1
                                            if true {
                                                var v131 int = 0
                                                var v132 bool = true
                                                for fi133 := 0; fi133 < 14; fi133++ {
                                                    var v134 uint64 = v1
                                                    if false {
                                                        v102 -= 0
                                                        v130 += 4294967295
                                                    }
                                                    switch (v3 + "foo bar") {
                                                    case "hello":
                                                        fmt.Println(v132, fi133)
                                                    case "a":
                                                        v134 += (v130 << ((0) & 63))
                                                    default:
                                                        var v135 int8 = 49
                                                        _ = v135
                                                    }
                                                    _ = fi133
                                                    _ = v134
                                                }
                                                _ = v131
                                                _ = v132
                                            } else {
                                                if (false && true) {
                                                    var v136 string = (v3 + (v3 + "a"))
                                                    if (true && (((false && (655 >= -32768)) || (0.1 != (3.14159 / 1.0))) && ((v1 ^ 2) < 2))) {
                                                        fmt.Println(fi22, v65, v1)
                                                        for fi137 := 0; fi137 < 17; fi137++ {
                                                            var v138 int16 = -1
                                                            _ = len(v2)
                                                            v65 += v65
                                                            _ = fi137
                                                            _ = v138
                                                        }
                                                        v1 += (v130 << ((4294967295) & 63))
                                                    }
                                                    if false {
                                                        v1 = (v130 - (v130 | (v1 ^ v130)))
                                                        switch (v136 + v3) {
                                                        case "a":
                                                            var v139 uint32 = 1
                                                            _ = v139
                                                        case "\n\t":
                                                            switch v3 {
                                                            case "\n\t":
                                                                if true {
                                                                    fmt.Println(fi22)
                                                                } else {
                                                                    v64 = (v64 / ((127) | 1))
                                                                    _ = len(v2)
                                                                }
                                                            case "":
                                                                var v140 int64 = -76
                                                                _ = v140
                                                            default:
                                                                v102 = int32(v1)
                                                            }
                                                        }
                                                    }
                                                    _ = v136
                                                } else {
                                                    v141 := []int{4294967295, 1, 32767}
                                                    if true {
                                                        for fi142 := 0; fi142 < 12; fi142++ {
                                                            _ = len(v141)
                                                            delete(v2, 65535)
                                                            v143 := GenS2{F0: 42, F1: 2147483647, F2: v64}
                                                            v102 += v102
                                                            _ = fi142
                                                            _ = v143
                                                        }
                                                    } else {
                                                        v2[4294967295] = (fi22 << ((4294967295) & 63))
                                                        var v144 uint64 = v1
                                                        for fi145 := 0; fi145 < 19; fi145++ {
                                                            var v146 bool = true
                                                            _ = fi145
                                                            _ = v146
                                                        }
                                                        _ = v144
                                                    }
                                                    _ = v141
                                                }
                                            }
                                            var v147 int16 = 42
                                            _ = v130
                                            _ = v147
                                        } else {
                                            v148 := GenS5{F0: 216, F1: -9223372036854775808}
                                            v2[2] = 0
                                            var v149 uint8 = (v64 * v64)
                                            _ = v148
                                            _ = v149
                                        }
                                    }
                                    var v150 bool = (false || (((v3 + v3) <= "\n\t") || true))
                                    _ = v150
                                }
                                _ = v102
                            }
                        }
                    case 5:
                        for fi151 := 0; fi151 < 10; fi151++ {
                            fmt.Println(fi151)
                            var v152 uint16 = 65535
                            v153 := []int{-564, 42, 1, 1}
                            _ = fi151
                            _ = v152
                            _ = v153
                        }
                    case -301:
                        v1 += (^(v1))
                    default:
                        v1 = (^(v1))
                    }
                    var v154 float32 = -2.25
                    _ = v154
                } else {
                    for fi155 := 0; fi155 < 1; fi155++ {
                        if ((v1 / ((690) | 1)) != v1) {
                            v156 := GenS11{F0: v64, F1: 42}
                            var v157 float32 = 0.0
                            _ = v156
                            _ = v157
                        }
                        fmt.Println(v3, v1, fi22)
                        switch v3 {
                        case "foo bar":
                            var v158 int16 = -1
                            _ = v158
                        case "a":
                            v159 := []int{int(v64), fi22}
                            _ = v159
                        default:
                            v64 = v64
                        }
                        _ = fi155
                    }
                    var v160 float64 = 0.0
                    v161 := GenS12{F0: 958, F1: true}
                    _ = v160
                    _ = v161
                }
            }
            _ = v64
            _ = v65
        }
        v162 := GenS9{F0: 65535}
        _ = fi22
        _ = v162
    }
    var v163 string = v3
    if true {
        for fi164 := 0; fi164 < 19; fi164++ {
            v2[4294967295] = (fi164 + fi164)
            v163 = (v3 + v3)
            _ = fi164
        }
    } else {
        for fi165 := 0; fi165 < 19; fi165++ {
            if false {
                var v166 string = v163
                if false {
                    v3 = ("hello" + "foo bar")
                }
                _ = v166
            }
            if true {
                if true {
                    for fi167 := 0; fi167 < 14; fi167++ {
                        v168 := []int{(fi167 << ((728) & 63))}
                        var v169 uint64 = (v1 % ((v1) | 1))
                        _ = fi167
                        _ = v168
                        _ = v169
                    }
                    v2[1] = (fi165 ^ (fi165 << ((42) & 63)))
                    var v170 int16 = -32768
                    _ = v170
                } else {
                    v1 += v1
                    v1 = v1
                    switch fi165 {
                    case 2147483647:
                        v171 := GenS11{F0: 1, F1: 2}
                        _ = v171
                    case 98:
                        v172 := []int{(^(fi165))}
                        _ = v172
                    default:
                        for fi173 := 0; fi173 < 4; fi173++ {
                            fmt.Println(v1, v163)
                            v3 = ("foo bar" + v163)
                            _ = fi173
                        }
                    }
                }
                for fi174 := 0; fi174 < 13; fi174++ {
                    var v175 bool = (true || false)
                    fmt.Println(fi174, v3)
                    var v176 bool = (!v175)
                    _ = fi174
                    _ = v175
                    _ = v176
                }
            }
            _ = fi165
        }
    }
    v3 = "hello"
    v1 -= (v1 & (^(v1)))
    v1 -= (v1 % ((v1) | 1))
    if true {
        fmt.Println(v3, v163)
    }
    v177 := GenS4{F0: 196, F1: 123, F2: 2}
    v178 := []int{1}
    v1 -= 0
    fmt.Println("v1:", v1)
    fmt.Println("v163:", v163)
    fmt.Println("slicelen v178:", len(v178), v178)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v163
    _ = v177
    _ = v178
    _ = v2
    _ = v3
}
