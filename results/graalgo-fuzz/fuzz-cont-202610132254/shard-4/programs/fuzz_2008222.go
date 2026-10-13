package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16; F1 int; F2 uint }
type GenS4 struct { F0 uint32; F1 int8 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint16; F1 int16 }
type GenS8 struct { F0 int32; F1 uint32 }
type GenS9 struct { F0 int32; F1 int8; F2 int16 }
type GenS10 struct { F0 int; F1 int }
func genFn0() int8 {
    if (2147483648 < 65535) {
        if (42 > 2) {
            var v5 bool = (("a" + "") == "")
            for fi6 := 0; fi6 < 5; fi6++ {
                var v7 float32 = (1.5 * 0.0)
                _ = fi6
                _ = v7
            }
            switch "hello" {
            case "foo bar":
                v8 := []int{}
                _ = v8
            case "a":
                v9 := GenS1{F0: 255}
                _ = v9
            default:
                if (!v5) {
                    var v10 float64 = -2.25
                    switch 4294967295 {
                    case 2147483647:
                        v11 := []int{128, -9223372036854775808}
                        _ = v11
                    case 41:
                        var v12 string = ("foo bar" + "")
                        _ = v12
                    case 127:
                        var v13 int32 = 0
                        _ = v13
                    }
                    _ = v10
                } else {
                    if true {
                        var v14 int8 = -128
                        for fi15 := 0; fi15 < 8; fi15++ {
                            if ("foo bar" >= "a") {
                                if ((1e10 < 1.5) && v5) {
                                    v16 := []int{1, (fi15 + fi15), 32767, fi15}
                                    if len(v16) > 0 {
                                        v16[0] = (fi15 & (fi15 >> ((4294967295) & 63)))
                                    }
                                    _ = v16
                                }
                                var v17 uint16 = 406
                                _ = v17
                            }
                            _ = fi15
                        }
                        _ = v14
                    } else {
                        v18 := GenS2{F0: 1}
                        v5 = v5
                        if v5 {
                            for fi19 := 0; fi19 < 3; fi19++ {
                                v20 := GenS3{F0: 883, F1: (^(fi19)), F2: 65535}
                                _ = fi19
                                _ = v20
                            }
                        }
                        _ = v18
                    }
                }
            }
            _ = v5
        }
    } else {
        var v21 uint8 = 255
        _ = v21
    }
    var v22 uint32 = 1
    fmt.Println(v22)
    v23 := GenS3{F0: 9, F1: 221, F2: 1}
    _ = v22
    _ = v23
    return -128
}
func genFn1(p0 string, p1 int16) uint16 {
    for fi24 := 0; fi24 < 12; fi24++ {
        p1 = (p1 * (p1 + (^(p1))))
        for fi25 := 0; fi25 < 7; fi25++ {
            if (false || true) {
                var v26 uint8 = 42
                _ = v26
            }
            if false {
                v27 := GenS4{F0: 2147483647, F1: -1}
                switch ("hello" + (p0 + p0)) {
                case "foo bar":
                    p1 = (p1 << ((2) & 63))
                case "a":
                    if true {
                        var v28 float32 = (-2.25 + (0.1 + 1e10))
                        _ = v28
                    }
                case "":
                    if (0 < 255) {
                        p0 = p0
                        if true {
                            if false {
                                v29 := GenS4{F0: 4294967295, F1: 72}
                                v30 := GenS5{F0: 2147483647}
                                _ = v29
                                _ = v30
                            } else {
                                switch (p0 + ("\n\t" + (p0 + p0))) {
                                case "a":
                                    v31 := GenS5{F0: 148}
                                    _ = v31
                                case "foo bar":
                                    p1 -= p1
                                    p1 += (p1 >> ((4294967295) & 63))
                                    fmt.Println(p1, fi24, p0)
                                    if true {
                                        var v32 uint64 = 117
                                        _ = v32
                                    }
                                    if true {
                                        v33 := []int{}
                                        _ = v33
                                    }
                                }
                                if (true && false) {
                                    p1 += (p1 - (p1 * (p1 & -17)))
                                } else {
                                    p0 = (p0 + (p0 + p0))
                                }
                            }
                            v34 := []int{}
                            _ = v34
                        }
                    }
                }
                switch 32767 {
                case -706:
                    fmt.Println(fi25, p1)
                case -135:
                    switch fi25 {
                    case -1:
                        p1 += (p1 | (p1 << ((137) & 63)))
                        var v35 float32 = 1e10
                        _ = v35
                    case 774:
                        if (((p1 | (p1 * p1)) == (p1 / (((p1 - p1)) | 1))) || true) {
                            var v36 uint = 365
                            switch fi25 {
                            case 42:
                                p1 = p1
                            }
                            if false {
                                switch "hello" {
                                case "\n\t":
                                    if false {
                                        p1 -= (^(p1))
                                        v36 -= (v36 ^ v36)
                                        v37 := GenS6{F0: 118}
                                        _ = v37
                                    } else {
                                        fmt.Println(fi24, p1)
                                        var v38 int16 = (p1 | p1)
                                        if (fi24 > (fi25 & (fi25 ^ (fi25 - (fi25 >> (((v36 & 837)) & 63)))))) {
                                            var v39 uint32 = 65535
                                            v40 := GenS6{F0: v39}
                                            _ = v39
                                            _ = v40
                                        } else {
                                            fmt.Println(p1, v36, p0)
                                            fmt.Println(p0, v36)
                                            var v41 float32 = 0.0
                                            _ = v41
                                        }
                                        _ = v38
                                    }
                                case "hello":
                                    var v42 int8 = 42
                                    _ = v42
                                case "a":
                                    if true {
                                        fmt.Println(v36, fi24)
                                        var v43 int = fi25
                                        _ = v43
                                    }
                                default:
                                    if (-2147483648 <= -1) {
                                        p0 = p0
                                    }
                                }
                                var v44 string = "foo bar"
                                fmt.Println(fi25)
                                _ = v44
                            }
                            _ = v36
                        }
                    default:
                        var v45 int8 = genFn0()
                        _ = v45
                    }
                case 32768:
                    v46 := GenS5{F0: -334}
                    _ = v46
                }
                switch "hello" {
                case "\n\t":
                    fmt.Println(p0)
                case "foo bar":
                    var v47 bool = false
                    _ = v47
                case "":
                    fmt.Println(fi24)
                }
                _ = v27
            }
            fmt.Println(fi24)
            _ = fi25
        }
        _ = fi24
    }
    switch -265 {
    case -1:
        fmt.Println(p1, p0)
    case 4294967295:
        if (255 < 0) {
            for fi48 := 0; fi48 < 19; fi48++ {
                v49 := GenS7{F0: 312, F1: -32768}
                _ = fi48
                _ = v49
            }
            switch -502 {
            case -224:
                var v50 string = p0
                _ = v50
            case 127:
                p0 = ((p0 + "hello") + (p0 + p0))
            case 1:
                fmt.Println(p0, p0)
            default:
                v51 := []int{2147483647}
                _ = v51
            }
            switch 32767 {
            case 882:
                var v52 int16 = (p1 ^ p1)
                _ = v52
            case -9223372036854775808:
                p1 += (-(p1))
            case -603:
                v53 := GenS1{F0: 65535}
                _ = v53
            default:
                v54 := []int{-437, 263, 42}
                _ = v54
            }
        }
    }
    p1 = (p1 % (((-(p1))) | 1))
    switch 128 {
    case 965:
        for fi55 := 0; fi55 < 7; fi55++ {
            for fi56 := 0; fi56 < 5; fi56++ {
                var v57 string = (p0 + ((((p0 + "a") + p0) + (p0 + (p0 + p0))) + p0))
                _ = fi56
                _ = v57
            }
            fmt.Println(p1)
            if ((fi55 | (fi55 | fi55)) < (fi55 - (fi55 % ((fi55) | 1)))) {
                fmt.Println(p1)
                if true {
                    if (false || true) {
                        for fi58 := 0; fi58 < 19; fi58++ {
                            var v59 uint8 = 127
                            v59 += (v59 % (((v59 >> ((728) & 63))) | 1))
                            p1 = 1000
                            v60 := GenS8{F0: 42, F1: 1}
                            switch p0 {
                            case "foo bar":
                                var v61 int8 = genFn0()
                                _ = v61
                            }
                            _ = fi58
                            _ = v59
                            _ = v60
                        }
                        switch (p0 + p0) {
                        case "a":
                            fmt.Println(p1, p0)
                        default:
                            var v62 int64 = -1
                            _ = v62
                        }
                        var v63 uint16 = 65535
                        _ = v63
                    } else {
                        p0 = p0
                    }
                    var v64 int = fi55
                    switch v64 {
                    case 32768:
                        for fi65 := 0; fi65 < 18; fi65++ {
                            v66 := []int{}
                            v64 += fi65
                            v67 := GenS4{F0: 2, F1: -96}
                            var v68 uint8 = 18
                            v69 := GenS3{F0: 2, F1: (^(fi55)), F2: 724}
                            _ = fi65
                            _ = v66
                            _ = v67
                            _ = v68
                            _ = v69
                        }
                    }
                    _ = v64
                }
                if (!false) {
                    p0 = ""
                    var v70 int32 = 116
                    _ = v70
                } else {
                    for fi71 := 0; fi71 < 9; fi71++ {
                        v72 := []int{fi55, 127, fi71}
                        p1 = (p1 ^ -810)
                        var v73 int = (^(fi55))
                        _ = fi71
                        _ = v72
                        _ = v73
                    }
                    var v74 int64 = 42
                    p0 = "\n\t"
                    _ = v74
                }
            } else {
                var v75 int16 = (^(p1))
                _ = v75
            }
            _ = fi55
        }
    case 606:
        var v76 int64 = 9223372036854775807
        _ = v76
    case 2147483647:
        fmt.Println(p0, p1)
    }
    _ = p0
    _ = p1
    return 65535
}

func main() {
    var v1 uint16 = 1
    v2 := make(map[uint32]string)
    var v3 float64 = (((0.1 - (1.5 + (1.5 * 1e10))) + 0.1) + 3.14159)
    v4 := make(map[int]int)
    delete(v2, 760)
    v3 += (float64(1.5) * v3)
    var v77 string = "\n\t"
    var v78 uint64 = 6
    var v79 uint = 65535
    v79 = (^(v79))
    v1 -= (v1 % (((v1 / ((v1) | 1))) | 1))
    switch -1 {
    case 703:
        if false {
            v80 := []int{2147483647}
            fmt.Println(v79, v79)
            _ = v80
        }
    default:
        v81 := GenS7{F0: 2, F1: 505}
        _ = v81
    }
    if (true || false) {
        if true {
            _ = len(v2)
            if (287 > 0) {
                switch 32768 {
                case 2147483647:
                    v3 = v3
                }
            }
            if true {
                v82 := GenS9{F0: 165, F1: -1, F2: 723}
                _ = v82
            }
        }
        fmt.Println(v78, v77)
    } else {
        var v83 int16 = -902
        for fi84 := 0; fi84 < 19; fi84++ {
            for fi85 := 0; fi85 < 2; fi85++ {
                v3 = 1e10
                switch v77 {
                case "a":
                    v1 = 1
                case "":
                    fmt.Println(v1, v79)
                case "hello":
                    v86 := v2[0]
                    _ = v86
                }
                v87 := []int{(fi85 | 127), (-(fi84)), fi85}
                _ = fi85
                _ = v87
            }
            _ = fi84
        }
        if true {
            v88 := GenS4{F0: 1, F1: -1}
            var v89 int64 = 9223372036854775807
            for fi90 := 0; fi90 < 6; fi90++ {
                if ((!(249 != 0)) && true) {
                    for fi91 := 0; fi91 < 12; fi91++ {
                        var v92 uint8 = 0
                        _ = fi91
                        _ = v92
                    }
                }
                var v93 string = v77
                _ = fi90
                _ = v93
            }
            _ = v88
            _ = v89
        } else {
            if false {
                v94 := []int{1, 4294967295, 32768}
                if (0 >= (v83 - 32767)) {
                    v79 = v79
                }
                _ = v94
            } else {
                var v95 uint32 = uint32(v79)
                if false {
                    if false {
                        var v96 int16 = v83
                        switch "\n\t" {
                        case "foo bar":
                            switch v77 {
                            case "foo bar":
                                switch -108 {
                                case 0:
                                    fmt.Println(v95)
                                case -396:
                                    fmt.Println(v1)
                                default:
                                    var v97 string = v77
                                    _ = v97
                                }
                            }
                        }
                        _ = v96
                    }
                    switch v77 {
                    case "a":
                        var v98 uint8 = 253
                        _ = v98
                    default:
                        _ = len(v4)
                    }
                }
                var v99 int16 = 32767
                _ = v95
                _ = v99
            }
            for fi100 := 0; fi100 < 15; fi100++ {
                var v101 float32 = float32(1e10)
                _ = fi100
                _ = v101
            }
            v78 -= (v78 + (v78 >> (((v79 ^ 255)) & 63)))
        }
        _ = v83
    }
    switch ("" + (v77 + v77)) {
    case "hello":
        if (42 != 9223372036854775807) {
            if false {
                for fi102 := 0; fi102 < 16; fi102++ {
                    var v103 int32 = -2147483648
                    v104 := GenS10{F0: -879, F1: (fi102 - (fi102 - (fi102 >> ((v79) & 63))))}
                    if true {
                        if (v103 < 504) {
                            for fi105 := 0; fi105 < 7; fi105++ {
                                var v106 int = (fi105 / ((-65) | 1))
                                if false {
                                    v107 := GenS6{F0: 65535}
                                    v108 := GenS7{F0: genFn1("", 32767), F1: 38}
                                    v77 = v77
                                    _ = v107
                                    _ = v108
                                } else {
                                    v106 -= 128
                                    v78 = (v78 & v78)
                                    v78 = v78
                                    v109 := []int{2147483647, (fi102 >> ((2) & 63)), v106}
                                    var v110 int16 = -503
                                    _ = v109
                                    _ = v110
                                }
                                _ = fi105
                                _ = v106
                            }
                        } else {
                            var v111 uint64 = (v78 % ((v78) | 1))
                            v3 -= float64(1.5)
                            var v112 bool = false
                            _ = v111
                            _ = v112
                        }
                        var v113 uint64 = (v78 << ((v79) & 63))
                        _ = v113
                    }
                    _ = fi102
                    _ = v103
                    _ = v104
                }
                var v114 uint32 = 954
                var v115 uint16 = uint16(v78)
                _ = v114
                _ = v115
            }
            v116 := v4[591]
            _ = v116
        } else {
            delete(v4, 0)
            var v117 int8 = -100
            v3 += (-2.25 / v3)
            _ = v117
        }
    }
    switch -356 {
    case -985:
        for fi118 := 0; fi118 < 16; fi118++ {
            for fi119 := 0; fi119 < 11; fi119++ {
                switch (v77 + v77) {
                case "a":
                    var v120 int8 = 0
                    _ = v120
                case "foo bar":
                    fmt.Println(fi119)
                case "":
                    v121 := GenS10{F0: fi119, F1: (fi118 & -9223372036854775808)}
                    _ = v121
                }
                var v122 uint16 = 255
                switch fi118 {
                case -1:
                    var v123 bool = ((^(v78)) >= v78)
                    _ = v123
                case 1:
                    v79 -= v79
                default:
                    v4[-896] = (fi118 + 2147483647)
                }
                _ = fi119
                _ = v122
            }
            switch (v77 + (v77 + v77)) {
            case "hello":
                fmt.Println(v78, v3)
            }
            _ = fi118
        }
    case 127:
        var v124 int16 = -1
        _ = v124
    case 377:
        fmt.Println(v3, v1, v79)
    }
    if ((1 >= 4294967295) || false) {
        v125 := []int{261, 42, 42}
        fmt.Println(v77)
        _ = v125
    } else {
        fmt.Println(v79)
        switch "foo bar" {
        case "a":
            if false {
                for fi126 := 0; fi126 < 7; fi126++ {
                    for fi127 := 0; fi127 < 8; fi127++ {
                        v128 := v4[fi127]
                        fmt.Println(v128)
                        v129 := GenS4{F0: 42, F1: -47}
                        _ = fi127
                        _ = v128
                        _ = v129
                    }
                    if (!true) {
                        var v130 bool = (v77 > v77)
                        v130 = v130
                        _ = v130
                    }
                    for fi131 := 0; fi131 < 18; fi131++ {
                        fmt.Println(v78, v77, fi131)
                        _ = fi131
                    }
                    _ = fi126
                }
                delete(v2, 2147483648)
                v132 := []int{32767, 2147483647}
                _ = v132
            } else {
                if (!true) {
                    for fi133 := 0; fi133 < 19; fi133++ {
                        for fi134 := 0; fi134 < 17; fi134++ {
                            if (65535 < (v1 - (v1 * v1))) {
                                _ = len(v2)
                            }
                            if ((fi133 * -1) >= fi134) {
                                fmt.Println(v1)
                                v3 = float64(float32(v3))
                            }
                            v1 = v1
                            _ = fi134
                        }
                        _ = fi133
                    }
                }
                v78 = 4294967295
                v135 := v4[42]
                _ = v135
            }
        default:
            for fi136 := 0; fi136 < 18; fi136++ {
                fmt.Println(v1)
                _ = fi136
            }
        }
        v78 += v78
    }
    for fi137 := 0; fi137 < 2; fi137++ {
        if false {
            v1 = v1
            v77 = "a"
            v78 -= (v78 - (v78 << ((2) & 63)))
        } else {
            v78 -= 18446744073709551615
            v138 := GenS7{F0: v1, F1: 32767}
            _ = v138
        }
        v139 := []int{int(v78), (fi137 << ((412) & 63))}
        _ = fi137
        _ = v139
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v77
    _ = v78
    _ = v79
}
