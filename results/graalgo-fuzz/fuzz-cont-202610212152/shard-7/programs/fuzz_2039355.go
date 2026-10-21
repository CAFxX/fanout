package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 int16 }
type GenS2 struct { F0 int8; F1 uint64; F2 int16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int32; F1 int8 }
type GenS6 struct { F0 int16; F1 int8; F2 uint8 }
type GenS7 struct { F0 uint64; F1 string }
type GenS8 struct { F0 bool; F1 uint16; F2 uint16 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 int64; F1 uint64; F2 bool }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 string }
type GenS14 struct { F0 int32; F1 uint8; F2 uint16 }
type GenS15 struct { F0 uint64; F1 int32 }
type GenS16 struct { F0 string; F1 int8; F2 uint64 }
type GenS17 struct { F0 bool; F1 uint64; F2 uint16 }
type GenS18 struct { F0 int16; F1 string }
type GenS19 struct { F0 string; F1 string; F2 int8 }
func genFn0(p0 int16, p1 int8, p2 int64) uint8 {
    if true {
        v7 := GenS1{F0: 0, F1: -2147483648, F2: p0}
        for fi8 := 0; fi8 < 2; fi8++ {
            p0 -= (p0 << ((2) & 63))
            _ = fi8
        }
        _ = v7
    }
    for fi9 := 0; fi9 < 1; fi9++ {
        if (!(255 > 255)) {
            fmt.Println(p2, p2)
            switch (^(fi9)) {
            case 1:
                v10 := GenS1{F0: 42, F1: (fi9 - fi9), F2: -406}
                _ = v10
            case 2147483647:
                v11 := GenS2{F0: p1, F1: 2, F2: p0}
                _ = v11
            case -211:
                for fi12 := 0; fi12 < 19; fi12++ {
                    v13 := GenS3{F0: 887}
                    if true {
                        v14 := GenS4{F0: 1}
                        p2 += -9223372036854775808
                        _ = v14
                    } else {
                        v15 := []int{fi12, (fi9 + fi12), 0, (-(fi12))}
                        p2 = p2
                        if (!true) {
                            v15 = append(v15, -1)
                        } else {
                            var v16 uint = 18446744073709551615
                            if len(v15) > 0 {
                                v15[0] = int(v16)
                            }
                            p0 -= (p0 - p0)
                            _ = v16
                        }
                        _ = v15
                    }
                    _ = fi12
                    _ = v13
                }
            default:
                p0 -= p0
            }
            p0 = p0
        }
        _ = fi9
    }
    p1 -= (p1 ^ 0)
    fmt.Println(p1)
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    return 81
}
func genFn1() uint {
    var v17 uint16 = 42
    v18 := GenS2{F0: 127, F1: 0, F2: 340}
    fmt.Println(v17)
    for fi19 := 0; fi19 < 1; fi19++ {
        var v20 float32 = (0.0 / 1.0)
        _ = fi19
        _ = v20
    }
    _ = v17
    _ = v18
    return 65535
}

func main() {
    v1 := make(map[int]uint32)
    var v2 string = "foo bar"
    v3 := []string{}
    var v4 float32 = 3.14159
    v5 := make(map[int]bool)
    var v6 uint16 = 42
    for fi21 := 0; fi21 < 17; fi21++ {
        var v22 uint64 = 871
        _ = fi21
        _ = v22
    }
    if true {
        fmt.Println(v6, v4, v4)
        var v23 uint16 = (^(v6))
        _ = v23
    } else {
        if true {
            if (1000 == -32768) {
                var v24 uint64 = 381
                switch 32767 {
                case 437:
                    var v25 uint64 = v24
                    _ = v25
                }
                var v26 uint64 = 1
                _ = v24
                _ = v26
            }
        } else {
            var v27 int8 = 0
            _ = v27
        }
        _ = len(v5)
        if true {
            switch v2 {
            case "\n\t":
                var v28 uint64 = 42
                _ = v28
            case "a":
                for fi29 := 0; fi29 < 12; fi29++ {
                    if false {
                        var v30 uint = 34
                        v4 = ((((((3.14159 / -2.25) * -2.25) - v4) / (0.0 * v4)) + float32(float64(1.5))) - 1e10)
                        _ = v30
                    }
                    switch (fi29 * (fi29 % ((fi29) | 1))) {
                    case 32768:
                        for fi31 := 0; fi31 < 1; fi31++ {
                            var v32 int = -284
                            _ = fi31
                            _ = v32
                        }
                    case 9223372036854775807:
                        if (!(0 != 95)) {
                            v6 = (v6 + 42)
                        }
                    case -9223372036854775808:
                        if (188 != 805) {
                            fmt.Println(v4, v2)
                            v2 = "hello"
                        } else {
                            var v33 string = v2
                            v4 = (float32(0.1) * 0.0)
                            v2 = v33
                            _ = v33
                        }
                    }
                    _ = fi29
                }
            case "foo bar":
                var v34 uint = 893
                _ = v34
            default:
                v6 = v6
            }
            for fi35 := 0; fi35 < 2; fi35++ {
                if ((v6 + 42) == (v6 + 817)) {
                    if false {
                        v3 = append(v3, v2)
                        if false {
                            for fi36 := 0; fi36 < 16; fi36++ {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                var v37 int64 = 9223372036854775807
                                _ = fi36
                                _ = v37
                            }
                            v38 := GenS5{F0: 1, F1: 1}
                            fmt.Println(v6)
                            _ = v38
                        }
                    } else {
                        if ((255 != 2) || false) {
                            for fi39 := 0; fi39 < 7; fi39++ {
                                switch v2 {
                                case "":
                                    var v40 float32 = (v4 + v4)
                                    _ = v40
                                case "foo bar":
                                    var v41 uint8 = 0
                                    _ = v41
                                }
                                v5[fi39] = (!((3.14159 != (1e10 + 3.14159)) && ((fi39 > (fi39 / ((int(v6)) | 1))) || true)))
                                v3 = append(v3, (v2 + v2))
                                _ = fi39
                            }
                            fmt.Println(v2, v6)
                            for fi42 := 0; fi42 < 12; fi42++ {
                                v43 := GenS5{F0: 0, F1: -128}
                                if (-438 < 32767) {
                                    switch (fi35 >> ((uint(v6)) & 63)) {
                                    case 305:
                                        var v44 uint64 = 0
                                        _ = v44
                                    case 4294967295:
                                        switch fi35 {
                                        case 0:
                                            if (42 <= 2) {
                                                if true {
                                                    v2 = (v2 + (v2 + v2))
                                                }
                                            } else {
                                                var v45 float64 = (1.5 + 0.0)
                                                switch (((v2 + v2) + (v2 + "\n\t")) + (v2 + v2)) {
                                                case "":
                                                    if true {
                                                        if (!true) {
                                                            switch fi42 {
                                                            case 1:
                                                                if (fi42 >= fi35) {
                                                                    var v46 uint8 = 128
                                                                    _ = v46
                                                                }
                                                            }
                                                            fmt.Println(fi42)
                                                            if (!(0 != -9223372036854775808)) {
                                                                if false {
                                                                    var v47 uint = 18446744073709551615
                                                                    switch (fi42 << (((^(v47))) & 63)) {
                                                                    case 4294967295:
                                                                        var v48 float32 = (((v4 + float32(float64(v4))) + (float32(v45) / 1.0)) / ((3.14159 - v4) - v4))
                                                                        _ = v48
                                                                    case 9223372036854775807:
                                                                        switch "\n\t" {
                                                                        case "a":
                                                                            if false {
                                                                                v4 -= (0.0 - v4)
                                                                            } else {
                                                                                var v49 uint = (v47 % ((v47) | 1))
                                                                                _ = v49
                                                                            }
                                                                        }
                                                                    }
                                                                    _ = v47
                                                                } else {
                                                                    v6 = v6
                                                                }
                                                                if (42 <= 232) {
                                                                    v50 := GenS1{F0: 2, F1: 2147483647, F2: 789}
                                                                    switch v2 {
                                                                    case "a":
                                                                        var v51 float32 = v4
                                                                        _ = v51
                                                                    case "\n\t":
                                                                        var v52 uint8 = 42
                                                                        _ = v52
                                                                    default:
                                                                        var v53 int = 9223372036854775807
                                                                        _ = v53
                                                                    }
                                                                    var v54 uint = 255
                                                                    _ = v50
                                                                    _ = v54
                                                                } else {
                                                                    if (861 > -319) {
                                                                        switch v2 {
                                                                        case "":
                                                                            if true {
                                                                                var v55 int8 = 127
                                                                                switch v2 {
                                                                                case "hello":
                                                                                    v2 = v2
                                                                                default:
                                                                                    fmt.Println(v2, v2, v2)
                                                                                }
                                                                                switch (^(fi35)) {
                                                                                case -9223372036854775808:
                                                                                    if false {
                                                                                        fmt.Println(fi35, v55)
                                                                                        if false {
                                                                                            v56 := GenS6{F0: int16(v55), F1: v55, F2: 2}
                                                                                            if len(v3) >= 2 {
                                                                                                v3 = v3[1:len(v3)]
                                                                                            }
                                                                                            var v57 int16 = 0
                                                                                            _ = v56
                                                                                            _ = v57
                                                                                        }
                                                                                    }
                                                                                    v58 := GenS3{F0: fi42}
                                                                                    switch 2147483647 {
                                                                                    case 4294967295:
                                                                                        v2 = v2
                                                                                    default:
                                                                                        v2 = "foo bar"
                                                                                    }
                                                                                    _ = v58
                                                                                case -827:
                                                                                    var v59 uint64 = 1
                                                                                    _ = v59
                                                                                }
                                                                                _ = v55
                                                                            }
                                                                        case "\n\t":
                                                                            if true {
                                                                                if false {
                                                                                    delete(v5, fi42)
                                                                                    var v60 bool = (v45 > 3.14159)
                                                                                    _ = v60
                                                                                }
                                                                                v61 := v5[128]
                                                                                fmt.Println(v4)
                                                                                v45 = v45
                                                                                var v62 int32 = -1
                                                                                _ = v61
                                                                                _ = v62
                                                                            } else {
                                                                                v63 := v5[127]
                                                                                var v64 int = (fi42 - int(v6))
                                                                                v65 := GenS2{F0: 0, F1: 18446744073709551615, F2: -740}
                                                                                _ = v63
                                                                                _ = v64
                                                                                _ = v65
                                                                            }
                                                                        default:
                                                                            var v66 uint8 = 255
                                                                            _ = v66
                                                                        }
                                                                        v67 := GenS7{F0: 1, F1: ""}
                                                                        fmt.Println(v4, fi42, v45)
                                                                        _ = v67
                                                                    }
                                                                }
                                                                switch v2 {
                                                                case "hello":
                                                                    var v68 int32 = -2147483648
                                                                    _ = v68
                                                                case "\n\t":
                                                                    if (v6 < v6) {
                                                                        fmt.Println(v4, v2, v45)
                                                                        var v69 int16 = -249
                                                                        _ = v69
                                                                    } else {
                                                                        var v70 int = fi42
                                                                        v71 := GenS3{F0: v70}
                                                                        _ = v70
                                                                        _ = v71
                                                                    }
                                                                default:
                                                                    fmt.Println(v4, v2, fi42)
                                                                }
                                                            }
                                                        }
                                                    } else {
                                                        switch (fi35 & 834) {
                                                        case 127:
                                                            if (65535 <= v6) {
                                                                v2 = v2
                                                                fmt.Println(fi42, v6)
                                                            } else {
                                                                fmt.Println(v45, v2, fi35)
                                                                v6 = 42
                                                                delete(v5, 42)
                                                            }
                                                        case -98:
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            v6 = (v6 - v6)
                                                            var v72 int = (fi35 * (fi35 & fi35))
                                                            _ = v72
                                                        case 2147483647:
                                                            v73 := GenS6{F0: 32767, F1: 39, F2: 39}
                                                            _ = v73
                                                        default:
                                                            var v74 int64 = -104
                                                            _ = v74
                                                        }
                                                        v75 := GenS8{F0: (1000 <= 42), F1: (v6 | (v6 - (v6 - v6))), F2: (v6 + v6)}
                                                        var v76 string = v2
                                                        _ = v75
                                                        _ = v76
                                                    }
                                                case "\n\t":
                                                    var v77 uint64 = 682
                                                    _ = v77
                                                case "foo bar":
                                                    var v78 int32 = 338
                                                    _ = v78
                                                }
                                                var v79 int32 = 2147483647
                                                _ = v45
                                                _ = v79
                                            }
                                        case 963:
                                            var v80 bool = false
                                            _ = v80
                                        }
                                    case 9223372036854775807:
                                        delete(v1, (fi35 << ((381) & 63)))
                                    default:
                                        v2 = ("hello" + (v2 + (v2 + v2)))
                                    }
                                    fmt.Println(fi35, v2)
                                } else {
                                    if true {
                                        v4 -= v4
                                        v4 -= -2.25
                                        v81 := GenS5{F0: -2147483648, F1: 98}
                                        switch 0 {
                                        case 141:
                                            _ = len(v3)
                                        }
                                        _ = v81
                                    }
                                    var v82 uint8 = 127
                                    _ = v82
                                }
                                var v83 int64 = 1
                                _ = fi42
                                _ = v43
                                _ = v83
                            }
                        } else {
                            var v84 uint8 = 127
                            for fi85 := 0; fi85 < 20; fi85++ {
                                switch (v2 + v2) {
                                case "a":
                                    delete(v1, fi85)
                                case "\n\t":
                                    switch fi35 {
                                    case -2147483648:
                                        if true {
                                            var v86 string = v2
                                            _ = v86
                                        } else {
                                            v87 := v1[(fi85 & fi35)]
                                            v1[(-(fi85))] = (v87 ^ 76)
                                            _ = v87
                                        }
                                        v6 += (v6 ^ (v6 + v6))
                                    case 369:
                                        v84 = genFn0(1000, int8(v84), 42)
                                    default:
                                        if (!true) {
                                            v84 += uint8(v6)
                                            v4 += v4
                                            if len(v3) > 0 {
                                                v3[0] = v2
                                            }
                                        } else {
                                            switch (fi85 & (fi35 << ((4294967295) & 63))) {
                                            case -1:
                                                switch "foo bar" {
                                                case "\n\t":
                                                    v6 = v6
                                                    if (558 != 0) {
                                                        var v88 bool = ((true || (42 != uint(v84))) || false)
                                                        switch "hello" {
                                                        case "hello":
                                                            v6 = 1
                                                        }
                                                        _ = v88
                                                    } else {
                                                        var v89 int64 = 0
                                                        _ = v89
                                                    }
                                                    var v90 int8 = 127
                                                    _ = v90
                                                case "hello":
                                                    v6 = 65535
                                                case "":
                                                    v91 := GenS9{F0: fi35}
                                                    _ = v91
                                                }
                                            case 32767:
                                                switch fi85 {
                                                case 4294967295:
                                                    if false {
                                                        v92 := GenS1{F0: (^(v84)), F1: -260, F2: 0}
                                                        fmt.Println(v2)
                                                        _ = v92
                                                    }
                                                case 419:
                                                    v2 = ("foo bar" + (v2 + v2))
                                                    if (4294967295 > 18446744073709551615) {
                                                        v93 := GenS8{F0: (212 <= 42), F1: 1, F2: 255}
                                                        _ = len(v3)
                                                        _ = v93
                                                    }
                                                case 2147483647:
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                default:
                                                    v4 += ((float32(1e10) - v4) - 1.5)
                                                }
                                            case -872:
                                                _ = len(v3)
                                            default:
                                                v94 := GenS5{F0: 1, F1: -1}
                                                _ = v94
                                            }
                                        }
                                    }
                                    if true {
                                        var v95 bool = ((((((1.5 + (-2.25 - 0.0)) * -2.25) / 1e10) - 3.14159) / (-2.25 / 1.0)) <= 3.14159)
                                        _ = v95
                                    } else {
                                        switch "hello" {
                                        case "":
                                            fmt.Println(v4)
                                        case "\n\t":
                                            switch fi85 {
                                            case 128:
                                                var v96 uint64 = 0
                                                _ = v96
                                            }
                                        case "foo bar":
                                            var v97 int32 = 1
                                            _ = v97
                                        default:
                                            var v98 bool = (42 <= 28)
                                            _ = v98
                                        }
                                        if true {
                                            if false {
                                                v84 = 255
                                                v6 += v6
                                                if false {
                                                    v6 = (v6 + v6)
                                                    v4 -= (((0.0 - v4) - (v4 + v4)) / ((v4 / v4) - v4))
                                                    var v99 uint64 = 668
                                                    _ = v99
                                                } else {
                                                    fmt.Println(v6, fi35, v2)
                                                    var v100 string = v2
                                                    _ = v100
                                                }
                                            } else {
                                                fmt.Println(v2, v2, fi85)
                                            }
                                            _ = len(v3)
                                        } else {
                                            v101 := GenS10{F0: 0}
                                            if (2147483648 == 4294967295) {
                                                var v102 int64 = -1
                                                if true {
                                                    v2 = ("\n\t" + v2)
                                                    fmt.Println(fi85)
                                                    var v103 int64 = v102
                                                    _ = v103
                                                }
                                                v6 = 32767
                                                _ = v102
                                            } else {
                                                v2 = "\n\t"
                                            }
                                            v2 = ("\n\t" + v2)
                                            _ = v101
                                        }
                                        switch (fi85 / (((^(fi85))) | 1)) {
                                        case 32768:
                                            fmt.Println(v4, v6, v2)
                                        case 42:
                                            var v104 int = 112
                                            _ = v104
                                        case 32767:
                                            var v105 uint = 65
                                            _ = v105
                                        }
                                        if ((!false) || ((true && false) && false)) {
                                            _ = len(v3)
                                            v84 -= 79
                                        } else {
                                            fmt.Println(v84)
                                        }
                                    }
                                    v106 := GenS5{F0: 175, F1: -1}
                                    if len(v3) > 0 {
                                        v3[0] = (v2 + v2)
                                    }
                                    _ = v106
                                case "foo bar":
                                    v107 := GenS6{F0: 32767, F1: -128, F2: (^(v84))}
                                    _ = v107
                                default:
                                    if ((!true) && true) {
                                        var v108 uint8 = v84
                                        switch (fi85 - fi35) {
                                        case 32767:
                                            var v109 uint8 = (v108 + v84)
                                            _ = v109
                                        case 2147483647:
                                            var v110 uint64 = 928
                                            _ = v110
                                        }
                                        v111 := GenS11{F0: 42, F1: 647, F2: (4294967295 == 1)}
                                        var v112 float32 = (((1e10 + 0.1) - v4) * v4)
                                        var v113 float64 = ((3.14159 - 0.0) * 3.14159)
                                        _ = v108
                                        _ = v111
                                        _ = v112
                                        _ = v113
                                    }
                                }
                                var v114 uint64 = 18446744073709551615
                                _ = fi85
                                _ = v114
                            }
                            _ = v84
                        }
                        var v115 int16 = -450
                        var v116 bool = ((v6 << ((179) & 63)) == v6)
                        _ = v115
                        _ = v116
                    }
                    var v117 float64 = -2.25
                    _ = v117
                }
                switch (fi35 ^ (fi35 / ((fi35) | 1))) {
                case 4294967295:
                    fmt.Println(v6)
                case -2147483648:
                    if (255 >= 4294967295) {
                        switch fi35 {
                        case -2147483648:
                            for fi118 := 0; fi118 < 8; fi118++ {
                                if true {
                                    v2 = v2
                                    if (4294967295 == 42) {
                                        fmt.Println(fi35, v6, v4)
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        v119 := GenS5{F0: 0, F1: 127}
                                        switch fi35 {
                                        case 32768:
                                            v120 := GenS12{F0: -213}
                                            _ = v120
                                        case 42:
                                            var v121 bool = false
                                            _ = v121
                                        case 32767:
                                            if false {
                                                var v122 string = v2
                                                _ = v122
                                            } else {
                                                if false {
                                                    fmt.Println(fi35)
                                                    switch fi35 {
                                                    case 32768:
                                                        if ((0.1 / 3.14159) < 3.14159) {
                                                            delete(v5, fi35)
                                                        } else {
                                                            if len(v3) > 0 {
                                                                v3[0] = ((v2 + v2) + v2)
                                                            }
                                                        }
                                                    case 2147483647:
                                                        v6 = 1
                                                    default:
                                                        switch v2 {
                                                        case "a":
                                                            if true {
                                                                v6 += v6
                                                                if (-709 != -392) {
                                                                    v6 += v6
                                                                    v123 := GenS13{F0: "hello"}
                                                                    _ = v123
                                                                }
                                                            }
                                                        }
                                                    }
                                                } else {
                                                    var v124 uint = 494
                                                    delete(v5, fi35)
                                                    var v125 uint64 = 455
                                                    _ = v124
                                                    _ = v125
                                                }
                                                switch v2 {
                                                case "a":
                                                    var v126 int8 = 127
                                                    _ = v126
                                                }
                                                var v127 uint16 = v6
                                                _ = v127
                                            }
                                        default:
                                            var v128 uint16 = (v6 / ((v6) | 1))
                                            _ = v128
                                        }
                                        _ = v119
                                    }
                                } else {
                                    var v129 uint8 = 156
                                    _ = v129
                                }
                                _ = fi118
                            }
                        default:
                            var v130 uint32 = 65535
                            _ = v130
                        }
                        if (0 < 0) {
                            for fi131 := 0; fi131 < 20; fi131++ {
                                _ = len(v1)
                                v6 = v6
                                var v132 string = (v2 + (v2 + (((v2 + v2) + ((v2 + (v2 + "a")) + "\n\t")) + v2)))
                                v133 := GenS13{F0: (v2 + ((v132 + "a") + ""))}
                                fmt.Println(v6, v132)
                                _ = fi131
                                _ = v132
                                _ = v133
                            }
                            for fi134 := 0; fi134 < 5; fi134++ {
                                switch v2 {
                                case "a":
                                    v6 = v6
                                case "\n\t":
                                    var v135 string = v2
                                    _ = v135
                                case "foo bar":
                                    v136 := GenS12{F0: 1}
                                    _ = v136
                                }
                                v2 = v2
                                _ = fi134
                            }
                            if len(v3) > 0 {
                                v3[0] = v2
                            }
                        }
                        switch v2 {
                        case "\n\t":
                            if len(v3) > 0 {
                                v3[0] = v2
                            }
                        }
                    } else {
                        v137 := GenS14{F0: 1, F1: 128, F2: (v6 | v6)}
                        _ = v137
                    }
                case -156:
                    var v138 int8 = -53
                    _ = v138
                }
                v139 := v5[fi35]
                _ = fi35
                _ = v139
            }
            v140 := GenS15{F0: 618, F1: -147}
            _ = v140
        } else {
            v4 -= ((v4 * v4) - -2.25)
            v141 := GenS3{F0: 4294967295}
            switch v2 {
            case "foo bar":
                var v142 float64 = 0.0
                _ = v142
            case "a":
                var v143 int64 = 195
                _ = v143
            default:
                switch ((v2 + ("a" + ("foo bar" + (v2 + v2)))) + v2) {
                case "\n\t":
                    for fi144 := 0; fi144 < 13; fi144++ {
                        if true {
                            _ = len(v1)
                            var v145 uint32 = 2147483648
                            _ = v145
                        } else {
                            var v146 int = (^(fi144))
                            for fi147 := 0; fi147 < 17; fi147++ {
                                v148 := GenS16{F0: v2, F1: 127, F2: 2}
                                var v149 int64 = 0
                                _ = fi147
                                _ = v148
                                _ = v149
                            }
                            fmt.Println(v2, v146, fi144)
                            _ = v146
                        }
                        var v150 int32 = 1
                        _ = fi144
                        _ = v150
                    }
                case "foo bar":
                    var v151 int32 = -2147483648
                    _ = v151
                case "a":
                    v4 = 0.0
                }
            }
            _ = v141
        }
    }
    var v152 uint32 = 818
    fmt.Println(v2, v2, v4)
    v6 = v6
    v153 := GenS17{F0: false, F1: uint64(v152), F2: v6}
    var v154 int8 = 42
    v155 := GenS18{F0: 1, F1: v2}
    switch -2147483648 {
    case 127:
        v156 := GenS6{F0: 767, F1: (v154 % (((-(v154))) | 1)), F2: 44}
        _ = v156
    case -124:
        fmt.Println(v2)
    case 1:
        switch -1 {
        case -436:
            var v157 int = 0
            _ = v157
        case -219:
            v6 = (v6 & v6)
        }
    }
    switch (v2 + ((v2 + v2) + ("foo bar" + v2))) {
    case "a":
        var v158 float32 = ((v4 - (v4 - v4)) / (v4 + v4))
        _ = v158
    default:
        v1[42] = 2
    }
    for fi159 := 0; fi159 < 14; fi159++ {
        v160 := GenS19{F0: (v2 + (v2 + "hello")), F1: "foo bar", F2: 1}
        _ = fi159
        _ = v160
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v152:", v152)
    fmt.Println("v154:", v154)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v152
    _ = v153
    _ = v154
    _ = v155
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
