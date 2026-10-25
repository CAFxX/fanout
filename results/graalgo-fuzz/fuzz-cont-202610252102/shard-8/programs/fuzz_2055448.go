package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 uint8; F1 int64; F2 uint64 }
type GenS4 struct { F0 uint64; F1 uint32; F2 int8 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint16; F1 int32; F2 int8 }
type GenS7 struct { F0 uint64; F1 uint; F2 uint16 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint16; F1 uint64; F2 int8 }
type GenS10 struct { F0 int16; F1 uint16; F2 int16 }
type GenS11 struct { F0 uint8; F1 int8; F2 int }
type GenS12 struct { F0 int }
type GenS13 struct { F0 int64; F1 int16 }
type GenS14 struct { F0 string }
type GenS15 struct { F0 int32 }
type GenS16 struct { F0 uint8; F1 int16 }
type GenS17 struct { F0 uint8; F1 uint8 }
type GenS18 struct { F0 uint64; F1 uint16 }
type GenS19 struct { F0 bool; F1 int8 }
type GenS20 struct { F0 uint32; F1 int }
type GenS21 struct { F0 uint64; F1 uint16 }
type GenS22 struct { F0 string; F1 uint16; F2 bool }
func genFn0() uint8 {
    switch ("\n\t" + "") {
    case "foo bar":
        var v7 int8 = 10
        _ = v7
    case "a":
        switch -733 {
        case 4294967295:
            if (!false) {
                if false {
                    v8 := []int{127, 32767, 709, -1}
                    switch 958 {
                    case 128:
                        var v9 int8 = 17
                        _ = v9
                    default:
                        var v10 uint8 = 2
                        _ = v10
                    }
                    _ = v8
                }
                var v11 uint32 = 1
                v12 := GenS1{F0: (128 < 128)}
                _ = v11
                _ = v12
            } else {
                var v13 uint16 = 1
                v14 := []int{4294967295, 548}
                _ = v13
                _ = v14
            }
        case 0:
            v15 := []int{-599, 32767, 494}
            _ = v15
        }
    case "\n\t":
        if true {
            var v16 float32 = float32(float64(1e10))
            if (false && false) {
                switch -193 {
                case 472:
                    var v17 uint64 = 42
                    _ = v17
                case -2147483648:
                    switch 9223372036854775807 {
                    case -9223372036854775808:
                        for fi18 := 0; fi18 < 13; fi18++ {
                            var v19 uint = 42
                            if true {
                                for fi20 := 0; fi20 < 11; fi20++ {
                                    fmt.Println(v19, v16)
                                    _ = fi20
                                }
                                for fi21 := 0; fi21 < 2; fi21++ {
                                    v22 := []int{}
                                    _ = fi21
                                    _ = v22
                                }
                            } else {
                                v23 := GenS1{F0: false}
                                _ = v23
                            }
                            _ = fi18
                            _ = v19
                        }
                    case 1:
                        v24 := []int{-1}
                        _ = v24
                    }
                }
            } else {
                var v25 int32 = -2147483648
                fmt.Println(v16, v16)
                _ = v25
            }
            switch "foo bar" {
            case "":
                var v26 string = ("" + "\n\t")
                _ = v26
            case "a":
                if (!(-93 < 226)) {
                    var v27 bool = (2 != 18446744073709551615)
                    fmt.Println(v16, v27)
                    v28 := GenS2{F0: 1000, F1: 1}
                    _ = v27
                    _ = v28
                } else {
                    v16 -= v16
                }
            default:
                var v29 int32 = -574
                _ = v29
            }
            _ = v16
        } else {
            v30 := GenS2{F0: 1000, F1: 127}
            _ = v30
        }
    }
    var v31 uint64 = 18446744073709551615
    v32 := []int{}
    _ = v31
    _ = v32
    return 27
}

func main() {
    var v1 uint8 = 189
    v2 := make(map[string]string)
    v3 := make(map[uint32]bool)
    var v4 uint32 = 42
    var v5 uint64 = 2
    var v6 float64 = 0.1
    v6 -= v6
    v5 -= (v5 >> ((656) & 63))
    v1 = 136
    for fi33 := 0; fi33 < 2; fi33++ {
        if false {
            for fi34 := 0; fi34 < 15; fi34++ {
                v35 := GenS1{F0: (v1 <= 0)}
                v36 := GenS3{F0: 42, F1: 42, F2: (v5 ^ uint64(v4))}
                _ = fi34
                _ = v35
                _ = v36
            }
            for fi37 := 0; fi37 < 17; fi37++ {
                if (true || (!true)) {
                    v4 = 0
                    if false {
                        var v38 int = fi37
                        fmt.Println(v5, v1, v6)
                        v39 := GenS4{F0: (v5 | uint64(v4)), F1: (v4 - v4), F2: 1}
                        _ = v38
                        _ = v39
                    }
                    var v40 bool = false
                    _ = v40
                } else {
                    v41 := []int{(-(fi33)), 4294967295}
                    fmt.Println(fi37, fi33)
                    _ = len(v3)
                    _ = v41
                }
                _ = fi37
            }
        } else {
            if false {
                var v42 uint16 = 9
                var v43 float32 = 0.1
                switch "\n\t" {
                case "hello":
                    switch "a" {
                    case "\n\t":
                        var v44 int64 = -9223372036854775808
                        _ = v44
                    }
                }
                _ = v42
                _ = v43
            }
            fmt.Println(v1, fi33)
        }
        _ = fi33
    }
    var v45 int8 = 48
    for fi46 := 0; fi46 < 13; fi46++ {
        v6 -= v6
        _ = fi46
    }
    for fi47 := 0; fi47 < 7; fi47++ {
        v48 := GenS2{F0: 492, F1: v45}
        if (2 != uint16(v5)) {
            for fi49 := 0; fi49 < 2; fi49++ {
                switch 9223372036854775807 {
                case 9223372036854775807:
                    v50 := GenS3{F0: (v1 + (^(v1))), F1: -9223372036854775808, F2: 952}
                    _ = v50
                case 374:
                    v51 := GenS5{F0: 18446744073709551615}
                    _ = v51
                case 844:
                    fmt.Println(v4, v6, v4)
                default:
                    fmt.Println(v1, fi49, fi49)
                }
                delete(v3, (^(v4)))
                _ = fi49
            }
        } else {
            v6 -= (v6 + v6)
            _ = len(v2)
            var v52 uint8 = (v1 | v1)
            _ = v52
        }
        _ = fi47
        _ = v48
    }
    v6 -= (3.14159 - 3.14159)
    for fi53 := 0; fi53 < 13; fi53++ {
        v54 := GenS6{F0: 1, F1: -1, F2: -23}
        _ = len(v2)
        _ = fi53
        _ = v54
    }
    switch ("hello" + "a") {
    case "hello":
        if ((984 > 1) || false) {
            var v55 uint16 = 65535
            v56 := GenS7{F0: (v5 * v5), F1: 567, F2: (v55 ^ v55)}
            _ = v55
            _ = v56
        } else {
            v5 = v5
            if false {
                for fi57 := 0; fi57 < 12; fi57++ {
                    v58 := v2[("" + "a")]
                    for fi59 := 0; fi59 < 5; fi59++ {
                        fmt.Println(v6, v1)
                        _ = fi59
                    }
                    _ = fi57
                    _ = v58
                }
                v1 += v1
                v60 := GenS8{F0: -20}
                _ = v60
            } else {
                switch -9223372036854775808 {
                case -438:
                    _ = len(v3)
                case -962:
                    var v61 int32 = -2147483648
                    _ = v61
                case 478:
                    var v62 float32 = -2.25
                    _ = v62
                default:
                    var v63 int32 = -974
                    _ = v63
                }
                switch "\n\t" {
                case "a":
                    switch "a" {
                    case "a":
                        for fi64 := 0; fi64 < 12; fi64++ {
                            v45 += (v45 & v45)
                            _ = fi64
                        }
                    case "\n\t":
                        switch "" {
                        case "\n\t":
                            v5 = (v5 % (((v5 | (v5 << ((0) & 63)))) | 1))
                        case "foo bar":
                            v65 := GenS2{F0: int16(v5), F1: (v45 ^ 42)}
                            _ = v65
                        case "hello":
                            v66 := []int{127}
                            _ = v66
                        default:
                            if true {
                                if true {
                                    if true {
                                        var v67 bool = (1 != 1)
                                        v68 := GenS2{F0: 32767, F1: v45}
                                        _ = v67
                                        _ = v68
                                    } else {
                                        if (genFn0() <= 0) {
                                            if false {
                                                switch 1 {
                                                case -232:
                                                    var v69 float64 = 3.14159
                                                    _ = v69
                                                }
                                                var v70 uint32 = v4
                                                fmt.Println(v4)
                                                _ = v70
                                            } else {
                                                v45 += -105
                                                fmt.Println(v6, v5, v45)
                                            }
                                            for fi71 := 0; fi71 < 18; fi71++ {
                                                var v72 uint8 = 51
                                                v45 = v45
                                                v6 = v6
                                                _ = fi71
                                                _ = v72
                                            }
                                            var v73 uint8 = (v1 / ((v1) | 1))
                                            _ = v73
                                        } else {
                                            switch 127 {
                                            case 127:
                                                v74 := GenS9{F0: 1, F1: 4294967295, F2: (v45 << ((0) & 63))}
                                                _ = v74
                                            }
                                            var v75 bool = (!false)
                                            _ = v75
                                        }
                                    }
                                    delete(v2, "")
                                    v76 := []int{9223372036854775807}
                                    _ = v76
                                }
                                var v77 int32 = 2147483647
                                _ = v77
                            }
                        }
                    }
                case "foo bar":
                    var v78 int8 = v45
                    _ = v78
                case "hello":
                    switch -956 {
                    case -2147483648:
                        switch ((("foo bar" + "a") + "\n\t") + ("a" + "a")) {
                        case "":
                            v79 := []int{0}
                            _ = v79
                        case "\n\t":
                            var v80 bool = false
                            _ = v80
                        case "hello":
                            v81 := GenS10{F0: 42, F1: 42, F2: -32768}
                            _ = v81
                        }
                    case -272:
                        var v82 float32 = float32(v6)
                        _ = v82
                    default:
                        if true {
                            if (false && (true && true)) {
                                for fi83 := 0; fi83 < 20; fi83++ {
                                    v5 = (v5 + v5)
                                    v84 := GenS11{F0: (v1 % ((v1) | 1)), F1: (v45 + -107), F2: fi83}
                                    _ = fi83
                                    _ = v84
                                }
                                if false {
                                    fmt.Println(v4, v5, v6)
                                    for fi85 := 0; fi85 < 7; fi85++ {
                                        switch "" {
                                        case "hello":
                                            for fi86 := 0; fi86 < 12; fi86++ {
                                                v87 := []int{}
                                                _ = fi86
                                                _ = v87
                                            }
                                        case "":
                                            for fi88 := 0; fi88 < 10; fi88++ {
                                                var v89 int64 = 706
                                                _ = fi88
                                                _ = v89
                                            }
                                        case "foo bar":
                                            v45 = 127
                                        }
                                        _ = fi85
                                    }
                                    v4 += (v4 / (((^(v4))) | 1))
                                }
                                fmt.Println(v5, v1, v4)
                            } else {
                                for fi90 := 0; fi90 < 16; fi90++ {
                                    v91 := GenS3{F0: 123, F1: 42, F2: v5}
                                    _ = fi90
                                    _ = v91
                                }
                                v92 := GenS12{F0: 2147483647}
                                _ = v92
                            }
                            for fi93 := 0; fi93 < 19; fi93++ {
                                for fi94 := 0; fi94 < 9; fi94++ {
                                    var v95 int16 = -341
                                    if true {
                                        v1 += v1
                                        var v96 string = "\n\t"
                                        v97 := []int{fi94, int(v45)}
                                        _ = v96
                                        _ = v97
                                    } else {
                                        switch "foo bar" {
                                        case "":
                                            fmt.Println(fi93)
                                        case "a":
                                            fmt.Println(fi94)
                                        case "\n\t":
                                            var v98 uint8 = genFn0()
                                            _ = v98
                                        }
                                        var v99 int = (fi94 | (fi94 & int(v45)))
                                        v4 = v4
                                        _ = v99
                                    }
                                    switch fi94 {
                                    case 9223372036854775807:
                                        fmt.Println(fi94, v95)
                                        var v100 string = "\n\t"
                                        _ = v100
                                    default:
                                        var v101 uint8 = 2
                                        _ = v101
                                    }
                                    _ = fi94
                                    _ = v95
                                }
                                fmt.Println(v6)
                                v5 = v5
                                _ = fi93
                            }
                            if true {
                                v102 := v3[v4]
                                _ = v102
                            } else {
                                if false {
                                    v45 -= (v45 - v45)
                                    for fi103 := 0; fi103 < 13; fi103++ {
                                        var v104 int = 2147483647
                                        if false {
                                            v105 := GenS13{F0: 42, F1: -1}
                                            _ = v105
                                        }
                                        _ = fi103
                                        _ = v104
                                    }
                                    v106 := GenS8{F0: 278}
                                    _ = v106
                                }
                                for fi107 := 0; fi107 < 4; fi107++ {
                                    v108 := GenS13{F0: 713, F1: int16(v45)}
                                    var v109 uint = 150
                                    _ = fi107
                                    _ = v108
                                    _ = v109
                                }
                            }
                        } else {
                            for fi110 := 0; fi110 < 1; fi110++ {
                                var v111 uint = 0
                                var v112 int = (-(fi110))
                                _ = fi110
                                _ = v111
                                _ = v112
                            }
                        }
                    }
                }
            }
        }
    case "\n\t":
        if false {
            v113 := GenS2{F0: -32768, F1: v45}
            v114 := []int{int(v5), 9223372036854775807, -2147483648, -964}
            _ = v113
            _ = v114
        }
    case "":
        if (3.14159 != 0.1) {
            var v115 int64 = -9223372036854775808
            _ = v115
        } else {
            if (126 >= -1) {
                v116 := v2["hello"]
                fmt.Println(v5, v6)
                _ = v116
            }
            if true {
                if true {
                    if true {
                        if false {
                            v117 := []int{-9223372036854775808, -1}
                            _ = v117
                        } else {
                            for fi118 := 0; fi118 < 10; fi118++ {
                                var v119 uint16 = 0
                                var v120 float64 = (v6 / 1.0)
                                _ = fi118
                                _ = v119
                                _ = v120
                            }
                        }
                        if (42 == 1) {
                            v4 = 283
                            v45 = -1
                        } else {
                            if (1e10 == float32(v6)) {
                                var v121 float32 = (1.5 * 1e10)
                                fmt.Println(v4, v45, v5)
                                if true {
                                    var v122 int8 = (v45 ^ v45)
                                    switch "\n\t" {
                                    case "a":
                                        v5 = (v5 / ((v5) | 1))
                                    default:
                                        v123 := []int{9223372036854775807, 9223372036854775807, int(v45)}
                                        _ = v123
                                    }
                                    for fi124 := 0; fi124 < 2; fi124++ {
                                        switch -926 {
                                        case -2147483648:
                                            if (v45 > v45) {
                                                v125 := GenS12{F0: fi124}
                                                var v126 int8 = (v122 & v45)
                                                _ = v125
                                                _ = v126
                                            }
                                        case 128:
                                            if (!((v1 | v1) >= uint8(v5))) {
                                                for fi127 := 0; fi127 < 16; fi127++ {
                                                    v121 -= v121
                                                    var v128 float64 = float64(v121)
                                                    switch 32767 {
                                                    case 4294967295:
                                                        v121 = float32((v128 / v128))
                                                    case 393:
                                                        var v129 float64 = 1e10
                                                        _ = v129
                                                    default:
                                                        if false {
                                                            if (!true) {
                                                                v130 := GenS14{F0: ("\n\t" + ("foo bar" + "foo bar"))}
                                                                if (int64(fi124) != -9223372036854775808) {
                                                                    if true {
                                                                        var v131 float64 = v6
                                                                        if true {
                                                                            var v132 uint32 = (^(v4))
                                                                            v133 := GenS6{F0: 299, F1: -1, F2: (v122 % (((v122 * (v122 ^ 1))) | 1))}
                                                                            _ = v132
                                                                            _ = v133
                                                                        } else {
                                                                            v122 = 20
                                                                            _ = len(v2)
                                                                            if false {
                                                                                switch fi124 {
                                                                                case -129:
                                                                                    var v134 uint64 = (v5 - 42)
                                                                                    _ = v134
                                                                                default:
                                                                                    v1 = v1
                                                                                }
                                                                                v121 = v121
                                                                                v1 = 2
                                                                                switch (fi127 | (fi124 + fi127)) {
                                                                                case 32767:
                                                                                    if true {
                                                                                        if false {
                                                                                            fmt.Println(v4, v4, fi127)
                                                                                            v135 := []int{}
                                                                                            _ = v135
                                                                                        }
                                                                                        v136 := GenS15{F0: 544}
                                                                                        _ = v136
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                var v137 bool = false
                                                                                v138 := []int{}
                                                                                v139 := GenS16{F0: genFn0(), F1: -32768}
                                                                                _ = v137
                                                                                _ = v138
                                                                                _ = v139
                                                                            }
                                                                            var v140 int32 = -238
                                                                            _ = v140
                                                                        }
                                                                        _ = v131
                                                                    } else {
                                                                        if ((v1 * v1) >= uint8(v5)) {
                                                                            switch fi127 {
                                                                            case -2147483648:
                                                                                switch "\n\t" {
                                                                                case "":
                                                                                    fmt.Println(v4, v1, fi127)
                                                                                case "\n\t":
                                                                                    var v141 uint32 = v4
                                                                                    _ = v141
                                                                                case "hello":
                                                                                    v142 := GenS17{F0: 128, F1: v1}
                                                                                    _ = v142
                                                                                }
                                                                            }
                                                                            var v143 uint = 2
                                                                            _ = v143
                                                                        } else {
                                                                            var v144 int = (fi124 ^ (fi124 - fi127))
                                                                            v145 := []int{-11, (-(v144)), 0}
                                                                            switch (v144 * v144) {
                                                                            case 32767:
                                                                                fmt.Println(fi127, v45, v4)
                                                                            case 95:
                                                                                if true {
                                                                                    fmt.Println(v45, v45)
                                                                                    var v146 int = (fi124 << ((0) & 63))
                                                                                    v145 = append(v145, v144)
                                                                                    _ = v146
                                                                                }
                                                                            case 42:
                                                                                v147 := GenS11{F0: (v1 >> ((65535) & 63)), F1: (v122 >> ((255) & 63)), F2: fi127}
                                                                                _ = v147
                                                                            }
                                                                            _ = v144
                                                                            _ = v145
                                                                        }
                                                                        fmt.Println(fi127, v1)
                                                                    }
                                                                    var v148 int = fi127
                                                                    switch "a" {
                                                                    case "foo bar":
                                                                        v122 += v122
                                                                    default:
                                                                        var v149 uint16 = 32767
                                                                        _ = v149
                                                                    }
                                                                    _ = v148
                                                                } else {
                                                                    if (!(true && false)) {
                                                                        v6 += ((1.5 / 1.0) * v128)
                                                                        v3[uint32(v1)] = (true || true)
                                                                    }
                                                                }
                                                                switch "hello" {
                                                                case "\n\t":
                                                                    v150 := []int{}
                                                                    _ = v150
                                                                default:
                                                                    v151 := GenS18{F0: uint64(v1), F1: 272}
                                                                    v6 -= v128
                                                                    fmt.Println(v128, v1, fi124)
                                                                    _ = v151
                                                                }
                                                                _ = v130
                                                            }
                                                            switch (fi124 * fi127) {
                                                            case 32768:
                                                                delete(v3, (v4 << ((42) & 63)))
                                                            default:
                                                                v45 = 42
                                                            }
                                                        }
                                                    }
                                                    _ = fi127
                                                    _ = v128
                                                }
                                                switch ("foo bar" + ("a" + ("foo bar" + "foo bar"))) {
                                                case "a":
                                                    if true {
                                                        var v152 int8 = (v122 + -128)
                                                        _ = v152
                                                    } else {
                                                        v153 := GenS3{F0: v1, F1: -921, F2: (v5 % ((v5) | 1))}
                                                        _ = v153
                                                    }
                                                case "\n\t":
                                                    var v154 uint64 = v5
                                                    _ = v154
                                                case "":
                                                    switch "" {
                                                    case "foo bar":
                                                        v4 = v4
                                                    default:
                                                        v155 := GenS19{F0: (0 > 427), F1: (v122 % ((109) | 1))}
                                                        _ = v155
                                                    }
                                                default:
                                                    v122 += v122
                                                }
                                                for fi156 := 0; fi156 < 20; fi156++ {
                                                    fmt.Println(v45)
                                                    _ = fi156
                                                }
                                            } else {
                                                if false {
                                                    v121 -= -2.25
                                                    v157 := GenS20{F0: v4, F1: fi124}
                                                    _ = v157
                                                }
                                                for fi158 := 0; fi158 < 8; fi158++ {
                                                    v6 -= v6
                                                    v4 = (v4 + (v4 / ((v4) | 1)))
                                                    _ = fi158
                                                }
                                            }
                                        case 127:
                                            var v159 uint8 = (v1 % ((0) | 1))
                                            _ = v159
                                        default:
                                            var v160 int8 = (v122 % (((v122 - 0)) | 1))
                                            _ = v160
                                        }
                                        v1 -= v1
                                        var v161 uint8 = v1
                                        _ = fi124
                                        _ = v161
                                    }
                                    _ = v122
                                } else {
                                    var v162 float32 = 3.14159
                                    for fi163 := 0; fi163 < 4; fi163++ {
                                        v45 = (v45 * (v45 - v45))
                                        v164 := GenS1{F0: (!true)}
                                        _ = fi163
                                        _ = v164
                                    }
                                    v6 = float64(v162)
                                    _ = v162
                                }
                                _ = v121
                            } else {
                                var v165 int16 = 484
                                _ = v165
                            }
                            switch "a" {
                            case "":
                                for fi166 := 0; fi166 < 3; fi166++ {
                                    fmt.Println(v6, v45)
                                    for fi167 := 0; fi167 < 17; fi167++ {
                                        switch int(v4) {
                                        case 777:
                                            switch fi166 {
                                            case 0:
                                                v45 += 127
                                            case 127:
                                                v168 := GenS17{F0: 1, F1: v1}
                                                _ = v168
                                            default:
                                                v5 = 415
                                            }
                                        default:
                                            var v169 int = fi166
                                            _ = v169
                                        }
                                        v170 := GenS10{F0: 42, F1: 336, F2: -32768}
                                        var v171 uint16 = 65535
                                        _ = fi167
                                        _ = v170
                                        _ = v171
                                    }
                                    _ = fi166
                                }
                            case "foo bar":
                                v6 += float64(1e10)
                            case "\n\t":
                                var v172 float32 = 0.0
                                _ = v172
                            }
                            for fi173 := 0; fi173 < 9; fi173++ {
                                v174 := GenS14{F0: "a"}
                                for fi175 := 0; fi175 < 8; fi175++ {
                                    var v176 float32 = -2.25
                                    v6 += float64(((1.5 - (v176 / v176)) / ((v176 + (((v176 + 0.1) / v176) / 1e10)) - ((1e10 + 0.0) - v176))))
                                    fmt.Println(v5)
                                    _ = fi175
                                    _ = v176
                                }
                                switch ("a" + ("a" + ("" + ("foo bar" + "hello")))) {
                                case "hello":
                                    _ = len(v3)
                                case "foo bar":
                                    if false {
                                        v6 = v6
                                        for fi177 := 0; fi177 < 18; fi177++ {
                                            v178 := GenS13{F0: 9223372036854775807, F1: 32767}
                                            _ = fi177
                                            _ = v178
                                        }
                                    }
                                }
                                _ = fi173
                                _ = v174
                            }
                        }
                        var v179 uint = 1
                        _ = v179
                    }
                    if true {
                        v5 = (v5 << ((632) & 63))
                        fmt.Println(v5)
                        if true {
                            var v180 int64 = -424
                            v5 = (v5 << ((766) & 63))
                            if (v45 < 127) {
                                switch (("\n\t" + "hello") + "foo bar") {
                                case "":
                                    for fi181 := 0; fi181 < 19; fi181++ {
                                        if true {
                                            fmt.Println(v6)
                                            switch fi181 {
                                            case 1:
                                                _ = len(v2)
                                            case -9223372036854775808:
                                                var v182 bool = false
                                                _ = v182
                                            case 951:
                                                v180 -= 42
                                            default:
                                                fmt.Println(v180)
                                            }
                                        }
                                        var v183 float64 = (v6 + v6)
                                        _ = fi181
                                        _ = v183
                                    }
                                case "foo bar":
                                    var v184 float64 = (v6 / v6)
                                    _ = v184
                                }
                                fmt.Println(v5)
                            } else {
                                fmt.Println(v45, v180, v1)
                                v185 := GenS21{F0: v5, F1: 0}
                                _ = v185
                            }
                            _ = v180
                        }
                    } else {
                        if (42 <= -627) {
                            var v186 uint16 = 672
                            v186 += 32767
                            v4 = v4
                            _ = v186
                        } else {
                            fmt.Println(v45)
                            v187 := []int{914, 4294967295, 717, -760}
                            var v188 int = -9223372036854775808
                            _ = v187
                            _ = v188
                        }
                        var v189 int8 = v45
                        var v190 float64 = v6
                        _ = v189
                        _ = v190
                    }
                }
                v191 := []int{433, 127, 2147483647, 1}
                _ = v191
            } else {
                for fi192 := 0; fi192 < 7; fi192++ {
                    v5 += (v5 << ((255) & 63))
                    v193 := GenS22{F0: "foo bar", F1: 255, F2: true}
                    var v194 int = 211
                    _ = fi192
                    _ = v193
                    _ = v194
                }
                for fi195 := 0; fi195 < 11; fi195++ {
                    for fi196 := 0; fi196 < 10; fi196++ {
                        v45 -= (v45 & (v45 ^ v45))
                        fmt.Println(v45, fi195)
                        switch ("a" + "\n\t") {
                        case "hello":
                            v45 -= -1
                        }
                        switch fi196 {
                        case -1:
                            v45 = (^(v45))
                        }
                        var v197 string = "\n\t"
                        var v198 float64 = 1e10
                        var v199 uint8 = genFn0()
                        switch "" {
                        case "hello":
                            var v200 int32 = 42
                            _ = v200
                        }
                        switch "" {
                        case "a":
                            v201 := []int{}
                            _ = v201
                        case "\n\t":
                            v1 = (v199 | v1)
                        }
                        _ = fi196
                        _ = v197
                        _ = v198
                        _ = v199
                    }
                    _ = len(v3)
                    _ = fi195
                }
                for fi202 := 0; fi202 < 10; fi202++ {
                    delete(v2, ("\n\t" + "hello"))
                    v45 -= v45
                    _ = fi202
                }
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v6
}
