package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int32 }
type GenS2 struct { F0 string; F1 int64; F2 uint32 }
type GenS3 struct { F0 string; F1 int8; F2 uint8 }
type GenS4 struct { F0 uint32; F1 int32; F2 uint }
type GenS5 struct { F0 uint64; F1 string }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int16; F1 int8 }
type GenS8 struct { F0 uint8; F1 uint64 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int64; F1 int16; F2 int32 }
type GenS11 struct { F0 uint64; F1 int64 }
type GenS12 struct { F0 int16; F1 int8 }
type GenS13 struct { F0 int16 }
func genFn0(p0 uint32, p1 uint64) uint32 {
    p0 = p0
    var v6 int64 = -9223372036854775808
    var v7 uint8 = 42
    v8 := []int{619, 32767, int(p1)}
    if false {
        v8 = append(v8, 128)
        if true {
            if (!(65535 == 2)) {
                v9 := GenS1{F0: 786, F1: (p0 & (p0 << ((42) & 63))), F2: 42}
                switch 954 {
                case 4294967295:
                    _ = len(v8)
                case -96:
                    v8 = append(v8, int(p0))
                }
                if len(v8) > 0 {
                    v8[0] = 128
                }
                _ = v9
            } else {
                if (-993 == 1) {
                    var v10 int16 = -705
                    var v11 int32 = -912
                    _ = v10
                    _ = v11
                }
            }
            switch "" {
            case "hello":
                if true {
                    if true {
                        v12 := GenS1{F0: uint16(p0), F1: p0, F2: -782}
                        _ = v12
                    } else {
                        var v13 uint8 = v7
                        _ = v13
                    }
                } else {
                    p0 += 42
                }
            case "a":
                v7 += (^(v7))
            }
        }
        _ = len(v8)
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return p0
}
func genFn1(p0 uint32) int8 {
    for fi14 := 0; fi14 < 3; fi14++ {
        p0 = genFn0(p0, 0)
        var v15 uint = 405
        _ = fi14
        _ = v15
    }
    if false {
        p0 += (p0 * 250)
        p0 = genFn0(p0, 2)
        fmt.Println(p0)
    } else {
        v16 := []int{9223372036854775807, 32767, -2147483648}
        _ = v16
    }
    switch -1 {
    case -9223372036854775808:
        var v17 int32 = 881
        _ = v17
    case 9223372036854775807:
        p0 = (p0 >> ((uint(p0)) & 63))
    }
    p0 -= p0
    for fi18 := 0; fi18 < 19; fi18++ {
        var v19 float64 = 1.5
        _ = fi18
        _ = v19
    }
    _ = p0
    return -99
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []int{}
    v3 := []string{}
    var v4 int32 = -271
    var v5 float32 = 0.1
    switch -1 {
    case -9223372036854775808:
        v20 := GenS1{F0: 32767, F1: 0, F2: 42}
        _ = v20
    case 127:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case 1:
        if false {
            for fi21 := 0; fi21 < 5; fi21++ {
                switch ("foo bar" + "a") {
                case "\n\t":
                    if true {
                        if len(v3) > 0 {
                            v3[0] = "hello"
                        }
                        switch "a" {
                        case "hello":
                            var v22 bool = false
                            _ = v22
                        default:
                            var v23 uint = 588
                            _ = v23
                        }
                    } else {
                        for fi24 := 0; fi24 < 16; fi24++ {
                            v5 = v5
                            v4 = -739
                            _ = fi24
                        }
                        v25 := GenS1{F0: 7, F1: 1, F2: 2147483647}
                        fmt.Println(v5)
                        _ = v25
                    }
                case "foo bar":
                    switch -1 {
                    case 128:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                default:
                    if false {
                        fmt.Println(fi21, v5)
                    } else {
                        var v26 uint16 = 378
                        var v27 uint = 65535
                        fmt.Println(v27, fi21)
                        _ = v26
                        _ = v27
                    }
                }
                if false {
                    v3 = append(v3, "\n\t")
                    if (genFn1(399) <= -128) {
                        for fi28 := 0; fi28 < 11; fi28++ {
                            v29 := GenS1{F0: 682, F1: 65535, F2: v4}
                            fmt.Println(v5, fi21, fi28)
                            _ = fi28
                            _ = v29
                        }
                    }
                }
                v30 := GenS1{F0: 2, F1: 4294967295, F2: (v4 / ((v4) | 1))}
                _ = fi21
                _ = v30
            }
        } else {
            var v31 int32 = 814
            for fi32 := 0; fi32 < 12; fi32++ {
                var v33 uint16 = 230
                if ((!(128 >= 2)) && false) {
                    if (false && false) {
                        var v34 uint32 = 1
                        v5 = (1.5 - (3.14159 - v5))
                        _ = v34
                    }
                    v33 -= 881
                    if (false || ((float64(-2.25) - 0.1) < (-2.25 + (1.5 + float64(1e10))))) {
                        for fi35 := 0; fi35 < 6; fi35++ {
                            fmt.Println(v5, v4)
                            switch (fi35 * fi32) {
                            case 451:
                                v36 := GenS1{F0: (v33 ^ (v33 % (((v33 / ((v33) | 1))) | 1))), F1: 2147483647, F2: v4}
                                _ = v36
                            case 254:
                                if (30 >= 42) {
                                    var v37 uint = 2
                                    v38 := GenS1{F0: (v33 % ((v33) | 1)), F1: genFn0(83, 42), F2: v31}
                                    _ = v37
                                    _ = v38
                                } else {
                                    v39 := GenS1{F0: 0, F1: 0, F2: (v31 * -2147483648)}
                                    var v40 float64 = ((((0.1 * float64(-2.25)) / (float64(((1e10 / v5) - (v5 - float32((((((1.5 + 1.5) * (1.5 * 0.0)) - 0.1) + 0.0) - 0.1))))) + ((0.1 - 1e10) / 1.0))) / 1e10) - 1.5)
                                    var v41 int32 = 0
                                    _ = v39
                                    _ = v40
                                    _ = v41
                                }
                            }
                            var v42 bool = (-1 >= -9223372036854775808)
                            _ = fi35
                            _ = v42
                        }
                        if true {
                            var v43 float32 = (v5 + (((v5 * (v5 / 1.0)) - (v5 - ((float32(1e10) * float32((0.0 / 1.0))) - -2.25))) - 1.5))
                            switch fi32 {
                            case -863:
                                switch "a" {
                                case "\n\t":
                                    v3 = append(v3, ("hello" + "foo bar"))
                                case "":
                                    fmt.Println(v43)
                                case "hello":
                                    fmt.Println(v43, v5, v31)
                                default:
                                    if (!(-25 > 0)) {
                                        delete(v1, fi32)
                                        switch (fi32 - fi32) {
                                        case 127:
                                            var v44 int16 = 42
                                            _ = v44
                                        case -9223372036854775808:
                                            var v45 int = fi32
                                            _ = v45
                                        }
                                    } else {
                                        var v46 int32 = v31
                                        v47 := GenS1{F0: v33, F1: 826, F2: (^(v31))}
                                        _ = v46
                                        _ = v47
                                    }
                                }
                            case -2147483648:
                                if true {
                                    if ((true || ((!true) && ((genFn0(139, 4294967295) < 2147483648) && true))) || true) {
                                        var v48 float64 = 1e10
                                        if false {
                                            switch fi32 {
                                            case -2147483648:
                                                fmt.Println(v48, v48)
                                            case 73:
                                                v49 := GenS2{F0: "hello", F1: 9223372036854775807, F2: 2}
                                                _ = v49
                                            default:
                                                v33 = (v33 >> ((52) & 63))
                                            }
                                        }
                                        _ = v48
                                    } else {
                                        for fi50 := 0; fi50 < 15; fi50++ {
                                            if (v4 == (v31 ^ (v31 ^ v4))) {
                                                v33 = (v33 >> ((65535) & 63))
                                            } else {
                                                switch "foo bar" {
                                                case "a":
                                                    v4 = v31
                                                case "":
                                                    var v51 int64 = 1
                                                    _ = v51
                                                default:
                                                    var v52 uint8 = 128
                                                    _ = v52
                                                }
                                                v53 := GenS3{F0: ("foo bar" + "a"), F1: -1, F2: 1}
                                                _ = v53
                                            }
                                            _ = fi50
                                        }
                                        if (v33 == (v33 << ((0) & 63))) {
                                            if ((fi32 >> ((4294967295) & 63)) > 2147483647) {
                                                v31 = (v31 / (((^(v31))) | 1))
                                                var v54 int = 1
                                                _ = v54
                                            }
                                            switch 0 {
                                            case -2147483648:
                                                v33 = 1
                                            case 9223372036854775807:
                                                fmt.Println(fi32)
                                            default:
                                                if true {
                                                    _ = len(v2)
                                                }
                                            }
                                            v31 = (v4 * v31)
                                        } else {
                                            for fi55 := 0; fi55 < 1; fi55++ {
                                                if (!false) {
                                                    v1[fi55] = 698
                                                    var v56 int8 = 120
                                                    v57 := GenS1{F0: (v33 & v33), F1: 304, F2: -2147483648}
                                                    _ = v56
                                                    _ = v57
                                                }
                                                fmt.Println(v5, fi32, fi32)
                                                var v58 bool = ((0 != 1) && false)
                                                _ = fi55
                                                _ = v58
                                            }
                                            if false {
                                                if (650 != 0) {
                                                    _ = len(v2)
                                                    v59 := v1[-9223372036854775808]
                                                    var v60 uint8 = 2
                                                    _ = v59
                                                    _ = v60
                                                } else {
                                                    switch "" {
                                                    case "a":
                                                        switch fi32 {
                                                        case 32767:
                                                            v61 := GenS4{F0: 266, F1: v31, F2: 18446744073709551615}
                                                            _ = v61
                                                        case -9223372036854775808:
                                                            fmt.Println(v4, v43)
                                                        }
                                                    case "foo bar":
                                                        v3 = append(v3, "hello")
                                                    }
                                                }
                                                v5 -= (v43 * v43)
                                            } else {
                                                v62 := GenS5{F0: 909, F1: (("foo bar" + ("foo bar" + "")) + "\n\t")}
                                                v4 = (v31 - (v31 & -476))
                                                _ = v62
                                            }
                                            for fi63 := 0; fi63 < 8; fi63++ {
                                                switch (fi32 << ((18446744073709551615) & 63)) {
                                                case 9223372036854775807:
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        v64 := GenS2{F0: "", F1: 731, F2: 0}
                                                        _ = v64
                                                    }
                                                case -277:
                                                    v31 = 0
                                                case 0:
                                                    if ((^(fi32)) >= int(v33)) {
                                                        delete(v1, int(v33))
                                                        v4 += (v31 % ((0) | 1))
                                                        fmt.Println(v31)
                                                    } else {
                                                        if false {
                                                            if (genFn0(1, 0) == 42) {
                                                                switch (fi63 % (((^(fi63))) | 1)) {
                                                                case -505:
                                                                    v43 = v43
                                                                case -900:
                                                                    if false {
                                                                        if true {
                                                                            v65 := v1[-2147483648]
                                                                            switch "a" {
                                                                            case "":
                                                                                if len(v3) >= 2 {
                                                                                    v3 = v3[1:len(v3)]
                                                                                }
                                                                            case "a":
                                                                                var v66 uint64 = 983
                                                                                _ = v66
                                                                            case "hello":
                                                                                v1[(fi63 * fi32)] = (v65 + v65)
                                                                            default:
                                                                                var v67 int8 = 127
                                                                                switch fi32 {
                                                                                case -1:
                                                                                    delete(v1, fi32)
                                                                                case 0:
                                                                                    v43 -= ((0.0 - v5) / float32(1e10))
                                                                                case 292:
                                                                                    v5 = (1e10 + v43)
                                                                                    var v68 float64 = -2.25
                                                                                    _ = v68
                                                                                default:
                                                                                    var v69 string = ""
                                                                                    _ = v69
                                                                                }
                                                                                _ = v67
                                                                            }
                                                                            switch "\n\t" {
                                                                            case "a":
                                                                                v70 := GenS6{F0: ("\n\t" + "\n\t")}
                                                                                var v71 int = -1
                                                                                _ = v70
                                                                                _ = v71
                                                                            case "\n\t":
                                                                                switch fi32 {
                                                                                case 127:
                                                                                    fmt.Println(fi63, v65, fi63)
                                                                                case 9223372036854775807:
                                                                                    if (!(!(42 != 2))) {
                                                                                        v33 = 682
                                                                                        if (!(0 > 18446744073709551615)) {
                                                                                            var v72 uint32 = v65
                                                                                            _ = v72
                                                                                        }
                                                                                    } else {
                                                                                        switch fi63 {
                                                                                        case 32768:
                                                                                            fmt.Println(v43, v5, v33)
                                                                                        case 0:
                                                                                            if (!(v5 > v43)) {
                                                                                                var v73 uint32 = genFn0(65535, 4294967295)
                                                                                                _ = v73
                                                                                            } else {
                                                                                                var v74 bool = (uint16(v65) == (v33 >> ((194) & 63)))
                                                                                                delete(v1, -1)
                                                                                                v31 = v31
                                                                                                v75 := GenS7{F0: int16(v33), F1: 127}
                                                                                                fmt.Println(v5, v5, v5)
                                                                                                _ = v74
                                                                                                _ = v75
                                                                                            }
                                                                                        case 4294967295:
                                                                                            if false {
                                                                                                if (!true) {
                                                                                                    if true {
                                                                                                        var v76 string = ("a" + ("a" + ("" + ("" + "foo bar"))))
                                                                                                        _ = len(v1)
                                                                                                        if false {
                                                                                                            if ((v65 & (v65 & 4294967295)) >= genFn0(65535, 234)) {
                                                                                                                fmt.Println(v65, v5)
                                                                                                                v77 := GenS8{F0: 103, F1: 4294967295}
                                                                                                                v65 += (^(v65))
                                                                                                                _ = v77
                                                                                                            }
                                                                                                            if len(v2) > 0 {
                                                                                                                v2[0] = -9223372036854775808
                                                                                                            }
                                                                                                            _ = len(v2)
                                                                                                        }
                                                                                                        _ = v76
                                                                                                    }
                                                                                                }
                                                                                                var v78 uint8 = 127
                                                                                                var v79 uint64 = 4294967295
                                                                                                fmt.Println(v79)
                                                                                                v4 -= (v4 / (((v31 - (v31 - (v31 / ((-322) | 1))))) | 1))
                                                                                                _ = v78
                                                                                                _ = v79
                                                                                            } else {
                                                                                                v80 := GenS6{F0: ""}
                                                                                                if (0 != -32768) {
                                                                                                    v1[fi32] = v65
                                                                                                }
                                                                                                _ = v80
                                                                                            }
                                                                                        default:
                                                                                            v81 := GenS9{F0: -9223372036854775808}
                                                                                            _ = v81
                                                                                        }
                                                                                        var v82 uint = 255
                                                                                        _ = v82
                                                                                    }
                                                                                }
                                                                            }
                                                                            _ = v65
                                                                        } else {
                                                                            var v83 bool = false
                                                                            if (91 <= 2) {
                                                                                delete(v1, 42)
                                                                                var v84 uint = 74
                                                                                _ = v84
                                                                            }
                                                                            _ = v83
                                                                        }
                                                                        v43 = 1e10
                                                                        fmt.Println(fi63, v43)
                                                                    } else {
                                                                        v4 = -2147483648
                                                                        v33 -= 2
                                                                    }
                                                                case -9223372036854775808:
                                                                    v85 := GenS10{F0: -9223372036854775808, F1: 0, F2: (v31 ^ v4)}
                                                                    v43 = v5
                                                                    v5 = 1e10
                                                                    var v86 int64 = -1
                                                                    var v87 bool = (!("" == "hello"))
                                                                    var v88 int32 = -225
                                                                    _ = v85
                                                                    _ = v86
                                                                    _ = v87
                                                                    _ = v88
                                                                }
                                                            }
                                                            if false {
                                                                v31 = v4
                                                                if len(v2) > 0 {
                                                                    v2[0] = fi32
                                                                }
                                                            }
                                                            switch fi63 {
                                                            case -1:
                                                                fmt.Println(v5, fi63, fi63)
                                                            }
                                                        }
                                                        var v89 uint64 = 42
                                                        _ = v89
                                                    }
                                                }
                                                _ = fi63
                                            }
                                        }
                                        switch fi32 {
                                        case -161:
                                            var v90 bool = (((float64(v5) * 3.14159) != (1.5 + (((-2.25 / 0.1) / (3.14159 / (1.5 * 0.1))) / float64(v5)))) || false)
                                            _ = v90
                                        case 42:
                                            v91 := GenS2{F0: ("\n\t" + "a"), F1: 42, F2: 1}
                                            _ = v91
                                        default:
                                            switch "" {
                                            case "hello":
                                                var v92 uint16 = v33
                                                _ = v92
                                            case "\n\t":
                                                for fi93 := 0; fi93 < 4; fi93++ {
                                                    var v94 int = fi32
                                                    fmt.Println(v33, fi32, fi32)
                                                    delete(v1, 42)
                                                    v33 -= v33
                                                    var v95 bool = (false && true)
                                                    _ = fi93
                                                    _ = v94
                                                    _ = v95
                                                }
                                            case "a":
                                                v96 := v1[fi32]
                                                _ = v96
                                            default:
                                                fmt.Println(v5, v43)
                                            }
                                        }
                                    }
                                    switch "" {
                                    case "":
                                        fmt.Println(fi32, v31)
                                    }
                                    var v97 float32 = 1e10
                                    _ = v97
                                } else {
                                    fmt.Println(v4)
                                    switch "\n\t" {
                                    case "\n\t":
                                        v31 = v31
                                    case "foo bar":
                                        v33 = 505
                                    }
                                    for fi98 := 0; fi98 < 11; fi98++ {
                                        v99 := GenS3{F0: (("" + "hello") + "\n\t"), F1: 0, F2: 127}
                                        fmt.Println(v4, v31)
                                        _ = fi98
                                        _ = v99
                                    }
                                }
                            case 0:
                                var v100 int32 = (v31 | (-(v31)))
                                _ = v100
                            default:
                                switch "hello" {
                                case "hello":
                                    for fi101 := 0; fi101 < 20; fi101++ {
                                        fmt.Println(fi101, v4, v5)
                                        _ = len(v3)
                                        if (fi32 < (fi101 / (((fi32 / ((fi101) | 1))) | 1))) {
                                            var v102 bool = (v43 >= (float32(float64((float32(1e10) - float32(3.14159)))) * -2.25))
                                            v103 := GenS1{F0: v33, F1: 0, F2: (v31 + 0)}
                                            v33 = (v33 & v33)
                                            _ = v102
                                            _ = v103
                                        }
                                        _ = fi101
                                    }
                                case "":
                                    v2 = append(v2, 4294967295)
                                }
                            }
                            var v104 int8 = 1
                            _ = v104
                            _ = v43
                        }
                    }
                }
                _ = fi32
                _ = v33
            }
            _ = v31
        }
    }
    switch "\n\t" {
    case "hello":
        fmt.Println(v4, v5)
    case "":
        switch "a" {
        case "foo bar":
            var v105 uint8 = 42
            _ = v105
        case "":
            switch "foo bar" {
            case "a":
                var v106 uint16 = 2
                _ = v106
            }
        case "hello":
            var v107 uint64 = 120
            _ = v107
        }
    default:
        v5 = 3.14159
    }
    if (3.14159 <= 1e10) {
        var v108 float32 = 1.5
        var v109 int = 0
        fmt.Println(v4, v5, v4)
        _ = v108
        _ = v109
    }
    if true {
        v110 := GenS6{F0: "foo bar"}
        _ = v110
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        if (v4 <= 193) {
            if false {
                switch 1 {
                case 42:
                    v4 -= (v4 & (^(v4)))
                case 32767:
                    if (1 < 4294967295) {
                        var v111 int16 = 1
                        _ = v111
                    } else {
                        if (!(((!(-858 > 1)) || false) || true)) {
                            if true {
                                for fi112 := 0; fi112 < 4; fi112++ {
                                    for fi113 := 0; fi113 < 5; fi113++ {
                                        if true {
                                            var v114 uint16 = 1
                                            var v115 string = "a"
                                            switch (^(fi112)) {
                                            case 32768:
                                                switch 128 {
                                                case 32768:
                                                    v116 := GenS8{F0: 0, F1: 4294967295}
                                                    _ = v116
                                                case -9223372036854775808:
                                                    switch "foo bar" {
                                                    case "a":
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        var v117 float64 = 3.14159
                                                        _ = v117
                                                    }
                                                case 127:
                                                    switch (("foo bar" + v115) + "\n\t") {
                                                    case "a":
                                                        var v118 float64 = (3.14159 / ((0.0 + float64(1e10)) - 3.14159))
                                                        _ = v118
                                                    case "foo bar":
                                                        var v119 uint8 = 42
                                                        _ = v119
                                                    case "hello":
                                                        fmt.Println(fi113, fi112, fi113)
                                                    default:
                                                        v1[(fi112 / ((fi113) | 1))] = 2147483648
                                                    }
                                                default:
                                                    v4 = (v4 % ((v4) | 1))
                                                }
                                            case -9223372036854775808:
                                                if (2 > 42) {
                                                    var v120 int8 = -128
                                                    var v121 bool = (false && true)
                                                    fmt.Println(fi112, v115, v115)
                                                    if (true && v121) {
                                                        var v122 string = (v115 + (v115 + ((v115 + v115) + v115)))
                                                        _ = v122
                                                    }
                                                    _ = v120
                                                    _ = v121
                                                } else {
                                                    var v123 float64 = -2.25
                                                    var v124 uint = 1
                                                    _ = v123
                                                    _ = v124
                                                }
                                            case 2147483647:
                                                v125 := GenS8{F0: 151, F1: 18446744073709551615}
                                                switch fi113 {
                                                case 1:
                                                    v115 = (v115 + "hello")
                                                case 32768:
                                                    v5 += (3.14159 - 1e10)
                                                }
                                                _ = v125
                                            default:
                                                v126 := GenS11{F0: 784, F1: 1}
                                                _ = v126
                                            }
                                            _ = v114
                                            _ = v115
                                        } else {
                                            var v127 int = int(v4)
                                            _ = v127
                                        }
                                        v128 := GenS11{F0: 383, F1: 42}
                                        if false {
                                            var v129 int16 = 1
                                            fmt.Println(fi112)
                                            _ = v129
                                        } else {
                                            delete(v1, (fi112 & (^(fi112))))
                                            _ = len(v1)
                                        }
                                        _ = fi113
                                        _ = v128
                                    }
                                    var v130 bool = true
                                    _ = fi112
                                    _ = v130
                                }
                            } else {
                                var v131 float32 = v5
                                fmt.Println(v131, v131, v131)
                                _ = v131
                            }
                            v132 := GenS10{F0: 1, F1: -870, F2: 800}
                            _ = v132
                        }
                        _ = len(v2)
                    }
                default:
                    for fi133 := 0; fi133 < 15; fi133++ {
                        v134 := GenS5{F0: 2, F1: "\n\t"}
                        if false {
                            v135 := GenS12{F0: 32767, F1: 0}
                            var v136 float64 = 1e10
                            _ = v135
                            _ = v136
                        } else {
                            if (false && (fi133 >= 42)) {
                                for fi137 := 0; fi137 < 16; fi137++ {
                                    var v138 int64 = -9223372036854775808
                                    if ((v138 * v138) == (v138 ^ -9223372036854775808)) {
                                        v139 := GenS8{F0: 128, F1: 896}
                                        v2 = append(v2, fi137)
                                        var v140 uint = 892
                                        _ = v139
                                        _ = v140
                                    } else {
                                        var v141 uint16 = 972
                                        var v142 uint32 = 4294967295
                                        if len(v3) > 0 {
                                            v3[0] = "hello"
                                        }
                                        _ = v141
                                        _ = v142
                                    }
                                    v138 += v138
                                    _ = fi137
                                    _ = v138
                                }
                                switch "hello" {
                                case "hello":
                                    switch "hello" {
                                    case "a":
                                        var v143 string = "a"
                                        _ = v143
                                    case "foo bar":
                                        if false {
                                            fmt.Println(fi133, v4)
                                        }
                                    case "":
                                        for fi144 := 0; fi144 < 2; fi144++ {
                                            var v145 int32 = v4
                                            _ = fi144
                                            _ = v145
                                        }
                                    default:
                                        var v146 int = fi133
                                        _ = v146
                                    }
                                case "a":
                                    fmt.Println(v5)
                                case "\n\t":
                                    v147 := GenS13{F0: 1000}
                                    _ = v147
                                }
                            }
                            v5 -= (((0.1 + (float32(0.0) + (v5 + v5))) - -2.25) * v5)
                        }
                        _ = fi133
                        _ = v134
                    }
                }
                if false {
                    if ("foo bar" == "hello") {
                        var v148 uint16 = 255
                        _ = v148
                    } else {
                        _ = len(v2)
                        var v149 int16 = 32767
                        _ = v149
                    }
                    v3 = append(v3, "hello")
                    v4 = 42
                }
            } else {
                for fi150 := 0; fi150 < 20; fi150++ {
                    var v151 int = fi150
                    if len(v3) > 0 {
                        v3[0] = ("hello" + (("" + "foo bar") + "foo bar"))
                    }
                    _ = fi150
                    _ = v151
                }
                v4 += (v4 % ((v4) | 1))
            }
        }
    }
    if true {
        v152 := v1[42]
        for fi153 := 0; fi153 < 16; fi153++ {
            v4 = -2147483648
            if (uint8(v152) != 28) {
                v152 -= v152
                for fi154 := 0; fi154 < 5; fi154++ {
                    var v155 float64 = (0.1 * 0.1)
                    _ = len(v1)
                    _ = fi154
                    _ = v155
                }
            }
            _ = fi153
        }
        v156 := v1[2147483647]
        _ = v152
        _ = v156
    } else {
        for fi157 := 0; fi157 < 3; fi157++ {
            if true {
                fmt.Println(fi157, v4, v4)
            } else {
                v4 = v4
                for fi158 := 0; fi158 < 11; fi158++ {
                    if len(v2) > 0 {
                        v2[0] = (fi158 | (fi158 * (fi157 | (fi158 - (fi157 % ((fi157) | 1))))))
                    }
                    v5 += v5
                    _ = fi158
                }
            }
            if true {
                var v159 uint = 334
                for fi160 := 0; fi160 < 17; fi160++ {
                    var v161 int16 = 437
                    v4 = 42
                    if true {
                        v159 += v159
                        v1[(fi157 % ((fi157) | 1))] = genFn0(2147483648, 18446744073709551615)
                        v161 = (v161 * 0)
                        switch -312 {
                        case -581:
                            v159 += v159
                        default:
                            if ((65535 == 42) || (-867 > fi157)) {
                                switch "\n\t" {
                                case "\n\t":
                                    if ((30 <= 0) && false) {
                                        v159 -= 645
                                        switch ("\n\t" + "") {
                                        case "\n\t":
                                            fmt.Println(v5, v159)
                                        case "a":
                                            fmt.Println(fi160)
                                        case "hello":
                                            v162 := GenS2{F0: (("a" + "foo bar") + ("foo bar" + ("\n\t" + ""))), F1: -9223372036854775808, F2: 1}
                                            _ = v162
                                        }
                                        v4 -= (v4 ^ 765)
                                        switch "\n\t" {
                                        case "a":
                                            fmt.Println(fi160, v159, v161)
                                        }
                                    } else {
                                        var v163 int64 = -1
                                        _ = v163
                                    }
                                default:
                                    var v164 uint16 = 255
                                    _ = v164
                                }
                            }
                        }
                        v4 = 1
                    } else {
                        _ = len(v2)
                    }
                    _ = fi160
                    _ = v161
                }
                _ = v159
            } else {
                var v165 uint = 255
                _ = v165
            }
            v4 = (v4 - 859)
            _ = fi157
        }
        v166 := GenS11{F0: 4294967295, F1: 9223372036854775807}
        _ = v166
    }
    v2 = append(v2, 605)
    fmt.Println(v5, v5)
    if len(v2) > 0 {
        v2[0] = -593
    }
    v4 = (-(v4))
    fmt.Println(v5, v5)
    if ((116 == genFn1(65535)) || true) {
        var v167 int16 = 42
        var v168 int = -1
        _ = v167
        _ = v168
    } else {
        v4 = (-(v4))
        switch "hello" {
        case "hello":
            _ = len(v3)
        case "foo bar":
            v1[32768] = 2147483647
        }
        v169 := GenS7{F0: 364, F1: genFn1(4294967295)}
        _ = v169
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
