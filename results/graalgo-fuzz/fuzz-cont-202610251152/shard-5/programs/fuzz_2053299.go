package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 uint64; F1 uint64; F2 string }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int8; F1 bool; F2 uint16 }
type GenS6 struct { F0 uint16; F1 int64; F2 int32 }
type GenS7 struct { F0 uint8; F1 bool }
type GenS8 struct { F0 uint; F1 bool }
type GenS9 struct { F0 int64; F1 uint16; F2 uint }
type GenS10 struct { F0 int8; F1 bool }
type GenS11 struct { F0 int8; F1 uint32; F2 uint }
type GenS12 struct { F0 bool }
func genFn0(p0 int16) int32 {
    p0 -= (^(p0))
    v4 := GenS1{F0: -128, F1: 2, F2: 161}
    _ = p0
    _ = v4
    return -1
}
func genFn1() uint64 {
    v5 := []int{872}
    var v6 int = -9223372036854775808
    _ = v5
    _ = v6
    return 42
}

func main() {
    v1 := []uint32{2147483648}
    var v2 string = "foo bar"
    var v3 int8 = -1
    for fi7 := 0; fi7 < 13; fi7++ {
        v2 = v2
        _ = fi7
    }
    v2 = v2
    v3 = -128
    var v8 float64 = 1.5
    var v9 uint16 = 65535
    switch v2 {
    case "a":
        if false {
            v10 := GenS2{F0: v9, F1: 373}
            _ = v10
        } else {
            for fi11 := 0; fi11 < 19; fi11++ {
                switch (fi11 % ((fi11) | 1)) {
                case -1:
                    if true {
                        v12 := GenS3{F0: 2, F1: genFn1(), F2: v2}
                        switch (v2 + (v2 + v2)) {
                        case "foo bar":
                            for fi13 := 0; fi13 < 1; fi13++ {
                                switch fi11 {
                                case 127:
                                    if true {
                                        v2 = "a"
                                    }
                                case 710:
                                    v14 := GenS3{F0: 1, F1: 445, F2: v2}
                                    _ = v14
                                }
                                if (((v2 + "\n\t") + (v2 + (v2 + "hello"))) > (v2 + (v2 + ("" + v2)))) {
                                    var v15 float32 = -2.25
                                    v16 := GenS3{F0: 2, F1: 941, F2: v2}
                                    if len(v1) > 0 {
                                        v1[0] = 2147483647
                                    }
                                    if len(v1) > 0 {
                                        v1[0] = 2147483647
                                    }
                                    _ = v15
                                    _ = v16
                                }
                                fmt.Println(v9, v9, fi13)
                                _ = fi13
                            }
                        }
                        _ = v12
                    }
                case -640:
                    v3 -= int8(v9)
                case 621:
                    fmt.Println(v9)
                }
                _ = fi11
            }
            switch 127 {
            case -2147483648:
                if true {
                    var v17 bool = false
                    if true {
                        var v18 uint8 = 2
                        var v19 uint32 = 2
                        if (v17 || (426 >= genFn1())) {
                            for fi20 := 0; fi20 < 18; fi20++ {
                                var v21 uint = 528
                                v18 = 0
                                _ = fi20
                                _ = v21
                            }
                            for fi22 := 0; fi22 < 6; fi22++ {
                                var v23 uint32 = uint32(v18)
                                _ = fi22
                                _ = v23
                            }
                        } else {
                            v24 := GenS1{F0: (v3 / ((int8(v18)) | 1)), F1: uint64(v19), F2: v18}
                            _ = v24
                        }
                        _ = v18
                        _ = v19
                    } else {
                        v25 := GenS1{F0: v3, F1: 4294967295, F2: 234}
                        fmt.Println(v17, v3, v2)
                        for fi26 := 0; fi26 < 1; fi26++ {
                            v27 := GenS4{F0: v2}
                            v3 = v3
                            var v28 uint32 = 1
                            _ = fi26
                            _ = v27
                            _ = v28
                        }
                        _ = v25
                    }
                    var v29 int64 = -9223372036854775808
                    _ = v17
                    _ = v29
                } else {
                    var v30 uint16 = v9
                    for fi31 := 0; fi31 < 13; fi31++ {
                        var v32 uint8 = 255
                        switch 694 {
                        case -9223372036854775808:
                            switch (fi31 >> ((255) & 63)) {
                            case 42:
                                var v33 int32 = -729
                                _ = v33
                            default:
                                v1 = append(v1, 183)
                            }
                        case 42:
                            switch v2 {
                            case "foo bar":
                                for fi34 := 0; fi34 < 14; fi34++ {
                                    if false {
                                        v35 := GenS2{F0: 42, F1: -1}
                                        _ = v35
                                    } else {
                                        v36 := GenS5{F0: (-(v3)), F1: ((0 < -1) && true), F2: 227}
                                        var v37 int16 = 1
                                        v37 = 0
                                        var v38 float32 = (0.0 - -2.25)
                                        _ = v36
                                        _ = v37
                                        _ = v38
                                    }
                                    v32 = (v32 + v32)
                                    if len(v1) > 0 {
                                        v1[0] = 51
                                    }
                                    _ = fi34
                                }
                            default:
                                if false {
                                    var v39 float32 = 0.1
                                    v40 := GenS2{F0: v9, F1: 0}
                                    var v41 uint8 = (v32 ^ v32)
                                    _ = v39
                                    _ = v40
                                    _ = v41
                                } else {
                                    if len(v1) > 0 {
                                        v1[0] = 2147483648
                                    }
                                    if false {
                                        v2 = v2
                                        for fi42 := 0; fi42 < 10; fi42++ {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            _ = fi42
                                        }
                                    } else {
                                        var v43 int64 = 0
                                        var v44 int32 = 1
                                        _ = v43
                                        _ = v44
                                    }
                                }
                            }
                        default:
                            fmt.Println(v2)
                        }
                        _ = fi31
                        _ = v32
                    }
                    v30 = v9
                    _ = v30
                }
            case 42:
                switch 9223372036854775807 {
                case -1:
                    var v45 float64 = 1e10
                    _ = v45
                }
            case 128:
                switch 0 {
                case 359:
                    for fi46 := 0; fi46 < 8; fi46++ {
                        for fi47 := 0; fi47 < 14; fi47++ {
                            switch fi46 {
                            case 42:
                                fmt.Println(v3, v9, v2)
                            case 128:
                                v2 = v2
                            case -567:
                                if (!(uint32(v9) >= 2147483648)) {
                                    if false {
                                        switch fi47 {
                                        case 0:
                                            v1 = append(v1, 194)
                                        case 4294967295:
                                            if (-796 == 9223372036854775807) {
                                                v3 = v3
                                                v2 = v2
                                                v48 := GenS6{F0: (v9 >> ((65535) & 63)), F1: 9223372036854775807, F2: 2147483647}
                                                _ = v48
                                            }
                                        default:
                                            if true {
                                                _ = len(v1)
                                                _ = len(v1)
                                            } else {
                                                switch -457 {
                                                case 32767:
                                                    v49 := GenS1{F0: (v3 ^ (v3 >> ((65535) & 63))), F1: 4294967295, F2: uint8(v9)}
                                                    _ = v49
                                                case 9223372036854775807:
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                case 141:
                                                    v3 += (v3 % (((v3 / ((v3) | 1))) | 1))
                                                }
                                                v8 = (v8 - (v8 * ((v8 * float64(-2.25)) - float64((((-2.25 / (1.5 - (float32(v8) / 3.14159))) + -2.25) - 3.14159)))))
                                                if len(v1) > 0 {
                                                    v1[0] = 65535
                                                }
                                            }
                                        }
                                        var v50 uint8 = 42
                                        v51 := GenS7{F0: (v50 ^ 127), F1: (!(!(42 >= 298)))}
                                        _ = v50
                                        _ = v51
                                    } else {
                                        var v52 float32 = float32(float64(3.14159))
                                        v52 += v52
                                        if true {
                                            var v53 int16 = -32768
                                            _ = v53
                                        } else {
                                            var v54 bool = (true && (v2 > v2))
                                            v8 = v8
                                            var v55 int16 = -897
                                            _ = v54
                                            _ = v55
                                        }
                                        switch v2 {
                                        case "\n\t":
                                            var v56 uint = 2
                                            _ = v56
                                        default:
                                            if (!true) {
                                                fmt.Println(v8, v8, v3)
                                            }
                                        }
                                        _ = v52
                                    }
                                    v57 := GenS1{F0: (^(v3)), F1: 760, F2: 0}
                                    _ = v57
                                }
                                if (-9223372036854775808 == 1) {
                                    if len(v1) > 0 {
                                        v1[0] = 2147483647
                                    }
                                    if true {
                                        switch (fi47 >> ((1) & 63)) {
                                        case 2147483647:
                                            v58 := GenS7{F0: 42, F1: (0 <= genFn0(0))}
                                            _ = v58
                                        default:
                                            fmt.Println(v2, v2, v9)
                                        }
                                        v59 := GenS2{F0: (v9 + 1), F1: genFn0(1)}
                                        v60 := GenS8{F0: 2, F1: true}
                                        _ = v59
                                        _ = v60
                                    } else {
                                        v2 = v2
                                    }
                                    v61 := GenS7{F0: 127, F1: (((!false) || true) && true)}
                                    _ = v61
                                } else {
                                    if false {
                                        var v62 uint = 576
                                        var v63 uint64 = uint64(v9)
                                        v8 += v8
                                        _ = v62
                                        _ = v63
                                    }
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                            }
                            _ = fi47
                        }
                        switch fi46 {
                        case -1:
                            var v64 uint32 = 4294967295
                            _ = v64
                        case 1:
                            if len(v1) > 0 {
                                v1[0] = 42
                            }
                        case -9223372036854775808:
                            switch (fi46 >> ((4294967295) & 63)) {
                            case 664:
                                var v65 int64 = 9223372036854775807
                                _ = v65
                            default:
                                var v66 int64 = 530
                                _ = v66
                            }
                        default:
                            if true {
                                for fi67 := 0; fi67 < 3; fi67++ {
                                    var v68 int8 = v3
                                    if len(v1) > 0 {
                                        v1[0] = 0
                                    }
                                    v8 += (3.14159 - (v8 / v8))
                                    _ = fi67
                                    _ = v68
                                }
                                switch (fi46 ^ 428) {
                                case 32768:
                                    v2 = ("hello" + v2)
                                case 32767:
                                    v8 += ((float64(float32(v8)) + float64((0.0 / 1.0))) * v8)
                                }
                                if ((v2 + v2) <= "a") {
                                    var v69 uint32 = 42
                                    _ = v69
                                }
                            }
                        }
                        _ = fi46
                    }
                case 103:
                    if true {
                        if (997 == 42) {
                            if false {
                                var v70 int32 = 0
                                for fi71 := 0; fi71 < 1; fi71++ {
                                    var v72 uint = 1
                                    if (!(!(!(!false)))) {
                                        if (!false) {
                                            switch (((v2 + v2) + "foo bar") + "") {
                                            case "hello":
                                                v70 = (v70 ^ v70)
                                            case "foo bar":
                                                var v73 int = (fi71 ^ (fi71 % ((42) | 1)))
                                                _ = v73
                                            }
                                            for fi74 := 0; fi74 < 11; fi74++ {
                                                v9 -= (^(v9))
                                                _ = fi74
                                            }
                                        }
                                        v75 := GenS6{F0: 0, F1: 42, F2: v70}
                                        _ = v75
                                    } else {
                                        if (336 <= 42) {
                                            v8 -= -2.25
                                            var v76 float32 = (0.0 / 1.5)
                                            _ = v76
                                        }
                                        if (9223372036854775807 >= -25) {
                                            switch v2 {
                                            case "a":
                                                v77 := GenS7{F0: 2, F1: (v70 >= -180)}
                                                _ = v77
                                            case "foo bar":
                                                for fi78 := 0; fi78 < 1; fi78++ {
                                                    fmt.Println(fi78, v72, v9)
                                                    if (104 >= 238) {
                                                        if (0 > (v9 / ((879) | 1))) {
                                                            if ((!((1.5 == v8) && (uint64(v72) <= 590))) && (0.0 >= -2.25)) {
                                                                var v79 float32 = (1e10 + (0.0 * (-2.25 / 1.0)))
                                                                v9 += v9
                                                                v3 = v3
                                                                _ = v79
                                                            }
                                                        } else {
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                            if false {
                                                                v3 = -100
                                                            }
                                                            var v80 float32 = float32(-2.25)
                                                            _ = v80
                                                        }
                                                        var v81 int = (fi78 % ((0) | 1))
                                                        _ = v81
                                                    } else {
                                                        var v82 int8 = 42
                                                        v8 = float64(1.5)
                                                        _ = v82
                                                    }
                                                    v3 -= v3
                                                    v72 -= (^(v72))
                                                    _ = fi78
                                                }
                                            case "\n\t":
                                                if len(v1) > 0 {
                                                    v1[0] = 65535
                                                }
                                            default:
                                                switch v2 {
                                                case "hello":
                                                    v70 -= (v70 / ((-2147483648) | 1))
                                                case "":
                                                    fmt.Println(v70, v70)
                                                case "foo bar":
                                                    v1 = append(v1, 738)
                                                }
                                            }
                                            var v83 bool = (4294967295 >= 2)
                                            _ = v83
                                        }
                                    }
                                    _ = fi71
                                    _ = v72
                                }
                                _ = v70
                            } else {
                                var v84 uint16 = (v9 % (((^(v9))) | 1))
                                v85 := GenS6{F0: 2, F1: 626, F2: -1}
                                var v86 uint64 = 805
                                _ = v84
                                _ = v85
                                _ = v86
                            }
                            v87 := GenS9{F0: -1, F1: 65535, F2: 687}
                            _ = v87
                        } else {
                            v88 := GenS10{F0: 127, F1: (42 > -2147483648)}
                            for fi89 := 0; fi89 < 16; fi89++ {
                                v90 := GenS10{F0: (v3 << ((18446744073709551615) & 63)), F1: false}
                                _ = fi89
                                _ = v90
                            }
                            _ = v88
                        }
                        if true {
                            fmt.Println(v9, v9)
                            if (2 >= 42) {
                                v9 += 50
                            } else {
                                for fi91 := 0; fi91 < 18; fi91++ {
                                    if ((-1 > 42) || true) {
                                        for fi92 := 0; fi92 < 7; fi92++ {
                                            var v93 int32 = 42
                                            v93 -= v93
                                            v94 := GenS4{F0: ("" + ("" + "hello"))}
                                            _ = fi92
                                            _ = v93
                                            _ = v94
                                        }
                                        _ = len(v1)
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    }
                                    var v95 int64 = 42
                                    fmt.Println(fi91)
                                    _ = fi91
                                    _ = v95
                                }
                                if (!((v9 == 2) || false)) {
                                    for fi96 := 0; fi96 < 4; fi96++ {
                                        var v97 uint8 = 255
                                        var v98 int64 = 0
                                        switch v2 {
                                        case "hello":
                                            fmt.Println(fi96)
                                        default:
                                            var v99 int16 = 1
                                            _ = v99
                                        }
                                        _ = fi96
                                        _ = v97
                                        _ = v98
                                    }
                                    v100 := GenS1{F0: -46, F1: genFn1(), F2: 42}
                                    for fi101 := 0; fi101 < 17; fi101++ {
                                        v102 := GenS2{F0: 2, F1: 0}
                                        switch ("\n\t" + "foo bar") {
                                        case "":
                                            var v103 float32 = 0.1
                                            _ = v103
                                        case "foo bar":
                                            var v104 bool = true
                                            _ = v104
                                        default:
                                            var v105 int16 = 0
                                            _ = v105
                                        }
                                        _ = fi101
                                        _ = v102
                                    }
                                    _ = v100
                                }
                            }
                            v1 = append(v1, 1)
                        }
                        v8 = v8
                    } else {
                        if len(v1) > 0 {
                            v1[0] = 603
                        }
                        if true {
                            v1 = append(v1, 2147483648)
                            var v106 int8 = (-(v3))
                            if (true || true) {
                                switch v2 {
                                case "a":
                                    if (true || (v2 == (v2 + ((v2 + v2) + v2)))) {
                                        v106 = (v106 % ((v3) | 1))
                                    } else {
                                        for fi107 := 0; fi107 < 9; fi107++ {
                                            var v108 float64 = v8
                                            _ = fi107
                                            _ = v108
                                        }
                                        v9 = (v9 >> ((1) & 63))
                                    }
                                default:
                                    switch (("" + "\n\t") + ((v2 + (v2 + (("\n\t" + "") + v2))) + ("a" + v2))) {
                                    case "foo bar":
                                        if false {
                                            for fi109 := 0; fi109 < 2; fi109++ {
                                                var v110 uint32 = 2
                                                var v111 int8 = (v3 * -82)
                                                _ = fi109
                                                _ = v110
                                                _ = v111
                                            }
                                            v106 = v106
                                        }
                                    case "hello":
                                        _ = len(v1)
                                    default:
                                        switch v2 {
                                        case "":
                                            var v112 uint16 = (v9 - v9)
                                            _ = v112
                                        case "foo bar":
                                            if false {
                                                for fi113 := 0; fi113 < 9; fi113++ {
                                                    v9 += 2
                                                    if false {
                                                        if (true || false) {
                                                            if (!true) {
                                                                switch (fi113 - (fi113 / ((0) | 1))) {
                                                                case 32767:
                                                                    v1 = append(v1, 65535)
                                                                }
                                                            } else {
                                                                for fi114 := 0; fi114 < 9; fi114++ {
                                                                    switch v2 {
                                                                    case "\n\t":
                                                                        v1 = append(v1, 2)
                                                                    }
                                                                    switch 127 {
                                                                    case -64:
                                                                        var v115 float64 = ((1.5 * (float64((((1.5 + ((3.14159 * 1.5) - 1e10)) * (0.1 + 0.0)) / 0.1)) * v8)) * v8)
                                                                        _ = v115
                                                                    case 2147483647:
                                                                        var v116 string = (v2 + v2)
                                                                        _ = v116
                                                                    case 0:
                                                                        var v117 uint8 = 128
                                                                        _ = v117
                                                                    default:
                                                                        if ("a" < "a") {
                                                                            v118 := GenS10{F0: v106, F1: (!(false && (((fi113 | fi113) == fi114) || (42 > 4294967295))))}
                                                                            if false {
                                                                                fmt.Println(v106, fi114)
                                                                                v3 -= v106
                                                                            }
                                                                            _ = v118
                                                                        } else {
                                                                            v9 -= 2
                                                                            v8 = v8
                                                                        }
                                                                    }
                                                                    v119 := GenS11{F0: v106, F1: 998, F2: 835}
                                                                    _ = fi114
                                                                    _ = v119
                                                                }
                                                                v2 = v2
                                                                var v120 int32 = -1
                                                                _ = v120
                                                            }
                                                            v121 := GenS10{F0: (v106 * (v3 - v106)), F1: (v9 > (v9 >> ((376) & 63)))}
                                                            _ = v121
                                                        } else {
                                                            switch (fi113 << ((18446744073709551615) & 63)) {
                                                            case 127:
                                                                v2 = ((("\n\t" + v2) + v2) + ((v2 + "a") + "a"))
                                                            default:
                                                                if (400 < 18446744073709551615) {
                                                                    v106 -= v106
                                                                }
                                                            }
                                                            var v122 string = v2
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                            _ = v122
                                                        }
                                                    }
                                                    var v123 float64 = float64((-2.25 * 0.0))
                                                    _ = fi113
                                                    _ = v123
                                                }
                                            }
                                        }
                                    }
                                }
                                v1 = append(v1, 948)
                                var v124 float32 = (float32(v8) * float32(v8))
                                _ = v124
                            } else {
                                if len(v1) > 0 {
                                    v1[0] = 4294967295
                                }
                            }
                            _ = v106
                        }
                        var v125 uint32 = 460
                        _ = v125
                    }
                case 32767:
                    var v126 int16 = -1
                    _ = v126
                default:
                    v127 := GenS6{F0: 1, F1: -1, F2: 1}
                    _ = v127
                }
            default:
                var v128 uint16 = (v9 + (v9 >> ((1) & 63)))
                _ = v128
            }
            v129 := GenS12{F0: (true || (-2147483648 >= -84))}
            _ = v129
        }
    }
    v3 = v3
    for fi130 := 0; fi130 < 11; fi130++ {
        for fi131 := 0; fi131 < 16; fi131++ {
            v1 = append(v1, 916)
            fmt.Println(v2)
            _ = fi131
        }
        v2 = "foo bar"
        _ = fi130
    }
    v8 = (v8 * float64(1.5))
    for fi132 := 0; fi132 < 13; fi132++ {
        if true {
            v133 := GenS1{F0: v3, F1: genFn1(), F2: 1}
            var v134 float64 = v8
            _ = v133
            _ = v134
        }
        fmt.Println(v8, v2, v3)
        var v135 int = fi132
        _ = fi132
        _ = v135
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
