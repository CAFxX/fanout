package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 int }
type GenS2 struct { F0 uint16; F1 bool; F2 int16 }
type GenS3 struct { F0 uint32; F1 int8 }
type GenS4 struct { F0 string; F1 int8; F2 uint64 }
type GenS5 struct { F0 uint8; F1 int16; F2 int32 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint16; F1 bool }
type GenS8 struct { F0 uint64; F1 uint8 }
type GenS9 struct { F0 uint16; F1 int }
type GenS10 struct { F0 uint16; F1 uint; F2 string }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 int8; F1 bool; F2 uint8 }
type GenS14 struct { F0 uint64; F1 uint64; F2 int32 }
type GenS15 struct { F0 uint; F1 int16; F2 int32 }
type GenS16 struct { F0 uint32; F1 bool }
type GenS17 struct { F0 uint64; F1 uint; F2 int32 }
func genFn0() bool {
    switch "\n\t" {
    case "hello":
        for fi4 := 0; fi4 < 6; fi4++ {
            if true {
                var v5 int8 = 127
                _ = v5
            } else {
                for fi6 := 0; fi6 < 12; fi6++ {
                    var v7 float64 = 0.1
                    _ = fi6
                    _ = v7
                }
                for fi8 := 0; fi8 < 3; fi8++ {
                    if (false && false) {
                        var v9 string = "hello"
                        if false {
                            v9 = v9
                            v9 = (v9 + v9)
                        }
                        switch (fi4 ^ fi4) {
                        case 345:
                            v10 := GenS1{F0: 2, F1: false, F2: -1}
                            _ = v10
                        default:
                            fmt.Println(fi8, v9, fi4)
                        }
                        _ = v9
                    } else {
                        v11 := GenS2{F0: 255, F1: (3.14159 < -2.25), F2: 42}
                        _ = v11
                    }
                    var v12 int8 = 100
                    if false {
                        var v13 int = 9223372036854775807
                        if (true || true) {
                            var v14 float32 = (float32(0.1) - 3.14159)
                            if true {
                                if false {
                                    var v15 bool = true
                                    _ = v15
                                } else {
                                    v16 := []int{-39, fi4}
                                    if (0 < fi4) {
                                        v16 = append(v16, v13)
                                    }
                                    _ = v16
                                }
                                v12 -= (v12 & (v12 + (v12 * 70)))
                            } else {
                                v13 = 128
                                v17 := GenS3{F0: 0, F1: 0}
                                _ = v17
                            }
                            v18 := []int{(fi8 << ((0) & 63))}
                            v12 = (v12 << ((255) & 63))
                            _ = v14
                            _ = v18
                        } else {
                            fmt.Println(v12)
                        }
                        switch fi8 {
                        case 127:
                            var v19 float32 = (1.5 / (0.0 + ((0.1 - -2.25) / 1e10)))
                            _ = v19
                        case 128:
                            if (-841 != 1) {
                                if false {
                                    var v20 uint64 = 990
                                    var v21 uint = 255
                                    var v22 int64 = -1
                                    _ = v20
                                    _ = v21
                                    _ = v22
                                }
                            } else {
                                if true {
                                    var v23 string = "hello"
                                    v24 := GenS2{F0: 255, F1: true, F2: 42}
                                    var v25 int = -9223372036854775808
                                    _ = v23
                                    _ = v24
                                    _ = v25
                                } else {
                                    var v26 uint = 482
                                    var v27 uint32 = 842
                                    _ = v26
                                    _ = v27
                                }
                            }
                        case 32768:
                            var v28 uint64 = 4294967295
                            _ = v28
                        default:
                            var v29 string = ("hello" + "a")
                            _ = v29
                        }
                        if (v12 >= (v12 ^ (v12 >> ((2) & 63)))) {
                            v30 := GenS4{F0: ("" + "a"), F1: v12, F2: 582}
                            var v31 bool = true
                            v31 = true
                            _ = v30
                            _ = v31
                        }
                        v32 := []int{(fi8 ^ (v13 | fi4))}
                        _ = v13
                        _ = v32
                    }
                    _ = fi8
                    _ = v12
                }
                var v33 uint = 0
                _ = v33
            }
            switch 92 {
            case -407:
                v34 := GenS5{F0: 42, F1: 42, F2: 2147483647}
                _ = v34
            case -349:
                if true {
                    var v35 float64 = 3.14159
                    if true {
                        if true {
                            var v36 uint16 = 2
                            _ = v36
                        } else {
                            if (1 > -1) {
                                v37 := []int{-264}
                                var v38 int16 = -838
                                _ = v37
                                _ = v38
                            }
                        }
                    }
                    var v39 string = "hello"
                    _ = v35
                    _ = v39
                }
            case 127:
                fmt.Println(fi4)
            default:
                var v40 int8 = -110
                _ = v40
            }
            _ = fi4
        }
    }
    var v41 int = 42
    var v42 uint = 508
    v43 := GenS2{F0: 471, F1: false, F2: 536}
    _ = v41
    _ = v42
    _ = v43
    return false
}
func genFn1() uint8 {
    var v44 float64 = ((-2.25 / 1.0) * (((0.1 / -2.25) - 1.5) * 1e10))
    v45 := GenS3{F0: 65535, F1: 0}
    for fi46 := 0; fi46 < 7; fi46++ {
        v47 := GenS6{F0: -1}
        for fi48 := 0; fi48 < 6; fi48++ {
            v49 := []int{}
            _ = fi48
            _ = v49
        }
        _ = fi46
        _ = v47
    }
    v50 := []int{42, 41}
    _ = v44
    _ = v45
    _ = v50
    return 2
}

func main() {
    var v1 string = ""
    v2 := []uint32{272}
    v3 := make(map[int]bool)
    v1 = v1
    v1 = ((((v1 + v1) + (v1 + v1)) + v1) + v1)
    for fi51 := 0; fi51 < 16; fi51++ {
        for fi52 := 0; fi52 < 10; fi52++ {
            v53 := v3[(fi51 ^ (fi52 % (((fi51 & fi52)) | 1)))]
            v3[-465] = true
            _ = fi52
            _ = v53
        }
        _ = fi51
    }
    var v54 uint16 = 1
    v1 = v1
    v54 += 255
    v55 := GenS2{F0: (v54 ^ v54), F1: true, F2: 42}
    switch (("hello" + v1) + v1) {
    case "hello":
        v3[259] = false
    default:
        v1 = (v1 + v1)
    }
    var v56 float64 = 0.0
    var v57 int = -311
    if (true && ((false || (v56 < v56)) || false)) {
        var v58 string = (((v1 + v1) + "") + (v1 + v1))
        var v59 float64 = (((-2.25 - 3.14159) * v56) + (v56 / v56))
        var v60 int64 = 9223372036854775807
        _ = v58
        _ = v59
        _ = v60
    }
    v54 = (v54 ^ (v54 + v54))
    switch (v1 + (v1 + v1)) {
    case "a":
        fmt.Println(v1)
    case "hello":
        v61 := GenS7{F0: 65535, F1: ((float32(1.5) - (0.1 * 1e10)) <= 0.1)}
        _ = v61
    case "":
        fmt.Println(v54, v57)
    }
    var v62 uint8 = 0
    for fi63 := 0; fi63 < 7; fi63++ {
        fmt.Println(v54, v62)
        var v64 int = fi63
        var v65 int32 = 509
        _ = fi63
        _ = v64
        _ = v65
    }
    switch v1 {
    case "hello":
        var v66 int64 = -29
        _ = v66
    case "":
        for fi67 := 0; fi67 < 12; fi67++ {
            for fi68 := 0; fi68 < 7; fi68++ {
                var v69 int32 = -2147483648
                fmt.Println(v69, v56)
                if (1 < int32(v62)) {
                    if true {
                        v70 := GenS7{F0: (v54 | (v54 | v54)), F1: ((v56 != v56) && false)}
                        _ = v70
                    } else {
                        switch v57 {
                        case -370:
                            if (false || (1.5 >= 3.14159)) {
                                v62 = 72
                                switch fi68 {
                                case 478:
                                    var v71 float32 = 0.1
                                    _ = v71
                                case 128:
                                    if ((3.14159 < 0.1) || true) {
                                        var v72 bool = (815 < 65535)
                                        v1 = ""
                                        _ = v72
                                    }
                                case 9223372036854775807:
                                    var v73 int8 = 100
                                    _ = v73
                                default:
                                    switch (-(fi68)) {
                                    case -9223372036854775808:
                                        v74 := v3[(^(v57))]
                                        _ = v74
                                    case 32767:
                                        v57 -= 127
                                    case 670:
                                        var v75 uint = 4294967295
                                        _ = v75
                                    }
                                }
                                fmt.Println(fi67)
                            } else {
                                var v76 int32 = v69
                                if true {
                                    v77 := v3[fi68]
                                    _ = v77
                                } else {
                                    var v78 float64 = v56
                                    v56 = v56
                                    var v79 int64 = 42
                                    _ = v78
                                    _ = v79
                                }
                                _ = v76
                            }
                        case 127:
                            switch 64 {
                            case -140:
                                var v80 float32 = (((3.14159 / (float32((1.5 + 0.1)) - (0.0 * 0.0))) - float32((float64(3.14159) - (float64(float32(v56)) * 1e10)))) * 0.0)
                                if ((fi68 <= fi67) || true) {
                                    if (84 <= 1) {
                                        v54 += v54
                                    } else {
                                        v81 := v3[(v57 - (v57 | (fi68 & 1)))]
                                        _ = v81
                                    }
                                    if true {
                                        if false {
                                            if true {
                                                var v82 uint32 = 792
                                                _ = v82
                                            } else {
                                                switch fi68 {
                                                case 9223372036854775807:
                                                    delete(v3, 2147483647)
                                                    var v83 int8 = -1
                                                    fmt.Println(v83)
                                                    _ = v83
                                                default:
                                                    if ((true && (!false)) || (65535 < 42)) {
                                                        v84 := v3[-2147483648]
                                                        fmt.Println(v1)
                                                        var v85 string = v1
                                                        _ = v84
                                                        _ = v85
                                                    } else {
                                                        if false {
                                                            if (1 != 42) {
                                                                if (v54 < (^(v54))) {
                                                                    if true {
                                                                        if (!false) {
                                                                            fmt.Println(v62, fi68)
                                                                            v86 := GenS4{F0: "foo bar", F1: 0, F2: 69}
                                                                            if (!true) {
                                                                                v1 = "foo bar"
                                                                                if true {
                                                                                    v87 := GenS7{F0: 65535, F1: false}
                                                                                    _ = v87
                                                                                }
                                                                                if (!(!(0 >= 18446744073709551615))) {
                                                                                    v62 = v62
                                                                                    var v88 bool = true
                                                                                    v57 -= (fi68 * fi67)
                                                                                    _ = v88
                                                                                }
                                                                                var v89 int32 = (-(v69))
                                                                                _ = v89
                                                                            }
                                                                            _ = v86
                                                                        } else {
                                                                            v62 = v62
                                                                            v1 = v1
                                                                            v62 = v62
                                                                            if (!true) {
                                                                                fmt.Println(fi68)
                                                                                if len(v2) >= 2 {
                                                                                    v2 = v2[1:len(v2)]
                                                                                }
                                                                                v1 = v1
                                                                                if (!((fi67 % (((fi67 >> ((411) & 63))) | 1)) >= (fi67 ^ (fi68 << ((969) & 63))))) {
                                                                                    v56 += float64(v80)
                                                                                    if len(v2) > 0 {
                                                                                        v2[0] = 838
                                                                                    }
                                                                                } else {
                                                                                    switch ("a" + "hello") {
                                                                                    case "hello":
                                                                                        v54 = (v54 + (v54 << ((1) & 63)))
                                                                                    case "foo bar":
                                                                                        var v90 int64 = 1
                                                                                        _ = v90
                                                                                    default:
                                                                                        if ((-(v69)) >= (v69 | 839)) {
                                                                                            _ = len(v2)
                                                                                            if len(v2) >= 2 {
                                                                                                v2 = v2[1:len(v2)]
                                                                                            }
                                                                                        }
                                                                                        fmt.Println(v54, v54, v69)
                                                                                        fmt.Println(v1, v62, v1)
                                                                                    }
                                                                                    if true {
                                                                                        var v91 string = ""
                                                                                        if (false && true) {
                                                                                            v2 = append(v2, 2147483648)
                                                                                        } else {
                                                                                            fmt.Println(v62, v80, v80)
                                                                                            if true {
                                                                                                fmt.Println(v80)
                                                                                                var v92 int8 = -16
                                                                                                v93 := GenS8{F0: 2, F1: v62}
                                                                                                _ = v92
                                                                                                _ = v93
                                                                                            } else {
                                                                                                if false {
                                                                                                    v57 -= -80
                                                                                                    var v94 uint16 = v54
                                                                                                    v2 = append(v2, 1)
                                                                                                    switch (fi67 - v57) {
                                                                                                    case 127:
                                                                                                        if false {
                                                                                                            fmt.Println(v56)
                                                                                                        } else {
                                                                                                            var v95 int16 = -63
                                                                                                            _ = v95
                                                                                                        }
                                                                                                    case -1:
                                                                                                        _ = len(v2)
                                                                                                    case -918:
                                                                                                        var v96 int32 = (v69 ^ 336)
                                                                                                        _ = v96
                                                                                                    }
                                                                                                    _ = v94
                                                                                                }
                                                                                                if (0 == 4294967295) {
                                                                                                    v91 = v91
                                                                                                } else {
                                                                                                    var v97 uint8 = (v62 ^ genFn1())
                                                                                                    if len(v2) > 0 {
                                                                                                        v2[0] = 0
                                                                                                    }
                                                                                                    _ = v97
                                                                                                }
                                                                                            }
                                                                                            v62 += v62
                                                                                        }
                                                                                        v91 = v91
                                                                                        v98 := GenS6{F0: 127}
                                                                                        switch fi68 {
                                                                                        case 0:
                                                                                            var v99 uint32 = 2147483647
                                                                                            _ = v99
                                                                                        case 128:
                                                                                            var v100 uint8 = (v62 << ((1) & 63))
                                                                                            _ = v100
                                                                                        }
                                                                                        _ = v91
                                                                                        _ = v98
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                        var v101 int8 = -76
                                                                        v3[fi67] = (!false)
                                                                        _ = v101
                                                                    }
                                                                    fmt.Println(v69)
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                    switch (^(fi68)) {
                                                                    case 0:
                                                                        if true {
                                                                            var v102 uint64 = 2
                                                                            if (!((v102 - (v102 * 18446744073709551615)) < v102)) {
                                                                                if false {
                                                                                    switch (v1 + v1) {
                                                                                    case "":
                                                                                        v103 := GenS2{F0: uint16(v102), F1: true, F2: 1000}
                                                                                        _ = v103
                                                                                    }
                                                                                } else {
                                                                                    v57 -= (fi67 ^ (-(fi68)))
                                                                                }
                                                                                v3[4294967295] = true
                                                                                switch "a" {
                                                                                case "\n\t":
                                                                                    if true {
                                                                                        var v104 int = v57
                                                                                        fmt.Println(fi68)
                                                                                        _ = v104
                                                                                    }
                                                                                case "foo bar":
                                                                                    v105 := GenS8{F0: v102, F1: 42}
                                                                                    _ = v105
                                                                                }
                                                                            } else {
                                                                                var v106 float64 = float64(v80)
                                                                                _ = v106
                                                                            }
                                                                            _ = v102
                                                                        } else {
                                                                            v107 := GenS3{F0: 42, F1: 1}
                                                                            var v108 uint32 = 65535
                                                                            v108 -= v108
                                                                            _ = v107
                                                                            _ = v108
                                                                        }
                                                                    default:
                                                                        fmt.Println(v80, fi67)
                                                                    }
                                                                }
                                                                v109 := GenS9{F0: v54, F1: v57}
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                                var v110 int8 = -1
                                                                if false {
                                                                    var v111 uint32 = 1
                                                                    if (641 >= -9223372036854775808) {
                                                                        fmt.Println(v110)
                                                                    } else {
                                                                        switch (v57 & (fi68 * fi67)) {
                                                                        case 128:
                                                                            fmt.Println(v110)
                                                                        }
                                                                        if false {
                                                                            v111 = (v111 >> ((65535) & 63))
                                                                            fmt.Println(v57, v56)
                                                                            var v112 uint16 = (v54 ^ 32767)
                                                                            _ = v112
                                                                        } else {
                                                                            fmt.Println(v80)
                                                                            if true {
                                                                                v62 += (v62 / ((v62) | 1))
                                                                                if false {
                                                                                    v113 := v3[(v57 << ((527) & 63))]
                                                                                    _ = v113
                                                                                } else {
                                                                                    if len(v2) >= 2 {
                                                                                        v2 = v2[1:len(v2)]
                                                                                    }
                                                                                    v114 := GenS9{F0: v54, F1: fi67}
                                                                                    v54 = (v54 + v54)
                                                                                    _ = v114
                                                                                }
                                                                                fmt.Println(v54)
                                                                            } else {
                                                                                v115 := GenS2{F0: v54, F1: (!(v110 != (v110 << ((0) & 63)))), F2: 1000}
                                                                                _ = len(v2)
                                                                                v110 = v110
                                                                                v56 = ((v56 + 3.14159) * float64((1.5 * ((float32(v56) / v80) - 1e10))))
                                                                                fmt.Println(v69)
                                                                                _ = v115
                                                                            }
                                                                            if true {
                                                                                if false {
                                                                                    if true {
                                                                                        v57 -= (fi68 / (((^(v57))) | 1))
                                                                                        var v116 int32 = v69
                                                                                        fmt.Println(v110, v69)
                                                                                        _ = v116
                                                                                    } else {
                                                                                        v3[fi68] = (!(1 != -1))
                                                                                        v117 := GenS10{F0: v54, F1: 255, F2: ((v1 + v1) + v1)}
                                                                                        _ = v117
                                                                                    }
                                                                                    if false {
                                                                                        fmt.Println(v62, v57, v69)
                                                                                        v118 := GenS11{F0: v111}
                                                                                        v56 = (1e10 * 1.5)
                                                                                        _ = v118
                                                                                    } else {
                                                                                        v119 := v3[fi67]
                                                                                        var v120 int64 = 196
                                                                                        _ = v119
                                                                                        _ = v120
                                                                                    }
                                                                                    if false {
                                                                                        delete(v3, 0)
                                                                                        if false {
                                                                                            switch (v57 % (((fi67 / (((-(fi68))) | 1))) | 1)) {
                                                                                            case 128:
                                                                                                v80 = float32(v56)
                                                                                            case 32768:
                                                                                                var v121 uint16 = v54
                                                                                                var v122 string = "foo bar"
                                                                                                _ = v121
                                                                                                _ = v122
                                                                                            case 127:
                                                                                                switch (-(v57)) {
                                                                                                case 128:
                                                                                                    v123 := GenS1{F0: genFn1(), F1: (int8(v69) >= (v110 % (((v110 & (^(v110)))) | 1))), F2: v57}
                                                                                                    _ = v123
                                                                                                case -9223372036854775808:
                                                                                                    v124 := GenS12{F0: (v54 ^ (v54 + v54))}
                                                                                                    _ = v124
                                                                                                default:
                                                                                                    v54 = 448
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    } else {
                                                                                        fmt.Println(v57)
                                                                                        v80 += v80
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                v69 = v69
                                                                            }
                                                                        }
                                                                        _ = len(v2)
                                                                        switch ("" + "") {
                                                                        case "a":
                                                                            switch v1 {
                                                                            case "foo bar":
                                                                                if ((v80 >= v80) && true) {
                                                                                    v111 += (v111 ^ v111)
                                                                                    v111 -= (^(v111))
                                                                                    var v125 uint32 = v111
                                                                                    _ = v125
                                                                                }
                                                                            case "":
                                                                                if (false || true) {
                                                                                    var v126 uint32 = (^(v111))
                                                                                    var v127 int32 = (^(v69))
                                                                                    _ = v126
                                                                                    _ = v127
                                                                                }
                                                                            default:
                                                                                if (0 >= (v110 % ((-19) | 1))) {
                                                                                    var v128 int64 = 989
                                                                                    v129 := GenS4{F0: v1, F1: v110, F2: 0}
                                                                                    v130 := GenS11{F0: 42}
                                                                                    _ = v128
                                                                                    _ = v129
                                                                                    _ = v130
                                                                                } else {
                                                                                    v131 := v3[v57]
                                                                                    if (42 > -595) {
                                                                                        v111 -= 13
                                                                                        v132 := GenS10{F0: (v54 | uint16(v62)), F1: 454, F2: v1}
                                                                                        v3[int(v62)] = true
                                                                                        _ = v132
                                                                                    } else {
                                                                                        var v133 float64 = -2.25
                                                                                        var v134 float32 = (((v80 * -2.25) + (v80 + ((v80 - float32(0.0)) - 3.14159))) + -2.25)
                                                                                        _ = v133
                                                                                        _ = v134
                                                                                    }
                                                                                    fmt.Println(v131, v80, v69)
                                                                                    _ = v131
                                                                                }
                                                                            }
                                                                        default:
                                                                            v57 = int(v111)
                                                                        }
                                                                    }
                                                                    v80 = v80
                                                                    fmt.Println(v1)
                                                                    _ = v111
                                                                }
                                                                _ = v109
                                                                _ = v110
                                                            } else {
                                                                fmt.Println(v54, v1)
                                                            }
                                                        }
                                                        v62 = (v62 + (v62 - 162))
                                                        v135 := GenS13{F0: 79, F1: true, F2: genFn1()}
                                                        var v136 float64 = v56
                                                        if false {
                                                            var v137 int = 127
                                                            var v138 uint8 = 0
                                                            if (640 > 2147483648) {
                                                                fmt.Println(v56, v1)
                                                                if (v62 == (v62 + 40)) {
                                                                    if ((!((255 == 743) || false)) || false) {
                                                                        var v139 int = (fi68 | fi68)
                                                                        _ = v139
                                                                    }
                                                                    v140 := GenS14{F0: 18446744073709551615, F1: 4294967295, F2: -135}
                                                                    if true {
                                                                        if false {
                                                                            switch (fi68 + (^(v137))) {
                                                                            case 1:
                                                                                var v141 uint16 = (v54 | 155)
                                                                                _ = v141
                                                                            case -641:
                                                                                var v142 bool = false
                                                                                _ = v142
                                                                            case -575:
                                                                                v62 += 0
                                                                            }
                                                                            if false {
                                                                                v138 = genFn1()
                                                                                if true {
                                                                                    if false {
                                                                                        v56 = v136
                                                                                        switch fi67 {
                                                                                        case -9223372036854775808:
                                                                                            v143 := GenS14{F0: 18446744073709551615, F1: 0, F2: v69}
                                                                                            _ = v143
                                                                                        default:
                                                                                            var v144 int32 = -42
                                                                                            _ = v144
                                                                                        }
                                                                                        var v145 string = v1
                                                                                        _ = v145
                                                                                    }
                                                                                    v80 = v80
                                                                                }
                                                                            } else {
                                                                                if ((fi68 << ((255) & 63)) < -826) {
                                                                                    _ = len(v3)
                                                                                }
                                                                                switch fi67 {
                                                                                case 4294967295:
                                                                                    v3[-9223372036854775808] = (v62 < genFn1())
                                                                                case 27:
                                                                                    fmt.Println(v138)
                                                                                default:
                                                                                    v69 += v69
                                                                                }
                                                                            }
                                                                            fmt.Println(fi68, v137, v56)
                                                                        }
                                                                        switch (fi67 << ((65535) & 63)) {
                                                                        case 993:
                                                                            if false {
                                                                                v56 += v56
                                                                            }
                                                                        case 42:
                                                                            var v146 uint = 2
                                                                            _ = v146
                                                                        case -374:
                                                                            var v147 int8 = 127
                                                                            _ = v147
                                                                        default:
                                                                            fmt.Println(v69)
                                                                        }
                                                                    } else {
                                                                        var v148 uint64 = 0
                                                                        _ = v148
                                                                    }
                                                                    v69 = (^(v69))
                                                                    var v149 float32 = float32(float64(v80))
                                                                    _ = v140
                                                                    _ = v149
                                                                } else {
                                                                    var v150 int64 = 0
                                                                    _ = v150
                                                                }
                                                            }
                                                            _ = v137
                                                            _ = v138
                                                        } else {
                                                            _ = len(v2)
                                                        }
                                                        _ = v135
                                                        _ = v136
                                                    }
                                                }
                                            }
                                        }
                                        var v151 uint16 = (^(v54))
                                        v152 := GenS15{F0: 4294967295, F1: 1, F2: -983}
                                        _ = v151
                                        _ = v152
                                    } else {
                                        fmt.Println(fi68)
                                    }
                                }
                                _ = v80
                            case 522:
                                fmt.Println(v54, v54)
                            case 42:
                                v2 = append(v2, 684)
                            default:
                                _ = len(v3)
                            }
                        case 9223372036854775807:
                            if (true || (float32(v56) > float32(((float64(-2.25) - v56) + 1.5)))) {
                                var v153 uint16 = (^(v54))
                                var v154 bool = (0.1 > v56)
                                _ = v153
                                _ = v154
                            }
                        }
                    }
                    v54 += (v54 + 388)
                } else {
                    v2 = append(v2, 2)
                    fmt.Println(v54, v54)
                    fmt.Println(v56, v69)
                    if (!(((v1 < v1) && true) && true)) {
                        switch v57 {
                        case 9223372036854775807:
                            var v155 uint16 = v54
                            _ = v155
                        case 932:
                            if false {
                                v156 := GenS16{F0: 4294967295, F1: (0 >= 2147483647)}
                                var v157 int16 = 42
                                _ = v156
                                _ = v157
                            } else {
                                if true {
                                    var v158 int8 = int8(v62)
                                    switch (v1 + ((((((((v1 + ((((((("" + "hello") + (v1 + v1)) + v1) + v1) + ((v1 + v1) + (v1 + v1))) + ("a" + (v1 + "a"))) + "hello")) + v1) + v1) + v1) + "a") + v1) + v1) + "foo bar")) {
                                    case "a":
                                        v69 = (v69 ^ 1)
                                    case "\n\t":
                                        v159 := GenS3{F0: 873, F1: v158}
                                        switch fi68 {
                                        case 32767:
                                            switch fi68 {
                                            case 128:
                                                v160 := v3[(fi67 / (((^(fi67))) | 1))]
                                                _ = v160
                                            default:
                                                fmt.Println(v69, v56)
                                            }
                                        case 9223372036854775807:
                                            var v161 uint = 687
                                            _ = v161
                                        case 42:
                                            var v162 string = ((v1 + (v1 + v1)) + ((v1 + ((v1 + "foo bar") + v1)) + (("" + (v1 + "hello")) + ("" + v1))))
                                            _ = v162
                                        }
                                        _ = v159
                                    default:
                                        var v163 uint32 = 2147483648
                                        _ = v163
                                    }
                                    if true {
                                        v57 += 127
                                        v158 += v158
                                        fmt.Println(v158, v56)
                                    }
                                    _ = v158
                                }
                                if true {
                                    fmt.Println(v54, v62, v69)
                                }
                                if true {
                                    var v164 string = "\n\t"
                                    _ = v164
                                } else {
                                    if (false && (!(!(!false)))) {
                                        v165 := GenS17{F0: 761, F1: 836, F2: v69}
                                        if (127 == 127) {
                                            if (65535 > 65535) {
                                                v69 += v69
                                            }
                                            v54 -= v54
                                            v57 = (-(fi68))
                                            fmt.Println(v69, v1)
                                            var v166 int = v57
                                            _ = v166
                                        }
                                        _ = v165
                                    }
                                    v167 := GenS9{F0: (^(v54)), F1: fi67}
                                    v69 = 0
                                    fmt.Println(fi68, v62, v54)
                                    v2 = append(v2, 2147483647)
                                    v57 = v57
                                    _ = v167
                                }
                            }
                        default:
                            v69 = (v69 & v69)
                        }
                    }
                }
                _ = fi68
                _ = v69
            }
            _ = fi67
        }
    case "foo bar":
        v62 = (v62 ^ (v62 * (^(v62))))
    default:
        v1 = v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v54:", v54)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v54
    _ = v55
    _ = v56
    _ = v57
    _ = v62
}
