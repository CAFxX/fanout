package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int; F1 bool; F2 uint64 }
type GenS6 struct { F0 uint16; F1 string; F2 uint }
type GenS7 struct { F0 int64; F1 uint32; F2 int64 }
type GenS8 struct { F0 int64; F1 bool; F2 int }
type GenS9 struct { F0 int16; F1 uint32; F2 int64 }
type GenS10 struct { F0 string; F1 int32; F2 uint }
type GenS11 struct { F0 int }
type GenS12 struct { F0 int; F1 int32 }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 string; F1 int64 }
type GenS15 struct { F0 bool; F1 bool }
type GenS16 struct { F0 int; F1 uint64 }
type GenS17 struct { F0 uint }
type GenS18 struct { F0 int; F1 int8; F2 string }
type GenS19 struct { F0 uint64 }
type GenS20 struct { F0 string; F1 uint16 }
type GenS21 struct { F0 bool }
type GenS22 struct { F0 uint64; F1 uint64 }
type GenS23 struct { F0 string }

func main() {
    var v1 uint16 = 2
    var v2 int8 = -9
    var v3 int32 = -1
    var v4 int = 0
    for fi5 := 0; fi5 < 20; fi5++ {
        v6 := []int{v4, fi5, (v4 * (v4 % ((fi5) | 1)))}
        if (!(("hello" + "foo bar") != "hello")) {
            v6 = append(v6, fi5)
            fmt.Println(v1, fi5)
            if false {
                if ((true && (true && (false && true))) || false) {
                    if true {
                        var v7 uint64 = 4294967295
                        _ = v7
                    } else {
                        switch "foo bar" {
                        case "a":
                            if (!(255 != 1)) {
                                v8 := GenS1{F0: "hello", F1: 144, F2: 42}
                                fmt.Println(v2, v4)
                                _ = v8
                            }
                        case "hello":
                            for fi9 := 0; fi9 < 14; fi9++ {
                                fmt.Println(fi9, fi9)
                                if true {
                                    fmt.Println(fi5)
                                    switch fi9 {
                                    case 32768:
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                    default:
                                        v10 := GenS2{F0: 2}
                                        _ = v10
                                    }
                                    if false {
                                        var v11 uint64 = 4294967295
                                        _ = v11
                                    } else {
                                        v12 := GenS3{F0: (v1 | (v1 % (((v1 / ((v1) | 1))) | 1)))}
                                        _ = v12
                                    }
                                }
                                var v13 uint8 = 42
                                switch fi5 {
                                case 42:
                                    var v14 float32 = (0.0 * 0.1)
                                    _ = v14
                                case -9223372036854775808:
                                    if (v4 <= (fi9 ^ fi9)) {
                                        if true {
                                            var v15 uint16 = 133
                                            v16 := GenS4{F0: -9223372036854775808}
                                            _ = v15
                                            _ = v16
                                        } else {
                                            v17 := GenS2{F0: 42}
                                            fmt.Println(v1)
                                            var v18 string = "\n\t"
                                            var v19 bool = false
                                            var v20 uint32 = 2147483648
                                            _ = v17
                                            _ = v18
                                            _ = v19
                                            _ = v20
                                        }
                                    }
                                default:
                                    var v21 int = 1
                                    _ = v21
                                }
                                _ = fi9
                                _ = v13
                            }
                        case "":
                            v2 = (v2 % ((int8(v3)) | 1))
                        }
                        v2 -= (v2 - v2)
                    }
                    v1 = (v1 | v1)
                }
                if false {
                    v3 = (v3 / (((v3 * (v3 >> ((255) & 63)))) | 1))
                    var v22 float64 = 3.14159
                    _ = v22
                } else {
                    if len(v6) > 0 {
                        v6[0] = (fi5 | fi5)
                    }
                    if (false && (!true)) {
                        v23 := GenS4{F0: 582}
                        _ = v23
                    }
                }
                var v24 int32 = 2147483647
                _ = v24
            } else {
                v1 = v1
                if len(v6) > 0 {
                    v6[0] = (v4 ^ (fi5 << ((4294967295) & 63)))
                }
                var v25 int32 = v3
                _ = v25
            }
        } else {
            fmt.Println(v2, v4, v1)
            fmt.Println(v3, v4)
        }
        v6 = append(v6, v4)
        _ = fi5
        _ = v6
    }
    var v26 int64 = 9223372036854775807
    if false {
        v26 -= (v26 & 42)
        v27 := []int{v4, v4, (v4 % ((v4) | 1)), 9223372036854775807}
        v3 -= (^(v3))
        _ = v27
    }
    v26 += (v26 & 1)
    if false {
        var v28 int = v4
        for fi29 := 0; fi29 < 11; fi29++ {
            var v30 int16 = 125
            _ = fi29
            _ = v30
        }
        if (65535 == 292) {
            v31 := GenS5{F0: 4294967295, F1: true, F2: 4294967295}
            v32 := GenS1{F0: "", F1: 0, F2: 42}
            switch ("foo bar" + ("" + "hello")) {
            case "a":
                for fi33 := 0; fi33 < 6; fi33++ {
                    fmt.Println(v26, fi33, v3)
                    if false {
                        v3 = (v3 << ((42) & 63))
                        v4 = -9223372036854775808
                        switch "hello" {
                        case "a":
                            fmt.Println(v1)
                        }
                    } else {
                        var v34 uint8 = 221
                        var v35 float32 = 1.5
                        v28 = 32768
                        _ = v34
                        _ = v35
                    }
                    _ = fi33
                }
            case "":
                var v36 int8 = (v2 + (v2 | -36))
                _ = v36
            case "foo bar":
                v26 += v26
            default:
                v37 := GenS6{F0: 42, F1: "foo bar", F2: 2}
                _ = v37
            }
            _ = v31
            _ = v32
        } else {
            v38 := GenS5{F0: 668, F1: false, F2: 18446744073709551615}
            v4 += v28
            v39 := []int{v4, (v28 << ((0) & 63))}
            _ = v38
            _ = v39
        }
        _ = v28
    } else {
        v40 := GenS1{F0: "a", F1: 1, F2: 738}
        v1 = (v1 >> ((164) & 63))
        _ = v40
    }
    switch v4 {
    case 127:
        switch "hello" {
        case "foo bar":
            v26 += v26
        case "":
            v41 := GenS7{F0: (v26 + (v26 & -737)), F1: 2147483648, F2: (^(v26))}
            _ = v41
        default:
            if (2147483647 <= 2147483648) {
                for fi42 := 0; fi42 < 7; fi42++ {
                    fmt.Println(fi42)
                    _ = fi42
                }
            }
        }
    default:
        v43 := []int{(v4 - v4), (v4 - v4), 127, (v4 << ((0) & 63))}
        _ = v43
    }
    v1 -= v1
    if (v1 <= (v1 % (((v1 & 1)) | 1))) {
        fmt.Println(v4, v2)
    } else {
        v26 -= v26
        for fi44 := 0; fi44 < 2; fi44++ {
            v45 := GenS8{F0: (v26 | int64(fi44)), F1: true, F2: -947}
            if ((128 == 0) || (-264 >= v26)) {
                if ((fi44 < v4) && (231 <= v1)) {
                    v3 += (v3 | v3)
                    var v46 uint8 = 1
                    var v47 uint8 = v46
                    _ = v46
                    _ = v47
                } else {
                    v26 -= (v26 + (-(v26)))
                    for fi48 := 0; fi48 < 2; fi48++ {
                        v49 := []int{-199}
                        _ = len(v49)
                        _ = fi48
                        _ = v49
                    }
                    if true {
                        v50 := []int{fi44, v4, (fi44 - (fi44 & fi44)), int(v2)}
                        _ = v50
                    }
                }
                for fi51 := 0; fi51 < 3; fi51++ {
                    var v52 uint64 = 315
                    _ = fi51
                    _ = v52
                }
            } else {
                var v53 int8 = 0
                v54 := GenS9{F0: -939, F1: 4294967295, F2: 1}
                for fi55 := 0; fi55 < 11; fi55++ {
                    v56 := []int{9223372036854775807, v4, v4}
                    _ = len(v56)
                    v4 = 551
                    _ = fi55
                    _ = v56
                }
                _ = v53
                _ = v54
            }
            var v57 int16 = int16(v3)
            _ = fi44
            _ = v45
            _ = v57
        }
        var v58 uint32 = 10
        _ = v58
    }
    for fi59 := 0; fi59 < 19; fi59++ {
        fmt.Println(v4, v4)
        _ = fi59
    }
    if (!true) {
        for fi60 := 0; fi60 < 20; fi60++ {
            v61 := GenS8{F0: 9223372036854775807, F1: (-117 > 32767), F2: v4}
            v62 := GenS1{F0: "hello", F1: 42, F2: 42}
            v63 := []int{v4, -9223372036854775808}
            _ = fi60
            _ = v61
            _ = v62
            _ = v63
        }
        v26 = v26
    }
    var v64 bool = true
    if (((^(v4)) > v4) && v64) {
        v65 := []int{}
        v66 := GenS9{F0: 365, F1: 42, F2: 42}
        _ = v65
        _ = v66
    }
    v4 += 2147483647
    if false {
        switch "\n\t" {
        case "\n\t":
            for fi67 := 0; fi67 < 7; fi67++ {
                v68 := GenS10{F0: (("\n\t" + "") + ""), F1: (v3 * v3), F2: 4294967295}
                switch 42 {
                case 42:
                    v2 = 127
                case -67:
                    var v69 float32 = 0.1
                    _ = v69
                default:
                    v70 := GenS4{F0: (v26 + int64(fi67))}
                    _ = v70
                }
                _ = fi67
                _ = v68
            }
        }
        for fi71 := 0; fi71 < 20; fi71++ {
            switch fi71 {
            case 127:
                if (v26 < (v26 | v26)) {
                    if false {
                        fmt.Println(v64)
                        fmt.Println(v26, v2)
                        v72 := []int{372}
                        _ = v72
                    }
                    for fi73 := 0; fi73 < 15; fi73++ {
                        var v74 float32 = 0.1
                        v75 := GenS3{F0: 2}
                        _ = fi73
                        _ = v74
                        _ = v75
                    }
                } else {
                    for fi76 := 0; fi76 < 11; fi76++ {
                        if (v64 || true) {
                            switch fi71 {
                            case -180:
                                var v77 float32 = (0.0 - 1.5)
                                _ = v77
                            case -868:
                                v78 := GenS11{F0: fi76}
                                _ = v78
                            case -277:
                                v79 := GenS3{F0: 2}
                                _ = v79
                            }
                            var v80 uint16 = 65535
                            v81 := GenS5{F0: fi76, F1: v64, F2: 619}
                            _ = v80
                            _ = v81
                        } else {
                            if ((-2.25 - 1e10) <= 1.5) {
                                if true {
                                    if v64 {
                                        v82 := GenS12{F0: (-(fi71)), F1: (v3 + (v3 - (v3 | -1)))}
                                        if ("" == "") {
                                            v83 := GenS13{F0: v3}
                                            var v84 int = (fi76 % ((v4) | 1))
                                            if true {
                                                var v85 bool = true
                                                v4 += 1
                                                v1 -= v1
                                                _ = v85
                                            }
                                            _ = v83
                                            _ = v84
                                        } else {
                                            v64 = v64
                                            var v86 uint8 = 122
                                            _ = v86
                                        }
                                        var v87 int16 = 32767
                                        var v88 int8 = (v2 % ((v2) | 1))
                                        fmt.Println(v1, fi76, v64)
                                        v89 := []int{(fi71 ^ fi76)}
                                        fmt.Println(v64)
                                        if true {
                                            var v90 uint16 = 255
                                            _ = v90
                                        } else {
                                            v4 = (fi71 / ((-2147483648) | 1))
                                            var v91 string = "foo bar"
                                            v4 += 2147483647
                                            v88 = v88
                                            _ = v91
                                        }
                                        _ = v82
                                        _ = v87
                                        _ = v88
                                        _ = v89
                                    } else {
                                        v3 -= (v3 + v3)
                                        v92 := []int{2147483647, v4}
                                        if (v64 && v64) {
                                            if (!(true || (42 >= 90))) {
                                                v64 = (!v64)
                                                var v93 uint32 = 2
                                                if (v93 > 0) {
                                                    var v94 int8 = v2
                                                    _ = v94
                                                }
                                                _ = v93
                                            } else {
                                                fmt.Println(v4, fi71)
                                            }
                                        } else {
                                            v3 = (v3 | 442)
                                            v95 := GenS14{F0: ("foo bar" + ("hello" + ((("\n\t" + "hello") + "\n\t") + ("hello" + "\n\t")))), F1: v26}
                                            v4 -= v4
                                            _ = v95
                                        }
                                        _ = v92
                                    }
                                    var v96 int32 = (v3 ^ 42)
                                    v97 := GenS9{F0: 1, F1: 456, F2: (v26 / (((v26 % ((v26) | 1))) | 1))}
                                    _ = v96
                                    _ = v97
                                }
                                v98 := GenS2{F0: 42}
                                fmt.Println(v3)
                                _ = v98
                            } else {
                                if (!true) {
                                    var v99 int = (fi76 & (fi76 & (v4 ^ (v4 << ((0) & 63)))))
                                    v100 := []int{v4, (fi76 >> ((255) & 63)), v99, fi76}
                                    var v101 float32 = 1.5
                                    _ = v100
                                    _ = v101
                                    _ = v99
                                } else {
                                    v102 := []int{fi71, (fi71 * (v4 / (((fi76 << ((18446744073709551615) & 63))) | 1))), v4}
                                    v102 = append(v102, -854)
                                    _ = v102
                                }
                                if (42 == 90) {
                                    v103 := GenS11{F0: (v4 << ((2) & 63))}
                                    v104 := []int{-9223372036854775808, 128}
                                    fmt.Println(v4, v26)
                                    _ = v103
                                    _ = v104
                                }
                                var v105 int32 = v3
                                _ = v105
                            }
                            v106 := []int{}
                            if len(v106) > 0 {
                                v106[0] = (fi71 ^ 127)
                            }
                            _ = v106
                        }
                        switch (("a" + "foo bar") + (("\n\t" + "foo bar") + (("foo bar" + "") + (("foo bar" + "hello") + ("" + (((("a" + "foo bar") + "foo bar") + ("" + ("foo bar" + ("\n\t" + ("foo bar" + "hello"))))) + ("" + "\n\t"))))))) {
                        case "foo bar":
                            v107 := []int{(-(fi76)), (fi71 << ((331) & 63)), v4}
                            _ = v107
                        case "hello":
                            v108 := []int{(v4 / ((32768) | 1)), (fi76 - v4), 2147483647, 0}
                            _ = v108
                        case "":
                            v109 := GenS15{F0: (!false), F1: v64}
                            _ = v109
                        default:
                            v110 := GenS10{F0: ("foo bar" + "foo bar"), F1: (v3 << ((4294967295) & 63)), F2: 18446744073709551615}
                            _ = v110
                        }
                        _ = fi76
                    }
                    fmt.Println(v4, v2)
                    fmt.Println(v1, v1, v4)
                }
            case 0:
                var v111 float64 = 1e10
                _ = v111
            default:
                if v64 {
                    if false {
                        v112 := GenS11{F0: fi71}
                        v113 := GenS8{F0: 801, F1: true, F2: (fi71 % ((-610) | 1))}
                        v114 := []int{fi71, (v4 & (^(fi71))), (v4 * (v4 / ((4294967295) | 1)))}
                        _ = v112
                        _ = v113
                        _ = v114
                    } else {
                        if (!((61 >= 112) && (-2.25 >= 0.1))) {
                            switch (-(v4)) {
                            case 32767:
                                for fi115 := 0; fi115 < 3; fi115++ {
                                    v3 = v3
                                    if (-32768 != 47) {
                                        fmt.Println(v1, fi71)
                                        switch fi71 {
                                        case -200:
                                            if true {
                                                if ((239 < 4294967295) && ((v26 << ((448) & 63)) > -866)) {
                                                    switch ("\n\t" + "a") {
                                                    case "hello":
                                                        fmt.Println(v3, v1)
                                                    default:
                                                        v2 = v2
                                                    }
                                                } else {
                                                    var v116 uint64 = 0
                                                    var v117 uint = 940
                                                    fmt.Println(v1, v1, v117)
                                                    _ = v116
                                                    _ = v117
                                                }
                                                switch v4 {
                                                case 128:
                                                    var v118 int = fi71
                                                    _ = v118
                                                case -940:
                                                    var v119 int16 = -1
                                                    _ = v119
                                                case 9223372036854775807:
                                                    v120 := GenS11{F0: v4}
                                                    _ = v120
                                                }
                                                v121 := []int{}
                                                _ = v121
                                            } else {
                                                fmt.Println(v2)
                                                var v122 int = fi115
                                                _ = v122
                                            }
                                        case 128:
                                            switch (fi115 << ((65535) & 63)) {
                                            case 2147483647:
                                                v123 := GenS15{F0: v64, F1: v64}
                                                _ = v123
                                            case 32767:
                                                v26 = v26
                                                switch (("hello" + "\n\t") + "hello") {
                                                case "":
                                                    v124 := []int{fi115, 42, (fi115 ^ -1)}
                                                    _ = v124
                                                default:
                                                    v3 -= (v3 + v3)
                                                }
                                                var v125 int8 = v2
                                                _ = v125
                                            case 1:
                                                if (4294967295 == 527) {
                                                    v26 = v26
                                                    switch "foo bar" {
                                                    case "":
                                                        switch 127 {
                                                        case 42:
                                                            if v64 {
                                                                switch ("hello" + "hello") {
                                                                case "hello":
                                                                    var v126 int32 = -20
                                                                    _ = v126
                                                                case "foo bar":
                                                                    v127 := []int{fi115, (fi71 * (v4 >> ((65535) & 63)))}
                                                                    _ = v127
                                                                }
                                                                if (-2.25 > 1e10) {
                                                                    v128 := GenS16{F0: fi71, F1: 1}
                                                                    v26 = 0
                                                                    switch "\n\t" {
                                                                    case "a":
                                                                        v129 := GenS3{F0: (v1 / ((v1) | 1))}
                                                                        _ = v129
                                                                    case "\n\t":
                                                                        v1 = v1
                                                                    case "hello":
                                                                        var v130 uint8 = 127
                                                                        _ = v130
                                                                    }
                                                                    _ = v128
                                                                }
                                                            }
                                                        case -716:
                                                            if (2 <= 4294967295) {
                                                                var v131 uint8 = 0
                                                                _ = v131
                                                            } else {
                                                                v132 := GenS17{F0: 65535}
                                                                if (int32(v26) < v3) {
                                                                    v133 := []int{int(v2), 42}
                                                                    v4 = (v4 << ((0) & 63))
                                                                    _ = v133
                                                                } else {
                                                                    v134 := []int{}
                                                                    var v135 uint16 = v1
                                                                    fmt.Println(v26, v26, fi115)
                                                                    var v136 int64 = (v26 + 0)
                                                                    _ = v134
                                                                    _ = v135
                                                                    _ = v136
                                                                }
                                                                switch "hello" {
                                                                case "":
                                                                    v4 = -768
                                                                case "foo bar":
                                                                    if (1 > 18446744073709551615) {
                                                                        v2 = v2
                                                                        v137 := GenS18{F0: (fi71 * (fi115 % (((fi71 & int(v2))) | 1))), F1: v2, F2: "\n\t"}
                                                                        _ = v137
                                                                    }
                                                                case "\n\t":
                                                                    v138 := GenS19{F0: 18446744073709551615}
                                                                    _ = v138
                                                                default:
                                                                    v3 -= (v3 | v3)
                                                                }
                                                                _ = v132
                                                            }
                                                            switch (v4 & -2147483648) {
                                                            case 1:
                                                                var v139 int16 = 42
                                                                _ = v139
                                                            }
                                                            if v64 {
                                                                switch (v4 | 2147483647) {
                                                                case 0:
                                                                    switch v4 {
                                                                    case 2147483647:
                                                                        v140 := []int{(v4 - fi115), int(v3)}
                                                                        _ = v140
                                                                    }
                                                                case -2147483648:
                                                                    var v141 uint8 = 127
                                                                    _ = v141
                                                                default:
                                                                    switch "a" {
                                                                    case "foo bar":
                                                                        v1 = (v1 | (v1 - (v1 & (^(v1)))))
                                                                        fmt.Println(fi71, fi71, fi71)
                                                                        v3 += (v3 | v3)
                                                                    case "":
                                                                        var v142 bool = v64
                                                                        v3 = v3
                                                                        if (("hello" + "hello") > "a") {
                                                                            fmt.Println(v2, v26)
                                                                            if v142 {
                                                                                v4 = -91
                                                                            } else {
                                                                                var v143 uint64 = 341
                                                                                v1 += (v1 & 216)
                                                                                _ = v143
                                                                            }
                                                                        } else {
                                                                            v144 := []int{fi115}
                                                                            _ = v144
                                                                        }
                                                                        _ = v142
                                                                    }
                                                                    v145 := GenS5{F0: -688, F1: (3.14159 > 3.14159), F2: 2}
                                                                    var v146 int64 = -9223372036854775808
                                                                    _ = v145
                                                                    _ = v146
                                                                }
                                                                var v147 int32 = v3
                                                                _ = v147
                                                            } else {
                                                                var v148 uint16 = v1
                                                                var v149 bool = (255 > 128)
                                                                _ = v148
                                                                _ = v149
                                                            }
                                                        case -1:
                                                            v3 += (v3 * 0)
                                                        }
                                                    case "foo bar":
                                                        fmt.Println(v26)
                                                    default:
                                                        fmt.Println(fi71, v1)
                                                    }
                                                    fmt.Println(v64)
                                                } else {
                                                    var v150 float32 = (((-2.25 - (3.14159 + 0.0)) - (((0.0 + 1e10) / -2.25) + -2.25)) / ((0.1 + 0.0) / float32(3.14159)))
                                                    _ = v150
                                                }
                                            }
                                            var v151 uint16 = 138
                                            _ = v151
                                        }
                                        switch fi115 {
                                        case 32768:
                                            var v152 int32 = (v3 & v3)
                                            _ = v152
                                        }
                                    } else {
                                        if (3.14159 == 1.5) {
                                            v64 = v64
                                            if v64 {
                                                if (-54 != 42) {
                                                    var v153 float32 = 0.0
                                                    v153 = float32((1.5 * 3.14159))
                                                    var v154 uint32 = 752
                                                    _ = v153
                                                    _ = v154
                                                } else {
                                                    v155 := []int{}
                                                    if false {
                                                        var v156 uint16 = (v1 >> ((482) & 63))
                                                        if (false && v64) {
                                                            v157 := GenS1{F0: "\n\t", F1: 1, F2: 0}
                                                            var v158 bool = ((fi71 % ((-450) | 1)) == 42)
                                                            _ = v157
                                                            _ = v158
                                                        } else {
                                                            switch ("" + "") {
                                                            case "foo bar":
                                                                fmt.Println(v156)
                                                            }
                                                            var v159 int32 = v3
                                                            var v160 uint = 107
                                                            v161 := GenS20{F0: ((((("hello" + ((("foo bar" + (("hello" + "a") + (("hello" + "foo bar") + "foo bar"))) + ("\n\t" + "a")) + (("\n\t" + "\n\t") + ("\n\t" + "")))) + "\n\t") + ((("a" + "foo bar") + ((("foo bar" + "a") + "foo bar") + "")) + "")) + "foo bar") + ("hello" + "hello")), F1: (v156 | (v1 ^ v156))}
                                                            fmt.Println(v1, v3)
                                                            _ = v159
                                                            _ = v160
                                                            _ = v161
                                                        }
                                                        v1 += v156
                                                        _ = v156
                                                    } else {
                                                        v3 -= (v3 % ((int32(v4)) | 1))
                                                        var v162 uint32 = 42
                                                        var v163 uint64 = 481
                                                        _ = v162
                                                        _ = v163
                                                    }
                                                    switch (("" + "a") + ("foo bar" + "hello")) {
                                                    case "\n\t":
                                                        var v164 uint32 = 42
                                                        _ = v164
                                                    case "foo bar":
                                                        var v165 uint16 = v1
                                                        _ = v165
                                                    default:
                                                        _ = len(v155)
                                                    }
                                                    _ = v155
                                                }
                                                v64 = v64
                                            }
                                        } else {
                                            switch "\n\t" {
                                            case "foo bar":
                                                var v166 uint = 375
                                                var v167 int64 = (v26 % (((v26 % (((v26 % ((v26) | 1))) | 1))) | 1))
                                                _ = v166
                                                _ = v167
                                            case "\n\t":
                                                switch (v4 ^ v4) {
                                                case -1:
                                                    v3 = (v3 & v3)
                                                case 786:
                                                    v2 = (^(v2))
                                                case 42:
                                                    v26 = v26
                                                }
                                            }
                                            if true {
                                                v2 -= (^(v2))
                                                fmt.Println(v26)
                                            } else {
                                                v26 += (-(v26))
                                                v168 := GenS13{F0: v3}
                                                if v64 {
                                                    v169 := GenS15{F0: v64, F1: ((v2 >> ((4294967295) & 63)) <= v2)}
                                                    fmt.Println(v2, v2, fi71)
                                                    v170 := []int{}
                                                    _ = v169
                                                    _ = v170
                                                }
                                                var v171 int64 = v26
                                                var v172 string = (("foo bar" + "\n\t") + "")
                                                _ = v168
                                                _ = v171
                                                _ = v172
                                            }
                                        }
                                        if v64 {
                                            fmt.Println(v3, fi115)
                                            if v64 {
                                                var v173 int16 = 42
                                                if v64 {
                                                    var v174 string = (("hello" + "") + "foo bar")
                                                    _ = v174
                                                }
                                                _ = v173
                                            } else {
                                                var v175 int16 = -1
                                                _ = v175
                                            }
                                            v4 += (fi115 + 2147483647)
                                            v176 := GenS9{F0: 32767, F1: 1, F2: v26}
                                            v177 := []int{(-(fi71))}
                                            _ = v176
                                            _ = v177
                                        }
                                        v3 = -1
                                    }
                                    _ = fi115
                                }
                            }
                            for fi178 := 0; fi178 < 4; fi178++ {
                                var v179 int16 = -130
                                var v180 uint64 = 2
                                v26 = (v26 - v26)
                                v181 := GenS1{F0: "hello", F1: 0, F2: (v179 / ((-371) | 1))}
                                if ((0 <= v3) || v64) {
                                    v179 = 32767
                                    switch (fi178 * (^(fi71))) {
                                    case -1:
                                        if (-2.25 >= 0.0) {
                                            v3 = -878
                                            v64 = (fi178 > fi178)
                                            if true {
                                                var v182 int32 = int32(v180)
                                                switch (fi178 << ((18446744073709551615) & 63)) {
                                                case -427:
                                                    fmt.Println(fi71)
                                                case -2147483648:
                                                    v183 := []int{}
                                                    _ = v183
                                                default:
                                                    switch -2147483648 {
                                                    case 686:
                                                        var v184 float64 = 1.5
                                                        _ = v184
                                                    }
                                                }
                                                _ = v182
                                            }
                                        }
                                    case 32768:
                                        v180 = 2
                                    }
                                }
                                _ = fi178
                                _ = v179
                                _ = v180
                                _ = v181
                            }
                        } else {
                            var v185 uint16 = (v1 >> ((553) & 63))
                            v186 := GenS20{F0: "a", F1: v1}
                            _ = v185
                            _ = v186
                        }
                    }
                    if false {
                        v3 += v3
                    } else {
                        if v64 {
                            fmt.Println(fi71)
                            switch fi71 {
                            case -929:
                                v187 := []int{(v4 | -602)}
                                _ = v187
                            case 211:
                                switch v4 {
                                case 32768:
                                    switch (fi71 - (fi71 | fi71)) {
                                    case -1:
                                        v188 := GenS17{F0: 18446744073709551615}
                                        _ = v188
                                    case -9223372036854775808:
                                        v189 := []int{(v4 % (((-(fi71))) | 1)), -552, fi71}
                                        _ = v189
                                    default:
                                        if (!false) {
                                            var v190 uint32 = 0
                                            switch ("" + "a") {
                                            case "":
                                                for fi191 := 0; fi191 < 4; fi191++ {
                                                    switch (v4 + fi71) {
                                                    case 998:
                                                        var v192 int16 = 1
                                                        _ = v192
                                                    case 42:
                                                        v193 := []int{9223372036854775807, (fi71 ^ -1)}
                                                        _ = v193
                                                    case 9223372036854775807:
                                                        var v194 float32 = 0.1
                                                        _ = v194
                                                    }
                                                    _ = fi191
                                                }
                                            case "foo bar":
                                                if v64 {
                                                    v26 += v26
                                                }
                                            }
                                            _ = v190
                                        } else {
                                            v195 := []int{2147483647, 360}
                                            v196 := GenS20{F0: "\n\t", F1: v1}
                                            _ = v195
                                            _ = v196
                                        }
                                    }
                                case -2147483648:
                                    if false {
                                        fmt.Println(v2)
                                    }
                                }
                            case 32767:
                                v1 = v1
                            default:
                                v1 -= 65535
                            }
                        } else {
                            var v197 string = "a"
                            if v64 {
                                fmt.Println(v26, v3)
                            } else {
                                if v64 {
                                    var v198 uint = 2
                                    var v199 bool = true
                                    for fi200 := 0; fi200 < 18; fi200++ {
                                        var v201 int64 = (v26 - (v26 % (((v26 << ((0) & 63))) | 1)))
                                        v201 -= (v26 ^ (-(v201)))
                                        _ = fi200
                                        _ = v201
                                    }
                                    _ = v198
                                    _ = v199
                                } else {
                                    v202 := []int{(fi71 >> ((65535) & 63)), fi71, v4, (v4 - fi71)}
                                    v2 -= v2
                                    v2 += v2
                                    _ = v202
                                }
                                v203 := GenS12{F0: -1, F1: 227}
                                _ = v203
                            }
                            if v64 {
                                for fi204 := 0; fi204 < 4; fi204++ {
                                    if (((float32(float64(0.0)) / (3.14159 * (3.14159 - float32((1.5 - 0.1))))) - ((float32(1e10) * float32(1.5)) * (1.5 + ((0.1 - ((3.14159 + 1e10) - 0.1)) - (1.5 + (0.1 - 1.5)))))) != ((1e10 / 1.0) / float32(0.1))) {
                                        v197 = "foo bar"
                                        v205 := []int{fi204, 1, (fi204 ^ (fi71 * (v4 >> ((2) & 63)))), -9223372036854775808}
                                        switch fi71 {
                                        case 44:
                                            fmt.Println(v197, v197)
                                        case -93:
                                            v3 = (v3 / (((v3 - 1)) | 1))
                                        }
                                        switch v197 {
                                        case "hello":
                                            v4 += (fi204 >> ((65535) & 63))
                                        case "foo bar":
                                            var v206 int = v4
                                            _ = v206
                                        case "a":
                                            fmt.Println(v2, v2, v64)
                                        }
                                        v207 := GenS10{F0: "", F1: 1, F2: 696}
                                        v205 = append(v205, (fi71 * -1))
                                        var v208 int16 = 308
                                        switch (^(v4)) {
                                        case 128:
                                            if ((v3 << ((18446744073709551615) & 63)) < int32(v2)) {
                                                var v209 uint8 = 65
                                                if len(v205) >= 2 {
                                                    v205 = v205[1:len(v205)]
                                                }
                                                _ = len(v205)
                                                _ = v209
                                            } else {
                                                var v210 float64 = (float64(0.1) + 3.14159)
                                                v211 := GenS11{F0: 2147483647}
                                                _ = v210
                                                _ = v211
                                            }
                                        }
                                        _ = v205
                                        _ = v207
                                        _ = v208
                                    }
                                    var v212 string = v197
                                    _ = fi204
                                    _ = v212
                                }
                                var v213 bool = (((!(1 >= 1)) && true) && ((float32(3.14159) - float32(1e10)) >= ((3.14159 / 3.14159) - 0.0)))
                                v214 := GenS21{F0: false}
                                _ = v213
                                _ = v214
                            }
                            _ = v197
                        }
                        if ((v2 > v2) && (("" + "\n\t") < ((("hello" + "hello") + ("hello" + "foo bar")) + "foo bar"))) {
                            v64 = false
                            v215 := []int{(fi71 | (-(v4))), int(v2), fi71, v4}
                            if (!((v2 & v2) < 113)) {
                                switch "hello" {
                                case "a":
                                    if len(v215) >= 2 {
                                        v215 = v215[1:len(v215)]
                                    }
                                case "foo bar":
                                    switch -444 {
                                    case -9223372036854775808:
                                        var v216 int32 = (v3 + 825)
                                        _ = v216
                                    default:
                                        var v217 float64 = float64(3.14159)
                                        _ = v217
                                    }
                                }
                            } else {
                                for fi218 := 0; fi218 < 19; fi218++ {
                                    v3 = v3
                                    if v64 {
                                        v219 := GenS22{F0: 2, F1: 2}
                                        var v220 float32 = 1e10
                                        var v221 int = fi71
                                        _ = v219
                                        _ = v220
                                        _ = v221
                                    } else {
                                        v222 := GenS17{F0: 163}
                                        _ = v222
                                    }
                                    _ = fi218
                                }
                            }
                            _ = v215
                        } else {
                            v223 := []int{(fi71 / ((v4) | 1)), -9223372036854775808, v4, (fi71 | fi71)}
                            var v224 float64 = ((-2.25 - 3.14159) - (1e10 - 1e10))
                            _ = v223
                            _ = v224
                        }
                    }
                } else {
                    switch fi71 {
                    case 9223372036854775807:
                        v4 = fi71
                    case 128:
                        for fi225 := 0; fi225 < 4; fi225++ {
                            if true {
                                var v226 string = "hello"
                                var v227 float64 = ((1e10 + -2.25) - 1e10)
                                _ = v226
                                _ = v227
                            }
                            _ = fi225
                        }
                    default:
                        switch "" {
                        case "foo bar":
                            v4 += 919
                        case "a":
                            v4 += fi71
                        }
                    }
                    var v228 string = "a"
                    _ = v228
                }
            }
            v1 += v1
            v229 := GenS11{F0: fi71}
            _ = fi71
            _ = v229
        }
    } else {
        var v230 int = v4
        if v64 {
            fmt.Println(v3)
            v231 := []int{(v4 / ((128) | 1)), -9223372036854775808}
            _ = v231
        } else {
            if (65535 != 65535) {
                v3 += -1
            }
            switch ("hello" + ("foo bar" + "a")) {
            case "a":
                v1 += (v1 & (v1 - v1))
            case "":
                v2 = (v2 - v2)
            }
            if (v64 && ((("foo bar" + "\n\t") + (("foo bar" + "hello") + "a")) <= "a")) {
                for fi232 := 0; fi232 < 15; fi232++ {
                    v230 = (v230 ^ (-(fi232)))
                    v26 -= 0
                    v4 -= fi232
                    _ = fi232
                }
                v233 := []int{v230, v4}
                _ = v233
            }
        }
        if true {
            var v234 bool = (42 != int8(v3))
            switch 416 {
            case 32767:
                fmt.Println(v64, v4, v234)
            case 128:
                switch "hello" {
                case "\n\t":
                    fmt.Println(v4, v26, v2)
                }
            case 42:
                var v235 float64 = 0.0
                _ = v235
            }
            _ = v234
        } else {
            v2 = int8(v26)
            fmt.Println(v2, v64)
            if v64 {
                v230 = 210
            } else {
                if v64 {
                    v236 := GenS23{F0: "\n\t"}
                    v230 = (v230 * (v4 | v230))
                    _ = v236
                }
                var v237 uint64 = 221
                if (!v64) {
                    var v238 float32 = 1.5
                    for fi239 := 0; fi239 < 17; fi239++ {
                        var v240 uint16 = v1
                        _ = fi239
                        _ = v240
                    }
                    _ = v238
                }
                _ = v237
            }
        }
        _ = v230
    }
    fmt.Println(v64, v1)
    v241 := GenS17{F0: uint(v1)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v241
    _ = v26
    _ = v3
    _ = v4
    _ = v64
}
