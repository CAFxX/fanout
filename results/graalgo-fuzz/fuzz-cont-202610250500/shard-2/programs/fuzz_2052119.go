package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 uint32; F1 int }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint64; F1 uint32; F2 uint32 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int; F1 uint32; F2 uint8 }
type GenS9 struct { F0 bool; F1 string; F2 bool }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 uint32; F1 int8 }
type GenS12 struct { F0 int64; F1 uint16 }
type GenS13 struct { F0 string; F1 uint64; F2 int64 }
type GenS14 struct { F0 uint; F1 int }
type GenS15 struct { F0 string; F1 int }
func genFn0(p0 bool, p1 string, p2 int32) uint {
    if p0 {
        v6 := GenS1{F0: "hello"}
        _ = v6
    }
    v7 := GenS1{F0: (p1 + (("\n\t" + p1) + (p1 + (p1 + "foo bar"))))}
    if false {
        switch 128 {
        case -9223372036854775808:
            if (32767 > -32768) {
                if (127 <= 127) {
                    if (2 != 4294967295) {
                        fmt.Println(p1, p2)
                    }
                    switch p1 {
                    case "":
                        if (1 > 141) {
                            for fi8 := 0; fi8 < 19; fi8++ {
                                for fi9 := 0; fi9 < 13; fi9++ {
                                    v10 := GenS1{F0: (p1 + p1)}
                                    v11 := GenS1{F0: ("hello" + p1)}
                                    v12 := []int{-2147483648, (fi8 ^ fi8), (fi9 * (fi8 & (fi8 ^ fi9))), (fi8 | (fi8 & (fi9 >> ((547) & 63))))}
                                    _ = fi9
                                    _ = v10
                                    _ = v11
                                    _ = v12
                                }
                                var v13 uint64 = 42
                                v14 := GenS1{F0: p1}
                                _ = fi8
                                _ = v13
                                _ = v14
                            }
                        }
                    case "foo bar":
                        p2 += 927
                    default:
                        p1 = (p1 + "")
                    }
                    switch 2147483647 {
                    case 565:
                        fmt.Println(p2, p1, p2)
                    case 9223372036854775807:
                        if p0 {
                            fmt.Println(p0)
                            v15 := GenS1{F0: (p1 + p1)}
                            _ = v15
                        } else {
                            v16 := GenS1{F0: p1}
                            switch p1 {
                            case "hello":
                                var v17 int = 32768
                                _ = v17
                            case "":
                                var v18 uint16 = 620
                                _ = v18
                            }
                            _ = v16
                        }
                    }
                } else {
                    if false {
                        v19 := []int{128, 54, 42}
                        if p0 {
                            p0 = (65535 <= 2)
                        } else {
                            var v20 int32 = (p2 ^ p2)
                            _ = v20
                        }
                        fmt.Println(p2, p0)
                        _ = v19
                    } else {
                        switch 128 {
                        case 0:
                            var v21 int = 2147483647
                            _ = v21
                        case 127:
                            if p0 {
                                v22 := GenS1{F0: p1}
                                _ = v22
                            } else {
                                for fi23 := 0; fi23 < 4; fi23++ {
                                    v24 := []int{}
                                    _ = fi23
                                    _ = v24
                                }
                                for fi25 := 0; fi25 < 16; fi25++ {
                                    switch "\n\t" {
                                    case "foo bar":
                                        var v26 string = p1
                                        _ = v26
                                    case "":
                                        var v27 float64 = ((((-2.25 - -2.25) + (-2.25 * -2.25)) - float64(0.1)) - float64(1e10))
                                        _ = v27
                                    case "hello":
                                        for fi28 := 0; fi28 < 9; fi28++ {
                                            switch fi25 {
                                            case 32767:
                                                fmt.Println(fi28)
                                            case 32768:
                                                fmt.Println(p1, p0)
                                            case -9223372036854775808:
                                                var v29 uint64 = 42
                                                _ = v29
                                            default:
                                                fmt.Println(fi25, fi28, p0)
                                            }
                                            _ = fi28
                                        }
                                    }
                                    if (float32((((-2.25 - 3.14159) - 0.1) - (0.0 - (((1e10 * float64(1.5)) - (0.1 - 1.5)) * float64((1.5 - -2.25)))))) == 1.5) {
                                        if ((true && true) || (p0 || false)) {
                                            switch 127 {
                                            case 1:
                                                p0 = p0
                                            case 4294967295:
                                                if (!true) {
                                                    fmt.Println(p0)
                                                } else {
                                                    var v30 int = (fi25 ^ fi25)
                                                    v31 := []int{v30, 127, (-(fi25))}
                                                    p2 -= (p2 << ((255) & 63))
                                                    _ = v30
                                                    _ = v31
                                                }
                                            }
                                            v32 := []int{(^(fi25)), fi25, (^(fi25))}
                                            if len(v32) >= 2 {
                                                v32 = v32[1:len(v32)]
                                            }
                                            _ = v32
                                        }
                                        for fi33 := 0; fi33 < 17; fi33++ {
                                            fmt.Println(fi33, p0, p0)
                                            var v34 int = fi33
                                            _ = fi33
                                            _ = v34
                                        }
                                    } else {
                                        v35 := []int{32767}
                                        _ = v35
                                    }
                                    if (0 <= -121) {
                                        p0 = p0
                                        v36 := GenS2{F0: (p2 % ((p2) | 1)), F1: 2147483647}
                                        _ = v36
                                    }
                                    _ = fi25
                                }
                                for fi37 := 0; fi37 < 13; fi37++ {
                                    if false {
                                        for fi38 := 0; fi38 < 13; fi38++ {
                                            p1 = ((("" + (p1 + "foo bar")) + (p1 + ((p1 + p1) + p1))) + ((("a" + "foo bar") + (p1 + p1)) + (p1 + p1)))
                                            _ = fi38
                                        }
                                        v39 := GenS1{F0: ("hello" + p1)}
                                        _ = v39
                                    } else {
                                        v40 := GenS3{F0: 42, F1: 0}
                                        var v41 uint16 = 678
                                        var v42 int = fi37
                                        _ = v40
                                        _ = v41
                                        _ = v42
                                    }
                                    _ = fi37
                                }
                            }
                        case -2147483648:
                            for fi43 := 0; fi43 < 4; fi43++ {
                                switch fi43 {
                                case -682:
                                    for fi44 := 0; fi44 < 4; fi44++ {
                                        switch fi43 {
                                        case 32768:
                                            var v45 int16 = -1
                                            _ = v45
                                        default:
                                            v46 := GenS3{F0: 410, F1: (fi43 * fi44)}
                                            _ = v46
                                        }
                                        v47 := GenS1{F0: "hello"}
                                        _ = fi44
                                        _ = v47
                                    }
                                case 2147483647:
                                    var v48 int8 = 0
                                    _ = v48
                                case 9223372036854775807:
                                    v49 := GenS3{F0: 2147483648, F1: fi43}
                                    _ = v49
                                default:
                                    if p0 {
                                        for fi50 := 0; fi50 < 2; fi50++ {
                                            v51 := []int{}
                                            var v52 int = fi50
                                            switch (fi50 - (fi43 + (-(fi50)))) {
                                            case 672:
                                                var v53 uint8 = 128
                                                _ = v53
                                            case 247:
                                                fmt.Println(fi50, p1, p2)
                                            default:
                                                if len(v51) > 0 {
                                                    v51[0] = (v52 * v52)
                                                }
                                            }
                                            _ = fi50
                                            _ = v51
                                            _ = v52
                                        }
                                    }
                                }
                                _ = fi43
                            }
                        default:
                            var v54 float64 = 0.1
                            _ = v54
                        }
                        var v55 string = "foo bar"
                        _ = v55
                    }
                    v56 := []int{32767, 644}
                    _ = v56
                }
                if (127 < 118) {
                    for fi57 := 0; fi57 < 1; fi57++ {
                        if (p0 || p0) {
                            if p0 {
                                for fi58 := 0; fi58 < 3; fi58++ {
                                    var v59 int16 = 1
                                    v59 -= v59
                                    var v60 int8 = -128
                                    v59 -= 1000
                                    if p0 {
                                        v61 := GenS4{F0: 147}
                                        p1 = "hello"
                                        switch int(v59) {
                                        case 1:
                                            var v62 string = p1
                                            _ = v62
                                        default:
                                            p2 += -2147483648
                                        }
                                        _ = v61
                                    }
                                    _ = fi58
                                    _ = v59
                                    _ = v60
                                }
                                fmt.Println(p0, p2, p0)
                            }
                        } else {
                            v63 := GenS2{F0: int32(fi57), F1: 300}
                            var v64 int = fi57
                            var v65 string = p1
                            _ = v63
                            _ = v64
                            _ = v65
                        }
                        if (fi57 != fi57) {
                            p0 = false
                            v66 := GenS2{F0: (-(p2)), F1: 2147483647}
                            _ = v66
                        } else {
                            v67 := GenS4{F0: 4294967295}
                            if p0 {
                                var v68 float64 = 1.5
                                v69 := []int{0, (fi57 * fi57), fi57}
                                _ = v68
                                _ = v69
                            } else {
                                if false {
                                    var v70 uint32 = 1
                                    v71 := []int{4294967295, (fi57 * -2147483648), fi57, (fi57 & (fi57 + (fi57 - -197)))}
                                    v70 = (v70 >> ((18446744073709551615) & 63))
                                    _ = v70
                                    _ = v71
                                }
                            }
                            v72 := []int{}
                            _ = v67
                            _ = v72
                        }
                        p0 = false
                        _ = fi57
                    }
                    v73 := GenS2{F0: 1, F1: (p2 ^ 42)}
                    var v74 string = p1
                    _ = v73
                    _ = v74
                }
            }
        }
        var v75 float64 = (1.5 - (3.14159 - 3.14159))
        var v76 int64 = 983
        _ = v75
        _ = v76
    }
    var v77 uint = 4294967295
    fmt.Println(p0, p0, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v77
    return v77
}

func main() {
    var v1 int32 = 97
    v2 := []uint32{414, 1, 2}
    var v3 int64 = -9223372036854775808
    var v4 float64 = 0.1
    var v5 uint16 = 255
    v2 = append(v2, 476)
    fmt.Println(v3, v3)
    v78 := GenS5{F0: 42, F1: 4294967295, F2: 174}
    v79 := GenS6{F0: 1}
    var v80 uint16 = (^(v5))
    if (42 > 32768) {
        v81 := GenS5{F0: 0, F1: 833, F2: 834}
        _ = v81
    } else {
        if (v4 == (1.5 + -2.25)) {
            fmt.Println(v1, v4)
            for fi82 := 0; fi82 < 9; fi82++ {
                fmt.Println(v1)
                v1 = v1
                if true {
                    v3 += -9223372036854775808
                    v3 += (v3 >> ((255) & 63))
                }
                _ = fi82
            }
        } else {
            var v83 uint = 4294967295
            _ = v83
        }
        for fi84 := 0; fi84 < 4; fi84++ {
            v85 := GenS7{F0: 884}
            v86 := GenS2{F0: (v1 | v1), F1: 42}
            _ = fi84
            _ = v85
            _ = v86
        }
    }
    v87 := GenS6{F0: 1000}
    v3 = (v3 % (((v3 << ((0) & 63))) | 1))
    v88 := GenS8{F0: 2147483647, F1: 2147483648, F2: 104}
    var v89 uint = genFn0(false, "hello", (-(v1)))
    v90 := GenS9{F0: false, F1: (("foo bar" + "") + ""), F2: false}
    if false {
        fmt.Println(v3, v1)
        var v91 float64 = v4
        v92 := GenS10{F0: 550}
        _ = v91
        _ = v92
    }
    if false {
        v4 = (float64(3.14159) + float64(0.0))
        if (128 == 42) {
            v89 = (v89 | v89)
        } else {
            var v93 uint16 = v80
            var v94 float64 = (0.0 / v4)
            _ = v93
            _ = v94
        }
        for fi95 := 0; fi95 < 16; fi95++ {
            v3 -= (v3 + v3)
            _ = fi95
        }
    } else {
        var v96 uint = v89
        var v97 uint32 = 36
        _ = v96
        _ = v97
    }
    var v98 uint32 = 369
    if false {
        v99 := GenS5{F0: 112, F1: v98, F2: v98}
        var v100 int16 = 42
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v100
        _ = v99
    }
    if true {
        if (!false) {
            switch ("" + "a") {
            case "foo bar":
                for fi101 := 0; fi101 < 8; fi101++ {
                    v89 = (v89 * genFn0((-357 > v1), "", v1))
                    if false {
                        v2 = append(v2, (^(v98)))
                        v89 -= 581
                    }
                    v102 := GenS4{F0: v98}
                    _ = fi101
                    _ = v102
                }
            case "hello":
                v2 = append(v2, v98)
            default:
                v103 := GenS6{F0: 1000}
                _ = v103
            }
            var v104 uint8 = 128
            _ = v104
        } else {
            fmt.Println(v1)
            v105 := GenS11{F0: v98, F1: -128}
            _ = v105
        }
    } else {
        if true {
            v106 := GenS12{F0: (-(v3)), F1: (^(v80))}
            if false {
                if (742 >= 1) {
                    var v107 float32 = -2.25
                    var v108 bool = false
                    switch 9223372036854775807 {
                    case 42:
                        for fi109 := 0; fi109 < 4; fi109++ {
                            for fi110 := 0; fi110 < 12; fi110++ {
                                if (("" + ("foo bar" + "a")) <= ("foo bar" + (("" + "hello") + ("" + "\n\t")))) {
                                    v80 = (^(v80))
                                    v111 := GenS11{F0: v98, F1: 42}
                                    _ = v111
                                } else {
                                    v112 := GenS2{F0: v1, F1: v1}
                                    v98 -= (v98 | v98)
                                    var v113 float64 = 0.0
                                    _ = v112
                                    _ = v113
                                }
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                v2 = append(v2, v98)
                                _ = fi110
                            }
                            for fi114 := 0; fi114 < 11; fi114++ {
                                var v115 string = (("foo bar" + "foo bar") + (("hello" + "\n\t") + "hello"))
                                v98 += (v98 / ((4294967295) | 1))
                                _ = fi114
                                _ = v115
                            }
                            v2 = append(v2, (v98 / ((v98) | 1)))
                            _ = fi109
                        }
                    case -15:
                        for fi116 := 0; fi116 < 18; fi116++ {
                            if false {
                                if v108 {
                                    v117 := GenS13{F0: "hello", F1: 2, F2: (v3 ^ (v3 % ((v3) | 1)))}
                                    var v118 float64 = 3.14159
                                    fmt.Println(v1)
                                    _ = v117
                                    _ = v118
                                } else {
                                    if (!(563 == fi116)) {
                                        v80 -= v80
                                        switch "\n\t" {
                                        case "\n\t":
                                            for fi119 := 0; fi119 < 7; fi119++ {
                                                var v120 int32 = v1
                                                fmt.Println(v98, v89)
                                                v3 -= (v3 | (v3 ^ int64(fi119)))
                                                _ = fi119
                                                _ = v120
                                            }
                                        }
                                        var v121 int16 = 0
                                        _ = v121
                                    }
                                    v3 -= (v3 / ((-933) | 1))
                                    for fi122 := 0; fi122 < 16; fi122++ {
                                        v123 := GenS11{F0: (v98 % (((v98 >> ((genFn0((true || true), "foo bar", (v1 % ((0) | 1)))) & 63))) | 1)), F1: int8(v89)}
                                        _ = fi122
                                        _ = v123
                                    }
                                }
                                v124 := GenS8{F0: fi116, F1: 95, F2: 0}
                                v4 = (v4 + (0.1 - (v4 / v4)))
                                _ = v124
                            } else {
                                var v125 bool = v108
                                _ = v125
                            }
                            _ = fi116
                        }
                    case -9223372036854775808:
                        if len(v2) > 0 {
                            v2[0] = (^(v98))
                        }
                    }
                    _ = v107
                    _ = v108
                } else {
                    switch 4294967295 {
                    case -1:
                        switch 982 {
                        case 127:
                            v126 := GenS14{F0: genFn0((false || false), (((("a" + ("a" + ("a" + ("" + (("hello" + (("foo bar" + "a") + "hello")) + ("a" + "hello")))))) + ("\n\t" + (("a" + "\n\t") + ("\n\t" + "\n\t")))) + ("\n\t" + ("hello" + (("foo bar" + "hello") + "\n\t")))) + ""), 1), F1: 4294967295}
                            _ = v126
                        default:
                            var v127 int32 = (v1 & -2147483648)
                            _ = v127
                        }
                    case 1:
                        v128 := GenS13{F0: "", F1: 0, F2: -9223372036854775808}
                        _ = v128
                    case 32767:
                        fmt.Println(v98)
                    default:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v129 uint16 = 65535
                    _ = v129
                }
                fmt.Println(v3, v5, v1)
                v130 := GenS6{F0: int16(v3)}
                _ = v130
            } else {
                fmt.Println(v1, v5)
                if (672 <= 91) {
                    if (true || false) {
                        _ = len(v2)
                    } else {
                        if true {
                            var v131 uint16 = v5
                            fmt.Println(v4)
                            _ = v131
                        }
                    }
                    v132 := GenS8{F0: -956, F1: 1, F2: 127}
                    if (true && true) {
                        v133 := GenS15{F0: "hello", F1: -2147483648}
                        var v134 bool = (!true)
                        if (v134 && v134) {
                            _ = len(v2)
                        } else {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            fmt.Println(v134, v80, v3)
                            var v135 int64 = v3
                            _ = v135
                        }
                        _ = v133
                        _ = v134
                    }
                    _ = v132
                } else {
                    var v136 int16 = 0
                    v98 -= 86
                    if ("\n\t" < ("hello" + "foo bar")) {
                        var v137 int8 = 100
                        switch "a" {
                        case "foo bar":
                            fmt.Println(v137)
                        case "":
                            if len(v2) > 0 {
                                v2[0] = (v98 % (((v98 >> ((genFn0(((float32(v4) != ((((((0.0 + 0.1) + 3.14159) + (0.0 / float32(float64(float32(1e10))))) + 1.5) + 0.0) * (-2.25 + 3.14159))) && false), "foo bar", (v1 * 0))) & 63))) | 1))
                            }
                        }
                        _ = v137
                    }
                    _ = v136
                }
            }
            _ = v106
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        var v138 int64 = (v3 % ((v3) | 1))
        _ = v138
    }
    var v139 int16 = 819
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v139:", v139)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v80:", v80)
    fmt.Println("v89:", v89)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v139
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v78
    _ = v79
    _ = v80
    _ = v87
    _ = v88
    _ = v89
    _ = v90
    _ = v98
}
