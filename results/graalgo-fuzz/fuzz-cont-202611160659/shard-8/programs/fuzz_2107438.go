package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64; F1 uint8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int64; F1 int; F2 int16 }
type GenS6 struct { F0 int16; F1 uint8; F2 int64 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int; F1 int32 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 int8; F1 bool }
type GenS11 struct { F0 uint16; F1 string }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 bool; F1 int64 }
type GenS14 struct { F0 uint32; F1 bool }
type GenS15 struct { F0 int64 }
type GenS16 struct { F0 uint64 }
type GenS17 struct { F0 uint8; F1 string }
type GenS18 struct { F0 string; F1 int8 }
type GenS19 struct { F0 uint; F1 int16 }
type GenS20 struct { F0 int16 }
type GenS21 struct { F0 uint32; F1 bool }
type GenS22 struct { F0 bool; F1 bool; F2 bool }
type GenS23 struct { F0 int64 }
type GenS24 struct { F0 uint; F1 bool; F2 string }
type GenS25 struct { F0 int16; F1 uint32; F2 bool }
type GenS26 struct { F0 int32 }
type GenS27 struct { F0 int32; F1 uint8; F2 bool }
type GenS28 struct { F0 uint8; F1 int32; F2 uint16 }
type GenS29 struct { F0 int64 }
type GenS30 struct { F0 int64; F1 uint8; F2 string }
type GenS31 struct { F0 int64; F1 uint16 }
type GenS32 struct { F0 int; F1 uint8 }
type GenS33 struct { F0 int8 }

func main() {
    v1 := []uint32{}
    var v2 int64 = 0
    v3 := []string{}
    var v4 float64 = 0.1
    v4 -= 0.0
    if true {
        if len(v3) > 0 {
            v3[0] = ("hello" + "foo bar")
        }
        v5 := GenS1{F0: 121, F1: 32767}
        _ = v5
    } else {
        if ((-428 != 0) && (("\n\t" + "a") == "a")) {
            v2 = (-(v2))
            switch "" {
            case "":
                switch "" {
                case "\n\t":
                    if (true && (-9223372036854775808 != 128)) {
                        fmt.Println(v2)
                        fmt.Println(v4, v2)
                    }
                }
            case "\n\t":
                var v6 uint64 = 42
                _ = v6
            case "foo bar":
                v2 = v2
            }
            fmt.Println(v4, v2)
        }
        var v7 int32 = -2147483648
        _ = v7
    }
    fmt.Println(v2, v2)
    var v8 int = 32768
    for fi9 := 0; fi9 < 15; fi9++ {
        for fi10 := 0; fi10 < 9; fi10++ {
            if (v4 > 0.1) {
                if (1 >= 486) {
                    fmt.Println(fi10, fi9, fi10)
                    var v11 uint16 = 2
                    switch ((((("\n\t" + "a") + "\n\t") + (("hello" + ("foo bar" + "a")) + "")) + (("hello" + "hello") + "hello")) + "a") {
                    case "hello":
                        var v12 int8 = 89
                        _ = v12
                    case "foo bar":
                        var v13 int = (v8 >> ((51) & 63))
                        _ = v13
                    }
                    _ = v11
                }
                if (100 <= -128) {
                    if true {
                        v1 = append(v1, 0)
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v14 := GenS1{F0: 0, F1: (v8 % ((fi10) | 1))}
                        _ = v14
                    }
                    switch 1 {
                    case 493:
                        v15 := GenS1{F0: 0, F1: fi10}
                        _ = v15
                    case 56:
                        v2 -= -9223372036854775808
                    case 42:
                        switch fi10 {
                        case 382:
                            fmt.Println(v8, v2, v4)
                        }
                    }
                    switch ("hello" + "") {
                    case "foo bar":
                        fmt.Println(v4)
                    case "":
                        v2 = 0
                    case "\n\t":
                        switch "hello" {
                        case "\n\t":
                            v2 += v2
                        case "hello":
                            v4 += (v4 - -2.25)
                        case "a":
                            var v16 int32 = -1
                            var v17 uint32 = 0
                            if (-32768 <= -893) {
                                if false {
                                    var v18 int32 = v16
                                    v16 += (v18 + 42)
                                    if (111 < 102) {
                                        fmt.Println(fi10)
                                    } else {
                                        v18 -= 0
                                        fmt.Println(v4, v4)
                                        var v19 string = "foo bar"
                                        _ = v19
                                    }
                                    _ = v18
                                }
                                v20 := GenS1{F0: 551, F1: fi10}
                                var v21 uint32 = v17
                                v21 -= (v17 / ((v21) | 1))
                                var v22 uint8 = 220
                                _ = v20
                                _ = v21
                                _ = v22
                            }
                            v4 -= v4
                            _ = v16
                            _ = v17
                        }
                    default:
                        v8 = 4294967295
                    }
                } else {
                    v23 := GenS1{F0: 162, F1: (^(fi9))}
                    _ = v23
                }
            } else {
                switch (^(fi10)) {
                case -1:
                    if true {
                        if true {
                            fmt.Println(fi10, fi10, v8)
                            var v24 int32 = 42
                            var v25 int16 = 32767
                            if (1 >= -128) {
                                var v26 uint = 1
                                v27 := GenS2{F0: 0}
                                _ = v26
                                _ = v27
                            } else {
                                var v28 int32 = (v24 / ((v24) | 1))
                                _ = v28
                            }
                            v29 := GenS1{F0: 65535, F1: 0}
                            _ = v24
                            _ = v25
                            _ = v29
                        } else {
                            v30 := GenS3{F0: 707, F1: 54}
                            _ = v30
                        }
                        switch ((("hello" + ("a" + "a")) + ("a" + "")) + "a") {
                        case "\n\t":
                            v8 = 32767
                            fmt.Println(v8, v4)
                        case "foo bar":
                            var v31 int8 = 78
                            _ = v31
                        }
                    } else {
                        v32 := GenS4{F0: (fi9 / (((v8 - (fi10 << ((18446744073709551615) & 63)))) | 1))}
                        var v33 bool = true
                        v34 := GenS3{F0: (v2 & -9223372036854775808), F1: 42}
                        _ = v32
                        _ = v33
                        _ = v34
                    }
                case -106:
                    if false {
                        switch (("" + ("foo bar" + "a")) + "") {
                        case "foo bar":
                            var v35 int16 = 0
                            _ = v35
                        case "a":
                            v8 = (fi10 % (((fi9 << ((0) & 63))) | 1))
                            v36 := GenS1{F0: 65535, F1: (fi9 & 9223372036854775807)}
                            v2 += v2
                            _ = v36
                        case "":
                            switch "foo bar" {
                            case "hello":
                                if (!false) {
                                    v8 -= (-(v8))
                                } else {
                                    fmt.Println(v2)
                                }
                            }
                        default:
                            if (v2 > -1) {
                                var v37 uint16 = 938
                                var v38 int32 = -2147483648
                                _ = v37
                                _ = v38
                            } else {
                                switch "a" {
                                case "hello":
                                    _ = len(v3)
                                case "\n\t":
                                    var v39 uint = 1
                                    if true {
                                        switch (fi10 + (fi10 * 9223372036854775807)) {
                                        case 978:
                                            var v40 uint = (^(v39))
                                            _ = v40
                                        default:
                                            var v41 float32 = 3.14159
                                            _ = v41
                                        }
                                    } else {
                                        v42 := GenS2{F0: 32767}
                                        var v43 int64 = v2
                                        if (128 == 42) {
                                            v2 = v43
                                            var v44 int64 = v43
                                            _ = len(v3)
                                            switch (-(fi9)) {
                                            case -835:
                                                var v45 bool = true
                                                _ = v45
                                            case 65:
                                                v46 := GenS2{F0: 65535}
                                                _ = v46
                                            default:
                                                var v47 bool = (!true)
                                                _ = v47
                                            }
                                            _ = v44
                                        } else {
                                            var v48 uint8 = 150
                                            v49 := GenS4{F0: (fi10 / ((1) | 1))}
                                            _ = v48
                                            _ = v49
                                        }
                                        v4 -= float64(-2.25)
                                        v3 = append(v3, "\n\t")
                                        _ = v42
                                        _ = v43
                                    }
                                    _ = v39
                                case "a":
                                    v50 := GenS2{F0: 945}
                                    if true {
                                        switch 435 {
                                        case 0:
                                            v8 = -2147483648
                                        case 626:
                                            fmt.Println(fi10, fi9, v8)
                                        case 42:
                                            v51 := GenS3{F0: (^(v2)), F1: 169}
                                            _ = v51
                                        default:
                                            var v52 int16 = 1
                                            _ = v52
                                        }
                                        if (!(((v4 / v4) - v4) == 3.14159)) {
                                            v2 = 9223372036854775807
                                            v53 := GenS5{F0: -1, F1: v8, F2: 1}
                                            _ = v53
                                        }
                                    }
                                    _ = v50
                                default:
                                    fmt.Println(v2, fi10, fi9)
                                }
                                switch "" {
                                case "a":
                                    fmt.Println(fi9, fi10)
                                case "hello":
                                    v3 = append(v3, "foo bar")
                                    var v54 bool = (!(!(!true)))
                                    _ = v54
                                case "\n\t":
                                    v4 -= v4
                                default:
                                    v3 = append(v3, (("\n\t" + "hello") + ""))
                                }
                            }
                        }
                        if (0 == 4294967295) {
                            v55 := GenS6{F0: 42, F1: 42, F2: v2}
                            var v56 uint8 = 169
                            fmt.Println(v56)
                            _ = v55
                            _ = v56
                        } else {
                            _ = len(v1)
                            var v57 uint = 65535
                            _ = v57
                        }
                    } else {
                        if true {
                            var v58 uint8 = 179
                            if true {
                                if ((("hello" + "") + (("foo bar" + "foo bar") + ("foo bar" + "hello"))) >= "") {
                                    v8 = fi9
                                }
                            }
                            if false {
                                v4 += (0.1 * ((1.5 - 1e10) - v4))
                                var v59 bool = true
                                v8 += (v8 + 10)
                                _ = v59
                            }
                            v58 = (v58 & (v58 ^ (v58 << ((752) & 63))))
                            if (100 > -128) {
                                if (false || (503 > 4294967295)) {
                                    v60 := GenS1{F0: 42, F1: -9223372036854775808}
                                    _ = len(v1)
                                    switch ("foo bar" + "\n\t") {
                                    case "\n\t":
                                        if false {
                                            v8 = fi10
                                            if len(v3) > 0 {
                                                v3[0] = (("hello" + "hello") + "a")
                                            }
                                        } else {
                                            v3 = append(v3, "\n\t")
                                            if false {
                                                fmt.Println(v4)
                                                v61 := GenS5{F0: -488, F1: 0, F2: -1}
                                                _ = v61
                                            } else {
                                                if (!(v2 > (-(v2)))) {
                                                    if (122 == 42) {
                                                        v62 := GenS4{F0: fi10}
                                                        var v63 uint = 139
                                                        v64 := GenS7{F0: -18}
                                                        _ = v62
                                                        _ = v63
                                                        _ = v64
                                                    }
                                                    if len(v3) > 0 {
                                                        v3[0] = ("a" + ("\n\t" + ("foo bar" + "a")))
                                                    }
                                                }
                                                switch "\n\t" {
                                                case "a":
                                                    if true {
                                                        v4 -= 0.0
                                                        v4 = v4
                                                    } else {
                                                        var v65 bool = (100 != 1)
                                                        _ = v65
                                                    }
                                                case "foo bar":
                                                    v2 = v2
                                                }
                                                if (611 < 619) {
                                                    if (false && false) {
                                                        var v66 int8 = -128
                                                        v66 -= int8(fi10)
                                                        var v67 float32 = 1.5
                                                        _ = len(v3)
                                                        var v68 int64 = (v2 / ((v2) | 1))
                                                        _ = v66
                                                        _ = v67
                                                        _ = v68
                                                    }
                                                }
                                            }
                                            var v69 int8 = 1
                                            _ = v69
                                        }
                                    case "a":
                                        if len(v1) > 0 {
                                            v1[0] = 2147483648
                                        }
                                        if true {
                                            if len(v1) > 0 {
                                                v1[0] = 2147483647
                                            }
                                        }
                                        v70 := GenS8{F0: fi9, F1: 1}
                                        _ = v70
                                    default:
                                        v71 := GenS8{F0: v8, F1: int32(v2)}
                                        _ = v71
                                    }
                                    switch "a" {
                                    case "foo bar":
                                        if len(v3) > 0 {
                                            v3[0] = "foo bar"
                                        }
                                    case "hello":
                                        if true {
                                            v58 = 1
                                            fmt.Println(v2)
                                            if true {
                                                v72 := GenS3{F0: 0, F1: v58}
                                                _ = v72
                                            } else {
                                                if (32767 == 32767) {
                                                    var v73 bool = (!false)
                                                    fmt.Println(v8)
                                                    _ = v73
                                                } else {
                                                    v58 = (v58 << ((1) & 63))
                                                }
                                            }
                                        }
                                    case "\n\t":
                                        if false {
                                            switch 0 {
                                            case 9223372036854775807:
                                                var v74 int = (v8 * (v8 - (fi10 * (fi10 ^ v8))))
                                                _ = v74
                                            }
                                            switch 0 {
                                            case -9223372036854775808:
                                                v75 := GenS1{F0: uint(v58), F1: fi9}
                                                var v76 uint64 = 102
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                _ = v75
                                                _ = v76
                                            case -1:
                                                if ((2147483647 < 65535) && true) {
                                                    v77 := GenS9{F0: 127}
                                                    v2 -= -9223372036854775808
                                                    var v78 int16 = 1000
                                                    _ = v77
                                                    _ = v78
                                                } else {
                                                    if (true || false) {
                                                        var v79 int64 = v2
                                                        _ = v79
                                                    } else {
                                                        v80 := GenS10{F0: 100, F1: true}
                                                        _ = v80
                                                    }
                                                    if false {
                                                        v4 = 1e10
                                                        switch v8 {
                                                        case -2147483648:
                                                            v3 = append(v3, "")
                                                        case 2147483647:
                                                            v81 := GenS5{F0: (-(v2)), F1: (v8 % (((-(v8))) | 1)), F2: 32767}
                                                            _ = v81
                                                        case 127:
                                                            var v82 int8 = -1
                                                            _ = v82
                                                        default:
                                                            v83 := GenS4{F0: v8}
                                                            _ = v83
                                                        }
                                                    } else {
                                                        v3 = append(v3, ("\n\t" + (((("foo bar" + (("a" + "hello") + "a")) + "a") + "a") + ("a" + "foo bar"))))
                                                        var v84 uint64 = 2
                                                        if true {
                                                            var v85 float32 = 3.14159
                                                            if len(v3) > 0 {
                                                                v3[0] = ""
                                                            }
                                                            v86 := GenS5{F0: (^(v2)), F1: 42, F2: -1}
                                                            _ = v85
                                                            _ = v86
                                                        } else {
                                                            switch ("hello" + "\n\t") {
                                                            case "\n\t":
                                                                v87 := GenS6{F0: 1000, F1: v58, F2: -202}
                                                                _ = v87
                                                            case "a":
                                                                var v88 uint32 = uint32(v84)
                                                                _ = v88
                                                            }
                                                            _ = len(v1)
                                                        }
                                                        _ = v84
                                                    }
                                                }
                                            case 4294967295:
                                                v2 += v2
                                            }
                                            if (true && true) {
                                                var v89 int32 = 1
                                                var v90 string = "hello"
                                                if ((!false) && false) {
                                                    if false {
                                                        var v91 uint32 = 2147483647
                                                        v89 -= (v89 / (((v89 >> ((65535) & 63))) | 1))
                                                        _ = v91
                                                    } else {
                                                        var v92 float32 = 0.0
                                                        var v93 string = "a"
                                                        _ = v92
                                                        _ = v93
                                                    }
                                                    v4 -= v4
                                                    v90 = v90
                                                }
                                                _ = v89
                                                _ = v90
                                            }
                                            switch (^(v8)) {
                                            case 642:
                                                switch "hello" {
                                                case "":
                                                    switch "hello" {
                                                    case "hello":
                                                        if (751 <= 924) {
                                                            v2 += v2
                                                        }
                                                    case "foo bar":
                                                        fmt.Println(v58)
                                                    case "\n\t":
                                                        v8 += v8
                                                    }
                                                case "hello":
                                                    fmt.Println(v58, v58, v58)
                                                    v4 = (v4 * float64(((-2.25 * (-2.25 * 0.0)) - 0.0)))
                                                    if (!true) {
                                                        v94 := GenS3{F0: v2, F1: (v58 + (v58 & v58))}
                                                        var v95 bool = true
                                                        v96 := GenS1{F0: 42, F1: fi10}
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            v97 := GenS11{F0: 0, F1: ("a" + ("" + ("a" + "foo bar")))}
                                                            _ = v97
                                                        default:
                                                            var v98 uint8 = 213
                                                            _ = v98
                                                        }
                                                        switch v8 {
                                                        case 0:
                                                            switch "\n\t" {
                                                            case "a":
                                                                if len(v3) > 0 {
                                                                    v3[0] = "foo bar"
                                                                }
                                                            default:
                                                                if (v95 && (!v95)) {
                                                                    v99 := GenS8{F0: -406, F1: 2147483647}
                                                                    switch 790 {
                                                                    case -746:
                                                                        if (2 >= 246) {
                                                                            v100 := GenS3{F0: (v2 * 1), F1: 255}
                                                                            switch "hello" {
                                                                            case "foo bar":
                                                                                if ((v95 || true) || ((v4 * (0.0 / v4)) < ((v4 - ((v4 * float64(0.0)) * 1.5)) - 1e10))) {
                                                                                    v58 -= (^(v58))
                                                                                }
                                                                            case "\n\t":
                                                                                switch 0 {
                                                                                case 4294967295:
                                                                                    var v101 int32 = -2147483648
                                                                                    switch fi9 {
                                                                                    case 945:
                                                                                        v4 = ((v4 + v4) * (1.5 + (v4 / v4)))
                                                                                    case -2147483648:
                                                                                        v102 := GenS12{F0: 42}
                                                                                        _ = v102
                                                                                    default:
                                                                                        v103 := GenS12{F0: -128}
                                                                                        _ = v103
                                                                                    }
                                                                                    switch ("a" + "hello") {
                                                                                    case "hello":
                                                                                        if ((v2 % (((v2 % ((9223372036854775807) | 1))) | 1)) >= 0) {
                                                                                            var v104 int8 = -72
                                                                                            _ = v104
                                                                                        } else {
                                                                                            var v105 float32 = (3.14159 / 1.0)
                                                                                            _ = v105
                                                                                        }
                                                                                    case "":
                                                                                        if v95 {
                                                                                            v4 = (((v4 + v4) + float64((0.1 * 0.0))) - (0.1 / v4))
                                                                                            v4 -= 1e10
                                                                                        }
                                                                                    default:
                                                                                        switch "\n\t" {
                                                                                        case "\n\t":
                                                                                            var v106 string = "\n\t"
                                                                                            _ = v106
                                                                                        case "hello":
                                                                                            var v107 int64 = (v2 - -765)
                                                                                            if v95 {
                                                                                                v101 += v101
                                                                                            }
                                                                                            v108 := GenS6{F0: 42, F1: v58, F2: (v107 | -9223372036854775808)}
                                                                                            _ = v107
                                                                                            _ = v108
                                                                                        default:
                                                                                            var v109 float64 = (float64(-2.25) * (v4 + (1e10 - (v4 * 0.1))))
                                                                                            if len(v3) >= 2 {
                                                                                                v3 = v3[1:len(v3)]
                                                                                            }
                                                                                            _ = v109
                                                                                        }
                                                                                    }
                                                                                    _ = v101
                                                                                case 2147483647:
                                                                                    if v95 {
                                                                                        var v110 uint8 = (v58 * (v58 << ((2) & 63)))
                                                                                        _ = len(v1)
                                                                                        v111 := GenS13{F0: true, F1: v2}
                                                                                        _ = v110
                                                                                        _ = v111
                                                                                    }
                                                                                case -589:
                                                                                    if (1e10 >= 3.14159) {
                                                                                        if v95 {
                                                                                            if (!true) {
                                                                                                switch "\n\t" {
                                                                                                case "a":
                                                                                                    switch ("foo bar" + ("foo bar" + "\n\t")) {
                                                                                                    case "hello":
                                                                                                        v112 := GenS14{F0: 42, F1: (302 >= -1)}
                                                                                                        _ = v112
                                                                                                    case "\n\t":
                                                                                                        var v113 int32 = 42
                                                                                                        _ = v113
                                                                                                    case "a":
                                                                                                        v114 := GenS15{F0: (v2 - 9223372036854775807)}
                                                                                                        _ = v114
                                                                                                    }
                                                                                                case "":
                                                                                                    var v115 uint8 = 93
                                                                                                    _ = v115
                                                                                                default:
                                                                                                    var v116 bool = (18446744073709551615 != 0)
                                                                                                    _ = v116
                                                                                                }
                                                                                            }
                                                                                            switch ("foo bar" + ("hello" + ("a" + ((("a" + ("foo bar" + "foo bar")) + "foo bar") + "foo bar")))) {
                                                                                            case "":
                                                                                                v95 = v95
                                                                                            default:
                                                                                                if len(v1) >= 2 {
                                                                                                    v1 = v1[1:len(v1)]
                                                                                                }
                                                                                            }
                                                                                        } else {
                                                                                            if v95 {
                                                                                                v95 = v95
                                                                                            } else {
                                                                                                var v117 int = (fi9 >> ((468) & 63))
                                                                                                _ = v117
                                                                                            }
                                                                                            switch fi10 {
                                                                                            case -9223372036854775808:
                                                                                                _ = len(v3)
                                                                                            }
                                                                                            switch 9223372036854775807 {
                                                                                            case -9223372036854775808:
                                                                                                var v118 int32 = 803
                                                                                                _ = v118
                                                                                            default:
                                                                                                switch (v8 * (v8 >> ((4294967295) & 63))) {
                                                                                                case 42:
                                                                                                    v58 += 128
                                                                                                    var v119 bool = (-724 >= 914)
                                                                                                    _ = v119
                                                                                                case 128:
                                                                                                    if v95 {
                                                                                                        var v120 int8 = 127
                                                                                                        v120 += (v120 << ((42) & 63))
                                                                                                        var v121 int32 = 0
                                                                                                        _ = v120
                                                                                                        _ = v121
                                                                                                    }
                                                                                                default:
                                                                                                    v122 := GenS16{F0: 0}
                                                                                                    switch "hello" {
                                                                                                    case "\n\t":
                                                                                                        if len(v1) > 0 {
                                                                                                            v1[0] = 565
                                                                                                        }
                                                                                                    case "hello":
                                                                                                        switch "a" {
                                                                                                        case "a":
                                                                                                            var v123 float64 = v4
                                                                                                            _ = v123
                                                                                                        case "":
                                                                                                            if (498 <= -1) {
                                                                                                                v95 = v95
                                                                                                                var v124 uint64 = 1
                                                                                                                _ = v124
                                                                                                            } else {
                                                                                                                fmt.Println(v4, fi9)
                                                                                                                v125 := GenS17{F0: v58, F1: "a"}
                                                                                                                switch "\n\t" {
                                                                                                                case "\n\t":
                                                                                                                    if (v95 || v95) {
                                                                                                                        fmt.Println(v8)
                                                                                                                        var v126 uint32 = 383
                                                                                                                        var v127 uint32 = (v126 & v126)
                                                                                                                        v4 -= 0.1
                                                                                                                        _ = v126
                                                                                                                        _ = v127
                                                                                                                    }
                                                                                                                case "hello":
                                                                                                                    v3 = append(v3, "")
                                                                                                                }
                                                                                                                v4 -= 0.1
                                                                                                                _ = len(v1)
                                                                                                                _ = v125
                                                                                                            }
                                                                                                        case "foo bar":
                                                                                                            v95 = ((v58 & 2) <= 106)
                                                                                                            _ = len(v1)
                                                                                                            var v128 int16 = 42
                                                                                                            _ = v128
                                                                                                        default:
                                                                                                            v4 -= 3.14159
                                                                                                        }
                                                                                                    case "foo bar":
                                                                                                        if v95 {
                                                                                                            fmt.Println(v58)
                                                                                                        } else {
                                                                                                            var v129 uint32 = 1
                                                                                                            var v130 float64 = 1.5
                                                                                                            switch (v8 & (-(v8))) {
                                                                                                            case 42:
                                                                                                                var v131 int32 = 42
                                                                                                                _ = v131
                                                                                                            case 0:
                                                                                                                v1 = append(v1, v129)
                                                                                                            case 9223372036854775807:
                                                                                                                switch "" {
                                                                                                                case "foo bar":
                                                                                                                    if v95 {
                                                                                                                        v130 = v130
                                                                                                                        v3 = append(v3, "\n\t")
                                                                                                                    } else {
                                                                                                                        var v132 float64 = float64(1e10)
                                                                                                                        fmt.Println(v8)
                                                                                                                        var v133 int64 = (v2 % (((v2 & (v2 ^ (^(v2))))) | 1))
                                                                                                                        var v134 int8 = 1
                                                                                                                        _ = v132
                                                                                                                        _ = v133
                                                                                                                        _ = v134
                                                                                                                    }
                                                                                                                case "\n\t":
                                                                                                                    if len(v1) > 0 {
                                                                                                                        v1[0] = v129
                                                                                                                    }
                                                                                                                default:
                                                                                                                    var v135 uint = 957
                                                                                                                    var v136 uint8 = uint8(v135)
                                                                                                                    if len(v1) >= 2 {
                                                                                                                        v1 = v1[1:len(v1)]
                                                                                                                    }
                                                                                                                    _ = v135
                                                                                                                    _ = v136
                                                                                                                }
                                                                                                            default:
                                                                                                                v3 = append(v3, "\n\t")
                                                                                                            }
                                                                                                            _ = v129
                                                                                                            _ = v130
                                                                                                        }
                                                                                                    default:
                                                                                                        switch "\n\t" {
                                                                                                        case "":
                                                                                                            if v95 {
                                                                                                                if (v95 || true) {
                                                                                                                    v137 := GenS12{F0: 1}
                                                                                                                    var v138 int = fi10
                                                                                                                    _ = v137
                                                                                                                    _ = v138
                                                                                                                }
                                                                                                            } else {
                                                                                                                if v95 {
                                                                                                                    var v139 uint = 688
                                                                                                                    fmt.Println(v58)
                                                                                                                    _ = v139
                                                                                                                }
                                                                                                            }
                                                                                                        default:
                                                                                                            switch (-(fi9)) {
                                                                                                            case 2147483647:
                                                                                                                var v140 int32 = 638
                                                                                                                _ = v140
                                                                                                            case -1:
                                                                                                                v4 = v4
                                                                                                            case 32768:
                                                                                                                var v141 uint64 = 553
                                                                                                                _ = v141
                                                                                                            default:
                                                                                                                var v142 int16 = -32768
                                                                                                                _ = v142
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v122
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    } else {
                                                                                        switch "foo bar" {
                                                                                        case "\n\t":
                                                                                            v143 := GenS18{F0: (("" + ("a" + ("" + "hello"))) + "foo bar"), F1: 88}
                                                                                            _ = v143
                                                                                        default:
                                                                                            var v144 uint = 196
                                                                                            _ = v144
                                                                                        }
                                                                                        if len(v1) > 0 {
                                                                                            v1[0] = 4294967295
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    var v145 uint64 = 4294967295
                                                                                    _ = v145
                                                                                }
                                                                            }
                                                                            _ = v100
                                                                        }
                                                                    case -1:
                                                                        v146 := GenS19{F0: 328, F1: 91}
                                                                        _ = v146
                                                                    case 9223372036854775807:
                                                                        switch fi10 {
                                                                        case 973:
                                                                            switch ("\n\t" + "\n\t") {
                                                                            case "foo bar":
                                                                                v95 = (-565 >= -307)
                                                                            case "":
                                                                                fmt.Println(v8, v2, v2)
                                                                            case "hello":
                                                                                v95 = v95
                                                                            default:
                                                                                if (v4 != v4) {
                                                                                    switch 127 {
                                                                                    case 32767:
                                                                                        v147 := GenS20{F0: -1}
                                                                                        _ = v147
                                                                                    }
                                                                                }
                                                                            }
                                                                        case -817:
                                                                            if v95 {
                                                                                switch "a" {
                                                                                case "":
                                                                                    var v148 int16 = -992
                                                                                    _ = v148
                                                                                default:
                                                                                    v4 = v4
                                                                                }
                                                                                switch fi9 {
                                                                                case 366:
                                                                                    var v149 int64 = v2
                                                                                    _ = v149
                                                                                case 140:
                                                                                    var v150 uint16 = 32767
                                                                                    _ = v150
                                                                                }
                                                                                if (((1.5 + 1.5) + (3.14159 + 3.14159)) <= 1.5) {
                                                                                    var v151 uint8 = v58
                                                                                    switch "hello" {
                                                                                    case "foo bar":
                                                                                        if false {
                                                                                            v1 = append(v1, 884)
                                                                                        } else {
                                                                                            v151 = 42
                                                                                            var v152 int32 = 2147483647
                                                                                            _ = v152
                                                                                        }
                                                                                    case "\n\t":
                                                                                        v4 = v4
                                                                                    default:
                                                                                        if (-67 >= 42) {
                                                                                            v8 = (^(v8))
                                                                                            var v153 int32 = 1
                                                                                            _ = v153
                                                                                        }
                                                                                    }
                                                                                    _ = v151
                                                                                } else {
                                                                                    var v154 string = ""
                                                                                    if v95 {
                                                                                        v154 = "hello"
                                                                                    } else {
                                                                                        if v95 {
                                                                                            var v155 int32 = -641
                                                                                            v156 := GenS21{F0: 1, F1: v95}
                                                                                            _ = v155
                                                                                            _ = v156
                                                                                        }
                                                                                        switch (v154 + v154) {
                                                                                        case "":
                                                                                            if len(v3) > 0 {
                                                                                                v3[0] = v154
                                                                                            }
                                                                                        case "hello":
                                                                                            v95 = (!true)
                                                                                        }
                                                                                        v157 := GenS18{F0: "hello", F1: 97}
                                                                                        var v158 uint8 = v58
                                                                                        _ = v157
                                                                                        _ = v158
                                                                                    }
                                                                                    v154 = (((v154 + v154) + v154) + ("a" + "a"))
                                                                                    _ = v154
                                                                                }
                                                                            } else {
                                                                                v4 = -2.25
                                                                            }
                                                                        case 78:
                                                                            v4 += v4
                                                                        default:
                                                                            var v159 int16 = -32768
                                                                            _ = v159
                                                                        }
                                                                        var v160 float64 = -2.25
                                                                        if false {
                                                                            v1 = append(v1, 449)
                                                                        }
                                                                        _ = v160
                                                                    default:
                                                                        if false {
                                                                            if v95 {
                                                                                _ = len(v3)
                                                                                if (568 <= 8) {
                                                                                    var v161 bool = false
                                                                                    v4 += 0.1
                                                                                    _ = v161
                                                                                }
                                                                            } else {
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = 2147483647
                                                                                }
                                                                            }
                                                                            if len(v1) > 0 {
                                                                                v1[0] = 2147483648
                                                                            }
                                                                            _ = len(v1)
                                                                            var v162 uint32 = 0
                                                                            if v95 {
                                                                                switch (fi10 + fi10) {
                                                                                case 32767:
                                                                                    switch ("a" + ("\n\t" + "hello")) {
                                                                                    case "foo bar":
                                                                                        var v163 float64 = 1.5
                                                                                        _ = v163
                                                                                    }
                                                                                case 42:
                                                                                    if ((fi9 % ((-2147483648) | 1)) > (v8 + (fi10 & (fi9 >> ((453) & 63))))) {
                                                                                        _ = len(v1)
                                                                                        v164 := GenS22{F0: (!true), F1: (515 < 255), F2: false}
                                                                                        fmt.Println(v2, fi10)
                                                                                        _ = v164
                                                                                    } else {
                                                                                        v165 := GenS23{F0: -1}
                                                                                        _ = v165
                                                                                    }
                                                                                default:
                                                                                    if v95 {
                                                                                        v166 := GenS24{F0: 255, F1: true, F2: ("foo bar" + "a")}
                                                                                        var v167 int = 366
                                                                                        _ = len(v1)
                                                                                        v8 = -1
                                                                                        if v95 {
                                                                                            v58 -= (v58 | 255)
                                                                                        } else {
                                                                                            switch v167 {
                                                                                            case 32767:
                                                                                                v8 = 42
                                                                                            case 0:
                                                                                                v167 = v8
                                                                                            default:
                                                                                                v1 = append(v1, 1)
                                                                                            }
                                                                                            fmt.Println(v8, v8, v167)
                                                                                        }
                                                                                        _ = v166
                                                                                        _ = v167
                                                                                    }
                                                                                }
                                                                                switch v8 {
                                                                                case 42:
                                                                                    var v168 uint = 0
                                                                                    _ = v168
                                                                                case 0:
                                                                                    v58 += (v58 ^ (v58 * 128))
                                                                                case -207:
                                                                                    v4 += 3.14159
                                                                                    v162 = (v162 + 2147483648)
                                                                                    fmt.Println(v162, fi10)
                                                                                }
                                                                                v95 = (68 > 132)
                                                                            }
                                                                            _ = v162
                                                                        }
                                                                    }
                                                                    var v169 int = (fi9 << ((2) & 63))
                                                                    _ = v169
                                                                    _ = v99
                                                                }
                                                            }
                                                        case 32767:
                                                            if (v8 != (fi10 % (((-(fi10))) | 1))) {
                                                                fmt.Println(v58, fi10)
                                                            }
                                                            var v170 int = fi10
                                                            v4 = v4
                                                            _ = v170
                                                        }
                                                        _ = v94
                                                        _ = v95
                                                        _ = v96
                                                    } else {
                                                        var v171 uint8 = v58
                                                        v8 = 32767
                                                        v58 -= v171
                                                        _ = v171
                                                    }
                                                }
                                            case -9223372036854775808:
                                                var v172 int64 = v2
                                                _ = v172
                                            }
                                        }
                                    }
                                    var v173 bool = true
                                    fmt.Println(fi10, fi9)
                                    _ = v173
                                    _ = v60
                                } else {
                                    if (!false) {
                                        v174 := GenS25{F0: -595, F1: 1, F2: (((1e10 * (1.5 + v4)) * 3.14159) > ((0.0 * (float64(0.1) * (3.14159 / v4))) / 1.0))}
                                        switch ((("\n\t" + "") + ("" + "a")) + "hello") {
                                        case "foo bar":
                                            v175 := GenS13{F0: false, F1: v2}
                                            _ = v175
                                        case "\n\t":
                                            switch "a" {
                                            case "foo bar":
                                                var v176 uint32 = 2147483648
                                                _ = v176
                                            }
                                        case "a":
                                            if true {
                                                fmt.Println(fi9, fi9)
                                                v177 := GenS26{F0: -2147483648}
                                                var v178 int16 = -737
                                                _ = v177
                                                _ = v178
                                            } else {
                                                var v179 int32 = -2147483648
                                                v180 := GenS27{F0: int32(v58), F1: v58, F2: false}
                                                _ = v179
                                                _ = v180
                                            }
                                        }
                                        _ = v174
                                    } else {
                                        v8 -= v8
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        var v181 uint64 = 42
                                        if (2 < (^(v58))) {
                                            if (1e10 >= 3.14159) {
                                                v4 += (float64(0.0) - ((v4 / v4) * (v4 / 1e10)))
                                                v3 = append(v3, "\n\t")
                                                v8 += 32768
                                            } else {
                                                if ((v2 > v2) || false) {
                                                    switch "hello" {
                                                    case "hello":
                                                        v181 -= (v181 >> ((73) & 63))
                                                    case "\n\t":
                                                        switch fi10 {
                                                        case 32767:
                                                            var v182 uint = 0
                                                            _ = v182
                                                        case -2147483648:
                                                            var v183 int = -718
                                                            _ = v183
                                                        }
                                                    case "foo bar":
                                                        var v184 uint8 = (^(v58))
                                                        _ = v184
                                                    default:
                                                        if false {
                                                            v4 -= 3.14159
                                                        } else {
                                                            v185 := GenS28{F0: 2, F1: -1, F2: 2}
                                                            v4 = (0.1 - v4)
                                                            _ = v185
                                                        }
                                                    }
                                                    switch (fi10 - -195) {
                                                    case 127:
                                                        var v186 int32 = -2147483648
                                                        _ = v186
                                                    case -124:
                                                        if len(v1) > 0 {
                                                            v1[0] = uint32(v181)
                                                        }
                                                    case -753:
                                                        v58 += (v58 - (v58 % (((^(v58))) | 1)))
                                                    default:
                                                        var v187 int = fi9
                                                        _ = v187
                                                    }
                                                    v188 := GenS19{F0: 210, F1: 1000}
                                                    _ = v188
                                                } else {
                                                    if true {
                                                        v4 = v4
                                                        if (255 == 255) {
                                                            if (true && true) {
                                                                v181 += 989
                                                            } else {
                                                                fmt.Println(fi10, v58, fi10)
                                                                if (true && false) {
                                                                    fmt.Println(fi9, v58)
                                                                    v189 := GenS29{F0: v2}
                                                                    fmt.Println(v58, v58, v2)
                                                                    v181 = uint64(v58)
                                                                    if true {
                                                                        if true {
                                                                            v4 = (0.1 * (v4 * (1e10 + float64(3.14159))))
                                                                            var v190 uint = 65535
                                                                            _ = v190
                                                                        }
                                                                        v181 += (v181 >> ((1) & 63))
                                                                        v3 = append(v3, ("\n\t" + "\n\t"))
                                                                    }
                                                                    v2 -= 401
                                                                    _ = v189
                                                                }
                                                                switch fi9 {
                                                                case 799:
                                                                    fmt.Println(v181, v181, v181)
                                                                default:
                                                                    var v191 int64 = 555
                                                                    v191 = (v191 << ((797) & 63))
                                                                    if ((-2.25 != -2.25) && (!true)) {
                                                                        v192 := GenS30{F0: (v191 | (v191 - (v191 / ((-30) | 1)))), F1: (v58 >> ((0) & 63)), F2: "\n\t"}
                                                                        fmt.Println(v8, fi9)
                                                                        _ = v192
                                                                    } else {
                                                                        switch "\n\t" {
                                                                        case "hello":
                                                                            if len(v1) >= 2 {
                                                                                v1 = v1[1:len(v1)]
                                                                            }
                                                                        }
                                                                        v193 := GenS17{F0: v58, F1: "hello"}
                                                                        _ = v193
                                                                    }
                                                                    _ = v191
                                                                }
                                                            }
                                                            var v194 int8 = 127
                                                            _ = v194
                                                        }
                                                        if ((v181 << ((612) & 63)) != v181) {
                                                            var v195 float64 = v4
                                                            var v196 int16 = 1000
                                                            if false {
                                                                v58 = (v58 | 114)
                                                            } else {
                                                                fmt.Println(v181, v2)
                                                                v196 -= v196
                                                                if false {
                                                                    var v197 int8 = 58
                                                                    var v198 int16 = v196
                                                                    _ = v197
                                                                    _ = v198
                                                                }
                                                            }
                                                            _ = v195
                                                            _ = v196
                                                        }
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        if (!(42 == 2147483647)) {
                                                            if (!false) {
                                                                switch "\n\t" {
                                                                case "\n\t":
                                                                    v8 -= v8
                                                                default:
                                                                    v199 := GenS3{F0: v2, F1: (^(v58))}
                                                                    _ = v199
                                                                }
                                                                if true {
                                                                    v200 := GenS2{F0: 32767}
                                                                    if len(v3) > 0 {
                                                                        v3[0] = (("hello" + (("" + "") + "\n\t")) + "hello")
                                                                    }
                                                                    _ = v200
                                                                } else {
                                                                    var v201 uint8 = v58
                                                                    v58 = (v58 << ((42) & 63))
                                                                    v202 := GenS15{F0: (v2 % (((v2 + 42)) | 1))}
                                                                    v58 += (^(v201))
                                                                    fmt.Println(v181, v201)
                                                                    _ = v201
                                                                    _ = v202
                                                                }
                                                            }
                                                            switch (^(fi9)) {
                                                            case 18:
                                                                if (0 < 854) {
                                                                    switch "a" {
                                                                    case "hello":
                                                                        v203 := GenS18{F0: ("\n\t" + (("hello" + ("a" + ("foo bar" + "foo bar"))) + "")), F1: -66}
                                                                        _ = v203
                                                                    case "\n\t":
                                                                        v58 += v58
                                                                    default:
                                                                        v204 := GenS2{F0: 2}
                                                                        _ = v204
                                                                    }
                                                                    if true {
                                                                        v4 -= ((0.0 + v4) * 0.1)
                                                                        if ((((((-2.25 + (0.1 + v4)) + (v4 / float64(-2.25))) - (float64(((1.5 - ((-2.25 + 1e10) + 0.0)) + 0.1)) - float64(3.14159))) == v4) || (false || ("a" <= (("" + ("foo bar" + "\n\t")) + "foo bar")))) || (1 == 0)) {
                                                                            var v205 uint = 255
                                                                            v206 := GenS30{F0: (v2 & int64(v205)), F1: (v58 % (((v58 << (((v205 % ((v205) | 1))) & 63))) | 1)), F2: (("" + (("" + (("foo bar" + "foo bar") + "hello")) + "\n\t")) + ("foo bar" + ("" + (("" + "") + ("" + ("hello" + "a"))))))}
                                                                            var v207 int16 = 42
                                                                            v208 := GenS14{F0: 1, F1: (((0.1 / ((float32(1.5) / 3.14159) / (((float32(v4) + float32(v4)) + float32((float64(1e10) - 0.0))) / 1.5))) * (1.5 / (float32(3.14159) * (1e10 * 0.1)))) >= 0.1)}
                                                                            if (100 > -1) {
                                                                                v3 = append(v3, "")
                                                                            } else {
                                                                                v207 -= v207
                                                                                if (!false) {
                                                                                    v3 = append(v3, "hello")
                                                                                    switch (fi9 % ((0) | 1)) {
                                                                                    case -1:
                                                                                        var v209 string = ((("foo bar" + "a") + ("foo bar" + "hello")) + "")
                                                                                        _ = v209
                                                                                    case 4294967295:
                                                                                        var v210 int16 = 217
                                                                                        _ = v210
                                                                                    }
                                                                                    if len(v1) >= 2 {
                                                                                        v1 = v1[1:len(v1)]
                                                                                    }
                                                                                } else {
                                                                                    v211 := GenS6{F0: (^(v207)), F1: uint8(v181), F2: (v2 >> (((v205 >> ((v205) & 63))) & 63))}
                                                                                    if false {
                                                                                        if (false || (v181 >= v181)) {
                                                                                            v212 := GenS17{F0: v58, F1: ""}
                                                                                            v4 = v4
                                                                                            _ = v212
                                                                                        }
                                                                                        var v213 int16 = (v207 & (v207 | -1))
                                                                                        var v214 int32 = 669
                                                                                        _ = v213
                                                                                        _ = v214
                                                                                    } else {
                                                                                        v58 += v58
                                                                                        v4 = ((3.14159 + 3.14159) - v4)
                                                                                        v181 = (v181 % (((v181 << ((uint(v58)) & 63))) | 1))
                                                                                    }
                                                                                    _ = v211
                                                                                }
                                                                            }
                                                                            _ = v205
                                                                            _ = v206
                                                                            _ = v207
                                                                            _ = v208
                                                                        }
                                                                    }
                                                                }
                                                            case 0:
                                                                if false {
                                                                    v181 += v181
                                                                }
                                                            default:
                                                                var v215 float32 = -2.25
                                                                v58 = (v58 - (^(v58)))
                                                                v215 = v215
                                                                var v216 int16 = 0
                                                                _ = v215
                                                                _ = v216
                                                            }
                                                            switch fi10 {
                                                            case 32768:
                                                                switch ("" + ("a" + "foo bar")) {
                                                                case "\n\t":
                                                                    switch -1 {
                                                                    case -247:
                                                                        switch fi10 {
                                                                        case 9223372036854775807:
                                                                            var v217 uint16 = 161
                                                                            _ = v217
                                                                        case 32768:
                                                                            var v218 uint8 = v58
                                                                            _ = v218
                                                                        case -9223372036854775808:
                                                                            v8 -= (v8 / (((v8 - v8)) | 1))
                                                                        }
                                                                    default:
                                                                        fmt.Println(fi9, v8, v181)
                                                                    }
                                                                }
                                                            case -292:
                                                                if true {
                                                                    if ("" < ("a" + "")) {
                                                                        var v219 uint8 = (^(v58))
                                                                        v220 := GenS31{F0: -795, F1: 2}
                                                                        v4 += v4
                                                                        _ = v219
                                                                        _ = v220
                                                                    }
                                                                } else {
                                                                    if true {
                                                                        fmt.Println(v8, v2)
                                                                        switch ((("\n\t" + (("hello" + ("hello" + "\n\t")) + "")) + ("hello" + "a")) + ("" + ("" + ""))) {
                                                                        case "foo bar":
                                                                            v221 := GenS32{F0: fi10, F1: v58}
                                                                            _ = v221
                                                                        }
                                                                        v2 = (v2 & (v2 - (^(v2))))
                                                                    }
                                                                    v222 := GenS9{F0: 7}
                                                                    _ = v222
                                                                }
                                                            case 2147483647:
                                                                v223 := GenS26{F0: 363}
                                                                _ = v223
                                                            }
                                                        } else {
                                                            fmt.Println(v58, fi10)
                                                        }
                                                    }
                                                    var v224 uint8 = uint8(v181)
                                                    v4 = v4
                                                    _ = v224
                                                }
                                            }
                                        } else {
                                            v2 -= (v2 << ((4294967295) & 63))
                                            var v225 int8 = 127
                                            var v226 uint = 1
                                            _ = v225
                                            _ = v226
                                        }
                                        _ = v181
                                    }
                                    var v227 int16 = 32767
                                    v228 := GenS23{F0: v2}
                                    _ = v227
                                    _ = v228
                                }
                                var v229 int32 = 2147483647
                                var v230 float32 = (1e10 - 0.0)
                                _ = v229
                                _ = v230
                            }
                            var v231 uint = 2
                            var v232 string = (("hello" + (("foo bar" + ((("\n\t" + "hello") + "foo bar") + ("" + ((((("hello" + "a") + "hello") + (("" + "\n\t") + ("\n\t" + "a"))) + ("\n\t" + ("hello" + ""))) + "hello")))) + "a")) + "\n\t")
                            _ = v231
                            _ = v232
                            _ = v58
                        }
                    }
                case 1:
                    v2 = v2
                default:
                    var v233 string = "\n\t"
                    _ = v233
                }
                var v234 int8 = -128
                _ = v234
            }
            var v235 bool = false
            v2 += v2
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v236 := GenS18{F0: "\n\t", F1: 100}
            v235 = (fi9 <= v8)
            _ = fi10
            _ = v235
            _ = v236
        }
        for fi237 := 0; fi237 < 12; fi237++ {
            v2 = (v2 - (-(v2)))
            v238 := GenS10{F0: -40, F1: (!false)}
            _ = fi237
            _ = v238
        }
        _ = fi9
    }
    for fi239 := 0; fi239 < 4; fi239++ {
        v240 := GenS22{F0: true, F1: false, F2: false}
        _ = fi239
        _ = v240
    }
    for fi241 := 0; fi241 < 10; fi241++ {
        v4 += (float64(0.1) / (3.14159 / 1.0))
        _ = fi241
    }
    v242 := GenS33{F0: 0}
    v3 = append(v3, "")
    var v243 float64 = v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v243:", v243)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v242
    _ = v243
    _ = v3
    _ = v4
    _ = v8
}
