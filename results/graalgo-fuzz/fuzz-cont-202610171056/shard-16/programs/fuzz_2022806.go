package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 int8; F1 uint32; F2 uint64 }
type GenS3 struct { F0 int64; F1 string }
type GenS4 struct { F0 uint64; F1 int16 }
type GenS5 struct { F0 uint64; F1 int }
type GenS6 struct { F0 int }
type GenS7 struct { F0 string; F1 int }
type GenS8 struct { F0 int16; F1 int32 }
type GenS9 struct { F0 uint8; F1 uint64; F2 string }
type GenS10 struct { F0 int8; F1 uint32; F2 int8 }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 int16; F1 uint16 }
type GenS13 struct { F0 uint8; F1 int64; F2 int }
type GenS14 struct { F0 uint64; F1 int32 }
type GenS15 struct { F0 int8; F1 int; F2 int8 }
type GenS16 struct { F0 uint; F1 uint32 }
type GenS17 struct { F0 int; F1 int }
type GenS18 struct { F0 int32; F1 int }
type GenS19 struct { F0 uint8; F1 uint32 }
type GenS20 struct { F0 int32; F1 string; F2 uint16 }
type GenS21 struct { F0 bool; F1 uint }
type GenS22 struct { F0 int16; F1 uint16; F2 int16 }
func genFn0(p0 int8) int16 {
    if true {
        v6 := GenS1{F0: 0, F1: 65535}
        _ = v6
    } else {
        var v7 int64 = 42
        _ = v7
    }
    for fi8 := 0; fi8 < 19; fi8++ {
        for fi9 := 0; fi9 < 16; fi9++ {
            if true {
                if false {
                    if (true && true) {
                        var v10 int8 = (p0 * p0)
                        p0 = (p0 & (p0 >> ((296) & 63)))
                        _ = v10
                    } else {
                        p0 -= (p0 - (^(p0)))
                        p0 += -1
                    }
                } else {
                    var v11 uint64 = 851
                    v12 := []int{-378, fi8}
                    _ = v11
                    _ = v12
                }
                fmt.Println(fi9, fi8)
                var v13 uint = 65535
                _ = v13
            }
            var v14 uint64 = 0
            _ = fi9
            _ = v14
        }
        _ = fi8
    }
    p0 = (^(p0))
    _ = p0
    return -1
}
func genFn1() uint16 {
    if false {
        var v15 int8 = 127
        var v16 bool = (!(-705 > 32767))
        _ = v15
        _ = v16
    } else {
        if (-32768 == 1000) {
            if (1 <= 0) {
                var v17 float32 = 0.0
                _ = v17
            } else {
                var v18 uint16 = 65535
                v19 := GenS1{F0: 42, F1: 65535}
                v18 += (v18 - (v18 / ((2) | 1)))
                _ = v18
                _ = v19
            }
            var v20 int8 = -89
            if false {
                v21 := GenS2{F0: v20, F1: 65535, F2: 18446744073709551615}
                var v22 int64 = 1
                _ = v21
                _ = v22
            }
            _ = v20
        }
        v23 := []int{42, 42}
        var v24 uint = 42
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 20; fi25++ {
        if false {
            for fi26 := 0; fi26 < 4; fi26++ {
                switch (fi26 / ((fi26) | 1)) {
                case 9223372036854775807:
                    if (42 >= 18446744073709551615) {
                        fmt.Println(fi26)
                    } else {
                        v27 := GenS3{F0: -772, F1: "hello"}
                        if ((1.5 - 0.0) < (float32((0.1 + 1.5)) - 3.14159)) {
                            var v28 int = 358
                            if ((1 <= 2) || (!true)) {
                                switch (fi26 / (((-(fi25))) | 1)) {
                                case 822:
                                    var v29 uint = 308
                                    _ = v29
                                case -757:
                                    v28 = (fi26 & (v28 / (((fi26 << ((55) & 63))) | 1)))
                                }
                                v30 := []int{(v28 >> ((1) & 63)), (v28 & (fi25 * -2147483648)), (fi26 - fi26)}
                                _ = v30
                            }
                            _ = v28
                        } else {
                            switch ("foo bar" + ("a" + ("a" + "a"))) {
                            case "\n\t":
                                if false {
                                    if (!true) {
                                        var v31 float32 = 3.14159
                                        v31 += v31
                                        var v32 float32 = float32(float64(0.0))
                                        _ = v31
                                        _ = v32
                                    } else {
                                        var v33 uint16 = 292
                                        fmt.Println(v33)
                                        var v34 float64 = ((((float64((float32(0.1) + (3.14159 * -2.25))) + 3.14159) / ((1.5 * float64(1e10)) - ((1.5 - 1e10) + 0.1))) - ((-2.25 * 0.1) / 1.5)) / ((3.14159 + (0.0 + (0.0 - ((1e10 - (1.5 / 1e10)) * 0.1)))) * 1.5))
                                        switch (fi26 * (fi25 & (fi26 - 0))) {
                                        case 32768:
                                            v35 := GenS2{F0: 100, F1: 2, F2: 2}
                                            _ = v35
                                        }
                                        _ = v33
                                        _ = v34
                                    }
                                    v36 := GenS4{F0: 1, F1: -1}
                                    _ = v36
                                } else {
                                    var v37 float32 = 1.5
                                    _ = v37
                                }
                            case "a":
                                var v38 int8 = 0
                                var v39 bool = (!true)
                                _ = v38
                                _ = v39
                            default:
                                var v40 int8 = -101
                                _ = v40
                            }
                            var v41 int = 9223372036854775807
                            var v42 uint16 = 42
                            _ = v41
                            _ = v42
                        }
                        _ = v27
                    }
                default:
                    var v43 string = ("hello" + "foo bar")
                    _ = v43
                }
                v44 := GenS4{F0: 2, F1: 0}
                _ = fi26
                _ = v44
            }
            for fi45 := 0; fi45 < 17; fi45++ {
                var v46 uint64 = 2
                if (!((-635 == 1) && true)) {
                    v46 = v46
                }
                v47 := GenS5{F0: v46, F1: (fi25 ^ (fi45 & fi25))}
                _ = fi45
                _ = v46
                _ = v47
            }
            if (false && false) {
                v48 := []int{(fi25 + 0), fi25, (fi25 + (-(fi25))), fi25}
                var v49 uint32 = 0
                v48 = append(v48, (fi25 - 128))
                _ = v48
                _ = v49
            }
        } else {
            for fi50 := 0; fi50 < 13; fi50++ {
                v51 := GenS6{F0: (fi25 + (fi50 / (((fi50 << ((923) & 63))) | 1)))}
                var v52 int32 = -615
                v52 = (v52 & v52)
                _ = fi50
                _ = v51
                _ = v52
            }
            if false {
                for fi53 := 0; fi53 < 3; fi53++ {
                    fmt.Println(fi53)
                    var v54 uint16 = 32767
                    _ = fi53
                    _ = v54
                }
            }
        }
        _ = fi25
    }
    v55 := GenS4{F0: 18446744073709551615, F1: 42}
    var v56 int = -9223372036854775808
    if true {
        var v57 bool = false
        v56 = v56
        _ = v57
    } else {
        v58 := []int{v56, (v56 ^ v56), v56}
        for fi59 := 0; fi59 < 17; fi59++ {
            switch "" {
            case "\n\t":
                _ = len(v58)
            case "a":
                for fi60 := 0; fi60 < 12; fi60++ {
                    if true {
                        if len(v58) >= 2 {
                            v58 = v58[1:len(v58)]
                        }
                        switch "a" {
                        case "foo bar":
                            if false {
                                fmt.Println(fi60, v56)
                            } else {
                                switch -1 {
                                case 42:
                                    if (42 < 196) {
                                        v58 = append(v58, (fi60 >> ((511) & 63)))
                                        if true {
                                            var v61 uint64 = 464
                                            var v62 uint32 = 42
                                            if true {
                                                var v63 string = "hello"
                                                switch (v63 + v63) {
                                                case "hello":
                                                    v61 -= (v61 * (v61 ^ v61))
                                                case "foo bar":
                                                    v64 := GenS7{F0: "hello", F1: v56}
                                                    _ = v64
                                                case "\n\t":
                                                    fmt.Println(v61, v62, fi59)
                                                default:
                                                    v58 = append(v58, (fi59 % ((32767) | 1)))
                                                }
                                                _ = v63
                                            }
                                            _ = v61
                                            _ = v62
                                        } else {
                                            v65 := GenS8{F0: genFn0(-88), F1: 0}
                                            switch fi60 {
                                            case 553:
                                                v56 -= v56
                                            case 2147483647:
                                                fmt.Println(fi59)
                                            case 127:
                                                v58 = append(v58, (^(fi59)))
                                            }
                                            fmt.Println(v56, fi59)
                                            _ = v65
                                        }
                                    } else {
                                        if (222 <= -1) {
                                            if len(v58) > 0 {
                                                v58[0] = -729
                                            }
                                        }
                                    }
                                case 32767:
                                    v66 := GenS9{F0: 128, F1: 2, F2: "foo bar"}
                                    _ = v66
                                default:
                                    var v67 int8 = 0
                                    _ = v67
                                }
                                if (42 == 65535) {
                                    if false {
                                        var v68 string = "foo bar"
                                        var v69 string = "hello"
                                        var v70 uint = 518
                                        _ = v68
                                        _ = v69
                                        _ = v70
                                    }
                                } else {
                                    var v71 int64 = 9223372036854775807
                                    _ = v71
                                }
                            }
                            var v72 uint64 = 303
                            _ = v72
                        case "":
                            v73 := GenS10{F0: 100, F1: 89, F2: -86}
                            _ = v73
                        }
                        if len(v58) >= 2 {
                            v58 = v58[1:len(v58)]
                        }
                    }
                    var v74 int64 = -172
                    switch fi59 {
                    case -1:
                        var v75 string = ("hello" + (("foo bar" + "a") + (("a" + ("a" + "a")) + ((("foo bar" + "a") + "foo bar") + "hello"))))
                        _ = v75
                    case 1:
                        if true {
                            v58 = append(v58, 1)
                            var v76 int = -1
                            _ = v76
                        } else {
                            if (2 <= 255) {
                                var v77 int32 = -763
                                _ = v77
                            } else {
                                v56 = (fi59 / (((fi59 / ((fi60) | 1))) | 1))
                                v78 := GenS7{F0: "", F1: v56}
                                switch "hello" {
                                case "a":
                                    var v79 uint32 = 233
                                    _ = v79
                                case "hello":
                                    fmt.Println(v74)
                                }
                                _ = v78
                            }
                            switch "" {
                            case "\n\t":
                                if false {
                                    if (42 < -1) {
                                        switch "a" {
                                        case "hello":
                                            v74 = (v74 % ((v74) | 1))
                                        case "\n\t":
                                            v80 := GenS11{F0: 42}
                                            _ = v80
                                        case "":
                                            _ = len(v58)
                                        }
                                    } else {
                                        if len(v58) >= 2 {
                                            v58 = v58[1:len(v58)]
                                        }
                                        v74 -= (v74 % ((-9223372036854775808) | 1))
                                    }
                                    v58 = append(v58, (^(v56)))
                                    if false {
                                        if (("" + "foo bar") >= ((("hello" + "foo bar") + "\n\t") + "foo bar")) {
                                            v81 := GenS12{F0: 1000, F1: 711}
                                            if (-2.25 < 1e10) {
                                                var v82 int = 4294967295
                                                _ = v82
                                            } else {
                                                v83 := GenS13{F0: 191, F1: (-(v74)), F2: (-(v56))}
                                                v56 += 42
                                                v74 = (-(v74))
                                                var v84 uint16 = 448
                                                switch "a" {
                                                case "hello":
                                                    v85 := GenS1{F0: 614, F1: 95}
                                                    _ = v85
                                                default:
                                                    if ((^(v74)) < (-(v74))) {
                                                        _ = len(v58)
                                                        v86 := GenS14{F0: 946, F1: 535}
                                                        _ = v86
                                                    }
                                                }
                                                v84 = (v84 << ((4294967295) & 63))
                                                if true {
                                                    v58 = append(v58, fi59)
                                                }
                                                _ = v83
                                                _ = v84
                                            }
                                            if true {
                                                if len(v58) > 0 {
                                                    v58[0] = -2147483648
                                                }
                                            }
                                            _ = v81
                                        }
                                        var v87 uint16 = 65535
                                        _ = v87
                                    }
                                    var v88 uint16 = 255
                                    v56 = v56
                                    _ = v88
                                } else {
                                    var v89 float64 = -2.25
                                    _ = v89
                                }
                            case "a":
                                if (100 <= -1) {
                                    if true {
                                        v90 := GenS1{F0: 920, F1: 42}
                                        _ = v90
                                    }
                                } else {
                                    v56 -= v56
                                    if len(v58) >= 2 {
                                        v58 = v58[1:len(v58)]
                                    }
                                }
                            case "":
                                var v91 float32 = 1e10
                                _ = v91
                            }
                            var v92 int16 = -32768
                            _ = v92
                        }
                        if false {
                            v93 := GenS4{F0: 4294967295, F1: 804}
                            var v94 int = v56
                            _ = v93
                            _ = v94
                        } else {
                            var v95 uint64 = 2
                            var v96 int32 = int32(fi60)
                            _ = v95
                            _ = v96
                        }
                        fmt.Println(fi60, fi60)
                    case 42:
                        switch (v56 + (fi59 | fi60)) {
                        case -572:
                            var v97 uint32 = 2147483648
                            _ = v97
                        }
                    default:
                        v56 -= v56
                    }
                    _ = fi60
                    _ = v74
                }
            case "":
                var v98 bool = false
                _ = v98
            default:
                if len(v58) > 0 {
                    v58[0] = (fi59 >> ((4294967295) & 63))
                }
            }
            if (false && true) {
                v56 = v56
                var v99 float32 = 0.1
                var v100 uint8 = 128
                _ = v100
                _ = v99
            } else {
                var v101 int32 = -806
                v56 -= (fi59 % ((fi59) | 1))
                var v102 uint32 = 4294967295
                _ = v101
                _ = v102
            }
            _ = fi59
        }
        if (65535 >= 2) {
            if len(v58) > 0 {
                v58[0] = 32767
            }
        } else {
            if true {
                v58 = append(v58, -689)
                fmt.Println(v56)
                for fi103 := 0; fi103 < 20; fi103++ {
                    for fi104 := 0; fi104 < 11; fi104++ {
                        var v105 float64 = ((((1e10 - (1e10 * (3.14159 / 1.0))) + ((3.14159 * 1.5) / ((1.5 - 3.14159) - float64(float32(((1e10 + float64(float32((1e10 + -2.25)))) + 3.14159)))))) / float64((3.14159 - -2.25))) + (float64(0.1) - 1e10))
                        v58 = append(v58, fi104)
                        v58 = append(v58, -1)
                        _ = fi104
                        _ = v105
                    }
                    v56 = -2147483648
                    for fi106 := 0; fi106 < 10; fi106++ {
                        var v107 uint8 = 127
                        v107 = v107
                        if false {
                            v108 := GenS2{F0: 100, F1: 65535, F2: 638}
                            fmt.Println(v107, v56, fi106)
                            v107 = 1
                            _ = v108
                        }
                        _ = fi106
                        _ = v107
                    }
                    _ = fi103
                }
            }
            if (0.1 > 1.5) {
                switch (v56 | 0) {
                case 127:
                    v56 -= v56
                case 9223372036854775807:
                    v109 := GenS15{F0: 100, F1: 32767, F2: 0}
                    _ = v109
                case 1:
                    for fi110 := 0; fi110 < 19; fi110++ {
                        if true {
                            var v111 uint16 = 651
                            fmt.Println(fi110, v56)
                            _ = v111
                        }
                        switch v56 {
                        case -706:
                            var v112 uint32 = 2
                            _ = v112
                        case -9223372036854775808:
                            var v113 uint64 = 0
                            _ = v113
                        case 42:
                            v114 := GenS16{F0: 18446744073709551615, F1: 0}
                            _ = v114
                        default:
                            switch v56 {
                            case -2147483648:
                                if len(v58) > 0 {
                                    v58[0] = v56
                                }
                            default:
                                var v115 int = v56
                                _ = v115
                            }
                        }
                        _ = fi110
                    }
                }
                v116 := GenS17{F0: (v56 >> ((1) & 63)), F1: v56}
                _ = v116
            }
        }
        _ = v58
    }
    _ = v55
    _ = v56
    return 357
}

func main() {
    v1 := []int{4294967295, 4294967295, 127, 128}
    v2 := make(map[int]bool)
    v3 := make(map[int]int)
    var v4 int8 = -1
    var v5 uint = 556
    for fi117 := 0; fi117 < 2; fi117++ {
        if len(v1) > 0 {
            v1[0] = (-(fi117))
        }
        if (!(0 < -807)) {
            v5 += v5
            for fi118 := 0; fi118 < 17; fi118++ {
                v5 += v5
                fmt.Println(v4)
                switch 4294967295 {
                case 1:
                    var v119 float64 = ((0.0 * 1.5) + 1.5)
                    _ = v119
                case -2147483648:
                    v4 = 42
                default:
                    var v120 int64 = 42
                    _ = v120
                }
                _ = len(v1)
                v121 := v2[(fi118 ^ fi118)]
                fmt.Println(v121)
                fmt.Println(v121)
                switch "a" {
                case "foo bar":
                    v1 = append(v1, (fi117 << ((v5) & 63)))
                    var v122 uint32 = 65535
                    _ = v122
                case "":
                    v123 := GenS18{F0: -1, F1: (fi117 - 1)}
                    _ = v123
                }
                _ = fi118
                _ = v121
            }
            v124 := GenS2{F0: v4, F1: 1, F2: 1}
            _ = v124
        } else {
            for fi125 := 0; fi125 < 9; fi125++ {
                var v126 uint32 = 2147483648
                _ = fi125
                _ = v126
            }
        }
        for fi127 := 0; fi127 < 20; fi127++ {
            switch "\n\t" {
            case "\n\t":
                switch "\n\t" {
                case "a":
                    var v128 uint = v5
                    _ = v128
                case "":
                    var v129 int64 = 0
                    _ = v129
                case "hello":
                    switch (fi127 * fi127) {
                    case -32:
                        var v130 uint64 = 0
                        _ = v130
                    case 27:
                        _ = len(v1)
                    }
                default:
                    var v131 uint32 = 42
                    _ = v131
                }
            case "":
                fmt.Println(v4, fi127, v4)
            case "hello":
                var v132 int16 = genFn0(v4)
                v133 := GenS7{F0: ("" + ((("" + (("" + "") + ("\n\t" + "\n\t"))) + "hello") + ("foo bar" + ("foo bar" + "foo bar")))), F1: 2147483647}
                _ = v132
                _ = v133
            default:
                delete(v2, 32768)
                switch (fi117 >> (((v5 - (^(v5)))) & 63)) {
                case 127:
                    v134 := v3[int(v4)]
                    _ = v134
                case 128:
                    v4 -= (v4 - (v4 + v4))
                default:
                    var v135 float64 = 1.5
                    _ = v135
                }
            }
            switch "" {
            case "foo bar":
                var v136 int64 = 42
                _ = v136
            case "":
                fmt.Println(v4, fi127)
            case "hello":
                if len(v1) > 0 {
                    v1[0] = -2147483648
                }
            default:
                if false {
                    var v137 int32 = -677
                    _ = v137
                }
            }
            var v138 string = ""
            switch (^(fi127)) {
            case 0:
                v139 := GenS19{F0: 127, F1: 1}
                if (!(!(3.14159 != 1e10))) {
                    var v140 uint32 = 0
                    var v141 uint32 = (v140 + v140)
                    var v142 float64 = (0.1 * float64(float32(0.1)))
                    switch fi117 {
                    case 807:
                        if (360 >= 2) {
                            if (0.1 < 1.5) {
                                var v143 uint8 = 128
                                v143 = 222
                                v141 = (v140 & 12)
                                _ = v143
                            } else {
                                v1 = append(v1, fi127)
                                if (v138 != v138) {
                                    _ = len(v3)
                                    delete(v2, fi117)
                                } else {
                                    var v144 int64 = -547
                                    switch fi127 {
                                    case 42:
                                        v141 = (v140 ^ v141)
                                    case 2147483647:
                                        v138 = "a"
                                    case 0:
                                        var v145 float32 = 1.5
                                        _ = v145
                                    }
                                    _ = v144
                                }
                                var v146 int = -9223372036854775808
                                var v147 int64 = -1
                                if (2 <= 32767) {
                                    v1 = append(v1, (-(v146)))
                                    v141 = v140
                                    var v148 int8 = v4
                                    if (v146 < 127) {
                                        v5 -= uint(v140)
                                    } else {
                                        if (true || false) {
                                            v149 := v3[v146]
                                            v150 := GenS19{F0: 116, F1: v140}
                                            switch v138 {
                                            case "foo bar":
                                                _ = len(v2)
                                            case "a":
                                                var v151 int64 = 1
                                                _ = v151
                                            }
                                            _ = v149
                                            _ = v150
                                        } else {
                                            _ = len(v3)
                                        }
                                        var v152 uint = v5
                                        fmt.Println(v147)
                                        _ = v152
                                    }
                                    v1 = append(v1, (fi127 & v146))
                                    _ = v148
                                }
                                _ = v146
                                _ = v147
                            }
                            fmt.Println(v140)
                        }
                    case -365:
                        v5 += 1
                    case 729:
                        var v153 int64 = 9223372036854775807
                        _ = v153
                    default:
                        v138 = v138
                    }
                    var v154 uint = (v5 / (((^(v5))) | 1))
                    _ = v140
                    _ = v141
                    _ = v142
                    _ = v154
                } else {
                    var v155 int8 = v4
                    _ = v155
                }
                _ = v139
            default:
                if true {
                    if len(v1) > 0 {
                        v1[0] = (fi117 ^ (fi127 | fi117))
                    }
                    var v156 uint32 = 254
                    v157 := GenS3{F0: 42, F1: "a"}
                    _ = v156
                    _ = v157
                }
            }
            _ = fi127
            _ = v138
        }
        _ = fi117
    }
    v4 += v4
    v5 += v5
    var v158 string = ""
    if len(v1) > 0 {
        v1[0] = 32768
    }
    v159 := v2[128]
    for fi160 := 0; fi160 < 7; fi160++ {
        var v161 float64 = 1e10
        _ = fi160
        _ = v161
    }
    if len(v1) > 0 {
        v1[0] = -1
    }
    var v162 float32 = ((0.0 * (3.14159 - 0.1)) * -2.25)
    v1 = append(v1, 42)
    var v163 uint8 = 116
    var v164 float64 = 0.1
    for fi165 := 0; fi165 < 7; fi165++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi165
    }
    v164 -= (v164 * 1.5)
    if v159 {
        v159 = true
    } else {
        switch (v158 + v158) {
        case "":
            v166 := GenS6{F0: 32767}
            _ = v166
        case "a":
            v162 = -2.25
        case "\n\t":
            v167 := GenS20{F0: 42, F1: "", F2: 255}
            _ = v167
        default:
            var v168 int16 = genFn0(-128)
            _ = v168
        }
        v169 := GenS10{F0: (v4 / ((v4) | 1)), F1: 65535, F2: -128}
        _ = v169
    }
    if (false || v159) {
        v170 := GenS21{F0: false, F1: (v5 - 255)}
        var v171 int16 = 42
        var v172 uint32 = uint32(v5)
        _ = v170
        _ = v171
        _ = v172
    } else {
        var v173 int8 = (^(v4))
        v174 := GenS22{F0: 1, F1: 65535, F2: 1000}
        _ = v173
        _ = v174
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v158:", v158)
    fmt.Println("v159:", v159)
    fmt.Println("v162:", v162)
    fmt.Println("v163:", v163)
    fmt.Println("v164:", v164)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v158
    _ = v159
    _ = v162
    _ = v163
    _ = v164
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
