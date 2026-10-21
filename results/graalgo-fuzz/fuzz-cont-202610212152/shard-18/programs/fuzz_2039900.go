package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 bool; F1 int64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int16; F1 int8 }
type GenS6 struct { F0 int16; F1 int64; F2 uint8 }
type GenS7 struct { F0 int32; F1 uint8 }
func genFn0(p0 uint, p1 int8) int64 {
    var v4 float32 = 0.0
    v5 := GenS1{F0: 42, F1: true}
    v6 := []int{431}
    p1 = 1
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 363
}
func genFn1(p0 int64, p1 uint64) string {
    if false {
        var v7 uint16 = 1
        _ = v7
    }
    switch 42 {
    case 127:
        v8 := []int{32768, 42, 1}
        _ = v8
    case 32768:
        if false {
            v9 := GenS1{F0: 10, F1: true}
            if (((p1 == 1) || (68 <= 255)) && true) {
                switch -247 {
                case 128:
                    v10 := GenS1{F0: 86, F1: true}
                    _ = v10
                case 0:
                    switch "\n\t" {
                    case "hello":
                        var v11 uint = 0
                        _ = v11
                    case "a":
                        var v12 int16 = -943
                        _ = v12
                    case "\n\t":
                        v13 := []int{32768, 713, 0, 32768}
                        _ = v13
                    default:
                        for fi14 := 0; fi14 < 10; fi14++ {
                            switch "a" {
                            case "":
                                v15 := GenS2{F0: ("a" + ""), F1: p1}
                                _ = v15
                            case "a":
                                p0 = p0
                            default:
                                v16 := GenS1{F0: 0, F1: true}
                                _ = v16
                            }
                            p1 = p1
                            _ = fi14
                        }
                    }
                default:
                    switch "hello" {
                    case "":
                        v17 := GenS2{F0: "\n\t", F1: (^(p1))}
                        _ = v17
                    }
                }
                if true {
                    p1 -= p1
                    v18 := []int{}
                    switch "" {
                    case "hello":
                        if true {
                            fmt.Println(p0, p1)
                            p0 += -885
                        }
                    case "a":
                        if ("" != "a") {
                            v19 := GenS2{F0: ("foo bar" + "foo bar"), F1: p1}
                            var v20 bool = (!(1 > 2147483647))
                            _ = v19
                            _ = v20
                        } else {
                            if (1.5 != (1e10 - float32(((3.14159 / ((-2.25 * 1e10) / 0.1)) * 3.14159)))) {
                                switch "a" {
                                case "\n\t":
                                    v21 := GenS2{F0: "\n\t", F1: 295}
                                    _ = v21
                                case "a":
                                    if true {
                                        p1 += p1
                                        var v22 uint64 = (p1 << ((255) & 63))
                                        var v23 float64 = 1e10
                                        _ = v22
                                        _ = v23
                                    }
                                case "":
                                    v24 := GenS2{F0: "a", F1: 193}
                                    _ = v24
                                default:
                                    switch "\n\t" {
                                    case "a":
                                        var v25 uint = uint(p1)
                                        _ = v25
                                    case "hello":
                                        v26 := GenS1{F0: 244, F1: (!false)}
                                        _ = v26
                                    }
                                }
                            }
                        }
                    case "foo bar":
                        v27 := GenS1{F0: 0, F1: (true || false)}
                        _ = v27
                    }
                    _ = v18
                }
            } else {
                if false {
                    v28 := []int{}
                    if (!(!false)) {
                        v28 = append(v28, 32768)
                    } else {
                        if false {
                            v28 = append(v28, 548)
                            for fi29 := 0; fi29 < 11; fi29++ {
                                if (!true) {
                                    _ = len(v28)
                                    fmt.Println(p0, p1, p0)
                                    if true {
                                        if len(v28) > 0 {
                                            v28[0] = fi29
                                        }
                                    } else {
                                        if len(v28) > 0 {
                                            v28[0] = (fi29 & fi29)
                                        }
                                    }
                                }
                                _ = fi29
                            }
                        } else {
                            if len(v28) >= 2 {
                                v28 = v28[1:len(v28)]
                            }
                            switch ((("" + ("a" + "foo bar")) + "foo bar") + ("hello" + ("hello" + "foo bar"))) {
                            case "a":
                                p1 -= 1
                            case "\n\t":
                                if len(v28) >= 2 {
                                    v28 = v28[1:len(v28)]
                                }
                            default:
                                var v30 float64 = (((1e10 + (1e10 + -2.25)) / float64(1e10)) * float64(1e10))
                                _ = v30
                            }
                        }
                    }
                    _ = v28
                } else {
                    var v31 int32 = 0
                    for fi32 := 0; fi32 < 12; fi32++ {
                        var v33 int64 = 9223372036854775807
                        _ = fi32
                        _ = v33
                    }
                    if (0 < 255) {
                        var v34 int16 = 1
                        if true {
                            switch 32768 {
                            case 42:
                                fmt.Println(p1, p0, v34)
                            case -1:
                                fmt.Println(p0, p0, v31)
                            }
                            var v35 int16 = (v34 ^ v34)
                            _ = v35
                        }
                        _ = v34
                    }
                    _ = v31
                }
                v36 := []int{961}
                _ = v36
            }
            fmt.Println(p0)
            _ = v9
        } else {
            if (!(1.5 != 1.5)) {
                p0 = genFn0(2, 42)
            }
            var v37 uint = 4294967295
            _ = v37
        }
    case 9223372036854775807:
        v38 := []int{}
        _ = v38
    default:
        for fi39 := 0; fi39 < 7; fi39++ {
            switch (fi39 / ((fi39) | 1)) {
            case 42:
                var v40 int64 = genFn0(1, 127)
                _ = v40
            case -9223372036854775808:
                v41 := GenS2{F0: (("" + "\n\t") + ("\n\t" + ("\n\t" + (("foo bar" + ("foo bar" + "")) + "foo bar")))), F1: (p1 - p1)}
                _ = v41
            default:
                var v42 uint16 = 2
                _ = v42
            }
            p0 += 178
            _ = fi39
        }
    }
    for fi43 := 0; fi43 < 1; fi43++ {
        p1 -= p1
        fmt.Println(p0, p0)
        for fi44 := 0; fi44 < 8; fi44++ {
            v45 := []int{fi43, (fi43 & (-(fi43)))}
            var v46 uint8 = 2
            p1 = p1
            _ = fi44
            _ = v45
            _ = v46
        }
        _ = fi43
    }
    for fi47 := 0; fi47 < 2; fi47++ {
        if (0 == (p1 >> ((0) & 63))) {
            if (42 != -32768) {
                fmt.Println(p0)
                switch "foo bar" {
                case "":
                    var v48 int16 = -32768
                    _ = v48
                default:
                    var v49 int32 = 482
                    _ = v49
                }
            } else {
                p1 = (p1 ^ p1)
                fmt.Println(p0, p1, fi47)
            }
            fmt.Println(p1)
            switch (fi47 * (fi47 % ((4294967295) | 1))) {
            case 0:
                for fi50 := 0; fi50 < 6; fi50++ {
                    v51 := GenS3{F0: (-1 == fi50), F1: (^(p0))}
                    _ = fi50
                    _ = v51
                }
            case 127:
                v52 := []int{(fi47 - (fi47 << ((255) & 63))), fi47, 2147483647}
                _ = v52
            default:
                var v53 uint32 = 1
                _ = v53
            }
        }
        var v54 int32 = int32(fi47)
        _ = fi47
        _ = v54
    }
    p0 += p0
    _ = p0
    _ = p1
    return ("foo bar" + ("hello" + "hello"))
}

func main() {
    var v1 int8 = 3
    var v2 int16 = 651
    var v3 uint8 = 131
    var v55 float32 = 0.0
    if ((v2 ^ v2) >= v2) {
        switch 369 {
        case 3:
            v56 := []int{32768, 0, 229, -209}
            _ = v56
        default:
            switch "a" {
            case "\n\t":
                var v57 int = -692
                _ = v57
            default:
                v58 := GenS4{F0: 9223372036854775807}
                _ = v58
            }
        }
        if ((v3 >> ((734) & 63)) != 1) {
            if false {
                v55 += v55
            }
            var v59 float64 = (0.1 / 0.1)
            v55 = v55
            _ = v59
        }
        var v60 uint32 = 2147483647
        _ = v60
    } else {
        v2 += (v2 & v2)
        v1 = (v1 - v1)
        for fi61 := 0; fi61 < 2; fi61++ {
            if ((true || (false || true)) && false) {
                var v62 int8 = v1
                _ = v62
            }
            v2 -= 1
            if (541 < -1) {
                var v63 int64 = 9223372036854775807
                var v64 int = (fi61 << ((2) & 63))
                _ = v63
                _ = v64
            }
            _ = fi61
        }
    }
    var v65 uint32 = 65535
    if (!false) {
        for fi66 := 0; fi66 < 14; fi66++ {
            var v67 uint8 = v3
            if (65535 <= 0) {
                for fi68 := 0; fi68 < 2; fi68++ {
                    var v69 bool = ((1 <= uint16(v67)) && (255 < (v67 | v3)))
                    _ = fi68
                    _ = v69
                }
                v70 := GenS5{F0: int16(v3), F1: (v1 - 127)}
                v71 := []int{(-(fi66))}
                _ = v70
                _ = v71
            }
            _ = fi66
            _ = v67
        }
    } else {
        if true {
            v3 = 127
            v3 -= (^(v3))
        }
    }
    switch int(v3) {
    case 779:
        var v72 uint16 = 255
        _ = v72
    case 580:
        for fi73 := 0; fi73 < 13; fi73++ {
            v2 = (v2 - v2)
            v65 = v65
            var v74 int64 = genFn0(11, (v1 ^ (v1 >> ((65535) & 63))))
            _ = fi73
            _ = v74
        }
    default:
        if false {
            for fi75 := 0; fi75 < 10; fi75++ {
                v65 = uint32(v3)
                if (-2.25 >= float64(v55)) {
                    v65 -= v65
                    var v76 int16 = (v2 - 985)
                    _ = v76
                }
                _ = fi75
            }
            v2 = (v2 | (v2 % ((v2) | 1)))
        }
    }
    v77 := GenS5{F0: v2, F1: (v1 >> ((255) & 63))}
    if false {
        for fi78 := 0; fi78 < 18; fi78++ {
            var v79 float32 = (v55 - (v55 - float32((1e10 * (0.1 * 0.1)))))
            if ((v2 + v2) == int16(fi78)) {
                v2 = v2
            } else {
                var v80 bool = false
                fmt.Println(v1)
                v81 := GenS4{F0: 42}
                _ = v80
                _ = v81
            }
            _ = fi78
            _ = v79
        }
        if (0.1 == (1.5 + 0.1)) {
            var v82 int = 9223372036854775807
            if true {
                var v83 float32 = float32(3.14159)
                if (42 != 65535) {
                    if (42 <= uint16(v3)) {
                        if (127 != (v3 / ((v3) | 1))) {
                            v84 := []int{(v82 & 4294967295), (v82 & (v82 ^ int(v2))), 2147483647}
                            v85 := GenS5{F0: (v2 * (v2 - (v2 / (((v2 / (((v2 * -1)) | 1))) | 1)))), F1: v1}
                            _ = v84
                            _ = v85
                        }
                    }
                    if ((v65 / ((42) | 1)) >= (v65 << ((uint(v65)) & 63))) {
                        v86 := []int{(-(v82)), v82}
                        fmt.Println(v3)
                        _ = v86
                    } else {
                        v87 := GenS5{F0: v2, F1: v1}
                        v88 := []int{(v82 ^ 2147483647), v82, (v82 << ((222) & 63))}
                        var v89 uint8 = (v3 * (^(v3)))
                        _ = v87
                        _ = v88
                        _ = v89
                    }
                }
                var v90 uint = 42
                _ = v83
                _ = v90
            } else {
                if true {
                    var v91 bool = true
                    _ = v91
                }
            }
            if false {
                var v92 uint = 704
                var v93 uint = (v92 - v92)
                v1 -= 1
                _ = v92
                _ = v93
            }
            _ = v82
        } else {
            var v94 int16 = v2
            var v95 int8 = (-(v1))
            var v96 string = "foo bar"
            _ = v94
            _ = v95
            _ = v96
        }
        fmt.Println(v65, v3)
    } else {
        if false {
            var v97 float32 = 1.5
            if (1e10 == 0.1) {
                v98 := GenS6{F0: v2, F1: 0, F2: (^(v3))}
                var v99 bool = ((1 >= 0) && false)
                var v100 bool = (0 <= 42)
                _ = v100
                _ = v98
                _ = v99
            }
            _ = v97
        } else {
            v55 = 1e10
            if false {
                var v101 uint32 = v65
                if (false && (200 > 2)) {
                    fmt.Println(v65, v55, v55)
                } else {
                    v65 = v65
                    var v102 uint64 = 2
                    fmt.Println(v3, v101)
                    _ = v102
                }
                v103 := []int{-1, -462, 1, 550}
                _ = v101
                _ = v103
            } else {
                var v104 string = "foo bar"
                v105 := []int{}
                v106 := GenS6{F0: (v2 << ((255) & 63)), F1: 1, F2: uint8(v65)}
                _ = v104
                _ = v105
                _ = v106
            }
        }
        fmt.Println(v65, v55, v1)
    }
    v107 := []int{32767}
    if false {
        switch "\n\t" {
        case "foo bar":
            for fi108 := 0; fi108 < 17; fi108++ {
                if ((2 == 2) && true) {
                    if false {
                        var v109 uint = 1
                        v1 -= 127
                        var v110 uint8 = v3
                        _ = v109
                        _ = v110
                    } else {
                        fmt.Println(v2, v55)
                    }
                    v65 += 2147483648
                }
                if (-32768 != -1) {
                    fmt.Println(v55)
                    var v111 int64 = 42
                    fmt.Println(v2, v1)
                    _ = v111
                } else {
                    fmt.Println(v65)
                    fmt.Println(v55, fi108, v55)
                }
                _ = fi108
            }
        case "a":
            if len(v107) > 0 {
                v107[0] = 32768
            }
        case "\n\t":
            switch "\n\t" {
            case "hello":
                v3 = (v3 - (v3 ^ (v3 - v3)))
            case "":
                if false {
                    if len(v107) >= 2 {
                        v107 = v107[1:len(v107)]
                    }
                    for fi112 := 0; fi112 < 10; fi112++ {
                        switch fi112 {
                        case 128:
                            v3 -= v3
                        case -1:
                            v65 -= v65
                        case 9223372036854775807:
                            v65 = (^(v65))
                        default:
                            for fi113 := 0; fi113 < 2; fi113++ {
                                fmt.Println(v65, fi113, fi113)
                                var v114 string = "\n\t"
                                switch fi112 {
                                case -85:
                                    var v115 float32 = v55
                                    _ = v115
                                case -1:
                                    if (((4294967295 >= 465) || false) && ("\n\t" != "a")) {
                                        var v116 uint8 = v3
                                        _ = v116
                                    } else {
                                        var v117 uint = 521
                                        _ = v117
                                    }
                                default:
                                    if false {
                                        var v118 uint8 = 42
                                        fmt.Println(fi112, fi112, v114)
                                        var v119 float32 = v55
                                        _ = v118
                                        _ = v119
                                    } else {
                                        v107 = append(v107, (-(fi112)))
                                    }
                                }
                                v3 += uint8(v65)
                                _ = fi113
                                _ = v114
                            }
                        }
                        var v120 float64 = 3.14159
                        _ = fi112
                        _ = v120
                    }
                }
            }
        }
    } else {
        for fi121 := 0; fi121 < 17; fi121++ {
            for fi122 := 0; fi122 < 10; fi122++ {
                var v123 uint = 255
                _ = fi122
                _ = v123
            }
            var v124 int8 = 0
            switch (("" + ("" + "\n\t")) + "a") {
            case "\n\t":
                var v125 string = "a"
                _ = v125
            }
            _ = fi121
            _ = v124
        }
        v126 := GenS7{F0: -1, F1: (v3 % ((v3) | 1))}
        v3 = v3
        _ = v126
    }
    for fi127 := 0; fi127 < 16; fi127++ {
        v1 += (v1 | (v1 ^ v1))
        var v128 uint = 65535
        var v129 uint32 = (v65 | v65)
        _ = fi127
        _ = v128
        _ = v129
    }
    switch 1 {
    case 1:
        if len(v107) > 0 {
            v107[0] = 492
        }
    case -9223372036854775808:
        v2 -= -669
    case -653:
        if len(v107) > 0 {
            v107[0] = 4294967295
        }
    default:
        v130 := GenS6{F0: (v2 / ((v2) | 1)), F1: -9223372036854775808, F2: (v3 * v3)}
        _ = v130
    }
    v55 += 3.14159
    v55 -= v55
    var v131 uint32 = v65
    var v132 uint = 263
    var v133 float32 = float32((1.5 * 1e10))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v107:", len(v107), v107)
    fmt.Println("v131:", v131)
    fmt.Println("v132:", v132)
    fmt.Println("v133:", v133)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v55:", v55)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v107
    _ = v131
    _ = v132
    _ = v133
    _ = v2
    _ = v3
    _ = v55
    _ = v65
    _ = v77
}
