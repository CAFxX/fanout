package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 int64 }
type GenS2 struct { F0 int32; F1 uint32; F2 uint }
type GenS3 struct { F0 int }
type GenS4 struct { F0 string; F1 uint32 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int16; F1 int32; F2 bool }
type GenS7 struct { F0 int64; F1 int }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int64; F1 uint16; F2 int }
type GenS10 struct { F0 int; F1 uint64; F2 uint16 }
type GenS11 struct { F0 uint; F1 bool }
type GenS12 struct { F0 uint32; F1 uint64 }
type GenS13 struct { F0 int8; F1 int64; F2 bool }
type GenS14 struct { F0 string }
type GenS15 struct { F0 int }
func genFn0(p0 int32, p1 uint8) int16 {
    var v6 float32 = 1.5
    if (2 != 0) {
        var v7 int16 = 1000
        fmt.Println(p1, v6, v7)
        _ = v7
    }
    var v8 uint32 = 450
    _ = p0
    _ = p1
    _ = v6
    _ = v8
    return -32768
}
func genFn1(p0 int16, p1 uint64) uint {
    v9 := GenS1{F0: (p1 | 18446744073709551615), F1: 2, F2: -318}
    p0 += p0
    fmt.Println(p1)
    v10 := GenS1{F0: 18446744073709551615, F1: 1, F2: -1}
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 255
}

func main() {
    v1 := []int{}
    v2 := []string{"hello", (("\n\t" + "\n\t") + (("a" + "") + "")), ("hello" + ("a" + "\n\t"))}
    v3 := make(map[int]int)
    v4 := make(map[int]string)
    v5 := make(map[int]uint32)
    var v11 int8 = 127
    v11 -= 42
    if (!false) {
        v12 := GenS1{F0: 579, F1: 128, F2: 0}
        _ = len(v5)
        _ = v12
    }
    if (-9223372036854775808 > 9223372036854775807) {
        if true {
            v13 := GenS1{F0: 1, F1: 2, F2: 9223372036854775807}
            _ = v13
        } else {
            if (42 <= 0) {
                v14 := GenS1{F0: 815, F1: 128, F2: 1}
                var v15 uint64 = 18446744073709551615
                _ = v14
                _ = v15
            }
            switch 4294967295 {
            case 779:
                v16 := GenS1{F0: 4294967295, F1: 26, F2: -302}
                _ = v16
            case 229:
                var v17 uint64 = 4294967295
                _ = v17
            }
            v11 -= 1
        }
        var v18 int = 4294967295
        _ = v18
    }
    v11 -= (v11 ^ v11)
    v11 = (v11 | 1)
    v11 += v11
    var v19 uint = 465
    v20 := GenS2{F0: 27, F1: 2147483648, F2: (v19 >> (((v19 | v19)) & 63))}
    v21 := GenS1{F0: 18446744073709551615, F1: 1, F2: -9223372036854775808}
    if (v19 < 18446744073709551615) {
        v22 := GenS1{F0: 2, F1: 137, F2: -1}
        switch ("" + "") {
        case "foo bar":
            for fi23 := 0; fi23 < 16; fi23++ {
                v24 := GenS1{F0: 2, F1: 127, F2: 179}
                _ = fi23
                _ = v24
            }
        case "hello":
            v19 += (v19 ^ genFn1(852, 777))
        case "\n\t":
            var v25 uint = v19
            _ = v25
        default:
            switch "hello" {
            case "\n\t":
                if (!false) {
                    for fi26 := 0; fi26 < 16; fi26++ {
                        v11 -= (v11 * v11)
                        if (-1 <= genFn0(-1, 2)) {
                            if (!true) {
                                var v27 uint64 = 0
                                if false {
                                    _ = len(v5)
                                    for fi28 := 0; fi28 < 5; fi28++ {
                                        var v29 int32 = int32(v19)
                                        v1 = append(v1, (fi28 | fi28))
                                        _ = fi28
                                        _ = v29
                                    }
                                }
                                _ = v27
                            } else {
                                if true {
                                    if true {
                                        v1 = append(v1, (fi26 - fi26))
                                    } else {
                                        if false {
                                            v30 := GenS1{F0: 2, F1: 1, F2: 0}
                                            fmt.Println(v19)
                                            var v31 int8 = v11
                                            _ = v30
                                            _ = v31
                                        } else {
                                            fmt.Println(fi26)
                                            v11 = (v11 << ((v19) & 63))
                                            for fi32 := 0; fi32 < 4; fi32++ {
                                                fmt.Println(v11, v19)
                                                _ = fi32
                                            }
                                        }
                                        for fi33 := 0; fi33 < 17; fi33++ {
                                            fmt.Println(v11, fi26, fi33)
                                            _ = fi33
                                        }
                                    }
                                    var v34 int16 = 32767
                                    _ = v34
                                } else {
                                    v35 := GenS3{F0: 42}
                                    for fi36 := 0; fi36 < 5; fi36++ {
                                        var v37 int16 = 42
                                        v37 = genFn0(1, 210)
                                        fmt.Println(v11, fi36)
                                        _ = fi36
                                        _ = v37
                                    }
                                    if len(v1) > 0 {
                                        v1[0] = (fi26 << (((v19 & 2)) & 63))
                                    }
                                    _ = v35
                                }
                            }
                        }
                        var v38 int16 = 1
                        _ = fi26
                        _ = v38
                    }
                } else {
                    v19 = (^(v19))
                    v19 = (v19 - v19)
                }
            case "hello":
                switch -458 {
                case 98:
                    v1 = append(v1, 42)
                }
            default:
                v4[128] = "foo bar"
            }
        }
        _ = v22
    } else {
        for fi39 := 0; fi39 < 13; fi39++ {
            v40 := GenS1{F0: 18446744073709551615, F1: 0, F2: -1}
            v19 += v19
            _ = fi39
            _ = v40
        }
    }
    if false {
        if (-493 >= -32768) {
            var v41 float32 = -2.25
            var v42 int32 = 0
            fmt.Println(v19, v19)
            _ = v41
            _ = v42
        }
        var v43 float32 = 1e10
        _ = v43
    } else {
        for fi44 := 0; fi44 < 8; fi44++ {
            for fi45 := 0; fi45 < 17; fi45++ {
                fmt.Println(v11, fi44, fi44)
                var v46 int32 = 433
                _ = fi45
                _ = v46
            }
            switch "foo bar" {
            case "a":
                if true {
                    for fi47 := 0; fi47 < 13; fi47++ {
                        v11 = (v11 + v11)
                        _ = fi47
                    }
                    if (294 != 0) {
                        switch "a" {
                        case "hello":
                            if true {
                                for fi48 := 0; fi48 < 18; fi48++ {
                                    if false {
                                        var v49 float64 = 0.0
                                        v19 -= (v19 & v19)
                                        var v50 uint16 = 32767
                                        _ = v49
                                        _ = v50
                                    }
                                    _ = fi48
                                }
                            }
                        case "foo bar":
                            for fi51 := 0; fi51 < 1; fi51++ {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                fmt.Println(v11, fi51, v19)
                                var v52 float32 = (3.14159 / -2.25)
                                _ = fi51
                                _ = v52
                            }
                        case "":
                            var v53 int8 = (v11 * 114)
                            _ = v53
                        default:
                            for fi54 := 0; fi54 < 17; fi54++ {
                                if true {
                                    v19 = v19
                                    var v55 uint16 = 255
                                    if (1.5 != 1.5) {
                                        var v56 int8 = (v11 << (((^(v19))) & 63))
                                        v4[fi54] = "a"
                                        switch 0 {
                                        case 42:
                                            var v57 uint32 = 542
                                            _ = v57
                                        default:
                                            switch (fi44 & (fi54 | 4294967295)) {
                                            case 0:
                                                switch "\n\t" {
                                                case "\n\t":
                                                    switch 32767 {
                                                    case 4294967295:
                                                        v55 += uint16(v19)
                                                    case 2147483647:
                                                        var v58 float64 = (-2.25 + 1.5)
                                                        _ = v58
                                                    case 587:
                                                        var v59 uint = genFn1(1000, 147)
                                                        _ = v59
                                                    default:
                                                        fmt.Println(v55, fi54)
                                                    }
                                                case "hello":
                                                    if (("foo bar" + ("hello" + "\n\t")) == "") {
                                                        var v60 int8 = 127
                                                        v61 := v5[fi54]
                                                        _ = v60
                                                        _ = v61
                                                    }
                                                default:
                                                    delete(v5, fi54)
                                                }
                                            case -9223372036854775808:
                                                var v62 uint32 = 4294967295
                                                _ = v62
                                            default:
                                                v55 = v55
                                            }
                                            var v63 uint64 = 233
                                            _ = v63
                                        }
                                        _ = v56
                                    }
                                    _ = v55
                                } else {
                                    v19 = 1
                                    var v64 uint16 = 2
                                    if (0 > uint64(v64)) {
                                        var v65 uint64 = 42
                                        switch "foo bar" {
                                        case "a":
                                            fmt.Println(v11)
                                        case "":
                                            if false {
                                                v4[95] = "\n\t"
                                            }
                                        default:
                                            fmt.Println(fi44, fi44)
                                        }
                                        switch fi54 {
                                        case 32768:
                                            if false {
                                                _ = len(v5)
                                                v66 := GenS2{F0: 1, F1: 1, F2: (v19 ^ (v19 - 951))}
                                                _ = v66
                                            } else {
                                                v11 = v11
                                            }
                                        case 2147483647:
                                            var v67 uint64 = (^(v65))
                                            _ = v67
                                        case -1:
                                            v19 = genFn1(0, v65)
                                        default:
                                            if true {
                                                if ((-32768 != genFn0(-2147483648, 128)) && ((v64 != 264) || (fi54 < fi54))) {
                                                    var v68 int64 = 881
                                                    if (v65 != (v65 / (((v65 << (((^(v19))) & 63))) | 1))) {
                                                        v69 := v5[(fi54 - 553)]
                                                        _ = v69
                                                    } else {
                                                        if false {
                                                            _ = len(v1)
                                                        } else {
                                                            var v70 float64 = ((1e10 + 0.0) - 1.5)
                                                            v71 := GenS3{F0: fi54}
                                                            v64 = v64
                                                            _ = v70
                                                            _ = v71
                                                        }
                                                        v72 := GenS3{F0: (fi54 / (((fi44 | -2147483648)) | 1))}
                                                        v73 := GenS1{F0: (v65 / ((4294967295) | 1)), F1: uint8(v65), F2: -1}
                                                        _ = v72
                                                        _ = v73
                                                    }
                                                    v3[-50] = fi54
                                                    _ = v68
                                                } else {
                                                    if (227 <= 127) {
                                                        var v74 float32 = (float32(3.14159) * (1.5 * (-2.25 / (1.5 / 1e10))))
                                                        _ = v74
                                                    } else {
                                                        if (("a" == "hello") || false) {
                                                            if len(v2) > 0 {
                                                                v2[0] = "\n\t"
                                                            }
                                                        } else {
                                                            fmt.Println(v65, v65, v11)
                                                        }
                                                        switch (((("" + "a") + ("\n\t" + "foo bar")) + ("a" + "")) + ((("\n\t" + ("\n\t" + "a")) + "a") + "a")) {
                                                        case "foo bar":
                                                            v75 := GenS4{F0: ("a" + (("hello" + "a") + "a")), F1: 65535}
                                                            _ = v75
                                                        case "a":
                                                            v19 = v19
                                                        default:
                                                            fmt.Println(v11, v11)
                                                        }
                                                        var v76 uint8 = 199
                                                        v77 := GenS2{F0: -808, F1: 2147483648, F2: (^(v19))}
                                                        _ = v76
                                                        _ = v77
                                                    }
                                                    var v78 uint16 = v64
                                                    switch "\n\t" {
                                                    case "":
                                                        var v79 int32 = int32(v65)
                                                        _ = v79
                                                    case "a":
                                                        if true {
                                                            if false {
                                                                switch "" {
                                                                case "hello":
                                                                    v19 -= uint(v65)
                                                                case "foo bar":
                                                                    var v80 int16 = 1
                                                                    _ = v80
                                                                default:
                                                                    v81 := GenS5{F0: v78}
                                                                    _ = v81
                                                                }
                                                                switch "" {
                                                                case "":
                                                                    var v82 int32 = 0
                                                                    _ = v82
                                                                case "\n\t":
                                                                    v83 := GenS6{F0: 1000, F1: int32(v11), F2: (1.5 == 0.1)}
                                                                    _ = v83
                                                                }
                                                                v84 := GenS5{F0: (v64 / ((2) | 1))}
                                                                _ = v84
                                                            } else {
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                                v85 := GenS3{F0: 128}
                                                                _ = v85
                                                            }
                                                            v19 += 65535
                                                            if false {
                                                                switch "a" {
                                                                case "a":
                                                                    v86 := GenS6{F0: 0, F1: int32(fi44), F2: false}
                                                                    _ = v86
                                                                }
                                                            } else {
                                                                if (!(3.14159 >= 0.1)) {
                                                                    fmt.Println(fi44, v64)
                                                                } else {
                                                                    switch "foo bar" {
                                                                    case "":
                                                                        v87 := GenS4{F0: ((("a" + "a") + ("hello" + "\n\t")) + (("" + "foo bar") + "")), F1: 31}
                                                                        _ = v87
                                                                    case "\n\t":
                                                                        var v88 uint = v19
                                                                        _ = v88
                                                                    default:
                                                                        fmt.Println(fi44, v64)
                                                                        if true {
                                                                            fmt.Println(fi54, v19)
                                                                            var v89 float64 = -2.25
                                                                            var v90 int64 = 448
                                                                            _ = v89
                                                                            _ = v90
                                                                        }
                                                                    }
                                                                }
                                                                delete(v5, (fi54 ^ fi44))
                                                                v91 := GenS3{F0: fi44}
                                                                _ = len(v2)
                                                                _ = v91
                                                            }
                                                        } else {
                                                            fmt.Println(v65, fi54, v65)
                                                            var v92 uint32 = 65535
                                                            _ = v92
                                                        }
                                                    case "\n\t":
                                                        switch (("a" + "foo bar") + "hello") {
                                                        case "a":
                                                            v78 = (^(v64))
                                                        default:
                                                            var v93 int16 = 32767
                                                            _ = v93
                                                        }
                                                    }
                                                    _ = v78
                                                }
                                                v19 = genFn1(42, (v65 % (((^(v65))) | 1)))
                                                v19 += (v19 / (((v19 & genFn1(1000, v65))) | 1))
                                            }
                                        }
                                        _ = v65
                                    } else {
                                        if true {
                                            fmt.Println(v11)
                                            switch fi54 {
                                            case 32767:
                                                v64 = v64
                                                var v94 int64 = 42
                                                _ = v94
                                            }
                                            var v95 uint = (v19 ^ v19)
                                            if false {
                                                fmt.Println(v95, v19, v11)
                                                switch ((("a" + ("\n\t" + ("" + "\n\t"))) + "") + "hello") {
                                                case "hello":
                                                    var v96 uint8 = 190
                                                    _ = v96
                                                case "a":
                                                    if (1e10 != 1.5) {
                                                        v19 -= v95
                                                        switch "hello" {
                                                        case "hello":
                                                            v64 = (v64 | v64)
                                                        case "foo bar":
                                                            fmt.Println(v11)
                                                        }
                                                    } else {
                                                        switch ("foo bar" + "foo bar") {
                                                        case "hello":
                                                            var v97 uint16 = v64
                                                            _ = v97
                                                        case "foo bar":
                                                            var v98 int16 = 241
                                                            _ = v98
                                                        case "":
                                                            var v99 int32 = 303
                                                            _ = v99
                                                        default:
                                                            switch "" {
                                                            case "\n\t":
                                                                switch (fi54 << ((uint(v64)) & 63)) {
                                                                case 32767:
                                                                    fmt.Println(fi44, v11)
                                                                case 529:
                                                                    if (1 != 42) {
                                                                        if (1 >= -77) {
                                                                            fmt.Println(fi54)
                                                                            v95 = (v95 << ((255) & 63))
                                                                            v100 := GenS7{F0: -1, F1: fi54}
                                                                            var v101 float64 = 0.1
                                                                            _ = v100
                                                                            _ = v101
                                                                        } else {
                                                                            if (fi44 >= fi44) {
                                                                                if true {
                                                                                    fmt.Println(fi44, v19)
                                                                                }
                                                                                var v102 int16 = -1
                                                                                _ = v102
                                                                            }
                                                                            fmt.Println(v95)
                                                                            if true {
                                                                                switch (("\n\t" + "") + "a") {
                                                                                case "hello":
                                                                                    if (220 < 2) {
                                                                                        v103 := GenS4{F0: "", F1: 42}
                                                                                        if false {
                                                                                            var v104 int = 185
                                                                                            if len(v2) >= 2 {
                                                                                                v2 = v2[1:len(v2)]
                                                                                            }
                                                                                            _ = v104
                                                                                        }
                                                                                        var v105 uint64 = 533
                                                                                        _ = v103
                                                                                        _ = v105
                                                                                    } else {
                                                                                        v95 = (v19 | 42)
                                                                                        var v106 bool = (!(!(1 != 42)))
                                                                                        if len(v1) > 0 {
                                                                                            v1[0] = (fi44 ^ fi54)
                                                                                        }
                                                                                        delete(v3, (fi44 ^ fi44))
                                                                                        v19 = (v95 >> ((499) & 63))
                                                                                        var v107 uint32 = 730
                                                                                        var v108 int32 = 455
                                                                                        _ = v106
                                                                                        _ = v107
                                                                                        _ = v108
                                                                                    }
                                                                                case "\n\t":
                                                                                    var v109 int64 = 1
                                                                                    _ = v109
                                                                                default:
                                                                                    fmt.Println(v11, v64, fi44)
                                                                                }
                                                                                fmt.Println(v19, fi54)
                                                                            }
                                                                        }
                                                                        v110 := v5[(fi44 << (((v95 - genFn1(genFn0(0, 42), 0))) & 63))]
                                                                        _ = v110
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        if len(v1) > 0 {
                                                            v1[0] = (fi54 + fi54)
                                                        }
                                                    }
                                                case "\n\t":
                                                    var v111 uint64 = 18446744073709551615
                                                    _ = v111
                                                }
                                            } else {
                                                if (false || (507 >= 9223372036854775807)) {
                                                    var v112 float32 = 1.5
                                                    if ((v64 * 752) <= (v64 - v64)) {
                                                        switch "a" {
                                                        case "foo bar":
                                                            v113 := GenS8{F0: true}
                                                            _ = v113
                                                        case "\n\t":
                                                            var v114 uint = genFn1(0, 0)
                                                            if ((0.1 > v112) && false) {
                                                                var v115 int = fi54
                                                                v116 := GenS9{F0: 0, F1: v64, F2: (fi54 + v115)}
                                                                var v117 float32 = ((v112 - (1.5 - v112)) / ((v112 - 1e10) + v112))
                                                                _ = v115
                                                                _ = v116
                                                                _ = v117
                                                            } else {
                                                                var v118 bool = (true || false)
                                                                _ = v118
                                                            }
                                                            _ = v114
                                                        }
                                                    } else {
                                                        v95 = (v19 | v95)
                                                    }
                                                    if len(v1) > 0 {
                                                        v1[0] = (fi44 >> (((^(v95))) & 63))
                                                    }
                                                    _ = v112
                                                } else {
                                                    v119 := GenS10{F0: (fi54 - int(v11)), F1: 4294967295, F2: v64}
                                                    _ = v119
                                                }
                                                if (1000 != -949) {
                                                    fmt.Println(fi54)
                                                    var v120 int64 = -1
                                                    v19 = v19
                                                    if false {
                                                        switch "hello" {
                                                        case "hello":
                                                            var v121 float64 = 0.1
                                                            _ = v121
                                                        case "":
                                                            var v122 uint = (^(v95))
                                                            _ = v122
                                                        case "a":
                                                            v1 = append(v1, fi44)
                                                        }
                                                        var v123 uint32 = 339
                                                        v124 := GenS7{F0: (v120 + (v120 / (((v120 & -563)) | 1))), F1: (-(fi54))}
                                                        _ = v123
                                                        _ = v124
                                                    } else {
                                                        v125 := GenS4{F0: "a", F1: 42}
                                                        _ = v125
                                                    }
                                                    _ = v120
                                                }
                                                fmt.Println(v95)
                                            }
                                            if true {
                                                fmt.Println(fi44, fi54, v19)
                                            } else {
                                                var v126 uint32 = 65535
                                                fmt.Println(fi54)
                                                _ = v126
                                            }
                                            var v127 uint = v19
                                            _ = v127
                                            _ = v95
                                        } else {
                                            delete(v4, -2147483648)
                                        }
                                    }
                                    _ = v64
                                }
                                if false {
                                    fmt.Println(fi44, fi54, fi54)
                                } else {
                                    v128 := GenS11{F0: (v19 % ((18446744073709551615) | 1)), F1: true}
                                    v129 := v5[fi44]
                                    _ = v128
                                    _ = v129
                                }
                                _ = fi54
                            }
                        }
                        for fi130 := 0; fi130 < 18; fi130++ {
                            var v131 int64 = -9223372036854775808
                            fmt.Println(fi130)
                            _ = fi130
                            _ = v131
                        }
                        switch ("foo bar" + "a") {
                        case "foo bar":
                            v11 = (v11 ^ -1)
                        case "a":
                            switch "\n\t" {
                            case "foo bar":
                                for fi132 := 0; fi132 < 10; fi132++ {
                                    _ = len(v3)
                                    _ = fi132
                                }
                            case "a":
                                if false {
                                    var v133 int8 = v11
                                    _ = v133
                                }
                            case "hello":
                                fmt.Println(fi44)
                            default:
                                if true {
                                    v134 := GenS1{F0: 1, F1: 42, F2: -9223372036854775808}
                                    _ = v134
                                } else {
                                    for fi135 := 0; fi135 < 16; fi135++ {
                                        switch ("hello" + ((("hello" + ("foo bar" + (("a" + "a") + "\n\t"))) + "a") + ("" + ("\n\t" + ("\n\t" + ("foo bar" + "a")))))) {
                                        case "a":
                                            var v136 int8 = -8
                                            _ = v136
                                        case "hello":
                                            fmt.Println(v19)
                                        case "\n\t":
                                            switch (-(fi44)) {
                                            case 42:
                                                v19 = v19
                                            case 4294967295:
                                                if (42 > -2147483648) {
                                                    v19 = (v19 >> ((v19) & 63))
                                                    v5[(fi135 * 9223372036854775807)] = 65535
                                                    if (!false) {
                                                        if false {
                                                            if false {
                                                                switch ("foo bar" + "") {
                                                                case "hello":
                                                                    _ = len(v4)
                                                                case "foo bar":
                                                                    v11 += int8(v19)
                                                                default:
                                                                    var v137 bool = true
                                                                    _ = v137
                                                                }
                                                                switch (fi135 % ((fi135) | 1)) {
                                                                case 921:
                                                                    fmt.Println(fi44, v11)
                                                                case 32767:
                                                                    if true {
                                                                        if (true && true) {
                                                                            v138 := GenS12{F0: 0, F1: 18446744073709551615}
                                                                            switch (("foo bar" + "foo bar") + "\n\t") {
                                                                            case "a":
                                                                                switch int(v11) {
                                                                                case -1:
                                                                                    v139 := GenS13{F0: -128, F1: 0, F2: (41 > 1)}
                                                                                    _ = v139
                                                                                case 558:
                                                                                    v5[fi135] = 733
                                                                                default:
                                                                                    v11 += (-(v11))
                                                                                }
                                                                            default:
                                                                                delete(v5, (fi44 ^ (^(fi44))))
                                                                            }
                                                                            _ = v138
                                                                        } else {
                                                                            var v140 bool = false
                                                                            _ = v140
                                                                        }
                                                                    } else {
                                                                        if (false || false) {
                                                                            v141 := GenS11{F0: 0, F1: true}
                                                                            _ = v141
                                                                        }
                                                                        fmt.Println(v19, fi135, v11)
                                                                        var v142 float32 = (1.5 * (0.1 - -2.25))
                                                                        _ = v142
                                                                    }
                                                                    v19 += (v19 * 0)
                                                                    var v143 uint32 = 4294967295
                                                                    _ = v143
                                                                default:
                                                                    var v144 int = fi135
                                                                    _ = v144
                                                                }
                                                            }
                                                            var v145 int32 = 157
                                                            v146 := GenS14{F0: ("\n\t" + "foo bar")}
                                                            _ = v145
                                                            _ = v146
                                                        } else {
                                                            v19 -= (v19 >> ((v19) & 63))
                                                            if false {
                                                                switch (fi44 << ((v19) & 63)) {
                                                                case 32767:
                                                                    var v147 int16 = 42
                                                                    _ = v147
                                                                }
                                                                v11 -= v11
                                                                var v148 float64 = -2.25
                                                                if false {
                                                                    var v149 int = (fi135 / (((fi135 * (fi135 - -2147483648))) | 1))
                                                                    if true {
                                                                        v148 -= (3.14159 - v148)
                                                                        if ((fi135 - fi44) != (-(v149))) {
                                                                            v150 := GenS15{F0: fi135}
                                                                            if true {
                                                                                var v151 uint8 = 85
                                                                                _ = v151
                                                                            }
                                                                            _ = v150
                                                                        } else {
                                                                            v19 = v19
                                                                            switch "a" {
                                                                            case "foo bar":
                                                                                v149 = 668
                                                                            case "":
                                                                                if (false || false) {
                                                                                    v11 = v11
                                                                                }
                                                                            case "a":
                                                                                var v152 int32 = -1
                                                                                if (627 > int16(v19)) {
                                                                                    if len(v1) > 0 {
                                                                                        v1[0] = (v149 << (((v19 - genFn1(-256, 2))) & 63))
                                                                                    }
                                                                                    v2 = append(v2, "foo bar")
                                                                                } else {
                                                                                    v148 = (0.0 / v148)
                                                                                    v19 = (^(v19))
                                                                                    v149 -= (-(v149))
                                                                                }
                                                                                switch -1 {
                                                                                case 4294967295:
                                                                                    switch "foo bar" {
                                                                                    case "foo bar":
                                                                                        v153 := GenS5{F0: 2}
                                                                                        _ = v153
                                                                                    case "hello":
                                                                                        v154 := GenS10{F0: fi135, F1: 4294967295, F2: 0}
                                                                                        _ = v154
                                                                                    case "\n\t":
                                                                                        _ = len(v4)
                                                                                    }
                                                                                case 32768:
                                                                                    v148 = v148
                                                                                case 108:
                                                                                    v11 -= -51
                                                                                }
                                                                                fmt.Println(v11)
                                                                                var v155 uint64 = 509
                                                                                _ = v152
                                                                                _ = v155
                                                                            }
                                                                        }
                                                                        v156 := GenS13{F0: 42, F1: 0, F2: ((225 != -2147483648) && ((true || (2 != 2)) && true))}
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                        _ = v156
                                                                    } else {
                                                                        switch v149 {
                                                                        case 32767:
                                                                            var v157 int16 = -842
                                                                            _ = v157
                                                                        }
                                                                        switch "foo bar" {
                                                                        case "\n\t":
                                                                            switch "hello" {
                                                                            case "foo bar":
                                                                                v1 = append(v1, v149)
                                                                            case "\n\t":
                                                                                var v158 uint64 = 42
                                                                                _ = v158
                                                                            default:
                                                                                v149 += (^(v149))
                                                                            }
                                                                        }
                                                                        v159 := v4[(fi44 + (fi135 + fi44))]
                                                                        _ = v159
                                                                    }
                                                                    var v160 int16 = 1
                                                                    var v161 uint = 453
                                                                    v149 = fi135
                                                                    _ = v149
                                                                    _ = v160
                                                                    _ = v161
                                                                }
                                                                var v162 uint64 = 349
                                                                _ = v148
                                                                _ = v162
                                                            }
                                                            switch "" {
                                                            case "hello":
                                                                v163 := GenS5{F0: 1}
                                                                _ = v163
                                                            case "foo bar":
                                                                var v164 int64 = 1
                                                                _ = v164
                                                            default:
                                                                var v165 uint32 = 0
                                                                _ = v165
                                                            }
                                                        }
                                                    } else {
                                                        if false {
                                                            var v166 uint32 = 122
                                                            v3[fi135] = fi135
                                                            if (-2.25 == 1e10) {
                                                                v167 := v5[(fi44 | (fi44 >> ((v19) & 63)))]
                                                                _ = v167
                                                            }
                                                            _ = v166
                                                        } else {
                                                            var v168 int = fi44
                                                            _ = v168
                                                        }
                                                    }
                                                } else {
                                                    var v169 uint64 = 42
                                                    var v170 float64 = 0.0
                                                    _ = v169
                                                    _ = v170
                                                }
                                            case -1:
                                                v11 = (-(v11))
                                            default:
                                                var v171 uint16 = 65535
                                                _ = v171
                                            }
                                            v11 = (^(v11))
                                        }
                                        _ = fi135
                                    }
                                }
                            }
                        default:
                            v172 := GenS1{F0: 1, F1: 32, F2: 1}
                            _ = v172
                        }
                    }
                    for fi173 := 0; fi173 < 12; fi173++ {
                        fmt.Println(v11, v19)
                        _ = fi173
                    }
                } else {
                    if true {
                        v19 = v19
                    } else {
                        fmt.Println(v19, fi44)
                    }
                    switch "\n\t" {
                    case "a":
                        switch fi44 {
                        case 976:
                            var v174 string = "a"
                            _ = v174
                        case 0:
                            for fi175 := 0; fi175 < 2; fi175++ {
                                var v176 int16 = 32767
                                fmt.Println(v19)
                                _ = len(v2)
                                _ = fi175
                                _ = v176
                            }
                        case 263:
                            v19 = (v19 ^ (v19 >> (((^(v19))) & 63)))
                        default:
                            var v177 uint64 = 629
                            _ = v177
                        }
                    default:
                        _ = len(v3)
                    }
                }
            case "":
                fmt.Println(fi44, v19, v19)
            default:
                v11 = v11
            }
            _ = fi44
        }
    }
    fmt.Println(v19)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
