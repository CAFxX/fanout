package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int16; F1 int8 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint64; F1 uint16; F2 uint32 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 uint; F1 uint8; F2 int32 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 uint16; F1 uint8; F2 int64 }
type GenS13 struct { F0 int32; F1 uint16; F2 int64 }
type GenS14 struct { F0 uint64; F1 int }
func genFn0(p0 int8) uint8 {
    var v5 int64 = 691
    for fi6 := 0; fi6 < 9; fi6++ {
        var v7 uint16 = 0
        for fi8 := 0; fi8 < 4; fi8++ {
            v5 = (v5 << ((466) & 63))
            _ = fi8
        }
        var v9 int64 = -9223372036854775808
        _ = fi6
        _ = v7
        _ = v9
    }
    fmt.Println(v5)
    for fi10 := 0; fi10 < 3; fi10++ {
        switch (fi10 << ((803) & 63)) {
        case 128:
            v5 = v5
        default:
            p0 -= (p0 - (p0 % (((p0 * p0)) | 1)))
        }
        if true {
            var v11 int32 = int32(v5)
            if false {
                var v12 bool = ("foo bar" == ((("" + "") + "\n\t") + "\n\t"))
                if (v5 < (v5 % (((^(v5))) | 1))) {
                    var v13 bool = v12
                    if (200 > 65535) {
                        p0 -= (p0 & (p0 | (p0 % (((^(p0))) | 1))))
                    }
                    _ = v13
                }
                for fi14 := 0; fi14 < 20; fi14++ {
                    if (541 == 65535) {
                        var v15 uint64 = 4294967295
                        if v12 {
                            v16 := []int{(fi10 % (((-(fi10))) | 1)), (fi14 / (((fi10 - (fi14 + (-(fi10))))) | 1))}
                            fmt.Println(v11, p0)
                            _ = v16
                        } else {
                            var v17 int16 = 32767
                            _ = v17
                        }
                        v18 := []int{fi10, -2147483648, fi10, int(p0)}
                        _ = v15
                        _ = v18
                    } else {
                        v19 := GenS1{F0: (v5 * (v5 * int64(v11)))}
                        var v20 uint64 = 1
                        v21 := GenS1{F0: v5}
                        switch fi10 {
                        case -620:
                            var v22 string = ("hello" + ((("foo bar" + ("a" + "")) + ("" + "a")) + "a"))
                            _ = v22
                        case 290:
                            v23 := []int{(fi10 + int(p0)), (fi14 >> ((255) & 63)), (fi14 ^ fi10), (-(fi14))}
                            _ = v23
                        case 478:
                            v24 := GenS1{F0: v5}
                            _ = v24
                        default:
                            var v25 uint64 = 4294967295
                            _ = v25
                        }
                        _ = v19
                        _ = v20
                        _ = v21
                    }
                    fmt.Println(fi14, p0)
                    switch "a" {
                    case "":
                        v26 := []int{fi14}
                        _ = v26
                    case "hello":
                        if (false && v12) {
                            v5 = int64(p0)
                            var v27 uint64 = 18446744073709551615
                            _ = v27
                        }
                    }
                    _ = fi14
                }
                _ = v12
            } else {
                v28 := GenS2{F0: "hello", F1: (v11 * 42)}
                var v29 float32 = 3.14159
                var v30 uint16 = 2
                _ = v28
                _ = v29
                _ = v30
            }
            v11 -= v11
            _ = v11
        }
        _ = fi10
    }
    _ = p0
    _ = v5
    return 2
}
func genFn1(p0 int8, p1 int16, p2 uint16) int16 {
    v31 := GenS1{F0: 9223372036854775807}
    var v32 uint8 = 127
    _ = p0
    _ = p1
    _ = p2
    _ = v31
    _ = v32
    return 1000
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 string = ("foo bar" + "\n\t")
    var v3 int16 = 0
    var v4 string = v2
    if false {
        v33 := v1[42]
        if ((v3 >> ((0) & 63)) < (v3 << ((289) & 63))) {
            if v33 {
                var v34 int16 = v3
                _ = v34
            }
        }
        _ = v33
    }
    v2 = v2
    if true {
        var v35 float64 = 1e10
        v36 := GenS2{F0: v4, F1: -512}
        _ = v35
        _ = v36
    }
    v37 := GenS2{F0: v4, F1: 0}
    for fi38 := 0; fi38 < 16; fi38++ {
        var v39 uint16 = 523
        switch "a" {
        case "\n\t":
            if false {
                v40 := []int{128, (^(fi38)), (fi38 / ((fi38) | 1))}
                var v41 int32 = 42
                v2 = ("hello" + v4)
                _ = v40
                _ = v41
            } else {
                var v42 uint16 = (v39 & (v39 ^ v39))
                if (true && true) {
                    if (1 <= -38) {
                        v43 := GenS1{F0: 1}
                        v44 := GenS3{F0: -9223372036854775808}
                        _ = v43
                        _ = v44
                    }
                    for fi45 := 0; fi45 < 3; fi45++ {
                        switch v2 {
                        case "hello":
                            v46 := []int{}
                            _ = v46
                        case "foo bar":
                            fmt.Println(fi38)
                            fmt.Println(v42)
                            v47 := []int{(-(fi45)), fi38, fi38}
                            v4 = (v4 + "")
                            _ = v47
                        case "a":
                            v48 := GenS4{F0: -830}
                            _ = v48
                        default:
                            if true {
                                v2 = "foo bar"
                            } else {
                                switch (v2 + "\n\t") {
                                case "\n\t":
                                    delete(v1, 42)
                                case "a":
                                    if false {
                                        v42 -= v39
                                        var v49 float32 = 3.14159
                                        switch fi45 {
                                        case 940:
                                            if (!(!false)) {
                                                var v50 float64 = -2.25
                                                _ = v50
                                            }
                                        case -1:
                                            v51 := GenS5{F0: v3, F1: 50}
                                            _ = v51
                                        default:
                                            var v52 uint64 = 2
                                            _ = v52
                                        }
                                        _ = v49
                                    }
                                case "hello":
                                    v39 += 255
                                }
                                fmt.Println(v2, v2, v42)
                                var v53 uint64 = 328
                                _ = v53
                            }
                        }
                        var v54 uint32 = 4294967295
                        v39 += 0
                        _ = fi45
                        _ = v54
                    }
                } else {
                    if (true && true) {
                        if ((!true) && true) {
                            switch (fi38 | 684) {
                            case -100:
                                if true {
                                    var v55 int16 = v3
                                    var v56 float32 = 3.14159
                                    _ = v55
                                    _ = v56
                                }
                            case 983:
                                var v57 uint = 0
                                _ = v57
                            default:
                                for fi58 := 0; fi58 < 9; fi58++ {
                                    var v59 string = v2
                                    if (v2 >= "foo bar") {
                                        v60 := GenS6{F0: v2}
                                        v61 := GenS4{F0: -545}
                                        v62 := GenS7{F0: 817}
                                        var v63 int64 = -1
                                        _ = v60
                                        _ = v61
                                        _ = v62
                                        _ = v63
                                    }
                                    switch v59 {
                                    case "":
                                        fmt.Println(fi58, v3, v2)
                                    case "foo bar":
                                        var v64 uint64 = 773
                                        _ = v64
                                    case "\n\t":
                                        if (true || true) {
                                            v4 = "a"
                                            v59 = v4
                                        } else {
                                            v4 = v4
                                        }
                                    default:
                                        v65 := []int{(fi58 ^ fi38), (fi38 ^ -601)}
                                        _ = v65
                                    }
                                    _ = fi58
                                    _ = v59
                                }
                            }
                            v4 = (("" + "hello") + "hello")
                        } else {
                            for fi66 := 0; fi66 < 6; fi66++ {
                                var v67 uint16 = 390
                                var v68 uint32 = 1
                                _ = len(v1)
                                if (!(!(uint8(v67) >= 1))) {
                                    v68 += v68
                                } else {
                                    v42 -= (v39 + (v42 >> ((2) & 63)))
                                    if (!true) {
                                        switch "foo bar" {
                                        case "a":
                                            v2 = ("\n\t" + "\n\t")
                                        default:
                                            if false {
                                                switch fi66 {
                                                case 32768:
                                                    fmt.Println(v42, v2)
                                                    v2 = "hello"
                                                case -1:
                                                    var v69 uint16 = (v42 % ((v67) | 1))
                                                    if true {
                                                        if true {
                                                            if (!false) {
                                                                switch "foo bar" {
                                                                case "a":
                                                                    v70 := GenS6{F0: v2}
                                                                    var v71 int = int(v68)
                                                                    _ = v70
                                                                    _ = v71
                                                                default:
                                                                    var v72 int16 = 42
                                                                    _ = v72
                                                                }
                                                                switch v4 {
                                                                case "\n\t":
                                                                    v67 -= v42
                                                                    var v73 int32 = 1
                                                                    _ = v73
                                                                case "hello":
                                                                    var v74 int16 = (v3 - v3)
                                                                    _ = v74
                                                                default:
                                                                    var v75 uint16 = v42
                                                                    _ = v75
                                                                }
                                                                if false {
                                                                    v76 := GenS8{F0: 18446744073709551615, F1: v39, F2: 943}
                                                                    _ = v76
                                                                }
                                                            }
                                                            fmt.Println(v39)
                                                            v3 = -1
                                                            fmt.Println(v42, v4, v68)
                                                            var v77 string = "a"
                                                            _ = v77
                                                        } else {
                                                            if (42 >= 1) {
                                                                if true {
                                                                    var v78 uint = 369
                                                                    _ = v78
                                                                } else {
                                                                    var v79 int64 = 1
                                                                    v80 := []int{(fi66 & (^(fi38)))}
                                                                    v80 = append(v80, fi38)
                                                                    _ = v79
                                                                    _ = v80
                                                                }
                                                                if (51 <= 119) {
                                                                    if ((((false || true) && ("hello" != v2)) && (!((v4 != "a") || true))) && false) {
                                                                        var v81 float64 = 0.0
                                                                        _ = v81
                                                                    }
                                                                    var v82 uint = 18446744073709551615
                                                                    _ = v82
                                                                } else {
                                                                    var v83 uint64 = 18446744073709551615
                                                                    _ = v83
                                                                }
                                                                var v84 uint64 = 18446744073709551615
                                                                var v85 int32 = 596
                                                                v69 += 879
                                                                switch ((v2 + (v4 + (v4 + (("" + "hello") + "a")))) + "foo bar") {
                                                                case "":
                                                                    v85 -= (v85 & 0)
                                                                default:
                                                                    if ((v68 != (v68 + 4294967295)) && (true && (!(fi66 >= fi66)))) {
                                                                        var v86 float32 = 0.1
                                                                        if false {
                                                                            v68 = (v68 - (v68 ^ v68))
                                                                            var v87 uint32 = v68
                                                                            var v88 uint64 = 632
                                                                            fmt.Println(v85, v84, v4)
                                                                            _ = v87
                                                                            _ = v88
                                                                        }
                                                                        _ = v86
                                                                    }
                                                                }
                                                                _ = v84
                                                                _ = v85
                                                            } else {
                                                                v68 = v68
                                                                var v89 int8 = 9
                                                                _ = v89
                                                            }
                                                            switch fi66 {
                                                            case -366:
                                                                v67 -= v67
                                                            }
                                                        }
                                                    }
                                                    var v90 uint64 = 1
                                                    _ = v69
                                                    _ = v90
                                                }
                                                switch (fi66 - (fi66 * (fi66 >> ((552) & 63)))) {
                                                case 445:
                                                    switch ((v4 + (v2 + v4)) + (("a" + v4) + "hello")) {
                                                    case "a":
                                                        var v91 int32 = -18
                                                        _ = v91
                                                    default:
                                                        v39 = (v39 * (^(v42)))
                                                    }
                                                    v92 := GenS4{F0: -1}
                                                    _ = v92
                                                case 1:
                                                    if (4294967295 >= 1) {
                                                        var v93 uint32 = 42
                                                        v94 := []int{fi38, fi66, (-(fi66))}
                                                        _ = v93
                                                        _ = v94
                                                    } else {
                                                        v95 := []int{(fi66 & fi38), (fi38 >> ((18446744073709551615) & 63))}
                                                        if ((255 > 18446744073709551615) || true) {
                                                            var v96 uint16 = (v39 * v67)
                                                            _ = v96
                                                        } else {
                                                            v2 = ("\n\t" + v2)
                                                            _ = len(v95)
                                                            if ((^(v3)) > (v3 ^ genFn1(127, (v3 - v3), (^(v39))))) {
                                                                _ = len(v1)
                                                                v97 := GenS4{F0: -2147483648}
                                                                _ = v97
                                                            } else {
                                                                var v98 string = "\n\t"
                                                                fmt.Println(fi38, v39)
                                                                var v99 int64 = 351
                                                                _ = v98
                                                                _ = v99
                                                            }
                                                            switch "\n\t" {
                                                            case "a":
                                                                switch "" {
                                                                case "hello":
                                                                    if ((v2 + "hello") != v4) {
                                                                        var v100 string = v2
                                                                        v101 := GenS1{F0: 9223372036854775807}
                                                                        _ = v100
                                                                        _ = v101
                                                                    } else {
                                                                        fmt.Println(v68)
                                                                        fmt.Println(v68)
                                                                    }
                                                                case "a":
                                                                    v1[(v68 + (v68 / ((uint32(v39)) | 1)))] = (0 < 812)
                                                                    var v102 bool = ((!true) || false)
                                                                    var v103 int8 = -1
                                                                    _ = v102
                                                                    _ = v103
                                                                case "foo bar":
                                                                    if ((42 <= 0) || false) {
                                                                        var v104 string = v4
                                                                        if true {
                                                                            switch v2 {
                                                                            case "\n\t":
                                                                                switch -1 {
                                                                                case 42:
                                                                                    v42 -= (v39 / ((1) | 1))
                                                                                case 2147483647:
                                                                                    var v105 string = (v4 + "")
                                                                                    _ = v105
                                                                                case 421:
                                                                                    v106 := GenS2{F0: v104, F1: -1}
                                                                                    _ = v106
                                                                                default:
                                                                                    v104 = v104
                                                                                }
                                                                            case "a":
                                                                                if ((!true) && (255 == 1)) {
                                                                                    delete(v1, v68)
                                                                                    v104 = v2
                                                                                    v3 = v3
                                                                                } else {
                                                                                    var v107 int8 = -128
                                                                                    _ = v107
                                                                                }
                                                                            }
                                                                            var v108 float64 = -2.25
                                                                            v109 := GenS9{F0: (fi66 >> ((1) & 63))}
                                                                            _ = v108
                                                                            _ = v109
                                                                        } else {
                                                                            var v110 uint = 1
                                                                            _ = v110
                                                                        }
                                                                        fmt.Println(fi66, v4, v68)
                                                                        v42 = v42
                                                                        _ = v104
                                                                    }
                                                                }
                                                            case "foo bar":
                                                                v3 += (v3 * (-(v3)))
                                                            }
                                                            var v111 uint = 1
                                                            v112 := GenS9{F0: 44}
                                                            v39 = v67
                                                            _ = v111
                                                            _ = v112
                                                        }
                                                        switch 0 {
                                                        case 9223372036854775807:
                                                            var v113 string = v4
                                                            _ = v113
                                                        case 1:
                                                            if (((0.0 * 0.0) * 0.0) == (((1.5 + 1e10) - 0.0) * 1e10)) {
                                                                v2 = (v4 + "\n\t")
                                                                fmt.Println(v4, v39)
                                                            } else {
                                                                var v114 int32 = -2147483648
                                                                var v115 int8 = 1
                                                                v2 = ("a" + "\n\t")
                                                                _ = v114
                                                                _ = v115
                                                            }
                                                        }
                                                        if len(v95) > 0 {
                                                            v95[0] = (fi66 / ((-334) | 1))
                                                        }
                                                        _ = v95
                                                    }
                                                case 32767:
                                                    var v116 int8 = int8(fi38)
                                                    _ = v116
                                                default:
                                                    switch v4 {
                                                    case "hello":
                                                        if true {
                                                            v4 = v4
                                                            v117 := []int{}
                                                            _ = v117
                                                        } else {
                                                            v1[(v68 / (((v68 << ((887) & 63))) | 1))] = true
                                                            var v118 int8 = -128
                                                            var v119 int32 = -2147483648
                                                            if (false && false) {
                                                                switch ((((("" + v4) + (v4 + v4)) + v2) + ((("a" + (v4 + ("\n\t" + ((v2 + ("\n\t" + v2)) + ("hello" + ("" + v4)))))) + (("\n\t" + ((v2 + "\n\t") + ("" + v2))) + "hello")) + ("" + v2))) + v4) {
                                                                case "":
                                                                    fmt.Println(v3, fi38, v68)
                                                                case "hello":
                                                                    v120 := v1[876]
                                                                    _ = v120
                                                                case "\n\t":
                                                                    v4 = "a"
                                                                default:
                                                                    var v121 uint64 = 4294967295
                                                                    _ = v121
                                                                }
                                                            } else {
                                                                var v122 int32 = 0
                                                                _ = v122
                                                            }
                                                            _ = v118
                                                            _ = v119
                                                        }
                                                    case "":
                                                        v67 = (v39 ^ v39)
                                                    case "\n\t":
                                                        v39 = v42
                                                    default:
                                                        if (3.14159 <= float64(0.0)) {
                                                            if (18446744073709551615 > 2) {
                                                                if (0.1 == -2.25) {
                                                                    delete(v1, v68)
                                                                    v123 := GenS10{F0: uint(v39), F1: 1, F2: 0}
                                                                    if ((v3 != -677) || true) {
                                                                        v67 = (v67 & 0)
                                                                        if false {
                                                                            var v124 uint32 = v68
                                                                            v125 := GenS11{F0: -9223372036854775808}
                                                                            v4 = (("foo bar" + v4) + v4)
                                                                            _ = v124
                                                                            _ = v125
                                                                        } else {
                                                                            v3 -= (^(v3))
                                                                            switch ((((v4 + (v4 + v2)) + v2) + v4) + ("hello" + v4)) {
                                                                            case "":
                                                                                fmt.Println(v3, v67, v39)
                                                                            case "\n\t":
                                                                                delete(v1, v68)
                                                                            }
                                                                        }
                                                                        if true {
                                                                            fmt.Println(v67, v67)
                                                                            fmt.Println(v39, v42, v4)
                                                                        }
                                                                    } else {
                                                                        v2 = v2
                                                                    }
                                                                    fmt.Println(v67, v39)
                                                                    _ = v123
                                                                }
                                                                v2 = v2
                                                            }
                                                            v67 -= (v67 % ((v39) | 1))
                                                        }
                                                    }
                                                    if (1 == uint(v68)) {
                                                        if true {
                                                            switch fi38 {
                                                            case 204:
                                                                v126 := GenS12{F0: v39, F1: 0, F2: 726}
                                                                _ = v126
                                                            case 440:
                                                                fmt.Println(fi66)
                                                            case -722:
                                                                v127 := GenS5{F0: 42, F1: 13}
                                                                _ = v127
                                                            }
                                                            v128 := []int{(fi38 & -9223372036854775808), fi38, fi38}
                                                            if (v4 != (v4 + ((v4 + v4) + "foo bar"))) {
                                                                v129 := v1[v68]
                                                                _ = v129
                                                            }
                                                            if false {
                                                                fmt.Println(v2, v67, v3)
                                                                var v130 int32 = int32(fi38)
                                                                var v131 int8 = -1
                                                                var v132 string = "a"
                                                                _ = v130
                                                                _ = v131
                                                                _ = v132
                                                            } else {
                                                                v128 = append(v128, (fi38 * (fi38 + 738)))
                                                                var v133 int16 = v3
                                                                switch (fi38 - (fi38 >> ((1) & 63))) {
                                                                case 1:
                                                                    var v134 uint = 18446744073709551615
                                                                    _ = v134
                                                                case -9223372036854775808:
                                                                    var v135 uint8 = 255
                                                                    _ = v135
                                                                default:
                                                                    _ = len(v128)
                                                                }
                                                                fmt.Println(v4, fi38)
                                                                _ = v133
                                                            }
                                                            _ = v128
                                                        }
                                                    }
                                                    fmt.Println(fi66, v42)
                                                }
                                            } else {
                                                var v136 uint = 349
                                                v137 := GenS12{F0: v39, F1: genFn0(100), F2: 1}
                                                _ = v136
                                                _ = v137
                                            }
                                            if false {
                                                if (-1 >= -125) {
                                                    v138 := v1[v68]
                                                    switch (^(fi38)) {
                                                    case 9223372036854775807:
                                                        fmt.Println(v42)
                                                    case 127:
                                                        var v139 float32 = 1.5
                                                        _ = v139
                                                    }
                                                    _ = v138
                                                }
                                                v67 -= 42
                                                if true {
                                                    var v140 int8 = -128
                                                    v141 := []int{}
                                                    _ = v140
                                                    _ = v141
                                                } else {
                                                    fmt.Println(v2, v3)
                                                    var v142 int16 = genFn1(-1, v3, v67)
                                                    _ = v142
                                                }
                                            }
                                        }
                                        v3 += (v3 - v3)
                                        v67 -= v39
                                    } else {
                                        fmt.Println(v42, fi38, v68)
                                        v4 = (("\n\t" + ("" + (v4 + v4))) + (v2 + "a"))
                                        var v143 int32 = 139
                                        v39 = (v39 | (v42 & v67))
                                        _ = v143
                                    }
                                }
                                v3 = (v3 >> ((1) & 63))
                                v144 := GenS6{F0: ""}
                                var v145 string = v2
                                _ = fi66
                                _ = v144
                                _ = v145
                                _ = v67
                                _ = v68
                            }
                        }
                        v146 := []int{(fi38 / ((fi38) | 1)), fi38}
                        var v147 float64 = 3.14159
                        _ = v146
                        _ = v147
                    }
                    for fi148 := 0; fi148 < 16; fi148++ {
                        var v149 int8 = 100
                        v150 := []int{(-(fi148)), -9223372036854775808}
                        _ = fi148
                        _ = v149
                        _ = v150
                    }
                }
                _ = v42
            }
        default:
            var v151 int8 = 0
            _ = v151
        }
        v152 := GenS13{F0: 2147483647, F1: 961, F2: 42}
        _ = fi38
        _ = v152
        _ = v39
    }
    v153 := GenS14{F0: 2, F1: 2147483647}
    fmt.Println(v3)
    v3 -= (v3 - genFn1(1, genFn1(119, (v3 ^ v3), 604), 255))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v153
    _ = v2
    _ = v3
    _ = v37
    _ = v4
}
