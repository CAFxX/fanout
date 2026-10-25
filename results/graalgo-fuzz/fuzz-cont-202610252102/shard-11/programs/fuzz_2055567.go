package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 uint; F1 int32; F2 string }
type GenS5 struct { F0 int8; F1 string; F2 uint16 }
type GenS6 struct { F0 int16; F1 int32 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint8; F1 uint; F2 uint16 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 int16; F1 int32 }

func main() {
    var v1 uint32 = 1
    v2 := make(map[uint32]bool)
    v3 := make(map[int]int)
    for fi4 := 0; fi4 < 18; fi4++ {
        v1 -= 2147483647
        _ = fi4
    }
    if (255 != 65535) {
        switch "hello" {
        case "\n\t":
            if false {
                v5 := []int{}
                v5 = append(v5, 0)
                switch 391 {
                case 9223372036854775807:
                    if len(v5) > 0 {
                        v5[0] = 703
                    }
                case -1:
                    v6 := GenS1{F0: (v1 | v1)}
                    _ = v6
                default:
                    var v7 uint32 = (v1 * v1)
                    _ = v7
                }
                _ = v5
            }
        default:
            v8 := GenS1{F0: 2147483648}
            _ = v8
        }
    }
    delete(v3, -87)
    _ = len(v2)
    fmt.Println(v1)
    var v9 int = 486
    var v10 int = v9
    if (42 < -904) {
        var v11 uint64 = 2
        _ = v11
    } else {
        fmt.Println(v10, v10, v9)
        var v12 int64 = -974
        _ = v12
    }
    switch ("" + "hello") {
    case "hello":
        v13 := GenS1{F0: (v1 % ((v1) | 1))}
        _ = v13
    case "foo bar":
        for fi14 := 0; fi14 < 17; fi14++ {
            for fi15 := 0; fi15 < 6; fi15++ {
                if true {
                    v16 := GenS2{F0: ("\n\t" + "hello"), F1: v10}
                    switch fi15 {
                    case 933:
                        fmt.Println(fi14)
                    case 127:
                        if (1.5 > 3.14159) {
                            var v17 int8 = 1
                            _ = v17
                        }
                        if false {
                            v18 := []int{fi15, v10, (fi15 & v9)}
                            var v19 uint = 18446744073709551615
                            v19 -= v19
                            _ = v18
                            _ = v19
                        }
                        v10 = (fi14 ^ (v9 / (((fi15 ^ (fi14 * -1))) | 1)))
                    default:
                        if (false || (v1 <= 4294967295)) {
                            var v20 int32 = -473
                            v21 := []int{fi14, 137, (fi15 - 127)}
                            fmt.Println(v10, v20, v10)
                            _ = v20
                            _ = v21
                        } else {
                            if ((18446744073709551615 < 1) && ((0 < 1) || (v9 > v9))) {
                                var v22 uint32 = 42
                                v23 := []int{}
                                _ = v22
                                _ = v23
                            }
                        }
                    }
                    _ = v16
                }
                v24 := GenS2{F0: "hello", F1: v10}
                v1 = (^(v1))
                _ = fi15
                _ = v24
            }
            if (true || true) {
                v25 := GenS3{F0: 1, F1: -2147483648}
                _ = v25
            } else {
                v26 := GenS1{F0: 0}
                _ = v26
            }
            v27 := []int{v9, 1, 32767, (^(v10))}
            _ = fi14
            _ = v27
        }
    case "a":
        v28 := GenS1{F0: v1}
        _ = v28
    }
    if (-1 != int32(v10)) {
        _ = len(v2)
        for fi29 := 0; fi29 < 12; fi29++ {
            v1 -= (v1 * v1)
            _ = fi29
        }
    }
    if false {
        v30 := GenS4{F0: 2, F1: 1, F2: "foo bar"}
        fmt.Println(v10)
        if true {
            var v31 bool = (v10 < 44)
            var v32 uint8 = 116
            _ = v31
            _ = v32
        }
        _ = v30
    } else {
        var v33 uint64 = uint64(v1)
        var v34 uint = 0
        switch (((("" + "hello") + (("foo bar" + "a") + "")) + "a") + (("" + "\n\t") + "hello")) {
        case "hello":
            fmt.Println(v1, v34)
        case "":
            if ((1e10 + 0.0) <= -2.25) {
                v35 := GenS4{F0: v34, F1: 435, F2: "hello"}
                v36 := GenS5{F0: 49, F1: ("hello" + "a"), F2: 2}
                v9 += (v9 - 128)
                _ = v35
                _ = v36
            }
        case "a":
            if true {
                for fi37 := 0; fi37 < 10; fi37++ {
                    var v38 int8 = -1
                    if true {
                        var v39 float32 = (0.0 / -2.25)
                        _ = v39
                    } else {
                        switch ("a" + "") {
                        case "a":
                            for fi40 := 0; fi40 < 10; fi40++ {
                                v41 := []int{(v9 - 2147483647), int(v34), (fi40 / (((v9 * fi40)) | 1))}
                                var v42 string = ""
                                var v43 uint8 = 1
                                var v44 bool = (!(!(3.14159 > (((0.0 + 1.5) * 0.0) / 3.14159))))
                                _ = fi40
                                _ = v41
                                _ = v42
                                _ = v43
                                _ = v44
                            }
                        case "hello":
                            fmt.Println(v33, v10, v34)
                        case "foo bar":
                            v10 += 487
                        }
                        if ((v10 ^ fi37) > (-(fi37))) {
                            v9 = 4294967295
                            var v45 float32 = 0.1
                            _ = v45
                        }
                        if (4294967295 <= -292) {
                            fmt.Println(v33, fi37, v34)
                            for fi46 := 0; fi46 < 8; fi46++ {
                                var v47 float32 = 3.14159
                                v48 := GenS6{F0: 0, F1: -39}
                                _ = fi46
                                _ = v47
                                _ = v48
                            }
                            v49 := GenS5{F0: (v38 / ((int8(v10)) | 1)), F1: ("" + "hello"), F2: 65535}
                            _ = v49
                        } else {
                            if ((!((0 < 9223372036854775807) && (uint8(v33) <= 0))) && true) {
                                if true {
                                    if ((v33 | 196) <= v33) {
                                        v50 := []int{fi37, v10, (fi37 & 4294967295), (v9 / ((42) | 1))}
                                        _ = v50
                                    }
                                    switch v9 {
                                    case 0:
                                        _ = len(v3)
                                    case 42:
                                        v51 := []int{(v9 / ((v10) | 1)), fi37, (v10 + fi37)}
                                        _ = v51
                                    case -885:
                                        switch (v9 - v9) {
                                        case 32767:
                                            v52 := []int{0, -301}
                                            _ = v52
                                        case 32768:
                                            for fi53 := 0; fi53 < 4; fi53++ {
                                                var v54 int64 = 342
                                                v55 := GenS3{F0: v33, F1: int32(v33)}
                                                var v56 float64 = 0.0
                                                var v57 float64 = v56
                                                v10 += 2147483647
                                                _ = fi53
                                                _ = v54
                                                _ = v55
                                                _ = v56
                                                _ = v57
                                            }
                                        case 42:
                                            switch ("\n\t" + "") {
                                            case "\n\t":
                                                v9 = v9
                                            case "foo bar":
                                                switch 863 {
                                                case 42:
                                                    v10 = fi37
                                                case -151:
                                                    v58 := []int{(fi37 ^ (-(v10))), (v9 & (v9 >> (((v34 | 0)) & 63))), v9, 344}
                                                    _ = v58
                                                case 637:
                                                    fmt.Println(fi37, v1)
                                                default:
                                                    var v59 string = "foo bar"
                                                    _ = v59
                                                }
                                            default:
                                                if (true || (!(!false))) {
                                                    v60 := []int{v10, fi37, (^(v9))}
                                                    var v61 uint32 = 50
                                                    _ = v60
                                                    _ = v61
                                                } else {
                                                    v62 := GenS5{F0: v38, F1: "a", F2: 65535}
                                                    if false {
                                                        var v63 uint32 = (^(v1))
                                                        if (v10 != v9) {
                                                            var v64 float64 = -2.25
                                                            var v65 float32 = 1.5
                                                            v38 -= (v38 | v38)
                                                            _ = v64
                                                            _ = v65
                                                        } else {
                                                            for fi66 := 0; fi66 < 12; fi66++ {
                                                                if true {
                                                                    v9 = (v10 | (fi37 * 449))
                                                                } else {
                                                                    switch "foo bar" {
                                                                    case "":
                                                                        _ = len(v3)
                                                                        v67 := []int{}
                                                                        fmt.Println(v33)
                                                                        _ = v67
                                                                    case "hello":
                                                                        v34 = 643
                                                                    }
                                                                    if (false && false) {
                                                                        v33 -= v33
                                                                        fmt.Println(v9, v33)
                                                                        var v68 int16 = -1
                                                                        var v69 int16 = v68
                                                                        _ = v68
                                                                        _ = v69
                                                                    } else {
                                                                        _ = len(v3)
                                                                        var v70 int64 = 0
                                                                        _ = v70
                                                                    }
                                                                }
                                                                var v71 bool = (255 != v34)
                                                                fmt.Println(v9, v33)
                                                                _ = fi66
                                                                _ = v71
                                                            }
                                                        }
                                                        delete(v3, 2147483647)
                                                        _ = v63
                                                    }
                                                    v72 := v2[v1]
                                                    _ = v62
                                                    _ = v72
                                                }
                                            }
                                        }
                                    default:
                                        v38 += int8(v9)
                                    }
                                    if (true || (((0 > 42) && ((-(v38)) > v38)) || ((!true) && true))) {
                                        v34 -= uint(v33)
                                        v10 -= -9223372036854775808
                                    } else {
                                        v73 := v3[fi37]
                                        v74 := GenS7{F0: v38}
                                        var v75 int = 2147483647
                                        _ = v73
                                        _ = v74
                                        _ = v75
                                    }
                                }
                            }
                            var v76 int = 128
                            _ = v76
                        }
                    }
                    switch "foo bar" {
                    case "":
                        if true {
                            var v77 int32 = -1
                            if (!(127 >= 128)) {
                                fmt.Println(v38)
                            } else {
                                if (!true) {
                                    for fi78 := 0; fi78 < 6; fi78++ {
                                        var v79 uint64 = 635
                                        _ = fi78
                                        _ = v79
                                    }
                                    var v80 string = ""
                                    v81 := GenS6{F0: 42, F1: 919}
                                    _ = v80
                                    _ = v81
                                }
                                var v82 int = fi37
                                switch "\n\t" {
                                case "\n\t":
                                    v83 := GenS3{F0: v33, F1: v77}
                                    _ = v83
                                default:
                                    var v84 int8 = v38
                                    _ = v84
                                }
                                _ = v82
                            }
                            switch fi37 {
                            case 264:
                                v85 := GenS8{F0: 128, F1: v34, F2: 0}
                                _ = v85
                            }
                            _ = v77
                        } else {
                            if ((fi37 * (^(fi37))) >= (-(v9))) {
                                if (2 < v1) {
                                    v1 -= 2147483647
                                    switch fi37 {
                                    case -355:
                                        if ((3.14159 < ((1e10 - -2.25) / float64((0.1 + ((((3.14159 - (float32(((((3.14159 / 1e10) * 0.1) / (((-2.25 - 1e10) - float64((1.5 - 3.14159))) + float64(3.14159))) * float64(float32((0.0 + 1.5))))) + 0.1)) * 1.5) + 1e10) / 1e10))))) && false) {
                                            switch "a" {
                                            case "foo bar":
                                                var v86 bool = false
                                                _ = v86
                                            default:
                                                v10 += (-(v9))
                                            }
                                            switch "foo bar" {
                                            case "hello":
                                                v87 := GenS6{F0: 683, F1: 42}
                                                _ = v87
                                            case "":
                                                delete(v2, 65535)
                                            case "\n\t":
                                                for fi88 := 0; fi88 < 8; fi88++ {
                                                    fmt.Println(v38, v10)
                                                    v9 += (v10 << ((1) & 63))
                                                    var v89 int8 = (v38 | v38)
                                                    _ = fi88
                                                    _ = v89
                                                }
                                            default:
                                                v90 := GenS1{F0: v1}
                                                _ = v90
                                            }
                                            var v91 int = fi37
                                            _ = v91
                                        } else {
                                            var v92 int32 = 42
                                            for fi93 := 0; fi93 < 8; fi93++ {
                                                v33 = (^(v33))
                                                var v94 uint = (^(v34))
                                                _ = fi93
                                                _ = v94
                                            }
                                            if true {
                                                for fi95 := 0; fi95 < 5; fi95++ {
                                                    var v96 float64 = 1.5
                                                    v33 = (v33 * (v33 ^ v33))
                                                    var v97 uint16 = uint16(v1)
                                                    var v98 int = 782
                                                    switch fi37 {
                                                    case 32767:
                                                        v99 := GenS9{F0: v1}
                                                        var v100 int = (^(v10))
                                                        _ = v100
                                                        _ = v99
                                                    default:
                                                        _ = len(v2)
                                                    }
                                                    v101 := GenS8{F0: 189, F1: (^(v34)), F2: v97}
                                                    fmt.Println(v33)
                                                    _ = fi95
                                                    _ = v101
                                                    _ = v96
                                                    _ = v97
                                                    _ = v98
                                                }
                                            }
                                            _ = v92
                                        }
                                    }
                                }
                                fmt.Println(v10)
                            }
                            v33 = (v33 & 871)
                        }
                    case "foo bar":
                        var v102 uint = v34
                        _ = v102
                    case "\n\t":
                        v103 := GenS5{F0: v38, F1: "foo bar", F2: 0}
                        _ = v103
                    default:
                        v2[v1] = (326 == 2)
                    }
                    _ = fi37
                    _ = v38
                }
                var v104 uint64 = 18446744073709551615
                _ = v104
            } else {
                var v105 bool = (!false)
                if (!((v1 > 1) && true)) {
                    v33 += 42
                    for fi106 := 0; fi106 < 6; fi106++ {
                        v107 := []int{v9, -2147483648, int(v33)}
                        _ = fi106
                        _ = v107
                    }
                    for fi108 := 0; fi108 < 6; fi108++ {
                        if false {
                            if ((^(v33)) == 18446744073709551615) {
                                if (1 > uint16(v33)) {
                                    fmt.Println(v10)
                                    var v109 uint = v34
                                    v110 := GenS10{F0: v34}
                                    _ = v109
                                    _ = v110
                                } else {
                                    switch "" {
                                    case "hello":
                                        if (v34 != (v34 / ((4294967295) | 1))) {
                                            var v111 uint64 = (^(v33))
                                            v1 = (v1 ^ v1)
                                            fmt.Println(v1)
                                            _ = v111
                                        }
                                    case "a":
                                        var v112 uint32 = v1
                                        _ = v112
                                    case "\n\t":
                                        var v113 int64 = -156
                                        _ = v113
                                    default:
                                        if true {
                                            for fi114 := 0; fi114 < 13; fi114++ {
                                                v105 = v105
                                                switch "a" {
                                                case "a":
                                                    var v115 bool = ((0 != 733) || false)
                                                    _ = v115
                                                case "":
                                                    var v116 int16 = 77
                                                    _ = v116
                                                case "foo bar":
                                                    var v117 uint64 = v33
                                                    _ = v117
                                                default:
                                                    v34 -= v34
                                                }
                                                var v118 int8 = 42
                                                v119 := []int{(fi114 + (v9 * (fi108 << ((v34) & 63)))), -1, fi114, (^(v9))}
                                                v9 = (^(fi108))
                                                _ = fi114
                                                _ = v118
                                                _ = v119
                                            }
                                            v120 := []int{(v10 + (v10 | v10)), v9, -2147483648}
                                            _ = v120
                                        } else {
                                            v121 := GenS1{F0: (v1 << (((^(v34))) & 63))}
                                            fmt.Println(v33)
                                            for fi122 := 0; fi122 < 8; fi122++ {
                                                var v123 int16 = -645
                                                if (!v105) {
                                                    if v105 {
                                                        if v105 {
                                                            delete(v2, 4294967295)
                                                            v124 := GenS5{F0: -128, F1: ((("a" + "\n\t") + "foo bar") + (("hello" + ("hello" + "\n\t")) + ("hello" + (("" + ("" + ("\n\t" + ("foo bar" + "hello")))) + ((("a" + ("foo bar" + "hello")) + "foo bar") + ("a" + "")))))), F2: 42}
                                                            _ = v124
                                                        } else {
                                                            var v125 int8 = 0
                                                            _ = v125
                                                        }
                                                        switch fi108 {
                                                        case -514:
                                                            v105 = ((309 <= 850) || true)
                                                        default:
                                                            switch "hello" {
                                                            case "foo bar":
                                                                var v126 uint = (v34 ^ v34)
                                                                _ = v126
                                                            }
                                                        }
                                                        v123 = v123
                                                        switch "" {
                                                        case "hello":
                                                            v127 := GenS1{F0: v1}
                                                            _ = v127
                                                        default:
                                                            fmt.Println(v123, v10, v123)
                                                        }
                                                        var v128 bool = (!true)
                                                        _ = v128
                                                    } else {
                                                        var v129 bool = v105
                                                        _ = v129
                                                    }
                                                    var v130 bool = true
                                                    _ = v130
                                                }
                                                var v131 string = ("" + "\n\t")
                                                if ((1.5 * 3.14159) != 1.5) {
                                                    var v132 uint64 = (^(v33))
                                                    _ = v132
                                                } else {
                                                    var v133 string = "foo bar"
                                                    v133 = v131
                                                    if true {
                                                        v134 := GenS11{F0: 1}
                                                        _ = v134
                                                    }
                                                    _ = v133
                                                }
                                                v9 = (fi108 << ((v34) & 63))
                                                _ = fi122
                                                _ = v123
                                                _ = v131
                                            }
                                            _ = v121
                                        }
                                    }
                                }
                                v33 = uint64(v1)
                            }
                            fmt.Println(v1, fi108)
                        }
                        fmt.Println(v10)
                        fmt.Println(fi108, v34, v1)
                        _ = fi108
                    }
                }
                if (((1000 > -1) || (954 == 41)) && ((v33 >> ((v34) & 63)) > (v33 * 2))) {
                    var v135 float32 = 1.5
                    v33 = (v33 | (^(v33)))
                    var v136 float64 = (1.5 / (float64(((v135 + v135) / v135)) * 0.0))
                    _ = v135
                    _ = v136
                }
                _ = v105
            }
        default:
            v10 = v9
        }
        _ = v33
        _ = v34
    }
    v137 := GenS8{F0: 42, F1: 42, F2: 2}
    fmt.Println(v9)
    switch ("hello" + "") {
    case "foo bar":
        var v138 uint8 = 255
        _ = v138
    case "":
        if (!(true && (0.1 <= (float32(((0.0 + (float64(-2.25) + 0.0)) / float64((-2.25 / 1.5)))) + (((0.0 + 3.14159) / 3.14159) - float32(float64(-2.25))))))) {
            var v139 float32 = (1.5 * (0.0 * -2.25))
            _ = v139
        } else {
            var v140 uint8 = 123
            _ = v140
        }
    default:
        v1 = v1
    }
    for fi141 := 0; fi141 < 9; fi141++ {
        switch v10 {
        case 4294967295:
            var v142 int64 = int64(v1)
            _ = v142
        case 2147483647:
            var v143 int32 = 15
            _ = v143
        case -1:
            v144 := GenS12{F0: -2147483648}
            _ = v144
        default:
            var v145 uint8 = 197
            _ = v145
        }
        for fi146 := 0; fi146 < 5; fi146++ {
            v10 = -1
            if false {
                switch (fi141 + -2147483648) {
                case -9223372036854775808:
                    v1 += (v1 * v1)
                case 128:
                    fmt.Println(fi146, fi146, fi146)
                default:
                    v2[(v1 & (v1 + v1))] = ((4294967295 != 0) && (136 < 42))
                }
                switch (v10 * (v9 + 2147483647)) {
                case 127:
                    if (1 > -128) {
                        var v147 float32 = -2.25
                        if false {
                            fmt.Println(v10, fi141)
                        }
                        var v148 uint32 = v1
                        _ = v147
                        _ = v148
                    } else {
                        v149 := []int{-1, 4294967295}
                        _ = v149
                    }
                case 2147483647:
                    if true {
                        v150 := GenS13{F0: 0, F1: 1}
                        v1 -= v1
                        _ = v150
                    }
                case 4294967295:
                    v151 := GenS7{F0: -128}
                    _ = v151
                }
            } else {
                v152 := v2[v1]
                v153 := v3[4294967295]
                delete(v2, v1)
                if (false || (1 > 1)) {
                    var v154 int64 = -1
                    var v155 uint16 = 918
                    v156 := []int{v153, (v10 >> ((848) & 63)), int(v155), 128}
                    _ = v154
                    _ = v155
                    _ = v156
                } else {
                    var v157 int16 = 643
                    _ = v157
                }
                v152 = v152
                v158 := GenS8{F0: 41, F1: 2, F2: 1}
                var v159 string = "a"
                _ = v152
                _ = v153
                _ = v158
                _ = v159
            }
            _ = fi146
        }
        _ = fi141
    }
    for fi160 := 0; fi160 < 6; fi160++ {
        v9 += v9
        var v161 int32 = 0
        if true {
            fmt.Println(v10)
            v162 := GenS10{F0: 585}
            _ = v162
        }
        _ = fi160
        _ = v161
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v137
    _ = v2
    _ = v3
    _ = v9
}
