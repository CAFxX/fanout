package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16; F1 uint64; F2 int32 }
type GenS4 struct { F0 int64; F1 int16 }
type GenS5 struct { F0 int8; F1 uint16 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 bool }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 bool; F1 string; F2 uint64 }
func genFn0() uint16 {
    var v5 int8 = 0
    fmt.Println(v5)
    var v6 int64 = -1
    for fi7 := 0; fi7 < 10; fi7++ {
        v8 := []int{}
        v9 := GenS1{F0: 2147483647}
        var v10 int8 = -128
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    v6 -= 438
    _ = v5
    _ = v6
    return 2
}

func main() {
    v1 := make(map[string]bool)
    v2 := []uint32{1, 1}
    v3 := make(map[uint32]int)
    var v4 float32 = float32(-2.25)
    v4 += ((v4 - v4) * v4)
    var v11 uint16 = 65535
    switch -9223372036854775808 {
    case -195:
        if len(v2) > 0 {
            v2[0] = 2147483648
        }
    case -435:
        v4 -= v4
    case 32767:
        for fi12 := 0; fi12 < 4; fi12++ {
            if ((false && (42 <= 677)) && true) {
                for fi13 := 0; fi13 < 5; fi13++ {
                    if true {
                        if (true && true) {
                            if (!false) {
                                v2 = append(v2, 4294967295)
                                var v14 int8 = 79
                                _ = v14
                            }
                            v11 += 539
                        }
                        var v15 bool = (1e10 <= 0.1)
                        var v16 float32 = float32((float64((((1.5 + v4) - 3.14159) + ((v4 * (3.14159 + v4)) / 1.5))) / 0.1))
                        _ = v15
                        _ = v16
                    } else {
                        var v17 uint32 = 2
                        if (110 >= uint32(v11)) {
                            var v18 int = (fi12 & (fi13 ^ (fi12 - fi13)))
                            var v19 uint = 0
                            switch fi12 {
                            case 32767:
                                if len(v2) > 0 {
                                    v2[0] = (v17 % (((v17 | uint32(v11))) | 1))
                                }
                            case 493:
                                v11 = uint16(v17)
                            default:
                                fmt.Println(v19, fi13, fi13)
                            }
                            var v20 int16 = 1
                            var v21 float64 = (0.0 - -2.25)
                            _ = v18
                            _ = v19
                            _ = v20
                            _ = v21
                        } else {
                            if false {
                                v3[v17] = (fi12 - (fi12 + (fi13 & fi12)))
                            }
                            switch "foo bar" {
                            case "hello":
                                var v22 int = -9223372036854775808
                                _ = v22
                            case "":
                                v17 = v17
                            case "a":
                                var v23 uint = 1
                                _ = v23
                            }
                            if (("hello" != ("a" + "")) && true) {
                                v24 := GenS2{F0: (v17 >> ((851) & 63))}
                                var v25 uint16 = (v11 - genFn0())
                                _ = v24
                                _ = v25
                            }
                            switch "a" {
                            case "\n\t":
                                v26 := GenS3{F0: v11, F1: 4294967295, F2: -2147483648}
                                _ = v26
                            case "foo bar":
                                if true {
                                    fmt.Println(v11, fi13)
                                    if true {
                                        v27 := v3[4294967295]
                                        fmt.Println(v4)
                                        v28 := GenS3{F0: (v11 << ((18446744073709551615) & 63)), F1: 42, F2: 700}
                                        var v29 uint32 = v17
                                        if (!false) {
                                            v30 := GenS4{F0: -9223372036854775808, F1: 1000}
                                            var v31 int16 = 1
                                            if true {
                                                v32 := GenS2{F0: (v29 & v17)}
                                                var v33 uint32 = (v17 >> ((965) & 63))
                                                fmt.Println(v33)
                                                _ = v32
                                                _ = v33
                                            }
                                            _ = v30
                                            _ = v31
                                        }
                                        var v34 uint8 = 0
                                        _ = v27
                                        _ = v28
                                        _ = v29
                                        _ = v34
                                    } else {
                                        switch "\n\t" {
                                        case "":
                                            switch fi12 {
                                            case -1:
                                                var v35 uint64 = 760
                                                _ = v35
                                            case 128:
                                                var v36 bool = true
                                                _ = len(v1)
                                                _ = v36
                                            case 0:
                                                switch 4294967295 {
                                                case 0:
                                                    v1["foo bar"] = true
                                                case 9223372036854775807:
                                                    v37 := GenS1{F0: 553}
                                                    _ = v37
                                                case 771:
                                                    v2 = append(v2, v17)
                                                }
                                            }
                                        case "a":
                                            v17 = 602
                                        case "hello":
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        }
                                        v38 := GenS4{F0: -1, F1: -204}
                                        v39 := GenS5{F0: 127, F1: (v11 >> ((65535) & 63))}
                                        _ = v38
                                        _ = v39
                                    }
                                    switch 231 {
                                    case -9223372036854775808:
                                        if true {
                                            var v40 int = fi13
                                            _ = v40
                                        } else {
                                            var v41 float64 = 1.5
                                            _ = v41
                                        }
                                        v11 -= v11
                                    case 586:
                                        switch (fi12 % (((fi13 + (fi12 - fi13))) | 1)) {
                                        case 128:
                                            delete(v3, (v17 | 590))
                                        default:
                                            var v42 string = ("foo bar" + "hello")
                                            _ = v42
                                        }
                                    default:
                                        _ = len(v2)
                                    }
                                }
                            case "hello":
                                var v43 bool = (1 == 4294967295)
                                _ = v43
                            default:
                                var v44 int8 = 111
                                _ = v44
                            }
                        }
                        _ = v17
                    }
                    delete(v3, 333)
                    _ = fi13
                }
                fmt.Println(v4)
                v45 := GenS6{F0: 2}
                _ = v45
            }
            _ = fi12
        }
    }
    for fi46 := 0; fi46 < 5; fi46++ {
        var v47 int16 = 0
        _ = fi46
        _ = v47
    }
    v48 := GenS1{F0: 958}
    v1[("" + "foo bar")] = (false && ((68 >= 774) && false))
    switch "\n\t" {
    case "a":
        if (627 < 42) {
            fmt.Println(v4)
        } else {
            for fi49 := 0; fi49 < 8; fi49++ {
                var v50 int16 = -1
                if true {
                    v51 := v3[42]
                    _ = v51
                } else {
                    var v52 string = "foo bar"
                    var v53 int16 = (v50 >> ((41) & 63))
                    var v54 float32 = (v4 - (1e10 / (v4 + (1.5 + ((((0.0 + v4) * v4) / v4) + v4)))))
                    _ = v52
                    _ = v53
                    _ = v54
                }
                v55 := v3[42]
                _ = fi49
                _ = v50
                _ = v55
            }
            switch "" {
            case "":
                v4 -= ((((v4 * v4) - float32(0.0)) - v4) / v4)
            case "a":
                if (4294967295 <= 828) {
                    if len(v2) > 0 {
                        v2[0] = 190
                    }
                    fmt.Println(v11, v11)
                } else {
                    v11 += (^(v11))
                    var v56 uint64 = 4294967295
                    _ = v56
                }
            case "\n\t":
                if len(v2) > 0 {
                    v2[0] = 2
                }
            }
            switch -9223372036854775808 {
            case 127:
                for fi57 := 0; fi57 < 17; fi57++ {
                    fmt.Println(fi57, fi57, v4)
                    _ = fi57
                }
            case 1:
                if false {
                    if false {
                        fmt.Println(v4, v4)
                        _ = len(v3)
                    }
                    v58 := GenS7{F0: 65535}
                    if ((255 >= 0) || false) {
                        if true {
                            switch (("foo bar" + ("hello" + "a")) + "hello") {
                            case "hello":
                                for fi59 := 0; fi59 < 18; fi59++ {
                                    v1["\n\t"] = true
                                    _ = fi59
                                }
                            }
                            v60 := GenS7{F0: 346}
                            if true {
                                var v61 uint64 = 686
                                if false {
                                    for fi62 := 0; fi62 < 10; fi62++ {
                                        fmt.Println(v4)
                                        _ = fi62
                                    }
                                } else {
                                    v63 := GenS8{F0: 329}
                                    v64 := GenS7{F0: 18446744073709551615}
                                    if false {
                                        v61 += (v61 * 18446744073709551615)
                                    } else {
                                        switch ("\n\t" + "") {
                                        case "\n\t":
                                            v61 += (v61 / (((v61 % ((1) | 1))) | 1))
                                        case "foo bar":
                                            v4 += v4
                                        case "":
                                            var v65 float32 = v4
                                            _ = v65
                                        }
                                        v61 += (v61 % ((v61) | 1))
                                    }
                                    _ = v63
                                    _ = v64
                                }
                                _ = v61
                            } else {
                                for fi66 := 0; fi66 < 17; fi66++ {
                                    if true {
                                        v67 := GenS4{F0: -1, F1: -1}
                                        fmt.Println(v4, fi66)
                                        _ = v67
                                    }
                                    switch "foo bar" {
                                    case "a":
                                        for fi68 := 0; fi68 < 3; fi68++ {
                                            v69 := GenS6{F0: 793}
                                            if len(v2) > 0 {
                                                v2[0] = 746
                                            }
                                            v11 = 839
                                            _ = fi68
                                            _ = v69
                                        }
                                    case "":
                                        if false {
                                            var v70 int8 = 54
                                            switch fi66 {
                                            case 9223372036854775807:
                                                if (true || (!(113 < 372))) {
                                                    for fi71 := 0; fi71 < 11; fi71++ {
                                                        var v72 uint = 126
                                                        if (!(!false)) {
                                                            var v73 int32 = 0
                                                            _ = len(v2)
                                                            var v74 bool = ((v73 ^ (^(v73))) != (v73 - v73))
                                                            _ = v73
                                                            _ = v74
                                                        } else {
                                                            v75 := GenS3{F0: v11, F1: 4294967295, F2: -748}
                                                            fmt.Println(v72, v11, v4)
                                                            fmt.Println(v11)
                                                            _ = v75
                                                        }
                                                        v76 := GenS6{F0: 915}
                                                        _ = fi71
                                                        _ = v72
                                                        _ = v76
                                                    }
                                                    for fi77 := 0; fi77 < 4; fi77++ {
                                                        fmt.Println(v11)
                                                        fmt.Println(v70)
                                                        var v78 uint64 = 1
                                                        _ = fi77
                                                        _ = v78
                                                    }
                                                } else {
                                                    v79 := GenS3{F0: 833, F1: 4294967295, F2: -682}
                                                    if (228 < 2) {
                                                        var v80 uint8 = 42
                                                        if true {
                                                            v80 -= (v80 % ((128) | 1))
                                                            v80 -= (^(v80))
                                                            switch "a" {
                                                            case "a":
                                                                var v81 int16 = 161
                                                                _ = v81
                                                            }
                                                        } else {
                                                            var v82 uint16 = v11
                                                            for fi83 := 0; fi83 < 9; fi83++ {
                                                                if true {
                                                                    v4 = (v4 / 0.1)
                                                                } else {
                                                                    var v84 string = "a"
                                                                    var v85 uint8 = v80
                                                                    _ = v84
                                                                    _ = v85
                                                                }
                                                                v86 := GenS9{F0: ((("foo bar" + ("" + "a")) + "foo bar") + "a")}
                                                                if true {
                                                                    fmt.Println(v4)
                                                                }
                                                                _ = fi83
                                                                _ = v86
                                                            }
                                                            for fi87 := 0; fi87 < 19; fi87++ {
                                                                var v88 int16 = 42
                                                                _ = fi87
                                                                _ = v88
                                                            }
                                                            _ = v82
                                                        }
                                                        delete(v3, 577)
                                                        _ = v80
                                                    }
                                                    _ = v79
                                                }
                                            default:
                                                if (808 <= 638) {
                                                    v11 = (v11 % ((genFn0()) | 1))
                                                } else {
                                                    for fi89 := 0; fi89 < 9; fi89++ {
                                                        v90 := GenS10{F0: 1}
                                                        var v91 bool = true
                                                        _ = fi89
                                                        _ = v90
                                                        _ = v91
                                                    }
                                                    for fi92 := 0; fi92 < 6; fi92++ {
                                                        if ((v70 == 100) && (139 >= 604)) {
                                                            var v93 int = fi66
                                                            v3[4294967295] = (-(fi92))
                                                            _ = v93
                                                        }
                                                        if (2 >= 7) {
                                                            v2 = append(v2, 248)
                                                        } else {
                                                            v4 -= (v4 * float32((0.1 + 1e10)))
                                                            v70 -= -1
                                                            var v94 uint = 4294967295
                                                            _ = v94
                                                        }
                                                        _ = fi92
                                                    }
                                                    var v95 int8 = (v70 / (((v70 ^ v70)) | 1))
                                                    _ = v95
                                                }
                                            }
                                            v70 -= v70
                                            _ = v70
                                        } else {
                                            if (v4 == v4) {
                                                switch "\n\t" {
                                                case "\n\t":
                                                    if false {
                                                        if false {
                                                            if true {
                                                                v11 += (^(v11))
                                                            } else {
                                                                v11 = (v11 ^ genFn0())
                                                            }
                                                        } else {
                                                            v11 += 0
                                                        }
                                                        var v96 int = fi66
                                                        var v97 uint16 = v11
                                                        _ = v96
                                                        _ = v97
                                                    }
                                                case "hello":
                                                    v4 = float32(((1e10 - -2.25) - -2.25))
                                                }
                                                fmt.Println(v4, v4)
                                            }
                                        }
                                    case "hello":
                                        v11 += v11
                                    default:
                                        delete(v3, 4294967295)
                                    }
                                    _ = fi66
                                }
                                for fi98 := 0; fi98 < 14; fi98++ {
                                    for fi99 := 0; fi99 < 17; fi99++ {
                                        var v100 int = 128
                                        if (!true) {
                                            var v101 int16 = -1
                                            _ = v101
                                        }
                                        _ = fi99
                                        _ = v100
                                    }
                                    _ = len(v2)
                                    if (0 != 9223372036854775807) {
                                        var v102 bool = false
                                        switch 0 {
                                        case 521:
                                            v11 = (v11 ^ (v11 * 1))
                                        case 913:
                                            if false {
                                                v4 -= v4
                                            }
                                        case 405:
                                            switch "hello" {
                                            case "a":
                                                for fi103 := 0; fi103 < 1; fi103++ {
                                                    switch ("" + "hello") {
                                                    case "foo bar":
                                                        if false {
                                                            v4 = v4
                                                            fmt.Println(v4)
                                                        }
                                                    case "a":
                                                        v11 = v11
                                                    case "":
                                                        if (!(-402 < 42)) {
                                                            v4 -= v4
                                                            if (v102 || (!(-9223372036854775808 < 1))) {
                                                                v104 := GenS3{F0: v11, F1: 4294967295, F2: 42}
                                                                v105 := GenS4{F0: 0, F1: -629}
                                                                var v106 float64 = 1e10
                                                                _ = v104
                                                                _ = v105
                                                                _ = v106
                                                            } else {
                                                                v11 = genFn0()
                                                            }
                                                        }
                                                    }
                                                    v11 += (v11 ^ (v11 % (((v11 - 255)) | 1)))
                                                    _ = fi103
                                                }
                                            default:
                                                v11 += v11
                                            }
                                        default:
                                            for fi107 := 0; fi107 < 4; fi107++ {
                                                var v108 int = (fi107 / ((fi98) | 1))
                                                var v109 uint = 827
                                                _ = fi107
                                                _ = v108
                                                _ = v109
                                            }
                                        }
                                        var v110 int = fi98
                                        _ = v102
                                        _ = v110
                                    } else {
                                        var v111 bool = ((v11 % ((v11) | 1)) < v11)
                                        v111 = ((!v111) || (-1 == 0))
                                        _ = v111
                                    }
                                    _ = fi98
                                }
                            }
                            _ = v60
                        }
                    }
                    _ = v58
                } else {
                    v112 := GenS11{F0: 42}
                    switch ("a" + (((("a" + (("a" + "hello") + "")) + "hello") + "\n\t") + "hello")) {
                    case "a":
                        v11 = (v11 >> ((2) & 63))
                    case "foo bar":
                        var v113 int8 = 123
                        _ = v113
                    case "":
                        for fi114 := 0; fi114 < 13; fi114++ {
                            v11 = v11
                            if (18446744073709551615 >= 65535) {
                                if true {
                                    v115 := GenS12{F0: 173}
                                    var v116 int16 = 32767
                                    _ = v115
                                    _ = v116
                                }
                                if true {
                                    fmt.Println(v4)
                                    v11 = v11
                                    v117 := GenS4{F0: 42, F1: 0}
                                    _ = v117
                                } else {
                                    for fi118 := 0; fi118 < 7; fi118++ {
                                        if (!true) {
                                            switch ("" + ((("a" + ((("a" + "\n\t") + (("\n\t" + "") + ((("" + ("foo bar" + "\n\t")) + "") + (("a" + "hello") + "hello")))) + "foo bar")) + "\n\t") + "hello")) {
                                            case "":
                                                if (!(!(2 > 885))) {
                                                    switch 42 {
                                                    case 1:
                                                        var v119 string = ""
                                                        _ = v119
                                                    case 628:
                                                        v120 := GenS5{F0: 1, F1: (^(v11))}
                                                        fmt.Println(fi114)
                                                        _ = v120
                                                    default:
                                                        var v121 int64 = 0
                                                        v122 := GenS13{F0: false}
                                                        v121 = 738
                                                        var v123 uint8 = 57
                                                        v1[("\n\t" + "")] = (!(!true))
                                                        _ = v121
                                                        _ = v122
                                                        _ = v123
                                                    }
                                                    var v124 string = ""
                                                    var v125 int8 = 127
                                                    v126 := GenS9{F0: "\n\t"}
                                                    v125 = v125
                                                    _ = v124
                                                    _ = v125
                                                    _ = v126
                                                }
                                            case "a":
                                                var v127 int64 = 895
                                                _ = v127
                                            default:
                                                var v128 int64 = 288
                                                _ = v128
                                            }
                                            var v129 int8 = -1
                                            _ = v129
                                        } else {
                                            switch (fi118 ^ -9223372036854775808) {
                                            case 127:
                                                fmt.Println(v11, v11, fi114)
                                            }
                                            if true {
                                                if ((v11 | v11) >= (v11 % ((v11) | 1))) {
                                                    if false {
                                                        fmt.Println(fi114, v4)
                                                        fmt.Println(fi114, v4)
                                                    } else {
                                                        var v130 bool = true
                                                        _ = v130
                                                    }
                                                    var v131 int64 = 42
                                                    _ = v131
                                                }
                                                var v132 uint8 = 42
                                                v2 = append(v2, 4294967295)
                                                fmt.Println(v4)
                                                switch fi114 {
                                                case 533:
                                                    var v133 int = (fi118 - (fi118 | fi118))
                                                    _ = v133
                                                case -9223372036854775808:
                                                    var v134 string = ("hello" + "\n\t")
                                                    var v135 int32 = 1
                                                    _ = v134
                                                    _ = v135
                                                }
                                                var v136 int = fi114
                                                var v137 uint16 = (v11 | 1)
                                                _ = v132
                                                _ = v136
                                                _ = v137
                                            } else {
                                                switch (fi114 / ((-757) | 1)) {
                                                case -1:
                                                    v138 := GenS9{F0: "\n\t"}
                                                    _ = v138
                                                case -9223372036854775808:
                                                    v4 = (1.5 - float32(3.14159))
                                                case 1:
                                                    if false {
                                                        if (!(65535 >= 64)) {
                                                            var v139 uint8 = 206
                                                            v11 += (v11 % ((v11) | 1))
                                                            _ = v139
                                                        }
                                                        var v140 float64 = -2.25
                                                        _ = v140
                                                    }
                                                default:
                                                    var v141 string = "foo bar"
                                                    _ = v141
                                                }
                                                v11 -= v11
                                                v11 -= v11
                                                v142 := GenS9{F0: "foo bar"}
                                                v143 := v3[812]
                                                _ = v142
                                                _ = v143
                                            }
                                            switch fi114 {
                                            case -9223372036854775808:
                                                _ = len(v1)
                                            case 127:
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            default:
                                                var v144 string = "hello"
                                                _ = v144
                                            }
                                        }
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        _ = fi118
                                    }
                                }
                            }
                            _ = fi114
                        }
                    default:
                        if ((128 <= 0) && true) {
                            var v145 uint32 = 282
                            for fi146 := 0; fi146 < 4; fi146++ {
                                v11 -= v11
                                _ = fi146
                            }
                            v147 := GenS14{F0: -32768}
                            _ = v145
                            _ = v147
                        }
                    }
                    v4 = v4
                    _ = v112
                }
            case 32767:
                v11 -= 65535
            default:
                var v148 string = "foo bar"
                _ = v148
            }
        }
    }
    for fi149 := 0; fi149 < 8; fi149++ {
        v150 := v3[471]
        delete(v3, 0)
        if ((42 >= 18446744073709551615) && true) {
            v151 := GenS15{F0: true, F1: "hello", F2: 786}
            _ = v151
        }
        _ = fi149
        _ = v150
    }
    var v152 uint = 558
    v152 = (v152 % (((v152 / ((v152) | 1))) | 1))
    var v153 uint32 = 625
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi154 := 0; fi154 < 15; fi154++ {
        if false {
            v1[""] = false
        }
        switch ("a" + "hello") {
        case "\n\t":
            delete(v3, 0)
        case "foo bar":
            fmt.Println(v11, v4, v153)
        case "":
            _ = len(v3)
        }
        var v155 int64 = 626
        _ = fi154
        _ = v155
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v152:", v152)
    fmt.Println("v153:", v153)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v152
    _ = v153
    _ = v2
    _ = v3
    _ = v4
    _ = v48
}
