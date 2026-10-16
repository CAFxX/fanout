package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 uint64; F1 bool }
type GenS4 struct { F0 uint32; F1 int16 }
type GenS5 struct { F0 int16; F1 int32; F2 int64 }
type GenS6 struct { F0 uint8; F1 int16; F2 string }
type GenS7 struct { F0 uint16; F1 int32 }
type GenS8 struct { F0 int64; F1 int16; F2 int64 }
type GenS9 struct { F0 int32; F1 string; F2 int32 }
type GenS10 struct { F0 uint64; F1 uint32; F2 int32 }
func genFn0() uint16 {
    for fi6 := 0; fi6 < 12; fi6++ {
        var v7 uint64 = 1
        var v8 uint16 = 2
        v9 := []int{fi6, 9223372036854775807, (fi6 >> ((4294967295) & 63)), 1}
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 9; fi10++ {
        v11 := GenS1{F0: 0}
        v12 := []int{(fi10 | fi10), fi10, 4294967295, (-(fi10))}
        var v13 int64 = 0
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    return 1
}

func main() {
    var v1 float32 = 3.14159
    var v2 int64 = 0
    var v3 uint32 = 304
    var v4 int32 = 2147483647
    var v5 bool = (v3 < (v3 + v3))
    v14 := []int{2147483647, 4294967295}
    for fi15 := 0; fi15 < 20; fi15++ {
        fmt.Println(v3, v4, v1)
        _ = fi15
    }
    switch 472 {
    case -2147483648:
        var v16 string = ""
        _ = v16
    default:
        var v17 uint32 = 65535
        _ = v17
    }
    v18 := GenS1{F0: 585}
    v4 = 42
    if false {
        v14 = append(v14, 2147483647)
    } else {
        var v19 string = (("a" + "\n\t") + ("a" + ("foo bar" + "hello")))
        _ = v19
    }
    _ = len(v14)
    v14 = append(v14, -420)
    if false {
        switch 4294967295 {
        case 128:
            for fi20 := 0; fi20 < 3; fi20++ {
                v21 := GenS1{F0: 386}
                fmt.Println(fi20)
                _ = fi20
                _ = v21
            }
        case 289:
            v2 = (v2 ^ (-(v2)))
        default:
            for fi22 := 0; fi22 < 13; fi22++ {
                switch (("hello" + "") + ("\n\t" + (((("" + "hello") + ("a" + "\n\t")) + "") + "a"))) {
                case "a":
                    switch (fi22 & -9223372036854775808) {
                    case -609:
                        switch "foo bar" {
                        case "foo bar":
                            if v5 {
                                for fi23 := 0; fi23 < 4; fi23++ {
                                    var v24 bool = true
                                    v1 = 0.0
                                    _ = fi23
                                    _ = v24
                                }
                                if v5 {
                                    v4 = v4
                                    if (v1 <= 1.5) {
                                        if (732 >= 1) {
                                            if (-109 > 1) {
                                                if false {
                                                    v25 := GenS1{F0: 42}
                                                    if (v5 && ((((v3 ^ v3) >= 511) && v5) || v5)) {
                                                        var v26 bool = true
                                                        _ = v26
                                                    }
                                                    _ = v25
                                                }
                                            }
                                            var v27 uint16 = uint16(v3)
                                            _ = v27
                                        } else {
                                            switch fi22 {
                                            case 32767:
                                                if true {
                                                    v28 := GenS1{F0: 979}
                                                    v1 += v1
                                                    _ = v28
                                                }
                                            case -2147483648:
                                                fmt.Println(v4)
                                            default:
                                                v14 = append(v14, (fi22 / (((fi22 + fi22)) | 1)))
                                            }
                                            for fi29 := 0; fi29 < 6; fi29++ {
                                                var v30 int32 = 2147483647
                                                var v31 string = "a"
                                                switch fi29 {
                                                case 0:
                                                    var v32 int64 = (v2 / ((-365) | 1))
                                                    _ = v32
                                                case 309:
                                                    var v33 bool = v5
                                                    _ = v33
                                                default:
                                                    if (42 != 101) {
                                                        if (!v5) {
                                                            fmt.Println(v3, v31)
                                                            fmt.Println(v1)
                                                            if (1 <= -890) {
                                                                var v34 uint32 = (v3 & 93)
                                                                _ = v34
                                                            }
                                                            var v35 uint8 = 1
                                                            _ = v35
                                                        } else {
                                                            var v36 float32 = (3.14159 + v1)
                                                            var v37 float64 = float64(v1)
                                                            v36 = (3.14159 / (-2.25 * v1))
                                                            if (!v5) {
                                                                var v38 int32 = v30
                                                                _ = v38
                                                            }
                                                            if (v2 <= 0) {
                                                                if ((^(v3)) >= v3) {
                                                                    if ((1000 > 1000) || (!v5)) {
                                                                        _ = len(v14)
                                                                    }
                                                                    if true {
                                                                        v3 = (v3 ^ (v3 * 0))
                                                                        var v39 bool = (!v5)
                                                                        _ = v39
                                                                    } else {
                                                                        fmt.Println(fi29)
                                                                        var v40 int32 = 0
                                                                        v37 -= 0.0
                                                                        _ = v40
                                                                    }
                                                                }
                                                                v30 = (v30 * (v4 | 2147483647))
                                                            }
                                                            _ = v36
                                                            _ = v37
                                                        }
                                                    } else {
                                                        var v41 float32 = v1
                                                        var v42 uint16 = 255
                                                        var v43 string = v31
                                                        v3 = v3
                                                        switch fi29 {
                                                        case 128:
                                                            switch (fi22 % (((-(fi22))) | 1)) {
                                                            case -2147483648:
                                                                v30 += 2147483647
                                                            default:
                                                                fmt.Println(v30, v4)
                                                            }
                                                        case 758:
                                                            v44 := GenS1{F0: 1}
                                                            _ = v44
                                                        case 127:
                                                            var v45 uint32 = v3
                                                            _ = v45
                                                        }
                                                        v1 += 1e10
                                                        _ = v41
                                                        _ = v42
                                                        _ = v43
                                                    }
                                                }
                                                _ = fi29
                                                _ = v30
                                                _ = v31
                                            }
                                        }
                                        var v46 int64 = (v2 | (-(v2)))
                                        _ = v46
                                    } else {
                                        if len(v14) > 0 {
                                            v14[0] = 9223372036854775807
                                        }
                                        var v47 int64 = v2
                                        _ = v47
                                    }
                                    for fi48 := 0; fi48 < 3; fi48++ {
                                        var v49 uint = 183
                                        _ = fi48
                                        _ = v49
                                    }
                                }
                                if v5 {
                                    var v50 string = ""
                                    _ = v50
                                } else {
                                    switch ("hello" + "") {
                                    case "foo bar":
                                        var v51 float32 = v1
                                        _ = v51
                                    default:
                                        switch ("a" + ((("\n\t" + ("hello" + "a")) + (("a" + (("" + "") + "a")) + "hello")) + (((("\n\t" + "") + "a") + (("\n\t" + ((("foo bar" + "\n\t") + "hello") + "a")) + ("" + ("a" + "\n\t")))) + "\n\t"))) {
                                        case "\n\t":
                                            var v52 string = (("a" + "foo bar") + (("foo bar" + "a") + "hello"))
                                            _ = v52
                                        case "foo bar":
                                            v2 -= (v2 & v2)
                                        case "hello":
                                            v5 = ((366 < 1000) && (v1 < (v1 / float32(1e10))))
                                        default:
                                            v2 = v2
                                        }
                                    }
                                    fmt.Println(v4)
                                }
                            }
                        case "a":
                            switch "a" {
                            case "":
                                switch "hello" {
                                case "":
                                    fmt.Println(v5, v4)
                                default:
                                    _ = len(v14)
                                }
                            case "foo bar":
                                fmt.Println(v2)
                            }
                        default:
                            for fi53 := 0; fi53 < 10; fi53++ {
                                v4 -= (v4 >> ((837) & 63))
                                var v54 uint8 = 42
                                v55 := GenS1{F0: 2}
                                _ = fi53
                                _ = v54
                                _ = v55
                            }
                        }
                    case 32768:
                        for fi56 := 0; fi56 < 7; fi56++ {
                            v3 = (v3 + v3)
                            _ = fi56
                        }
                    case 0:
                        if len(v14) > 0 {
                            v14[0] = -995
                        }
                    }
                default:
                    fmt.Println(v2, v2, fi22)
                }
                for fi57 := 0; fi57 < 8; fi57++ {
                    v58 := GenS1{F0: 2}
                    if false {
                        v1 -= (1e10 * (((float32(1e10) * 1e10) - v1) * (v1 * (v1 - v1))))
                        if false {
                            var v59 int = -1
                            v1 += (0.0 * float32(0.1))
                            v60 := GenS1{F0: 0}
                            var v61 uint16 = 1
                            _ = v59
                            _ = v60
                            _ = v61
                        } else {
                            var v62 uint16 = genFn0()
                            v62 = (v62 * genFn0())
                            v14 = append(v14, fi22)
                            _ = v62
                        }
                    } else {
                        v4 += (v4 + v4)
                        switch "hello" {
                        case "\n\t":
                            switch (fi57 - fi57) {
                            case -2147483648:
                                v3 += (v3 & (v3 & (v3 & v3)))
                            case -791:
                                v4 -= (v4 / ((-1) | 1))
                            case -152:
                                v2 -= v2
                            }
                        case "":
                            v14 = append(v14, (fi22 >> ((255) & 63)))
                        case "hello":
                            var v63 bool = false
                            _ = v63
                        default:
                            switch (-(fi57)) {
                            case -204:
                                if v5 {
                                    switch "a" {
                                    case "a":
                                        if (v5 || v5) {
                                            if true {
                                                v2 -= (v2 << ((479) & 63))
                                                switch (("a" + "a") + ("hello" + "hello")) {
                                                case "a":
                                                    switch "hello" {
                                                    case "\n\t":
                                                        v64 := GenS1{F0: 0}
                                                        _ = v64
                                                    default:
                                                        var v65 uint16 = 199
                                                        _ = v65
                                                    }
                                                case "":
                                                    v5 = (!v5)
                                                case "foo bar":
                                                    if (!((v4 << ((614) & 63)) <= v4)) {
                                                        v66 := GenS2{F0: (fi22 >> ((255) & 63)), F1: 2147483647}
                                                        _ = v66
                                                    }
                                                }
                                                var v67 int8 = -60
                                                switch (fi22 >> ((310) & 63)) {
                                                case 4294967295:
                                                    if true {
                                                        var v68 string = "hello"
                                                        switch (^(fi57)) {
                                                        case 9223372036854775807:
                                                            var v69 uint8 = 2
                                                            _ = v69
                                                        case 993:
                                                            fmt.Println(v3)
                                                        default:
                                                            if len(v14) > 0 {
                                                                v14[0] = (fi22 * fi57)
                                                            }
                                                        }
                                                        var v70 uint32 = v3
                                                        _ = v68
                                                        _ = v70
                                                    } else {
                                                        var v71 uint = 4294967295
                                                        _ = v71
                                                    }
                                                case 1:
                                                    if (38 <= 412) {
                                                        if v5 {
                                                            if v5 {
                                                                v72 := GenS3{F0: 683, F1: v5}
                                                                v5 = (false && v5)
                                                                var v73 uint8 = 1
                                                                _ = v72
                                                                _ = v73
                                                            }
                                                            fmt.Println(fi22, v1)
                                                        }
                                                        var v74 int8 = 127
                                                        _ = len(v14)
                                                        _ = v74
                                                    }
                                                }
                                                _ = v67
                                            } else {
                                                var v75 int8 = 53
                                                v76 := GenS2{F0: (fi22 / (((fi22 + -1)) | 1)), F1: -2147483648}
                                                _ = len(v14)
                                                _ = v75
                                                _ = v76
                                            }
                                            if (255 > 255) {
                                                if (!false) {
                                                    var v77 int = 42
                                                    v77 += 2147483647
                                                    v77 = (fi57 | (v77 & v77))
                                                    _ = v77
                                                } else {
                                                    if len(v14) >= 2 {
                                                        v14 = v14[1:len(v14)]
                                                    }
                                                }
                                                var v78 uint8 = 0
                                                switch (fi57 ^ fi57) {
                                                case 32768:
                                                    var v79 bool = false
                                                    _ = v79
                                                default:
                                                    fmt.Println(v4, v78)
                                                }
                                                _ = v78
                                            } else {
                                                var v80 bool = v5
                                                if (((v1 + v1) / v1) != v1) {
                                                    var v81 int16 = 0
                                                    _ = v81
                                                } else {
                                                    var v82 bool = false
                                                    if false {
                                                        v82 = ((fi57 + int(v3)) <= (fi22 ^ (fi57 | (fi22 & fi57))))
                                                    } else {
                                                        var v83 int = (fi22 | (fi57 ^ 2147483647))
                                                        var v84 uint = 18446744073709551615
                                                        if false {
                                                            if len(v14) > 0 {
                                                                v14[0] = 2147483647
                                                            }
                                                            v85 := GenS4{F0: v3, F1: 478}
                                                            v1 += (v1 - (((-2.25 * 0.0) + v1) / 1.0))
                                                            _ = v85
                                                        }
                                                        _ = v83
                                                        _ = v84
                                                    }
                                                    v2 = (-(v2))
                                                    _ = v82
                                                }
                                                _ = v80
                                            }
                                            v5 = v5
                                        }
                                    case "":
                                        if len(v14) > 0 {
                                            v14[0] = fi22
                                        }
                                    default:
                                        _ = len(v14)
                                        v4 += v4
                                    }
                                    var v86 float64 = 0.0
                                    var v87 uint16 = 255
                                    _ = v86
                                    _ = v87
                                } else {
                                    var v88 uint16 = 255
                                    var v89 uint = 395
                                    if len(v14) > 0 {
                                        v14[0] = -2147483648
                                    }
                                    v90 := GenS5{F0: 915, F1: -2147483648, F2: (-(v2))}
                                    var v91 int = (fi22 >> ((v89) & 63))
                                    _ = v88
                                    _ = v89
                                    _ = v90
                                    _ = v91
                                }
                            case 356:
                                var v92 float32 = float32(((((-2.25 + (0.0 / 1.5)) + 1.5) / ((1.5 / 1.0) + 3.14159)) * (1.5 * (0.1 * 1e10))))
                                _ = v92
                            case 0:
                                if (int32(v3) == (v4 + int32(v3))) {
                                    v93 := GenS6{F0: 42, F1: -812, F2: ("a" + "a")}
                                    v1 += (v1 - float32(1.5))
                                    var v94 float32 = v1
                                    _ = v93
                                    _ = v94
                                } else {
                                    if (18446744073709551615 == 0) {
                                        v95 := GenS2{F0: -1, F1: fi22}
                                        switch (^(fi22)) {
                                        case 128:
                                            if true {
                                                v2 = int64(v4)
                                                var v96 string = "\n\t"
                                                _ = v96
                                            }
                                        case 433:
                                            var v97 int8 = 0
                                            _ = v97
                                        case 1:
                                            switch (("foo bar" + ("hello" + "foo bar")) + "foo bar") {
                                            case "hello":
                                                var v98 bool = (0 >= 503)
                                                _ = v98
                                            case "foo bar":
                                                if false {
                                                    if (false || (722 < -743)) {
                                                        var v99 string = ("foo bar" + "foo bar")
                                                        var v100 uint = 17
                                                        _ = v100
                                                        _ = v99
                                                    } else {
                                                        var v101 float64 = 3.14159
                                                        _ = v101
                                                    }
                                                    fmt.Println(v3)
                                                }
                                            default:
                                                var v102 uint8 = 2
                                                _ = v102
                                            }
                                        }
                                        _ = v95
                                    }
                                }
                            }
                        }
                        if true {
                            fmt.Println(v1, fi22)
                            if true {
                                v1 -= float32(float64(1.5))
                                if (!(true && v5)) {
                                    if len(v14) >= 2 {
                                        v14 = v14[1:len(v14)]
                                    }
                                } else {
                                    if false {
                                        if v5 {
                                            var v103 uint16 = 434
                                            switch "" {
                                            case "":
                                                switch (fi57 * int(v4)) {
                                                case -85:
                                                    v2 = v2
                                                case 927:
                                                    if len(v14) >= 2 {
                                                        v14 = v14[1:len(v14)]
                                                    }
                                                default:
                                                    v2 -= v2
                                                }
                                            case "foo bar":
                                                var v104 int32 = v4
                                                _ = v104
                                            case "\n\t":
                                                v105 := GenS2{F0: fi22, F1: 1}
                                                _ = v105
                                            default:
                                                fmt.Println(v3, v3)
                                            }
                                            var v106 string = ("foo bar" + "hello")
                                            fmt.Println(v3, v4)
                                            v14 = append(v14, -2147483648)
                                            _ = v103
                                            _ = v106
                                        }
                                        var v107 uint64 = 743
                                        _ = v107
                                    }
                                    if len(v14) >= 2 {
                                        v14 = v14[1:len(v14)]
                                    }
                                }
                                v5 = (2 <= 65535)
                            } else {
                                if ((v4 | (v4 - v4)) >= (v4 | 42)) {
                                    v108 := GenS7{F0: 99, F1: v4}
                                    var v109 uint8 = 127
                                    _ = v108
                                    _ = v109
                                }
                                switch fi22 {
                                case 1:
                                    fmt.Println(fi22, fi57, v2)
                                }
                                fmt.Println(v2, v4)
                            }
                        }
                    }
                    v5 = v5
                    switch "\n\t" {
                    case "foo bar":
                        if len(v14) >= 2 {
                            v14 = v14[1:len(v14)]
                        }
                    case "\n\t":
                        var v110 float64 = 1.5
                        _ = v110
                    }
                    _ = fi57
                    _ = v58
                }
                _ = fi22
            }
        }
        switch -2147483648 {
        case 32767:
            var v111 uint32 = v3
            _ = v111
        default:
            v14 = append(v14, 32768)
        }
        for fi112 := 0; fi112 < 5; fi112++ {
            v14 = append(v14, fi112)
            for fi113 := 0; fi113 < 12; fi113++ {
                v114 := GenS8{F0: v2, F1: 1, F2: v2}
                fmt.Println(v4)
                _ = fi113
                _ = v114
            }
            if (!true) {
                if (!v5) {
                    fmt.Println(fi112, v3, v4)
                    if len(v14) > 0 {
                        v14[0] = 128
                    }
                } else {
                    v115 := GenS6{F0: 0, F1: 42, F2: "a"}
                    var v116 float64 = 0.1
                    if true {
                        v116 = v116
                        fmt.Println(v5, v116, v2)
                    }
                    _ = v115
                    _ = v116
                }
            }
            _ = fi112
        }
    }
    var v117 int16 = 0
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    if v5 {
        var v118 uint64 = 18446744073709551615
        v3 += 42
        v118 = v118
        _ = v118
    }
    v119 := GenS9{F0: int32(v3), F1: "a", F2: v4}
    v120 := GenS2{F0: 127, F1: 128}
    switch "foo bar" {
    case "\n\t":
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
    case "hello":
        for fi121 := 0; fi121 < 6; fi121++ {
            v117 = (v117 | (v117 - v117))
            for fi122 := 0; fi122 < 16; fi122++ {
                var v123 int32 = 700
                _ = fi122
                _ = v123
            }
            var v124 int32 = (v4 / ((v4) | 1))
            _ = fi121
            _ = v124
        }
    default:
        fmt.Println(v117, v117, v5)
    }
    switch 42 {
    case 42:
        for fi125 := 0; fi125 < 4; fi125++ {
            if (fi125 <= (^(fi125))) {
                v126 := GenS10{F0: 221, F1: (v3 & v3), F2: 1}
                for fi127 := 0; fi127 < 6; fi127++ {
                    fmt.Println(v117)
                    v128 := GenS7{F0: 976, F1: v4}
                    _ = fi127
                    _ = v128
                }
                v129 := GenS6{F0: 42, F1: (v117 / ((v117) | 1)), F2: "foo bar"}
                _ = v126
                _ = v129
            }
            if (0 < -71) {
                for fi130 := 0; fi130 < 16; fi130++ {
                    var v131 uint32 = v3
                    _ = fi130
                    _ = v131
                }
                switch "a" {
                case "hello":
                    var v132 float64 = (1.5 * -2.25)
                    _ = v132
                }
                var v133 float32 = (v1 + v1)
                _ = v133
            } else {
                switch "\n\t" {
                case "\n\t":
                    var v134 float32 = v1
                    _ = v134
                case "foo bar":
                    v135 := GenS5{F0: -1, F1: v4, F2: v2}
                    _ = v135
                default:
                    v136 := GenS8{F0: (v2 + v2), F1: v117, F2: (v2 & v2)}
                    _ = v136
                }
            }
            v137 := GenS6{F0: 1, F1: (^(v117)), F2: "foo bar"}
            _ = fi125
            _ = v137
        }
    case 128:
        fmt.Println(v117)
    case 0:
        for fi138 := 0; fi138 < 17; fi138++ {
            var v139 int8 = 0
            _ = len(v14)
            var v140 int8 = (v139 & int8(v2))
            _ = fi138
            _ = v139
            _ = v140
        }
    default:
        var v141 uint8 = 2
        _ = v141
    }
    _ = len(v14)
    fmt.Println("v1:", v1)
    fmt.Println("v117:", v117)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v117
    _ = v119
    _ = v120
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
