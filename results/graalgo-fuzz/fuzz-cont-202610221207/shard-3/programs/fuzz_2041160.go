package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int8; F1 int32; F2 bool }
type GenS5 struct { F0 bool; F1 uint64 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint64; F1 int32; F2 uint }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 uint8; F1 int16 }
type GenS11 struct { F0 bool; F1 uint32; F2 int16 }
type GenS12 struct { F0 uint; F1 uint64; F2 int16 }
type GenS13 struct { F0 bool }
type GenS14 struct { F0 uint }
type GenS15 struct { F0 int8; F1 uint64 }
type GenS16 struct { F0 uint32; F1 bool }
type GenS17 struct { F0 uint16; F1 uint32 }
type GenS18 struct { F0 uint32; F1 int8 }
type GenS19 struct { F0 int; F1 int8; F2 int64 }
type GenS20 struct { F0 int; F1 int; F2 string }
type GenS21 struct { F0 int32 }
type GenS22 struct { F0 uint32; F1 string; F2 bool }
func genFn0(p0 bool, p1 uint64) uint8 {
    v4 := GenS1{F0: 506, F1: 42}
    for fi5 := 0; fi5 < 12; fi5++ {
        p0 = p0
        _ = fi5
    }
    var v6 int64 = 1
    _ = p0
    _ = p1
    _ = v4
    _ = v6
    return 134
}
func genFn1(p0 int64) int8 {
    for fi7 := 0; fi7 < 20; fi7++ {
        for fi8 := 0; fi8 < 14; fi8++ {
            if false {
                var v9 int16 = 1
                _ = v9
            }
            if (true || (210 >= 1)) {
                fmt.Println(fi8, p0)
            }
            switch 127 {
            case 127:
                var v10 uint32 = 579
                _ = v10
            case -275:
                fmt.Println(fi7, fi7, fi7)
            case -9223372036854775808:
                var v11 uint = 42
                v12 := []int{(fi7 | fi8), fi8, 2147483647}
                if (!(!false)) {
                    v11 -= v11
                    if true {
                        v13 := GenS1{F0: 65535, F1: 42}
                        fmt.Println(p0, fi7, p0)
                        _ = v13
                    }
                } else {
                    v11 -= (v11 ^ v11)
                    if (!false) {
                        p0 = -1
                        v11 += v11
                    }
                }
                _ = v11
                _ = v12
            default:
                var v14 int = 9223372036854775807
                _ = v14
            }
            p0 = p0
            _ = fi8
        }
        var v15 uint32 = 2147483647
        switch "foo bar" {
        case "":
            v16 := []int{int(v15), (fi7 >> ((452) & 63))}
            _ = v16
        case "foo bar":
            var v17 uint64 = 18446744073709551615
            _ = v17
        default:
            switch "a" {
            case "a":
                v18 := []int{}
                _ = v18
            }
        }
        _ = fi7
        _ = v15
    }
    p0 += p0
    p0 = (p0 | -1)
    v19 := GenS2{F0: -1}
    _ = p0
    _ = v19
    return 1
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float32 = 0.1
    var v3 int = -824
    v1[2147483648] = false
    for fi20 := 0; fi20 < 3; fi20++ {
        var v21 int64 = 814
        v3 -= (v3 % ((fi20) | 1))
        switch v3 {
        case -1:
            fmt.Println(v2, v21)
        default:
            var v22 uint = 4294967295
            _ = v22
        }
        _ = fi20
        _ = v21
    }
    v3 = (v3 << ((180) & 63))
    if (true && false) {
        if (false || false) {
            v3 = 32767
            v2 = (v2 / v2)
            var v23 string = "a"
            _ = v23
        } else {
            var v24 uint = 65535
            if (1 <= 522) {
                if ((!true) && false) {
                    v3 -= v3
                }
            } else {
                v1[674] = false
                if true {
                    var v25 int = 0
                    if (false && false) {
                        delete(v1, 2147483647)
                        v26 := []int{(v3 << (((v24 % ((0) | 1))) & 63)), 32768, v3, v3}
                        var v27 int64 = 0
                        _ = v26
                        _ = v27
                    } else {
                        v28 := GenS2{F0: -39}
                        v25 += (v3 % ((v25) | 1))
                        _ = v28
                    }
                    var v29 int8 = -19
                    _ = v25
                    _ = v29
                } else {
                    v30 := GenS3{F0: uint16(v24)}
                    switch -750 {
                    case 133:
                        var v31 int16 = -1
                        _ = v31
                    default:
                        if false {
                            for fi32 := 0; fi32 < 9; fi32++ {
                                v24 = (v24 % (((^(v24))) | 1))
                                _ = fi32
                            }
                        } else {
                            for fi33 := 0; fi33 < 4; fi33++ {
                                v34 := GenS4{F0: 48, F1: -414, F2: true}
                                for fi35 := 0; fi35 < 9; fi35++ {
                                    fmt.Println(v2, v2, v3)
                                    v36 := GenS2{F0: 127}
                                    var v37 int8 = 0
                                    _ = fi35
                                    _ = v36
                                    _ = v37
                                }
                                _ = fi33
                                _ = v34
                            }
                            fmt.Println(v2, v3, v2)
                            v3 -= (-(v3))
                        }
                    }
                    _ = v30
                }
            }
            if false {
                for fi38 := 0; fi38 < 4; fi38++ {
                    if true {
                        var v39 string = "\n\t"
                        v24 = v24
                        _ = v39
                    } else {
                        v40 := []int{fi38, fi38, 42}
                        _ = len(v40)
                        v41 := v1[2]
                        _ = v40
                        _ = v41
                    }
                    var v42 int16 = 1
                    v43 := []int{v3, fi38}
                    _ = fi38
                    _ = v42
                    _ = v43
                }
                var v44 uint8 = 53
                var v45 uint64 = 0
                _ = v44
                _ = v45
            } else {
                v46 := GenS5{F0: (!true), F1: 2}
                v47 := GenS3{F0: 0}
                if true {
                    v3 += v3
                } else {
                    switch ("a" + "hello") {
                    case "hello":
                        v2 = v2
                    case "":
                        var v48 uint = v24
                        _ = v48
                    case "foo bar":
                        fmt.Println(v3, v3, v3)
                    default:
                        for fi49 := 0; fi49 < 13; fi49++ {
                            v50 := v1[2147483647]
                            _ = fi49
                            _ = v50
                        }
                    }
                    for fi51 := 0; fi51 < 12; fi51++ {
                        if (false && true) {
                            v2 += 0.0
                        }
                        var v52 int = (fi51 >> ((0) & 63))
                        v2 += (((v2 * 1.5) + v2) - v2)
                        _ = fi51
                        _ = v52
                    }
                }
                _ = v46
                _ = v47
            }
            _ = v24
        }
    }
    for fi53 := 0; fi53 < 10; fi53++ {
        if true {
            var v54 float32 = (float32(-2.25) / ((v2 + (3.14159 / float32((1.5 / float64(-2.25))))) / v2))
            for fi55 := 0; fi55 < 8; fi55++ {
                if true {
                    if false {
                        _ = len(v1)
                        var v56 string = ""
                        switch 9223372036854775807 {
                        case 32767:
                            var v57 uint32 = 42
                            if (65535 >= 61) {
                                v58 := GenS6{F0: ("" + v56)}
                                _ = v58
                            } else {
                                _ = len(v1)
                                var v59 string = v56
                                fmt.Println(v54)
                                _ = v59
                            }
                            switch (("" + "hello") + v56) {
                            case "":
                                var v60 string = "\n\t"
                                _ = v60
                            }
                            _ = v57
                        case 2147483647:
                            if (127 <= 127) {
                                v61 := GenS2{F0: genFn1(1)}
                                _ = v61
                            }
                        case 783:
                            switch 1 {
                            case -2147483648:
                                if ((((-86 < 38) && true) || ((false && false) || (-1 <= -956))) || (!false)) {
                                    v62 := []int{(fi53 << ((868) & 63)), fi53}
                                    v63 := GenS7{F0: 1, F1: 1, F2: 255}
                                    v64 := GenS6{F0: v56}
                                    _ = v62
                                    _ = v63
                                    _ = v64
                                } else {
                                    v65 := GenS7{F0: 0, F1: 431, F2: 877}
                                    _ = v65
                                }
                                v56 = "\n\t"
                            case 1:
                                v66 := GenS8{F0: 42}
                                _ = v66
                            case -165:
                                v67 := GenS9{F0: -650}
                                _ = v67
                            }
                        default:
                            var v68 bool = true
                            _ = v68
                        }
                        var v69 bool = (true || true)
                        v70 := []int{}
                        var v71 int = (v3 | fi53)
                        _ = v56
                        _ = v69
                        _ = v70
                        _ = v71
                    }
                    var v72 int16 = 42
                    fmt.Println(v54, v72, fi55)
                    _ = v72
                }
                _ = fi55
            }
            _ = v54
        } else {
            switch (v3 - fi53) {
            case -9223372036854775808:
                var v73 uint16 = 65535
                _ = v73
            case 222:
                v74 := GenS10{F0: 1, F1: -1}
                _ = v74
            case 9223372036854775807:
                for fi75 := 0; fi75 < 6; fi75++ {
                    v76 := []int{635}
                    fmt.Println(v3, v3, v3)
                    v3 += 128
                    _ = fi75
                    _ = v76
                }
            }
            fmt.Println(v3)
            v3 -= fi53
        }
        _ = fi53
    }
    v3 -= (v3 | 32767)
    v77 := GenS3{F0: 2}
    var v78 string = (("hello" + "hello") + "a")
    if (!false) {
        if (!(2147483647 != -1)) {
            if false {
                v78 = ("hello" + v78)
                if true {
                    if false {
                        fmt.Println(v3, v78)
                    }
                    if true {
                        v79 := []int{-405, v3, v3, v3}
                        var v80 float32 = (v2 + (float32(3.14159) - v2))
                        _ = v79
                        _ = v80
                    } else {
                        fmt.Println(v3)
                        if false {
                            var v81 int32 = 0
                            _ = v81
                        } else {
                            v2 -= (v2 / (v2 + v2))
                            for fi82 := 0; fi82 < 6; fi82++ {
                                for fi83 := 0; fi83 < 8; fi83++ {
                                    v2 = v2
                                    var v84 uint16 = 65535
                                    if false {
                                        var v85 int16 = 0
                                        if (255 < 42) {
                                            v85 = v85
                                            v84 -= (v84 ^ (v84 & (^(v84))))
                                        }
                                        _ = v85
                                    } else {
                                        fmt.Println(v2, fi83)
                                        var v86 int32 = 42
                                        v87 := v1[4294967295]
                                        _ = v86
                                        _ = v87
                                    }
                                    _ = fi83
                                    _ = v84
                                }
                                if (!false) {
                                    v2 -= 0.0
                                } else {
                                    fmt.Println(v2)
                                }
                                v88 := v1[887]
                                _ = fi82
                                _ = v88
                            }
                            _ = len(v1)
                        }
                        if ((1e10 * (1e10 - (0.1 * (((3.14159 + 3.14159) + 1e10) - 1e10)))) <= ((((float64(((((v2 + ((v2 + 1e10) * float32((0.1 * (-2.25 + 1.5))))) - v2) - -2.25) + v2)) * float64((v2 * -2.25))) - (0.1 + 1.5)) * 0.1) * 1.5)) {
                            var v89 float64 = float64(v2)
                            var v90 string = v78
                            var v91 bool = true
                            _ = v89
                            _ = v90
                            _ = v91
                        }
                    }
                }
            } else {
                v92 := GenS8{F0: 0}
                v78 = "\n\t"
                _ = v92
            }
        } else {
            if (!(v3 > (v3 * 32768))) {
                var v93 uint8 = 127
                switch v78 {
                case "foo bar":
                    v2 -= v2
                case "":
                    for fi94 := 0; fi94 < 15; fi94++ {
                        var v95 int = fi94
                        var v96 bool = true
                        var v97 uint64 = 543
                        _ = fi94
                        _ = v95
                        _ = v96
                        _ = v97
                    }
                }
                var v98 uint8 = v93
                _ = v93
                _ = v98
            } else {
                v3 = v3
                if false {
                    for fi99 := 0; fi99 < 15; fi99++ {
                        var v100 int = v3
                        v3 = (-(v3))
                        _ = fi99
                        _ = v100
                    }
                    v1[1] = true
                    var v101 uint8 = genFn0(((v3 * v3) != v3), 4294967295)
                    _ = v101
                } else {
                    if (-1 >= 42) {
                        v102 := GenS11{F0: false, F1: 2147483648, F2: -562}
                        var v103 int = 895
                        v103 = v103
                        _ = v102
                        _ = v103
                    }
                    switch v78 {
                    case "foo bar":
                        v2 = v2
                    case "\n\t":
                        var v104 uint32 = 174
                        _ = v104
                    case "hello":
                        var v105 uint8 = 127
                        _ = v105
                    }
                    for fi106 := 0; fi106 < 2; fi106++ {
                        if (true && (0 > 1)) {
                            v78 = "a"
                            var v107 int16 = 0
                            for fi108 := 0; fi108 < 19; fi108++ {
                                fmt.Println(v2, fi108, v107)
                                _ = len(v1)
                                if (4294967295 != 65535) {
                                    v109 := GenS12{F0: 18446744073709551615, F1: 480, F2: v107}
                                    var v110 uint = 948
                                    _ = v109
                                    _ = v110
                                }
                                v2 = float32(3.14159)
                                v111 := GenS11{F0: (235 > genFn0(((v78 + "") < ("hello" + v78)), 2)), F1: 4294967295, F2: (v107 | 1000)}
                                _ = fi108
                                _ = v111
                            }
                            _ = v107
                        } else {
                            fmt.Println(v2)
                            for fi112 := 0; fi112 < 13; fi112++ {
                                fmt.Println(v78, v2)
                                v113 := []int{fi112, fi112, (fi112 + 4294967295)}
                                v114 := GenS12{F0: 18446744073709551615, F1: 410, F2: -978}
                                _ = fi112
                                _ = v113
                                _ = v114
                            }
                            fmt.Println(v2, v2, v3)
                        }
                        v2 += v2
                        _ = fi106
                    }
                }
                for fi115 := 0; fi115 < 1; fi115++ {
                    if (!(1e10 > 3.14159)) {
                        v116 := GenS8{F0: 0}
                        var v117 string = (v78 + v78)
                        v118 := v1[1]
                        _ = v116
                        _ = v117
                        _ = v118
                    } else {
                        for fi119 := 0; fi119 < 9; fi119++ {
                            fmt.Println(fi115, v3, v2)
                            if (!((!false) && false)) {
                                v78 = v78
                                v120 := []int{(v3 * 0), (^(fi119)), fi119}
                                v121 := GenS13{F0: true}
                                v122 := GenS13{F0: false}
                                v2 = v2
                                _ = v120
                                _ = v121
                                _ = v122
                            }
                            if false {
                                var v123 string = (v78 + (v78 + ""))
                                var v124 string = (((v78 + (v123 + ((v78 + ("foo bar" + "a")) + "hello"))) + v78) + ("hello" + ("a" + v123)))
                                var v125 uint64 = 711
                                _ = v123
                                _ = v124
                                _ = v125
                            }
                            _ = fi119
                        }
                        if (1e10 == 0.1) {
                            v2 -= 0.0
                        }
                    }
                    _ = fi115
                }
            }
        }
        if true {
            for fi126 := 0; fi126 < 9; fi126++ {
                if true {
                    v2 = v2
                    v127 := GenS10{F0: 255, F1: 1}
                    _ = v127
                } else {
                    v128 := GenS3{F0: 2}
                    switch v78 {
                    case "\n\t":
                        v2 = -2.25
                    case "hello":
                        switch fi126 {
                        case -193:
                            var v129 bool = true
                            _ = v129
                        case 32767:
                            var v130 int8 = 29
                            _ = v130
                        case 9223372036854775807:
                            for fi131 := 0; fi131 < 3; fi131++ {
                                v1[2147483647] = (42 != -2147483648)
                                if false {
                                    var v132 int = (v3 % (((^(fi126))) | 1))
                                    v78 = v78
                                    v133 := []int{(v132 << ((18446744073709551615) & 63)), (v3 / ((v3) | 1)), fi126, 42}
                                    _ = v132
                                    _ = v133
                                }
                                switch v3 {
                                case 1:
                                    if false {
                                        v78 = v78
                                    } else {
                                        var v134 float64 = 1e10
                                        _ = v134
                                    }
                                case -2147483648:
                                    var v135 int = fi131
                                    _ = v135
                                case -908:
                                    v136 := []int{(fi126 | fi131), (v3 + -1), (fi126 - v3), (v3 / (((fi131 >> ((281) & 63))) | 1))}
                                    _ = v136
                                }
                                _ = fi131
                            }
                        default:
                            for fi137 := 0; fi137 < 1; fi137++ {
                                _ = len(v1)
                                fmt.Println(fi126)
                                _ = fi137
                            }
                        }
                    default:
                        v138 := GenS14{F0: 702}
                        _ = v138
                    }
                    v78 = ("foo bar" + v78)
                    _ = v128
                }
                for fi139 := 0; fi139 < 12; fi139++ {
                    if true {
                        switch "hello" {
                        case "foo bar":
                            v140 := GenS9{F0: 9223372036854775807}
                            v78 = ((v78 + (v78 + "hello")) + "a")
                            switch "" {
                            case "hello":
                                v78 = ((((("" + (((v78 + v78) + "\n\t") + v78)) + "") + (v78 + "")) + v78) + "foo bar")
                            case "":
                                var v141 uint16 = 510
                                _ = v141
                            case "foo bar":
                                var v142 int64 = -9223372036854775808
                                _ = v142
                            default:
                                v143 := []int{}
                                _ = v143
                            }
                            _ = v140
                        case "a":
                            var v144 uint = 490
                            _ = v144
                        default:
                            switch -1 {
                            case 9223372036854775807:
                                if false {
                                    v145 := []int{(v3 & v3)}
                                    if true {
                                        switch fi139 {
                                        case 861:
                                            var v146 int16 = 152
                                            _ = v146
                                        case 769:
                                            if len(v145) > 0 {
                                                v145[0] = fi139
                                            }
                                        }
                                    } else {
                                        var v147 uint8 = 0
                                        switch "" {
                                        case "\n\t":
                                            if true {
                                                var v148 uint32 = 190
                                                if (0.1 > ((3.14159 - 0.0) + 1e10)) {
                                                    v149 := v1[(v148 / ((v148) | 1))]
                                                    _ = v149
                                                } else {
                                                    v150 := v1[(v148 << ((2) & 63))]
                                                    v151 := GenS6{F0: v78}
                                                    v150 = v150
                                                    _ = v150
                                                    _ = v151
                                                }
                                                var v152 uint = 18446744073709551615
                                                _ = v148
                                                _ = v152
                                            }
                                        case "":
                                            v2 += float32((1e10 / 0.1))
                                        default:
                                            if (!true) {
                                                var v153 int32 = 2147483647
                                                fmt.Println(v153, v3)
                                                _ = v153
                                            }
                                        }
                                        _ = v147
                                    }
                                    _ = v145
                                }
                            }
                        }
                        var v154 uint = 0
                        var v155 uint16 = 255
                        _ = v154
                        _ = v155
                    } else {
                        var v156 float64 = (0.1 - 0.0)
                        _ = v156
                    }
                    if (false && (128 < 232)) {
                        fmt.Println(v78)
                        v157 := GenS15{F0: 127, F1: 10}
                        v2 -= v2
                        _ = v157
                    } else {
                        var v158 bool = (32768 > fi139)
                        switch v78 {
                        case "hello":
                            v3 = -9223372036854775808
                        case "\n\t":
                            v3 += 4294967295
                        default:
                            v159 := []int{}
                            _ = v159
                        }
                        if ((0 >= 2) && v158) {
                            var v160 int = fi126
                            fmt.Println(v3, v78, v160)
                            switch v78 {
                            case "":
                                switch v3 {
                                case 9223372036854775807:
                                    v161 := []int{(v3 >> ((48) & 63)), (fi126 >> ((2) & 63))}
                                    v162 := GenS15{F0: 55, F1: 178}
                                    v2 -= v2
                                    _ = v161
                                    _ = v162
                                default:
                                    if v158 {
                                        v2 -= float32(1.5)
                                        v160 += (fi139 + (^(v160)))
                                    } else {
                                        v78 = ((v78 + (v78 + "")) + "")
                                        var v163 int8 = -128
                                        _ = v163
                                    }
                                }
                            case "foo bar":
                                v164 := GenS16{F0: 533, F1: (!v158)}
                                v160 -= fi126
                                var v165 int32 = -92
                                var v166 uint16 = 1
                                _ = v164
                                _ = v165
                                _ = v166
                            case "hello":
                                if (!v158) {
                                    delete(v1, 2147483647)
                                    switch "hello" {
                                    case "hello":
                                        var v167 bool = ((!v158) || v158)
                                        _ = v167
                                    case "":
                                        v2 += (v2 + float32(((1.5 / ((0.0 - 1.5) + 3.14159)) / (-2.25 + 0.0))))
                                    }
                                }
                            }
                            v160 = 32768
                            _ = v160
                        }
                        _ = v158
                    }
                    var v168 float32 = float32((-2.25 - 1e10))
                    switch fi139 {
                    case -9223372036854775808:
                        v2 -= (v2 - -2.25)
                    }
                    switch "\n\t" {
                    case "foo bar":
                        v78 = ("a" + ("hello" + ("hello" + v78)))
                    case "\n\t":
                        v169 := GenS11{F0: false, F1: 2, F2: 32767}
                        _ = v169
                    case "hello":
                        fmt.Println(fi139, v2, v168)
                    }
                    _ = fi139
                    _ = v168
                }
                var v170 int8 = 115
                _ = fi126
                _ = v170
            }
        }
        v171 := GenS16{F0: 98, F1: (!(v78 > ((("a" + "\n\t") + v78) + ("hello" + (v78 + "foo bar")))))}
        _ = v171
    }
    if true {
        v78 = (v78 + (v78 + "foo bar"))
        var v172 uint16 = 65535
        if (!false) {
            fmt.Println(v78)
            var v173 int = (v3 & (v3 << ((998) & 63)))
            var v174 float32 = ((v2 - (((3.14159 / -2.25) + v2) - v2)) + (3.14159 - v2))
            _ = v173
            _ = v174
        }
        _ = v172
    } else {
        v2 -= ((v2 * (v2 / 1.5)) / 1.5)
        v175 := GenS15{F0: 42, F1: 1}
        var v176 uint = 42
        _ = v175
        _ = v176
    }
    for fi177 := 0; fi177 < 12; fi177++ {
        var v178 int64 = 0
        for fi179 := 0; fi179 < 19; fi179++ {
            switch (fi177 >> ((415) & 63)) {
            case 379:
                if true {
                    var v180 string = "foo bar"
                    _ = v180
                } else {
                    switch v78 {
                    case "\n\t":
                        switch (v78 + (((v78 + v78) + v78) + "foo bar")) {
                        case "a":
                            switch 4294967295 {
                            case 1:
                                fmt.Println(fi179)
                            case -1:
                                v181 := GenS3{F0: 2}
                                _ = v181
                            case 32768:
                                var v182 uint = 885
                                _ = v182
                            default:
                                var v183 int16 = 1000
                                _ = v183
                            }
                        case "foo bar":
                            switch fi179 {
                            case 127:
                                v2 = v2
                            case 0:
                                switch "\n\t" {
                                case "foo bar":
                                    if true {
                                        var v184 uint32 = 1
                                        var v185 int8 = 42
                                        _ = v184
                                        _ = v185
                                    } else {
                                        v3 = int(v178)
                                        switch fi177 {
                                        case 32768:
                                            var v186 float64 = 1e10
                                            _ = v186
                                        case 42:
                                            v3 -= fi179
                                        case 673:
                                            var v187 int16 = 32767
                                            _ = v187
                                        }
                                    }
                                case "a":
                                    v188 := GenS14{F0: 42}
                                    _ = v188
                                case "":
                                    var v189 bool = (-1 != fi177)
                                    _ = v189
                                }
                            case -9223372036854775808:
                                if (!true) {
                                    v190 := GenS7{F0: 678, F1: 2147483647, F2: 0}
                                    v3 -= fi177
                                    fmt.Println(fi177)
                                    v191 := []int{1, (-(fi177)), (fi179 + fi179)}
                                    _ = v190
                                    _ = v191
                                }
                            default:
                                var v192 float32 = v2
                                _ = v192
                            }
                        case "hello":
                            if (1 <= 4294967295) {
                                switch v3 {
                                case 2147483647:
                                    v3 = (-(v3))
                                case -1:
                                    var v193 uint = 1
                                    _ = v193
                                default:
                                    v178 = -1
                                }
                                v194 := GenS17{F0: 458, F1: 358}
                                v1[500] = (false || (100 == 127))
                                var v195 uint32 = 162
                                v78 = ((v78 + ("hello" + v78)) + "foo bar")
                                _ = v194
                                _ = v195
                            } else {
                                var v196 int32 = -667
                                if false {
                                    var v197 bool = false
                                    if (0 <= v196) {
                                        if v197 {
                                            v196 -= 42
                                        } else {
                                            v3 += fi177
                                            if v197 {
                                                var v198 uint8 = genFn0(v197, 0)
                                                _ = v198
                                            }
                                        }
                                    }
                                    _ = v197
                                }
                                _ = v196
                            }
                            v199 := GenS18{F0: 732, F1: genFn1(1)}
                            _ = v199
                        default:
                            if (!true) {
                                var v200 uint8 = 0
                                if false {
                                    v201 := GenS17{F0: 2, F1: 4294967295}
                                    var v202 string = v78
                                    _ = v201
                                    _ = v202
                                }
                                _ = v200
                            } else {
                                if false {
                                    v3 += (fi177 * -365)
                                    if (0 <= 42) {
                                        v203 := []int{(-(v3))}
                                        _ = v203
                                    } else {
                                        switch "" {
                                        case "\n\t":
                                            var v204 uint16 = 0
                                            _ = v204
                                        case "":
                                            var v205 uint32 = 724
                                            _ = v205
                                        default:
                                            var v206 uint16 = 255
                                            _ = v206
                                        }
                                        fmt.Println(v178, fi179, fi177)
                                        var v207 int8 = -51
                                        _ = v207
                                    }
                                    if true {
                                        _ = len(v1)
                                        fmt.Println(v3, v78)
                                        var v208 float32 = v2
                                        _ = v208
                                    } else {
                                        v209 := GenS2{F0: 100}
                                        _ = len(v1)
                                        _ = v209
                                    }
                                    if (((v2 == ((v2 * float32(0.1)) + (0.0 * (v2 / v2)))) && (((v178 / (((v178 / (((v178 & -1)) | 1))) | 1)) == (v178 & v178)) || true)) && (!false)) {
                                        v78 = (v78 + v78)
                                        v210 := GenS1{F0: 833, F1: -1}
                                        fmt.Println(v3)
                                        _ = v210
                                    }
                                }
                                v1[608] = (v78 >= ("foo bar" + ""))
                            }
                        }
                    case "a":
                        var v211 float32 = 0.0
                        _ = v211
                    }
                    var v212 uint32 = 4294967295
                    switch v78 {
                    case "":
                        v213 := GenS18{F0: v212, F1: -128}
                        _ = v213
                    case "\n\t":
                        v214 := []int{}
                        _ = v214
                    }
                    _ = v212
                }
            case 2147483647:
                fmt.Println(v3, v3, fi179)
            default:
                var v215 uint = 255
                _ = v215
            }
            _ = fi179
        }
        v78 = "\n\t"
        _ = fi177
        _ = v178
    }
    for fi216 := 0; fi216 < 9; fi216++ {
        if true {
            v217 := GenS10{F0: genFn0(((0 > 2) && (!((v2 + v2) > float32(-2.25)))), 1), F1: -1}
            if true {
                fmt.Println(v3, fi216)
            }
            _ = v217
        } else {
            var v218 string = ("" + "hello")
            for fi219 := 0; fi219 < 11; fi219++ {
                if (91 >= 127) {
                    switch fi216 {
                    case 2147483647:
                        if true {
                            fmt.Println(v78, v78)
                        } else {
                            switch fi219 {
                            case -386:
                                v220 := []int{}
                                _ = v220
                            case -142:
                                if ((true || false) || true) {
                                    v3 -= (^(v3))
                                }
                            case -273:
                                var v221 int32 = 1
                                _ = v221
                            }
                            var v222 uint = 65535
                            fmt.Println(v222)
                            _ = v222
                        }
                    default:
                        if (127 == genFn0((true && true), 476)) {
                            _ = len(v1)
                            switch (fi216 << ((1) & 63)) {
                            case -249:
                                v2 += ((v2 / 1e10) * v2)
                            case 42:
                                switch v78 {
                                case "a":
                                    if (false && (false || ((true || false) && false))) {
                                        v223 := GenS3{F0: 2}
                                        v224 := GenS19{F0: 32768, F1: 0, F2: 1}
                                        _ = v223
                                        _ = v224
                                    } else {
                                        if true {
                                            var v225 float32 = 1e10
                                            _ = v225
                                        } else {
                                            v226 := []int{(v3 ^ 0)}
                                            fmt.Println(v218, v2)
                                            v227 := GenS20{F0: fi219, F1: fi216, F2: v218}
                                            _ = v226
                                            _ = v227
                                        }
                                    }
                                case "foo bar":
                                    v2 += (v2 * v2)
                                case "hello":
                                    v228 := GenS21{F0: 42}
                                    _ = v228
                                default:
                                    if (!true) {
                                        fmt.Println(v2, fi219, fi216)
                                    } else {
                                        switch v3 {
                                        case 4294967295:
                                            if false {
                                                delete(v1, 2147483648)
                                                v3 = fi219
                                                var v229 float32 = ((float32((0.0 - (3.14159 / 0.1))) / 1.0) - v2)
                                                _ = v229
                                            } else {
                                                switch v218 {
                                                case "":
                                                    var v230 string = v78
                                                    _ = v230
                                                }
                                                fmt.Println(v3)
                                                v231 := GenS2{F0: 0}
                                                _ = v231
                                            }
                                        case -210:
                                            v78 = v78
                                        }
                                        v1[42] = false
                                        switch v78 {
                                        case "\n\t":
                                            var v232 string = "a"
                                            _ = v232
                                        case "a":
                                            if true {
                                                var v233 string = "a"
                                                v234 := GenS1{F0: 82, F1: 2147483647}
                                                var v235 int8 = 100
                                                _ = v233
                                                _ = v234
                                                _ = v235
                                            }
                                        case "hello":
                                            var v236 int = (-(v3))
                                            _ = v236
                                        }
                                        var v237 string = (v78 + v218)
                                        _ = v237
                                    }
                                    if (true || false) {
                                        fmt.Println(v78, v3, v2)
                                    }
                                }
                            default:
                                if (42 != 1) {
                                    fmt.Println(fi216, fi216)
                                }
                            }
                            switch v3 {
                            case 4294967295:
                                fmt.Println(v218, v2, v218)
                                var v238 uint8 = 0
                                v78 = v78
                                _ = v238
                            }
                        }
                    }
                } else {
                    v218 = ((v78 + v218) + v218)
                }
                v239 := GenS5{F0: (142 != 2), F1: 838}
                _ = fi219
                _ = v239
            }
            var v240 uint = 255
            _ = v218
            _ = v240
        }
        v241 := []int{v3}
        delete(v1, 2147483647)
        _ = fi216
        _ = v241
    }
    var v242 bool = (0 <= 0)
    if (genFn0(true, 18446744073709551615) <= 1) {
        switch (v3 ^ (-(v3))) {
        case -2147483648:
            v242 = true
        case 127:
            var v243 bool = v242
            _ = v243
        case 32767:
            if (!(1 >= -389)) {
                v244 := GenS5{F0: false, F1: 331}
                _ = v244
            }
        }
        if (128 > 128) {
            v245 := GenS22{F0: 4294967295, F1: "a", F2: v242}
            _ = v245
        } else {
            delete(v1, 2)
            switch "a" {
            case "foo bar":
                v2 -= 0.0
            case "hello":
                fmt.Println(v78)
            case "":
                switch (v3 << ((42) & 63)) {
                case 1:
                    fmt.Println(v3)
                case 32768:
                    v3 = (v3 - (v3 % ((v3) | 1)))
                }
            }
        }
    }
    var v246 int8 = genFn1(-1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v242:", v242)
    fmt.Println("v246:", v246)
    fmt.Println("v3:", v3)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v2
    _ = v242
    _ = v246
    _ = v3
    _ = v77
    _ = v78
}
