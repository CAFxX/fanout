package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint32; F1 string }
type GenS8 struct { F0 uint16; F1 uint; F2 uint }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 uint32; F1 int64; F2 int32 }
type GenS11 struct { F0 uint16; F1 int16; F2 int8 }
type GenS12 struct { F0 uint32; F1 uint16; F2 int }
type GenS13 struct { F0 uint; F1 int32; F2 int32 }
type GenS14 struct { F0 int; F1 int32 }
type GenS15 struct { F0 int8; F1 uint64; F2 uint16 }
type GenS16 struct { F0 int16; F1 int8 }
func genFn0(p0 uint64, p1 int64) uint8 {
    p0 = (^(p0))
    for fi6 := 0; fi6 < 12; fi6++ {
        for fi7 := 0; fi7 < 4; fi7++ {
            var v8 uint8 = 244
            v9 := GenS1{F0: (48 > 28), F1: p1}
            _ = fi7
            _ = v8
            _ = v9
        }
        switch (fi6 - fi6) {
        case 0:
            switch (^(fi6)) {
            case 128:
                var v10 uint16 = 822
                _ = v10
            default:
                v11 := GenS1{F0: (!(fi6 <= fi6)), F1: 0}
                _ = v11
            }
        case 42:
            for fi12 := 0; fi12 < 20; fi12++ {
                v13 := []int{(fi12 * (fi12 % ((fi6) | 1)))}
                var v14 int64 = -1
                v15 := GenS1{F0: (true || (false && false)), F1: v14}
                _ = fi12
                _ = v13
                _ = v14
                _ = v15
            }
        case 4294967295:
            fmt.Println(p0)
        }
        p1 = (p1 % ((p1) | 1))
        _ = fi6
    }
    switch "\n\t" {
    case "":
        p1 = (p1 ^ p1)
    case "a":
        v16 := GenS2{F0: 32767, F1: 2}
        _ = v16
    case "foo bar":
        v17 := GenS3{F0: 482, F1: (65535 <= 42)}
        _ = v17
    }
    _ = p0
    _ = p1
    return 127
}

func main() {
    var v1 float64 = (0.0 - ((3.14159 * float64((-2.25 - 0.1))) / 3.14159))
    var v2 int32 = 0
    var v3 uint32 = 983
    var v4 bool = true
    var v5 int = 83
    v3 += 65535
    if v4 {
        switch (v5 & v5) {
        case 371:
            for fi18 := 0; fi18 < 11; fi18++ {
                v19 := GenS3{F0: 128, F1: (v4 || v4)}
                switch -1 {
                case 132:
                    var v20 int16 = -32768
                    _ = v20
                case 42:
                    if (110 < 128) {
                        v21 := GenS4{F0: -9223372036854775808}
                        _ = v21
                    } else {
                        v22 := GenS1{F0: (-1 > v2), F1: -1}
                        v23 := GenS2{F0: v5, F1: 2}
                        v24 := GenS2{F0: 42, F1: 558}
                        _ = v22
                        _ = v23
                        _ = v24
                    }
                default:
                    fmt.Println(fi18)
                }
                _ = fi18
                _ = v19
            }
        default:
            fmt.Println(v5, v2)
        }
    } else {
        v25 := []int{(v5 >> ((65535) & 63)), int(v3)}
        v5 += v5
        _ = v25
    }
    v26 := []int{4294967295}
    v27 := GenS5{F0: (v2 - v2)}
    fmt.Println(v1, v1, v5)
    switch (v5 % ((v5) | 1)) {
    case 0:
        for fi28 := 0; fi28 < 19; fi28++ {
            if len(v26) >= 2 {
                v26 = v26[1:len(v26)]
            }
            for fi29 := 0; fi29 < 10; fi29++ {
                v2 = 1
                _ = fi29
            }
            _ = fi28
        }
    }
    for fi30 := 0; fi30 < 12; fi30++ {
        v26 = append(v26, v5)
        if v4 {
            var v31 uint = 255
            for fi32 := 0; fi32 < 13; fi32++ {
                if len(v26) > 0 {
                    v26[0] = int(v2)
                }
                if v4 {
                    fmt.Println(v3, v31)
                    v33 := GenS5{F0: v2}
                    fmt.Println(fi32, fi30)
                    _ = v33
                }
                if v4 {
                    var v34 uint = v31
                    _ = v34
                } else {
                    v35 := GenS6{F0: 89}
                    v26 = append(v26, 0)
                    if (!v4) {
                        v36 := GenS2{F0: (fi32 / (((v5 | (-(fi32)))) | 1)), F1: 713}
                        v2 = v2
                        if len(v26) > 0 {
                            v26[0] = 403
                        }
                        _ = v36
                    } else {
                        if false {
                            var v37 int16 = -32768
                            _ = v37
                        } else {
                            fmt.Println(fi30)
                            var v38 uint32 = v3
                            if (42 <= 255) {
                                if len(v26) >= 2 {
                                    v26 = v26[1:len(v26)]
                                }
                                var v39 uint = v31
                                fmt.Println(v39)
                                _ = v39
                            } else {
                                if (false || v4) {
                                    fmt.Println(v3, v5, v3)
                                }
                                var v40 uint8 = 1
                                fmt.Println(v4, v4)
                                if len(v26) > 0 {
                                    v26[0] = fi30
                                }
                                var v41 uint32 = (^(v38))
                                if v4 {
                                    var v42 float32 = (float32(v1) * 1.5)
                                    if len(v26) >= 2 {
                                        v26 = v26[1:len(v26)]
                                    }
                                    switch "foo bar" {
                                    case "a":
                                        var v43 uint32 = v3
                                        _ = v43
                                    case "":
                                        v44 := GenS7{F0: uint32(v31), F1: "foo bar"}
                                        _ = v44
                                    default:
                                        if len(v26) > 0 {
                                            v26[0] = fi32
                                        }
                                    }
                                    _ = v42
                                }
                                _ = v40
                                _ = v41
                            }
                            _ = v38
                        }
                        fmt.Println(v3, v3, v4)
                        v45 := GenS5{F0: -2147483648}
                        _ = v45
                    }
                    switch "hello" {
                    case "\n\t":
                        if len(v26) >= 2 {
                            v26 = v26[1:len(v26)]
                        }
                    case "a":
                        var v46 uint8 = 167
                        _ = v46
                    default:
                        switch "" {
                        case "hello":
                            v3 -= v3
                        case "":
                            var v47 string = "\n\t"
                            _ = v47
                        default:
                            if ((^(v2)) == -2147483648) {
                                if (fi32 > int(v31)) {
                                    var v48 uint8 = uint8(v3)
                                    if len(v26) > 0 {
                                        v26[0] = fi32
                                    }
                                    var v49 uint32 = v3
                                    var v50 uint = uint(v48)
                                    v3 += v3
                                    _ = v48
                                    _ = v49
                                    _ = v50
                                }
                                _ = len(v26)
                            }
                            var v51 bool = true
                            _ = v51
                        }
                    }
                    _ = v35
                }
                _ = fi32
            }
            v4 = (!(1 < 42))
            _ = v31
        } else {
            for fi52 := 0; fi52 < 11; fi52++ {
                if len(v26) > 0 {
                    v26[0] = fi52
                }
                v3 -= 42
                _ = fi52
            }
        }
        _ = fi30
    }
    if (false || (false || v4)) {
        fmt.Println(v4, v5)
        v53 := GenS8{F0: 255, F1: 65535, F2: 1}
        fmt.Println(v5, v5, v1)
        _ = v53
    }
    if (255 == 388) {
        if (v4 && v4) {
            for fi54 := 0; fi54 < 18; fi54++ {
                switch (^(v5)) {
                case 2147483647:
                    for fi55 := 0; fi55 < 11; fi55++ {
                        var v56 uint = 255
                        if len(v26) >= 2 {
                            v26 = v26[1:len(v26)]
                        }
                        v1 = 0.0
                        _ = fi55
                        _ = v56
                    }
                }
                for fi57 := 0; fi57 < 17; fi57++ {
                    v1 += (v1 - (v1 * v1))
                    v2 += (v2 - (v2 >> ((uint(v3)) & 63)))
                    _ = fi57
                }
                if v4 {
                    v26 = append(v26, v5)
                    fmt.Println(v5)
                    switch v5 {
                    case 9223372036854775807:
                        var v58 uint64 = 721
                        _ = v58
                    case -362:
                        if ((int16(v2) >= 1000) || (v4 || (((!v4) && (!v4)) || v4))) {
                            v59 := GenS9{F0: 101}
                            var v60 int = (fi54 * 0)
                            var v61 uint16 = 65535
                            _ = v59
                            _ = v60
                            _ = v61
                        }
                    case -2147483648:
                        for fi62 := 0; fi62 < 8; fi62++ {
                            var v63 int8 = -97
                            var v64 float64 = (((v1 * 1.5) + (v1 - v1)) / v1)
                            if ((v5 / (((fi62 & fi54)) | 1)) != -481) {
                                var v65 uint = 18446744073709551615
                                var v66 uint8 = 20
                                _ = v65
                                _ = v66
                            }
                            _ = fi62
                            _ = v63
                            _ = v64
                        }
                    default:
                        _ = len(v26)
                    }
                }
                _ = fi54
            }
        } else {
            fmt.Println(v1, v5, v5)
            switch v5 {
            case -109:
                v1 -= v1
            case 769:
                for fi67 := 0; fi67 < 1; fi67++ {
                    for fi68 := 0; fi68 < 6; fi68++ {
                        var v69 int8 = 83
                        var v70 float32 = 0.1
                        var v71 string = (("hello" + "a") + "")
                        _ = fi68
                        _ = v69
                        _ = v70
                        _ = v71
                    }
                    _ = fi67
                }
            case 128:
                var v72 float32 = 0.0
                _ = v72
            }
        }
        if true {
            for fi73 := 0; fi73 < 10; fi73++ {
                fmt.Println(v5, v2)
                v4 = (884 >= 18446744073709551615)
                _ = fi73
            }
            switch (v5 ^ 0) {
            case 2147483647:
                for fi74 := 0; fi74 < 5; fi74++ {
                    v75 := GenS2{F0: (-(v5)), F1: 65535}
                    _ = fi74
                    _ = v75
                }
            case -851:
                if v4 {
                    if len(v26) >= 2 {
                        v26 = v26[1:len(v26)]
                    }
                    switch "" {
                    case "a":
                        _ = len(v26)
                    default:
                        for fi76 := 0; fi76 < 4; fi76++ {
                            v77 := GenS7{F0: v3, F1: ("foo bar" + ((("" + "") + "\n\t") + ("foo bar" + "a")))}
                            for fi78 := 0; fi78 < 5; fi78++ {
                                if true {
                                    v1 = -2.25
                                    var v79 float64 = (v1 + 1e10)
                                    v3 = 2147483648
                                    _ = v79
                                } else {
                                    v4 = (v4 && v4)
                                }
                                v80 := GenS10{F0: (v3 + (^(v3))), F1: -1, F2: (v2 >> ((970) & 63))}
                                _ = fi78
                                _ = v80
                            }
                            _ = fi76
                            _ = v77
                        }
                    }
                    v3 = 1
                } else {
                    var v81 int = v5
                    var v82 uint8 = 0
                    _ = v81
                    _ = v82
                }
            }
            v2 += (-(v2))
        }
        if v4 {
            v5 = v5
            v83 := GenS4{F0: (-(v5))}
            _ = v83
        }
    }
    if v4 {
        var v84 float64 = v1
        var v85 uint8 = 1
        for fi86 := 0; fi86 < 11; fi86++ {
            for fi87 := 0; fi87 < 2; fi87++ {
                v5 -= -949
                _ = fi87
            }
            if v4 {
                switch -572 {
                case 0:
                    v5 = fi86
                case -520:
                    if ((v3 * (^(v3))) > v3) {
                        v4 = ("hello" >= "a")
                        for fi88 := 0; fi88 < 18; fi88++ {
                            v2 += int32(v85)
                            v4 = ((-1 < -305) && false)
                            _ = fi88
                        }
                        v89 := GenS5{F0: v2}
                        _ = v89
                    } else {
                        _ = len(v26)
                        if v4 {
                            v3 -= v3
                            fmt.Println(v85)
                        } else {
                            if ((42 != 0) && (v4 || true)) {
                                v1 = v84
                                fmt.Println(v5, v1)
                                v1 = 1e10
                            }
                            if v4 {
                                v90 := GenS11{F0: 1, F1: 32767, F2: 42}
                                _ = v90
                            }
                            var v91 uint8 = (v85 & v85)
                            _ = v91
                        }
                        for fi92 := 0; fi92 < 4; fi92++ {
                            fmt.Println(v1, fi92, v84)
                            switch "a" {
                            case "foo bar":
                                v1 = 0.0
                            default:
                                v2 += (v2 % ((-1) | 1))
                                if (!v4) {
                                    if (v4 && (157 <= 2)) {
                                        var v93 uint32 = (v3 << ((627) & 63))
                                        _ = v93
                                    } else {
                                        v26 = append(v26, (fi86 / ((fi92) | 1)))
                                        v5 += fi86
                                        if (!v4) {
                                            v2 = (v2 | (v2 >> ((603) & 63)))
                                        } else {
                                            var v94 uint32 = (v3 | (v3 >> ((2) & 63)))
                                            v95 := GenS4{F0: (fi86 % ((-2147483648) | 1))}
                                            if (520 >= 875) {
                                                fmt.Println(v5, v85)
                                                v85 = (v85 ^ (v85 * genFn0(2, 9223372036854775807)))
                                                v5 = (fi86 * (fi92 - (v5 + fi86)))
                                            }
                                            switch ("a" + "hello") {
                                            case "\n\t":
                                                var v96 int8 = 100
                                                _ = v96
                                            case "a":
                                                v3 += 762
                                            }
                                            var v97 float32 = (0.1 - ((0.0 * float32(float64(0.1))) - 3.14159))
                                            _ = v94
                                            _ = v95
                                            _ = v97
                                        }
                                    }
                                    if (true || ((v85 | v85) != 94)) {
                                        if v4 {
                                            if len(v26) >= 2 {
                                                v26 = v26[1:len(v26)]
                                            }
                                            var v98 uint = 18446744073709551615
                                            _ = v98
                                        } else {
                                            var v99 uint32 = 1
                                            _ = v99
                                        }
                                        v2 = v2
                                        var v100 int16 = 1000
                                        var v101 int32 = (v2 * v2)
                                        _ = v100
                                        _ = v101
                                    } else {
                                        v3 = v3
                                    }
                                    fmt.Println(v5, v1)
                                    switch "" {
                                    case "a":
                                        v102 := GenS12{F0: (v3 + v3), F1: 255, F2: 9223372036854775807}
                                        v84 += ((v1 + v1) / v84)
                                        _ = v102
                                    case "\n\t":
                                        if true {
                                            var v103 uint16 = 2
                                            _ = v103
                                        } else {
                                            fmt.Println(v4, v4, v85)
                                        }
                                    }
                                    v2 = (v2 ^ 2147483647)
                                }
                                var v104 int16 = 151
                                _ = v104
                            }
                            if ((v84 >= v1) && (0 > -32768)) {
                                if len(v26) >= 2 {
                                    v26 = v26[1:len(v26)]
                                }
                                v105 := GenS13{F0: 18446744073709551615, F1: (^(v2)), F2: (v2 >> ((663) & 63))}
                                _ = v105
                            } else {
                                v106 := GenS8{F0: 1, F1: 594, F2: 425}
                                _ = v106
                            }
                            _ = fi92
                        }
                    }
                default:
                    switch "hello" {
                    case "a":
                        switch "foo bar" {
                        case "hello":
                            v1 -= v84
                        case "\n\t":
                            for fi107 := 0; fi107 < 7; fi107++ {
                                fmt.Println(fi86, v84)
                                var v108 string = "a"
                                v109 := GenS10{F0: 0, F1: -1, F2: (v2 ^ v2)}
                                if len(v26) >= 2 {
                                    v26 = v26[1:len(v26)]
                                }
                                _ = fi107
                                _ = v108
                                _ = v109
                            }
                        }
                    case "hello":
                        var v110 float64 = (float64(0.0) * v1)
                        _ = v110
                    }
                }
                if v4 {
                    for fi111 := 0; fi111 < 20; fi111++ {
                        if false {
                            if true {
                                v84 = -2.25
                                var v112 int32 = v2
                                _ = v112
                            }
                            if (42 >= 0) {
                                v113 := GenS1{F0: v4, F1: 42}
                                switch "foo bar" {
                                case "":
                                    var v114 int8 = 100
                                    var v115 float32 = 1.5
                                    v85 = (v85 / ((v85) | 1))
                                    fmt.Println(v4, v4)
                                    v26 = append(v26, v5)
                                    _ = v114
                                    _ = v115
                                }
                                _ = v113
                            }
                        } else {
                            if len(v26) >= 2 {
                                v26 = v26[1:len(v26)]
                            }
                        }
                        v116 := GenS14{F0: (-(v5)), F1: 1}
                        _ = fi111
                        _ = v116
                    }
                }
                for fi117 := 0; fi117 < 12; fi117++ {
                    if len(v26) > 0 {
                        v26[0] = fi117
                    }
                    if len(v26) >= 2 {
                        v26 = v26[1:len(v26)]
                    }
                    if (42 != 42) {
                        if ((42 <= 127) && (2 == 18446744073709551615)) {
                            var v118 uint8 = v85
                            _ = v118
                        } else {
                            if len(v26) > 0 {
                                v26[0] = 384
                            }
                        }
                    }
                    _ = fi117
                }
            }
            fmt.Println(fi86, v1, v4)
            _ = fi86
        }
        _ = v84
        _ = v85
    } else {
        v5 = -699
        v119 := GenS15{F0: 87, F1: 477, F2: 255}
        var v120 uint64 = 4294967295
        _ = v119
        _ = v120
    }
    var v121 int64 = 1
    v122 := GenS7{F0: (v3 & v3), F1: "a"}
    var v123 int64 = v121
    if v4 {
        var v124 float32 = 1e10
        var v125 uint64 = 0
        v2 = -337
        _ = v124
        _ = v125
    } else {
        v5 += (-(v5))
    }
    for fi126 := 0; fi126 < 10; fi126++ {
        if len(v26) > 0 {
            v26[0] = (fi126 ^ v5)
        }
        fmt.Println(v123, v3, v3)
        _ = fi126
    }
    v121 -= 9223372036854775807
    for fi127 := 0; fi127 < 15; fi127++ {
        if (1 <= 1000) {
            switch (fi127 % ((-437) | 1)) {
            case 1:
                v2 -= (v2 * int32(v5))
            case 127:
                if (128 >= 128) {
                    if (327 < 244) {
                        v128 := GenS16{F0: 42, F1: -1}
                        _ = v128
                    }
                    v129 := GenS10{F0: (v3 ^ (v3 & v3)), F1: -1, F2: (-(v2))}
                    v130 := GenS11{F0: 654, F1: 1, F2: 75}
                    _ = v129
                    _ = v130
                } else {
                    v1 = ((3.14159 / 0.1) - (-2.25 - (1.5 + (3.14159 * v1))))
                }
            case -2147483648:
                switch (-(fi127)) {
                case 489:
                    fmt.Println(v2, v5)
                case 128:
                    for fi131 := 0; fi131 < 19; fi131++ {
                        switch "" {
                        case "hello":
                            var v132 uint32 = (v3 >> ((2) & 63))
                            v4 = (v4 && v4)
                            _ = v132
                        case "\n\t":
                            var v133 uint64 = 42
                            _ = v133
                        case "":
                            switch "foo bar" {
                            case "hello":
                                var v134 int64 = -9223372036854775808
                                _ = v134
                            case "\n\t":
                                switch ((("\n\t" + ("" + (("\n\t" + "foo bar") + "hello"))) + "") + "foo bar") {
                                case "a":
                                    if len(v26) > 0 {
                                        v26[0] = (fi127 - (^(fi131)))
                                    }
                                case "":
                                    if ((v2 & (v2 + 1)) >= (v2 ^ v2)) {
                                        v26 = append(v26, (v5 * (v5 & fi127)))
                                        if (42 >= 980) {
                                            v123 -= (v123 / ((-1) | 1))
                                            if len(v26) >= 2 {
                                                v26 = v26[1:len(v26)]
                                            }
                                            var v135 float32 = float32(v1)
                                            _ = v135
                                        }
                                    } else {
                                        v4 = (v2 < v2)
                                        v136 := GenS14{F0: v5, F1: (v2 & v2)}
                                        _ = v136
                                    }
                                default:
                                    fmt.Println(v4, v123, fi127)
                                }
                            default:
                                v4 = true
                            }
                        default:
                            var v137 uint32 = v3
                            if (true || v4) {
                                switch fi131 {
                                case 34:
                                    v138 := GenS7{F0: (v3 << ((1) & 63)), F1: "hello"}
                                    _ = v138
                                case 842:
                                    fmt.Println(v5, fi127, v1)
                                case 0:
                                    if false {
                                        if false {
                                            switch fi127 {
                                            case 704:
                                                v123 += (v121 | v121)
                                            case 4294967295:
                                                v139 := GenS11{F0: 67, F1: 0, F2: -68}
                                                _ = v139
                                            default:
                                                var v140 int16 = 1000
                                                _ = v140
                                            }
                                            v2 += (v2 - (-(v2)))
                                        }
                                    }
                                    if v4 {
                                        var v141 int16 = -349
                                        var v142 uint16 = 698
                                        if ((int16(v5) > 0) || v4) {
                                            var v143 uint64 = 4294967295
                                            v3 = v137
                                            _ = v143
                                        } else {
                                            var v144 uint8 = genFn0(2, 0)
                                            _ = v144
                                        }
                                        _ = v141
                                        _ = v142
                                    } else {
                                        var v145 int8 = -128
                                        v146 := GenS3{F0: (v5 >> ((255) & 63)), F1: false}
                                        if len(v26) >= 2 {
                                            v26 = v26[1:len(v26)]
                                        }
                                        _ = v145
                                        _ = v146
                                    }
                                    v123 += 1
                                default:
                                    if (1 >= -119) {
                                        if (v4 && false) {
                                            switch ("" + "") {
                                            case "hello":
                                                switch "" {
                                                case "a":
                                                    v26 = append(v26, (-(v5)))
                                                default:
                                                    v1 -= v1
                                                }
                                            case "foo bar":
                                                v5 -= (fi127 + (fi127 << ((0) & 63)))
                                            case "a":
                                                if len(v26) > 0 {
                                                    v26[0] = (fi127 & fi131)
                                                }
                                            }
                                            var v147 int = (v5 | 128)
                                            _ = v147
                                        }
                                        v1 = v1
                                    } else {
                                        var v148 int = fi127
                                        _ = v148
                                    }
                                }
                                var v149 int16 = 42
                                _ = v149
                            }
                            _ = v137
                        }
                        _ = fi131
                    }
                case -563:
                    var v150 uint64 = 0
                    _ = v150
                default:
                    fmt.Println(v121, v3, v5)
                }
            }
            switch "\n\t" {
            case "foo bar":
                switch "\n\t" {
                case "\n\t":
                    if len(v26) > 0 {
                        v26[0] = fi127
                    }
                case "a":
                    fmt.Println(v1)
                case "foo bar":
                    var v151 uint16 = 42
                    _ = v151
                default:
                    for fi152 := 0; fi152 < 9; fi152++ {
                        switch "\n\t" {
                        case "foo bar":
                            v5 = 32767
                        case "\n\t":
                            var v153 int = 42
                            _ = v153
                        default:
                            var v154 uint16 = 979
                            _ = v154
                        }
                        fmt.Println(v2, fi152)
                        _ = fi152
                    }
                }
            case "":
                if len(v26) >= 2 {
                    v26 = v26[1:len(v26)]
                }
            default:
                for fi155 := 0; fi155 < 2; fi155++ {
                    if len(v26) >= 2 {
                        v26 = v26[1:len(v26)]
                    }
                    _ = fi155
                }
            }
            v121 -= v123
        }
        if false {
            v156 := GenS8{F0: 343, F1: 328, F2: 18446744073709551615}
            v26 = append(v26, -291)
            _ = v156
        }
        _ = fi127
    }
    var v157 float64 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v121:", v121)
    fmt.Println("v123:", v123)
    fmt.Println("v157:", v157)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v121
    _ = v122
    _ = v123
    _ = v157
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
