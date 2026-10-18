package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 uint8; F2 uint8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint64; F1 uint64 }
type GenS7 struct { F0 uint32; F1 int }
type GenS8 struct { F0 int8; F1 int64; F2 uint8 }
type GenS9 struct { F0 uint32; F1 int; F2 uint }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 int64; F1 int8 }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 uint; F1 uint32; F2 uint8 }
type GenS14 struct { F0 uint8; F1 uint8; F2 uint }
type GenS15 struct { F0 int32; F1 int64; F2 int32 }
type GenS16 struct { F0 uint64; F1 int64 }
type GenS17 struct { F0 uint16 }
type GenS18 struct { F0 uint64 }
type GenS19 struct { F0 uint16; F1 int32 }
type GenS20 struct { F0 uint16; F1 uint8 }
type GenS21 struct { F0 string; F1 int16 }
type GenS22 struct { F0 int64 }
func genFn0(p0 uint64, p1 uint16, p2 int64) int32 {
    for fi7 := 0; fi7 < 18; fi7++ {
        p1 += p1
        _ = fi7
    }
    v8 := GenS1{F0: -1}
    for fi9 := 0; fi9 < 3; fi9++ {
        if (!true) {
            p1 += p1
        }
        _ = fi9
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        for fi11 := 0; fi11 < 6; fi11++ {
            if (!false) {
                if ((uint8(p0) >= 255) && false) {
                    switch 2147483647 {
                    case -1:
                        if (true && true) {
                            p2 += (-(p2))
                            p2 += p2
                        }
                    case 2147483647:
                        fmt.Println(p2)
                    case -867:
                        p2 -= p2
                    default:
                        if (float64(1e10) > ((3.14159 + 0.1) / 1.0)) {
                            if (true || (502 <= 0)) {
                                var v12 int32 = 725
                                if false {
                                    if true {
                                        if true {
                                            p2 -= (p2 % ((p2) | 1))
                                            var v13 uint32 = 1
                                            v14 := []int{(fi11 ^ 32768), 0, (-(fi10))}
                                            _ = v13
                                            _ = v14
                                        }
                                    } else {
                                        var v15 string = "foo bar"
                                        v16 := GenS2{F0: 276, F1: 195, F2: 42}
                                        _ = v15
                                        _ = v16
                                    }
                                    var v17 uint32 = 1
                                    _ = v17
                                }
                                var v18 bool = true
                                _ = v12
                                _ = v18
                            }
                        }
                    }
                    p0 += 571
                    if false {
                        p0 -= 42
                    } else {
                        p1 -= (p1 & p1)
                        p1 = (^(p1))
                    }
                    var v19 uint64 = (p0 + p0)
                    v20 := GenS2{F0: 134, F1: 123, F2: 255}
                    fmt.Println(p0)
                    _ = v19
                    _ = v20
                }
                if true {
                    if (-128 != 100) {
                        v21 := GenS2{F0: 65535, F1: 128, F2: 29}
                        p1 = p1
                        v22 := []int{}
                        v23 := GenS3{F0: 2}
                        var v24 int16 = -1
                        _ = v21
                        _ = v22
                        _ = v23
                        _ = v24
                    }
                } else {
                    v25 := []int{(fi11 & fi11)}
                    _ = v25
                }
                if (-711 <= 795) {
                    if (0 != 42) {
                        p2 = p2
                        v26 := []int{}
                        _ = v26
                    }
                    var v27 uint16 = 1
                    fmt.Println(p0)
                    if (p2 < (p2 ^ (^(p2)))) {
                        fmt.Println(v27)
                        var v28 int32 = -1
                        _ = v28
                    } else {
                        var v29 uint = 4294967295
                        _ = v29
                    }
                    _ = v27
                } else {
                    if true {
                        v30 := []int{fi10, fi11, (fi11 ^ (-(fi10))), 2147483647}
                        v31 := GenS4{F0: -456}
                        v32 := GenS2{F0: 255, F1: uint8(p0), F2: 2}
                        _ = v30
                        _ = v31
                        _ = v32
                    } else {
                        var v33 uint64 = (p0 * (p0 ^ 4294967295))
                        v33 += (p0 + (^(p0)))
                        _ = v33
                    }
                    p2 = p2
                    v34 := []int{fi11, (fi10 >> ((1) & 63)), -8}
                    _ = v34
                }
            } else {
                var v35 int64 = 9223372036854775807
                _ = v35
            }
            _ = fi11
        }
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return -1
}
func genFn1(p0 uint64, p1 uint8) uint {
    var v36 uint = 0
    for fi37 := 0; fi37 < 17; fi37++ {
        p1 = p1
        switch "hello" {
        case "foo bar":
            v38 := []int{}
            _ = v38
        default:
            v39 := []int{(fi37 | (fi37 / ((fi37) | 1)))}
            _ = v39
        }
        _ = fi37
    }
    if true {
        p0 -= 42
        p1 += (p1 | p1)
        for fi40 := 0; fi40 < 7; fi40++ {
            fmt.Println(fi40, fi40, v36)
            _ = fi40
        }
    }
    _ = p0
    _ = p1
    _ = v36
    return v36
}

func main() {
    var v1 float32 = -2.25
    var v2 uint64 = 136
    var v3 string = "foo bar"
    v4 := make(map[string]uint32)
    var v5 uint32 = 2
    v6 := make(map[string]uint32)
    var v41 uint16 = 2
    switch 1 {
    case 2147483647:
        v5 -= 0
    case -989:
        v41 = (^(v41))
    case 32767:
        switch -2147483648 {
        case 127:
            fmt.Println(v41)
        case 0:
            v42 := v6["\n\t"]
            _ = v42
        case 856:
            if (true || false) {
                switch v3 {
                case "a":
                    switch 32768 {
                    case 128:
                        var v43 float64 = -2.25
                        _ = v43
                    }
                case "hello":
                    v44 := GenS4{F0: 0}
                    _ = v44
                case "":
                    var v45 int32 = 2147483647
                    _ = v45
                default:
                    v46 := GenS5{F0: (true || true)}
                    _ = v46
                }
                v2 = (v2 | uint64(v5))
                switch "\n\t" {
                case "":
                    v47 := GenS6{F0: 1, F1: v2}
                    _ = v47
                }
            }
        }
    default:
        v2 = (^(v2))
    }
    var v48 uint16 = 688
    for fi49 := 0; fi49 < 6; fi49++ {
        switch -2147483648 {
        case 1:
            for fi50 := 0; fi50 < 1; fi50++ {
                v41 += 374
                var v51 float64 = -2.25
                var v52 int64 = 9223372036854775807
                _ = fi50
                _ = v51
                _ = v52
            }
        case 4294967295:
            v53 := []int{(fi49 - 42), fi49, 128, (fi49 - -647)}
            _ = v53
        case 127:
            v54 := []int{fi49, (fi49 << ((2) & 63))}
            _ = v54
        }
        if (0 >= 0) {
            var v55 uint8 = 127
            _ = len(v4)
            _ = v55
        } else {
            var v56 float32 = float32(((0.0 - ((0.1 + 1.5) - 0.0)) / (1.5 * 3.14159)))
            v1 += (float32(-2.25) - (float32(((float64(v1) / 1.0) / 0.1)) * v1))
            _ = v56
        }
        var v57 int8 = -36
        _ = fi49
        _ = v57
    }
    v1 = (1.5 - v1)
    v41 += (v48 >> ((uint(v2)) & 63))
    fmt.Println(v48)
    switch 42 {
    case 4294967295:
        if false {
            if true {
                if (!true) {
                    _ = len(v6)
                    v2 -= v2
                } else {
                    for fi58 := 0; fi58 < 14; fi58++ {
                        for fi59 := 0; fi59 < 9; fi59++ {
                            if (!(!false)) {
                                v60 := GenS7{F0: (v5 - (v5 >> ((genFn1((v2 * (v2 + (^(v2)))), 193)) & 63))), F1: (fi58 + fi59)}
                                v61 := GenS3{F0: 65535}
                                var v62 int32 = 2147483647
                                _ = v60
                                _ = v61
                                _ = v62
                            }
                            var v63 uint32 = 4294967295
                            _ = fi59
                            _ = v63
                        }
                        v48 = (v41 | (v48 | 65535))
                        v64 := GenS1{F0: -16}
                        _ = fi58
                        _ = v64
                    }
                    if false {
                        for fi65 := 0; fi65 < 9; fi65++ {
                            v2 = v2
                            var v66 string = (v3 + (v3 + "\n\t"))
                            _ = fi65
                            _ = v66
                        }
                        v2 = (^(v2))
                    } else {
                        var v67 uint32 = v5
                        v68 := GenS5{F0: true}
                        _ = v67
                        _ = v68
                    }
                }
            } else {
                for fi69 := 0; fi69 < 19; fi69++ {
                    if (-727 < -9223372036854775808) {
                        v70 := []int{fi69}
                        _ = v70
                    }
                    if (!(32768 > fi69)) {
                        var v71 string = v3
                        if (0 >= -966) {
                            v41 += uint16(v5)
                            for fi72 := 0; fi72 < 5; fi72++ {
                                var v73 int8 = 0
                                v48 = v48
                                v71 = v71
                                _ = fi72
                                _ = v73
                            }
                            var v74 float32 = v1
                            _ = v74
                        }
                        fmt.Println(v48, v3, v5)
                        _ = v71
                    } else {
                        switch v3 {
                        case "a":
                            if true {
                                v75 := GenS5{F0: (v1 != v1)}
                                v5 += 2
                                _ = v75
                            }
                        default:
                            v76 := GenS8{F0: 50, F1: -9223372036854775808, F2: 42}
                            _ = v76
                        }
                        v5 = v5
                    }
                    _ = fi69
                }
            }
        } else {
            v77 := []int{-575, 32768}
            fmt.Println(v48, v3, v2)
            _ = v77
        }
    case -570:
        var v78 uint = 18446744073709551615
        _ = v78
    case -386:
        switch v3 {
        case "\n\t":
            if false {
                switch -9223372036854775808 {
                case 0:
                    switch -692 {
                    case 4294967295:
                        if true {
                            var v79 float64 = float64(v1)
                            _ = v79
                        } else {
                            v2 = v2
                            var v80 uint8 = 127
                            _ = v80
                        }
                    case 127:
                        v81 := []int{2147483647, 127, 1, 1}
                        _ = v81
                    case 42:
                        v82 := GenS9{F0: 454, F1: 496, F2: 54}
                        _ = v82
                    default:
                        switch "foo bar" {
                        case "hello":
                            for fi83 := 0; fi83 < 17; fi83++ {
                                for fi84 := 0; fi84 < 6; fi84++ {
                                    var v85 string = "\n\t"
                                    _ = fi84
                                    _ = v85
                                }
                                if (v41 < (v41 ^ 42)) {
                                    var v86 string = v3
                                    v87 := v6[v86]
                                    v1 = v1
                                    _ = v86
                                    _ = v87
                                } else {
                                    v1 += 0.0
                                }
                                _ = fi83
                            }
                        }
                    }
                case -9223372036854775808:
                    for fi88 := 0; fi88 < 8; fi88++ {
                        v89 := v6[v3]
                        var v90 int32 = 42
                        _ = fi88
                        _ = v89
                        _ = v90
                    }
                case 128:
                    var v91 float32 = 0.0
                    _ = v91
                default:
                    if false {
                        if (v48 > 0) {
                            if true {
                                for fi92 := 0; fi92 < 17; fi92++ {
                                    switch 9223372036854775807 {
                                    case 811:
                                        v93 := []int{}
                                        _ = v93
                                    case 32768:
                                        v1 -= v1
                                    case 128:
                                        var v94 int16 = -1
                                        _ = v94
                                    }
                                    _ = fi92
                                }
                            }
                            v41 = (v41 & (^(v48)))
                        }
                        if (!false) {
                            for fi95 := 0; fi95 < 10; fi95++ {
                                var v96 uint8 = 2
                                for fi97 := 0; fi97 < 16; fi97++ {
                                    v98 := GenS10{F0: v48}
                                    v4["foo bar"] = v5
                                    v1 = float32(((-2.25 - 1e10) / (float64(0.1) + 0.1)))
                                    v99 := []int{(^(fi95)), fi97, (-(fi97)), (fi97 | (fi97 ^ (fi95 << ((2) & 63))))}
                                    fmt.Println(fi95, v3)
                                    _ = fi97
                                    _ = v98
                                    _ = v99
                                }
                                _ = fi95
                                _ = v96
                            }
                            for fi100 := 0; fi100 < 18; fi100++ {
                                v101 := GenS11{F0: 1, F1: 73}
                                for fi102 := 0; fi102 < 11; fi102++ {
                                    if false {
                                        switch -1 {
                                        case 1:
                                            fmt.Println(v3)
                                        case -1:
                                            if (false && ((((v1 / (1.5 + v1)) - ((((1e10 - v1) / 3.14159) - 0.1) * (v1 - -2.25))) / ((float32(1e10) / v1) - 3.14159)) >= v1)) {
                                                var v103 int32 = genFn0(191, v48, 411)
                                                v4[(v3 + v3)] = 1
                                                fmt.Println(v3)
                                                _ = v103
                                            } else {
                                                v104 := GenS3{F0: 1}
                                                _ = v104
                                            }
                                        default:
                                            switch fi102 {
                                            case 0:
                                                v105 := []int{(fi100 | (^(fi102)))}
                                                _ = v105
                                            default:
                                                switch (fi102 << ((208) & 63)) {
                                                case 4294967295:
                                                    v41 = v41
                                                case 1:
                                                    v106 := GenS12{F0: 1}
                                                    _ = v106
                                                }
                                                v107 := GenS13{F0: genFn1((v2 / (((v2 & v2)) | 1)), 128), F1: (v5 % ((v5) | 1)), F2: 1}
                                                _ = v107
                                            }
                                        }
                                        _ = len(v6)
                                        fmt.Println(v3, fi102, v5)
                                        v3 = ((v3 + v3) + "hello")
                                    } else {
                                        if ((!false) && (111 > -68)) {
                                            fmt.Println(v3, v3, v48)
                                            var v108 int64 = -1
                                            var v109 uint8 = 6
                                            _ = v108
                                            _ = v109
                                        } else {
                                            var v110 int = (fi100 & (fi102 >> ((0) & 63)))
                                            switch (v3 + (v3 + "")) {
                                            case "a":
                                                v111 := GenS14{F0: 0, F1: 255, F2: 710}
                                                _ = v111
                                            }
                                            v48 += v48
                                            _ = v110
                                        }
                                    }
                                    switch (v3 + v3) {
                                    case "":
                                        v41 -= (v41 % ((900) | 1))
                                    case "foo bar":
                                        v112 := []int{fi102}
                                        _ = v112
                                    default:
                                        var v113 uint16 = (v41 ^ (v41 >> ((529) & 63)))
                                        _ = v113
                                    }
                                    _ = fi102
                                }
                                _ = fi100
                                _ = v101
                            }
                            for fi114 := 0; fi114 < 14; fi114++ {
                                switch ((((v3 + "foo bar") + (v3 + v3)) + v3) + "foo bar") {
                                case "":
                                    fmt.Println(v1, v1, v41)
                                default:
                                    _ = len(v4)
                                }
                                _ = fi114
                            }
                        }
                    }
                }
                var v115 bool = (!true)
                var v116 uint = 211
                _ = v115
                _ = v116
            }
        case "":
            var v117 uint16 = (v48 >> ((702) & 63))
            _ = v117
        case "foo bar":
            for fi118 := 0; fi118 < 1; fi118++ {
                for fi119 := 0; fi119 < 14; fi119++ {
                    v1 = v1
                    v120 := []int{2147483647, fi118, fi119, fi119}
                    v121 := GenS3{F0: 255}
                    _ = fi119
                    _ = v120
                    _ = v121
                }
                for fi122 := 0; fi122 < 15; fi122++ {
                    if (!(((!(490 != (v48 ^ v41))) || ((((v3 + "hello") > (("a" + "foo bar") + ((v3 + v3) + "a"))) && true) && true)) && (-9223372036854775808 < -9223372036854775808))) {
                        if true {
                            v123 := []int{(fi122 >> ((395) & 63)), (^(fi122)), fi118, 1}
                            if len(v123) >= 2 {
                                v123 = v123[1:len(v123)]
                            }
                            var v124 float32 = 0.1
                            v1 = v124
                            _ = v123
                            _ = v124
                        }
                    } else {
                        var v125 uint16 = v41
                        var v126 uint16 = (v41 << ((1) & 63))
                        v5 -= (v5 + (v5 + v5))
                        _ = v125
                        _ = v126
                    }
                    var v127 float32 = (1e10 * v1)
                    _ = fi122
                    _ = v127
                }
                _ = fi118
            }
        }
    default:
        _ = len(v4)
    }
    v128 := []int{32767, -539, 215, -450}
    for fi129 := 0; fi129 < 1; fi129++ {
        v130 := GenS14{F0: 2, F1: 0, F2: 436}
        if true {
            delete(v4, v3)
        } else {
            v131 := GenS15{F0: -1, F1: -991, F2: 2147483647}
            v132 := GenS3{F0: genFn1((^(v2)), 0)}
            switch 128 {
            case 66:
                var v133 uint = genFn1(1, uint8(v5))
                _ = v133
            case -450:
                if false {
                    if (!((2147483647 <= 0) || false)) {
                        switch "\n\t" {
                        case "foo bar":
                            v128 = append(v128, fi129)
                        }
                    } else {
                        v134 := v6[v3]
                        _ = v134
                    }
                }
            default:
                v135 := GenS16{F0: (v2 / ((0) | 1)), F1: -1}
                _ = v135
            }
            _ = v131
            _ = v132
        }
        if false {
            fmt.Println(v3, v48)
            if len(v128) > 0 {
                v128[0] = -9223372036854775808
            }
        }
        _ = fi129
        _ = v130
    }
    var v136 float64 = 1.5
    for fi137 := 0; fi137 < 2; fi137++ {
        v138 := GenS17{F0: (v48 % (((v48 & 32767)) | 1))}
        if (42 <= 9223372036854775807) {
            switch (v3 + "") {
            case "":
                fmt.Println(v136, v5)
            case "\n\t":
                var v139 string = ("hello" + v3)
                _ = v139
            case "foo bar":
                var v140 bool = (!false)
                _ = v140
            }
            var v141 int8 = 100
            _ = v141
        } else {
            switch "\n\t" {
            case "hello":
                if false {
                    var v142 string = v3
                    var v143 float64 = float64(-2.25)
                    fmt.Println(v2)
                    _ = v142
                    _ = v143
                } else {
                    fmt.Println(fi137)
                }
            case "\n\t":
                fmt.Println(v41)
            default:
                var v144 bool = false
                _ = v144
            }
            for fi145 := 0; fi145 < 7; fi145++ {
                var v146 string = (v3 + "hello")
                switch v146 {
                case "":
                    fmt.Println(v2)
                    if (0.0 <= 1.5) {
                        var v147 bool = (1 > (v41 & (v48 - v48)))
                        fmt.Println(fi137, v41, fi145)
                        var v148 int32 = 520
                        _ = v147
                        _ = v148
                    }
                case "hello":
                    v149 := v6[v146]
                    var v150 uint32 = v149
                    if true {
                        v136 -= (-2.25 - v136)
                        var v151 string = v146
                        _ = v151
                    } else {
                        if (0 == v41) {
                            var v152 uint32 = v150
                            v146 = (v146 + ("hello" + ""))
                            _ = v152
                        } else {
                            if false {
                                var v153 uint16 = 253
                                v154 := GenS8{F0: 100, F1: 42, F2: uint8(v2)}
                                _ = v153
                                _ = v154
                            }
                            switch (fi145 ^ fi137) {
                            case 32768:
                                if false {
                                    switch v3 {
                                    case "\n\t":
                                        v150 = 92
                                    case "":
                                        if true {
                                            v128 = append(v128, 564)
                                            v41 -= v41
                                            if false {
                                                v155 := GenS8{F0: 100, F1: 1, F2: 128}
                                                var v156 int32 = 2147483647
                                                _ = v155
                                                _ = v156
                                            } else {
                                                var v157 int64 = 1
                                                _ = v157
                                            }
                                        }
                                    }
                                } else {
                                    fmt.Println(v3, fi137, v3)
                                }
                                var v158 int16 = -32768
                                var v159 int8 = 42
                                _ = v158
                                _ = v159
                            case -339:
                                switch (fi145 << ((2) & 63)) {
                                case -9223372036854775808:
                                    if true {
                                        v2 = v2
                                    } else {
                                        v150 = (v150 & 2)
                                        var v160 uint8 = 42
                                        _ = v160
                                    }
                                case 9223372036854775807:
                                    var v161 int = -237
                                    if true {
                                        var v162 uint8 = 0
                                        switch -845 {
                                        case -494:
                                            v150 = v149
                                        case 9223372036854775807:
                                            v146 = v3
                                        case 127:
                                            if false {
                                                var v163 bool = false
                                                _ = v163
                                            }
                                        }
                                        if false {
                                            var v164 uint16 = v41
                                            switch "foo bar" {
                                            case "\n\t":
                                                var v165 uint64 = (v2 - (^(v2)))
                                                _ = v165
                                            case "hello":
                                                var v166 float32 = (-2.25 * (v1 + v1))
                                                _ = v166
                                            default:
                                                v167 := GenS18{F0: 42}
                                                _ = v167
                                            }
                                            fmt.Println(v164, fi145)
                                            if len(v128) >= 2 {
                                                v128 = v128[1:len(v128)]
                                            }
                                            _ = v164
                                        }
                                        var v168 float32 = float32(v136)
                                        _ = v162
                                        _ = v168
                                    } else {
                                        if (!false) {
                                            if true {
                                                if ((v48 << ((771) & 63)) > 946) {
                                                    v169 := GenS7{F0: (v5 % ((v150) | 1)), F1: (v161 + (fi137 / ((fi137) | 1)))}
                                                    if (838 >= (v5 % ((2) | 1))) {
                                                        var v170 int = 128
                                                        fmt.Println(fi145, v5, v5)
                                                        _ = v170
                                                    } else {
                                                        var v171 int16 = 32767
                                                        fmt.Println(v146, v161)
                                                        if ((float64(v1) / -2.25) >= v136) {
                                                            fmt.Println(fi137, fi145, v171)
                                                            var v172 uint32 = (v149 & (v150 + 839))
                                                            _ = v172
                                                        } else {
                                                            if len(v128) >= 2 {
                                                                v128 = v128[1:len(v128)]
                                                            }
                                                            switch 4294967295 {
                                                            case -2147483648:
                                                                switch v3 {
                                                                case "a":
                                                                    switch fi137 {
                                                                    case 4294967295:
                                                                        fmt.Println(v136, v150)
                                                                        v173 := GenS19{F0: v48, F1: 567}
                                                                        _ = v173
                                                                    }
                                                                case "\n\t":
                                                                    v1 = 0.0
                                                                    if false {
                                                                        var v174 float64 = (v136 * ((v136 / v136) / v136))
                                                                        if false {
                                                                            var v175 uint16 = v48
                                                                            if len(v128) >= 2 {
                                                                                v128 = v128[1:len(v128)]
                                                                            }
                                                                            _ = v175
                                                                        } else {
                                                                            switch "\n\t" {
                                                                            case "foo bar":
                                                                                var v176 float64 = ((v136 - (0.1 - v136)) * (((v136 / (v136 + ((v136 * v174) * 3.14159))) - v136) / 3.14159))
                                                                                _ = v176
                                                                            case "":
                                                                                v41 = v48
                                                                            case "hello":
                                                                                if true {
                                                                                    if false {
                                                                                        v4["hello"] = 937
                                                                                    }
                                                                                    var v177 string = v146
                                                                                    v48 = v48
                                                                                    _ = v177
                                                                                } else {
                                                                                    if ((fi137 << ((764) & 63)) >= fi145) {
                                                                                        v178 := GenS20{F0: v41, F1: 2}
                                                                                        v149 -= 4294967295
                                                                                        _ = v178
                                                                                    } else {
                                                                                        if (!(42 <= 154)) {
                                                                                            fmt.Println(v150, v150, v146)
                                                                                            var v179 uint32 = (v149 / ((2) | 1))
                                                                                            _ = v179
                                                                                        }
                                                                                        var v180 bool = ((v149 & v149) != 0)
                                                                                        _ = v180
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                        var v181 uint16 = (v48 << ((255) & 63))
                                                                        _ = v174
                                                                        _ = v181
                                                                    } else {
                                                                        fmt.Println(v3)
                                                                    }
                                                                case "hello":
                                                                    _ = len(v4)
                                                                default:
                                                                    v136 += (1e10 + float64(v1))
                                                                }
                                                            case 32768:
                                                                var v182 string = v146
                                                                _ = v182
                                                            default:
                                                                var v183 uint16 = v41
                                                                _ = v183
                                                            }
                                                        }
                                                        _ = v171
                                                    }
                                                    v128 = append(v128, fi145)
                                                    _ = v169
                                                }
                                                var v184 uint8 = 0
                                                _ = v184
                                            } else {
                                                v146 = ("a" + (((("foo bar" + "") + "\n\t") + "foo bar") + ((v3 + "hello") + (("foo bar" + (v146 + ((("hello" + (((v3 + (v3 + (v146 + (((("" + v3) + v3) + v3) + v3)))) + v146) + v3)) + (("foo bar" + "a") + ("foo bar" + (("\n\t" + "a") + "a")))) + v3))) + ""))))
                                                if (9223372036854775807 > -9223372036854775808) {
                                                    fmt.Println(v41, v48, fi145)
                                                    if false {
                                                        var v185 float64 = v136
                                                        var v186 uint8 = 1
                                                        _ = v185
                                                        _ = v186
                                                    } else {
                                                        v187 := GenS21{F0: (v146 + "a"), F1: 42}
                                                        _ = v187
                                                    }
                                                    fmt.Println(v161, v5)
                                                }
                                                var v188 uint8 = 2
                                                _ = v188
                                            }
                                            var v189 uint16 = 65535
                                            v190 := GenS20{F0: (^(v48)), F1: 1}
                                            fmt.Println(v150)
                                            v189 = (v48 * 1)
                                            _ = v189
                                            _ = v190
                                        } else {
                                            v149 += v150
                                            v41 += (v41 << ((255) & 63))
                                            var v191 int32 = -2147483648
                                            _ = v191
                                        }
                                        v192 := v6[v3]
                                        if (uint32(v2) == 42) {
                                            var v193 float32 = v1
                                            fmt.Println(v2)
                                            v194 := v6[v3]
                                            _ = v193
                                            _ = v194
                                        } else {
                                            v195 := GenS22{F0: -370}
                                            var v196 int = fi145
                                            if (v136 > ((v136 * float64(0.1)) / v136)) {
                                                v150 = uint32(v2)
                                            }
                                            _ = v195
                                            _ = v196
                                        }
                                        _ = v192
                                    }
                                    _ = v161
                                }
                            }
                            var v197 int16 = 1
                            _ = v197
                        }
                    }
                    _ = v149
                    _ = v150
                case "foo bar":
                    var v198 float32 = 3.14159
                    _ = v198
                default:
                    var v199 uint32 = 606
                    if (!true) {
                        var v200 int8 = 0
                        v199 = (v5 << ((405) & 63))
                        v201 := GenS16{F0: 270, F1: -9223372036854775808}
                        fmt.Println(v3)
                        if len(v128) >= 2 {
                            v128 = v128[1:len(v128)]
                        }
                        fmt.Println(fi137, v2)
                        var v202 bool = (-619 >= 0)
                        if false {
                            v200 -= v200
                        }
                        var v203 int8 = v200
                        _ = v200
                        _ = v201
                        _ = v202
                        _ = v203
                    } else {
                        if len(v128) >= 2 {
                            v128 = v128[1:len(v128)]
                        }
                        if true {
                            v146 = v146
                            fmt.Println(v146)
                        }
                        if false {
                            if ((v1 <= 0.0) && (-1 > 1)) {
                                _ = len(v4)
                                var v204 uint8 = 108
                                if len(v128) > 0 {
                                    v128[0] = (fi137 ^ int(v199))
                                }
                                _ = v204
                            } else {
                                v48 -= (^(v41))
                                switch v3 {
                                case "hello":
                                    v205 := GenS21{F0: v3, F1: int16(v199)}
                                    _ = v205
                                case "foo bar":
                                    var v206 int = (fi145 + fi137)
                                    _ = v206
                                case "a":
                                    if false {
                                        var v207 string = ""
                                        _ = v207
                                    }
                                }
                                _ = len(v128)
                            }
                        }
                        v208 := GenS2{F0: 2, F1: 42, F2: 255}
                        var v209 int32 = -851
                        _ = v208
                        _ = v209
                    }
                    v136 -= 0.1
                    var v210 float32 = ((v1 + v1) - float32((1e10 / (1e10 + v136))))
                    _ = v199
                    _ = v210
                }
                var v211 string = "\n\t"
                _ = fi145
                _ = v146
                _ = v211
            }
            v48 = (v41 + uint16(v5))
        }
        v41 = (v41 >> ((255) & 63))
        _ = fi137
        _ = v138
    }
    v1 -= (0.1 - (v1 / (v1 + v1)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v128:", len(v128), v128)
    fmt.Println("v136:", v136)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v128
    _ = v136
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v48
    _ = v5
    _ = v6
}
