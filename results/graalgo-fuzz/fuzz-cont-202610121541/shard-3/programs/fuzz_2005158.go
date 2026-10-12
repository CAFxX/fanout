package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 uint }
type GenS4 struct { F0 uint8; F1 uint64; F2 bool }
type GenS5 struct { F0 int64; F1 int64 }
type GenS6 struct { F0 uint16; F1 uint64; F2 int32 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int16; F1 int8; F2 int8 }
type GenS9 struct { F0 uint32; F1 int8 }
type GenS10 struct { F0 int64; F1 int32 }
type GenS11 struct { F0 int64; F1 int8; F2 bool }
type GenS12 struct { F0 uint8; F1 bool }
func genFn0(p0 string, p1 uint32) int8 {
    p1 -= (p1 ^ p1)
    for fi4 := 0; fi4 < 18; fi4++ {
        v5 := GenS1{F0: 193}
        _ = fi4
        _ = v5
    }
    var v6 uint32 = 372
    _ = p0
    _ = p1
    _ = v6
    return -57
}
func genFn1() int8 {
    v7 := []int{157, 42, 1, 9223372036854775807}
    var v8 uint32 = 1
    var v9 bool = (127 > 127)
    v10 := GenS1{F0: 65535}
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 30
}

func main() {
    var v1 float64 = (0.0 - 0.0)
    var v2 uint = 65535
    var v3 uint = v2
    var v11 uint32 = 839
    if (127 <= 4294967295) {
        var v12 uint = (v3 - v3)
        if ((-1 > -123) || (-434 >= 1)) {
            v13 := []int{369, -1, 4294967295, -9223372036854775808}
            _ = v13
        } else {
            var v14 string = (((("" + ("" + "a")) + ("a" + "a")) + "\n\t") + "hello")
            for fi15 := 0; fi15 < 19; fi15++ {
                v16 := GenS1{F0: 2}
                v17 := []int{}
                _ = fi15
                _ = v16
                _ = v17
            }
            _ = v14
        }
        switch ("foo bar" + ((("a" + ("" + "")) + ("" + ("\n\t" + ("" + "a")))) + "")) {
        case "\n\t":
            if (v11 > (v11 % (((^(v11))) | 1))) {
                var v18 uint16 = 65535
                v12 -= v3
                _ = v18
            } else {
                var v19 int16 = 1
                _ = v19
            }
        case "foo bar":
            v20 := GenS2{F0: v2}
            _ = v20
        default:
            for fi21 := 0; fi21 < 19; fi21++ {
                v11 += (^(v11))
                _ = fi21
            }
        }
        _ = v12
    } else {
        for fi22 := 0; fi22 < 12; fi22++ {
            v11 -= v11
            _ = fi22
        }
        var v23 string = ""
        for fi24 := 0; fi24 < 15; fi24++ {
            v25 := []int{9223372036854775807, fi24}
            _ = fi24
            _ = v25
        }
        _ = v23
    }
    var v26 float32 = (((float32(float64(1e10)) + 0.1) * 0.0) - 0.1)
    switch 397 {
    case 0:
        v27 := []int{42, int(v2), 42}
        _ = v27
    case 1:
        var v28 uint64 = 2
        _ = v28
    default:
        v1 += (v1 / v1)
    }
    if ((0.1 + v1) == 3.14159) {
        var v29 float64 = ((-2.25 + v1) - (1e10 / 1.0))
        v2 = uint(v11)
        _ = v29
    } else {
        v30 := GenS1{F0: 42}
        _ = v30
    }
    var v31 float64 = (v1 + v1)
    var v32 int64 = 1
    switch "\n\t" {
    case "a":
        for fi33 := 0; fi33 < 2; fi33++ {
            switch "" {
            case "\n\t":
                fmt.Println(v32, fi33, fi33)
            case "a":
                v32 += int64(v11)
            case "hello":
                if false {
                    v34 := []int{}
                    var v35 int = (fi33 ^ fi33)
                    _ = len(v34)
                    _ = v34
                    _ = v35
                } else {
                    if true {
                        var v36 int = (fi33 / ((fi33) | 1))
                        for fi37 := 0; fi37 < 6; fi37++ {
                            switch fi37 {
                            case 681:
                                if true {
                                    v38 := []int{128, (v36 | 1), fi37, (fi37 - (fi37 - (^(fi37))))}
                                    _ = v38
                                } else {
                                    v2 += 990
                                    var v39 int8 = -99
                                    switch (("foo bar" + ("" + "hello")) + "a") {
                                    case "":
                                        v40 := GenS2{F0: (v2 % ((v2) | 1))}
                                        _ = v40
                                    }
                                    _ = v39
                                }
                            case -2147483648:
                                switch "hello" {
                                case "\n\t":
                                    v41 := GenS3{F0: 944, F1: v2}
                                    _ = v41
                                case "a":
                                    fmt.Println(v3)
                                    v11 += v11
                                    var v42 uint32 = v11
                                    _ = v42
                                case "":
                                    v43 := GenS1{F0: 790}
                                    _ = v43
                                default:
                                    fmt.Println(v3, v36)
                                }
                            case -9223372036854775808:
                                if true {
                                    if false {
                                        var v44 uint16 = 42
                                        switch "hello" {
                                        case "":
                                            v1 = float64(v26)
                                        case "foo bar":
                                            v45 := GenS2{F0: v3}
                                            _ = v45
                                        case "a":
                                            v36 += fi37
                                        }
                                        v3 = v3
                                        _ = v44
                                    } else {
                                        v11 = v11
                                        fmt.Println(v11)
                                        var v46 int64 = 1
                                        _ = v46
                                    }
                                    var v47 float64 = v31
                                    if (18446744073709551615 <= 2) {
                                        v48 := []int{53, v36, fi33, (-(fi37))}
                                        _ = v48
                                    }
                                    _ = v47
                                } else {
                                    v1 -= v31
                                    switch (-(fi33)) {
                                    case 9223372036854775807:
                                        v31 -= (v31 + (v31 + float64(v26)))
                                    case -2147483648:
                                        if false {
                                            switch (fi33 + (fi33 + fi33)) {
                                            case 1:
                                                var v49 uint16 = 32767
                                                _ = v49
                                            case 0:
                                                var v50 uint16 = 818
                                                _ = v50
                                            case -887:
                                                v51 := GenS1{F0: 2}
                                                _ = v51
                                            default:
                                                fmt.Println(v3, v31)
                                            }
                                            if false {
                                                fmt.Println(fi33, v31, v32)
                                                v52 := []int{}
                                                var v53 float64 = (v1 + v31)
                                                var v54 bool = (v11 == (v11 / (((v11 & v11)) | 1)))
                                                _ = v52
                                                _ = v53
                                                _ = v54
                                            }
                                        }
                                    }
                                    switch "" {
                                    case "\n\t":
                                        v3 = (v2 - (v2 >> ((uint(v11)) & 63)))
                                    case "":
                                        if false {
                                            fmt.Println(v3, v3)
                                            v55 := []int{(fi33 * 434), 0, (fi33 / (((fi37 >> (((v3 % ((42) | 1))) & 63))) | 1))}
                                            _ = v55
                                        } else {
                                            v56 := GenS2{F0: 18446744073709551615}
                                            if (false || (false || true)) {
                                                var v57 uint = 704
                                                if true {
                                                    v58 := GenS4{F0: 128, F1: 483, F2: false}
                                                    _ = v58
                                                }
                                                switch ("hello" + ("\n\t" + ("foo bar" + "foo bar"))) {
                                                case "foo bar":
                                                    if true {
                                                        var v59 string = ("foo bar" + "hello")
                                                        _ = v59
                                                    }
                                                case "\n\t":
                                                    var v60 float64 = 0.0
                                                    _ = v60
                                                case "a":
                                                    v1 += (v31 / v1)
                                                default:
                                                    var v61 int64 = -968
                                                    _ = v61
                                                }
                                                _ = v57
                                            } else {
                                                v62 := []int{42, 32768}
                                                _ = v62
                                            }
                                            _ = v56
                                        }
                                    }
                                }
                            default:
                                v63 := GenS1{F0: 32767}
                                switch 0 {
                                case 1:
                                    var v64 int16 = 42
                                    v65 := GenS5{F0: (v32 * -1), F1: (v32 / ((-1) | 1))}
                                    _ = v64
                                    _ = v65
                                default:
                                    v26 += ((v26 * v26) / (v26 / (v26 / v26)))
                                }
                                _ = v63
                            }
                            fmt.Println(v2, v36)
                            v66 := []int{fi33}
                            _ = fi37
                            _ = v66
                        }
                        _ = v36
                    } else {
                        fmt.Println(v26, v32)
                        v31 = v31
                        v67 := GenS3{F0: 155, F1: 4294967295}
                        _ = v67
                    }
                    v68 := GenS3{F0: 2, F1: 2}
                    _ = v68
                }
            }
            _ = fi33
        }
    case "foo bar":
        switch "a" {
        case "":
            for fi69 := 0; fi69 < 14; fi69++ {
                fmt.Println(v31, v3, fi69)
                for fi70 := 0; fi70 < 4; fi70++ {
                    switch -1 {
                    case 2147483647:
                        fmt.Println(v3, fi70, v26)
                        if ((1e10 < v31) && ((!true) || (128 < 226))) {
                            var v71 float32 = 1.5
                            v3 -= 2
                            _ = v71
                        }
                        v26 = 1.5
                    case 32767:
                        if (!false) {
                            fmt.Println(v31)
                            var v72 int8 = -1
                            v72 -= 42
                            _ = v72
                        }
                    }
                    v32 += (v32 << ((1) & 63))
                    v73 := []int{}
                    _ = fi70
                    _ = v73
                }
                _ = fi69
            }
        case "a":
            if (v3 < (^(v3))) {
                var v74 float64 = 0.0
                _ = v74
            } else {
                v75 := []int{-677, -824, 2147483647}
                for fi76 := 0; fi76 < 12; fi76++ {
                    v32 += -9223372036854775808
                    v77 := GenS2{F0: v3}
                    for fi78 := 0; fi78 < 7; fi78++ {
                        var v79 int = fi78
                        _ = fi78
                        _ = v79
                    }
                    _ = fi76
                    _ = v77
                }
                _ = v75
            }
        case "\n\t":
            fmt.Println(v26, v2)
        default:
            v80 := []int{}
            _ = v80
        }
    }
    if true {
        fmt.Println(v11, v31)
        var v81 int8 = -128
        v11 += (v11 / ((v11) | 1))
        _ = v81
    } else {
        v82 := []int{}
        v83 := GenS6{F0: uint16(v11), F1: 1, F2: 0}
        v3 -= v3
        _ = v82
        _ = v83
    }
    switch -1 {
    case -277:
        if true {
            switch 9223372036854775807 {
            case -2147483648:
                v84 := []int{}
                _ = v84
            case -1:
                var v85 uint = (v2 + 4294967295)
                _ = v85
            }
            v3 += 835
            if (v11 == (v11 - 552)) {
                for fi86 := 0; fi86 < 12; fi86++ {
                    switch "hello" {
                    case "\n\t":
                        var v87 uint32 = (v11 | 689)
                        _ = v87
                    case "":
                        var v88 uint64 = 2
                        _ = v88
                    case "a":
                        var v89 uint32 = 657
                        _ = v89
                    }
                    if (v32 > (v32 + (v32 | v32))) {
                        var v90 float32 = (0.0 + (v26 + (3.14159 / (v26 * v26))))
                        var v91 uint16 = 255
                        v92 := []int{-9223372036854775808}
                        _ = v90
                        _ = v91
                        _ = v92
                    } else {
                        switch ("foo bar" + "") {
                        case "a":
                            v93 := []int{42, 253, 4294967295}
                            _ = v93
                        case "hello":
                            if (127 < 42) {
                                var v94 uint16 = 32767
                                var v95 int16 = 1000
                                _ = v94
                                _ = v95
                            }
                        default:
                            v31 = 3.14159
                        }
                        if (false || (3.14159 >= 0.1)) {
                            var v96 bool = false
                            if (v26 <= v26) {
                                v32 = int64(fi86)
                                if (v96 && (210 <= 18446744073709551615)) {
                                    v3 -= v3
                                    v97 := GenS7{F0: (v11 << ((v3) & 63))}
                                    _ = v97
                                }
                            } else {
                                var v98 bool = (!v96)
                                if true {
                                    if (248 <= 2) {
                                        if (v98 || (!false)) {
                                            for fi99 := 0; fi99 < 12; fi99++ {
                                                var v100 int = -9223372036854775808
                                                v101 := []int{(fi99 * fi99)}
                                                switch ("a" + (((("\n\t" + "foo bar") + "foo bar") + ("" + ((("hello" + ((("foo bar" + "\n\t") + "foo bar") + "")) + "\n\t") + "\n\t"))) + "a")) {
                                                case "a":
                                                    if v96 {
                                                        fmt.Println(v3, fi99)
                                                    } else {
                                                        v101 = append(v101, (fi86 ^ fi86))
                                                    }
                                                    var v102 uint64 = 0
                                                    var v103 uint32 = (^(v11))
                                                    _ = v102
                                                    _ = v103
                                                case "foo bar":
                                                    v100 = -779
                                                case "\n\t":
                                                    fmt.Println(v32, v98, v11)
                                                default:
                                                    v104 := GenS8{F0: 459, F1: 0, F2: int8(v11)}
                                                    _ = v104
                                                }
                                                _ = fi99
                                                _ = v100
                                                _ = v101
                                            }
                                        } else {
                                            for fi105 := 0; fi105 < 9; fi105++ {
                                                if false {
                                                    var v106 uint64 = 18446744073709551615
                                                    _ = v106
                                                }
                                                v32 += (v32 | -1)
                                                fmt.Println(fi86)
                                                fmt.Println(v11, v3)
                                                var v107 int32 = int32(v11)
                                                _ = fi105
                                                _ = v107
                                            }
                                        }
                                        for fi108 := 0; fi108 < 17; fi108++ {
                                            v31 = (v31 + v31)
                                            switch ("\n\t" + "foo bar") {
                                            case "":
                                                v98 = (-698 > 1)
                                            }
                                            v109 := []int{fi108, -9223372036854775808}
                                            _ = fi108
                                            _ = v109
                                        }
                                    } else {
                                        fmt.Println(v31)
                                        v32 = 1
                                        if v98 {
                                            for fi110 := 0; fi110 < 19; fi110++ {
                                                v111 := GenS9{F0: (v11 | v11), F1: -128}
                                                var v112 int = (fi110 & (^(fi110)))
                                                if (!(4294967295 != 18446744073709551615)) {
                                                    v113 := []int{154}
                                                    fmt.Println(v11)
                                                    _ = v113
                                                } else {
                                                    v32 = (^(v32))
                                                    v3 -= v2
                                                    fmt.Println(v32, v26, fi86)
                                                    var v114 int16 = 42
                                                    _ = v114
                                                }
                                                _ = fi110
                                                _ = v111
                                                _ = v112
                                            }
                                        }
                                    }
                                    v96 = (((1000 < -32768) || true) && true)
                                    v115 := GenS10{F0: (^(v32)), F1: 1}
                                    _ = v115
                                } else {
                                    var v116 int = fi86
                                    for fi117 := 0; fi117 < 12; fi117++ {
                                        v116 += -1
                                        _ = fi117
                                    }
                                    if v96 {
                                        fmt.Println(v2)
                                        var v118 string = (("\n\t" + "foo bar") + (("\n\t" + "hello") + "\n\t"))
                                        v119 := GenS6{F0: 280, F1: 0, F2: 139}
                                        _ = v118
                                        _ = v119
                                    }
                                    _ = v116
                                }
                                _ = v98
                            }
                            _ = v96
                        } else {
                            if ((-32768 != -32768) || false) {
                                v31 += v31
                            } else {
                                var v120 uint32 = v11
                                for fi121 := 0; fi121 < 5; fi121++ {
                                    v122 := GenS11{F0: v32, F1: 42, F2: false}
                                    if false {
                                        v123 := []int{}
                                        _ = v123
                                    } else {
                                        fmt.Println(fi121, v26)
                                        switch ("foo bar" + "hello") {
                                        case "":
                                            v1 = ((1.5 / v1) / v1)
                                        }
                                        var v124 float32 = (v26 + v26)
                                        var v125 int64 = (-(v32))
                                        _ = v124
                                        _ = v125
                                    }
                                    _ = fi121
                                    _ = v122
                                }
                                _ = v120
                            }
                        }
                        v2 += (v3 >> ((v3) & 63))
                    }
                    if true {
                        var v126 uint = (v2 + (v2 + v2))
                        fmt.Println(v32)
                        _ = v126
                    }
                    _ = fi86
                }
            } else {
                v127 := GenS3{F0: 65535, F1: (v3 ^ 233)}
                if true {
                    var v128 uint32 = (v11 >> ((866) & 63))
                    switch -602 {
                    case 1:
                        v2 += v3
                    default:
                        v11 = v128
                    }
                    fmt.Println(v31)
                    _ = v128
                }
                _ = v127
            }
        }
    default:
        if (false || (0 == 2147483647)) {
            if ((false && (v3 == v3)) && (!((3.14159 - v26) >= 0.1))) {
                for fi129 := 0; fi129 < 10; fi129++ {
                    var v130 int8 = 1
                    var v131 float64 = (v1 - 0.1)
                    for fi132 := 0; fi132 < 11; fi132++ {
                        v31 = (v1 + ((0.0 / (v131 + -2.25)) / v131))
                        v131 += ((v31 - v1) - (0.1 * v1))
                        var v133 uint64 = 119
                        switch fi132 {
                        case 127:
                            var v134 float32 = 1.5
                            _ = v134
                        default:
                            var v135 float32 = -2.25
                            _ = v135
                        }
                        v136 := []int{}
                        _ = fi132
                        _ = v133
                        _ = v136
                    }
                    _ = fi129
                    _ = v130
                    _ = v131
                }
                v137 := GenS11{F0: v32, F1: -128, F2: true}
                v138 := []int{42}
                _ = v137
                _ = v138
            }
            switch 4294967295 {
            case 32767:
                if (!false) {
                    fmt.Println(v31)
                }
            case 32768:
                if true {
                    for fi139 := 0; fi139 < 10; fi139++ {
                        v1 = (v31 * v1)
                        _ = fi139
                    }
                } else {
                    v140 := []int{128, -2147483648, 128, 32767}
                    var v141 int16 = 1
                    if false {
                        var v142 uint16 = 2
                        _ = v142
                    } else {
                        v143 := GenS1{F0: 2}
                        var v144 string = ""
                        _ = v143
                        _ = v144
                    }
                    _ = v140
                    _ = v141
                }
            }
            if true {
                if false {
                    fmt.Println(v3, v1)
                    v2 = 18446744073709551615
                    if (32767 == 32767) {
                        for fi145 := 0; fi145 < 16; fi145++ {
                            fmt.Println(v32, fi145)
                            var v146 uint8 = 255
                            _ = fi145
                            _ = v146
                        }
                        var v147 int16 = -856
                        if true {
                            var v148 uint32 = v11
                            v149 := GenS12{F0: 1, F1: (269 > 4294967295)}
                            _ = v148
                            _ = v149
                        }
                        _ = v147
                    } else {
                        var v150 int = 42
                        if true {
                            fmt.Println(v26)
                            var v151 float64 = 1.5
                            _ = v151
                        }
                        fmt.Println(v2, v150, v2)
                        _ = v150
                    }
                }
                fmt.Println(v31, v3)
                for fi152 := 0; fi152 < 19; fi152++ {
                    v153 := []int{int(v32)}
                    for fi154 := 0; fi154 < 9; fi154++ {
                        if true {
                            if true {
                                if len(v153) >= 2 {
                                    v153 = v153[1:len(v153)]
                                }
                            }
                            v2 -= v3
                            v26 -= v26
                        } else {
                            var v155 float32 = 0.1
                            if len(v153) > 0 {
                                v153[0] = fi152
                            }
                            switch (fi154 % (((fi152 * fi152)) | 1)) {
                            case 0:
                                if false {
                                    v155 += v26
                                    v156 := GenS3{F0: uint16(v3), F1: v3}
                                    _ = v156
                                } else {
                                    var v157 float32 = 3.14159
                                    var v158 int64 = (v32 | v32)
                                    _ = v157
                                    _ = v158
                                }
                            case 32768:
                                var v159 int8 = 100
                                v160 := GenS10{F0: (^(v32)), F1: -572}
                                var v161 uint8 = 255
                                _ = v159
                                _ = v160
                                _ = v161
                            case 128:
                                v162 := GenS12{F0: 128, F1: false}
                                _ = v162
                            default:
                                fmt.Println(v2, v2, v11)
                            }
                            v163 := GenS8{F0: 1, F1: genFn0("hello", v11), F2: 74}
                            _ = v155
                            _ = v163
                        }
                        v3 += (v3 | v3)
                        _ = fi154
                    }
                    fmt.Println(v31)
                    _ = fi152
                    _ = v153
                }
            }
        } else {
            for fi164 := 0; fi164 < 15; fi164++ {
                v32 = v32
                var v165 int32 = -2147483648
                _ = fi164
                _ = v165
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v11
    _ = v2
    _ = v26
    _ = v3
    _ = v31
    _ = v32
}
