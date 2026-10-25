package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 int8; F1 uint64; F2 int64 }
type GenS4 struct { F0 uint16; F1 uint32 }
type GenS5 struct { F0 int16; F1 int32; F2 uint32 }
type GenS6 struct { F0 int32; F1 bool; F2 int16 }
type GenS7 struct { F0 int32; F1 int32 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint64; F1 uint64 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 int32; F1 uint8; F2 int16 }
type GenS12 struct { F0 int64; F1 bool; F2 uint64 }
type GenS13 struct { F0 string }
type GenS14 struct { F0 int8; F1 int; F2 uint16 }
type GenS15 struct { F0 bool; F1 bool }
type GenS16 struct { F0 int32 }
type GenS17 struct { F0 uint64 }
type GenS18 struct { F0 uint64; F1 int32; F2 int8 }
type GenS19 struct { F0 uint16 }
type GenS20 struct { F0 string; F1 int64 }
type GenS21 struct { F0 uint8 }
func genFn0() int8 {
    switch "hello" {
    case "a":
        for fi6 := 0; fi6 < 11; fi6++ {
            for fi7 := 0; fi7 < 15; fi7++ {
                switch "hello" {
                case "":
                    var v8 int16 = 42
                    fmt.Println(fi6, fi7)
                    fmt.Println(fi6, fi6, v8)
                    v8 = int16(fi6)
                    v8 = (-(v8))
                    var v9 uint16 = 868
                    v9 += v9
                    var v10 int16 = (v8 + -350)
                    _ = v10
                    _ = v8
                    _ = v9
                case "a":
                    v11 := GenS1{F0: (!true)}
                    _ = v11
                case "hello":
                    v12 := []int{fi6, (fi7 - fi6), fi6, fi6}
                    _ = v12
                }
                fmt.Println(fi6, fi6)
                fmt.Println(fi6)
                var v13 uint64 = 0
                v14 := GenS1{F0: true}
                _ = fi7
                _ = v13
                _ = v14
            }
            for fi15 := 0; fi15 < 11; fi15++ {
                var v16 bool = (true || (int16(fi15) != 32767))
                var v17 int16 = 0
                _ = fi15
                _ = v16
                _ = v17
            }
            var v18 uint = 65535
            _ = fi6
            _ = v18
        }
    case "foo bar":
        var v19 uint32 = 0
        _ = v19
    case "hello":
        var v20 int64 = 1
        _ = v20
    default:
        switch 128 {
        case 32768:
            for fi21 := 0; fi21 < 19; fi21++ {
                var v22 int64 = 965
                fmt.Println(fi21, fi21)
                _ = fi21
                _ = v22
            }
        case 415:
            for fi23 := 0; fi23 < 14; fi23++ {
                var v24 int16 = 32767
                fmt.Println(v24, v24)
                _ = fi23
                _ = v24
            }
        case -9223372036854775808:
            var v25 int16 = 1000
            _ = v25
        default:
            var v26 int32 = 79
            _ = v26
        }
    }
    var v27 float32 = (1.5 + -2.25)
    if true {
        if true {
            var v28 float32 = (((float32(1e10) * float32(1e10)) / v27) / v27)
            if (-9223372036854775808 == 42) {
                if (0 != 0) {
                    var v29 int16 = -1
                    var v30 uint = 4294967295
                    _ = v29
                    _ = v30
                }
            }
            _ = v28
        }
        for fi31 := 0; fi31 < 9; fi31++ {
            fmt.Println(v27)
            _ = fi31
        }
        switch "foo bar" {
        case "a":
            var v32 uint32 = 2
            _ = v32
        default:
            v33 := []int{32768, 66, -2147483648, -139}
            _ = v33
        }
    }
    fmt.Println(v27)
    if (!false) {
        var v34 int8 = -77
        _ = v34
    } else {
        var v35 bool = ((true && false) || false)
        _ = v35
    }
    _ = v27
    return 42
}
func genFn1(p0 int8, p1 uint16, p2 uint32) int {
    switch 4294967295 {
    case -9223372036854775808:
        fmt.Println(p2)
    case 32768:
        p2 -= p2
    default:
        var v36 float64 = -2.25
        _ = v36
    }
    v37 := GenS1{F0: (!true)}
    if (!(!true)) {
        var v38 uint8 = 127
        var v39 int64 = -9223372036854775808
        _ = v38
        _ = v39
    } else {
        v40 := []int{-424, 2147483647, -1, 504}
        v41 := GenS1{F0: false}
        v42 := GenS2{F0: 0, F1: -1}
        _ = v40
        _ = v41
        _ = v42
    }
    if false {
        var v43 int64 = 14
        _ = v43
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v37
    return 196
}

func main() {
    var v1 bool = true
    v2 := make(map[string]uint32)
    v3 := make(map[string]uint32)
    var v4 uint = 42
    v5 := []uint32{65535, 1}
    v5 = append(v5, 65535)
    switch "" {
    case "a":
        if true {
            switch 127 {
            case 32768:
                if v1 {
                    if ("a" > "foo bar") {
                        var v44 uint = (v4 & (^(v4)))
                        for fi45 := 0; fi45 < 1; fi45++ {
                            switch (("foo bar" + "foo bar") + "") {
                            case "hello":
                                switch "" {
                                case "a":
                                    if ((2 != 352) || (!v1)) {
                                        v44 += (v44 / (((v44 * 757)) | 1))
                                    } else {
                                        var v46 float32 = ((1.5 - float32(0.1)) + 1.5)
                                        _ = v46
                                    }
                                case "":
                                    var v47 uint = 589
                                    _ = v47
                                case "\n\t":
                                    for fi48 := 0; fi48 < 2; fi48++ {
                                        fmt.Println(fi48, fi45)
                                        var v49 float64 = -2.25
                                        var v50 int16 = 32767
                                        _ = fi48
                                        _ = v49
                                        _ = v50
                                    }
                                }
                            }
                            switch (("foo bar" + ("hello" + "\n\t")) + "hello") {
                            case "foo bar":
                                for fi51 := 0; fi51 < 19; fi51++ {
                                    if true {
                                        v1 = v1
                                    }
                                    fmt.Println(fi45)
                                    _ = fi51
                                }
                            case "a":
                                v52 := GenS2{F0: 65535, F1: (fi45 << ((v44) & 63))}
                                _ = v52
                            case "\n\t":
                                v53 := GenS1{F0: v1}
                                _ = v53
                            }
                            v54 := GenS1{F0: (222 == 320)}
                            _ = fi45
                            _ = v54
                        }
                        v1 = v1
                        _ = v44
                    } else {
                        if v1 {
                            fmt.Println(v1)
                            var v55 uint16 = 65535
                            _ = v55
                        } else {
                            v4 -= (^(v4))
                            fmt.Println(v1)
                            var v56 uint8 = 127
                            _ = v56
                        }
                        v1 = v1
                        for fi57 := 0; fi57 < 16; fi57++ {
                            fmt.Println(v1)
                            _ = fi57
                        }
                    }
                    var v58 uint = 798
                    if len(v5) > 0 {
                        v5[0] = 2147483648
                    }
                    _ = v58
                }
            case -2147483648:
                if (42 < 65535) {
                    fmt.Println(v4, v1)
                } else {
                    v59 := GenS2{F0: 57, F1: -833}
                    for fi60 := 0; fi60 < 17; fi60++ {
                        for fi61 := 0; fi61 < 15; fi61++ {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            _ = fi61
                        }
                        for fi62 := 0; fi62 < 13; fi62++ {
                            v4 = (v4 & (v4 & v4))
                            v1 = v1
                            _ = fi62
                        }
                        _ = fi60
                    }
                    _ = v59
                }
            case 437:
                v4 = v4
            default:
                v3[(("foo bar" + "hello") + (((((("foo bar" + "") + "") + "foo bar") + "hello") + "a") + ""))] = 42
            }
            switch -37 {
            case 1:
                for fi63 := 0; fi63 < 19; fi63++ {
                    if false {
                        fmt.Println(fi63)
                        v4 -= (v4 % ((338) | 1))
                        var v64 uint64 = 641
                        _ = v64
                    }
                    v1 = (("hello" > "") || false)
                    v1 = (2 != 68)
                    _ = fi63
                }
            case 4294967295:
                for fi65 := 0; fi65 < 3; fi65++ {
                    for fi66 := 0; fi66 < 9; fi66++ {
                        v67 := GenS2{F0: 65535, F1: (fi66 & fi65)}
                        v5 = append(v5, 2147483648)
                        v68 := GenS3{F0: 49, F1: 0, F2: -1}
                        _ = fi66
                        _ = v67
                        _ = v68
                    }
                    var v69 int8 = 42
                    if ("hello" >= ("foo bar" + (("foo bar" + "") + ("\n\t" + "foo bar")))) {
                        for fi70 := 0; fi70 < 3; fi70++ {
                            if (-2.25 > -2.25) {
                                var v71 float64 = (float64(0.0) * float64(0.0))
                                _ = v71
                            }
                            _ = fi70
                        }
                    } else {
                        var v72 int32 = 42
                        var v73 uint8 = 6
                        fmt.Println(fi65)
                        _ = v72
                        _ = v73
                    }
                    _ = fi65
                    _ = v69
                }
            case -9223372036854775808:
                if false {
                    if true {
                        switch "a" {
                        case "":
                            v5 = append(v5, 1)
                        }
                    } else {
                        var v74 int = genFn1(genFn0(), 0, 2147483647)
                        v75 := v2[""]
                        _ = v74
                        _ = v75
                    }
                }
            default:
                if true {
                    v76 := GenS4{F0: 614, F1: 2147483647}
                    _ = v76
                } else {
                    v77 := GenS4{F0: 255, F1: 0}
                    for fi78 := 0; fi78 < 18; fi78++ {
                        switch "" {
                        case "a":
                            var v79 uint = 2
                            _ = v79
                        case "":
                            fmt.Println(v4, fi78, v4)
                        case "hello":
                            v80 := GenS1{F0: v1}
                            _ = v80
                        }
                        fmt.Println(v4, v1)
                        var v81 float32 = 0.0
                        _ = fi78
                        _ = v81
                    }
                    var v82 float64 = 3.14159
                    _ = v77
                    _ = v82
                }
            }
        }
    default:
        switch 9223372036854775807 {
        case -1:
            var v83 uint8 = 128
            _ = v83
        case 42:
            fmt.Println(v4)
        }
    }
    fmt.Println(v1, v4)
    if ((240 != 1) || v1) {
        if true {
            switch 42 {
            case 42:
                if v1 {
                    for fi84 := 0; fi84 < 7; fi84++ {
                        v5 = append(v5, 4294967295)
                        _ = fi84
                    }
                    for fi85 := 0; fi85 < 1; fi85++ {
                        switch (^(fi85)) {
                        case 0:
                            var v86 int32 = 42
                            _ = v86
                        case -9223372036854775808:
                            switch "a" {
                            case "a":
                                v87 := GenS4{F0: 80, F1: 2147483648}
                                _ = v87
                            case "foo bar":
                                switch ("foo bar" + ("foo bar" + "\n\t")) {
                                case "foo bar":
                                    switch "" {
                                    case "foo bar":
                                        fmt.Println(v1, v4, v4)
                                    case "":
                                        v88 := GenS5{F0: -143, F1: -1, F2: 2147483647}
                                        _ = v88
                                    default:
                                        var v89 uint8 = 127
                                        _ = v89
                                    }
                                default:
                                    v90 := GenS6{F0: 487, F1: v1, F2: 42}
                                    _ = v90
                                }
                            case "\n\t":
                                fmt.Println(v4, v4, fi85)
                            default:
                                if v1 {
                                    if (!v1) {
                                        fmt.Println(fi85)
                                    } else {
                                        var v91 int8 = 100
                                        _ = v91
                                    }
                                    if (((v1 || (!(!(!v1)))) || (42 > -983)) && true) {
                                        for fi92 := 0; fi92 < 19; fi92++ {
                                            v4 = (v4 - 0)
                                            var v93 uint32 = 79
                                            if (268 <= -879) {
                                                switch "hello" {
                                                case "":
                                                    var v94 int64 = 0
                                                    _ = v94
                                                default:
                                                    var v95 uint32 = (v93 / ((v93) | 1))
                                                    _ = v95
                                                }
                                                var v96 uint64 = 811
                                                _ = v96
                                            } else {
                                                var v97 uint = (v4 & (^(v4)))
                                                var v98 int16 = 32767
                                                _ = v97
                                                _ = v98
                                            }
                                            _ = fi92
                                            _ = v93
                                        }
                                        if ((fi85 - 128) < fi85) {
                                            var v99 uint64 = 898
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            _ = v99
                                        }
                                        if v1 {
                                            switch (fi85 >> ((18446744073709551615) & 63)) {
                                            case 270:
                                                if (613 != -1) {
                                                    v100 := GenS7{F0: 42, F1: 0}
                                                    if (!v1) {
                                                        switch "a" {
                                                        case "a":
                                                            var v101 uint32 = 106
                                                            _ = v101
                                                        case "\n\t":
                                                            var v102 int8 = 1
                                                            _ = v102
                                                        case "":
                                                            v3["a"] = 4294967295
                                                        }
                                                        var v103 int = (^(fi85))
                                                        _ = v103
                                                    } else {
                                                        v4 += (^(v4))
                                                        switch (("hello" + ((("a" + ("hello" + ("foo bar" + "hello"))) + ("a" + "foo bar")) + "hello")) + "foo bar") {
                                                        case "\n\t":
                                                            for fi104 := 0; fi104 < 4; fi104++ {
                                                                v4 -= v4
                                                                v105 := GenS6{F0: 42, F1: v1, F2: 1000}
                                                                var v106 float64 = 3.14159
                                                                _ = fi104
                                                                _ = v105
                                                                _ = v106
                                                            }
                                                        }
                                                    }
                                                    v2[("a" + "")] = 290
                                                    _ = v100
                                                } else {
                                                    v4 += v4
                                                    switch fi85 {
                                                    case 0:
                                                        for fi107 := 0; fi107 < 20; fi107++ {
                                                            v5 = append(v5, 1)
                                                            _ = fi107
                                                        }
                                                    case 32767:
                                                        var v108 uint32 = 0
                                                        _ = v108
                                                    default:
                                                        if true {
                                                            v1 = (true || (true || (!false)))
                                                        } else {
                                                            var v109 uint = v4
                                                            _ = v109
                                                        }
                                                    }
                                                    for fi110 := 0; fi110 < 12; fi110++ {
                                                        fmt.Println(fi110)
                                                        var v111 uint = 0
                                                        if v1 {
                                                            _ = len(v5)
                                                            if false {
                                                                v112 := GenS6{F0: -2147483648, F1: true, F2: 42}
                                                                _ = v112
                                                            }
                                                        } else {
                                                            fmt.Println(v1, v4, v4)
                                                            v4 = v4
                                                        }
                                                        v113 := GenS8{F0: 0}
                                                        _ = fi110
                                                        _ = v111
                                                        _ = v113
                                                    }
                                                }
                                            case 32767:
                                                for fi114 := 0; fi114 < 14; fi114++ {
                                                    var v115 uint32 = 65535
                                                    _ = fi114
                                                    _ = v115
                                                }
                                            default:
                                                if ((-2.25 - -2.25) == (3.14159 - 1.5)) {
                                                    v116 := v2[""]
                                                    for fi117 := 0; fi117 < 14; fi117++ {
                                                        switch (("a" + "\n\t") + "hello") {
                                                        case "a":
                                                            switch "a" {
                                                            case "foo bar":
                                                                fmt.Println(v4)
                                                            default:
                                                                var v118 uint64 = 126
                                                                _ = v118
                                                            }
                                                        case "\n\t":
                                                            if ((0 <= -9223372036854775808) || (!(!v1))) {
                                                                switch 127 {
                                                                case 42:
                                                                    if ((v1 || ((!v1) || ((v1 && (-2147483648 != 42)) || v1))) || (2147483647 <= 969)) {
                                                                        switch "a" {
                                                                        case "hello":
                                                                            v116 += v116
                                                                        case "":
                                                                            v119 := GenS2{F0: 270, F1: fi85}
                                                                            _ = v119
                                                                        }
                                                                        switch 869 {
                                                                        case 2147483647:
                                                                            switch ("hello" + "") {
                                                                            case "\n\t":
                                                                                v4 = v4
                                                                            case "":
                                                                                var v120 int8 = -128
                                                                                _ = v120
                                                                            }
                                                                        case 0:
                                                                            v121 := GenS2{F0: 32767, F1: fi117}
                                                                            _ = v121
                                                                        case -1:
                                                                            var v122 string = "a"
                                                                            _ = v122
                                                                        default:
                                                                            v1 = (1.5 < (0.0 / ((0.0 + -2.25) / 1.0)))
                                                                        }
                                                                    } else {
                                                                        var v123 bool = (!(v1 || (0 == -760)))
                                                                        _ = v123
                                                                    }
                                                                }
                                                                v4 = v4
                                                                var v124 uint64 = 0
                                                                _ = len(v3)
                                                                _ = v124
                                                            } else {
                                                                var v125 string = "a"
                                                                _ = v125
                                                            }
                                                        case "hello":
                                                            if ((^(v4)) != v4) {
                                                                v126 := GenS9{F0: 42, F1: uint64(v116)}
                                                                v116 = (v116 & v116)
                                                                _ = v126
                                                            }
                                                        default:
                                                            switch (-(fi85)) {
                                                            case 9223372036854775807:
                                                                var v127 int8 = genFn0()
                                                                _ = v127
                                                            case 2147483647:
                                                                var v128 int8 = 127
                                                                _ = v128
                                                            }
                                                        }
                                                        var v129 float32 = float32(1e10)
                                                        _ = fi117
                                                        _ = v129
                                                    }
                                                    switch 270 {
                                                    case 2147483647:
                                                        var v130 float32 = 1.5
                                                        _ = v130
                                                    case -869:
                                                        for fi131 := 0; fi131 < 8; fi131++ {
                                                            switch -698 {
                                                            case -1:
                                                                _ = len(v2)
                                                            case 42:
                                                                v116 = (v116 ^ (v116 >> ((v4) & 63)))
                                                            default:
                                                                v4 = (v4 + v4)
                                                            }
                                                            if v1 {
                                                                v4 += (v4 ^ v4)
                                                                _ = len(v2)
                                                                v132 := GenS10{F0: 2147483647}
                                                                _ = v132
                                                            } else {
                                                                v133 := GenS11{F0: -913, F1: 255, F2: 756}
                                                                _ = v133
                                                            }
                                                            _ = fi131
                                                        }
                                                    case 1:
                                                        for fi134 := 0; fi134 < 10; fi134++ {
                                                            v4 = 353
                                                            _ = fi134
                                                        }
                                                    default:
                                                        v135 := GenS2{F0: 510, F1: 4294967295}
                                                        _ = v135
                                                    }
                                                    _ = v116
                                                }
                                            }
                                            if false {
                                                var v136 int32 = 1
                                                _ = v136
                                            }
                                        } else {
                                            v1 = v1
                                            var v137 int16 = 0
                                            v1 = true
                                            _ = v137
                                        }
                                    } else {
                                        if len(v5) > 0 {
                                            v5[0] = 4294967295
                                        }
                                    }
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                } else {
                                    for fi138 := 0; fi138 < 15; fi138++ {
                                        if v1 {
                                            var v139 int = fi85
                                            v4 = v4
                                            var v140 int64 = -318
                                            var v141 uint64 = 18446744073709551615
                                            _ = v139
                                            _ = v140
                                            _ = v141
                                        }
                                        fmt.Println(v4, v1, v1)
                                        _ = fi138
                                    }
                                    if len(v5) > 0 {
                                        v5[0] = 2
                                    }
                                }
                            }
                        }
                        _ = fi85
                    }
                }
            case -1:
                if (2 == (v4 + (^(v4)))) {
                    _ = len(v5)
                    v4 -= v4
                } else {
                    v4 = v4
                    switch 335 {
                    case 42:
                        v1 = true
                    default:
                        for fi142 := 0; fi142 < 19; fi142++ {
                            v4 = v4
                            _ = fi142
                        }
                    }
                    if (1 < -555) {
                        for fi143 := 0; fi143 < 10; fi143++ {
                            var v144 int64 = 640
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            v145 := GenS9{F0: 1, F1: 0}
                            _ = fi143
                            _ = v144
                            _ = v145
                        }
                        if true {
                            v5 = append(v5, 74)
                        } else {
                            if false {
                                fmt.Println(v1, v4)
                                var v146 bool = true
                                for fi147 := 0; fi147 < 11; fi147++ {
                                    for fi148 := 0; fi148 < 4; fi148++ {
                                        v149 := v3[("\n\t" + ("a" + "hello"))]
                                        if v146 {
                                            switch "a" {
                                            case "":
                                                switch (("\n\t" + "") + "a") {
                                                case "":
                                                    v149 = v149
                                                }
                                                if v1 {
                                                    v150 := GenS12{F0: -744, F1: v1, F2: uint64(v149)}
                                                    var v151 uint64 = 42
                                                    _ = len(v5)
                                                    _ = v150
                                                    _ = v151
                                                }
                                            case "\n\t":
                                                _ = len(v5)
                                            case "foo bar":
                                                if v146 {
                                                    v152 := GenS6{F0: 0, F1: ((3.14159 - 1.5) < 0.1), F2: -693}
                                                    v4 = v4
                                                    var v153 float64 = 0.0
                                                    var v154 int16 = 1000
                                                    _ = v152
                                                    _ = v153
                                                    _ = v154
                                                }
                                            default:
                                                v146 = v1
                                            }
                                            var v155 int16 = 191
                                            _ = v155
                                        }
                                        _ = fi148
                                        _ = v149
                                    }
                                    v1 = v146
                                    _ = fi147
                                }
                                _ = v146
                            } else {
                                if v1 {
                                    v156 := GenS7{F0: -762, F1: 0}
                                    _ = v156
                                } else {
                                    var v157 uint32 = 11
                                    _ = v157
                                }
                            }
                            v4 += (v4 & 58)
                        }
                        var v158 uint = v4
                        _ = v158
                    }
                }
            }
        } else {
            delete(v2, (("hello" + "foo bar") + "\n\t"))
            v4 = (v4 ^ 18446744073709551615)
        }
    }
    var v159 float64 = float64((0.1 * 3.14159))
    if v1 {
        v160 := GenS1{F0: (4294967295 <= 4294967295)}
        switch 1 {
        case -686:
            _ = len(v3)
        case 128:
            fmt.Println(v1)
        case 32768:
            var v161 string = ""
            _ = v161
        }
        if (v1 || ("hello" < "\n\t")) {
            if ((127 <= 0) && (!(!(625 <= 65535)))) {
                for fi162 := 0; fi162 < 15; fi162++ {
                    var v163 uint = v4
                    fmt.Println(v159, v1)
                    _ = fi162
                    _ = v163
                }
                switch 253 {
                case 128:
                    var v164 float64 = v159
                    _ = v164
                case 127:
                    for fi165 := 0; fi165 < 12; fi165++ {
                        for fi166 := 0; fi166 < 10; fi166++ {
                            v1 = true
                            v1 = (!true)
                            if true {
                                var v167 int16 = 0
                                var v168 uint32 = 2147483648
                                switch -1 {
                                case 127:
                                    fmt.Println(fi165, v167)
                                }
                                _ = v167
                                _ = v168
                            } else {
                                v169 := GenS8{F0: -32768}
                                var v170 uint64 = 733
                                v4 -= 42
                                _ = v169
                                _ = v170
                            }
                            _ = len(v2)
                            v171 := GenS12{F0: 0, F1: true, F2: 4294967295}
                            if true {
                                v4 = v4
                                if v1 {
                                    if ((("" + ("\n\t" + "foo bar")) > ("hello" + "hello")) && v1) {
                                        v172 := GenS12{F0: -9223372036854775808, F1: (2147483648 <= 2147483648), F2: 45}
                                        switch (("\n\t" + "a") + ("" + "foo bar")) {
                                        case "\n\t":
                                            var v173 int = 995
                                            _ = v173
                                        }
                                        _ = v172
                                    } else {
                                        v174 := GenS13{F0: (("foo bar" + ("" + "\n\t")) + ("\n\t" + ""))}
                                        if v1 {
                                            switch ("a" + "foo bar") {
                                            case "\n\t":
                                                v1 = false
                                            case "hello":
                                                if v1 {
                                                    var v175 float64 = (0.0 + (v159 + -2.25))
                                                    v175 += v175
                                                    _ = v175
                                                } else {
                                                    v176 := GenS3{F0: genFn0(), F1: 863, F2: -9223372036854775808}
                                                    if false {
                                                        fmt.Println(v1, v1, v4)
                                                        v1 = v1
                                                        var v177 int16 = 1000
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                        v159 = float64(0.1)
                                                        _ = v177
                                                    }
                                                    var v178 uint8 = 127
                                                    _ = v176
                                                    _ = v178
                                                }
                                            case "":
                                                v4 -= (v4 * (v4 * (v4 / ((v4) | 1))))
                                            }
                                            switch "a" {
                                            case "a":
                                                v159 = -2.25
                                            case "foo bar":
                                                fmt.Println(v4, v159)
                                            case "\n\t":
                                                var v179 uint32 = 0
                                                _ = v179
                                            default:
                                                v4 += (v4 % ((v4) | 1))
                                            }
                                        }
                                        switch "\n\t" {
                                        case "hello":
                                            switch "a" {
                                            case "":
                                                v180 := v3[("a" + (("a" + ((("" + "a") + ((((("\n\t" + "foo bar") + ("hello" + (("\n\t" + ("\n\t" + ("a" + "hello"))) + ""))) + "a") + ("a" + ("foo bar" + ("hello" + ("\n\t" + "\n\t"))))) + "a")) + (("foo bar" + "foo bar") + "a"))) + ""))]
                                                _ = v180
                                            case "a":
                                                fmt.Println(v1)
                                            case "hello":
                                                fmt.Println(fi166)
                                                var v181 float64 = (v159 + float64(0.0))
                                                _ = v181
                                            default:
                                                v182 := v3["foo bar"]
                                                _ = v182
                                            }
                                        default:
                                            v5 = append(v5, 0)
                                        }
                                        _ = v174
                                    }
                                }
                            }
                            _ = fi166
                            _ = v171
                        }
                        _ = len(v5)
                        if true {
                            var v183 int64 = 1
                            _ = v183
                        } else {
                            var v184 uint16 = 2
                            v185 := GenS14{F0: 73, F1: (fi165 + (fi165 * 42)), F2: v184}
                            _ = v184
                            _ = v185
                        }
                        _ = fi165
                    }
                case 2147483647:
                    fmt.Println(v1, v4)
                default:
                    v5 = append(v5, 42)
                }
            }
            if (v1 || true) {
                for fi186 := 0; fi186 < 5; fi186++ {
                    var v187 bool = (!("hello" < "a"))
                    if v1 {
                        v188 := GenS15{F0: (!((!v187) || v187)), F1: ("hello" >= "\n\t")}
                        fmt.Println(fi186)
                        for fi189 := 0; fi189 < 18; fi189++ {
                            v1 = false
                            if ("a" == "hello") {
                                fmt.Println(v159, fi189, v4)
                                v2["a"] = 42
                                switch "hello" {
                                case "\n\t":
                                    if ((1e10 + 0.0) > 0.1) {
                                        if v187 {
                                            fmt.Println(fi189)
                                            fmt.Println(v187, fi189)
                                            fmt.Println(v159, fi189)
                                        }
                                        v4 -= v4
                                        if (4294967295 != 2147483648) {
                                            v190 := GenS11{F0: 0, F1: 255, F2: 1}
                                            v191 := GenS14{F0: 127, F1: (fi186 - 0), F2: 65535}
                                            v1 = true
                                            _ = v190
                                            _ = v191
                                        } else {
                                            fmt.Println(v159, v4, fi186)
                                            v159 += v159
                                            v192 := GenS16{F0: 42}
                                            _ = v192
                                        }
                                        switch "" {
                                        case "":
                                            if v187 {
                                                v1 = v187
                                                v159 = float64(((3.14159 + 3.14159) / float32(1.5)))
                                            } else {
                                                v193 := GenS11{F0: 577, F1: 2, F2: 0}
                                                v1 = (0 > 0)
                                                var v194 uint16 = 65535
                                                _ = v193
                                                _ = v194
                                            }
                                        default:
                                            if true {
                                                switch fi186 {
                                                case 64:
                                                    var v195 uint64 = 18446744073709551615
                                                    _ = v195
                                                case 0:
                                                    v196 := GenS1{F0: false}
                                                    _ = v196
                                                case -2147483648:
                                                    v197 := v2["\n\t"]
                                                    _ = v197
                                                }
                                                v1 = (77 <= int8(v4))
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                            } else {
                                                if true {
                                                    v198 := GenS3{F0: 1, F1: 0, F2: -5}
                                                    _ = v198
                                                }
                                                fmt.Println(fi186, fi189)
                                                if true {
                                                    if v1 {
                                                        var v199 float32 = (0.0 * ((float32(3.14159) / 3.14159) * 3.14159))
                                                        fmt.Println(fi186, v187, v187)
                                                        v4 = 4294967295
                                                        _ = v199
                                                    } else {
                                                        if (18446744073709551615 != 2) {
                                                            v200 := GenS16{F0: 42}
                                                            if (158 <= v4) {
                                                                v159 += 1.5
                                                                switch 163 {
                                                                case 127:
                                                                    v201 := GenS7{F0: -2147483648, F1: 168}
                                                                    _ = v201
                                                                case 2147483647:
                                                                    if len(v5) > 0 {
                                                                        v5[0] = 42
                                                                    }
                                                                case 32768:
                                                                    v202 := GenS9{F0: 42, F1: 2}
                                                                    _ = v202
                                                                default:
                                                                    if (!(v159 >= (0.1 * v159))) {
                                                                        switch "" {
                                                                        case "a":
                                                                            fmt.Println(v4, v4)
                                                                        }
                                                                    }
                                                                }
                                                                if v187 {
                                                                    delete(v2, "hello")
                                                                } else {
                                                                    var v203 int32 = 765
                                                                    _ = v203
                                                                }
                                                            }
                                                            _ = v200
                                                        }
                                                        v159 += v159
                                                    }
                                                } else {
                                                    delete(v2, (("\n\t" + ("hello" + "")) + "\n\t"))
                                                    v204 := GenS10{F0: 0}
                                                    _ = v204
                                                }
                                            }
                                        }
                                    } else {
                                        var v205 uint64 = 596
                                        switch (fi186 + (fi186 % (((^(fi186))) | 1))) {
                                        case -2147483648:
                                            delete(v3, "")
                                        case -907:
                                            v4 -= v4
                                        case 1:
                                            v4 -= (^(v4))
                                            v159 -= v159
                                            if true {
                                                var v206 float64 = 1e10
                                                switch "hello" {
                                                case "foo bar":
                                                    fmt.Println(v206, v4, v159)
                                                    v4 += v4
                                                case "a":
                                                    var v207 uint16 = 255
                                                    _ = v207
                                                case "\n\t":
                                                    if false {
                                                        v208 := GenS2{F0: 1, F1: (fi189 - (fi189 ^ (^(fi186))))}
                                                        _ = v208
                                                    } else {
                                                        var v209 float32 = 3.14159
                                                        fmt.Println(v209, v206)
                                                        v210 := GenS17{F0: (v205 >> ((9) & 63))}
                                                        _ = v209
                                                        _ = v210
                                                    }
                                                }
                                                var v211 uint8 = 21
                                                _ = v206
                                                _ = v211
                                            }
                                        default:
                                            var v212 int64 = 42
                                            _ = v212
                                        }
                                        if v187 {
                                            var v213 string = "a"
                                            switch genFn1(-1, 1, 742) {
                                            case -2147483648:
                                                v1 = false
                                                switch fi186 {
                                                case 32767:
                                                    var v214 int = genFn1(100, 255, 0)
                                                    _ = v214
                                                case 335:
                                                    switch (fi186 ^ genFn1(1, 65535, 711)) {
                                                    case -664:
                                                        var v215 uint32 = 0
                                                        _ = v215
                                                    case 127:
                                                        switch ((v213 + v213) + "foo bar") {
                                                        case "":
                                                            switch fi186 {
                                                            case -153:
                                                                delete(v2, "\n\t")
                                                            case 2147483647:
                                                                v187 = (v187 && v1)
                                                            case -2147483648:
                                                                v205 = (^(v205))
                                                            default:
                                                                if v187 {
                                                                    v1 = (47 == 127)
                                                                    v1 = (0.1 == 1e10)
                                                                }
                                                            }
                                                        case "\n\t":
                                                            var v216 int32 = 441
                                                            var v217 float32 = 1.5
                                                            var v218 int64 = 9223372036854775807
                                                            _ = v216
                                                            _ = v217
                                                            _ = v218
                                                        case "hello":
                                                            var v219 int16 = 42
                                                            _ = v219
                                                        }
                                                    case -631:
                                                        v1 = v187
                                                    }
                                                case 32768:
                                                    v205 += 4294967295
                                                default:
                                                    v220 := GenS2{F0: 65535, F1: -604}
                                                    _ = v220
                                                }
                                            case -688:
                                                v159 = (1.5 * ((v159 / ((float64(0.1) + 1e10) - v159)) - v159))
                                            default:
                                                fmt.Println(v4, v213, v187)
                                            }
                                            _ = v213
                                        } else {
                                            fmt.Println(v159)
                                        }
                                        _ = v205
                                    }
                                case "foo bar":
                                    v4 = 39
                                default:
                                    _ = len(v5)
                                }
                            } else {
                                _ = len(v3)
                                var v221 int64 = 9223372036854775807
                                _ = v221
                            }
                            if v187 {
                                if (v187 || ((362 < 391) && (42 <= 0))) {
                                    var v222 uint8 = 127
                                    var v223 uint8 = 42
                                    var v224 uint8 = (^(v222))
                                    _ = v222
                                    _ = v223
                                    _ = v224
                                }
                                fmt.Println(fi189)
                                if (1000 < -32768) {
                                    v4 += (^(v4))
                                    var v225 int8 = -81
                                    _ = v225
                                } else {
                                    if v1 {
                                        v226 := GenS15{F0: ((((0.1 * 0.0) - float32(v159)) > 3.14159) && v187), F1: (fi189 >= fi186)}
                                        var v227 int8 = 127
                                        _ = v226
                                        _ = v227
                                    }
                                    if (128 < (-(fi186))) {
                                        switch fi186 {
                                        case 2147483647:
                                            v228 := GenS3{F0: 27, F1: 184, F2: 42}
                                            _ = v228
                                        case -1:
                                            if true {
                                                v5 = append(v5, 2)
                                                _ = len(v3)
                                            } else {
                                                var v229 uint16 = 65535
                                                if true {
                                                    v1 = true
                                                    v4 = (v4 | v4)
                                                    v159 = 1.5
                                                    fmt.Println(v159, fi186, fi189)
                                                    v230 := GenS4{F0: v229, F1: 65535}
                                                    _ = v230
                                                }
                                                delete(v3, ("hello" + (("" + ("hello" + "\n\t")) + ((("hello" + "a") + ("a" + "foo bar")) + ""))))
                                                _ = v229
                                            }
                                        case 9223372036854775807:
                                            fmt.Println(v1, v187)
                                        default:
                                            v187 = v1
                                        }
                                        switch ("hello" + "foo bar") {
                                        case "foo bar":
                                            delete(v3, ("foo bar" + ("hello" + "\n\t")))
                                        case "hello":
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                        }
                                        delete(v3, ("hello" + "\n\t"))
                                    } else {
                                        v159 = float64(1.5)
                                    }
                                    fmt.Println(v187, fi186)
                                    var v231 int32 = 2147483647
                                    if v187 {
                                        v1 = v187
                                        fmt.Println(v187)
                                        v232 := GenS12{F0: -718, F1: v1, F2: 18446744073709551615}
                                        _ = v232
                                    } else {
                                        v231 -= (v231 - v231)
                                    }
                                    v233 := GenS2{F0: 610, F1: fi189}
                                    _ = v231
                                    _ = v233
                                }
                            } else {
                                var v234 float64 = v159
                                _ = v234
                            }
                            v235 := GenS15{F0: v1, F1: v1}
                            v159 = -2.25
                            _ = fi189
                            _ = v235
                        }
                        _ = v188
                    }
                    _ = fi186
                    _ = v187
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
        }
        _ = v160
    }
    v236 := GenS18{F0: 4294967295, F1: -1, F2: -2}
    for fi237 := 0; fi237 < 20; fi237++ {
        switch fi237 {
        case 9223372036854775807:
            fmt.Println(v4)
        case 32768:
            switch genFn1(3, 463, 1) {
            case 127:
                v238 := GenS19{F0: 42}
                _ = v238
            case 128:
                for fi239 := 0; fi239 < 5; fi239++ {
                    v240 := GenS20{F0: "hello", F1: 1}
                    var v241 uint8 = 255
                    var v242 float64 = (-2.25 * 3.14159)
                    if v1 {
                        if v1 {
                            var v243 int = genFn1(127, 265, 1)
                            var v244 float64 = ((v159 * v242) / v242)
                            fmt.Println(v1)
                            _ = v243
                            _ = v244
                        }
                    } else {
                        var v245 uint = 18446744073709551615
                        _ = v245
                    }
                    var v246 float32 = (1.5 - 1.5)
                    _ = fi239
                    _ = v240
                    _ = v241
                    _ = v242
                    _ = v246
                }
            case 549:
                fmt.Println(v159, v159)
            }
        case -9223372036854775808:
            if (!("hello" > ("a" + ""))) {
                var v247 int16 = -32768
                _ = v247
            }
        }
        v248 := GenS21{F0: 42}
        _ = fi237
        _ = v248
    }
    delete(v2, ((("a" + (("hello" + "") + "foo bar")) + "foo bar") + "foo bar"))
    v249 := GenS20{F0: "", F1: -1}
    for fi250 := 0; fi250 < 9; fi250++ {
        var v251 uint = (v4 & v4)
        _ = fi250
        _ = v251
    }
    if (127 == genFn1(-92, 0, 975)) {
        var v252 int = 128
        var v253 uint32 = 0
        _ = v252
        _ = v253
    }
    v254 := GenS2{F0: 1, F1: -651}
    var v255 int = -276
    var v256 uint = 255
    v257 := GenS11{F0: 2147483647, F1: 255, F2: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v159:", v159)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v255:", v255)
    fmt.Println("v256:", v256)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v159
    _ = v2
    _ = v236
    _ = v249
    _ = v254
    _ = v255
    _ = v256
    _ = v257
    _ = v3
    _ = v4
    _ = v5
}
