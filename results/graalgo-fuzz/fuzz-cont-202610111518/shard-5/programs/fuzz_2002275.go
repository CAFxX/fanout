package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 uint64 }
type GenS2 struct { F0 int8; F1 int64; F2 int }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int64; F1 uint; F2 uint32 }
type GenS5 struct { F0 uint16; F1 uint16; F2 uint64 }
type GenS6 struct { F0 int; F1 uint64 }
type GenS7 struct { F0 uint8; F1 int8 }
type GenS8 struct { F0 int64; F1 uint64 }
type GenS9 struct { F0 int8; F1 uint16 }
type GenS10 struct { F0 uint32; F1 int8; F2 string }
type GenS11 struct { F0 uint16; F1 uint8 }
type GenS12 struct { F0 bool; F1 int }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 uint32 }
type GenS15 struct { F0 int64; F1 uint8 }
type GenS16 struct { F0 int64 }
type GenS17 struct { F0 uint32 }
type GenS18 struct { F0 int8 }
type GenS19 struct { F0 int32 }
type GenS20 struct { F0 int16; F1 bool; F2 int8 }
type GenS21 struct { F0 int8; F1 uint16; F2 uint64 }
type GenS22 struct { F0 int64; F1 int16; F2 bool }
func genFn0(p0 int32, p1 uint, p2 string) int32 {
    p2 = (p2 + ("\n\t" + ""))
    switch 2147483647 {
    case 32768:
        for fi6 := 0; fi6 < 20; fi6++ {
            if (((false && false) && true) || (p1 == 65535)) {
                for fi7 := 0; fi7 < 15; fi7++ {
                    var v8 int16 = 42
                    _ = fi7
                    _ = v8
                }
                var v9 int8 = 127
                var v10 bool = true
                _ = v10
                _ = v9
            }
            for fi11 := 0; fi11 < 17; fi11++ {
                v12 := GenS1{F0: (fi11 >> ((1) & 63)), F1: 0, F2: 42}
                fmt.Println(fi6)
                _ = fi11
                _ = v12
            }
            _ = fi6
        }
    case -2147483648:
        fmt.Println(p2)
    default:
        v13 := []int{-913, 0}
        _ = v13
    }
    v14 := []int{4294967295}
    var v15 uint = p1
    switch 588 {
    case -105:
        switch 261 {
        case -1:
            var v16 uint64 = 866
            _ = v16
        }
    case 42:
        for fi17 := 0; fi17 < 3; fi17++ {
            if ((1.5 + 3.14159) != float32(1.5)) {
                v18 := GenS1{F0: fi17, F1: 42, F2: 1}
                var v19 uint = p1
                _ = v18
                _ = v19
            }
            _ = fi17
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    return p0
}
func genFn1(p0 int64, p1 string) uint32 {
    v20 := []int{-672, -675, 1}
    if len(v20) > 0 {
        v20[0] = 1
    }
    p0 += p0
    for fi21 := 0; fi21 < 4; fi21++ {
        switch (^(fi21)) {
        case 4294967295:
            var v22 uint = 235
            _ = v22
        }
        if false {
            for fi23 := 0; fi23 < 20; fi23++ {
                var v24 int64 = p0
                _ = fi23
                _ = v24
            }
        }
        _ = fi21
    }
    _ = p0
    _ = p1
    _ = v20
    return 4294967295
}

func main() {
    var v1 string = "a"
    v2 := []int{42, 583}
    v3 := []string{v1, ((v1 + "hello") + v1), "a"}
    var v4 float64 = (1.5 + 0.0)
    var v5 float32 = (1e10 + 3.14159)
    if false {
        switch 32768 {
        case 84:
            if true {
                v25 := GenS1{F0: -9223372036854775808, F1: 62, F2: 42}
                if (false && (-2147483648 >= -996)) {
                    switch v1 {
                    case "hello":
                        fmt.Println(v5)
                    case "":
                        var v26 uint16 = 1
                        _ = v26
                    default:
                        fmt.Println(v4, v5)
                    }
                    if len(v2) > 0 {
                        v2[0] = 128
                    }
                    var v27 uint64 = 1
                    _ = v27
                } else {
                    v5 += v5
                    var v28 bool = (!(-128 >= 42))
                    v28 = (42 < -521)
                    _ = v28
                }
                _ = v25
            } else {
                switch 670 {
                case -1:
                    var v29 float64 = -2.25
                    _ = v29
                }
            }
        case -2147483648:
            fmt.Println(v1, v1)
        default:
            switch ("hello" + ((v1 + "hello") + "a")) {
            case "a":
                switch ("a" + (v1 + "foo bar")) {
                case "foo bar":
                    if (-9223372036854775808 < 0) {
                        v4 -= ((3.14159 * v4) / ((((v4 - ((v4 - (v4 + float64(float32((v4 / (v4 * float64(0.1))))))) * (v4 * ((1e10 * 3.14159) * 1.5)))) - 1.5) / v4) * 0.0))
                    }
                case "":
                    v5 = float32((v4 / 3.14159))
                case "hello":
                    for fi30 := 0; fi30 < 17; fi30++ {
                        v4 += v4
                        v5 = v5
                        var v31 int16 = 0
                        _ = fi30
                        _ = v31
                    }
                }
            case "hello":
                v4 += v4
            case "foo bar":
                v3 = append(v3, "a")
            default:
                v2 = append(v2, 860)
            }
        }
    } else {
        v4 = (v4 / v4)
    }
    if len(v3) > 0 {
        v3[0] = (v1 + "foo bar")
    }
    if (float64((0.0 + (-2.25 - float32(((v4 - 0.1) + ((float64(v5) + 0.1) * v4)))))) > float64(v5)) {
        var v32 int32 = 0
        switch v1 {
        case "a":
            if (-128 < 0) {
                v5 = (0.0 / v5)
                switch -685 {
                case 42:
                    var v33 uint = 4294967295
                    _ = v33
                case 128:
                    if false {
                        v34 := GenS1{F0: -9223372036854775808, F1: 42, F2: 2}
                        v32 -= genFn0(0, 921, v1)
                        v1 = v1
                        _ = v34
                    } else {
                        var v35 bool = true
                        _ = v35
                    }
                }
            }
        default:
            fmt.Println(v32, v5)
        }
        _ = v32
    } else {
        v5 = (float32((0.0 / 0.1)) / float32(0.1))
        for fi36 := 0; fi36 < 1; fi36++ {
            var v37 uint = 0
            if (42 == 1) {
                var v38 uint64 = 2
                for fi39 := 0; fi39 < 1; fi39++ {
                    v4 += 0.1
                    _ = fi39
                }
                for fi40 := 0; fi40 < 1; fi40++ {
                    var v41 int64 = 9223372036854775807
                    var v42 int8 = int8(v38)
                    _ = fi40
                    _ = v41
                    _ = v42
                }
                _ = v38
            } else {
                if false {
                    var v43 uint8 = 128
                    _ = v43
                } else {
                    fmt.Println(v37, v5)
                }
            }
            _ = fi36
            _ = v37
        }
    }
    fmt.Println(v4)
    switch "hello" {
    case "foo bar":
        if (-95 <= -1) {
            v5 = v5
            var v44 uint64 = 910
            _ = v44
        } else {
            if true {
                v1 = ((v1 + "a") + (v1 + v1))
                switch ("foo bar" + v1) {
                case "foo bar":
                    for fi45 := 0; fi45 < 18; fi45++ {
                        v46 := GenS1{F0: fi45, F1: 1, F2: 4294967295}
                        for fi47 := 0; fi47 < 7; fi47++ {
                            switch v1 {
                            case "foo bar":
                                if (2147483647 <= 628) {
                                    if (0 != 438) {
                                        v48 := GenS1{F0: (fi45 - fi47), F1: 65535, F2: 1}
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        var v49 int8 = -1
                                        _ = v48
                                        _ = v49
                                    } else {
                                        if (((-301 >= -32768) && true) || (65535 <= 2)) {
                                            var v50 string = ((v1 + v1) + v1)
                                            fmt.Println(v5)
                                            _ = v50
                                        } else {
                                            fmt.Println(fi47, fi45)
                                            switch (fi45 ^ fi45) {
                                            case -454:
                                                v1 = v1
                                            default:
                                                var v51 uint64 = 4294967295
                                                _ = v51
                                            }
                                        }
                                    }
                                    var v52 int8 = -128
                                    _ = v52
                                } else {
                                    switch 9223372036854775807 {
                                    case -9223372036854775808:
                                        if len(v3) > 0 {
                                            v3[0] = "\n\t"
                                        }
                                    default:
                                        if len(v3) > 0 {
                                            v3[0] = ""
                                        }
                                        v3 = append(v3, "\n\t")
                                        var v53 int = fi45
                                        if true {
                                            switch v1 {
                                            case "a":
                                                v4 -= v4
                                            case "\n\t":
                                                v53 = (v53 >> ((18446744073709551615) & 63))
                                            }
                                        }
                                        _ = v53
                                    }
                                    v54 := GenS2{F0: 100, F1: -854, F2: 32767}
                                    _ = v54
                                }
                            case "hello":
                                v1 = v1
                            case "a":
                                v55 := GenS1{F0: (^(fi47)), F1: 1, F2: 2}
                                _ = v55
                            default:
                                var v56 uint8 = 2
                                _ = v56
                            }
                            _ = fi47
                        }
                        _ = fi45
                        _ = v46
                    }
                case "":
                    for fi57 := 0; fi57 < 3; fi57++ {
                        switch 2147483647 {
                        case -9223372036854775808:
                            v1 = ((v1 + v1) + v1)
                        default:
                            for fi58 := 0; fi58 < 12; fi58++ {
                                v5 = ((-2.25 * v5) / v5)
                                var v59 uint32 = 2147483648
                                _ = fi58
                                _ = v59
                            }
                        }
                        v60 := GenS3{F0: (2147483647 >= genFn0(57, 255, "\n\t"))}
                        var v61 uint16 = 473
                        _ = fi57
                        _ = v60
                        _ = v61
                    }
                default:
                    var v62 int64 = 376
                    _ = v62
                }
                if (856 < 32767) {
                    v5 = 0.0
                    if ((32767 != 15) || (false || (-41 > 951))) {
                        var v63 uint8 = 255
                        v64 := GenS3{F0: (940 != 1)}
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = v63
                        _ = v64
                    }
                    v5 = 0.1
                } else {
                    if ((!(1 > 18446744073709551615)) && (!(2 <= 18446744073709551615))) {
                        for fi65 := 0; fi65 < 20; fi65++ {
                            var v66 uint64 = 9
                            _ = fi65
                            _ = v66
                        }
                        if (false && false) {
                            v4 += v4
                        }
                    } else {
                        for fi67 := 0; fi67 < 16; fi67++ {
                            if (!(genFn1(223, v1) == 855)) {
                                v4 -= v4
                                if true {
                                    if (1.5 != 1e10) {
                                        v68 := GenS3{F0: false}
                                        v5 = v5
                                        var v69 float64 = ((v4 - v4) - v4)
                                        _ = v68
                                        _ = v69
                                    }
                                    switch v1 {
                                    case "hello":
                                        if (!(252 <= 124)) {
                                            var v70 int32 = 808
                                            if (!false) {
                                                v71 := GenS4{F0: 42, F1: 0, F2: 343}
                                                for fi72 := 0; fi72 < 9; fi72++ {
                                                    var v73 uint32 = genFn1(527, ("\n\t" + ("" + ("foo bar" + ("" + ((v1 + ((v1 + "a") + v1)) + ("\n\t" + v1)))))))
                                                    if false {
                                                        fmt.Println(v4, v73, fi67)
                                                    }
                                                    var v74 int = fi72
                                                    _ = fi72
                                                    _ = v73
                                                    _ = v74
                                                }
                                                _ = v71
                                            }
                                            for fi75 := 0; fi75 < 14; fi75++ {
                                                var v76 uint16 = 1
                                                v1 = (v1 + v1)
                                                _ = fi75
                                                _ = v76
                                            }
                                            _ = v70
                                        } else {
                                            fmt.Println(v4)
                                            v77 := GenS3{F0: true}
                                            if ((65535 <= 201) || true) {
                                                var v78 uint8 = 255
                                                _ = v78
                                            } else {
                                                switch (fi67 * fi67) {
                                                case 1:
                                                    if true {
                                                        for fi79 := 0; fi79 < 12; fi79++ {
                                                            if true {
                                                                if false {
                                                                    v80 := GenS2{F0: -46, F1: -9223372036854775808, F2: (fi79 | 32768)}
                                                                    switch "foo bar" {
                                                                    case "":
                                                                        var v81 int = -9223372036854775808
                                                                        _ = v81
                                                                    }
                                                                    _ = v80
                                                                } else {
                                                                    var v82 uint16 = 32767
                                                                    var v83 uint8 = 2
                                                                    _ = v82
                                                                    _ = v83
                                                                }
                                                                v84 := GenS5{F0: 0, F1: 42, F2: 530}
                                                                _ = v84
                                                            } else {
                                                                var v85 uint8 = 1
                                                                fmt.Println(fi67, v1)
                                                                if (-98 <= -1) {
                                                                    v86 := GenS6{F0: fi79, F1: 2}
                                                                    _ = v86
                                                                }
                                                                _ = v85
                                                            }
                                                            _ = fi79
                                                        }
                                                        var v87 uint16 = 0
                                                        _ = v87
                                                    }
                                                case -1:
                                                    if len(v2) > 0 {
                                                        v2[0] = (fi67 / ((fi67) | 1))
                                                    }
                                                case 845:
                                                    v5 += v5
                                                default:
                                                    var v88 int32 = 42
                                                    _ = v88
                                                }
                                                if false {
                                                    v89 := GenS7{F0: 255, F1: -69}
                                                    _ = v89
                                                } else {
                                                    fmt.Println(v1, v5, v1)
                                                    switch "" {
                                                    case "a":
                                                        for fi90 := 0; fi90 < 17; fi90++ {
                                                            _ = len(v3)
                                                            var v91 int16 = -32768
                                                            var v92 uint16 = 1
                                                            _ = fi90
                                                            _ = v91
                                                            _ = v92
                                                        }
                                                    case "":
                                                        var v93 float64 = (1.5 * ((1e10 - v4) * ((((v4 - v4) - v4) * v4) * v4)))
                                                        _ = v93
                                                    case "foo bar":
                                                        v94 := GenS8{F0: -9223372036854775808, F1: 0}
                                                        _ = v94
                                                    }
                                                    var v95 uint16 = 42
                                                    _ = v95
                                                }
                                                var v96 uint8 = 255
                                                _ = v96
                                            }
                                            _ = v77
                                        }
                                    }
                                    v5 = (((1e10 + 1e10) - v5) / (3.14159 * v5))
                                } else {
                                    v97 := GenS9{F0: -1, F1: 42}
                                    _ = v97
                                }
                            }
                            _ = fi67
                        }
                        var v98 int64 = -1
                        for fi99 := 0; fi99 < 17; fi99++ {
                            v100 := GenS10{F0: 439, F1: 42, F2: ((v1 + "a") + "\n\t")}
                            v2 = append(v2, fi99)
                            _ = fi99
                            _ = v100
                        }
                        _ = v98
                    }
                }
            } else {
                if (100 == -128) {
                    _ = len(v2)
                    v5 -= (v5 + 1.5)
                    for fi101 := 0; fi101 < 8; fi101++ {
                        if false {
                            for fi102 := 0; fi102 < 7; fi102++ {
                                v2 = append(v2, fi102)
                                if (!true) {
                                    if len(v2) > 0 {
                                        v2[0] = (fi102 + (fi101 ^ (fi102 & 127)))
                                    }
                                    v103 := GenS11{F0: 0, F1: 0}
                                    _ = v103
                                } else {
                                    var v104 int = fi101
                                    var v105 int32 = 837
                                    v105 += v105
                                    _ = v104
                                    _ = v105
                                }
                                _ = fi102
                            }
                            v2 = append(v2, fi101)
                            for fi106 := 0; fi106 < 4; fi106++ {
                                if len(v2) > 0 {
                                    v2[0] = (fi101 | (^(fi106)))
                                }
                                v5 += v5
                                _ = fi106
                            }
                        } else {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                        switch (fi101 % ((fi101) | 1)) {
                        case -2147483648:
                            var v107 int = (^(fi101))
                            _ = v107
                        case -1:
                            v5 = 3.14159
                        case 9223372036854775807:
                            if (42 >= 850) {
                                if (!true) {
                                    var v108 uint16 = 255
                                    if ((4294967295 == 2) && false) {
                                        var v109 float32 = v5
                                        _ = v109
                                    }
                                    _ = v108
                                }
                            }
                        }
                        var v110 int = -1
                        _ = fi101
                        _ = v110
                    }
                }
            }
            switch (((v1 + v1) + v1) + v1) {
            case "hello":
                for fi111 := 0; fi111 < 9; fi111++ {
                    v2 = append(v2, (fi111 * (fi111 - (fi111 | fi111))))
                    var v112 int16 = 1000
                    _ = fi111
                    _ = v112
                }
            default:
                v5 += float32(v4)
            }
        }
    case "a":
        var v113 int8 = -1
        _ = v113
    case "hello":
        var v114 string = v1
        _ = v114
    default:
        if len(v3) > 0 {
            v3[0] = v1
        }
    }
    for fi115 := 0; fi115 < 4; fi115++ {
        if ((v1 + "") != v1) {
            _ = len(v3)
        } else {
            var v116 int = (fi115 | 383)
            _ = v116
        }
        var v117 int64 = 9223372036854775807
        if true {
            fmt.Println(v4, v4)
        } else {
            v118 := GenS7{F0: 42, F1: 100}
            _ = v118
        }
        _ = fi115
        _ = v117
    }
    var v119 int32 = 1
    var v120 bool = false
    if false {
        switch 32768 {
        case 842:
            v1 = (("foo bar" + ("\n\t" + v1)) + v1)
        }
    } else {
        if false {
            if (!(-1 < (v119 - (^(v119))))) {
                for fi121 := 0; fi121 < 1; fi121++ {
                    if (fi121 < 0) {
                        var v122 float64 = v4
                        v119 = (v119 % ((v119) | 1))
                        v120 = (((v1 + v1) + v1) <= v1)
                        _ = v122
                    }
                    _ = fi121
                }
                v5 = 1e10
                var v123 uint64 = 4294967295
                _ = v123
            } else {
                v120 = (18446744073709551615 == 1)
                var v124 int16 = 669
                for fi125 := 0; fi125 < 7; fi125++ {
                    for fi126 := 0; fi126 < 17; fi126++ {
                        v120 = false
                        if len(v2) > 0 {
                            v2[0] = -125
                        }
                        if true {
                            var v127 int32 = genFn0((v119 + (v119 * genFn0(v119, 2, ""))), 255, v1)
                            if (-32768 >= 32767) {
                                v128 := GenS11{F0: 255, F1: 78}
                                var v129 int32 = 0
                                v3 = append(v3, v1)
                                _ = v128
                                _ = v129
                            } else {
                                v130 := GenS1{F0: (^(fi126)), F1: 1, F2: 18446744073709551615}
                                v131 := GenS2{F0: -98, F1: 42, F2: fi125}
                                _ = v130
                                _ = v131
                            }
                            v119 = v119
                            _ = v127
                        } else {
                            var v132 uint32 = genFn1(-348, "")
                            _ = v132
                        }
                        _ = fi126
                    }
                    var v133 float32 = v5
                    _ = fi125
                    _ = v133
                }
                _ = v124
            }
        } else {
            v5 += (v5 * (v5 + (((float32(v4) * float32(v4)) + v5) * 0.0)))
        }
    }
    if len(v2) > 0 {
        v2[0] = -9223372036854775808
    }
    v120 = false
    v1 = (v1 + v1)
    var v134 int16 = 1000
    fmt.Println(v1, v119)
    switch 4294967295 {
    case 581:
        if false {
            v135 := GenS6{F0: -1, F1: 2}
            var v136 int64 = 0
            for fi137 := 0; fi137 < 16; fi137++ {
                if v120 {
                    switch "a" {
                    case "":
                        v138 := GenS7{F0: 72, F1: -1}
                        _ = v138
                    case "hello":
                        v139 := GenS1{F0: (fi137 | 127), F1: 319, F2: 1}
                        _ = v139
                    case "\n\t":
                        v140 := GenS2{F0: 127, F1: 0, F2: -202}
                        _ = v140
                    default:
                        if ((v119 / ((1) | 1)) < v119) {
                            var v141 uint16 = 1
                            _ = v141
                        } else {
                            var v142 int64 = v136
                            for fi143 := 0; fi143 < 11; fi143++ {
                                v5 -= (v5 - 3.14159)
                                v142 += (-(v142))
                                var v144 uint16 = 42
                                _ = fi143
                                _ = v144
                            }
                            switch fi137 {
                            case 42:
                                if v120 {
                                    _ = len(v2)
                                    v4 = v4
                                    if (65535 >= 511) {
                                        if v120 {
                                            if v120 {
                                                v120 = true
                                                switch (fi137 << ((318) & 63)) {
                                                case -9223372036854775808:
                                                    var v145 uint16 = 718
                                                    _ = v145
                                                case 2147483647:
                                                    for fi146 := 0; fi146 < 19; fi146++ {
                                                        if (((!v120) && (float64(float32((((v4 / v4) / float64(v5)) - ((1e10 - v4) + v4)))) == float64(v5))) && v120) {
                                                            var v147 float32 = (v5 / (0.0 - v5))
                                                            _ = v147
                                                        }
                                                        v148 := GenS12{F0: (v5 == (0.0 - v5)), F1: fi146}
                                                        _ = fi146
                                                        _ = v148
                                                    }
                                                case 4294967295:
                                                    fmt.Println(v119, v134)
                                                }
                                                if ((255 != 65535) || v120) {
                                                    fmt.Println(v120)
                                                    v2 = append(v2, -831)
                                                } else {
                                                    fmt.Println(fi137, v142, v136)
                                                }
                                            }
                                            if false {
                                                for fi149 := 0; fi149 < 19; fi149++ {
                                                    var v150 bool = true
                                                    if (v150 && true) {
                                                        v120 = false
                                                    }
                                                    _ = fi149
                                                    _ = v150
                                                }
                                                switch (v1 + v1) {
                                                case "foo bar":
                                                    fmt.Println(v5, v134, v120)
                                                default:
                                                    v4 += 0.1
                                                }
                                            } else {
                                                switch (fi137 % (((-(fi137))) | 1)) {
                                                case 42:
                                                    fmt.Println(v136, v1, v120)
                                                case 127:
                                                    v120 = (0 < 4294967295)
                                                case -2147483648:
                                                    for fi151 := 0; fi151 < 17; fi151++ {
                                                        var v152 uint8 = 1
                                                        var v153 float32 = 3.14159
                                                        _ = fi151
                                                        _ = v152
                                                        _ = v153
                                                    }
                                                }
                                                var v154 int = 139
                                                _ = v154
                                            }
                                        } else {
                                            v155 := GenS13{F0: 621}
                                            if v120 {
                                                var v156 int32 = (v119 % ((-619) | 1))
                                                for fi157 := 0; fi157 < 12; fi157++ {
                                                    v158 := GenS14{F0: 0}
                                                    var v159 float32 = ((v5 * ((v5 * v5) / (v5 / ((float32(3.14159) * v5) * v5)))) + 0.1)
                                                    if true {
                                                        v160 := GenS15{F0: (v136 << ((1) & 63)), F1: 42}
                                                        _ = v160
                                                    } else {
                                                        v134 = (v134 & v134)
                                                        v161 := GenS2{F0: -126, F1: v136, F2: fi157}
                                                        _ = len(v3)
                                                        _ = v161
                                                    }
                                                    _ = fi157
                                                    _ = v158
                                                    _ = v159
                                                }
                                                if v120 {
                                                    v134 += (v134 | (v134 << ((1) & 63)))
                                                } else {
                                                    if (!(((v120 && v120) && (42 > 4294967295)) || ((!false) || (2147483648 > 1)))) {
                                                        v136 += (v136 / ((42) | 1))
                                                    } else {
                                                        switch (-(fi137)) {
                                                        case -32:
                                                            fmt.Println(v5, v134)
                                                        }
                                                        switch fi137 {
                                                        case 32767:
                                                            switch "foo bar" {
                                                            case "":
                                                                switch -556 {
                                                                case 32767:
                                                                    if (!(!v120)) {
                                                                        var v162 uint64 = 2
                                                                        v2 = append(v2, fi137)
                                                                        _ = v162
                                                                    } else {
                                                                        v5 = v5
                                                                    }
                                                                case -9223372036854775808:
                                                                    v142 -= int64(v156)
                                                                case -865:
                                                                    if len(v2) > 0 {
                                                                        v2[0] = 550
                                                                    }
                                                                }
                                                            case "\n\t":
                                                                for fi163 := 0; fi163 < 3; fi163++ {
                                                                    v164 := GenS16{F0: (v142 << ((65535) & 63))}
                                                                    fmt.Println(v120, v120)
                                                                    _ = fi163
                                                                    _ = v164
                                                                }
                                                            default:
                                                                var v165 int = (fi137 & (fi137 ^ -261))
                                                                _ = v165
                                                            }
                                                        case -415:
                                                            var v166 uint16 = 42
                                                            _ = v166
                                                        }
                                                        if (4294967295 != 255) {
                                                            var v167 int64 = (v142 % ((v142) | 1))
                                                            _ = v167
                                                        }
                                                    }
                                                    var v168 int16 = (v134 / ((v134) | 1))
                                                    for fi169 := 0; fi169 < 4; fi169++ {
                                                        v3 = append(v3, "\n\t")
                                                        if true {
                                                            if (((!v120) || (v120 || true)) && v120) {
                                                                var v170 uint = 430
                                                                v171 := GenS17{F0: 589}
                                                                v136 -= v136
                                                                var v172 int8 = -114
                                                                if v120 {
                                                                    var v173 bool = (genFn1((v136 ^ int64(v119)), "foo bar") != 134)
                                                                    _ = v173
                                                                }
                                                                _ = v170
                                                                _ = v171
                                                                _ = v172
                                                            } else {
                                                                if v120 {
                                                                    var v174 int16 = -1
                                                                    v119 = (-(v156))
                                                                    v175 := GenS13{F0: 4294967295}
                                                                    _ = v174
                                                                    _ = v175
                                                                } else {
                                                                    v3 = append(v3, v1)
                                                                    fmt.Println(v156)
                                                                    if (false || v120) {
                                                                        v156 = (-(v156))
                                                                        var v176 uint64 = 138
                                                                        fmt.Println(v136)
                                                                        _ = v176
                                                                    }
                                                                }
                                                                v142 = v136
                                                                if ((int32(v142) > 998) && ((v168 | (v134 / ((42) | 1))) != -986)) {
                                                                    if (65535 != 0) {
                                                                        if v120 {
                                                                            if (2 > 127) {
                                                                                switch fi169 {
                                                                                case 32768:
                                                                                    if v120 {
                                                                                        v134 = -559
                                                                                        var v177 bool = (18446744073709551615 > 0)
                                                                                        _ = v177
                                                                                    } else {
                                                                                        if (648 > 660) {
                                                                                            v134 = 32767
                                                                                            v4 += 0.1
                                                                                        }
                                                                                        v178 := GenS2{F0: int8(v156), F1: 1, F2: -96}
                                                                                        _ = v178
                                                                                    }
                                                                                case 314:
                                                                                    if len(v2) > 0 {
                                                                                        v2[0] = (-(fi137))
                                                                                    }
                                                                                default:
                                                                                    v134 = -723
                                                                                    fmt.Println(v168, fi169, v4)
                                                                                }
                                                                                _ = len(v3)
                                                                                var v179 float32 = ((float32((v4 + v4)) * v5) / 1.0)
                                                                                _ = v179
                                                                            } else {
                                                                                v180 := GenS9{F0: 44, F1: 0}
                                                                                _ = v180
                                                                            }
                                                                            fmt.Println(v168)
                                                                        } else {
                                                                            fmt.Println(v119, v119, v168)
                                                                            switch fi137 {
                                                                            case 0:
                                                                                v120 = (v1 < ("a" + (v1 + "hello")))
                                                                            case 128:
                                                                                if len(v2) > 0 {
                                                                                    v2[0] = fi137
                                                                                }
                                                                            case 2147483647:
                                                                                v142 -= 42
                                                                            }
                                                                            if false {
                                                                                if ("foo bar" >= ((v1 + "foo bar") + v1)) {
                                                                                    var v181 string = (v1 + v1)
                                                                                    v5 = 1.5
                                                                                    if (0 < 2) {
                                                                                        var v182 uint64 = 2
                                                                                        _ = v182
                                                                                    }
                                                                                    _ = v181
                                                                                } else {
                                                                                    switch (fi169 & fi169) {
                                                                                    case 0:
                                                                                        var v183 int32 = (v119 - v156)
                                                                                        _ = v183
                                                                                    default:
                                                                                        if (fi169 < 42) {
                                                                                            var v184 uint64 = 2
                                                                                            _ = v184
                                                                                        } else {
                                                                                            v185 := GenS4{F0: v142, F1: 1, F2: 42}
                                                                                            var v186 float64 = v4
                                                                                            var v187 uint64 = 1
                                                                                            _ = v185
                                                                                            _ = v186
                                                                                            _ = v187
                                                                                        }
                                                                                        v4 += v4
                                                                                    }
                                                                                    v156 = -2147483648
                                                                                    v142 += int64(v156)
                                                                                }
                                                                            }
                                                                        }
                                                                        if v120 {
                                                                            v119 += v119
                                                                            v188 := GenS18{F0: 6}
                                                                            v189 := GenS19{F0: 0}
                                                                            _ = v188
                                                                            _ = v189
                                                                        } else {
                                                                            if false {
                                                                                switch v1 {
                                                                                case "foo bar":
                                                                                    v142 += 739
                                                                                case "\n\t":
                                                                                    v190 := GenS10{F0: 65535, F1: -1, F2: v1}
                                                                                    _ = v190
                                                                                }
                                                                                v191 := GenS11{F0: 42, F1: 255}
                                                                                _ = v191
                                                                            }
                                                                            v1 = "foo bar"
                                                                            v5 += 1e10
                                                                            if (genFn1(v136, "foo bar") <= 721) {
                                                                                if v120 {
                                                                                    v168 = (v168 >> ((355) & 63))
                                                                                    v192 := GenS4{F0: (v142 / (((^(v136))) | 1)), F1: 351, F2: 4294967295}
                                                                                    var v193 int16 = v168
                                                                                    _ = v192
                                                                                    _ = v193
                                                                                } else {
                                                                                    v156 += 0
                                                                                    v136 -= (^(v136))
                                                                                    v194 := GenS20{F0: v168, F1: (255 >= 1), F2: 0}
                                                                                    _ = v194
                                                                                }
                                                                                var v195 uint = 255
                                                                                _ = v195
                                                                            }
                                                                            var v196 int8 = -128
                                                                            _ = v196
                                                                        }
                                                                    } else {
                                                                        v5 = (-2.25 / 1e10)
                                                                        _ = len(v3)
                                                                    }
                                                                }
                                                                var v197 string = (v1 + v1)
                                                                _ = v197
                                                            }
                                                            v156 = 1
                                                            var v198 float32 = float32(v4)
                                                            _ = v198
                                                        }
                                                        _ = fi169
                                                    }
                                                    _ = v168
                                                }
                                                _ = v156
                                            }
                                            _ = v155
                                        }
                                        fmt.Println(v120, v120, fi137)
                                        if len(v3) > 0 {
                                            v3[0] = v1
                                        }
                                    } else {
                                        var v199 int32 = (v119 * (v119 - 2147483647))
                                        if v120 {
                                            for fi200 := 0; fi200 < 19; fi200++ {
                                                v201 := GenS21{F0: 42, F1: 717, F2: 42}
                                                var v202 int64 = v142
                                                _ = fi200
                                                _ = v201
                                                _ = v202
                                            }
                                        }
                                        _ = v199
                                    }
                                } else {
                                    _ = len(v3)
                                    var v203 string = ((("foo bar" + (("hello" + "hello") + "hello")) + ((v1 + "foo bar") + v1)) + (v1 + v1))
                                    _ = v203
                                }
                            case 127:
                                v204 := GenS1{F0: (^(fi137)), F1: 2, F2: 4294967295}
                                _ = v204
                            case 32767:
                                switch (v1 + "foo bar") {
                                case "foo bar":
                                    if len(v3) > 0 {
                                        v3[0] = ""
                                    }
                                case "\n\t":
                                    v134 = v134
                                case "hello":
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                default:
                                    for fi205 := 0; fi205 < 9; fi205++ {
                                        v120 = (32767 > 1)
                                        fmt.Println(v142, v1)
                                        _ = fi205
                                    }
                                }
                            default:
                                if false {
                                    var v206 int64 = (v142 & v142)
                                    fmt.Println(v120, v5)
                                    v207 := GenS6{F0: fi137, F1: 304}
                                    _ = v206
                                    _ = v207
                                } else {
                                    var v208 uint8 = 128
                                    _ = v208
                                }
                            }
                            _ = v142
                        }
                    }
                } else {
                    fmt.Println(v136)
                }
                if (18446744073709551615 != 18446744073709551615) {
                    if (!("hello" != ("foo bar" + v1))) {
                        v209 := GenS15{F0: 765, F1: 127}
                        _ = v209
                    }
                    v210 := GenS22{F0: (v136 - (v136 ^ (v136 & (v136 << ((65535) & 63))))), F1: (-(v134)), F2: (v136 <= -563)}
                    _ = v210
                } else {
                    for fi211 := 0; fi211 < 17; fi211++ {
                        var v212 uint16 = 255
                        v213 := GenS1{F0: (-(fi137)), F1: v212, F2: 18446744073709551615}
                        var v214 float32 = ((v5 - -2.25) + ((v5 - -2.25) / -2.25))
                        _ = fi211
                        _ = v212
                        _ = v213
                        _ = v214
                    }
                    var v215 int32 = (v119 ^ genFn0((v119 >> ((0) & 63)), 4294967295, v1))
                    var v216 bool = false
                    _ = v215
                    _ = v216
                }
                _ = fi137
            }
            _ = v135
            _ = v136
        } else {
            v119 -= genFn0(genFn0(42, 18446744073709551615, (v1 + v1)), 1, (((v1 + v1) + "foo bar") + "a"))
        }
    case 32767:
        var v217 uint32 = 65535
        _ = v217
    default:
        v120 = (152 == 0)
    }
    var v218 uint64 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v119:", v119)
    fmt.Println("v120:", v120)
    fmt.Println("v134:", v134)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v218:", v218)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v119
    _ = v120
    _ = v134
    _ = v2
    _ = v218
    _ = v3
    _ = v4
    _ = v5
}
