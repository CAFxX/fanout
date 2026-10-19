package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 int8; F1 int8; F2 uint32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64; F1 bool; F2 uint64 }
type GenS5 struct { F0 int8; F1 uint64 }
type GenS6 struct { F0 int64; F1 int32 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int8; F1 int64; F2 int16 }
type GenS9 struct { F0 uint; F1 int16; F2 int16 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 uint16; F1 uint16; F2 int32 }
type GenS12 struct { F0 int; F1 int32; F2 uint }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 int }
type GenS16 struct { F0 uint8; F1 bool; F2 int32 }
type GenS17 struct { F0 uint64; F1 bool }
type GenS18 struct { F0 int64; F1 bool }
type GenS19 struct { F0 int64 }
type GenS20 struct { F0 int16 }
type GenS21 struct { F0 int64; F1 int; F2 uint32 }
type GenS22 struct { F0 int32; F1 uint16 }
type GenS23 struct { F0 int16; F1 uint32; F2 bool }
type GenS24 struct { F0 int8 }
type GenS25 struct { F0 bool; F1 uint32; F2 uint8 }
type GenS26 struct { F0 int8 }
type GenS27 struct { F0 int16 }
type GenS28 struct { F0 int8; F1 int32; F2 int64 }
func genFn0(p0 uint8, p1 int64) uint32 {
    v4 := GenS1{F0: 4294967295, F1: -1}
    for fi5 := 0; fi5 < 3; fi5++ {
        fmt.Println(p1, p0, p0)
        p1 += p1
        for fi6 := 0; fi6 < 14; fi6++ {
            var v7 int32 = -885
            if false {
                v8 := []int{fi5, 32768}
                if true {
                    var v9 uint = 42
                    fmt.Println(p1, v7, p0)
                    _ = v9
                } else {
                    v10 := GenS2{F0: 1, F1: 42, F2: 1}
                    _ = v10
                }
                switch fi5 {
                case 4294967295:
                    if (9223372036854775807 < (fi5 << ((743) & 63))) {
                        var v11 uint16 = 352
                        var v12 uint64 = 0
                        _ = v11
                        _ = v12
                    } else {
                        if false {
                            var v13 int64 = p1
                            var v14 int64 = (p1 / (((p1 * p1)) | 1))
                            fmt.Println(v7, fi6, v7)
                            if true {
                                v13 -= -9223372036854775808
                                v15 := GenS1{F0: 65535, F1: 533}
                                _ = v15
                            } else {
                                v14 += 1
                                v16 := GenS2{F0: 0, F1: -1, F2: 0}
                                fmt.Println(v7)
                                v17 := GenS3{F0: 865}
                                v7 = v7
                                _ = v16
                                _ = v17
                            }
                            v7 += v7
                            _ = v13
                            _ = v14
                        }
                    }
                case -794:
                    fmt.Println(p0)
                }
                p1 -= (^(p1))
                _ = v8
            } else {
                switch (fi6 / (((fi6 - (fi5 | 127))) | 1)) {
                case 42:
                    var v18 float32 = 1.5
                    _ = v18
                case -137:
                    v19 := []int{(fi6 - (-(fi6)))}
                    _ = v19
                default:
                    var v20 bool = (!false)
                    var v21 uint32 = 65535
                    _ = v20
                    _ = v21
                }
                if true {
                    var v22 string = ("\n\t" + ("" + "foo bar"))
                    v23 := GenS2{F0: 127, F1: -13, F2: 2147483647}
                    var v24 string = ("" + v22)
                    _ = v22
                    _ = v23
                    _ = v24
                }
                var v25 bool = (0 == 912)
                _ = v25
            }
            var v26 int16 = 42
            var v27 int16 = 147
            _ = fi6
            _ = v26
            _ = v27
            _ = v7
        }
        _ = fi5
    }
    _ = p0
    _ = p1
    _ = v4
    return 424
}
func genFn1(p0 int16) int32 {
    fmt.Println(p0)
    switch "foo bar" {
    case "\n\t":
        if (false && (4294967295 != 1)) {
            var v28 int16 = p0
            p0 = v28
            _ = v28
        }
    case "foo bar":
        fmt.Println(p0)
    case "a":
        v29 := GenS1{F0: 578, F1: 0}
        _ = v29
    default:
        var v30 float64 = 0.0
        _ = v30
    }
    _ = p0
    return 0
}

func main() {
    var v1 string = "hello"
    v2 := make(map[uint32]uint32)
    v3 := []int{-564}
    for fi31 := 0; fi31 < 19; fi31++ {
        switch v1 {
        case "hello":
            var v32 int = fi31
            _ = v32
        default:
            switch (fi31 & fi31) {
            case 2147483647:
                var v33 uint16 = 0
                _ = v33
            case 210:
                if (!false) {
                    for fi34 := 0; fi34 < 3; fi34++ {
                        switch v1 {
                        case "hello":
                            switch ("hello" + v1) {
                            case "\n\t":
                                v1 = v1
                            default:
                                _ = len(v2)
                            }
                        }
                        var v35 float32 = 3.14159
                        if ((^(fi31)) == (fi34 * (fi31 << ((105) & 63)))) {
                            if false {
                                if (false && ((!false) || (!true))) {
                                    if true {
                                        if true {
                                            v35 = (0.0 + 3.14159)
                                            v1 = ("" + v1)
                                            v2[384] = 0
                                        }
                                    } else {
                                        _ = len(v3)
                                        switch (fi31 % ((fi31) | 1)) {
                                        case 128:
                                            v35 = ((0.1 / 1.0) + float32(0.0))
                                        case -2147483648:
                                            delete(v2, 65535)
                                        default:
                                            var v36 string = ((v1 + "hello") + "hello")
                                            _ = v36
                                        }
                                        var v37 float32 = (v35 - v35)
                                        switch fi31 {
                                        case 32768:
                                            var v38 uint8 = 93
                                            _ = v38
                                        case -2147483648:
                                            switch "foo bar" {
                                            case "foo bar":
                                                v37 -= (0.0 + float32(((0.1 + 3.14159) + 0.0)))
                                            }
                                        case -9223372036854775808:
                                            v39 := GenS4{F0: -9223372036854775808, F1: (0 == 42), F2: 42}
                                            _ = v39
                                        default:
                                            var v40 int64 = -1
                                            _ = v40
                                        }
                                        fmt.Println(v1, v35)
                                        _ = v37
                                    }
                                    delete(v2, 839)
                                    v35 -= (0.1 + float32(1.5))
                                } else {
                                    v41 := v2[genFn0(101, 0)]
                                    _ = v41
                                }
                                switch v1 {
                                case "foo bar":
                                    var v42 float64 = float64((v35 * (float32(1.5) + ((0.1 * float32(3.14159)) + (((3.14159 * float32(0.1)) - 0.0) - (float32(0.0) - v35))))))
                                    _ = v42
                                case "hello":
                                    if (genFn0(0, 0) < 42) {
                                        var v43 int64 = -330
                                        v35 = ((((v35 + v35) - (v35 * (((v35 / 1e10) / 3.14159) / float32(((1.5 * -2.25) * -2.25))))) * (((1.5 / float32(1.5)) - v35) / v35)) + v35)
                                        _ = v43
                                    }
                                default:
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                }
                                fmt.Println(fi34)
                            } else {
                                v2[2147483647] = 1
                                v1 = v1
                            }
                        } else {
                            if (fi34 <= -1) {
                                if (-128 >= -128) {
                                    switch fi31 {
                                    case 9223372036854775807:
                                        v1 = "\n\t"
                                    case 2147483647:
                                        v44 := GenS5{F0: 0, F1: 4294967295}
                                        _ = v44
                                    case 42:
                                        var v45 int32 = -2147483648
                                        _ = v45
                                    }
                                    switch "a" {
                                    case "foo bar":
                                        v1 = v1
                                    case "":
                                        fmt.Println(v35, v35, fi34)
                                    case "\n\t":
                                        v46 := GenS6{F0: -699, F1: 2147483647}
                                        _ = v46
                                    }
                                }
                            } else {
                                v47 := GenS7{F0: ("" + "\n\t")}
                                switch fi31 {
                                case -904:
                                    v48 := GenS8{F0: -47, F1: -9223372036854775808, F2: 845}
                                    _ = v48
                                default:
                                    var v49 int8 = 0
                                    _ = v49
                                }
                                fmt.Println(fi34)
                                _ = v47
                            }
                            if false {
                                if (83 != -5) {
                                    var v50 uint64 = 2
                                    _ = len(v2)
                                    var v51 int16 = -109
                                    _ = v50
                                    _ = v51
                                } else {
                                    if false {
                                        if true {
                                            if (-1 >= 100) {
                                                if true {
                                                    v1 = "foo bar"
                                                    var v52 int = (fi34 ^ fi31)
                                                    _ = v52
                                                }
                                                v35 += (1.5 - (v35 + v35))
                                            } else {
                                                switch (fi31 + (fi31 % (((fi31 & (fi31 + (^(fi31))))) | 1))) {
                                                case 32767:
                                                    switch fi34 {
                                                    case 127:
                                                        v2[2147483648] = 2147483648
                                                    default:
                                                        switch fi34 {
                                                        case -1:
                                                            if false {
                                                                v35 -= (3.14159 / -2.25)
                                                                var v53 uint8 = 1
                                                                v54 := GenS9{F0: 42, F1: -2, F2: -32768}
                                                                _ = v53
                                                                _ = v54
                                                            }
                                                            var v55 int64 = -9223372036854775808
                                                            _ = v55
                                                        default:
                                                            if true {
                                                                fmt.Println(fi34)
                                                                _ = len(v3)
                                                            }
                                                        }
                                                    }
                                                case 616:
                                                    if len(v3) > 0 {
                                                        v3[0] = fi34
                                                    }
                                                default:
                                                    switch fi34 {
                                                    case 4294967295:
                                                        v56 := GenS8{F0: 42, F1: 42, F2: -1}
                                                        _ = v56
                                                    case 42:
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    case -2147483648:
                                                        v57 := v2[4294967295]
                                                        var v58 uint = 791
                                                        v57 = v57
                                                        _ = v57
                                                        _ = v58
                                                    default:
                                                        v59 := GenS10{F0: -25}
                                                        var v60 string = v1
                                                        _ = v59
                                                        _ = v60
                                                    }
                                                }
                                                v2[4294967295] = 42
                                                if true {
                                                    fmt.Println(v35, v35, v35)
                                                    fmt.Println(v1, v35, fi34)
                                                    if false {
                                                        if true {
                                                            var v61 uint32 = 422
                                                            v62 := GenS5{F0: 46, F1: 18446744073709551615}
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            _ = v61
                                                            _ = v62
                                                        } else {
                                                            var v63 uint64 = 200
                                                            _ = v63
                                                        }
                                                        if (!(!(false || false))) {
                                                            v64 := GenS11{F0: 730, F1: 1, F2: 1}
                                                            var v65 string = "\n\t"
                                                            if false {
                                                                v65 = "foo bar"
                                                                v66 := GenS11{F0: 255, F1: 255, F2: 925}
                                                                v67 := GenS10{F0: 0}
                                                                if false {
                                                                    var v68 uint16 = 2
                                                                    var v69 uint64 = 1
                                                                    var v70 uint16 = (v68 ^ (v68 << ((42) & 63)))
                                                                    _ = v68
                                                                    _ = v69
                                                                    _ = v70
                                                                }
                                                                _ = v66
                                                                _ = v67
                                                            } else {
                                                                v35 += 3.14159
                                                                if true {
                                                                    fmt.Println(fi31, fi34, fi34)
                                                                } else {
                                                                    _ = len(v2)
                                                                }
                                                            }
                                                            _ = v64
                                                            _ = v65
                                                        }
                                                    }
                                                }
                                            }
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            var v71 uint64 = 0
                                            _ = v71
                                        }
                                        delete(v2, 2)
                                        switch (fi34 % ((fi31) | 1)) {
                                        case 32768:
                                            _ = len(v3)
                                        case 9223372036854775807:
                                            fmt.Println(fi31, v35)
                                            if true {
                                                v72 := GenS4{F0: -504, F1: (-2147483648 <= 1), F2: 2}
                                                if (true && true) {
                                                    _ = len(v3)
                                                    switch v1 {
                                                    case "a":
                                                        v73 := GenS12{F0: fi34, F1: -782, F2: 4294967295}
                                                        _ = v73
                                                    default:
                                                        var v74 uint16 = 771
                                                        _ = v74
                                                    }
                                                    if false {
                                                        v1 = (((v1 + "hello") + v1) + v1)
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    }
                                                }
                                                var v75 uint8 = 255
                                                if true {
                                                    if len(v3) > 0 {
                                                        v3[0] = (fi31 >> ((2) & 63))
                                                    }
                                                    var v76 float64 = 1.5
                                                    var v77 int16 = -1
                                                    _ = v76
                                                    _ = v77
                                                } else {
                                                    v75 = 42
                                                }
                                                _ = v72
                                                _ = v75
                                            } else {
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                if true {
                                                    v3 = append(v3, 2147483647)
                                                    fmt.Println(fi34, fi31)
                                                    fmt.Println(v1, v1, v35)
                                                }
                                                fmt.Println(v35)
                                                var v78 int = -9223372036854775808
                                                _ = v78
                                            }
                                            v2[2147483647] = 65535
                                            v1 = v1
                                        case 127:
                                            v1 = "\n\t"
                                        default:
                                            var v79 int8 = 0
                                            _ = v79
                                        }
                                    }
                                    v80 := GenS3{F0: 0}
                                    _ = v80
                                }
                                v1 = v1
                                v35 = 3.14159
                            }
                            var v81 int16 = 17
                            _ = v81
                        }
                        _ = fi34
                        _ = v35
                    }
                } else {
                    var v82 uint8 = 2
                    fmt.Println(v1, v82)
                    fmt.Println(fi31, v82)
                    _ = v82
                }
            default:
                v83 := GenS13{F0: 89}
                _ = v83
            }
        }
        _ = fi31
    }
    v1 = v1
    v1 = v1
    if (!false) {
        if false {
            var v84 float32 = float32((((float64(float32((0.1 + 0.1))) + float64(1e10)) - 0.0) * 3.14159))
            fmt.Println(v84, v1)
            _ = v84
        }
    } else {
        v85 := GenS14{F0: 1}
        _ = v85
    }
    if false {
        var v86 int8 = 0
        var v87 uint8 = 2
        _ = v86
        _ = v87
    } else {
        v88 := GenS15{F0: -355}
        v1 = (("\n\t" + ("\n\t" + ((v1 + (v1 + v1)) + v1))) + v1)
        if (681 < 0) {
            if false {
                v1 = v1
                v1 = v1
                if (0.1 > (1e10 / 1.0)) {
                    var v89 uint = 1
                    var v90 string = v1
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = v89
                    _ = v90
                } else {
                    switch "\n\t" {
                    case "hello":
                        v1 = v1
                    case "foo bar":
                        v91 := GenS11{F0: 607, F1: 1, F2: 423}
                        _ = v91
                    default:
                        var v92 float64 = 1.5
                        _ = v92
                    }
                    for fi93 := 0; fi93 < 14; fi93++ {
                        for fi94 := 0; fi94 < 10; fi94++ {
                            fmt.Println(v1, v1)
                            _ = fi94
                        }
                        _ = fi93
                    }
                }
            }
            v1 = (v1 + "foo bar")
        }
        _ = v88
    }
    if (!(42 < 42)) {
        fmt.Println(v1)
        v1 = ""
    }
    v95 := GenS14{F0: 65535}
    switch 32768 {
    case -9223372036854775808:
        fmt.Println(v1)
    case 42:
        var v96 uint = 508
        _ = v96
    case 128:
        switch 1 {
        case -829:
            switch 32767 {
            case -868:
                switch 42 {
                case 32767:
                    if false {
                        v1 = v1
                        v1 = "a"
                        v97 := GenS16{F0: 112, F1: false, F2: genFn1(-803)}
                        _ = v97
                    }
                case 1:
                    var v98 float64 = (-2.25 / 0.1)
                    _ = v98
                case 2147483647:
                    if true {
                        var v99 int = -9223372036854775808
                        if ((184 <= 255) && ((!false) && false)) {
                            v100 := GenS17{F0: 18446744073709551615, F1: true}
                            v1 = v1
                            _ = v100
                        }
                        _ = v99
                    } else {
                        for fi101 := 0; fi101 < 13; fi101++ {
                            if ("a" <= "foo bar") {
                                v1 = (v1 + "\n\t")
                                for fi102 := 0; fi102 < 9; fi102++ {
                                    v1 = v1
                                    var v103 bool = false
                                    fmt.Println(v103, v103, v1)
                                    _ = fi102
                                    _ = v103
                                }
                            }
                            _ = fi101
                        }
                        delete(v2, 0)
                        fmt.Println(v1)
                    }
                default:
                    var v104 uint16 = 483
                    _ = v104
                }
            case -2147483648:
                switch 264 {
                case 208:
                    for fi105 := 0; fi105 < 3; fi105++ {
                        var v106 bool = (1.5 >= 1e10)
                        var v107 int = fi105
                        for fi108 := 0; fi108 < 18; fi108++ {
                            var v109 uint16 = 255
                            if false {
                                v109 += (v109 % ((v109) | 1))
                                var v110 uint8 = 127
                                _ = v110
                            } else {
                                v107 -= fi108
                                if (0.0 < 0.1) {
                                    v107 += 128
                                    if len(v3) > 0 {
                                        v3[0] = (v107 / ((fi108) | 1))
                                    }
                                    v111 := GenS18{F0: 0, F1: false}
                                    _ = v111
                                } else {
                                    var v112 uint32 = genFn0(0, 9223372036854775807)
                                    switch ("hello" + v1) {
                                    case "":
                                        if true {
                                            fmt.Println(fi105)
                                        } else {
                                            if (1 < -1) {
                                                _ = len(v2)
                                            }
                                            if v106 {
                                                v113 := GenS7{F0: v1}
                                                switch "" {
                                                case "hello":
                                                    var v114 int32 = 0
                                                    v109 = 284
                                                    v109 += 630
                                                    v114 += (v114 & (v114 + v114))
                                                    v115 := GenS9{F0: 568, F1: 42, F2: -724}
                                                    _ = v114
                                                    _ = v115
                                                default:
                                                    if ((v109 == (v109 - (v109 - (v109 >> ((1) & 63))))) || true) {
                                                        v116 := GenS6{F0: -1, F1: 2147483647}
                                                        _ = len(v2)
                                                        _ = v116
                                                    }
                                                }
                                                var v117 int8 = -102
                                                _ = v113
                                                _ = v117
                                            }
                                        }
                                    case "hello":
                                        v112 -= 752
                                    case "a":
                                        v1 = "a"
                                    default:
                                        switch (v1 + v1) {
                                        case "foo bar":
                                            v2[genFn0(0, -845)] = genFn0(147, -9223372036854775808)
                                        case "\n\t":
                                            if false {
                                                _ = len(v2)
                                            } else {
                                                switch -283 {
                                                case -508:
                                                    var v118 int64 = -9223372036854775808
                                                    _ = v118
                                                case -1:
                                                    switch (v1 + "") {
                                                    case "":
                                                        if ((32767 != 300) || v106) {
                                                            if (150 < 308) {
                                                                v109 = 1
                                                                if v106 {
                                                                    switch (fi108 << ((1) & 63)) {
                                                                    case 42:
                                                                        switch v107 {
                                                                        case 9223372036854775807:
                                                                            var v119 int32 = 0
                                                                            _ = v119
                                                                        case 575:
                                                                            var v120 uint64 = 1
                                                                            _ = v120
                                                                        }
                                                                    case -972:
                                                                        if v106 {
                                                                            _ = len(v2)
                                                                        } else {
                                                                            v1 = (v1 + v1)
                                                                        }
                                                                    case 128:
                                                                        fmt.Println(fi108)
                                                                    default:
                                                                        var v121 float32 = -2.25
                                                                        _ = v121
                                                                    }
                                                                    var v122 uint8 = 128
                                                                    _ = v122
                                                                } else {
                                                                    v123 := GenS19{F0: 0}
                                                                    _ = v123
                                                                }
                                                                var v124 uint16 = (v109 << ((uint(v112)) & 63))
                                                                switch 654 {
                                                                case -9223372036854775808:
                                                                    var v125 uint8 = 1
                                                                    _ = v125
                                                                case 128:
                                                                    var v126 uint = 2
                                                                    _ = v126
                                                                case 2147483647:
                                                                    switch "foo bar" {
                                                                    case "hello":
                                                                        if (1.5 >= -2.25) {
                                                                            switch fi105 {
                                                                            case -2147483648:
                                                                                v107 -= v107
                                                                            default:
                                                                                var v127 int = (fi108 * v107)
                                                                                _ = v127
                                                                            }
                                                                            var v128 uint64 = 797
                                                                            _ = v128
                                                                        }
                                                                    default:
                                                                        v3 = append(v3, fi108)
                                                                        fmt.Println(v112)
                                                                        if false {
                                                                            v3 = append(v3, 32768)
                                                                            v129 := GenS20{F0: -1}
                                                                            if (false && v106) {
                                                                                v112 -= 1
                                                                            } else {
                                                                                switch -544 {
                                                                                case -405:
                                                                                    v130 := GenS21{F0: 0, F1: fi108, F2: v112}
                                                                                    _ = v130
                                                                                default:
                                                                                    var v131 int32 = 0
                                                                                    _ = v131
                                                                                }
                                                                                v132 := GenS18{F0: 9223372036854775807, F1: v106}
                                                                                var v133 string = "a"
                                                                                _ = v132
                                                                                _ = v133
                                                                            }
                                                                            if (42 == -1) {
                                                                                if false {
                                                                                    var v134 float32 = float32(1e10)
                                                                                    var v135 int16 = 453
                                                                                    _ = v134
                                                                                    _ = v135
                                                                                } else {
                                                                                    v124 = (v109 - v109)
                                                                                    v106 = v106
                                                                                }
                                                                                if (1000 != 1000) {
                                                                                    v112 += v112
                                                                                    var v136 int32 = genFn1(113)
                                                                                    _ = v136
                                                                                } else {
                                                                                    if true {
                                                                                        var v137 string = "hello"
                                                                                        var v138 bool = (v112 <= 2147483647)
                                                                                        delete(v2, 1)
                                                                                        _ = v137
                                                                                        _ = v138
                                                                                    }
                                                                                }
                                                                                v107 = 127
                                                                            } else {
                                                                                v106 = (false && v106)
                                                                            }
                                                                            if (-128 < 54) {
                                                                                v109 += (v124 & v124)
                                                                                if (1 != -32768) {
                                                                                    fmt.Println(v109, v112, v109)
                                                                                } else {
                                                                                    var v139 uint16 = v109
                                                                                    var v140 int64 = 549
                                                                                    var v141 int16 = -1
                                                                                    _ = v139
                                                                                    _ = v140
                                                                                    _ = v141
                                                                                }
                                                                                switch 42 {
                                                                                case -506:
                                                                                    v142 := GenS18{F0: 9223372036854775807, F1: (((0.0 / 1e10) > float64(-2.25)) && (v106 && (-198 <= -938)))}
                                                                                    _ = v142
                                                                                default:
                                                                                    v143 := GenS22{F0: genFn1(0), F1: v109}
                                                                                    _ = v143
                                                                                }
                                                                                v112 = 928
                                                                            } else {
                                                                                var v144 uint = 18446744073709551615
                                                                                v145 := GenS2{F0: 0, F1: 42, F2: v112}
                                                                                if len(v3) > 0 {
                                                                                    v3[0] = 128
                                                                                }
                                                                                _ = v144
                                                                                _ = v145
                                                                            }
                                                                            v146 := GenS23{F0: -32768, F1: (^(v112)), F2: v106}
                                                                            if true {
                                                                                v109 -= (v109 - v109)
                                                                                v106 = (v1 == v1)
                                                                                var v147 uint = uint(v124)
                                                                                _ = v147
                                                                            }
                                                                            v148 := GenS7{F0: ("\n\t" + v1)}
                                                                            var v149 float32 = 1e10
                                                                            v150 := GenS24{F0: -1}
                                                                            _ = len(v3)
                                                                            switch (v1 + v1) {
                                                                            case "hello":
                                                                                var v151 uint32 = v112
                                                                                _ = v151
                                                                            }
                                                                            _ = v129
                                                                            _ = v146
                                                                            _ = v148
                                                                            _ = v149
                                                                            _ = v150
                                                                        }
                                                                    }
                                                                }
                                                                _ = v124
                                                            } else {
                                                                var v152 float64 = (0.0 - 0.0)
                                                                switch v1 {
                                                                case "foo bar":
                                                                    var v153 int64 = -9223372036854775808
                                                                    _ = v153
                                                                case "hello":
                                                                    fmt.Println(v1)
                                                                default:
                                                                    var v154 int64 = 172
                                                                    _ = v154
                                                                }
                                                                _ = v152
                                                            }
                                                            v3 = append(v3, (fi105 | (-(fi105))))
                                                            var v155 string = v1
                                                            var v156 float64 = ((1e10 - (float64(0.0) * float64(float32(3.14159)))) / float64((float32(0.1) / 1e10)))
                                                            var v157 int16 = 1000
                                                            switch v155 {
                                                            case "hello":
                                                                v107 = fi105
                                                            case "foo bar":
                                                                var v158 int32 = -2147483648
                                                                _ = v158
                                                            }
                                                            _ = v155
                                                            _ = v156
                                                            _ = v157
                                                        }
                                                    default:
                                                        v159 := GenS9{F0: 255, F1: 214, F2: -32768}
                                                        _ = v159
                                                    }
                                                    v107 = (-(v107))
                                                    if (681 == 276) {
                                                        if v106 {
                                                            var v160 int8 = 127
                                                            var v161 int64 = 328
                                                            _ = v160
                                                            _ = v161
                                                        }
                                                        v107 = fi108
                                                    }
                                                case -932:
                                                    switch (^(fi105)) {
                                                    case -799:
                                                        fmt.Println(v112, v112)
                                                    case -1:
                                                        if false {
                                                            v162 := GenS23{F0: -32768, F1: v112, F2: ((3.14159 - 1e10) >= 0.1)}
                                                            switch v1 {
                                                            case "hello":
                                                                var v163 uint64 = 382
                                                                _ = v163
                                                            case "\n\t":
                                                                var v164 uint = 42
                                                                var v165 int = (^(fi105))
                                                                if (32767 > 1000) {
                                                                    var v166 float32 = ((1e10 * 0.1) * 3.14159)
                                                                    v112 -= v112
                                                                    _ = v166
                                                                }
                                                                if len(v3) > 0 {
                                                                    v3[0] = fi108
                                                                }
                                                                _ = v164
                                                                _ = v165
                                                            }
                                                            v167 := GenS24{F0: int8(v109)}
                                                            v2[(v112 + 148)] = 2147483647
                                                            var v168 int8 = 1
                                                            _ = v162
                                                            _ = v167
                                                            _ = v168
                                                        } else {
                                                            if true {
                                                                switch v1 {
                                                                case "hello":
                                                                    fmt.Println(fi108, fi105, fi105)
                                                                }
                                                                v106 = (v106 || (4294967295 != 42))
                                                                v169 := GenS25{F0: (1 == 0), F1: (v112 & genFn0(1, -1)), F2: 255}
                                                                switch (v1 + (v1 + (v1 + v1))) {
                                                                case "hello":
                                                                    v107 = 0
                                                                case "":
                                                                    if (255 <= 1) {
                                                                        v107 += fi105
                                                                    }
                                                                case "foo bar":
                                                                    v170 := v2[v112]
                                                                    _ = v170
                                                                default:
                                                                    var v171 uint64 = 42
                                                                    _ = v171
                                                                }
                                                                v172 := GenS26{F0: 42}
                                                                _ = v169
                                                                _ = v172
                                                            } else {
                                                                if (0 < -2147483648) {
                                                                    if v106 {
                                                                        var v173 float64 = 3.14159
                                                                        _ = v173
                                                                    }
                                                                    if ((128 <= 1) && v106) {
                                                                        var v174 uint = 0
                                                                        fmt.Println(v112, v109, v107)
                                                                        _ = v174
                                                                    }
                                                                    v175 := GenS27{F0: 1}
                                                                    _ = v175
                                                                }
                                                            }
                                                            _ = len(v3)
                                                            v112 = (v112 + 65535)
                                                        }
                                                    default:
                                                        v176 := GenS10{F0: 1}
                                                        _ = v176
                                                    }
                                                default:
                                                    v1 = (v1 + "foo bar")
                                                }
                                                v177 := GenS1{F0: 2, F1: -137}
                                                switch (v1 + (v1 + "")) {
                                                case "a":
                                                    v1 = "hello"
                                                case "\n\t":
                                                    if true {
                                                        v109 = (^(v109))
                                                        v178 := GenS8{F0: -99, F1: -9223372036854775808, F2: 32767}
                                                        _ = v178
                                                    }
                                                }
                                                _ = v177
                                            }
                                        default:
                                            if v106 {
                                                var v179 uint64 = 0
                                                v109 -= (v109 % ((v109) | 1))
                                                v3 = append(v3, v107)
                                                _ = v179
                                            } else {
                                                v180 := GenS4{F0: 12, F1: (v106 || (0 == 100)), F2: 4294967295}
                                                v3 = append(v3, fi108)
                                                _ = v180
                                            }
                                        }
                                    }
                                    var v181 uint64 = 2
                                    _ = v112
                                    _ = v181
                                }
                                if (false && v106) {
                                    var v182 float64 = -2.25
                                    if (127 <= 42) {
                                        fmt.Println(fi105)
                                        var v183 uint8 = 47
                                        v184 := v2[655]
                                        var v185 uint8 = v183
                                        var v186 int32 = 42
                                        fmt.Println(v107)
                                        if v106 {
                                            v184 = (v184 * (v184 + v184))
                                            var v187 float64 = v182
                                            v183 -= v185
                                            _ = v187
                                        }
                                        _ = v183
                                        _ = v184
                                        _ = v185
                                        _ = v186
                                    }
                                    v107 = -2147483648
                                    _ = v182
                                }
                            }
                            _ = fi108
                            _ = v109
                        }
                        _ = fi105
                        _ = v106
                        _ = v107
                    }
                case -2147483648:
                    var v188 uint64 = 2
                    _ = v188
                default:
                    if true {
                        delete(v2, 2)
                        var v189 int64 = -280
                        _ = v189
                    } else {
                        var v190 int8 = -1
                        _ = v190
                    }
                }
            case -9223372036854775808:
                var v191 uint16 = 42
                _ = v191
            default:
                var v192 int32 = -827
                _ = v192
            }
        case 197:
            var v193 int = -2147483648
            _ = v193
        default:
            fmt.Println(v1)
        }
    default:
        var v194 bool = (32767 >= 42)
        _ = v194
    }
    for fi195 := 0; fi195 < 3; fi195++ {
        switch v1 {
        case "foo bar":
            v1 = (v1 + (("a" + ((("\n\t" + v1) + ("a" + (v1 + v1))) + v1)) + "foo bar"))
        case "":
            var v196 int64 = 607
            _ = v196
        case "\n\t":
            v1 = v1
        }
        if len(v3) > 0 {
            v3[0] = (fi195 + fi195)
        }
        var v197 uint = 97
        _ = fi195
        _ = v197
    }
    for fi198 := 0; fi198 < 8; fi198++ {
        if ((false && (false || true)) && true) {
            _ = len(v3)
        } else {
            if false {
                var v199 string = v1
                _ = v199
            } else {
                fmt.Println(fi198)
                var v200 int = fi198
                fmt.Println(v200)
                _ = v200
            }
        }
        _ = len(v3)
        _ = fi198
    }
    if (42 == 42) {
        fmt.Println(v1)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi201 := 0; fi201 < 14; fi201++ {
            v1 = v1
            if false {
                var v202 string = v1
                fmt.Println(v202, fi201, v1)
                v1 = v1
                _ = v202
            } else {
                v1 = v1
                switch (fi201 << ((0) & 63)) {
                case 4294967295:
                    var v203 uint8 = 2
                    _ = v203
                }
                var v204 float64 = 3.14159
                _ = v204
            }
            _ = fi201
        }
    }
    v205 := GenS28{F0: -104, F1: 2147483647, F2: 0}
    var v206 string = v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v206:", v206)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v205
    _ = v206
    _ = v3
    _ = v95
}
