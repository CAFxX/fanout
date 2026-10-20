package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 uint64; F1 uint16; F2 int64 }
type GenS3 struct { F0 uint16; F1 string }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint64; F1 string; F2 int }
type GenS8 struct { F0 uint16; F1 int8; F2 int }
type GenS9 struct { F0 int8; F1 uint16; F2 uint64 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 uint64; F1 int64 }
type GenS12 struct { F0 uint64; F1 bool; F2 uint32 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint64 = 399
    var v3 int16 = 0
    v4 := make(map[string]uint32)
    var v5 float32 = ((0.1 * 0.1) + -2.25)
    v6 := make(map[string]bool)
    v3 -= 871
    var v7 string = (("foo bar" + "foo bar") + "")
    v3 = 1
    if ((v3 - (v3 | -32768)) < v3) {
        v5 = v5
        v8 := v6[(v7 + (v7 + "foo bar"))]
        v9 := GenS1{F0: "a", F1: 127}
        _ = v8
        _ = v9
    } else {
        if false {
            if true {
                if true {
                    fmt.Println(v2, v3, v5)
                    _ = len(v6)
                }
                v10 := []int{2147483647, 643, 128, -1}
                _ = v10
            }
            var v11 int8 = 116
            v12 := GenS2{F0: (v2 % (((v2 | (v2 << ((uint(v2)) & 63)))) | 1)), F1: 106, F2: -329}
            _ = v11
            _ = v12
        } else {
            v2 -= (v2 & v2)
            v5 += v5
        }
        var v13 uint16 = 32767
        if false {
            if (0.0 <= 1.5) {
                fmt.Println(v5)
                fmt.Println(v2, v7, v7)
                v7 = ((v7 + v7) + (v7 + v7))
            } else {
                v14 := []int{128, 369}
                _ = v14
            }
        } else {
            v13 += (v13 % (((v13 * v13)) | 1))
            switch v7 {
            case "a":
                for fi15 := 0; fi15 < 9; fi15++ {
                    for fi16 := 0; fi16 < 17; fi16++ {
                        v17 := []int{42, 4294967295}
                        fmt.Println(v13)
                        _ = fi16
                        _ = v17
                    }
                    _ = fi15
                }
            case "hello":
                v18 := GenS3{F0: (v13 - v13), F1: v7}
                _ = v18
            default:
                fmt.Println(v3, v3)
            }
            var v19 float64 = 1e10
            _ = v19
        }
        _ = v13
    }
    for fi20 := 0; fi20 < 2; fi20++ {
        v1[2147483648] = fi20
        var v21 int8 = 127
        _ = fi20
        _ = v21
    }
    v3 -= 32767
    if (!true) {
        for fi22 := 0; fi22 < 2; fi22++ {
            var v23 float32 = v5
            v7 = v7
            _ = fi22
            _ = v23
        }
    }
    var v24 uint64 = v2
    _ = len(v4)
    if (false && false) {
        for fi25 := 0; fi25 < 6; fi25++ {
            v26 := v6[v7]
            var v27 int = (-(fi25))
            _ = fi25
            _ = v26
            _ = v27
        }
        switch (v7 + "\n\t") {
        case "\n\t":
            if false {
                v2 -= v24
                var v28 int32 = -2147483648
                switch (((("\n\t" + (v7 + (("" + v7) + "a"))) + "") + (("" + ("hello" + ("hello" + v7))) + v7)) + (v7 + v7)) {
                case "a":
                    if ((16 >= -108) && ((v28 <= -443) && (false && (true && false)))) {
                        for fi29 := 0; fi29 < 1; fi29++ {
                            v28 = 2147483647
                            v28 = 1
                            _ = fi29
                        }
                        v7 = v7
                        var v30 uint8 = 55
                        _ = v30
                    } else {
                        for fi31 := 0; fi31 < 10; fi31++ {
                            fmt.Println(v7, fi31, v5)
                            v3 += int16(v2)
                            _ = fi31
                        }
                        if false {
                            v32 := GenS3{F0: 2, F1: v7}
                            switch 230 {
                            case -9223372036854775808:
                                for fi33 := 0; fi33 < 10; fi33++ {
                                    if false {
                                        for fi34 := 0; fi34 < 13; fi34++ {
                                            v35 := GenS2{F0: v2, F1: 2, F2: 102}
                                            v36 := []int{}
                                            v37 := GenS3{F0: 0, F1: (((v7 + "") + (v7 + "a")) + "a")}
                                            _ = fi34
                                            _ = v35
                                            _ = v36
                                            _ = v37
                                        }
                                        switch (fi33 & (fi33 * 0)) {
                                        case -868:
                                            var v38 int8 = -1
                                            _ = v38
                                        case 4294967295:
                                            fmt.Println(v24, v3, v24)
                                        }
                                        v3 -= (v3 - (v3 ^ v3))
                                    } else {
                                        for fi39 := 0; fi39 < 6; fi39++ {
                                            if (42 >= -128) {
                                                v2 = (v2 % ((v2) | 1))
                                                var v40 uint = 97
                                                v7 = "hello"
                                                _ = v40
                                            }
                                            var v41 int16 = v3
                                            fmt.Println(v5, v2)
                                            _ = fi39
                                            _ = v41
                                        }
                                        var v42 string = v7
                                        var v43 int = fi33
                                        _ = v42
                                        _ = v43
                                    }
                                    v7 = v7
                                    _ = fi33
                                }
                            default:
                                v5 = (v5 + 1.5)
                            }
                            for fi44 := 0; fi44 < 20; fi44++ {
                                var v45 string = v7
                                v46 := GenS1{F0: ((((v7 + ("hello" + v7)) + "foo bar") + v45) + "\n\t"), F1: 42}
                                _ = fi44
                                _ = v45
                                _ = v46
                            }
                            _ = v32
                        } else {
                            var v47 int8 = 1
                            v1[1] = 0
                            _ = v47
                        }
                    }
                }
                _ = v28
            }
        case "foo bar":
            if true {
                var v48 int64 = -1
                if (443 >= 73) {
                    v49 := GenS4{F0: 0}
                    var v50 int32 = 0
                    _ = v49
                    _ = v50
                } else {
                    v48 = (v48 - (v48 + 42))
                }
                _ = v48
            } else {
                var v51 uint = 255
                var v52 uint16 = 1
                _ = v51
                _ = v52
            }
        default:
            v53 := GenS3{F0: 905, F1: v7}
            _ = v53
        }
        delete(v6, "foo bar")
    }
    v2 = (v24 / ((v24) | 1))
    for fi54 := 0; fi54 < 3; fi54++ {
        switch "hello" {
        case "hello":
            var v55 int = fi54
            _ = v55
        default:
            if true {
                for fi56 := 0; fi56 < 20; fi56++ {
                    v57 := v4["\n\t"]
                    v57 += v57
                    _ = fi56
                    _ = v57
                }
                var v58 int8 = 100
                var v59 float64 = 0.0
                _ = v58
                _ = v59
            } else {
                if false {
                    v60 := GenS3{F0: 255, F1: "foo bar"}
                    fmt.Println(v24, fi54)
                    fmt.Println(v24)
                    _ = v60
                } else {
                    if false {
                        v6[((v7 + v7) + (v7 + "\n\t"))] = false
                        var v61 bool = (982 >= 657)
                        _ = v61
                    } else {
                        if true {
                            fmt.Println(fi54, v3, v24)
                            if ((v3 != v3) || true) {
                                for fi62 := 0; fi62 < 7; fi62++ {
                                    v63 := GenS4{F0: 876}
                                    var v64 int32 = 291
                                    var v65 uint = 18446744073709551615
                                    var v66 bool = (1.5 >= 0.0)
                                    _ = fi62
                                    _ = v63
                                    _ = v64
                                    _ = v65
                                    _ = v66
                                }
                                if false {
                                    v5 = (float32(-2.25) + (float32(1e10) * v5))
                                    if true {
                                        if (128 > 42) {
                                            v67 := []int{fi54, fi54, int(v2), (fi54 ^ fi54)}
                                            v2 = (v2 + (v2 >> ((999) & 63)))
                                            v2 = (v2 + 797)
                                            _ = v67
                                        }
                                        for fi68 := 0; fi68 < 20; fi68++ {
                                            fmt.Println(v3, fi68, v7)
                                            var v69 uint16 = 0
                                            switch (v7 + (("hello" + "hello") + v7)) {
                                            case "\n\t":
                                                v70 := v6[((((v7 + v7) + "foo bar") + v7) + v7)]
                                                _ = v70
                                            }
                                            _ = fi68
                                            _ = v69
                                        }
                                    } else {
                                        v2 = 4294967295
                                        for fi71 := 0; fi71 < 3; fi71++ {
                                            delete(v6, "hello")
                                            _ = fi71
                                        }
                                    }
                                } else {
                                    v5 = 3.14159
                                    v2 += v24
                                    v72 := GenS3{F0: 65535, F1: ("foo bar" + v7)}
                                    _ = v72
                                }
                            } else {
                                for fi73 := 0; fi73 < 8; fi73++ {
                                    _ = len(v1)
                                    v24 += 172
                                    v1[uint32(v24)] = fi73
                                    fmt.Println(v7)
                                    v74 := GenS5{F0: 0}
                                    _ = fi73
                                    _ = v74
                                }
                                v24 += v24
                                var v75 int16 = (v3 / (((v3 | v3)) | 1))
                                _ = v75
                            }
                            v5 += v5
                        } else {
                            var v76 bool = (3.14159 > (1.5 / ((-2.25 + v5) / (v5 / ((v5 * (0.1 * v5)) + (v5 / v5))))))
                            _ = v76
                        }
                    }
                }
                for fi77 := 0; fi77 < 15; fi77++ {
                    v78 := GenS2{F0: v24, F1: 42, F2: 9223372036854775807}
                    _ = fi77
                    _ = v78
                }
                v5 = v5
            }
        }
        _ = fi54
    }
    if false {
        v1[4294967295] = 128
        delete(v6, (v7 + ""))
    } else {
        v79 := GenS1{F0: ((v7 + v7) + ""), F1: 42}
        _ = v79
    }
    switch 1 {
    case -831:
        for fi80 := 0; fi80 < 15; fi80++ {
            var v81 uint16 = 932
            _ = fi80
            _ = v81
        }
    default:
        if (42 == 2147483647) {
            if false {
                fmt.Println(v5)
            }
            delete(v4, (v7 + "a"))
            v82 := GenS5{F0: 0}
            _ = v82
        }
    }
    v83 := GenS1{F0: v7, F1: 2}
    if (0 == -9223372036854775808) {
        v24 += v2
        if true {
            var v84 float64 = 3.14159
            v85 := v1[42]
            _ = v84
            _ = v85
        } else {
            v86 := []int{-1, -1, 128}
            for fi87 := 0; fi87 < 12; fi87++ {
                if (false || true) {
                    _ = len(v86)
                    var v88 int8 = -61
                    for fi89 := 0; fi89 < 19; fi89++ {
                        v90 := GenS4{F0: -2147483648}
                        switch (fi87 % (((fi89 * 2147483647)) | 1)) {
                        case 0:
                            var v91 int = 127
                            _ = v91
                        default:
                            var v92 float32 = float32((0.1 - 0.0))
                            _ = v92
                        }
                        if (v3 <= (-(v3))) {
                            v7 = v7
                            v93 := GenS3{F0: 255, F1: v7}
                            _ = v93
                        }
                        if ("" > v7) {
                            var v94 int16 = v3
                            _ = v94
                        } else {
                            _ = len(v86)
                            switch (fi89 * (fi87 & (fi87 | fi87))) {
                            case 9223372036854775807:
                                if (true && true) {
                                    var v95 uint64 = (^(v2))
                                    if true {
                                        switch (^(fi89)) {
                                        case 0:
                                            var v96 float32 = (v5 - (-2.25 - -2.25))
                                            v7 = ((("\n\t" + (v7 + (v7 + "a"))) + v7) + ("a" + v7))
                                            var v97 uint32 = 635
                                            switch "a" {
                                            case "":
                                                v98 := GenS6{F0: 171}
                                                _ = v98
                                            case "hello":
                                                v86 = append(v86, (fi89 % ((fi87) | 1)))
                                            case "a":
                                                v99 := GenS7{F0: v2, F1: (v7 + v7), F2: (fi87 / (((fi87 - int(v3))) | 1))}
                                                _ = v99
                                            default:
                                                var v100 int = (fi87 ^ 9223372036854775807)
                                                _ = v100
                                            }
                                            _ = v96
                                            _ = v97
                                        case 32767:
                                            var v101 int64 = -9223372036854775808
                                            _ = v101
                                        }
                                    } else {
                                        switch (fi89 << ((1) & 63)) {
                                        case 4294967295:
                                            v102 := GenS3{F0: 32767, F1: v7}
                                            _ = v102
                                        case 2147483647:
                                            if true {
                                                if false {
                                                    var v103 uint32 = 1
                                                    _ = v103
                                                }
                                                v3 = v3
                                                var v104 uint16 = 487
                                                _ = v104
                                            } else {
                                                fmt.Println(fi87, fi89, v3)
                                            }
                                        default:
                                            switch 32768 {
                                            case 389:
                                                delete(v6, v7)
                                            case 32767:
                                                v7 = "foo bar"
                                            }
                                        }
                                    }
                                    _ = v95
                                } else {
                                    v5 += v5
                                }
                            case 127:
                                fmt.Println(v88)
                            case -634:
                                v5 += v5
                            default:
                                v105 := GenS4{F0: 603}
                                _ = v105
                            }
                            v106 := GenS1{F0: v7, F1: 0}
                            if false {
                                fmt.Println(v5)
                                if false {
                                    v2 = 954
                                    v7 = ((v7 + ((("a" + v7) + v7) + "hello")) + "a")
                                } else {
                                    fmt.Println(v88, v7, v3)
                                }
                                var v107 int16 = (v3 * v3)
                                _ = v107
                            } else {
                                v7 = "a"
                                v6["hello"] = (0 >= 2147483647)
                                v24 += (v24 * 510)
                            }
                            _ = v106
                        }
                        _ = fi89
                        _ = v90
                    }
                    _ = v88
                }
                v4[""] = 65535
                _ = fi87
            }
            if false {
                fmt.Println(v3, v3)
                if (42 != 255) {
                    for fi108 := 0; fi108 < 1; fi108++ {
                        if len(v86) > 0 {
                            v86[0] = fi108
                        }
                        v2 -= 18446744073709551615
                        if len(v86) > 0 {
                            v86[0] = 32767
                        }
                        _ = fi108
                    }
                    if false {
                        v3 += (v3 % ((0) | 1))
                        fmt.Println(v3, v5)
                        v24 -= 406
                    } else {
                        for fi109 := 0; fi109 < 1; fi109++ {
                            v110 := GenS6{F0: 42}
                            if false {
                                if (!false) {
                                    if (!(-128 > -118)) {
                                        var v111 int = 9223372036854775807
                                        v24 = (v2 ^ v24)
                                        var v112 uint32 = 42
                                        _ = v111
                                        _ = v112
                                    }
                                }
                                v6[v7] = true
                            } else {
                                v113 := GenS8{F0: 0, F1: -128, F2: fi109}
                                _ = v113
                            }
                            if true {
                                var v114 string = (((v7 + "a") + ("" + v7)) + "hello")
                                var v115 string = "\n\t"
                                v24 = (v2 / (((v2 - (v2 % ((v2) | 1)))) | 1))
                                _ = v114
                                _ = v115
                            }
                            _ = fi109
                            _ = v110
                        }
                        fmt.Println(v5, v24, v24)
                        v86 = append(v86, 2147483647)
                    }
                } else {
                    v116 := GenS9{F0: 0, F1: 0, F2: (v24 - (v2 << ((980) & 63)))}
                    delete(v6, "a")
                    v117 := GenS10{F0: 18446744073709551615}
                    _ = v116
                    _ = v117
                }
            } else {
                if (!((false || (2 <= 2)) && (4294967295 == 38))) {
                    for fi118 := 0; fi118 < 5; fi118++ {
                        v119 := GenS7{F0: 1, F1: v7, F2: (fi118 << ((147) & 63))}
                        v7 = v7
                        _ = fi118
                        _ = v119
                    }
                    fmt.Println(v5, v24, v2)
                    if (true || true) {
                        _ = len(v86)
                        if (798 > 4294967295) {
                            for fi120 := 0; fi120 < 20; fi120++ {
                                v121 := GenS1{F0: (v7 + v7), F1: 255}
                                if len(v86) >= 2 {
                                    v86 = v86[1:len(v86)]
                                }
                                _ = len(v86)
                                _ = fi120
                                _ = v121
                            }
                            for fi122 := 0; fi122 < 5; fi122++ {
                                if len(v86) >= 2 {
                                    v86 = v86[1:len(v86)]
                                }
                                for fi123 := 0; fi123 < 11; fi123++ {
                                    fmt.Println(v3, v3, v7)
                                    if (-2147483648 >= -49) {
                                        v5 = (v5 + float32(-2.25))
                                        if (!true) {
                                            fmt.Println(v3, fi122, v2)
                                        } else {
                                            if (65535 >= 45) {
                                                v124 := GenS4{F0: -188}
                                                _ = v124
                                            }
                                            if (float64((v5 / v5)) < 0.0) {
                                                v125 := GenS11{F0: v24, F1: -811}
                                                v24 = v2
                                                _ = v125
                                            } else {
                                                v5 = v5
                                                if (!(!(!(!(932 >= 291))))) {
                                                    v7 = ""
                                                }
                                            }
                                            v24 += v2
                                        }
                                    }
                                    fmt.Println(v7, v5, v7)
                                    _ = fi123
                                }
                                _ = fi122
                            }
                        } else {
                            v126 := GenS7{F0: 4294967295, F1: v7, F2: 128}
                            v1[0] = 746
                            var v127 float32 = v5
                            _ = v126
                            _ = v127
                        }
                        if len(v86) > 0 {
                            v86[0] = 400
                        }
                    } else {
                        v3 -= (v3 ^ v3)
                        v3 = v3
                    }
                } else {
                    v128 := v1[42]
                    _ = v128
                }
                _ = len(v1)
                v3 = v3
            }
            _ = v86
        }
        switch 50 {
        case -43:
            fmt.Println(v3, v7, v3)
        case 2147483647:
            for fi129 := 0; fi129 < 20; fi129++ {
                fmt.Println(v7, v24, v24)
                for fi130 := 0; fi130 < 6; fi130++ {
                    var v131 float32 = ((((1.5 - float32(1.5)) + v5) + 1e10) + 1e10)
                    switch "hello" {
                    case "hello":
                        switch "\n\t" {
                        case "":
                            var v132 uint64 = (v2 - 54)
                            _ = v132
                        case "a":
                            var v133 int32 = 0
                            _ = v133
                        case "foo bar":
                            if (true || false) {
                                v24 = 2
                                fmt.Println(v3, v3)
                                var v134 uint = 255
                                _ = v134
                            }
                        }
                    case "\n\t":
                        v7 = ("" + v7)
                    case "":
                        v135 := GenS3{F0: 255, F1: ("hello" + v7)}
                        _ = v135
                    default:
                        v131 = v131
                    }
                    v3 -= (-(v3))
                    _ = fi130
                    _ = v131
                }
                switch fi129 {
                case 172:
                    v136 := []int{-500, (fi129 - fi129), (fi129 >> ((18446744073709551615) & 63)), (-(fi129))}
                    _ = v136
                case -216:
                    for fi137 := 0; fi137 < 7; fi137++ {
                        switch "foo bar" {
                        case "hello":
                            v138 := GenS11{F0: 536, F1: 42}
                            _ = v138
                        case "a":
                            fmt.Println(v3)
                            v3 -= (v3 * (v3 * 1))
                        case "":
                            var v139 bool = false
                            _ = v139
                        }
                        if true {
                            var v140 int16 = (^(v3))
                            var v141 uint = 65535
                            _ = v140
                            _ = v141
                        }
                        v142 := []int{fi137, fi129}
                        _ = fi137
                        _ = v142
                    }
                default:
                    switch v7 {
                    case "hello":
                        delete(v6, v7)
                    case "\n\t":
                        fmt.Println(fi129)
                    case "":
                        switch ("" + "\n\t") {
                        case "\n\t":
                            v143 := GenS12{F0: (v24 >> ((2) & 63)), F1: true, F2: 2147483648}
                            _ = v143
                        default:
                            if false {
                                v2 += v24
                                for fi144 := 0; fi144 < 2; fi144++ {
                                    v145 := GenS4{F0: -2147483648}
                                    _ = fi144
                                    _ = v145
                                }
                                for fi146 := 0; fi146 < 15; fi146++ {
                                    fmt.Println(fi129, v7)
                                    v7 = ("hello" + ((v7 + v7) + v7))
                                    _ = fi146
                                }
                            }
                        }
                    }
                }
                _ = fi129
            }
        case -1:
            fmt.Println(v3, v24)
        default:
            v5 -= (0.1 + (((-2.25 + ((0.1 + -2.25) + 0.1)) * ((v5 + ((3.14159 * (float32(1.5) + 3.14159)) / float32((-2.25 * -2.25)))) + (v5 / 1e10))) * 1e10))
        }
    }
    fmt.Println(v3, v2, v7)
    v147 := v1[626]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v147:", v147)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v147
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v83
}
