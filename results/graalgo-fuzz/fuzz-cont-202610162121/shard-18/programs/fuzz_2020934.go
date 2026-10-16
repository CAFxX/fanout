package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint16; F1 uint32; F2 uint64 }
type GenS5 struct { F0 uint64; F1 int64 }
type GenS6 struct { F0 string; F1 int16; F2 bool }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 bool; F1 int64 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int64; F1 uint32; F2 uint16 }
type GenS11 struct { F0 uint32 }
func genFn0(p0 bool, p1 uint32, p2 string) bool {
    p1 -= (p1 | 65535)
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = p2
    return p0
}

func main() {
    v1 := make(map[string]string)
    v2 := []int{1, -242, 524, -1}
    var v3 int32 = 0
    v4 := make(map[int]bool)
    v5 := make(map[int]string)
    if (42 > 0) {
        if (1 >= 127) {
            var v6 int = 9223372036854775807
            if false {
                for fi7 := 0; fi7 < 10; fi7++ {
                    fmt.Println(v6, fi7, v6)
                    var v8 int8 = 100
                    for fi9 := 0; fi9 < 7; fi9++ {
                        var v10 int32 = (v3 | v3)
                        _ = fi9
                        _ = v10
                    }
                    _ = fi7
                    _ = v8
                }
            }
            var v11 int = v6
            _ = v11
            _ = v6
        }
        if true {
            var v12 float32 = 1e10
            var v13 string = ((("a" + "a") + ("hello" + ("foo bar" + "foo bar"))) + "\n\t")
            var v14 int64 = 728
            _ = v12
            _ = v13
            _ = v14
        }
        if false {
            var v15 string = "hello"
            for fi16 := 0; fi16 < 6; fi16++ {
                v17 := GenS1{F0: -1, F1: 0}
                var v18 uint64 = 4294967295
                _ = fi16
                _ = v17
                _ = v18
            }
            _ = v15
        } else {
            var v19 int16 = 679
            _ = v19
        }
    } else {
        switch -1 {
        case 32768:
            v3 = (v3 + (v3 % ((v3) | 1)))
        case 2147483647:
            if true {
                for fi20 := 0; fi20 < 17; fi20++ {
                    if (42 > 78) {
                        switch ("\n\t" + "\n\t") {
                        case "hello":
                            switch (-(fi20)) {
                            case -1:
                                v2 = append(v2, -9223372036854775808)
                            case 127:
                                v3 = (^(v3))
                            case 4294967295:
                                v3 -= v3
                            }
                        case "\n\t":
                            var v21 uint = 42
                            _ = v21
                        default:
                            for fi22 := 0; fi22 < 11; fi22++ {
                                var v23 bool = ((0.1 / ((0.0 - (1e10 - (((float64(0.0) + (1e10 * 3.14159)) - (0.0 - (-2.25 - 0.0))) - 3.14159))) / 3.14159)) >= (-2.25 - -2.25))
                                if (436 > 42) {
                                    if (1 >= 907) {
                                        v23 = (("foo bar" + "\n\t") == "\n\t")
                                        var v24 uint64 = 989
                                        _ = v24
                                    } else {
                                        _ = len(v2)
                                        v25 := GenS2{F0: (!v23), F1: 28}
                                        _ = v25
                                    }
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    fmt.Println(v23, fi20, fi22)
                                } else {
                                    _ = len(v1)
                                    delete(v5, fi20)
                                }
                                v23 = true
                                _ = fi22
                                _ = v23
                            }
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    switch "hello" {
                    case "":
                        for fi26 := 0; fi26 < 2; fi26++ {
                            v3 += v3
                            var v27 uint = 18446744073709551615
                            v27 += 42
                            _ = fi26
                            _ = v27
                        }
                    case "hello":
                        switch fi20 {
                        case 4294967295:
                            if true {
                                v28 := GenS3{F0: 0}
                                _ = v28
                            } else {
                                v29 := GenS2{F0: (!false), F1: 1}
                                _ = v29
                            }
                        case 1:
                            if (false || ((fi20 ^ -1) < fi20)) {
                                var v30 int64 = 0
                                _ = v30
                            } else {
                                _ = len(v4)
                                v3 = 42
                                var v31 int64 = -9223372036854775808
                                _ = v31
                            }
                        default:
                            if ((-119 == 100) || (false || true)) {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                if true {
                                    for fi32 := 0; fi32 < 5; fi32++ {
                                        if len(v2) > 0 {
                                            v2[0] = (fi32 & 4294967295)
                                        }
                                        delete(v5, 9223372036854775807)
                                        fmt.Println(fi20, fi20)
                                        _ = fi32
                                    }
                                    var v33 uint8 = 255
                                    _ = v33
                                } else {
                                    v3 = (v3 & v3)
                                    v3 -= (v3 - 42)
                                    v34 := GenS3{F0: 42}
                                    _ = v34
                                }
                                if (813 <= (fi20 * fi20)) {
                                    v35 := GenS4{F0: 32767, F1: 42, F2: 607}
                                    _ = v35
                                } else {
                                    var v36 int = fi20
                                    var v37 int16 = 294
                                    v37 = v37
                                    _ = v36
                                    _ = v37
                                }
                            }
                        }
                    }
                    if (104 > 96) {
                        fmt.Println(fi20)
                    } else {
                        v38 := GenS4{F0: 408, F1: 2147483648, F2: 937}
                        fmt.Println(v3, v3)
                        _ = v38
                    }
                    _ = fi20
                }
                fmt.Println(v3)
                if true {
                    fmt.Println(v3)
                    for fi39 := 0; fi39 < 19; fi39++ {
                        fmt.Println(v3, v3)
                        switch "\n\t" {
                        case "":
                            v1[(("foo bar" + "\n\t") + "foo bar")] = "foo bar"
                        case "hello":
                            for fi40 := 0; fi40 < 19; fi40++ {
                                var v41 int32 = 1
                                fmt.Println(v41, fi39, v41)
                                _ = fi40
                                _ = v41
                            }
                        case "a":
                            if (false && (-1 < -38)) {
                                for fi42 := 0; fi42 < 15; fi42++ {
                                    v3 -= (v3 + (v3 - v3))
                                    _ = fi42
                                }
                            }
                        default:
                            switch "" {
                            case "":
                                v3 += (v3 + v3)
                            case "a":
                                if (int32(fi39) < (v3 << ((1) & 63))) {
                                    if false {
                                        var v43 uint64 = 18446744073709551615
                                        v2 = append(v2, fi39)
                                        _ = v43
                                    }
                                    v4[fi39] = (false && (true || true))
                                    if (18446744073709551615 < 779) {
                                        if true {
                                            for fi44 := 0; fi44 < 2; fi44++ {
                                                v3 -= 42
                                                _ = fi44
                                            }
                                            if true {
                                                var v45 int16 = -32768
                                                _ = v45
                                            }
                                        }
                                        v3 = 1
                                        if len(v2) > 0 {
                                            v2[0] = (fi39 * (fi39 | fi39))
                                        }
                                    }
                                }
                            case "hello":
                                for fi46 := 0; fi46 < 5; fi46++ {
                                    v3 = (v3 * v3)
                                    _ = fi46
                                }
                            default:
                                if false {
                                    v3 = 1
                                } else {
                                    v47 := GenS2{F0: true, F1: 0}
                                    for fi48 := 0; fi48 < 12; fi48++ {
                                        v49 := GenS5{F0: 0, F1: 1}
                                        if true {
                                            v3 += 65
                                            v3 = 2147483647
                                            if false {
                                                if true {
                                                    var v50 uint8 = 255
                                                    v51 := GenS6{F0: "\n\t", F1: -32768, F2: (true && (!false))}
                                                    var v52 uint = 255
                                                    _ = v50
                                                    _ = v51
                                                    _ = v52
                                                }
                                                if true {
                                                    v53 := GenS7{F0: 864}
                                                    if len(v2) > 0 {
                                                        v2[0] = fi48
                                                    }
                                                    _ = v53
                                                } else {
                                                    v3 += v3
                                                    var v54 int16 = 1000
                                                    _ = v54
                                                }
                                            } else {
                                                if false {
                                                    v3 = (-(v3))
                                                } else {
                                                    var v55 uint32 = 2147483648
                                                    _ = v55
                                                }
                                            }
                                        }
                                        _ = fi48
                                        _ = v49
                                    }
                                    _ = v47
                                }
                            }
                        }
                        var v56 int32 = v3
                        _ = fi39
                        _ = v56
                    }
                } else {
                    v2 = append(v2, 336)
                    var v57 int16 = 42
                    _ = v57
                }
            } else {
                if ((float32(1e10) <= 1e10) && false) {
                    v58 := GenS8{F0: true, F1: 42}
                    v3 -= (v3 * (v3 / (((v3 * (v3 << ((0) & 63)))) | 1)))
                    v59 := GenS6{F0: "", F1: 1000, F2: ("foo bar" > ("\n\t" + "hello"))}
                    _ = v58
                    _ = v59
                }
            }
        }
        if (-128 >= -1) {
            for fi60 := 0; fi60 < 15; fi60++ {
                delete(v4, fi60)
                for fi61 := 0; fi61 < 19; fi61++ {
                    v3 += 1
                    _ = len(v5)
                    if ((v3 != v3) || false) {
                        if (0 >= (v3 & v3)) {
                            v62 := GenS4{F0: 610, F1: 4294967295, F2: 0}
                            if false {
                                switch "\n\t" {
                                case "\n\t":
                                    var v63 uint16 = 292
                                    _ = v63
                                case "":
                                    switch fi61 {
                                    case 127:
                                        fmt.Println(fi60, v3)
                                        v64 := GenS9{F0: 976}
                                        _ = v64
                                    default:
                                        v65 := GenS8{F0: true, F1: 1}
                                        _ = v65
                                    }
                                case "a":
                                    var v66 uint8 = 0
                                    _ = v66
                                }
                                v3 -= (v3 & (v3 * v3))
                            }
                            _ = len(v1)
                            _ = v62
                        } else {
                            v67 := GenS5{F0: 580, F1: 905}
                            _ = v67
                        }
                        var v68 int64 = 0
                        _ = len(v4)
                        _ = v68
                    }
                    _ = fi61
                }
                v69 := GenS4{F0: 584, F1: 330, F2: 937}
                _ = fi60
                _ = v69
            }
            v70 := GenS2{F0: true, F1: 2}
            _ = v70
        }
    }
    if true {
        var v71 int32 = v3
        v71 -= 332
        fmt.Println(v71, v3)
        _ = v71
    } else {
        var v72 int32 = v3
        if false {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            if (0 > 32767) {
                for fi73 := 0; fi73 < 4; fi73++ {
                    var v74 uint = 2
                    _ = fi73
                    _ = v74
                }
                v75 := v4[-2147483648]
                _ = v75
            }
            if true {
                fmt.Println(v3, v3)
            }
        } else {
            if ((65535 == 729) && true) {
                var v76 int32 = v3
                if false {
                    var v77 int32 = v3
                    _ = v77
                }
                v72 += 42
                _ = v76
            } else {
                switch 32768 {
                case -2147483648:
                    var v78 bool = (true || true)
                    _ = v78
                case 721:
                    v3 += v3
                case 127:
                    switch 42 {
                    case -1:
                        var v79 int32 = 0
                        _ = v79
                    default:
                        v80 := GenS6{F0: ((("\n\t" + (("hello" + (("a" + (("foo bar" + ("foo bar" + "\n\t")) + ("a" + "hello"))) + "a")) + "foo bar")) + ("a" + (((("foo bar" + ((("" + "foo bar") + "") + ("hello" + "foo bar"))) + "a") + ("hello" + ("" + ("a" + "")))) + (("foo bar" + "") + "")))) + ("a" + "foo bar")), F1: 1, F2: true}
                        _ = v80
                    }
                default:
                    fmt.Println(v72)
                }
                if (!true) {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v81 uint8 = 127
                    var v82 float32 = (0.1 - ((0.1 - float32(-2.25)) + (float32(1.5) - 3.14159)))
                    _ = v81
                    _ = v82
                }
                switch 32768 {
                case 127:
                    if (9223372036854775807 == -1) {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    } else {
                        var v83 float32 = 3.14159
                        v83 = 0.1
                        for fi84 := 0; fi84 < 5; fi84++ {
                            var v85 uint64 = 18446744073709551615
                            fmt.Println(v85)
                            var v86 int64 = -1
                            _ = fi84
                            _ = v85
                            _ = v86
                        }
                        _ = v83
                    }
                case 690:
                    for fi87 := 0; fi87 < 16; fi87++ {
                        fmt.Println(v72, fi87, v3)
                        if true {
                            v72 += v72
                            switch "hello" {
                            case "hello":
                                _ = len(v5)
                            case "a":
                                var v88 uint8 = 66
                                _ = v88
                            default:
                                v3 = v72
                            }
                        } else {
                            switch fi87 {
                            case 616:
                                for fi89 := 0; fi89 < 17; fi89++ {
                                    if false {
                                        _ = len(v2)
                                        if len(v2) > 0 {
                                            v2[0] = fi87
                                        }
                                        var v90 int8 = -13
                                        _ = v90
                                    }
                                    v3 += (v3 & 2147483647)
                                    _ = fi89
                                }
                            case -897:
                                v72 += -842
                            default:
                                v91 := GenS2{F0: true, F1: 251}
                                _ = v91
                            }
                        }
                        _ = fi87
                    }
                case 196:
                    var v92 uint32 = 2
                    _ = v92
                }
            }
            fmt.Println(v72, v3)
        }
        _ = v72
    }
    var v93 float32 = (float32(((1e10 * 3.14159) - 3.14159)) - -2.25)
    v93 = (3.14159 + v93)
    for fi94 := 0; fi94 < 7; fi94++ {
        var v95 uint = 255
        for fi96 := 0; fi96 < 12; fi96++ {
            if true {
                var v97 uint32 = 2
                v93 = v93
                if false {
                    if (v3 > v3) {
                        var v98 string = ""
                        if false {
                            v95 += (^(v95))
                        }
                        _ = v98
                    } else {
                        if true {
                            v2 = append(v2, -1)
                            v3 += v3
                        }
                        v93 -= ((((((v93 / 1.0) * (1e10 * float32((0.0 / (1e10 + float64(v93)))))) + v93) / v93) - float32(float64(1e10))) * 1.5)
                        v3 = -1
                    }
                    fmt.Println(fi94, v93, v95)
                    v99 := v4[(fi96 >> (((v95 % (((v95 + (v95 >> ((0) & 63)))) | 1))) & 63))]
                    _ = v99
                } else {
                    var v100 uint = (v95 & 130)
                    _ = v100
                }
                _ = v97
            } else {
                switch "" {
                case "a":
                    _ = len(v2)
                default:
                    v95 = (^(v95))
                }
                _ = len(v1)
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            _ = fi96
        }
        var v101 int32 = v3
        _ = fi94
        _ = v101
        _ = v95
    }
    if (!false) {
        fmt.Println(v93, v3)
        delete(v4, 757)
        v93 -= (v93 / 1.0)
    } else {
        if true {
            if (3.14159 == v93) {
                v93 -= v93
                var v102 uint = 42
                for fi103 := 0; fi103 < 5; fi103++ {
                    v102 += (^(v102))
                    var v104 bool = (fi103 == 9223372036854775807)
                    v105 := v5[fi103]
                    _ = fi103
                    _ = v104
                    _ = v105
                }
                _ = v102
            } else {
                for fi106 := 0; fi106 < 19; fi106++ {
                    if (1 > -67) {
                        v93 = ((v93 - 3.14159) + float32(((float64(1.5) * -2.25) / 1.5)))
                        if true {
                            v3 += v3
                        }
                    }
                    _ = fi106
                }
                v107 := GenS10{F0: 0, F1: 2147483647, F2: 65535}
                if (9223372036854775807 >= 42) {
                    v108 := GenS11{F0: 496}
                    var v109 uint16 = 65535
                    _ = v108
                    _ = v109
                }
                _ = v107
            }
        }
    }
    for fi110 := 0; fi110 < 11; fi110++ {
        var v111 uint8 = 0
        _ = fi110
        _ = v111
    }
    for fi112 := 0; fi112 < 1; fi112++ {
        var v113 int = 32768
        var v114 int16 = 42
        v115 := GenS7{F0: 954}
        _ = fi112
        _ = v113
        _ = v114
        _ = v115
    }
    if true {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    if (-120 > -128) {
        fmt.Println(v3)
        var v116 string = "hello"
        _ = v116
    }
    fmt.Println(v3, v3)
    v3 -= v3
    _ = len(v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v93:", v93)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v93
}
