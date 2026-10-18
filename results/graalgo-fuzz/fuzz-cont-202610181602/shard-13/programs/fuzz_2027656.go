package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 int8; F1 int32; F2 uint8 }
type GenS4 struct { F0 int; F1 string; F2 uint8 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 bool; F1 uint }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int; F1 int16 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int16; F1 uint; F2 uint32 }
type GenS11 struct { F0 uint8; F1 uint32 }
type GenS12 struct { F0 uint16; F1 uint64 }
type GenS13 struct { F0 uint64 }

func main() {
    var v1 int8 = 42
    v2 := make(map[int]bool)
    v3 := []uint32{}
    var v4 int16 = 1
    v5 := []int{128, -768, -9223372036854775808}
    var v6 int = 4294967295
    v1 = v1
    for fi7 := 0; fi7 < 18; fi7++ {
        if (0 < 42) {
            var v8 uint64 = 2
            fmt.Println(v6, v8)
            for fi9 := 0; fi9 < 15; fi9++ {
                v8 = (v8 ^ 18446744073709551615)
                _ = fi9
            }
            _ = v8
        } else {
            var v10 int16 = v4
            var v11 int = (fi7 * 2147483647)
            var v12 int32 = -1
            _ = v10
            _ = v11
            _ = v12
        }
        v3 = append(v3, 470)
        for fi13 := 0; fi13 < 15; fi13++ {
            if ((!true) && (0 > 655)) {
                switch "" {
                case "\n\t":
                    if (((("\n\t" + "\n\t") <= "hello") && (!true)) || true) {
                        switch 436 {
                        case 254:
                            var v14 float32 = float32(1.5)
                            _ = v14
                        case 9223372036854775807:
                            v1 = v1
                        case 1:
                            var v15 int64 = 262
                            _ = v15
                        }
                        if (2147483647 == 42) {
                            var v16 int = (^(fi7))
                            switch (("foo bar" + "foo bar") + "a") {
                            case "foo bar":
                                v1 -= (v1 % ((-128) | 1))
                            case "":
                                fmt.Println(v4)
                            default:
                                fmt.Println(v16, v6)
                            }
                            _ = v16
                        } else {
                            switch (^(fi13)) {
                            case -44:
                                if (0 > 1) {
                                    var v17 uint64 = 4294967295
                                    _ = v17
                                } else {
                                    switch "hello" {
                                    case "hello":
                                        switch fi7 {
                                        case 127:
                                            v18 := GenS1{F0: v4}
                                            _ = v18
                                        default:
                                            var v19 int64 = 0
                                            _ = v19
                                        }
                                        fmt.Println(v4, v4, v6)
                                        switch "foo bar" {
                                        case "":
                                            if ((!true) || false) {
                                                v4 = (v4 + -32768)
                                                v20 := GenS1{F0: (-(v4))}
                                                var v21 float32 = 0.1
                                                _ = v20
                                                _ = v21
                                            }
                                        case "a":
                                            switch "hello" {
                                            case "\n\t":
                                                var v22 int64 = -246
                                                _ = v22
                                            case "a":
                                                if len(v3) > 0 {
                                                    v3[0] = 2147483648
                                                }
                                            default:
                                                switch fi7 {
                                                case 1:
                                                    var v23 int8 = v1
                                                    _ = v23
                                                case 185:
                                                    v6 = (fi13 + v6)
                                                    var v24 int8 = 42
                                                    _ = v24
                                                case 32768:
                                                    switch fi7 {
                                                    case -1:
                                                        var v25 int32 = 1
                                                        _ = v25
                                                    case -9223372036854775808:
                                                        var v26 float32 = 1.5
                                                        _ = v26
                                                    default:
                                                        v1 += v1
                                                        switch (v6 >> ((1) & 63)) {
                                                        case 740:
                                                            v4 = 213
                                                            var v27 float32 = 3.14159
                                                            v27 += (float32(float64(((0.1 / v27) + float32(3.14159)))) * (3.14159 / 0.1))
                                                            var v28 int16 = v4
                                                            var v29 int16 = v4
                                                            _ = v27
                                                            _ = v28
                                                            _ = v29
                                                        }
                                                        v6 += (fi13 >> ((255) & 63))
                                                        v6 = (fi7 * fi7)
                                                        fmt.Println(fi7)
                                                    }
                                                default:
                                                    var v30 int16 = (v4 | v4)
                                                    _ = v30
                                                }
                                            }
                                        }
                                    case "foo bar":
                                        fmt.Println(v4, v6)
                                    }
                                }
                            case -9223372036854775808:
                                if true {
                                    v4 = (v4 & v4)
                                    switch fi7 {
                                    case 890:
                                        fmt.Println(v6, fi7)
                                    case -9223372036854775808:
                                        v4 = v4
                                    case 32768:
                                        v1 -= v1
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    default:
                                        if ((!false) || ((v1 < (^(v1))) || false)) {
                                            switch "" {
                                            case "foo bar":
                                                var v31 uint32 = 778
                                                _ = v31
                                            case "a":
                                                if true {
                                                    switch fi7 {
                                                    case 128:
                                                        var v32 int32 = 415
                                                        _ = v32
                                                    case -2147483648:
                                                        v4 -= 1000
                                                    case 672:
                                                        if false {
                                                            v33 := GenS1{F0: (v4 << ((65535) & 63))}
                                                            v4 = (v4 & v4)
                                                            _ = v33
                                                        }
                                                    default:
                                                        if ((v6 * 1) <= int(v1)) {
                                                            switch "" {
                                                            case "foo bar":
                                                                var v34 int32 = 496
                                                                _ = v34
                                                            case "\n\t":
                                                                v2[(fi7 & fi7)] = (v4 >= (v4 << ((161) & 63)))
                                                            case "":
                                                                v6 -= fi13
                                                            default:
                                                                v4 = (v4 + (v4 | 65))
                                                                var v35 uint64 = 289
                                                                _ = v35
                                                            }
                                                        } else {
                                                            if false {
                                                                var v36 uint64 = 955
                                                                v36 = v36
                                                                if (1000 < -32768) {
                                                                    var v37 uint32 = 2147483648
                                                                    _ = v37
                                                                }
                                                                _ = v36
                                                            } else {
                                                                if len(v5) >= 2 {
                                                                    v5 = v5[1:len(v5)]
                                                                }
                                                                v38 := GenS2{F0: 2147483648, F1: 32767}
                                                                delete(v2, 127)
                                                                fmt.Println(v1, fi13)
                                                                v4 += (v4 - 118)
                                                                _ = v38
                                                            }
                                                        }
                                                    }
                                                    var v39 bool = true
                                                    _ = v39
                                                } else {
                                                    var v40 uint = 4294967295
                                                    _ = v40
                                                }
                                            default:
                                                switch (-(fi7)) {
                                                case -894:
                                                    var v41 int16 = 419
                                                    _ = v41
                                                }
                                            }
                                            fmt.Println(fi7)
                                        }
                                    }
                                    if (1 > 2) {
                                        v42 := GenS1{F0: v4}
                                        _ = v42
                                    } else {
                                        v4 -= v4
                                        v43 := v2[fi7]
                                        if v43 {
                                            var v44 float32 = (0.0 / 1.0)
                                            if (!(165 >= 255)) {
                                                _ = len(v2)
                                                v43 = (0 < 0)
                                            } else {
                                                if (!(!v43)) {
                                                    if (!v43) {
                                                        var v45 int8 = -1
                                                        switch "a" {
                                                        case "":
                                                            v46 := v2[v6]
                                                            v47 := GenS2{F0: 1, F1: 32767}
                                                            _ = v46
                                                            _ = v47
                                                        case "foo bar":
                                                            switch (fi7 >> ((632) & 63)) {
                                                            case -9223372036854775808:
                                                                fmt.Println(fi13, v6, v44)
                                                            case 904:
                                                                fmt.Println(fi7, v43)
                                                            case 9223372036854775807:
                                                                if false {
                                                                    var v48 uint64 = 932
                                                                    var v49 uint8 = 1
                                                                    if false {
                                                                        var v50 uint64 = v48
                                                                        v5 = append(v5, (fi7 / (((fi7 - fi7)) | 1)))
                                                                        _ = v50
                                                                    }
                                                                    _ = v48
                                                                    _ = v49
                                                                } else {
                                                                    v1 += v1
                                                                }
                                                            }
                                                        }
                                                        v51 := v2[v6]
                                                        var v52 float64 = -2.25
                                                        v53 := GenS1{F0: v4}
                                                        _ = v45
                                                        _ = v51
                                                        _ = v52
                                                        _ = v53
                                                    }
                                                } else {
                                                    if (-2.25 < 3.14159) {
                                                        if true {
                                                            var v54 float32 = v44
                                                            _ = v54
                                                        } else {
                                                            var v55 uint8 = 127
                                                            delete(v2, -1)
                                                            var v56 int32 = -197
                                                            v57 := v2[(fi13 % (((^(v6))) | 1))]
                                                            _ = v55
                                                            _ = v56
                                                            _ = v57
                                                        }
                                                        fmt.Println(fi7, v43)
                                                        v3 = append(v3, 745)
                                                    } else {
                                                        v58 := GenS1{F0: (v4 - (v4 + 32767))}
                                                        if (v43 && v43) {
                                                            var v59 string = ((("foo bar" + ("\n\t" + "hello")) + "foo bar") + ("foo bar" + "foo bar"))
                                                            delete(v2, 32767)
                                                            v1 = 0
                                                            _ = v59
                                                        } else {
                                                            switch "foo bar" {
                                                            case "":
                                                                if (0 >= 65535) {
                                                                    var v60 string = ""
                                                                    v1 -= (v1 - 100)
                                                                    v44 = (0.1 + v44)
                                                                    _ = v60
                                                                } else {
                                                                    fmt.Println(v4, v6)
                                                                }
                                                            case "foo bar":
                                                                if ("foo bar" > "a") {
                                                                    v1 = (v1 + (v1 + (v1 >> ((363) & 63))))
                                                                    v6 = fi7
                                                                } else {
                                                                    v1 = (v1 ^ -55)
                                                                    v43 = v43
                                                                    v6 = (v6 & (^(fi7)))
                                                                }
                                                            case "a":
                                                                v61 := GenS2{F0: 65535, F1: 65535}
                                                                _ = v61
                                                            }
                                                            if (((v44 - (-2.25 - 0.0)) / 1.0) == v44) {
                                                                fmt.Println(v6)
                                                            }
                                                        }
                                                        _ = v58
                                                    }
                                                    switch (fi7 >> ((1) & 63)) {
                                                    case 4294967295:
                                                        if (v6 == -2147483648) {
                                                            v62 := GenS3{F0: (v1 - -128), F1: -1, F2: 127}
                                                            _ = v62
                                                        }
                                                    case -129:
                                                        switch "\n\t" {
                                                        case "":
                                                            if v43 {
                                                                switch ("" + ("a" + "hello")) {
                                                                case "\n\t":
                                                                    if (!v43) {
                                                                        v43 = v43
                                                                        v63 := GenS4{F0: (fi13 % ((v6) | 1)), F1: ("foo bar" + "hello"), F2: 42}
                                                                        v6 += (fi13 >> ((1) & 63))
                                                                        _ = v63
                                                                    } else {
                                                                        var v64 int = (fi13 % ((fi7) | 1))
                                                                        v43 = false
                                                                        var v65 int16 = (v4 + 407)
                                                                        _ = v64
                                                                        _ = v65
                                                                    }
                                                                case "a":
                                                                    fmt.Println(v1, v4)
                                                                case "":
                                                                    if true {
                                                                        switch (^(fi7)) {
                                                                        case -2147483648:
                                                                            if ((false || v43) && v43) {
                                                                                fmt.Println(v1, v1, v43)
                                                                            }
                                                                        case 127:
                                                                            switch "a" {
                                                                            case "a":
                                                                                if true {
                                                                                    switch -862 {
                                                                                    case 0:
                                                                                        switch "" {
                                                                                        case "hello":
                                                                                            v44 = ((v44 + v44) - v44)
                                                                                        case "foo bar":
                                                                                            var v66 int32 = 243
                                                                                            _ = v66
                                                                                        }
                                                                                    case 9223372036854775807:
                                                                                        switch v6 {
                                                                                        case -863:
                                                                                            _ = len(v2)
                                                                                        case 0:
                                                                                            var v67 int8 = (-(v1))
                                                                                            _ = v67
                                                                                        case 2147483647:
                                                                                            var v68 float32 = float32((1.5 + 3.14159))
                                                                                            if len(v3) > 0 {
                                                                                                v3[0] = 4294967295
                                                                                            }
                                                                                            _ = v68
                                                                                        default:
                                                                                            var v69 bool = v43
                                                                                            _ = v69
                                                                                        }
                                                                                        var v70 uint8 = 2
                                                                                        _ = v70
                                                                                    default:
                                                                                        v1 -= v1
                                                                                    }
                                                                                    v71 := GenS2{F0: 2147483648, F1: 291}
                                                                                    _ = v71
                                                                                } else {
                                                                                    if (v43 && (((fi13 >> ((989) & 63)) > v6) && ((-318 <= 234) && false))) {
                                                                                        v4 += (v4 - 42)
                                                                                        v6 = v6
                                                                                        if (!((0 > 65535) && ((!false) && v43))) {
                                                                                            var v72 uint64 = 2
                                                                                            if (false && ((0 != 4294967295) || v43)) {
                                                                                                var v73 uint16 = 32767
                                                                                                _ = v73
                                                                                            } else {
                                                                                                fmt.Println(v72)
                                                                                                v72 += (v72 % ((1) | 1))
                                                                                            }
                                                                                            _ = len(v5)
                                                                                            _ = v72
                                                                                        }
                                                                                        var v74 int8 = v1
                                                                                        _ = v74
                                                                                    } else {
                                                                                        v75 := GenS5{F0: -312}
                                                                                        fmt.Println(fi7, fi13, v43)
                                                                                        _ = v75
                                                                                    }
                                                                                    if (v43 || (!(1 <= 185))) {
                                                                                        var v76 int8 = 1
                                                                                        v1 -= v76
                                                                                        v77 := GenS6{F0: true, F1: 255}
                                                                                        _ = v76
                                                                                        _ = v77
                                                                                    }
                                                                                    var v78 int = 32767
                                                                                    v1 -= (v1 >> ((1) & 63))
                                                                                    if (true || v43) {
                                                                                        switch (("hello" + "foo bar") + "a") {
                                                                                        case "a":
                                                                                            var v79 int64 = 1
                                                                                            _ = v79
                                                                                        case "foo bar":
                                                                                            var v80 int8 = v1
                                                                                            _ = v80
                                                                                        case "":
                                                                                            var v81 uint32 = 2
                                                                                            _ = v81
                                                                                        default:
                                                                                            v44 -= v44
                                                                                        }
                                                                                    } else {
                                                                                        v82 := GenS1{F0: v4}
                                                                                        if true {
                                                                                            fmt.Println(fi13, v44)
                                                                                        }
                                                                                        if (0 != 0) {
                                                                                            var v83 uint8 = 164
                                                                                            _ = v83
                                                                                        } else {
                                                                                            if v43 {
                                                                                                v43 = (42 == 2147483647)
                                                                                                v43 = (v44 != float32(0.1))
                                                                                                v1 += (v1 * (v1 << ((632) & 63)))
                                                                                            }
                                                                                        }
                                                                                        _ = v82
                                                                                    }
                                                                                    _ = v78
                                                                                }
                                                                            default:
                                                                                if v43 {
                                                                                    v1 = v1
                                                                                    fmt.Println(v4, v44)
                                                                                    v43 = v43
                                                                                } else {
                                                                                    if (42 != (v1 << ((0) & 63))) {
                                                                                        switch 2147483647 {
                                                                                        case 599:
                                                                                            if true {
                                                                                                var v84 int16 = (v4 % ((0) | 1))
                                                                                                _ = v84
                                                                                            } else {
                                                                                                v85 := GenS4{F0: 42, F1: (("foo bar" + "hello") + "a"), F2: 127}
                                                                                                _ = v85
                                                                                            }
                                                                                        }
                                                                                        _ = len(v2)
                                                                                        var v86 uint64 = 18446744073709551615
                                                                                        _ = v86
                                                                                    }
                                                                                    var v87 uint16 = 1
                                                                                    if true {
                                                                                        var v88 bool = false
                                                                                        switch "\n\t" {
                                                                                        case "\n\t":
                                                                                            if v43 {
                                                                                                v89 := GenS7{F0: 4294967295}
                                                                                                v90 := GenS8{F0: (fi7 * v6), F1: (v4 | v4)}
                                                                                                if ((v43 && v88) || v43) {
                                                                                                    switch fi7 {
                                                                                                    case 4294967295:
                                                                                                        var v91 int = v6
                                                                                                        _ = v91
                                                                                                    case -1:
                                                                                                        v87 -= v87
                                                                                                    case -2147483648:
                                                                                                        v87 = 509
                                                                                                    default:
                                                                                                        v3 = append(v3, 2)
                                                                                                    }
                                                                                                    v88 = ((0.1 + (3.14159 / (3.14159 + (1.5 - (0.0 / 1.5))))) >= (0.0 + 0.1))
                                                                                                }
                                                                                                _ = v89
                                                                                                _ = v90
                                                                                            }
                                                                                        case "":
                                                                                            v44 += 1e10
                                                                                        case "hello":
                                                                                            if (227 != 255) {
                                                                                                var v92 uint16 = v87
                                                                                                v93 := GenS5{F0: -650}
                                                                                                v6 -= fi13
                                                                                                var v94 uint8 = 127
                                                                                                var v95 int = fi13
                                                                                                v96 := GenS6{F0: (!false), F1: 65535}
                                                                                                _ = v92
                                                                                                _ = v93
                                                                                                _ = v94
                                                                                                _ = v95
                                                                                                _ = v96
                                                                                            }
                                                                                        }
                                                                                        var v97 int = (v6 ^ fi13)
                                                                                        _ = v88
                                                                                        _ = v97
                                                                                    } else {
                                                                                        var v98 uint = 18446744073709551615
                                                                                        fmt.Println(fi7, fi13)
                                                                                        _ = v98
                                                                                    }
                                                                                    _ = v87
                                                                                }
                                                                            }
                                                                        default:
                                                                            v99 := GenS9{F0: v43}
                                                                            _ = v99
                                                                        }
                                                                    } else {
                                                                        var v100 uint8 = 1
                                                                        _ = v100
                                                                    }
                                                                }
                                                                v1 += (v1 % ((-1) | 1))
                                                            }
                                                        case "foo bar":
                                                            var v101 uint = 1
                                                            _ = v101
                                                        case "a":
                                                            fmt.Println(v1, v4, fi7)
                                                        default:
                                                            if v43 {
                                                                var v102 int = (fi7 / ((fi7) | 1))
                                                                var v103 int = v102
                                                                _ = v102
                                                                _ = v103
                                                            }
                                                        }
                                                    default:
                                                        v44 += float32(0.0)
                                                    }
                                                    if (v43 && (584 > 42)) {
                                                        var v104 uint32 = 2147483648
                                                        if (!(true || v43)) {
                                                            v105 := GenS10{F0: 884, F1: 441, F2: v104}
                                                            var v106 int32 = 2147483647
                                                            _ = v105
                                                            _ = v106
                                                        } else {
                                                            if (42 >= 970) {
                                                                var v107 uint16 = 255
                                                                _ = v107
                                                            } else {
                                                                var v108 int64 = int64(v6)
                                                                _ = v108
                                                            }
                                                            var v109 float32 = float32(1.5)
                                                            fmt.Println(v1, v43)
                                                            _ = v109
                                                        }
                                                        var v110 uint = 18446744073709551615
                                                        _ = v104
                                                        _ = v110
                                                    }
                                                }
                                                fmt.Println(fi13, v43)
                                                var v111 int64 = 963
                                                _ = v111
                                            }
                                            _ = v44
                                        }
                                        _ = v43
                                    }
                                    v1 = 127
                                } else {
                                    v3 = append(v3, 111)
                                    var v112 float64 = 0.0
                                    v1 -= (-(v1))
                                    _ = v112
                                }
                            case 1:
                                v113 := GenS11{F0: 42, F1: 2}
                                _ = v113
                            }
                            v6 -= fi13
                        }
                    } else {
                        v3 = append(v3, 471)
                        v4 += int16(fi7)
                        var v114 int32 = 42
                        _ = v114
                    }
                }
                var v115 uint = 4294967295
                var v116 int8 = (v1 << ((65535) & 63))
                delete(v2, fi7)
                v117 := GenS11{F0: 248, F1: 2147483647}
                _ = v115
                _ = v116
                _ = v117
            }
            v3 = append(v3, 42)
            v118 := GenS12{F0: 8, F1: 4294967295}
            var v119 int16 = v4
            v120 := GenS13{F0: 730}
            _ = fi13
            _ = v118
            _ = v119
            _ = v120
        }
        _ = fi7
    }
    var v121 int32 = -1
    if false {
        v121 -= v121
    }
    v4 = -710
    v2[(v6 >> ((0) & 63))] = true
    v122 := GenS2{F0: 0, F1: 0}
    var v123 string = "foo bar"
    for fi124 := 0; fi124 < 16; fi124++ {
        var v125 bool = false
        _ = fi124
        _ = v125
    }
    var v126 uint = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v121:", v121)
    fmt.Println("v123:", v123)
    fmt.Println("v126:", v126)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v121
    _ = v122
    _ = v123
    _ = v126
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
