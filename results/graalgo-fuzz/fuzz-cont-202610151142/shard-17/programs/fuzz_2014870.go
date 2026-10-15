package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 uint32 }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 uint32; F1 uint64; F2 int8 }
type GenS4 struct { F0 uint8; F1 uint32; F2 uint16 }
type GenS5 struct { F0 uint16; F1 int16 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint64; F1 int; F2 uint8 }
type GenS8 struct { F0 int64; F1 uint8; F2 int32 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 uint8; F1 int32; F2 uint16 }
type GenS11 struct { F0 string; F1 int32; F2 int32 }
type GenS12 struct { F0 int16; F1 int16; F2 int }
func genFn0(p0 uint, p1 int32) uint {
    var v5 int16 = 0
    switch "\n\t" {
    case "":
        v6 := GenS1{F0: "foo bar", F1: "a", F2: 235}
        _ = v6
    default:
        v7 := []int{0}
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v5
    return (p0 >> (((^(p0))) & 63))
}
func genFn1(p0 int16, p1 uint, p2 bool) int64 {
    for fi8 := 0; fi8 < 17; fi8++ {
        p2 = p2
        _ = fi8
    }
    p2 = p2
    fmt.Println(p0, p2)
    v9 := GenS1{F0: "foo bar", F1: ("\n\t" + "a"), F2: 2147483648}
    var v10 int64 = 363
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return -9223372036854775808
}

func main() {
    var v1 string = ("a" + "")
    var v2 int = 516
    var v3 bool = true
    var v4 uint8 = 127
    if v3 {
        var v11 bool = false
        v12 := GenS2{F0: 9223372036854775807, F1: 2147483647}
        _ = v11
        _ = v12
    } else {
        switch v1 {
        case "a":
            for fi13 := 0; fi13 < 7; fi13++ {
                v1 = (v1 + (v1 + v1))
                for fi14 := 0; fi14 < 8; fi14++ {
                    v1 = v1
                    fmt.Println(v3)
                    var v15 uint16 = 2
                    _ = fi14
                    _ = v15
                }
                v4 += (v4 + v4)
                _ = fi13
            }
        case "":
            var v16 int8 = 111
            _ = v16
        case "\n\t":
            for fi17 := 0; fi17 < 9; fi17++ {
                v18 := []int{128}
                v1 = v1
                _ = fi17
                _ = v18
            }
        default:
            fmt.Println(v3, v3)
        }
    }
    var v19 uint64 = 18446744073709551615
    var v20 uint32 = 4294967295
    fmt.Println(v20, v1, v4)
    var v21 uint = 42
    if false {
        if (!false) {
            if false {
                v4 -= v4
                v22 := GenS1{F0: ((("hello" + ("hello" + v1)) + "\n\t") + v1), F1: "a", F2: (^(v20))}
                v20 += 42
                _ = v22
            }
            switch v1 {
            case "":
                var v23 float32 = 1.5
                _ = v23
            case "a":
                v24 := GenS3{F0: (v20 * 42), F1: (v19 * 18446744073709551615), F2: 5}
                _ = v24
            case "hello":
                for fi25 := 0; fi25 < 11; fi25++ {
                    var v26 uint = (^(v21))
                    for fi27 := 0; fi27 < 3; fi27++ {
                        v28 := []int{int(v19), 127, fi27, -1}
                        v2 += (-(fi27))
                        _ = fi27
                        _ = v28
                    }
                    _ = fi25
                    _ = v26
                }
            default:
                if false {
                    var v29 int64 = 575
                    var v30 uint16 = 255
                    _ = v29
                    _ = v30
                } else {
                    for fi31 := 0; fi31 < 14; fi31++ {
                        v3 = ((!v3) || (52 >= fi31))
                        _ = fi31
                    }
                    var v32 int = v2
                    var v33 string = ""
                    _ = v32
                    _ = v33
                }
            }
            if (-2.25 == 3.14159) {
                switch 815 {
                case 564:
                    var v34 uint8 = v4
                    _ = v34
                case 9223372036854775807:
                    var v35 uint8 = (v4 & (^(v4)))
                    _ = v35
                case 1:
                    if false {
                        v19 = v19
                        v3 = (v20 > 681)
                    } else {
                        v36 := []int{}
                        _ = v36
                    }
                default:
                    v3 = ((v20 / ((v20) | 1)) > v20)
                }
                var v37 int64 = int64(v20)
                var v38 int16 = -605
                _ = v37
                _ = v38
            }
        } else {
            v2 += (v2 * 2147483647)
        }
        v19 += (v19 >> ((v21) & 63))
    } else {
        var v39 int32 = 452
        var v40 float64 = 1e10
        var v41 uint8 = v4
        _ = v39
        _ = v40
        _ = v41
    }
    for fi42 := 0; fi42 < 13; fi42++ {
        if false {
            for fi43 := 0; fi43 < 17; fi43++ {
                v44 := GenS4{F0: 128, F1: v20, F2: 2}
                _ = fi43
                _ = v44
            }
            var v45 uint32 = 0
            _ = v45
        }
        if v3 {
            v46 := []int{(fi42 ^ (fi42 & -948)), fi42, v2}
            v47 := GenS2{F0: -605, F1: -907}
            v48 := GenS5{F0: 0, F1: 42}
            _ = v46
            _ = v47
            _ = v48
        } else {
            fmt.Println(v19)
        }
        switch ("a" + ("foo bar" + v1)) {
        case "foo bar":
            var v49 uint16 = 259
            _ = v49
        case "":
            fmt.Println(v4, v4)
        case "a":
            var v50 int32 = 0
            _ = v50
        default:
            switch "\n\t" {
            case "\n\t":
                for fi51 := 0; fi51 < 8; fi51++ {
                    v4 += v4
                    v52 := GenS6{F0: 204}
                    v21 -= (v21 * 4294967295)
                    _ = fi51
                    _ = v52
                }
            case "foo bar":
                var v53 string = (v1 + ((v1 + (v1 + ("" + v1))) + "\n\t"))
                _ = v53
            case "":
                if ((((fi42 - (v2 - (^(fi42)))) > (^(fi42))) || v3) || v3) {
                    var v54 int32 = 0
                    fmt.Println(v21, v2, v4)
                    _ = v54
                }
            default:
                if ((^(v4)) > v4) {
                    fmt.Println(v21)
                    var v55 int8 = 42
                    _ = v55
                }
            }
        }
        _ = fi42
    }
    for fi56 := 0; fi56 < 8; fi56++ {
        v57 := GenS7{F0: (v19 ^ 42), F1: (^(fi56)), F2: 59}
        v58 := GenS8{F0: -9223372036854775808, F1: 0, F2: 1}
        _ = fi56
        _ = v57
        _ = v58
    }
    if (v3 && v3) {
        if v3 {
            v59 := GenS1{F0: v1, F1: "", F2: (v20 * (v20 - (v20 & v20)))}
            switch ("hello" + v1) {
            case "hello":
                for fi60 := 0; fi60 < 19; fi60++ {
                    var v61 uint16 = 1
                    var v62 int = v2
                    for fi63 := 0; fi63 < 13; fi63++ {
                        if v3 {
                            v64 := GenS9{F0: v2}
                            if (v19 != v19) {
                                var v65 int16 = 32767
                                v66 := GenS4{F0: 42, F1: (v20 & v20), F2: v61}
                                if v3 {
                                    v67 := GenS6{F0: 2147483647}
                                    v65 += 42
                                    _ = v67
                                }
                                var v68 int64 = 42
                                _ = v65
                                _ = v66
                                _ = v68
                            } else {
                                switch v1 {
                                case "\n\t":
                                    v21 -= v21
                                case "":
                                    v62 = fi63
                                }
                                switch (fi63 << (((v21 << ((v21) & 63))) & 63)) {
                                case 2147483647:
                                    var v69 int = 690
                                    _ = v69
                                case 0:
                                    var v70 float32 = 1e10
                                    _ = v70
                                default:
                                    switch (fi60 / ((v2) | 1)) {
                                    case -2147483648:
                                        if v3 {
                                            switch (v2 & 32767) {
                                            case 32768:
                                                var v71 bool = (((v3 && (1000 >= int16(v20))) || (v1 <= v1)) || v3)
                                                _ = v71
                                            case 876:
                                                if ((^(v21)) != (v21 + (v21 ^ 2))) {
                                                    var v72 uint64 = uint64(v20)
                                                    _ = v72
                                                } else {
                                                    v73 := []int{(v62 >> ((v21) & 63)), (v2 ^ (^(fi60))), v2, 150}
                                                    var v74 uint32 = v20
                                                    v20 -= (v74 << ((v21) & 63))
                                                    _ = v73
                                                    _ = v74
                                                }
                                            default:
                                                v61 = (^(v61))
                                            }
                                            switch "\n\t" {
                                            case "foo bar":
                                                fmt.Println(v4, v19)
                                                switch v2 {
                                                case -198:
                                                    var v75 float64 = (3.14159 * float64(0.0))
                                                    _ = v75
                                                case 9223372036854775807:
                                                    v76 := []int{}
                                                    _ = v76
                                                case 0:
                                                    var v77 int64 = -9223372036854775808
                                                    _ = v77
                                                }
                                            case "a":
                                                fmt.Println(v62, v4)
                                                if ((v19 | uint64(v61)) != v19) {
                                                    v78 := GenS1{F0: (("a" + (v1 + ("" + v1))) + "foo bar"), F1: (v1 + v1), F2: v20}
                                                    var v79 uint16 = 620
                                                    var v80 int32 = 42
                                                    _ = v78
                                                    _ = v79
                                                    _ = v80
                                                }
                                            }
                                            if v3 {
                                                switch ("a" + "hello") {
                                                case "\n\t":
                                                    var v81 int8 = 100
                                                    _ = v81
                                                default:
                                                    if (1 == -734) {
                                                        fmt.Println(v19, fi63)
                                                        fmt.Println(v62, v19)
                                                        v20 = uint32(v21)
                                                    }
                                                }
                                            }
                                        }
                                    }
                                }
                                v82 := []int{(v2 ^ fi60), (fi60 & (fi63 / ((42) | 1)))}
                                var v83 uint64 = 237
                                _ = v82
                                _ = v83
                            }
                            _ = v64
                        } else {
                            v84 := GenS6{F0: 1}
                            _ = v84
                        }
                        _ = fi63
                    }
                    _ = fi60
                    _ = v61
                    _ = v62
                }
            case "":
                if v3 {
                    for fi85 := 0; fi85 < 1; fi85++ {
                        v86 := []int{}
                        var v87 bool = v3
                        for fi88 := 0; fi88 < 20; fi88++ {
                            v4 = (v4 + (^(v4)))
                            _ = fi88
                        }
                        _ = fi85
                        _ = v86
                        _ = v87
                    }
                    for fi89 := 0; fi89 < 9; fi89++ {
                        fmt.Println(v19, v21, v21)
                        for fi90 := 0; fi90 < 2; fi90++ {
                            var v91 uint16 = 953
                            v19 = (^(v19))
                            if (1.5 == 0.0) {
                                if v3 {
                                    fmt.Println(fi89)
                                    var v92 float32 = 3.14159
                                    _ = v92
                                }
                                var v93 int32 = 0
                                var v94 float32 = 0.1
                                _ = v93
                                _ = v94
                            }
                            _ = fi90
                            _ = v91
                        }
                        if (100 == -1) {
                            var v95 int8 = -1
                            switch fi89 {
                            case -9223372036854775808:
                                switch (fi89 + -2147483648) {
                                case -2147483648:
                                    if v3 {
                                        switch ("" + v1) {
                                        case "foo bar":
                                            v96 := GenS9{F0: (fi89 | 4294967295)}
                                            _ = v96
                                        }
                                        v97 := []int{(v2 - (fi89 ^ (fi89 / ((-45) | 1))))}
                                        _ = v97
                                    }
                                case 309:
                                    if v3 {
                                        for fi98 := 0; fi98 < 16; fi98++ {
                                            if (42 > (v19 - v19)) {
                                                if false {
                                                    var v99 int16 = -32768
                                                    _ = v99
                                                }
                                                fmt.Println(v1, v4, v19)
                                            } else {
                                                var v100 uint64 = v19
                                                v101 := []int{-475, fi89}
                                                switch (v1 + ("hello" + "\n\t")) {
                                                case "foo bar":
                                                    switch 42 {
                                                    case 128:
                                                        switch (fi98 % ((2147483647) | 1)) {
                                                        case 128:
                                                            v21 = 42
                                                        case 42:
                                                            v101 = append(v101, (fi98 + 32768))
                                                        case 1:
                                                            var v102 float64 = 1.5
                                                            if len(v101) >= 2 {
                                                                v101 = v101[1:len(v101)]
                                                            }
                                                            switch ("hello" + v1) {
                                                            case "hello":
                                                                v2 -= (fi89 % ((int(v19)) | 1))
                                                            case "a":
                                                                fmt.Println(v2, v4, v21)
                                                            case "foo bar":
                                                                v103 := GenS1{F0: "\n\t", F1: "foo bar", F2: (v20 % ((0) | 1))}
                                                                _ = v103
                                                            default:
                                                                if (915 < -440) {
                                                                    fmt.Println(v2)
                                                                    fmt.Println(v21, v3, v19)
                                                                    v104 := GenS10{F0: 153, F1: -19, F2: 283}
                                                                    v21 += v21
                                                                    fmt.Println(v19, v95)
                                                                    v19 += (^(v100))
                                                                    _ = v104
                                                                }
                                                            }
                                                            _ = v102
                                                        default:
                                                            var v105 float64 = (0.0 * (0.1 / 1e10))
                                                            _ = v105
                                                        }
                                                    case -362:
                                                        if len(v101) > 0 {
                                                            v101[0] = fi98
                                                        }
                                                    case -731:
                                                        _ = len(v101)
                                                    default:
                                                        var v106 uint8 = 128
                                                        _ = v106
                                                    }
                                                case "a":
                                                    v95 -= v95
                                                case "hello":
                                                    v107 := GenS1{F0: v1, F1: v1, F2: v20}
                                                    _ = v107
                                                }
                                                _ = v100
                                                _ = v101
                                            }
                                            v108 := GenS3{F0: (v20 & 2147483647), F1: uint64(v20), F2: (^(v95))}
                                            v109 := GenS11{F0: v1, F1: 0, F2: -1}
                                            v110 := []int{fi89, (fi89 >> ((255) & 63)), (-(fi89)), 2147483647}
                                            switch v2 {
                                            case 128:
                                                var v111 string = v1
                                                _ = v111
                                            }
                                            _ = fi98
                                            _ = v108
                                            _ = v109
                                            _ = v110
                                        }
                                        var v112 int = v2
                                        _ = v112
                                    } else {
                                        for fi113 := 0; fi113 < 1; fi113++ {
                                            if (!false) {
                                                switch (fi113 >> ((genFn0((v21 | v21), -459)) & 63)) {
                                                case -9223372036854775808:
                                                    v2 -= (fi89 ^ v2)
                                                    if (!true) {
                                                        switch (v1 + ("a" + v1)) {
                                                        case "foo bar":
                                                            var v114 float64 = 1e10
                                                            _ = v114
                                                        case "\n\t":
                                                            fmt.Println(fi113, v19, v20)
                                                        }
                                                    }
                                                case 9223372036854775807:
                                                    if false {
                                                        fmt.Println(fi113)
                                                        fmt.Println(v4, v4, fi89)
                                                        v4 = (v4 | (v4 % (((^(v4))) | 1)))
                                                    }
                                                }
                                            } else {
                                                v3 = v3
                                                fmt.Println(v4, fi113, fi113)
                                                if false {
                                                    var v115 float64 = -2.25
                                                    v95 = v95
                                                    if (-797 < 179) {
                                                        v116 := []int{-9223372036854775808, (fi113 | fi89), (v2 % ((2147483647) | 1))}
                                                        var v117 int8 = (v95 * -128)
                                                        _ = v116
                                                        _ = v117
                                                    }
                                                    _ = v115
                                                }
                                            }
                                            var v118 uint64 = v19
                                            v20 -= v20
                                            _ = fi113
                                            _ = v118
                                        }
                                        if false {
                                            v119 := []int{fi89, (v2 - (v2 / ((v2) | 1))), (fi89 * fi89)}
                                            if (997 >= 1) {
                                                var v120 uint16 = 435
                                                if v3 {
                                                    for fi121 := 0; fi121 < 17; fi121++ {
                                                        var v122 float32 = (3.14159 * 3.14159)
                                                        switch (fi89 + (fi121 * (fi121 >> ((910) & 63)))) {
                                                        case 128:
                                                            fmt.Println(v19)
                                                        case 4294967295:
                                                            v123 := GenS10{F0: v4, F1: -44, F2: (^(v120))}
                                                            _ = v123
                                                        case 175:
                                                            if len(v119) > 0 {
                                                                v119[0] = (^(v2))
                                                            }
                                                        }
                                                        _ = fi121
                                                        _ = v122
                                                    }
                                                } else {
                                                    for fi124 := 0; fi124 < 19; fi124++ {
                                                        var v125 int8 = v95
                                                        _ = fi124
                                                        _ = v125
                                                    }
                                                    if v3 {
                                                        switch (v1 + "a") {
                                                        case "hello":
                                                            var v126 uint32 = 4294967295
                                                            _ = v126
                                                        case "\n\t":
                                                            if v3 {
                                                                var v127 int = 0
                                                                var v128 float64 = 3.14159
                                                                if len(v119) >= 2 {
                                                                    v119 = v119[1:len(v119)]
                                                                }
                                                                _ = v127
                                                                _ = v128
                                                            }
                                                        case "":
                                                            switch ((v1 + v1) + ((v1 + v1) + ("a" + ""))) {
                                                            case "hello":
                                                                for fi129 := 0; fi129 < 14; fi129++ {
                                                                    if len(v119) > 0 {
                                                                        v119[0] = (fi89 ^ fi89)
                                                                    }
                                                                    _ = fi129
                                                                }
                                                            case "foo bar":
                                                                v95 += v95
                                                            default:
                                                                for fi130 := 0; fi130 < 14; fi130++ {
                                                                    v19 = (v19 - (v19 ^ (v19 ^ 1)))
                                                                    _ = fi130
                                                                }
                                                            }
                                                        }
                                                        v19 = v19
                                                        if (v3 && (!v3)) {
                                                            v3 = v3
                                                        } else {
                                                            v21 -= (v21 | (v21 << (((v21 % (((v21 ^ (^(v21)))) | 1))) & 63)))
                                                            switch (v1 + (v1 + "hello")) {
                                                            case "\n\t":
                                                                switch "" {
                                                                case "a":
                                                                    if v3 {
                                                                        v19 += (v19 & v19)
                                                                    }
                                                                }
                                                            case "":
                                                                var v131 int16 = -32768
                                                                _ = v131
                                                            case "a":
                                                                v119 = append(v119, (^(fi89)))
                                                            default:
                                                                var v132 bool = false
                                                                _ = v132
                                                            }
                                                            switch (^(v2)) {
                                                            case 32767:
                                                                v133 := GenS7{F0: v19, F1: (^(v2)), F2: (v4 >> (((^(v21))) & 63))}
                                                                _ = v133
                                                            case 127:
                                                                for fi134 := 0; fi134 < 9; fi134++ {
                                                                    v119 = append(v119, 127)
                                                                    v135 := GenS7{F0: (^(v19)), F1: (^(v2)), F2: (v4 + (^(v4)))}
                                                                    v2 = (fi134 - (v2 - (v2 - (^(fi89)))))
                                                                    _ = fi134
                                                                    _ = v135
                                                                }
                                                            default:
                                                                v20 += 201
                                                            }
                                                        }
                                                    }
                                                }
                                                _ = len(v119)
                                                _ = v120
                                            }
                                            _ = v119
                                        }
                                    }
                                case 32767:
                                    var v136 float64 = (3.14159 + (1e10 - (float64(0.0) * float64((((0.0 * 0.0) - 0.1) * 0.0)))))
                                    _ = v136
                                default:
                                    v19 -= v19
                                }
                            case 32768:
                                var v137 float32 = 1.5
                                _ = v137
                            default:
                                switch "hello" {
                                case "\n\t":
                                    v138 := GenS11{F0: ((v1 + "a") + v1), F1: 42, F2: -2147483648}
                                    _ = v138
                                case "a":
                                    v4 += v4
                                case "":
                                    var v139 bool = true
                                    _ = v139
                                }
                            }
                            _ = v95
                        } else {
                            fmt.Println(v1)
                            if v3 {
                                var v140 uint64 = 509
                                v141 := []int{(fi89 ^ v2)}
                                _ = v140
                                _ = v141
                            } else {
                                fmt.Println(v19, v4, v4)
                                v142 := GenS12{F0: 1000, F1: -440, F2: (v2 << (((v21 / (((v21 ^ (^(v21)))) | 1))) & 63))}
                                _ = v142
                            }
                        }
                        _ = fi89
                    }
                    v143 := []int{(^(v2))}
                    _ = v143
                } else {
                    var v144 int64 = 9223372036854775807
                    v4 += uint8(v19)
                    var v145 uint = genFn0(v21, 2147483647)
                    _ = v144
                    _ = v145
                }
            case "\n\t":
                for fi146 := 0; fi146 < 16; fi146++ {
                    var v147 int8 = -36
                    for fi148 := 0; fi148 < 9; fi148++ {
                        v149 := GenS12{F0: 1000, F1: 41, F2: 541}
                        fmt.Println(v21)
                        _ = fi148
                        _ = v149
                    }
                    v147 = (v147 + (v147 & (v147 % ((v147) | 1))))
                    _ = fi146
                    _ = v147
                }
            }
            _ = v59
        } else {
            v150 := []int{v2, v2, -2147483648}
            switch (v1 + (v1 + "")) {
            case "hello":
                for fi151 := 0; fi151 < 4; fi151++ {
                    v1 = v1
                    v21 = (v21 ^ (v21 ^ 42))
                    if v3 {
                        var v152 float32 = (3.14159 * 0.1)
                        fmt.Println(v3, v2)
                        v3 = v3
                        _ = v152
                    }
                    _ = fi151
                }
            case "a":
                v1 = ((v1 + ((((v1 + "\n\t") + ((((v1 + ("hello" + v1)) + ("" + "\n\t")) + (v1 + (v1 + v1))) + "")) + ("a" + v1)) + v1)) + v1)
            }
            _ = v150
        }
        var v153 int16 = 960
        var v154 uint16 = 710
        _ = v153
        _ = v154
    } else {
        var v155 uint16 = 32767
        v20 -= (v20 | (v20 & uint32(v21)))
        _ = v155
    }
    switch v1 {
    case "hello":
        fmt.Println(v21, v1)
    }
    v156 := GenS9{F0: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v156
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
