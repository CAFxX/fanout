package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint32; F1 uint32 }
type GenS5 struct { F0 uint32; F1 uint32; F2 string }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint8; F1 bool }
type GenS8 struct { F0 int16; F1 uint }
type GenS9 struct { F0 uint8; F1 int16 }
type GenS10 struct { F0 string; F1 int16 }
type GenS11 struct { F0 int32; F1 string; F2 int32 }
type GenS12 struct { F0 uint; F1 uint32 }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 int }
type GenS15 struct { F0 bool }
type GenS16 struct { F0 uint64; F1 int64; F2 bool }
type GenS17 struct { F0 int8; F1 uint8; F2 uint8 }
type GenS18 struct { F0 bool; F1 int16 }
type GenS19 struct { F0 string; F1 uint16; F2 int16 }
type GenS20 struct { F0 uint8; F1 int }
type GenS21 struct { F0 uint64; F1 int }
type GenS22 struct { F0 int64 }
func genFn0() string {
    v5 := GenS1{F0: -910}
    var v6 uint = 65535
    v7 := []int{32767}
    var v8 uint32 = 1
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return ("foo bar" + "foo bar")
}
func genFn1(p0 bool, p1 int) int32 {
    if ("foo bar" > "\n\t") {
        var v9 int8 = -32
        v10 := GenS1{F0: 937}
        var v11 string = "\n\t"
        _ = v10
        _ = v11
        _ = v9
    }
    fmt.Println(p0, p1)
    if (!(p0 && (p0 && (1.5 != 3.14159)))) {
        var v12 uint64 = 0
        fmt.Println(v12)
        var v13 string = "hello"
        _ = v12
        _ = v13
    } else {
        v14 := GenS1{F0: 32767}
        if ((0.0 + float64(0.0)) > -2.25) {
            if (!(!p0)) {
                v15 := []int{(p1 * p1), (^(p1))}
                v16 := GenS1{F0: 42}
                _ = v15
                _ = v16
            }
            p0 = p0
            for fi17 := 0; fi17 < 5; fi17++ {
                var v18 string = "a"
                v19 := []int{128, (p1 | (p1 << ((881) & 63))), (-(fi17))}
                _ = fi17
                _ = v18
                _ = v19
            }
        } else {
            p0 = (82 >= 42)
            v20 := GenS2{F0: true, F1: (("a" + "a") + "\n\t")}
            var v21 int64 = -688
            _ = v20
            _ = v21
        }
        _ = v14
    }
    var v22 int16 = 1000
    _ = p0
    _ = p1
    _ = v22
    return 42
}

func main() {
    var v1 bool = (true || true)
    var v2 bool = (-2.25 < (0.1 * 1.5))
    var v3 uint16 = 0
    var v4 int64 = 325
    switch (("foo bar" + "\n\t") + "") {
    case "a":
        for fi23 := 0; fi23 < 19; fi23++ {
            for fi24 := 0; fi24 < 6; fi24++ {
                fmt.Println(v2)
                v25 := GenS3{F0: false}
                _ = fi24
                _ = v25
            }
            for fi26 := 0; fi26 < 2; fi26++ {
                v27 := GenS2{F0: (!true), F1: ("\n\t" + "")}
                v28 := GenS4{F0: uint32(v3), F1: 259}
                v29 := GenS2{F0: v1, F1: "hello"}
                fmt.Println(fi23, v4)
                _ = fi26
                _ = v27
                _ = v28
                _ = v29
            }
            v30 := []int{}
            _ = fi23
            _ = v30
        }
    case "\n\t":
        var v31 float64 = 1e10
        _ = v31
    default:
        v4 = -100
    }
    var v32 float32 = float32(3.14159)
    var v33 int = 9223372036854775807
    switch (v33 / ((1) | 1)) {
    case 32768:
        for fi34 := 0; fi34 < 7; fi34++ {
            v35 := GenS5{F0: 716, F1: 2, F2: "hello"}
            if false {
                if (((!((42 < 100) && true)) && (1 > 127)) && (v3 != v3)) {
                    for fi36 := 0; fi36 < 8; fi36++ {
                        v1 = v2
                        if (v2 || (0 == (v3 / ((v3) | 1)))) {
                            var v37 uint8 = 42
                            var v38 string = ("\n\t" + ("foo bar" + (("a" + ("" + ("foo bar" + ("hello" + "")))) + "\n\t")))
                            var v39 uint8 = v37
                            _ = v37
                            _ = v38
                            _ = v39
                        }
                        _ = fi36
                    }
                    v3 = v3
                } else {
                    if true {
                        v40 := []int{fi34, (v33 & v33)}
                        _ = v40
                    } else {
                        for fi41 := 0; fi41 < 4; fi41++ {
                            if (18446744073709551615 < 434) {
                                switch "a" {
                                case "a":
                                    var v42 uint64 = 1
                                    _ = v42
                                default:
                                    fmt.Println(fi41, v32)
                                }
                                switch "" {
                                case "\n\t":
                                    var v43 uint = 157
                                    _ = v43
                                case "":
                                    if ((true || v2) || false) {
                                        if v1 {
                                            fmt.Println(v1)
                                            if ((v2 && (((v32 * v32) / v32) != 0.1)) || (((((v32 + -2.25) + v32) * v32) - v32) > 0.0)) {
                                                fmt.Println(v1, v32, v33)
                                                if ((v3 % ((65535) | 1)) > (v3 ^ v3)) {
                                                    v44 := GenS6{F0: 969}
                                                    if v1 {
                                                        if ((0.0 - float64(v32)) < ((1e10 / 1e10) + 0.1)) {
                                                            var v45 int32 = -375
                                                            v2 = false
                                                            v46 := []int{(v33 - fi41), (fi41 / ((fi41) | 1)), fi41}
                                                            var v47 uint8 = 42
                                                            if v1 {
                                                                fmt.Println(v47, v32)
                                                                v3 = v3
                                                            }
                                                            v2 = (v45 < v45)
                                                            _ = v45
                                                            _ = v46
                                                            _ = v47
                                                        } else {
                                                            fmt.Println(v33)
                                                            v2 = v1
                                                            v48 := GenS3{F0: true}
                                                            _ = v48
                                                        }
                                                        v2 = (0 > int16(v3))
                                                        var v49 uint32 = 2147483647
                                                        v4 -= 42
                                                        _ = v49
                                                    }
                                                    var v50 int8 = 42
                                                    _ = v44
                                                    _ = v50
                                                } else {
                                                    var v51 uint8 = 37
                                                    if (((-(v4)) > (v4 * 0)) && (623 < 2)) {
                                                        if v2 {
                                                            var v52 uint32 = 202
                                                            if v1 {
                                                                var v53 int8 = 42
                                                                v32 = ((((float32(0.1) * float32((1e10 - 1e10))) / 1.0) * float32(0.1)) / v32)
                                                                _ = v53
                                                            }
                                                            _ = v52
                                                        }
                                                    }
                                                    v54 := GenS2{F0: v1, F1: "hello"}
                                                    _ = v51
                                                    _ = v54
                                                }
                                                v55 := []int{(fi41 ^ fi34), (fi41 - (v33 >> ((1) & 63))), v33}
                                                var v56 uint8 = 255
                                                fmt.Println(v2, v4)
                                                v57 := GenS2{F0: false, F1: ("a" + "foo bar")}
                                                if len(v55) >= 2 {
                                                    v55 = v55[1:len(v55)]
                                                }
                                                var v58 uint8 = (v56 - 228)
                                                if (!true) {
                                                    var v59 bool = ((0 == 2147483647) || (!true))
                                                    v60 := GenS3{F0: (((!(42 < 42)) && v1) || (1 < uint32(v58)))}
                                                    _ = v59
                                                    _ = v60
                                                } else {
                                                    v2 = v1
                                                }
                                                _ = v55
                                                _ = v56
                                                _ = v57
                                                _ = v58
                                            }
                                        } else {
                                            v32 = v32
                                        }
                                        if (-1 < 2147483647) {
                                            v2 = v2
                                            v3 += (v3 % ((v3) | 1))
                                        }
                                    }
                                case "a":
                                    fmt.Println(v32)
                                default:
                                    var v61 int = (-(fi34))
                                    _ = v61
                                }
                            } else {
                                switch 32767 {
                                case 42:
                                    v4 -= v4
                                case -2147483648:
                                    var v62 int8 = -124
                                    _ = v62
                                case 4294967295:
                                    v63 := GenS3{F0: v2}
                                    _ = v63
                                }
                                if false {
                                    var v64 int16 = -1
                                    v65 := GenS4{F0: 65535, F1: 816}
                                    switch fi41 {
                                    case 128:
                                        v33 = (fi34 % (((^(fi41))) | 1))
                                    case 127:
                                        fmt.Println(v3)
                                    default:
                                        switch v33 {
                                        case -1:
                                            var v66 uint = 845
                                            _ = v66
                                        default:
                                            v67 := GenS2{F0: ((v64 >> ((255) & 63)) >= v64), F1: "hello"}
                                            _ = v67
                                        }
                                    }
                                    _ = v64
                                    _ = v65
                                }
                                if (!v1) {
                                    v68 := []int{(fi41 << ((14) & 63)), (v33 - fi34), (v33 - (fi34 + fi34))}
                                    var v69 string = "hello"
                                    _ = v68
                                    _ = v69
                                }
                                switch "a" {
                                case "foo bar":
                                    v32 -= 1.5
                                case "\n\t":
                                    fmt.Println(v4, v2)
                                case "hello":
                                    switch ("" + ("" + "")) {
                                    case "hello":
                                        if v2 {
                                            v4 = (v4 | 165)
                                            var v70 float64 = -2.25
                                            _ = v70
                                        } else {
                                            switch (fi34 - -931) {
                                            case -2147483648:
                                                v1 = ("hello" >= "\n\t")
                                            default:
                                                v4 += -1
                                            }
                                            var v71 int64 = v4
                                            _ = v71
                                        }
                                    case "foo bar":
                                        switch "\n\t" {
                                        case "hello":
                                            v33 = fi41
                                        case "":
                                            fmt.Println(v4, fi34, v2)
                                        case "\n\t":
                                            fmt.Println(fi41, v32)
                                        default:
                                            var v72 int = fi34
                                            _ = v72
                                        }
                                    case "\n\t":
                                        if true {
                                            if v1 {
                                                if (v1 || ("foo bar" < "a")) {
                                                    var v73 uint16 = v3
                                                    v74 := []int{fi34, (fi41 << ((347) & 63))}
                                                    if (!((v2 && v1) || v2)) {
                                                        if (620 <= 2) {
                                                            switch -1 {
                                                            case 128:
                                                                var v75 int16 = -32768
                                                                v76 := GenS1{F0: (v75 + v75)}
                                                                _ = v75
                                                                _ = v76
                                                            default:
                                                                v77 := GenS7{F0: 0, F1: v1}
                                                                _ = v77
                                                            }
                                                        } else {
                                                            fmt.Println(v73, v33, v4)
                                                            var v78 float32 = v32
                                                            _ = v78
                                                        }
                                                        fmt.Println(v33)
                                                    } else {
                                                        v79 := GenS8{F0: -282, F1: 259}
                                                        switch (fi34 % ((32768) | 1)) {
                                                        case -383:
                                                            var v80 int8 = -128
                                                            _ = v80
                                                        case -9223372036854775808:
                                                            if v1 {
                                                                switch "foo bar" {
                                                                case "hello":
                                                                    var v81 int = fi34
                                                                    _ = v81
                                                                case "a":
                                                                    var v82 int64 = (v4 / ((v4) | 1))
                                                                    _ = v82
                                                                }
                                                            }
                                                        default:
                                                            var v83 int64 = v4
                                                            _ = v83
                                                        }
                                                        var v84 uint = 65535
                                                        v74 = append(v74, 127)
                                                        _ = v79
                                                        _ = v84
                                                    }
                                                    _ = v73
                                                    _ = v74
                                                } else {
                                                    v85 := GenS9{F0: 2, F1: 970}
                                                    _ = v85
                                                }
                                                v32 += v32
                                                v86 := []int{(fi41 >> ((0) & 63)), v33, fi34, 1}
                                                var v87 int8 = -1
                                                if v2 {
                                                    fmt.Println(v1, v33, v32)
                                                    if len(v86) >= 2 {
                                                        v86 = v86[1:len(v86)]
                                                    }
                                                }
                                                _ = v86
                                                _ = v87
                                            }
                                            fmt.Println(v4)
                                            if (v32 >= (v32 / 3.14159)) {
                                                v88 := GenS2{F0: v1, F1: "foo bar"}
                                                v89 := GenS10{F0: (("hello" + "\n\t") + ("" + "a")), F1: -32768}
                                                _ = v88
                                                _ = v89
                                            } else {
                                                v90 := GenS9{F0: 127, F1: 1}
                                                var v91 int64 = 0
                                                _ = v90
                                                _ = v91
                                            }
                                        }
                                    }
                                }
                                var v92 int32 = genFn1(((v4 << ((42) & 63)) == v4), (fi41 & (fi41 / ((int(v4)) | 1))))
                                v93 := GenS9{F0: 1, F1: 42}
                                _ = v92
                                _ = v93
                            }
                            v2 = true
                            _ = fi41
                        }
                        var v94 uint32 = 2147483648
                        _ = v94
                    }
                    var v95 uint16 = (^(v3))
                    for fi96 := 0; fi96 < 11; fi96++ {
                        var v97 uint8 = 127
                        _ = fi96
                        _ = v97
                    }
                    _ = v95
                }
                v33 = -584
            } else {
                if (931 >= 2) {
                    switch int(v4) {
                    case -193:
                        v4 -= v4
                    default:
                        switch "hello" {
                        case "hello":
                            v98 := GenS11{F0: 0, F1: ("foo bar" + ""), F2: genFn1(v2, (fi34 * fi34))}
                            _ = v98
                        case "":
                            var v99 int32 = 1
                            _ = v99
                        default:
                            var v100 float64 = (((-2.25 - 0.0) - 1.5) * 0.0)
                            _ = v100
                        }
                    }
                } else {
                    v3 -= v3
                    v101 := []int{127}
                    var v102 bool = ((float32(-2.25) / (float32(0.1) - v32)) <= v32)
                    _ = v101
                    _ = v102
                }
                v103 := []int{}
                if (genFn1((((int8(v3) == -88) || ((212 >= 128) || (!true))) || true), (-(fi34))) == -1) {
                    var v104 int16 = 586
                    _ = v104
                }
                _ = v103
            }
            if (2 != 867) {
                var v105 uint8 = 0
                switch 32767 {
                case 32767:
                    var v106 bool = true
                    _ = v106
                default:
                    var v107 int8 = -104
                    _ = v107
                }
                var v108 string = (("" + ("foo bar" + "hello")) + "a")
                _ = v105
                _ = v108
            }
            _ = fi34
            _ = v35
        }
    case 128:
        for fi109 := 0; fi109 < 12; fi109++ {
            var v110 uint8 = 76
            fmt.Println(v1, v4, v33)
            for fi111 := 0; fi111 < 4; fi111++ {
                switch 127 {
                case 127:
                    v2 = v1
                case 4294967295:
                    v1 = (((v32 - (1.5 / v32)) * (v32 - v32)) <= v32)
                default:
                    var v112 uint16 = (v3 ^ (v3 % ((0) | 1)))
                    var v113 float64 = -2.25
                    if v2 {
                        var v114 int8 = int8(v112)
                        switch fi109 {
                        case 2147483647:
                            switch (fi111 | fi109) {
                            case -9223372036854775808:
                                if (128 != v110) {
                                    switch (fi111 - 32768) {
                                    case -564:
                                        switch "" {
                                        case "hello":
                                            if (!((v4 - (-(v4))) > v4)) {
                                                v115 := GenS12{F0: 2, F1: 753}
                                                v116 := GenS13{F0: -2147483648}
                                                if v1 {
                                                    v32 = v32
                                                }
                                                _ = v115
                                                _ = v116
                                            }
                                        case "":
                                            if (false || (v114 == v114)) {
                                                v1 = v1
                                            } else {
                                                v117 := GenS14{F0: int(v4)}
                                                v32 = v32
                                                _ = v117
                                            }
                                        }
                                    case -2147483648:
                                        switch "a" {
                                        case "":
                                            if ((1 == 42) || v1) {
                                                var v118 string = ("a" + "\n\t")
                                                fmt.Println(v4)
                                                switch ((v118 + "hello") + ((v118 + "a") + ("hello" + v118))) {
                                                case "hello":
                                                    var v119 uint64 = 291
                                                    _ = v119
                                                case "a":
                                                    var v120 uint64 = 2
                                                    _ = v120
                                                default:
                                                    v121 := GenS15{F0: (!(v113 == v113))}
                                                    _ = v121
                                                }
                                                _ = v118
                                            } else {
                                                v122 := []int{-172, (^(fi109)), 668}
                                                _ = v122
                                            }
                                        default:
                                            fmt.Println(v3, v114, v2)
                                        }
                                        v113 = 1.5
                                        switch (fi109 * 264) {
                                        case 32768:
                                            var v123 uint = 142
                                            _ = v123
                                        case -9223372036854775808:
                                            var v124 uint32 = 2147483648
                                            fmt.Println(v110, v32)
                                            if true {
                                                v125 := []int{(fi109 * 128), fi109}
                                                v3 = v3
                                                var v126 int8 = (v114 * v114)
                                                _ = v125
                                                _ = v126
                                            }
                                            _ = v124
                                        }
                                    }
                                }
                                var v127 float32 = -2.25
                                var v128 float32 = float32((((v113 + v113) + (v113 / ((0.0 + v113) + (1.5 - 0.1)))) * 3.14159))
                                var v129 int64 = v4
                                _ = v127
                                _ = v128
                                _ = v129
                            case -36:
                                var v130 bool = v2
                                _ = v130
                            case 1:
                                v110 = 237
                            default:
                                var v131 uint16 = (^(v3))
                                _ = v131
                            }
                        default:
                            fmt.Println(v112, v110)
                        }
                        _ = v114
                    } else {
                        v132 := []int{(fi109 | (-(v33)))}
                        if false {
                            if (3.14159 > v32) {
                                switch "\n\t" {
                                case "\n\t":
                                    v133 := GenS16{F0: 2, F1: v4, F2: false}
                                    _ = v133
                                default:
                                    if (2 >= uint(v3)) {
                                        var v134 uint16 = 42
                                        if v2 {
                                            v3 -= (v112 << ((18446744073709551615) & 63))
                                            var v135 uint16 = v112
                                            fmt.Println(v110, v1, v135)
                                            _ = v135
                                        } else {
                                            if (889 != 65535) {
                                                v132 = append(v132, (v33 * 9223372036854775807))
                                            }
                                        }
                                        if (-32768 == -1) {
                                            if (-2.25 > 1.5) {
                                                fmt.Println(v113, v4, v4)
                                                v4 = v4
                                                var v136 uint16 = v134
                                                fmt.Println(v136)
                                                v137 := GenS17{F0: 105, F1: (^(v110)), F2: (v110 & (v110 / ((2) | 1)))}
                                                _ = v136
                                                _ = v137
                                            }
                                            if v1 {
                                                v134 = v134
                                                var v138 uint16 = v134
                                                _ = v138
                                            }
                                            v113 = 1.5
                                        } else {
                                            switch 813 {
                                            case 128:
                                                switch v33 {
                                                case 128:
                                                    v139 := GenS18{F0: v1, F1: 303}
                                                    _ = v139
                                                }
                                                v140 := GenS7{F0: v110, F1: v1}
                                                _ = v140
                                            case 42:
                                                fmt.Println(v4, v32)
                                                var v141 int8 = -7
                                                var v142 float32 = (0.0 / v32)
                                                switch (("" + ("" + "")) + "foo bar") {
                                                case "":
                                                    v143 := GenS16{F0: 4294967295, F1: -580, F2: v1}
                                                    _ = v143
                                                case "foo bar":
                                                    v2 = v2
                                                case "a":
                                                    v144 := GenS1{F0: 0}
                                                    _ = v144
                                                }
                                                if (v1 || (!false)) {
                                                    var v145 int64 = v4
                                                    _ = v145
                                                }
                                                v32 = float32(((1e10 - 3.14159) * -2.25))
                                                _ = v141
                                                _ = v142
                                            default:
                                                var v146 uint32 = 280
                                                _ = v146
                                            }
                                            if ((v4 | v4) <= (v4 ^ -996)) {
                                                switch ("\n\t" + "") {
                                                case "foo bar":
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        fmt.Println(v2)
                                                    }
                                                case "\n\t":
                                                    v33 += (v33 + -564)
                                                case "hello":
                                                    var v147 uint16 = v134
                                                    _ = v147
                                                }
                                            } else {
                                                v134 -= 32767
                                                v110 = v110
                                                fmt.Println(fi109)
                                            }
                                            var v148 uint32 = 760
                                            fmt.Println(v113, v32)
                                            var v149 uint8 = (^(v110))
                                            v149 += v149
                                            v132 = append(v132, -477)
                                            _ = v148
                                            _ = v149
                                        }
                                        _ = v134
                                    } else {
                                        v150 := GenS19{F0: "", F1: 935, F2: 1}
                                        if false {
                                            v1 = v2
                                        } else {
                                            v4 = v4
                                        }
                                        _ = v150
                                    }
                                }
                                _ = len(v132)
                                var v151 uint16 = (v3 | v112)
                                _ = v151
                            } else {
                                v152 := GenS7{F0: (v110 * v110), F1: (false || v2)}
                                v1 = (!((-2.25 > ((v113 / (v113 - float64(v32))) + v113)) || ("foo bar" < (("foo bar" + "") + (("a" + ("" + "a")) + ("foo bar" + ""))))))
                                _ = v152
                            }
                            v153 := GenS4{F0: 4294967295, F1: 791}
                            var v154 uint64 = 42
                            _ = v153
                            _ = v154
                        } else {
                            if len(v132) > 0 {
                                v132[0] = fi109
                            }
                            v2 = ((uint(v110) > 4294967295) || false)
                        }
                        _ = v132
                    }
                    if v1 {
                        v155 := []int{v33}
                        v156 := GenS20{F0: (v110 | 42), F1: 1}
                        _ = v155
                        _ = v156
                    } else {
                        v2 = v2
                    }
                    var v157 int8 = 42
                    var v158 bool = (2 > 65535)
                    if ((v4 & v4) <= (-(v4))) {
                        v4 = v4
                    }
                    fmt.Println(v33, v157, v157)
                    _ = v112
                    _ = v113
                    _ = v157
                    _ = v158
                }
                _ = fi111
            }
            _ = fi109
            _ = v110
        }
    case -2147483648:
        v159 := []int{v33, (-(v33)), (v33 | 1)}
        _ = v159
    }
    v32 = v32
    v1 = (177 == 127)
    v3 = v3
    v160 := GenS21{F0: 2, F1: (v33 >> ((37) & 63))}
    if ("" < "hello") {
        for fi161 := 0; fi161 < 13; fi161++ {
            for fi162 := 0; fi162 < 11; fi162++ {
                v4 += int64(fi161)
                v163 := []int{(fi161 | fi162)}
                v33 -= (v33 / ((127) | 1))
                _ = fi162
                _ = v163
            }
            v164 := GenS7{F0: 0, F1: v2}
            var v165 uint64 = 18446744073709551615
            _ = fi161
            _ = v164
            _ = v165
        }
        switch (-(v33)) {
        case 32768:
            v166 := GenS10{F0: ("\n\t" + "\n\t"), F1: 100}
            _ = v166
        case 32767:
            switch (v33 + -9223372036854775808) {
            case 32768:
                for fi167 := 0; fi167 < 10; fi167++ {
                    v1 = true
                    if (1000 < -1) {
                        switch ("" + "") {
                        case "\n\t":
                            var v168 float64 = float64(v32)
                            _ = v168
                        default:
                            if (v1 && ((fi167 << ((255) & 63)) >= -826)) {
                                var v169 int8 = 42
                                fmt.Println(v169, v33, v33)
                                v169 = 127
                                _ = v169
                            }
                        }
                        v170 := []int{}
                        v4 = (v4 >> ((42) & 63))
                        _ = v170
                    } else {
                        for fi171 := 0; fi171 < 3; fi171++ {
                            fmt.Println(fi171, v4, v2)
                            fmt.Println(v33, v1)
                            v3 -= v3
                            _ = fi171
                        }
                    }
                    _ = fi167
                }
            case 2147483647:
                v4 -= (v4 + (v4 + (v4 - v4)))
            }
        case 9223372036854775807:
            var v172 string = (("foo bar" + "hello") + "foo bar")
            _ = v172
        }
    }
    var v173 int64 = 772
    for fi174 := 0; fi174 < 2; fi174++ {
        for fi175 := 0; fi175 < 1; fi175++ {
            if true {
                v176 := []int{(fi174 + v33), (v33 ^ (fi174 * 0)), 1}
                _ = v176
            } else {
                v2 = (2147483647 < 42)
                v3 += 42
                var v177 uint64 = 2
                _ = v177
            }
            if (uint64(v3) != uint64(v3)) {
                switch fi175 {
                case -436:
                    var v178 uint = 1
                    _ = v178
                case -332:
                    v179 := GenS9{F0: 102, F1: 1}
                    _ = v179
                case 2147483647:
                    v180 := GenS4{F0: 4294967295, F1: 65535}
                    _ = v180
                }
                var v181 uint = 227
                _ = v181
            } else {
                fmt.Println(v33)
                var v182 float64 = 0.1
                _ = v182
            }
            var v183 uint = 735
            _ = fi175
            _ = v183
        }
        v184 := GenS22{F0: v4}
        _ = fi174
        _ = v184
    }
    fmt.Println("v1:", v1)
    fmt.Println("v173:", v173)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v160
    _ = v173
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
}
