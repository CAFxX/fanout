package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint16; F1 string; F2 uint64 }
type GenS6 struct { F0 string; F1 uint32 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int64; F1 uint32; F2 uint64 }
type GenS9 struct { F0 uint8; F1 uint }
type GenS10 struct { F0 uint; F1 uint16; F2 uint8 }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 int8; F1 int32; F2 int32 }
type GenS13 struct { F0 uint16; F1 uint8; F2 uint64 }
type GenS14 struct { F0 string; F1 uint32; F2 uint64 }
type GenS15 struct { F0 int64; F1 int64; F2 bool }
type GenS16 struct { F0 int64 }
type GenS17 struct { F0 int64; F1 int32; F2 int64 }
type GenS18 struct { F0 string; F1 uint64; F2 uint }
func genFn0(p0 int16, p1 int64) int64 {
    v5 := GenS1{F0: true, F1: -1}
    switch 408 {
    case -9223372036854775808:
        for fi6 := 0; fi6 < 18; fi6++ {
            for fi7 := 0; fi7 < 13; fi7++ {
                p1 -= p1
                _ = fi7
            }
            _ = fi6
        }
    case -1:
        for fi8 := 0; fi8 < 16; fi8++ {
            var v9 uint64 = 2
            for fi10 := 0; fi10 < 9; fi10++ {
                var v11 float32 = 1.5
                v9 = v9
                if true {
                    var v12 bool = ((!false) && true)
                    _ = v12
                }
                _ = fi10
                _ = v11
            }
            switch fi8 {
            case 1:
                for fi13 := 0; fi13 < 16; fi13++ {
                    p0 -= 1
                    p0 -= p0
                    _ = fi13
                }
            case 4294967295:
                v9 += 42
            default:
                if (26 >= 128) {
                    if true {
                        fmt.Println(v9)
                        fmt.Println(fi8, fi8, v9)
                    }
                    if (int8(fi8) <= 0) {
                        if true {
                            v14 := GenS2{F0: false, F1: (p0 % ((p0) | 1))}
                            _ = v14
                        } else {
                            for fi15 := 0; fi15 < 6; fi15++ {
                                v16 := GenS3{F0: (p1 - -9223372036854775808)}
                                v17 := GenS2{F0: false, F1: p0}
                                v18 := []int{fi8, 9223372036854775807, 298}
                                fmt.Println(p1)
                                _ = fi15
                                _ = v16
                                _ = v17
                                _ = v18
                            }
                        }
                        var v19 int = fi8
                        v19 = fi8
                        _ = v19
                    } else {
                        fmt.Println(fi8)
                        var v20 int = 85
                        _ = v20
                    }
                } else {
                    fmt.Println(fi8, v9)
                }
            }
            _ = fi8
            _ = v9
        }
    case 587:
        p0 -= p0
    }
    v21 := []int{}
    if (false || (!false)) {
        fmt.Println(p0, p0)
    } else {
        v21 = append(v21, 4294967295)
    }
    for fi22 := 0; fi22 < 1; fi22++ {
        if (746 != 32767) {
            if len(v21) > 0 {
                v21[0] = fi22
            }
            p0 += p0
            var v23 uint32 = 65535
            _ = v23
        } else {
            var v24 uint16 = 2
            if false {
                var v25 uint16 = (^(v24))
                switch (((((("hello" + "a") + "") + "\n\t") + "hello") + ("\n\t" + ("foo bar" + ""))) + "\n\t") {
                case "\n\t":
                    if (!false) {
                        var v26 int16 = p0
                        _ = v26
                    }
                case "a":
                    for fi27 := 0; fi27 < 5; fi27++ {
                        if false {
                            p0 += (^(p0))
                            _ = len(v21)
                        }
                        _ = fi27
                    }
                case "":
                    var v28 int64 = p1
                    _ = v28
                }
                for fi29 := 0; fi29 < 5; fi29++ {
                    fmt.Println(p0, p1)
                    fmt.Println(fi22, p1)
                    v25 -= 255
                    _ = fi29
                }
                _ = v25
            } else {
                var v30 int64 = (p1 ^ (p1 & (^(p1))))
                _ = v30
            }
            _ = v24
        }
        _ = len(v21)
        for fi31 := 0; fi31 < 6; fi31++ {
            if (!true) {
                var v32 uint32 = 2147483647
                _ = v32
            }
            if (!(!(!false))) {
                var v33 int16 = p0
                _ = v33
            }
            v34 := GenS4{F0: 127}
            var v35 int32 = 0
            p1 = 1
            _ = fi31
            _ = v34
            _ = v35
        }
        _ = fi22
    }
    _ = p0
    _ = p1
    _ = v21
    _ = v5
    return (p1 >> ((255) & 63))
}
func genFn1(p0 uint16, p1 int) bool {
    fmt.Println(p0, p1)
    v36 := []int{p1}
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = v36
    return false
}

func main() {
    var v1 uint64 = 664
    var v2 string = ("a" + ((((("" + "foo bar") + "") + "hello") + ("a" + ("a" + ("hello" + ("\n\t" + "hello"))))) + "\n\t"))
    var v3 int64 = 0
    v4 := make(map[string]bool)
    for fi37 := 0; fi37 < 2; fi37++ {
        v3 = (v3 / ((1) | 1))
        v1 = (^(v1))
        _ = fi37
    }
    if (!(0 > -600)) {
        if false {
            delete(v4, ("a" + (v2 + v2)))
            if false {
                v3 += v3
                for fi38 := 0; fi38 < 18; fi38++ {
                    var v39 int8 = -90
                    _ = fi38
                    _ = v39
                }
            }
            switch 4294967295 {
            case 4294967295:
                var v40 uint = 2
                _ = v40
            case -1:
                switch "hello" {
                case "a":
                    for fi41 := 0; fi41 < 5; fi41++ {
                        if false {
                            switch (fi41 * fi41) {
                            case -712:
                                v2 = v2
                            case 2147483647:
                                fmt.Println(v1, v3)
                            case 9223372036854775807:
                                if (false && ((314 < 0) && false)) {
                                    if false {
                                        delete(v4, v2)
                                        var v42 float32 = 0.0
                                        v43 := []int{}
                                        _ = v42
                                        _ = v43
                                    }
                                }
                            default:
                                v44 := GenS5{F0: 636, F1: v2, F2: v1}
                                _ = v44
                            }
                            for fi45 := 0; fi45 < 4; fi45++ {
                                delete(v4, v2)
                                var v46 bool = (!true)
                                _ = fi45
                                _ = v46
                            }
                        } else {
                            fmt.Println(v2, v2)
                            v1 = v1
                        }
                        if false {
                            var v47 uint = 42
                            _ = v47
                        }
                        var v48 int32 = 2147483647
                        _ = fi41
                        _ = v48
                    }
                case "foo bar":
                    v3 += 1
                }
            default:
                for fi49 := 0; fi49 < 12; fi49++ {
                    if false {
                        v50 := GenS4{F0: 185}
                        for fi51 := 0; fi51 < 16; fi51++ {
                            var v52 uint = 287
                            _ = fi51
                            _ = v52
                        }
                        _ = v50
                    } else {
                        for fi53 := 0; fi53 < 7; fi53++ {
                            var v54 uint = 2
                            _ = fi53
                            _ = v54
                        }
                        delete(v4, "hello")
                        v55 := []int{(-(fi49)), (fi49 / ((42) | 1))}
                        _ = v55
                    }
                    v56 := GenS4{F0: 127}
                    v2 = v2
                    _ = fi49
                    _ = v56
                }
            }
        } else {
            var v57 uint64 = 4294967295
            var v58 int32 = -827
            var v59 int8 = 100
            _ = v57
            _ = v58
            _ = v59
        }
        switch (((v2 + "foo bar") + "foo bar") + ("hello" + v2)) {
        case "a":
            v2 = ((v2 + ((v2 + v2) + v2)) + ((v2 + "\n\t") + v2))
        case "hello":
            v60 := []int{9223372036854775807, 42, -138, -807}
            _ = v60
        default:
            for fi61 := 0; fi61 < 15; fi61++ {
                switch -1 {
                case -9223372036854775808:
                    v2 = v2
                case -637:
                    v3 -= v3
                }
                for fi62 := 0; fi62 < 9; fi62++ {
                    v3 += (v3 % ((int64(v1)) | 1))
                    var v63 float32 = 0.0
                    fmt.Println(v2)
                    v64 := []int{}
                    var v65 float64 = 0.0
                    _ = fi62
                    _ = v63
                    _ = v64
                    _ = v65
                }
                _ = fi61
            }
        }
        if (false || true) {
            for fi66 := 0; fi66 < 13; fi66++ {
                for fi67 := 0; fi67 < 16; fi67++ {
                    var v68 uint64 = 4294967295
                    v2 = v2
                    if (v68 == v68) {
                        v2 = "hello"
                        v68 = (v68 ^ v1)
                        v1 = v1
                        v69 := []int{(fi67 + (fi66 + (fi66 << ((0) & 63)))), fi67, (fi66 * (fi66 * (fi67 & (fi66 ^ fi67)))), (^(fi66))}
                        _ = v69
                    }
                    _ = fi67
                    _ = v68
                }
                if ((1e10 < (1.5 - (((1e10 + 0.1) / 0.1) / 1.5))) || (65535 != 65535)) {
                    if false {
                        v70 := []int{fi66, fi66, fi66}
                        _ = v70
                    }
                }
                if (18446744073709551615 < 18446744073709551615) {
                    v71 := []int{fi66, fi66, (fi66 | -2147483648), (-(fi66))}
                    _ = v71
                }
                _ = fi66
            }
            if false {
                switch 9223372036854775807 {
                case 185:
                    for fi72 := 0; fi72 < 11; fi72++ {
                        var v73 int64 = v3
                        var v74 float32 = float32((1.5 / ((((3.14159 / 1e10) - 0.1) + 0.1) + 1e10)))
                        v75 := GenS1{F0: true, F1: 109}
                        _ = fi72
                        _ = v73
                        _ = v74
                        _ = v75
                    }
                default:
                    v76 := []int{-114, 2147483647, 42}
                    _ = v76
                }
                v77 := GenS2{F0: false, F1: 525}
                _ = v77
            } else {
                v78 := []int{}
                _ = v78
            }
        } else {
            if true {
                var v79 float32 = (0.1 - -2.25)
                var v80 int16 = -338
                _ = v79
                _ = v80
            } else {
                v3 = v3
                switch 32767 {
                case -16:
                    v81 := []int{128, 42, 1}
                    _ = v81
                case 32767:
                    v1 = v1
                default:
                    var v82 float64 = -2.25
                    _ = v82
                }
            }
        }
    }
    v3 -= 1
    fmt.Println(v1, v2, v1)
    var v83 uint64 = (v1 ^ v1)
    for fi84 := 0; fi84 < 3; fi84++ {
        for fi85 := 0; fi85 < 8; fi85++ {
            v1 += v1
            v2 = v2
            v86 := GenS3{F0: (v3 ^ v3)}
            v87 := GenS6{F0: "", F1: 42}
            fmt.Println(fi85, fi85)
            var v88 bool = (42 >= 2147483648)
            _ = fi85
            _ = v86
            _ = v87
            _ = v88
        }
        v89 := []int{fi84, fi84, fi84, fi84}
        fmt.Println(v83)
        _ = fi84
        _ = v89
    }
    v90 := GenS2{F0: false, F1: 566}
    v91 := []int{}
    if ((1 < 115) && (880 <= 1)) {
        v2 = ((v2 + ("" + "")) + v2)
    } else {
        for fi92 := 0; fi92 < 18; fi92++ {
            for fi93 := 0; fi93 < 12; fi93++ {
                var v94 int64 = 42
                switch "a" {
                case "hello":
                    if (!false) {
                        if false {
                            if true {
                                switch v2 {
                                case "foo bar":
                                    v95 := v4["a"]
                                    _ = v95
                                case "a":
                                    var v96 int8 = -122
                                    _ = v96
                                default:
                                    if (337 == 0) {
                                        if len(v91) > 0 {
                                            v91[0] = fi93
                                        }
                                        fmt.Println(v3)
                                        fmt.Println(v2)
                                    } else {
                                        if false {
                                            var v97 float32 = 1e10
                                            _ = v97
                                        }
                                        if len(v91) >= 2 {
                                            v91 = v91[1:len(v91)]
                                        }
                                        var v98 float64 = 0.0
                                        fmt.Println(v83, v98, fi92)
                                        _ = v98
                                    }
                                }
                                v1 -= 1
                                if (!false) {
                                    switch (fi93 * (-(fi92))) {
                                    case 32767:
                                        v99 := GenS7{F0: 0}
                                        _ = v99
                                    case -9223372036854775808:
                                        var v100 uint64 = v83
                                        _ = v100
                                    case 4294967295:
                                        v4[""] = (0 != 2147483647)
                                    }
                                    v101 := v4["foo bar"]
                                    _ = v101
                                } else {
                                    if (v3 != v3) {
                                        fmt.Println(v3, v94)
                                        var v102 uint16 = 32767
                                        if (0 != 229) {
                                            v103 := GenS4{F0: 30}
                                            _ = v103
                                        }
                                        _ = v102
                                    } else {
                                        var v104 string = (((v2 + ("\n\t" + (v2 + ""))) + ("foo bar" + ("" + (v2 + "foo bar")))) + v2)
                                        var v105 float64 = float64(float32(float64(float32(float64(1.5)))))
                                        v91 = append(v91, (-(fi92)))
                                        _ = v104
                                        _ = v105
                                    }
                                    v94 -= -738
                                    v106 := GenS8{F0: (v3 ^ v94), F1: 882, F2: (v83 | 2)}
                                    _ = v106
                                }
                            }
                            if false {
                                if true {
                                    if true {
                                        fmt.Println(fi93, v2, v2)
                                        switch ("" + (v2 + v2)) {
                                        case "hello":
                                            var v107 uint32 = 456
                                            _ = v107
                                        case "\n\t":
                                            if true {
                                                if (true && false) {
                                                    fmt.Println(fi93, v94)
                                                    v2 = v2
                                                    switch v2 {
                                                    case "":
                                                        v1 += v83
                                                    default:
                                                        fmt.Println(fi93)
                                                    }
                                                } else {
                                                    var v108 bool = (-1 > 1)
                                                    _ = v108
                                                }
                                                v94 = (v94 & 197)
                                                if (1 <= 42) {
                                                    switch "foo bar" {
                                                    case "hello":
                                                        if (!false) {
                                                            v1 += (^(v83))
                                                            fmt.Println(v83)
                                                        }
                                                    case "":
                                                        v91 = append(v91, (fi92 + (fi92 | (fi93 | -9223372036854775808))))
                                                    case "foo bar":
                                                        if (!(!true)) {
                                                            var v109 float32 = 3.14159
                                                            var v110 int32 = 1
                                                            if (((!true) || true) || (1 <= 127)) {
                                                                var v111 bool = (fi93 < 960)
                                                                v112 := GenS3{F0: v94}
                                                                _ = v111
                                                                _ = v112
                                                            }
                                                            _ = v109
                                                            _ = v110
                                                        }
                                                    default:
                                                        delete(v4, ("hello" + "hello"))
                                                    }
                                                } else {
                                                    fmt.Println(v94, v1)
                                                }
                                            } else {
                                                v113 := GenS9{F0: 255, F1: 266}
                                                _ = v113
                                            }
                                        case "a":
                                            switch v2 {
                                            case "a":
                                                v114 := v4[("a" + (v2 + v2))]
                                                _ = v114
                                            case "\n\t":
                                                var v115 uint = 1
                                                _ = v115
                                            case "":
                                                var v116 uint8 = 255
                                                _ = v116
                                            }
                                        }
                                    }
                                    v94 += (v94 & v3)
                                    fmt.Println(fi93, v94)
                                    fmt.Println(v83)
                                    v94 += -9223372036854775808
                                    v4[v2] = false
                                } else {
                                    var v117 uint32 = 496
                                    v118 := GenS10{F0: 65535, F1: 412, F2: 152}
                                    _ = v117
                                    _ = v118
                                }
                                if true {
                                    v119 := GenS9{F0: 174, F1: 633}
                                    if false {
                                        v83 = 2
                                    }
                                    if false {
                                        if len(v91) > 0 {
                                            v91[0] = (fi92 + fi92)
                                        }
                                    } else {
                                        if (255 >= 32767) {
                                            var v120 int32 = 1
                                            _ = v120
                                        }
                                        switch v2 {
                                        case "foo bar":
                                            var v121 float32 = (0.1 / -2.25)
                                            _ = v121
                                        case "a":
                                            if true {
                                                v1 = 0
                                                v94 += -9223372036854775808
                                            }
                                        }
                                        v91 = append(v91, (fi93 & (fi92 / ((0) | 1))))
                                        switch 128 {
                                        case 32768:
                                            var v122 uint32 = 54
                                            _ = v122
                                        case 333:
                                            fmt.Println(fi93)
                                        }
                                    }
                                    _ = v119
                                }
                                fmt.Println(v83, v83)
                                v123 := v4[("" + v2)]
                                switch "" {
                                case "\n\t":
                                    _ = len(v91)
                                case "a":
                                    var v124 uint64 = 4294967295
                                    _ = v124
                                case "foo bar":
                                    var v125 uint8 = 128
                                    _ = v125
                                default:
                                    var v126 uint32 = 65535
                                    _ = v126
                                }
                                v1 -= v1
                                _ = v123
                            }
                            fmt.Println(v83, v94, v1)
                        } else {
                            if (!false) {
                                fmt.Println(v1)
                            } else {
                                v83 += (v83 << ((uint(v83)) & 63))
                            }
                            var v127 uint64 = v1
                            var v128 float64 = 0.1
                            _ = v127
                            _ = v128
                        }
                        if true {
                            switch 128 {
                            case -839:
                                v129 := GenS11{F0: 570}
                                _ = v129
                            case -384:
                                v1 = (v83 % ((v83) | 1))
                            default:
                                if (!true) {
                                    var v130 int8 = -128
                                    v3 = 1
                                    var v131 int64 = (v94 << ((uint(v83)) & 63))
                                    _ = v130
                                    _ = v131
                                } else {
                                    v132 := GenS8{F0: (v94 << ((443) & 63)), F1: 2147483648, F2: (v1 ^ (v1 >> ((589) & 63)))}
                                    if (false && false) {
                                        var v133 uint32 = 76
                                        var v134 int16 = 928
                                        _ = v133
                                        _ = v134
                                    } else {
                                        v3 = (-(v3))
                                        if (!(((3.14159 / 3.14159) <= (0.0 + (-2.25 * 0.0))) && (!(!(true || ((v2 + (v2 + ((v2 + v2) + "\n\t"))) >= v2)))))) {
                                            v1 += (v83 << ((4294967295) & 63))
                                        } else {
                                            v94 -= -1
                                            var v135 uint32 = 712
                                            v3 = v3
                                            _ = v135
                                        }
                                        var v136 int = fi92
                                        _ = v136
                                    }
                                    _ = v132
                                }
                            }
                        } else {
                            if true {
                                v137 := GenS12{F0: -108, F1: -918, F2: 2147483647}
                                v138 := GenS13{F0: 899, F1: 255, F2: 0}
                                var v139 uint8 = 174
                                _ = v137
                                _ = v138
                                _ = v139
                            }
                        }
                        var v140 float64 = (1.5 - 1.5)
                        _ = v140
                    } else {
                        var v141 int = 32768
                        v142 := GenS14{F0: "hello", F1: 2147483648, F2: (^(v83))}
                        v143 := v4["hello"]
                        _ = v141
                        _ = v142
                        _ = v143
                    }
                case "\n\t":
                    v1 = v83
                default:
                    _ = len(v4)
                }
                v144 := GenS15{F0: genFn0(39, v3), F1: (v94 << ((843) & 63)), F2: false}
                var v145 uint32 = 2
                v83 += v1
                _ = fi93
                _ = v144
                _ = v145
                _ = v94
            }
            _ = fi92
        }
        v146 := GenS16{F0: v3}
        v3 += genFn0(-1, (v3 / ((genFn0(-1, genFn0(32767, v3))) | 1)))
        _ = v146
    }
    if false {
        var v147 uint = 0
        _ = v147
    }
    switch -943 {
    case -9223372036854775808:
        for fi148 := 0; fi148 < 15; fi148++ {
            v149 := GenS17{F0: (v3 | v3), F1: 1, F2: v3}
            var v150 uint64 = v83
            fmt.Println(v83)
            _ = fi148
            _ = v149
            _ = v150
        }
    default:
        switch 4294967295 {
        case 0:
            if len(v91) >= 2 {
                v91 = v91[1:len(v91)]
            }
        case 32768:
            var v151 float64 = -2.25
            _ = v151
        default:
            v152 := GenS18{F0: v2, F1: (v83 - 517), F2: 255}
            _ = v152
        }
    }
    if false {
        v4[("hello" + ((v2 + v2) + v2))] = ((2 <= 127) && (!(-677 != genFn0(883, 1))))
        var v153 uint8 = 4
        for fi154 := 0; fi154 < 18; fi154++ {
            var v155 int64 = v3
            _ = fi154
            _ = v155
        }
        _ = v153
    } else {
        var v156 uint8 = 42
        _ = v156
    }
    switch 32768 {
    case -535:
        switch 32768 {
        case 1:
            if true {
                for fi157 := 0; fi157 < 6; fi157++ {
                    fmt.Println(v83)
                    var v158 int64 = 9223372036854775807
                    v83 -= (v83 / ((v83) | 1))
                    _ = fi157
                    _ = v158
                }
            }
        }
    case -9223372036854775808:
        for fi159 := 0; fi159 < 7; fi159++ {
            v2 = v2
            _ = fi159
        }
    default:
        fmt.Println(v3, v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v83:", v83)
    fmt.Println("slicelen v91:", len(v91), v91)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v83
    _ = v90
    _ = v91
}
