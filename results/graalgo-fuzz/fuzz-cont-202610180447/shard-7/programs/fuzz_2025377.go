package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 int32; F1 bool; F2 uint32 }
type GenS3 struct { F0 uint16; F1 int8 }
type GenS4 struct { F0 uint32; F1 int16; F2 int8 }
type GenS5 struct { F0 uint32; F1 uint8 }
type GenS6 struct { F0 bool; F1 uint; F2 uint8 }
type GenS7 struct { F0 uint8; F1 int32; F2 string }
type GenS8 struct { F0 string; F1 int16; F2 int8 }
type GenS9 struct { F0 string; F1 int32 }
type GenS10 struct { F0 string; F1 string; F2 int32 }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 uint32; F1 int }

func main() {
    var v1 bool = (79 >= 1)
    var v2 int32 = -18
    var v3 int64 = -535
    switch (("foo bar" + "hello") + "hello") {
    case "foo bar":
        switch -168 {
        case 61:
            v4 := GenS1{F0: v3, F1: 0}
            _ = v4
        case 42:
            switch -731 {
            case 603:
                var v5 int8 = -1
                _ = v5
            case -9223372036854775808:
                switch "" {
                case "foo bar":
                    v6 := []int{-1, -21, 32767}
                    _ = v6
                default:
                    v7 := []int{9223372036854775807, 4294967295, 9223372036854775807}
                    _ = v7
                }
            case 4294967295:
                switch 25 {
                case 32767:
                    v2 += v2
                case -969:
                    v8 := []int{38}
                    _ = v8
                case 2147483647:
                    v2 -= -299
                default:
                    for fi9 := 0; fi9 < 13; fi9++ {
                        v10 := []int{fi9, fi9}
                        if (0 != 728) {
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                            fmt.Println(v2, v1)
                            if (731 <= (v2 % ((v2) | 1))) {
                                switch (("hello" + "") + "hello") {
                                case "a":
                                    switch fi9 {
                                    case 128:
                                        v2 = v2
                                    case -866:
                                        fmt.Println(v3)
                                    case -1:
                                        v3 = 1
                                    default:
                                        if false {
                                            if (1 == 234) {
                                                switch (("foo bar" + "hello") + "") {
                                                case "hello":
                                                    v3 -= v3
                                                case "":
                                                    if false {
                                                        if (((!(!true)) || (1 >= 251)) && (v1 && v1)) {
                                                            var v11 int64 = v3
                                                            fmt.Println(fi9, v11)
                                                            _ = v11
                                                        }
                                                        v12 := GenS1{F0: v3, F1: 42}
                                                        _ = v12
                                                    }
                                                }
                                            }
                                            fmt.Println(v3)
                                            v13 := GenS2{F0: (-(v2)), F1: v1, F2: 2147483647}
                                            _ = v13
                                        } else {
                                            v1 = (-48 > -1)
                                            v2 = v2
                                            if (v1 && true) {
                                                switch (("" + "a") + (("foo bar" + ((("\n\t" + "\n\t") + "") + ("foo bar" + "a"))) + ("hello" + ""))) {
                                                case "\n\t":
                                                    for fi14 := 0; fi14 < 19; fi14++ {
                                                        switch ("foo bar" + "hello") {
                                                        case "hello":
                                                            v2 -= 957
                                                        case "a":
                                                            if (0.1 > -2.25) {
                                                                v1 = true
                                                                switch ("" + ("hello" + ("\n\t" + "\n\t"))) {
                                                                case "":
                                                                    v15 := GenS2{F0: (v2 / (((-(v2))) | 1)), F1: v1, F2: 42}
                                                                    _ = v15
                                                                case "foo bar":
                                                                    v1 = false
                                                                }
                                                                var v16 uint = 1
                                                                _ = v16
                                                            }
                                                        }
                                                        fmt.Println(fi9)
                                                        var v17 int8 = 60
                                                        _ = fi14
                                                        _ = v17
                                                    }
                                                default:
                                                    var v18 uint = 18446744073709551615
                                                    _ = v18
                                                }
                                            }
                                        }
                                    }
                                }
                            }
                        }
                        _ = fi9
                        _ = v10
                    }
                }
            }
        default:
            if (330 > 297) {
                v3 = 1
                v1 = true
                if v1 {
                    for fi19 := 0; fi19 < 8; fi19++ {
                        var v20 int16 = 952
                        for fi21 := 0; fi21 < 1; fi21++ {
                            fmt.Println(v3)
                            _ = fi21
                        }
                        _ = fi19
                        _ = v20
                    }
                    fmt.Println(v2, v1, v1)
                    for fi22 := 0; fi22 < 16; fi22++ {
                        fmt.Println(fi22, v2)
                        for fi23 := 0; fi23 < 12; fi23++ {
                            v24 := []int{fi22}
                            v2 += -2147483648
                            v3 = v3
                            _ = fi23
                            _ = v24
                        }
                        _ = fi22
                    }
                }
            } else {
                v2 = 464
                if (v3 >= -833) {
                    v25 := GenS1{F0: (v3 % (((v3 - 0)) | 1)), F1: 42}
                    _ = v25
                } else {
                    v1 = (v1 || true)
                    v26 := GenS2{F0: v2, F1: v1, F2: 2147483648}
                    _ = v26
                }
                v27 := []int{128, 128, 4294967295}
                _ = v27
            }
        }
    case "":
        v2 -= (v2 | v2)
    default:
        v2 = (v2 ^ v2)
    }
    var v28 int32 = (v2 >> ((1) & 63))
    v3 = 42
    for fi29 := 0; fi29 < 4; fi29++ {
        v30 := []int{fi29, fi29, (fi29 | fi29), -735}
        _ = fi29
        _ = v30
    }
    var v31 uint8 = 2
    v32 := []int{884}
    fmt.Println(v28, v31)
    var v33 int16 = 32767
    if v1 {
        if (("" + "a") >= ("foo bar" + "")) {
            for fi34 := 0; fi34 < 16; fi34++ {
                if (18446744073709551615 < 783) {
                    v35 := GenS1{F0: v3, F1: 101}
                    if (!(true || ((v3 - v3) > v3))) {
                        if (v1 || (0.0 == (0.0 - 0.0))) {
                            v28 -= v28
                            if len(v32) >= 2 {
                                v32 = v32[1:len(v32)]
                            }
                        } else {
                            fmt.Println(v31, v3, v2)
                            var v36 bool = (!(float64(0.1) <= 0.0))
                            _ = v36
                        }
                    } else {
                        var v37 uint64 = 1
                        if (v1 || (117 >= (v31 & (^(v31))))) {
                            v28 = v28
                            v38 := GenS3{F0: 1, F1: 127}
                            _ = v38
                        } else {
                            v1 = true
                            if (v1 || true) {
                                switch ("a" + ((("" + "\n\t") + "hello") + ("\n\t" + "\n\t"))) {
                                case "\n\t":
                                    var v39 int8 = 36
                                    _ = v39
                                case "":
                                    for fi40 := 0; fi40 < 2; fi40++ {
                                        _ = len(v32)
                                        _ = fi40
                                    }
                                case "hello":
                                    if (!((v2 & v28) >= 2147483647)) {
                                        v2 += -1
                                    }
                                }
                            }
                        }
                        _ = v37
                    }
                    v33 += (v33 - 1)
                    _ = v35
                }
                switch "\n\t" {
                case "\n\t":
                    var v41 float32 = 0.0
                    _ = v41
                }
                _ = fi34
            }
        } else {
            switch ("" + "hello") {
            case "foo bar":
                if true {
                    var v42 int = int(v31)
                    _ = v42
                }
            case "\n\t":
                var v43 int16 = (v33 / (((v33 % (((v33 - (v33 - 1))) | 1))) | 1))
                _ = v43
            case "a":
                v1 = true
            }
            v44 := GenS4{F0: 4, F1: v33, F2: 1}
            _ = v44
        }
        if (v2 >= v2) {
            v3 = (v3 - v3)
            var v45 int64 = v3
            _ = v45
        }
        if v1 {
            for fi46 := 0; fi46 < 16; fi46++ {
                var v47 uint8 = v31
                if true {
                    if v1 {
                        fmt.Println(v3, v28)
                    }
                    _ = len(v32)
                }
                v48 := GenS2{F0: v28, F1: v1, F2: 42}
                _ = fi46
                _ = v47
                _ = v48
            }
            if v1 {
                switch "hello" {
                case "foo bar":
                    v49 := GenS5{F0: 393, F1: (v31 | 2)}
                    _ = v49
                case "hello":
                    v1 = (64 >= 42)
                default:
                    fmt.Println(v3)
                }
                if (65535 > 2) {
                    var v50 string = "\n\t"
                    _ = v50
                }
                for fi51 := 0; fi51 < 8; fi51++ {
                    var v52 uint8 = v31
                    _ = fi51
                    _ = v52
                }
            } else {
                if ((v28 / ((v28) | 1)) >= (^(v2))) {
                    var v53 int32 = (v2 ^ (v2 / ((v28) | 1)))
                    switch 861 {
                    case 4294967295:
                        var v54 float64 = 1.5
                        _ = v54
                    default:
                        var v55 string = "foo bar"
                        _ = v55
                    }
                    if len(v32) > 0 {
                        v32[0] = 0
                    }
                    _ = v53
                }
            }
        } else {
            v32 = append(v32, -1)
            for fi56 := 0; fi56 < 2; fi56++ {
                for fi57 := 0; fi57 < 13; fi57++ {
                    v1 = v1
                    _ = fi57
                }
                for fi58 := 0; fi58 < 7; fi58++ {
                    if true {
                        v2 = v28
                    } else {
                        var v59 uint32 = 2147483648
                        _ = v59
                    }
                    var v60 uint = 326
                    v2 += int32(v33)
                    _ = fi58
                    _ = v60
                }
                if ((-2.25 * 3.14159) <= float32((3.14159 * 0.1))) {
                    _ = len(v32)
                }
                _ = fi56
            }
            v31 = v31
        }
    }
    v3 = v3
    var v61 uint8 = (^(v31))
    if ((18446744073709551615 == 2) && true) {
        v2 += -2147483648
    } else {
        v62 := GenS6{F0: true, F1: 532, F2: (v31 - (v31 - v31))}
        fmt.Println(v31, v61)
        _ = v62
    }
    var v63 float32 = ((float32(0.1) * (0.0 - (3.14159 / float32((0.1 / 1.0))))) / 1.0)
    fmt.Println(v28)
    if (255 <= 161) {
        fmt.Println(v2, v1, v33)
        v32 = append(v32, 4294967295)
    } else {
        if v1 {
            var v64 uint16 = 0
            if (((0.0 / v63) - (v63 + v63)) > 1.5) {
                v65 := GenS4{F0: 2147483647, F1: v33, F2: 93}
                _ = v65
            } else {
                if len(v32) >= 2 {
                    v32 = v32[1:len(v32)]
                }
                v32 = append(v32, -2147483648)
                var v66 uint = 1
                _ = v66
            }
            for fi67 := 0; fi67 < 13; fi67++ {
                v1 = ((^(v3)) < (v3 & (^(v3))))
                fmt.Println(fi67, v61, v1)
                if v1 {
                    if true {
                        for fi68 := 0; fi68 < 14; fi68++ {
                            var v69 int = (fi68 % ((fi68) | 1))
                            if v1 {
                                if len(v32) > 0 {
                                    v32[0] = 0
                                }
                                var v70 uint32 = 2147483648
                                if len(v32) > 0 {
                                    v32[0] = (^(fi67))
                                }
                                var v71 uint16 = (v64 % (((^(v64))) | 1))
                                switch ("" + "foo bar") {
                                case "\n\t":
                                    switch ((("foo bar" + (("\n\t" + "") + "")) + "a") + "foo bar") {
                                    case "\n\t":
                                        if v1 {
                                            v31 += 42
                                            fmt.Println(fi67, v71, v2)
                                            v33 = -1
                                        } else {
                                            if ((v28 ^ v28) != (^(v28))) {
                                                if v1 {
                                                    var v72 int64 = (v3 % ((v3) | 1))
                                                    fmt.Println(v61, v63, v70)
                                                    fmt.Println(v61, v72)
                                                    if v1 {
                                                        var v73 uint64 = 700
                                                        _ = v73
                                                    } else {
                                                        if (v1 || (!v1)) {
                                                            _ = len(v32)
                                                        } else {
                                                            v32 = append(v32, (v69 - fi68))
                                                        }
                                                    }
                                                    var v74 uint = 4294967295
                                                    _ = v72
                                                    _ = v74
                                                } else {
                                                    v69 -= -260
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        var v75 float64 = (3.14159 / ((float64(v63) + 1e10) - -2.25))
                                                        _ = v75
                                                    }
                                                    v76 := GenS2{F0: (v2 / ((v2) | 1)), F1: false, F2: v70}
                                                    var v77 int32 = v2
                                                    _ = v76
                                                    _ = v77
                                                }
                                                if (!(int64(fi68) == (v3 | (v3 + v3)))) {
                                                    v78 := GenS7{F0: 146, F1: (-(v2)), F2: ("\n\t" + "hello")}
                                                    v33 += 1000
                                                    var v79 uint = 42
                                                    if v1 {
                                                        fmt.Println(v70, v79)
                                                    } else {
                                                        _ = len(v32)
                                                        var v80 uint8 = 195
                                                        _ = len(v32)
                                                        _ = v80
                                                    }
                                                    _ = v78
                                                    _ = v79
                                                } else {
                                                    var v81 int = 1
                                                    _ = v81
                                                }
                                                fmt.Println(v31)
                                            } else {
                                                v82 := GenS2{F0: -545, F1: true, F2: (v70 ^ (v70 & (v70 ^ v70)))}
                                                var v83 int64 = (v3 * (v3 | (v3 / ((-1) | 1))))
                                                fmt.Println(fi68, v28, v1)
                                                _ = v82
                                                _ = v83
                                            }
                                            var v84 float32 = 0.1
                                            var v85 uint8 = 0
                                            _ = v84
                                            _ = v85
                                        }
                                    case "":
                                        switch (fi67 * v69) {
                                        case 4294967295:
                                            if (615 == (v69 << ((69) & 63))) {
                                                switch "" {
                                                case "hello":
                                                    var v86 int32 = -2147483648
                                                    _ = v86
                                                case "":
                                                    v87 := GenS8{F0: "\n\t", F1: v33, F2: 42}
                                                    _ = v87
                                                case "\n\t":
                                                    if (v63 == -2.25) {
                                                        v31 = (v61 / ((v31) | 1))
                                                    } else {
                                                        v88 := GenS3{F0: 208, F1: -34}
                                                        var v89 uint32 = (v70 % (((v70 ^ (v70 << ((526) & 63)))) | 1))
                                                        v90 := GenS3{F0: v64, F1: -128}
                                                        _ = v88
                                                        _ = v89
                                                        _ = v90
                                                    }
                                                }
                                                var v91 bool = false
                                                switch ("\n\t" + "hello") {
                                                case "\n\t":
                                                    fmt.Println(v63)
                                                case "a":
                                                    if v1 {
                                                        var v92 uint8 = (v31 / (((v61 + v31)) | 1))
                                                        v31 += uint8(v70)
                                                        _ = v92
                                                    } else {
                                                        _ = len(v32)
                                                        fmt.Println(v2, fi67, v1)
                                                        switch (fi67 + (fi67 % (((fi68 - 0)) | 1))) {
                                                        case 356:
                                                            if ((v28 + (v2 - (v28 + (v28 % ((int32(v70)) | 1))))) <= (v28 - (v28 % ((0) | 1)))) {
                                                                _ = len(v32)
                                                            } else {
                                                                var v93 int32 = v28
                                                                var v94 uint64 = 764
                                                                var v95 uint8 = (v61 >> ((18446744073709551615) & 63))
                                                                _ = v93
                                                                _ = v94
                                                                _ = v95
                                                            }
                                                        default:
                                                            v96 := GenS6{F0: v91, F1: 21, F2: 115}
                                                            _ = v96
                                                        }
                                                        var v97 uint64 = 1
                                                        fmt.Println(v3)
                                                        if v91 {
                                                            v2 += (v2 - (v28 - v2))
                                                            fmt.Println(v69, v70, v63)
                                                            v98 := GenS9{F0: "\n\t", F1: v28}
                                                            _ = v98
                                                        } else {
                                                            if ((true && (666 <= 4294967295)) || ((v3 * 0) <= (v3 - (v3 ^ v3)))) {
                                                                if len(v32) >= 2 {
                                                                    v32 = v32[1:len(v32)]
                                                                }
                                                                v61 -= (^(v31))
                                                                v91 = (!(((3.14159 - (0.0 - ((0.1 + -2.25) - 3.14159))) + (float64(((1.5 / v63) - v63)) / (-2.25 * 1.5))) <= 0.1))
                                                            } else {
                                                                fmt.Println(v64, fi67, v2)
                                                                fmt.Println(v69, v63)
                                                            }
                                                        }
                                                        _ = v97
                                                    }
                                                case "hello":
                                                    if len(v32) > 0 {
                                                        v32[0] = (^(v69))
                                                    }
                                                default:
                                                    v99 := GenS9{F0: "\n\t", F1: int32(v64)}
                                                    _ = v99
                                                }
                                                _ = v91
                                            }
                                            var v100 uint8 = v61
                                            _ = v100
                                        case -2147483648:
                                            v101 := GenS4{F0: v70, F1: 1, F2: -128}
                                            _ = v101
                                        case 1:
                                            v3 += 65
                                        }
                                        fmt.Println(fi68)
                                    default:
                                        v102 := GenS5{F0: 108, F1: v31}
                                        _ = v102
                                    }
                                default:
                                    fmt.Println(v70)
                                }
                                _ = v70
                                _ = v71
                            } else {
                                var v103 string = "foo bar"
                                var v104 bool = v1
                                _ = v103
                                _ = v104
                            }
                            _ = fi68
                            _ = v69
                        }
                        var v105 int16 = (v33 | int16(fi67))
                        for fi106 := 0; fi106 < 16; fi106++ {
                            v107 := GenS7{F0: v31, F1: v28, F2: "\n\t"}
                            if v1 {
                                var v108 string = "hello"
                                var v109 uint32 = 589
                                if (false || v1) {
                                    if ((42 <= v64) && ((v2 + v28) > v28)) {
                                        v110 := GenS4{F0: 1, F1: (v105 - v33), F2: -128}
                                        _ = v110
                                    }
                                    if (46 > 665) {
                                        var v111 int64 = -9223372036854775808
                                        _ = len(v32)
                                        _ = v111
                                    }
                                    if (0.0 >= v63) {
                                        var v112 float32 = float32(0.1)
                                        var v113 uint32 = 1
                                        var v114 float32 = v112
                                        if ((v64 + (v64 ^ 520)) != v64) {
                                            v115 := GenS8{F0: v108, F1: v105, F2: 1}
                                            v112 -= float32((1e10 + (1.5 * float64(((1.5 * (v63 - v112)) - ((v114 / (v63 + v63)) + ((float32(3.14159) - float32(((float64(0.1) - 0.1) - 1e10))) - ((v112 / 1e10) - (3.14159 / 1.5)))))))))
                                            _ = v115
                                        } else {
                                            var v116 uint64 = 4294967295
                                            v117 := GenS6{F0: v1, F1: 119, F2: (v31 << ((392) & 63))}
                                            var v118 uint32 = v109
                                            _ = v116
                                            _ = v117
                                            _ = v118
                                        }
                                        v1 = v1
                                        _ = v112
                                        _ = v113
                                        _ = v114
                                    } else {
                                        _ = len(v32)
                                        _ = len(v32)
                                        var v119 uint = 2
                                        if v1 {
                                            var v120 uint = v119
                                            v121 := GenS7{F0: (^(v61)), F1: v28, F2: ("a" + v108)}
                                            var v122 float32 = 1e10
                                            _ = v120
                                            _ = v121
                                            _ = v122
                                        }
                                        _ = v119
                                    }
                                }
                                _ = v108
                                _ = v109
                            } else {
                                v63 = v63
                                v61 += v31
                                var v123 int64 = (^(v3))
                                _ = v123
                            }
                            _ = fi106
                            _ = v107
                        }
                        _ = v105
                    } else {
                        switch "hello" {
                        case "foo bar":
                            switch ("foo bar" + "\n\t") {
                            case "":
                                fmt.Println(v28, v33)
                            case "hello":
                                v124 := GenS10{F0: ("hello" + "hello"), F1: "", F2: v28}
                                _ = v124
                            }
                        default:
                            for fi125 := 0; fi125 < 14; fi125++ {
                                v126 := GenS3{F0: (v64 * v64), F1: -36}
                                switch 2147483647 {
                                case -626:
                                    fmt.Println(v1, v31, v33)
                                default:
                                    v127 := GenS11{F0: 1}
                                    _ = v127
                                }
                                v32 = append(v32, (fi67 / (((fi125 << ((255) & 63))) | 1)))
                                var v128 uint8 = (^(v61))
                                _ = fi125
                                _ = v126
                                _ = v128
                            }
                        }
                        v61 = v31
                    }
                }
                _ = fi67
            }
            _ = v64
        } else {
            var v129 uint16 = 0
            v130 := GenS12{F0: 2, F1: int(v3)}
            v131 := GenS4{F0: 647, F1: (^(v33)), F2: 100}
            _ = v129
            _ = v130
            _ = v131
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v33:", v33)
    fmt.Println("v61:", v61)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v61
    _ = v63
}
