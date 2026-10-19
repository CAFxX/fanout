package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 uint8; F1 uint16; F2 int64 }
type GenS4 struct { F0 uint64; F1 int32 }
type GenS5 struct { F0 uint32; F1 uint }
type GenS6 struct { F0 int64; F1 int }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int64; F1 int }
type GenS9 struct { F0 bool; F1 int16; F2 uint }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 bool; F1 uint8 }
type GenS12 struct { F0 string; F1 int16 }
type GenS13 struct { F0 uint8; F1 int32; F2 int }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 int16; F1 int8 }
type GenS16 struct { F0 bool }
type GenS17 struct { F0 bool; F1 uint8; F2 uint }
type GenS18 struct { F0 uint64; F1 int16; F2 int32 }
type GenS19 struct { F0 uint; F1 int16; F2 int32 }
func genFn0(p0 uint32) int64 {
    v5 := GenS1{F0: 0, F1: 0}
    for fi6 := 0; fi6 < 16; fi6++ {
        var v7 int8 = 105
        var v8 int = 127
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 17; fi9++ {
        v10 := []int{1}
        for fi11 := 0; fi11 < 10; fi11++ {
            switch "a" {
            case "\n\t":
                var v12 float32 = 1.5
                _ = v12
            default:
                switch (("a" + "\n\t") + "foo bar") {
                case "\n\t":
                    if true {
                        p0 -= p0
                        v10 = append(v10, fi9)
                    }
                default:
                    var v13 uint32 = 2147483647
                    _ = v13
                }
            }
            p0 += (p0 ^ p0)
            if false {
                if true {
                    var v14 uint = 65535
                    v15 := GenS1{F0: 42, F1: -1}
                    switch -1 {
                    case 32767:
                        var v16 float64 = 0.0
                        _ = v16
                    case -206:
                        v14 -= (v14 + 255)
                    }
                    _ = v14
                    _ = v15
                } else {
                    p0 += (^(p0))
                }
                if true {
                    if false {
                        var v17 bool = (-70 == -128)
                        _ = v17
                    }
                    var v18 float64 = (float64(((1e10 * ((0.1 - (float32(-2.25) + 1.5)) * 1e10)) + 1.5)) * (1e10 / ((float64(0.1) * (0.0 - 1e10)) + 0.1)))
                    if false {
                        var v19 float32 = 0.1
                        _ = v19
                    } else {
                        v18 = 1.5
                    }
                    v20 := GenS2{F0: -9223372036854775808, F1: 1000}
                    v18 = float64(float32(0.0))
                    if len(v10) > 0 {
                        v10[0] = (fi11 & fi11)
                    }
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                    switch (fi9 & (fi9 * (^(fi9)))) {
                    case 127:
                        var v21 int = (fi11 * fi11)
                        _ = v21
                    }
                    v10 = append(v10, fi9)
                    var v22 string = "a"
                    _ = v18
                    _ = v20
                    _ = v22
                } else {
                    var v23 float32 = (3.14159 - 3.14159)
                    if true {
                        var v24 string = "a"
                        switch (("\n\t" + v24) + "") {
                        case "\n\t":
                            switch fi11 {
                            case 128:
                                fmt.Println(v24, v24)
                            }
                        case "hello":
                            v23 -= v23
                        default:
                            if false {
                                switch (("foo bar" + (v24 + v24)) + "\n\t") {
                                case "foo bar":
                                    if true {
                                        switch "foo bar" {
                                        case "a":
                                            v25 := GenS1{F0: 0, F1: -1}
                                            _ = v25
                                        case "":
                                            if false {
                                                fmt.Println(v23)
                                            }
                                        case "\n\t":
                                            v10 = append(v10, fi11)
                                            if true {
                                                fmt.Println(fi9)
                                                if true {
                                                    v26 := GenS2{F0: (^(fi9)), F1: 9}
                                                    v10 = append(v10, (fi9 / (((fi11 & (fi9 & fi9))) | 1)))
                                                    _ = v26
                                                }
                                                v27 := GenS3{F0: 255, F1: 2, F2: -1}
                                                _ = v27
                                            }
                                            switch fi11 {
                                            case 1:
                                                var v28 int8 = 0
                                                var v29 uint = 65535
                                                var v30 uint8 = 128
                                                _ = v28
                                                _ = v29
                                                _ = v30
                                            case 279:
                                                fmt.Println(fi11, fi9, v23)
                                            case -385:
                                                v23 -= v23
                                            }
                                            p0 = (^(p0))
                                        }
                                        fmt.Println(v24)
                                        if (42 != 42) {
                                            v31 := GenS4{F0: 1, F1: -840}
                                            _ = v31
                                        } else {
                                            switch v24 {
                                            case "foo bar":
                                                v32 := GenS3{F0: 255, F1: 747, F2: 9223372036854775807}
                                                _ = v32
                                            case "":
                                                if (255 != 127) {
                                                    if (true && ((p0 + (p0 * p0)) < 2)) {
                                                        p0 = p0
                                                        v33 := GenS3{F0: 42, F1: 255, F2: -1}
                                                        var v34 uint8 = 128
                                                        _ = v33
                                                        _ = v34
                                                    } else {
                                                        p0 = p0
                                                    }
                                                    fmt.Println(v24, v24)
                                                    var v35 uint = 4294967295
                                                    _ = v35
                                                } else {
                                                    var v36 float64 = (0.0 - (0.0 * ((1e10 / float64(v23)) / (0.1 * 0.1))))
                                                    v37 := GenS5{F0: p0, F1: 132}
                                                    _ = v36
                                                    _ = v37
                                                }
                                                var v38 float32 = ((-2.25 * ((v23 / float32(((-2.25 + 1e10) + 1e10))) + v23)) / 1.5)
                                                fmt.Println(v24)
                                                _ = v38
                                            default:
                                                var v39 uint32 = 385
                                                _ = v39
                                            }
                                            if false {
                                                switch fi9 {
                                                case 631:
                                                    v23 = v23
                                                case 201:
                                                    var v40 int16 = 0
                                                    _ = v40
                                                case 32768:
                                                    p0 -= 220
                                                default:
                                                    if false {
                                                        if true {
                                                            if false {
                                                                v41 := GenS6{F0: -1, F1: fi11}
                                                                if true {
                                                                    var v42 uint16 = 517
                                                                    if true {
                                                                        var v43 float64 = (float64((1e10 - 0.0)) / 1.5)
                                                                        _ = v43
                                                                    }
                                                                    _ = v42
                                                                }
                                                                _ = v41
                                                            } else {
                                                                switch v24 {
                                                                case "a":
                                                                    switch 4294967295 {
                                                                    case -64:
                                                                        if false {
                                                                            var v44 uint8 = 2
                                                                            v44 = v44
                                                                            _ = v44
                                                                        }
                                                                    default:
                                                                        if true {
                                                                            fmt.Println(fi9, v24)
                                                                            fmt.Println(fi11, v24)
                                                                            switch fi11 {
                                                                            case -368:
                                                                                v45 := GenS3{F0: 127, F1: 967, F2: -1}
                                                                                _ = v45
                                                                            case 128:
                                                                                var v46 int64 = 0
                                                                                _ = v46
                                                                            }
                                                                        }
                                                                    }
                                                                case "":
                                                                    v24 = ("foo bar" + v24)
                                                                }
                                                                _ = len(v10)
                                                            }
                                                            var v47 uint32 = (p0 >> ((65535) & 63))
                                                            fmt.Println(fi9, p0)
                                                            if len(v10) >= 2 {
                                                                v10 = v10[1:len(v10)]
                                                            }
                                                            v47 += v47
                                                            var v48 int = fi11
                                                            _ = v47
                                                            _ = v48
                                                        }
                                                        var v49 uint32 = p0
                                                        _ = v49
                                                    } else {
                                                        if (40 == 127) {
                                                            p0 = 42
                                                            if (-1 <= -391) {
                                                                switch v24 {
                                                                case "foo bar":
                                                                    var v50 uint32 = p0
                                                                    _ = v50
                                                                case "a":
                                                                    var v51 int64 = 0
                                                                    _ = v51
                                                                case "\n\t":
                                                                    if (float32(0.0) == (1e10 + 1.5)) {
                                                                        switch "" {
                                                                        case "a":
                                                                            v23 = v23
                                                                            if (v24 >= "a") {
                                                                                if len(v10) > 0 {
                                                                                    v10[0] = 4294967295
                                                                                }
                                                                                v24 = v24
                                                                                switch "foo bar" {
                                                                                case "hello":
                                                                                    var v52 int = 990
                                                                                    _ = v52
                                                                                }
                                                                            } else {
                                                                                var v53 string = "\n\t"
                                                                                v10 = append(v10, (-(fi9)))
                                                                                _ = v53
                                                                            }
                                                                            if len(v10) >= 2 {
                                                                                v10 = v10[1:len(v10)]
                                                                            }
                                                                        case "hello":
                                                                            var v54 uint = 0
                                                                            _ = v54
                                                                        case "foo bar":
                                                                            var v55 uint16 = 151
                                                                            _ = v55
                                                                        }
                                                                        v23 -= float32((float64(1.5) * 1e10))
                                                                    }
                                                                default:
                                                                    v56 := GenS7{F0: (-12 != 0)}
                                                                    _ = v56
                                                                }
                                                            } else {
                                                                if true {
                                                                    fmt.Println(fi9, fi11, v23)
                                                                }
                                                            }
                                                            fmt.Println(fi11)
                                                            p0 -= 42
                                                        }
                                                        v57 := GenS4{F0: 331, F1: 1}
                                                        var v58 int8 = -47
                                                        var v59 int8 = (-(v58))
                                                        if len(v10) >= 2 {
                                                            v10 = v10[1:len(v10)]
                                                        }
                                                        var v60 bool = true
                                                        _ = v57
                                                        _ = v58
                                                        _ = v59
                                                        _ = v60
                                                    }
                                                }
                                                v23 = 1e10
                                                if (0 >= 42) {
                                                    if false {
                                                        var v61 int8 = 42
                                                        var v62 uint32 = p0
                                                        _ = v61
                                                        _ = v62
                                                    }
                                                } else {
                                                    if false {
                                                        fmt.Println(v24, p0)
                                                    } else {
                                                        fmt.Println(p0, p0, v23)
                                                        v63 := GenS5{F0: 2, F1: 200}
                                                        var v64 uint32 = (p0 * p0)
                                                        v64 += (p0 ^ p0)
                                                        _ = len(v10)
                                                        v24 = v24
                                                        fmt.Println(v24)
                                                        _ = len(v10)
                                                        if true {
                                                            var v65 int = 32767
                                                            switch v65 {
                                                            case 9223372036854775807:
                                                                var v66 bool = (v23 != v23)
                                                                v64 -= v64
                                                                _ = v66
                                                            case 4294967295:
                                                                if false {
                                                                    if ((!((v64 + p0) == p0)) || false) {
                                                                        v67 := GenS8{F0: -1, F1: v65}
                                                                        if true {
                                                                            var v68 float64 = (-2.25 + 3.14159)
                                                                            var v69 int64 = 1
                                                                            v64 = v64
                                                                            v69 = (^(v69))
                                                                            _ = v68
                                                                            _ = v69
                                                                        } else {
                                                                            p0 = p0
                                                                            v23 = v23
                                                                            if true {
                                                                                var v70 uint32 = (v64 - p0)
                                                                                v70 += p0
                                                                                if len(v10) >= 2 {
                                                                                    v10 = v10[1:len(v10)]
                                                                                }
                                                                                _ = v70
                                                                            }
                                                                        }
                                                                        _ = len(v10)
                                                                        _ = v67
                                                                    }
                                                                    _ = len(v10)
                                                                } else {
                                                                    v64 += v64
                                                                }
                                                            case 713:
                                                                var v71 int32 = -958
                                                                _ = v71
                                                            default:
                                                                switch v65 {
                                                                case -9223372036854775808:
                                                                    v72 := GenS3{F0: 128, F1: 65535, F2: 42}
                                                                    _ = v72
                                                                case 2147483647:
                                                                    if len(v10) > 0 {
                                                                        v10[0] = 818
                                                                    }
                                                                default:
                                                                    v24 = ""
                                                                }
                                                            }
                                                            var v73 int8 = 1
                                                            _ = v65
                                                            _ = v73
                                                        }
                                                        _ = v63
                                                        _ = v64
                                                    }
                                                    if len(v10) >= 2 {
                                                        v10 = v10[1:len(v10)]
                                                    }
                                                    var v74 uint16 = 365
                                                    _ = v74
                                                }
                                            } else {
                                                p0 += (p0 % ((p0) | 1))
                                            }
                                            var v75 bool = ((3.14159 / 1.5) <= 0.1)
                                            _ = v75
                                        }
                                    }
                                case "a":
                                    if len(v10) >= 2 {
                                        v10 = v10[1:len(v10)]
                                    }
                                case "hello":
                                    if (!(!true)) {
                                        p0 = 1
                                        var v76 uint = 659
                                        fmt.Println(fi9, v23)
                                        var v77 bool = true
                                        switch "foo bar" {
                                        case "":
                                            if (3.14159 < ((-2.25 - 1.5) / 1.0)) {
                                                v78 := GenS9{F0: v77, F1: -848, F2: v76}
                                                switch "hello" {
                                                case "\n\t":
                                                    v24 = v24
                                                case "":
                                                    v77 = v77
                                                case "foo bar":
                                                    if true {
                                                        _ = len(v10)
                                                        if (1 <= 65535) {
                                                            if len(v10) >= 2 {
                                                                v10 = v10[1:len(v10)]
                                                            }
                                                            var v79 uint16 = 880
                                                            var v80 int64 = -1
                                                            _ = v79
                                                            _ = v80
                                                        } else {
                                                            if (float64(0.1) != float64(v23)) {
                                                                var v81 uint32 = 2
                                                                _ = v81
                                                            } else {
                                                                v24 = v24
                                                            }
                                                        }
                                                        switch (fi11 << ((v76) & 63)) {
                                                        case 9223372036854775807:
                                                            var v82 bool = ((236 >= (^(p0))) && ((p0 / ((4294967295) | 1)) >= p0))
                                                            if len(v10) >= 2 {
                                                                v10 = v10[1:len(v10)]
                                                            }
                                                            _ = v82
                                                        case -522:
                                                            if ((!(0 == 0)) && false) {
                                                                v10 = append(v10, fi11)
                                                                if (v77 && (((v76 >> (((v76 << ((v76) & 63))) & 63)) < (^(v76))) || v77)) {
                                                                    var v83 uint32 = p0
                                                                    _ = v83
                                                                } else {
                                                                    switch v24 {
                                                                    case "a":
                                                                        var v84 string = ((v24 + v24) + (v24 + ((v24 + v24) + (("hello" + v24) + v24))))
                                                                        _ = v84
                                                                    }
                                                                    var v85 uint8 = 111
                                                                    _ = v85
                                                                }
                                                                var v86 uint = (v76 + (v76 / ((255) | 1)))
                                                                _ = v86
                                                            } else {
                                                                var v87 int32 = 355
                                                                var v88 int = (fi9 | 1)
                                                                var v89 bool = true
                                                                _ = v87
                                                                _ = v88
                                                                _ = v89
                                                            }
                                                        }
                                                    } else {
                                                        if (1 >= 896) {
                                                            _ = len(v10)
                                                            _ = len(v10)
                                                        }
                                                        v90 := GenS6{F0: -846, F1: fi11}
                                                        var v91 int8 = 66
                                                        if false {
                                                            if len(v10) > 0 {
                                                                v10[0] = fi9
                                                            }
                                                            var v92 int = fi9
                                                            _ = v92
                                                        } else {
                                                            if len(v10) > 0 {
                                                                v10[0] = (fi9 + (fi9 * (fi11 - -1)))
                                                            }
                                                            p0 = p0
                                                            var v93 uint = (v76 * (v76 * v76))
                                                            _ = v93
                                                        }
                                                        var v94 int8 = (v91 ^ v91)
                                                        _ = v90
                                                        _ = v91
                                                        _ = v94
                                                    }
                                                default:
                                                    if len(v10) >= 2 {
                                                        v10 = v10[1:len(v10)]
                                                    }
                                                }
                                                fmt.Println(v24, v24)
                                                _ = v78
                                            }
                                        }
                                        v95 := GenS6{F0: 9223372036854775807, F1: fi11}
                                        fmt.Println(v24)
                                        if (v77 && (v77 || (0.1 >= 3.14159))) {
                                            v96 := GenS10{F0: 42}
                                            _ = v96
                                        } else {
                                            if len(v10) > 0 {
                                                v10[0] = 142
                                            }
                                            fmt.Println(v23)
                                        }
                                        _ = v76
                                        _ = v77
                                        _ = v95
                                    } else {
                                        switch (((v24 + v24) + ("a" + "foo bar")) + "\n\t") {
                                        case "":
                                            p0 -= p0
                                        }
                                        if (2147483647 < 2147483647) {
                                            v97 := GenS1{F0: 42, F1: 42}
                                            _ = v97
                                        } else {
                                            v98 := GenS11{F0: true, F1: 94}
                                            var v99 uint32 = (p0 ^ (p0 << ((1) & 63)))
                                            v24 = ((v24 + ((("a" + ((v24 + (((v24 + v24) + (v24 + v24)) + (v24 + ("a" + "hello")))) + "hello")) + "\n\t") + "a")) + v24)
                                            var v100 uint = 111
                                            var v101 bool = (-1 < 930)
                                            _ = v100
                                            _ = v101
                                            _ = v98
                                            _ = v99
                                        }
                                        var v102 int32 = -2147483648
                                        _ = v102
                                    }
                                }
                                v24 = (v24 + v24)
                                if false {
                                    var v103 float64 = (0.0 - -2.25)
                                    _ = v103
                                }
                            } else {
                                if (false || true) {
                                    if true {
                                        var v104 uint8 = 42
                                        switch (fi11 + 4294967295) {
                                        case 728:
                                            var v105 int = -563
                                            _ = v105
                                        case 128:
                                            v106 := GenS12{F0: v24, F1: 806}
                                            _ = v106
                                        default:
                                            v23 = float32(-2.25)
                                        }
                                        v23 = float32((0.0 / (1.5 + 1.5)))
                                        if ((fi11 != fi11) || true) {
                                            fmt.Println(v23)
                                            fmt.Println(fi9, fi11)
                                        }
                                        _ = v104
                                    }
                                }
                                v24 = v24
                                v23 = v23
                            }
                        }
                        v23 = v23
                        _ = v24
                    } else {
                        v107 := GenS11{F0: false, F1: 42}
                        var v108 int8 = 42
                        if len(v10) >= 2 {
                            v10 = v10[1:len(v10)]
                        }
                        _ = v107
                        _ = v108
                    }
                    _ = v23
                }
                if len(v10) > 0 {
                    v10[0] = (-(fi11))
                }
            } else {
                switch fi9 {
                case 1:
                    var v109 uint64 = 0
                    _ = v109
                case -103:
                    if (true || (42 == 32767)) {
                        p0 -= p0
                        p0 = 4294967295
                        switch (fi9 * (fi9 >> ((1) & 63))) {
                        case 787:
                            p0 = p0
                        case 298:
                            p0 = (p0 | p0)
                        case -440:
                            p0 += 1
                        }
                    } else {
                        fmt.Println(fi9)
                        switch fi11 {
                        case 9223372036854775807:
                            fmt.Println(fi9)
                        case 968:
                            var v110 string = "\n\t"
                            _ = v110
                        case 4294967295:
                            if false {
                                v111 := GenS3{F0: 2, F1: 1, F2: 9223372036854775807}
                                var v112 uint8 = 42
                                if (!(-1 >= 127)) {
                                    var v113 int8 = 100
                                    if len(v10) > 0 {
                                        v10[0] = (fi9 >> ((65535) & 63))
                                    }
                                    switch (fi11 << ((42) & 63)) {
                                    case -560:
                                        var v114 int16 = -912
                                        _ = v114
                                    case 32767:
                                        var v115 float64 = 3.14159
                                        _ = v115
                                    case -2147483648:
                                        p0 -= (p0 + 2147483648)
                                    }
                                    _ = v113
                                }
                                _ = v111
                                _ = v112
                            }
                        default:
                            p0 -= (p0 << ((0) & 63))
                        }
                        var v116 uint8 = 32
                        _ = v116
                    }
                default:
                    p0 = 2147483647
                }
                v10 = append(v10, (fi11 ^ (fi11 >> ((2) & 63))))
                v10 = append(v10, (^(fi9)))
            }
            var v117 uint16 = 2
            var v118 uint8 = 71
            _ = fi11
            _ = v117
            _ = v118
        }
        var v119 int16 = 0
        _ = fi9
        _ = v10
        _ = v119
    }
    p0 = p0
    _ = p0
    _ = v5
    return -9223372036854775808
}
func genFn1() string {
    switch "\n\t" {
    case "":
        for fi120 := 0; fi120 < 9; fi120++ {
            v121 := GenS3{F0: 231, F1: 2, F2: -9223372036854775808}
            _ = fi120
            _ = v121
        }
    case "hello":
        var v122 float64 = -2.25
        _ = v122
    case "\n\t":
        switch 42 {
        case 32768:
            var v123 uint16 = 921
            _ = v123
        case 0:
            v124 := []int{}
            _ = v124
        }
    default:
        var v125 float64 = (3.14159 / 1.5)
        _ = v125
    }
    for fi126 := 0; fi126 < 18; fi126++ {
        v127 := GenS13{F0: 226, F1: int32(fi126), F2: fi126}
        var v128 float32 = -2.25
        for fi129 := 0; fi129 < 3; fi129++ {
            fmt.Println(fi129)
            v130 := GenS14{F0: 2}
            v128 = (((v128 * v128) + (v128 * v128)) / v128)
            fmt.Println(fi129, fi129, v128)
            _ = fi129
            _ = v130
        }
        _ = fi126
        _ = v127
        _ = v128
    }
    switch "a" {
    case "\n\t":
        switch 32768 {
        case -134:
            switch -743 {
            case 9223372036854775807:
                switch "a" {
                case "\n\t":
                    var v131 float64 = -2.25
                    _ = v131
                case "hello":
                    switch -414 {
                    case 9223372036854775807:
                        var v132 uint8 = 128
                        _ = v132
                    case 0:
                        var v133 uint64 = 2
                        _ = v133
                    case 42:
                        v134 := GenS15{F0: 32767, F1: 109}
                        _ = v134
                    }
                case "a":
                    if true {
                        var v135 uint16 = 1
                        switch "hello" {
                        case "\n\t":
                            if true {
                                v135 -= (v135 % (((v135 * v135)) | 1))
                            } else {
                                var v136 int64 = -62
                                v136 = 0
                                for fi137 := 0; fi137 < 13; fi137++ {
                                    if (!(17 <= 1000)) {
                                        v138 := GenS13{F0: uint8(v135), F1: 2147483647, F2: fi137}
                                        if (-2147483648 < 1) {
                                            switch ("\n\t" + "hello") {
                                            case "\n\t":
                                                for fi139 := 0; fi139 < 6; fi139++ {
                                                    var v140 int64 = (-(v136))
                                                    var v141 float32 = 0.1
                                                    v142 := []int{-1, (fi139 / ((1) | 1))}
                                                    _ = fi139
                                                    _ = v140
                                                    _ = v141
                                                    _ = v142
                                                }
                                            }
                                            for fi143 := 0; fi143 < 8; fi143++ {
                                                v135 += (v135 | v135)
                                                v136 -= v136
                                                var v144 float64 = 1.5
                                                if true {
                                                    fmt.Println(v136, fi143, v136)
                                                    var v145 int8 = -1
                                                    _ = v145
                                                } else {
                                                    if true {
                                                        var v146 int = (fi143 ^ fi143)
                                                        v147 := GenS11{F0: true, F1: 8}
                                                        _ = v146
                                                        _ = v147
                                                    } else {
                                                        v148 := GenS14{F0: v135}
                                                        v149 := []int{}
                                                        _ = v148
                                                        _ = v149
                                                    }
                                                    v135 -= v135
                                                    v150 := GenS6{F0: (v136 % ((v136) | 1)), F1: (-(fi137))}
                                                    _ = v150
                                                }
                                                _ = fi143
                                                _ = v144
                                            }
                                        }
                                        _ = v138
                                    }
                                    for fi151 := 0; fi151 < 16; fi151++ {
                                        if (-808 > -554) {
                                            var v152 int16 = 0
                                            v153 := GenS16{F0: (int32(fi151) != -1)}
                                            _ = v152
                                            _ = v153
                                        }
                                        _ = fi151
                                    }
                                    _ = fi137
                                }
                                _ = v136
                            }
                        case "":
                            for fi154 := 0; fi154 < 3; fi154++ {
                                var v155 int32 = -1
                                if (!true) {
                                    for fi156 := 0; fi156 < 1; fi156++ {
                                        v157 := []int{fi156, fi154, fi154, fi156}
                                        _ = fi156
                                        _ = v157
                                    }
                                }
                                _ = fi154
                                _ = v155
                            }
                        case "foo bar":
                            switch -9223372036854775808 {
                            case -2147483648:
                                for fi158 := 0; fi158 < 6; fi158++ {
                                    v159 := GenS10{F0: int8(v135)}
                                    for fi160 := 0; fi160 < 9; fi160++ {
                                        v135 = 943
                                        v161 := GenS17{F0: (-2.25 > 1e10), F1: 1, F2: 1}
                                        if (!true) {
                                            v135 -= (v135 / ((602) | 1))
                                            switch (fi158 * fi158) {
                                            case -258:
                                                var v162 uint = 1
                                                _ = v162
                                            case -2147483648:
                                                if false {
                                                    v135 = v135
                                                    var v163 uint8 = 127
                                                    var v164 float32 = ((1.5 + (1.5 - ((-2.25 * 0.0) + -2.25))) - 0.0)
                                                    _ = v163
                                                    _ = v164
                                                } else {
                                                    v165 := []int{(fi158 << ((65535) & 63)), fi158, fi160, fi160}
                                                    _ = v165
                                                }
                                            default:
                                                switch "" {
                                                case "foo bar":
                                                    v166 := []int{4294967295}
                                                    _ = v166
                                                case "\n\t":
                                                    if (float32(0.1) < -2.25) {
                                                        switch (fi160 - (fi160 | (fi160 >> ((186) & 63)))) {
                                                        case 626:
                                                            var v167 uint64 = 18446744073709551615
                                                            v168 := GenS15{F0: -1, F1: -82}
                                                            if (((1.5 + 1e10) <= 1.5) || (((-9223372036854775808 > 42) || (((fi158 & fi160) >= -125) && true)) || ((-2.25 <= (3.14159 + 3.14159)) && (!(396 != genFn0(1)))))) {
                                                                v169 := GenS11{F0: false, F1: 126}
                                                                var v170 string = ("\n\t" + "\n\t")
                                                                v170 = "a"
                                                                _ = v169
                                                                _ = v170
                                                            } else {
                                                                switch (fi160 << ((65535) & 63)) {
                                                                case 42:
                                                                    var v171 int = (-(fi160))
                                                                    _ = v171
                                                                }
                                                            }
                                                            _ = v167
                                                            _ = v168
                                                        }
                                                        v135 -= (^(v135))
                                                        var v172 uint64 = 42
                                                        _ = v172
                                                    } else {
                                                        v173 := GenS18{F0: 42, F1: -1, F2: 1}
                                                        _ = v173
                                                    }
                                                }
                                            }
                                        } else {
                                            var v174 float64 = (0.1 / ((1e10 * -2.25) / -2.25))
                                            if ((fi158 & fi158) == fi160) {
                                                if true {
                                                    var v175 uint16 = 42
                                                    _ = v175
                                                }
                                            } else {
                                                fmt.Println(fi158)
                                                v174 = (((0.0 * (1.5 * v174)) / v174) / v174)
                                            }
                                            _ = v174
                                        }
                                        if false {
                                            var v176 uint8 = 42
                                            var v177 string = "\n\t"
                                            if false {
                                                var v178 uint16 = 720
                                                if false {
                                                    var v179 int64 = 42
                                                    if false {
                                                        v176 = v176
                                                    } else {
                                                        if true {
                                                            switch v177 {
                                                            case "\n\t":
                                                                v180 := GenS17{F0: true, F1: 2, F2: 65535}
                                                                _ = v180
                                                            case "foo bar":
                                                                fmt.Println(fi158, v176, v135)
                                                            case "":
                                                                fmt.Println(v176)
                                                            default:
                                                                v179 = (-(v179))
                                                            }
                                                        }
                                                    }
                                                    _ = v179
                                                }
                                                if false {
                                                    if true {
                                                        var v181 uint = 65535
                                                        fmt.Println(v135, fi160, v135)
                                                        v182 := GenS6{F0: 581, F1: fi158}
                                                        var v183 string = "hello"
                                                        v135 += v135
                                                        if (false || false) {
                                                            v184 := GenS3{F0: 1, F1: (v135 * 65535), F2: genFn0(2)}
                                                            var v185 uint16 = v178
                                                            var v186 int8 = 100
                                                            _ = v184
                                                            _ = v185
                                                            _ = v186
                                                        } else {
                                                            v177 = "hello"
                                                            v177 = ("\n\t" + "\n\t")
                                                        }
                                                        v135 += v135
                                                        _ = v181
                                                        _ = v182
                                                        _ = v183
                                                    } else {
                                                        if ((false || (0.0 >= (1e10 * ((((0.0 + float32((3.14159 + -2.25))) * float32(3.14159)) / (1.5 / (3.14159 / 1.0))) - 0.0)))) || false) {
                                                            v187 := []int{-558, fi158}
                                                            v176 = (v176 ^ (v176 ^ 240))
                                                            var v188 int64 = 1
                                                            _ = v187
                                                            _ = v188
                                                        }
                                                        var v189 int8 = 100
                                                        v190 := []int{-2147483648, fi158}
                                                        v135 = 42
                                                        _ = v189
                                                        _ = v190
                                                    }
                                                    v135 += (^(v178))
                                                } else {
                                                    var v191 int = 1
                                                    v192 := GenS19{F0: 65535, F1: 0, F2: 0}
                                                    var v193 int = v191
                                                    _ = v191
                                                    _ = v192
                                                    _ = v193
                                                }
                                                switch -9223372036854775808 {
                                                case -2147483648:
                                                    v177 = "a"
                                                case 127:
                                                    var v194 uint = 2
                                                    _ = v194
                                                }
                                                _ = v178
                                            }
                                            var v195 int16 = 0
                                            _ = v176
                                            _ = v177
                                            _ = v195
                                        } else {
                                            v196 := GenS13{F0: 128, F1: 2147483647, F2: (fi160 >> ((4294967295) & 63))}
                                            v135 -= (v135 ^ v135)
                                            _ = v196
                                        }
                                        _ = fi160
                                        _ = v161
                                    }
                                    v135 += v135
                                    _ = fi158
                                    _ = v159
                                }
                            default:
                                fmt.Println(v135)
                            }
                        default:
                            fmt.Println(v135)
                        }
                        v135 = 65535
                        _ = v135
                    }
                default:
                    if (!true) {
                        for fi197 := 0; fi197 < 9; fi197++ {
                            fmt.Println(fi197)
                            var v198 int8 = 100
                            _ = fi197
                            _ = v198
                        }
                    } else {
                        v199 := []int{285, 216}
                        var v200 float64 = (-2.25 / (0.0 + -2.25))
                        _ = v199
                        _ = v200
                    }
                }
            case 785:
                var v201 int64 = 42
                _ = v201
            }
        case 127:
            if true {
                for fi202 := 0; fi202 < 8; fi202++ {
                    switch fi202 {
                    case 127:
                        var v203 uint64 = 18446744073709551615
                        _ = v203
                    }
                    _ = fi202
                }
            }
        default:
            for fi204 := 0; fi204 < 12; fi204++ {
                var v205 int16 = -32768
                _ = fi204
                _ = v205
            }
        }
    case "a":
        var v206 uint8 = 1
        _ = v206
    case "":
        var v207 float32 = 3.14159
        _ = v207
    default:
        if true {
            var v208 int = -1
            v209 := []int{127}
            _ = v208
            _ = v209
        }
    }
    switch (("\n\t" + ("" + "\n\t")) + "") {
    case "\n\t":
        var v210 uint64 = 4294967295
        _ = v210
    case "a":
        var v211 uint16 = 1
        _ = v211
    case "":
        var v212 uint16 = 354
        _ = v212
    }
    if false {
        var v213 uint32 = 2147483647
        v213 = v213
        _ = v213
    }
    return ""
}

func main() {
    var v1 int16 = 1000
    v2 := make(map[uint32]string)
    v3 := make(map[int]bool)
    v4 := make(map[uint32]int)
    switch "" {
    case "\n\t":
        fmt.Println(v1)
    default:
        var v214 int = 2147483647
        _ = v214
    }
    v215 := GenS13{F0: 127, F1: -1, F2: -2147483648}
    v1 += (^(v1))
    v1 -= 0
    v1 += v1
    var v216 float32 = 1.5
    for fi217 := 0; fi217 < 17; fi217++ {
        v1 = v1
        if ((0.0 / 1.0) > 1.5) {
            for fi218 := 0; fi218 < 9; fi218++ {
                switch (fi218 - (fi218 * fi217)) {
                case -105:
                    fmt.Println(v216, v216, fi217)
                case 2147483647:
                    if false {
                        var v219 int8 = -128
                        _ = v219
                    }
                }
                _ = fi218
            }
        }
        v220 := GenS16{F0: (((3.14159 + v216) * 1.5) == v216)}
        _ = fi217
        _ = v220
    }
    for fi221 := 0; fi221 < 16; fi221++ {
        var v222 uint = 2
        fmt.Println(v216, v222)
        _ = fi221
        _ = v222
    }
    var v223 bool = (4294967295 != 42)
    v216 += 0.1
    var v224 uint16 = 1
    v1 = v1
    var v225 string = ""
    for fi226 := 0; fi226 < 4; fi226++ {
        v227 := GenS9{F0: (v224 == v224), F1: (-(v1)), F2: 905}
        var v228 uint32 = 2147483648
        _ = fi226
        _ = v227
        _ = v228
    }
    var v229 int16 = (v1 + 42)
    v1 += 1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v216:", v216)
    fmt.Println("v223:", v223)
    fmt.Println("v224:", v224)
    fmt.Println("v225:", v225)
    fmt.Println("v229:", v229)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v215
    _ = v216
    _ = v223
    _ = v224
    _ = v225
    _ = v229
    _ = v3
    _ = v4
}
