package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int32; F1 int64; F2 uint16 }
type GenS5 struct { F0 uint32; F1 uint16 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int64; F1 int64 }
type GenS8 struct { F0 int; F1 bool; F2 uint8 }
type GenS9 struct { F0 string; F1 int32; F2 int8 }
type GenS10 struct { F0 int16; F1 uint32 }
type GenS11 struct { F0 uint }

func main() {
    var v1 int32 = -1
    var v2 uint64 = 18446744073709551615
    v3 := []uint32{}
    v4 := make(map[uint32]bool)
    var v5 float64 = 1.5
    v6 := make(map[uint32]string)
    var v7 int8 = 42
    if false {
        switch 42 {
        case -9223372036854775808:
            for fi8 := 0; fi8 < 20; fi8++ {
                v7 = (v7 ^ (v7 ^ 100))
                if false {
                    v7 = v7
                    v9 := GenS1{F0: "foo bar", F1: 162}
                    _ = v9
                }
                if true {
                    v3 = append(v3, uint32(v2))
                    var v10 int64 = 1
                    fmt.Println(v1, v5)
                    _ = v10
                } else {
                    if (false && true) {
                        fmt.Println(v1, v7)
                        var v11 int32 = (v1 & v1)
                        switch ((((("\n\t" + "") + "a") + "\n\t") + "") + "") {
                        case "a":
                            fmt.Println(fi8, v1, v5)
                        }
                        _ = v11
                    }
                    var v12 string = "hello"
                    var v13 float32 = 0.0
                    _ = v12
                    _ = v13
                }
                _ = fi8
            }
        case 42:
            var v14 float32 = 0.1
            _ = v14
        default:
            if true {
                if true {
                    var v15 uint64 = 2
                    switch 32767 {
                    case -9223372036854775808:
                        if true {
                            v2 += 193
                        } else {
                            v7 += v7
                        }
                    case 1:
                        v5 = v5
                    default:
                        var v16 bool = (((694 < 505) && false) && ("\n\t" == "a"))
                        _ = v16
                    }
                    for fi17 := 0; fi17 < 7; fi17++ {
                        if false {
                            switch 128 {
                            case 1:
                                v6[2] = "\n\t"
                            case -2147483648:
                                if len(v3) > 0 {
                                    v3[0] = 724
                                }
                            default:
                                if false {
                                    switch "\n\t" {
                                    case "":
                                        v1 += 42
                                    case "\n\t":
                                        var v18 int = int(v2)
                                        _ = v18
                                    case "hello":
                                        if (247 != 65535) {
                                            if (42 > int8(v2)) {
                                                if (v1 != -1) {
                                                    var v19 int64 = 1
                                                    for fi20 := 0; fi20 < 20; fi20++ {
                                                        if (4294967295 == 946) {
                                                            if true {
                                                                fmt.Println(v1, v7, v1)
                                                                v21 := v4[2]
                                                                _ = v21
                                                            } else {
                                                                v7 = 82
                                                            }
                                                            if (false || true) {
                                                                if false {
                                                                    v7 += v7
                                                                }
                                                                v5 = (float64(((0.1 - 1e10) + 1.5)) - float64((0.0 - (((3.14159 / 1e10) + (3.14159 / 0.1)) - float32(v5)))))
                                                            }
                                                            if len(v3) > 0 {
                                                                v3[0] = 468
                                                            }
                                                        } else {
                                                            var v22 float32 = float32(v5)
                                                            _ = v22
                                                        }
                                                        if false {
                                                            var v23 uint32 = 849
                                                            v24 := GenS1{F0: "a", F1: 128}
                                                            v23 -= uint32(v15)
                                                            _ = v23
                                                            _ = v24
                                                        } else {
                                                            fmt.Println(fi20, v19, v1)
                                                        }
                                                        _ = fi20
                                                    }
                                                    _ = v19
                                                }
                                            }
                                            switch 296 {
                                            case -1:
                                                for fi25 := 0; fi25 < 8; fi25++ {
                                                    if (!false) {
                                                        fmt.Println(v1, fi25)
                                                        switch 42 {
                                                        case 128:
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                        default:
                                                            if false {
                                                                switch "" {
                                                                case "\n\t":
                                                                    v15 = 1
                                                                default:
                                                                    v26 := GenS2{F0: -32768, F1: (^(v1))}
                                                                    _ = v26
                                                                }
                                                                v27 := GenS1{F0: "", F1: 173}
                                                                _ = v27
                                                            }
                                                        }
                                                        var v28 int = fi17
                                                        _ = v28
                                                    } else {
                                                        v29 := GenS2{F0: 118, F1: v1}
                                                        v30 := GenS2{F0: -32768, F1: (v1 + v1)}
                                                        _ = v29
                                                        _ = v30
                                                    }
                                                    var v31 int32 = (v1 % ((v1) | 1))
                                                    v32 := GenS3{F0: (v15 % (((v15 + (v15 + 18446744073709551615))) | 1))}
                                                    _ = fi25
                                                    _ = v31
                                                    _ = v32
                                                }
                                            case -9223372036854775808:
                                                switch fi17 {
                                                case -825:
                                                    var v33 bool = true
                                                    _ = v33
                                                case -2147483648:
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                }
                                            case 128:
                                                if true {
                                                    var v34 string = ""
                                                    _ = v34
                                                }
                                            default:
                                                var v35 uint8 = 157
                                                _ = v35
                                            }
                                        } else {
                                            for fi36 := 0; fi36 < 2; fi36++ {
                                                var v37 uint32 = 583
                                                _ = fi36
                                                _ = v37
                                            }
                                            var v38 uint8 = 120
                                            v38 += (^(v38))
                                            _ = v38
                                        }
                                    default:
                                        v39 := v4[2147483648]
                                        _ = v39
                                    }
                                    _ = len(v6)
                                } else {
                                    v15 = v2
                                    fmt.Println(v2, v5)
                                }
                            }
                            if ((v5 / v5) <= v5) {
                                switch "hello" {
                                case "hello":
                                    fmt.Println(v2)
                                case "foo bar":
                                    var v40 uint64 = v15
                                    _ = v40
                                default:
                                    var v41 int16 = -130
                                    _ = v41
                                }
                                v42 := GenS2{F0: -32768, F1: (v1 * v1)}
                                v7 -= (v7 >> ((42) & 63))
                                _ = v42
                            } else {
                                v1 -= (^(v1))
                                v2 = v2
                            }
                            v3 = append(v3, 0)
                        }
                        var v43 uint8 = 1
                        _ = fi17
                        _ = v43
                    }
                    _ = v15
                } else {
                    fmt.Println(v1)
                    v3 = append(v3, 42)
                    for fi44 := 0; fi44 < 11; fi44++ {
                        for fi45 := 0; fi45 < 10; fi45++ {
                            var v46 int = fi45
                            v47 := GenS2{F0: 1, F1: (v1 - v1)}
                            v7 = v7
                            _ = fi45
                            _ = v46
                            _ = v47
                        }
                        _ = fi44
                    }
                }
                for fi48 := 0; fi48 < 13; fi48++ {
                    fmt.Println(fi48, v2, v5)
                    v1 += v1
                    v5 = -2.25
                    _ = fi48
                }
            }
        }
    }
    var v49 float32 = (1e10 * 3.14159)
    v50 := GenS4{F0: (v1 * (-(v1))), F1: -1, F2: 268}
    v51 := GenS4{F0: v1, F1: 893, F2: 255}
    v2 += v2
    if len(v3) > 0 {
        v3[0] = 18
    }
    v1 = v1
    if len(v3) > 0 {
        v3[0] = 2
    }
    var v52 bool = true
    v49 += 1.5
    for fi53 := 0; fi53 < 20; fi53++ {
        if true {
            if (v52 || (2 < 42)) {
                v1 = (v1 - 42)
                v52 = v52
                v1 += v1
            } else {
                if ((v7 ^ (^(v7))) > v7) {
                    switch (fi53 / (((-(fi53))) | 1)) {
                    case 0:
                        var v54 float64 = (v5 * v5)
                        _ = v54
                    }
                    switch (^(fi53)) {
                    case 760:
                        for fi55 := 0; fi55 < 8; fi55++ {
                            if v52 {
                                v52 = (-32768 != 0)
                                v2 -= (v2 << ((445) & 63))
                            } else {
                                var v56 string = "\n\t"
                                if (v2 == v2) {
                                    if (!(!v52)) {
                                        if (2 == 42) {
                                            v57 := GenS5{F0: 4294967295, F1: 2}
                                            v58 := GenS6{F0: 0}
                                            _ = v57
                                            _ = v58
                                        }
                                        if v52 {
                                            _ = len(v3)
                                            v5 = (1.5 * (v5 * (float64(v49) + 0.1)))
                                        } else {
                                            v3 = append(v3, 0)
                                            v56 = v56
                                            fmt.Println(v7, v2)
                                            if ((v2 > (^(v2))) || false) {
                                                var v59 int8 = (v7 / ((-1) | 1))
                                                _ = v59
                                            }
                                        }
                                        v1 -= v1
                                    }
                                    v60 := GenS4{F0: 1, F1: -9223372036854775808, F2: 393}
                                    _ = v60
                                }
                                var v61 uint = 65535
                                _ = v56
                                _ = v61
                            }
                            fmt.Println(fi53, v5, v49)
                            _ = fi55
                        }
                    case 372:
                        switch "" {
                        case "foo bar":
                            v5 += 0.1
                        case "hello":
                            v62 := GenS7{F0: 211, F1: 9223372036854775807}
                            _ = v62
                        default:
                            v52 = v52
                        }
                    default:
                        _ = len(v6)
                    }
                    var v63 int64 = 9223372036854775807
                    _ = v63
                }
                fmt.Println(v2, v7)
            }
            v64 := v6[2147483647]
            _ = v64
        } else {
            var v65 int64 = -9223372036854775808
            var v66 uint64 = (v2 % ((v2) | 1))
            switch "foo bar" {
            case "":
                switch ("" + (("" + "foo bar") + "hello")) {
                case "hello":
                    var v67 int8 = (v7 << ((975) & 63))
                    _ = v67
                case "":
                    for fi68 := 0; fi68 < 5; fi68++ {
                        if (v1 <= (^(v1))) {
                            fmt.Println(v49, v66)
                            if v52 {
                                var v69 int = 2147483647
                                var v70 uint8 = 1
                                _ = v69
                                _ = v70
                            } else {
                                v7 = v7
                                v71 := GenS8{F0: (fi53 | (fi53 - int(v65))), F1: ((9223372036854775807 <= v65) || v52), F2: 42}
                                switch 0 {
                                case -2147483648:
                                    if (-590 >= 42) {
                                        v72 := v6[4294967295]
                                        _ = v72
                                    } else {
                                        if v52 {
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        }
                                        v73 := GenS7{F0: 9223372036854775807, F1: (-(v65))}
                                        if len(v3) > 0 {
                                            v3[0] = 42
                                        }
                                        _ = v73
                                    }
                                case -2:
                                    switch (fi53 - (fi53 + fi53)) {
                                    case 1:
                                        if v52 {
                                            switch "foo bar" {
                                            case "hello":
                                                v7 += v7
                                            case "foo bar":
                                                if (-32768 != 378) {
                                                    var v74 uint = 255
                                                    switch "foo bar" {
                                                    case "a":
                                                        if (!(v74 == 65535)) {
                                                            if ((v66 > (v2 - (^(v66)))) || true) {
                                                                switch "\n\t" {
                                                                case "\n\t":
                                                                    var v75 int8 = (v7 | int8(v65))
                                                                    _ = v75
                                                                case "foo bar":
                                                                    if (32767 >= -1) {
                                                                        v3 = append(v3, 2147483648)
                                                                        var v76 int8 = int8(v65)
                                                                        _ = v76
                                                                    }
                                                                default:
                                                                    var v77 float32 = v49
                                                                    _ = v77
                                                                }
                                                                var v78 uint32 = 2147483647
                                                                var v79 uint16 = 299
                                                                _ = v78
                                                                _ = v79
                                                            }
                                                        } else {
                                                            _ = len(v3)
                                                            if len(v3) > 0 {
                                                                v3[0] = 237
                                                            }
                                                        }
                                                    case "\n\t":
                                                        v80 := GenS2{F0: 1000, F1: int32(v66)}
                                                        _ = v80
                                                    }
                                                    v1 += v1
                                                    _ = v74
                                                } else {
                                                    var v81 string = "a"
                                                    var v82 int8 = -64
                                                    _ = v81
                                                    _ = v82
                                                }
                                                if len(v3) > 0 {
                                                    v3[0] = 0
                                                }
                                            default:
                                                var v83 uint32 = 2147483648
                                                _ = v83
                                            }
                                        } else {
                                            if len(v3) > 0 {
                                                v3[0] = 24
                                            }
                                            var v84 float32 = v49
                                            switch fi53 {
                                            case 4294967295:
                                                var v85 uint = 1
                                                _ = v85
                                            case 128:
                                                if false {
                                                    switch 995 {
                                                    case 765:
                                                        var v86 int64 = v65
                                                        _ = v86
                                                    case 4294967295:
                                                        v87 := GenS6{F0: -32768}
                                                        _ = v87
                                                    }
                                                    v66 = (v2 << ((4294967295) & 63))
                                                    var v88 float32 = (v49 / v84)
                                                    _ = v88
                                                }
                                            case -503:
                                                if ((fi68 - (-(fi68))) != fi68) {
                                                    switch "" {
                                                    case "hello":
                                                        var v89 uint16 = 779
                                                        _ = v89
                                                    case "\n\t":
                                                        if ((!((!((v52 || v52) && (v52 || true))) && v52)) || (float64(1.5) <= (float64(v84) + 0.0))) {
                                                            v90 := GenS9{F0: "foo bar", F1: (v1 - (v1 - (^(v1)))), F2: v7}
                                                            _ = v90
                                                        } else {
                                                            fmt.Println(v5, v1, fi68)
                                                        }
                                                    }
                                                }
                                            }
                                            _ = v84
                                        }
                                    case 0:
                                        v5 -= float64(v49)
                                    default:
                                        v5 = float64(v49)
                                    }
                                case -9223372036854775808:
                                    fmt.Println(v1, v7, v5)
                                }
                                _ = v71
                            }
                            fmt.Println(fi68, v5)
                            v66 -= (v66 % ((v2) | 1))
                        } else {
                            v65 = (v65 >> ((0) & 63))
                            switch fi68 {
                            case -76:
                                var v91 string = ("foo bar" + "hello")
                                _ = v91
                            case 42:
                                v92 := v6[2147483648]
                                _ = v92
                            case 9223372036854775807:
                                v93 := GenS3{F0: (v66 << ((0) & 63))}
                                _ = v93
                            }
                            if len(v3) > 0 {
                                v3[0] = 4294967295
                            }
                        }
                        _ = fi68
                    }
                default:
                    for fi94 := 0; fi94 < 18; fi94++ {
                        switch fi53 {
                        case 9223372036854775807:
                            v5 = v5
                        case -907:
                            v95 := GenS8{F0: (fi94 - fi53), F1: (v49 == (v49 - v49)), F2: 127}
                            _ = v95
                        }
                        var v96 string = "foo bar"
                        _ = fi94
                        _ = v96
                    }
                }
            case "\n\t":
                v65 += (-(v65))
            case "a":
                for fi97 := 0; fi97 < 11; fi97++ {
                    v98 := GenS2{F0: -324, F1: v1}
                    v66 = 2
                    _ = fi97
                    _ = v98
                }
            default:
                if len(v3) > 0 {
                    v3[0] = 720
                }
            }
            _ = v65
            _ = v66
        }
        for fi99 := 0; fi99 < 3; fi99++ {
            v7 = v7
            v100 := GenS5{F0: 42, F1: 255}
            switch "" {
            case "\n\t":
                v1 = v1
            case "a":
                v3 = append(v3, 0)
            default:
                v5 = 1e10
                var v101 bool = (4294967295 < 2147483647)
                _ = v101
            }
            _ = fi99
            _ = v100
        }
        fmt.Println(fi53)
        _ = fi53
    }
    switch 32767 {
    case 127:
        for fi102 := 0; fi102 < 5; fi102++ {
            fmt.Println(fi102, v1, v49)
            v2 -= v2
            var v103 float32 = v49
            _ = fi102
            _ = v103
        }
    default:
        v7 -= v7
    }
    v52 = v52
    if (!v52) {
        switch "foo bar" {
        case "":
            var v104 bool = (v52 && v52)
            _ = v104
        case "\n\t":
            if (836 != 1) {
                switch ("\n\t" + "hello") {
                case "a":
                    for fi105 := 0; fi105 < 16; fi105++ {
                        if true {
                            v7 += (v7 / (((v7 % ((v7) | 1))) | 1))
                            for fi106 := 0; fi106 < 8; fi106++ {
                                var v107 float32 = float32((0.1 * v5))
                                _ = fi106
                                _ = v107
                            }
                        } else {
                            v108 := GenS10{F0: 0, F1: uint32(v2)}
                            _ = v108
                        }
                        fmt.Println(v52)
                        _ = len(v4)
                        _ = fi105
                    }
                case "":
                    var v109 int64 = -716
                    _ = v109
                }
                fmt.Println(v2)
            }
        default:
            fmt.Println(v1, v1)
        }
        var v110 int32 = (v1 - v1)
        _ = v110
    } else {
        var v111 int64 = 9223372036854775807
        for fi112 := 0; fi112 < 5; fi112++ {
            for fi113 := 0; fi113 < 4; fi113++ {
                fmt.Println(fi113, v111, v7)
                var v114 float64 = (((v5 + (((v5 + v5) - ((v5 + 1.5) * v5)) / v5)) + (float64(v49) + 3.14159)) * 0.1)
                if ("\n\t" == "foo bar") {
                    v4[0] = v52
                    switch ("a" + "hello") {
                    case "hello":
                        _ = len(v4)
                    case "a":
                        var v115 int16 = 0
                        _ = v115
                    default:
                        _ = len(v6)
                    }
                }
                _ = fi113
                _ = v114
            }
            _ = fi112
        }
        _ = v111
    }
    v116 := GenS11{F0: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v116
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v6
    _ = v7
}
