package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 int8; F1 uint16; F2 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint16; F1 string; F2 int }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int16; F1 int64 }
type GenS8 struct { F0 string; F1 string }
type GenS9 struct { F0 int64; F1 int16; F2 int16 }
type GenS10 struct { F0 int; F1 uint16 }
type GenS11 struct { F0 uint; F1 int }
type GenS12 struct { F0 int64 }
func genFn0(p0 uint16, p1 uint32) int32 {
    v6 := []int{4294967295, -9223372036854775808, 9223372036854775807, 88}
    var v7 float64 = (1.5 + 0.0)
    switch "" {
    case "":
        p1 -= p1
    default:
        v8 := GenS1{F0: "\n\t", F1: p0}
        _ = v8
    }
    for fi9 := 0; fi9 < 20; fi9++ {
        if (87 < 1) {
            var v10 int = (fi9 * -9223372036854775808)
            fmt.Println(v10, fi9)
            _ = v10
        } else {
            v11 := GenS1{F0: "a", F1: 255}
            var v12 uint64 = 0
            if true {
                for fi13 := 0; fi13 < 4; fi13++ {
                    v12 = (v12 / (((v12 & v12)) | 1))
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    _ = fi13
                }
                for fi14 := 0; fi14 < 19; fi14++ {
                    v12 = 0
                    v12 = (^(v12))
                    _ = fi14
                }
            }
            _ = v11
            _ = v12
        }
        var v15 int8 = 0
        v16 := GenS1{F0: "hello", F1: (p0 - (p0 / (((p0 << ((65535) & 63))) | 1)))}
        _ = fi9
        _ = v15
        _ = v16
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 1
}
func genFn1() uint {
    if (!(false || true)) {
        var v17 uint64 = 42
        var v18 float64 = float64(-2.25)
        v17 += 0
        _ = v17
        _ = v18
    } else {
        var v19 uint16 = 1
        var v20 uint8 = 136
        _ = v19
        _ = v20
    }
    switch 42 {
    case 1:
        var v21 int32 = 0
        _ = v21
    case -1:
        if false {
            for fi22 := 0; fi22 < 14; fi22++ {
                v23 := GenS1{F0: ("\n\t" + "foo bar"), F1: 255}
                var v24 float64 = 3.14159
                _ = fi22
                _ = v23
                _ = v24
            }
            v25 := GenS2{F0: 1, F1: 65535, F2: 0}
            var v26 uint64 = 1
            _ = v25
            _ = v26
        } else {
            var v27 uint = 160
            if true {
                if (true && (false && true)) {
                    v28 := GenS3{F0: -128}
                    for fi29 := 0; fi29 < 5; fi29++ {
                        fmt.Println(fi29, v27)
                        fmt.Println(v27)
                        v27 = 1
                        _ = fi29
                    }
                    _ = v28
                }
            } else {
                if true {
                    if (v27 <= 263) {
                        v30 := GenS1{F0: ("\n\t" + "\n\t"), F1: 0}
                        _ = v30
                    }
                    switch -260 {
                    case 4294967295:
                        v31 := []int{9223372036854775807}
                        _ = v31
                    }
                    v27 = (^(v27))
                } else {
                    if ((79 <= 1) && (false && false)) {
                        var v32 uint64 = 1
                        v33 := []int{1, 2147483647, -1}
                        _ = v32
                        _ = v33
                    }
                    if false {
                        switch ("\n\t" + ((("foo bar" + "\n\t") + ((("\n\t" + "hello") + "\n\t") + ((("a" + "a") + "a") + "hello"))) + "hello")) {
                        case "hello":
                            if true {
                                switch "\n\t" {
                                case "hello":
                                    for fi34 := 0; fi34 < 6; fi34++ {
                                        var v35 uint = (v27 * (^(v27)))
                                        _ = fi34
                                        _ = v35
                                    }
                                case "a":
                                    fmt.Println(v27)
                                default:
                                    v27 -= (v27 & (^(v27)))
                                }
                                v36 := GenS2{F0: -1, F1: 1, F2: 0}
                                _ = v36
                            } else {
                                v37 := GenS4{F0: 1, F1: "\n\t", F2: 303}
                                v38 := GenS2{F0: 94, F1: 255, F2: -2147483648}
                                var v39 uint = 21
                                _ = v37
                                _ = v38
                                _ = v39
                            }
                        case "":
                            v27 = (v27 * 0)
                        case "a":
                            for fi40 := 0; fi40 < 5; fi40++ {
                                var v41 uint64 = 18446744073709551615
                                v27 -= v27
                                _ = fi40
                                _ = v41
                            }
                        }
                        v42 := []int{}
                        _ = v42
                    }
                    switch ("a" + "a") {
                    case "foo bar":
                        switch -478 {
                        case 9223372036854775807:
                            v43 := []int{921, 127, 9223372036854775807, 0}
                            _ = v43
                        }
                    case "hello":
                        var v44 int64 = -768
                        _ = v44
                    }
                }
                fmt.Println(v27)
            }
            _ = v27
        }
    }
    return 0
}

func main() {
    var v1 int8 = 127
    var v2 int64 = -336
    var v3 uint16 = 0
    var v4 float64 = (-2.25 / 0.1)
    v5 := []string{(("hello" + "foo bar") + (("foo bar" + "a") + "")), "a", "", "a"}
    v45 := GenS5{F0: 1}
    var v46 uint16 = 255
    var v47 int = 42
    v47 = (v47 / ((int(v46)) | 1))
    if ((-2.25 >= (v4 / float64(0.1))) && (1 <= -81)) {
        for fi48 := 0; fi48 < 1; fi48++ {
            var v49 uint8 = 0
            _ = fi48
            _ = v49
        }
        v3 -= (v46 & (v3 & 65535))
    } else {
        var v50 bool = true
        v5 = append(v5, "foo bar")
        _ = v50
    }
    if len(v5) > 0 {
        v5[0] = (("foo bar" + ((("\n\t" + (("hello" + ("foo bar" + ("hello" + "foo bar"))) + "hello")) + "foo bar") + "hello")) + "hello")
    }
    if false {
        switch v47 {
        case 4294967295:
            v51 := GenS6{F0: 255}
            _ = v51
        }
    }
    var v52 int16 = 0
    if (0 >= int16(v2)) {
        v3 += (^(v3))
        if false {
            var v53 uint = 65535
            for fi54 := 0; fi54 < 15; fi54++ {
                v55 := GenS5{F0: (^(v52))}
                if true {
                    v3 = (v3 | (v46 + 479))
                }
                if true {
                    if false {
                        for fi56 := 0; fi56 < 15; fi56++ {
                            var v57 float32 = 1e10
                            v5 = append(v5, "")
                            v46 += v3
                            _ = fi56
                            _ = v57
                        }
                        if (0.1 < 0.1) {
                            if ((v1 | v1) == (^(v1))) {
                                if true {
                                    fmt.Println(v4)
                                    for fi58 := 0; fi58 < 9; fi58++ {
                                        var v59 int32 = 2147483647
                                        var v60 uint = (v53 * (v53 ^ (v53 % ((4294967295) | 1))))
                                        _ = fi58
                                        _ = v59
                                        _ = v60
                                    }
                                } else {
                                    if (42 != (v2 + 42)) {
                                        if (!true) {
                                            for fi61 := 0; fi61 < 20; fi61++ {
                                                v4 = v4
                                                var v62 uint16 = (^(v3))
                                                switch ("hello" + "\n\t") {
                                                case "\n\t":
                                                    v1 -= 89
                                                case "a":
                                                    var v63 bool = true
                                                    _ = v63
                                                case "foo bar":
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                }
                                                _ = fi61
                                                _ = v62
                                            }
                                            fmt.Println(v2, v46)
                                            for fi64 := 0; fi64 < 6; fi64++ {
                                                if (fi54 > (v47 & (-(v47)))) {
                                                    v1 = v1
                                                }
                                                var v65 int = (-(fi54))
                                                _ = fi64
                                                _ = v65
                                            }
                                        } else {
                                            var v66 float64 = v4
                                            _ = v66
                                        }
                                        if true {
                                            for fi67 := 0; fi67 < 20; fi67++ {
                                                var v68 int32 = 2147483647
                                                if ((v3 << ((uint(v3)) & 63)) > 255) {
                                                    var v69 uint = (^(v53))
                                                    switch "\n\t" {
                                                    case "":
                                                        v47 = fi54
                                                    case "\n\t":
                                                        var v70 int32 = -1
                                                        _ = v70
                                                    default:
                                                        v71 := GenS3{F0: v1}
                                                        _ = v71
                                                    }
                                                    switch (fi67 - -2147483648) {
                                                    case 9223372036854775807:
                                                        var v72 float64 = v4
                                                        _ = v72
                                                    case 0:
                                                        if len(v5) > 0 {
                                                            v5[0] = ("foo bar" + ("a" + ""))
                                                        }
                                                    case 2147483647:
                                                        v47 = 1
                                                    }
                                                    _ = v69
                                                } else {
                                                    if (true || ((!false) && true)) {
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                        var v73 int32 = v68
                                                        var v74 uint8 = 228
                                                        v74 -= (v74 ^ (v74 % ((v74) | 1)))
                                                        _ = v73
                                                        _ = v74
                                                    }
                                                }
                                                var v75 float64 = (v4 + (v4 * -2.25))
                                                _ = fi67
                                                _ = v68
                                                _ = v75
                                            }
                                        }
                                    } else {
                                        var v76 bool = false
                                        v77 := GenS4{F0: (v3 + v46), F1: "\n\t", F2: (^(v47))}
                                        _ = v76
                                        _ = v77
                                    }
                                    for fi78 := 0; fi78 < 13; fi78++ {
                                        var v79 uint64 = 42
                                        v80 := GenS2{F0: (-(v1)), F1: v3, F2: 568}
                                        v4 = v4
                                        _ = fi78
                                        _ = v79
                                        _ = v80
                                    }
                                }
                                for fi81 := 0; fi81 < 15; fi81++ {
                                    var v82 uint = genFn1()
                                    var v83 uint8 = 2
                                    var v84 int64 = -426
                                    _ = fi81
                                    _ = v82
                                    _ = v83
                                    _ = v84
                                }
                                v5 = append(v5, ("hello" + (("" + "\n\t") + ("hello" + "a"))))
                            } else {
                                var v85 bool = (2 < 65535)
                                _ = v85
                            }
                            if (v3 < 32767) {
                                var v86 uint16 = uint16(v53)
                                if true {
                                    var v87 float32 = (1.5 - float32((0.0 + (((v4 / v4) * float64(((1e10 + ((1.5 / (float32(v4) * ((-2.25 + 3.14159) - (3.14159 * float32(1.5))))) * ((3.14159 / (3.14159 / 1.0)) / 0.1))) * -2.25))) * ((float64((((((((-2.25 + -2.25) - (1.5 * 3.14159)) / float32(v4)) * (((0.0 - 0.0) - 0.0) + (float32(v4) - 1.5))) / -2.25) - float32(v4)) / (1.5 + ((0.1 * float32(0.0)) - 0.1)))) / v4) - v4)))))
                                    v5 = append(v5, ("" + ("" + "")))
                                    _ = v87
                                }
                                v88 := GenS7{F0: int16(v53), F1: v2}
                                _ = v86
                                _ = v88
                            } else {
                                var v89 string = "a"
                                var v90 int32 = 1
                                _ = v89
                                _ = v90
                            }
                        } else {
                            if false {
                                if true {
                                    var v91 int = (v47 / ((fi54) | 1))
                                    if true {
                                        var v92 int = v91
                                        fmt.Println(v4, v46)
                                        _ = v92
                                    }
                                    if (42 <= 1) {
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        v4 += v4
                                        fmt.Println(v1)
                                    } else {
                                        v4 -= v4
                                    }
                                    _ = v91
                                } else {
                                    switch "" {
                                    case "\n\t":
                                        v93 := GenS2{F0: v1, F1: 2, F2: genFn0(255, 42)}
                                        _ = v93
                                    case "foo bar":
                                        fmt.Println(v4, v46, fi54)
                                    default:
                                        switch ("" + "foo bar") {
                                        case "hello":
                                            if true {
                                                for fi94 := 0; fi94 < 15; fi94++ {
                                                    if ((^(fi54)) < (v47 * fi54)) {
                                                        v95 := GenS8{F0: "hello", F1: "\n\t"}
                                                        var v96 float64 = (v4 / v4)
                                                        _ = v95
                                                        _ = v96
                                                    }
                                                    var v97 int64 = int64(v46)
                                                    _ = fi94
                                                    _ = v97
                                                }
                                                for fi98 := 0; fi98 < 19; fi98++ {
                                                    var v99 int64 = int64(v1)
                                                    _ = fi98
                                                    _ = v99
                                                }
                                            }
                                        case "":
                                            switch 127 {
                                            case -9223372036854775808:
                                                var v100 bool = (v4 < (v4 + v4))
                                                _ = v100
                                            default:
                                                for fi101 := 0; fi101 < 7; fi101++ {
                                                    v102 := GenS9{F0: (v2 * v2), F1: v52, F2: int16(v2)}
                                                    var v103 int16 = (^(v52))
                                                    var v104 uint16 = v46
                                                    _ = fi101
                                                    _ = v102
                                                    _ = v103
                                                    _ = v104
                                                }
                                            }
                                        default:
                                            fmt.Println(v52)
                                        }
                                    }
                                    if (("" + ("foo bar" + "a")) >= ("a" + ("hello" + ((((("" + ("hello" + (("hello" + "a") + "\n\t"))) + "") + "hello") + "\n\t") + "a")))) {
                                        v105 := GenS4{F0: v46, F1: ("" + "\n\t"), F2: (fi54 - 32768)}
                                        v47 = fi54
                                        _ = v105
                                    } else {
                                        var v106 float32 = 1e10
                                        _ = v106
                                    }
                                    v47 = v47
                                }
                                v53 = 383
                            } else {
                                v5 = append(v5, "hello")
                                v4 = 0.0
                            }
                            switch "foo bar" {
                            case "foo bar":
                                fmt.Println(v4)
                            case "":
                                v53 += (v53 / ((v53) | 1))
                            case "a":
                                fmt.Println(v52, v1, v1)
                            default:
                                for fi107 := 0; fi107 < 7; fi107++ {
                                    var v108 uint64 = 4294967295
                                    var v109 uint8 = 230
                                    _ = fi107
                                    _ = v108
                                    _ = v109
                                }
                            }
                        }
                        if true {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        } else {
                            v47 = v47
                            v46 += (v3 % ((42) | 1))
                        }
                    } else {
                        var v110 int16 = v52
                        var v111 float64 = 0.0
                        _ = v110
                        _ = v111
                    }
                    switch 32768 {
                    case 128:
                        v5 = append(v5, "")
                    default:
                        if true {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            var v112 float64 = (-2.25 + (1e10 + v4))
                            if false {
                                switch (v47 >> ((535) & 63)) {
                                case 9223372036854775807:
                                    for fi113 := 0; fi113 < 7; fi113++ {
                                        switch fi113 {
                                        case 202:
                                            if true {
                                                v114 := GenS6{F0: (^(v46))}
                                                fmt.Println(v4, fi113)
                                                v52 += v52
                                                _ = v114
                                            }
                                        case 0:
                                            switch "a" {
                                            case "hello":
                                                var v115 int = 1
                                                _ = v115
                                            case "foo bar":
                                                v116 := GenS10{F0: -878, F1: (^(v46))}
                                                _ = v116
                                            }
                                        case 42:
                                            var v117 int = fi54
                                            var v118 float32 = -2.25
                                            _ = v117
                                            _ = v118
                                        }
                                        _ = fi113
                                    }
                                case -848:
                                    for fi119 := 0; fi119 < 13; fi119++ {
                                        v5 = append(v5, ("foo bar" + "hello"))
                                        _ = fi119
                                    }
                                default:
                                    for fi120 := 0; fi120 < 8; fi120++ {
                                        v121 := GenS11{F0: (^(v53)), F1: (-(v47))}
                                        var v122 string = "hello"
                                        _ = fi120
                                        _ = v121
                                        _ = v122
                                    }
                                }
                            } else {
                                var v123 float32 = -2.25
                                v46 = (^(v3))
                                _ = v123
                            }
                            _ = v112
                        }
                    }
                }
                _ = fi54
                _ = v55
            }
            switch "hello" {
            case "":
                if false {
                    var v124 uint8 = 80
                    for fi125 := 0; fi125 < 18; fi125++ {
                        var v126 uint = (v53 - (^(v53)))
                        _ = fi125
                        _ = v126
                    }
                    fmt.Println(v46)
                    _ = v124
                }
            }
            _ = v53
        } else {
            v52 -= v52
            fmt.Println(v52, v46)
            v5 = append(v5, "\n\t")
        }
        var v127 string = ("" + "")
        _ = v127
    } else {
        if ("hello" > "a") {
            var v128 uint64 = 789
            for fi129 := 0; fi129 < 4; fi129++ {
                v130 := GenS2{F0: v1, F1: (^(v46)), F2: 1}
                var v131 int = -2147483648
                _ = fi129
                _ = v130
                _ = v131
            }
            var v132 uint32 = 501
            _ = v128
            _ = v132
        }
        var v133 uint = 18446744073709551615
        var v134 string = ""
        _ = v133
        _ = v134
    }
    for fi135 := 0; fi135 < 7; fi135++ {
        v136 := GenS2{F0: (v1 * (v1 * 1)), F1: v46, F2: 171}
        _ = fi135
        _ = v136
    }
    v137 := GenS12{F0: (v2 | v2)}
    for fi138 := 0; fi138 < 2; fi138++ {
        _ = len(v5)
        _ = fi138
    }
    v47 = v47
    var v139 int = -425
    fmt.Println("v1:", v1)
    fmt.Println("v139:", v139)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v137
    _ = v139
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v52
}
