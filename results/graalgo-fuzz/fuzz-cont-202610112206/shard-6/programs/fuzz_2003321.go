package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 string; F2 uint16 }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int64; F1 uint8 }
type GenS6 struct { F0 int16; F1 uint; F2 uint64 }
type GenS7 struct { F0 bool; F1 int8 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int16; F1 int64; F2 int32 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 int16; F1 string; F2 uint64 }
type GenS12 struct { F0 uint16; F1 string }
type GenS13 struct { F0 uint64; F1 int; F2 string }
func genFn0() uint {
    v5 := []int{743}
    if true {
        for fi6 := 0; fi6 < 4; fi6++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi6
        }
        var v7 float64 = float64((1.5 / 3.14159))
        for fi8 := 0; fi8 < 18; fi8++ {
            v7 = (v7 - 0.0)
            for fi9 := 0; fi9 < 16; fi9++ {
                if false {
                    v10 := GenS1{F0: 42}
                    var v11 string = ""
                    var v12 int8 = 0
                    v13 := GenS2{F0: 816, F1: v11, F2: 143}
                    _ = v10
                    _ = v11
                    _ = v12
                    _ = v13
                }
                _ = fi9
            }
            _ = fi8
        }
        _ = v7
    } else {
        switch "hello" {
        case "a":
            if len(v5) > 0 {
                v5[0] = 34
            }
        case "foo bar":
            var v14 int64 = 215
            _ = v14
        }
    }
    _ = v5
    return 18446744073709551615
}
func genFn1(p0 string, p1 int64, p2 bool) int16 {
    for fi15 := 0; fi15 < 1; fi15++ {
        if (0 != 127) {
            p1 = p1
        } else {
            switch (fi15 ^ (fi15 % (((-(fi15))) | 1))) {
            case -1:
                var v16 uint16 = 42
                _ = v16
            case -2147483648:
                var v17 int8 = -1
                _ = v17
            default:
                switch (-(fi15)) {
                case -515:
                    v18 := GenS3{F0: 42, F1: -2147483648}
                    _ = v18
                case 2147483647:
                    v19 := []int{}
                    _ = v19
                case 4294967295:
                    if false {
                        var v20 bool = p2
                        switch (fi15 % ((128) | 1)) {
                        case 42:
                            var v21 uint = 2
                            _ = v21
                        case 127:
                            p0 = p0
                        case 4294967295:
                            var v22 uint32 = 320
                            _ = v22
                        }
                        _ = v20
                    } else {
                        v23 := []int{fi15, 61, 128}
                        _ = v23
                    }
                default:
                    var v24 uint16 = 255
                    _ = v24
                }
            }
            switch p0 {
            case "":
                var v25 uint8 = 128
                _ = v25
            case "\n\t":
                p2 = p2
            }
        }
        _ = fi15
    }
    if p2 {
        p1 -= p1
        fmt.Println(p0)
    }
    p2 = ((("hello" + (p0 + ((((p0 + p0) + p0) + (p0 + p0)) + p0))) + "") < p0)
    for fi26 := 0; fi26 < 6; fi26++ {
        v27 := GenS3{F0: 245, F1: 824}
        _ = fi26
        _ = v27
    }
    v28 := []int{0}
    _ = p0
    _ = p1
    _ = p2
    _ = v28
    return 53
}

func main() {
    var v1 string = ("" + "hello")
    var v2 int32 = 42
    v3 := []int{36}
    v4 := make(map[int]uint32)
    for fi29 := 0; fi29 < 18; fi29++ {
        v30 := v4[(fi29 * (fi29 >> ((255) & 63)))]
        v30 -= v30
        switch -509 {
        case 0:
            var v31 uint32 = 816
            _ = v31
        default:
            fmt.Println(v1)
        }
        _ = fi29
        _ = v30
    }
    fmt.Println(v2, v2)
    switch "" {
    case "":
        if true {
            v2 = v2
        } else {
            v1 = ("a" + "\n\t")
            v32 := GenS4{F0: 0}
            _ = v32
        }
    case "\n\t":
        var v33 uint32 = 734
        _ = v33
    case "hello":
        for fi34 := 0; fi34 < 6; fi34++ {
            v1 = (("hello" + v1) + v1)
            switch "foo bar" {
            case "":
                v1 = "hello"
            case "hello":
                if ((v2 != v2) || (false || (1 != -645))) {
                    var v35 bool = (!(false || true))
                    _ = v35
                } else {
                    v36 := GenS3{F0: -528, F1: v2}
                    if (!false) {
                        if (222 < 42) {
                            switch ("" + "a") {
                            case "":
                                if (3.14159 != 1.5) {
                                    for fi37 := 0; fi37 < 15; fi37++ {
                                        v1 = ((v1 + (v1 + (("a" + v1) + v1))) + (v1 + "foo bar"))
                                        switch 32767 {
                                        case 2147483647:
                                            fmt.Println(fi34)
                                        case 9223372036854775807:
                                            v2 -= v2
                                        default:
                                            var v38 int = fi34
                                            var v39 float64 = 3.14159
                                            if (v1 != (v1 + "")) {
                                                v2 -= v2
                                            }
                                            _ = v38
                                            _ = v39
                                        }
                                        _ = fi37
                                    }
                                    v1 = v1
                                    if true {
                                        v40 := GenS1{F0: 90}
                                        var v41 bool = (255 <= 32767)
                                        _ = v40
                                        _ = v41
                                    } else {
                                        v1 = ("hello" + v1)
                                        switch fi34 {
                                        case 128:
                                            var v42 int32 = (-(v2))
                                            _ = v42
                                        }
                                        var v43 int32 = (v2 | v2)
                                        _ = v43
                                    }
                                }
                            case "hello":
                                fmt.Println(v1, fi34)
                            default:
                                if false {
                                    switch (("" + "foo bar") + "") {
                                    case "":
                                        if false {
                                            switch fi34 {
                                            case -2147483648:
                                                fmt.Println(fi34, v1, v1)
                                            case 415:
                                                v44 := GenS4{F0: -1}
                                                _ = v44
                                            default:
                                                if (1e10 <= (1.5 - 3.14159)) {
                                                    if true {
                                                        _ = len(v4)
                                                        var v45 float64 = 3.14159
                                                        _ = v45
                                                    } else {
                                                        if (16 != 4294967295) {
                                                            for fi46 := 0; fi46 < 6; fi46++ {
                                                                fmt.Println(fi46, v2, fi34)
                                                                v47 := GenS4{F0: -1}
                                                                _ = fi46
                                                                _ = v47
                                                            }
                                                            v2 += v2
                                                        } else {
                                                            var v48 string = v1
                                                            fmt.Println(v48, fi34, fi34)
                                                            _ = v48
                                                        }
                                                    }
                                                }
                                            }
                                            var v49 uint8 = 255
                                            _ = v49
                                        } else {
                                            if false {
                                                if (!false) {
                                                    fmt.Println(fi34, v1)
                                                    v50 := GenS4{F0: 0}
                                                    _ = v50
                                                }
                                                v1 = v1
                                                var v51 int16 = 42
                                                _ = v51
                                            }
                                            switch (fi34 | fi34) {
                                            case 127:
                                                var v52 int8 = 1
                                                _ = v52
                                            case -42:
                                                var v53 uint64 = 0
                                                _ = v53
                                            case 32767:
                                                if len(v3) > 0 {
                                                    v3[0] = (^(fi34))
                                                }
                                            default:
                                                var v54 uint64 = 4294967295
                                                _ = v54
                                            }
                                        }
                                    case "hello":
                                        v4[fi34] = 2147483647
                                    }
                                    if (1 < 18446744073709551615) {
                                        for fi55 := 0; fi55 < 11; fi55++ {
                                            v56 := GenS5{F0: -9223372036854775808, F1: 2}
                                            if false {
                                                var v57 uint64 = 18446744073709551615
                                                switch (-(fi55)) {
                                                case 1:
                                                    fmt.Println(v2, v1)
                                                default:
                                                    v2 += v2
                                                }
                                                var v58 uint32 = 538
                                                _ = v57
                                                _ = v58
                                            } else {
                                                var v59 uint64 = 2
                                                _ = v59
                                            }
                                            _ = fi55
                                            _ = v56
                                        }
                                    } else {
                                        if ((v1 + (v1 + v1)) < ("foo bar" + "")) {
                                            for fi60 := 0; fi60 < 11; fi60++ {
                                                var v61 uint8 = 255
                                                _ = fi60
                                                _ = v61
                                            }
                                            if ((fi34 << ((255) & 63)) <= (fi34 >> ((18446744073709551615) & 63))) {
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                var v62 uint64 = 896
                                                if (fi34 != 0) {
                                                    for fi63 := 0; fi63 < 2; fi63++ {
                                                        v64 := GenS2{F0: 603, F1: (((v1 + v1) + "\n\t") + "hello"), F2: 887}
                                                        _ = fi63
                                                        _ = v64
                                                    }
                                                    v2 = (v2 * (v2 >> ((18446744073709551615) & 63)))
                                                } else {
                                                    if (!(609 != 1)) {
                                                        v3 = append(v3, 9223372036854775807)
                                                        fmt.Println(v62, v62, fi34)
                                                        fmt.Println(fi34, v62, v62)
                                                    }
                                                }
                                                _ = v62
                                            }
                                            v65 := GenS5{F0: -1, F1: 255}
                                            _ = v65
                                        } else {
                                            if (!false) {
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            }
                                            switch (fi34 & 127) {
                                            case -2147483648:
                                                var v66 uint32 = 160
                                                _ = v66
                                            case -643:
                                                v1 = v1
                                            case 75:
                                                delete(v4, fi34)
                                            }
                                            var v67 float32 = 0.0
                                            _ = v67
                                        }
                                    }
                                }
                            }
                        } else {
                            var v68 uint32 = 0
                            var v69 float64 = -2.25
                            _ = v68
                            _ = v69
                        }
                        v2 -= 42
                        var v70 int8 = -48
                        _ = v70
                    }
                    var v71 int32 = 2147483647
                    _ = v36
                    _ = v71
                }
            }
            _ = fi34
        }
    default:
        v2 = (v2 | (-(v2)))
    }
    if false {
        for fi72 := 0; fi72 < 8; fi72++ {
            var v73 uint32 = 42
            v4[fi72] = v73
            _ = fi72
            _ = v73
        }
        fmt.Println(v1)
        var v74 uint32 = 1000
        _ = v74
    }
    if false {
        var v75 int16 = 353
        switch "a" {
        case "foo bar":
            v76 := GenS6{F0: (v75 / (((v75 ^ genFn1(v1, 42, false))) | 1)), F1: 255, F2: 4294967295}
            _ = v76
        default:
            if false {
                var v77 int8 = 19
                fmt.Println(v2)
                _ = v77
            }
        }
        var v78 int = 42
        _ = v75
        _ = v78
    }
    v2 = (v2 >> ((242) & 63))
    if (!true) {
        if true {
            if len(v3) > 0 {
                v3[0] = 814
            }
            var v79 string = v1
            switch 504 {
            case 128:
                fmt.Println(v79, v79, v1)
            case 127:
                fmt.Println(v79, v2)
            }
            _ = v79
        }
    } else {
        v80 := GenS7{F0: (v1 >= ("\n\t" + (v1 + (v1 + "hello")))), F1: -128}
        switch ("a" + "hello") {
        case "foo bar":
            v1 = v1
        case "a":
            v2 = (-(v2))
        default:
            var v81 int = 2147483647
            _ = v81
        }
        if (true && (false || false)) {
            v2 += v2
        } else {
            for fi82 := 0; fi82 < 13; fi82++ {
                if (0 == 255) {
                    fmt.Println(fi82, v1)
                } else {
                    if true {
                        var v83 int16 = 1
                        var v84 uint8 = 255
                        _ = v83
                        _ = v84
                    }
                    v85 := GenS8{F0: (649 <= (fi82 & (fi82 / ((fi82) | 1))))}
                    _ = v85
                }
                v1 = ("hello" + v1)
                v2 = 1
                _ = fi82
            }
            switch "a" {
            case "":
                v2 = (v2 - 42)
            case "foo bar":
                var v86 bool = (0 == 58)
                _ = v86
            case "a":
                for fi87 := 0; fi87 < 13; fi87++ {
                    for fi88 := 0; fi88 < 7; fi88++ {
                        v2 += -2147483648
                        switch v1 {
                        case "\n\t":
                            if len(v3) > 0 {
                                v3[0] = fi87
                            }
                        case "":
                            var v89 string = "\n\t"
                            _ = v89
                        case "foo bar":
                            if (0 < 0) {
                                var v90 int8 = 1
                                _ = v90
                            } else {
                                v1 = (v1 + ((("" + v1) + ("a" + "hello")) + v1))
                                v2 += v2
                                var v91 float64 = (1.5 + 1e10)
                                _ = v91
                            }
                            var v92 float64 = 1.5
                            _ = v92
                        default:
                            v93 := GenS9{F0: -625, F1: 9223372036854775807, F2: (v2 ^ v2)}
                            _ = v93
                        }
                        _ = fi88
                    }
                    if true {
                        var v94 uint32 = 2
                        var v95 int = fi87
                        if len(v3) > 0 {
                            v3[0] = v95
                        }
                        _ = v94
                        _ = v95
                    } else {
                        fmt.Println(fi87, v2, v1)
                    }
                    _ = fi87
                }
            }
        }
        _ = v80
    }
    if true {
        switch (v1 + v1) {
        case "hello":
            v1 = ""
        case "\n\t":
            if (0.1 <= (-2.25 / ((3.14159 / (1.5 + (((0.0 / 1e10) * 0.1) - (0.0 * float64(0.0))))) / (0.0 - ((float64(-2.25) / (float64(3.14159) - 1e10)) - (1e10 + -2.25)))))) {
                v2 = (^(v2))
            } else {
                for fi96 := 0; fi96 < 12; fi96++ {
                    if false {
                        v2 -= v2
                        var v97 int8 = 109
                        _ = v97
                    } else {
                        v98 := GenS4{F0: -32768}
                        if true {
                            v99 := GenS3{F0: -32768, F1: v2}
                            _ = v99
                        }
                        _ = v98
                    }
                    var v100 string = v1
                    _ = fi96
                    _ = v100
                }
            }
        }
        fmt.Println(v1, v2)
        fmt.Println(v2, v1)
    }
    v2 = (v2 << ((65535) & 63))
    v101 := GenS9{F0: 42, F1: 9223372036854775807, F2: (v2 % ((0) | 1))}
    if true {
        v102 := GenS10{F0: 999}
        _ = v102
    } else {
        if (42 > -9223372036854775808) {
            for fi103 := 0; fi103 < 8; fi103++ {
                var v104 float64 = (3.14159 * 0.1)
                fmt.Println(fi103)
                for fi105 := 0; fi105 < 11; fi105++ {
                    v106 := GenS3{F0: genFn1((v1 + ("hello" + v1)), 9223372036854775807, (true || false)), F1: (v2 << ((0) & 63))}
                    if (true && true) {
                        v104 -= 3.14159
                        v1 = (v1 + v1)
                    } else {
                        var v107 uint16 = 255
                        _ = v107
                    }
                    if (2 == 195) {
                        fmt.Println(v104, v2, v104)
                        v1 = "\n\t"
                    } else {
                        switch -867 {
                        case 42:
                            v2 += (v2 / ((v2) | 1))
                        case 1:
                            v108 := GenS10{F0: 22}
                            _ = v108
                        }
                        v1 = "foo bar"
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        fmt.Println(v1, v2)
                    }
                    if len(v3) > 0 {
                        v3[0] = fi103
                    }
                    _ = fi105
                    _ = v106
                }
                _ = fi103
                _ = v104
            }
            if false {
                switch 288 {
                case 32767:
                    v109 := GenS5{F0: 0, F1: 2}
                    _ = v109
                case 2147483647:
                    var v110 int16 = -1
                    _ = v110
                case -879:
                    v2 = 2147483647
                }
                if (42 < 2) {
                    if ((363 < 1) && true) {
                        v111 := GenS11{F0: 32767, F1: v1, F2: 2}
                        var v112 string = v1
                        if ((!((((v1 + (v112 + (v1 + ("hello" + (v1 + v1))))) + ((((("foo bar" + (v1 + (v1 + v112))) + ("foo bar" + (((v112 + "") + v1) + v1))) + v1) + ("\n\t" + "hello")) + v1)) + "foo bar") >= (v112 + ("" + v1)))) && ((2 == genFn0()) || false)) {
                            var v113 float64 = (3.14159 / -2.25)
                            v1 = ((v112 + v1) + (v1 + "hello"))
                            v113 -= v113
                            _ = v113
                        } else {
                            var v114 float32 = float32(1e10)
                            v115 := v4[9223372036854775807]
                            var v116 int32 = (v2 / (((v2 & (v2 * (v2 * (^(v2)))))) | 1))
                            _ = v114
                            _ = v115
                            _ = v116
                        }
                        _ = v111
                        _ = v112
                    }
                }
                fmt.Println(v2, v1)
            } else {
                var v117 uint64 = 42
                var v118 uint8 = 0
                _ = v117
                _ = v118
            }
        }
        var v119 float64 = 1.5
        for fi120 := 0; fi120 < 3; fi120++ {
            v1 = v1
            if true {
                switch fi120 {
                case 128:
                    v121 := v4[(fi120 << ((661) & 63))]
                    _ = v121
                case 2147483647:
                    var v122 int64 = 42
                    _ = v122
                case 4294967295:
                    var v123 uint16 = 0
                    _ = v123
                }
                fmt.Println(v2, fi120, v1)
                v124 := v4[4294967295]
                _ = v124
            } else {
                for fi125 := 0; fi125 < 19; fi125++ {
                    if ((218 == 0) || (v1 >= "")) {
                        v126 := GenS12{F0: 32767, F1: (v1 + v1)}
                        _ = v126
                    } else {
                        v127 := GenS8{F0: (3.14159 <= ((1e10 - -2.25) + 1e10))}
                        _ = v127
                    }
                    _ = fi125
                }
                v128 := GenS13{F0: 271, F1: fi120, F2: (((v1 + "hello") + "\n\t") + (("" + "foo bar") + v1))}
                _ = v128
            }
            _ = fi120
        }
        _ = v119
    }
    if (836 <= -9223372036854775808) {
        v129 := GenS6{F0: 42, F1: genFn0(), F2: 42}
        _ = v129
    } else {
        fmt.Println(v1, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v101
    _ = v2
    _ = v3
    _ = v4
}
