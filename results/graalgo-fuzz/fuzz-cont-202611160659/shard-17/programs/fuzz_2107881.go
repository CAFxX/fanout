package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 uint16; F1 uint32; F2 uint }
type GenS4 struct { F0 uint16; F1 uint }
type GenS5 struct { F0 uint; F1 uint; F2 int8 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int8; F1 uint16 }
type GenS11 struct { F0 bool }
func genFn0(p0 uint, p1 int, p2 uint64) uint16 {
    p0 -= p0
    var v5 int = (p1 * (p1 / (((p1 ^ (p1 + 32767))) | 1)))
    p2 = p2
    v5 = 87
    v6 := []int{(p1 >> ((1) & 63)), 229}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 255
}
func genFn1(p0 int16, p1 uint64, p2 int16) uint64 {
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 uint32 = 947
        _ = fi7
        _ = v8
    }
    if true {
        switch int(p2) {
        case 128:
            v9 := GenS1{F0: -1}
            _ = v9
        default:
            var v10 string = "foo bar"
            _ = v10
        }
        v11 := GenS1{F0: 0}
        p2 = (^(p2))
        _ = v11
    }
    p0 += (p0 % ((p0) | 1))
    for fi12 := 0; fi12 < 2; fi12++ {
        p0 = 1000
        v13 := GenS2{F0: 4294967295, F1: 416}
        p1 += p1
        _ = fi12
        _ = v13
    }
    v14 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return p1
}

func main() {
    var v1 uint8 = 42
    v2 := []int{-2147483648, 127}
    var v3 uint16 = 1
    v4 := make(map[uint32]bool)
    v1 += (^(v1))
    switch "a" {
    case "foo bar":
        if (v1 <= v1) {
            if (3 != 431) {
                v2 = append(v2, 42)
            } else {
                v3 += (v3 << ((435) & 63))
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            switch -633 {
            case 127:
                v3 -= (^(v3))
            }
        }
    case "a":
        var v15 float32 = -2.25
        _ = v15
    default:
        v16 := GenS3{F0: 564, F1: 2147483648, F2: 337}
        _ = v16
    }
    delete(v4, 81)
    fmt.Println(v1, v1)
    for fi17 := 0; fi17 < 9; fi17++ {
        v4[95] = (-2147483648 > 36)
        _ = fi17
    }
    if false {
        switch "" {
        case "":
            var v18 uint32 = 0
            _ = v18
        case "hello":
            v19 := v4[42]
            _ = v19
        case "\n\t":
            v20 := GenS3{F0: v3, F1: 2147483648, F2: 26}
            _ = v20
        default:
            v21 := GenS1{F0: 100}
            _ = v21
        }
        var v22 float64 = 1.5
        if (((2 != 255) && (42 != -500)) && (true && false)) {
            var v23 int64 = -9223372036854775808
            _ = v23
        } else {
            var v24 uint16 = (v3 % (((^(v3))) | 1))
            if true {
                if (0.0 >= 0.1) {
                    fmt.Println(v3)
                    v25 := GenS1{F0: 127}
                    v26 := GenS3{F0: v3, F1: 2147483647, F2: 65535}
                    _ = v25
                    _ = v26
                } else {
                    var v27 int32 = 42
                    _ = v27
                }
                if true {
                    switch -1 {
                    case 419:
                        _ = len(v4)
                    }
                } else {
                    switch "" {
                    case "":
                        v4[4294967295] = (true || (-129 < 32768))
                    }
                    if (false && (42 > 42)) {
                        switch "hello" {
                        case "foo bar":
                            var v28 uint = 0
                            _ = v28
                        case "a":
                            if true {
                                if false {
                                    for fi29 := 0; fi29 < 8; fi29++ {
                                        for fi30 := 0; fi30 < 14; fi30++ {
                                            if false {
                                                v1 -= (^(v1))
                                                v4[0] = ((v22 * (0.1 + v22)) >= v22)
                                            } else {
                                                fmt.Println(fi29, fi30, fi30)
                                            }
                                            _ = fi30
                                        }
                                        var v31 string = "foo bar"
                                        v22 -= (-2.25 / (v22 / -2.25))
                                        _ = fi29
                                        _ = v31
                                    }
                                    for fi32 := 0; fi32 < 9; fi32++ {
                                        v33 := GenS3{F0: (v24 - (^(v3))), F1: uint32(v1), F2: 255}
                                        if (false || false) {
                                            var v34 uint8 = v1
                                            v34 -= (^(v34))
                                            _ = v34
                                        }
                                        var v35 uint16 = 751
                                        _ = fi32
                                        _ = v33
                                        _ = v35
                                    }
                                }
                                v3 += (v3 >> ((678) & 63))
                                for fi36 := 0; fi36 < 13; fi36++ {
                                    if ((32767 != int16(v1)) || true) {
                                        v2 = append(v2, 32767)
                                    }
                                    _ = fi36
                                }
                            } else {
                                if (365 <= 4294967295) {
                                    v24 -= 1
                                    _ = len(v2)
                                } else {
                                    v37 := GenS4{F0: genFn0(0, 0, 2), F1: 0}
                                    _ = v37
                                }
                                var v38 string = ("" + "")
                                _ = v38
                            }
                        }
                    }
                    v39 := GenS5{F0: 2, F1: 18446744073709551615, F2: 127}
                    _ = v39
                }
            }
            v40 := GenS3{F0: v24, F1: 940, F2: 378}
            _ = v24
            _ = v40
        }
        _ = v22
    } else {
        if false {
            fmt.Println(v1, v3)
            for fi41 := 0; fi41 < 11; fi41++ {
                v42 := GenS1{F0: -1}
                var v43 int = (fi41 >> ((42) & 63))
                _ = fi41
                _ = v42
                _ = v43
            }
            for fi44 := 0; fi44 < 5; fi44++ {
                switch fi44 {
                case -1:
                    var v45 int8 = 100
                    _ = v45
                case 0:
                    var v46 uint16 = v3
                    _ = v46
                default:
                    v3 += (v3 % ((255) | 1))
                }
                v1 = 0
                _ = fi44
            }
        } else {
            if (true || true) {
                for fi47 := 0; fi47 < 16; fi47++ {
                    v4[426] = false
                    _ = fi47
                }
            } else {
                _ = len(v4)
                var v48 string = ("" + (("" + "foo bar") + ("a" + (("" + ("a" + ("" + "\n\t"))) + "\n\t"))))
                delete(v4, 595)
                _ = v48
            }
        }
    }
    if ((-953 != 0) || ("hello" != ("foo bar" + ""))) {
        v1 = v1
        v2 = append(v2, 2147483647)
    } else {
        for fi49 := 0; fi49 < 2; fi49++ {
            switch "hello" {
            case "hello":
                var v50 float32 = 0.0
                _ = v50
            }
            _ = fi49
        }
        if false {
            if true {
                v51 := GenS4{F0: v3, F1: 653}
                _ = v51
            }
            v52 := GenS6{F0: v1}
            var v53 uint32 = 4294967295
            _ = v52
            _ = v53
        }
        if ((!(true && false)) && (!false)) {
            var v54 int64 = 42
            v3 = (^(v3))
            for fi55 := 0; fi55 < 2; fi55++ {
                for fi56 := 0; fi56 < 1; fi56++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v57 uint8 = v1
                    v1 -= (v57 ^ 42)
                    _ = fi56
                    _ = v57
                }
                v1 = 42
                var v58 uint64 = 523
                _ = fi55
                _ = v58
            }
            _ = v54
        } else {
            if true {
                var v59 float64 = 3.14159
                for fi60 := 0; fi60 < 4; fi60++ {
                    _ = len(v4)
                    fmt.Println(v3, v1)
                    var v61 uint16 = (v3 / ((v3) | 1))
                    _ = fi60
                    _ = v61
                }
                v59 = ((float64((0.1 + (3.14159 + (float32(v59) / float32(v59))))) - 1e10) * v59)
                _ = v59
            } else {
                var v62 int64 = -1
                _ = v62
            }
            var v63 bool = true
            _ = v63
        }
    }
    if true {
        if false {
            var v64 uint16 = 0
            v1 += 42
            if (-2147483648 != 380) {
                if (((true || (true || true)) && ("hello" >= "a")) && (v64 > 0)) {
                    switch -2147483648 {
                    case 374:
                        switch ("" + "hello") {
                        case "foo bar":
                            if true {
                                v1 = (v1 ^ 0)
                                v3 = (^(v3))
                                if false {
                                    if false {
                                        var v65 string = "hello"
                                        _ = v65
                                    } else {
                                        v66 := GenS2{F0: 18446744073709551615, F1: (v3 / (((v3 | uint16(v1))) | 1))}
                                        _ = v66
                                    }
                                }
                            }
                        }
                    case 2147483647:
                        v3 = 698
                    }
                    for fi67 := 0; fi67 < 2; fi67++ {
                        v68 := GenS3{F0: v3, F1: 2147483648, F2: 251}
                        var v69 int16 = -338
                        _ = fi67
                        _ = v68
                        _ = v69
                    }
                    for fi70 := 0; fi70 < 18; fi70++ {
                        v64 = (v64 << ((350) & 63))
                        v3 -= (v3 / (((v64 - v64)) | 1))
                        _ = fi70
                    }
                } else {
                    for fi71 := 0; fi71 < 8; fi71++ {
                        if false {
                            var v72 int8 = -1
                            var v73 float32 = 1e10
                            _ = v72
                            _ = v73
                        }
                        _ = fi71
                    }
                }
                switch 42 {
                case 339:
                    switch 632 {
                    case -997:
                        var v74 uint32 = 974
                        _ = v74
                    }
                case 2147483647:
                    if false {
                        var v75 float64 = 1.5
                        if false {
                            fmt.Println(v3, v75)
                            var v76 uint16 = v3
                            _ = v76
                        }
                        v64 -= (v64 << ((4294967295) & 63))
                        _ = v75
                    } else {
                        var v77 string = ("foo bar" + ("foo bar" + "hello"))
                        _ = v77
                    }
                case -944:
                    var v78 bool = false
                    _ = v78
                }
                _ = len(v4)
            } else {
                var v79 uint32 = 4294967295
                _ = v79
            }
            _ = v64
        } else {
            for fi80 := 0; fi80 < 2; fi80++ {
                switch fi80 {
                case 613:
                    _ = len(v2)
                case 4294967295:
                    fmt.Println(v1)
                }
                for fi81 := 0; fi81 < 4; fi81++ {
                    switch fi81 {
                    case 42:
                        fmt.Println(v1, v1)
                    case -514:
                        var v82 float64 = -2.25
                        fmt.Println(v3)
                        _ = v82
                    case -2147483648:
                        if ((-1 <= 32767) || false) {
                            fmt.Println(v3, fi80, fi81)
                            v83 := GenS3{F0: v3, F1: 1, F2: 340}
                            _ = v83
                        }
                    }
                    _ = fi81
                }
                _ = fi80
            }
            var v84 uint64 = genFn1(1000, 1, 0)
            for fi85 := 0; fi85 < 2; fi85++ {
                var v86 int32 = -375
                if len(v2) > 0 {
                    v2[0] = fi85
                }
                v4[2147483648] = false
                _ = fi85
                _ = v86
            }
            _ = v84
        }
        v87 := v4[4294967295]
        if true {
            if (!true) {
                if ("a" <= "\n\t") {
                    for fi88 := 0; fi88 < 19; fi88++ {
                        for fi89 := 0; fi89 < 9; fi89++ {
                            var v90 float32 = 1e10
                            var v91 uint8 = (v1 << ((65535) & 63))
                            _ = fi89
                            _ = v90
                            _ = v91
                        }
                        v1 = (v1 % ((v1) | 1))
                        v3 += genFn0(385, (^(fi88)), 2)
                        _ = fi88
                    }
                } else {
                    for fi92 := 0; fi92 < 4; fi92++ {
                        fmt.Println(v87)
                        _ = fi92
                    }
                    for fi93 := 0; fi93 < 2; fi93++ {
                        for fi94 := 0; fi94 < 10; fi94++ {
                            v95 := GenS6{F0: 42}
                            v4[2] = v87
                            if v87 {
                                var v96 uint64 = 2
                                v97 := GenS5{F0: 2, F1: 4294967295, F2: 1}
                                _ = v96
                                _ = v97
                            } else {
                                var v98 bool = (!(!(!(1 < -128))))
                                v99 := GenS7{F0: -32768}
                                _ = v98
                                _ = v99
                            }
                            _ = fi94
                            _ = v95
                        }
                        for fi100 := 0; fi100 < 3; fi100++ {
                            switch fi100 {
                            case -602:
                                v87 = true
                            case 9223372036854775807:
                                if (4294967295 != fi100) {
                                    v3 = 42
                                    var v101 int64 = 9223372036854775807
                                    _ = v101
                                } else {
                                    fmt.Println(fi100, v3, v3)
                                }
                            case 0:
                                fmt.Println(v87, v3, v1)
                            }
                            v102 := GenS1{F0: 100}
                            _ = fi100
                            _ = v102
                        }
                        if v87 {
                            v103 := GenS5{F0: 0, F1: 42, F2: 100}
                            _ = v103
                        } else {
                            switch 127 {
                            case 596:
                                v104 := GenS1{F0: 1}
                                _ = v104
                            case 713:
                                v1 += (v1 | (v1 ^ (v1 ^ v1)))
                            case 834:
                                var v105 int = -936
                                _ = v105
                            }
                            var v106 float64 = float64(1.5)
                            _ = v106
                        }
                        _ = fi93
                    }
                    _ = len(v4)
                }
                if true {
                    v107 := GenS2{F0: 255, F1: (v3 * (v3 - (v3 / ((288) | 1))))}
                    var v108 uint8 = 42
                    _ = v107
                    _ = v108
                } else {
                    for fi109 := 0; fi109 < 7; fi109++ {
                        switch "foo bar" {
                        case "\n\t":
                            if true {
                                v110 := GenS8{F0: 942}
                                _ = v110
                            }
                        case "foo bar":
                            if (42 > 65535) {
                                for fi111 := 0; fi111 < 11; fi111++ {
                                    v1 += v1
                                    _ = fi111
                                }
                                v87 = v87
                            }
                        default:
                            _ = len(v2)
                        }
                        if (255 > 65535) {
                            for fi112 := 0; fi112 < 20; fi112++ {
                                var v113 uint8 = (^(v1))
                                var v114 int64 = 9223372036854775807
                                if true {
                                    var v115 bool = (!(((v3 | (v3 * (^(v3)))) >= v3) && (2 != 2)))
                                    if (!((!true) || (-1 > -128))) {
                                        if (698 <= 18446744073709551615) {
                                            fmt.Println(fi109, v114)
                                            var v116 uint64 = genFn1(32767, 841, 1)
                                            _ = v116
                                        }
                                        var v117 float32 = float32(0.1)
                                        v3 += (^(v3))
                                        _ = v117
                                    } else {
                                        fmt.Println(v114)
                                    }
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = v115
                                } else {
                                    var v118 int16 = 356
                                    v119 := GenS3{F0: 32767, F1: 42, F2: 255}
                                    v2 = append(v2, fi112)
                                    _ = v118
                                    _ = v119
                                }
                                switch "" {
                                case "":
                                    v120 := GenS4{F0: genFn0(0, fi109, 4294967295), F1: 0}
                                    _ = v120
                                case "\n\t":
                                    _ = len(v4)
                                case "foo bar":
                                    if (!true) {
                                        var v121 bool = v87
                                        v1 = 42
                                        if v121 {
                                            fmt.Println(v87, fi109, v87)
                                        }
                                        _ = v121
                                    } else {
                                        if v87 {
                                            v122 := GenS9{F0: 127}
                                            _ = v122
                                        }
                                        delete(v4, 4294967295)
                                    }
                                }
                                _ = fi112
                                _ = v113
                                _ = v114
                            }
                            var v123 uint16 = 2
                            fmt.Println(v87)
                            _ = v123
                        } else {
                            var v124 int16 = 0
                            _ = v124
                        }
                        v1 -= 42
                        _ = fi109
                    }
                    v3 = 65535
                    for fi125 := 0; fi125 < 14; fi125++ {
                        v1 = 2
                        _ = fi125
                    }
                }
            } else {
                var v126 int64 = 42
                _ = v126
            }
            v3 -= v3
            var v127 uint8 = (v1 & 128)
            _ = v127
        }
        _ = v87
    } else {
        v128 := GenS9{F0: 485}
        _ = v128
    }
    v1 = v1
    v4[uint32(v1)] = (false && false)
    var v129 int32 = 584
    var v130 uint16 = 62
    var v131 uint16 = (v3 ^ v3)
    v3 = (v3 + 622)
    v132 := GenS10{F0: 0, F1: 900}
    if true {
        var v133 string = ("foo bar" + "\n\t")
        switch -2147483648 {
        case -649:
            if (false || false) {
                if true {
                    _ = len(v4)
                    v134 := GenS1{F0: 1}
                    switch 67 {
                    case -1:
                        v135 := GenS11{F0: (0.0 == float64(float32(1e10)))}
                        _ = v135
                    case -340:
                        v136 := v4[376]
                        _ = v136
                    default:
                        v137 := v4[2147483647]
                        _ = v137
                    }
                    _ = v134
                } else {
                    v138 := GenS5{F0: 255, F1: 18446744073709551615, F2: -17}
                    _ = v138
                }
            } else {
                _ = len(v2)
                v1 = 42
                var v139 float64 = (0.1 * (3.14159 + 3.14159))
                _ = v139
            }
        case 2147483647:
            for fi140 := 0; fi140 < 14; fi140++ {
                for fi141 := 0; fi141 < 15; fi141++ {
                    v129 += 1
                    var v142 uint16 = genFn0(42, (fi140 - (fi140 & (fi140 >> ((255) & 63)))), 18446744073709551615)
                    _ = fi141
                    _ = v142
                }
                if (42 != 9223372036854775807) {
                    var v143 uint32 = 688
                    for fi144 := 0; fi144 < 5; fi144++ {
                        v1 += v1
                        v130 = (^(v130))
                        switch "" {
                        case "hello":
                            v1 -= (v1 % (((v1 % (((v1 ^ 42)) | 1))) | 1))
                        case "\n\t":
                            fmt.Println(fi144, v1)
                        }
                        _ = fi144
                    }
                    var v145 string = (v133 + "\n\t")
                    _ = v143
                    _ = v145
                }
                for fi146 := 0; fi146 < 14; fi146++ {
                    v129 = v129
                    _ = fi146
                }
                _ = fi140
            }
        case 220:
            var v147 uint = 65535
            _ = v147
        default:
            fmt.Println(v131, v133, v1)
        }
        if false {
            v2 = append(v2, 127)
        }
        _ = v133
    } else {
        v148 := GenS3{F0: (v130 - (v131 & 0)), F1: 2, F2: 255}
        if false {
            v130 += v131
            v130 = (v131 - v131)
        }
        var v149 string = "a"
        _ = v148
        _ = v149
    }
    var v150 int8 = 42
    v1 -= (v1 % ((139) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v129:", v129)
    fmt.Println("v130:", v130)
    fmt.Println("v131:", v131)
    fmt.Println("v150:", v150)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v129
    _ = v130
    _ = v131
    _ = v132
    _ = v150
    _ = v2
    _ = v3
    _ = v4
}
