package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int; F1 uint32; F2 int64 }
type GenS4 struct { F0 int64; F1 uint64 }
type GenS5 struct { F0 uint64; F1 int64; F2 int64 }
type GenS6 struct { F0 int32; F1 int16; F2 int64 }
type GenS7 struct { F0 uint8; F1 uint16 }
type GenS8 struct { F0 uint; F1 int16 }
type GenS9 struct { F0 uint8; F1 int64; F2 uint32 }
type GenS10 struct { F0 uint; F1 int32; F2 uint16 }
type GenS11 struct { F0 bool; F1 bool }
func genFn0(p0 bool) int64 {
    v4 := []int{32768, 4294967295, 0, 4294967295}
    switch 42 {
    case 42:
        var v5 float32 = (1e10 * ((0.1 - (-2.25 - 0.1)) / (-2.25 - 0.0)))
        _ = v5
    case 2147483647:
        var v6 int64 = 1
        _ = v6
    case 1:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    default:
        _ = len(v4)
    }
    for fi7 := 0; fi7 < 11; fi7++ {
        _ = len(v4)
        fmt.Println(p0, fi7)
        _ = fi7
    }
    _ = p0
    _ = v4
    return -1
}
func genFn1(p0 int16, p1 uint32, p2 uint8) int16 {
    var v8 int = 127
    switch "foo bar" {
    case "hello":
        v9 := []int{(^(v8)), (v8 & 708), v8}
        _ = v9
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return (p0 << ((255) & 63))
}

func main() {
    var v1 uint = 1
    var v2 int = 127
    v3 := make(map[int]bool)
    switch "foo bar" {
    case "\n\t":
        if false {
            fmt.Println(v2, v2)
            var v10 uint = v1
            v11 := GenS1{F0: 827}
            _ = v10
            _ = v11
        }
    case "hello":
        v3[v2] = true
    case "a":
        if ((v1 - v1) <= (v1 << (((v1 / ((v1) | 1))) & 63))) {
            v2 = 649
            var v12 uint = v1
            delete(v3, (v2 - 4294967295))
            _ = v12
        } else {
            fmt.Println(v2, v2)
            if false {
                if false {
                    if true {
                        var v13 int8 = -80
                        if (636 != 2) {
                            v14 := GenS2{F0: 42}
                            _ = v14
                        }
                        if false {
                            var v15 uint64 = 18446744073709551615
                            var v16 uint16 = 65535
                            _ = v15
                            _ = v16
                        }
                        _ = v13
                    }
                    v3[v2] = true
                }
                v17 := GenS2{F0: v1}
                v18 := []int{42, v2, v2, (v2 ^ 361)}
                _ = v17
                _ = v18
            } else {
                fmt.Println(v1, v1)
                var v19 int16 = genFn1(-42, 2147483647, 128)
                var v20 float32 = 3.14159
                _ = v19
                _ = v20
            }
            v3[(v2 ^ 0)] = false
        }
    default:
        fmt.Println(v1, v2)
    }
    var v21 uint32 = 2147483648
    v22 := GenS2{F0: 652}
    var v23 uint32 = (v21 * 620)
    fmt.Println(v1)
    v21 = (v21 * v23)
    if (!false) {
        if false {
            switch "foo bar" {
            case "":
                for fi24 := 0; fi24 < 6; fi24++ {
                    switch "foo bar" {
                    case "\n\t":
                        if false {
                            v25 := v3[-2147483648]
                            _ = v25
                        } else {
                            v2 -= fi24
                            var v26 int8 = 88
                            switch (fi24 & fi24) {
                            case 1:
                                for fi27 := 0; fi27 < 15; fi27++ {
                                    var v28 int32 = -459
                                    v29 := GenS2{F0: (v1 ^ 4294967295)}
                                    _ = fi27
                                    _ = v28
                                    _ = v29
                                }
                            default:
                                v23 = v21
                            }
                            _ = v26
                        }
                    default:
                        v30 := GenS2{F0: v1}
                        _ = v30
                    }
                    _ = fi24
                }
            default:
                for fi31 := 0; fi31 < 13; fi31++ {
                    fmt.Println(fi31)
                    v32 := []int{-2147483648, fi31, fi31}
                    _ = fi31
                    _ = v32
                }
            }
            for fi33 := 0; fi33 < 10; fi33++ {
                fmt.Println(v1, v21)
                _ = fi33
            }
        } else {
            fmt.Println(v1, v21, v1)
            for fi34 := 0; fi34 < 17; fi34++ {
                var v35 uint = 18446744073709551615
                if true {
                    var v36 uint16 = 65535
                    var v37 int16 = genFn1(-32768, 42, 0)
                    _ = v36
                    _ = v37
                } else {
                    switch "a" {
                    case "foo bar":
                        var v38 float32 = 1e10
                        _ = v38
                    case "\n\t":
                        v35 = (v35 - 946)
                    case "hello":
                        v35 -= v1
                    }
                }
                _ = fi34
                _ = v35
            }
            if (-2.25 != 0.0) {
                fmt.Println(v21, v2)
            }
        }
    }
    for fi39 := 0; fi39 < 20; fi39++ {
        v2 -= -1000
        v21 = v23
        v1 = v1
        _ = fi39
    }
    switch v2 {
    case 678:
        v21 = 0
    case -622:
        v40 := GenS3{F0: v2, F1: 108, F2: genFn0((!true))}
        _ = v40
    case 127:
        if true {
            v41 := GenS1{F0: (^(v23))}
            _ = v41
        } else {
            v42 := GenS2{F0: (v1 % ((0) | 1))}
            if true {
                for fi43 := 0; fi43 < 8; fi43++ {
                    switch ("" + "\n\t") {
                    case "\n\t":
                        var v44 int32 = 1
                        _ = v44
                    default:
                        if ((v2 >> ((uint(v21)) & 63)) <= (fi43 ^ (^(fi43)))) {
                            v2 += v2
                        } else {
                            if ((-2.25 + 0.0) > float64((1.5 * ((-2.25 * (1.5 * float32((1.5 * 0.1)))) - (3.14159 + ((-2.25 - 0.1) - 3.14159)))))) {
                                switch 0 {
                                case 127:
                                    switch fi43 {
                                    case 501:
                                        fmt.Println(fi43, v1)
                                    case 9223372036854775807:
                                        if (42 != 42) {
                                            v45 := GenS4{F0: -1, F1: uint64(v1)}
                                            v1 += (v1 % ((v1) | 1))
                                            _ = v45
                                        } else {
                                            fmt.Println(v1, v23)
                                            switch "\n\t" {
                                            case "hello":
                                                switch fi43 {
                                                case 388:
                                                    var v46 bool = ((false && (!false)) && false)
                                                    _ = v46
                                                }
                                            case "a":
                                                v2 = (fi43 >> (((v1 / ((v1) | 1))) & 63))
                                            }
                                            for fi47 := 0; fi47 < 8; fi47++ {
                                                v3[fi47] = (66 == -128)
                                                var v48 int64 = -1
                                                _ = fi47
                                                _ = v48
                                            }
                                        }
                                    case -555:
                                        var v49 int16 = 1000
                                        _ = v49
                                    }
                                case 32767:
                                    var v50 int16 = 32767
                                    _ = v50
                                case 140:
                                    var v51 float32 = (float32((1.5 - 1e10)) / 1e10)
                                    _ = v51
                                default:
                                    var v52 bool = true
                                    _ = v52
                                }
                            } else {
                                v21 += 437
                                if true {
                                    switch v2 {
                                    case -599:
                                        v23 = (v21 - (^(v23)))
                                    case 127:
                                        _ = len(v3)
                                    case 32767:
                                        v3[(fi43 | 32768)] = (fi43 < v2)
                                    }
                                    v21 -= (^(v23))
                                } else {
                                    if true {
                                        v23 -= v21
                                        if (255 >= 32767) {
                                            var v53 uint16 = 51
                                            _ = v53
                                        }
                                        v1 -= v1
                                    }
                                    if true {
                                        switch (fi43 | 32767) {
                                        case -2147483648:
                                            if true {
                                                var v54 uint64 = 42
                                                switch -1 {
                                                case 1:
                                                    var v55 uint = v1
                                                    _ = v55
                                                }
                                                if (65535 <= 754) {
                                                    switch (("\n\t" + "\n\t") + "\n\t") {
                                                    case "":
                                                        if true {
                                                            v1 -= v1
                                                            v56 := GenS2{F0: 4294967295}
                                                            _ = v56
                                                        }
                                                    case "hello":
                                                        v57 := GenS3{F0: -770, F1: v21, F2: 42}
                                                        _ = v57
                                                    }
                                                }
                                                _ = v54
                                            } else {
                                                var v58 string = ""
                                                var v59 int8 = 0
                                                v60 := GenS3{F0: fi43, F1: v23, F2: 360}
                                                _ = v58
                                                _ = v59
                                                _ = v60
                                            }
                                        case -9223372036854775808:
                                            v61 := []int{}
                                            _ = v61
                                        }
                                    } else {
                                        for fi62 := 0; fi62 < 14; fi62++ {
                                            if (v1 >= (v1 << ((18446744073709551615) & 63))) {
                                                var v63 int32 = 0
                                                v64 := GenS5{F0: 0, F1: 0, F2: genFn0((!(true || true)))}
                                                if ("hello" <= (("a" + "hello") + (((("foo bar" + "") + "\n\t") + (("a" + "\n\t") + (("\n\t" + (("foo bar" + "") + "foo bar")) + "\n\t"))) + "a"))) {
                                                    fmt.Println(v63, v23, v1)
                                                    fmt.Println(v21, v23, fi43)
                                                    v63 += v63
                                                    switch "" {
                                                    case "a":
                                                        delete(v3, (fi62 >> ((v1) & 63)))
                                                    case "\n\t":
                                                        v3[(v2 * (-(v2)))] = (2 != 2)
                                                    case "":
                                                        v65 := GenS2{F0: v1}
                                                        _ = v65
                                                    }
                                                }
                                                var v66 int32 = (v63 - v63)
                                                v66 -= 42
                                                v63 -= (v66 >> ((v1) & 63))
                                                _ = v63
                                                _ = v64
                                                _ = v66
                                            }
                                            v67 := GenS6{F0: -2147483648, F1: 32767, F2: 9223372036854775807}
                                            fmt.Println(fi62, v21)
                                            _ = fi62
                                            _ = v67
                                        }
                                        switch "" {
                                        case "":
                                            v68 := GenS4{F0: 0, F1: 42}
                                            _ = v68
                                        }
                                        var v69 int32 = 1
                                        _ = v69
                                    }
                                    fmt.Println(v1, v21, v2)
                                }
                            }
                            v23 = v21
                        }
                    }
                    _ = fi43
                }
                for fi70 := 0; fi70 < 19; fi70++ {
                    if (0 <= -603) {
                        for fi71 := 0; fi71 < 5; fi71++ {
                            v2 += fi71
                            switch "\n\t" {
                            case "foo bar":
                                if (255 != 32767) {
                                    v72 := GenS2{F0: 17}
                                    var v73 uint32 = v23
                                    v1 -= 2
                                    fmt.Println(v21)
                                    v21 -= v23
                                    _ = v72
                                    _ = v73
                                } else {
                                    var v74 int = (v2 + 0)
                                    _ = v74
                                }
                            }
                            fmt.Println(v1, fi71)
                            _ = fi71
                        }
                        if (-32768 < 42) {
                            var v75 uint16 = uint16(v1)
                            for fi76 := 0; fi76 < 4; fi76++ {
                                fmt.Println(v21, v2)
                                v21 += v23
                                v77 := GenS7{F0: 0, F1: (^(v75))}
                                _ = fi76
                                _ = v77
                            }
                            _ = v75
                        }
                    } else {
                        v78 := v3[(^(fi70))]
                        _ = v78
                    }
                    if (128 <= 0) {
                        for fi79 := 0; fi79 < 12; fi79++ {
                            v80 := GenS4{F0: 42, F1: 4294967295}
                            var v81 float64 = (((1e10 / 1.5) + ((3.14159 - ((((3.14159 - float64(float32(float64(0.0)))) - ((3.14159 * 1e10) / 1.0)) * 1e10) - (1e10 - (float64(0.1) - 1.5)))) * (1.5 * -2.25))) + 1e10)
                            switch 127 {
                            case -1:
                                v82 := []int{(fi70 & (-(fi79))), v2, (v2 | fi70), fi79}
                                _ = v82
                            case 0:
                                v21 -= (v23 - (v21 / (((^(v21))) | 1)))
                            case -460:
                                v83 := []int{fi70, 0, v2}
                                _ = v83
                            default:
                                if true {
                                    if (uint32(v1) != (v23 * uint32(v1))) {
                                        var v84 float32 = (1.5 + 0.0)
                                        var v85 bool = true
                                        _ = v84
                                        _ = v85
                                    }
                                    var v86 uint16 = 680
                                    _ = v86
                                } else {
                                    v87 := GenS2{F0: (v1 / ((v1) | 1))}
                                    _ = v87
                                }
                                fmt.Println(v1, v2)
                                var v88 int64 = 9223372036854775807
                                _ = v88
                            }
                            var v89 uint8 = 128
                            v90 := GenS6{F0: 299, F1: -227, F2: -296}
                            _ = fi79
                            _ = v80
                            _ = v81
                            _ = v89
                            _ = v90
                        }
                        v91 := []int{}
                        v21 += 4294967295
                        _ = v91
                    }
                    v92 := []int{fi70}
                    _ = fi70
                    _ = v92
                }
            } else {
                fmt.Println(v23, v2)
                if (-878 < -337) {
                    if (311 < 65535) {
                        var v93 uint64 = 42
                        _ = v93
                    } else {
                        if (2147483647 > -1) {
                            fmt.Println(v2, v23)
                            switch "\n\t" {
                            case "\n\t":
                                var v94 int32 = 42
                                _ = v94
                            case "a":
                                for fi95 := 0; fi95 < 10; fi95++ {
                                    if (false || true) {
                                        v2 = fi95
                                        v23 = 42
                                    }
                                    v96 := GenS3{F0: fi95, F1: v23, F2: 9223372036854775807}
                                    _ = fi95
                                    _ = v96
                                }
                            case "":
                                for fi97 := 0; fi97 < 4; fi97++ {
                                    v1 -= (v1 << ((v1) & 63))
                                    for fi98 := 0; fi98 < 7; fi98++ {
                                        switch "a" {
                                        case "a":
                                            var v99 int64 = 42
                                            _ = v99
                                        case "":
                                            v100 := v3[(v2 >> ((523) & 63))]
                                            _ = v100
                                        case "\n\t":
                                            v101 := []int{fi97, 32767, (fi97 % (((fi97 / (((fi97 % (((fi97 * fi98)) | 1))) | 1))) | 1))}
                                            _ = v101
                                        }
                                        _ = fi98
                                    }
                                    _ = fi97
                                }
                            default:
                                for fi102 := 0; fi102 < 6; fi102++ {
                                    v103 := []int{(fi102 >> ((uint(v21)) & 63)), (fi102 / ((833) | 1)), v2, (^(fi102))}
                                    _ = fi102
                                    _ = v103
                                }
                            }
                            if false {
                                for fi104 := 0; fi104 < 15; fi104++ {
                                    for fi105 := 0; fi105 < 12; fi105++ {
                                        v106 := GenS8{F0: v1, F1: 605}
                                        _ = fi105
                                        _ = v106
                                    }
                                    var v107 int16 = genFn1(genFn1(-860, v21, 184), v21, 1)
                                    v107 = v107
                                    _ = fi104
                                    _ = v107
                                }
                                var v108 int16 = -76
                                _ = v108
                            }
                        }
                        var v109 string = "foo bar"
                        _ = v109
                    }
                    v23 += 307
                    v110 := GenS7{F0: 42, F1: uint16(v23)}
                    _ = v110
                } else {
                    switch v2 {
                    case 128:
                        if (1 != (v2 - 911)) {
                            v111 := GenS7{F0: 0, F1: 28}
                            fmt.Println(v1, v21, v21)
                            _ = v111
                        } else {
                            var v112 int8 = 1
                            v113 := GenS9{F0: 255, F1: 0, F2: 382}
                            switch "" {
                            case "a":
                                switch v2 {
                                case 2147483647:
                                    v114 := GenS4{F0: genFn0(false), F1: 18446744073709551615}
                                    _ = v114
                                default:
                                    fmt.Println(v112, v23)
                                }
                            }
                            _ = v112
                            _ = v113
                        }
                    default:
                        v1 += 0
                    }
                    var v115 int16 = 1000
                    v116 := GenS1{F0: v21}
                    _ = v115
                    _ = v116
                }
            }
            _ = v42
        }
    default:
        fmt.Println(v21)
    }
    var v117 uint64 = 2
    v118 := GenS10{F0: v1, F1: 0, F2: 255}
    for fi119 := 0; fi119 < 16; fi119++ {
        v120 := []int{(v2 ^ 1), fi119, v2, -1}
        var v121 uint32 = v23
        _ = fi119
        _ = v120
        _ = v121
    }
    switch (v2 ^ v2) {
    case 2147483647:
        v122 := []int{v2, (^(v2)), (v2 / (((v2 + (-(v2)))) | 1))}
        _ = v122
    }
    if true {
        for fi123 := 0; fi123 < 15; fi123++ {
            var v124 uint16 = 2
            for fi125 := 0; fi125 < 9; fi125++ {
                v1 = (^(v1))
                var v126 uint64 = (v117 / (((v117 % ((1) | 1))) | 1))
                _ = fi125
                _ = v126
            }
            var v127 float64 = 0.0
            _ = fi123
            _ = v124
            _ = v127
        }
        for fi128 := 0; fi128 < 11; fi128++ {
            for fi129 := 0; fi129 < 16; fi129++ {
                var v130 int64 = int64(v1)
                _ = fi129
                _ = v130
            }
            _ = fi128
        }
    } else {
        v131 := GenS11{F0: true, F1: true}
        _ = v131
    }
    switch (v2 - v2) {
    case 1:
        switch v2 {
        case 9223372036854775807:
            fmt.Println(v21, v2, v117)
        case 32767:
            var v132 uint8 = 2
            _ = v132
        }
    }
    for fi133 := 0; fi133 < 18; fi133++ {
        for fi134 := 0; fi134 < 6; fi134++ {
            v135 := GenS2{F0: v1}
            _ = fi134
            _ = v135
        }
        var v136 int16 = 0
        for fi137 := 0; fi137 < 15; fi137++ {
            v2 -= v2
            _ = fi137
        }
        _ = fi133
        _ = v136
    }
    fmt.Println("v1:", v1)
    fmt.Println("v117:", v117)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v117
    _ = v118
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
