package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint16; F1 int8 }
type GenS5 struct { F0 int; F1 int16 }
type GenS6 struct { F0 int8; F1 uint }
type GenS7 struct { F0 int16; F1 uint8; F2 string }
type GenS8 struct { F0 int32; F1 uint8 }
type GenS9 struct { F0 int8; F1 int; F2 int64 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 string; F1 bool; F2 uint8 }
type GenS12 struct { F0 string; F1 int; F2 uint8 }
type GenS13 struct { F0 uint8; F1 uint32 }
type GenS14 struct { F0 uint64 }
type GenS15 struct { F0 uint16 }
type GenS16 struct { F0 uint64 }
type GenS17 struct { F0 bool }
type GenS18 struct { F0 int; F1 bool }
type GenS19 struct { F0 uint16 }
type GenS20 struct { F0 int16 }
type GenS21 struct { F0 int32; F1 uint16; F2 int16 }
type GenS22 struct { F0 uint8; F1 uint64 }
type GenS23 struct { F0 uint32; F1 int16; F2 int64 }
type GenS24 struct { F0 int16; F1 int }
type GenS25 struct { F0 uint8 }
type GenS26 struct { F0 int16 }
type GenS27 struct { F0 uint16; F1 int }
type GenS28 struct { F0 int64; F1 int16; F2 uint64 }
type GenS29 struct { F0 int32; F1 int16 }
type GenS30 struct { F0 uint64 }
type GenS31 struct { F0 uint32 }
type GenS32 struct { F0 int64 }
type GenS33 struct { F0 int64 }
type GenS34 struct { F0 uint64; F1 uint; F2 uint16 }
type GenS35 struct { F0 uint64; F1 uint32; F2 uint32 }
type GenS36 struct { F0 int64; F1 uint }
type GenS37 struct { F0 int64 }
type GenS38 struct { F0 bool; F1 uint32 }
type GenS39 struct { F0 int64 }
type GenS40 struct { F0 uint16 }
func genFn0() int {
    var v5 uint8 = 122
    var v6 uint32 = 260
    v7 := []int{2147483647, -491, 1}
    v6 = (v6 >> ((517) & 63))
    _ = v5
    _ = v6
    _ = v7
    return int(v5)
}

func main() {
    var v1 uint = 42
    var v2 uint64 = 4294967295
    var v3 int32 = -1
    var v4 uint16 = 65535
    fmt.Println(v2)
    fmt.Println(v4, v4, v4)
    if true {
        for fi8 := 0; fi8 < 5; fi8++ {
            switch 0 {
            case 2147483647:
                var v9 uint64 = (v2 + (v2 | 2))
                _ = v9
            }
            _ = fi8
        }
    } else {
        v1 += (v1 | (v1 ^ 18446744073709551615))
        var v10 int32 = (v3 | v3)
        _ = v10
    }
    switch "hello" {
    case "foo bar":
        v11 := []int{int(v3), 4294967295, 32767}
        _ = v11
    case "":
        var v12 uint32 = 2147483648
        _ = v12
    case "\n\t":
        v13 := GenS1{F0: 1000, F1: v1, F2: -492}
        _ = v13
    default:
        for fi14 := 0; fi14 < 16; fi14++ {
            if (true && ((v2 / ((v2) | 1)) < v2)) {
                switch fi14 {
                case -9223372036854775808:
                    v15 := GenS2{F0: (v2 >> ((v1) & 63))}
                    _ = v15
                case 32767:
                    v3 += v3
                case 42:
                    v16 := []int{fi14, (fi14 + (fi14 / ((fi14) | 1))), (fi14 >> ((784) & 63))}
                    _ = v16
                default:
                    v17 := GenS2{F0: (v2 - v2)}
                    _ = v17
                }
                switch "foo bar" {
                case "a":
                    var v18 uint16 = v4
                    _ = v18
                case "hello":
                    v3 -= (^(v3))
                default:
                    if false {
                        v3 -= v3
                        var v19 uint32 = 909
                        if true {
                            v20 := GenS3{F0: uint64(v1)}
                            _ = v20
                        }
                        _ = v19
                    } else {
                        if true {
                            if (127 != 42) {
                                if ((^(fi14)) < -9223372036854775808) {
                                    switch "a" {
                                    case "\n\t":
                                        var v21 int32 = (-(v3))
                                        _ = v21
                                    case "hello":
                                        v22 := GenS1{F0: 32767, F1: v1, F2: 1}
                                        _ = v22
                                    default:
                                        if true {
                                            var v23 uint64 = v2
                                            fmt.Println(v23, v23)
                                            _ = v23
                                        }
                                    }
                                    switch (fi14 % (((fi14 | fi14)) | 1)) {
                                    case 127:
                                        v24 := []int{fi14, genFn0(), (fi14 + fi14), fi14}
                                        _ = v24
                                    }
                                    if (false || true) {
                                        v25 := []int{}
                                        v4 = v4
                                        _ = v25
                                    }
                                } else {
                                    fmt.Println(v4, v1, v4)
                                }
                                if false {
                                    v26 := GenS3{F0: 2}
                                    _ = v26
                                } else {
                                    v27 := GenS3{F0: (v2 >> (((v1 << ((v1) & 63))) & 63))}
                                    _ = v27
                                }
                            }
                            v1 = 65535
                        }
                        v28 := []int{int(v4)}
                        _ = v28
                    }
                }
            } else {
                var v29 uint8 = 2
                switch (fi14 / ((genFn0()) | 1)) {
                case -506:
                    if (("hello" + "a") <= ("\n\t" + ("\n\t" + "hello"))) {
                        v30 := GenS4{F0: 255, F1: 1}
                        if true {
                            v29 = 0
                        } else {
                            for fi31 := 0; fi31 < 20; fi31++ {
                                v32 := GenS5{F0: genFn0(), F1: 381}
                                switch (fi31 ^ fi14) {
                                case 42:
                                    switch "" {
                                    case "foo bar":
                                        switch (fi14 & int(v3)) {
                                        case 972:
                                            if false {
                                                switch ("hello" + "") {
                                                case "foo bar":
                                                    v33 := GenS1{F0: 42, F1: v1, F2: -54}
                                                    _ = v33
                                                case "a":
                                                    if (true && true) {
                                                        if true {
                                                            var v34 uint16 = 2
                                                            if true {
                                                                switch genFn0() {
                                                                case 32768:
                                                                    switch "\n\t" {
                                                                    case "\n\t":
                                                                        v3 += -2147483648
                                                                    case "a":
                                                                        fmt.Println(fi14)
                                                                        var v35 uint16 = v34
                                                                        if (true && ((^(v2)) >= v2)) {
                                                                            v36 := GenS2{F0: v2}
                                                                            v37 := GenS5{F0: genFn0(), F1: int16(v3)}
                                                                            v38 := []int{fi14, fi14, (fi31 / (((fi14 ^ (fi31 % (((-(fi31))) | 1)))) | 1))}
                                                                            v38 = append(v38, -1)
                                                                            v3 = 1
                                                                            _ = v36
                                                                            _ = v37
                                                                            _ = v38
                                                                        } else {
                                                                            var v39 uint16 = (v4 % ((v4) | 1))
                                                                            var v40 uint64 = 131
                                                                            _ = v39
                                                                            _ = v40
                                                                        }
                                                                        _ = v35
                                                                    default:
                                                                        if false {
                                                                            v3 += -124
                                                                            switch (fi31 % (((-(fi14))) | 1)) {
                                                                            case 1:
                                                                                v1 -= v1
                                                                            }
                                                                            if (fi31 <= 556) {
                                                                                v41 := GenS4{F0: 0, F1: -100}
                                                                                v2 += v2
                                                                                _ = v41
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                                fmt.Println(v4, v1, v2)
                                                                v34 += (^(v34))
                                                            }
                                                            v2 = v2
                                                            _ = v34
                                                        } else {
                                                            switch "\n\t" {
                                                            case "a":
                                                                v42 := GenS2{F0: (v2 / ((0) | 1))}
                                                                _ = v42
                                                            case "hello":
                                                                fmt.Println(v29, v2)
                                                            }
                                                            fmt.Println(v3)
                                                        }
                                                        v43 := GenS3{F0: (v2 % ((v2) | 1))}
                                                        if false {
                                                            v44 := []int{(fi14 << ((2) & 63)), (fi31 ^ (fi31 | (^(fi14)))), (fi31 >> (((v1 | (v1 >> ((v1) & 63)))) & 63)), (fi31 / ((genFn0()) | 1))}
                                                            switch genFn0() {
                                                            case 127:
                                                                v3 += v3
                                                            }
                                                            var v45 uint = v1
                                                            _ = v44
                                                            _ = v45
                                                        } else {
                                                            var v46 bool = true
                                                            if (v29 != (v29 + (v29 % ((184) | 1)))) {
                                                                var v47 int8 = 51
                                                                v46 = (!v46)
                                                                var v48 bool = v46
                                                                var v49 float64 = 1.5
                                                                var v50 float32 = (0.1 * 0.0)
                                                                if false {
                                                                    if v48 {
                                                                        v2 += v2
                                                                        switch ("a" + "hello") {
                                                                        case "foo bar":
                                                                            if v46 {
                                                                                if (!(-323 > -1)) {
                                                                                    v51 := GenS3{F0: (v2 | v2)}
                                                                                    _ = v51
                                                                                } else {
                                                                                    fmt.Println(v4, v49, v47)
                                                                                    switch "foo bar" {
                                                                                    case "":
                                                                                        switch (fi14 + fi31) {
                                                                                        case 668:
                                                                                            v47 = (v47 | (v47 | v47))
                                                                                        }
                                                                                    case "hello":
                                                                                        if v48 {
                                                                                            switch "hello" {
                                                                                            case "a":
                                                                                                v52 := []int{(fi14 % ((int(v1)) | 1)), (fi14 << ((4294967295) & 63))}
                                                                                                _ = v52
                                                                                            case "":
                                                                                                switch fi31 {
                                                                                                case 2147483647:
                                                                                                    v53 := []int{(fi14 & fi14), (fi31 << ((uint(v2)) & 63)), 1}
                                                                                                    _ = v53
                                                                                                case 32768:
                                                                                                    v54 := GenS6{F0: 100, F1: (^(v1))}
                                                                                                    _ = v54
                                                                                                case -1:
                                                                                                    if v48 {
                                                                                                        fmt.Println(v3)
                                                                                                    }
                                                                                                }
                                                                                                switch (fi14 | -961) {
                                                                                                case 453:
                                                                                                    v3 = (v3 / ((v3) | 1))
                                                                                                case -1:
                                                                                                    switch genFn0() {
                                                                                                    case 37:
                                                                                                        v55 := GenS7{F0: -127, F1: v29, F2: "\n\t"}
                                                                                                        _ = v55
                                                                                                    case -531:
                                                                                                        var v56 int16 = 42
                                                                                                        _ = v56
                                                                                                    default:
                                                                                                        v57 := []int{int(v1), fi31, -534}
                                                                                                        _ = v57
                                                                                                    }
                                                                                                case 4294967295:
                                                                                                    v58 := GenS5{F0: genFn0(), F1: int16(v2)}
                                                                                                    _ = v58
                                                                                                }
                                                                                            case "foo bar":
                                                                                                var v59 string = "a"
                                                                                                _ = v59
                                                                                            }
                                                                                            if false {
                                                                                                switch "" {
                                                                                                case "hello":
                                                                                                    var v60 uint16 = v4
                                                                                                    _ = v60
                                                                                                case "foo bar":
                                                                                                    fmt.Println(v46, v29)
                                                                                                    v61 := GenS8{F0: (v3 | int32(v2)), F1: v29}
                                                                                                    _ = v61
                                                                                                }
                                                                                            } else {
                                                                                                var v62 int16 = 1
                                                                                                v63 := []int{(fi14 * -754), 128, (fi31 * fi31)}
                                                                                                _ = v62
                                                                                                _ = v63
                                                                                            }
                                                                                            v64 := []int{int(v47), 2147483647, fi31, int(v29)}
                                                                                            _ = v64
                                                                                        }
                                                                                        v4 = 0
                                                                                        var v65 int8 = (v47 - v47)
                                                                                        _ = v65
                                                                                    case "a":
                                                                                        var v66 float32 = 1.5
                                                                                        _ = v66
                                                                                    default:
                                                                                        v49 += v49
                                                                                        v67 := GenS5{F0: genFn0(), F1: -1}
                                                                                        var v68 int8 = v47
                                                                                        if v46 {
                                                                                            if ((^(v2)) == (v2 + v2)) {
                                                                                                fmt.Println(v49, v3)
                                                                                                var v69 uint32 = 2147483647
                                                                                                var v70 uint32 = v69
                                                                                                _ = v69
                                                                                                _ = v70
                                                                                            } else {
                                                                                                var v71 float64 = (((v49 * (1e10 / (v49 - v49))) / v49) * (v49 * v49))
                                                                                                _ = v71
                                                                                            }
                                                                                            v72 := []int{-1, (fi14 / (((fi14 % (((fi14 << ((4294967295) & 63))) | 1))) | 1)), (fi31 + genFn0()), (fi31 & fi14)}
                                                                                            _ = len(v72)
                                                                                            var v73 uint = (v1 * v1)
                                                                                            v2 += (^(v2))
                                                                                            fmt.Println(v29, v73)
                                                                                            _ = v72
                                                                                            _ = v73
                                                                                        }
                                                                                        _ = v67
                                                                                        _ = v68
                                                                                    }
                                                                                }
                                                                            }
                                                                        case "a":
                                                                            v4 -= (v4 & uint16(v2))
                                                                        case "hello":
                                                                            var v74 float32 = (v50 - v50)
                                                                            _ = v74
                                                                        default:
                                                                            v2 += 0
                                                                        }
                                                                    } else {
                                                                        if v46 {
                                                                            if v46 {
                                                                                var v75 int64 = 0
                                                                                if v46 {
                                                                                    var v76 int32 = (-(v3))
                                                                                    _ = v76
                                                                                } else {
                                                                                    fmt.Println(v3, v1, fi31)
                                                                                    v77 := GenS9{F0: (v47 ^ v47), F1: (fi14 / (((-(fi14))) | 1)), F2: int64(fi14)}
                                                                                    _ = v77
                                                                                }
                                                                                v47 += int8(fi14)
                                                                                _ = v75
                                                                            } else {
                                                                                var v78 int32 = 42
                                                                                _ = v78
                                                                            }
                                                                        } else {
                                                                            var v79 float64 = (v49 / (float64(v50) * (v49 / v49)))
                                                                            switch (fi31 | fi31) {
                                                                            case -396:
                                                                                v80 := []int{32768, 32767, (-(fi14)), -894}
                                                                                _ = v80
                                                                            }
                                                                            _ = v79
                                                                        }
                                                                        v81 := []int{(^(fi14)), (fi14 >> ((18446744073709551615) & 63))}
                                                                        var v82 int = (fi14 | fi14)
                                                                        _ = v81
                                                                        _ = v82
                                                                    }
                                                                }
                                                                _ = v47
                                                                _ = v48
                                                                _ = v49
                                                                _ = v50
                                                            } else {
                                                                fmt.Println(v3, fi31, fi14)
                                                                switch (fi14 << ((v1) & 63)) {
                                                                case 1:
                                                                    v83 := []int{fi31, genFn0()}
                                                                    _ = v83
                                                                default:
                                                                    var v84 uint64 = (^(v2))
                                                                    _ = v84
                                                                }
                                                            }
                                                            _ = v46
                                                        }
                                                        _ = v43
                                                    } else {
                                                        v85 := []int{(fi14 / ((fi31) | 1)), fi14, -2147483648, (fi31 % ((127) | 1))}
                                                        _ = v85
                                                    }
                                                case "\n\t":
                                                    v86 := GenS10{F0: -32768}
                                                    _ = v86
                                                default:
                                                    if true {
                                                        switch fi31 {
                                                        case 0:
                                                            if false {
                                                                v2 += v2
                                                            }
                                                        default:
                                                            v87 := []int{}
                                                            _ = v87
                                                        }
                                                    } else {
                                                        v88 := GenS11{F0: "a", F1: (float32(0.1) > (3.14159 - (((3.14159 + (0.1 * (-2.25 * -2.25))) / float32(3.14159)) - 1.5))), F2: v29}
                                                        v89 := GenS8{F0: v3, F1: v29}
                                                        _ = v88
                                                        _ = v89
                                                    }
                                                }
                                            } else {
                                                if true {
                                                    if true {
                                                        v3 = (v3 >> ((42) & 63))
                                                        v1 = v1
                                                    }
                                                } else {
                                                    v90 := GenS5{F0: fi31, F1: 902}
                                                    v91 := []int{(-(fi31)), (-(fi31)), (fi31 << ((v1) & 63))}
                                                    _ = v90
                                                    _ = v91
                                                }
                                            }
                                        case 52:
                                            if true {
                                                var v92 uint16 = 623
                                                _ = v92
                                            } else {
                                                v93 := GenS12{F0: "hello", F1: -215, F2: 128}
                                                var v94 uint = (v1 ^ (v1 - v1))
                                                var v95 int16 = -1
                                                switch (-(fi31)) {
                                                case -2147483648:
                                                    switch (^(fi31)) {
                                                    case -607:
                                                        if true {
                                                            switch genFn0() {
                                                            case 4294967295:
                                                                switch fi31 {
                                                                case 478:
                                                                    if false {
                                                                        if ((true && (!((1.5 / (3.14159 + 1.5)) != (1e10 * 1.5)))) && true) {
                                                                            switch "\n\t" {
                                                                            case "\n\t":
                                                                                v96 := GenS3{F0: v2}
                                                                                _ = v96
                                                                            default:
                                                                                v97 := []int{(fi31 / ((fi14) | 1)), -9223372036854775808, int(v2), int(v1)}
                                                                                _ = v97
                                                                            }
                                                                        } else {
                                                                            v1 = v1
                                                                            var v98 uint64 = v2
                                                                            var v99 int32 = v3
                                                                            v1 += (v94 / ((v1) | 1))
                                                                            _ = v98
                                                                            _ = v99
                                                                        }
                                                                        v100 := []int{int(v95), 32767, (fi31 >> ((v94) & 63)), genFn0()}
                                                                        _ = v100
                                                                    }
                                                                case -9223372036854775808:
                                                                    var v101 uint32 = 42
                                                                    _ = v101
                                                                case 1:
                                                                    fmt.Println(fi31)
                                                                default:
                                                                    var v102 int = 42
                                                                    _ = v102
                                                                }
                                                            case -2147483648:
                                                                var v103 uint = (v1 << ((v94) & 63))
                                                                _ = v103
                                                            default:
                                                                v3 = v3
                                                            }
                                                            var v104 int64 = -468
                                                            _ = v104
                                                        }
                                                    case 854:
                                                        switch fi31 {
                                                        case -1:
                                                            v29 += (v29 % ((v29) | 1))
                                                        }
                                                    default:
                                                        if ((-1 != -9223372036854775808) || (v29 >= v29)) {
                                                            if false {
                                                                if true {
                                                                    var v105 float32 = (3.14159 * (3.14159 / 1.5))
                                                                    var v106 uint64 = (^(v2))
                                                                    _ = v105
                                                                    _ = v106
                                                                } else {
                                                                    switch ("a" + "") {
                                                                    case "\n\t":
                                                                        if (825 == 485) {
                                                                            v4 = (^(v4))
                                                                            v107 := GenS12{F0: "foo bar", F1: 32767, F2: v29}
                                                                            var v108 uint16 = (v4 | v4)
                                                                            _ = v107
                                                                            _ = v108
                                                                        } else {
                                                                            v95 = v95
                                                                        }
                                                                    case "foo bar":
                                                                        if false {
                                                                            if (110 != 0) {
                                                                                var v109 float32 = 0.0
                                                                                var v110 int16 = (v95 + v95)
                                                                                if false {
                                                                                    if true {
                                                                                        fmt.Println(v29)
                                                                                    } else {
                                                                                        fmt.Println(v95)
                                                                                        var v111 int = genFn0()
                                                                                        _ = v111
                                                                                    }
                                                                                    switch ("" + "foo bar") {
                                                                                    case "":
                                                                                        if true {
                                                                                            switch "hello" {
                                                                                            case "a":
                                                                                                fmt.Println(v110, v95)
                                                                                            case "foo bar":
                                                                                                v112 := GenS2{F0: v2}
                                                                                                v94 -= v94
                                                                                                v113 := GenS7{F0: (v95 << ((v94) & 63)), F1: 42, F2: ""}
                                                                                                fmt.Println(v94, v4, v4)
                                                                                                v114 := []int{946}
                                                                                                _ = v112
                                                                                                _ = v113
                                                                                                _ = v114
                                                                                            case "\n\t":
                                                                                                var v115 float32 = -2.25
                                                                                                _ = v115
                                                                                            }
                                                                                            v116 := GenS2{F0: (v2 | 1)}
                                                                                            fmt.Println(v4, v1, v2)
                                                                                            _ = v116
                                                                                        }
                                                                                    case "a":
                                                                                        if (!(((v109 * 1e10) * v109) < 0.0)) {
                                                                                            switch (fi31 % (((^(fi14))) | 1)) {
                                                                                            case 1:
                                                                                                v117 := []int{(fi31 % ((-9223372036854775808) | 1))}
                                                                                                v118 := GenS2{F0: v2}
                                                                                                _ = v117
                                                                                                _ = v118
                                                                                            case 42:
                                                                                                if (0.1 < 3.14159) {
                                                                                                    v119 := []int{(fi31 & 9223372036854775807), int(v4), 32767}
                                                                                                    var v120 uint32 = 362
                                                                                                    if (!((!true) && true)) {
                                                                                                        switch "\n\t" {
                                                                                                        case "a":
                                                                                                            if true {
                                                                                                                if len(v119) > 0 {
                                                                                                                    v119[0] = (fi31 | (fi14 ^ int(v4)))
                                                                                                                }
                                                                                                            } else {
                                                                                                                var v121 float32 = (0.0 / v109)
                                                                                                                var v122 uint8 = (v29 / ((42) | 1))
                                                                                                                var v123 string = ""
                                                                                                                _ = v121
                                                                                                                _ = v122
                                                                                                                _ = v123
                                                                                                            }
                                                                                                        default:
                                                                                                            if (v29 != v29) {
                                                                                                                _ = len(v119)
                                                                                                            }
                                                                                                        }
                                                                                                        var v124 uint32 = v120
                                                                                                        switch -117 {
                                                                                                        case -9223372036854775808:
                                                                                                            if (!false) {
                                                                                                                var v125 uint8 = (v29 | (v29 >> ((v1) & 63)))
                                                                                                                _ = v125
                                                                                                            }
                                                                                                        }
                                                                                                        switch 127 {
                                                                                                        case 4294967295:
                                                                                                            v126 := GenS9{F0: 0, F1: (-(fi31)), F2: 0}
                                                                                                            _ = v126
                                                                                                        case -1:
                                                                                                            if true {
                                                                                                                v127 := GenS9{F0: -8, F1: int(v2), F2: -9223372036854775808}
                                                                                                                _ = v127
                                                                                                            } else {
                                                                                                                switch fi31 {
                                                                                                                case -9223372036854775808:
                                                                                                                    if false {
                                                                                                                        switch "\n\t" {
                                                                                                                        case "hello":
                                                                                                                            v3 -= (v3 ^ v3)
                                                                                                                        default:
                                                                                                                            v128 := GenS8{F0: (^(v3)), F1: (v29 - v29)}
                                                                                                                            _ = v128
                                                                                                                        }
                                                                                                                        var v129 int = 128
                                                                                                                        _ = v129
                                                                                                                    } else {
                                                                                                                        var v130 int8 = 127
                                                                                                                        fmt.Println(v3, v4, v94)
                                                                                                                        _ = v130
                                                                                                                    }
                                                                                                                    fmt.Println(v94)
                                                                                                                    v110 -= (v110 | (v110 * int16(v2)))
                                                                                                                }
                                                                                                            }
                                                                                                        case 1:
                                                                                                            var v131 int16 = (v110 ^ v95)
                                                                                                            _ = v131
                                                                                                        }
                                                                                                        var v132 int8 = -128
                                                                                                        v133 := GenS2{F0: (v2 % (((^(v2))) | 1))}
                                                                                                        _ = v124
                                                                                                        _ = v132
                                                                                                        _ = v133
                                                                                                    } else {
                                                                                                        if (false || (1e10 >= 1.5)) {
                                                                                                            v119 = append(v119, (fi31 >> (((v94 - v1)) & 63)))
                                                                                                            v95 = -32768
                                                                                                            var v134 bool = (false && false)
                                                                                                            if v134 {
                                                                                                                fmt.Println(v29, v29)
                                                                                                                switch 178 {
                                                                                                                case 2147483647:
                                                                                                                    var v135 uint64 = (v2 % ((0) | 1))
                                                                                                                    _ = v135
                                                                                                                case -370:
                                                                                                                    var v136 int8 = -1
                                                                                                                    _ = v136
                                                                                                                case -606:
                                                                                                                    var v137 uint64 = (^(v2))
                                                                                                                    _ = v137
                                                                                                                }
                                                                                                                var v138 int8 = 127
                                                                                                                _ = v138
                                                                                                            } else {
                                                                                                                v139 := GenS13{F0: (v29 - v29), F1: (v120 | v120)}
                                                                                                                if len(v119) >= 2 {
                                                                                                                    v119 = v119[1:len(v119)]
                                                                                                                }
                                                                                                                _ = v139
                                                                                                            }
                                                                                                            _ = v134
                                                                                                        } else {
                                                                                                            var v140 uint32 = 65535
                                                                                                            v141 := GenS10{F0: v110}
                                                                                                            if len(v119) >= 2 {
                                                                                                                v119 = v119[1:len(v119)]
                                                                                                            }
                                                                                                            _ = v140
                                                                                                            _ = v141
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v119
                                                                                                    _ = v120
                                                                                                }
                                                                                            }
                                                                                            var v142 int = fi31
                                                                                            _ = v142
                                                                                        } else {
                                                                                            if (false || true) {
                                                                                                var v143 bool = (true || ("foo bar" != ("\n\t" + ("foo bar" + "hello"))))
                                                                                                v109 -= v109
                                                                                                _ = v143
                                                                                            } else {
                                                                                                var v144 int = -1
                                                                                                switch fi14 {
                                                                                                case 730:
                                                                                                    switch fi14 {
                                                                                                    case 9223372036854775807:
                                                                                                        v2 = (v2 + v2)
                                                                                                    case 1:
                                                                                                        var v145 int16 = (v110 | (v95 / (((-(v110))) | 1)))
                                                                                                        _ = v145
                                                                                                    }
                                                                                                default:
                                                                                                    if false {
                                                                                                        if (true || (65535 != (v4 ^ (v4 | v4)))) {
                                                                                                            switch "foo bar" {
                                                                                                            case "\n\t":
                                                                                                                switch ("" + ("a" + ("\n\t" + "hello"))) {
                                                                                                                case "\n\t":
                                                                                                                    fmt.Println(v29, v3)
                                                                                                                    v1 += uint(v2)
                                                                                                                    if false {
                                                                                                                        if true {
                                                                                                                            v146 := []int{fi31, -140}
                                                                                                                            var v147 bool = (0.0 == 1e10)
                                                                                                                            _ = v146
                                                                                                                            _ = v147
                                                                                                                        } else {
                                                                                                                            v148 := []int{fi31, genFn0(), 128, 42}
                                                                                                                            _ = v148
                                                                                                                        }
                                                                                                                    }
                                                                                                                    fmt.Println(fi14, v94, v94)
                                                                                                                case "hello":
                                                                                                                    v149 := GenS14{F0: v2}
                                                                                                                    if (v2 < (v2 | (v2 % ((v2) | 1)))) {
                                                                                                                        switch "hello" {
                                                                                                                        case "hello":
                                                                                                                            if false {
                                                                                                                                var v150 int = (fi14 / ((1) | 1))
                                                                                                                                switch -177 {
                                                                                                                                case 911:
                                                                                                                                    v144 += (^(fi14))
                                                                                                                                case 127:
                                                                                                                                    fmt.Println(v144, v29)
                                                                                                                                    v151 := GenS15{F0: 0}
                                                                                                                                    v144 -= v144
                                                                                                                                    v2 = v2
                                                                                                                                    v110 -= 42
                                                                                                                                    var v152 uint16 = 1
                                                                                                                                    _ = v151
                                                                                                                                    _ = v152
                                                                                                                                case 0:
                                                                                                                                    if false {
                                                                                                                                        v153 := GenS16{F0: (^(v2))}
                                                                                                                                        _ = v153
                                                                                                                                    } else {
                                                                                                                                        if (((true && (false && (!false))) && ((^(v94)) <= 42)) || (-128 != 0)) {
                                                                                                                                            v154 := GenS1{F0: 531, F1: (v1 >> ((158) & 63)), F2: 153}
                                                                                                                                            var v155 uint8 = (v29 | (^(v29)))
                                                                                                                                            var v156 float32 = 3.14159
                                                                                                                                            _ = v154
                                                                                                                                            _ = v155
                                                                                                                                            _ = v156
                                                                                                                                        } else {
                                                                                                                                            v157 := []int{(v144 - genFn0()), v150}
                                                                                                                                            v144 = v150
                                                                                                                                            _ = v157
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                default:
                                                                                                                                    if ((((0.0 * (v109 * v109)) / -2.25) * 3.14159) == 1e10) {
                                                                                                                                        if (!(v2 != 0)) {
                                                                                                                                            var v158 uint32 = 2147483648
                                                                                                                                            fmt.Println(v1, v110)
                                                                                                                                            _ = v158
                                                                                                                                        } else {
                                                                                                                                            var v159 string = "hello"
                                                                                                                                            _ = v159
                                                                                                                                        }
                                                                                                                                    } else {
                                                                                                                                        v160 := []int{-925, (v150 + (fi14 + fi14)), genFn0()}
                                                                                                                                        if len(v160) >= 2 {
                                                                                                                                            v160 = v160[1:len(v160)]
                                                                                                                                        }
                                                                                                                                        _ = v160
                                                                                                                                    }
                                                                                                                                }
                                                                                                                                var v161 bool = true
                                                                                                                                _ = v150
                                                                                                                                _ = v161
                                                                                                                            }
                                                                                                                        case "foo bar":
                                                                                                                            switch "" {
                                                                                                                            case "a":
                                                                                                                                var v162 uint16 = 32767
                                                                                                                                _ = v162
                                                                                                                            default:
                                                                                                                                if (-1 <= 259) {
                                                                                                                                    v95 = v95
                                                                                                                                    v163 := []int{(fi14 % (((fi14 / (((fi14 >> ((v94) & 63))) | 1))) | 1)), 1, -865, -466}
                                                                                                                                    if true {
                                                                                                                                        fmt.Println(v29, v110, v94)
                                                                                                                                        if len(v163) > 0 {
                                                                                                                                            v163[0] = v144
                                                                                                                                        }
                                                                                                                                        var v164 uint8 = v29
                                                                                                                                        _ = v164
                                                                                                                                    } else {
                                                                                                                                        switch "" {
                                                                                                                                        case "hello":
                                                                                                                                            switch -445 {
                                                                                                                                            case 1:
                                                                                                                                                if len(v163) > 0 {
                                                                                                                                                    v163[0] = (-(fi31))
                                                                                                                                                }
                                                                                                                                            case -2147483648:
                                                                                                                                                var v165 uint64 = v2
                                                                                                                                                _ = v165
                                                                                                                                            case 604:
                                                                                                                                                v166 := GenS17{F0: ((v109 * 0.1) != (v109 + (0.1 - v109)))}
                                                                                                                                                _ = v166
                                                                                                                                            default:
                                                                                                                                                v167 := GenS18{F0: (v144 % ((v144) | 1)), F1: true}
                                                                                                                                                var v168 float32 = v109
                                                                                                                                                _ = v167
                                                                                                                                                _ = v168
                                                                                                                                            }
                                                                                                                                            switch v144 {
                                                                                                                                            case -1:
                                                                                                                                                var v169 uint32 = 172
                                                                                                                                                _ = v169
                                                                                                                                            case -253:
                                                                                                                                                var v170 uint64 = (^(v2))
                                                                                                                                                _ = v170
                                                                                                                                            default:
                                                                                                                                                var v171 float32 = (v109 * -2.25)
                                                                                                                                                _ = v171
                                                                                                                                            }
                                                                                                                                        case "\n\t":
                                                                                                                                            switch (("foo bar" + "\n\t") + ("" + ("hello" + "hello"))) {
                                                                                                                                            case "a":
                                                                                                                                                v4 = (v4 | v4)
                                                                                                                                            }
                                                                                                                                        }
                                                                                                                                        v3 = v3
                                                                                                                                        fmt.Println(v4)
                                                                                                                                    }
                                                                                                                                    var v172 float64 = (1.5 / (0.0 - 3.14159))
                                                                                                                                    if (127 != int8(v94)) {
                                                                                                                                        fmt.Println(v94, v110, fi31)
                                                                                                                                    } else {
                                                                                                                                        switch ("a" + "foo bar") {
                                                                                                                                        case "":
                                                                                                                                            var v173 string = ""
                                                                                                                                            _ = v173
                                                                                                                                        }
                                                                                                                                        v174 := GenS19{F0: 982}
                                                                                                                                        _ = v174
                                                                                                                                    }
                                                                                                                                    v95 = v95
                                                                                                                                    _ = v163
                                                                                                                                    _ = v172
                                                                                                                                } else {
                                                                                                                                    var v175 int64 = 0
                                                                                                                                    _ = v175
                                                                                                                                }
                                                                                                                                v176 := []int{2147483647, -1}
                                                                                                                                _ = v176
                                                                                                                            }
                                                                                                                        default:
                                                                                                                            v177 := GenS20{F0: v110}
                                                                                                                            fmt.Println(v95)
                                                                                                                            v178 := GenS21{F0: (v3 * (v3 * (v3 ^ 426))), F1: 1, F2: (v110 % (((-(v110))) | 1))}
                                                                                                                            fmt.Println(v1)
                                                                                                                            if true {
                                                                                                                                var v179 string = "hello"
                                                                                                                                v3 -= 160
                                                                                                                                v180 := []int{(v144 - (v144 | fi31)), fi14}
                                                                                                                                _ = v179
                                                                                                                                _ = v180
                                                                                                                            } else {
                                                                                                                                var v181 int32 = int32(v1)
                                                                                                                                _ = v181
                                                                                                                            }
                                                                                                                            var v182 int64 = -9223372036854775808
                                                                                                                            _ = v177
                                                                                                                            _ = v178
                                                                                                                            _ = v182
                                                                                                                        }
                                                                                                                        var v183 uint8 = v29
                                                                                                                        v109 = (0.0 - -2.25)
                                                                                                                        v184 := GenS18{F0: 42, F1: (!false)}
                                                                                                                        _ = v183
                                                                                                                        _ = v184
                                                                                                                    } else {
                                                                                                                        v95 += v95
                                                                                                                        v110 = (v110 >> (((v94 ^ v1)) & 63))
                                                                                                                        v185 := []int{4294967295, fi14, genFn0()}
                                                                                                                        _ = v185
                                                                                                                    }
                                                                                                                    _ = v149
                                                                                                                }
                                                                                                            default:
                                                                                                                v186 := GenS22{F0: 42, F1: 42}
                                                                                                                _ = v186
                                                                                                            }
                                                                                                            if ((v94 & v1) <= 0) {
                                                                                                                v187 := GenS23{F0: uint32(v94), F1: -718, F2: 41}
                                                                                                                _ = v187
                                                                                                            } else {
                                                                                                                v188 := []int{4294967295}
                                                                                                                var v189 uint = (^(v94))
                                                                                                                _ = v188
                                                                                                                _ = v189
                                                                                                            }
                                                                                                            v190 := []int{(v144 * fi31)}
                                                                                                            var v191 float64 = -2.25
                                                                                                            var v192 float64 = 0.1
                                                                                                            _ = v190
                                                                                                            _ = v191
                                                                                                            _ = v192
                                                                                                        }
                                                                                                        switch "" {
                                                                                                        case "hello":
                                                                                                            v193 := GenS24{F0: (v110 >> ((v1) & 63)), F1: genFn0()}
                                                                                                            _ = v193
                                                                                                        case "a":
                                                                                                            switch ("\n\t" + "\n\t") {
                                                                                                            case "foo bar":
                                                                                                                var v194 int8 = -128
                                                                                                                _ = v194
                                                                                                            case "hello":
                                                                                                                switch "\n\t" {
                                                                                                                case "hello":
                                                                                                                    var v195 int8 = -80
                                                                                                                    _ = v195
                                                                                                                case "\n\t":
                                                                                                                    v196 := []int{(fi14 & (fi14 | 173)), (fi14 & genFn0()), v144, (fi31 ^ (^(v144)))}
                                                                                                                    _ = v196
                                                                                                                }
                                                                                                            case "\n\t":
                                                                                                                fmt.Println(fi14, fi14, v1)
                                                                                                            default:
                                                                                                                switch ("foo bar" + "a") {
                                                                                                                case "\n\t":
                                                                                                                    if (false && (-441 >= 0)) {
                                                                                                                        var v197 uint8 = v29
                                                                                                                        v144 += (fi31 + fi31)
                                                                                                                        _ = v197
                                                                                                                    } else {
                                                                                                                        if ((v3 >= (v3 << (((v1 | 1)) & 63))) && ((v29 ^ v29) != 26)) {
                                                                                                                            if true {
                                                                                                                                v110 += v95
                                                                                                                                switch v144 {
                                                                                                                                case -709:
                                                                                                                                    v1 += 574
                                                                                                                                default:
                                                                                                                                    v198 := []int{273, (fi14 % (((fi14 * fi31)) | 1))}
                                                                                                                                    _ = v198
                                                                                                                                }
                                                                                                                                if false {
                                                                                                                                    var v199 uint64 = (v2 | 18446744073709551615)
                                                                                                                                    _ = v199
                                                                                                                                } else {
                                                                                                                                    var v200 int16 = v95
                                                                                                                                    if false {
                                                                                                                                        var v201 int32 = (v3 ^ (^(v3)))
                                                                                                                                        v202 := GenS25{F0: v29}
                                                                                                                                        var v203 int32 = -349
                                                                                                                                        _ = v201
                                                                                                                                        _ = v202
                                                                                                                                        _ = v203
                                                                                                                                    } else {
                                                                                                                                        v204 := GenS26{F0: -936}
                                                                                                                                        v205 := []int{-511, (fi31 % ((-787) | 1))}
                                                                                                                                        _ = len(v205)
                                                                                                                                        _ = v204
                                                                                                                                        _ = v205
                                                                                                                                    }
                                                                                                                                    _ = v200
                                                                                                                                }
                                                                                                                            } else {
                                                                                                                                var v206 uint64 = (v2 & 610)
                                                                                                                                _ = v206
                                                                                                                            }
                                                                                                                            v207 := []int{1, 405, genFn0(), 1}
                                                                                                                            _ = v207
                                                                                                                        }
                                                                                                                        v95 += (v110 & v95)
                                                                                                                        v208 := []int{}
                                                                                                                        if (v29 > v29) {
                                                                                                                            switch 32767 {
                                                                                                                            case 32767:
                                                                                                                                if len(v208) > 0 {
                                                                                                                                    v208[0] = fi14
                                                                                                                                }
                                                                                                                            default:
                                                                                                                                var v209 uint8 = (v29 ^ 128)
                                                                                                                                _ = v209
                                                                                                                            }
                                                                                                                            var v210 float64 = 3.14159
                                                                                                                            v94 += v94
                                                                                                                            _ = v210
                                                                                                                        }
                                                                                                                        _ = v208
                                                                                                                    }
                                                                                                                case "a":
                                                                                                                    fmt.Println(v94, v94)
                                                                                                                default:
                                                                                                                    switch "hello" {
                                                                                                                    case "":
                                                                                                                        v4 += (v4 << ((v1) & 63))
                                                                                                                    case "\n\t":
                                                                                                                        switch "foo bar" {
                                                                                                                        case "foo bar":
                                                                                                                            fmt.Println(fi14, v1)
                                                                                                                        default:
                                                                                                                            v211 := []int{32768}
                                                                                                                            _ = v211
                                                                                                                        }
                                                                                                                    case "a":
                                                                                                                        v212 := GenS23{F0: 358, F1: 246, F2: 1}
                                                                                                                        _ = v212
                                                                                                                    default:
                                                                                                                        var v213 int8 = 42
                                                                                                                        _ = v213
                                                                                                                    }
                                                                                                                }
                                                                                                            }
                                                                                                        case "":
                                                                                                            switch "a" {
                                                                                                            case "a":
                                                                                                                fmt.Println(fi31, v29, v1)
                                                                                                                v214 := []int{(fi31 >> (((v1 % ((v94) | 1))) & 63)), fi14, genFn0(), (-(fi31))}
                                                                                                                if len(v214) > 0 {
                                                                                                                    v214[0] = fi31
                                                                                                                }
                                                                                                                _ = v214
                                                                                                            }
                                                                                                        default:
                                                                                                            v215 := GenS8{F0: v3, F1: v29}
                                                                                                            _ = v215
                                                                                                        }
                                                                                                        v1 += (v1 >> ((v1) & 63))
                                                                                                        var v216 bool = true
                                                                                                        switch genFn0() {
                                                                                                        case 143:
                                                                                                            v217 := []int{v144}
                                                                                                            _ = v217
                                                                                                        case -1:
                                                                                                            v218 := GenS27{F0: 65535, F1: (fi31 << (((v1 & v94)) & 63))}
                                                                                                            _ = v218
                                                                                                        }
                                                                                                        _ = v216
                                                                                                    }
                                                                                                }
                                                                                                _ = v144
                                                                                            }
                                                                                        }
                                                                                    case "foo bar":
                                                                                        if ((v95 + (v110 % (((v110 << ((v1) & 63))) | 1))) == (v95 << ((v94) & 63))) {
                                                                                            v94 += (v94 & v1)
                                                                                        }
                                                                                    }
                                                                                    var v219 uint64 = v2
                                                                                    if (!false) {
                                                                                        v1 = v1
                                                                                        var v220 int = 136
                                                                                        var v221 bool = false
                                                                                        _ = v220
                                                                                        _ = v221
                                                                                    }
                                                                                    _ = v219
                                                                                } else {
                                                                                    v222 := []int{fi14, fi14, fi31}
                                                                                    _ = v222
                                                                                }
                                                                                _ = v109
                                                                                _ = v110
                                                                            } else {
                                                                                v223 := GenS16{F0: v2}
                                                                                fmt.Println(v94)
                                                                                v29 = (v29 ^ uint8(v94))
                                                                                v95 += (^(v95))
                                                                                v1 = 278
                                                                                _ = v223
                                                                            }
                                                                            var v224 float64 = float64(3.14159)
                                                                            _ = v224
                                                                        }
                                                                    case "":
                                                                        if (!(255 >= (v1 | 171))) {
                                                                            if true {
                                                                                v29 -= 127
                                                                                var v225 uint16 = (v4 << (((v94 - uint(v2))) & 63))
                                                                                _ = v225
                                                                            } else {
                                                                                if true {
                                                                                    v226 := []int{int(v29), (fi31 & (^(fi14))), genFn0(), int(v2)}
                                                                                    _ = v226
                                                                                } else {
                                                                                    fmt.Println(v95, v1)
                                                                                    switch "\n\t" {
                                                                                    case "\n\t":
                                                                                        var v227 float32 = (1e10 - 3.14159)
                                                                                        _ = v227
                                                                                    }
                                                                                    v228 := []int{0, (fi14 / ((178) | 1)), (fi14 / ((-2147483648) | 1))}
                                                                                    switch fi31 {
                                                                                    case 9223372036854775807:
                                                                                        v229 := GenS28{F0: -710, F1: (v95 - (v95 % ((v95) | 1))), F2: 42}
                                                                                        _ = v229
                                                                                    }
                                                                                    _ = v228
                                                                                }
                                                                                switch (fi14 ^ (fi31 + genFn0())) {
                                                                                case 9223372036854775807:
                                                                                    v230 := []int{fi14, 4294967295, (fi14 & (^(fi31))), (fi14 >> (((^(v94))) & 63))}
                                                                                    _ = v230
                                                                                case 127:
                                                                                    v231 := GenS16{F0: v2}
                                                                                    if (2 > 2147483647) {
                                                                                        switch genFn0() {
                                                                                        case -1:
                                                                                            switch (fi31 * (fi14 % (((fi31 & (fi14 * int(v4)))) | 1))) {
                                                                                            case 0:
                                                                                                v232 := GenS29{F0: 2147483647, F1: v95}
                                                                                                _ = v232
                                                                                            }
                                                                                        case 42:
                                                                                            fmt.Println(v95, fi31, v94)
                                                                                        case 2147483647:
                                                                                            fmt.Println(v3, v94, v94)
                                                                                        }
                                                                                        v233 := GenS30{F0: (v2 * v2)}
                                                                                        v3 = (v3 & (v3 - v3))
                                                                                        _ = v233
                                                                                    }
                                                                                    v234 := []int{fi14}
                                                                                    _ = v231
                                                                                    _ = v234
                                                                                case 820:
                                                                                    if true {
                                                                                        var v235 uint = v94
                                                                                        v236 := []int{fi14, genFn0()}
                                                                                        v235 -= 958
                                                                                        _ = v235
                                                                                        _ = v236
                                                                                    } else {
                                                                                        if (fi31 != (fi14 ^ 0)) {
                                                                                            if (true && (float64(float32(1e10)) == 3.14159)) {
                                                                                                fmt.Println(v3, v1)
                                                                                                v2 -= 42
                                                                                                var v237 string = "hello"
                                                                                                _ = v237
                                                                                            } else {
                                                                                                var v238 uint64 = (v2 / (((v2 - uint64(v29))) | 1))
                                                                                                _ = v238
                                                                                            }
                                                                                            v239 := GenS30{F0: (v2 * (^(v2)))}
                                                                                            if (true || false) {
                                                                                                v95 -= (-(v95))
                                                                                            }
                                                                                            _ = v239
                                                                                        } else {
                                                                                            v240 := GenS6{F0: 1, F1: (v1 >> ((v94) & 63))}
                                                                                            v29 += v29
                                                                                            if true {
                                                                                                var v241 int8 = 0
                                                                                                switch "hello" {
                                                                                                case "foo bar":
                                                                                                    if ((fi14 >> ((18446744073709551615) & 63)) != (fi14 | (fi14 * -232))) {
                                                                                                        if true {
                                                                                                            var v242 float32 = (1e10 + 0.1)
                                                                                                            var v243 uint32 = 2147483648
                                                                                                            _ = v242
                                                                                                            _ = v243
                                                                                                        } else {
                                                                                                            v244 := []int{fi14, 451, (fi31 - 327)}
                                                                                                            _ = v244
                                                                                                        }
                                                                                                        v245 := GenS12{F0: ("" + "a"), F1: fi31, F2: v29}
                                                                                                        _ = v245
                                                                                                    }
                                                                                                default:
                                                                                                    var v246 bool = (true || true)
                                                                                                    _ = v246
                                                                                                }
                                                                                                v247 := []int{127, 128, int(v29), genFn0()}
                                                                                                _ = v241
                                                                                                _ = v247
                                                                                            } else {
                                                                                                v248 := GenS31{F0: 42}
                                                                                                var v249 bool = true
                                                                                                v2 = 1
                                                                                                _ = v248
                                                                                                _ = v249
                                                                                            }
                                                                                            _ = v240
                                                                                        }
                                                                                        var v250 int64 = 42
                                                                                        _ = v250
                                                                                    }
                                                                                    if true {
                                                                                        var v251 uint = (v1 | 0)
                                                                                        v29 += (v29 / (((v29 + (^(v29)))) | 1))
                                                                                        if (int(v95) < (^(fi31))) {
                                                                                            switch (fi14 / (((fi31 | fi14)) | 1)) {
                                                                                            case -1:
                                                                                                if false {
                                                                                                    var v252 float32 = 1.5
                                                                                                    fmt.Println(v251, v252, v29)
                                                                                                    _ = v252
                                                                                                }
                                                                                            case -2147483648:
                                                                                                var v253 int32 = (v3 | (v3 + 916))
                                                                                                _ = v253
                                                                                            }
                                                                                            var v254 bool = false
                                                                                            var v255 int = (fi14 / (((^(fi31))) | 1))
                                                                                            _ = v254
                                                                                            _ = v255
                                                                                        }
                                                                                        _ = v251
                                                                                    } else {
                                                                                        v256 := []int{genFn0(), int(v29)}
                                                                                        if len(v256) >= 2 {
                                                                                            v256 = v256[1:len(v256)]
                                                                                        }
                                                                                        var v257 int32 = (v3 + (v3 - v3))
                                                                                        _ = v256
                                                                                        _ = v257
                                                                                    }
                                                                                    v258 := []int{genFn0(), (fi31 << ((99) & 63)), (fi31 << (((^(v1))) & 63)), 128}
                                                                                    _ = v258
                                                                                default:
                                                                                    if (true && ((v29 >> ((uint(v4)) & 63)) > 128)) {
                                                                                        v3 -= (v3 * (v3 % (((v3 >> (((v94 / (((^(v94))) | 1))) & 63))) | 1)))
                                                                                    } else {
                                                                                        v95 = (^(v95))
                                                                                        v259 := GenS15{F0: v4}
                                                                                        _ = v259
                                                                                    }
                                                                                }
                                                                            }
                                                                            switch -43 {
                                                                            case -9223372036854775808:
                                                                                v260 := GenS9{F0: -1, F1: fi14, F2: 9223372036854775807}
                                                                                _ = v260
                                                                            case -295:
                                                                                var v261 float64 = 1e10
                                                                                _ = v261
                                                                            case 9223372036854775807:
                                                                                v3 -= -757
                                                                            }
                                                                        }
                                                                    }
                                                                    if (genFn0() < (fi31 + 0)) {
                                                                        v4 -= (v4 * (v4 / (((v4 - v4)) | 1)))
                                                                        if ("\n\t" <= "\n\t") {
                                                                            fmt.Println(v4, v95, v3)
                                                                            if true {
                                                                                if (1 >= -623) {
                                                                                    v4 += v4
                                                                                } else {
                                                                                    var v262 string = "foo bar"
                                                                                    switch v262 {
                                                                                    case "\n\t":
                                                                                        var v263 uint8 = (v29 ^ v29)
                                                                                        _ = v263
                                                                                    case "hello":
                                                                                        v264 := GenS2{F0: (v2 >> (((^(v1))) & 63))}
                                                                                        _ = v264
                                                                                    case "a":
                                                                                        var v265 int64 = 1
                                                                                        _ = v265
                                                                                    default:
                                                                                        var v266 uint = 42
                                                                                        _ = v266
                                                                                    }
                                                                                    _ = v262
                                                                                }
                                                                                v267 := []int{fi31, genFn0(), fi31}
                                                                                _ = v267
                                                                            } else {
                                                                                if ((v3 - (v3 >> ((v1) & 63))) < 42) {
                                                                                    v268 := []int{(fi31 % ((genFn0()) | 1)), (fi31 * fi14), 42, genFn0()}
                                                                                    var v269 float64 = 3.14159
                                                                                    _ = v268
                                                                                    _ = v269
                                                                                } else {
                                                                                    if true {
                                                                                        var v270 float32 = 3.14159
                                                                                        _ = v270
                                                                                    }
                                                                                    var v271 uint64 = 42
                                                                                    _ = v271
                                                                                }
                                                                            }
                                                                        }
                                                                        var v272 int = genFn0()
                                                                        _ = v272
                                                                    }
                                                                    fmt.Println(v95, v95)
                                                                }
                                                                v29 = v29
                                                            }
                                                            if false {
                                                                if (!(!true)) {
                                                                    var v273 int32 = -684
                                                                    switch "\n\t" {
                                                                    case "a":
                                                                        var v274 int16 = 42
                                                                        _ = v274
                                                                    case "":
                                                                        v4 -= (v4 & (v4 + (v4 - 65535)))
                                                                        if (0.1 < 1.5) {
                                                                            switch "a" {
                                                                            case "\n\t":
                                                                                var v275 int16 = v95
                                                                                _ = v275
                                                                            case "":
                                                                                var v276 int64 = -469
                                                                                _ = v276
                                                                            case "hello":
                                                                                v3 = v3
                                                                                v94 = (v94 / (((^(v94))) | 1))
                                                                                v277 := GenS32{F0: 42}
                                                                                _ = v277
                                                                            }
                                                                            v4 -= (v4 % ((v4) | 1))
                                                                        } else {
                                                                            var v278 uint8 = 255
                                                                            var v279 uint = v1
                                                                            _ = v278
                                                                            _ = v279
                                                                        }
                                                                        if (int64(v1) <= 1) {
                                                                            v280 := []int{fi31}
                                                                            v94 -= 42
                                                                            _ = v280
                                                                        }
                                                                    default:
                                                                        if ((true && (((3.14159 + -2.25) * float64(0.0)) == (((-2.25 + ((0.1 + ((1e10 * (float64(1.5) + float64(-2.25))) + (float64(3.14159) + (1e10 * 3.14159)))) / float64(0.1))) * -2.25) * 0.1))) && false) {
                                                                            if false {
                                                                                v94 = uint(v29)
                                                                            }
                                                                            if false {
                                                                                switch "a" {
                                                                                case "foo bar":
                                                                                    if true {
                                                                                        if ("foo bar" <= ("\n\t" + (("foo bar" + "hello") + "a"))) {
                                                                                            if (-128 >= 0) {
                                                                                                v2 = v2
                                                                                                var v281 uint8 = v29
                                                                                                _ = v281
                                                                                            } else {
                                                                                                v3 = -2147483648
                                                                                                if false {
                                                                                                    v29 -= 2
                                                                                                    switch "hello" {
                                                                                                    case "hello":
                                                                                                        v282 := GenS25{F0: 255}
                                                                                                        _ = v282
                                                                                                    case "foo bar":
                                                                                                        v95 = v95
                                                                                                    default:
                                                                                                        if true {
                                                                                                            if (false || ((v3 + (v273 - v273)) != (v3 % ((int32(v95)) | 1)))) {
                                                                                                                var v283 uint32 = 516
                                                                                                                _ = v283
                                                                                                            } else {
                                                                                                                v284 := GenS33{F0: 9223372036854775807}
                                                                                                                _ = v284
                                                                                                            }
                                                                                                            v285 := []int{genFn0(), (^(fi31)), (fi31 / (((fi14 % (((fi14 + -1)) | 1))) | 1)), genFn0()}
                                                                                                            switch (fi31 ^ fi14) {
                                                                                                            case 127:
                                                                                                                var v286 string = ("hello" + "foo bar")
                                                                                                                _ = v286
                                                                                                            case 4294967295:
                                                                                                                var v287 int8 = -1
                                                                                                                _ = v287
                                                                                                            }
                                                                                                            _ = v285
                                                                                                        }
                                                                                                    }
                                                                                                    if ((fi31 - genFn0()) != genFn0()) {
                                                                                                        v288 := []int{(-(fi14)), fi14, fi31, (fi31 % ((fi31) | 1))}
                                                                                                        switch ("\n\t" + "\n\t") {
                                                                                                        case "":
                                                                                                            var v289 float64 = 3.14159
                                                                                                            _ = v289
                                                                                                        }
                                                                                                        _ = v288
                                                                                                    } else {
                                                                                                        if ((!((v95 ^ (v95 / (((^(v95))) | 1))) == 42)) && (3.14159 < (3.14159 * ((0.1 / (float64(0.1) + 1.5)) * 0.0)))) {
                                                                                                            v290 := []int{128}
                                                                                                            v4 -= (v4 << (((v94 - v94)) & 63))
                                                                                                            v29 -= 23
                                                                                                            _ = v290
                                                                                                        } else {
                                                                                                            if ((v29 / (((v29 * 255)) | 1)) <= v29) {
                                                                                                                var v291 bool = (true && true)
                                                                                                                switch "foo bar" {
                                                                                                                case "":
                                                                                                                    switch "a" {
                                                                                                                    case "foo bar":
                                                                                                                        var v292 uint64 = (v2 / (((^(v2))) | 1))
                                                                                                                        _ = v292
                                                                                                                    default:
                                                                                                                        if (17 <= 2147483647) {
                                                                                                                            var v293 int32 = v273
                                                                                                                            var v294 float64 = -2.25
                                                                                                                            _ = v293
                                                                                                                            _ = v294
                                                                                                                        }
                                                                                                                    }
                                                                                                                case "foo bar":
                                                                                                                    switch (("foo bar" + "\n\t") + "") {
                                                                                                                    case "hello":
                                                                                                                        var v295 uint32 = 295
                                                                                                                        _ = v295
                                                                                                                    case "foo bar":
                                                                                                                        v296 := GenS34{F0: 2, F1: (v94 | (v94 ^ (v94 / ((v1) | 1)))), F2: 255}
                                                                                                                        _ = v296
                                                                                                                    }
                                                                                                                default:
                                                                                                                    v1 += v1
                                                                                                                }
                                                                                                                _ = v291
                                                                                                            } else {
                                                                                                                if true {
                                                                                                                    var v297 uint8 = v29
                                                                                                                    fmt.Println(v297, v95)
                                                                                                                    v298 := []int{fi31, fi31, (fi14 % ((0) | 1)), genFn0()}
                                                                                                                    _ = v297
                                                                                                                    _ = v298
                                                                                                                }
                                                                                                            }
                                                                                                            v299 := GenS35{F0: v2, F1: 1, F2: 900}
                                                                                                            v300 := GenS31{F0: uint32(v1)}
                                                                                                            _ = v299
                                                                                                            _ = v300
                                                                                                        }
                                                                                                    }
                                                                                                } else {
                                                                                                    var v301 int32 = v273
                                                                                                    v4 = (^(v4))
                                                                                                    _ = v301
                                                                                                }
                                                                                            }
                                                                                        } else {
                                                                                            if (1 > 42) {
                                                                                                var v302 uint32 = 0
                                                                                                v303 := []int{9223372036854775807}
                                                                                                if (!false) {
                                                                                                    var v304 float64 = 1.5
                                                                                                    var v305 uint32 = v302
                                                                                                    _ = v304
                                                                                                    _ = v305
                                                                                                }
                                                                                                switch "a" {
                                                                                                case "\n\t":
                                                                                                    if (float32(0.1) >= 1.5) {
                                                                                                        fmt.Println(v94)
                                                                                                    }
                                                                                                }
                                                                                                v306 := GenS2{F0: (v2 | (v2 & (v2 ^ v2)))}
                                                                                                _ = v302
                                                                                                _ = v303
                                                                                                _ = v306
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    v307 := []int{4294967295, fi14}
                                                                                    _ = v307
                                                                                }
                                                                                if (true || true) {
                                                                                    v94 -= (^(v1))
                                                                                    var v308 uint8 = v29
                                                                                    var v309 uint = v94
                                                                                    _ = v308
                                                                                    _ = v309
                                                                                } else {
                                                                                    fmt.Println(v3, fi31)
                                                                                }
                                                                            } else {
                                                                                var v310 bool = true
                                                                                v311 := []int{}
                                                                                _ = v310
                                                                                _ = v311
                                                                            }
                                                                        }
                                                                    }
                                                                    _ = v273
                                                                } else {
                                                                    fmt.Println(v2, v1, v29)
                                                                    if (69 != 0) {
                                                                        if (!((v3 >> ((255) & 63)) > -575)) {
                                                                            switch "\n\t" {
                                                                            case "foo bar":
                                                                                v312 := []int{2147483647, int(v95)}
                                                                                _ = v312
                                                                            case "":
                                                                                switch 1 {
                                                                                case 47:
                                                                                    v3 -= (v3 * (v3 >> ((uint(v2)) & 63)))
                                                                                    switch "hello" {
                                                                                    case "a":
                                                                                        var v313 uint32 = 42
                                                                                        _ = v313
                                                                                    case "foo bar":
                                                                                        fmt.Println(v29, v1)
                                                                                    case "":
                                                                                        v29 -= v29
                                                                                    }
                                                                                    v29 = 106
                                                                                case -9223372036854775808:
                                                                                    switch 9223372036854775807 {
                                                                                    case 393:
                                                                                        if (896 != 0) {
                                                                                            var v314 bool = false
                                                                                            var v315 bool = true
                                                                                            switch "a" {
                                                                                            case "foo bar":
                                                                                                fmt.Println(fi14, fi31, v95)
                                                                                            case "\n\t":
                                                                                                v316 := GenS36{F0: -1, F1: 465}
                                                                                                _ = v316
                                                                                            case "":
                                                                                                if v315 {
                                                                                                    v3 -= (v3 & (^(v3)))
                                                                                                } else {
                                                                                                    var v317 int8 = 127
                                                                                                    _ = v317
                                                                                                }
                                                                                            }
                                                                                            _ = v314
                                                                                            _ = v315
                                                                                        } else {
                                                                                            v318 := []int{}
                                                                                            var v319 uint = v1
                                                                                            var v320 int64 = -786
                                                                                            _ = v318
                                                                                            _ = v319
                                                                                            _ = v320
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    var v321 int16 = 0
                                                                                    _ = v321
                                                                                }
                                                                            case "a":
                                                                                if true {
                                                                                    var v322 float64 = float64(float32(0.0))
                                                                                    v4 += (v4 >> ((471) & 63))
                                                                                    var v323 bool = false
                                                                                    _ = v322
                                                                                    _ = v323
                                                                                } else {
                                                                                    v2 += (v2 + (v2 | (^(v2))))
                                                                                    if ((3.14159 / 0.1) >= float64(3.14159)) {
                                                                                        v324 := GenS1{F0: v95, F1: (^(v94)), F2: 6}
                                                                                        fmt.Println(fi31, v3)
                                                                                        switch 32768 {
                                                                                        case 1:
                                                                                            v325 := []int{int(v2)}
                                                                                            _ = v325
                                                                                        case -483:
                                                                                            if false {
                                                                                                switch (-(fi14)) {
                                                                                                case 2147483647:
                                                                                                    var v326 string = "foo bar"
                                                                                                    _ = v326
                                                                                                case 1:
                                                                                                    var v327 bool = true
                                                                                                    _ = v327
                                                                                                }
                                                                                                v3 = (-(v3))
                                                                                                var v328 int64 = -1
                                                                                                _ = v328
                                                                                            }
                                                                                            fmt.Println(v95, v95, fi14)
                                                                                            v1 -= (v1 & (v1 ^ (v1 | (v94 - (v94 % (((^(v1))) | 1))))))
                                                                                        case -529:
                                                                                            if false {
                                                                                                v329 := []int{(fi31 / (((fi31 | (fi14 / (((-(fi31))) | 1)))) | 1)), -9223372036854775808, (fi31 ^ (fi31 % ((genFn0()) | 1))), fi14}
                                                                                                var v330 int = fi31
                                                                                                var v331 bool = (!((true && (v2 == v2)) || false))
                                                                                                v1 -= (^(v1))
                                                                                                fmt.Println(v1, v1, v331)
                                                                                                _ = v329
                                                                                                _ = v330
                                                                                                _ = v331
                                                                                            } else {
                                                                                                var v332 string = (("" + (("a" + "foo bar") + ("foo bar" + "\n\t"))) + "hello")
                                                                                                var v333 int = (fi14 ^ fi14)
                                                                                                _ = v332
                                                                                                _ = v333
                                                                                            }
                                                                                        }
                                                                                        v334 := GenS27{F0: (v4 + (v4 << (((v1 + (v94 + v1))) & 63))), F1: int(v4)}
                                                                                        v335 := []int{(fi14 ^ 4294967295), (fi31 << (((v1 & 42)) & 63)), 32767}
                                                                                        _ = v324
                                                                                        _ = v334
                                                                                        _ = v335
                                                                                    } else {
                                                                                        v336 := GenS37{F0: -9223372036854775808}
                                                                                        _ = v336
                                                                                    }
                                                                                }
                                                                            default:
                                                                                v337 := GenS38{F0: false, F1: 2}
                                                                                _ = v337
                                                                            }
                                                                        }
                                                                        var v338 string = "\n\t"
                                                                        switch 4294967295 {
                                                                        case 871:
                                                                            fmt.Println(v95, v338, v338)
                                                                        case 4294967295:
                                                                            var v339 int8 = -53
                                                                            _ = v339
                                                                        default:
                                                                            v340 := GenS12{F0: "", F1: (fi31 << ((v1) & 63)), F2: (v29 * v29)}
                                                                            _ = v340
                                                                        }
                                                                        v341 := GenS38{F0: false, F1: 42}
                                                                        v342 := GenS33{F0: 9223372036854775807}
                                                                        _ = v338
                                                                        _ = v341
                                                                        _ = v342
                                                                    } else {
                                                                        var v343 uint = v1
                                                                        v344 := []int{127, genFn0(), fi14, genFn0()}
                                                                        v3 = (v3 << (((v1 | (v94 + 0))) & 63))
                                                                        _ = v343
                                                                        _ = v344
                                                                    }
                                                                }
                                                                var v345 bool = false
                                                                _ = v345
                                                            }
                                                            var v346 int8 = 3
                                                            _ = v346
                                                        }
                                                    }
                                                case -1:
                                                    fmt.Println(v1, v4)
                                                }
                                                _ = v93
                                                _ = v94
                                                _ = v95
                                            }
                                        case 388:
                                            var v347 int64 = -9223372036854775808
                                            _ = v347
                                        }
                                        if true {
                                            v348 := GenS39{F0: 0}
                                            var v349 uint8 = 146
                                            var v350 float32 = 0.0
                                            fmt.Println(v1, v29, v2)
                                            _ = v348
                                            _ = v349
                                            _ = v350
                                        }
                                        switch 32768 {
                                        case 587:
                                            v351 := []int{fi14, (fi14 % ((genFn0()) | 1)), 1, (fi14 - -228)}
                                            _ = v351
                                        default:
                                            v352 := []int{-14, (fi31 & genFn0()), fi14, 42}
                                            _ = v352
                                        }
                                        v1 += 65535
                                    case "":
                                        if ((!((!("" >= (("foo bar" + "a") + ("hello" + "foo bar")))) && ((v2 < v2) || (false && ((((0.1 - (0.1 - 0.0)) / 0.1) - float32((((float64(0.0) - -2.25) / 1.5) * (3.14159 / float64(3.14159))))) > 0.0))))) && true) {
                                            v353 := GenS25{F0: v29}
                                            if (((uint8(v1) >= v29) || (!(65535 >= 817))) || (("foo bar" < "hello") || false)) {
                                                v4 = (^(v4))
                                                var v354 int8 = 100
                                                _ = v354
                                            } else {
                                                var v355 int64 = 1
                                                switch 4294967295 {
                                                case 32767:
                                                    v355 -= 202
                                                case 965:
                                                    v356 := []int{fi31}
                                                    _ = v356
                                                }
                                                _ = v355
                                            }
                                            v29 = 247
                                            _ = v353
                                        } else {
                                            v357 := GenS31{F0: 1}
                                            _ = v357
                                        }
                                    case "a":
                                        if ((v3 | v3) > (v3 >> ((v1) & 63))) {
                                            if false {
                                                v358 := GenS20{F0: -1}
                                                _ = v358
                                            } else {
                                                var v359 string = ("hello" + "\n\t")
                                                _ = v359
                                            }
                                        } else {
                                            v2 = (v2 << ((uint(v4)) & 63))
                                            switch (fi31 % ((551) | 1)) {
                                            case 42:
                                                v360 := []int{int(v3)}
                                                _ = v360
                                            case -1:
                                                v361 := GenS40{F0: v4}
                                                _ = v361
                                            }
                                        }
                                    default:
                                        v362 := []int{(fi14 * (fi14 / (((fi31 * (fi31 % (((fi31 ^ (fi31 / (((fi31 % ((-301) | 1))) | 1)))) | 1)))) | 1))), (fi31 / ((fi14) | 1)), fi31, (fi14 << ((uint(v4)) & 63))}
                                        _ = v362
                                    }
                                case 32768:
                                    var v363 uint = 2
                                    _ = v363
                                default:
                                    v2 = 945
                                    var v364 int16 = -1
                                    _ = v364
                                }
                                _ = fi31
                                _ = v32
                            }
                            switch ("" + (("\n\t" + "hello") + ("a" + ("" + "\n\t")))) {
                            case "foo bar":
                                var v365 float64 = 0.1
                                _ = v365
                            case "":
                                switch genFn0() {
                                case 9223372036854775807:
                                    v366 := []int{(fi14 % (((-(fi14))) | 1)), (fi14 & (-(fi14))), fi14}
                                    _ = v366
                                case 127:
                                    var v367 int = (fi14 / (((fi14 % (((fi14 | (-(fi14)))) | 1))) | 1))
                                    _ = v367
                                default:
                                    v4 = (v4 * v4)
                                }
                            }
                            for fi368 := 0; fi368 < 5; fi368++ {
                                fmt.Println(v2)
                                v4 += v4
                                _ = fi368
                            }
                        }
                        switch (("foo bar" + "\n\t") + "foo bar") {
                        case "hello":
                            fmt.Println(v3)
                        case "foo bar":
                            var v369 uint16 = 2
                            _ = v369
                        case "\n\t":
                            v3 = (v3 - v3)
                        }
                        _ = v30
                    }
                }
                var v370 uint = v1
                _ = v29
                _ = v370
            }
            _ = fi14
        }
    }
    var v371 uint64 = (v2 ^ v2)
    switch 260 {
    case -9223372036854775808:
        switch "\n\t" {
        case "\n\t":
            v3 = (v3 ^ (v3 + 0))
        case "a":
            v372 := []int{-2147483648, -1}
            _ = v372
        }
    }
    var v373 uint32 = 1
    fmt.Println(v4, v373)
    v374 := []int{4294967295, -9223372036854775808, genFn0(), 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v371:", v371)
    fmt.Println("v373:", v373)
    fmt.Println("slicelen v374:", len(v374), v374)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v371
    _ = v373
    _ = v374
    _ = v4
}
