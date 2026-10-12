package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint; F1 bool; F2 int }
type GenS4 struct { F0 int32; F1 string; F2 uint }
type GenS5 struct { F0 string; F1 int64 }
type GenS6 struct { F0 uint; F1 int16 }
type GenS7 struct { F0 uint32; F1 uint; F2 uint64 }
type GenS8 struct { F0 uint32; F1 uint64 }
type GenS9 struct { F0 bool; F1 uint8; F2 int8 }
type GenS10 struct { F0 uint; F1 int16 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 int16 }
type GenS13 struct { F0 uint16; F1 bool; F2 int8 }
type GenS14 struct { F0 uint }
func genFn0() string {
    var v5 int32 = 1
    v6 := []int{32767}
    fmt.Println(v5)
    for fi7 := 0; fi7 < 15; fi7++ {
        v5 = 42
        _ = fi7
    }
    _ = v5
    _ = v6
    return ""
}
func genFn1(p0 int32, p1 int16, p2 uint) uint8 {
    var v8 string = "a"
    var v9 float32 = -2.25
    v10 := []int{127, -2147483648}
    var v11 float64 = (1e10 + -2.25)
    for fi12 := 0; fi12 < 11; fi12++ {
        var v13 float64 = v11
        for fi14 := 0; fi14 < 6; fi14++ {
            if (false && true) {
                v10 = append(v10, 1)
                var v15 int32 = (p0 / ((411) | 1))
                _ = v15
            } else {
                switch (v8 + v8) {
                case "":
                    v16 := GenS1{F0: -1, F1: (fi14 | (-(fi14))), F2: true}
                    _ = v16
                case "hello":
                    fmt.Println(fi12)
                case "\n\t":
                    var v17 int16 = p1
                    _ = v17
                default:
                    var v18 uint8 = 0
                    _ = v18
                }
                var v19 uint16 = 651
                fmt.Println(p1, p2, v9)
                _ = v19
            }
            v8 = ("foo bar" + "")
            p1 = 42
            _ = fi14
        }
        switch (fi12 % (((fi12 | 648)) | 1)) {
        case 128:
            v13 -= ((-2.25 * float64(v9)) * (v13 - v11))
        case -74:
            v20 := GenS2{F0: (p1 - (p1 & 32767))}
            _ = v20
        case 1:
            var v21 int = fi12
            _ = v21
        }
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return 127
}

func main() {
    var v1 int32 = 1
    var v2 int16 = 42
    v3 := make(map[uint32]bool)
    var v4 bool = (1.5 <= -2.25)
    if (false && ((42 != 65535) && ((2 <= 255) && v4))) {
        switch 963 {
        case 128:
            if true {
                var v22 string = ("hello" + "a")
                if false {
                    if false {
                        v22 = "a"
                        v23 := GenS1{F0: 1, F1: -540, F2: v4}
                        _ = v23
                    }
                    if false {
                        v24 := GenS3{F0: 18446744073709551615, F1: v4, F2: -582}
                        if (!(v22 < "\n\t")) {
                            var v25 int8 = 127
                            var v26 int8 = v25
                            _ = v25
                            _ = v26
                        } else {
                            for fi27 := 0; fi27 < 17; fi27++ {
                                if (42 != 339) {
                                    if (v4 && true) {
                                        fmt.Println(v2, v2)
                                    }
                                } else {
                                    fmt.Println(v1)
                                }
                                var v28 int64 = 607
                                _ = fi27
                                _ = v28
                            }
                            for fi29 := 0; fi29 < 16; fi29++ {
                                v4 = v4
                                for fi30 := 0; fi30 < 9; fi30++ {
                                    var v31 string = v22
                                    if (!(v4 || true)) {
                                        var v32 uint32 = 994
                                        v33 := GenS4{F0: 808, F1: (v22 + ((v22 + ((v31 + v31) + (((v22 + "hello") + "\n\t") + (v31 + (v22 + ((v31 + "a") + "\n\t")))))) + "a")), F2: 42}
                                        v22 = v31
                                        _ = v32
                                        _ = v33
                                    } else {
                                        var v34 int = (-(fi30))
                                        v35 := GenS1{F0: 0, F1: (fi29 * 32768), F2: v4}
                                        if false {
                                            v22 = v31
                                        } else {
                                            switch "" {
                                            case "foo bar":
                                                v36 := GenS4{F0: (v1 & v1), F1: "", F2: 2}
                                                _ = v36
                                            default:
                                                if v4 {
                                                    v4 = false
                                                    v22 = v22
                                                }
                                            }
                                            v37 := GenS5{F0: v31, F1: 0}
                                            var v38 int16 = (v2 % ((v2) | 1))
                                            _ = v37
                                            _ = v38
                                        }
                                        _ = v34
                                        _ = v35
                                    }
                                    _ = fi30
                                    _ = v31
                                }
                                _ = fi29
                            }
                            var v39 int8 = 1
                            _ = v39
                        }
                        if true {
                            for fi40 := 0; fi40 < 17; fi40++ {
                                v4 = v4
                                if (v22 <= ((v22 + (((v22 + ("foo bar" + v22)) + ("\n\t" + "")) + ("\n\t" + v22))) + v22)) {
                                    v41 := GenS1{F0: 100, F1: (fi40 / ((-673) | 1)), F2: v4}
                                    var v42 uint = 362
                                    v43 := GenS6{F0: v42, F1: v2}
                                    _ = v41
                                    _ = v42
                                    _ = v43
                                } else {
                                    for fi44 := 0; fi44 < 18; fi44++ {
                                        switch v22 {
                                        case "a":
                                            v45 := GenS6{F0: 18446744073709551615, F1: v2}
                                            _ = v45
                                        case "foo bar":
                                            var v46 int8 = -51
                                            _ = v46
                                        case "hello":
                                            v47 := []int{}
                                            _ = v47
                                        default:
                                            v48 := GenS7{F0: 660, F1: 255, F2: 149}
                                            _ = v48
                                        }
                                        switch "\n\t" {
                                        case "foo bar":
                                            v49 := GenS3{F0: 18446744073709551615, F1: (!(798 != -2147483648)), F2: (^(fi40))}
                                            _ = v49
                                        case "hello":
                                            if (!(1 != 42)) {
                                                var v50 float64 = -2.25
                                                v51 := GenS8{F0: 38, F1: 823}
                                                fmt.Println(fi40)
                                                _ = v50
                                                _ = v51
                                            }
                                        case "\n\t":
                                            switch v22 {
                                            case "":
                                                var v52 float64 = -2.25
                                                _ = v52
                                            }
                                        default:
                                            var v53 bool = (541 > 18446744073709551615)
                                            _ = v53
                                        }
                                        _ = fi44
                                    }
                                    v54 := []int{127}
                                    switch fi40 {
                                    case -143:
                                        v4 = false
                                    case 917:
                                        v2 = v2
                                    default:
                                        v55 := v3[2]
                                        _ = v55
                                    }
                                    _ = v54
                                }
                                _ = fi40
                            }
                            v22 = ("foo bar" + (("hello" + ((("\n\t" + "") + "a") + v22)) + "hello"))
                            v56 := []int{32767, 42, 186, 32767}
                            _ = v56
                        }
                        _ = v24
                    } else {
                        var v57 uint8 = 67
                        v58 := GenS9{F0: (!v4), F1: v57, F2: 42}
                        _ = v57
                        _ = v58
                    }
                    var v59 string = v22
                    _ = v59
                }
                _ = v22
            } else {
                var v60 uint8 = 0
                switch "foo bar" {
                case "":
                    for fi61 := 0; fi61 < 12; fi61++ {
                        v62 := []int{fi61, -1}
                        v2 += -1
                        _ = fi61
                        _ = v62
                    }
                case "hello":
                    fmt.Println(v4)
                case "\n\t":
                    v2 += (v2 - 858)
                }
                var v63 int16 = v2
                _ = v60
                _ = v63
            }
        case 32767:
            fmt.Println(v1)
        case 2147483647:
            v64 := GenS2{F0: v2}
            _ = v64
        default:
            switch -2147483648 {
            case 400:
                v65 := []int{-957, -701, -9223372036854775808, -765}
                _ = v65
            case 32767:
                for fi66 := 0; fi66 < 7; fi66++ {
                    v67 := []int{fi66, fi66, (fi66 ^ (fi66 & (fi66 >> ((255) & 63))))}
                    _ = fi66
                    _ = v67
                }
            case 127:
                v2 = (v2 * v2)
            default:
                v68 := GenS10{F0: 18446744073709551615, F1: v2}
                _ = v68
            }
        }
        var v69 int32 = (v1 % (((^(v1))) | 1))
        _ = v69
    } else {
        v2 += (v2 ^ v2)
        switch ("hello" + ((((("foo bar" + "foo bar") + (("\n\t" + ("" + ("hello" + "a"))) + "")) + "hello") + "") + "")) {
        case "\n\t":
            var v70 int = 2147483647
            _ = v70
        }
        for fi71 := 0; fi71 < 6; fi71++ {
            for fi72 := 0; fi72 < 6; fi72++ {
                v1 -= (v1 % (((v1 | (v1 % (((v1 - v1)) | 1)))) | 1))
                _ = fi72
            }
            for fi73 := 0; fi73 < 19; fi73++ {
                var v74 int16 = 1
                v75 := GenS6{F0: 18446744073709551615, F1: (v2 - (v2 ^ (v74 * 1000)))}
                if (-128 < -107) {
                    v76 := []int{(-(fi73)), (fi71 + fi71), 1}
                    v2 += v2
                    _ = v76
                } else {
                    v77 := GenS5{F0: "", F1: 1}
                    _ = v77
                }
                _ = fi73
                _ = v74
                _ = v75
            }
            _ = fi71
        }
    }
    var v78 uint = 1
    var v79 uint = (v78 << (((v78 << ((4294967295) & 63))) & 63))
    fmt.Println(v79)
    if v4 {
        fmt.Println(v79)
    }
    for fi80 := 0; fi80 < 17; fi80++ {
        var v81 uint16 = 0
        _ = fi80
        _ = v81
    }
    var v82 int32 = v1
    v79 += (v78 >> ((v78) & 63))
    if v4 {
        v83 := GenS9{F0: (!v4), F1: 255, F2: 1}
        v79 -= 4294967295
        for fi84 := 0; fi84 < 3; fi84++ {
            for fi85 := 0; fi85 < 15; fi85++ {
                switch "foo bar" {
                case "":
                    if true {
                        var v86 int8 = -48
                        _ = v86
                    }
                    v87 := GenS9{F0: (1 != 805), F1: 176, F2: 0}
                    if (!(-1 > v1)) {
                        if (v2 < (v2 % (((v2 ^ v2)) | 1))) {
                            var v88 int8 = 42
                            v89 := []int{int(v88), 128}
                            v1 = -659
                            _ = v88
                            _ = v89
                        } else {
                            var v90 int = -233
                            _ = v90
                        }
                        if true {
                            if v4 {
                                v78 = (v78 * (v78 >> (((v78 << ((255) & 63))) & 63)))
                            } else {
                                if true {
                                    v91 := GenS11{F0: 1}
                                    _ = v91
                                } else {
                                    fmt.Println(v78, v1)
                                    var v92 uint8 = genFn1((-(v1)), (v2 << ((2) & 63)), v78)
                                    var v93 float64 = 1.5
                                    _ = v92
                                    _ = v93
                                }
                                var v94 int32 = v82
                                if ((!v4) && v4) {
                                    v94 = v82
                                    v95 := v3[65535]
                                    var v96 int8 = -94
                                    _ = v95
                                    _ = v96
                                } else {
                                    v97 := GenS5{F0: ((("hello" + ("\n\t" + "a")) + "") + (("" + "a") + "foo bar")), F1: 1}
                                    _ = v97
                                }
                                _ = v94
                            }
                            delete(v3, 65535)
                            v98 := GenS4{F0: (v1 << ((v78) & 63)), F1: "", F2: (v79 * (v78 + v78))}
                            _ = v98
                        } else {
                            switch (fi84 | (fi85 | fi85)) {
                            case 0:
                                v1 = (^(v82))
                            case -369:
                                v99 := []int{(^(fi85)), (fi85 - fi84)}
                                _ = v99
                            case 42:
                                v4 = (!v4)
                            default:
                                var v100 int16 = (v2 + 1000)
                                _ = v100
                            }
                        }
                        v101 := v3[65535]
                        _ = v101
                    } else {
                        v102 := []int{(fi85 | fi85)}
                        var v103 int8 = -128
                        _ = v102
                        _ = v103
                    }
                    _ = v87
                }
                fmt.Println(v4, v79, v82)
                var v104 float64 = 0.1
                _ = fi85
                _ = v104
            }
            v105 := []int{fi84, 1, -292, fi84}
            _ = fi84
            _ = v105
        }
        _ = v83
    } else {
        var v106 float32 = (1.5 + 0.0)
        v107 := GenS2{F0: (v2 * v2)}
        _ = v106
        _ = v107
    }
    v108 := []int{-9223372036854775808, 4294967295, -704, int(v1)}
    if false {
        fmt.Println(v79)
    } else {
        for fi109 := 0; fi109 < 4; fi109++ {
            for fi110 := 0; fi110 < 18; fi110++ {
                v78 = 718
                _ = fi110
            }
            _ = fi109
        }
        var v111 int64 = 1
        _ = v111
    }
    if v4 {
        v79 = (v78 << ((v79) & 63))
        var v112 bool = v4
        _ = v112
    }
    switch "foo bar" {
    case "a":
        var v113 int64 = -9223372036854775808
        _ = v113
    case "\n\t":
        v114 := GenS12{F0: 1}
        _ = v114
    case "hello":
        if (v2 < (-(v2))) {
            for fi115 := 0; fi115 < 1; fi115++ {
                _ = len(v108)
                fmt.Println(v2)
                var v116 int64 = 1
                _ = fi115
                _ = v116
            }
        }
    }
    switch ("hello" + "foo bar") {
    case "foo bar":
        if (v4 || ((1 != -1) || true)) {
            v82 += (^(v82))
        }
    }
    switch "\n\t" {
    case "\n\t":
        switch 128 {
        case -775:
            v117 := GenS2{F0: v2}
            _ = v117
        case -9223372036854775808:
            v1 = (v82 / ((v82) | 1))
        case 2147483647:
            v82 -= v1
        }
    case "":
        if ((!v4) || (!(32767 != 1))) {
            fmt.Println(v79)
            for fi118 := 0; fi118 < 3; fi118++ {
                for fi119 := 0; fi119 < 10; fi119++ {
                    var v120 uint32 = 2147483648
                    switch fi118 {
                    case 42:
                        v120 = v120
                    case 32767:
                        v121 := v3[v120]
                        _ = v121
                    }
                    if (!(!((473 <= (^(v120))) || v4))) {
                        if (false || (2 == 583)) {
                            v120 = v120
                            if v4 {
                                var v122 uint = (^(v79))
                                v108 = append(v108, fi118)
                                fmt.Println(v1, v78, v82)
                                _ = v122
                            }
                            v123 := GenS7{F0: (v120 / (((^(v120))) | 1)), F1: (v79 & v78), F2: 18446744073709551615}
                            v79 += (v78 - (^(v79)))
                            fmt.Println(fi119, fi118, v4)
                            switch ("a" + "hello") {
                            case "hello":
                                fmt.Println(v78, v1)
                            case "foo bar":
                                v1 -= (v82 * 1)
                            case "":
                                fmt.Println(v2, v78)
                            }
                            var v124 int64 = 9223372036854775807
                            _ = v123
                            _ = v124
                        } else {
                            var v125 uint = 42
                            v126 := GenS12{F0: v2}
                            _ = len(v108)
                            v127 := GenS13{F0: 0, F1: (false || v4), F2: 53}
                            _ = v125
                            _ = v126
                            _ = v127
                        }
                    } else {
                        if (v4 && true) {
                            var v128 int32 = 129
                            _ = v128
                        } else {
                            v82 = v82
                            fmt.Println(v78, v82, v78)
                            var v129 float32 = 0.1
                            _ = v129
                        }
                    }
                    _ = fi119
                    _ = v120
                }
                v130 := GenS14{F0: v78}
                _ = fi118
                _ = v130
            }
            _ = len(v108)
        } else {
            for fi131 := 0; fi131 < 5; fi131++ {
                v1 = (v82 / ((-521) | 1))
                _ = len(v3)
                for fi132 := 0; fi132 < 13; fi132++ {
                    switch fi132 {
                    case -229:
                        var v133 uint8 = 42
                        _ = v133
                    }
                    v1 -= (v82 / ((-970) | 1))
                    fmt.Println(v2)
                    _ = fi132
                }
                _ = fi131
            }
        }
    case "foo bar":
        v134 := GenS1{F0: -128, F1: 127, F2: v4}
        _ = v134
    default:
        switch "hello" {
        case "\n\t":
            if (42 != 504) {
                var v135 uint16 = 257
                v136 := GenS3{F0: 0, F1: v4, F2: -9223372036854775808}
                var v137 uint64 = 358
                _ = v135
                _ = v136
                _ = v137
            } else {
                if v4 {
                    v138 := GenS2{F0: 340}
                    v79 -= v78
                    _ = v138
                } else {
                    for fi139 := 0; fi139 < 5; fi139++ {
                        fmt.Println(v4, v4)
                        _ = fi139
                    }
                    var v140 int16 = (v2 / ((289) | 1))
                    if len(v108) >= 2 {
                        v108 = v108[1:len(v108)]
                    }
                    _ = v140
                }
                if len(v108) >= 2 {
                    v108 = v108[1:len(v108)]
                }
                fmt.Println(v4, v1)
            }
        case "":
            v1 -= -2147483648
        default:
            v2 = -1
        }
    }
    if ((9223372036854775807 <= 1) || v4) {
        if (genFn1(v82, (v2 << ((v79) & 63)), 65535) <= 42) {
            var v141 uint16 = 0
            _ = v141
        }
        var v142 int16 = v2
        v143 := GenS1{F0: 1, F1: -9223372036854775808, F2: (!true)}
        _ = v142
        _ = v143
    }
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v108:", len(v108), v108)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v108
    _ = v2
    _ = v3
    _ = v4
    _ = v78
    _ = v79
    _ = v82
}
