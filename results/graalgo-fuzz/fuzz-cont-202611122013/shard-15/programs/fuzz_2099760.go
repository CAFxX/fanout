package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool; F1 int8 }
type GenS4 struct { F0 uint16; F1 uint8 }
type GenS5 struct { F0 bool; F1 int32 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int; F1 uint32; F2 int8 }
type GenS9 struct { F0 uint16; F1 uint; F2 uint32 }
type GenS10 struct { F0 bool; F1 uint16; F2 uint8 }
type GenS11 struct { F0 uint16; F1 uint; F2 bool }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 uint64 }
type GenS15 struct { F0 uint; F1 uint16 }
type GenS16 struct { F0 bool }
type GenS17 struct { F0 int32; F1 string }
type GenS18 struct { F0 int }
type GenS19 struct { F0 uint16; F1 int }

func main() {
    var v1 uint16 = 1
    var v2 uint = 65535
    var v3 int32 = -2147483648
    for fi4 := 0; fi4 < 8; fi4++ {
        v5 := GenS1{F0: -461, F1: uint8(v1), F2: (("a" + "hello") + ((("\n\t" + "hello") + "\n\t") + "foo bar"))}
        var v6 float64 = -2.25
        if false {
            var v7 uint32 = 394
            var v8 int = (fi4 << (((v2 / ((uint(v1)) | 1))) & 63))
            if (!(v7 == v7)) {
                if (((v2 >> ((65535) & 63)) < (^(v2))) && ((!false) || (!(("" + "") < "\n\t")))) {
                    var v9 uint8 = 1
                    for fi10 := 0; fi10 < 10; fi10++ {
                        v11 := GenS1{F0: 1000, F1: (v9 - v9), F2: ""}
                        _ = fi10
                        _ = v11
                    }
                    v12 := GenS1{F0: int16(v7), F1: 143, F2: ("hello" + ("" + (("\n\t" + ("\n\t" + "")) + ("\n\t" + (("\n\t" + "a") + "hello")))))}
                    _ = v12
                    _ = v9
                } else {
                    for fi13 := 0; fi13 < 14; fi13++ {
                        var v14 string = "\n\t"
                        fmt.Println(v7, fi13)
                        _ = fi13
                        _ = v14
                    }
                }
                fmt.Println(v7, v1, v6)
                v15 := GenS1{F0: -32768, F1: uint8(v7), F2: "\n\t"}
                _ = v15
            } else {
                var v16 float32 = 1.5
                for fi17 := 0; fi17 < 18; fi17++ {
                    v18 := GenS1{F0: 233, F1: 127, F2: ("\n\t" + (("hello" + "foo bar") + "\n\t"))}
                    _ = fi17
                    _ = v18
                }
                _ = v16
            }
            _ = v7
            _ = v8
        } else {
            var v19 float32 = (1.5 * 0.1)
            for fi20 := 0; fi20 < 9; fi20++ {
                v2 = (v2 | 42)
                if (!false) {
                    switch (fi20 ^ (fi4 + fi20)) {
                    case 0:
                        v21 := []int{(fi4 * 1), -2147483648, (^(fi20))}
                        _ = v21
                    }
                    v22 := []int{-9223372036854775808, fi20, (fi4 << ((v2) & 63)), fi4}
                    _ = len(v22)
                    _ = v22
                } else {
                    var v23 uint8 = uint8(v1)
                    switch "hello" {
                    case "a":
                        v24 := GenS1{F0: -1, F1: 214, F2: "hello"}
                        _ = v24
                    default:
                        v23 -= (v23 & (v23 % ((127) | 1)))
                    }
                    _ = v23
                }
                _ = fi20
            }
            _ = v19
        }
        _ = fi4
        _ = v5
        _ = v6
    }
    v3 = v3
    var v25 string = "a"
    v26 := []int{-470, 127, -633}
    v27 := GenS1{F0: 1000, F1: 255, F2: (v25 + "foo bar")}
    var v28 uint64 = 167
    if (float64(1.5) == 0.1) {
        var v29 string = (((v25 + "hello") + (v25 + v25)) + v25)
        if true {
            fmt.Println(v25, v25, v25)
            if len(v26) >= 2 {
                v26 = v26[1:len(v26)]
            }
        } else {
            v30 := GenS2{F0: 9223372036854775807}
            _ = v30
        }
        var v31 uint16 = v1
        _ = v29
        _ = v31
    } else {
        if (v25 < (v25 + "\n\t")) {
            var v32 uint = (^(v2))
            v33 := GenS2{F0: 42}
            if (-1 < 42) {
                fmt.Println(v1, v2)
                for fi34 := 0; fi34 < 4; fi34++ {
                    v32 -= v32
                    _ = fi34
                }
                switch v25 {
                case "\n\t":
                    if false {
                        switch 444 {
                        case 128:
                            for fi35 := 0; fi35 < 16; fi35++ {
                                v2 = uint(v28)
                                switch (fi35 - 4294967295) {
                                case 618:
                                    for fi36 := 0; fi36 < 18; fi36++ {
                                        fmt.Println(v25, v32, fi35)
                                        _ = fi36
                                    }
                                case 32768:
                                    if (!(981 <= 65535)) {
                                        v28 += v28
                                        fmt.Println(v1, v25)
                                        v37 := GenS1{F0: 1000, F1: 42, F2: v25}
                                        _ = v37
                                    }
                                case -1:
                                    if (409 >= v28) {
                                        for fi38 := 0; fi38 < 7; fi38++ {
                                            if true {
                                                var v39 float64 = float64(((0.1 - 1e10) / 3.14159))
                                                _ = v39
                                            }
                                            _ = fi38
                                        }
                                        var v40 uint8 = uint8(v32)
                                        var v41 int8 = 100
                                        _ = v40
                                        _ = v41
                                    }
                                default:
                                    if (0 == 42) {
                                        if false {
                                            if false {
                                                fmt.Println(v3, v1, v1)
                                                for fi42 := 0; fi42 < 2; fi42++ {
                                                    v26 = append(v26, -978)
                                                    var v43 int32 = v3
                                                    var v44 int16 = int16(v32)
                                                    _ = fi42
                                                    _ = v43
                                                    _ = v44
                                                }
                                                switch v25 {
                                                case "a":
                                                    if len(v26) >= 2 {
                                                        v26 = v26[1:len(v26)]
                                                    }
                                                case "hello":
                                                    for fi45 := 0; fi45 < 5; fi45++ {
                                                        var v46 int = fi45
                                                        if (-9223372036854775808 != 618) {
                                                            v47 := GenS3{F0: (!false), F1: -58}
                                                            if len(v26) > 0 {
                                                                v26[0] = (^(fi45))
                                                            }
                                                            v26 = append(v26, fi35)
                                                            _ = v47
                                                        } else {
                                                            if false {
                                                                var v48 uint = 18446744073709551615
                                                                switch "hello" {
                                                                case "a":
                                                                    v25 = v25
                                                                case "":
                                                                    v49 := GenS4{F0: (v1 & (v1 | (^(v1)))), F1: 242}
                                                                    _ = v49
                                                                }
                                                                _ = v48
                                                            } else {
                                                                if true {
                                                                    var v50 float32 = 1.5
                                                                    _ = v50
                                                                } else {
                                                                    if ((fi35 ^ (fi45 * (fi45 << (((v2 - (v2 / (((v2 >> (((v32 << (((v32 + (v32 ^ v2))) & 63))) & 63))) | 1)))) & 63)))) != 32768) {
                                                                        v51 := GenS5{F0: false, F1: int32(v32)}
                                                                        v32 = (v32 >> ((v32) & 63))
                                                                        _ = v51
                                                                    }
                                                                    switch -2147483648 {
                                                                    case 128:
                                                                        var v52 uint = v32
                                                                        _ = v52
                                                                    default:
                                                                        var v53 uint32 = 42
                                                                        _ = v53
                                                                    }
                                                                }
                                                            }
                                                            v54 := GenS6{F0: 2}
                                                            if true {
                                                                if true {
                                                                    v46 -= v46
                                                                }
                                                                if ((false && false) && (0 != 100)) {
                                                                    v55 := GenS7{F0: v1}
                                                                    if (v46 <= fi35) {
                                                                        if len(v26) >= 2 {
                                                                            v26 = v26[1:len(v26)]
                                                                        }
                                                                        v25 = v25
                                                                        switch "\n\t" {
                                                                        case "\n\t":
                                                                            var v56 uint16 = 2
                                                                            _ = v56
                                                                        case "":
                                                                            if false {
                                                                                if false {
                                                                                    if (1.5 > ((0.1 + float64(0.1)) / 1.0)) {
                                                                                        var v57 int8 = 100
                                                                                        if (v3 < -2147483648) {
                                                                                            var v58 int16 = 375
                                                                                            _ = v58
                                                                                        }
                                                                                        _ = v57
                                                                                    } else {
                                                                                        var v59 int32 = int32(v2)
                                                                                        var v60 uint16 = 743
                                                                                        v2 = 1
                                                                                        _ = v59
                                                                                        _ = v60
                                                                                    }
                                                                                    v25 = "\n\t"
                                                                                    v26 = append(v26, (fi35 + fi45))
                                                                                } else {
                                                                                    var v61 string = v25
                                                                                    _ = v61
                                                                                }
                                                                                switch (v25 + ("foo bar" + (v25 + (v25 + v25)))) {
                                                                                case "a":
                                                                                    v2 = (v32 * v2)
                                                                                case "foo bar":
                                                                                    var v62 bool = true
                                                                                    _ = v62
                                                                                default:
                                                                                    var v63 uint32 = 241
                                                                                    _ = v63
                                                                                }
                                                                            }
                                                                            var v64 int8 = 0
                                                                            v65 := GenS4{F0: v1, F1: 239}
                                                                            _ = v64
                                                                            _ = v65
                                                                        case "hello":
                                                                            var v66 uint8 = 1
                                                                            _ = v66
                                                                        }
                                                                        var v67 string = v25
                                                                        _ = v67
                                                                    } else {
                                                                        v26 = append(v26, v46)
                                                                        if (false && true) {
                                                                            if len(v26) >= 2 {
                                                                                v26 = v26[1:len(v26)]
                                                                            }
                                                                            v32 -= (v32 ^ (v2 - (v32 ^ (^(v2)))))
                                                                            v25 = v25
                                                                            v32 -= v32
                                                                        }
                                                                        var v68 bool = true
                                                                        _ = v68
                                                                    }
                                                                    _ = v55
                                                                } else {
                                                                    v28 += v28
                                                                    v28 -= 467
                                                                    var v69 uint = (v32 >> ((65535) & 63))
                                                                    _ = v69
                                                                }
                                                                fmt.Println(v1)
                                                            }
                                                            _ = v54
                                                        }
                                                        switch -181 {
                                                        case -761:
                                                            switch "" {
                                                            case "a":
                                                                v26 = append(v26, (fi35 % ((fi45) | 1)))
                                                            }
                                                            switch v25 {
                                                            case "foo bar":
                                                                var v70 uint = v32
                                                                _ = v70
                                                            }
                                                            var v71 float64 = ((((0.0 - 1.5) + -2.25) / (float64(float32((float64((1.5 * 0.0)) * (0.0 / 1.5)))) - (float64(3.14159) / (float64(1e10) - (((1e10 * float64(-2.25)) - ((((-2.25 / 1.5) - 0.0) + float64(1e10)) + ((((float64(-2.25) + (((1e10 + float64(1.5)) * 0.0) / (float64(0.1) + 0.1))) * 3.14159) + 1.5) - (3.14159 / 1.0)))) / (1.5 / 1e10)))))) + 3.14159)
                                                            _ = v71
                                                        default:
                                                            if (!(1 >= 128)) {
                                                                var v72 int8 = 0
                                                                v72 -= v72
                                                                if len(v26) >= 2 {
                                                                    v26 = v26[1:len(v26)]
                                                                }
                                                                _ = v72
                                                            }
                                                        }
                                                        v32 -= v32
                                                        switch "hello" {
                                                        case "\n\t":
                                                            v1 = v1
                                                        case "a":
                                                            v25 = v25
                                                        case "foo bar":
                                                            var v73 uint8 = uint8(v32)
                                                            _ = v73
                                                        default:
                                                            v74 := GenS5{F0: (788 != 9223372036854775807), F1: (v3 >> (((v32 * v32)) & 63))}
                                                            _ = v74
                                                        }
                                                        _ = fi45
                                                        _ = v46
                                                    }
                                                }
                                            } else {
                                                for fi75 := 0; fi75 < 13; fi75++ {
                                                    v76 := GenS8{F0: (fi35 << (((v2 ^ 4294967295)) & 63)), F1: 65535, F2: 0}
                                                    var v77 int = fi35
                                                    v26 = append(v26, fi35)
                                                    v3 -= v3
                                                    _ = fi75
                                                    _ = v76
                                                    _ = v77
                                                }
                                                for fi78 := 0; fi78 < 3; fi78++ {
                                                    var v79 int = int(v3)
                                                    v2 = (^(v2))
                                                    _ = fi78
                                                    _ = v79
                                                }
                                                v2 = 245
                                            }
                                        } else {
                                            if false {
                                                switch ("hello" + "\n\t") {
                                                case "a":
                                                    for fi80 := 0; fi80 < 2; fi80++ {
                                                        v1 += 255
                                                        v81 := GenS9{F0: (v1 >> (((^(v32))) & 63)), F1: v2, F2: 599}
                                                        var v82 string = (v25 + v25)
                                                        _ = fi80
                                                        _ = v81
                                                        _ = v82
                                                    }
                                                case "hello":
                                                    if true {
                                                        var v83 int16 = 0
                                                        fmt.Println(v25, v25)
                                                        var v84 uint = (v2 - 255)
                                                        _ = v83
                                                        _ = v84
                                                    } else {
                                                        var v85 uint64 = v28
                                                        var v86 float64 = (1.5 - 1.5)
                                                        _ = v85
                                                        _ = v86
                                                    }
                                                default:
                                                    v28 += v28
                                                }
                                                if len(v26) > 0 {
                                                    v26[0] = 854
                                                }
                                            }
                                            for fi87 := 0; fi87 < 2; fi87++ {
                                                var v88 uint8 = 52
                                                v89 := GenS9{F0: v1, F1: (v32 << ((v32) & 63)), F2: 2147483647}
                                                var v90 uint16 = v1
                                                var v91 int8 = -34
                                                v90 += (v1 << ((v32) & 63))
                                                _ = fi87
                                                _ = v88
                                                _ = v89
                                                _ = v90
                                                _ = v91
                                            }
                                        }
                                        var v92 int32 = (v3 * -645)
                                        if len(v26) >= 2 {
                                            v26 = v26[1:len(v26)]
                                        }
                                        _ = v92
                                    }
                                }
                                _ = fi35
                            }
                        case -559:
                            var v93 int64 = 0
                            _ = v93
                        case 32768:
                            fmt.Println(v1, v3)
                        default:
                            v25 = ((v25 + v25) + (v25 + v25))
                        }
                        if false {
                            if len(v26) >= 2 {
                                v26 = v26[1:len(v26)]
                            }
                        }
                    }
                case "a":
                    var v94 uint8 = 1
                    _ = v94
                default:
                    for fi95 := 0; fi95 < 17; fi95++ {
                        v2 = uint(v28)
                        var v96 int8 = -97
                        var v97 bool = true
                        _ = fi95
                        _ = v96
                        _ = v97
                    }
                }
            } else {
                var v98 uint32 = 4294967295
                if (v1 <= v1) {
                    if false {
                        if false {
                            if ((true || (!(!(120 != 279)))) && (-2.25 <= 0.0)) {
                                var v99 float64 = float64((0.0 * float32((0.1 - (3.14159 - (3.14159 + float64((-2.25 + (3.14159 * 1.5)))))))))
                                _ = v99
                            } else {
                                if false {
                                    if ((!(4294967295 > uint32(v28))) || (!true)) {
                                        var v100 int32 = (v3 + v3)
                                        v26 = append(v26, 127)
                                        _ = v100
                                    }
                                    v101 := GenS8{F0: -2147483648, F1: (v98 % ((v98) | 1)), F2: -1}
                                    v1 -= v1
                                    _ = v101
                                }
                                if (1e10 >= -2.25) {
                                    switch 9223372036854775807 {
                                    case 128:
                                        v25 = "\n\t"
                                    }
                                    v32 = (v32 + (^(v32)))
                                }
                                v28 += (v28 | (v28 << (((^(v2))) & 63)))
                            }
                            if false {
                                switch v25 {
                                case "foo bar":
                                    var v102 uint64 = uint64(v1)
                                    _ = v102
                                case "hello":
                                    switch "a" {
                                    case "":
                                        for fi103 := 0; fi103 < 20; fi103++ {
                                            for fi104 := 0; fi104 < 15; fi104++ {
                                                _ = len(v26)
                                                v2 = (v32 ^ (v32 / ((v2) | 1)))
                                                if len(v26) >= 2 {
                                                    v26 = v26[1:len(v26)]
                                                }
                                                _ = fi104
                                            }
                                            v105 := GenS4{F0: v1, F1: 80}
                                            v106 := GenS10{F0: (-128 < 0), F1: 2, F2: 0}
                                            _ = fi103
                                            _ = v105
                                            _ = v106
                                        }
                                    case "a":
                                        if len(v26) > 0 {
                                            v26[0] = -9223372036854775808
                                        }
                                    }
                                default:
                                    v2 -= 18446744073709551615
                                }
                                v107 := GenS6{F0: v32}
                                _ = v107
                            }
                        }
                        v108 := GenS11{F0: (^(v1)), F1: v32, F2: false}
                        v109 := GenS8{F0: 32767, F1: v98, F2: 0}
                        _ = v108
                        _ = v109
                    }
                    var v110 float64 = ((-2.25 + ((1.5 + 1.5) + ((-2.25 + 1.5) * 3.14159))) + (-2.25 / 1e10))
                    v110 += v110
                    _ = v110
                } else {
                    _ = len(v26)
                    if ((0 >= int32(v2)) || false) {
                        var v111 uint32 = v98
                        _ = v111
                    } else {
                        switch "hello" {
                        case "foo bar":
                            if len(v26) >= 2 {
                                v26 = v26[1:len(v26)]
                            }
                        case "hello":
                            switch v25 {
                            case "":
                                var v112 uint16 = v1
                                _ = v112
                            case "\n\t":
                                var v113 uint32 = (v98 * (v98 >> (((v32 + v2)) & 63)))
                                _ = v113
                            case "hello":
                                var v114 uint64 = 951
                                _ = v114
                            default:
                                var v115 bool = true
                                _ = v115
                            }
                        case "":
                            v2 = (^(v32))
                        }
                    }
                }
                _ = v98
            }
            _ = v32
            _ = v33
        }
        var v116 int8 = 114
        _ = v116
    }
    v117 := GenS12{F0: 65535}
    if (v25 != "\n\t") {
        if true {
            switch 127 {
            case 9223372036854775807:
                if (!true) {
                    for fi118 := 0; fi118 < 16; fi118++ {
                        v26 = append(v26, fi118)
                        v119 := GenS5{F0: true, F1: (v3 ^ v3)}
                        v120 := GenS7{F0: v1}
                        _ = fi118
                        _ = v119
                        _ = v120
                    }
                    var v121 int64 = -1
                    if true {
                        fmt.Println(v2, v121)
                    } else {
                        for fi122 := 0; fi122 < 5; fi122++ {
                            for fi123 := 0; fi123 < 16; fi123++ {
                                var v124 int16 = -478
                                fmt.Println(v3, v1)
                                _ = fi123
                                _ = v124
                            }
                            if (true || (uint64(v1) == 2)) {
                                v121 += (v121 & v121)
                            } else {
                                var v125 uint64 = (v28 * (^(v28)))
                                var v126 uint16 = (v1 / ((v1) | 1))
                                _ = v125
                                _ = v126
                            }
                            if false {
                                var v127 uint = v2
                                _ = len(v26)
                                _ = v127
                            } else {
                                var v128 int = fi122
                                var v129 bool = false
                                _ = v128
                                _ = v129
                            }
                            _ = fi122
                        }
                        switch 1 {
                        case 564:
                            var v130 float64 = (3.14159 / 3.14159)
                            _ = v130
                        }
                        switch (v25 + "a") {
                        case "":
                            v25 = "foo bar"
                        case "hello":
                            for fi131 := 0; fi131 < 7; fi131++ {
                                v1 += (^(v1))
                                _ = fi131
                            }
                        }
                    }
                    _ = v121
                }
            }
            if (127 < -675) {
                _ = len(v26)
                var v132 uint32 = 0
                _ = v132
            } else {
                v28 = (v28 << ((v2) & 63))
                if ((!((9223372036854775807 == 4294967295) && true)) || false) {
                    if (!true) {
                        var v133 bool = false
                        _ = v133
                    } else {
                        if ((948 == 46) && false) {
                            var v134 int64 = 0
                            _ = v134
                        }
                        var v135 float32 = (-2.25 / 0.1)
                        _ = v135
                    }
                    if len(v26) >= 2 {
                        v26 = v26[1:len(v26)]
                    }
                    for fi136 := 0; fi136 < 2; fi136++ {
                        switch "hello" {
                        case "a":
                            v137 := GenS1{F0: int16(v1), F1: 255, F2: ((((((v25 + (v25 + (v25 + v25))) + v25) + v25) + v25) + "foo bar") + ("a" + v25))}
                            _ = v137
                        case "\n\t":
                            for fi138 := 0; fi138 < 15; fi138++ {
                                switch ("hello" + v25) {
                                case "hello":
                                    v25 = (("" + v25) + v25)
                                case "a":
                                    var v139 int64 = 798
                                    _ = v139
                                default:
                                    var v140 float64 = (float64(0.0) + 1e10)
                                    _ = v140
                                }
                                v141 := GenS4{F0: v1, F1: 128}
                                _ = fi138
                                _ = v141
                            }
                        case "foo bar":
                            v2 = v2
                        default:
                            v142 := GenS3{F0: (-32768 > 1000), F1: -1}
                            _ = v142
                        }
                        if true {
                            v143 := GenS13{F0: 2147483647}
                            v1 += (^(v1))
                            _ = v143
                        }
                        _ = fi136
                    }
                }
            }
            if len(v26) > 0 {
                v26[0] = 9223372036854775807
            }
        }
        switch (v25 + v25) {
        case "\n\t":
            var v144 float64 = 1e10
            _ = v144
        case "":
            v28 += v28
        case "hello":
            for fi145 := 0; fi145 < 19; fi145++ {
                if (((((v28 >> ((0) & 63)) != 42) && true) && (float64((-2.25 / 1e10)) < -2.25)) || (fi145 == fi145)) {
                    if (505 < (^(v1))) {
                        switch v25 {
                        case "a":
                            if (-1 != 9223372036854775807) {
                                if true {
                                    v146 := GenS9{F0: (v1 & 396), F1: (v2 + 1), F2: 2}
                                    _ = v146
                                }
                                if len(v26) >= 2 {
                                    v26 = v26[1:len(v26)]
                                }
                            }
                        case "foo bar":
                            v147 := GenS14{F0: v28}
                            _ = v147
                        case "\n\t":
                            for fi148 := 0; fi148 < 8; fi148++ {
                                _ = len(v26)
                                v149 := GenS15{F0: 4294967295, F1: v1}
                                fmt.Println(v25)
                                if len(v26) > 0 {
                                    v26[0] = fi148
                                }
                                _ = fi148
                                _ = v149
                            }
                        }
                    }
                } else {
                    var v150 int = (fi145 + (fi145 - fi145))
                    _ = v150
                }
                var v151 int16 = 1000
                _ = fi145
                _ = v151
            }
        default:
            for fi152 := 0; fi152 < 8; fi152++ {
                for fi153 := 0; fi153 < 1; fi153++ {
                    v154 := GenS16{F0: false}
                    if len(v26) > 0 {
                        v26[0] = (fi153 - fi153)
                    }
                    v1 = 65535
                    if false {
                        var v155 string = "a"
                        _ = v155
                    }
                    if true {
                        switch v25 {
                        case "a":
                            fmt.Println(fi152, v1, fi153)
                        case "":
                            v2 = (^(v2))
                        }
                        var v156 int64 = -317
                        _ = v156
                    }
                    _ = len(v26)
                    _ = fi153
                    _ = v154
                }
                _ = fi152
            }
        }
        fmt.Println(v25, v28)
    }
    var v157 int8 = 100
    v158 := GenS16{F0: (true && (-2.25 > 1.5))}
    v3 += 0
    fmt.Println(v1, v3)
    v159 := GenS13{F0: 42}
    if (128 > 0) {
        _ = len(v26)
        v3 += -904
    }
    v28 += 0
    for fi160 := 0; fi160 < 11; fi160++ {
        v26 = append(v26, (fi160 - 0))
        if ((!(2 >= 1)) || (true && false)) {
            if false {
                fmt.Println(v28, v157)
            }
            v25 = v25
            v1 -= v1
        }
        _ = fi160
    }
    if (-284 == -997) {
        for fi161 := 0; fi161 < 8; fi161++ {
            if len(v26) > 0 {
                v26[0] = (fi161 | fi161)
            }
            fmt.Println(v25, v2)
            _ = fi161
        }
        v162 := GenS14{F0: (v28 | (v28 + (v28 >> (((v2 ^ 1000)) & 63))))}
        switch 2147483647 {
        case 127:
            if (true || (!false)) {
                var v163 uint16 = 42
                v164 := GenS13{F0: 1}
                _ = v163
                _ = v164
            }
        case 9223372036854775807:
            if true {
                _ = len(v26)
                if ((v28 & (v28 / ((uint64(v2)) | 1))) > 4294967295) {
                    v157 = 127
                    var v165 uint = v2
                    if false {
                        switch -2147483648 {
                        case -2147483648:
                            v26 = append(v26, 128)
                        default:
                            v166 := GenS13{F0: 1}
                            _ = v166
                        }
                        var v167 int32 = v3
                        fmt.Println(v28, v165, v2)
                        _ = v167
                    }
                    _ = v165
                }
                for fi168 := 0; fi168 < 7; fi168++ {
                    if (!(0 < 0)) {
                        v169 := GenS17{F0: -94, F1: v25}
                        switch "hello" {
                        case "a":
                            switch 128 {
                            case 32767:
                                fmt.Println(v25, v3)
                            case 462:
                                v2 = (v2 >> ((255) & 63))
                            case -167:
                                switch (("\n\t" + ("foo bar" + v25)) + "\n\t") {
                                case "foo bar":
                                    _ = len(v26)
                                }
                            }
                        }
                        v170 := GenS9{F0: (v1 & (v1 / ((v1) | 1))), F1: 255, F2: 505}
                        _ = v169
                        _ = v170
                    }
                    _ = fi168
                }
            } else {
                v171 := GenS18{F0: -691}
                v172 := GenS19{F0: (v1 >> (((v2 / (((v2 + 4294967295)) | 1))) & 63)), F1: 0}
                _ = v171
                _ = v172
            }
        }
        _ = v162
    } else {
        var v173 float32 = 0.1
        _ = v173
    }
    fmt.Println("v1:", v1)
    fmt.Println("v157:", v157)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v117
    _ = v157
    _ = v158
    _ = v159
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
}
