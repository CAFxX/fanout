package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 uint; F2 uint }
type GenS3 struct { F0 int; F1 int16; F2 int }
type GenS4 struct { F0 int32; F1 string }
type GenS5 struct { F0 string; F1 int8 }
type GenS6 struct { F0 uint16; F1 uint8; F2 bool }
type GenS7 struct { F0 string }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int; F1 bool; F2 int16 }
type GenS10 struct { F0 int16; F1 uint16; F2 uint8 }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 int8; F1 string; F2 uint16 }
type GenS13 struct { F0 int; F1 uint64 }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 int64; F1 uint64 }
type GenS16 struct { F0 uint32; F1 uint8; F2 int16 }
type GenS17 struct { F0 int64 }
type GenS18 struct { F0 uint }
type GenS19 struct { F0 int32; F1 int32 }
func genFn0(p0 uint16, p1 uint32) uint8 {
    if (true || (!false)) {
        if (!true) {
            if (32767 > -150) {
                var v6 int32 = 889
                _ = v6
            } else {
                p0 -= (p0 / ((p0) | 1))
                fmt.Println(p0, p0)
                var v7 int = 32768
                _ = v7
            }
            var v8 int32 = -1
            _ = v8
        }
    }
    if false {
        p0 = 42
        if (48 <= 185) {
            v9 := GenS1{F0: 2}
            p0 = (p0 / ((p0) | 1))
            for fi10 := 0; fi10 < 8; fi10++ {
                v11 := GenS1{F0: 464}
                _ = fi10
                _ = v11
            }
            _ = v9
        } else {
            for fi12 := 0; fi12 < 6; fi12++ {
                var v13 uint = 18446744073709551615
                _ = fi12
                _ = v13
            }
        }
    }
    var v14 uint8 = 0
    var v15 string = "hello"
    v16 := []int{978, 42, 1, 9223372036854775807}
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    _ = v16
    return v14
}
func genFn1(p0 uint, p1 uint, p2 uint32) uint32 {
    v17 := GenS2{F0: 127, F1: 4294967295, F2: 42}
    switch "hello" {
    case "a":
        fmt.Println(p0)
    case "foo bar":
        if false {
            p1 -= uint(p2)
            switch int(p2) {
            case 105:
                var v18 string = "foo bar"
                _ = v18
            }
        } else {
            v19 := GenS3{F0: 9223372036854775807, F1: 32767, F2: 2147483647}
            _ = v19
        }
    default:
        var v20 uint32 = (p2 % ((p2) | 1))
        _ = v20
    }
    var v21 int32 = -2147483648
    for fi22 := 0; fi22 < 15; fi22++ {
        if false {
            for fi23 := 0; fi23 < 7; fi23++ {
                var v24 int = (fi22 + fi23)
                _ = fi23
                _ = v24
            }
        } else {
            v25 := GenS1{F0: 37}
            switch int(p1) {
            case -9223372036854775808:
                switch -757 {
                case 32768:
                    fmt.Println(v21)
                case -9223372036854775808:
                    fmt.Println(v21, p2)
                case 4294967295:
                    for fi26 := 0; fi26 < 19; fi26++ {
                        if (!true) {
                            v27 := []int{(fi22 >> ((255) & 63)), (fi26 + fi22), 894, (-(fi26))}
                            _ = v27
                        } else {
                            if false {
                                v28 := GenS4{F0: v21, F1: ""}
                                v29 := GenS5{F0: "\n\t", F1: 100}
                                _ = v28
                                _ = v29
                            }
                            var v30 int = -1
                            _ = v30
                        }
                        v31 := []int{(fi26 + 176), (-(fi22)), fi26}
                        _ = fi26
                        _ = v31
                    }
                default:
                    var v32 int32 = -1
                    _ = v32
                }
            case 128:
                v33 := GenS4{F0: (v21 ^ v21), F1: "foo bar"}
                _ = v33
            case -888:
                if ((true || (-9223372036854775808 > -750)) || (false && true)) {
                    v34 := []int{(^(fi22)), 573}
                    for fi35 := 0; fi35 < 16; fi35++ {
                        var v36 int64 = 9223372036854775807
                        _ = fi35
                        _ = v36
                    }
                    if true {
                        var v37 uint = p1
                        _ = v37
                    } else {
                        p1 = 1
                        if ("" <= "a") {
                            var v38 int8 = 1
                            if (4294967295 == 86) {
                                v39 := GenS3{F0: 9223372036854775807, F1: 37, F2: fi22}
                                var v40 float32 = (0.1 - 1e10)
                                var v41 int = (fi22 & fi22)
                                _ = v39
                                _ = v40
                                _ = v41
                            } else {
                                for fi42 := 0; fi42 < 18; fi42++ {
                                    switch (fi42 - fi22) {
                                    case -9223372036854775808:
                                        v43 := GenS2{F0: fi22, F1: 972, F2: 581}
                                        _ = v43
                                    case 32768:
                                        var v44 int64 = 42
                                        _ = v44
                                    default:
                                        var v45 string = "hello"
                                        _ = v45
                                    }
                                    var v46 int32 = (v21 / ((int32(p2)) | 1))
                                    var v47 uint8 = genFn0(369, (p2 >> (((p1 | p0)) & 63)))
                                    _ = fi42
                                    _ = v46
                                    _ = v47
                                }
                                v34 = append(v34, 1)
                            }
                            _ = v38
                        }
                    }
                    _ = v34
                }
            }
            _ = v25
        }
        _ = fi22
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v21
    return p2
}

func main() {
    var v1 int = -2147483648
    var v2 bool = true
    var v3 int64 = -9223372036854775808
    var v4 float64 = -2.25
    v5 := []uint32{42, 42, 65535}
    for fi48 := 0; fi48 < 8; fi48++ {
        v4 += float64(0.0)
        v1 = fi48
        _ = fi48
    }
    var v49 float64 = v4
    var v50 float64 = 0.1
    var v51 uint8 = 1
    for fi52 := 0; fi52 < 10; fi52++ {
        v1 += v1
        if (-98 == 0) {
            v53 := GenS6{F0: 189, F1: 127, F2: v2}
            fmt.Println(v1, v49, v2)
            for fi54 := 0; fi54 < 18; fi54++ {
                v55 := GenS7{F0: (("hello" + "\n\t") + "a")}
                v4 = v49
                _ = fi54
                _ = v55
            }
            _ = v53
        } else {
            switch "foo bar" {
            case "a":
                fmt.Println(v2, v3)
            case "foo bar":
                v56 := GenS6{F0: 262, F1: genFn0(255, 65535), F2: v2}
                _ = v56
            case "\n\t":
                for fi57 := 0; fi57 < 6; fi57++ {
                    v58 := GenS8{F0: v2}
                    if v2 {
                        v4 -= -2.25
                        var v59 int8 = 42
                        var v60 int64 = v3
                        _ = v59
                        _ = v60
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = fi57
                    _ = v58
                }
            }
        }
        _ = fi52
    }
    if (v2 && (false || true)) {
        var v61 int16 = 32767
        v2 = ((-(v61)) <= 32767)
        v4 += v4
        _ = v61
    }
    v51 = (v51 | 128)
    switch (v1 * -712) {
    case 724:
        if v2 {
            v4 = v4
            for fi62 := 0; fi62 < 18; fi62++ {
                var v63 uint32 = 1
                var v64 string = "hello"
                switch v64 {
                case "":
                    switch ((v64 + "a") + "foo bar") {
                    case "hello":
                        switch v64 {
                        case "\n\t":
                            _ = len(v5)
                        case "a":
                            var v65 uint32 = (^(v63))
                            _ = v65
                        case "hello":
                            switch v64 {
                            case "\n\t":
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            case "hello":
                                v66 := GenS9{F0: fi62, F1: v2, F2: 42}
                                _ = v66
                            }
                        default:
                            v64 = v64
                        }
                    case "a":
                        var v67 uint = 255
                        _ = v67
                    default:
                        v49 -= (v49 / v49)
                    }
                case "hello":
                    v50 -= 0.0
                }
                _ = fi62
                _ = v63
                _ = v64
            }
        } else {
            var v68 int8 = 127
            for fi69 := 0; fi69 < 5; fi69++ {
                if (!v2) {
                    for fi70 := 0; fi70 < 2; fi70++ {
                        fmt.Println(fi70, v68)
                        _ = fi70
                    }
                    v5 = append(v5, 0)
                    if (false && v2) {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        for fi71 := 0; fi71 < 9; fi71++ {
                            v5 = append(v5, genFn1(4294967295, 255, 800))
                            _ = fi71
                        }
                    } else {
                        var v72 uint8 = (v51 % ((genFn0(657, 2147483647)) | 1))
                        v2 = (v2 && false)
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = v72
                    }
                }
                var v73 uint32 = 2
                if (9223372036854775807 < v3) {
                    var v74 int64 = v3
                    switch "hello" {
                    case "hello":
                        if (!false) {
                            for fi75 := 0; fi75 < 14; fi75++ {
                                v76 := GenS6{F0: 255, F1: genFn0(725, genFn1(4294967295, 65, (v73 & v73))), F2: (1 < 1)}
                                if (1 >= -32768) {
                                    var v77 int8 = 100
                                    _ = v77
                                } else {
                                    v68 += v68
                                }
                                var v78 int32 = -1
                                v4 += v49
                                _ = fi75
                                _ = v76
                                _ = v78
                            }
                        } else {
                            _ = len(v5)
                        }
                    case "":
                        for fi79 := 0; fi79 < 8; fi79++ {
                            v5 = append(v5, 2147483648)
                            v80 := GenS5{F0: ("" + ""), F1: 100}
                            v68 -= v68
                            var v81 bool = false
                            var v82 int32 = -279
                            var v83 bool = ((fi69 * v1) < fi79)
                            _ = fi79
                            _ = v80
                            _ = v81
                            _ = v82
                            _ = v83
                        }
                    case "\n\t":
                        v5 = append(v5, genFn1(686, 42, (v73 + v73)))
                    }
                    _ = v74
                } else {
                    v84 := GenS8{F0: v2}
                    v2 = v2
                    _ = v84
                }
                _ = fi69
                _ = v73
            }
            v51 -= genFn0(324, 42)
            _ = v68
        }
    default:
        fmt.Println(v3, v4, v4)
    }
    for fi85 := 0; fi85 < 14; fi85++ {
        if len(v5) > 0 {
            v5[0] = 4294967295
        }
        _ = fi85
    }
    if (!(!v2)) {
        if v2 {
            for fi86 := 0; fi86 < 6; fi86++ {
                v50 += v4
                var v87 bool = true
                _ = fi86
                _ = v87
            }
            v88 := GenS10{F0: 1, F1: 0, F2: 0}
            v89 := GenS11{F0: 51}
            _ = v88
            _ = v89
        } else {
            v90 := GenS5{F0: (("foo bar" + "\n\t") + ("hello" + "foo bar")), F1: 100}
            var v91 uint = 65535
            if true {
                if false {
                    v5 = append(v5, 116)
                    _ = len(v5)
                }
            }
            _ = v90
            _ = v91
        }
        v92 := GenS10{F0: -236, F1: 32767, F2: (v51 << ((569) & 63))}
        _ = v92
    }
    if false {
        v93 := GenS4{F0: 2147483647, F1: ("a" + "a")}
        var v94 string = "\n\t"
        if false {
            switch v1 {
            case 32768:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            case 1:
                v1 = v1
            default:
                fmt.Println(v3)
            }
        } else {
            if (!(324 >= 2147483647)) {
                var v95 float32 = 3.14159
                fmt.Println(v50, v3, v49)
                for fi96 := 0; fi96 < 9; fi96++ {
                    v51 = 1
                    for fi97 := 0; fi97 < 20; fi97++ {
                        switch (v94 + v94) {
                        case "\n\t":
                            fmt.Println(v95)
                        case "":
                            var v98 uint = 18446744073709551615
                            _ = v98
                        case "a":
                            if true {
                                _ = len(v5)
                            } else {
                                v99 := GenS12{F0: -128, F1: v94, F2: 42}
                                _ = v99
                            }
                        }
                        var v100 uint32 = 2
                        _ = fi97
                        _ = v100
                    }
                    fmt.Println(v2)
                    _ = fi96
                }
                _ = v95
            }
            switch v94 {
            case "foo bar":
                if (!v2) {
                    fmt.Println(v50, v49, v51)
                    var v101 uint32 = 4294967295
                    for fi102 := 0; fi102 < 2; fi102++ {
                        v103 := GenS13{F0: (v1 >> ((4294967295) & 63)), F1: 1}
                        if (int64(v101) == -9223372036854775808) {
                            for fi104 := 0; fi104 < 13; fi104++ {
                                if v2 {
                                    fmt.Println(v2, v4)
                                    switch v1 {
                                    case 180:
                                        switch -459 {
                                        case -918:
                                            v101 += (^(v101))
                                        case 476:
                                            v50 = ((1e10 / v50) + v4)
                                        }
                                    default:
                                        if ((-(fi102)) <= v1) {
                                            v3 += v3
                                            switch "" {
                                            case "a":
                                                var v105 int32 = 42
                                                _ = v105
                                            default:
                                                _ = len(v5)
                                            }
                                            switch v94 {
                                            case "\n\t":
                                                v106 := GenS2{F0: fi102, F1: 18446744073709551615, F2: 18446744073709551615}
                                                _ = v106
                                            case "foo bar":
                                                if v2 {
                                                    var v107 uint8 = (^(v51))
                                                    v108 := GenS2{F0: v1, F1: 2, F2: 65535}
                                                    _ = v107
                                                    _ = v108
                                                }
                                            default:
                                                v49 += v49
                                                _ = len(v5)
                                            }
                                        } else {
                                            var v109 int32 = -390
                                            if (0 > 714) {
                                                v109 = v109
                                            }
                                            if (int64(v101) == v3) {
                                                v1 = (v1 & -2147483648)
                                                fmt.Println(v49, v1, fi102)
                                                var v110 float32 = 3.14159
                                                _ = v110
                                            } else {
                                                if len(v5) > 0 {
                                                    v5[0] = (^(v101))
                                                }
                                                v50 = v4
                                            }
                                            var v111 int64 = 42
                                            _ = v109
                                            _ = v111
                                        }
                                    }
                                } else {
                                    v112 := GenS14{F0: 2}
                                    _ = v112
                                }
                                var v113 float64 = 0.1
                                var v114 int8 = -35
                                if false {
                                    v115 := GenS14{F0: 0}
                                    switch v94 {
                                    case "a":
                                        var v116 uint = 42
                                        _ = v116
                                    case "\n\t":
                                        _ = len(v5)
                                    default:
                                        v51 = v51
                                    }
                                    fmt.Println(v4)
                                    if (0 < 2) {
                                        fmt.Println(v2)
                                        v117 := GenS12{F0: (-(v114)), F1: "foo bar", F2: 0}
                                        _ = v117
                                    }
                                    _ = v115
                                }
                                var v118 int64 = 0
                                v1 += fi102
                                var v119 bool = false
                                var v120 int64 = v3
                                _ = fi104
                                _ = v113
                                _ = v114
                                _ = v118
                                _ = v119
                                _ = v120
                            }
                        }
                        for fi121 := 0; fi121 < 16; fi121++ {
                            switch ((v94 + "foo bar") + v94) {
                            case "foo bar":
                                var v122 int64 = (v3 + -1)
                                _ = v122
                            }
                            var v123 int32 = int32(v101)
                            _ = fi121
                            _ = v123
                        }
                        _ = fi102
                        _ = v103
                    }
                    _ = v101
                } else {
                    v124 := GenS11{F0: 0}
                    if (0 < 457) {
                        for fi125 := 0; fi125 < 14; fi125++ {
                            v126 := GenS9{F0: (v1 * (v1 ^ -9223372036854775808)), F1: (!v2), F2: 42}
                            v2 = (325 != 423)
                            for fi127 := 0; fi127 < 12; fi127++ {
                                var v128 int32 = -2147483648
                                var v129 uint16 = 0
                                if (!(347 >= 2147483648)) {
                                    v130 := GenS15{F0: v3, F1: 0}
                                    v131 := GenS15{F0: 9223372036854775807, F1: 18446744073709551615}
                                    _ = v130
                                    _ = v131
                                } else {
                                    switch (fi127 % ((624) | 1)) {
                                    case 2147483647:
                                        v132 := GenS8{F0: (!false)}
                                        _ = v132
                                    case -560:
                                        var v133 uint64 = 42
                                        _ = v133
                                    default:
                                        if len(v5) > 0 {
                                            v5[0] = genFn1(0, 255, uint32(v129))
                                        }
                                    }
                                    if false {
                                        var v134 int16 = -559
                                        fmt.Println(v94)
                                        _ = v134
                                    }
                                }
                                fmt.Println(v51, v2, v50)
                                _ = fi127
                                _ = v128
                                _ = v129
                            }
                            _ = fi125
                            _ = v126
                        }
                        var v135 int8 = 127
                        if ((!v2) && v2) {
                            for fi136 := 0; fi136 < 10; fi136++ {
                                v137 := GenS16{F0: 161, F1: 1, F2: 561}
                                if v2 {
                                    fmt.Println(fi136, fi136)
                                    for fi138 := 0; fi138 < 2; fi138++ {
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        var v139 int64 = v3
                                        _ = fi138
                                        _ = v139
                                    }
                                }
                                _ = fi136
                                _ = v137
                            }
                        } else {
                            var v140 uint8 = (^(v51))
                            _ = v140
                        }
                        _ = v135
                    }
                    _ = v124
                }
            case "":
                fmt.Println(v51)
            default:
                switch v1 {
                case 2147483647:
                    for fi141 := 0; fi141 < 3; fi141++ {
                        v4 = float64((-2.25 + float32(v50)))
                        _ = fi141
                    }
                case 430:
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                case 4294967295:
                    if false {
                        var v142 int64 = v3
                        fmt.Println(v94, v50)
                        _ = v142
                    }
                }
            }
        }
        _ = v93
        _ = v94
    } else {
        for fi143 := 0; fi143 < 11; fi143++ {
            v50 -= float64(0.0)
            if false {
                if false {
                    v4 += (v50 / 3.14159)
                } else {
                    for fi144 := 0; fi144 < 12; fi144++ {
                        if v2 {
                            v1 = v1
                        }
                        _ = fi144
                    }
                    fmt.Println(v1)
                }
                fmt.Println(v51)
                if (1 < 0) {
                    v145 := GenS17{F0: v3}
                    _ = v145
                }
            } else {
                v4 = v4
                var v146 int32 = -1
                _ = v146
            }
            if v2 {
                if v2 {
                    var v147 int8 = 0
                    _ = v147
                }
                v4 = 1e10
            }
            _ = fi143
        }
        v148 := GenS18{F0: 4294967295}
        var v149 int = 32768
        _ = v148
        _ = v149
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v5 = append(v5, genFn1(65535, 18446744073709551615, 1))
    switch (v1 - 32768) {
    case -12:
        v4 += (v4 / 1e10)
    case 32767:
        if v2 {
            v150 := GenS16{F0: 916, F1: (v51 % ((v51) | 1)), F2: 1}
            v151 := GenS8{F0: (false && v2)}
            _ = v150
            _ = v151
        } else {
            v152 := GenS14{F0: 1}
            var v153 uint = 2
            v3 = -9223372036854775808
            _ = v152
            _ = v153
        }
    }
    v154 := GenS17{F0: (v3 + (v3 + int64(v1)))}
    switch v1 {
    case 169:
        fmt.Println(v50, v3, v4)
    case 149:
        switch "" {
        case "foo bar":
            switch "foo bar" {
            case "a":
                v155 := GenS19{F0: 994, F1: 407}
                _ = v155
            case "":
                fmt.Println(v3, v2, v51)
            case "\n\t":
                var v156 int32 = 0
                _ = v156
            default:
                v50 += 1e10
            }
        case "a":
            var v157 float32 = (3.14159 / 1.0)
            _ = v157
        }
    }
    switch (-(v1)) {
    case 0:
        fmt.Println(v49, v49)
    default:
        var v158 uint8 = v51
        _ = v158
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v154
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v51
}
