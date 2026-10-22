package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16; F1 int64; F2 int }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 string; F1 int32; F2 int32 }
type GenS7 struct { F0 int; F1 uint16 }
type GenS8 struct { F0 int64; F1 int8 }
type GenS9 struct { F0 uint; F1 int32; F2 string }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int32; F1 uint16 }
type GenS12 struct { F0 uint32; F1 uint64 }
type GenS13 struct { F0 int16; F1 uint; F2 uint32 }
type GenS14 struct { F0 int }
type GenS15 struct { F0 string; F1 uint }
func genFn0() int16 {
    var v4 int16 = -1
    var v5 string = ""
    var v6 int8 = 127
    if (!false) {
        if (42 != 1) {
            for fi7 := 0; fi7 < 14; fi7++ {
                if ((-(v4)) == 1000) {
                    v8 := []int{fi7, fi7, fi7, fi7}
                    v6 = -35
                    _ = v8
                }
                v9 := []int{-886, 0}
                var v10 uint8 = 128
                _ = fi7
                _ = v10
                _ = v9
            }
            v5 = v5
        } else {
            if false {
                for fi11 := 0; fi11 < 5; fi11++ {
                    for fi12 := 0; fi12 < 9; fi12++ {
                        v13 := []int{}
                        _ = fi12
                        _ = v13
                    }
                    _ = fi11
                }
            } else {
                if true {
                    v6 = (v6 % (((v6 & v6)) | 1))
                    if ((0.1 + 3.14159) <= (3.14159 / 0.1)) {
                        for fi14 := 0; fi14 < 3; fi14++ {
                            switch fi14 {
                            case -9223372036854775808:
                                v5 = ((v5 + "") + "\n\t")
                            case 128:
                                v5 = v5
                            case 32768:
                                v15 := GenS1{F0: -2147483648, F1: 733}
                                _ = v15
                            default:
                                fmt.Println(v6, v5, fi14)
                            }
                            for fi16 := 0; fi16 < 7; fi16++ {
                                switch (fi16 & fi14) {
                                case -63:
                                    switch "a" {
                                    case "foo bar":
                                        if (true && true) {
                                            v17 := GenS2{F0: 4294967295}
                                            _ = v17
                                        }
                                    }
                                case 925:
                                    v5 = (v5 + v5)
                                }
                                fmt.Println(fi14, fi16)
                                var v18 uint8 = 1
                                switch (v5 + "hello") {
                                case "\n\t":
                                    var v19 int16 = (-(v4))
                                    _ = v19
                                case "hello":
                                    v20 := []int{fi14, fi16, (-(fi16))}
                                    v20 = append(v20, (-(fi14)))
                                    _ = v20
                                default:
                                    var v21 uint64 = 853
                                    _ = v21
                                }
                                v22 := []int{42}
                                _ = fi16
                                _ = v18
                                _ = v22
                            }
                            var v23 int64 = 9223372036854775807
                            _ = fi14
                            _ = v23
                        }
                        v4 = v4
                        switch (v5 + v5) {
                        case "foo bar":
                            v24 := []int{32767, -784, 323, 128}
                            _ = v24
                        case "\n\t":
                            var v25 uint8 = 1
                            _ = v25
                        }
                    }
                    fmt.Println(v6)
                }
            }
        }
    }
    v6 = (-(v6))
    _ = v4
    _ = v5
    _ = v6
    return v4
}
func genFn1() uint {
    var v26 int8 = -1
    if (false && (true || true)) {
        if (!true) {
            var v27 int = -651
            if true {
                if false {
                    switch ((("foo bar" + "a") + "") + "\n\t") {
                    case "a":
                        switch "foo bar" {
                        case "foo bar":
                            v26 += v26
                        default:
                            for fi28 := 0; fi28 < 12; fi28++ {
                                for fi29 := 0; fi29 < 2; fi29++ {
                                    switch "hello" {
                                    case "foo bar":
                                        var v30 uint = 4294967295
                                        _ = v30
                                    }
                                    if (1 < 1) {
                                        v31 := GenS3{F0: 1, F1: 0, F2: (fi29 << ((986) & 63))}
                                        var v32 float64 = (0.0 * float64((((0.0 * ((((1e10 + ((0.0 / ((1.5 + (float32((-2.25 + 0.0)) + 0.1)) * 3.14159)) + (-2.25 / (3.14159 + (0.1 + 0.0))))) + 1.5) + 0.0) / ((((0.0 - ((float32(0.0) + (0.1 * 1e10)) * 3.14159)) - 0.0) - -2.25) / 1.5))) * float32((0.0 + -2.25))) - 3.14159)))
                                        var v33 uint = 65535
                                        var v34 uint32 = 1
                                        _ = v31
                                        _ = v32
                                        _ = v33
                                        _ = v34
                                    } else {
                                        v35 := []int{fi28, 4294967295, (^(fi28)), -2147483648}
                                        fmt.Println(v26, fi28, v27)
                                        switch ("" + "\n\t") {
                                        case "\n\t":
                                            if true {
                                                v27 = v27
                                            }
                                        default:
                                            if false {
                                                if len(v35) >= 2 {
                                                    v35 = v35[1:len(v35)]
                                                }
                                                v35 = append(v35, (fi28 * fi28))
                                                switch "hello" {
                                                case "":
                                                    fmt.Println(fi28)
                                                case "foo bar":
                                                    if (-2147483648 > -786) {
                                                        v26 = 0
                                                    }
                                                case "\n\t":
                                                    if len(v35) >= 2 {
                                                        v35 = v35[1:len(v35)]
                                                    }
                                                default:
                                                    if false {
                                                        if len(v35) > 0 {
                                                            v35[0] = fi29
                                                        }
                                                        if len(v35) > 0 {
                                                            v35[0] = -9223372036854775808
                                                        }
                                                    } else {
                                                        var v36 int32 = 2147483647
                                                        fmt.Println(fi28)
                                                        v37 := GenS3{F0: 32767, F1: -9223372036854775808, F2: (^(fi29))}
                                                        _ = v36
                                                        _ = v37
                                                    }
                                                }
                                                var v38 int64 = 0
                                                _ = v38
                                            }
                                        }
                                        _ = v35
                                    }
                                    _ = fi29
                                }
                                v39 := GenS1{F0: 0, F1: 9223372036854775807}
                                _ = fi28
                                _ = v39
                            }
                        }
                    default:
                        v40 := GenS2{F0: 0}
                        _ = v40
                    }
                    v41 := GenS4{F0: 0}
                    _ = v41
                }
            } else {
                v42 := []int{}
                var v43 int8 = v26
                _ = v42
                _ = v43
            }
            _ = v27
        } else {
            v44 := []int{128}
            _ = v44
        }
        var v45 uint32 = 2147483648
        fmt.Println(v26, v26)
        _ = v45
    } else {
        var v46 bool = true
        _ = v46
    }
    if true {
        for fi47 := 0; fi47 < 10; fi47++ {
            if (true && (1000 == 0)) {
                v26 += v26
                v48 := GenS5{F0: 4294967295}
                _ = v48
            } else {
                var v49 uint8 = 42
                var v50 uint64 = 268
                v26 = -1
                _ = v49
                _ = v50
            }
            if false {
                v51 := []int{(fi47 - (fi47 & (fi47 ^ (fi47 / ((fi47) | 1))))), fi47, (fi47 % (((fi47 ^ (fi47 - fi47))) | 1)), (-(fi47))}
                fmt.Println(fi47)
                v52 := GenS6{F0: (("a" + "hello") + ""), F1: 1, F2: 42}
                _ = v51
                _ = v52
            } else {
                switch (fi47 + (fi47 + 4294967295)) {
                case 32767:
                    var v53 float64 = float64(float32(-2.25))
                    _ = v53
                case 1:
                    v54 := GenS7{F0: (fi47 & fi47), F1: 65535}
                    _ = v54
                }
            }
            if (!(10 >= 65535)) {
                for fi55 := 0; fi55 < 18; fi55++ {
                    v56 := []int{}
                    fmt.Println(fi47, v26)
                    _ = fi55
                    _ = v56
                }
                switch (fi47 ^ (fi47 / ((4294967295) | 1))) {
                case 9223372036854775807:
                    for fi57 := 0; fi57 < 12; fi57++ {
                        if true {
                            v26 = (v26 * -1)
                            v26 = (-(v26))
                            var v58 string = "foo bar"
                            _ = v58
                        }
                        var v59 int32 = -43
                        _ = fi57
                        _ = v59
                    }
                case -493:
                    var v60 bool = (fi47 < fi47)
                    _ = v60
                case 1:
                    v26 = v26
                }
            } else {
                for fi61 := 0; fi61 < 6; fi61++ {
                    v26 = 1
                    _ = fi61
                }
                if (!false) {
                    fmt.Println(v26, v26)
                    fmt.Println(v26)
                }
                if false {
                    v26 = v26
                } else {
                    v62 := []int{-474}
                    if false {
                        v63 := GenS6{F0: (("hello" + "foo bar") + "foo bar"), F1: -845, F2: 1}
                        if (2 >= 0) {
                            if (("foo bar" + "\n\t") > "hello") {
                                for fi64 := 0; fi64 < 17; fi64++ {
                                    switch "" {
                                    case "a":
                                        v62 = append(v62, fi47)
                                    case "\n\t":
                                        var v65 float32 = ((-2.25 / 1e10) - 3.14159)
                                        _ = v65
                                    case "hello":
                                        v26 += v26
                                    }
                                    _ = fi64
                                }
                                if ((v26 & v26) < (^(v26))) {
                                    for fi66 := 0; fi66 < 4; fi66++ {
                                        v62 = append(v62, (-(fi66)))
                                        _ = fi66
                                    }
                                } else {
                                    if true {
                                        for fi67 := 0; fi67 < 1; fi67++ {
                                            var v68 uint64 = 1
                                            if (false || true) {
                                                v68 = v68
                                            }
                                            if ((1 <= 419) || ((("\n\t" + ("" + ("a" + (("a" + "\n\t") + "hello")))) + ((("hello" + "hello") + "a") + "hello")) >= "hello")) {
                                                _ = len(v62)
                                                fmt.Println(v26)
                                            }
                                            v26 = (v26 & v26)
                                            _ = fi67
                                            _ = v68
                                        }
                                        var v69 uint = 255
                                        if false {
                                            v62 = append(v62, (fi47 / ((0) | 1)))
                                            fmt.Println(fi47, v69, fi47)
                                            v69 = v69
                                        } else {
                                            var v70 uint = (v69 << (((^(v69))) & 63))
                                            _ = v70
                                        }
                                        _ = v69
                                    } else {
                                        for fi71 := 0; fi71 < 7; fi71++ {
                                            v26 += 127
                                            fmt.Println(fi71, v26, fi71)
                                            _ = fi71
                                        }
                                    }
                                }
                                v26 = 100
                            } else {
                                for fi72 := 0; fi72 < 15; fi72++ {
                                    var v73 int64 = -1
                                    fmt.Println(v26)
                                    _ = fi72
                                    _ = v73
                                }
                            }
                        } else {
                            var v74 uint16 = 510
                            fmt.Println(v74, v26, fi47)
                            _ = v74
                        }
                        var v75 int64 = 1
                        _ = v63
                        _ = v75
                    } else {
                        if true {
                            var v76 uint64 = 4294967295
                            v77 := GenS2{F0: (v76 * (v76 * v76))}
                            _ = v76
                            _ = v77
                        } else {
                            v78 := GenS3{F0: 255, F1: 852, F2: fi47}
                            if (1 == 42) {
                                if true {
                                    for fi79 := 0; fi79 < 12; fi79++ {
                                        var v80 uint16 = 255
                                        var v81 float64 = -2.25
                                        _ = fi79
                                        _ = v80
                                        _ = v81
                                    }
                                    switch "foo bar" {
                                    case "hello":
                                        v26 -= v26
                                    case "\n\t":
                                        v26 = 127
                                    }
                                    var v82 int = fi47
                                    _ = v82
                                }
                                for fi83 := 0; fi83 < 19; fi83++ {
                                    v62 = append(v62, fi83)
                                    _ = fi83
                                }
                            } else {
                                if false {
                                    v26 += v26
                                }
                                v62 = append(v62, 4294967295)
                            }
                            _ = v78
                        }
                        if len(v62) > 0 {
                            v62[0] = fi47
                        }
                        var v84 int = fi47
                        _ = v84
                    }
                    var v85 uint32 = 2147483647
                    _ = v62
                    _ = v85
                }
            }
            _ = fi47
        }
        v26 += v26
    } else {
        if (1 <= 42) {
            for fi86 := 0; fi86 < 12; fi86++ {
                if ((v26 * (v26 >> ((42) & 63))) > 1) {
                    for fi87 := 0; fi87 < 4; fi87++ {
                        switch "a" {
                        case "\n\t":
                            v26 = v26
                        case "a":
                            v26 = v26
                        }
                        var v88 uint64 = 4294967295
                        _ = fi87
                        _ = v88
                    }
                    for fi89 := 0; fi89 < 12; fi89++ {
                        v26 -= 0
                        _ = fi89
                    }
                    if ((true || true) || (255 < 255)) {
                        v90 := GenS2{F0: 1}
                        if (v26 >= v26) {
                            for fi91 := 0; fi91 < 3; fi91++ {
                                fmt.Println(fi86, v26, v26)
                                v26 -= v26
                                if (396 < 32767) {
                                    v92 := GenS4{F0: 2}
                                    v93 := []int{}
                                    _ = v92
                                    _ = v93
                                } else {
                                    v94 := []int{fi91, fi86}
                                    var v95 int32 = -213
                                    var v96 uint32 = 704
                                    _ = v94
                                    _ = v95
                                    _ = v96
                                }
                                _ = fi91
                            }
                        } else {
                            var v97 int16 = 1000
                            v98 := []int{}
                            switch fi86 {
                            case 2147483647:
                                if true {
                                    v26 = 1
                                }
                            case 32768:
                                for fi99 := 0; fi99 < 15; fi99++ {
                                    switch fi86 {
                                    case -9223372036854775808:
                                        fmt.Println(fi99, v26, v26)
                                    }
                                    if true {
                                        var v100 float64 = -2.25
                                        _ = v100
                                    }
                                    var v101 int = (fi86 - (^(fi99)))
                                    _ = fi99
                                    _ = v101
                                }
                            case -910:
                                for fi102 := 0; fi102 < 7; fi102++ {
                                    v103 := GenS5{F0: 65535}
                                    _ = fi102
                                    _ = v103
                                }
                            default:
                                if (1 == 0) {
                                    v26 = v26
                                    var v104 uint = 42
                                    _ = v104
                                } else {
                                    var v105 int = -232
                                    _ = v105
                                }
                            }
                            _ = v97
                            _ = v98
                        }
                        _ = v90
                    }
                } else {
                    v106 := GenS2{F0: 42}
                    v26 = (v26 + (v26 + (^(v26))))
                    if ("" < "foo bar") {
                        fmt.Println(v26, v26)
                        v26 = (v26 & v26)
                        if true {
                            v107 := []int{1, 9223372036854775807}
                            _ = v107
                        } else {
                            fmt.Println(fi86, fi86)
                            v26 += 127
                            v108 := GenS8{F0: 1, F1: (v26 % ((127) | 1))}
                            _ = v108
                        }
                    }
                    _ = v106
                }
                if ("foo bar" > "a") {
                    v109 := []int{(fi86 * fi86), (fi86 >> ((65535) & 63)), -528, fi86}
                    _ = v109
                } else {
                    v110 := GenS9{F0: 2, F1: 426, F2: "\n\t"}
                    var v111 int32 = -1
                    _ = v110
                    _ = v111
                }
                _ = fi86
            }
        } else {
            v112 := GenS8{F0: 9223372036854775807, F1: v26}
            _ = v112
        }
        if ((-32768 == genFn0()) || ((2 != 686) || true)) {
            for fi113 := 0; fi113 < 15; fi113++ {
                var v114 int = (fi113 - 42)
                v26 -= v26
                _ = fi113
                _ = v114
            }
            var v115 int32 = -2147483648
            _ = v115
        } else {
            v116 := GenS4{F0: 485}
            _ = v116
        }
        var v117 uint64 = 2
        _ = v117
    }
    if true {
        switch 996 {
        case 2147483647:
            if (!true) {
                v118 := GenS10{F0: 1}
                if (-9223372036854775808 == 600) {
                    v26 += 4
                } else {
                    v26 += v26
                }
                v119 := []int{268, 4294967295, 32768}
                _ = v118
                _ = v119
            }
        case 42:
            for fi120 := 0; fi120 < 14; fi120++ {
                if false {
                    if (!true) {
                        v121 := GenS9{F0: 578, F1: 1, F2: ("\n\t" + ("a" + (("" + "a") + "a")))}
                        v122 := GenS7{F0: fi120, F1: 2}
                        _ = v121
                        _ = v122
                    }
                    v123 := []int{fi120}
                    _ = v123
                }
                if (42 >= -204) {
                    for fi124 := 0; fi124 < 14; fi124++ {
                        v26 += 96
                        v26 += (v26 * (v26 >> ((184) & 63)))
                        _ = fi124
                    }
                    var v125 int8 = v26
                    for fi126 := 0; fi126 < 18; fi126++ {
                        v26 += v26
                        _ = fi126
                    }
                    _ = v125
                } else {
                    v127 := []int{(fi120 - 621), (fi120 + -889), (fi120 ^ (fi120 << ((18446744073709551615) & 63)))}
                    _ = v127
                }
                for fi128 := 0; fi128 < 6; fi128++ {
                    if (fi128 >= fi120) {
                        v129 := []int{}
                        v26 += v26
                        _ = v129
                    }
                    var v130 string = ("foo bar" + "")
                    _ = fi128
                    _ = v130
                }
                _ = fi120
            }
        default:
            switch "foo bar" {
            case "a":
                for fi131 := 0; fi131 < 15; fi131++ {
                    var v132 int = (-(fi131))
                    var v133 uint16 = 1
                    _ = fi131
                    _ = v132
                    _ = v133
                }
            case "":
                for fi134 := 0; fi134 < 20; fi134++ {
                    switch fi134 {
                    case -932:
                        fmt.Println(fi134, fi134)
                    case 2147483647:
                        var v135 int32 = 2147483647
                        _ = v135
                    case 127:
                        if true {
                            var v136 int32 = -2147483648
                            v136 = (v136 ^ 0)
                            for fi137 := 0; fi137 < 19; fi137++ {
                                v26 += -95
                                var v138 int16 = 0
                                fmt.Println(v138, v26, fi134)
                                _ = fi137
                                _ = v138
                            }
                            _ = v136
                        } else {
                            v26 += v26
                            var v139 uint32 = 1
                            var v140 int64 = int64(v139)
                            _ = v139
                            _ = v140
                        }
                    }
                    for fi141 := 0; fi141 < 1; fi141++ {
                        switch (fi141 | -1) {
                        case 9223372036854775807:
                            v142 := []int{}
                            _ = v142
                        default:
                            v143 := GenS5{F0: 1}
                            _ = v143
                        }
                        var v144 int64 = 1
                        _ = fi141
                        _ = v144
                    }
                    for fi145 := 0; fi145 < 14; fi145++ {
                        var v146 uint = 42
                        v147 := GenS11{F0: -640, F1: 42}
                        v148 := []int{}
                        _ = fi145
                        _ = v146
                        _ = v147
                        _ = v148
                    }
                    _ = fi134
                }
            case "foo bar":
                for fi149 := 0; fi149 < 3; fi149++ {
                    var v150 uint8 = 72
                    _ = fi149
                    _ = v150
                }
            }
        }
    }
    _ = v26
    return 18446744073709551615
}

func main() {
    var v1 int64 = 906
    var v2 uint = 0
    v3 := []string{"foo bar", ""}
    _ = len(v3)
    fmt.Println(v1, v2)
    if (2 == 255) {
        var v151 uint64 = 0
        _ = v151
    }
    v152 := GenS7{F0: 127, F1: 65535}
    switch ("a" + "a") {
    case "hello":
        var v153 uint8 = 255
        _ = v153
    case "foo bar":
        for fi154 := 0; fi154 < 16; fi154++ {
            v1 += v1
            _ = fi154
        }
    case "a":
        v3 = append(v3, "a")
    default:
        if true {
            v155 := GenS12{F0: 41, F1: 1}
            switch 1 {
            case 4294967295:
                for fi156 := 0; fi156 < 19; fi156++ {
                    fmt.Println(v1, v1, v1)
                    v1 = (-(v1))
                    v2 = (v2 - v2)
                    _ = fi156
                }
            case 128:
                v2 -= v2
            case -9223372036854775808:
                for fi157 := 0; fi157 < 9; fi157++ {
                    for fi158 := 0; fi158 < 9; fi158++ {
                        v1 += -9223372036854775808
                        var v159 int8 = 0
                        _ = fi158
                        _ = v159
                    }
                    _ = fi157
                }
            }
            _ = len(v3)
            _ = v155
        }
    }
    if true {
        v2 = 904
        for fi160 := 0; fi160 < 14; fi160++ {
            v161 := GenS13{F0: 0, F1: v2, F2: 2147483648}
            fmt.Println(v1, fi160)
            v162 := GenS8{F0: -238, F1: 42}
            _ = fi160
            _ = v161
            _ = v162
        }
    }
    for fi163 := 0; fi163 < 11; fi163++ {
        v164 := GenS14{F0: (fi163 << (((v2 / (((v2 % (((v2 * (v2 << ((genFn1()) & 63)))) | 1))) | 1))) & 63))}
        var v165 string = "a"
        _ = fi163
        _ = v164
        _ = v165
    }
    _ = len(v3)
    fmt.Println(v1, v1)
    var v166 float32 = 0.0
    if (0 == 42) {
        var v167 float32 = ((v166 / ((-2.25 + ((v166 + float32((((0.0 - ((1e10 + 0.1) - -2.25)) + (float64(v166) / 0.1)) * 1e10))) + (v166 * ((float32(3.14159) + v166) * (((((1.5 - (v166 - v166)) + (v166 - 0.0)) * v166) - v166) * v166))))) - (0.1 * (-2.25 + v166)))) / v166)
        switch 623 {
        case 128:
            for fi168 := 0; fi168 < 12; fi168++ {
                v167 = v167
                if true {
                    v169 := GenS14{F0: -600}
                    _ = v169
                }
                _ = fi168
            }
        case 2147483647:
            fmt.Println(v166, v167, v167)
        default:
            var v170 int64 = v1
            _ = v170
        }
        v171 := GenS15{F0: "a", F1: (^(v2))}
        _ = v167
        _ = v171
    }
    fmt.Println("v1:", v1)
    fmt.Println("v166:", v166)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v152
    _ = v166
    _ = v2
    _ = v3
}
