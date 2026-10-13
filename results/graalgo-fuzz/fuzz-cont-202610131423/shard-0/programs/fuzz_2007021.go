package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 uint16; F1 int8; F2 uint8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int64; F1 uint32; F2 uint32 }
type GenS7 struct { F0 int16; F1 uint32 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 uint; F1 uint32 }
type GenS12 struct { F0 string; F1 string; F2 int64 }
type GenS13 struct { F0 int }
type GenS14 struct { F0 uint; F1 uint16; F2 uint32 }
type GenS15 struct { F0 uint32; F1 bool }
type GenS16 struct { F0 int8 }
type GenS17 struct { F0 int64 }
type GenS18 struct { F0 uint32 }
type GenS19 struct { F0 int16; F1 uint32; F2 int8 }
type GenS20 struct { F0 int8; F1 int16; F2 int8 }
type GenS21 struct { F0 int16 }

func main() {
    v1 := []string{"", "\n\t", ((("a" + "hello") + "foo bar") + "a"), ("\n\t" + "a")}
    var v2 int16 = 32767
    var v3 int8 = -9
    var v4 uint8 = 222
    var v5 int32 = 294
    if false {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        var v6 bool = (((false && false) && (((0.0 + 0.1) == -2.25) || (((!true) && (695 >= 1)) || (0 < -9223372036854775808)))) && ((1 == 18446744073709551615) || (!false)))
        switch 127 {
        case 4294967295:
            if true {
                switch 128 {
                case 0:
                    var v7 float32 = -2.25
                    _ = v7
                case 393:
                    for fi8 := 0; fi8 < 5; fi8++ {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = fi8
                    }
                case 927:
                    var v9 int16 = v2
                    _ = v9
                }
                v10 := GenS1{F0: (v2 - 0)}
                for fi11 := 0; fi11 < 12; fi11++ {
                    var v12 int32 = v5
                    if ((-(v3)) <= (v3 % ((20) | 1))) {
                        if (!v6) {
                            fmt.Println(v12)
                            v3 = v3
                        } else {
                            var v13 uint = 520
                            for fi14 := 0; fi14 < 15; fi14++ {
                                fmt.Println(v3)
                                var v15 int64 = 9223372036854775807
                                _ = fi14
                                _ = v15
                            }
                            _ = v13
                        }
                        var v16 uint8 = 171
                        v17 := GenS1{F0: (v2 - (v2 & v2))}
                        _ = v16
                        _ = v17
                    }
                    _ = fi11
                    _ = v12
                }
                _ = v10
            } else {
                v2 = 32767
            }
        case 32767:
            var v18 int32 = 2147483647
            _ = v18
        case 0:
            if (4294967295 > 0) {
                v3 += v3
                _ = len(v1)
                var v19 int16 = v2
                _ = v19
            }
        default:
            var v20 int32 = (v5 / (((^(v5))) | 1))
            _ = v20
        }
        _ = v6
    }
    v4 = (^(v4))
    v4 -= v4
    if (497 == 42) {
        if (("foo bar" + "") < "a") {
            switch 32767 {
            case -9223372036854775808:
                if true {
                    if (18446744073709551615 == 597) {
                        v5 -= 42
                    }
                    for fi21 := 0; fi21 < 4; fi21++ {
                        for fi22 := 0; fi22 < 3; fi22++ {
                            var v23 int32 = v5
                            v23 += 1
                            _ = fi22
                            _ = v23
                        }
                        if false {
                            if (v2 == v2) {
                                var v24 uint64 = 4294967295
                                if (0.0 > (-2.25 - -2.25)) {
                                    v25 := GenS2{F0: (fi21 % ((2147483647) | 1)), F1: "a"}
                                    _ = v25
                                }
                                var v26 bool = true
                                _ = v24
                                _ = v26
                            }
                        }
                        _ = fi21
                    }
                } else {
                    if true {
                        if (18446744073709551615 <= 615) {
                            v27 := GenS2{F0: 127, F1: "\n\t"}
                            if (65535 <= 750) {
                                if false {
                                    for fi28 := 0; fi28 < 2; fi28++ {
                                        switch ("hello" + ("hello" + "")) {
                                        case "foo bar":
                                            v29 := GenS1{F0: -398}
                                            _ = v29
                                        case "hello":
                                            switch (fi28 & fi28) {
                                            case -9223372036854775808:
                                                var v30 float32 = -2.25
                                                _ = v30
                                            case 1:
                                                if false {
                                                    fmt.Println(v5, fi28)
                                                    if (v5 <= (v5 / ((v5) | 1))) {
                                                        v31 := GenS2{F0: 2147483647, F1: "foo bar"}
                                                        if false {
                                                            var v32 uint32 = 2
                                                            fmt.Println(v5, v2, fi28)
                                                            if false {
                                                                for fi33 := 0; fi33 < 15; fi33++ {
                                                                    var v34 string = ""
                                                                    _ = fi33
                                                                    _ = v34
                                                                }
                                                                v35 := GenS2{F0: (fi28 * (^(fi28))), F1: ("hello" + "hello")}
                                                                _ = v35
                                                            }
                                                            _ = v32
                                                        }
                                                        _ = v31
                                                    }
                                                } else {
                                                    for fi36 := 0; fi36 < 2; fi36++ {
                                                        v4 += (v4 + v4)
                                                        v5 = 0
                                                        _ = fi36
                                                    }
                                                    var v37 bool = true
                                                    var v38 string = "foo bar"
                                                    _ = v37
                                                    _ = v38
                                                }
                                            }
                                        case "":
                                            fmt.Println(fi28)
                                        }
                                        v5 = v5
                                        fmt.Println(v4)
                                        _ = fi28
                                    }
                                }
                            } else {
                                var v39 int8 = (v3 << ((2) & 63))
                                for fi40 := 0; fi40 < 19; fi40++ {
                                    switch "hello" {
                                    case "\n\t":
                                        switch fi40 {
                                        case 2147483647:
                                            var v41 float32 = float32(-2.25)
                                            _ = v41
                                        case -249:
                                            v42 := GenS2{F0: (fi40 * (fi40 / ((fi40) | 1))), F1: "\n\t"}
                                            _ = v42
                                        }
                                    default:
                                        var v43 int = (fi40 - int(v2))
                                        _ = v43
                                    }
                                    _ = fi40
                                }
                                fmt.Println(v4, v2, v2)
                                _ = v39
                            }
                            switch 1 {
                            case 32767:
                                var v44 int32 = (v5 ^ v5)
                                _ = v44
                            case 2147483647:
                                v2 += v2
                            case 128:
                                var v45 uint16 = 1
                                _ = v45
                            }
                            _ = v27
                        } else {
                            for fi46 := 0; fi46 < 10; fi46++ {
                                var v47 uint32 = 0
                                for fi48 := 0; fi48 < 11; fi48++ {
                                    v49 := GenS2{F0: (fi46 | (^(fi48))), F1: "\n\t"}
                                    _ = fi48
                                    _ = v49
                                }
                                _ = fi46
                                _ = v47
                            }
                            v1 = append(v1, "a")
                        }
                        var v50 int8 = v3
                        v50 = (v50 % (((^(v3))) | 1))
                        _ = v50
                    } else {
                        fmt.Println(v4)
                        switch "hello" {
                        case "\n\t":
                            if false {
                                fmt.Println(v4)
                                for fi51 := 0; fi51 < 13; fi51++ {
                                    var v52 uint64 = 2
                                    if true {
                                        if false {
                                            _ = len(v1)
                                            v3 = v3
                                        } else {
                                            fmt.Println(v52, v3)
                                            if ((-(v3)) <= 54) {
                                                switch fi51 {
                                                case 32767:
                                                    for fi53 := 0; fi53 < 13; fi53++ {
                                                        var v54 float32 = 3.14159
                                                        _ = fi53
                                                        _ = v54
                                                    }
                                                default:
                                                    if true {
                                                        switch "foo bar" {
                                                        case "foo bar":
                                                            for fi55 := 0; fi55 < 8; fi55++ {
                                                                var v56 int8 = (v3 / (((v3 << ((uint(v52)) & 63))) | 1))
                                                                _ = fi55
                                                                _ = v56
                                                            }
                                                        }
                                                        var v57 int = (fi51 & fi51)
                                                        _ = v57
                                                    } else {
                                                        if false {
                                                            if (0.0 < 1e10) {
                                                                var v58 uint64 = 864
                                                                switch fi51 {
                                                                case 488:
                                                                    switch "\n\t" {
                                                                    case "":
                                                                        var v59 int = fi51
                                                                        _ = v59
                                                                    case "\n\t":
                                                                        v4 = v4
                                                                    default:
                                                                        v5 = v5
                                                                    }
                                                                case -207:
                                                                    switch 2147483647 {
                                                                    case -119:
                                                                        var v60 int64 = 42
                                                                        _ = v60
                                                                    default:
                                                                        if len(v1) > 0 {
                                                                            v1[0] = "\n\t"
                                                                        }
                                                                    }
                                                                case 75:
                                                                    v61 := GenS1{F0: (v2 + v2)}
                                                                    _ = v61
                                                                default:
                                                                    if (1 != 285) {
                                                                        switch (("\n\t" + "foo bar") + (("a" + ("a" + ("\n\t" + "a"))) + "hello")) {
                                                                        case "hello":
                                                                            switch ("\n\t" + ((("hello" + "") + "foo bar") + "\n\t")) {
                                                                            case "\n\t":
                                                                                if (-2.25 > (1.5 + (-2.25 - float64(1e10)))) {
                                                                                    switch "hello" {
                                                                                    case "hello":
                                                                                        switch ("foo bar" + "") {
                                                                                        case "":
                                                                                            v3 = (v3 % (((-(v3))) | 1))
                                                                                        }
                                                                                    case "":
                                                                                        v2 -= int16(v52)
                                                                                    case "foo bar":
                                                                                        v58 = v58
                                                                                    }
                                                                                    var v62 int16 = (v2 / ((v2) | 1))
                                                                                    _ = v62
                                                                                } else {
                                                                                    v52 = 4294967295
                                                                                }
                                                                            case "":
                                                                                var v63 uint = 42
                                                                                _ = v63
                                                                            case "hello":
                                                                                for fi64 := 0; fi64 < 6; fi64++ {
                                                                                    v4 -= v4
                                                                                    var v65 int16 = v2
                                                                                    _ = fi64
                                                                                    _ = v65
                                                                                }
                                                                            }
                                                                        default:
                                                                            if true {
                                                                                v58 += uint64(v4)
                                                                            }
                                                                        }
                                                                        var v66 uint8 = uint8(v58)
                                                                        _ = v66
                                                                    }
                                                                }
                                                                _ = v58
                                                            }
                                                            for fi67 := 0; fi67 < 12; fi67++ {
                                                                if (712 <= 0) {
                                                                    if len(v1) > 0 {
                                                                        v1[0] = "\n\t"
                                                                    }
                                                                    fmt.Println(v2, v5)
                                                                } else {
                                                                    v68 := GenS3{F0: 8, F1: (v3 * v3), F2: v4}
                                                                    _ = v68
                                                                }
                                                                v5 += v5
                                                                fmt.Println(v5, v5)
                                                                _ = fi67
                                                            }
                                                            var v69 float32 = (3.14159 / float32(1.5))
                                                            _ = v69
                                                        } else {
                                                            v4 += v4
                                                            v70 := GenS4{F0: 65535}
                                                            _ = v70
                                                        }
                                                    }
                                                }
                                                fmt.Println(fi51, v2, fi51)
                                                v71 := GenS5{F0: v5}
                                                _ = v71
                                            }
                                        }
                                        if false {
                                            if len(v1) > 0 {
                                                v1[0] = "\n\t"
                                            }
                                        } else {
                                            v52 += v52
                                            _ = len(v1)
                                        }
                                    } else {
                                        v5 += v5
                                        v72 := GenS1{F0: (v2 + v2)}
                                        v5 += int32(v3)
                                        _ = v72
                                    }
                                    _ = fi51
                                    _ = v52
                                }
                            } else {
                                fmt.Println(v3, v3, v3)
                                if false {
                                    var v73 uint16 = 1
                                    if true {
                                        for fi74 := 0; fi74 < 4; fi74++ {
                                            for fi75 := 0; fi75 < 18; fi75++ {
                                                v73 += 88
                                                var v76 int64 = 9223372036854775807
                                                v3 += int8(v73)
                                                v3 = 1
                                                var v77 uint64 = 376
                                                _ = fi75
                                                _ = v76
                                                _ = v77
                                            }
                                            _ = fi74
                                        }
                                        var v78 int8 = int8(v2)
                                        v73 = (^(v73))
                                        _ = v78
                                    }
                                    var v79 uint16 = 0
                                    _ = v73
                                    _ = v79
                                }
                            }
                        case "a":
                            if false {
                                var v80 bool = (42 <= -689)
                                if v80 {
                                    var v81 int = -541
                                    if (v3 > (v3 | (v3 | v3))) {
                                        v3 = v3
                                    } else {
                                        v81 = v81
                                    }
                                    switch (v81 / ((v81) | 1)) {
                                    case 767:
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    case 32767:
                                        v80 = (!(!(v80 && (v4 > (v4 << ((0) & 63))))))
                                    case 673:
                                        v3 = (v3 >> ((255) & 63))
                                    default:
                                        fmt.Println(v4, v4, v2)
                                    }
                                    _ = v81
                                }
                                _ = v80
                            } else {
                                var v82 float64 = 0.1
                                var v83 int8 = v3
                                _ = v82
                                _ = v83
                            }
                        case "foo bar":
                            for fi84 := 0; fi84 < 5; fi84++ {
                                for fi85 := 0; fi85 < 6; fi85++ {
                                    switch 128 {
                                    case -1:
                                        fmt.Println(v5, v3, fi85)
                                    case 2147483647:
                                        v86 := GenS5{F0: (v5 + v5)}
                                        _ = v86
                                    }
                                    var v87 float32 = -2.25
                                    switch ("\n\t" + "") {
                                    case "hello":
                                        v4 = 128
                                    case "\n\t":
                                        _ = len(v1)
                                    }
                                    _ = fi85
                                    _ = v87
                                }
                                _ = len(v1)
                                if (((((float32(0.0) + (0.0 / 1.5)) + -2.25) / 3.14159) / 1.5) < 0.0) {
                                    v4 = (v4 / ((v4) | 1))
                                    switch "" {
                                    case "hello":
                                        v88 := GenS5{F0: (v5 * (v5 - (v5 * v5)))}
                                        _ = v88
                                    default:
                                        var v89 float32 = 0.0
                                        _ = v89
                                    }
                                    for fi90 := 0; fi90 < 17; fi90++ {
                                        var v91 float64 = -2.25
                                        if len(v1) > 0 {
                                            v1[0] = "foo bar"
                                        }
                                        _ = fi90
                                        _ = v91
                                    }
                                }
                                _ = fi84
                            }
                        default:
                            v92 := GenS4{F0: 42}
                            _ = v92
                        }
                        v4 = (^(v4))
                    }
                    v93 := GenS5{F0: -1}
                    _ = v93
                }
            case -34:
                switch ("foo bar" + (("\n\t" + ("\n\t" + "foo bar")) + "\n\t")) {
                case "a":
                    v2 += v2
                case "hello":
                    v3 += v3
                default:
                    if false {
                        var v94 bool = false
                        _ = v94
                    }
                }
            default:
                if (0 != 4294967295) {
                    fmt.Println(v3, v3)
                }
            }
        }
        fmt.Println(v2)
        var v95 float32 = 3.14159
        _ = v95
    } else {
        fmt.Println(v2, v5)
        if false {
            var v96 float64 = (0.1 * (3.14159 / 1.5))
            v4 -= v4
            if true {
                v97 := GenS6{F0: 42, F1: 292, F2: 2147483648}
                if (0 <= v5) {
                    var v98 uint = 88
                    var v99 string = ""
                    v100 := GenS4{F0: 255}
                    _ = v100
                    _ = v98
                    _ = v99
                }
                for fi101 := 0; fi101 < 14; fi101++ {
                    v1 = append(v1, ("\n\t" + "\n\t"))
                    for fi102 := 0; fi102 < 6; fi102++ {
                        v4 -= (v4 + (v4 / ((v4) | 1)))
                        switch (fi102 & fi101) {
                        case -9223372036854775808:
                            if true {
                                if true {
                                    fmt.Println(fi102, fi102, v4)
                                } else {
                                    if true {
                                        var v103 uint32 = 2
                                        _ = v103
                                    } else {
                                        switch "" {
                                        case "hello":
                                            v1 = append(v1, "hello")
                                        case "a":
                                            switch -226 {
                                            case -12:
                                                if (fi102 > (fi102 | (-(fi101)))) {
                                                    var v104 int16 = v2
                                                    v5 += 2147483647
                                                    v105 := GenS3{F0: 42, F1: -1, F2: (v4 << ((0) & 63))}
                                                    v104 = v2
                                                    if ((1 < v3) || false) {
                                                        v5 = v5
                                                        fmt.Println(fi102, v104, fi102)
                                                    } else {
                                                        var v106 uint = 4294967295
                                                        if false {
                                                            if (0.1 >= -2.25) {
                                                                v1 = append(v1, "")
                                                            } else {
                                                                var v107 float64 = v96
                                                                _ = v107
                                                            }
                                                            if (!true) {
                                                                v106 += v106
                                                                if len(v1) > 0 {
                                                                    v1[0] = "a"
                                                                }
                                                                if true {
                                                                    var v108 int = (-(fi102))
                                                                    _ = v108
                                                                }
                                                            } else {
                                                                var v109 int64 = -590
                                                                _ = len(v1)
                                                                _ = v109
                                                            }
                                                            v110 := GenS7{F0: (v104 | (v104 >> ((v106) & 63))), F1: 1}
                                                            if (-1 == -1) {
                                                                v111 := GenS7{F0: v2, F1: 65535}
                                                                _ = len(v1)
                                                                fmt.Println(v3, v3)
                                                                var v112 string = "hello"
                                                                v1 = append(v1, "a")
                                                                _ = v111
                                                                _ = v112
                                                            }
                                                            _ = v110
                                                        }
                                                        var v113 uint16 = 320
                                                        _ = v106
                                                        _ = v113
                                                    }
                                                    _ = v104
                                                    _ = v105
                                                }
                                            case -346:
                                                var v114 int64 = -790
                                                _ = v114
                                            case 4294967295:
                                                var v115 uint = 255
                                                _ = v115
                                            }
                                            v1 = append(v1, "foo bar")
                                            v116 := GenS5{F0: (v5 - (-(v5)))}
                                            if false {
                                                fmt.Println(v5)
                                                var v117 uint8 = 255
                                                v117 -= v4
                                                _ = v117
                                            }
                                            _ = v116
                                        case "":
                                            switch "foo bar" {
                                            case "\n\t":
                                                v2 = v2
                                            case "a":
                                                v118 := GenS8{F0: 42}
                                                _ = v118
                                            default:
                                                v3 = (v3 ^ (v3 + 22))
                                            }
                                        default:
                                            v96 += v96
                                        }
                                        var v119 int = (fi102 << ((18446744073709551615) & 63))
                                        _ = v119
                                    }
                                    if true {
                                        if false {
                                            v120 := GenS5{F0: v5}
                                            _ = v120
                                        } else {
                                            v96 = v96
                                            v5 += 2147483647
                                        }
                                        if len(v1) > 0 {
                                            v1[0] = "foo bar"
                                        }
                                    } else {
                                        v2 = (^(v2))
                                    }
                                    if true {
                                        if false {
                                            _ = len(v1)
                                        }
                                        if false {
                                            switch ("" + ("hello" + "hello")) {
                                            case "":
                                                v121 := GenS9{F0: true}
                                                _ = v121
                                            case "\n\t":
                                                fmt.Println(v4, v3)
                                            default:
                                                switch (^(fi101)) {
                                                case -2147483648:
                                                    fmt.Println(v5, fi101, v4)
                                                }
                                            }
                                            var v122 uint64 = 42
                                            _ = v122
                                        } else {
                                            var v123 float32 = 1e10
                                            if (v2 > 0) {
                                                var v124 uint = 42
                                                if len(v1) > 0 {
                                                    v1[0] = "\n\t"
                                                }
                                                _ = v124
                                            } else {
                                                var v125 int64 = -643
                                                if (fi101 == fi102) {
                                                    if (true || (147 <= 879)) {
                                                        v126 := GenS5{F0: v5}
                                                        _ = v126
                                                    } else {
                                                        v127 := GenS10{F0: 4294967295}
                                                        v128 := GenS11{F0: 542, F1: 4294967295}
                                                        v129 := GenS3{F0: 309, F1: int8(v4), F2: (v4 ^ (^(v4)))}
                                                        fmt.Println(fi101, v96, v125)
                                                        _ = v127
                                                        _ = v128
                                                        _ = v129
                                                    }
                                                    v130 := GenS3{F0: 42, F1: 47, F2: v4}
                                                    _ = v130
                                                } else {
                                                    var v131 int8 = 100
                                                    if true {
                                                        var v132 uint32 = 2147483648
                                                        fmt.Println(v3, v125)
                                                        _ = v132
                                                    } else {
                                                        var v133 uint16 = 947
                                                        _ = v133
                                                    }
                                                    _ = v131
                                                }
                                                _ = v125
                                            }
                                            _ = v123
                                        }
                                        v96 = v96
                                    } else {
                                        if true {
                                            switch fi102 {
                                            case 447:
                                                fmt.Println(fi102)
                                            case 1:
                                                v134 := GenS1{F0: 32767}
                                                _ = v134
                                            default:
                                                var v135 float32 = (((0.0 + (0.0 + 0.1)) + 0.1) - 1.5)
                                                _ = v135
                                            }
                                            var v136 string = ""
                                            _ = v136
                                        } else {
                                            var v137 int = int(v2)
                                            _ = v137
                                        }
                                        v1 = append(v1, ("\n\t" + ((("a" + ("hello" + "hello")) + "\n\t") + ("hello" + ""))))
                                        v2 = 1
                                    }
                                }
                                if (3.14159 < (float32(v96) * 1.5)) {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    v138 := GenS11{F0: 570, F1: 2}
                                    if ((^(v4)) < 247) {
                                        var v139 bool = (true || false)
                                        var v140 uint64 = 958
                                        _ = v139
                                        _ = v140
                                    } else {
                                        switch "a" {
                                        case "hello":
                                            v141 := GenS11{F0: 2, F1: 2147483647}
                                            _ = v141
                                        case "foo bar":
                                            switch (fi102 >> ((42) & 63)) {
                                            case -2147483648:
                                                v1 = append(v1, "")
                                            }
                                        case "":
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        }
                                        _ = len(v1)
                                        if ("foo bar" > "a") {
                                            if true {
                                                var v142 uint8 = (^(v4))
                                                _ = v142
                                            }
                                        } else {
                                            v3 += (v3 % ((v3) | 1))
                                            switch "a" {
                                            case "foo bar":
                                                v2 = (v2 ^ (v2 | v2))
                                            }
                                        }
                                    }
                                    _ = v138
                                } else {
                                    var v143 uint32 = 2
                                    var v144 float32 = (0.1 * ((((float32((v96 + v96)) - float32(0.1)) + 1.5) / -2.25) * -2.25))
                                    v145 := GenS12{F0: "", F1: "foo bar", F2: 1}
                                    _ = v143
                                    _ = v144
                                    _ = v145
                                }
                            } else {
                                v4 = v4
                                var v146 float64 = ((v96 / float64(1.5)) / v96)
                                _ = v146
                            }
                        case -338:
                            if ((v96 - float64(((1e10 + ((((0.1 - 0.0) - (0.0 - 0.1)) / 1.0) + (0.1 - 1e10))) * (1.5 + 3.14159)))) < (v96 * (v96 * v96))) {
                                if true {
                                    v147 := GenS4{F0: 288}
                                    _ = v147
                                }
                                v148 := GenS11{F0: 262, F1: 4294967295}
                                v149 := GenS13{F0: 232}
                                if true {
                                    var v150 bool = (!(int8(v4) == v3))
                                    var v151 int = fi101
                                    _ = v150
                                    _ = v151
                                } else {
                                    v152 := GenS3{F0: 1, F1: v3, F2: 0}
                                    _ = len(v1)
                                    v153 := GenS5{F0: (v5 % (((v5 << ((0) & 63))) | 1))}
                                    _ = v152
                                    _ = v153
                                }
                                v2 = (v2 | 42)
                                var v154 uint32 = 42
                                v154 = 42
                                _ = v148
                                _ = v149
                                _ = v154
                            }
                        default:
                            var v155 int32 = v5
                            _ = v155
                        }
                        v156 := GenS5{F0: (^(v5))}
                        _ = fi102
                        _ = v156
                    }
                    var v157 uint = 1
                    _ = fi101
                    _ = v157
                }
                _ = v97
            } else {
                var v158 uint64 = 0
                if len(v1) > 0 {
                    v1[0] = "\n\t"
                }
                _ = v158
            }
            _ = v96
        } else {
            v159 := GenS14{F0: 4294967295, F1: 415, F2: 620}
            var v160 float32 = 0.1
            _ = v159
            _ = v160
        }
    }
    v5 = (v5 / ((1) | 1))
    for fi161 := 0; fi161 < 20; fi161++ {
        v162 := GenS15{F0: 42, F1: (1.5 < 3.14159)}
        _ = fi161
        _ = v162
    }
    v163 := GenS10{F0: 4294967295}
    for fi164 := 0; fi164 < 8; fi164++ {
        for fi165 := 0; fi165 < 19; fi165++ {
            v166 := GenS10{F0: 0}
            var v167 int64 = 1
            if ((float32(1.5) / float32((3.14159 + -2.25))) == float32(0.0)) {
                v168 := GenS12{F0: ("a" + ""), F1: "hello", F2: -9223372036854775808}
                v1 = append(v1, ("a" + "hello"))
                v169 := GenS3{F0: 397, F1: v3, F2: (v4 % ((128) | 1))}
                _ = v168
                _ = v169
            } else {
                v170 := GenS3{F0: 0, F1: v3, F2: v4}
                v5 = (v5 * -515)
                if ((fi165 ^ fi165) <= (fi164 * (^(fi165)))) {
                    if len(v1) > 0 {
                        v1[0] = "foo bar"
                    }
                    switch 1 {
                    case 127:
                        var v171 uint8 = v4
                        _ = v171
                    }
                    if ((65535 > 32767) || false) {
                        v1 = append(v1, "\n\t")
                        v167 = (v167 * (v167 & v167))
                        v167 = (v167 * (v167 - (v167 / ((v167) | 1))))
                        v172 := GenS4{F0: 2}
                        _ = v172
                    } else {
                        var v173 uint32 = 475
                        _ = v173
                    }
                }
                _ = v170
            }
            _ = fi165
            _ = v166
            _ = v167
        }
        _ = fi164
    }
    switch "a" {
    case "hello":
        var v174 int16 = (v2 & 290)
        _ = v174
    case "a":
        _ = len(v1)
    default:
        if len(v1) > 0 {
            v1[0] = ("\n\t" + "\n\t")
        }
    }
    if len(v1) > 0 {
        v1[0] = ("hello" + "hello")
    }
    if (v3 <= v3) {
        if ((^(v4)) >= v4) {
            v175 := GenS16{F0: (v3 & v3)}
            switch "a" {
            case "hello":
                for fi176 := 0; fi176 < 8; fi176++ {
                    for fi177 := 0; fi177 < 11; fi177++ {
                        switch ("\n\t" + ("a" + "hello")) {
                        case "":
                            v1 = append(v1, "hello")
                        case "a":
                            if true {
                                v4 += v4
                                switch ("a" + "hello") {
                                case "\n\t":
                                    v5 -= 42
                                default:
                                    v3 -= (v3 % ((v3) | 1))
                                }
                                v178 := GenS10{F0: 0}
                                _ = v178
                            } else {
                                fmt.Println(v2, v5, v2)
                                v3 -= (v3 << ((4294967295) & 63))
                            }
                        case "\n\t":
                            switch "hello" {
                            case "a":
                                v5 = v5
                            default:
                                if len(v1) > 0 {
                                    v1[0] = "foo bar"
                                }
                            }
                        default:
                            v179 := GenS6{F0: -562, F1: 42, F2: 2147483648}
                            _ = v179
                        }
                        _ = fi177
                    }
                    var v180 int64 = -1
                    var v181 uint32 = 340
                    _ = fi176
                    _ = v180
                    _ = v181
                }
            case "foo bar":
                if false {
                    var v182 uint = 11
                    if len(v1) > 0 {
                        v1[0] = "a"
                    }
                    _ = v182
                } else {
                    v183 := GenS17{F0: 9223372036854775807}
                    switch "a" {
                    case "hello":
                        v184 := GenS18{F0: 880}
                        _ = v184
                    case "foo bar":
                        if (float32(-2.25) < 1e10) {
                            for fi185 := 0; fi185 < 16; fi185++ {
                                for fi186 := 0; fi186 < 10; fi186++ {
                                    v187 := GenS19{F0: v2, F1: 762, F2: v3}
                                    var v188 string = ("hello" + "a")
                                    _ = fi186
                                    _ = v187
                                    _ = v188
                                }
                                _ = fi185
                            }
                        } else {
                            var v189 uint8 = 127
                            var v190 uint16 = 1
                            _ = v189
                            _ = v190
                        }
                    default:
                        v4 += (v4 * 2)
                    }
                    _ = v183
                }
            case "":
                var v191 string = (("\n\t" + "hello") + "a")
                _ = v191
            default:
                v2 -= -1
            }
            _ = v175
        }
        var v192 int8 = v3
        _ = v192
    } else {
        for fi193 := 0; fi193 < 14; fi193++ {
            v194 := GenS14{F0: 111, F1: 0, F2: 42}
            fmt.Println(v4, v3)
            for fi195 := 0; fi195 < 14; fi195++ {
                var v196 int32 = v5
                _ = fi195
                _ = v196
            }
            _ = fi193
            _ = v194
        }
        for fi197 := 0; fi197 < 11; fi197++ {
            if false {
                switch fi197 {
                case 9223372036854775807:
                    switch -9223372036854775808 {
                    case 32768:
                        for fi198 := 0; fi198 < 3; fi198++ {
                            v2 += v2
                            _ = fi198
                        }
                    case 9223372036854775807:
                        switch fi197 {
                        case 9223372036854775807:
                            for fi199 := 0; fi199 < 10; fi199++ {
                                var v200 uint64 = 4294967295
                                if (!((v4 % ((128) | 1)) != v4)) {
                                    var v201 uint8 = v4
                                    _ = v201
                                }
                                var v202 string = "hello"
                                v3 = (v3 * 42)
                                _ = fi199
                                _ = v200
                                _ = v202
                            }
                        case -628:
                            if (0.0 < float32((1e10 + (1.5 * 0.1)))) {
                                var v203 int = 179
                                var v204 uint32 = 694
                                _ = v203
                                _ = v204
                            }
                        default:
                            for fi205 := 0; fi205 < 16; fi205++ {
                                fmt.Println(v2, v2)
                                if ("hello" > "hello") {
                                    if (int8(fi205) != (-(v3))) {
                                        var v206 bool = (false && (v5 == (-(v5))))
                                        var v207 int16 = v2
                                        var v208 uint8 = v4
                                        switch "foo bar" {
                                        case "foo bar":
                                            switch fi205 {
                                            case -1:
                                                var v209 bool = true
                                                _ = v209
                                            case 2147483647:
                                                v2 = (v207 ^ 1000)
                                            default:
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                v206 = false
                                            }
                                        default:
                                            v2 += (v2 & v2)
                                        }
                                        fmt.Println(v4)
                                        _ = v206
                                        _ = v207
                                        _ = v208
                                    }
                                    var v210 int64 = -925
                                    _ = v210
                                } else {
                                    if (42 == 1) {
                                        fmt.Println(fi197, fi197, v4)
                                        v1 = append(v1, "\n\t")
                                        if (false && true) {
                                            var v211 string = "a"
                                            _ = len(v1)
                                            _ = v211
                                        }
                                        v4 = 207
                                    } else {
                                        v4 -= 2
                                        v5 += -2147483648
                                        v212 := GenS9{F0: true}
                                        _ = v212
                                    }
                                    var v213 uint = 18446744073709551615
                                    _ = v213
                                }
                                _ = fi205
                            }
                        }
                    default:
                        fmt.Println(v4)
                    }
                case 316:
                    if ((v4 % ((2) | 1)) < v4) {
                        fmt.Println(v2, fi197, v2)
                    } else {
                        v5 -= v5
                        for fi214 := 0; fi214 < 16; fi214++ {
                            v2 = (v2 >> ((42) & 63))
                            switch ("\n\t" + ("" + "a")) {
                            case "\n\t":
                                switch (fi214 << ((657) & 63)) {
                                case 9223372036854775807:
                                    var v215 int8 = v3
                                    _ = v215
                                case -2147483648:
                                    switch "hello" {
                                    case "a":
                                        switch "" {
                                        case "a":
                                            fmt.Println(fi197, v5, v2)
                                        case "hello":
                                            v5 = -837
                                        case "\n\t":
                                            fmt.Println(v5, v5, fi197)
                                        }
                                    default:
                                        if len(v1) > 0 {
                                            v1[0] = (("\n\t" + ("foo bar" + "a")) + "a")
                                        }
                                    }
                                }
                            case "":
                                if true {
                                    if true {
                                        v4 -= (v4 >> ((885) & 63))
                                        var v216 uint16 = 255
                                        v217 := GenS19{F0: (v2 & 867), F1: 1, F2: (^(v3))}
                                        _ = v216
                                        _ = v217
                                    }
                                    v5 = v5
                                    var v218 int32 = v5
                                    _ = v218
                                }
                                fmt.Println(fi197, fi197, fi197)
                                fmt.Println(fi197, fi197, fi214)
                            default:
                                var v219 uint64 = 4294967295
                                _ = v219
                            }
                            _ = fi214
                        }
                    }
                default:
                    switch "" {
                    case "\n\t":
                        v2 = v2
                    case "a":
                        v4 = v4
                    default:
                        _ = len(v1)
                    }
                }
            } else {
                var v220 string = (("\n\t" + "a") + "\n\t")
                _ = v220
            }
            for fi221 := 0; fi221 < 8; fi221++ {
                v222 := GenS20{F0: 42, F1: (v2 & v2), F2: v3}
                v2 = -1
                _ = fi221
                _ = v222
            }
            var v223 int64 = -9223372036854775808
            _ = fi197
            _ = v223
        }
        var v224 int64 = -9223372036854775808
        _ = v224
    }
    if true {
        if (false && (!(!((v2 * (v2 ^ v2)) != 42)))) {
            v4 = (v4 * 2)
            _ = len(v1)
        } else {
            var v225 int8 = v3
            if true {
                var v226 int32 = v5
                if len(v1) > 0 {
                    v1[0] = "hello"
                }
                for fi227 := 0; fi227 < 4; fi227++ {
                    switch fi227 {
                    case -9223372036854775808:
                        v228 := GenS1{F0: int16(v225)}
                        _ = v228
                    case 32768:
                        for fi229 := 0; fi229 < 6; fi229++ {
                            v2 = 1
                            if true {
                                if true {
                                    if (true || false) {
                                        v2 = -622
                                    }
                                }
                            } else {
                                var v230 int64 = -9223372036854775808
                                var v231 bool = true
                                v232 := GenS21{F0: int16(v5)}
                                _ = v230
                                _ = v231
                                _ = v232
                            }
                            v1 = append(v1, "a")
                            _ = fi229
                        }
                    default:
                        v225 += v225
                    }
                    if true {
                        var v233 uint32 = 313
                        fmt.Println(fi227, v5, v3)
                        switch "" {
                        case "foo bar":
                            v4 = 77
                        case "\n\t":
                            v226 = -2147483648
                        }
                        _ = v233
                    } else {
                        v226 -= v226
                        if true {
                            v225 = v3
                        } else {
                            fmt.Println(v225, v226)
                            v226 = v5
                            var v234 int8 = 100
                            _ = v234
                        }
                    }
                    v5 += 1
                    _ = fi227
                }
                _ = v226
            }
            _ = v225
        }
    }
    var v235 float64 = (float64(0.1) - -2.25)
    v1 = append(v1, "hello")
    v5 -= (v5 - (v5 | (-(v5))))
    v2 += (v2 | (v2 * (-(v2))))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v235:", v235)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v163
    _ = v2
    _ = v235
    _ = v3
    _ = v4
    _ = v5
}
