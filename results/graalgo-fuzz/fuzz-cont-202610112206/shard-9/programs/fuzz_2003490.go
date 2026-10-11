package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 int }
type GenS2 struct { F0 int64; F1 uint; F2 int64 }
type GenS3 struct { F0 uint; F1 int; F2 uint32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 bool; F1 int32 }
type GenS6 struct { F0 int16; F1 uint32 }
type GenS7 struct { F0 uint8; F1 uint64 }
type GenS8 struct { F0 uint16; F1 bool }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 int8; F1 int8 }
type GenS11 struct { F0 int8; F1 int }
type GenS12 struct { F0 string; F1 int16; F2 uint64 }
type GenS13 struct { F0 bool; F1 int64 }
type GenS14 struct { F0 uint; F1 string; F2 int16 }
type GenS15 struct { F0 string; F1 int32; F2 bool }
type GenS16 struct { F0 uint; F1 uint }
func genFn0() int16 {
    var v7 int8 = -128
    v7 += v7
    var v8 int8 = v7
    _ = v7
    _ = v8
    return 285
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint64 = 178
    var v3 uint32 = 2147483648
    var v4 uint64 = v2
    v5 := make(map[int]uint32)
    var v6 uint16 = 42
    var v9 int = 717
    v10 := []int{(v9 - (v9 - (^(v9)))), -543, (v9 / ((v9) | 1))}
    var v11 uint16 = v6
    var v12 float32 = (1.5 + 1.5)
    v12 += v12
    var v13 float64 = 1.5
    v11 -= v11
    switch ("foo bar" + ("\n\t" + "foo bar")) {
    case "foo bar":
        switch "hello" {
        case "foo bar":
            for fi14 := 0; fi14 < 16; fi14++ {
                var v15 int32 = 1
                switch (^(v9)) {
                case 1:
                    var v16 uint8 = 221
                    _ = v16
                default:
                    var v17 int16 = genFn0()
                    _ = v17
                }
                fmt.Println(v12)
                _ = fi14
                _ = v15
            }
        case "":
            switch "\n\t" {
            case "":
                for fi18 := 0; fi18 < 7; fi18++ {
                    switch "foo bar" {
                    case "":
                        v19 := GenS1{F0: 324, F1: 32767, F2: 591}
                        _ = v19
                    case "hello":
                        v12 = v12
                    default:
                        v20 := GenS2{F0: 9223372036854775807, F1: 18446744073709551615, F2: -1}
                        _ = v20
                    }
                    v3 -= (v3 ^ (^(v3)))
                    if (0.0 <= v13) {
                        for fi21 := 0; fi21 < 10; fi21++ {
                            if false {
                                var v22 uint = 65535
                                v22 -= v22
                                v22 -= (^(v22))
                                _ = v22
                            }
                            _ = fi21
                        }
                        var v23 uint = 1
                        for fi24 := 0; fi24 < 15; fi24++ {
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                            if true {
                                var v25 float32 = v12
                                _ = v25
                            } else {
                                v2 = 708
                                fmt.Println(fi24, v6)
                                v26 := GenS3{F0: 74, F1: v9, F2: 2}
                                v4 += 42
                                _ = v26
                            }
                            v12 -= v12
                            if false {
                                var v27 uint16 = v11
                                v5[32767] = 140
                                _ = v27
                            } else {
                                var v28 int8 = 26
                                _ = v28
                            }
                            switch 128 {
                            case 0:
                                var v29 float64 = v13
                                v3 = v3
                                if (!true) {
                                    var v30 uint = (^(v23))
                                    _ = v30
                                }
                                _ = v29
                            case 42:
                                _ = len(v10)
                            }
                            _ = fi24
                        }
                        _ = v23
                    } else {
                        if (-2147483648 > 2147483647) {
                            if (!(!false)) {
                                if ((42 == -128) || (!false)) {
                                    for fi31 := 0; fi31 < 7; fi31++ {
                                        if len(v10) >= 2 {
                                            v10 = v10[1:len(v10)]
                                        }
                                        v1[(^(v3))] = v3
                                        var v32 uint64 = (v2 | (v2 | 259))
                                        _ = fi31
                                        _ = v32
                                    }
                                    fmt.Println(v3, v3, v3)
                                    v4 += (v2 ^ v4)
                                } else {
                                    v9 = (^(v9))
                                    switch 0 {
                                    case 4294967295:
                                        delete(v1, v3)
                                    default:
                                        v12 -= (1e10 + v12)
                                    }
                                    for fi33 := 0; fi33 < 14; fi33++ {
                                        var v34 uint = uint(v2)
                                        v35 := GenS3{F0: v34, F1: fi18, F2: (v3 + 902)}
                                        _ = fi33
                                        _ = v34
                                        _ = v35
                                    }
                                }
                            }
                            v13 = -2.25
                            v36 := GenS3{F0: 108, F1: -432, F2: (^(v3))}
                            _ = v36
                        } else {
                            v12 = (((v12 / v12) * -2.25) / (-2.25 / ((3.14159 + v12) / ((((((-2.25 * v12) / v12) - (v12 * (v12 + ((((v12 / ((((1.5 - (1.5 - (0.0 - v12))) - 1.5) - (v12 * v12)) - v12)) - v12) * -2.25) / (1.5 + v12))))) - ((float32((float64((((0.1 + 1.5) - 1e10) - 0.1)) - (1e10 * float64(v12)))) - v12) / v12)) * float32(((1.5 - (v13 / 3.14159)) - v13))) + (1e10 + v12)))))
                            if (((v3 != (v3 / (((^(v3))) | 1))) || true) || (496 != -1)) {
                                _ = len(v1)
                                v37 := v5[(^(v9))]
                                _ = v37
                            }
                        }
                    }
                    _ = fi18
                }
            case "a":
                if false {
                    if false {
                        for fi38 := 0; fi38 < 1; fi38++ {
                            v39 := GenS3{F0: 42, F1: -535, F2: uint32(v2)}
                            _ = fi38
                            _ = v39
                        }
                        var v40 uint32 = 42
                        _ = v40
                    } else {
                        v41 := GenS4{F0: v3}
                        _ = v41
                    }
                    switch (v9 - v9) {
                    case -414:
                        if (true || true) {
                            switch "a" {
                            case "a":
                                _ = len(v10)
                            }
                            v42 := GenS2{F0: -1, F1: 925, F2: 9223372036854775807}
                            switch "a" {
                            case "hello":
                                if true {
                                    v1[(v3 * (v3 | (^(v3))))] = v3
                                    fmt.Println(v12, v6, v9)
                                    v13 = v13
                                }
                            case "":
                                for fi43 := 0; fi43 < 6; fi43++ {
                                    v4 = (v4 - 0)
                                    if false {
                                        for fi44 := 0; fi44 < 6; fi44++ {
                                            switch (("a" + ("a" + "a")) + ("foo bar" + "foo bar")) {
                                            case "":
                                                v3 = 39
                                            case "\n\t":
                                                if len(v10) >= 2 {
                                                    v10 = v10[1:len(v10)]
                                                }
                                            case "hello":
                                                fmt.Println(v9, v12, v4)
                                            }
                                            v4 -= v4
                                            v45 := GenS5{F0: true, F1: 42}
                                            v13 = 0.0
                                            var v46 float64 = v13
                                            v13 = v13
                                            v10 = append(v10, (v9 / (((-(fi43))) | 1)))
                                            _ = fi44
                                            _ = v45
                                            _ = v46
                                        }
                                    }
                                    _ = fi43
                                }
                            default:
                                var v47 string = "foo bar"
                                _ = v47
                            }
                            _ = v42
                        } else {
                            fmt.Println(v2, v4, v13)
                            if true {
                                switch "foo bar" {
                                case "\n\t":
                                    var v48 int32 = -47
                                    _ = v48
                                }
                            } else {
                                v49 := GenS4{F0: v3}
                                _ = v49
                            }
                            var v50 uint64 = v4
                            _ = v50
                        }
                    case 32768:
                        if true {
                            if (("a" + "foo bar") <= "hello") {
                                var v51 uint32 = v3
                                v52 := GenS3{F0: 949, F1: (v9 + v9), F2: v3}
                                _ = v51
                                _ = v52
                            }
                            switch v9 {
                            case 1:
                                if ((v11 >> ((42) & 63)) == v6) {
                                    if false {
                                        if true {
                                            if false {
                                                var v53 uint32 = 1
                                                v54 := GenS6{F0: 918, F1: v3}
                                                if (-770 <= 1) {
                                                    var v55 uint16 = (v11 % ((v6) | 1))
                                                    var v56 int16 = 0
                                                    _ = v55
                                                    _ = v56
                                                } else {
                                                    var v57 float64 = v13
                                                    _ = v57
                                                }
                                                _ = v53
                                                _ = v54
                                            }
                                        }
                                        if true {
                                            v13 += v13
                                            delete(v5, v9)
                                        } else {
                                            var v58 string = ""
                                            for fi59 := 0; fi59 < 14; fi59++ {
                                                if ((v3 + v3) < uint32(v2)) {
                                                    var v60 uint64 = (^(v4))
                                                    _ = v60
                                                } else {
                                                    switch (^(fi59)) {
                                                    case -2147483648:
                                                        if (!false) {
                                                            if len(v10) > 0 {
                                                                v10[0] = fi59
                                                            }
                                                            for fi61 := 0; fi61 < 11; fi61++ {
                                                                v62 := GenS4{F0: v3}
                                                                _ = fi61
                                                                _ = v62
                                                            }
                                                            v63 := GenS6{F0: 32767, F1: (v3 >> ((2) & 63))}
                                                            _ = v63
                                                        }
                                                    default:
                                                        fmt.Println(v4, v12)
                                                    }
                                                }
                                                v9 = (v9 + (fi59 / (((fi59 >> ((42) & 63))) | 1)))
                                                var v64 int64 = int64(v3)
                                                _ = fi59
                                                _ = v64
                                            }
                                            _ = v58
                                        }
                                    }
                                    if (!(0 == 127)) {
                                        _ = len(v1)
                                        var v65 int16 = 98
                                        _ = v65
                                    }
                                    var v66 int = v9
                                    _ = v66
                                }
                            }
                            switch ("a" + "a") {
                            case "":
                                if ((v3 >> ((1) & 63)) >= 2) {
                                    var v67 bool = (123 > 255)
                                    for fi68 := 0; fi68 < 7; fi68++ {
                                        v11 = (v6 * 42)
                                        var v69 int = 373
                                        _ = fi68
                                        _ = v69
                                    }
                                    if v67 {
                                        v67 = true
                                        for fi70 := 0; fi70 < 5; fi70++ {
                                            v71 := GenS7{F0: 133, F1: v2}
                                            _ = fi70
                                            _ = v71
                                        }
                                        fmt.Println(v6)
                                    } else {
                                        if (!v67) {
                                            if v67 {
                                                var v72 float64 = ((-2.25 * (v13 + v13)) + v13)
                                                for fi73 := 0; fi73 < 7; fi73++ {
                                                    switch "hello" {
                                                    case "a":
                                                        if (362 <= uint(v3)) {
                                                            v13 += (v72 / 1e10)
                                                        }
                                                    case "foo bar":
                                                        fmt.Println(v3)
                                                    case "":
                                                        v5[-9223372036854775808] = 2
                                                    }
                                                    v74 := GenS5{F0: true, F1: 0}
                                                    v11 = v11
                                                    _ = fi73
                                                    _ = v74
                                                }
                                                var v75 int8 = 127
                                                _ = v72
                                                _ = v75
                                            } else {
                                                v76 := GenS4{F0: (v3 >> ((uint(v3)) & 63))}
                                                _ = v76
                                            }
                                        }
                                    }
                                    _ = v67
                                }
                            default:
                                switch ("" + "") {
                                case "a":
                                    if true {
                                        _ = len(v1)
                                        if true {
                                            for fi77 := 0; fi77 < 6; fi77++ {
                                                var v78 int64 = 1
                                                switch (fi77 << ((0) & 63)) {
                                                case 127:
                                                    var v79 int16 = 42
                                                    _ = v79
                                                case 0:
                                                    fmt.Println(v4, v2)
                                                }
                                                if len(v10) > 0 {
                                                    v10[0] = (fi77 / (((fi77 ^ 0)) | 1))
                                                }
                                                _ = fi77
                                                _ = v78
                                            }
                                            fmt.Println(v2, v4)
                                        } else {
                                            for fi80 := 0; fi80 < 6; fi80++ {
                                                for fi81 := 0; fi81 < 19; fi81++ {
                                                    fmt.Println(v13, v12)
                                                    if (!true) {
                                                        switch (v9 % ((915) | 1)) {
                                                        case 32768:
                                                            v13 = (1.5 / v13)
                                                        case -731:
                                                            var v82 float32 = v12
                                                            _ = v82
                                                        default:
                                                            switch 127 {
                                                            case -9223372036854775808:
                                                                fmt.Println(v2, v11)
                                                            case 560:
                                                                if true {
                                                                    var v83 string = "a"
                                                                    fmt.Println(v4, v6)
                                                                    _ = v83
                                                                }
                                                                v84 := GenS8{F0: v11, F1: ((v11 ^ 255) >= (v11 / ((2) | 1)))}
                                                                _ = v84
                                                            default:
                                                                v85 := GenS9{F0: v3}
                                                                _ = v85
                                                            }
                                                        }
                                                    } else {
                                                        var v86 int16 = genFn0()
                                                        var v87 int16 = (v86 | v86)
                                                        delete(v5, fi81)
                                                        _ = v86
                                                        _ = v87
                                                    }
                                                    _ = fi81
                                                }
                                                switch "" {
                                                case "\n\t":
                                                    v88 := GenS10{F0: 127, F1: -1}
                                                    _ = v88
                                                case "hello":
                                                    var v89 int = (fi80 * v9)
                                                    _ = v89
                                                }
                                                if len(v10) >= 2 {
                                                    v10 = v10[1:len(v10)]
                                                }
                                                _ = fi80
                                            }
                                        }
                                        v90 := GenS11{F0: 0, F1: (v9 - 0)}
                                        _ = v90
                                    }
                                case "":
                                    v6 -= 65535
                                default:
                                    fmt.Println(v11, v9)
                                }
                            }
                        }
                    case 0:
                        for fi91 := 0; fi91 < 8; fi91++ {
                            v11 -= (v11 / ((v6) | 1))
                            for fi92 := 0; fi92 < 20; fi92++ {
                                var v93 uint8 = 127
                                _ = fi92
                                _ = v93
                            }
                            _ = len(v10)
                            _ = fi91
                        }
                    }
                }
            case "hello":
                delete(v5, v9)
            default:
                v4 = (v4 + 2)
            }
        default:
            v4 = 4294967295
        }
    case "hello":
        for fi94 := 0; fi94 < 1; fi94++ {
            for fi95 := 0; fi95 < 8; fi95++ {
                if true {
                    if false {
                        v12 = -2.25
                        var v96 int = 382
                        _ = v96
                    } else {
                        v9 -= fi94
                        if (false && false) {
                            switch (v9 ^ (-(v9))) {
                            case -9223372036854775808:
                                v97 := GenS5{F0: (false && (1e10 >= 0.1)), F1: -630}
                                _ = v97
                            case -2147483648:
                                var v98 float64 = v13
                                _ = v98
                            case -124:
                                var v99 uint = 380
                                _ = v99
                            }
                            if (!true) {
                                v13 += (0.0 * 1e10)
                                var v100 int32 = 2147483647
                                _ = v100
                            } else {
                                v101 := GenS12{F0: "foo bar", F1: int16(v6), F2: (v2 & v2)}
                                v102 := GenS13{F0: (127 < 251), F1: -541}
                                _ = v101
                                _ = v102
                            }
                            var v103 int8 = 100
                            _ = v103
                        } else {
                            v9 = fi95
                            if false {
                                var v104 bool = (float64(v12) <= (v13 + -2.25))
                                if v104 {
                                    v10 = append(v10, (v9 - fi95))
                                    v105 := GenS7{F0: 42, F1: (v2 * (v2 ^ v4))}
                                    var v106 uint8 = 191
                                    _ = v105
                                    _ = v106
                                }
                                _ = v104
                            } else {
                                if (!true) {
                                    v11 = 0
                                    if true {
                                        switch 0 {
                                        case 316:
                                            if true {
                                                v107 := GenS14{F0: 209, F1: "\n\t", F2: genFn0()}
                                                var v108 float32 = 0.1
                                                _ = v107
                                                _ = v108
                                            } else {
                                                var v109 string = ""
                                                v9 += 9223372036854775807
                                                _ = v109
                                            }
                                        case 127:
                                            var v110 int32 = 42
                                            switch "foo bar" {
                                            case "hello":
                                                v2 += v4
                                            case "foo bar":
                                                if len(v10) >= 2 {
                                                    v10 = v10[1:len(v10)]
                                                }
                                            default:
                                                if len(v10) > 0 {
                                                    v10[0] = 4294967295
                                                }
                                            }
                                            _ = v110
                                        }
                                        v10 = append(v10, int(v3))
                                        fmt.Println(v12, v6)
                                        v111 := GenS15{F0: "", F1: 42, F2: true}
                                        switch int(v11) {
                                        case 39:
                                            if false {
                                                v12 = (v12 / 1.0)
                                                switch "foo bar" {
                                                case "":
                                                    switch fi95 {
                                                    case -9223372036854775808:
                                                        var v112 float32 = v12
                                                        _ = v112
                                                    default:
                                                        if false {
                                                            switch "a" {
                                                            case "a":
                                                                if (!((true || true) && true)) {
                                                                    var v113 int16 = 1
                                                                    _ = v113
                                                                } else {
                                                                    if true {
                                                                        v9 -= -9223372036854775808
                                                                    }
                                                                    if len(v10) > 0 {
                                                                        v10[0] = fi95
                                                                    }
                                                                }
                                                            case "hello":
                                                                if false {
                                                                    switch fi94 {
                                                                    case 127:
                                                                        var v114 float32 = float32(v13)
                                                                        _ = v114
                                                                    case -2147483648:
                                                                        v115 := GenS7{F0: 255, F1: v2}
                                                                        _ = v115
                                                                    case 32768:
                                                                        v116 := GenS12{F0: ("foo bar" + ""), F1: genFn0(), F2: v2}
                                                                        _ = v116
                                                                    default:
                                                                        var v117 float64 = (float64(1.5) - (v13 + float64(v12)))
                                                                        _ = v117
                                                                    }
                                                                }
                                                                var v118 uint64 = (^(v2))
                                                                var v119 uint8 = 12
                                                                fmt.Println(v3)
                                                                _ = v118
                                                                _ = v119
                                                            }
                                                            var v120 int8 = 1
                                                            _ = v120
                                                        }
                                                    }
                                                default:
                                                    var v121 string = ""
                                                    _ = v121
                                                }
                                                if false {
                                                    if true {
                                                        v6 -= v6
                                                    }
                                                } else {
                                                    if len(v10) > 0 {
                                                        v10[0] = fi94
                                                    }
                                                    var v122 int = -2147483648
                                                    fmt.Println(v13, v122, v11)
                                                    _ = v122
                                                }
                                            }
                                        default:
                                            if (!(!false)) {
                                                v10 = append(v10, (fi95 >> ((255) & 63)))
                                            } else {
                                                switch (^(v9)) {
                                                case -591:
                                                    fmt.Println(v6)
                                                case -134:
                                                    v123 := v1[(v3 * 2147483647)]
                                                    _ = v123
                                                case 2147483647:
                                                    v10 = append(v10, 0)
                                                default:
                                                    v4 += v4
                                                }
                                            }
                                        }
                                        switch "a" {
                                        case "\n\t":
                                            v124 := GenS16{F0: 255, F1: 18446744073709551615}
                                            _ = v124
                                        case "hello":
                                            var v125 string = "a"
                                            _ = v125
                                        case "":
                                            var v126 int64 = 0
                                            _ = v126
                                        default:
                                            delete(v5, (fi94 & v9))
                                        }
                                        _ = v111
                                    } else {
                                        v12 = -2.25
                                        var v127 uint16 = v6
                                        _ = v127
                                    }
                                } else {
                                    v3 -= 715
                                }
                            }
                        }
                        var v128 string = "\n\t"
                        _ = v128
                    }
                    v4 -= v2
                }
                fmt.Println(v3, v12, fi95)
                _ = fi95
            }
            _ = fi94
        }
    default:
        fmt.Println(v3, v12)
    }
    var v129 int = (-(v9))
    _ = len(v5)
    switch v9 {
    case -9223372036854775808:
        var v130 uint8 = 42
        _ = v130
    case -507:
        fmt.Println(v3, v3)
    case 202:
        if false {
            var v131 float64 = ((3.14159 / -2.25) * (0.1 / (float64(float32(v13)) / v13)))
            _ = v131
        } else {
            if true {
                for fi132 := 0; fi132 < 19; fi132++ {
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                    delete(v5, (v129 | v9))
                    for fi133 := 0; fi133 < 7; fi133++ {
                        fmt.Println(v12, v3)
                        var v134 uint16 = 154
                        _ = fi133
                        _ = v134
                    }
                    _ = fi132
                }
                for fi135 := 0; fi135 < 15; fi135++ {
                    var v136 bool = false
                    _ = fi135
                    _ = v136
                }
                var v137 uint32 = 2147483647
                _ = v137
            } else {
                var v138 uint32 = v3
                v139 := GenS8{F0: v6, F1: (0 >= 0)}
                v1[(^(v3))] = 42
                _ = v138
                _ = v139
            }
            fmt.Println(v3, v11)
        }
    default:
        var v140 uint8 = 42
        _ = v140
    }
    if len(v10) > 0 {
        v10[0] = 2147483647
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v129:", v129)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v129
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
