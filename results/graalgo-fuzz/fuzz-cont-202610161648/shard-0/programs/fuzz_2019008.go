package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 uint16 }
type GenS2 struct { F0 int32; F1 string; F2 int16 }
type GenS3 struct { F0 uint16; F1 int; F2 int16 }
type GenS4 struct { F0 uint8; F1 int8; F2 uint }
type GenS5 struct { F0 int16; F1 int32 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int16; F1 uint64; F2 uint64 }
type GenS9 struct { F0 uint16; F1 int; F2 uint64 }
type GenS10 struct { F0 uint32; F1 int64; F2 string }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 string; F1 uint64; F2 bool }
type GenS13 struct { F0 bool; F1 int8; F2 uint8 }
func genFn0() int {
    v6 := GenS1{F0: 1, F1: false, F2: 499}
    v7 := []int{593, 4294967295, -9223372036854775808}
    var v8 float32 = 1.5
    for fi9 := 0; fi9 < 13; fi9++ {
        switch "hello" {
        case "foo bar":
            if false {
                var v10 uint16 = 255
                var v11 uint16 = v10
                _ = v10
                _ = v11
            }
        }
        var v12 float64 = 1e10
        fmt.Println(v8, v8, v12)
        _ = fi9
        _ = v12
    }
    _ = v6
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 uint) uint {
    var v13 uint64 = 42
    if true {
        var v14 uint64 = 2
        _ = v14
    } else {
        if true {
            for fi15 := 0; fi15 < 3; fi15++ {
                v16 := GenS1{F0: -1, F1: ((42 < 32767) && false), F2: 65535}
                _ = fi15
                _ = v16
            }
            v13 -= (v13 & (^(v13)))
        }
    }
    v17 := GenS1{F0: -615, F1: true, F2: 65535}
    _ = p0
    _ = v13
    _ = v17
    return p0
}

func main() {
    var v1 int64 = 42
    var v2 uint = 2
    var v3 int32 = -1
    var v4 uint64 = 785
    var v5 int16 = 37
    if false {
        v18 := GenS1{F0: 42, F1: (1 <= -1), F2: 531}
        v19 := []int{-2147483648, 9223372036854775807}
        _ = v18
        _ = v19
    } else {
        var v20 string = "foo bar"
        for fi21 := 0; fi21 < 7; fi21++ {
            if ((!(!true)) && false) {
                if (v2 < v2) {
                    v22 := []int{int(v2)}
                    _ = v22
                }
                switch fi21 {
                case -1:
                    for fi23 := 0; fi23 < 11; fi23++ {
                        var v24 bool = ((1e10 + -2.25) != 1e10)
                        var v25 uint16 = 442
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                }
                switch "hello" {
                case "foo bar":
                    v26 := GenS2{F0: 0, F1: ("\n\t" + "\n\t"), F2: v5}
                    _ = v26
                case "":
                    v27 := GenS1{F0: v3, F1: false, F2: 65535}
                    _ = v27
                case "hello":
                    v28 := []int{-27, (fi21 + (fi21 / (((fi21 / ((fi21) | 1))) | 1)))}
                    _ = v28
                }
            }
            v29 := GenS3{F0: 0, F1: (fi21 % ((fi21) | 1)), F2: (v5 + 1000)}
            if true {
                var v30 int8 = 1
                var v31 int8 = -1
                _ = v30
                _ = v31
            } else {
                v32 := []int{fi21, (fi21 & 2147483647), -9223372036854775808, (fi21 ^ fi21)}
                _ = v32
            }
            _ = fi21
            _ = v29
        }
        _ = v20
    }
    for fi33 := 0; fi33 < 18; fi33++ {
        switch "a" {
        case "":
            switch "foo bar" {
            case "foo bar":
                for fi34 := 0; fi34 < 13; fi34++ {
                    var v35 int = -147
                    v36 := GenS2{F0: v3, F1: ("" + "foo bar"), F2: (v5 + (^(v5)))}
                    v37 := []int{32768, genFn0(), (v35 & -224)}
                    _ = fi34
                    _ = v35
                    _ = v36
                    _ = v37
                }
            }
        case "foo bar":
            if false {
                if (42 > v2) {
                    v38 := GenS4{F0: 62, F1: 67, F2: 65535}
                    switch "\n\t" {
                    case "":
                        v39 := GenS1{F0: (v3 >> ((v2) & 63)), F1: (65535 != 2), F2: 255}
                        _ = v39
                    case "\n\t":
                        v2 = (^(v2))
                    default:
                        for fi40 := 0; fi40 < 12; fi40++ {
                            var v41 int64 = v1
                            _ = fi40
                            _ = v41
                        }
                    }
                    for fi42 := 0; fi42 < 14; fi42++ {
                        if (int64(v4) != (v1 % ((v1) | 1))) {
                            var v43 int16 = (-(v5))
                            _ = v43
                        } else {
                            v3 += 42
                            if true {
                                v44 := []int{int(v1)}
                                _ = v44
                            }
                        }
                        v45 := GenS5{F0: 42, F1: v3}
                        fmt.Println(v1, v2)
                        _ = fi42
                        _ = v45
                    }
                    _ = v38
                }
                var v46 int16 = v5
                if false {
                    fmt.Println(v2)
                    var v47 uint8 = 42
                    _ = v47
                }
                _ = v46
            }
        case "hello":
            v2 = (v2 % (((^(v2))) | 1))
        default:
            fmt.Println(v4, v4)
        }
        if ((true && false) && ((v2 >> ((genFn1(genFn1(genFn1(65535)))) & 63)) >= (^(v2)))) {
            if false {
                var v48 int16 = v5
                v49 := []int{(-(fi33)), -1, genFn0()}
                _ = v48
                _ = v49
            }
            v50 := GenS6{F0: 450}
            v4 = v4
            _ = v50
        }
        _ = fi33
    }
    var v51 int8 = 100
    v4 = v4
    if true {
        v52 := []int{-257}
        v53 := GenS7{F0: 65535}
        _ = v52
        _ = v53
    }
    if true {
        switch 2147483647 {
        case -9223372036854775808:
            v2 += (^(v2))
        case 1:
            v54 := []int{}
            _ = v54
        case -1:
            v2 += 2
        default:
            v2 = 18446744073709551615
        }
    }
    if false {
        v55 := GenS1{F0: int32(v2), F1: true, F2: 32767}
        var v56 int = 4294967295
        _ = v55
        _ = v56
    } else {
        for fi57 := 0; fi57 < 17; fi57++ {
            fmt.Println(v3)
            _ = fi57
        }
        var v58 float64 = 1.5
        _ = v58
    }
    if (!(!true)) {
        for fi59 := 0; fi59 < 11; fi59++ {
            v60 := []int{(fi59 << ((v2) & 63)), int(v51), (fi59 - fi59)}
            switch "" {
            case "\n\t":
                v60 = append(v60, (fi59 << ((genFn1(370)) & 63)))
            case "foo bar":
                for fi61 := 0; fi61 < 18; fi61++ {
                    if ((v5 >> ((1) & 63)) <= (v5 % ((v5) | 1))) {
                        fmt.Println(v1, fi61, fi59)
                        switch "foo bar" {
                        case "":
                            if ((-2.25 < float32(3.14159)) && false) {
                                var v62 float32 = 0.0
                                _ = v62
                            }
                        case "foo bar":
                            v63 := GenS4{F0: uint8(v2), F1: v51, F2: (v2 % ((18446744073709551615) | 1))}
                            _ = v63
                        case "a":
                            if (true && true) {
                                var v64 uint = genFn1(18446744073709551615)
                                _ = v64
                            } else {
                                v5 = (v5 << ((genFn1(v2)) & 63))
                            }
                        }
                        _ = len(v60)
                        if (!(uint16(v4) <= 2)) {
                            v3 -= v3
                        }
                    } else {
                        v3 += v3
                    }
                    _ = fi61
                }
            case "a":
                switch fi59 {
                case -1:
                    switch fi59 {
                    case 1:
                        fmt.Println(v2, v3)
                    case -2147483648:
                        switch (fi59 << ((genFn1((v2 | (v2 / ((1) | 1))))) & 63)) {
                        case -2147483648:
                            for fi65 := 0; fi65 < 19; fi65++ {
                                v3 += 1
                                _ = fi65
                            }
                        case 1:
                            var v66 float64 = ((float64(-2.25) / ((1e10 / 1e10) + (((1e10 * 1e10) - ((1.5 - -2.25) - -2.25)) * ((float64(((3.14159 - 0.0) - 0.0)) * 0.0) / 1.0)))) / float64((0.0 - 0.1)))
                            _ = v66
                        case -316:
                            if len(v60) > 0 {
                                v60[0] = (fi59 - fi59)
                            }
                        }
                    default:
                        _ = len(v60)
                    }
                }
            default:
                var v67 int8 = 100
                _ = v67
            }
            v5 = v5
            _ = fi59
            _ = v60
        }
        for fi68 := 0; fi68 < 17; fi68++ {
            switch ("hello" + "hello") {
            case "a":
                v3 -= -2147483648
            case "\n\t":
                v51 = int8(v2)
            case "hello":
                if true {
                    if ("" < ("hello" + ("" + "\n\t"))) {
                        var v69 float32 = -2.25
                        _ = v69
                    } else {
                        for fi70 := 0; fi70 < 7; fi70++ {
                            v71 := GenS2{F0: 2147483647, F1: ("a" + "\n\t"), F2: v5}
                            v72 := GenS8{F0: (v5 | (-(v5))), F1: (v4 >> ((0) & 63)), F2: 42}
                            _ = fi70
                            _ = v71
                            _ = v72
                        }
                        var v73 uint64 = (v4 ^ (v4 << (((v2 & (^(v2)))) & 63)))
                        _ = v73
                    }
                    for fi74 := 0; fi74 < 5; fi74++ {
                        v51 = int8(v5)
                        v4 += 0
                        v75 := []int{(fi68 ^ (fi74 & int(v4))), fi74, genFn0(), (fi74 * fi74)}
                        _ = fi74
                        _ = v75
                    }
                    v76 := GenS4{F0: 255, F1: (v51 | v51), F2: genFn1((v2 % ((genFn1((v2 << ((v2) & 63)))) | 1)))}
                    _ = v76
                }
            default:
                v51 = 60
            }
            _ = fi68
        }
    } else {
        var v77 uint = v2
        _ = v77
    }
    switch 0 {
    case 9223372036854775807:
        v78 := GenS9{F0: 0, F1: 769, F2: 4294967295}
        _ = v78
    case 42:
        v5 += v5
    default:
        for fi79 := 0; fi79 < 4; fi79++ {
            var v80 uint = 2
            _ = fi79
            _ = v80
        }
    }
    v1 = 42
    var v81 uint32 = 65535
    var v82 float32 = float32(((0.1 * -2.25) - -2.25))
    var v83 uint8 = 42
    for fi84 := 0; fi84 < 5; fi84++ {
        fmt.Println(fi84, fi84, v2)
        if false {
            v85 := GenS1{F0: (v3 * v3), F1: (false && false), F2: 32767}
            _ = v85
        }
        v86 := []int{}
        _ = fi84
        _ = v86
    }
    switch "" {
    case "hello":
        v3 = (^(v3))
    case "a":
        var v87 int8 = v51
        _ = v87
    case "":
        for fi88 := 0; fi88 < 7; fi88++ {
            for fi89 := 0; fi89 < 17; fi89++ {
                v90 := GenS10{F0: v81, F1: (v1 % ((42) | 1)), F2: ""}
                switch (fi88 | (fi88 - fi89)) {
                case -428:
                    if true {
                        if ((v3 | v3) > int32(v1)) {
                            var v91 uint8 = v83
                            fmt.Println(fi88)
                            switch fi88 {
                            case -9223372036854775808:
                                v3 = (v3 % ((v3) | 1))
                            case 4294967295:
                                var v92 bool = false
                                fmt.Println(v81)
                                switch "hello" {
                                case "hello":
                                    var v93 uint = (v2 & (v2 - v2))
                                    _ = v93
                                default:
                                    if v92 {
                                        v4 = (v4 & v4)
                                        switch 32767 {
                                        case 9223372036854775807:
                                            v92 = false
                                        }
                                        v91 = v83
                                    } else {
                                        if (("" < "") && (!(v82 == (v82 - float32(((((-2.25 - -2.25) * 1e10) - float64((3.14159 / ((v82 * (-2.25 / ((v82 - (v82 + 1.5)) / v82))) / (v82 - v82))))) / -2.25)))))) {
                                            switch "foo bar" {
                                            case "hello":
                                                var v94 int64 = (v1 / ((v1) | 1))
                                                var v95 uint64 = 1
                                                var v96 uint = 1
                                                _ = v94
                                                _ = v95
                                                _ = v96
                                            case "foo bar":
                                                v97 := []int{fi89, genFn0()}
                                                if (false || (v1 >= v1)) {
                                                    fmt.Println(v83, v2, v1)
                                                    var v98 uint8 = v83
                                                    v92 = v92
                                                    fmt.Println(v92, v91, v4)
                                                    v92 = v92
                                                    _ = v98
                                                }
                                                _ = v97
                                            case "\n\t":
                                                v2 += (v2 >> ((v2) & 63))
                                                if (!(v82 <= (3.14159 * float32(-2.25)))) {
                                                    if (!(false && (v92 && false))) {
                                                        var v99 int32 = v3
                                                        v100 := GenS7{F0: v81}
                                                        if v92 {
                                                            fmt.Println(v5, v5, v5)
                                                            v91 += v91
                                                            v101 := []int{int(v2), (^(fi89)), (fi88 / (((-(fi88))) | 1))}
                                                            _ = v101
                                                        } else {
                                                            v99 -= (v3 & 1)
                                                            var v102 int64 = v1
                                                            switch (fi88 + fi89) {
                                                            case -9223372036854775808:
                                                                v103 := []int{}
                                                                _ = v103
                                                            case -7:
                                                                v1 -= v102
                                                            case 0:
                                                                if (int32(v81) <= 0) {
                                                                    v104 := []int{(fi88 + (fi89 ^ (fi88 ^ fi88)))}
                                                                    if v92 {
                                                                        var v105 int32 = (^(v3))
                                                                        switch (fi89 % ((fi89) | 1)) {
                                                                        case 42:
                                                                            switch (-(fi88)) {
                                                                            case 4294967295:
                                                                                v81 = (v81 % (((^(v81))) | 1))
                                                                            }
                                                                        case 127:
                                                                            var v106 int32 = (v105 | 0)
                                                                            _ = v106
                                                                        default:
                                                                            switch "" {
                                                                            case "":
                                                                                v92 = ((229 != 65535) && (!(false || (!v92))))
                                                                            case "hello":
                                                                                v107 := GenS8{F0: -32768, F1: (^(v4)), F2: 2}
                                                                                _ = v107
                                                                            case "a":
                                                                                var v108 float32 = (v82 + 3.14159)
                                                                                _ = v108
                                                                            default:
                                                                                v102 = 1
                                                                            }
                                                                        }
                                                                        _ = v105
                                                                    }
                                                                    fmt.Println(fi89)
                                                                    v2 = v2
                                                                    fmt.Println(fi89, v1)
                                                                    _ = v104
                                                                } else {
                                                                    if false {
                                                                        var v109 uint64 = uint64(v83)
                                                                        switch genFn0() {
                                                                        case -1:
                                                                            v110 := GenS3{F0: 42, F1: genFn0(), F2: (v5 << (((v2 << ((uint(v81)) & 63))) & 63))}
                                                                            _ = v110
                                                                        }
                                                                        _ = v109
                                                                    }
                                                                }
                                                            default:
                                                                fmt.Println(v83, v4)
                                                            }
                                                            _ = v102
                                                        }
                                                        var v111 int64 = 0
                                                        var v112 uint64 = 2
                                                        _ = v100
                                                        _ = v111
                                                        _ = v112
                                                        _ = v99
                                                    } else {
                                                        var v113 float64 = (1.5 - 0.1)
                                                        _ = v113
                                                    }
                                                    v92 = (!false)
                                                    v114 := []int{(-(fi88)), (fi89 | genFn0())}
                                                    _ = v114
                                                }
                                            default:
                                                fmt.Println(v5, v83)
                                            }
                                            var v115 uint8 = v83
                                            if (v92 && v92) {
                                                var v116 uint = uint(v115)
                                                v1 = (v1 | int64(v3))
                                                _ = v116
                                            }
                                            v117 := GenS6{F0: 386}
                                            v118 := []int{fi89, fi89, 128, (fi88 ^ fi88)}
                                            switch fi88 {
                                            case -851:
                                                if v92 {
                                                    var v119 uint8 = v83
                                                    var v120 float64 = ((1e10 / 3.14159) + 0.1)
                                                    var v121 int = (fi88 & (-(fi88)))
                                                    _ = v119
                                                    _ = v120
                                                    _ = v121
                                                }
                                            case -801:
                                                v122 := GenS1{F0: (v3 ^ (v3 / (((^(v3))) | 1))), F1: true, F2: 701}
                                                _ = v122
                                            }
                                            var v123 float32 = 0.1
                                            _ = v115
                                            _ = v117
                                            _ = v118
                                            _ = v123
                                        }
                                        v124 := []int{(-(fi89)), 1}
                                        if true {
                                            v125 := GenS4{F0: (^(v83)), F1: v51, F2: (v2 | (v2 | v2))}
                                            _ = v125
                                        } else {
                                            var v126 float32 = 0.1
                                            var v127 string = ("foo bar" + "\n\t")
                                            v3 = -1
                                            _ = v126
                                            _ = v127
                                        }
                                        _ = v124
                                    }
                                }
                                _ = v92
                            }
                            if (203 < (v5 >> (((v2 >> ((v2) & 63))) & 63))) {
                                var v128 string = "a"
                                _ = v128
                            }
                            var v129 string = ((((("" + ("a" + "hello")) + "") + "\n\t") + ("hello" + "hello")) + ("" + ""))
                            v130 := GenS11{F0: 2147483647}
                            v131 := []int{183, fi88, (-(fi89))}
                            v91 = v91
                            var v132 uint16 = 255
                            _ = v129
                            _ = v130
                            _ = v131
                            _ = v132
                            _ = v91
                        } else {
                            fmt.Println(v5)
                        }
                    } else {
                        v133 := GenS1{F0: v3, F1: true, F2: 0}
                        v1 = v1
                        _ = v133
                    }
                }
                v134 := GenS1{F0: 2147483647, F1: (("foo bar" + ("" + ("a" + "a"))) <= (("\n\t" + (("\n\t" + (("a" + "") + "\n\t")) + ("hello" + "hello"))) + ("" + ""))), F2: uint16(v2)}
                _ = fi89
                _ = v134
                _ = v90
            }
            var v135 int64 = v1
            _ = fi88
            _ = v135
        }
    default:
        v1 = v1
    }
    switch ((("hello" + ("" + (("\n\t" + "\n\t") + ("\n\t" + "\n\t")))) + (("\n\t" + "foo bar") + ("foo bar" + ""))) + "\n\t") {
    case "\n\t":
        switch "foo bar" {
        case "foo bar":
            if (9223372036854775807 != 2147483647) {
                var v136 float32 = (3.14159 - (v82 * (0.0 * v82)))
                _ = v136
            } else {
                if false {
                    var v137 float32 = ((float32(float64(v82)) + v82) / v82)
                    var v138 uint8 = v83
                    v138 = (^(v138))
                    _ = v137
                    _ = v138
                } else {
                    v3 -= int32(v5)
                    switch 9223372036854775807 {
                    case 4294967295:
                        var v139 float64 = float64(1.5)
                        _ = v139
                    case 500:
                        var v140 uint32 = v81
                        _ = v140
                    case 9223372036854775807:
                        v141 := []int{-956}
                        _ = v141
                    }
                }
                v142 := GenS8{F0: 32767, F1: v4, F2: (v4 & v4)}
                switch "\n\t" {
                case "\n\t":
                    v143 := GenS12{F0: ("hello" + "foo bar"), F1: (v4 ^ v4), F2: ((v83 * uint8(v81)) != v83)}
                    _ = v143
                case "foo bar":
                    if false {
                        var v144 uint = uint(v83)
                        v145 := []int{-9223372036854775808, int(v83), 9223372036854775807, -130}
                        _ = v144
                        _ = v145
                    }
                default:
                    v82 = (0.0 + 3.14159)
                }
                _ = v142
            }
        }
    case "foo bar":
        v146 := []int{32768, genFn0(), int(v51)}
        _ = v146
    default:
        v147 := GenS13{F0: true, F1: (-(v51)), F2: (^(v83))}
        _ = v147
    }
    v82 -= float32(((1.5 / 1.5) - 1.5))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v81:", v81)
    fmt.Println("v82:", v82)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v81
    _ = v82
    _ = v83
}
