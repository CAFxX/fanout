package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint; F1 int64; F2 uint32 }
type GenS4 struct { F0 int8; F1 uint32 }
type GenS5 struct { F0 int; F1 string }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 bool; F1 string }
type GenS8 struct { F0 int16; F1 int64 }
type GenS9 struct { F0 uint64; F1 int32; F2 int64 }
type GenS10 struct { F0 int32; F1 int32; F2 uint16 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 uint16; F1 uint32 }
type GenS13 struct { F0 uint32; F1 uint8 }
func genFn0(p0 int8, p1 int64) int64 {
    v5 := []int{}
    for fi6 := 0; fi6 < 2; fi6++ {
        v5 = append(v5, 1)
        if len(v5) > 0 {
            v5[0] = fi6
        }
        for fi7 := 0; fi7 < 13; fi7++ {
            if (2 != 345) {
                _ = len(v5)
                p1 = 413
                fmt.Println(fi7, fi6)
            } else {
                var v8 uint = 65535
                var v9 uint = 255
                v10 := GenS1{F0: -920, F1: (-(fi6))}
                _ = v10
                _ = v8
                _ = v9
            }
            var v11 float32 = 0.0
            var v12 string = "\n\t"
            switch v12 {
            case "":
                v13 := GenS1{F0: 1, F1: (fi6 % (((fi7 << ((2) & 63))) | 1))}
                _ = v13
            case "hello":
                var v14 float32 = float32(1.5)
                _ = v14
            }
            _ = fi7
            _ = v11
            _ = v12
        }
        _ = fi6
    }
    for fi15 := 0; fi15 < 1; fi15++ {
        if (128 < 255) {
            var v16 int = fi15
            var v17 int8 = p0
            if (!((!(1 >= 1)) && false)) {
                if ("" > "\n\t") {
                    var v18 uint = 0
                    if len(v5) > 0 {
                        v5[0] = -439
                    }
                    _ = v18
                }
            }
            _ = v16
            _ = v17
        } else {
            if ((1e10 - 1.5) < float32(3.14159)) {
                v19 := GenS2{F0: 1}
                fmt.Println(fi15, fi15)
                _ = v19
            } else {
                _ = len(v5)
                var v20 string = ""
                _ = v20
            }
            for fi21 := 0; fi21 < 11; fi21++ {
                var v22 int64 = (p1 % ((p1) | 1))
                if (((0.0 - 3.14159) - float32(0.0)) != (((((0.0 + 0.1) / 1.0) - 0.0) / 0.1) - 3.14159)) {
                    p1 = (p1 ^ (-(p1)))
                    p1 = p1
                    v22 = 0
                } else {
                    p1 = 9223372036854775807
                }
                _ = fi21
                _ = v22
            }
            var v23 float64 = 1e10
            _ = v23
        }
        if true {
            for fi24 := 0; fi24 < 18; fi24++ {
                p1 += (p1 + 0)
                v25 := GenS1{F0: 42, F1: 127}
                if true {
                    var v26 string = "\n\t"
                    var v27 int8 = (-(p0))
                    var v28 int16 = 42
                    _ = v26
                    _ = v27
                    _ = v28
                }
                var v29 int32 = 42
                p0 = p0
                switch -724 {
                case 128:
                    var v30 int8 = p0
                    _ = v30
                case 1:
                    var v31 bool = ((((float32((0.1 - float64((0.0 * -2.25)))) / 1.5) + (1e10 / -2.25)) <= 3.14159) && false)
                    _ = v31
                case 42:
                    v29 = (v29 & 42)
                }
                _ = fi24
                _ = v25
                _ = v29
            }
            p0 += (p0 >> ((255) & 63))
            var v32 int16 = 145
            _ = v32
        }
        if ((((1.5 / ((float32(float64(1.5)) * -2.25) / 0.1)) * 1.5) <= 3.14159) && true) {
            _ = len(v5)
        }
        _ = fi15
    }
    if true {
        v33 := GenS1{F0: 1, F1: -112}
        if false {
            p0 = p0
            p0 -= p0
        } else {
            fmt.Println(p1, p0)
            switch 1 {
            case 32767:
                for fi34 := 0; fi34 < 10; fi34++ {
                    v35 := GenS3{F0: 18446744073709551615, F1: p1, F2: 2147483648}
                    p1 = (p1 / ((p1) | 1))
                    _ = fi34
                    _ = v35
                }
            case 2147483647:
                if false {
                    switch "hello" {
                    case "hello":
                        switch ("hello" + "\n\t") {
                        case "a":
                            for fi36 := 0; fi36 < 15; fi36++ {
                                _ = len(v5)
                                fmt.Println(p0, p1)
                                var v37 int32 = -2147483648
                                _ = fi36
                                _ = v37
                            }
                        case "hello":
                            _ = len(v5)
                        }
                    }
                    v38 := GenS4{F0: (p0 | (p0 << ((302) & 63))), F1: 4294967295}
                    p1 += p1
                    _ = v38
                }
            }
        }
        if (!false) {
            switch 390 {
            case 171:
                v39 := GenS4{F0: 42, F1: 1}
                _ = v39
            case -521:
                for fi40 := 0; fi40 < 7; fi40++ {
                    _ = len(v5)
                    v41 := GenS3{F0: 177, F1: (^(p1)), F2: 793}
                    v42 := GenS5{F0: (fi40 ^ fi40), F1: ""}
                    _ = fi40
                    _ = v41
                    _ = v42
                }
            default:
                var v43 int32 = 1
                _ = v43
            }
            fmt.Println(p1)
            for fi44 := 0; fi44 < 18; fi44++ {
                var v45 uint8 = 255
                var v46 int = fi44
                _ = fi44
                _ = v45
                _ = v46
            }
        } else {
            switch -9223372036854775808 {
            case 845:
                v47 := GenS6{F0: -2147483648}
                _ = v47
            case 42:
                if false {
                    switch "" {
                    case "":
                        if (!true) {
                            v48 := GenS2{F0: 255}
                            var v49 uint64 = 4294967295
                            p1 = 226
                            _ = v48
                            _ = v49
                        } else {
                            fmt.Println(p1, p1)
                            v50 := GenS6{F0: 2147483647}
                            if false {
                                switch 0 {
                                case 32767:
                                    var v51 float64 = (-2.25 - 0.1)
                                    _ = v51
                                case 9223372036854775807:
                                    for fi52 := 0; fi52 < 13; fi52++ {
                                        p1 = (p1 & (p1 | -1))
                                        _ = len(v5)
                                        var v53 int16 = 263
                                        _ = fi52
                                        _ = v53
                                    }
                                case 985:
                                    v54 := GenS7{F0: true, F1: "a"}
                                    _ = v54
                                }
                            }
                            _ = v50
                        }
                    case "\n\t":
                        v55 := GenS8{F0: 0, F1: (p1 << ((255) & 63))}
                        _ = v55
                    case "a":
                        switch 2147483647 {
                        case 128:
                            if (true && (!(-2147483648 <= -1))) {
                                p0 = (p0 - 42)
                            }
                        case 741:
                            p0 = (p0 ^ p0)
                        default:
                            p0 -= 0
                        }
                    }
                } else {
                    fmt.Println(p0, p0)
                    p1 = p1
                }
            }
            v56 := GenS2{F0: 1}
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = v56
        }
        _ = v33
    }
    for fi57 := 0; fi57 < 12; fi57++ {
        var v58 float64 = -2.25
        if (32768 != 264) {
            if (!(true && (528 == 840))) {
                var v59 uint16 = 65535
                _ = v59
            } else {
                fmt.Println(p0, v58)
                if ((false && false) && false) {
                    if false {
                        switch "\n\t" {
                        case "foo bar":
                            for fi60 := 0; fi60 < 4; fi60++ {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                var v61 int8 = -76
                                v5 = append(v5, 32768)
                                _ = fi60
                                _ = v61
                            }
                        case "\n\t":
                            fmt.Println(v58)
                        case "":
                            var v62 int16 = -653
                            _ = v62
                        default:
                            if false {
                                fmt.Println(fi57)
                                switch "" {
                                case "hello":
                                    v58 = -2.25
                                case "":
                                    if (-41 == (^(fi57))) {
                                        for fi63 := 0; fi63 < 14; fi63++ {
                                            v64 := GenS9{F0: 2, F1: -2147483648, F2: p1}
                                            _ = fi63
                                            _ = v64
                                        }
                                        var v65 int = fi57
                                        _ = v65
                                    } else {
                                        var v66 float32 = 3.14159
                                        if (true && (42 == 203)) {
                                            if len(v5) > 0 {
                                                v5[0] = fi57
                                            }
                                        } else {
                                            v66 = v66
                                            fmt.Println(fi57, v66)
                                            switch (fi57 * 32768) {
                                            case 622:
                                                var v67 bool = (!true)
                                                _ = v67
                                            case 42:
                                                v58 += (v58 * v58)
                                            case 32768:
                                                switch (fi57 % (((fi57 / ((fi57) | 1))) | 1)) {
                                                case -9223372036854775808:
                                                    fmt.Println(p0, p0)
                                                case 920:
                                                    switch fi57 {
                                                    case 1:
                                                        if false {
                                                            p1 = (p1 * (p1 & p1))
                                                        } else {
                                                            v66 += ((v66 / 1e10) * v66)
                                                        }
                                                    }
                                                case 2147483647:
                                                    if (!true) {
                                                        switch ("a" + "hello") {
                                                        case "":
                                                            p1 -= p1
                                                        case "hello":
                                                            if false {
                                                                fmt.Println(v58)
                                                                for fi68 := 0; fi68 < 17; fi68++ {
                                                                    if false {
                                                                        v58 = -2.25
                                                                        v69 := GenS7{F0: (v58 == (0.1 / 1.0)), F1: (((("" + "a") + "a") + "a") + "\n\t")}
                                                                        _ = v69
                                                                    } else {
                                                                        var v70 bool = true
                                                                        var v71 float32 = v66
                                                                        _ = v70
                                                                        _ = v71
                                                                    }
                                                                    _ = fi68
                                                                }
                                                                if ((fi57 * fi57) >= (fi57 - fi57)) {
                                                                    switch "foo bar" {
                                                                    case "hello":
                                                                        fmt.Println(p1, fi57, p0)
                                                                    case "foo bar":
                                                                        fmt.Println(fi57)
                                                                    case "\n\t":
                                                                        var v72 int = fi57
                                                                        _ = v72
                                                                    }
                                                                    if len(v5) > 0 {
                                                                        v5[0] = (fi57 >> ((718) & 63))
                                                                    }
                                                                    fmt.Println(v66, p0)
                                                                }
                                                            }
                                                        case "\n\t":
                                                            var v73 uint16 = 42
                                                            _ = v73
                                                        default:
                                                            v74 := GenS8{F0: -32768, F1: 90}
                                                            _ = v74
                                                        }
                                                    }
                                                }
                                            }
                                        }
                                        p0 -= p0
                                        _ = v66
                                    }
                                case "\n\t":
                                    v5 = append(v5, fi57)
                                default:
                                    fmt.Println(v58, p1, v58)
                                }
                                var v75 int8 = (p0 >> ((65535) & 63))
                                _ = v75
                            }
                        }
                        var v76 float32 = float32(float64(3.14159))
                        _ = v76
                    }
                    if false {
                        var v77 uint64 = 768
                        _ = v77
                    } else {
                        v78 := GenS8{F0: 798, F1: (p1 + (p1 | int64(p0)))}
                        var v79 uint32 = 1
                        _ = v78
                        _ = v79
                    }
                }
            }
        }
        _ = fi57
        _ = v58
    }
    _ = p0
    _ = p1
    _ = v5
    return 0
}
func genFn1(p0 int32, p1 int32, p2 int32) int {
    var v80 int64 = 1
    var v81 uint32 = 1
    var v82 int64 = (v80 * (v80 >> ((2) & 63)))
    v83 := []int{4294967295, 32767}
    var v84 int16 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v80
    _ = v81
    _ = v82
    _ = v83
    _ = v84
    return -2147483648
}

func main() {
    v1 := []uint32{2147483648, 42, 4294967295, 4294967295}
    v2 := make(map[int]bool)
    var v3 int64 = -957
    var v4 float32 = 0.0
    v85 := GenS10{F0: 1, F1: -877, F2: 0}
    if false {
        var v86 bool = (!false)
        switch (("" + ("a" + "foo bar")) + "\n\t") {
        case "a":
            var v87 uint64 = 42
            _ = v87
        case "\n\t":
            v4 += (v4 - float32(0.0))
        case "":
            var v88 uint = 98
            _ = v88
        default:
            if false {
                v89 := v2[-147]
                for fi90 := 0; fi90 < 7; fi90++ {
                    switch "foo bar" {
                    case "a":
                        var v91 uint32 = 2147483648
                        _ = v91
                    default:
                        _ = len(v2)
                    }
                    v86 = false
                    _ = fi90
                }
                if len(v1) > 0 {
                    v1[0] = 2
                }
                _ = v89
            } else {
                v92 := GenS3{F0: 65535, F1: -1, F2: 471}
                v93 := GenS8{F0: 32767, F1: (v3 * -1)}
                var v94 string = "\n\t"
                _ = v92
                _ = v93
                _ = v94
            }
        }
        for fi95 := 0; fi95 < 11; fi95++ {
            var v96 string = "a"
            var v97 int64 = v3
            for fi98 := 0; fi98 < 16; fi98++ {
                v4 -= v4
                v96 = (("\n\t" + (v96 + "a")) + v96)
                v86 = (v4 != -2.25)
                _ = fi98
            }
            _ = fi95
            _ = v96
            _ = v97
        }
        _ = v86
    }
    var v99 int32 = -1
    if ((^(v99)) > v99) {
        switch genFn1(910, v99, 1) {
        case 196:
            v3 -= (v3 - (v3 / ((-9223372036854775808) | 1)))
        default:
            switch "" {
            case "\n\t":
                if (1 == 42) {
                    v3 -= genFn0(0, (^(v3)))
                    var v100 uint64 = 2
                    v101 := GenS2{F0: 0}
                    _ = v100
                    _ = v101
                }
            case "a":
                v3 = (v3 >> ((0) & 63))
            case "foo bar":
                for fi102 := 0; fi102 < 2; fi102++ {
                    var v103 int32 = (v99 * v99)
                    _ = fi102
                    _ = v103
                }
            }
        }
        v104 := v2[2147483647]
        v3 = v3
        _ = v104
    } else {
        switch 32768 {
        case -9223372036854775808:
            v4 = 0.1
        case 0:
            delete(v2, -93)
        case 9223372036854775807:
            delete(v2, 128)
        default:
            v105 := GenS11{F0: -1}
            _ = v105
        }
    }
    v106 := GenS1{F0: v99, F1: 4294967295}
    if true {
        v4 += 3.14159
    } else {
        var v107 uint16 = 1
        _ = v107
    }
    if false {
        for fi108 := 0; fi108 < 13; fi108++ {
            for fi109 := 0; fi109 < 16; fi109++ {
                v110 := GenS3{F0: 1, F1: (v3 << ((761) & 63)), F2: 252}
                _ = fi109
                _ = v110
            }
            fmt.Println(v3, v99)
            v4 = v4
            _ = fi108
        }
    } else {
        for fi111 := 0; fi111 < 18; fi111++ {
            fmt.Println(v99)
            if (1 > 65535) {
                switch ("hello" + "foo bar") {
                case "hello":
                    switch ("\n\t" + "foo bar") {
                    case "foo bar":
                        for fi112 := 0; fi112 < 19; fi112++ {
                            var v113 float32 = 1.5
                            if (((v113 - v4) + v4) >= float32((float64(((float32(((3.14159 / 1.0) + float64(0.0))) * v4) + 3.14159)) + float64(v4)))) {
                                var v114 uint32 = 2147483647
                                v115 := GenS1{F0: (-(v99)), F1: (fi111 % (((fi111 & (fi111 ^ (fi111 - int(v3))))) | 1))}
                                switch ("a" + "") {
                                case "\n\t":
                                    var v116 int8 = 49
                                    _ = v116
                                case "foo bar":
                                    v99 = v99
                                case "hello":
                                    switch genFn1((v99 << ((18446744073709551615) & 63)), v99, 42) {
                                    case -137:
                                        switch ("\n\t" + "\n\t") {
                                        case "foo bar":
                                            var v117 float64 = (float64(v113) * 1e10)
                                            _ = v117
                                        case "hello":
                                            v2[(^(fi111))] = true
                                        case "":
                                            var v118 string = ("a" + "\n\t")
                                            _ = v118
                                        }
                                    case 4294967295:
                                        var v119 int = (fi112 % ((4294967295) | 1))
                                        _ = v119
                                    case 260:
                                        v113 -= v113
                                    default:
                                        v113 = v113
                                    }
                                default:
                                    fmt.Println(v4)
                                }
                                _ = v114
                                _ = v115
                            }
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = fi112
                            _ = v113
                        }
                    case "a":
                        for fi120 := 0; fi120 < 2; fi120++ {
                            var v121 uint16 = 1
                            if ((1e10 - ((v4 + 0.1) * (((float32(1.5) / (v4 + 1.5)) * v4) + -2.25))) != (1e10 - -2.25)) {
                                var v122 int = int(v121)
                                v3 -= v3
                                fmt.Println(fi120, fi111)
                                _ = v122
                            } else {
                                fmt.Println(v99, v121)
                            }
                            _ = fi120
                            _ = v121
                        }
                    default:
                        v123 := GenS2{F0: 2}
                        _ = v123
                    }
                case "foo bar":
                    v4 = v4
                default:
                    var v124 string = (("\n\t" + "a") + "\n\t")
                    _ = v124
                }
                for fi125 := 0; fi125 < 6; fi125++ {
                    v3 -= genFn0(1, genFn0(107, v3))
                    v126 := GenS12{F0: 349, F1: 863}
                    var v127 uint8 = 128
                    fmt.Println(fi111, fi125)
                    v3 = v3
                    _ = fi125
                    _ = v126
                    _ = v127
                }
                switch "foo bar" {
                case "a":
                    var v128 uint64 = 581
                    _ = v128
                case "":
                    var v129 uint32 = 378
                    _ = v129
                case "hello":
                    for fi130 := 0; fi130 < 17; fi130++ {
                        if (genFn0(-1, genFn0(13, v3)) >= 0) {
                            v99 += int32(v3)
                        }
                        var v131 uint = 669
                        _ = fi130
                        _ = v131
                    }
                }
            } else {
                v2[fi111] = (-474 < v99)
                for fi132 := 0; fi132 < 10; fi132++ {
                    v99 += 0
                    v133 := GenS2{F0: 2}
                    _ = fi132
                    _ = v133
                }
            }
            _ = fi111
        }
    }
    var v134 uint64 = 18446744073709551615
    if false {
        for fi135 := 0; fi135 < 5; fi135++ {
            for fi136 := 0; fi136 < 7; fi136++ {
                if true {
                    switch fi135 {
                    case 32768:
                        var v137 float64 = (1e10 * 1.5)
                        _ = v137
                    case -20:
                        switch "" {
                        case "\n\t":
                            v138 := GenS9{F0: v134, F1: v99, F2: v3}
                            _ = v138
                        case "hello":
                            v139 := GenS4{F0: 1, F1: 471}
                            _ = v139
                        }
                    default:
                        var v140 float32 = float32(0.1)
                        _ = v140
                    }
                    switch (((((("a" + (("foo bar" + "hello") + "\n\t")) + "\n\t") + "foo bar") + (("" + "") + "foo bar")) + "a") + "") {
                    case "foo bar":
                        if true {
                            var v141 uint64 = 18446744073709551615
                            _ = v141
                        } else {
                            fmt.Println(v4, fi136)
                        }
                    case "":
                        fmt.Println(v3, v99, fi136)
                    }
                }
                v142 := GenS12{F0: 2, F1: 1}
                v143 := GenS13{F0: 0, F1: 0}
                _ = fi136
                _ = v142
                _ = v143
            }
            var v144 uint8 = 2
            _ = fi135
            _ = v144
        }
    }
    for fi145 := 0; fi145 < 11; fi145++ {
        var v146 float64 = (1.5 + 0.1)
        if true {
            for fi147 := 0; fi147 < 16; fi147++ {
                var v148 int16 = 160
                var v149 uint32 = 2147483647
                v150 := v2[(fi147 * fi145)]
                _ = fi147
                _ = v148
                _ = v149
                _ = v150
            }
        }
        switch fi145 {
        case 32767:
            switch "a" {
            case "foo bar":
                v3 = (v3 / ((-83) | 1))
            case "a":
                if true {
                    v151 := v2[genFn1(-2147483648, 0, v99)]
                    fmt.Println(v99, fi145)
                    if (!true) {
                        v152 := GenS10{F0: (v99 >> ((255) & 63)), F1: v99, F2: 65535}
                        _ = len(v1)
                        _ = v152
                    }
                    _ = v151
                } else {
                    var v153 uint8 = 60
                    v153 += (v153 | (v153 * (v153 ^ (v153 % ((v153) | 1)))))
                    _ = v153
                }
            case "":
                var v154 int32 = v99
                _ = v154
            }
        case -2147483648:
            v155 := GenS5{F0: fi145, F1: ("\n\t" + ((("hello" + ("a" + "hello")) + "foo bar") + ""))}
            _ = v155
        case 42:
            fmt.Println(v146)
        default:
            if (0 >= 2) {
                v156 := GenS3{F0: 0, F1: (v3 + (v3 ^ 9223372036854775807)), F2: 42}
                v2[genFn1((v99 ^ 2147483647), (v99 * -1), 0)] = (0 > 100)
                v157 := GenS5{F0: 42, F1: "foo bar"}
                _ = v156
                _ = v157
            } else {
                for fi158 := 0; fi158 < 12; fi158++ {
                    var v159 int8 = -113
                    v3 = (^(v3))
                    v1 = append(v1, 4294967295)
                    _ = fi158
                    _ = v159
                }
            }
        }
        _ = fi145
        _ = v146
    }
    _ = len(v2)
    var v160 uint8 = 128
    switch genFn1((-(v99)), v99, (v99 / ((0) | 1))) {
    case -1:
        var v161 int64 = genFn0(18, (v3 / ((v3) | 1)))
        _ = v161
    case 275:
        v162 := GenS5{F0: 2147483647, F1: "a"}
        _ = v162
    case 2147483647:
        v4 -= (3.14159 - v4)
    }
    if len(v1) > 0 {
        v1[0] = 65535
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v134:", v134)
    fmt.Println("v160:", v160)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v106
    _ = v134
    _ = v160
    _ = v2
    _ = v3
    _ = v4
    _ = v85
    _ = v99
}
