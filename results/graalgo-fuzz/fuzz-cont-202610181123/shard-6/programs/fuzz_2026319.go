package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32; F1 int8; F2 uint }
type GenS4 struct { F0 uint32; F1 uint8 }
type GenS5 struct { F0 int32; F1 int32 }
type GenS6 struct { F0 string; F1 uint8; F2 uint64 }
type GenS7 struct { F0 int32; F1 string }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 uint32; F1 uint16 }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 uint16; F1 int8 }
type GenS14 struct { F0 uint32 }
func genFn0(p0 int32, p1 int32, p2 int16) int16 {
    p1 = (p0 - 1)
    v7 := GenS1{F0: 451, F1: (p2 ^ (p2 * p2)), F2: 2147483648}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p2
}
func genFn1(p0 int32) string {
    fmt.Println(p0)
    if false {
        for fi8 := 0; fi8 < 16; fi8++ {
            var v9 int = (fi8 * (fi8 >> ((255) & 63)))
            for fi10 := 0; fi10 < 11; fi10++ {
                if (361 != 2) {
                    switch fi8 {
                    case -156:
                        v11 := []int{v9, v9}
                        _ = v11
                    case 0:
                        v12 := GenS2{F0: (p0 % (((^(p0))) | 1))}
                        _ = v12
                    default:
                        v13 := GenS3{F0: (p0 % ((p0) | 1)), F1: -128, F2: 2}
                        v14 := []int{int(p0), (v9 & (fi10 + -9223372036854775808))}
                        v15 := GenS4{F0: 984, F1: 208}
                        var v16 int32 = (p0 * (p0 | (p0 & p0)))
                        _ = v13
                        _ = v14
                        _ = v15
                        _ = v16
                    }
                } else {
                    v17 := GenS1{F0: 87, F1: -160, F2: 2147483647}
                    if false {
                        if (!(1 < 1)) {
                            p0 = -2147483648
                            var v18 uint16 = 320
                            _ = v18
                        } else {
                            v19 := GenS2{F0: -680}
                            switch "\n\t" {
                            case "hello":
                                var v20 int8 = 21
                                _ = v20
                            }
                            _ = v19
                        }
                    }
                    switch "a" {
                    case "a":
                        v9 = fi8
                    case "foo bar":
                        switch "\n\t" {
                        case "foo bar":
                            v21 := GenS3{F0: (p0 >> ((2) & 63)), F1: 90, F2: 255}
                            _ = v21
                        case "":
                            switch "\n\t" {
                            case "":
                                var v22 bool = (42 == 22)
                                _ = v22
                            case "hello":
                                switch -1 {
                                case 2147483647:
                                    fmt.Println(p0)
                                default:
                                    v23 := GenS5{F0: p0, F1: 2147483647}
                                    _ = v23
                                }
                            default:
                                v24 := []int{(fi10 % ((2147483647) | 1))}
                                _ = v24
                            }
                        default:
                            var v25 float32 = 3.14159
                            _ = v25
                        }
                    case "":
                        p0 = p0
                    default:
                        fmt.Println(p0)
                    }
                    var v26 int = fi10
                    _ = v17
                    _ = v26
                }
                switch (("" + "\n\t") + "hello") {
                case "\n\t":
                    switch (("a" + ("" + "")) + ("foo bar" + "\n\t")) {
                    case "\n\t":
                        v9 += 32768
                    default:
                        v27 := GenS2{F0: p0}
                        _ = v27
                    }
                case "a":
                    v28 := []int{(fi8 ^ (-(v9)))}
                    _ = v28
                }
                fmt.Println(fi8)
                _ = fi10
            }
            for fi29 := 0; fi29 < 5; fi29++ {
                v30 := GenS6{F0: ("foo bar" + ("foo bar" + "hello")), F1: 167, F2: 42}
                var v31 uint = 1
                if (-962 == 1) {
                    p0 = p0
                    fmt.Println(v9, fi29)
                    var v32 int64 = -113
                    _ = v32
                } else {
                    fmt.Println(p0, p0)
                    var v33 uint = (v31 * v31)
                    var v34 int = -343
                    _ = v33
                    _ = v34
                }
                var v35 int64 = -495
                _ = fi29
                _ = v30
                _ = v31
                _ = v35
            }
            _ = fi8
            _ = v9
        }
        if false {
            fmt.Println(p0)
        } else {
            fmt.Println(p0)
            v36 := []int{}
            _ = v36
        }
        p0 = 1
    } else {
        if true {
            if false {
                fmt.Println(p0)
                if true {
                    var v37 int32 = (p0 >> ((18446744073709551615) & 63))
                    if (((!(722 <= 42)) || ((1 >= 2) && true)) || (0 <= 68)) {
                        switch "a" {
                        case "\n\t":
                            v38 := GenS5{F0: (p0 % ((p0) | 1)), F1: p0}
                            _ = v38
                        case "hello":
                            v37 -= -702
                        case "":
                            for fi39 := 0; fi39 < 8; fi39++ {
                                switch 127 {
                                case 32767:
                                    v37 += p0
                                case 127:
                                    v40 := []int{32768}
                                    _ = v40
                                default:
                                    if false {
                                        if true {
                                            v41 := []int{fi39, (fi39 << ((0) & 63)), (-(fi39))}
                                            for fi42 := 0; fi42 < 7; fi42++ {
                                                v43 := GenS2{F0: (v37 - v37)}
                                                v37 += 42
                                                if true {
                                                    switch fi42 {
                                                    case 908:
                                                        v37 = -1
                                                        if true {
                                                            var v44 int8 = -1
                                                            _ = v44
                                                        } else {
                                                            if len(v41) > 0 {
                                                                v41[0] = 32768
                                                            }
                                                        }
                                                        var v45 string = "a"
                                                        _ = v45
                                                    default:
                                                        var v46 int = 127
                                                        _ = v46
                                                    }
                                                    switch "" {
                                                    case "hello":
                                                        fmt.Println(fi42, fi42, p0)
                                                    case "a":
                                                        switch ("foo bar" + ("\n\t" + "")) {
                                                        case "hello":
                                                            fmt.Println(fi39, v37, fi39)
                                                        case "":
                                                            if len(v41) > 0 {
                                                                v41[0] = 32767
                                                            }
                                                        }
                                                    default:
                                                        var v47 string = "\n\t"
                                                        _ = v47
                                                    }
                                                    v37 = v37
                                                }
                                                if (("hello" + "a") <= "a") {
                                                    switch "" {
                                                    case "hello":
                                                        if len(v41) >= 2 {
                                                            v41 = v41[1:len(v41)]
                                                        }
                                                    case "\n\t":
                                                        p0 += p0
                                                        p0 = (p0 << ((18446744073709551615) & 63))
                                                    }
                                                    v37 = (p0 * v37)
                                                    if false {
                                                        v37 = (p0 + p0)
                                                        if len(v41) > 0 {
                                                            v41[0] = -534
                                                        }
                                                        if false {
                                                            if false {
                                                                p0 -= p0
                                                                v41 = append(v41, int(p0))
                                                                var v48 uint64 = 1
                                                                _ = v48
                                                            }
                                                            v41 = append(v41, 32767)
                                                            v37 += -999
                                                            if false {
                                                                v37 += -2147483648
                                                                if len(v41) > 0 {
                                                                    v41[0] = (fi42 & 871)
                                                                }
                                                                if (32767 < fi42) {
                                                                    fmt.Println(v37, fi42)
                                                                }
                                                            }
                                                        }
                                                    }
                                                } else {
                                                    p0 = (p0 >> ((2) & 63))
                                                }
                                                var v49 uint16 = 1
                                                _ = fi42
                                                _ = v43
                                                _ = v49
                                            }
                                            _ = v41
                                        }
                                        p0 += p0
                                    }
                                }
                                v50 := []int{fi39, (fi39 - (fi39 & fi39)), fi39, 32768}
                                _ = fi39
                                _ = v50
                            }
                        }
                        var v51 uint16 = 2
                        v52 := GenS6{F0: ("\n\t" + "a"), F1: 42, F2: 42}
                        _ = v51
                        _ = v52
                    }
                    _ = v37
                } else {
                    v53 := []int{-585}
                    _ = len(v53)
                    for fi54 := 0; fi54 < 11; fi54++ {
                        if true {
                            _ = len(v53)
                            fmt.Println(fi54, p0)
                        } else {
                            p0 = -2147483648
                        }
                        if (true && (-2.25 < float32(3.14159))) {
                            switch fi54 {
                            case -9223372036854775808:
                                fmt.Println(fi54, p0)
                            }
                            p0 += (p0 << ((667) & 63))
                            switch (("hello" + (("\n\t" + "") + "a")) + "a") {
                            case "\n\t":
                                var v55 uint64 = 26
                                _ = v55
                            case "a":
                                fmt.Println(fi54, fi54)
                            default:
                                var v56 int8 = 100
                                _ = v56
                            }
                        }
                        p0 = p0
                        _ = fi54
                    }
                    _ = v53
                }
            } else {
                v57 := GenS2{F0: -519}
                for fi58 := 0; fi58 < 7; fi58++ {
                    p0 += (^(p0))
                    if true {
                        if true {
                            var v59 uint64 = 4294967295
                            _ = v59
                        }
                    }
                    _ = fi58
                }
                switch 105 {
                case 4294967295:
                    var v60 int8 = 100
                    _ = v60
                case 9223372036854775807:
                    v61 := []int{4294967295}
                    _ = v61
                case 32768:
                    for fi62 := 0; fi62 < 6; fi62++ {
                        if (32767 <= -32768) {
                            var v63 uint32 = 65535
                            v64 := []int{}
                            p0 += (p0 & p0)
                            _ = v63
                            _ = v64
                        } else {
                            switch (fi62 - (fi62 - fi62)) {
                            case 32767:
                                p0 = (p0 % ((p0) | 1))
                            case 128:
                                v65 := []int{(fi62 ^ fi62), fi62, -2147483648, (fi62 | 0)}
                                _ = v65
                            default:
                                var v66 int8 = 127
                                _ = v66
                            }
                            fmt.Println(fi62, fi62)
                            switch (fi62 & fi62) {
                            case 128:
                                fmt.Println(p0, fi62)
                            case -573:
                                var v67 uint16 = 32767
                                _ = v67
                            }
                        }
                        p0 -= (p0 + p0)
                        var v68 uint = 294
                        _ = fi62
                        _ = v68
                    }
                default:
                    fmt.Println(p0)
                }
                _ = v57
            }
            if ((("hello" != "a") && false) && (!true)) {
                v69 := []int{128, 1, 32768, -9223372036854775808}
                p0 -= p0
                switch 127 {
                case 0:
                    p0 += (p0 >> ((0) & 63))
                case 1:
                    p0 = p0
                }
                _ = v69
            }
            v70 := GenS7{F0: -1, F1: ""}
            _ = v70
        } else {
            for fi71 := 0; fi71 < 18; fi71++ {
                p0 = (p0 & (^(p0)))
                _ = fi71
            }
            var v72 int64 = -866
            _ = v72
        }
        for fi73 := 0; fi73 < 14; fi73++ {
            switch (-(fi73)) {
            case -2147483648:
                v74 := []int{(fi73 >> ((18446744073709551615) & 63)), 32768, 42}
                _ = v74
            }
            p0 -= p0
            var v75 uint = 321
            _ = fi73
            _ = v75
        }
        var v76 float32 = 1e10
        _ = v76
    }
    var v77 int8 = -121
    _ = p0
    _ = v77
    return ("" + "hello")
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint64 = 2
    var v3 uint16 = 976
    var v4 uint32 = 861
    var v5 uint8 = 42
    var v6 float64 = 1.5
    var v78 float64 = v6
    var v79 float32 = -2.25
    switch 4294967295 {
    case 9223372036854775807:
        if true {
            var v80 int8 = 0
            switch 9223372036854775807 {
            case 125:
                v81 := GenS2{F0: int32(v5)}
                _ = v81
            case 127:
                var v82 uint16 = (v3 - 0)
                _ = v82
            default:
                for fi83 := 0; fi83 < 18; fi83++ {
                    for fi84 := 0; fi84 < 2; fi84++ {
                        var v85 int32 = int32(v5)
                        _ = fi84
                        _ = v85
                    }
                    var v86 float32 = (((1.5 * 1.5) + (v79 + float32(float64((float32((v78 * v6)) * 0.1))))) + v79)
                    _ = fi83
                    _ = v86
                }
            }
            var v87 int16 = 1
            _ = v80
            _ = v87
        } else {
            var v88 int32 = -1
            if true {
                var v89 int16 = -32768
                v90 := GenS4{F0: (^(v4)), F1: (v5 ^ (^(v5)))}
                v91 := GenS4{F0: v4, F1: v5}
                _ = v89
                _ = v90
                _ = v91
            } else {
                v79 = v79
            }
            _ = v88
        }
    default:
        fmt.Println(v5, v3, v79)
    }
    v1[("a" + ("foo bar" + "a"))] = v4
    for fi92 := 0; fi92 < 20; fi92++ {
        var v93 float64 = float64(0.1)
        _ = fi92
        _ = v93
    }
    v94 := GenS4{F0: 948, F1: (v5 / ((v5) | 1))}
    var v95 uint16 = v3
    v96 := []int{32767, 1}
    switch 32767 {
    case 42:
        v78 = v6
    case 4294967295:
        for fi97 := 0; fi97 < 20; fi97++ {
            if false {
                for fi98 := 0; fi98 < 17; fi98++ {
                    v3 -= 246
                    var v99 float64 = (((float64(v79) - (0.1 / float64(1e10))) + 0.0) / v78)
                    var v100 uint16 = 823
                    _ = fi98
                    _ = v100
                    _ = v99
                }
            }
            _ = fi97
        }
    case 32767:
        fmt.Println(v79, v95)
    default:
        v101 := GenS7{F0: 2147483647, F1: ""}
        _ = v101
    }
    if false {
        for fi102 := 0; fi102 < 12; fi102++ {
            v4 -= 747
            var v103 uint = 0
            _ = fi102
            _ = v103
        }
        switch 9223372036854775807 {
        case 2147483647:
            fmt.Println(v4, v5)
        default:
            if ((((((("hello" + ("foo bar" + "hello")) + "hello") + "hello") + "foo bar") + "\n\t") + "\n\t") >= (("\n\t" + (("\n\t" + ((("hello" + (("\n\t" + "a") + "a")) + "foo bar") + "a")) + "hello")) + "hello")) {
                v104 := GenS8{F0: ("foo bar" + "")}
                switch 127 {
                case -1:
                    delete(v1, ("" + ("hello" + ("a" + ""))))
                case 4294967295:
                    v105 := GenS9{F0: 18446744073709551615}
                    _ = v105
                }
                _ = v104
            } else {
                var v106 int64 = -292
                _ = v106
            }
        }
        var v107 int8 = -16
        _ = v107
    } else {
        var v108 uint = 96
        if false {
            for fi109 := 0; fi109 < 20; fi109++ {
                v4 -= 2
                _ = fi109
            }
            if (2 <= 518) {
                for fi110 := 0; fi110 < 3; fi110++ {
                    switch ("\n\t" + "foo bar") {
                    case "foo bar":
                        switch ("" + "hello") {
                        case "foo bar":
                            _ = len(v96)
                        default:
                            if (false || true) {
                                fmt.Println(v4)
                            } else {
                                switch ("a" + "\n\t") {
                                case "a":
                                    for fi111 := 0; fi111 < 19; fi111++ {
                                        var v112 int32 = 0
                                        _ = fi111
                                        _ = v112
                                    }
                                case "hello":
                                    v108 = v108
                                default:
                                    if true {
                                        v113 := GenS10{F0: true}
                                        _ = v113
                                    } else {
                                        var v114 int16 = 42
                                        _ = v114
                                    }
                                }
                                switch fi110 {
                                case 2147483647:
                                    if false {
                                        if false {
                                            if (799 <= (v2 << (((v108 ^ (v108 % ((v108) | 1)))) & 63))) {
                                                v5 -= v5
                                                v5 = (v5 ^ v5)
                                            } else {
                                                switch "a" {
                                                case "foo bar":
                                                    switch (fi110 % ((127) | 1)) {
                                                    case -601:
                                                        v96 = append(v96, fi110)
                                                    case 32767:
                                                        fmt.Println(v78)
                                                    default:
                                                        var v115 bool = false
                                                        _ = v115
                                                    }
                                                case "\n\t":
                                                    switch "a" {
                                                    case "a":
                                                        for fi116 := 0; fi116 < 2; fi116++ {
                                                            _ = len(v96)
                                                            v117 := v1[("\n\t" + "a")]
                                                            var v118 int64 = -9223372036854775808
                                                            _ = fi116
                                                            _ = v117
                                                            _ = v118
                                                        }
                                                    case "hello":
                                                        var v119 uint32 = uint32(v95)
                                                        _ = v119
                                                    default:
                                                        fmt.Println(v78, v2, v78)
                                                    }
                                                case "hello":
                                                    if false {
                                                        var v120 uint = uint(v5)
                                                        if (65535 <= v4) {
                                                            if true {
                                                                v120 = (v108 & 832)
                                                                var v121 int64 = 9223372036854775807
                                                                _ = v121
                                                            } else {
                                                                for fi122 := 0; fi122 < 19; fi122++ {
                                                                    if (fi122 != (fi110 << (((^(v120))) & 63))) {
                                                                        switch 32768 {
                                                                        case 32768:
                                                                            if false {
                                                                                v123 := GenS9{F0: 4294967295}
                                                                                _ = v123
                                                                            }
                                                                        case 747:
                                                                            if true {
                                                                                var v124 uint = 0
                                                                                _ = v124
                                                                            }
                                                                        }
                                                                        var v125 int16 = 1
                                                                        if len(v96) >= 2 {
                                                                            v96 = v96[1:len(v96)]
                                                                        }
                                                                        _ = v125
                                                                    } else {
                                                                        v96 = append(v96, 9223372036854775807)
                                                                        if (42 > -773) {
                                                                            _ = len(v96)
                                                                            var v126 uint64 = (v2 * 2)
                                                                            _ = v126
                                                                        } else {
                                                                            fmt.Println(v6, v3)
                                                                        }
                                                                        v127 := GenS3{F0: 2147483647, F1: int8(v2), F2: v120}
                                                                        _ = v127
                                                                    }
                                                                    v128 := GenS3{F0: -1, F1: 15, F2: 103}
                                                                    _ = fi122
                                                                    _ = v128
                                                                }
                                                            }
                                                        }
                                                        _ = v120
                                                    } else {
                                                        for fi129 := 0; fi129 < 1; fi129++ {
                                                            v3 -= (^(v3))
                                                            switch fi110 {
                                                            case 0:
                                                                fmt.Println(v95, v6, v6)
                                                            case -9223372036854775808:
                                                                var v130 float32 = v79
                                                                _ = v130
                                                            default:
                                                                switch "" {
                                                                case "foo bar":
                                                                    v131 := GenS1{F0: v2, F1: 1, F2: v4}
                                                                    _ = v131
                                                                case "hello":
                                                                    switch 4294967295 {
                                                                    case 127:
                                                                        if false {
                                                                            var v132 string = ((("" + "a") + "\n\t") + "foo bar")
                                                                            _ = v132
                                                                        }
                                                                    case 226:
                                                                        if (false && (42 <= -726)) {
                                                                            v78 = (((float64(1.5) * (v6 + (-2.25 * (1e10 * 3.14159)))) - v6) + (-2.25 / v6))
                                                                        }
                                                                    case -176:
                                                                        if false {
                                                                            var v133 int8 = -128
                                                                            v2 -= 42
                                                                            if false {
                                                                                if true {
                                                                                    fmt.Println(v95, v79)
                                                                                    v5 = (v5 - (v5 >> ((v108) & 63)))
                                                                                }
                                                                                var v134 bool = true
                                                                                if true {
                                                                                    fmt.Println(v3, v79, v95)
                                                                                }
                                                                                _ = v134
                                                                            }
                                                                            var v135 string = ("foo bar" + "foo bar")
                                                                            if false {
                                                                                fmt.Println(v133, v2)
                                                                                v3 = 351
                                                                                v136 := GenS11{F0: v4, F1: uint16(v5)}
                                                                                _ = v136
                                                                            } else {
                                                                                v137 := GenS12{F0: (v108 >> (((v108 | v108)) & 63))}
                                                                                fmt.Println(v95, v3)
                                                                                v138 := v1[v135]
                                                                                _ = v137
                                                                                _ = v138
                                                                            }
                                                                            _ = v133
                                                                            _ = v135
                                                                        } else {
                                                                            if true {
                                                                                var v139 int32 = 104
                                                                                if false {
                                                                                    v6 -= (v78 - v6)
                                                                                }
                                                                                _ = v139
                                                                            }
                                                                            v140 := GenS5{F0: 42, F1: 0}
                                                                            _ = v140
                                                                        }
                                                                        v141 := GenS12{F0: v108}
                                                                        if (42 <= 42) {
                                                                            v5 = (^(v5))
                                                                            switch ("a" + "") {
                                                                            case "":
                                                                                var v142 int64 = 42
                                                                                _ = v142
                                                                            case "\n\t":
                                                                                fmt.Println(fi110, v78, v78)
                                                                            case "a":
                                                                                switch fi110 {
                                                                                case 128:
                                                                                    v143 := GenS8{F0: (("a" + (("a" + "") + (("a" + (((("foo bar" + "a") + "hello") + "hello") + "hello")) + "\n\t"))) + "\n\t")}
                                                                                    _ = v143
                                                                                }
                                                                            }
                                                                        } else {
                                                                            v144 := GenS6{F0: "hello", F1: 0, F2: (^(v2))}
                                                                            _ = v144
                                                                        }
                                                                        _ = v141
                                                                    }
                                                                default:
                                                                    switch int(v2) {
                                                                    case 1:
                                                                        var v145 float32 = (0.1 * v79)
                                                                        _ = v145
                                                                    case 862:
                                                                        switch "hello" {
                                                                        case "a":
                                                                            v79 = 1.5
                                                                        case "foo bar":
                                                                            if true {
                                                                                v3 = v95
                                                                                var v146 int64 = -1
                                                                                _ = v146
                                                                            }
                                                                        case "hello":
                                                                            fmt.Println(v78, v79)
                                                                        default:
                                                                            v2 = (v2 * (v2 ^ (^(v2))))
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                            var v147 int64 = 743
                                                            _ = fi129
                                                            _ = v147
                                                        }
                                                        if (!false) {
                                                            for fi148 := 0; fi148 < 9; fi148++ {
                                                                v4 += (v4 | v4)
                                                                v6 -= (v6 - (((1.5 * 1.5) * (0.0 - v78)) + (v6 * v6)))
                                                                var v149 string = "\n\t"
                                                                v2 = (v2 & 0)
                                                                _ = fi148
                                                                _ = v149
                                                            }
                                                            var v150 int32 = 2147483647
                                                            _ = len(v96)
                                                            _ = v150
                                                        }
                                                        v108 = 2
                                                    }
                                                }
                                                v1["foo bar"] = uint32(v108)
                                            }
                                            var v151 uint8 = 0
                                            if false {
                                                for fi152 := 0; fi152 < 1; fi152++ {
                                                    v5 -= 42
                                                    fmt.Println(v151, v3, fi152)
                                                    _ = fi152
                                                }
                                                switch fi110 {
                                                case -2147483648:
                                                    if ((-128 < 127) || false) {
                                                        if len(v96) > 0 {
                                                            v96[0] = (fi110 % ((fi110) | 1))
                                                        }
                                                        v153 := GenS11{F0: 42, F1: v95}
                                                        v78 += v78
                                                        _ = v153
                                                    } else {
                                                        if (((1e10 - v6) == float64(1e10)) && true) {
                                                            if (-1 <= -128) {
                                                                if false {
                                                                    if len(v96) >= 2 {
                                                                        v96 = v96[1:len(v96)]
                                                                    }
                                                                } else {
                                                                    v2 = (^(v2))
                                                                }
                                                            }
                                                            var v154 float32 = -2.25
                                                            var v155 uint = (^(v108))
                                                            _ = v154
                                                            _ = v155
                                                        }
                                                    }
                                                case 526:
                                                    v156 := GenS13{F0: 606, F1: -128}
                                                    _ = v156
                                                case -1:
                                                    var v157 uint = v108
                                                    _ = v157
                                                default:
                                                    for fi158 := 0; fi158 < 3; fi158++ {
                                                        var v159 string = "a"
                                                        fmt.Println(v2, v108)
                                                        _ = fi158
                                                        _ = v159
                                                    }
                                                }
                                            } else {
                                                if true {
                                                    if false {
                                                        for fi160 := 0; fi160 < 17; fi160++ {
                                                            switch "\n\t" {
                                                            case "hello":
                                                                var v161 uint16 = 542
                                                                _ = v161
                                                            case "\n\t":
                                                                v162 := GenS8{F0: (("" + "\n\t") + ("hello" + "hello"))}
                                                                _ = v162
                                                            default:
                                                                v163 := v1["a"]
                                                                _ = v163
                                                            }
                                                            _ = fi160
                                                        }
                                                    }
                                                }
                                                for fi164 := 0; fi164 < 6; fi164++ {
                                                    var v165 bool = false
                                                    v4 -= v4
                                                    delete(v1, "foo bar")
                                                    _ = fi164
                                                    _ = v165
                                                }
                                                for fi166 := 0; fi166 < 5; fi166++ {
                                                    var v167 int16 = 0
                                                    fmt.Println(fi110)
                                                    fmt.Println(v78)
                                                    _ = fi166
                                                    _ = v167
                                                }
                                            }
                                            _ = v151
                                        }
                                    }
                                default:
                                    var v168 float32 = v79
                                    _ = v168
                                }
                                if (-536 <= int64(v2)) {
                                    for fi169 := 0; fi169 < 5; fi169++ {
                                        var v170 string = ("foo bar" + ("\n\t" + ("hello" + "")))
                                        v78 -= v6
                                        var v171 uint16 = (^(v95))
                                        _ = fi169
                                        _ = v170
                                        _ = v171
                                    }
                                } else {
                                    var v172 float64 = (((3.14159 - v6) + (3.14159 + float64(3.14159))) + (float64(3.14159) - 1.5))
                                    var v173 int8 = 127
                                    _ = len(v96)
                                    _ = v172
                                    _ = v173
                                }
                            }
                        }
                    }
                    v174 := GenS14{F0: uint32(v2)}
                    fmt.Println(v6, v5, v3)
                    _ = fi110
                    _ = v174
                }
            }
        }
        v108 = (v108 % ((v108) | 1))
        _ = v108
    }
    v3 = (v3 * 372)
    var v175 string = ("" + "\n\t")
    fmt.Println(v95)
    v176 := GenS7{F0: 1, F1: v175}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v175:", v175)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    fmt.Println("v95:", v95)
    fmt.Println("slicelen v96:", len(v96), v96)
    _ = v1
    _ = v175
    _ = v176
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v78
    _ = v79
    _ = v94
    _ = v95
    _ = v96
}
