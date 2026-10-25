package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 int64; F2 int }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint64; F1 bool; F2 uint32 }
type GenS5 struct { F0 uint; F1 int64 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 string; F1 int }
type GenS10 struct { F0 uint; F1 int16; F2 int8 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 int32; F1 uint32 }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 bool; F1 uint; F2 int32 }
type GenS15 struct { F0 bool }
type GenS16 struct { F0 int32; F1 int8; F2 uint32 }
type GenS17 struct { F0 int16; F1 uint64 }
type GenS18 struct { F0 int16; F1 int64; F2 uint64 }
func genFn0(p0 int64) int16 {
    p0 -= -9223372036854775808
    var v5 uint8 = 128
    if false {
        v5 = (v5 - (^(v5)))
    } else {
        v6 := []int{}
        var v7 int8 = -6
        p0 = (-(p0))
        _ = v6
        _ = v7
    }
    v8 := []int{-570, int(v5), 4294967295, 32768}
    v9 := GenS1{F0: 1}
    _ = p0
    _ = v5
    _ = v8
    _ = v9
    return -1
}
func genFn1(p0 int64) int16 {
    v10 := GenS2{F0: 255, F1: (-(p0)), F2: -9223372036854775808}
    var v11 int16 = -190
    for fi12 := 0; fi12 < 6; fi12++ {
        v13 := GenS1{F0: (p0 << ((65535) & 63))}
        if true {
            for fi14 := 0; fi14 < 14; fi14++ {
                if true {
                    v15 := GenS1{F0: (p0 / (((p0 % (((p0 ^ -873)) | 1))) | 1))}
                    _ = v15
                } else {
                    p0 = 9223372036854775807
                    var v16 int32 = 2147483647
                    _ = v16
                }
                v17 := []int{(fi14 - (fi12 >> ((167) & 63))), (fi14 & (-(fi14))), fi14}
                _ = fi14
                _ = v17
            }
        }
        if false {
            p0 = (p0 * p0)
            var v18 int32 = 1
            _ = v18
        }
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = v10
    _ = v11
    return 0
}

func main() {
    var v1 int32 = 932
    var v2 int16 = -32768
    var v3 int32 = v1
    var v4 string = ("hello" + "hello")
    if (42 != 97) {
        switch v4 {
        case "foo bar":
            for fi19 := 0; fi19 < 13; fi19++ {
                v20 := []int{0, (fi19 << ((255) & 63)), (fi19 / ((fi19) | 1))}
                for fi21 := 0; fi21 < 4; fi21++ {
                    v2 += int16(v3)
                    _ = fi21
                }
                _ = fi19
                _ = v20
            }
        case "":
            v3 += 42
        case "a":
            v2 += genFn1(9223372036854775807)
        default:
            v22 := GenS1{F0: 29}
            _ = v22
        }
        var v23 int64 = -9223372036854775808
        fmt.Println(v3)
        _ = v23
    } else {
        v24 := GenS1{F0: 198}
        var v25 uint8 = 1
        var v26 int64 = 0
        _ = v24
        _ = v25
        _ = v26
    }
    if (!false) {
        switch 1 {
        case 0:
            v27 := GenS2{F0: 1, F1: -726, F2: 193}
            _ = v27
        case 128:
            var v28 int32 = (v3 * v1)
            _ = v28
        case -2147483648:
            switch 128 {
            case -2147483648:
                v29 := GenS2{F0: 0, F1: -1, F2: 0}
                _ = v29
            case -436:
                v30 := GenS2{F0: 4294967295, F1: -9223372036854775808, F2: -456}
                _ = v30
            }
        default:
            switch 1 {
            case 278:
                v4 = "foo bar"
            case 32768:
                var v31 string = "a"
                _ = v31
            }
        }
    } else {
        v2 += -1
    }
    fmt.Println(v2, v1)
    v32 := GenS3{F0: 999}
    for fi33 := 0; fi33 < 14; fi33++ {
        if true {
            if false {
                switch -9223372036854775808 {
                case 2147483647:
                    var v34 int16 = v2
                    _ = v34
                case 0:
                    for fi35 := 0; fi35 < 3; fi35++ {
                        var v36 float64 = ((0.1 - 0.0) - 0.0)
                        v37 := GenS4{F0: 2, F1: (190 < 18446744073709551615), F2: 2147483647}
                        v3 = v1
                        fmt.Println(v3, fi35, v2)
                        var v38 uint32 = 2147483647
                        var v39 float64 = v36
                        _ = fi35
                        _ = v36
                        _ = v37
                        _ = v38
                        _ = v39
                    }
                default:
                    var v40 bool = true
                    _ = v40
                }
                fmt.Println(v2)
                if (365 <= 0) {
                    var v41 uint8 = 164
                    _ = v41
                }
            }
        }
        v2 = (v2 ^ 32767)
        fmt.Println(v3, v4, v1)
        _ = fi33
    }
    for fi42 := 0; fi42 < 6; fi42++ {
        switch v4 {
        case "":
            for fi43 := 0; fi43 < 20; fi43++ {
                if false {
                    v3 = v3
                    v44 := []int{1, (^(fi43)), fi43, (fi43 | -309)}
                    _ = len(v44)
                    var v45 uint64 = 4294967295
                    _ = v44
                    _ = v45
                }
                v46 := GenS1{F0: 1}
                var v47 int16 = v2
                switch "" {
                case "":
                    var v48 int64 = 1
                    _ = v48
                case "hello":
                    switch ("hello" + v4) {
                    case "\n\t":
                        v49 := []int{}
                        _ = v49
                    case "":
                        if false {
                            if (526 >= 1) {
                                v1 -= 2147483647
                                switch "\n\t" {
                                case "a":
                                    v47 -= (^(v2))
                                case "hello":
                                    switch v4 {
                                    case "a":
                                        v50 := GenS2{F0: 42, F1: 266, F2: (fi42 * -2147483648)}
                                        v2 += (v2 & int16(v3))
                                        v1 = (v1 >> ((322) & 63))
                                        _ = v50
                                    case "":
                                        switch ("\n\t" + (v4 + v4)) {
                                        case "foo bar":
                                            if false {
                                                v1 = (v1 + v3)
                                                if true {
                                                    var v51 int64 = -1
                                                    _ = v51
                                                }
                                            }
                                        case "":
                                            v52 := []int{}
                                            _ = v52
                                        case "hello":
                                            v53 := []int{(fi42 >> ((0) & 63))}
                                            _ = v53
                                        default:
                                            v54 := GenS2{F0: 18446744073709551615, F1: 9223372036854775807, F2: (fi43 % ((-1) | 1))}
                                            _ = v54
                                        }
                                    case "foo bar":
                                        v55 := GenS5{F0: 2, F1: -713}
                                        _ = v55
                                    }
                                case "foo bar":
                                    var v56 float64 = 3.14159
                                    _ = v56
                                default:
                                    v1 = (v3 - -1)
                                }
                            } else {
                                if false {
                                    var v57 string = v4
                                    v58 := []int{}
                                    v59 := GenS6{F0: 128}
                                    _ = v57
                                    _ = v58
                                    _ = v59
                                }
                                var v60 bool = (4294967295 < 18446744073709551615)
                                var v61 int32 = (-(v1))
                                _ = v60
                                _ = v61
                            }
                            v62 := []int{}
                            if false {
                                if true {
                                    if len(v62) >= 2 {
                                        v62 = v62[1:len(v62)]
                                    }
                                    if true {
                                        var v63 uint8 = 197
                                        _ = len(v62)
                                        fmt.Println(v1, fi42, fi43)
                                        _ = v63
                                    } else {
                                        if (548 == 157) {
                                            if (!false) {
                                                var v64 float32 = (3.14159 * 0.1)
                                                var v65 bool = true
                                                switch "\n\t" {
                                                case "foo bar":
                                                    var v66 int8 = 1
                                                    _ = v66
                                                case "":
                                                    var v67 float32 = 0.0
                                                    _ = v67
                                                }
                                                _ = v64
                                                _ = v65
                                            } else {
                                                v62 = append(v62, (fi42 >> ((2) & 63)))
                                            }
                                            switch (fi42 & fi42) {
                                            case 2147483647:
                                                var v68 bool = true
                                                _ = v68
                                            case 169:
                                                if (-2.25 < ((-2.25 * -2.25) / 1.5)) {
                                                    v69 := GenS1{F0: -287}
                                                    _ = v69
                                                }
                                            case 127:
                                                if (!false) {
                                                    v70 := GenS7{F0: 0}
                                                    var v71 uint64 = 2
                                                    _ = v70
                                                    _ = v71
                                                } else {
                                                    if (((("foo bar" + v4) + (v4 + "a")) + v4) < ((v4 + v4) + "\n\t")) {
                                                        fmt.Println(v3)
                                                        switch (fi42 ^ (fi43 % (((fi43 & fi42)) | 1))) {
                                                        case 32768:
                                                            if false {
                                                                v72 := GenS3{F0: (fi43 & fi42)}
                                                                _ = v72
                                                            }
                                                        case -2147483648:
                                                            if false {
                                                                v4 = v4
                                                            }
                                                        }
                                                        var v73 string = "foo bar"
                                                        _ = v73
                                                    } else {
                                                        var v74 string = v4
                                                        _ = v74
                                                    }
                                                }
                                            }
                                            if len(v62) > 0 {
                                                v62[0] = 42
                                            }
                                            var v75 int32 = (v1 * v3)
                                            if true {
                                                var v76 float32 = -2.25
                                                var v77 float32 = v76
                                                _ = v76
                                                _ = v77
                                            } else {
                                                if ((!false) || false) {
                                                    if (1 < 2) {
                                                        if true {
                                                            switch fi42 {
                                                            case -2147483648:
                                                                var v78 uint32 = 4294967295
                                                                _ = v78
                                                            case 32767:
                                                                v4 = v4
                                                            default:
                                                                fmt.Println(v4, v47, v1)
                                                            }
                                                            v3 = v3
                                                            var v79 int = (fi42 & fi43)
                                                            v80 := GenS8{F0: 0}
                                                            _ = v79
                                                            _ = v80
                                                        }
                                                        if (false && true) {
                                                            v47 = -1
                                                        }
                                                        if len(v62) > 0 {
                                                            v62[0] = -9223372036854775808
                                                        }
                                                    }
                                                    if (0.0 > 0.0) {
                                                        v75 -= (v1 - v3)
                                                    } else {
                                                        fmt.Println(v47)
                                                    }
                                                    _ = len(v62)
                                                } else {
                                                    var v81 bool = (false || false)
                                                    fmt.Println(v75, fi43, fi43)
                                                    v82 := GenS4{F0: 2, F1: v81, F2: 2147483648}
                                                    _ = v81
                                                    _ = v82
                                                }
                                            }
                                            _ = v75
                                        }
                                        fmt.Println(v3, v3, v47)
                                    }
                                    switch ((((v4 + "") + "\n\t") + v4) + v4) {
                                    case "\n\t":
                                        switch (fi42 * fi42) {
                                        case 133:
                                            fmt.Println(fi42, v4)
                                        }
                                    case "a":
                                        if (v4 < v4) {
                                            v2 += 42
                                            v2 = (v47 & (v47 % ((v47) | 1)))
                                            var v83 uint64 = 4294967295
                                            _ = v83
                                        }
                                        v47 += (v47 >> ((18446744073709551615) & 63))
                                    default:
                                        v1 += v3
                                        v84 := GenS9{F0: v4, F1: 295}
                                        _ = v84
                                    }
                                    fmt.Println(v3, fi43, v4)
                                    switch v4 {
                                    case "hello":
                                        if false {
                                            v4 = v4
                                            v85 := GenS10{F0: 759, F1: v2, F2: -1}
                                            var v86 uint32 = 964
                                            _ = v85
                                            _ = v86
                                        } else {
                                            var v87 uint8 = 35
                                            _ = v87
                                        }
                                    case "a":
                                        v47 = (v2 % (((v47 << ((0) & 63))) | 1))
                                    case "foo bar":
                                        if (32767 != 42) {
                                            v2 = (v47 >> ((65535) & 63))
                                        } else {
                                            switch fi42 {
                                            case -2147483648:
                                                if len(v62) >= 2 {
                                                    v62 = v62[1:len(v62)]
                                                }
                                            case 128:
                                                v2 += (v2 | (v2 | v47))
                                            default:
                                                if false {
                                                    v4 = (("" + v4) + "a")
                                                    v4 = "\n\t"
                                                    var v88 int8 = 92
                                                    v2 = genFn1(1)
                                                    _ = v88
                                                } else {
                                                    switch ("hello" + v4) {
                                                    case "foo bar":
                                                        fmt.Println(v3, fi43, fi42)
                                                    case "":
                                                        var v89 int8 = 100
                                                        _ = v89
                                                    }
                                                }
                                            }
                                        }
                                    default:
                                        fmt.Println(fi42, v3)
                                    }
                                    fmt.Println(v47, v47, v1)
                                } else {
                                    if (1e10 == float32(0.0)) {
                                        fmt.Println(v1, v47, fi43)
                                        _ = len(v62)
                                        var v90 uint8 = 108
                                        _ = v90
                                    }
                                    var v91 float64 = (((float64(float32(1e10)) / 1.5) + 0.1) + (1e10 - (3.14159 / ((((0.1 / ((0.0 * (((float64(3.14159) - 1.5) + 1e10) - 0.1)) + ((((1.5 * 1.5) * -2.25) * -2.25) - (3.14159 + (3.14159 * (((1.5 - 1e10) + 1e10) * (-2.25 - (-2.25 * 0.0)))))))) + 1.5) / float64(1e10)) + (1e10 * float64(1.5))))))
                                    v92 := GenS7{F0: v1}
                                    _ = v91
                                    _ = v92
                                }
                                fmt.Println(v47)
                                if ("\n\t" < v4) {
                                    var v93 uint64 = 381
                                    if true {
                                        fmt.Println(v2, fi43)
                                    } else {
                                        if len(v62) >= 2 {
                                            v62 = v62[1:len(v62)]
                                        }
                                    }
                                    _ = v93
                                } else {
                                    v94 := GenS9{F0: v4, F1: (fi43 + (fi42 ^ (fi43 - fi42)))}
                                    _ = v94
                                }
                            }
                            fmt.Println(v2, fi42)
                            v2 -= v47
                            _ = v62
                        }
                    case "foo bar":
                        if (((0.1 + 1e10) >= -2.25) || (86 != fi43)) {
                            fmt.Println(v4)
                            if false {
                                var v95 int8 = 0
                                _ = v95
                            }
                        }
                    }
                default:
                    v96 := []int{(fi43 - (fi42 / ((-986) | 1)))}
                    _ = v96
                }
                switch (^(fi43)) {
                case -2147483648:
                    v97 := []int{(fi42 / ((fi43) | 1)), fi43, fi43, (fi43 + (fi43 / ((int(v3)) | 1)))}
                    _ = v97
                default:
                    if (1 == 0) {
                        var v98 uint = 2
                        switch (fi42 / ((fi42) | 1)) {
                        case -68:
                            v99 := []int{fi43, (fi42 | 128)}
                            _ = v99
                        case -1:
                            var v100 int8 = -1
                            _ = v100
                        default:
                            if (int64(fi43) > 42) {
                                var v101 uint64 = 1
                                _ = v101
                            } else {
                                v4 = v4
                                v102 := GenS11{F0: 1}
                                _ = v102
                            }
                        }
                        switch ((v4 + v4) + ("a" + v4)) {
                        case "foo bar":
                            v2 -= v47
                        case "\n\t":
                            var v103 float32 = -2.25
                            _ = v103
                        case "a":
                            fmt.Println(v47, v47, v3)
                        }
                        _ = v98
                    }
                }
                _ = fi43
                _ = v46
                _ = v47
            }
        case "a":
            if false {
                for fi104 := 0; fi104 < 3; fi104++ {
                    v105 := GenS12{F0: (v3 % (((^(v1))) | 1)), F1: 0}
                    v106 := []int{(^(fi42))}
                    if (1 > -9223372036854775808) {
                        switch "hello" {
                        case "a":
                            var v107 int32 = (-(v1))
                            _ = v107
                        case "hello":
                            var v108 uint64 = 636
                            _ = v108
                        case "foo bar":
                            switch (fi42 - (fi42 ^ fi42)) {
                            case 0:
                                v106 = append(v106, 42)
                            case -439:
                                var v109 bool = (!false)
                                _ = v109
                            case 127:
                                if true {
                                    v4 = (((("\n\t" + v4) + "a") + ("\n\t" + v4)) + "hello")
                                    switch v4 {
                                    case "hello":
                                        switch "a" {
                                        case "":
                                            v110 := GenS1{F0: -125}
                                            _ = v110
                                        case "a":
                                            v111 := GenS13{F0: 534}
                                            fmt.Println(v2, fi104, fi42)
                                            v3 = -2147483648
                                            _ = v111
                                        }
                                    case "\n\t":
                                        fmt.Println(v2, v3)
                                    case "foo bar":
                                        v2 = v2
                                        fmt.Println(fi42, v4)
                                        fmt.Println(v4, v3, v1)
                                        v4 = v4
                                        fmt.Println(v1, v2)
                                        fmt.Println(v4)
                                    }
                                    fmt.Println(fi42)
                                }
                            }
                        default:
                            v1 -= (v3 * v3)
                        }
                    }
                    _ = fi104
                    _ = v105
                    _ = v106
                }
            }
        }
        v4 = ("hello" + "\n\t")
        var v112 float64 = float64(3.14159)
        _ = fi42
        _ = v112
    }
    if false {
        fmt.Println(v4, v3, v3)
    } else {
        fmt.Println(v3)
        v2 -= v2
    }
    v113 := GenS10{F0: 2, F1: (v2 - v2), F2: -128}
    for fi114 := 0; fi114 < 8; fi114++ {
        fmt.Println(v2)
        switch (v4 + "hello") {
        case "foo bar":
            v115 := []int{fi114}
            _ = v115
        case "":
            v2 += -1
        case "a":
            v2 += (v2 | v2)
        default:
            var v116 float32 = (1.5 * (1.5 - 3.14159))
            _ = v116
        }
        _ = fi114
    }
    v4 = v4
    v117 := []int{9223372036854775807, -2147483648, -2147483648, -376}
    for fi118 := 0; fi118 < 11; fi118++ {
        if (470 < 227) {
            v119 := GenS14{F0: (!false), F1: 0, F2: 0}
            v4 = "\n\t"
            v4 = v4
            _ = v119
        } else {
            if len(v117) > 0 {
                v117[0] = fi118
            }
            if false {
                for fi120 := 0; fi120 < 2; fi120++ {
                    var v121 int8 = -72
                    _ = fi120
                    _ = v121
                }
                v2 = v2
                for fi122 := 0; fi122 < 19; fi122++ {
                    fmt.Println(fi118, v1, fi118)
                    _ = fi122
                }
            }
        }
        if (0.1 > 1.5) {
            if (!false) {
                switch (fi118 - fi118) {
                case 9223372036854775807:
                    var v123 int8 = 42
                    _ = v123
                case -2147483648:
                    if ((v3 >> ((18446744073709551615) & 63)) < (^(v1))) {
                        if len(v117) > 0 {
                            v117[0] = (fi118 % ((fi118) | 1))
                        }
                    }
                default:
                    var v124 string = ("a" + v4)
                    _ = v124
                }
                v2 = v2
            }
            switch ("a" + v4) {
            case "\n\t":
                switch fi118 {
                case -2147483648:
                    fmt.Println(v2, v1)
                default:
                    var v125 int8 = -1
                    _ = v125
                }
            case "hello":
                v126 := GenS1{F0: -1}
                _ = v126
            case "":
                fmt.Println(v4, v3)
            }
        } else {
            var v127 int = fi118
            var v128 int64 = 1
            switch ("\n\t" + (v4 + v4)) {
            case "foo bar":
                if ((fi118 / ((fi118) | 1)) >= fi118) {
                    for fi129 := 0; fi129 < 19; fi129++ {
                        if false {
                            v1 = 42
                            if len(v117) >= 2 {
                                v117 = v117[1:len(v117)]
                            }
                        } else {
                            if false {
                                if (!(127 > -1)) {
                                    v130 := GenS15{F0: ((v2 ^ (v2 * -1)) <= 42)}
                                    v131 := GenS16{F0: v3, F1: 84, F2: 287}
                                    fmt.Println(fi118, v3)
                                    _ = v130
                                    _ = v131
                                }
                                v3 += (v1 + (v3 % ((v1) | 1)))
                            }
                            v2 = 42
                            v117 = append(v117, (v127 - v127))
                        }
                        v132 := GenS17{F0: v2, F1: 2}
                        v1 = v3
                        var v133 string = "a"
                        v134 := GenS1{F0: (-(v128))}
                        v2 -= -32768
                        _ = fi129
                        _ = v132
                        _ = v133
                        _ = v134
                    }
                    var v135 uint = 65535
                    var v136 uint32 = 65535
                    _ = v135
                    _ = v136
                }
            case "\n\t":
                v137 := GenS18{F0: (v2 * (v2 << ((65535) & 63))), F1: 52, F2: 4294967295}
                _ = v137
            default:
                var v138 int32 = -2147483648
                _ = v138
            }
            _ = v127
            _ = v128
        }
        _ = fi118
    }
    fmt.Println(v2, v1, v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v117:", len(v117), v117)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v113
    _ = v117
    _ = v2
    _ = v3
    _ = v32
    _ = v4
}
