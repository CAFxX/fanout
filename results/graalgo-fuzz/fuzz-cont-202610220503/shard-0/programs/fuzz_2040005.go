package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 string }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint; F1 uint16 }
type GenS5 struct { F0 bool; F1 uint; F2 uint }
type GenS6 struct { F0 int64; F1 uint; F2 uint }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int16; F1 int16; F2 uint64 }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 uint8; F1 uint8; F2 int64 }
type GenS13 struct { F0 int; F1 uint16; F2 int16 }

func main() {
    var v1 int64 = -1
    v2 := []string{"\n\t"}
    v3 := make(map[uint32]string)
    v4 := make(map[string]uint32)
    if false {
        if true {
            if true {
                v5 := GenS1{F0: 1, F1: 42, F2: "\n\t"}
                _ = v5
            }
        } else {
            if true {
                for fi6 := 0; fi6 < 2; fi6++ {
                    v1 = v1
                    v1 = (v1 % ((v1) | 1))
                    v1 += (v1 | (v1 ^ 1))
                    _ = fi6
                }
                var v7 uint64 = 4294967295
                switch 94 {
                case -9223372036854775808:
                    for fi8 := 0; fi8 < 8; fi8++ {
                        v9 := GenS1{F0: 42, F1: 1, F2: "\n\t"}
                        _ = fi8
                        _ = v9
                    }
                }
                _ = v7
            }
        }
    }
    switch ("\n\t" + "a") {
    case "":
        if true {
            if len(v2) > 0 {
                v2[0] = ("a" + "\n\t")
            }
            for fi10 := 0; fi10 < 2; fi10++ {
                var v11 uint8 = 127
                v1 = (v1 & 0)
                for fi12 := 0; fi12 < 17; fi12++ {
                    var v13 uint32 = 4294967295
                    _ = fi12
                    _ = v13
                }
                _ = fi10
                _ = v11
            }
            switch "foo bar" {
            case "foo bar":
                v14 := GenS2{F0: 18446744073709551615, F1: ""}
                _ = v14
            case "":
                v4[(("\n\t" + "a") + ("foo bar" + ""))] = 65535
            case "a":
                v1 = v1
            default:
                var v15 uint32 = 2147483647
                _ = v15
            }
        }
    default:
        if false {
            if false {
                v1 += v1
                for fi16 := 0; fi16 < 2; fi16++ {
                    v3[1] = ("hello" + "foo bar")
                    _ = fi16
                }
                for fi17 := 0; fi17 < 4; fi17++ {
                    v1 = (v1 + 1)
                    _ = fi17
                }
            }
            v1 = -9223372036854775808
        }
    }
    fmt.Println(v1)
    for fi18 := 0; fi18 < 16; fi18++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi18
    }
    var v19 int16 = 42
    if (0 != 42) {
        if false {
            switch 127 {
            case -2147483648:
                switch "" {
                case "hello":
                    if true {
                        switch 2147483647 {
                        case -9223372036854775808:
                            var v20 int16 = (v19 & v19)
                            _ = v20
                        case 9223372036854775807:
                            if true {
                                v21 := GenS3{F0: -644}
                                var v22 uint = 1
                                _ = v21
                                _ = v22
                            }
                        case 4:
                            v23 := GenS4{F0: 917, F1: 0}
                            _ = v23
                        }
                        var v24 bool = ("" == ("a" + ((("\n\t" + "\n\t") + ("hello" + (((("" + (("foo bar" + "foo bar") + "\n\t")) + (("hello" + (((("a" + "hello") + "a") + ("" + "")) + "\n\t")) + (("hello" + ("" + "hello")) + ("a" + "foo bar")))) + "") + ((("a" + "foo bar") + ("a" + ("foo bar" + (("\n\t" + ("" + ("foo bar" + ""))) + (("hello" + "hello") + ""))))) + "")))) + ("" + "hello"))))
                        _ = len(v3)
                        _ = v24
                    } else {
                        v1 = v1
                        switch 1 {
                        case 994:
                            if true {
                                if (65 <= 1) {
                                    fmt.Println(v19, v19)
                                    var v25 float32 = (-2.25 * 0.1)
                                    var v26 bool = (!false)
                                    _ = v25
                                    _ = v26
                                }
                            }
                        case 715:
                            switch 0 {
                            case 42:
                                v1 = 42
                            default:
                                v19 -= v19
                            }
                        case 127:
                            var v27 int8 = -128
                            _ = v27
                        }
                        var v28 bool = (!(42 < 42))
                        _ = v28
                    }
                case "a":
                    var v29 uint64 = 711
                    _ = v29
                case "\n\t":
                    var v30 int = 408
                    _ = v30
                default:
                    for fi31 := 0; fi31 < 20; fi31++ {
                        var v32 string = "\n\t"
                        v1 = v1
                        _ = fi31
                        _ = v32
                    }
                }
            default:
                if false {
                    if (4294967295 < 128) {
                        var v33 float64 = ((0.1 - ((float64(((1e10 - ((((0.0 / (3.14159 + (-2.25 / 0.1))) * float32(1.5)) + (-2.25 / (-2.25 * 0.1))) + (0.0 - 1.5))) / (((0.0 / 1.0) + 1e10) / 3.14159))) - 0.1) - (0.0 / (float64(3.14159) * (((3.14159 - 1e10) / 1.5) * ((1e10 - (0.0 * (0.0 * float64(3.14159)))) / -2.25)))))) / float64(1e10))
                        v1 += v1
                        v1 += (v1 >> ((255) & 63))
                        _ = v33
                    }
                    if (-1 <= -73) {
                        switch ("a" + "foo bar") {
                        case "\n\t":
                            v1 = v1
                        }
                        for fi34 := 0; fi34 < 2; fi34++ {
                            v19 = 42
                            v35 := GenS3{F0: (fi34 + (fi34 << ((65535) & 63)))}
                            fmt.Println(v19, fi34)
                            _ = fi34
                            _ = v35
                        }
                    } else {
                        v36 := v4[("foo bar" + "\n\t")]
                        _ = v36
                    }
                }
            }
            switch int(v19) {
            case 128:
                var v37 float64 = (0.1 * ((3.14159 / (1e10 / 1.0)) / ((float64(-2.25) * (1e10 * (0.0 * -2.25))) - 1e10)))
                _ = v37
            case 32768:
                switch "" {
                case "\n\t":
                    v38 := GenS5{F0: false, F1: 255, F2: 193}
                    _ = v38
                case "hello":
                    fmt.Println(v19, v1)
                }
            case 286:
                var v39 int8 = -128
                _ = v39
            }
        }
        for fi40 := 0; fi40 < 17; fi40++ {
            var v41 uint32 = 65535
            if (32767 != 1) {
                switch "hello" {
                case "hello":
                    for fi42 := 0; fi42 < 6; fi42++ {
                        switch 9223372036854775807 {
                        case 9223372036854775807:
                            if (!((true && (uint(v41) > 65535)) && true)) {
                                switch 42 {
                                case 42:
                                    if false {
                                        if (1 < 0) {
                                            fmt.Println(fi40, fi42, v1)
                                            var v43 int64 = (v1 << ((4294967295) & 63))
                                            if false {
                                                var v44 bool = false
                                                _ = v44
                                            }
                                            fmt.Println(v19, fi42)
                                            fmt.Println(fi42)
                                            _ = v43
                                        } else {
                                            if false {
                                                if (true && true) {
                                                    var v45 uint8 = 127
                                                    var v46 string = ("hello" + "hello")
                                                    _ = v45
                                                    _ = v46
                                                }
                                                v47 := GenS6{F0: -899, F1: 0, F2: 42}
                                                if (0 >= 65535) {
                                                    fmt.Println(fi42)
                                                    var v48 int64 = v1
                                                    fmt.Println(v41, v41, fi40)
                                                    _ = v48
                                                }
                                                switch "" {
                                                case "":
                                                    var v49 int64 = (v1 >> ((2) & 63))
                                                    _ = v49
                                                case "\n\t":
                                                    var v50 uint16 = 2
                                                    _ = v50
                                                default:
                                                    if (false && true) {
                                                        fmt.Println(v41, v1)
                                                    }
                                                }
                                                if ("foo bar" < (((("foo bar" + "a") + "foo bar") + "") + "\n\t")) {
                                                    v4[("hello" + ("\n\t" + ""))] = v41
                                                    v19 = v19
                                                    if true {
                                                        v51 := GenS1{F0: v19, F1: 1, F2: "hello"}
                                                        _ = v51
                                                    } else {
                                                        var v52 uint64 = 42
                                                        switch fi42 {
                                                        case 0:
                                                            if len(v2) > 0 {
                                                                v2[0] = "\n\t"
                                                            }
                                                        case -9223372036854775808:
                                                            var v53 int64 = v1
                                                            _ = v53
                                                        case 405:
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        }
                                                        _ = v52
                                                    }
                                                } else {
                                                    switch fi40 {
                                                    case -575:
                                                        switch fi42 {
                                                        case 4294967295:
                                                            if (1e10 != -2.25) {
                                                                v1 += (v1 >> ((42) & 63))
                                                                if true {
                                                                    v54 := GenS6{F0: (v1 * (v1 & (v1 << ((4294967295) & 63)))), F1: 4294967295, F2: uint(v41)}
                                                                    _ = v54
                                                                }
                                                                var v55 float64 = -2.25
                                                                v19 -= 407
                                                                _ = v55
                                                            } else {
                                                                v1 -= (v1 % ((v1) | 1))
                                                                v1 += (v1 % ((int64(fi42)) | 1))
                                                            }
                                                        default:
                                                            var v56 float32 = 1.5
                                                            _ = v56
                                                        }
                                                    }
                                                    var v57 uint64 = 42
                                                    v1 = (v1 % ((v1) | 1))
                                                    if (false || (false || (true && (((0.1 + 3.14159) / (-2.25 * (0.1 * -2.25))) > 0.1)))) {
                                                        v58 := GenS3{F0: (fi42 + fi42)}
                                                        var v59 uint16 = 255
                                                        switch "\n\t" {
                                                        case "a":
                                                            var v60 uint = 65535
                                                            _ = v60
                                                        case "hello":
                                                            var v61 uint = 382
                                                            _ = v61
                                                        default:
                                                            fmt.Println(v59, fi40)
                                                        }
                                                        _ = v58
                                                        _ = v59
                                                    }
                                                    v62 := GenS1{F0: (v19 | v19), F1: (^(v41)), F2: ("hello" + ("\n\t" + "foo bar"))}
                                                    switch ("foo bar" + "") {
                                                    case "\n\t":
                                                        var v63 uint32 = (v41 | (v41 + v41))
                                                        _ = v63
                                                    }
                                                    var v64 float64 = 3.14159
                                                    v57 -= (v57 >> ((148) & 63))
                                                    var v65 uint32 = (v41 | v41)
                                                    _ = v57
                                                    _ = v62
                                                    _ = v64
                                                    _ = v65
                                                }
                                                switch ("" + ("hello" + "")) {
                                                case "":
                                                    v66 := GenS1{F0: -1, F1: v41, F2: (("hello" + "a") + "foo bar")}
                                                    switch fi40 {
                                                    case 244:
                                                        v67 := GenS7{F0: 1}
                                                        _ = v67
                                                    case 4294967295:
                                                        var v68 float32 = 0.0
                                                        _ = v68
                                                    case 32767:
                                                        fmt.Println(fi40, fi42, v41)
                                                        v69 := v3[v41]
                                                        var v70 float32 = (1.5 * 0.1)
                                                        _ = v69
                                                        _ = v70
                                                    }
                                                    if (1000 < (-(v19))) {
                                                        v41 += (v41 >> ((65535) & 63))
                                                    } else {
                                                        v71 := GenS7{F0: 42}
                                                        switch fi40 {
                                                        case 4294967295:
                                                            v19 += int16(v1)
                                                        case 2147483647:
                                                            var v72 string = (("\n\t" + "") + "hello")
                                                            _ = v72
                                                        case 0:
                                                            var v73 float64 = (0.1 + -2.25)
                                                            _ = v73
                                                        }
                                                        _ = v71
                                                    }
                                                    _ = v66
                                                default:
                                                    if len(v2) > 0 {
                                                        v2[0] = ""
                                                    }
                                                }
                                                _ = v47
                                            }
                                            v2 = append(v2, ("a" + "foo bar"))
                                        }
                                    } else {
                                        v41 -= 265
                                        v1 += int64(v19)
                                    }
                                default:
                                    switch "\n\t" {
                                    case "\n\t":
                                        fmt.Println(v19, v41, fi42)
                                    default:
                                        v1 = -9223372036854775808
                                        v74 := GenS7{F0: -128}
                                        v19 += v19
                                        _ = v74
                                    }
                                }
                                v1 = v1
                                v75 := GenS8{F0: 0}
                                delete(v4, "foo bar")
                                _ = v75
                            } else {
                                if true {
                                    switch fi42 {
                                    case 32768:
                                        switch fi42 {
                                        case 4294967295:
                                            v2 = append(v2, "foo bar")
                                        }
                                    case -925:
                                        var v76 uint16 = 868
                                        v41 = 698
                                        v4["hello"] = (v41 | v41)
                                        var v77 int8 = -1
                                        _ = v76
                                        _ = v77
                                    case 0:
                                        _ = len(v2)
                                    default:
                                        if (421 == v41) {
                                            fmt.Println(fi40, v1)
                                            fmt.Println(v1)
                                            fmt.Println(fi40, v19, fi42)
                                        }
                                    }
                                    v19 -= (v19 / ((v19) | 1))
                                }
                            }
                        case -1:
                            var v78 uint64 = 2
                            _ = v78
                        default:
                            fmt.Println(v41)
                            v19 = (-(v19))
                            fmt.Println(fi40, v19)
                            v41 = 880
                            if ((1 == 127) || (!false)) {
                                var v79 int16 = v19
                                fmt.Println(v19)
                                if false {
                                    v80 := GenS8{F0: 42}
                                    switch "" {
                                    case "hello":
                                        v81 := GenS9{F0: false}
                                        _ = v81
                                    case "foo bar":
                                        v79 = int16(v1)
                                    default:
                                        v2 = append(v2, (("\n\t" + "\n\t") + "hello"))
                                    }
                                    var v82 float64 = ((0.0 + (0.0 - (-2.25 - (1e10 / 1.0)))) * (0.0 / 0.1))
                                    _ = v80
                                    _ = v82
                                }
                                _ = v79
                            }
                            fmt.Println(v1, v41)
                        }
                        _ = fi42
                    }
                }
                v83 := GenS10{F0: 1, F1: v19, F2: 709}
                v1 += -9223372036854775808
                _ = v83
            } else {
                _ = len(v2)
            }
            _ = fi40
            _ = v41
        }
    }
    if true {
        switch "a" {
        case "hello":
            v84 := v4[("hello" + "foo bar")]
            _ = v84
        default:
            for fi85 := 0; fi85 < 2; fi85++ {
                v1 = 0
                _ = fi85
            }
        }
        if true {
            v86 := GenS6{F0: -9223372036854775808, F1: 828, F2: 564}
            fmt.Println(v19, v19)
            for fi87 := 0; fi87 < 20; fi87++ {
                v88 := GenS10{F0: (v19 * 998), F1: (-(v19)), F2: 18446744073709551615}
                for fi89 := 0; fi89 < 5; fi89++ {
                    var v90 int16 = v19
                    _ = fi89
                    _ = v90
                }
                _ = fi87
                _ = v88
            }
            _ = v86
        } else {
            v1 = (v1 | (^(v1)))
            var v91 float32 = 3.14159
            if true {
                v19 += (v19 ^ (v19 + -1))
                for fi92 := 0; fi92 < 15; fi92++ {
                    for fi93 := 0; fi93 < 7; fi93++ {
                        v91 -= v91
                        var v94 uint = 356
                        _ = fi93
                        _ = v94
                    }
                    switch "hello" {
                    case "\n\t":
                        v95 := GenS11{F0: 65535}
                        _ = v95
                    default:
                        if true {
                            for fi96 := 0; fi96 < 9; fi96++ {
                                if (true && false) {
                                    v4[(("\n\t" + "") + ("" + ("\n\t" + ("a" + ""))))] = 2
                                    if true {
                                        if (!false) {
                                            fmt.Println(v91, fi96, v91)
                                            switch ("" + (("a" + "foo bar") + "")) {
                                            case "hello":
                                                var v97 uint64 = 212
                                                _ = v97
                                            case "":
                                                if true {
                                                    v4[("foo bar" + ("hello" + ("a" + "a")))] = 1
                                                } else {
                                                    if len(v2) > 0 {
                                                        v2[0] = "foo bar"
                                                    }
                                                    fmt.Println(fi96)
                                                }
                                            }
                                        }
                                        v91 = v91
                                        v91 = (v91 - ((1.5 + 0.0) * ((v91 + (((v91 - 1e10) / ((1.5 / -2.25) * float32(1.5))) / float32(float64(((v91 * ((v91 - v91) / 1.5)) * v91))))) * (v91 / v91))))
                                        switch (-(fi92)) {
                                        case -9223372036854775808:
                                            switch (fi92 ^ (-(fi96))) {
                                            case 128:
                                                if ((v1 ^ (v1 << ((18446744073709551615) & 63))) <= v1) {
                                                    v98 := GenS2{F0: 18446744073709551615, F1: "foo bar"}
                                                    delete(v3, 2147483647)
                                                    _ = v98
                                                } else {
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        var v99 int8 = 0
                                                        _ = v99
                                                    }
                                                    v1 += (v1 * (-(v1)))
                                                    _ = len(v3)
                                                    switch fi96 {
                                                    case -173:
                                                        if (fi96 != fi96) {
                                                            if false {
                                                                var v100 float32 = v91
                                                                var v101 float64 = float64(v100)
                                                                fmt.Println(v19, v1)
                                                                _ = v100
                                                                _ = v101
                                                            } else {
                                                                fmt.Println(v91, fi96)
                                                            }
                                                        }
                                                    case -2147483648:
                                                        v3[2] = ((("a" + "a") + "") + ("" + "hello"))
                                                    }
                                                }
                                            }
                                        case -2147483648:
                                            if (true || (70 <= 109)) {
                                                if true {
                                                    delete(v3, 253)
                                                }
                                                var v102 string = "foo bar"
                                                _ = v102
                                            }
                                        }
                                    }
                                    switch fi92 {
                                    case 9223372036854775807:
                                        if false {
                                            if (18446744073709551615 >= 4294967295) {
                                                var v103 int32 = -217
                                                _ = len(v3)
                                                _ = v103
                                            }
                                            var v104 uint16 = 255
                                            _ = v104
                                        }
                                    case 0:
                                        if len(v2) > 0 {
                                            v2[0] = "hello"
                                        }
                                    default:
                                        var v105 int32 = -1
                                        _ = v105
                                    }
                                } else {
                                    v19 -= v19
                                    v106 := GenS2{F0: 1, F1: "\n\t"}
                                    fmt.Println(v91)
                                    _ = v106
                                }
                                var v107 int64 = (v1 * (v1 ^ (v1 >> ((0) & 63))))
                                _ = fi96
                                _ = v107
                            }
                            v1 += -9223372036854775808
                        }
                    }
                    _ = fi92
                }
            } else {
                if true {
                    if (2 <= 0) {
                        v91 = v91
                        if true {
                            var v108 int = 32767
                            if (!false) {
                                fmt.Println(v19)
                                v109 := GenS12{F0: 0, F1: 112, F2: 0}
                                for fi110 := 0; fi110 < 14; fi110++ {
                                    v108 += v108
                                    _ = fi110
                                }
                                _ = v109
                            } else {
                                var v111 uint64 = 2
                                _ = v111
                            }
                            _ = v108
                        } else {
                            var v112 int16 = 42
                            v113 := GenS5{F0: (572 < -1), F1: 2, F2: 65}
                            v112 = -118
                            _ = v112
                            _ = v113
                        }
                        fmt.Println(v19, v19)
                    }
                    v19 += v19
                } else {
                    v1 -= v1
                    var v114 int16 = 1
                    _ = v114
                }
                v91 = v91
            }
            _ = v91
        }
    } else {
        var v115 string = "a"
        v116 := GenS13{F0: 128, F1: 42, F2: 42}
        fmt.Println(v19, v1)
        _ = v115
        _ = v116
    }
    var v117 float32 = (0.1 - (0.0 / 3.14159))
    for fi118 := 0; fi118 < 13; fi118++ {
        fmt.Println(v117, v19)
        if true {
            v119 := GenS12{F0: 0, F1: 42, F2: v1}
            var v120 uint64 = 187
            var v121 int8 = 42
            _ = v119
            _ = v120
            _ = v121
        }
        if true {
            _ = len(v2)
        } else {
            fmt.Println(v1, v19)
            v2 = append(v2, (("\n\t" + "a") + "a"))
            var v122 uint = 4294967295
            _ = v122
        }
        _ = fi118
    }
    v123 := v3[10]
    fmt.Println("v1:", v1)
    fmt.Println("v117:", v117)
    fmt.Println("v123:", v123)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v117
    _ = v123
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
