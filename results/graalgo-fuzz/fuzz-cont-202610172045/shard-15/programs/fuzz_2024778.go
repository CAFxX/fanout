package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 uint32; F1 int; F2 uint }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 bool; F1 int32; F2 int32 }
type GenS6 struct { F0 uint16; F1 int }
type GenS7 struct { F0 bool; F1 uint16 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int64; F1 int }
type GenS10 struct { F0 int16; F1 bool; F2 uint16 }
type GenS11 struct { F0 int32; F1 uint8 }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 uint32; F1 uint8; F2 uint8 }

func main() {
    var v1 float32 = 0.1
    v2 := make(map[uint32]uint32)
    var v3 int32 = 42
    var v4 int32 = v3
    var v5 string = "\n\t"
    v1 -= 1e10
    var v6 float64 = 0.1
    var v7 int16 = 180
    if (534 >= -2147483648) {
        v8 := GenS1{F0: 32, F1: 1}
        for fi9 := 0; fi9 < 13; fi9++ {
            if false {
                for fi10 := 0; fi10 < 20; fi10++ {
                    var v11 int8 = -128
                    v12 := GenS1{F0: (v11 / (((v11 % ((108) | 1))) | 1)), F1: 2}
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            } else {
                switch "hello" {
                case "a":
                    v13 := GenS1{F0: -128, F1: 111}
                    _ = v13
                default:
                    fmt.Println(v5, v1)
                }
                switch fi9 {
                case -9223372036854775808:
                    for fi14 := 0; fi14 < 20; fi14++ {
                        var v15 int = fi14
                        _ = fi14
                        _ = v15
                    }
                }
            }
            v2[663] = 795
            v16 := []int{9223372036854775807, fi9, -43}
            _ = fi9
            _ = v16
        }
        v6 += v6
        _ = v8
    } else {
        if (2 <= 2) {
            fmt.Println(v3, v7)
            v5 = "\n\t"
        }
    }
    if (42 == 2) {
        var v17 int64 = 9223372036854775807
        _ = v17
    } else {
        v3 -= 2147483647
        var v18 uint = 2
        v5 = (v5 + v5)
        _ = v18
    }
    v19 := GenS2{F0: 42, F1: 1}
    v5 = v5
    var v20 string = v5
    if true {
        v21 := GenS1{F0: 1, F1: 185}
        _ = v21
    }
    if (1 > 667) {
        for fi22 := 0; fi22 < 7; fi22++ {
            v20 = (v20 + ("a" + (("hello" + ("a" + v5)) + v20)))
            var v23 uint = 255
            switch "\n\t" {
            case "foo bar":
                for fi24 := 0; fi24 < 1; fi24++ {
                    if false {
                        if true {
                            var v25 bool = ((true || false) && true)
                            _ = v25
                        } else {
                            fmt.Println(v7)
                            var v26 float32 = ((3.14159 * v1) / v1)
                            if (!(v5 >= ("hello" + v5))) {
                                if ((^(v3)) == v3) {
                                    if true {
                                        v27 := GenS3{F0: 977, F1: fi22, F2: (v23 << (((v23 / ((v23) | 1))) & 63))}
                                        _ = v27
                                    } else {
                                        v28 := GenS4{F0: 2147483647}
                                        v29 := v2[532]
                                        var v30 uint = v23
                                        _ = v28
                                        _ = v29
                                        _ = v30
                                    }
                                    var v31 uint16 = 42
                                    _ = v31
                                } else {
                                    v32 := GenS2{F0: 0, F1: 138}
                                    v26 += v1
                                    var v33 uint16 = 2
                                    var v34 uint64 = 42
                                    v5 = "\n\t"
                                    _ = v32
                                    _ = v33
                                    _ = v34
                                }
                            } else {
                                var v35 float32 = v1
                                var v36 int = -1
                                switch v20 {
                                case "hello":
                                    if (2 <= 364) {
                                        fmt.Println(v4, v3)
                                        var v37 int = 128
                                        if (!(!true)) {
                                            v38 := GenS5{F0: (!false), F1: 586, F2: v3}
                                            v39 := GenS6{F0: 791, F1: (v37 + v36)}
                                            _ = v38
                                            _ = v39
                                        } else {
                                            var v40 int8 = 127
                                            v23 = 393
                                            _ = v40
                                        }
                                        _ = v37
                                    }
                                case "":
                                    switch v36 {
                                    case -1:
                                        var v41 string = v20
                                        _ = v41
                                    case 128:
                                        v26 = v35
                                    default:
                                        v26 = (v26 * (3.14159 / v26))
                                    }
                                case "a":
                                    switch "a" {
                                    case "\n\t":
                                        var v42 uint32 = 2147483648
                                        _ = v42
                                    case "foo bar":
                                        v43 := GenS2{F0: 24, F1: 36}
                                        _ = v43
                                    }
                                }
                                _ = v35
                                _ = v36
                            }
                            var v44 int64 = -911
                            v45 := GenS4{F0: 437}
                            _ = v26
                            _ = v44
                            _ = v45
                        }
                    }
                    _ = fi24
                }
            case "":
                for fi46 := 0; fi46 < 2; fi46++ {
                    v47 := GenS7{F0: true, F1: 2}
                    _ = fi46
                    _ = v47
                }
            case "a":
                if false {
                    if false {
                        for fi48 := 0; fi48 < 10; fi48++ {
                            var v49 uint = 255
                            fmt.Println(v3, v3, v5)
                            var v50 int8 = 100
                            _ = fi48
                            _ = v49
                            _ = v50
                        }
                    } else {
                        var v51 bool = (-2.25 < (v1 * v1))
                        _ = v51
                    }
                    switch "" {
                    case "\n\t":
                        v1 -= ((0.1 + float32(1e10)) + (0.1 + 1.5))
                    case "a":
                        var v52 uint32 = 4294967295
                        _ = v52
                    }
                } else {
                    fmt.Println(v6, v20)
                    var v53 uint64 = 4294967295
                    _ = v53
                }
            }
            _ = fi22
            _ = v23
        }
        switch 519 {
        case 32768:
            v20 = v20
        case 128:
            switch 32768 {
            case 32767:
                v4 -= (v3 >> ((4294967295) & 63))
            case 128:
                v54 := []int{128, 127, -1, 1}
                _ = v54
            }
        case 32767:
            var v55 int8 = 1
            _ = v55
        default:
            v56 := GenS1{F0: -128, F1: 25}
            _ = v56
        }
        if true {
            if (!false) {
                if (v7 >= v7) {
                    v6 += 3.14159
                    for fi57 := 0; fi57 < 6; fi57++ {
                        delete(v2, 2147483648)
                        v58 := []int{(fi57 % (((fi57 ^ (-(fi57)))) | 1)), 2147483647, 127}
                        v6 = v6
                        _ = fi57
                        _ = v58
                    }
                } else {
                    if false {
                        var v59 int = -9223372036854775808
                        for fi60 := 0; fi60 < 5; fi60++ {
                            v61 := v2[907]
                            var v62 int16 = int16(v3)
                            _ = fi60
                            _ = v61
                            _ = v62
                        }
                        v63 := GenS4{F0: 1}
                        _ = v59
                        _ = v63
                    }
                    if true {
                        v64 := GenS8{F0: 65535}
                        _ = v64
                    } else {
                        for fi65 := 0; fi65 < 4; fi65++ {
                            v1 = v1
                            var v66 string = v5
                            _ = fi65
                            _ = v66
                        }
                        v20 = v20
                    }
                    if (!(true || true)) {
                        for fi67 := 0; fi67 < 13; fi67++ {
                            switch 0 {
                            case 0:
                                for fi68 := 0; fi68 < 1; fi68++ {
                                    fmt.Println(v4)
                                    fmt.Println(v20)
                                    _ = fi68
                                }
                            case 482:
                                if true {
                                    for fi69 := 0; fi69 < 11; fi69++ {
                                        if false {
                                            v20 = ""
                                            switch (fi67 % ((fi67) | 1)) {
                                            case 1:
                                                if (!((fi69 - fi67) != (fi69 & (fi67 * (fi69 + fi69))))) {
                                                    var v70 int16 = (v7 ^ (^(v7)))
                                                    if true {
                                                        var v71 int64 = 1
                                                        _ = v71
                                                    } else {
                                                        v7 = (v70 | 558)
                                                        v6 += ((1e10 / float64((float32(v6) - float32((0.1 / 1.5))))) + -2.25)
                                                    }
                                                    v72 := GenS2{F0: 127, F1: 0}
                                                    _ = v70
                                                    _ = v72
                                                } else {
                                                    delete(v2, 4294967295)
                                                    delete(v2, 42)
                                                    v6 = (v6 / (v6 + 0.0))
                                                }
                                                v2[2] = 1
                                                var v73 int64 = 0
                                                v5 = v5
                                                _ = v73
                                            case 127:
                                                if true {
                                                    var v74 uint8 = 176
                                                    fmt.Println(v4)
                                                    if false {
                                                        var v75 uint = 42
                                                        v76 := GenS2{F0: 255, F1: v74}
                                                        v77 := GenS7{F0: false, F1: 1}
                                                        _ = v75
                                                        _ = v76
                                                        _ = v77
                                                    } else {
                                                        v78 := []int{fi67, 128, 658}
                                                        var v79 uint32 = 1
                                                        v1 += v1
                                                        _ = v78
                                                        _ = v79
                                                    }
                                                    _ = v74
                                                } else {
                                                    v5 = "hello"
                                                    v20 = ("" + "a")
                                                    v4 = (^(v3))
                                                }
                                                v80 := GenS9{F0: 0, F1: -143}
                                                v20 = v20
                                                if true {
                                                    v2[42] = 474
                                                    v7 += (v7 | (v7 + -32768))
                                                } else {
                                                    fmt.Println(fi69, v1, v5)
                                                    fmt.Println(v3)
                                                    var v81 bool = ("\n\t" >= "foo bar")
                                                    _ = v81
                                                }
                                                _ = len(v2)
                                                var v82 int16 = v7
                                                _ = v80
                                                _ = v82
                                            }
                                            if false {
                                                var v83 uint16 = 54
                                                _ = v83
                                            }
                                        }
                                        var v84 int16 = (v7 - (v7 - v7))
                                        _ = fi69
                                        _ = v84
                                    }
                                    var v85 uint = 4294967295
                                    if false {
                                        if ((false || false) || false) {
                                            v86 := []int{}
                                            _ = v86
                                        } else {
                                            _ = len(v2)
                                        }
                                        if (v4 < (v4 >> (((v85 & 449)) & 63))) {
                                            v87 := v2[919]
                                            switch 0 {
                                            case -9223372036854775808:
                                                v88 := []int{9223372036854775807, 128}
                                                _ = v88
                                            case 641:
                                                v89 := GenS2{F0: 2, F1: 128}
                                                _ = v89
                                            case 32768:
                                                if true {
                                                    v90 := GenS4{F0: 42}
                                                    var v91 float64 = v6
                                                    _ = v90
                                                    _ = v91
                                                } else {
                                                    var v92 int32 = v4
                                                    var v93 int32 = v92
                                                    _ = v92
                                                    _ = v93
                                                }
                                            }
                                            if (!(!true)) {
                                                switch (^(fi67)) {
                                                case 878:
                                                    v4 = (v3 - 42)
                                                default:
                                                    if ((v3 + -1) > v3) {
                                                        if (v85 <= (v85 - v85)) {
                                                            v6 -= -2.25
                                                            v94 := GenS6{F0: 1, F1: fi67}
                                                            _ = v94
                                                        } else {
                                                            for fi95 := 0; fi95 < 16; fi95++ {
                                                                delete(v2, 2147483647)
                                                                v96 := GenS10{F0: (^(v7)), F1: (127 < 42), F2: uint16(v85)}
                                                                v97 := []int{}
                                                                _ = fi95
                                                                _ = v96
                                                                _ = v97
                                                            }
                                                            v98 := []int{(^(fi67)), (fi67 << ((1) & 63)), (fi67 - 1)}
                                                            fmt.Println(v87, v5, v20)
                                                            _ = v98
                                                        }
                                                        var v99 uint = 423
                                                        _ = v99
                                                    }
                                                }
                                            } else {
                                                v3 += (v3 >> ((v85) & 63))
                                                for fi100 := 0; fi100 < 8; fi100++ {
                                                    switch fi100 {
                                                    case 127:
                                                        v101 := GenS11{F0: (v3 | 283), F1: 255}
                                                        _ = v101
                                                    case -644:
                                                        var v102 uint32 = v87
                                                        _ = v102
                                                    case -937:
                                                        v3 -= v4
                                                        if false {
                                                            var v103 uint = (v85 / ((v85) | 1))
                                                            v104 := GenS12{F0: 4294967295}
                                                            var v105 uint = (v103 ^ v85)
                                                            _ = v103
                                                            _ = v104
                                                            _ = v105
                                                        } else {
                                                            switch -9223372036854775808 {
                                                            case -420:
                                                                v87 = 42
                                                            default:
                                                                if true {
                                                                    var v106 int64 = -479
                                                                    var v107 uint64 = 1
                                                                    _ = v106
                                                                    _ = v107
                                                                }
                                                            }
                                                            v108 := GenS10{F0: 42, F1: true, F2: uint16(v87)}
                                                            var v109 uint64 = 509
                                                            _ = v108
                                                            _ = v109
                                                        }
                                                        v85 = v85
                                                    }
                                                    var v110 uint16 = 42
                                                    _ = fi100
                                                    _ = v110
                                                }
                                                var v111 float64 = (float64(v1) + 3.14159)
                                                _ = v111
                                            }
                                            _ = v87
                                        } else {
                                            v112 := GenS13{F0: 1, F1: 127, F2: 128}
                                            _ = v112
                                        }
                                    }
                                    _ = v85
                                }
                            case 2147483647:
                                v2[2147483648] = 2
                            }
                            v3 = (v4 ^ -1)
                            for fi113 := 0; fi113 < 4; fi113++ {
                                switch "foo bar" {
                                case "a":
                                    v1 = v1
                                case "\n\t":
                                    var v114 uint16 = 1
                                    _ = v114
                                case "":
                                    if ((((v1 / (v1 + (v1 + 1.5))) / v1) == (v1 - (v1 - v1))) && (573 >= 65535)) {
                                        v115 := GenS13{F0: 65535, F1: 25, F2: 191}
                                        _ = v115
                                    } else {
                                        fmt.Println(fi113, v3)
                                    }
                                }
                                _ = fi113
                            }
                            _ = fi67
                        }
                        v116 := GenS13{F0: 64, F1: 127, F2: 0}
                        _ = v116
                    }
                }
            } else {
                v117 := GenS11{F0: v4, F1: 65}
                _ = v117
            }
            var v118 string = ""
            _ = v118
        }
    } else {
        if true {
            var v119 uint8 = 8
            if false {
                if (v119 > 255) {
                    v120 := GenS5{F0: (!true), F1: v4, F2: v3}
                    _ = v120
                }
                var v121 int16 = (v7 * v7)
                v122 := GenS2{F0: v119, F1: (v119 * v119)}
                _ = v121
                _ = v122
            } else {
                v1 -= 1.5
            }
            switch v20 {
            case "foo bar":
                for fi123 := 0; fi123 < 18; fi123++ {
                    v124 := GenS3{F0: 2147483647, F1: fi123, F2: 597}
                    _ = fi123
                    _ = v124
                }
            }
            _ = v119
        } else {
            for fi125 := 0; fi125 < 4; fi125++ {
                fmt.Println(v5, fi125, v5)
                if true {
                    for fi126 := 0; fi126 < 11; fi126++ {
                        fmt.Println(v6, v3)
                        v127 := []int{fi126, -1, 42}
                        _ = fi126
                        _ = v127
                    }
                } else {
                    switch (fi125 - fi125) {
                    case -619:
                        var v128 int16 = (v7 << ((95) & 63))
                        _ = v128
                    case 1:
                        var v129 uint64 = 4294967295
                        _ = v129
                    }
                    v130 := []int{(fi125 / ((127) | 1))}
                    v20 = (((v20 + ("" + (v5 + "\n\t"))) + (((v5 + v5) + "\n\t") + v20)) + v5)
                    _ = v130
                }
                var v131 int8 = 59
                _ = fi125
                _ = v131
            }
        }
    }
    for fi132 := 0; fi132 < 15; fi132++ {
        v133 := GenS8{F0: 0}
        v134 := []int{fi132, (fi132 | 1), fi132}
        _ = fi132
        _ = v133
        _ = v134
    }
    var v135 int = 127
    v7 = (^(v7))
    fmt.Println("v1:", v1)
    fmt.Println("v135:", v135)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v135
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
