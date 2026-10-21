package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 uint8 }
type GenS2 struct { F0 uint8; F1 uint8; F2 uint64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int16; F1 int; F2 int8 }
type GenS5 struct { F0 uint64; F1 uint8 }
type GenS6 struct { F0 uint64; F1 string; F2 int32 }
type GenS7 struct { F0 int16; F1 int; F2 bool }
type GenS8 struct { F0 int; F1 bool; F2 int32 }
type GenS9 struct { F0 int32; F1 uint; F2 string }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 int16; F1 uint8 }
type GenS12 struct { F0 uint32; F1 int16; F2 bool }
type GenS13 struct { F0 string }
type GenS14 struct { F0 int16; F1 uint64 }
type GenS15 struct { F0 uint32; F1 int16; F2 int }
func genFn0() uint32 {
    var v6 int = 18
    for fi7 := 0; fi7 < 13; fi7++ {
        switch "a" {
        case "a":
            v6 -= (fi7 | 32768)
        }
        _ = fi7
    }
    if (0.0 < float32(0.1)) {
        for fi8 := 0; fi8 < 15; fi8++ {
            v6 -= (-(fi8))
            switch ("hello" + "hello") {
            case "hello":
                v9 := []int{v6, fi8, (v6 * (v6 | 0)), 394}
                _ = v9
            case "foo bar":
                v10 := GenS1{F0: 413, F1: (100 == 127), F2: 55}
                _ = v10
            case "\n\t":
                switch 729 {
                case -125:
                    if ((0 < 2) && false) {
                        var v11 int64 = 9223372036854775807
                        _ = v11
                    }
                case 127:
                    switch "a" {
                    case "\n\t":
                        for fi12 := 0; fi12 < 14; fi12++ {
                            v13 := GenS2{F0: 74, F1: 127, F2: 4294967295}
                            _ = fi12
                            _ = v13
                        }
                    }
                default:
                    for fi14 := 0; fi14 < 17; fi14++ {
                        v15 := []int{(v6 + 642), fi14, -74, (fi14 & -2147483648)}
                        _ = fi14
                        _ = v15
                    }
                }
            default:
                v16 := []int{(-(fi8)), fi8}
                _ = v16
            }
            _ = fi8
        }
        v17 := []int{(v6 / (((v6 + v6)) | 1)), (v6 << ((0) & 63)), 127}
        _ = v17
    }
    switch "\n\t" {
    case "foo bar":
        var v18 string = ""
        _ = v18
    case "hello":
        v6 -= v6
    default:
        var v19 float64 = 0.1
        _ = v19
    }
    for fi20 := 0; fi20 < 18; fi20++ {
        v6 -= v6
        var v21 int8 = -1
        _ = fi20
        _ = v21
    }
    _ = v6
    return 1
}
func genFn1(p0 uint64, p1 uint32, p2 bool) uint16 {
    v22 := GenS3{F0: p0}
    if (p2 && (127 <= 218)) {
        var v23 float32 = 0.0
        for fi24 := 0; fi24 < 14; fi24++ {
            switch "hello" {
            case "a":
                var v25 string = "\n\t"
                _ = v25
            case "foo bar":
                p0 -= (p0 >> ((4294967295) & 63))
            default:
                for fi26 := 0; fi26 < 8; fi26++ {
                    var v27 float64 = 0.0
                    if (2 == 18446744073709551615) {
                        var v28 int16 = 478
                        _ = v28
                    } else {
                        fmt.Println(v23, v23, fi26)
                    }
                    _ = fi26
                    _ = v27
                }
            }
            var v29 int8 = -2
            _ = fi24
            _ = v29
        }
        v30 := GenS4{F0: 42, F1: 666, F2: -128}
        _ = v23
        _ = v30
    }
    switch "a" {
    case "foo bar":
        switch 32768 {
        case 46:
            var v31 string = "\n\t"
            _ = v31
        case -9223372036854775808:
            fmt.Println(p2)
        default:
            v32 := []int{-2147483648, 4294967295, 9223372036854775807, -9223372036854775808}
            _ = v32
        }
    case "hello":
        v33 := GenS3{F0: 42}
        _ = v33
    default:
        var v34 int16 = 0
        _ = v34
    }
    for fi35 := 0; fi35 < 2; fi35++ {
        switch "a" {
        case "":
            var v36 uint32 = p1
            _ = v36
        case "hello":
            for fi37 := 0; fi37 < 8; fi37++ {
                var v38 uint64 = (p0 / ((1) | 1))
                _ = fi37
                _ = v38
            }
        }
        switch "a" {
        case "a":
            if (-2147483648 >= 2147483647) {
                var v39 float32 = (1e10 + 0.0)
                _ = v39
            } else {
                var v40 uint32 = genFn0()
                _ = v40
            }
        default:
            v41 := GenS1{F0: -1, F1: (!true), F2: 2}
            _ = v41
        }
        _ = fi35
    }
    for fi42 := 0; fi42 < 11; fi42++ {
        v43 := []int{(fi42 ^ (fi42 | (fi42 ^ fi42))), fi42, int(p0), fi42}
        v44 := GenS5{F0: p0, F1: 127}
        _ = fi42
        _ = v43
        _ = v44
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    return 65535
}

func main() {
    var v1 uint32 = 969
    var v2 int8 = -62
    var v3 bool = (100 != v2)
    v4 := []uint32{}
    v5 := make(map[uint32]bool)
    fmt.Println(v3, v2)
    fmt.Println(v3)
    for fi45 := 0; fi45 < 6; fi45++ {
        v46 := GenS6{F0: 299, F1: "hello", F2: 2147483647}
        _ = fi45
        _ = v46
    }
    v5[2147483648] = false
    switch 32768 {
    case 32768:
        for fi47 := 0; fi47 < 19; fi47++ {
            v2 = -128
            switch "\n\t" {
            case "foo bar":
                var v48 int = (fi47 ^ (^(fi47)))
                _ = v48
            case "\n\t":
                if v3 {
                    for fi49 := 0; fi49 < 1; fi49++ {
                        switch "\n\t" {
                        case "foo bar":
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        case "\n\t":
                            switch fi49 {
                            case -957:
                                var v50 float64 = (1.5 * (1.5 - 1e10))
                                _ = v50
                            default:
                                v51 := GenS3{F0: 520}
                                _ = v51
                            }
                        case "hello":
                            v4 = append(v4, v1)
                        }
                        _ = fi49
                    }
                } else {
                    if (18446744073709551615 >= 0) {
                        if (1.5 <= 0.1) {
                            var v52 uint64 = 4294967295
                            v53 := GenS7{F0: 247, F1: fi47, F2: (!((-128 > (v2 & -128)) || false))}
                            if (-1 <= 42) {
                                if v3 {
                                    var v54 uint32 = (v1 ^ v1)
                                    _ = v54
                                } else {
                                    v2 -= (v2 % (((v2 - v2)) | 1))
                                    v3 = true
                                }
                                var v55 uint16 = 304
                                _ = v55
                            } else {
                                if (false && true) {
                                    v56 := GenS8{F0: fi47, F1: ((9223372036854775807 == 453) || (!v3)), F2: -1}
                                    _ = v56
                                }
                            }
                            _ = v52
                            _ = v53
                        }
                        for fi57 := 0; fi57 < 1; fi57++ {
                            _ = len(v5)
                            if false {
                                v1 = 319
                                switch fi57 {
                                case 127:
                                    v58 := GenS8{F0: fi57, F1: v3, F2: -58}
                                    _ = v58
                                default:
                                    if ("hello" == "hello") {
                                        var v59 string = ""
                                        v4 = append(v4, genFn0())
                                        var v60 uint64 = 18446744073709551615
                                        _ = v59
                                        _ = v60
                                    }
                                }
                                fmt.Println(v3)
                            }
                            v1 += (v1 - v1)
                            _ = fi57
                        }
                        v3 = true
                    } else {
                        v61 := v5[(v1 % ((v1) | 1))]
                        v3 = v61
                        if (true && true) {
                            fmt.Println(v2, v1)
                            v62 := GenS6{F0: 18446744073709551615, F1: ("\n\t" + ((("a" + "") + "hello") + ("foo bar" + "\n\t"))), F2: -608}
                            _ = v62
                        }
                        _ = v61
                    }
                    if v3 {
                        v63 := GenS9{F0: 0, F1: 908, F2: (((("\n\t" + "\n\t") + ("hello" + "a")) + "hello") + "a")}
                        for fi64 := 0; fi64 < 20; fi64++ {
                            if v3 {
                                if v3 {
                                    v65 := GenS10{F0: v3}
                                    v2 = v2
                                    var v66 int32 = 1
                                    _ = v65
                                    _ = v66
                                } else {
                                    var v67 bool = v3
                                    _ = len(v5)
                                    _ = v67
                                }
                                var v68 uint32 = 42
                                v5[v68] = true
                                var v69 int32 = 1
                                _ = v68
                                _ = v69
                            }
                            v70 := GenS3{F0: 18446744073709551615}
                            _ = fi64
                            _ = v70
                        }
                        for fi71 := 0; fi71 < 6; fi71++ {
                            v2 = v2
                            _ = fi71
                        }
                        _ = v63
                    } else {
                        v2 = 1
                        var v72 string = ("" + "hello")
                        v3 = v3
                        _ = v72
                    }
                }
            case "a":
                var v73 int64 = -9223372036854775808
                _ = v73
            default:
                for fi74 := 0; fi74 < 17; fi74++ {
                    v2 = v2
                    _ = fi74
                }
            }
            _ = fi47
        }
    case -9223372036854775808:
        v1 = (v1 | genFn0())
    default:
        v3 = v3
    }
    for fi75 := 0; fi75 < 2; fi75++ {
        v76 := GenS11{F0: 42, F1: 163}
        v77 := GenS3{F0: 2}
        for fi78 := 0; fi78 < 7; fi78++ {
            var v79 string = "hello"
            if (128 > 32768) {
                var v80 int = -282
                v2 = v2
                fmt.Println(v80, v2, v1)
                var v81 int32 = 42
                _ = v80
                _ = v81
            } else {
                v1 += (v1 << ((18446744073709551615) & 63))
                v82 := GenS8{F0: fi75, F1: (!(!(!v3))), F2: -393}
                _ = v82
            }
            v1 -= (v1 % ((2147483648) | 1))
            _ = fi78
            _ = v79
        }
        _ = fi75
        _ = v76
        _ = v77
    }
    for fi83 := 0; fi83 < 18; fi83++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v84 uint8 = 2
        for fi85 := 0; fi85 < 11; fi85++ {
            if (true && (-1 <= fi83)) {
                fmt.Println(v3, v84, v2)
                var v86 uint32 = v1
                v1 -= (v86 - (v1 / ((v86) | 1)))
                v87 := GenS7{F0: 0, F1: (fi85 | fi85), F2: (!(fi83 <= fi85))}
                if len(v4) > 0 {
                    v4[0] = (v86 << ((0) & 63))
                }
                _ = v86
                _ = v87
            } else {
                switch (fi83 | -903) {
                case -9223372036854775808:
                    fmt.Println(v2, v84, v2)
                case 42:
                    fmt.Println(fi83)
                case -1:
                    if len(v4) > 0 {
                        v4[0] = (^(v1))
                    }
                default:
                    v88 := v5[v1]
                    _ = v88
                }
            }
            _ = fi85
        }
        _ = fi83
        _ = v84
    }
    switch ("" + ("foo bar" + "hello")) {
    case "foo bar":
        for fi89 := 0; fi89 < 5; fi89++ {
            v90 := GenS7{F0: 257, F1: fi89, F2: (!v3)}
            _ = fi89
            _ = v90
        }
    case "a":
        v91 := GenS6{F0: 42, F1: ((("hello" + "\n\t") + "a") + (("a" + "") + ("foo bar" + ("" + "foo bar")))), F2: 42}
        _ = v91
    case "hello":
        fmt.Println(v2)
    }
    for fi92 := 0; fi92 < 8; fi92++ {
        v2 += int8(v1)
        var v93 int32 = 217
        if v3 {
            for fi94 := 0; fi94 < 19; fi94++ {
                delete(v5, 2147483647)
                v3 = true
                switch "" {
                case "foo bar":
                    var v95 int = (fi92 + -2147483648)
                    _ = v95
                case "\n\t":
                    var v96 float64 = ((0.0 - 3.14159) + 1.5)
                    _ = v96
                case "hello":
                    switch (fi92 & (fi92 & (^(fi94)))) {
                    case 718:
                        switch (fi94 + 1) {
                        case 32767:
                            _ = len(v4)
                        }
                    case 4294967295:
                        if v3 {
                            v97 := GenS8{F0: 367, F1: false, F2: 42}
                            if v3 {
                                v1 = v1
                                v98 := GenS12{F0: 2147483648, F1: 538, F2: (!v3)}
                                _ = v98
                            } else {
                                switch "a" {
                                case "hello":
                                    var v99 uint64 = 2
                                    _ = v99
                                case "a":
                                    if (!((v93 % (((v93 & 431)) | 1)) < -2147483648)) {
                                        if v3 {
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            switch "hello" {
                                            case "a":
                                                v2 += (v2 + (v2 ^ (v2 / (((v2 % ((v2) | 1))) | 1))))
                                            default:
                                                if v3 {
                                                    v3 = (v3 && (953 == 983))
                                                    var v100 float64 = 1e10
                                                    v100 = v100
                                                    _ = v100
                                                } else {
                                                    fmt.Println(fi92, v93, v93)
                                                }
                                            }
                                            v1 += genFn0()
                                        } else {
                                            var v101 bool = false
                                            switch (-(fi92)) {
                                            case 167:
                                                switch fi92 {
                                                case -1:
                                                    _ = len(v4)
                                                case -592:
                                                    if v101 {
                                                        var v102 uint64 = 438
                                                        _ = v102
                                                    } else {
                                                        v101 = true
                                                    }
                                                case 2147483647:
                                                    var v103 int = (fi92 << ((42) & 63))
                                                    _ = v103
                                                default:
                                                    if len(v4) > 0 {
                                                        v4[0] = genFn0()
                                                    }
                                                }
                                            case 0:
                                                var v104 int = fi92
                                                _ = v104
                                            case 254:
                                                fmt.Println(fi94, v2)
                                            }
                                            _ = v101
                                        }
                                        if v3 {
                                            if ((v3 || (!(!false))) && (20 != 255)) {
                                                v2 = v2
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                var v105 int64 = 42
                                                _ = v105
                                            } else {
                                                fmt.Println(v3)
                                                v2 = (v2 & 0)
                                                fmt.Println(fi94)
                                                v5[v1] = (!((!(!(!(!v3)))) || v3))
                                            }
                                            v93 -= v93
                                        } else {
                                            var v106 uint16 = 932
                                            v2 += (v2 % ((v2) | 1))
                                            v1 = (v1 * (v1 + genFn0()))
                                            _ = v106
                                        }
                                        switch ("" + "") {
                                        case "\n\t":
                                            switch "\n\t" {
                                            case "\n\t":
                                                v3 = ((3.14159 < 0.1) && (150 > 604))
                                            default:
                                                fmt.Println(fi94, v1, v1)
                                            }
                                        case "foo bar":
                                            fmt.Println(v93, v93, v1)
                                        }
                                    }
                                case "\n\t":
                                    v2 = v2
                                }
                                var v107 int = fi94
                                _ = len(v4)
                                _ = v107
                            }
                            _ = v97
                        }
                    case 789:
                        if v3 {
                            var v108 uint16 = 65535
                            v109 := GenS1{F0: 1, F1: v3, F2: 1}
                            _ = v108
                            _ = v109
                        } else {
                            var v110 string = ("foo bar" + "a")
                            var v111 uint16 = genFn1(18446744073709551615, v1, (v110 < (v110 + (("" + (v110 + v110)) + v110))))
                            _ = v110
                            _ = v111
                        }
                    }
                }
                var v112 int16 = -32768
                v113 := v5[v1]
                var v114 string = "hello"
                v115 := GenS11{F0: (v112 ^ (v112 & v112)), F1: 42}
                _ = fi94
                _ = v112
                _ = v113
                _ = v114
                _ = v115
            }
            var v116 uint16 = 1
            switch (((("\n\t" + (("hello" + "foo bar") + "hello")) + "a") + (("a" + "foo bar") + ("hello" + "hello"))) + "a") {
            case "\n\t":
                if false {
                    switch (fi92 & int(v1)) {
                    case -290:
                        if (!((fi92 ^ fi92) == fi92)) {
                            v2 = (^(v2))
                            for fi117 := 0; fi117 < 5; fi117++ {
                                if v3 {
                                    v3 = false
                                    switch "\n\t" {
                                    case "":
                                        v2 = (-(v2))
                                    case "a":
                                        var v118 int8 = (v2 ^ (^(v2)))
                                        _ = v118
                                    default:
                                        fmt.Println(fi117, v116, fi117)
                                    }
                                    var v119 uint = 2
                                    _ = v119
                                }
                                _ = fi117
                            }
                            v3 = (!v3)
                        }
                    case -1:
                        for fi120 := 0; fi120 < 2; fi120++ {
                            if (631 < 729) {
                                delete(v5, v1)
                            } else {
                                v121 := GenS13{F0: ("" + "\n\t")}
                                _ = v121
                            }
                            switch (fi120 - fi92) {
                            case 32767:
                                var v122 int64 = 9223372036854775807
                                _ = v122
                            case 127:
                                v123 := GenS5{F0: 928, F1: 2}
                                _ = v123
                            default:
                                switch fi120 {
                                case -2147483648:
                                    var v124 bool = false
                                    _ = v124
                                case 9223372036854775807:
                                    var v125 uint16 = genFn1(111, v1, ((!false) || (859 == 65535)))
                                    _ = v125
                                default:
                                    var v126 int32 = -2147483648
                                    _ = v126
                                }
                            }
                            switch "\n\t" {
                            case "":
                                if (v3 && (0 >= 2)) {
                                    if len(v4) > 0 {
                                        v4[0] = v1
                                    }
                                    v127 := GenS14{F0: -32768, F1: 1}
                                    fmt.Println(fi120)
                                    _ = v127
                                }
                            case "foo bar":
                                v128 := GenS15{F0: genFn0(), F1: 714, F2: fi92}
                                _ = v128
                            case "hello":
                                var v129 int64 = 1
                                _ = v129
                            }
                            _ = fi120
                        }
                    default:
                        var v130 int32 = 2147483647
                        _ = v130
                    }
                } else {
                    v131 := GenS1{F0: 1, F1: (132 != 224), F2: 1}
                    _ = v131
                }
            default:
                fmt.Println(v2, v1, v93)
            }
            _ = v116
        } else {
            switch ("a" + "a") {
            case "\n\t":
                _ = len(v4)
            case "hello":
                v1 -= 703
            }
            v3 = v3
            var v132 uint32 = (v1 % ((v1) | 1))
            _ = v132
        }
        _ = fi92
        _ = v93
    }
    fmt.Println(v3, v2, v3)
    for fi133 := 0; fi133 < 1; fi133++ {
        v3 = true
        fmt.Println(v1, v1)
        _ = fi133
    }
    switch 42 {
    case 42:
        for fi134 := 0; fi134 < 12; fi134++ {
            v3 = false
            var v135 uint32 = 2
            v135 -= v135
            _ = fi134
            _ = v135
        }
    }
    v3 = ((((1e10 - 3.14159) + -2.25) == 1.5) && v3)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v136 := GenS15{F0: v1, F1: 1, F2: -2147483648}
    var v137 int64 = -9223372036854775808
    if v3 {
        fmt.Println(v3, v1, v1)
        fmt.Println(v2, v3, v137)
    }
    for fi138 := 0; fi138 < 11; fi138++ {
        var v139 int = (-(fi138))
        for fi140 := 0; fi140 < 12; fi140++ {
            v2 = v2
            v4 = append(v4, v1)
            _ = fi140
        }
        _ = fi138
        _ = v139
    }
    fmt.Println("v1:", v1)
    fmt.Println("v137:", v137)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v136
    _ = v137
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
