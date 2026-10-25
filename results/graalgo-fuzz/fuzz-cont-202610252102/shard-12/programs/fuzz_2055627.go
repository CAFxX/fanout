package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint32; F1 string; F2 int64 }
type GenS4 struct { F0 uint; F1 int; F2 uint8 }
type GenS5 struct { F0 int64; F1 uint8 }
type GenS6 struct { F0 int32; F1 string; F2 bool }
type GenS7 struct { F0 int64; F1 int32 }
type GenS8 struct { F0 bool; F1 int; F2 int32 }
type GenS9 struct { F0 uint32; F1 int }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 uint8; F1 uint8; F2 int8 }
type GenS13 struct { F0 uint64; F1 int8 }
type GenS14 struct { F0 int8 }
type GenS15 struct { F0 int8; F1 uint32 }
type GenS16 struct { F0 bool }
type GenS17 struct { F0 int16; F1 int }
type GenS18 struct { F0 uint16; F1 uint; F2 string }
type GenS19 struct { F0 int8; F1 uint8; F2 int32 }
type GenS20 struct { F0 bool; F1 int16 }
type GenS21 struct { F0 uint8; F1 bool }
type GenS22 struct { F0 uint32; F1 uint16; F2 uint8 }
type GenS23 struct { F0 int; F1 bool; F2 int32 }
type GenS24 struct { F0 int16; F1 uint; F2 uint }
type GenS25 struct { F0 int32 }
type GenS26 struct { F0 string; F1 int8; F2 bool }
type GenS27 struct { F0 uint16 }
type GenS28 struct { F0 bool }
type GenS29 struct { F0 int16; F1 int16 }
type GenS30 struct { F0 int16 }
type GenS31 struct { F0 uint32 }
type GenS32 struct { F0 int }
func genFn0() int16 {
    var v6 int16 = 876
    if true {
        v7 := []int{9223372036854775807, -1, -277, 9223372036854775807}
        if (1.5 != ((3.14159 - -2.25) * -2.25)) {
            var v8 int = 759
            fmt.Println(v6, v6)
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            _ = v8
        }
        v6 += v6
        _ = v7
    }
    switch 1 {
    case 1:
        for fi9 := 0; fi9 < 20; fi9++ {
            var v10 bool = (!true)
            v11 := GenS1{F0: ("foo bar" + ("a" + ("\n\t" + ""))), F1: -868, F2: 122}
            _ = fi9
            _ = v10
            _ = v11
        }
    case -565:
        for fi12 := 0; fi12 < 5; fi12++ {
            v6 -= int16(fi12)
            if false {
                if (!false) {
                    fmt.Println(v6)
                }
                var v13 int = (fi12 >> ((4294967295) & 63))
                _ = v13
            }
            switch (("foo bar" + ("\n\t" + "")) + "") {
            case "foo bar":
                for fi14 := 0; fi14 < 1; fi14++ {
                    v15 := GenS2{F0: (fi12 << ((2) & 63))}
                    var v16 float64 = 0.1
                    switch ("\n\t" + (("" + ("a" + "hello")) + "hello")) {
                    case "a":
                        v6 = v6
                    default:
                        v6 -= v6
                    }
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            case "hello":
                if true {
                    switch "foo bar" {
                    case "hello":
                        if (-2.25 <= float32(-2.25)) {
                            var v17 float32 = 1.5
                            var v18 int64 = 9223372036854775807
                            _ = v17
                            _ = v18
                        }
                    case "foo bar":
                        v6 = v6
                    case "a":
                        v19 := []int{128, (fi12 | (fi12 * -115)), (^(fi12))}
                        _ = v19
                    default:
                        v6 -= v6
                    }
                }
            default:
                for fi20 := 0; fi20 < 3; fi20++ {
                    var v21 uint32 = 1
                    if false {
                        v22 := []int{}
                        _ = v22
                    } else {
                        switch (("\n\t" + "foo bar") + "foo bar") {
                        case "a":
                            switch "\n\t" {
                            case "a":
                                var v23 float32 = 1e10
                                _ = v23
                            }
                            v6 = v6
                            var v24 int64 = -9223372036854775808
                            _ = v24
                        }
                        v6 -= (v6 / ((v6) | 1))
                        if ((v6 - (v6 | 1)) < v6) {
                            switch ("hello" + ("foo bar" + ("\n\t" + "a"))) {
                            case "":
                                var v25 bool = (-62 >= 380)
                                var v26 float64 = (((float64(0.0) - 3.14159) - (3.14159 - 1.5)) + 0.1)
                                fmt.Println(v6, fi20)
                                _ = v25
                                _ = v26
                            case "a":
                                switch (fi20 | fi12) {
                                case 42:
                                    var v27 int64 = 42
                                    _ = v27
                                case 2147483647:
                                    if (0.1 <= 0.0) {
                                        v6 -= v6
                                        var v28 int64 = -1
                                        var v29 uint16 = 65535
                                        _ = v28
                                        _ = v29
                                    }
                                default:
                                    v30 := GenS1{F0: ("\n\t" + ("foo bar" + "\n\t")), F1: 1, F2: -1}
                                    _ = v30
                                }
                            case "\n\t":
                                v21 -= (v21 & 2147483647)
                            default:
                                v6 -= int16(fi12)
                            }
                            v31 := GenS3{F0: (v21 + 4294967295), F1: "", F2: -9223372036854775808}
                            _ = v31
                        }
                    }
                    _ = fi20
                    _ = v21
                }
            }
            _ = fi12
        }
    case 32767:
        var v32 int8 = 50
        _ = v32
    default:
        if false {
            if true {
                v33 := GenS4{F0: 65535, F1: 32767, F2: 2}
                _ = v33
            }
        } else {
            v6 -= v6
        }
    }
    _ = v6
    return v6
}

func main() {
    v1 := []string{"foo bar", "foo bar", (("\n\t" + "\n\t") + "\n\t"), "foo bar"}
    var v2 uint8 = 255
    v3 := make(map[int]uint32)
    v4 := make(map[string]bool)
    v5 := []string{}
    v2 += (v2 >> ((0) & 63))
    var v34 int32 = 758
    switch "\n\t" {
    case "foo bar":
        switch "a" {
        case "hello":
            if true {
                switch 42 {
                case 9223372036854775807:
                    var v35 int32 = 2147483647
                    _ = v35
                default:
                    v2 += (^(v2))
                }
                _ = len(v4)
            }
        case "":
            for fi36 := 0; fi36 < 9; fi36++ {
                var v37 int8 = -101
                var v38 int32 = v34
                var v39 uint64 = 856
                _ = fi36
                _ = v37
                _ = v38
                _ = v39
            }
        }
    case "":
        v40 := GenS5{F0: -9223372036854775808, F1: v2}
        _ = v40
    default:
        v34 += (^(v34))
    }
    for fi41 := 0; fi41 < 8; fi41++ {
        for fi42 := 0; fi42 < 17; fi42++ {
            if (!false) {
                v5 = append(v5, "\n\t")
            }
            v2 -= v2
            _ = fi42
        }
        var v43 uint = 188
        _ = fi41
        _ = v43
    }
    var v44 float64 = (-2.25 + ((1.5 - float64(1.5)) * -2.25))
    if true {
        v45 := GenS2{F0: 127}
        fmt.Println(v2, v44, v44)
        var v46 int32 = v34
        _ = v45
        _ = v46
    } else {
        var v47 uint32 = 2
        var v48 uint8 = 127
        _ = v47
        _ = v48
    }
    if len(v5) > 0 {
        v5[0] = ("foo bar" + "a")
    }
    switch 128 {
    case 127:
        v44 += v44
    case 32768:
        var v49 bool = (255 > 181)
        _ = v49
    default:
        fmt.Println(v44, v34)
    }
    for fi50 := 0; fi50 < 18; fi50++ {
        for fi51 := 0; fi51 < 15; fi51++ {
            if true {
                fmt.Println(fi51, v44, v2)
                v52 := GenS1{F0: (("\n\t" + ("a" + "foo bar")) + "foo bar"), F1: (v34 ^ (^(v34))), F2: -128}
                _ = v52
            } else {
                v53 := GenS6{F0: 42, F1: ("hello" + ("hello" + "")), F2: true}
                var v54 string = "\n\t"
                _ = v53
                _ = v54
            }
            var v55 uint64 = 18446744073709551615
            switch ("foo bar" + "hello") {
            case "foo bar":
                v56 := GenS2{F0: 1}
                _ = v56
            case "a":
                v55 += uint64(v2)
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                fmt.Println(fi51)
            case "hello":
                var v57 float32 = 1e10
                if false {
                    var v58 int = fi51
                    v2 = 1
                    _ = v58
                } else {
                    if false {
                        var v59 uint16 = 2
                        _ = v59
                    } else {
                        var v60 int32 = v34
                        if len(v1) > 0 {
                            v1[0] = ("hello" + "hello")
                        }
                        _ = v60
                    }
                    if true {
                        v5 = append(v5, "\n\t")
                        switch 9223372036854775807 {
                        case -10:
                            fmt.Println(v44, v2)
                        case 127:
                            v44 = v44
                        case 69:
                            v61 := GenS7{F0: -9223372036854775808, F1: 922}
                            _ = v61
                        }
                        if ((false && true) || true) {
                            v62 := GenS8{F0: true, F1: fi51, F2: v34}
                            _ = v62
                        } else {
                            v63 := GenS5{F0: -9223372036854775808, F1: v2}
                            v64 := GenS9{F0: 850, F1: 30}
                            if ((-128 > 127) && (989 != 2)) {
                                _ = len(v4)
                                if false {
                                    var v65 int64 = -1
                                    _ = v65
                                } else {
                                    var v66 float64 = v44
                                    _ = v66
                                }
                                switch "a" {
                                case "foo bar":
                                    v57 = v57
                                case "a":
                                    if (false && (276 < v34)) {
                                        v44 += (v44 / ((0.1 - 3.14159) + v44))
                                    } else {
                                        v2 -= (^(v2))
                                    }
                                case "\n\t":
                                    if false {
                                        var v67 uint8 = v2
                                        if false {
                                            var v68 uint = 42
                                            _ = v68
                                        } else {
                                            var v69 uint32 = 2
                                            fmt.Println(v2, v69)
                                            v69 = (v69 * (v69 - 2))
                                            fmt.Println(v2, v57)
                                            var v70 uint64 = v55
                                            var v71 uint16 = 1
                                            v55 -= (v70 / ((uint64(v71)) | 1))
                                            switch (^(fi50)) {
                                            case 42:
                                                fmt.Println(v71, fi50)
                                            case 2147483647:
                                                v71 = 32767
                                            }
                                            _ = v69
                                            _ = v70
                                            _ = v71
                                        }
                                        if len(v5) > 0 {
                                            v5[0] = (("foo bar" + (("foo bar" + ("foo bar" + "a")) + ("\n\t" + ""))) + ("foo bar" + "hello"))
                                        }
                                        _ = v67
                                    } else {
                                        var v72 uint64 = 0
                                        var v73 uint = 0
                                        _ = v72
                                        _ = v73
                                    }
                                }
                            } else {
                                v34 = (v34 - (-(v34)))
                                fmt.Println(fi51, v34)
                                fmt.Println(v55, fi51)
                                var v74 bool = (false && ((-128 > 1) && (2 > 354)))
                                _ = v74
                            }
                            fmt.Println(v55)
                            fmt.Println(fi50)
                            _ = v63
                            _ = v64
                        }
                    } else {
                        v44 -= v44
                        switch (fi51 | (fi50 | (fi50 * (fi51 ^ 0)))) {
                        case 127:
                            v75 := v4[("" + "hello")]
                            _ = v75
                        default:
                            fmt.Println(v55, v57, fi51)
                        }
                    }
                    if false {
                        if (((0 < 857) || (uint(v2) == 371)) || (v34 <= -207)) {
                            switch "foo bar" {
                            case "\n\t":
                                v34 += -1
                            case "a":
                                if true {
                                    var v76 float64 = (v44 / v44)
                                    v76 += float64((((1e10 / 1.0) - (0.0 - v57)) - 0.0))
                                    _ = v76
                                } else {
                                    fmt.Println(v44, v44, v34)
                                    v55 += v55
                                    var v77 float32 = ((v57 * (v57 + 0.0)) * 0.1)
                                    _ = v77
                                }
                            case "hello":
                                v57 = v57
                            }
                            v78 := GenS10{F0: 932}
                            v79 := GenS2{F0: (^(fi50))}
                            v80 := GenS2{F0: (fi50 & fi51)}
                            _ = v78
                            _ = v79
                            _ = v80
                        } else {
                            var v81 int32 = 2147483647
                            switch "a" {
                            case "hello":
                                if true {
                                    switch (fi50 - (fi50 >> ((588) & 63))) {
                                    case -158:
                                        if true {
                                            var v82 int = (fi51 << ((42) & 63))
                                            v34 = v81
                                            _ = v82
                                        }
                                    case 32767:
                                        var v83 int64 = -9223372036854775808
                                        _ = v83
                                    case -9223372036854775808:
                                        switch fi50 {
                                        case 127:
                                            var v84 int16 = genFn0()
                                            _ = v84
                                        case 0:
                                            v85 := GenS2{F0: (fi50 / (((fi50 << ((65535) & 63))) | 1))}
                                            _ = v85
                                        case -418:
                                            if (2147483647 >= -172) {
                                                if true {
                                                    switch "" {
                                                    case "foo bar":
                                                        v3[(fi51 * fi51)] = 2147483647
                                                    case "":
                                                        v57 += v57
                                                    default:
                                                        _ = len(v3)
                                                    }
                                                    v86 := GenS11{F0: 0}
                                                    v44 = v44
                                                    switch 800 {
                                                    case 128:
                                                        v87 := GenS12{F0: 1, F1: v2, F2: -104}
                                                        _ = v87
                                                    case 9223372036854775807:
                                                        fmt.Println(fi51, v57)
                                                    case -835:
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                    }
                                                    _ = v86
                                                } else {
                                                    v88 := GenS13{F0: (v55 + (v55 * uint64(v2))), F1: 127}
                                                    var v89 int16 = 1
                                                    fmt.Println(v81)
                                                    if ((fi51 * fi50) >= -976) {
                                                        if false {
                                                            if false {
                                                                v90 := GenS14{F0: 127}
                                                                if false {
                                                                    v55 = v55
                                                                    _ = len(v3)
                                                                    if (!(v81 >= (v34 * -2147483648))) {
                                                                        switch ("" + "foo bar") {
                                                                        case "":
                                                                            if (162 == (v55 / ((v55) | 1))) {
                                                                                if false {
                                                                                    if true {
                                                                                        v91 := v3[(fi51 >> ((18446744073709551615) & 63))]
                                                                                        var v92 float32 = 1.5
                                                                                        if true {
                                                                                            switch 4294967295 {
                                                                                            case -95:
                                                                                                v93 := GenS9{F0: v91, F1: fi51}
                                                                                                _ = v93
                                                                                            case -899:
                                                                                                v91 = v91
                                                                                            case -2147483648:
                                                                                                var v94 int = (fi50 / ((fi50) | 1))
                                                                                                _ = v94
                                                                                            }
                                                                                            v89 -= -1
                                                                                        } else {
                                                                                            switch (fi50 ^ (fi51 * (fi51 * (fi51 + (-(fi50)))))) {
                                                                                            case -9223372036854775808:
                                                                                                if (3.14159 > v44) {
                                                                                                    v5 = append(v5, "a")
                                                                                                    var v95 int = fi50
                                                                                                    fmt.Println(v91, v91)
                                                                                                    _ = v95
                                                                                                }
                                                                                            case 1:
                                                                                                switch 4294967295 {
                                                                                                case -9223372036854775808:
                                                                                                    v3[128] = v91
                                                                                                case 0:
                                                                                                    v96 := GenS12{F0: v2, F1: (v2 & (^(v2))), F2: 93}
                                                                                                    _ = v96
                                                                                                case -2147483648:
                                                                                                    var v97 bool = true
                                                                                                    _ = v97
                                                                                                }
                                                                                            default:
                                                                                                if (true && false) {
                                                                                                    var v98 bool = (0.0 > 1.5)
                                                                                                    _ = v98
                                                                                                } else {
                                                                                                    if (874 >= 255) {
                                                                                                        if false {
                                                                                                            var v99 string = ("\n\t" + "a")
                                                                                                            _ = v99
                                                                                                        }
                                                                                                        v44 += float64(v57)
                                                                                                        if len(v5) >= 2 {
                                                                                                            v5 = v5[1:len(v5)]
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            v100 := GenS11{F0: 42}
                                                                                            _ = v100
                                                                                        }
                                                                                        if len(v1) > 0 {
                                                                                            v1[0] = "a"
                                                                                        }
                                                                                        v3[(^(fi51))] = 421
                                                                                        _ = v91
                                                                                        _ = v92
                                                                                    }
                                                                                } else {
                                                                                    var v101 int8 = -1
                                                                                    _ = len(v3)
                                                                                    _ = v101
                                                                                }
                                                                                var v102 int16 = v89
                                                                                _ = v102
                                                                            } else {
                                                                                if ((v2 * v2) <= uint8(v55)) {
                                                                                    var v103 uint32 = 65535
                                                                                    fmt.Println(v81, v103, fi51)
                                                                                    fmt.Println(v34, v89, v55)
                                                                                    _ = v103
                                                                                }
                                                                                var v104 bool = true
                                                                                _ = v104
                                                                            }
                                                                        case "\n\t":
                                                                            switch "" {
                                                                            case "hello":
                                                                                v105 := v3[fi50]
                                                                                _ = v105
                                                                            case "":
                                                                                switch 0 {
                                                                                case -1:
                                                                                    var v106 uint = 528
                                                                                    _ = v106
                                                                                case -2147483648:
                                                                                    var v107 int32 = (^(v34))
                                                                                    v89 += v89
                                                                                    v1 = append(v1, "")
                                                                                    _ = v107
                                                                                case 127:
                                                                                    v57 = v57
                                                                                }
                                                                            }
                                                                        default:
                                                                            v108 := GenS8{F0: (fi50 >= 2147483647), F1: fi50, F2: v81}
                                                                            _ = v108
                                                                        }
                                                                        fmt.Println(v2, v2)
                                                                        v89 = -1
                                                                        if (true && true) {
                                                                            v34 -= v81
                                                                            v109 := GenS13{F0: 937, F1: -92}
                                                                            v34 = 1
                                                                            _ = v109
                                                                        }
                                                                    }
                                                                    switch 42 {
                                                                    case 9223372036854775807:
                                                                        fmt.Println(fi50, v55, v55)
                                                                    }
                                                                    v110 := GenS15{F0: 1, F1: 4294967295}
                                                                    v111 := GenS16{F0: (uint(v2) < 891)}
                                                                    switch "" {
                                                                    case "foo bar":
                                                                        switch "" {
                                                                        case "hello":
                                                                            var v112 string = (("\n\t" + ("\n\t" + "")) + "")
                                                                            v113 := GenS17{F0: -907, F1: (fi51 + (fi51 & (fi51 - (fi50 | fi51))))}
                                                                            delete(v4, v112)
                                                                            v114 := GenS8{F0: false, F1: fi51, F2: v34}
                                                                            _ = v112
                                                                            _ = v113
                                                                            _ = v114
                                                                        case "\n\t":
                                                                            v89 = (v89 + 32767)
                                                                        case "a":
                                                                            v81 = v34
                                                                        default:
                                                                            if true {
                                                                                v89 = (^(v89))
                                                                                v115 := GenS18{F0: uint16(v55), F1: 306, F2: ((("hello" + "\n\t") + "") + "")}
                                                                                var v116 float32 = v57
                                                                                _ = v115
                                                                                _ = v116
                                                                            } else {
                                                                                if ((0 != 42) || true) {
                                                                                    if (-128 != 42) {
                                                                                        var v117 int = fi50
                                                                                        _ = v117
                                                                                    }
                                                                                    switch (-(fi51)) {
                                                                                    case -9223372036854775808:
                                                                                        v44 = v44
                                                                                    case 4294967295:
                                                                                        fmt.Println(fi51, v34, v44)
                                                                                    default:
                                                                                        v118 := GenS19{F0: 100, F1: (v2 + (v2 / ((v2) | 1))), F2: 0}
                                                                                        _ = v118
                                                                                    }
                                                                                }
                                                                            }
                                                                            switch "" {
                                                                            case "hello":
                                                                                fmt.Println(v57, v44, v44)
                                                                                if true {
                                                                                    fmt.Println(v81, v89, v55)
                                                                                    _ = len(v5)
                                                                                    v119 := GenS20{F0: true, F1: (v89 & (v89 << ((616) & 63)))}
                                                                                    v81 = (v81 % (((v81 % ((v81) | 1))) | 1))
                                                                                    _ = v119
                                                                                } else {
                                                                                    var v120 int32 = 2147483647
                                                                                    _ = v120
                                                                                }
                                                                            case "\n\t":
                                                                                v121 := GenS9{F0: 4294967295, F1: fi51}
                                                                                _ = v121
                                                                            }
                                                                        }
                                                                    default:
                                                                        switch fi51 {
                                                                        case 320:
                                                                            switch fi51 {
                                                                            case 9223372036854775807:
                                                                                v55 += (^(v55))
                                                                            case -91:
                                                                                if (!((^(v81)) == v34)) {
                                                                                    if len(v1) > 0 {
                                                                                        v1[0] = "a"
                                                                                    }
                                                                                    v34 -= v81
                                                                                    v122 := GenS16{F0: (true && (2 > 0))}
                                                                                    switch -284 {
                                                                                    case 32767:
                                                                                        v123 := v3[9223372036854775807]
                                                                                        _ = v123
                                                                                    }
                                                                                    _ = v122
                                                                                } else {
                                                                                    var v124 int32 = v81
                                                                                    v125 := GenS21{F0: (v2 % (((^(v2))) | 1)), F1: ((-683 < genFn0()) && (!((1 == (v2 / ((v2) | 1))) && true)))}
                                                                                    if false {
                                                                                        switch ("\n\t" + "") {
                                                                                        case "":
                                                                                            v126 := v4[("hello" + "a")]
                                                                                            switch -2147483648 {
                                                                                            case 527:
                                                                                                v127 := GenS5{F0: 531, F1: 1}
                                                                                                _ = v127
                                                                                            case 128:
                                                                                                v128 := GenS22{F0: 441, F1: 956, F2: (v2 / ((v2) | 1))}
                                                                                                _ = v128
                                                                                            default:
                                                                                                var v129 float32 = v57
                                                                                                _ = v129
                                                                                            }
                                                                                            v57 += 1e10
                                                                                            var v130 uint16 = 65535
                                                                                            fmt.Println(v55, v44)
                                                                                            _ = v126
                                                                                            _ = v130
                                                                                        }
                                                                                    } else {
                                                                                        if true {
                                                                                            fmt.Println(v34)
                                                                                            v2 = 133
                                                                                        } else {
                                                                                            v81 = v81
                                                                                            v131 := GenS14{F0: 100}
                                                                                            var v132 string = "hello"
                                                                                            _ = v131
                                                                                            _ = v132
                                                                                        }
                                                                                    }
                                                                                    _ = v124
                                                                                    _ = v125
                                                                                }
                                                                            case 508:
                                                                                v44 = (v44 * v44)
                                                                            }
                                                                        }
                                                                    }
                                                                    _ = v110
                                                                    _ = v111
                                                                } else {
                                                                    v133 := GenS12{F0: v2, F1: v2, F2: 103}
                                                                    _ = v133
                                                                }
                                                                v89 = genFn0()
                                                                _ = v90
                                                            } else {
                                                                if (!(!true)) {
                                                                    delete(v3, 42)
                                                                }
                                                                switch (fi51 * 32768) {
                                                                case -2147483648:
                                                                    fmt.Println(fi50)
                                                                case 681:
                                                                    var v134 float64 = v44
                                                                    _ = v134
                                                                case 294:
                                                                    v55 += v55
                                                                    var v135 uint = 0
                                                                    _ = v135
                                                                default:
                                                                    v57 = float32(v44)
                                                                }
                                                                var v136 float64 = (v44 / (float64((v57 - 0.0)) + v44))
                                                                _ = v136
                                                            }
                                                        } else {
                                                            var v137 int32 = (^(v81))
                                                            _ = v137
                                                        }
                                                        if (42 > (v2 - v2)) {
                                                            v55 = v55
                                                            v81 -= (v34 & 522)
                                                            if true {
                                                                v89 = -229
                                                                switch (fi51 / (((^(fi50))) | 1)) {
                                                                case 0:
                                                                    fmt.Println(v89, v55, v44)
                                                                case 2147483647:
                                                                    v81 = v81
                                                                }
                                                                switch fi50 {
                                                                case 42:
                                                                    v81 = v81
                                                                default:
                                                                    var v138 float32 = ((v57 + v57) / 1.0)
                                                                    _ = v138
                                                                }
                                                                if false {
                                                                    _ = len(v3)
                                                                } else {
                                                                    var v139 int8 = 56
                                                                    var v140 uint8 = (v2 | (^(v2)))
                                                                    _ = v139
                                                                    _ = v140
                                                                }
                                                            }
                                                            if true {
                                                                v141 := GenS5{F0: 9223372036854775807, F1: (v2 ^ v2)}
                                                                v142 := GenS4{F0: 1, F1: (^(fi50)), F2: (v2 | v2)}
                                                                switch (fi50 + 2147483647) {
                                                                case 1:
                                                                    var v143 uint16 = 995
                                                                    _ = v143
                                                                case 128:
                                                                    v144 := GenS23{F0: fi51, F1: (v55 < 2), F2: (v81 * -4)}
                                                                    _ = v144
                                                                default:
                                                                    var v145 bool = false
                                                                    _ = v145
                                                                }
                                                                _ = v141
                                                                _ = v142
                                                            }
                                                        } else {
                                                            fmt.Println(v81, v55, fi50)
                                                            var v146 uint32 = 66
                                                            _ = v146
                                                        }
                                                        var v147 int64 = 42
                                                        _ = v147
                                                    }
                                                    fmt.Println(v57, v81)
                                                    _ = v88
                                                    _ = v89
                                                }
                                                v148 := GenS24{F0: -475, F1: 255, F2: 18446744073709551615}
                                                fmt.Println(v81)
                                                _ = v148
                                            }
                                        default:
                                            var v149 int = (fi50 * (fi50 << ((0) & 63)))
                                            _ = v149
                                        }
                                    default:
                                        v55 += v55
                                    }
                                    v150 := GenS25{F0: (v81 + (v34 ^ int32(v55)))}
                                    if true {
                                        if ((0.0 * 0.0) <= (0.1 / 1.5)) {
                                            switch fi51 {
                                            case -896:
                                                if false {
                                                    _ = len(v1)
                                                    v151 := GenS26{F0: ("a" + "\n\t"), F1: -1, F2: (1 == 2)}
                                                    _ = v151
                                                }
                                            default:
                                                v3[127] = 448
                                            }
                                            fmt.Println(v81, v44, v2)
                                            v81 += (v34 / ((-888) | 1))
                                        }
                                        if true {
                                            var v152 float32 = v57
                                            _ = v152
                                        } else {
                                            if ((true || ((^(v55)) != (v55 - v55))) || false) {
                                                v34 = (v81 - v34)
                                                fmt.Println(v57, v81)
                                                switch -9223372036854775808 {
                                                case 9223372036854775807:
                                                    if (10 == 0) {
                                                        if (!(!(!false))) {
                                                            v55 = (v55 / ((v55) | 1))
                                                            var v153 uint16 = 545
                                                            _ = v153
                                                        }
                                                    } else {
                                                        if true {
                                                            if false {
                                                                if len(v5) >= 2 {
                                                                    v5 = v5[1:len(v5)]
                                                                }
                                                                v154 := GenS27{F0: 918}
                                                                _ = v154
                                                            } else {
                                                                v57 -= v57
                                                                _ = len(v1)
                                                                switch fi51 {
                                                                case 337:
                                                                    var v155 int64 = 42
                                                                    _ = v155
                                                                case 806:
                                                                    fmt.Println(v81)
                                                                case -9223372036854775808:
                                                                    if len(v5) >= 2 {
                                                                        v5 = v5[1:len(v5)]
                                                                    }
                                                                default:
                                                                    var v156 uint64 = v55
                                                                    _ = v156
                                                                }
                                                            }
                                                            if (-127 > 1) {
                                                                var v157 int16 = genFn0()
                                                                _ = v157
                                                            }
                                                        }
                                                    }
                                                default:
                                                    var v158 bool = false
                                                    _ = v158
                                                }
                                            }
                                            var v159 uint8 = 174
                                            if true {
                                                if (((3.14159 - float64(0.1)) / 1e10) >= float64((v57 / ((v57 - v57) / 1.0)))) {
                                                    var v160 bool = true
                                                    var v161 int64 = -9223372036854775808
                                                    _ = v160
                                                    _ = v161
                                                }
                                                var v162 bool = false
                                                _ = v162
                                            } else {
                                                v163 := GenS19{F0: 100, F1: v159, F2: v81}
                                                var v164 int64 = -9223372036854775808
                                                if (v44 <= 0.1) {
                                                    v165 := GenS28{F0: (313 >= fi50)}
                                                    _ = v165
                                                } else {
                                                    _ = len(v1)
                                                }
                                                var v166 uint32 = 329
                                                _ = v163
                                                _ = v164
                                                _ = v166
                                            }
                                            _ = v159
                                        }
                                        var v167 string = "\n\t"
                                        _ = v167
                                    }
                                    _ = v150
                                } else {
                                    var v168 string = ""
                                    v55 = 42
                                    fmt.Println(v2, v57, v168)
                                    _ = v168
                                }
                            default:
                                if (v55 <= 145) {
                                    v57 += (float32(0.0) / v57)
                                    _ = len(v5)
                                }
                            }
                            _ = v81
                        }
                    } else {
                        if (42 != 1) {
                            v2 = (v2 % (((^(v2))) | 1))
                            var v169 int8 = 0
                            fmt.Println(v169, v55)
                            _ = v169
                        }
                        v4[("hello" + "foo bar")] = true
                        v1 = append(v1, "a")
                    }
                    var v170 bool = true
                    v171 := v4["a"]
                    if false {
                        delete(v4, ("foo bar" + "a"))
                        v55 -= (^(v55))
                        fmt.Println(v34, v55)
                    } else {
                        var v172 float64 = v44
                        fmt.Println(v44)
                        _ = v172
                    }
                    if true {
                        fmt.Println(v170, v2, v57)
                        v1 = append(v1, "")
                    }
                    _ = v170
                    _ = v171
                }
                switch (fi51 >> ((820) & 63)) {
                case -844:
                    v173 := GenS29{F0: genFn0(), F1: 0}
                    v174 := GenS30{F0: 42}
                    v57 += (0.1 * v57)
                    _ = v173
                    _ = v174
                case 2147483647:
                    var v175 uint16 = 2
                    _ = v175
                case -35:
                    switch "a" {
                    case "":
                        var v176 uint8 = v2
                        _ = v176
                    }
                }
                _ = v57
            }
            _ = fi51
            _ = v55
        }
        var v177 int = -9223372036854775808
        _ = fi50
        _ = v177
    }
    switch 127 {
    case -9223372036854775808:
        v178 := GenS31{F0: 65535}
        _ = v178
    case 4294967295:
        v44 += 0.1
    default:
        switch 779 {
        case 127:
            _ = len(v4)
        case 62:
            v179 := GenS27{F0: 340}
            _ = v179
        }
    }
    var v180 int8 = -1
    var v181 int8 = v180
    if false {
        fmt.Println(v44)
        if true {
            var v182 int8 = v181
            if len(v1) > 0 {
                v1[0] = ""
            }
            v183 := GenS5{F0: -9223372036854775808, F1: v2}
            _ = v182
            _ = v183
        }
        fmt.Println(v2, v181, v44)
    }
    v184 := GenS32{F0: 537}
    _ = len(v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v180:", v180)
    fmt.Println("v181:", v181)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v180
    _ = v181
    _ = v184
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v44
    _ = v5
}
