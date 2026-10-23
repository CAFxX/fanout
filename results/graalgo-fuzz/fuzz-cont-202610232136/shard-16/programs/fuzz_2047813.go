package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8; F1 int8 }
type GenS5 struct { F0 int32; F1 uint; F2 bool }
type GenS6 struct { F0 int8; F1 uint64 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 uint; F1 string; F2 bool }
type GenS12 struct { F0 uint }
func genFn0(p0 int32, p1 uint8, p2 uint) uint8 {
    var v5 float32 = (-2.25 * 3.14159)
    p1 = 131
    for fi6 := 0; fi6 < 3; fi6++ {
        for fi7 := 0; fi7 < 15; fi7++ {
            var v8 uint64 = 2
            v5 = v5
            _ = fi7
            _ = v8
        }
        v9 := GenS1{F0: 692, F1: p2}
        _ = fi6
        _ = v9
    }
    var v10 uint16 = 255
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v5
    return p1
}
func genFn1(p0 uint8, p1 int8) uint64 {
    switch ("hello" + (("hello" + ("foo bar" + ("" + "\n\t"))) + "foo bar")) {
    case "hello":
        var v11 int8 = (p1 / ((p1) | 1))
        _ = v11
    case "foo bar":
        var v12 uint = 2
        _ = v12
    case "\n\t":
        for fi13 := 0; fi13 < 7; fi13++ {
            v14 := []int{-225, (-(fi13)), fi13}
            _ = fi13
            _ = v14
        }
    default:
        if (46 != 4294967295) {
            v15 := []int{398, int(p0), 1}
            if false {
                fmt.Println(p1)
                if false {
                    switch "foo bar" {
                    case "hello":
                        p0 -= (p0 - p0)
                    default:
                        if true {
                            switch 503 {
                            case 4294967295:
                                for fi16 := 0; fi16 < 3; fi16++ {
                                    var v17 float64 = ((1e10 * (((3.14159 / float64(-2.25)) + 1.5) - 0.0)) / -2.25)
                                    for fi18 := 0; fi18 < 13; fi18++ {
                                        var v19 int = (fi16 | (fi18 >> ((336) & 63)))
                                        v15 = append(v15, (^(fi18)))
                                        if true {
                                            var v20 float32 = 1.5
                                            if (genFn0(-2147483648, (p0 ^ p0), 18446744073709551615) != (^(p0))) {
                                                if false {
                                                    v20 -= v20
                                                    var v21 int = fi16
                                                    var v22 float32 = 0.0
                                                    _ = v21
                                                    _ = v22
                                                } else {
                                                    var v23 string = "\n\t"
                                                    switch -1 {
                                                    case 318:
                                                        v19 -= (fi18 % ((v19) | 1))
                                                    case 1:
                                                        if false {
                                                            var v24 uint8 = (p0 / ((2) | 1))
                                                            fmt.Println(fi16, v17)
                                                            _ = v24
                                                        } else {
                                                            v25 := GenS2{F0: -523}
                                                            _ = v25
                                                        }
                                                    }
                                                    v26 := GenS3{F0: 255}
                                                    _ = v23
                                                    _ = v26
                                                }
                                                v19 = (v19 % (((fi18 - (fi18 + -906))) | 1))
                                                switch "\n\t" {
                                                case "foo bar":
                                                    v27 := GenS3{F0: uint(p0)}
                                                    _ = v27
                                                case "hello":
                                                    var v28 bool = false
                                                    _ = v28
                                                case "\n\t":
                                                    fmt.Println(fi18, fi18, p0)
                                                default:
                                                    var v29 uint8 = p0
                                                    _ = v29
                                                }
                                            }
                                            var v30 string = "foo bar"
                                            _ = v20
                                            _ = v30
                                        }
                                        _ = fi18
                                        _ = v19
                                    }
                                    _ = fi16
                                    _ = v17
                                }
                            case 128:
                                p1 += (p1 & -80)
                            case 32767:
                                if ("hello" <= "a") {
                                    for fi31 := 0; fi31 < 5; fi31++ {
                                        var v32 uint16 = 42
                                        var v33 int8 = p1
                                        _ = fi31
                                        _ = v32
                                        _ = v33
                                    }
                                    var v34 float32 = (((1e10 - (0.0 + 3.14159)) - 0.1) * float32(-2.25))
                                    v35 := GenS2{F0: 104}
                                    _ = v34
                                    _ = v35
                                } else {
                                    for fi36 := 0; fi36 < 17; fi36++ {
                                        for fi37 := 0; fi37 < 6; fi37++ {
                                            var v38 int16 = -752
                                            if false {
                                                var v39 uint64 = 42
                                                var v40 float32 = (0.1 * (((1.5 + -2.25) + 1e10) + 1.5))
                                                _ = v39
                                                _ = v40
                                            }
                                            _ = fi37
                                            _ = v38
                                        }
                                        v41 := GenS4{F0: p1, F1: 31}
                                        var v42 float32 = 1e10
                                        _ = fi36
                                        _ = v41
                                        _ = v42
                                    }
                                    fmt.Println(p0, p0)
                                    switch "a" {
                                    case "foo bar":
                                        if len(v15) > 0 {
                                            v15[0] = 1
                                        }
                                    case "a":
                                        p0 -= (p0 % ((genFn0(-1, genFn0(-2147483648, (p0 / ((p0) | 1)), 880), 65535)) | 1))
                                    case "":
                                        if true {
                                            v43 := GenS5{F0: 723, F1: 4294967295, F2: false}
                                            _ = v43
                                        } else {
                                            var v44 uint16 = 2
                                            p0 -= p0
                                            for fi45 := 0; fi45 < 5; fi45++ {
                                                v44 += v44
                                                fmt.Println(fi45)
                                                fmt.Println(p1, p0)
                                                _ = fi45
                                            }
                                            _ = v44
                                        }
                                    default:
                                        v46 := GenS3{F0: 598}
                                        _ = v46
                                    }
                                }
                            default:
                                switch -1 {
                                case 32768:
                                    switch 32768 {
                                    case 32767:
                                        p1 = p1
                                    case 9223372036854775807:
                                        if true {
                                            var v47 uint32 = 307
                                            switch "a" {
                                            case "\n\t":
                                                var v48 int16 = -1
                                                _ = v48
                                            case "hello":
                                                fmt.Println(p0, p0)
                                            }
                                            _ = v47
                                        } else {
                                            if false {
                                                p1 = p1
                                            } else {
                                                for fi49 := 0; fi49 < 6; fi49++ {
                                                    if true {
                                                        for fi50 := 0; fi50 < 6; fi50++ {
                                                            if true {
                                                                v15 = append(v15, fi50)
                                                                switch "a" {
                                                                case "foo bar":
                                                                    p1 = p1
                                                                }
                                                            } else {
                                                                p1 += p1
                                                            }
                                                            switch (fi50 % ((fi49) | 1)) {
                                                            case 128:
                                                                var v51 uint32 = 90
                                                                _ = v51
                                                            case 42:
                                                                if (1 < 9223372036854775807) {
                                                                    if len(v15) >= 2 {
                                                                        v15 = v15[1:len(v15)]
                                                                    }
                                                                    switch "foo bar" {
                                                                    case "\n\t":
                                                                        if true {
                                                                            var v52 uint64 = 240
                                                                            if true {
                                                                                if true {
                                                                                    var v53 bool = true
                                                                                    _ = v53
                                                                                } else {
                                                                                    switch fi50 {
                                                                                    case 4294967295:
                                                                                        var v54 int64 = -9223372036854775808
                                                                                        _ = v54
                                                                                    }
                                                                                    p1 = -128
                                                                                }
                                                                                if ((-2.25 * (0.0 * 0.1)) >= 3.14159) {
                                                                                    var v55 uint = 2
                                                                                    switch ("a" + "foo bar") {
                                                                                    case "\n\t":
                                                                                        v55 += (v55 * 677)
                                                                                    case "foo bar":
                                                                                        if (true && false) {
                                                                                            v52 += v52
                                                                                            v56 := GenS6{F0: p1, F1: (v52 << ((0) & 63))}
                                                                                            var v57 int32 = 2147483647
                                                                                            _ = v56
                                                                                            _ = v57
                                                                                        }
                                                                                    case "a":
                                                                                        if true {
                                                                                            var v58 bool = ((fi49 % ((fi49) | 1)) < (-(fi49)))
                                                                                            if len(v15) > 0 {
                                                                                                v15[0] = -29
                                                                                            }
                                                                                            fmt.Println(v55, fi49)
                                                                                            _ = v58
                                                                                        }
                                                                                    }
                                                                                    switch fi50 {
                                                                                    case 75:
                                                                                        p0 = 151
                                                                                    case 127:
                                                                                        _ = len(v15)
                                                                                    default:
                                                                                        var v59 int64 = -523
                                                                                        _ = v59
                                                                                    }
                                                                                    _ = v55
                                                                                } else {
                                                                                    var v60 int64 = 0
                                                                                    switch 788 {
                                                                                    case -510:
                                                                                        v61 := GenS6{F0: p1, F1: v52}
                                                                                        _ = v61
                                                                                    case 9223372036854775807:
                                                                                        v62 := GenS3{F0: uint(v52)}
                                                                                        _ = v62
                                                                                    }
                                                                                    _ = v60
                                                                                }
                                                                            }
                                                                            _ = v52
                                                                        } else {
                                                                            var v63 int64 = -674
                                                                            v64 := GenS3{F0: 0}
                                                                            switch (("a" + "\n\t") + "hello") {
                                                                            case "\n\t":
                                                                                if (false && false) {
                                                                                    var v65 bool = true
                                                                                    if (553 <= 0) {
                                                                                        fmt.Println(fi50, fi50, v65)
                                                                                        p0 += p0
                                                                                        var v66 float32 = ((3.14159 * float32((0.0 / 1.0))) / -2.25)
                                                                                        _ = v66
                                                                                    }
                                                                                    var v67 uint8 = p0
                                                                                    if ((fi49 >> ((570) & 63)) == fi49) {
                                                                                        v68 := GenS7{F0: "a"}
                                                                                        v67 += p0
                                                                                        _ = v68
                                                                                    } else {
                                                                                        if v65 {
                                                                                            if (42 <= 42) {
                                                                                                switch fi50 {
                                                                                                case 550:
                                                                                                    var v69 int8 = (p1 ^ 0)
                                                                                                    _ = v69
                                                                                                default:
                                                                                                    v63 = -1
                                                                                                }
                                                                                            } else {
                                                                                                var v70 float32 = -2.25
                                                                                                var v71 int32 = 648
                                                                                                _ = v70
                                                                                                _ = v71
                                                                                            }
                                                                                            var v72 int64 = (v63 % ((v63) | 1))
                                                                                            var v73 int8 = (-(p1))
                                                                                            _ = v72
                                                                                            _ = v73
                                                                                        }
                                                                                        if v65 {
                                                                                            if len(v15) >= 2 {
                                                                                                v15 = v15[1:len(v15)]
                                                                                            }
                                                                                            var v74 bool = (v65 && false)
                                                                                            if true {
                                                                                                var v75 int64 = v63
                                                                                                v15 = append(v15, fi49)
                                                                                                _ = v75
                                                                                            } else {
                                                                                                switch (fi49 + (fi49 & fi49)) {
                                                                                                case -9223372036854775808:
                                                                                                    switch fi49 {
                                                                                                    case -1:
                                                                                                        var v76 int64 = (^(v63))
                                                                                                        _ = v76
                                                                                                    }
                                                                                                }
                                                                                                var v77 uint = 42
                                                                                                if len(v15) >= 2 {
                                                                                                    v15 = v15[1:len(v15)]
                                                                                                }
                                                                                                _ = v77
                                                                                            }
                                                                                            if len(v15) >= 2 {
                                                                                                v15 = v15[1:len(v15)]
                                                                                            }
                                                                                            _ = v74
                                                                                        }
                                                                                        var v78 int64 = -616
                                                                                        fmt.Println(fi49)
                                                                                        _ = v78
                                                                                    }
                                                                                    _ = v65
                                                                                    _ = v67
                                                                                } else {
                                                                                    var v79 uint64 = 18446744073709551615
                                                                                    if true {
                                                                                        var v80 int8 = p1
                                                                                        if false {
                                                                                            v81 := GenS3{F0: 177}
                                                                                            v63 = 9223372036854775807
                                                                                            p1 = (^(p1))
                                                                                            _ = v81
                                                                                        } else {
                                                                                            v82 := GenS1{F0: 2, F1: 95}
                                                                                            v83 := GenS8{F0: 513}
                                                                                            _ = v82
                                                                                            _ = v83
                                                                                        }
                                                                                        _ = v80
                                                                                    } else {
                                                                                        var v84 uint8 = (^(p0))
                                                                                        v84 = v84
                                                                                        if len(v15) > 0 {
                                                                                            v15[0] = fi49
                                                                                        }
                                                                                        _ = v84
                                                                                    }
                                                                                    _ = v79
                                                                                }
                                                                                v85 := GenS7{F0: "a"}
                                                                                var v86 uint8 = 0
                                                                                switch "" {
                                                                                case "\n\t":
                                                                                    var v87 int64 = 9223372036854775807
                                                                                    _ = v87
                                                                                default:
                                                                                    v15 = append(v15, fi50)
                                                                                }
                                                                                if true {
                                                                                    var v88 uint16 = 341
                                                                                    if (4294967295 == 0) {
                                                                                        v86 = 128
                                                                                        p0 += v86
                                                                                        var v89 int = (fi49 & 2147483647)
                                                                                        _ = v89
                                                                                    }
                                                                                    if (fi50 <= (^(fi50))) {
                                                                                        var v90 int64 = -9223372036854775808
                                                                                        var v91 string = ""
                                                                                        v92 := GenS5{F0: -1, F1: 766, F2: false}
                                                                                        _ = v90
                                                                                        _ = v91
                                                                                        _ = v92
                                                                                    } else {
                                                                                        switch "hello" {
                                                                                        case "foo bar":
                                                                                            if len(v15) > 0 {
                                                                                                v15[0] = fi49
                                                                                            }
                                                                                        case "\n\t":
                                                                                            v63 += v63
                                                                                        case "hello":
                                                                                            _ = len(v15)
                                                                                        }
                                                                                        var v93 float64 = 1.5
                                                                                        _ = v93
                                                                                    }
                                                                                    var v94 int8 = int8(fi50)
                                                                                    _ = v88
                                                                                    _ = v94
                                                                                } else {
                                                                                    v95 := GenS1{F0: 32767, F1: 998}
                                                                                    var v96 uint = 838
                                                                                    _ = v95
                                                                                    _ = v96
                                                                                }
                                                                                _ = v85
                                                                                _ = v86
                                                                            default:
                                                                                if false {
                                                                                    v97 := GenS9{F0: 517}
                                                                                    _ = v97
                                                                                } else {
                                                                                    var v98 float32 = ((0.1 - 1.5) * float32((-2.25 + 3.14159)))
                                                                                    _ = v98
                                                                                }
                                                                            }
                                                                            _ = v63
                                                                            _ = v64
                                                                        }
                                                                    case "foo bar":
                                                                        p1 = -9
                                                                    default:
                                                                        var v99 int16 = -942
                                                                        _ = v99
                                                                    }
                                                                } else {
                                                                    var v100 int8 = p1
                                                                    _ = v100
                                                                }
                                                            default:
                                                                var v101 int = 9223372036854775807
                                                                _ = v101
                                                            }
                                                            _ = fi50
                                                        }
                                                    } else {
                                                        v15 = append(v15, 4294967295)
                                                        fmt.Println(p0)
                                                    }
                                                    for fi102 := 0; fi102 < 5; fi102++ {
                                                        p1 -= -128
                                                        p0 = p0
                                                        _ = fi102
                                                    }
                                                    v103 := GenS2{F0: 0}
                                                    _ = fi49
                                                    _ = v103
                                                }
                                                p1 = p1
                                                _ = len(v15)
                                            }
                                            v104 := GenS6{F0: 127, F1: 28}
                                            _ = v104
                                        }
                                    case 4294967295:
                                        if (14 <= 4294967295) {
                                            if true {
                                                switch "" {
                                                case "a":
                                                    switch "a" {
                                                    case "foo bar":
                                                        var v105 float64 = (-2.25 - ((1.5 - -2.25) - (float64((((((1e10 / 0.1) + 0.1) * 0.1) / float32((3.14159 * ((0.1 - 0.0) * (float64(-2.25) - 3.14159))))) / (float32(1.5) + (0.1 * 3.14159)))) - 1e10)))
                                                        _ = v105
                                                    default:
                                                        v106 := GenS3{F0: 2}
                                                        _ = v106
                                                    }
                                                case "foo bar":
                                                    var v107 int8 = (p1 & (p1 >> ((42) & 63)))
                                                    _ = v107
                                                case "\n\t":
                                                    if false {
                                                        switch 128 {
                                                        case -1:
                                                            if (p1 != p1) {
                                                                if len(v15) >= 2 {
                                                                    v15 = v15[1:len(v15)]
                                                                }
                                                            }
                                                        case 32767:
                                                            var v108 uint64 = 415
                                                            _ = v108
                                                        default:
                                                            if len(v15) >= 2 {
                                                                v15 = v15[1:len(v15)]
                                                            }
                                                        }
                                                        _ = len(v15)
                                                    } else {
                                                        fmt.Println(p0, p1)
                                                    }
                                                }
                                                if false {
                                                    var v109 int32 = 153
                                                    switch 32767 {
                                                    case -675:
                                                        v15 = append(v15, 42)
                                                    default:
                                                        if true {
                                                            var v110 uint = 439
                                                            _ = v110
                                                        } else {
                                                            p0 -= (p0 << ((1) & 63))
                                                        }
                                                    }
                                                    _ = v109
                                                }
                                                for fi111 := 0; fi111 < 10; fi111++ {
                                                    var v112 uint32 = 102
                                                    _ = fi111
                                                    _ = v112
                                                }
                                            } else {
                                                if (true && ((32767 < 42) && (!((0.0 - 0.0) < 1e10)))) {
                                                    for fi113 := 0; fi113 < 19; fi113++ {
                                                        p1 = 100
                                                        var v114 uint32 = 4294967295
                                                        if len(v15) > 0 {
                                                            v15[0] = fi113
                                                        }
                                                        _ = fi113
                                                        _ = v114
                                                    }
                                                    for fi115 := 0; fi115 < 16; fi115++ {
                                                        var v116 int64 = -879
                                                        var v117 int64 = (v116 / ((0) | 1))
                                                        _ = fi115
                                                        _ = v116
                                                        _ = v117
                                                    }
                                                    v15 = append(v15, 2147483647)
                                                } else {
                                                    var v118 int32 = 934
                                                    _ = len(v15)
                                                    _ = v118
                                                }
                                                var v119 string = "a"
                                                _ = v119
                                            }
                                        }
                                    }
                                }
                            }
                            if (!(127 != p0)) {
                                var v120 int8 = (p1 % ((p1) | 1))
                                switch -9223372036854775808 {
                                case 42:
                                    _ = len(v15)
                                case 32767:
                                    for fi121 := 0; fi121 < 10; fi121++ {
                                        p0 = (p0 * (^(p0)))
                                        v122 := GenS1{F0: 42, F1: 65535}
                                        var v123 float32 = float32(3.14159)
                                        _ = fi121
                                        _ = v122
                                        _ = v123
                                    }
                                }
                                _ = v120
                            } else {
                                fmt.Println(p1)
                                for fi124 := 0; fi124 < 10; fi124++ {
                                    var v125 bool = true
                                    p0 -= (p0 | (p0 & p0))
                                    var v126 int = (fi124 + fi124)
                                    _ = fi124
                                    _ = v125
                                    _ = v126
                                }
                            }
                            v15 = append(v15, -9223372036854775808)
                        } else {
                            v127 := GenS10{F0: 42}
                            p1 = (-(p1))
                            _ = v127
                        }
                    }
                }
                fmt.Println(p0, p0)
            }
            if false {
                var v128 int = 703
                p0 = genFn0(2147483647, genFn0(1, p0, 18446744073709551615), 2)
                var v129 int32 = 2147483647
                _ = v128
                _ = v129
            }
            _ = v15
        }
    }
    for fi130 := 0; fi130 < 3; fi130++ {
        for fi131 := 0; fi131 < 20; fi131++ {
            var v132 int8 = 127
            var v133 uint8 = genFn0(659, genFn0(-2147483648, (p0 / (((p0 + (p0 - 1))) | 1)), 65535), 668)
            _ = fi131
            _ = v132
            _ = v133
        }
        if (int64(p1) == 42) {
            for fi134 := 0; fi134 < 4; fi134++ {
                v135 := GenS4{F0: 0, F1: (p1 - 127)}
                if false {
                    p1 = -1
                }
                _ = fi134
                _ = v135
            }
            for fi136 := 0; fi136 < 12; fi136++ {
                if true {
                    if (!(0.0 <= 0.1)) {
                        if true {
                            var v137 string = ""
                            _ = v137
                        } else {
                            if true {
                                var v138 uint = 42
                                _ = v138
                            } else {
                                var v139 float32 = 1.5
                                _ = v139
                            }
                        }
                    }
                    fmt.Println(fi136, fi130)
                } else {
                    v140 := GenS11{F0: 4294967295, F1: ("hello" + ""), F2: (p0 > 127)}
                    _ = v140
                }
                v141 := []int{}
                v141 = append(v141, 32767)
                _ = fi136
                _ = v141
            }
        } else {
            var v142 uint32 = 65535
            fmt.Println(p0, p0)
            var v143 float32 = float32(-2.25)
            _ = v142
            _ = v143
        }
        _ = fi130
    }
    p1 -= p1
    v144 := []int{0, 127, 128, 87}
    for fi145 := 0; fi145 < 7; fi145++ {
        _ = len(v144)
        p1 += p1
        if true {
            if (true && true) {
                var v146 int64 = 42
                _ = v146
            }
            var v147 string = "foo bar"
            p1 = (p1 / (((p1 * 0)) | 1))
            _ = v147
        }
        _ = fi145
    }
    _ = p0
    _ = p1
    _ = v144
    return 0
}

func main() {
    v1 := []string{}
    v2 := make(map[uint32]string)
    var v3 uint32 = 2
    var v4 bool = true
    var v148 uint64 = 42
    fmt.Println(v4, v4, v3)
    for fi149 := 0; fi149 < 1; fi149++ {
        var v150 uint = 165
        _ = fi149
        _ = v150
    }
    v148 -= 0
    delete(v2, 42)
    var v151 uint32 = v3
    v152 := GenS12{F0: 18446744073709551615}
    for fi153 := 0; fi153 < 3; fi153++ {
        v154 := v2[(v3 & v151)]
        var v155 uint64 = (v148 + (v148 * genFn1(58, 0)))
        _ = fi153
        _ = v154
        _ = v155
    }
    fmt.Println(v151, v151)
    fmt.Println(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v148:", v148)
    fmt.Println("v151:", v151)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v148
    _ = v151
    _ = v152
    _ = v2
    _ = v3
    _ = v4
}
