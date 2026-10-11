package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16; F1 int8; F2 string }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 string; F1 uint }
type GenS6 struct { F0 uint32; F1 int32; F2 int64 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int8; F1 int16 }
type GenS10 struct { F0 string; F1 uint16 }
type GenS11 struct { F0 int; F1 uint16; F2 uint32 }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 int64; F1 bool }
type GenS14 struct { F0 uint; F1 uint64; F2 string }
type GenS15 struct { F0 int }
type GenS16 struct { F0 int8; F1 int16; F2 uint8 }
type GenS17 struct { F0 string }
func genFn0(p0 uint16, p1 string) uint8 {
    var v4 float32 = 3.14159
    for fi5 := 0; fi5 < 16; fi5++ {
        var v6 uint64 = 2
        p0 += p0
        _ = fi5
        _ = v6
    }
    var v7 float64 = 0.0
    _ = p0
    _ = p1
    _ = v4
    _ = v7
    return 1
}
func genFn1(p0 int32) int {
    p0 += (p0 * 1)
    for fi8 := 0; fi8 < 9; fi8++ {
        for fi9 := 0; fi9 < 20; fi9++ {
            var v10 int8 = -128
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    fmt.Println(p0)
    if (true && true) {
        for fi11 := 0; fi11 < 1; fi11++ {
            if true {
                for fi12 := 0; fi12 < 11; fi12++ {
                    p0 -= p0
                    v13 := GenS1{F0: 42, F1: 2147483648, F2: true}
                    _ = fi12
                    _ = v13
                }
                for fi14 := 0; fi14 < 16; fi14++ {
                    if true {
                        p0 = 828
                        if (42 < 1) {
                            switch "" {
                            case "foo bar":
                                if (1e10 != float64(3.14159)) {
                                    var v15 float32 = 1e10
                                    _ = v15
                                } else {
                                    if false {
                                        var v16 uint64 = 85
                                        switch fi14 {
                                        case 127:
                                            if true {
                                                v16 = (v16 * 96)
                                                v17 := GenS2{F0: 42}
                                                v18 := GenS1{F0: 1, F1: 407, F2: (65535 <= uint(v16))}
                                                var v19 uint = 0
                                                _ = v17
                                                _ = v18
                                                _ = v19
                                            } else {
                                                fmt.Println(v16)
                                                v20 := GenS3{F0: -275, F1: -128, F2: "hello"}
                                                if (!true) {
                                                    switch -2147483648 {
                                                    case 4294967295:
                                                        v21 := []int{}
                                                        _ = v21
                                                    }
                                                } else {
                                                    if false {
                                                        var v22 float64 = 1.5
                                                        v23 := GenS2{F0: uint8(v16)}
                                                        if false {
                                                            var v24 int16 = 1000
                                                            v25 := []int{fi11}
                                                            if (255 > 18446744073709551615) {
                                                                fmt.Println(fi11, v16, fi14)
                                                            } else {
                                                                var v26 uint64 = 18446744073709551615
                                                                v25 = append(v25, 1)
                                                                p0 = p0
                                                                _ = v26
                                                            }
                                                            switch ("a" + "") {
                                                            case "foo bar":
                                                                var v27 bool = true
                                                                _ = v27
                                                            default:
                                                                v22 -= 0.1
                                                            }
                                                            _ = v24
                                                            _ = v25
                                                        } else {
                                                            p0 += 1
                                                        }
                                                        if false {
                                                            var v28 uint = 105
                                                            if false {
                                                                if ((("foo bar" + "\n\t") >= "foo bar") || false) {
                                                                    var v29 uint64 = v16
                                                                    _ = v29
                                                                } else {
                                                                    var v30 int16 = 32767
                                                                    v31 := GenS2{F0: 42}
                                                                    switch fi14 {
                                                                    case -9223372036854775808:
                                                                        var v32 uint = (v28 ^ 4294967295)
                                                                        _ = v32
                                                                    case 2147483647:
                                                                        p0 = (p0 + p0)
                                                                    }
                                                                    if (true && (1e10 == 1.5)) {
                                                                        v28 += 4294967295
                                                                        switch ("foo bar" + ("\n\t" + ("foo bar" + (("\n\t" + ("\n\t" + "a")) + (("foo bar" + "") + ("hello" + "a")))))) {
                                                                        case "a":
                                                                            fmt.Println(v16, v16)
                                                                            var v33 uint64 = v16
                                                                            _ = v33
                                                                        case "foo bar":
                                                                            v34 := GenS2{F0: 127}
                                                                            _ = v34
                                                                        case "hello":
                                                                            v35 := GenS1{F0: int64(v16), F1: 1, F2: false}
                                                                            _ = v35
                                                                        default:
                                                                            v36 := GenS4{F0: 100}
                                                                            _ = v36
                                                                        }
                                                                        var v37 uint8 = genFn0(65535, ("" + "\n\t"))
                                                                        _ = v37
                                                                    }
                                                                    _ = v30
                                                                    _ = v31
                                                                }
                                                                v38 := GenS3{F0: -131, F1: 42, F2: ("" + "hello")}
                                                                v39 := []int{fi14, (fi11 + fi14), (fi14 << ((1) & 63))}
                                                                if (((v22 / (0.0 - ((v22 / -2.25) / 0.1))) <= v22) && (!((false && false) && (false || false)))) {
                                                                    var v40 uint16 = 0
                                                                    v22 = 3.14159
                                                                    if len(v39) > 0 {
                                                                        v39[0] = fi14
                                                                    }
                                                                    _ = v40
                                                                }
                                                                _ = v38
                                                                _ = v39
                                                            } else {
                                                                if ((v16 != v16) && false) {
                                                                    if (32767 <= -848) {
                                                                        v41 := []int{(^(fi11)), fi11}
                                                                        var v42 bool = (((!false) && ((!(!(v28 >= 18446744073709551615))) || false)) || false)
                                                                        v28 = v28
                                                                        _ = v41
                                                                        _ = v42
                                                                    } else {
                                                                        v43 := GenS4{F0: -21}
                                                                        _ = v43
                                                                    }
                                                                    var v44 int8 = 127
                                                                    if (!((true || ((v28 / ((v28) | 1)) < (v28 & v28))) || true)) {
                                                                        if (42 >= (p0 << ((v28) & 63))) {
                                                                            switch (fi11 | (fi14 | fi14)) {
                                                                            case 4294967295:
                                                                                var v45 uint = v28
                                                                                _ = v45
                                                                            case 127:
                                                                                v22 = (v22 * 0.1)
                                                                            default:
                                                                                switch "\n\t" {
                                                                                case "foo bar":
                                                                                    var v46 int8 = (v44 << ((v28) & 63))
                                                                                    var v47 float64 = v22
                                                                                    _ = v46
                                                                                    _ = v47
                                                                                case "hello":
                                                                                    v22 -= v22
                                                                                case "":
                                                                                    v48 := GenS5{F0: ((((("hello" + "") + "\n\t") + "hello") + "\n\t") + "a"), F1: v28}
                                                                                    _ = v48
                                                                                default:
                                                                                    fmt.Println(v22, v16)
                                                                                }
                                                                            }
                                                                            if true {
                                                                                v49 := []int{fi11, (fi11 << ((v28) & 63)), (fi11 % (((fi14 / ((fi11) | 1))) | 1))}
                                                                                var v50 int32 = p0
                                                                                p0 -= p0
                                                                                _ = v49
                                                                                _ = v50
                                                                            }
                                                                        }
                                                                        v51 := GenS6{F0: 1, F1: p0, F2: -9223372036854775808}
                                                                        fmt.Println(v16, p0)
                                                                        _ = v51
                                                                    } else {
                                                                        v52 := []int{-9223372036854775808, fi14}
                                                                        _ = v52
                                                                    }
                                                                    _ = v44
                                                                }
                                                                v28 -= 18446744073709551615
                                                                v53 := []int{int(v28), fi14}
                                                                _ = v53
                                                            }
                                                            v22 = (v22 * v22)
                                                            _ = v28
                                                        }
                                                        v54 := GenS7{F0: 4294967295}
                                                        _ = v22
                                                        _ = v23
                                                        _ = v54
                                                    }
                                                    v55 := []int{-98, fi11, 4294967295}
                                                    p0 -= -1
                                                    var v56 uint32 = 65535
                                                    switch int(v56) {
                                                    case -254:
                                                        var v57 int64 = -336
                                                        _ = v57
                                                    case 336:
                                                        v56 -= 120
                                                    case 337:
                                                        if true {
                                                            if (801 != 1) {
                                                                v58 := GenS8{F0: "foo bar"}
                                                                if false {
                                                                    fmt.Println(fi11, fi14)
                                                                    if (42 == 764) {
                                                                        if (649 < 2) {
                                                                            v56 -= (^(v56))
                                                                            v59 := GenS3{F0: 942, F1: 0, F2: ""}
                                                                            _ = v59
                                                                        } else {
                                                                            p0 = p0
                                                                            if (42 != 32767) {
                                                                                if len(v55) >= 2 {
                                                                                    v55 = v55[1:len(v55)]
                                                                                }
                                                                            } else {
                                                                                if (true || (2 == (^(v16)))) {
                                                                                    v56 = 164
                                                                                    if (743 != 927) {
                                                                                        if len(v55) > 0 {
                                                                                            v55[0] = 128
                                                                                        }
                                                                                    }
                                                                                }
                                                                                if len(v55) > 0 {
                                                                                    v55[0] = 849
                                                                                }
                                                                                p0 += p0
                                                                                v60 := GenS5{F0: "hello", F1: 255}
                                                                                p0 -= 2147483647
                                                                                _ = v60
                                                                            }
                                                                        }
                                                                        v61 := GenS9{F0: int8(v56), F1: -32768}
                                                                        _ = v61
                                                                    } else {
                                                                        p0 = p0
                                                                    }
                                                                    p0 += p0
                                                                } else {
                                                                    if len(v55) >= 2 {
                                                                        v55 = v55[1:len(v55)]
                                                                    }
                                                                    if ((v16 << ((42) & 63)) != v16) {
                                                                        var v62 uint = 2
                                                                        v56 -= (v56 - v56)
                                                                        _ = v62
                                                                    }
                                                                }
                                                                _ = v58
                                                            } else {
                                                                v56 += 561
                                                                v16 = v16
                                                            }
                                                            v63 := GenS10{F0: ("foo bar" + ("\n\t" + "\n\t")), F1: 65535}
                                                            _ = v63
                                                        } else {
                                                            fmt.Println(v16, v16, fi14)
                                                        }
                                                    }
                                                    _ = v55
                                                    _ = v56
                                                }
                                                _ = v20
                                            }
                                        case -991:
                                            v64 := GenS6{F0: 147, F1: (p0 - p0), F2: -1}
                                            _ = v64
                                        default:
                                            v65 := GenS11{F0: fi11, F1: 32767, F2: 2147483647}
                                            _ = v65
                                        }
                                        _ = v16
                                    } else {
                                        switch int(p0) {
                                        case -2147483648:
                                            v66 := GenS11{F0: fi14, F1: 773, F2: 65535}
                                            _ = v66
                                        case 32767:
                                            v67 := GenS2{F0: genFn0(32767, "foo bar")}
                                            _ = v67
                                        case 2147483647:
                                            v68 := GenS5{F0: (("foo bar" + ("foo bar" + "\n\t")) + "a"), F1: 944}
                                            _ = v68
                                        default:
                                            v69 := GenS1{F0: 0, F1: 65535, F2: false}
                                            _ = v69
                                        }
                                        v70 := []int{fi14}
                                        fmt.Println(fi11, fi14, fi14)
                                        _ = v70
                                    }
                                }
                            case "\n\t":
                                p0 -= 434
                            case "a":
                                p0 -= (p0 ^ 2147483647)
                            default:
                                fmt.Println(p0, fi14)
                            }
                            switch ((("hello" + "") + "\n\t") + ("hello" + ("hello" + ("a" + "foo bar")))) {
                            case "hello":
                                v71 := []int{(fi11 >> ((101) & 63)), (fi11 ^ (fi11 / ((fi11) | 1))), fi11, (fi14 + 2147483647)}
                                _ = v71
                            case "":
                                var v72 int = 127
                                _ = v72
                            case "foo bar":
                                if (!false) {
                                    p0 -= p0
                                    v73 := GenS12{F0: 0}
                                    v74 := GenS12{F0: 18446744073709551615}
                                    fmt.Println(p0)
                                    fmt.Println(fi11, fi14, fi11)
                                    var v75 int32 = (p0 ^ -1)
                                    _ = v73
                                    _ = v74
                                    _ = v75
                                } else {
                                    if true {
                                        v76 := GenS7{F0: 18446744073709551615}
                                        var v77 int16 = -1
                                        var v78 int8 = 100
                                        _ = v76
                                        _ = v77
                                        _ = v78
                                    } else {
                                        p0 -= (p0 ^ p0)
                                        v79 := GenS13{F0: 537, F1: (!(!(((!(-128 != 0)) && true) && (p0 >= 0))))}
                                        _ = v79
                                    }
                                    var v80 int8 = 1
                                    v81 := GenS11{F0: (fi11 / ((fi14) | 1)), F1: 32767, F2: 65535}
                                    _ = v80
                                    _ = v81
                                }
                            }
                        } else {
                            if (!true) {
                                var v82 int16 = 1
                                v83 := GenS9{F0: 36, F1: 1}
                                _ = v82
                                _ = v83
                            }
                            v84 := []int{fi14, fi14, (fi11 / (((fi14 << ((0) & 63))) | 1))}
                            if (true || false) {
                                var v85 int16 = 780
                                _ = v85
                            } else {
                                var v86 uint64 = 18446744073709551615
                                var v87 uint64 = v86
                                var v88 int16 = 42
                                if false {
                                    v86 += (v87 << ((42) & 63))
                                    v86 = (v87 ^ v86)
                                    v89 := GenS14{F0: 255, F1: 4294967295, F2: ""}
                                    if len(v84) >= 2 {
                                        v84 = v84[1:len(v84)]
                                    }
                                    var v90 uint64 = v86
                                    var v91 int8 = 41
                                    _ = v89
                                    _ = v90
                                    _ = v91
                                }
                                if ((!(true && (!false))) && ((v88 ^ -32768) <= v88)) {
                                    v92 := GenS2{F0: 137}
                                    fmt.Println(v88, fi11)
                                    v87 = v87
                                    if (p0 > p0) {
                                        var v93 float64 = 1e10
                                        if len(v84) >= 2 {
                                            v84 = v84[1:len(v84)]
                                        }
                                        var v94 uint32 = 983
                                        _ = v93
                                        _ = v94
                                    }
                                    _ = v92
                                } else {
                                    if ((!true) || false) {
                                        _ = len(v84)
                                        var v95 int64 = 42
                                        _ = v95
                                    }
                                    var v96 uint = 4294967295
                                    _ = v96
                                }
                                _ = v86
                                _ = v87
                                _ = v88
                            }
                            fmt.Println(fi14)
                            _ = v84
                        }
                        switch (-(fi14)) {
                        case 301:
                            v97 := []int{fi14, 349}
                            _ = v97
                        case -245:
                            if false {
                                p0 -= p0
                            }
                        case 0:
                            p0 -= -2147483648
                        default:
                            if ("\n\t" <= "\n\t") {
                                switch (fi11 << ((833) & 63)) {
                                case 42:
                                    switch (fi14 % ((0) | 1)) {
                                    case -691:
                                        var v98 float64 = -2.25
                                        _ = v98
                                    case 32767:
                                        p0 -= (p0 & -787)
                                    case 128:
                                        p0 -= -968
                                    default:
                                        p0 -= -743
                                    }
                                }
                            }
                        }
                    } else {
                        switch 9223372036854775807 {
                        case 0:
                            p0 += p0
                        case 127:
                            var v99 string = "foo bar"
                            _ = v99
                        case -753:
                            var v100 int = (-(fi11))
                            _ = v100
                        }
                    }
                    fmt.Println(fi11, fi11)
                    v101 := []int{fi14}
                    switch ("foo bar" + ("" + ("" + "a"))) {
                    case "hello":
                        p0 += (p0 ^ -1)
                    case "a":
                        p0 = (^(p0))
                    default:
                        p0 = 0
                    }
                    _ = fi14
                    _ = v101
                }
            } else {
                if false {
                    var v102 uint32 = 2147483647
                    _ = v102
                }
                v103 := []int{}
                switch fi11 {
                case 9223372036854775807:
                    if true {
                        var v104 uint = 455
                        var v105 bool = false
                        _ = v104
                        _ = v105
                    }
                case 42:
                    fmt.Println(p0)
                default:
                    var v106 uint8 = 1
                    _ = v106
                }
                _ = v103
            }
            v107 := []int{fi11, fi11, fi11}
            _ = fi11
            _ = v107
        }
        if (42 >= 32767) {
            var v108 int16 = 144
            _ = v108
        } else {
            var v109 string = ((("a" + ((("hello" + "") + "a") + "")) + ("hello" + ("a" + "hello"))) + ("foo bar" + "a"))
            p0 += p0
            switch 0 {
            case 4294967295:
                if true {
                    fmt.Println(v109)
                }
            case -684:
                v110 := GenS7{F0: 18446744073709551615}
                _ = v110
            }
            _ = v109
        }
    }
    var v111 int = 9223372036854775807
    _ = p0
    _ = v111
    return v111
}

func main() {
    v1 := []string{"\n\t", "hello", ((("" + "\n\t") + (("hello" + ("" + "")) + "a")) + ((("foo bar" + "foo bar") + "foo bar") + "\n\t"))}
    v2 := make(map[uint32]bool)
    v3 := []string{"hello", "\n\t", (("\n\t" + ("\n\t" + "\n\t")) + "foo bar"), "foo bar"}
    v112 := v2[670]
    var v113 uint = 255
    for fi114 := 0; fi114 < 19; fi114++ {
        fmt.Println(v112, fi114)
        _ = fi114
    }
    if v112 {
        _ = len(v2)
        v115 := GenS15{F0: genFn1(474)}
        v116 := GenS6{F0: 892, F1: -2147483648, F2: 128}
        _ = v115
        _ = v116
    } else {
        if true {
            for fi117 := 0; fi117 < 7; fi117++ {
                v1 = append(v1, (("\n\t" + "a") + "\n\t"))
                var v118 uint8 = 76
                _ = fi117
                _ = v118
            }
            var v119 bool = (65535 <= 2)
            _ = v119
        }
        v113 += v113
        for fi120 := 0; fi120 < 20; fi120++ {
            if false {
                fmt.Println(fi120, v113)
                v121 := GenS16{F0: -65, F1: 41, F2: 2}
                _ = v121
            } else {
                v112 = (100 > 127)
                v2[83] = v112
            }
            var v122 uint = v113
            switch fi120 {
            case -323:
                v3 = append(v3, ("" + ""))
            case 916:
                _ = len(v2)
            case -618:
                var v123 uint32 = 2
                _ = v123
            default:
                if (-1 <= 0) {
                    v124 := GenS17{F0: "\n\t"}
                    v1 = append(v1, "foo bar")
                    _ = v124
                } else {
                    v112 = v112
                    _ = len(v2)
                }
            }
            _ = fi120
            _ = v122
        }
    }
    v113 -= v113
    v125 := GenS10{F0: "\n\t", F1: 42}
    for fi126 := 0; fi126 < 2; fi126++ {
        for fi127 := 0; fi127 < 8; fi127++ {
            var v128 uint = 18446744073709551615
            if len(v3) > 0 {
                v3[0] = "\n\t"
            }
            v129 := GenS5{F0: ("hello" + ""), F1: 891}
            _ = fi127
            _ = v128
            _ = v129
        }
        var v130 uint64 = 202
        for fi131 := 0; fi131 < 11; fi131++ {
            fmt.Println(fi126, v130, v113)
            fmt.Println(v113)
            var v132 int16 = 1
            _ = fi131
            _ = v132
        }
        _ = fi126
        _ = v130
    }
    fmt.Println(v112)
    for fi133 := 0; fi133 < 12; fi133++ {
        switch ("foo bar" + ((("hello" + "hello") + "") + "")) {
        case "a":
            v113 += (v113 * (v113 | v113))
        case "":
            var v134 uint64 = 42
            _ = v134
        default:
            switch (^(fi133)) {
            case -2147483648:
                var v135 uint64 = 2
                _ = v135
            }
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v136 := GenS11{F0: fi133, F1: 32767, F2: 2147483647}
        _ = fi133
        _ = v136
    }
    v113 = 927
    v112 = (!(42 <= 1))
    for fi137 := 0; fi137 < 10; fi137++ {
        var v138 float32 = 1e10
        if len(v1) > 0 {
            v1[0] = "\n\t"
        }
        v139 := GenS8{F0: "hello"}
        _ = fi137
        _ = v138
        _ = v139
    }
    v113 = (v113 % ((v113) | 1))
    var v140 uint8 = 154
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v112:", v112)
    fmt.Println("v113:", v113)
    fmt.Println("v140:", v140)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v112
    _ = v113
    _ = v125
    _ = v140
    _ = v2
    _ = v3
}
