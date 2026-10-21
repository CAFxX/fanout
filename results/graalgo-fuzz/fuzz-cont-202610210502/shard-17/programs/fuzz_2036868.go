package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 int64; F2 string }
type GenS4 struct { F0 int16; F1 int8; F2 int }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint32; F1 uint8 }
type GenS8 struct { F0 uint32; F1 bool }
type GenS9 struct { F0 uint16; F1 uint8 }
type GenS10 struct { F0 int64; F1 int8 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 string; F1 string }
type GenS13 struct { F0 uint8; F1 uint64; F2 uint8 }
func genFn0() uint16 {
    var v4 bool = true
    fmt.Println(v4)
    var v5 float32 = 1.5
    _ = v4
    _ = v5
    return 255
}
func genFn1(p0 string, p1 int64, p2 int) int8 {
    var v6 uint = 403
    fmt.Println(p1)
    switch (^(p2)) {
    case 4294967295:
        switch (p2 | p2) {
        case -130:
            v6 += v6
        default:
            v7 := []int{}
            _ = v7
        }
    case 2147483647:
        v8 := GenS1{F0: p0, F1: true}
        _ = v8
    case 9223372036854775807:
        v9 := []int{(p2 >> (((v6 & v6)) & 63)), -584}
        _ = v9
    default:
        v10 := GenS2{F0: 170}
        _ = v10
    }
    var v11 float64 = (1.5 * (-2.25 + 0.0))
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v6
    return 1
}

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[string]int)
    var v3 float32 = 1.5
    v1 += v1
    var v12 float64 = 1.5
    v13 := GenS2{F0: 2}
    if (0 > 42) {
        v14 := []int{9223372036854775807, 9223372036854775807}
        if true {
            for fi15 := 0; fi15 < 11; fi15++ {
                v12 = v1
                var v16 uint8 = 54
                _ = fi15
                _ = v16
            }
            v17 := GenS2{F0: 371}
            v3 = (float32(0.0) / float32(1.5))
            _ = v17
        } else {
            v3 -= (1e10 + 0.0)
        }
        _ = v14
    }
    var v18 string = ((("\n\t" + ("a" + ("hello" + ("hello" + "\n\t")))) + "") + "\n\t")
    if (false && (0 > -580)) {
        if (v18 >= "a") {
            delete(v2, "\n\t")
            if (42 > 0) {
                delete(v2, v18)
                v19 := v2[v18]
                if false {
                    var v20 int16 = 100
                    v21 := GenS2{F0: 0}
                    v20 += 0
                    _ = v20
                    _ = v21
                } else {
                    if (true && false) {
                        if true {
                            v12 = 1e10
                            var v22 uint8 = 0
                            v23 := GenS2{F0: 0}
                            _ = v22
                            _ = v23
                        } else {
                            var v24 int8 = 67
                            v25 := v2[v18]
                            _ = v24
                            _ = v25
                        }
                        v18 = v18
                        if false {
                            fmt.Println(v18)
                        }
                    } else {
                        if (!(1 == 2147483648)) {
                            var v26 uint32 = 981
                            _ = v26
                        }
                        v27 := v2["hello"]
                        _ = v27
                    }
                    v28 := []int{-9223372036854775808, v19}
                    _ = v28
                }
                _ = v19
            } else {
                v29 := GenS3{F0: 127, F1: -812, F2: v18}
                if (false && true) {
                    if false {
                        fmt.Println(v12, v3)
                        if (true && true) {
                            var v30 uint64 = 42
                            var v31 int32 = 434
                            var v32 int = -249
                            _ = v30
                            _ = v31
                            _ = v32
                        } else {
                            var v33 int16 = -32768
                            v34 := GenS4{F0: -32768, F1: -1, F2: -994}
                            if true {
                                var v35 uint = 65535
                                for fi36 := 0; fi36 < 14; fi36++ {
                                    v35 += (v35 / ((v35) | 1))
                                    v33 -= int16(fi36)
                                    _ = fi36
                                }
                                _ = v35
                            }
                            _ = v33
                            _ = v34
                        }
                        var v37 int8 = genFn1(("foo bar" + v18), 1, -837)
                        _ = v37
                    }
                    if (-32768 >= 42) {
                        var v38 uint8 = 254
                        if true {
                            switch 127 {
                            case 745:
                                switch v18 {
                                case "foo bar":
                                    for fi39 := 0; fi39 < 6; fi39++ {
                                        if (v18 != v18) {
                                            for fi40 := 0; fi40 < 8; fi40++ {
                                                v12 = 1e10
                                                _ = fi40
                                            }
                                            for fi41 := 0; fi41 < 11; fi41++ {
                                                v42 := GenS1{F0: (v18 + v18), F1: true}
                                                var v43 uint64 = 4294967295
                                                v43 -= uint64(v38)
                                                _ = fi41
                                                _ = v42
                                                _ = v43
                                            }
                                        } else {
                                            v44 := GenS5{F0: -1}
                                            switch 1 {
                                            case 128:
                                                if false {
                                                    var v45 uint64 = 4294967295
                                                    var v46 uint32 = 810
                                                    _ = v45
                                                    _ = v46
                                                } else {
                                                    fmt.Println(v38)
                                                }
                                            default:
                                                v18 = v18
                                            }
                                            delete(v2, (("a" + v18) + v18))
                                            _ = v44
                                        }
                                        var v47 float64 = v12
                                        _ = fi39
                                        _ = v47
                                    }
                                case "hello":
                                    fmt.Println(v3, v3)
                                default:
                                    for fi48 := 0; fi48 < 15; fi48++ {
                                        v49 := []int{fi48, (fi48 ^ (fi48 & (fi48 ^ fi48))), 2147483647}
                                        for fi50 := 0; fi50 < 12; fi50++ {
                                            _ = len(v49)
                                            _ = fi50
                                        }
                                        _ = fi48
                                        _ = v49
                                    }
                                }
                            case 882:
                                var v51 float64 = v12
                                _ = v51
                            case 9223372036854775807:
                                for fi52 := 0; fi52 < 10; fi52++ {
                                    for fi53 := 0; fi53 < 11; fi53++ {
                                        v54 := []int{(fi52 + 1), fi52}
                                        _ = fi53
                                        _ = v54
                                    }
                                    v12 -= (float64((float32(v12) + 0.1)) * ((v1 * ((v12 - ((v12 * (3.14159 - float64((v3 + (v3 * (1.5 * v3)))))) + v1)) / v12)) / (0.1 * ((v1 * v12) / 3.14159))))
                                    var v55 uint32 = 0
                                    _ = fi52
                                    _ = v55
                                }
                            default:
                                switch ((v18 + ((v18 + ("foo bar" + "hello")) + (v18 + "a"))) + "hello") {
                                case "\n\t":
                                    _ = len(v2)
                                case "":
                                    if false {
                                        v56 := v2[("a" + v18)]
                                        for fi57 := 0; fi57 < 18; fi57++ {
                                            v56 -= (-(v56))
                                            var v58 string = v18
                                            v12 += float64(v3)
                                            _ = fi57
                                            _ = v58
                                        }
                                        _ = v56
                                    }
                                default:
                                    v59 := []int{127}
                                    _ = v59
                                }
                            }
                            if true {
                                switch 17 {
                                case -507:
                                    fmt.Println(v12)
                                case 567:
                                    v60 := GenS6{F0: -1}
                                    _ = v60
                                }
                                switch 9223372036854775807 {
                                case -9223372036854775808:
                                    v38 += 128
                                case -721:
                                    switch v18 {
                                    case "hello":
                                        v61 := GenS7{F0: 2147483647, F1: v38}
                                        _ = v61
                                    case "\n\t":
                                        v12 = v12
                                    case "":
                                        v3 = ((v3 + -2.25) * 0.0)
                                    }
                                case 1:
                                    v62 := GenS2{F0: 65535}
                                    _ = v62
                                }
                            }
                            var v63 uint16 = genFn0()
                            _ = v63
                        } else {
                            v38 -= v38
                            v64 := []int{955, -9223372036854775808}
                            v65 := GenS5{F0: -888}
                            _ = v64
                            _ = v65
                        }
                        if (false || false) {
                            var v66 uint = 506
                            var v67 float32 = v3
                            v68 := GenS7{F0: 0, F1: v38}
                            _ = v66
                            _ = v67
                            _ = v68
                        }
                        _ = v38
                    } else {
                        if (!false) {
                            fmt.Println(v1, v1, v12)
                            var v69 int16 = -1
                            for fi70 := 0; fi70 < 12; fi70++ {
                                v18 = v18
                                _ = fi70
                            }
                            _ = v69
                        } else {
                            var v71 int16 = 0
                            switch 32767 {
                            case -1:
                                v72 := []int{}
                                _ = v72
                            case -705:
                                fmt.Println(v12, v18)
                            default:
                                for fi73 := 0; fi73 < 13; fi73++ {
                                    if false {
                                        v1 = (v1 * 1.5)
                                        switch -503 {
                                        case 127:
                                            if (true && (65535 <= 713)) {
                                                var v74 float64 = v1
                                                _ = v74
                                            } else {
                                                var v75 float32 = (v3 + float32((v1 + v12)))
                                                var v76 uint64 = 42
                                                _ = v75
                                                _ = v76
                                            }
                                        default:
                                            var v77 int16 = (v71 ^ (v71 % ((v71) | 1)))
                                            _ = v77
                                        }
                                        var v78 uint8 = 0
                                        _ = v78
                                    } else {
                                        v79 := GenS4{F0: (v71 & (v71 + 1)), F1: 66, F2: fi73}
                                        if false {
                                            for fi80 := 0; fi80 < 11; fi80++ {
                                                v81 := GenS8{F0: 873, F1: (false && (false && (v1 < (((v1 - ((v12 * v1) / 1e10)) + 0.0) * v12))))}
                                                _ = fi80
                                                _ = v81
                                            }
                                            switch "\n\t" {
                                            case "a":
                                                v3 = -2.25
                                            case "":
                                                v3 = 0.0
                                            case "hello":
                                                switch v18 {
                                                case "\n\t":
                                                    fmt.Println(v12)
                                                case "hello":
                                                    v71 = (v71 >> ((18446744073709551615) & 63))
                                                case "":
                                                    if (!(42 <= 9223372036854775807)) {
                                                        var v82 uint16 = genFn0()
                                                        if (42 < 42) {
                                                            var v83 int = (fi73 * fi73)
                                                            _ = v83
                                                        }
                                                        v84 := []int{fi73}
                                                        _ = v82
                                                        _ = v84
                                                    } else {
                                                        v71 = v71
                                                        if (v18 > (v18 + "")) {
                                                            if true {
                                                                fmt.Println(v71)
                                                            }
                                                            switch (v18 + (("" + "hello") + v18)) {
                                                            case "a":
                                                                if (340 < 2) {
                                                                    v85 := []int{fi73, fi73}
                                                                    _ = v85
                                                                } else {
                                                                    var v86 int64 = -725
                                                                    _ = v86
                                                                }
                                                            case "\n\t":
                                                                var v87 int64 = -1
                                                                _ = v87
                                                            case "hello":
                                                                v88 := GenS6{F0: (fi73 >> ((706) & 63))}
                                                                _ = v88
                                                            default:
                                                                v12 -= float64(1.5)
                                                            }
                                                        }
                                                        for fi89 := 0; fi89 < 12; fi89++ {
                                                            var v90 float32 = ((0.1 + ((0.0 / v3) * 3.14159)) / v3)
                                                            v12 = v12
                                                            _ = fi89
                                                            _ = v90
                                                        }
                                                    }
                                                default:
                                                    var v91 uint16 = 32767
                                                    _ = v91
                                                }
                                            }
                                            v92 := GenS1{F0: "foo bar", F1: (255 != 286)}
                                            _ = v92
                                        } else {
                                            for fi93 := 0; fi93 < 8; fi93++ {
                                                var v94 float64 = float64((0.1 + -2.25))
                                                if true {
                                                    switch v18 {
                                                    case "":
                                                        var v95 int16 = (v71 + (v71 ^ (-(v71))))
                                                        _ = v95
                                                    case "a":
                                                        fmt.Println(fi93, fi73, v71)
                                                    default:
                                                        v3 += v3
                                                    }
                                                    v3 -= float32(v94)
                                                }
                                                _ = fi93
                                                _ = v94
                                            }
                                        }
                                        _ = v79
                                    }
                                    _ = fi73
                                }
                            }
                            switch 611 {
                            case 833:
                                v3 = 3.14159
                            }
                            _ = v71
                        }
                        switch (v18 + v18) {
                        case "hello":
                            v18 = (v18 + (v18 + (v18 + "hello")))
                        }
                        v1 += (v12 - v12)
                    }
                    v96 := v2[v18]
                    _ = v96
                }
                var v97 float64 = 1.5
                _ = v29
                _ = v97
            }
            var v98 int = 872
            _ = v98
        }
    } else {
        v1 = v12
        fmt.Println(v12)
        var v99 int16 = -1
        _ = v99
    }
    for fi100 := 0; fi100 < 3; fi100++ {
        for fi101 := 0; fi101 < 19; fi101++ {
            var v102 float32 = v3
            var v103 uint32 = 479
            _ = fi101
            _ = v102
            _ = v103
        }
        _ = fi100
    }
    var v104 float32 = v3
    var v105 uint = 1
    v106 := GenS7{F0: 280, F1: uint8(v105)}
    if true {
        if false {
            if (false || false) {
                v107 := GenS3{F0: -892, F1: 1, F2: v18}
                v1 += float64(1.5)
                v108 := v2["hello"]
                _ = v107
                _ = v108
            } else {
                fmt.Println(v12, v1)
                if false {
                    var v109 int8 = 100
                    v110 := v2[(((((v18 + (v18 + "a")) + "a") + "") + v18) + (v18 + (v18 + (("a" + v18) + ("foo bar" + v18)))))]
                    _ = v109
                    _ = v110
                } else {
                    if true {
                        var v111 int16 = 32767
                        _ = v111
                    } else {
                        for fi112 := 0; fi112 < 16; fi112++ {
                            v113 := GenS7{F0: 902, F1: 2}
                            if (42 >= 32) {
                                var v114 int = fi112
                                v115 := []int{fi112, (^(v114))}
                                if ((-(fi112)) != fi112) {
                                    var v116 uint = (v105 / (((v105 % ((v105) | 1))) | 1))
                                    for fi117 := 0; fi117 < 17; fi117++ {
                                        switch "" {
                                        case "":
                                            switch (fi117 >> ((v116) & 63)) {
                                            case -9223372036854775808:
                                                var v118 uint8 = 170
                                                _ = v118
                                            }
                                        }
                                        v2[((v18 + ((v18 + v18) + ((v18 + "") + ""))) + (v18 + ((v18 + (v18 + (((v18 + "foo bar") + v18) + v18))) + v18)))] = (fi112 % (((^(v114))) | 1))
                                        v119 := GenS7{F0: uint32(v105), F1: 0}
                                        _ = fi117
                                        _ = v119
                                    }
                                    v114 += (v114 / (((fi112 >> ((255) & 63))) | 1))
                                    _ = v116
                                } else {
                                    delete(v2, ((v18 + "hello") + (v18 + v18)))
                                    v2[v18] = 845
                                }
                                _ = v114
                                _ = v115
                            }
                            if false {
                                var v120 float64 = v1
                                v12 += v1
                                _ = v120
                            } else {
                                v121 := v2[((v18 + (v18 + v18)) + v18)]
                                if (-2.25 >= -2.25) {
                                    for fi122 := 0; fi122 < 4; fi122++ {
                                        v123 := GenS4{F0: 1000, F1: 100, F2: (v121 - 9223372036854775807)}
                                        var v124 bool = true
                                        _ = fi122
                                        _ = v123
                                        _ = v124
                                    }
                                    v125 := v2[((v18 + (v18 + v18)) + ((v18 + v18) + ("\n\t" + v18)))]
                                    switch v18 {
                                    case "hello":
                                        v104 = v104
                                    case "\n\t":
                                        for fi126 := 0; fi126 < 5; fi126++ {
                                            v127 := GenS3{F0: (fi126 | fi126), F1: 42, F2: v18}
                                            _ = fi126
                                            _ = v127
                                        }
                                    case "":
                                        switch 0 {
                                        case 244:
                                            v128 := GenS2{F0: 674}
                                            _ = v128
                                        case -1000:
                                            if true {
                                                v129 := GenS9{F0: 649, F1: 255}
                                                v3 -= (float32((v1 / v1)) + 1.5)
                                                _ = v129
                                            }
                                        case -2147483648:
                                            _ = len(v2)
                                        default:
                                            v130 := []int{-241, (-(fi112)), v121}
                                            _ = v130
                                        }
                                    default:
                                        v131 := v2[v18]
                                        _ = v131
                                    }
                                    _ = v125
                                }
                                var v132 float64 = float64((float32(1.5) - (float32(0.0) + 0.1)))
                                _ = v121
                                _ = v132
                            }
                            _ = fi112
                            _ = v113
                        }
                        v133 := GenS10{F0: -488, F1: 88}
                        _ = v133
                    }
                    var v134 int16 = 1
                    for fi135 := 0; fi135 < 16; fi135++ {
                        v12 += (1.5 - (((v1 / -2.25) - (1.5 + (1e10 + 0.0))) / (0.0 * float64(float32(v1)))))
                        var v136 int32 = 1
                        _ = fi135
                        _ = v136
                    }
                    _ = v134
                }
            }
            v1 -= v1
        } else {
            switch "" {
            case "a":
                v18 = "hello"
            case "foo bar":
                v137 := GenS1{F0: v18, F1: (true && ((v105 > v105) || false))}
                _ = v137
            default:
                fmt.Println(v3, v1, v18)
            }
            var v138 float32 = v104
            _ = v138
        }
        v139 := v2[v18]
        _ = v139
    } else {
        v104 = v3
    }
    var v140 int8 = 42
    v141 := GenS11{F0: -9223372036854775808}
    switch -200 {
    case -9223372036854775808:
        if (4294967295 > 4294967295) {
            if true {
                v140 = (v140 >> (((v105 % ((1) | 1))) & 63))
            } else {
                v142 := v2[v18]
                v1 -= v12
                v143 := []int{v142, v142}
                _ = v142
                _ = v143
            }
        } else {
            v12 += v1
            switch ("foo bar" + v18) {
            case "a":
                v144 := []int{}
                _ = v144
            }
        }
    default:
        fmt.Println(v140)
    }
    if true {
        v18 = (v18 + (v18 + v18))
    } else {
        for fi145 := 0; fi145 < 9; fi145++ {
            v2[v18] = fi145
            for fi146 := 0; fi146 < 15; fi146++ {
                if false {
                    v147 := GenS8{F0: 2, F1: true}
                    var v148 bool = ((3.14159 + v104) == (1e10 + (-2.25 * v104)))
                    _ = v147
                    _ = v148
                } else {
                    var v149 int32 = int32(v105)
                    _ = v149
                }
                v150 := []int{(fi146 * (fi145 | fi146)), fi146, -2147483648, 32768}
                v3 += 1.5
                _ = fi146
                _ = v150
            }
            _ = fi145
        }
        switch -327 {
        case -827:
            switch "foo bar" {
            case "":
                var v151 int64 = -925
                _ = v151
            default:
                for fi152 := 0; fi152 < 15; fi152++ {
                    v3 -= v3
                    _ = fi152
                }
            }
        case 32767:
            v105 = v105
        case -9223372036854775808:
            v153 := []int{-711, 116, 0, -302}
            _ = v153
        default:
            v18 = v18
        }
    }
    v154 := []int{9223372036854775807, 245, 2147483647, 0}
    for fi155 := 0; fi155 < 19; fi155++ {
        v105 -= v105
        var v156 float64 = float64(float32(((v12 + 1e10) / ((-2.25 * (v12 / (v12 / v12))) * (((v12 - 3.14159) * v12) / (v1 / v12))))))
        for fi157 := 0; fi157 < 18; fi157++ {
            v158 := GenS8{F0: 2147483648, F1: false}
            var v159 uint = 2
            if true {
                var v160 bool = false
                switch fi155 {
                case 730:
                    v161 := GenS12{F0: "a", F1: "hello"}
                    _ = v161
                case -62:
                    switch v18 {
                    case "":
                        var v162 int = -905
                        _ = v162
                    case "hello":
                        if false {
                            var v163 bool = (!false)
                            var v164 int32 = -2147483648
                            _ = v163
                            _ = v164
                        } else {
                            fmt.Println(v105, v18)
                            v2[v18] = fi155
                            var v165 float32 = v104
                            v166 := GenS13{F0: 128, F1: 18446744073709551615, F2: 42}
                            _ = v165
                            _ = v166
                        }
                    }
                default:
                    var v167 uint32 = 506
                    _ = v167
                }
                v154 = append(v154, (fi155 << (((v159 - (^(v105)))) & 63)))
                _ = v160
            } else {
                v104 = 0.1
                v140 = 50
                v12 = 0.0
            }
            _ = fi157
            _ = v158
            _ = v159
        }
        _ = fi155
        _ = v156
    }
    fmt.Println("v1:", v1)
    fmt.Println("v104:", v104)
    fmt.Println("v105:", v105)
    fmt.Println("v12:", v12)
    fmt.Println("v140:", v140)
    fmt.Println("slicelen v154:", len(v154), v154)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v104
    _ = v105
    _ = v106
    _ = v12
    _ = v13
    _ = v140
    _ = v141
    _ = v154
    _ = v18
    _ = v2
    _ = v3
}
