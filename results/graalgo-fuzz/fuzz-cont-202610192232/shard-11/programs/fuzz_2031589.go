package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 bool; F1 int64; F2 uint }
type GenS5 struct { F0 int16; F1 int32 }
type GenS6 struct { F0 uint; F1 int; F2 uint32 }
type GenS7 struct { F0 uint32; F1 bool }
type GenS8 struct { F0 int; F1 bool }
type GenS9 struct { F0 uint8; F1 bool }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 bool; F1 uint8; F2 uint }
type GenS12 struct { F0 uint64; F1 string; F2 int16 }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 bool; F1 uint32 }
type GenS16 struct { F0 bool; F1 string; F2 int32 }
type GenS17 struct { F0 int8; F1 bool }
type GenS18 struct { F0 int8; F1 uint8 }
type GenS19 struct { F0 uint; F1 int }
type GenS20 struct { F0 int32; F1 string }
type GenS21 struct { F0 uint8; F1 int8; F2 uint }
type GenS22 struct { F0 int8; F1 uint; F2 int8 }
type GenS23 struct { F0 int64; F1 int64 }
type GenS24 struct { F0 uint32 }
type GenS25 struct { F0 uint8; F1 uint8 }
type GenS26 struct { F0 int64; F1 int16 }
type GenS27 struct { F0 int; F1 uint; F2 uint64 }
type GenS28 struct { F0 int16; F1 uint16 }
type GenS29 struct { F0 bool; F1 uint8 }
type GenS30 struct { F0 uint32; F1 int32 }
type GenS31 struct { F0 int8; F1 int; F2 int8 }
func genFn0() int64 {
    for fi4 := 0; fi4 < 5; fi4++ {
        var v5 int = 2147483647
        v6 := []int{1, -2147483648, (v5 | 2147483647)}
        v7 := GenS1{F0: 2, F1: 0}
        _ = fi4
        _ = v5
        _ = v6
        _ = v7
    }
    switch 32767 {
    case 1:
        var v8 int = 32768
        _ = v8
    case 454:
        var v9 uint16 = 65535
        _ = v9
    default:
        v10 := GenS2{F0: 619}
        _ = v10
    }
    var v11 uint = 4294967295
    _ = v11
    return -529
}
func genFn1(p0 uint16, p1 int) int32 {
    fmt.Println(p1, p1)
    if true {
        if false {
            switch "hello" {
            case "a":
                p1 += (p1 % ((128) | 1))
            case "hello":
                v12 := GenS3{F0: -1}
                _ = v12
            default:
                var v13 int16 = -32768
                _ = v13
            }
        } else {
            var v14 uint8 = 255
            _ = v14
        }
        v15 := []int{(p1 | (p1 << ((42) & 63)))}
        _ = v15
    } else {
        var v16 int = (p1 ^ p1)
        switch (v16 ^ p1) {
        case -9223372036854775808:
            var v17 int64 = genFn0()
            _ = v17
        case 127:
            fmt.Println(v16)
        default:
            v18 := GenS4{F0: (0 > 42), F1: 9223372036854775807, F2: 2}
            _ = v18
        }
        for fi19 := 0; fi19 < 16; fi19++ {
            if (0 <= 100) {
                p0 -= (p0 % (((^(p0))) | 1))
                var v20 float32 = float32(float64(0.0))
                v20 -= (((v20 + 0.0) / v20) * 3.14159)
                _ = v20
            }
            _ = fi19
        }
        _ = v16
    }
    _ = p0
    _ = p1
    return 2147483647
}

func main() {
    var v1 uint8 = 232
    var v2 string = "foo bar"
    var v3 float64 = (0.1 - 3.14159)
    var v21 int32 = 1
    v1 = (v1 * v1)
    if false {
        v22 := GenS2{F0: 42}
        if true {
            if true {
                switch "\n\t" {
                case "":
                    v21 += v21
                }
            } else {
                for fi23 := 0; fi23 < 3; fi23++ {
                    var v24 uint64 = 42
                    var v25 uint64 = v24
                    v26 := []int{(^(fi23))}
                    _ = fi23
                    _ = v24
                    _ = v25
                    _ = v26
                }
            }
            v3 -= (v3 * float64((((3.14159 * -2.25) + float32(1.5)) + 1.5)))
            switch "hello" {
            case "hello":
                v27 := GenS5{F0: 1000, F1: (-(v21))}
                _ = v27
            case "\n\t":
                var v28 float64 = v3
                _ = v28
            case "":
                var v29 uint8 = v1
                _ = v29
            default:
                var v30 float64 = 1e10
                _ = v30
            }
        }
        v1 = (v1 / ((v1) | 1))
        _ = v22
    } else {
        var v31 string = v2
        for fi32 := 0; fi32 < 19; fi32++ {
            if false {
                var v33 int8 = 0
                for fi34 := 0; fi34 < 20; fi34++ {
                    v33 = v33
                    if false {
                        fmt.Println(v21, fi32)
                        var v35 string = (v2 + ((v2 + (v2 + v2)) + v2))
                        switch "foo bar" {
                        case "\n\t":
                            v31 = "foo bar"
                        }
                        _ = v35
                    }
                    _ = fi34
                }
                _ = v33
            }
            v36 := GenS5{F0: 1000, F1: v21}
            v37 := []int{(^(fi32)), (fi32 % ((4294967295) | 1))}
            _ = fi32
            _ = v36
            _ = v37
        }
        v38 := []int{128, 503, 828, 1}
        _ = v31
        _ = v38
    }
    switch -9223372036854775808 {
    case -725:
        var v39 int16 = 1
        _ = v39
    case 32767:
        switch 130 {
        case -820:
            if false {
                v3 = 1.5
                v40 := GenS6{F0: 42, F1: 128, F2: 1}
                _ = v40
            } else {
                var v41 bool = (0 >= 616)
                _ = v41
            }
        case -287:
            for fi42 := 0; fi42 < 4; fi42++ {
                v43 := GenS7{F0: 2147483648, F1: (65535 >= 65535)}
                for fi44 := 0; fi44 < 17; fi44++ {
                    v21 -= genFn1(18, fi44)
                    _ = fi44
                }
                if false {
                    switch fi42 {
                    case 4294967295:
                        v21 = -1
                    case 2147483647:
                        var v45 uint64 = 18446744073709551615
                        _ = v45
                    case 669:
                        if ((2147483647 != 1) || true) {
                            if (true && (9223372036854775807 != 42)) {
                                v21 += (-(v21))
                            } else {
                                v1 = (v1 & 0)
                            }
                            v3 -= v3
                            var v46 bool = false
                            _ = v46
                        }
                    }
                    switch (fi42 + fi42) {
                    case 2147483647:
                        var v47 int32 = (v21 - 42)
                        _ = v47
                    case -1:
                        var v48 uint32 = 2
                        _ = v48
                    }
                }
                _ = fi42
                _ = v43
            }
        default:
            switch v2 {
            case "a":
                switch v2 {
                case "hello":
                    if (false && true) {
                        var v49 float32 = 1.5
                        v2 = v2
                        _ = v49
                    }
                case "foo bar":
                    switch -2147483648 {
                    case 32767:
                        switch -281 {
                        case 0:
                            var v50 uint = 4294967295
                            _ = v50
                        case -144:
                            v51 := GenS1{F0: 200, F1: 0}
                            _ = v51
                        case 1:
                            if true {
                                v3 = -2.25
                                fmt.Println(v21)
                                for fi52 := 0; fi52 < 20; fi52++ {
                                    v53 := []int{fi52, 0, (fi52 * 1)}
                                    _ = fi52
                                    _ = v53
                                }
                            }
                        }
                    default:
                        if true {
                            for fi54 := 0; fi54 < 19; fi54++ {
                                switch 32768 {
                                case -9223372036854775808:
                                    for fi55 := 0; fi55 < 19; fi55++ {
                                        fmt.Println(fi54, v1, v2)
                                        if (false || false) {
                                            var v56 uint8 = v1
                                            v21 = genFn1(285, (fi55 & fi54))
                                            v21 = 992
                                            var v57 int8 = -1
                                            _ = v56
                                            _ = v57
                                        } else {
                                            var v58 float32 = 1e10
                                            _ = v58
                                        }
                                        switch (((v2 + "foo bar") + v2) + ("foo bar" + v2)) {
                                        case "a":
                                            if true {
                                                v59 := []int{(fi55 | fi54), (fi55 << ((42) & 63)), fi54}
                                                v2 = v2
                                                var v60 float32 = (3.14159 / (3.14159 / 0.1))
                                                _ = v59
                                                _ = v60
                                            } else {
                                                v61 := GenS3{F0: 42}
                                                v62 := GenS4{F0: true, F1: -9223372036854775808, F2: 255}
                                                switch "a" {
                                                case "":
                                                    if false {
                                                        v21 -= (v21 | int32(fi54))
                                                    } else {
                                                        v21 = v21
                                                    }
                                                case "hello":
                                                    if true {
                                                        v21 -= (v21 | v21)
                                                        v21 -= -678
                                                    } else {
                                                        if ((!true) || false) {
                                                            switch v2 {
                                                            case "\n\t":
                                                                v63 := GenS7{F0: 2147483647, F1: (v3 > v3)}
                                                                _ = v63
                                                            case "foo bar":
                                                                var v64 uint64 = 887
                                                                v65 := GenS8{F0: (fi54 >> ((284) & 63)), F1: (121 != -60)}
                                                                v66 := []int{(^(fi55)), (fi55 << ((0) & 63))}
                                                                _ = v64
                                                                _ = v65
                                                                _ = v66
                                                            }
                                                            if true {
                                                                var v67 float64 = v3
                                                                if (false && true) {
                                                                    var v68 uint64 = 4294967295
                                                                    if (v1 > 128) {
                                                                        v69 := []int{}
                                                                        v2 = v2
                                                                        if false {
                                                                            if (true || true) {
                                                                                if true {
                                                                                    v68 = v68
                                                                                    v70 := GenS9{F0: (v1 << ((18446744073709551615) & 63)), F1: true}
                                                                                    v1 += (v1 % ((91) | 1))
                                                                                    _ = v70
                                                                                } else {
                                                                                    switch fi55 {
                                                                                    case 127:
                                                                                        if (!true) {
                                                                                            v67 = 1e10
                                                                                            var v71 int8 = -54
                                                                                            v67 += (v3 / v67)
                                                                                            var v72 uint16 = 0
                                                                                            var v73 int16 = 1000
                                                                                            _ = v71
                                                                                            _ = v72
                                                                                            _ = v73
                                                                                        }
                                                                                    default:
                                                                                        v67 = ((1.5 - (((0.1 / v67) / 0.1) + (float64(3.14159) / float64(((-2.25 + float32((float64(0.0) / v3))) * (3.14159 + 0.0)))))) * v67)
                                                                                    }
                                                                                    v68 += (v68 - v68)
                                                                                }
                                                                                switch "" {
                                                                                case "\n\t":
                                                                                    switch (((v2 + v2) + "hello") + (("hello" + (v2 + v2)) + ((v2 + ("" + "")) + v2))) {
                                                                                    case "foo bar":
                                                                                        switch (fi54 | fi55) {
                                                                                        case 416:
                                                                                            var v74 bool = (255 >= 1)
                                                                                            _ = v74
                                                                                        case 32768:
                                                                                            var v75 string = "\n\t"
                                                                                            _ = v75
                                                                                        case 128:
                                                                                            var v76 int32 = genFn1(uint16(v68), 42)
                                                                                            _ = v76
                                                                                        default:
                                                                                            if (true && false) {
                                                                                                v68 += v68
                                                                                            } else {
                                                                                                v77 := GenS10{F0: 65535}
                                                                                                _ = v77
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                case "hello":
                                                                                    var v78 int32 = v21
                                                                                    _ = v78
                                                                                }
                                                                            } else {
                                                                                var v79 uint16 = 65535
                                                                                var v80 float32 = (1e10 * 1e10)
                                                                                var v81 uint8 = (v1 * v1)
                                                                                _ = v79
                                                                                _ = v80
                                                                                _ = v81
                                                                            }
                                                                            if len(v69) > 0 {
                                                                                v69[0] = (^(fi55))
                                                                            }
                                                                        } else {
                                                                            var v82 uint = 377
                                                                            _ = v82
                                                                        }
                                                                        _ = v69
                                                                    } else {
                                                                        var v83 uint64 = 763
                                                                        v84 := []int{(fi55 / ((fi55) | 1)), 127, -201}
                                                                        _ = v83
                                                                        _ = v84
                                                                    }
                                                                    _ = v68
                                                                } else {
                                                                    var v85 int8 = 100
                                                                    _ = v85
                                                                }
                                                                if (512 > 1) {
                                                                    switch fi55 {
                                                                    case 2147483647:
                                                                        var v86 uint8 = (v1 + v1)
                                                                        _ = v86
                                                                    case -1:
                                                                        fmt.Println(fi54)
                                                                    default:
                                                                        v3 = (((v67 - 1e10) * (v3 / v67)) + float64(1e10))
                                                                    }
                                                                    v87 := []int{-9223372036854775808}
                                                                    if true {
                                                                        v2 = "a"
                                                                        var v88 bool = false
                                                                        v89 := GenS11{F0: false, F1: v1, F2: 255}
                                                                        _ = v88
                                                                        _ = v89
                                                                    }
                                                                    var v90 uint64 = 42
                                                                    v91 := GenS12{F0: v90, F1: ("\n\t" + (v2 + v2)), F2: -777}
                                                                    v2 = v2
                                                                    switch v2 {
                                                                    case "hello":
                                                                        var v92 uint = 65535
                                                                        _ = v92
                                                                    default:
                                                                        _ = len(v87)
                                                                    }
                                                                    if (("foo bar" + v2) < v2) {
                                                                        fmt.Println(v67, v90)
                                                                    }
                                                                    var v93 uint64 = v90
                                                                    _ = v87
                                                                    _ = v90
                                                                    _ = v91
                                                                    _ = v93
                                                                }
                                                                if (fi54 != fi55) {
                                                                    var v94 string = (("a" + v2) + (v2 + v2))
                                                                    v21 -= v21
                                                                    v95 := []int{fi55, (fi55 >> ((0) & 63)), -651, -1}
                                                                    _ = v94
                                                                    _ = v95
                                                                }
                                                                if (112 <= 168) {
                                                                    v96 := []int{}
                                                                    _ = v96
                                                                }
                                                                v97 := GenS13{F0: 4294967295}
                                                                v98 := []int{(fi55 >> ((680) & 63))}
                                                                _ = v67
                                                                _ = v97
                                                                _ = v98
                                                            } else {
                                                                v99 := []int{}
                                                                var v100 uint = 65535
                                                                _ = v100
                                                                _ = v99
                                                            }
                                                        }
                                                        if false {
                                                            var v101 uint64 = 2
                                                            v1 = v1
                                                            _ = v101
                                                        }
                                                    }
                                                case "foo bar":
                                                    switch "a" {
                                                    case "a":
                                                        var v102 uint8 = (^(v1))
                                                        _ = v102
                                                    case "\n\t":
                                                        var v103 int64 = 9223372036854775807
                                                        fmt.Println(v21, v3, v2)
                                                        _ = v103
                                                    case "hello":
                                                        var v104 uint64 = 512
                                                        _ = v104
                                                    default:
                                                        if false {
                                                            if (!true) {
                                                                v105 := GenS11{F0: (fi54 >= (fi55 % ((42) | 1))), F1: 0, F2: 61}
                                                                _ = v105
                                                            }
                                                            var v106 int = 0
                                                            _ = v106
                                                        } else {
                                                            v107 := []int{(fi54 * (fi55 ^ 127))}
                                                            if false {
                                                                switch -1 {
                                                                case 9223372036854775807:
                                                                    var v108 uint = 18446744073709551615
                                                                    _ = v108
                                                                case 128:
                                                                    var v109 uint16 = 1
                                                                    _ = v109
                                                                }
                                                            } else {
                                                                fmt.Println(v3, fi55)
                                                                var v110 float64 = (((v3 * float64(1e10)) / float64(1.5)) - (v3 + v3))
                                                                _ = v110
                                                            }
                                                            _ = v107
                                                        }
                                                    }
                                                default:
                                                    var v111 int = (fi54 + 42)
                                                    _ = v111
                                                }
                                                _ = v61
                                                _ = v62
                                            }
                                            v112 := []int{}
                                            if (!((4294967295 > 65535) && (-32768 == 987))) {
                                                v1 = v1
                                                _ = len(v112)
                                                var v113 string = ("\n\t" + "hello")
                                                _ = v113
                                            }
                                            _ = v112
                                        case "foo bar":
                                            v3 = float64(((-2.25 + 0.0) * (3.14159 * float32(v3))))
                                        }
                                        _ = fi55
                                    }
                                default:
                                    v114 := []int{fi54}
                                    _ = v114
                                }
                                switch 85 {
                                case 32768:
                                    var v115 int32 = v21
                                    _ = v115
                                }
                                _ = fi54
                            }
                            v2 = v2
                            v1 = 60
                        } else {
                            v2 = v2
                        }
                    }
                }
            case "foo bar":
                var v116 string = (("foo bar" + v2) + "a")
                _ = v116
            }
        }
    case 128:
        var v117 uint8 = 56
        _ = v117
    }
    if (42 < 1) {
        for fi118 := 0; fi118 < 8; fi118++ {
            for fi119 := 0; fi119 < 19; fi119++ {
                if false {
                    v120 := GenS14{F0: 0}
                    var v121 int8 = -128
                    _ = v120
                    _ = v121
                }
                _ = fi119
            }
            fmt.Println(v2, v3)
            var v122 int = fi118
            _ = fi118
            _ = v122
        }
    } else {
        v1 -= (v1 ^ v1)
        v123 := GenS1{F0: 586, F1: 4294967295}
        fmt.Println(v1, v2, v1)
        _ = v123
    }
    var v124 int = 42
    for fi125 := 0; fi125 < 2; fi125++ {
        if (19 != (v124 - (v124 % ((v124) | 1)))) {
            v126 := []int{fi125, 42, v124, (v124 >> ((255) & 63))}
            _ = v126
        } else {
            v127 := GenS4{F0: (0 >= 717), F1: 1, F2: 18446744073709551615}
            v128 := GenS4{F0: true, F1: 369, F2: 255}
            _ = v127
            _ = v128
        }
        if (false || (0.0 >= (-2.25 / (1.5 + 3.14159)))) {
            var v129 int16 = 1
            _ = v129
        } else {
            v3 = 3.14159
            v130 := GenS15{F0: false, F1: 603}
            var v131 uint32 = 35
            _ = v130
            _ = v131
        }
        v21 = v21
        _ = fi125
    }
    var v132 uint16 = 851
    var v133 int8 = -91
    fmt.Println(v1)
    v133 = -1
    switch -9223372036854775808 {
    case 209:
        var v134 int = (v124 ^ v124)
        _ = v134
    case -1:
        var v135 uint8 = v1
        _ = v135
    default:
        for fi136 := 0; fi136 < 5; fi136++ {
            v1 = v1
            switch v124 {
            case 127:
                v137 := []int{(v124 * (fi136 << ((653) & 63))), v124}
                _ = v137
            case 9223372036854775807:
                v132 += (v132 * 1)
            }
            _ = fi136
        }
    }
    if (!((-128 > v133) && false)) {
        v3 = v3
        var v138 int64 = -1
        _ = v138
    } else {
        v139 := []int{(-(v124)), v124, -9223372036854775808}
        _ = v139
    }
    if false {
        switch (v124 % ((391) | 1)) {
        case 4294967295:
            if (v21 >= -1) {
                var v140 float32 = 1.5
                _ = v140
            } else {
                v21 += (^(v21))
                var v141 int = (v124 / (((v124 >> ((uint(v1)) & 63))) | 1))
                _ = v141
            }
        case 0:
            v142 := []int{(v124 * v124), (v124 ^ (v124 * (v124 | (v124 % (((v124 & (^(v124)))) | 1))))), 32767, (v124 % ((127) | 1))}
            _ = v142
        case 128:
            v2 = v2
        }
    } else {
        if false {
            if ((v124 + 72) != v124) {
                v2 = v2
            } else {
                if (true && false) {
                    if true {
                        if true {
                            for fi143 := 0; fi143 < 14; fi143++ {
                                v144 := []int{fi143, v124}
                                if len(v144) >= 2 {
                                    v144 = v144[1:len(v144)]
                                }
                                _ = fi143
                                _ = v144
                            }
                            for fi145 := 0; fi145 < 9; fi145++ {
                                v1 = v1
                                _ = fi145
                            }
                        } else {
                            v2 = v2
                        }
                        for fi146 := 0; fi146 < 4; fi146++ {
                            for fi147 := 0; fi147 < 12; fi147++ {
                                switch fi146 {
                                case 32768:
                                    switch v124 {
                                    case 0:
                                        v1 += 255
                                        var v148 bool = true
                                        if (float32((float64(3.14159) + float64((1.5 + 1.5)))) < (float32(float64(3.14159)) * 1.5)) {
                                            v1 = (v1 * (v1 / (((v1 ^ v1)) | 1)))
                                            v133 = v133
                                            switch "a" {
                                            case "foo bar":
                                                v149 := []int{(fi147 - -9223372036854775808)}
                                                _ = v149
                                            case "a":
                                                v132 -= 32767
                                            case "hello":
                                                v150 := []int{(fi147 & fi146), fi147}
                                                _ = v150
                                            }
                                        }
                                        _ = v148
                                    case 128:
                                        var v151 bool = true
                                        _ = v151
                                    }
                                case 706:
                                    v152 := GenS7{F0: 2147483648, F1: true}
                                    _ = v152
                                case 127:
                                    var v153 int16 = -179
                                    _ = v153
                                default:
                                    if (42 < 42) {
                                        v154 := []int{(v124 - (^(v124)))}
                                        switch "hello" {
                                        case "\n\t":
                                            if true {
                                                var v155 float64 = v3
                                                switch fi147 {
                                                case 9223372036854775807:
                                                    v132 = v132
                                                }
                                                v156 := GenS16{F0: (true && (((!true) || true) || true)), F1: v2, F2: v21}
                                                v21 -= 2147483647
                                                var v157 float64 = v3
                                                _ = v155
                                                _ = v156
                                                _ = v157
                                            }
                                        case "":
                                            var v158 int = (-(v124))
                                            _ = v158
                                        case "hello":
                                            v132 = v132
                                        }
                                        var v159 uint = 2
                                        _ = v154
                                        _ = v159
                                    }
                                }
                                _ = fi147
                            }
                            if true {
                                v160 := []int{(fi146 | fi146), -878, v124, (v124 / ((fi146) | 1))}
                                _ = v160
                            }
                            if (255 <= 0) {
                                for fi161 := 0; fi161 < 16; fi161++ {
                                    if false {
                                        v124 += (v124 | (fi161 ^ v124))
                                    } else {
                                        fmt.Println(v2, fi161)
                                    }
                                    v162 := GenS2{F0: 1}
                                    if true {
                                        var v163 uint8 = 127
                                        v164 := GenS17{F0: 127, F1: false}
                                        v165 := GenS8{F0: v124, F1: true}
                                        _ = v163
                                        _ = v164
                                        _ = v165
                                    }
                                    _ = fi161
                                    _ = v162
                                }
                                v3 -= -2.25
                            }
                            _ = fi146
                        }
                        for fi166 := 0; fi166 < 16; fi166++ {
                            v167 := GenS18{F0: v133, F1: (v1 * 0)}
                            v168 := GenS19{F0: 255, F1: v124}
                            var v169 uint64 = 42
                            _ = fi166
                            _ = v167
                            _ = v168
                            _ = v169
                        }
                    } else {
                        for fi170 := 0; fi170 < 19; fi170++ {
                            fmt.Println(v3, v133)
                            if false {
                                if ((v1 <= v1) || false) {
                                    v21 -= -2147483648
                                    v171 := GenS11{F0: (true && ((^(v133)) > (v133 ^ (^(v133))))), F1: v1, F2: 989}
                                    switch v124 {
                                    case 1:
                                        v132 += v132
                                    case 32768:
                                        v1 = 185
                                    case 9223372036854775807:
                                        if false {
                                            switch (fi170 % ((fi170) | 1)) {
                                            case -568:
                                                v1 = (^(v1))
                                            }
                                            for fi172 := 0; fi172 < 9; fi172++ {
                                                var v173 float32 = float32((v3 + 1.5))
                                                var v174 uint16 = (v132 - (v132 % (((v132 - (v132 % (((^(v132))) | 1)))) | 1)))
                                                _ = fi172
                                                _ = v173
                                                _ = v174
                                            }
                                        } else {
                                            if (42 >= 1) {
                                                v175 := GenS20{F0: v21, F1: "a"}
                                                for fi176 := 0; fi176 < 17; fi176++ {
                                                    v177 := GenS21{F0: (v1 / ((v1) | 1)), F1: (v133 * v133), F2: 255}
                                                    _ = fi176
                                                    _ = v177
                                                }
                                                _ = v175
                                            } else {
                                                var v178 string = ""
                                                for fi179 := 0; fi179 < 13; fi179++ {
                                                    fmt.Println(v1, v132)
                                                    v180 := GenS22{F0: (v133 ^ v133), F1: 255, F2: 1}
                                                    v178 = v2
                                                    _ = fi179
                                                    _ = v180
                                                }
                                                switch v178 {
                                                case "":
                                                    v181 := []int{fi170, -329}
                                                    _ = v181
                                                case "a":
                                                    fmt.Println(v124)
                                                case "foo bar":
                                                    v182 := GenS23{F0: int64(v1), F1: -1}
                                                    _ = v182
                                                default:
                                                    for fi183 := 0; fi183 < 8; fi183++ {
                                                        var v184 int8 = (v133 << ((699) & 63))
                                                        var v185 uint = 0
                                                        _ = fi183
                                                        _ = v184
                                                        _ = v185
                                                    }
                                                }
                                                _ = v178
                                            }
                                        }
                                    default:
                                        if true {
                                            v186 := GenS19{F0: 813, F1: (v124 & 128)}
                                            _ = v186
                                        } else {
                                            if (!false) {
                                                var v187 int64 = 0
                                                v188 := []int{}
                                                _ = v187
                                                _ = v188
                                            } else {
                                                v21 += (v21 / (((v21 | 638)) | 1))
                                            }
                                        }
                                    }
                                    _ = v171
                                } else {
                                    v189 := []int{}
                                    v190 := GenS24{F0: 2147483647}
                                    v191 := GenS6{F0: 2, F1: fi170, F2: 805}
                                    _ = v189
                                    _ = v190
                                    _ = v191
                                }
                                for fi192 := 0; fi192 < 11; fi192++ {
                                    v193 := GenS8{F0: (fi192 & v124), F1: (false && true)}
                                    switch fi192 {
                                    case -1:
                                        switch (v124 % (((v124 & fi192)) | 1)) {
                                        case 127:
                                            switch v2 {
                                            case "":
                                                v194 := []int{(^(fi170)), (^(v124))}
                                                _ = v194
                                            case "hello":
                                                if (true && true) {
                                                    if false {
                                                        v124 -= fi192
                                                        if true {
                                                            var v195 int32 = 0
                                                            switch 23 {
                                                            case -927:
                                                                v2 = ""
                                                            case 0:
                                                                if false {
                                                                    v196 := GenS20{F0: v195, F1: v2}
                                                                    v195 += v21
                                                                    _ = v196
                                                                }
                                                            case 4294967295:
                                                                v133 = (v133 * (v133 ^ (^(v133))))
                                                            }
                                                            _ = v195
                                                        }
                                                        var v197 int64 = genFn0()
                                                        _ = v197
                                                    }
                                                    var v198 bool = true
                                                    _ = v198
                                                } else {
                                                    var v199 int64 = 0
                                                    v21 -= (v21 * (v21 - v21))
                                                    _ = v199
                                                }
                                                v21 += v21
                                                if false {
                                                    v200 := GenS12{F0: 934, F1: (("\n\t" + (v2 + "\n\t")) + (v2 + ("foo bar" + ""))), F2: -1}
                                                    var v201 uint64 = 1
                                                    v124 = -2147483648
                                                    v21 -= 0
                                                    v202 := []int{-798, 32768, 2147483647}
                                                    _ = v200
                                                    _ = v201
                                                    _ = v202
                                                }
                                            default:
                                                var v203 int64 = 1
                                                _ = v203
                                            }
                                            switch fi170 {
                                            case 1:
                                                v21 = (v21 % ((int32(v132)) | 1))
                                            }
                                            switch fi170 {
                                            case 4294967295:
                                                if (!true) {
                                                    fmt.Println(v132, v132, v2)
                                                } else {
                                                    fmt.Println(v1)
                                                    var v204 string = v2
                                                    _ = v204
                                                }
                                            case -1:
                                                v205 := GenS5{F0: 1000, F1: (v21 + v21)}
                                                _ = v205
                                            case -2147483648:
                                                switch (v2 + v2) {
                                                case "":
                                                    var v206 int32 = 1
                                                    _ = v206
                                                default:
                                                    var v207 float32 = (0.1 / float32((0.1 * v3)))
                                                    _ = v207
                                                }
                                            default:
                                                v124 += (fi192 * 32767)
                                            }
                                        case -81:
                                            var v208 uint64 = 0
                                            _ = v208
                                        default:
                                            if ("\n\t" != v2) {
                                                v209 := []int{(-(v124)), 492, (v124 >> ((42) & 63)), v124}
                                                switch "a" {
                                                case "":
                                                    var v210 string = (((v2 + "a") + ((v2 + "") + "a")) + "hello")
                                                    _ = v210
                                                case "hello":
                                                    v2 = ((v2 + v2) + (v2 + v2))
                                                }
                                                _ = len(v209)
                                                v211 := GenS18{F0: (v133 >> ((733) & 63)), F1: 1}
                                                _ = v209
                                                _ = v211
                                            }
                                        }
                                    case 1:
                                        v132 += 2
                                    case -564:
                                        v212 := GenS19{F0: 18446744073709551615, F1: (v124 / ((fi192) | 1))}
                                        _ = v212
                                    }
                                    switch "a" {
                                    case "\n\t":
                                        var v213 uint8 = 128
                                        _ = v213
                                    }
                                    var v214 uint64 = 65
                                    v215 := GenS16{F0: (true && (!((!(v2 > (v2 + ""))) || (uint(v214) == 744)))), F1: "hello", F2: (v21 % ((-1) | 1))}
                                    if true {
                                        v216 := GenS12{F0: v214, F1: "", F2: -617}
                                        _ = v216
                                    } else {
                                        v1 = (v1 + 128)
                                        v124 += fi192
                                    }
                                    if (!false) {
                                        v2 = "a"
                                        v217 := []int{(fi192 >> ((807) & 63)), (fi192 % ((fi170) | 1)), 127, (v124 + (v124 ^ int(v133)))}
                                        var v218 uint8 = (^(v1))
                                        _ = v217
                                        _ = v218
                                    }
                                    var v219 int8 = (v133 >> ((4294967295) & 63))
                                    if (((v2 + v2) != ("foo bar" + v2)) && true) {
                                        v1 = v1
                                        var v220 int32 = genFn1(v132, (fi170 * 128))
                                        var v221 int8 = -128
                                        _ = v220
                                        _ = v221
                                    } else {
                                        var v222 int64 = -9223372036854775808
                                        v222 -= genFn0()
                                        v2 = v2
                                        _ = v222
                                    }
                                    var v223 uint8 = (^(v1))
                                    var v224 uint = 4294967295
                                    _ = fi192
                                    _ = v193
                                    _ = v214
                                    _ = v215
                                    _ = v219
                                    _ = v223
                                    _ = v224
                                }
                            }
                            _ = fi170
                        }
                        for fi225 := 0; fi225 < 19; fi225++ {
                            if false {
                                v133 -= v133
                                v226 := []int{(fi225 | (v124 << ((269) & 63))), v124, -2147483648}
                                for fi227 := 0; fi227 < 12; fi227++ {
                                    v1 -= 85
                                    _ = fi227
                                }
                                _ = v226
                            } else {
                                for fi228 := 0; fi228 < 17; fi228++ {
                                    var v229 int8 = (v133 << ((1) & 63))
                                    _ = fi228
                                    _ = v229
                                }
                                if true {
                                    switch v2 {
                                    case "hello":
                                        var v230 uint16 = v132
                                        _ = v230
                                    case "a":
                                        if (!false) {
                                            if (!false) {
                                                var v231 int16 = 175
                                                _ = v231
                                            }
                                            v133 = -14
                                            v232 := GenS17{F0: (v133 + v133), F1: false}
                                            _ = v232
                                        }
                                    }
                                    v133 -= v133
                                }
                            }
                            var v233 uint64 = 545
                            _ = fi225
                            _ = v233
                        }
                        var v234 float64 = ((((v3 + 1.5) - ((v3 / v3) - (v3 / float64(1e10)))) + 0.0) + v3)
                        _ = v234
                    }
                }
            }
        }
        switch (v124 & 648) {
        case 2147483647:
            if false {
                if ((!(uint(v132) <= 4294967295)) || (genFn0() < -290)) {
                    v235 := []int{127}
                    _ = v235
                } else {
                    v236 := []int{(v124 % ((712) | 1)), (^(v124)), v124}
                    for fi237 := 0; fi237 < 11; fi237++ {
                        for fi238 := 0; fi238 < 14; fi238++ {
                            var v239 float32 = float32(((v3 * ((v3 * (v3 * v3)) - ((v3 * v3) / (float64(((3.14159 / 1.5) + float32(1e10))) * (v3 / ((1.5 + v3) / v3)))))) / -2.25))
                            if len(v236) >= 2 {
                                v236 = v236[1:len(v236)]
                            }
                            v239 = v239
                            _ = fi238
                            _ = v239
                        }
                        v21 -= (^(v21))
                        for fi240 := 0; fi240 < 16; fi240++ {
                            v133 += v133
                            if true {
                                var v241 uint32 = 65535
                                if ((v1 & v1) <= v1) {
                                    if true {
                                        fmt.Println(v1, v124)
                                        fmt.Println(v132, v3)
                                        v133 -= v133
                                    }
                                    switch "\n\t" {
                                    case "\n\t":
                                        var v242 float64 = float64(0.0)
                                        _ = v242
                                    case "a":
                                        if len(v236) > 0 {
                                            v236[0] = (fi237 & fi240)
                                        }
                                        v132 -= 581
                                        fmt.Println(v2, v133)
                                    case "hello":
                                        v1 -= (v1 / ((v1) | 1))
                                    }
                                } else {
                                    var v243 uint = 1
                                    _ = v243
                                }
                                _ = v241
                            }
                            _ = fi240
                        }
                        _ = fi237
                    }
                    var v244 bool = true
                    _ = v236
                    _ = v244
                }
            }
        case -9223372036854775808:
            if (0 <= genFn1(v132, v124)) {
                fmt.Println(v133, v2)
                var v245 float64 = 3.14159
                _ = v245
            } else {
                if false {
                    var v246 uint16 = v132
                    var v247 float64 = 0.1
                    var v248 int16 = 229
                    _ = v246
                    _ = v247
                    _ = v248
                } else {
                    var v249 uint = 65535
                    v250 := []int{(-(v124)), v124}
                    v249 += (v249 >> ((1) & 63))
                    _ = v249
                    _ = v250
                }
            }
        case 1:
            v251 := []int{(v124 / ((-163) | 1)), v124}
            _ = v251
        }
        if false {
            switch v124 {
            case 1:
                switch (v124 - (v124 | v124)) {
                case -9223372036854775808:
                    var v252 uint32 = 4294967295
                    _ = v252
                }
            }
            v253 := GenS25{F0: v1, F1: v1}
            _ = v253
        } else {
            for fi254 := 0; fi254 < 19; fi254++ {
                for fi255 := 0; fi255 < 5; fi255++ {
                    var v256 int32 = (-(v21))
                    var v257 uint = 704
                    _ = fi255
                    _ = v256
                    _ = v257
                }
                if false {
                    switch 544 {
                    case -2147483648:
                        for fi258 := 0; fi258 < 6; fi258++ {
                            v259 := GenS18{F0: (v133 << ((18446744073709551615) & 63)), F1: v1}
                            _ = fi258
                            _ = v259
                        }
                    case 4294967295:
                        if (int8(v1) < v133) {
                            for fi260 := 0; fi260 < 17; fi260++ {
                                v3 += (float64(3.14159) + (v3 * 0.0))
                                _ = fi260
                            }
                            if true {
                                var v261 int32 = genFn1(v132, fi254)
                                v262 := []int{v124, v124, (v124 - (fi254 % ((v124) | 1))), fi254}
                                if true {
                                    if (1 > v261) {
                                        v263 := GenS13{F0: 2147483648}
                                        v132 = 2
                                        _ = v263
                                    }
                                    switch 4294967295 {
                                    case -359:
                                        v264 := GenS26{F0: genFn0(), F1: -32768}
                                        _ = v264
                                    case 406:
                                        if false {
                                            if (v132 < (v132 ^ v132)) {
                                                switch (^(fi254)) {
                                                case 32768:
                                                    v265 := GenS27{F0: 1, F1: 104, F2: 4294967295}
                                                    _ = v265
                                                case 645:
                                                    if ((v132 * 32767) >= 747) {
                                                        switch fi254 {
                                                        case -2147483648:
                                                            _ = len(v262)
                                                        case 112:
                                                            for fi266 := 0; fi266 < 5; fi266++ {
                                                                v3 += v3
                                                                _ = fi266
                                                            }
                                                        }
                                                    } else {
                                                        for fi267 := 0; fi267 < 13; fi267++ {
                                                            if false {
                                                                v268 := GenS1{F0: 0, F1: 4294967295}
                                                                _ = v268
                                                            } else {
                                                                v2 = v2
                                                                v133 = v133
                                                            }
                                                            _ = fi267
                                                        }
                                                        v124 += 2147483647
                                                        for fi269 := 0; fi269 < 20; fi269++ {
                                                            var v270 uint = 4294967295
                                                            if len(v262) >= 2 {
                                                                v262 = v262[1:len(v262)]
                                                            }
                                                            _ = fi269
                                                            _ = v270
                                                        }
                                                    }
                                                }
                                                if len(v262) >= 2 {
                                                    v262 = v262[1:len(v262)]
                                                }
                                            } else {
                                                fmt.Println(v124)
                                                switch fi254 {
                                                case 2147483647:
                                                    var v271 string = ""
                                                    _ = v271
                                                case 4294967295:
                                                    if ((((3.14159 * (((float64(0.0) + v3) * (v3 / ((v3 * v3) + 0.0))) + v3)) - float64(0.1)) + v3) > 3.14159) {
                                                        var v272 uint64 = 354
                                                        fmt.Println(v21, v133, fi254)
                                                        _ = v272
                                                    } else {
                                                        switch v2 {
                                                        case "\n\t":
                                                            if (0 > uint64(v1)) {
                                                                if ((4294967295 >= 65535) || false) {
                                                                    var v273 int8 = 127
                                                                    var v274 float64 = ((v3 - v3) * v3)
                                                                    _ = v273
                                                                    _ = v274
                                                                } else {
                                                                    v275 := GenS2{F0: 4294967295}
                                                                    _ = v275
                                                                }
                                                            }
                                                        }
                                                        if (!true) {
                                                            if false {
                                                                var v276 int32 = (v261 >> ((18446744073709551615) & 63))
                                                                var v277 bool = (1e10 > (float32(float64(-2.25)) - 0.1))
                                                                if len(v262) > 0 {
                                                                    v262[0] = -9223372036854775808
                                                                }
                                                                _ = v276
                                                                _ = v277
                                                            } else {
                                                                for fi278 := 0; fi278 < 8; fi278++ {
                                                                    if true {
                                                                        if (!(!true)) {
                                                                            if (!(((127 >= 0) || (!(255 < 255))) && false)) {
                                                                                var v279 uint64 = 4294967295
                                                                                if (0.0 <= v3) {
                                                                                    _ = len(v262)
                                                                                    v124 += fi278
                                                                                    if len(v262) >= 2 {
                                                                                        v262 = v262[1:len(v262)]
                                                                                    }
                                                                                    fmt.Println(v21, v133, v133)
                                                                                    switch fi278 {
                                                                                    case -9223372036854775808:
                                                                                        v132 -= (v132 & 326)
                                                                                    default:
                                                                                        v280 := GenS27{F0: (-(fi278)), F1: 1, F2: (v279 + v279)}
                                                                                        _ = v280
                                                                                    }
                                                                                }
                                                                                var v281 bool = true
                                                                                fmt.Println(v132)
                                                                                var v282 int = (fi278 << ((495) & 63))
                                                                                switch v2 {
                                                                                case "hello":
                                                                                    v132 -= 32767
                                                                                }
                                                                                _ = v279
                                                                                _ = v281
                                                                                _ = v282
                                                                            }
                                                                            v283 := GenS8{F0: (^(fi278)), F1: false}
                                                                            _ = v283
                                                                        }
                                                                    } else {
                                                                        v262 = append(v262, (fi254 % ((fi254) | 1)))
                                                                        var v284 float64 = (v3 - v3)
                                                                        _ = v284
                                                                    }
                                                                    _ = fi278
                                                                }
                                                                if len(v262) > 0 {
                                                                    v262[0] = 2147483647
                                                                }
                                                                v285 := GenS28{F0: -1, F1: (v132 + (v132 ^ (v132 | v132)))}
                                                                _ = v285
                                                            }
                                                        } else {
                                                            v286 := GenS29{F0: (!(true || false)), F1: (v1 << ((4294967295) & 63))}
                                                            fmt.Println(v132)
                                                            switch v124 {
                                                            case 623:
                                                                for fi287 := 0; fi287 < 8; fi287++ {
                                                                    v288 := GenS1{F0: 42, F1: 4294967295}
                                                                    _ = fi287
                                                                    _ = v288
                                                                }
                                                            case 127:
                                                                switch ("" + "") {
                                                                case "":
                                                                    if len(v262) >= 2 {
                                                                        v262 = v262[1:len(v262)]
                                                                    }
                                                                }
                                                            }
                                                            _ = v286
                                                        }
                                                    }
                                                case -225:
                                                    var v289 string = "a"
                                                    _ = v289
                                                default:
                                                    fmt.Println(fi254)
                                                }
                                            }
                                        }
                                    default:
                                        var v290 uint = 445
                                        _ = v290
                                    }
                                }
                                _ = v261
                                _ = v262
                            } else {
                                fmt.Println(v124, v1)
                                var v291 uint16 = (^(v132))
                                v1 += (v1 / ((255) | 1))
                                _ = v291
                            }
                            v133 += (v133 >> ((4294967295) & 63))
                        } else {
                            for fi292 := 0; fi292 < 18; fi292++ {
                                v293 := GenS30{F0: 228, F1: v21}
                                v124 = 520
                                _ = fi292
                                _ = v293
                            }
                        }
                    }
                    if (!((v21 % ((v21) | 1)) <= v21)) {
                        v3 = 1e10
                        var v294 bool = false
                        _ = v294
                    }
                    v295 := GenS31{F0: v133, F1: 4294967295, F2: v133}
                    _ = v295
                }
                _ = fi254
            }
            v296 := []int{v124, 32768, v124}
            _ = v296
        }
    }
    for fi297 := 0; fi297 < 11; fi297++ {
        v3 += (v3 * 0.1)
        _ = fi297
    }
    fmt.Println("v1:", v1)
    fmt.Println("v124:", v124)
    fmt.Println("v132:", v132)
    fmt.Println("v133:", v133)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v124
    _ = v132
    _ = v133
    _ = v2
    _ = v21
    _ = v3
}
