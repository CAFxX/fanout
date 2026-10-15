package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 uint64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32; F1 int }
type GenS4 struct { F0 string; F1 bool }
type GenS5 struct { F0 string; F1 int }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint; F1 int; F2 uint64 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 uint32; F1 int8; F2 uint16 }
type GenS12 struct { F0 int16; F1 uint64 }
type GenS13 struct { F0 int64; F1 int8 }
func genFn0(p0 int32, p1 int, p2 uint64) uint {
    var v4 int64 = 837
    p2 = 576
    var v5 int32 = 26
    v4 = v4
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 18446744073709551615
}
func genFn1(p0 uint8) int8 {
    p0 += 128
    if (0 < 14) {
        for fi6 := 0; fi6 < 6; fi6++ {
            fmt.Println(fi6, p0)
            _ = fi6
        }
    }
    switch ("hello" + "\n\t") {
    case "foo bar":
        fmt.Println(p0)
    }
    var v7 bool = true
    switch ("hello" + "") {
    case "a":
        p0 -= (p0 / ((p0) | 1))
    case "hello":
        v7 = (1 == 762)
    }
    _ = p0
    _ = v7
    return 42
}

func main() {
    v1 := make(map[int]int)
    var v2 float32 = 0.0
    var v3 string = "a"
    v8 := []int{2147483647, -1, 1, 248}
    if false {
        if (2 == 591) {
            fmt.Println(v2, v2)
            v3 = v3
            switch "foo bar" {
            case "":
                var v9 uint32 = 4294967295
                _ = v9
            case "hello":
                _ = len(v1)
            case "\n\t":
                var v10 bool = false
                _ = v10
            }
        } else {
            var v11 uint8 = 1
            _ = v11
        }
        v12 := GenS1{F0: "hello", F1: 924, F2: 509}
        v2 -= v2
        _ = v12
    } else {
        switch (v3 + "foo bar") {
        case "":
            if (42 > 793) {
                if (1 >= 4294967295) {
                    var v13 int16 = 1
                    if (18446744073709551615 != 4294967295) {
                        fmt.Println(v3, v3, v13)
                        if true {
                            for fi14 := 0; fi14 < 5; fi14++ {
                                v13 += 32767
                                fmt.Println(fi14, v13, v3)
                                fmt.Println(v3, v3)
                                _ = fi14
                            }
                            v3 = (("a" + "a") + ("\n\t" + ("hello" + v3)))
                            var v15 bool = false
                            _ = v15
                        }
                    }
                    var v16 uint64 = 560
                    _ = v13
                    _ = v16
                }
            } else {
                for fi17 := 0; fi17 < 18; fi17++ {
                    var v18 int8 = 0
                    _ = fi17
                    _ = v18
                }
            }
        case "\n\t":
            v3 = (v3 + (v3 + v3))
        case "hello":
            v3 = "foo bar"
        default:
            var v19 uint8 = 2
            _ = v19
        }
        for fi20 := 0; fi20 < 19; fi20++ {
            v21 := GenS2{F0: 18}
            _ = fi20
            _ = v21
        }
        v2 = 1.5
    }
    v22 := GenS2{F0: 127}
    for fi23 := 0; fi23 < 3; fi23++ {
        for fi24 := 0; fi24 < 2; fi24++ {
            v3 = "foo bar"
            fmt.Println(v2)
            _ = fi24
        }
        v3 = (v3 + v3)
        _ = fi23
    }
    switch 9223372036854775807 {
    case 32768:
        v2 -= (v2 * (((v2 + float32(1.5)) / v2) - v2))
    case 0:
        fmt.Println(v3, v2)
    }
    if false {
        var v25 int32 = 1
        var v26 int16 = -531
        _ = v25
        _ = v26
    } else {
        v27 := GenS1{F0: v3, F1: 65535, F2: 42}
        v2 += (v2 + (0.0 - v2))
        switch -224 {
        case -170:
            v28 := GenS3{F0: -606, F1: 4294967295}
            _ = v28
        case -64:
            v2 -= ((v2 / v2) / v2)
        case 2147483647:
            var v29 uint32 = 42
            _ = v29
        }
        _ = v27
    }
    for fi30 := 0; fi30 < 1; fi30++ {
        var v31 uint64 = 4294967295
        _ = fi30
        _ = v31
    }
    if (583 > genFn0(-1, -1, 725)) {
        if (2 <= 0) {
            if ((655 > 42) && (1 >= 1)) {
                fmt.Println(v2)
                switch -36 {
                case 0:
                    switch 9223372036854775807 {
                    case -1:
                        _ = len(v8)
                    case -401:
                        switch 9223372036854775807 {
                        case 129:
                            v2 -= float32((0.1 - -2.25))
                        case 32767:
                            v3 = (("hello" + ((v3 + "a") + v3)) + "a")
                        default:
                            if true {
                                v32 := GenS4{F0: "\n\t", F1: (2147483647 != 431)}
                                v3 = v3
                                _ = v32
                            }
                        }
                    default:
                        switch "a" {
                        case "hello":
                            v2 = (v2 - v2)
                        case "a":
                            fmt.Println(v3, v2)
                        case "foo bar":
                            var v33 int64 = 763
                            _ = v33
                        default:
                            switch "foo bar" {
                            case "":
                                if true {
                                    v3 = ("\n\t" + v3)
                                }
                            case "hello":
                                var v34 bool = false
                                _ = v34
                            }
                        }
                    }
                case -82:
                    var v35 uint16 = 311
                    _ = v35
                default:
                    var v36 uint64 = 4294967295
                    _ = v36
                }
                for fi37 := 0; fi37 < 9; fi37++ {
                    v2 += v2
                    for fi38 := 0; fi38 < 18; fi38++ {
                        switch v3 {
                        case "foo bar":
                            var v39 uint8 = 35
                            _ = v39
                        default:
                            var v40 uint64 = 1
                            _ = v40
                        }
                        var v41 uint16 = 2
                        switch (^(fi38)) {
                        case 867:
                            var v42 float64 = ((0.0 - 0.1) - 1.5)
                            _ = v42
                        case 127:
                            v41 -= 230
                        }
                        _ = fi38
                        _ = v41
                    }
                    var v43 uint32 = 883
                    _ = fi37
                    _ = v43
                }
            } else {
                for fi44 := 0; fi44 < 20; fi44++ {
                    _ = len(v8)
                    _ = fi44
                }
                v45 := GenS5{F0: v3, F1: -18}
                var v46 int8 = 1
                _ = v45
                _ = v46
            }
            var v47 string = v3
            _ = v47
        } else {
            v48 := v1[141]
            var v49 int16 = 0
            _ = v48
            _ = v49
        }
        v50 := GenS1{F0: v3, F1: 842, F2: 42}
        _ = v50
    } else {
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        var v51 uint64 = 683
        _ = v51
    }
    v2 = 0.0
    if true {
        v52 := GenS1{F0: v3, F1: 303, F2: 18446744073709551615}
        if (0 >= 1) {
            var v53 int16 = 1
            v2 += v2
            _ = v53
        }
        var v54 int16 = 42
        _ = v52
        _ = v54
    }
    switch 42 {
    case 9223372036854775807:
        var v55 uint16 = 65535
        _ = v55
    case -255:
        var v56 string = ""
        _ = v56
    default:
        switch v3 {
        case "hello":
            var v57 int32 = 2147483647
            _ = v57
        case "":
            var v58 int = 9223372036854775807
            _ = v58
        case "foo bar":
            if true {
                if false {
                    delete(v1, 455)
                    v3 = v3
                } else {
                    if true {
                        if (("\n\t" + ("hello" + "foo bar")) >= (v3 + v3)) {
                            var v59 float64 = -2.25
                            switch 128 {
                            case 4294967295:
                                if len(v8) >= 2 {
                                    v8 = v8[1:len(v8)]
                                }
                            default:
                                v3 = (("foo bar" + "\n\t") + "")
                            }
                            _ = v59
                        } else {
                            fmt.Println(v2, v3)
                            var v60 bool = true
                            _ = v60
                        }
                        var v61 float64 = ((-2.25 / 1.5) / 3.14159)
                        _ = v61
                    }
                }
            } else {
                if ((v3 + v3) == "foo bar") {
                    v62 := GenS2{F0: -128}
                    _ = v62
                } else {
                    fmt.Println(v3, v3)
                    for fi63 := 0; fi63 < 16; fi63++ {
                        fmt.Println(v3)
                        var v64 uint8 = 255
                        var v65 string = v3
                        _ = fi63
                        _ = v64
                        _ = v65
                    }
                    fmt.Println(v3)
                }
            }
        default:
            v2 += v2
        }
    }
    switch 127 {
    case 0:
        if false {
            for fi66 := 0; fi66 < 1; fi66++ {
                var v67 float32 = float32(float64(v2))
                if (true || true) {
                    for fi68 := 0; fi68 < 14; fi68++ {
                        var v69 uint8 = 33
                        var v70 float64 = (0.1 / (3.14159 / 1.0))
                        v3 = v3
                        _ = fi68
                        _ = v69
                        _ = v70
                    }
                    v2 += v67
                } else {
                    delete(v1, (fi66 ^ -2147483648))
                    if (true || ((fi66 & fi66) >= fi66)) {
                        for fi71 := 0; fi71 < 7; fi71++ {
                            if false {
                                v72 := GenS6{F0: false}
                                v73 := GenS7{F0: -32768}
                                v74 := v1[(fi66 | (fi71 + fi66))]
                                if (3.14159 == 1.5) {
                                    if true {
                                        switch (^(v74)) {
                                        case 42:
                                            var v75 float64 = (((0.0 / -2.25) + 0.1) / 3.14159)
                                            _ = v75
                                        case 0:
                                            v2 = v67
                                        case 2147483647:
                                            switch (fi66 | fi71) {
                                            case -2147483648:
                                                var v76 uint8 = 0
                                                _ = v76
                                            case 128:
                                                var v77 int64 = 290
                                                _ = v77
                                            case 9223372036854775807:
                                                v78 := GenS3{F0: -2147483648, F1: 1}
                                                fmt.Println(v67, fi71, v2)
                                                v2 -= v2
                                                _ = v78
                                            }
                                        default:
                                            if len(v8) > 0 {
                                                v8[0] = (fi66 / (((fi71 | 883)) | 1))
                                            }
                                        }
                                    } else {
                                        var v79 int64 = 967
                                        _ = v79
                                    }
                                    var v80 uint8 = 1
                                    v81 := GenS8{F0: 127}
                                    _ = v80
                                    _ = v81
                                } else {
                                    var v82 uint16 = 1
                                    if (float32((-2.25 * 3.14159)) < 1e10) {
                                        fmt.Println(v82, v74)
                                        fmt.Println(v82, fi66, v3)
                                        v83 := GenS2{F0: 87}
                                        _ = v83
                                    } else {
                                        v84 := GenS9{F0: 2, F1: 216, F2: 18446744073709551615}
                                        v85 := GenS5{F0: "hello", F1: (fi66 + (fi71 >> ((2) & 63)))}
                                        _ = v84
                                        _ = v85
                                    }
                                    var v86 string = (v3 + v3)
                                    _ = v82
                                    _ = v86
                                }
                                _ = v72
                                _ = v73
                                _ = v74
                            }
                            switch (fi66 % ((799) | 1)) {
                            case 4294967295:
                                fmt.Println(fi71, v2, fi71)
                            case 288:
                                var v87 int16 = 657
                                _ = v87
                            case -650:
                                var v88 int32 = 2147483647
                                _ = v88
                            default:
                                var v89 uint = 985
                                _ = v89
                            }
                            var v90 float32 = v67
                            _ = fi71
                            _ = v90
                        }
                        if ((fi66 & fi66) > fi66) {
                            v91 := GenS10{F0: 1}
                            if true {
                                if true {
                                    v2 -= ((((v2 + 1.5) / (v2 / (v2 * (float32((-2.25 * 0.1)) + (float32((0.0 + 0.0)) + float32((0.1 * 0.0))))))) - (v67 - (v2 / -2.25))) / 1.5)
                                } else {
                                    if false {
                                        v92 := GenS11{F0: 677, F1: -85, F2: 355}
                                        if false {
                                            var v93 string = v3
                                            if false {
                                                for fi94 := 0; fi94 < 9; fi94++ {
                                                    v95 := GenS3{F0: 482, F1: fi94}
                                                    v96 := GenS3{F0: -2147483648, F1: (fi66 / ((0) | 1))}
                                                    _ = fi94
                                                    _ = v95
                                                    _ = v96
                                                }
                                                v67 = 1.5
                                                v93 = "foo bar"
                                            }
                                            var v97 int32 = -2147483648
                                            _ = v93
                                            _ = v97
                                        } else {
                                            switch 805 {
                                            case 4294967295:
                                                v98 := GenS5{F0: "", F1: (fi66 | 127)}
                                                _ = v98
                                            case 32768:
                                                v1[(fi66 / ((1) | 1))] = (-(fi66))
                                            default:
                                                for fi99 := 0; fi99 < 14; fi99++ {
                                                    if false {
                                                        var v100 bool = (3.14159 <= ((v67 / 1e10) + v2))
                                                        if (v100 && true) {
                                                            if (v100 || (-689 >= 42)) {
                                                                fmt.Println(v3, v2, v3)
                                                                if (2 >= 18446744073709551615) {
                                                                    var v101 bool = (2 < 255)
                                                                    _ = v101
                                                                } else {
                                                                    v8 = append(v8, fi99)
                                                                }
                                                            }
                                                        } else {
                                                            switch fi99 {
                                                            case 32767:
                                                                if (genFn0(-2147483648, (fi66 >> ((18446744073709551615) & 63)), 0) >= 65535) {
                                                                    v3 = ("foo bar" + "a")
                                                                } else {
                                                                    var v102 int64 = -9223372036854775808
                                                                    delete(v1, (-(fi66)))
                                                                    if (15 <= 42) {
                                                                        switch fi99 {
                                                                        case -9223372036854775808:
                                                                            var v103 string = ""
                                                                            _ = v103
                                                                        case 127:
                                                                            var v104 uint64 = 42
                                                                            _ = v104
                                                                        }
                                                                    } else {
                                                                        var v105 int16 = -471
                                                                        var v106 int32 = -978
                                                                        fmt.Println(v100, v106)
                                                                        if (1 > 2) {
                                                                            var v107 bool = (0 < 2)
                                                                            _ = v107
                                                                        }
                                                                        v105 += v105
                                                                        _ = v105
                                                                        _ = v106
                                                                    }
                                                                    _ = v102
                                                                }
                                                            }
                                                            var v108 int32 = -648
                                                            _ = v108
                                                        }
                                                        v2 = 1.5
                                                        _ = v100
                                                    }
                                                    _ = fi99
                                                }
                                            }
                                            var v109 int32 = 247
                                            _ = v109
                                        }
                                        _ = v92
                                    }
                                }
                                for fi110 := 0; fi110 < 20; fi110++ {
                                    switch fi110 {
                                    case 489:
                                        v3 = (v3 + v3)
                                    }
                                    switch 203 {
                                    case 42:
                                        v1[(fi110 << ((65535) & 63))] = 4294967295
                                    case 2147483647:
                                        fmt.Println(fi66)
                                        if ((4294967295 == 4294967295) || (1.5 > 1e10)) {
                                            var v111 string = (("hello" + v3) + (v3 + v3))
                                            if false {
                                                var v112 uint8 = 128
                                                if (42 >= 0) {
                                                    fmt.Println(fi110, v112)
                                                    v2 += 1.5
                                                } else {
                                                    if ((42 < 85) || false) {
                                                        v67 += 1.5
                                                        v113 := v1[(fi110 * (fi66 & (fi66 & -9223372036854775808)))]
                                                        var v114 int32 = 381
                                                        _ = v113
                                                        _ = v114
                                                    } else {
                                                        if (true && true) {
                                                            var v115 float32 = (v2 + -2.25)
                                                            var v116 float64 = 0.0
                                                            _ = v115
                                                            _ = v116
                                                        } else {
                                                            v3 = ""
                                                            fmt.Println(v2, v112, fi110)
                                                            _ = len(v1)
                                                        }
                                                        _ = len(v8)
                                                    }
                                                    var v117 int64 = 585
                                                    if len(v8) > 0 {
                                                        v8[0] = fi110
                                                    }
                                                    _ = v117
                                                }
                                                _ = v112
                                            } else {
                                                var v118 int = 42
                                                if true {
                                                    v1[v118] = (fi110 % (((fi66 >> ((genFn0(684, (v118 * 127), 4294967295)) & 63))) | 1))
                                                }
                                                _ = v118
                                            }
                                            if false {
                                                switch (("foo bar" + v3) + "hello") {
                                                case "\n\t":
                                                    v119 := GenS9{F0: 793, F1: 2147483647, F2: 69}
                                                    _ = v119
                                                case "":
                                                    switch (^(fi66)) {
                                                    case -9223372036854775808:
                                                        switch (v111 + "foo bar") {
                                                        case "foo bar":
                                                            v111 = v111
                                                        default:
                                                            var v120 bool = (!false)
                                                            _ = v120
                                                        }
                                                    case 128:
                                                        if false {
                                                            switch fi110 {
                                                            case 32768:
                                                                var v121 uint8 = 255
                                                                _ = v121
                                                            case 32767:
                                                                v122 := GenS6{F0: true}
                                                                _ = v122
                                                            case -1:
                                                                v123 := GenS1{F0: "foo bar", F1: 753, F2: 0}
                                                                _ = v123
                                                            default:
                                                                fmt.Println(v67)
                                                            }
                                                            if false {
                                                                fmt.Println(fi66, fi66)
                                                            } else {
                                                                var v124 int16 = 42
                                                                fmt.Println(v111, v3, v111)
                                                                var v125 int16 = -32768
                                                                _ = v124
                                                                _ = v125
                                                            }
                                                        }
                                                    default:
                                                        switch ("foo bar" + ((v111 + (("\n\t" + "a") + v3)) + v3)) {
                                                        case "":
                                                            v126 := GenS10{F0: 0}
                                                            _ = v126
                                                        }
                                                    }
                                                    v111 = v111
                                                    v2 = v2
                                                case "a":
                                                    if true {
                                                        v67 -= v67
                                                        v127 := GenS12{F0: -32768, F1: 319}
                                                        _ = v127
                                                    }
                                                }
                                                v111 = v111
                                            } else {
                                                var v128 uint64 = 396
                                                var v129 float64 = float64(0.0)
                                                var v130 bool = (((^(v128)) > 443) || true)
                                                _ = v128
                                                _ = v129
                                                _ = v130
                                            }
                                            _ = v111
                                        }
                                        var v131 uint8 = 222
                                        _ = v131
                                    case 796:
                                        v8 = append(v8, (fi110 ^ fi110))
                                    }
                                    var v132 uint16 = 0
                                    _ = fi110
                                    _ = v132
                                }
                            }
                            _ = v91
                        }
                        v2 = 3.14159
                    } else {
                        v133 := v1[fi66]
                        _ = len(v8)
                        if (540 > 790) {
                            v2 += 3.14159
                            v134 := GenS10{F0: 988}
                            _ = v134
                        } else {
                            var v135 uint = 850
                            v136 := GenS6{F0: false}
                            for fi137 := 0; fi137 < 18; fi137++ {
                                fmt.Println(v67, v135)
                                v138 := GenS12{F0: 42, F1: 0}
                                _ = fi137
                                _ = v138
                            }
                            _ = v135
                            _ = v136
                        }
                        _ = v133
                    }
                }
                _ = fi66
                _ = v67
            }
            var v139 uint32 = 2
            _ = v139
        } else {
            var v140 bool = true
            switch 478 {
            case 4294967295:
                for fi141 := 0; fi141 < 20; fi141++ {
                    switch 128 {
                    case -54:
                        var v142 int16 = 0
                        _ = v142
                    case 471:
                        v8 = append(v8, (-(fi141)))
                    case 706:
                        v3 = v3
                    default:
                        for fi143 := 0; fi143 < 1; fi143++ {
                            v3 = v3
                            _ = fi143
                        }
                    }
                    var v144 int8 = -78
                    _ = fi141
                    _ = v144
                }
            case 54:
                for fi145 := 0; fi145 < 2; fi145++ {
                    switch fi145 {
                    case -9223372036854775808:
                        if (float64((v2 * 0.1)) >= 0.1) {
                            v3 = v3
                        }
                    }
                    v146 := GenS11{F0: 2147483648, F1: -38, F2: 1}
                    _ = fi145
                    _ = v146
                }
            }
            for fi147 := 0; fi147 < 20; fi147++ {
                v2 = v2
                _ = fi147
            }
            _ = v140
        }
    case 42:
        for fi148 := 0; fi148 < 8; fi148++ {
            if len(v8) > 0 {
                v8[0] = fi148
            }
            _ = fi148
        }
    default:
        switch 127 {
        case 673:
            var v149 bool = (-2.25 != float64(v2))
            _ = v149
        }
    }
    v8 = append(v8, 1)
    switch 127 {
    case 564:
        switch "a" {
        case "hello":
            switch ("hello" + v3) {
            case "hello":
                v150 := GenS6{F0: false}
                _ = v150
            case "foo bar":
                var v151 int64 = 1
                _ = v151
            }
        case "":
            v2 = 1.5
        default:
            switch -9223372036854775808 {
            case -2147483648:
                if true {
                    var v152 bool = (((1.5 * v2) + v2) != ((v2 - float32(-2.25)) + ((1e10 * v2) / 1.0)))
                    var v153 int64 = 563
                    _ = v152
                    _ = v153
                }
            case 1:
                var v154 int8 = 42
                _ = v154
            }
        }
    case 32768:
        v155 := GenS3{F0: -1, F1: 0}
        _ = v155
    case -963:
        v3 = v3
    }
    fmt.Println(v3, v3)
    var v156 int = 32768
    switch (v156 ^ v156) {
    case 32768:
        switch v3 {
        case "hello":
            for fi157 := 0; fi157 < 15; fi157++ {
                v158 := GenS13{F0: 988, F1: -1}
                _ = fi157
                _ = v158
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v156:", v156)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v156
    _ = v2
    _ = v22
    _ = v3
    _ = v8
}
