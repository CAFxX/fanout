package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 uint8; F1 uint32; F2 uint64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 bool; F1 string }
type GenS5 struct { F0 int64; F1 int16; F2 int32 }
type GenS6 struct { F0 bool; F1 int }
type GenS7 struct { F0 int64; F1 string; F2 bool }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int8; F1 uint16; F2 string }
type GenS10 struct { F0 int }

func main() {
    var v1 uint32 = 215
    var v2 uint = 2
    var v3 bool = false
    v4 := make(map[uint32]bool)
    v5 := make(map[int]uint32)
    var v6 uint8 = 67
    if ("a" != "hello") {
        if v3 {
            if v3 {
                var v7 int16 = 1
                if v3 {
                    var v8 int8 = -61
                    var v9 uint8 = v6
                    _ = v8
                    _ = v9
                }
                fmt.Println(v6, v2, v3)
                _ = v7
            }
            for fi10 := 0; fi10 < 16; fi10++ {
                delete(v5, 32768)
                var v11 int64 = -9223372036854775808
                v11 -= v11
                _ = fi10
                _ = v11
            }
            switch "" {
            case "":
                var v12 uint = 303
                _ = v12
            case "a":
                var v13 float64 = 0.0
                _ = v13
            case "hello":
                for fi14 := 0; fi14 < 8; fi14++ {
                    switch "" {
                    case "":
                        switch fi14 {
                        case 364:
                            var v15 int32 = int32(v2)
                            _ = v15
                        case 370:
                            if v3 {
                                v16 := []int{}
                                _ = len(v4)
                                for fi17 := 0; fi17 < 6; fi17++ {
                                    if (0.1 >= -2.25) {
                                        v18 := GenS1{F0: v6, F1: 42}
                                        fmt.Println(fi17, v6, v6)
                                        _ = v18
                                    }
                                    if (1 < 0) {
                                        if false {
                                            v6 += 255
                                            v3 = (true && v3)
                                        }
                                    } else {
                                        var v19 int32 = -2147483648
                                        switch (^(fi14)) {
                                        case -912:
                                            v20 := v5[(-(fi14))]
                                            _ = v20
                                        case 9223372036854775807:
                                            var v21 int = (fi14 ^ -113)
                                            _ = v21
                                        case 32768:
                                            v19 = 42
                                        default:
                                            v2 -= (^(v2))
                                        }
                                        v22 := GenS1{F0: v6, F1: 0}
                                        v1 -= 2147483648
                                        _ = v19
                                        _ = v22
                                    }
                                    _ = fi17
                                }
                                _ = v16
                            }
                        case -793:
                            v1 -= v1
                        }
                    case "hello":
                        delete(v5, 42)
                    default:
                        var v23 uint64 = 828
                        _ = v23
                    }
                    v3 = (((v6 << (((v2 << ((v2) & 63))) & 63)) != v6) && v3)
                    for fi24 := 0; fi24 < 1; fi24++ {
                        v25 := GenS1{F0: (v6 + v6), F1: 1}
                        _ = fi24
                        _ = v25
                    }
                    _ = fi14
                }
            }
        } else {
            v3 = (!v3)
            var v26 int8 = 100
            _ = v26
        }
    }
    for fi27 := 0; fi27 < 8; fi27++ {
        for fi28 := 0; fi28 < 20; fi28++ {
            if v3 {
                var v29 uint = (v2 * (v2 << ((0) & 63)))
                v30 := []int{fi28, int(v6), (fi27 ^ (fi28 ^ fi27)), (-(fi28))}
                _ = v29
                _ = v30
            } else {
                if v3 {
                    delete(v5, 1)
                    v3 = true
                    if (!v3) {
                        var v31 int8 = 1
                        _ = v31
                    } else {
                        if false {
                            switch ("a" + "a") {
                            case "foo bar":
                                var v32 uint64 = 18446744073709551615
                                _ = v32
                            case "":
                                var v33 uint32 = (^(v1))
                                _ = v33
                            default:
                                if v3 {
                                    var v34 uint8 = 2
                                    v3 = false
                                    if true {
                                        if false {
                                            v1 += 478
                                            switch (fi27 / (((fi27 >> (((v2 % ((v2) | 1))) & 63))) | 1)) {
                                            case 1:
                                                fmt.Println(fi27, v3)
                                            default:
                                                var v35 bool = true
                                                _ = v35
                                            }
                                            var v36 uint = v2
                                            _ = v36
                                        } else {
                                            v37 := GenS1{F0: 1, F1: 0}
                                            _ = v37
                                        }
                                    } else {
                                        fmt.Println(v1)
                                        var v38 int32 = 42
                                        var v39 uint32 = v1
                                        _ = v38
                                        _ = v39
                                    }
                                    _ = v34
                                } else {
                                    v40 := GenS1{F0: v6, F1: 0}
                                    var v41 int8 = 42
                                    v4[v1] = (((("" + "") + "hello") >= "\n\t") && ((((0.0 - 0.1) + (0.1 + (-2.25 + ((float32((((1e10 / 1.0) * (float64(1.5) / float64(-2.25))) / 1.5)) + (0.1 / -2.25)) + float32(0.0))))) != 1.5) || ((-(fi28)) == (fi27 << (((v2 - (^(v2)))) & 63)))))
                                    _ = v40
                                    _ = v41
                                }
                            }
                        }
                        v3 = v3
                    }
                }
            }
            var v42 int32 = -2147483648
            if true {
                v4[(v1 - v1)] = v3
                v43 := []int{(-(fi27)), -69}
                v2 -= 821
                _ = v43
            } else {
                switch ("\n\t" + ("\n\t" + (("" + "foo bar") + "a"))) {
                case "":
                    if true {
                        v2 = 42
                        v44 := []int{fi28, (fi27 ^ fi27)}
                        v1 = (v1 << ((v2) & 63))
                        _ = v44
                    } else {
                        if ((fi28 / ((fi28) | 1)) <= (fi28 % ((-9223372036854775808) | 1))) {
                            switch "foo bar" {
                            case "\n\t":
                                if v3 {
                                    var v45 uint = 636
                                    fmt.Println(v45)
                                    _ = v45
                                }
                            }
                        } else {
                            v46 := GenS1{F0: (v6 / (((^(v6))) | 1)), F1: 0}
                            if v3 {
                                v2 -= (v2 + (^(v2)))
                            }
                            var v47 uint = 4294967295
                            _ = v46
                            _ = v47
                        }
                        var v48 uint16 = 65535
                        _ = v48
                    }
                case "foo bar":
                    var v49 string = ""
                    _ = v49
                case "\n\t":
                    v1 -= (v1 | v1)
                    v50 := []int{fi27, fi27, 128, fi28}
                    _ = len(v5)
                    switch "a" {
                    case "a":
                        v6 = 80
                    case "hello":
                        if len(v50) >= 2 {
                            v50 = v50[1:len(v50)]
                        }
                    case "foo bar":
                        if v3 {
                            _ = len(v4)
                        }
                    }
                    _ = v50
                }
            }
            _ = fi28
            _ = v42
        }
        _ = fi27
    }
    switch 2147483647 {
    case 42:
        switch "hello" {
        case "hello":
            switch "foo bar" {
            case "hello":
                if v3 {
                    if (false && v3) {
                        var v51 float32 = (3.14159 / 0.1)
                        if ((v2 * (v2 - v2)) <= v2) {
                            v3 = true
                            for fi52 := 0; fi52 < 15; fi52++ {
                                v53 := v4[v1]
                                switch "\n\t" {
                                case "a":
                                    if v3 {
                                        v3 = true
                                        v4[65535] = (v3 && v3)
                                    } else {
                                        if ("a" == "hello") {
                                            v1 = v1
                                        } else {
                                            if (v53 && (!v3)) {
                                                var v54 int32 = -1
                                                _ = v54
                                            }
                                        }
                                    }
                                }
                                _ = fi52
                                _ = v53
                            }
                        } else {
                            v51 -= v51
                            if v3 {
                                switch -9223372036854775808 {
                                case 9223372036854775807:
                                    if (true && (!(!false))) {
                                        v55 := []int{int(v2), -22, 1, 4294967295}
                                        fmt.Println(v3, v6)
                                        _ = v55
                                    } else {
                                        v56 := GenS2{F0: (v6 >> ((180) & 63)), F1: 102, F2: 188}
                                        v57 := []int{0}
                                        _ = v56
                                        _ = v57
                                    }
                                case 32767:
                                    if (42 > 101) {
                                        var v58 int = -1
                                        _ = v58
                                    }
                                default:
                                    for fi59 := 0; fi59 < 15; fi59++ {
                                        _ = len(v4)
                                        switch 32768 {
                                        case 32768:
                                            var v60 int16 = 1
                                            _ = v60
                                        default:
                                            v61 := GenS3{F0: fi59}
                                            _ = v61
                                        }
                                        var v62 uint = v2
                                        _ = fi59
                                        _ = v62
                                    }
                                }
                            }
                            v4[2147483648] = true
                        }
                        v51 -= 3.14159
                        _ = v51
                    } else {
                        for fi63 := 0; fi63 < 9; fi63++ {
                            switch "a" {
                            case "\n\t":
                                v64 := []int{(fi63 ^ fi63), (fi63 ^ fi63), 32767, (fi63 >> (((v2 * (v2 - (v2 % ((2) | 1))))) & 63))}
                                _ = v64
                            }
                            for fi65 := 0; fi65 < 8; fi65++ {
                                if v3 {
                                    var v66 bool = (v3 && (42 > 1))
                                    v67 := []int{(fi65 | (fi63 >> (((v2 >> (((^(v2))) & 63))) & 63))), fi65}
                                    switch fi65 {
                                    case -350:
                                        delete(v4, v1)
                                        var v68 string = "\n\t"
                                        v68 = (v68 + v68)
                                        _ = v68
                                    case 42:
                                        v69 := v5[fi63]
                                        _ = v69
                                    case -2147483648:
                                        v6 = (v6 << (((v2 >> (((v2 ^ v2)) & 63))) & 63))
                                    }
                                    _ = v66
                                    _ = v67
                                }
                                v70 := GenS4{F0: ((^(v2)) > (v2 | 489)), F1: ("a" + "foo bar")}
                                v2 += (^(v2))
                                _ = fi65
                                _ = v70
                            }
                            v71 := []int{}
                            _ = fi63
                            _ = v71
                        }
                        var v72 int = 35
                        _ = v72
                    }
                    v73 := GenS4{F0: (!v3), F1: ""}
                    _ = v73
                }
            default:
                v3 = true
            }
        case "":
            switch ("foo bar" + "") {
            case "foo bar":
                for fi74 := 0; fi74 < 2; fi74++ {
                    if (42 != -1) {
                        if (!(123 == -81)) {
                            for fi75 := 0; fi75 < 13; fi75++ {
                                if (("" + "a") != "a") {
                                    switch ("\n\t" + "a") {
                                    case "":
                                        fmt.Println(v3, v2)
                                    case "a":
                                        var v76 int16 = -60
                                        _ = v76
                                    case "foo bar":
                                        if ((float64(float32(0.1)) >= 0.0) || (42 != int32(v1))) {
                                            v1 = (v1 ^ (v1 % ((65535) | 1)))
                                        }
                                    default:
                                        fmt.Println(v3, fi74)
                                    }
                                    v77 := GenS5{F0: 42, F1: 1, F2: 622}
                                    _ = v77
                                }
                                _ = fi75
                            }
                        } else {
                            if ("a" <= "hello") {
                                var v78 uint = v2
                                v2 -= v78
                                _ = v78
                            } else {
                                if (-78 != (fi74 ^ (-(fi74)))) {
                                    var v79 int32 = 1
                                    for fi80 := 0; fi80 < 7; fi80++ {
                                        v81 := v5[(fi80 ^ int(v79))]
                                        _ = fi80
                                        _ = v81
                                    }
                                    _ = v79
                                } else {
                                    v6 = (v6 - uint8(v1))
                                    for fi82 := 0; fi82 < 15; fi82++ {
                                        v3 = true
                                        switch fi74 {
                                        case 1:
                                            fmt.Println(v1, fi74)
                                            v4[v1] = false
                                            v6 -= (v6 + v6)
                                        case 83:
                                            if ((v6 > v6) && true) {
                                                switch fi74 {
                                                case 1:
                                                    v83 := GenS6{F0: false, F1: fi82}
                                                    _ = v83
                                                }
                                            } else {
                                                var v84 int = 0
                                                var v85 uint16 = 2
                                                var v86 float64 = ((3.14159 + (0.1 + ((1e10 - -2.25) - 0.1))) * float64(1.5))
                                                _ = v84
                                                _ = v85
                                                _ = v86
                                            }
                                        default:
                                            var v87 uint = v2
                                            _ = v87
                                        }
                                        _ = fi82
                                    }
                                    if (!v3) {
                                        var v88 int16 = 0
                                        _ = v88
                                    } else {
                                        v89 := GenS5{F0: 42, F1: 42, F2: 0}
                                        var v90 uint = v2
                                        for fi91 := 0; fi91 < 14; fi91++ {
                                            var v92 uint64 = 42
                                            fmt.Println(fi91)
                                            if v3 {
                                                var v93 string = ("foo bar" + "")
                                                _ = v93
                                            } else {
                                                var v94 string = "a"
                                                switch v94 {
                                                case "foo bar":
                                                    var v95 float32 = -2.25
                                                    _ = v95
                                                default:
                                                    var v96 string = v94
                                                    _ = v96
                                                }
                                                v6 = 143
                                                _ = v94
                                            }
                                            _ = fi91
                                            _ = v92
                                        }
                                        _ = v89
                                        _ = v90
                                    }
                                }
                                switch ("a" + "hello") {
                                case "\n\t":
                                    v6 += v6
                                }
                                fmt.Println(v2)
                            }
                        }
                    }
                    var v97 uint8 = v6
                    _ = fi74
                    _ = v97
                }
            case "":
                v98 := []int{-2147483648, -282}
                _ = v98
            case "\n\t":
                v1 = (v1 % ((1) | 1))
            }
        default:
            v99 := GenS1{F0: 118, F1: -487}
            _ = v99
        }
    case 127:
        fmt.Println(v3, v3, v3)
    case -474:
        var v100 uint32 = (v1 & (v1 & 65535))
        _ = v100
    }
    v101 := GenS6{F0: (!v3), F1: 234}
    for fi102 := 0; fi102 < 16; fi102++ {
        v6 = v6
        v3 = v3
        _ = fi102
    }
    var v103 uint32 = (v1 / ((0) | 1))
    v1 = 4294967295
    var v104 float64 = 0.0
    var v105 float64 = (0.1 - (3.14159 + 0.1))
    var v106 int16 = 42
    for fi107 := 0; fi107 < 18; fi107++ {
        switch "foo bar" {
        case "":
            v103 += v1
        case "a":
            fmt.Println(fi107, v105, v106)
        }
        _ = fi107
    }
    var v108 float64 = 1e10
    if ("a" > ("\n\t" + ("foo bar" + "a"))) {
        if v3 {
            for fi109 := 0; fi109 < 4; fi109++ {
                v3 = (792 >= 982)
                var v110 float64 = 1.5
                v111 := v5[(fi109 ^ fi109)]
                _ = fi109
                _ = v110
                _ = v111
            }
        } else {
            v6 -= 248
            v112 := v5[int(v1)]
            _ = v112
        }
        if (!(0.0 == (3.14159 / 1.0))) {
            fmt.Println(v105, v104, v6)
        } else {
            for fi113 := 0; fi113 < 18; fi113++ {
                for fi114 := 0; fi114 < 6; fi114++ {
                    v115 := GenS5{F0: -9223372036854775808, F1: 1000, F2: 42}
                    v103 = (v103 << (((v2 - 65535)) & 63))
                    var v116 int32 = 200
                    _ = fi114
                    _ = v115
                    _ = v116
                }
                _ = fi113
            }
            if ("" == "hello") {
                v117 := GenS7{F0: 1, F1: ("a" + "a"), F2: v3}
                v118 := []int{340, 1}
                _ = v117
                _ = v118
            }
            fmt.Println(v1, v2)
        }
        switch 212 {
        case 54:
            if (0 != 9223372036854775807) {
                v1 -= v1
            } else {
                if v3 {
                    v119 := GenS8{F0: 255}
                    _ = v119
                }
                v120 := GenS5{F0: 0, F1: -600, F2: -1}
                switch "hello" {
                case "hello":
                    v121 := GenS5{F0: -16, F1: -367, F2: 2147483647}
                    _ = v121
                default:
                    if ((v2 << (((v2 * v2)) & 63)) < (v2 << ((v2) & 63))) {
                        v122 := GenS9{F0: 1, F1: 255, F2: ("foo bar" + "")}
                        fmt.Println(v106)
                        _ = v122
                    } else {
                        var v123 uint8 = 255
                        v123 -= v123
                        var v124 int32 = 42
                        _ = v123
                        _ = v124
                    }
                }
                _ = v120
            }
        case -23:
            var v125 int = 514
            _ = v125
        default:
            v126 := GenS1{F0: v6, F1: 42}
            _ = v126
        }
    } else {
        var v127 int = 32768
        v103 = (v1 / (((v1 << ((v2) & 63))) | 1))
        _ = v127
    }
    for fi128 := 0; fi128 < 4; fi128++ {
        if true {
            var v129 uint64 = 0
            if v3 {
                var v130 int = fi128
                fmt.Println(v1, fi128)
                _ = v130
            } else {
                var v131 int32 = -1
                _ = v131
            }
            _ = v129
        }
        for fi132 := 0; fi132 < 7; fi132++ {
            switch "\n\t" {
            case "":
                v103 += 65535
            case "a":
                v133 := GenS7{F0: 366, F1: "a", F2: true}
                _ = v133
            case "foo bar":
                switch fi132 {
                case 9223372036854775807:
                    v134 := GenS10{F0: (-(fi132))}
                    _ = v134
                case 32768:
                    v3 = (!true)
                }
                if v3 {
                    v1 = (v103 - v1)
                    _ = len(v5)
                    v2 = v2
                }
            }
            _ = fi132
        }
        _ = fi128
    }
    if (33 == (^(v6))) {
        var v135 uint32 = 65535
        v103 = (v103 / ((116) | 1))
        for fi136 := 0; fi136 < 2; fi136++ {
            var v137 uint64 = 18446744073709551615
            _ = fi136
            _ = v137
        }
        _ = v135
    } else {
        v103 = v1
        fmt.Println(v3)
    }
    v3 = ((3.14159 - 3.14159) > (-2.25 + ((0.1 * float32((v108 - v105))) + 0.0)))
    fmt.Println("v1:", v1)
    fmt.Println("v103:", v103)
    fmt.Println("v104:", v104)
    fmt.Println("v105:", v105)
    fmt.Println("v106:", v106)
    fmt.Println("v108:", v108)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v101
    _ = v103
    _ = v104
    _ = v105
    _ = v106
    _ = v108
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
