package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 int8 }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint16; F1 int32 }
type GenS6 struct { F0 uint8; F1 int16; F2 bool }
type GenS7 struct { F0 int8; F1 string; F2 int32 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 uint16; F1 bool }
type GenS11 struct { F0 int16; F1 uint; F2 int64 }
type GenS12 struct { F0 uint16; F1 uint32; F2 string }
type GenS13 struct { F0 uint32; F1 int }
type GenS14 struct { F0 int8 }
type GenS15 struct { F0 uint8; F1 uint8 }
type GenS16 struct { F0 int16; F1 int8 }
func genFn0(p0 int16, p1 int) string {
    v5 := []int{p1}
    switch "" {
    case "\n\t":
        for fi6 := 0; fi6 < 6; fi6++ {
            v7 := GenS1{F0: 128, F1: 2, F2: 0}
            for fi8 := 0; fi8 < 2; fi8++ {
                p0 += (p0 << ((4294967295) & 63))
                v5 = append(v5, -423)
                v9 := GenS2{F0: -1, F1: (p1 | -9223372036854775808)}
                _ = fi8
                _ = v9
            }
            _ = fi6
            _ = v7
        }
    case "foo bar":
        var v10 uint32 = 2147483648
        _ = v10
    case "a":
        fmt.Println(p0, p1)
    default:
        var v11 int16 = (p0 - p0)
        _ = v11
    }
    switch ((("hello" + "\n\t") + "foo bar") + "hello") {
    case "a":
        var v12 int64 = 42
        _ = v12
    case "foo bar":
        for fi13 := 0; fi13 < 18; fi13++ {
            if true {
                if true {
                    var v14 uint8 = 42
                    for fi15 := 0; fi15 < 18; fi15++ {
                        v14 += v14
                        _ = fi15
                    }
                    _ = v14
                } else {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    switch -1 {
                    case 4294967295:
                        if (!((p1 ^ fi13) == 128)) {
                            p0 = p0
                            fmt.Println(fi13)
                            if len(v5) > 0 {
                                v5[0] = 32767
                            }
                        } else {
                            switch "foo bar" {
                            case "foo bar":
                                if true {
                                    if true {
                                        p0 += 32767
                                        p1 = (^(p1))
                                    }
                                    if (205 == 171) {
                                        if false {
                                            for fi16 := 0; fi16 < 20; fi16++ {
                                                var v17 uint16 = 65535
                                                _ = fi16
                                                _ = v17
                                            }
                                            var v18 uint64 = 2
                                            _ = v18
                                        } else {
                                            if len(v5) > 0 {
                                                v5[0] = (p1 >> ((42) & 63))
                                            }
                                        }
                                        if false {
                                            if ("foo bar" < "hello") {
                                                switch "" {
                                                case "\n\t":
                                                    if (4294967295 > 2) {
                                                        for fi19 := 0; fi19 < 14; fi19++ {
                                                            fmt.Println(fi19)
                                                            switch 32767 {
                                                            case 4294967295:
                                                                var v20 uint = 2
                                                                _ = v20
                                                            case 32768:
                                                                p1 -= 128
                                                            default:
                                                                p1 = -28
                                                            }
                                                            p1 = fi13
                                                            _ = fi19
                                                        }
                                                        v21 := GenS1{F0: 0, F1: 2, F2: 1}
                                                        _ = v21
                                                    }
                                                }
                                                for fi22 := 0; fi22 < 5; fi22++ {
                                                    fmt.Println(fi13, fi13, p0)
                                                    var v23 uint = 255
                                                    _ = fi22
                                                    _ = v23
                                                }
                                            } else {
                                                switch "" {
                                                case "":
                                                    if (2 < 128) {
                                                        if (false || (97 >= 192)) {
                                                            for fi24 := 0; fi24 < 4; fi24++ {
                                                                p0 = p0
                                                                v5 = append(v5, p1)
                                                                if ((false && true) && (551 == 1)) {
                                                                    p1 += 127
                                                                    var v25 int8 = -128
                                                                    v5 = append(v5, 9223372036854775807)
                                                                    _ = v25
                                                                }
                                                                p0 += (p0 | p0)
                                                                var v26 uint = 259
                                                                _ = fi24
                                                                _ = v26
                                                            }
                                                        }
                                                        v5 = append(v5, 128)
                                                        for fi27 := 0; fi27 < 14; fi27++ {
                                                            var v28 int32 = 1
                                                            _ = fi27
                                                            _ = v28
                                                        }
                                                    } else {
                                                        p1 = (p1 + (^(p1)))
                                                        var v29 int64 = 9223372036854775807
                                                        _ = v29
                                                    }
                                                case "a":
                                                    if false {
                                                        fmt.Println(p0)
                                                    } else {
                                                        p0 = p0
                                                        v30 := GenS3{F0: 2147483648}
                                                        _ = v30
                                                    }
                                                case "\n\t":
                                                    var v31 bool = (2 <= 42)
                                                    _ = v31
                                                }
                                            }
                                            var v32 int32 = 1
                                            v33 := GenS4{F0: 4294967295}
                                            _ = v32
                                            _ = v33
                                        }
                                    }
                                } else {
                                    for fi34 := 0; fi34 < 8; fi34++ {
                                        p0 -= p0
                                        _ = fi34
                                    }
                                    var v35 uint = 255
                                    _ = v35
                                }
                            case "a":
                                fmt.Println(p1, p0)
                            case "":
                                if ("a" < (((((("a" + ("a" + "a")) + "") + "foo bar") + "hello") + "\n\t") + ("" + ""))) {
                                    var v36 uint = 1
                                    switch "" {
                                    case "":
                                        if (-790 > (p0 - p0)) {
                                            if true {
                                                for fi37 := 0; fi37 < 7; fi37++ {
                                                    p0 = p0
                                                    if true {
                                                        if ((float32(-2.25) < 3.14159) && (0.0 == 0.1)) {
                                                            var v38 int = (p1 * 32767)
                                                            fmt.Println(v38)
                                                            _ = v38
                                                        }
                                                        var v39 int = fi37
                                                        p0 += 720
                                                        _ = v39
                                                    }
                                                    _ = len(v5)
                                                    _ = fi37
                                                }
                                                p0 = 1
                                            }
                                        } else {
                                            fmt.Println(p0)
                                            for fi40 := 0; fi40 < 11; fi40++ {
                                                _ = len(v5)
                                                v41 := GenS4{F0: fi13}
                                                var v42 string = (("foo bar" + "\n\t") + "")
                                                _ = fi40
                                                _ = v41
                                                _ = v42
                                            }
                                        }
                                    case "\n\t":
                                        var v43 int32 = 2147483647
                                        _ = v43
                                    default:
                                        switch ((("" + "\n\t") + "foo bar") + "foo bar") {
                                        case "hello":
                                            if true {
                                                for fi44 := 0; fi44 < 17; fi44++ {
                                                    v45 := GenS3{F0: 644}
                                                    if false {
                                                        v36 = (v36 | v36)
                                                        if ((!(!(18446744073709551615 < 0))) && (("foo bar" + "") >= ((("" + "") + "a") + ""))) {
                                                            v46 := GenS5{F0: 65535, F1: 87}
                                                            if len(v5) >= 2 {
                                                                v5 = v5[1:len(v5)]
                                                            }
                                                            var v47 int32 = 0
                                                            _ = v46
                                                            _ = v47
                                                        }
                                                    } else {
                                                        p1 = fi13
                                                        var v48 uint8 = 225
                                                        var v49 uint16 = 107
                                                        _ = v48
                                                        _ = v49
                                                    }
                                                    v50 := GenS6{F0: 128, F1: 42, F2: (255 != 2)}
                                                    fmt.Println(v36, p0, p0)
                                                    v51 := GenS6{F0: 0, F1: p0, F2: true}
                                                    _ = fi44
                                                    _ = v45
                                                    _ = v50
                                                    _ = v51
                                                }
                                                if true {
                                                    for fi52 := 0; fi52 < 1; fi52++ {
                                                        switch "a" {
                                                        case "foo bar":
                                                            var v53 int16 = (p0 - (p0 ^ p0))
                                                            _ = v53
                                                        case "":
                                                            var v54 uint = v36
                                                            _ = v54
                                                        case "a":
                                                            if (!((true || (!false)) && (103 < 2147483647))) {
                                                                v5 = append(v5, (p1 - p1))
                                                                v5 = append(v5, 32767)
                                                            } else {
                                                                var v55 uint8 = 0
                                                                var v56 int32 = -756
                                                                _ = v55
                                                                _ = v56
                                                            }
                                                        }
                                                        var v57 int64 = 9223372036854775807
                                                        _ = fi52
                                                        _ = v57
                                                    }
                                                    fmt.Println(fi13)
                                                    var v58 int32 = 2147483647
                                                    _ = v58
                                                }
                                                if ((true || false) || true) {
                                                    var v59 int8 = -27
                                                    _ = v59
                                                }
                                            }
                                        case "\n\t":
                                            for fi60 := 0; fi60 < 3; fi60++ {
                                                fmt.Println(fi13, p0)
                                                _ = fi60
                                            }
                                        case "a":
                                            fmt.Println(v36, p1, p1)
                                        }
                                    }
                                    _ = v36
                                }
                            default:
                                if (!(!(((true || (-705 <= 9223372036854775807)) && (false && false)) && false))) {
                                    var v61 bool = (!true)
                                    _ = v61
                                } else {
                                    switch fi13 {
                                    case 0:
                                        p1 = fi13
                                    default:
                                        var v62 int16 = (p0 + 1)
                                        _ = v62
                                    }
                                    var v63 uint = 255
                                    if (true && false) {
                                        var v64 uint = (v63 >> ((v63) & 63))
                                        v65 := GenS6{F0: 255, F1: (p0 >> ((v64) & 63)), F2: false}
                                        _ = v64
                                        _ = v65
                                    }
                                    _ = v63
                                }
                            }
                            for fi66 := 0; fi66 < 14; fi66++ {
                                _ = len(v5)
                                _ = fi66
                            }
                            p0 = (p0 / ((p0) | 1))
                        }
                    case 1:
                        _ = len(v5)
                    case -9223372036854775808:
                        var v67 uint64 = 42
                        _ = v67
                    }
                }
                var v68 uint = 65535
                v69 := GenS4{F0: (p1 * fi13)}
                _ = v68
                _ = v69
            }
            p0 -= p0
            p1 -= (^(fi13))
            _ = fi13
        }
    default:
        if ((p0 - -32768) <= (-(p0))) {
            p1 = -1
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v70 := GenS4{F0: p1}
            _ = v70
        }
    }
    if true {
        for fi71 := 0; fi71 < 10; fi71++ {
            var v72 uint = 243
            if (1 == 2147483647) {
                var v73 int64 = -9223372036854775808
                v72 -= (v72 >> ((v72) & 63))
                _ = v73
            }
            _ = fi71
            _ = v72
        }
    }
    p0 = (p0 >> ((2) & 63))
    _ = p0
    _ = p1
    _ = v5
    return ""
}
func genFn1(p0 uint64, p1 uint16, p2 int16) uint32 {
    p2 = p2
    var v74 int32 = -1
    v75 := []int{42, 9223372036854775807, -2147483648}
    fmt.Println(p1, p0, v74)
    if (9223372036854775807 == -9223372036854775808) {
        for fi76 := 0; fi76 < 6; fi76++ {
            for fi77 := 0; fi77 < 8; fi77++ {
                var v78 uint = 2
                switch fi77 {
                case 190:
                    var v79 int = 128
                    _ = v79
                case 551:
                    var v80 int64 = 9223372036854775807
                    v74 = (v74 - (v74 & (v74 / ((v74) | 1))))
                    switch (("hello" + ("\n\t" + "a")) + ((("a" + "\n\t") + ("" + "\n\t")) + (("foo bar" + "") + "a"))) {
                    case "":
                        switch "a" {
                        case "":
                            switch fi77 {
                            case 2147483647:
                                v81 := GenS1{F0: 0, F1: v78, F2: -1}
                                _ = v81
                            case -1:
                                if true {
                                    v75 = append(v75, 1)
                                    switch (fi77 % (((-(fi77))) | 1)) {
                                    case 9223372036854775807:
                                        v80 += (v80 % ((-900) | 1))
                                    default:
                                        v82 := GenS2{F0: 127, F1: (fi77 << ((18446744073709551615) & 63))}
                                        _ = v82
                                    }
                                    switch 2147483647 {
                                    case 373:
                                        var v83 float32 = 3.14159
                                        _ = v83
                                    default:
                                        v75 = append(v75, fi76)
                                    }
                                } else {
                                    var v84 uint32 = uint32(p0)
                                    if len(v75) > 0 {
                                        v75[0] = fi77
                                    }
                                    _ = v84
                                }
                            case -509:
                                v85 := GenS7{F0: 127, F1: "\n\t", F2: int32(v78)}
                                _ = v85
                            default:
                                v75 = append(v75, -296)
                            }
                        case "a":
                            v80 += v80
                        case "foo bar":
                            if (((p2 << (((^(v78))) & 63)) > (p2 * (p2 % (((p2 << (((^(v78))) & 63))) | 1)))) && true) {
                                switch 9223372036854775807 {
                                case 32768:
                                    var v86 int64 = v80
                                    _ = v86
                                case 4294967295:
                                    if len(v75) >= 2 {
                                        v75 = v75[1:len(v75)]
                                    }
                                    fmt.Println(v80)
                                    var v87 float64 = float64(-2.25)
                                    _ = v87
                                }
                                fmt.Println(fi76, fi76, v74)
                                var v88 uint16 = 172
                                v80 -= v80
                                _ = v88
                            } else {
                                v89 := GenS6{F0: 0, F1: (p2 + (p2 % ((-32768) | 1))), F2: false}
                                _ = v89
                            }
                        }
                    default:
                        fmt.Println(p2, p0)
                    }
                    v90 := GenS8{F0: "\n\t"}
                    var v91 float32 = ((1.5 - (1.5 * (float32(0.1) / 0.1))) / 1.0)
                    _ = v80
                    _ = v90
                    _ = v91
                default:
                    var v92 float64 = 0.0
                    _ = v92
                }
                _ = fi77
                _ = v78
            }
            if true {
                p2 = 42
                fmt.Println(fi76)
            }
            _ = fi76
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v74
    _ = v75
    return 1
}

func main() {
    v1 := make(map[string]string)
    var v2 string = ""
    var v3 int32 = 2147483647
    v4 := []string{v2, "foo bar", v2, v2}
    if false {
        var v93 bool = (122 >= 97)
        var v94 uint16 = 42
        _ = v93
        _ = v94
    } else {
        var v95 int32 = (v3 & (v3 + (^(v3))))
        _ = v95
    }
    switch 20 {
    case -561:
        for fi96 := 0; fi96 < 16; fi96++ {
            if (0.1 == -2.25) {
                if true {
                    var v97 float32 = 1e10
                    for fi98 := 0; fi98 < 3; fi98++ {
                        v99 := GenS1{F0: 42, F1: 580, F2: -95}
                        _ = fi98
                        _ = v99
                    }
                    _ = v97
                } else {
                    var v100 uint = 525
                    v100 = v100
                    _ = v100
                }
            } else {
                if (v3 > v3) {
                    var v101 uint64 = 42
                    v2 = ""
                    _ = v101
                }
            }
            v3 -= v3
            _ = fi96
        }
    case -1:
        for fi102 := 0; fi102 < 9; fi102++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if true {
                for fi103 := 0; fi103 < 7; fi103++ {
                    v2 = (v2 + "a")
                    v104 := GenS7{F0: -67, F1: v2, F2: 0}
                    var v105 uint64 = 664
                    _ = fi103
                    _ = v104
                    _ = v105
                }
            }
            var v106 int = fi102
            _ = fi102
            _ = v106
        }
    case 144:
        for fi107 := 0; fi107 < 4; fi107++ {
            var v108 uint = 2
            if (v2 <= v2) {
                v2 = (v2 + "\n\t")
            } else {
                for fi109 := 0; fi109 < 7; fi109++ {
                    if ((!true) && (-321 == 32768)) {
                        fmt.Println(v2, v2, fi109)
                        var v110 uint32 = 4294967295
                        _ = v110
                    } else {
                        if ((3.14159 + -2.25) > 0.0) {
                            var v111 float64 = 1e10
                            v112 := GenS4{F0: (-(fi107))}
                            _ = len(v4)
                            _ = v111
                            _ = v112
                        } else {
                            if true {
                                var v113 uint16 = 32767
                                v2 = "hello"
                                v2 = v2
                                v3 = (v3 & (v3 | -1))
                                v114 := GenS8{F0: v2}
                                _ = v113
                                _ = v114
                            }
                            switch (fi107 | fi109) {
                            case -82:
                                _ = len(v1)
                            case 2147483647:
                                switch 4294967295 {
                                case -9223372036854775808:
                                    var v115 int16 = -951
                                    _ = v115
                                case 784:
                                    fmt.Println(fi109, v2, v3)
                                case -773:
                                    v2 = ""
                                    var v116 bool = false
                                    v3 += (v3 | (v3 % ((-2147483648) | 1)))
                                    _ = v116
                                default:
                                    var v117 bool = true
                                    _ = v117
                                }
                            }
                        }
                    }
                    v108 = v108
                    _ = fi109
                }
                for fi118 := 0; fi118 < 17; fi118++ {
                    v3 -= (v3 >> ((v108) & 63))
                    var v119 string = (v2 + (v2 + ""))
                    v120 := GenS8{F0: v119}
                    _ = fi118
                    _ = v119
                    _ = v120
                }
            }
            _ = fi107
            _ = v108
        }
    default:
        var v121 uint32 = 42
        _ = v121
    }
    if false {
        _ = len(v1)
        if (-2.25 >= 1.5) {
            for fi122 := 0; fi122 < 5; fi122++ {
                for fi123 := 0; fi123 < 15; fi123++ {
                    var v124 int32 = v3
                    v125 := v1[v2]
                    var v126 float32 = 1e10
                    _ = fi123
                    _ = v124
                    _ = v125
                    _ = v126
                }
                fmt.Println(v3, v2)
                _ = fi122
            }
        } else {
            if true {
                v127 := GenS9{F0: v3}
                v128 := GenS10{F0: 255, F1: (1e10 != (float32(-2.25) * 0.1))}
                var v129 uint = 1
                _ = v127
                _ = v128
                _ = v129
            }
            for fi130 := 0; fi130 < 18; fi130++ {
                v131 := GenS9{F0: -617}
                _ = fi130
                _ = v131
            }
            var v132 uint16 = 996
            _ = v132
        }
        delete(v1, v2)
    }
    if (-9223372036854775808 == 42) {
        for fi133 := 0; fi133 < 5; fi133++ {
            fmt.Println(v3)
            if true {
                var v134 int8 = -1
                v135 := GenS8{F0: v2}
                var v136 float64 = (1.5 + 1.5)
                _ = v134
                _ = v135
                _ = v136
            } else {
                var v137 float64 = -2.25
                v2 = "hello"
                _ = v137
            }
            if (65535 > 32767) {
                switch (fi133 << ((4294967295) & 63)) {
                case 1:
                    for fi138 := 0; fi138 < 17; fi138++ {
                        fmt.Println(v2)
                        fmt.Println(fi133, fi133, v2)
                        v3 -= -1
                        _ = fi138
                    }
                case 127:
                    var v139 uint16 = 42
                    _ = v139
                }
            }
            _ = fi133
        }
        v3 = v3
    } else {
        var v140 int = -105
        _ = v140
    }
    switch 127 {
    case -2147483648:
        _ = len(v4)
    case 42:
        var v141 int64 = 1
        _ = v141
    default:
        if true {
            _ = len(v1)
            var v142 float32 = 1.5
            _ = v142
        } else {
            var v143 int = 161
            if (true && true) {
                switch (v143 & (v143 + 9223372036854775807)) {
                case -884:
                    switch ((v2 + (v2 + (("" + (v2 + ((v2 + "a") + ((v2 + ("\n\t" + v2)) + v2)))) + v2))) + v2) {
                    case "\n\t":
                        v144 := GenS11{F0: -40, F1: 65535, F2: 1}
                        _ = v144
                    case "":
                        if false {
                            for fi145 := 0; fi145 < 11; fi145++ {
                                v2 = v2
                                _ = fi145
                            }
                            switch "" {
                            case "\n\t":
                                switch (v143 + 4294967295) {
                                case 0:
                                    fmt.Println(v143)
                                default:
                                    if (v3 < (^(v3))) {
                                        for fi146 := 0; fi146 < 18; fi146++ {
                                            var v147 uint8 = 22
                                            for fi148 := 0; fi148 < 2; fi148++ {
                                                v2 = (v2 + "foo bar")
                                                v149 := GenS1{F0: (v147 % ((v147) | 1)), F1: 65535, F2: -16}
                                                v143 = (v143 % (((fi148 % ((fi146) | 1))) | 1))
                                                if (!false) {
                                                    var v150 int64 = -1
                                                    _ = v150
                                                }
                                                _ = fi148
                                                _ = v149
                                            }
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            _ = fi146
                                            _ = v147
                                        }
                                        v3 += (-(v3))
                                    }
                                }
                            case "a":
                                switch v2 {
                                case "a":
                                    switch v143 {
                                    case -975:
                                        if (1 > 42) {
                                            var v151 uint16 = 32767
                                            _ = v151
                                        } else {
                                            var v152 uint8 = 201
                                            v153 := GenS8{F0: ("a" + v2)}
                                            if (((0.1 + ((0.1 - float64(1e10)) - 3.14159)) / ((3.14159 - 0.1) - (-2.25 * 1e10))) <= -2.25) {
                                                v2 = "\n\t"
                                            } else {
                                                if (v3 < v3) {
                                                    v154 := GenS10{F0: 2, F1: (254 >= (v143 << ((156) & 63)))}
                                                    if false {
                                                        if (!((!(-92 >= 1)) || true)) {
                                                            if len(v4) > 0 {
                                                                v4[0] = (((v2 + v2) + v2) + v2)
                                                            }
                                                        }
                                                    } else {
                                                        v155 := v1[("\n\t" + ((v2 + "a") + v2))]
                                                        for fi156 := 0; fi156 < 5; fi156++ {
                                                            v155 = v2
                                                            _ = fi156
                                                        }
                                                        var v157 int = v143
                                                        _ = v155
                                                        _ = v157
                                                    }
                                                    _ = v154
                                                } else {
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                }
                                                v158 := GenS2{F0: 42, F1: v143}
                                                if (("\n\t" != v2) && ((false || true) && ((1 == -1) && (0 < 42)))) {
                                                    v159 := GenS12{F0: 862, F1: 526, F2: v2}
                                                    _ = len(v1)
                                                    if (false || false) {
                                                        if true {
                                                            var v160 float64 = (0.0 + 0.1)
                                                            _ = v160
                                                        } else {
                                                            switch "\n\t" {
                                                            case "":
                                                                v161 := GenS5{F0: 42, F1: v3}
                                                                _ = v161
                                                            case "hello":
                                                                var v162 int32 = (v3 % ((v3) | 1))
                                                                _ = v162
                                                            }
                                                            for fi163 := 0; fi163 < 3; fi163++ {
                                                                var v164 uint16 = 42
                                                                for fi165 := 0; fi165 < 11; fi165++ {
                                                                    var v166 uint = 65535
                                                                    v143 += (fi165 / (((^(v143))) | 1))
                                                                    v143 = fi165
                                                                    _ = fi165
                                                                    _ = v166
                                                                }
                                                                fmt.Println(v143, v143, fi163)
                                                                _ = fi163
                                                                _ = v164
                                                            }
                                                        }
                                                        var v167 float64 = (-2.25 / 3.14159)
                                                        for fi168 := 0; fi168 < 6; fi168++ {
                                                            v143 = 97
                                                            _ = len(v4)
                                                            var v169 uint64 = 565
                                                            _ = fi168
                                                            _ = v169
                                                        }
                                                        _ = v167
                                                    } else {
                                                        switch 0 {
                                                        case 9223372036854775807:
                                                            var v170 int16 = 0
                                                            _ = v170
                                                        case 841:
                                                            v171 := v1[v2]
                                                            _ = v171
                                                        default:
                                                            var v172 string = (v2 + v2)
                                                            _ = v172
                                                        }
                                                        v173 := GenS13{F0: 2147483648, F1: 9223372036854775807}
                                                        switch (v143 >> ((0) & 63)) {
                                                        case 400:
                                                            var v174 int8 = 127
                                                            _ = v174
                                                        case -1:
                                                            v175 := GenS6{F0: (^(v152)), F1: -923, F2: true}
                                                            _ = v175
                                                        default:
                                                            v2 = "\n\t"
                                                        }
                                                        _ = v173
                                                    }
                                                    _ = v159
                                                }
                                                _ = v158
                                            }
                                            _ = v152
                                            _ = v153
                                        }
                                    case -255:
                                        var v176 uint16 = 2
                                        _ = v176
                                    }
                                case "":
                                    var v177 int8 = 1
                                    _ = v177
                                case "foo bar":
                                    if false {
                                        switch v2 {
                                        case "":
                                            if (true || (-1 < -153)) {
                                                var v178 int16 = 0
                                                var v179 uint32 = 42
                                                if (v178 <= (v178 | (v178 + (v178 & v178)))) {
                                                    v179 += 753
                                                    for fi180 := 0; fi180 < 6; fi180++ {
                                                        v181 := GenS11{F0: 32767, F1: 898, F2: -1}
                                                        _ = fi180
                                                        _ = v181
                                                    }
                                                    var v182 bool = false
                                                    _ = v182
                                                }
                                                _ = v178
                                                _ = v179
                                            } else {
                                                v3 = v3
                                                v4 = append(v4, ((v2 + v2) + ("a" + v2)))
                                                var v183 int = (v143 << ((2) & 63))
                                                _ = v183
                                            }
                                        default:
                                            if true {
                                                v143 = v143
                                            }
                                        }
                                        fmt.Println(v143)
                                        for fi184 := 0; fi184 < 17; fi184++ {
                                            if (!false) {
                                                _ = len(v4)
                                                v3 += (v3 ^ (v3 % (((v3 | (v3 ^ (^(v3))))) | 1)))
                                                if (true || (1 != 32767)) {
                                                    var v185 int = (v143 ^ fi184)
                                                    _ = v185
                                                } else {
                                                    if (65535 < 2147483647) {
                                                        switch 1 {
                                                        case 1:
                                                            for fi186 := 0; fi186 < 18; fi186++ {
                                                                switch fi184 {
                                                                case -9223372036854775808:
                                                                    if ((!(0 >= -396)) || (!true)) {
                                                                        v187 := GenS1{F0: 1, F1: 42, F2: -1}
                                                                        v3 += -1
                                                                        _ = v187
                                                                    }
                                                                case -1:
                                                                    v2 = ("" + "hello")
                                                                }
                                                                _ = fi186
                                                            }
                                                        case 32767:
                                                            fmt.Println(fi184, v3)
                                                        default:
                                                            var v188 int32 = -930
                                                            _ = v188
                                                        }
                                                        v189 := GenS9{F0: (v3 ^ 2147483647)}
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        _ = v189
                                                    } else {
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        for fi190 := 0; fi190 < 8; fi190++ {
                                                            v2 = v2
                                                            _ = fi190
                                                        }
                                                        fmt.Println(v3, v2, v3)
                                                    }
                                                    var v191 string = v2
                                                    v2 = (v191 + v2)
                                                    _ = v191
                                                }
                                            }
                                            _ = fi184
                                        }
                                    }
                                default:
                                    if false {
                                        fmt.Println(v2, v143, v2)
                                        if (2 > 867) {
                                            v2 = (v2 + ((v2 + v2) + v2))
                                            v143 -= (^(v143))
                                            v3 = v3
                                        } else {
                                            v3 -= (v3 + v3)
                                            var v192 uint32 = 2
                                            v3 += v3
                                            _ = v192
                                        }
                                    } else {
                                        v3 -= (-(v3))
                                        fmt.Println(v2, v143, v2)
                                        fmt.Println(v2)
                                    }
                                }
                            }
                        } else {
                            var v193 uint64 = 16
                            switch v143 {
                            case -815:
                                var v194 bool = ((2147483647 >= 0) && (128 > 1))
                                _ = v194
                            case 32768:
                                switch ((v2 + v2) + v2) {
                                case "a":
                                    var v195 uint64 = (^(v193))
                                    _ = v195
                                case "\n\t":
                                    v143 = (v143 | -2147483648)
                                }
                            case 873:
                                fmt.Println(v193, v2, v143)
                            }
                            var v196 bool = ((0.1 * 0.0) >= ((-2.25 - 0.0) * (0.1 + -2.25)))
                            _ = v193
                            _ = v196
                        }
                    default:
                        v1[((v2 + "foo bar") + v2)] = (v2 + ((v2 + (v2 + v2)) + "\n\t"))
                    }
                case -9223372036854775808:
                    if false {
                        v197 := GenS14{F0: 127}
                        _ = v197
                    }
                case 0:
                    v1["a"] = v2
                }
                switch "" {
                case "":
                    v4 = append(v4, v2)
                case "hello":
                    v3 = v3
                default:
                    v3 = v3
                }
                fmt.Println(v3, v143, v143)
            } else {
                for fi198 := 0; fi198 < 7; fi198++ {
                    delete(v1, (v2 + ""))
                    switch 9223372036854775807 {
                    case 913:
                        switch v2 {
                        case "hello":
                            var v199 int16 = -32768
                            _ = v199
                        default:
                            v143 += 4294967295
                        }
                    }
                    _ = fi198
                }
                switch -1 {
                case -2147483648:
                    v143 = 9223372036854775807
                default:
                    if false {
                        fmt.Println(v3, v2)
                    } else {
                        _ = len(v1)
                    }
                }
            }
            if false {
                for fi200 := 0; fi200 < 18; fi200++ {
                    v201 := v1[(v2 + v2)]
                    v202 := GenS13{F0: 4294967295, F1: (-(v143))}
                    _ = fi200
                    _ = v201
                    _ = v202
                }
                switch (v143 * 42) {
                case 4294967295:
                    for fi203 := 0; fi203 < 2; fi203++ {
                        if (v3 < v3) {
                            var v204 int = v143
                            _ = v204
                        }
                        fmt.Println(v2, v143, fi203)
                        _ = fi203
                    }
                case 860:
                    v1[((v2 + v2) + (v2 + (v2 + v2)))] = "hello"
                case -882:
                    var v205 string = v2
                    _ = v205
                }
                var v206 int32 = (v3 / ((v3) | 1))
                _ = v206
            }
            _ = v143
        }
    }
    var v207 uint64 = 1
    v208 := GenS15{F0: 2, F1: 224}
    v2 = v2
    fmt.Println(v2)
    var v209 float32 = ((float32((0.0 * -2.25)) / (0.1 + -2.25)) + (((3.14159 * (1.5 * 3.14159)) * 1.5) + 3.14159))
    v207 = v207
    v210 := GenS16{F0: 1000, F1: 53}
    var v211 int = 32768
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v207:", v207)
    fmt.Println("v209:", v209)
    fmt.Println("v211:", v211)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v207
    _ = v208
    _ = v209
    _ = v210
    _ = v211
    _ = v3
    _ = v4
}
