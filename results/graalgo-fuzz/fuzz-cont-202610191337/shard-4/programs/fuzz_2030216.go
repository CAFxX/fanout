package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 int; F1 int; F2 uint32 }
type GenS3 struct { F0 uint8; F1 int32; F2 int8 }
type GenS4 struct { F0 uint64; F1 uint }
type GenS5 struct { F0 uint; F1 int16 }
type GenS6 struct { F0 uint8; F1 int16 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int }

func main() {
    var v1 int8 = -128
    v2 := []uint32{1, 293, 2}
    v3 := []int{717, -1, 32768, 127}
    v4 := make(map[uint32]string)
    var v5 int = 128
    switch "a" {
    case "":
        fmt.Println(v1, v5)
    case "a":
        v6 := GenS1{F0: -1, F1: 4294967295}
        _ = v6
    }
    for fi7 := 0; fi7 < 4; fi7++ {
        v1 = (v1 ^ (v1 << ((599) & 63)))
        v5 += 196
        _ = fi7
    }
    for fi8 := 0; fi8 < 6; fi8++ {
        v5 = (v5 / ((9223372036854775807) | 1))
        _ = fi8
    }
    var v9 uint32 = 257
    switch "hello" {
    case "foo bar":
        v10 := GenS1{F0: 1, F1: 2}
        _ = v10
    case "hello":
        var v11 int8 = -1
        _ = v11
    case "\n\t":
        var v12 bool = (true && (float64(float32(1e10)) != 0.1))
        _ = v12
    default:
        var v13 int16 = -959
        _ = v13
    }
    switch -9223372036854775808 {
    case 127:
        fmt.Println(v1, v5, v1)
    case -996:
        for fi14 := 0; fi14 < 17; fi14++ {
            switch "hello" {
            case "":
                var v15 uint8 = 205
                _ = v15
            case "a":
                v1 -= (v1 * v1)
            case "hello":
                _ = len(v4)
            }
            for fi16 := 0; fi16 < 5; fi16++ {
                v17 := GenS1{F0: -2147483648, F1: uint64(v9)}
                v1 = v1
                _ = fi16
                _ = v17
            }
            if (false && ((2 > (^(v9))) && (true || true))) {
                for fi18 := 0; fi18 < 8; fi18++ {
                    _ = len(v4)
                    switch ("" + "\n\t") {
                    case "":
                        v9 = v9
                    }
                    if len(v2) > 0 {
                        v2[0] = (^(v9))
                    }
                    v9 -= (^(v9))
                    _ = fi18
                }
            }
            _ = fi14
        }
    }
    for fi19 := 0; fi19 < 8; fi19++ {
        var v20 float32 = 3.14159
        for fi21 := 0; fi21 < 13; fi21++ {
            if true {
                switch 42 {
                case 127:
                    if ((v1 % ((42) | 1)) <= 48) {
                        if true {
                            v22 := GenS1{F0: -220, F1: 1}
                            if true {
                                v20 = v20
                                delete(v4, 1)
                            }
                            _ = v22
                        }
                        v23 := v4[v9]
                        v24 := GenS1{F0: 42, F1: 0}
                        _ = v23
                        _ = v24
                    }
                case -1:
                    var v25 uint16 = 192
                    _ = v25
                case 128:
                    fmt.Println(fi19)
                }
            }
            _ = fi21
        }
        _ = fi19
        _ = v20
    }
    var v26 bool = false
    if v26 {
        switch v5 {
        case 9223372036854775807:
            switch (v5 + v5) {
            case 201:
                if true {
                    for fi27 := 0; fi27 < 13; fi27++ {
                        if v26 {
                            v3 = append(v3, fi27)
                            v5 += (v5 - (v5 * -1))
                            switch "" {
                            case "a":
                                var v28 uint16 = 1
                                _ = v28
                            }
                        } else {
                            v1 -= 0
                            _ = len(v4)
                        }
                        v29 := GenS1{F0: 1, F1: 1}
                        fmt.Println(fi27)
                        _ = fi27
                        _ = v29
                    }
                    switch ("foo bar" + "") {
                    case "hello":
                        switch v5 {
                        case -2147483648:
                            switch (^(v5)) {
                            case -1:
                                v1 = v1
                            case 128:
                                v30 := v4[881]
                                _ = v30
                            }
                        default:
                            v31 := GenS1{F0: -386, F1: 18446744073709551615}
                            _ = v31
                        }
                    case "foo bar":
                        var v32 uint16 = 32767
                        _ = v32
                    default:
                        var v33 int = (v5 * v5)
                        _ = v33
                    }
                }
            default:
                if (!(!((657 >= 0) && true))) {
                    if (!(!v26)) {
                        if len(v2) > 0 {
                            v2[0] = v9
                        }
                    } else {
                        v34 := GenS1{F0: int32(v1), F1: 2}
                        _ = v34
                    }
                } else {
                    v35 := v4[v9]
                    delete(v4, v9)
                    v35 = v35
                    _ = v35
                }
            }
        case 32768:
            v36 := GenS1{F0: 0, F1: 806}
            _ = v36
        default:
            for fi37 := 0; fi37 < 5; fi37++ {
                v5 = (^(v5))
                var v38 string = ""
                v39 := GenS2{F0: (v5 & fi37), F1: -2147483648, F2: (v9 * (v9 << ((456) & 63)))}
                _ = fi37
                _ = v38
                _ = v39
            }
        }
        for fi40 := 0; fi40 < 18; fi40++ {
            for fi41 := 0; fi41 < 7; fi41++ {
                v5 -= (^(v5))
                var v42 int8 = v1
                _ = fi41
                _ = v42
            }
            _ = fi40
        }
        if (42 >= v5) {
            if ((((v26 || (2 >= 2)) || false) && v26) && v26) {
                if len(v3) > 0 {
                    v3[0] = -1
                }
                if (v9 != (^(v9))) {
                    v43 := GenS1{F0: -2147483648, F1: 2}
                    var v44 int8 = 42
                    v4[(v9 + v9)] = "hello"
                    _ = v43
                    _ = v44
                } else {
                    for fi45 := 0; fi45 < 3; fi45++ {
                        var v46 float32 = float32(3.14159)
                        v4[2147483648] = "hello"
                        var v47 uint64 = 1
                        _ = fi45
                        _ = v46
                        _ = v47
                    }
                    for fi48 := 0; fi48 < 11; fi48++ {
                        v49 := GenS2{F0: fi48, F1: 42, F2: v9}
                        _ = fi48
                        _ = v49
                    }
                }
            }
            switch (v5 ^ v5) {
            case 32767:
                if false {
                    var v50 uint64 = 0
                    v9 = (v9 * v9)
                    _ = v50
                } else {
                    var v51 int32 = 1
                    v1 = (-(v1))
                    _ = v51
                }
            case -543:
                var v52 int16 = -32768
                _ = v52
            case 42:
                v4[(v9 | v9)] = "a"
            }
        } else {
            if v26 {
                v4[v9] = "hello"
                var v53 int64 = -1
                v5 -= v5
                _ = v53
            } else {
                fmt.Println(v26, v1)
                for fi54 := 0; fi54 < 11; fi54++ {
                    switch "hello" {
                    case "":
                        v9 += v9
                    case "hello":
                        var v55 uint16 = 32767
                        _ = v55
                    }
                    for fi56 := 0; fi56 < 18; fi56++ {
                        v1 = v1
                        _ = fi56
                    }
                    _ = fi54
                }
                for fi57 := 0; fi57 < 17; fi57++ {
                    if false {
                        for fi58 := 0; fi58 < 14; fi58++ {
                            switch 805 {
                            case 32768:
                                delete(v4, v9)
                            case -9223372036854775808:
                                var v59 float32 = float32((3.14159 / (((0.1 / 3.14159) - 1.5) - 0.1)))
                                _ = v59
                            case 619:
                                fmt.Println(fi58, v9)
                            default:
                                var v60 int16 = -1
                                _ = v60
                            }
                            if v26 {
                                var v61 uint16 = 65535
                                _ = v61
                            } else {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                v1 += (v1 | 41)
                            }
                            _ = fi58
                        }
                        v26 = v26
                    }
                    _ = fi57
                }
            }
            for fi62 := 0; fi62 < 2; fi62++ {
                switch (((("\n\t" + "foo bar") + "a") + "hello") + "foo bar") {
                case "hello":
                    var v63 uint8 = 147
                    _ = v63
                case "foo bar":
                    v64 := GenS1{F0: 2147483647, F1: 149}
                    _ = v64
                case "":
                    if false {
                        v5 -= int(v9)
                        var v65 int64 = 1
                        switch (v5 << ((2) & 63)) {
                        case 995:
                            v66 := GenS2{F0: (fi62 * int(v9)), F1: 2147483647, F2: 205}
                            _ = v66
                        case -2147483648:
                            if (32767 < fi62) {
                                fmt.Println(v1)
                            } else {
                                var v67 int64 = -299
                                var v68 int64 = v67
                                switch (v5 / ((32768) | 1)) {
                                case -817:
                                    for fi69 := 0; fi69 < 6; fi69++ {
                                        v70 := GenS2{F0: (-(fi69)), F1: -9223372036854775808, F2: v9}
                                        switch ("foo bar" + "foo bar") {
                                        case "\n\t":
                                            if (v1 < -1) {
                                                if (v9 < 158) {
                                                    if v26 {
                                                        v65 = (v67 ^ 9223372036854775807)
                                                        v5 += fi69
                                                        fmt.Println(v68, v5, v9)
                                                        v71 := GenS3{F0: 255, F1: -383, F2: v1}
                                                        v9 -= 771
                                                        _ = v71
                                                    } else {
                                                        if true {
                                                            v67 -= v68
                                                            if len(v2) > 0 {
                                                                v2[0] = v9
                                                            }
                                                            fmt.Println(v65, v67)
                                                        } else {
                                                            if false {
                                                                switch (fi69 & (fi62 << ((2) & 63))) {
                                                                case 32768:
                                                                    fmt.Println(v68, v5)
                                                                default:
                                                                    fmt.Println(fi69, v9, v65)
                                                                }
                                                                fmt.Println(v65, v67)
                                                                if v26 {
                                                                    if (68 < 2) {
                                                                        v67 = 0
                                                                        v72 := GenS4{F0: 145, F1: 18446744073709551615}
                                                                        var v73 uint8 = 45
                                                                        v74 := GenS5{F0: 46, F1: 1}
                                                                        _ = v72
                                                                        _ = v73
                                                                        _ = v74
                                                                    } else {
                                                                        v75 := GenS2{F0: 9223372036854775807, F1: (v5 ^ 32768), F2: v9}
                                                                        _ = v75
                                                                    }
                                                                    v1 = 88
                                                                    v3 = append(v3, v5)
                                                                    v65 = -9223372036854775808
                                                                }
                                                                if (18446744073709551615 >= 18446744073709551615) {
                                                                    v67 = v67
                                                                    var v76 uint16 = 65535
                                                                    _ = v76
                                                                }
                                                            }
                                                            fmt.Println(fi69, fi69, fi62)
                                                            if true {
                                                                if v26 {
                                                                    var v77 int16 = -32768
                                                                    if true {
                                                                        v78 := GenS5{F0: 4294967295, F1: 1000}
                                                                        var v79 uint = 1
                                                                        v80 := GenS6{F0: uint8(v79), F1: v77}
                                                                        fmt.Println(v1)
                                                                        fmt.Println(v68, fi69)
                                                                        _ = v78
                                                                        _ = v79
                                                                        _ = v80
                                                                    }
                                                                    var v81 bool = (0 == (v1 & v1))
                                                                    _ = v77
                                                                    _ = v81
                                                                } else {
                                                                    v1 -= v1
                                                                    fmt.Println(v68, v68)
                                                                    var v82 uint64 = 354
                                                                    _ = v82
                                                                }
                                                            } else {
                                                                if v26 {
                                                                    var v83 uint32 = v9
                                                                    _ = v83
                                                                } else {
                                                                    var v84 int8 = 111
                                                                    fmt.Println(v65)
                                                                    v85 := GenS3{F0: 255, F1: 2147483647, F2: (^(v1))}
                                                                    _ = v84
                                                                    _ = v85
                                                                }
                                                                fmt.Println(v68, fi69, v5)
                                                            }
                                                        }
                                                        var v86 uint16 = 42
                                                        var v87 string = ""
                                                        fmt.Println(v65, v5, fi62)
                                                        v67 = (v68 | (v68 ^ 42))
                                                        _ = v86
                                                        _ = v87
                                                    }
                                                    if v26 {
                                                        var v88 uint16 = 1
                                                        if (((1 <= 0) || ((65535 == 306) || ("foo bar" > ("\n\t" + "hello")))) && v26) {
                                                            switch "" {
                                                            case "foo bar":
                                                                var v89 uint = 0
                                                                v1 = 127
                                                                _ = v89
                                                            case "a":
                                                                var v90 uint = 2
                                                                _ = v90
                                                            default:
                                                                var v91 uint64 = 0
                                                                _ = v91
                                                            }
                                                        }
                                                        _ = v88
                                                    } else {
                                                        switch 577 {
                                                        case 2147483647:
                                                            fmt.Println(fi69, v1)
                                                            var v92 float32 = (1.5 * -2.25)
                                                            _ = v92
                                                        case -581:
                                                            if v26 {
                                                                var v93 uint64 = 1
                                                                v94 := GenS3{F0: 128, F1: 1, F2: 0}
                                                                var v95 float32 = 3.14159
                                                                _ = v93
                                                                _ = v94
                                                                _ = v95
                                                            } else {
                                                                if v26 {
                                                                    v2 = append(v2, v9)
                                                                } else {
                                                                    delete(v4, v9)
                                                                }
                                                                v96 := GenS6{F0: 128, F1: 247}
                                                                switch -120 {
                                                                case -2147483648:
                                                                    switch (fi62 % (((v5 * (v5 >> ((2) & 63)))) | 1)) {
                                                                    case 9223372036854775807:
                                                                        if (42 == (v9 / ((v9) | 1))) {
                                                                            var v97 string = (("a" + "a") + "\n\t")
                                                                            var v98 int = fi69
                                                                            var v99 int16 = -1
                                                                            _ = v97
                                                                            _ = v98
                                                                            _ = v99
                                                                        }
                                                                        var v100 int16 = 1
                                                                        fmt.Println(v68, v67, v9)
                                                                        _ = v100
                                                                    case -1:
                                                                        if (!false) {
                                                                            v1 -= v1
                                                                            switch fi69 {
                                                                            case 32767:
                                                                                var v101 uint32 = 2147483647
                                                                                _ = v101
                                                                            case 9223372036854775807:
                                                                                var v102 uint16 = 255
                                                                                _ = v102
                                                                            default:
                                                                                v65 -= (v68 << ((961) & 63))
                                                                            }
                                                                        } else {
                                                                            fmt.Println(v1, v1, v68)
                                                                            v65 = (-(v68))
                                                                        }
                                                                        switch fi62 {
                                                                        case 127:
                                                                            v65 -= (v67 & 0)
                                                                        case 341:
                                                                            v1 = (v1 - -89)
                                                                        default:
                                                                            v103 := GenS3{F0: 0, F1: -334, F2: v1}
                                                                            _ = v103
                                                                        }
                                                                    default:
                                                                        if v26 {
                                                                            fmt.Println(v67)
                                                                        } else {
                                                                            var v104 uint16 = 65535
                                                                            _ = v104
                                                                        }
                                                                    }
                                                                default:
                                                                    switch (fi62 | fi62) {
                                                                    case 1:
                                                                        var v105 int64 = 42
                                                                        _ = v105
                                                                    }
                                                                }
                                                                _ = v96
                                                            }
                                                        case 32767:
                                                            var v106 uint16 = 22
                                                            _ = v106
                                                        }
                                                        var v107 uint8 = 38
                                                        var v108 uint64 = 1
                                                        _ = v107
                                                        _ = v108
                                                    }
                                                }
                                            } else {
                                                if v26 {
                                                    fmt.Println(fi62)
                                                }
                                                v67 = (v67 - (v67 >> ((0) & 63)))
                                            }
                                        case "":
                                            v2 = append(v2, v9)
                                        }
                                        _ = fi69
                                        _ = v70
                                    }
                                }
                                _ = v67
                                _ = v68
                            }
                        }
                        _ = v65
                    } else {
                        _ = len(v4)
                    }
                default:
                    var v109 uint8 = 101
                    _ = v109
                }
                v9 = v9
                if v26 {
                    var v110 int32 = 651
                    switch 212 {
                    case 32768:
                        delete(v4, 4294967295)
                    case 127:
                        var v111 uint8 = 127
                        _ = v111
                    default:
                        if v26 {
                            switch (v5 * (v5 >> ((4294967295) & 63))) {
                            case -388:
                                switch v5 {
                                case 42:
                                    var v112 uint8 = 0
                                    _ = v112
                                }
                            case 314:
                                v113 := GenS7{F0: false}
                                _ = v113
                            }
                        } else {
                            fmt.Println(v9, v26, v5)
                            var v114 float64 = 0.0
                            for fi115 := 0; fi115 < 1; fi115++ {
                                v116 := v4[(v9 / ((v9) | 1))]
                                var v117 bool = true
                                _ = fi115
                                _ = v116
                                _ = v117
                            }
                            _ = v114
                        }
                    }
                    _ = v110
                } else {
                    for fi118 := 0; fi118 < 2; fi118++ {
                        _ = len(v4)
                        var v119 int64 = 9223372036854775807
                        _ = fi118
                        _ = v119
                    }
                    for fi120 := 0; fi120 < 5; fi120++ {
                        var v121 bool = (18446744073709551615 > 42)
                        _ = fi120
                        _ = v121
                    }
                }
                _ = fi62
            }
        }
    } else {
        switch "a" {
        case "\n\t":
            var v122 int = v5
            _ = v122
        case "hello":
            v26 = (255 < 156)
        default:
            var v123 uint = 0
            _ = v123
        }
    }
    if false {
        var v124 float32 = float32(0.0)
        _ = v124
    } else {
        var v125 int8 = v1
        for fi126 := 0; fi126 < 16; fi126++ {
            var v127 int32 = -1
            _ = fi126
            _ = v127
        }
        v128 := GenS8{F0: 562}
        _ = v125
        _ = v128
    }
    switch "hello" {
    case "hello":
        delete(v4, (v9 * (^(v9))))
    }
    v129 := GenS2{F0: v5, F1: (v5 * v5), F2: (v9 - v9)}
    for fi130 := 0; fi130 < 7; fi130++ {
        v9 += v9
        _ = fi130
    }
    v131 := v4[v9]
    if ((v9 & v9) < 442) {
        v26 = ((v26 && v26) && v26)
    } else {
        var v132 uint64 = 42
        _ = v132
    }
    fmt.Println("v1:", v1)
    fmt.Println("v131:", v131)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v129
    _ = v131
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
