package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int16; F1 bool }
type GenS6 struct { F0 uint; F1 int; F2 int32 }
type GenS7 struct { F0 uint32; F1 bool }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int; F1 int8; F2 string }
type GenS11 struct { F0 int32; F1 uint32; F2 int16 }
type GenS12 struct { F0 int; F1 uint; F2 uint16 }
type GenS13 struct { F0 string }
type GenS14 struct { F0 int64 }
type GenS15 struct { F0 bool }
type GenS16 struct { F0 int; F1 uint8; F2 int16 }
type GenS17 struct { F0 uint32; F1 uint }
func genFn0() int16 {
    var v7 int8 = -1
    if true {
        v8 := []int{-2147483648, -1, 127}
        var v9 int32 = 949
        if true {
            var v10 uint = 18446744073709551615
            _ = v10
        }
        _ = v8
        _ = v9
    }
    switch -65 {
    case -71:
        if (0 < 42) {
            switch 0 {
            case 32768:
                for fi11 := 0; fi11 < 20; fi11++ {
                    v12 := GenS1{F0: 1}
                    v7 += -3
                    var v13 uint32 = 1
                    _ = fi11
                    _ = v12
                    _ = v13
                }
            }
            v14 := []int{9223372036854775807, 42, 32768, 128}
            if false {
                switch 0 {
                case -1:
                    v7 += v7
                }
                fmt.Println(v7)
                for fi15 := 0; fi15 < 3; fi15++ {
                    v16 := GenS1{F0: 4294967295}
                    _ = fi15
                    _ = v16
                }
            }
            _ = v14
        } else {
            v17 := GenS1{F0: 466}
            for fi18 := 0; fi18 < 7; fi18++ {
                var v19 uint32 = 820
                fmt.Println(v7)
                _ = fi18
                _ = v19
            }
            _ = v17
        }
    case -943:
        fmt.Println(v7)
    case 0:
        v20 := []int{0, -82, 42, 0}
        _ = v20
    }
    _ = v7
    return 0
}

func main() {
    var v1 string = ("\n\t" + ((("foo bar" + "a") + "\n\t") + "hello"))
    var v2 float32 = -2.25
    var v3 int = 1
    v4 := make(map[uint32]bool)
    var v5 float64 = 0.1
    var v6 int8 = -128
    if true {
        var v21 uint16 = 1
        if (346 == 32767) {
            for fi22 := 0; fi22 < 18; fi22++ {
                v23 := []int{v3, fi22}
                switch 2147483647 {
                case 4294967295:
                    if (155 == 2147483647) {
                        for fi24 := 0; fi24 < 9; fi24++ {
                            if len(v23) >= 2 {
                                v23 = v23[1:len(v23)]
                            }
                            _ = fi24
                        }
                    } else {
                        for fi25 := 0; fi25 < 16; fi25++ {
                            fmt.Println(v5, v21)
                            _ = fi25
                        }
                        _ = len(v4)
                        var v26 float64 = 1.5
                        _ = v26
                    }
                case 884:
                    v6 += v6
                }
                _ = fi22
                _ = v23
            }
            v6 = 127
            if false {
                switch (v3 + (v3 & v3)) {
                case 699:
                    for fi27 := 0; fi27 < 8; fi27++ {
                        v28 := v4[0]
                        v29 := []int{127}
                        if false {
                            switch (v3 * fi27) {
                            case -9223372036854775808:
                                for fi30 := 0; fi30 < 9; fi30++ {
                                    v31 := GenS2{F0: 2147483647}
                                    _ = fi30
                                    _ = v31
                                }
                            case 9223372036854775807:
                                var v32 string = v1
                                _ = v32
                            default:
                                for fi33 := 0; fi33 < 19; fi33++ {
                                    v1 = ("foo bar" + ("\n\t" + "foo bar"))
                                    _ = fi33
                                }
                            }
                            v3 += v3
                        } else {
                            for fi34 := 0; fi34 < 13; fi34++ {
                                switch -2147483648 {
                                case 42:
                                    if v28 {
                                        fmt.Println(v3, fi34)
                                    }
                                case -9223372036854775808:
                                    v35 := GenS1{F0: 4294967295}
                                    _ = v35
                                case 0:
                                    fmt.Println(v3, fi27)
                                default:
                                    fmt.Println(fi27, v6)
                                }
                                _ = fi34
                            }
                        }
                        _ = fi27
                        _ = v28
                        _ = v29
                    }
                }
                v4[0] = (0.1 == v2)
                if true {
                    var v36 int64 = -1
                    switch 9223372036854775807 {
                    case 826:
                        v37 := GenS2{F0: (v3 / ((v3) | 1))}
                        _ = v37
                    case 42:
                        if false {
                            v4[0] = true
                            if true {
                                v21 += (v21 & v21)
                                fmt.Println(v6)
                            } else {
                                fmt.Println(v21)
                                delete(v4, 276)
                                v4[42] = false
                            }
                        }
                    case 64:
                        fmt.Println(v2)
                    default:
                        fmt.Println(v5)
                    }
                    delete(v4, 2147483647)
                    _ = v36
                }
            }
        } else {
            v21 = v21
            switch v3 {
            case -867:
                for fi38 := 0; fi38 < 16; fi38++ {
                    fmt.Println(v3, v5, v2)
                    var v39 int32 = -1
                    _ = fi38
                    _ = v39
                }
            case 0:
                v6 += 5
            case 780:
                for fi40 := 0; fi40 < 20; fi40++ {
                    var v41 float64 = float64(((0.0 / v2) - 0.1))
                    var v42 uint = 255
                    if true {
                        switch -2147483648 {
                        case 238:
                            if false {
                                fmt.Println(v6, v42, v5)
                                v43 := GenS3{F0: ("" + v1)}
                                for fi44 := 0; fi44 < 3; fi44++ {
                                    v2 = ((v2 / (v2 * float32((float64(v2) + v5)))) + (float32(-2.25) / v2))
                                    var v45 int64 = 9223372036854775807
                                    _ = fi44
                                    _ = v45
                                }
                                _ = v43
                            } else {
                                var v46 string = "hello"
                                _ = v46
                            }
                        }
                    } else {
                        if true {
                            v47 := v4[2147483647]
                            for fi48 := 0; fi48 < 11; fi48++ {
                                var v49 uint8 = 144
                                _ = fi48
                                _ = v49
                            }
                            _ = v47
                        }
                    }
                    _ = fi40
                    _ = v41
                    _ = v42
                }
            default:
                v50 := GenS4{F0: 255}
                _ = v50
            }
        }
        v4[746] = false
        _ = v21
    }
    switch (v3 * v3) {
    case -189:
        var v51 uint8 = 128
        _ = v51
    case -85:
        var v52 uint64 = 2
        _ = v52
    }
    switch (v3 | (-(v3))) {
    case 811:
        if ((v3 / ((864) | 1)) >= v3) {
            if false {
                var v53 uint = 65535
                _ = v53
            } else {
                for fi54 := 0; fi54 < 17; fi54++ {
                    v55 := GenS2{F0: (fi54 & (fi54 - fi54))}
                    var v56 float32 = (v2 / float32((v5 + float64((v2 / ((-2.25 * float32(v5)) - v2))))))
                    switch v3 {
                    case 1:
                        fmt.Println(v6)
                    }
                    _ = fi54
                    _ = v55
                    _ = v56
                }
                v57 := v4[1]
                _ = v57
            }
            for fi58 := 0; fi58 < 7; fi58++ {
                switch v3 {
                case 2147483647:
                    for fi59 := 0; fi59 < 8; fi59++ {
                        v60 := GenS5{F0: -1, F1: (0 != 2)}
                        _ = fi59
                        _ = v60
                    }
                default:
                    v61 := GenS2{F0: v3}
                    _ = v61
                }
                _ = fi58
            }
            for fi62 := 0; fi62 < 15; fi62++ {
                switch v1 {
                case "":
                    v63 := []int{(fi62 | (^(v3))), (^(fi62))}
                    _ = v63
                case "hello":
                    v64 := []int{(v3 % (((fi62 & v3)) | 1))}
                    _ = v64
                case "a":
                    v65 := GenS4{F0: 141}
                    _ = v65
                default:
                    v5 = (float64(float32((v5 * (v5 - (float64((v2 + v2)) * 3.14159))))) / v5)
                }
                v4[65535] = false
                _ = fi62
            }
        }
    case -918:
        fmt.Println(v5, v3, v2)
    }
    v66 := GenS4{F0: 931}
    if false {
        for fi67 := 0; fi67 < 15; fi67++ {
            fmt.Println(fi67)
            var v68 int = v3
            switch (v68 ^ 2147483647) {
            case 957:
                if false {
                    fmt.Println(v5, v1)
                    var v69 int32 = -1
                    var v70 uint8 = 1
                    _ = v69
                    _ = v70
                } else {
                    var v71 int32 = 42
                    switch "foo bar" {
                    case "hello":
                        fmt.Println(v6, v5, fi67)
                    case "\n\t":
                        for fi72 := 0; fi72 < 5; fi72++ {
                            v73 := GenS6{F0: 0, F1: v3, F2: (v71 * v71)}
                            v4[949] = false
                            _ = fi72
                            _ = v73
                        }
                    case "":
                        var v74 int32 = (v71 + 1)
                        _ = v74
                    default:
                        fmt.Println(v6, v71)
                    }
                    v75 := []int{v68, (-(v68)), v68, 2147483647}
                    _ = v71
                    _ = v75
                }
            case 9223372036854775807:
                v6 -= (v6 ^ (v6 - 100))
            default:
                switch v1 {
                case "":
                    var v76 int8 = 100
                    _ = v76
                case "foo bar":
                    var v77 bool = true
                    _ = v77
                default:
                    var v78 uint = 1
                    _ = v78
                }
            }
            _ = fi67
            _ = v68
        }
        for fi79 := 0; fi79 < 7; fi79++ {
            var v80 uint = 18446744073709551615
            var v81 float64 = 1.5
            v82 := GenS7{F0: 1, F1: (!(122 < 0))}
            _ = fi79
            _ = v80
            _ = v81
            _ = v82
        }
    }
    if (255 >= 0) {
        for fi83 := 0; fi83 < 2; fi83++ {
            for fi84 := 0; fi84 < 8; fi84++ {
                v5 -= (-2.25 - 3.14159)
                switch "a" {
                case "\n\t":
                    if (4294967295 >= 2147483648) {
                        fmt.Println(fi83, fi84, v2)
                    } else {
                        switch (fi84 % ((1) | 1)) {
                        case 834:
                            fmt.Println(fi84)
                        case 322:
                            var v85 string = "a"
                            _ = v85
                        case -9223372036854775808:
                            var v86 uint32 = 509
                            _ = v86
                        default:
                            var v87 int = -562
                            _ = v87
                        }
                    }
                case "a":
                    switch "foo bar" {
                    case "\n\t":
                        if ((!(0 >= -1)) && (v5 < (float64(((((float32(-2.25) + -2.25) + (v2 / (float32(v5) * (v2 - (float32(v5) * v2))))) + ((0.0 * v2) / v2)) * float32(0.1))) - (v5 / v5)))) {
                            v2 = v2
                        }
                    case "":
                        switch 127 {
                        case 2147483647:
                            v88 := v4[42]
                            _ = v88
                        case -1:
                            if (42 > 18446744073709551615) {
                                fmt.Println(fi84, v6)
                                v2 -= v2
                                var v89 float64 = v5
                                _ = v89
                            } else {
                                v3 = 1
                                if false {
                                    v1 = (v1 + "a")
                                    if true {
                                        if true {
                                            v90 := []int{v3, 1, 2147483647, (fi84 << ((491) & 63))}
                                            switch v1 {
                                            case "":
                                                if false {
                                                    if true {
                                                        if true {
                                                            if false {
                                                                v90 = append(v90, -9223372036854775808)
                                                                if ((!false) || (((131 == 164) || false) && ((!((0 <= 2) && (!false))) || true))) {
                                                                    var v91 int = (-(v3))
                                                                    v92 := GenS1{F0: 1}
                                                                    var v93 uint = 749
                                                                    _ = v91
                                                                    _ = v92
                                                                    _ = v93
                                                                }
                                                            } else {
                                                                if true {
                                                                    fmt.Println(v2, v5, v2)
                                                                } else {
                                                                    fmt.Println(v6)
                                                                    fmt.Println(v1)
                                                                    fmt.Println(v2, fi83, v3)
                                                                }
                                                                var v94 bool = false
                                                                if len(v90) > 0 {
                                                                    v90[0] = 799
                                                                }
                                                                _ = v94
                                                            }
                                                        } else {
                                                            v95 := GenS8{F0: 2}
                                                            switch "a" {
                                                            case "foo bar":
                                                                var v96 string = v1
                                                                _ = v96
                                                            default:
                                                                var v97 int = (fi83 + (-(v3)))
                                                                _ = v97
                                                            }
                                                            var v98 float32 = v2
                                                            _ = v95
                                                            _ = v98
                                                        }
                                                    } else {
                                                        v5 += v5
                                                        v6 -= -1
                                                        v6 += (v6 / ((v6) | 1))
                                                    }
                                                }
                                            case "foo bar":
                                                switch fi84 {
                                                case 183:
                                                    var v99 uint = 1
                                                    _ = v99
                                                case 421:
                                                    v6 = v6
                                                case 9223372036854775807:
                                                    v1 = v1
                                                default:
                                                    var v100 uint8 = 128
                                                    _ = v100
                                                }
                                            }
                                            if false {
                                                switch -780 {
                                                case 9223372036854775807:
                                                    v101 := v4[112]
                                                    _ = v101
                                                case -2147483648:
                                                    var v102 uint8 = 42
                                                    _ = v102
                                                default:
                                                    var v103 uint32 = 626
                                                    _ = v103
                                                }
                                                var v104 uint16 = 604
                                                var v105 int8 = -26
                                                _ = v104
                                                _ = v105
                                            }
                                            _ = v90
                                        } else {
                                            if true {
                                                if true {
                                                    v2 = 0.1
                                                } else {
                                                    var v106 uint8 = 8
                                                    if false {
                                                        v1 = v1
                                                        v107 := []int{fi83}
                                                        if len(v107) >= 2 {
                                                            v107 = v107[1:len(v107)]
                                                        }
                                                        _ = v107
                                                    } else {
                                                        v108 := []int{}
                                                        if (false && false) {
                                                            v6 = (-(v6))
                                                        } else {
                                                            var v109 int = 4294967295
                                                            _ = len(v108)
                                                            _ = v109
                                                        }
                                                        v110 := GenS9{F0: 0}
                                                        var v111 bool = false
                                                        _ = v108
                                                        _ = v110
                                                        _ = v111
                                                    }
                                                    _ = v106
                                                }
                                                v2 += (v2 - v2)
                                            }
                                            v6 = 74
                                            v112 := []int{(v3 * (-(fi84))), (v3 >> ((2) & 63)), (-(fi84)), -621}
                                            _ = v112
                                        }
                                        if (-2147483648 == -1) {
                                            var v113 int16 = 32767
                                            v114 := GenS7{F0: 2147483647, F1: (0 >= -2147483648)}
                                            v113 += v113
                                            _ = v113
                                            _ = v114
                                        } else {
                                            if true {
                                                v115 := []int{-997, (fi83 ^ -975), -2147483648}
                                                _ = v115
                                            } else {
                                                v116 := []int{v3, fi83}
                                                if false {
                                                    v1 = (v1 + v1)
                                                } else {
                                                    switch ("foo bar" + v1) {
                                                    case "\n\t":
                                                        var v117 uint8 = 0
                                                        _ = v117
                                                    default:
                                                        _ = len(v116)
                                                    }
                                                    switch (fi84 * 2147483647) {
                                                    case 9223372036854775807:
                                                        var v118 int32 = -1
                                                        _ = v118
                                                    default:
                                                        fmt.Println(fi84, v3, fi83)
                                                    }
                                                }
                                                var v119 int = (fi84 / (((^(fi83))) | 1))
                                                _ = v116
                                                _ = v119
                                            }
                                            var v120 uint16 = 0
                                            var v121 int8 = 42
                                            if (-567 == -1) {
                                                v122 := GenS6{F0: 255, F1: v3, F2: -2147483648}
                                                _ = v122
                                            }
                                            _ = v120
                                            _ = v121
                                        }
                                    } else {
                                        fmt.Println(v1)
                                        v123 := []int{-2147483648}
                                        v2 = ((float32(v5) - v2) + (((v2 - (v2 - v2)) - v2) + (float32(v5) + v2)))
                                        _ = v123
                                    }
                                    v124 := GenS10{F0: 93, F1: -1, F2: "hello"}
                                    _ = v124
                                } else {
                                    v1 = v1
                                    if true {
                                        v1 = (v1 + "foo bar")
                                        var v125 uint = 826
                                        v126 := GenS3{F0: v1}
                                        switch "" {
                                        case "\n\t":
                                            v125 = (v125 % ((v125) | 1))
                                        case "foo bar":
                                            if (false && true) {
                                                var v127 float32 = (((v2 * (-2.25 + (v2 - -2.25))) - (v2 / v2)) * v2)
                                                _ = v127
                                            } else {
                                                v128 := GenS11{F0: -2147483648, F1: 278, F2: -1}
                                                fmt.Println(v6)
                                                _ = v128
                                            }
                                        }
                                        v129 := GenS12{F0: fi83, F1: 18446744073709551615, F2: 255}
                                        switch v1 {
                                        case "a":
                                            fmt.Println(v125, v2)
                                        }
                                        v3 = -9223372036854775808
                                        v130 := GenS13{F0: v1}
                                        v5 = (((v5 + v5) / (v5 + v5)) * (v5 / 1e10))
                                        _ = v125
                                        _ = v126
                                        _ = v129
                                        _ = v130
                                    }
                                }
                            }
                        }
                    }
                default:
                    if false {
                        if false {
                            v131 := GenS14{F0: -1}
                            var v132 uint16 = 1
                            _ = v131
                            _ = v132
                        }
                    } else {
                        switch v1 {
                        case "a":
                            fmt.Println(fi84, v1)
                        case "hello":
                            v133 := GenS15{F0: (587 > 65535)}
                            _ = v133
                        case "":
                            v4[638] = true
                        }
                        v134 := GenS16{F0: 128, F1: 127, F2: -32768}
                        if (2147483647 != -475) {
                            v3 -= -2147483648
                            v135 := []int{fi84, (fi84 % (((-(fi84))) | 1)), 1}
                            _ = v135
                        }
                        _ = v134
                    }
                }
                _ = fi84
            }
            v136 := GenS9{F0: 617}
            v2 = float32(((0.0 * v5) / float64(3.14159)))
            _ = fi83
            _ = v136
        }
    } else {
        for fi137 := 0; fi137 < 13; fi137++ {
            fmt.Println(v6, v3)
            var v138 int8 = (v6 * (v6 >> ((853) & 63)))
            if false {
                if (true && (!("a" >= (v1 + ("foo bar" + "\n\t"))))) {
                    v139 := v4[2147483648]
                    var v140 string = v1
                    v141 := GenS16{F0: (v3 & (-(fi137))), F1: 42, F2: 0}
                    _ = v139
                    _ = v140
                    _ = v141
                }
            } else {
                for fi142 := 0; fi142 < 11; fi142++ {
                    var v143 string = "a"
                    var v144 uint32 = 149
                    var v145 uint64 = 4294967295
                    fmt.Println(v6, fi142, v3)
                    var v146 float64 = 0.1
                    if ((3.14159 - ((-2.25 + v2) + 0.1)) <= v2) {
                        v1 = v1
                        if true {
                            v1 = v1
                        } else {
                            if true {
                                switch (fi142 | v3) {
                                case 1:
                                    fmt.Println(v143, v5, v138)
                                default:
                                    v138 -= (v6 ^ (v138 * v6))
                                }
                            }
                            var v147 uint16 = 255
                            if true {
                                v148 := GenS17{F0: 0, F1: 42}
                                switch v1 {
                                case "foo bar":
                                    fmt.Println(v146)
                                    switch ("" + v143) {
                                    case "\n\t":
                                        v149 := []int{-9223372036854775808, v3, fi137}
                                        _ = v149
                                    case "a":
                                        v1 = v143
                                    }
                                case "hello":
                                    var v150 float64 = -2.25
                                    _ = v150
                                default:
                                    v3 -= (^(fi142))
                                    if false {
                                        var v151 int = fi142
                                        switch v151 {
                                        case -723:
                                            v152 := GenS15{F0: true}
                                            _ = v152
                                        }
                                        v138 = v6
                                        _ = v151
                                    } else {
                                        var v153 int = (fi142 & fi142)
                                        var v154 uint = 4294967295
                                        if false {
                                            v155 := v4[65535]
                                            _ = v155
                                        }
                                        _ = v153
                                        _ = v154
                                    }
                                    var v156 uint8 = 211
                                    fmt.Println(v5, v156, v138)
                                    _ = v156
                                }
                                _ = v148
                            }
                            _ = v147
                        }
                        v145 = 2
                    }
                    _ = fi142
                    _ = v143
                    _ = v144
                    _ = v145
                    _ = v146
                }
                v6 = (v6 * (v138 | (v138 & 1)))
            }
            _ = fi137
            _ = v138
        }
        v157 := []int{4294967295, (v3 % ((v3) | 1))}
        _ = v157
    }
    var v158 int64 = -853
    var v159 int16 = 32767
    v160 := GenS9{F0: 4294967295}
    var v161 int16 = v159
    for fi162 := 0; fi162 < 16; fi162++ {
        v161 = (v159 * v161)
        fmt.Println(v161, v2)
        _ = fi162
    }
    var v163 int16 = v161
    var v164 int = v3
    v163 -= genFn0()
    fmt.Println("v1:", v1)
    fmt.Println("v158:", v158)
    fmt.Println("v159:", v159)
    fmt.Println("v161:", v161)
    fmt.Println("v163:", v163)
    fmt.Println("v164:", v164)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v158
    _ = v159
    _ = v160
    _ = v161
    _ = v163
    _ = v164
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v66
}
