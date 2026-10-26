package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 uint8 }
type GenS2 struct { F0 uint8; F1 bool; F2 int8 }
type GenS3 struct { F0 int16; F1 uint; F2 string }
type GenS4 struct { F0 uint8; F1 int64 }
type GenS5 struct { F0 int8; F1 int16 }
type GenS6 struct { F0 bool; F1 bool }
type GenS7 struct { F0 string; F1 int32 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 int8; F1 int32 }
type GenS10 struct { F0 uint; F1 string; F2 int32 }
type GenS11 struct { F0 uint64; F1 string }
type GenS12 struct { F0 uint8; F1 int64 }
type GenS13 struct { F0 int8; F1 int32 }
func genFn0(p0 int64, p1 string, p2 uint) int64 {
    v7 := []int{}
    for fi8 := 0; fi8 < 2; fi8++ {
        var v9 int = (fi8 << (((p2 % ((p2) | 1))) & 63))
        if (true || (730 <= 0)) {
            switch p1 {
            case "a":
                v7 = append(v7, 42)
            case "foo bar":
                if false {
                    if true {
                        if true {
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                            if false {
                                v10 := GenS1{F0: -2147483648, F1: 1, F2: 255}
                                _ = v10
                            } else {
                                if (p0 != p0) {
                                    v11 := GenS2{F0: 50, F1: ((2147483647 == -739) || true), F2: 42}
                                    _ = v11
                                }
                                v7 = append(v7, 32767)
                                v7 = append(v7, (-(fi8)))
                            }
                        } else {
                            for fi12 := 0; fi12 < 9; fi12++ {
                                switch (^(fi12)) {
                                case -234:
                                    if ((0.0 <= 1e10) && ((!(915 >= -32768)) && false)) {
                                        var v13 uint = 988
                                        _ = v13
                                    }
                                case -937:
                                    var v14 uint64 = 684
                                    _ = v14
                                case 2147483647:
                                    if false {
                                        var v15 int32 = int32(p2)
                                        _ = v15
                                    } else {
                                        fmt.Println(v9, fi8, fi12)
                                        _ = len(v7)
                                        fmt.Println(fi8, p0, v9)
                                    }
                                }
                                var v16 bool = (57 <= 1000)
                                _ = fi12
                                _ = v16
                            }
                            fmt.Println(v9)
                            for fi17 := 0; fi17 < 14; fi17++ {
                                var v18 string = p1
                                v19 := GenS3{F0: 1, F1: 2, F2: ("a" + (((("" + "foo bar") + "\n\t") + (p1 + "hello")) + (((v18 + ("\n\t" + "a")) + ((p1 + (v18 + ("foo bar" + p1))) + ((p1 + (p1 + "")) + "\n\t"))) + v18)))}
                                var v20 int8 = 42
                                _ = fi17
                                _ = v18
                                _ = v19
                                _ = v20
                            }
                        }
                        switch "" {
                        case "":
                            for fi21 := 0; fi21 < 2; fi21++ {
                                var v22 bool = (1 <= 42)
                                _ = fi21
                                _ = v22
                            }
                        case "foo bar":
                            var v23 string = ""
                            _ = v23
                        default:
                            fmt.Println(p0, v9)
                        }
                        for fi24 := 0; fi24 < 8; fi24++ {
                            fmt.Println(p0, p2, p2)
                            v25 := GenS4{F0: 255, F1: (p0 - p0)}
                            v26 := GenS3{F0: -32768, F1: (p2 << ((p2) & 63)), F2: p1}
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                    } else {
                        if (!true) {
                            var v27 uint64 = 18446744073709551615
                            if true {
                                p1 = p1
                            } else {
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                            }
                            p2 = (^(p2))
                            _ = v27
                        }
                    }
                } else {
                    var v28 float32 = 3.14159
                    _ = v28
                }
            case "":
                for fi29 := 0; fi29 < 6; fi29++ {
                    var v30 bool = true
                    fmt.Println(v30, p1, v30)
                    _ = fi29
                    _ = v30
                }
            default:
                switch (v9 << ((931) & 63)) {
                case 42:
                    if true {
                        v9 = -97
                    }
                case 4294967295:
                    for fi31 := 0; fi31 < 2; fi31++ {
                        switch (^(fi31)) {
                        case 2147483647:
                            var v32 int8 = 1
                            _ = v32
                        case -109:
                            v33 := GenS5{F0: -10, F1: -211}
                            _ = v33
                        case -638:
                            v9 += (fi8 / ((int(p2)) | 1))
                        }
                        v7 = append(v7, (v9 + -645))
                        _ = fi31
                    }
                }
            }
        }
        _ = fi8
        _ = v9
    }
    fmt.Println(p2, p1, p1)
    _ = len(v7)
    var v34 string = "hello"
    _ = p0
    _ = p1
    _ = p2
    _ = v34
    _ = v7
    return p0
}

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[string]int)
    var v3 int8 = 74
    var v4 uint64 = 42
    v5 := make(map[string]int)
    var v6 string = "foo bar"
    var v35 uint32 = 156
    for fi36 := 0; fi36 < 20; fi36++ {
        v37 := GenS5{F0: v3, F1: 32767}
        _ = fi36
        _ = v37
    }
    if ((^(v35)) != v35) {
        v2[v6] = -798
        var v38 uint = 903
        var v39 uint32 = v35
        _ = v38
        _ = v39
    }
    v40 := []int{}
    v41 := GenS2{F0: 154, F1: false, F2: v3}
    for fi42 := 0; fi42 < 19; fi42++ {
        for fi43 := 0; fi43 < 8; fi43++ {
            v44 := v5[v6]
            _ = fi43
            _ = v44
        }
        _ = fi42
    }
    switch "" {
    case "foo bar":
        if (true || ("" < "\n\t")) {
            if len(v40) > 0 {
                v40[0] = 1
            }
            if (255 != 42) {
                for fi45 := 0; fi45 < 3; fi45++ {
                    v35 = v35
                    switch v6 {
                    case "":
                        if ((-32768 < 32767) || false) {
                            var v46 float64 = (float64(v1) / 1.0)
                            v35 = (^(v35))
                            v3 += 84
                            _ = v46
                        }
                    case "hello":
                        v47 := GenS2{F0: 255, F1: false, F2: int8(v35)}
                        _ = v47
                    default:
                        delete(v2, (v6 + "a"))
                    }
                    v48 := GenS4{F0: 0, F1: 42}
                    _ = fi45
                    _ = v48
                }
                fmt.Println(v6, v6, v1)
                var v49 int = 32768
                _ = v49
            }
        } else {
            for fi50 := 0; fi50 < 5; fi50++ {
                v6 = v6
                fmt.Println(v4, v6, fi50)
                v51 := GenS5{F0: (v3 * v3), F1: -172}
                _ = fi50
                _ = v51
            }
            v52 := GenS2{F0: 1, F1: false, F2: int8(v35)}
            if true {
                var v53 uint16 = 165
                switch (v6 + (v6 + v6)) {
                case "hello":
                    for fi54 := 0; fi54 < 12; fi54++ {
                        v3 = 100
                        v55 := v2[("foo bar" + v6)]
                        if (!(v4 > v4)) {
                            if (int16(v55) == int16(fi54)) {
                                v56 := GenS3{F0: int16(fi54), F1: 615, F2: v6}
                                _ = v56
                            } else {
                                if (1000 < 42) {
                                    fmt.Println(v53)
                                } else {
                                    v57 := GenS1{F0: -2147483648, F1: -367, F2: 42}
                                    for fi58 := 0; fi58 < 12; fi58++ {
                                        switch -9223372036854775808 {
                                        case -171:
                                            var v59 float32 = v1
                                            v59 = v1
                                            _ = v59
                                        case -115:
                                            if true {
                                                var v60 bool = false
                                                if true {
                                                    if len(v40) >= 2 {
                                                        v40 = v40[1:len(v40)]
                                                    }
                                                    var v61 int8 = 100
                                                    _ = v61
                                                } else {
                                                    _ = len(v40)
                                                }
                                                var v62 int64 = 1
                                                _ = v60
                                                _ = v62
                                            } else {
                                                delete(v5, v6)
                                            }
                                        case -2147483648:
                                            switch v6 {
                                            case "\n\t":
                                                var v63 string = ("hello" + "hello")
                                                _ = v63
                                            default:
                                                if false {
                                                    var v64 int64 = -1
                                                    if (v3 < v3) {
                                                        if true {
                                                            var v65 uint8 = 255
                                                            v66 := GenS6{F0: (((v1 * float32(3.14159)) / (0.1 - 1.5)) == (v1 - v1)), F1: (v64 == (v64 ^ v64))}
                                                            var v67 uint = 0
                                                            _ = v65
                                                            _ = v66
                                                            _ = v67
                                                        }
                                                    }
                                                    var v68 uint32 = (^(v35))
                                                    fmt.Println(v4)
                                                    _ = v64
                                                    _ = v68
                                                } else {
                                                    v1 = (v1 - v1)
                                                    if (-2147483648 < 83) {
                                                        v69 := GenS7{F0: v6, F1: 535}
                                                        if (!(0.1 < 1.5)) {
                                                            fmt.Println(fi54)
                                                            v2["hello"] = -771
                                                        }
                                                        if (411 == 255) {
                                                            if (false && (1 != -362)) {
                                                                switch fi54 {
                                                                case -938:
                                                                    fmt.Println(v55)
                                                                default:
                                                                    v5[v6] = v55
                                                                }
                                                                if ((-2.25 / v1) != (float32(1e10) * v1)) {
                                                                    fmt.Println(v4)
                                                                    v70 := v2[(v6 + ("\n\t" + "\n\t"))]
                                                                    switch v55 {
                                                                    case 42:
                                                                        if false {
                                                                            switch (v6 + "\n\t") {
                                                                            case "foo bar":
                                                                                var v71 int8 = v3
                                                                                _ = v71
                                                                            }
                                                                            var v72 int64 = -1
                                                                            v1 -= (v1 + (v1 * (v1 * v1)))
                                                                            _ = v72
                                                                        } else {
                                                                            v3 += 47
                                                                        }
                                                                    case -136:
                                                                        var v73 string = v6
                                                                        _ = v73
                                                                    case -901:
                                                                        if true {
                                                                            switch v6 {
                                                                            case "\n\t":
                                                                                if (!(1 < 0)) {
                                                                                    if (-66 >= genFn0(9223372036854775807, (v6 + ""), 559)) {
                                                                                        var v74 uint = 832
                                                                                        _ = v74
                                                                                    } else {
                                                                                        var v75 int = (v70 | (fi54 + (v55 * fi58)))
                                                                                        var v76 uint8 = 0
                                                                                        _ = v75
                                                                                        _ = v76
                                                                                    }
                                                                                }
                                                                            }
                                                                            if len(v40) > 0 {
                                                                                v40[0] = (v55 * (v55 - (v70 << ((0) & 63))))
                                                                            }
                                                                        } else {
                                                                            v77 := GenS8{F0: v3}
                                                                            var v78 uint8 = 127
                                                                            _ = v77
                                                                            _ = v78
                                                                        }
                                                                    default:
                                                                        v70 += v55
                                                                    }
                                                                    _ = v70
                                                                } else {
                                                                    if len(v40) >= 2 {
                                                                        v40 = v40[1:len(v40)]
                                                                    }
                                                                }
                                                                if false {
                                                                    v6 = "hello"
                                                                }
                                                            } else {
                                                                var v79 uint8 = 127
                                                                v53 = v53
                                                                if (((v1 + v1) / v1) != -2.25) {
                                                                    var v80 int16 = -32768
                                                                    switch (v55 % (((v55 + -1)) | 1)) {
                                                                    case -477:
                                                                        v81 := GenS3{F0: v80, F1: 638, F2: (v6 + (v6 + v6))}
                                                                        _ = v81
                                                                    case 0:
                                                                        var v82 bool = false
                                                                        _ = v82
                                                                    case 101:
                                                                        var v83 uint8 = (v79 / (((^(v79))) | 1))
                                                                        _ = v83
                                                                    }
                                                                    _ = v80
                                                                } else {
                                                                    _ = len(v2)
                                                                    if (658 <= -9223372036854775808) {
                                                                        if (1 > 42) {
                                                                            var v84 float32 = (v1 - 3.14159)
                                                                            fmt.Println(fi58, v79, v4)
                                                                            _ = v84
                                                                        } else {
                                                                            v55 -= -2147483648
                                                                        }
                                                                        switch ((v6 + "") + "a") {
                                                                        case "a":
                                                                            var v85 bool = (-9223372036854775808 < int64(v4))
                                                                            _ = v85
                                                                        case "foo bar":
                                                                            v86 := GenS2{F0: v79, F1: (4294967295 != 37), F2: (v3 * (v3 / (((^(v3))) | 1)))}
                                                                            _ = v86
                                                                        }
                                                                        v87 := GenS3{F0: 42, F1: 255, F2: "foo bar"}
                                                                        _ = v87
                                                                    } else {
                                                                        v79 += v79
                                                                    }
                                                                }
                                                                _ = v79
                                                            }
                                                        }
                                                        var v88 uint8 = 42
                                                        _ = v69
                                                        _ = v88
                                                    } else {
                                                        var v89 uint32 = v35
                                                        var v90 string = v6
                                                        _ = v89
                                                        _ = v90
                                                    }
                                                }
                                            }
                                        }
                                        _ = fi58
                                    }
                                    fmt.Println(v53, v35, v6)
                                    _ = v57
                                }
                            }
                            v91 := v2[v6]
                            v92 := GenS4{F0: 255, F1: 0}
                            _ = v91
                            _ = v92
                        } else {
                            var v93 uint = 4294967295
                            _ = v93
                        }
                        _ = fi54
                        _ = v55
                    }
                default:
                    for fi94 := 0; fi94 < 18; fi94++ {
                        for fi95 := 0; fi95 < 17; fi95++ {
                            v40 = append(v40, (fi94 / (((fi94 >> ((387) & 63))) | 1)))
                            v40 = append(v40, fi94)
                            var v96 float32 = (v1 - ((v1 / -2.25) / v1))
                            _ = fi95
                            _ = v96
                        }
                        _ = fi94
                    }
                }
                _ = v53
            } else {
                fmt.Println(v6)
                switch (v6 + (v6 + v6)) {
                case "foo bar":
                    var v97 float64 = 0.1
                    _ = v97
                case "":
                    fmt.Println(v1, v6, v1)
                }
                v40 = append(v40, -1)
            }
            _ = v52
        }
    case "hello":
        var v98 uint16 = 32767
        _ = v98
    case "\n\t":
        fmt.Println(v4, v1, v35)
    default:
        var v99 int = 32768
        _ = v99
    }
    switch ((v6 + v6) + (v6 + v6)) {
    case "foo bar":
        var v100 uint32 = uint32(v4)
        _ = v100
    case "\n\t":
        var v101 uint32 = v35
        _ = v101
    case "a":
        var v102 uint8 = 0
        _ = v102
    }
    for fi103 := 0; fi103 < 12; fi103++ {
        v6 = ("hello" + ((v6 + v6) + v6))
        if len(v40) >= 2 {
            v40 = v40[1:len(v40)]
        }
        _ = fi103
    }
    if len(v40) >= 2 {
        v40 = v40[1:len(v40)]
    }
    for fi104 := 0; fi104 < 4; fi104++ {
        var v105 int16 = 42
        v106 := GenS1{F0: 1, F1: 0, F2: 169}
        fmt.Println(fi104, fi104, v6)
        _ = fi104
        _ = v105
        _ = v106
    }
    for fi107 := 0; fi107 < 6; fi107++ {
        v35 += (v35 >> ((4294967295) & 63))
        if (55 < 145) {
            if ("hello" != "\n\t") {
                for fi108 := 0; fi108 < 11; fi108++ {
                    _ = len(v40)
                    v109 := GenS2{F0: 128, F1: false, F2: 106}
                    if true {
                        v110 := GenS9{F0: v3, F1: 0}
                        if true {
                            switch (((((v6 + "") + (v6 + v6)) + "foo bar") + ("\n\t" + v6)) + v6) {
                            case "foo bar":
                                v3 = 0
                            case "a":
                                fmt.Println(v3, v35, fi107)
                                var v111 float64 = 1.5
                                _ = v111
                            case "hello":
                                if (34 < 255) {
                                    if (true || false) {
                                        var v112 uint = 336
                                        if (-1 == -2147483648) {
                                            if (true && true) {
                                                v6 = (v6 + "foo bar")
                                                v113 := GenS10{F0: v112, F1: v6, F2: 42}
                                                switch fi107 {
                                                case 567:
                                                    if false {
                                                        v3 -= -1
                                                        var v114 int32 = 0
                                                        _ = v114
                                                    } else {
                                                        switch -1 {
                                                        case 711:
                                                            v6 = "foo bar"
                                                        default:
                                                            var v115 int64 = -310
                                                            _ = v115
                                                        }
                                                        switch -1 {
                                                        case -2147483648:
                                                            var v116 int64 = genFn0(-757, v6, v112)
                                                            _ = v116
                                                        case -703:
                                                            var v117 uint64 = (^(v4))
                                                            _ = v117
                                                        }
                                                        v118 := GenS9{F0: (v3 / ((100) | 1)), F1: 415}
                                                        _ = v118
                                                    }
                                                    var v119 uint16 = 255
                                                    _ = v119
                                                case 42:
                                                    var v120 uint32 = (v35 >> ((v112) & 63))
                                                    _ = v120
                                                }
                                                _ = v113
                                            } else {
                                                var v121 uint32 = (v35 >> ((420) & 63))
                                                fmt.Println(v3, v6)
                                                if true {
                                                    v3 += (v3 + v3)
                                                    fmt.Println(v112, v121)
                                                    var v122 bool = true
                                                    var v123 uint16 = 65535
                                                    _ = v122
                                                    _ = v123
                                                } else {
                                                    v1 = ((v1 * 0.0) / float32((-2.25 * (-2.25 * float64(1.5)))))
                                                }
                                                _ = v121
                                            }
                                            switch (-(fi108)) {
                                            case 42:
                                                var v124 uint8 = 47
                                                _ = v124
                                            case 0:
                                                v4 = 0
                                            case 128:
                                                v125 := GenS11{F0: uint64(v35), F1: v6}
                                                _ = v125
                                            }
                                        }
                                        _ = len(v5)
                                        v1 = v1
                                        v3 += (^(v3))
                                        _ = v112
                                    }
                                    v126 := GenS12{F0: 0, F1: 9223372036854775807}
                                    var v127 uint8 = 128
                                    _ = v126
                                    _ = v127
                                } else {
                                    _ = len(v40)
                                    var v128 int16 = -1
                                    var v129 int16 = v128
                                    _ = v128
                                    _ = v129
                                }
                            }
                        }
                        if ((0 > 708) && ("foo bar" != v6)) {
                            delete(v2, v6)
                            var v130 int = (fi108 + (fi107 >> ((2) & 63)))
                            var v131 uint64 = v4
                            v131 -= v4
                            if true {
                                if (v1 > v1) {
                                    delete(v5, (v6 + "hello"))
                                }
                                var v132 int = fi108
                                if (1 > -669) {
                                    var v133 bool = (false && (false || (!((true || true) || true))))
                                    _ = v133
                                }
                                _ = v132
                            }
                            var v134 uint64 = v131
                            _ = v130
                            _ = v131
                            _ = v134
                        } else {
                            v135 := GenS13{F0: v3, F1: 0}
                            if (v4 <= 18446744073709551615) {
                                var v136 uint32 = 2147483648
                                _ = v136
                            } else {
                                fmt.Println(v1, v3)
                            }
                            _ = v135
                        }
                        _ = v110
                    } else {
                        var v137 int64 = 0
                        _ = len(v2)
                        var v138 uint = 4294967295
                        _ = v137
                        _ = v138
                    }
                    _ = fi108
                    _ = v109
                }
            }
        } else {
            fmt.Println(v1, v4)
            switch fi107 {
            case -119:
                for fi139 := 0; fi139 < 2; fi139++ {
                    switch 2147483647 {
                    case 4294967295:
                        v35 += 50
                    case -2147483648:
                        switch fi107 {
                        case 127:
                            var v140 uint = 1
                            _ = v140
                        default:
                            var v141 string = v6
                            _ = v141
                        }
                    case -675:
                        var v142 int8 = v3
                        _ = v142
                    }
                    var v143 int32 = 2147483647
                    _ = fi139
                    _ = v143
                }
            case 625:
                fmt.Println(v4)
            }
        }
        _ = fi107
    }
    var v144 string = ("\n\t" + ("foo bar" + (v6 + (("" + v6) + "foo bar"))))
    var v145 uint16 = 65535
    v4 += 362
    if ((v145 - v145) < v145) {
        var v146 uint = 1
        var v147 uint64 = v4
        _ = v146
        _ = v147
    } else {
        v4 = (v4 >> ((797) & 63))
    }
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v144:", v144)
    fmt.Println("v145:", v145)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v144
    _ = v145
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
