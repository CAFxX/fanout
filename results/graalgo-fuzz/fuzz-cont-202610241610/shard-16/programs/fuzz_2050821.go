package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 uint16; F2 int16 }
type GenS3 struct { F0 uint16; F1 int16; F2 int64 }
type GenS4 struct { F0 uint8; F1 int32; F2 uint }
type GenS5 struct { F0 int; F1 int }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint64; F1 uint; F2 bool }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int; F1 uint8 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 int8 }

func main() {
    v1 := []uint32{276, 1}
    var v2 int8 = 42
    v3 := make(map[int]uint32)
    v4 := []uint32{42, 2}
    var v5 bool = true
    var v6 float64 = 0.0
    fmt.Println(v6, v6, v5)
    for fi7 := 0; fi7 < 5; fi7++ {
        if v5 {
            switch (fi7 % ((32) | 1)) {
            case 4294967295:
                for fi8 := 0; fi8 < 18; fi8++ {
                    var v9 bool = true
                    v10 := GenS1{F0: 1}
                    if (16 > 2) {
                        fmt.Println(fi7)
                        if v5 {
                            switch fi7 {
                            case 2147483647:
                                v6 = float64(0.0)
                            case -875:
                                if (0 != -2147483648) {
                                    v9 = v9
                                } else {
                                    var v11 float64 = 0.1
                                    _ = v11
                                }
                            case 4294967295:
                                var v12 uint32 = 42
                                _ = v12
                            default:
                                fmt.Println(fi8)
                            }
                        } else {
                            v13 := GenS1{F0: -1}
                            var v14 uint8 = 189
                            _ = v13
                            _ = v14
                        }
                        if (!(-2147483648 != fi8)) {
                            _ = len(v3)
                            v5 = true
                            v15 := GenS1{F0: 0}
                            switch fi7 {
                            case 32768:
                                switch ((("hello" + (((("foo bar" + ("a" + "hello")) + ("foo bar" + ("a" + "hello"))) + "\n\t") + ("foo bar" + ("a" + "\n\t")))) + "hello") + ("\n\t" + ((("\n\t" + "") + (((("\n\t" + "foo bar") + "hello") + "foo bar") + "hello")) + ("a" + "foo bar")))) {
                                case "hello":
                                    fmt.Println(fi7)
                                    v2 = 42
                                    v2 += v2
                                case "\n\t":
                                    switch fi8 {
                                    case -37:
                                        if (((128 > 42) && true) && ((0 < 1) && true)) {
                                            if (float32(v6) < (-2.25 / 1e10)) {
                                                v16 := v3[fi7]
                                                _ = v16
                                            } else {
                                                var v17 int = 32768
                                                if v9 {
                                                    if v9 {
                                                        var v18 int32 = 638
                                                        if v5 {
                                                            var v19 int8 = (v2 ^ v2)
                                                            var v20 string = "foo bar"
                                                            v6 = (v6 * v6)
                                                            _ = v19
                                                            _ = v20
                                                        } else {
                                                            var v21 float32 = float32(v6)
                                                            if v5 {
                                                                var v22 uint16 = 255
                                                                var v23 uint32 = 637
                                                                _ = v22
                                                                _ = v23
                                                            } else {
                                                                var v24 uint8 = 152
                                                                _ = v24
                                                            }
                                                            _ = v21
                                                        }
                                                        v25 := GenS1{F0: -365}
                                                        _ = v18
                                                        _ = v25
                                                    }
                                                    fmt.Println(fi8, fi8)
                                                } else {
                                                    v26 := GenS1{F0: 1000}
                                                    _ = v26
                                                }
                                                var v27 uint16 = 1
                                                var v28 float64 = v6
                                                _ = v17
                                                _ = v27
                                                _ = v28
                                            }
                                            v9 = v5
                                        } else {
                                            var v29 uint = 1
                                            var v30 uint32 = 305
                                            _ = v29
                                            _ = v30
                                        }
                                    }
                                case "foo bar":
                                    var v31 int64 = -9223372036854775808
                                    _ = v31
                                }
                            default:
                                var v32 string = ((("foo bar" + (("hello" + "a") + "\n\t")) + "\n\t") + ("\n\t" + (("" + "hello") + (("foo bar" + "a") + ""))))
                                _ = v32
                            }
                            var v33 float64 = v6
                            if (((v6 + (v6 + (1.5 - 1e10))) + v33) <= v33) {
                                v3[fi8] = 263
                                delete(v3, (-(fi7)))
                                if (true || ((fi7 >> ((389) & 63)) < fi7)) {
                                    v5 = v5
                                    v34 := GenS2{F0: -1, F1: 805, F2: 1000}
                                    _ = v34
                                } else {
                                    if v9 {
                                        var v35 int64 = 927
                                        _ = v35
                                    }
                                    v6 = (v33 + v6)
                                    var v36 string = "a"
                                    var v37 uint64 = 954
                                    _ = v36
                                    _ = v37
                                }
                                v38 := GenS3{F0: 842, F1: 0, F2: -388}
                                if (-2147483648 == fi8) {
                                    if v9 {
                                        v3[fi8] = 2
                                    }
                                    v39 := GenS4{F0: 128, F1: 2147483647, F2: 65535}
                                    if true {
                                        v40 := v3[fi8]
                                        var v41 int = (fi7 / (((fi8 & -1)) | 1))
                                        _ = v40
                                        _ = v41
                                    } else {
                                        switch fi7 {
                                        case 42:
                                            if false {
                                                fmt.Println(v33)
                                                if (!(v9 && true)) {
                                                    if v9 {
                                                        v42 := v3[fi8]
                                                        var v43 int = fi7
                                                        var v44 uint8 = 0
                                                        _ = v42
                                                        _ = v43
                                                        _ = v44
                                                    } else {
                                                        v2 += (^(v2))
                                                    }
                                                    if true {
                                                        if false {
                                                            fmt.Println(fi8, fi8)
                                                            v9 = v9
                                                        }
                                                        var v45 int32 = 2147483647
                                                        if (177 != 2) {
                                                            v46 := GenS5{F0: (fi7 - fi7), F1: (^(fi8))}
                                                            _ = v46
                                                        }
                                                        _ = v45
                                                    } else {
                                                        var v47 uint32 = 4294967295
                                                        var v48 string = "a"
                                                        _ = v47
                                                        _ = v48
                                                    }
                                                    switch "a" {
                                                    case "hello":
                                                        switch "foo bar" {
                                                        case "hello":
                                                            var v49 uint16 = 2
                                                            _ = v49
                                                        default:
                                                            var v50 uint = 255
                                                            _ = v50
                                                        }
                                                    default:
                                                        var v51 int32 = 44
                                                        _ = v51
                                                    }
                                                }
                                                if true {
                                                    v5 = (766 >= 0)
                                                    v6 = v33
                                                    var v52 int = fi8
                                                    var v53 int = 32767
                                                    _ = v52
                                                    _ = v53
                                                }
                                            }
                                        case 4294967295:
                                            var v54 float32 = -2.25
                                            _ = v54
                                        case -1:
                                            v55 := GenS2{F0: -1, F1: 604, F2: 32767}
                                            _ = v55
                                        }
                                        var v56 bool = (!(1 > 2))
                                        if (true || false) {
                                            v57 := GenS1{F0: -32768}
                                            _ = len(v3)
                                            _ = v57
                                        }
                                        _ = v56
                                    }
                                    fmt.Println(v6, fi8)
                                    var v58 string = "hello"
                                    _ = v39
                                    _ = v58
                                }
                                _ = v38
                            }
                            _ = v15
                            _ = v33
                        } else {
                            switch "hello" {
                            case "a":
                                var v59 uint32 = 560
                                switch (fi7 * fi8) {
                                case 2147483647:
                                    v6 -= (v6 * float64(3.14159))
                                case 128:
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                }
                                v5 = (!true)
                                _ = v59
                            default:
                                switch (fi7 - -1) {
                                case 1:
                                    v60 := v3[-1]
                                    _ = v60
                                case -9223372036854775808:
                                    switch "hello" {
                                    case "a":
                                        _ = len(v1)
                                    case "foo bar":
                                        var v61 uint = 759
                                        _ = v61
                                    }
                                case 4294967295:
                                    var v62 int8 = v2
                                    _ = v62
                                }
                            }
                            v6 += (v6 + float64(float32(v6)))
                            v63 := GenS1{F0: 1}
                            _ = v63
                        }
                        fmt.Println(fi8, v2)
                    }
                    _ = fi8
                    _ = v10
                    _ = v9
                }
            case -227:
                v4 = append(v4, 230)
            default:
                var v64 int16 = 0
                _ = v64
            }
        }
        for fi65 := 0; fi65 < 3; fi65++ {
            v5 = (true || (!v5))
            var v66 int = (fi7 ^ 1)
            delete(v3, fi65)
            v3[fi65] = 42
            switch (fi7 >> ((4294967295) & 63)) {
            case 1:
                if v5 {
                    var v67 int = (fi65 * (fi7 >> ((0) & 63)))
                    var v68 int16 = -384
                    if (false || v5) {
                        v67 = v66
                        v4 = append(v4, 2147483647)
                    }
                    _ = v67
                    _ = v68
                } else {
                    v5 = (false && v5)
                    var v69 int32 = 356
                    fmt.Println(v2, v2, v2)
                    v2 += v2
                    _ = v69
                }
            case 32768:
                fmt.Println(v2)
            }
            _ = fi65
            _ = v66
        }
        _ = fi7
    }
    if true {
        fmt.Println(v6)
        v70 := GenS2{F0: 1, F1: 42, F2: 280}
        _ = v70
    }
    for fi71 := 0; fi71 < 14; fi71++ {
        for fi72 := 0; fi72 < 6; fi72++ {
            var v73 int16 = 1
            _ = fi72
            _ = v73
        }
        v6 = (v6 + 0.0)
        if v5 {
            for fi74 := 0; fi74 < 18; fi74++ {
                var v75 float64 = v6
                if (-2.25 < float32((3.14159 + -2.25))) {
                    var v76 string = ""
                    switch (v76 + ("\n\t" + v76)) {
                    case "a":
                        v1 = append(v1, 2)
                    }
                    fmt.Println(v2, fi71, fi74)
                    _ = v76
                } else {
                    var v77 int32 = 0
                    switch fi74 {
                    case -925:
                        if v5 {
                            v77 -= v77
                            switch "hello" {
                            case "a":
                                if ((v6 >= 1.5) || true) {
                                    if len(v4) > 0 {
                                        v4[0] = 2147483647
                                    }
                                    var v78 int8 = (^(v2))
                                    switch fi71 {
                                    case 9223372036854775807:
                                        if v5 {
                                            switch ("foo bar" + "hello") {
                                            case "hello":
                                                var v79 string = "foo bar"
                                                _ = v79
                                            }
                                        } else {
                                            if (32767 >= 32767) {
                                                v80 := GenS2{F0: (v77 - v77), F1: 1, F2: -1}
                                                fmt.Println(v78)
                                                fmt.Println(fi74)
                                                _ = v80
                                            }
                                        }
                                    }
                                    switch "hello" {
                                    case "\n\t":
                                        if len(v1) > 0 {
                                            v1[0] = 2147483647
                                        }
                                        v77 -= -1
                                        var v81 string = ""
                                        var v82 int = (fi71 << ((65535) & 63))
                                        _ = v81
                                        _ = v82
                                    case "a":
                                        v6 = v75
                                    default:
                                        var v83 int8 = (v2 + v2)
                                        _ = v83
                                    }
                                    v84 := GenS6{F0: 128}
                                    _ = v78
                                    _ = v84
                                } else {
                                    if v5 {
                                        if (((fi71 >> ((255) & 63)) <= -2147483648) || true) {
                                            v85 := v3[fi74]
                                            _ = v85
                                        }
                                        var v86 int8 = (v2 & v2)
                                        v87 := GenS7{F0: 426, F1: 878, F2: (!v5)}
                                        _ = v86
                                        _ = v87
                                    } else {
                                        v3[fi74] = 0
                                        var v88 float32 = 1.5
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        if true {
                                            _ = len(v3)
                                        }
                                        _ = v88
                                    }
                                }
                            case "":
                                fmt.Println(v5, v5)
                            }
                        }
                    case 32767:
                        _ = len(v3)
                        var v89 uint64 = 4294967295
                        _ = v89
                    case -9223372036854775808:
                        switch (("a" + "\n\t") + "\n\t") {
                        case "a":
                            if true {
                                var v90 uint16 = 700
                                v5 = v5
                                if (!(v5 && (v77 >= (^(v77))))) {
                                    switch (fi74 & (-(fi71))) {
                                    case -1:
                                        if true {
                                            fmt.Println(v6, v2, v75)
                                            if true {
                                                v2 = v2
                                                v4 = append(v4, 2)
                                                var v91 int = (fi71 % ((fi71) | 1))
                                                _ = v91
                                            } else {
                                                var v92 int32 = (v77 >> ((255) & 63))
                                                fmt.Println(v90, v90)
                                                _ = v92
                                            }
                                        }
                                    default:
                                        var v93 int16 = -32768
                                        _ = v93
                                    }
                                } else {
                                    v6 -= v75
                                    switch (fi74 | (fi71 + fi74)) {
                                    case 198:
                                        var v94 uint64 = 1
                                        _ = v94
                                    case 32767:
                                        var v95 float64 = v75
                                        _ = v95
                                    }
                                }
                                if true {
                                    var v96 uint32 = 4294967295
                                    if v5 {
                                        var v97 bool = v5
                                        _ = v97
                                    } else {
                                        var v98 int = fi74
                                        _ = v98
                                    }
                                    fmt.Println(fi74, v90, v2)
                                    _ = v96
                                } else {
                                    v5 = (765 != fi74)
                                }
                                var v99 uint8 = 202
                                _ = v90
                                _ = v99
                            } else {
                                v2 = -1
                            }
                            v77 = v77
                            v100 := GenS8{F0: 42}
                            _ = v100
                        case "":
                            v2 = 32
                        default:
                            v75 = ((v6 * v75) / 1.0)
                        }
                    default:
                        if (-309 < 1) {
                            v101 := GenS9{F0: 0, F1: 0}
                            _ = len(v1)
                            fmt.Println(v2)
                            _ = v101
                        }
                    }
                    if true {
                        v75 = (v6 * v75)
                        switch "\n\t" {
                        case "":
                            var v102 string = ("" + "foo bar")
                            _ = v102
                        default:
                            v2 += v2
                        }
                        fmt.Println(v75)
                        v5 = v5
                        v103 := GenS4{F0: 68, F1: -2147483648, F2: 0}
                        v77 -= v77
                        switch fi74 {
                        case -900:
                            var v104 uint64 = 1
                            _ = v104
                        case 32768:
                            v2 -= v2
                        case 42:
                            delete(v3, fi74)
                        default:
                            fmt.Println(fi74)
                        }
                        _ = v103
                    } else {
                        v105 := GenS7{F0: 4294967295, F1: 0, F2: (!(1000 <= 32767))}
                        _ = v105
                    }
                    var v106 bool = true
                    _ = v106
                    _ = v77
                }
                _ = fi74
                _ = v75
            }
            for fi107 := 0; fi107 < 4; fi107++ {
                switch "" {
                case "foo bar":
                    v108 := GenS9{F0: fi107, F1: 128}
                    _ = v108
                case "a":
                    v6 = v6
                }
                if (!(v5 || v5)) {
                    if (!v5) {
                        delete(v3, -952)
                    } else {
                        if v5 {
                            switch "" {
                            case "\n\t":
                                var v109 uint8 = 54
                                _ = v109
                            case "":
                                if v5 {
                                    var v110 uint = 830
                                    if len(v4) > 0 {
                                        v4[0] = 994
                                    }
                                    _ = v110
                                }
                            default:
                                if false {
                                    switch -126 {
                                    case 128:
                                        v2 -= 127
                                        fmt.Println(fi107, v5)
                                    default:
                                        if len(v1) > 0 {
                                            v1[0] = 360
                                        }
                                    }
                                    switch (-(fi107)) {
                                    case 128:
                                        var v111 uint = 211
                                        _ = v111
                                    }
                                } else {
                                    v6 = 3.14159
                                    fmt.Println(v2, fi107, fi107)
                                    fmt.Println(v2)
                                    v6 += v6
                                }
                            }
                            var v112 uint64 = 2
                            _ = len(v4)
                            var v113 bool = ((!true) || true)
                            _ = v112
                            _ = v113
                        }
                        v114 := v3[810]
                        var v115 float64 = ((v6 * 3.14159) * v6)
                        _ = v114
                        _ = v115
                    }
                    var v116 int8 = v2
                    if len(v1) > 0 {
                        v1[0] = 0
                    }
                    _ = v116
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    if false {
                        var v117 uint8 = 251
                        v6 -= float64(float32((v6 - 0.1)))
                        fmt.Println(v117, fi107)
                        switch ("a" + (("hello" + "hello") + (("hello" + "foo bar") + ("\n\t" + "")))) {
                        case "a":
                            v118 := GenS6{F0: v117}
                            _ = v118
                        case "hello":
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        _ = v117
                    } else {
                        var v119 float32 = ((0.1 * 0.0) / (0.1 - (1e10 / 0.1)))
                        if (2 == 183) {
                            v120 := GenS10{F0: 1}
                            v3[fi107] = 42
                            switch "foo bar" {
                            case "\n\t":
                                if (!true) {
                                    v6 += v6
                                    v5 = true
                                }
                            default:
                                _ = len(v3)
                            }
                            _ = v120
                        } else {
                            fmt.Println(fi71)
                        }
                        _ = v119
                    }
                    if (v5 && (v5 || (!v5))) {
                        switch (("a" + "a") + "a") {
                        case "hello":
                            v6 = (v6 * (v6 / (-2.25 / (v6 * v6))))
                        case "foo bar":
                            var v121 uint32 = 2
                            _ = v121
                        default:
                            var v122 uint32 = 917
                            _ = v122
                        }
                        v123 := GenS5{F0: (-(fi71)), F1: fi71}
                        fmt.Println(v2)
                        _ = v123
                    } else {
                        var v124 uint16 = 32767
                        var v125 uint8 = 154
                        _ = v124
                        _ = v125
                    }
                }
                _ = fi107
            }
            var v126 string = "hello"
            _ = v126
        } else {
            for fi127 := 0; fi127 < 10; fi127++ {
                var v128 uint = 2
                delete(v3, fi71)
                _ = fi127
                _ = v128
            }
        }
        _ = fi71
    }
    delete(v3, 32767)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi129 := 0; fi129 < 12; fi129++ {
        v130 := GenS11{F0: -1}
        var v131 int8 = 125
        _ = fi129
        _ = v130
        _ = v131
    }
    switch 659 {
    case 32768:
        for fi132 := 0; fi132 < 14; fi132++ {
            v5 = v5
            switch (("foo bar" + "foo bar") + "") {
            case "":
                v133 := v3[(^(fi132))]
                _ = v133
            case "foo bar":
                v2 = (v2 >> ((4294967295) & 63))
            default:
                v5 = true
            }
            _ = fi132
        }
    case 32767:
        var v134 int16 = 1
        _ = v134
    case 4294967295:
        v135 := GenS8{F0: 445}
        _ = v135
    default:
        v136 := GenS12{F0: (v2 + v2)}
        _ = v136
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
