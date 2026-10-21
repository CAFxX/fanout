package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int64; F1 int64; F2 uint16 }
type GenS5 struct { F0 uint16; F1 uint }
type GenS6 struct { F0 int; F1 string }
type GenS7 struct { F0 uint8; F1 int64 }
type GenS8 struct { F0 bool; F1 int32 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 uint8; F1 uint32; F2 int }
type GenS11 struct { F0 int; F1 string }
type GenS12 struct { F0 string; F1 uint; F2 uint16 }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 uint64 }
func genFn0(p0 uint32) int32 {
    for fi6 := 0; fi6 < 19; fi6++ {
        v7 := []int{(fi6 / (((fi6 * 128)) | 1)), -9223372036854775808, 1}
        _ = fi6
        _ = v7
    }
    p0 = (p0 % ((p0) | 1))
    _ = p0
    return 42
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]uint32)
    v3 := []int{}
    v4 := make(map[string]bool)
    v5 := make(map[int]uint32)
    var v8 uint16 = 0
    for fi9 := 0; fi9 < 6; fi9++ {
        v8 += v8
        var v10 int64 = -1
        for fi11 := 0; fi11 < 2; fi11++ {
            fmt.Println(fi9, fi9, v10)
            if (!true) {
                if (true && (true || true)) {
                    if (!true) {
                        v10 += -9223372036854775808
                    }
                    v10 += (v10 & v10)
                    v2[0] = 4294967295
                }
                switch "" {
                case "\n\t":
                    var v12 uint16 = (v8 - (v8 >> ((2) & 63)))
                    _ = v12
                default:
                    switch (fi9 ^ int(v10)) {
                    case 1:
                        var v13 string = "foo bar"
                        _ = v13
                    }
                }
                v3 = append(v3, -1)
            }
            _ = fi11
        }
        _ = fi9
        _ = v10
    }
    if true {
        v8 += v8
        switch ("hello" + "\n\t") {
        case "hello":
            for fi14 := 0; fi14 < 13; fi14++ {
                v15 := GenS1{F0: "\n\t", F1: 311, F2: 227}
                var v16 int8 = -86
                for fi17 := 0; fi17 < 19; fi17++ {
                    if true {
                        var v18 uint64 = 846
                        if false {
                            var v19 int16 = 574
                            fmt.Println(fi17, fi17, v18)
                            fmt.Println(v18)
                            _ = v19
                        }
                        _ = v18
                    } else {
                        var v20 bool = (fi17 <= fi14)
                        if (!(!v20)) {
                            v8 = (v8 * v8)
                            var v21 float64 = ((1e10 * -2.25) + 0.1)
                            v22 := v1[(fi17 % ((fi17) | 1))]
                            _ = v21
                            _ = v22
                        }
                        _ = v20
                    }
                    fmt.Println(fi14, fi14, fi17)
                    _ = fi17
                }
                _ = fi14
                _ = v15
                _ = v16
            }
        default:
            for fi23 := 0; fi23 < 9; fi23++ {
                if (0.1 > 3.14159) {
                    v8 = (^(v8))
                    var v24 uint32 = 2147483647
                    fmt.Println(v24, v8, v24)
                    _ = v24
                } else {
                    v25 := GenS1{F0: ("" + "hello"), F1: -121, F2: 255}
                    v26 := v2[2147483648]
                    var v27 uint32 = (v26 << ((270) & 63))
                    _ = v25
                    _ = v26
                    _ = v27
                }
                for fi28 := 0; fi28 < 4; fi28++ {
                    v29 := v2[2]
                    v8 = (^(v8))
                    _ = fi28
                    _ = v29
                }
                if (!(false && (193 < 252))) {
                    v8 = v8
                } else {
                    switch (fi23 % ((fi23) | 1)) {
                    case -1:
                        var v30 int16 = 179
                        _ = v30
                    case -466:
                        v3 = append(v3, (-(fi23)))
                    }
                    v8 -= v8
                    switch 2147483647 {
                    case -225:
                        switch "hello" {
                        case "\n\t":
                            var v31 uint64 = 2
                            _ = v31
                        default:
                            if false {
                                fmt.Println(v8, v8)
                                fmt.Println(fi23, v8)
                            } else {
                                var v32 bool = ((-2.25 * ((1e10 + ((3.14159 + 1.5) / 0.1)) / 1e10)) < 0.1)
                                _ = v32
                            }
                        }
                    case -507:
                        fmt.Println(fi23)
                    default:
                        v8 = v8
                    }
                }
                _ = fi23
            }
        }
    } else {
        if (42 < 0) {
            switch "" {
            case "\n\t":
                if true {
                    _ = len(v3)
                } else {
                    var v33 int8 = 42
                    _ = v33
                }
            }
        }
        if false {
            var v34 string = ("foo bar" + "\n\t")
            for fi35 := 0; fi35 < 8; fi35++ {
                var v36 uint16 = (v8 - 42)
                var v37 float32 = 0.0
                v8 = v36
                _ = fi35
                _ = v36
                _ = v37
            }
            if false {
                v38 := GenS2{F0: 73}
                var v39 float64 = (float64(((-2.25 + 0.0) - 1.5)) / -2.25)
                if false {
                    var v40 uint16 = (v8 / ((v8) | 1))
                    switch 127 {
                    case -311:
                        var v41 uint8 = 42
                        _ = v41
                    }
                    _ = v40
                }
                _ = v38
                _ = v39
            }
            _ = v34
        } else {
            v8 += v8
            for fi42 := 0; fi42 < 16; fi42++ {
                var v43 float32 = 1e10
                var v44 uint8 = 128
                v43 = ((1.5 + float32(0.1)) * v43)
                _ = fi42
                _ = v43
                _ = v44
            }
            v8 -= (v8 | 2)
        }
        var v45 int32 = -771
        _ = v45
    }
    switch "" {
    case "":
        v8 -= (^(v8))
    case "\n\t":
        switch -359 {
        case 42:
            if false {
                if (!true) {
                    if (-1 < 0) {
                        fmt.Println(v8)
                        var v46 int8 = 0
                        _ = v46
                    } else {
                        v8 += v8
                        v8 -= v8
                    }
                }
            }
        case -832:
            switch 678 {
            case 52:
                v47 := GenS3{F0: 182}
                _ = v47
            case -1:
                var v48 uint32 = 433
                _ = v48
            }
        case -476:
            switch "foo bar" {
            case "\n\t":
                v8 -= (v8 & v8)
            }
        }
    }
    for fi49 := 0; fi49 < 4; fi49++ {
        for fi50 := 0; fi50 < 13; fi50++ {
            var v51 int16 = -464
            _ = fi50
            _ = v51
        }
        v8 = v8
        v2[0] = 2147483647
        _ = fi49
    }
    v8 = (^(v8))
    switch "" {
    case "a":
        v52 := GenS2{F0: -125}
        _ = v52
    case "hello":
        if (42 > 18446744073709551615) {
            if true {
                for fi53 := 0; fi53 < 2; fi53++ {
                    v8 = v8
                    v54 := GenS2{F0: -87}
                    _ = fi53
                    _ = v54
                }
                v55 := GenS4{F0: 9223372036854775807, F1: 9223372036854775807, F2: v8}
                if (1 <= 51) {
                    for fi56 := 0; fi56 < 20; fi56++ {
                        fmt.Println(v8, v8)
                        _ = fi56
                    }
                    for fi57 := 0; fi57 < 7; fi57++ {
                        for fi58 := 0; fi58 < 9; fi58++ {
                            var v59 int16 = 32767
                            _ = fi58
                            _ = v59
                        }
                        for fi60 := 0; fi60 < 4; fi60++ {
                            var v61 uint32 = 866
                            var v62 uint8 = 245
                            switch "hello" {
                            case "\n\t":
                                switch (fi60 << ((0) & 63)) {
                                case 312:
                                    fmt.Println(v8)
                                case 128:
                                    if true {
                                        fmt.Println(fi57)
                                    }
                                default:
                                    var v63 int16 = -1
                                    _ = v63
                                }
                            case "foo bar":
                                v8 = 42
                            }
                            v62 = (^(v62))
                            v64 := GenS5{F0: (v8 - (v8 << ((406) & 63))), F1: 18446744073709551615}
                            _ = fi60
                            _ = v61
                            _ = v62
                            _ = v64
                        }
                        _ = fi57
                    }
                    for fi65 := 0; fi65 < 7; fi65++ {
                        v8 = v8
                        v8 -= v8
                        _ = fi65
                    }
                }
                _ = v55
            } else {
                if (false || (("foo bar" + ("\n\t" + ("" + ""))) > "\n\t")) {
                    v66 := GenS6{F0: -9223372036854775808, F1: "a"}
                    var v67 uint16 = 546
                    _ = v66
                    _ = v67
                } else {
                    var v68 int64 = -36
                    for fi69 := 0; fi69 < 8; fi69++ {
                        v70 := GenS6{F0: fi69, F1: "foo bar"}
                        var v71 uint8 = 115
                        _ = fi69
                        _ = v70
                        _ = v71
                    }
                    v8 -= v8
                    _ = v68
                }
                var v72 uint64 = 42
                for fi73 := 0; fi73 < 8; fi73++ {
                    for fi74 := 0; fi74 < 13; fi74++ {
                        v75 := GenS1{F0: "hello", F1: 0, F2: 1}
                        v3 = append(v3, (fi73 | (fi74 ^ fi74)))
                        v76 := GenS4{F0: 9223372036854775807, F1: -1, F2: (^(v8))}
                        _ = fi74
                        _ = v75
                        _ = v76
                    }
                    if ("" != ("" + ("hello" + "\n\t"))) {
                        for fi77 := 0; fi77 < 11; fi77++ {
                            v8 += 1
                            if len(v3) > 0 {
                                v3[0] = (-(fi73))
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            var v78 uint32 = 322
                            switch "a" {
                            case "foo bar":
                                fmt.Println(fi77, v78)
                            case "a":
                                v8 += (v8 * 61)
                            }
                            _ = fi77
                            _ = v78
                        }
                        v79 := v2[33]
                        _ = v79
                    }
                    for fi80 := 0; fi80 < 15; fi80++ {
                        var v81 int8 = -1
                        v8 += (v8 & v8)
                        v82 := GenS5{F0: 65535, F1: 255}
                        v8 -= (v8 | 410)
                        var v83 int32 = 2147483647
                        _ = fi80
                        _ = v81
                        _ = v82
                        _ = v83
                    }
                    _ = fi73
                }
                _ = v72
            }
            fmt.Println(v8)
            for fi84 := 0; fi84 < 17; fi84++ {
                v85 := GenS7{F0: 2, F1: -1}
                v8 += (v8 + (v8 & 749))
                _ = fi84
                _ = v85
            }
        }
    default:
        for fi86 := 0; fi86 < 12; fi86++ {
            if len(v3) > 0 {
                v3[0] = fi86
            }
            v87 := GenS8{F0: (!(2 < 42)), F1: -104}
            _ = fi86
            _ = v87
        }
    }
    v8 -= v8
    if ((genFn0(uint32(v8)) == -1) || false) {
        if false {
            if true {
                v88 := GenS8{F0: (421 >= uint(v8)), F1: 1}
                switch (("\n\t" + "\n\t") + "") {
                case "foo bar":
                    for fi89 := 0; fi89 < 9; fi89++ {
                        v8 = (v8 * (v8 * (^(v8))))
                        _ = fi89
                    }
                case "\n\t":
                    v90 := GenS1{F0: "\n\t", F1: -1, F2: 127}
                    _ = v90
                case "":
                    var v91 int8 = 0
                    _ = v91
                }
                _ = v88
            }
        } else {
            v8 -= (v8 - (^(v8)))
            if (!(1 >= -757)) {
                var v92 int32 = -496
                _ = v92
            } else {
                fmt.Println(v8)
                fmt.Println(v8)
            }
            for fi93 := 0; fi93 < 19; fi93++ {
                v94 := GenS9{F0: 113}
                var v95 int = (fi93 / (((fi93 + 2147483647)) | 1))
                _ = fi93
                _ = v94
                _ = v95
            }
        }
        if (2147483647 == 42) {
            v8 += v8
            v1[2147483647] = "\n\t"
            var v96 int = 9223372036854775807
            _ = v96
        }
    } else {
        var v97 uint8 = 255
        var v98 uint16 = (v8 & (v8 - (^(v8))))
        if false {
            var v99 uint32 = 4294967295
            if true {
                fmt.Println(v99, v8)
            }
            if true {
                if ("hello" >= ("" + "hello")) {
                    var v100 int8 = 7
                    v8 = (v8 - uint16(v97))
                    _ = v100
                } else {
                    for fi101 := 0; fi101 < 20; fi101++ {
                        var v102 bool = false
                        for fi103 := 0; fi103 < 19; fi103++ {
                            var v104 uint8 = (v97 * (v97 << ((42) & 63)))
                            switch -405 {
                            case 2147483647:
                                if len(v3) > 0 {
                                    v3[0] = (fi101 >> ((uint(v8)) & 63))
                                }
                            case 9223372036854775807:
                                delete(v2, (v99 % (((^(v99))) | 1)))
                            case 32768:
                                fmt.Println(fi103, v97)
                            }
                            var v105 uint16 = v8
                            switch fi103 {
                            case 1:
                                if v102 {
                                    v106 := GenS10{F0: v97, F1: 4294967295, F2: fi103}
                                    _ = v106
                                } else {
                                    v107 := GenS4{F0: 0, F1: -471, F2: 2}
                                    var v108 int32 = 0
                                    v104 = (v97 + 255)
                                    _ = v107
                                    _ = v108
                                }
                            default:
                                fmt.Println(v98, v104)
                                if true {
                                    v104 -= (v97 | (v97 + v104))
                                } else {
                                    fmt.Println(v105, fi103)
                                }
                            }
                            var v109 string = "a"
                            _ = fi103
                            _ = v104
                            _ = v105
                            _ = v109
                        }
                        _ = fi101
                        _ = v102
                    }
                    switch 1 {
                    case -200:
                        for fi110 := 0; fi110 < 4; fi110++ {
                            v111 := GenS7{F0: (v97 << ((0) & 63)), F1: 9223372036854775807}
                            _ = fi110
                            _ = v111
                        }
                    default:
                        for fi112 := 0; fi112 < 13; fi112++ {
                            v113 := GenS11{F0: -1, F1: "\n\t"}
                            v2[2] = 70
                            switch (-(fi112)) {
                            case -1:
                                if true {
                                    v114 := GenS5{F0: v8, F1: 2}
                                    _ = v114
                                }
                            case 32768:
                                v8 += (^(v8))
                            default:
                                var v115 uint16 = (v8 << ((1) & 63))
                                _ = v115
                            }
                            _ = fi112
                            _ = v113
                        }
                    }
                    fmt.Println(v8, v97, v99)
                }
            } else {
                v116 := GenS12{F0: ("" + "a"), F1: 2, F2: 2}
                v117 := GenS1{F0: (("" + "") + "foo bar"), F1: 636, F2: 0}
                var v118 uint32 = v99
                _ = v116
                _ = v117
                _ = v118
            }
            _ = v99
        }
        _ = v97
        _ = v98
    }
    if false {
        if true {
            fmt.Println(v8)
            v119 := GenS13{F0: 18446744073709551615}
            switch 128 {
            case 42:
                for fi120 := 0; fi120 < 1; fi120++ {
                    v121 := GenS9{F0: 4294967295}
                    _ = fi120
                    _ = v121
                }
            }
            _ = v119
        }
        var v122 bool = false
        _ = v122
    }
    v4[""] = (false || ((((false && (691 != 184)) && true) || ("foo bar" <= "hello")) || false))
    var v123 float64 = (0.0 - (0.1 * 0.1))
    var v124 int = 2147483647
    v125 := v2[65535]
    for fi126 := 0; fi126 < 6; fi126++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        if ((4294967295 != 589) && false) {
            for fi127 := 0; fi127 < 16; fi127++ {
                if false {
                    if (0.0 != -2.25) {
                        v128 := GenS9{F0: 0}
                        _ = v128
                    }
                }
                _ = fi127
            }
        } else {
            if false {
                v8 -= v8
                v129 := GenS14{F0: 359}
                switch v124 {
                case -9223372036854775808:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                _ = v129
            }
            fmt.Println(v125, v125)
        }
        _ = fi126
    }
    v125 = v125
    var v130 uint8 = 233
    var v131 float32 = float32(float64(1e10))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v123:", v123)
    fmt.Println("v124:", v124)
    fmt.Println("v125:", v125)
    fmt.Println("v130:", v130)
    fmt.Println("v131:", v131)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v123
    _ = v124
    _ = v125
    _ = v130
    _ = v131
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
