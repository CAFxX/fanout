package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 bool }
type GenS4 struct { F0 uint64; F1 int32; F2 uint }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 string; F1 uint; F2 uint16 }
type GenS9 struct { F0 int32; F1 uint32; F2 uint64 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 int32; F1 uint; F2 int16 }
type GenS12 struct { F0 uint; F1 uint32; F2 uint64 }
type GenS13 struct { F0 uint64; F1 bool; F2 uint }
type GenS14 struct { F0 uint8; F1 uint16 }

func main() {
    v1 := make(map[string]bool)
    var v2 uint8 = 127
    v3 := make(map[uint32]bool)
    var v4 int16 = 165
    v5 := make(map[uint32]bool)
    v6 := GenS1{F0: 1, F1: -2147483648}
    if (-107 < 4294967295) {
        fmt.Println(v4)
        v2 = 1
        if (-2147483648 > 1) {
            _ = len(v1)
            v2 = (v2 + v2)
        }
    } else {
        var v7 uint = 1
        v2 = uint8(v7)
        fmt.Println(v4)
        _ = v7
    }
    for fi8 := 0; fi8 < 17; fi8++ {
        v4 = v4
        for fi9 := 0; fi9 < 4; fi9++ {
            var v10 uint = 42
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    if false {
        v11 := []int{}
        var v12 int16 = -893
        v5[958] = (675 >= 0)
        _ = v11
        _ = v12
    } else {
        v4 -= v4
        for fi13 := 0; fi13 < 13; fi13++ {
            v2 += v2
            _ = fi13
        }
    }
    if false {
        v14 := GenS1{F0: 42, F1: -1}
        switch -1 {
        case 128:
            var v15 int32 = -745
            _ = v15
        default:
            v16 := GenS1{F0: -9223372036854775808, F1: 2147483647}
            _ = v16
        }
        var v17 int32 = -2147483648
        _ = v14
        _ = v17
    } else {
        v18 := GenS1{F0: 0, F1: -853}
        var v19 uint32 = 588
        _ = v18
        _ = v19
    }
    v20 := GenS2{F0: -610}
    delete(v1, ("\n\t" + "\n\t"))
    v21 := GenS2{F0: 9223372036854775807}
    for fi22 := 0; fi22 < 11; fi22++ {
        if (!false) {
            if (true && (!true)) {
                var v23 float64 = 0.0
                v24 := GenS3{F0: "hello", F1: (0 < -2147483648)}
                var v25 uint = 18446744073709551615
                _ = v23
                _ = v24
                _ = v25
            } else {
                var v26 int8 = 127
                switch (("a" + "\n\t") + "\n\t") {
                case "":
                    v3[186] = (-2.25 == 3.14159)
                default:
                    var v27 bool = (v4 != (v4 + 169))
                    _ = v27
                }
                if true {
                    for fi28 := 0; fi28 < 4; fi28++ {
                        v1[("" + "hello")] = false
                        _ = fi28
                    }
                } else {
                    var v29 float64 = 3.14159
                    for fi30 := 0; fi30 < 10; fi30++ {
                        _ = len(v3)
                        if true {
                            v3[621] = (470 != 569)
                            switch -430 {
                            case 9223372036854775807:
                                fmt.Println(fi22)
                            case -754:
                                var v31 int64 = -33
                                _ = v31
                            default:
                                var v32 uint64 = 908
                                _ = v32
                            }
                        } else {
                            var v33 uint64 = 2
                            if false {
                                switch fi30 {
                                case 4294967295:
                                    if true {
                                        v34 := GenS2{F0: -9223372036854775808}
                                        var v35 uint8 = v2
                                        v33 -= v33
                                        _ = v34
                                        _ = v35
                                    } else {
                                        v33 -= (v33 / (((v33 - (v33 % ((1) | 1)))) | 1))
                                        switch "a" {
                                        case "a":
                                            switch (fi30 - fi30) {
                                            case 32767:
                                                var v36 uint = 29
                                                _ = v36
                                            case -972:
                                                var v37 uint8 = 2
                                                _ = v37
                                            }
                                        case "foo bar":
                                            if true {
                                                v33 += v33
                                                if false {
                                                    v29 = (v29 * 0.1)
                                                    v33 = v33
                                                }
                                                v38 := GenS2{F0: fi22}
                                                _ = v38
                                            } else {
                                                v33 += v33
                                                if (fi22 != 806) {
                                                    v4 += v4
                                                    fmt.Println(v2, v26, fi22)
                                                    if false {
                                                        fmt.Println(v33)
                                                    }
                                                }
                                            }
                                        case "hello":
                                            v4 -= v4
                                        }
                                    }
                                case -2147483648:
                                    switch int(v2) {
                                    case 9223372036854775807:
                                        switch fi22 {
                                        case 32767:
                                            if (fi30 != (fi30 ^ fi30)) {
                                                var v39 uint32 = uint32(v2)
                                                if false {
                                                    v40 := []int{(fi22 - fi30)}
                                                    v4 = 0
                                                    _ = v40
                                                } else {
                                                    v4 = -408
                                                }
                                                _ = v39
                                            }
                                        case 128:
                                            delete(v5, 861)
                                        case 21:
                                            var v41 int16 = 1
                                            v42 := GenS4{F0: (v33 / ((v33) | 1)), F1: 408, F2: 0}
                                            _ = v41
                                            _ = v42
                                        default:
                                            if false {
                                                if true {
                                                    v43 := []int{4294967295}
                                                    switch (fi30 - 1) {
                                                    case -174:
                                                        v2 = (v2 | v2)
                                                    }
                                                    _ = v43
                                                }
                                            }
                                        }
                                    case -9223372036854775808:
                                        var v44 string = "\n\t"
                                        _ = v44
                                    case -299:
                                        var v45 uint8 = 42
                                        _ = v45
                                    }
                                case -9223372036854775808:
                                    v46 := GenS1{F0: -9223372036854775808, F1: 209}
                                    _ = v46
                                }
                            }
                            _ = v33
                        }
                        _ = fi30
                    }
                    if (false || false) {
                        v47 := GenS5{F0: v4}
                        for fi48 := 0; fi48 < 5; fi48++ {
                            var v49 uint16 = 349
                            _ = fi48
                            _ = v49
                        }
                        _ = v47
                    } else {
                        for fi50 := 0; fi50 < 16; fi50++ {
                            var v51 int64 = 42
                            fmt.Println(v51, v29, fi50)
                            _ = fi50
                            _ = v51
                        }
                        v52 := []int{}
                        _ = v52
                    }
                    _ = v29
                }
                _ = v26
            }
            var v53 int16 = 1
            v54 := []int{fi22, fi22, 127, (fi22 << ((uint(v2)) & 63))}
            _ = v53
            _ = v54
        }
        v55 := GenS6{F0: 128}
        if false {
            var v56 int32 = -2147483648
            if true {
                v56 = (v56 & (v56 | (v56 / ((0) | 1))))
            } else {
                v57 := GenS3{F0: "", F1: false}
                var v58 int = -592
                _ = v57
                _ = v58
            }
            _ = v56
        }
        _ = fi22
        _ = v55
    }
    switch 42 {
    case 0:
        switch 127 {
        case -266:
            var v59 int8 = int8(v2)
            _ = v59
        case 9223372036854775807:
            if (v2 >= (v2 << ((156) & 63))) {
                delete(v1, "")
                var v60 int32 = 0
                v61 := v3[2]
                _ = v60
                _ = v61
            } else {
                v62 := GenS4{F0: 18446744073709551615, F1: 42, F2: 65535}
                if false {
                    switch (("\n\t" + "foo bar") + "") {
                    case "\n\t":
                        delete(v5, 4294967295)
                    case "":
                        v63 := GenS7{F0: v2}
                        _ = v63
                    case "hello":
                        for fi64 := 0; fi64 < 18; fi64++ {
                            var v65 int = 127
                            var v66 string = "hello"
                            _ = fi64
                            _ = v65
                            _ = v66
                        }
                    default:
                        var v67 uint32 = 2
                        _ = v67
                    }
                }
                for fi68 := 0; fi68 < 2; fi68++ {
                    switch (fi68 >> ((255) & 63)) {
                    case 128:
                        var v69 uint64 = 14
                        _ = v69
                    case 1:
                        v4 -= (v4 % (((v4 & (v4 - (^(v4))))) | 1))
                    }
                    _ = fi68
                }
                _ = v62
            }
        }
    default:
        for fi70 := 0; fi70 < 11; fi70++ {
            var v71 uint64 = 18446744073709551615
            _ = fi70
            _ = v71
        }
    }
    for fi72 := 0; fi72 < 2; fi72++ {
        fmt.Println(fi72, fi72)
        _ = fi72
    }
    if ((4294967295 <= 0) && true) {
        if (!true) {
            var v73 uint16 = 42
            v74 := GenS1{F0: 42, F1: 470}
            _ = v73
            _ = v74
        } else {
            for fi75 := 0; fi75 < 19; fi75++ {
                if (2147483647 == 2) {
                    if true {
                        var v76 float32 = 0.1
                        v77 := GenS6{F0: v2}
                        _ = v76
                        _ = v77
                    } else {
                        _ = len(v5)
                    }
                    v78 := GenS5{F0: v4}
                    v4 = (^(v4))
                    _ = v78
                } else {
                    if true {
                        v79 := GenS5{F0: v4}
                        fmt.Println(v4, v2)
                        _ = v79
                    }
                    for fi80 := 0; fi80 < 18; fi80++ {
                        switch "" {
                        case "a":
                            var v81 float64 = ((float64(1e10) - (((0.0 * float64(float32(-2.25))) / -2.25) + (1e10 - 0.0))) / float64(0.1))
                            _ = v81
                        case "hello":
                            var v82 float64 = 1e10
                            _ = v82
                        case "\n\t":
                            v2 += 2
                        default:
                            v1[("\n\t" + "")] = false
                        }
                        _ = fi80
                    }
                    v83 := []int{(fi75 - 1), (^(fi75)), fi75, (fi75 | fi75)}
                    _ = v83
                }
                v2 = v2
                v2 = v2
                _ = fi75
            }
            v84 := v5[2147483647]
            _ = v84
        }
        v85 := GenS6{F0: v2}
        _ = v85
    } else {
        switch 127 {
        case -9223372036854775808:
            fmt.Println(v2)
        case 2147483647:
            var v86 uint8 = v2
            _ = v86
        case 0:
            if true {
                for fi87 := 0; fi87 < 12; fi87++ {
                    v2 = (v2 / ((v2) | 1))
                    fmt.Println(fi87, v4)
                    if (42 < 42) {
                        if (1.5 == 1e10) {
                            var v88 string = "a"
                            _ = v88
                        } else {
                            for fi89 := 0; fi89 < 19; fi89++ {
                                switch (^(fi89)) {
                                case 9223372036854775807:
                                    v90 := GenS6{F0: v2}
                                    var v91 uint64 = 922
                                    _ = v90
                                    _ = v91
                                case -669:
                                    v1["\n\t"] = true
                                default:
                                    if true {
                                        var v92 uint32 = 1
                                        switch (fi89 & (fi89 & 32767)) {
                                        case 32767:
                                            if true {
                                                _ = len(v3)
                                                v92 -= 1
                                                fmt.Println(v4, v92)
                                            } else {
                                                var v93 uint64 = 608
                                                v92 -= 738
                                                v4 += -32768
                                                v4 -= int16(v2)
                                                v92 += (^(v92))
                                                _ = v93
                                            }
                                        }
                                        if false {
                                            v94 := []int{(fi89 & fi89), (fi87 * fi89), -1, (fi87 >> ((515) & 63))}
                                            if (0 != 4294967295) {
                                                var v95 float64 = (3.14159 * (0.1 / ((1.5 + -2.25) + 3.14159)))
                                                if false {
                                                    var v96 uint64 = 18446744073709551615
                                                    v97 := GenS8{F0: ("a" + "\n\t"), F1: uint(v2), F2: 0}
                                                    _ = v96
                                                    _ = v97
                                                }
                                                if (false || true) {
                                                    v92 = (v92 >> ((42) & 63))
                                                } else {
                                                    delete(v3, (v92 >> ((9) & 63)))
                                                    switch "hello" {
                                                    case "hello":
                                                        if false {
                                                            fmt.Println(v2)
                                                            v92 = (^(v92))
                                                        } else {
                                                            var v98 float64 = float64(1.5)
                                                            v99 := GenS7{F0: v2}
                                                            _ = v98
                                                            _ = v99
                                                        }
                                                    case "":
                                                        fmt.Println(v95, fi87, fi87)
                                                    default:
                                                        var v100 uint32 = v92
                                                        _ = v100
                                                    }
                                                }
                                                _ = v95
                                            }
                                            var v101 float64 = float64((3.14159 + 0.0))
                                            _ = v101
                                            _ = v94
                                        } else {
                                            v102 := GenS9{F0: 0, F1: (v92 >> ((644) & 63)), F2: 4294967295}
                                            fmt.Println(fi87)
                                            var v103 uint32 = v92
                                            var v104 bool = (2147483647 >= -2147483648)
                                            if v104 {
                                                v105 := []int{4294967295, (fi87 - (fi89 | fi89)), 1, fi89}
                                                _ = v105
                                            }
                                            _ = v102
                                            _ = v103
                                            _ = v104
                                        }
                                        _ = v92
                                    }
                                }
                                if ((0 >= 0) && true) {
                                    if (true && (2147483647 <= -1)) {
                                        v4 += (v4 & v4)
                                        fmt.Println(v2, v2)
                                        if (false && true) {
                                            v4 -= v4
                                            switch "foo bar" {
                                            case "foo bar":
                                                if false {
                                                    var v106 float32 = (((0.1 + 1.5) / 1.0) * ((0.0 - 1e10) + float32(0.1)))
                                                    v107 := GenS8{F0: "a", F1: 0, F2: 1}
                                                    _ = v106
                                                    _ = v107
                                                }
                                            case "\n\t":
                                                if true {
                                                    v2 -= (v2 + 2)
                                                    v108 := v1[("a" + "foo bar")]
                                                    v108 = v108
                                                    var v109 int64 = 42
                                                    _ = v108
                                                    _ = v109
                                                } else {
                                                    var v110 uint32 = 65535
                                                    var v111 uint64 = 0
                                                    _ = v110
                                                    _ = v111
                                                }
                                            default:
                                                v112 := GenS2{F0: (fi87 & (-(fi87)))}
                                                _ = v112
                                            }
                                            v4 = (v4 >> ((34) & 63))
                                            v113 := v3[4294967295]
                                            _ = v113
                                        } else {
                                            v2 -= v2
                                        }
                                    } else {
                                        v2 = (^(v2))
                                    }
                                } else {
                                    v2 = v2
                                    if true {
                                        var v114 float32 = (0.0 - 0.0)
                                        var v115 string = "\n\t"
                                        v116 := v5[2147483648]
                                        v2 += 1
                                        v3[833] = (54 <= -9223372036854775808)
                                        v117 := []int{fi87, (fi89 >> ((18446744073709551615) & 63)), (fi87 | fi87)}
                                        v118 := v5[2]
                                        var v119 uint32 = 170
                                        _ = v114
                                        _ = v115
                                        _ = v116
                                        _ = v117
                                        _ = v118
                                        _ = v119
                                    }
                                }
                                v120 := GenS10{F0: 204}
                                _ = fi89
                                _ = v120
                            }
                            v2 = 0
                        }
                    } else {
                        v121 := GenS11{F0: 1, F1: 65535, F2: v4}
                        for fi122 := 0; fi122 < 15; fi122++ {
                            if true {
                                if (false && (-2147483648 <= 1)) {
                                    var v123 uint32 = 2147483647
                                    v124 := GenS7{F0: 42}
                                    v4 = v4
                                    _ = v123
                                    _ = v124
                                } else {
                                    var v125 int = (fi87 % (((fi87 - (fi122 << ((344) & 63)))) | 1))
                                    v126 := GenS8{F0: ("a" + "\n\t"), F1: 42, F2: 1}
                                    _ = v125
                                    _ = v126
                                }
                            } else {
                                var v127 uint8 = (^(v2))
                                fmt.Println(v2)
                                if false {
                                    delete(v1, "foo bar")
                                    v128 := []int{-2147483648, (fi122 << ((18446744073709551615) & 63))}
                                    _ = v128
                                } else {
                                    v129 := GenS1{F0: 9223372036854775807, F1: 0}
                                    _ = v129
                                }
                                _ = v127
                            }
                            switch ("foo bar" + "foo bar") {
                            case "foo bar":
                                if ((!true) && true) {
                                    var v130 uint8 = 127
                                    _ = v130
                                } else {
                                    switch "\n\t" {
                                    case "a":
                                        var v131 int64 = 1
                                        _ = v131
                                    case "hello":
                                        v132 := GenS7{F0: 1}
                                        _ = v132
                                    }
                                }
                            case "":
                                if true {
                                    fmt.Println(v4)
                                }
                            case "a":
                                v133 := []int{fi122}
                                _ = v133
                            default:
                                switch fi87 {
                                case 128:
                                    fmt.Println(fi87, v2)
                                case -507:
                                    if (462 >= 185) {
                                        v134 := GenS12{F0: 129, F1: 2, F2: 4294967295}
                                        v5[8] = (!false)
                                        _ = v134
                                    }
                                    var v135 int64 = 1
                                    switch "hello" {
                                    case "":
                                        var v136 int = fi87
                                        _ = v136
                                    }
                                    _ = v135
                                case -168:
                                    var v137 float32 = float32((float64((-2.25 / ((-2.25 / float32((0.0 - -2.25))) + -2.25))) * 0.1))
                                    _ = v137
                                }
                            }
                            _ = fi122
                        }
                        if false {
                            v2 = v2
                        } else {
                            var v138 float64 = ((((3.14159 - 0.1) + (0.1 + -2.25)) / 1.0) * 0.0)
                            v138 = 0.0
                            switch "a" {
                            case "foo bar":
                                v139 := []int{29, fi87}
                                _ = v139
                            case "hello":
                                v140 := GenS8{F0: "a", F1: 42, F2: 1}
                                _ = v140
                            default:
                                switch fi87 {
                                case 4294967295:
                                    var v141 int16 = 1
                                    _ = v141
                                case -2147483648:
                                    var v142 int16 = v4
                                    _ = v142
                                default:
                                    if (!("" == "\n\t")) {
                                        _ = len(v1)
                                        v138 = v138
                                        var v143 uint64 = 42
                                        _ = v143
                                    } else {
                                        var v144 bool = (!false)
                                        if ("foo bar" < "hello") {
                                            v144 = ((18446744073709551615 <= 375) || ((825 > 138) || (0 > 4294967295)))
                                            v4 -= (v4 & (v4 ^ (v4 + v4)))
                                            var v145 uint16 = 65535
                                            _ = v145
                                        }
                                        switch ("hello" + "a") {
                                        case "":
                                            if v144 {
                                                var v146 uint = 496
                                                _ = v146
                                            } else {
                                                v138 = 1.5
                                                var v147 uint = 1
                                                _ = v147
                                            }
                                        default:
                                            v148 := GenS11{F0: 0, F1: 42, F2: 32767}
                                            _ = v148
                                        }
                                        _ = v144
                                    }
                                }
                            }
                            _ = v138
                        }
                        _ = v121
                    }
                    _ = fi87
                }
                for fi149 := 0; fi149 < 17; fi149++ {
                    switch "hello" {
                    case "":
                        fmt.Println(v4, v2, fi149)
                    }
                    if ((("hello" + "\n\t") <= "foo bar") || (780 != 91)) {
                        v4 -= (v4 / ((v4) | 1))
                        if false {
                            v4 += v4
                            v1[("" + "hello")] = (true || (-2147483648 < 0))
                            for fi150 := 0; fi150 < 10; fi150++ {
                                fmt.Println(v2, fi149)
                                v151 := []int{}
                                _ = fi150
                                _ = v151
                            }
                        } else {
                            for fi152 := 0; fi152 < 1; fi152++ {
                                var v153 uint32 = 42
                                _ = fi152
                                _ = v153
                            }
                            if ((668 > 18446744073709551615) && (-1 <= -1)) {
                                for fi154 := 0; fi154 < 12; fi154++ {
                                    v4 = 543
                                    _ = fi154
                                }
                                switch "foo bar" {
                                case "":
                                    v155 := GenS4{F0: 0, F1: -340, F2: 4294967295}
                                    _ = v155
                                case "\n\t":
                                    v3[2147483648] = (42 < v2)
                                case "foo bar":
                                    fmt.Println(fi149, v4)
                                }
                                for fi156 := 0; fi156 < 7; fi156++ {
                                    fmt.Println(fi149)
                                    delete(v1, ("foo bar" + "\n\t"))
                                    v4 += v4
                                    _ = fi156
                                }
                            }
                            delete(v1, "hello")
                        }
                        var v157 uint16 = 167
                        _ = v157
                    }
                    _ = fi149
                }
            } else {
                for fi158 := 0; fi158 < 4; fi158++ {
                    var v159 string = "foo bar"
                    v160 := GenS13{F0: 4294967295, F1: (!true), F2: 790}
                    v159 = v159
                    _ = fi158
                    _ = v159
                    _ = v160
                }
            }
        default:
            var v161 float64 = -2.25
            _ = v161
        }
    }
    for fi162 := 0; fi162 < 15; fi162++ {
        v163 := GenS5{F0: (v4 * (v4 << ((255) & 63)))}
        v164 := GenS14{F0: v2, F1: 1}
        for fi165 := 0; fi165 < 3; fi165++ {
            if (34 != 4294967295) {
                v5[2] = (-128 == -128)
                var v166 int64 = 42
                _ = v166
            }
            _ = fi165
        }
        _ = fi162
        _ = v163
        _ = v164
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
