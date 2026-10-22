package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 int16; F1 uint64 }
type GenS4 struct { F0 int8; F1 uint; F2 int }
type GenS5 struct { F0 int64; F1 uint16 }
type GenS6 struct { F0 uint32; F1 int64; F2 bool }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int16; F1 int8 }
type GenS11 struct { F0 int; F1 int64; F2 int }

func main() {
    v1 := []string{"\n\t", ("" + ("foo bar" + "\n\t")), ("\n\t" + "foo bar"), "\n\t"}
    var v2 uint64 = 2
    var v3 float64 = 0.0
    var v4 int32 = 1
    v5 := []string{"hello"}
    v3 -= 1.5
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi6 := 0; fi6 < 2; fi6++ {
        v4 -= v4
        _ = fi6
    }
    switch "hello" {
    case "":
        v7 := GenS1{F0: 2, F1: 1}
        _ = v7
    case "a":
        switch "foo bar" {
        case "":
            var v8 int8 = 100
            _ = v8
        case "\n\t":
            var v9 uint64 = 18446744073709551615
            _ = v9
        }
    default:
        v10 := GenS1{F0: 521, F1: (v2 - v2)}
        _ = v10
    }
    v11 := GenS1{F0: 215, F1: 18446744073709551615}
    if false {
        var v12 int = -2147483648
        _ = v12
    }
    var v13 uint64 = v2
    var v14 int8 = 127
    fmt.Println(v14, v3)
    fmt.Println(v3, v3)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi15 := 0; fi15 < 20; fi15++ {
        v2 += v2
        switch fi15 {
        case -9223372036854775808:
            v5 = append(v5, "foo bar")
        case 1:
            switch fi15 {
            case 2147483647:
                switch (fi15 | -350) {
                case 32767:
                    _ = len(v1)
                case 2147483647:
                    v14 = (v14 * v14)
                case 128:
                    v14 = -1
                }
            default:
                for fi16 := 0; fi16 < 3; fi16++ {
                    v14 += 127
                    var v17 int16 = 32767
                    switch fi15 {
                    case 515:
                        v18 := GenS2{F0: (2 == 42), F1: 127}
                        var v19 int8 = 1
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = v18
                        _ = v19
                    }
                    if true {
                        var v20 uint8 = 2
                        var v21 int64 = 1
                        _ = v20
                        _ = v21
                    }
                    var v22 string = ""
                    _ = fi16
                    _ = v17
                    _ = v22
                }
            }
        default:
            if (65535 != 2) {
                for fi23 := 0; fi23 < 9; fi23++ {
                    var v24 int8 = (v14 % ((-33) | 1))
                    _ = fi23
                    _ = v24
                }
            } else {
                if (!false) {
                    v13 -= (v13 >> ((2) & 63))
                    if true {
                        if (1 == 42) {
                            switch "" {
                            case "\n\t":
                                var v25 uint8 = 1
                                _ = v25
                            case "a":
                                var v26 int16 = -32768
                                _ = v26
                            case "foo bar":
                                v14 -= (v14 + (v14 ^ v14))
                            default:
                                for fi27 := 0; fi27 < 18; fi27++ {
                                    fmt.Println(v14, v4)
                                    var v28 float64 = v3
                                    _ = fi27
                                    _ = v28
                                }
                            }
                            if (v13 < v2) {
                                var v29 float64 = v3
                                _ = v29
                            } else {
                                v3 += (1e10 - v3)
                                switch "" {
                                case "hello":
                                    switch (-(fi15)) {
                                    case -580:
                                        var v30 int16 = -948
                                        _ = v30
                                    }
                                default:
                                    fmt.Println(v14, v4, v4)
                                }
                                for fi31 := 0; fi31 < 8; fi31++ {
                                    if len(v1) > 0 {
                                        v1[0] = "hello"
                                    }
                                    _ = fi31
                                }
                            }
                        } else {
                            fmt.Println(v2)
                            switch fi15 {
                            case 4294967295:
                                v32 := GenS2{F0: (!(!(!false))), F1: 1}
                                _ = v32
                            case 0:
                                var v33 uint16 = 940
                                _ = v33
                            }
                        }
                    }
                } else {
                    for fi34 := 0; fi34 < 3; fi34++ {
                        v2 = v2
                        _ = fi34
                    }
                }
            }
        }
        _ = fi15
    }
    switch "hello" {
    case "a":
        if true {
            v14 = (-(v14))
            v4 = -1
            v4 = 1
        } else {
            v35 := GenS3{F0: 42, F1: v13}
            if (((-(v14)) != (v14 - v14)) && true) {
                for fi36 := 0; fi36 < 15; fi36++ {
                    var v37 int8 = (v14 + 98)
                    switch "\n\t" {
                    case "\n\t":
                        v13 -= (v2 & (^(v13)))
                    case "foo bar":
                        v1 = append(v1, ("\n\t" + "a"))
                    case "hello":
                        for fi38 := 0; fi38 < 2; fi38++ {
                            v5 = append(v5, "hello")
                            v13 -= v2
                            _ = fi38
                        }
                    default:
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    }
                    _ = fi36
                    _ = v37
                }
                if false {
                    v13 = (v13 + 2)
                }
            }
            var v39 int8 = 42
            _ = v35
            _ = v39
        }
    case "":
        for fi40 := 0; fi40 < 5; fi40++ {
            v13 += v13
            _ = fi40
        }
    default:
        if true {
            if true {
                fmt.Println(v2, v13)
            }
        } else {
            switch "a" {
            case "hello":
                for fi41 := 0; fi41 < 8; fi41++ {
                    switch fi41 {
                    case 106:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    case 128:
                        if len(v5) > 0 {
                            v5[0] = "a"
                        }
                    }
                    v42 := GenS1{F0: 1, F1: 2}
                    v43 := GenS2{F0: false, F1: 185}
                    _ = fi41
                    _ = v42
                    _ = v43
                }
            case "\n\t":
                var v44 uint16 = 65535
                _ = v44
            case "a":
                fmt.Println(v2, v2, v2)
            default:
                if false {
                    if true {
                        var v45 uint64 = v13
                        v46 := GenS2{F0: true, F1: 42}
                        _ = v45
                        _ = v46
                    } else {
                        for fi47 := 0; fi47 < 16; fi47++ {
                            if (int32(v14) <= (^(v4))) {
                                if false {
                                    var v48 int64 = 42
                                    _ = v48
                                } else {
                                    v49 := GenS2{F0: (42 < 162), F1: 42}
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    _ = v49
                                }
                            } else {
                                fmt.Println(v4)
                                v50 := GenS2{F0: (255 == 1), F1: 128}
                                var v51 int = fi47
                                _ = v50
                                _ = v51
                            }
                            _ = fi47
                        }
                        v2 -= (v13 << ((42) & 63))
                        for fi52 := 0; fi52 < 13; fi52++ {
                            var v53 uint64 = (v2 & (v13 | (v13 / ((18446744073709551615) | 1))))
                            var v54 int8 = (v14 + (v14 / (((v14 | (-(v14)))) | 1)))
                            if (((v3 - v3) / v3) >= v3) {
                                for fi55 := 0; fi55 < 6; fi55++ {
                                    var v56 float32 = ((1e10 * 0.1) + 0.1)
                                    if len(v1) > 0 {
                                        v1[0] = ("" + "foo bar")
                                    }
                                    _ = fi55
                                    _ = v56
                                }
                                v57 := GenS4{F0: v14, F1: 1, F2: (fi52 ^ (fi52 % (((^(fi52))) | 1)))}
                                _ = v57
                            }
                            _ = fi52
                            _ = v53
                            _ = v54
                        }
                    }
                }
            }
            if (!(271 != 1)) {
                var v58 bool = false
                if v58 {
                    var v59 uint16 = 0
                    fmt.Println(v4, v3, v4)
                    if ("\n\t" < "hello") {
                        var v60 uint8 = 2
                        fmt.Println(v13, v58, v58)
                        _ = v60
                    }
                    _ = v59
                } else {
                    v2 += v13
                }
                for fi61 := 0; fi61 < 15; fi61++ {
                    v58 = true
                    switch fi61 {
                    case 1:
                        if false {
                            fmt.Println(v3)
                            var v62 uint64 = 0
                            if v58 {
                                var v63 bool = ((-(v14)) != -1)
                                v5 = append(v5, "foo bar")
                                _ = v63
                            } else {
                                if v58 {
                                    v64 := GenS3{F0: 32767, F1: (v13 * (^(v13)))}
                                    _ = v64
                                } else {
                                    var v65 int = 42
                                    switch fi61 {
                                    case 127:
                                        if len(v1) > 0 {
                                            v1[0] = ("\n\t" + ("hello" + ("\n\t" + "a")))
                                        }
                                    case 0:
                                        if (1 < v2) {
                                            if (v58 || (true || true)) {
                                                for fi66 := 0; fi66 < 10; fi66++ {
                                                    var v67 uint = 65535
                                                    v3 = v3
                                                    _ = fi66
                                                    _ = v67
                                                }
                                                if false {
                                                    if (!(false && (1 != 1))) {
                                                        v68 := GenS5{F0: 758, F1: 0}
                                                        switch int(v14) {
                                                        case 248:
                                                            fmt.Println(v14)
                                                        case 4294967295:
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                        case 32767:
                                                            if (!(-9223372036854775808 == -1)) {
                                                                for fi69 := 0; fi69 < 9; fi69++ {
                                                                    v58 = v58
                                                                    if len(v5) > 0 {
                                                                        v5[0] = "hello"
                                                                    }
                                                                    var v70 int8 = -115
                                                                    _ = fi69
                                                                    _ = v70
                                                                }
                                                                v5 = append(v5, ("\n\t" + "a"))
                                                            } else {
                                                                switch (v65 / (((fi61 | v65)) | 1)) {
                                                                case 382:
                                                                    var v71 uint8 = 116
                                                                    _ = v71
                                                                case -9223372036854775808:
                                                                    v65 = fi61
                                                                default:
                                                                    if (!(!(v14 <= 120))) {
                                                                        if (!v58) {
                                                                            var v72 float32 = 1.5
                                                                            var v73 uint16 = 32767
                                                                            if len(v1) >= 2 {
                                                                                v1 = v1[1:len(v1)]
                                                                            }
                                                                            _ = v72
                                                                            _ = v73
                                                                        } else {
                                                                            var v74 uint8 = 88
                                                                            _ = v74
                                                                        }
                                                                        if len(v1) > 0 {
                                                                            v1[0] = "\n\t"
                                                                        }
                                                                    } else {
                                                                        v65 = int(v14)
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        _ = v68
                                                    } else {
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                        _ = len(v5)
                                                    }
                                                }
                                                for fi75 := 0; fi75 < 19; fi75++ {
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                    _ = fi75
                                                }
                                            }
                                        } else {
                                            fmt.Println(v14)
                                        }
                                    case 600:
                                        switch fi61 {
                                        case -618:
                                            fmt.Println(v65)
                                        case 1:
                                            _ = len(v1)
                                        case 128:
                                            if true {
                                                var v76 int16 = -1
                                                _ = v76
                                            } else {
                                                v65 -= (fi61 + (v65 / ((fi61) | 1)))
                                                v77 := GenS2{F0: v58, F1: 255}
                                                _ = v77
                                            }
                                        }
                                    }
                                    _ = len(v5)
                                    _ = v65
                                }
                            }
                            _ = v62
                        }
                    case -9223372036854775808:
                        if false {
                            _ = len(v5)
                            for fi78 := 0; fi78 < 11; fi78++ {
                                if v58 {
                                    v13 += v13
                                } else {
                                    v79 := GenS4{F0: v14, F1: 113, F2: (fi61 | (-(fi78)))}
                                    if (!v58) {
                                        v13 -= 4294967295
                                        if v58 {
                                            v4 += v4
                                            v80 := GenS1{F0: 2, F1: 0}
                                            if (-2.25 >= 1e10) {
                                                fmt.Println(v13, v58)
                                                _ = len(v5)
                                                fmt.Println(v2)
                                            } else {
                                                if (!v58) {
                                                    v81 := GenS6{F0: 689, F1: 954, F2: v58}
                                                    v13 -= v13
                                                    fmt.Println(fi78, v2, v2)
                                                    _ = v81
                                                }
                                                v2 = (v13 % (((v2 % ((18446744073709551615) | 1))) | 1))
                                                v4 = (v4 * (v4 / ((-2147483648) | 1)))
                                            }
                                            _ = v80
                                        }
                                        switch -2147483648 {
                                        case 777:
                                            var v82 int = -453
                                            _ = v82
                                        default:
                                            if true {
                                                var v83 float32 = 0.0
                                                if len(v1) > 0 {
                                                    v1[0] = "a"
                                                }
                                                _ = v83
                                            } else {
                                                v58 = v58
                                            }
                                        }
                                    }
                                    switch "foo bar" {
                                    case "foo bar":
                                        var v84 int = (fi61 << ((255) & 63))
                                        _ = v84
                                    case "":
                                        switch fi78 {
                                        case 2147483647:
                                            fmt.Println(v4, v4, v2)
                                        default:
                                            v85 := GenS3{F0: 32767, F1: 911}
                                            _ = v85
                                        }
                                    default:
                                        v13 = (v2 >> ((1) & 63))
                                    }
                                    v86 := GenS3{F0: -755, F1: (^(v13))}
                                    _ = v79
                                    _ = v86
                                }
                                if v58 {
                                    var v87 uint32 = 0
                                    switch fi61 {
                                    case 32767:
                                        var v88 uint8 = 128
                                        _ = v88
                                    case -9223372036854775808:
                                        v13 -= v2
                                    }
                                    var v89 float64 = 0.0
                                    _ = v87
                                    _ = v89
                                }
                                v14 = -1
                                _ = fi78
                            }
                            var v90 uint16 = 178
                            _ = v90
                        } else {
                            v2 -= v2
                            var v91 uint16 = uint16(v2)
                            switch ("" + "foo bar") {
                            case "a":
                                if false {
                                    v2 -= v13
                                    v3 -= 0.0
                                } else {
                                    switch -9223372036854775808 {
                                    case 9223372036854775807:
                                        _ = len(v5)
                                    case -547:
                                        v92 := GenS7{F0: uint8(v91)}
                                        _ = v92
                                    default:
                                        v93 := GenS2{F0: (v58 && v58), F1: 127}
                                        _ = v93
                                    }
                                }
                            case "foo bar":
                                v94 := GenS7{F0: 2}
                                _ = v94
                            case "":
                                var v95 bool = (!v58)
                                _ = v95
                            default:
                                var v96 int = fi61
                                _ = v96
                            }
                            _ = v91
                        }
                    case 32767:
                        v97 := GenS8{F0: 2147483647}
                        _ = v97
                    }
                    _ = fi61
                }
                _ = v58
            } else {
                switch "hello" {
                case "\n\t":
                    _ = len(v5)
                case "hello":
                    fmt.Println(v2, v14)
                case "foo bar":
                    switch "a" {
                    case "foo bar":
                        v13 -= v13
                    case "\n\t":
                        fmt.Println(v3)
                    }
                default:
                    fmt.Println(v14)
                }
                _ = len(v5)
                var v98 bool = true
                _ = v98
            }
            if (true && (!true)) {
                var v99 bool = ((true && ((v3 < v3) && true)) || true)
                v100 := GenS9{F0: 255}
                switch 965 {
                case -453:
                    v101 := GenS6{F0: 42, F1: -9223372036854775808, F2: v99}
                    _ = v101
                case -9223372036854775808:
                    for fi102 := 0; fi102 < 10; fi102++ {
                        v2 -= (v13 * v2)
                        _ = fi102
                    }
                case -1:
                    v1 = append(v1, "hello")
                }
                _ = v100
                _ = v99
            } else {
                if true {
                    for fi103 := 0; fi103 < 8; fi103++ {
                        if false {
                            var v104 float64 = v3
                            _ = v104
                        }
                        switch "a" {
                        case "a":
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        case "hello":
                            var v105 int32 = v4
                            _ = v105
                        case "":
                            v106 := GenS6{F0: 2147483647, F1: -838, F2: true}
                            _ = v106
                        }
                        for fi107 := 0; fi107 < 18; fi107++ {
                            var v108 uint32 = 1
                            v1 = append(v1, "a")
                            fmt.Println(v108, v14, v13)
                            _ = fi107
                            _ = v108
                        }
                        _ = fi103
                    }
                    v14 -= -56
                }
            }
        }
    }
    switch 42 {
    case 4294967295:
        fmt.Println(v4, v4)
    case 42:
        fmt.Println(v14, v3, v2)
    case 87:
        if (false || false) {
            for fi109 := 0; fi109 < 5; fi109++ {
                v2 = 18446744073709551615
                if (v4 > 0) {
                    switch (fi109 << ((18446744073709551615) & 63)) {
                    case 128:
                        var v110 uint64 = (v13 + v13)
                        _ = v110
                    case 903:
                        for fi111 := 0; fi111 < 18; fi111++ {
                            fmt.Println(v3, v14)
                            v2 -= (v2 | (v2 | v2))
                            v14 -= -128
                            _ = fi111
                        }
                    case 4294967295:
                        if false {
                            if true {
                                v112 := GenS10{F0: 820, F1: 42}
                                _ = v112
                            }
                            v113 := GenS11{F0: fi109, F1: 42, F2: (^(fi109))}
                            _ = v113
                        } else {
                            if false {
                                v13 = v2
                            }
                        }
                    }
                    fmt.Println(v4)
                    for fi114 := 0; fi114 < 13; fi114++ {
                        var v115 string = "hello"
                        _ = fi114
                        _ = v115
                    }
                }
                _ = fi109
            }
        } else {
            v5 = append(v5, "hello")
            if true {
                fmt.Println(v14, v13)
                for fi116 := 0; fi116 < 1; fi116++ {
                    switch "\n\t" {
                    case "":
                        var v117 int8 = (v14 % ((int8(v4)) | 1))
                        _ = v117
                    }
                    if (uint8(v13) <= 1) {
                        fmt.Println(fi116, v14, v3)
                    } else {
                        v118 := GenS9{F0: 592}
                        v4 = int32(fi116)
                        _ = v118
                    }
                    _ = fi116
                }
            } else {
                switch "\n\t" {
                case "hello":
                    v3 = v3
                default:
                    var v119 uint32 = 65535
                    _ = v119
                }
                v120 := GenS4{F0: (v14 % ((-1) | 1)), F1: 0, F2: 9223372036854775807}
                _ = v120
            }
        }
    default:
        v2 += (v2 * 5)
    }
    for fi121 := 0; fi121 < 4; fi121++ {
        var v122 int8 = (v14 << ((42) & 63))
        v123 := GenS2{F0: false, F1: 0}
        _ = fi121
        _ = v122
        _ = v123
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
