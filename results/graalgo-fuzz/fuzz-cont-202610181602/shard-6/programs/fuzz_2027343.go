package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 string; F1 int64; F2 int16 }
type GenS3 struct { F0 bool; F1 bool }
type GenS4 struct { F0 uint64; F1 int }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint; F1 int16; F2 uint32 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint8; F1 int16; F2 uint }
type GenS9 struct { F0 int8; F1 uint64 }
type GenS10 struct { F0 bool; F1 uint16; F2 uint }
type GenS11 struct { F0 int8; F1 int16; F2 uint8 }
func genFn0(p0 uint64) uint16 {
    var v5 int8 = 42
    if false {
        v6 := GenS1{F0: 496, F1: 2147483647}
        var v7 float32 = 0.1
        v8 := []int{2147483647}
        _ = v6
        _ = v7
        _ = v8
    } else {
        if false {
            var v9 int8 = v5
            _ = v9
        } else {
            v5 = -128
        }
    }
    v10 := []int{5}
    v11 := GenS2{F0: "", F1: 1, F2: 848}
    _ = p0
    _ = v10
    _ = v11
    _ = v5
    return 1
}
func genFn1() int16 {
    v12 := []int{-9223372036854775808}
    var v13 int32 = -2147483648
    for fi14 := 0; fi14 < 6; fi14++ {
        _ = len(v12)
        for fi15 := 0; fi15 < 16; fi15++ {
            if true {
                fmt.Println(v13, fi15, fi14)
                v12 = append(v12, (fi14 ^ fi14))
                var v16 float32 = 3.14159
                _ = v16
            }
            v13 += int32(fi14)
            switch (fi15 & (fi15 >> ((1) & 63))) {
            case 2147483647:
                var v17 float32 = float32(0.1)
                v13 -= v13
                if false {
                    if len(v12) > 0 {
                        v12[0] = fi15
                    }
                    fmt.Println(fi15)
                    if (128 == 170) {
                        if false {
                            v18 := GenS1{F0: 42, F1: 2147483647}
                            var v19 uint = 855
                            fmt.Println(v19, v19, v19)
                            _ = v18
                            _ = v19
                        }
                        var v20 uint8 = 168
                        _ = v20
                    }
                    var v21 float32 = (float32(0.1) + 1e10)
                    _ = v21
                }
                _ = v17
            case -460:
                var v22 uint = 1
                _ = v22
            case -1:
                var v23 int = (-(fi14))
                _ = v23
            default:
                if (((-(v13)) < (v13 * (v13 & (^(v13))))) && true) {
                    if true {
                        fmt.Println(v13, fi15, fi14)
                        if (true || (0.0 <= 0.1)) {
                            if len(v12) >= 2 {
                                v12 = v12[1:len(v12)]
                            }
                            var v24 bool = true
                            _ = v24
                        }
                    }
                    v12 = append(v12, fi14)
                    v13 = 42
                    if (genFn0(4294967295) == 953) {
                        v12 = append(v12, (^(fi15)))
                        v12 = append(v12, 42)
                        var v25 int64 = -471
                        v25 = v25
                        _ = v25
                    } else {
                        v26 := GenS1{F0: 1, F1: (^(v13))}
                        _ = v26
                    }
                    switch (fi15 ^ 128) {
                    case 156:
                        v13 = -2147483648
                    case 32767:
                        switch -379 {
                        case -1:
                            v27 := GenS2{F0: "a", F1: 42, F2: -32768}
                            _ = v27
                        default:
                            v28 := GenS2{F0: (("foo bar" + ("\n\t" + ("" + ""))) + ("\n\t" + "")), F1: 42, F2: 1}
                            _ = v28
                        }
                    }
                } else {
                    v29 := GenS1{F0: 42, F1: v13}
                    var v30 float64 = 0.0
                    var v31 uint16 = 2
                    switch fi15 {
                    case -632:
                        v13 = (-(v13))
                    case 4294967295:
                        v31 += v31
                    }
                    _ = v29
                    _ = v30
                    _ = v31
                }
            }
            _ = fi15
        }
        _ = fi14
    }
    switch ("a" + "a") {
    case "\n\t":
        fmt.Println(v13)
    }
    var v32 int8 = 32
    _ = v12
    _ = v13
    _ = v32
    return -32768
}

func main() {
    var v1 uint = 0
    var v2 int64 = -455
    var v3 int64 = v2
    var v4 float64 = (-2.25 + (1.5 * ((((-2.25 * (((1.5 * (0.0 - 0.0)) + 0.0) - float64(1e10))) / 0.1) - 3.14159) + 0.0)))
    for fi33 := 0; fi33 < 7; fi33++ {
        if (genFn0(0) <= 0) {
            var v34 float32 = 0.0
            _ = v34
        } else {
            if true {
                var v35 int32 = -712
                fmt.Println(v4, v3)
                v36 := GenS1{F0: 1000, F1: (v35 << ((v1) & 63))}
                _ = v35
                _ = v36
            } else {
                v1 = v1
            }
        }
        v1 -= v1
        _ = fi33
    }
    v4 = v4
    if (49 <= 0) {
        var v37 int64 = (v3 * (v2 * (v3 / ((v2) | 1))))
        _ = v37
    } else {
        var v38 string = ("\n\t" + "hello")
        v39 := []int{42, 238, 691}
        _ = v38
        _ = v39
    }
    var v40 float32 = -2.25
    var v41 uint64 = 1
    if (2 != 191) {
        v2 = v3
        for fi42 := 0; fi42 < 1; fi42++ {
            if false {
                v1 = (v1 | v1)
                if false {
                    switch ("hello" + ("\n\t" + "hello")) {
                    case "hello":
                        if ((42 <= (v41 / ((1) | 1))) || true) {
                            fmt.Println(v2, fi42)
                            var v43 int16 = int16(fi42)
                            var v44 int32 = 1
                            _ = v43
                            _ = v44
                        } else {
                            v45 := []int{fi42, -576}
                            _ = v45
                        }
                    case "":
                        if (0 <= 1) {
                            for fi46 := 0; fi46 < 9; fi46++ {
                                var v47 uint = v1
                                _ = fi46
                                _ = v47
                            }
                            var v48 float64 = v4
                            var v49 string = ("a" + (("hello" + (((("" + "\n\t") + "a") + "foo bar") + ("" + "foo bar"))) + (("\n\t" + "") + "a")))
                            _ = v48
                            _ = v49
                        } else {
                            for fi50 := 0; fi50 < 3; fi50++ {
                                v51 := []int{fi42}
                                var v52 uint8 = 60
                                if len(v51) >= 2 {
                                    v51 = v51[1:len(v51)]
                                }
                                _ = fi50
                                _ = v51
                                _ = v52
                            }
                        }
                    case "\n\t":
                        v1 = v1
                    default:
                        fmt.Println(v40, v4, v4)
                    }
                }
            }
            _ = fi42
        }
    } else {
        var v53 int8 = -93
        switch 414 {
        case 32768:
            v3 += (v3 + 944)
        case 1:
            for fi54 := 0; fi54 < 2; fi54++ {
                if false {
                    fmt.Println(v3, v41)
                } else {
                    if true {
                        v55 := GenS3{F0: (!false), F1: (869 > (^(fi54)))}
                        if false {
                            var v56 float32 = 3.14159
                            var v57 bool = (127 == 127)
                            switch fi54 {
                            case -1:
                                var v58 int64 = (v3 % (((v3 | (v2 % (((v2 ^ v3)) | 1)))) | 1))
                                _ = v58
                            case 940:
                                switch fi54 {
                                case -244:
                                    var v59 int32 = 1
                                    _ = v59
                                case 9223372036854775807:
                                    v40 = v56
                                case 128:
                                    switch ("" + "a") {
                                    case "":
                                        if v57 {
                                            if true {
                                                var v60 string = "\n\t"
                                                switch v60 {
                                                case "a":
                                                    v61 := GenS2{F0: v60, F1: (v3 + 1), F2: -805}
                                                    _ = v61
                                                }
                                                _ = v60
                                            } else {
                                                v41 += 42
                                                var v62 float64 = v4
                                                fmt.Println(v2, v40)
                                                _ = v62
                                            }
                                        }
                                    default:
                                        var v63 uint8 = 63
                                        _ = v63
                                    }
                                default:
                                    v64 := GenS4{F0: v41, F1: 42}
                                    _ = v64
                                }
                            default:
                                v56 -= v40
                            }
                            _ = v56
                            _ = v57
                        } else {
                            var v65 int = (fi54 + fi54)
                            _ = v65
                        }
                        _ = v55
                    }
                }
                v66 := []int{fi54, (^(fi54))}
                _ = fi54
                _ = v66
            }
        case 4294967295:
            if false {
                switch "" {
                case "hello":
                    var v67 uint64 = v41
                    _ = v67
                case "":
                    if (((v4 / v4) == v4) || true) {
                        if (2147483647 <= 127) {
                            if (1 > 255) {
                                v2 = v3
                                v68 := GenS5{F0: 2}
                                _ = v68
                            } else {
                                var v69 uint = (v1 | (^(v1)))
                                _ = v69
                            }
                            fmt.Println(v4)
                            v1 = v1
                        } else {
                            if ((2 >= 236) || (127 >= 2)) {
                                var v70 int16 = -32768
                                _ = v70
                            } else {
                                var v71 uint8 = 99
                                fmt.Println(v53)
                                _ = v71
                            }
                            var v72 string = "a"
                            var v73 int32 = -1
                            _ = v72
                            _ = v73
                        }
                        for fi74 := 0; fi74 < 15; fi74++ {
                            var v75 uint = (^(v1))
                            var v76 uint64 = (^(v41))
                            fmt.Println(fi74)
                            _ = fi74
                            _ = v75
                            _ = v76
                        }
                    } else {
                        var v77 int64 = int64(v41)
                        if (false && false) {
                            for fi78 := 0; fi78 < 15; fi78++ {
                                v79 := []int{}
                                _ = fi78
                                _ = v79
                            }
                            if (v3 <= 914) {
                                if true {
                                    if ((761 <= (v77 * (v2 | v3))) && false) {
                                        v4 = v4
                                        var v80 string = ("\n\t" + "foo bar")
                                        _ = v80
                                    }
                                    for fi81 := 0; fi81 < 16; fi81++ {
                                        v2 = -420
                                        var v82 uint8 = 0
                                        var v83 uint = (v1 % (((v1 + 241)) | 1))
                                        _ = fi81
                                        _ = v82
                                        _ = v83
                                    }
                                    for fi84 := 0; fi84 < 10; fi84++ {
                                        if false {
                                            v4 = (-2.25 - (float64(v40) / v4))
                                            v85 := []int{fi84, (-(fi84)), fi84}
                                            v85 = append(v85, (-(fi84)))
                                            _ = v85
                                        } else {
                                            if true {
                                                var v86 uint16 = 845
                                                v87 := GenS2{F0: "hello", F1: 0, F2: -32768}
                                                _ = v86
                                                _ = v87
                                            }
                                            v2 -= v3
                                            v88 := GenS6{F0: (v1 + (v1 + (v1 ^ 18446744073709551615))), F1: 1, F2: 2147483647}
                                            _ = v88
                                        }
                                        v89 := []int{(fi84 | fi84), fi84}
                                        _ = fi84
                                        _ = v89
                                    }
                                } else {
                                    if (127 > 32767) {
                                        var v90 float64 = (v4 + (((v4 + v4) - float64(0.0)) + v4))
                                        _ = v90
                                    }
                                    var v91 int16 = 1
                                    _ = v91
                                }
                                v92 := GenS1{F0: 32767, F1: 154}
                                _ = v92
                            }
                        }
                        v93 := GenS1{F0: 32767, F1: 0}
                        _ = v77
                        _ = v93
                    }
                case "a":
                    var v94 float64 = 0.1
                    _ = v94
                }
            } else {
                var v95 uint32 = 65535
                v95 = 0
                v1 = 18446744073709551615
                _ = v95
            }
        }
        _ = v53
    }
    v41 = (^(v41))
    switch 0 {
    case 32768:
        if (110 <= -4) {
            if ((-1 >= 2147483647) && (v41 < 18446744073709551615)) {
                v1 = v1
            }
            for fi96 := 0; fi96 < 16; fi96++ {
                v40 += v40
                _ = fi96
            }
        } else {
            v3 = int64(v1)
        }
    case -250:
        if ((^(v41)) <= 328) {
            var v97 float32 = v40
            _ = v97
        }
    case -1:
        for fi98 := 0; fi98 < 19; fi98++ {
            if true {
                for fi99 := 0; fi99 < 18; fi99++ {
                    fmt.Println(fi99, v41)
                    v100 := []int{-1, -699, fi98}
                    _ = fi99
                    _ = v100
                }
            } else {
                v2 = (v3 << ((962) & 63))
                v101 := GenS7{F0: 2}
                var v102 bool = false
                _ = v101
                _ = v102
            }
            v103 := []int{(fi98 + fi98), 36}
            if true {
                v41 += (v41 ^ 207)
                switch (fi98 + (fi98 | (fi98 - fi98))) {
                case 32768:
                    var v104 bool = true
                    _ = v104
                case -9223372036854775808:
                    var v105 bool = (-1 == 127)
                    _ = v105
                case 2147483647:
                    if (!true) {
                        if (false || (!true)) {
                            fmt.Println(v2)
                            if len(v103) > 0 {
                                v103[0] = (fi98 & fi98)
                            }
                            v40 += v40
                        } else {
                            fmt.Println(v41, v3)
                            for fi106 := 0; fi106 < 10; fi106++ {
                                var v107 uint64 = (v41 ^ 798)
                                switch 4294967295 {
                                case 696:
                                    if (!false) {
                                        var v108 uint16 = 42
                                        var v109 float32 = (1.5 - v40)
                                        v108 = (v108 & 42)
                                        if false {
                                            if false {
                                                v4 -= (-2.25 + (3.14159 + v4))
                                                if len(v103) > 0 {
                                                    v103[0] = fi98
                                                }
                                                v1 -= v1
                                            }
                                        } else {
                                            var v110 uint64 = (v41 * 2)
                                            v111 := GenS3{F0: (false && (-32768 >= -328)), F1: ("" != "foo bar")}
                                            _ = v110
                                            _ = v111
                                        }
                                        _ = v108
                                        _ = v109
                                    }
                                case -9223372036854775808:
                                    v112 := GenS8{F0: 1, F1: genFn1(), F2: v1}
                                    _ = v112
                                case -263:
                                    fmt.Println(v2)
                                }
                                var v113 int16 = -32768
                                fmt.Println(v1, fi98, v2)
                                v103 = append(v103, (-(fi106)))
                                _ = fi106
                                _ = v107
                                _ = v113
                            }
                            var v114 uint16 = 456
                            _ = v114
                        }
                        v40 = v40
                    }
                default:
                    fmt.Println(v1)
                }
            }
            _ = fi98
            _ = v103
        }
    }
    for fi115 := 0; fi115 < 2; fi115++ {
        if false {
            if true {
                for fi116 := 0; fi116 < 19; fi116++ {
                    fmt.Println(fi115)
                    var v117 int8 = -128
                    v118 := []int{(fi116 >> ((4294967295) & 63))}
                    _ = fi116
                    _ = v117
                    _ = v118
                }
                if (889 != 255) {
                    v119 := []int{-1, -1, (-(fi115)), fi115}
                    if true {
                        var v120 uint = (v1 * (^(v1)))
                        for fi121 := 0; fi121 < 10; fi121++ {
                            _ = len(v119)
                            _ = fi121
                        }
                        _ = v120
                    } else {
                        v122 := GenS8{F0: 255, F1: 32767, F2: 4294967295}
                        for fi123 := 0; fi123 < 9; fi123++ {
                            v124 := GenS3{F0: (65535 <= 2147483647), F1: (!((v4 * v4) < 1.5))}
                            _ = fi123
                            _ = v124
                        }
                        v125 := GenS9{F0: 0, F1: uint64(v1)}
                        _ = v122
                        _ = v125
                    }
                    _ = v119
                }
                v126 := GenS10{F0: (true && false), F1: 2, F2: v1}
                _ = v126
            }
            v127 := []int{(fi115 % ((-9223372036854775808) | 1)), fi115}
            v127 = append(v127, (-(fi115)))
            _ = v127
        } else {
            v128 := []int{(fi115 - 42), (fi115 ^ 1), -832}
            switch "foo bar" {
            case "a":
                var v129 uint64 = v41
                _ = v129
            default:
                _ = len(v128)
            }
            _ = v128
        }
        v4 = v4
        _ = fi115
    }
    if (v1 <= v1) {
        if false {
            for fi130 := 0; fi130 < 9; fi130++ {
                v4 = 0.1
                if true {
                    switch "\n\t" {
                    case "hello":
                        v40 += -2.25
                    case "foo bar":
                        fmt.Println(fi130, v2)
                    case "":
                        v4 += v4
                    default:
                        v131 := []int{-1, -2147483648}
                        _ = v131
                    }
                    var v132 float32 = v40
                    var v133 uint16 = 65535
                    _ = v132
                    _ = v133
                }
                v134 := []int{(fi130 - fi130)}
                _ = fi130
                _ = v134
            }
            v4 = (v4 + v4)
        }
        var v135 uint8 = 127
        v136 := []int{42, 32768, -9223372036854775808}
        _ = v135
        _ = v136
    } else {
        fmt.Println(v2, v2, v2)
    }
    var v137 string = ("hello" + ("a" + "a"))
    if false {
        v138 := []int{127, 504, 1, 4294967295}
        _ = v138
    }
    v2 += (v3 - 9223372036854775807)
    v139 := GenS7{F0: 1}
    switch v137 {
    case "hello":
        if true {
            if false {
                fmt.Println(v3, v137, v40)
                var v140 uint64 = v41
                _ = v140
            }
            for fi141 := 0; fi141 < 19; fi141++ {
                v142 := GenS11{F0: 0, F1: genFn1(), F2: 2}
                v143 := []int{fi141, fi141, -2147483648}
                _ = fi141
                _ = v142
                _ = v143
            }
        }
    case "":
        for fi144 := 0; fi144 < 4; fi144++ {
            v145 := GenS7{F0: 163}
            v146 := []int{}
            var v147 float32 = ((float32(v4) * v40) - 1e10)
            _ = fi144
            _ = v145
            _ = v146
            _ = v147
        }
    case "foo bar":
        var v148 float32 = (v40 - v40)
        _ = v148
    default:
        var v149 int8 = 42
        _ = v149
    }
    v40 = v40
    var v150 int8 = 39
    fmt.Println("v1:", v1)
    fmt.Println("v137:", v137)
    fmt.Println("v150:", v150)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v137
    _ = v139
    _ = v150
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
}
