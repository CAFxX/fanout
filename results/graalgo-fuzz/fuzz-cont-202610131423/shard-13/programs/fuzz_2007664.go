package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 int }
type GenS2 struct { F0 uint; F1 int8; F2 uint64 }
type GenS3 struct { F0 int8; F1 string; F2 int64 }
type GenS4 struct { F0 int32; F1 uint8; F2 int32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int; F1 uint16; F2 int8 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint; F1 uint8 }
type GenS9 struct { F0 int8; F1 int8; F2 uint64 }
type GenS10 struct { F0 uint8; F1 int32 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 string; F1 int16 }
type GenS13 struct { F0 int32; F1 int8; F2 int32 }
type GenS14 struct { F0 uint16; F1 uint64 }
type GenS15 struct { F0 uint8 }
type GenS16 struct { F0 string }
type GenS17 struct { F0 int }
type GenS18 struct { F0 uint64 }

func main() {
    var v1 int8 = -128
    var v2 uint8 = 42
    var v3 int32 = 42
    v4 := make(map[int]int)
    var v5 float64 = float64((1.5 / (0.1 + 0.1)))
    for fi6 := 0; fi6 < 20; fi6++ {
        v1 += (v1 >> ((850) & 63))
        _ = fi6
    }
    switch "hello" {
    case "\n\t":
        v1 = (v1 >> ((2) & 63))
    case "a":
        if false {
            switch 4294967295 {
            case 791:
                for fi7 := 0; fi7 < 5; fi7++ {
                    for fi8 := 0; fi8 < 11; fi8++ {
                        fmt.Println(v2)
                        if false {
                            fmt.Println(fi8)
                            v9 := GenS1{F0: 531, F1: (v1 % (((v1 ^ v1)) | 1)), F2: fi7}
                            if true {
                                v5 = 1.5
                                if true {
                                    switch (^(fi8)) {
                                    case 128:
                                        v10 := GenS1{F0: 426, F1: v1, F2: (fi8 | (-(fi8)))}
                                        _ = v10
                                    case 2147483647:
                                        fmt.Println(v1, v1, fi8)
                                        v5 = ((v5 - float64(0.1)) - ((0.1 - 3.14159) + v5))
                                        v1 -= (v1 & (v1 / ((0) | 1)))
                                    case -9223372036854775808:
                                        _ = len(v4)
                                    default:
                                        var v11 int32 = 42
                                        _ = v11
                                    }
                                } else {
                                    if (42 <= 0) {
                                        var v12 uint64 = 42
                                        v2 += 58
                                        switch (("a" + (("" + "a") + ("\n\t" + ("a" + ("" + "\n\t"))))) + "foo bar") {
                                        case "a":
                                            v13 := GenS1{F0: 4294967295, F1: v1, F2: fi7}
                                            _ = v13
                                        default:
                                            v14 := GenS1{F0: 4294967295, F1: (v1 ^ (^(v1))), F2: (-(fi8))}
                                            _ = v14
                                        }
                                        v5 += -2.25
                                        if true {
                                            var v15 int16 = -1
                                            if true {
                                                fmt.Println(v3, v15, fi7)
                                                if (1 < 9223372036854775807) {
                                                    fmt.Println(fi7, v12, v5)
                                                    fmt.Println(v5, v1, v5)
                                                    v16 := GenS1{F0: 441, F1: v1, F2: (fi7 << ((65535) & 63))}
                                                    var v17 uint16 = 1
                                                    _ = v16
                                                    _ = v17
                                                } else {
                                                    var v18 int16 = v15
                                                    _ = v18
                                                }
                                            } else {
                                                var v19 float64 = (-2.25 * 3.14159)
                                                _ = v19
                                            }
                                            if false {
                                                switch "hello" {
                                                case "\n\t":
                                                    fmt.Println(v3, v2, v5)
                                                default:
                                                    v20 := GenS1{F0: 1, F1: (v1 ^ (v1 | v1)), F2: 32767}
                                                    _ = v20
                                                }
                                                var v21 uint64 = (v12 % ((450) | 1))
                                                if ((v15 & v15) < (v15 + v15)) {
                                                    fmt.Println(v15, v3)
                                                } else {
                                                    var v22 uint64 = v21
                                                    var v23 uint64 = (v12 ^ v21)
                                                    _ = v22
                                                    _ = v23
                                                }
                                                _ = v21
                                            } else {
                                                if false {
                                                    fmt.Println(fi8)
                                                    var v24 int = fi8
                                                    switch "hello" {
                                                    case "":
                                                        v25 := GenS1{F0: 65535, F1: (v1 >> ((65535) & 63)), F2: (^(fi7))}
                                                        _ = v25
                                                    case "foo bar":
                                                        var v26 int = (v24 * fi7)
                                                        _ = v26
                                                    default:
                                                        v27 := GenS1{F0: 253, F1: int8(v3), F2: (v24 & fi8)}
                                                        _ = v27
                                                    }
                                                    switch -1 {
                                                    case 654:
                                                        v28 := GenS1{F0: 1, F1: (^(v1)), F2: fi7}
                                                        var v29 float32 = 0.1
                                                        if false {
                                                            var v30 int16 = (v15 << ((2) & 63))
                                                            v3 += int32(v1)
                                                            _ = v30
                                                        } else {
                                                            fmt.Println(v29, v24)
                                                        }
                                                        _ = v28
                                                        _ = v29
                                                    default:
                                                        fmt.Println(v15, v2)
                                                    }
                                                    v1 += v1
                                                    _ = v24
                                                } else {
                                                    if (1.5 == 1e10) {
                                                        fmt.Println(v2)
                                                        v31 := GenS1{F0: 1, F1: (v1 + -128), F2: fi7}
                                                        v2 = (v2 >> ((2) & 63))
                                                        v32 := []int{fi7, (fi8 ^ fi7)}
                                                        _ = v31
                                                        _ = v32
                                                    }
                                                }
                                                switch "" {
                                                case "hello":
                                                    v33 := GenS1{F0: 483, F1: -127, F2: (fi8 >> ((421) & 63))}
                                                    _ = v33
                                                }
                                            }
                                            _ = v15
                                        } else {
                                            var v34 int32 = v3
                                            _ = v34
                                        }
                                        _ = v12
                                    }
                                    var v35 bool = (false || false)
                                    if v35 {
                                        if false {
                                            if (!("a" != ("" + "hello"))) {
                                                v36 := GenS1{F0: 1, F1: (v1 + (^(v1))), F2: fi7}
                                                _ = v36
                                            }
                                        } else {
                                            switch (fi7 - fi8) {
                                            case 32768:
                                                fmt.Println(fi7, fi8, fi7)
                                            case 12:
                                                _ = len(v4)
                                            default:
                                                v3 += v3
                                            }
                                        }
                                    } else {
                                        if (!v35) {
                                            fmt.Println(fi7, v2)
                                            if v35 {
                                                var v37 float64 = 0.1
                                                fmt.Println(v5, fi7)
                                                _ = v37
                                            } else {
                                                v38 := []int{(fi8 & 1), fi7}
                                                if len(v38) > 0 {
                                                    v38[0] = (fi8 / ((fi7) | 1))
                                                }
                                                _ = v38
                                            }
                                            var v39 int16 = 675
                                            _ = v39
                                        }
                                    }
                                    _ = v35
                                }
                            } else {
                                v2 = (v2 | (v2 + v2))
                            }
                            _ = v9
                        } else {
                            var v40 uint32 = 2147483648
                            if (!(false && (904 >= -9223372036854775808))) {
                                switch (^(fi8)) {
                                case 128:
                                    if false {
                                        var v41 float32 = float32(float64((1.5 * 1.5)))
                                        var v42 string = ("a" + "")
                                        if ((v1 * (v1 + 1)) <= int8(v2)) {
                                            var v43 int = -782
                                            _ = v43
                                        } else {
                                            switch 32768 {
                                            case 32768:
                                                var v44 uint64 = 2
                                                _ = v44
                                            case 4294967295:
                                                switch v42 {
                                                case "a":
                                                    if (v42 >= (v42 + (v42 + v42))) {
                                                        if true {
                                                            v3 = v3
                                                            fmt.Println(v5, fi7)
                                                        }
                                                        var v45 float64 = (v5 - ((v5 + v5) + -2.25))
                                                        v46 := GenS1{F0: v40, F1: (v1 - v1), F2: (^(fi7))}
                                                        _ = v45
                                                        _ = v46
                                                    }
                                                }
                                                v40 = (^(v40))
                                            case -683:
                                                v2 = v2
                                            default:
                                                v47 := v4[(fi8 >> ((2) & 63))]
                                                _ = v47
                                            }
                                            _ = len(v4)
                                            v48 := GenS2{F0: 2, F1: v1, F2: 0}
                                            _ = v48
                                        }
                                        switch fi7 {
                                        case 968:
                                            fmt.Println(v40)
                                        case 4294967295:
                                            var v49 int16 = -577
                                            _ = v49
                                        default:
                                            if false {
                                                var v50 float64 = v5
                                                _ = v50
                                            } else {
                                                v51 := GenS3{F0: v1, F1: "\n\t", F2: 1}
                                                switch (fi8 ^ fi8) {
                                                case 128:
                                                    v42 = v42
                                                default:
                                                    var v52 int32 = v3
                                                    fmt.Println(v52, v52)
                                                    v53 := GenS2{F0: 936, F1: (v1 ^ v1), F2: 190}
                                                    _ = v52
                                                    _ = v53
                                                }
                                                fmt.Println(v1)
                                                _ = v51
                                            }
                                        }
                                        var v54 uint = 0
                                        _ = v41
                                        _ = v42
                                        _ = v54
                                    }
                                case 0:
                                    v2 -= v2
                                    v55 := GenS4{F0: -2147483648, F1: (v2 ^ (v2 >> ((814) & 63))), F2: (v3 / (((v3 >> ((718) & 63))) | 1))}
                                    _ = v55
                                default:
                                    switch -556 {
                                    case 2147483647:
                                        v5 -= (-2.25 / v5)
                                    }
                                }
                                v56 := GenS2{F0: 22, F1: (^(v1)), F2: 619}
                                v57 := GenS5{F0: 2}
                                var v58 int64 = 1
                                _ = v56
                                _ = v57
                                _ = v58
                            }
                            _ = v40
                        }
                        _ = fi8
                    }
                    v2 = v2
                    _ = fi7
                }
            case 191:
                if false {
                    fmt.Println(v3)
                } else {
                    if ((42 < 384) && (v1 <= 50)) {
                        for fi59 := 0; fi59 < 3; fi59++ {
                            v1 -= (v1 - (v1 << ((42) & 63)))
                            _ = fi59
                        }
                        if false {
                            switch 42 {
                            case -55:
                                for fi60 := 0; fi60 < 16; fi60++ {
                                    v2 = (v2 % ((v2) | 1))
                                    v4[(fi60 - fi60)] = fi60
                                    _ = fi60
                                }
                            case 483:
                                var v61 uint16 = 434
                                _ = v61
                            default:
                                v2 += 127
                            }
                            v5 -= ((v5 * (v5 + v5)) + (3.14159 * (0.0 - v5)))
                            v4[4294967295] = -1
                        }
                        v62 := GenS6{F0: 239, F1: 980, F2: (v1 + v1)}
                        _ = v62
                    } else {
                        for fi63 := 0; fi63 < 11; fi63++ {
                            v64 := GenS5{F0: 18446744073709551615}
                            v65 := []int{fi63, 77, fi63, fi63}
                            v2 = v2
                            _ = fi63
                            _ = v64
                            _ = v65
                        }
                    }
                }
            }
            switch ("foo bar" + "a") {
            case "a":
                v5 -= v5
            case "hello":
                if true {
                    for fi66 := 0; fi66 < 18; fi66++ {
                        if false {
                            if false {
                                v67 := GenS2{F0: 2, F1: v1, F2: 42}
                                var v68 uint64 = 2
                                _ = v67
                                _ = v68
                            }
                            switch ("foo bar" + ("" + "foo bar")) {
                            case "foo bar":
                                if false {
                                    switch fi66 {
                                    case 2147483647:
                                        if (2 <= 4294967295) {
                                            v69 := []int{fi66, (fi66 ^ fi66), fi66}
                                            var v70 uint = 446
                                            v1 = v1
                                            _ = v69
                                            _ = v70
                                        } else {
                                            switch fi66 {
                                            case 0:
                                                v71 := []int{1, fi66}
                                                _ = v71
                                            case 9223372036854775807:
                                                v72 := GenS7{F0: (-(v3))}
                                                _ = v72
                                            case -9223372036854775808:
                                                for fi73 := 0; fi73 < 2; fi73++ {
                                                    fmt.Println(v2, fi73)
                                                    _ = fi73
                                                }
                                            default:
                                                if (972 < 89) {
                                                    if (!(!(false && (42 != 4294967295)))) {
                                                        if (fi66 != fi66) {
                                                            for fi74 := 0; fi74 < 9; fi74++ {
                                                                v75 := v4[fi66]
                                                                v75 = (fi74 >> ((255) & 63))
                                                                v76 := []int{v75, (fi66 & -9223372036854775808), (-(v75))}
                                                                _ = fi74
                                                                _ = v75
                                                                _ = v76
                                                            }
                                                            for fi77 := 0; fi77 < 4; fi77++ {
                                                                switch 711 {
                                                                case 122:
                                                                    delete(v4, (fi66 - 127))
                                                                case 314:
                                                                    v78 := v4[fi77]
                                                                    _ = v78
                                                                }
                                                                var v79 uint16 = 32767
                                                                switch "" {
                                                                case "foo bar":
                                                                    var v80 bool = true
                                                                    _ = v80
                                                                case "hello":
                                                                    v81 := GenS2{F0: 27, F1: (-(v1)), F2: 0}
                                                                    _ = v81
                                                                case "\n\t":
                                                                    v79 = v79
                                                                }
                                                                v82 := GenS4{F0: v3, F1: 42, F2: (v3 & (v3 % ((549) | 1)))}
                                                                _ = fi77
                                                                _ = v79
                                                                _ = v82
                                                            }
                                                            switch "" {
                                                            case "":
                                                                if true {
                                                                    v83 := GenS4{F0: (v3 + v3), F1: v2, F2: int32(fi66)}
                                                                    v84 := GenS7{F0: v3}
                                                                    v85 := []int{(fi66 / (((fi66 ^ (fi66 - (^(fi66))))) | 1)), 1}
                                                                    _ = v83
                                                                    _ = v84
                                                                    _ = v85
                                                                } else {
                                                                    var v86 int = (fi66 + fi66)
                                                                    _ = v86
                                                                }
                                                            case "hello":
                                                                if (float32((v5 - v5)) <= -2.25) {
                                                                    switch "" {
                                                                    case "a":
                                                                        v2 = (v2 & (v2 & (v2 - 0)))
                                                                    default:
                                                                        v5 = v5
                                                                    }
                                                                    v5 = 1e10
                                                                }
                                                            }
                                                        } else {
                                                            var v87 float32 = 0.1
                                                            for fi88 := 0; fi88 < 20; fi88++ {
                                                                if true {
                                                                    v2 -= 152
                                                                    v89 := []int{}
                                                                    v90 := GenS2{F0: 65535, F1: (v1 & (v1 + v1)), F2: 4294967295}
                                                                    v91 := GenS6{F0: 544, F1: 1, F2: (v1 >> ((65535) & 63))}
                                                                    if len(v89) >= 2 {
                                                                        v89 = v89[1:len(v89)]
                                                                    }
                                                                    _ = v89
                                                                    _ = v90
                                                                    _ = v91
                                                                } else {
                                                                    v92 := GenS2{F0: 2, F1: v1, F2: 4294967295}
                                                                    v1 += 0
                                                                    _ = v92
                                                                }
                                                                var v93 float32 = (1.5 - ((v87 / (v87 / (3.14159 * 0.1))) / 0.1))
                                                                var v94 int = (fi66 / (((fi88 % (((fi66 * 407)) | 1))) | 1))
                                                                _ = fi88
                                                                _ = v93
                                                                _ = v94
                                                            }
                                                            v87 -= (v87 - v87)
                                                            _ = v87
                                                        }
                                                        var v95 float32 = float32((1.5 / (((v5 * -2.25) / -2.25) + 3.14159)))
                                                        _ = v95
                                                    }
                                                } else {
                                                    fmt.Println(v3, v5, fi66)
                                                }
                                            }
                                            var v96 uint64 = 4294967295
                                            _ = v96
                                        }
                                    case -96:
                                        var v97 int8 = v1
                                        _ = v97
                                    default:
                                        v5 = v5
                                    }
                                    var v98 float64 = (v5 * (((v5 * (float64(0.0) * 0.0)) + (v5 / 0.1)) + -2.25))
                                    fmt.Println(v1, v1)
                                    _ = v98
                                }
                            case "":
                                var v99 uint64 = 2
                                _ = v99
                            case "a":
                                if (false && false) {
                                    for fi100 := 0; fi100 < 8; fi100++ {
                                        var v101 uint64 = 18446744073709551615
                                        v2 -= (^(v2))
                                        var v102 uint64 = 4294967295
                                        _ = fi100
                                        _ = v101
                                        _ = v102
                                    }
                                    var v103 int16 = -1
                                    _ = v103
                                } else {
                                    v1 = (v1 ^ v1)
                                }
                            default:
                                if true {
                                    v104 := GenS7{F0: (v3 / ((v3) | 1))}
                                    var v105 uint16 = 255
                                    for fi106 := 0; fi106 < 12; fi106++ {
                                        v107 := GenS8{F0: 1, F1: (v2 + v2)}
                                        switch "a" {
                                        case "a":
                                            var v108 int64 = 1
                                            _ = v108
                                        case "":
                                            v2 -= 72
                                        case "foo bar":
                                            v3 += int32(fi66)
                                        default:
                                            var v109 uint = 839
                                            _ = v109
                                        }
                                        v110 := GenS3{F0: v1, F1: "hello", F2: 0}
                                        _ = fi106
                                        _ = v107
                                        _ = v110
                                    }
                                    _ = v104
                                    _ = v105
                                } else {
                                    switch (-(fi66)) {
                                    case 128:
                                        v3 += (v3 & 564)
                                    case 0:
                                        fmt.Println(fi66)
                                    case 2147483647:
                                        switch fi66 {
                                        case 32767:
                                            for fi111 := 0; fi111 < 2; fi111++ {
                                                v1 += (v1 + (v1 << ((520) & 63)))
                                                _ = len(v4)
                                                _ = fi111
                                            }
                                        case 1:
                                            v5 = v5
                                        }
                                    default:
                                        var v112 int32 = 0
                                        _ = v112
                                    }
                                }
                            }
                        } else {
                            for fi113 := 0; fi113 < 5; fi113++ {
                                v114 := GenS4{F0: 42, F1: (^(v2)), F2: v3}
                                switch (fi113 * (fi66 ^ fi113)) {
                                case -1:
                                    v115 := GenS9{F0: int8(v2), F1: (v1 << ((0) & 63)), F2: 18446744073709551615}
                                    _ = v115
                                case 1:
                                    switch (^(fi66)) {
                                    case -9223372036854775808:
                                        var v116 int16 = -233
                                        v1 = 0
                                        switch fi113 {
                                        case -2147483648:
                                            var v117 int32 = 0
                                            fmt.Println(v5, fi113)
                                            v118 := GenS10{F0: (v2 + (v2 ^ v2)), F1: -2147483648}
                                            _ = v117
                                            _ = v118
                                        case 565:
                                            var v119 int = (fi113 / (((fi66 + (fi66 + (^(fi113))))) | 1))
                                            _ = v119
                                        }
                                        _ = v116
                                    case 32767:
                                        if false {
                                            v120 := GenS11{F0: (^(v2))}
                                            _ = v120
                                        } else {
                                            var v121 uint8 = (^(v2))
                                            switch (fi66 | fi66) {
                                            case 32768:
                                                fmt.Println(v5, v121, v121)
                                                v2 -= 255
                                                if (((-2.25 - (1e10 * v5)) / -2.25) > float64((float32(v5) + float32(v5)))) {
                                                    if false {
                                                        switch "hello" {
                                                        case "a":
                                                            var v122 uint32 = 2
                                                            _ = v122
                                                        case "\n\t":
                                                            v123 := v4[-9223372036854775808]
                                                            _ = v123
                                                        case "foo bar":
                                                            v4[-765] = fi113
                                                        default:
                                                            v124 := GenS12{F0: "hello", F1: -1}
                                                            _ = v124
                                                        }
                                                        delete(v4, fi113)
                                                        switch fi66 {
                                                        case 1:
                                                            _ = len(v4)
                                                            var v125 int8 = -128
                                                            if ((v121 * v121) <= v121) {
                                                                if (!(uint16(v2) == 2)) {
                                                                    switch (fi113 | 128) {
                                                                    case 62:
                                                                        switch (^(fi66)) {
                                                                        case 127:
                                                                            v126 := GenS3{F0: v1, F1: "", F2: -9223372036854775808}
                                                                            _ = v126
                                                                        case -350:
                                                                            if false {
                                                                                var v127 uint32 = 339
                                                                                _ = v127
                                                                            }
                                                                        case -1:
                                                                            fmt.Println(fi66)
                                                                            delete(v4, (-(fi113)))
                                                                        default:
                                                                            v2 = (v121 << ((4294967295) & 63))
                                                                        }
                                                                    case -1:
                                                                        var v128 uint16 = uint16(v121)
                                                                        _ = v128
                                                                    default:
                                                                        if ((("hello" + "hello") + "foo bar") == "a") {
                                                                            v129 := GenS5{F0: 585}
                                                                            _ = v129
                                                                        } else {
                                                                            switch (("foo bar" + "\n\t") + "a") {
                                                                            case "foo bar":
                                                                                v4[582] = (fi113 ^ 32767)
                                                                            case "\n\t":
                                                                                var v130 bool = (false && false)
                                                                                v131 := GenS12{F0: "hello", F1: 42}
                                                                                _ = v130
                                                                                _ = v131
                                                                            case "hello":
                                                                                var v132 float32 = ((float32(((1.5 - v5) / (v5 + (((1.5 - 0.0) / v5) / float64(-2.25))))) / ((-2.25 - 0.1) - 1.5)) - (-2.25 / 3.14159))
                                                                                _ = v132
                                                                            }
                                                                            v133 := GenS13{F0: v3, F1: v125, F2: -2147483648}
                                                                            _ = v133
                                                                        }
                                                                    }
                                                                }
                                                                var v134 uint64 = 18446744073709551615
                                                                var v135 float64 = 1.5
                                                                if false {
                                                                    switch fi113 {
                                                                    case 468:
                                                                        fmt.Println(v135, v121)
                                                                    case -50:
                                                                        switch "" {
                                                                        case "a":
                                                                            var v136 int32 = (^(v3))
                                                                            _ = v136
                                                                        }
                                                                    case 42:
                                                                        v137 := GenS14{F0: 488, F1: (v134 << ((255) & 63))}
                                                                        _ = v137
                                                                    }
                                                                    var v138 uint = 558
                                                                    _ = v138
                                                                }
                                                                _ = v134
                                                                _ = v135
                                                            }
                                                            var v139 uint = 1
                                                            _ = v125
                                                            _ = v139
                                                        case 0:
                                                            v140 := GenS15{F0: 2}
                                                            _ = v140
                                                        case 32767:
                                                            var v141 uint32 = 2147483647
                                                            _ = v141
                                                        default:
                                                            var v142 int16 = 32767
                                                            _ = v142
                                                        }
                                                    } else {
                                                        v1 = 0
                                                        if true {
                                                            v4[(fi113 | (fi113 & 127))] = fi66
                                                        }
                                                        v1 = (v1 & v1)
                                                    }
                                                    var v143 int8 = v1
                                                    v1 -= v1
                                                    _ = v143
                                                } else {
                                                    v144 := []int{(fi113 >> ((614) & 63)), (fi66 - 32767), int(v121), fi113}
                                                    _ = v144
                                                }
                                            default:
                                                var v145 float32 = ((1.5 - 3.14159) * float32(float64((1.5 - (0.1 - 0.0)))))
                                                _ = v145
                                            }
                                            _ = v121
                                        }
                                    default:
                                        var v146 uint = 464
                                        _ = v146
                                    }
                                }
                                _ = fi113
                                _ = v114
                            }
                            v2 -= (v2 + (^(v2)))
                        }
                        v1 += (v1 ^ (-(v1)))
                        v5 -= float64(float32(1.5))
                        _ = fi66
                    }
                } else {
                    if (!true) {
                        v147 := GenS2{F0: 354, F1: (v1 % ((v1) | 1)), F2: 0}
                        _ = v147
                    }
                }
            default:
                delete(v4, 4294967295)
            }
            var v148 float32 = 1.5
            _ = v148
        } else {
            v149 := GenS13{F0: v3, F1: 26, F2: v3}
            for fi150 := 0; fi150 < 4; fi150++ {
                var v151 float32 = 3.14159
                switch "hello" {
                case "":
                    if (0 == 1) {
                        v3 += -1
                        v1 -= v1
                    } else {
                        if true {
                            if true {
                                v152 := []int{}
                                _ = v152
                            }
                        } else {
                            if false {
                                if true {
                                    v5 -= 0.1
                                }
                                var v153 uint = 18446744073709551615
                                _ = v153
                            } else {
                                _ = len(v4)
                            }
                            if true {
                                v154 := GenS16{F0: "hello"}
                                var v155 uint8 = 255
                                var v156 string = ""
                                _ = v154
                                _ = v155
                                _ = v156
                            }
                            fmt.Println(v2, v1, v151)
                        }
                        for fi157 := 0; fi157 < 14; fi157++ {
                            fmt.Println(v1, fi157)
                            var v158 uint32 = 743
                            _ = fi157
                            _ = v158
                        }
                        var v159 int64 = -9223372036854775808
                        _ = v159
                    }
                default:
                    var v160 float32 = ((v151 - v151) * v151)
                    _ = v160
                }
                for fi161 := 0; fi161 < 15; fi161++ {
                    switch "a" {
                    case "a":
                        if false {
                            if false {
                                fmt.Println(v5, v5, v1)
                            }
                            if true {
                                _ = len(v4)
                            }
                            v162 := []int{fi161, -1, (-(fi150)), fi161}
                            _ = v162
                        } else {
                            v3 -= 0
                            v163 := GenS17{F0: fi161}
                            _ = v163
                        }
                    }
                    _ = fi161
                }
                _ = fi150
                _ = v151
            }
            _ = v149
        }
    case "foo bar":
        v164 := []int{4294967295, 890, -2147483648, 0}
        _ = v164
    default:
        if true {
            v165 := []int{0, 2147483647, 32768}
            _ = v165
        } else {
            v1 = (v1 | (v1 / (((v1 ^ 0)) | 1)))
            v4[42] = -9223372036854775808
            v166 := []int{32767}
            _ = v166
        }
    }
    v167 := GenS18{F0: 0}
    fmt.Println(v3)
    _ = len(v4)
    v3 -= 42
    v168 := GenS3{F0: 127, F1: ("foo bar" + "a"), F2: -631}
    var v169 uint8 = 255
    var v170 int = 32768
    var v171 uint64 = 0
    var v172 uint32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v169:", v169)
    fmt.Println("v170:", v170)
    fmt.Println("v171:", v171)
    fmt.Println("v172:", v172)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v167
    _ = v168
    _ = v169
    _ = v170
    _ = v171
    _ = v172
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
