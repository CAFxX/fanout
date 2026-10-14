package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 int32 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 int8; F1 int8 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint; F1 uint64 }
type GenS7 struct { F0 string; F1 uint; F2 uint }
type GenS8 struct { F0 int; F1 int8 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 int8; F1 uint64; F2 int }
type GenS11 struct { F0 uint16; F1 bool; F2 int32 }
type GenS12 struct { F0 int16; F1 int; F2 bool }
type GenS13 struct { F0 uint32; F1 int64; F2 int8 }
type GenS14 struct { F0 uint16; F1 string }
type GenS15 struct { F0 uint64; F1 int16 }
type GenS16 struct { F0 int; F1 uint8; F2 uint16 }
type GenS17 struct { F0 uint16 }
type GenS18 struct { F0 uint32; F1 uint16 }
type GenS19 struct { F0 int64 }
type GenS20 struct { F0 string; F1 string }
type GenS21 struct { F0 uint8 }
type GenS22 struct { F0 string }
type GenS23 struct { F0 int64; F1 string }
type GenS24 struct { F0 uint }
type GenS25 struct { F0 uint8; F1 uint8; F2 int }
type GenS26 struct { F0 int; F1 bool }
type GenS27 struct { F0 int16 }
type GenS28 struct { F0 bool; F1 uint8; F2 int64 }
type GenS29 struct { F0 int16; F1 int16; F2 int }
func genFn0(p0 int16) uint64 {
    var v6 bool = ((false || (117 < 42)) || true)
    for fi7 := 0; fi7 < 10; fi7++ {
        if v6 {
            if (!(42 < (p0 + p0))) {
                v6 = (1 != 448)
                v8 := []int{-317}
                _ = v8
            } else {
                var v9 int8 = 34
                _ = v9
            }
        }
        var v10 bool = v6
        _ = fi7
        _ = v10
    }
    _ = p0
    _ = v6
    return 0
}
func genFn1(p0 uint16, p1 int, p2 bool) int64 {
    fmt.Println(p2)
    for fi11 := 0; fi11 < 9; fi11++ {
        switch "\n\t" {
        case "a":
            v12 := []int{(p1 % ((997) | 1)), 128, fi11, fi11}
            _ = v12
        default:
            if p2 {
                p2 = true
                switch ("hello" + "hello") {
                case "foo bar":
                    var v13 int8 = -1
                    _ = v13
                case "":
                    var v14 uint8 = 42
                    _ = v14
                case "a":
                    for fi15 := 0; fi15 < 6; fi15++ {
                        v16 := []int{fi11, (fi11 >> ((970) & 63)), fi11}
                        fmt.Println(p0, fi11, p1)
                        p2 = false
                        _ = fi15
                        _ = v16
                    }
                }
            }
        }
        for fi17 := 0; fi17 < 13; fi17++ {
            p2 = false
            switch fi17 {
            case 0:
                switch "a" {
                case "a":
                    switch "\n\t" {
                    case "":
                        var v18 uint64 = 4294967295
                        _ = v18
                    case "\n\t":
                        if ("hello" > "") {
                            v19 := GenS1{F0: fi11, F1: 9223372036854775807, F2: 1}
                            var v20 string = "foo bar"
                            _ = v19
                            _ = v20
                        }
                    case "foo bar":
                        p2 = p2
                    }
                case "hello":
                    v21 := []int{(p1 + (fi11 / (((p1 % ((fi11) | 1))) | 1)))}
                    _ = v21
                case "\n\t":
                    var v22 string = ""
                    _ = v22
                default:
                    p2 = p2
                }
            default:
                v23 := []int{}
                _ = v23
            }
            _ = fi17
        }
        fmt.Println(p1)
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = p2
    return -9223372036854775808
}

func main() {
    var v1 int = 4294967295
    v2 := make(map[string]string)
    var v3 uint32 = 497
    var v4 int32 = 0
    v5 := make(map[int]uint32)
    fmt.Println(v3)
    var v24 uint = 2
    v1 = v1
    v4 = (-(v4))
    v25 := GenS1{F0: (v1 << (((v24 / ((v24) | 1))) & 63)), F1: 42, F2: (v4 * (v4 - v4))}
    if false {
        if (int64(v4) <= 0) {
            v4 = -1
            fmt.Println(v1)
            if false {
                v24 += v24
                for fi26 := 0; fi26 < 19; fi26++ {
                    var v27 int32 = 2147483647
                    if false {
                        for fi28 := 0; fi28 < 12; fi28++ {
                            if (147 == 128) {
                                switch (v1 & (fi28 % ((128) | 1))) {
                                case 4294967295:
                                    if (!true) {
                                        switch "hello" {
                                        case "a":
                                            v1 = (v1 >> ((v24) & 63))
                                        case "foo bar":
                                            v29 := GenS1{F0: fi26, F1: 42, F2: (v4 / ((-1) | 1))}
                                            _ = v29
                                        case "hello":
                                            var v30 int64 = 0
                                            _ = v30
                                        default:
                                            if (!true) {
                                                v31 := GenS1{F0: int(v3), F1: 218, F2: v4}
                                                _ = v31
                                            } else {
                                                var v32 float32 = 1.5
                                                var v33 int64 = 454
                                                v34 := GenS2{F0: -32768, F1: -128}
                                                _ = v32
                                                _ = v33
                                                _ = v34
                                            }
                                        }
                                        switch (fi28 / ((fi28) | 1)) {
                                        case 32767:
                                            v24 += v24
                                        case 1:
                                            v35 := v5[(fi28 + 127)]
                                            _ = v35
                                        case -1:
                                            v36 := GenS2{F0: -32768, F1: 0}
                                            _ = v36
                                        default:
                                            switch "hello" {
                                            case "foo bar":
                                                fmt.Println(fi26, v4)
                                            default:
                                                switch v1 {
                                                case -2147483648:
                                                    var v37 float64 = ((((0.1 / (0.0 + (1.5 + 0.0))) - float64(float32((1.5 + -2.25)))) * 3.14159) / 3.14159)
                                                    _ = v37
                                                case 9223372036854775807:
                                                    var v38 string = ""
                                                    _ = v38
                                                default:
                                                    v39 := GenS1{F0: v1, F1: -1, F2: v4}
                                                    _ = v39
                                                }
                                            }
                                        }
                                        v4 = v4
                                    } else {
                                        var v40 bool = (-1 > 127)
                                        _ = v40
                                    }
                                case -414:
                                    _ = len(v2)
                                case -437:
                                    v27 -= v4
                                }
                                if true {
                                    v24 -= v24
                                    if true {
                                        var v41 uint = v24
                                        fmt.Println(fi28, fi28)
                                        _ = v41
                                    } else {
                                        if ((-2.25 * 1e10) == 1e10) {
                                            if (false && ((fi26 % ((fi28) | 1)) < (v1 >> (((v24 ^ (^(v24)))) & 63)))) {
                                                var v42 int64 = int64(fi26)
                                                _ = v42
                                            }
                                            v4 = (v27 + v4)
                                        } else {
                                            v43 := GenS3{F0: 42, F1: "a"}
                                            v44 := v5[(v1 | -9223372036854775808)]
                                            switch "a" {
                                            case "foo bar":
                                                v24 += v24
                                            case "a":
                                                if true {
                                                    v45 := []int{-9223372036854775808, -590}
                                                    _ = v45
                                                }
                                            default:
                                                v27 -= (^(v4))
                                            }
                                            _ = v43
                                            _ = v44
                                        }
                                        v27 -= (v4 * v27)
                                    }
                                    var v46 int64 = -9223372036854775808
                                    _ = v46
                                } else {
                                    var v47 uint64 = 203
                                    _ = v47
                                }
                            }
                            v27 = 2147483647
                            if (false && ((v1 % ((fi26) | 1)) <= fi26)) {
                                v4 = (v4 << ((v24) & 63))
                            }
                            _ = fi28
                        }
                        if true {
                            if ((18446744073709551615 > 2) && (true || false)) {
                                for fi48 := 0; fi48 < 17; fi48++ {
                                    var v49 int32 = (v27 - 390)
                                    var v50 int8 = -36
                                    _ = fi48
                                    _ = v49
                                    _ = v50
                                }
                                var v51 int32 = v27
                                _ = v51
                            } else {
                                var v52 uint = 4294967295
                                var v53 int = -2147483648
                                delete(v2, "hello")
                                _ = v52
                                _ = v53
                            }
                            v5[32768] = (v3 + v3)
                        } else {
                            v24 -= 2
                            fmt.Println(v4)
                        }
                    } else {
                        switch v1 {
                        case 4294967295:
                            _ = len(v5)
                        }
                        v3 -= v3
                    }
                    if false {
                        var v54 uint32 = v3
                        v3 = 2147483647
                        var v55 int8 = -99
                        _ = v54
                        _ = v55
                    }
                    _ = fi26
                    _ = v27
                }
                for fi56 := 0; fi56 < 15; fi56++ {
                    for fi57 := 0; fi57 < 20; fi57++ {
                        v58 := []int{fi57, fi57, (v1 | v1), (fi56 ^ 1)}
                        delete(v2, ("a" + (("" + ("" + (("" + "") + "foo bar"))) + ("foo bar" + "a"))))
                        fmt.Println(fi56, fi57)
                        if false {
                            v59 := GenS4{F0: 127, F1: 80}
                            v60 := GenS3{F0: 68, F1: ""}
                            _ = v59
                            _ = v60
                        } else {
                            if true {
                                var v61 int32 = (v4 >> (((v24 | v24)) & 63))
                                switch (-(fi57)) {
                                case 99:
                                    if (genFn0(-1) <= 18446744073709551615) {
                                        var v62 float64 = (1.5 - 1e10)
                                        _ = v62
                                    } else {
                                        v3 = (^(v3))
                                        if ("a" <= "a") {
                                            v63 := GenS1{F0: (fi56 - 674), F1: 42, F2: 1}
                                            _ = v63
                                        }
                                        fmt.Println(fi57, v61)
                                    }
                                case 127:
                                    _ = len(v58)
                                case 4294967295:
                                    if (65535 <= 255) {
                                        if true {
                                            var v64 int8 = -1
                                            v65 := GenS5{F0: (v3 / ((v3) | 1))}
                                            _ = v64
                                            _ = v65
                                        }
                                    } else {
                                        v1 = fi56
                                        var v66 uint8 = 6
                                        v66 += 0
                                        _ = v66
                                    }
                                }
                                _ = v61
                            }
                            var v67 int64 = int64(v4)
                            v68 := GenS6{F0: (v24 * (v24 + (v24 - v24))), F1: 2}
                            _ = v67
                            _ = v68
                        }
                        var v69 int32 = v4
                        _ = fi57
                        _ = v58
                        _ = v69
                    }
                    _ = fi56
                }
            } else {
                if (!false) {
                    var v70 int8 = 100
                    v71 := GenS2{F0: -1, F1: v70}
                    _ = v70
                    _ = v71
                } else {
                    v4 -= v4
                    v24 = 4294967295
                }
            }
        }
        switch ("\n\t" + "foo bar") {
        case "":
            var v72 float64 = (-2.25 * 1.5)
            _ = v72
        default:
            var v73 uint16 = 65535
            _ = v73
        }
        fmt.Println(v4)
    } else {
        switch "foo bar" {
        case "hello":
            if (0 >= uint16(v3)) {
                v74 := GenS7{F0: ("a" + "hello"), F1: (v24 % (((v24 * (v24 * v24))) | 1)), F2: (^(v24))}
                _ = v74
            } else {
                if true {
                    for fi75 := 0; fi75 < 10; fi75++ {
                        v24 += (^(v24))
                        fmt.Println(v3, v3)
                        _ = fi75
                    }
                    v4 += (^(v4))
                    v3 = uint32(v24)
                }
                if true {
                    switch 42 {
                    case 32767:
                        var v76 string = "a"
                        _ = v76
                    }
                } else {
                    switch "a" {
                    case "":
                        switch v1 {
                        case -2147483648:
                            v4 += v4
                        case 9223372036854775807:
                            v3 += (v3 - v3)
                        default:
                            for fi77 := 0; fi77 < 3; fi77++ {
                                fmt.Println(v1, fi77)
                                if false {
                                    fmt.Println(v4)
                                    var v78 int = fi77
                                    var v79 int = (-(v78))
                                    _ = v78
                                    _ = v79
                                }
                                var v80 string = ("a" + "hello")
                                _ = fi77
                                _ = v80
                            }
                        }
                    default:
                        _ = len(v5)
                    }
                    v81 := v5[(v1 & v1)]
                    _ = v81
                }
                var v82 uint32 = (v3 * 962)
                _ = v82
            }
        case "":
            v83 := GenS2{F0: -1, F1: 1}
            _ = v83
        }
        if false {
            switch (v1 << ((946) & 63)) {
            case -1:
                v84 := []int{-1, 910, (-(v1))}
                _ = v84
            }
            fmt.Println(v24, v24, v1)
        } else {
            if (v1 == (v1 + (-(v1)))) {
                var v85 float32 = ((float32(0.0) - (-2.25 - 0.0)) * (3.14159 + ((1e10 / float32(-2.25)) / 1.0)))
                switch -137 {
                case 9223372036854775807:
                    var v86 int16 = 32767
                    _ = v86
                case -1:
                    if true {
                        if false {
                            var v87 int = (v1 + v1)
                            switch (("hello" + (("a" + "\n\t") + "\n\t")) + "") {
                            case "foo bar":
                                v24 += 1
                            case "\n\t":
                                v85 += ((v85 * (float32(1.5) / ((0.0 * v85) / (float32(0.0) / (float32(-2.25) - ((float32(1e10) + v85) + v85)))))) / float32(1.5))
                            case "hello":
                                for fi88 := 0; fi88 < 5; fi88++ {
                                    v24 = v24
                                    v89 := GenS1{F0: (v1 + (-(fi88))), F1: 0, F2: -2147483648}
                                    _ = fi88
                                    _ = v89
                                }
                            default:
                                if (((-2.25 - v85) == 0.0) || (!true)) {
                                    v90 := GenS8{F0: v1, F1: 127}
                                    _ = v90
                                } else {
                                    v24 = (v24 % (((v24 & 312)) | 1))
                                    if (-1 > 841) {
                                        switch "a" {
                                        case "\n\t":
                                            if true {
                                                v91 := GenS6{F0: 156, F1: 2}
                                                switch 4294967295 {
                                                case 42:
                                                    var v92 uint8 = 246
                                                    _ = v92
                                                case 440:
                                                    var v93 uint32 = (^(v3))
                                                    _ = v93
                                                default:
                                                    var v94 int8 = -128
                                                    _ = v94
                                                }
                                                var v95 int32 = (^(v4))
                                                _ = v91
                                                _ = v95
                                            }
                                        case "foo bar":
                                            if (v87 < v1) {
                                                if true {
                                                    var v96 int64 = -261
                                                    switch (v1 | 0) {
                                                    case -250:
                                                        for fi97 := 0; fi97 < 14; fi97++ {
                                                            fmt.Println(v87)
                                                            switch (v1 % (((v87 * v1)) | 1)) {
                                                            case 0:
                                                                var v98 uint32 = (v3 | (v3 << (((v24 & 2)) & 63)))
                                                                _ = v98
                                                            case 4294967295:
                                                                v99 := []int{(fi97 >> ((963) & 63)), v1, (-(fi97)), (v87 - -683)}
                                                                _ = v99
                                                            case -1:
                                                                var v100 int32 = (v4 | v4)
                                                                _ = v100
                                                            default:
                                                                fmt.Println(v96, fi97)
                                                            }
                                                            switch "" {
                                                            case "hello":
                                                                v87 = (fi97 << (((v24 & (v24 & v24))) & 63))
                                                            case "":
                                                                switch ("a" + ("" + ("a" + "foo bar"))) {
                                                                case "\n\t":
                                                                    var v101 uint16 = 32767
                                                                    _ = v101
                                                                case "a":
                                                                    if false {
                                                                        v1 += int(v96)
                                                                        var v102 int32 = (-(v4))
                                                                        v96 -= v96
                                                                        _ = v102
                                                                    } else {
                                                                        var v103 float64 = 3.14159
                                                                        var v104 uint64 = genFn0(1000)
                                                                        for fi105 := 0; fi105 < 2; fi105++ {
                                                                            var v106 string = "hello"
                                                                            _ = fi105
                                                                            _ = v106
                                                                        }
                                                                        _ = v103
                                                                        _ = v104
                                                                    }
                                                                default:
                                                                    fmt.Println(v96, v1)
                                                                }
                                                            default:
                                                                var v107 uint = (v24 & v24)
                                                                _ = v107
                                                            }
                                                            _ = fi97
                                                        }
                                                    case -9223372036854775808:
                                                        var v108 int16 = 42
                                                        _ = v108
                                                    case -128:
                                                        if (!true) {
                                                            v109 := GenS9{F0: 100}
                                                            delete(v5, v87)
                                                            v110 := []int{4294967295, v87}
                                                            _ = v109
                                                            _ = v110
                                                        } else {
                                                            fmt.Println(v96, v24)
                                                            var v111 bool = ((v3 % ((v3) | 1)) <= (v3 & 4294967295))
                                                            _ = v111
                                                        }
                                                    }
                                                    var v112 uint8 = 128
                                                    _ = v112
                                                    _ = v96
                                                }
                                            } else {
                                                for fi113 := 0; fi113 < 5; fi113++ {
                                                    var v114 float64 = 0.0
                                                    var v115 int8 = 1
                                                    _ = fi113
                                                    _ = v114
                                                    _ = v115
                                                }
                                            }
                                        }
                                    }
                                    if false {
                                        var v116 uint = 0
                                        _ = v116
                                    }
                                }
                            }
                            _ = v87
                        }
                    }
                case 42:
                    if (((float64(v85) / 1e10) > 0.1) || (!false)) {
                        v117 := GenS6{F0: v24, F1: 42}
                        var v118 int32 = v4
                        if true {
                            if (3.14159 == v85) {
                                v119 := []int{1, 4294967295}
                                switch "a" {
                                case "\n\t":
                                    switch v1 {
                                    case -9223372036854775808:
                                        v119 = append(v119, v1)
                                    case 0:
                                        var v120 int = 1
                                        _ = v120
                                    default:
                                        switch (v1 % (((v1 % (((v1 + 9223372036854775807)) | 1))) | 1)) {
                                        case -33:
                                            v121 := GenS10{F0: 127, F1: 918, F2: (-(v1))}
                                            _ = v121
                                        case -919:
                                            v24 = (v24 - (^(v24)))
                                        case 32767:
                                            if (!(1 > -41)) {
                                                var v122 int64 = 9223372036854775807
                                                if ((v1 % (((v1 & 1)) | 1)) > (v1 ^ (v1 - -505))) {
                                                    var v123 float64 = (3.14159 + float64(-2.25))
                                                    _ = v123
                                                } else {
                                                    if (!(104 >= 255)) {
                                                        if len(v119) > 0 {
                                                            v119[0] = (v1 + 2147483647)
                                                        }
                                                        var v124 uint32 = (v3 ^ 42)
                                                        delete(v5, v1)
                                                        _ = v124
                                                    } else {
                                                        v125 := GenS11{F0: 42, F1: (-1 >= 0), F2: (v4 - (v4 & (v4 ^ v118)))}
                                                        if false {
                                                            var v126 int32 = v118
                                                            v127 := GenS6{F0: uint(v3), F1: 920}
                                                            _ = v126
                                                            _ = v127
                                                        }
                                                        var v128 int8 = 0
                                                        _ = v125
                                                        _ = v128
                                                    }
                                                }
                                                if true {
                                                    _ = len(v2)
                                                    var v129 uint16 = 0
                                                    _ = v129
                                                }
                                                _ = v122
                                            }
                                        }
                                    }
                                case "foo bar":
                                    if true {
                                        switch "\n\t" {
                                        case "a":
                                            var v130 uint64 = 0
                                            _ = v130
                                        case "hello":
                                            v3 -= v3
                                        default:
                                            fmt.Println(v3, v3)
                                        }
                                        var v131 uint = (^(v24))
                                        var v132 uint8 = 248
                                        _ = v131
                                        _ = v132
                                    } else {
                                        var v133 float64 = (3.14159 + -2.25)
                                        for fi134 := 0; fi134 < 14; fi134++ {
                                            switch "\n\t" {
                                            case "hello":
                                                var v135 uint32 = v3
                                                _ = v135
                                            default:
                                                v136 := GenS12{F0: 1, F1: -9223372036854775808, F2: false}
                                                _ = v136
                                            }
                                            if false {
                                                fmt.Println(v1, v85)
                                                v133 = float64(v85)
                                            } else {
                                                v4 += (v4 - v4)
                                                if true {
                                                    var v137 bool = true
                                                    for fi138 := 0; fi138 < 2; fi138++ {
                                                        fmt.Println(v137, v85)
                                                        var v139 string = ("\n\t" + "foo bar")
                                                        if v137 {
                                                            switch (v139 + "a") {
                                                            case "\n\t":
                                                                if v137 {
                                                                    v140 := GenS13{F0: 4294967295, F1: 9223372036854775807, F2: int8(v24)}
                                                                    _ = v140
                                                                }
                                                            case "a":
                                                                v141 := GenS11{F0: 691, F1: v137, F2: (v4 ^ (v4 ^ v4))}
                                                                _ = v141
                                                            default:
                                                                var v142 uint16 = 42
                                                                _ = v142
                                                            }
                                                        }
                                                        _ = fi138
                                                        _ = v139
                                                    }
                                                    for fi143 := 0; fi143 < 7; fi143++ {
                                                        switch "" {
                                                        case "":
                                                            v144 := GenS14{F0: 725, F1: "a"}
                                                            _ = v144
                                                        case "\n\t":
                                                            if true {
                                                                v145 := GenS4{F0: -93, F1: 42}
                                                                _ = v145
                                                            }
                                                        default:
                                                            fmt.Println(fi134, v118)
                                                        }
                                                        if (float32((-2.25 - ((((0.0 - 0.1) + float64(-2.25)) + (v133 / float64((0.1 + v85)))) / 1.5))) != (float32(v133) / float32(v133))) {
                                                            v146 := GenS7{F0: "", F1: (v24 << ((737) & 63)), F2: v24}
                                                            v24 = (v24 | (v24 & (v24 + (^(v24)))))
                                                            var v147 float32 = -2.25
                                                            _ = v146
                                                            _ = v147
                                                        } else {
                                                            var v148 int64 = -1
                                                            _ = v148
                                                        }
                                                        var v149 int16 = 1000
                                                        _ = fi143
                                                        _ = v149
                                                    }
                                                    _ = v137
                                                }
                                            }
                                            v1 = v1
                                            _ = fi134
                                        }
                                        _ = v133
                                    }
                                case "hello":
                                    if true {
                                        v85 = v85
                                        var v150 uint64 = 705
                                        for fi151 := 0; fi151 < 4; fi151++ {
                                            switch "hello" {
                                            case "foo bar":
                                                var v152 int16 = -32768
                                                _ = v152
                                            case "a":
                                                for fi153 := 0; fi153 < 4; fi153++ {
                                                    switch -9223372036854775808 {
                                                    case 113:
                                                        if false {
                                                            fmt.Println(v150, v118, v118)
                                                            v1 = 9223372036854775807
                                                        }
                                                    case -2147483648:
                                                        var v154 int64 = -193
                                                        _ = v154
                                                    default:
                                                        fmt.Println(v150, fi153, v1)
                                                    }
                                                    var v155 int32 = v118
                                                    v150 += v150
                                                    _ = fi153
                                                    _ = v155
                                                }
                                            case "hello":
                                                if true {
                                                    _ = len(v119)
                                                    v156 := GenS15{F0: 1, F1: 307}
                                                    _ = v156
                                                } else {
                                                    if len(v119) > 0 {
                                                        v119[0] = 553
                                                    }
                                                }
                                            }
                                            v3 += v3
                                            if ("hello" >= "hello") {
                                                for fi157 := 0; fi157 < 9; fi157++ {
                                                    var v158 uint32 = v3
                                                    if len(v119) > 0 {
                                                        v119[0] = fi151
                                                    }
                                                    _ = fi157
                                                    _ = v158
                                                }
                                                v1 += v1
                                            }
                                            _ = fi151
                                        }
                                        _ = v150
                                    } else {
                                        if false {
                                            v4 = v4
                                            fmt.Println(v24, v24)
                                            var v159 uint32 = v3
                                            _ = v159
                                        } else {
                                            if true {
                                                if false {
                                                    v24 -= (v24 ^ (^(v24)))
                                                    fmt.Println(v118, v118, v118)
                                                    switch (v1 % ((v1) | 1)) {
                                                    case 0:
                                                        var v160 uint64 = 18446744073709551615
                                                        _ = v160
                                                    case 9223372036854775807:
                                                        if true {
                                                            switch v1 {
                                                            case 4294967295:
                                                                v161 := GenS14{F0: 731, F1: ("" + "hello")}
                                                                _ = v161
                                                            case -802:
                                                                var v162 bool = ((v118 ^ (v4 << ((v24) & 63))) < v118)
                                                                _ = v162
                                                            case 32768:
                                                                var v163 float32 = 1.5
                                                                _ = v163
                                                            }
                                                            v164 := GenS12{F0: 1000, F1: (v1 % (((v1 - v1)) | 1)), F2: ((0 != genFn1(1, (v1 & v1), (2 == 42))) || false)}
                                                            _ = v164
                                                        } else {
                                                            var v165 int32 = -166
                                                            fmt.Println(v85, v4)
                                                            _ = len(v2)
                                                            _ = v165
                                                        }
                                                    case 170:
                                                        v166 := GenS8{F0: v1, F1: -128}
                                                        _ = v166
                                                    default:
                                                        v4 -= v4
                                                    }
                                                }
                                                v4 -= (-(v4))
                                                if ("\n\t" != "") {
                                                    var v167 int16 = -32768
                                                    _ = v167
                                                }
                                            } else {
                                                _ = len(v119)
                                                switch "" {
                                                case "foo bar":
                                                    v85 = 3.14159
                                                case "hello":
                                                    fmt.Println(v3, v3, v118)
                                                }
                                            }
                                            fmt.Println(v3)
                                            switch "hello" {
                                            case "\n\t":
                                                var v168 int8 = -1
                                                _ = v168
                                            case "hello":
                                                for fi169 := 0; fi169 < 15; fi169++ {
                                                    v1 -= 0
                                                    _ = fi169
                                                }
                                            case "a":
                                                var v170 int64 = 9223372036854775807
                                                _ = v170
                                            default:
                                                var v171 uint32 = (^(v3))
                                                _ = v171
                                            }
                                        }
                                    }
                                default:
                                    for fi172 := 0; fi172 < 16; fi172++ {
                                        switch 42 {
                                        case 4294967295:
                                            switch (((("hello" + "hello") + ("" + "\n\t")) + ("hello" + "a")) + ("foo bar" + "a")) {
                                            case "":
                                                var v173 int8 = 100
                                                _ = v173
                                            case "\n\t":
                                                v4 += v118
                                            }
                                        case 9223372036854775807:
                                            fmt.Println(v4)
                                        case -1:
                                            for fi174 := 0; fi174 < 7; fi174++ {
                                                v4 -= v118
                                                v4 = v118
                                                _ = fi174
                                            }
                                        }
                                        if true {
                                            v175 := GenS4{F0: 127, F1: 42}
                                            if true {
                                                if true {
                                                    fmt.Println(v85, fi172, v85)
                                                    fmt.Println(v1, v118, v1)
                                                }
                                                for fi176 := 0; fi176 < 19; fi176++ {
                                                    v177 := v2["hello"]
                                                    _ = fi176
                                                    _ = v177
                                                }
                                                for fi178 := 0; fi178 < 7; fi178++ {
                                                    delete(v5, (fi178 % ((350) | 1)))
                                                    _ = fi178
                                                }
                                            } else {
                                                var v179 int32 = (v118 >> ((v24) & 63))
                                                v118 -= (v4 >> (((v24 * v24)) & 63))
                                                _ = v179
                                            }
                                            fmt.Println(v118, v4)
                                            _ = v175
                                        } else {
                                            for fi180 := 0; fi180 < 14; fi180++ {
                                                v85 += (1.5 + (v85 + v85))
                                                if true {
                                                    v181 := GenS16{F0: (v1 % (((fi180 | (v1 ^ -787))) | 1)), F1: 2, F2: 255}
                                                    _ = v181
                                                } else {
                                                    v1 = -9223372036854775808
                                                    var v182 float32 = (((v85 - ((v85 * (0.0 - v85)) * v85)) + v85) / (v85 * v85))
                                                    v119 = append(v119, (v1 + (fi172 % (((fi172 + (v1 * (^(fi180))))) | 1))))
                                                    if (!true) {
                                                        v85 -= (v182 / 3.14159)
                                                        v183 := v5[9223372036854775807]
                                                        v4 -= -424
                                                        _ = v183
                                                    } else {
                                                        if len(v119) > 0 {
                                                            v119[0] = (fi172 / ((v1) | 1))
                                                        }
                                                        if len(v119) >= 2 {
                                                            v119 = v119[1:len(v119)]
                                                        }
                                                        switch 4294967295 {
                                                        case 1:
                                                            v4 = (v118 | v118)
                                                            if (472 < 192) {
                                                                switch fi172 {
                                                                case -9223372036854775808:
                                                                    if true {
                                                                        if true {
                                                                            if (1e10 >= -2.25) {
                                                                                var v184 uint64 = 4294967295
                                                                                v5[v1] = v3
                                                                                if false {
                                                                                    var v185 int = fi172
                                                                                    _ = v185
                                                                                } else {
                                                                                    v118 = v4
                                                                                }
                                                                                _ = v184
                                                                            }
                                                                            if len(v119) >= 2 {
                                                                                v119 = v119[1:len(v119)]
                                                                            }
                                                                            fmt.Println(fi180, v24, v1)
                                                                        } else {
                                                                            if false {
                                                                                v186 := GenS13{F0: v3, F1: 0, F2: -78}
                                                                                _ = v186
                                                                            } else {
                                                                                _ = len(v5)
                                                                                _ = len(v119)
                                                                            }
                                                                            fmt.Println(fi180, v118)
                                                                            var v187 uint64 = 0
                                                                            _ = v187
                                                                        }
                                                                        if false {
                                                                            var v188 int64 = -9223372036854775808
                                                                            _ = v188
                                                                        }
                                                                    }
                                                                case -130:
                                                                    var v189 bool = true
                                                                    _ = v189
                                                                }
                                                                if len(v119) > 0 {
                                                                    v119[0] = 32767
                                                                }
                                                            } else {
                                                                var v190 string = "\n\t"
                                                                fmt.Println(v3, v1, fi180)
                                                                _ = v190
                                                            }
                                                            delete(v5, fi180)
                                                        }
                                                    }
                                                    v85 = (v182 * 0.0)
                                                    _ = v182
                                                }
                                                _ = fi180
                                            }
                                            if len(v119) > 0 {
                                                v119[0] = (v1 >> (((v24 >> (((v24 % (((^(v24))) | 1))) & 63))) & 63))
                                            }
                                            var v191 uint32 = v3
                                            _ = v191
                                        }
                                        _ = fi172
                                    }
                                }
                                _ = v119
                            } else {
                                if ((!("a" <= "foo bar")) && (8 == 127)) {
                                    if (!false) {
                                        if false {
                                            v192 := GenS7{F0: "\n\t", F1: v24, F2: v24}
                                            _ = v192
                                        } else {
                                            var v193 int32 = v118
                                            if false {
                                                var v194 int64 = -9223372036854775808
                                                if true {
                                                    var v195 float64 = ((0.0 / 1.5) - 1e10)
                                                    var v196 uint = v24
                                                    var v197 int64 = v194
                                                    _ = v195
                                                    _ = v196
                                                    _ = v197
                                                } else {
                                                    v85 += 1.5
                                                    switch ("hello" + ("foo bar" + "foo bar")) {
                                                    case "foo bar":
                                                        v3 = (v3 << ((v24) & 63))
                                                    case "hello":
                                                        for fi198 := 0; fi198 < 10; fi198++ {
                                                            v199 := GenS12{F0: -32768, F1: fi198, F2: (v85 != v85)}
                                                            fmt.Println(v118)
                                                            var v200 string = "hello"
                                                            _ = fi198
                                                            _ = v199
                                                            _ = v200
                                                        }
                                                    default:
                                                        v118 += 0
                                                    }
                                                }
                                                _ = v194
                                            }
                                            for fi201 := 0; fi201 < 13; fi201++ {
                                                if (!false) {
                                                    var v202 uint8 = 64
                                                    for fi203 := 0; fi203 < 7; fi203++ {
                                                        var v204 uint16 = 91
                                                        fmt.Println(fi201, fi203)
                                                        v1 -= fi201
                                                        _ = fi203
                                                        _ = v204
                                                    }
                                                    _ = v202
                                                }
                                                v205 := v2["hello"]
                                                _ = fi201
                                                _ = v205
                                            }
                                            _ = v193
                                        }
                                        v206 := GenS17{F0: 400}
                                        v207 := []int{-1}
                                        _ = v206
                                        _ = v207
                                    } else {
                                        var v208 float32 = v85
                                        for fi209 := 0; fi209 < 3; fi209++ {
                                            for fi210 := 0; fi210 < 7; fi210++ {
                                                v211 := GenS14{F0: 0, F1: "a"}
                                                v212 := []int{fi209, fi210, 1}
                                                var v213 bool = true
                                                _ = fi210
                                                _ = v211
                                                _ = v212
                                                _ = v213
                                            }
                                            if false {
                                                v118 = v118
                                            } else {
                                                v214 := []int{1, (fi209 - (v1 % (((^(v1))) | 1))), 4294967295}
                                                switch "hello" {
                                                case "a":
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        var v215 uint16 = 0
                                                        _ = v215
                                                    case "a":
                                                        fmt.Println(v3, v118)
                                                    case "":
                                                        for fi216 := 0; fi216 < 14; fi216++ {
                                                            switch 128 {
                                                            case 810:
                                                                v217 := GenS18{F0: v3, F1: 255}
                                                                _ = v217
                                                            case 595:
                                                                if ((!(133 < 416)) && (127 > 0)) {
                                                                    v218 := GenS19{F0: genFn1(0, 128, (false || false))}
                                                                    v208 -= ((3.14159 + v85) + (v85 / v85))
                                                                    var v219 uint = v24
                                                                    _ = v218
                                                                    _ = v219
                                                                }
                                                            case 0:
                                                                switch ("" + "") {
                                                                case "foo bar":
                                                                    fmt.Println(fi209, fi209)
                                                                default:
                                                                    var v220 bool = false
                                                                    _ = v220
                                                                }
                                                            default:
                                                                v85 += 0.0
                                                            }
                                                            v85 -= 1e10
                                                            v3 -= 2147483647
                                                            _ = fi216
                                                        }
                                                    default:
                                                        v221 := GenS15{F0: genFn0(-32768), F1: 1000}
                                                        _ = v221
                                                    }
                                                case "hello":
                                                    v3 = v3
                                                case "foo bar":
                                                    var v222 uint64 = 1
                                                    _ = v222
                                                }
                                                switch (^(v1)) {
                                                case 0:
                                                    for fi223 := 0; fi223 < 6; fi223++ {
                                                        v224 := GenS20{F0: "foo bar", F1: "hello"}
                                                        _ = fi223
                                                        _ = v224
                                                    }
                                                case 1:
                                                    fmt.Println(v118, v118, v1)
                                                }
                                                _ = v214
                                            }
                                            _ = fi209
                                        }
                                        _ = v208
                                    }
                                }
                            }
                            if true {
                                v1 -= v1
                                for fi225 := 0; fi225 < 16; fi225++ {
                                    var v226 string = ""
                                    fmt.Println(v226)
                                    v24 += 18446744073709551615
                                    _ = fi225
                                    _ = v226
                                }
                                v227 := GenS21{F0: 2}
                                _ = v227
                            } else {
                                var v228 uint8 = 128
                                for fi229 := 0; fi229 < 8; fi229++ {
                                    for fi230 := 0; fi230 < 9; fi230++ {
                                        v231 := []int{}
                                        var v232 int32 = 0
                                        _ = fi230
                                        _ = v231
                                        _ = v232
                                    }
                                    _ = fi229
                                }
                                if (!false) {
                                    if (1 <= 32767) {
                                        if (true && false) {
                                            v118 -= (v118 % ((-2147483648) | 1))
                                            v233 := GenS15{F0: 2, F1: -1}
                                            v3 = (v3 - 76)
                                            _ = v233
                                        } else {
                                            fmt.Println(v3, v4)
                                            for fi234 := 0; fi234 < 9; fi234++ {
                                                if (false || false) {
                                                    v235 := v5[fi234]
                                                    _ = v235
                                                } else {
                                                    v236 := GenS22{F0: ("" + "\n\t")}
                                                    var v237 uint16 = 0
                                                    delete(v5, v1)
                                                    _ = v236
                                                    _ = v237
                                                }
                                                _ = fi234
                                            }
                                            switch "hello" {
                                            case "\n\t":
                                                v228 = (v228 >> ((754) & 63))
                                            case "foo bar":
                                                v238 := GenS14{F0: 807, F1: "hello"}
                                                _ = v238
                                            case "":
                                                if true {
                                                    var v239 int16 = -941
                                                    _ = v239
                                                } else {
                                                    if (534 < genFn0(-1)) {
                                                        for fi240 := 0; fi240 < 11; fi240++ {
                                                            v241 := GenS8{F0: v1, F1: -127}
                                                            if (42 <= 2) {
                                                                var v242 int = (-(v1))
                                                                switch "hello" {
                                                                case "":
                                                                    var v243 int64 = 985
                                                                    _ = v243
                                                                default:
                                                                    var v244 float64 = 1e10
                                                                    _ = v244
                                                                }
                                                                switch fi240 {
                                                                case 32768:
                                                                    for fi245 := 0; fi245 < 6; fi245++ {
                                                                        v246 := v2["\n\t"]
                                                                        if true {
                                                                            if false {
                                                                                var v247 uint16 = 1
                                                                                if ((v246 + v246) <= "\n\t") {
                                                                                    var v248 int8 = 0
                                                                                    v249 := GenS23{F0: -9223372036854775808, F1: ("foo bar" + "")}
                                                                                    _ = v248
                                                                                    _ = v249
                                                                                } else {
                                                                                    v246 = v246
                                                                                    v1 += int(v3)
                                                                                }
                                                                                v250 := GenS18{F0: v3, F1: (v247 >> (((v24 * (v24 + (v24 ^ (v24 << ((v24) & 63)))))) & 63))}
                                                                                _ = v247
                                                                                _ = v250
                                                                            }
                                                                        }
                                                                        if false {
                                                                            _ = len(v5)
                                                                        } else {
                                                                            v251 := GenS18{F0: (v3 | uint32(v228)), F1: 50}
                                                                            switch v246 {
                                                                            case "foo bar":
                                                                                v252 := GenS24{F0: v24}
                                                                                _ = v252
                                                                            case "a":
                                                                                v253 := []int{(v242 - (v1 & fi245)), 127}
                                                                                _ = v253
                                                                            default:
                                                                                v228 -= (v228 << (((^(v24))) & 63))
                                                                            }
                                                                            v4 -= 1
                                                                            _ = v251
                                                                        }
                                                                        _ = fi245
                                                                        _ = v246
                                                                    }
                                                                case 386:
                                                                    fmt.Println(v228, v1, v242)
                                                                }
                                                                _ = v242
                                                            }
                                                            for fi254 := 0; fi254 < 18; fi254++ {
                                                                var v255 int = 4294967295
                                                                var v256 bool = true
                                                                fmt.Println(v228, v4)
                                                                v257 := GenS12{F0: 1000, F1: v1, F2: (18446744073709551615 != 42)}
                                                                _ = fi254
                                                                _ = v255
                                                                _ = v256
                                                                _ = v257
                                                            }
                                                            _ = fi240
                                                            _ = v241
                                                        }
                                                        if true {
                                                            var v258 uint32 = 1
                                                            v259 := GenS25{F0: (^(v228)), F1: v228, F2: (v1 / (((v1 % (((v1 - v1)) | 1))) | 1))}
                                                            _ = v258
                                                            _ = v259
                                                        } else {
                                                            var v260 float64 = 3.14159
                                                            v118 += (v4 * v118)
                                                            _ = v260
                                                        }
                                                        v118 = (v118 / ((v4) | 1))
                                                    }
                                                }
                                            }
                                        }
                                        v118 -= (v4 - (^(v4)))
                                        for fi261 := 0; fi261 < 9; fi261++ {
                                            var v262 uint32 = (v3 + 2)
                                            _ = fi261
                                            _ = v262
                                        }
                                    }
                                }
                                _ = v228
                            }
                        }
                        _ = v117
                        _ = v118
                    }
                default:
                    fmt.Println(v24, v4)
                }
                _ = v85
            }
            switch -1 {
            case 32768:
                if true {
                    if false {
                        if false {
                            for fi263 := 0; fi263 < 8; fi263++ {
                                v3 = (^(v3))
                                _ = fi263
                            }
                        }
                    } else {
                        v2[""] = "hello"
                    }
                    for fi264 := 0; fi264 < 15; fi264++ {
                        v265 := GenS7{F0: ((("" + "hello") + ("\n\t" + "\n\t")) + ("foo bar" + "hello")), F1: 0, F2: (v24 % ((v24) | 1))}
                        _ = fi264
                        _ = v265
                    }
                } else {
                    fmt.Println(v3, v4, v24)
                }
            case 32767:
                for fi266 := 0; fi266 < 6; fi266++ {
                    if false {
                        switch "a" {
                        case "\n\t":
                            v3 = v3
                        default:
                            v267 := GenS26{F0: fi266, F1: (("a" + "") >= ("\n\t" + ("a" + "hello")))}
                            _ = v267
                        }
                        v268 := GenS8{F0: (fi266 << ((v24) & 63)), F1: 17}
                        var v269 int8 = -69
                        _ = v268
                        _ = v269
                    }
                    _ = fi266
                }
            case 127:
                v3 = (v3 | v3)
            }
        }
        var v270 int64 = 9223372036854775807
        _ = v270
    }
    v271 := GenS2{F0: 1000, F1: -112}
    for fi272 := 0; fi272 < 9; fi272++ {
        v273 := GenS5{F0: (v3 ^ (v3 % ((v3) | 1)))}
        v274 := GenS27{F0: 446}
        _ = fi272
        _ = v273
        _ = v274
    }
    if (!(!("a" != "a"))) {
        v275 := []int{}
        _ = v275
    } else {
        if true {
            v4 += 1
            var v276 bool = (v24 > v24)
            for fi277 := 0; fi277 < 6; fi277++ {
                fmt.Println(v1, fi277, fi277)
                _ = fi277
            }
            _ = v276
        }
        if (0 < -32768) {
            v24 = (^(v24))
            if true {
                for fi278 := 0; fi278 < 9; fi278++ {
                    if (!false) {
                        var v279 int8 = 42
                        _ = v279
                    }
                    var v280 int32 = -2147483648
                    if false {
                        switch (fi278 / ((-176) | 1)) {
                        case 42:
                            if ((v280 ^ (v4 >> (((v24 << (((^(v24))) & 63))) & 63))) <= -770) {
                                v24 = (v24 - (v24 * v24))
                                fmt.Println(fi278, v1)
                            } else {
                                var v281 int = (fi278 & (^(fi278)))
                                var v282 int32 = (v4 * 42)
                                _ = v281
                                _ = v282
                            }
                        default:
                            switch (fi278 ^ v1) {
                            case -413:
                                v24 = 65535
                            case 42:
                                switch (^(fi278)) {
                                case 9223372036854775807:
                                    var v283 float64 = (1.5 / 3.14159)
                                    _ = v283
                                case 1:
                                    switch (v1 | (v1 | v1)) {
                                    case 9223372036854775807:
                                        fmt.Println(v280, v4, v3)
                                    }
                                }
                            default:
                                v1 = -2147483648
                            }
                        }
                        v4 = (v4 | -2147483648)
                        var v284 uint8 = 2
                        _ = v284
                    } else {
                        fmt.Println(v4, v4)
                        v280 -= 1
                        switch "foo bar" {
                        case "hello":
                            v280 = v280
                        default:
                            v285 := GenS1{F0: (^(v1)), F1: genFn1(0, 42, (243 > 64)), F2: v4}
                            _ = v285
                        }
                    }
                    _ = fi278
                    _ = v280
                }
                var v286 uint = v24
                _ = v286
            }
        }
        if false {
            v3 -= v3
            if false {
                switch "" {
                case "hello":
                    var v287 bool = (((501 >= (v1 % (((-(v1))) | 1))) || false) || (((((((^(v1)) < 1) && (-2.25 >= (0.0 / float32(((float64((1e10 / 3.14159)) + ((-2.25 + -2.25) + (float64((0.1 * 1.5)) + 3.14159))) - 0.1))))) && true) || (true || (!false))) || false) || (32767 > 42)))
                    _ = v287
                case "":
                    if (!false) {
                        fmt.Println(v1)
                        if (!false) {
                            var v288 uint16 = 65535
                            v1 = (v1 - (v1 / ((v1) | 1)))
                            _ = v288
                        } else {
                            for fi289 := 0; fi289 < 8; fi289++ {
                                delete(v5, v1)
                                _ = fi289
                            }
                            for fi290 := 0; fi290 < 13; fi290++ {
                                v291 := []int{-569}
                                var v292 int64 = 42
                                switch (^(fi290)) {
                                case 127:
                                    for fi293 := 0; fi293 < 13; fi293++ {
                                        v294 := GenS23{F0: (v292 * v292), F1: ("hello" + "hello")}
                                        var v295 uint = v24
                                        _ = fi293
                                        _ = v294
                                        _ = v295
                                    }
                                case -1:
                                    switch "a" {
                                    case "":
                                        if (!((0 >= 652) || false)) {
                                            if false {
                                                if ((0 != 255) || (!(-1 >= 42))) {
                                                    v296 := GenS1{F0: -990, F1: (v292 % ((genFn1(389, (v1 & fi290), true)) | 1)), F2: v4}
                                                    _ = v296
                                                }
                                            } else {
                                                if (!true) {
                                                    for fi297 := 0; fi297 < 6; fi297++ {
                                                        v24 += v24
                                                        v1 = (fi290 % ((fi290) | 1))
                                                        v5[(-(fi297))] = (v3 & v3)
                                                        _ = fi297
                                                    }
                                                    if (1 <= v292) {
                                                        fmt.Println(fi290, v4)
                                                        v4 = (v4 * 2147483647)
                                                    }
                                                } else {
                                                    var v298 int64 = v292
                                                    v299 := GenS28{F0: true, F1: 42, F2: (v292 / ((v298) | 1))}
                                                    var v300 float32 = 3.14159
                                                    _ = v298
                                                    _ = v299
                                                    _ = v300
                                                }
                                                fmt.Println(fi290)
                                            }
                                        }
                                    default:
                                        for fi301 := 0; fi301 < 9; fi301++ {
                                            fmt.Println(v292, v4)
                                            var v302 int8 = -1
                                            _ = fi301
                                            _ = v302
                                        }
                                    }
                                default:
                                    v5[v1] = (^(v3))
                                }
                                _ = fi290
                                _ = v291
                                _ = v292
                            }
                        }
                        v303 := GenS22{F0: ("\n\t" + "a")}
                        _ = v303
                    }
                default:
                    for fi304 := 0; fi304 < 3; fi304++ {
                        v305 := []int{(fi304 % ((v1) | 1)), 863, v1}
                        v4 = v4
                        _ = fi304
                        _ = v305
                    }
                }
                v306 := []int{128, int(v4), v1, (v1 >> ((636) & 63))}
                if ((891 <= -359) || false) {
                    v1 = (-(v1))
                    var v307 string = ""
                    v308 := GenS18{F0: v3, F1: 65535}
                    _ = v307
                    _ = v308
                }
                _ = v306
            } else {
                v1 -= (v1 % (((^(v1))) | 1))
                var v309 string = ("" + "a")
                v310 := v5[(v1 % ((v1) | 1))]
                _ = v309
                _ = v310
            }
        } else {
            if true {
                var v311 int16 = -908
                _ = v311
            } else {
                v1 -= v1
                if (!false) {
                    if true {
                        var v312 float32 = 0.1
                        _ = v312
                    } else {
                        var v313 int64 = 42
                        if ((v3 % ((165) | 1)) > (^(v3))) {
                            var v314 uint = v24
                            for fi315 := 0; fi315 < 17; fi315++ {
                                for fi316 := 0; fi316 < 20; fi316++ {
                                    var v317 float32 = (3.14159 * 1e10)
                                    var v318 float64 = (((float64(v317) - 0.0) / 1e10) - (-2.25 + float64(v317)))
                                    _ = fi316
                                    _ = v317
                                    _ = v318
                                }
                                v319 := GenS23{F0: (v313 >> ((1) & 63)), F1: "foo bar"}
                                _ = fi315
                                _ = v319
                            }
                            v5[(^(v1))] = v3
                            _ = v314
                        }
                        _ = v313
                    }
                    v320 := []int{(v1 << ((v24) & 63)), -1}
                    _ = v320
                } else {
                    delete(v5, 2147483647)
                }
            }
            var v321 int16 = 0
            v322 := GenS3{F0: 0, F1: "foo bar"}
            _ = v321
            _ = v322
        }
    }
    switch -355 {
    case -804:
        fmt.Println(v3, v1, v1)
    case 32768:
        delete(v5, v1)
    }
    v1 += -1
    if (795 == 559) {
        for fi323 := 0; fi323 < 1; fi323++ {
            v3 = v3
            if (1 == -32768) {
                v1 = (v1 * 0)
            }
            _ = fi323
        }
        var v324 int = (v1 >> (((v24 ^ (^(v24)))) & 63))
        _ = v324
    } else {
        switch v1 {
        case 309:
            v325 := GenS3{F0: 94, F1: "a"}
            _ = v325
        }
        fmt.Println(v3)
        var v326 int = 128
        _ = v326
    }
    switch (v1 << ((v24) & 63)) {
    case 608:
        v327 := []int{}
        _ = v327
    case -150:
        v328 := []int{v1, (^(v1)), (^(v1)), 0}
        _ = v328
    case 128:
        v329 := v5[(v1 * -1)]
        _ = v329
    default:
        for fi330 := 0; fi330 < 13; fi330++ {
            var v331 string = ("foo bar" + "")
            var v332 int8 = -80
            fmt.Println(v3, v332)
            _ = fi330
            _ = v331
            _ = v332
        }
    }
    v24 = 356
    for fi333 := 0; fi333 < 10; fi333++ {
        switch fi333 {
        case -772:
            v3 = 0
        default:
            v1 = v1
        }
        _ = fi333
    }
    for fi334 := 0; fi334 < 4; fi334++ {
        fmt.Println(fi334, fi334)
        fmt.Println(fi334)
        for fi335 := 0; fi335 < 7; fi335++ {
            v1 -= (fi335 + 88)
            v1 = fi334
            v336 := GenS29{F0: 32767, F1: 850, F2: 455}
            _ = fi335
            _ = v336
        }
        _ = fi334
    }
    var v337 int64 = -707
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v337:", v337)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v271
    _ = v3
    _ = v337
    _ = v4
    _ = v5
}
