package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 string; F2 bool }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 string; F1 int64; F2 uint }
type GenS7 struct { F0 uint; F1 uint32; F2 int16 }
type GenS8 struct { F0 int64; F1 int; F2 uint8 }
type GenS9 struct { F0 int32; F1 int32 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 string; F1 uint64 }
func genFn0(p0 int16) int16 {
    p0 = (-(p0))
    for fi7 := 0; fi7 < 13; fi7++ {
        v8 := GenS1{F0: ((3.14159 <= 1.5) && false)}
        if (573 < 2) {
            switch (fi7 / ((128) | 1)) {
            case -1:
                for fi9 := 0; fi9 < 14; fi9++ {
                    p0 = (p0 >> ((436) & 63))
                    v10 := GenS1{F0: false}
                    p0 -= p0
                    var v11 float32 = 3.14159
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            default:
                fmt.Println(fi7, p0)
            }
            fmt.Println(p0, p0)
            v12 := GenS2{F0: "hello", F1: "\n\t", F2: ((^(fi7)) <= (fi7 ^ 215))}
            _ = v12
        } else {
            p0 = (p0 + p0)
        }
        _ = fi7
        _ = v8
    }
    for fi13 := 0; fi13 < 9; fi13++ {
        var v14 bool = false
        _ = fi13
        _ = v14
    }
    v15 := GenS2{F0: "", F1: (("hello" + "foo bar") + ""), F2: ((-2.25 + 0.0) == 1.5)}
    _ = p0
    _ = v15
    return 809
}

func main() {
    var v1 int16 = 0
    var v2 uint = 295
    var v3 int64 = -679
    var v4 uint8 = 197
    v5 := []uint32{1, 4294967295}
    var v6 int64 = -156
    if (true || false) {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    } else {
        for fi16 := 0; fi16 < 8; fi16++ {
            if true {
                var v17 int16 = (^(v1))
                _ = v17
            } else {
                v2 = (v2 ^ v2)
                if false {
                    v6 += (v6 * v6)
                    v3 -= v6
                    switch "foo bar" {
                    case "foo bar":
                        _ = len(v5)
                    case "a":
                        v6 += (v6 | (-(v6)))
                    case "":
                        var v18 int = 128
                        _ = v18
                    }
                } else {
                    switch ("foo bar" + "hello") {
                    case "a":
                        _ = len(v5)
                    case "\n\t":
                        var v19 uint32 = 65535
                        _ = v19
                    case "hello":
                        var v20 string = "hello"
                        _ = v20
                    }
                    var v21 uint16 = 32767
                    var v22 int8 = 0
                    _ = v21
                    _ = v22
                }
            }
            var v23 float64 = ((float64(0.1) - 1.5) * (-2.25 / 1e10))
            if (!false) {
                v3 = int64(fi16)
            } else {
                v5 = append(v5, 679)
            }
            _ = fi16
            _ = v23
        }
        for fi24 := 0; fi24 < 1; fi24++ {
            switch fi24 {
            case 818:
                var v25 uint64 = 905
                _ = v25
            case 1:
                var v26 float32 = 1.5
                _ = v26
            case 9223372036854775807:
                fmt.Println(v6)
            }
            var v27 uint64 = 1
            if len(v5) > 0 {
                v5[0] = 683
            }
            _ = fi24
            _ = v27
        }
        for fi28 := 0; fi28 < 6; fi28++ {
            if (0 == 4294967295) {
                var v29 float64 = 0.1
                _ = v29
            } else {
                fmt.Println(v1, fi28, v2)
                for fi30 := 0; fi30 < 15; fi30++ {
                    v4 = v4
                    v31 := GenS1{F0: ((!true) && (v4 >= (^(v4))))}
                    switch fi30 {
                    case 32767:
                        switch "foo bar" {
                        case "\n\t":
                            var v32 bool = true
                            _ = v32
                        case "hello":
                            if (!false) {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            }
                        case "a":
                            fmt.Println(v2)
                        }
                    case -1:
                        if false {
                            var v33 uint64 = 2
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            _ = v33
                        }
                    case 493:
                        var v34 int = 253
                        _ = v34
                    default:
                        var v35 float32 = 3.14159
                        _ = v35
                    }
                    _ = fi30
                    _ = v31
                }
            }
            _ = fi28
        }
    }
    v6 -= v6
    var v36 uint16 = 65535
    for fi37 := 0; fi37 < 9; fi37++ {
        var v38 float64 = 1.5
        if len(v5) > 0 {
            v5[0] = 65535
        }
        v36 -= 0
        _ = fi37
        _ = v38
    }
    v2 = v2
    switch "hello" {
    case "":
        if true {
            v5 = append(v5, 4294967295)
            v2 += (v2 * 255)
            v36 = 42
        } else {
            for fi39 := 0; fi39 < 2; fi39++ {
                if false {
                    v6 = v6
                    v6 += 0
                    for fi40 := 0; fi40 < 7; fi40++ {
                        v1 = v1
                        _ = fi40
                    }
                }
                var v41 float32 = float32(0.0)
                _ = fi39
                _ = v41
            }
            v2 = (^(v2))
            for fi42 := 0; fi42 < 19; fi42++ {
                switch 128 {
                case 128:
                    switch "foo bar" {
                    case "":
                        var v43 int16 = v1
                        _ = v43
                    case "a":
                        for fi44 := 0; fi44 < 12; fi44++ {
                            switch "\n\t" {
                            case "a":
                                if true {
                                    var v45 int = (fi42 / (((-(fi44))) | 1))
                                    fmt.Println(v3)
                                    _ = v45
                                }
                            }
                            _ = fi44
                        }
                    }
                case -2147483648:
                    switch -1 {
                    case 2147483647:
                        var v46 int = 907
                        _ = v46
                    case 127:
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    case -1:
                        for fi47 := 0; fi47 < 16; fi47++ {
                            if false {
                                if (uint8(v36) == (v4 << (((v2 % ((0) | 1))) & 63))) {
                                    fmt.Println(v6)
                                } else {
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                }
                                var v48 int8 = 100
                                v49 := GenS2{F0: "\n\t", F1: "hello", F2: (v6 != (-(v6)))}
                                _ = v48
                                _ = v49
                            } else {
                                var v50 float32 = 3.14159
                                var v51 int64 = (v3 << ((v2) & 63))
                                if false {
                                    var v52 int8 = 0
                                    _ = v52
                                }
                                _ = v50
                                _ = v51
                            }
                            _ = fi47
                        }
                    }
                }
                v53 := GenS3{F0: 2, F1: ""}
                _ = fi42
                _ = v53
            }
        }
    case "foo bar":
        v54 := GenS4{F0: v36}
        _ = v54
    case "hello":
        if false {
            switch "" {
            case "":
                var v55 uint64 = 803
                _ = v55
            }
            for fi56 := 0; fi56 < 18; fi56++ {
                switch fi56 {
                case 669:
                    if (!(!true)) {
                        v2 = (v2 >> ((65535) & 63))
                        if false {
                            for fi57 := 0; fi57 < 17; fi57++ {
                                var v58 float32 = (float32(0.0) / 1e10)
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                _ = fi57
                                _ = v58
                            }
                            for fi59 := 0; fi59 < 7; fi59++ {
                                fmt.Println(v36)
                                if (!(!true)) {
                                    fmt.Println(v2)
                                    var v60 uint64 = 42
                                    _ = v60
                                }
                                if true {
                                    var v61 uint8 = (v4 - 189)
                                    _ = v61
                                } else {
                                    v36 = v36
                                }
                                v62 := GenS5{F0: 255}
                                if true {
                                    _ = len(v5)
                                    v63 := GenS6{F0: "", F1: v3, F2: v2}
                                    v3 = (^(v6))
                                    switch 32767 {
                                    case 0:
                                        v36 = v36
                                    }
                                    _ = v63
                                }
                                _ = fi59
                                _ = v62
                            }
                        }
                    }
                case -9223372036854775808:
                    v64 := GenS2{F0: (("" + ((("a" + ("foo bar" + "")) + ("hello" + ("a" + "hello"))) + "\n\t")) + "foo bar"), F1: "hello", F2: (!((true || (!(v36 != (v36 << ((940) & 63))))) || true))}
                    _ = v64
                }
                _ = fi56
            }
            if (!(22 >= 4294967295)) {
                if len(v5) > 0 {
                    v5[0] = 0
                }
            }
        } else {
            v1 += v1
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v4 -= (v4 % (((v4 + v4)) | 1))
        }
    default:
        fmt.Println(v36)
    }
    if false {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v65 := GenS5{F0: (^(v4))}
        v3 += v6
        _ = v65
    }
    if true {
        for fi66 := 0; fi66 < 9; fi66++ {
            if true {
                for fi67 := 0; fi67 < 2; fi67++ {
                    var v68 uint8 = (^(v4))
                    _ = fi67
                    _ = v68
                }
            } else {
                for fi69 := 0; fi69 < 8; fi69++ {
                    v70 := GenS6{F0: "", F1: (-(v6)), F2: v2}
                    if true {
                        var v71 uint = (v2 | (v2 + v2))
                        v72 := GenS2{F0: "foo bar", F1: ("a" + "hello"), F2: (-2.25 >= (float32(0.1) + ((0.0 / 3.14159) * -2.25)))}
                        _ = v71
                        _ = v72
                    } else {
                        var v73 uint16 = 582
                        if false {
                            var v74 uint16 = v73
                            fmt.Println(v2, v6, v74)
                            if (!((1 >= 2) && (!(v4 == 2)))) {
                                v75 := GenS7{F0: (v2 << ((4294967295) & 63)), F1: 1, F2: genFn0(-1)}
                                v5 = append(v5, 2147483647)
                                _ = v75
                            } else {
                                v73 -= (v74 ^ 255)
                                if true {
                                    v5 = append(v5, 2)
                                    v2 += v2
                                } else {
                                    if true {
                                        var v76 float32 = 3.14159
                                        v74 = v74
                                        fmt.Println(v3, v74)
                                        v77 := GenS4{F0: v73}
                                        var v78 float64 = 1e10
                                        _ = v76
                                        _ = v77
                                        _ = v78
                                    } else {
                                        v4 += (v4 | (v4 % ((v4) | 1)))
                                        var v79 int8 = -128
                                        v80 := GenS3{F0: 2, F1: "foo bar"}
                                        if false {
                                            if (((0.1 + 1e10) < 1.5) && ((((v79 * (v79 * (v79 >> ((0) & 63)))) > v79) && (1e10 > (-2.25 / 0.1))) || ((0.0 * 3.14159) > (1e10 / 3.14159)))) {
                                                var v81 uint16 = v73
                                                _ = v81
                                            } else {
                                                if len(v5) > 0 {
                                                    v5[0] = 8
                                                }
                                                switch (fi69 | (fi66 % ((fi69) | 1))) {
                                                case 42:
                                                    var v82 string = "foo bar"
                                                    _ = v82
                                                case 9223372036854775807:
                                                    v3 -= (v3 ^ (v3 ^ int64(v2)))
                                                }
                                                if len(v5) > 0 {
                                                    v5[0] = 0
                                                }
                                            }
                                            var v83 int16 = (v1 / (((v1 << ((v2) & 63))) | 1))
                                            switch (fi66 << ((v2) & 63)) {
                                            case -2147483648:
                                                var v84 int = 32768
                                                _ = v84
                                            }
                                            v74 += (v73 % (((^(v73))) | 1))
                                            _ = v83
                                        }
                                        if true {
                                            v85 := GenS8{F0: (v3 >> ((2) & 63)), F1: (fi69 ^ (fi69 % ((2147483647) | 1))), F2: v4}
                                            v1 += (v1 + (v1 + (v1 / ((v1) | 1))))
                                            if len(v5) > 0 {
                                                v5[0] = 2
                                            }
                                            _ = v85
                                        }
                                        _ = v79
                                        _ = v80
                                    }
                                    if false {
                                        _ = len(v5)
                                    } else {
                                        switch "\n\t" {
                                        case "foo bar":
                                            switch fi66 {
                                            case 62:
                                                var v86 float64 = 0.1
                                                _ = v86
                                            case 32767:
                                                v87 := GenS7{F0: (v2 ^ (v2 / ((768) | 1))), F1: 42, F2: v1}
                                                _ = v87
                                            }
                                        case "a":
                                            v36 = uint16(v2)
                                        case "hello":
                                            v2 = v2
                                        }
                                        if ((v1 != v1) && ((^(v74)) > v74)) {
                                            var v88 float32 = -2.25
                                            _ = v88
                                        } else {
                                            fmt.Println(v2)
                                            fmt.Println(v36, v3, v4)
                                        }
                                        var v89 uint32 = 626
                                        _ = v89
                                    }
                                }
                                if true {
                                    if ((v3 << ((2) & 63)) == v3) {
                                        var v90 uint64 = 541
                                        v36 -= v73
                                        _ = v90
                                    }
                                    v91 := GenS8{F0: v6, F1: (fi66 + 127), F2: v4}
                                    var v92 uint16 = v74
                                    if (!true) {
                                        v93 := GenS9{F0: 42, F1: -2147483648}
                                        var v94 uint64 = 4294967295
                                        _ = v93
                                        _ = v94
                                    }
                                    v6 -= v6
                                    _ = v91
                                    _ = v92
                                }
                            }
                            v3 -= v3
                            switch (fi66 - -9223372036854775808) {
                            case -2147483648:
                                fmt.Println(fi66, v3, v2)
                            case 112:
                                fmt.Println(v6, v1)
                            }
                            if false {
                                if false {
                                    if true {
                                        if false {
                                            v95 := GenS5{F0: (v4 << ((v2) & 63))}
                                            _ = v95
                                        }
                                        fmt.Println(v6, v36, v3)
                                    } else {
                                        v3 += v3
                                        if (((fi69 == fi69) || ((("" + (("foo bar" + "hello") + "hello")) + ("hello" + "hello")) > "\n\t")) || ((-(fi66)) != fi66)) {
                                            if false {
                                                v73 = (v73 << ((1) & 63))
                                            } else {
                                                var v96 int64 = (-(v3))
                                                if (!true) {
                                                    switch 2147483647 {
                                                    case 32768:
                                                        switch int(v2) {
                                                        case 0:
                                                            switch ("hello" + "\n\t") {
                                                            case "hello":
                                                                v3 -= v6
                                                            case "foo bar":
                                                                switch ("foo bar" + (("" + "foo bar") + "\n\t")) {
                                                                case "hello":
                                                                    var v97 uint8 = (v4 & v4)
                                                                    _ = v97
                                                                default:
                                                                    _ = len(v5)
                                                                }
                                                            }
                                                        case 32767:
                                                            _ = len(v5)
                                                        default:
                                                            var v98 uint8 = 2
                                                            _ = v98
                                                        }
                                                    case 731:
                                                        v1 += (v1 | v1)
                                                    default:
                                                        var v99 uint8 = v4
                                                        _ = v99
                                                    }
                                                    v2 += v2
                                                    var v100 uint16 = 255
                                                    v96 += (v3 >> (((v2 % (((v2 % (((v2 / ((v2) | 1))) | 1))) | 1))) & 63))
                                                    _ = v100
                                                } else {
                                                    v3 += v6
                                                    fmt.Println(v36, v6, v73)
                                                }
                                                var v101 int8 = -106
                                                _ = v101
                                                _ = v96
                                            }
                                            fmt.Println(v74, v4)
                                        }
                                    }
                                    switch "a" {
                                    case "foo bar":
                                        var v102 uint32 = 165
                                        _ = v102
                                    case "a":
                                        var v103 int64 = (^(v6))
                                        _ = v103
                                    default:
                                        var v104 uint8 = uint8(v2)
                                        v74 -= (v73 + 65535)
                                        var v105 uint8 = v4
                                        if (false && (((^(v4)) >= (v104 + (v4 >> ((v2) & 63)))) || true)) {
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                        }
                                        var v106 int8 = 32
                                        _ = v104
                                        _ = v105
                                        _ = v106
                                    }
                                    _ = len(v5)
                                } else {
                                    v107 := GenS6{F0: (("\n\t" + "foo bar") + "hello"), F1: -9223372036854775808, F2: (v2 / ((uint(v4)) | 1))}
                                    var v108 float32 = float32(1.5)
                                    v109 := GenS6{F0: "foo bar", F1: -9223372036854775808, F2: (v2 + (v2 % ((0) | 1)))}
                                    _ = v107
                                    _ = v108
                                    _ = v109
                                }
                                var v110 string = "foo bar"
                                var v111 bool = false
                                v112 := GenS1{F0: false}
                                v1 -= genFn0(-32768)
                                var v113 int8 = -1
                                v114 := GenS3{F0: 0, F1: v110}
                                _ = v110
                                _ = v111
                                _ = v112
                                _ = v113
                                _ = v114
                            }
                            _ = v74
                        }
                        _ = v73
                    }
                    if true {
                        var v115 string = ("a" + (("\n\t" + "hello") + "\n\t"))
                        var v116 bool = (float32(0.1) < 0.1)
                        _ = v115
                        _ = v116
                    } else {
                        v6 -= v6
                        var v117 int64 = (v3 >> ((v2) & 63))
                        _ = v117
                    }
                    _ = fi69
                    _ = v70
                }
            }
            switch fi66 {
            case -79:
                var v118 float64 = (-2.25 + (-2.25 + (-2.25 * 3.14159)))
                _ = v118
            }
            var v119 uint = 976
            _ = fi66
            _ = v119
        }
    } else {
        for fi120 := 0; fi120 < 19; fi120++ {
            for fi121 := 0; fi121 < 12; fi121++ {
                if (true && true) {
                    var v122 uint8 = (v4 << ((2) & 63))
                    v4 -= (v122 | (v4 & 174))
                    _ = v122
                }
                v6 = 9223372036854775807
                _ = fi121
            }
            v2 = (v2 | v2)
            _ = fi120
        }
        for fi123 := 0; fi123 < 9; fi123++ {
            _ = len(v5)
            _ = fi123
        }
        var v124 int64 = v3
        _ = v124
    }
    v3 += v6
    v5 = append(v5, uint32(v2))
    v125 := GenS10{F0: (v2 | v2)}
    var v126 uint8 = 255
    for fi127 := 0; fi127 < 18; fi127++ {
        for fi128 := 0; fi128 < 10; fi128++ {
            var v129 uint64 = 42
            if ((v4 / ((v126) | 1)) < (v126 / ((v126) | 1))) {
                var v130 uint = v2
                v5 = append(v5, 1)
                v5 = append(v5, 221)
                v6 = v3
                v3 += 0
                _ = v130
            } else {
                var v131 uint = v2
                var v132 uint32 = 0
                var v133 uint8 = v126
                _ = v131
                _ = v132
                _ = v133
            }
            v129 += 2
            _ = fi128
            _ = v129
        }
        if (v1 > (v1 << ((v2) & 63))) {
            v134 := GenS11{F0: 2147483647}
            v135 := GenS12{F0: 1}
            _ = v134
            _ = v135
        }
        _ = fi127
    }
    var v136 int = 32768
    v6 += -158
    if (-459 > 0) {
        v1 = genFn0(v1)
        for fi137 := 0; fi137 < 12; fi137++ {
            var v138 uint32 = 841
            for fi139 := 0; fi139 < 14; fi139++ {
                if true {
                    v140 := GenS13{F0: "foo bar", F1: 42}
                    var v141 string = ("\n\t" + "")
                    fmt.Println(v36)
                    _ = v140
                    _ = v141
                }
                _ = fi139
            }
            switch "a" {
            case "foo bar":
                for fi142 := 0; fi142 < 8; fi142++ {
                    var v143 float64 = (3.14159 / 0.1)
                    if true {
                        var v144 uint16 = 1
                        v143 -= 1.5
                        _ = v144
                    }
                    _ = fi142
                    _ = v143
                }
            }
            _ = fi137
            _ = v138
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v126:", v126)
    fmt.Println("v136:", v136)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v125
    _ = v126
    _ = v136
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
