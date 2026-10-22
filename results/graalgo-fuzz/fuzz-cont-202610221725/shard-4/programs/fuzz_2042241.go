package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 int32; F2 int }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 int16; F1 int8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int32; F1 int32; F2 bool }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint32; F1 int64; F2 int16 }
type GenS9 struct { F0 int16; F1 bool; F2 bool }
type GenS10 struct { F0 uint; F1 int16; F2 uint32 }
type GenS11 struct { F0 uint16; F1 int; F2 uint }
type GenS12 struct { F0 string }
type GenS13 struct { F0 uint8 }
func genFn0() uint8 {
    v5 := []int{42}
    var v6 bool = true
    for fi7 := 0; fi7 < 2; fi7++ {
        v6 = false
        for fi8 := 0; fi8 < 13; fi8++ {
            v6 = ((4294967295 != 0) || (-534 > -2147483648))
            switch fi8 {
            case 42:
                v9 := GenS1{F0: 255}
                v10 := GenS1{F0: 879}
                if (708 != 307) {
                    switch (fi7 >> ((2) & 63)) {
                    case -2147483648:
                        fmt.Println(v6, fi7, v6)
                    }
                } else {
                    v6 = (482 > 507)
                }
                _ = v10
                _ = v9
            }
            if ((!true) && (fi8 > -158)) {
                v11 := GenS1{F0: 65535}
                _ = v11
            } else {
                var v12 uint8 = 127
                switch ("hello" + ("" + "foo bar")) {
                case "foo bar":
                    if ("foo bar" <= ("\n\t" + "a")) {
                        v5 = append(v5, 127)
                        var v13 uint32 = 0
                        _ = v13
                    } else {
                        v6 = v6
                        v14 := GenS2{F0: "hello", F1: 1, F2: -99}
                        _ = v14
                    }
                case "\n\t":
                    if (true && v6) {
                        var v15 int64 = 9223372036854775807
                        v16 := GenS2{F0: ("\n\t" + ("\n\t" + ("" + "\n\t"))), F1: -1, F2: (fi8 / (((fi8 + (^(fi8)))) | 1))}
                        _ = v15
                        _ = v16
                    } else {
                        fmt.Println(v6, fi7)
                    }
                default:
                    var v17 uint8 = 255
                    _ = v17
                }
                v12 = (^(v12))
                _ = v12
            }
            switch "hello" {
            case "foo bar":
                if v6 {
                    var v18 int = (fi8 << ((0) & 63))
                    var v19 int = (fi8 & (fi8 - 4294967295))
                    _ = v18
                    _ = v19
                } else {
                    var v20 int16 = -32768
                    v21 := GenS3{F0: -2147483648, F1: 750}
                    _ = v20
                    _ = v21
                }
            }
            v22 := GenS3{F0: 2147483647, F1: 1}
            switch "hello" {
            case "\n\t":
                if (42 < 255) {
                    var v23 uint = 0
                    v24 := GenS4{F0: 1, F1: 100}
                    v6 = v6
                    _ = v23
                    _ = v24
                }
            case "foo bar":
                v6 = (!(-706 != 2147483647))
            default:
                switch "foo bar" {
                case "":
                    var v25 uint64 = 0
                    _ = v25
                case "\n\t":
                    var v26 uint8 = 127
                    _ = v26
                case "hello":
                    v6 = false
                default:
                    if false {
                        fmt.Println(v6, fi7, v6)
                        v27 := GenS5{F0: 114}
                        v28 := GenS1{F0: 178}
                        _ = v27
                        _ = v28
                    } else {
                        v5 = append(v5, (^(fi8)))
                        v29 := GenS6{F0: 618, F1: 791, F2: ((float64(-2.25) / (float64(float32(-2.25)) / 3.14159)) != ((0.1 - 3.14159) + (0.0 * float64((1e10 + 0.0)))))}
                        if len(v5) > 0 {
                            v5[0] = (fi8 % (((^(fi8))) | 1))
                        }
                        _ = v29
                    }
                    v6 = (128 != 0)
                }
            }
            v6 = v6
            var v30 int64 = 9223372036854775807
            _ = fi8
            _ = v22
            _ = v30
        }
        v6 = v6
        _ = fi7
    }
    _ = v5
    _ = v6
    return 0
}
func genFn1(p0 int, p1 bool) int16 {
    var v31 uint32 = 65535
    var v32 int = p0
    v33 := []int{}
    for fi34 := 0; fi34 < 3; fi34++ {
        p1 = (true || ("" != ""))
        if false {
            for fi35 := 0; fi35 < 10; fi35++ {
                if (-27 <= 42) {
                    var v36 uint16 = 702
                    v37 := GenS6{F0: 0, F1: -1, F2: p1}
                    _ = v36
                    _ = v37
                }
                var v38 int = fi35
                if len(v33) >= 2 {
                    v33 = v33[1:len(v33)]
                }
                _ = fi35
                _ = v38
            }
            var v39 uint32 = 372
            if len(v33) >= 2 {
                v33 = v33[1:len(v33)]
            }
            _ = v39
        } else {
            p0 = (v32 * fi34)
            p0 -= (p0 - fi34)
            v33 = append(v33, -9223372036854775808)
        }
        fmt.Println(p1)
        _ = fi34
    }
    p0 -= 32768
    _ = p0
    _ = p1
    _ = v31
    _ = v32
    _ = v33
    return 974
}

func main() {
    var v1 int = 32768
    v2 := make(map[int]string)
    v3 := make(map[uint32]bool)
    var v4 uint64 = 2
    _ = len(v3)
    for fi40 := 0; fi40 < 16; fi40++ {
        v1 -= -9223372036854775808
        var v41 string = "a"
        if false {
            for fi42 := 0; fi42 < 10; fi42++ {
                v43 := v2[fi42]
                v44 := GenS7{F0: v4}
                _ = fi42
                _ = v43
                _ = v44
            }
        } else {
            switch v1 {
            case 0:
                v45 := []int{fi40, 128, fi40, (fi40 << ((65535) & 63))}
                _ = v45
            case 127:
                for fi46 := 0; fi46 < 7; fi46++ {
                    var v47 int32 = 1
                    if true {
                        var v48 int32 = (^(v47))
                        _ = v48
                    } else {
                        v49 := GenS8{F0: 97, F1: 42, F2: 299}
                        var v50 int = -1
                        delete(v3, 65535)
                        _ = v49
                        _ = v50
                    }
                    fmt.Println(v4, fi40)
                    v41 = v41
                    _ = fi46
                    _ = v47
                }
            case -9223372036854775808:
                fmt.Println(v41)
            }
            switch (fi40 * 128) {
            case 69:
                for fi51 := 0; fi51 < 14; fi51++ {
                    if (-958 > fi40) {
                        switch 403 {
                        case -2147483648:
                            v52 := []int{(fi51 / (((v1 - v1)) | 1)), (v1 / ((v1) | 1))}
                            _ = v52
                        default:
                            if ((1 != 65535) || (v4 < (v4 & v4))) {
                                v53 := []int{(fi51 >> ((255) & 63))}
                                if false {
                                    fmt.Println(v41, v1)
                                    switch (fi40 & fi51) {
                                    case 750:
                                        v41 = v41
                                    default:
                                        fmt.Println(v4)
                                    }
                                }
                                if true {
                                    _ = len(v2)
                                    var v54 bool = (0 != 82)
                                    v4 = (^(v4))
                                    v55 := GenS8{F0: 42, F1: -1, F2: -32768}
                                    _ = v54
                                    _ = v55
                                }
                                _ = v53
                            } else {
                                var v56 int8 = int8(v4)
                                v1 += (^(fi51))
                                v41 = "foo bar"
                                _ = v56
                            }
                        }
                        v57 := GenS9{F0: -32768, F1: true, F2: (-885 >= 42)}
                        _ = v57
                    }
                    _ = fi51
                }
            case -9223372036854775808:
                fmt.Println(fi40, fi40)
            case 32767:
                var v58 float32 = 0.1
                _ = v58
            }
            var v59 int32 = 1
            _ = v59
        }
        _ = fi40
        _ = v41
    }
    switch "" {
    case "\n\t":
        delete(v3, 464)
    }
    switch -60 {
    case -884:
        v1 += (v1 * v1)
    case 42:
        v60 := GenS7{F0: (v4 << ((1) & 63))}
        _ = v60
    }
    delete(v2, -1)
    fmt.Println(v4)
    if (!(-32768 != 540)) {
        if (-817 >= -394) {
            var v61 int64 = 520
            if (false || (84 < 25)) {
                for fi62 := 0; fi62 < 9; fi62++ {
                    if (!false) {
                        v4 = (v4 * (v4 / (((v4 & v4)) | 1)))
                        switch 4294967295 {
                        case -274:
                            v4 = v4
                        case 9223372036854775807:
                            v63 := v3[874]
                            _ = v63
                        }
                    } else {
                        if true {
                            v64 := GenS6{F0: -2147483648, F1: -2147483648, F2: true}
                            if (255 > 189) {
                                v4 += (v4 + 710)
                                var v65 uint32 = 2147483648
                                _ = v65
                            } else {
                                var v66 uint32 = 405
                                v1 -= (v1 + (-(v1)))
                                _ = v66
                            }
                            _ = v64
                        }
                        var v67 uint64 = v4
                        for fi68 := 0; fi68 < 15; fi68++ {
                            v69 := GenS9{F0: 345, F1: false, F2: (!(true || true))}
                            v70 := GenS10{F0: 65535, F1: -1, F2: 2}
                            if (-520 < 286) {
                                switch -335 {
                                case -346:
                                    switch ("" + "hello") {
                                    case "\n\t":
                                        var v71 float64 = -2.25
                                        _ = v71
                                    case "hello":
                                        var v72 uint32 = 675
                                        _ = v72
                                    case "foo bar":
                                        v73 := GenS8{F0: 2147483647, F1: (v61 >> ((255) & 63)), F2: 1}
                                        _ = v73
                                    }
                                }
                            }
                            _ = fi68
                            _ = v69
                            _ = v70
                        }
                        _ = v67
                    }
                    _ = fi62
                }
                for fi74 := 0; fi74 < 11; fi74++ {
                    v75 := GenS7{F0: 219}
                    _ = fi74
                    _ = v75
                }
            }
            if (v4 == 21) {
                var v76 uint16 = 32767
                if true {
                    var v77 int32 = -128
                    _ = v77
                } else {
                    if (false || true) {
                        switch ("a" + (("a" + ("" + "foo bar")) + "a")) {
                        case "":
                            if false {
                                for fi78 := 0; fi78 < 13; fi78++ {
                                    for fi79 := 0; fi79 < 9; fi79++ {
                                        fmt.Println(v4, fi78)
                                        if false {
                                            fmt.Println(v1, v61, fi79)
                                            var v80 string = "hello"
                                            if (v80 != v80) {
                                                var v81 uint64 = (v4 + 0)
                                                if (1 < 2147483648) {
                                                    var v82 uint8 = genFn0()
                                                    var v83 uint32 = 65535
                                                    v84 := []int{(fi78 | v1)}
                                                    fmt.Println(v1, v80, v61)
                                                    _ = v82
                                                    _ = v83
                                                    _ = v84
                                                }
                                                fmt.Println(v80)
                                                _ = v81
                                            }
                                            _ = v80
                                        } else {
                                            var v85 float32 = -2.25
                                            var v86 uint64 = v4
                                            _ = v85
                                            _ = v86
                                        }
                                        if (!true) {
                                            switch ("hello" + "hello") {
                                            case "foo bar":
                                                v87 := []int{9223372036854775807}
                                                _ = v87
                                            }
                                        }
                                        var v88 bool = true
                                        _ = fi79
                                        _ = v88
                                    }
                                    for fi89 := 0; fi89 < 19; fi89++ {
                                        switch v1 {
                                        case -696:
                                            if false {
                                                v90 := []int{(fi78 | fi89), 4294967295, (^(v1)), fi89}
                                                var v91 int8 = -128
                                                var v92 string = ("" + "foo bar")
                                                if false {
                                                    v92 = ""
                                                    v61 = (^(v61))
                                                    v76 = v76
                                                } else {
                                                    var v93 float64 = -2.25
                                                    var v94 uint16 = 42
                                                    if len(v90) > 0 {
                                                        v90[0] = int(v94)
                                                    }
                                                    _ = v93
                                                    _ = v94
                                                }
                                                v1 -= 42
                                                _ = v90
                                                _ = v91
                                                _ = v92
                                            } else {
                                                v61 -= v61
                                            }
                                        case -2147483648:
                                            if true {
                                                v95 := GenS5{F0: 2}
                                                v96 := []int{fi89, fi78, (fi89 ^ fi78), fi89}
                                                switch "" {
                                                case "hello":
                                                    var v97 int64 = v61
                                                    _ = v97
                                                case "\n\t":
                                                    var v98 string = ("" + "")
                                                    _ = v98
                                                case "":
                                                    var v99 uint64 = 193
                                                    _ = v99
                                                }
                                                _ = v95
                                                _ = v96
                                            } else {
                                                var v100 uint32 = 1
                                                if (uint8(v4) >= genFn0()) {
                                                    v76 = 979
                                                    switch fi78 {
                                                    case 128:
                                                        switch "foo bar" {
                                                        case "a":
                                                            var v101 int32 = -2147483648
                                                            _ = v101
                                                        }
                                                    case 2147483647:
                                                        if (127 != 86) {
                                                            var v102 int64 = (v61 & (v61 * -421))
                                                            _ = v102
                                                        }
                                                    case 0:
                                                        v76 += v76
                                                    }
                                                } else {
                                                    v103 := v2[(^(fi78))]
                                                    fmt.Println(v4)
                                                    _ = v103
                                                }
                                                _ = v100
                                            }
                                        }
                                        v3[0] = (!(((2 <= 1) || false) || true))
                                        v76 = (v76 << ((42) & 63))
                                        _ = fi89
                                    }
                                    _ = fi78
                                }
                            } else {
                                for fi104 := 0; fi104 < 15; fi104++ {
                                    for fi105 := 0; fi105 < 15; fi105++ {
                                        fmt.Println(v4, v61)
                                        var v106 bool = false
                                        v107 := GenS3{F0: -2147483648, F1: 42}
                                        v106 = (("a" + (("hello" + "\n\t") + "a")) <= ("" + "\n\t"))
                                        v1 -= (-(v1))
                                        v106 = v106
                                        if v106 {
                                            v108 := GenS1{F0: 42}
                                            _ = v108
                                        }
                                        _ = fi105
                                        _ = v106
                                        _ = v107
                                    }
                                    _ = fi104
                                }
                                if (((!false) && (!(!(1 < 2147483648)))) && false) {
                                    var v109 int64 = v61
                                    var v110 uint = 617
                                    v4 += (v4 - (v4 >> ((v110) & 63)))
                                    _ = v109
                                    _ = v110
                                }
                                v111 := GenS11{F0: v76, F1: 128, F2: uint(v76)}
                                _ = v111
                            }
                        default:
                            for fi112 := 0; fi112 < 12; fi112++ {
                                v113 := GenS7{F0: 0}
                                v114 := GenS12{F0: "\n\t"}
                                var v115 uint16 = v76
                                _ = fi112
                                _ = v113
                                _ = v114
                                _ = v115
                            }
                        }
                        if false {
                            var v116 int16 = 386
                            if true {
                                v76 = (^(v76))
                                v4 += (v4 ^ (v4 * (v4 - (v4 % ((v4) | 1)))))
                            } else {
                                var v117 uint32 = 4294967295
                                _ = v117
                            }
                            if (2147483648 < 273) {
                                v118 := GenS1{F0: 137}
                                v119 := []int{}
                                _ = v118
                                _ = v119
                            } else {
                                switch ("" + ("" + "hello")) {
                                case "hello":
                                    v120 := []int{(^(v1)), int(v61), (-(v1))}
                                    _ = v120
                                case "a":
                                    var v121 int8 = 0
                                    _ = v121
                                case "\n\t":
                                    fmt.Println(v116)
                                }
                                switch "" {
                                case "\n\t":
                                    v122 := v2[v1]
                                    _ = v122
                                case "hello":
                                    for fi123 := 0; fi123 < 8; fi123++ {
                                        var v124 string = ("foo bar" + "hello")
                                        _ = fi123
                                        _ = v124
                                    }
                                case "":
                                    v116 -= v116
                                }
                            }
                            _ = v116
                        } else {
                            v76 -= (v76 | 32767)
                            if (29 >= 62) {
                                v61 += -314
                            }
                        }
                        var v125 float32 = (1e10 - float32(0.1))
                        _ = v125
                    } else {
                        var v126 uint = 4294967295
                        _ = v126
                    }
                }
                _ = v76
            }
            _ = v61
        } else {
            for fi127 := 0; fi127 < 6; fi127++ {
                if false {
                    var v128 int64 = -9223372036854775808
                    _ = v128
                }
                fmt.Println(v4, fi127, v4)
                _ = fi127
            }
            v4 -= 465
        }
    }
    fmt.Println(v4, v4)
    var v129 string = "\n\t"
    v130 := v2[-830]
    for fi131 := 0; fi131 < 20; fi131++ {
        for fi132 := 0; fi132 < 4; fi132++ {
            var v133 uint = 0
            switch (^(fi131)) {
            case 325:
                v130 = (v129 + (v129 + "a"))
            case 0:
                v129 = ("" + v130)
            case 127:
                switch 1 {
                case 4294967295:
                    fmt.Println(v130, v129, v1)
                case 32767:
                    v134 := []int{(fi132 / ((fi131) | 1))}
                    v1 += fi132
                    switch v129 {
                    case "\n\t":
                        if true {
                            v4 = (v4 / (((v4 - v4)) | 1))
                        }
                        v135 := v3[2147483648]
                        var v136 float64 = (1.5 + (1e10 * 1.5))
                        v137 := GenS9{F0: 1000, F1: false, F2: (v135 || (v135 && (v135 || (21 >= -96))))}
                        _ = v135
                        _ = v136
                        _ = v137
                    }
                    var v138 uint = (v133 + 255)
                    _ = v134
                    _ = v138
                case 0:
                    if (1 > -995) {
                        delete(v3, 732)
                    }
                }
            default:
                var v139 uint32 = 0
                _ = v139
            }
            if false {
                var v140 uint8 = 42
                var v141 int16 = -32768
                _ = v140
                _ = v141
            } else {
                switch (^(fi131)) {
                case -1:
                    var v142 int16 = 32767
                    _ = v142
                }
                v143 := GenS2{F0: v130, F1: 317, F2: (v1 % (((fi132 ^ -963)) | 1))}
                v144 := GenS11{F0: uint16(v4), F1: 42, F2: 18446744073709551615}
                _ = v143
                _ = v144
            }
            switch (v1 / ((-9223372036854775808) | 1)) {
            case 32767:
                v2[(fi131 & fi132)] = v130
            case 9223372036854775807:
                var v145 uint8 = 225
                _ = v145
            default:
                v146 := []int{0, (v1 % ((fi131) | 1))}
                _ = v146
            }
            _ = fi132
            _ = v133
        }
        _ = fi131
    }
    var v147 bool = (42 < genFn0())
    switch v1 {
    case 2147483647:
        if false {
            v4 -= v4
        } else {
            var v148 uint64 = v4
            if v147 {
                switch v130 {
                case "foo bar":
                    fmt.Println(v130, v1, v147)
                case "hello":
                    v149 := GenS13{F0: genFn0()}
                    _ = v149
                case "a":
                    v4 += (v148 / (((^(v148))) | 1))
                }
                if v147 {
                    v4 -= (^(v148))
                    if (0.1 <= 1.5) {
                        var v150 int16 = 0
                        v3[582] = v147
                        _ = v150
                    } else {
                        var v151 uint32 = 42
                        _ = v151
                    }
                    for fi152 := 0; fi152 < 19; fi152++ {
                        v1 = (fi152 << ((42) & 63))
                        _ = fi152
                    }
                }
                v153 := v3[2147483647]
                _ = v153
            } else {
                v148 += v148
                v154 := v3[2147483647]
                for fi155 := 0; fi155 < 11; fi155++ {
                    fmt.Println(v148)
                    switch v1 {
                    case 416:
                        fmt.Println(v154, v148, v130)
                    default:
                        fmt.Println(v147, v4)
                    }
                    v148 += (v148 / ((v148) | 1))
                    _ = fi155
                }
                _ = v154
            }
            _ = v148
        }
    default:
        for fi156 := 0; fi156 < 8; fi156++ {
            fmt.Println(v129, v130)
            switch 42 {
            case 32768:
                switch (-(fi156)) {
                case -9223372036854775808:
                    if v147 {
                        _ = len(v3)
                    }
                default:
                    var v157 uint8 = 42
                    _ = v157
                }
            case 563:
                v1 += 32767
            case 4294967295:
                var v158 int = 835
                _ = v158
            default:
                var v159 uint32 = 880
                _ = v159
            }
            _ = fi156
        }
    }
    fmt.Println(v129)
    fmt.Println("v1:", v1)
    fmt.Println("v129:", v129)
    fmt.Println("v130:", v130)
    fmt.Println("v147:", v147)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v129
    _ = v130
    _ = v147
    _ = v2
    _ = v3
    _ = v4
}
