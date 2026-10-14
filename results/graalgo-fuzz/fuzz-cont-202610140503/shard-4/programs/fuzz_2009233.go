package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 uint32; F2 uint8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int; F1 uint8; F2 int64 }
type GenS5 struct { F0 bool; F1 uint8; F2 uint64 }
type GenS6 struct { F0 string; F1 bool; F2 string }
type GenS7 struct { F0 uint16; F1 uint }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 uint; F1 uint16 }
type GenS14 struct { F0 int64; F1 uint8; F2 bool }
type GenS15 struct { F0 string; F1 int64 }
type GenS16 struct { F0 int32; F1 int; F2 string }
func genFn0(p0 int16, p1 uint32) uint32 {
    if (!false) {
        fmt.Println(p1, p1)
    } else {
        fmt.Println(p0, p0)
        fmt.Println(p0, p1)
        var v6 uint32 = p1
        _ = v6
    }
    if (42 >= 100) {
        for fi7 := 0; fi7 < 9; fi7++ {
            v8 := GenS1{F0: 128}
            var v9 uint64 = 803
            var v10 int = fi7
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
        for fi11 := 0; fi11 < 19; fi11++ {
            switch (^(fi11)) {
            case -190:
                p0 += (p0 % (((p0 + p0)) | 1))
            case 32767:
                fmt.Println(fi11)
            }
            switch ((("a" + "") + "\n\t") + "foo bar") {
            case "foo bar":
                p1 -= 847
            default:
                fmt.Println(fi11, p1, p0)
            }
            _ = fi11
        }
        var v12 int16 = p0
        _ = v12
    }
    v13 := []int{-2147483648, 32767, -9223372036854775808}
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    for fi14 := 0; fi14 < 3; fi14++ {
        v15 := GenS2{F0: -51, F1: (p1 << ((1) & 63)), F2: 1}
        _ = fi14
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v13
    return p1
}
func genFn1(p0 int8) int {
    v16 := []int{-332, 605, -966}
    p0 -= (p0 >> ((0) & 63))
    switch 42 {
    case 0:
        for fi17 := 0; fi17 < 16; fi17++ {
            var v18 int64 = -1
            if true {
                fmt.Println(v18, p0)
                var v19 uint8 = 255
                _ = v19
            } else {
                v20 := GenS3{F0: fi17}
                _ = v20
            }
            switch ("a" + ("" + "a")) {
            case "hello":
                if len(v16) > 0 {
                    v16[0] = int(v18)
                }
            case "\n\t":
                var v21 int8 = -103
                _ = v21
            case "":
                p0 += 75
            default:
                v22 := GenS4{F0: fi17, F1: 42, F2: v18}
                _ = v22
            }
            _ = fi17
            _ = v18
        }
    }
    var v23 int32 = 0
    _ = p0
    _ = v16
    _ = v23
    return 566
}

func main() {
    var v1 int16 = 1
    var v2 uint32 = 65535
    var v3 float32 = 1.5
    var v4 float32 = 3.14159
    var v5 string = ("a" + "hello")
    var v24 uint = 18446744073709551615
    var v25 bool = true
    for fi26 := 0; fi26 < 19; fi26++ {
        if v25 {
            if true {
                for fi27 := 0; fi27 < 16; fi27++ {
                    v3 += v4
                    v25 = (!false)
                    fmt.Println(v24)
                    _ = fi27
                }
                v28 := GenS5{F0: (v25 && v25), F1: 42, F2: 4294967295}
                _ = v28
            }
        }
        v29 := []int{(fi26 % ((fi26) | 1)), fi26}
        _ = fi26
        _ = v29
    }
    var v30 int8 = 100
    v2 = (v2 | (v2 - v2))
    if false {
        switch -765 {
        case 369:
            for fi31 := 0; fi31 < 2; fi31++ {
                if (v25 && false) {
                    for fi32 := 0; fi32 < 14; fi32++ {
                        v33 := GenS6{F0: v5, F1: v25, F2: v5}
                        switch "foo bar" {
                        case "\n\t":
                            v1 -= 1000
                        case "foo bar":
                            fmt.Println(fi32, v5)
                        case "hello":
                            var v34 uint8 = 42
                            _ = v34
                        }
                        _ = fi32
                        _ = v33
                    }
                    if (0.0 == v4) {
                        v35 := []int{genFn1(v30), fi31, (fi31 / (((fi31 << ((v24) & 63))) | 1)), (fi31 - (fi31 | fi31))}
                        _ = v35
                    }
                    if (genFn1((v30 - (v30 & (v30 << (((v24 + v24)) & 63))))) != fi31) {
                        var v36 int32 = -1
                        switch ("a" + (v5 + (("a" + v5) + v5))) {
                        case "hello":
                            var v37 uint = 97
                            _ = v37
                        case "":
                            var v38 uint = (v24 >> ((0) & 63))
                            _ = v38
                        case "\n\t":
                            var v39 int8 = v30
                            _ = v39
                        default:
                            switch fi31 {
                            case 128:
                                v24 = (v24 + (v24 % ((v24) | 1)))
                            case -792:
                                v40 := GenS5{F0: v25, F1: 0, F2: 105}
                                _ = v40
                            }
                        }
                        if ((float32(0.1) / (v3 + v4)) > float32(1.5)) {
                            var v41 uint8 = 0
                            if v25 {
                                var v42 string = v5
                                for fi43 := 0; fi43 < 17; fi43++ {
                                    v44 := []int{(fi43 ^ (fi31 >> ((v24) & 63))), fi31}
                                    _ = fi43
                                    _ = v44
                                }
                                _ = v42
                            }
                            var v45 float64 = (((1e10 * 1e10) - 1e10) / (((((3.14159 + -2.25) * -2.25) - 1e10) / 3.14159) + 1.5))
                            _ = v41
                            _ = v45
                        } else {
                            var v46 float32 = v3
                            if v25 {
                                v47 := []int{fi31, genFn1(v30), (fi31 * (^(fi31)))}
                                var v48 uint64 = 605
                                if v25 {
                                    v47 = append(v47, (fi31 << ((18446744073709551615) & 63)))
                                } else {
                                    if (127 != v30) {
                                        var v49 int = 42
                                        if len(v47) >= 2 {
                                            v47 = v47[1:len(v47)]
                                        }
                                        if len(v47) > 0 {
                                            v47[0] = int(v48)
                                        }
                                        _ = v49
                                    }
                                    var v50 uint64 = (v48 >> (((v24 * (v24 - (^(v24))))) & 63))
                                    _ = v50
                                }
                                _ = v47
                                _ = v48
                            } else {
                                v51 := []int{(fi31 + (fi31 / ((fi31) | 1))), fi31, fi31, 554}
                                for fi52 := 0; fi52 < 2; fi52++ {
                                    fmt.Println(v1, v36, v46)
                                    v5 = v5
                                    if v25 {
                                        v25 = (v46 > float32(-2.25))
                                    }
                                    var v53 float64 = (0.1 * 0.0)
                                    if (251 != 2) {
                                        if v25 {
                                            v53 = v53
                                            var v54 bool = (0 > 9223372036854775807)
                                            fmt.Println(v54, v25, v30)
                                            _ = v54
                                        }
                                    } else {
                                        v36 = v36
                                    }
                                    _ = len(v51)
                                    fmt.Println(v4, v24, fi31)
                                    _ = fi52
                                    _ = v53
                                }
                                _ = v51
                            }
                            _ = v46
                        }
                        _ = v36
                    }
                } else {
                    for fi55 := 0; fi55 < 17; fi55++ {
                        if ((float32(1e10) <= v4) || (v25 && ((!v25) || (v25 || v25)))) {
                            var v56 uint = v24
                            var v57 uint32 = v2
                            _ = v56
                            _ = v57
                        }
                        fmt.Println(v24, v30)
                        fmt.Println(v30)
                        _ = fi55
                    }
                    switch 32768 {
                    case 9223372036854775807:
                        for fi58 := 0; fi58 < 13; fi58++ {
                            v4 = v4
                            _ = fi58
                        }
                    case -929:
                        if (!true) {
                            v30 = (v30 >> ((42) & 63))
                            var v59 float32 = 1e10
                            v60 := []int{}
                            _ = v59
                            _ = v60
                        } else {
                            v30 = (v30 / ((0) | 1))
                        }
                    }
                }
                v4 = ((1.5 * (v4 - v3)) - float32(((float64((3.14159 / v3)) + ((-2.25 / (1.5 - (float64(v3) - (3.14159 + 1e10)))) + 1.5)) * float64((float32(3.14159) - 1e10)))))
                var v61 bool = v25
                _ = fi31
                _ = v61
            }
        case -1:
            v62 := GenS7{F0: 1, F1: v24}
            _ = v62
        case -106:
            var v63 uint = 42
            _ = v63
        default:
            fmt.Println(v5, v25, v4)
        }
    }
    var v64 int32 = 42
    var v65 uint8 = 230
    if (genFn0(-80, genFn0(-32768, v2)) > genFn0(42, (v2 % (((v2 * genFn0((^(v1)), 501))) | 1)))) {
        for fi66 := 0; fi66 < 16; fi66++ {
            var v67 int = -9223372036854775808
            switch "" {
            case "":
                v1 += (-(v1))
            case "hello":
                fmt.Println(v1)
            case "a":
                for fi68 := 0; fi68 < 2; fi68++ {
                    v25 = v25
                    var v69 float32 = v4
                    _ = fi68
                    _ = v69
                }
            default:
                for fi70 := 0; fi70 < 14; fi70++ {
                    v71 := GenS8{F0: genFn1(v30)}
                    fmt.Println(v2)
                    _ = fi70
                    _ = v71
                }
            }
            v72 := GenS9{F0: -9223372036854775808}
            _ = fi66
            _ = v67
            _ = v72
        }
        var v73 bool = true
        _ = v73
    } else {
        for fi74 := 0; fi74 < 17; fi74++ {
            v5 = v5
            for fi75 := 0; fi75 < 17; fi75++ {
                v76 := GenS10{F0: fi74}
                v77 := []int{int(v1), fi75}
                if v25 {
                    v78 := GenS9{F0: 534}
                    var v79 uint = (v24 % (((v24 * v24)) | 1))
                    var v80 int32 = (v64 ^ v64)
                    _ = v78
                    _ = v79
                    _ = v80
                } else {
                    if v25 {
                        var v81 bool = v25
                        _ = v81
                    } else {
                        v65 += (^(v65))
                        switch fi75 {
                        case 128:
                            var v82 bool = (!true)
                            _ = v82
                        case -1:
                            _ = len(v77)
                        case 9223372036854775807:
                            fmt.Println(v2, v64, v24)
                        default:
                            fmt.Println(v4, fi74, v25)
                        }
                        if v25 {
                            var v83 float32 = v3
                            var v84 uint32 = 51
                            var v85 int = (fi74 & fi75)
                            _ = v83
                            _ = v84
                            _ = v85
                        } else {
                            fmt.Println(fi74, v64)
                            var v86 string = v5
                            _ = v86
                        }
                    }
                    v5 = v5
                }
                _ = fi75
                _ = v76
                _ = v77
            }
            v24 -= v24
            _ = fi74
        }
        switch "foo bar" {
        case "a":
            var v87 int16 = v1
            _ = v87
        case "hello":
            if ((v5 != ((v5 + "hello") + v5)) && (v5 < (v5 + ("hello" + ("" + (v5 + (v5 + v5))))))) {
                fmt.Println(v24)
                var v88 int64 = 1
                _ = v88
            } else {
                for fi89 := 0; fi89 < 12; fi89++ {
                    if v25 {
                        v1 += 898
                        switch genFn1((v30 % ((-128) | 1))) {
                        case 2147483647:
                            for fi90 := 0; fi90 < 13; fi90++ {
                                v25 = (false || true)
                                var v91 uint = v24
                                _ = fi90
                                _ = v91
                            }
                        case 9223372036854775807:
                            for fi92 := 0; fi92 < 8; fi92++ {
                                switch (v5 + (v5 + v5)) {
                                case "\n\t":
                                    v93 := []int{fi92, fi89}
                                    _ = v93
                                case "foo bar":
                                    switch v5 {
                                    case "hello":
                                        fmt.Println(v25, v2, v65)
                                    case "foo bar":
                                        var v94 uint16 = 2
                                        _ = v94
                                    default:
                                        v3 -= 0.1
                                    }
                                case "a":
                                    fmt.Println(v30, v3)
                                default:
                                    switch (fi92 | fi89) {
                                    case 9223372036854775807:
                                        if ((v2 - (v2 - v2)) < v2) {
                                            var v95 bool = ((v24 - (^(v24))) != v24)
                                            _ = v95
                                        } else {
                                            var v96 int8 = v30
                                            v2 += (v2 | (v2 << ((v24) & 63)))
                                            _ = v96
                                        }
                                    default:
                                        fmt.Println(v65, fi92)
                                        v97 := GenS9{F0: 9223372036854775807}
                                        v98 := []int{fi89, fi89, (fi92 / ((fi89) | 1)), fi92}
                                        _ = v97
                                        _ = v98
                                    }
                                }
                                var v99 int64 = 0
                                v1 -= (v1 | v1)
                                v100 := GenS11{F0: 54}
                                v1 += v1
                                var v101 int8 = int8(v24)
                                v102 := GenS6{F0: ("a" + v5), F1: v25, F2: "\n\t"}
                                switch (fi92 & (fi92 * fi89)) {
                                case 42:
                                    fmt.Println(v25, v64)
                                case 32767:
                                    var v103 float64 = (float64(0.0) / ((0.0 + 0.0) + 0.1))
                                    _ = v103
                                case -556:
                                    v104 := GenS3{F0: (fi89 * (fi89 >> (((v24 - (v24 / ((v24) | 1)))) & 63)))}
                                    _ = v104
                                }
                                fmt.Println(fi89, v5)
                                var v105 int64 = -569
                                _ = fi92
                                _ = v100
                                _ = v101
                                _ = v102
                                _ = v105
                                _ = v99
                            }
                        }
                        v106 := GenS10{F0: fi89}
                        _ = v106
                    }
                    for fi107 := 0; fi107 < 20; fi107++ {
                        var v108 uint64 = 42
                        fmt.Println(v4, fi107)
                        var v109 float32 = 0.1
                        _ = fi107
                        _ = v108
                        _ = v109
                    }
                    _ = fi89
                }
            }
        }
        var v110 int32 = 0
        _ = v110
    }
    var v111 uint32 = (v2 >> ((v24) & 63))
    v111 = 1
    switch 32767 {
    case 0:
        if (!true) {
            if (uint(v65) < (^(v24))) {
                var v112 string = ("a" + "foo bar")
                if (-400 != (^(v64))) {
                    v113 := GenS12{F0: (v2 * v111)}
                    _ = v113
                } else {
                    fmt.Println(v112)
                    if false {
                        var v114 uint8 = (^(v65))
                        if v25 {
                            switch 1 {
                            case 127:
                                switch "hello" {
                                case "\n\t":
                                    v115 := []int{int(v24), 22, 2147483647}
                                    _ = v115
                                case "foo bar":
                                    v1 += -32768
                                }
                            }
                            if (!false) {
                                for fi116 := 0; fi116 < 8; fi116++ {
                                    var v117 uint16 = 42
                                    _ = fi116
                                    _ = v117
                                }
                            } else {
                                v118 := []int{-1, genFn1((v30 + -1)), -4, 0}
                                var v119 int64 = -9223372036854775808
                                switch 1 {
                                case 0:
                                    fmt.Println(v119, v114)
                                case 127:
                                    var v120 bool = false
                                    _ = v120
                                }
                                _ = v118
                                _ = v119
                            }
                        } else {
                            v121 := []int{-1, 128, genFn1(-90), 427}
                            _ = v121
                        }
                        var v122 float64 = 3.14159
                        _ = v114
                        _ = v122
                    } else {
                        var v123 int64 = int64(v30)
                        for fi124 := 0; fi124 < 16; fi124++ {
                            if v25 {
                                v125 := []int{(fi124 & (^(fi124))), (^(fi124)), genFn1(int8(fi124))}
                                var v126 int32 = v64
                                for fi127 := 0; fi127 < 4; fi127++ {
                                    var v128 int8 = v30
                                    v65 -= v65
                                    _ = fi127
                                    _ = v128
                                }
                                _ = v125
                                _ = v126
                            } else {
                                v129 := GenS13{F0: v24, F1: 255}
                                _ = v129
                            }
                            _ = fi124
                        }
                        if v25 {
                            fmt.Println(v1, v4)
                        } else {
                            var v130 uint = (^(v24))
                            v131 := GenS14{F0: 42, F1: v65, F2: false}
                            var v132 int8 = int8(v65)
                            _ = v130
                            _ = v131
                            _ = v132
                        }
                        _ = v123
                    }
                }
                var v133 string = "hello"
                _ = v112
                _ = v133
            } else {
                switch genFn1(v30) {
                case -815:
                    v134 := []int{}
                    _ = v134
                case 847:
                    if (v25 && (3.14159 > (v4 * v3))) {
                        v135 := []int{127, -14}
                        v25 = v25
                        v64 += v64
                        _ = v135
                    }
                case 2147483647:
                    switch v5 {
                    case "foo bar":
                        if false {
                            for fi136 := 0; fi136 < 18; fi136++ {
                                var v137 int64 = -427
                                v111 += genFn0(-1, (v111 | v2))
                                _ = fi136
                                _ = v137
                            }
                            var v138 int8 = v30
                            switch -2147483648 {
                            case 2147483647:
                                if false {
                                    v5 = "a"
                                    v65 = 156
                                    switch ("hello" + "\n\t") {
                                    case "foo bar":
                                        var v139 float64 = 1e10
                                        _ = v139
                                    case "hello":
                                        v140 := GenS8{F0: 751}
                                        _ = v140
                                    case "\n\t":
                                        v141 := []int{int(v65), 9223372036854775807}
                                        _ = v141
                                    default:
                                        v142 := []int{128, 289, genFn1((v138 & (v138 / (((v30 - (v138 * v30))) | 1)))), 9223372036854775807}
                                        _ = v142
                                    }
                                } else {
                                    var v143 uint = (v24 % (((^(v24))) | 1))
                                    if (!false) {
                                        if (v65 == (v65 % (((v65 | (v65 >> ((v143) & 63)))) | 1))) {
                                            if (v138 < (v30 | v30)) {
                                                var v144 string = ("foo bar" + v5)
                                                if false {
                                                    var v145 uint64 = uint64(v24)
                                                    switch genFn1(int8(v24)) {
                                                    case -16:
                                                        var v146 bool = false
                                                        _ = v146
                                                    case 2147483647:
                                                        if (!true) {
                                                            v147 := []int{-286, -643, 9223372036854775807}
                                                            _ = v147
                                                        } else {
                                                            switch "" {
                                                            case "foo bar":
                                                                v111 = genFn0(v1, v2)
                                                            }
                                                            fmt.Println(v145, v145, v65)
                                                        }
                                                    case -2147483648:
                                                        for fi148 := 0; fi148 < 5; fi148++ {
                                                            v5 = v5
                                                            v149 := GenS15{F0: v5, F1: -351}
                                                            v25 = ((3.14159 > 0.0) || v25)
                                                            _ = fi148
                                                            _ = v149
                                                        }
                                                    }
                                                    v111 = (v2 / ((v2) | 1))
                                                    _ = v145
                                                }
                                                var v150 int8 = v30
                                                _ = v144
                                                _ = v150
                                            } else {
                                                fmt.Println(v4, v138)
                                            }
                                            v151 := []int{genFn1(v30)}
                                            _ = v151
                                        } else {
                                            if (!v25) {
                                                for fi152 := 0; fi152 < 18; fi152++ {
                                                    var v153 uint32 = 2147483648
                                                    v154 := []int{fi152}
                                                    var v155 uint64 = 18446744073709551615
                                                    _ = fi152
                                                    _ = v153
                                                    _ = v154
                                                    _ = v155
                                                }
                                            } else {
                                                switch "a" {
                                                case "a":
                                                    var v156 bool = ((v1 - 32767) <= (v1 | (v1 ^ (v1 | (v1 * (-(v1)))))))
                                                    _ = v156
                                                default:
                                                    v138 = int8(v1)
                                                }
                                                for fi157 := 0; fi157 < 11; fi157++ {
                                                    v2 -= v2
                                                    v158 := GenS16{F0: v64, F1: genFn1(v30), F2: (v5 + v5)}
                                                    for fi159 := 0; fi159 < 4; fi159++ {
                                                        v65 -= (v65 << ((840) & 63))
                                                        _ = fi159
                                                    }
                                                    _ = fi157
                                                    _ = v158
                                                }
                                            }
                                            var v160 uint8 = 0
                                            var v161 int = genFn1((v30 / ((v138) | 1)))
                                            _ = v160
                                            _ = v161
                                        }
                                        v2 = 65535
                                    }
                                    _ = v143
                                }
                            case -2147483648:
                                var v162 uint16 = 628
                                _ = v162
                            }
                            _ = v138
                        } else {
                            v163 := GenS11{F0: -805}
                            _ = v163
                        }
                    }
                default:
                    var v164 uint16 = 0
                    _ = v164
                }
            }
        }
    case 128:
        switch 32767 {
        case 2147483647:
            var v165 uint64 = 378
            _ = v165
        }
    }
    fmt.Println(v64, v111, v3)
    v166 := []int{-1, genFn1(19), 2147483647, genFn1(v30)}
    v1 = v1
    v4 += 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v111:", v111)
    fmt.Println("slicelen v166:", len(v166), v166)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v111
    _ = v166
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v64
    _ = v65
}
