package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 int64; F1 uint32; F2 int16 }
type GenS3 struct { F0 string; F1 bool }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int8; F1 int; F2 int16 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int32; F1 uint }
type GenS8 struct { F0 uint16; F1 int8; F2 bool }
type GenS9 struct { F0 int32; F1 uint8 }
type GenS10 struct { F0 bool; F1 int64 }
type GenS11 struct { F0 int16 }

func main() {
    var v1 uint = 1
    var v2 float64 = 1.5
    var v3 uint32 = 2147483647
    var v4 int32 = -2147483648
    v5 := []int{-1, -9223372036854775808}
    if len(v5) > 0 {
        v5[0] = 290
    }
    if (1e10 >= -2.25) {
        if (1 <= 4294967295) {
            if len(v5) > 0 {
                v5[0] = -654
            }
            fmt.Println(v2, v3, v2)
        } else {
            v4 -= int32(v1)
        }
        if true {
            for fi6 := 0; fi6 < 19; fi6++ {
                fmt.Println(fi6, v2, fi6)
                _ = fi6
            }
            if true {
                if ((1 == 127) || false) {
                    v7 := GenS1{F0: "a", F1: 836}
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v8 := GenS1{F0: ("foo bar" + "a"), F1: 991}
                    _ = v7
                    _ = v8
                }
                v2 -= (v2 - 1e10)
                if len(v5) > 0 {
                    v5[0] = 162
                }
            } else {
                if (929 > 2147483647) {
                    v9 := GenS1{F0: "\n\t", F1: (v1 & (v1 + 18446744073709551615))}
                    _ = v9
                } else {
                    for fi10 := 0; fi10 < 1; fi10++ {
                        if (5 < 127) {
                            var v11 bool = true
                            for fi12 := 0; fi12 < 3; fi12++ {
                                v13 := GenS2{F0: 670, F1: (v3 ^ 616), F2: 475}
                                _ = fi12
                                _ = v13
                            }
                            _ = v11
                        }
                        v2 -= 0.1
                        fmt.Println(v1)
                        _ = fi10
                    }
                }
            }
            v1 = (v1 + v1)
        } else {
            if false {
                if true {
                    v14 := GenS2{F0: 9223372036854775807, F1: uint32(v1), F2: int16(v3)}
                    if true {
                        for fi15 := 0; fi15 < 17; fi15++ {
                            var v16 uint = 255
                            if len(v5) > 0 {
                                v5[0] = fi15
                            }
                            if true {
                                var v17 float32 = 1.5
                                if (v3 != 2147483648) {
                                    v3 -= (v3 - v3)
                                    switch ("foo bar" + ("\n\t" + "")) {
                                    case "":
                                        _ = len(v5)
                                    case "foo bar":
                                        v1 = v16
                                    case "hello":
                                        if false {
                                            var v18 float64 = (v2 / (1.5 - v2))
                                            _ = v18
                                        } else {
                                            var v19 uint8 = 255
                                            v3 -= v3
                                            if false {
                                                if (false && false) {
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                    var v20 bool = false
                                                    var v21 uint16 = 2
                                                    _ = v20
                                                    _ = v21
                                                } else {
                                                    v1 = (v16 ^ 932)
                                                    switch "" {
                                                    case "\n\t":
                                                        v3 = v3
                                                    case "":
                                                        switch "a" {
                                                        case "a":
                                                            v5 = append(v5, -150)
                                                        case "":
                                                            for fi22 := 0; fi22 < 9; fi22++ {
                                                                if true {
                                                                    var v23 int32 = 42
                                                                    v19 -= 100
                                                                    _ = v23
                                                                }
                                                                var v24 float32 = v17
                                                                _ = fi22
                                                                _ = v24
                                                            }
                                                        case "hello":
                                                            if (2 != 720) {
                                                                fmt.Println(v16)
                                                                if true {
                                                                    v4 -= v4
                                                                    var v25 int = fi15
                                                                    _ = v25
                                                                }
                                                                v5 = append(v5, (fi15 - (fi15 - (fi15 * fi15))))
                                                            }
                                                        }
                                                    case "foo bar":
                                                        if ((!(!(fi15 > int(v4)))) && false) {
                                                            var v26 int = 4294967295
                                                            switch (-(fi15)) {
                                                            case 128:
                                                                var v27 uint32 = (^(v3))
                                                                _ = v27
                                                            case 42:
                                                                for fi28 := 0; fi28 < 16; fi28++ {
                                                                    switch v26 {
                                                                    case -1:
                                                                        v4 = v4
                                                                    case 32767:
                                                                        var v29 int8 = 16
                                                                        _ = v29
                                                                    default:
                                                                        if ((v1 ^ 1) <= (v16 ^ v16)) {
                                                                            var v30 int32 = v4
                                                                            if len(v5) > 0 {
                                                                                v5[0] = fi28
                                                                            }
                                                                            var v31 string = ""
                                                                            v26 = fi28
                                                                            _ = v30
                                                                            _ = v31
                                                                        }
                                                                    }
                                                                    _ = len(v5)
                                                                    _ = fi28
                                                                }
                                                            }
                                                            _ = v26
                                                        }
                                                    default:
                                                        v16 = (v16 >> ((v1) & 63))
                                                    }
                                                    switch fi15 {
                                                    case 2147483647:
                                                        v16 = (v1 * (v16 - v16))
                                                    }
                                                }
                                                v16 += (^(v1))
                                            } else {
                                                v1 = v16
                                                v32 := GenS3{F0: ("\n\t" + "hello"), F1: (v19 == v19)}
                                                _ = v32
                                            }
                                            _ = v19
                                        }
                                    }
                                    var v33 uint32 = v3
                                    _ = v33
                                }
                                _ = v17
                            } else {
                                v4 += 42
                            }
                            _ = fi15
                            _ = v16
                        }
                    }
                    switch 127 {
                    case -142:
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    case 4294967295:
                        if false {
                            if false {
                                var v34 uint32 = 882
                                for fi35 := 0; fi35 < 2; fi35++ {
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    for fi36 := 0; fi36 < 6; fi36++ {
                                        v3 = v34
                                        if len(v5) > 0 {
                                            v5[0] = 1
                                        }
                                        _ = fi36
                                    }
                                    _ = fi35
                                }
                                v5 = append(v5, 726)
                                _ = v34
                            }
                            v5 = append(v5, 0)
                        }
                    case -777:
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    default:
                        if (-1 <= 9223372036854775807) {
                            v5 = append(v5, -364)
                        }
                    }
                    _ = v14
                } else {
                    for fi37 := 0; fi37 < 6; fi37++ {
                        v3 -= 517
                        for fi38 := 0; fi38 < 13; fi38++ {
                            _ = len(v5)
                            v39 := GenS4{F0: 42}
                            _ = fi38
                            _ = v39
                        }
                        _ = fi37
                    }
                    if (true && (1000 > 113)) {
                        for fi40 := 0; fi40 < 14; fi40++ {
                            v2 += v2
                            v41 := GenS5{F0: -1, F1: (fi40 - (fi40 / (((fi40 % ((fi40) | 1))) | 1))), F2: 1}
                            _ = fi40
                            _ = v41
                        }
                        v2 += v2
                        var v42 uint16 = 1
                        _ = v42
                    } else {
                        fmt.Println(v3)
                        _ = len(v5)
                        var v43 float32 = ((-2.25 * float32(v2)) - 0.0)
                        _ = v43
                    }
                }
                if true {
                    fmt.Println(v1, v4, v4)
                } else {
                    v44 := GenS6{F0: v1}
                    var v45 uint64 = 2
                    fmt.Println(v2, v1, v4)
                    _ = v44
                    _ = v45
                }
                v4 = (v4 << ((18446744073709551615) & 63))
            } else {
                var v46 int64 = 0
                if ((-(v4)) <= v4) {
                    for fi47 := 0; fi47 < 19; fi47++ {
                        if (v4 < (v4 / (((v4 + (v4 | v4))) | 1))) {
                            if ((!true) || (true || ((v3 | 2147483647) == v3))) {
                                if ((((v3 & v3) > (v3 | 2)) && (false && false)) && (1 >= 425)) {
                                    switch -1 {
                                    case 42:
                                        for fi48 := 0; fi48 < 16; fi48++ {
                                            switch "hello" {
                                            case "":
                                                _ = len(v5)
                                            }
                                            if len(v5) > 0 {
                                                v5[0] = (fi47 ^ fi47)
                                            }
                                            var v49 int8 = -56
                                            _ = fi48
                                            _ = v49
                                        }
                                    case -343:
                                        var v50 int = -589
                                        _ = v50
                                    }
                                } else {
                                    v46 -= (^(v46))
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    fmt.Println(fi47, v46)
                                }
                                v51 := GenS1{F0: "hello", F1: (v1 << ((42) & 63))}
                                _ = v51
                            } else {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                var v52 uint64 = 4294967295
                                _ = v52
                            }
                            v53 := GenS7{F0: v4, F1: 65535}
                            _ = v53
                        }
                        _ = fi47
                    }
                }
                v54 := GenS4{F0: (v3 >> ((2) & 63))}
                _ = v46
                _ = v54
            }
        }
        v4 = v4
    }
    switch 42 {
    case 1:
        for fi55 := 0; fi55 < 9; fi55++ {
            switch ("hello" + ("hello" + "a")) {
            case "\n\t":
                var v56 int16 = 1
                _ = v56
            case "hello":
                if ((true || ("\n\t" != "hello")) && true) {
                    for fi57 := 0; fi57 < 18; fi57++ {
                        var v58 uint64 = 42
                        v5 = append(v5, fi57)
                        _ = fi57
                        _ = v58
                    }
                    v59 := GenS3{F0: ("\n\t" + (("foo bar" + "\n\t") + ("a" + "a"))), F1: ((uint64(v1) > 724) && true)}
                    _ = v59
                }
            case "foo bar":
                v1 -= v1
            default:
                var v60 uint8 = 127
                _ = v60
            }
            _ = fi55
        }
    case -752:
        v4 = 1
    }
    var v61 int64 = 17
    if true {
        var v62 uint32 = (^(v3))
        if (false && false) {
            v61 = v61
            var v63 float64 = (v2 + (-2.25 + (v2 + (v2 / (v2 - float64(0.1))))))
            _ = v63
        } else {
            switch -9223372036854775808 {
            case 32768:
                var v64 uint32 = v3
                _ = v64
            }
            v65 := GenS6{F0: (^(v1))}
            _ = v65
        }
        switch -2147483648 {
        case 127:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        case 42:
            v2 += v2
        default:
            fmt.Println(v1, v3, v61)
        }
        _ = v62
    } else {
        switch "foo bar" {
        case "":
            if len(v5) > 0 {
                v5[0] = 4294967295
            }
        }
    }
    switch -1 {
    case 2147483647:
        v5 = append(v5, 9223372036854775807)
    case 857:
        for fi66 := 0; fi66 < 12; fi66++ {
            var v67 int32 = (v4 + v4)
            var v68 int = 32767
            _ = fi66
            _ = v67
            _ = v68
        }
    case -2147483648:
        var v69 uint16 = 32767
        _ = v69
    default:
        v3 -= 128
    }
    fmt.Println(v3)
    for fi70 := 0; fi70 < 8; fi70++ {
        for fi71 := 0; fi71 < 19; fi71++ {
            if (false && ((v3 & v3) < v3)) {
                fmt.Println(v4, fi71, v3)
                v72 := GenS2{F0: v61, F1: (v3 ^ v3), F2: int16(v61)}
                switch ("foo bar" + "\n\t") {
                case "hello":
                    var v73 uint16 = 1
                    _ = v73
                }
                switch (fi70 | (fi71 * (fi70 % ((fi71) | 1)))) {
                case 0:
                    var v74 int64 = -1
                    _ = v74
                }
                _ = v72
            } else {
                v75 := GenS5{F0: 1, F1: fi71, F2: 1}
                _ = v75
            }
            switch (fi70 >> ((v1) & 63)) {
            case 32768:
                switch 987 {
                case -9223372036854775808:
                    v4 -= v4
                case 369:
                    v4 = 0
                    v2 = float64(1e10)
                    switch int(v4) {
                    case -953:
                        var v76 int32 = v4
                        if false {
                            v4 = (v76 / (((v76 << (((v1 & (v1 >> (((v1 - v1)) & 63)))) & 63))) | 1))
                            v4 -= (v76 * v4)
                            var v77 uint = v1
                            v77 -= (v1 | v1)
                            v77 -= v77
                            _ = len(v5)
                            _ = v77
                        } else {
                            v78 := GenS7{F0: (v4 ^ (-(v4))), F1: (v1 % ((512) | 1))}
                            if true {
                                if true {
                                    if true {
                                        fmt.Println(v2, fi71)
                                        fmt.Println(fi70, fi70, fi70)
                                    } else {
                                        switch (fi71 << (((v1 | (v1 / ((v1) | 1)))) & 63)) {
                                        case -2147483648:
                                            switch fi71 {
                                            case 128:
                                                v79 := GenS3{F0: "", F1: false}
                                                _ = v79
                                            case 0:
                                                v5 = append(v5, (fi70 & fi70))
                                                fmt.Println(v76, v4, v76)
                                            case -819:
                                                v76 = v4
                                            }
                                        case 42:
                                            var v80 uint = 2
                                            _ = v80
                                        default:
                                            var v81 uint8 = 128
                                            _ = v81
                                        }
                                        v82 := GenS8{F0: 633, F1: 1, F2: false}
                                        _ = v82
                                    }
                                    switch fi70 {
                                    case 143:
                                        if true {
                                            switch fi71 {
                                            case -233:
                                                fmt.Println(fi71, fi70, v4)
                                            case -26:
                                                if true {
                                                    var v83 bool = (1 > (v61 + (v61 ^ (v61 & 9223372036854775807))))
                                                    _ = v83
                                                }
                                            default:
                                                if len(v5) > 0 {
                                                    v5[0] = fi71
                                                }
                                            }
                                            var v84 bool = true
                                            fmt.Println(fi71, v84, v2)
                                            _ = v84
                                        }
                                    }
                                    if ((0.1 <= 1.5) || false) {
                                        fmt.Println(v2, v61)
                                        switch (fi70 - fi71) {
                                        case 127:
                                            var v85 uint = v1
                                            switch "\n\t" {
                                            case "foo bar":
                                                var v86 float32 = (1.5 * (3.14159 * -2.25))
                                                _ = v86
                                            case "hello":
                                                switch fi71 {
                                                case -2147483648:
                                                    fmt.Println(v1, v4, v1)
                                                default:
                                                    var v87 bool = ((!((v2 + 3.14159) == 0.1)) && (172 > uint8(v3)))
                                                    _ = v87
                                                }
                                            default:
                                                v76 += -396
                                            }
                                            v88 := GenS9{F0: -1, F1: 42}
                                            _ = v85
                                            _ = v88
                                        case 42:
                                            v3 = 4294967295
                                        default:
                                            if false {
                                                if (0.0 > 0.1) {
                                                    v3 = (v3 % (((v3 & v3)) | 1))
                                                    if (0 < 65535) {
                                                        _ = len(v5)
                                                        if len(v5) > 0 {
                                                            v5[0] = (fi71 % (((fi70 | fi70)) | 1))
                                                        }
                                                    } else {
                                                        v89 := GenS10{F0: true, F1: (v61 * v61)}
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                        _ = v89
                                                    }
                                                    fmt.Println(fi70, v4, fi70)
                                                }
                                                if ((!((2 > 127) || (65535 < 254))) && ("" == "foo bar")) {
                                                    v2 = (3.14159 * 1e10)
                                                    v5 = append(v5, int(v61))
                                                    if true {
                                                        v3 += (v3 | 653)
                                                        fmt.Println(v3, v2)
                                                    } else {
                                                        var v90 bool = (0 != -32768)
                                                        _ = v90
                                                    }
                                                }
                                                if len(v5) > 0 {
                                                    v5[0] = fi70
                                                }
                                            }
                                        }
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    }
                                }
                            } else {
                                if true {
                                    v91 := GenS5{F0: -17, F1: (fi70 | fi70), F2: 1}
                                    if ((v61 << ((65535) & 63)) < v61) {
                                        var v92 uint32 = 0
                                        _ = len(v5)
                                        v92 += 730
                                        _ = v92
                                    } else {
                                        var v93 string = "\n\t"
                                        var v94 float64 = 1e10
                                        if true {
                                            v93 = ""
                                            v1 -= (v1 - (v1 ^ v1))
                                        }
                                        _ = v93
                                        _ = v94
                                    }
                                    if ((!true) && (!(-213 == (fi71 % ((int(v1)) | 1))))) {
                                        _ = len(v5)
                                        var v95 uint64 = 42
                                        v96 := GenS11{F0: -122}
                                        _ = v95
                                        _ = v96
                                    } else {
                                        var v97 float64 = (1e10 / v2)
                                        v98 := GenS4{F0: v3}
                                        _ = v97
                                        _ = v98
                                    }
                                    _ = v91
                                }
                            }
                            _ = v78
                        }
                        _ = v76
                    case -399:
                        switch "" {
                        case "hello":
                            fmt.Println(fi71)
                        case "a":
                            var v99 uint32 = v3
                            _ = v99
                        case "foo bar":
                            if len(v5) > 0 {
                                v5[0] = (^(fi71))
                            }
                        }
                    default:
                        if len(v5) > 0 {
                            v5[0] = -9
                        }
                    }
                default:
                    v100 := GenS1{F0: ("a" + "a"), F1: (v1 * (v1 * v1))}
                    _ = v100
                }
            case 127:
                var v101 float64 = 1.5
                _ = v101
            }
            var v102 float32 = ((float32(((v2 / 1.0) / v2)) / -2.25) - -2.25)
            _ = fi71
            _ = v102
        }
        _ = fi70
    }
    if (!("hello" > ("\n\t" + ""))) {
        var v103 int = 239
        switch "a" {
        case "a":
            var v104 uint8 = 127
            _ = v104
        case "":
            v3 -= (v3 + 2)
        default:
            v4 = v4
        }
        _ = v103
    } else {
        v105 := GenS11{F0: 42}
        _ = v105
    }
    switch "foo bar" {
    case "hello":
        switch "a" {
        case "\n\t":
            fmt.Println(v4, v4)
        case "hello":
            for fi106 := 0; fi106 < 15; fi106++ {
                switch ("foo bar" + ("" + "hello")) {
                case "hello":
                    if true {
                        for fi107 := 0; fi107 < 8; fi107++ {
                            var v108 float64 = 1.5
                            _ = fi107
                            _ = v108
                        }
                        var v109 uint8 = 0
                        _ = v109
                    }
                }
                for fi110 := 0; fi110 < 4; fi110++ {
                    var v111 uint = (v1 * 325)
                    switch "hello" {
                    case "foo bar":
                        v5 = append(v5, fi110)
                    case "":
                        var v112 int8 = 100
                        _ = v112
                    default:
                        var v113 uint16 = 65535
                        _ = v113
                    }
                    _ = fi110
                    _ = v111
                }
                var v114 string = "\n\t"
                _ = fi106
                _ = v114
            }
        case "a":
            if false {
                var v115 int64 = v61
                fmt.Println(v4)
                _ = v115
            }
        default:
            v2 -= v2
        }
    }
    v3 = (^(v3))
    if ("hello" < ("a" + "")) {
        fmt.Println(v61)
        v1 -= 0
        for fi116 := 0; fi116 < 9; fi116++ {
            switch 32768 {
            case 2147483647:
                for fi117 := 0; fi117 < 4; fi117++ {
                    if len(v5) > 0 {
                        v5[0] = (-(fi117))
                    }
                    _ = fi117
                }
            case 4294967295:
                v2 = v2
            case 32768:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
            var v118 float32 = float32((0.0 - float64(float32(v2))))
            v3 += v3
            _ = fi116
            _ = v118
        }
    } else {
        if false {
            switch ((("" + ("foo bar" + "")) + "foo bar") + "") {
            case "hello":
                fmt.Println(v2, v1, v61)
            }
        }
        v1 += v1
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    var v119 uint = (^(v1))
    var v120 string = "a"
    v2 = float64(3.14159)
    v119 = (^(v119))
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v119:", v119)
    fmt.Println("v120:", v120)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v119
    _ = v120
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v61
}
