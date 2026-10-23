package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 uint64 }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 uint16; F1 string; F2 int8 }
type GenS4 struct { F0 uint; F1 int64; F2 bool }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int8; F1 int16 }
type GenS7 struct { F0 int32; F1 string; F2 int64 }
type GenS8 struct { F0 uint64; F1 uint64 }
type GenS9 struct { F0 string; F1 bool }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 uint64; F1 string }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 bool; F1 string }
func genFn0(p0 string, p1 uint16, p2 uint32) uint8 {
    v5 := []int{487}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println(p2)
    if len(v5) > 0 {
        v5[0] = -2147483648
    }
    v5 = append(v5, 42)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 127
}
func genFn1(p0 uint, p1 int8, p2 uint64) int8 {
    var v6 uint32 = 2147483647
    v7 := GenS1{F0: 1, F1: 9223372036854775807, F2: 4294967295}
    var v8 uint32 = v6
    var v9 int16 = 353
    v10 := GenS2{F0: true, F1: v6}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return p1
}

func main() {
    var v1 int8 = -1
    v2 := make(map[string]int)
    v3 := []uint32{65535, 4294967295}
    var v4 float32 = (0.0 - 1e10)
    v4 -= 0.0
    v3 = append(v3, 659)
    switch int(v1) {
    case 127:
        if true {
            fmt.Println(v4, v1)
            fmt.Println(v4)
        } else {
            switch (("a" + "a") + "foo bar") {
            case "foo bar":
                if false {
                    if false {
                        switch 2147483647 {
                        case -307:
                            if false {
                                if false {
                                    switch "hello" {
                                    case "a":
                                        v4 = (-2.25 - float32(1e10))
                                    case "foo bar":
                                        v1 = v1
                                    }
                                    var v11 uint16 = 32767
                                    v3 = append(v3, 0)
                                    _ = v11
                                }
                                switch -9223372036854775808 {
                                case -212:
                                    switch (("foo bar" + "") + ("" + ("a" + ("a" + "hello")))) {
                                    case "foo bar":
                                        switch "hello" {
                                        case "a":
                                            if ((false || (!(4294967295 > 4294967295))) && false) {
                                                v12 := GenS2{F0: ((128 != 127) || (42 < 9223372036854775807)), F1: 641}
                                                fmt.Println(v4, v4)
                                                v1 -= 0
                                                _ = v12
                                            }
                                        default:
                                            v3 = append(v3, 1)
                                        }
                                    }
                                default:
                                    _ = len(v3)
                                }
                                var v13 uint8 = 2
                                _ = v13
                            }
                        }
                        var v14 int64 = 847
                        _ = v14
                    } else {
                        switch "\n\t" {
                        case "a":
                            for fi15 := 0; fi15 < 16; fi15++ {
                                v4 -= v4
                                for fi16 := 0; fi16 < 8; fi16++ {
                                    delete(v2, "a")
                                    _ = fi16
                                }
                                _ = fi15
                            }
                        case "\n\t":
                            v17 := GenS2{F0: (4294967295 <= 42), F1: 2}
                            _ = v17
                        default:
                            v4 = v4
                        }
                        v4 = v4
                        if len(v3) > 0 {
                            v3[0] = 645
                        }
                    }
                    fmt.Println(v1, v1)
                }
            case "\n\t":
                if (127 <= 69) {
                    v4 = v4
                    switch 560 {
                    case 42:
                        switch 42 {
                        case 749:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        case 32767:
                            fmt.Println(v1)
                        case 116:
                            var v18 float32 = (-2.25 - v4)
                            _ = v18
                        }
                    case 4294967295:
                        v4 += ((((v4 / -2.25) / -2.25) - v4) - 1e10)
                    case 905:
                        fmt.Println(v4)
                    }
                    fmt.Println(v4)
                }
            default:
                if (1 <= 2) {
                    v1 = (^(v1))
                    var v19 int64 = 187
                    if false {
                        fmt.Println(v19, v4)
                        v1 += v1
                        if true {
                            for fi20 := 0; fi20 < 14; fi20++ {
                                var v21 uint8 = 0
                                var v22 uint8 = (v21 / ((v21) | 1))
                                _ = fi20
                                _ = v21
                                _ = v22
                            }
                            for fi23 := 0; fi23 < 18; fi23++ {
                                for fi24 := 0; fi24 < 20; fi24++ {
                                    if true {
                                        fmt.Println(fi24, v1)
                                    } else {
                                        if false {
                                            var v25 uint64 = 340
                                            if true {
                                                var v26 int32 = 718
                                                fmt.Println(v25)
                                                var v27 int8 = 1
                                                _ = len(v3)
                                                v25 -= 0
                                                v1 -= (v1 ^ (-(v1)))
                                                v4 += (v4 + v4)
                                                fmt.Println(v26, fi23, v27)
                                                var v28 bool = (genFn1(42, (-(v27)), 869) < 49)
                                                _ = v26
                                                _ = v27
                                                _ = v28
                                            }
                                            v29 := GenS3{F0: 65535, F1: ("" + "hello"), F2: v1}
                                            var v30 uint64 = (v25 << ((1) & 63))
                                            _ = v25
                                            _ = v29
                                            _ = v30
                                        }
                                        v3 = append(v3, 42)
                                    }
                                    v31 := GenS4{F0: 964, F1: (v19 / (((v19 >> ((65535) & 63))) | 1)), F2: (!(2 == 686))}
                                    switch fi23 {
                                    case 9223372036854775807:
                                        v4 += 1e10
                                    case 42:
                                        v32 := GenS2{F0: true, F1: 2}
                                        _ = v32
                                    }
                                    _ = fi24
                                    _ = v31
                                }
                                v19 += (v19 ^ (^(v19)))
                                if ((0.1 - 1e10) <= 1.5) {
                                    switch (((("" + "") + "") + "a") + ("a" + "hello")) {
                                    case "a":
                                        if (2 < 0) {
                                            fmt.Println(v4, v4, v19)
                                        }
                                    default:
                                        var v33 uint16 = 32767
                                        _ = v33
                                    }
                                    var v34 bool = (449 != -1)
                                    if (fi23 < (fi23 >> ((65535) & 63))) {
                                        for fi35 := 0; fi35 < 7; fi35++ {
                                            var v36 uint8 = 255
                                            _ = fi35
                                            _ = v36
                                        }
                                    }
                                    _ = v34
                                }
                                _ = fi23
                            }
                            if ("a" <= ((("a" + "hello") + "hello") + "hello")) {
                                if true {
                                    if false {
                                        switch 32767 {
                                        case 128:
                                            v3 = append(v3, 2147483647)
                                        case 32767:
                                            if false {
                                                if false {
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    v19 += v19
                                                }
                                            } else {
                                                v2["\n\t"] = 32768
                                                var v37 string = ""
                                                var v38 int16 = -1
                                                _ = v37
                                                _ = v38
                                            }
                                        }
                                        v39 := GenS4{F0: 630, F1: 0, F2: (v4 > v4)}
                                        _ = v39
                                    }
                                    if true {
                                        for fi40 := 0; fi40 < 7; fi40++ {
                                            var v41 int16 = 0
                                            fmt.Println(v19)
                                            _ = fi40
                                            _ = v41
                                        }
                                    } else {
                                        var v42 uint8 = 42
                                        _ = v42
                                    }
                                } else {
                                    var v43 int64 = v19
                                    _ = v43
                                }
                                v44 := GenS5{F0: int8(v19)}
                                _ = v44
                            }
                        } else {
                            _ = len(v3)
                            if true {
                                if false {
                                    v45 := GenS6{F0: v1, F1: -32768}
                                    _ = v45
                                } else {
                                    switch "\n\t" {
                                    case "":
                                        v46 := GenS4{F0: 893, F1: (v19 % (((v19 ^ -9223372036854775808)) | 1)), F2: (false || (!(false || false)))}
                                        _ = v46
                                    }
                                    var v47 uint64 = 18446744073709551615
                                    for fi48 := 0; fi48 < 14; fi48++ {
                                        v49 := v2["\n\t"]
                                        _ = fi48
                                        _ = v49
                                    }
                                    _ = v47
                                }
                                fmt.Println(v1)
                            } else {
                                fmt.Println(v4, v4, v1)
                            }
                            if (!false) {
                                v50 := GenS2{F0: true, F1: 2147483647}
                                for fi51 := 0; fi51 < 6; fi51++ {
                                    if true {
                                        v1 = v1
                                    }
                                    _ = fi51
                                }
                                _ = v50
                            } else {
                                switch "foo bar" {
                                case "foo bar":
                                    fmt.Println(v19, v19, v4)
                                default:
                                    if (true || ((true || true) && (!false))) {
                                        v19 = (^(v19))
                                        for fi52 := 0; fi52 < 8; fi52++ {
                                            if (1 == -728) {
                                                var v53 uint32 = 1
                                                _ = v53
                                            }
                                            for fi54 := 0; fi54 < 1; fi54++ {
                                                fmt.Println(v1)
                                                _ = fi54
                                            }
                                            var v55 uint = 0
                                            _ = fi52
                                            _ = v55
                                        }
                                        v4 += v4
                                    }
                                }
                            }
                        }
                    } else {
                        for fi56 := 0; fi56 < 8; fi56++ {
                            fmt.Println(v1)
                            if false {
                                if (2 == 665) {
                                    fmt.Println(v1)
                                    v57 := GenS5{F0: genFn1(18446744073709551615, (v1 % (((-(v1))) | 1)), 588)}
                                    _ = v57
                                }
                                var v58 int8 = v1
                                _ = len(v3)
                                _ = v58
                            }
                            var v59 uint32 = 4294967295
                            _ = fi56
                            _ = v59
                        }
                    }
                    _ = v19
                }
            }
            for fi60 := 0; fi60 < 14; fi60++ {
                fmt.Println(v1)
                var v61 uint32 = 199
                for fi62 := 0; fi62 < 17; fi62++ {
                    v61 += (^(v61))
                    switch (fi60 - 32768) {
                    case 4294967295:
                        v63 := GenS7{F0: 42, F1: "hello", F2: 1}
                        fmt.Println(fi62, fi60, fi60)
                        _ = v63
                    case 1:
                        var v64 uint8 = 92
                        v61 = (v61 | (^(v61)))
                        _ = v64
                    case 9223372036854775807:
                        v65 := GenS4{F0: 255, F1: 1, F2: false}
                        _ = v65
                    default:
                        v1 += v1
                    }
                    v2["hello"] = 137
                    _ = fi62
                }
                _ = fi60
                _ = v61
            }
        }
    }
    v66 := v2[("hello" + (("a" + "foo bar") + "hello"))]
    if (0 == 0) {
        var v67 float64 = ((1.5 + (float64(v4) * 3.14159)) / (float64(v4) - 3.14159))
        for fi68 := 0; fi68 < 10; fi68++ {
            if (42 <= 32767) {
                var v69 uint16 = 255
                if true {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    if ((v69 != v69) && (v69 < (v69 << ((138) & 63)))) {
                        if false {
                            fmt.Println(fi68)
                            v67 = float64(1.5)
                            v4 = 0.0
                        }
                    }
                    v1 = 42
                } else {
                    if false {
                        for fi70 := 0; fi70 < 6; fi70++ {
                            var v71 uint16 = (v69 ^ (v69 * 244))
                            fmt.Println(v4)
                            _ = fi70
                            _ = v71
                        }
                        v69 = v69
                        v72 := GenS3{F0: v69, F1: "a", F2: v1}
                        _ = v72
                    }
                    if true {
                        if len(v3) > 0 {
                            v3[0] = 42
                        }
                        var v73 uint16 = v69
                        fmt.Println(fi68)
                        _ = v73
                    } else {
                        var v74 int32 = -747
                        switch "\n\t" {
                        case "":
                            switch v66 {
                            case 627:
                                v75 := GenS6{F0: (v1 & (v1 % (((v1 ^ (v1 * (v1 & (v1 + 42))))) | 1))), F1: 42}
                                _ = v75
                            case 127:
                                var v76 uint8 = 2
                                _ = v76
                            case 2147483647:
                                var v77 uint64 = 4294967295
                                _ = v77
                            }
                        case "\n\t":
                            if false {
                                fmt.Println(fi68, v67)
                                for fi78 := 0; fi78 < 8; fi78++ {
                                    var v79 bool = (!(!((3.14159 - v67) < 3.14159)))
                                    if v79 {
                                        if ((((true && (v79 && v79)) && v79) || (-359 == -1)) && (-445 < 750)) {
                                            switch "foo bar" {
                                            case "hello":
                                                var v80 uint8 = 42
                                                _ = v80
                                            default:
                                                var v81 uint64 = 295
                                                _ = v81
                                            }
                                        } else {
                                            v74 += int32(v1)
                                            switch "" {
                                            case "":
                                                if false {
                                                    var v82 int64 = 0
                                                    var v83 string = (("\n\t" + "\n\t") + "a")
                                                    _ = v82
                                                    _ = v83
                                                } else {
                                                    var v84 int = (fi78 ^ (^(fi78)))
                                                    v85 := v2[("hello" + "foo bar")]
                                                    _ = v84
                                                    _ = v85
                                                }
                                            case "a":
                                                if len(v3) > 0 {
                                                    v3[0] = 65535
                                                }
                                            default:
                                                var v86 int = (fi78 >> ((25) & 63))
                                                _ = v86
                                            }
                                        }
                                        switch "hello" {
                                        case "\n\t":
                                            if len(v3) > 0 {
                                                v3[0] = 42
                                            }
                                        default:
                                            v3 = append(v3, 65535)
                                        }
                                        v69 = v69
                                    } else {
                                        var v87 uint = 0
                                        v79 = ((-237 >= -1) || (v87 < (v87 << ((v87) & 63))))
                                        _ = v87
                                    }
                                    if len(v3) > 0 {
                                        v3[0] = 65535
                                    }
                                    _ = fi78
                                    _ = v79
                                }
                            }
                        default:
                            for fi88 := 0; fi88 < 6; fi88++ {
                                if false {
                                    if (!(255 != 127)) {
                                        var v89 int = 1
                                        v74 = -130
                                        v90 := GenS6{F0: 0, F1: 628}
                                        v69 = v69
                                        v91 := v2["foo bar"]
                                        _ = v89
                                        _ = v90
                                        _ = v91
                                    }
                                    v74 = -2147483648
                                }
                                if false {
                                    v4 += ((((v4 - 1e10) / v4) / v4) - (float32(v67) - v4))
                                }
                                _ = fi88
                            }
                        }
                        v92 := GenS3{F0: (v69 ^ (v69 / (((v69 * v69)) | 1))), F1: "", F2: v1}
                        _ = v74
                        _ = v92
                    }
                    for fi93 := 0; fi93 < 20; fi93++ {
                        if len(v3) > 0 {
                            v3[0] = 42
                        }
                        switch (fi68 % (((v66 & (v66 >> ((0) & 63)))) | 1)) {
                        case 100:
                            if (v4 <= v4) {
                                var v94 int8 = -1
                                v94 += v1
                                _ = v94
                            } else {
                                var v95 string = "hello"
                                fmt.Println(v66)
                                _ = v95
                            }
                            v96 := GenS3{F0: v69, F1: ((("hello" + "a") + ("\n\t" + "a")) + "hello"), F2: (v1 & v1)}
                            _ = v96
                        }
                        _ = fi93
                    }
                }
                _ = v69
            } else {
                fmt.Println(v66, v66, v4)
                v97 := v2["foo bar"]
                _ = v97
            }
            switch fi68 {
            case 85:
                if false {
                    if (v66 > (fi68 - 470)) {
                        v3 = append(v3, 42)
                        fmt.Println(fi68)
                        var v98 int64 = 0
                        _ = v98
                    } else {
                        if ((0 <= 42) || ((^(v66)) != 0)) {
                            v66 += (v66 / (((v66 | (fi68 - (v66 / ((-2147483648) | 1))))) | 1))
                            v99 := v2[""]
                            switch "a" {
                            case "foo bar":
                                for fi100 := 0; fi100 < 18; fi100++ {
                                    _ = len(v2)
                                    var v101 uint8 = 127
                                    _ = fi100
                                    _ = v101
                                }
                            case "hello":
                                var v102 int32 = -191
                                _ = v102
                            case "a":
                                _ = len(v3)
                            }
                            _ = v99
                        } else {
                            for fi103 := 0; fi103 < 18; fi103++ {
                                v104 := GenS8{F0: 285, F1: 0}
                                var v105 float64 = v67
                                _ = fi103
                                _ = v104
                                _ = v105
                            }
                            var v106 int = -768
                            _ = v106
                        }
                    }
                    for fi107 := 0; fi107 < 6; fi107++ {
                        v2["a"] = v66
                        switch "foo bar" {
                        case "":
                            var v108 float32 = 1e10
                            v3 = append(v3, 300)
                            fmt.Println(v66, v4, v67)
                            _ = v108
                        case "a":
                            delete(v2, ("" + (("foo bar" + ("\n\t" + "\n\t")) + "hello")))
                        case "\n\t":
                            if (!false) {
                                v4 -= v4
                            } else {
                                switch v66 {
                                case 1:
                                    _ = len(v2)
                                }
                            }
                        default:
                            fmt.Println(v1)
                        }
                        delete(v2, "\n\t")
                        _ = fi107
                    }
                    if false {
                        for fi109 := 0; fi109 < 9; fi109++ {
                            var v110 float64 = (v67 + 1.5)
                            var v111 string = "foo bar"
                            v112 := GenS3{F0: 671, F1: "a", F2: v1}
                            v113 := GenS9{F0: (v111 + "a"), F1: (!(((v4 * 0.0) + v4) != v4))}
                            _ = fi109
                            _ = v110
                            _ = v111
                            _ = v112
                            _ = v113
                        }
                    } else {
                        switch v66 {
                        case 127:
                            v4 = v4
                        case 2147483647:
                            if (32767 < 628) {
                                v114 := GenS10{F0: 82}
                                _ = v114
                            }
                        default:
                            var v115 uint = 750
                            _ = v115
                        }
                        v116 := GenS3{F0: 0, F1: "hello", F2: v1}
                        _ = v116
                    }
                }
            }
            _ = fi68
        }
        _ = v67
    } else {
        var v117 string = ""
        for fi118 := 0; fi118 < 10; fi118++ {
            for fi119 := 0; fi119 < 16; fi119++ {
                var v120 uint8 = 127
                _ = fi119
                _ = v120
            }
            fmt.Println(v117, v1, v66)
            _ = fi118
        }
        _ = v117
    }
    v66 += (-(v66))
    var v121 int64 = 9223372036854775807
    for fi122 := 0; fi122 < 8; fi122++ {
        v123 := v2["hello"]
        if true {
            for fi124 := 0; fi124 < 3; fi124++ {
                v4 -= (float32(3.14159) * ((v4 * (v4 - (v4 / v4))) * v4))
                v125 := GenS2{F0: true, F1: 2147483647}
                _ = fi124
                _ = v125
            }
        } else {
            if false {
                if false {
                    for fi126 := 0; fi126 < 4; fi126++ {
                        if (!(("" != "foo bar") || (((v66 >> ((294) & 63)) > fi122) || (false || (2 >= 255))))) {
                            v1 = v1
                            var v127 int64 = 9223372036854775807
                            _ = v127
                        }
                        _ = fi126
                    }
                    v128 := GenS11{F0: 657, F1: "hello"}
                    _ = v128
                }
            }
        }
        _ = fi122
        _ = v123
    }
    if (true && true) {
        for fi129 := 0; fi129 < 19; fi129++ {
            v4 = (v4 - -2.25)
            for fi130 := 0; fi130 < 14; fi130++ {
                v131 := GenS12{F0: 855}
                fmt.Println(v66, v4, v4)
                v66 = -143
                _ = fi130
                _ = v131
            }
            _ = fi129
        }
    }
    v132 := GenS3{F0: 0, F1: (("hello" + "a") + ""), F2: v1}
    for fi133 := 0; fi133 < 14; fi133++ {
        fmt.Println(v4, fi133, v4)
        v134 := GenS13{F0: true, F1: "\n\t"}
        _ = fi133
        _ = v134
    }
    for fi135 := 0; fi135 < 2; fi135++ {
        for fi136 := 0; fi136 < 19; fi136++ {
            var v137 int32 = 42
            var v138 uint32 = 65535
            _ = fi136
            _ = v137
            _ = v138
        }
        if (605 == 251) {
            v4 -= v4
        }
        var v139 uint16 = 42
        _ = fi135
        _ = v139
    }
    v140 := GenS1{F0: 1, F1: v121, F2: 1}
    v141 := GenS3{F0: 580, F1: "hello", F2: (v1 | (v1 / (((v1 << ((457) & 63))) | 1)))}
    v121 -= v121
    v4 -= v4
    switch 1 {
    case -793:
        if ((42 > 42) || true) {
            v121 += (v121 / ((v121) | 1))
        }
    }
    v1 -= (v1 + v1)
    fmt.Println("v1:", v1)
    fmt.Println("v121:", v121)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v121
    _ = v132
    _ = v140
    _ = v141
    _ = v2
    _ = v3
    _ = v4
    _ = v66
}
