package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 uint64; F2 uint8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int64; F1 uint8 }
type GenS5 struct { F0 bool; F1 uint32; F2 string }
type GenS6 struct { F0 uint; F1 int8 }
type GenS7 struct { F0 uint16; F1 int64; F2 uint }
type GenS8 struct { F0 bool; F1 int; F2 uint8 }
type GenS9 struct { F0 int64; F1 uint8 }
type GenS10 struct { F0 uint32; F1 uint16 }
type GenS11 struct { F0 uint; F1 int8; F2 bool }
type GenS12 struct { F0 string }
type GenS13 struct { F0 uint32; F1 bool; F2 int64 }
type GenS14 struct { F0 uint32; F1 uint32; F2 int8 }
type GenS15 struct { F0 string }
type GenS16 struct { F0 int32 }
type GenS17 struct { F0 string; F1 uint16 }
type GenS18 struct { F0 uint32; F1 uint; F2 uint16 }
type GenS19 struct { F0 int; F1 int8; F2 string }
type GenS20 struct { F0 int32; F1 uint; F2 int16 }
type GenS21 struct { F0 uint32 }
type GenS22 struct { F0 int16; F1 uint16; F2 uint64 }
func genFn0(p0 uint) uint8 {
    fmt.Println(p0)
    var v7 float32 = (1e10 + 3.14159)
    fmt.Println(v7, v7)
    for fi8 := 0; fi8 < 20; fi8++ {
        v9 := GenS1{F0: -1}
        var v10 uint16 = 65535
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = v7
    return 74
}
func genFn1(p0 uint8, p1 int64, p2 uint16) uint64 {
    switch 0 {
    case -2147483648:
        v11 := GenS2{F0: 67, F1: 0, F2: (p0 - p0)}
        _ = v11
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        if (!(42 > 1)) {
            var v13 float32 = 0.0
            if true {
                var v14 uint16 = 255
                fmt.Println(p1, p2, p0)
                var v15 string = ("hello" + "\n\t")
                _ = v14
                _ = v15
            }
            _ = v13
        } else {
            if ("hello" < (("a" + "hello") + "\n\t")) {
                v16 := GenS1{F0: (p1 % (((p1 + p1)) | 1))}
                v17 := GenS3{F0: 42}
                _ = v16
                _ = v17
            }
        }
        _ = fi12
    }
    p2 = uint16(p0)
    p0 = (^(p0))
    switch -52 {
    case -902:
        for fi18 := 0; fi18 < 2; fi18++ {
            fmt.Println(p0, p0, p1)
            for fi19 := 0; fi19 < 5; fi19++ {
                v20 := []int{}
                _ = fi19
                _ = v20
            }
            p1 = (p1 * (p1 - (p1 + p1)))
            _ = fi18
        }
    case 304:
        switch "\n\t" {
        case "a":
            if (728 != 825) {
                var v21 float64 = (0.1 / 1.5)
                _ = v21
            }
        default:
            if (true && false) {
                for fi22 := 0; fi22 < 13; fi22++ {
                    if false {
                        if (!((1e10 + (float64((0.0 - (0.1 * float32((0.1 + (3.14159 * 0.1)))))) * (-2.25 / 3.14159))) != (0.0 / ((3.14159 - -2.25) / 0.1)))) {
                            for fi23 := 0; fi23 < 1; fi23++ {
                                if (533 > 2147483647) {
                                    p0 -= (p0 ^ (^(p0)))
                                } else {
                                    p0 = p0
                                    var v24 int8 = 0
                                    _ = v24
                                }
                                if (0.1 == 0.0) {
                                    if (!((((0 <= uint(p2)) || true) || (!(!false))) && false)) {
                                        switch "hello" {
                                        case "a":
                                            p1 = int64(p2)
                                            p0 += p0
                                            p1 = (p1 + -9223372036854775808)
                                        case "":
                                            v25 := GenS1{F0: 42}
                                            _ = v25
                                        }
                                        p1 -= p1
                                        v26 := GenS1{F0: 53}
                                        _ = v26
                                    }
                                    var v27 string = "foo bar"
                                    v28 := GenS3{F0: 146}
                                    _ = v27
                                    _ = v28
                                } else {
                                    fmt.Println(fi23)
                                }
                                fmt.Println(fi23, p1)
                                _ = fi23
                            }
                            for fi29 := 0; fi29 < 4; fi29++ {
                                var v30 uint8 = (p0 ^ (p0 % ((42) | 1)))
                                var v31 float64 = -2.25
                                v32 := []int{fi22, fi22, fi29, fi22}
                                _ = fi29
                                _ = v30
                                _ = v31
                                _ = v32
                            }
                        }
                        fmt.Println(p1, p1)
                        p1 -= (p1 << ((2) & 63))
                    }
                    _ = fi22
                }
                v33 := []int{}
                v33 = append(v33, 9223372036854775807)
                _ = v33
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 440
}

func main() {
    var v1 int = 42
    var v2 string = "a"
    var v3 bool = false
    var v4 uint32 = 65535
    var v5 int = (v1 + v1)
    v6 := make(map[uint32]bool)
    var v34 uint64 = 42
    switch (v1 & v1) {
    case 9223372036854775807:
        fmt.Println(v5, v1, v1)
    case -257:
        for fi35 := 0; fi35 < 2; fi35++ {
            v34 = (^(v34))
            var v36 int16 = 0
            _ = fi35
            _ = v36
        }
    default:
        v6[v4] = false
    }
    if ((v5 ^ (v1 / ((v1) | 1))) >= (v1 * -201)) {
        v37 := v6[15]
        switch (v2 + "") {
        case "a":
            for fi38 := 0; fi38 < 13; fi38++ {
                var v39 int8 = 52
                var v40 float64 = ((-2.25 / 3.14159) - 1e10)
                _ = fi38
                _ = v39
                _ = v40
            }
        case "":
            var v41 int8 = -1
            _ = v41
        case "foo bar":
            var v42 uint8 = 248
            _ = v42
        default:
            switch v2 {
            case "hello":
                v34 = (^(v34))
            }
        }
        switch int(v34) {
        case -389:
            if v37 {
                var v43 uint8 = 1
                _ = v43
            }
        default:
            v44 := []int{(v1 << ((2) & 63)), -9223372036854775808, 128}
            _ = v44
        }
        _ = v37
    }
    for fi45 := 0; fi45 < 1; fi45++ {
        for fi46 := 0; fi46 < 5; fi46++ {
            v1 = (v5 & v1)
            var v47 int64 = -1
            v4 = (v4 | (v4 % ((1) | 1)))
            switch v2 {
            case "":
                if v3 {
                    v3 = true
                    switch v2 {
                    case "a":
                        if (v1 <= fi46) {
                            fmt.Println(v3, fi46, v5)
                            var v48 float32 = ((((0.0 * (-2.25 / float32(((0.1 - 0.0) + 0.1)))) / 1.0) - 0.0) / 3.14159)
                            fmt.Println(fi46, v1, v2)
                            v49 := GenS2{F0: 255, F1: 865, F2: genFn0(673)}
                            _ = v48
                            _ = v49
                        } else {
                            var v50 string = v2
                            v1 = 42
                            _ = v50
                        }
                    case "foo bar":
                        v51 := v6[(v4 / ((v4) | 1))]
                        fmt.Println(v34, v5)
                        v52 := []int{v5, 1, (v1 | (v5 << ((746) & 63)))}
                        _ = v51
                        _ = v52
                    default:
                        if (0 <= 1) {
                            var v53 uint64 = v34
                            _ = v53
                        } else {
                            v2 = v2
                            _ = len(v6)
                            v54 := GenS2{F0: 42, F1: v34, F2: 127}
                            _ = v54
                        }
                    }
                    var v55 float64 = 1.5
                    var v56 uint64 = (^(v34))
                    _ = v55
                    _ = v56
                } else {
                    switch (fi45 + fi45) {
                    case 9223372036854775807:
                        var v57 int = 1
                        _ = v57
                    case 895:
                        var v58 uint64 = (^(v34))
                        _ = v58
                    case -981:
                        _ = len(v6)
                    }
                    if true {
                        var v59 int8 = -41
                        var v60 uint64 = (^(v34))
                        v61 := GenS1{F0: (v47 + (v47 << ((978) & 63)))}
                        _ = v59
                        _ = v60
                        _ = v61
                    } else {
                        var v62 int64 = -684
                        _ = v62
                    }
                    var v63 int8 = 1
                    _ = v63
                }
            default:
                v64 := v6[1]
                _ = v64
            }
            v65 := GenS2{F0: 2, F1: v34, F2: 128}
            v1 = 9223372036854775807
            switch v2 {
            case "hello":
                var v66 uint8 = genFn0(65535)
                _ = v66
            }
            _ = fi46
            _ = v47
            _ = v65
        }
        if v3 {
            v67 := GenS2{F0: 4294967295, F1: 18446744073709551615, F2: 255}
            v68 := GenS4{F0: 1, F1: 25}
            var v69 string = v2
            _ = v67
            _ = v68
            _ = v69
        } else {
            v1 -= fi45
            v70 := GenS3{F0: (v4 | v4)}
            _ = v70
        }
        _ = fi45
    }
    if (42 != -40) {
        v2 = ("hello" + "foo bar")
        for fi71 := 0; fi71 < 12; fi71++ {
            if false {
                v6[(v4 + (v4 * v4))] = (v3 && (!true))
                v6[(v4 ^ uint32(v34))] = v3
                var v72 uint = uint(v34)
                _ = v72
            }
            if ((((((v2 + v2) + ((v2 + v2) + v2)) + v2) < v2) && (42 < 0)) && v3) {
                fmt.Println(v34, v1, v1)
                var v73 float32 = (((1e10 - 0.0) / 3.14159) + (0.0 / -2.25))
                _ = v73
            } else {
                v74 := GenS2{F0: 712, F1: v34, F2: 127}
                _ = v74
            }
            _ = fi71
        }
        if v3 {
            v5 += v1
            v34 += 141
            if (v5 == v1) {
                v6[(v4 << ((2) & 63))] = v3
            }
        } else {
            for fi75 := 0; fi75 < 17; fi75++ {
                for fi76 := 0; fi76 < 17; fi76++ {
                    delete(v6, (^(v4)))
                    fmt.Println(v3, v34, v34)
                    var v77 int32 = 2147483647
                    _ = fi76
                    _ = v77
                }
                var v78 string = ("foo bar" + v2)
                var v79 string = v78
                _ = fi75
                _ = v78
                _ = v79
            }
            if (!(!v3)) {
                switch (v2 + v2) {
                case "foo bar":
                    v80 := GenS5{F0: (!v3), F1: (^(v4)), F2: v2}
                    _ = v80
                default:
                    switch v2 {
                    case "hello":
                        v5 = v5
                    default:
                        for fi81 := 0; fi81 < 19; fi81++ {
                            var v82 uint64 = (^(v34))
                            _ = fi81
                            _ = v82
                        }
                    }
                }
            } else {
                v4 = (v4 << ((0) & 63))
                if v3 {
                    v83 := GenS3{F0: (v4 % ((65535) | 1))}
                    var v84 uint16 = 42
                    _ = v83
                    _ = v84
                } else {
                    for fi85 := 0; fi85 < 20; fi85++ {
                        v86 := []int{32767, v5, 2147483647}
                        switch v2 {
                        case "a":
                            fmt.Println(v4, v5)
                        default:
                            _ = len(v86)
                        }
                        _ = fi85
                        _ = v86
                    }
                }
                if (!(1000 != 812)) {
                    switch (^(v5)) {
                    case 735:
                        var v87 float64 = float64((3.14159 * 0.1))
                        _ = v87
                    case 2147483647:
                        if (-2147483648 == 863) {
                            for fi88 := 0; fi88 < 5; fi88++ {
                                for fi89 := 0; fi89 < 5; fi89++ {
                                    var v90 uint16 = 760
                                    if (!v3) {
                                        var v91 uint = 308
                                        _ = v91
                                    }
                                    v92 := GenS6{F0: 836, F1: -1}
                                    _ = fi89
                                    _ = v90
                                    _ = v92
                                }
                                _ = fi88
                            }
                            switch (v5 | v1) {
                            case -1:
                                switch (v5 >> ((18446744073709551615) & 63)) {
                                case 32767:
                                    if v3 {
                                        for fi93 := 0; fi93 < 19; fi93++ {
                                            v94 := []int{}
                                            var v95 bool = v3
                                            switch (v2 + v2) {
                                            case "\n\t":
                                                fmt.Println(v4, v95, v2)
                                            case "hello":
                                                var v96 int = fi93
                                                _ = v96
                                            case "":
                                                v2 = "hello"
                                            default:
                                                var v97 uint8 = 251
                                                _ = v97
                                            }
                                            _ = fi93
                                            _ = v94
                                            _ = v95
                                        }
                                        v34 = v34
                                        v98 := []int{(v1 ^ v1)}
                                        _ = v98
                                    }
                                }
                            default:
                                if (true || false) {
                                    v99 := GenS5{F0: (0.0 <= float64(1e10)), F1: v4, F2: v2}
                                    var v100 int = 387
                                    _ = v100
                                    _ = v99
                                } else {
                                    var v101 string = "foo bar"
                                    v102 := GenS6{F0: 4294967295, F1: -128}
                                    _ = v101
                                    _ = v102
                                }
                            }
                            v103 := []int{(v5 >> ((172) & 63)), v5, (v1 - v1), (v5 * 165)}
                            _ = v103
                        } else {
                            v104 := []int{(v5 - (v1 - 9223372036854775807))}
                            if len(v104) > 0 {
                                v104[0] = v1
                            }
                            var v105 uint32 = 65535
                            _ = v104
                            _ = v105
                        }
                    }
                } else {
                    fmt.Println(v1)
                    v2 = (((v2 + (v2 + "\n\t")) + v2) + "")
                    v34 = (^(v34))
                }
            }
        }
    } else {
        v4 = 42
        for fi106 := 0; fi106 < 15; fi106++ {
            if v3 {
                fmt.Println(fi106, v3)
            } else {
                fmt.Println(v3)
                for fi107 := 0; fi107 < 16; fi107++ {
                    var v108 int64 = 42
                    _ = fi107
                    _ = v108
                }
            }
            for fi109 := 0; fi109 < 6; fi109++ {
                var v110 int16 = 1
                v3 = v3
                v111 := GenS5{F0: v3, F1: 2147483647, F2: v2}
                _ = fi109
                _ = v110
                _ = v111
            }
            var v112 string = ("hello" + "foo bar")
            _ = fi106
            _ = v112
        }
    }
    for fi113 := 0; fi113 < 12; fi113++ {
        v5 = v1
        _ = fi113
    }
    for fi114 := 0; fi114 < 5; fi114++ {
        v4 -= 565
        _ = fi114
    }
    v1 = v5
    v115 := GenS3{F0: (v4 << ((785) & 63))}
    v116 := GenS3{F0: v4}
    v3 = v3
    switch v2 {
    case "":
        v4 = v4
    }
    var v117 int8 = 100
    switch v2 {
    case "\n\t":
        if v3 {
            if ((v5 / (((v5 ^ (v1 >> ((484) & 63)))) | 1)) <= v5) {
                if v3 {
                    var v118 string = ""
                    _ = v118
                }
                switch ("a" + "hello") {
                case "a":
                    switch v1 {
                    case 4294967295:
                        switch v2 {
                        case "\n\t":
                            v1 -= -2147483648
                        case "hello":
                            v117 = (v117 << ((1) & 63))
                        default:
                            for fi119 := 0; fi119 < 10; fi119++ {
                                for fi120 := 0; fi120 < 15; fi120++ {
                                    if ((0.1 * (3.14159 / float32(0.1))) <= -2.25) {
                                        v121 := GenS7{F0: 405, F1: -1, F2: 255}
                                        switch v1 {
                                        case -175:
                                            v122 := GenS8{F0: (-9223372036854775808 != 9223372036854775807), F1: 42, F2: 143}
                                            _ = v122
                                        case -995:
                                            switch v2 {
                                            case "hello":
                                                switch fi119 {
                                                case -923:
                                                    var v123 int64 = -584
                                                    _ = v123
                                                }
                                            }
                                        default:
                                            v124 := GenS5{F0: v3, F1: (v4 | (^(v4))), F2: "a"}
                                            _ = v124
                                        }
                                        if v3 {
                                            var v125 uint = 401
                                            var v126 uint32 = 2147483647
                                            switch v5 {
                                            case 32768:
                                                if v3 {
                                                    switch int(v117) {
                                                    case -747:
                                                        var v127 uint = v125
                                                        _ = v127
                                                    default:
                                                        var v128 float32 = -2.25
                                                        _ = v128
                                                    }
                                                    _ = len(v6)
                                                    var v129 string = (v2 + v2)
                                                    var v130 float32 = (1e10 * 1.5)
                                                    _ = v129
                                                    _ = v130
                                                }
                                            case -2147483648:
                                                v6[(v126 * (v4 / ((2147483648) | 1)))] = v3
                                            }
                                            _ = v125
                                            _ = v126
                                        }
                                        v34 += (v34 << ((18446744073709551615) & 63))
                                        var v131 int32 = -1
                                        var v132 float64 = 0.0
                                        if true {
                                            fmt.Println(v4, v117, v5)
                                            v133 := []int{fi120, 654}
                                            v134 := GenS5{F0: v3, F1: (^(v4)), F2: v2}
                                            _ = v133
                                            _ = v134
                                        } else {
                                            v135 := GenS3{F0: (^(v4))}
                                            _ = v135
                                        }
                                        _ = v121
                                        _ = v131
                                        _ = v132
                                    } else {
                                        var v136 string = v2
                                        if true {
                                            var v137 bool = false
                                            _ = v137
                                        } else {
                                            if true {
                                                v34 = 42
                                                if v3 {
                                                    v138 := GenS9{F0: 1, F1: 255}
                                                    var v139 int = -2147483648
                                                    v4 = v4
                                                    var v140 uint8 = 2
                                                    _ = v138
                                                    _ = v139
                                                    _ = v140
                                                } else {
                                                    switch v136 {
                                                    case "foo bar":
                                                        v141 := GenS5{F0: true, F1: (v4 >> ((2) & 63)), F2: "a"}
                                                        _ = v141
                                                    }
                                                    v136 = ((("foo bar" + v136) + v2) + ("\n\t" + v2))
                                                }
                                                if false {
                                                    switch v5 {
                                                    case -9223372036854775808:
                                                        v34 = (v34 % ((uint64(v4)) | 1))
                                                    case 2147483647:
                                                        if (float64(float32((float64(float32(0.0)) * float64(0.0)))) > (((-2.25 * 0.0) / (((1e10 + float64(0.1)) + 3.14159) - 0.1)) + 0.1)) {
                                                            _ = len(v6)
                                                            fmt.Println(v1, v2)
                                                            var v142 string = ""
                                                            var v143 int64 = 0
                                                            if v3 {
                                                                v142 = v142
                                                                var v144 int8 = v117
                                                                if v3 {
                                                                    fmt.Println(v34)
                                                                    if (v3 && (0 != 127)) {
                                                                        var v145 float32 = 1e10
                                                                        v136 = (v136 + ("foo bar" + ((v136 + (v142 + v136)) + v2)))
                                                                        _ = v145
                                                                    }
                                                                }
                                                                _ = v144
                                                            }
                                                            _ = v142
                                                            _ = v143
                                                        }
                                                    default:
                                                        v4 += v4
                                                    }
                                                    v117 -= -42
                                                }
                                                v136 = (v136 + (v2 + ("hello" + v136)))
                                                v2 = v136
                                                fmt.Println(fi120, v34, fi119)
                                                if ((v34 ^ v34) == (v34 % ((v34) | 1))) {
                                                    v146 := v6[428]
                                                    v34 = v34
                                                    _ = v146
                                                } else {
                                                    v147 := []int{1, fi119, 962}
                                                    var v148 float32 = 1e10
                                                    _ = v147
                                                    _ = v148
                                                }
                                                if v3 {
                                                    v136 = "a"
                                                } else {
                                                    v4 -= (v4 * v4)
                                                }
                                            }
                                        }
                                        v149 := v6[v4]
                                        _ = v136
                                        _ = v149
                                    }
                                    v150 := []int{(v5 & v5), (-(fi120)), (-(v5)), fi119}
                                    if v3 {
                                        v151 := GenS10{F0: 42, F1: 881}
                                        v34 = (v34 / ((genFn1(255, 661, 0)) | 1))
                                        _ = v151
                                    } else {
                                        var v152 uint = 210
                                        var v153 int32 = 42
                                        var v154 string = v2
                                        _ = v152
                                        _ = v153
                                        _ = v154
                                    }
                                    _ = fi120
                                    _ = v150
                                }
                                fmt.Println(v117)
                                _ = fi119
                            }
                        }
                    }
                default:
                    v3 = false
                }
            }
            v155 := GenS2{F0: 312, F1: v34, F2: 42}
            var v156 int8 = (v117 >> ((1) & 63))
            _ = v155
            _ = v156
        } else {
            switch v2 {
            case "a":
                for fi157 := 0; fi157 < 14; fi157++ {
                    switch ("hello" + v2) {
                    case "\n\t":
                        for fi158 := 0; fi158 < 12; fi158++ {
                            fmt.Println(v34, fi158, v117)
                            v2 = (("" + (v2 + "hello")) + v2)
                            v159 := []int{fi158, 127, (^(v5)), (^(v5))}
                            _ = fi158
                            _ = v159
                        }
                    case "a":
                        if ("a" == ("foo bar" + v2)) {
                            for fi160 := 0; fi160 < 18; fi160++ {
                                v5 = (fi157 | v5)
                                if ((v117 - int8(v34)) >= v117) {
                                    v161 := []int{(fi157 - (v1 | (-(fi157)))), (-(fi160))}
                                    v162 := GenS11{F0: 296, F1: v117, F2: v3}
                                    if false {
                                        if (((1.5 * (0.1 + 1.5)) + (((1e10 + 0.1) * 0.1) * float32(3.14159))) <= 3.14159) {
                                            _ = len(v6)
                                            v3 = v3
                                            fmt.Println(fi160, v1)
                                            v163 := GenS2{F0: 1, F1: (v34 | (v34 / ((v34) | 1))), F2: 128}
                                            _ = v163
                                        }
                                    } else {
                                        v34 = v34
                                    }
                                    _ = v161
                                    _ = v162
                                }
                                if true {
                                    v164 := GenS12{F0: v2}
                                    v6[(v4 % ((v4) | 1))] = false
                                    _ = v164
                                } else {
                                    var v165 uint8 = 173
                                    fmt.Println(v34, v5)
                                    _ = v165
                                }
                                _ = fi160
                            }
                            v166 := GenS13{F0: 2, F1: (!(-2147483648 >= 143)), F2: 42}
                            if (false && ((fi157 == (fi157 + v5)) || ((v3 || true) && false))) {
                                if v3 {
                                    for fi167 := 0; fi167 < 1; fi167++ {
                                        fmt.Println(v117)
                                        v168 := GenS4{F0: 42, F1: 8}
                                        v1 -= v1
                                        _ = fi167
                                        _ = v168
                                    }
                                    if (832 < 42) {
                                        var v169 uint16 = 0
                                        _ = v169
                                    } else {
                                        v117 += v117
                                        switch (v5 * 1) {
                                        case 152:
                                            switch ((v2 + v2) + (v2 + v2)) {
                                            case "a":
                                                if v3 {
                                                    var v170 uint8 = genFn0(400)
                                                    if v3 {
                                                        v171 := GenS11{F0: 2, F1: v117, F2: v3}
                                                        switch (v2 + v2) {
                                                        case "foo bar":
                                                            if ((v117 | v117) == (^(v117))) {
                                                                var v172 uint64 = uint64(v4)
                                                                _ = v172
                                                            }
                                                        }
                                                        v1 -= fi157
                                                        _ = v171
                                                    }
                                                    _ = v170
                                                }
                                            case "foo bar":
                                                v4 -= v4
                                            default:
                                                v34 += v34
                                            }
                                        case -9223372036854775808:
                                            if (!true) {
                                                var v173 uint16 = 1
                                                var v174 int8 = v117
                                                if v3 {
                                                    if (128 > 1) {
                                                        v175 := v6[v4]
                                                        if v3 {
                                                            v2 = (((v2 + ("foo bar" + v2)) + "") + "hello")
                                                            fmt.Println(v5, v173)
                                                            fmt.Println(v1, fi157, v173)
                                                        }
                                                        v6[2147483648] = (1e10 == 1.5)
                                                        _ = v175
                                                    } else {
                                                        if (93 >= 42) {
                                                            for fi176 := 0; fi176 < 9; fi176++ {
                                                                if false {
                                                                    fmt.Println(v5, v2, v1)
                                                                    v2 = v2
                                                                } else {
                                                                    fmt.Println(fi176, v5, v173)
                                                                    v177 := GenS14{F0: v4, F1: (v4 - v4), F2: -1}
                                                                    _ = v177
                                                                }
                                                                if v3 {
                                                                    var v178 int64 = 1
                                                                    switch (fi157 * 519) {
                                                                    case -1:
                                                                        var v179 uint8 = 2
                                                                        _ = v179
                                                                    case 2147483647:
                                                                        if (fi176 == 42) {
                                                                            v174 -= (-(v117))
                                                                            v1 = v5
                                                                        } else {
                                                                            if v3 {
                                                                                if (v173 < v173) {
                                                                                    if v3 {
                                                                                        var v180 uint16 = v173
                                                                                        v181 := GenS4{F0: 0, F1: 53}
                                                                                        switch (v2 + v2) {
                                                                                        case "":
                                                                                            v182 := GenS15{F0: (v2 + v2)}
                                                                                            _ = v182
                                                                                        case "a":
                                                                                            var v183 float64 = 3.14159
                                                                                            _ = v183
                                                                                        }
                                                                                        _ = v180
                                                                                        _ = v181
                                                                                    }
                                                                                    if ((v117 | v117) <= (v174 * -1)) {
                                                                                        if v3 {
                                                                                            var v184 int8 = v117
                                                                                            _ = v184
                                                                                        }
                                                                                        v185 := []int{(fi157 / (((v5 * v5)) | 1))}
                                                                                        _ = v185
                                                                                    }
                                                                                    var v186 int = fi157
                                                                                    _ = v186
                                                                                } else {
                                                                                    v4 = (v4 >> ((4294967295) & 63))
                                                                                }
                                                                                if true {
                                                                                    v187 := GenS16{F0: 42}
                                                                                    _ = v187
                                                                                }
                                                                                var v188 uint = 0
                                                                                _ = v188
                                                                            }
                                                                            v174 -= (^(v174))
                                                                        }
                                                                    case 32768:
                                                                        v2 = "\n\t"
                                                                    }
                                                                    _ = v178
                                                                }
                                                                _ = fi176
                                                            }
                                                            switch v2 {
                                                            case "hello":
                                                                v189 := GenS17{F0: v2, F1: v173}
                                                                _ = v189
                                                            case "\n\t":
                                                                var v190 int = (v5 & (v5 % (((v1 * (v1 ^ v5))) | 1)))
                                                                _ = v190
                                                            case "":
                                                                if false {
                                                                    var v191 uint8 = 1
                                                                    var v192 bool = (v173 > v173)
                                                                    _ = v191
                                                                    _ = v192
                                                                } else {
                                                                    if (-1 < 32767) {
                                                                        v193 := GenS6{F0: 4294967295, F1: 42}
                                                                        v194 := GenS18{F0: v4, F1: 18446744073709551615, F2: (v173 % (((^(v173))) | 1))}
                                                                        v5 -= (fi157 >> ((608) & 63))
                                                                        _ = v193
                                                                        _ = v194
                                                                    }
                                                                    for fi195 := 0; fi195 < 13; fi195++ {
                                                                        var v196 int = v5
                                                                        _ = fi195
                                                                        _ = v196
                                                                    }
                                                                }
                                                            }
                                                            var v197 int64 = -9223372036854775808
                                                            _ = v197
                                                        }
                                                    }
                                                } else {
                                                    fmt.Println(v4, v34)
                                                    switch v5 {
                                                    case -557:
                                                        v117 = -1
                                                    case -369:
                                                        v198 := []int{v5, (fi157 + (v1 - v1))}
                                                        _ = v198
                                                    case -1:
                                                        if false {
                                                            fmt.Println(v2)
                                                        }
                                                    }
                                                    for fi199 := 0; fi199 < 14; fi199++ {
                                                        switch "" {
                                                        case "\n\t":
                                                            var v200 uint16 = (v173 | (v173 + 973))
                                                            _ = v200
                                                        }
                                                        if v3 {
                                                            switch (v2 + v2) {
                                                            case "":
                                                                if v3 {
                                                                    var v201 int = v1
                                                                    _ = v201
                                                                } else {
                                                                    var v202 float32 = 3.14159
                                                                    _ = v202
                                                                }
                                                            case "a":
                                                                fmt.Println(v34, v174)
                                                            case "foo bar":
                                                                v203 := v6[(v4 - v4)]
                                                                _ = v203
                                                            default:
                                                                var v204 int64 = 1
                                                                fmt.Println(v174)
                                                                _ = v204
                                                            }
                                                            if (((0 <= -806) && (-9223372036854775808 <= (^(fi199)))) && (32767 >= 128)) {
                                                                fmt.Println(v173, v3, v5)
                                                                var v205 int16 = -32768
                                                                _ = v205
                                                            }
                                                        }
                                                        if true {
                                                            v206 := GenS11{F0: 18446744073709551615, F1: (v174 >> ((896) & 63)), F2: true}
                                                            if (((0.0 <= 1e10) && (!true)) || (0 != 1)) {
                                                                switch v5 {
                                                                case 345:
                                                                    var v207 string = v2
                                                                    _ = v207
                                                                case -882:
                                                                    v208 := GenS19{F0: fi199, F1: int8(v173), F2: (("hello" + v2) + v2)}
                                                                    _ = v208
                                                                case 32767:
                                                                    if v3 {
                                                                        v209 := []int{(v5 >> ((uint(v4)) & 63)), fi157, fi157}
                                                                        _ = v209
                                                                    }
                                                                default:
                                                                    switch v1 {
                                                                    case 312:
                                                                        var v210 uint32 = v4
                                                                        _ = v210
                                                                    }
                                                                }
                                                                var v211 int = (-(v1))
                                                                _ = v211
                                                            } else {
                                                                v5 += (v5 ^ v1)
                                                            }
                                                            v212 := GenS20{F0: -2147483648, F1: 2, F2: 0}
                                                            _ = v206
                                                            _ = v212
                                                        } else {
                                                            v3 = ((false || (!v3)) && (!true))
                                                            v213 := GenS21{F0: (v4 / ((v4) | 1))}
                                                            var v214 string = v2
                                                            switch v2 {
                                                            case "":
                                                                v6[(v4 ^ 890)] = true
                                                            case "hello":
                                                                switch v2 {
                                                                case "foo bar":
                                                                    var v215 uint = 18446744073709551615
                                                                    var v216 int32 = 230
                                                                    v5 = v5
                                                                    _ = v215
                                                                    _ = v216
                                                                case "\n\t":
                                                                    if (-9223372036854775808 >= 42) {
                                                                        fmt.Println(v34)
                                                                    }
                                                                case "a":
                                                                    delete(v6, v4)
                                                                }
                                                                if true {
                                                                    fmt.Println(v174, v173)
                                                                    var v217 string = (((v2 + "\n\t") + v2) + v2)
                                                                    _ = v217
                                                                } else {
                                                                    v218 := []int{(-(v5)), v5}
                                                                    v3 = v3
                                                                    _ = v218
                                                                }
                                                            default:
                                                                fmt.Println(v173)
                                                            }
                                                            if (!v3) {
                                                                v5 = 9223372036854775807
                                                                v214 = "a"
                                                            } else {
                                                                var v219 int32 = 732
                                                                _ = v219
                                                            }
                                                            _ = v213
                                                            _ = v214
                                                        }
                                                        _ = fi199
                                                    }
                                                }
                                                _ = v173
                                                _ = v174
                                            } else {
                                                v220 := []int{int(v117), 127}
                                                if v3 {
                                                    v5 -= (v1 - v1)
                                                } else {
                                                    for fi221 := 0; fi221 < 18; fi221++ {
                                                        v2 = v2
                                                        v3 = v3
                                                        v117 = (v117 + (v117 ^ v117))
                                                        _ = fi221
                                                    }
                                                    v5 -= v1
                                                    var v222 string = ((("\n\t" + (v2 + "hello")) + v2) + v2)
                                                    _ = v222
                                                }
                                                if v3 {
                                                    var v223 int64 = 919
                                                    _ = v223
                                                }
                                                _ = v220
                                            }
                                        case 4294967295:
                                            v224 := GenS22{F0: -1, F1: 42, F2: 2}
                                            _ = v224
                                        default:
                                            fmt.Println(v1, v4)
                                        }
                                        v225 := []int{int(v34), 127}
                                        _ = v225
                                    }
                                } else {
                                    var v226 int32 = -2147483648
                                    _ = v226
                                }
                            } else {
                                var v227 uint64 = 916
                                if ((v4 << ((2) & 63)) == 286) {
                                    v4 = (v4 * 4294967295)
                                } else {
                                    fmt.Println(v34, v227)
                                }
                                var v228 int64 = 1
                                _ = v227
                                _ = v228
                            }
                            _ = v166
                        } else {
                            v229 := GenS10{F0: (v4 * (v4 % ((v4) | 1))), F1: 21}
                            _ = v229
                        }
                    }
                    _ = fi157
                }
            }
            fmt.Println(v3, v4, v5)
            for fi230 := 0; fi230 < 17; fi230++ {
                for fi231 := 0; fi231 < 16; fi231++ {
                    v2 = (v2 + v2)
                    _ = fi231
                }
                v2 = (v2 + "foo bar")
                var v232 int = fi230
                _ = fi230
                _ = v232
            }
        }
    case "":
        fmt.Println(v3)
    case "foo bar":
        v233 := []int{(v1 ^ v5), int(v4), -1, v5}
        _ = v233
    }
    v2 = v2
    v2 = (v2 + v2)
    fmt.Println("v1:", v1)
    fmt.Println("v117:", v117)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v115
    _ = v116
    _ = v117
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
