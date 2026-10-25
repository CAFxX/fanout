package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 uint8; F1 int16; F2 uint }
type GenS4 struct { F0 int8; F1 int16; F2 int }
type GenS5 struct { F0 uint32; F1 int; F2 bool }
type GenS6 struct { F0 int; F1 int64 }
type GenS7 struct { F0 uint8; F1 bool }
type GenS8 struct { F0 int64; F1 uint64 }
type GenS9 struct { F0 uint16; F1 uint16 }
type GenS10 struct { F0 int64; F1 uint32 }
type GenS11 struct { F0 uint32; F1 int }
type GenS12 struct { F0 int; F1 uint64 }
type GenS13 struct { F0 bool; F1 uint32; F2 int32 }
type GenS14 struct { F0 uint8; F1 uint64 }
type GenS15 struct { F0 uint; F1 int64 }
type GenS16 struct { F0 bool }
type GenS17 struct { F0 int16 }
type GenS18 struct { F0 int16; F1 int64 }
type GenS19 struct { F0 uint8; F1 int8; F2 uint }
func genFn0(p0 uint8) uint {
    var v4 uint8 = (p0 - 2)
    for fi5 := 0; fi5 < 5; fi5++ {
        if false {
            var v6 string = "hello"
            _ = v6
        } else {
            fmt.Println(fi5, fi5)
            v7 := GenS1{F0: int32(fi5), F1: 2147483647}
            v8 := GenS1{F0: 2147483647, F1: -2147483648}
            _ = v7
            _ = v8
        }
        var v9 int16 = -461
        if true {
            fmt.Println(p0, fi5)
            v4 = (v4 & (p0 % ((p0) | 1)))
        } else {
            for fi10 := 0; fi10 < 1; fi10++ {
                if (v4 > p0) {
                    if (-1 >= 9223372036854775807) {
                        v11 := []int{}
                        _ = v11
                    }
                    v9 = (v9 ^ v9)
                    if ("\n\t" < ("foo bar" + "hello")) {
                        v12 := []int{}
                        var v13 uint8 = (v4 ^ (v4 << ((766) & 63)))
                        v12 = append(v12, (-(fi5)))
                        v14 := GenS1{F0: -704, F1: -470}
                        _ = v12
                        _ = v13
                        _ = v14
                    } else {
                        if false {
                            var v15 float32 = (3.14159 / (3.14159 / (-2.25 + -2.25)))
                            if ((3.14159 > 1e10) && (!(true || false))) {
                                v16 := []int{}
                                if len(v16) >= 2 {
                                    v16 = v16[1:len(v16)]
                                }
                                v15 -= (1.5 - (v15 + v15))
                                _ = v16
                            } else {
                                p0 = p0
                                v9 = v9
                            }
                            var v17 uint64 = 904
                            _ = v15
                            _ = v17
                        }
                        fmt.Println(fi5, v4, fi10)
                        var v18 uint32 = 42
                        _ = v18
                    }
                    v19 := []int{1, int(v4), (fi10 - int(p0)), fi5}
                    _ = v19
                }
                v4 -= p0
                _ = fi10
            }
            var v20 float64 = float64(0.1)
            _ = v20
        }
        _ = fi5
        _ = v9
    }
    v4 = p0
    _ = p0
    _ = v4
    return 0
}

func main() {
    v1 := []uint32{472, 2147483647}
    v2 := make(map[uint32]bool)
    var v3 uint64 = 18446744073709551615
    switch "hello" {
    case "":
        if ((true && (!(127 == 174))) || false) {
            v3 -= (^(v3))
            v21 := v2[2]
            v22 := GenS1{F0: 2147483647, F1: -1}
            _ = v21
            _ = v22
        }
    case "hello":
        for fi23 := 0; fi23 < 19; fi23++ {
            switch ("foo bar" + ("a" + "\n\t")) {
            case "foo bar":
                if (true || (-49 <= -1)) {
                    for fi24 := 0; fi24 < 10; fi24++ {
                        v25 := GenS2{F0: v3, F1: genFn0(12)}
                        _ = fi24
                        _ = v25
                    }
                    v26 := GenS3{F0: 128, F1: -1, F2: 4294967295}
                    _ = v26
                }
            case "":
                v3 = 0
            case "hello":
                switch (-(fi23)) {
                case 0:
                    _ = len(v1)
                case 42:
                    v27 := GenS4{F0: 100, F1: -32768, F2: fi23}
                    _ = v27
                case -2147483648:
                    v28 := GenS5{F0: 75, F1: (^(fi23)), F2: false}
                    _ = v28
                default:
                    switch "" {
                    case "foo bar":
                        fmt.Println(v3)
                    case "hello":
                        if false {
                            fmt.Println(v3, v3)
                        } else {
                            v29 := GenS1{F0: 0, F1: 1}
                            _ = v29
                        }
                    }
                }
            default:
                switch "foo bar" {
                case "foo bar":
                    var v30 float64 = (1.5 - ((((1e10 - 1e10) * 3.14159) / 1.0) / 3.14159))
                    _ = v30
                case "a":
                    for fi31 := 0; fi31 < 19; fi31++ {
                        v32 := GenS3{F0: 47, F1: 723, F2: 164}
                        fmt.Println(v3)
                        v1 = append(v1, 4294967295)
                        v3 += (v3 >> ((4294967295) & 63))
                        _ = fi31
                        _ = v32
                    }
                case "\n\t":
                    var v33 int16 = 1000
                    _ = v33
                default:
                    v34 := GenS6{F0: fi23, F1: 9223372036854775807}
                    _ = v34
                }
            }
            if false {
                var v35 float32 = ((0.1 + (-2.25 + 0.0)) + (3.14159 * 0.1))
                var v36 uint16 = 32767
                var v37 uint64 = (v3 << ((genFn0(127)) & 63))
                _ = v35
                _ = v36
                _ = v37
            } else {
                for fi38 := 0; fi38 < 18; fi38++ {
                    var v39 uint16 = 32767
                    v40 := GenS7{F0: 127, F1: (!true)}
                    v41 := GenS2{F0: v3, F1: 2}
                    v3 = v3
                    v3 -= 4294967295
                    var v42 uint64 = 1
                    if (((0.0 / float64(-2.25)) - float64(float32(float64(1.5)))) < 3.14159) {
                        var v43 string = ("foo bar" + ("hello" + "foo bar"))
                        fmt.Println(v39, v3, v3)
                        fmt.Println(fi38)
                        _ = v43
                    } else {
                        var v44 uint64 = (^(v3))
                        _ = v44
                    }
                    _ = fi38
                    _ = v39
                    _ = v40
                    _ = v41
                    _ = v42
                }
                var v45 float32 = 0.1
                if ((2 == (v3 / (((v3 & v3)) | 1))) || true) {
                    var v46 int32 = 0
                    v45 = v45
                    _ = v46
                }
                _ = v45
            }
            fmt.Println(v3, v3)
            _ = fi23
        }
    default:
        v3 = v3
    }
    if false {
        var v47 float32 = (0.1 / 1.0)
        if ((("" + ("a" + "foo bar")) + "") == (("" + "foo bar") + "")) {
            v48 := GenS4{F0: -52, F1: 42, F2: 2147483647}
            _ = v48
        } else {
            v1 = append(v1, 2147483647)
            switch -1 {
            case -906:
                var v49 int = -2147483648
                _ = v49
            case -423:
                if (false && true) {
                    fmt.Println(v47)
                    if ((v3 ^ (v3 | v3)) != 1) {
                        if len(v1) > 0 {
                            v1[0] = 881
                        }
                        if true {
                            v50 := GenS2{F0: 2, F1: 65535}
                            _ = v50
                        }
                        v51 := v2[499]
                        _ = v51
                    }
                    v52 := GenS6{F0: 32767, F1: 9223372036854775807}
                    _ = v52
                } else {
                    v3 = 0
                }
            default:
                var v53 int32 = 0
                _ = v53
            }
            switch "\n\t" {
            case "hello":
                for fi54 := 0; fi54 < 12; fi54++ {
                    var v55 int8 = -122
                    if ((((v47 * (((1.5 / v47) * v47) - 1.5)) / 1.0) / 1.5) <= 3.14159) {
                        var v56 uint64 = (v3 << ((1) & 63))
                        _ = v56
                    }
                    _ = fi54
                    _ = v55
                }
            }
        }
        v57 := GenS8{F0: -1, F1: (v3 - v3)}
        _ = v47
        _ = v57
    }
    for fi58 := 0; fi58 < 12; fi58++ {
        for fi59 := 0; fi59 < 8; fi59++ {
            if (18 > 0) {
                fmt.Println(fi59, fi58, fi59)
                if (!false) {
                    v3 = (v3 << ((genFn0(222)) & 63))
                    if true {
                        v60 := GenS6{F0: -2147483648, F1: -1}
                        v61 := GenS9{F0: 255, F1: 0}
                        fmt.Println(fi58, fi58)
                        var v62 float32 = -2.25
                        fmt.Println(fi58, v62, v3)
                        _ = v60
                        _ = v61
                        _ = v62
                    }
                    v63 := GenS9{F0: 2, F1: 528}
                    _ = v63
                } else {
                    v2[2147483647] = true
                    fmt.Println(fi59)
                    var v64 string = "a"
                    v65 := GenS10{F0: 0, F1: 2147483648}
                    _ = v64
                    _ = v65
                }
                fmt.Println(v3, v3, fi58)
                v66 := v2[42]
                _ = v66
            }
            var v67 bool = (-542 >= 42)
            _ = fi59
            _ = v67
        }
        v3 = 4294967295
        for fi68 := 0; fi68 < 5; fi68++ {
            v69 := GenS11{F0: 4294967295, F1: (fi58 >> ((uint(v3)) & 63))}
            v3 += v3
            _ = fi68
            _ = v69
        }
        _ = fi58
    }
    if ((518 == uint16(v3)) || (-38 < -128)) {
        fmt.Println(v3)
        if len(v1) > 0 {
            v1[0] = 2
        }
        if ((-113 <= 42) && true) {
            for fi70 := 0; fi70 < 9; fi70++ {
                var v71 int = fi70
                v71 = fi70
                fmt.Println(v3, v3)
                _ = fi70
                _ = v71
            }
            _ = len(v1)
            v3 -= 1
        }
    } else {
        var v72 float64 = float64(1e10)
        for fi73 := 0; fi73 < 15; fi73++ {
            var v74 int64 = 995
            if len(v1) > 0 {
                v1[0] = 42
            }
            if (3.14159 != (0.1 / 1.0)) {
                for fi75 := 0; fi75 < 15; fi75++ {
                    fmt.Println(v3, fi73, fi75)
                    var v76 uint16 = 42
                    _ = fi75
                    _ = v76
                }
                v1 = append(v1, 0)
            } else {
                v72 -= (3.14159 / v72)
                for fi77 := 0; fi77 < 1; fi77++ {
                    v74 += 0
                    if true {
                        if (128 > 127) {
                            var v78 bool = false
                            fmt.Println(v3)
                            v72 += (v72 * v72)
                            v72 = 1.5
                            var v79 int16 = 1
                            v74 = (v74 ^ v74)
                            _ = v78
                            _ = v79
                        } else {
                            fmt.Println(v72)
                            var v80 uint32 = 65535
                            v80 -= (^(v80))
                            _ = v80
                        }
                    }
                    _ = fi77
                }
                v72 += ((((0.0 - (v72 * float64((1.5 / (0.1 / 3.14159))))) - (0.1 + v72)) + ((float64(1.5) / 1.5) * (v72 - v72))) / float64(1e10))
            }
            _ = fi73
            _ = v74
        }
        v81 := GenS12{F0: 32767, F1: v3}
        _ = v72
        _ = v81
    }
    if (false || true) {
        v3 = v3
    } else {
        switch -576 {
        case 32767:
            for fi82 := 0; fi82 < 2; fi82++ {
                var v83 uint8 = 154
                _ = fi82
                _ = v83
            }
        }
        v3 -= (v3 | (v3 | v3))
        v84 := v2[4294967295]
        _ = v84
    }
    _ = len(v1)
    if true {
        if (-1 < 122) {
            v3 = v3
        } else {
            v85 := GenS13{F0: (435 <= -1), F1: 2, F2: -1}
            var v86 int32 = 138
            if ((v3 ^ v3) == v3) {
                if true {
                    for fi87 := 0; fi87 < 3; fi87++ {
                        _ = len(v1)
                        _ = fi87
                    }
                } else {
                    for fi88 := 0; fi88 < 15; fi88++ {
                        if false {
                            switch "hello" {
                            case "hello":
                                switch (fi88 << ((255) & 63)) {
                                case -615:
                                    fmt.Println(fi88, fi88)
                                case 409:
                                    v86 = (v86 % ((864) | 1))
                                case -9223372036854775808:
                                    switch (fi88 * 4294967295) {
                                    case 128:
                                        for fi89 := 0; fi89 < 1; fi89++ {
                                            switch (fi89 + 642) {
                                            case 2147483647:
                                                v3 -= v3
                                            }
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            v1 = append(v1, 65535)
                                            if false {
                                                v86 -= (v86 ^ 1)
                                                var v90 uint = 255
                                                var v91 float64 = 0.0
                                                if (740 < 0) {
                                                    v86 += (v86 + (^(v86)))
                                                    switch fi88 {
                                                    case 0:
                                                        if true {
                                                            var v92 float32 = 1.5
                                                            _ = v92
                                                        } else {
                                                            if true {
                                                                if true {
                                                                    v93 := GenS14{F0: 255, F1: v3}
                                                                    _ = v93
                                                                }
                                                                if (0.0 != (3.14159 - 1.5)) {
                                                                    if (!false) {
                                                                        var v94 uint64 = v3
                                                                        var v95 uint16 = 65535
                                                                        v1 = append(v1, 0)
                                                                        fmt.Println(fi88)
                                                                        if (!false) {
                                                                            if len(v1) > 0 {
                                                                                v1[0] = 0
                                                                            }
                                                                        } else {
                                                                            fmt.Println(v91, v86, v95)
                                                                        }
                                                                        _ = v94
                                                                        _ = v95
                                                                    } else {
                                                                        v91 = float64(((0.0 / 1.0) - ((0.1 + (float32((1.5 - (((v91 - v91) - -2.25) / v91))) / (0.1 * 1.5))) * float32((0.0 * v91)))))
                                                                        _ = len(v1)
                                                                        v90 = (v90 | (^(v90)))
                                                                    }
                                                                    v2[2147483647] = (42 > 42)
                                                                } else {
                                                                    if (-168 > v86) {
                                                                        v3 = (v3 >> (((^(v90))) & 63))
                                                                        v96 := v2[uint32(v90)]
                                                                        _ = v96
                                                                    }
                                                                }
                                                                var v97 float32 = 1.5
                                                                _ = v97
                                                            } else {
                                                                v98 := GenS15{F0: (v90 | genFn0(255)), F1: 504}
                                                                if true {
                                                                    var v99 float32 = float32((1.5 / (v91 / 1e10)))
                                                                    v100 := GenS11{F0: 65535, F1: (fi88 - fi88)}
                                                                    var v101 int = (fi88 >> (((v90 ^ (v90 / ((834) | 1)))) & 63))
                                                                    _ = v100
                                                                    _ = v101
                                                                    _ = v99
                                                                } else {
                                                                    v102 := GenS8{F0: -1, F1: (^(v3))}
                                                                    _ = v102
                                                                }
                                                                _ = len(v1)
                                                                _ = v98
                                                            }
                                                        }
                                                    case 978:
                                                        var v103 int64 = -9223372036854775808
                                                        _ = v103
                                                    default:
                                                        v104 := v2[2147483647]
                                                        var v105 uint32 = 42
                                                        switch "a" {
                                                        case "":
                                                            var v106 int8 = 42
                                                            _ = v106
                                                        case "hello":
                                                            var v107 float64 = ((((((v91 - ((3.14159 - (v91 / float64(0.1))) + v91)) - float64(((0.1 + 1.5) + 0.0))) * (v91 / v91)) * (v91 - ((-2.25 / float64(float32(-2.25))) + 0.0))) / (v91 - v91)) + (0.0 - (1.5 * (float64(1e10) / v91))))
                                                            _ = v107
                                                        case "a":
                                                            _ = len(v2)
                                                        }
                                                        _ = v104
                                                        _ = v105
                                                    }
                                                } else {
                                                    var v108 bool = false
                                                    if len(v1) > 0 {
                                                        v1[0] = 864
                                                    }
                                                    _ = v108
                                                }
                                                var v109 int16 = 388
                                                switch "hello" {
                                                case "a":
                                                    var v110 int32 = (v86 % ((v86) | 1))
                                                    _ = v110
                                                case "\n\t":
                                                    var v111 int64 = -1
                                                    _ = v111
                                                }
                                                _ = v109
                                                _ = v90
                                                _ = v91
                                            }
                                            var v112 bool = (uint32(v3) >= 0)
                                            _ = fi89
                                            _ = v112
                                        }
                                    case 32768:
                                        var v113 int16 = 720
                                        _ = v113
                                    default:
                                        if true {
                                            v114 := GenS1{F0: (v86 * (v86 * 2147483647)), F1: v86}
                                            _ = v114
                                        }
                                    }
                                }
                            case "\n\t":
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            case "a":
                                fmt.Println(fi88, v3)
                            }
                            v115 := v2[224]
                            v3 = (^(v3))
                            _ = v115
                        } else {
                            var v116 int8 = -1
                            _ = v116
                        }
                        v117 := v2[983]
                        _ = fi88
                        _ = v117
                    }
                }
                v3 -= (v3 * (v3 % ((678) | 1)))
                var v118 uint = genFn0(128)
                _ = v118
            } else {
                v86 -= v86
                fmt.Println(v3, v86)
                v119 := GenS16{F0: (false || ((("a" + "foo bar") + ("\n\t" + "a")) == ""))}
                _ = v119
            }
            _ = v85
            _ = v86
        }
        switch 32767 {
        case 42:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case 32768:
            for fi120 := 0; fi120 < 4; fi120++ {
                v3 += v3
                if (!(v3 > v3)) {
                    if (!false) {
                        for fi121 := 0; fi121 < 1; fi121++ {
                            if (2147483647 != 1) {
                                v3 = v3
                                if (!(613 <= 1)) {
                                    var v122 uint32 = 2
                                    v122 += 42
                                    _ = v122
                                }
                                delete(v2, 2147483647)
                            }
                            if true {
                                switch -279 {
                                case 42:
                                    fmt.Println(v3, fi120, v3)
                                case 0:
                                    if (42 == 65535) {
                                        switch 1 {
                                        case -9223372036854775808:
                                            v123 := GenS3{F0: 2, F1: 32767, F2: 65535}
                                            _ = v123
                                        case 32768:
                                            var v124 int8 = 0
                                            _ = v124
                                        case 127:
                                            v3 -= v3
                                        default:
                                            switch (-(fi121)) {
                                            case 0:
                                                v3 += (v3 | 834)
                                            case 4294967295:
                                                var v125 float32 = ((((1.5 + 1e10) + (float32(1.5) * float32(-2.25))) * (1.5 + (0.1 / 1.0))) - 1e10)
                                                _ = v125
                                            default:
                                                var v126 int32 = 42
                                                _ = v126
                                            }
                                        }
                                    }
                                case -9223372036854775808:
                                    v127 := GenS17{F0: 32767}
                                    var v128 string = "\n\t"
                                    switch v128 {
                                    case "foo bar":
                                        v128 = v128
                                    case "hello":
                                        var v129 bool = (((v128 + v128) < ((v128 + v128) + (v128 + ((v128 + (("hello" + v128) + "foo bar")) + v128)))) || false)
                                        _ = v129
                                    case "\n\t":
                                        var v130 uint16 = 42
                                        _ = v130
                                    }
                                    _ = v127
                                    _ = v128
                                }
                            }
                            v1 = append(v1, 42)
                            _ = fi121
                        }
                        if (42 < -1) {
                            v3 -= (v3 ^ v3)
                            v3 = (v3 / (((v3 % ((v3) | 1))) | 1))
                        } else {
                            var v131 uint16 = 32767
                            var v132 string = "foo bar"
                            _ = v131
                            _ = v132
                        }
                        v133 := GenS18{F0: 194, F1: 9223372036854775807}
                        _ = v133
                    }
                }
                var v134 float64 = ((0.1 / 1.0) + 3.14159)
                _ = fi120
                _ = v134
            }
        }
    } else {
        v3 += (v3 * v3)
    }
    if (false || (false || ((-2.25 * 3.14159) < -2.25))) {
        for fi135 := 0; fi135 < 4; fi135++ {
            v3 = 308
            switch (fi135 | fi135) {
            case -864:
                for fi136 := 0; fi136 < 20; fi136++ {
                    var v137 uint32 = 2147483647
                    var v138 uint8 = 130
                    var v139 int16 = 281
                    v140 := GenS12{F0: (fi135 >> ((18446744073709551615) & 63)), F1: v3}
                    _ = fi136
                    _ = v137
                    _ = v138
                    _ = v139
                    _ = v140
                }
            case 127:
                v3 += v3
            default:
                if (((float64(1e10) - float64(1e10)) * (((1.5 * -2.25) + 1.5) + 1.5)) > (0.1 + 0.0)) {
                    v141 := GenS19{F0: 128, F1: 0, F2: genFn0(0)}
                    var v142 string = "\n\t"
                    _ = v141
                    _ = v142
                }
            }
            _ = fi135
        }
        var v143 int64 = -911
        for fi144 := 0; fi144 < 12; fi144++ {
            v145 := v2[628]
            _ = fi144
            _ = v145
        }
        _ = v143
    }
    var v146 uint = 255
    v147 := GenS9{F0: 65535, F1: 65535}
    if true {
        v3 = (v3 >> ((v146) & 63))
        var v148 int32 = -2147483648
        if (!(true && true)) {
            v149 := GenS19{F0: 1, F1: 100, F2: 4294967295}
            if len(v1) > 0 {
                v1[0] = 286
            }
            _ = v149
        }
        _ = v148
    } else {
        for fi150 := 0; fi150 < 6; fi150++ {
            for fi151 := 0; fi151 < 1; fi151++ {
                var v152 int64 = -9223372036854775808
                v153 := GenS1{F0: -765, F1: 785}
                _ = fi151
                _ = v152
                _ = v153
            }
            var v154 uint = v146
            v155 := GenS10{F0: int64(v146), F1: 526}
            _ = fi150
            _ = v154
            _ = v155
        }
        var v156 bool = false
        v3 -= (v3 ^ 18446744073709551615)
        _ = v156
    }
    v3 = (v3 - (v3 >> ((v146) & 63)))
    v157 := GenS18{F0: -50, F1: 0}
    v158 := GenS3{F0: 1, F1: 32767, F2: v146}
    for fi159 := 0; fi159 < 16; fi159++ {
        switch fi159 {
        case 2147483647:
            for fi160 := 0; fi160 < 5; fi160++ {
                v161 := GenS15{F0: v146, F1: 1}
                v146 = 18446744073709551615
                _ = fi160
                _ = v161
            }
        case -9223372036854775808:
            v162 := GenS19{F0: 127, F1: 100, F2: v146}
            _ = v162
        case -1:
            v3 = v3
        default:
            fmt.Println(v3, v146, fi159)
        }
        if len(v1) > 0 {
            v1[0] = 0
        }
        if (!true) {
            if false {
                var v163 uint = (v146 >> ((580) & 63))
                fmt.Println(v163, v163, v3)
                _ = v163
            } else {
                if (!true) {
                    v146 -= 0
                    v164 := GenS19{F0: 2, F1: -76, F2: v146}
                    var v165 uint64 = (v3 << (((v146 - genFn0(247))) & 63))
                    _ = v164
                    _ = v165
                }
            }
            v146 -= genFn0(1)
            v166 := GenS16{F0: (false && (1 > 0))}
            _ = v166
        }
        _ = fi159
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v146:", v146)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v146
    _ = v147
    _ = v157
    _ = v158
    _ = v2
    _ = v3
}
