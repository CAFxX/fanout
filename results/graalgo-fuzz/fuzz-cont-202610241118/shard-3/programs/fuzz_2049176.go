package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 int; F1 int32; F2 bool }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 int8; F1 string }
type GenS5 struct { F0 int32; F1 uint8 }
type GenS6 struct { F0 uint; F1 bool; F2 uint }
type GenS7 struct { F0 uint32; F1 int16; F2 int64 }
type GenS8 struct { F0 int; F1 uint16 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint; F1 int32 }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 uint16; F1 bool; F2 uint }
type GenS14 struct { F0 bool; F1 uint32; F2 uint64 }
type GenS15 struct { F0 uint; F1 int8 }
func genFn0(p0 int64) uint64 {
    var v7 int = 9223372036854775807
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 float64 = 0.0
        fmt.Println(p0, v9, v7)
        _ = fi8
        _ = v9
    }
    fmt.Println(p0, v7)
    _ = p0
    _ = v7
    return 143
}

func main() {
    var v1 int32 = -1
    v2 := make(map[string]string)
    var v3 float32 = (-2.25 + (0.1 / (0.1 - (3.14159 / 3.14159))))
    var v4 uint64 = 872
    var v5 float64 = 1.5
    var v6 int64 = 196
    if (((v5 / 0.1) < 0.1) || ("" == ((("\n\t" + "\n\t") + "hello") + ""))) {
        if true {
            v4 = v4
        }
    } else {
        v10 := []int{32768, 0}
        _ = v10
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        if (127 < 1) {
            if true {
                switch 9223372036854775807 {
                case 32767:
                    v12 := []int{(fi11 >> ((65535) & 63)), fi11, fi11, fi11}
                    _ = v12
                }
                v13 := GenS1{F0: 2147483648, F1: (v6 * v6)}
                for fi14 := 0; fi14 < 5; fi14++ {
                    v15 := GenS1{F0: 2147483648, F1: -1}
                    _ = fi14
                    _ = v15
                }
                _ = v13
            } else {
                v6 = (-(v6))
                v16 := GenS1{F0: 0, F1: -1}
                if (337 > 65535) {
                    if (fi11 != fi11) {
                        var v17 float64 = 0.1
                        _ = v17
                    }
                    for fi18 := 0; fi18 < 17; fi18++ {
                        if ((v1 + -139) < v1) {
                            if false {
                                v19 := []int{(fi11 / ((0) | 1)), fi18, fi18, (fi18 >> ((543) & 63))}
                                v4 += (v4 / ((1) | 1))
                                var v20 int8 = -1
                                _ = v19
                                _ = v20
                            }
                        }
                        v5 += (v5 + v5)
                        switch fi11 {
                        case 32767:
                            var v21 uint8 = 42
                            _ = v21
                        case 9223372036854775807:
                            var v22 int64 = (v6 - 292)
                            _ = v22
                        }
                        _ = fi18
                    }
                }
                _ = v16
            }
            v23 := v2[("\n\t" + "a")]
            _ = v23
        }
        var v24 int32 = v1
        var v25 int32 = (v1 % ((-1) | 1))
        _ = fi11
        _ = v24
        _ = v25
    }
    v26 := GenS1{F0: 65535, F1: 0}
    fmt.Println(v6, v6)
    v5 += v5
    v27 := v2[("a" + "\n\t")]
    switch 127 {
    case -297:
        for fi28 := 0; fi28 < 15; fi28++ {
            var v29 uint64 = v4
            _ = fi28
            _ = v29
        }
    case -9223372036854775808:
        var v30 int32 = v1
        _ = v30
    case 127:
        if false {
            switch v27 {
            case "":
                v27 = ""
            case "foo bar":
                v31 := GenS1{F0: 42, F1: (v6 | (v6 | (v6 << ((4294967295) & 63))))}
                _ = v31
            case "hello":
                if (!(-2147483648 > -912)) {
                    v5 = 3.14159
                    for fi32 := 0; fi32 < 15; fi32++ {
                        v3 -= 0.1
                        switch fi32 {
                        case 127:
                            v33 := GenS1{F0: 2147483647, F1: v6}
                            _ = v33
                        case 42:
                            var v34 float64 = (v5 / v5)
                            _ = v34
                        case -39:
                            v5 = -2.25
                        default:
                            v35 := []int{128, -853}
                            _ = v35
                        }
                        _ = fi32
                    }
                } else {
                    for fi36 := 0; fi36 < 15; fi36++ {
                        v37 := []int{127, (fi36 << ((255) & 63)), (fi36 / (((fi36 ^ 32767)) | 1)), (fi36 / ((fi36) | 1))}
                        if (!(!false)) {
                            var v38 string = ("" + ("foo bar" + v27))
                            v39 := v2[("" + ("foo bar" + "foo bar"))]
                            var v40 uint64 = (v4 & 2)
                            _ = v38
                            _ = v39
                            _ = v40
                        } else {
                            fmt.Println(v1, v5, v6)
                            var v41 float64 = v5
                            if ((v6 & v6) < 9223372036854775807) {
                                for fi42 := 0; fi42 < 4; fi42++ {
                                    if true {
                                        if ((false || false) || true) {
                                            v6 = 42
                                            v41 = float64(v3)
                                            v1 -= (v1 + v1)
                                        }
                                    } else {
                                        var v43 uint = 4294967295
                                        var v44 int = fi36
                                        v6 += -9223372036854775808
                                        _ = v43
                                        _ = v44
                                    }
                                    _ = fi42
                                }
                                var v45 int64 = (v6 * 905)
                                _ = v45
                            } else {
                                var v46 float32 = -2.25
                                for fi47 := 0; fi47 < 3; fi47++ {
                                    if (32767 != -590) {
                                        if false {
                                            v48 := GenS2{F0: fi36, F1: (v1 | (v1 + (v1 / ((v1) | 1)))), F2: true}
                                            if (!false) {
                                                var v49 uint = 255
                                                _ = v49
                                            }
                                            _ = v48
                                        }
                                        if true {
                                            v41 = v5
                                        }
                                        v1 -= v1
                                    }
                                    var v50 int32 = v1
                                    _ = fi47
                                    _ = v50
                                }
                                var v51 int = 707
                                _ = v46
                                _ = v51
                            }
                            _ = v41
                        }
                        switch ((((v27 + v27) + "\n\t") + "a") + v27) {
                        case "":
                            for fi52 := 0; fi52 < 8; fi52++ {
                                if (255 <= 0) {
                                    v53 := GenS2{F0: (-(fi36)), F1: (v1 + -897), F2: ((v27 + (v27 + ((("\n\t" + v27) + ("foo bar" + v27)) + ("" + ("\n\t" + ((v27 + v27) + (v27 + (("foo bar" + v27) + v27)))))))) > ((("hello" + v27) + ("hello" + v27)) + v27))}
                                    var v54 int32 = v1
                                    _ = v53
                                    _ = v54
                                }
                                if (!(0.0 != float64((v3 * 1.5)))) {
                                    if (-128 >= -67) {
                                        switch "\n\t" {
                                        case "\n\t":
                                            var v55 uint = 4294967295
                                            _ = v55
                                        default:
                                            v6 -= 0
                                        }
                                    } else {
                                        v56 := GenS2{F0: (fi52 * 9223372036854775807), F1: 2147483647, F2: false}
                                        v1 -= int32(v4)
                                        _ = v56
                                    }
                                    _ = len(v2)
                                } else {
                                    if len(v37) > 0 {
                                        v37[0] = (fi36 & (fi52 | fi36))
                                    }
                                    if (!(4294967295 == 326)) {
                                        var v57 int64 = 9223372036854775807
                                        v58 := v2["foo bar"]
                                        v27 = "hello"
                                        _ = v57
                                        _ = v58
                                    }
                                    switch fi36 {
                                    case -23:
                                        v3 -= (float32(0.0) + 0.1)
                                    case 32767:
                                        var v59 string = v27
                                        _ = v59
                                    default:
                                        var v60 uint64 = 1
                                        _ = v60
                                    }
                                }
                                _ = fi52
                            }
                        default:
                            fmt.Println(v3, v5, v4)
                        }
                        _ = fi36
                        _ = v37
                    }
                }
            default:
                v61 := []int{}
                _ = v61
            }
        }
    }
    v62 := GenS3{F0: -988, F1: -47}
    switch 128 {
    case 127:
        if (v4 != (v4 / ((genFn0(9223372036854775807)) | 1))) {
            if ((((float32((v5 + 1e10)) / v3) / ((((float32(3.14159) / ((v3 / float32(v5)) - v3)) + v3) / (v3 - v3)) / 0.1)) != (v3 / v3)) || false) {
                v63 := []int{-162}
                v4 = v4
                _ = v63
            }
            v1 = v1
            v27 = ""
        } else {
            v64 := GenS4{F0: -128, F1: (v27 + (v27 + ((v27 + "a") + (v27 + "\n\t"))))}
            fmt.Println(v27)
            _ = v64
        }
    case 42:
        if false {
            v4 += 1
            v65 := GenS2{F0: -1, F1: v1, F2: false}
            v27 = "a"
            _ = v65
        }
    case -302:
        var v66 int16 = 32767
        _ = v66
    default:
        switch v27 {
        case "foo bar":
            var v67 int64 = -9223372036854775808
            _ = v67
        case "a":
            for fi68 := 0; fi68 < 1; fi68++ {
                var v69 int = fi68
                v70 := []int{-944, fi68, fi68}
                var v71 uint16 = uint16(v4)
                _ = fi68
                _ = v69
                _ = v70
                _ = v71
            }
        }
    }
    for fi72 := 0; fi72 < 19; fi72++ {
        for fi73 := 0; fi73 < 6; fi73++ {
            fmt.Println(fi72, v27, v1)
            if (2 >= 2) {
                if false {
                    fmt.Println(v1)
                    if (v4 == v4) {
                        v74 := []int{fi73, fi73, (fi73 % ((-1) | 1)), (fi73 << ((255) & 63))}
                        _ = v74
                    } else {
                        if (!(true || (fi72 >= -1))) {
                            if (32767 != 733) {
                                v27 = "foo bar"
                                if false {
                                    if ((v1 + v1) > 88) {
                                        fmt.Println(v5, v5, v27)
                                        switch "\n\t" {
                                        case "\n\t":
                                            var v75 uint8 = 127
                                            _ = v75
                                        case "a":
                                            v27 = v27
                                        case "":
                                            var v76 uint32 = 358
                                            _ = v76
                                        }
                                        v77 := []int{}
                                        _ = v77
                                    }
                                } else {
                                    v3 = v3
                                }
                                v3 = v3
                            } else {
                                switch (fi72 * 42) {
                                case 116:
                                    v78 := v2[v27]
                                    _ = v78
                                case -297:
                                    if (true && (255 != 82)) {
                                        if false {
                                            fmt.Println(v27)
                                            v4 = v4
                                            switch v27 {
                                            case "hello":
                                                if ("hello" <= v27) {
                                                    v4 += 556
                                                }
                                            default:
                                                var v79 bool = false
                                                _ = v79
                                            }
                                            var v80 int64 = int64(v1)
                                            switch ((v27 + (("" + ("foo bar" + v27)) + ("" + (((("foo bar" + "hello") + (("" + v27) + v27)) + "") + v27)))) + ((v27 + "foo bar") + (v27 + v27))) {
                                            case "foo bar":
                                                v1 -= (^(v1))
                                            case "\n\t":
                                                v80 = (v80 | v6)
                                            case "a":
                                                if false {
                                                    v80 -= v6
                                                }
                                            }
                                            var v81 int64 = (-(v80))
                                            _ = v80
                                            _ = v81
                                        }
                                        v82 := []int{(fi72 - (-(fi72)))}
                                        delete(v2, ((v27 + v27) + v27))
                                        _ = v82
                                    } else {
                                        var v83 int = fi72
                                        v83 += (v83 ^ (v83 + 1))
                                        fmt.Println(v1, v83, v4)
                                        v84 := GenS3{F0: 128, F1: 1}
                                        _ = v83
                                        _ = v84
                                    }
                                }
                                var v85 uint32 = 617
                                var v86 uint8 = 128
                                var v87 int16 = 0
                                var v88 uint64 = (v4 | v4)
                                _ = v85
                                _ = v86
                                _ = v87
                                _ = v88
                            }
                            v5 = (3.14159 - (v5 * v5))
                            v89 := GenS4{F0: -30, F1: "\n\t"}
                            _ = v89
                        }
                    }
                    if true {
                        v1 = v1
                        v6 = (v6 % ((v6) | 1))
                        var v90 bool = ((v4 + v4) < (v4 - 291))
                        _ = v90
                    }
                } else {
                    v91 := GenS5{F0: (v1 ^ (-(v1))), F1: 128}
                    v6 -= v6
                    v92 := GenS4{F0: 100, F1: "hello"}
                    _ = v91
                    _ = v92
                }
                v27 = "a"
                if ((!false) || true) {
                    v93 := GenS6{F0: 970, F1: true, F2: 342}
                    v4 = v4
                    var v94 int16 = int16(v4)
                    _ = v93
                    _ = v94
                }
            }
            if false {
                fmt.Println(v1, v6, fi72)
                var v95 uint64 = 4294967295
                _ = v95
            } else {
                if (true && ((!(!false)) && ((v4 - genFn0(v6)) <= (v4 / (((v4 - (^(v4)))) | 1))))) {
                    if true {
                        v96 := GenS7{F0: 867, F1: 0, F2: v6}
                        if (!false) {
                            var v97 bool = (42 > 0)
                            v98 := []int{fi73, -9223372036854775808, fi73, fi73}
                            v1 = (v1 + v1)
                            _ = v97
                            _ = v98
                        }
                        _ = v96
                    } else {
                        var v99 uint8 = 124
                        _ = v99
                    }
                }
                var v100 uint16 = 0
                if true {
                    var v101 bool = ("foo bar" == v27)
                    _ = v101
                } else {
                    var v102 bool = (461 <= v100)
                    switch v27 {
                    case "\n\t":
                        fmt.Println(v102, v102)
                    default:
                        if v102 {
                            if (1 > 4294967295) {
                                if ((0 < -755) || v102) {
                                    v103 := GenS8{F0: fi72, F1: 950}
                                    _ = v103
                                } else {
                                    var v104 string = (v27 + v27)
                                    var v105 uint16 = 32767
                                    var v106 int32 = (v1 ^ (v1 + (v1 % ((2147483647) | 1))))
                                    _ = v104
                                    _ = v105
                                    _ = v106
                                }
                                v107 := GenS5{F0: -1, F1: 200}
                                _ = v107
                            } else {
                                if (42 > int64(v4)) {
                                    v6 = -1
                                    v5 = v5
                                }
                                fmt.Println(v1, v1)
                                fmt.Println(v6, v5, v100)
                            }
                        } else {
                            var v108 uint16 = (v100 * (v100 ^ v100))
                            var v109 int = fi72
                            if v102 {
                                var v110 string = "hello"
                                var v111 float32 = (v3 + 1.5)
                                v27 = v27
                                _ = v110
                                _ = v111
                            } else {
                                if (!(!v102)) {
                                    v112 := []int{}
                                    _ = len(v2)
                                    _ = v112
                                }
                                fmt.Println(fi72)
                                switch (fi72 % ((fi73) | 1)) {
                                case -495:
                                    var v113 float64 = v5
                                    _ = v113
                                case -9223372036854775808:
                                    fmt.Println(v5, fi72, v6)
                                default:
                                    _ = len(v2)
                                }
                            }
                            _ = v108
                            _ = v109
                        }
                    }
                    _ = v102
                }
                _ = v100
            }
            v3 = (0.1 - (v3 + ((-2.25 * 0.1) - 3.14159)))
            switch v27 {
            case "a":
                if true {
                    fmt.Println(fi73, fi72, fi73)
                    v1 = v1
                }
            case "\n\t":
                var v114 uint32 = 2147483648
                _ = v114
            case "foo bar":
                switch "" {
                case "hello":
                    if (!true) {
                        var v115 uint8 = 245
                        _ = v115
                    }
                }
            default:
                v116 := GenS4{F0: 0, F1: ("" + (v27 + ""))}
                fmt.Println(v1, v6, v27)
                _ = v116
            }
            v3 -= 0.1
            _ = fi73
        }
        _ = fi72
    }
    v117 := GenS9{F0: (v4 | (^(v4)))}
    switch 9223372036854775807 {
    case 220:
        v118 := GenS6{F0: 1, F1: ((float32(float64(v3)) + float32(v5)) == v3), F2: 4294967295}
        _ = v118
    case -68:
        for fi119 := 0; fi119 < 17; fi119++ {
            switch (v27 + (v27 + (((v27 + v27) + v27) + "a"))) {
            case "\n\t":
                var v120 uint64 = v4
                _ = v120
            }
            if true {
                v6 += v6
            } else {
                v121 := GenS3{F0: fi119, F1: 107}
                _ = v121
            }
            _ = fi119
        }
    case 985:
        var v122 int16 = -595
        _ = v122
    }
    v27 = v27
    for fi123 := 0; fi123 < 9; fi123++ {
        v3 += ((v3 - -2.25) * 0.1)
        fmt.Println(v3)
        _ = fi123
    }
    var v124 float64 = v5
    switch v27 {
    case "a":
        if false {
            fmt.Println(v6, v1)
        } else {
            if false {
                var v125 int16 = 32767
                _ = v125
            } else {
                v1 = (v1 + v1)
            }
            v126 := GenS10{F0: 4294967295, F1: v1}
            for fi127 := 0; fi127 < 4; fi127++ {
                v5 = 1.5
                if false {
                    var v128 uint64 = 18446744073709551615
                    var v129 uint = 1
                    _ = v128
                    _ = v129
                } else {
                    var v130 string = (v27 + v27)
                    if (0 != 32767) {
                        var v131 float64 = (((((v124 - v124) * (float64(1.5) * 3.14159)) + v124) / ((0.0 / 1.5) * (3.14159 + (v5 * v5)))) / 3.14159)
                        v132 := GenS11{F0: -128}
                        var v133 uint32 = 791
                        _ = v131
                        _ = v132
                        _ = v133
                    } else {
                        switch (fi127 ^ (-(fi127))) {
                        case 818:
                            switch "" {
                            case "\n\t":
                                v134 := GenS12{F0: 127}
                                _ = v134
                            case "hello":
                                switch v130 {
                                case "a":
                                    v5 += v124
                                default:
                                    if true {
                                        for fi135 := 0; fi135 < 8; fi135++ {
                                            v130 = "a"
                                            switch "\n\t" {
                                            case "a":
                                                fmt.Println(v124)
                                            case "":
                                                v136 := GenS13{F0: 0, F1: (!false), F2: 255}
                                                _ = v136
                                            case "\n\t":
                                                var v137 bool = false
                                                _ = v137
                                            }
                                            _ = fi135
                                        }
                                        var v138 uint32 = 883
                                        _ = v138
                                    }
                                }
                            default:
                                var v139 int64 = v6
                                _ = v139
                            }
                        case 2147483647:
                            v140 := GenS14{F0: (!false), F1: 35, F2: v4}
                            _ = v140
                        }
                        v3 = ((v3 - v3) / ((v3 * (0.1 * (((v3 + v3) + v3) * v3))) + float32(v5)))
                    }
                    for fi141 := 0; fi141 < 18; fi141++ {
                        v142 := GenS15{F0: 1, F1: 1}
                        var v143 bool = false
                        _ = fi141
                        _ = v142
                        _ = v143
                    }
                    _ = v130
                }
                switch (^(fi127)) {
                case 42:
                    var v144 uint32 = 4294967295
                    _ = v144
                case 32767:
                    fmt.Println(v124, fi127, fi127)
                case 4294967295:
                    v4 = 0
                default:
                    if true {
                        var v145 int16 = -32768
                        var v146 int64 = v6
                        switch (-(fi127)) {
                        case 32768:
                            v3 = (v3 + ((float32(0.0) + v3) + (v3 + v3)))
                        }
                        _ = v145
                        _ = v146
                    } else {
                        fmt.Println(v4)
                        delete(v2, v27)
                        for fi147 := 0; fi147 < 15; fi147++ {
                            fmt.Println(fi127)
                            if false {
                                var v148 uint32 = 1
                                v124 = -2.25
                                var v149 int = (fi147 >> ((65535) & 63))
                                _ = v148
                                _ = v149
                            }
                            v2[((v27 + v27) + v27)] = "foo bar"
                            _ = fi147
                        }
                    }
                }
                _ = fi127
            }
            _ = v126
        }
    case "foo bar":
        if (false && true) {
            v2["a"] = (v27 + v27)
            delete(v2, v27)
        }
    case "\n\t":
        switch (v27 + ("\n\t" + (("foo bar" + "") + v27))) {
        case "hello":
            var v150 int64 = v6
            _ = v150
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v124:", v124)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v117
    _ = v124
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v62
}
