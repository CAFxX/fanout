package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 uint8 }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 string; F1 uint64 }
type GenS5 struct { F0 uint8; F1 int32 }
type GenS6 struct { F0 int8; F1 uint16; F2 string }
type GenS7 struct { F0 uint8; F1 int64; F2 int }
type GenS8 struct { F0 uint32; F1 uint; F2 uint8 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 uint8; F1 uint32 }
type GenS12 struct { F0 uint32; F1 uint32 }
type GenS13 struct { F0 int64; F1 uint; F2 int }
type GenS14 struct { F0 uint8; F1 int }
func genFn0() uint {
    if true {
        for fi4 := 0; fi4 < 4; fi4++ {
            var v5 int = 4294967295
            _ = fi4
            _ = v5
        }
        for fi6 := 0; fi6 < 12; fi6++ {
            var v7 uint16 = 962
            _ = fi6
            _ = v7
        }
        if true {
            if (("hello" >= (("foo bar" + "hello") + "foo bar")) && false) {
                v8 := []int{}
                var v9 int16 = -605
                v9 = v9
                _ = v8
                _ = v9
            }
        } else {
            var v10 uint32 = 825
            v11 := []int{-480}
            v11 = append(v11, 4294967295)
            _ = v10
            _ = v11
        }
    }
    var v12 uint = 0
    if true {
        v13 := []int{32768, -1, 0, -2147483648}
        _ = v13
    } else {
        if (1 < 1) {
            v14 := GenS1{F0: 1, F1: 2, F2: 0}
            if ((false || (2147483648 <= 15)) && false) {
                var v15 uint8 = 0
                v16 := []int{32767, 167, 127}
                _ = v15
                _ = v16
            } else {
                fmt.Println(v12)
                if true {
                    v17 := []int{-2147483648, 1, 32767, 4294967295}
                    v18 := GenS1{F0: 4294967295, F1: 32767, F2: 255}
                    if len(v17) >= 2 {
                        v17 = v17[1:len(v17)]
                    }
                    _ = v17
                    _ = v18
                }
                for fi19 := 0; fi19 < 13; fi19++ {
                    v12 = 18446744073709551615
                    v20 := GenS2{F0: 620, F1: 91}
                    var v21 uint64 = 0
                    _ = fi19
                    _ = v20
                    _ = v21
                }
            }
            if true {
                if (255 > 255) {
                    var v22 int = 128
                    _ = v22
                }
            }
            _ = v14
        } else {
            if (42 >= 1) {
                v23 := []int{127}
                var v24 int8 = 127
                for fi25 := 0; fi25 < 3; fi25++ {
                    switch (fi25 | (fi25 + (-(fi25)))) {
                    case -2147483648:
                        var v26 float32 = -2.25
                        _ = v26
                    default:
                        fmt.Println(fi25, v24, fi25)
                    }
                    v27 := GenS3{F0: -128}
                    for fi28 := 0; fi28 < 9; fi28++ {
                        v29 := GenS2{F0: 2, F1: (v24 ^ -128)}
                        var v30 uint8 = 28
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                    _ = fi25
                    _ = v27
                }
                _ = v23
                _ = v24
            } else {
                v31 := GenS4{F0: "hello", F1: 42}
                _ = v31
            }
            var v32 uint64 = 632
            v12 += (^(v12))
            _ = v32
        }
        if false {
            for fi33 := 0; fi33 < 16; fi33++ {
                for fi34 := 0; fi34 < 8; fi34++ {
                    var v35 int8 = 0
                    v12 -= v12
                    v12 = 1
                    _ = fi34
                    _ = v35
                }
                switch 128 {
                case 240:
                    for fi36 := 0; fi36 < 19; fi36++ {
                        var v37 uint64 = 0
                        v12 = (v12 & v12)
                        if (fi36 == 788) {
                            var v38 int64 = 0
                            v37 += (v37 | 240)
                            v39 := []int{fi36}
                            _ = v38
                            _ = v39
                        } else {
                            v37 -= (^(v37))
                        }
                        _ = fi36
                        _ = v37
                    }
                default:
                    for fi40 := 0; fi40 < 4; fi40++ {
                        v12 = (v12 & 197)
                        var v41 uint = (v12 / ((2) | 1))
                        _ = fi40
                        _ = v41
                    }
                }
                _ = fi33
            }
            if ((true && (((!(958 < 42)) || (!(0 == 0))) && false)) && true) {
                var v42 uint64 = 1
                _ = v42
            } else {
                v43 := GenS5{F0: 127, F1: 2147483647}
                v44 := []int{-144, 641, 4294967295}
                for fi45 := 0; fi45 < 10; fi45++ {
                    v46 := GenS3{F0: 2}
                    v12 = v12
                    v47 := GenS6{F0: 1, F1: 943, F2: ("hello" + ("\n\t" + "foo bar"))}
                    _ = fi45
                    _ = v46
                    _ = v47
                }
                _ = v43
                _ = v44
            }
            v12 += (v12 & 255)
        }
        var v48 uint16 = 1
        _ = v48
    }
    v12 -= 505
    _ = v12
    return v12
}

func main() {
    var v1 uint64 = 822
    var v2 int8 = -81
    v3 := make(map[string]string)
    switch (("a" + "") + "foo bar") {
    case "a":
        v49 := []int{}
        _ = v49
    case "hello":
        var v50 uint8 = 255
        _ = v50
    case "":
        delete(v3, "")
    default:
        if true {
            var v51 uint64 = v1
            delete(v3, "a")
            _ = v51
        } else {
            v52 := []int{481, 216, -899}
            _ = v52
        }
    }
    if ((^(v2)) >= (^(v2))) {
        for fi53 := 0; fi53 < 19; fi53++ {
            if true {
                delete(v3, ("" + "a"))
                v54 := GenS1{F0: 42, F1: 255, F2: 1}
                _ = v54
            } else {
                if ((!true) || (1000 == -963)) {
                    var v55 bool = ("hello" >= "foo bar")
                    _ = v55
                } else {
                    var v56 int32 = 42
                    if (0 <= 1000) {
                        switch (^(fi53)) {
                        case -2147483648:
                            var v57 uint8 = 127
                            _ = v57
                        case 127:
                            for fi58 := 0; fi58 < 19; fi58++ {
                                if (!(true && (!false))) {
                                    switch "hello" {
                                    case "":
                                        fmt.Println(v1, fi53)
                                    case "\n\t":
                                        v59 := GenS6{F0: v2, F1: 32767, F2: "a"}
                                        _ = v59
                                    case "hello":
                                        var v60 uint = 451
                                        _ = v60
                                    default:
                                        switch ("" + "hello") {
                                        case "":
                                            switch fi58 {
                                            case 42:
                                                v56 -= (v56 << ((255) & 63))
                                            case -1:
                                                v61 := []int{fi58, (fi53 | (^(fi58)))}
                                                if false {
                                                    var v62 int64 = -81
                                                    var v63 uint = 18446744073709551615
                                                    fmt.Println(fi53)
                                                    _ = v62
                                                    _ = v63
                                                } else {
                                                    switch (fi58 | fi58) {
                                                    case 1:
                                                        v56 = (-(v56))
                                                    case 2147483647:
                                                        v2 = 1
                                                        switch "foo bar" {
                                                        case "hello":
                                                            var v64 int16 = 42
                                                            _ = v64
                                                        }
                                                        if len(v61) >= 2 {
                                                            v61 = v61[1:len(v61)]
                                                        }
                                                    case -263:
                                                        var v65 int64 = 0
                                                        _ = v65
                                                    }
                                                    switch fi58 {
                                                    case 9223372036854775807:
                                                        var v66 int8 = (-(v2))
                                                        v1 = 4294967295
                                                        _ = v66
                                                    default:
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            if (1 > 9223372036854775807) {
                                                                var v67 uint = 376
                                                                if true {
                                                                    fmt.Println(v2, fi58, v2)
                                                                    if (42 < 764) {
                                                                        if true {
                                                                            if true {
                                                                                v56 = (v56 << ((v67) & 63))
                                                                                fmt.Println(v2, v67, v2)
                                                                                delete(v3, "a")
                                                                            } else {
                                                                                _ = len(v61)
                                                                                switch ("\n\t" + "hello") {
                                                                                case "":
                                                                                    switch 2147483647 {
                                                                                    case 735:
                                                                                        var v68 uint = (^(v67))
                                                                                        _ = v68
                                                                                    case 0:
                                                                                        var v69 float64 = 0.0
                                                                                        _ = v69
                                                                                    case 128:
                                                                                        v56 += -357
                                                                                    default:
                                                                                        var v70 int64 = -1
                                                                                        _ = v70
                                                                                    }
                                                                                default:
                                                                                    if len(v61) >= 2 {
                                                                                        v61 = v61[1:len(v61)]
                                                                                    }
                                                                                }
                                                                            }
                                                                            var v71 float64 = 1e10
                                                                            _ = v71
                                                                        }
                                                                        _ = len(v3)
                                                                        var v72 int8 = (v2 + -128)
                                                                        if true {
                                                                            v72 += v72
                                                                            v73 := GenS7{F0: 42, F1: 9223372036854775807, F2: (fi58 - fi58)}
                                                                            _ = v73
                                                                        }
                                                                        v3[""] = "foo bar"
                                                                        _ = v72
                                                                    } else {
                                                                        if ((v1 % ((0) | 1)) == (v1 >> (((v67 >> (((v67 & v67)) & 63))) & 63))) {
                                                                            fmt.Println(v67, v1, v2)
                                                                            if true {
                                                                                var v74 uint8 = 1
                                                                                v3["a"] = "a"
                                                                                _ = v74
                                                                            }
                                                                            if (208 != 1) {
                                                                                v75 := v3["a"]
                                                                                v67 -= v67
                                                                                _ = v75
                                                                            }
                                                                        }
                                                                        switch "\n\t" {
                                                                        case "a":
                                                                            v1 = v1
                                                                        case "hello":
                                                                            v76 := GenS4{F0: ("" + "\n\t"), F1: (^(v1))}
                                                                            var v77 string = ("hello" + ("a" + "hello"))
                                                                            var v78 bool = (!false)
                                                                            _ = v76
                                                                            _ = v77
                                                                            _ = v78
                                                                        }
                                                                        var v79 uint = genFn0()
                                                                        _ = v79
                                                                    }
                                                                    var v80 uint16 = 42
                                                                    _ = v80
                                                                }
                                                                var v81 float32 = ((-2.25 + ((0.1 * 3.14159) * 3.14159)) * 0.0)
                                                                _ = v67
                                                                _ = v81
                                                            }
                                                        case "a":
                                                            if true {
                                                                fmt.Println(fi53)
                                                                v82 := GenS6{F0: -1, F1: 32767, F2: "hello"}
                                                                _ = v82
                                                            }
                                                        }
                                                    }
                                                    if len(v61) > 0 {
                                                        v61[0] = (fi53 | 128)
                                                    }
                                                    switch -2147483648 {
                                                    case 72:
                                                        switch 32767 {
                                                        case 127:
                                                            switch 32768 {
                                                            case -1:
                                                                _ = len(v61)
                                                            case -9223372036854775808:
                                                                var v83 float64 = float64(3.14159)
                                                                _ = v83
                                                            case -257:
                                                                if ("" <= ("" + ((((("" + "a") + "\n\t") + "a") + "\n\t") + "\n\t"))) {
                                                                    if (!(2 >= 32767)) {
                                                                        var v84 uint8 = 42
                                                                        _ = v84
                                                                    }
                                                                    var v85 float32 = 1.5
                                                                    _ = v85
                                                                } else {
                                                                    v2 -= v2
                                                                    switch ("foo bar" + "a") {
                                                                    case "\n\t":
                                                                        v86 := v3["\n\t"]
                                                                        _ = v86
                                                                    case "foo bar":
                                                                        var v87 uint64 = (v1 << ((42) & 63))
                                                                        _ = v87
                                                                    }
                                                                }
                                                            }
                                                        case 2147483647:
                                                            if true {
                                                                if (!((v56 % ((2147483647) | 1)) != v56)) {
                                                                    v88 := GenS8{F0: 957, F1: 4294967295, F2: 122}
                                                                    v56 = 53
                                                                    fmt.Println(v1, v1, v1)
                                                                    v89 := GenS8{F0: 0, F1: 2, F2: 239}
                                                                    var v90 uint32 = 716
                                                                    _ = v88
                                                                    _ = v89
                                                                    _ = v90
                                                                }
                                                                if true {
                                                                    v91 := v3[("\n\t" + "\n\t")]
                                                                    _ = len(v3)
                                                                    var v92 uint = genFn0()
                                                                    _ = v91
                                                                    _ = v92
                                                                }
                                                                if true {
                                                                    v93 := GenS1{F0: 2, F1: 683, F2: 127}
                                                                    _ = v93
                                                                }
                                                            }
                                                        default:
                                                            v2 = (v2 | -123)
                                                        }
                                                    case -1:
                                                        var v94 float64 = 3.14159
                                                        _ = v94
                                                    default:
                                                        v2 -= v2
                                                    }
                                                    var v95 float32 = float32(0.0)
                                                    _ = v95
                                                }
                                                _ = v61
                                            case 127:
                                                v2 -= (v2 | (v2 << ((780) & 63)))
                                            }
                                        case "a":
                                            if ((false || false) && true) {
                                                var v96 int64 = -1
                                                v97 := GenS1{F0: 865, F1: 255, F2: 42}
                                                _ = v96
                                                _ = v97
                                            }
                                        default:
                                            fmt.Println(fi58, v2)
                                        }
                                    }
                                    v1 -= 877
                                } else {
                                    var v98 bool = ((0 == -1) || (128 <= 128))
                                    _ = v98
                                }
                                switch "" {
                                case "foo bar":
                                    var v99 bool = true
                                    _ = v99
                                default:
                                    var v100 int16 = int16(fi53)
                                    _ = v100
                                }
                                _ = fi58
                            }
                        default:
                            var v101 int32 = v56
                            _ = v101
                        }
                    } else {
                        v2 = v2
                        fmt.Println(v1, fi53, fi53)
                    }
                    _ = v56
                }
                v1 = 346
                if (true && (2147483648 < 2)) {
                    v3["\n\t"] = ("" + ("foo bar" + (("hello" + (((("hello" + "a") + ("hello" + "hello")) + "a") + "a")) + "")))
                    v2 += v2
                }
            }
            switch "" {
            case "a":
                v102 := GenS9{F0: v2}
                _ = v102
            case "":
                v2 = (v2 ^ v2)
            case "hello":
                var v103 int = 0
                _ = v103
            }
            var v104 float64 = 0.1
            _ = fi53
            _ = v104
        }
        switch "\n\t" {
        case "\n\t":
            for fi105 := 0; fi105 < 13; fi105++ {
                if (!(42 <= 212)) {
                    var v106 uint64 = v1
                    _ = v106
                } else {
                    fmt.Println(v2)
                    var v107 int32 = -4
                    _ = v107
                }
                if true {
                    if (((0.1 + 1e10) <= ((3.14159 + (-2.25 - 1e10)) * -2.25)) && (-2.25 > float32(((-2.25 / 3.14159) * (float64(float32(float64(1.5))) + ((1e10 * (0.1 * -2.25)) - 3.14159)))))) {
                        switch fi105 {
                        case 127:
                            if (42 >= 1) {
                                if true {
                                    var v108 uint16 = 987
                                    v108 = v108
                                    _ = v108
                                }
                            }
                        case -9223372036854775808:
                            if false {
                                if (!false) {
                                    if (true && false) {
                                        switch (fi105 * fi105) {
                                        case 128:
                                            var v109 int64 = 376
                                            _ = v109
                                        case -1:
                                            for fi110 := 0; fi110 < 18; fi110++ {
                                                v3[(("" + ("hello" + "")) + "a")] = ""
                                                _ = fi110
                                            }
                                        case 32767:
                                            for fi111 := 0; fi111 < 9; fi111++ {
                                                if true {
                                                    if (0 != 657) {
                                                        v112 := GenS1{F0: 30, F1: 65535, F2: 40}
                                                        fmt.Println(fi105)
                                                        _ = v112
                                                    }
                                                    v2 = 42
                                                    delete(v3, (("\n\t" + ("hello" + ("a" + ""))) + "\n\t"))
                                                } else {
                                                    delete(v3, "\n\t")
                                                    var v113 int32 = 701
                                                    _ = v113
                                                }
                                                var v114 uint16 = 1
                                                var v115 uint64 = v1
                                                _ = fi111
                                                _ = v114
                                                _ = v115
                                            }
                                        default:
                                            v2 = (v2 + v2)
                                        }
                                        v2 = (v2 >> ((162) & 63))
                                    } else {
                                        var v116 int64 = -9223372036854775808
                                        var v117 bool = false
                                        _ = v116
                                        _ = v117
                                    }
                                    switch fi105 {
                                    case 42:
                                        var v118 uint64 = (v1 ^ v1)
                                        _ = v118
                                    }
                                } else {
                                    if true {
                                        v119 := v3["\n\t"]
                                        for fi120 := 0; fi120 < 15; fi120++ {
                                            switch (v119 + "\n\t") {
                                            case "hello":
                                                v1 = 18446744073709551615
                                            }
                                            _ = fi120
                                        }
                                        v121 := GenS10{F0: (v2 + v2)}
                                        _ = v119
                                        _ = v121
                                    }
                                    switch -2147483648 {
                                    case 2147483647:
                                        v122 := GenS6{F0: v2, F1: 42, F2: "hello"}
                                        _ = v122
                                    case -9223372036854775808:
                                        v1 = 108
                                    case 32767:
                                        var v123 int16 = 32767
                                        _ = v123
                                    }
                                    v124 := []int{fi105, (fi105 * (fi105 / ((fi105) | 1))), -9223372036854775808}
                                    _ = v124
                                }
                                var v125 uint16 = 255
                                _ = v125
                            } else {
                                if true {
                                    var v126 int32 = 42
                                    v127 := GenS10{F0: v2}
                                    _ = v126
                                    _ = v127
                                } else {
                                    var v128 int32 = 0
                                    _ = v128
                                }
                                fmt.Println(v1, v1)
                            }
                        case -812:
                            v2 = (v2 | v2)
                        }
                    }
                }
                _ = fi105
            }
        }
    } else {
        for fi129 := 0; fi129 < 10; fi129++ {
            var v130 uint16 = 65535
            v131 := GenS6{F0: v2, F1: v130, F2: "\n\t"}
            if true {
                var v132 uint = 786
                _ = v132
            } else {
                delete(v3, "foo bar")
                var v133 uint8 = 255
                v134 := GenS1{F0: 778, F1: v130, F2: 240}
                _ = v133
                _ = v134
            }
            _ = fi129
            _ = v130
            _ = v131
        }
        if (-753 <= -2147483648) {
            var v135 bool = true
            for fi136 := 0; fi136 < 10; fi136++ {
                if v135 {
                    for fi137 := 0; fi137 < 3; fi137++ {
                        delete(v3, "a")
                        if (111 == (^(fi136))) {
                            v2 = (v2 >> ((1) & 63))
                            var v138 uint64 = (v1 / (((v1 << ((69) & 63))) | 1))
                            v139 := GenS6{F0: 5, F1: 65535, F2: "a"}
                            _ = v138
                            _ = v139
                        } else {
                            v2 += v2
                            if (0 >= 42) {
                                var v140 int64 = -9223372036854775808
                                _ = v140
                            } else {
                                if (((1 != 42) && false) && (!((v2 * (v2 + v2)) > -64))) {
                                    v141 := []int{(fi136 % ((fi136) | 1)), 614, 778}
                                    _ = v141
                                } else {
                                    var v142 int = fi137
                                    _ = v142
                                }
                                if v135 {
                                    fmt.Println(fi136, v135, fi136)
                                    v143 := []int{127}
                                    fmt.Println(fi137, v2, fi137)
                                    _ = v143
                                }
                                v144 := []int{fi137}
                                v144 = append(v144, (^(fi136)))
                                _ = v144
                            }
                            var v145 int64 = 1
                            _ = v145
                        }
                        if true {
                            v146 := GenS5{F0: 238, F1: 2147483647}
                            _ = v146
                        }
                        _ = fi137
                    }
                    v135 = v135
                } else {
                    v147 := GenS11{F0: 127, F1: 0}
                    switch "" {
                    case "\n\t":
                        var v148 float32 = (1.5 - -2.25)
                        _ = v148
                    case "foo bar":
                        fmt.Println(v135)
                    default:
                        v149 := GenS8{F0: 65535, F1: 902, F2: 255}
                        _ = v149
                    }
                    var v150 uint8 = 127
                    _ = v147
                    _ = v150
                }
                _ = fi136
            }
            for fi151 := 0; fi151 < 5; fi151++ {
                switch fi151 {
                case 127:
                    if (v135 && (65535 > 42)) {
                        var v152 int64 = -9223372036854775808
                        v153 := []int{(fi151 >> ((0) & 63)), fi151, fi151}
                        _ = v152
                        _ = v153
                    }
                }
                if (("" + (("\n\t" + ("a" + "hello")) + ("hello" + (("" + (("" + "foo bar") + ("foo bar" + ""))) + ("a" + "\n\t"))))) > "\n\t") {
                    var v154 uint8 = 206
                    v2 += v2
                    _ = v154
                } else {
                    var v155 string = (("a" + "a") + ((((("hello" + "") + "") + "hello") + ((("\n\t" + "hello") + "\n\t") + (("a" + "foo bar") + ("hello" + ("\n\t" + ("foo bar" + "\n\t")))))) + ("hello" + "\n\t")))
                    for fi156 := 0; fi156 < 6; fi156++ {
                        fmt.Println(fi156)
                        switch (-(fi151)) {
                        case 92:
                            v3[v155] = v155
                        case 150:
                            delete(v3, v155)
                        case 176:
                            var v157 uint = 255
                            v158 := GenS8{F0: 2, F1: (v157 - 18446744073709551615), F2: 102}
                            _ = v157
                            _ = v158
                        }
                        _ = fi156
                    }
                    _ = v155
                }
                _ = fi151
            }
            _ = v135
        }
        v2 = -1
    }
    v159 := GenS7{F0: 42, F1: 1, F2: 127}
    switch -7 {
    case -2147483648:
        if (!(-23 <= (v2 - (v2 >> ((42) & 63))))) {
            fmt.Println(v1, v2)
        }
    default:
        if (!(!(2 == v1))) {
            if ((v1 << ((255) & 63)) <= v1) {
                var v160 uint16 = 65535
                _ = v160
            }
            for fi161 := 0; fi161 < 16; fi161++ {
                for fi162 := 0; fi162 < 14; fi162++ {
                    if (("a" + "hello") < ((((("a" + ("" + "\n\t")) + "") + ("foo bar" + "hello")) + "\n\t") + ("hello" + ((((("foo bar" + "") + "hello") + (("foo bar" + "hello") + "foo bar")) + "hello") + "hello")))) {
                        var v163 float64 = (1.5 + 1.5)
                        _ = v163
                    }
                    _ = fi162
                }
                if false {
                    for fi164 := 0; fi164 < 18; fi164++ {
                        switch "a" {
                        case "hello":
                            var v165 int16 = 1
                            _ = v165
                        }
                        v166 := GenS12{F0: 42, F1: 167}
                        _ = fi164
                        _ = v166
                    }
                    var v167 int64 = 985
                    switch 340 {
                    case 4294967295:
                        v167 = v167
                    case 878:
                        v168 := []int{}
                        _ = v168
                    case -744:
                        for fi169 := 0; fi169 < 6; fi169++ {
                            v1 -= (v1 / ((v1) | 1))
                            v2 += 127
                            v170 := GenS13{F0: (^(v167)), F1: 0, F2: (fi161 & 522)}
                            _ = fi169
                            _ = v170
                        }
                    default:
                        var v171 int = -9223372036854775808
                        _ = v171
                    }
                    _ = v167
                }
                var v172 uint64 = (^(v1))
                _ = fi161
                _ = v172
            }
        }
    }
    v1 -= v1
    v3[""] = "\n\t"
    for fi173 := 0; fi173 < 5; fi173++ {
        for fi174 := 0; fi174 < 18; fi174++ {
            var v175 int64 = int64(v2)
            v3[("a" + ("\n\t" + (("hello" + ("\n\t" + "")) + ("\n\t" + "\n\t"))))] = ""
            fmt.Println(fi174, v175)
            _ = fi174
            _ = v175
        }
        switch "\n\t" {
        case "foo bar":
            for fi176 := 0; fi176 < 5; fi176++ {
                v1 += (v1 - v1)
                _ = fi176
            }
        default:
            fmt.Println(v1, v2)
        }
        v177 := v3["hello"]
        _ = fi173
        _ = v177
    }
    v178 := GenS14{F0: 219, F1: 827}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v159
    _ = v178
    _ = v2
    _ = v3
}
