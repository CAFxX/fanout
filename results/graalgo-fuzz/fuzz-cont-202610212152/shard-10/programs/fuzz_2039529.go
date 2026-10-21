package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 int64; F1 int32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int16; F1 int16; F2 uint16 }
type GenS6 struct { F0 uint16; F1 bool; F2 int8 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int8; F1 int; F2 uint }
type GenS9 struct { F0 uint; F1 int16; F2 bool }
type GenS10 struct { F0 int64; F1 uint8; F2 uint8 }
type GenS11 struct { F0 uint64; F1 string }
type GenS12 struct { F0 int64; F1 int32; F2 uint16 }
type GenS13 struct { F0 uint16; F1 string; F2 int8 }
type GenS14 struct { F0 uint16; F1 uint16 }
type GenS15 struct { F0 int }
type GenS16 struct { F0 uint8; F1 uint8; F2 uint16 }
type GenS17 struct { F0 int8; F1 uint; F2 int }
type GenS18 struct { F0 int16; F1 uint64; F2 uint64 }
type GenS19 struct { F0 int8; F1 int16; F2 int }
func genFn0() int {
    if true {
        for fi7 := 0; fi7 < 9; fi7++ {
            for fi8 := 0; fi8 < 13; fi8++ {
                var v9 float32 = 1e10
                v9 -= (v9 * v9)
                _ = fi8
                _ = v9
            }
            v10 := []int{(fi7 & 159), fi7, fi7}
            var v11 int64 = 157
            _ = fi7
            _ = v10
            _ = v11
        }
        if false {
            v12 := GenS1{F0: 2, F1: -945}
            var v13 int16 = -32768
            var v14 float64 = 0.0
            _ = v12
            _ = v13
            _ = v14
        } else {
            if ("foo bar" <= "foo bar") {
                v15 := GenS2{F0: 1, F1: -32768}
                for fi16 := 0; fi16 < 1; fi16++ {
                    var v17 uint32 = 100
                    v17 -= (v17 >> ((372) & 63))
                    for fi18 := 0; fi18 < 11; fi18++ {
                        v19 := GenS2{F0: -128, F1: 32767}
                        _ = fi18
                        _ = v19
                    }
                    _ = fi16
                    _ = v17
                }
                for fi20 := 0; fi20 < 11; fi20++ {
                    v21 := GenS2{F0: 0, F1: -32768}
                    if (!(true && (false || (255 < 0)))) {
                        v22 := GenS2{F0: -128, F1: 1000}
                        if true {
                            fmt.Println(fi20)
                            var v23 float64 = -2.25
                            fmt.Println(v23, v23)
                            _ = v23
                        } else {
                            switch ("foo bar" + "a") {
                            case "a":
                                v24 := []int{}
                                _ = v24
                            case "\n\t":
                                v25 := GenS1{F0: 42, F1: 0}
                                _ = v25
                            case "foo bar":
                                v26 := []int{(-(fi20)), 4294967295}
                                _ = v26
                            }
                            var v27 float64 = (1.5 * -2.25)
                            if (!(((v27 / 1.5) / v27) == v27)) {
                                v28 := GenS1{F0: 2, F1: -2147483648}
                                v29 := GenS3{F0: 9223372036854775807, F1: 42}
                                v27 = (0.0 - ((((float64(1.5) * 0.1) * ((float64(0.0) + 0.0) * 3.14159)) - 1.5) - 1e10))
                                _ = v28
                                _ = v29
                            }
                            _ = v27
                        }
                        _ = v22
                    } else {
                        var v30 uint8 = 255
                        _ = v30
                    }
                    _ = fi20
                    _ = v21
                }
                _ = v15
            } else {
                var v31 uint32 = 0
                fmt.Println(v31)
                _ = v31
            }
            switch "foo bar" {
            case "":
                v32 := GenS4{F0: 2147483647}
                _ = v32
            case "foo bar":
                v33 := []int{42, 128}
                _ = v33
            }
        }
    }
    if (false && false) {
        if (false && (255 != 1)) {
            v34 := []int{128, 735}
            _ = v34
        }
        v35 := []int{42, -9223372036854775808, 42}
        var v36 string = "hello"
        _ = v35
        _ = v36
    } else {
        switch "hello" {
        case "hello":
            var v37 bool = true
            _ = v37
        case "foo bar":
            if (1 <= -1) {
                var v38 uint = 18446744073709551615
                _ = v38
            }
        case "":
            for fi39 := 0; fi39 < 13; fi39++ {
                var v40 uint8 = 253
                _ = fi39
                _ = v40
            }
        default:
            var v41 uint64 = 2
            _ = v41
        }
        var v42 float32 = (0.1 + (float32(1e10) + 0.1))
        if ((true || true) && ("a" <= "a")) {
            v42 = (1e10 + 3.14159)
            if (!false) {
                for fi43 := 0; fi43 < 7; fi43++ {
                    v42 = (0.0 / 1.5)
                    v44 := []int{(fi43 / (((fi43 << ((65535) & 63))) | 1)), fi43, -9223372036854775808}
                    var v45 uint = 42
                    _ = fi43
                    _ = v44
                    _ = v45
                }
                if (float32(float64(v42)) >= (0.0 / v42)) {
                    var v46 string = "hello"
                    _ = v46
                }
                v47 := []int{-1, 128, 957, 32767}
                _ = v47
            }
            if true {
                for fi48 := 0; fi48 < 13; fi48++ {
                    if true {
                        if false {
                            var v49 int32 = -1
                            _ = v49
                        }
                        fmt.Println(fi48)
                        switch "" {
                        case "foo bar":
                            var v50 int16 = -193
                            _ = v50
                        }
                    } else {
                        v42 = 0.1
                        var v51 uint64 = 4294967295
                        _ = v51
                    }
                    v52 := GenS4{F0: -2147483648}
                    _ = fi48
                    _ = v52
                }
            }
        }
        _ = v42
    }
    var v53 int16 = 32767
    fmt.Println(v53)
    for fi54 := 0; fi54 < 16; fi54++ {
        for fi55 := 0; fi55 < 18; fi55++ {
            fmt.Println(fi55)
            if true {
                if (18446744073709551615 >= 488) {
                    switch (fi54 & int(v53)) {
                    case -9223372036854775808:
                        v56 := []int{}
                        _ = v56
                    case 109:
                        v57 := GenS2{F0: 100, F1: v53}
                        _ = v57
                    case 128:
                        v58 := GenS5{F0: v53, F1: v53, F2: 65535}
                        _ = v58
                    default:
                        if (true || false) {
                            v53 = (v53 << ((653) & 63))
                            v59 := GenS5{F0: v53, F1: (v53 | v53), F2: 990}
                            if ((!((false || ((((("hello" + "") + ("a" + "")) + "") != ("\n\t" + "")) || false)) && true)) && ("a" != "foo bar")) {
                                if true {
                                    if ("hello" != "foo bar") {
                                        v53 += -32768
                                        v60 := GenS6{F0: 713, F1: true, F2: 100}
                                        _ = v60
                                    } else {
                                        fmt.Println(fi55)
                                        if (227 != 167) {
                                            v61 := []int{(fi55 << ((4294967295) & 63)), (fi54 ^ int(v53)), fi55, fi54}
                                            v53 += -32768
                                            v53 += -32768
                                            v62 := GenS2{F0: -128, F1: v53}
                                            v63 := GenS4{F0: 42}
                                            _ = v61
                                            _ = v62
                                            _ = v63
                                        } else {
                                            var v64 int64 = -1
                                            _ = v64
                                        }
                                    }
                                } else {
                                    switch (fi54 - (^(fi55))) {
                                    case -2147483648:
                                        var v65 int32 = 1
                                        _ = v65
                                    case -275:
                                        var v66 float64 = 1e10
                                        var v67 int8 = -1
                                        var v68 int8 = (v67 % ((-128) | 1))
                                        _ = v66
                                        _ = v67
                                        _ = v68
                                    default:
                                        v69 := []int{fi55}
                                        _ = v69
                                    }
                                }
                                v70 := GenS7{F0: 0}
                                _ = v70
                            } else {
                                v71 := GenS8{F0: 51, F1: (fi55 / ((-1) | 1)), F2: 2}
                                fmt.Println(v53, fi54)
                                _ = v71
                            }
                            _ = v59
                        }
                    }
                    if false {
                        v53 = v53
                        if false {
                            v53 -= (v53 ^ (v53 + (v53 - (v53 ^ (v53 / ((v53) | 1))))))
                        } else {
                            if false {
                                v72 := GenS9{F0: 916, F1: v53, F2: (1 < 1)}
                                var v73 int16 = v53
                                _ = v72
                                _ = v73
                            } else {
                                if (!(!(true || true))) {
                                    var v74 int16 = (v53 << ((126) & 63))
                                    var v75 int = fi55
                                    _ = v74
                                    _ = v75
                                } else {
                                    v53 -= 1
                                    v53 -= v53
                                }
                            }
                        }
                        var v76 int16 = 42
                        switch (fi55 / ((fi54) | 1)) {
                        case 32767:
                            if true {
                                var v77 float32 = (1e10 + (0.1 / 1e10))
                                v77 = v77
                                var v78 uint = 2
                                v79 := GenS2{F0: -1, F1: v76}
                                _ = v77
                                _ = v78
                                _ = v79
                            } else {
                                v53 = v53
                                switch (fi55 + (fi55 ^ 127)) {
                                case 128:
                                    var v80 uint = 255
                                    _ = v80
                                case 42:
                                    var v81 uint8 = 0
                                    _ = v81
                                }
                            }
                            var v82 bool = true
                            v83 := GenS8{F0: 42, F1: 801, F2: 1}
                            _ = v82
                            _ = v83
                        case 0:
                            v84 := GenS10{F0: 1, F1: 98, F2: 127}
                            _ = v84
                        default:
                            v85 := []int{(fi55 ^ fi55)}
                            _ = v85
                        }
                        _ = v76
                    } else {
                        v86 := GenS11{F0: 115, F1: "\n\t"}
                        _ = v86
                    }
                    var v87 float32 = 1e10
                    _ = v87
                }
            }
            var v88 string = "hello"
            if (!((65535 >= 2) || (!true))) {
                switch v88 {
                case "":
                    if true {
                        if false {
                            var v89 int64 = 0
                            if ((1e10 / float64(float32(1e10))) >= 1e10) {
                                var v90 uint16 = 369
                                v91 := GenS10{F0: (v89 - -9223372036854775808), F1: 127, F2: 42}
                                _ = v90
                                _ = v91
                            } else {
                                var v92 uint8 = 1
                                v93 := GenS12{F0: v89, F1: int32(v92), F2: 1}
                                _ = v92
                                _ = v93
                            }
                            v94 := []int{(fi54 ^ (fi54 & (fi54 << ((65535) & 63)))), 128, 450}
                            _ = v89
                            _ = v94
                        }
                    }
                }
            } else {
                if true {
                    v95 := []int{}
                    v95 = append(v95, (fi54 + (fi54 + 42)))
                    if false {
                        v95 = append(v95, (-(fi54)))
                        v88 = "foo bar"
                        if ((v53 << ((397) & 63)) == (^(v53))) {
                            var v96 float32 = -2.25
                            v88 = "foo bar"
                            var v97 uint64 = 42
                            if len(v95) > 0 {
                                v95[0] = (^(fi55))
                            }
                            _ = v96
                            _ = v97
                        } else {
                            var v98 int64 = -9223372036854775808
                            if true {
                                fmt.Println(v98)
                                v99 := GenS13{F0: 1, F1: "\n\t", F2: -1}
                                if (!((v53 | 1000) <= (v53 | -1))) {
                                    var v100 int32 = 1
                                    _ = v100
                                } else {
                                    var v101 int = fi54
                                    switch fi54 {
                                    case -1:
                                        fmt.Println(v88, v98)
                                    case -641:
                                        var v102 bool = true
                                        _ = v102
                                    case 9223372036854775807:
                                        if (!false) {
                                            v103 := GenS12{F0: v98, F1: -37, F2: 0}
                                            v104 := GenS14{F0: 1, F1: 1}
                                            switch (fi55 ^ (fi54 / (((-(fi55))) | 1))) {
                                            case 4294967295:
                                                var v105 float32 = 3.14159
                                                _ = v105
                                            }
                                            _ = v103
                                            _ = v104
                                        }
                                    default:
                                        var v106 uint8 = 2
                                        _ = v106
                                    }
                                    _ = v101
                                }
                                _ = v99
                            }
                            v53 += v53
                            _ = v98
                        }
                    } else {
                        switch v88 {
                        case "\n\t":
                            if false {
                                _ = len(v95)
                                var v107 int = (fi55 & (fi54 >> ((4294967295) & 63)))
                                fmt.Println(fi55, v53)
                                fmt.Println(v107)
                                if false {
                                    switch fi54 {
                                    case 127:
                                        switch ("foo bar" + (v88 + (v88 + ("" + ((("a" + "foo bar") + (v88 + v88)) + v88))))) {
                                        case "a":
                                            fmt.Println(v88, fi54, fi54)
                                        case "\n\t":
                                            var v108 float32 = 1e10
                                            _ = v108
                                        }
                                    case 1:
                                        var v109 uint16 = 255
                                        _ = v109
                                    }
                                    if true {
                                        v107 += fi55
                                        var v110 int8 = -128
                                        _ = v110
                                    } else {
                                        fmt.Println(fi55, fi54)
                                        var v111 int8 = 92
                                        v88 = v88
                                        v107 = 394
                                        _ = v111
                                    }
                                    switch "foo bar" {
                                    case "\n\t":
                                        _ = len(v95)
                                    case "a":
                                        v112 := GenS7{F0: 2}
                                        _ = v112
                                    case "":
                                        v107 = -640
                                    }
                                } else {
                                    fmt.Println(fi55, fi54, fi54)
                                    v113 := GenS15{F0: (v107 << ((2) & 63))}
                                    _ = v113
                                }
                                _ = v107
                            }
                        default:
                            var v114 int16 = 42
                            _ = v114
                        }
                    }
                    _ = v95
                } else {
                    var v115 string = (v88 + v88)
                    fmt.Println(v115)
                    switch (fi55 >> ((855) & 63)) {
                    case 9223372036854775807:
                        var v116 uint8 = 251
                        _ = v116
                    default:
                        v115 = (v88 + (("hello" + "a") + (((v115 + v88) + v88) + v115)))
                    }
                    _ = v115
                }
                v117 := GenS14{F0: 2, F1: 2}
                v118 := GenS16{F0: 16, F1: 78, F2: 255}
                _ = v117
                _ = v118
            }
            _ = fi55
            _ = v88
        }
        if (105 < 896) {
            v119 := GenS12{F0: 128, F1: 2147483647, F2: 628}
            v120 := GenS17{F0: 100, F1: 1, F2: (fi54 % ((fi54) | 1))}
            _ = v119
            _ = v120
        }
        _ = fi54
    }
    _ = v53
    return 4294967295
}
func genFn1(p0 int32, p1 uint, p2 int32) int8 {
    var v121 int = -9223372036854775808
    if (true && false) {
        switch "" {
        case "\n\t":
            for fi122 := 0; fi122 < 16; fi122++ {
                switch 128 {
                case 111:
                    var v123 uint16 = 0
                    _ = v123
                case 32768:
                    if (!(!((v121 >> ((p1) & 63)) < (v121 - v121)))) {
                        p0 += (p2 >> ((p1) & 63))
                        if false {
                            switch ("" + "a") {
                            case "":
                                for fi124 := 0; fi124 < 12; fi124++ {
                                    fmt.Println(fi122, fi124)
                                    if false {
                                        if false {
                                            v121 += fi124
                                        }
                                        v121 -= v121
                                        switch "hello" {
                                        case "foo bar":
                                            v121 += fi124
                                        case "\n\t":
                                            p2 = (p0 & p0)
                                        case "hello":
                                            var v125 uint64 = 42
                                            _ = v125
                                        }
                                        var v126 uint64 = 42
                                        _ = v126
                                    }
                                    _ = fi124
                                }
                            case "\n\t":
                                v127 := []int{708}
                                _ = v127
                            default:
                                fmt.Println(p0)
                            }
                            p0 = 2147483647
                            if (!(false && false)) {
                                var v128 uint16 = 255
                                v129 := GenS10{F0: -9223372036854775808, F1: 0, F2: uint8(p1)}
                                _ = v128
                                _ = v129
                            }
                        } else {
                            if (-128 < 117) {
                                var v130 int64 = 660
                                _ = v130
                            } else {
                                fmt.Println(fi122, v121)
                                var v131 int64 = 9223372036854775807
                                _ = v131
                            }
                            v121 = genFn0()
                        }
                    } else {
                        switch fi122 {
                        case 127:
                            p1 = 2
                        }
                    }
                }
                v121 += fi122
                _ = fi122
            }
        default:
            var v132 float32 = (((0.0 / 3.14159) / 1.0) + 1.5)
            _ = v132
        }
        var v133 int = (v121 ^ (v121 * v121))
        var v134 uint16 = 615
        _ = v133
        _ = v134
    }
    v135 := GenS16{F0: 244, F1: 14, F2: 775}
    if (9223372036854775807 <= 589) {
        for fi136 := 0; fi136 < 1; fi136++ {
            fmt.Println(fi136, fi136, p0)
            for fi137 := 0; fi137 < 12; fi137++ {
                var v138 int64 = 42
                _ = fi137
                _ = v138
            }
            _ = fi136
        }
        var v139 int32 = p0
        v139 -= -411
        _ = v139
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v121
    _ = v135
    return 0
}

func main() {
    var v1 uint16 = 670
    var v2 int = 0
    var v3 int16 = -1
    v4 := []string{}
    var v5 int16 = v3
    v6 := make(map[uint32]uint32)
    for fi140 := 0; fi140 < 1; fi140++ {
        var v141 int = genFn0()
        _ = fi140
        _ = v141
    }
    var v142 int16 = (v5 & v3)
    var v143 int64 = 0
    v144 := GenS6{F0: (^(v1)), F1: (false && true), F2: genFn1(42, 4294967295, 104)}
    v145 := GenS13{F0: v1, F1: "a", F2: 127}
    fmt.Println(v142, v2, v1)
    v146 := GenS18{F0: (-(v3)), F1: 229, F2: uint64(v1)}
    var v147 uint16 = v1
    v1 += v1
    if false {
        var v148 uint16 = 255
        _ = v148
    } else {
        v149 := GenS14{F0: v147, F1: 255}
        var v150 uint64 = 184
        _ = v149
        _ = v150
    }
    var v151 float32 = 1e10
    var v152 float64 = 1.5
    switch "foo bar" {
    case "":
        switch 127 {
        case -2147483648:
            _ = len(v6)
        case 32768:
            _ = len(v4)
        }
    case "hello":
        v153 := GenS19{F0: -126, F1: (v5 * (v3 | 1000)), F2: 42}
        _ = v153
    case "a":
        var v154 int8 = -78
        _ = v154
    default:
        fmt.Println(v143, v1)
    }
    for fi155 := 0; fi155 < 17; fi155++ {
        fmt.Println(v147, v3, v151)
        _ = fi155
    }
    fmt.Println("v1:", v1)
    fmt.Println("v142:", v142)
    fmt.Println("v143:", v143)
    fmt.Println("v147:", v147)
    fmt.Println("v151:", v151)
    fmt.Println("v152:", v152)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v142
    _ = v143
    _ = v144
    _ = v145
    _ = v146
    _ = v147
    _ = v151
    _ = v152
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
