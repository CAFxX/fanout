package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 int64; F1 uint16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int32; F1 int64 }
type GenS7 struct { F0 int; F1 int16 }
type GenS8 struct { F0 uint32; F1 uint32 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 uint32; F1 int16 }
type GenS11 struct { F0 uint64; F1 uint16; F2 int32 }
type GenS12 struct { F0 int; F1 uint8; F2 int64 }
type GenS13 struct { F0 int32; F1 string }
type GenS14 struct { F0 uint32; F1 uint8; F2 int16 }
type GenS15 struct { F0 uint16; F1 int64 }
type GenS16 struct { F0 int8 }
type GenS17 struct { F0 uint8 }
type GenS18 struct { F0 string; F1 bool; F2 uint8 }
type GenS19 struct { F0 uint64 }
func genFn0() uint8 {
    switch "\n\t" {
    case "foo bar":
        var v4 int16 = 1
        _ = v4
    default:
        for fi5 := 0; fi5 < 1; fi5++ {
            for fi6 := 0; fi6 < 2; fi6++ {
                if false {
                    var v7 string = "a"
                    _ = v7
                } else {
                    fmt.Println(fi5, fi6)
                }
                v8 := GenS1{F0: 0}
                _ = fi6
                _ = v8
            }
            _ = fi5
        }
    }
    var v9 int16 = 1
    _ = v9
    return 2
}
func genFn1(p0 bool) uint32 {
    if (-1 != -128) {
        v10 := GenS2{F0: 32767, F1: 832}
        _ = v10
    } else {
        var v11 float32 = (0.0 / 0.1)
        v11 = v11
        _ = v11
    }
    if (false && (87 < 772)) {
        var v12 int8 = -56
        if (v12 < (v12 % (((-(v12))) | 1))) {
            for fi13 := 0; fi13 < 8; fi13++ {
                if p0 {
                    switch ("\n\t" + "\n\t") {
                    case "hello":
                        v12 = v12
                    }
                    for fi14 := 0; fi14 < 20; fi14++ {
                        var v15 int64 = 1
                        _ = fi14
                        _ = v15
                    }
                    p0 = p0
                }
                _ = fi13
            }
        } else {
            switch "foo bar" {
            case "\n\t":
                v16 := []int{0, -2147483648, 0, 9223372036854775807}
                _ = v16
            default:
                if (!p0) {
                    var v17 int8 = (v12 / ((-123) | 1))
                    _ = v17
                }
            }
            var v18 int = -2147483648
            _ = v18
        }
        if p0 {
            var v19 float32 = 0.1
            _ = v19
        } else {
            if p0 {
                switch "\n\t" {
                case "foo bar":
                    if (genFn0() > 128) {
                        p0 = (!(true || (!p0)))
                        switch "\n\t" {
                        case "foo bar":
                            v12 += v12
                        case "hello":
                            var v20 uint32 = 42
                            _ = v20
                        case "a":
                            v21 := GenS1{F0: 65535}
                            _ = v21
                        }
                        p0 = p0
                    } else {
                        v12 = 100
                    }
                }
                p0 = (42 != 453)
                var v22 string = ("a" + "a")
                _ = v22
            }
            v23 := []int{-9223372036854775808, -863}
            _ = v23
        }
        _ = v12
    } else {
        var v24 float64 = float64((1e10 * 1.5))
        var v25 int = -479
        v26 := GenS1{F0: 2}
        _ = v24
        _ = v25
        _ = v26
    }
    _ = p0
    return 42
}

func main() {
    var v1 uint32 = 464
    var v2 float32 = float32((3.14159 + 0.1))
    v3 := make(map[string]string)
    v1 -= (v1 >> ((18446744073709551615) & 63))
    v27 := GenS3{F0: -1, F1: 2}
    if false {
        if (2 > 2) {
            if (!(1 != 0)) {
                if (true && (255 > 0)) {
                    var v28 int8 = -128
                    v29 := GenS4{F0: 1}
                    var v30 int8 = v28
                    _ = v28
                    _ = v29
                    _ = v30
                }
                if (false && true) {
                    if ((!(!((v1 + 0) >= 578))) && false) {
                        var v31 bool = true
                        _ = v31
                    } else {
                        switch 127 {
                        case -3:
                            for fi32 := 0; fi32 < 2; fi32++ {
                                delete(v3, "foo bar")
                                v33 := GenS5{F0: 588}
                                _ = fi32
                                _ = v33
                            }
                        case -1:
                            var v34 float64 = -2.25
                            _ = v34
                        }
                    }
                    switch 335 {
                    case 4294967295:
                        switch ("\n\t" + "a") {
                        case "\n\t":
                            switch 127 {
                            case 32767:
                                v35 := []int{-2147483648, 9223372036854775807, 32767, 539}
                                _ = v35
                            }
                        }
                    case 225:
                        if (!("hello" >= "\n\t")) {
                            v3[""] = "hello"
                            fmt.Println(v2, v1)
                            v36 := []int{9223372036854775807, 0, 2147483647}
                            _ = v36
                        }
                    }
                } else {
                    fmt.Println(v1, v2)
                    fmt.Println(v1, v1)
                    v37 := GenS5{F0: 683}
                    _ = v37
                }
                var v38 float32 = v2
                _ = v38
            } else {
                fmt.Println(v1, v2)
                for fi39 := 0; fi39 < 12; fi39++ {
                    if true {
                        var v40 float32 = v2
                        if true {
                            v41 := []int{}
                            if false {
                                for fi42 := 0; fi42 < 13; fi42++ {
                                    v3[("hello" + ("\n\t" + "\n\t"))] = "a"
                                    var v43 int = (fi39 << ((4294967295) & 63))
                                    _ = fi42
                                    _ = v43
                                }
                                var v44 int = 32767
                                _ = v44
                            } else {
                                v45 := v3[("foo bar" + "foo bar")]
                                var v46 bool = false
                                v2 = v40
                                _ = v45
                                _ = v46
                            }
                            if (true || (false || true)) {
                                v47 := GenS4{F0: 634}
                                _ = v47
                            }
                            _ = v41
                        } else {
                            _ = len(v3)
                            fmt.Println(v40, v1, fi39)
                        }
                        for fi48 := 0; fi48 < 19; fi48++ {
                            v1 = 2147483648
                            v49 := []int{554}
                            _ = fi48
                            _ = v49
                        }
                        _ = v40
                    } else {
                        for fi50 := 0; fi50 < 10; fi50++ {
                            v2 = ((v2 + v2) / v2)
                            v51 := []int{(fi50 << ((4294967295) & 63)), (fi39 * (^(fi39)))}
                            var v52 uint = 741
                            _ = fi50
                            _ = v51
                            _ = v52
                        }
                        v53 := []int{(fi39 % (((fi39 / ((fi39) | 1))) | 1)), fi39}
                        fmt.Println(fi39, v1, fi39)
                        _ = v53
                    }
                    for fi54 := 0; fi54 < 20; fi54++ {
                        if (-575 < 0) {
                            var v55 int8 = 0
                            _ = v55
                        } else {
                            v1 += genFn1(false)
                        }
                        _ = fi54
                    }
                    _ = fi39
                }
                fmt.Println(v2)
            }
        } else {
            v2 = (v2 / 1e10)
            v56 := []int{32768, 128, -765, -482}
            v1 += v1
            _ = v56
        }
        v57 := GenS6{F0: 6, F1: int64(v1)}
        if false {
            v2 += 1e10
            v58 := GenS7{F0: -2147483648, F1: 1}
            _ = v58
        }
        _ = v57
    }
    delete(v3, (("\n\t" + ("" + "foo bar")) + "hello"))
    fmt.Println(v1)
    if true {
        v59 := GenS8{F0: 4294967295, F1: 2147483648}
        var v60 string = "foo bar"
        var v61 int64 = 9223372036854775807
        _ = v59
        _ = v60
        _ = v61
    }
    v2 -= v2
    v62 := GenS2{F0: -251, F1: -620}
    if true {
        if false {
            var v63 uint16 = 10
            _ = v63
        }
        var v64 int32 = -2147483648
        v65 := GenS9{F0: 1}
        _ = v64
        _ = v65
    }
    var v66 int = 1
    if true {
        var v67 int16 = 32767
        var v68 uint = 313
        var v69 int32 = -1
        _ = v67
        _ = v68
        _ = v69
    }
    fmt.Println(v2, v66)
    for fi70 := 0; fi70 < 12; fi70++ {
        if true {
            v2 -= v2
            v66 = -898
            var v71 float64 = 3.14159
            _ = v71
        } else {
            fmt.Println(v1)
            switch 9223372036854775807 {
            case 128:
                var v72 float32 = float32(1.5)
                _ = v72
            case 4294967295:
                if false {
                    v73 := []int{v66, -68}
                    for fi74 := 0; fi74 < 4; fi74++ {
                        v75 := GenS5{F0: 18446744073709551615}
                        _ = fi74
                        _ = v75
                    }
                    for fi76 := 0; fi76 < 19; fi76++ {
                        _ = len(v73)
                        _ = len(v73)
                        _ = fi76
                    }
                    _ = v73
                }
            default:
                switch fi70 {
                case 32768:
                    v77 := GenS3{F0: -1, F1: 948}
                    _ = v77
                }
            }
        }
        var v78 uint = 65535
        _ = fi70
        _ = v78
    }
    switch (v66 ^ (v66 + v66)) {
    case 373:
        if false {
            if false {
                var v79 bool = true
                _ = v79
            } else {
                for fi80 := 0; fi80 < 20; fi80++ {
                    for fi81 := 0; fi81 < 11; fi81++ {
                        v82 := []int{(v66 % ((32768) | 1)), 1000, -1, 904}
                        v2 += v2
                        if ((-10 > 1) && (((v1 | v1) > genFn1(true)) || ((v1 | 65535) >= 2147483647))) {
                            if (32767 == 1000) {
                                v1 -= v1
                                var v83 int8 = -128
                                if false {
                                    if false {
                                        v83 -= (v83 >> ((42) & 63))
                                        if false {
                                            v83 = v83
                                            var v84 float64 = (3.14159 - 0.1)
                                            _ = v84
                                        } else {
                                            var v85 uint16 = 1
                                            switch "hello" {
                                            case "a":
                                                var v86 bool = true
                                                _ = v86
                                            case "":
                                                v87 := v3[""]
                                                _ = v87
                                            case "hello":
                                                var v88 float32 = v2
                                                _ = v88
                                            default:
                                                v3[("hello" + "\n\t")] = "hello"
                                                v2 -= (((float32((-2.25 - (3.14159 * (0.0 + 3.14159)))) / ((v2 + v2) - (float32(3.14159) - 0.1))) - 0.0) / v2)
                                                fmt.Println(v1)
                                            }
                                            _ = v85
                                        }
                                        var v89 uint64 = 42
                                        _ = v89
                                    }
                                    v90 := v3["hello"]
                                    _ = v90
                                } else {
                                    var v91 float32 = (1.5 * v2)
                                    var v92 int64 = 1
                                    switch ("foo bar" + ("hello" + ((("" + ("\n\t" + ("a" + ""))) + "hello") + ("foo bar" + "hello")))) {
                                    case "a":
                                        v1 -= 2147483648
                                    case "\n\t":
                                        var v93 uint = 721
                                        _ = v93
                                    case "foo bar":
                                        v92 += (v92 - (v92 << ((4294967295) & 63)))
                                    default:
                                        switch "" {
                                        case "a":
                                            var v94 int32 = 2147483647
                                            _ = v94
                                        case "\n\t":
                                            v92 += 1
                                        default:
                                            v82 = append(v82, fi81)
                                        }
                                    }
                                    v95 := GenS9{F0: -32768}
                                    v83 -= 127
                                    _ = v91
                                    _ = v92
                                    _ = v95
                                }
                                _ = v83
                            } else {
                                if false {
                                    if (2 > 1) {
                                        switch "\n\t" {
                                        case "a":
                                            var v96 uint16 = 65535
                                            _ = v96
                                        case "\n\t":
                                            v1 = genFn1(false)
                                        }
                                        if (1 >= 255) {
                                            v2 += 3.14159
                                            v2 += (v2 + (-2.25 * (((3.14159 * -2.25) + float32(0.1)) / v2)))
                                        }
                                    } else {
                                        var v97 int = (fi81 ^ v66)
                                        _ = v97
                                    }
                                } else {
                                    fmt.Println(fi80, fi80, v2)
                                }
                                switch 4294967295 {
                                case 2147483647:
                                    v98 := GenS3{F0: 9223372036854775807, F1: 32767}
                                    _ = v98
                                }
                                fmt.Println(v2)
                                fmt.Println(fi81, fi80, fi80)
                            }
                            if false {
                                var v99 int8 = 42
                                v2 = (v2 / (v2 / ((-2.25 + v2) * v2)))
                                switch (fi81 - 1) {
                                case -2147483648:
                                    v2 = 1.5
                                case 42:
                                    var v100 uint32 = 350
                                    _ = v100
                                default:
                                    v101 := v3[("\n\t" + "foo bar")]
                                    _ = v101
                                }
                                var v102 int = -677
                                _ = v102
                                _ = v99
                            }
                            if false {
                                v66 = 42
                            } else {
                                var v103 float64 = float64(float32(3.14159))
                                v2 = v2
                                v103 += 0.1
                                _ = v103
                            }
                        } else {
                            v82 = append(v82, v66)
                            v66 = (fi81 * -373)
                        }
                        _ = fi81
                        _ = v82
                    }
                    switch v66 {
                    case -1:
                        var v104 int = 32768
                        _ = v104
                    case -9223372036854775808:
                        v105 := []int{(-(fi80)), v66, (v66 << ((65535) & 63)), (v66 / (((fi80 / ((1) | 1))) | 1))}
                        _ = v105
                    default:
                        var v106 float32 = v2
                        _ = v106
                    }
                    v107 := GenS6{F0: 42, F1: 1}
                    _ = fi80
                    _ = v107
                }
            }
            v1 = v1
            for fi108 := 0; fi108 < 19; fi108++ {
                v66 = v66
                v109 := GenS1{F0: 0}
                v110 := v3[""]
                _ = fi108
                _ = v109
                _ = v110
            }
        }
    default:
        for fi111 := 0; fi111 < 5; fi111++ {
            v112 := GenS10{F0: (v1 - (^(v1))), F1: 32767}
            if (42 == 100) {
                if false {
                    var v113 uint64 = 229
                    var v114 int16 = 1
                    v115 := GenS11{F0: v113, F1: 648, F2: -1}
                    _ = v113
                    _ = v114
                    _ = v115
                }
                if false {
                    if true {
                        v116 := GenS10{F0: v1, F1: 559}
                        v117 := GenS8{F0: (v1 >> ((1) & 63)), F1: v1}
                        if (1e10 > float64((v2 * ((0.1 - 3.14159) * v2)))) {
                            v118 := GenS1{F0: 42}
                            _ = v118
                        } else {
                            v2 = (0.0 / ((v2 - (v2 - v2)) + 3.14159))
                            fmt.Println(v2)
                        }
                        _ = v116
                        _ = v117
                    } else {
                        if false {
                            v2 = v2
                            v119 := v3["\n\t"]
                            _ = v119
                        } else {
                            if (290 != 1) {
                                _ = len(v3)
                                for fi120 := 0; fi120 < 1; fi120++ {
                                    v1 += v1
                                    v121 := GenS7{F0: fi120, F1: -32768}
                                    v122 := GenS12{F0: (^(fi120)), F1: genFn0(), F2: -697}
                                    _ = fi120
                                    _ = v121
                                    _ = v122
                                }
                            } else {
                                v123 := GenS13{F0: 852, F1: "hello"}
                                v66 = 1
                                _ = v123
                            }
                            var v124 int16 = -536
                            _ = v124
                        }
                        switch "hello" {
                        case "":
                            var v125 uint32 = 2147483647
                            _ = v125
                        case "\n\t":
                            for fi126 := 0; fi126 < 17; fi126++ {
                                v127 := GenS1{F0: 2}
                                _ = fi126
                                _ = v127
                            }
                        case "a":
                            var v128 uint32 = 2147483647
                            _ = v128
                        default:
                            v1 -= 42
                        }
                        var v129 int16 = 0
                        _ = v129
                    }
                    fmt.Println(v66)
                    var v130 float64 = -2.25
                    _ = v130
                }
            } else {
                switch fi111 {
                case 0:
                    var v131 int64 = int64(v1)
                    _ = v131
                case 578:
                    if false {
                        if true {
                            v132 := GenS5{F0: 1}
                            for fi133 := 0; fi133 < 14; fi133++ {
                                var v134 uint64 = 232
                                switch fi111 {
                                case 0:
                                    v2 -= v2
                                case 939:
                                    v135 := v3[""]
                                    _ = v135
                                case -6:
                                    var v136 int32 = -2147483648
                                    _ = v136
                                }
                                fmt.Println(v66, v1, v66)
                                switch "a" {
                                case "foo bar":
                                    if true {
                                        var v137 int64 = 0
                                        if true {
                                            if (true && ((18446744073709551615 >= 0) || false)) {
                                                switch -2147483648 {
                                                case 526:
                                                    fmt.Println(v134, fi111)
                                                case -1:
                                                    var v138 uint = 4294967295
                                                    _ = v138
                                                default:
                                                    v139 := GenS4{F0: fi111}
                                                    _ = v139
                                                }
                                                v140 := GenS6{F0: -338, F1: (v137 - (v137 - 42))}
                                                _ = v140
                                            } else {
                                                v141 := GenS10{F0: 712, F1: -943}
                                                v142 := GenS14{F0: (v1 | v1), F1: 2, F2: 1000}
                                                v1 = v1
                                                _ = v141
                                                _ = v142
                                            }
                                            v143 := GenS7{F0: (v66 % (((^(v66))) | 1)), F1: 946}
                                            fmt.Println(v134, fi133, v2)
                                            _ = v143
                                        }
                                        v66 = (fi111 >> ((42) & 63))
                                        switch (fi111 + fi111) {
                                        case 42:
                                            v144 := v3[("a" + ("foo bar" + ("foo bar" + (("" + "\n\t") + ("foo bar" + "hello")))))]
                                            _ = v144
                                        case 32768:
                                            v66 += (v66 % (((^(v66))) | 1))
                                        default:
                                            v145 := GenS3{F0: v137, F1: 1}
                                            var v146 uint32 = genFn1(((float64((v2 - v2)) / 1.0) < ((0.0 / 1.5) * float64(-2.25))))
                                            v134 += 628
                                            _ = v145
                                            _ = v146
                                        }
                                        if false {
                                            v147 := v3[""]
                                            v134 = v134
                                            v148 := []int{1, 2147483647, v66}
                                            if ((2 != uint(v134)) || true) {
                                                switch "" {
                                                case "hello":
                                                    v149 := GenS15{F0: uint16(v1), F1: -1}
                                                    v148 = append(v148, fi111)
                                                    _ = v149
                                                case "a":
                                                    switch v147 {
                                                    case "foo bar":
                                                        v147 = ""
                                                    }
                                                case "foo bar":
                                                    var v150 float32 = v2
                                                    _ = v150
                                                default:
                                                    fmt.Println(v1, v2)
                                                }
                                                _ = len(v3)
                                                var v151 int = v66
                                                _ = v151
                                            }
                                            var v152 uint64 = 1
                                            if (42 <= -1) {
                                                var v153 string = v147
                                                v2 += v2
                                                var v154 int8 = -128
                                                v2 = v2
                                                var v155 float32 = v2
                                                _ = v153
                                                _ = v154
                                                _ = v155
                                            } else {
                                                v156 := GenS16{F0: 0}
                                                if (!(127 <= 0)) {
                                                    if (176 > -2147483648) {
                                                        v148 = append(v148, -443)
                                                    }
                                                    v157 := GenS6{F0: 2147483647, F1: (v137 % (((v137 >> ((1000) & 63))) | 1))}
                                                    var v158 int32 = 1
                                                    _ = v157
                                                    _ = v158
                                                } else {
                                                    fmt.Println(v2, fi111)
                                                    v159 := GenS13{F0: 475, F1: v147}
                                                    _ = v159
                                                }
                                                fmt.Println(fi133)
                                                _ = v156
                                            }
                                            _ = v147
                                            _ = v148
                                            _ = v152
                                        } else {
                                            var v160 uint8 = 255
                                            v161 := GenS9{F0: 1}
                                            var v162 int16 = 32767
                                            _ = v160
                                            _ = v161
                                            _ = v162
                                        }
                                        _ = v137
                                    }
                                case "a":
                                    switch 0 {
                                    case 9223372036854775807:
                                        var v163 int = 0
                                        _ = v163
                                    case -1:
                                        v164 := GenS17{F0: uint8(v1)}
                                        _ = v164
                                    }
                                case "hello":
                                    var v165 string = "\n\t"
                                    _ = v165
                                default:
                                    v166 := GenS14{F0: v1, F1: genFn0(), F2: 1000}
                                    _ = v166
                                }
                                _ = fi133
                                _ = v134
                            }
                            _ = v132
                        } else {
                            fmt.Println(v66)
                            _ = len(v3)
                            for fi167 := 0; fi167 < 17; fi167++ {
                                v168 := GenS18{F0: "", F1: false, F2: 2}
                                _ = fi167
                                _ = v168
                            }
                        }
                    } else {
                        v169 := v3[""]
                        var v170 int = (v66 % (((v66 & -80)) | 1))
                        for fi171 := 0; fi171 < 9; fi171++ {
                            if (3.14159 > (0.0 * float32(-2.25))) {
                                var v172 float64 = 0.0
                                if (!false) {
                                    var v173 int = fi111
                                    _ = v173
                                } else {
                                    var v174 int8 = -88
                                    v172 += (0.1 - 0.1)
                                    _ = v174
                                }
                                _ = v172
                            } else {
                                switch "foo bar" {
                                case "hello":
                                    var v175 uint32 = 2147483647
                                    _ = v175
                                case "foo bar":
                                    v176 := GenS19{F0: 1}
                                    _ = v176
                                case "":
                                    v66 = v66
                                }
                                v177 := v3[(v169 + v169)]
                                _ = v177
                            }
                            _ = fi171
                        }
                        _ = v169
                        _ = v170
                    }
                default:
                    v1 = 42
                }
                v178 := []int{v66, (fi111 / ((v66) | 1)), fi111, 0}
                _ = v178
            }
            if (!true) {
                var v179 string = ""
                for fi180 := 0; fi180 < 13; fi180++ {
                    v181 := GenS9{F0: 1}
                    _ = fi180
                    _ = v181
                }
                _ = v179
            } else {
                var v182 int64 = -9223372036854775808
                for fi183 := 0; fi183 < 15; fi183++ {
                    v182 = v182
                    delete(v3, "a")
                    v184 := GenS13{F0: 0, F1: ""}
                    if true {
                        fmt.Println(v66)
                        v185 := GenS11{F0: 18446744073709551615, F1: 699, F2: 498}
                        _ = v185
                    }
                    _ = fi183
                    _ = v184
                }
                _ = v182
            }
            _ = fi111
            _ = v112
        }
    }
    _ = len(v3)
    for fi186 := 0; fi186 < 11; fi186++ {
        if false {
            v187 := GenS12{F0: 517, F1: 127, F2: 42}
            _ = v187
        }
        v3["hello"] = "a"
        _ = fi186
    }
    switch "hello" {
    case "hello":
        v188 := GenS3{F0: 584, F1: 0}
        _ = v188
    }
    var v189 bool = (false && (v1 >= 1))
    fmt.Println("v1:", v1)
    fmt.Println("v189:", v189)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v66:", v66)
    _ = v1
    _ = v189
    _ = v2
    _ = v27
    _ = v3
    _ = v62
    _ = v66
}
