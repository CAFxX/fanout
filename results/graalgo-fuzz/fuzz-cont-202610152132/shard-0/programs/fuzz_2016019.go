package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16; F1 int32; F2 int8 }
type GenS4 struct { F0 int32; F1 int64 }
type GenS5 struct { F0 int64; F1 int64; F2 int }
type GenS6 struct { F0 int8; F1 int16; F2 bool }
type GenS7 struct { F0 int; F1 int16; F2 int8 }
type GenS8 struct { F0 bool; F1 int16; F2 int }
type GenS9 struct { F0 uint; F1 string; F2 uint16 }
func genFn0(p0 string) uint16 {
    v7 := []int{-525, 0, 128, -1}
    for fi8 := 0; fi8 < 18; fi8++ {
        if false {
            var v9 string = ("a" + (p0 + (((p0 + "a") + ("hello" + "\n\t")) + p0)))
            if (-1 >= -56) {
                var v10 float32 = 1.5
                v11 := GenS1{F0: 9223372036854775807}
                _ = v10
                _ = v11
            }
            for fi12 := 0; fi12 < 15; fi12++ {
                v13 := GenS2{F0: 100}
                if true {
                    var v14 int32 = -199
                    v15 := GenS1{F0: -9223372036854775808}
                    _ = v14
                    _ = v15
                } else {
                    var v16 int64 = 42
                    _ = v16
                }
                if true {
                    if (false && (true && (!(463 <= 352)))) {
                        if false {
                            var v17 uint = 255
                            var v18 bool = true
                            p0 = ("" + p0)
                            _ = v17
                            _ = v18
                        }
                        var v19 float32 = (0.0 * (1e10 + 3.14159))
                        if (32767 == 65535) {
                            var v20 bool = false
                            _ = v20
                        } else {
                            fmt.Println(fi12, v19, fi8)
                            v21 := GenS1{F0: -908}
                            switch v9 {
                            case "":
                                v19 += v19
                            case "\n\t":
                                p0 = (v9 + v9)
                            case "foo bar":
                                var v22 float64 = (1.5 * 3.14159)
                                _ = v22
                            }
                            v19 -= (-2.25 - v19)
                            _ = v21
                        }
                        switch p0 {
                        case "foo bar":
                            v23 := GenS1{F0: -9223372036854775808}
                            _ = v23
                        case "hello":
                            var v24 int = fi12
                            _ = v24
                        case "":
                            _ = len(v7)
                        }
                        _ = v19
                    } else {
                        v25 := GenS1{F0: -632}
                        _ = v25
                    }
                } else {
                    p0 = p0
                    var v26 uint = 255
                    _ = v26
                }
                var v27 bool = false
                v7 = append(v7, (fi12 & fi8))
                switch "" {
                case "":
                    switch fi8 {
                    case 1:
                        if (0 == 18446744073709551615) {
                            var v28 float32 = (3.14159 + 0.0)
                            v28 += (v28 / float32((-2.25 * (float64((1.5 + (v28 - float32(1e10)))) / (float64((v28 / v28)) / 3.14159)))))
                            _ = v28
                        } else {
                            var v29 float32 = (3.14159 + 1.5)
                            if (218 < 128) {
                                v30 := GenS3{F0: 32767, F1: -739, F2: 127}
                                switch 2147483647 {
                                case -473:
                                    switch "\n\t" {
                                    case "hello":
                                        var v31 uint64 = 226
                                        _ = v31
                                    case "a":
                                        switch (fi12 % ((fi12) | 1)) {
                                        case 1:
                                            switch (fi12 ^ fi8) {
                                            case 842:
                                                if true {
                                                    var v32 bool = (!false)
                                                    if v32 {
                                                        fmt.Println(v32, v27, v9)
                                                        var v33 string = v9
                                                        if (4294967295 >= 0) {
                                                            v29 -= (v29 - (0.1 * (1.5 / v29)))
                                                            v9 = "foo bar"
                                                        }
                                                        _ = v33
                                                    }
                                                    _ = v32
                                                }
                                            case -9223372036854775808:
                                                v34 := GenS2{F0: 100}
                                                _ = v34
                                            case -956:
                                                if len(v7) >= 2 {
                                                    v7 = v7[1:len(v7)]
                                                }
                                            default:
                                                var v35 uint16 = 1
                                                v36 := GenS4{F0: -716, F1: 104}
                                                _ = v35
                                                _ = v36
                                            }
                                        case 127:
                                            v37 := GenS5{F0: -9223372036854775808, F1: -869, F2: fi12}
                                            _ = v37
                                        default:
                                            if (100 != 42) {
                                                fmt.Println(v9, v29)
                                                switch "" {
                                                case "\n\t":
                                                    p0 = v9
                                                default:
                                                    var v38 float32 = v29
                                                    _ = v38
                                                }
                                            }
                                        }
                                    case "":
                                        var v39 uint = 18446744073709551615
                                        fmt.Println(v29, fi12, v9)
                                        switch fi8 {
                                        case 1:
                                            var v40 uint32 = 281
                                            _ = v40
                                        case -396:
                                            v39 += (v39 | v39)
                                        }
                                        fmt.Println(fi8, fi8)
                                        var v41 string = p0
                                        _ = v39
                                        _ = v41
                                    }
                                case 127:
                                    var v42 int32 = 42
                                    _ = v42
                                }
                                v29 = v29
                                _ = v30
                            } else {
                                var v43 bool = (35 == 127)
                                _ = v43
                            }
                            v44 := GenS5{F0: 42, F1: 703, F2: fi8}
                            _ = v29
                            _ = v44
                        }
                    case -656:
                        p0 = (p0 + "")
                    default:
                        v7 = append(v7, fi8)
                    }
                case "a":
                    v45 := GenS2{F0: 42}
                    _ = v45
                default:
                    if (true || v27) {
                        if v27 {
                            _ = len(v7)
                            v46 := GenS3{F0: 124, F1: 42, F2: 100}
                            _ = v46
                        }
                        switch "" {
                        case "foo bar":
                            var v47 int64 = -9223372036854775808
                            _ = v47
                        case "\n\t":
                            if len(v7) > 0 {
                                v7[0] = fi8
                            }
                        default:
                            p0 = p0
                        }
                        fmt.Println(v27, p0, fi8)
                    }
                }
                _ = fi12
                _ = v13
                _ = v27
            }
            _ = v9
        }
        switch -927 {
        case 9223372036854775807:
            var v48 float64 = (((0.1 - 0.0) + 3.14159) + -2.25)
            _ = v48
        case 127:
            switch p0 {
            case "":
                fmt.Println(p0, fi8)
            case "\n\t":
                v49 := GenS2{F0: -128}
                _ = v49
            default:
                fmt.Println(fi8)
            }
        case -9223372036854775808:
            if true {
                p0 = (("\n\t" + p0) + (p0 + "hello"))
                var v50 float64 = (0.0 / 0.1)
                p0 = ("\n\t" + (p0 + "hello"))
                _ = v50
            } else {
                p0 = "foo bar"
                v51 := GenS2{F0: 127}
                _ = v51
            }
        }
        _ = fi8
    }
    var v52 int32 = 0
    p0 = (p0 + p0)
    _ = p0
    _ = v52
    _ = v7
    return 65535
}
func genFn1(p0 uint8, p1 uint8) uint {
    if (true || false) {
        for fi53 := 0; fi53 < 14; fi53++ {
            p0 = p1
            _ = fi53
        }
    } else {
        for fi54 := 0; fi54 < 4; fi54++ {
            var v55 uint = 65535
            for fi56 := 0; fi56 < 6; fi56++ {
                var v57 int64 = 0
                if false {
                    switch -113 {
                    case -346:
                        if (1 != 913) {
                            if false {
                                switch ("" + ("a" + "")) {
                                case "foo bar":
                                    p1 -= 42
                                case "hello":
                                    v58 := []int{}
                                    _ = v58
                                default:
                                    fmt.Println(fi56)
                                }
                                p1 += (^(p1))
                                var v59 float64 = 0.1
                                _ = v59
                            }
                            var v60 uint32 = 65535
                            _ = v60
                        } else {
                            v61 := []int{}
                            var v62 float32 = (1.5 + 0.1)
                            v55 -= v55
                            _ = v61
                            _ = v62
                        }
                    case -2147483648:
                        v55 = 211
                    case 128:
                        v63 := GenS2{F0: -73}
                        _ = v63
                    }
                } else {
                    fmt.Println(v57, fi56, v55)
                    fmt.Println(v57)
                    p1 += 42
                }
                _ = fi56
                _ = v57
            }
            _ = fi54
            _ = v55
        }
    }
    var v64 int64 = 42
    v65 := []int{-600, -1, 153}
    p1 = p0
    _ = p0
    _ = p1
    _ = v64
    _ = v65
    return 4294967295
}

func main() {
    var v1 uint16 = 42
    var v2 int16 = 0
    v3 := make(map[string]string)
    var v4 bool = false
    v5 := []string{"hello", "foo bar", "foo bar"}
    v6 := []int{128}
    if (!v4) {
        v4 = (2147483647 < 2147483647)
    } else {
        for fi66 := 0; fi66 < 9; fi66++ {
            fmt.Println(fi66)
            var v67 string = ((("a" + "foo bar") + "hello") + ("" + "a"))
            _ = fi66
            _ = v67
        }
        switch "\n\t" {
        case "a":
            v4 = (v4 || false)
        }
        v2 = v2
    }
    for fi68 := 0; fi68 < 4; fi68++ {
        v69 := GenS4{F0: 0, F1: 9223372036854775807}
        _ = fi68
        _ = v69
    }
    delete(v3, "")
    for fi70 := 0; fi70 < 10; fi70++ {
        var v71 uint64 = 42
        if true {
            v4 = v4
            var v72 int = fi70
            _ = v72
        } else {
            _ = len(v6)
        }
        _ = fi70
        _ = v71
    }
    if v4 {
        _ = len(v5)
    } else {
        switch 950 {
        case -9223372036854775808:
            if true {
                v1 = 65535
                v2 = -32768
            } else {
                fmt.Println(v2)
            }
        case -2147483648:
            v4 = true
        default:
            v1 += (^(v1))
        }
    }
    for fi73 := 0; fi73 < 12; fi73++ {
        fmt.Println(fi73, v4)
        _ = fi73
    }
    if (200 <= 479) {
        _ = len(v6)
    } else {
        if (1 > 42) {
            for fi74 := 0; fi74 < 19; fi74++ {
                if v4 {
                    for fi75 := 0; fi75 < 15; fi75++ {
                        var v76 int32 = -613
                        v4 = false
                        v3["hello"] = ("foo bar" + ((("" + "a") + "foo bar") + "a"))
                        _ = fi75
                        _ = v76
                    }
                    if v4 {
                        if v4 {
                            for fi77 := 0; fi77 < 16; fi77++ {
                                var v78 uint32 = 2147483648
                                _ = fi77
                                _ = v78
                            }
                        }
                        var v79 bool = (289 != 549)
                        _ = v79
                    }
                } else {
                    for fi80 := 0; fi80 < 17; fi80++ {
                        v2 = (v2 & (v2 << ((149) & 63)))
                        var v81 uint8 = 128
                        v1 = (v1 | v1)
                        var v82 int = (fi80 - (fi74 + (fi80 - int(v1))))
                        _ = fi80
                        _ = v81
                        _ = v82
                    }
                    for fi83 := 0; fi83 < 20; fi83++ {
                        var v84 uint64 = 863
                        var v85 uint8 = 231
                        if v4 {
                            v85 += 127
                            v4 = v4
                            v86 := GenS4{F0: 2147483647, F1: 191}
                            v4 = true
                            _ = v86
                        }
                        _ = fi83
                        _ = v84
                        _ = v85
                    }
                    v1 = (v1 / ((255) | 1))
                }
                for fi87 := 0; fi87 < 17; fi87++ {
                    if ((1.5 + 1.5) > (0.1 + 0.0)) {
                        if (-1 == 42) {
                            if false {
                                v88 := GenS3{F0: v1, F1: 2147483647, F2: int8(fi87)}
                                switch fi74 {
                                case 32768:
                                    v89 := GenS4{F0: -2147483648, F1: 0}
                                    _ = v89
                                case -2147483648:
                                    var v90 float32 = float32(float64(float32((1e10 * (float64(1e10) / 1.0)))))
                                    _ = v90
                                }
                                fmt.Println(v2, v1)
                                _ = v88
                            } else {
                                var v91 uint32 = 2147483647
                                if v4 {
                                    var v92 uint8 = 127
                                    var v93 bool = v4
                                    _ = v92
                                    _ = v93
                                } else {
                                    v1 = (v1 * genFn0("foo bar"))
                                }
                                var v94 int64 = 42
                                _ = v91
                                _ = v94
                            }
                            v4 = (v2 == v2)
                            var v95 int64 = 9223372036854775807
                            _ = v95
                        } else {
                            var v96 float64 = (3.14159 + 3.14159)
                            v97 := GenS1{F0: 469}
                            _ = v96
                            _ = v97
                        }
                    }
                    var v98 string = ("foo bar" + "")
                    _ = fi87
                    _ = v98
                }
                _ = fi74
            }
            for fi99 := 0; fi99 < 19; fi99++ {
                if (171 >= 2147483648) {
                    switch (((((("" + "foo bar") + "foo bar") + "a") + (("" + "\n\t") + ("a" + "hello"))) + "") + "a") {
                    case "foo bar":
                        var v100 uint16 = (v1 | 42)
                        _ = v100
                    default:
                        var v101 uint16 = (v1 | genFn0(""))
                        _ = v101
                    }
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                }
                _ = len(v6)
                for fi102 := 0; fi102 < 8; fi102++ {
                    v2 = v2
                    v4 = ((!(((true || (0 != int32(fi99))) && false) && ((!(9223372036854775807 >= -1)) || true))) && false)
                    _ = fi102
                }
                _ = fi99
            }
            v103 := GenS2{F0: 127}
            _ = v103
        }
    }
    for fi104 := 0; fi104 < 8; fi104++ {
        if (!v4) {
            v105 := GenS6{F0: 100, F1: (v2 % ((v2) | 1)), F2: ("foo bar" >= ("\n\t" + ""))}
            v106 := v3[("" + ((((("foo bar" + "a") + "a") + "\n\t") + (("a" + (("hello" + ("foo bar" + "\n\t")) + "hello")) + "\n\t")) + ""))]
            _ = v105
            _ = v106
        } else {
            switch "a" {
            case "foo bar":
                var v107 int8 = 38
                _ = v107
            default:
                for fi108 := 0; fi108 < 20; fi108++ {
                    _ = len(v3)
                    v3[(((("a" + ("hello" + ("" + "foo bar"))) + (("\n\t" + "") + ("\n\t" + ""))) + "\n\t") + "foo bar")] = "foo bar"
                    _ = fi108
                }
            }
        }
        if len(v5) > 0 {
            v5[0] = "hello"
        }
        if v4 {
            for fi109 := 0; fi109 < 1; fi109++ {
                _ = len(v3)
                _ = fi109
            }
        } else {
            v2 = v2
            var v110 uint32 = 4294967295
            _ = v110
        }
        _ = fi104
    }
    for fi111 := 0; fi111 < 15; fi111++ {
        var v112 uint32 = uint32(v1)
        switch (fi111 | (^(fi111))) {
        case -284:
            for fi113 := 0; fi113 < 14; fi113++ {
                switch fi113 {
                case 42:
                    v2 = -294
                    var v114 uint64 = 4294967295
                    v115 := v3["foo bar"]
                    _ = v114
                    _ = v115
                }
                _ = fi113
            }
        }
        var v116 bool = true
        _ = fi111
        _ = v112
        _ = v116
    }
    fmt.Println(v1)
    switch 32768 {
    case -1:
        switch "\n\t" {
        case "foo bar":
            v117 := GenS1{F0: 9223372036854775807}
            _ = v117
        case "hello":
            if v4 {
                v118 := GenS7{F0: 0, F1: 326, F2: -128}
                switch 779 {
                case 845:
                    fmt.Println(v4, v1, v2)
                }
                _ = v118
            } else {
                fmt.Println(v1)
            }
        default:
            var v119 uint = genFn1(42, 42)
            _ = v119
        }
    case 32768:
        for fi120 := 0; fi120 < 1; fi120++ {
            var v121 bool = false
            v122 := GenS8{F0: ((("foo bar" + "") + ("hello" + "foo bar")) < ((((("foo bar" + "a") + ("\n\t" + (("foo bar" + ("\n\t" + ("hello" + ("foo bar" + "a")))) + ""))) + "a") + "a") + ("hello" + (((("\n\t" + ("a" + "a")) + "a") + "") + ("hello" + "hello"))))), F1: v2, F2: 9223372036854775807}
            if (v4 && (1e10 == 1.5)) {
                if v121 {
                    v3[""] = "a"
                }
                v123 := GenS8{F0: false, F1: (v2 >> ((18446744073709551615) & 63)), F2: fi120}
                delete(v3, "foo bar")
                _ = v123
            } else {
                for fi124 := 0; fi124 < 3; fi124++ {
                    fmt.Println(fi124, v1)
                    v2 = (v2 - (v2 >> ((241) & 63)))
                    v1 = v1
                    v125 := v3[("a" + "hello")]
                    _ = fi124
                    _ = v125
                }
            }
            _ = fi120
            _ = v121
            _ = v122
        }
    case 9223372036854775807:
        v1 += v1
    default:
        v126 := GenS3{F0: 1, F1: 2147483647, F2: -89}
        _ = v126
    }
    switch "\n\t" {
    case "foo bar":
        for fi127 := 0; fi127 < 1; fi127++ {
            for fi128 := 0; fi128 < 1; fi128++ {
                var v129 bool = (false && v4)
                v1 = (v1 ^ (^(v1)))
                var v130 int = (fi128 % (((fi128 + -662)) | 1))
                var v131 int16 = v2
                var v132 uint8 = 252
                _ = fi128
                _ = v129
                _ = v130
                _ = v131
                _ = v132
            }
            if len(v6) > 0 {
                v6[0] = fi127
            }
            _ = fi127
        }
    case "a":
        v4 = (v4 && v4)
    default:
        var v133 int16 = v2
        _ = v133
    }
    fmt.Println(v4, v4, v2)
    var v134 int8 = -128
    for fi135 := 0; fi135 < 15; fi135++ {
        switch fi135 {
        case 32767:
            for fi136 := 0; fi136 < 11; fi136++ {
                v137 := GenS9{F0: 1, F1: "a", F2: (^(v1))}
                _ = fi136
                _ = v137
            }
        case 32768:
            v138 := v3["\n\t"]
            _ = v138
        case -161:
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        default:
            for fi139 := 0; fi139 < 15; fi139++ {
                var v140 float64 = -2.25
                var v141 uint16 = (v1 % (((v1 ^ genFn0("hello"))) | 1))
                v142 := GenS2{F0: v134}
                _ = fi139
                _ = v140
                _ = v141
                _ = v142
            }
        }
        var v143 int64 = 0
        _ = fi135
        _ = v143
    }
    for fi144 := 0; fi144 < 17; fi144++ {
        v145 := GenS7{F0: (fi144 + (fi144 & -9223372036854775808)), F1: (v2 & v2), F2: v134}
        var v146 uint8 = 127
        v147 := GenS2{F0: v134}
        _ = fi144
        _ = v145
        _ = v146
        _ = v147
    }
    var v148 int8 = v134
    for fi149 := 0; fi149 < 1; fi149++ {
        _ = len(v5)
        if v4 {
            v1 = genFn0(((("a" + (("\n\t" + ("" + "")) + "a")) + (("a" + (("" + ("foo bar" + "foo bar")) + "foo bar")) + (("\n\t" + "foo bar") + ""))) + ""))
        }
        _ = fi149
    }
    fmt.Println("v1:", v1)
    fmt.Println("v134:", v134)
    fmt.Println("v148:", v148)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v134
    _ = v148
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
