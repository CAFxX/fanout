package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint64; F1 int }
type GenS5 struct { F0 int32; F1 int32 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint; F1 uint16 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int32; F1 uint32; F2 int64 }
type GenS11 struct { F0 string; F1 uint32 }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 int8; F1 int8 }
type GenS14 struct { F0 int; F1 int16; F2 bool }
type GenS15 struct { F0 int8 }
type GenS16 struct { F0 uint64; F1 bool }

func main() {
    var v1 int32 = 42
    v2 := make(map[int]bool)
    v3 := make(map[string]int)
    fmt.Println(v1)
    v1 -= (v1 & -37)
    if (false || false) {
        v4 := GenS1{F0: 274}
        _ = v4
    }
    if true {
        switch "" {
        case "\n\t":
            v1 = v1
        case "hello":
            fmt.Println(v1)
        case "a":
            v1 += v1
        }
        v1 += (v1 + (v1 << ((65535) & 63)))
    }
    if (1 != 196) {
        var v5 uint16 = 32767
        var v6 int16 = -1
        _ = v5
        _ = v6
    }
    fmt.Println(v1)
    switch -856 {
    case 127:
        switch "foo bar" {
        case "":
            v7 := GenS2{F0: 42, F1: 55}
            _ = v7
        case "a":
            fmt.Println(v1)
        case "\n\t":
            var v8 int16 = int16(v1)
            _ = v8
        default:
            v9 := []int{}
            _ = v9
        }
    case 42:
        v1 = (v1 ^ v1)
    default:
        fmt.Println(v1)
    }
    fmt.Println(v1)
    var v10 uint16 = 65535
    if (1.5 > (-2.25 + 0.1)) {
        for fi11 := 0; fi11 < 13; fi11++ {
            for fi12 := 0; fi12 < 2; fi12++ {
                v13 := GenS1{F0: v10}
                _ = fi12
                _ = v13
            }
            _ = len(v3)
            _ = fi11
        }
        switch "foo bar" {
        case "":
            var v14 int8 = 100
            _ = v14
        case "\n\t":
            v15 := []int{32767, -2147483648}
            _ = v15
        }
        if ("" <= ("foo bar" + "a")) {
            v16 := v3[""]
            var v17 int = 1
            _ = v16
            _ = v17
        } else {
            for fi18 := 0; fi18 < 17; fi18++ {
                switch (fi18 | (fi18 ^ (^(fi18)))) {
                case 32768:
                    v19 := GenS3{F0: 0}
                    _ = v19
                case 2147483647:
                    var v20 uint8 = 91
                    _ = v20
                default:
                    v1 += v1
                }
                _ = fi18
            }
            v21 := []int{}
            _ = v21
        }
    }
    fmt.Println(v10, v1)
    v10 += (v10 % ((42) | 1))
    v22 := GenS4{F0: 0, F1: 731}
    var v23 string = ("" + "foo bar")
    if (843 <= 2147483647) {
        v10 -= (v10 * (v10 >> ((890) & 63)))
        fmt.Println(v23)
    } else {
        v1 = v1
        if false {
            for fi24 := 0; fi24 < 20; fi24++ {
                var v25 uint16 = v10
                for fi26 := 0; fi26 < 3; fi26++ {
                    if false {
                        v1 += v1
                    } else {
                        _ = len(v2)
                        v23 = "foo bar"
                    }
                    fmt.Println(fi24, v23, v10)
                    if true {
                        v27 := v2[(fi26 & 127)]
                        v28 := GenS5{F0: v1, F1: -1}
                        _ = v27
                        _ = v28
                    } else {
                        fmt.Println(v1, fi26, v23)
                        v29 := []int{(fi24 >> ((4294967295) & 63)), fi24, fi26, 0}
                        var v30 uint64 = 971
                        _ = v29
                        _ = v30
                    }
                    _ = fi26
                }
                v23 = v23
                _ = fi24
                _ = v25
            }
            for fi31 := 0; fi31 < 1; fi31++ {
                for fi32 := 0; fi32 < 11; fi32++ {
                    var v33 int64 = -1
                    var v34 int8 = 42
                    _ = fi32
                    _ = v33
                    _ = v34
                }
                v35 := []int{9223372036854775807, fi31, (fi31 + (-(fi31)))}
                _ = fi31
                _ = v35
            }
        }
        v36 := GenS1{F0: (v10 * (v10 / ((v10) | 1)))}
        _ = v36
    }
    if true {
        if (-2.25 <= float32(0.0)) {
            _ = len(v3)
        } else {
            v37 := []int{}
            var v38 int8 = -1
            fmt.Println(v10, v38)
            _ = v37
            _ = v38
        }
        var v39 int64 = 9223372036854775807
        if false {
            fmt.Println(v1, v10, v39)
        } else {
            var v40 bool = false
            for fi41 := 0; fi41 < 4; fi41++ {
                switch (("foo bar" + v23) + ("foo bar" + v23)) {
                case "":
                    var v42 int = (^(fi41))
                    _ = v42
                case "hello":
                    var v43 int = 128
                    _ = v43
                case "\n\t":
                    fmt.Println(v10)
                }
                var v44 int64 = (-(v39))
                if (v10 > (^(v10))) {
                    switch v23 {
                    case "\n\t":
                        v2[fi41] = v40
                    case "hello":
                        switch (fi41 & -672) {
                        case 363:
                            v40 = ((0.1 <= 0.1) && (false || false))
                        case 127:
                            v40 = (!v40)
                        default:
                            for fi45 := 0; fi45 < 5; fi45++ {
                                v10 = (v10 | (v10 ^ (v10 / (((v10 - (v10 ^ v10))) | 1))))
                                v2[int(v44)] = (4294967295 == 52)
                                v46 := GenS2{F0: 42, F1: 42}
                                _ = fi45
                                _ = v46
                            }
                        }
                    default:
                        if ((!(1 != 683)) && v40) {
                            if v40 {
                                switch v23 {
                                case "":
                                    v47 := GenS5{F0: 2147483647, F1: (v1 + (v1 ^ (^(v1))))}
                                    _ = v47
                                case "foo bar":
                                    switch 42 {
                                    case 42:
                                        var v48 bool = false
                                        _ = v48
                                    case -366:
                                        fmt.Println(v40, v39, v40)
                                    default:
                                        for fi49 := 0; fi49 < 13; fi49++ {
                                            if (((-2.25 / 1.5) * float32(3.14159)) > (0.0 + ((float32((((float64((float32((0.1 * ((-2.25 - 0.0) * -2.25))) * 1e10)) - (((3.14159 + (1.5 / 0.1)) / 3.14159) + (1.5 + 1e10))) - 1e10) + (((-2.25 + (0.0 + 0.0)) + float64(float32((-2.25 - 3.14159)))) / 1e10))) - 0.1) - float32(1.5)))) {
                                                v44 = -962
                                                var v50 uint = 128
                                                v39 = (v44 + v39)
                                                _ = v50
                                            } else {
                                                v51 := GenS3{F0: fi41}
                                                switch (v23 + "hello") {
                                                case "a":
                                                    if true {
                                                        switch (((v23 + "") + v23) + v23) {
                                                        case "hello":
                                                            var v52 uint32 = 1
                                                            _ = v52
                                                        case "foo bar":
                                                            if (((52 <= 54) || true) || v40) {
                                                                var v53 float32 = ((0.0 / -2.25) - 3.14159)
                                                                if (!v40) {
                                                                    var v54 uint64 = 4294967295
                                                                    delete(v2, (^(fi49)))
                                                                    var v55 int32 = v1
                                                                    var v56 int16 = 42
                                                                    _ = v54
                                                                    _ = v55
                                                                    _ = v56
                                                                }
                                                                v57 := GenS1{F0: 32767}
                                                                _ = v53
                                                                _ = v57
                                                            }
                                                        default:
                                                            var v58 int32 = v1
                                                            _ = v58
                                                        }
                                                        if true {
                                                            var v59 int64 = v44
                                                            _ = v59
                                                        }
                                                    } else {
                                                        v40 = (2147483647 >= 2)
                                                        var v60 float32 = 0.1
                                                        v61 := GenS5{F0: v1, F1: -1}
                                                        _ = v60
                                                        _ = v61
                                                    }
                                                case "hello":
                                                    var v62 float64 = (0.0 - -2.25)
                                                    _ = v62
                                                default:
                                                    v40 = false
                                                }
                                                _ = v51
                                            }
                                            _ = fi49
                                        }
                                    }
                                case "hello":
                                    var v63 string = ("" + ((v23 + ("foo bar" + "foo bar")) + v23))
                                    _ = v63
                                default:
                                    v10 += (v10 | v10)
                                }
                                if false {
                                    if ((-68 < 127) && false) {
                                        v44 -= 93
                                        var v64 int8 = 55
                                        _ = v64
                                    }
                                } else {
                                    v2[fi41] = (!false)
                                }
                            } else {
                                fmt.Println(v10, v23, v39)
                                v39 = (v44 ^ 1)
                            }
                            v44 += v39
                            switch v23 {
                            case "hello":
                                if ((v23 + v23) < v23) {
                                    for fi65 := 0; fi65 < 4; fi65++ {
                                        if (!((!v40) && v40)) {
                                            if (-19 > 1) {
                                                var v66 float64 = (float64((3.14159 * (3.14159 * (1.5 * (3.14159 - 0.0))))) + -2.25)
                                                _ = v66
                                            }
                                            var v67 int32 = -1
                                            var v68 string = v23
                                            _ = v67
                                            _ = v68
                                        } else {
                                            if v40 {
                                                v39 -= 1
                                            } else {
                                                v39 += 42
                                            }
                                            if true {
                                                if ((v1 >= 1) && (127 <= 26)) {
                                                    var v69 uint32 = 712
                                                    v10 = 0
                                                    _ = v69
                                                }
                                            } else {
                                                if v40 {
                                                    var v70 float32 = 3.14159
                                                    if false {
                                                        if (!((v1 + (-(v1))) >= -2147483648)) {
                                                            fmt.Println(v10, v1)
                                                        } else {
                                                            v71 := v2[(fi41 * fi41)]
                                                            v72 := GenS6{F0: 18446744073709551615}
                                                            _ = v71
                                                            _ = v72
                                                        }
                                                    } else {
                                                        switch (fi41 % ((-595) | 1)) {
                                                        case 32767:
                                                            v73 := []int{1}
                                                            _ = v73
                                                        case -9223372036854775808:
                                                            if ("hello" >= "foo bar") {
                                                                v40 = (1 == 2)
                                                            }
                                                            var v74 uint32 = 4294967295
                                                            _ = v74
                                                        case 32768:
                                                            if (v70 <= (v70 / 1e10)) {
                                                                if ((!v40) && (!(v44 != v39))) {
                                                                    v75 := GenS7{F0: 807, F1: 1}
                                                                    v1 += (v1 ^ v1)
                                                                    _ = v75
                                                                }
                                                                switch (fi41 / (((fi65 | fi65)) | 1)) {
                                                                case 127:
                                                                    if (!v40) {
                                                                        var v76 int64 = (v39 % ((0) | 1))
                                                                        v44 = (-(v44))
                                                                        _ = v76
                                                                    }
                                                                case -9223372036854775808:
                                                                    switch (^(fi65)) {
                                                                    case 128:
                                                                        v2[575] = (!v40)
                                                                    }
                                                                case -49:
                                                                    v77 := GenS2{F0: 65535, F1: 73}
                                                                    var v78 float32 = v70
                                                                    _ = v77
                                                                    _ = v78
                                                                }
                                                                v39 = v44
                                                            }
                                                        }
                                                        v79 := []int{fi41}
                                                        v80 := GenS1{F0: v10}
                                                        _ = v79
                                                        _ = v80
                                                    }
                                                    if v40 {
                                                        fmt.Println(v10, fi65)
                                                    } else {
                                                        v70 += 1.5
                                                        switch fi65 {
                                                        case 1:
                                                            v10 = (v10 ^ 237)
                                                        default:
                                                            if true {
                                                                fmt.Println(v70, fi65)
                                                                v23 = ((((("hello" + v23) + "") + ("hello" + "a")) + (v23 + v23)) + (v23 + v23))
                                                                v81 := GenS3{F0: (fi65 ^ 0)}
                                                                _ = v81
                                                            }
                                                        }
                                                    }
                                                    v40 = v40
                                                    _ = v70
                                                } else {
                                                    v82 := GenS4{F0: 4294967295, F1: fi65}
                                                    _ = v82
                                                }
                                                fmt.Println(fi65)
                                            }
                                            v10 = (v10 + v10)
                                        }
                                        var v83 uint32 = 0
                                        _ = fi65
                                        _ = v83
                                    }
                                    switch fi41 {
                                    case 1:
                                        v84 := GenS8{F0: (v10 / (((v10 % ((1) | 1))) | 1))}
                                        _ = v84
                                    default:
                                        if v40 {
                                            fmt.Println(v10, v39)
                                            for fi85 := 0; fi85 < 12; fi85++ {
                                                fmt.Println(v23, v23, fi41)
                                                v1 = v1
                                                if v40 {
                                                    switch fi85 {
                                                    case 634:
                                                        var v86 uint64 = 943
                                                        _ = v86
                                                    case 0:
                                                        v23 = "hello"
                                                        delete(v3, (("" + v23) + v23))
                                                        if (v40 && v40) {
                                                            switch -361 {
                                                            case 4294967295:
                                                                v39 -= -1
                                                            case 2147483647:
                                                                switch (fi41 | -2147483648) {
                                                                case 0:
                                                                    var v87 float64 = 3.14159
                                                                    _ = v87
                                                                case 42:
                                                                    switch v23 {
                                                                    case "":
                                                                        if true {
                                                                            var v88 int8 = 0
                                                                            var v89 string = v23
                                                                            _ = v88
                                                                            _ = v89
                                                                        }
                                                                    case "a":
                                                                        var v90 bool = v40
                                                                        _ = v90
                                                                    default:
                                                                        switch (fi85 * (fi85 * (fi41 | -817))) {
                                                                        case 128:
                                                                            v91 := GenS9{F0: 932}
                                                                            _ = v91
                                                                        case -9223372036854775808:
                                                                            fmt.Println(v23)
                                                                        }
                                                                    }
                                                                case 1:
                                                                    v92 := GenS10{F0: -593, F1: 0, F2: (v44 << ((10) & 63))}
                                                                    _ = v92
                                                                }
                                                            case 0:
                                                                if v40 {
                                                                    var v93 uint8 = 195
                                                                    _ = v93
                                                                } else {
                                                                    var v94 int = (fi85 >> ((42) & 63))
                                                                    if v40 {
                                                                        v1 += v1
                                                                        if (v40 || v40) {
                                                                            v95 := []int{fi41, v94, (-(fi41))}
                                                                            if (!v40) {
                                                                                fmt.Println(fi85, v23)
                                                                                if (((v23 + "a") + "\n\t") >= "") {
                                                                                    if (112 > 127) {
                                                                                        fmt.Println(v94)
                                                                                    } else {
                                                                                        fmt.Println(v44)
                                                                                        v44 += v39
                                                                                        var v96 int8 = 1
                                                                                        _ = v96
                                                                                    }
                                                                                    _ = len(v95)
                                                                                    if (v40 && (0 < -9223372036854775808)) {
                                                                                        v97 := GenS11{F0: v23, F1: 1}
                                                                                        var v98 uint64 = 0
                                                                                        v98 = (v98 * v98)
                                                                                        _ = v97
                                                                                        _ = v98
                                                                                    }
                                                                                } else {
                                                                                    if v40 {
                                                                                        var v99 int64 = v44
                                                                                        fmt.Println(v10, v44)
                                                                                        _ = v99
                                                                                    }
                                                                                    if true {
                                                                                        delete(v2, v94)
                                                                                        switch "hello" {
                                                                                        case "\n\t":
                                                                                            var v100 int64 = v39
                                                                                            var v101 string = v23
                                                                                            v95 = append(v95, fi41)
                                                                                            var v102 int8 = 127
                                                                                            var v103 uint64 = 4294967295
                                                                                            var v104 int64 = v44
                                                                                            _ = v100
                                                                                            _ = v101
                                                                                            _ = v102
                                                                                            _ = v103
                                                                                            _ = v104
                                                                                        }
                                                                                    } else {
                                                                                        v105 := GenS3{F0: 2147483647}
                                                                                        fmt.Println(fi85, v39, fi85)
                                                                                        var v106 float32 = (-2.25 / 1.5)
                                                                                        if false {
                                                                                            v107 := GenS12{F0: 0}
                                                                                            _ = v107
                                                                                        } else {
                                                                                            switch ("\n\t" + v23) {
                                                                                            case "foo bar":
                                                                                                var v108 uint16 = v10
                                                                                                _ = v108
                                                                                            case "hello":
                                                                                                var v109 int64 = (^(v44))
                                                                                                _ = v109
                                                                                            default:
                                                                                                if (v106 >= (v106 - ((v106 / ((v106 + (1.5 * v106)) - (((v106 * v106) / (v106 + (v106 + 0.1))) / 1.0))) - (1.5 / 3.14159)))) {
                                                                                                    v106 -= (0.1 - (0.1 / ((v106 / (0.0 - (v106 - v106))) / v106)))
                                                                                                    fmt.Println(v1, v39)
                                                                                                    fmt.Println(fi85, v40, v23)
                                                                                                    v110 := GenS2{F0: 65535, F1: 239}
                                                                                                    _ = v110
                                                                                                } else {
                                                                                                    fmt.Println(fi41, v106)
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        _ = v105
                                                                                        _ = v106
                                                                                    }
                                                                                    var v111 uint32 = 60
                                                                                    _ = v111
                                                                                }
                                                                            }
                                                                            switch v23 {
                                                                            case "":
                                                                                switch fi85 {
                                                                                case 32767:
                                                                                    var v112 float64 = (0.0 + 3.14159)
                                                                                    _ = v112
                                                                                case 156:
                                                                                    v113 := GenS2{F0: 2147483647, F1: 255}
                                                                                    _ = v113
                                                                                default:
                                                                                    if (v40 || v40) {
                                                                                        var v114 int32 = (v1 * -151)
                                                                                        fmt.Println(v114)
                                                                                        v115 := GenS11{F0: v23, F1: 815}
                                                                                        _ = v114
                                                                                        _ = v115
                                                                                    }
                                                                                }
                                                                            case "\n\t":
                                                                                _ = len(v2)
                                                                            case "foo bar":
                                                                                fmt.Println(fi41, fi41, fi41)
                                                                                v94 += fi85
                                                                                var v116 uint64 = 953
                                                                                _ = v116
                                                                            }
                                                                            _ = v95
                                                                        } else {
                                                                            if (v40 || (65535 < 392)) {
                                                                                v117 := GenS9{F0: 4294967295}
                                                                                if (2 == 1) {
                                                                                    var v118 uint8 = 255
                                                                                    delete(v2, fi41)
                                                                                    var v119 int8 = 100
                                                                                    _ = v118
                                                                                    _ = v119
                                                                                } else {
                                                                                    var v120 uint64 = 42
                                                                                    fmt.Println(v44)
                                                                                    v121 := []int{32768, (v94 & (fi41 + 9223372036854775807))}
                                                                                    _ = v120
                                                                                    _ = v121
                                                                                }
                                                                                v40 = v40
                                                                                _ = v117
                                                                            }
                                                                        }
                                                                        v122 := v3[v23]
                                                                        _ = v122
                                                                    } else {
                                                                        v123 := GenS13{F0: 1, F1: -16}
                                                                        _ = v123
                                                                    }
                                                                    _ = v94
                                                                }
                                                            }
                                                            fmt.Println(fi85, fi41, v39)
                                                        }
                                                        var v124 uint16 = v10
                                                        if ("hello" <= v23) {
                                                            v125 := []int{(fi85 ^ fi41), (fi41 - fi41)}
                                                            _ = v125
                                                        }
                                                        v39 -= (v44 - (v39 << ((112) & 63)))
                                                        _ = v124
                                                    case 42:
                                                        fmt.Println(v40, v44, fi41)
                                                    }
                                                    var v126 int32 = 191
                                                    _ = v126
                                                }
                                                _ = fi85
                                            }
                                            v3[(("" + ("\n\t" + v23)) + v23)] = fi41
                                        }
                                    }
                                }
                            case "foo bar":
                                switch "" {
                                case "":
                                    v40 = v40
                                default:
                                    var v127 int16 = -32768
                                    _ = v127
                                }
                            default:
                                var v128 bool = (fi41 >= fi41)
                                _ = v128
                            }
                        } else {
                            v129 := GenS12{F0: 0}
                            fmt.Println(v39, v44, v39)
                            if true {
                                v130 := GenS2{F0: 2147483647, F1: 0}
                                if (4294967295 < 361) {
                                    fmt.Println(v1)
                                    for fi131 := 0; fi131 < 7; fi131++ {
                                        if true {
                                            if ((fi41 - fi131) == int(v44)) {
                                                v23 = v23
                                                if v40 {
                                                    v132 := []int{}
                                                    v133 := GenS12{F0: 42}
                                                    fmt.Println(fi41)
                                                    var v134 float64 = (((0.0 - (-2.25 + 3.14159)) * float64((1.5 - ((0.0 * (-2.25 - (0.0 / 1.5))) / ((0.0 + ((0.1 - 0.0) * (0.1 + 0.0))) * (0.0 - 0.1)))))) - 3.14159)
                                                    v135 := GenS2{F0: 42, F1: 87}
                                                    switch -122 {
                                                    case 1:
                                                        switch (^(fi41)) {
                                                        case -77:
                                                            var v136 int32 = v1
                                                            _ = v136
                                                        case 127:
                                                            v137 := GenS8{F0: (^(v10))}
                                                            _ = v137
                                                        case 32767:
                                                            v39 += v39
                                                        }
                                                    case 751:
                                                        var v138 int8 = -3
                                                        _ = v138
                                                    case 700:
                                                        v40 = v40
                                                    }
                                                    if false {
                                                        switch -716 {
                                                        case 127:
                                                            _ = len(v132)
                                                        case -1:
                                                            if ((v10 / (((v10 / ((v10) | 1))) | 1)) <= v10) {
                                                                v139 := GenS9{F0: 42}
                                                                v134 = v134
                                                                var v140 float64 = (v134 / 3.14159)
                                                                if ((fi131 <= fi41) || (("foo bar" + v23) < ("foo bar" + v23))) {
                                                                    if ((0 < (^(v44))) || v40) {
                                                                        v10 += 0
                                                                        v1 += -840
                                                                        var v141 uint32 = 2147483647
                                                                        _ = v141
                                                                    }
                                                                    v140 = (-2.25 / 1e10)
                                                                    if (0 == (v10 / (((^(v10))) | 1))) {
                                                                        if ((99 > 98) && (!false)) {
                                                                            v132 = append(v132, (fi41 << ((822) & 63)))
                                                                        }
                                                                        v3[((((("" + ((v23 + ("\n\t" + "\n\t")) + "hello")) + "a") + v23) + (v23 + "")) + v23)] = fi41
                                                                    } else {
                                                                        v23 = v23
                                                                        v40 = (v10 >= v10)
                                                                        fmt.Println(v10)
                                                                    }
                                                                }
                                                                _ = v139
                                                                _ = v140
                                                            } else {
                                                                switch v23 {
                                                                case "foo bar":
                                                                    v142 := GenS3{F0: -1}
                                                                    _ = v142
                                                                case "a":
                                                                    v44 += -1
                                                                case "\n\t":
                                                                    if (v40 && (v40 || false)) {
                                                                        fmt.Println(v10, v10)
                                                                        switch "a" {
                                                                        case "foo bar":
                                                                            var v143 uint32 = 65535
                                                                            v1 -= (v1 / ((v1) | 1))
                                                                            fmt.Println(v10, v44)
                                                                            if v40 {
                                                                                delete(v3, "\n\t")
                                                                            } else {
                                                                                switch (fi41 << ((255) & 63)) {
                                                                                case 32767:
                                                                                    v1 += (v1 | v1)
                                                                                case 42:
                                                                                    if false {
                                                                                        var v144 uint32 = (v143 & (v143 - (v143 & (^(v143)))))
                                                                                        v145 := GenS14{F0: fi41, F1: 1, F2: v40}
                                                                                        _ = v144
                                                                                        _ = v145
                                                                                    }
                                                                                }
                                                                                var v146 int64 = (v44 - v44)
                                                                                _ = v146
                                                                            }
                                                                            fmt.Println(fi131, v39, fi41)
                                                                            _ = v143
                                                                        }
                                                                    }
                                                                default:
                                                                    if v40 {
                                                                        var v147 int = -469
                                                                        _ = v147
                                                                    } else {
                                                                        if (!(v40 || (!(!(false || v40))))) {
                                                                            v39 -= -201
                                                                        }
                                                                        v1 -= (v1 ^ v1)
                                                                    }
                                                                }
                                                            }
                                                        case -2147483648:
                                                            fmt.Println(v39)
                                                        default:
                                                            v148 := GenS6{F0: 0}
                                                            _ = v148
                                                        }
                                                        if true {
                                                            var v149 float64 = (v134 - ((v134 - v134) + v134))
                                                            _ = v149
                                                        } else {
                                                            fmt.Println(v40, v39)
                                                            var v150 float32 = (1e10 * (1.5 - float32((1.5 + ((float64(1.5) - v134) * (v134 * v134))))))
                                                            v23 = ((v23 + v23) + v23)
                                                            var v151 uint16 = (v10 + v10)
                                                            _ = v150
                                                            _ = v151
                                                        }
                                                        switch "\n\t" {
                                                        case "":
                                                            var v152 int8 = -62
                                                            _ = v152
                                                        case "\n\t":
                                                            if true {
                                                                var v153 uint8 = 0
                                                                _ = v153
                                                            } else {
                                                                var v154 float32 = 0.0
                                                                fmt.Println(v10)
                                                                var v155 int64 = 445
                                                                _ = v154
                                                                _ = v155
                                                            }
                                                        }
                                                        v156 := GenS15{F0: 127}
                                                        _ = v156
                                                    } else {
                                                        v44 = (v44 / (((-(v44))) | 1))
                                                        var v157 int32 = v1
                                                        v1 -= (v157 / ((1) | 1))
                                                        _ = v157
                                                    }
                                                    fmt.Println(v40, v23)
                                                    v23 = v23
                                                    _ = v132
                                                    _ = v133
                                                    _ = v134
                                                    _ = v135
                                                } else {
                                                    var v158 uint8 = 183
                                                    _ = v158
                                                }
                                                var v159 int8 = 40
                                                _ = v159
                                            }
                                            var v160 uint8 = 255
                                            var v161 float64 = 0.1
                                            _ = v160
                                            _ = v161
                                        }
                                        if v40 {
                                            if v40 {
                                                fmt.Println(fi131, fi41, v10)
                                                var v162 uint8 = 127
                                                if (-446 < 42) {
                                                    if (((0.1 != -2.25) && v40) || (fi131 >= (fi41 << ((0) & 63)))) {
                                                        v1 = v1
                                                        v39 = v44
                                                    }
                                                    fmt.Println(fi41)
                                                }
                                                _ = v162
                                            }
                                            v3[("hello" + ((v23 + "\n\t") + v23))] = fi131
                                        } else {
                                            v40 = false
                                        }
                                        var v163 uint32 = 4294967295
                                        _ = fi131
                                        _ = v163
                                    }
                                    _ = len(v2)
                                }
                                var v164 int64 = (v39 | v39)
                                _ = v130
                                _ = v164
                            }
                            _ = v129
                        }
                    }
                    v165 := GenS13{F0: 1, F1: 88}
                    var v166 uint64 = 1
                    _ = v165
                    _ = v166
                }
                _ = fi41
                _ = v44
            }
            v2[9223372036854775807] = ((202 == uint8(v10)) && (105 < 2))
            _ = v40
        }
        _ = v39
    } else {
        for fi167 := 0; fi167 < 20; fi167++ {
            if false {
                var v168 uint8 = 42
                _ = v168
            }
            fmt.Println(v23, v23, v10)
            switch v23 {
            case "foo bar":
                for fi169 := 0; fi169 < 19; fi169++ {
                    var v170 int8 = 127
                    _ = fi169
                    _ = v170
                }
            case "\n\t":
                v1 -= -146
            case "a":
                fmt.Println(v23, v10)
            }
            _ = fi167
        }
        v23 = ""
    }
    if (2147483647 >= 2147483647) {
        var v171 int = -541
        v172 := GenS3{F0: v171}
        _ = v171
        _ = v172
    } else {
        switch v23 {
        case "hello":
            var v173 uint = 401
            _ = v173
        default:
            for fi174 := 0; fi174 < 8; fi174++ {
                if true {
                    v175 := GenS16{F0: 416, F1: (v1 == -125)}
                    _ = v175
                } else {
                    var v176 int32 = v1
                    if false {
                        var v177 int16 = 544
                        var v178 uint16 = v10
                        _ = v177
                        _ = v178
                    } else {
                        if (2 == 2) {
                            v23 = ("a" + (v23 + ""))
                            var v179 int16 = -243
                            _ = v179
                        }
                    }
                    _ = v176
                }
                _ = fi174
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v23
    _ = v3
}
