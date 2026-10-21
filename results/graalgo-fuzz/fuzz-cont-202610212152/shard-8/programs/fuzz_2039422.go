package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint16; F1 int32; F2 int64 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint; F1 int64; F2 int16 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int }
type GenS11 struct { F0 bool; F1 int64 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int32 = 0
    var v3 int8 = 42
    switch 32767 {
    case 438:
        v4 := []int{32768, 9223372036854775807, 422, 2147483647}
        _ = v4
    case 1:
        var v5 int = -2147483648
        _ = v5
    }
    switch "foo bar" {
    case "\n\t":
        for fi6 := 0; fi6 < 20; fi6++ {
            v7 := GenS1{F0: -1, F1: 9223372036854775807, F2: fi6}
            if true {
                v3 = (v3 + v3)
                fmt.Println(v2)
            } else {
                if (!true) {
                    var v8 bool = (!((false || (0 >= 1)) || (634 < 499)))
                    if (127 >= 127) {
                        var v9 bool = (fi6 == (fi6 * (^(fi6))))
                        for fi10 := 0; fi10 < 18; fi10++ {
                            v3 += 100
                            switch "a" {
                            case "\n\t":
                                v11 := []int{fi10, (fi10 & (fi6 + (fi6 & (^(fi10))))), fi10, (fi6 >> ((1) & 63))}
                                _ = v11
                            case "a":
                                fmt.Println(fi6, v9)
                            }
                            var v12 float32 = ((1e10 - 0.1) - (((1.5 + -2.25) - 0.0) / 0.1))
                            fmt.Println(fi10, v8, v12)
                            var v13 int = 1
                            _ = fi10
                            _ = v12
                            _ = v13
                        }
                        _ = v9
                    }
                    _ = v8
                }
            }
            v2 += v2
            _ = fi6
            _ = v7
        }
    case "":
        for fi14 := 0; fi14 < 10; fi14++ {
            for fi15 := 0; fi15 < 17; fi15++ {
                var v16 float32 = 0.0
                _ = fi15
                _ = v16
            }
            _ = fi14
        }
    case "hello":
        for fi17 := 0; fi17 < 5; fi17++ {
            v18 := GenS2{F0: ("hello" + "foo bar")}
            var v19 string = ""
            for fi20 := 0; fi20 < 15; fi20++ {
                v19 = "foo bar"
                switch "\n\t" {
                case "a":
                    fmt.Println(v3, fi20)
                case "\n\t":
                    if true {
                        var v21 bool = true
                        var v22 float32 = (((1e10 / ((float32(1.5) + (-2.25 + float32(1.5))) * float32(3.14159))) * float32((1e10 / 0.1))) * 1.5)
                        _ = v21
                        _ = v22
                    }
                case "":
                    var v23 float64 = (float64(1e10) * (3.14159 * 1e10))
                    _ = v23
                default:
                    v24 := GenS3{F0: v2}
                    _ = v24
                }
                _ = fi20
            }
            _ = fi17
            _ = v18
            _ = v19
        }
    default:
        var v25 int = 9223372036854775807
        _ = v25
    }
    v3 = (v3 + v3)
    var v26 int64 = -821
    v27 := v1[""]
    if (128 != 0) {
        var v28 int32 = v2
        if ((42 != -1) || (v3 < 100)) {
            if (v28 != v2) {
                v1["a"] = v27
            }
            for fi29 := 0; fi29 < 10; fi29++ {
                if ((v27 * v27) > 42) {
                    var v30 bool = true
                    v3 -= v3
                    _ = v30
                } else {
                    v27 -= v27
                    if (-2.25 >= (-2.25 * 0.0)) {
                        if false {
                            var v31 int = (fi29 / ((128) | 1))
                            v32 := []int{(-(fi29))}
                            _ = len(v1)
                            _ = v31
                            _ = v32
                        }
                    }
                    if true {
                        for fi33 := 0; fi33 < 12; fi33++ {
                            if false {
                                v34 := v1[("a" + ("hello" + (("\n\t" + "a") + ("\n\t" + ("hello" + "\n\t")))))]
                                if true {
                                    if ((v26 ^ v26) > (v26 >> ((635) & 63))) {
                                        if false {
                                            if true {
                                                var v35 int64 = (-(v26))
                                                var v36 int8 = (v3 >> ((255) & 63))
                                                v2 += v2
                                                _ = v35
                                                _ = v36
                                            } else {
                                                switch (-(fi33)) {
                                                case 856:
                                                    if true {
                                                        fmt.Println(v3)
                                                    } else {
                                                        v37 := []int{32768, -57, fi33, (-(fi29))}
                                                        v1["\n\t"] = 2
                                                        _ = v37
                                                    }
                                                case 916:
                                                    v27 += v34
                                                    fmt.Println(v28, v34)
                                                case 2147483647:
                                                    var v38 string = "a"
                                                    _ = v38
                                                default:
                                                    v2 += v28
                                                }
                                                switch (fi29 / ((fi29) | 1)) {
                                                case 2147483647:
                                                    var v39 uint = 42
                                                    _ = v39
                                                case -142:
                                                    var v40 string = ("\n\t" + "foo bar")
                                                    _ = v40
                                                }
                                                switch int(v26) {
                                                case -23:
                                                    if (65535 == 255) {
                                                        if ((0.0 * (1.5 / ((1.5 + 0.0) - 0.0))) <= 3.14159) {
                                                            fmt.Println(fi29, v34, fi29)
                                                        }
                                                        v28 = (v2 % ((v28) | 1))
                                                        if false {
                                                            var v41 float32 = 1e10
                                                            _ = v41
                                                        }
                                                    }
                                                case -345:
                                                    v34 += (v27 << ((775) & 63))
                                                case 0:
                                                    v42 := GenS4{F0: true}
                                                    switch "\n\t" {
                                                    case "a":
                                                        if (v26 != 517) {
                                                            switch ("hello" + "hello") {
                                                            case "a":
                                                                switch "a" {
                                                                case "hello":
                                                                    var v43 uint32 = (v27 ^ 1)
                                                                    _ = v43
                                                                case "\n\t":
                                                                    fmt.Println(v34, fi29, v27)
                                                                case "a":
                                                                    var v44 uint16 = 521
                                                                    _ = v44
                                                                }
                                                            case "hello":
                                                                if true {
                                                                    var v45 uint = 65535
                                                                    switch (("\n\t" + "foo bar") + "foo bar") {
                                                                    case "":
                                                                        _ = len(v1)
                                                                    default:
                                                                        v46 := GenS5{F0: 32767, F1: (v28 & (v2 & v28)), F2: v26}
                                                                        _ = v46
                                                                    }
                                                                    v3 = (-(v3))
                                                                    _ = v45
                                                                } else {
                                                                    v3 -= (^(v3))
                                                                    v47 := []int{}
                                                                    _ = v47
                                                                }
                                                            default:
                                                                if ((v3 - v3) > 127) {
                                                                    var v48 float64 = float64(0.1)
                                                                    _ = v48
                                                                }
                                                            }
                                                            v27 -= (v34 + v27)
                                                        } else {
                                                            v49 := GenS6{F0: 65535}
                                                            fmt.Println(v34, v26)
                                                            _ = v49
                                                        }
                                                    case "\n\t":
                                                        v28 = 2147483647
                                                    case "foo bar":
                                                        var v50 int64 = (-(v26))
                                                        _ = v50
                                                    }
                                                    _ = v42
                                                }
                                                if ((9223372036854775807 >= (fi29 >> ((699) & 63))) || true) {
                                                    switch fi33 {
                                                    case -1:
                                                        var v51 uint16 = 255
                                                        _ = v51
                                                    case 2147483647:
                                                        v28 += (v28 * v28)
                                                    default:
                                                        fmt.Println(v26, v26)
                                                        var v52 int = -62
                                                        _ = v52
                                                    }
                                                    var v53 int64 = -933
                                                    v3 = (v3 << ((464) & 63))
                                                    var v54 int16 = -764
                                                    _ = v53
                                                    _ = v54
                                                }
                                                v55 := GenS1{F0: 1, F1: v26, F2: 128}
                                                _ = v55
                                            }
                                        } else {
                                            v56 := GenS5{F0: 924, F1: v28, F2: 95}
                                            var v57 int8 = 100
                                            _ = v56
                                            _ = v57
                                        }
                                        v28 = (v2 >> ((2) & 63))
                                        v34 -= v34
                                    } else {
                                        var v58 uint64 = 768
                                        if (-519 <= 203) {
                                            v34 = (v34 - (v27 ^ v27))
                                            _ = len(v1)
                                        } else {
                                            v59 := GenS7{F0: -710}
                                            _ = len(v1)
                                            v2 = v28
                                            var v60 uint64 = v58
                                            _ = v59
                                            _ = v60
                                        }
                                        _ = v58
                                    }
                                    var v61 float64 = 0.1
                                    var v62 int32 = -2147483648
                                    _ = v61
                                    _ = v62
                                } else {
                                    var v63 uint8 = 42
                                    _ = v63
                                }
                                _ = v34
                            } else {
                                switch "a" {
                                case "a":
                                    v64 := v1["\n\t"]
                                    _ = v64
                                case "":
                                    if (int64(v27) <= 636) {
                                        v65 := v1["foo bar"]
                                        if false {
                                            var v66 bool = ((-1 >= 32767) && false)
                                            v28 -= (v28 * (v2 & (v28 >> ((117) & 63))))
                                            var v67 int = (fi29 << ((uint(v27)) & 63))
                                            _ = v66
                                            _ = v67
                                        } else {
                                            fmt.Println(fi29, v28)
                                        }
                                        _ = v65
                                    }
                                case "foo bar":
                                    v68 := GenS8{F0: 65535, F1: v26, F2: -32768}
                                    var v69 string = "hello"
                                    var v70 int8 = v3
                                    v71 := GenS9{F0: (!false)}
                                    v72 := GenS3{F0: (v28 ^ -911)}
                                    v73 := []int{-780, int(v70)}
                                    delete(v1, v69)
                                    _ = v68
                                    _ = v69
                                    _ = v70
                                    _ = v71
                                    _ = v72
                                    _ = v73
                                default:
                                    _ = len(v1)
                                }
                                fmt.Println(v2, fi33, fi29)
                                switch (fi29 * fi33) {
                                case 128:
                                    if (201 >= 42) {
                                        switch "a" {
                                        case "":
                                            var v74 uint16 = 2
                                            _ = v74
                                        case "\n\t":
                                            fmt.Println(v26, fi33)
                                        default:
                                            v75 := GenS4{F0: false}
                                            _ = v75
                                        }
                                    } else {
                                        var v76 float64 = 3.14159
                                        var v77 int = fi29
                                        _ = v76
                                        _ = v77
                                    }
                                }
                                v26 += v26
                            }
                            v3 = -119
                            _ = fi33
                        }
                    } else {
                        v27 = (v27 ^ (v27 % ((v27) | 1)))
                        v78 := []int{fi29, (^(fi29))}
                        _ = v78
                    }
                }
                for fi79 := 0; fi79 < 6; fi79++ {
                    var v80 uint8 = 255
                    _ = fi79
                    _ = v80
                }
                v2 = (-(v2))
                _ = fi29
            }
        } else {
            v27 -= v27
        }
        v2 -= (v28 / ((v2) | 1))
        _ = v28
    }
    switch "hello" {
    case "":
        v2 += (v2 >> ((0) & 63))
    case "foo bar":
        v27 -= 4294967295
    default:
        switch 127 {
        case -652:
            fmt.Println(v26)
        case 316:
            switch (("a" + ("a" + "a")) + ((("\n\t" + "foo bar") + "hello") + ("" + (("\n\t" + "foo bar") + "hello")))) {
            case "a":
                switch -312 {
                case 299:
                    v81 := v1["\n\t"]
                    _ = v81
                case -114:
                    for fi82 := 0; fi82 < 6; fi82++ {
                        var v83 string = ""
                        _ = fi82
                        _ = v83
                    }
                }
            case "hello":
                for fi84 := 0; fi84 < 12; fi84++ {
                    switch "hello" {
                    case "\n\t":
                        var v85 int16 = 1
                        _ = v85
                    }
                    v2 -= -2147483648
                    v26 = v26
                    _ = fi84
                }
            case "foo bar":
                v3 -= (v3 << ((42) & 63))
            }
        case 32767:
            if (!false) {
                v2 = (v2 + v2)
            }
        default:
            var v86 float64 = float64(((3.14159 + 1.5) * 0.0))
            _ = v86
        }
    }
    if true {
        v87 := v1["a"]
        v88 := []int{9223372036854775807, 528, 1, 128}
        v87 -= (v27 | v87)
        _ = v87
        _ = v88
    } else {
        switch 42 {
        case -9223372036854775808:
            switch "" {
            case "hello":
                v2 = v2
            case "\n\t":
                for fi89 := 0; fi89 < 2; fi89++ {
                    fmt.Println(v26, v27)
                    fmt.Println(v2, v27)
                    var v90 uint32 = (v27 >> ((2) & 63))
                    _ = fi89
                    _ = v90
                }
            case "foo bar":
                v91 := GenS5{F0: 42, F1: -2147483648, F2: 722}
                _ = v91
            }
        case 42:
            v92 := GenS6{F0: 149}
            _ = v92
        case 127:
            v93 := []int{}
            _ = v93
        default:
            switch "a" {
            case "a":
                var v94 uint64 = 4294967295
                _ = v94
            case "hello":
                var v95 uint64 = 2
                _ = v95
            case "\n\t":
                if false {
                    var v96 uint64 = 4294967295
                    v3 -= (^(v3))
                    if (("hello" + "a") < ("\n\t" + "hello")) {
                        v96 = 0
                    } else {
                        v97 := GenS3{F0: -2147483648}
                        v1["\n\t"] = (v27 | (v27 % ((2147483647) | 1)))
                        _ = len(v1)
                        _ = v97
                    }
                    _ = v96
                }
            default:
                v98 := GenS10{F0: 492}
                _ = v98
            }
        }
        if true {
            if true {
                v27 -= (v27 >> ((uint(v27)) & 63))
                v99 := GenS6{F0: 18446744073709551615}
                _ = v99
            } else {
                for fi100 := 0; fi100 < 5; fi100++ {
                    switch fi100 {
                    case 0:
                        var v101 int8 = (v3 - (v3 + 127))
                        _ = v101
                    case 26:
                        v26 = (v26 / ((-295) | 1))
                    case -2147483648:
                        for fi102 := 0; fi102 < 20; fi102++ {
                            v103 := GenS11{F0: (1 != 42), F1: (v26 % ((v26) | 1))}
                            var v104 float64 = (1e10 + 0.0)
                            v105 := GenS5{F0: 863, F1: v2, F2: (v26 << ((2) & 63))}
                            _ = fi102
                            _ = v103
                            _ = v104
                            _ = v105
                        }
                    default:
                        var v106 int8 = v3
                        _ = v106
                    }
                    fmt.Println(v2)
                    _ = fi100
                }
                switch -330 {
                case -954:
                    fmt.Println(v26, v3)
                }
            }
            if ((true && true) && false) {
                switch "" {
                case "foo bar":
                    for fi107 := 0; fi107 < 19; fi107++ {
                        v108 := GenS1{F0: int32(fi107), F1: (v26 / (((v26 ^ (v26 ^ (v26 | (v26 ^ v26))))) | 1)), F2: 127}
                        var v109 float64 = (1.5 / -2.25)
                        v26 = (-(v26))
                        _ = fi107
                        _ = v108
                        _ = v109
                    }
                case "":
                    if (65 != 2) {
                        fmt.Println(v27)
                        v110 := GenS5{F0: 65535, F1: 187, F2: (v26 * (-(v26)))}
                        _ = v110
                    }
                case "hello":
                    v111 := GenS6{F0: 255}
                    _ = v111
                default:
                    var v112 uint8 = 128
                    _ = v112
                }
                if ((v2 % ((v2) | 1)) >= v2) {
                    v3 -= (v3 % ((v3) | 1))
                    for fi113 := 0; fi113 < 8; fi113++ {
                        for fi114 := 0; fi114 < 15; fi114++ {
                            v115 := []int{(fi113 * (^(fi114))), fi113, -1}
                            if (v27 <= v27) {
                                if true {
                                    var v116 int32 = -2147483648
                                    switch ((("a" + "hello") + "hello") + ("a" + "\n\t")) {
                                    case "a":
                                        var v117 float64 = -2.25
                                        _ = v117
                                    case "foo bar":
                                        v26 = (-(v26))
                                    default:
                                        v118 := GenS8{F0: 18446744073709551615, F1: v26, F2: 1}
                                        _ = v118
                                    }
                                    _ = v116
                                }
                                fmt.Println(v27, v26, fi114)
                            } else {
                                v2 = v2
                                if (true && (!false)) {
                                    var v119 string = ""
                                    _ = v119
                                }
                            }
                            _ = fi114
                            _ = v115
                        }
                        v120 := []int{int(v27)}
                        _ = fi113
                        _ = v120
                    }
                } else {
                    var v121 float32 = (3.14159 / 1e10)
                    if (321 >= uint(v27)) {
                        v122 := []int{32767, 177, 42, 9223372036854775807}
                        var v123 float32 = (((-2.25 + 1.5) - (v121 * (0.1 * float32(3.14159)))) / -2.25)
                        v124 := GenS11{F0: false, F1: v26}
                        _ = v122
                        _ = v123
                        _ = v124
                    }
                    _ = v121
                }
            }
        } else {
            if (-371 >= 2147483647) {
                fmt.Println(v3, v3)
                v27 += 4294967295
            }
        }
    }
    v125 := []int{4294967295, 42}
    if true {
        for fi126 := 0; fi126 < 7; fi126++ {
            if false {
                v26 -= v26
            }
            _ = fi126
        }
        v127 := GenS7{F0: 1000}
        _ = v127
    }
    var v128 int32 = -1
    v129 := GenS11{F0: (1.5 == (0.1 / (1e10 + float32((1.5 - (1e10 * (-2.25 * 1e10))))))), F1: v26}
    for fi130 := 0; fi130 < 20; fi130++ {
        for fi131 := 0; fi131 < 15; fi131++ {
            if (false && true) {
                v132 := v1[((("a" + "hello") + (((("foo bar" + "\n\t") + "\n\t") + ("" + ("foo bar" + "foo bar"))) + ("" + "a"))) + "hello")]
                _ = v132
            } else {
                if false {
                    v3 -= (-(v3))
                    var v133 uint64 = 1
                    _ = v133
                } else {
                    var v134 uint16 = 0
                    v128 = (v2 + v128)
                    _ = v134
                }
                fmt.Println(v2)
                var v135 int64 = (-(v26))
                _ = v135
            }
            if (0 <= (-(v2))) {
                switch "" {
                case "":
                    var v136 uint64 = 18446744073709551615
                    _ = v136
                case "a":
                    v137 := GenS11{F0: ((v128 / (((v2 | -1)) | 1)) < (v128 + v2)), F1: (v26 << ((0) & 63))}
                    _ = v137
                case "hello":
                    var v138 int8 = 100
                    _ = v138
                }
                v125 = append(v125, (-(fi130)))
                if true {
                    _ = len(v125)
                    if (280 >= 0) {
                        var v139 uint32 = v27
                        if len(v125) >= 2 {
                            v125 = v125[1:len(v125)]
                        }
                        _ = v139
                    } else {
                        if false {
                            var v140 float32 = 1.5
                            var v141 uint16 = 2
                            _ = v140
                            _ = v141
                        }
                    }
                } else {
                    switch "hello" {
                    case "":
                        if true {
                            var v142 string = ("\n\t" + "foo bar")
                            _ = v142
                        }
                    }
                    v27 -= v27
                }
            } else {
                if len(v125) >= 2 {
                    v125 = v125[1:len(v125)]
                }
            }
            v128 -= (v128 / ((v128) | 1))
            v27 = 2147483648
            _ = fi131
        }
        v3 = -1
        var v143 uint64 = uint64(v27)
        _ = fi130
        _ = v143
    }
    var v144 bool = (true || true)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v125:", len(v125), v125)
    fmt.Println("v128:", v128)
    fmt.Println("v144:", v144)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v125
    _ = v128
    _ = v129
    _ = v144
    _ = v2
    _ = v26
    _ = v27
    _ = v3
}
