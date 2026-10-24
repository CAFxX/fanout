package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 bool; F1 uint64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint16; F1 uint16 }
type GenS6 struct { F0 int64; F1 uint; F2 uint16 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint64; F1 int16; F2 int8 }
type GenS9 struct { F0 uint8; F1 string }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 int32; F1 int; F2 int }
type GenS12 struct { F0 int16; F1 int16 }
type GenS13 struct { F0 uint32; F1 int8 }
type GenS14 struct { F0 int8 }
type GenS15 struct { F0 int8 }
type GenS16 struct { F0 uint8; F1 uint64; F2 int }
type GenS17 struct { F0 int64; F1 bool }
func genFn0(p0 uint8, p1 uint16) uint64 {
    p0 -= (^(p0))
    v5 := []int{int(p0), 889, 0, 1}
    v6 := GenS1{F0: 1}
    v5 = append(v5, 903)
    switch "foo bar" {
    case "":
        v7 := GenS1{F0: 4294967295}
        _ = v7
    default:
        p1 -= 2
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 18446744073709551615
}
func genFn1(p0 uint32, p1 string, p2 uint32) uint {
    p1 = ""
    v8 := GenS2{F0: 1, F1: 32767}
    if (750 >= 2147483647) {
        switch 4294967295 {
        case 32768:
            var v9 bool = false
            _ = v9
        case -1:
            v10 := []int{-2147483648, 1}
            _ = v10
        default:
            if (255 == 0) {
                v11 := []int{-1, 0, 935, 128}
                _ = v11
            } else {
                p2 += (p0 | p0)
            }
        }
        switch 4294967295 {
        case 127:
            for fi12 := 0; fi12 < 15; fi12++ {
                fmt.Println(p2)
                p0 += p2
                var v13 string = (p1 + p1)
                _ = fi12
                _ = v13
            }
        default:
            var v14 uint64 = 197
            _ = v14
        }
        if false {
            p0 += 537
            fmt.Println(p2)
        } else {
            p0 += 594
            v15 := GenS1{F0: 18446744073709551615}
            fmt.Println(p2)
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 5; fi16++ {
        for fi17 := 0; fi17 < 20; fi17++ {
            switch p1 {
            case "\n\t":
                var v18 int32 = 1
                _ = v18
            case "hello":
                var v19 uint16 = 32767
                _ = v19
            default:
                var v20 float64 = 1e10
                _ = v20
            }
            v21 := []int{}
            _ = len(v21)
            _ = fi17
            _ = v21
        }
        _ = fi16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 18446744073709551615
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint8 = 1
    v3 := []uint32{}
    var v4 int8 = 109
    var v22 float64 = -2.25
    var v23 uint64 = 1
    v2 -= (^(v2))
    v22 = -2.25
    delete(v1, 1)
    _ = len(v3)
    if (1 >= 42) {
        fmt.Println(v22, v23, v23)
        for fi24 := 0; fi24 < 10; fi24++ {
            fmt.Println(v4, v22)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if true {
                v25 := GenS3{F0: (false || false), F1: 4294967295}
                for fi26 := 0; fi26 < 3; fi26++ {
                    switch -1 {
                    case -1:
                        if (false && (("hello" + ("hello" + "a")) == ("\n\t" + ""))) {
                            switch (fi26 + fi26) {
                            case -9223372036854775808:
                                if (!((4294967295 == 2147483648) && true)) {
                                    if false {
                                        v22 += 1e10
                                        v23 += (^(v23))
                                        var v27 float64 = v22
                                        _ = v27
                                    } else {
                                        if false {
                                            v28 := GenS4{F0: -132}
                                            _ = v28
                                        } else {
                                            v29 := GenS2{F0: 18446744073709551615, F1: 2}
                                            var v30 uint64 = (^(v23))
                                            v31 := GenS5{F0: 42, F1: 990}
                                            v32 := GenS4{F0: -420}
                                            _ = v29
                                            _ = v30
                                            _ = v31
                                            _ = v32
                                        }
                                    }
                                    v33 := GenS4{F0: 9223372036854775807}
                                    v34 := GenS1{F0: 1}
                                    _ = v33
                                    _ = v34
                                } else {
                                    _ = len(v1)
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    delete(v1, 477)
                                }
                            case -608:
                                var v35 int = (-(fi24))
                                fmt.Println(v35, fi24, fi26)
                                var v36 float64 = v22
                                _ = v35
                                _ = v36
                            case 127:
                                var v37 int32 = -16
                                _ = v37
                            }
                        } else {
                            _ = len(v1)
                            fmt.Println(v4)
                            v38 := GenS1{F0: v23}
                            v3 = append(v3, 2147483648)
                            _ = v38
                        }
                    case 632:
                        _ = len(v1)
                    case -9223372036854775808:
                        if true {
                            if true {
                                var v39 uint = 18446744073709551615
                                _ = v39
                            } else {
                                if (!true) {
                                    if (!(266 >= 1000)) {
                                        var v40 uint = 42
                                        _ = v40
                                    } else {
                                        switch -453 {
                                        case 1:
                                            if true {
                                                switch fi24 {
                                                case -1:
                                                    delete(v1, 523)
                                                case -2147483648:
                                                    v2 = 107
                                                case 32768:
                                                    v41 := GenS3{F0: ((2 < 42) && false), F1: uint64(v2)}
                                                    _ = v41
                                                }
                                            } else {
                                                var v42 int8 = -128
                                                _ = v42
                                            }
                                        default:
                                            v43 := GenS5{F0: 1, F1: 32767}
                                            _ = v43
                                        }
                                        if ((true && (2 == 0)) && true) {
                                            var v44 uint64 = (v23 ^ genFn0((v2 + 163), 42))
                                            _ = v44
                                        }
                                    }
                                    var v45 int8 = v4
                                    switch "" {
                                    case "\n\t":
                                        var v46 int64 = -9223372036854775808
                                        _ = v46
                                    default:
                                        v23 = genFn0(v2, 0)
                                    }
                                    _ = v45
                                } else {
                                    fmt.Println(v22, v23)
                                }
                            }
                            v23 = (^(v23))
                        } else {
                            var v47 string = "a"
                            switch v47 {
                            case "\n\t":
                                var v48 string = v47
                                _ = v48
                            case "":
                                switch v47 {
                                case "\n\t":
                                    if false {
                                        fmt.Println(fi26)
                                        fmt.Println(fi24, v2, v4)
                                        v49 := GenS3{F0: ((1.5 + -2.25) != 0.1), F1: (v23 / ((v23) | 1))}
                                        _ = v49
                                    }
                                    fmt.Println(fi24, v4, v22)
                                default:
                                    var v50 int = (fi26 & fi24)
                                    _ = v50
                                }
                            default:
                                _ = len(v3)
                            }
                            _ = v47
                        }
                    default:
                        v51 := GenS2{F0: 255, F1: 16}
                        _ = v51
                    }
                    var v52 bool = true
                    _ = fi26
                    _ = v52
                }
                if len(v3) > 0 {
                    v3[0] = 659
                }
                _ = v25
            }
            _ = fi24
        }
    } else {
        v53 := GenS3{F0: (-339 > -400), F1: 2}
        var v54 bool = true
        v55 := GenS6{F0: -493, F1: 0, F2: 0}
        _ = v53
        _ = v54
        _ = v55
    }
    switch 32767 {
    case -163:
        fmt.Println(v23, v22)
    case -2147483648:
        v1[42] = 422
    default:
        var v56 int64 = -1
        _ = v56
    }
    switch "" {
    case "foo bar":
        if (!true) {
            if false {
                for fi57 := 0; fi57 < 2; fi57++ {
                    if (true && true) {
                        if (65535 != uint16(v2)) {
                            v58 := GenS7{F0: 453}
                            v23 = (v23 * v23)
                            if (false && (-2.25 > (-2.25 + 0.0))) {
                                v59 := GenS8{F0: v23, F1: 810, F2: v4}
                                v60 := GenS4{F0: -1}
                                switch "" {
                                case "":
                                    fmt.Println(v4, v23)
                                }
                                _ = v59
                                _ = v60
                            } else {
                                if (4294967295 >= 255) {
                                    var v61 float32 = 0.1
                                    switch ("hello" + "\n\t") {
                                    case "hello":
                                        var v62 string = "\n\t"
                                        _ = v62
                                    case "foo bar":
                                        var v63 int16 = -632
                                        _ = v63
                                    case "":
                                        delete(v1, uint32(v23))
                                    default:
                                        var v64 string = "\n\t"
                                        _ = v64
                                    }
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    _ = v61
                                } else {
                                    v23 -= 4294967295
                                }
                            }
                            _ = v58
                        } else {
                            for fi65 := 0; fi65 < 11; fi65++ {
                                if (32767 >= 42) {
                                    v66 := GenS9{F0: v2, F1: "hello"}
                                    v4 = (^(v4))
                                    var v67 uint32 = 1
                                    _ = len(v1)
                                    _ = v66
                                    _ = v67
                                }
                                fmt.Println(fi57, fi57, fi57)
                                if (!true) {
                                    v68 := GenS10{F0: 476}
                                    v2 = v2
                                    fmt.Println(fi57, fi65)
                                    v23 += (^(v23))
                                    if false {
                                        v69 := GenS6{F0: -1, F1: 547, F2: 255}
                                        _ = v69
                                    }
                                    var v70 int64 = -114
                                    _ = v68
                                    _ = v70
                                }
                                _ = fi65
                            }
                            switch ("foo bar" + "hello") {
                            case "a":
                                v2 -= v2
                            case "foo bar":
                                var v71 uint8 = (v2 - (v2 + uint8(v23)))
                                _ = v71
                            case "":
                                v22 = (v22 / v22)
                            }
                        }
                        v72 := GenS11{F0: -277, F1: (fi57 / (((^(fi57))) | 1)), F2: fi57}
                        switch "hello" {
                        case "hello":
                            if (v23 == (v23 >> ((1) & 63))) {
                                var v73 uint = genFn1(4294967295, ("foo bar" + "a"), uint32(v2))
                                _ = v73
                            } else {
                                v22 = (v22 - v22)
                                if (v4 >= v4) {
                                    _ = len(v1)
                                }
                                switch fi57 {
                                case 2147483647:
                                    v22 -= v22
                                case 9223372036854775807:
                                    for fi74 := 0; fi74 < 5; fi74++ {
                                        if ((fi74 % (((fi74 + (fi74 ^ (^(fi57))))) | 1)) == fi57) {
                                            v22 -= float64(1.5)
                                            if true {
                                                v3 = append(v3, uint32(v23))
                                                v23 -= genFn0(152, 65535)
                                                var v75 uint16 = 858
                                                _ = v75
                                            } else {
                                                v22 += float64((1.5 / (0.1 + 1.5)))
                                                var v76 bool = true
                                                v22 = ((v22 * v22) / (((v22 + v22) - 0.0) - 3.14159))
                                                if ((v2 ^ v2) >= 0) {
                                                    if (!(414 <= 18446744073709551615)) {
                                                        if v76 {
                                                            v4 -= (^(v4))
                                                            v3 = append(v3, 42)
                                                            v2 += (v2 % (((v2 >> ((18446744073709551615) & 63))) | 1))
                                                            if (0 != -1) {
                                                                var v77 int16 = 135
                                                                _ = len(v3)
                                                                var v78 int8 = (^(v4))
                                                                v77 = v77
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                                _ = v77
                                                                _ = v78
                                                            }
                                                            v79 := GenS12{F0: -757, F1: 140}
                                                            if (v22 != float64(3.14159)) {
                                                                var v80 uint = 522
                                                                switch "hello" {
                                                                case "":
                                                                    fmt.Println(v2, v23)
                                                                case "foo bar":
                                                                    fmt.Println(v2, v80, fi74)
                                                                case "\n\t":
                                                                    _ = len(v3)
                                                                }
                                                                var v81 int32 = -549
                                                                _ = v80
                                                                _ = v81
                                                            }
                                                            _ = v79
                                                        } else {
                                                            switch "\n\t" {
                                                            case "a":
                                                                v76 = v76
                                                            case "\n\t":
                                                                v2 += 1
                                                            }
                                                        }
                                                        if v76 {
                                                            switch "\n\t" {
                                                            case "hello":
                                                                if v76 {
                                                                    v76 = v76
                                                                } else {
                                                                    var v82 int64 = 265
                                                                    fmt.Println(fi74)
                                                                    switch fi74 {
                                                                    case 32767:
                                                                        v1[463] = (fi74 * (fi74 >> ((255) & 63)))
                                                                    case -5:
                                                                        v23 = (v23 << ((2) & 63))
                                                                    }
                                                                    v4 -= (v4 & 127)
                                                                    _ = v82
                                                                }
                                                            case "":
                                                                var v83 uint16 = 65535
                                                                _ = v83
                                                            default:
                                                                var v84 int = 1
                                                                _ = v84
                                                            }
                                                            var v85 int16 = -32768
                                                            _ = v85
                                                        } else {
                                                            v22 = float64(3.14159)
                                                            var v86 uint64 = (v23 + (v23 % ((v23) | 1)))
                                                            _ = v86
                                                        }
                                                    }
                                                    v1[0] = (fi74 * fi57)
                                                    fmt.Println(v23)
                                                } else {
                                                    var v87 uint16 = 520
                                                    v88 := GenS9{F0: (v2 + v2), F1: ""}
                                                    v23 += 264
                                                    _ = v87
                                                    _ = v88
                                                }
                                                switch (fi57 % (((-(fi74))) | 1)) {
                                                case 32768:
                                                    v76 = v76
                                                }
                                                _ = v76
                                            }
                                        }
                                        _ = fi74
                                    }
                                case 128:
                                    if true {
                                        v2 -= v2
                                        var v89 int64 = 9223372036854775807
                                        var v90 uint32 = 42
                                        _ = v89
                                        _ = v90
                                    } else {
                                        for fi91 := 0; fi91 < 4; fi91++ {
                                            v23 += (v23 & 0)
                                            fmt.Println(fi57)
                                            v22 = v22
                                            switch (-(fi91)) {
                                            case 42:
                                                v92 := GenS13{F0: 0, F1: -128}
                                                _ = v92
                                            case -626:
                                                _ = len(v1)
                                            default:
                                                switch "\n\t" {
                                                case "":
                                                    if true {
                                                        switch fi57 {
                                                        case 127:
                                                            fmt.Println(v2, v23)
                                                        default:
                                                            v22 = v22
                                                        }
                                                    } else {
                                                        v1[2] = fi57
                                                    }
                                                case "foo bar":
                                                    var v93 float32 = 1e10
                                                    _ = v93
                                                case "hello":
                                                    switch fi91 {
                                                    case 2147483647:
                                                        v94 := GenS14{F0: (^(v4))}
                                                        _ = v94
                                                    case 4294967295:
                                                        v1[0] = (fi91 ^ (fi57 - fi91))
                                                    case 1:
                                                        var v95 int16 = 42
                                                        _ = v95
                                                    default:
                                                        v96 := GenS13{F0: 2, F1: (v4 % ((0) | 1))}
                                                        _ = v96
                                                    }
                                                default:
                                                    var v97 uint64 = 0
                                                    _ = v97
                                                }
                                            }
                                            v23 = (v23 * (v23 % (((^(v23))) | 1)))
                                            _ = fi91
                                        }
                                        switch "hello" {
                                        case "":
                                            var v98 int16 = -1
                                            _ = v98
                                        case "hello":
                                            if len(v3) > 0 {
                                                v3[0] = 832
                                            }
                                        case "a":
                                            _ = len(v3)
                                        }
                                        switch 2147483647 {
                                        case -58:
                                            switch ((("hello" + ("hello" + "")) + ("" + "")) + "foo bar") {
                                            case "foo bar":
                                                v2 -= (v2 * (v2 - v2))
                                            }
                                        case 42:
                                            switch ((((("\n\t" + ("foo bar" + ("a" + ""))) + "") + ("\n\t" + "hello")) + "\n\t") + ("" + ("\n\t" + ""))) {
                                            case "\n\t":
                                                var v99 int16 = 1
                                                _ = v99
                                            case "hello":
                                                v2 -= (v2 & v2)
                                            case "a":
                                                var v100 int64 = 42
                                                _ = v100
                                            default:
                                                if (false || true) {
                                                    var v101 uint16 = 0
                                                    for fi102 := 0; fi102 < 8; fi102++ {
                                                        if true {
                                                            v103 := GenS15{F0: v4}
                                                            _ = v103
                                                        }
                                                        _ = fi102
                                                    }
                                                    _ = v101
                                                }
                                            }
                                        }
                                    }
                                }
                            }
                        case "":
                            v104 := GenS14{F0: v4}
                            _ = v104
                        default:
                            v105 := GenS9{F0: 128, F1: ("hello" + "")}
                            _ = v105
                        }
                        _ = v72
                    }
                    _ = fi57
                }
                var v106 string = ("\n\t" + "foo bar")
                var v107 float32 = 3.14159
                _ = v106
                _ = v107
            }
            switch -1 {
            case 2147483647:
                var v108 bool = false
                _ = v108
            case -9223372036854775808:
                for fi109 := 0; fi109 < 6; fi109++ {
                    switch ("a" + "hello") {
                    case "":
                        v2 = uint8(v23)
                    default:
                        delete(v1, 49)
                    }
                    var v110 int8 = v4
                    if (0.0 <= 0.1) {
                        if (0.0 != float32((v22 / v22))) {
                            switch -251 {
                            case -1:
                                _ = len(v3)
                            case -684:
                                if (!(v22 < v22)) {
                                    if true {
                                        v110 = (-(v4))
                                    }
                                } else {
                                    switch fi109 {
                                    case -1:
                                        v1[4294967295] = (-(fi109))
                                    case 32768:
                                        for fi111 := 0; fi111 < 17; fi111++ {
                                            var v112 string = "\n\t"
                                            _ = fi111
                                            _ = v112
                                        }
                                    }
                                    v23 -= 786
                                }
                            case -729:
                                v113 := GenS12{F0: -1, F1: -32768}
                                _ = v113
                            }
                        } else {
                            v114 := GenS13{F0: 42, F1: (v4 >> ((1) & 63))}
                            v4 -= 100
                            _ = v114
                        }
                        fmt.Println(v22, fi109, v22)
                    } else {
                        v4 = -1
                        switch "a" {
                        case "a":
                            v110 = v4
                        case "foo bar":
                            for fi115 := 0; fi115 < 12; fi115++ {
                                var v116 int8 = -23
                                _ = fi115
                                _ = v116
                            }
                        default:
                            if true {
                                fmt.Println(v2, v23)
                                fmt.Println(v110, v4)
                            }
                        }
                        if len(v3) > 0 {
                            v3[0] = 4294967295
                        }
                    }
                    _ = fi109
                    _ = v110
                }
            default:
                var v117 int64 = 90
                _ = v117
            }
        } else {
            if true {
                v22 += v22
            } else {
                fmt.Println(v2)
            }
            for fi118 := 0; fi118 < 3; fi118++ {
                if len(v3) > 0 {
                    v3[0] = 607
                }
                var v119 bool = (fi118 >= int(v2))
                _ = fi118
                _ = v119
            }
            var v120 int = 837
            _ = v120
        }
    case "a":
        var v121 int = 42
        _ = v121
    case "hello":
        v22 = v22
    }
    if false {
        fmt.Println(v2)
        var v122 int = 4294967295
        _ = v122
    }
    v123 := GenS16{F0: v2, F1: (v23 | (v23 | 1)), F2: -528}
    for fi124 := 0; fi124 < 12; fi124++ {
        var v125 uint64 = (^(v23))
        _ = fi124
        _ = v125
    }
    var v126 int = 4294967295
    for fi127 := 0; fi127 < 17; fi127++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi128 := 0; fi128 < 16; fi128++ {
            v2 += (^(v2))
            var v129 uint16 = 32767
            v130 := GenS17{F0: 1, F1: (!(!true))}
            _ = fi128
            _ = v129
            _ = v130
        }
        v4 = (v4 * v4)
        _ = fi127
    }
    if len(v3) > 0 {
        v3[0] = 36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v126:", v126)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v123
    _ = v126
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
