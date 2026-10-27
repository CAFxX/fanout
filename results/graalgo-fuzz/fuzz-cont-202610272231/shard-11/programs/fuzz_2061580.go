package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32; F1 string; F2 int8 }
type GenS5 struct { F0 int16; F1 int32 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int16; F1 uint32; F2 int16 }
type GenS8 struct { F0 uint8; F1 uint; F2 int64 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 uint32; F1 uint32; F2 bool }
type GenS11 struct { F0 uint; F1 int64 }
type GenS12 struct { F0 int64; F1 int; F2 uint64 }
type GenS13 struct { F0 int; F1 int8 }
type GenS14 struct { F0 uint32 }
type GenS15 struct { F0 uint8; F1 int }

func main() {
    v1 := []string{"\n\t", "hello", "a", ("hello" + ("" + "a"))}
    var v2 int8 = 1
    var v3 int32 = -2147483648
    v4 := []int{32767, 4294967295, 1, 1}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi5 := 0; fi5 < 2; fi5++ {
        v3 = -1
        _ = fi5
    }
    var v6 bool = false
    for fi7 := 0; fi7 < 4; fi7++ {
        v2 -= (-(v2))
        _ = fi7
    }
    v6 = v6
    switch 268 {
    case 451:
        for fi8 := 0; fi8 < 3; fi8++ {
            for fi9 := 0; fi9 < 7; fi9++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                if v6 {
                    switch "" {
                    case "\n\t":
                        v10 := GenS1{F0: 42}
                        _ = v10
                    case "hello":
                        v1 = append(v1, ((((("a" + "a") + "hello") + "foo bar") + ("hello" + "a")) + "a"))
                        if v6 {
                            var v11 uint32 = 2147483647
                            var v12 string = "\n\t"
                            v2 -= v2
                            _ = v11
                            _ = v12
                        } else {
                            if ((2147483648 <= 433) || (v6 && (-1 >= 42))) {
                                v13 := GenS1{F0: -1}
                                _ = v13
                            }
                            v1 = append(v1, ("" + "hello"))
                            fmt.Println(v6, v3)
                        }
                        v3 = (v3 << ((255) & 63))
                    }
                    v14 := GenS1{F0: 0}
                    v15 := GenS2{F0: 611}
                    _ = v14
                    _ = v15
                }
                _ = fi9
            }
            for fi16 := 0; fi16 < 9; fi16++ {
                var v17 int8 = (v2 + v2)
                _ = fi16
                _ = v17
            }
            _ = fi8
        }
    case -2147483648:
        if len(v1) > 0 {
            v1[0] = ""
        }
    case 127:
        fmt.Println(v3, v3, v6)
    }
    if (true && v6) {
        v4 = append(v4, 2147483647)
        v6 = v6
    } else {
        var v18 int = -2147483648
        fmt.Println(v3)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v18
    }
    var v19 float32 = (float32((0.0 / 1.5)) - 3.14159)
    if v6 {
        for fi20 := 0; fi20 < 16; fi20++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            for fi21 := 0; fi21 < 10; fi21++ {
                if (-1 > -2) {
                    var v22 int8 = (-(v2))
                    v2 = (^(v22))
                    v23 := GenS2{F0: 957}
                    switch (fi20 + fi21) {
                    case 9223372036854775807:
                        var v24 uint32 = 274
                        _ = v24
                    case 203:
                        v6 = false
                    case -9223372036854775808:
                        switch "foo bar" {
                        case "a":
                            v25 := GenS1{F0: 9223372036854775807}
                            _ = v25
                        case "foo bar":
                            var v26 uint16 = 2
                            _ = v26
                        case "\n\t":
                            fmt.Println(v6, v19, v6)
                        default:
                            fmt.Println(v3, v6)
                        }
                    }
                    _ = v22
                    _ = v23
                } else {
                    var v27 int8 = (v2 & v2)
                    v2 += (v27 * v2)
                    switch fi20 {
                    case 32767:
                        fmt.Println(fi21, v27, fi20)
                    case -117:
                        v28 := GenS2{F0: 42}
                        _ = v28
                    default:
                        var v29 float32 = v19
                        _ = v29
                    }
                    _ = v27
                }
                var v30 uint32 = 65535
                _ = fi21
                _ = v30
            }
            v19 = -2.25
            _ = fi20
        }
        fmt.Println(v19, v2)
    } else {
        switch "hello" {
        case "foo bar":
            if len(v4) > 0 {
                v4[0] = 1
            }
        }
        v2 = v2
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if v6 {
        _ = len(v1)
        if len(v4) > 0 {
            v4[0] = 2147483647
        }
        for fi31 := 0; fi31 < 15; fi31++ {
            var v32 float64 = 0.0
            for fi33 := 0; fi33 < 15; fi33++ {
                if true {
                    if true {
                        switch "\n\t" {
                        case "hello":
                            var v34 float32 = 1.5
                            _ = v34
                        }
                        var v35 uint64 = 18446744073709551615
                        var v36 uint64 = (v35 / (((v35 ^ 2)) | 1))
                        _ = v35
                        _ = v36
                    } else {
                        v3 = v3
                    }
                    v37 := GenS1{F0: -1}
                    v3 = (-(v3))
                    v6 = v6
                    if (v19 != 0.1) {
                        if (v6 || v6) {
                            if (!v6) {
                                v6 = ((^(fi31)) >= 42)
                            } else {
                                var v38 uint32 = 205
                                switch 1 {
                                case -589:
                                    v6 = v6
                                case -853:
                                    if ("hello" == "hello") {
                                        v6 = v6
                                        v4 = append(v4, fi33)
                                        var v39 uint = 18446744073709551615
                                        _ = v39
                                    }
                                default:
                                    fmt.Println(v32, v3, v2)
                                }
                                _ = v38
                            }
                        }
                        v19 -= float32((1e10 + v32))
                    }
                    _ = v37
                } else {
                    fmt.Println(fi31, v32, fi33)
                }
                var v40 uint16 = 0
                switch fi31 {
                case -9223372036854775808:
                    v32 += v32
                case 652:
                    v32 -= 0.1
                case 994:
                    var v41 float64 = ((0.0 / v32) + (((-2.25 * 0.0) - -2.25) / float64(v19)))
                    _ = v41
                }
                switch ("" + "hello") {
                case "foo bar":
                    v6 = (((v32 - (((float64(v19) + ((v32 * ((float64(1.5) + ((3.14159 + 1e10) / (v32 / 1.0))) - 3.14159)) + float64(v19))) - float64(v19)) - v32)) * v32) >= 3.14159)
                case "hello":
                    if v6 {
                        var v42 int64 = 9223372036854775807
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = v42
                    }
                }
                _ = len(v4)
                var v43 int = (^(fi31))
                _ = fi33
                _ = v40
                _ = v43
            }
            _ = fi31
            _ = v32
        }
    } else {
        fmt.Println(v6, v3)
        switch "foo bar" {
        case "foo bar":
            v44 := GenS1{F0: 0}
            _ = v44
        case "":
            for fi45 := 0; fi45 < 3; fi45++ {
                var v46 int32 = (v3 ^ v3)
                for fi47 := 0; fi47 < 15; fi47++ {
                    var v48 uint8 = 42
                    v6 = v6
                    v6 = (42 > 1)
                    _ = fi47
                    _ = v48
                }
                _ = fi45
                _ = v46
            }
        }
        switch -1 {
        case 1:
            if true {
                if true {
                    if len(v1) > 0 {
                        v1[0] = (((("" + "a") + ("\n\t" + "")) + ("" + "")) + ("" + ("a" + "a")))
                    }
                    if (!v6) {
                        v19 = 0.1
                        switch 1 {
                        case 128:
                            fmt.Println(v19)
                        case 1:
                            switch (((((("a" + "hello") + "foo bar") + "") + "") + ("\n\t" + ("hello" + "a"))) + ("foo bar" + ("a" + ("foo bar" + "\n\t")))) {
                            case "foo bar":
                                if (!false) {
                                    v49 := GenS1{F0: -794}
                                    if false {
                                        for fi50 := 0; fi50 < 3; fi50++ {
                                            v19 = 1e10
                                            _ = fi50
                                        }
                                        for fi51 := 0; fi51 < 19; fi51++ {
                                            v3 -= 0
                                            _ = fi51
                                        }
                                        if ((v2 % ((-10) | 1)) >= (v2 ^ v2)) {
                                            if false {
                                                if (238 <= -2147483648) {
                                                    var v52 int = 128
                                                    _ = v52
                                                } else {
                                                    for fi53 := 0; fi53 < 8; fi53++ {
                                                        v54 := GenS2{F0: 580}
                                                        _ = fi53
                                                        _ = v54
                                                    }
                                                    for fi55 := 0; fi55 < 15; fi55++ {
                                                        switch "" {
                                                        case "a":
                                                            var v56 int64 = 9223372036854775807
                                                            _ = v56
                                                        case "hello":
                                                            v57 := GenS1{F0: 1}
                                                            _ = v57
                                                        case "":
                                                            for fi58 := 0; fi58 < 3; fi58++ {
                                                                v59 := GenS2{F0: 2}
                                                                switch ("hello" + "foo bar") {
                                                                case "a":
                                                                    v2 = v2
                                                                case "hello":
                                                                    v6 = true
                                                                }
                                                                _ = fi58
                                                                _ = v59
                                                            }
                                                        default:
                                                            var v60 int64 = -1
                                                            _ = v60
                                                        }
                                                        v2 -= 100
                                                        _ = fi55
                                                    }
                                                }
                                            } else {
                                                var v61 float64 = (1e10 / (3.14159 - 0.0))
                                                fmt.Println(v6, v6, v61)
                                                _ = v61
                                            }
                                        } else {
                                            switch ("foo bar" + ("foo bar" + "a")) {
                                            case "foo bar":
                                                if len(v4) > 0 {
                                                    v4[0] = 4294967295
                                                }
                                            case "\n\t":
                                                for fi62 := 0; fi62 < 10; fi62++ {
                                                    for fi63 := 0; fi63 < 4; fi63++ {
                                                        if false {
                                                            v64 := GenS2{F0: 943}
                                                            v19 += 1e10
                                                            v2 = v2
                                                            _ = v64
                                                        } else {
                                                            switch (("foo bar" + ("foo bar" + "hello")) + "hello") {
                                                            case "hello":
                                                                if false {
                                                                    switch fi62 {
                                                                    case -868:
                                                                        v65 := GenS2{F0: 1}
                                                                        _ = v65
                                                                    case 127:
                                                                        if ((v6 || true) && true) {
                                                                            switch "a" {
                                                                            case "hello":
                                                                                if v6 {
                                                                                    var v66 uint64 = 932
                                                                                    switch "\n\t" {
                                                                                    case "a":
                                                                                        var v67 int16 = -387
                                                                                        _ = v67
                                                                                    }
                                                                                    v19 = v19
                                                                                    var v68 float64 = 0.1
                                                                                    _ = v66
                                                                                    _ = v68
                                                                                } else {
                                                                                    if (v6 && v6) {
                                                                                        _ = len(v1)
                                                                                        v69 := GenS3{F0: 2}
                                                                                        _ = v69
                                                                                    } else {
                                                                                        v70 := GenS4{F0: 4294967295, F1: "foo bar", F2: 1}
                                                                                        switch fi63 {
                                                                                        case 687:
                                                                                            v71 := GenS1{F0: -549}
                                                                                            _ = v71
                                                                                        case 4294967295:
                                                                                            v72 := GenS4{F0: 288, F1: "hello", F2: v2}
                                                                                            _ = v72
                                                                                        }
                                                                                        v2 -= 42
                                                                                        switch (("a" + "hello") + ((("\n\t" + ("\n\t" + ("a" + "foo bar"))) + ("a" + ("a" + "foo bar"))) + "hello")) {
                                                                                        case "\n\t":
                                                                                            var v73 int32 = (v3 * (v3 << ((255) & 63)))
                                                                                            _ = v73
                                                                                        case "hello":
                                                                                            fmt.Println(v6, v2)
                                                                                            var v74 int8 = (v2 + v2)
                                                                                            _ = v74
                                                                                        case "a":
                                                                                            fmt.Println(v3, v2, v19)
                                                                                        }
                                                                                        _ = v70
                                                                                    }
                                                                                    v75 := GenS5{F0: 32767, F1: (v3 | (v3 << ((0) & 63)))}
                                                                                    v76 := GenS4{F0: 42, F1: "hello", F2: v2}
                                                                                    _ = v75
                                                                                    _ = v76
                                                                                }
                                                                            case "":
                                                                                v1 = append(v1, "\n\t")
                                                                            case "\n\t":
                                                                                v77 := GenS1{F0: 0}
                                                                                if len(v1) >= 2 {
                                                                                    v1 = v1[1:len(v1)]
                                                                                }
                                                                                if len(v4) > 0 {
                                                                                    v4[0] = 773
                                                                                }
                                                                                _ = v77
                                                                            default:
                                                                                v3 -= (v3 << ((255) & 63))
                                                                            }
                                                                        }
                                                                    default:
                                                                        _ = len(v4)
                                                                    }
                                                                } else {
                                                                    switch (("foo bar" + "") + "a") {
                                                                    case "foo bar":
                                                                        if v6 {
                                                                            if v6 {
                                                                                fmt.Println(fi62)
                                                                                var v78 int = 8
                                                                                switch "" {
                                                                                case "\n\t":
                                                                                    var v79 int = (^(v78))
                                                                                    _ = v79
                                                                                case "hello":
                                                                                    var v80 float32 = v19
                                                                                    _ = v80
                                                                                case "a":
                                                                                    var v81 int = v78
                                                                                    _ = v81
                                                                                default:
                                                                                    v82 := GenS4{F0: 487, F1: "foo bar", F2: int8(v3)}
                                                                                    _ = v82
                                                                                }
                                                                                v3 += v3
                                                                                fmt.Println(v3, v2)
                                                                                _ = v78
                                                                            }
                                                                        } else {
                                                                            var v83 string = "\n\t"
                                                                            if (!(500 <= -2147483648)) {
                                                                                var v84 string = "a"
                                                                                _ = v84
                                                                            }
                                                                            _ = v83
                                                                        }
                                                                    default:
                                                                        if (!v6) {
                                                                            if v6 {
                                                                                v85 := GenS4{F0: 740, F1: "hello", F2: (v2 ^ v2)}
                                                                                var v86 int16 = 956
                                                                                if ((1 < 4294967295) && v6) {
                                                                                    switch fi63 {
                                                                                    case 127:
                                                                                        fmt.Println(v19, fi62, fi63)
                                                                                        var v87 uint64 = 0
                                                                                        if v6 {
                                                                                            if false {
                                                                                                switch "\n\t" {
                                                                                                case "a":
                                                                                                    _ = len(v4)
                                                                                                default:
                                                                                                    var v88 uint16 = 837
                                                                                                    _ = v88
                                                                                                }
                                                                                                var v89 uint16 = 115
                                                                                                v86 = (v86 | (-(v86)))
                                                                                                _ = v89
                                                                                            }
                                                                                            fmt.Println(v3, v3, v19)
                                                                                        }
                                                                                        _ = v87
                                                                                    default:
                                                                                        fmt.Println(fi62, v19, v2)
                                                                                    }
                                                                                    var v90 int32 = (v3 * 0)
                                                                                    v3 -= -2147483648
                                                                                    _ = v90
                                                                                }
                                                                                _ = v85
                                                                                _ = v86
                                                                            }
                                                                            if len(v4) >= 2 {
                                                                                v4 = v4[1:len(v4)]
                                                                            }
                                                                            v91 := GenS3{F0: 127}
                                                                            _ = v91
                                                                        }
                                                                    }
                                                                    fmt.Println(v6, fi63, v19)
                                                                }
                                                            case "foo bar":
                                                                if (1.5 < v19) {
                                                                    if false {
                                                                        v4 = append(v4, fi63)
                                                                        switch "\n\t" {
                                                                        case "hello":
                                                                            var v92 string = ""
                                                                            _ = v92
                                                                        case "\n\t":
                                                                            v2 = 100
                                                                        default:
                                                                            fmt.Println(v6, v6)
                                                                        }
                                                                        var v93 int64 = 0
                                                                        _ = v93
                                                                    } else {
                                                                        var v94 int16 = 177
                                                                        _ = v94
                                                                    }
                                                                } else {
                                                                    fmt.Println(fi62, v3, v19)
                                                                    if (false || v6) {
                                                                        if ((!(0 >= 0)) || ((-32768 > 83) && v6)) {
                                                                            v19 = ((v19 - v19) / 0.1)
                                                                            var v95 bool = (45 != 1)
                                                                            var v96 bool = (v19 != v19)
                                                                            _ = v95
                                                                            _ = v96
                                                                        } else {
                                                                            v97 := GenS5{F0: -810, F1: -1}
                                                                            if (((0.0 - 1e10) == 0.0) && ((!(!v6)) || (352 == 0))) {
                                                                                if true {
                                                                                    v98 := GenS6{F0: fi62}
                                                                                    fmt.Println(v2)
                                                                                    v19 -= (float32(3.14159) - 0.0)
                                                                                    v99 := GenS7{F0: 1000, F1: 2147483648, F2: 1}
                                                                                    switch "foo bar" {
                                                                                    case "foo bar":
                                                                                        v3 += (v3 << ((18446744073709551615) & 63))
                                                                                    case "a":
                                                                                        v3 += v3
                                                                                    case "\n\t":
                                                                                        var v100 uint8 = 214
                                                                                        _ = v100
                                                                                    }
                                                                                    v101 := GenS1{F0: -516}
                                                                                    _ = v101
                                                                                    _ = v98
                                                                                    _ = v99
                                                                                } else {
                                                                                    v102 := GenS8{F0: 2, F1: 65535, F2: 1}
                                                                                    if ("hello" < "\n\t") {
                                                                                        v103 := GenS9{F0: ("a" + "foo bar")}
                                                                                        fmt.Println(v3, v19, v2)
                                                                                        v104 := GenS8{F0: 1, F1: 0, F2: -232}
                                                                                        _ = v103
                                                                                        _ = v104
                                                                                    } else {
                                                                                        var v105 uint8 = 167
                                                                                        _ = len(v1)
                                                                                        fmt.Println(v19)
                                                                                        _ = v105
                                                                                    }
                                                                                    v106 := GenS10{F0: 42, F1: 17, F2: (v6 && true)}
                                                                                    _ = v102
                                                                                    _ = v106
                                                                                }
                                                                                v6 = v6
                                                                            }
                                                                            v107 := GenS11{F0: 423, F1: 32}
                                                                            _ = v107
                                                                            _ = v97
                                                                        }
                                                                        _ = len(v1)
                                                                    }
                                                                }
                                                            default:
                                                                v3 += v3
                                                            }
                                                            var v108 int16 = int16(fi63)
                                                            _ = v108
                                                        }
                                                        _ = fi63
                                                    }
                                                    var v109 uint32 = 2
                                                    var v110 uint = 18446744073709551615
                                                    _ = fi62
                                                    _ = v109
                                                    _ = v110
                                                }
                                            default:
                                                for fi111 := 0; fi111 < 11; fi111++ {
                                                    if false {
                                                        v19 = (v19 - (1e10 / v19))
                                                        var v112 uint16 = 2
                                                        v112 = v112
                                                        _ = v112
                                                    } else {
                                                        var v113 float32 = -2.25
                                                        _ = v113
                                                    }
                                                    v114 := GenS8{F0: 1, F1: 1, F2: 0}
                                                    v1 = append(v1, ("a" + "\n\t"))
                                                    _ = fi111
                                                    _ = v114
                                                }
                                            }
                                            if (v2 < (v2 % (((v2 % (((v2 & v2)) | 1))) | 1))) {
                                                var v115 int = 0
                                                switch "a" {
                                                case "hello":
                                                    var v116 int16 = 1
                                                    _ = v116
                                                case "":
                                                    var v117 bool = (311 < 0)
                                                    _ = v117
                                                }
                                                v118 := GenS12{F0: -1, F1: v115, F2: 4294967295}
                                                _ = v115
                                                _ = v118
                                            } else {
                                                for fi119 := 0; fi119 < 13; fi119++ {
                                                    v19 += 1.5
                                                    _ = fi119
                                                }
                                                v120 := GenS8{F0: 1, F1: 18446744073709551615, F2: 1}
                                                _ = v120
                                            }
                                        }
                                    }
                                    for fi121 := 0; fi121 < 17; fi121++ {
                                        var v122 uint = 800
                                        switch ("a" + "hello") {
                                        case "hello":
                                            v6 = (32767 == 67)
                                        case "":
                                            var v123 float32 = -2.25
                                            _ = v123
                                        case "foo bar":
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        }
                                        _ = fi121
                                        _ = v122
                                    }
                                    _ = v49
                                }
                            }
                        default:
                            v124 := GenS10{F0: 660, F1: 414, F2: (1000 < -1)}
                            _ = v124
                        }
                    } else {
                        if (255 >= 0) {
                            var v125 uint32 = 2147483647
                            v19 -= ((v19 + v19) - 0.1)
                            _ = v125
                        } else {
                            v126 := GenS7{F0: -1, F1: 222, F2: 1000}
                            _ = v126
                        }
                        v2 -= (v2 * (v2 << ((2) & 63)))
                    }
                    v127 := GenS13{F0: 127, F1: 108}
                    _ = v127
                }
            }
        case -1:
            v128 := GenS3{F0: 0}
            _ = v128
        }
    }
    v129 := GenS14{F0: 503}
    for fi130 := 0; fi130 < 19; fi130++ {
        for fi131 := 0; fi131 < 15; fi131++ {
            if (0 == v3) {
                if v6 {
                    fmt.Println(v6, fi131, v6)
                    if (!false) {
                        v2 += v2
                        _ = len(v4)
                    }
                    var v132 uint = 18446744073709551615
                    _ = v132
                } else {
                    var v133 int64 = 9223372036854775807
                    if (-1 == 32767) {
                        fmt.Println(v3)
                        if len(v4) > 0 {
                            v4[0] = fi131
                        }
                    } else {
                        v134 := GenS15{F0: 14, F1: 128}
                        _ = v134
                    }
                    v19 = 0.0
                    _ = v133
                }
            }
            _ = fi131
        }
        _ = fi130
    }
    fmt.Println(v19, v3, v3)
    switch 4294967295 {
    case -308:
        switch -9223372036854775808 {
        case -293:
            for fi135 := 0; fi135 < 6; fi135++ {
                switch fi135 {
                case 32767:
                    var v136 float64 = 1.5
                    _ = v136
                }
                fmt.Println(v19, v6)
                _ = fi135
            }
        }
    }
    var v137 int32 = (-(v3))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch -2147483648 {
    case 127:
        var v138 uint64 = 0
        _ = v138
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v137:", v137)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v129
    _ = v137
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
