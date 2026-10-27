package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 string; F1 int }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint32; F1 uint64 }
type GenS8 struct { F0 int32; F1 int64 }
type GenS9 struct { F0 bool; F1 uint32 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 int8; F1 int }
type GenS12 struct { F0 int16; F1 int; F2 uint16 }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 int; F1 int64 }
type GenS15 struct { F0 uint32; F1 uint64; F2 int8 }
type GenS16 struct { F0 uint32 }
type GenS17 struct { F0 string; F1 bool }
type GenS18 struct { F0 uint; F1 int8; F2 int }
type GenS19 struct { F0 int64; F1 bool; F2 uint8 }
type GenS20 struct { F0 int8; F1 uint16; F2 bool }
type GenS21 struct { F0 int }
type GenS22 struct { F0 uint32; F1 uint64 }
type GenS23 struct { F0 int64; F1 int16; F2 uint }
type GenS24 struct { F0 int8; F1 uint }
type GenS25 struct { F0 uint32; F1 int16 }
type GenS26 struct { F0 int; F1 int8 }
type GenS27 struct { F0 int16; F1 uint8 }
type GenS28 struct { F0 bool; F1 int16; F2 uint64 }
type GenS29 struct { F0 string }
type GenS30 struct { F0 int8; F1 uint8; F2 uint32 }
type GenS31 struct { F0 int32 }
type GenS32 struct { F0 bool }
type GenS33 struct { F0 int32; F1 int8 }
func genFn0(p0 uint64, p1 uint64, p2 uint64) uint {
    if false {
        switch -320 {
        case 0:
            for fi7 := 0; fi7 < 7; fi7++ {
                p0 -= p2
                _ = fi7
            }
        case 2147483647:
            fmt.Println(p2, p2, p0)
        case 32768:
            var v8 uint32 = 65535
            _ = v8
        }
        p2 += p1
        var v9 uint8 = 127
        _ = v9
    } else {
        var v10 uint32 = 2147483648
        var v11 uint8 = 128
        v12 := GenS1{F0: 466, F1: 42}
        _ = v10
        _ = v11
        _ = v12
    }
    v13 := []int{-194, -966, 32767, 0}
    fmt.Println(p1, p2, p2)
    switch ("foo bar" + "foo bar") {
    case "":
        if (981 < 42) {
            v14 := GenS2{F0: 255, F1: 65535}
            _ = v14
        } else {
            v15 := GenS1{F0: (p0 << ((18446744073709551615) & 63)), F1: 0}
            _ = v15
        }
    case "hello":
        fmt.Println(p1)
    }
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 2
}
func genFn1(p0 uint32) int64 {
    var v16 uint = 255
    var v17 string = "\n\t"
    fmt.Println(p0, v16, p0)
    _ = p0
    _ = v16
    _ = v17
    return 898
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int8 = 1
    var v3 int8 = 1
    v4 := make(map[uint32]bool)
    var v5 uint64 = 42
    var v6 string = ""
    if (!false) {
        var v18 int64 = genFn1(1)
        var v19 uint = 396
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 17; fi20++ {
        for fi21 := 0; fi21 < 2; fi21++ {
            var v22 int16 = -1
            v23 := GenS1{F0: v5, F1: (-(v2))}
            v24 := GenS1{F0: v5, F1: 1}
            _ = fi21
            _ = v22
            _ = v23
            _ = v24
        }
        switch -179 {
        case -906:
            if (false && false) {
                var v25 uint16 = 255
                v26 := GenS2{F0: 532, F1: 18}
                _ = v25
                _ = v26
            }
        case 409:
            var v27 uint64 = v5
            _ = v27
        case 32767:
            for fi28 := 0; fi28 < 14; fi28++ {
                v2 = -1
                _ = fi28
            }
        }
        _ = fi20
    }
    var v29 uint64 = v5
    v2 += v3
    v30 := GenS1{F0: v29, F1: 119}
    var v31 float32 = ((0.0 / (-2.25 / 0.1)) - 1.5)
    v32 := GenS3{F0: (v6 + "\n\t"), F1: -510}
    v2 = (v3 + (v3 + v2))
    for fi33 := 0; fi33 < 17; fi33++ {
        v34 := []int{(^(fi33)), 4294967295, (fi33 + -799)}
        v6 = (v6 + v6)
        _ = fi33
        _ = v34
    }
    fmt.Println(v2)
    v31 += (float32(0.1) / (v31 * v31))
    for fi35 := 0; fi35 < 9; fi35++ {
        if (458 == 515) {
            var v36 uint = genFn0(4294967295, 0, 18446744073709551615)
            if false {
                var v37 int = 32767
                if ((v2 / ((v3) | 1)) >= 127) {
                    for fi38 := 0; fi38 < 18; fi38++ {
                        v39 := GenS3{F0: "a", F1: v37}
                        _ = fi38
                        _ = v39
                    }
                    for fi40 := 0; fi40 < 19; fi40++ {
                        var v41 uint64 = v29
                        switch (v6 + v6) {
                        case "foo bar":
                            fmt.Println(v31)
                        case "\n\t":
                            var v42 int8 = v3
                            _ = v42
                        }
                        _ = fi40
                        _ = v41
                    }
                    if false {
                        v2 = v3
                    }
                } else {
                    var v43 int16 = 32767
                    _ = v43
                }
                _ = v37
            }
            v44 := GenS4{F0: true}
            _ = v36
            _ = v44
        } else {
            for fi45 := 0; fi45 < 9; fi45++ {
                var v46 uint = 65535
                var v47 bool = true
                switch 664 {
                case 42:
                    switch (v6 + v6) {
                    case "hello":
                        v31 = 0.0
                    }
                    if true {
                        fmt.Println(fi45)
                        v47 = false
                        var v48 float64 = ((3.14159 - float64(1.5)) * 0.1)
                        _ = v48
                    } else {
                        fmt.Println(v31, v46)
                        if v47 {
                            if v47 {
                                v5 += 607
                            } else {
                                fmt.Println(v5)
                                v49 := GenS5{F0: genFn0((v5 % ((2) | 1)), v29, v29)}
                                v50 := GenS1{F0: 18446744073709551615, F1: v3}
                                _ = v49
                                _ = v50
                            }
                            v51 := []int{(fi35 + (fi45 >> ((v46) & 63))), (fi35 % (((fi35 << ((v46) & 63))) | 1)), fi35}
                            _ = v51
                        }
                    }
                    v47 = v47
                case 573:
                    switch 127 {
                    case 127:
                        if (0 != 630) {
                            v52 := GenS2{F0: v46, F1: 2147483648}
                            _ = v52
                        }
                    case -939:
                        switch -102 {
                        case 113:
                            v53 := GenS5{F0: (v46 - 65535)}
                            _ = v53
                        }
                    default:
                        var v54 float64 = (-2.25 / -2.25)
                        _ = v54
                    }
                }
                _ = fi45
                _ = v46
                _ = v47
            }
            if ((v31 < 0.0) && (2147483647 >= 728)) {
                var v55 uint8 = 172
                v56 := GenS3{F0: v6, F1: fi35}
                delete(v4, uint32(v5))
                _ = v55
                _ = v56
            }
        }
        _ = fi35
    }
    if false {
        if (!false) {
            for fi57 := 0; fi57 < 10; fi57++ {
                if ((1.5 == 3.14159) || true) {
                    if (0 < 0) {
                        for fi58 := 0; fi58 < 12; fi58++ {
                            switch (("hello" + "a") + ("foo bar" + v6)) {
                            case "a":
                                v59 := GenS5{F0: 2}
                                _ = v59
                            case "hello":
                                if ((4294967295 == (v29 % (((^(v5))) | 1))) && false) {
                                    v6 = v6
                                    if (true && true) {
                                        var v60 bool = true
                                        v61 := []int{}
                                        fmt.Println(v2, v2)
                                        _ = v60
                                        _ = v61
                                    }
                                    v29 = (v5 & (v29 - v29))
                                }
                            case "foo bar":
                                v62 := GenS6{F0: 2147483647}
                                _ = v62
                            }
                            _ = fi58
                        }
                        var v63 int = fi57
                        v64 := GenS7{F0: 2, F1: (v5 + 0)}
                        _ = v63
                        _ = v64
                    }
                }
                var v65 int16 = 1
                for fi66 := 0; fi66 < 3; fi66++ {
                    if true {
                        v67 := GenS8{F0: -2147483648, F1: 9223372036854775807}
                        var v68 int64 = 0
                        v69 := GenS9{F0: (0.0 < (1e10 * 0.1)), F1: 612}
                        v70 := GenS10{F0: v68}
                        fmt.Println(fi57, v2, v6)
                        _ = v67
                        _ = v68
                        _ = v69
                        _ = v70
                    } else {
                        var v71 bool = (!false)
                        var v72 float32 = v31
                        _ = v71
                        _ = v72
                    }
                    v73 := GenS11{F0: (v2 + v3), F1: 128}
                    _ = fi66
                    _ = v73
                }
                _ = fi57
                _ = v65
            }
            var v74 float64 = 3.14159
            _ = v74
        } else {
            if true {
                v75 := []int{}
                if true {
                    if (false && (v31 > float32((float64(float32(((0.0 * 3.14159) - -2.25))) + 0.1)))) {
                        if true {
                            var v76 float32 = -2.25
                            _ = v76
                        }
                        for fi77 := 0; fi77 < 17; fi77++ {
                            var v78 uint8 = 1
                            for fi79 := 0; fi79 < 7; fi79++ {
                                switch fi77 {
                                case -2147483648:
                                    if (v29 > v5) {
                                        switch v6 {
                                        case "hello":
                                            switch (fi77 ^ fi79) {
                                            case -759:
                                                var v80 uint32 = 812
                                                _ = v80
                                            case 431:
                                                fmt.Println(v78)
                                            }
                                        case "":
                                            if (32767 != 88) {
                                                v6 = ""
                                                if ((true && (true || true)) || true) {
                                                    if ((v29 * v29) < v29) {
                                                        v81 := GenS2{F0: 18446744073709551615, F1: 0}
                                                        var v82 uint16 = 32767
                                                        if len(v75) > 0 {
                                                            v75[0] = 127
                                                        }
                                                        if false {
                                                            v83 := GenS2{F0: 65535, F1: 2}
                                                            v84 := GenS12{F0: 1000, F1: (fi79 + (fi77 * int(v2))), F2: 1}
                                                            switch "hello" {
                                                            case "a":
                                                                switch -9223372036854775808 {
                                                                case 32768:
                                                                    var v85 uint16 = (^(v82))
                                                                    _ = v85
                                                                case -350:
                                                                    var v86 int32 = 42
                                                                    _ = v86
                                                                default:
                                                                    if false {
                                                                        fmt.Println(v82)
                                                                        v87 := GenS13{F0: v82}
                                                                        _ = v87
                                                                    }
                                                                }
                                                                if len(v75) > 0 {
                                                                    v75[0] = fi77
                                                                }
                                                            case "\n\t":
                                                                switch v6 {
                                                                case "a":
                                                                    v6 = "\n\t"
                                                                default:
                                                                    if true {
                                                                        if false {
                                                                            if false {
                                                                                v29 += v5
                                                                                var v88 float64 = 1.5
                                                                                v2 = (v2 - (v2 / ((v3) | 1)))
                                                                                _ = v88
                                                                            } else {
                                                                                switch ("\n\t" + "hello") {
                                                                                case "":
                                                                                    switch (v6 + v6) {
                                                                                    case "\n\t":
                                                                                        if false {
                                                                                            switch (fi79 * fi79) {
                                                                                            case 127:
                                                                                                if (!(-2.25 == -2.25)) {
                                                                                                    var v89 int = 128
                                                                                                    _ = v89
                                                                                                }
                                                                                            default:
                                                                                                if false {
                                                                                                    var v90 float32 = v31
                                                                                                    delete(v1, fi79)
                                                                                                    switch 42 {
                                                                                                    case -2147483648:
                                                                                                        if (656 != -1) {
                                                                                                            if true {
                                                                                                                v29 = (v5 / (((v5 - v5)) | 1))
                                                                                                            } else {
                                                                                                                switch (fi79 / ((32767) | 1)) {
                                                                                                                case -9223372036854775808:
                                                                                                                    if (v82 < 2) {
                                                                                                                        if (1 >= 0) {
                                                                                                                            var v91 uint = 255
                                                                                                                            switch v6 {
                                                                                                                            case "a":
                                                                                                                                switch (fi77 | fi77) {
                                                                                                                                case -698:
                                                                                                                                    delete(v1, (fi77 | (fi79 >> ((v91) & 63))))
                                                                                                                                default:
                                                                                                                                    switch fi79 {
                                                                                                                                    case -9223372036854775808:
                                                                                                                                        v92 := GenS14{F0: (fi79 | (-(fi79))), F1: 0}
                                                                                                                                        _ = v92
                                                                                                                                    default:
                                                                                                                                        var v93 int = fi77
                                                                                                                                        _ = v93
                                                                                                                                    }
                                                                                                                                }
                                                                                                                            case "":
                                                                                                                                v94 := GenS5{F0: (v91 | (v91 << (((v91 + genFn0((v5 | v5), uint64(v82), v5))) & 63)))}
                                                                                                                                _ = v94
                                                                                                                            case "foo bar":
                                                                                                                                v82 += v82
                                                                                                                            default:
                                                                                                                                if (false || (42 != -2147483648)) {
                                                                                                                                    if true {
                                                                                                                                        if true {
                                                                                                                                            fmt.Println(v78)
                                                                                                                                            var v95 int64 = 0
                                                                                                                                            var v96 float64 = (3.14159 * float64(((v90 * v90) + 0.0)))
                                                                                                                                            _ = v95
                                                                                                                                            _ = v96
                                                                                                                                        } else {
                                                                                                                                            switch v6 {
                                                                                                                                            case "\n\t":
                                                                                                                                                var v97 int32 = -1
                                                                                                                                                _ = v97
                                                                                                                                            case "":
                                                                                                                                                fmt.Println(v3, v3, fi79)
                                                                                                                                            }
                                                                                                                                            v91 += (^(v91))
                                                                                                                                        }
                                                                                                                                        v90 = v31
                                                                                                                                        v98 := GenS8{F0: -1, F1: -129}
                                                                                                                                        v99 := GenS3{F0: (("\n\t" + v6) + "a"), F1: (fi79 / (((^(fi79))) | 1))}
                                                                                                                                        _ = v98
                                                                                                                                        _ = v99
                                                                                                                                    } else {
                                                                                                                                        var v100 string = "foo bar"
                                                                                                                                        _ = v100
                                                                                                                                    }
                                                                                                                                    v4[2] = (int16(fi77) <= 141)
                                                                                                                                }
                                                                                                                            }
                                                                                                                            fmt.Println(v90, v3, fi79)
                                                                                                                            v31 = (0.0 * v90)
                                                                                                                            var v101 float64 = (1.5 * 0.0)
                                                                                                                            var v102 int = 0
                                                                                                                            switch 2147483647 {
                                                                                                                            case 32767:
                                                                                                                                v101 = (v101 - 0.0)
                                                                                                                            case 9223372036854775807:
                                                                                                                                var v103 uint64 = v29
                                                                                                                                _ = v103
                                                                                                                            }
                                                                                                                            _ = v101
                                                                                                                            _ = v102
                                                                                                                            _ = v91
                                                                                                                        }
                                                                                                                    }
                                                                                                                case 42:
                                                                                                                    var v104 int16 = 42
                                                                                                                    _ = v104
                                                                                                                case -758:
                                                                                                                    v105 := v1[fi77]
                                                                                                                    _ = v105
                                                                                                                }
                                                                                                            }
                                                                                                            v2 += v2
                                                                                                        }
                                                                                                    case 0:
                                                                                                        v90 += (3.14159 - float32(((float64(float32(float64(((v90 * float32(0.1)) * v90)))) * 0.1) + (1e10 + (1.5 - 1e10)))))
                                                                                                    default:
                                                                                                        if len(v75) >= 2 {
                                                                                                            v75 = v75[1:len(v75)]
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v90
                                                                                                }
                                                                                            }
                                                                                            v82 = uint16(v78)
                                                                                        } else {
                                                                                            fmt.Println(v6)
                                                                                            _ = len(v75)
                                                                                            var v106 uint8 = v78
                                                                                            _ = v106
                                                                                        }
                                                                                    case "":
                                                                                        var v107 uint64 = 607
                                                                                        _ = v107
                                                                                    }
                                                                                default:
                                                                                    fmt.Println(v2, fi79, fi77)
                                                                                }
                                                                                var v108 uint8 = v78
                                                                                if true {
                                                                                    var v109 bool = (390 >= v5)
                                                                                    var v110 bool = true
                                                                                    _ = v109
                                                                                    _ = v110
                                                                                }
                                                                                _ = v108
                                                                            }
                                                                            v6 = ""
                                                                            v31 += 1.5
                                                                        } else {
                                                                            if ((^(v82)) < (v82 - v82)) {
                                                                                if (-49 <= (v3 - 127)) {
                                                                                    if (43 >= 102) {
                                                                                        var v111 float32 = (v31 + (1e10 * v31))
                                                                                        var v112 float64 = 0.0
                                                                                        _ = v111
                                                                                        _ = v112
                                                                                    }
                                                                                }
                                                                            }
                                                                            v2 = v3
                                                                        }
                                                                    } else {
                                                                        v113 := GenS15{F0: 242, F1: v29, F2: v2}
                                                                        switch -2147483648 {
                                                                        case -837:
                                                                            var v114 uint32 = 965
                                                                            _ = v114
                                                                        case 127:
                                                                            if false {
                                                                                v115 := GenS13{F0: (v82 - 0)}
                                                                                v116 := GenS1{F0: v29, F1: 42}
                                                                                v117 := GenS16{F0: 237}
                                                                                _ = v115
                                                                                _ = v116
                                                                                _ = v117
                                                                            } else {
                                                                                v75 = append(v75, fi77)
                                                                                var v118 bool = ((v78 == (v78 & (v78 - 35))) && false)
                                                                                var v119 string = v6
                                                                                _ = v118
                                                                                _ = v119
                                                                            }
                                                                        }
                                                                        if false {
                                                                            fmt.Println(fi79, fi79)
                                                                            var v120 uint32 = 2147483648
                                                                            if (1 < 0) {
                                                                                if len(v75) >= 2 {
                                                                                    v75 = v75[1:len(v75)]
                                                                                }
                                                                                switch "" {
                                                                                case "hello":
                                                                                    _ = len(v75)
                                                                                case "foo bar":
                                                                                    if ((0.0 <= float32(1.5)) || (v78 <= uint8(v29))) {
                                                                                        v3 = (^(v2))
                                                                                        if false {
                                                                                            var v121 uint32 = (^(v120))
                                                                                            v122 := v1[9223372036854775807]
                                                                                            var v123 int8 = (v3 | -128)
                                                                                            _ = v121
                                                                                            _ = v122
                                                                                            _ = v123
                                                                                        } else {
                                                                                            fmt.Println(v29, v82)
                                                                                            _ = len(v1)
                                                                                            v124 := GenS6{F0: uint32(v29)}
                                                                                            _ = v124
                                                                                        }
                                                                                        var v125 bool = (2147483647 != -1)
                                                                                        _ = v125
                                                                                    }
                                                                                    switch (v6 + (v6 + v6)) {
                                                                                    case "a":
                                                                                        if ((!false) && (!(true && false))) {
                                                                                            if (42 != -32768) {
                                                                                                var v126 uint32 = (v120 << ((uint(v82)) & 63))
                                                                                                v127 := GenS12{F0: 32767, F1: (fi79 % (((fi77 ^ fi77)) | 1)), F2: v82}
                                                                                                switch (v6 + (("a" + (v6 + v6)) + "foo bar")) {
                                                                                                case "foo bar":
                                                                                                    v126 -= v120
                                                                                                case "":
                                                                                                    var v128 float64 = 1e10
                                                                                                    _ = v128
                                                                                                default:
                                                                                                    v82 = v82
                                                                                                }
                                                                                                _ = v126
                                                                                                _ = v127
                                                                                            }
                                                                                            fmt.Println(v82)
                                                                                        }
                                                                                    case "hello":
                                                                                        switch "" {
                                                                                        case "":
                                                                                            v31 += v31
                                                                                        default:
                                                                                            switch (-(fi79)) {
                                                                                            case 9223372036854775807:
                                                                                                switch v6 {
                                                                                                case "a":
                                                                                                    var v129 uint32 = (^(v120))
                                                                                                    _ = v129
                                                                                                }
                                                                                                if (!true) {
                                                                                                    _ = len(v75)
                                                                                                    fmt.Println(v120, v31)
                                                                                                    if ((v82 - v82) == (^(v82))) {
                                                                                                        switch -782 {
                                                                                                        case 32768:
                                                                                                            if true {
                                                                                                                var v130 int = (fi79 | fi77)
                                                                                                                v3 -= 122
                                                                                                                _ = v130
                                                                                                            }
                                                                                                        case -535:
                                                                                                            var v131 int8 = (v2 | -123)
                                                                                                            _ = v131
                                                                                                        case 128:
                                                                                                            if (!(true && (!true))) {
                                                                                                                var v132 uint8 = 205
                                                                                                                v133 := GenS4{F0: (!(58 > 545))}
                                                                                                                v2 = (v3 * 49)
                                                                                                                if (float32(-2.25) > v31) {
                                                                                                                    var v134 float64 = float64(v31)
                                                                                                                    fmt.Println(v31, v120)
                                                                                                                    var v135 uint8 = (v78 >> ((2) & 63))
                                                                                                                    _ = v134
                                                                                                                    _ = v135
                                                                                                                }
                                                                                                                _ = v132
                                                                                                                _ = v133
                                                                                                            }
                                                                                                        default:
                                                                                                            switch ("a" + ((((v6 + v6) + "\n\t") + v6) + v6)) {
                                                                                                            case "":
                                                                                                                if false {
                                                                                                                    switch (((v6 + (v6 + "\n\t")) + "") + v6) {
                                                                                                                    case "\n\t":
                                                                                                                        if true {
                                                                                                                            delete(v1, (^(fi77)))
                                                                                                                            if true {
                                                                                                                                switch fi79 {
                                                                                                                                case 315:
                                                                                                                                    if (!true) {
                                                                                                                                        _ = len(v4)
                                                                                                                                        var v136 int16 = 1000
                                                                                                                                        fmt.Println(v6)
                                                                                                                                        _ = v136
                                                                                                                                    } else {
                                                                                                                                        switch fi77 {
                                                                                                                                        case 4294967295:
                                                                                                                                            switch "\n\t" {
                                                                                                                                            case "hello":
                                                                                                                                                v137 := GenS14{F0: (fi79 << ((genFn0((v5 % (((v5 % (((v5 ^ (v5 / ((v5) | 1)))) | 1))) | 1)), v5, 161)) & 63)), F1: int64(v29)}
                                                                                                                                                _ = v137
                                                                                                                                            case "\n\t":
                                                                                                                                                v78 -= (v78 & v78)
                                                                                                                                            case "foo bar":
                                                                                                                                                var v138 uint16 = uint16(v78)
                                                                                                                                                _ = v138
                                                                                                                                            default:
                                                                                                                                                var v139 uint64 = v29
                                                                                                                                                _ = v139
                                                                                                                                            }
                                                                                                                                        case -639:
                                                                                                                                            if false {
                                                                                                                                                if len(v75) > 0 {
                                                                                                                                                    v75[0] = (fi79 + fi79)
                                                                                                                                                }
                                                                                                                                                fmt.Println(v3)
                                                                                                                                                v140 := v1[fi79]
                                                                                                                                                _ = v140
                                                                                                                                            }
                                                                                                                                        case 32768:
                                                                                                                                            fmt.Println(fi79)
                                                                                                                                            if true {
                                                                                                                                                fmt.Println(v2, v29, v29)
                                                                                                                                                v75 = append(v75, fi79)
                                                                                                                                                var v141 uint64 = (^(v5))
                                                                                                                                                _ = v141
                                                                                                                                            } else {
                                                                                                                                                var v142 bool = true
                                                                                                                                                v143 := v1[fi77]
                                                                                                                                                switch fi79 {
                                                                                                                                                case 128:
                                                                                                                                                    v2 -= int8(fi77)
                                                                                                                                                case 1:
                                                                                                                                                    switch "" {
                                                                                                                                                    case "\n\t":
                                                                                                                                                        v1[fi77] = (v120 % ((v143) | 1))
                                                                                                                                                    case "a":
                                                                                                                                                        var v144 int16 = 189
                                                                                                                                                        _ = v144
                                                                                                                                                    }
                                                                                                                                                case 4294967295:
                                                                                                                                                    v145 := v1[fi79]
                                                                                                                                                    _ = v145
                                                                                                                                                default:
                                                                                                                                                    var v146 int32 = 2147483647
                                                                                                                                                    _ = v146
                                                                                                                                                }
                                                                                                                                                var v147 string = "hello"
                                                                                                                                                _ = v142
                                                                                                                                                _ = v143
                                                                                                                                                _ = v147
                                                                                                                                            }
                                                                                                                                        }
                                                                                                                                        var v148 uint16 = (v82 % (((v82 | v82)) | 1))
                                                                                                                                        _ = v148
                                                                                                                                    }
                                                                                                                                case 32767:
                                                                                                                                    v149 := GenS14{F0: 291, F1: genFn1(v120)}
                                                                                                                                    _ = v149
                                                                                                                                case 4294967295:
                                                                                                                                    if ((v82 + v82) == (v82 % ((v82) | 1))) {
                                                                                                                                        switch (fi77 * (fi77 >> ((948) & 63))) {
                                                                                                                                        case -9223372036854775808:
                                                                                                                                            fmt.Println(v3, v31, v2)
                                                                                                                                        case 127:
                                                                                                                                            var v150 uint32 = v120
                                                                                                                                            _ = v150
                                                                                                                                        case 248:
                                                                                                                                            v78 += 255
                                                                                                                                        }
                                                                                                                                        if len(v75) >= 2 {
                                                                                                                                            v75 = v75[1:len(v75)]
                                                                                                                                        }
                                                                                                                                    } else {
                                                                                                                                        if true {
                                                                                                                                            v3 += v3
                                                                                                                                            var v151 string = v6
                                                                                                                                            _ = v151
                                                                                                                                        } else {
                                                                                                                                            var v152 uint = genFn0(uint64(v78), (v5 - v5), 817)
                                                                                                                                            if true {
                                                                                                                                                v82 -= uint16(v120)
                                                                                                                                                if true {
                                                                                                                                                    fmt.Println(v2)
                                                                                                                                                    v3 = v2
                                                                                                                                                } else {
                                                                                                                                                    _ = len(v1)
                                                                                                                                                }
                                                                                                                                                v153 := GenS17{F0: (v6 + (v6 + ((v6 + ("hello" + v6)) + ((v6 + "foo bar") + (v6 + ("" + v6)))))), F1: (-941 == fi77)}
                                                                                                                                                _ = v153
                                                                                                                                            }
                                                                                                                                            _ = v152
                                                                                                                                        }
                                                                                                                                        var v154 float64 = 1.5
                                                                                                                                        v31 += ((-2.25 + (v31 / (float32((v154 + 3.14159)) / v31))) * ((0.0 + v31) + float32((float64(((0.1 * 3.14159) - float32(((0.1 * v154) * 0.0)))) + v154))))
                                                                                                                                        _ = v154
                                                                                                                                    }
                                                                                                                                    _ = len(v4)
                                                                                                                                    var v155 string = (v6 + "\n\t")
                                                                                                                                    _ = v155
                                                                                                                                }
                                                                                                                            } else {
                                                                                                                                switch (fi79 << ((18446744073709551615) & 63)) {
                                                                                                                                case 886:
                                                                                                                                    _ = len(v75)
                                                                                                                                }
                                                                                                                                v2 += v3
                                                                                                                            }
                                                                                                                        } else {
                                                                                                                            if (0 >= 1000) {
                                                                                                                                var v156 uint32 = 0
                                                                                                                                v82 = 2
                                                                                                                                v157 := GenS18{F0: 65535, F1: (v3 - -1), F2: (fi79 & (^(fi79)))}
                                                                                                                                _ = v156
                                                                                                                                _ = v157
                                                                                                                            } else {
                                                                                                                                switch fi77 {
                                                                                                                                case 42:
                                                                                                                                    var v158 uint = 0
                                                                                                                                    _ = v158
                                                                                                                                case -341:
                                                                                                                                    if len(v75) >= 2 {
                                                                                                                                        v75 = v75[1:len(v75)]
                                                                                                                                    }
                                                                                                                                case 9223372036854775807:
                                                                                                                                    fmt.Println(v5, v6)
                                                                                                                                    if false {
                                                                                                                                        var v159 int64 = 9223372036854775807
                                                                                                                                        v160 := v4[(v120 & v120)]
                                                                                                                                        _ = v159
                                                                                                                                        _ = v160
                                                                                                                                    } else {
                                                                                                                                        switch fi77 {
                                                                                                                                        case 2147483647:
                                                                                                                                            _ = len(v4)
                                                                                                                                        case 128:
                                                                                                                                            v161 := v4[v120]
                                                                                                                                            _ = v161
                                                                                                                                        case 32768:
                                                                                                                                            var v162 int = fi77
                                                                                                                                            _ = v162
                                                                                                                                        default:
                                                                                                                                            fmt.Println(v78, v120)
                                                                                                                                        }
                                                                                                                                        switch 42 {
                                                                                                                                        case 370:
                                                                                                                                            v163 := GenS19{F0: 815, F1: true, F2: (v78 | (v78 - v78))}
                                                                                                                                            _ = v163
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                    if true {
                                                                                                                                        var v164 uint32 = (v120 * v120)
                                                                                                                                        _ = v164
                                                                                                                                    } else {
                                                                                                                                        fmt.Println(v6)
                                                                                                                                        v6 = "\n\t"
                                                                                                                                        var v165 int16 = 402
                                                                                                                                        _ = v165
                                                                                                                                    }
                                                                                                                                default:
                                                                                                                                    if len(v75) > 0 {
                                                                                                                                        v75[0] = (fi77 >> ((4294967295) & 63))
                                                                                                                                    }
                                                                                                                                }
                                                                                                                                switch ("\n\t" + ("foo bar" + ("foo bar" + ("" + v6)))) {
                                                                                                                                case "":
                                                                                                                                    var v166 uint8 = uint8(v5)
                                                                                                                                    if (false && true) {
                                                                                                                                        switch v6 {
                                                                                                                                        case "a":
                                                                                                                                            v78 -= 1
                                                                                                                                        case "foo bar":
                                                                                                                                            v4[uint32(v29)] = false
                                                                                                                                        }
                                                                                                                                    } else {
                                                                                                                                        fmt.Println(v6, v3)
                                                                                                                                        if false {
                                                                                                                                            if len(v75) >= 2 {
                                                                                                                                                v75 = v75[1:len(v75)]
                                                                                                                                            }
                                                                                                                                            v82 = v82
                                                                                                                                            v167 := GenS14{F0: fi79, F1: 9223372036854775807}
                                                                                                                                            _ = v167
                                                                                                                                        }
                                                                                                                                        var v168 int32 = -1
                                                                                                                                        _ = v168
                                                                                                                                    }
                                                                                                                                    _ = v166
                                                                                                                                case "\n\t":
                                                                                                                                    var v169 int64 = 42
                                                                                                                                    _ = v169
                                                                                                                                }
                                                                                                                            }
                                                                                                                        }
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    var v170 uint32 = 739
                                                                                                                    if len(v75) >= 2 {
                                                                                                                        v75 = v75[1:len(v75)]
                                                                                                                    }
                                                                                                                    _ = v170
                                                                                                                }
                                                                                                            case "foo bar":
                                                                                                                if (!true) {
                                                                                                                    var v171 int = (fi77 / ((-9223372036854775808) | 1))
                                                                                                                    _ = v171
                                                                                                                } else {
                                                                                                                    v3 -= v3
                                                                                                                    switch v6 {
                                                                                                                    case "hello":
                                                                                                                        var v172 uint32 = 635
                                                                                                                        _ = v172
                                                                                                                    }
                                                                                                                }
                                                                                                            case "\n\t":
                                                                                                                var v173 string = (v6 + "")
                                                                                                                _ = v173
                                                                                                            default:
                                                                                                                v29 += v29
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                v29 = 473
                                                                                            case 32767:
                                                                                                if ((("" + "foo bar") + ("\n\t" + ("hello" + "\n\t"))) == v6) {
                                                                                                    if true {
                                                                                                        var v174 string = "a"
                                                                                                        _ = v174
                                                                                                    }
                                                                                                    if (-36 == -1) {
                                                                                                        v175 := GenS20{F0: v2, F1: (v82 ^ v82), F2: (true && (!((v78 ^ v78) < 2)))}
                                                                                                        if false {
                                                                                                            var v176 float32 = (v31 - ((((v31 / (-2.25 - (1e10 + v31))) * (v31 * 1e10)) - v31) + ((v31 / v31) - -2.25)))
                                                                                                            _ = v176
                                                                                                        } else {
                                                                                                            if (!((genFn1((v120 % ((395) | 1))) < 9223372036854775807) && (v6 != "hello"))) {
                                                                                                                var v177 int64 = -888
                                                                                                                v120 = (^(v120))
                                                                                                                if len(v75) >= 2 {
                                                                                                                    v75 = v75[1:len(v75)]
                                                                                                                }
                                                                                                                _ = v177
                                                                                                            } else {
                                                                                                                v78 = (v78 << ((887) & 63))
                                                                                                                v5 = (v5 ^ (^(v5)))
                                                                                                                switch (v6 + "\n\t") {
                                                                                                                case "":
                                                                                                                    if (fi79 > (fi79 % (((fi77 >> ((uint(v120)) & 63))) | 1))) {
                                                                                                                        switch v6 {
                                                                                                                        case "hello":
                                                                                                                            if true {
                                                                                                                                v120 = (v120 + v120)
                                                                                                                                switch "foo bar" {
                                                                                                                                case "hello":
                                                                                                                                    switch (^(fi77)) {
                                                                                                                                    case 9223372036854775807:
                                                                                                                                        var v178 int8 = 0
                                                                                                                                        _ = v178
                                                                                                                                    case 4294967295:
                                                                                                                                        var v179 string = ("a" + (v6 + "\n\t"))
                                                                                                                                        _ = v179
                                                                                                                                    case -1:
                                                                                                                                        if (true && (!true)) {
                                                                                                                                            v82 += v82
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                case "foo bar":
                                                                                                                                    if (42 >= -903) {
                                                                                                                                        v180 := GenS12{F0: -669, F1: (fi77 << ((354) & 63)), F2: (^(v82))}
                                                                                                                                        v29 = (v5 ^ (v29 | 0))
                                                                                                                                        if false {
                                                                                                                                            v181 := GenS13{F0: v82}
                                                                                                                                            if (!true) {
                                                                                                                                                v31 += (((v31 + v31) * ((v31 * float32(-2.25)) / (0.0 + 1e10))) * (v31 * (1.5 / (((v31 / 3.14159) - (float32(-2.25) / v31)) * v31))))
                                                                                                                                                var v182 bool = (!(0 == 9223372036854775807))
                                                                                                                                                _ = v182
                                                                                                                                            } else {
                                                                                                                                                var v183 int32 = 1
                                                                                                                                                _ = v183
                                                                                                                                            }
                                                                                                                                            v78 = uint8(v82)
                                                                                                                                            _ = v181
                                                                                                                                        } else {
                                                                                                                                            switch (v6 + (((("" + v6) + "\n\t") + "a") + "foo bar")) {
                                                                                                                                            case "":
                                                                                                                                                _ = len(v75)
                                                                                                                                            default:
                                                                                                                                                if ((v3 & (v2 | (v3 * v3))) > (v2 ^ v3)) {
                                                                                                                                                    if (uint8(v5) != (v78 ^ v78)) {
                                                                                                                                                        v75 = append(v75, fi79)
                                                                                                                                                    }
                                                                                                                                                    v120 -= 844
                                                                                                                                                }
                                                                                                                                            }
                                                                                                                                        }
                                                                                                                                        _ = v180
                                                                                                                                    } else {
                                                                                                                                        var v184 bool = true
                                                                                                                                        v31 = (v31 + v31)
                                                                                                                                        v3 -= v3
                                                                                                                                        _ = v184
                                                                                                                                    }
                                                                                                                                case "a":
                                                                                                                                    if true {
                                                                                                                                        v185 := GenS21{F0: (fi79 + (fi79 | (fi77 | (fi77 / (((fi79 * (fi77 / (((fi79 - (fi79 ^ fi79))) | 1)))) | 1)))))}
                                                                                                                                        fmt.Println(v120)
                                                                                                                                        _ = v185
                                                                                                                                    } else {
                                                                                                                                        var v186 int = fi79
                                                                                                                                        var v187 int64 = int64(v29)
                                                                                                                                        _ = v186
                                                                                                                                        _ = v187
                                                                                                                                    }
                                                                                                                                default:
                                                                                                                                    v188 := GenS14{F0: (fi79 + fi79), F1: -829}
                                                                                                                                    _ = v188
                                                                                                                                }
                                                                                                                                if false {
                                                                                                                                    switch v6 {
                                                                                                                                    case "":
                                                                                                                                        var v189 int64 = -990
                                                                                                                                        _ = v189
                                                                                                                                    case "a":
                                                                                                                                        if (!(!true)) {
                                                                                                                                            delete(v4, (v120 % (((v120 * v120)) | 1)))
                                                                                                                                            v6 = "foo bar"
                                                                                                                                        } else {
                                                                                                                                            v29 -= 4294967295
                                                                                                                                        }
                                                                                                                                        v190 := GenS22{F0: uint32(v82), F1: (^(v29))}
                                                                                                                                        _ = v190
                                                                                                                                    case "foo bar":
                                                                                                                                        switch fi79 {
                                                                                                                                        case -465:
                                                                                                                                            var v191 float64 = 0.0
                                                                                                                                            _ = v191
                                                                                                                                        case 1:
                                                                                                                                            var v192 uint16 = v82
                                                                                                                                            _ = v192
                                                                                                                                        case -2147483648:
                                                                                                                                            fmt.Println(v29)
                                                                                                                                        default:
                                                                                                                                            _ = len(v1)
                                                                                                                                        }
                                                                                                                                    default:
                                                                                                                                        var v193 uint16 = 2
                                                                                                                                        _ = v193
                                                                                                                                    }
                                                                                                                                }
                                                                                                                            } else {
                                                                                                                                v78 -= 1
                                                                                                                            }
                                                                                                                        case "a":
                                                                                                                            var v194 int8 = (v2 - 0)
                                                                                                                            _ = v194
                                                                                                                        case "\n\t":
                                                                                                                            if ((v120 >= v120) && (v5 < 18446744073709551615)) {
                                                                                                                                fmt.Println(v78, fi77, v78)
                                                                                                                                var v195 float32 = 0.1
                                                                                                                                _ = v195
                                                                                                                            } else {
                                                                                                                                if (fi77 == fi79) {
                                                                                                                                    v196 := GenS20{F0: (v2 * 98), F1: uint16(v29), F2: (!true)}
                                                                                                                                    _ = v196
                                                                                                                                }
                                                                                                                                var v197 uint16 = (v82 / (((v82 % ((255) | 1))) | 1))
                                                                                                                                fmt.Println(v29)
                                                                                                                                _ = v197
                                                                                                                            }
                                                                                                                        }
                                                                                                                        fmt.Println(v120, v2)
                                                                                                                        var v198 int8 = int8(fi79)
                                                                                                                        var v199 int64 = 757
                                                                                                                        switch "hello" {
                                                                                                                        case "a":
                                                                                                                            var v200 uint8 = 42
                                                                                                                            _ = v200
                                                                                                                        case "":
                                                                                                                            v201 := GenS13{F0: v82}
                                                                                                                            _ = v201
                                                                                                                        }
                                                                                                                        _ = len(v1)
                                                                                                                        var v202 int32 = -880
                                                                                                                        v31 = float32(3.14159)
                                                                                                                        fmt.Println(v29)
                                                                                                                        v31 = v31
                                                                                                                        _ = v198
                                                                                                                        _ = v199
                                                                                                                        _ = v202
                                                                                                                    }
                                                                                                                case "\n\t":
                                                                                                                    var v203 uint = 527
                                                                                                                    _ = v203
                                                                                                                default:
                                                                                                                    var v204 uint16 = (^(v82))
                                                                                                                    _ = v204
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                        v120 += 593
                                                                                                        _ = v175
                                                                                                    } else {
                                                                                                        v205 := GenS23{F0: -1, F1: 956, F2: genFn0(uint64(v120), (v5 % ((v29) | 1)), v29)}
                                                                                                        var v206 uint = genFn0(uint64(v82), 871, 755)
                                                                                                        _ = v205
                                                                                                        _ = v206
                                                                                                    }
                                                                                                    var v207 int32 = 0
                                                                                                    _ = v207
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    case "foo bar":
                                                                                        v78 = 42
                                                                                    default:
                                                                                        switch v6 {
                                                                                        case "a":
                                                                                            if len(v75) >= 2 {
                                                                                                v75 = v75[1:len(v75)]
                                                                                            }
                                                                                            var v208 float64 = 1.5
                                                                                            v31 = (v31 * v31)
                                                                                            var v209 string = "a"
                                                                                            _ = v208
                                                                                            _ = v209
                                                                                        default:
                                                                                            v2 = 1
                                                                                        }
                                                                                        var v210 int8 = (v3 >> ((1) & 63))
                                                                                        v211 := GenS1{F0: 4294967295, F1: 9}
                                                                                        _ = v210
                                                                                        _ = v211
                                                                                    }
                                                                                case "a":
                                                                                    v120 = v120
                                                                                }
                                                                                var v212 string = ((v6 + v6) + (v6 + v6))
                                                                                _ = v212
                                                                            } else {
                                                                                v3 = (-(v3))
                                                                            }
                                                                            var v213 uint = 718
                                                                            if (v120 <= 4294967295) {
                                                                                var v214 int64 = -1
                                                                                var v215 int64 = (v214 + v214)
                                                                                var v216 uint = genFn0(v5, v29, (v29 / (((^(v29))) | 1)))
                                                                                _ = v214
                                                                                _ = v215
                                                                                _ = v216
                                                                            }
                                                                            v213 += v213
                                                                            _ = v120
                                                                            _ = v213
                                                                        } else {
                                                                            v217 := GenS24{F0: v3, F1: 4294967295}
                                                                            _ = v217
                                                                        }
                                                                        v5 = (v29 ^ (v5 * (v5 - v29)))
                                                                        _ = v113
                                                                    }
                                                                }
                                                            }
                                                            _ = v83
                                                            _ = v84
                                                        }
                                                        var v218 uint64 = v29
                                                        _ = v218
                                                        _ = v81
                                                        _ = v82
                                                    }
                                                    v29 = (v5 + v5)
                                                    var v219 int32 = 42
                                                    _ = v219
                                                }
                                            }
                                        case "foo bar":
                                            if (float32((0.0 * (1e10 - (3.14159 * (0.0 + (0.1 - ((0.0 + 3.14159) * -2.25))))))) <= v31) {
                                                var v220 int16 = -672
                                                v221 := GenS20{F0: -128, F1: 0, F2: (fi77 != fi77)}
                                                _ = v220
                                                _ = v221
                                            } else {
                                                v222 := v4[uint32(v78)]
                                                if len(v75) > 0 {
                                                    v75[0] = (fi77 / ((fi79) | 1))
                                                }
                                                v4[263] = v222
                                                v29 = 2
                                                v4[920] = (float32(3.14159) >= v31)
                                                v3 += (v3 - (v3 >> ((255) & 63)))
                                                _ = v222
                                            }
                                            var v223 uint8 = v78
                                            _ = v223
                                        }
                                        var v224 int16 = 1000
                                        var v225 int8 = (-(v3))
                                        _ = v224
                                        _ = v225
                                    }
                                case 32768:
                                    v226 := v1[(fi79 + -2147483648)]
                                    _ = v226
                                case -682:
                                    var v227 uint32 = 369
                                    _ = v227
                                }
                                if false {
                                    var v228 uint32 = 1
                                    switch "foo bar" {
                                    case "\n\t":
                                        var v229 string = "\n\t"
                                        _ = v229
                                    case "hello":
                                        v5 = v5
                                    default:
                                        v31 -= (v31 + float32(1e10))
                                    }
                                    _ = v228
                                }
                                if (((((float32((0.1 / 3.14159)) - v31) - (1.5 / (v31 * float32(3.14159)))) + (float32(float64(float32(-2.25))) * (-2.25 - 0.0))) / v31) <= v31) {
                                    if true {
                                        _ = len(v75)
                                        v6 = v6
                                    }
                                }
                                _ = fi79
                            }
                            _ = len(v1)
                            _ = fi77
                            _ = v78
                        }
                        v31 -= v31
                    } else {
                        v3 = v3
                    }
                    v3 = 34
                } else {
                    if false {
                        fmt.Println(v2, v29)
                    } else {
                        v31 = 1e10
                        for fi230 := 0; fi230 < 13; fi230++ {
                            v3 = 42
                            for fi231 := 0; fi231 < 3; fi231++ {
                                v31 -= 3.14159
                                v3 -= v3
                                _ = fi231
                            }
                            _ = fi230
                        }
                        fmt.Println(v2, v29)
                    }
                }
                _ = v75
            } else {
                if (1 > 32767) {
                    var v232 uint64 = v29
                    _ = v232
                }
                var v233 float64 = ((float64(3.14159) / 1.0) + float64(0.0))
                for fi234 := 0; fi234 < 19; fi234++ {
                    v233 = v233
                    var v235 string = v6
                    if false {
                        for fi236 := 0; fi236 < 18; fi236++ {
                            switch v6 {
                            case "\n\t":
                                switch v6 {
                                case "\n\t":
                                    v2 = (v2 - -128)
                                case "hello":
                                    if false {
                                        v237 := []int{fi234, (^(fi234))}
                                        _ = v237
                                    } else {
                                        switch (fi234 * -264) {
                                        case 4294967295:
                                            fmt.Println(fi234)
                                        case -9223372036854775808:
                                            var v238 uint64 = (v29 / (((^(v29))) | 1))
                                            _ = v238
                                        case -830:
                                            if true {
                                                v29 = 4294967295
                                                v233 = 1e10
                                            }
                                        }
                                        v239 := GenS25{F0: 42, F1: -701}
                                        _ = v239
                                    }
                                default:
                                    v3 = v3
                                }
                            case "foo bar":
                                if true {
                                    fmt.Println(v5, fi234)
                                    var v240 int64 = -9223372036854775808
                                    switch (fi236 | (fi234 % (((fi234 << ((738) & 63))) | 1))) {
                                    case -2147483648:
                                        switch (fi236 & 2147483647) {
                                        case -2147483648:
                                            delete(v1, fi234)
                                        default:
                                            fmt.Println(v235, v233)
                                        }
                                    case -692:
                                        v241 := GenS26{F0: int(v29), F1: 127}
                                        _ = v241
                                    case 1:
                                        v240 -= v240
                                    }
                                    var v242 uint64 = 18446744073709551615
                                    v243 := []int{}
                                    v31 += v31
                                    if len(v243) >= 2 {
                                        v243 = v243[1:len(v243)]
                                    }
                                    _ = v240
                                    _ = v242
                                    _ = v243
                                }
                            case "hello":
                                fmt.Println(fi234)
                            default:
                                var v244 uint = 0
                                _ = v244
                            }
                            v245 := GenS7{F0: 174, F1: (v29 + 436)}
                            if (uint(v5) < 42) {
                                switch ((v6 + "foo bar") + v6) {
                                case "":
                                    v29 = (v5 | v29)
                                    var v246 float32 = 3.14159
                                    _ = v246
                                case "a":
                                    var v247 float32 = (v31 * (1e10 + (v31 * ((v31 / float32(v233)) - 1.5))))
                                    _ = v247
                                case "\n\t":
                                    v248 := []int{fi236, (fi236 >> ((907) & 63)), 2147483647}
                                    _ = v248
                                default:
                                    if ((v5 << ((uint(v29)) & 63)) != 4294967295) {
                                        switch "\n\t" {
                                        case "hello":
                                            if false {
                                                if (!true) {
                                                    var v249 uint = 338
                                                    var v250 int32 = 243
                                                    var v251 int32 = int32(v249)
                                                    _ = v249
                                                    _ = v250
                                                    _ = v251
                                                } else {
                                                    v252 := []int{(fi234 & fi236), (fi234 / ((-1) | 1)), (fi236 | fi234), int(v5)}
                                                    _ = len(v1)
                                                    _ = v252
                                                }
                                                v253 := []int{128, (fi236 % ((42) | 1))}
                                                v254 := GenS27{F0: 0, F1: uint8(v29)}
                                                _ = v253
                                                _ = v254
                                            }
                                        case "":
                                            v255 := GenS24{F0: v3, F1: 163}
                                            _ = v255
                                        case "\n\t":
                                            switch v6 {
                                            case "":
                                                switch (("hello" + "a") + "") {
                                                case "":
                                                    v4[uint32(v5)] = false
                                                }
                                            case "a":
                                                switch ("foo bar" + v235) {
                                                case "":
                                                    var v256 uint8 = 0
                                                    switch v235 {
                                                    case "":
                                                        v6 = (v235 + "hello")
                                                    case "foo bar":
                                                        var v257 int8 = v2
                                                        _ = v257
                                                    default:
                                                        v258 := []int{(fi234 / ((4294967295) | 1)), (fi234 % (((fi236 - (fi234 - fi236))) | 1))}
                                                        _ = v258
                                                    }
                                                    v259 := GenS1{F0: (^(v5)), F1: v2}
                                                    _ = v256
                                                    _ = v259
                                                default:
                                                    v260 := GenS14{F0: (fi236 >> ((928) & 63)), F1: 1}
                                                    _ = v260
                                                }
                                            }
                                        }
                                    } else {
                                        v261 := GenS28{F0: false, F1: 0, F2: v29}
                                        v6 = ("\n\t" + "\n\t")
                                        _ = v261
                                    }
                                }
                                v262 := []int{fi234}
                                v3 -= v3
                                if (true && false) {
                                    switch v6 {
                                    case "":
                                        fmt.Println(v31, fi234)
                                        v5 = 0
                                    case "\n\t":
                                        fmt.Println(v235, v5, v235)
                                    case "a":
                                        var v263 int16 = 1
                                        _ = v263
                                    default:
                                        fmt.Println(v31)
                                        delete(v4, 42)
                                    }
                                }
                                _ = v262
                            } else {
                                v264 := GenS29{F0: (v235 + v6)}
                                v29 += (v29 % ((361) | 1))
                                var v265 float64 = ((v233 * v233) - (v233 + v233))
                                _ = v264
                                _ = v265
                            }
                            _ = fi236
                            _ = v245
                        }
                        v235 = v6
                    } else {
                        var v266 int64 = 0
                        _ = v266
                    }
                    _ = fi234
                    _ = v235
                }
                _ = v233
            }
            v267 := GenS20{F0: 70, F1: 0, F2: true}
            for fi268 := 0; fi268 < 9; fi268++ {
                v269 := GenS3{F0: v6, F1: 2147483647}
                v270 := GenS2{F0: 2, F1: 52}
                var v271 int64 = 42
                _ = fi268
                _ = v269
                _ = v270
                _ = v271
            }
            _ = v267
        }
        switch (v6 + v6) {
        case "":
            switch 128 {
            case 128:
                v3 -= (v2 - -92)
            }
        case "a":
            for fi272 := 0; fi272 < 15; fi272++ {
                var v273 float64 = 3.14159
                switch v6 {
                case "":
                    v29 = (v29 & v29)
                case "a":
                    var v274 int = fi272
                    _ = v274
                default:
                    switch -2147483648 {
                    case 879:
                        v273 -= v273
                    }
                }
                switch fi272 {
                case -339:
                    v275 := []int{(fi272 << ((4294967295) & 63))}
                    _ = v275
                case -855:
                    if false {
                        v2 -= (v3 ^ v2)
                        v276 := GenS20{F0: v3, F1: 65535, F2: (-2147483648 < 42)}
                        _ = v276
                    } else {
                        delete(v4, 2147483647)
                        v5 += (v5 + v5)
                    }
                }
                _ = fi272
                _ = v273
            }
        default:
            switch -489 {
            case -1:
                if true {
                    v277 := []int{-9223372036854775808, 4294967295, 2147483647}
                    for fi278 := 0; fi278 < 18; fi278++ {
                        v279 := GenS30{F0: v3, F1: 11, F2: 42}
                        switch "hello" {
                        case "":
                            var v280 uint32 = 65535
                            _ = v280
                        case "hello":
                            v3 = v2
                        case "a":
                            v281 := GenS31{F0: 2147483647}
                            _ = v281
                        default:
                            var v282 int8 = 1
                            _ = v282
                        }
                        _ = fi278
                        _ = v279
                    }
                    v6 = ("" + ((v6 + (v6 + (v6 + v6))) + ("foo bar" + "a")))
                    _ = v277
                }
            case 128:
                if true {
                    v6 = v6
                }
            }
        }
        v2 = v3
    }
    for fi283 := 0; fi283 < 12; fi283++ {
        var v284 string = v6
        _ = fi283
        _ = v284
    }
    if (genFn1(2147483647) == 46) {
        var v285 int32 = 99
        _ = v285
    } else {
        if (2 < 739) {
            fmt.Println(v29, v2, v5)
            v29 = (v5 - (v5 - (^(v5))))
            for fi286 := 0; fi286 < 19; fi286++ {
                v287 := GenS32{F0: false}
                fmt.Println(v31)
                if true {
                    var v288 uint64 = (v29 / (((v5 >> ((18446744073709551615) & 63))) | 1))
                    v289 := GenS20{F0: v2, F1: 2, F2: false}
                    v2 += (v2 % (((v2 >> ((85) & 63))) | 1))
                    _ = v288
                    _ = v289
                } else {
                    v290 := []int{(fi286 << ((825) & 63)), (fi286 % (((fi286 % (((fi286 & (-(fi286)))) | 1))) | 1)), fi286, (fi286 % (((fi286 % ((fi286) | 1))) | 1))}
                    v291 := GenS2{F0: 139, F1: 2147483647}
                    _ = v290
                    _ = v291
                }
                _ = fi286
                _ = v287
            }
        } else {
            v5 -= (^(v5))
            for fi292 := 0; fi292 < 20; fi292++ {
                v29 += 164
                v293 := GenS33{F0: 42, F1: -115}
                var v294 int8 = (v3 ^ (v3 << ((42) & 63)))
                _ = fi292
                _ = v293
                _ = v294
            }
            var v295 uint = 18446744073709551615
            _ = v295
        }
        for fi296 := 0; fi296 < 11; fi296++ {
            var v297 uint = 4294967295
            for fi298 := 0; fi298 < 13; fi298++ {
                var v299 uint16 = 788
                v300 := []int{-1}
                v3 += v3
                _ = fi298
                _ = v299
                _ = v300
            }
            _ = fi296
            _ = v297
        }
        var v301 string = (v6 + "")
        _ = v301
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
