package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int; F1 uint16; F2 uint16 }
type GenS4 struct { F0 int8; F1 uint }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint; F1 int32 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint64; F1 bool }
type GenS10 struct { F0 int16; F1 int16; F2 uint }
type GenS11 struct { F0 uint8; F1 uint8 }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 string; F1 uint8; F2 uint64 }
type GenS14 struct { F0 int; F1 uint32 }
type GenS15 struct { F0 uint32; F1 string }
type GenS16 struct { F0 int16 }
type GenS17 struct { F0 string; F1 int }

func main() {
    var v1 float32 = 0.1
    var v2 string = ("a" + "\n\t")
    v3 := []int{}
    var v4 uint8 = 128
    var v5 int16 = -1
    v6 := make(map[string]uint32)
    for fi7 := 0; fi7 < 7; fi7++ {
        fmt.Println(v5)
        var v8 uint16 = 0
        _ = fi7
        _ = v8
    }
    if false {
        var v9 uint16 = 0
        var v10 int8 = 42
        _ = v10
        _ = v9
    }
    if (!(-463 < v5)) {
        if true {
            if (255 <= 886) {
                for fi11 := 0; fi11 < 4; fi11++ {
                    v2 = v2
                    _ = fi11
                }
            } else {
                v6[(v2 + "")] = 225
            }
        }
    }
    fmt.Println(v4)
    for fi12 := 0; fi12 < 8; fi12++ {
        fmt.Println(v2, v4, fi12)
        delete(v6, v2)
        for fi13 := 0; fi13 < 20; fi13++ {
            v4 -= v4
            _ = fi13
        }
        _ = fi12
    }
    for fi14 := 0; fi14 < 17; fi14++ {
        for fi15 := 0; fi15 < 13; fi15++ {
            var v16 uint = 35
            v17 := GenS1{F0: v2, F1: 0, F2: 42}
            _ = fi15
            _ = v16
            _ = v17
        }
        if true {
            for fi18 := 0; fi18 < 16; fi18++ {
                var v19 string = v2
                switch (fi18 / ((32767) | 1)) {
                case 285:
                    switch v19 {
                    case "foo bar":
                        var v20 int = (fi18 >> ((1) & 63))
                        _ = v20
                    default:
                        var v21 int = 32768
                        _ = v21
                    }
                case 607:
                    delete(v6, v19)
                case -569:
                    delete(v6, (("foo bar" + "foo bar") + "\n\t"))
                default:
                    switch 871 {
                    case 32767:
                        if true {
                            if (265 < 2) {
                                var v22 uint = 315
                                var v23 uint32 = 0
                                v2 = v19
                                _ = v22
                                _ = v23
                            }
                            var v24 uint = 4294967295
                            switch fi14 {
                            case -851:
                                v2 = v19
                            case -661:
                                v5 = (v5 >> (((v24 % ((v24) | 1))) & 63))
                            default:
                                if false {
                                    if len(v3) > 0 {
                                        v3[0] = 9223372036854775807
                                    }
                                    v4 = 128
                                    fmt.Println(v1, fi14, fi14)
                                } else {
                                    v2 = ("" + v19)
                                }
                            }
                            _ = v24
                        } else {
                            v2 = v19
                        }
                    case -2147483648:
                        fmt.Println(v5, v2)
                    }
                }
                _ = fi18
                _ = v19
            }
        } else {
            var v25 float32 = (v1 + v1)
            _ = v25
        }
        _ = fi14
    }
    if (0 <= 4294967295) {
        for fi26 := 0; fi26 < 6; fi26++ {
            delete(v6, v2)
            _ = fi26
        }
        if true {
            var v27 int16 = (^(v5))
            for fi28 := 0; fi28 < 18; fi28++ {
                for fi29 := 0; fi29 < 12; fi29++ {
                    var v30 float32 = (v1 / 1.0)
                    _ = fi29
                    _ = v30
                }
                if false {
                    if true {
                        switch v2 {
                        case "foo bar":
                            v31 := v6[v2]
                            _ = v31
                        case "a":
                            v32 := GenS2{F0: 0}
                            _ = v32
                        }
                    } else {
                        if true {
                            fmt.Println(v4, v5)
                        }
                        var v33 uint64 = 18446744073709551615
                        if false {
                            v6[""] = 432
                            if (2147483647 == 42) {
                                switch ("a" + "hello") {
                                case "\n\t":
                                    v34 := v6[("\n\t" + "")]
                                    _ = v34
                                case "a":
                                    var v35 int = (fi28 ^ -257)
                                    _ = v35
                                case "foo bar":
                                    fmt.Println(v1, v1)
                                }
                            } else {
                                if (!true) {
                                    var v36 int = (-(fi28))
                                    var v37 int64 = 9223372036854775807
                                    var v38 uint64 = (v33 | (v33 | v33))
                                    _ = v36
                                    _ = v37
                                    _ = v38
                                }
                                var v39 string = (v2 + "a")
                                _ = v39
                            }
                        }
                        _ = v33
                    }
                    v40 := GenS1{F0: v2, F1: 9223372036854775807, F2: 0}
                    switch (fi28 % ((fi28) | 1)) {
                    case 127:
                        var v41 uint64 = 0
                        _ = v41
                    case 2147483647:
                        fmt.Println(v4, v4)
                    case 4294967295:
                        for fi42 := 0; fi42 < 9; fi42++ {
                            v6[""] = 0
                            if (!("" <= (v2 + (((v2 + ("a" + v2)) + (v2 + v2)) + v2)))) {
                                v43 := GenS2{F0: -2147483648}
                                _ = v43
                            } else {
                                v44 := GenS3{F0: (fi28 * 278), F1: 42, F2: 255}
                                v4 += (v4 * (v4 & (v4 % ((v4) | 1))))
                                if true {
                                    v27 -= 0
                                    if true {
                                        var v45 float32 = (-2.25 - ((v1 * v1) + v1))
                                        if (!true) {
                                            v46 := GenS2{F0: -577}
                                            switch ((((v2 + (v2 + ((("hello" + ((v2 + v2) + ((((v2 + v2) + "\n\t") + v2) + v2))) + "\n\t") + v2))) + (v2 + (v2 + v2))) + v2) + "") {
                                            case "hello":
                                                var v47 uint32 = 330
                                                _ = v47
                                            }
                                            if (true || (true || false)) {
                                                switch "foo bar" {
                                                case "a":
                                                    v2 = "\n\t"
                                                default:
                                                    if true {
                                                        var v48 int8 = -76
                                                        fmt.Println(v48)
                                                        fmt.Println(v4, v5)
                                                        _ = v48
                                                    } else {
                                                        var v49 uint = 307
                                                        fmt.Println(fi28)
                                                        _ = v49
                                                    }
                                                }
                                                fmt.Println(v2, v4)
                                                v50 := GenS4{F0: int8(v4), F1: 65535}
                                                fmt.Println(v2, v1)
                                                if false {
                                                    var v51 uint32 = 681
                                                    v52 := GenS5{F0: 18446744073709551615}
                                                    _ = v51
                                                    _ = v52
                                                } else {
                                                    switch v2 {
                                                    case "\n\t":
                                                        v53 := GenS6{F0: 18446744073709551615, F1: 486}
                                                        _ = v53
                                                    default:
                                                        var v54 int64 = 1
                                                        _ = v54
                                                    }
                                                }
                                                _ = v50
                                            } else {
                                                if ((true || (((3.14159 + -2.25) + float64((0.0 / -2.25))) == 0.1)) && (!(!false))) {
                                                    var v55 uint8 = v4
                                                    v4 += v4
                                                    _ = v55
                                                }
                                                v56 := GenS6{F0: 0, F1: -599}
                                                v57 := GenS2{F0: -2147483648}
                                                _ = v56
                                                _ = v57
                                            }
                                            _ = v46
                                        }
                                        _ = v45
                                    } else {
                                        if ((0 != 255) || (true || false)) {
                                            if ((fi28 + (fi42 & 128)) != (fi42 / (((fi28 - fi28)) | 1))) {
                                                var v58 int32 = -554
                                                switch v2 {
                                                case "foo bar":
                                                    fmt.Println(v2, v58, v27)
                                                }
                                                fmt.Println(v4, v5, v58)
                                                _ = v58
                                            } else {
                                                v3 = append(v3, fi28)
                                                switch fi42 {
                                                case -9223372036854775808:
                                                    v6[v2] = 42
                                                }
                                            }
                                            v27 -= 0
                                            var v59 int = 0
                                            v60 := v6[v2]
                                            switch "" {
                                            case "foo bar":
                                                v2 = "\n\t"
                                            }
                                            _ = v59
                                            _ = v60
                                        }
                                        v27 = (v5 - 405)
                                        var v61 float64 = 3.14159
                                        _ = v61
                                    }
                                } else {
                                    v62 := GenS7{F0: 42}
                                    _ = v62
                                }
                                _ = v44
                            }
                            v27 = (v27 << ((65535) & 63))
                            _ = fi42
                        }
                    }
                    _ = v40
                }
                _ = fi28
            }
            _ = v27
        } else {
            v6[(v2 + v2)] = 0
            if false {
                v1 = 0.0
                _ = len(v3)
            } else {
                var v63 int32 = -293
                v1 -= (v1 + v1)
                _ = v63
            }
        }
    }
    v4 = 42
    v64 := GenS8{F0: (v4 + (^(v4)))}
    for fi65 := 0; fi65 < 3; fi65++ {
        v5 += v5
        _ = len(v3)
        v66 := GenS6{F0: 65535, F1: 2147483647}
        _ = fi65
        _ = v66
    }
    if false {
        v2 = "foo bar"
        for fi67 := 0; fi67 < 7; fi67++ {
            _ = len(v6)
            fmt.Println(fi67, v2)
            switch "foo bar" {
            case "foo bar":
                if (false && false) {
                    switch v2 {
                    case "":
                        v68 := GenS5{F0: 543}
                        _ = v68
                    case "\n\t":
                        for fi69 := 0; fi69 < 4; fi69++ {
                            var v70 uint8 = 127
                            _ = fi69
                            _ = v70
                        }
                    case "a":
                        v5 = 51
                    default:
                        var v71 float64 = 0.0
                        _ = v71
                    }
                    v72 := GenS9{F0: 2, F1: false}
                    fmt.Println(v1, v2)
                    _ = v72
                } else {
                    fmt.Println(v4, v1)
                    for fi73 := 0; fi73 < 5; fi73++ {
                        if false {
                            if true {
                                v74 := GenS5{F0: 167}
                                v5 = (v5 | v5)
                                v4 = (v4 ^ v4)
                                _ = v74
                            }
                        } else {
                            if true {
                                switch 42 {
                                case -9223372036854775808:
                                    if false {
                                        v4 -= v4
                                        switch -1 {
                                        case 2147483647:
                                            v2 = (v2 + v2)
                                        }
                                        v75 := GenS10{F0: v5, F1: -1, F2: 255}
                                        _ = v75
                                    }
                                case 330:
                                    var v76 uint16 = 1
                                    switch -2147483648 {
                                    case 32768:
                                        v2 = (v2 + "foo bar")
                                    }
                                    v1 -= (v1 * float32(((0.0 + (0.1 * ((3.14159 + 1e10) / 1.5))) * (((((((-2.25 / 1.0) + 0.0) - 0.0) + (1e10 / 0.1)) * 3.14159) - (0.1 - 0.1)) - -2.25))))
                                    _ = v76
                                case -2147483648:
                                    _ = len(v3)
                                    if true {
                                        v77 := GenS11{F0: 0, F1: 42}
                                        _ = len(v3)
                                        var v78 int64 = -971
                                        if len(v3) > 0 {
                                            v3[0] = fi73
                                        }
                                        _ = v77
                                        _ = v78
                                    }
                                }
                            }
                            v2 = (v2 + "foo bar")
                            switch 128 {
                            case 2147483647:
                                _ = len(v3)
                            }
                        }
                        v79 := GenS12{F0: 0}
                        _ = fi73
                        _ = v79
                    }
                    v2 = "a"
                }
            case "\n\t":
                var v80 bool = (((!false) && (!true)) && (int8(fi67) == 42))
                _ = v80
            case "a":
                if (!true) {
                    var v81 uint64 = 2
                    fmt.Println(v4, v4, v81)
                    _ = v81
                }
            }
            _ = fi67
        }
        switch ((v2 + ((("a" + ((("hello" + v2) + "") + ("a" + "a"))) + v2) + "a")) + "") {
        case "a":
            for fi82 := 0; fi82 < 11; fi82++ {
                if (((v1 + 1e10) / (1.5 + (((v1 + v1) - (1e10 * (3.14159 * (v1 - (v1 - ((0.0 + v1) / (v1 - v1))))))) - -2.25))) >= v1) {
                    if (!true) {
                        for fi83 := 0; fi83 < 18; fi83++ {
                            switch (v2 + v2) {
                            case "hello":
                                var v84 float32 = 0.1
                                _ = v84
                            case "":
                                v2 = "foo bar"
                            case "\n\t":
                                var v85 int8 = 0
                                _ = v85
                            }
                            switch fi83 {
                            case -9223372036854775808:
                                fmt.Println(v5)
                            }
                            var v86 int64 = 0
                            _ = fi83
                            _ = v86
                        }
                        var v87 bool = (2147483647 != 999)
                        if (!(v87 || (-2147483648 != -675))) {
                            for fi88 := 0; fi88 < 8; fi88++ {
                                v89 := GenS13{F0: v2, F1: (v4 & (^(v4))), F2: 0}
                                var v90 uint32 = 2
                                fmt.Println(v1, v87, v4)
                                v5 -= v5
                                switch ("hello" + (v2 + v2)) {
                                case "\n\t":
                                    var v91 int = fi88
                                    _ = v91
                                case "hello":
                                    v92 := GenS14{F0: -2147483648, F1: 2147483648}
                                    v93 := GenS1{F0: ((v2 + v2) + ""), F1: 289, F2: 2}
                                    if (("foo bar" == ("foo bar" + v2)) || false) {
                                        v87 = v87
                                    } else {
                                        switch 0 {
                                        case 42:
                                            fmt.Println(v2)
                                        case 0:
                                            v94 := GenS15{F0: (v90 | v90), F1: v2}
                                            _ = v94
                                        case 118:
                                            switch (fi88 & (fi88 + 42)) {
                                            case -584:
                                                v4 = (v4 << ((65535) & 63))
                                                fmt.Println(v5, v2, fi82)
                                            case 2147483647:
                                                var v95 int8 = 1
                                                _ = v95
                                            }
                                        }
                                        v5 = (v5 ^ -32768)
                                    }
                                    _ = v92
                                    _ = v93
                                default:
                                    if (0 > 2) {
                                        if len(v3) > 0 {
                                            v3[0] = fi82
                                        }
                                        var v96 uint = 363
                                        switch v2 {
                                        case "a":
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        case "\n\t":
                                            switch fi88 {
                                            case -622:
                                                if true {
                                                    v87 = v87
                                                    switch (fi82 * (fi82 / ((127) | 1))) {
                                                    case -1:
                                                        var v97 string = "\n\t"
                                                        _ = v97
                                                    case 138:
                                                        v5 = (v5 << ((v96) & 63))
                                                        delete(v6, (v2 + v2))
                                                        switch ((v2 + v2) + (("foo bar" + "\n\t") + ((v2 + (v2 + v2)) + (v2 + v2)))) {
                                                        case "":
                                                            var v98 int8 = 100
                                                            _ = v98
                                                        }
                                                    case 42:
                                                        var v99 int16 = v5
                                                        _ = v99
                                                    default:
                                                        if (3.14159 <= 0.1) {
                                                            switch "hello" {
                                                            case "foo bar":
                                                                var v100 int32 = 477
                                                                _ = v100
                                                            case "hello":
                                                                switch (fi82 + fi88) {
                                                                case 32768:
                                                                    var v101 int32 = 2147483647
                                                                    _ = v101
                                                                case 1:
                                                                    var v102 int = 4294967295
                                                                    _ = v102
                                                                default:
                                                                    v96 = (v96 % (((v96 ^ (v96 >> (((^(v96))) & 63)))) | 1))
                                                                }
                                                            case "a":
                                                                if len(v3) > 0 {
                                                                    v3[0] = fi82
                                                                }
                                                            default:
                                                                if (((32768 >= fi82) && ((-1 <= -89) || v87)) || v87) {
                                                                    fmt.Println(fi82)
                                                                    fmt.Println(v2)
                                                                    var v103 int8 = -128
                                                                    _ = v103
                                                                }
                                                            }
                                                            if ((v5 % ((v5) | 1)) != v5) {
                                                                var v104 uint = v96
                                                                _ = v104
                                                            } else {
                                                                var v105 uint32 = v90
                                                                _ = v105
                                                            }
                                                            v106 := GenS6{F0: (v96 * (v96 - (^(v96)))), F1: -803}
                                                            _ = v106
                                                        } else {
                                                            var v107 bool = v87
                                                            fmt.Println(fi82, v2, v87)
                                                            _ = v107
                                                        }
                                                    }
                                                } else {
                                                    v4 += (v4 % ((v4) | 1))
                                                }
                                            default:
                                                var v108 bool = v87
                                                fmt.Println(v96, v87, v108)
                                                var v109 int32 = 0
                                                _ = v108
                                                _ = v109
                                            }
                                        case "foo bar":
                                            if len(v3) > 0 {
                                                v3[0] = fi88
                                            }
                                        default:
                                            if len(v3) > 0 {
                                                v3[0] = (fi82 + fi88)
                                            }
                                            switch 1 {
                                            case -2147483648:
                                                var v110 float64 = 0.0
                                                _ = v110
                                            }
                                        }
                                        _ = v96
                                    } else {
                                        v111 := GenS13{F0: v2, F1: (v4 - 204), F2: 21}
                                        v112 := GenS16{F0: 276}
                                        v1 = v1
                                        _ = v111
                                        _ = v112
                                    }
                                }
                                _ = fi88
                                _ = v89
                                _ = v90
                            }
                            var v113 uint16 = 255
                            _ = v113
                        } else {
                            switch fi82 {
                            case 32768:
                                _ = len(v3)
                            case 2147483647:
                                var v114 int16 = (^(v5))
                                _ = v114
                            }
                            v115 := GenS11{F0: (v4 % ((v4) | 1)), F1: v4}
                            _ = v115
                        }
                        _ = v87
                    } else {
                        v116 := GenS12{F0: -9223372036854775808}
                        var v117 int8 = 0
                        _ = v116
                        _ = v117
                    }
                    for fi118 := 0; fi118 < 4; fi118++ {
                        v2 = (((v2 + (v2 + (v2 + "foo bar"))) + ((v2 + ((v2 + v2) + v2)) + v2)) + v2)
                        if false {
                            _ = len(v3)
                            switch (fi82 + 1) {
                            case -1:
                                fmt.Println(v5, v1)
                            case 4294967295:
                                var v119 int16 = v5
                                v4 += 1
                                _ = v119
                            case 246:
                                switch "\n\t" {
                                case "hello":
                                    if len(v3) > 0 {
                                        v3[0] = (fi82 / ((fi118) | 1))
                                    }
                                default:
                                    fmt.Println(v4, fi118)
                                }
                            }
                            switch v2 {
                            case "foo bar":
                                var v120 int32 = 42
                                _ = v120
                            case "hello":
                                v121 := GenS17{F0: (v2 + v2), F1: fi82}
                                switch (-(fi82)) {
                                case 557:
                                    fmt.Println(fi118)
                                case -49:
                                    if false {
                                        var v122 float64 = -2.25
                                        _ = v122
                                    } else {
                                        var v123 float32 = ((0.1 - ((1e10 - 1.5) + ((v1 + v1) / v1))) / v1)
                                        switch (v2 + v2) {
                                        case "foo bar":
                                            var v124 int32 = 42
                                            _ = v124
                                        }
                                        var v125 uint = 65535
                                        _ = v123
                                        _ = v125
                                    }
                                case 1:
                                    fmt.Println(fi118)
                                default:
                                    var v126 uint64 = 1
                                    _ = v126
                                }
                                v5 -= (v5 & (^(v5)))
                                v5 += (^(v5))
                                _ = v121
                            default:
                                fmt.Println(v5, v5)
                            }
                            v5 += 10
                        } else {
                            fmt.Println(fi118, v2, v2)
                        }
                        _ = fi118
                    }
                    fmt.Println(v1, v1, v2)
                } else {
                    v3 = append(v3, fi82)
                    var v127 int32 = int32(v4)
                    _ = v127
                }
                _ = fi82
            }
        case "hello":
            for fi128 := 0; fi128 < 8; fi128++ {
                var v129 uint = uint(v4)
                delete(v6, ((v2 + v2) + (v2 + (v2 + v2))))
                _ = fi128
                _ = v129
            }
        case "\n\t":
            v2 = v2
        default:
            var v130 uint16 = 65535
            _ = v130
        }
    }
    var v131 int8 = 1
    switch -2147483648 {
    case 32768:
        v132 := GenS3{F0: 32767, F1: 324, F2: 98}
        _ = v132
    case 4294967295:
        var v133 float64 = 0.1
        _ = v133
    default:
        for fi134 := 0; fi134 < 8; fi134++ {
            var v135 uint64 = 42
            v6[v2] = 2147483647
            _ = fi134
            _ = v135
        }
    }
    for fi136 := 0; fi136 < 17; fi136++ {
        var v137 uint8 = v4
        _ = fi136
        _ = v137
    }
    v5 -= (v5 >> ((184) & 63))
    for fi138 := 0; fi138 < 1; fi138++ {
        var v139 int64 = 9223372036854775807
        v140 := v6[(v2 + v2)]
        var v141 int32 = -2147483648
        _ = fi138
        _ = v139
        _ = v140
        _ = v141
    }
    fmt.Println("v1:", v1)
    fmt.Println("v131:", v131)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v131
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v64
}
