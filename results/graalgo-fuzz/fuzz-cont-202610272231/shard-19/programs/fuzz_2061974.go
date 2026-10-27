package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 uint; F2 uint }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint8; F1 int64; F2 int64 }
type GenS5 struct { F0 uint; F1 bool }
type GenS6 struct { F0 uint32; F1 int64; F2 int8 }
type GenS7 struct { F0 uint32; F1 int32 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint16; F1 uint64 }
type GenS10 struct { F0 int16; F1 string; F2 uint64 }
type GenS11 struct { F0 int16; F1 string }
type GenS12 struct { F0 int8; F1 int64; F2 uint64 }
type GenS13 struct { F0 int8; F1 int32; F2 int8 }
type GenS14 struct { F0 uint16; F1 int32 }
type GenS15 struct { F0 uint }
type GenS16 struct { F0 uint; F1 int16 }
type GenS17 struct { F0 bool }

func main() {
    v1 := []uint32{2147483648, 0}
    v2 := make(map[uint32]uint32)
    var v3 int32 = 0
    v4 := make(map[int]int)
    v5 := []string{"foo bar", "a", ("" + "foo bar"), "\n\t"}
    var v6 int32 = (v3 + (v3 / ((v3) | 1)))
    if len(v5) > 0 {
        v5[0] = "\n\t"
    }
    switch ("foo bar" + ("\n\t" + (((("hello" + (((("" + ("" + "")) + ("\n\t" + "foo bar")) + ((("hello" + ("\n\t" + "\n\t")) + "hello") + "")) + "a")) + "hello") + "") + "hello"))) {
    case "":
        _ = len(v4)
    }
    if false {
        fmt.Println(v3, v6)
        var v7 uint64 = 1
        v8 := GenS1{F0: 1}
        _ = v7
        _ = v8
    }
    if (2 >= 38) {
        switch 32767 {
        case 42:
            if (!(false && (("hello" < ("\n\t" + "\n\t")) || false))) {
                if true {
                    v6 -= 880
                    var v9 uint8 = 128
                    _ = v9
                }
                var v10 uint16 = 723
                _ = v10
            } else {
                v11 := GenS2{F0: 4294967295, F1: 65535, F2: 672}
                var v12 int16 = -801
                var v13 int32 = v6
                _ = v11
                _ = v12
                _ = v13
            }
        case 481:
            delete(v2, 1)
        default:
            v14 := GenS1{F0: 0}
            _ = v14
        }
    }
    v6 -= v3
    if (0 > 24) {
        v15 := GenS3{F0: 100}
        delete(v2, 65535)
        for fi16 := 0; fi16 < 15; fi16++ {
            v17 := GenS4{F0: 159, F1: 798, F2: -1}
            _ = len(v2)
            if false {
                switch "a" {
                case "a":
                    for fi18 := 0; fi18 < 12; fi18++ {
                        if false {
                            v6 = (-(v3))
                            v3 = 263
                        } else {
                            var v19 int8 = 0
                            _ = v19
                        }
                        _ = fi18
                    }
                }
                switch "hello" {
                case "":
                    v20 := v4[fi16]
                    _ = v20
                case "a":
                    var v21 float32 = (0.1 + 1e10)
                    _ = v21
                }
                var v22 int8 = -128
                _ = v22
            } else {
                v6 += -2147483648
                var v23 uint8 = 255
                _ = v23
            }
            _ = fi16
            _ = v17
        }
        _ = v15
    }
    if false {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        var v24 int32 = (-(v3))
        _ = v24
    }
    v6 += (v6 << ((349) & 63))
    if (false || true) {
        v25 := GenS3{F0: -94}
        var v26 uint = 220
        _ = v25
        _ = v26
    }
    if (true || ((float64((float32(0.1) * 0.0)) < 3.14159) || (!(32768 <= 367)))) {
        var v27 int8 = 100
        for fi28 := 0; fi28 < 13; fi28++ {
            for fi29 := 0; fi29 < 3; fi29++ {
                var v30 int64 = 9223372036854775807
                _ = fi29
                _ = v30
            }
            _ = fi28
        }
        _ = v27
    }
    if false {
        fmt.Println(v3)
        switch 1 {
        case 4294967295:
            for fi31 := 0; fi31 < 10; fi31++ {
                var v32 uint8 = 42
                if (2 >= 255) {
                    v33 := GenS4{F0: v32, F1: 42, F2: 0}
                    var v34 float32 = 1e10
                    v3 += -2147483648
                    _ = v33
                    _ = v34
                } else {
                    v3 -= (v6 / (((v3 << ((255) & 63))) | 1))
                    for fi35 := 0; fi35 < 10; fi35++ {
                        if (!(true && (false && ((v6 % (((^(v3))) | 1)) >= v6)))) {
                            if (-1 == -1) {
                                var v36 uint = 42
                                v32 = v32
                                v37 := GenS1{F0: fi31}
                                _ = v36
                                _ = v37
                            } else {
                                var v38 float64 = 1.5
                                _ = v38
                            }
                            var v39 int8 = 0
                            var v40 int = 4294967295
                            _ = v39
                            _ = v40
                        } else {
                            fmt.Println(v6)
                            switch "a" {
                            case "hello":
                                v6 -= (v6 % ((-2147483648) | 1))
                            }
                        }
                        v3 = v6
                        _ = fi35
                    }
                    var v41 uint64 = 378
                    _ = v41
                }
                v42 := GenS5{F0: 18446744073709551615, F1: true}
                _ = fi31
                _ = v32
                _ = v42
            }
        case 2147483647:
            switch 32767 {
            case 127:
                v1 = append(v1, 42)
            case 32768:
                var v43 uint8 = 0
                _ = v43
            }
        default:
            if true {
                v3 -= (^(v3))
            } else {
                for fi44 := 0; fi44 < 8; fi44++ {
                    for fi45 := 0; fi45 < 3; fi45++ {
                        v46 := GenS6{F0: 2147483648, F1: 198, F2: -128}
                        _ = fi45
                        _ = v46
                    }
                    v47 := GenS7{F0: 1, F1: (v3 | v3)}
                    _ = fi44
                    _ = v47
                }
                v3 = v3
            }
        }
    } else {
        if (32767 <= 0) {
            switch "foo bar" {
            case "a":
                var v48 uint32 = 95
                _ = v48
            case "hello":
                if true {
                    for fi49 := 0; fi49 < 14; fi49++ {
                        if (!false) {
                            var v50 uint16 = 65535
                            var v51 uint32 = 65535
                            v50 = (v50 + 1)
                            _ = v50
                            _ = v51
                        }
                        for fi52 := 0; fi52 < 6; fi52++ {
                            if true {
                                switch (fi49 >> ((0) & 63)) {
                                case -436:
                                    v3 += v3
                                case 0:
                                    var v53 uint32 = 909
                                    _ = v53
                                case -234:
                                    switch fi52 {
                                    case 32768:
                                        v6 = 42
                                    case 1:
                                        if (1 >= -1) {
                                            if false {
                                                if (0.1 > -2.25) {
                                                    v54 := v2[0]
                                                    v6 += v3
                                                    _ = v54
                                                } else {
                                                    v5 = append(v5, "\n\t")
                                                }
                                            }
                                            v6 = 2147483647
                                        }
                                    default:
                                        if false {
                                            var v55 string = "foo bar"
                                            v6 += (v6 % ((v6) | 1))
                                            var v56 int = 783
                                            if ((42 >= 87) || (1 > 174)) {
                                                v56 = fi52
                                                var v57 uint16 = 65535
                                                _ = v57
                                            }
                                            fmt.Println(v3, v55, v55)
                                            _ = v55
                                            _ = v56
                                        } else {
                                            var v58 uint = 42
                                            v3 = (v3 - -267)
                                            _ = v58
                                        }
                                    }
                                default:
                                    var v59 float64 = (float64(1.5) + (0.0 + 1e10))
                                    _ = v59
                                }
                            } else {
                                var v60 string = "foo bar"
                                _ = v60
                            }
                            _ = fi52
                        }
                        _ = fi49
                    }
                    switch -709 {
                    case 10:
                        for fi61 := 0; fi61 < 13; fi61++ {
                            v6 += v3
                            var v62 int32 = 0
                            _ = fi61
                            _ = v62
                        }
                    case 2147483647:
                        for fi63 := 0; fi63 < 19; fi63++ {
                            for fi64 := 0; fi64 < 18; fi64++ {
                                if true {
                                    var v65 float32 = (0.1 * 1.5)
                                    switch (fi63 & fi64) {
                                    case -1:
                                        switch "hello" {
                                        case "foo bar":
                                            var v66 string = ""
                                            _ = v66
                                        case "a":
                                            if true {
                                                if (true && (!true)) {
                                                    var v67 int32 = v6
                                                    v68 := GenS6{F0: 65535, F1: -975, F2: 127}
                                                    _ = v67
                                                    _ = v68
                                                } else {
                                                    v69 := GenS1{F0: (-(fi63))}
                                                    _ = v69
                                                }
                                                if true {
                                                    var v70 string = ""
                                                    if false {
                                                        var v71 bool = true
                                                        fmt.Println(v6)
                                                        _ = v71
                                                    }
                                                    _ = v70
                                                } else {
                                                    var v72 uint32 = 2147483647
                                                    v6 -= (v6 ^ (^(v3)))
                                                    fmt.Println(v72)
                                                    _ = v72
                                                }
                                                v73 := GenS3{F0: 100}
                                                _ = v73
                                            } else {
                                                delete(v4, (fi63 ^ -250))
                                                v3 += v3
                                            }
                                        default:
                                            v6 = (v3 | v3)
                                        }
                                    default:
                                        switch fi64 {
                                        case -840:
                                            v6 = v3
                                            v74 := GenS2{F0: 65535, F1: 810, F2: 65535}
                                            fmt.Println(fi63, v65, v6)
                                            v75 := GenS4{F0: 128, F1: 42, F2: -1}
                                            fmt.Println(fi64, v65, fi63)
                                            _ = v74
                                            _ = v75
                                        case 32767:
                                            if true {
                                                if false {
                                                    switch fi64 {
                                                    case 2147483647:
                                                        v76 := GenS8{F0: fi64}
                                                        _ = v76
                                                    case 227:
                                                        v77 := GenS4{F0: 42, F1: -9223372036854775808, F2: 155}
                                                        _ = v77
                                                    case 42:
                                                        v78 := v2[65535]
                                                        _ = v78
                                                    }
                                                    v65 = (v65 - ((v65 * 0.0) + v65))
                                                    var v79 uint = 4294967295
                                                    fmt.Println(fi64, fi64)
                                                    _ = v79
                                                }
                                            }
                                        case 42:
                                            fmt.Println(fi63, v3, fi63)
                                        default:
                                            v65 = 0.0
                                        }
                                    }
                                    _ = v65
                                } else {
                                    v6 = v6
                                    if false {
                                        if true {
                                            var v80 uint16 = 0
                                            v80 = v80
                                            v81 := v2[1]
                                            v82 := GenS9{F0: (^(v80)), F1: 492}
                                            v80 = 255
                                            _ = v80
                                            _ = v81
                                            _ = v82
                                        }
                                        if true {
                                            var v83 int = fi64
                                            if (127 < 678) {
                                                var v84 uint8 = 1
                                                v85 := GenS9{F0: 32767, F1: 42}
                                                _ = v84
                                                _ = v85
                                            } else {
                                                fmt.Println(v3)
                                                v3 = -1
                                                v6 = (v6 << ((311) & 63))
                                            }
                                            var v86 int32 = (v3 * (-(v6)))
                                            var v87 int = 4294967295
                                            _ = v83
                                            _ = v86
                                            _ = v87
                                        }
                                        if (!true) {
                                            var v88 uint16 = 32767
                                            if ((true && true) || (!true)) {
                                                var v89 uint = 42
                                                v90 := GenS10{F0: 32767, F1: "\n\t", F2: 18446744073709551615}
                                                _ = v89
                                                _ = v90
                                            } else {
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                v91 := GenS11{F0: -32768, F1: "foo bar"}
                                                _ = v91
                                            }
                                            _ = v88
                                        } else {
                                            if (!false) {
                                                v6 += (v6 - 42)
                                                fmt.Println(v3)
                                            }
                                            fmt.Println(v3)
                                        }
                                    } else {
                                        v92 := GenS12{F0: 0, F1: 1, F2: 0}
                                        _ = v92
                                    }
                                    v93 := GenS13{F0: -128, F1: -959, F2: 0}
                                    _ = v93
                                }
                                if true {
                                    if (!(127 > 42)) {
                                        v94 := GenS8{F0: fi63}
                                        _ = len(v1)
                                        switch (fi64 >> ((1) & 63)) {
                                        case -357:
                                            v95 := GenS6{F0: 4294967295, F1: -716, F2: int8(v3)}
                                            _ = v95
                                        case 127:
                                            var v96 int16 = 185
                                            _ = v96
                                        case 1:
                                            v97 := GenS11{F0: -989, F1: "foo bar"}
                                            _ = v97
                                        default:
                                            var v98 uint16 = 65535
                                            _ = v98
                                        }
                                        if false {
                                            fmt.Println(v6)
                                            var v99 float64 = 3.14159
                                            _ = v99
                                        } else {
                                            fmt.Println(fi63, fi63, fi64)
                                            var v100 float64 = 0.0
                                            v101 := v4[fi63]
                                            _ = v100
                                            _ = v101
                                        }
                                        var v102 float32 = 1.5
                                        var v103 int8 = 37
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        if (128 <= 255) {
                                            v104 := GenS14{F0: 2, F1: (^(v3))}
                                            _ = len(v1)
                                            _ = v104
                                        }
                                        _ = v102
                                        _ = v103
                                        _ = v94
                                    } else {
                                        v3 = v6
                                    }
                                } else {
                                    if true {
                                        v3 += 2147483647
                                        v105 := GenS7{F0: 2147483647, F1: (v3 / ((v6) | 1))}
                                        if true {
                                            v106 := GenS15{F0: 922}
                                            switch -750 {
                                            case 32767:
                                                if (255 != 963) {
                                                    v6 = (v6 - (v3 | (^(v6))))
                                                }
                                            }
                                            v3 = (v6 + 0)
                                            _ = v106
                                        }
                                        _ = v105
                                    }
                                    _ = len(v5)
                                }
                                _ = fi64
                            }
                            if (!(-9223372036854775808 >= -605)) {
                                var v107 float64 = 3.14159
                                _ = v107
                            }
                            _ = fi63
                        }
                    case 9223372036854775807:
                        for fi108 := 0; fi108 < 6; fi108++ {
                            fmt.Println(v3, fi108, fi108)
                            _ = fi108
                        }
                    default:
                        if (true && ((!(v3 <= 0)) || false)) {
                            fmt.Println(v6)
                            switch 9223372036854775807 {
                            case 196:
                                switch "foo bar" {
                                case "foo bar":
                                    for fi109 := 0; fi109 < 2; fi109++ {
                                        switch fi109 {
                                        case -1:
                                            v3 = -468
                                        case 9223372036854775807:
                                            fmt.Println(v6, v3, v6)
                                        }
                                        fmt.Println(fi109, v6, fi109)
                                        for fi110 := 0; fi110 < 4; fi110++ {
                                            fmt.Println(fi109, fi110, fi110)
                                            _ = fi110
                                        }
                                        _ = fi109
                                    }
                                case "":
                                    if (((!(true || (!true))) && true) || false) {
                                        fmt.Println(v6)
                                        v3 -= 1
                                    } else {
                                        var v111 float32 = 3.14159
                                        for fi112 := 0; fi112 < 8; fi112++ {
                                            v6 = 42
                                            var v113 uint32 = 1
                                            v3 = -807
                                            _ = fi112
                                            _ = v113
                                        }
                                        _ = v111
                                    }
                                case "\n\t":
                                    if (false || true) {
                                        v114 := GenS8{F0: 2147483647}
                                        _ = v114
                                    }
                                }
                            }
                            v115 := GenS7{F0: 2147483648, F1: (^(v6))}
                            _ = v115
                        } else {
                            v6 += (v3 / (((v3 % ((2147483647) | 1))) | 1))
                            v2[2147483647] = 0
                        }
                    }
                    fmt.Println(v6, v6)
                } else {
                    v6 -= v6
                    for fi116 := 0; fi116 < 19; fi116++ {
                        v3 = v3
                        v3 = 397
                        var v117 int32 = (v3 >> ((623) & 63))
                        _ = fi116
                        _ = v117
                    }
                }
            }
            switch 1 {
            case -1:
                v3 = (v3 - 765)
            case 32767:
                for fi118 := 0; fi118 < 20; fi118++ {
                    v3 = (v6 / (((v6 % ((v6) | 1))) | 1))
                    switch "hello" {
                    case "a":
                        v119 := GenS2{F0: 246, F1: 0, F2: 2}
                        _ = v119
                    case "hello":
                        v120 := GenS12{F0: 28, F1: -1, F2: 2}
                        _ = v120
                    case "":
                        v3 = v3
                    default:
                        if true {
                            var v121 int16 = 42
                            var v122 uint8 = 0
                            _ = v121
                            _ = v122
                        }
                    }
                    _ = fi118
                }
            default:
                if (v3 < (v3 % ((v6) | 1))) {
                    switch "a" {
                    case "":
                        for fi123 := 0; fi123 < 5; fi123++ {
                            v3 = (v3 ^ (v3 + (v6 | -1)))
                            if (false || true) {
                                for fi124 := 0; fi124 < 16; fi124++ {
                                    var v125 uint64 = 0
                                    _ = fi124
                                    _ = v125
                                }
                                for fi126 := 0; fi126 < 2; fi126++ {
                                    v5 = append(v5, "hello")
                                    switch ("foo bar" + "a") {
                                    case "hello":
                                        v3 -= (-(v6))
                                    }
                                    if false {
                                        var v127 int8 = 0
                                        v6 -= v6
                                        _ = v127
                                    } else {
                                        if ((0 == 4294967295) || (false && false)) {
                                            var v128 uint = 65535
                                            var v129 uint32 = 717
                                            _ = v128
                                            _ = v129
                                        }
                                        v6 = v3
                                        v130 := GenS6{F0: 2, F1: -1, F2: 1}
                                        _ = v130
                                    }
                                    var v131 int16 = -1
                                    v131 = -463
                                    _ = fi126
                                    _ = v131
                                }
                            }
                            v4[(fi123 * (fi123 & (fi123 + fi123)))] = (fi123 * fi123)
                            _ = fi123
                        }
                    }
                } else {
                    var v132 string = ""
                    var v133 uint32 = 1
                    for fi134 := 0; fi134 < 4; fi134++ {
                        switch -711 {
                        case -9223372036854775808:
                            switch "foo bar" {
                            case "foo bar":
                                var v135 int32 = v3
                                _ = v135
                            case "\n\t":
                                fmt.Println(v133, fi134, v3)
                            default:
                                v132 = ""
                            }
                        }
                        var v136 int8 = 26
                        _ = fi134
                        _ = v136
                    }
                    _ = v132
                    _ = v133
                }
            }
        } else {
            for fi137 := 0; fi137 < 6; fi137++ {
                v138 := GenS16{F0: 65535, F1: 1000}
                v139 := GenS14{F0: 65535, F1: 2147483647}
                _ = fi137
                _ = v138
                _ = v139
            }
            if true {
                var v140 string = "\n\t"
                _ = v140
            }
            var v141 uint32 = 42
            _ = v141
        }
        if (!true) {
            v142 := GenS17{F0: (true || false)}
            _ = v142
        }
    }
    v6 -= (v6 + (v3 % (((v6 << ((42) & 63))) | 1)))
    fmt.Println(v6)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
