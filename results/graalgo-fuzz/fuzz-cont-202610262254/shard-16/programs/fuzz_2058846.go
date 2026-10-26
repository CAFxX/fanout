package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 int64; F1 int }
type GenS5 struct { F0 uint; F1 uint32; F2 uint64 }
type GenS6 struct { F0 int16; F1 int32 }
type GenS7 struct { F0 int; F1 bool }
type GenS8 struct { F0 int16; F1 int8 }
func genFn0(p0 int16, p1 int64, p2 uint64) int8 {
    if true {
        if true {
            switch 311 {
            case 295:
                v6 := GenS1{F0: 42, F1: 168, F2: (p0 >> ((0) & 63))}
                _ = v6
            default:
                var v7 uint64 = (p2 ^ 4294967295)
                _ = v7
            }
        }
        var v8 string = "foo bar"
        p1 = (p1 >> ((65535) & 63))
        _ = v8
    }
    v9 := []int{int(p0), 32768, -9223372036854775808}
    var v10 uint64 = 832
    for fi11 := 0; fi11 < 12; fi11++ {
        p2 += p2
        v10 += (^(v10))
        if false {
            p1 -= p1
        }
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 100
}
func genFn1(p0 int32, p1 int) uint {
    fmt.Println(p0)
    var v12 bool = (((18446744073709551615 >= 1) || true) && false)
    if (v12 || v12) {
        v13 := []int{(p1 / ((-519) | 1))}
        if v12 {
            v12 = true
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
        }
        _ = v13
    } else {
        for fi14 := 0; fi14 < 3; fi14++ {
            if v12 {
                if ((v12 && false) || (!(!(!(!v12))))) {
                    switch "hello" {
                    case "foo bar":
                        if (32767 > 1) {
                            for fi15 := 0; fi15 < 9; fi15++ {
                                v16 := []int{p1}
                                p1 -= fi14
                                _ = fi15
                                _ = v16
                            }
                            if false {
                                if false {
                                    v17 := GenS2{F0: 32767}
                                    _ = v17
                                }
                            }
                        }
                    case "hello":
                        if (2 != 2) {
                            for fi18 := 0; fi18 < 6; fi18++ {
                                fmt.Println(p0, fi14, p0)
                                v19 := []int{fi18, (fi14 << ((0) & 63))}
                                _ = fi18
                                _ = v19
                            }
                            switch "" {
                            case "a":
                                var v20 int8 = 42
                                _ = v20
                            case "foo bar":
                                var v21 string = ""
                                _ = v21
                            case "\n\t":
                                if (((("foo bar" + ("" + "")) + ("" + "")) + (("hello" + "foo bar") + "a")) != "hello") {
                                    var v22 uint8 = 38
                                    for fi23 := 0; fi23 < 1; fi23++ {
                                        switch (p1 + fi14) {
                                        case -879:
                                            var v24 uint16 = uint16(v22)
                                            _ = v24
                                        case 32768:
                                            v25 := []int{-378}
                                            _ = v25
                                        default:
                                            var v26 uint64 = 828
                                            _ = v26
                                        }
                                        v27 := GenS2{F0: -32768}
                                        v28 := []int{fi23, (p1 + (fi23 >> ((1) & 63)))}
                                        _ = fi23
                                        _ = v27
                                        _ = v28
                                    }
                                    _ = v22
                                } else {
                                    for fi29 := 0; fi29 < 9; fi29++ {
                                        p0 = (p0 * -735)
                                        fmt.Println(fi29, p0, fi29)
                                        _ = fi29
                                    }
                                }
                            }
                            var v30 int = 32767
                            _ = v30
                        }
                    default:
                        v12 = (v12 || (0.0 < -2.25))
                    }
                } else {
                    switch (fi14 * (p1 | fi14)) {
                    case 461:
                        if (243 != 127) {
                            v31 := []int{fi14, 32768}
                            _ = v31
                        }
                    case 1:
                        v32 := GenS1{F0: 543, F1: 255, F2: 1000}
                        _ = v32
                    case 502:
                        v12 = ((0.0 == float32((1e10 * (float64(0.0) + 0.1)))) && v12)
                    }
                    fmt.Println(fi14, p1)
                    if false {
                        if false {
                            fmt.Println(p0, p1)
                        }
                        p0 += p0
                    }
                }
                fmt.Println(p0, p0)
                for fi33 := 0; fi33 < 6; fi33++ {
                    switch (^(fi14)) {
                    case 32767:
                        v34 := GenS1{F0: 594, F1: 0, F2: 1000}
                        _ = v34
                    default:
                        v35 := []int{p1, (p1 % (((fi14 << ((65535) & 63))) | 1)), p1, fi33}
                        _ = v35
                    }
                    fmt.Println(p1, p0)
                    _ = fi33
                }
            }
            _ = fi14
        }
        v36 := []int{32767, p1, p1}
        var v37 string = "a"
        _ = v36
        _ = v37
    }
    switch ((("a" + "a") + ("a" + "\n\t")) + "") {
    case "":
        v12 = (!false)
    case "a":
        var v38 int32 = (p0 & p0)
        _ = v38
    case "foo bar":
        v39 := []int{p1, 9223372036854775807, p1, (p1 % (((^(p1))) | 1))}
        _ = v39
    default:
        v12 = v12
    }
    _ = p0
    _ = p1
    _ = v12
    return 1
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]bool)
    v3 := make(map[uint32]bool)
    v4 := make(map[int]int)
    var v5 float32 = 0.1
    switch "hello" {
    case "foo bar":
        var v40 int16 = 1
        _ = v40
    case "\n\t":
        switch -2147483648 {
        case 721:
            for fi41 := 0; fi41 < 2; fi41++ {
                var v42 int8 = -1
                v5 += (v5 * -2.25)
                _ = fi41
                _ = v42
            }
        case 0:
            v43 := GenS3{F0: 42, F1: genFn0(-32768, 42, 0)}
            _ = v43
        case 181:
            v44 := []int{0, -2147483648}
            _ = v44
        default:
            var v45 int = 127
            _ = v45
        }
    case "a":
        v46 := GenS1{F0: 599, F1: 42, F2: -262}
        _ = v46
    }
    for fi47 := 0; fi47 < 1; fi47++ {
        if (735 >= -1) {
            v48 := GenS4{F0: 9223372036854775807, F1: fi47}
            v49 := []int{4294967295, 127, fi47, (^(fi47))}
            _ = v48
            _ = v49
        }
        _ = fi47
    }
    v5 = ((float32((3.14159 * 1e10)) + v5) - (v5 + v5))
    var v50 string = ("foo bar" + "foo bar")
    if false {
        if ((3.14159 > ((0.0 - (((1e10 + 3.14159) + (float64((-2.25 + 0.0)) + (3.14159 + 3.14159))) / 0.1)) / -2.25)) && true) {
            v5 -= v5
            var v51 int8 = 42
            v52 := GenS2{F0: -1}
            _ = v51
            _ = v52
        } else {
            for fi53 := 0; fi53 < 20; fi53++ {
                var v54 int64 = 1
                _ = fi53
                _ = v54
            }
            var v55 int8 = -128
            if true {
                v55 = v55
                fmt.Println(v55)
            }
            _ = v55
        }
        if (0 >= 4294967295) {
            fmt.Println(v50, v50)
            var v56 int32 = 49
            _ = v56
        }
        fmt.Println(v50, v50)
    } else {
        var v57 string = v50
        var v58 uint8 = 255
        _ = v57
        _ = v58
    }
    var v59 int64 = -107
    v60 := GenS4{F0: (-(v59)), F1: 227}
    v61 := []int{}
    v61 = append(v61, 32767)
    if (false && (1 <= 42)) {
        for fi62 := 0; fi62 < 5; fi62++ {
            for fi63 := 0; fi63 < 17; fi63++ {
                v50 = "\n\t"
                v64 := GenS5{F0: 0, F1: 1, F2: 0}
                _ = fi63
                _ = v64
            }
            _ = fi62
        }
        for fi65 := 0; fi65 < 15; fi65++ {
            switch v50 {
            case "a":
                var v66 int64 = 1
                _ = v66
            case "\n\t":
                var v67 uint8 = 0
                _ = v67
            case "":
                var v68 uint16 = 42
                _ = v68
            default:
                if (!false) {
                    v59 -= 42
                    if true {
                        v69 := v2[(fi65 % (((fi65 / ((fi65) | 1))) | 1))]
                        if false {
                            if len(v61) >= 2 {
                                v61 = v61[1:len(v61)]
                            }
                            var v70 uint16 = 0
                            for fi71 := 0; fi71 < 5; fi71++ {
                                var v72 uint = 2
                                var v73 int = (fi71 * fi71)
                                var v74 int64 = (v59 - (^(v59)))
                                _ = fi71
                                _ = v72
                                _ = v73
                                _ = v74
                            }
                            _ = v70
                        } else {
                            if len(v61) > 0 {
                                v61[0] = (fi65 | fi65)
                            }
                            v75 := v3[4294967295]
                            _ = v75
                        }
                        v76 := v4[-861]
                        _ = v69
                        _ = v76
                    }
                    v77 := GenS1{F0: 0, F1: 196, F2: 1}
                    _ = v77
                } else {
                    fmt.Println(fi65, v5, fi65)
                    v78 := GenS5{F0: 42, F1: 65535, F2: 635}
                    v3[2147483647] = (3.14159 >= 0.1)
                    _ = v78
                }
            }
            var v79 float32 = 1.5
            _ = fi65
            _ = v79
        }
        for fi80 := 0; fi80 < 11; fi80++ {
            if (v5 == ((float32((3.14159 + 1e10)) * 0.1) + (((3.14159 + float32(-2.25)) + 1e10) - ((v5 / v5) / (0.1 * (v5 - 1.5)))))) {
                if false {
                    v5 -= 0.0
                    v5 = v5
                } else {
                    switch fi80 {
                    case 42:
                        var v81 int32 = -1
                        _ = v81
                    case 1:
                        switch "\n\t" {
                        case "a":
                            if true {
                                v5 += v5
                                v82 := GenS6{F0: 580, F1: 0}
                                _ = v82
                            }
                        case "foo bar":
                            for fi83 := 0; fi83 < 16; fi83++ {
                                v50 = v50
                                _ = fi83
                            }
                        case "\n\t":
                            switch "hello" {
                            case "hello":
                                var v84 int8 = 0
                                _ = v84
                            case "a":
                                v50 = v50
                            case "\n\t":
                                switch fi80 {
                                case 32767:
                                    var v85 uint = 4294967295
                                    _ = v85
                                case 400:
                                    for fi86 := 0; fi86 < 15; fi86++ {
                                        var v87 int16 = 42
                                        var v88 string = (v50 + "hello")
                                        v89 := GenS5{F0: 440, F1: 42, F2: 4294967295}
                                        _ = fi86
                                        _ = v87
                                        _ = v88
                                        _ = v89
                                    }
                                default:
                                    v90 := GenS4{F0: v59, F1: fi80}
                                    _ = v90
                                }
                            }
                        }
                    case 4294967295:
                        if (221 != 255) {
                            for fi91 := 0; fi91 < 7; fi91++ {
                                v92 := v1[1]
                                if v92 {
                                    var v93 int8 = 0
                                    _ = v93
                                }
                                _ = fi91
                                _ = v92
                            }
                            var v94 string = "foo bar"
                            _ = v94
                        } else {
                            switch fi80 {
                            case 0:
                                var v95 uint16 = 0
                                _ = v95
                            case 1:
                                _ = len(v1)
                            case -668:
                                v96 := GenS7{F0: 2147483647, F1: true}
                                _ = v96
                            default:
                                var v97 int64 = (v59 << ((18446744073709551615) & 63))
                                _ = v97
                            }
                            v50 = "\n\t"
                        }
                    default:
                        v98 := GenS6{F0: 1, F1: -2147483648}
                        _ = v98
                    }
                }
            }
            _ = fi80
        }
    } else {
        fmt.Println(v59, v50)
        var v99 string = v50
        var v100 int16 = 247
        _ = v100
        _ = v99
    }
    v59 = (v59 | 1)
    var v101 float64 = (0.0 + -2.25)
    var v102 uint64 = 42
    switch -647 {
    case 1:
        for fi103 := 0; fi103 < 11; fi103++ {
            var v104 int64 = v59
            _ = fi103
            _ = v104
        }
    case 0:
        v5 = (0.0 + (v5 / v5))
    default:
        switch ((v50 + v50) + (v50 + (("foo bar" + "") + v50))) {
        case "":
            if (true && (100 <= int8(v59))) {
                v50 = v50
                v50 = v50
                v61 = append(v61, -460)
            } else {
                var v105 string = v50
                for fi106 := 0; fi106 < 4; fi106++ {
                    var v107 uint64 = v102
                    _ = fi106
                    _ = v107
                }
                _ = v105
            }
        case "\n\t":
            v108 := GenS6{F0: -361, F1: -1}
            _ = v108
        default:
            if false {
                var v109 uint16 = 967
                v50 = (v50 + v50)
                if ((v102 / (((v102 + 42)) | 1)) >= (v102 % ((v102) | 1))) {
                    var v110 uint16 = (^(v109))
                    if true {
                        for fi111 := 0; fi111 < 20; fi111++ {
                            v101 += ((v101 / v101) * float64(1.5))
                            _ = fi111
                        }
                        switch 913 {
                        case 32767:
                            v1[4294967295] = (0.0 > v101)
                        }
                    }
                    var v112 float32 = v5
                    _ = v110
                    _ = v112
                } else {
                    v61 = append(v61, 32768)
                    v109 -= v109
                    if (42 == -385) {
                        for fi113 := 0; fi113 < 10; fi113++ {
                            var v114 uint64 = 4294967295
                            v115 := GenS6{F0: 32767, F1: 1}
                            v102 = (v102 + v102)
                            _ = fi113
                            _ = v114
                            _ = v115
                        }
                    }
                }
                _ = v109
            }
        }
    }
    v116 := GenS7{F0: -9223372036854775808, F1: false}
    if false {
        if true {
            v5 += ((v5 * -2.25) - (-2.25 - (v5 * 1.5)))
            v117 := GenS6{F0: 32767, F1: 2147483647}
            if true {
                v102 = (v102 / ((v102) | 1))
                if len(v61) >= 2 {
                    v61 = v61[1:len(v61)]
                }
            } else {
                var v118 uint16 = 42
                v59 = (v59 / (((v59 - (v59 << ((4294967295) & 63)))) | 1))
                fmt.Println(v102)
                _ = v118
            }
            _ = v117
        }
        var v119 bool = ((2147483647 == -129) && (!true))
        _ = v119
    } else {
        if (true && false) {
            v101 = 0.0
            for fi120 := 0; fi120 < 11; fi120++ {
                if (true || false) {
                    if true {
                        var v121 int16 = 1
                        _ = v121
                    }
                    v102 -= (^(v102))
                    v5 = ((((v5 / v5) * 0.1) * 0.0) / ((float32((v101 * v101)) * v5) - (-2.25 / ((-2.25 - v5) * v5))))
                } else {
                    var v122 int = (fi120 - fi120)
                    _ = v122
                }
                for fi123 := 0; fi123 < 7; fi123++ {
                    var v124 int = (^(fi120))
                    switch fi120 {
                    case 0:
                        if false {
                            v125 := GenS3{F0: 126, F1: -12}
                            switch (v50 + ((v50 + ("hello" + v50)) + ((((v50 + (v50 + v50)) + ("foo bar" + (v50 + v50))) + v50) + (v50 + ((v50 + v50) + v50))))) {
                            case "foo bar":
                                switch 42 {
                                case 962:
                                    v126 := v3[209]
                                    _ = v126
                                case 762:
                                    v124 = v124
                                case 32768:
                                    if len(v61) >= 2 {
                                        v61 = v61[1:len(v61)]
                                    }
                                }
                            }
                            _ = v125
                        } else {
                            if len(v61) >= 2 {
                                v61 = v61[1:len(v61)]
                            }
                            fmt.Println(v5, v124)
                        }
                        v127 := GenS1{F0: 255, F1: 150, F2: 727}
                        fmt.Println(v124)
                        _ = v127
                    case 127:
                        if false {
                            var v128 int64 = v59
                            _ = v128
                        } else {
                            _ = len(v4)
                            _ = len(v61)
                        }
                    case 42:
                        var v129 int8 = 1
                        _ = v129
                    default:
                        if (!(255 >= 148)) {
                            v59 -= v59
                            var v130 int8 = 0
                            _ = v130
                        } else {
                            var v131 uint16 = 65535
                            var v132 uint64 = v102
                            if false {
                                v133 := GenS5{F0: 42, F1: 2147483648, F2: (v102 / (((v132 * uint64(v131))) | 1))}
                                if true {
                                    var v134 uint64 = (v102 % ((653) | 1))
                                    _ = v134
                                } else {
                                    if false {
                                        switch v50 {
                                        case "a":
                                            v4[32768] = v124
                                        case "":
                                            v50 = v50
                                        case "\n\t":
                                            v61 = append(v61, fi120)
                                        }
                                        var v135 uint32 = 2147483647
                                        _ = v135
                                    }
                                    v50 = (v50 + v50)
                                }
                                var v136 uint32 = 65535
                                _ = v133
                                _ = v136
                            }
                            _ = v131
                            _ = v132
                        }
                    }
                    v137 := GenS5{F0: 255, F1: 2, F2: (v102 - 2)}
                    v5 = -2.25
                    _ = fi123
                    _ = v124
                    _ = v137
                }
                switch (v50 + v50) {
                case "\n\t":
                    switch 9223372036854775807 {
                    case -1:
                        if (!(65535 >= 177)) {
                            for fi138 := 0; fi138 < 2; fi138++ {
                                var v139 int = int(v59)
                                if len(v61) > 0 {
                                    v61[0] = (-(v139))
                                }
                                if true {
                                    var v140 float64 = (float64(((0.1 + v5) - v5)) * 1e10)
                                    var v141 float64 = v101
                                    _ = v140
                                    _ = v141
                                }
                                _ = fi138
                                _ = v139
                            }
                            var v142 uint16 = 773
                            _ = v142
                        }
                    case -2147483648:
                        for fi143 := 0; fi143 < 12; fi143++ {
                            v59 -= v59
                            switch 9223372036854775807 {
                            case 32767:
                                _ = len(v61)
                            case 413:
                                v59 = -9223372036854775808
                            }
                            _ = fi143
                        }
                    }
                case "hello":
                    if ((v101 / v101) > 3.14159) {
                        if (int32(v59) <= 42) {
                            fmt.Println(v59)
                            var v144 float64 = (float64((v5 - v5)) + v101)
                            _ = v144
                        } else {
                            v145 := GenS8{F0: -597, F1: 100}
                            _ = v145
                        }
                        fmt.Println(v50, v101, v102)
                    } else {
                        v61 = append(v61, fi120)
                        if true {
                            for fi146 := 0; fi146 < 5; fi146++ {
                                var v147 string = "\n\t"
                                v61 = append(v61, 770)
                                var v148 int16 = 42
                                _ = fi146
                                _ = v147
                                _ = v148
                            }
                            var v149 uint8 = 42
                            _ = v149
                        } else {
                            fmt.Println(v101, v50, v59)
                        }
                    }
                case "a":
                    switch fi120 {
                    case 1:
                        delete(v2, (fi120 % ((370) | 1)))
                    default:
                        v2[fi120] = (int32(fi120) == 1)
                    }
                default:
                    fmt.Println(v50, v101)
                }
                _ = fi120
            }
            if (false || ((42 < -323) && (("" + (("" + ("foo bar" + ((v50 + ("a" + (v50 + "foo bar"))) + v50))) + "a")) > (v50 + "")))) {
                fmt.Println(v50)
            }
        } else {
            var v150 uint8 = 127
            _ = len(v1)
            var v151 uint16 = 32767
            _ = v150
            _ = v151
        }
        v152 := GenS3{F0: 127, F1: -128}
        _ = v152
    }
    fmt.Println(v102, v50, v101)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v101:", v101)
    fmt.Println("v102:", v102)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v59:", v59)
    fmt.Println("slicelen v61:", len(v61), v61)
    _ = v1
    _ = v101
    _ = v102
    _ = v116
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v59
    _ = v60
    _ = v61
}
