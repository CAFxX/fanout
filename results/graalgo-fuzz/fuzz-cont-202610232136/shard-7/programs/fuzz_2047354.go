package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 uint8; F1 int32 }
type GenS5 struct { F0 uint8; F1 int8; F2 uint64 }
type GenS6 struct { F0 uint16; F1 uint32 }
type GenS7 struct { F0 string; F1 uint8; F2 string }
type GenS8 struct { F0 int32; F1 uint }
type GenS9 struct { F0 uint64; F1 string; F2 string }
type GenS10 struct { F0 int64; F1 int64; F2 int8 }
type GenS11 struct { F0 bool; F1 uint64 }
type GenS12 struct { F0 uint16; F1 int; F2 string }
type GenS13 struct { F0 int64; F1 uint; F2 int }
type GenS14 struct { F0 uint32; F1 uint; F2 bool }
type GenS15 struct { F0 uint; F1 uint16 }
type GenS16 struct { F0 string }
type GenS17 struct { F0 int64; F1 int32; F2 uint16 }
type GenS18 struct { F0 uint32; F1 bool; F2 int16 }
type GenS19 struct { F0 int16; F1 uint; F2 uint }
type GenS20 struct { F0 int8; F1 bool; F2 int64 }
type GenS21 struct { F0 int8; F1 int8 }
type GenS22 struct { F0 int16 }
func genFn0(p0 string) uint16 {
    v7 := []int{564}
    p0 = p0
    for fi8 := 0; fi8 < 19; fi8++ {
        fmt.Println(fi8)
        var v9 float32 = (3.14159 / 0.1)
        if (false && false) {
            _ = len(v7)
            if (706 < 345) {
                var v10 int = fi8
                v10 += (v10 << ((471) & 63))
                p0 = p0
                _ = v10
            } else {
                p0 = (p0 + p0)
                for fi11 := 0; fi11 < 9; fi11++ {
                    p0 = p0
                    p0 = p0
                    var v12 uint8 = 0
                    var v13 int64 = -1
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                for fi14 := 0; fi14 < 16; fi14++ {
                    var v15 uint8 = 28
                    if (((42 >= 2) || true) || (2147483648 < 902)) {
                        v15 += (v15 % ((1) | 1))
                    } else {
                        v16 := GenS1{F0: 545, F1: 0}
                        if false {
                            v17 := GenS2{F0: true, F1: -73}
                            v15 = (v15 + v15)
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                            _ = v17
                        } else {
                            var v18 string = (p0 + "")
                            v15 = (v15 | v15)
                            _ = v18
                        }
                        var v19 float64 = (0.0 / 1.5)
                        _ = v16
                        _ = v19
                    }
                    if true {
                        _ = len(v7)
                        switch (fi14 * (^(fi8))) {
                        case -2147483648:
                            var v20 uint32 = 63
                            _ = v20
                        case -413:
                            fmt.Println(v9, fi14, p0)
                        case 4294967295:
                            switch fi8 {
                            case 128:
                                v15 = (v15 >> ((617) & 63))
                            case 2147483647:
                                v21 := GenS3{F0: 0, F1: 100}
                                _ = v21
                            case 42:
                                var v22 int64 = 647
                                _ = v22
                            default:
                                v15 -= 228
                            }
                        default:
                            v7 = append(v7, 128)
                        }
                        v23 := GenS4{F0: 128, F1: -1}
                        _ = v23
                    }
                    v9 = float32(0.0)
                    fmt.Println(fi14, v9, v9)
                    _ = fi14
                    _ = v15
                }
            }
        }
        _ = fi8
        _ = v9
    }
    for fi24 := 0; fi24 < 19; fi24++ {
        if (!true) {
            for fi25 := 0; fi25 < 18; fi25++ {
                p0 = "\n\t"
                fmt.Println(fi24, p0, fi24)
                _ = fi25
            }
            p0 = "\n\t"
            p0 = p0
        } else {
            p0 = p0
            fmt.Println(fi24, fi24)
            if (!((42 != 604) || true)) {
                p0 = p0
            } else {
                p0 = (("hello" + "hello") + "")
                var v26 int16 = -1
                var v27 float64 = (0.0 * 3.14159)
                _ = v26
                _ = v27
            }
        }
        if (!(((^(fi24)) != (^(fi24))) && false)) {
            p0 = (p0 + "hello")
        }
        _ = fi24
    }
    switch p0 {
    case "a":
        fmt.Println(p0)
    }
    _ = p0
    _ = v7
    return 65535
}
func genFn1(p0 int64, p1 uint32, p2 int16) uint {
    for fi28 := 0; fi28 < 7; fi28++ {
        var v29 float64 = 0.1
        if false {
            fmt.Println(p2, fi28, p0)
            v30 := GenS5{F0: 0, F1: -128, F2: 342}
            _ = v30
        } else {
            for fi31 := 0; fi31 < 17; fi31++ {
                fmt.Println(p2, p1, fi31)
                _ = fi31
            }
            v32 := GenS2{F0: false, F1: 1}
            p2 -= p2
            _ = v32
        }
        _ = fi28
        _ = v29
    }
    for fi33 := 0; fi33 < 8; fi33++ {
        fmt.Println(p2, p2, p1)
        _ = fi33
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    v1 := make(map[int]string)
    var v2 uint64 = 4294967295
    var v3 uint32 = 4294967295
    var v4 float32 = float32(3.14159)
    var v5 int32 = 36
    var v6 uint64 = v2
    v6 = 896
    switch "" {
    case "foo bar":
        var v34 uint = genFn1(42, v3, 42)
        _ = v34
    }
    if (123 < 255) {
        switch 420 {
        case 427:
            if true {
                switch 9223372036854775807 {
                case 522:
                    delete(v1, 370)
                case 1:
                    fmt.Println(v6)
                }
            }
        case 102:
            var v35 int8 = 85
            _ = v35
        default:
            v3 -= (^(v3))
        }
        if true {
            v36 := GenS2{F0: (42 <= 1), F1: -28}
            v37 := []int{}
            var v38 int8 = -128
            _ = v36
            _ = v37
            _ = v38
        } else {
            v39 := GenS6{F0: 646, F1: 4294967295}
            delete(v1, 9223372036854775807)
            _ = v39
        }
        for fi40 := 0; fi40 < 3; fi40++ {
            var v41 int8 = 1
            switch (fi40 >> ((4294967295) & 63)) {
            case -9223372036854775808:
                v42 := GenS3{F0: 882, F1: v41}
                _ = v42
            case -764:
                v4 = (0.0 * 3.14159)
            default:
                fmt.Println(v5)
            }
            v43 := GenS4{F0: 127, F1: v5}
            _ = fi40
            _ = v41
            _ = v43
        }
    }
    if true {
        if false {
            for fi44 := 0; fi44 < 2; fi44++ {
                switch "hello" {
                case "\n\t":
                    v6 = 667
                }
                _ = fi44
            }
        }
        v6 -= v6
        v45 := []int{1, 42, -775}
        _ = v45
    }
    v46 := GenS7{F0: "a", F1: 42, F2: "\n\t"}
    switch "foo bar" {
    case "a":
        _ = len(v1)
    case "foo bar":
        v5 -= (v5 % (((v5 ^ (v5 ^ (^(v5))))) | 1))
    case "":
        var v47 string = ("" + "foo bar")
        _ = v47
    default:
        for fi48 := 0; fi48 < 15; fi48++ {
            for fi49 := 0; fi49 < 12; fi49++ {
                if false {
                    v3 = v3
                    delete(v1, (fi48 % ((127) | 1)))
                    switch (fi48 % ((128) | 1)) {
                    case 127:
                        v5 += -1
                    case -198:
                        v50 := GenS8{F0: (v5 % ((v5) | 1)), F1: 282}
                        _ = v50
                    default:
                        fmt.Println(v2, v5)
                    }
                }
                if false {
                    var v51 int = fi48
                    switch (fi49 / ((fi48) | 1)) {
                    case 42:
                        v52 := []int{(v51 ^ (fi49 << ((uint(v2)) & 63))), (v51 + (v51 ^ 42))}
                        if (false || (int16(v3) == 245)) {
                            if len(v52) >= 2 {
                                v52 = v52[1:len(v52)]
                            }
                            if len(v52) > 0 {
                                v52[0] = fi48
                            }
                            v53 := v1[32767]
                            _ = v53
                        } else {
                            var v54 uint = 2
                            v55 := GenS9{F0: v6, F1: "", F2: "foo bar"}
                            _ = v54
                            _ = v55
                        }
                        _ = v52
                    }
                    if false {
                        var v56 string = ("foo bar" + ("hello" + "a"))
                        v5 = -2147483648
                        delete(v1, (fi49 | (fi49 << ((genFn1(470, 2147483648, 42)) & 63))))
                        fmt.Println(v2)
                        switch (^(fi48)) {
                        case -166:
                            switch "foo bar" {
                            case "\n\t":
                                var v57 bool = (false && true)
                                _ = v57
                            case "a":
                                v58 := GenS10{F0: 42, F1: 842, F2: 42}
                                _ = v58
                            case "":
                                if false {
                                    switch fi48 {
                                    case 42:
                                        switch v56 {
                                        case "a":
                                            v59 := []int{(-(v51)), fi48, 4294967295, 310}
                                            _ = v59
                                        }
                                    case 9223372036854775807:
                                        fmt.Println(v56)
                                        fmt.Println(v4)
                                        fmt.Println(v4, v51, v56)
                                    }
                                    var v60 uint = 1
                                    _ = v60
                                } else {
                                    v61 := []int{fi48}
                                    var v62 uint = uint(v6)
                                    _ = v61
                                    _ = v62
                                }
                            }
                        default:
                            switch (v56 + (v56 + (v56 + ((v56 + ((("" + v56) + (v56 + v56)) + v56)) + (("hello" + v56) + v56))))) {
                            case "hello":
                                var v63 int16 = 1
                                _ = v63
                            default:
                                v64 := GenS11{F0: (true || false), F1: v2}
                                _ = v64
                            }
                        }
                        _ = v56
                    } else {
                        v65 := v1[fi48]
                        v51 = (fi49 / (((fi49 + 2147483647)) | 1))
                        _ = v65
                    }
                    _ = len(v1)
                    switch "hello" {
                    case "a":
                        v6 = (v2 | (v2 - (v6 | v2)))
                    case "\n\t":
                        v66 := v1[fi49]
                        _ = v66
                    }
                    var v67 uint = 1
                    if true {
                        v68 := GenS12{F0: uint16(v67), F1: (v51 << ((v67) & 63)), F2: ("\n\t" + "\n\t")}
                        v3 += 219
                        _ = len(v1)
                        _ = v68
                    }
                    switch "hello" {
                    case "":
                        fmt.Println(v4, fi48, v5)
                    case "hello":
                        v2 = (^(v6))
                    default:
                        v69 := []int{128, v51, (v51 + (fi49 / ((-9223372036854775808) | 1)))}
                        _ = v69
                    }
                    _ = v51
                    _ = v67
                }
                fmt.Println(v4, v3)
                var v70 int64 = 0
                delete(v1, int(v3))
                v71 := GenS8{F0: v5, F1: 907}
                _ = fi49
                _ = v70
                _ = v71
            }
            switch int(v5) {
            case 282:
                fmt.Println(v5, v2, v6)
            case 4294967295:
                var v72 string = "hello"
                _ = v72
            default:
                for fi73 := 0; fi73 < 7; fi73++ {
                    v2 += (v6 & v6)
                    switch "foo bar" {
                    case "\n\t":
                        switch (fi48 | fi48) {
                        case 32767:
                            v74 := []int{-2147483648}
                            _ = v74
                        default:
                            fmt.Println(v2, v4, fi48)
                        }
                    }
                    switch (fi48 * fi73) {
                    case 128:
                        var v75 float32 = v4
                        if false {
                            var v76 float64 = 0.1
                            _ = v76
                        }
                        switch (fi48 >> ((255) & 63)) {
                        case -9223372036854775808:
                            v77 := []int{42}
                            _ = v77
                        case 0:
                            switch "" {
                            case "a":
                                v6 = v2
                            case "foo bar":
                                if false {
                                    v78 := GenS5{F0: 0, F1: 116, F2: v2}
                                    v79 := []int{fi73, 127}
                                    _ = v78
                                    _ = v79
                                } else {
                                    v2 = v6
                                    v4 = 1e10
                                    var v80 uint64 = v6
                                    v81 := []int{42}
                                    _ = v80
                                    _ = v81
                                }
                            case "":
                                switch 127 {
                                case 42:
                                    v82 := []int{fi48}
                                    _ = v82
                                case -71:
                                    if (1 == 9223372036854775807) {
                                        if ((fi73 / ((-1) | 1)) == -728) {
                                            v83 := v1[(fi48 * fi48)]
                                            switch (fi73 / (((fi48 + (fi48 & (fi48 >> ((695) & 63))))) | 1)) {
                                            case 128:
                                                var v84 int = fi73
                                                _ = v84
                                            case 4294967295:
                                                var v85 int8 = -1
                                                _ = v85
                                            }
                                            _ = v83
                                        } else {
                                            v86 := []int{fi48}
                                            v87 := GenS13{F0: int64(v6), F1: 1, F2: (fi48 >> ((genFn1(42, v3, 102)) & 63))}
                                            _ = v86
                                            _ = v87
                                        }
                                        if false {
                                            var v88 int8 = 127
                                            _ = v88
                                        } else {
                                            if true {
                                                fmt.Println(fi48)
                                                if true {
                                                    fmt.Println(v5, fi48)
                                                    switch ("" + "a") {
                                                    case "\n\t":
                                                        if (0 != -1) {
                                                            fmt.Println(v2, v5)
                                                            fmt.Println(v2, fi73)
                                                        } else {
                                                            v1[835] = "\n\t"
                                                        }
                                                    default:
                                                        _ = len(v1)
                                                    }
                                                    v89 := GenS1{F0: -9223372036854775808, F1: 0}
                                                    _ = v89
                                                }
                                            }
                                        }
                                        var v90 string = "foo bar"
                                        _ = v90
                                    }
                                default:
                                    if true {
                                        if (!false) {
                                            var v91 float32 = v75
                                            _ = v91
                                        } else {
                                            if false {
                                                var v92 uint64 = (^(v2))
                                                _ = v92
                                            }
                                            switch "a" {
                                            case "\n\t":
                                                fmt.Println(v75, v6, v4)
                                            case "a":
                                                var v93 uint = 4294967295
                                                _ = v93
                                            case "foo bar":
                                                v94 := []int{(-(fi48))}
                                                _ = v94
                                            default:
                                                v95 := GenS14{F0: (v3 % (((v3 << ((2) & 63))) | 1)), F1: 207, F2: true}
                                                _ = v95
                                            }
                                        }
                                        fmt.Println(v75, v75)
                                        _ = len(v1)
                                        switch "a" {
                                        case "a":
                                            switch (-(fi48)) {
                                            case -2147483648:
                                                var v96 uint = 4294967295
                                                _ = v96
                                            }
                                        case "\n\t":
                                            if (v2 >= v2) {
                                                var v97 float32 = v4
                                                _ = v97
                                            } else {
                                                v2 = (v2 | (v6 | v6))
                                                if (-756 != 99) {
                                                    if (false && false) {
                                                        var v98 float32 = v75
                                                        _ = v98
                                                    } else {
                                                        v75 += -2.25
                                                        if (true && (-882 >= 1)) {
                                                            v3 = (v3 + (v3 >> ((0) & 63)))
                                                            if (((!false) && false) || false) {
                                                                v99 := v1[int(v6)]
                                                                if (true || false) {
                                                                    v6 += (^(v2))
                                                                    fmt.Println(v5, v5, v75)
                                                                    v99 = v99
                                                                }
                                                                var v100 float32 = float32(0.1)
                                                                _ = v100
                                                                _ = v99
                                                            }
                                                        }
                                                        if (-9223372036854775808 != -9223372036854775808) {
                                                            var v101 bool = false
                                                            var v102 uint16 = genFn0(("hello" + (("hello" + ("hello" + ("\n\t" + "hello"))) + ("\n\t" + (("" + (("foo bar" + "a") + (("foo bar" + "a") + "\n\t"))) + ("" + (("" + "\n\t") + "")))))))
                                                            var v103 float64 = (1e10 - ((0.0 - -2.25) * 1.5))
                                                            if (float64((v75 / v75)) <= float64(((v4 / 1.0) + v75))) {
                                                                v104 := GenS10{F0: 1, F1: 0, F2: 127}
                                                                v105 := []int{}
                                                                _ = v104
                                                                _ = v105
                                                            } else {
                                                                switch ((("\n\t" + "a") + ("hello" + "hello")) + "") {
                                                                case "foo bar":
                                                                    v75 = v75
                                                                case "\n\t":
                                                                    var v106 float32 = v75
                                                                    _ = v106
                                                                case "":
                                                                    var v107 uint8 = 255
                                                                    _ = v107
                                                                default:
                                                                    var v108 uint64 = (v6 >> ((4294967295) & 63))
                                                                    v109 := v1[fi73]
                                                                    _ = v108
                                                                    _ = v109
                                                                }
                                                                v102 = v102
                                                            }
                                                            if true {
                                                                v5 = 2147483647
                                                                var v110 float64 = ((float64(((float32((-2.25 + 1.5)) * (v75 - float32(v103))) - v75)) * (v103 / 1.0)) - v103)
                                                                fmt.Println(v110, fi73, v75)
                                                                _ = v110
                                                            } else {
                                                                switch "" {
                                                                case "a":
                                                                    v101 = false
                                                                case "\n\t":
                                                                    if v101 {
                                                                        var v111 uint = 65535
                                                                        _ = v111
                                                                    } else {
                                                                        var v112 int = fi48
                                                                        v4 += v75
                                                                        fmt.Println(v112, v2, v3)
                                                                        _ = v112
                                                                    }
                                                                case "foo bar":
                                                                    switch "" {
                                                                    case "":
                                                                        fmt.Println(v3, v3)
                                                                    case "hello":
                                                                        switch fi73 {
                                                                        case -9223372036854775808:
                                                                            var v113 int8 = 75
                                                                            _ = v113
                                                                        default:
                                                                            if (v3 == 42) {
                                                                                fmt.Println(v6, v103, v103)
                                                                                v75 = (v4 - v4)
                                                                            } else {
                                                                                if v101 {
                                                                                    v114 := GenS15{F0: 669, F1: (v102 << ((903) & 63))}
                                                                                    if (0.1 != v75) {
                                                                                        v5 = v5
                                                                                        if ((v6 < (v2 >> ((502) & 63))) && (!false)) {
                                                                                            if ((!false) || (-72 == 100)) {
                                                                                                v103 = float64(-2.25)
                                                                                                v115 := []int{(fi73 % ((127) | 1)), (fi48 << ((922) & 63))}
                                                                                                _ = v115
                                                                                            } else {
                                                                                                if true {
                                                                                                    fmt.Println(v103, fi73)
                                                                                                    var v116 int = 32767
                                                                                                    v3 = (v3 << ((209) & 63))
                                                                                                    _ = v116
                                                                                                } else {
                                                                                                    var v117 string = ""
                                                                                                    v101 = v101
                                                                                                    _ = v117
                                                                                                }
                                                                                                v101 = ((v6 * v2) != v2)
                                                                                            }
                                                                                            v102 = 1
                                                                                        } else {
                                                                                            if (144 < genFn0("hello")) {
                                                                                                var v118 uint64 = (v6 / (((v6 ^ v2)) | 1))
                                                                                                var v119 string = "a"
                                                                                                _ = v118
                                                                                                _ = v119
                                                                                            } else {
                                                                                                if (v103 == ((v103 - v103) + ((v103 - (float64(v4) / (v103 / v103))) + -2.25))) {
                                                                                                    v120 := GenS8{F0: 2147483647, F1: 18446744073709551615}
                                                                                                    v5 += (v5 + v5)
                                                                                                    _ = v120
                                                                                                }
                                                                                                switch "\n\t" {
                                                                                                case "foo bar":
                                                                                                    switch fi48 {
                                                                                                    case 127:
                                                                                                        v102 -= v102
                                                                                                    default:
                                                                                                        switch (fi73 + (fi48 ^ fi48)) {
                                                                                                        case 306:
                                                                                                            v121 := GenS15{F0: 413, F1: v102}
                                                                                                            _ = v121
                                                                                                        default:
                                                                                                            v6 = (v2 & 590)
                                                                                                        }
                                                                                                    }
                                                                                                case "\n\t":
                                                                                                    switch fi73 {
                                                                                                    case 2147483647:
                                                                                                        var v122 int64 = 9223372036854775807
                                                                                                        _ = v122
                                                                                                    case 42:
                                                                                                        fmt.Println(v4, v3)
                                                                                                    default:
                                                                                                        v4 = float32((float64(v75) - v103))
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            if false {
                                                                                                delete(v1, fi48)
                                                                                                if v101 {
                                                                                                    v123 := GenS16{F0: "foo bar"}
                                                                                                    _ = v123
                                                                                                } else {
                                                                                                    v124 := v1[-9223372036854775808]
                                                                                                    var v125 uint8 = 1
                                                                                                    if true {
                                                                                                        v75 -= 3.14159
                                                                                                    }
                                                                                                    _ = v124
                                                                                                    _ = v125
                                                                                                }
                                                                                                fmt.Println(v4, v3, v3)
                                                                                                if (-1 < -1) {
                                                                                                    v126 := []int{(fi73 | (fi73 & fi73)), (fi73 & fi48), fi48, (fi48 & (fi48 << ((0) & 63)))}
                                                                                                    v101 = v101
                                                                                                    _ = v126
                                                                                                }
                                                                                            } else {
                                                                                                if false {
                                                                                                    v127 := GenS17{F0: -9223372036854775808, F1: v5, F2: v102}
                                                                                                    v128 := GenS18{F0: v3, F1: (v101 || ((375 <= v3) && v101)), F2: int16(v3)}
                                                                                                    v103 = (3.14159 + v103)
                                                                                                    _ = v127
                                                                                                    _ = v128
                                                                                                } else {
                                                                                                    if ((-(v5)) < 2147483647) {
                                                                                                        v129 := GenS19{F0: -886, F1: 355, F2: uint(v2)}
                                                                                                        fmt.Println(v2, fi73)
                                                                                                        _ = len(v1)
                                                                                                        v130 := GenS19{F0: 20, F1: 18446744073709551615, F2: 65535}
                                                                                                        _ = v129
                                                                                                        _ = v130
                                                                                                    }
                                                                                                    v4 -= (-2.25 / 0.1)
                                                                                                }
                                                                                                var v131 uint8 = 174
                                                                                                _ = v131
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    v3 += v3
                                                                                    _ = v114
                                                                                }
                                                                            }
                                                                        }
                                                                    default:
                                                                        switch -9223372036854775808 {
                                                                        case 32768:
                                                                            fmt.Println(v4, v101)
                                                                        }
                                                                    }
                                                                default:
                                                                    var v132 string = "hello"
                                                                    _ = v132
                                                                }
                                                            }
                                                            _ = v101
                                                            _ = v102
                                                            _ = v103
                                                        } else {
                                                            v133 := GenS18{F0: 495, F1: ((!true) || false), F2: 245}
                                                            fmt.Println(v2, v5)
                                                            _ = v133
                                                        }
                                                    }
                                                }
                                                v2 = (v2 + v6)
                                            }
                                        case "hello":
                                            if true {
                                                switch "foo bar" {
                                                case "foo bar":
                                                    v1[fi48] = "a"
                                                case "hello":
                                                    if false {
                                                        if ((v4 > 1.5) || (float32(((((0.1 / (0.1 / (((3.14159 * -2.25) * (0.1 / -2.25)) * -2.25))) * 1e10) + ((-2.25 * 1e10) * -2.25)) / -2.25)) != 0.0)) {
                                                            var v134 float32 = float32(-2.25)
                                                            var v135 uint = 2
                                                            v136 := GenS20{F0: 0, F1: (genFn1(int64(v135), 65535, 1000) != 18446744073709551615), F2: -137}
                                                            _ = v134
                                                            _ = v135
                                                            _ = v136
                                                        } else {
                                                            if (!(42 > uint(v2))) {
                                                                v3 -= 0
                                                                var v137 uint16 = 42
                                                                _ = v137
                                                            } else {
                                                                var v138 float32 = float32((0.0 / float64(0.1)))
                                                                v4 -= (v4 / 3.14159)
                                                                switch fi48 {
                                                                case -9223372036854775808:
                                                                    v139 := GenS21{F0: 98, F1: 42}
                                                                    _ = v139
                                                                default:
                                                                    v4 += 3.14159
                                                                }
                                                                _ = v138
                                                            }
                                                            var v140 int8 = 1
                                                            if (false || true) {
                                                                delete(v1, (fi48 << ((2) & 63)))
                                                                v4 = (3.14159 + (1.5 - 1e10))
                                                                var v141 uint = 42
                                                                v4 = v4
                                                                var v142 uint = (^(v141))
                                                                _ = v141
                                                                _ = v142
                                                            } else {
                                                                var v143 float32 = v75
                                                                _ = v143
                                                            }
                                                            _ = v140
                                                        }
                                                    }
                                                    v144 := GenS18{F0: 2147483648, F1: (!(!true)), F2: 32767}
                                                    v145 := []int{(fi48 ^ (fi73 ^ (fi73 & -504)))}
                                                    _ = v144
                                                    _ = v145
                                                }
                                                var v146 int16 = -32768
                                                _ = v146
                                            } else {
                                                v147 := GenS2{F0: (-20 != 127), F1: -31}
                                                fmt.Println(v2, v2)
                                                _ = v147
                                            }
                                        default:
                                            fmt.Println(fi48, fi73)
                                        }
                                        fmt.Println(v6, v75, v75)
                                    } else {
                                        var v148 uint8 = 42
                                        _ = v148
                                    }
                                }
                            default:
                                v3 = 2147483647
                            }
                        case 9223372036854775807:
                            v149 := []int{}
                            _ = v149
                        default:
                            v150 := GenS22{F0: -382}
                            _ = v150
                        }
                        v5 += (v5 ^ (v5 - (v5 & v5)))
                        _ = v75
                    case -171:
                        var v151 int32 = (v5 & (-(v5)))
                        _ = v151
                    }
                    if ((255 != 166) && (!false)) {
                        fmt.Println(v3, v2)
                        var v152 bool = (212 != 153)
                        fmt.Println(v4)
                        _ = v152
                    }
                    if true {
                        v153 := GenS21{F0: 1, F1: 0}
                        v154 := GenS14{F0: (v3 << ((genFn1(1, 2147483647, 0)) & 63)), F1: 65535, F2: ("" == ("a" + ("a" + "a")))}
                        v155 := []int{(fi48 ^ (fi73 - fi48)), (fi48 << ((genFn1(-950, (v3 | (v3 + v3)), -32768)) & 63))}
                        switch ("hello" + ("foo bar" + "foo bar")) {
                        case "\n\t":
                            v156 := GenS19{F0: 42, F1: 803, F2: 1}
                            _ = v156
                        default:
                            delete(v1, 128)
                        }
                        _ = v153
                        _ = v154
                        _ = v155
                    }
                    _ = fi73
                }
            }
            _ = fi48
        }
    }
    v157 := []int{128, -1, -639}
    v158 := GenS13{F0: -821, F1: 155, F2: 501}
    var v159 float64 = 1.5
    switch 1 {
    case 127:
        v6 += v2
    case -1:
        if len(v157) > 0 {
            v157[0] = -2147483648
        }
    }
    var v160 uint = 100
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v157:", len(v157), v157)
    fmt.Println("v159:", v159)
    fmt.Println("v160:", v160)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v157
    _ = v158
    _ = v159
    _ = v160
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v6
}
