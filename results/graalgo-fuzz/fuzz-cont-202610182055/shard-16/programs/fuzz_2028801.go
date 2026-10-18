package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 int8 }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int16; F1 int32 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint8; F1 uint }
type GenS10 struct { F0 int64; F1 uint }
type GenS11 struct { F0 int32; F1 bool; F2 bool }
type GenS12 struct { F0 uint32; F1 string; F2 bool }
type GenS13 struct { F0 uint64; F1 int8 }
type GenS14 struct { F0 uint; F1 int64; F2 uint }
type GenS15 struct { F0 int16; F1 int }
type GenS16 struct { F0 uint }
type GenS17 struct { F0 uint; F1 int; F2 int8 }
type GenS18 struct { F0 int; F1 string; F2 int16 }
type GenS19 struct { F0 int; F1 int; F2 uint }
type GenS20 struct { F0 bool; F1 uint; F2 bool }
type GenS21 struct { F0 int32; F1 int8; F2 uint32 }

func main() {
    var v1 uint32 = 1
    var v2 float32 = 1.5
    var v3 uint8 = 177
    for fi4 := 0; fi4 < 6; fi4++ {
        v3 = (v3 >> ((745) & 63))
        _ = fi4
    }
    var v5 float64 = (-2.25 - -2.25)
    v6 := []int{-2147483648}
    v2 -= v2
    v1 -= 2147483647
    v7 := GenS1{F0: v3, F1: -2147483648, F2: -1}
    v8 := GenS1{F0: (^(v3)), F1: -926, F2: 127}
    if (!(115 >= 1)) {
        var v9 uint16 = 255
        _ = v9
    }
    v3 += uint8(v1)
    _ = len(v6)
    var v10 uint = 4294967295
    var v11 int8 = 1
    switch 42 {
    case 845:
        if false {
            switch "\n\t" {
            case "a":
                v6 = append(v6, 532)
            case "foo bar":
                v12 := GenS2{F0: (true || (!false)), F1: v3}
                _ = v12
            }
            var v13 uint16 = 455
            v13 -= (v13 | (v13 + (v13 >> ((v10) & 63))))
            _ = v13
        }
    default:
        if (v1 < v1) {
            switch -2147483648 {
            case 42:
                switch ("\n\t" + ("hello" + "\n\t")) {
                case "\n\t":
                    for fi14 := 0; fi14 < 2; fi14++ {
                        switch (-(fi14)) {
                        case -1:
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        case 128:
                            if true {
                                var v15 float32 = (v2 + (v2 + ((v2 + v2) * v2)))
                                var v16 uint = v10
                                _ = v15
                                _ = v16
                            }
                        default:
                            v5 = v5
                        }
                        var v17 float64 = 0.0
                        v17 -= ((1.5 * v5) + float64(v2))
                        _ = fi14
                        _ = v17
                    }
                case "a":
                    if len(v6) > 0 {
                        v6[0] = 148
                    }
                }
            case 128:
                for fi18 := 0; fi18 < 17; fi18++ {
                    var v19 float32 = v2
                    _ = fi18
                    _ = v19
                }
            default:
                if (float64((0.0 * ((((3.14159 - v2) - v2) - 0.1) + v2))) > v5) {
                    var v20 float64 = 0.0
                    if ((false && (!(1 != 80))) && false) {
                        v21 := GenS1{F0: (v3 | 127), F1: -401, F2: (v11 - (v11 & v11))}
                        _ = v21
                    } else {
                        if (1e10 >= float32(0.0)) {
                            v22 := GenS2{F0: true, F1: v3}
                            _ = v22
                        } else {
                            for fi23 := 0; fi23 < 9; fi23++ {
                                for fi24 := 0; fi24 < 2; fi24++ {
                                    v11 -= v11
                                    _ = fi24
                                }
                                v6 = append(v6, fi23)
                                var v25 uint16 = 65535
                                _ = fi23
                                _ = v25
                            }
                            var v26 int8 = v11
                            if true {
                                switch ("foo bar" + ("" + ("hello" + "foo bar"))) {
                                case "":
                                    for fi27 := 0; fi27 < 10; fi27++ {
                                        for fi28 := 0; fi28 < 8; fi28++ {
                                            fmt.Println(v5, v11)
                                            _ = fi28
                                        }
                                        v29 := GenS2{F0: false, F1: 127}
                                        _ = fi27
                                        _ = v29
                                    }
                                case "a":
                                    for fi30 := 0; fi30 < 16; fi30++ {
                                        if len(v6) > 0 {
                                            v6[0] = fi30
                                        }
                                        v11 -= v11
                                        _ = fi30
                                    }
                                default:
                                    switch 4294967295 {
                                    case 32767:
                                        switch "a" {
                                        case "foo bar":
                                            if (v2 >= 3.14159) {
                                                var v31 string = "\n\t"
                                                var v32 uint16 = 0
                                                for fi33 := 0; fi33 < 4; fi33++ {
                                                    v34 := GenS3{F0: (v10 ^ 558)}
                                                    switch v31 {
                                                    case "\n\t":
                                                        fmt.Println(v1)
                                                    }
                                                    if (!false) {
                                                        for fi35 := 0; fi35 < 7; fi35++ {
                                                            v36 := GenS4{F0: 1, F1: -1}
                                                            _ = fi35
                                                            _ = v36
                                                        }
                                                    } else {
                                                        v37 := GenS5{F0: (v1 | (v1 & v1))}
                                                        var v38 uint = (v10 + v10)
                                                        v39 := GenS6{F0: (^(v11))}
                                                        _ = v37
                                                        _ = v38
                                                        _ = v39
                                                    }
                                                    _ = fi33
                                                    _ = v34
                                                }
                                                _ = v31
                                                _ = v32
                                            }
                                        default:
                                            var v40 int32 = -96
                                            _ = v40
                                        }
                                    default:
                                        var v41 bool = false
                                        _ = v41
                                    }
                                }
                                if len(v6) > 0 {
                                    v6[0] = -9223372036854775808
                                }
                                v42 := GenS7{F0: (v1 - (v1 - (v1 >> ((42) & 63))))}
                                _ = v42
                            } else {
                                v3 = (v3 >> ((v10) & 63))
                                var v43 uint64 = 0
                                _ = v43
                            }
                            _ = v26
                        }
                        v44 := GenS5{F0: 2147483647}
                        _ = v44
                    }
                    _ = v20
                } else {
                    var v45 int16 = 1000
                    _ = v45
                }
            }
        } else {
            if (true || true) {
                var v46 int64 = 42
                v47 := GenS8{F0: v3}
                _ = v46
                _ = v47
            } else {
                for fi48 := 0; fi48 < 7; fi48++ {
                    if (!(!true)) {
                        if (true && true) {
                            var v49 int = (fi48 + int(v10))
                            v50 := GenS4{F0: 42, F1: 42}
                            if false {
                                if ((v49 - (fi48 & v49)) > 9223372036854775807) {
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    for fi51 := 0; fi51 < 4; fi51++ {
                                        if false {
                                            v10 -= (^(v10))
                                            var v52 int64 = -9223372036854775808
                                            v52 -= (v52 | v52)
                                            _ = v52
                                        } else {
                                            if true {
                                                v53 := GenS3{F0: v10}
                                                var v54 uint32 = (^(v1))
                                                v5 -= v5
                                                _ = v53
                                                _ = v54
                                            }
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                            if ((^(v10)) <= 976) {
                                                if len(v6) > 0 {
                                                    v6[0] = fi51
                                                }
                                                v6 = append(v6, 9223372036854775807)
                                            }
                                            switch (fi48 & fi51) {
                                            case 128:
                                                if false {
                                                    _ = len(v6)
                                                }
                                            case 4294967295:
                                                v2 = 1.5
                                            default:
                                                if len(v6) > 0 {
                                                    v6[0] = (fi48 * (fi51 & v49))
                                                }
                                            }
                                            v55 := GenS9{F0: 13, F1: v10}
                                            var v56 float64 = (float64(0.1) * v5)
                                            _ = v55
                                            _ = v56
                                        }
                                        var v57 string = "foo bar"
                                        _ = fi51
                                        _ = v57
                                    }
                                    var v58 string = "\n\t"
                                    _ = v58
                                } else {
                                    var v59 uint16 = 42
                                    _ = len(v6)
                                    _ = v59
                                }
                                fmt.Println(v11, fi48, v2)
                                var v60 uint = v10
                                _ = v60
                            } else {
                                var v61 int8 = (v11 & -128)
                                _ = v61
                            }
                            _ = v49
                            _ = v50
                        }
                    }
                    v11 -= (^(v11))
                    _ = fi48
                }
                switch "\n\t" {
                case "a":
                    var v62 int = 9223372036854775807
                    _ = v62
                case "foo bar":
                    var v63 uint = 4294967295
                    _ = v63
                default:
                    switch (("hello" + "foo bar") + "hello") {
                    case "\n\t":
                        var v64 float32 = v2
                        _ = v64
                    case "a":
                        switch "" {
                        case "a":
                            if len(v6) > 0 {
                                v6[0] = 4294967295
                            }
                        case "\n\t":
                            v65 := GenS6{F0: (v11 ^ 0)}
                            _ = v65
                        case "":
                            _ = len(v6)
                        }
                    case "hello":
                        if len(v6) > 0 {
                            v6[0] = 118
                        }
                    default:
                        fmt.Println(v2)
                    }
                }
                switch 0 {
                case 0:
                    for fi66 := 0; fi66 < 19; fi66++ {
                        for fi67 := 0; fi67 < 19; fi67++ {
                            v68 := GenS10{F0: 1, F1: (^(v10))}
                            if true {
                                var v69 string = "a"
                                if false {
                                    var v70 uint64 = 519
                                    _ = v70
                                } else {
                                    v71 := GenS5{F0: v1}
                                    v72 := GenS11{F0: -2147483648, F1: true, F2: true}
                                    v5 = (-2.25 * ((1e10 + (1e10 - (3.14159 - 1e10))) * 1e10))
                                    _ = v71
                                    _ = v72
                                }
                                switch v69 {
                                case "a":
                                    v1 -= v1
                                default:
                                    v73 := GenS12{F0: (v1 >> ((42) & 63)), F1: "", F2: false}
                                    _ = v73
                                }
                                switch "" {
                                case "":
                                    _ = len(v6)
                                case "a":
                                    switch ("" + (("a" + ((((((v69 + v69) + "foo bar") + ((("a" + v69) + v69) + ((v69 + ("foo bar" + "hello")) + (v69 + ("hello" + ("a" + v69)))))) + v69) + v69) + (v69 + (v69 + v69)))) + v69)) {
                                    case "hello":
                                        if (((v2 + float32((v5 + 3.14159))) > v2) || (345 <= 255)) {
                                            var v74 float32 = v2
                                            if false {
                                                var v75 int8 = (^(v11))
                                                v2 = v74
                                                switch -351 {
                                                case 32767:
                                                    var v76 bool = (2147483647 <= 2147483647)
                                                    _ = v76
                                                default:
                                                    v77 := GenS9{F0: v3, F1: v10}
                                                    _ = v77
                                                }
                                                _ = v75
                                            } else {
                                                var v78 int = -278
                                                var v79 int = fi67
                                                _ = v78
                                                _ = v79
                                            }
                                            var v80 uint = 268
                                            _ = v74
                                            _ = v80
                                        } else {
                                            v3 -= (v3 & 9)
                                            v69 = v69
                                        }
                                    case "\n\t":
                                        if len(v6) > 0 {
                                            v6[0] = (fi67 - (fi66 | -838))
                                        }
                                    case "":
                                        switch v69 {
                                        case "":
                                            fmt.Println(fi67, fi66, v69)
                                        case "hello":
                                            var v81 uint8 = v3
                                            _ = v81
                                        default:
                                            if false {
                                                if false {
                                                    if false {
                                                        if false {
                                                            switch (fi66 & fi67) {
                                                            case -9223372036854775808:
                                                                v3 += v3
                                                            case 1:
                                                                fmt.Println(v11)
                                                            }
                                                            if true {
                                                                var v82 int16 = 32767
                                                                _ = v82
                                                            }
                                                            v69 = v69
                                                            v83 := GenS13{F0: 151, F1: (v11 << ((v10) & 63))}
                                                            _ = v83
                                                        } else {
                                                            fmt.Println(v10, v2)
                                                            if ((false && (42 < 993)) && false) {
                                                                if len(v6) > 0 {
                                                                    v6[0] = (^(fi66))
                                                                }
                                                                v5 += v5
                                                                var v84 uint = 65535
                                                                v10 -= (v10 << (((v10 + v84)) & 63))
                                                                v85 := GenS14{F0: v84, F1: 9223372036854775807, F2: v10}
                                                                _ = v84
                                                                _ = v85
                                                            } else {
                                                                var v86 uint = v10
                                                                var v87 int32 = 1
                                                                _ = v86
                                                                _ = v87
                                                            }
                                                        }
                                                        var v88 int = (fi67 / (((fi66 | int(v3))) | 1))
                                                        fmt.Println(v69)
                                                        _ = v88
                                                    } else {
                                                        v10 -= (v10 - v10)
                                                        var v89 uint8 = (v3 & (v3 | (v3 % ((v3) | 1))))
                                                        _ = v89
                                                    }
                                                    v5 = 1e10
                                                } else {
                                                    var v90 bool = (!true)
                                                    v3 = (^(v3))
                                                    var v91 int = (fi67 + fi67)
                                                    _ = v90
                                                    _ = v91
                                                }
                                                switch ("\n\t" + v69) {
                                                case "":
                                                    v92 := GenS8{F0: v3}
                                                    _ = v92
                                                case "\n\t":
                                                    var v93 bool = false
                                                    _ = v93
                                                case "a":
                                                    v6 = append(v6, (^(fi66)))
                                                default:
                                                    if (1 > 0) {
                                                        v94 := GenS11{F0: 0, F1: false, F2: true}
                                                        v95 := GenS15{F0: 393, F1: 32768}
                                                        _ = v94
                                                        _ = v95
                                                    }
                                                }
                                                var v96 int64 = 9223372036854775807
                                                _ = v96
                                            } else {
                                                var v97 uint16 = 0
                                                v5 = (float64(v2) / 1.0)
                                                if len(v6) > 0 {
                                                    v6[0] = (fi66 - (^(fi66)))
                                                }
                                                _ = v97
                                            }
                                        }
                                    default:
                                        v1 = v1
                                    }
                                case "foo bar":
                                    switch v69 {
                                    case "a":
                                        switch "a" {
                                        case "\n\t":
                                            if ((v10 * v10) != uint(v1)) {
                                                if (((v3 % ((v3) | 1)) <= v3) && ("" < v69)) {
                                                    if len(v6) > 0 {
                                                        v6[0] = fi66
                                                    }
                                                }
                                            }
                                            v98 := GenS16{F0: 558}
                                            _ = v98
                                        default:
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                            v10 -= (v10 & (v10 | 591))
                                        }
                                    case "foo bar":
                                        if ((!false) && true) {
                                            switch 4294967295 {
                                            case 127:
                                                var v99 int32 = -2147483648
                                                _ = v99
                                            default:
                                                var v100 int8 = v11
                                                _ = v100
                                            }
                                        } else {
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                            switch 32767 {
                                            case -2147483648:
                                                if false {
                                                    v101 := GenS17{F0: 1, F1: (fi66 & fi66), F2: (v11 << (((^(v10))) & 63))}
                                                    v3 = 1
                                                    _ = v101
                                                } else {
                                                    var v102 string = v69
                                                    v103 := GenS11{F0: 42, F1: false, F2: false}
                                                    v1 = (v1 >> ((0) & 63))
                                                    _ = v102
                                                    _ = v103
                                                }
                                            default:
                                                var v104 uint = uint(v3)
                                                _ = v104
                                            }
                                            v105 := GenS4{F0: 42, F1: -2147483648}
                                            v106 := GenS18{F0: (fi67 << (((v10 >> ((v10) & 63))) & 63)), F1: "foo bar", F2: 366}
                                            _ = v105
                                            _ = v106
                                        }
                                    }
                                }
                                v6 = append(v6, (fi67 + fi67))
                                v10 += v10
                                _ = v69
                            }
                            v107 := GenS19{F0: fi66, F1: (fi67 % ((0) | 1)), F2: uint(v1)}
                            _ = fi67
                            _ = v107
                            _ = v68
                        }
                        v3 += uint8(v1)
                        _ = fi66
                    }
                case 295:
                    if false {
                        if true {
                            switch 1 {
                            case -594:
                                for fi108 := 0; fi108 < 10; fi108++ {
                                    v11 = (v11 * -1)
                                    _ = fi108
                                }
                            case 9223372036854775807:
                                switch ("a" + ("" + ("a" + "hello"))) {
                                case "foo bar":
                                    v11 -= -4
                                default:
                                    var v109 bool = (727 <= 939)
                                    _ = v109
                                }
                            }
                        }
                    } else {
                        if false {
                            var v110 uint16 = 572
                            var v111 float32 = ((v2 - (-2.25 + (((((3.14159 * ((float32(0.1) + 3.14159) * (float32(((1.5 / 1.0) - (0.1 + (3.14159 / v5)))) - v2))) / ((v2 - float32((1.5 + v5))) - v2)) - v2) / (-2.25 - 1.5)) + -2.25))) * v2)
                            _ = v110
                            _ = v111
                        }
                        v1 -= 691
                    }
                case 4294967295:
                    if (v5 <= v5) {
                        fmt.Println(v2, v1)
                        if false {
                            switch ("\n\t" + "\n\t") {
                            case "foo bar":
                                for fi112 := 0; fi112 < 4; fi112++ {
                                    var v113 uint32 = 907
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    v3 = (^(v3))
                                    _ = fi112
                                    _ = v113
                                }
                            case "a":
                                fmt.Println(v10)
                            default:
                                for fi114 := 0; fi114 < 13; fi114++ {
                                    v10 = (v10 >> ((v10) & 63))
                                    switch (fi114 % (((fi114 - 128)) | 1)) {
                                    case 32768:
                                        v115 := GenS20{F0: (42 < 273), F1: (v10 + 748), F2: false}
                                        _ = v115
                                    }
                                    _ = fi114
                                }
                            }
                        } else {
                            if false {
                                v1 = (v1 % ((v1) | 1))
                            } else {
                                var v116 uint32 = (v1 << (((v10 & v10)) & 63))
                                _ = v116
                            }
                            var v117 int8 = (v11 * (v11 & int8(v10)))
                            var v118 string = "\n\t"
                            _ = v117
                            _ = v118
                        }
                    } else {
                        fmt.Println(v1)
                        for fi119 := 0; fi119 < 3; fi119++ {
                            switch "a" {
                            case "\n\t":
                                if false {
                                    v5 = (v5 + v5)
                                    for fi120 := 0; fi120 < 13; fi120++ {
                                        switch "" {
                                        case "\n\t":
                                            if ("\n\t" >= ("" + "")) {
                                                var v121 string = "\n\t"
                                                _ = v121
                                            } else {
                                                v122 := GenS5{F0: (v1 * 1)}
                                                v1 = v1
                                                var v123 int32 = 1
                                                _ = v122
                                                _ = v123
                                            }
                                        case "":
                                            v2 -= ((3.14159 + ((float32(v5) + (v2 + float32(v5))) * (v2 + ((float32(((0.0 * (1.5 * (((3.14159 * v5) * float64(v2)) - 1e10))) / (v5 - float64((v2 + v2))))) + v2) * (((1e10 - v2) * float32((v5 + v5))) + float32(((float64(3.14159) * 0.0) - (0.0 * v5)))))))) - (v2 + v2))
                                        case "hello":
                                            v6 = append(v6, (fi119 << ((42) & 63)))
                                        }
                                        switch ("a" + "") {
                                        case "":
                                            fmt.Println(v5)
                                        case "hello":
                                            fmt.Println(v3, fi120)
                                        case "a":
                                            var v124 int32 = 2147483647
                                            _ = v124
                                        default:
                                            if false {
                                                if (v11 != int8(v1)) {
                                                    switch "foo bar" {
                                                    case "a":
                                                        v2 += v2
                                                    case "foo bar":
                                                        var v125 uint64 = 192
                                                        _ = v125
                                                    default:
                                                        fmt.Println(v2, v3)
                                                    }
                                                    var v126 uint32 = (v1 << ((v10) & 63))
                                                    _ = v126
                                                } else {
                                                    v10 = (^(v10))
                                                }
                                                if true {
                                                    v127 := GenS13{F0: 42, F1: 74}
                                                    if true {
                                                        _ = len(v6)
                                                        var v128 bool = true
                                                        _ = len(v6)
                                                        v10 = v10
                                                        _ = v128
                                                    } else {
                                                        var v129 uint = 101
                                                        v2 += 1.5
                                                        _ = v129
                                                    }
                                                    fmt.Println(v1)
                                                    fmt.Println(fi119, v5, v3)
                                                    _ = v127
                                                }
                                                v5 -= 0.0
                                                if false {
                                                    fmt.Println(v11, v1, v3)
                                                } else {
                                                    fmt.Println(v3, v1, fi119)
                                                    fmt.Println(v10, fi119)
                                                    fmt.Println(fi119, fi120)
                                                }
                                            } else {
                                                v130 := GenS21{F0: 2147483647, F1: 1, F2: v1}
                                                v2 -= (v2 - (((1e10 - v2) + (((float32(float64(v2)) - (v2 + float32(float64(float32((float64(v2) * (float64(1.5) + v5))))))) / (float32(v5) - (-2.25 + ((float32(v5) * float32(float64((float32(0.0) - (v2 - 3.14159))))) - (v2 / 0.1))))) - (1.5 - float32(1e10)))) + (-2.25 + (float32(v5) - ((1e10 * v2) + v2)))))
                                                if (v3 > 127) {
                                                    var v131 int16 = 0
                                                    v10 = v10
                                                    _ = v131
                                                } else {
                                                    var v132 uint16 = 97
                                                    v6 = append(v6, (fi120 - 4294967295))
                                                    v1 -= v1
                                                    _ = v132
                                                }
                                                _ = v130
                                            }
                                        }
                                        _ = fi120
                                    }
                                } else {
                                    v2 = v2
                                    for fi133 := 0; fi133 < 9; fi133++ {
                                        if (true && true) {
                                            if false {
                                                v10 = (^(v10))
                                            }
                                            v1 -= (v1 % (((^(v1))) | 1))
                                        }
                                        v6 = append(v6, 9223372036854775807)
                                        fmt.Println(v11, fi119)
                                        _ = fi133
                                    }
                                }
                            default:
                                v134 := GenS9{F0: v3, F1: v10}
                                _ = v134
                            }
                            _ = fi119
                        }
                    }
                }
            }
            switch "foo bar" {
            case "\n\t":
                v6 = append(v6, -9223372036854775808)
            case "":
                var v135 int32 = 1
                _ = v135
            }
        }
    }
    switch "a" {
    case "":
        var v136 uint16 = 65535
        _ = v136
    default:
        v3 = v3
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
