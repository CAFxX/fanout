package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 string; F2 uint64 }
type GenS3 struct { F0 uint64; F1 int8; F2 int8 }
type GenS4 struct { F0 uint; F1 int8 }
type GenS5 struct { F0 int64; F1 int32 }
type GenS6 struct { F0 int64; F1 string; F2 int }
type GenS7 struct { F0 uint8; F1 uint }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int8; F1 uint }
func genFn0() int16 {
    var v5 uint16 = 885
    if true {
        if true {
            var v6 int32 = -2147483648
            var v7 int64 = 9223372036854775807
            v8 := []int{-180}
            _ = v6
            _ = v7
            _ = v8
        }
        v9 := []int{-1, 214, 488, 42}
        _ = v9
    }
    if false {
        var v10 int64 = 503
        v5 -= v5
        _ = v10
    } else {
        v11 := []int{}
        var v12 uint16 = (^(v5))
        _ = v11
        _ = v12
    }
    v13 := []int{190}
    _ = v13
    _ = v5
    return 518
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint32 = 4294967295
    v3 := make(map[string]int)
    var v4 string = ("a" + ("foo bar" + ""))
    var v14 int16 = 32767
    switch 2147483647 {
    case -684:
        v15 := GenS1{F0: 32767}
        _ = v15
    case 938:
        if false {
            v16 := GenS1{F0: 595}
            var v17 uint16 = 0
            _ = v16
            _ = v17
        }
    }
    v18 := []int{32767, -785, -9223372036854775808, -9223372036854775808}
    for fi19 := 0; fi19 < 13; fi19++ {
        _ = len(v18)
        if true {
            if false {
                for fi20 := 0; fi20 < 19; fi20++ {
                    var v21 int64 = -1
                    v4 = ((v4 + "hello") + "\n\t")
                    var v22 uint = 4294967295
                    v14 -= v14
                    var v23 uint16 = 438
                    _ = fi20
                    _ = v21
                    _ = v22
                    _ = v23
                }
                if false {
                    var v24 int64 = -1
                    _ = v24
                } else {
                    for fi25 := 0; fi25 < 7; fi25++ {
                        var v26 uint8 = 217
                        var v27 float64 = float64(3.14159)
                        fmt.Println(fi19, v27)
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                    switch v4 {
                    case "":
                        v4 = (("a" + v4) + "")
                    case "a":
                        if (!(976 > 2)) {
                            if true {
                                v2 += (^(v2))
                                var v28 int8 = 1
                                _ = v28
                            }
                            var v29 string = v4
                            var v30 int16 = (v14 & v14)
                            _ = v29
                            _ = v30
                        } else {
                            if (3.14159 > 0.1) {
                                switch (fi19 + (fi19 * 156)) {
                                case -116:
                                    if (!(-662 > 1)) {
                                        switch v4 {
                                        case "hello":
                                            v2 += v2
                                        case "foo bar":
                                            v2 -= 908
                                        case "\n\t":
                                            for fi31 := 0; fi31 < 10; fi31++ {
                                                v14 += (-(v14))
                                                var v32 bool = false
                                                _ = fi31
                                                _ = v32
                                            }
                                        default:
                                            if (!false) {
                                                var v33 uint8 = 184
                                                _ = v33
                                            } else {
                                                var v34 int8 = 91
                                                _ = v34
                                            }
                                        }
                                        if false {
                                            fmt.Println(v2)
                                        }
                                        for fi35 := 0; fi35 < 20; fi35++ {
                                            var v36 uint = 0
                                            if true {
                                                v4 = v4
                                                v37 := GenS2{F0: v14, F1: "\n\t", F2: 808}
                                                var v38 uint32 = v2
                                                v2 -= 65535
                                                _ = v37
                                                _ = v38
                                            }
                                            if len(v18) >= 2 {
                                                v18 = v18[1:len(v18)]
                                            }
                                            _ = fi35
                                            _ = v36
                                        }
                                    } else {
                                        v2 = v2
                                    }
                                case -2147483648:
                                    switch "foo bar" {
                                    case "a":
                                        for fi39 := 0; fi39 < 2; fi39++ {
                                            if (!(42 < 65535)) {
                                                _ = len(v3)
                                                v18 = append(v18, fi39)
                                            }
                                            if len(v18) > 0 {
                                                v18[0] = 1
                                            }
                                            if (-128 != 100) {
                                                v2 -= 42
                                            } else {
                                                v14 = (^(v14))
                                                v4 = v4
                                            }
                                            switch v4 {
                                            case "":
                                                v40 := GenS2{F0: (v14 % (((-(v14))) | 1)), F1: v4, F2: 1}
                                                _ = v40
                                            }
                                            v1[fi19] = 1
                                            _ = fi39
                                        }
                                    case "foo bar":
                                        if false {
                                            v4 = v4
                                            v41 := GenS2{F0: v14, F1: v4, F2: uint64(v2)}
                                            var v42 uint16 = 32767
                                            _ = v41
                                            _ = v42
                                        } else {
                                            for fi43 := 0; fi43 < 9; fi43++ {
                                                v14 = 1
                                                v4 = ((v4 + (v4 + ((v4 + v4) + "foo bar"))) + (v4 + v4))
                                                _ = fi43
                                            }
                                        }
                                    case "hello":
                                        var v44 uint8 = 2
                                        _ = v44
                                    }
                                case 0:
                                    v45 := GenS3{F0: 0, F1: 42, F2: -12}
                                    _ = v45
                                default:
                                    for fi46 := 0; fi46 < 19; fi46++ {
                                        v3[(v4 + (v4 + "foo bar"))] = int(v2)
                                        switch v4 {
                                        case "\n\t":
                                            v47 := GenS4{F0: 1, F1: -43}
                                            _ = v47
                                        case "a":
                                            v48 := GenS5{F0: -897, F1: 42}
                                            _ = v48
                                        case "":
                                            if true {
                                                var v49 int16 = v14
                                                _ = v49
                                            } else {
                                                var v50 float32 = 0.0
                                                var v51 uint64 = 823
                                                var v52 uint = 2
                                                _ = v50
                                                _ = v51
                                                _ = v52
                                            }
                                        }
                                        var v53 string = v4
                                        _ = fi46
                                        _ = v53
                                    }
                                }
                            }
                            if (false && (!false)) {
                                var v54 int64 = 0
                                v14 = (-(v14))
                                _ = v54
                            }
                            v55 := GenS3{F0: 2, F1: 94, F2: 42}
                            _ = v55
                        }
                    case "foo bar":
                        var v56 int32 = -2147483648
                        _ = v56
                    default:
                        v57 := GenS1{F0: 255}
                        _ = v57
                    }
                    v14 += -518
                }
                var v58 uint64 = 1
                _ = v58
            } else {
                _ = len(v1)
            }
            switch (^(fi19)) {
            case -2147483648:
                switch ((((v4 + "a") + "hello") + "") + v4) {
                case "foo bar":
                    v2 -= v2
                default:
                    for fi59 := 0; fi59 < 12; fi59++ {
                        v14 -= 0
                        if false {
                            var v60 int = (fi59 >> ((18446744073709551615) & 63))
                            _ = v60
                        }
                        v2 = v2
                        if false {
                            var v61 int32 = 0
                            v2 = (v2 ^ (v2 << ((2) & 63)))
                            _ = v61
                        }
                        v2 = v2
                        _ = fi59
                    }
                }
            default:
                if ((v4 + v4) <= v4) {
                    v4 = ("hello" + (v4 + "\n\t"))
                } else {
                    for fi62 := 0; fi62 < 19; fi62++ {
                        switch (v4 + ((((v4 + v4) + v4) + v4) + "foo bar")) {
                        case "\n\t":
                            v4 = (v4 + (("hello" + "\n\t") + ((v4 + v4) + v4)))
                        }
                        switch "a" {
                        case "\n\t":
                            v63 := GenS3{F0: 1, F1: -128, F2: 42}
                            v4 = v4
                            var v64 int16 = v14
                            switch (fi62 ^ 9223372036854775807) {
                            case 986:
                                var v65 float64 = ((0.0 + 3.14159) - -2.25)
                                fmt.Println(v65, v14)
                                if (192 < 1) {
                                    v4 = (("hello" + (((v4 + ("\n\t" + v4)) + v4) + v4)) + v4)
                                    var v66 uint8 = 197
                                    var v67 int64 = 0
                                    _ = v66
                                    _ = v67
                                }
                                _ = v65
                            case 0:
                                var v68 string = v4
                                if true {
                                    fmt.Println(v68, v4, v64)
                                    var v69 string = ((v68 + ("" + v68)) + v4)
                                    v70 := GenS6{F0: -503, F1: "", F2: 4294967295}
                                    _ = v69
                                    _ = v70
                                } else {
                                    v68 = "foo bar"
                                }
                                fmt.Println(v64, v14)
                                _ = v68
                            case 127:
                                var v71 int = fi62
                                _ = v71
                            }
                            _ = len(v18)
                            if true {
                                fmt.Println(v64, v14, v2)
                                if len(v18) > 0 {
                                    v18[0] = fi62
                                }
                            } else {
                                var v72 string = "foo bar"
                                _ = v72
                            }
                            var v73 float64 = 0.0
                            v74 := GenS6{F0: -1, F1: v4, F2: fi62}
                            v14 -= (v64 - v14)
                            v14 = genFn0()
                            if (-128 <= 0) {
                                if (-2147483648 == (fi19 & (-(fi62)))) {
                                    var v75 int32 = 1
                                    _ = v75
                                }
                                var v76 uint64 = 560
                                _ = len(v18)
                                _ = v76
                            }
                            _ = v63
                            _ = v64
                            _ = v73
                            _ = v74
                        case "":
                            var v77 uint16 = 1
                            _ = v77
                        }
                        v4 = ("a" + "")
                        _ = fi62
                    }
                    var v78 float64 = 1.5
                    _ = v78
                }
            }
            v79 := GenS7{F0: 127, F1: 2}
            _ = v79
        }
        _ = fi19
    }
    for fi80 := 0; fi80 < 17; fi80++ {
        if (84 < -2147483648) {
            v81 := GenS3{F0: 0, F1: 1, F2: 42}
            switch fi80 {
            case -9223372036854775808:
                v4 = (v4 + "")
            }
            _ = v81
        }
        switch (v4 + "a") {
        case "hello":
            if false {
                var v82 uint64 = 42
                for fi83 := 0; fi83 < 19; fi83++ {
                    if true {
                        var v84 int32 = -320
                        v4 = "foo bar"
                        _ = v84
                    } else {
                        switch ((v4 + v4) + v4) {
                        case "":
                            var v85 string = v4
                            _ = v85
                        default:
                            if len(v18) > 0 {
                                v18[0] = 32768
                            }
                        }
                        var v86 uint32 = (v2 >> ((2) & 63))
                        v87 := GenS8{F0: 255}
                        v88 := v3[(v4 + (v4 + v4))]
                        _ = v86
                        _ = v87
                        _ = v88
                    }
                    var v89 string = v4
                    fmt.Println(fi83)
                    v89 = v4
                    _ = fi83
                    _ = v89
                }
                var v90 uint = 0
                _ = v82
                _ = v90
            }
        case "a":
            switch "\n\t" {
            case "hello":
                fmt.Println(fi80)
            case "\n\t":
                var v91 string = v4
                _ = v91
            default:
                v4 = v4
            }
        case "\n\t":
            for fi92 := 0; fi92 < 4; fi92++ {
                _ = len(v3)
                if false {
                    delete(v1, (fi80 - (fi80 << ((255) & 63))))
                    var v93 int64 = 9223372036854775807
                    _ = v93
                } else {
                    var v94 int16 = 1
                    fmt.Println(fi92)
                    var v95 int8 = 70
                    _ = v94
                    _ = v95
                }
                v4 = ((("a" + ("" + v4)) + v4) + v4)
                v14 -= (^(v14))
                _ = fi92
            }
        }
        fmt.Println(v2, fi80)
        _ = fi80
    }
    if (-2.25 <= 0.1) {
        if true {
            var v96 int64 = 0
            if (false && false) {
                v97 := GenS8{F0: 572}
                v4 = (v4 + v4)
                _ = v97
            } else {
                var v98 uint16 = 42
                if false {
                    switch -655 {
                    case 0:
                        v98 = (v98 + v98)
                    case 410:
                        switch 573 {
                        case -205:
                            for fi99 := 0; fi99 < 18; fi99++ {
                                for fi100 := 0; fi100 < 9; fi100++ {
                                    _ = len(v18)
                                    _ = fi100
                                }
                                v96 = v96
                                for fi101 := 0; fi101 < 16; fi101++ {
                                    var v102 float64 = float64(0.0)
                                    var v103 float64 = (v102 * 0.0)
                                    var v104 float32 = 3.14159
                                    var v105 uint32 = 580
                                    var v106 int32 = 2147483647
                                    _ = fi101
                                    _ = v102
                                    _ = v103
                                    _ = v104
                                    _ = v105
                                    _ = v106
                                }
                                _ = fi99
                            }
                        case -2147483648:
                            var v107 float32 = ((((0.1 / (1e10 * 3.14159)) + 1e10) / 1e10) / 0.1)
                            _ = v107
                        default:
                            v108 := GenS7{F0: 95, F1: 1}
                            _ = v108
                        }
                    case 128:
                        var v109 int16 = (-(v14))
                        _ = v109
                    default:
                        if (2 <= 1) {
                            if false {
                                for fi110 := 0; fi110 < 16; fi110++ {
                                    for fi111 := 0; fi111 < 7; fi111++ {
                                        v14 -= v14
                                        v4 = "a"
                                        if true {
                                            fmt.Println(v14)
                                            v112 := GenS7{F0: 101, F1: 255}
                                            if ((("\n\t" + v4) + v4) != (v4 + (((v4 + ("hello" + (("foo bar" + (((v4 + (("foo bar" + v4) + v4)) + v4) + "hello")) + v4))) + ((v4 + (((((("" + "a") + "a") + "a") + v4) + v4) + "foo bar")) + v4)) + "a"))) {
                                                v2 = v2
                                                delete(v3, (v4 + (v4 + v4)))
                                                v98 -= 255
                                            }
                                            fmt.Println(v4, fi110)
                                            _ = v112
                                        }
                                        var v113 int32 = 759
                                        switch fi111 {
                                        case -9223372036854775808:
                                            fmt.Println(fi110, v98, v96)
                                        case -97:
                                            var v114 float32 = 1e10
                                            _ = v114
                                        }
                                        _ = fi111
                                        _ = v113
                                    }
                                    var v115 float32 = (0.1 + (-2.25 * 0.0))
                                    _ = fi110
                                    _ = v115
                                }
                            }
                            fmt.Println(v4, v96)
                        } else {
                            switch v4 {
                            case "\n\t":
                                var v116 uint16 = (v98 & v98)
                                _ = v116
                            case "":
                                switch 0 {
                                case 9223372036854775807:
                                    var v117 float64 = ((float64((3.14159 - float32(-2.25))) + (-2.25 / ((0.0 - (1e10 / (((0.0 * (-2.25 + -2.25)) * ((0.0 - 0.1) + 1e10)) - (((0.1 / 3.14159) / (3.14159 / (0.1 / 1.5))) + (0.1 / ((float64(((-2.25 / (float32(0.1) / 1e10)) / 0.1)) / 0.1) / -2.25)))))) - (-2.25 / 1e10)))) + 1.5)
                                    _ = v117
                                case 0:
                                    var v118 int32 = -1
                                    _ = v118
                                case 127:
                                    var v119 uint16 = (v98 >> ((255) & 63))
                                    _ = v119
                                default:
                                    v120 := v3[(v4 + (("" + v4) + v4))]
                                    _ = v120
                                }
                            }
                            if (((v2 - (v2 / (((^(v2))) | 1))) >= uint32(v98)) && (!((v2 >> ((0) & 63)) < (v2 & v2)))) {
                                var v121 float64 = 1.5
                                _ = v121
                            } else {
                                if len(v18) >= 2 {
                                    v18 = v18[1:len(v18)]
                                }
                                v14 -= genFn0()
                                v4 = ""
                            }
                            fmt.Println(v14)
                        }
                    }
                }
                _ = v98
            }
            v4 = ""
            _ = v96
        }
        switch 404 {
        case 4294967295:
            if false {
                for fi122 := 0; fi122 < 17; fi122++ {
                    if false {
                        var v123 float32 = -2.25
                        _ = v123
                    } else {
                        switch fi122 {
                        case -109:
                            fmt.Println(fi122)
                        case -2147483648:
                            switch (fi122 * fi122) {
                            case -747:
                                v4 = v4
                            default:
                                if len(v18) > 0 {
                                    v18[0] = (fi122 % ((32768) | 1))
                                }
                            }
                        case 9223372036854775807:
                            for fi124 := 0; fi124 < 12; fi124++ {
                                v125 := GenS5{F0: -1, F1: 2147483647}
                                v4 = ""
                                v14 += (v14 & v14)
                                v126 := GenS2{F0: genFn0(), F1: "\n\t", F2: 0}
                                _ = fi124
                                _ = v125
                                _ = v126
                            }
                        default:
                            v14 = -403
                        }
                    }
                    if (-684 <= 2147483647) {
                        switch (v4 + ("\n\t" + v4)) {
                        case "\n\t":
                            switch v4 {
                            case "foo bar":
                                v4 = v4
                            }
                        case "hello":
                            v127 := GenS5{F0: 9223372036854775807, F1: 2147483647}
                            _ = v127
                        default:
                            _ = len(v3)
                        }
                        var v128 bool = (1.5 < 1e10)
                        _ = v128
                    }
                    v3[("hello" + "\n\t")] = fi122
                    _ = fi122
                }
                v14 = (v14 - genFn0())
            }
        default:
            var v129 uint = 18446744073709551615
            _ = v129
        }
    } else {
        for fi130 := 0; fi130 < 16; fi130++ {
            v2 += v2
            var v131 float32 = 1.5
            var v132 string = (("foo bar" + (v4 + v4)) + (v4 + "a"))
            _ = fi130
            _ = v131
            _ = v132
        }
    }
    var v133 float32 = ((((3.14159 - 0.0) * 0.1) * -2.25) / 1e10)
    fmt.Println(v133)
    v3["hello"] = -2147483648
    _ = len(v1)
    if ((v4 + v4) != (v4 + v4)) {
        _ = len(v18)
        if (v2 >= v2) {
            v133 = 0.1
            fmt.Println(v14, v14)
        }
        if len(v18) > 0 {
            v18[0] = -772
        }
    } else {
        if len(v18) >= 2 {
            v18 = v18[1:len(v18)]
        }
    }
    switch (v4 + (v4 + "a")) {
    case "foo bar":
        var v134 uint16 = 255
        _ = v134
    case "a":
        v135 := GenS9{F0: 84, F1: 42}
        _ = v135
    case "":
        var v136 int64 = 456
        _ = v136
    default:
        if true {
            v137 := GenS3{F0: 18446744073709551615, F1: -1, F2: -1}
            v138 := GenS6{F0: -1, F1: (v4 + v4), F2: 2147483647}
            _ = v137
            _ = v138
        }
    }
    v139 := GenS1{F0: 42}
    fmt.Println(v133)
    var v140 bool = (!false)
    if len(v18) >= 2 {
        v18 = v18[1:len(v18)]
    }
    for fi141 := 0; fi141 < 12; fi141++ {
        if v140 {
            v2 = (v2 >> ((1) & 63))
        }
        _ = fi141
    }
    if (255 > 115) {
        _ = len(v3)
        v142 := GenS6{F0: 42, F1: ((v4 + (v4 + v4)) + (v4 + v4)), F2: -2147483648}
        _ = len(v3)
        _ = v142
    } else {
        var v143 uint = 4294967295
        _ = v143
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v133:", v133)
    fmt.Println("v14:", v14)
    fmt.Println("v140:", v140)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v133
    _ = v139
    _ = v14
    _ = v140
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
