package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 string; F1 uint64; F2 int }
type GenS4 struct { F0 int; F1 int16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint32; F1 int16 }
type GenS7 struct { F0 string; F1 int; F2 int }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint32; F1 uint16 }
type GenS10 struct { F0 uint16; F1 uint }
func genFn0(p0 bool) bool {
    v5 := []int{128, 413, 42, 42}
    for fi6 := 0; fi6 < 15; fi6++ {
        if p0 {
            var v7 uint = 2
            p0 = (p0 || p0)
            v7 -= (v7 ^ v7)
            _ = v7
        }
        _ = fi6
    }
    fmt.Println(p0)
    for fi8 := 0; fi8 < 7; fi8++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi8
    }
    _ = p0
    _ = v5
    return p0
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]uint32)
    var v3 bool = (0 == 127)
    var v4 uint16 = 353
    v9 := GenS1{F0: 255, F1: v4}
    fmt.Println(v4, v3)
    switch "" {
    case "a":
        if v3 {
            v10 := []int{232, 0}
            v3 = true
            _ = v10
        } else {
            if true {
                _ = len(v2)
                switch ("foo bar" + "\n\t") {
                case "a":
                    for fi11 := 0; fi11 < 5; fi11++ {
                        v12 := GenS1{F0: uint(v4), F1: v4}
                        if (0.0 >= (float32(1e10) - (-2.25 - (0.1 / ((-2.25 * 1.5) + 3.14159))))) {
                            v1[2] = ((-437 >= 42) && v3)
                            var v13 int8 = -88
                            var v14 uint = 2
                            _ = v13
                            _ = v14
                        } else {
                            v15 := []int{(-(fi11))}
                            if v3 {
                                for fi16 := 0; fi16 < 8; fi16++ {
                                    var v17 int64 = -1
                                    _ = fi16
                                    _ = v17
                                }
                            } else {
                                if true {
                                    if (32767 != (fi11 >> ((4294967295) & 63))) {
                                        v3 = (198 == (v4 ^ v4))
                                    } else {
                                        if (127 != -128) {
                                            if false {
                                                v18 := GenS1{F0: 18446744073709551615, F1: 255}
                                                var v19 bool = true
                                                _ = v18
                                                _ = v19
                                            }
                                        }
                                        if v3 {
                                            for fi20 := 0; fi20 < 10; fi20++ {
                                                v4 += (v4 << ((476) & 63))
                                                fmt.Println(fi20, v3, v4)
                                                v2[""] = 2147483648
                                                _ = fi20
                                            }
                                        }
                                        v3 = true
                                    }
                                    var v21 uint32 = 1
                                    v22 := GenS2{F0: "a", F1: fi11}
                                    _ = v21
                                    _ = v22
                                }
                                for fi23 := 0; fi23 < 19; fi23++ {
                                    var v24 int = fi23
                                    var v25 bool = true
                                    if v25 {
                                        var v26 uint8 = 127
                                        var v27 int8 = 127
                                        _ = v26
                                        _ = v27
                                    }
                                    _ = fi23
                                    _ = v24
                                    _ = v25
                                }
                                var v28 uint16 = (v4 - (v4 | 2))
                                _ = v28
                            }
                            _ = v15
                        }
                        _ = fi11
                        _ = v12
                    }
                case "foo bar":
                    var v29 uint8 = 255
                    _ = v29
                default:
                    var v30 uint8 = 255
                    _ = v30
                }
                var v31 uint = 479
                _ = v31
            } else {
                if (0 > 127) {
                    v32 := GenS2{F0: "hello", F1: -9223372036854775808}
                    for fi33 := 0; fi33 < 1; fi33++ {
                        var v34 uint = 42
                        if (1 != 960) {
                            v34 = (v34 / (((^(v34))) | 1))
                            for fi35 := 0; fi35 < 7; fi35++ {
                                v34 = (v34 + (v34 * (^(v34))))
                                _ = fi35
                            }
                        } else {
                            v36 := GenS3{F0: "a", F1: 4294967295, F2: (fi33 >> ((869) & 63))}
                            _ = v36
                        }
                        var v37 uint = (^(v34))
                        _ = fi33
                        _ = v34
                        _ = v37
                    }
                    _ = v32
                } else {
                    if (42 > 1) {
                        if false {
                            if ("foo bar" > "a") {
                                var v38 int8 = 0
                                v39 := []int{128, 2147483647}
                                _ = v38
                                _ = v39
                            } else {
                                switch 32767 {
                                case -222:
                                    v40 := []int{9223372036854775807}
                                    _ = v40
                                case 32767:
                                    v2[("hello" + ("\n\t" + "hello"))] = uint32(v4)
                                default:
                                    var v41 uint16 = 2
                                    _ = v41
                                }
                                v42 := []int{-1, 128, 1, 406}
                                var v43 int8 = 127
                                _ = v42
                                _ = v43
                            }
                            if (555 > 158) {
                                fmt.Println(v3, v3)
                                var v44 uint32 = 1
                                v45 := v2["\n\t"]
                                _ = v44
                                _ = v45
                            }
                        } else {
                            var v46 uint8 = 42
                            _ = v46
                        }
                    } else {
                        if v3 {
                            if (-818 >= 2147483647) {
                                v3 = v3
                                fmt.Println(v3, v3)
                            }
                            v2[("foo bar" + ("hello" + "hello"))] = 458
                        } else {
                            v3 = v3
                            var v47 float64 = 3.14159
                            v48 := GenS3{F0: "\n\t", F1: 826, F2: 223}
                            _ = v47
                            _ = v48
                        }
                    }
                    fmt.Println(v4)
                }
                var v49 uint = 42
                v50 := v1[0]
                _ = v49
                _ = v50
            }
            if (-1 >= -9223372036854775808) {
                if (true && (v3 || false)) {
                    v3 = (1 > 1)
                }
                fmt.Println(v3)
            } else {
                var v51 float64 = float64(0.0)
                for fi52 := 0; fi52 < 6; fi52++ {
                    var v53 bool = ((3.14159 * 0.1) > -2.25)
                    var v54 bool = (0.0 >= ((3.14159 / float32(((v51 + float64(0.0)) + (((v51 / float64(0.1)) - v51) + 1e10)))) * (((0.1 / -2.25) + (float32(v51) / 0.1)) / 3.14159)))
                    _ = fi52
                    _ = v53
                    _ = v54
                }
                for fi55 := 0; fi55 < 4; fi55++ {
                    var v56 float32 = 1.5
                    _ = len(v2)
                    var v57 int64 = 1
                    _ = fi55
                    _ = v56
                    _ = v57
                }
                _ = v51
            }
            if ((0 > 1) && (32768 == 4294967295)) {
                if v3 {
                    var v58 uint = 18446744073709551615
                    _ = v58
                } else {
                    v3 = (!false)
                }
            }
        }
    case "foo bar":
        v4 = (v4 - v4)
    case "":
        v4 = v4
    }
    for fi59 := 0; fi59 < 17; fi59++ {
        v60 := []int{fi59, (fi59 ^ 32768), fi59, fi59}
        for fi61 := 0; fi61 < 13; fi61++ {
            var v62 uint = 904
            var v63 uint16 = (^(v4))
            _ = fi61
            _ = v62
            _ = v63
        }
        _ = len(v1)
        _ = fi59
        _ = v60
    }
    v3 = true
    fmt.Println(v3, v4)
    v64 := GenS2{F0: "foo bar", F1: 0}
    if (v3 || false) {
        switch "\n\t" {
        case "hello":
            v65 := GenS4{F0: 2147483647, F1: int16(v4)}
            _ = v65
        case "foo bar":
            for fi66 := 0; fi66 < 19; fi66++ {
                if (!(!((!(0 <= 2147483647)) || v3))) {
                    fmt.Println(fi66, v3)
                    var v67 uint8 = 2
                    _ = v67
                }
                for fi68 := 0; fi68 < 7; fi68++ {
                    switch "\n\t" {
                    case "":
                        if (!(!v3)) {
                            v69 := GenS4{F0: fi68, F1: 42}
                            switch ("" + "\n\t") {
                            case "\n\t":
                                v2[("hello" + "hello")] = 4294967295
                            case "a":
                                var v70 int64 = -9223372036854775808
                                _ = v70
                            }
                            var v71 int16 = 32767
                            _ = v69
                            _ = v71
                        } else {
                            v72 := GenS1{F0: 1, F1: 296}
                            if v3 {
                                v4 = v4
                            }
                            _ = v72
                        }
                    default:
                        switch "hello" {
                        case "a":
                            fmt.Println(v4, v3)
                        }
                    }
                    if false {
                        if v3 {
                            v73 := []int{fi68, -2147483648}
                            _ = v73
                        } else {
                            v74 := GenS4{F0: fi68, F1: -32768}
                            v4 -= v4
                            if v3 {
                                if false {
                                    v75 := GenS2{F0: ("foo bar" + "\n\t"), F1: fi68}
                                    if false {
                                        v4 += v4
                                        delete(v1, 2)
                                    } else {
                                        if false {
                                            var v76 int64 = -1
                                            var v77 uint64 = 4294967295
                                            if (((v4 - v4) > v4) || false) {
                                                var v78 uint64 = v77
                                                _ = v78
                                            }
                                            _ = v76
                                            _ = v77
                                        }
                                    }
                                    _ = v75
                                }
                            } else {
                                if v3 {
                                    fmt.Println(fi68, v3, v3)
                                    v4 += v4
                                } else {
                                    var v79 int64 = 924
                                    v80 := []int{fi66, fi66, (fi68 % ((32768) | 1))}
                                    v3 = true
                                    _ = v79
                                    _ = v80
                                }
                            }
                            _ = v74
                        }
                        v3 = v3
                        switch (fi68 + fi66) {
                        case 4294967295:
                            if v3 {
                                v3 = v3
                            }
                        case 0:
                            switch (("a" + (("a" + (("hello" + "foo bar") + ("a" + "a"))) + (("" + (("\n\t" + "") + "foo bar")) + (("\n\t" + ("hello" + "hello")) + (("\n\t" + ("foo bar" + "")) + "foo bar"))))) + "\n\t") {
                            case "":
                                if (197 == 496) {
                                    v3 = v3
                                    v3 = v3
                                    fmt.Println(fi68)
                                } else {
                                    v81 := GenS3{F0: "", F1: 336, F2: fi68}
                                    if (v3 || (128 <= 0)) {
                                        if v3 {
                                            var v82 uint16 = v4
                                            _ = v82
                                        } else {
                                            fmt.Println(v4, v3)
                                            v3 = v3
                                        }
                                        v3 = (!v3)
                                        fmt.Println(fi68, fi68)
                                    } else {
                                        if (4294967295 > 18446744073709551615) {
                                            v4 = (v4 & (v4 / ((v4) | 1)))
                                        } else {
                                            if (1 != 65535) {
                                                var v83 float32 = -2.25
                                                v4 += v4
                                                var v84 int = -2147483648
                                                delete(v1, 2147483648)
                                                v85 := GenS5{F0: -2147483648}
                                                switch (fi66 | 128) {
                                                case -2147483648:
                                                    var v86 float64 = float64((v83 * (0.0 - v83)))
                                                    _ = v86
                                                case 128:
                                                    v84 -= 128
                                                    v87 := GenS6{F0: 795, F1: -439}
                                                    v83 -= (v83 - (0.0 - (v83 + (0.1 * ((((3.14159 - float32(-2.25)) - v83) * (v83 + v83)) - v83)))))
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        v3 = (v3 && (!v3))
                                                    case "a":
                                                        v88 := GenS7{F0: ("\n\t" + ""), F1: (fi66 + fi66), F2: -826}
                                                        _ = v88
                                                    default:
                                                        v89 := GenS3{F0: ("hello" + ("hello" + "foo bar")), F1: 1, F2: fi66}
                                                        _ = v89
                                                    }
                                                    _ = v87
                                                case 42:
                                                    if ("" <= "\n\t") {
                                                        var v90 int64 = 9223372036854775807
                                                        var v91 int64 = 0
                                                        switch ("" + "a") {
                                                        case "\n\t":
                                                            var v92 uint = 0
                                                            _ = v92
                                                        case "hello":
                                                            var v93 int32 = 0
                                                            _ = v93
                                                        case "a":
                                                            switch "\n\t" {
                                                            case "foo bar":
                                                                v3 = true
                                                            case "hello":
                                                                v1[2147483647] = (!true)
                                                            }
                                                        default:
                                                            v94 := []int{fi66}
                                                            _ = v94
                                                        }
                                                        _ = v90
                                                        _ = v91
                                                    } else {
                                                        if v3 {
                                                            var v95 int64 = 9223372036854775807
                                                            _ = v95
                                                        } else {
                                                            var v96 bool = (v3 || v3)
                                                            if v3 {
                                                                if (!(100 >= -63)) {
                                                                    v84 = -76
                                                                    v96 = (v3 || true)
                                                                    v96 = (-388 > 0)
                                                                } else {
                                                                    fmt.Println(fi68, v84, v84)
                                                                    var v97 uint16 = (v4 & v4)
                                                                    _ = v97
                                                                }
                                                                var v98 uint16 = 32767
                                                                _ = v98
                                                            }
                                                            _ = v96
                                                        }
                                                    }
                                                }
                                                v3 = (1e10 != 0.0)
                                                _ = v83
                                                _ = v84
                                                _ = v85
                                            } else {
                                                switch "foo bar" {
                                                case "a":
                                                    v99 := GenS6{F0: 2147483647, F1: 1}
                                                    _ = v99
                                                case "hello":
                                                    if v3 {
                                                        delete(v1, 2)
                                                    } else {
                                                        v100 := GenS4{F0: fi68, F1: 0}
                                                        switch fi66 {
                                                        case -425:
                                                            v3 = (-1 != 0)
                                                        case 42:
                                                            var v101 bool = ((^(v4)) <= 32767)
                                                            _ = v101
                                                        case -808:
                                                            var v102 uint = 575
                                                            _ = v102
                                                        }
                                                        v103 := GenS8{F0: -9223372036854775808}
                                                        _ = v100
                                                        _ = v103
                                                    }
                                                    v104 := []int{fi68}
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        var v105 int32 = 42
                                                        var v106 bool = v3
                                                        _ = v105
                                                        _ = v106
                                                    case "foo bar":
                                                        v4 = (v4 & v4)
                                                    case "":
                                                        var v107 float32 = -2.25
                                                        _ = v107
                                                    default:
                                                        v3 = v3
                                                    }
                                                    fmt.Println(fi68, v3)
                                                    _ = v104
                                                }
                                            }
                                        }
                                    }
                                    var v108 string = "hello"
                                    _ = v108
                                    _ = v81
                                }
                            default:
                                v109 := GenS1{F0: 0, F1: v4}
                                _ = v109
                            }
                            v3 = true
                            if (127 == 127) {
                                v110 := GenS2{F0: "foo bar", F1: (^(fi68))}
                                _ = v110
                            }
                        case 2147483647:
                            v111 := GenS8{F0: 42}
                            _ = v111
                        default:
                            v112 := GenS2{F0: ("" + (("hello" + ("a" + "")) + "a")), F1: fi68}
                            v113 := GenS9{F0: 531, F1: v4}
                            _ = v112
                            _ = v113
                        }
                        var v114 uint32 = 717
                        _ = v114
                    } else {
                        if v3 {
                            if (!v3) {
                                v3 = ((2 >= 0) || v3)
                            }
                            fmt.Println(v4, v3)
                        }
                        v3 = v3
                        v4 += (v4 | (^(v4)))
                    }
                    if (!((1 != -620) || false)) {
                        v3 = v3
                        switch "a" {
                        case "":
                            switch fi66 {
                            case 4294967295:
                                fmt.Println(fi66, v3, v3)
                            case -138:
                                v115 := []int{(fi68 / ((32767) | 1)), 0}
                                _ = v115
                            default:
                                v116 := []int{}
                                _ = v116
                            }
                        case "a":
                            if (-1 >= 100) {
                                v4 = (^(v4))
                                v1[42] = (uint(v4) <= 65535)
                                if false {
                                    v117 := []int{(fi66 - 0)}
                                    _ = v117
                                }
                            }
                        default:
                            fmt.Println(fi68, v3)
                        }
                        if false {
                            fmt.Println(fi66, v4, v3)
                        }
                    } else {
                        fmt.Println(fi66, fi66)
                        fmt.Println(fi68, fi66)
                    }
                    _ = fi68
                }
                _ = fi66
            }
        case "\n\t":
            v118 := []int{}
            _ = v118
        }
        v119 := GenS9{F0: 506, F1: (^(v4))}
        _ = v119
    }
    delete(v1, 2147483648)
    if (v3 && true) {
        v120 := []int{}
        switch "a" {
        case "hello":
            v4 = (v4 & (v4 % ((65535) | 1)))
        case "\n\t":
            v4 += (v4 | (^(v4)))
        case "foo bar":
            var v121 uint = 141
            _ = v121
        default:
            var v122 uint32 = 2
            _ = v122
        }
        v123 := GenS6{F0: 1, F1: -32768}
        _ = v120
        _ = v123
    } else {
        fmt.Println(v3, v4)
        for fi124 := 0; fi124 < 1; fi124++ {
            v125 := v2["\n\t"]
            var v126 uint64 = 142
            switch ("hello" + ("" + "\n\t")) {
            case "hello":
                v3 = (!((float32(-2.25) < 0.1) && (32767 == -32768)))
            }
            _ = fi124
            _ = v125
            _ = v126
        }
    }
    for fi127 := 0; fi127 < 13; fi127++ {
        fmt.Println(v4, v4, v3)
        v3 = v3
        if true {
            var v128 int16 = -105
            for fi129 := 0; fi129 < 18; fi129++ {
                var v130 uint = 4294967295
                _ = fi129
                _ = v130
            }
            v4 = v4
            _ = v128
        } else {
            v3 = v3
            if v3 {
                if (v4 != v4) {
                    v4 = (^(v4))
                    v131 := GenS5{F0: -1}
                    for fi132 := 0; fi132 < 8; fi132++ {
                        v4 -= (v4 >> ((0) & 63))
                        _ = fi132
                    }
                    _ = v131
                }
                if (100 < -128) {
                    var v133 bool = (false || ((^(v4)) == 1))
                    v4 -= (v4 ^ (v4 + (v4 >> ((538) & 63))))
                    _ = v133
                }
            }
            if (32767 != 116) {
                if false {
                    _ = len(v2)
                    v2["\n\t"] = 362
                    fmt.Println(v3)
                } else {
                    v134 := GenS10{F0: 2, F1: 0}
                    _ = v134
                }
            } else {
                var v135 int64 = -894
                _ = v135
            }
        }
        _ = fi127
    }
    if (!(255 > (v4 * v4))) {
        if (!v3) {
            v4 = (^(v4))
            var v136 int64 = 1
            v137 := GenS10{F0: v4, F1: 0}
            _ = v136
            _ = v137
        } else {
            v4 = v4
            var v138 uint32 = 2147483648
            _ = v138
        }
        v1[507] = (false || (v4 >= v4))
        var v139 uint8 = 0
        _ = v139
    } else {
        var v140 uint64 = 0
        if (1e10 == ((1e10 * 3.14159) - 0.1)) {
            v141 := GenS7{F0: "", F1: 32767, F2: 32768}
            _ = v141
        } else {
            v140 = 149
        }
        _ = len(v1)
        _ = v140
    }
    fmt.Println(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v64
    _ = v9
}
