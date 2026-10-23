package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16; F1 int32; F2 int32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint32; F1 uint8; F2 int64 }
type GenS6 struct { F0 uint8; F1 uint16; F2 uint64 }
type GenS7 struct { F0 uint32; F1 uint32 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint16; F1 int16; F2 int8 }
type GenS10 struct { F0 int16; F1 int; F2 int32 }
type GenS11 struct { F0 uint16; F1 string; F2 int }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 uint64; F1 int }
type GenS15 struct { F0 uint8 }
type GenS16 struct { F0 int32; F1 int64 }
type GenS17 struct { F0 int }
type GenS18 struct { F0 int64 }
type GenS19 struct { F0 int32; F1 bool }
type GenS20 struct { F0 int32; F1 uint16; F2 int16 }
type GenS21 struct { F0 string; F1 uint16; F2 int32 }
type GenS22 struct { F0 uint8; F1 string }
type GenS23 struct { F0 uint16 }
type GenS24 struct { F0 uint64; F1 uint16 }
type GenS25 struct { F0 int16 }
type GenS26 struct { F0 string; F1 uint; F2 int32 }
type GenS27 struct { F0 uint; F1 bool }
type GenS28 struct { F0 uint }
type GenS29 struct { F0 bool }
func genFn0(p0 uint8, p1 int64, p2 uint) bool {
    p2 -= (^(p2))
    v5 := GenS1{F0: 42}
    for fi6 := 0; fi6 < 6; fi6++ {
        if true {
            var v7 uint = p2
            v8 := GenS2{F0: 127}
            _ = v7
            _ = v8
        }
        v9 := []int{-1}
        _ = fi6
        _ = v9
    }
    var v10 uint64 = 799
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v5
    return false
}
func genFn1(p0 uint16, p1 uint8) int {
    v11 := []int{42, 2147483647, -9223372036854775808}
    v11 = append(v11, 0)
    for fi12 := 0; fi12 < 14; fi12++ {
        _ = len(v11)
        _ = fi12
    }
    var v13 float64 = (1e10 - 0.1)
    switch -2147483648 {
    case -346:
        if false {
            v14 := GenS3{F0: (p0 & p0), F1: 327, F2: -504}
            _ = v14
        }
    case -2147483648:
        v15 := GenS2{F0: -1}
        _ = v15
    case -9223372036854775808:
        v16 := GenS4{F0: 42}
        _ = v16
    default:
        fmt.Println(p0, p0, p0)
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v13
    return 46
}

func main() {
    var v1 uint = 42
    var v2 float32 = ((0.1 - (1e10 / 3.14159)) / -2.25)
    v3 := make(map[uint32]bool)
    var v4 int8 = -27
    var v17 uint16 = 65535
    for fi18 := 0; fi18 < 12; fi18++ {
        if (true || (!false)) {
            v19 := GenS5{F0: 2147483648, F1: uint8(v17), F2: -651}
            v17 = v17
            _ = v19
        }
        if true {
            for fi20 := 0; fi20 < 8; fi20++ {
                v21 := []int{fi18, fi20}
                v22 := GenS5{F0: 2, F1: uint8(v17), F2: 1}
                if true {
                    var v23 int64 = 9223372036854775807
                    var v24 int16 = int16(v1)
                    var v25 uint64 = 233
                    _ = v23
                    _ = v24
                    _ = v25
                }
                _ = fi20
                _ = v21
                _ = v22
            }
            var v26 string = "a"
            _ = v26
        }
        if false {
            switch fi18 {
            case -866:
                _ = len(v3)
            }
            v27 := GenS2{F0: v4}
            _ = v27
        } else {
            v2 -= 1.5
            v2 = 1e10
            switch 1 {
            case 27:
                var v28 uint64 = 811
                _ = v28
            default:
                fmt.Println(v17, v2)
            }
        }
        _ = fi18
    }
    var v29 float32 = -2.25
    if (65535 != 4294967295) {
        if true {
            switch "\n\t" {
            case "a":
                if false {
                    for fi30 := 0; fi30 < 6; fi30++ {
                        if true {
                            var v31 uint = (^(v1))
                            v32 := GenS5{F0: 4294967295, F1: 127, F2: 1}
                            switch "\n\t" {
                            case "\n\t":
                                switch fi30 {
                                case 1:
                                    switch ((("" + "a") + ("a" + "\n\t")) + "\n\t") {
                                    case "":
                                        delete(v3, 368)
                                    case "\n\t":
                                        for fi33 := 0; fi33 < 7; fi33++ {
                                            var v34 float32 = 0.0
                                            v2 = (v34 * v29)
                                            var v35 int8 = (v4 >> ((525) & 63))
                                            _ = fi33
                                            _ = v34
                                            _ = v35
                                        }
                                    }
                                case -693:
                                    var v36 int64 = 9223372036854775807
                                    _ = v36
                                case -1:
                                    for fi37 := 0; fi37 < 4; fi37++ {
                                        switch (fi37 - genFn1(v17, 255)) {
                                        case 261:
                                            fmt.Println(fi37)
                                        case 42:
                                            v2 = 1e10
                                        case -519:
                                            if ((v17 >> ((259) & 63)) <= (v17 + v17)) {
                                                v31 = 279
                                                v38 := []int{int(v4), genFn1(v17, 146), (fi30 >> ((42) & 63)), 42}
                                                _ = v38
                                            } else {
                                                fmt.Println(v29, fi37)
                                            }
                                        }
                                        var v39 int64 = 1
                                        _ = fi37
                                        _ = v39
                                    }
                                }
                            case "a":
                                v40 := GenS6{F0: 127, F1: (v17 << (((v31 ^ v31)) & 63)), F2: 82}
                                _ = v40
                            case "":
                                v41 := []int{genFn1((v17 - 0), 42)}
                                _ = v41
                            default:
                                switch 4294967295 {
                                case 82:
                                    var v42 uint8 = 255
                                    _ = v42
                                case 654:
                                    var v43 uint = v1
                                    _ = v43
                                default:
                                    fmt.Println(v4, v2, v17)
                                }
                            }
                            _ = v31
                            _ = v32
                        } else {
                            var v44 int16 = -32768
                            for fi45 := 0; fi45 < 9; fi45++ {
                                switch "foo bar" {
                                case "a":
                                    switch -353 {
                                    case -1:
                                        v46 := []int{fi45, (fi45 % (((fi30 & -146)) | 1)), -1}
                                        _ = v46
                                    }
                                case "":
                                    var v47 float32 = (v2 * v2)
                                    _ = v47
                                default:
                                    v48 := GenS4{F0: (v1 / (((v1 << ((v1) & 63))) | 1))}
                                    _ = v48
                                }
                                _ = fi45
                            }
                            for fi49 := 0; fi49 < 1; fi49++ {
                                v50 := []int{(fi49 * (fi49 * fi49))}
                                var v51 uint32 = 0
                                switch 966 {
                                case 833:
                                    v52 := v3[v51]
                                    _ = v52
                                case 1:
                                    v53 := GenS7{F0: (v51 ^ (v51 * v51)), F1: uint32(v17)}
                                    _ = v53
                                case 0:
                                    switch (fi30 << (((^(v1))) & 63)) {
                                    case 42:
                                        if (uint64(v51) > 0) {
                                            if (v1 == 1) {
                                                v2 = float32(3.14159)
                                                v4 += v4
                                            } else {
                                                fmt.Println(v2, v17, v44)
                                            }
                                            delete(v3, (v51 >> (((^(v1))) & 63)))
                                            var v54 string = "\n\t"
                                            _ = v54
                                        }
                                    default:
                                        if (false && false) {
                                            if true {
                                                if true {
                                                    var v55 float32 = v29
                                                    _ = v55
                                                }
                                            } else {
                                                v56 := GenS8{F0: "a"}
                                                v44 += (^(v44))
                                                _ = v56
                                            }
                                            switch "" {
                                            case "hello":
                                                v57 := GenS9{F0: v17, F1: 0, F2: (v4 - int8(v17))}
                                                _ = v57
                                            case "":
                                                v44 = (v44 >> ((762) & 63))
                                            case "a":
                                                if len(v50) > 0 {
                                                    v50[0] = genFn1(v17, 0)
                                                }
                                            }
                                            v4 += (^(v4))
                                            v50 = append(v50, -2147483648)
                                        } else {
                                            fmt.Println(v4, v4, v51)
                                            var v58 uint64 = 194
                                            _ = v58
                                        }
                                        _ = len(v50)
                                    }
                                }
                                _ = fi49
                                _ = v50
                                _ = v51
                            }
                            _ = v44
                        }
                        var v59 bool = ("foo bar" != ("" + "\n\t"))
                        var v60 int32 = -2147483648
                        _ = fi30
                        _ = v59
                        _ = v60
                    }
                    switch 143 {
                    case 0:
                        var v61 float64 = ((0.0 - 0.1) / -2.25)
                        _ = v61
                    }
                }
            case "foo bar":
                for fi62 := 0; fi62 < 9; fi62++ {
                    for fi63 := 0; fi63 < 11; fi63++ {
                        if false {
                            v64 := GenS10{F0: 42, F1: 1, F2: 1}
                            _ = v64
                        }
                        _ = fi63
                    }
                    v2 = (((v2 + (v29 * 3.14159)) - (((v2 + v29) * 3.14159) + v29)) / v2)
                    _ = fi62
                }
            case "":
                if (!((v4 - (v4 * (v4 | v4))) > (v4 >> ((v1) & 63)))) {
                    v65 := GenS11{F0: 0, F1: "\n\t", F2: -1}
                    _ = v65
                }
            default:
                v66 := v3[2]
                _ = v66
            }
            for fi67 := 0; fi67 < 20; fi67++ {
                switch "hello" {
                case "foo bar":
                    v1 = (v1 | 1)
                case "a":
                    fmt.Println(v1, v2, v17)
                case "hello":
                    if false {
                        switch (fi67 & (fi67 >> ((v1) & 63))) {
                        case -825:
                            for fi68 := 0; fi68 < 9; fi68++ {
                                v2 = v29
                                switch 32767 {
                                case 4294967295:
                                    v2 = (v2 * 1e10)
                                case 128:
                                    if (-589 >= 1) {
                                        v69 := GenS4{F0: 189}
                                        var v70 uint16 = (v17 ^ v17)
                                        _ = v69
                                        _ = v70
                                    } else {
                                        var v71 int8 = (^(v4))
                                        var v72 string = ((("" + "") + "a") + "foo bar")
                                        _ = v71
                                        _ = v72
                                    }
                                default:
                                    var v73 int16 = 32767
                                    _ = v73
                                }
                                _ = fi68
                            }
                        case 32767:
                            if false {
                                v74 := []int{(fi67 + (fi67 - fi67)), (fi67 / ((-1) | 1)), fi67}
                                var v75 int16 = -1
                                _ = v74
                                _ = v75
                            } else {
                                var v76 string = ""
                                v4 = v4
                                var v77 uint32 = 4294967295
                                _ = v76
                                _ = v77
                            }
                        }
                        v2 -= v29
                    }
                default:
                    if (-1 != -1) {
                        v78 := GenS12{F0: 4294967295}
                        _ = v78
                    } else {
                        var v79 uint16 = v17
                        _ = v79
                    }
                }
                _ = fi67
            }
            var v80 uint64 = 60
            _ = v80
        } else {
            fmt.Println(v4)
            v81 := GenS13{F0: 1}
            _ = v81
        }
        var v82 int32 = -2147483648
        _ = v82
    } else {
        var v83 int = -1
        for fi84 := 0; fi84 < 16; fi84++ {
            var v85 uint64 = uint64(v17)
            if false {
                var v86 uint = v1
                for fi87 := 0; fi87 < 6; fi87++ {
                    v85 = 350
                    var v88 uint = (v1 << (((^(v86))) & 63))
                    _ = fi87
                    _ = v88
                }
                fmt.Println(v2, v4, fi84)
                _ = v86
            } else {
                if false {
                    fmt.Println(v4)
                    fmt.Println(v4, v4)
                    switch "" {
                    case "a":
                        v89 := []int{(fi84 | (v83 >> (((v1 >> (((v1 | (v1 % (((v1 << ((18446744073709551615) & 63))) | 1)))) & 63))) & 63))), (v83 >> (((v1 & v1)) & 63))}
                        _ = v89
                    default:
                        fmt.Println(v83, fi84, fi84)
                    }
                } else {
                    var v90 float64 = 0.0
                    if ((^(v1)) != 428) {
                        var v91 bool = true
                        var v92 float64 = -2.25
                        var v93 uint64 = (v85 & v85)
                        _ = v91
                        _ = v92
                        _ = v93
                    } else {
                        if (255 >= 608) {
                            var v94 bool = true
                            v94 = (0 < 255)
                            var v95 uint8 = 127
                            _ = v94
                            _ = v95
                        }
                        if true {
                            for fi96 := 0; fi96 < 11; fi96++ {
                                switch (("" + "a") + ("foo bar" + "\n\t")) {
                                case "hello":
                                    v29 += v29
                                case "\n\t":
                                    v90 -= float64(1e10)
                                default:
                                    var v97 int32 = 1
                                    _ = v97
                                }
                                var v98 int32 = int32(v17)
                                _ = fi96
                                _ = v98
                            }
                            if (v17 > 255) {
                                if (2 != 2147483648) {
                                    var v99 float64 = (v90 / (1.5 / v90))
                                    _ = v99
                                }
                                v100 := GenS14{F0: v85, F1: v83}
                                _ = v100
                            } else {
                                var v101 int16 = 1
                                _ = v101
                            }
                        }
                    }
                    switch "" {
                    case "a":
                        v85 = (v85 + 42)
                    case "hello":
                        switch (-(fi84)) {
                        case 0:
                            switch "" {
                            case "hello":
                                v102 := GenS15{F0: 95}
                                _ = v102
                            case "a":
                                if (!(true || (!(!true)))) {
                                    v1 = v1
                                    fmt.Println(v29, v2)
                                } else {
                                    v17 = v17
                                }
                            }
                        case 128:
                            v29 -= 0.0
                        case -2147483648:
                            v103 := GenS5{F0: 2147483648, F1: 2, F2: 825}
                            _ = v103
                        }
                    }
                    _ = v90
                }
            }
            if (1 == 2) {
                fmt.Println(fi84, v85)
                var v104 uint = v1
                var v105 uint8 = 2
                _ = v104
                _ = v105
            }
            _ = fi84
            _ = v85
        }
        _ = v83
    }
    v2 -= float32(float64(v29))
    if (((v4 - v4) > v4) && true) {
        v106 := v3[2147483647]
        if v106 {
            for fi107 := 0; fi107 < 3; fi107++ {
                v106 = (("a" + "") < "a")
                v17 = 2
                _ = fi107
            }
            v106 = (-1 > 152)
        }
        var v108 uint64 = 42
        _ = v106
        _ = v108
    }
    v29 = float32(-2.25)
    v4 += (v4 / (((v4 >> ((v1) & 63))) | 1))
    fmt.Println(v29, v2, v17)
    v29 += 1.5
    _ = len(v3)
    for fi109 := 0; fi109 < 6; fi109++ {
        if false {
            if (((-9223372036854775808 >= fi109) && true) || (-1 > v4)) {
                for fi110 := 0; fi110 < 12; fi110++ {
                    switch "a" {
                    case "hello":
                        v111 := GenS6{F0: 255, F1: (v17 / (((v17 / ((v17) | 1))) | 1)), F2: 0}
                        _ = v111
                    case "a":
                        v112 := GenS5{F0: 4294967295, F1: 42, F2: 0}
                        _ = v112
                    case "":
                        fmt.Println(v17, fi109)
                    default:
                        var v113 uint = 100
                        _ = v113
                    }
                    v4 = v4
                    v114 := GenS16{F0: 2147483647, F1: 42}
                    _ = fi110
                    _ = v114
                }
                var v115 float64 = 1e10
                v116 := GenS3{F0: v17, F1: 0, F2: -519}
                _ = v115
                _ = v116
            } else {
                v117 := GenS5{F0: 65535, F1: 255, F2: 1}
                v1 = v1
                var v118 float32 = v2
                _ = v117
                _ = v118
            }
            if (true && true) {
                var v119 string = "hello"
                _ = v119
            } else {
                for fi120 := 0; fi120 < 20; fi120++ {
                    v121 := GenS12{F0: 262}
                    var v122 float64 = 0.0
                    _ = fi120
                    _ = v121
                    _ = v122
                }
                v123 := GenS17{F0: (fi109 * (fi109 >> (((^(v1))) & 63)))}
                if (3.14159 <= -2.25) {
                    v124 := []int{genFn1((v17 << (((^(v1))) & 63)), 1), (fi109 & (fi109 & fi109)), fi109, fi109}
                    v29 += v29
                    var v125 int8 = v4
                    _ = v124
                    _ = v125
                } else {
                    var v126 uint16 = v17
                    for fi127 := 0; fi127 < 9; fi127++ {
                        var v128 uint8 = 128
                        _ = fi127
                        _ = v128
                    }
                    if true {
                        var v129 uint32 = 65535
                        v130 := GenS9{F0: (v126 - (v126 % (((v17 & (^(v17)))) | 1))), F1: 0, F2: v4}
                        var v131 float32 = v2
                        _ = v129
                        _ = v130
                        _ = v131
                    }
                    _ = v126
                }
                _ = v123
            }
            var v132 bool = false
            _ = v132
        }
        for fi133 := 0; fi133 < 4; fi133++ {
            var v134 bool = (v17 >= 277)
            v4 -= 127
            switch (fi109 << ((v1) & 63)) {
            case 9223372036854775807:
                var v135 int64 = 1
                _ = v135
            default:
                fmt.Println(v29)
            }
            _ = fi133
            _ = v134
        }
        _ = fi109
    }
    v136 := GenS15{F0: 0}
    v2 = 3.14159
    v3[1] = true
    v137 := GenS7{F0: 2, F1: 0}
    v2 = (v2 / (((float32(float64((v2 * v29))) + v29) + (((v29 + ((((((0.1 + v2) + v29) - v29) + (v2 + float32(1e10))) / 1e10) - v2)) + 1.5) - v2)) * 0.1))
    for fi138 := 0; fi138 < 1; fi138++ {
        switch (fi138 & fi138) {
        case 4294967295:
            switch "hello" {
            case "a":
                if (v1 == (v1 >> ((v1) & 63))) {
                    var v139 uint8 = 15
                    fmt.Println(v1, v4)
                    switch "\n\t" {
                    case "hello":
                        fmt.Println(fi138, v4, v1)
                    }
                    _ = v139
                }
            default:
                v140 := v3[uint32(v1)]
                _ = v140
            }
        case 42:
            if (((^(v1)) >= 65535) && (!(!true))) {
                var v141 uint8 = 42
                if (-9223372036854775808 == -1) {
                    for fi142 := 0; fi142 < 15; fi142++ {
                        v143 := GenS13{F0: 4294967295}
                        v141 = (^(v141))
                        switch (fi138 - -9223372036854775808) {
                        case -2147483648:
                            if true {
                                switch ("hello" + "a") {
                                case "a":
                                    var v144 int16 = -32768
                                    _ = v144
                                case "":
                                    v145 := GenS2{F0: 42}
                                    _ = v145
                                case "hello":
                                    if true {
                                        v17 = v17
                                        v146 := GenS16{F0: 2147483647, F1: 42}
                                        if true {
                                            switch "\n\t" {
                                            case "\n\t":
                                                var v147 uint8 = 178
                                                _ = v147
                                            case "hello":
                                                switch genFn1((^(v17)), v141) {
                                                case -12:
                                                    v29 -= 0.0
                                                }
                                            default:
                                                var v148 uint8 = (v141 | 1)
                                                _ = v148
                                            }
                                            switch genFn1((v17 - v17), 1) {
                                            case 9223372036854775807:
                                                var v149 uint64 = 0
                                                _ = v149
                                            case 128:
                                                v2 = v29
                                            case 32767:
                                                if true {
                                                    if ((v2 / (float32(3.14159) + v29)) >= (0.1 + float32(0.1))) {
                                                        if (0 > 801) {
                                                            v3[42] = (((0.0 <= float64((v29 * v29))) || (42 > 2)) || true)
                                                            var v150 bool = (4294967295 < 1)
                                                            switch (fi138 * 32767) {
                                                            case -360:
                                                                v151 := GenS18{F0: -9223372036854775808}
                                                                _ = v151
                                                            }
                                                            _ = v150
                                                        } else {
                                                            if true {
                                                                if true {
                                                                    switch 847 {
                                                                    case -1:
                                                                        v152 := []int{fi138}
                                                                        _ = v152
                                                                    case 128:
                                                                        var v153 uint32 = 2147483647
                                                                        _ = v153
                                                                    default:
                                                                        v141 = uint8(v1)
                                                                        var v154 uint32 = 1
                                                                        var v155 string = "foo bar"
                                                                        _ = v154
                                                                        _ = v155
                                                                    }
                                                                    if true {
                                                                        v1 = v1
                                                                        var v156 int64 = -1
                                                                        _ = v156
                                                                    } else {
                                                                        if (((!((v2 + v29) < (v29 / 1.5))) && false) && false) {
                                                                            v157 := GenS8{F0: "a"}
                                                                            v158 := GenS13{F0: 622}
                                                                            _ = v157
                                                                            _ = v158
                                                                        } else {
                                                                            var v159 int8 = (v4 % ((42) | 1))
                                                                            v160 := []int{(fi138 * (fi142 << ((829) & 63))), fi142, (fi138 & int(v1))}
                                                                            var v161 float32 = (-2.25 - 3.14159)
                                                                            if (true || ((v1 ^ 402) > (v1 & (v1 % (((v1 % ((v1) | 1))) | 1))))) {
                                                                                var v162 uint32 = 1
                                                                                v141 -= (^(v141))
                                                                                _ = v162
                                                                            } else {
                                                                                v163 := GenS19{F0: 42, F1: (v1 > v1)}
                                                                                v164 := GenS20{F0: -992, F1: v17, F2: 1000}
                                                                                var v165 uint16 = (v17 >> ((2) & 63))
                                                                                _ = v163
                                                                                _ = v164
                                                                                _ = v165
                                                                            }
                                                                            _ = v159
                                                                            _ = v160
                                                                            _ = v161
                                                                        }
                                                                        if true {
                                                                            var v166 int = (fi138 | genFn1((v17 * v17), v141))
                                                                            var v167 int16 = 1000
                                                                            fmt.Println(v4)
                                                                            _ = v166
                                                                            _ = v167
                                                                        } else {
                                                                            v168 := GenS15{F0: v141}
                                                                            v29 = v2
                                                                            var v169 float64 = (-2.25 + 1e10)
                                                                            _ = v168
                                                                            _ = v169
                                                                        }
                                                                    }
                                                                    v29 = v2
                                                                }
                                                            }
                                                            var v170 int8 = v4
                                                            v171 := []int{32767, -925, -9223372036854775808, fi138}
                                                            v170 += (^(v170))
                                                            _ = v170
                                                            _ = v171
                                                        }
                                                        var v172 uint = uint(v17)
                                                        _ = v172
                                                    } else {
                                                        v173 := GenS15{F0: (^(v141))}
                                                        fmt.Println(v4, v1, fi138)
                                                        _ = v173
                                                    }
                                                } else {
                                                    fmt.Println(v2, fi138, v4)
                                                    var v174 float64 = 3.14159
                                                    _ = v174
                                                }
                                            default:
                                                v175 := []int{(fi138 ^ -85), (fi142 * (-(fi142))), (fi138 * fi138)}
                                                _ = v175
                                            }
                                        } else {
                                            if false {
                                                if (0 != -325) {
                                                    var v176 float32 = v2
                                                    var v177 uint16 = (^(v17))
                                                    _ = v176
                                                    _ = v177
                                                }
                                                var v178 float64 = 3.14159
                                                _ = len(v3)
                                                _ = v178
                                            } else {
                                                if (!((fi138 - int(v17)) != (fi142 * fi142))) {
                                                    switch fi138 {
                                                    case 9223372036854775807:
                                                        switch "\n\t" {
                                                        case "foo bar":
                                                            fmt.Println(v29, fi138, v2)
                                                        }
                                                    case 861:
                                                        var v179 bool = false
                                                        _ = v179
                                                    case -2147483648:
                                                        switch (^(fi138)) {
                                                        case 9223372036854775807:
                                                            v29 = -2.25
                                                        default:
                                                            if false {
                                                                var v180 int8 = v4
                                                                _ = v180
                                                            }
                                                        }
                                                    default:
                                                        if false {
                                                            v181 := GenS9{F0: v17, F1: -1, F2: v4}
                                                            switch "a" {
                                                            case "":
                                                                var v182 uint = uint(v141)
                                                                v183 := GenS15{F0: 2}
                                                                v182 += 123
                                                                _ = v182
                                                                _ = v183
                                                            case "a":
                                                                if (((v17 / ((v17) | 1)) < (v17 * uint16(v141))) && true) {
                                                                    v184 := []int{9223372036854775807, fi138}
                                                                    _ = v184
                                                                } else {
                                                                    var v185 string = (("foo bar" + "\n\t") + "hello")
                                                                    _ = v185
                                                                }
                                                            case "hello":
                                                                if (!false) {
                                                                    if true {
                                                                        v186 := []int{(fi138 & (fi142 << ((766) & 63))), fi138}
                                                                        switch fi142 {
                                                                        case 1:
                                                                            v2 = 1.5
                                                                        case 670:
                                                                            delete(v3, 642)
                                                                        default:
                                                                            if true {
                                                                                v2 = (-2.25 * (1.5 / (v29 + float32(-2.25))))
                                                                                var v187 string = ""
                                                                                _ = v187
                                                                            }
                                                                        }
                                                                        delete(v3, 65535)
                                                                        _ = v186
                                                                    } else {
                                                                        switch genFn1(v17, (v141 >> (((v1 << ((1) & 63))) & 63))) {
                                                                        case -977:
                                                                            var v188 float64 = 1.5
                                                                            v189 := GenS8{F0: ((("" + ("\n\t" + ("hello" + "hello"))) + ((("" + "") + (("foo bar" + "foo bar") + "\n\t")) + ("foo bar" + ""))) + "a")}
                                                                            v190 := GenS6{F0: v141, F1: v17, F2: 43}
                                                                            var v191 uint16 = v17
                                                                            _ = v188
                                                                            _ = v189
                                                                            _ = v190
                                                                            _ = v191
                                                                        case 1:
                                                                            var v192 int8 = (-(v4))
                                                                            _ = v192
                                                                        case 32767:
                                                                            v193 := []int{genFn1((v17 - v17), (v141 ^ v141)), int(v17), (fi142 / (((fi142 % ((127) | 1))) | 1))}
                                                                            _ = v193
                                                                        default:
                                                                            switch "hello" {
                                                                            case "hello":
                                                                                var v194 int8 = 0
                                                                                _ = v194
                                                                            case "\n\t":
                                                                                switch fi142 {
                                                                                case -9223372036854775808:
                                                                                    var v195 uint64 = 18446744073709551615
                                                                                    _ = v195
                                                                                case -71:
                                                                                    switch "" {
                                                                                    case "\n\t":
                                                                                        fmt.Println(v1)
                                                                                    case "a":
                                                                                        switch ("hello" + "hello") {
                                                                                        case "hello":
                                                                                            v196 := GenS12{F0: 18446744073709551615}
                                                                                            _ = v196
                                                                                        case "\n\t":
                                                                                            var v197 uint32 = 42
                                                                                            _ = v197
                                                                                        case "a":
                                                                                            v141 -= v141
                                                                                        }
                                                                                    case "foo bar":
                                                                                        v198 := GenS20{F0: -1, F1: v17, F2: -526}
                                                                                        v2 = 1.5
                                                                                        if (1 > 198) {
                                                                                            v17 += 42
                                                                                            var v199 uint8 = v141
                                                                                            _ = v199
                                                                                        } else {
                                                                                            v1 = 65535
                                                                                            switch ("foo bar" + "\n\t") {
                                                                                            case "hello":
                                                                                                v3[42] = (1e10 <= 1e10)
                                                                                            default:
                                                                                                var v200 float64 = float64(v29)
                                                                                                _ = v200
                                                                                            }
                                                                                            fmt.Println(v17, v1, v141)
                                                                                        }
                                                                                        _ = v198
                                                                                    }
                                                                                default:
                                                                                    if (!(!(!true))) {
                                                                                        if (!(870 == 32767)) {
                                                                                            v201 := v3[42]
                                                                                            _ = v201
                                                                                        }
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                        v202 := GenS21{F0: "\n\t", F1: 255, F2: -1}
                                                                        if false {
                                                                            fmt.Println(fi138)
                                                                            var v203 float64 = 1.5
                                                                            _ = v203
                                                                        } else {
                                                                            var v204 float64 = 3.14159
                                                                            if (32768 > 0) {
                                                                                v1 = 0
                                                                            }
                                                                            _ = v204
                                                                        }
                                                                        v205 := []int{fi142, int(v4), (^(fi138))}
                                                                        _ = v202
                                                                        _ = v205
                                                                    }
                                                                    if false {
                                                                        v29 = v2
                                                                        v29 -= float32((1.5 / 1.0))
                                                                        switch "a" {
                                                                        case "a":
                                                                            switch (-(fi138)) {
                                                                            case 128:
                                                                                var v206 string = "hello"
                                                                                _ = v206
                                                                            case 32768:
                                                                                v29 = (v29 * (-2.25 / v2))
                                                                            }
                                                                        case "foo bar":
                                                                            fmt.Println(v141, fi138, v2)
                                                                        }
                                                                    }
                                                                } else {
                                                                    v2 = ((((v29 + 1e10) * v29) + -2.25) / (v2 / v2))
                                                                    if (!false) {
                                                                        switch (^(fi138)) {
                                                                        case 42:
                                                                            var v207 int8 = (v4 << ((v1) & 63))
                                                                            _ = v207
                                                                        default:
                                                                            v141 = (v141 | v141)
                                                                        }
                                                                        switch (fi138 - (fi138 % ((genFn1(v17, (v141 & 42))) | 1))) {
                                                                        case 32767:
                                                                            switch genFn1(255, uint8(v17)) {
                                                                            case 4294967295:
                                                                                if false {
                                                                                    var v208 int = (fi142 >> ((2) & 63))
                                                                                    _ = v208
                                                                                } else {
                                                                                    v3[4294967295] = true
                                                                                    v209 := GenS12{F0: 42}
                                                                                    if false {
                                                                                        v29 -= v29
                                                                                        var v210 float64 = (0.0 - (1.5 + 0.1))
                                                                                        fmt.Println(fi138, v29, fi138)
                                                                                        if false {
                                                                                            v211 := GenS22{F0: v141, F1: "a"}
                                                                                            _ = v211
                                                                                        } else {
                                                                                            var v212 uint64 = 42
                                                                                            _ = v212
                                                                                        }
                                                                                        _ = v210
                                                                                    } else {
                                                                                        var v213 float32 = (v2 * ((v29 / float32(float64(v2))) * v29))
                                                                                        if (0.1 >= -2.25) {
                                                                                            v214 := GenS23{F0: (v17 + 661)}
                                                                                            _ = v214
                                                                                        }
                                                                                        if false {
                                                                                            var v215 float32 = v29
                                                                                            fmt.Println(v17, v213, v17)
                                                                                            v216 := GenS1{F0: (^(v4))}
                                                                                            _ = v215
                                                                                            _ = v216
                                                                                        } else {
                                                                                            v141 = (v141 & 127)
                                                                                            v217 := GenS24{F0: 917, F1: 1}
                                                                                            fmt.Println(v1, v213)
                                                                                            _ = v217
                                                                                        }
                                                                                        _ = v213
                                                                                    }
                                                                                    _ = v209
                                                                                }
                                                                            }
                                                                        case 4294967295:
                                                                            switch fi142 {
                                                                            case -2147483648:
                                                                                v218 := []int{genFn1(v17, 252), (fi138 & (^(fi138))), fi142}
                                                                                v4 += (v4 % (((v4 * v4)) | 1))
                                                                                v17 -= v17
                                                                                _ = v218
                                                                            case 140:
                                                                                v17 = 0
                                                                            }
                                                                        case -32:
                                                                            var v219 int16 = -523
                                                                            _ = v219
                                                                        }
                                                                        var v220 uint32 = 241
                                                                        _ = v220
                                                                    } else {
                                                                        if ((v141 * 27) == v141) {
                                                                            v29 -= 0.0
                                                                            v3[2147483647] = (((v29 - 1.5) * v29) == v2)
                                                                            var v221 int = int(v17)
                                                                            _ = v221
                                                                        } else {
                                                                            v17 -= (v17 % ((v17) | 1))
                                                                            v222 := []int{-1}
                                                                            var v223 uint8 = (v141 - v141)
                                                                            var v224 float32 = v2
                                                                            _ = v222
                                                                            _ = v223
                                                                            _ = v224
                                                                        }
                                                                        fmt.Println(fi142, fi138)
                                                                    }
                                                                }
                                                            }
                                                            _ = v181
                                                        } else {
                                                            switch "" {
                                                            case "a":
                                                                switch "hello" {
                                                                case "foo bar":
                                                                    v225 := []int{fi138}
                                                                    _ = v225
                                                                }
                                                            case "":
                                                                v2 = v2
                                                            case "hello":
                                                                delete(v3, 65535)
                                                            }
                                                            v17 = v17
                                                            v226 := GenS25{F0: -1}
                                                            var v227 uint64 = 4294967295
                                                            _ = v226
                                                            _ = v227
                                                        }
                                                    }
                                                } else {
                                                    if (0.1 < float64(((float32(0.1) * v29) * 1.5))) {
                                                        v228 := []int{fi138, fi138, genFn1(1, v141)}
                                                        var v229 float64 = (-2.25 * (0.1 + float64((-2.25 / float32((3.14159 / 1.0))))))
                                                        v2 += v2
                                                        switch (fi138 * (-(fi138))) {
                                                        case 198:
                                                            v230 := GenS26{F0: "\n\t", F1: (v1 - v1), F2: 0}
                                                            _ = v230
                                                        case -490:
                                                            var v231 uint64 = 4294967295
                                                            _ = v231
                                                        default:
                                                            var v232 bool = false
                                                            _ = v232
                                                        }
                                                        _ = v228
                                                        _ = v229
                                                    }
                                                    v2 = v29
                                                    v141 += v141
                                                }
                                                v233 := GenS1{F0: (v4 << ((0) & 63))}
                                                v234 := []int{(fi142 ^ (-(fi142))), (fi142 >> ((v1) & 63)), (fi142 ^ 930), 4294967295}
                                                _ = v233
                                                _ = v234
                                            }
                                            v235 := GenS19{F0: 2147483647, F1: true}
                                            if (4294967295 >= 273) {
                                                var v236 uint64 = 4294967295
                                                v237 := []int{-9223372036854775808, 983}
                                                v236 = (v236 << ((v1) & 63))
                                                v17 = v17
                                                var v238 uint8 = v141
                                                _ = v236
                                                _ = v237
                                                _ = v238
                                            }
                                            _ = v235
                                        }
                                        _ = v146
                                    }
                                }
                                v1 += (v1 % ((v1) | 1))
                                fmt.Println(fi138, fi142, v2)
                            } else {
                                switch "hello" {
                                case "hello":
                                    v17 -= 454
                                default:
                                    delete(v3, 42)
                                    var v239 float64 = float64(0.0)
                                    v141 += 128
                                    v240 := GenS13{F0: 4294967295}
                                    _ = v239
                                    _ = v240
                                }
                            }
                        case 2147483647:
                            v141 = uint8(v17)
                        case 42:
                            var v241 float32 = (float32(-2.25) - 3.14159)
                            _ = v241
                        }
                        _ = fi142
                        _ = v143
                    }
                    if false {
                        for fi242 := 0; fi242 < 4; fi242++ {
                            if false {
                                var v243 uint32 = 2147483647
                                v244 := GenS15{F0: v141}
                                if false {
                                    if (!(!true)) {
                                        v245 := GenS8{F0: "\n\t"}
                                        v4 += (v4 / ((v4) | 1))
                                        _ = v245
                                    }
                                }
                                _ = v243
                                _ = v244
                            }
                            _ = fi242
                        }
                        for fi246 := 0; fi246 < 8; fi246++ {
                            v247 := GenS18{F0: -43}
                            fmt.Println(v1, v4)
                            var v248 int = fi246
                            _ = fi246
                            _ = v247
                            _ = v248
                        }
                        v249 := GenS27{F0: (v1 << ((v1) & 63)), F1: true}
                        _ = v249
                    }
                    v250 := GenS28{F0: 42}
                    _ = v250
                } else {
                    switch "foo bar" {
                    case "foo bar":
                        if (!(32767 != 1)) {
                            fmt.Println(v2, v17, v1)
                            v29 -= v29
                            var v251 uint16 = (v17 % ((v17) | 1))
                            _ = v251
                        } else {
                            delete(v3, 567)
                            v3[2147483648] = (!true)
                        }
                    default:
                        switch (-(fi138)) {
                        case 0:
                            v4 = (v4 ^ 0)
                        case -284:
                            v252 := GenS17{F0: fi138}
                            _ = v252
                        case 42:
                            v2 -= 3.14159
                        }
                    }
                    v253 := GenS29{F0: true}
                    var v254 int16 = 42
                    _ = v253
                    _ = v254
                }
                _ = v141
            }
        }
        _ = fi138
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v136
    _ = v137
    _ = v17
    _ = v2
    _ = v29
    _ = v3
    _ = v4
}
