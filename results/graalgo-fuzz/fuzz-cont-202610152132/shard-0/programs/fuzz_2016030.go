package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 uint8; F1 uint8; F2 int }
type GenS4 struct { F0 uint32; F1 uint64 }
type GenS5 struct { F0 uint; F1 uint }
type GenS6 struct { F0 int64; F1 uint16 }
type GenS7 struct { F0 int; F1 uint8 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int32; F1 uint8 }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 uint; F1 uint32 }
type GenS14 struct { F0 uint; F1 int16 }
type GenS15 struct { F0 int16; F1 int; F2 int32 }
type GenS16 struct { F0 uint16; F1 int64 }
type GenS17 struct { F0 int; F1 int64 }
type GenS18 struct { F0 int32; F1 int32 }
type GenS19 struct { F0 int32; F1 uint; F2 bool }
type GenS20 struct { F0 uint8 }
type GenS21 struct { F0 uint32 }
type GenS22 struct { F0 int64; F1 uint8; F2 uint32 }
type GenS23 struct { F0 uint32 }

func main() {
    v1 := make(map[string]bool)
    var v2 int32 = -2147483648
    v3 := []uint32{640}
    var v4 string = ("a" + ("a" + "hello"))
    v5 := make(map[int]uint32)
    v6 := make(map[string]bool)
    var v7 uint8 = 1
    if (-9223372036854775808 != 824) {
        var v8 string = ""
        var v9 float64 = 1.5
        _ = v8
        _ = v9
    }
    if (42 >= 42) {
        for fi10 := 0; fi10 < 11; fi10++ {
            var v11 int64 = -9223372036854775808
            _ = fi10
            _ = v11
        }
    } else {
        for fi12 := 0; fi12 < 18; fi12++ {
            var v13 int64 = -805
            if (false || false) {
                for fi14 := 0; fi14 < 20; fi14++ {
                    v15 := GenS1{F0: true, F1: 18446744073709551615}
                    _ = fi14
                    _ = v15
                }
                v2 = (v2 >> ((65535) & 63))
                var v16 uint32 = 610
                _ = v16
            }
            _ = fi12
            _ = v13
        }
        v2 = 0
    }
    if (!(((v7 | (^(v7))) >= v7) && ((0.0 > 1e10) && (!true)))) {
        fmt.Println(v4)
        v4 = v4
    } else {
        if true {
            v17 := GenS2{F0: 0, F1: v4}
            var v18 int = 863
            _ = v17
            _ = v18
        } else {
            v3 = append(v3, 537)
            if true {
                v19 := GenS3{F0: (v7 + v7), F1: v7, F2: 0}
                _ = v19
            }
        }
    }
    for fi20 := 0; fi20 < 10; fi20++ {
        for fi21 := 0; fi21 < 17; fi21++ {
            v22 := GenS4{F0: 32, F1: 2}
            if true {
                v23 := GenS3{F0: 128, F1: v7, F2: fi20}
                var v24 uint16 = 264
                _ = v23
                _ = v24
            } else {
                _ = len(v3)
                v7 = v7
                var v25 uint8 = (^(v7))
                _ = v25
            }
            _ = fi21
            _ = v22
        }
        v26 := GenS4{F0: 779, F1: 980}
        var v27 uint32 = 462
        _ = fi20
        _ = v26
        _ = v27
    }
    var v28 uint16 = 32767
    v7 -= 0
    if false {
        fmt.Println(v4)
        for fi29 := 0; fi29 < 9; fi29++ {
            v6[v4] = (655 < 18446744073709551615)
            _ = fi29
        }
        if ((0.1 - 1.5) <= 0.0) {
            var v30 int64 = 9223372036854775807
            _ = v30
        } else {
            switch v4 {
            case "":
                var v31 uint64 = 4294967295
                _ = v31
            case "foo bar":
                fmt.Println(v28)
            }
            if true {
                var v32 int16 = 1
                if len(v3) > 0 {
                    v3[0] = 545
                }
                _ = v32
            } else {
                var v33 uint16 = v28
                _ = v33
            }
            v28 = v28
        }
    }
    v34 := GenS2{F0: 2147483647, F1: v4}
    if false {
        v35 := GenS4{F0: 4294967295, F1: 18446744073709551615}
        _ = v35
    } else {
        v28 += 32767
    }
    if len(v3) > 0 {
        v3[0] = 2147483648
    }
    if (true || (431 > 188)) {
        v36 := GenS5{F0: 1, F1: 42}
        switch -1 {
        case 372:
            if true {
                if (!false) {
                    if true {
                        var v37 uint = 255
                        v38 := GenS6{F0: -1, F1: (v28 | (^(v28)))}
                        var v39 string = v4
                        _ = v37
                        _ = v38
                        _ = v39
                    }
                }
                var v40 uint32 = 1
                _ = v40
            } else {
                for fi41 := 0; fi41 < 16; fi41++ {
                    if true {
                        for fi42 := 0; fi42 < 18; fi42++ {
                            v3 = append(v3, 4294967295)
                            fmt.Println(v4, fi41)
                            switch "hello" {
                            case "":
                                v28 += (v28 + 841)
                            default:
                                v43 := GenS6{F0: -1, F1: (v28 - 230)}
                                _ = v43
                            }
                            _ = fi42
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    if true {
                        v44 := GenS2{F0: 42, F1: "\n\t"}
                        if true {
                            var v45 int64 = 183
                            switch (("a" + (v4 + "")) + "\n\t") {
                            case "a":
                                var v46 int64 = (^(v45))
                                _ = v46
                            case "":
                                if false {
                                    v4 = v4
                                    fmt.Println(v4, v28, fi41)
                                    var v47 bool = ((0 < -1) || false)
                                    _ = v47
                                } else {
                                    v48 := GenS7{F0: fi41, F1: (v7 ^ (v7 | v7))}
                                    var v49 int8 = 127
                                    var v50 float64 = 0.0
                                    _ = v48
                                    _ = v49
                                    _ = v50
                                }
                            default:
                                var v51 int64 = v45
                                _ = v51
                            }
                            _ = v45
                        }
                        if (1 > 2147483648) {
                            if (int16(v7) == 0) {
                                if ((v4 >= v4) && (("a" + (v4 + v4)) <= ((("" + (("" + v4) + "a")) + v4) + "hello"))) {
                                    if true {
                                        switch v4 {
                                        case "foo bar":
                                            v7 -= v7
                                        case "":
                                            for fi52 := 0; fi52 < 4; fi52++ {
                                                var v53 string = v4
                                                v53 = ""
                                                var v54 uint32 = 2147483647
                                                _ = fi52
                                                _ = v53
                                                _ = v54
                                            }
                                        }
                                        for fi55 := 0; fi55 < 13; fi55++ {
                                            var v56 uint32 = 737
                                            if true {
                                                v57 := GenS8{F0: 2}
                                                v58 := GenS9{F0: "foo bar"}
                                                _ = v57
                                                _ = v58
                                            } else {
                                                var v59 uint16 = (v28 / ((v28) | 1))
                                                _ = v59
                                            }
                                            v60 := GenS10{F0: v2, F1: v7}
                                            _ = fi55
                                            _ = v56
                                            _ = v60
                                        }
                                        for fi61 := 0; fi61 < 6; fi61++ {
                                            var v62 uint = 1
                                            var v63 uint16 = (^(v28))
                                            v7 -= v7
                                            _ = len(v3)
                                            var v64 bool = (!true)
                                            _ = fi61
                                            _ = v62
                                            _ = v63
                                            _ = v64
                                        }
                                    } else {
                                        switch -2147483648 {
                                        case 0:
                                            var v65 float32 = 3.14159
                                            _ = v65
                                        default:
                                            if (-9223372036854775808 != 0) {
                                                if false {
                                                    if true {
                                                        var v66 string = v4
                                                        _ = v66
                                                    }
                                                    switch v4 {
                                                    case "":
                                                        fmt.Println(v2)
                                                    case "a":
                                                        if false {
                                                            if (false && true) {
                                                                for fi67 := 0; fi67 < 18; fi67++ {
                                                                    if false {
                                                                        var v68 int = (fi41 / (((fi41 ^ -2147483648)) | 1))
                                                                        var v69 uint8 = (v7 ^ v7)
                                                                        v68 = fi41
                                                                        var v70 string = v4
                                                                        _ = v68
                                                                        _ = v69
                                                                        _ = v70
                                                                    } else {
                                                                        v71 := GenS11{F0: 65535}
                                                                        _ = v71
                                                                    }
                                                                    _ = len(v1)
                                                                    if false {
                                                                        if ((v28 >= 65535) && (false && true)) {
                                                                            if (-266 < 1) {
                                                                                var v72 bool = true
                                                                                var v73 int16 = -848
                                                                                switch fi67 {
                                                                                case -9223372036854775808:
                                                                                    var v74 int64 = -40
                                                                                    _ = v74
                                                                                default:
                                                                                    v7 -= (v7 >> ((875) & 63))
                                                                                }
                                                                                v7 = 128
                                                                                v4 = "a"
                                                                                var v75 uint16 = v28
                                                                                v72 = v72
                                                                                _ = len(v3)
                                                                                _ = v72
                                                                                _ = v73
                                                                                _ = v75
                                                                            } else {
                                                                                if (false || false) {
                                                                                    var v76 bool = false
                                                                                    var v77 int8 = 1
                                                                                    var v78 uint32 = 0
                                                                                    _ = v76
                                                                                    _ = v77
                                                                                    _ = v78
                                                                                } else {
                                                                                    switch 86 {
                                                                                    case -535:
                                                                                        switch v4 {
                                                                                        case "a":
                                                                                            var v79 uint32 = 1
                                                                                            _ = v79
                                                                                        default:
                                                                                            v3 = append(v3, 42)
                                                                                        }
                                                                                    case -1:
                                                                                        switch fi41 {
                                                                                        case 42:
                                                                                            v5[fi41] = 904
                                                                                        case 32768:
                                                                                            v80 := GenS7{F0: fi67, F1: v7}
                                                                                            _ = v80
                                                                                        }
                                                                                    }
                                                                                }
                                                                                var v81 float64 = (3.14159 - 1.5)
                                                                                fmt.Println(v2, fi67, v2)
                                                                                _ = v81
                                                                            }
                                                                        }
                                                                    }
                                                                    v7 = v7
                                                                    _ = fi67
                                                                }
                                                            }
                                                            fmt.Println(v7, v4, v4)
                                                            fmt.Println(v7)
                                                        }
                                                    }
                                                    fmt.Println(v4)
                                                }
                                                if true {
                                                    if true {
                                                        v7 += 128
                                                        if true {
                                                            for fi82 := 0; fi82 < 11; fi82++ {
                                                                var v83 float64 = 3.14159
                                                                _ = fi82
                                                                _ = v83
                                                            }
                                                            for fi84 := 0; fi84 < 15; fi84++ {
                                                                v7 = v7
                                                                v2 = 725
                                                                switch v4 {
                                                                case "hello":
                                                                    switch int(v2) {
                                                                    case -104:
                                                                        if (((3.14159 * 0.0) - (3.14159 - ((float32((float64((0.0 - 1.5)) * 0.0)) / 1.0) / float32(1.5)))) == ((float32(-2.25) - 0.1) * 1.5)) {
                                                                            v2 = 0
                                                                            _ = len(v6)
                                                                        } else {
                                                                            var v85 uint16 = v28
                                                                            if (((408 == 1) || (v7 <= (^(v7)))) || true) {
                                                                                if true {
                                                                                    v85 = v85
                                                                                    fmt.Println(fi41)
                                                                                } else {
                                                                                    var v86 uint32 = 1
                                                                                    fmt.Println(fi84, v2, v86)
                                                                                    _ = v86
                                                                                }
                                                                                v87 := v5[(^(fi41))]
                                                                                delete(v6, v4)
                                                                                if (!(18446744073709551615 >= 0)) {
                                                                                    v87 = (v87 % ((uint32(v28)) | 1))
                                                                                    v88 := GenS12{F0: 794}
                                                                                    _ = v88
                                                                                }
                                                                                _ = v87
                                                                            } else {
                                                                                switch (^(fi41)) {
                                                                                case 42:
                                                                                    if false {
                                                                                        var v89 bool = false
                                                                                        var v90 uint8 = 255
                                                                                        _ = v89
                                                                                        _ = v90
                                                                                    }
                                                                                }
                                                                            }
                                                                            _ = v85
                                                                        }
                                                                    case 9223372036854775807:
                                                                        delete(v6, "hello")
                                                                    case 128:
                                                                        var v91 int16 = -32768
                                                                        _ = v91
                                                                    }
                                                                default:
                                                                    var v92 uint32 = 2147483647
                                                                    _ = v92
                                                                }
                                                                _ = fi84
                                                            }
                                                            fmt.Println(v4)
                                                        } else {
                                                            var v93 bool = ((304 >= 4294967295) && false)
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            _ = v93
                                                        }
                                                    }
                                                    v94 := GenS2{F0: 555, F1: v4}
                                                    fmt.Println(v2, fi41, fi41)
                                                    _ = v94
                                                }
                                            } else {
                                                switch (v4 + (v4 + ((v4 + ((v4 + "") + v4)) + (v4 + v4)))) {
                                                case "hello":
                                                    fmt.Println(fi41)
                                                case "foo bar":
                                                    v2 = (v2 * -965)
                                                case "":
                                                    switch ((v4 + v4) + v4) {
                                                    case "\n\t":
                                                        v4 = ((v4 + v4) + v4)
                                                    case "hello":
                                                        v7 = (^(v7))
                                                    default:
                                                        if (((v28 - v28) != 32767) && false) {
                                                            if false {
                                                                v3 = append(v3, 1)
                                                            }
                                                        }
                                                    }
                                                }
                                                fmt.Println(v28, fi41)
                                                if true {
                                                    for fi95 := 0; fi95 < 18; fi95++ {
                                                        var v96 bool = ((!(((v4 + v4) <= (v4 + "foo bar")) || true)) && (true || true))
                                                        var v97 int16 = 1
                                                        _ = fi95
                                                        _ = v96
                                                        _ = v97
                                                    }
                                                    var v98 uint = 65535
                                                    _ = v98
                                                } else {
                                                    v99 := GenS1{F0: false, F1: 0}
                                                    _ = v99
                                                }
                                            }
                                        }
                                    }
                                    for fi100 := 0; fi100 < 1; fi100++ {
                                        fmt.Println(fi41, v28, v2)
                                        v4 = (v4 + (("hello" + (v4 + (((v4 + v4) + v4) + (v4 + v4)))) + ((v4 + "a") + v4)))
                                        _ = len(v6)
                                        _ = len(v6)
                                        var v101 int16 = -32768
                                        if true {
                                            var v102 float64 = 1.5
                                            var v103 int64 = 9223372036854775807
                                            _ = v102
                                            _ = v103
                                        } else {
                                            var v104 string = v4
                                            v105 := GenS4{F0: 2147483647, F1: 1}
                                            switch "\n\t" {
                                            case "":
                                                switch v4 {
                                                case "":
                                                    var v106 bool = ((-(fi41)) <= (-(fi41)))
                                                    _ = v106
                                                }
                                            case "hello":
                                                v28 += (v28 >> ((0) & 63))
                                            case "a":
                                                var v107 uint8 = (v7 ^ v7)
                                                _ = v107
                                            default:
                                                switch "hello" {
                                                case "a":
                                                    if (true && true) {
                                                        v3 = append(v3, 235)
                                                        v108 := GenS1{F0: false, F1: 42}
                                                        _ = v108
                                                    } else {
                                                        _ = len(v6)
                                                        if (!true) {
                                                            v109 := GenS12{F0: 662}
                                                            v101 += v101
                                                            _ = v109
                                                        }
                                                        var v110 float32 = 0.0
                                                        _ = v110
                                                    }
                                                case "\n\t":
                                                    if false {
                                                        if true {
                                                            switch "hello" {
                                                            case "a":
                                                                if false {
                                                                    v111 := GenS7{F0: fi41, F1: v7}
                                                                    _ = v111
                                                                } else {
                                                                    v112 := GenS11{F0: v28}
                                                                    switch fi100 {
                                                                    case -536:
                                                                        switch v104 {
                                                                        case "":
                                                                            v28 = v28
                                                                        case "foo bar":
                                                                            v113 := GenS13{F0: 65535, F1: 0}
                                                                            _ = v113
                                                                        }
                                                                    case -720:
                                                                        fmt.Println(v28)
                                                                    default:
                                                                        var v114 int8 = -1
                                                                        _ = v114
                                                                    }
                                                                    v104 = ((v104 + ("hello" + (v4 + v4))) + v104)
                                                                    _ = v112
                                                                }
                                                            case "foo bar":
                                                                fmt.Println(v4, fi41)
                                                            case "\n\t":
                                                                var v115 uint = 487
                                                                _ = v115
                                                            default:
                                                                switch 0 {
                                                                case 1:
                                                                    if (v7 > 1) {
                                                                        v7 = (v7 ^ v7)
                                                                        var v116 uint16 = (v28 / ((v28) | 1))
                                                                        var v117 int = fi100
                                                                        _ = v116
                                                                        _ = v117
                                                                    } else {
                                                                        v7 = (v7 - 153)
                                                                    }
                                                                    v118 := GenS14{F0: 208, F1: int16(v7)}
                                                                    _ = v118
                                                                case 4294967295:
                                                                    _ = len(v3)
                                                                }
                                                            }
                                                            fmt.Println(v7, v4, v28)
                                                            delete(v5, 4294967295)
                                                        }
                                                    } else {
                                                        fmt.Println(v7, v104)
                                                        switch -9223372036854775808 {
                                                        case 686:
                                                            v119 := GenS15{F0: (v101 ^ v101), F1: fi100, F2: int32(v7)}
                                                            _ = v119
                                                        case 32767:
                                                            v7 -= (v7 & 236)
                                                        case -9223372036854775808:
                                                            var v120 string = v104
                                                            _ = v120
                                                        }
                                                        v121 := GenS13{F0: 4294967295, F1: 2}
                                                        var v122 bool = false
                                                        if (-2.25 >= 0.0) {
                                                            v28 = 1
                                                        }
                                                        _ = v121
                                                        _ = v122
                                                    }
                                                }
                                            }
                                            _ = v104
                                            _ = v105
                                        }
                                        v28 = (v28 | v28)
                                        _ = fi100
                                        _ = v101
                                    }
                                    switch fi41 {
                                    case 128:
                                        switch "" {
                                        case "foo bar":
                                            for fi123 := 0; fi123 < 15; fi123++ {
                                                v3 = append(v3, 162)
                                                v124 := GenS12{F0: 4294967295}
                                                if (!(65535 != 2147483648)) {
                                                    var v125 float64 = 3.14159
                                                    var v126 int = 32767
                                                    var v127 int = v126
                                                    _ = v125
                                                    _ = v126
                                                    _ = v127
                                                }
                                                _ = fi123
                                                _ = v124
                                            }
                                        }
                                    case 32:
                                        delete(v6, "")
                                    default:
                                        switch v4 {
                                        case "\n\t":
                                            v7 = uint8(v28)
                                        case "hello":
                                            if false {
                                                v2 = 1
                                                v128 := GenS8{F0: 18446744073709551615}
                                                var v129 uint32 = 585
                                                _ = v128
                                                _ = v129
                                            }
                                        default:
                                            v130 := GenS6{F0: -548, F1: (v28 ^ (v28 ^ (^(v28))))}
                                            _ = v130
                                        }
                                    }
                                } else {
                                    _ = len(v3)
                                    switch v4 {
                                    case "hello":
                                        v131 := GenS16{F0: v28, F1: -9223372036854775808}
                                        _ = v131
                                    default:
                                        switch fi41 {
                                        case 0:
                                            var v132 int32 = (-(v2))
                                            _ = v132
                                        case 752:
                                            if (!(0 <= 42)) {
                                                if len(v3) > 0 {
                                                    v3[0] = 0
                                                }
                                            } else {
                                                var v133 uint = 1
                                                switch fi41 {
                                                case 585:
                                                    switch fi41 {
                                                    case 0:
                                                        var v134 uint16 = 2
                                                        _ = v134
                                                    default:
                                                        var v135 int8 = 0
                                                        _ = v135
                                                    }
                                                case -43:
                                                    v136 := GenS17{F0: (fi41 >> ((4294967295) & 63)), F1: -396}
                                                    _ = v136
                                                case 375:
                                                    v137 := GenS18{F0: v2, F1: v2}
                                                    _ = v137
                                                }
                                                _ = v133
                                            }
                                        default:
                                            switch v4 {
                                            case "a":
                                                for fi138 := 0; fi138 < 6; fi138++ {
                                                    v139 := GenS19{F0: v2, F1: 2, F2: false}
                                                    if false {
                                                        v5[(fi138 % ((-1) | 1))] = 1
                                                        v2 = v2
                                                        if (!true) {
                                                            var v140 uint8 = (v7 << ((0) & 63))
                                                            var v141 float64 = (3.14159 / 1.0)
                                                            _ = v140
                                                            _ = v141
                                                        } else {
                                                            var v142 uint16 = v28
                                                            if ((v2 < v2) && true) {
                                                                v7 += (v7 - 127)
                                                                v7 = 126
                                                                v143 := GenS20{F0: 255}
                                                                _ = v143
                                                            }
                                                            _ = v142
                                                        }
                                                    }
                                                    if (true && true) {
                                                        if (false || true) {
                                                            if true {
                                                                var v144 uint64 = 42
                                                                if (-128 == 42) {
                                                                    var v145 float32 = 0.1
                                                                    if len(v3) > 0 {
                                                                        v3[0] = 2147483648
                                                                    }
                                                                    v146 := GenS20{F0: (v7 >> ((42) & 63))}
                                                                    _ = v145
                                                                    _ = v146
                                                                } else {
                                                                    var v147 int = (fi138 - 128)
                                                                    var v148 string = v4
                                                                    v147 = (fi138 / ((fi41) | 1))
                                                                    if (752 != 391) {
                                                                        switch fi41 {
                                                                        case 4294967295:
                                                                            fmt.Println(v2)
                                                                        case 42:
                                                                            var v149 float64 = 1e10
                                                                            _ = v149
                                                                        case 32767:
                                                                            if true {
                                                                                var v150 bool = ((v144 & (v144 << ((42) & 63))) > v144)
                                                                                if len(v3) > 0 {
                                                                                    v3[0] = 182
                                                                                }
                                                                                switch "foo bar" {
                                                                                case "hello":
                                                                                    v150 = v150
                                                                                case "\n\t":
                                                                                    v151 := v6[(v148 + "foo bar")]
                                                                                    _ = v151
                                                                                default:
                                                                                    var v152 int32 = (v2 - (v2 >> ((995) & 63)))
                                                                                    _ = v152
                                                                                }
                                                                                var v153 uint32 = 735
                                                                                switch 1 {
                                                                                case -2147483648:
                                                                                    var v154 int8 = -128
                                                                                    _ = v154
                                                                                case 0:
                                                                                    var v155 uint16 = v28
                                                                                    _ = v155
                                                                                case 42:
                                                                                    v156 := GenS15{F0: 0, F1: 32768, F2: v2}
                                                                                    _ = v156
                                                                                }
                                                                                _ = v150
                                                                                _ = v153
                                                                            } else {
                                                                                var v157 bool = (2147483647 != fi138)
                                                                                switch (v147 - (fi41 / ((v147) | 1))) {
                                                                                case -835:
                                                                                    fmt.Println(fi138, v4)
                                                                                case 461:
                                                                                    v4 = (v148 + ("a" + v4))
                                                                                case 156:
                                                                                    v4 = v4
                                                                                default:
                                                                                    if v157 {
                                                                                        var v158 string = (v4 + v148)
                                                                                        _ = v158
                                                                                    } else {
                                                                                        v159 := GenS4{F0: 42, F1: (v144 * uint64(v7))}
                                                                                        _ = v159
                                                                                    }
                                                                                }
                                                                                _ = v157
                                                                            }
                                                                        default:
                                                                            if len(v3) > 0 {
                                                                                v3[0] = 619
                                                                            }
                                                                        }
                                                                        v160 := GenS21{F0: 4294967295}
                                                                        _ = v160
                                                                    } else {
                                                                        v2 += v2
                                                                        v148 = "hello"
                                                                    }
                                                                    _ = v147
                                                                    _ = v148
                                                                }
                                                                var v161 int32 = v2
                                                                v162 := v6[v4]
                                                                switch (fi138 & fi41) {
                                                                case -65:
                                                                    v162 = true
                                                                    v7 -= (v7 | v7)
                                                                case 1:
                                                                    v3 = append(v3, 42)
                                                                }
                                                                _ = v144
                                                                _ = v161
                                                                _ = v162
                                                            }
                                                            v28 = 984
                                                            var v163 int64 = -347
                                                            _ = v163
                                                        } else {
                                                            if (true || false) {
                                                                if len(v3) > 0 {
                                                                    v3[0] = 826
                                                                }
                                                            }
                                                        }
                                                    }
                                                    _ = fi138
                                                    _ = v139
                                                }
                                            case "":
                                                v7 -= (v7 * 1)
                                            case "hello":
                                                switch 683 {
                                                case 352:
                                                    fmt.Println(fi41)
                                                case 4294967295:
                                                    v28 += 0
                                                case 643:
                                                    v164 := GenS22{F0: -627, F1: 134, F2: 193}
                                                    _ = v164
                                                default:
                                                    var v165 int64 = -459
                                                    _ = v165
                                                }
                                            default:
                                                if true {
                                                    for fi166 := 0; fi166 < 6; fi166++ {
                                                        var v167 float32 = (float32((1e10 * 3.14159)) / float32(1e10))
                                                        v4 = ""
                                                        _ = fi166
                                                        _ = v167
                                                    }
                                                    for fi168 := 0; fi168 < 19; fi168++ {
                                                        var v169 uint8 = (v7 << ((0) & 63))
                                                        v3 = append(v3, 2147483647)
                                                        v1["foo bar"] = true
                                                        _ = fi168
                                                        _ = v169
                                                    }
                                                }
                                            }
                                        }
                                    }
                                }
                            } else {
                                v2 -= v2
                            }
                            fmt.Println(v7, v4)
                        } else {
                            v7 -= (^(v7))
                            v4 = v4
                        }
                        _ = v44
                    }
                    _ = fi41
                }
                v7 = v7
                var v170 uint16 = v28
                _ = v170
            }
        }
        var v171 float64 = 1.5
        _ = v171
        _ = v36
    } else {
        for fi172 := 0; fi172 < 1; fi172++ {
            v173 := GenS2{F0: 42, F1: "hello"}
            if (v2 == -345) {
                for fi174 := 0; fi174 < 17; fi174++ {
                    v28 = v28
                    _ = fi174
                }
                v4 = "a"
                for fi175 := 0; fi175 < 2; fi175++ {
                    v28 = v28
                    var v176 uint16 = 1
                    _ = fi175
                    _ = v176
                }
            } else {
                v177 := GenS23{F0: 2147483647}
                _ = v177
            }
            if (float64(0.1) < float64(float32((-2.25 - 0.1)))) {
                if false {
                    v7 = v7
                    var v178 uint8 = (^(v7))
                    v1[v4] = (!true)
                    _ = v178
                }
                var v179 int8 = 42
                _ = v179
            } else {
                for fi180 := 0; fi180 < 9; fi180++ {
                    switch 0 {
                    case 394:
                        var v181 int32 = v2
                        _ = v181
                    default:
                        var v182 string = "a"
                        _ = v182
                    }
                    var v183 uint8 = (^(v7))
                    _ = fi180
                    _ = v183
                }
                v1["\n\t"] = false
                var v184 float32 = ((float32(float64((0.1 + 1.5))) * (float32((-2.25 + -2.25)) - 0.1)) * -2.25)
                _ = v184
            }
            _ = fi172
            _ = v173
        }
    }
    if false {
        var v185 int8 = 42
        var v186 bool = (1.5 < 1e10)
        _ = v185
        _ = v186
    } else {
        v28 = (v28 >> ((uint(v28)) & 63))
    }
    var v187 int16 = -1
    v1["hello"] = (false && true)
    v7 = 239
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v187:", v187)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v187
    _ = v2
    _ = v28
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
