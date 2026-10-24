package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 uint32 }
type GenS2 struct { F0 int8; F1 uint32; F2 int32 }
type GenS3 struct { F0 string; F1 int64; F2 uint16 }
type GenS4 struct { F0 string; F1 uint64; F2 uint32 }
type GenS5 struct { F0 uint32; F1 int8 }
type GenS6 struct { F0 int; F1 uint; F2 int }
type GenS7 struct { F0 int; F1 int }
type GenS8 struct { F0 uint32; F1 bool; F2 uint16 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 string; F1 uint16; F2 int16 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 uint64 }

func main() {
    var v1 uint = 2
    var v2 int8 = -1
    v3 := make(map[int]int)
    if (-1 != -9223372036854775808) {
        var v4 string = "\n\t"
        var v5 uint8 = 42
        var v6 int32 = -1
        _ = v4
        _ = v5
        _ = v6
    }
    var v7 float64 = 0.1
    _ = len(v3)
    v8 := GenS1{F0: 528, F1: 1, F2: 2147483647}
    var v9 int64 = -9223372036854775808
    v10 := GenS1{F0: 42, F1: v9, F2: 2}
    var v11 int64 = v9
    v9 = (^(v9))
    for fi12 := 0; fi12 < 6; fi12++ {
        var v13 float64 = v7
        for fi14 := 0; fi14 < 18; fi14++ {
            if (("" + "hello") <= "foo bar") {
                if (1.5 <= -2.25) {
                    v1 = v1
                } else {
                    v15 := GenS2{F0: (-(v2)), F1: 0, F2: 1}
                    var v16 int8 = (^(v2))
                    var v17 int64 = -1
                    _ = v15
                    _ = v16
                    _ = v17
                }
                var v18 uint16 = 42
                v3[fi14] = 2147483647
                _ = v18
            }
            fmt.Println(v1)
            if true {
                if false {
                    if (!(v11 != -333)) {
                        v19 := []int{(fi12 & 42), fi12, fi12, fi12}
                        if (int8(fi12) >= int8(fi12)) {
                            v1 = v1
                        } else {
                            _ = len(v3)
                        }
                        v13 -= v13
                        _ = v19
                    }
                    switch (fi12 % ((int(v1)) | 1)) {
                    case 4294967295:
                        v1 = 858
                    }
                }
                if false {
                    switch fi12 {
                    case 4294967295:
                        switch "" {
                        case "\n\t":
                            if false {
                                if false {
                                    v20 := []int{fi12, (^(fi14))}
                                    v9 -= 1
                                    v1 = (v1 + (v1 % ((v1) | 1)))
                                    if len(v20) >= 2 {
                                        v20 = v20[1:len(v20)]
                                    }
                                    _ = v20
                                } else {
                                    if ((v1 % (((v1 & v1)) | 1)) <= (v1 >> ((42) & 63))) {
                                        fmt.Println(fi12, v2, fi12)
                                        var v21 int16 = 599
                                        v22 := GenS1{F0: 0, F1: -1, F2: 2147483648}
                                        _ = v21
                                        _ = v22
                                    } else {
                                        v7 += 3.14159
                                        if ((!(-982 != 1)) && true) {
                                            if ((42 > 123) || true) {
                                                v23 := GenS1{F0: 586, F1: (v11 % (((v9 * (v11 ^ v11))) | 1)), F2: 2147483648}
                                                _ = v23
                                            }
                                            switch (fi14 % (((fi12 & (fi12 % ((fi12) | 1)))) | 1)) {
                                            case 9223372036854775807:
                                                v24 := GenS1{F0: 1, F1: 0, F2: 739}
                                                _ = v24
                                            case 0:
                                                var v25 int8 = v2
                                                _ = v25
                                            }
                                            switch "\n\t" {
                                            case "a":
                                                switch "a" {
                                                case "":
                                                    var v26 float64 = 0.1
                                                    _ = v26
                                                case "hello":
                                                    v27 := GenS2{F0: (v2 & v2), F1: 0, F2: -2147483648}
                                                    _ = v27
                                                case "foo bar":
                                                    switch "a" {
                                                    case "\n\t":
                                                        v2 = (v2 ^ 0)
                                                    default:
                                                        if (82 >= 32767) {
                                                            v2 = v2
                                                            v11 += (v11 >> (((v1 % (((^(v1))) | 1))) & 63))
                                                            if false {
                                                                var v28 int16 = 1
                                                                _ = v28
                                                            } else {
                                                                if ("" < "a") {
                                                                    switch "\n\t" {
                                                                    case "a":
                                                                        v13 = (v13 + (v7 * v7))
                                                                    default:
                                                                        switch "a" {
                                                                        case "":
                                                                            var v29 int64 = (v11 | 0)
                                                                            _ = v29
                                                                        default:
                                                                            v30 := GenS2{F0: (v2 % (((v2 - (v2 << ((v1) & 63)))) | 1)), F1: 2, F2: 0}
                                                                            _ = v30
                                                                        }
                                                                        if (42 >= 0) {
                                                                            v31 := []int{2147483647, 0}
                                                                            if (250 > 31) {
                                                                                var v32 float32 = (0.1 / (float32(0.1) / ((((1.5 - (0.0 / (3.14159 + (-2.25 + -2.25)))) / (3.14159 + 3.14159)) - 1e10) / 0.1)))
                                                                                _ = v32
                                                                            }
                                                                            delete(v3, fi12)
                                                                            _ = v31
                                                                        } else {
                                                                            v7 += v7
                                                                        }
                                                                        switch fi14 {
                                                                        case 0:
                                                                            fmt.Println(v7, v11)
                                                                        }
                                                                    }
                                                                    if (2 >= 741) {
                                                                        var v33 uint64 = 4294967295
                                                                        var v34 int8 = v2
                                                                        var v35 string = ("a" + "a")
                                                                        _ = v33
                                                                        _ = v34
                                                                        _ = v35
                                                                    } else {
                                                                        v7 = float64(1.5)
                                                                    }
                                                                } else {
                                                                    if false {
                                                                        v2 = -128
                                                                        var v36 int16 = 0
                                                                        _ = v36
                                                                    }
                                                                }
                                                                var v37 uint64 = 207
                                                                v38 := GenS1{F0: (v37 / (((v37 - 0)) | 1)), F1: v9, F2: uint32(v37)}
                                                                _ = v37
                                                                _ = v38
                                                            }
                                                            v39 := []int{fi12}
                                                            _ = v39
                                                        }
                                                    }
                                                }
                                            case "foo bar":
                                                switch "foo bar" {
                                                case "foo bar":
                                                    v40 := []int{}
                                                    v41 := v3[int(v2)]
                                                    _ = v40
                                                    _ = v41
                                                case "\n\t":
                                                    if (v2 > v2) {
                                                        var v42 uint = (v1 & (v1 ^ 255))
                                                        if false {
                                                            if false {
                                                                switch fi14 {
                                                                case 629:
                                                                    v11 += (v11 % (((v9 * v11)) | 1))
                                                                }
                                                            }
                                                        } else {
                                                            if (v2 <= 100) {
                                                                v43 := v3[(fi14 ^ -1)]
                                                                _ = v43
                                                            } else {
                                                                switch (fi14 & -9223372036854775808) {
                                                                case 127:
                                                                    fmt.Println(v2, v11)
                                                                case 149:
                                                                    var v44 int64 = (^(v9))
                                                                    v3[1] = fi12
                                                                    var v45 int64 = v9
                                                                    _ = v44
                                                                    _ = v45
                                                                case -9223372036854775808:
                                                                    fmt.Println(fi12, fi14)
                                                                }
                                                                v42 -= (^(v1))
                                                            }
                                                            v46 := GenS2{F0: -61, F1: 97, F2: 2147483647}
                                                            var v47 uint8 = 2
                                                            v48 := []int{fi12}
                                                            _ = v46
                                                            _ = v47
                                                            _ = v48
                                                        }
                                                        var v49 uint64 = 1
                                                        if true {
                                                            switch (("" + "\n\t") + "hello") {
                                                            case "":
                                                                var v50 uint8 = 0
                                                                _ = v50
                                                            default:
                                                                v11 -= int64(v1)
                                                            }
                                                            fmt.Println(v11, v42, v42)
                                                        }
                                                        _ = v42
                                                        _ = v49
                                                    }
                                                }
                                            }
                                        }
                                        v51 := GenS3{F0: "a", F1: v11, F2: uint16(v1)}
                                        v52 := GenS3{F0: (("hello" + "hello") + (("foo bar" + "") + "hello")), F1: 1, F2: 688}
                                        _ = v51
                                        _ = v52
                                    }
                                }
                                v7 -= float64(float32(((float64(1.5) * v13) - ((float64(3.14159) - ((float64(-2.25) - 0.1) + v7)) - v13))))
                                v53 := GenS3{F0: "foo bar", F1: (-(v11)), F2: 65535}
                                v54 := []int{}
                                var v55 int64 = v11
                                fmt.Println(v11, v55, v2)
                                fmt.Println(v9)
                                _ = v53
                                _ = v54
                                _ = v55
                            } else {
                                v9 += v11
                                var v56 int16 = 8
                                var v57 float32 = 3.14159
                                _ = v56
                                _ = v57
                            }
                        case "foo bar":
                            if true {
                                v58 := []int{fi12, int(v9), (fi14 + 127)}
                                v9 = (v11 << ((v1) & 63))
                                _ = v58
                            }
                        case "a":
                            var v59 string = "hello"
                            _ = v59
                        }
                        _ = len(v3)
                        if true {
                            v11 -= v11
                            var v60 float32 = -2.25
                            v60 -= v60
                            _ = v60
                        }
                    case 27:
                        v13 -= v7
                    case 32767:
                        if (!(!(0 >= v2))) {
                            if (1 < -1) {
                                var v61 float64 = v13
                                _ = v61
                            }
                            v62 := GenS2{F0: (^(v2)), F1: 2147483647, F2: 0}
                            v63 := GenS2{F0: v2, F1: 2, F2: 42}
                            _ = v62
                            _ = v63
                        } else {
                            v64 := GenS4{F0: "a", F1: 18446744073709551615, F2: 65535}
                            v65 := GenS1{F0: 542, F1: 42, F2: 2147483648}
                            var v66 int16 = 1
                            _ = v64
                            _ = v65
                            _ = v66
                        }
                    }
                    var v67 string = "hello"
                    var v68 bool = true
                    v69 := GenS5{F0: 126, F1: (v2 + 127)}
                    v70 := GenS3{F0: (v67 + "a"), F1: (^(v9)), F2: 2}
                    _ = v67
                    _ = v68
                    _ = v69
                    _ = v70
                }
                if false {
                    if ((^(v1)) >= (v1 | (^(v1)))) {
                        var v71 uint64 = 42
                        var v72 bool = (!(3.14159 == 3.14159))
                        _ = v71
                        _ = v72
                    } else {
                        var v73 uint = (v1 / ((42) | 1))
                        v9 -= (^(v9))
                        _ = v73
                    }
                    if ((v2 > 1) || true) {
                        v74 := GenS1{F0: 1, F1: v11, F2: 4294967295}
                        var v75 int8 = v2
                        _ = v74
                        _ = v75
                    }
                } else {
                    if ((!false) && false) {
                        v76 := GenS6{F0: (fi14 | 0), F1: v1, F2: (fi14 >> ((v1) & 63))}
                        var v77 int64 = (^(v9))
                        var v78 int16 = -441
                        _ = v76
                        _ = v77
                        _ = v78
                    }
                    v79 := v3[(fi12 % (((fi12 + fi14)) | 1))]
                    _ = v79
                }
                fmt.Println(v7, v11)
                var v80 bool = true
                _ = v80
            } else {
                switch "a" {
                case "hello":
                    if (1 <= 490) {
                        var v81 uint16 = 809
                        _ = v81
                    } else {
                        switch "\n\t" {
                        case "":
                            switch fi14 {
                            case 1:
                                if (false && false) {
                                    var v82 string = "hello"
                                    _ = v82
                                }
                            case -404:
                                v9 -= 9223372036854775807
                            case 0:
                                v7 -= (((-2.25 + ((-2.25 * (float64((-2.25 - ((0.0 * ((3.14159 + 1.5) + (3.14159 * -2.25))) + -2.25))) * v13)) + (((3.14159 / 1.5) / v7) - v7))) - v13) - 0.1)
                            }
                        case "a":
                            switch "\n\t" {
                            case "a":
                                fmt.Println(v13, v2)
                            case "\n\t":
                                _ = len(v3)
                            case "foo bar":
                                if false {
                                    v2 = v2
                                    v83 := GenS5{F0: 2147483648, F1: (v2 % ((v2) | 1))}
                                    _ = v83
                                } else {
                                    v84 := []int{-700, (-(fi12)), fi12, -1}
                                    v3[fi12] = -9223372036854775808
                                    if true {
                                        if true {
                                            v9 -= (v11 & 9223372036854775807)
                                        } else {
                                            if len(v84) > 0 {
                                                v84[0] = fi14
                                            }
                                            if len(v84) >= 2 {
                                                v84 = v84[1:len(v84)]
                                            }
                                            if false {
                                                var v85 float32 = 3.14159
                                                v85 -= float32(v7)
                                                if false {
                                                    if (912 < 32767) {
                                                        var v86 int = (-(fi12))
                                                        v85 = 1.5
                                                        if (fi12 > fi12) {
                                                            if (0 == 1) {
                                                                v9 = (v11 << ((v1) & 63))
                                                                var v87 uint32 = 2147483648
                                                                _ = v87
                                                            } else {
                                                                var v88 uint64 = 4294967295
                                                                _ = v88
                                                            }
                                                        } else {
                                                            v89 := GenS3{F0: "a", F1: -515, F2: 65535}
                                                            v90 := GenS7{F0: (v86 ^ 815), F1: fi14}
                                                            if len(v84) > 0 {
                                                                v84[0] = (fi14 / (((v86 >> ((v1) & 63))) | 1))
                                                            }
                                                            _ = v89
                                                            _ = v90
                                                        }
                                                        _ = v86
                                                    } else {
                                                        v2 += 100
                                                        var v91 bool = true
                                                        switch fi12 {
                                                        case -2147483648:
                                                            switch fi12 {
                                                            case 9223372036854775807:
                                                                v3[-647] = fi12
                                                            case 625:
                                                                v84 = append(v84, 32767)
                                                            default:
                                                                v91 = (((v9 * v9) < v9) && false)
                                                            }
                                                        }
                                                        _ = v91
                                                    }
                                                    switch ("a" + "foo bar") {
                                                    case "hello":
                                                        if (false || true) {
                                                            _ = len(v84)
                                                        }
                                                    case "":
                                                        v92 := v3[fi12]
                                                        _ = v92
                                                    case "\n\t":
                                                        if true {
                                                            var v93 int32 = 2147483647
                                                            _ = v93
                                                        } else {
                                                            _ = len(v3)
                                                        }
                                                    }
                                                } else {
                                                    if (((("foo bar" + "") + "") + "hello") < "foo bar") {
                                                        v2 += (v2 * (v2 + 0))
                                                        v94 := GenS8{F0: 2, F1: (v1 < v1), F2: 32767}
                                                        _ = len(v3)
                                                        fmt.Println(v2, v13)
                                                        _ = v94
                                                    }
                                                }
                                                _ = v85
                                            } else {
                                                var v95 uint64 = 769
                                                _ = v95
                                            }
                                        }
                                        switch int(v11) {
                                        case 127:
                                            v9 += (v9 * v9)
                                        case 1:
                                            var v96 float64 = v13
                                            _ = v96
                                        default:
                                            if (int8(fi12) >= 100) {
                                                if true {
                                                    if true {
                                                        fmt.Println(v11, v7)
                                                        v84 = append(v84, fi14)
                                                        _ = len(v84)
                                                    }
                                                    delete(v3, (fi14 * fi14))
                                                    v13 = (v13 + -2.25)
                                                    switch "hello" {
                                                    case "foo bar":
                                                        switch fi14 {
                                                        case 9223372036854775807:
                                                            switch (fi12 ^ -660) {
                                                            case 128:
                                                                var v97 int8 = v2
                                                                _ = v97
                                                            default:
                                                                fmt.Println(v1)
                                                            }
                                                        default:
                                                            switch fi14 {
                                                            case 32768:
                                                                var v98 uint16 = 424
                                                                _ = v98
                                                            case 456:
                                                                if false {
                                                                    v84 = append(v84, (fi14 % ((fi14) | 1)))
                                                                    v1 += v1
                                                                    var v99 int16 = int16(v11)
                                                                    _ = v99
                                                                } else {
                                                                    v84 = append(v84, (^(fi12)))
                                                                    fmt.Println(v2, v11, fi12)
                                                                    v3[(fi12 << ((v1) & 63))] = (fi12 ^ 2147483647)
                                                                }
                                                            default:
                                                                switch fi14 {
                                                                case -2147483648:
                                                                    v11 += 9223372036854775807
                                                                default:
                                                                    var v100 uint32 = 318
                                                                    _ = v100
                                                                }
                                                            }
                                                        }
                                                    case "":
                                                        switch (fi12 - (fi12 % ((fi12) | 1))) {
                                                        case 9223372036854775807:
                                                            v7 = (1e10 + v13)
                                                        case 127:
                                                            var v101 bool = false
                                                            _ = v101
                                                        default:
                                                            switch 4294967295 {
                                                            case 9223372036854775807:
                                                                if (0 == v2) {
                                                                    fmt.Println(v11, v13)
                                                                }
                                                            }
                                                        }
                                                    case "hello":
                                                        var v102 uint64 = 637
                                                        _ = v102
                                                    }
                                                    fmt.Println(v2)
                                                    var v103 uint = v1
                                                    v1 -= v1
                                                    v13 += v7
                                                    _ = v103
                                                } else {
                                                    v104 := GenS3{F0: (("a" + "a") + "hello"), F1: -9223372036854775808, F2: 2}
                                                    v13 += (0.0 * -2.25)
                                                    if (0 != 1000) {
                                                        v13 -= v13
                                                        var v105 bool = true
                                                        fmt.Println(v13, v11, v13)
                                                        _ = v105
                                                    } else {
                                                        v2 += (v2 - v2)
                                                        v7 -= (0.1 / (0.1 - 0.0))
                                                    }
                                                    v11 = 0
                                                    _ = v104
                                                }
                                            } else {
                                                if len(v84) >= 2 {
                                                    v84 = v84[1:len(v84)]
                                                }
                                                if true {
                                                    v106 := GenS9{F0: 18446744073709551615}
                                                    switch ("a" + "hello") {
                                                    case "foo bar":
                                                        var v107 int64 = 1
                                                        _ = v107
                                                    }
                                                    _ = v106
                                                } else {
                                                    if len(v84) >= 2 {
                                                        v84 = v84[1:len(v84)]
                                                    }
                                                }
                                                fmt.Println(fi12, v9)
                                                v7 += (v13 / (1e10 - v13))
                                            }
                                        }
                                    }
                                    _ = v84
                                }
                            default:
                                var v108 int32 = 1
                                _ = v108
                            }
                        case "hello":
                            var v109 int64 = 0
                            _ = v109
                        default:
                            fmt.Println(v11, v1)
                        }
                        if true {
                            if (v2 == int8(fi14)) {
                                if false {
                                    v110 := []int{(fi14 ^ 870), 42, 1, (fi14 + fi12)}
                                    var v111 uint32 = 2
                                    _ = v110
                                    _ = v111
                                } else {
                                    v112 := []int{fi12}
                                    switch "hello" {
                                    case "\n\t":
                                        var v113 bool = true
                                        if v113 {
                                            var v114 string = "a"
                                            _ = v114
                                        }
                                        if (128 == (fi12 | 128)) {
                                            v115 := GenS9{F0: 1}
                                            _ = v115
                                        } else {
                                            var v116 bool = (!(1e10 != ((0.1 * ((float32(0.1) * (3.14159 / (-2.25 + ((1e10 - float32(v13)) + 3.14159)))) + (3.14159 - -2.25))) + 1e10)))
                                            _ = v116
                                        }
                                        _ = v113
                                    }
                                    if (127 >= 107) {
                                        var v117 uint32 = 65535
                                        var v118 uint8 = 128
                                        var v119 string = "hello"
                                        _ = v117
                                        _ = v118
                                        _ = v119
                                    }
                                    _ = len(v112)
                                    _ = v112
                                }
                                var v120 int32 = 2147483647
                                v121 := GenS3{F0: ("hello" + (("\n\t" + (("\n\t" + "a") + ("hello" + "\n\t"))) + ("\n\t" + "hello"))), F1: v11, F2: 0}
                                _ = v120
                                _ = v121
                            } else {
                                var v122 int16 = int16(v1)
                                switch (-(fi14)) {
                                case 32768:
                                    var v123 int8 = 60
                                    _ = v123
                                case 128:
                                    switch 42 {
                                    case 32767:
                                        v124 := GenS10{F0: "foo bar", F1: 1, F2: (v122 - v122)}
                                        _ = v124
                                    case -230:
                                        var v125 int64 = (v11 ^ 0)
                                        _ = v125
                                    case 0:
                                        if (!false) {
                                            v122 = 0
                                            var v126 int64 = (v9 % (((v11 - v11)) | 1))
                                            var v127 int32 = 42
                                            _ = v126
                                            _ = v127
                                        } else {
                                            switch fi12 {
                                            case -1:
                                                switch "foo bar" {
                                                case "a":
                                                    v1 = v1
                                                case "hello":
                                                    v128 := GenS11{F0: 0}
                                                    _ = v128
                                                default:
                                                    v7 = 1.5
                                                    switch "" {
                                                    case "a":
                                                        fmt.Println(v11, v7, v1)
                                                    case "hello":
                                                        var v129 int16 = v122
                                                        _ = v129
                                                    case "":
                                                        switch (fi14 << (((^(v1))) & 63)) {
                                                        case 0:
                                                            var v130 float64 = (((v13 + v13) + float64(1.5)) * v13)
                                                            _ = v130
                                                        }
                                                    }
                                                }
                                            }
                                        }
                                    }
                                    switch int(v122) {
                                    case 32768:
                                        if true {
                                            v2 = v2
                                            fmt.Println(fi14)
                                            var v131 bool = false
                                            _ = v131
                                        }
                                    case 1:
                                        v1 -= (v1 & v1)
                                    default:
                                        v2 = int8(v11)
                                    }
                                case 1:
                                    if true {
                                        var v132 string = "\n\t"
                                        var v133 int = (fi14 ^ fi12)
                                        _ = v132
                                        _ = v133
                                    } else {
                                        var v134 int8 = v2
                                        if (v1 < v1) {
                                            delete(v3, fi12)
                                        }
                                        _ = v134
                                    }
                                default:
                                    v135 := GenS4{F0: ("hello" + "a"), F1: 2, F2: 2147483647}
                                    _ = v135
                                }
                                v136 := []int{(fi12 / (((fi12 << ((v1) & 63))) | 1)), 127, (-(fi12))}
                                _ = v122
                                _ = v136
                            }
                        }
                        v137 := []int{(fi12 ^ fi12), fi12, fi14}
                        _ = v137
                    }
                case "\n\t":
                    var v138 int64 = (v11 | 9223372036854775807)
                    _ = v138
                case "a":
                    v139 := GenS12{F0: 42}
                    _ = v139
                default:
                    v140 := GenS13{F0: 783}
                    _ = v140
                }
                if false {
                    delete(v3, 310)
                }
                fmt.Println(v2, v7, v1)
            }
            _ = fi14
        }
        v11 -= v11
        _ = fi12
        _ = v13
    }
    for fi141 := 0; fi141 < 8; fi141++ {
        v11 += (v11 / ((v9) | 1))
        v1 = v1
        _ = fi141
    }
    v3[127] = 4294967295
    _ = len(v3)
    if (!false) {
        v9 = v9
        var v142 int16 = 1
        _ = v142
    } else {
        v1 += (v1 / ((v1) | 1))
    }
    v3[127] = 9223372036854775807
    var v143 int = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v143:", v143)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v143
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
