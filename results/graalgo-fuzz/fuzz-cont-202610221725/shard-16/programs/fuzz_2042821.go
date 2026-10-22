package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 uint8; F1 string; F2 uint8 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 string; F1 int }
type GenS7 struct { F0 uint8; F1 bool; F2 string }
type GenS8 struct { F0 bool; F1 uint32; F2 uint64 }
type GenS9 struct { F0 uint16; F1 string }
type GenS10 struct { F0 string }
type GenS11 struct { F0 int64; F1 string }
type GenS12 struct { F0 int16 }
type GenS13 struct { F0 int16; F1 uint8 }
type GenS14 struct { F0 int64; F1 int }
type GenS15 struct { F0 uint8; F1 uint16; F2 bool }
type GenS16 struct { F0 string; F1 uint8; F2 string }
type GenS17 struct { F0 bool; F1 uint16; F2 int32 }
type GenS18 struct { F0 uint; F1 uint32; F2 uint }
type GenS19 struct { F0 int8; F1 int32; F2 uint8 }
type GenS20 struct { F0 uint16 }
type GenS21 struct { F0 uint; F1 uint; F2 int32 }
type GenS22 struct { F0 uint64 }
type GenS23 struct { F0 uint32 }
type GenS24 struct { F0 int16 }

func main() {
    var v1 uint8 = 2
    var v2 int8 = 127
    var v3 int16 = 32767
    var v4 int32 = -324
    var v5 uint = 390
    var v6 uint32 = 636
    for fi7 := 0; fi7 < 9; fi7++ {
        fmt.Println(v5, v3, v5)
        v5 -= (^(v5))
        if (((42 <= 616) || (!true)) && false) {
            switch "\n\t" {
            case "a":
                if false {
                    var v8 int32 = 519
                    if ("a" < "a") {
                        switch -388 {
                        case 897:
                            v1 = v1
                        default:
                            if true {
                                switch fi7 {
                                case 373:
                                    v6 = (v6 % ((v6) | 1))
                                case 32767:
                                    switch fi7 {
                                    case 0:
                                        for fi9 := 0; fi9 < 13; fi9++ {
                                            if false {
                                                var v10 int8 = (v2 + (v2 ^ (v2 + (v2 / ((v2) | 1)))))
                                                v10 += (v2 * -30)
                                                _ = v10
                                            }
                                            fmt.Println(v8, fi9, v3)
                                            _ = fi9
                                        }
                                    case -2147483648:
                                        v4 += (v8 % (((v8 & (^(v4)))) | 1))
                                    }
                                case 1:
                                    v8 -= v4
                                }
                            }
                        }
                        if (((1e10 + 0.1) > 0.0) && true) {
                            if true {
                                v11 := GenS1{F0: v2}
                                var v12 uint = 255
                                var v13 int = fi7
                                _ = v11
                                _ = v12
                                _ = v13
                            } else {
                                var v14 int = fi7
                                if (!(((v14 * (fi7 / ((int(v8)) | 1))) == -582) || (2 != 0))) {
                                    var v15 float32 = float32((((1.5 + 1e10) + -2.25) * float64((1e10 + float32(1.5)))))
                                    var v16 int16 = (^(v3))
                                    var v17 int8 = (v2 - 127)
                                    _ = v15
                                    _ = v16
                                    _ = v17
                                }
                                _ = v14
                            }
                        }
                    } else {
                        switch ("\n\t" + "hello") {
                        case "foo bar":
                            switch (("a" + "a") + "") {
                            case "":
                                for fi18 := 0; fi18 < 5; fi18++ {
                                    v19 := GenS1{F0: (v2 / (((-(v2))) | 1))}
                                    _ = fi18
                                    _ = v19
                                }
                            case "\n\t":
                                if true {
                                    v20 := []int{(fi7 + fi7), 302}
                                    v6 = v6
                                    switch "foo bar" {
                                    case "a":
                                        var v21 float32 = float32(3.14159)
                                        _ = v21
                                    case "foo bar":
                                        for fi22 := 0; fi22 < 5; fi22++ {
                                            var v23 uint8 = (v1 ^ 235)
                                            v24 := GenS1{F0: v2}
                                            var v25 uint8 = v23
                                            _ = fi22
                                            _ = v23
                                            _ = v24
                                            _ = v25
                                        }
                                    case "":
                                        if (int32(fi7) > -2147483648) {
                                            if (-128 == (-(v2))) {
                                                v6 = 4294967295
                                                _ = len(v20)
                                            } else {
                                                var v26 float64 = ((0.0 + 1e10) + (1e10 + float64(0.1)))
                                                switch "\n\t" {
                                                case "hello":
                                                    for fi27 := 0; fi27 < 17; fi27++ {
                                                        fmt.Println(v3, v5)
                                                        v26 += v26
                                                        _ = fi27
                                                    }
                                                }
                                                fmt.Println(v5)
                                                _ = v26
                                            }
                                        }
                                    }
                                    _ = v20
                                } else {
                                    var v28 int8 = (v2 % ((v2) | 1))
                                    v1 += (v1 & v1)
                                    _ = v28
                                }
                            }
                        }
                        if true {
                            v29 := GenS2{F0: (v3 >> ((18446744073709551615) & 63)), F1: 32767}
                            v30 := []int{fi7, (fi7 & (fi7 & fi7))}
                            _ = v29
                            _ = v30
                        } else {
                            var v31 float32 = 1e10
                            fmt.Println(v4, v31)
                            _ = v31
                        }
                        v4 += (v4 | v8)
                    }
                    fmt.Println(v4, fi7, fi7)
                    _ = v8
                } else {
                    v32 := []int{int(v5), 9223372036854775807}
                    _ = v32
                }
            case "\n\t":
                v2 = v2
            case "":
                v33 := []int{128, (fi7 >> (((v5 & (v5 + (v5 >> (((v5 + v5)) & 63))))) & 63)), (fi7 * (fi7 + 1))}
                _ = v33
            }
            v4 += -1
        }
        _ = fi7
    }
    for fi34 := 0; fi34 < 17; fi34++ {
        for fi35 := 0; fi35 < 5; fi35++ {
            switch fi35 {
            case -1:
                if (4294967295 >= (v5 + v5)) {
                    v4 -= v4
                    fmt.Println(v4, v4)
                }
            case 1:
                switch "hello" {
                case "a":
                    if false {
                        if false {
                            var v36 string = "\n\t"
                            v2 = v2
                            _ = v36
                        }
                        if true {
                            if ((v5 != v5) || false) {
                                v37 := []int{}
                                fmt.Println(v6)
                                _ = v37
                            } else {
                                if true {
                                    var v38 string = "a"
                                    if false {
                                        v5 = v5
                                    } else {
                                        v39 := GenS1{F0: -11}
                                        _ = v39
                                    }
                                    if true {
                                        var v40 int16 = v3
                                        v41 := GenS1{F0: int8(v4)}
                                        _ = v40
                                        _ = v41
                                    } else {
                                        var v42 int16 = v3
                                        v43 := GenS3{F0: 255, F1: v1}
                                        _ = v42
                                        _ = v43
                                    }
                                    _ = v38
                                }
                                v6 = 2147483647
                                var v44 uint8 = v1
                                _ = v44
                            }
                            var v45 uint16 = 0
                            v46 := GenS2{F0: (-(v3)), F1: v45}
                            _ = v45
                            _ = v46
                        } else {
                            fmt.Println(v5, v1)
                        }
                    } else {
                        v47 := []int{127, (^(fi35)), fi34}
                        v3 -= v3
                        v48 := GenS4{F0: (^(v1)), F1: ("hello" + ("\n\t" + "\n\t")), F2: v1}
                        _ = v47
                        _ = v48
                    }
                }
            case -2147483648:
                var v49 int8 = (v2 | (v2 * (v2 >> ((0) & 63))))
                _ = v49
            }
            v50 := GenS2{F0: -529, F1: 255}
            v6 += uint32(v5)
            v4 = v4
            _ = fi35
            _ = v50
        }
        switch -9223372036854775808 {
        case -9223372036854775808:
            fmt.Println(v6, v5)
        default:
            for fi51 := 0; fi51 < 15; fi51++ {
                v52 := GenS5{F0: v2}
                var v53 int = 2147483647
                var v54 string = "foo bar"
                _ = fi51
                _ = v52
                _ = v53
                _ = v54
            }
        }
        v55 := GenS6{F0: "\n\t", F1: (fi34 >> (((^(v5))) & 63))}
        _ = fi34
        _ = v55
    }
    if (int64(v6) <= -560) {
        fmt.Println(v4)
        for fi56 := 0; fi56 < 19; fi56++ {
            v57 := []int{fi56}
            var v58 uint16 = 255
            if len(v57) >= 2 {
                v57 = v57[1:len(v57)]
            }
            _ = fi56
            _ = v57
            _ = v58
        }
    } else {
        v59 := GenS7{F0: uint8(v6), F1: true, F2: ""}
        switch "" {
        case "hello":
            if true {
                switch 9223372036854775807 {
                case 32767:
                    v3 = v3
                case 0:
                    v60 := []int{32768, 4294967295, 9223372036854775807}
                    _ = v60
                default:
                    v2 = v2
                }
            }
        case "foo bar":
            for fi61 := 0; fi61 < 6; fi61++ {
                if (1 > v4) {
                    for fi62 := 0; fi62 < 9; fi62++ {
                        if true {
                            if true {
                                v63 := GenS8{F0: (float32(((1e10 + (1.5 - 1e10)) - -2.25)) != ((1e10 * 1e10) * ((0.1 + 1.5) + 0.0))), F1: (v6 & (^(v6))), F2: 42}
                                switch "\n\t" {
                                case "":
                                    var v64 float32 = 1.5
                                    v65 := []int{fi62, -2147483648}
                                    var v66 float64 = (float64((v64 / v64)) - ((float64(3.14159) * (0.1 + (((0.1 + (((0.1 + 0.1) + 0.1) - (0.0 - (((0.0 * (-2.25 / 0.1)) - float64(v64)) + 1e10)))) + (-2.25 * 1e10)) - 1e10))) + ((1.5 - -2.25) - 3.14159)))
                                    _ = v64
                                    _ = v65
                                    _ = v66
                                default:
                                    v67 := GenS9{F0: 42, F1: "foo bar"}
                                    _ = v67
                                }
                                if (!false) {
                                    fmt.Println(v2, fi61)
                                }
                                _ = v63
                            }
                            var v68 uint64 = 434
                            switch (fi62 + 4294967295) {
                            case -216:
                                var v69 uint = (v5 ^ 605)
                                _ = v69
                            case -1:
                                v70 := []int{int(v1), (fi62 * int(v5)), 854, -464}
                                _ = v70
                            default:
                                fmt.Println(v3, v3)
                            }
                            _ = v68
                        } else {
                            v71 := []int{fi61, (fi61 % ((fi61) | 1)), fi61}
                            _ = v71
                        }
                        _ = fi62
                    }
                    var v72 float64 = 1.5
                    v73 := GenS10{F0: ""}
                    _ = v72
                    _ = v73
                }
                if false {
                    v74 := GenS7{F0: v1, F1: ((!false) && (true || false)), F2: (("foo bar" + "foo bar") + (("\n\t" + "hello") + "\n\t"))}
                    for fi75 := 0; fi75 < 5; fi75++ {
                        v76 := []int{fi61, fi75, fi75}
                        var v77 string = ("" + "")
                        fmt.Println(v3, v1)
                        v78 := GenS5{F0: int8(v4)}
                        fmt.Println(v2)
                        _ = fi75
                        _ = v76
                        _ = v77
                        _ = v78
                    }
                    _ = v74
                }
                for fi79 := 0; fi79 < 18; fi79++ {
                    var v80 int8 = (v2 << ((v5) & 63))
                    v2 += (v2 >> (((v5 | (v5 / ((v5) | 1)))) & 63))
                    v2 -= (^(v80))
                    _ = fi79
                    _ = v80
                }
                _ = fi61
            }
        case "":
            for fi81 := 0; fi81 < 3; fi81++ {
                var v82 int = fi81
                _ = fi81
                _ = v82
            }
        default:
            if (false || false) {
                for fi83 := 0; fi83 < 20; fi83++ {
                    switch "hello" {
                    case "hello":
                        if (0 >= 318) {
                            for fi84 := 0; fi84 < 16; fi84++ {
                                if true {
                                    switch fi83 {
                                    case 0:
                                        if ((0.1 + (0.1 + 0.1)) > (-2.25 / (-2.25 * 0.1))) {
                                            switch "" {
                                            case "foo bar":
                                                v85 := []int{fi84}
                                                _ = v85
                                            case "\n\t":
                                                v5 = 2
                                                if true {
                                                    if false {
                                                        v86 := GenS7{F0: uint8(v6), F1: (((1e10 * 1.5) + (-2.25 + 3.14159)) >= 0.0), F2: "foo bar"}
                                                        _ = v86
                                                    } else {
                                                        var v87 uint = 0
                                                        if false {
                                                            v88 := GenS2{F0: v3, F1: 592}
                                                            v1 = (v1 / ((v1) | 1))
                                                            v89 := GenS10{F0: (("" + "\n\t") + "\n\t")}
                                                            _ = v88
                                                            _ = v89
                                                        } else {
                                                            var v90 uint16 = 32767
                                                            fmt.Println(v4, v87, v87)
                                                            var v91 uint8 = (v1 << ((0) & 63))
                                                            _ = v90
                                                            _ = v91
                                                        }
                                                        _ = v87
                                                    }
                                                }
                                            case "a":
                                                var v92 string = (("a" + "\n\t") + ("foo bar" + ""))
                                                _ = v92
                                            }
                                            fmt.Println(v2, v3)
                                            v93 := GenS11{F0: -1, F1: ""}
                                            _ = v93
                                        } else {
                                            if true {
                                                var v94 string = "\n\t"
                                                fmt.Println(v6, v2, v3)
                                                _ = v94
                                            }
                                        }
                                    }
                                }
                                if false {
                                    switch "\n\t" {
                                    case "hello":
                                        v95 := GenS12{F0: (v3 * (v3 - 1))}
                                        _ = v95
                                    }
                                }
                                _ = fi84
                            }
                        } else {
                            var v96 float64 = ((0.0 + float64(0.0)) + (float64(1.5) * 1.5))
                            var v97 uint16 = uint16(v6)
                            _ = v96
                            _ = v97
                        }
                    }
                    fmt.Println(v5)
                    _ = fi83
                }
            } else {
                v4 = (v4 & v4)
                if true {
                    if true {
                        v98 := []int{}
                        _ = v98
                    } else {
                        v99 := GenS7{F0: (v1 - v1), F1: false, F2: ""}
                        v1 = v1
                        fmt.Println(v6, v1, v1)
                        _ = v99
                    }
                } else {
                    fmt.Println(v5, v2)
                    var v100 string = "hello"
                    for fi101 := 0; fi101 < 6; fi101++ {
                        if ((^(v3)) == (v3 % (((v3 - (^(v3)))) | 1))) {
                            v1 -= v1
                            v102 := []int{fi101, 632, 2147483647, 0}
                            v4 -= (^(v4))
                            _ = v102
                        } else {
                            v103 := []int{(fi101 & fi101)}
                            _ = v103
                        }
                        if false {
                            for fi104 := 0; fi104 < 11; fi104++ {
                                v105 := []int{fi104}
                                v5 -= (^(v5))
                                _ = fi104
                                _ = v105
                            }
                        }
                        if true {
                            fmt.Println(v5, v1, v1)
                        }
                        _ = fi101
                    }
                    _ = v100
                }
            }
        }
        var v106 int64 = 42
        _ = v106
        _ = v59
    }
    var v107 uint64 = 2
    var v108 uint32 = v6
    switch (((("" + ("hello" + "a")) + "foo bar") + "\n\t") + ("" + ("hello" + "foo bar"))) {
    case "\n\t":
        if true {
            for fi109 := 0; fi109 < 19; fi109++ {
                var v110 uint = (v5 & uint(v1))
                v4 += 51
                v108 += v108
                _ = fi109
                _ = v110
            }
            for fi111 := 0; fi111 < 19; fi111++ {
                switch fi111 {
                case 0:
                    v5 -= uint(v107)
                case -2147483648:
                    switch 32768 {
                    case 2147483647:
                        v112 := []int{fi111}
                        _ = v112
                    case 364:
                        v113 := []int{fi111, (fi111 + fi111)}
                        _ = v113
                    }
                case 32768:
                    v107 -= (v107 ^ v107)
                default:
                    v3 = 1000
                }
                _ = fi111
            }
            v114 := []int{1, 9223372036854775807, 32767, 127}
            _ = v114
        } else {
            fmt.Println(v108)
            fmt.Println(v5)
        }
    }
    var v115 uint64 = v107
    for fi116 := 0; fi116 < 7; fi116++ {
        if false {
            for fi117 := 0; fi117 < 9; fi117++ {
                if true {
                    v118 := []int{-299, 32768, fi117}
                    fmt.Println(v107)
                    _ = v118
                } else {
                    v119 := []int{}
                    switch "hello" {
                    case "foo bar":
                        fmt.Println(fi116)
                    case "a":
                        if (!false) {
                            if false {
                                var v120 int64 = -9223372036854775808
                                v2 += (v2 - v2)
                                if len(v119) > 0 {
                                    v119[0] = (fi116 % (((^(fi117))) | 1))
                                }
                                _ = v120
                            } else {
                                if len(v119) > 0 {
                                    v119[0] = (fi117 ^ -48)
                                }
                                switch (("\n\t" + "") + "a") {
                                case "hello":
                                    var v121 bool = false
                                    _ = v121
                                default:
                                    switch fi116 {
                                    case -9223372036854775808:
                                        var v122 uint16 = 65535
                                        _ = v122
                                    case 32767:
                                        if (v1 >= 155) {
                                            var v123 int64 = 0
                                            var v124 uint16 = 65535
                                            if false {
                                                v119 = append(v119, int(v124))
                                                var v125 uint32 = v6
                                                if true {
                                                    _ = len(v119)
                                                } else {
                                                    if (0.1 == 3.14159) {
                                                        if false {
                                                            _ = len(v119)
                                                            if (v5 == 42) {
                                                                switch (fi117 - (^(fi117))) {
                                                                case 9223372036854775807:
                                                                    v126 := GenS7{F0: v1, F1: (!false), F2: "hello"}
                                                                    _ = v126
                                                                case 32767:
                                                                    fmt.Println(v4, v115)
                                                                default:
                                                                    if false {
                                                                        if (!(v107 >= (v107 | v107))) {
                                                                            if true {
                                                                                var v127 uint8 = (v1 / ((v1) | 1))
                                                                                _ = v127
                                                                            } else {
                                                                                var v128 int32 = (v4 ^ v4)
                                                                                fmt.Println(v107)
                                                                                if (true || (1e10 <= float32(0.0))) {
                                                                                    fmt.Println(v5, v6, v115)
                                                                                    var v129 int = -2147483648
                                                                                    if (v1 < (v1 / ((v1) | 1))) {
                                                                                        if ((^(v115)) != v107) {
                                                                                            if ("hello" < ("hello" + "hello")) {
                                                                                                var v130 uint = uint(v124)
                                                                                                if len(v119) >= 2 {
                                                                                                    v119 = v119[1:len(v119)]
                                                                                                }
                                                                                                v131 := GenS13{F0: v3, F1: v1}
                                                                                                _ = v130
                                                                                                _ = v131
                                                                                            } else {
                                                                                                fmt.Println(fi117, v1, v124)
                                                                                            }
                                                                                            v119 = append(v119, v129)
                                                                                        } else {
                                                                                            fmt.Println(v125, v3)
                                                                                            if len(v119) >= 2 {
                                                                                                v119 = v119[1:len(v119)]
                                                                                            }
                                                                                            if ("a" > "foo bar") {
                                                                                                _ = len(v119)
                                                                                                var v132 uint32 = 2147483648
                                                                                                _ = v132
                                                                                            } else {
                                                                                                v133 := GenS14{F0: v123, F1: (fi116 - int(v107))}
                                                                                                fmt.Println(v108, fi117)
                                                                                                v119 = append(v119, -62)
                                                                                                _ = v133
                                                                                            }
                                                                                        }
                                                                                        v3 = 1000
                                                                                        v108 = 131
                                                                                        v123 += 1
                                                                                    } else {
                                                                                        fmt.Println(v129)
                                                                                        v115 = v107
                                                                                        var v134 uint16 = v124
                                                                                        _ = v134
                                                                                    }
                                                                                    v3 += v3
                                                                                    switch (fi117 * 128) {
                                                                                    case -1:
                                                                                        var v135 uint16 = 287
                                                                                        _ = v135
                                                                                    default:
                                                                                        var v136 uint16 = 42
                                                                                        _ = v136
                                                                                    }
                                                                                    _ = v129
                                                                                }
                                                                                _ = v128
                                                                            }
                                                                            fmt.Println(v6, v3, v5)
                                                                            fmt.Println(v115, v107)
                                                                        }
                                                                        var v137 int32 = v4
                                                                        var v138 float32 = -2.25
                                                                        _ = v137
                                                                        _ = v138
                                                                    }
                                                                }
                                                            }
                                                        } else {
                                                            if true {
                                                                var v139 uint8 = (v1 % ((128) | 1))
                                                                _ = v139
                                                            }
                                                        }
                                                        fmt.Println(v108, v125, v108)
                                                        if false {
                                                            var v140 float32 = (0.0 / (float32((1e10 / 1.5)) - (1.5 - float32(float64(-2.25)))))
                                                            v119 = append(v119, (fi116 >> ((v5) & 63)))
                                                            _ = v140
                                                        } else {
                                                            v5 = (^(v5))
                                                        }
                                                    }
                                                    v124 = v124
                                                    var v141 uint = 42
                                                    v3 = -1
                                                    _ = v141
                                                }
                                                _ = v125
                                            } else {
                                                fmt.Println(v1)
                                                v6 = 706
                                                fmt.Println(fi116)
                                                switch 32768 {
                                                case 0:
                                                    v142 := GenS3{F0: uint16(v115), F1: (v1 << ((v5) & 63))}
                                                    _ = v142
                                                case 66:
                                                    v2 += int8(v1)
                                                default:
                                                    _ = len(v119)
                                                }
                                                fmt.Println(v4, v1)
                                                v143 := GenS15{F0: 2, F1: (v124 - v124), F2: true}
                                                v119 = append(v119, (fi117 * (fi116 & 0)))
                                                if len(v119) > 0 {
                                                    v119[0] = fi116
                                                }
                                                switch (fi116 | (-(fi117))) {
                                                case 329:
                                                    if (true && false) {
                                                        fmt.Println(v4)
                                                    }
                                                case 32767:
                                                    _ = len(v119)
                                                default:
                                                    if true {
                                                        v115 = (v115 & (v115 * uint64(v108)))
                                                        var v144 int = -698
                                                        _ = v144
                                                    } else {
                                                        var v145 uint8 = (v1 + (v1 * (^(v1))))
                                                        v115 = uint64(v124)
                                                        v146 := GenS10{F0: ("" + ("foo bar" + "hello"))}
                                                        switch ("foo bar" + "a") {
                                                        case "\n\t":
                                                            fmt.Println(v2, v3, v124)
                                                        case "foo bar":
                                                            if true {
                                                                var v147 bool = false
                                                                var v148 int64 = int64(v4)
                                                                var v149 uint16 = 980
                                                                _ = v147
                                                                _ = v148
                                                                _ = v149
                                                            } else {
                                                                var v150 uint = 4294967295
                                                                _ = v150
                                                            }
                                                        }
                                                        if len(v119) >= 2 {
                                                            v119 = v119[1:len(v119)]
                                                        }
                                                        _ = v145
                                                        _ = v146
                                                    }
                                                }
                                                var v151 int8 = int8(v1)
                                                _ = v143
                                                _ = v151
                                            }
                                            if false {
                                                if false {
                                                    if false {
                                                        v152 := GenS16{F0: "\n\t", F1: 140, F2: "a"}
                                                        var v153 uint32 = 100
                                                        if false {
                                                            fmt.Println(v153)
                                                            if false {
                                                                v154 := GenS16{F0: "hello", F1: v1, F2: "\n\t"}
                                                                v155 := GenS17{F0: (false && false), F1: v124, F2: v4}
                                                                v119 = append(v119, -9223372036854775808)
                                                                var v156 int32 = (v4 - v4)
                                                                _ = v154
                                                                _ = v155
                                                                _ = v156
                                                            }
                                                        }
                                                        var v157 uint32 = (v108 >> ((42) & 63))
                                                        _ = v152
                                                        _ = v153
                                                        _ = v157
                                                    } else {
                                                        _ = len(v119)
                                                    }
                                                    v108 -= v108
                                                    switch (fi117 % (((fi116 % ((fi116) | 1))) | 1)) {
                                                    case 42:
                                                        if (false && false) {
                                                            if false {
                                                                var v158 int64 = (v123 | (v123 << ((v5) & 63)))
                                                                v115 = 18446744073709551615
                                                                fmt.Println(fi116)
                                                                var v159 uint32 = v108
                                                                _ = v158
                                                                _ = v159
                                                            } else {
                                                                if (0.0 <= (0.1 / 1.0)) {
                                                                    if len(v119) >= 2 {
                                                                        v119 = v119[1:len(v119)]
                                                                    }
                                                                    if len(v119) > 0 {
                                                                        v119[0] = (^(fi117))
                                                                    }
                                                                    var v160 uint16 = (v124 / ((uint16(v1)) | 1))
                                                                    _ = v160
                                                                } else {
                                                                    var v161 uint32 = 2147483648
                                                                    _ = v161
                                                                }
                                                                v162 := GenS8{F0: true, F1: (v108 << (((^(v5))) & 63)), F2: uint64(v1)}
                                                                if (477 == v124) {
                                                                    switch (-(fi117)) {
                                                                    case -178:
                                                                        if len(v119) > 0 {
                                                                            v119[0] = int(v4)
                                                                        }
                                                                    case 818:
                                                                        v107 -= 1
                                                                    case 0:
                                                                        var v163 bool = (true && true)
                                                                        _ = v163
                                                                    }
                                                                    if len(v119) >= 2 {
                                                                        v119 = v119[1:len(v119)]
                                                                    }
                                                                    if ("a" < "foo bar") {
                                                                        fmt.Println(v4)
                                                                    }
                                                                } else {
                                                                    v3 = (v3 | v3)
                                                                    _ = len(v119)
                                                                }
                                                                _ = v162
                                                            }
                                                            v1 += v1
                                                            switch fi117 {
                                                            case 0:
                                                                v164 := GenS1{F0: (-(v2))}
                                                                _ = v164
                                                            default:
                                                                v115 = v115
                                                            }
                                                            v165 := GenS1{F0: (v2 / ((-1) | 1))}
                                                            _ = v165
                                                        } else {
                                                            v6 = (^(v108))
                                                        }
                                                    case 127:
                                                        if (!((v124 ^ uint16(v5)) >= (v124 << ((814) & 63)))) {
                                                            if (231 >= v1) {
                                                                v124 -= uint16(v107)
                                                                fmt.Println(v1)
                                                                v166 := GenS12{F0: (^(v3))}
                                                                _ = v166
                                                            } else {
                                                                v167 := GenS13{F0: (v3 - v3), F1: v1}
                                                                var v168 int8 = v2
                                                                switch fi116 {
                                                                case -9223372036854775808:
                                                                    v2 = (v2 >> ((255) & 63))
                                                                default:
                                                                    fmt.Println(v6)
                                                                }
                                                                _ = v167
                                                                _ = v168
                                                            }
                                                            if (-128 > v2) {
                                                                v119 = append(v119, (fi117 ^ fi116))
                                                                var v169 uint8 = v1
                                                                _ = v169
                                                            }
                                                        }
                                                    case 32767:
                                                        fmt.Println(v4)
                                                    }
                                                    if true {
                                                        fmt.Println(v115, fi117, v108)
                                                        var v170 float32 = 1e10
                                                        var v171 uint16 = (v124 + uint16(v107))
                                                        var v172 float32 = v170
                                                        v107 = 2
                                                        _ = v170
                                                        _ = v171
                                                        _ = v172
                                                    } else {
                                                        var v173 uint8 = uint8(v124)
                                                        _ = v173
                                                    }
                                                    var v174 int = (-(fi117))
                                                    _ = v174
                                                } else {
                                                    var v175 float32 = 1e10
                                                    _ = v175
                                                }
                                                switch ((("foo bar" + (("" + "\n\t") + (("hello" + "") + "a"))) + "\n\t") + "\n\t") {
                                                case "":
                                                    var v176 float32 = 0.0
                                                    _ = v176
                                                case "a":
                                                    switch (("foo bar" + "a") + "hello") {
                                                    case "foo bar":
                                                        v177 := GenS6{F0: "hello", F1: fi116}
                                                        _ = v177
                                                    case "":
                                                        if ((v124 >> ((v5) & 63)) < v124) {
                                                            var v178 int = 4294967295
                                                            if (((v124 << ((v5) & 63)) < v124) && (true || ((v2 - v2) != (v2 / ((v2) | 1))))) {
                                                                fmt.Println(v6, v6)
                                                                v179 := GenS9{F0: (v124 + (v124 << ((441) & 63))), F1: "hello"}
                                                                var v180 uint = v5
                                                                _ = v179
                                                                _ = v180
                                                            } else {
                                                                switch ("hello" + ("hello" + "a")) {
                                                                case "":
                                                                    switch ("\n\t" + "hello") {
                                                                    case "hello":
                                                                        v6 = (^(v6))
                                                                    case "a":
                                                                        if (!(!false)) {
                                                                            v119 = append(v119, (fi117 | -310))
                                                                            if false {
                                                                                switch 128 {
                                                                                case -331:
                                                                                    v181 := GenS18{F0: 42, F1: (v108 / (((v6 - v108)) | 1)), F2: (v5 + (v5 ^ v5))}
                                                                                    _ = v181
                                                                                case -9223372036854775808:
                                                                                    var v182 uint8 = 51
                                                                                    _ = v182
                                                                                case 9223372036854775807:
                                                                                    switch (("a" + ("hello" + "")) + "\n\t") {
                                                                                    case "":
                                                                                        if (0.1 == 3.14159) {
                                                                                            v3 = v3
                                                                                        }
                                                                                    case "foo bar":
                                                                                        var v183 int = int(v1)
                                                                                        _ = v183
                                                                                    case "\n\t":
                                                                                        var v184 int64 = int64(v4)
                                                                                        _ = v184
                                                                                    default:
                                                                                        v185 := GenS18{F0: uint(v1), F1: (^(v6)), F2: (v5 - uint(v1))}
                                                                                        _ = v185
                                                                                    }
                                                                                default:
                                                                                    if false {
                                                                                        switch (^(fi117)) {
                                                                                        case 2147483647:
                                                                                            v115 = v107
                                                                                        case 42:
                                                                                            _ = len(v119)
                                                                                        }
                                                                                    } else {
                                                                                        var v186 string = "foo bar"
                                                                                        var v187 string = (v186 + "foo bar")
                                                                                        if len(v119) > 0 {
                                                                                            v119[0] = 4294967295
                                                                                        }
                                                                                        fmt.Println(v3, v178)
                                                                                        if len(v119) > 0 {
                                                                                            v119[0] = fi116
                                                                                        }
                                                                                        _ = v186
                                                                                        _ = v187
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                fmt.Println(v4, v107)
                                                                                switch v178 {
                                                                                case 42:
                                                                                    if true {
                                                                                        var v188 uint = (v5 << ((uint(v124)) & 63))
                                                                                        fmt.Println(v123, v4)
                                                                                        _ = v188
                                                                                    }
                                                                                case -606:
                                                                                    v107 -= uint64(v124)
                                                                                    if ((!true) && (true || true)) {
                                                                                        v2 -= (v2 | int8(v5))
                                                                                        v119 = append(v119, (fi116 - (fi117 * 128)))
                                                                                    }
                                                                                case 645:
                                                                                    v3 += (v3 >> (((v5 | v5)) & 63))
                                                                                }
                                                                            }
                                                                            var v189 int64 = (v123 % (((^(v123))) | 1))
                                                                            _ = v189
                                                                        }
                                                                    case "foo bar":
                                                                        var v190 int16 = (v3 + v3)
                                                                        _ = v190
                                                                    }
                                                                case "foo bar":
                                                                    if false {
                                                                        v108 = (v108 | v6)
                                                                        v123 = 42
                                                                        var v191 int32 = (v4 - (v4 << (((v5 - (v5 + (v5 * 211)))) & 63)))
                                                                        _ = v191
                                                                    }
                                                                case "\n\t":
                                                                    if false {
                                                                        _ = len(v119)
                                                                        var v192 uint32 = (v108 / ((v108) | 1))
                                                                        _ = v192
                                                                    } else {
                                                                        v119 = append(v119, (v178 % (((fi117 >> ((v5) & 63))) | 1)))
                                                                        var v193 int = (fi117 | (fi116 * fi117))
                                                                        _ = v193
                                                                    }
                                                                }
                                                            }
                                                            v115 = uint64(v1)
                                                            _ = v178
                                                        }
                                                    case "\n\t":
                                                        v4 = (v4 & v4)
                                                    default:
                                                        var v194 int64 = int64(v107)
                                                        _ = v194
                                                    }
                                                }
                                            } else {
                                                fmt.Println(v1)
                                                v108 = v108
                                                v1 = uint8(v5)
                                            }
                                            fmt.Println(fi116)
                                            _ = v123
                                            _ = v124
                                        }
                                        v1 -= (v1 % (((^(v1))) | 1))
                                    case 120:
                                        v195 := GenS9{F0: 0, F1: ("a" + "\n\t")}
                                        _ = v195
                                    default:
                                        var v196 int = 9223372036854775807
                                        _ = v196
                                    }
                                }
                                v107 += v115
                                v197 := GenS19{F0: (v2 + v2), F1: -2147483648, F2: v1}
                                v108 = v6
                                _ = v197
                            }
                            if (!false) {
                                if ("\n\t" >= "a") {
                                    var v198 string = ("a" + "\n\t")
                                    v199 := GenS20{F0: 606}
                                    v3 = v3
                                    _ = v198
                                    _ = v199
                                }
                                var v200 int16 = v3
                                _ = v200
                            }
                        }
                    }
                    _ = v119
                }
                var v201 uint64 = v107
                if false {
                    v1 = 111
                    fmt.Println(v108, v5)
                    if (!((v108 - (v108 >> ((v5) & 63))) > v108)) {
                        v202 := GenS21{F0: v5, F1: (^(v5)), F2: (v4 | v4)}
                        _ = v202
                    } else {
                        if false {
                            v203 := GenS19{F0: v2, F1: -406, F2: v1}
                            v2 = v2
                            var v204 float64 = 0.0
                            _ = v203
                            _ = v204
                        } else {
                            if true {
                                v1 += v1
                                v205 := GenS15{F0: (v1 ^ 199), F1: 42, F2: true}
                                switch fi117 {
                                case 925:
                                    v1 = v1
                                case 127:
                                    fmt.Println(v4, fi117, v4)
                                case 128:
                                    switch -117 {
                                    case 9223372036854775807:
                                        v115 -= uint64(v1)
                                    case 42:
                                        if (!((v108 + 462) >= 42)) {
                                            switch "a" {
                                            case "\n\t":
                                                v206 := []int{fi117, fi116}
                                                _ = v206
                                            case "foo bar":
                                                if false {
                                                    v207 := GenS22{F0: (^(v107))}
                                                    var v208 float32 = 3.14159
                                                    _ = v207
                                                    _ = v208
                                                }
                                            }
                                            if ((v5 == 738) || true) {
                                                v209 := GenS12{F0: v3}
                                                fmt.Println(v5)
                                                var v210 uint16 = 0
                                                _ = v209
                                                _ = v210
                                            }
                                            v211 := GenS23{F0: v6}
                                            _ = v211
                                        } else {
                                            v212 := GenS24{F0: -305}
                                            v2 -= 52
                                            _ = v212
                                        }
                                    }
                                default:
                                    if true {
                                        v213 := []int{(fi116 - fi117)}
                                        _ = v213
                                    }
                                }
                                if (((v201 != (v107 ^ (v201 - 168))) || (("a" <= "hello") && false)) && (true && false)) {
                                    var v214 uint16 = 65535
                                    _ = v214
                                }
                                _ = v205
                            }
                        }
                        v215 := []int{fi117, (fi116 >> ((255) & 63))}
                        _ = v215
                    }
                } else {
                    if true {
                        v216 := []int{fi117}
                        _ = v216
                    }
                    if false {
                        v201 = (v107 | uint64(v6))
                        v217 := []int{}
                        _ = v217
                    }
                    v218 := GenS5{F0: 89}
                    _ = v218
                }
                switch "foo bar" {
                case "hello":
                    v115 = v107
                default:
                    v219 := GenS24{F0: (v3 * v3)}
                    _ = v219
                }
                _ = fi117
                _ = v201
            }
        }
        var v220 bool = true
        _ = fi116
        _ = v220
    }
    for fi221 := 0; fi221 < 9; fi221++ {
        v222 := []int{fi221, fi221}
        if len(v222) > 0 {
            v222[0] = 32768
        }
        _ = fi221
        _ = v222
    }
    v223 := []int{-9223372036854775808}
    v115 = 561
    fmt.Println("v1:", v1)
    fmt.Println("v107:", v107)
    fmt.Println("v108:", v108)
    fmt.Println("v115:", v115)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v223:", len(v223), v223)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v107
    _ = v108
    _ = v115
    _ = v2
    _ = v223
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
