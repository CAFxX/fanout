package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int16 }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 uint8; F1 int32; F2 int16 }
type GenS4 struct { F0 string; F1 string }
type GenS5 struct { F0 uint8; F1 int32; F2 uint32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint64; F1 uint64; F2 uint64 }
type GenS8 struct { F0 uint16; F1 uint16 }
type GenS9 struct { F0 int; F1 string; F2 int32 }
type GenS10 struct { F0 string; F1 int32 }
type GenS11 struct { F0 uint64; F1 int64 }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 uint; F1 int16 }
type GenS14 struct { F0 bool; F1 int }
type GenS15 struct { F0 uint8 }
func genFn0() int {
    var v7 uint64 = 0
    fmt.Println(v7)
    fmt.Println(v7)
    v7 = 4294967295
    v8 := GenS1{F0: 42, F1: -1, F2: 0}
    _ = v7
    _ = v8
    return 9223372036854775807
}

func main() {
    v1 := make(map[int]int)
    var v2 int16 = 1000
    var v3 float64 = -2.25
    var v4 uint = 1
    var v5 int16 = v2
    v6 := []int{1, 127, -1, 127}
    v4 = v4
    switch 4294967295 {
    case 132:
        for fi9 := 0; fi9 < 18; fi9++ {
            var v10 uint8 = uint8(v4)
            for fi11 := 0; fi11 < 16; fi11++ {
                fmt.Println(v10, v10)
                v12 := GenS2{F0: (fi11 * (^(fi11))), F1: 42}
                if (true && ((^(v10)) <= (v10 & 2))) {
                    v13 := GenS2{F0: (fi9 - (^(fi9))), F1: 4294967295}
                    var v14 uint16 = 1
                    _ = v13
                    _ = v14
                }
                _ = fi11
                _ = v12
            }
            v15 := GenS1{F0: genFn0(), F1: -69, F2: v2}
            _ = fi9
            _ = v10
            _ = v15
        }
    }
    if ((false || (!false)) && true) {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        var v16 int = genFn0()
        _ = v16
    } else {
        for fi17 := 0; fi17 < 12; fi17++ {
            fmt.Println(v5, fi17, v4)
            _ = len(v6)
            _ = fi17
        }
        v18 := GenS1{F0: 4294967295, F1: -128, F2: v2}
        v3 += v3
        _ = v18
    }
    for fi19 := 0; fi19 < 16; fi19++ {
        switch "\n\t" {
        case "hello":
            var v20 uint = v4
            _ = v20
        case "foo bar":
            var v21 uint8 = 2
            _ = v21
        case "a":
            switch ((("hello" + "a") + "a") + "foo bar") {
            case "foo bar":
                v1[(fi19 << ((4294967295) & 63))] = fi19
            case "a":
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            case "":
                if false {
                    var v22 int16 = 42
                    _ = v22
                } else {
                    v23 := GenS3{F0: 0, F1: -995, F2: v5}
                    if (-956 > -9223372036854775808) {
                        var v24 uint32 = 42
                        var v25 int16 = v2
                        fmt.Println(v24, v25)
                        _ = v24
                        _ = v25
                    } else {
                        if true {
                            var v26 int32 = 0
                            _ = v26
                        } else {
                            switch "\n\t" {
                            case "\n\t":
                                switch (fi19 >> ((v4) & 63)) {
                                case 127:
                                    var v27 uint16 = 181
                                    _ = v27
                                case -9223372036854775808:
                                    v28 := v1[(-(fi19))]
                                    _ = v28
                                case 4294967295:
                                    switch (("foo bar" + ((("hello" + "") + "a") + ("" + ("a" + "\n\t")))) + "foo bar") {
                                    case "foo bar":
                                        switch fi19 {
                                        case -46:
                                            for fi29 := 0; fi29 < 17; fi29++ {
                                                v3 -= (((v3 + ((float64(float32(v3)) / 1e10) / (1.5 + 1.5))) - v3) + v3)
                                                _ = fi29
                                            }
                                        }
                                    case "":
                                        switch ("foo bar" + ("hello" + ("\n\t" + ("" + "\n\t")))) {
                                        case "hello":
                                            v30 := GenS4{F0: ("foo bar" + ""), F1: "\n\t"}
                                            _ = v30
                                        default:
                                            var v31 uint64 = 564
                                            _ = v31
                                        }
                                    }
                                }
                            default:
                                if (-1 < -2147483648) {
                                    switch fi19 {
                                    case 1:
                                        fmt.Println(v2)
                                    default:
                                        for fi32 := 0; fi32 < 11; fi32++ {
                                            if ("foo bar" <= "hello") {
                                                if true {
                                                    v5 = (-(v2))
                                                    v33 := GenS3{F0: 128, F1: 204, F2: (v2 ^ 1)}
                                                    _ = v33
                                                } else {
                                                    v34 := GenS5{F0: uint8(v4), F1: 353, F2: 2147483648}
                                                    if true {
                                                        fmt.Println(v5)
                                                        var v35 int64 = -9223372036854775808
                                                        _ = v35
                                                    } else {
                                                        v36 := GenS4{F0: "a", F1: "foo bar"}
                                                        _ = v36
                                                    }
                                                    v4 += (v4 | (v4 - (v4 % ((641) | 1))))
                                                    _ = v34
                                                }
                                                var v37 bool = (1 <= -2147483648)
                                                var v38 float64 = float64(float32((((float64((float32((v3 / (v3 - v3))) + -2.25)) - ((v3 - float64(0.1)) + (0.0 / (v3 + (v3 / 1.5))))) / float64(3.14159)) - ((v3 + 0.1) * 1e10))))
                                                _ = v37
                                                _ = v38
                                            } else {
                                                var v39 int8 = 20
                                                fmt.Println(v5, v5, v2)
                                                var v40 uint32 = 1
                                                _ = v39
                                                _ = v40
                                            }
                                            if true {
                                                fmt.Println(v2, fi19, v5)
                                            }
                                            _ = fi32
                                        }
                                    }
                                    delete(v1, -9223372036854775808)
                                    if false {
                                        for fi41 := 0; fi41 < 5; fi41++ {
                                            v2 = v2
                                            fmt.Println(v2, fi19)
                                            _ = fi41
                                        }
                                    } else {
                                        for fi42 := 0; fi42 < 20; fi42++ {
                                            delete(v1, 32767)
                                            _ = fi42
                                        }
                                    }
                                } else {
                                    var v43 string = "a"
                                    v43 = "foo bar"
                                    _ = v43
                                }
                            }
                        }
                    }
                    switch (fi19 >> (((v4 + v4)) & 63)) {
                    case 1:
                        v4 = v4
                    default:
                        fmt.Println(fi19, v3)
                    }
                    _ = v23
                }
            }
        default:
            if ((1 <= -1) || (-1 != fi19)) {
                v3 = v3
                if false {
                    if (57 <= 32767) {
                        v2 = v5
                        fmt.Println(v5)
                    }
                    if (!(true || (false && (0 == 147)))) {
                        var v44 int64 = int64(v4)
                        _ = v44
                    } else {
                        var v45 int32 = -696
                        fmt.Println(v3)
                        _ = v45
                    }
                    v6 = append(v6, (fi19 + 4294967295))
                }
                switch -1 {
                case 32768:
                    v2 -= v2
                default:
                    var v46 uint8 = 42
                    _ = v46
                }
            } else {
                fmt.Println(fi19, v2, v3)
            }
        }
        var v47 int8 = 127
        _ = fi19
        _ = v47
    }
    for fi48 := 0; fi48 < 12; fi48++ {
        fmt.Println(v5, v3)
        var v49 uint = 2
        _ = fi48
        _ = v49
    }
    v3 = (v3 + 1.5)
    for fi50 := 0; fi50 < 4; fi50++ {
        var v51 float64 = (1.5 * (v3 + float64(1e10)))
        v52 := v1[fi50]
        if true {
            var v53 int32 = 42
            _ = v53
        } else {
            var v54 int64 = 941
            v52 -= fi50
            _ = v54
        }
        _ = fi50
        _ = v51
        _ = v52
    }
    switch (((("a" + "a") + "") + "hello") + "") {
    case "a":
        if true {
            for fi55 := 0; fi55 < 4; fi55++ {
                if (0 <= 167) {
                    _ = len(v6)
                    var v56 float32 = (((float32(v3) - 3.14159) + 0.1) + (3.14159 * 0.0))
                    _ = v56
                } else {
                    v2 = (-(v2))
                    for fi57 := 0; fi57 < 2; fi57++ {
                        var v58 int = (fi55 / ((fi57) | 1))
                        v59 := GenS6{F0: 795}
                        if ((1 < 65535) && false) {
                            if (382 >= 680) {
                                var v60 uint16 = 922
                                _ = v60
                            }
                            v58 = genFn0()
                            if false {
                                fmt.Println(v2, v5)
                                v4 -= (v4 & v4)
                                v61 := GenS7{F0: 2, F1: 2, F2: 0}
                                _ = v61
                            } else {
                                var v62 bool = true
                                switch 573 {
                                case 32768:
                                    fmt.Println(v5, v62, v5)
                                }
                                _ = v62
                            }
                        } else {
                            v1[(v58 << (((v4 + 65535)) & 63))] = 466
                            v63 := GenS5{F0: 118, F1: -414, F2: 2147483647}
                            var v64 uint32 = 2147483647
                            _ = v63
                            _ = v64
                        }
                        _ = fi57
                        _ = v58
                        _ = v59
                    }
                    v65 := GenS8{F0: 855, F1: 42}
                    _ = v65
                }
                _ = len(v1)
                _ = fi55
            }
            v3 -= (v3 / -2.25)
            if false {
                v4 += v4
                var v66 uint8 = 128
                _ = v66
            }
        } else {
            for fi67 := 0; fi67 < 20; fi67++ {
                switch ("" + ("\n\t" + ("\n\t" + "foo bar"))) {
                case "foo bar":
                    v4 = v4
                }
                _ = fi67
            }
        }
    case "foo bar":
        if true {
            var v68 uint8 = 255
            for fi69 := 0; fi69 < 8; fi69++ {
                v68 = 1
                var v70 uint64 = 0
                v71 := GenS4{F0: ("foo bar" + ""), F1: "foo bar"}
                _ = fi69
                _ = v70
                _ = v71
            }
            var v72 uint64 = 39
            _ = v68
            _ = v72
        }
    }
    v4 += 1
    v5 -= 32767
    v73 := GenS6{F0: 960}
    var v74 int64 = -427
    v75 := GenS6{F0: v74}
    fmt.Println(v3, v5, v5)
    for fi76 := 0; fi76 < 12; fi76++ {
        if (!false) {
            for fi77 := 0; fi77 < 20; fi77++ {
                v78 := GenS9{F0: genFn0(), F1: "foo bar", F2: -2147483648}
                switch fi76 {
                case -1:
                    if (42 <= (v4 / ((v4) | 1))) {
                        var v79 int16 = v2
                        _ = len(v1)
                        _ = v79
                    } else {
                        v80 := GenS4{F0: "\n\t", F1: "\n\t"}
                        if (true && (false && (!false))) {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            if (fi77 == 32768) {
                                v2 = v5
                                _ = len(v6)
                            }
                            var v81 int = fi77
                            v4 += (v4 >> (((v4 << ((v4) & 63))) & 63))
                            fmt.Println(fi77, fi77, v5)
                            v3 = 1.5
                            var v82 int = genFn0()
                            _ = v81
                            _ = v82
                        }
                        _ = v80
                    }
                    var v83 uint64 = 4294967295
                    if (!(!false)) {
                        if true {
                            switch "a" {
                            case "\n\t":
                                if (0.0 > v3) {
                                    fmt.Println(v5)
                                    switch (fi76 << ((v4) & 63)) {
                                    case 4294967295:
                                        if false {
                                            _ = len(v6)
                                            fmt.Println(fi77, fi77, v4)
                                        } else {
                                            v84 := GenS10{F0: "", F1: -2147483648}
                                            _ = v84
                                        }
                                        var v85 int16 = 32767
                                        _ = v85
                                    case 127:
                                        v83 += v83
                                    }
                                } else {
                                    v4 += (^(v4))
                                }
                            case "foo bar":
                                v6 = append(v6, fi76)
                            case "a":
                                v2 = 32767
                            default:
                                if (116 > 1) {
                                    v74 = 42
                                    v86 := GenS6{F0: v74}
                                    _ = v86
                                } else {
                                    if true {
                                        v87 := GenS7{F0: 378, F1: 106, F2: (v83 << ((v4) & 63))}
                                        if len(v6) > 0 {
                                            v6[0] = fi76
                                        }
                                        delete(v1, fi77)
                                        switch fi76 {
                                        case 127:
                                            switch "\n\t" {
                                            case "\n\t":
                                                v74 = v74
                                            case "foo bar":
                                                switch fi76 {
                                                case 125:
                                                    var v88 int8 = -1
                                                    _ = v88
                                                }
                                            }
                                        case 1:
                                            if true {
                                                v2 = 1000
                                            } else {
                                                var v89 int32 = 339
                                                fmt.Println(v3, v83, v74)
                                                _ = v89
                                            }
                                        default:
                                            if false {
                                                v6 = append(v6, fi77)
                                                switch "hello" {
                                                case "":
                                                    var v90 uint8 = 62
                                                    if (!true) {
                                                        if (0.0 < -2.25) {
                                                            v91 := GenS11{F0: (^(v83)), F1: v74}
                                                            _ = v91
                                                        }
                                                        switch "" {
                                                        case "a":
                                                            var v92 uint64 = v83
                                                            _ = v92
                                                        default:
                                                            v90 = (v90 << ((v4) & 63))
                                                        }
                                                        v3 = (((v3 + v3) + 0.0) + 1e10)
                                                    }
                                                    var v93 int32 = 1
                                                    _ = v90
                                                    _ = v93
                                                case "hello":
                                                    v6 = append(v6, fi76)
                                                case "a":
                                                    switch genFn0() {
                                                    case 127:
                                                        if (!true) {
                                                            var v94 int64 = 516
                                                            v95 := GenS12{F0: 42}
                                                            _ = v94
                                                            _ = v95
                                                        }
                                                    case -69:
                                                        if false {
                                                            v2 += -32768
                                                        } else {
                                                            fmt.Println(v74, v83, fi76)
                                                            v5 = (-(v2))
                                                            var v96 int32 = -157
                                                            v2 += (v5 - (v5 | 32767))
                                                            if (42 != 0) {
                                                                if (!(int64(v4) != v74)) {
                                                                    var v97 float64 = (v3 - v3)
                                                                    _ = v97
                                                                } else {
                                                                    v5 -= v5
                                                                }
                                                            }
                                                            _ = v96
                                                        }
                                                    case 369:
                                                        switch "" {
                                                        case "hello":
                                                            if false {
                                                                _ = len(v1)
                                                            } else {
                                                                v98 := GenS13{F0: v4, F1: (v2 + 0)}
                                                                v3 -= -2.25
                                                                var v99 uint16 = 0
                                                                _ = v98
                                                                _ = v99
                                                            }
                                                        case "\n\t":
                                                            if ((v4 << ((0) & 63)) < 194) {
                                                                var v100 uint8 = 255
                                                                if true {
                                                                    v101 := GenS14{F0: (!false), F1: (fi77 % ((573) | 1))}
                                                                    var v102 uint64 = 0
                                                                    var v103 uint = v4
                                                                    _ = v101
                                                                    _ = v102
                                                                    _ = v103
                                                                } else {
                                                                    v1[(fi76 - fi77)] = fi76
                                                                }
                                                                _ = v100
                                                            }
                                                        default:
                                                            var v104 int16 = v2
                                                            _ = v104
                                                        }
                                                    default:
                                                        var v105 uint16 = 255
                                                        v106 := GenS1{F0: (fi77 - fi76), F1: 127, F2: v2}
                                                        v107 := v1[int(v2)]
                                                        _ = v105
                                                        _ = v106
                                                        _ = v107
                                                    }
                                                default:
                                                    v108 := v1[(fi76 >> (((v4 * 773)) & 63))]
                                                    _ = v108
                                                }
                                                v3 += (v3 - (float64(0.0) + 1.5))
                                            } else {
                                                fmt.Println(fi77, v74)
                                            }
                                            var v109 uint32 = 42
                                            _ = len(v6)
                                            _ = v109
                                        }
                                        v2 = v5
                                        _ = v87
                                    }
                                    switch ((((("\n\t" + "foo bar") + "") + ("" + "")) + "foo bar") + ("foo bar" + "a")) {
                                    case "a":
                                        fmt.Println(v74, fi76, v2)
                                    }
                                    var v110 float32 = (float32(float64(-2.25)) + 1.5)
                                    _ = v110
                                }
                            }
                            if len(v6) > 0 {
                                v6[0] = genFn0()
                            }
                            var v111 float64 = 0.0
                            _ = v111
                        } else {
                            v112 := GenS11{F0: 337, F1: (v74 % (((v74 >> (((^(v4))) & 63))) | 1))}
                            var v113 uint32 = 927
                            _ = v112
                            _ = v113
                        }
                        if true {
                            switch genFn0() {
                            case 2147483647:
                                v4 = 1
                            default:
                                var v114 float32 = 3.14159
                                _ = v114
                            }
                            v4 += v4
                        }
                        if true {
                            var v115 int8 = 0
                            v4 -= 691
                            _ = v115
                        }
                    }
                    _ = v83
                case 4294967295:
                    if false {
                        _ = len(v6)
                        fmt.Println(v74)
                        fmt.Println(v5, v2, fi77)
                    } else {
                        switch fi77 {
                        case 128:
                            if (!false) {
                                v116 := GenS7{F0: 4294967295, F1: 857, F2: 4294967295}
                                _ = v116
                            } else {
                                var v117 int8 = -111
                                fmt.Println(v3)
                                var v118 bool = (-2.25 == (-2.25 - (0.0 * 0.0)))
                                _ = v117
                                _ = v118
                            }
                        }
                        var v119 uint32 = 2147483648
                        v1[fi77] = 32767
                        v120 := GenS15{F0: 5}
                        _ = v119
                        _ = v120
                    }
                case 357:
                    if false {
                        fmt.Println(v4, v2)
                        v121 := GenS2{F0: genFn0(), F1: 42}
                        var v122 uint8 = 64
                        _ = v121
                        _ = v122
                    }
                }
                _ = fi77
                _ = v78
            }
        } else {
            if len(v6) > 0 {
                v6[0] = fi76
            }
            v4 = v4
        }
        if (true && false) {
            v123 := GenS11{F0: 1, F1: (v74 / ((v74) | 1))}
            if (18446744073709551615 >= 4294967295) {
                for fi124 := 0; fi124 < 9; fi124++ {
                    v6 = append(v6, fi76)
                    if len(v6) > 0 {
                        v6[0] = fi124
                    }
                    v125 := GenS7{F0: 1, F1: 18446744073709551615, F2: 0}
                    _ = fi124
                    _ = v125
                }
                switch (fi76 % (((fi76 / (((^(fi76))) | 1))) | 1)) {
                case 2147483647:
                    var v126 float64 = (v3 - ((((v3 / v3) * v3) / v3) - v3))
                    _ = v126
                }
            }
            var v127 uint8 = 216
            _ = v123
            _ = v127
        } else {
            var v128 int8 = 121
            var v129 int16 = (v5 >> ((v4) & 63))
            _ = v128
            _ = v129
        }
        fmt.Println(v2)
        _ = fi76
    }
    for fi130 := 0; fi130 < 16; fi130++ {
        v5 = v2
        _ = fi130
    }
    for fi131 := 0; fi131 < 20; fi131++ {
        var v132 int = (^(fi131))
        fmt.Println(v2, v5)
        _ = fi131
        _ = v132
    }
    var v133 int = genFn0()
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v133:", v133)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v133
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v73
    _ = v74
    _ = v75
}
