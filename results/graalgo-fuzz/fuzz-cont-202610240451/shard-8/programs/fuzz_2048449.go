package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 string }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 uint8; F1 int64; F2 int32 }
type GenS4 struct { F0 int8; F1 int64; F2 int32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int32; F1 bool }
type GenS7 struct { F0 int8; F1 bool }
type GenS8 struct { F0 uint16; F1 string; F2 int }
type GenS9 struct { F0 uint64; F1 int16; F2 int }
type GenS10 struct { F0 uint32; F1 int8; F2 uint8 }
type GenS11 struct { F0 uint; F1 bool; F2 int8 }
type GenS12 struct { F0 int; F1 uint; F2 int32 }
type GenS13 struct { F0 int16 }
type GenS14 struct { F0 uint8 }
type GenS15 struct { F0 int64 }
type GenS16 struct { F0 uint64; F1 int8; F2 uint16 }
type GenS17 struct { F0 int8; F1 int64 }
type GenS18 struct { F0 uint16 }
type GenS19 struct { F0 uint8; F1 int8; F2 int64 }
type GenS20 struct { F0 uint64; F1 int64 }
type GenS21 struct { F0 uint32 }
type GenS22 struct { F0 int16 }
type GenS23 struct { F0 uint64 }
type GenS24 struct { F0 uint32; F1 uint8; F2 int }
type GenS25 struct { F0 string; F1 bool }
type GenS26 struct { F0 int8; F1 int; F2 bool }
type GenS27 struct { F0 int16; F1 int64 }
type GenS28 struct { F0 uint16; F1 uint; F2 bool }
func genFn0() int16 {
    if (true || true) {
        v4 := []int{548, -9223372036854775808}
        switch "hello" {
        case "\n\t":
            var v5 uint = 18446744073709551615
            _ = v5
        case "hello":
            for fi6 := 0; fi6 < 13; fi6++ {
                for fi7 := 0; fi7 < 8; fi7++ {
                    switch (fi6 - 32767) {
                    case 32767:
                        fmt.Println(fi7, fi6)
                    case 4294967295:
                        var v8 bool = ((-1 <= 1) || (!(!false)))
                        var v9 bool = (!(0 >= 0))
                        v9 = v8
                        _ = v8
                        _ = v9
                    }
                    _ = fi7
                }
                v10 := GenS1{F0: 0, F1: 0, F2: "foo bar"}
                _ = fi6
                _ = v10
            }
        case "foo bar":
            var v11 uint32 = 2
            _ = v11
        default:
            var v12 string = ""
            _ = v12
        }
        _ = v4
    } else {
        v13 := []int{-462, 32768}
        _ = v13
    }
    switch -9223372036854775808 {
    case 127:
        var v14 string = ("hello" + ("" + "foo bar"))
        _ = v14
    case 32767:
        switch 32768 {
        case 9223372036854775807:
            switch 0 {
            case -9223372036854775808:
                v15 := []int{563, 949, 895, 9223372036854775807}
                _ = v15
            case 2147483647:
                v16 := []int{}
                _ = v16
            }
        case 128:
            var v17 uint8 = 113
            _ = v17
        }
    case -436:
        if (42 != 1) {
            if true {
                for fi18 := 0; fi18 < 6; fi18++ {
                    var v19 int64 = 173
                    v20 := GenS2{F0: 42, F1: 2}
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                var v21 uint = 65535
                _ = v21
            }
            switch -9223372036854775808 {
            case -2147483648:
                for fi22 := 0; fi22 < 8; fi22++ {
                    if (!true) {
                        var v23 uint16 = 800
                        var v24 uint16 = (v23 >> ((767) & 63))
                        for fi25 := 0; fi25 < 12; fi25++ {
                            var v26 string = "foo bar"
                            if true {
                                fmt.Println(v23, v23, fi25)
                                v27 := GenS3{F0: 0, F1: 486, F2: 2147483647}
                                if (!true) {
                                    if true {
                                        v28 := GenS3{F0: 255, F1: 9223372036854775807, F2: 733}
                                        if false {
                                            var v29 uint64 = 686
                                            if false {
                                                if true {
                                                    v24 = (v23 >> ((711) & 63))
                                                    v24 = v23
                                                    v23 = uint16(v29)
                                                } else {
                                                    v30 := GenS1{F0: 1, F1: 895, F2: ("foo bar" + "hello")}
                                                    v23 -= v24
                                                    v23 += 1
                                                    _ = v30
                                                }
                                            } else {
                                                v23 = (^(v24))
                                                switch "a" {
                                                case "\n\t":
                                                    if (100 != 94) {
                                                        v31 := []int{(-(fi25)), 127, fi25, 32767}
                                                        v29 = (v29 ^ 2)
                                                        switch 0 {
                                                        case 4294967295:
                                                            v26 = v26
                                                            var v32 string = (v26 + "hello")
                                                            var v33 uint16 = (v23 << ((4294967295) & 63))
                                                            v26 = ""
                                                            _ = v32
                                                            _ = v33
                                                        }
                                                        _ = v31
                                                    }
                                                case "":
                                                    v29 += 1
                                                }
                                            }
                                            switch (v26 + v26) {
                                            case "foo bar":
                                                var v34 float64 = 3.14159
                                                _ = v34
                                            default:
                                                fmt.Println(fi25, fi25)
                                            }
                                            _ = v29
                                        } else {
                                            if true {
                                                switch "a" {
                                                case "":
                                                    v24 = (v23 % ((0) | 1))
                                                case "hello":
                                                    v35 := GenS4{F0: 0, F1: 1, F2: -2147483648}
                                                    _ = v35
                                                }
                                                v26 = "foo bar"
                                            }
                                            fmt.Println(fi25, v23, v23)
                                        }
                                        var v36 uint16 = (^(v24))
                                        _ = v28
                                        _ = v36
                                    }
                                }
                                _ = v27
                            }
                            v37 := []int{(-(fi25))}
                            _ = fi25
                            _ = v26
                            _ = v37
                        }
                        _ = v23
                        _ = v24
                    } else {
                        if false {
                            var v38 string = "\n\t"
                            v39 := GenS5{F0: 1}
                            var v40 int64 = 1
                            _ = v38
                            _ = v39
                            _ = v40
                        }
                    }
                    for fi41 := 0; fi41 < 9; fi41++ {
                        if false {
                            var v42 string = ("foo bar" + "\n\t")
                            var v43 uint = 255
                            _ = v42
                            _ = v43
                        } else {
                            v44 := []int{fi22, -9223372036854775808}
                            var v45 uint = 18446744073709551615
                            if (int8(fi41) != 100) {
                                v46 := GenS3{F0: 0, F1: 1, F2: 593}
                                fmt.Println(fi41, fi22)
                                _ = v46
                            }
                            _ = v44
                            _ = v45
                        }
                        var v47 uint = 0
                        var v48 int = fi41
                        var v49 int = (fi41 * (fi41 - 0))
                        var v50 string = "a"
                        _ = fi41
                        _ = v47
                        _ = v48
                        _ = v49
                        _ = v50
                    }
                    switch (fi22 | (fi22 & 2147483647)) {
                    case 970:
                        var v51 int16 = -721
                        _ = v51
                    default:
                        fmt.Println(fi22)
                    }
                    _ = fi22
                }
            case -1:
                for fi52 := 0; fi52 < 3; fi52++ {
                    v53 := GenS6{F0: -2147483648, F1: true}
                    var v54 int32 = -368
                    if (2 < 42) {
                        if (2 != 2) {
                            v54 = (-(v54))
                            var v55 uint8 = 255
                            _ = v55
                        } else {
                            v54 = v54
                            var v56 uint8 = 255
                            _ = v56
                        }
                        v54 = (v54 + (v54 % ((1) | 1)))
                        v57 := GenS4{F0: 1, F1: 256, F2: v54}
                        _ = v57
                    }
                    _ = fi52
                    _ = v53
                    _ = v54
                }
            }
            var v58 uint8 = 0
            _ = v58
        } else {
            for fi59 := 0; fi59 < 16; fi59++ {
                if true {
                    if true {
                        if true {
                            v60 := GenS5{F0: 42}
                            _ = v60
                        }
                        var v61 uint8 = 209
                        _ = v61
                    }
                    if (!(1 != 65535)) {
                        var v62 uint8 = 138
                        v62 = v62
                        v63 := GenS5{F0: 1}
                        _ = v62
                        _ = v63
                    }
                }
                _ = fi59
            }
            v64 := []int{921}
            _ = v64
        }
    default:
        v65 := []int{0, 32768, 432}
        _ = v65
    }
    var v66 string = ""
    if false {
        for fi67 := 0; fi67 < 2; fi67++ {
            var v68 float64 = 3.14159
            switch fi67 {
            case -169:
                for fi69 := 0; fi69 < 1; fi69++ {
                    v70 := GenS5{F0: 1}
                    var v71 string = v66
                    switch fi67 {
                    case 0:
                        var v72 string = v71
                        _ = v72
                    case 128:
                        var v73 int8 = -128
                        _ = v73
                    }
                    _ = fi69
                    _ = v70
                    _ = v71
                }
            case 1:
                if (!(65535 >= 0)) {
                    for fi74 := 0; fi74 < 3; fi74++ {
                        v75 := []int{(fi67 % ((fi74) | 1)), fi74, (fi67 - (fi74 | (fi67 << ((255) & 63))))}
                        var v76 int64 = 42
                        _ = fi74
                        _ = v75
                        _ = v76
                    }
                } else {
                    v68 = ((float64((((1.5 + float32(0.1)) - 0.0) - 1.5)) + (v68 / (float64((float32((v68 * v68)) - 1.5)) + v68))) + (v68 - (1e10 - v68)))
                    v77 := []int{fi67}
                    var v78 uint32 = 2147483647
                    _ = v77
                    _ = v78
                }
            case 0:
                var v79 uint16 = 624
                _ = v79
            }
            var v80 int16 = 745
            _ = fi67
            _ = v68
            _ = v80
        }
        var v81 int8 = 127
        _ = v81
    } else {
        fmt.Println(v66)
        v82 := []int{}
        _ = v82
    }
    v83 := GenS7{F0: 62, F1: true}
    _ = v66
    _ = v83
    return 750
}
func genFn1(p0 uint16, p1 int) int {
    if false {
        fmt.Println(p0)
        for fi84 := 0; fi84 < 6; fi84++ {
            if (209 <= 255) {
                switch 32768 {
                case 292:
                    var v85 float32 = 1.5
                    _ = v85
                case 9223372036854775807:
                    switch (-(p1)) {
                    case 795:
                        var v86 float32 = (0.1 - (float32(1e10) / 1.0))
                        _ = v86
                    case 973:
                        p1 -= fi84
                    default:
                        p0 -= p0
                    }
                }
                if false {
                    v87 := GenS8{F0: p0, F1: "\n\t", F2: p1}
                    for fi88 := 0; fi88 < 18; fi88++ {
                        v89 := []int{(fi84 ^ fi84), fi88, 9223372036854775807}
                        switch "hello" {
                        case "a":
                            if (2 != 2) {
                                if (1 < 1) {
                                    var v90 float64 = 1e10
                                    p1 = (fi84 | (fi88 & 370))
                                    if (-971 < 0) {
                                        if (false || (!false)) {
                                            if (1 >= 1) {
                                                p0 -= p0
                                            }
                                            var v91 int = 285
                                            _ = v91
                                        }
                                    } else {
                                        var v92 uint32 = 42
                                        var v93 uint = 837
                                        if (128 > 255) {
                                            var v94 uint64 = 18446744073709551615
                                            _ = v94
                                        } else {
                                            switch p1 {
                                            case 9223372036854775807:
                                                if true {
                                                    if true {
                                                        switch "a" {
                                                        case "a":
                                                            v95 := GenS9{F0: 4294967295, F1: -290, F2: int(v92)}
                                                            _ = v95
                                                        }
                                                        v92 -= v92
                                                        var v96 int8 = 127
                                                        _ = v96
                                                    }
                                                    v97 := GenS2{F0: 42, F1: 2}
                                                    _ = v97
                                                } else {
                                                    var v98 uint8 = 137
                                                    _ = len(v89)
                                                    _ = v98
                                                }
                                            }
                                            v99 := GenS10{F0: (v92 - v92), F1: 20, F2: 42}
                                            _ = v99
                                        }
                                        _ = v92
                                        _ = v93
                                    }
                                    _ = v90
                                }
                                v89 = append(v89, fi88)
                            } else {
                                if len(v89) > 0 {
                                    v89[0] = (p1 | -785)
                                }
                                var v100 int = 4294967295
                                _ = v100
                            }
                        case "hello":
                            v101 := GenS7{F0: 42, F1: true}
                            _ = v101
                        }
                        if false {
                            var v102 float64 = 0.0
                            _ = v102
                        } else {
                            if (-721 > 960) {
                                fmt.Println(p1, p0)
                            } else {
                                switch "\n\t" {
                                case "a":
                                    if len(v89) >= 2 {
                                        v89 = v89[1:len(v89)]
                                    }
                                default:
                                    switch "hello" {
                                    case "":
                                        var v103 int8 = 1
                                        _ = v103
                                    case "foo bar":
                                        if (!(("foo bar" + "foo bar") != ("\n\t" + ""))) {
                                            switch "a" {
                                            case "\n\t":
                                                if false {
                                                    var v104 int16 = 855
                                                    if (("hello" <= "\n\t") || (true && false)) {
                                                        fmt.Println(p1)
                                                    }
                                                    if (18446744073709551615 < 1) {
                                                        var v105 uint32 = 42
                                                        var v106 int64 = 78
                                                        p1 = (^(fi88))
                                                        _ = v105
                                                        _ = v106
                                                    } else {
                                                        _ = len(v89)
                                                        fmt.Println(fi84, p1, p0)
                                                        fmt.Println(fi88, v104)
                                                        var v107 uint16 = p0
                                                        _ = v107
                                                    }
                                                    _ = v104
                                                } else {
                                                    var v108 int64 = -999
                                                    v89 = append(v89, -2147483648)
                                                    if true {
                                                        v109 := GenS11{F0: 255, F1: (191 != 127), F2: 29}
                                                        v110 := GenS12{F0: (fi84 ^ 553), F1: 255, F2: 588}
                                                        var v111 uint32 = 245
                                                        v112 := GenS5{F0: -240}
                                                        p0 -= (p0 - p0)
                                                        _ = v109
                                                        _ = v110
                                                        _ = v111
                                                        _ = v112
                                                    }
                                                    _ = v108
                                                }
                                            case "a":
                                                switch "" {
                                                case "hello":
                                                    v89 = append(v89, 127)
                                                case "a":
                                                    switch fi88 {
                                                    case 4294967295:
                                                        p0 += (p0 % ((899) | 1))
                                                    }
                                                case "\n\t":
                                                    switch (("hello" + "\n\t") + "") {
                                                    case "a":
                                                        v113 := GenS13{F0: -1}
                                                        _ = v113
                                                    case "\n\t":
                                                        if (true || false) {
                                                            if len(v89) > 0 {
                                                                v89[0] = fi84
                                                            }
                                                            p1 += (fi88 ^ (fi88 % (((fi88 / ((1) | 1))) | 1)))
                                                            p1 += p1
                                                        }
                                                    case "":
                                                        if (2147483647 <= 42) {
                                                            v114 := GenS5{F0: -453}
                                                            v115 := GenS10{F0: 309, F1: -97, F2: 128}
                                                            v89 = append(v89, 9223372036854775807)
                                                            _ = v114
                                                            _ = v115
                                                        } else {
                                                            var v116 uint8 = 42
                                                            _ = v116
                                                        }
                                                    default:
                                                        switch (fi84 | fi84) {
                                                        case 128:
                                                            fmt.Println(fi84)
                                                        default:
                                                            if ("\n\t" == "hello") {
                                                                if len(v89) >= 2 {
                                                                    v89 = v89[1:len(v89)]
                                                                }
                                                                v117 := GenS12{F0: (^(fi84)), F1: 2, F2: 0}
                                                                _ = v117
                                                            }
                                                            var v118 float32 = (1.5 / 3.14159)
                                                            fmt.Println(v118, fi84)
                                                            var v119 int8 = 27
                                                            v120 := GenS13{F0: -1}
                                                            if false {
                                                                fmt.Println(fi88, v118, v118)
                                                                var v121 uint = 584
                                                                _ = v121
                                                            } else {
                                                                v122 := GenS14{F0: 6}
                                                                switch (fi88 << ((4294967295) & 63)) {
                                                                case 1:
                                                                    switch ("a" + "\n\t") {
                                                                    case "":
                                                                        p0 = 265
                                                                    case "\n\t":
                                                                        _ = len(v89)
                                                                    }
                                                                default:
                                                                    fmt.Println(v119, p0, p0)
                                                                }
                                                                p0 = (p0 << ((752) & 63))
                                                                _ = v122
                                                            }
                                                            _ = v118
                                                            _ = v119
                                                            _ = v120
                                                        }
                                                    }
                                                    v123 := GenS13{F0: 1}
                                                    _ = v123
                                                default:
                                                    if true {
                                                        if len(v89) >= 2 {
                                                            v89 = v89[1:len(v89)]
                                                        }
                                                    } else {
                                                        if true {
                                                            var v124 string = "\n\t"
                                                            _ = v124
                                                        }
                                                    }
                                                }
                                            case "hello":
                                                fmt.Println(p0)
                                            }
                                            var v125 uint8 = 42
                                            switch (fi88 * fi88) {
                                            case 9223372036854775807:
                                                var v126 int16 = 32767
                                                _ = v126
                                            case -9223372036854775808:
                                                v127 := GenS15{F0: 9223372036854775807}
                                                _ = v127
                                            case -752:
                                                v128 := GenS16{F0: 756, F1: 115, F2: (p0 | (p0 * 164))}
                                                _ = v128
                                            default:
                                                v129 := GenS17{F0: -1, F1: 1}
                                                _ = v129
                                            }
                                            _ = v125
                                        } else {
                                            p1 += -210
                                            var v130 uint64 = 385
                                            if (2 < 1) {
                                                v131 := GenS18{F0: p0}
                                                if (-556 <= -2147483648) {
                                                    fmt.Println(p1, p0)
                                                    fmt.Println(v130, fi88, p1)
                                                }
                                                _ = v131
                                            } else {
                                                var v132 int16 = -375
                                                if len(v89) > 0 {
                                                    v89[0] = p1
                                                }
                                                _ = v132
                                            }
                                            _ = v130
                                        }
                                    }
                                }
                                p1 -= fi84
                                var v133 int8 = 1
                                switch "foo bar" {
                                case "a":
                                    v134 := GenS15{F0: 1}
                                    _ = v134
                                }
                                _ = v133
                            }
                            if len(v89) > 0 {
                                v89[0] = (fi88 - fi88)
                            }
                            p0 += (^(p0))
                        }
                        if len(v89) >= 2 {
                            v89 = v89[1:len(v89)]
                        }
                        var v135 int = fi84
                        var v136 float32 = 1e10
                        var v137 float64 = (3.14159 - (0.1 / float64(v136)))
                        _ = fi88
                        _ = v135
                        _ = v136
                        _ = v137
                        _ = v89
                    }
                    p0 = p0
                    _ = v87
                } else {
                    fmt.Println(p1)
                    for fi138 := 0; fi138 < 3; fi138++ {
                        if true {
                            p0 = 0
                            if true {
                                switch "hello" {
                                case "\n\t":
                                    var v139 int8 = 42
                                    var v140 int16 = 985
                                    var v141 uint32 = 0
                                    _ = v139
                                    _ = v140
                                    _ = v141
                                case "hello":
                                    var v142 uint32 = 2
                                    _ = v142
                                default:
                                    fmt.Println(fi138, p1)
                                }
                            }
                            if false {
                                v143 := []int{(fi84 % (((fi84 ^ (-(p1)))) | 1)), (fi138 % ((fi84) | 1))}
                                var v144 bool = (-1 < 42)
                                _ = v143
                                _ = v144
                            }
                        } else {
                            if false {
                                fmt.Println(p1, fi84, fi138)
                            }
                        }
                        _ = fi138
                    }
                }
                fmt.Println(fi84, p0, p1)
            } else {
                v145 := []int{(p1 - p1), (p1 >> ((523) & 63)), -2147483648}
                _ = v145
            }
            _ = fi84
        }
        fmt.Println(p0, p0)
    } else {
        for fi146 := 0; fi146 < 14; fi146++ {
            fmt.Println(fi146)
            v147 := []int{(fi146 & p1), p1, fi146, fi146}
            for fi148 := 0; fi148 < 1; fi148++ {
                v149 := GenS10{F0: 2147483648, F1: 0, F2: 148}
                fmt.Println(p1, p0)
                p1 -= fi148
                _ = fi148
                _ = v149
            }
            _ = fi146
            _ = v147
        }
        if (-150 < 0) {
            for fi150 := 0; fi150 < 2; fi150++ {
                fmt.Println(fi150, fi150)
                var v151 int32 = 2147483647
                switch "foo bar" {
                case "hello":
                    if (!(false || true)) {
                        var v152 int32 = -623
                        if (!(!true)) {
                            switch (-(p1)) {
                            case 326:
                                if (!false) {
                                    if (!false) {
                                        v153 := []int{p1, p1, (-(p1))}
                                        var v154 string = "hello"
                                        _ = v153
                                        _ = v154
                                    }
                                    p0 = (p0 << ((4294967295) & 63))
                                } else {
                                    var v155 uint16 = p0
                                    if (false || (-1 > (fi150 / (((fi150 << ((620) & 63))) | 1)))) {
                                        if false {
                                            v156 := []int{fi150, p1}
                                            switch (fi150 - p1) {
                                            case 1:
                                                var v157 uint32 = 2147483648
                                                _ = v157
                                            case 2147483647:
                                                if (803 < 42) {
                                                    v156 = append(v156, fi150)
                                                }
                                            }
                                            _ = v156
                                        } else {
                                            v158 := GenS19{F0: 0, F1: 42, F2: 42}
                                            for fi159 := 0; fi159 < 12; fi159++ {
                                                if true {
                                                    p1 -= (fi159 ^ fi150)
                                                    fmt.Println(v155, fi150, p1)
                                                    var v160 int64 = -9223372036854775808
                                                    _ = v160
                                                } else {
                                                    v161 := GenS18{F0: p0}
                                                    v152 += v152
                                                    _ = v161
                                                }
                                                if (0.1 == 1.5) {
                                                    var v162 uint = 913
                                                    var v163 int64 = 1
                                                    _ = v162
                                                    _ = v163
                                                }
                                                switch 42 {
                                                case -1:
                                                    v164 := []int{(fi150 | fi159)}
                                                    _ = v164
                                                }
                                                _ = fi159
                                            }
                                            _ = v158
                                        }
                                        v152 = v152
                                        for fi165 := 0; fi165 < 8; fi165++ {
                                            var v166 uint32 = 774
                                            _ = fi165
                                            _ = v166
                                        }
                                    }
                                    _ = v155
                                }
                            case -2147483648:
                                v167 := []int{(-(p1)), int(v151), p1, fi150}
                                _ = v167
                            case 128:
                                p0 -= (p0 & 32767)
                            default:
                                switch (((((("\n\t" + "hello") + "hello") + "hello") + ((("hello" + ((("" + ("\n\t" + "foo bar")) + ("hello" + (("hello" + "") + (("a" + "hello") + "foo bar")))) + "a")) + (("a" + ("hello" + ("a" + ""))) + ("\n\t" + "a"))) + "hello")) + "\n\t") + "hello") {
                                case "a":
                                    var v168 uint = 18446744073709551615
                                    _ = v168
                                case "":
                                    v169 := GenS3{F0: 127, F1: 42, F2: (^(v151))}
                                    _ = v169
                                default:
                                    v170 := []int{(fi150 - fi150)}
                                    _ = v170
                                }
                            }
                            if false {
                                if (!true) {
                                    var v171 uint64 = 2
                                    var v172 int32 = 0
                                    if false {
                                        v172 -= (v172 | (-(v152)))
                                    }
                                    _ = v171
                                    _ = v172
                                } else {
                                    for fi173 := 0; fi173 < 18; fi173++ {
                                        v174 := []int{}
                                        switch "hello" {
                                        case "":
                                            var v175 int16 = 0
                                            _ = v175
                                        case "\n\t":
                                            v151 += (v152 >> ((65535) & 63))
                                        default:
                                            if (!true) {
                                                if ((!((-32768 <= int16(v151)) && false)) && true) {
                                                    v151 += 2147483647
                                                    if false {
                                                        fmt.Println(fi150)
                                                        v176 := GenS3{F0: 2, F1: -623, F2: 634}
                                                        _ = v176
                                                    } else {
                                                        p0 -= 32767
                                                        v152 += 897
                                                        var v177 uint64 = 18446744073709551615
                                                        _ = v177
                                                    }
                                                    var v178 bool = (!false)
                                                    _ = v178
                                                } else {
                                                    v179 := GenS20{F0: 200, F1: 139}
                                                    p1 = fi150
                                                    fmt.Println(v151, p0)
                                                    if (!false) {
                                                        var v180 float32 = 0.0
                                                        _ = v180
                                                    } else {
                                                        v151 = v152
                                                    }
                                                    fmt.Println(v152, fi150)
                                                    _ = v179
                                                }
                                            } else {
                                                var v181 float32 = 1e10
                                                if ((fi173 << ((0) & 63)) != fi173) {
                                                    v182 := GenS15{F0: -9223372036854775808}
                                                    if len(v174) >= 2 {
                                                        v174 = v174[1:len(v174)]
                                                    }
                                                    _ = v182
                                                } else {
                                                    if len(v174) >= 2 {
                                                        v174 = v174[1:len(v174)]
                                                    }
                                                }
                                                v183 := GenS10{F0: 2147483648, F1: 127, F2: 183}
                                                _ = v181
                                                _ = v183
                                            }
                                        }
                                        _ = fi173
                                        _ = v174
                                    }
                                    p1 += (fi150 - 0)
                                }
                                p0 = p0
                            }
                            v184 := GenS21{F0: 0}
                            _ = v184
                        }
                        _ = v152
                    }
                default:
                    fmt.Println(p0, p0)
                }
                _ = fi150
                _ = v151
            }
            for fi185 := 0; fi185 < 18; fi185++ {
                v186 := GenS14{F0: 183}
                p1 += fi185
                _ = fi185
                _ = v186
            }
        }
        for fi187 := 0; fi187 < 14; fi187++ {
            var v188 uint32 = 632
            v188 += 65535
            _ = fi187
            _ = v188
        }
    }
    for fi189 := 0; fi189 < 20; fi189++ {
        switch fi189 {
        case 9223372036854775807:
            if true {
                p0 -= p0
                fmt.Println(p0, p0)
            }
        case 0:
            p1 = p1
        case 206:
            p0 = 810
        default:
            var v190 int64 = -1
            _ = v190
        }
        var v191 int = 773
        v192 := []int{}
        _ = fi189
        _ = v191
        _ = v192
    }
    _ = p0
    _ = p1
    return 231
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]uint32)
    var v3 int8 = -113
    var v193 float32 = 0.1
    var v194 float32 = float32((0.0 * float64(float32(-2.25))))
    if (!true) {
        if (18446744073709551615 > 42) {
            var v195 uint64 = 119
            v3 = (v3 << ((4294967295) & 63))
            _ = v195
        } else {
            var v196 uint8 = 128
            if (!false) {
                v196 = 1
                v197 := GenS22{F0: genFn0()}
                _ = v197
            } else {
                v198 := GenS22{F0: -1}
                _ = v198
            }
            switch ("hello" + "a") {
            case "":
                if false {
                    fmt.Println(v3, v193)
                    var v199 int16 = 1
                    _ = v199
                } else {
                    var v200 int32 = 1
                    v196 = 42
                    _ = v200
                }
            }
            _ = v196
        }
    } else {
        v194 -= 3.14159
    }
    v201 := GenS23{F0: 747}
    var v202 int16 = 1000
    for fi203 := 0; fi203 < 20; fi203++ {
        for fi204 := 0; fi204 < 11; fi204++ {
            var v205 uint32 = 0
            if false {
                var v206 uint32 = (v205 | (^(v205)))
                fmt.Println(v202, fi204)
                v202 = genFn0()
                var v207 uint32 = 42
                switch fi203 {
                case -377:
                    var v208 uint16 = 428
                    _ = v208
                case 127:
                    var v209 int8 = (v3 - v3)
                    _ = v209
                case -683:
                    v210 := GenS10{F0: (v206 | 4294967295), F1: v3, F2: 59}
                    _ = v210
                default:
                    v3 = -92
                }
                _ = len(v1)
                if true {
                    v193 -= float32(((1e10 - 0.0) * (((3.14159 + 0.0) / ((((1e10 + float64(v193)) - -2.25) - ((float64(float32((0.1 + 1.5))) - 1.5) - 1.5)) / float64(((((0.1 - v193) - -2.25) * float32(-2.25)) / v193)))) * (1.5 + (0.0 * (-2.25 + 0.0))))))
                    var v211 uint = 2
                    v212 := []int{genFn1(uint16(v211), 1), 32767}
                    _ = v211
                    _ = v212
                }
                v205 = v205
                fmt.Println(v202)
                var v213 bool = true
                v214 := v1["\n\t"]
                _ = v206
                _ = v207
                _ = v213
                _ = v214
            }
            v202 -= (v202 >> ((189) & 63))
            _ = fi204
            _ = v205
        }
        var v215 int32 = -1
        if true {
            v202 += v202
            for fi216 := 0; fi216 < 19; fi216++ {
                if (4294967295 <= 261) {
                    v217 := GenS24{F0: 249, F1: 41, F2: 32768}
                    v202 -= (v202 % ((v202) | 1))
                    v218 := []int{(fi203 >> ((806) & 63)), (fi203 >> ((808) & 63))}
                    v219 := GenS25{F0: "", F1: true}
                    var v220 uint64 = 0
                    _ = v217
                    _ = v218
                    _ = v219
                    _ = v220
                }
                v3 -= v3
                v221 := GenS6{F0: (v215 ^ v215), F1: false}
                var v222 uint = 42
                if true {
                    var v223 string = ((("a" + (("a" + "foo bar") + "hello")) + "") + "hello")
                    v193 += v193
                    switch v223 {
                    case "hello":
                        fmt.Println(fi216)
                    case "a":
                        var v224 uint8 = 255
                        _ = v224
                    case "foo bar":
                        var v225 string = v223
                        _ = v225
                    }
                    _ = v223
                }
                _ = fi216
                _ = v221
                _ = v222
            }
            for fi226 := 0; fi226 < 6; fi226++ {
                switch (("hello" + ("\n\t" + ("" + "\n\t"))) + "") {
                case "\n\t":
                    var v227 uint8 = 0
                    _ = v227
                case "foo bar":
                    switch int(v215) {
                    case -1:
                        v193 = 1e10
                    default:
                        v228 := GenS20{F0: 424, F1: -605}
                        _ = v228
                    }
                case "":
                    v193 = (v194 - (v193 / (v193 - (v194 - ((3.14159 - float32(float64(v194))) * v193)))))
                }
                v3 -= -45
                v229 := GenS26{F0: -122, F1: genFn1(0, (fi226 & (fi203 & fi226))), F2: ((false || (true && true)) && false)}
                _ = fi226
                _ = v229
            }
        }
        _ = fi203
        _ = v215
    }
    if (!(!(42 >= 2))) {
        for fi230 := 0; fi230 < 1; fi230++ {
            switch 9223372036854775807 {
            case 128:
                if (false || true) {
                    v193 += (-2.25 * v193)
                    for fi231 := 0; fi231 < 14; fi231++ {
                        v232 := []int{fi230, 42, (fi230 + (fi230 & fi230))}
                        _ = fi231
                        _ = v232
                    }
                }
            case 0:
                fmt.Println(fi230)
            case -761:
                var v233 int64 = 1
                _ = v233
            default:
                for fi234 := 0; fi234 < 15; fi234++ {
                    switch ("a" + ("hello" + "\n\t")) {
                    case "foo bar":
                        v235 := []int{-2147483648, fi230, genFn1(761, fi230)}
                        var v236 uint32 = 2147483648
                        var v237 float32 = (-2.25 - v193)
                        switch fi234 {
                        case -1:
                            switch fi230 {
                            case 32768:
                                v235 = append(v235, (fi230 % ((127) | 1)))
                                _ = len(v235)
                            }
                        }
                        if (((float64(v237) + float64(v193)) * 1e10) > 3.14159) {
                            var v238 uint8 = 127
                            var v239 uint16 = 60
                            if false {
                                if false {
                                    var v240 uint8 = (v238 & v238)
                                    v240 -= uint8(v239)
                                    fmt.Println(v240, v239, v239)
                                    _ = v240
                                }
                            }
                            _ = v238
                            _ = v239
                        }
                        _ = v235
                        _ = v236
                        _ = v237
                    case "hello":
                        v241 := v1["hello"]
                        _ = v241
                    case "\n\t":
                        if false {
                            var v242 uint64 = 4294967295
                            _ = v242
                        } else {
                            fmt.Println(v202)
                        }
                    default:
                        switch ("\n\t" + "hello") {
                        case "":
                            v194 += (1.5 * v194)
                        default:
                            var v243 int = fi234
                            _ = v243
                        }
                    }
                    if false {
                        var v244 uint64 = 2
                        _ = v244
                    }
                    v245 := GenS27{F0: (v202 ^ -32768), F1: 877}
                    _ = fi234
                    _ = v245
                }
            }
            var v246 uint8 = 209
            _ = fi230
            _ = v246
        }
        _ = len(v1)
    } else {
        var v247 string = "hello"
        v248 := GenS28{F0: 42, F1: 42, F2: false}
        _ = v247
        _ = v248
    }
    var v249 uint32 = 11
    var v250 uint = 637
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v193:", v193)
    fmt.Println("v194:", v194)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v202:", v202)
    fmt.Println("v249:", v249)
    fmt.Println("v250:", v250)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v193
    _ = v194
    _ = v2
    _ = v201
    _ = v202
    _ = v249
    _ = v250
    _ = v3
}
